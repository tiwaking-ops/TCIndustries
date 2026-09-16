// Phase 10 balance validation report (B3): fork-observable values beside the
// HISTORICAL §33 bands (predecessor GDD, read-only targets — never tune-to).
//
// BAL-001 discipline: this tool OBSERVES and REPORTS. Every row is one of:
//   - OBSERVED — read from the live DB (ledger, missions, snapshots, items)
//   - CODE-CONSTANT — a fork source constant, cited file:line
//   - DERIVED — arithmetic over cited constants (formula shown)
//   - NOT-MEASURABLE — the testbed cannot honestly observe it (say why)
//
// It changes nothing: no numeric tuning, no gameplay writes, no levers.
// Output: docs/phase10_balance_report.md (repo-relative to the testbed root).
//
// Usage: go run ./cmd/phase10report [db-path]   (default swg.db)
package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

// band is one HISTORICAL §33 target; value is the fork's observation.
type row struct {
	area    string
	item    string
	fork    string // fork-observable value or NOT-MEASURABLE reason
	target  string // HISTORICAL band, cited
	note    string // observation note (gap framing, never a fix)
	verdict string // IN-BAND | OUT-OF-BAND | NOT-MEASURABLE | OBSERVED
}

func main() {
	dbPath := "swg.db"
	if len(os.Args) > 1 {
		dbPath = os.Args[1]
	}
	// Same driver as the server (modernc, pure-Go). The tool only ever runs
	// SELECTs; the DB is opened after the suite has shut the server down.
	db, err := sql.Open("sqlite", dbPath)
	must(err, "open DB")
	defer db.Close()

	var rows []row

	// ---------------- Combat: sidearm damage scale (§9.7) ----------------
	// Fork unarmed baseline: handlers/combat.go:27-30 — MinDamage 5,
	// MaxDamage 15, BaseAccuracy 30 ([PROVISIONAL]).
	rows = append(rows,
		row{"Combat §9.7", "unarmed/sidearm damage (start scale)",
			"5–15 per swing (unarmedWeapon, handlers/combat.go:29 [PROVISIONAL])",
			"50–150 start band (predecessor GDD §9.7, HISTORICAL)",
			"fork scale sits BELOW the start band by 10× — a DPS-parity gap is not a failure (BAL-002 lens); recorded as evidence",
			"OUT-OF-BAND"})

	// Creature drop faucet: handlers/combat.go:311-314 — 10 × CLMax capped 200.
	rows = append(rows,
		row{"Economic §12.2.2", "creature-drop cap",
			"10 × CLMax, capped 200 (handlers/combat.go:312-314 [PROVISIONAL])",
			"10–200 band (§12.2.2, HISTORICAL)",
			"fork construction lands inside the band by design",
			"IN-BAND"})

	// Mission rewards (missions.go:183,193,201,214 [PROVISIONAL]):
	// destroy-lair 1500+500×CL; delivery 500×qty (qty 2–4); sample 750;
	// recon 500.
	rows = append(rows,
		row{"Economic §12.2.1", "mission reward band (destroy-lair, CL1)",
			"2000 credits (= 1500 + 500×CL1, missions.go:183 [PROVISIONAL])",
			"500–10,000 band (§12.2.1, HISTORICAL)",
			"low-CL floor sits inside the band",
			"IN-BAND"})
	rows = append(rows,
		row{"Economic §12.2.1", "mission reward band (destroy-lair, CL10)",
			"6500 credits (= 1500 + 500×CL10, derived)",
			"500–10,000 band (§12.2.1, HISTORICAL)",
			"upper-CL construction stays inside the band",
			"IN-BAND"})
	rows = append(rows,
		row{"Economic §12.2.1", "mission reward band (delivery, qty 2–4)",
			"1000–2000 credits (= 500×qty, missions.go:193 [PROVISIONAL])",
			"500–10,000 band (§12.2.1, HISTORICAL)",
			"inside the band",
			"IN-BAND"})
	rows = append(rows,
		row{"Economic §12.2.1", "mission reward band (sample / recon)",
			"750 / 500 credits (missions.go:201,214 [PROVISIONAL])",
			"500–10,000 band (§12.2.1, HISTORICAL)",
			"inside the band",
			"IN-BAND"})

	// Faction thresholds: §15.2.3 gives 2,500/10,000/30,000/75,000. The fork
	// declares alignments with no ranked point thresholds — search proves it.
	// (Declared here from verified source reading, 2026-09-17.)
	rows = append(rows,
		row{"Faction §15.2.3", "ranked faction thresholds",
			"NOT-BUILT: fork has alignment declaration + overt flagging only; no ranked point thresholds exist",
			"2,500 / 10,000 / 30,000 / 75,000 (§15.2.3, HISTORICAL)",
			"recorded: nothing to compare — an unbuilt mechanic, not a gap (ability-system-class residual)",
			"NOT-MEASURABLE"})

	// ---------------- Observed DB rows ------------------------------------
	obs := observe(db)
	rows = append(rows, obs...)

	// ---------------- Progression horizons (§33.3) ------------------------
	rows = append(rows,
		row{"Progression §33.3", "basic mastery 40–60 h / elite 80–120 h / full cap 200–300 h",
			"NOT-MEASURABLE live: XP rates are [PROVISIONAL] flat awards (service 50/25/100, combat 50×CLMax) and fast-cycle testbed cadences distort any wall-clock projection",
			"40–60 h basic / 80–120 h elite / 200–300 h cap (§33.3, HISTORICAL)",
			"honest projection requires the B4 rate redesign (proposal §6); characterized in the B4 options paper, not estimated here",
			"NOT-MEASURABLE"})

	// ---------------- TTK (§9.7) ------------------------------------------
	rows = append(rows,
		row{"Combat §9.7", "PvP 1v1 TTK 30–60 s; solo white-con 20–40 s",
			"NOT-MEASURABLE from this report: TTK is behavior, observable only from live-resolved sparring data (phase10test Proxy 2 records per-swing outcomes at the observed 5–15 dmg scale)",
			"30–60 s PvP / 20–40 s solo (§9.7, HISTORICAL)",
			"at 5–15 dmg vs 1000 health the arithmetic TTK is ~2–3 minutes, far outside 30–60 s — recorded as DERIVED observation, no fix",
			"OUT-OF-BAND (derived)"})

	writeReport(rows)
	writeServiceXPSection(db)
	fmt.Println("Phase 10 balance report written: docs/phase10_balance_report.md")
}

