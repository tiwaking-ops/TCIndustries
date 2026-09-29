# TCIndustries — Salvage, Destruction & Contents Handling — Structural Proposal

**Filename:** `TCIndustries_6d_Salvage_Destruction_Contents_Proposal_v0.1_2026-09-20.md`
**Status:** PROPOSED. No authority. Ruling-ready candidate only. Nothing here is
LOCKED, and no element of this proposal changes any GDD status.
**Author:** Buffy — Codebuff agent (Freebuff), autonomous-loop drafter, Cycle 14.
**Assessor:** None — single-session authoring; no independent review pass (per GR-004, recorded as such rather than guessed).
**Date:** 2026-09-20
**Queue position:** designated Cycle-14 next draft to restore 5/5 (backlog priority 6d; see `proposals/TCIndustries_Continuous_Queue_Tracker.md`).
**Pre-flight:** the salvage/destruction/contents surface is cited as a newly-surfaced adjacent gap under HD-ITM-01 (Option C — no baseline decay; destruction/loss is now the recurring item-destruction source) and HD-BLD-02 (zero-credit structure removal; contents-handling on removal is open question 6e(a)). No prior proposal or investigation structures this surface as a decision. The adjacent OQ-010 ruling is recorded; the adjacent BLD-004 ruling is recorded; no prior artifact bundles these surfaces into one composition. No duplicate pass. The 2026-09-17 duplicate-pass rule is respected — pre-flight grep of `proposals/` + `investigations/` + `governance/project_memory.md` for "salvage", "destruction", "contents-handling", "6d" returned zero prior proposal/investigation hits.

---

## Status wall

- Every option below is **PROPOSED** and requires an explicit HD ruling to advance.
- **Zero numeric values.** No salvage fractions, destruction-probabilities, contents-loss curves, or recovery windows appear anywhere (BAL-001).
- **Composition note (recorded, not authority):** this proposal was designated next-draft at Cycle 13 specifically because two already-ruled items make it load-bearing immediately: HD-ITM-01 makes destruction/loss the *only* recurring source of item destruction, and HD-BLD-02 makes zero-credit removal a structure-removal instrument whose contents-handling surface is an open question (6e(a)). This proposal treats the two as one composing draft where the *ownership/salvage* surface is shared. It does **not** re-propose either HD-ITM-01 (decay structure) or HD-BLD-02 (abandonment instrument) — both are already ruled.
- **Non-canonical list guard:** "Fully tradeable by default / no account or character soul-binding" is on the non-canonical list (Open Questions Register §6; GDD §36). No option assumes tradeability rules; no option defaults to soul-binding as the protection mechanism. Any option that would touch transfer/ownership tooling does so only as a *surfaces-for-review* reference to SAFE-001's still-open mechanism surface, not as a tradeability ruling.
- **SAFE-001 is HUMAN-LOCKED principle-only (HD-GDD-01):** this proposal is part of that mechanism surface — it structures the *structural handling* of destruction/loss/removal for owner ruling; it does not authorise the safeguard mechanism in full (transfer mechanics, ownership tooling, fraud protections remain open).
- **INV-010 (approved, conditional HD-ITM-01):** clauses (c) and (d) bind on any destruction/loss system as ownership/meaningful-recourse obligations; clause (b) binds the provenance requirements of any repair/restoration mechanism; clause (a) attaches to whatever systems generate repair demand in the ruled no-decay model. This proposal's options must be read against that approved conditional invariant.
- **INV-BLD-004b (approved HD-BLD-02):** no ghost equilibrium — zero-sustenance settlements converge to non-functional state. This proposal's removal-contents handling must not quietly violate that approved invariant by making zero-credit removal a lossless escape hatch from ownership consequences.
- **PROVENANCE ENGINE IS DEFERRED (HD-EIC-08):** provenance *depth/visibility requirements* are lawful to state (INV-017 candidate, INV-016a approved, INV-010(b) approved) but no provenance *engine* design is proposed or implied.

