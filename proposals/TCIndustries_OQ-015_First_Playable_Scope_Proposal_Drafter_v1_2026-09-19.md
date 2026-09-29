# TCIndustries — OQ-015 First-Playable vs Persistent-Alpha Scoping — Structural Proposal

**Filename:** `TCIndustries_OQ-015_First_Playable_Scope_Proposal_Drafter_v1_2026-09-19.md`
**Status:** PROPOSED. No authority. Ruling-ready candidate only. Nothing here is
LOCKED, and no element of this proposal changes any GDD status. No effort
estimates and no calendar commitments (Gap Analysis v0.1 §6.3 pattern — OQ-015
is a scope definition, not a schedule).
**Author:** Buffy — Codebuff agent (Freebuff), autonomous-loop drafter, Cycle 1.
**Assessor:** None — single-session authoring; no independent review pass (per GR-004, recorded as such rather than guessed).
**Date:** 2026-09-19
**Queue position:** 5 of 5 (see `proposals/TCIndustries_Continuous_Queue_Tracker.md`)
**Pre-flight:** no prior proposal or investigation addresses OQ-015 scoping; the gap ID appears only in invariant and evidence rows (cited below). No duplicate pass.
**Provenance note:** this file's first write attempt was corrupted mid-generation and was fully rewritten in this clean version; no content from the corrupted draft is retained or relied upon.

---

## Status wall

- Every option below is **PROPOSED** and requires an explicit HD ruling to advance.
- **Zero numeric values** — no system counts, player counts, durations, or
  thresholds (BAL-001). Scope *enumerations* name systems; they set no numbers.
  Boundary case flagged for the owner's numeric-touch screen (Multi-LLM Plan §4):
  a systems list is enumeration, not tuning.
- **EIC architecture untouched.** The core loop touches interdependence by
  definition; nothing here selects any interdependence mechanism
  (HD-EIC-07 untouched; Architecture C stays retired, HD-EIC-05). INV-015
  exists precisely so a scope ruling cannot silently select one.
- **Species/setting content (OQ-014) is DEFERRED** (GDD §30) — none enters any option.
- **The scoping decision itself remains human.** The Invariant Register's OQ-015
  entry records: "the actual scoping decision remains human (T-08 brief covers
  it)". The T-08 decision-support brief is NOT approved and NOT assumed here;
  this proposal structures the ruling independently of it.
- **GDD §34 sequence respected.** The LOCKED-layer programme order (Core first;
  then Provenance [DEFERRED], Commerce & Services, Society, World & Logistics,
  Combat, simulation/numerics last) is not reordered by any option — an option
  only decides which subset of systems forms the first playable and which are
  deferred to persistent alpha.

## Anchors (controlling references)

| Anchor | Status | What it locks/constrains at the scope ruling |
|---|---|---|
| OQ-015 (GDD §31) | TBD | "Which systems are required for first playable prototype vs. persistent alpha?" — the open question structured here |
| INV-015 (Invariant Register v0.1, OQ-015 entry) | Candidate | '*Any "first playable" definition must include enough economy for at least one interdependence relation to be real — i.e., the minimal slice cannot consist of systems a single player exercises alone.*' (derives from VIS-004 + LOOP-002) |
| GDD §2 central loop | LOCKED layer | "Discover resources → acquire and process materials → craft goods or provide services → specialise → establish reputation → participate in trade, society, and regional economies" — the spine any scope must trace |
| LOOP-001 (GDD §7) | PROPOSED | The eight-loop table — a menu for scope composition, not itself canon |
| VIS-004 / PIL-003 (GDD §5/§6) | LOCKED | Interdependence — INV-015's source; the slice must make at least one between-player relation real |
| LOOP-002 (GDD §7) | LOCKED | Non-combat viability — no scope may gate careers behind combat |
| RET-001 / SERV-001 / BLD-001 / SOC-001 (GDD §§16/17/19/20) | LOCKED | Retail, services, structures, organisations — candidate in-scope LOCKED capabilities |
| MFG-001 (GDD §13) | LOCKED | Controlled automation — factories stay OUT of first playable unless separately ruled in (anti-drift guard at scope time) |
| E-10 (Design Evidence Register §1.1/§2 OQ-015 row) | PROTOTYPE | Full 10-phase build chain — effort-shape reference and a worked single-player-closure counterexample; implementability evidence only (PROT-001) |

