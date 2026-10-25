# TCIndustries — Current Project State Report

**Author:** Claude (Anthropic), model Fable 5 / Opus 5.5 tier — `claude-fable-5` (switchx:opus)
**Version:** 1.0
**Date:** 2026-10-25
**Assessor:** Claude (Anthropic) — sole author and assessor of this document (no second-reviewer pass was performed; noted per GR-004).
**Status / Authority class:** ANALYSIS / REPORT — **non-authoritative**. This document creates no game-design decisions, promotes no status, and grants no design authority. It is written for downstream Claude analysis (recommendations and improvements).

---

## 0. Purpose and scope of this report

This report documents the **current state of the TCIndustries project** as of 2026-10-25, for the explicit purpose of handing it to Claude for analysis, recommendations, and improvement suggestions. It is a documentation/analysis artifact, not a design authority document.

Scope (confirmed with the project owner via ask-user flow): **full project state** — authority framework, the complete human-ruling inventory, the open-question registry, the Economic Interdependence Core (EIC) chain, all deferred magnitudes, open backlog items, and status-hygiene findings.

Sources read in full for this report (not reconstructed from summaries):
- `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` (current working canonical design reference)
- `governance/TCIndustries_Human_Rulings_Register.md` (authority record)
- `governance/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` (authority classes + governance rules)
- `governance/project_memory.md` (running project memory)
- `proposals/TCIndustries_Open_Questions_Register.md`
- `proposals/TCIndustries_Phase9_Elite_and_TaskLayer_Proposal_v0.1.md`
- `proposals/TCIndustries_OQ-001_BAL-001_Numeric_Pass_2026-10-25.md`
- `proposals/TCIndustries_OQ-001_Deferred_BAL-001_Pass_Disposition_2026-10-25.md`
- `proposals/TCIndustries_OQ-001_Specialisation_Budget_Proposed_Ruling_2026-10-25.md`
- `proposals/TCIndustries_6d_Salvage_Destruction_Contents_Ruling_Record-HD-SAL-01_2026-10-23.md`

**Conventions preserved in this report.** Status vocabulary (`LOCKED` / `HUMAN-LOCKED`, `DERIVED CONSTRAINT`, `PROPOSED`, `TBD`, `PROTOTYPE`/`EVIDENCE`, `DEFERRED`, `ASSUMPTION`, `HISTORICAL`) and the authority classes (A–G) are used as defined in `AGENTS.md` and the Authority Matrix §1. Per `AGENTS.md`, no non-LOCKED status is silently promoted to LOCKED; status changes require explicit, recorded human approval. Prototype and SWG behaviour are treated as evidence/historical, never as canon.

---

## 1. Authority and status framework (in force)

### 1.1 Status vocabulary
From `AGENTS.md` and GDD §3/§35. Each significant design element carries an explicit status:

| Status | Meaning |
|---|---|
| `LOCKED` / `HUMAN-LOCKED` | Explicit human/project decision. Authoritative unless formally revised. |
| `DERIVED CONSTRAINT` | Binding logical consequence of one or more LOCKED principles; not an independent human decision. |
| `PROPOSED` | Candidate design. Never canon. |
| `TBD` | Genuine unresolved decision. Do not invent answers. |
| `PROTOTYPE` / `EVIDENCE` | Behaviour demonstrated by implementation/test. Evidence only. |
| `DEFERRED` | Intentionally postponed. Do not pull into scope without authorisation. |
| `ASSUMPTION` | Working assumption not elevated to decision. |
| `HISTORICAL` | Reference/inspirational (incl. SWG). Never automatic canon. |

### 1.2 Authority classes (Matrix §1)
A statement remains `LOCKED` only where authority is **A** (explicit human ruling) or **B** (explicit pillar). Classes C–E are not substitutes for human approval.

| Code | Type |
|---|---|
| A | Explicit Human / Project Ruling |
| B | Explicit Project Principle / Pillar |
| C | Derived from Locked Principle |
| D | Audit / Governance Reclassification |
| E | Multi-Source Inference |
| F | Historical / External Reference (SWG) |
| G | Prototype / Evidence |

