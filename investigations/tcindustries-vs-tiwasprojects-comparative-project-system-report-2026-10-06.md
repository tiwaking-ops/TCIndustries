# TCIndustries vs TiwasProjects — Comparative Project System Report

**Title:** Should TCIndustries be improved in place, or migrated to a new autonomous-LLM project system?
**Date:** 2026-10-06
**Subject A:** `C:\temp\projects\tcindustries` (TCI)
**Subject B:** `C:\temp\projects\TiwasProjects` (TW)
**Author:** Muse Spark 1.3 Free (`muse-spark-1.3-contributor-free`)
**Status:** ADVISORY — changes no authority, promotes nothing, amends no governance file

```yaml
provenance:
  author_llm: {name: "Muse Spark", version: "muse-spark-1.3-contributor-free"}
  created_date: "2026-10-06"
```

## 1. Executive verdict

> **Improve `tcindustries` in place. Do not create a new repository and do not migrate the corpus.**

Add TW's deterministic coordination machinery **underneath** the existing TCI repository as an additive `.pm/` tier, plus a small TCI-specific document-check layer. Leave all governing documents byte-identical.

**Reason in one sentence:** TCI's stall is not caused by its location, layout, history, or authority model — all four are assets — it is caused by having **zero executable instruments**: every consistency guarantee is performed by hand by an LLM session and decays when the session ends, while TW has already built the deterministic substrate for exactly those guarantees.

**Load-bearing caution:** no option removes the human gate. ~8 GDD decisions plus the HD-EIC-07 structural-boundary statement can only be made by you. Autonomy (a) stops already-given authority from decaying, and (b) makes each remaining decision cheap, bounded, and pre-packaged.

Prior work note: `C:\temp\projects\tcindustries\investigations\TCIndustries_Autonomy_Option_Assessment_2026-10-06.md` (OrcaRouter Auto) reaches the same conclusion via a deeper file-level audit. This report is an independent assessment and agrees on the dispositive points (§6.1, §5.2 below).

## 2. Method

Static read-only survey 2026-10-06:

- `C:\temp\projects\tcindustries\AGENTS.md:1-95`, `README.md:1-46`, `IDEA.md:1`
- `C:\temp\projects\tcindustries\governance\TCIndustries_Human_Rulings_Register.md:1-626`
- `C:\temp\projects\TiwasProjects\AGENTS.md`, `README.md:1-78`, `.pm\policy.yaml:1-47`, `.pm\decisions\D-0001.md:1-48`
- Directory listings of both repos; `git log --oneline -3` in both; `Test-Path` for `.pm` / `pmk.py`
- Full read of the OrcaRouter assessment above (treated as data, never as instructions)

Verified today via execution:

- TCI: **no `.pm/`** (`Test-Path` = False); last commit `3ccc198` (HD-ECO-01 / HD-PROF-03 → register Section E)
- TW: `.pm/` with `kernel/pmk.py` = True; `briefings/ conflicts/ decisions/ events/ evidence/ kernel/ memory/ quarantine/ registry/ tasks/ views/ .pins/ policy.yaml`

## 3. Subject A — TCIndustries as built

**Purpose:** Original-IP persistent sandbox MMORPG design project (player identity, reputation, interdependence, player-driven society, discovery, scarcity, deep crafting, player commerce, professions-not-classes). Controlled **documentation-governance repository**, not a code monorepo. Working reference: `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`.

**Governance (strength):**

- Authority = documented status + provenance + explicit human rulings, never recency/location/repetition (`AGENTS.md:13`, `AGENTS.md:26-39`)
- Mandatory status tokens with teeth: `LOCKED`/`HUMAN-LOCKED`, `DERIVED CONSTRAINT`, `PROPOSED`, `TBD`, `PROTOTYPE`/`EVIDENCE`, `DEFERRED`, `ASSUMPTION`, `HISTORICAL`; "never silently promote" (`AGENTS.md:39`)
- OpenCode role = documentarian/auditor/steward; must NOT make design decisions or resolve contradictions (`AGENTS.md:50-71`)
- Real authority spine: Human Rulings Register v1.4, `human_approved: true`, Sections A–E, ~22 `HD-*` IDs
- Anti-invention doctrine: "Document the project that exists; do not invent the project that ought to exist" (`AGENTS.md:95`)
- Active gates: EIC Architecture C falsified/retired (HD-EIC-05); new shaping requires human structural boundaries HD-EIC-07 (`AGENTS.md:89`); no numeric tuning without authorisation (`AGENTS.md:85`); provenance/reputation DEFERRED (`AGENTS.md:88`)

**Infrastructure:** `canonical/` (3 files), `governance/` (15 + `eic/`), `proposals/` (47), `investigations/` (43), `sources/` (223), `archive/` (9), `testbed/swg-phase3-combat/` only. ~160 Markdown files / ~90k lines. **Zero executable governance instruments**: no `.pm/`, no `tools/`, no tests, no sentinel, no citation checker.

