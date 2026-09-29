# TCIndustries — OQ-011 Acquisition Half / PROG-003 Skill Acquisition — Human Ruling Record (HD-PROF-03)

**Status:** HUMAN RULED (PARTIAL) 2026-09-28 (owner decision recorded this
session). The acquisition model is ruled; INV-011c and the scope confirmation are
**NOT** ruled and are recorded as open. This document preserves the owner's
decision verbatim for filing.
**Filing state:** decision record filed in `proposals/` by the recorder
(Freebuff/OpenCode, per owner's instruction). Ready-to-file register entry in §3.
**Author:** OpenCode (space-bunny-free) — ruling-session recorder, 2026-09-28.
**Assessor:** none (self-recorded; no independent assessment).
**Gives effect to:** `proposals/TCIndustries_Skill_Acquisition_Proposal_Drafter_v1_2026-09-19.md` (retained unaltered as provenance).
**Ruling ID:** **HD-PROF-03** — register's HD-<Area>-<NN> pattern (area: PROF).
ID confirmed free at filing (HD-PROF-01 and HD-PROF-02 are respecialisation
rulings and are not amended).

---

## 1. The ruling (as decided by the owner)

| # | Question | Owner decision |
|---|---|---|
| 1 | OQ-011 acquisition half / PROG-003 acquisition model | **Option A — Trainer-gated study.** |
| 2 | Candidate invariant INV-011c | **NOT RULED — OPEN.** See §1c. |
| 3 | Scope confirmation | **NOT RULED — OPEN.** The three sub-confirmations were not answered this session. |

### 1a — Owner's stated reason, verbatim

> *"Historicity. SWG had NPC trainers you could pay 1000 credits to to learn
> skills. You could also have other players teach you skills for free if they
> already had that skill."*

### 1b — Owner's elaboration of the model's mechanics, verbatim

Recorded because it specifies the acquisition rule more precisely than the option
label alone, and because the recorder had put an incorrect characterisation to
the owner (see §1d).

> *"In Star Wars Galaxies you chose a starting class and recieved appropriate
> starting equipment."*
>
> *"In SWG you picked up new classes by going to the class trainer. You pay
> money. No money no training. no training no gain skill. Another player can
> train you for free."*

**Structural content recorded from the above:**

1. A character is created into a **starting class** and receives appropriate
   starting equipment.
2. **Additional classes are acquired by attending a class trainer.**
3. The trainer path is **credit-gated**: no credits means no training.
4. **No training means no skill gain.** Skill is not obtainable outside the
   training path.
5. **A second, non-monetary path exists:** another player who already holds the
   class can train the character **for free**. The gate is therefore not purely
   monetary.

**Numeric discipline:** the 1000-credit figure is recorded as **Pre-CU SWG
historical precedent cited by the owner as design rationale.** It is **not** an
authorised TCIndustries value. All training, credit, and point magnitudes are
BAL-001-gated and remain undecided.

### 1c — INV-011c: what was asked, what was answered, what remains

The recorder raised the concern that trainer-gating might make some careers
unreachable, and quoted INV-011c — *"alternative specialisations remain reachable
from any starting state; no acquisition path locks a character out of any
legitimate career permanently"* — as the guard rail against exactly that.

**The owner addressed the substance directly** (§1b): a character is never locked
out, because training is always available — for a fee from an NPC trainer, or
free from a player who holds the class. No approval or decline of INV-011c was
given, however, and per-invariant granularity forbids inferring one from an
answer to the recorder's question. **INV-011c is recorded as NOT RULED.** It
requires an explicit approve/decline.

**Note for the ruling session:** the owner's elaboration establishes that the
model satisfies INV-011c's reachability condition, subject to one residual
dependency — a character needs *either* credits *or* a willing player trainer.
Because HD-ECO-01 rules a faucet-exceeds-sink money-supply policy, the credits
branch is not structurally scarce. This is recorded as context, not as a reason
to presume the owner's answer.

### 1d — Correction of a recorder mis-framing, recorded

The recorder asked: *"if you never trained as a weaponsmith, and trainers are the
gate, can you still become one?"* The owner replied: *"I dont understand about
your 'If you never trained as a weaponsmith'. question."*

Recorded for provenance accuracy: the question was poorly framed. It implied a
state where becoming a weaponsmith might be unavailable, which is **not** what
trainer-gating produces under this ruling — attendance at a trainer is always
available. The question is withdrawn. No ruling content turns on it.

## 2. What this ruling explicitly does NOT authorise

- Any numeric value: training costs, credit amounts, skill-point grants, the
  1000-credit SWG precedent figure, or the point-return arithmetic underlying
  respecialisation (BAL-001).
- Any **budget** or specialisation-cap model (OQ-001 — **EIC-coupled, and the
  HD-EIC-07 gate remains open and blocking it**).
- Any change to the respecialisation cost instrument. That is ruled by HD-PROF-02
  (free-drop: dropping a class-skill tier returns the points spent on that tier).
  HD-PROF-02 is not amended.
- Any change to PROF-004 (LOCKED: respecialisation must exist; no permanent
  traps). Unchanged.
- Any NPC design, dialogue, or trainer roster content. Trainer existence is
  ruled as an acquisition gate; the NPCs are not designed here.
- Any Provenance & Reputation work (DEFERRED, HD-EIC-08).
- Any EIC shaping (HD-EIC-05/07 in force). **Specifically: this ruling does not
  supply, imply, or satisfy the HD-EIC-07 structural boundaries.** A
  credit-gated entry path is a candidate structural commitment, not a ruling on
  it.
- Any implementation, simulation, or testbed work.
- Approval of INV-011c or the scope sub-confirmations.

## 3. Ready-to-file register entry (verbatim, for the Human Rulings Register)

```markdown
## HD-PROF-03 — OQ-011 Acquisition Half / PROG-003 Skill Acquisition Model (PARTIAL: model; INV-011c and scope open)

**Source:** Owner ruling session, 2026-09-28, on
`proposals/TCIndustries_Skill_Acquisition_Proposal_Drafter_v1_2026-09-19.md`;
decision record `proposals/TCIndustries_Skill_Acquisition_Proposal_Ruling_Record_HD-PROF-03_2026-09-28.md`.

**Status:** HUMAN-LOCKED (acquisition model only). Complements HD-EIC-01–09,
HD-GDD-01, HD-TST-01, HD-RET-01, HD-BLD-01, HD-SCOPE-01, HD-ITM-01, HD-BLD-02,
HD-PROF-01, HD-PROF-02, HD-ITM-02, HD-BLD-03, HD-ECO-01, which remain in force.
Supersedes nothing; amends nothing.

### Decision

1. **PRIMARY — Option A (trainer-gated study).** Owner's stated reason, verbatim:
   *"Historicity. SWG had NPC trainers you could pay 1000 credits to to learn
   skills. You could also have other players teach you skills for free if they
   already had that skill."* The owner's elaboration, recorded verbatim in the
   decision record §1b: a character is created into a **starting class** with
   appropriate starting equipment; **additional classes are acquired by attending
   a class trainer**; the trainer path is **credit-gated** — *"You pay money. No
   money no training. no training no gain skill"* — and **a second,
   non-monetary path exists**: *"Another player can train you for free."* The
   acquisition gate is therefore not purely monetary. Answers the OQ-011
   acquisition half / PROG-003 at the structural layer.
2. **NUMERIC DISCIPLINE.** The 1000-credit figure is recorded as Pre-CU SWG
   historical precedent cited as design rationale, **not** as an authorised
   value. All training, credit, and point magnitudes are BAL-001-gated and
   undecided.
3. **INVARIANT — INV-011c NOT RULED; REMAINS OPEN.** The owner addressed the
   reachability concern substantively, confirming that training is always
   available — for a fee from an NPC trainer, or free from a player who holds the
   class — but gave no explicit approve/decline. Per-invariant granularity: no
   approval is inferred from an answer to a recorder's question.
4. **SCOPE — NOT CONFIRMED, REMAINS OPEN.** The three sub-confirmations (no OQ-001
   budget model selected or implied; all magnitudes to BAL-001; respecialisation
   cost instrument remains with HD-PROF-02) were not ruled this session. Default
   boundaries remain in force: numerics are BAL-001-gated, and OQ-001 remains
   open and EIC-coupled.

**Explicitly does not authorise:** any numeric value (BAL-001, including the
1000-credit SWG precedent figure); any budget or specialisation-cap model (OQ-001,
EIC-coupled and blocked behind the open HD-EIC-07 gate); any change to the
respecialisation cost instrument (HD-PROF-02 free-drop stands) or to PROF-004;
any NPC, dialogue, or trainer-roster design; provenance or reputation work
(DEFERRED, HD-EIC-08); **any EIC shaping — this ruling does not supply or satisfy
the HD-EIC-07 structural boundaries**; any implementation, simulation, or testbed
work.

### Traceability
| Item | Source |
|---|---|
| Proposal given effect | `proposals/TCIndustries_Skill_Acquisition_Proposal_Drafter_v1_2026-09-19.md` (retained as provenance) |
| Decision record | `proposals/TCIndustries_Skill_Acquisition_Proposal_Ruling_Record_HD-PROF-03_2026-09-28.md` |
| Invariant | INV-011c — NOT RULED, remains a candidate (Invariant Register v0.1, OQ-011 entry) |
| Composing rulings relied upon | HD-PROF-02 (free-drop respec — the exit cost is zero, so a wrong class choice is recoverable); HD-ECO-01 (faucet-exceeds-sink policy — the condition under which the credit-gated training path stays affordable) |
| Prior rulings affected | None |
```

## 4. Filing-time acts (reserved; owner/coordinator executes at register filing)

1. Paste §3 into `governance/TCIndustries_Human_Rulings_Register.md` as part of **Section E**, after Section D.
2. Update `proposals/TCIndustries_Open_Questions_Register.md` §2 OQ-011 status cell → "RULED (acquisition: trainer-gated; cost: free-drop) — see HD-PROF-01, HD-PROF-02, HD-PROF-03; INV-011c and scope still open".
3. Annotate `proposals/TCIndustries_Invariant_Register_v0.1.md` OQ-011 entry → INV-011c "NOT RULED, remains candidate (HD-PROF-03, 2026-09-28)".
4. **Master GDD status line** for OQ-011 / PROG-003 (canonical edit — owner's act). Note: OQ-011's GDD row is already marked RULED by HD-PROF-01 + HD-PROF-02; this ruling extends that pointer to the acquisition half.
5. Update `proposals/TCIndustries_Continuous_Queue_Tracker.md` — skill acquisition exits the live queue; add a RULED line; record that the live queue is now **empty** and the 3–5 band requires new drafts.
6. `governance/project_memory.md` event line.
7. Run the **GR-006.5** mechanical check for HD-PROF-03.

## 5. Immediate consequences recorded by the recorder

- **Queue:** the skill-acquisition item exits the live queue. With OQ-009 also
  ruled, the live queue is **empty** (0 items) against the tracker's 3–5 band.
  This is a deliberate owner decision, not drift; the tracker's method requires
  new backlog candidates to be drafted to restore the band.
- **OQ-001 remains blocked**, and this ruling does not help it. OQ-001 (the
  specialisation budget/cap) is EIC-coupled and still waits on the HD-EIC-07
  structural boundaries. **Recording a credit-gated acquisition path is not a
  substitute for those boundaries and must not be read as one.**
- **6d interaction:** none. Salvage and destruction handling are unaffected.
- **Post-Phase-10 interaction:** none. Testbed direction remains undecided and
  separate.
- **Still open from this proposal:** INV-011c (one-word confirmation) and the
  three scope sub-confirmations.

---
*End of decision record. Authority = the owner's 2026-09-28 ruling as recorded in §1.*
