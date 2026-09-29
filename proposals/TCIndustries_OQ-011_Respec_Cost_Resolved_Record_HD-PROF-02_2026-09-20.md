# TCIndustries — OQ-011 Respec Cost Instrument Resolved — Human Ruling Record (HD-PROF-02)

**Status:** HUMAN RULED 2026-09-20 (owner decision recorded this session). This
document is the decision record; authority attaches to the owner's explicit
ruling, and this record preserves it verbatim for filing.
**Filing state:** decision record filed in `proposals/` by the drafter
(hard boundary: no `governance/` writes). The ready-to-file register entry is
in §3; `governance/TCIndustries_Human_Rulings_Register.md` is updated only by
the owner/coordinator at filing time.
**Author:** Buffy — Codebuff agent (Freebuff), ruling-session recorder, Cycle 14.
**Date:** 2026-09-20
**Gives effect to:** owner's in-session ruling (re-stated verbatim in §1); no
prior proposal is given effect, because the respec-cost question was deferred
under HD-PROF-01 and never reached a determination. This record closes the
defer.
**Ruling ID:** **HD-PROF-02** — register's HD-<Area>-<NN> pattern (area: PROF;
second of area).

---

## 1. The ruling (three parts, as decided by the owner)

|| # | Question | Owner decision |
||---|---|---|
|| 1 | OQ-011 cost instrument class (the deferred item from HD-PROF-01) | **FREE-DROP — a variant of Option C with no surrender cost.** Skill tiers are gained by spending skill points (per whatever acquisition model is later ruled for PROG-003/OQ-011 acquisition half). Dropping a class-skill tier returns the points the player spent on that tier. No fee, no time cost, no capability penalty beyond the tier drop itself. Reason stated by owner: authentic historical accuracy to Pre-CU SWG game rules. |
|| 2 | Candidate invariants | **INV-011a remains APPROVED** as a design invariant (TEST-001 class) per the existing record — under any respecialisation mechanism, name, appearance, item provenance, business ownership, organisation membership, and reputation history survive intact. **INV-011b ("cost is real but bounded") is NOT approved** in its cost form — the owner chose a no-cost respec model; the candidate is recorded as not-approved-under-this-model (it remains a candidate in the register because per-invariant granularity is the project convention, and a future respec ruling could adopt a different instrument). **INV-011c (acquisition-route equivalence) remains a candidate** — approval is independent and was not requested this session. |
|| 3 | Scope confirmation | **Confirmed:** no cost instrument, no surrender cost, no time cost, no economic cost. Magnitudes (point-return arithmetic, any transition rules) are deferred to the acquisition model ruling (PROG-003/OQ-011 acquisition half) and any BAL-001 pass; nothing here sets a magnitude. The acquisition model and the specialisation budget model (OQ-001) remain separately open. |

**Explicitly does not authorise:**
- Any cost beyond the tier-drop itself (no fee, no time window, no capability decay).
- Any change to PROF-004 (LOCKED: respecialisation must exist; no permanent traps) — unchanged.
- Any acquisition model (PROG-003 — open) or budget model (OQ-001 — EIC-coupled, HD-EIC-07-gated).
- Any implementation, numeric magnitude, or testbed work. GDD text edits and status-pointer updates are reserved controlled acts at register filing (§4).

---

## 2. What this ruling explicitly does NOT authorise