// writeServiceXPSection is the B4 characterization telemetry (Phase 9
// decision 6c): the flat [PROVISIONAL] service-XP sites with citations, and
// the observed per-pool XP totals from character_xp. NO rate is changed and
// none is proposed here — the options paper (docs/phase10_service_xp_options.md)
// holds the redesign shapes with explicit numeric asks.
func writeServiceXPSection(db *sql.DB) {
	var b []byte
	add := func(s string) { b = append(b, s...) }
	add("\n## B4 — Service-XP characterization (Phase 9 decision 6c residual)\n\n")
	add("All service-XP awards in the fork are flat [PROVISIONAL] constants — the\n")
	add("same award regardless of target state, magnitude, or repetition. Sites\n")
	add("(verified 2026-09-17):\n\n")
	add("| Pool | Award | Shape | Site |\n")
	add("|---|---|---|---|\n")
	for _, s := range []struct{ pool, award, shape, site string }{
		{"medical", "50", "per wound-heal action", "services.go:136 (services.HealXP=50)"},
		{"medical", "50", "per buff application", "services.go:190 (services.BuffXP=50)"},
		{"medical", "25", "per stim use", "services.go:427 [PROVISIONAL]"},
		{"medical", "100", "per successful revive", "combat.go:458 (services.ReviveXP=100)"},
		{"entertaining", "25", "per tip received", "services.go:350 (services.TipXP=25)"},
		{"image_designer", "50", "per image-design application", "elite_http.go:327 [PROVISIONAL]"},
		{"image_designer", "50", "per holoemote created", "elite_http.go:367 [PROVISIONAL]"},
		{"smuggler", "50", "per slice accepted", "elite_http.go:512 [PROVISIONAL]"},
		{"(stim pool varies)", "25", "per stim applied", "elite_http.go:613 [PROVISIONAL]"},
		{"scouting", "10", "per survey", "resources.go:265 [PROVISIONAL]"},
		{"scouting", "25", "per sample", "resources.go:347 [PROVISIONAL]"},
		{"scouting", "25", "per corpse harvest", "resources.go:503 [PROVISIONAL]"},
		{"ranger", "10", "per occupied camp per harvest tick", "ranger.go:101 (CampXPPerMember=10)"},
		{"(camp action pool)", "25", "per camp-related action", "ranger.go:176 [PROVISIONAL]"},
		{"combat", "50", "per pet kill assist (flat)", "ranger.go:266 [PROVISIONAL flat]"},
		{"creature_handling", "25/100", "per tame attempt / success", "pets.go:129,153 (constants pets.go:29-30)"},
		{"bio_engineering", "50", "per DNA sample", "pets.go:313 (DNASampleXP=50)"},
		{"mentoring", "50+50", "per mentorship session, both sides", "civic_http2.go:599-600 (civic.MentorshipXP=50)"},
		{"merchant", "price/100", "per sale — unique buyer full, repeat 25%", "economy.go:297-303 (rate-shaped, not flat)"},
		{"structure_crafting", "100", "per house/vendor/base placement", "economy.go:157,196; civic.go:96; faction.go:367"},
		{"combat", "50×CLMax", "per kill (reference rate)", "combat.go:297 [PROVISIONAL]"},
	} {
		add(fmt.Sprintf("| %s | %s | %s | %s |\n", s.pool, s.award, s.shape, s.site))
	}
	add("\nObserved per-pool totals (character_xp, live DB):\n\n")
	add("| XP pool | characters holding XP | total XP |\n")
	add("|---|---|---|\n")
	xrows, err := db.Query(
		`SELECT xp_type, COUNT(DISTINCT character_id), SUM(amount)
		 FROM character_xp GROUP BY xp_type ORDER BY xp_type`)
	if err != nil {
		add("(query failed: " + err.Error() + ")\n")
	} else {
		defer xrows.Close()
		any := false
		for xrows.Next() {
			var pool string
			var n, sum int64
			if err := xrows.Scan(&pool, &n, &sum); err == nil {
				any = true
				add(fmt.Sprintf("| %s | %d | %d |\n", pool, n, sum))
			}
		}
		if !any {
			add("(no XP rows yet — fresh world)\n")
		}
	}
	add("\nCharacterization conclusion: every award above is action-count-driven,\n")
	add("not outcome-driven — repeating the action repeats the XP (the merchant\n")
	add("sale hook is the only rate-shaped exception). This is the flat-rate shape\n")
	add("the Phase 9 decision 6c residual refers to. Any redesign is an option in\n")
	add("docs/phase10_service_xp_options.md and requires explicit numeric\n")
	add("authorization (proposal §6); the default if declined is that these\n")
	add("provisionals stand, characterized by this section.\n")
	_ = os.MkdirAll("docs", 0o755)
	appendOrCreate("docs/phase10_balance_report.md", b)
}