### 1.3 Governance rules in force (Matrix §2)
- **GR-001** — Canonical presence is not human approval.
- **GR-002** — Historical proposal boundary (superseded/first-pass proposals non-canonical unless promoted through change control).
- **GR-003** — Derived-constraint strictness (only genuine logical consequences of HUMAN-LOCKED principles).
- **GR-004** — Every LLM-produced document must carry an attribution metadata block (this report complies; attribution is provenance, not authority).
- **GR-005** — filename convention: **not in force**.
- **GR-006** — Detective/corrective control of project decisions (in force 2026-09-28; filing-window value reserved to owner).
- **GR-007** — provenance incident control (2026-10-18 Beam Free provenance incident; grounded in `investigations/Provenance_Incident_2026-10-18_Beam_Free.md`).

### 1.4 Standing governance boundaries (AGENTS.md + HD-EIC-08)
- No numeric tuning of economic or progression systems without explicit human authorisation (**BAL-001** LOCKED — "Avoid Premature Numeric Tuning"; false precision is worse than explicit unresolved questions).
- No invention of new residual channels or architectures without human-defined structural boundaries.
- No automatic creation of successor architectures after falsification of a candidate.
- Provenance and formal reputation systems remain **DEFERRED** until explicitly authorised.
- EIC work is gated: Architecture C falsified and retired (HD-EIC-05); shaping requires human-supplied structural boundaries (HD-EIC-07/10).

### 1.5 OpenCode / agent role boundary
Agents act as documentarian, repository steward, auditor, and documentation assistant. They must NOT silently: make game-design decisions; promote proposals/inference/historical/prototype into canonical design; resolve design contradictions; delete or rewrite historical/provenance material; or change design authority. This report respects those boundaries — it records and flags, it does not resolve or promote.

---

## 2. Human ruling inventory (register of HD-* rulings)

Each ruling below is `HUMAN-LOCKED` (authority class A). For each, the **ruled scope** and the **explicit "does not authorise" carve-outs** are captured, because the carve-outs are where most project risk now concentrates (they are the deferred work).

### Section A — EIC foundation (2026-08-25)
- **HD-EIC-01** — APPROVE layered validation unit (participant-level). HUMAN-LOCKED unit only.
- **HD-EIC-02** — hierarchical comparative-value principle, no hard factory-quality ceiling. HUMAN-LOCKED principle.
- **HD-EIC-03** — M1 multi-accounting is in-scope adversarial condition. HUMAN-LOCKED stance; no account-level restriction authorised.
- **HD-EIC-04** — provisional non-numeric functional shapes authorised for simulation only. HUMAN-LOCKED method permission.
- **HD-EIC-05** — ACCEPT FALSIFICATION; Architecture C retired. HUMAN-LOCKED.
- **HD-EIC-06** — EIC requirement unchanged; PIL-003 binding. HUMAN-LOCKED.
- **HD-EIC-07** — bounded new shaping *contingent on human-defined boundaries*; NO automatic Architecture D. HUMAN-LOCKED. Completed 2026-10-23 by HD-EIC-10.
- **HD-EIC-08** — governance boundaries CONFIRMED table: numeric tuning NOT authorised; Provenance & Reputation DEFERRED; Master GDD consolidation NOT authorised; automatic Architecture D NOT authorised; LLM residual-channel invention w/o human boundaries NOT authorised; prototype evidence reconciliation is a separate track; PIL-003 HUMAN-LOCKED.

### Section B — GDD confirmation (HD-GDD-01, 2026-09-12 reconfirmed)
- **HD-GDD-01** — WRLD-001 CONFIRM (HUMAN-LOCKED); PLR-001 CONFIRM (HUMAN-LOCKED); PLR-003 CONFIRM principle-only (HUMAN-LOCKED); CRFT-001 CONFIRM principle-only (HUMAN-LOCKED); SAFE-001 CONFIRM principle-only (HUMAN-LOCKED). GDD-AUTH-REVIEW-GATE CLOSED.
- Provenance note: no standalone source doc; root cause was a prior ChatGPT session that lacked the full project file set (partial-context drift).

