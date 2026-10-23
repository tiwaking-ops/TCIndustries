# Filing-Time Acts for HD-SAL-01 and 2026-10-18/2026-10-23 Events

These are **owner-only filing acts** reserved per GR-006 and each decision record's §4/§7.
They update `canonical/` and `governance/` documents, which OpenCode cannot edit per AGENTS.md boundaries.

The agent has already completed the proposal-side annotations:
- ✅ Invariant Register updated (INV-010 CONFIRMED + INV-SAL-01 APPROVED, filing note added)
- ✅ Open Questions Register §4 updated (items 7, 8, 11 annotated as RULED)
- ✅ Queue tracker updated (6d RULED, Cycle 19 entry)
- ✅ HD-SAL-01 decision record created in `proposals/`

The agent has NOT done (cannot do) the following canonical/governance edits. **These are for the owner/coordinator to apply.**

---

## Action A: Master GDD Status-Patch Table (canonical/)

**File:** `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`
**Location:** §29 status-patch table, after line 955 (after the BLD-004 row)

**Add two rows:**

```
| OQ-011 / PROG-003 | HD-PROF-03 | Skill acquisition **RULED** (trainer-gated with free player-teacher path; starting class + additional classes via trainer; credit-gated); INV-011c still open; magnitudes TBD (BAL-001) |
| Backlog 6d | HD-SAL-01 | Salvage/destruction/contents handling **RULED** (Option A: salvage-to-owner with provenance continuity; contents sub-rule selection deferred to BAL-001; INV-SAL-01 approved); backlog 6e(a) closed; backlog 6e(b) remains open |
```

**Note:** HD-PROF-03 was already filed in the register Section E on 2026-09-28, and its Master GDD status line was supposed to be applied per GR-006. This is now caught up. HD-SAL-01 is new (2026-10-23).

---

## Action B: project_memory.md Event Lines (governance/)

**File:** `governance/project_memory.md`
**Location:** Append after the last existing entry (line 40, the "TWO RULINGS RECORDED AND FILED 2026-09-28" line)

**Append these event lines:**

```markdown
- GOVERNANCE FILED 2026-10-18 — HD-TST-01 final approval text filed in governance/TCIndustries_Testbed_Reuse_Human_Ruling-HD-TST-01_2026-09-14.md at the owner's direction via SWG Post-Phase-10 §9 item 9; mechanical act only, adds no new ruling or scope; preconditions (due-diligence steps 1-4, de-SWG fork, human review gate) remain Pending until separately executed and evidenced.
- GOVERNANCE RULE GR-007 IN FORCE 2026-10-18 — provenance attribution must cite a written convention (not infer one); agents must record their true identity or `Author LLM Unknown`. Added to Authority Matrix §GR-007 by owner approval. Grounding: provenance incident 2026-10-18 (Beam Free mis-attribution corrected, `investigations/Provenance_Incident_2026-10-18_Beam_Free.md`). No design authority created.
- OWNER DECISION 2026-10-18 — Post-Phase-10 program direction: **Option A — Consolidate** authorized (§9.1 answers recorded in `proposals/SWG_Post_Phase10_Program_Proposal_v0.1.md`). Adopt ~65 pre-existing testbed paths into git; close heal-evidence gap; port overrides for phases 1-9 clients; HD-TST-01 filing. Proxy 4 DEFERRED; Historical Accuracy questionnaire (`proposals/SWG_Historical_Accuracy-v1.md`) authored (pending human answers). B/C/D remain unselected and unauthorized.
- HISTORICAL ACCURACY INSTRUMENT AUTHORED 2026-10-18 — `proposals/SWG_Historical_Accuracy-v1.md` (v1.0, 223 questions) created as the comprehensive human-recollection questionnaire; owner to answer from recollection; answers are EVIDENCE only. Proxy 4 (Sandbox Validation) DEFERRED (requires live player cohort; no simulated players permitted).
- 6D PROPOSAL FORMAL QUEUE ENTRY 2026-10-23 — `proposals/TCIndustries_6d_Salvage_Destruction_Contents_Proposal_v0.1_2026-09-20.md` formally queued as Cycle-15 live item (was designated Cycle-14 next draft but not queued as other items were still live). Composes HD-ITM-01 (no decay → destruction/loss sole item-destruction source) + HD-BLD-02/03 (contents handling on zero-credit removal). Owner-approved this queuing.
- **HD-SAL-01 RULED 2026-10-23** — salvage, destruction & contents handling: **Option A** (salvage-to-owner with provenance continuity; contents sub-rule selection deferred to BAL-001 — owner picks from salvage-to-owner / world-debris / destroyed-with-structure, no default). **INV-010 CONFIRMED** in conditional reading under no-decay model (c/d bind destruction/loss systems; a attaches to repair-demand sources). **INV-SAL-01 APPROVED** as new TEST-001 invariant: "salvage returns meaningful ownership-relevant substance; removal contents resolved by stated rule." All four scope confirmations upheld (params → BAL-001; provenance-engine DEFERRED; 6e(b) separate; SAFE-001/002 informed not closed). Backlog 6d closed; backlog 6e(a) already closed by HD-BLD-03; backlog 6e(b) remains open. Contents sub-rule selection (salvage-to-owner vs world-debris vs destroyed-with-structure) remains a BAL-001-gated follow-up.
```
