# TCI–B5–TW Comparative Project-Management Report

**Title:** Comparing the two TCI-vs-TW assessments, with Babylon 5 as third reference; improvements toward an autonomous LLM system for TCIndustries
**Date:** 2026-10-06
**Author:** Muse Spark 1.3 Free (`muse-spark-1.3-contributor-free`)
**Status:** ADVISORY — changes no authority, promotes nothing, amends no governance file

```yaml
provenance:
  author_llm: {name: "Muse Spark", version: "muse-spark-1.3-contributor-free"}
  created_date: "2026-10-06"
```

## 0. Premise correction (read first)

The request states that `TCIndustries_Autonomy_Option_Assessment_2026-10-06.md` "is a comparison of Babylon5CCGGameDesign and tcindustries project infrastructure." That is **incorrect**, verified by grep today:

- Line 19: "**Comparative assessment against the `TiwasProjects` `.pm` system**, with a recommendation."
- Line 24: "Subject B: `C:\temp\projects\TiwasProjects` (hereafter **TW**)"
- `Babylon|B5` matches: only line 83, a passing note that "a Babylon 5 comparison exists in TW; a mirrored TW comparison exists in the Babylon 5 repo."

So the two assessments compared here are **two independent assessments of the same pair (TCI vs TW)**:

- **Assessment A:** OrcaRouter Auto, `C:\temp\projects\tcindustries\investigations\TCIndustries_Autonomy_Option_Assessment_2026-10-06.md` (deep file-level audit, §§0–9 + Appendices A–B)
- **Assessment B:** Muse Spark 1.3 Free, `C:\temp\projects\TiwasProjects\docs\tcindustries-vs-tiwasprojects-comparative-project-system-report-2026-10-06.md` (independent, same verdict)

Babylon 5 enters legitimately as a **third reference system**, via:

- Big Pickle's `C:\temp\projects\TiwasProjects\docs\babylon5-system-comparison-report-2026-10-06.md` (B5 vs TW, with mirror at `C:\temp\projects\Babylon5CCGGameDesign\docs\reports\tiwas-system-comparison-2026-10-06.md`)
- B5's own `AGENTS.md:1-80` (autonomous, lightweight; provenance §1 + compaction §1a; adapted from Tiwas TTRPG governance "stripped of its heavy human promotion process," line 24–26)

This report therefore does three things: (§§2–4) compare A vs B, (§5–6) add B5 as the battle-tested third pole neither A nor B fully exploited, (§7) consolidate improvements for TCI.

## 1. Method