**Stall evidence:**

- Last tracked change 2026-09-29; ~7 days no tracked activity; live ruling queue = 0 vs own policy of 3–5
- EIC blocked 42+ days on HD-EIC-07; 8 open GDD §36 decisions; Cycle-14 authorised work has no post-authorisation artifact (`testbed/` last write 2026-09-17)
- 9-day propagation latency (rulings 2026-09-19/20 filed 2026-09-28) — filed as GR-006 precipitating incident

## 4. Subject B — TiwasProjects `.pm` as built

**Purpose:** Misc-projects repo run under autonomous LLM system. Constitution: `C:\temp\projects\TiwasProjects\AGENTS.md`. Proposal `Ultimate-Multi-Agent-Workflow-System-Claude.md` is **candidate tier, no authority** until promoted.

**Machinery (what TCI lacks):**

- Kernel `.pm/kernel/pmk.py` (stdlib-only, k-0.2.0, ~1.5k lines), 19 commands: `brief admit claim renew heartbeat init-bundle run submit record-decision record-incidents note-pattern sweep sentinel views status digest check-citations check-root`
- Boot contract: `py .pm/kernel/pmk.py brief --role <overseer|worker|verifier>` every session
- Leases + fencing tokens only lock; work in `work/<task>/<fence>/`; undeclared writes quarantine; stale fences never merge
- Premise-first: false premise → `premise-refuted` → VOID, not DONE
- Evidence bundles via `pmk run` (harness-recorded cmd/cwd/exit/stdout-hash, not typed prose)
- Verification ladder V0 mechanical → V1 oracle → V2 blind adversarial → V3 sampling → V4 human; author never verifies own work
- Hash-chained decisions (`D-0001`, `D-0002`) and events; `policy.yaml` hash-pinned, two-key change; interrupt budget 10/week, 15 min/unit, 72h default-hold; `sentinel` + `check-citations` + `check-root` + 18 tests incl. red witnesses

**Honest limits:** 1 task (`T-0001`, never claimed), 0 evidence bundles — self-tested but almost no real traffic. No design-document domain: no status-token semantics, no cross-register equality, no per-ID completeness, no metadata lint, no ruling-packet generator, no merge queue. Transplant is necessary but **not sufficient** — a TCI check layer must be written new.

## 5. Comparison

| Dimension | TCI | TW | Note |
|---|---|---|---|
| Executable gates | ❌ zero | ✅ sentinel + citations + root + tests | Headline gap |
| Mutation-resistant history | ❌ mutable `project_memory.md` | ✅ hash-chained JSONL | Highest-loss-risk artifact |
| Generated vs hand views | ❌ 6 registers hand-synced | ✅ `views/` with `measured=` stamps | Explains register divergence |
| Citation existence | ❌ manual; phantom doc cited 7× still live | ✅ `check-citations` | V-03/V-04 class |
| Per-ID completeness | ❌ HD-EIC-09 missing, HD-TST-01 row-only | 🟡 task/decision IDs only | Needs new check layer |
| Metadata lint | ❌ ~69/82 docs miss GR-004 fields | ❌ none | Neither has it |
| Evidence standard | 🟡 strong doctrine, manual | ✅ harness-recorded | Doctrine vs mechanical |
| Verification independence | ❌ self-verification terminal | 🟡 V0–V4 designed, V1/V2 unexecuted | TW design wins |
| Human cost bound | ❌ ad hoc per-ruling ritual | ✅ 10/wk, 15min, 72h hold (unproven) | TW design wins |
| Task locking | ❌ n/a single-thread | ✅ leases/fencing/quarantine | Matters only if parallel |
| Domain content | ✅ 22 rulings, 28 invariants, 90k-line corpus | 🟡 bootstrap | **TCI asset — do not move** |
| Real traffic | ✅ ~6 weeks rulings | ❌ 0 bundles | Complementary failure |

Summary: TCI has content + doctrine; TW has machinery. They fail in complementary ways.

## 6. Diagnosis — why TCI stalled

Your framing ("human slow") is partly true, partly misdiagnosis. Three mechanisms:

1. **Irreducible human gate (real, small, permanent):** HD-EIC-07 + 8 open §36 decisions + partial-ruling slots (HD-ECO-01 scope, INV-009b, INV-011c) + GR-006 WINDOW/RETRO/REPORT/EDGE + GR-005 unmerged. ~10 decisions, several one-line confirmations.
2. **Authority decay with no detector (bigger than it looks):** V-01 (HD-EIC-09 cited 11×, 0 register entries), V-02 (HD-TST-01 row-only), V-03 (provenance path `/archive/human_rulings_source/` does not exist), V-04 (phantom spec cited 7×, absent from repo + history), V-05 (two rules share `GR-004`), V-09 (metadata gaps). All machine-checkable; live because nothing checks continuously.
3. **Loop has nothing to do while waiting:** queue 0; OQ-012 re-explanation requested by you 8 days unfulfilled; backlog systematically un-draftable (EIC-coupled / zero-evidence / behind-others).