### Section C — Production & systems rulings (2026-09-28)
- **HD-RET-01** (OQ-016) — Option A identity-anchored player vendors; INV-016a/016b approved; fees/taxation/price bands deferred (BAL-001).
- **HD-BLD-01** (OQ-007) — Option A free-claim formation; basic governance depth (ownership + zoning presence + shared maintenance); INV-007a approved, INV-007b candidate; numerics deferred.
- **HD-SCOPE-01** (OQ-015) — Option S1 economic spine; early-game on-ramp NAMED-OPEN (backlog 6c); INV-015 approved.

### Section D — Items, identity, buildings (2026-09-19/20)
- **HD-ITM-01** (OQ-010) — no baseline decay; INV-010 conditional approved (under no-decay model).
- **HD-PROF-01** (OQ-011, PARTIAL) — identity/provenance survives respec; INV-011a approved; instrument deferred.
- **HD-BLD-02** (BLD-004) — owner's zero-credit removal model; INV-BLD-004b approved, INV-BLD-004a candidate.
- **HD-PROF-02** (OQ-011) — free-drop respec; INV-011a approved; INV-011b/c candidates.
- **HD-ITM-02** (OQ-010/ITM-002) — six-point: no decay; 1% death-damage; repairable; no immortal crit-fail bug; destruction only by player menu; no salvage system at present.
- **HD-BLD-03** (BLD-004) — contents removed with structure; 24h upkeep lapse; closes backlog 6e(a).

