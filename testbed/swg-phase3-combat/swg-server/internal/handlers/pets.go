// Pet handlers for Phase 9: taming, pet commands, pet ticks, DNA/tissue
// sampling, and DNA combining (GDD 17.2.4; Bio-Engineer fuller-DNA scope per
// the §9.3 decision). Testbed fork; generic content only. GDD-silent numbers
// [PROVISIONAL] per the owner-approved convention; GDD-given values cited.
package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"swg-server/internal/creatures"
	"swg-server/internal/protocol"
	"swg-server/internal/world"
)

// Pet tuning [PROVISIONAL where GDD-silent; GDD-given cited].
const (
	// TameRangeM: taming/DNA interaction range (8 m melee-gate basis).
	TameRangeM = 8.0
	// PetFollowRadiusM: follow engages outside this radius (flagged).
	PetFollowRadiusM = 5.0
	// PetStepM: follow step per 1 s tick (flagged; matches client move steps).
	PetStepM = 4.0
	// PetSightM: attack reach from the pet (8 m gate basis).
	PetSightM = 8.0
	// TameXP: creature_handling XP per attempt/success [PROVISIONAL].
	TameXPAttempt = 25
	TameXPSuccess = 100
	// DNASampleXP: bio_engineering XP per sample [PROVISIONAL].
	DNASampleXP = 50
)

// tameCLCaps maps Command-tree tiers to CL caps (GDD 8.3.7: up to CL 20 at
// Master). [PROVISIONAL distribution]
var tameCLCaps = [5]int{3, 8, 12, 16, 20}

// petCapByTier maps Command tiers to simultaneous pets (GDD 8.3.7: up to 3 at
// Master). [PROVISIONAL distribution]
var petCapByTier = [5]int{1, 1, 2, 2, 3}

// PetRuntime mirrors a pet for ticks + visibility (liveCreatures precedent).
type PetRuntime struct {
	ID         string
	OwnerID    string
	TemplateID string
	Name       string
	Quality    int
	Zone       string
	PosX, PosZ float64
}

func petsMap(h *WorldHandler) map[string]*PetRuntime {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.livePets
}

// SeedPhase9World prepares Phase 9 persistence. Idempotent.
func (h *WorldHandler) SeedPhase9World() error {
	if err := h.db.EnsurePhase9Schema(); err != nil {
		return err
	}
	log.Printf("Phase 9 world seeded: missions, contracts, pets, camps")
	return nil
}

// --- Taming (WS positional) ---