## Principle → Constraint → Invariant → Failure Condition → Test

- **Principle (PROPOSED):** the first playable is the smallest slice of the
  LOCKED vision that is genuinely multiplayer-economic: it must let at least
  one interdependence relation occur between real participants (INV-015),
  trace the GDD §2 central loop end-to-end, and leave every excluded system
  explicitly deferred rather than silently assumed.
- **Constraint (derives from LOCKED layer):** any scope must (i) satisfy
  INV-015 (no single-player-closable loop), (ii) respect non-combat viability
  (LOOP-002 — combat may be present but nothing in the slice may require it),
  (iii) exclude factories unless separately ruled in (MFG-001 guard),
  (iv) exclude species/setting content (OQ-014 DEFERRED), (v) record every
  deferred system as TBD-not-assumed (no silent local assumptions — the
  documentation standard this repository exists to enforce).
- **Invariant (candidate INV-015, quoted):** '*Any "first playable" definition
  must include enough economy for at least one interdependence relation to be
  real — i.e., the minimal slice cannot consist of systems a single player
  exercises alone.*'
- **Failure condition:** a proposed minimal scope whose loop closes within one
  character (the exact anti-goals #6/#7 degeneration; E-10's fork chain is the
  worked example of single-player closure to design against).
- **Test (shape-level):** scope-proposal review against INV-015
  (checklist-class): enumerate the slice's systems, then name the real
  between-player relations it enables and verify at least one is unavoidable
  in normal play (not optional decoration). The INV-015 check table below
  applies it to each option now.

## Options (structural only; tradeoffs and failure modes stated)

### Option S1 — Economic spine
Discover → acquire → process → craft → sell via player vendors, plus one
player service (e.g. medical) generating recurring demand. Out of scope:
combat, cities, organisations, transport, factories.
- *Real interdependence relations:* craft→vendor→buyer; service demand
  (combatant-optional, entertainment/medical consumers). INV-015 satisfiable.
- *Serves:* the GDD §2 loop nearly verbatim; smallest LOCKED-capability set
  (RET-001, SERV-001) that is still economic; aligns with GDD §34 (Commerce &
  Services precedes Society/World/Combat).
- *Tradeoffs:* no social layer — "player-created society" (PIL-004) is absent
  from first playable, deferring that pillar's first evidence to alpha; the
  NPC-supply floor assumed for early-game shop viability is exactly the OQ-016
  Option B shape — flagged: if OQ-016 is ruled Option A (player-only
  vendors), S1's baseline assumption needs its own ruling.
  **[Cycle-4 annotation, 2026-09-19]** OQ-016 was ruled 2026-09-19 (HD-RET-01,
  Option A — player-only vendors; decision record
  `proposals/TCIndustries_OQ-016_Ruling_Record_HD-RET-01_2026-09-19.md`):
  the NPC-supply-floor assumption is **dissolved**, not adopted — under the
  ruled model there is no NPC supply of player-producible goods at all, so
  early-game shop viability becomes an explicit OQ-015 scope design question,
  answerable at ruling time without importing any OQ-016 assumption.
- *Failure modes:* if the service loop is cut for size, the slice can drift
  toward single-player closure (INV-015 breach — the check table guards this);
  skill-acquisition machinery must exist but its model is open (OQ-011
  acquisition half) — slice must ship a placeholder acquisition shape labeled
  PROTOTYPE, not canon.

### Option S2 — Settlement slice
S1 plus minimal settlement presence: one placeable civic structure set and a
basic shared-maintenance obligation (no governance instruments, no taxation).
Out of scope: combat, full city governance (OQ-007), org rights (OQ-013),
transport, factories.
- *Real interdependence relations:* S1's relations plus settlement
  maintenance requiring distributed contribution (INV-007a-shaped behavior
  appears early). INV-015 satisfiable.
- *Serves:* gives PIL-004 its first evidence in the playable; natural home for
  vendor clustering (RET-003 embryonic); co-ruled cleanly with OQ-007 if the
  owner rules "basic depth" there (the two rulings compose).