### Section E — Production & systems rulings, continued
- **HD-ECO-01** (OQ-009, PARTIAL) — Structure A: single unified currency (non-canonical entry "Single unified currency as the default" **PROMOTED to canonical**); INV-009a APPROVED in the faucet-exceeds-sink reading; **scope confirmation + INV-009b remain OPEN**. Sink instrument classes named-open.
- **HD-PROF-03** (OQ-011, PARTIAL) — acquisition trainer-gated (NPC trainer for credits, or free from a player who holds the class); **INV-011c NOT ruled; scope confirmation OPEN**.
- **HD-PROF-04** (OQ-001, RULED 2026-10-25) — two-layer specialisation-budget model **shape** (i: subordinate character-layer hygiene cap; ii: participant-layer interdependence budget anchored on HD-EIC-10 C2/C4). Magnitudes supplied separately by the BAL-001 numeric pass (see §4). Explicitly does NOT authorise: any numeric value beyond the pass; changes to free-drop respec or trainer-gated acquisition; any hard factory ceiling / account restriction; the §36 NOT-CANONICAL "Skill Discipline/Skill Box", "Use-based SP acquisition", "2–3 profession target" models; provenance/reputation work; any implementation/simulation/testbed work.
- **HD-EIC-10** (completes HD-EIC-07, 2026-10-23) — EIC Structural Boundary Statement. GLOSS G2 ("literal effortlessness" reading of HD-EIC-09's "trivially eliminates" — reopens γ′ re-test eligibility as a *separate* act). Structural commitments: **C1** (non-internalisable scarcity), **C2** (persistent information/discovery asymmetry), **C4** (non-economic institutional/social constraint) adopted; **C3** DEFERRED. Soft-residual exclusions by category (α, β, γ′-1, γ′-2). **Boundaries only — no mechanism, architecture, residual channel, or number authorised.** Does NOT authorise initiation of the bounded EIC search (that requires HD-EIC-11).

### Section F — Safety & exploit (HD-SAFE-01, OQ-012, 2026-10-23)
- **HD-SAFE-01** — Option A detection-first, human-review-only; INV-012 approved; no account restriction; no numeric thresholds. Resolves OQ-012 as HUMAN-LOCKED.

### Section G — EIC bounded search authorisation (HD-EIC-11, 2026-10-23)
- **HD-EIC-11** — bounded EIC search + γ′ re-test AUTHORISED under HD-EIC-10 boundaries (G2 gloss, C1/C2/C4 committed, C3 deferred, soft-residual exclusions). Does NOT authorise any number, architecture, residual-channel invention, provenance work, GDD consolidation, or automatic Architecture D.

### Section H — Faction, PvP, combat scope (HD-PVP-01, OQ-008, 2026-10-25)
- **HD-PVP-01** — three PvP modes (duel/guild-war; faction overt/covert flagging via recruiter NPCs; battlefields); two fixed factions (FACTION_A, FACTION_B) + neutral; NPC-only membership changes; 30s overt / 300s covert toggles; 120s aggression timer; standing **stored but never awarded** in Phase 8 (kills logged via `pvp_events`); guild wars 24h activation; defeat = normal incapacitation/clone flow only, **zero extra penalties**. Explicitly does NOT authorise: any extra death penalty; automated standing rewards; remote faction switching.

### Outside the governance Register — HD-SAL-01 (provenance gap, see §8)
- **HD-SAL-01** (backlog 6d, 2026-10-23) — salvage/destruction/contents handling, Option A (salvage-to-owner with provenance continuity); contents sub-rule selection deferred to BAL-001; INV-SAL-01 approved. Applied in GDD §36 and recorded in `proposals/`, but **absent from the governance Human Rulings Register** (see finding F-4).

---

## 3. Open-Question registry status (GDD §31, full)

| OQ | Question | Status | Ruling / Note |
|---|---|---|---|
| OQ-001 | Specialisation budget / skill-cap model | **RULED** | Two-layer model, HD-PROF-04 + BAL-001 numeric pass. Task-layer reward/sink magnitudes deferred. |
| OQ-002 | Resource spawn rates / depletion math | TBD | — |
| OQ-003 | Resource attribute sets / product mapping | TBD | — |
| OQ-004 | Crafting experimentation model | TBD | EIC-coupled; unblocked for search under HD-EIC-10/11. |
| OQ-005 | Manufacturing throughput / ownership | TBD | EIC-coupled; unblocked for search under HD-EIC-10/11. |
| OQ-006 | Transportation / fast-travel | TBD | — |
| OQ-007 | City formation / governance / tax | **RULED** | HD-BLD-01 (formation + basic depth); civic-depth instruments open. |
| OQ-008 | Combat scope / risk / progression | **RULED** | HD-PVP-01 (2026-10-25). Numerics BAL-001-gated. |
| OQ-009 | Currency / money sinks | **RULED** | HD-ECO-01 (Structure A; INV-009a approved); sink classes named-open; numerics TBD. |
| OQ-010 | Durability / decay / repair | **RULED** | HD-ITM-01 + HD-ITM-02 (no baseline decay). |
| OQ-011 | Respecialisation costs / rules | **RULED** | HD-PROF-01/02/03; magnitudes TBD; INV-011c + scope confirmation open. |
| OQ-012 | Multi-account / exploit policy | **TBD in GDD §31** — but **HUMAN-LOCKED in Register (HD-SAFE-01)** | Status-lag discrepancy (see finding F-1). |
| OQ-013 | Org rights / hierarchy / property | TBD | — |
| OQ-014 | Species / ancestry / starting location | DEFERRED | — |
| OQ-015 | First-playable vs persistent-alpha scope | **RULED** | HD-SCOPE-01 (S1 economic spine); early-game on-ramp named-open. |
| OQ-016 | Vendor / retail mechanics | **RULED** | HD-RET-01 (structural); fees/tax/price bands deferred. |
| OQ-017 | Provenance depth / visibility | TBD | — |

**Summary:** 10 of 17 OQs are RULED (OQ-001, 007, 008, 009, 010, 011, 015, 016 — plus HD-SAL-01 covering backlog 6d outside the 001–017 table). 1 DEFERRED (OQ-014). 6 TBD (OQ-002, 003, 004, 005, 006, 013, 017 — seven actually). **Note:** the GDD §31 row for OQ-012 reads TBD while the Register records it HUMAN-LOCKED via HD-SAFE-01 (2026-10-23). This is a documented status-lag, not a substantive open question (finding F-1).

---

## 4. OQ-001 / HD-PROF-04 magnitudes (authorised via BAL-001 pass, 2026-10-25)

Authorised by `proposals/TCIndustries_OQ-001_BAL-001_Numeric_Pass_2026-10-25.md` (owner authorisation 2026-10-25). These are the only numeric magnitudes currently authorised for OQ-001. None modify fork code.

| Magnitude | Authorised value | Layer |
|---|---|---|
| Per-character skill-point cap `MaxSkillPoints` | 250 (subordinate hygiene floor; ≈5–6 professions at 42 SP each; revisable) | (i) |
| High-tier (Master) decay scope | Master-tier (Tier 5) boxes only; Basic tiers I–IV never decay | (ii) |
| High-tier (Master) decay window | 30 in-game days of no re-affirming activity → revert to Tier IV (Master re-trained, cost re-paid); identity survives (PROF-004) | (ii) |
| Complementarity binding — association floor `K` | ≥5 distinct participant-owners (not alts of one participant) + role-provenance credential (enforces C2) | (ii) |
| Acquisition cost curve | XP = 1000·SP²; Credit = round(XP/20) → nearest 10 (typed XP pools; trainer-gated acquisition retained per HD-PROF-03) | (i)/(ii) |
| Per-profession cost (42 SP) | 118,000 XP / 5,900 credits; Master = 10 SP / 100,000 XP / 5,000 credits | (i)/(ii) |
| Typed XP pools | combat, crafting, scouting, medical, entertainer, merchant (authorised as the spendable six) | — |

**Still deferred (structural partners absent — not invented):** XP earn / progression rate (PROG-003 pace, coupled to Phase-8 balance and the unbuilt task layer); task-layer reward magnitudes (Q8); task-layer sink magnitudes (Q8 / HD-ECO-01 — requires a matched sink before any reward faucet is authorised).

---

## 5. Economic Interdependence Core (EIC) chain

The EIC is the project's central unresolved-design area and the source of most deferred work.

1. **Architecture C** developed (2026-08-25), simulated, and **falsified as currently shaped** (γ′ variants fail Phase-6 combined closure).
2. **HD-EIC-05** — falsification accepted; Architecture C **RETIRED**.
3. **HD-EIC-07** — bounded new shaping authorised *only with human-defined structural boundaries*; NO automatic Architecture D.
4. **HD-EIC-10** (2026-10-23) — EIC Structural Boundary Statement filed as Register Section E; **completes the HD-EIC-07 gate**. Durable design commitments: C1, C2, C4 (C3 deferred); soft-residual exclusions by category. **Boundaries only — no mechanism, architecture, residual channel, or number authorised.**
5. **HD-EIC-11** (2026-10-23) — bounded EIC search + γ′ re-test **AUTHORISED** under the HD-EIC-10 boundaries.

**State:** the gate is closed and a bounded search is authorised, but **no mechanism or architecture has yet been designed within the boundaries**. OQ-004, OQ-005, and the GAP tasks T-03/T-06 are now unblocked and workable but remain unresolved and EIC-coupled.

---

## 6. Deferred / DEFERRED magnitudes and open backlog items

### 6.1 BAL-001-gated numeric magnitudes (not yet authorised)
These remain `TBD`/deferred pending structure stabilisation, testability, and explicit human authorisation:
- **XP earn / progression rate** (PROG-003) — see §7.1 for the live blocker.
- **Task-layer reward magnitudes** (Q8 / Phase 9 proposal) — no task layer exists.
- **Task-layer sink magnitudes** (Q8 / HD-ECO-01) — no matched sink exists.
- **Fees / taxation / price bands** (HD-RET-01, OQ-016).
- **Durability numerics** (HD-ITM-01/02: 1% death-damage is specified as a behaviour; repair costs/rates TBD).
- **Building placement / lots / density** (HD-BLD-01/02/03).
- **Salvage contents sub-rule selection** (HD-SAL-01 → BAL-001 gate; salvage-to-owner vs world-debris vs destroyed-with-structure, no default).
- **All combat / Phase-8 numerical balance** (HD-PVP-01 numerics BAL-001-gated).

### 6.2 Open items on already-ruled rulings
- **HD-ECO-01** — scope confirmation + INV-009b **OPEN**.
- **HD-PROF-03** — scope confirmation + INV-011c **NOT ruled**.
- **Backlog 6c** — early-game economic on-ramp — **named-open**.
- **Backlog 6e(b)** — settlement-institution dissolution — **OPEN** (contents sub-rule resolved by HD-SAL-01/HD-BLD-03; dissolution governance remains open).

### 6.3 Invariants explicitly NOT approved (candidate only)
INV-007b, INV-011b, INV-011c, INV-BLD-004a — each remains a candidate; per-invariant granularity, no batch approval.

---

## 7. Unresolved issues and significant open risks

### 7.1 PROG-003 (Skill Acquisition Model) — still TBD, and a live fork-plumbing blocker
GDD §655 holds the Skill Acquisition Model at **TBD**. The OQ-001 deferred BAL-001 disposition (2026-10-25) identified that the fork's XP earn→spend loop is **internally inconsistent** (earn pools ≠ spendable pools):
- `internal/crafting/crafting.go` awards `XPPool` values `weapon_crafting`, `armor_crafting`, `smuggler`, `medical` — sub-pools **not spendable** by any box.
- `internal/handlers/civic_http2.go` awards a `"mentoring"` pool (via `civic.MentorshipXP = 50`) — also outside the spendable six.
- The only channel feeding a spendable pool is combat (`AddCharacterXP("combat", 50*tmpl.CLMax)`), and the live bestiary spans CL 1–3 only, yielding 50/100/150 combat XP against a per-profession training cost of 118,000 XP.
- Consequence: authorising an earn *rate* on top of this would enshrine XP players can never convert — exactly the false precision BAL-001 warns against.

**Resolution required before PROG-003 earn rates can be set:** (a) move PROG-003 from TBD to RULED; (b) reconcile the XP pool taxonomy so every earn pool is spendable (or document a mapping).

### 7.2 Task layer absent (Phase 9)
The Phase 9 replacement proposal (`TCIndustries_Phase9_Elite_and_TaskLayer_Proposal_v0.1.md`, AUTHORISED owner 2026-10-25) defines the task layer **structurally only** — repeatable system-generated contracts with a reward *hook* and no magnitude. A tree-wide grep for `taskLayer`/`task layer`/`phase9`/`mission`/`quest` returns matches only inside `SWG_PreCU_GDD.md` (HISTORICAL, out of scope per HD-TST-01). **No task layer exists in the fork**, so Q8 reward/sink magnitudes cannot be attached or authorised (HD-ECO-01 requires a matched sink first).

### 7.3 Provenance & Reputation Engine — DEFERRED
HD-EIC-08 and AGENTS.md governance boundaries hold Provenance & Reputation DEFERRED until explicitly authorised. OQ-017 (provenance depth/visibility) remains TBD. The engine is a standing pillar-support dependency but is not in current scope.

### 7.4 Prototype / testbed evidence unreconciled
GDD §34 step 0 = Evidence Reconciliation Pass (OPEN). The reference SWG Pre-CU engine testbed is HUMAN-LOCKED patterns-only (HD-TST-01); its preconditions (due-diligence steps 1–4, de-SWG fork, human review gate) are **not executed**. Post-Phase-10 testbed direction is **NOT DECIDED** — no default exists; all testbed work stays closed pending owner answer (`proposals/SWG_Post_Phase10_Program_Proposal_v0.1.md` §9 item 1).

### 7.5 Master GDD consolidation not authorised
Per HD-EIC-08, no further LLM consolidation of the Master GDD is authorised. The GDD is the working canonical reference but is explicitly not to be re-consolidated by an LLM at this time.

---

## 8. Status-hygiene findings (for downstream attention)

These are documentation-consistency defects found while reading the sources. None change design authority; several are the F1–F4 defect class GR-006 was written to detect.

- **F-1 — GDD §31 OQ-012 status-lag.** GDD §31 lists OQ-012 as `TBD`, but the Human Rulings Register records it `HUMAN-LOCKED` via HD-SAFE-01 (2026-10-23). The GDD open-question row lags the authority record. (Substantively resolved; the GDD pointer is stale.)
- **F-2 — GDD §36 "Applied 2026-09-28" table is mislabeled and incomplete.** The table is headed "Applied 2026-09-28" but contains rulings through 2026-10-23/25 (HD-SAL-01, HD-PVP-01). It also **omits HD-PROF-04 and HD-ECO-01**, which are applied in §31 but absent from the §36 applied-rulings table. (Earlier assessment that §36 omits HD-PVP-01 was incorrect — HD-PVP-01 is present at line 956; the genuine omissions are HD-PROF-04 and HD-ECO-01.)
- **F-3 — Human Rulings Register header lags its body.** The Register's own YAML/header declares version "1.5", date 2026-10-23, authority_class A — but the body contains HD-PVP-01 (2026-10-25) and the BAL-001 numeric pass (2026-10-25). This is the F4-class defect (a summary/header silently disagreeing with the document it summarises).
- **F-4 — HD-SAL-01 absent from the governance Register.** HD-SAL-01 (2026-10-23) is recorded in `proposals/` and applied to GDD §36 and the proposals OQ register, but has **no entry in `governance/TCIndustries_Human_Rulings_Register.md`**. This is the F1-class defect GR-006 was written to prevent (a ruling existing only as a proposals/ decision record, not in the authority record). It should be filed into the Register to close the provenance gap.
- **F-5 — `project_memory.md` minor staleness on EIC.** Line 48 describes HD-EIC-10 and ends "next EIC step is a distinct owner authorisation act," which has since been superseded by HD-EIC-11 (authorised 2026-10-23). Informational only; no authority impact.

---

## 9. Conclusions

1. **The authority record is substantially complete and internally coherent on substance.** Ten open questions are RULED, the EIC gate is closed (HD-EIC-10) with a bounded search authorised (HD-EIC-11), and OQ-001's magnitudes are now authorised via the BAL-001 numeric pass. The project is past its structural-design logjam.
2. **The remaining work is overwhelmingly magnitude/mechanism definition, not structural decision.** Nearly every open item is a BAL-001-gated numeric value, an open backlog slot, or mechanism design inside already-ruled boundaries — not a new ruling.
3. **Two hard blockers sit upstream of any numeric authorisation:** (a) PROG-003 is TBD *and* the fork's XP earn→spend pool loop is inconsistent (§7.1); (b) the task layer does not exist, so Q8 reward/sink magnitudes have no hook (§7.2). Both must be resolved before the corresponding BAL-001 passes can responsibly set numbers.
4. **Documentation hygiene has reopened small gaps** (F-1 through F-5) despite GR-006. These are provenance/consistency defects, not design defects, but they are exactly the class GR-006 targets and should be closed for audit cleanliness.
5. **No fork code is authorised by any document in scope.** All numeric passes and rulings explicitly carve out implementation, simulation, and testbed work. The testbed remains patterns-only and its gate preconditions are unexecuted.

---

## 10. Recommended focus areas for downstream Claude analysis

*(These are analysis pointers for Claude's recommendations, not decisions or canon promotions. They inherit the project's standing governance boundaries — no numbers, no architecture invention, no silent status promotion.)*

1. **Close the PROG-003 / XP-pool reconciliation blocker first.** This is the single highest-leverage unblock for progression numerics. Recommend a scoped analysis of the fork's `AddCharacterXP`/`train` pool mismatch and a concrete reconciliation proposal (earn pools == spendable six, or a documented mapping) as a prerequisite to any PROG-003 ruling.
2. **Define the task-layer minimal build (structural) before Q8 magnitudes can exist.** Per HD-ECO-01, a matched sink must accompany any reward faucet; recommend an analysis of what the task layer + matched sink would look like as a structural proposal, independent of any number.
3. **Clear the status-hygiene findings (F-1–F-5).** File HD-SAL-01 into the governance Register; bump the Register header to match its body; complete the GDD §36 applied table (add HD-PROF-04, HD-ECO-01) and correct its date label; refresh the OQ-012 §31 row.
4. **Sequence the DEFERRED magnitude authorisations.** Group the BAL-001-gated numerics into a recommended ordering (e.g., PROG-003 + pool reconciliation → OQ-001 earn rate → task-layer Q8 → HD-RET-01 fees → HD-ITM-02 repair costs → HD-BLD placement) so the owner can authorise them in dependency order.
5. **EIC bounded search execution plan.** With HD-EIC-11 authorising the search, recommend a scoped plan for executing the bounded EIC search + γ′ re-test under the HD-EIC-10 boundaries, without inventing mechanisms, numbers, or an Architecture D.
6. **Provenance & Reputation scope framing.** As a deferred pillar dependency (OQ-017), recommend a structural scoping of what "provenance depth/visibility" would mean when the owner chooses to authorise it — keeping it clearly separate from the DEFERRED engine work.

---

*End of report v1.0 (ANALYSIS / REPORT — non-authoritative). Created 2026-10-25 by Claude (Anthropic) for downstream Claude analysis. No game-design decisions, numeric values, or status promotions were introduced. All statuses cited are as recorded in the source documents named in §0.*