- Full prior read of Assessment A (585 lines) and Assessment B (this author's own report).
- Grep of Assessment A for `Babylon|B5|Subject B|TiwasProjects` (4 matches, cited above).
- Read of B5 `AGENTS.md:1-80`, `README.md:1-80`; listing of `C:\temp\projects\Babylon5CCGGameDesign` (top level), `docs/`, `investigations/`; file-match for `ledger|claim|task|heartbeat` at root.
- Read of Big Pickle B5-vs-TW report §§1–4 (lines 1–120).
- `Test-Path`/directory verification from the prior session: TCI has no `.pm/`; TW has `.pm/kernel/pmk.py`; TCI last commit `3ccc198`; TW kernel k-0.2.0.
- No files modified except this report. No kernels or gates executed; TW gate results quoted from records.

## 2. Assessment A (OrcaRouter) — summary

- **Verdict:** improve TCI in place via a **third option (C)**: transplant TW's kernel underneath TCI byte-identically, plus a new TCI document-check layer. Explicitly not (A) docs-only improvement, not (B) new-repo migration.
- **Evidence base (strongest of the three reports):** full directory census with sizes/last-write, git history + dirty tree + ahead/behind, largest-tracked-file check, full reads of TCI `AGENTS.md`/`README.md`/`governance/README.md`/`project_memory.md`/queue tracker/OQR/2026-09-28 corpus audit, full reads of TW `AGENTS.md`/`README.md`/`policy.yaml`/D-0002/roadmap/events/views/task card/incidents/patterns, kernel subcommand + source-size inventory, and **re-verification today of the 2026-09-28 defects** (V-01…V-11 table, §5.2 — V-01…V-05 still live).
- **Diagnosis:** three mechanisms — (1) irreducible human gate (~10 decisions: HD-EIC-07 + 8× §36 + partial slots + GR-006 slots + GR-005 merge); (2) authority decay with no detector (every V-01…V-05 defect is machine-checkable); (3) loop has nothing to do while waiting (queue 0, OQ-012 re-explanation owed 8 days, backlog un-draftable).
- **Anti-migration argument (decisive, from TCI's own record):** a prior prefix-match consolidation silently dropped 4 non-`TCIndustries_`-prefixed docs; migration reproduces that silent-drop mode across 160 path-coupled docs.
- **Implementation:** P0 preservation (content-hash corpus, verified-write for Register + memory, root hygiene, testbed-DB policy) → P1 kernel transplant with TCI policy values → P2 six-check TCI layer (`check-rulings`, `check-refs`, `check-id-unique`, `check-registers`, `check-metadata`, `check-boundaries`, each with real red witnesses) → P3 ruling packets (one screen, options + invariant consequences, marked default, pre-rendered filing, atomic multi-target file, ≤15 min/ruling). Plus "what NOT to import" list and revisit conditions.

## 3. Assessment B (Muse Spark) — summary

- **Verdict:** identical — improve in place (Option C). Independent evidence path, same dispositive reasoning (zero executable instruments; location is not the defect).
- **Evidence base (thinner than A):** top-level listings, `AGENTS.md`/`README.md`/`IDEA.md`/Register/D-0001/policy reads, two `explore` subagents, `Test-Path` + `git log -3` verification. No per-defect re-verification, no size/census data, no V-table.
- **Value-add over A:** independent convergence (two models, same verdict, different paths — weak but real corroboration); explicit acknowledgment of A as prior advisory; blind-spot naming of TW's own gaps (no status-token semantics, no cross-register equality, no metadata lint, no ruling-packet generator, no merge queue).

## 4. A vs B — agreement and divergence

| Point | A (OrcaRouter) | B (Muse Spark) | Standing |
|---|---|---|---|
| Verdict (C: transplant, no move) | ✅ | ✅ | **Convergent** |
| Cause = zero instruments, not location | ✅ | ✅ | Convergent |
| Human gate irreducible (~10 decisions) | ✅ quantified | ✅ | Convergent |
| V-01…V-05 still live, machine-checkable | ✅ re-verified today | cited from A, not re-verified | A stronger; B dependent |
| Migration silent-drop risk (prefix-match precedent) | ✅ decisive argument | ✅ cited | Convergent, A primary |
| P0→P3 phased plan + revisit conditions | ✅ detailed | ✅ summarised | A normative; B endorses |
| "What NOT to import" guardrails | ✅ 4 items | ✅ cited | Convergent |
| B5 battle-tested lessons | ⚠️ one-line pointer only | ⚠️ same | **Gap in both — filled in §5–7 below** |
| TW staleness awareness | current (k-0.2.0, 8 events, R1–R15) | current | Both current; Big Pickle report is not (see §5) |

**Net:** B corroborates A; A remains the evidentiary authority. Neither assessment mined B5 for transferable mechanisms. That is this report's contribution.

## 5. Third pole: Babylon 5 (what neither assessment used)

B5 (`C:\temp\projects\Babylon5CCGGameDesign`) is a Java 6 CCG implementation run by ~200 autonomous LLM identities under an executable file-based protocol (`.agent/`), per Big Pickle's 2026-10-06 census: **1,153 ledger rows / 1,153 distinct IDs / 0 duplicates** (DONE 963 · BLOCKED 149 · OPEN 16 · VOID 17 · SUPERSEDED 7 · CLAIMED 1), **1,299 close-out reports**, **1,023 advisory lessons across 352 namespaces**, **472 decision-log entries (1.56 MB)**, **47 coordination tools each with exit-code contracts, most with `-SelfTest`/red-witness fixtures**, 10 live claims / 209 heartbeats, 92 commits. B5 `AGENTS.md:22-26` declares itself fully autonomous on shared live files, adapted from TTRPG governance minus the heavy human promotion process, with mandatory provenance (§1) and the §1a assessor-compaction rule (one entry per agent, pass counts, append-only at entry level, no retro-compaction).

Transferable findings (Big Pickle F1+; verified against B5 `AGENTS.md` shape):

1. **Gates are executable, self-tested, hostile to themselves.** Exit-code contracts (0/1/2), `-SelfTest` fixtures, red-witness mutants that must fail before trust; a never-executed green tool was itself a filed defect. TW-at-time-of-Big-Pickle had `tests/` empty and one `sentinel`; TW since adopted R1–R15 (k-0.2.0, 18 tests with red witnesses) — but TCI has **none of this**.
2. **Verified-write / compare-and-swap for shared files.** B5 measured read-modify-write destroying concurrent appends while reporting success, and answered with re-read-and-refuse-success tooling ("an assertion you read past is not a gate"). Directly applicable to TCI's Register + `project_memory.md` five-act filing ritual and GR-006 atomic-filing obligation.
3. **Tool-stamped clocks + payload-vs-clock acceptance.** B5-1931 forensics (UTC+13 hand-stamps reading LIVE forever). TCI has no clock discipline; ruling/filing dates are hand-written prose.
4. **Three-signal liveness (claim + heartbeat + report; absent ⇒ UNKNOWN, never STALE) + recorded reaps.** TCI's queue-emptiness and Cycle-14 authorised-but-unevidenced work are liveness failures with no detector.
5. **Failure IDs + lessons at point of use.** Every B5 defect carries a `B5-####` id; lessons embedded in `AGENTS.md` §5; mandatory one-line "Reusable lesson" at close-out feeding PATTERNS. TCI has excellent narrative failure history but **no incident objects**; TW has `INC-0001`/`PAT-0001` shape but one instance.
6. **Human surface that works.** B5: `human-decision-brief.md` + `human-approved (B5-NNNN)` annotations, no budget, but exercised at scale. TW: budget + 72h hold specified, unexercised. TCI: ad hoc per-ruling round-trips. The synthesis is TW's bounds + B5's exercised briefing format = the ruling packet both A and B recommend, now with a proven template to copy.
7. **Staleness warning:** Big Pickle describes TW at 508 lines / 4 events / 0 tests / 1 decision — TW is now k-0.2.0 / 8 events / 18 tests / D-0002. Quote Big Pickle's TW column as historical, not current. Its B5 column stands.

## 6. Three-way contrast

| Dimension | TCI (content, no machine) | TW (machine, no traffic) | B5 (machine + traffic) |
|---|---|---|---|
| Executable gates | ❌ zero | ✅ sentinel + citations + root + 18 tests (k-0.2.0) | ✅ ~20 instruments, self-test + red witnesses |
| Shared-file safety | ❌ hand-pasted five-act filing | ✅ append lock + hash chain (views/decisions/bundles lack verify-write) | ✅ CAS re-read verified append, measured clobber history |
| Clock discipline | ❌ hand dates | ✅ tool-stamped + 3-signal liveness rule | ✅ tool-stamped + payload-vs-clock forensics |
| Liveness/reaping | ❌ queue 0 undetected; authorised work unevidenced | 🟡 leases/fencing/quarantine; sweep weak | ✅ 3-signal rule + reap notes at scale |
| Verification independence | ❌ self-verification terminal | 🟡 V0–V4 ladder, V1/V2 unexecuted | ❌ self-verification only (B5's worst gap) |
| Authority model | ✅ status tokens + human rulings (violation-free 2026-09-28) | ✅ R0–R4 + two-key self-modification | ✅ status-by-location + claims-are-authority |
| Failure memory | ✅ narrative, no objects | 🟡 INC/PAT shape, 1 instance | ✅ IDs + PATTERNS + point-of-use lessons |
| Human surface | ❌ ad hoc | ✅ budgeted, unexercised | ✅ exercised, unbudgeted |
| Provenance | ✅ rich conventions | ✅ global rules, sparse use | ✅ frontmatter + compaction at scale |
| Content asset | ✅ 22 rulings, 28 invariants, 90k-line corpus | 🟡 bootstrap | ✅ engine + 643-check conformance + card pool |

## 7. Suggested improvements to TCI (consolidated, prioritized)

Ordered so each phase is independently useful with an exit test. R3 two-key (agent proposal + your approval) per TW convention. Items marked **[+B5]** are new beyond Assessments A/B, sourced from §5.

**P0 — Preservation (no decision from you; do first).**

- P0.1 Content-hash the corpus (`canonical/ governance/ proposals/ investigations/ archive/`); fail on unexplained change.
- P0.2 Verified-write for Register + `project_memory.md`: write → re-read → compare → refuse success otherwise **[+B5]** (CAS pattern, not just re-read).
- P0.3 Root hygiene: quarantine `swg.db*`, `.freebuff/project-id`, 0-byte `fs.statSync(x).isFile()))` fossil; `check-root` green.
- P0.4 Move 2,168 MB testbed-DB exclusion from vendored nested `.gitignore` into project policy.
- P0.5 **[+B5]** Tool-stamp all future filing/claim timestamps (tool clock, payload-vs-clock acceptance); hand dates become advisory. Exit: a future-stamped fixture is rejected by the checker.

**P1 — Transplant kernel (additive, reversible).**

- Copy `pmk.py` (stdlib-only) + `policy.yaml` + `views/ events/ memory/ registry/` beside TCI's seven directories; no existing governance doc edited.
- TCI policy values: `stdlib` only; ceiling R2; any `canonical/`/`governance/` write = R3; status promotion = R3; number-invention forbidden + machine-checked.
- Adopt interrupt budget (10/wk, 15 min/unit) + 72h default-hold, then populate the queue with the real backlog: HD-EIC-07, 8× §36, GR-006 slots, GR-005 merge.
- Exit: `sentinel` + `check-root` + `check-citations` green on fresh clone; `views/*` stamped `measured=`.

**P2 — TCI check layer (the new code; each with a real red witness).**

- `check-rulings` (V-01/V-02: every cited `HD-*` has a dedicated record) · `check-refs` (V-03/V-04: every asserted path resolves) · `check-id-unique` (V-05) · `check-registers` (OQ/invariant equality across GDD §31, OQR, tracker) · `check-metadata` (GR-004 frontmatter) · `check-boundaries` (no numerics / no resurrected non-canon).
- Declare known-red with inline finding IDs (B5 practice); never ship the layer as green.
- **[+B5]** Every check gets exit-code contract + `-SelfTest` + red-witness mutant; a never-executed green check is itself a filed defect.

**P3 — Liveness + lessons [+B5] (neither A nor B specified these; B5 proves them).**

- Three-signal liveness for every authorised-but-unexecuted item (claim + heartbeat + close-out; absent ⇒ UNKNOWN with recorded reap). Cycle-14 class cannot recur silently.
- Incident objects (`INC-####`) + advisory patterns (`PAT-####` with `review_by`), one per sentinel/check finding; mandatory one-line reusable lesson on every ruling filing.
- Failure IDs embedded at point of use in the check messages (B5-#### style), so the next audit starts from IDs, not prose.

**P4 — Reshape the ask (fixes "I am slow").**

- Ruling packet per open decision, in B5's exercised `human-decision-brief` format under TW's budget: one screen, exact options + invariant consequences, marked recommended default (no authority), pre-rendered register entry; one command files all dependent acts atomically (GR-006 mechanical at last). Target ≤15 min/ruling.
- Companion act (owed regardless): produce the OQ-012 re-explanation requested 2026-09-28.

**Explicitly NOT imported:** unread-policy values; bundle ritual for one-line checks; premature merge queue (until ≥2 agents work concurrently); any change to `AGENTS.md` status semantics; B5's self-verification-only habit (keep TW's V-ladder + never-verify-own-work).

**Revisit conditions:** check layer finds nothing beyond V-01…V-05 → stop at lightweight (A); GDD rewrite wanted → scoped v1.2 status-patch, not migration; genuine concurrency → full lease/fence/merge; EIC blocked 60 more days → one ruling de-scopes EIC, no migration.

## 8. Direct answers

1. **Is Assessment A a B5-vs-TCI comparison?** No (§0). It is TCI-vs-TW and remains the evidentiary authority on that pair.
2. **Do A and B agree?** Yes on verdict, cause, gate inventory, and plan shape; A carries the evidence, B the independent convergence. Trust A for facts, the pair for direction.
3. **Should TCI move to a new autonomous project?** No. Transplant TW's kernel + TCI check layer + B5's verified-write/clock/liveness/lesson practices into TCI in place. Corpus never moves except by history-preserving `git clone`.
4. **What is genuinely new here vs A/B?** P0.5 (clock discipline), P2 self-test/red-witness requirement, and all of P3 (liveness, incident/pattern objects, failure IDs) — the B5-sourced half of the autonomy story that the TCI-vs-TW framing could not see.

## 9. Caveats

- Single-day snapshot; all three trees drift; Big Pickle's TW column is stale (pre-R1–R15) while its B5 column stands.
- No kernel, gate, or B5 tool executed here; TW/B5 gate results quoted from their records — re-run before relying.
- TW machinery self-tested, near-zero traffic; B5 machinery traffic-tested, self-verification-only; TCI content traffic-tested, machine-zero. Discount each accordingly.
- No design opinion expressed; no EIC boundary, number, or architecture proposed.
- This report has no independent assessor yet — assess by a different model or yourself before structural action.

*End. Author: Muse Spark 1.3 Free (`muse-spark-1.3-contributor-free`), 2026-10-06. ADVISORY — no authority; no status changed; no file amended except this report.*