---

## Anchors (controlling references)

| Anchor | Status | What it locks/constrains |
|---|---|---|
| HD-ITM-01 / OQ-010 (decision record `proposals/TCIndustries_OQ-010_Ruling_Record_HD-ITM-01_2026-09-19.md`) | HUMAN-LOCKED (durability structure only) | No baseline decay; destruction/loss/obsolescence/upgrade are the item-destruction sources; INV-010 approved conditional (clauses a-d) |
| HD-BLD-02 / BLD-004 (decision record `proposals/TCIndustries_BLD-004_Ruling_Record_HD-BLD-02_2026-09-19.md`) | HUMAN-LOCKED (abandonment handling only) | Zero-credit structure removal is the ruled instrument; contents-handling on removal is open (backlog 6e(a)); INV-BLD-004b approved |
| SAFE-001 (GDD §26; HUMAN-LOCKED principle-only, HD-GDD-01) | HUMAN-LOCKED principle-only | Safeguards against fraud, abandoned assets, exploitative transfer — this proposal is part of that mechanism surface (mechanism still open). |
| SAFE-002 (GDD §26; PROPOSED/TBD) | PROPOSED/TBD | Exploit policies require explicit design — destruction/removal edge cases (race conditions on removal, contents appropriation) are an implementation-policy surface the OQ-012 ruling would inform, not a separate path. |
| SAFE-003 (GDD §26; PROPOSED) | PROPOSED | Economic observability — destruction/loss/removal events are observable economic events (relevant to INV-009a surface and any future observability requirements brief T-07, unapproved). |
| PLR-003 (GDD §9; HUMAN-LOCKED principle-only, HD-GDD-01) | HUMAN-LOCKED principle-only | Meaningful persistent ownership — destruction/loss must not make ownership meaningless; this is INV-010(c)'s subject and binds structure, not numerics. |
| CRFT-001 (GDD §12; HUMAN-LOCKED principle-only, HD-GDD-01) | HUMAN-LOCKED principle-only | Differentiated products — provenance-bearing items may be destroyed; their provenance record status on destruction is a design surface, not a settled rule. |
| ITM-002 (GDD §14) | PROPOSED/TBD | Durability/decay/repair model TBD — ruled by HD-ITM-01 at structure layer; this proposal covers the *adjacent destruction/loss* surface, not the (ruled) decay question. |
| ITM-001 (GDD §13) | PROPOSED | Item identity/provenance — destruction/loss/removal must state what happens to that identity record (open). |
| BLD-004 (GDD §19) | TBD (partial) | Abandonment handling now ruled (HD-BLD-02); placement/lots/density remain TBD; contents-handling on removal is the composing open surface. |
| INV-010 (Invariant Register v0.1, OQ-010 entry; APPROVED conditional, HD-ITM-01) | APPROVED (conditional) | Binding on any destruction/loss system: (c) ownership remains meaningful; (d) no irreplaceable provenance-bearing asset destroyed without recourse; (a) repair demand attaches to destruction/loss; (b) provenance survives repair. |
| INV-BLD-004b (new candidate authored in the HD-BLD-02 proposal; APPROVED, HD-BLD-02) | APPROVED | No ghost equilibrium — satisfied by construction under zero-credit removal; contents handling must not void this. |
| INV-017 (Invariant Register v0.1, OQ-017 entry; candidate) | Candidate | Provenance at three granularities (creator/org/resource-origin); visibility decision-relevant — referenced as a requirement-shape for what provenance does on destruction, not as a ruling. |
| Backlog 6e (from HD-BLD-02 §1.1(d); recorded, open) | Open | (a) contents handling on zero-credit removal; (b) settlement-institution dissolution rule — 6e(a) is the composition point this proposal folds in as one surface; 6e(b) is named but deferred in §"What this does not resolve". |
| Non-canonical list (Open Questions Register §6; GDD §36) | Confirmed boundary | Single unified currency as default excluded; soul-binding default excluded; etc. — relevant to any transfer/ownership default an option might imply. |
| T-07 (Gap Analysis §7; NOT STARTED, not approved) | — | Economic observability requirements brief would derive observability requirements from E-09/E-08 patterns — cited as informing, not assumed; not needed for this proposal to structure the surface. |