// appendOrCreate appends b to path (the B4 section trails the B3 report).
func appendOrCreate(path string, b []byte) {
	existing, err := os.ReadFile(path)
	if err != nil {
		_ = os.WriteFile(path, b, 0o644)
		return
	}
	_ = os.WriteFile(path, append(existing, b...), 0o644)
}

// observe pulls the OBSERVED rows from the live DB (read-only).
func observe(db *sql.DB) []row {
	var rows []row

	// Snapshot chain: latest economic snapshot (faucet/sink/net).
	var circ, f30, s30 int64
	var net float64
	var snapTime string
	err := db.QueryRow(
		`SELECT circulation, faucet_30d, sink_30d, net_pct, taken_at
		 FROM economic_snapshots ORDER BY id DESC LIMIT 1`,
	).Scan(&circ, &f30, &s30, &net, &snapTime)
	switch {
	case err == sql.ErrNoRows:
		rows = append(rows, row{"Economic §12.4", "30-day net inflation window",
			"OBSERVED: no economic_snapshots row yet (fresh world; the Phase 10 job writes these on its cadence)",
			"~5–10% annualized (§12.4, HISTORICAL)",
			"fresh testbed cannot produce a 30-day window; observed when a window exists",
			"NOT-MEASURABLE"})
	case err != nil:
		fmt.Fprintf(os.Stderr, "snapshot query failed: %v\n", err)
	default:
		rows = append(rows, row{"Economic §12.4", "30-day net inflation window",
			fmt.Sprintf("OBSERVED: circulation=%d faucet30d=%d sink30d=%d net=%.2f%% (snapshot at %s)", circ, f30, s30, net, snapTime),
			"~5–10% annualized (§12.4, HISTORICAL)",
			"testbed windows are hours old, not 30 days — read as pipeline proof, not as an inflation measurement",
			"OBSERVED"})
	}

	// Wealth distribution (diagnostic Gini; §34.3 — never target-capped).
	total, players, gini := wealth(db)
	if players > 0 {
		rows = append(rows, row{"Wealth §34.3", "credit distribution (diagnostic Gini)",
			fmt.Sprintf("OBSERVED: players=%d total_credits=%d gini=%.3f", players, total, gini),
			"no §33 target — Gini is diagnostic only (§34.3)",
			"tracked, never target-capped, never a lever",
			"OBSERVED"})
	}

	// Maintenance share of weekly sink (maintenance_fee vs all sinks, 7d).
	mFee, allSink := maintenanceShare(db)
	if allSink > 0 {
		share := float64(mFee) / float64(allSink) * 100
		rows = append(rows, row{"Economic §33.4", "maintenance share of weekly income/sink",
			fmt.Sprintf("OBSERVED: maintenance_fee=%d of %d total sink over trailing 7d (%.1f%%)", mFee, allSink, share),
			"5–15% of weekly income (§33.4, HISTORICAL)",
			"fresh-world windows are dominated by training sinks; the share is recorded, not judged",
			"OBSERVED"})
	}

	// Inventory crafted share (§34 player-driven economy, raw counts).
	cTotal, cWith := craftedShare(db)
	if cTotal > 0 {
		pct := float64(cWith) / float64(cTotal) * 100
		rows = append(rows, row{"Economy §34", "crafted provenance share",
			fmt.Sprintf("OBSERVED: %d/%d items carry schematic provenance (%.0f%%)", cWith, cTotal, pct),
			"≥95% crafted (§1.3/§34, HISTORICAL)",
			"raw counts beside the percentage (small-cohort rule)",
			"OBSERVED"})
	}

	return rows
}

