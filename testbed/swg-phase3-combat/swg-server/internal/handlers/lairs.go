// Lair handlers for Phase 9: lair damage via combat_action, population
// regeneration, and destroyed-lair relocation (GDD 9.5.3/17.2.3 — GDD-given
// mechanics the fork never built). Testbed fork; generic content only.
//
// Regen/relocation also close the root cause of the Phase 6 shared-DB
// depletion flake (recorded as evidence in HYGIENE_NOTE.md, not scope creep:
// Destroy Lair missions and taming both require population accounting).
package handlers

import (
	"log"
	"math/rand"
	"os"
	"time"

	"swg-server/internal/combat"
	"swg-server/internal/creatures"
	"swg-server/internal/protocol"
	"swg-server/internal/world"
)

// LairDestroyXP: bonus XP for destroying a lair (GDD 7.2.1 "Destroying lairs
// (bonus XP)" — amount GDD-silent). [PROVISIONAL]
const LairDestroyXP = 150

// LairRelocationWindow: destroyed lairs respawn elsewhere after this long
// (GDD 17.2.3/5.4.2: 12–24 h window; lower bound taken, flagged).
// Fast-cycle compresses to minutes so relocation executes live in tests.
var LairRelocationWindow = 12 * time.Hour

// LairRelocateModifier is the SQLite datetime modifier twin of the window.
var LairRelocateModifier = "-12 hours"

func init() {
	if os.Getenv("TESTBED_FAST_CYCLE") == "1" {
		LairRelocationWindow = 10 * time.Minute
		LairRelocateModifier = "-10 minutes"
	}
}

// handleLairAttack resolves one accepted lair hit (base-attack precedent:
// skill + range + conscious gates already apply; no faction/window).
func (h *WorldHandler) handleLairAttack(client *Client, lairID string) {
	lair, err := h.db.GetLair(lairID)
	if err != nil {
		h.sendError(client, "lair unknown")
		return
	}
	if lair.DestroyedAt.Valid {
		h.sendError(client, "lair already destroyed")
		return
	}
	if lair.Zone != client.Pos.Planet {
		h.sendError(client, "lair not in this zone")
		return
	}
	dx := client.Pos.X - lair.PosX
	dz := client.Pos.Z - lair.PosZ
	if dx*dx+dz*dz > AttackRangeM*AttackRangeM {
		h.sendError(client, "lair out of range")
		return
	}
	attacker := combat.Attacker{
		AccuracySkill: 0, DamageBonus: UnarmedDamageBonus,
		Posture: combat.PostureStanding, Stance: combat.StanceNormal,
	}
	if st, err := h.db.GetCombatState(client.CharacterID); err == nil {
		attacker.Posture = combat.Posture(st.Posture)
		attacker.Stance = combat.Stance(st.Stance)
	}
	weapon := h.attackerWeapon(client)
	// Lairs do not dodge (same dummy-defender shape as bases — flagged).
	defender := combat.Defender{
		DefenseSkill: 0, Posture: combat.PostureStanding,
		ArmorMitigationPct: 0, ResistanceMultiplier: 1.0,
	}
	result := combat.ResolveAttack(attacker, defender, weapon,
		rand.New(rand.NewSource(time.Now().UnixNano())))
	crMsg := protocol.CombatResultMsg{
		AttackerID: client.CharacterID, TargetID: lair.ID,
		Hit: result.Hit, Damage: result.Damage, HitChance: result.HitChance,
	}
	h.broadcastToInterested(client, protocol.MsgCombatResult, crMsg)
	h.send(client, protocol.MsgCombatResult, crMsg)
	if !result.Hit {
		return
	}
	lair.LairHP -= result.Damage
	if lair.LairHP < 0 {
		lair.LairHP = 0
	}
	_ = h.db.UpdateLairHP(lair.ID, lair.LairHP)
	_ = h.db.LogLairDamage(lair.ID, client.CharacterID, result.Damage, time.Now().Unix())
	if lair.LairHP <= 0 {
		_ = h.db.DestroyLair(lair.ID)
		_ = h.db.AddCharacterXP(client.CharacterID, "combat", LairDestroyXP)
		h.broadcastToInterested(client, protocol.MsgLairDestroyed,
			protocol.LairDestroyedMsg{LairID: lair.ID, Destroyed: true})
		h.send(client, protocol.MsgLairDestroyed,
			protocol.LairDestroyedMsg{LairID: lair.ID, Destroyed: true})
	}
}