func (h *WorldHandler) handleTameCreature(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var tm protocol.TameMsg
	if err := json.Unmarshal(dataBytes, &tm); err != nil {
		h.sendError(client, "invalid tame data")
		return
	}
	if !h.hasBox(client.CharacterID, "creaturehandler_novice") {
		h.sendError(client, "requires Novice Creature Handler")
		return
	}
	target, ok := h.getLiveCreature(tm.TargetID)
	if !ok {
		h.sendError(client, "target not found")
		return
	}
	tmpl := creatures.TemplateByID(target.TemplateID)
	if tmpl == nil || !tmpl.Tamable {
		h.sendError(client, "target not tamable")
		return
	}
	cmdTier := h.skillTier(client.CharacterID, "creaturehandler_command_")
	cap := tameCLCaps[4]
	if cmdTier >= 0 && cmdTier <= 4 {
		cap = tameCLCaps[cmdTier]
	}
	if tmpl.CLMax > cap {
		h.sendError(client, "target too difficult to command")
		return
	}
	dx := client.Pos.X - target.PosX
	dz := client.Pos.Z - target.PosZ
	if target.Zone != client.Pos.Planet ||
		dx*dx+dz*dz > TameRangeM*TameRangeM {
		h.sendError(client, "target out of range")
		return
	}
	// Strict reading of 17.5: blocked while aggro'd (at anyone — the tame
	// interaction must not interrupt combat).
	if h.creatureTarget(target.ID) != "" {
		h.sendError(client, "target is fighting")
		return
	}
	owned, _ := h.db.PetsOf(client.CharacterID)
	maxPets := petCapByTier[4]
	if cmdTier >= 0 && cmdTier <= 4 {
		maxPets = petCapByTier[cmdTier]
	}
	if len(owned) >= maxPets {
		h.sendError(client, "pet cap reached")
		return
	}
	_ = h.db.AddCharacterXP(client.CharacterID, "creature_handling", TameXPAttempt)
	// Success 50% + 10%/Command-tier (provisional); fail aggros tamer
	// (CONFIRMED risk rule — taming danger is meaningful).
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	chance := 50 + 10*cmdTier
	if rng.Intn(100) >= chance {
		h.setCreatureTarget(target.ID, client.CharacterID)
		h.sendError(client, "taming failed — the creature turns on you")
		return
	}
	// Convert: remove from the wild (no corpse — taming isn't killing, flagged)
	// and bind a pet. Quality rolls the corpse band (flagged reuse).
	quality := 300 + rng.Intn(501)
	name := tm.TargetID
	if tm.Name != "" && len(tm.Name) <= 40 {
		name = tm.Name
	}
	petID, err := h.db.CreatePet(client.CharacterID, target.TemplateID, name,
		quality, "", client.Pos.Planet, target.PosX, target.PosZ)
	if err != nil {
		h.sendError(client, "taming failed")
		return
	}
	h.removeLiveCreature(target.ID)
	_ = h.db.AddCharacterXP(client.CharacterID, "creature_handling", TameXPSuccess)
	h.mu.Lock()
	if h.livePets == nil {
		h.livePets = make(map[string]*PetRuntime)
	}
	h.livePets[petID] = &PetRuntime{
		ID: petID, OwnerID: client.CharacterID, TemplateID: target.TemplateID,
		Name: name, Quality: quality, Zone: client.Pos.Planet,
		PosX: target.PosX, PosZ: target.PosZ,
	}
	h.mu.Unlock()
	h.grid.AddEntity(&world.Entity{
		CharacterID: petID, Name: name, Species: "pet",
		Pos: world.Position{Planet: client.Pos.Planet, X: target.PosX, Y: 5.0, Z: target.PosZ},
	})
	h.send(client, protocol.MsgTamed, protocol.TamedMsg{PetID: petID, TemplateID: target.TemplateID})
}

// addPetWorld registers a pet in the runtime mirror + grid (tame/combine).
func (h *WorldHandler) addPetWorld(petID, ownerID, templateID, name string, quality int, zone string, x, z float64) {
	h.mu.Lock()
	if h.livePets == nil {
		h.livePets = make(map[string]*PetRuntime)
	}
	h.livePets[petID] = &PetRuntime{
		ID: petID, OwnerID: ownerID, TemplateID: templateID,
		Name: name, Quality: quality, Zone: zone, PosX: x, PosZ: z,
	}
	h.mu.Unlock()
	h.grid.AddEntity(&world.Entity{
		CharacterID: petID, Name: name, Species: "pet",
		Pos: world.Position{Planet: zone, X: x, Y: 5.0, Z: z},
	})
}

// removePetWorld unregisters a pet (release path).
func (h *WorldHandler) removePetWorld(petID string) {
	h.mu.Lock()
	delete(h.livePets, petID)
	h.mu.Unlock()
	h.grid.RemoveEntity(petID, world.Position{})
}

// creatureTarget reads a live instance's current target (aggression check).
func (h *WorldHandler) creatureTarget(instanceID string) string {
	rows, err := h.db.GetLivingCreatureInstances()
	if err != nil {
		return ""
	}
	for i := range rows {
		if rows[i].ID == instanceID {
			return rows[i].TargetCharacterID
		}
	}
	return ""
}

// setCreatureTarget forces aggro (fail-aggro rule + interrupt helper).
func (h *WorldHandler) setCreatureTarget(instanceID, charID string) {
	rows, err := h.db.GetLivingCreatureInstances()
	if err != nil {
		return
	}
	for i := range rows {
		if rows[i].ID == instanceID {
			rows[i].TargetCharacterID = charID
			rows[i].State = "aggro"
			_ = h.db.UpdateCreatureInstance(&rows[i])
			return
		}
	}
}