- *Tradeoffs:* embeds a sub-decision of OQ-007 (formation model) — flagged:
  ruling S2 before OQ-007 risks a city-formation assumption entering by the
  back door; recommended sequencing if chosen is OQ-007 first.
  **[Cycle-6 annotation, 2026-09-19]** OQ-007 was ruled 2026-09-19 (HD-BLD-01:
  Option A free-claim formation + basic governance depth — ownership, zoning
  presence, shared maintenance responsibility; decision record
  `proposals/TCIndustries_OQ-007_Ruling_Record_HD-BLD-01_2026-09-19.md`):
  S2's settlement presence + shared-maintenance obligation now composes with
  ruled structure — no import remains. OQ-015 is **fully unblocked** (both
  embedded sub-decisions ruled: OQ-016 → HD-RET-01; OQ-007 → HD-BLD-01).
- *Failure modes:* settlement scaffolding without governance can become
  decorative (no decision-relevant placement value — INV-007b pressure);
  abandonment handling (BLD-004 surface) becomes visible immediately and is
  unresolved.

### Option S3 — Shortest loop
Survey → harvest → craft → vend only, with acquisition placeholder and no
service layer. Explicitly *fails* INV-015 as a standalone "first playable"
candidate — listed deliberately as the reductio: its loop closes within one
character plus a buyer, which is the degenerate slice the invariant exists to
catch.
- *Purpose:* included so the owner can see the boundary of INV-015 applied;
  selectable only together with an explicit invariant waiver (recorded as
  such, never silently).
- *Tradeoffs:* fastest to stand up; but every downstream interdependence
  invariant (INV-001…017 family) would be untestable in it, making it a weak
  testbed for exactly the properties the project cares about.
- *Failure modes:* normalises the anti-goal "single-player crafting game with
  multiplayer chat" (VIS-006) if promoted by repetition — the strongest reason
  to record any S3 selection as waiver-with-rationale.

## INV-015 check per option

| Option | Real between-player relations | INV-015 verdict |
|---|---|---|
| S1 Economic spine | craft→vendor→buyer; service demand | Satisfies (relations unavoidable in normal play) |
| S2 Settlement slice | S1's relations + distributed settlement maintenance | Satisfies |
| S3 Shortest loop | One optional buyer relation only | **Fails** — selectable only with explicit recorded waiver |

## What this proposal does NOT resolve

- The scope itself (that is the requested ruling) and any schedule or effort
  estimate (deliberately absent).
- Which acquisition model the placeholder uses (OQ-011 acquisition half, open;
  PROTOTYPE-labelled placeholder assumed shapeless until ruled).
- City formation (OQ-007, queued), org rights (OQ-013), transport (OQ-006 —
  zero evidence anywhere, backlog), currency (OQ-009), combat scope (OQ-008) —
  each stays with its own queued/backlog proposal.
- Whether factories exist at all (MFG-001 LOCKED constrains them; OQ-005
  facility rules are open and EIC-coupled).
- Any implementation work — this proposal decides scope structure only; actual
  build work would be separate tasks under separate approvals (per the
  Gap Analysis approval mechanics, §7.2).

## Ruling requested (HD-<Area>-<NN> slot reserved in the Human Rulings Register)

**Question for the owner — rule yes / no / amend:**

1. **Primary ruling:** Select the first-playable scope: **S1** (economic
   spine), **S2** (settlement slice; recommend ruling OQ-007 first), or
   **S3 with explicit INV-015 waiver** (rationale recorded), or **defer**.
2. **Invariant ruling (independent):** Do you approve candidate invariant
   INV-015 as a design invariant (TEST-001 class) binding all future scope
   proposals? (If S3-with-waiver is chosen, INV-015 is still approvable — the
   waiver then names the specific exception and rationale.)
3. **Scope confirmation:** Confirm that systems named out-of-scope are recorded
   as deferred-TBD (not assumed absent from the design), and that no schedule
   or numeric commitment is implied by this ruling.

*Every claim above cites its file and anchor. On any ruling, the tracker updates
and the queue maintains 3–5 live items per the loop rules.*

---
*End of proposal. PROPOSED — awaiting HD ruling. No status changed anywhere by this document.*
