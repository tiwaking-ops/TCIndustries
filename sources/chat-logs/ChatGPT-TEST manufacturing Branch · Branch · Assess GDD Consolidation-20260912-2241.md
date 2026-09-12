# TEST manufacturing Branch · Branch · Assess GDD Consolidation

**User:** Anonymous  
**Created:** 8/25/2026 18:27:49  
**Updated:** 9/12/2026 22:41:00  
**Exported:** 9/12/2026 22:41:25  
**Link:** [https://chatgpt.com/g/g-p-6a8605e707e4819196e35afd1d292bd4/c/6a8d35e2-ac90-83ec-b8b5-890541cf5f48?messageId=c08e0934-d1ae-47c7-8ed9-c602cb773b46&historySearchQuery=PLR-003](https://chatgpt.com/g/g-p-6a8605e707e4819196e35afd1d292bd4/c/6a8d35e2-ac90-83ec-b8b5-890541cf5f48?messageId=c08e0934-d1ae-47c7-8ed9-c602cb773b46&historySearchQuery=PLR-003)  

## Prompt:
25/08/2026, 15:07:49

Grok has created a consolidated TCIndustries Master GDD document. analyze, assess, advise &#x20;

[Attachment: TCIndustries_Master_GDD_v1.0_Consolidated-grok.md]

## Response:
25/08/2026, 15:07:51

> **Assessing Consolidated TCIndustries Master GDD**
>
> 
>
> Worked for a few seconds

Yes. I’ve reviewed the consolidated Grok GDD against the material available from the earlier TCIndustries GDD work. My assessment is that **Grok did a good consolidation job, but this should not yet be treated as the final canonical GDD without one human-controlled validation pass.**

## Executive assessment

**Overall: 8/10 as a consolidation document.  
~6/10 as a design-ready canonical GDD.**

The important distinction is that **the document is much better at preserving uncertainty than it is at resolving the underlying design problems**.

That is actually a good outcome for a consolidation pass.

Grok correctly establishes:

- a single Master GDD;
- explicit authority/status discipline;
- separation of LOCKED vs PROPOSED vs TBD;
- provenance of the source material;
- separation of SWG inspiration from TCIndustries rules;
- preservation of unresolved decisions;
- a sensible dependency-oriented design sequence.

Those are significant improvements over having several competing GDDs. The document explicitly states that agreement between LLM-generated sources does not constitute human approval, and that proposed material must not silently become canonical. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

### My main concern

**Grok has occasionally turned repeated LLM recommendations into something that looks more authoritative than the evidence warrants.**

The document says it did this conservatively, but there are still places where the distinction between:

> "This is genuinely locked by the project"

and

> "This is a very consistent design recommendation"

needs another human audit.

That matters enormously because this GDD is now intended to become the **single primary design authority**. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

---

# 1. What Grok got right

## A. The consolidation architecture is excellent

This is probably the strongest part of the document.

The introduction establishes:

> LOCKED = authoritative  
> PROPOSED = recommendation  
> TBD = unresolved  
> PROTOTYPE = evidence  
> DEFERRED = deliberately postponed  
> ASSUMPTION = working assumption  
> HISTORICAL = reference material

That is considerably stronger than simply merging several GDDs into one document. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

More importantly, Grok explicitly says:

> "Never silently promote PROPOSED, TBD, PROTOTYPE, DEFERRED, or ASSUMPTION material into LOCKED."

That should remain a **canonical governance rule**.

I would actually preserve this almost unchanged in the eventual master document.

---

# 2. The core TCIndustries vision survived consolidation well

The important design identity appears intact.

The consolidated document retains:

- persistent virtual world;
- citizen rather than chosen hero;
- player-driven economy;
- profession interdependence;
- emergence;
- player-created society;
- non-combat viability;
- flexible skill-based professions;
- player ownership;
- player commerce;
- player organisations;
- meaningful services.

These are appropriately treated as the architectural foundation rather than merely individual mechanics. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

That is important because the project could very easily drift into:

**"MMO + crafting system"**

instead of:

**"persistent player society in which crafting/economics/combat are components of the society."**

Grok preserved the latter.

### Verdict

**PASS - retain.**

---

# 3. The Anti-Goal framework is particularly strong

The consolidated document doesn't merely say:

> "Don't become Factorio."

It attempts to establish mechanical guardrails.

For example:

- automation cannot make players irrelevant;
- NPC commerce cannot replace player commerce;
- one character shouldn't efficiently master everything;
- combat shouldn't become the only meaningful progression route;
- the economy shouldn't reduce to spreadsheet optimisation.

The document explicitly ties these anti-goals back to system proposals. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

This is exactly the kind of thing a large MMO design needs.

### But there is an important caveat

An anti-goal is **not a mechanic**.

"Don't become Factorio" isn't sufficient.

You eventually need measurable invariants such as:

> Under defined conditions, a solo player cannot efficiently operate the entire production chain for high-value goods.

Grok has started moving in that direction with its testing section. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

That needs to become much more rigorous later.

---

# 4. The biggest improvement: it recognizes that the project is not ready for numeric balancing

This was absolutely correct.

Grok explicitly avoids inventing:

- skill caps;
- spawn rates;
- resource durations;
- prices;
- decay curves;
- combat coefficients;
- economic values.

Instead, it defines **structural shapes first** and leaves numbers for later simulation. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

I strongly agree.

Given the previous Tiwas work, this is exactly the right philosophy:

**Do not tune a mechanic whose structural behaviour has not been established.**

A mathematically precise bad design is still a bad design.

### Verdict

**PASS - this should become a formal TCIndustries design principle.**

---

# 5. The resource/crafting/manufacturing relationship is promising

The consolidated GDD has correctly identified the most important economic chain:

**Resource discovery → extraction → processing → crafting → manufacturing → retail → consumption/services**

That gives TCIndustries an actual systemic backbone.

The resource system is also correctly treated as more than:

> "iron + 10 = sword."

It is supposed to create:

- scarcity;
- geographic differentiation;
- information asymmetry;
- temporary opportunities;
- quality variation;
- trade;
- player reputation.

`TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

Likewise, differentiated crafting is explicitly locked:

> products of the same schematic should not simply become interchangeable commodities.

`TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

That is strongly aligned with the project's identity/reputation pillar.

---

# 6. However: there is one potentially dangerous design decision buried in the document

This needs attention before we treat the GDD as canonical.

The manufacturing section appears to move toward:

> **Factories can never produce Exceptional-tier output.**

This is presented in the earlier source material as a proposed early-lock direction, because it preserves the distinction between mass production and master craftsmanship.

Conceptually, I understand the reasoning.

But I **would not lock this yet.**

Why?

Because it may create an artificial dichotomy:

**Crafting = good**  
**Manufacturing = inferior**

rather than:

**Crafting = discovery/design/mastery**  
**Manufacturing = scale/consistency/capacity**

Those are not the same thing.

A manufacturing organisation should potentially be able to produce **extremely good** products, but perhaps it should not be able to reproduce the *creative/master-crafter optimisation process* indefinitely or autonomously.

That is a much richer design space.

### Recommendation

Keep the anti-automation principle LOCKED:

> Automation must not replace the economically/socially meaningful role of skilled players.

But keep the specific:

> "Factories cannot produce Exceptional"

**PROPOSED**, until it has been simulated.

This is an example of why the status discipline needs to remain ruthless.

---

# 7. The specialisation problem is correctly identified as the #1 foundational problem

Grok's OQ-001 is:

> What is the exact specialisation budget / skill-cap model?

That is absolutely the correct question.

The document correctly recognises that if the cap is too generous:

**everyone eventually becomes everything.**

Then:

**interdependence collapses.**

And once interdependence collapses:

**player society becomes optional.**

The GDD explicitly identifies this dependency. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

### I would elevate this even further

Before designing detailed profession trees, determine the **mathematical/structural meaning of "specialist"**.

Not:

> "How many skill points do players get?"

but:

> **What must a character be unable to do simultaneously for TCIndustries' economy to remain socially interdependent?**

That should be the actual S-1 question.

---

# 8. The GDD still has a major unresolved circular dependency

There is a deeper problem than the document currently acknowledges.

Consider:

**Specialisation**

depends on:

**Profession design**

which depends on:

**Crafting complexity**

which depends on:

**Resource attributes**

which depends on:

**Resource taxonomy**

which affects:

**Economic value**

which affects:

**Manufacturing**

which affects:

**Skill demand**

which feeds back into:

**Specialisation.**

So the proposed sequence:

> skills → resources → crafting → manufacturing...

is useful, but not truly linear.

It is a **design dependency graph**, not a simple sequence.

This matters because locking the skill model too early could be just as dangerous as locking resource mathematics too early.

### Recommendation

The next design stage should use **cross-system design packages**, not isolated system documents.

For example:

### Package A - Economic Interdependence Core

Resolve together:

- specialisation;
- resource quality;
- crafting quality;
- manufacturing;
- NPC substitution;
- player demand;
- service demand.

Then test the entire loop.

---

# 9. The prototype problem is still unresolved

This is the biggest factual weakness in the current GDD.

Grok explicitly says:

> "No prototype was inspected during this consolidation."

It therefore correctly retains prototype information as evidence/assumption rather than canon. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

That is good methodology.

But it means **we do not actually know how much of the project already exists.**

This matters because the previous project history references a Seed-2.1 TypeScript/Vite/React prototype.

The GDD itself acknowledges the issue:

> "Reconciliation of any existing prototype against this GDD."

is one of the remaining human decisions. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

### Recommendation

Do **not** immediately start designing another large batch of systems.

First establish an:

**Evidence Reconciliation Pass**

with:

| Evidence | Question |
|---|---|
| Master GDD | What is supposed to exist? |
| Prototype | What actually exists? |
| Prototype tests | What has been demonstrated? |
| Design decisions | What has actually been approved? |
| Audit | What can independently be verified? |

This will prevent the project from accidentally designing systems that have already been prototyped - or assuming prototype mechanics are canonical when they aren't.

---

# 10. The GDD's biggest structural weakness: "LOCKED" is still too broad

This is subtle.

The document has excellent status discipline **within individual rules**.

But there is still a distinction between:

### Locked principle

> Player interdependence is required.

and:

### Locked implementation consequence

> A particular skill-cap mechanism is required.

The former is safe.

The latter may not be.

I see some places where "LOCKED" is justified as a derived consequence of a pillar rather than an explicit human decision.

For example:

> CRFT-006 - No Instant Mastery of Everything  
> LOCKED (derived from interdependence and specialisation principles)

`TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

I would **not automatically classify derived implementation rules as LOCKED.**

The principle is locked.

The implementation should remain proposed until explicitly approved.

### I recommend adding a fourth authority distinction

Something like:

- **LOCKED - HUMAN** = explicitly approved by you.
- **LOCKED - DERIVED** = logically follows from an approved principle.
- **PROPOSED** = LLM/design recommendation.
- **TBD** = unresolved.

Even better:

**Do not call DERIVED rules LOCKED at all.**

Use:

> **DERIVED CONSTRAINT**

That eliminates ambiguity.

---

# 11. Provenance is correctly identified as important - but under-designed

The GDD repeatedly relies on provenance to create identity:

> "Who made this?"

That's excellent.

But provenance needs considerably more design.

For example:

- Does the original resource source remain visible?
- Does processing preserve provenance?
- Does a component remember its creator?
- Does a final product contain the identities of all major contributors?
- Can an organisation manufacture using another player's template?
- Who receives reputation?
- What happens when a company employs 50 crafters?
- Can a brand become more important than the individual?
- Can reputation be transferred?
- What happens when the player quits?

The current GDD correctly lists provenance depth as OQ-017. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

### Recommendation

Provenance deserves its own dedicated design pass because it connects:

**resource → processor → crafter → organisation → manufacturer → retailer → owner → reputation**

That could become one of TCIndustries' defining systems.

---

# 12. The "player society" pillar needs more mechanical depth

The GDD is strongest on:

**resources + crafting + economy**

It is less developed on:

**society.**

Cities, organisations, governance and services remain mostly TBD. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

That's understandable at this stage.

But eventually TCIndustries needs to answer:

> Why does a player need a society rather than simply needing a marketplace?

Those aren't equivalent.

A marketplace gives you:

> buy/sell.

A society gives you:

> trust, institutions, politics, cooperation, status, social obligations, collective projects, reputation, belonging.

The current document has the right aspiration but not yet the mechanical machinery.

---

# 13. The economic model is still conceptual rather than systemic

This isn't a criticism of the consolidation itself.

It's simply where the project currently stands.

The GDD says:

> player-driven economy

and:

> regional economies

but currency, sinks, stabilisation, market manipulation and vendor mechanics remain TBD. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

That means we cannot yet answer critical questions such as:

- Why does money retain value?
- What destroys money?
- What creates money?
- What prevents inflation?
- What prevents resource monopolies?
- What prevents market cornering?
- What happens when a major corporation controls a resource?
- How does a new player enter the economy?
- How does a poor player participate?
- What creates demand for ordinary-quality products?

Those are not merely "economy tuning" questions.

Some are **structural design questions**.

---

# 14. The GDD needs one more thing: explicit design invariants

Grok has started this correctly.

For example:

> a specialised weaponsmith should produce higher-quality weapons than a generalist under equivalent resource quality.

`TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

I recommend making this a major formal framework.

Every major system should eventually have:

### Design Principle
What we want.

### Design Constraint
What must never happen.

### Invariant
What must remain true under testing.

### Failure Condition
What tells us the design is broken.

### Test
How we measure it.

For example:

**Interdependence**

> Principle: specialists should matter.

> Constraint: no single character should efficiently dominate all major economic roles.

> Invariant: high-end production requires at least N distinct capability domains.

> Failure condition: economically optimal solo production exceeds X% of specialist-network production.

> Test: simulated economy.

That is much stronger than prose.

---

# 15. The recommended sequence is good, but I would change it

Grok recommends:

1. Confirm LOCKED items.
2. Specialisation.
3. Resources.
4. Crafting.
5. Manufacturing.
6. Ownership/provenance/vendors.
7. Cities/orgs.
8. Combat.
9. Transport.
10. Numeric simulation.

`TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

I would modify this.

## Recommended next sequence

### Phase 0 - Canonical Audit
Before designing anything:

**Audit the consolidated GDD itself.**

Determine:

- genuine human decisions;
- derived constraints;
- LLM proposals;
- unresolved contradictions;
- missing source material;
- prototype evidence;
- questionable LOCKED classifications.

### Phase 1 - Core Economic Interdependence

Design these together:

**Specialisation ↔ Resources ↔ Crafting ↔ Manufacturing**

This is TCIndustries' economic engine.

### Phase 2 - Provenance & Reputation

Design:

**Resource → creator → product → organisation → retailer → reputation**

This is TCIndustries' identity engine.

### Phase 3 - Commerce & Services

Design:

**Retail → demand → services → recurring player interaction**

This is the social/economic engine.

### Phase 4 - Society

Design:

**Organisations → property → cities → governance → institutions**

This is the civilisation engine.

### Phase 5 - World & Logistics

Design:

**Geography → transportation → regional markets → resource movement**

This makes the economic system spatial.

### Phase 6 - Combat

Combat should be designed **into the existing economy**, rather than designed as a separate game and then connected afterward.

### Phase 7 - Simulation

Only after the structural models exist:

- resource simulation;
- production simulation;
- economic simulation;
- specialisation simulation;
- player behaviour modelling.

### Phase 8 - Prototype reconciliation

Then compare:

**GDD ↔ prototype ↔ simulation**

and decide what survives.

---

# 16. One thing I would explicitly NOT do

I would **not** ask Grok, Claude, Gemini, etc. to simply "improve" this GDD next.

The consolidation phase has achieved its purpose.

The next phase should be **design adjudication**, not more synthesis.

Otherwise you risk this cycle:

> LLM A proposes  
> → LLM B expands  
> → LLM C consolidates  
> → LLM D validates  
> → repeated agreement starts looking like authority.

That is exactly the false-consensus problem the GDD itself is trying to prevent.

---

# 17. My recommendation for the current status

I would classify the document:

### **TCIndustries Master GDD v1.0 - Consolidation Candidate**

rather than:

### **TCIndustries Master GDD v1.0 - Fully Canonical**

The distinction is important.

It is **good enough to become the single working reference**, but not yet good enough to freeze as unquestionable project canon.

Grok itself acknowledges that several human decisions remain, including specialisation, resources, crafting, manufacturing, cities, combat, currency, respecialisation, exploit policy, and prototype reconciliation. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

---

# 18. What I would do next

I recommend **one dedicated Canonical GDD Audit before any new system design.**

That audit should produce four things:

### A. Canonicality Audit
Identify every statement currently marked LOCKED and classify it:

- **HUMAN-LOCKED**
- **DERIVED**
- **LLM-PROPOSED**
- **QUESTIONABLE**

### B. Consolidation Integrity Audit
Check whether Grok:

- lost anything important from the source GDDs;
- accidentally promoted anything;
- accidentally demoted anything;
- merged distinct concepts;
- removed useful unresolved alternatives;
- introduced new design decisions during consolidation.

### C. Dependency Audit
Map which unresolved decisions block which other decisions.

### D. Master GDD Readiness Verdict

Something like:

> **GREEN - suitable as working canonical reference**

or

> **YELLOW - suitable as working reference but requires corrections**

or

> **RED - consolidation introduced material canonical errors.**

---

## Bottom line

**Grok's consolidation is substantially successful.**

I would **keep the document** rather than regenerate it.

Its strongest achievement is not the actual game mechanics; it's the establishment of a disciplined **design authority/status framework** and preservation of uncertainty. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

Its biggest weakness is that some **derived design recommendations are wearing LOCKED clothing**, while several foundational systems remain insufficiently defined to validate whether the central promise of interdependence will actually work.

So my advice is:

> **Do not rewrite the GDD yet. Audit it first.**

After that audit, we should have a clean **Master GDD v1.1 Canonical Baseline**, and *then* start the first serious system-design pass - most likely **Specialisation + Resource + Crafting + Manufacturing as one interconnected design problem**.

That would be a much more rigorous progression than asking another LLM to produce yet another consolidated GDD.

## Prompt:
25/08/2026, 15:11:30

# TCIndustries Master GDD v1.0 — Canonical Audit

**Document:** Canonical GDD Audit of `TCIndustries_Master_GDD_v1.0_Consolidated.md`\
**Audit Date:** 2026-08-25\
**Auditor Role:** Senior Game Systems Architect / Design Auditor\
**Audited Document Version:** 1.0 — Consolidated Master GDD\
**Recommended Status of Audited Document:** **Consolidation Candidate** (not yet Fully Canonical)

---

## Executive Verdict

| Dimension                | Rating                       | Notes                                                                                    |
| ------------------------ | ---------------------------- | ---------------------------------------------------------------------------------------- |
| Consolidation quality    | Strong                       | Single coherent reference; excellent status discipline; uncertainty preserved            |
| Design readiness         | Incomplete                   | Core economic interdependence engine not yet structurally defined                        |
| Authority hygiene        | Mostly good, with exceptions | Several derived rules incorrectly wear LOCKED clothing                                   |
| Prototype reconciliation | Incomplete                   | No prototype inspected; AS-001 remains open                                              |
| Readiness for freeze     | **YELLOW**                   | Suitable as the single working reference; requires corrections before Canonical Baseline |

**Overall Recommendation:**\
Do **not** regenerate the GDD. Do **not** begin large new system design.\
Treat the current document as the working reference and produce a **Master GDD v1.1 Canonical Baseline** after applying the corrections and authority re-classifications identified in this audit.

---

## A. Canonicality Audit

Every statement currently marked **LOCKED** is classified below.

### Classification Legend

| Class                            | Meaning                                                                                                                                                                                                  |
| -------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **HUMAN-LOCKED**                 | Explicitly established by project brief / human direction / consistent core identity language across sources that is treated as foundational. Safe to retain as LOCKED pending final human confirmation. |
| **DERIVED CONSTRAINT**           | Logically follows from a HUMAN-LOCKED principle but is itself an implementation consequence. Should **not** be called LOCKED. Recommend reclassification.                                                |
| **LLM-PROPOSED (over-promoted)** | Appears primarily as consistent LLM recommendation. Should be demoted to PROPOSED.                                                                                                                       |
| **QUESTIONABLE**                 | Authority or wording is ambiguous; requires human ruling.                                                                                                                                                |

### Vision & Pillars

| ID      | Statement                                | Current Status | Audit Class  | Recommendation |
| ------- | ---------------------------------------- | -------------- | ------------ | -------------- |
| VIS-001 | Persistent player-driven virtual society | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| VIS-002 | Citizen rather than chosen hero          | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| VIS-003 | Player-driven economy                    | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| VIS-004 | Interdependence of professions           | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| VIS-005 | Emergence over scripting                 | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| VIS-006 | Anti-goals list                          | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| PIL-001 | Discovery and the Gold Rush              | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| PIL-002 | Identity and Reputation                  | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| PIL-003 | Interdependence                          | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| PIL-004 | Player-Created Society                   | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |

### Core Experience & World

| ID       | Statement                                     | Current Status | Audit Class  | Recommendation |
| -------- | --------------------------------------------- | -------------- | ------------ | -------------- |
| LOOP-002 | Non-combat long-term viability                | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| LOOP-003 | Progression expands options, not linear power | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| WRLD-001 | Persistent shared world                       | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| PLR-001  | Persistent character identity                 | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |
| PLR-003  | Persistent ownership                          | LOCKED         | HUMAN-LOCKED | Retain LOCKED  |

### Skills & Professions

| ID       | Statement                                            | Current Status | Audit Class  | Recommendation                         |
| -------- | ---------------------------------------------------- | -------------- | ------------ | -------------------------------------- |
| PROF-001 | Flexible skill-based professions (not rigid classes) | LOCKED         | HUMAN-LOCKED | Retain LOCKED                          |
| PROF-004 | Respecialisation must be allowed                     | LOCKED         | HUMAN-LOCKED | Retain LOCKED (exact costs remain TBD) |

### Resources, Crafting, Manufacturing

| ID       | Statement                                                   | Current Status                   | Audit Class            | Recommendation                                                                                                        |
| -------- | ----------------------------------------------------------- | -------------------------------- | ---------------------- | --------------------------------------------------------------------------------------------------------------------- |
| RES-002  | Resources create discovery, scarcity, events                | LOCKED                           | HUMAN-LOCKED           | Retain LOCKED                                                                                                         |
| CRFT-001 | Differentiated products (not pure commodities)              | LOCKED                           | HUMAN-LOCKED           | Retain LOCKED                                                                                                         |
| CRFT-006 | No Instant Mastery of Everything                            | LOCKED (derived)                 | **DERIVED CONSTRAINT** | **Reclassify to DERIVED CONSTRAINT or PROPOSED**. Principle is locked; this wording is an implementation consequence. |
| MFG-001  | Controlled automation — must not eliminate player relevance | LOCKED                           | HUMAN-LOCKED           | Retain LOCKED (principle)                                                                                             |
| MFG-004  | Anti-Factorio drift                                         | LOCKED (anti-goal reinforcement) | **DERIVED CONSTRAINT** | Reclassify to DERIVED CONSTRAINT. The anti-goal is already covered by VIS-006 / MFG-001.                              |

**Note on manufacturing “Exceptional” restriction:**\
The consolidated GDD does **not** contain a LOCKED rule that “factories can never produce Exceptional-tier output.” That specific restriction (present in some earlier source material) was correctly left out of LOCKED status. No action required on this point.

### Economy, Retail, Services, Combat

| ID       | Statement                          | Current Status     | Audit Class            | Recommendation                                     |
| -------- | ---------------------------------- | ------------------ | ---------------------- | -------------------------------------------------- |
| ECO-001  | Player-driven economy              | LOCKED             | HUMAN-LOCKED           | Retain (overlaps VIS-003; acceptable)              |
| ECO-003  | No pure spreadsheet optimisation   | LOCKED (anti-goal) | **DERIVED CONSTRAINT** | Reclassify to DERIVED CONSTRAINT                   |
| RET-001  | Player retail must exist           | LOCKED             | HUMAN-LOCKED           | Retain LOCKED                                      |
| SERV-001 | Meaningful non-combat services     | LOCKED             | HUMAN-LOCKED           | Retain LOCKED                                      |
| CMBT-001 | Combat as one lifestyle among many | LOCKED             | HUMAN-LOCKED           | Retain LOCKED                                      |
| CMBT-003 | No Forced Combat Path              | LOCKED (derived)   | **DERIVED CONSTRAINT** | Reclassify. Already covered by LOOP-002 + VIS-006. |

### Society, Progression, Emergence, Safety

| ID                | Statement                                  | Current Status    | Audit Class                | Recommendation                                                             |
| ----------------- | ------------------------------------------ | ----------------- | -------------------------- | -------------------------------------------------------------------------- |
| BLD-001           | Player-owned structures                    | LOCKED            | HUMAN-LOCKED               | Retain LOCKED                                                              |
| SOC-001           | Player organisations                       | LOCKED            | HUMAN-LOCKED               | Retain LOCKED                                                              |
| PROG-001          | Option expansion (restatement of LOOP-003) | LOCKED            | HUMAN-LOCKED / restatement | Retain or fold into LOOP-003                                               |
| PROG-002          | Multiple viable paths                      | LOCKED (derived)  | **DERIVED CONSTRAINT**     | Reclassify. Covered by LOOP-002.                                           |
| EMRG-001          | Conditions for emergence                   | LOCKED            | HUMAN-LOCKED               | Retain LOCKED                                                              |
| SAFE-001          | Ownership/transfer safeguards              | LOCKED            | HUMAN-LOCKED               | Retain LOCKED                                                              |
| EXP-001           | Data-driven extensibility                  | LOCKED / PROPOSED | QUESTIONABLE               | Keep principle as HUMAN-LOCKED aspiration; implementation remains PROPOSED |
| SWG-001 / SWG-002 | Inspiration only; no Star Wars IP          | LOCKED            | HUMAN-LOCKED               | Retain LOCKED                                                              |
| PROT-001          | GDD rules authoritative over prototype     | LOCKED            | HUMAN-LOCKED               | Retain LOCKED                                                              |
| BAL-001           | Avoid premature numeric tuning             | LOCKED            | HUMAN-LOCKED               | Retain LOCKED — elevate to formal design principle                         |

### Summary of Required Reclassifications

The following should be demoted from LOCKED to **DERIVED CONSTRAINT** (or PROPOSED) in v1.1:

1. CRFT-006 — No Instant Mastery of Everything
2. MFG-004 — Anti-Factorio Drift (redundant with VIS-006 + MFG-001)
3. ECO-003 — No pure spreadsheet optimisation
4. CMBT-003 — No Forced Combat Path
5. PROG-002 — Multiple viable paths

These are true statements that follow from higher principles, but labelling them LOCKED creates false precision about what has been explicitly decided versus what is logically implied.

**Recommended new status vocabulary addition for v1.1:**

| Status                 | Meaning                                                                                                                                                                         |
| ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **DERIVED CONSTRAINT** | Follows from one or more HUMAN-LOCKED principles. Binding as a design constraint, but not an independent human decision. May be refined when the parent principles are refined. |

---

## B. Consolidation Integrity Audit

### What Grok got right

- Established a single Master GDD with clear authority/status discipline.
- Explicitly stated that LLM agreement ≠ human approval.
- Preserved the four pillars and citizen fantasy without dilution.
- Correctly avoided inventing numbers.
- Correctly left city, combat, currency, transport, and most mechanical detail as TBD/PROPOSED.
- Did **not** promote the “factories cannot produce Exceptional” rule.
- Preserved open questions and assumptions as registers.
- Separated HISTORICAL SWG material cleanly.
- Recognised specialisation budget (OQ-001) as foundational.

### Potential integrity issues

| Issue                           | Severity                 | Finding                                                                                                                                                                      |
| ------------------------------- | ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Derived rules marked LOCKED     | Medium                   | CRFT-006, CMBT-003, PROG-002, MFG-004, ECO-003 wear LOCKED clothing they have not earned as independent decisions.                                                           |
| Overlap / restatement inflation | Low                      | Several LOCKED items are near-duplicates of pillars (ECO-001 ≈ VIS-003; PROG-001 ≈ LOOP-003). Acceptable for readability; can be tightened later.                            |
| Lost material                   | Low–Medium               | No major unique LOCKED content appears lost. Some richer provenance discussion and manufacturing quality-tier language from sources was correctly left PROPOSED or excluded. |
| Accidental new design           | Low                      | Consolidation did not invent major new mechanics. The design sequence is a recommendation, not a rule.                                                                       |
| Prototype blindness             | High (process)           | Document correctly flags that no prototype was inspected. This remains a blocking process gap for any claim of implementation awareness.                                     |
| Society pillar thinness         | Medium (design maturity) | Aspiration is LOCKED; mechanical machinery (governance, institutions, trust, collective projects) is almost entirely TBD. Not a consolidation error, but a design gap.       |

### Material that was correctly excluded or kept non-LOCKED

- Specific skill-point costs, spawn rates, durations, prices, decay curves.
- Any claim that factories are permanently barred from top-tier quality output.
- Detailed profession trees.
- Concrete city governance models.
- Concrete combat systems.
- Concrete currency/sink models.

---

## C. Dependency Audit

The current recommended sequence in the GDD is useful but overly linear. The real structure is a dependency graph.

### Critical Circular / Tightly Coupled Cluster
```
Specialisation Budget (OQ-001)
        ↕
Resource Attributes & Quality (OQ-003)
        ↕
Crafting Differentiation & Experimentation (OQ-004)
        ↕
Manufacturing Constraints (OQ-005)
        ↕
NPC Substitution Limits
        ↕
Player Demand / Service Demand
        ↕
Economic Value Signals
```

This cluster is the **Economic Interdependence Core**. Resolving any single node in isolation risks locking the wrong shape.

### Blocking Relationships

| Unresolved Decision              | Blocks or Strongly Constrains                                                                                    |
| -------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| OQ-001 Specialisation budget     | Almost everything downstream: profession design, crafting depth, manufacturing viability, interdependence claims |
| OQ-003 Resource attributes       | Crafting outcomes, product differentiation, experimentation model, provenance value                              |
| OQ-004 Crafting experimentation  | Crafter reputation, manufacturing vs. craft distinction, identity pillar                                         |
| OQ-005 Manufacturing constraints | Anti-automation claims, organisation vs. individual production, scale economics                                  |
| OQ-017 Provenance depth          | Identity/reputation pillar realisation, brand vs. individual, organisation employment models                     |
| OQ-009 Currency & sinks          | Economic stability, vendor design, long-term demand                                                              |
| OQ-007 City/governance           | Player-Created Society pillar realisation                                                                        |
| Prototype evidence (AS-001)      | Any claim about what already works or what has been tested                                                       |

### Soft vs Hard Blockers

- **Hard blocker for design validity of interdependence claims:** OQ-001 + OQ-003 + OQ-004 + OQ-005 treated as one package.
- **Hard blocker for identity/reputation claims:** OQ-017 (provenance).
- **Soft blocker for first playable:** many society and combat questions can remain TBD longer.
- **Process hard blocker:** absence of prototype inspection (AS-001 / PROT reconciliation).

---

## D. Master GDD Readiness Verdict

**YELLOW — Suitable as the single working reference, but requires corrections before Canonical Baseline.**

### Conditions to reach GREEN (Master GDD v1.1 Canonical Baseline)

1. Reclassify the five derived items listed in Section A from LOCKED to DERIVED CONSTRAINT (or PROPOSED).
2. Add DERIVED CONSTRAINT to the official status vocabulary.
3. Explicitly label the document status as “Working Canonical Reference — Consolidation Candidate” until the above is done and human confirmation of the HUMAN-LOCKED set is recorded.
4. Record that BAL-001 (no premature numeric tuning) is a formal project design principle.
5. Leave all mechanical detail (resources, crafting, manufacturing, cities, combat, currency) as PROPOSED/TBD — do not attempt to resolve them inside the audit.
6. Schedule an Evidence Reconciliation Pass for the prototype as a separate work item.

### Explicitly rejected next actions

- Do **not** ask another LLM to “improve” or re-consolidate the GDD.
- Do **not** begin isolated detailed design of skill trees, resource tables, or combat before the Economic Interdependence Core package is addressed.
- Do **not** treat the current Recommended Design Sequence as binding; replace it with the phase model below after v1.1.

---

## Recommended Path Forward

### Immediate (this audit cycle)

1. Accept this audit.
2. Produce **Master GDD v1.1 — Canonical Baseline** that:
   - Applies the reclassifications above.
   - Adds DERIVED CONSTRAINT status.
   - Updates Document Control to “Working Canonical Reference”.
   - Does **not** expand mechanical design.

### Next Design Phase (after v1.1)

**Phase 0 — Evidence Reconciliation** (parallel or immediately after v1.1)\
Inspect any existing prototype. Produce a short evidence matrix: GDD claim ↔ prototype behaviour ↔ test status ↔ decision status.

**Phase 1 — Economic Interdependence Core (cross-system package)**\
Design together, not sequentially:

- Specialisation meaning and budget shape (what a character must be *unable* to do simultaneously)
- Resource quality / attribute model (shape only)
- Crafting differentiation and experimentation shape
- Manufacturing constraints that preserve skilled-player relevance
- NPC substitution limits
- Basic demand generation for ordinary and high-quality goods

Success criterion: a set of design invariants that can be simulated.

**Phase 2 — Provenance & Reputation Engine**\
Resource → processor → crafter → organisation → manufacturer → retailer → owner → reputation.\
Answer OQ-017 with concrete rules for multi-contributor products and brands.

**Phase 3 — Commerce & Services**\
Retail, recurring services, demand loops.

**Phase 4 — Society**\
Organisations, property, cities, governance, institutions — the machinery that makes “society” more than a marketplace.

**Phase 5 — World & Logistics**\
Geography, transport, regional markets.

**Phase 6 — Combat**\
Designed *into* the existing economy, not as a separate game.

**Phase 7 — Simulation**\
Only after structural shapes exist.

---

## Design Invariant Framework (Recommended Adoption)

Every major system should eventually carry:

| Element               | Purpose                                     |
| --------------------- | ------------------------------------------- |
| **Principle**         | What we want                                |
| **Constraint**        | What must never happen                      |
| **Invariant**         | What must remain true under testing         |
| **Failure Condition** | Observable signal that the design is broken |
| **Test**              | How we measure it                           |

Example for Interdependence (illustrative only — not yet locked):

- **Principle:** Specialists should matter.
- **Constraint:** No single character should efficiently dominate all major economic roles.
- **Invariant:** High-end production of defined critical goods requires contribution from at least N distinct capability domains.
- **Failure Condition:** Economically optimal solo production exceeds a defined threshold of specialist-network production for those goods.
- **Test:** Simulated multi-agent economy under controlled specialisation budgets.

This framework should be adopted as a documentation standard in v1.1 or immediately thereafter.

---

## Final Assessment

Grok’s consolidation is a successful process document. It correctly prioritised authority hygiene and uncertainty preservation over false completeness. That is the right outcome for a consolidation pass.

The remaining work is **adjudication and structural design**, not further synthesis.

The highest-leverage next human decision is not “which skill costs more points.” It is:

> What must a single character be structurally unable to do at the same time for TCIndustries’ economy and society to remain interdependent?

Until that question has a testable answer, claims about interdependence, anti-automation, and meaningful professions remain aspirational.

---

**Audit complete.**\
Recommended immediate action: apply the reclassifications and status vocabulary update to produce Master GDD v1.1 Canonical Baseline, then proceed to Phase 0 / Phase 1 as defined above.

**Agreed.** Your assessment is correct, and I have executed the recommended next step rather than rewriting or expanding the GDD.

### Deliverable

**Canonical GDD Audit** saved to:

---

### Audit Summary

**Verdict: YELLOW — Consolidation Candidate**

Suitable as the single working reference. Not yet ready to freeze as Fully Canonical.

#### What the consolidation got right

- Status discipline and the explicit rule against silent promotion.
- Preservation of the citizen / interdependence / player-society identity.
- Anti-goal framing with mechanical intent.
- Refusal to invent numbers.
- Correct identification of specialisation (OQ-001) as foundational.
- No promotion of the “factories cannot produce Exceptional” rule (that specific restriction is **not** present as LOCKED in the consolidated document).

#### Material corrections required before v1.1

Five items currently marked LOCKED are **derived constraints**, not independent human decisions:

| ID       | Issue                                                                        |
| -------- | ---------------------------------------------------------------------------- |
| CRFT-006 | “No Instant Mastery of Everything” — derived from interdependence principles |
| MFG-004  | Anti-Factorio drift — already covered by VIS-006 + MFG-001                   |
| ECO-003  | No pure spreadsheet optimisation — anti-goal restatement                     |
| CMBT-003 | No Forced Combat Path — covered by LOOP-002 + VIS-006                        |
| PROG-002 | Multiple viable paths — covered by LOOP-002                                  |

**Recommendation:** Introduce a **DERIVED CONSTRAINT** status and reclassify these five. The parent principles remain HUMAN-LOCKED.

#### Structural finding

The recommended linear sequence in the GDD understates a real dependency graph. Specialisation, resource attributes, crafting differentiation, and manufacturing constraints form a tightly coupled **Economic Interdependence Core**. They should be designed as one package, not sequentially.

#### Explicitly rejected next actions

- Do not ask another LLM to “improve” or re-consolidate the GDD.
- Do not begin isolated detailed design of skill trees, resource tables, or combat.
- Do not treat the current document as frozen canon.

#### Recommended path

1. **Immediate:** Produce Master GDD v1.1 Canonical Baseline — status hygiene only (reclassifications + DERIVED CONSTRAINT vocabulary + document status label). No new mechanical design.
2. **Phase 0 (parallel or next):** Evidence Reconciliation Pass against any existing prototype.
3. **Phase 1:** Economic Interdependence Core package (specialisation meaning + resource quality shape + crafting differentiation + manufacturing constraints + NPC substitution limits), expressed as testable design invariants.

The highest-leverage open question remains:

> What must a single character be structurally unable to do at the same time for TCIndustries’ economy and society to remain interdependent?

Until that has a testable answer, interdependence claims stay aspirational.

I am ready to produce the v1.1 Canonical Baseline (status corrections only) or to wait for your direction on the audit findings.

## Response:
25/08/2026, 15:11:33

Yes. **I accept this audit as the controlling assessment of the Grok-consolidated GDD.**

It is materially more rigorous than simply accepting the consolidation, and I agree with its central conclusion:

> **The GDD should not be regenerated. It should be corrected and promoted to a v1.1 Canonical Baseline through a status/authority pass only.**

### One important refinement

I would preserve the audit's distinction between:

- **HUMAN-LOCKED** - explicitly authoritative.
- **DERIVED CONSTRAINT** - binding consequence of an authoritative principle.
- **PROPOSED** - design recommendation awaiting approval.
- **TBD** - unresolved.
- **PROTOTYPE / EVIDENCE** - demonstrated implementation evidence.
- **ASSUMPTION** - working assumption.
- **DEFERRED** - deliberately postponed.
- **HISTORICAL** - reference only.

I agree that **DERIVED CONSTRAINT should not be represented as LOCKED**. That distinction is particularly important for future cross-LLM audits because otherwise different models can mistake logical consequences for explicit human decisions.

### I also agree with the sequencing

The next substantive design problem should **not** be skill trees.

It should be the coupled:

**Specialisation ↔ Resources ↔ Crafting ↔ Manufacturing ↔ NPC substitution ↔ Demand**

package.

And the proposed **Principle → Constraint → Invariant → Failure Condition → Test** framework is worth adopting as a formal TCIndustries design-documentation standard.

### Current project state

I would therefore record the project conceptually as:

**TCIndustries Master GDD**
- Current consolidated document: **v1.0**
- Status: **Consolidation Candidate / Working Reference**
- Canonical freeze: **NOT YET**
- Required immediate artifact: **v1.1 Canonical Baseline**
- v1.1 scope: **authority/status corrections only**
- First substantive design phase after v1.1: **Economic Interdependence Core**
- Prototype reconciliation: **open process requirement**

And importantly:

**No further LLM consolidation should occur.**

The next operation should be controlled correction of the existing document, followed by evidence reconciliation and then actual systems design.

If you want to proceed, the appropriate next step is to **create `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md` from the Grok v1.0 document, applying only the audit-approved authority/status changes and making no new game-design decisions.**

## Prompt:
25/08/2026, 15:16:36

# TCIndustries — Master Game Design Document

## Version 1.1 — Canonical Baseline

**Project:** Tiwakings Craftworld Industries (TCIndustries)\
**Document Title:** Master Game Design Document\
**Version:** 1.1 — Canonical Baseline\
**Baseline Date:** 2026-08-25\
**Document Purpose:** Working canonical design reference for TCIndustries. This version applies authority/status corrections identified in the Canonical Audit of the v1.0 Consolidated Master GDD. No new game-design decisions are introduced.\
**Authority Statement:** This document is the current working canonical design reference for TCIndustries. It supersedes the v1.0 Consolidated Master GDD. Status tags are authoritative. No status is silently promoted. Derived constraints are explicitly distinguished from human-locked decisions.

**Lineage:**

- Source consolidation: `TCIndustries_Master_GDD_v1.0_Consolidated.md`
- Controlling audit: `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md`
- Original first-pass sources remain as provenance only.

---

# 1. Document Control

| Field             | Value                                                                                                                          |
| ----------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| Project           | TCIndustries (Tiwakings Craftworld Industries)                                                                                 |
| Document          | Master Game Design Document                                                                                                    |
| Version           | 1.1 — Canonical Baseline                                                                                                       |
| Status            | Working Canonical Reference (status/authority pass complete)                                                                   |
| Scope             | Game design and rules only. Technical implementation belongs in a future Technical Design Document (TDD).                      |
| Change Control    | Status changes require explicit project/human approval. Derived constraints may be refined when parent principles are refined. |
| Prior Version     | 1.0 — Consolidated Master GDD                                                                                                  |
| Controlling Audit | TCIndustries\_Master\_GDD\_v1.0\_Canonical\_Audit.md                                                                           |

---

# 2. Executive Summary

TCIndustries is a persistent sandbox MMORPG centred on living in a player-driven society. It is inspired by the virtual-world philosophy and systemic depth of early *Star Wars Galaxies* — dynamic resources, interdependent professions, player-made goods, recognisable crafters, player commerce, housing, settlements, social services, and emergent economic life — under entirely original intellectual property.

It is **not** intended to reproduce Star Wars lore, characters, factions, names, or protected content.

The intended player fantasy is not “become the chosen hero.” It is:

> Become somebody known in a persistent world.

A player may become a resource prospector, farmer, medical practitioner, entertainer, master craftsperson, merchant, trader, architect, combatant, researcher, city organiser, manufacturer, transport operator, or a hybrid of several identities.

The game’s central loop is:

> Discover resources → acquire and process materials → craft goods or provide services → specialise → establish reputation → participate in trade, society, and regional economies.

TCIndustries must avoid becoming:

- an automation-first industrial game,
- a conventional class-based MMO,
- a quest-driven theme park,
- a single-player crafting game with multiplayer features,
- a pure market spreadsheet simulator,
- a combat-only MMO,
- a world where NPCs make meaningful player economic roles irrelevant,
- a system where one character can efficiently perform every economically important role.

Industry, production, combat, and economic simulation exist to create meaningful social roles and player interdependence. The ultimate measure of success is whether the systems combine to produce a world that feels **alive**.

This document deliberately identifies many mechanics as **PROPOSED**, **TBD**, or **DEFERRED** rather than inventing premature false precision. Agreement among LLM-generated source documents does not constitute human approval.

---

# 3. How to Read This Document

Every significant design element carries an explicit status. Ordinary explanatory prose does not override the status of individual mechanics.

| Status                       | Meaning                                                                                                                                                                    |
| ---------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **LOCKED**                   | Explicit project/human decision; authoritative unless formally revised. Also called HUMAN-LOCKED where distinction from derived material is needed.                        |
| **DERIVED CONSTRAINT**       | Binding consequence of one or more LOCKED principles. Not an independent human decision. May be refined when parent principles are refined. Must not be treated as LOCKED. |
| **PROPOSED**                 | Recommended design direction that has not been explicitly approved.                                                                                                        |
| **TBD**                      | Genuine unresolved design decision.                                                                                                                                        |
| **PROTOTYPE** / **EVIDENCE** | Behaviour demonstrated or explored by prototype implementation; not canonical merely because it exists.                                                                    |
| **DEFERRED**                 | Recognized future system intentionally postponed.                                                                                                                          |
| **ASSUMPTION**               | Working assumption used by one or more sources but not established as fact or decision.                                                                                    |
| **HISTORICAL**               | Historical/reference information (particularly from SWG) that is not itself a TCIndustries rule.                                                                           |

**Critical Rule:** Never silently promote PROPOSED, TBD, PROTOTYPE, DEFERRED, ASSUMPTION, or DERIVED CONSTRAINT material into LOCKED. Status changes require explicit documented human/project approval. Logical consequences of LOCKED principles are recorded as DERIVED CONSTRAINT, not LOCKED.

Rule identifiers (e.g., VIS-001, RES-002, CRFT-001) are retained or normalised for traceability. Collisions between source identifiers are resolved in the Provenance Appendix.

---

# 4. Authority and Status Rules

1. Explicit human/project rulings have highest authority (LOCKED / HUMAN-LOCKED).
2. LOCKED rules from sources that explicitly attribute them to project brief or human direction are treated as LOCKED.
3. Logical consequences of LOCKED principles are recorded as **DERIVED CONSTRAINT**, not LOCKED.
4. Consistent design decisions independently proposed across multiple documents remain PROPOSED unless higher authority exists.
5. Prototype behaviour is evidence, not design authority. GDD rules govern over prototype behaviour unless a later explicit human ruling changes that relationship.
6. Historical SWG behaviour is inspiration only. It is never automatically a TCIndustries rule.
7. Where authority cannot be established, uncertainty is preserved as TBD, PROPOSED, or ASSUMPTION.

---

# 5. Vision

## VIS-001 — Living Virtual World

**Status: LOCKED**\
**Authority:** Explicit project brief / consistent across all sources as core objective.

TCIndustries is a persistent multiplayer virtual world in which players create their own economic, social, occupational, and reputational identities.

The world should feel inhabited not only by NPCs and systems, but by player-created institutions: shops, workshops, supply networks, social venues, cities, organisations, specialist communities, commercial districts, brands, and player reputations.

The world should appear to continue existing independently of any particular player’s story. Players participate in the world rather than being the centre of it.

The game’s primary content is the interaction between players, world systems, and player-created consequences.

## VIS-002 — Citizen Rather Than Chosen Hero

**Status: LOCKED**\
**Authority:** Explicit project brief / consistent core player fantasy.

Players should feel like citizens of a living society rather than protagonists of a predetermined story.

A player’s prestige should be able to derive from non-combat achievements, including:

- discovering a rare resource,
- producing respected goods,
- operating a reliable shop,
- becoming a famous doctor or entertainer,
- coordinating a city,
- supplying a war effort,
- operating a trusted trade network,
- designing desirable buildings,
- becoming an expert in a narrow specialisation.

A player should be able to say: “This is who I am in the world,” rather than “This is the quest line I completed.”

## VIS-003 — Player-Driven Economy

**Status: LOCKED**\
**Authority:** Explicit project direction across sources.

Players should be major participants in the production and distribution of economic goods and services. The economy must be materially connected to the world. Resources, labour, transport, crafting, consumption, degradation, services, and regional differences should matter.

## VIS-004 — Interdependence

**Status: LOCKED**\
**Authority:** Explicit project pillar.

Important professions should depend upon one another. A complete economy should contain useful roles that cannot all be replaced by a single generalist player.

## VIS-005 — Emergence

**Status: LOCKED**\
**Authority:** Explicit project direction.

The game should create conditions from which interesting events emerge rather than attempting to script every important player experience.

Desired player thought: “This world changed because of what players did, and other players know who matters here.”

## VIS-006 — Anti-Goals

**Status: LOCKED**\
**Authority:** Explicit project direction / consistent anti-drift language.

TCIndustries must not drift into:

- a linear quest-driven MMORPG,
- a rigid class-based MMORPG,
- a single-player crafting game with multiplayer chat,
- an infinite personal-factory automation game,
- a pure market spreadsheet simulator,
- a combat-only MMORPG,
- an economy where all meaningful production is performed by NPCs,
- a world where all players can independently perform every role efficiently.

---

# 6. Design Pillars

## PIL-001 — Discovery and the Gold Rush

**Status: LOCKED**\
**Authority:** Explicit project pillar.

Resources are temporary, geographically distributed, variable in quality, and economically meaningful.

Exceptional resource discoveries should create player-driven events: prospecting activity, mining camps, trade opportunities, speculation, guild coordination, regional traffic, competitive supply chains, and demand for extractors, transport, processing, and manufacturing.

A resource discovery should be an **event**, not merely an inventory update.

**Historical SWG Inspiration (HISTORICAL):** Original SWG used temporary resource spawns with variable statistics, regional availability, and resource surveying.

**TCIndustries Adaptation (PROPOSED — RES-001):** Retain the principle of dynamic, temporary, statistically variable resource occurrences, while improving information clarity, anti-monopoly controls, and long-term economic observability.

## PIL-002 — Identity and Reputation

**Status: LOCKED**\
**Authority:** Explicit project pillar.

Player identity must be meaningful, visible, and socially consequential.

Products, services, businesses, and buildings should preserve meaningful association with their creator, provider, organisation, or brand.

Reputation should emerge from player behaviour and economic history rather than merely from an abstract reputation score. Ideal outcome: players seek out products made by a particular player or organisation.

## PIL-003 — Interdependence

**Status: LOCKED**\
**Authority:** Explicit project pillar.

No major profession category should be fully self-sufficient at high effectiveness.

Interdependence should arise through: distinct skills, specialised resources, intermediate components, service needs, regional logistics, knowledge gaps, reputation, production capacity limits, and social organisation.

Interdependence must create opportunity rather than coercive inconvenience.

## PIL-004 — Player-Created Society

**Status: LOCKED**\
**Authority:** Explicit project pillar.

Players must be able to form organisations, businesses, settlements, social hubs, trade networks, and reputational institutions with meaningful world presence.

---

# 7. Player Experience and Core Gameplay Loops

## LOOP-001 — Interconnected Loops

**Status: PROPOSED**

The game supports several interconnected loops rather than a single mandatory progression path.

| Loop              | Core Activities                                   | Social Outcome                                |
| ----------------- | ------------------------------------------------- | --------------------------------------------- |
| Discovery         | Survey, explore, identify resources               | Information trade, rushes, regional knowledge |
| Extraction        | Harvest, operate extraction equipment, gather     | Resource supply and logistics                 |
| Processing        | Refine, convert, prepare materials                | Intermediate-goods markets                    |
| Crafting          | Build items, experiment, specialise               | Brand identity and product differentiation    |
| Retail            | Stock vendors, price goods, market services       | Local commerce and commercial districts       |
| Service           | Heal, entertain, repair, transport, construct     | Recurring player interaction                  |
| Combat            | Fight threats, protect territory, acquire rewards | Demand for equipment, medicine, consumables   |
| Social Governance | Form associations, cities, events, policies       | Community identity and institutions           |

## LOOP-002 — Non-Combat Viability

**Status: LOCKED**\
**Authority:** Explicit project direction.

A player must be able to pursue a meaningful long-term career without engaging in combat as their primary activity.

Combat may offer demand, risk, materials, access, and opportunities, but must not be the sole source of progression, wealth, or social status.

## LOOP-003 — Progression Philosophy

**Status: LOCKED**\
**Authority:** Explicit project philosophy.

Progression should expand choices, capability, specialisation, reputation, and access to complex roles. It should not force all players through one linear power ladder.

---

# 8. World Philosophy and Structure

## WRLD-001 — Persistent Shared World

**Status: LOCKED**

TCIndustries is a persistent shared world. Player actions should have durable consequences where practical, including: built structures, business inventories, market prices, resource depletion, city development, organisational ownership, local reputation, and product provenance.

## WRLD-002 — Geographic Structure

**Status: PROPOSED**

The world should consist of multiple regions with distinct combinations of terrain, climate, resource pools, settlement potential, dangers, transportation constraints, local demand, aesthetic identity, and environmental conditions.

Regions should not merely be level bands. They should create economic and social differentiation.

## WRLD-003 — Settlements, Cities, and Wilderness

**Status: PROPOSED**

The world should contain:

1. **NPC settlements** — Stable service, trade, tutorial, and baseline social locations.
2. **Player settlements/cities** — Player-founded areas with governance, zoning, commerce, civic structures, and community identity.
3. **Wilderness regions** — Exploration, harvesting, combat, resource discovery, and construction opportunities.
4. **Economic corridors** — Roads, transport routes, ports, stations, or equivalent systems connecting markets and resource regions.

Exact city-formation rules are **TBD**.

## WRLD-004 — Time and Environmental Simulation

**Status: PROPOSED**

The world should use a persistent world clock and support at least basic environmental variation (day/night, weather, seasons, ecological changes, temporary hazards, travel conditions, resource extraction modifiers).

Weather and seasons must only be implemented where they create meaningful decisions rather than visual noise.

## WRLD-005 — Transportation

**Status: PROPOSED**

Transportation should create meaningful geography without turning routine trade into excessive travel friction.

Support a spectrum: local movement, personal vehicles or mounts, public transit, organisation-operated transport, freight logistics, and constrained fast travel.

Fast travel must not eliminate the economic relevance of location, local markets, or transport services.

Exact transportation mechanics are **TBD**.

---

# 9. Character, Identity, and Ownership

## PLR-001 — Persistent Character Identity

**Status: LOCKED**

Each player character has a persistent identity that can accumulate: name, appearance, ownership, skill history, profession specialisations, reputation, business affiliation, organisation membership, product/service provenance, and social relationships.

## PLR-002 — Character Creation

**Status: PROPOSED**

Character creation should prioritise identity and aesthetic expression rather than granting permanent mechanical superiority.

Initial mechanical differences, if any, must not trap players into inferior long-term choices.

Species, ancestry, culture, and starting location design are **DEFERRED** pending setting development.

## PLR-003 — Ownership

**Status: LOCKED**

Players and organisations must be able to own meaningful persistent assets, including: items, resources, structures, vendors, business inventory, manufacturing permissions, organisation property, brands, and contracts where implemented.

Ownership systems must include safeguards against fraud, abandoned assets, and exploitative transfer mechanics.

---

# 10. Skills, Professions, and Progression

## PROF-001 — Flexible Skill-Based Profession Architecture

**Status: LOCKED**\
**Authority:** Explicit project requirement.

TCIndustries uses a flexible skill-based profession architecture rather than rigid immutable classes.

A profession is an emergent identity resulting from acquired skills, specialisation choices, equipment, social reputation, and player behaviour.

## PROF-002 — Skill Domains

**Status: PROPOSED**

Initial skill domains should include:

- resource surveying and prospecting,
- extraction and harvesting,
- processing and refining,
- crafting disciplines,
- manufacturing,
- architecture and construction,
- commerce and retail,
- medical treatment,
- entertainment and social services,
- combat disciplines,
- survival and exploration,
- research and analysis,
- transport and logistics.

Specific trees, ranks, prerequisites, and skill-point costs are **TBD**.

## PROF-003 — Specialisation

**Status: PROPOSED**

Players may develop broad competency across domains but should require meaningful investment to reach high-tier specialisation.

A master weaponsmith should not automatically also be a top medical practitioner, architect, merchant, and combatant. Specialisation creates markets for experts.

## PROF-004 — Respecialisation

**Status: LOCKED**

Players must be able to respecialise. Permanent character traps are contrary to sandbox identity development.

**PROPOSED implementation principle:** Respecialisation should involve meaningful opportunity cost, cooldown, retraining time, loss of active specialisation capacity, or economic cost — but should not erase player identity, item provenance, business ownership, or reputation history.

Exact respecialisation rules are **TBD**.

## PROF-005 — Progression Caps / Specialisation Budget

**Status: PROPOSED**

The game should have a capability cap or specialisation budget that prevents one character from mastering every economically important role at once.

The cap must preserve experimentation, casual participation, hybrid roles, respecialisation, and player autonomy.

The exact cap model is **TBD**.

---

# 11. Resource System

## RES-002 — Resource Design Goal

**Status: LOCKED**

The resource system exists to create discovery, scarcity, regional trade, product differentiation, temporary opportunity, and social/economic events.

Resources must not function solely as generic crafting currencies.

## RES-003 — Resource Taxonomy

**Status: PROPOSED**

Resources should be defined through extensible data categories: resource family, resource class, subtype, origin, material state, applicable uses, measurable attributes, rarity characteristics, spawn rules, extraction requirements, and processing options.

Example high-level families (examples only, not locked content): minerals, metals, stone, chemical compounds, organics, fibers, woods, fuels, gases, biological materials, water and liquids, rare technical materials.

## RES-004 — Resource Instances

**Status: PROPOSED**

A resource spawn is an individual temporary resource instance with unique identifier, name, taxonomy classification, spawn time, expected end time or lifecycle state, geographic distribution, density, attribute values, rarity, extraction difficulty, discoverability, and remaining accessible quantity where applicable.

Each instance should be distinguishable from generic material categories.

Example form: “Kavren Ferritic Alloy, discovered in the North Ember Plains, cycle 184.”

## RES-005 — Resource Attributes

**Status: PROPOSED**

Resources should possess measurable attributes relevant to recipes and product outcomes (e.g., purity, hardness, conductivity, toughness, density, reactivity, nutritional value, thermal properties, etc.).

Exact attribute sets, ranges, and mapping to product outcomes are **TBD**.

## RES-006 — Lifecycle and Scarcity

**Status: PROPOSED**

Resources should have finite or cycling availability. Discovery, depletion, and rediscovery should create economic events. Exact spawn rates, durations, and depletion mathematics are **TBD** (do not invent numbers).

## RES-007 — Surveying and Information

**Status: PROPOSED**

Surveying/prospecting skills should create meaningful information asymmetry and opportunity. Information itself can have economic value.

---

# 12. Crafting System

## CRFT-001 — Differentiated Products

**Status: LOCKED**

Crafting must create differentiated products. Products of the same schematic should not be interchangeable commodities when resource quality, experimentation, and crafter skill differ meaningfully.

## CRFT-002 — Schematics and Recipes

**Status: PROPOSED**

Crafting uses schematics/recipes that specify required resources (with attribute requirements or preferences), skill requirements, tools/stations, and possible outcomes including experimentation results.

Exact schematic structure is **TBD**.

## CRFT-003 — Experimentation

**Status: PROPOSED**

Experimentation (or equivalent) should allow skilled crafters to influence product attributes within the constraints of the resources used and schematic limits. This supports product differentiation and crafter reputation.

Exact experimentation model is **TBD**.

## CRFT-004 — Provenance

**Status: PROPOSED**

Crafted items should retain meaningful provenance linking them to the crafter, organisation, and/or resource sources where practical. This supports the Identity and Reputation pillar.

## CRFT-005 — Skill and Tool Requirements

**Status: PROPOSED**

Higher-tier or specialised crafting should require appropriate skills, tools, and/or facilities. Exact requirements are **TBD**.

## CRFT-006 — No Instant Mastery of Everything

**Status: DERIVED CONSTRAINT**\
**Derives from:** PIL-003 (Interdependence), PROF-001 / specialisation principles

Crafting depth must reinforce specialisation rather than allow one character to efficiently dominate all high-value production.

---

# 13. Manufacturing and Automation

## MFG-001 — Controlled Automation

**Status: LOCKED**\
**Authority:** Explicit anti-automation-drift direction.

Automation and manufacturing systems must not eliminate the economic or social relevance of other players.

Factories or automated production, if present, must operate under meaningful constraints of inputs, maintenance, skill, ownership, capacity, logistics, or risk so that pure automation cannot replace player interdependence at scale.

## MFG-002 — Manufacturing Facilities

**Status: PROPOSED**

Manufacturing facilities (factories, workshops, processing plants) may exist as player-owned or organisation-owned structures that convert inputs into outputs over time under defined rules.

Exact facility rules, throughput, and maintenance are **TBD**.

## MFG-003 — Automation Limits

**Status: PROPOSED**

Any automation must be designed so that:

- input quality and scarcity still matter,
- specialised human oversight or skilled setup retains value,
- logistics and regional factors remain relevant,
- one player cannot efficiently operate a complete closed industrial empire without other players.

## MFG-004 — Anti-Factorio Drift

**Status: DERIVED CONSTRAINT**\
**Derives from:** VIS-006 (Anti-Goals), MFG-001 (Controlled Automation)

The game must not become an infinite personal-factory automation game where the optimal strategy is to build self-sufficient automated production that renders other players optional.

---

# 14. Items

## ITM-001 — Item Identity and Provenance

**Status: PROPOSED**

Items should support meaningful identity: quality attributes, provenance, condition/decay where applicable, and association with creators or owners.

## ITM-002 — Durability and Maintenance

**Status: PROPOSED / TBD**

Whether and how items decay, require repair, or have limited lifespan is **TBD**. Any such system must create demand for repair services and materials without becoming pure friction.

## ITM-003 — Stacking and Uniqueness

**Status: PROPOSED**

Generic materials may stack. High-quality or unique crafted goods should retain individual identity where this supports reputation and differentiation.

---

# 15. Economy

## ECO-001 — Player-Driven Economy

**Status: LOCKED**

The economy is primarily player-driven. Players produce, consume, trade, and specialise. NPC vendors, if present, must not make player production irrelevant.

## ECO-002 — Regional Economies

**Status: PROPOSED**

Regional differences in resources, demand, transport costs, and local conditions should create trade opportunities and local market character.

## ECO-003 — No Pure Spreadsheet Optimisation

**Status: DERIVED CONSTRAINT**\
**Derives from:** VIS-006 (Anti-Goals), VIS-003 (Player-Driven Economy)

The economy must remain grounded in the world. Pure spreadsheet optimisation without meaningful world interaction, location, risk, or social factors is contrary to design intent.

## ECO-004 — Currency and Pricing

**Status: TBD**

Exact currency systems, money sinks, and pricing mechanics are unresolved. Multiple currencies or regional media of exchange remain possible design directions.

---

# 16. Retail and Commerce

## RET-001 — Player Retail

**Status: LOCKED**

Players must be able to operate shops, vendors, and commercial spaces that sell goods and services to other players.

## RET-002 — Vendor Systems

**Status: PROPOSED**

Vendor systems should support stock management, pricing, and association with player or organisation identity. Exact vendor mechanics are **TBD**.

## RET-003 — Commercial Districts and Branding

**Status: PROPOSED**

Player commerce should support the formation of commercial districts, recognised brands, and trusted suppliers.

---

# 17. Services

## SERV-001 — Meaningful Non-Combat Services

**Status: LOCKED**

Services such as medical treatment, entertainment, repair, transport, construction, and other player-provided services must be viable long-term careers.

## SERV-002 — Service Interdependence

**Status: PROPOSED**

Services should create recurring player interaction and reinforce interdependence (e.g., combatants needing medical care, social venues needing entertainers, structures needing architects/builders).

Exact service lists and mechanics are **TBD**.

---

# 18. Combat

## CMBT-001 — Combat as One Lifestyle Among Many

**Status: LOCKED**

Combat is a valid and important lifestyle, but not the default or mandatory path. It generates demand for equipment, medicine, consumables, and related services.

## CMBT-002 — Combat Scope

**Status: PROPOSED / TBD**

Exact combat systems (PvE, PvP, territorial, etc.), risk models, and progression integration remain largely **TBD**. Any combat system must respect non-combat viability and interdependence pillars.

## CMBT-003 — No Forced Combat Path

**Status: DERIVED CONSTRAINT**\
**Derives from:** LOOP-002 (Non-Combat Viability), VIS-006 (Anti-Goals)

Players must not be forced into combat as the sole meaningful progression or status path.

---

# 19. Buildings, Housing, and Cities

## BLD-001 — Player-Owned Structures

**Status: LOCKED**

Players and organisations must be able to own and place meaningful structures (housing, workshops, shops, civic buildings, etc.) with persistent presence.

## BLD-002 — Housing and Function

**Status: PROPOSED**

Structures should support functional roles (storage, crafting stations, vendors, social space, manufacturing) as well as aesthetic and identity expression.

## BLD-003 — Cities and Governance

**Status: PROPOSED / TBD**

Player cities should support governance, zoning, taxation or maintenance models, civic structures, and community identity. Exact city formation, governance, and rules are **TBD**.

## BLD-004 — Placement and Limits

**Status: TBD**

Structure placement rules, lot systems, density limits, and abandonment handling are unresolved.

---

# 20. Organisations and Social Systems

## SOC-001 — Player Organisations

**Status: LOCKED**

Players must be able to form organisations with meaningful collective identity, ownership, and presence.

## SOC-002 — Organisation Capabilities

**Status: PROPOSED**

Organisations should be able to own property, operate businesses, manage membership, and participate in cities or larger social structures. Exact organisation systems are **TBD**.

## SOC-003 — Social Institutions

**Status: PROPOSED**

The design should support emergent social institutions (guilds, trade associations, city councils, entertainment venues, etc.) without scripting every institution.

---

# 21. Progression

## PROG-001 — Option Expansion, Not Linear Power

**Status: LOCKED** (restatement of LOOP-003)

Progression expands choices, specialisation depth, reputation, and access rather than forcing a single linear power ladder.

## PROG-002 — Multiple Viable Paths

**Status: DERIVED CONSTRAINT**\
**Derives from:** LOOP-002 (Non-Combat Viability)

Multiple long-term identities and careers must remain viable.

## PROG-003 — Skill Acquisition Model

**Status: TBD**

Exact skill acquisition (points, experience, training, discovery, etc.) remains unresolved. Any model must support specialisation budgets and respecialisation principles.

---

# 22. Emergent Gameplay

## EMRG-001 — Conditions for Emergence

**Status: LOCKED**

The design prioritises creating conditions from which interesting player-driven events emerge over scripting every important experience.

Resource discoveries, market shifts, city politics, reputation events, and supply crises should arise from system interactions.

---

# 23. Content Expansion Framework

## EXP-001 — Data-Driven Extensibility

**Status: LOCKED** / **PROPOSED** (principle locked; implementation PROPOSED)

Systems should be designed for data-driven content expansion (new resources, schematics, structures, skills) without requiring core system redesign for every addition.

Exact content pipeline is **TBD**.

---

# 24. Anti-Goals / Drift Prevention

**Status: LOCKED** (restatement and reinforcement of VIS-006)

The project actively guards against:

1. Theme-park / quest-driven drift.
2. Rigid class-based progression.
3. Combat-as-only-path.
4. Automation-first industrial simulation (Factorio drift).
5. Pure market spreadsheet simulation (EVE-spreadsheet drift without world grounding).
6. NPC economic dominance that makes players optional.
7. Single-character mastery of all economically important roles.
8. Design that treats non-combat careers as secondary or incomplete.

All major system proposals must be checked against these anti-goals.

---

# 25. Balance Philosophy

## BAL-001 — Avoid Premature Numeric Tuning

**Status: LOCKED** (consistent philosophy across sources)

Do not invent skill caps, spawn rates, prices, decay curves, combat coefficients, or other numeric balance values until the structural design is stable and the values can be tested.

False precision is worse than explicit unresolved questions.

## BAL-002 — Balance for Interdependence and Viability

**Status: PROPOSED**

Balance should prioritise the viability of multiple careers, the value of specialisation, and the health of interdependence rather than pure combat DPS parity or identical time-to-effectiveness across all roles.

---

# 26. Exploit, Fraud, and Economic Stability

## SAFE-001 — Ownership and Transfer Safeguards

**Status: LOCKED**

Ownership systems must include safeguards against fraud, abandoned assets, and exploitative transfer mechanics.

## SAFE-002 — Multi-Accounting and Economic Exploits

**Status: PROPOSED / TBD**

Policies and systems regarding multi-accounting, market manipulation, and economic exploits require explicit design. Exact rules are **TBD**.

## SAFE-003 — Economic Observability

**Status: PROPOSED**

The design should support observability of economic health so that designers can detect and respond to systemic problems without relying solely on player reports.

---

# 27. Testing and Validation Requirements

## TEST-001 — Design Invariants

**Status: PROPOSED**

Major systems should be accompanied by testable design invariants (e.g., “a specialised weaponsmith must produce higher-quality weapons than a generalist under equivalent resource quality,” “non-combat careers must remain economically viable over long play sessions”).

## TEST-002 — Simulation Preference

**Status: PROPOSED**

Where practical, economic and resource systems should be simulatable before full implementation to validate scarcity, interdependence, and anti-monopoly properties.

Exact testing protocols are **TBD**.

---

# 28. Prototype Relationship

## PROT-001 — Prototype Authority

**Status: LOCKED** (principle)

GDD rules are authoritative over prototype behaviour unless a later explicit human ruling changes that relationship.

Prototype behaviour is evidence of experimentation, not automatically approved design.

## PROT-002 — Existing Prototype

**Status: PROTOTYPE / ASSUMPTION**

Sources reference a Seed-2.1 TypeScript/Vite/React prototype. No prototype source was available for direct inspection during this consolidation. Any specific prototype behaviours remain **PROTOTYPE** evidence only and must be reconciled in a future pass.

---

# 29. Historical SWG Relationship

## SWG-001 — Inspiration, Not Reproduction

**Status: LOCKED**

TCIndustries is inspired by the virtual-world philosophy and systemic lessons of early Star Wars Galaxies. It is not a reproduction of Star Wars intellectual property.

## SWG-002 — Strict Separation

**Status: LOCKED**

Historical SWG behaviour is recorded as **HISTORICAL**. TCIndustries adaptations are explicitly marked. No Star Wars-specific lore, characters, factions, names, or protected content is introduced as TCIndustries content.

---

# 30. Systems Requiring Deeper Design

The following systems require dedicated future design passes (not exhaustive):

- Detailed skill trees, costs, and specialisation budgets (**TBD**)
- Exact resource spawn, attribute, and lifecycle mathematics (**TBD**)
- Crafting experimentation model and schematic structure (**TBD**)
- Manufacturing facility rules and automation constraints (**TBD**)
- Transportation and logistics (**TBD**)
- City formation, governance, and maintenance (**TBD**)
- Combat systems (PvE/PvP scope, risk, progression) (**TBD**)
- Currency, money sinks, and economic stabilisers (**TBD**)
- Organisation governance and rights (**TBD**)
- Item durability, repair, and decay (**TBD**)
- Species/ancestry/setting content (**DEFERRED**)
- Full service profession definitions (**TBD**)

---

# 31. Open Question Register

| ID     | Question                                                                      | Affected Systems                     | Why It Matters                                 | Current Status |
| ------ | ----------------------------------------------------------------------------- | ------------------------------------ | ---------------------------------------------- | -------------- |
| OQ-001 | Exact specialisation budget / skill-cap model?                                | Skills, Professions, Interdependence | Prevents single-character mastery of all roles | TBD            |
| OQ-002 | Resource spawn rates, durations, and depletion mathematics?                   | Resources, Economy, Discovery        | Core scarcity and event generation             | TBD            |
| OQ-003 | Exact attribute sets for resources and mapping to product outcomes?           | Resources, Crafting                  | Product differentiation and experimentation    | TBD            |
| OQ-004 | Crafting experimentation model details?                                       | Crafting, Identity                   | Crafter skill expression and reputation        | TBD            |
| OQ-005 | Manufacturing/factory throughput, maintenance, and ownership rules?           | Manufacturing, Automation            | Anti-automation-drift enforcement              | TBD            |
| OQ-006 | Transportation model and fast-travel constraints?                             | World, Economy, Logistics            | Geographic meaning vs. friction                | TBD            |
| OQ-007 | City formation, governance, taxation/maintenance models?                      | Cities, Society, Buildings           | Player-created society pillar                  | TBD            |
| OQ-008 | Combat scope (PvE/PvP), risk model, and progression integration?              | Combat, Economy, Services            | Demand generation without forced path          | TBD            |
| OQ-009 | Currency system(s) and money sinks?                                           | Economy                              | Economic stability                             | TBD            |
| OQ-010 | Item durability, decay, and repair model?                                     | Items, Services, Economy             | Service demand and item identity               | TBD            |
| OQ-011 | Respecialisation exact costs and rules?                                       | Progression, Identity                | Balance of flexibility vs. commitment          | TBD            |
| OQ-012 | Multi-accounting and economic exploit policies?                               | Economy, Safety                      | Fairness and stability                         | TBD            |
| OQ-013 | Organisation rights, hierarchy, and property model?                           | Social Systems                       | Collective ownership                           | TBD            |
| OQ-014 | Species/ancestry/starting location design?                                    | Character                            | Setting and identity expression                | DEFERRED       |
| OQ-015 | Which systems are required for first playable prototype vs. persistent alpha? | Production planning                  | Scope control                                  | TBD            |
| OQ-016 | Exact vendor and retail mechanics?                                            | Retail, Economy                      | Player commerce viability                      | TBD            |
| OQ-017 | Provenance depth and visibility rules?                                        | Identity, Crafting, Items            | Reputation pillar support                      | TBD            |

---

# 32. Assumption Register

| ID     | Assumption                                                                     | Source                    | Status                | Consequence if Wrong                                           |
| ------ | ------------------------------------------------------------------------------ | ------------------------- | --------------------- | -------------------------------------------------------------- |
| AS-001 | A functional prototype (Seed-2.1 or similar) exists or will exist for testing  | Multiple sources          | ASSUMPTION            | Testing approach and reconciliation process must be redesigned |
| AS-002 | Data-driven content expansion is preferred over hard-coded systems             | Consistent across sources | ASSUMPTION / PROPOSED | Content pipeline and technical architecture change             |
| AS-003 | Regional economic differentiation is desirable and feasible                    | Consistent proposals      | ASSUMPTION            | World design and transport systems simplify or change          |
| AS-004 | Meaningful provenance tracking is technically and design-feasible at scale     | Identity pillar proposals | ASSUMPTION            | Reputation systems may need alternative approaches             |
| AS-005 | Non-combat careers can generate sufficient player demand without combat gating | Core vision               | ASSUMPTION            | Demand generation systems may need redesign                    |

---

# 33. Status / Decision Register (Selected Key Rules)

| ID                                               | Rule Summary                                     | Status              | Authority / Evidence Class                      | Dependencies                 | Notes                |
| ------------------------------------------------ | ------------------------------------------------ | ------------------- | ----------------------------------------------- | ---------------------------- | -------------------- |
| VIS-001                                          | Persistent player-driven virtual society         | LOCKED              | Explicit project objective                      | All systems                  | Core                 |
| VIS-002                                          | Citizen identity over hero narrative             | LOCKED              | Explicit player fantasy                         | Progression, social, economy | Core                 |
| VIS-003                                          | Player-driven economy                            | LOCKED              | Explicit project direction                      | Economy, production          | Core                 |
| VIS-004                                          | Interdependence of professions                   | LOCKED              | Explicit pillar                                 | Skills, economy, services    | Core                 |
| VIS-005                                          | Emergence over scripting                         | LOCKED              | Explicit project direction                      | All systems                  | Core                 |
| VIS-006                                          | Anti-goals (theme-park, class, automation, etc.) | LOCKED              | Explicit anti-drift                             | All systems                  | Core                 |
| PIL-001                                          | Discovery / Gold Rush                            | LOCKED              | Explicit pillar                                 | Resources, economy           | Core                 |
| PIL-002                                          | Identity and Reputation                          | LOCKED              | Explicit pillar                                 | Crafting, retail, services   | Core                 |
| PIL-003                                          | Interdependence                                  | LOCKED              | Explicit pillar                                 | Skills, economy              | Core                 |
| PIL-004                                          | Player-Created Society                           | LOCKED              | Explicit pillar                                 | Buildings, cities, orgs      | Core                 |
| LOOP-002                                         | Non-combat long-term viability                   | LOCKED              | Explicit direction                              | Professions, economy         | Core                 |
| LOOP-003                                         | Progression expands options                      | LOCKED              | Explicit philosophy                             | Skills                       | Core                 |
| WRLD-001                                         | Persistent shared world                          | LOCKED              | Required by vision                              | World, ownership             | Core                 |
| PLR-001                                          | Persistent character identity                    | LOCKED              | Required by pillars                             | Player systems               | Core                 |
| PLR-003                                          | Persistent ownership                             | LOCKED              | Required by society/economy                     | Buildings, items             | Core                 |
| PROF-001                                         | Flexible skill-based professions                 | LOCKED              | Explicit requirement                            | Progression                  | Core                 |
| PROF-004                                         | Respecialisation allowed                         | LOCKED              | Explicit philosophy                             | Progression                  | Core                 |
| RES-002                                          | Resources create discovery/scarcity              | LOCKED              | Pillar implementation                           | Economy, crafting            | Core                 |
| CRFT-001                                         | Differentiated products                          | LOCKED              | Identity/reputation requirement                 | Resources, skills            | Core                 |
| MFG-001                                          | Controlled automation                            | LOCKED              | Anti-drift                                      | Manufacturing                | Core                 |
| MFG-004                                          | Anti-Factorio drift                              | DERIVED CONSTRAINT  | Derives from VIS-006, MFG-001                   | Manufacturing                | Reclassified in v1.1 |
| ECO-001                                          | Player-driven economy                            | LOCKED              | Explicit                                        | Production, retail           | Core                 |
| ECO-003                                          | No pure spreadsheet optimisation                 | DERIVED CONSTRAINT  | Derives from VIS-006, VIS-003                   | Economy                      | Reclassified in v1.1 |
| RET-001                                          | Player retail                                    | LOCKED              | Explicit                                        | Economy, identity            | Core                 |
| SERV-001                                         | Meaningful non-combat services                   | LOCKED              | Explicit                                        | Interdependence              | Core                 |
| CMBT-001                                         | Combat as one lifestyle among many               | LOCKED              | Explicit                                        | Economy, services            | Core                 |
| CMBT-003                                         | No Forced Combat Path                            | DERIVED CONSTRAINT  | Derives from LOOP-002, VIS-006                  | Combat                       | Reclassified in v1.1 |
| BLD-001                                          | Player-owned structures                          | LOCKED              | Explicit                                        | Society, economy             | Core                 |
| SOC-001                                          | Player organisations                             | LOCKED              | Explicit                                        | Society                      | Core                 |
| PROG-002                                         | Multiple viable paths                            | DERIVED CONSTRAINT  | Derives from LOOP-002                           | Progression                  | Reclassified in v1.1 |
| CRFT-006                                         | No Instant Mastery of Everything                 | DERIVED CONSTRAINT  | Derives from PIL-003, specialisation principles | Crafting                     | Reclassified in v1.1 |
| EMRG-001                                         | Conditions for emergence                         | LOCKED              | Explicit                                        | All systems                  | Core                 |
| SAFE-001                                         | Ownership/transfer safeguards                    | LOCKED              | Explicit                                        | Ownership                    | Core                 |
| SWG-001/002                                      | Inspiration only; no SW IP                       | LOCKED              | Explicit                                        | Setting                      | Core                 |
| PROT-001                                         | GDD over prototype                               | LOCKED              | Explicit principle                              | All                          | Core                 |
| Most detailed taxonomies, formulas, trees, rates | PROPOSED or TBD                                  | LLM recommendations | Various                                         | Not promoted                 |                      |

---

# 34. Recommended Design Sequence

This sequence reflects the controlling Canonical Audit. It replaces the linear list from v1.0.

**Completed in this version (v1.1):**

- Authority/status hygiene pass (HUMAN-LOCKED vs DERIVED CONSTRAINT distinction).

**Next required work (process, not design expansion):**
0\. **Evidence Reconciliation Pass** — Inspect any existing prototype; produce GDD ↔ prototype ↔ decision evidence matrix. (Open process requirement.)

**First substantive design phase after evidence reconciliation:**

1. **Economic Interdependence Core (cross-system package)** — Design together, not sequentially:
   - Specialisation meaning and budget shape (what a character must be unable to do simultaneously)
   - Resource quality / attribute model (shape only; no premature numbers)
   - Crafting differentiation and experimentation shape
   - Manufacturing constraints that preserve skilled-player relevance
   - NPC substitution limits
   - Basic demand generation for ordinary and high-quality goods
   - Express as Principle → Constraint → Invariant → Failure Condition → Test

**Subsequent phases (order after Core is stable):**
2\. Provenance & Reputation Engine
3\. Commerce & Services
4\. Society (organisations, cities, governance, institutions)
5\. World & Logistics
6\. Combat (designed into the existing economy)
7\. Simulation and numeric tuning only after structural shapes exist

No further LLM consolidation of the Master GDD should occur. The next operations are controlled correction (this document), evidence reconciliation, and then actual systems design.

---

# 35. Glossary

| Term                  | Definition                                                                                                                   |
| --------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| Citizen               | Player fantasy of being an ordinary but consequential participant in a living society rather than a chosen hero.             |
| Gold Rush             | Player-driven economic and social event triggered by exceptional temporary resource discoveries.                             |
| Interdependence       | Design condition in which high-effectiveness roles require other specialised players.                                        |
| Provenance            | Association of an item, product, or service with its creator, organisation, or resource origins.                             |
| Profession            | Emergent identity arising from skills, specialisation, equipment, reputation, and behaviour — not a rigid class.             |
| Specialisation Budget | Constraint (skill points, capacity, or equivalent) preventing one character from mastering all economically important roles. |
| Status Tag            | Explicit marker (LOCKED, DERIVED CONSTRAINT, PROPOSED, TBD, etc.) controlling the authority of a design statement.           |
| DERIVED CONSTRAINT    | Binding consequence of LOCKED principle(s); not an independent human decision.                                               |

---

# 36. Provenance and Change Notes

## v1.1 Canonical Baseline — Change Scope

This version applies **only** the authority/status corrections approved in the Canonical Audit of v1.0. No new game-design decisions, mechanics, numbers, or system expansions were introduced.

### Status reclassifications (LOCKED → DERIVED CONSTRAINT)

| ID       | Rule                             | Parent Principle(s)                |
| -------- | -------------------------------- | ---------------------------------- |
| CRFT-006 | No Instant Mastery of Everything | PIL-003, specialisation principles |
| MFG-004  | Anti-Factorio Drift              | VIS-006, MFG-001                   |
| ECO-003  | No Pure Spreadsheet Optimisation | VIS-006, VIS-003                   |
| CMBT-003 | No Forced Combat Path            | LOOP-002, VIS-006                  |
| PROG-002 | Multiple Viable Paths            | LOOP-002                           |

### Vocabulary addition

- **DERIVED CONSTRAINT** added to the official status vocabulary.
- Critical Rule updated to forbid silent promotion of DERIVED CONSTRAINT into LOCKED.

### Document status

- v1.0: Consolidation Candidate / Working Reference
- v1.1: Working Canonical Reference (status/authority pass complete)
- Fully frozen project canon: still requires human confirmation of the HUMAN-LOCKED set and prototype evidence reconciliation.

## Inherited from v1.0 Consolidation

1. **LOCKED core vision and pillars** retained from consistent explicit project-level language across source documents.
2. **Detailed taxonomies, skill lists, attribute examples, and mechanical shapes** remain **PROPOSED**. LLM agreement does not constitute human approval.
3. **Numeric values** remain universally avoided / TBD.
4. **Prototype references** remain PROTOTYPE/ASSUMPTION. No prototype was inspected during consolidation or this status pass.
5. **Historical SWG material** remains strictly separated and marked HISTORICAL.
6. Rule IDs normalised; no material ID collisions required renumbering beyond synthesis.

## Remaining Human Decisions

The following still require explicit human/project approval before they can be treated as LOCKED:

1. Final specialisation budget / skill-cap model (structural meaning of “specialist”).
2. Resource lifecycle mathematics and attribute set.
3. Crafting experimentation model.
4. Manufacturing and automation constraint details.
5. City formation and governance rules.
6. Combat scope and risk model.
7. Currency and economic stabiliser design.
8. Exact respecialisation costs and rules.
9. Multi-accounting and exploit policies.
10. Confirmation that the current HUMAN-LOCKED vision/pillar statements remain accurate after human review.
11. Reconciliation of any existing prototype against this GDD (Evidence Reconciliation Pass).

## Design Documentation Standard (Adopted)

Major systems should eventually document:

- **Principle** — What we want
- **Constraint** — What must never happen
- **Invariant** — What must remain true under testing
- **Failure Condition** — Observable signal that the design is broken
- **Test** — How we measure it

This standard was recommended by the Canonical Audit and is adopted here as a documentation expectation. It does not itself introduce new game mechanics.

---

**End of Master GDD v1.1 — Canonical Baseline**

This document is the current working canonical design reference. Future system-specific specifications should build upon, and remain consistent with, the LOCKED principles, DERIVED CONSTRAINTS, and status discipline established here. No further LLM consolidation of the Master GDD should occur; the next operations are evidence reconciliation and substantive systems design (Economic Interdependence Core package).

**Master GDD v1.1 Canonical Baseline** is complete.

**Deliverable:**

---

### Scope of this pass (status/authority only)

| Change          | Detail                                                                                                |
| --------------- | ----------------------------------------------------------------------------------------------------- |
| Version         | 1.0 Consolidated → **1.1 Canonical Baseline**                                                         |
| Document status | Working Canonical Reference (status/authority pass complete)                                          |
| New status      | **DERIVED CONSTRAINT** added to vocabulary and Critical Rule                                          |
| Reclassified    | CRFT-006, MFG-004, ECO-003, CMBT-003, PROG-002 → DERIVED CONSTRAINT (with parent principles recorded) |
| Sequence        | Updated to match audit phase model (Evidence Reconciliation → Economic Interdependence Core package)  |
| Provenance      | v1.1 change notes + inherited consolidation notes                                                     |
| Design standard | Principle → Constraint → Invariant → Failure Condition → Test adopted as documentation expectation    |

**No new game-design decisions, mechanics, numbers, or system expansions were introduced.**

---

### Current project state

| Item                           | State                                                                                                             |
| ------------------------------ | ----------------------------------------------------------------------------------------------------------------- |
| Master GDD                     | **v1.1 Canonical Baseline** — Working Canonical Reference                                                         |
| Fully frozen canon             | Not yet (HUMAN-LOCKED set still needs explicit human confirmation; prototype reconciliation open)                 |
| Next process step              | Evidence Reconciliation Pass (prototype)                                                                          |
| First substantive design phase | Economic Interdependence Core (Specialisation ↔ Resources ↔ Crafting ↔ Manufacturing ↔ NPC substitution ↔ Demand) |
| Further LLM consolidation      | **Do not perform**                                                                                                |

Ready for your direction on the Evidence Reconciliation Pass or the Economic Interdependence Core package.

## Response:
25/08/2026, 15:16:37

> **Auditing document consistency**
>
> I’m comparing the pasted specification with the uploaded version, and I’m finding that the “no new decisions” claim is too strong. Several mechanics appear newly reintroduced, while authority labels and lock status need clarification.
>
> **Audited document consistency**
>
> 
>
> **Assessed and summarized TCIndustries’ canonical game design baseline**
>
> 
>
> Worked for 8s

I’ve reviewed the material against the uploaded project files and the v1.0 audit history. There is an important distinction to make first:

**The document you pasted as “v1.1 - Canonical Baseline” is materially different from the uploaded `TCIndustries_Master_GDD-chatgptv1-1.md`.** The uploaded file is actually the older **First-Pass v0.1** document, despite the filename. `TCIndustries_Master_GDD-chatgptv1-1.md`

So my assessment below is of the **v1.1 text you pasted**, cross-checked against the v1.0 consolidated GDD and its source lineage.

# Executive assessment

**Verdict: YELLOW/RED - structurally improved, but I would NOT accept this as a true “status/authority-only” v1.1 canonical baseline yet.**

The document is substantially better than v1.0 in one respect: the authority model is clearer and the five identified over-promoted rules have been reclassified.

However, there is a more important problem:

> **v1.1 claims to contain no new game-design decisions, but it actually reintroduces a significant amount of detailed design that was deliberately left PROPOSED or excluded during the v1.0 consolidation.**

That is the main issue I would correct before treating v1.1 as the canonical baseline.

---

# 1. What v1.1 gets right

The core audit correction was implemented correctly.

### Correctly added

**DERIVED CONSTRAINT** has been added to the status vocabulary, with an explicit rule that it is not equivalent to LOCKED.

That is exactly the right direction.

### Correctly reclassified

These five were changed:

| Rule | v1.0 | v1.1 | Assessment |
|---|---|---|---|
| CRFT-006 | LOCKED | DERIVED CONSTRAINT | Correct direction |
| MFG-004 | LOCKED | DERIVED CONSTRAINT | Correct |
| ECO-003 | LOCKED | DERIVED CONSTRAINT | Correct |
| CMBT-003 | LOCKED | DERIVED CONSTRAINT | Correct |
| PROG-002 | LOCKED | DERIVED CONSTRAINT | Correct direction |

This addresses the central authority-hygiene problem identified by the audit.

The v1.0 consolidation explicitly said that detailed mechanical material remained PROPOSED and that agreement among LLM documents did not constitute human approval. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

### Good preservation of uncertainty

The v1.1 document continues to keep:

- resource mathematics TBD,
- skill costs TBD,
- crafting formulas TBD,
- manufacturing throughput TBD,
- cities/governance TBD,
- combat scope TBD,
- currency TBD,
- respecialisation costs TBD.

That is good design governance.

### Good evidence discipline

The prototype relationship is also appropriately cautious. The consolidated GDD explicitly stated that no prototype was inspected and that prototype behaviour remained evidence rather than authority. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

That remains the correct position.

---

# 2. The major problem: v1.1 is NOT actually status-only

This is the most important finding.

The document says:

> “No new game-design decisions are introduced.”

But that is not strictly true.

It contains substantial mechanical content that is more specific than the v1.0 consolidated GDD.

The v1.0 consolidation deliberately said that detailed taxonomies, skill lists, attribute examples and mechanical shapes were being retained conservatively as PROPOSED, while more speculative material was excluded. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

The v1.1 you pasted now contains things such as:

### Resource system

It specifies:

- Resource Domain → Category → Specific Resource Class
- a universal attribute pool
- **Toxicity**
- time-limited rather than extraction-limited resource pools
- short/medium/long resource lifespans
- surveying information tiers
- manual harvesting vs installed Harvesters
- processing as a separate player activity.

Those are **design decisions/shapes**, even if marked PROPOSED.

For example, the statement that resources are deliberately **time-limited rather than extraction-limited** is not merely authority hygiene. It chooses one economic model over another.

That is a substantive design decision.

---

# 3. Manufacturing has an even bigger problem

The v1.1 text says:

> “Factories can never produce Exceptional-tier output.”

The document correctly labels this **PROPOSED** rather than LOCKED.

That is good status discipline.

But the problem is that this rule had specifically been identified in the prior audit as material that was **not present as LOCKED in the consolidated GDD**.

The audit explicitly noted:

> “The consolidated GDD does not contain a LOCKED rule that ‘factories can never produce Exceptional-tier output.’”

So reintroducing it in v1.1 is not a status correction.

It is a **design proposal being reintroduced into the canonical document**.

That may ultimately be a very good design decision.

But it violates the stated purpose of v1.1:

> “No new game-design decisions are introduced.”

### My recommendation

Do **not** delete the idea.

Instead, either:

**Option A - preferred for a true baseline**

Move it into a clearly labelled:

> **Candidate Design Proposals / Deferred Proposals**

section outside the canonical rule body.

Or:

**Option B**

Keep it in the GDD as PROPOSED, but explicitly state:

> “Inherited from pre-consolidation source material; not part of the v1.1 authority correction.”

That would make the provenance honest.

---

# 4. The same problem exists in several other places

The v1.1 document contains many substantive mechanisms that are stronger than the consolidated v1.0.

Examples include:

### Character system

- Vitality / Stamina / Focus pools
- Skill Discipline / Skill Box architecture
- use-based Skill Point acquisition
- two-to-three-profession target
- friction-based respecialisation

These were present in the older first-pass source. `TCIndustries_Master_GDD-chatgptv1-1.md`

They are legitimate proposals, but they aren't simply authority corrections.

### Crafting

v1.1 specifies:

- experimentation point allocation,
- success/critical success/failure,
- experimentation risk,
- exceptional outcomes,
- permanent provenance.

Again: these are **design proposals**, not mere canonical-status corrections.

### Economy

v1.1 introduces/retains:

- single unified currency as the default,
- player-set pricing,
- baseline NPC goods,
- degradation as replacement-demand engine,
- Commerce Directory.

These are substantive economic architecture choices.

### Items

It introduces a particularly strong proposal:

> “fully tradeable by default, with no account/character soul-binding.”

That's a real design policy.

### Services

It proposes:

> services must be recurring/consumable by design.

Again, substantive design.

---

# 5. This creates a provenance problem

The v1.1 document says:

> “This version applies only the authority/status corrections approved in the Canonical Audit of v1.0.”

That statement is therefore too strong.

The actual document is closer to:

> **v1.0 Consolidated GDD + authority corrections + selected reintroduced material from earlier source GDDs.**

That distinction matters enormously for your cross-LLM audit system.

Otherwise, six months from now another LLM could look at v1.1 and conclude:

> “This was already part of the canonical baseline.”

When in reality some of it was merely an old LLM proposal that has been carried forward.

That is exactly the kind of **authority contamination** your audit framework is designed to prevent.

---

# 6. There is a deeper DERIVED CONSTRAINT problem

I would also revise one of the five classifications.

## CRFT-006

It says:

> derives from PIL-003 + specialisation principles.

But the “specialisation principles” themselves are not LOCKED.

The v1.1 GDD explicitly keeps the actual specialisation budget as TBD/PROPOSED.

Therefore:

**LOCKED Interdependence does not necessarily logically entail a particular implementation rule such as “No Instant Mastery of Everything.”**

Interdependence could theoretically be achieved through:

- resource dependencies,
- service dependencies,
- geographic dependencies,
- organisational dependencies,
- equipment dependencies,
- reputation,
- production chains,

without imposing a character skill cap.

So I would classify CRFT-006 as:

**PROPOSED / DESIGN CONSTRAINT**

rather than a strict DERIVED CONSTRAINT.

The same reasoning applies, to a lesser extent, to:

## PROG-002 - Multiple Viable Paths

“Non-combat viability” does not logically prove that **multiple** paths must be viable.

It proves at least one meaningful non-combat path must exist.

Multiple viable professions is a stronger design proposition.

Therefore I'd also reconsider its DERIVED CONSTRAINT status.

---

# 7. The distinction you need is actually three-layered

I think your project would benefit from separating:

### 1. HUMAN-LOCKED

> The human explicitly decided this.

Example:

**“Combat must not be the only meaningful career.”**

### 2. DERIVED CONSTRAINT

> This necessarily follows from a HUMAN-LOCKED principle.

Example:

If combat is not allowed to be the only meaningful career, then:

> “A system that makes combat mandatory for accessing all meaningful progression violates the design.”

That's genuinely derived.

### 3. DESIGN PROPOSAL

> This is one possible implementation of the principle.

Example:

> “Factories cannot produce Exceptional items.”

That may be an excellent solution.

But it is **not logically inevitable** from “automation must not replace player relevance.”

Another design could achieve the same goal differently.

This distinction is extremely valuable.

---

# 8. The current GDD sometimes blurs those layers

For example:

> “Interdependence should arise through distinct skills, specialised resources, intermediate components, service needs, regional logistics, knowledge gaps, reputation, production capacity limits…”

That's actually a **design hypothesis**, not necessarily a derived rule.

Likewise:

> “The cap should be generous enough to support roughly two to three concurrent professions…”

That's a concrete design target.

Even though the numerical value is TBD, the **structural target itself is a design choice**.

This should remain clearly PROPOSED.

---

# 9. Another important issue: “LOCKED” authority is being inferred too broadly

The v1.1 document states things such as:

> “WRLD-001 - Persistent Shared World - LOCKED”

and

> “PLR-001 - Persistent Character Identity - LOCKED”

Those may indeed be valid consequences of the project's core vision.

But your own authority rules say:

> Logical consequences of LOCKED principles are DERIVED CONSTRAINT, not LOCKED.

So there is an internal tension.

If:

> “Persistent virtual society”

is human-locked,

then:

> “persistent character identity”

might be a reasonable **derived constraint**.

It does not necessarily follow that the human independently decided:

> “Character identity must persist in precisely this way.”

This is a subtle but important distinction.

I would not automatically demote all of these, because some may have genuinely been explicitly decided in the original project brief.

But they need **provenance evidence**, not just:

> “required by pillars.”

“Required by pillars” is an inference.

---

# 10. The biggest conceptual improvement I recommend

Do not make v1.2 another giant GDD rewrite.

Instead, create a small **Authority Matrix**.

Something like:

| ID | Statement | Status | Authority Type | Source | Human-confirmed? |
|---|---|---|---|---|---|
| VIS-001 | Living virtual world | LOCKED | Explicit human | Project brief | YES |
| VIS-004 | Interdependence | LOCKED | Explicit human | Project brief | YES |
| CRFT-006 | No instant mastery | PROPOSED | Design inference | LLM sources | NO |
| MFG-002 | Factories cannot make Exceptional | PROPOSED | Design proposal | Earlier GDD | NO |
| RES-004 | Time-limited resource pools | PROPOSED | Design proposal | Earlier GDD | NO |
| CRF-004 | Permanent provenance | PROPOSED | Design proposal | Earlier GDD | NO |

That would make the canonical state dramatically easier to audit.

---

# 11. I would change the document status slightly

Currently:

> **Working Canonical Reference (status/authority pass complete)**

I think that's slightly premature.

I'd use:

> **Working Canonical Reference - Authority Baseline**

and explicitly state:

> “This document is canonical as a reference structure and authority ledger. It contains both HUMAN-LOCKED rules and unapproved PROPOSED design material. ‘Canonical’ does not mean every design statement is approved.”

That prevents a future LLM from interpreting the word **Canonical** as meaning “everything in this document is approved.”

---

# 12. What I would NOT do

I agree strongly with the previous audit on these points.

Do **not**:

- regenerate the GDD again;
- ask Grok/Claude/Gemini/etc. to “improve” it;
- start designing individual skill trees;
- start assigning numerical values;
- start designing combat independently;
- lock manufacturing before the interdependence problem is understood;
- treat multiple LLMs agreeing as human approval.

The original consolidation correctly preserved uncertainty rather than pretending the game was more mature than it is. `TCIndustries_Master_GDD_v1.0_Consolidated-grok.md`

---

# 13. Where the project actually stands

My current assessment is:

| Area | State |
|---|---|
| Vision | 🟢 Strong |
| Core pillars | 🟢 Strong |
| Anti-goals | 🟢 Strong |
| Authority framework | 🟢 Strong after v1.1 changes |
| Status vocabulary | 🟢 Strong |
| Provenance discipline | 🟡 Needs tightening |
| Canonicality claim | 🟡 Needs clarification |
| Mechanical design maturity | 🟡 Early |
| Economic interdependence | 🔴 Not yet demonstrated |
| Prototype evidence | 🔴 Unreconciled |
| Numerical balance | 🔴 Intentionally premature |
| Ready for major system design | 🟡 Almost |
| Ready for canonical freeze | 🔴 No |

---

# 14. My recommended next move

I would **not create another Master GDD version yet.**

Instead, perform a short:

## **TCIndustries Authority & Provenance Reconciliation Pass**

Its sole purpose would be to answer:

> **For every LOCKED and DERIVED CONSTRAINT statement in v1.1, can we prove why it has that authority?**

And simultaneously identify every substantive PROPOSED mechanic that was reintroduced from older GDDs.

The output should be a compact matrix, not another giant GDD.

Then we can establish:

### Layer A - Canon

Human decisions.

### Layer B - Derived constraints

Logical consequences of Layer A.

### Layer C - Design proposals

Candidate mechanics awaiting human approval.

### Layer D - Evidence

Prototype/simulation/playtest results.

### Layer E - Historical/reference

SWG and other external inspiration.

That would give TCIndustries a **very robust multi-LLM governance architecture**.

---

## Bottom line

**Grok's v1.1 is better than v1.0, but I would not yet call it a clean Canonical Baseline.**

The five status corrections were correctly applied. The real flaw is that **the document simultaneously claims to be a status-only correction while reintroducing a large amount of substantive design from earlier LLM drafts.**

That is fixable without another consolidation.

**My recommendation is therefore:**

> **Freeze v1.1 as a working draft, do not redesign the GDD, and perform one Authority & Provenance Reconciliation Pass before beginning the Economic Interdependence Core.**

And there is one especially important correction I would make during that pass:

> **Do not allow “logical consequence” to become a back door for turning a desirable design proposal into a pseudo-LOCKED rule.**

That distinction is likely to become the single most important governance mechanism for TCIndustries as multiple LLMs continue contributing to the project.

## Prompt:
25/08/2026, 18:27:45

This chat highlighted something I missed in the GDD:
Manufacturing facilities (factories, workshops, processing plants) may exist as player-owned or organisation-owned structures that convert inputs into outputs over time under defined rules.

“Factories can never produce Exceptional-tier output.” &#x20;

In star wars galaxes pre-cu factories were player created structures which performed the following tasks:
You can use a [factory](https://swg.fandom.com/wiki/Factory) to manufacture crates of identical items from a [Manufacturing Schematic](https://swg.fandom.com/wiki/Manufacturing_Schematic). These items may be for sale or to be used to create other items; some schematics require items created in the same factory run (identical serial numbers). A common mistake is to make a schematic with one resource you have on hand, and then try to run it through the factory with another resource. They have to be exactly the same, down to the specific name. The same applies for subcomponents; the serial numbers must be exactly the same. &#x20;

The process is:
Players create items
Players create schematic using recipe created from items. The recipe contains multiple different t ypes of resources and / or components. When creating the item the player can choose to create a schematic as well as the item.
The schematic's quality is the same quality as the item.
Factories are given schematics and can hold a number of resources and components.
Over time, if they have enough resources and components, they will create 1 item. The item build time is based on difficulty. The quality and item are identical. These identical items are then held in item crates. item crates hold a number of identical items.

Can you verify this process is true from the GDD?

## Response:
25/08/2026, 18:27:47

> **Verifying Factory Rules Against GDD v1.1**
>
> 
>
> Worked for 11s

Yes - **you identified a real gap**, but there is an important distinction between what the current v1.1 GDD actually establishes and what you are proposing to add.

## Verdict

Your proposed manufacturing process is **only partially represented in Master GDD v1.1**.

The GDD establishes the **conceptual framework**, but it does **not currently specify the full factory workflow you described**.

More importantly:

> **“Factories can never produce Exceptional-tier output” is NOT currently LOCKED in v1.1.**

It existed in earlier source material as a **PROPOSED** rule recommended for promotion, but v1.1 deliberately did not introduce that decision. The earlier source explicitly identifies it as proposed/recommended rather than approved. `TCIndustries_Master_GDD-chatgptv1-1.md`

### What v1.1 currently establishes

| Your proposed rule/process | v1.1 status | Assessment |
|---|---|---|
| Player/organisation-owned manufacturing facilities | **PROPOSED / implied** | Supported conceptually |
| Factory converts inputs into outputs | **PROPOSED** | Explicitly supported |
| Factory operates over time | **PROPOSED** | Queue time is mentioned |
| Factory uses a schematic | **PROPOSED** | Explicit |
| Schematic specifies inputs/components | **PROPOSED** | Explicit |
| Player first creates/crafts the item | **NOT specified** | Missing |
| Player creates a schematic from that crafted item | **NOT specified** | Missing |
| Schematic captures the quality of the crafted item | **NOT specified** | Missing |
| Factory receives that exact configured schematic/template | **NOT specified in v1.1** | Missing |
| Factory requires exact matching resources/components | **NOT specified** | Missing |
| Factory produces identical items | **Partially implied** | "Repeatable items", but not formally defined as identical |
| Items are produced individually over time | **Partially implied** | Queue time exists, but production unit/process is unspecified |
| Identical output is accumulated into crates | **NOT specified** | Missing |
| Factory reproduces the template's quality | **NOT specified** | Missing |
| Factory can never produce Exceptional | **NOT LOCKED** | Important missing rule |
| Manual crafting can produce Exceptional | **PROPOSED/partially represented** | Signature exceptional products are supported, but exact quality tiers are unresolved |

The current GDD does say that factories may produce **repeatable items from approved schematics and supplied inputs**, with constraints including input materials, component availability, factory capacity and queue time. It also says manufacturing quality should generally have a lower ceiling than direct specialist crafting. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

It also explicitly makes **manufacturing quality** one component of the eventual product-quality model. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

So the *shape* is there, but the **actual manufacturing primitive is underspecified**.

---

# The important discovery: this was already present in the source lineage

What you remembered is not something we would be inventing from scratch.

The earlier Claude consolidation contains a much more explicit manufacturing model:

> A factory mass-produces a schematic only after a crafter has manually crafted and "locked in" a specific configured version - a **Production Template**.

That document then states that the factory repeats that configuration at volume. `TCIndustries_Master_GDD-claude1.md`

It also explicitly proposed:

> **Factories can never produce Exceptional-tier output.**

But critically, that source marked it:

**PROPOSED - flagged for early lock**

and recommended that it be promoted to LOCKED. `TCIndustries_Master_GDD-claude1.md`

Likewise, the earlier Claude material explicitly proposed throughput, maintenance, rising marginal upkeep, and land scarcity as factory constraints. `TCIndustries_Master_GDD-claude1.md`

So this is **not a new idea appearing out of nowhere**. It is an unresolved manufacturing design element that was present in the source material but was intentionally not promoted into v1.1.

---

# Your SWG-derived process is more specific than the current GDD

The process you described is essentially:

```text
CRAFTING
   │
   ├─ Player selects recipe/schematic
   ├─ Player selects exact resources/components
   ├─ Player crafts item
   ├─ Player experiments/refines
   └─ Resulting item has specific quality
            │
            ▼
      PRODUCTION TEMPLATE
            │
            │ captures exact configuration
            ▼
         FACTORY
            │
     ┌──────┴────────┐
     │               │
 exact inputs     exact components
     │               │
     └──────┬────────┘
            ▼
      Production Queue
            │
       time passes
            │
            ▼
       identical item
            │
            ▼
       another item
            │
            ▼
       another item
            │
            ▼
       ITEM CRATE
```

That is **considerably more precise** than v1.1.

And it creates an important conceptual distinction:

### Manual crafting

**Discovery → design → experimentation → exceptional product**

versus

### Manufacturing

**Repetition → scale → consistency → volume**

That distinction fits extremely well with the GDD's existing anti-automation philosophy.

The earlier source actually articulated this as:

> Automation may exist for repetition; it must never exist for discovery, for breaking a quality ceiling, or for identity. `TCIndustries_Master_GDD-claude1.md`

That is arguably the cleanest articulation of what your proposed system is trying to achieve.

---

# One correction to your proposed terminology

I would **not yet say**:

> "The schematic's quality is the same quality as the item."

That's potentially ambiguous.

The stronger design concept is:

> **A Production Template captures a specific manually crafted configuration and its permitted manufacturing quality/output state.**

Why?

Because a **schematic** and a **Production Template** should probably be different concepts.

### Schematic

Defines:

- what can be made;
- required resources;
- component slots;
- skill requirements;
- experimentation possibilities;
- quality boundaries;
- output definition.

### Production Template

Defines:

- the specific product configuration selected for repetition;
- the actual resource/component configuration;
- the applicable schematic version;
- the manufacturing-eligible quality/output state;
- possibly cosmetic/customisation choices;
- provenance of the originating crafter.

That distinction is much cleaner and is already latent in the earlier source's "Production Template" proposal. `TCIndustries_Master_GDD-claude1.md`

---

# The Exceptional rule is especially important

I agree that this is something the GDD needs to explicitly resolve.

The current v1.1 wording:

> "Manufactured products should generally have a lower quality ceiling..."

is **too weak** if the intended rule is actually:

> **Factories can never produce Exceptional-tier output.**

"Generally" and "never" have radically different authority.

The current v1.1 rule permits a future designer to argue:

> "This factory is sufficiently advanced that it can produce Exceptional items."

That would be inconsistent with the intended manufacturing philosophy if the actual design decision is that **Exceptional is intrinsically a manual crafter achievement**.

The earlier source correctly identified this as a load-bearing anti-Factorio rule. `TCIndustries_Master_GDD-claude1.md`

---

# I would classify the discovery this way

### Already supported by v1.1

**LOCKED**

- Manufacturing exists to support production scale rather than replace players.
- Manufacturing must preserve player specialisation/social interaction.
- Increasing scale must require coordination.
- Crafting produces differentiated products.
- Resource/component quality matters.
- Exceptional products associated with players are part of the intended identity direction.

### Existing but unresolved

**PROPOSED**

- Factories produce repeatable products.
- Factories use schematics and supplied inputs.
- Factories have production constraints.
- Manufacturing has a lower quality/customisation ceiling.
- Manufacturing operates over time/queue.
- Factory maintenance/capacity/ownership/organisation constraints.

### Missing from v1.1

**Design gap**

- Manual item → Production Template conversion.
- Exact template semantics.
- Template quality capture.
- Exact input identity requirements.
- Exact component identity requirements.
- Whether inputs must be identical to the template's original inputs.
- One-at-a-time production semantics.
- Identical-output semantics.
- Item crate/container semantics.
- Crate capacity.
- Factory run lifecycle.
- Failure behaviour when inputs don't match.
- Whether factory output retains original crafter provenance.
- Whether factory output retains template/resource provenance.
- **Explicit factory Exceptional-quality prohibition.**

---

## The key conclusion

So **yes, your observation is correct**:

> The GDD has the **manufacturing philosophy**, but it does not yet have the **manufacturing process specification**.

And the particularly important point is that **the "Factories can never produce Exceptional-tier output" rule was already present in the project lineage, but v1.1 correctly left it PROPOSED rather than silently promoting it to LOCKED.** `TCIndustries_Master_GDD-chatgptv1-1.md`

Therefore, I would **not modify the Master GDD yet**. This is exactly the kind of thing that belongs in the upcoming **Economic Interdependence Core** design work.

The manufacturing package should explicitly resolve:

**Manual Craft → Production Template → Factory Input Validation → Timed Production → Identical Output → Crating → Provenance → Quality Ceiling**

and then determine whether:

> **Factories can never produce Exceptional-tier output**

is a human LOCKED rule or merely a design recommendation.

That would turn the current vague `MFG-002/MFG-003` framework into an actual manufacturing system specification without violating the v1.1 rule against premature consolidation.



---
Powered by [ChatGPT Exporter](https://www.chatgptexporter.com)