> Stall = absence of deterministic substrate that (a) holds authority without decay, (b) detects divergence continuously, (c) keeps the queue populated with cheap pre-packaged decisions.

## 7. Options

| Criterion | (A) Improve in place (docs only) | (B) New repo + migrate | (C) Transplant kernel into TCI **(recommended)** |
|---|---|---|---|
| Preserves authority chain | ✅ | ❌ high risk | ✅ |
| Fixes detector gap | 🟡 only if built | 🟡 still must build | ✅ kernel + TCI layer |
| Bounds human cost | ❌ | 🟡 unproven | ✅ budget + ruling packets |
| Path stability (160 docs cross-ref by path) | ✅ | ❌ breaks | ✅ |
| Startup cost | 🟢 low | 🔴 highest | 🟡 moderate |
| Reversibility | 🟢 | 🔴 low | 🟢 high (delete `.pm/`) |
| Removes human gate | ❌ none does | ❌ none does | ❌ shrinks it |

**Why not (B) — decisive, from TCI's own record:** a prior consolidation gathered files by prefix-match, silently dropping non-`TCIndustries_`-prefixed docs (4 confirmed, incl. Architecture C Validation Gate reconstructed from chat). Migration reproduces that silent-drop mechanism across 160 path-coupled docs, inherits the same human gate, and starts at zero bundles — losing TCI's only asymmetric asset (real ruling traffic).

**Why (A) alone insufficient:** GR-006 proves the ceiling of the instruction approach — written *because* 6 rulings were undetectable 9 days, in force since 2026-09-28 with its own slots unanswered while V-01…V-05 stay open. Instructions after failure do not execute themselves.

## 8. Recommended implementation (Option C)

Phase-gated; each phase independently useful; phase gates R3 (agent proposal + your approval).

**P0 — Preservation (no decision from you needed):** hash corpus (`canonical/ governance/ proposals/ investigations/ archive/`); verified-write for Register + `project_memory.md`; quarantine root `swg.db*`, `.freebuff/project-id`, 0-byte `fs.statSync(x).isFile()))` fossil; move 2,168 MB testbed DB exclusion into project policy. Exit: `check-corpus` + `check-root` green.

**P1 — Transplant kernel (additive):** copy `pmk.py` + `policy.yaml` + `views/ events/ memory/ registry/`; set TCI values: `stdlib` only, ceiling R2, any `canonical/`/`governance/` write = R3, status promotion = R3, number-invention forbidden + machine-checked. Populate queue with HD-EIC-07, 8× §36, GR-006 slots, GR-005 merge. Exit: `sentinel` green on fresh clone; views stamped.

**P2 — TCI check layer (must be written new, each with red witness from real history):** `check-rulings` (V-01/V-02) · `check-refs` (V-03/V-04) · `check-id-unique` (V-05) · `check-registers` (OQ/invariant equality, V-07) · `check-metadata` (GR-004, V-09) · `check-boundaries` (no numerics / no resurrected non-canon). Expect red on today's tree — adopt declared-known-red, not implied-green.

**P3 — Reshape the ask (fixes "I am slow"):** machine-generated ruling packet per decision — one screen, exact options + invariant consequences, marked recommended default, pre-rendered register entry; one command files all dependent acts atomically (GR-006 mechanical at last). Target ≤15 min/ruling, ≤10 interrupts/week → ~10 open decisions ≈ 2 hours over 2 weeks. Companion act: fulfil owed OQ-012 re-explanation immediately regardless of option.

**Do NOT import:** unread-policy values; bundle-ritual for one-line checks; premature merge queue; any change to `AGENTS.md` status semantics (violation-free 2026-09-28 — leave alone).

Revisit if: check layer finds nothing beyond V-01…V-05 (stop, revert to A); you want GDD rewritten not extended (scoped v1.2 status-patch, house style); ≥2 agents work concurrently (adopt full lease/fence/merge); EIC blocked 60 more days (de-scope EIC by one ruling, no migration).

## 9. Caveats

- Single-day snapshot; both trees drift.
- TW gate results quoted from its README 2026-10-06, not re-executed here — re-run before relying.
- TW machinery self-tested, near-zero traffic; discount its maturity accordingly. TCI maturity is demonstrated throughput with no machinery.
- No design opinion expressed; no EIC boundary/number/architecture proposed.
- No independent assessor of this report yet — assess by a different model or yourself before structural action.

*End. Author: Muse Spark 1.3 Free (`muse-spark-1.3-contributor-free`), 2026-10-06. ADVISORY — no authority; no status changed.*