func wealth(db *sql.DB) (int64, int, float64) {
	rowsx, err := db.Query(`SELECT credits FROM characters WHERE credits >= 0 ORDER BY credits ASC`)
	if err != nil {
		return 0, 0, 0
	}
	defer rowsx.Close()
	var vals []int64
	for rowsx.Next() {
		var v int64
		_ = rowsx.Scan(&v)
		vals = append(vals, v)
	}
	n := len(vals)
	if n == 0 {
		return 0, 0, 0
	}
	var sum, weighted int64
	for i, v := range vals {
		sum += v
		weighted += int64(i+1) * v
	}
	if sum == 0 {
		return 0, n, 0
	}
	g := (2*float64(weighted))/(float64(n)*float64(sum)) - (float64(n)+1)/float64(n)
	if g < 0 {
		g = 0
	}
	if g > 1 {
		g = 1
	}
	return sum, n, g
}

func maintenanceShare(db *sql.DB) (int64, int64) {
	var mFee, allSink int64
	_ = db.QueryRow(
		`SELECT COALESCE(SUM(CASE WHEN category='maintenance_fee' THEN -amount ELSE 0 END),0),
		        COALESCE(SUM(CASE WHEN flow='sink' THEN -amount ELSE 0 END),0)
		 FROM ledger_entries
		 WHERE created_at >= datetime('now','-7 days')`).Scan(&mFee, &allSink)
	return mFee, allSink
}

func craftedShare(db *sql.DB) (int, int) {
	var total, with int
	_ = db.QueryRow(
		`SELECT COUNT(*), COALESCE(SUM(CASE WHEN schematic_id != '' THEN 1 ELSE 0 END),0)
		 FROM crafted_items`).Scan(&total, &with)
	return total, with
}

func writeReport(rows []row) {
	_ = os.MkdirAll("docs", 0o755)
	var b []byte
	add := func(s string) { b = append(b, s...) }
	add("# Phase 10 Balance Validation Report (B3)\n\n")
	add("Generated: " + time.Now().UTC().Format(time.RFC3339) + "\n\n")
	add("Author: Buffy — Codebuff agent, Phase 10 implementation (proposal v0.2 §3/B3).\n\n")
	add("> **Status discipline:** every fork value below is OBSERVED from the live DB,\n")
	add("> read from fork source with a file:line citation, DERIVED with the formula\n")
	add("> shown, or honestly NOT-MEASURABLE. The HISTORICAL §33 bands are targets to\n")
	add("> validate against, never tune-to (BAL-001; Master GDD \"Inherited from prior\n")
	add("> versions\" item 3). A gap recorded here is EVIDENCE, not a work item: any fix\n")
	add("> is a separate numeric-authorization decision (proposal §5.1).\n\n")
	add("| Area | Item | Fork value | HISTORICAL target | Verdict | Note |\n")
	add("|---|---|---|---|---|---|\n")
	for _, r := range rows {
		note := r.note
		if len(note) > 0 {
			note = escapePipe(note)
		}
		add(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
			r.area, r.item, escapePipe(r.fork), escapePipe(r.target), r.verdict, note))
	}
	add("\n## Reading this report\n\n")
	add("- **BAL-002 lens:** a DPS-parity gap is not a failure; a collapsed profession\n  category is. Out-of-band rows are recorded observations.\n")
	add("- **NOT-MEASURABLE rows are honest exits** (§2.1 disjunction): instrumentation\n  plus an honest failure signal satisfies the Phase 10 exit criteria.\n")
	add("- **No numeric changes** were made by or proposed inside this report.\n")
	must(os.WriteFile("docs/phase10_balance_report.md", b, 0o644), "write report")
}

func escapePipe(s string) string {
	out := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			out = append(out, '\\', '|')
		} else {
			out = append(out, s[i])
		}
	}
	return string(out)
}

func must(err error, what string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "phase10report: %s: %v\n", what, err)
		os.Exit(1)
	}
}