// tickLairs runs each harvest interval: population regen to max for standing
// lairs (GDD 17.2.3 — keeps farming viable), plus destroyed-lair relocation.
// Regen rate is PROVISIONAL +1 instance per harvest tick while living < max
// (incremental shape; full-repopulation alternative stays TBD). The 30-minute
// production tick (HISTORICAL-USER input, 2026-09-16) is deliberately NOT
// applied in the testbed per owner direction recorded in the Phase 9 proposal
// §10 amendment — correct for production, unusable for the test.
func (h *WorldHandler) tickLairs() {
	all, err := h.db.AllLairs()
	if err != nil {
		log.Printf("lair tick: list failed: %v", err)
		return
	}
	for i := range all {
		l := &all[i]
		if l.DestroyedAt.Valid {
			continue
		}
		living, err := h.db.LivingCountForLair(l.ID)
		if err != nil {
			continue
		}
		if living >= l.MaxPopulation {
			continue
		}
		tmpl := creatures.TemplateByID(l.TemplateID)
		if tmpl == nil {
			continue
		}
		id, err := h.db.SpawnLairInstance(l.ID, l.Zone, l.TemplateID, l.PosX, l.PosZ, tmpl.HealthMax)
		if err != nil {
			continue
		}
		h.mu.Lock()
		h.liveCreatures[id] = &LiveCreature{
			ID: id, TemplateID: l.TemplateID, Name: tmpl.Name,
			Zone: l.Zone, PosX: l.PosX, PosZ: l.PosZ,
		}
		h.mu.Unlock()
		h.grid.AddEntity(&world.Entity{
			CharacterID: id, Name: tmpl.Name, Species: "creature",
			Pos: world.Position{Planet: l.Zone, X: l.PosX, Y: 5.0, Z: l.PosZ},
		})
	}
	h.tickLairRelocation()
}

// tickLairRelocation respawns long-destroyed lairs elsewhere (GDD 17.2.3:
// "a new lair (of a valid type for the biome) spawns elsewhere on the planet").
// Single zone: template pool is the starter bestiary (flagged); placement is
// random outside active cities with structure separation (flagged).
func (h *WorldHandler) tickLairRelocation() {
	due, err := h.db.RelocatableLairs(LairRelocateModifier)
	if err != nil {
		return
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range due {
		l := &due[i]
		var x, z float64
		placed := false
		for attempt := 0; attempt < 10 && !placed; attempt++ {
			x = (rng.Float64()*2 - 1) * 4000
			z = (rng.Float64()*2 - 1) * 4000
			if _, cerr := h.db.CityContaining(l.Zone, x, z); cerr == nil {
				continue
			}
			if nerr := h.checkNoBuild(l.Zone, x, z); nerr != nil {
				continue
			}
			placed = true
		}
		if !placed {
			continue
		}
		tmpl := creatures.TemplateByID(l.TemplateID)
		hpMax, maxPop := 200, 3
		if tmpl != nil {
			hpMax = 100 * tmpl.CLMax
			if hpMax < 200 {
				hpMax = 200
			}
		}
		if _, err := h.db.InsertLair(l.TemplateID, l.Zone, x, z, hpMax, maxPop); err != nil {
			continue
		}
		_ = h.db.DeleteLair(l.ID)
		log.Printf("lair %s relocated to (%.0f, %.0f)", l.ID, x, z)
	}
}