- Any respec cost instrument other than the free-drop model ruled above (a future ruling could amend this; nothing here locks the free-drop model as permanent canon beyond this ruling's scope).
- Any surrender-cost mechanism (the Pre-CU SWG "drop = get points back" model is the ruled shape; the historic SWG Skill Box architecture is not promoted — the ruling is about point-return behaviour, not box architecture).
- Any magnitude: point-return arithmetic, transition timing, or cap interaction — deferred to the acquisition model ruling and BAL-001.
- Any change to the acquisition half of OQ-011 / PROG-003 — that is a separate open question.
- Any implementation work; this is a design ruling (implementation is separately authorized per the Cycle-14 implementation authorization record for the other systems — the respec free-drop behaviour, if implemented, would be implemented under that same authorization, but no implementation commitment is made by this record alone).

---

## 3. Ready-to-file register entry (verbatim, for `governance/TCIndustries_Human_Rulings_Register.md`)

```markdown
## HD-PROF-02 — OQ-011 Respec Cost Instrument Resolved (Free-Drop)

**Source:** Owner ruling session, 2026-09-20, on owner's in-session ruling;
closes the OQ-011 cost-instrument defer left open by HD-PROF-01
(`proposals/TCIndustries_OQ-011_Ruling_Record_HD-PROF-01_2026-09-19.md`).

**Status:** HUMAN-LOCKED (respec cost instrument only). Complements HD-EIC-01–09,
HD-GDD-01, HD-TST-01, HD-RET-01, HD-BLD-01, HD-SCOPE-01, HD-ITM-01,
HD-BLD-02, HD-PROF-01, which remain in force. Supersedes nothing; resolves the
OQ-011 cost-instrument defer left by HD-PROF-01.

### Decision

1. **PRIMARY — Free-drop respec (variant of Option C, no surrender cost):**
   skill tiers are gained by spending skill points (per the later-ruled
   acquisition model); dropping a class-skill tier returns the points the
   player spent on that tier. No fee, no time cost, no capability penalty
   beyond the tier drop itself. Reason: authentic historical accuracy to
   Pre-CU SWG game rules.
2. **INVARIANT — INV-011a remains APPROVED** as a design invariant (TEST-001
   class): under any respecialisation mechanism, name, appearance, item
   provenance, business ownership, organisation membership, and reputation
   history survive intact. **INV-011b (cost is real but bounded) is NOT
   approved** in its cost form under this ruling — a no-cost model was chosen;
   INV-011b remains a candidate in the register (per-invariant granularity;
   a future respec ruling could adopt a different instrument). **INV-011c
   (acquisition-route equivalence) remains a candidate** — approval was not
   requested this session.
3. **SCOPE — confirmed:** no cost instrument beyond the tier drop; no magnitude
   set; acquisition model (PROG-003) and budget model (OQ-001) remain
   separately open; magnitudes deferred to the acquisition ruling and any
   BAL-001 pass.

**Explicitly does not authorise:** any respec cost other than the free-drop
model; any surrender-cost mechanism; any magnitude; any acquisition or budget
model selection; any implementation (separately authorized per the Cycle-14
implementation authorization record, where applicable).

### Traceability

|| Item | Source |
||---|---|
|| Ruling that closes the defer | Owner in-session ruling, 2026-09-20 (re-stated verbatim in this record §1) |
|| Prior related ruling | HD-PROF-01 (`proposals/TCIndustries_OQ-011_Ruling_Record_HD-PROF-01_2026-09-19.md`) — deferred the cost instrument; this record resolves it |
|| Invariant register state | `proposals/TCIndustries_Invariant_Register_v0.1.md`, OQ-011 entry — INV-011a approved; INV-011b/c candidates |
|| Prior rulings affected | None (resolves a defer; amends no prior ruling) |
```

---

## 4. Filing-time acts (reserved; owner/coordinator executes at register filing)

1. Paste §3 into `governance/TCIndustries_Human_Rulings_Register.md` (chronological entry, after the pending HD-PROF-01 entry).
2. Update `proposals/TCIndustries_Open_Questions_Register.md` §2 OQ-011 status cell → "RULED (cost instrument: free-drop) — see HD-PROF-01 + HD-PROF-02" (register convention: pointer only; ruling remains the authority).
3. Annotate `proposals/TCIndustries_Invariant_Register_v0.1.md` OQ-011 entry → INV-011a "remains APPROVED (HD-PROF-01, 2026-09-19)"; INV-011b "NOT approved in cost form under HD-PROF-02 (2026-09-20); remains candidate"; INV-011c "remains candidate".
4. Master GDD status-patch line for OQ-011 (canonical edit — owner's act, per GDD change-control).
5. `governance/project_memory.md` event line.

---

## 5. Immediate consequences recorded by the drafter (proposals/ side)

- **Queue:** the OQ-011 cost-instrument item is no longer queued as an open/deferred question — the cost instrument is ruled. The OQ-011 *acquisition half* (skill acquisition model) remains a separate live queued item (`proposals/TCIndustries_Skill_Acquisition_Proposal_Drafter_v1_2026-09-19.md`) and is not affected by this ruling.
- **HD-PROF-01 relationship:** HD-PROF-01 deferred the cost instrument and ruled INV-011a only. This record resolves the defer. HD-PROF-01 is not amended; its approved-invariant and scope statements stand.
- **Pre-CU SWG basis (recorded for the file, not canonised beyond the ruling):** the ruled model is anchored in the owner's stated Pre-CU SWG precedent (point spend → tier gain; tier drop → point return; no surrender cost). The Evidence Register's E-17 (Pre-CU skill-calc pattern, HISTORICAL) is the closest evidence reference; it is cited as historical pattern evidence only, not as TCIndustries design authority (SWG-002).
- **INV-011b status:** recorded as not-approved-in-cost-form under this ruling. If the owner ever revisits respec cost, INV-011b is a candidate for re-approval or further amendment; nothing here silently promotes it.

---

*End of decision record. Authority = the owner's 2026-09-20 ruling as recorded in §1.*