---

## Principle → Constraint → Invariant → Failure Condition → Test

- **Principle (PROPOSED):** destruction, loss, and removal are the economic-event surfaces that replace decay as the source of item attrition (under HD-ITM-01's ruled no-decay model) — they must create meaningful consequences for player investment (otherwise the no-decay model produces an attrition-free economy that drifts toward ECO-003 spreadsheet optimism and against VIS-003 player-driven stakes) **without** making ownership meaningless (PLR-003; INV-010(c)), without silently destroying irreplaceable provenance-bearing assets (INV-010(d)), and without creating exploitable race/exploit surfaces on the ownership transition (SAFE-001; SAFE-002; INV-013 adjacency).

- **Constraint (derives from LOCKED/ruled layer):** any destruction/loss/removal handling must (i) respect INV-010(c)/(d) — ownership and recourse remain meaningful, irreplaceable provenance-bearing items are not silently lost; (ii) respect HD-BLD-02's ruled instrument that zero-credit removal is terminal (no post-zero reclamation), while still answering *what happens to contents*; (iii) not invent placement/lot/density rules (BLD-004 remainder stays TBD); (iv) state what happens to provenance records on destruction/loss without engaging the DEFERRED provenance engine (PROvenance depth/visibility requirements only, per the non-canonical/deferred boundary); (v) carry no numeric parameters (BAL-001); (vi) not assume any transfer/ownership mechanism default from the non-canonical list.

- **Invariant (binding, already approved — not re-proposed):** INV-010 (conditional): if decay exists (it does not, under HD-ITM-01), its four clauses hold; under the ruled no-decay model, (c) and (d) bind as ownership/recourse obligations on any destruction/loss system; (b) binds the provenance requirements of any restoration mechanism. INV-BLD-004b (approved): no ghost equilibrium.

- **Candidate invariant for this surface (if approved):** *Salvage/restoration, where it exists, returns a meaningful and non-trivial fraction of an item's ownership-relevant substance (materials, usable structure, provenance continuity) to a player with standing, so that destruction/loss is consequential but never ownership-confiscation; and zero-credit removal of a structure resolves its contents through a defined rule (salvage-to-claimant, salvage-to-world, or destruction) that is stated, not left to race/exploit timing.*

- **Failure condition (this surface):** (a) destruction/loss that silently confiscates player assets with no salvage or recovery path (INV-010(c)/(d) pressure); (b) a removal rule whose contents outcome is "undefined / whatever the next player who clicks wins" — an exploit surface (SAFE-001; SAFE-002; per the HD-BLD-02 open question 6e(a)); (c) zero-credit removal that quietly becomes lossless because contents are auto-salvaged in full to the structure owner, eliminating the ownership consequence INV-BLD-004b's model depends on (ghost-equilibrium revival).

- **Test (shape-level):** checklist-class walk-through of the ruled model against the failure conditions above; the walk-through enumerates at least: (1) combat-destroyed equipment item; (2) lost/stolen item; (3) obsolescence-driven retirement; (4) voluntary upgrade replacing an old item; (5) zero-credit structure removal with a non-empty contents list; (6) zero-credit structure removal of a settlement's last structure (6e(b) adjacency); (7) provenance-bearing high-value crafted item destroyed; (8) irreplaceable item (one of a kind) destroyed. Full behavioural testing requires an implementation or simulation host (none exists today; the fork's teardown/loot patterns are PROTOTYPE evidence of shape only, not TCIndustries design — E-02/E-09 where relevant).