// removeLiveCreature deletes a wild instance without a corpse (taming path).
func (h *WorldHandler) removeLiveCreature(instanceID string) {
	h.mu.Lock()
	delete(h.liveCreatures, instanceID)
	h.mu.Unlock()
	h.grid.RemoveEntity(instanceID, world.Position{})
	_ = h.db.DeleteCreatureInstance(instanceID)
}

// --- DNA + tissue sampling (WS positional, 8 m uniform) ---

func (h *WorldHandler) handleDNASample(client *Client, raw []byte) {
	if client.CharacterID == "" {
		h.sendError(client, "not in world")
		return
	}
	var msg protocol.WSMessage
	json.Unmarshal(raw, &msg)
	dataBytes, _ := json.Marshal(msg.Data)
	var dm protocol.DNASampleMsg
	if err := json.Unmarshal(dataBytes, &dm); err != nil {
		h.sendError(client, "invalid dna data")
		return
	}
	if !h.hasBox(client.CharacterID, "bioengineer_novice") {
		h.sendError(client, "requires Novice Bio-Engineer")
		return
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	quality := 300 + rng.Intn(501)
	templateID := ""
	// Owned pets sample anytime, in range (17.2.4 "tamed" branch).
	if pet, err := h.db.GetPet(dm.TargetID); err == nil {
		if pet.OwnerID != client.CharacterID {
			h.sendError(client, "not your pet")
			return
		}
		dx := client.Pos.X - pet.PosX
		dz := client.Pos.Z - pet.PosZ
		if pet.Zone != client.Pos.Planet || dx*dx+dz*dz > TameRangeM*TameRangeM {
			h.sendError(client, "pet out of range")
			return
		}
		quality = pet.Quality
		templateID = pet.TemplateID
	} else {
		// Wild branch: in-window corpses only (17.2.4 "harvestable window").
		corpses, err := h.db.HarvestableCorpses(
			client.Pos.Planet, client.Pos.X, client.Pos.Z, TameRangeM, time.Now())
		if err != nil {
			h.sendError(client, "sample failed")
			return
		}
		for i := range corpses {
			if corpses[i].InstanceID == dm.TargetID {
				templateID = corpses[i].TemplateID
				break
			}
		}
		if templateID == "" {
			h.sendError(client, "no sampleable corpse in range")
			return
		}
	}
	// DNA rides crafted_items (schematic dna_sample): mail/vendor/claim
	// machinery applies with no new tables (flagged reuse). Template serials
	// are numeric, so the template rides a numeric stat (flagged).
	templateNum := float64(templateSerial(templateID))
	var issued []string
	dnaID, err := h.db.InsertItem(client.CharacterID, "dna_sample",
		"DNA Sample ("+templateID+")",
		map[string]float64{"dna_template": templateNum, "dna_quality": float64(quality)})
	if err != nil {
		h.sendError(client, "sampling failed")
		return
	}
	issued = append(issued, dnaID)
	// Tissue Engineering tiers additionally yield a tissue sample (tradeable
	// resource for elite armor slots).
	if h.skillTier(client.CharacterID, "bioengineer_tissue_engineering_") >= 1 {
		tissueID, err := h.db.InsertItem(client.CharacterID, "tissue_sample",
			"Tissue Sample ("+templateID+")",
			map[string]float64{"tissue_template": templateNum, "tissue_quality": float64(quality)})
		if err == nil {
			issued = append(issued, tissueID)
		}
	}
	_ = h.db.AddCharacterXP(client.CharacterID, "bio_engineering", DNASampleXP)
	h.send(client, protocol.MsgSampled, protocol.DNASampledMsg{ItemIDs: issued})
}

// templateSerial parses numeric template serials ("0001" → 1).
func templateSerial(id string) int {
	n := 0
	for _, c := range id {
		if c < '0' || c > '9' {
			return 0
		}
		n = n*10 + int(c-'0')
	}
	return n
}