---

## Options (structural only; tradeoffs and failure modes stated)

### Option A — Salvage-to-owner with provenance continuity (the ownership-consequence-preserving default)

Destruction/loss of an owned item produces a salvageable remnant whose recovery is a **defined player-facing surface**: the item's owner (or anointed successor) can recover a meaningful fraction of the original (materials, repairable structure, or equivalent) through a salvage action; provenance continuity is preserved where the item was provenance-bearing (the salvage retains the identity link so that restoration, if later provided, does not lose provenance). Structure contents on zero-credit removal are handled by **explicit rule**: contents are either (i) salvaged to the structure owner before removal, (ii) left in-world as salvageable debris at the removal site, or (iii) destroyed alongside the structure — each a stated option for the owner to pick, not a default born of silence.

- *Serves:* INV-010(c)/(d) directly (ownership remains meaningful; irreplaceable items have a stated recourse path — at minimum the salvage remnant and the provenance record); SAFE-001's "abandoned assets and exploitative transfer" surface gets a defined answer instead of silence; INV-BLD-004b kept clean (removal is still terminal for the structure; contents rule is separate and can be chosen to preserve consequence); HD-BLD-02's open question 6e(a) answered in this proposal rather than left for a separate micro-ruling.
- *Tradeoffs:* salvage-to-owner is the option most in tension with "loss has teeth" — if salvage recovery is too generous, destruction becomes a mild inconvenience rather than an economic event (ECO-003/INV-009a pressure on the sink side); the *fraction* question is exactly the BAL-001-gated numeric that this proposal deliberately does not answer. Provenance continuity on salvage is a requirements statement (INV-010(b); INV-017) not an engine design — it says "if restoration exists, it must not break provenance"; it does not design the restoration engine (DEFERRED).
- *Failure modes:* salvage too generous → destruction ceases to be load-bearing (the no-decay model then has no economic teeth; the world drifts toward attrition-free); salvage too harsh → INV-010(c)/(d) pressure (loss becomes de-facto confiscation, especially for irreplaceable items); contents rule left as "owner auto-recovers everything" → ghost-equilibrium revival under INV-BLD-004b (the structure disappears but the owner lost nothing — that is the failure this proposal's contents sub-options are named to prevent).
- *Non-canonical-list check:* clean — no default tradeability assumption; no soul-binding default; salvage is a recovery *surface*, not a transfer-rights ruling.

### Option B — Salvage-as-world-resource (destruction feeds the economy, not the owner)

Destruction/loss returns materials to the world as salvageable debris/loot rather than (or in addition to) a direct owner recovery; recovery is a separate player activity (scavenging, retrieval, salvage service) that creates demand for an additional player role; provenance-bearing items that are destroyed lose their active provenance record (or leave a "lost/destroyed" marker) rather than retaining a restoration-ready continuity. Structure contents on zero-credit removal fall to the world (debris at the site, claimable by others / scavengeable) or are destroyed with the structure.

- *Serves:* strongest economic-event generation from loss (PIL-001 event shape: a lost/destroyed item becomes a scavenger opportunity, a service demand, possibly a territorial retrieval event); aligns with VIS-003 player-driven stakes (loss creates other players' opportunities); degradation of provenance on destruction is an honest expression of "destroyed items are gone" — avoids the "provanance-everything-forever" assumption.
- *Tradeoffs:* direct tension with INV-010(c)/(d) *as written* — the approved invariant says ownership remains meaningful and irreplaceable provenance-bearing assets are not destroyed without recourse; under Option B, "recourse" is not recovery-to-owner but the possibility of salvage competition, which is a weaker and more contestable reading; the owner's ruled invariant reading (recorded HD-ITM-01 §5a) leaned toward recourse-on-balance, so Option B is the option most likely to require an invariant *amendment* rather than clean adoption; provenance loss on destruction may be acceptable or may not — that is a ruling question, not a default.
- *Failure modes:* INV-010(c)/(d) failure if the rule is read as "destroyed = gone, no recourse"; scavenge turns loss into a free respawn for whoever arrives first (race surface, SAFE-002); provenance loss on destruction may silently defeat the identity/reputation pillar's downstream value (PIL-002) if high-value crafted items can be removed from the world's memory by whoever destroys them — needs explicit ruling, not silence.
- *Non-canonical-list check:* clean on tradeability/soul-binding; provenance-loss-on-destruction is a *new* design stance this option takes, so it must be ruled, not assumed.

### Option C — Restitution/contingency framing (no salvage; escalation-and-recourse only)

Destruction/loss does not auto-generate salvage; instead the design provides a **recourse/contingency** surface: lost or destroyed items may be recoverable through explicitly designed contingency channels (insurance-like player instruments, org protection, escrow, or replacement-on-condition) whose *existence and eligibility* are ruled here, while the contingency mechanics themselves remain open for later ruling. Structure contents on zero-credit removal are handled by a **contents-to-owner restitution** default (the structure's contents are retrievable by the owner for a defined window or condition before terminal removal, or held in custody) — giving ownership a real consequence (the window can be missed) without forfeiture-by-race.

- *Serves:* cleanest reading of INV-010(c)/(d) as "recourse exists and is meaningful" without making salvage a consumption surface; aligns with SAFE-001's "safeguards" framing (protection against exploitative loss); contents-to-owner restitution is the option most naturally aligned with the owner's HD-BLD-02 model (the credit balance is the recourse path; contents restitution extends that logic to the structure's holdings).
- *Tradeoffs:* heaviest "depends on later rulings" surface — contingency instruments (insurance, escrow, org protection, replacement) are genuinely separate design questions that this proposal names but cannot answer; the "window/condition" for contents retrieval is a structural shape whose *parameters* are BAL-001-gated; risks the recourse surface becoming the *only* economic event from loss (if destruction produces no salvage and no world resource, the only event is the contingency demand — which may be too narrow).
- *Failure modes:* contingency channels empty/unimplemented → destruction has no consequence and no recourse (the worst of both — neither salvage nor restitution materialises); contents restitution window too generous → ghost-equilibrium revival (INV-BLD-004b pressure); contingency design that defaults to a single mandatory insurer (likely an NPC or org) risks ECO-001 NPC-dominance pressure if players cannot enter the contingency market.
- *Non-canonical-list check:* clean; but any contingency instrument that requires a transfer/ownership default must not silently pick the non-canonical defaults.

### Cross-option notes (recorded, not decided)

- **Salvage *fraction*, retrieval window, debris persistence, and any "recoverable fraction" number are BAL-001-gated** — this proposal rules the *structural handling class* only (does the owner recover? does the world get debris? does provenance survive? is there a restitution window?), not the numbers.
- **The contents-handling sub-rule on zero-credit removal is the explicit composition point with backlog 6e(a).** The owner's HD-BLD-02 model made contents-handling an open question; this proposal folds it in as part of the salvage/destruction surface because the *ownership-consequence* design is shared — losing a carried item and losing a structure's contents are the same design question (what does the owner keep/recover/lose?). The proposal does not *require* the owner to rule them together; it offers the composition for the ruling session.
- **Backlog 6e(b) settlement-institution dissolution is named but deferred** — when a settlement's last structure is removed, the civic institution has no structures left. INV-007a is satisfied at structure scale by HD-BLD-02's model, but the institution's end-state (does it persist as a legal entity with no seat? does it dissolve? who inherits its obligations?) is a separate small ruling that this proposal flags and does not answer.
- **The salvage/destruction surface is where SAFE-002 (exploit implementation) and INV-013 (org asset ownership on membership events) touch the same ownership-transition machinery.** This proposal does not authorise either; it states the surface so that whichever option is ruled here is *checkable* against the still-open SAFE-002 policy ruling (live queue) and the still-open INV-013 candidate.

---

## What this proposal does NOT resolve

- Any numeric: salvage fractions, recovery rates, debris persistence timers, retrieval windows, restitution eligibility thresholds (BAL-001; separate authorised passes).
- The full SAFE-001 mechanism (transfer rules, ownership tooling, fraud protections remain open per HD-GDD-01 principle-only lock) — this proposal only states the destruction/loss/removal *handling* surface.
- Whether irreplaceable provenance-bearing items are recoverable in principle vs. destroyable with provenance record only — that is the INV-010(d) reading the owner gave at HD-ITM-01 §5a; it is recorded, and this proposal offers options consistent with both readings, but the *final reading* is a ruling question if it was not fully settled.
- The provenance *engine* (DEFERRED); provenance *depth/visibility requirements* (INV-017 candidate) are referenced as a requirement-shape only.
- Backlog 6e(b) settlement-dissolution rule; backlog 6c early-game on-ramp; OQ-006 transport; OQ-013 org rights — all named as adjacent, none resolved.
- Any implementation, simulation, or testbed work; the fork's teardown/loot and market_records patterns (E-02, E-09) are PROTOTYPE evidence of shape only (PROT-001).
- The OQ-010/HD-ITM-01 decay question (already ruled Option C — no baseline decay; not re-opened).
- The BLD-004 abandonment instrument (already ruled HD-BLD-02; not re-opened).

---

## Ruling requested (HD-<Area>-<NN> slot reserved in the Human Rulings Register)

**Question for the owner — rule yes / no / amend on each axis:**

1. **Primary ruling — salvage/destruction/contents handling class:** Pick the structural handling class for (a) item destruction/loss and (b) structure-contents-on-removal (the two compose under one ownership-consequence surface; the owner may split them):
   - **A** — Salvage-to-owner with provenance continuity (Option A above); contents rule chosen from the stated sub-options (salvage-to-owner / world-debris / destroyed-with-structure).
   - **B** — Salvage-as-world-resource (Option B); provenance on destruction handled as stated (or amended).
   - **C** — Restitution/contingency framing (Option C); contents restitution default as stated (or amended).
   - **Amend** — any hybrid or correction (e.g. A for items, C for structure contents; or a contents sub-rule not listed).
   - **Defer** — keep the surface open (note: under HD-ITM-01 + HD-BLD-02 the surface is load-bearing; defer leaves an explicit unresolved ownership-consequence gap).

2. **Invariant ruling (independent):** Do you confirm that the binding approved invariant INV-010 (conditional) reads, under the no-decay model, as (c)/(d) binding on any destruction/loss system and (a) attaching to the destruction/loss/obsolescence/upgrade demand sources — and do you approve the *candidate* invariant stated in this proposal's P→C→I→F→T block (salvage returns meaningful ownership-relevant substance; removal contents resolved by stated rule) as an additional TEST-001-class invariant for this surface?

3. **Scope confirmation:** Confirm that (a) all salvage/retrieval/contents parameters are deferred to BAL-001-authorised passes; (b) the provenance-engine remains DEFERRED (depth/visibility requirements only); (c) backlog 6e(b) settlement-dissolution remains a separate open small ruling; (d) SAFE-001/002 mechanism surfaces remain open and are informed, not closed, by this ruling.

4. **Reading clarification (if needed):** If Option B is under consideration, does the owner intend INV-010(c)/(d) to be read as "recourse = salvage competition / world resource" rather than "recourse = recovery-to-owner", or does Option B require an INV-010 amendment to be lawful? (Recorded so the ruling session can resolve the exact tension; the proposal does not pre-answer it.)

*Every claim above cites its file and anchor. On any ruling, the tracker updates and the queue maintains 3–5 live items per the loop rules. No status changes anywhere by this document.*
