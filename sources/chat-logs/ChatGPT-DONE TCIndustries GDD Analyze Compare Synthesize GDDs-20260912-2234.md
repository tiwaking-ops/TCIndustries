# DONE TCIndustries GDD Analyze Compare Synthesize GDDs

**User:** Anonymous  
**Created:** 8/20/2026 22:00:28  
**Updated:** 8/22/2026 13:22:44  
**Exported:** 9/12/2026 22:34:25  
**Link:** [https://chatgpt.com/g/g-p-6a8605e707e4819196e35afd1d292bd4/c/6a86cfaa-4e08-83ec-b00a-a9fd727c5a22?src=history_search&messageId=f4490d09-f91b-4338-a296-840199104ccc&historySearchQuery=PLR-003](https://chatgpt.com/g/g-p-6a8605e707e4819196e35afd1d292bd4/c/6a86cfaa-4e08-83ec-b00a-a9fd727c5a22?src=history_search&messageId=f4490d09-f91b-4338-a296-840199104ccc&historySearchQuery=PLR-003)  

## Prompt:
20/08/2026, 22:00:26

I want to think this through properly before you answer: I want you to Analyze, compare, and contrast the three TCIndustries GDD's in the project files. identify which is the best GDD. output in markdown a formal report of your findings. if possible, synthesize all three and output in markdown the result as a formal document.

1. Interview me first. Ask the smallest set of questions that would remove real ambiguity about my goal, audience, constraints, and what "great" looks like here. Skip anything you can reasonably infer.
2. Once you have enough, draft a tailored prompt with:
   - Role — only if it changes vocabulary, judgment calls, or tone
   - Context — what I'm trying to accomplish and why
   - Audience — who this ultimately reaches
   - Format & length — structure and how long
   - Success criteria — 2–3 things separating a great answer from a mediocre one here
   - Constraints — hard boundaries or things to avoid
   - Examples — reference anything I've shared, or note that none exists
3. Show me the prompt, then answer it — flag anywhere you had to guess.

## Response:
20/08/2026, 22:00:27

Yes. Before I compare the three GDDs, I want to establish what **“best”** means for this project, because the documents appear to differ partly in *purpose and level of specification*, not merely quality.

I can infer most of the background from the documents and our project history, so I only need **four decisions**:

1. **The three documents:** I currently have two TCIndustries GDDs directly available. What is the **third GDD** you want included? If it is in the project File Library, give me its filename or distinctive phrase and I’ll locate it.

2. **Primary purpose of the synthesis:** Which should the resulting GDD optimize for?
   - **A. Canonical design authority** — the definitive source of truth from which all later system specifications derive.
   - **B. Development blueprint** — optimized for actually building the game, with enough specificity to guide prototypes and implementation.
   - **C. Both**, with a clear separation between canonical design decisions and implementation/development guidance.

3. **Audience:** Who is the eventual primary reader?
   - **A. You / yourself as project owner**
   - **B. AI collaborators (ChatGPT, Claude, Gemini, etc.)**
   - **C. Human game designers/developers**
   - **D. All three**

4. **How aggressive should the synthesis be?**
   - **A. Conservative:** preserve the strongest existing material and reconcile contradictions, but don't introduce much new design.
   - **B. Editorial:** combine the best material and improve structure, terminology, precision, and identified weaknesses, while clearly marking anything newly inferred.
   - **C. Architectural:** treat the three GDDs as competing design proposals, select the strongest ideas, resolve conflicts, identify missing architecture, and produce the best coherent canonical GDD—even where that requires making reasoned design decisions.

Once you answer those four, I’ll:

1. Locate and fully inspect all three GDDs.
2. Compare them systematically across **vision, design philosophy, system architecture, economy, resources, crafting, manufacturing, progression, social systems, world simulation, extensibility, design governance, completeness, ambiguity, and development usefulness**.
3. Identify **which GDD is actually strongest and why**, rather than simply choosing the longest or most detailed one.
4. Identify where each of the other GDDs contains superior material.
5. Produce the tailored prompt and show it to you **before execution**, as requested.
6. Execute that prompt and produce the formal Markdown report.
7. Then produce a **synthesized canonical TCIndustries GDD** in Markdown, clearly flagging every place where I had to infer, reconcile, or make a design judgment.

I will distinguish **what the three source documents actually say** from my comparative judgment and from any new synthesis.

## Prompt:
20/08/2026, 22:31:23

1. There are 4 now: TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md , Tiwakings Craftworld Industries — Master Game Design Document v0.1-chatgpt1.md, TCIndustries\_Master\_GDD\_v1.0-claude-opus-4-6-thiking.md, TCIndustries\_Master\_GDD-claude1
2. C
3. B
4. A

## Response:
20/08/2026, 22:31:24

> **Compared and synthesized four TCIndustries game design documents**
>
> 
>
> Worked for 10s

# TCIndustries GDD Comparative Analysis and Canonical Synthesis

## 0. Executive finding

I have now treated the four files as **four competing design documents**, rather than assuming that the newest, longest, or most polished document is automatically the best.

My conclusion is:

> **The Claude Opus 4.6 v1.0 GDD is the best existing GDD overall.**
>
> **However, it should not become the final canonical GDD unchanged.**
>
> The strongest final document should be a **synthesis of the Opus v1.0 and the Terra-medium v0.1, with selected architectural material from Claude1 and selected clarity/material from ChatGPT1.**

In particular:

- **Claude Opus v1.0** is the best *overall architecture and governance document*.
- **GPT-5.6 Terra-medium v0.1** is extremely strong and contributes some of the best **design invariants, open-question structure, system boundaries, and canonical discipline**.
- **Claude1** contains some of the most useful **mechanical elaboration**, particularly where it converts principles into explicit system interactions and failure modes.
- **ChatGPT1** is the cleanest and most conservative **foundational GDD**, but is now substantially superseded by the more developed documents.

The important conclusion is therefore **not "pick Claude Opus and discard the others."**

It is:

> **Use Claude Opus v1.0 as the architectural spine, Terra as the governance/testing spine, Claude1 as a source of carefully selected mechanical elaboration, and ChatGPT1 as a clean baseline/reference.**

That distinction matters considerably.

---

# 1. Tailored prompt

As requested, I am showing the prompt I would use to perform the actual synthesis.

```markdown
# Role

Act as a senior game-systems architect and GDD editor specialising in persistent-world sandbox games, player-driven economies, systemic MMORPGs, and design-governance documentation.

Your task is not merely to merge documents. Treat the four existing TCIndustries GDDs as competing design drafts and determine which ideas, structures, and rules deserve to survive into a canonical document.

# Context

TCIndustries is intended to be a persistent virtual-world sandbox MMORPG inspired by the systemic philosophy of the original Star Wars Galaxies, but it is an original IP and must not reproduce Star Wars content.

The project's central fantasy is that players live as consequential citizens of a persistent world rather than as predetermined heroes.

The game is intended to support:

- player-driven economic activity;
- dynamic and geographically variable resources;
- resource quality and differentiation;
- crafting and manufacturing;
- meaningful player specialisation;
- player services;
- reputation and provenance;
- player commerce;
- ownership;
- player organisations;
- settlements and cities;
- combat as one important activity rather than the mandatory centre of play;
- emergent gameplay;
- meaningful interdependence between players;
- expandable content architecture.

The project deliberately does not want to become a conventional theme-park MMO, rigid class MMO, Factorio/Satisfactory-style automation game, Capitalism-style spreadsheet simulator, or a single-player industrial game with multiplayer attached.

Four GDDs must be analysed:

1. TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md
2. Tiwakings Craftworld Industries - Master Game Design Document v0.1-chatgpt1.md
3. TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md
4. TCIndustries_Master_GDD-claude1.md

The desired final architecture is BOTH:
- a canonical design authority; and
- a development/design blueprint.

The primary audience is AI collaborators such as ChatGPT, Claude, Gemini, and other future design/development agents.

The synthesis should be conservative: preserve strong existing design decisions and reconcile contradictions, but do not gratuitously invent new game mechanics.

# Audience

The final document should be usable by:

1. AI collaborators as a reliable source of truth;
2. the project owner as a design-governance document;
3. future human designers/developers as a clear statement of intended game behaviour.

AI readability and unambiguous status classification are particularly important.

# Format and length

Produce two formal Markdown documents:

## Document A - Comparative Assessment

Include:

1. Executive finding
2. Evaluation methodology
3. Summary comparison of all four GDDs
4. Detailed strengths and weaknesses of each
5. Comparison by major design domain
6. Design-governance comparison
7. System-completeness comparison
8. Identification of contradictions and unresolved tensions
9. Best material contributed by each GDD
10. Overall ranking
11. Recommendation for the canonical source
12. Required synthesis strategy

## Document B - Synthesized Canonical Master GDD

Create a coherent canonical TCIndustries Master GDD using the best compatible material from all four documents.

Preserve explicit status distinctions:

- LOCKED
- PROPOSED
- TBD
- PROTOTYPE
- DEFERRED

Do not silently convert proposed material into locked design.

Where source documents disagree:

1. prefer explicit project decisions over designer invention;
2. prefer compatible established principles over newly invented mechanics;
3. preserve genuinely unresolved questions as TBD;
4. use the strongest existing recommendation where the project has not decided;
5. record consequential assumptions rather than hiding them;
6. do not invent numerical values merely to make the document appear complete.

The final GDD should include, where supported:

- executive summary;
- design status system;
- vision;
- design pillars;
- player experience and core loops;
- anti-goals;
- world philosophy;
- world simulation;
- character identity and ownership;
- skills/professions/progression;
- resources;
- crafting;
- manufacturing and automation;
- items/provenance/degradation;
- economy;
- retail/commerce;
- services;
- combat;
- buildings/housing/cities;
- social organisations;
- emergent gameplay;
- content expansion;
- balance philosophy;
- exploit/fraud/economic safeguards;
- testing requirements and invariants;
- historical SWG relationship;
- prototype relationship;
- deeper-design roadmap;
- status register;
- assumption register;
- open-question register;
- glossary.

# Success criteria

A great result must:

1. Identify the genuinely strongest design architecture rather than simply selecting the longest or newest document.
2. Preserve the distinction between actual project decisions and designer recommendations.
3. Produce a synthesis that is internally coherent and more useful than any individual source document.
4. Prevent accidental canonisation of speculative mechanics.

# Constraints

- Do not treat document version number as proof of superiority.
- Do not reward verbosity merely for being verbose.
- Do not erase useful disagreements merely to make the synthesis look clean.
- Do not invent numerical rules without source support or explicit approval.
- Do not allow prototype behaviour to override the GDD.
- Do not introduce Star Wars intellectual property into the design.
- Do not make technical implementation decisions that belong in a future TDD.
- Clearly identify any judgement call made during synthesis.
- Treat the four documents as design evidence, not as equally authoritative sources.
- Be conservative in changing established project direction.

# Examples

Relevant source material is contained in the four uploaded TCIndustries GDDs and the project context.

No external examples are required.

# Required final judgement

Explicitly answer:

- Which GDD is best overall?
- Which GDD has the best architecture?
- Which has the best mechanical detail?
- Which has the best design governance?
- Which has the best testing/balance discipline?
- Which has the best readability for AI collaboration?
- What should be retained from each?
- What should be rejected?
- What should become the canonical synthesis?
```

## Important note on assumptions

I made **one important interpretation** rather than asking another question:

I have treated your answer **"A" for conservative synthesis** as meaning *do not substantially redesign TCIndustries merely because one draft proposes a potentially better mechanic*. The synthesis can improve structure and reconcile contradictions, but it should not become a fourth designer inventing an entirely new game.

That seems the most faithful interpretation of your stated goal.

---

# 2. Evaluation methodology

I evaluated the four documents against seven criteria:

| Criterion | Weight |
|---|---:|
| Design architecture and coherence | 20% |
| Fidelity to the project's core vision | 15% |
| Design-status/governance discipline | 15% |
| System coverage and completeness | 15% |
| Mechanical usefulness | 15% |
| Testing/balance/exploit discipline | 10% |
| AI collaboration usefulness | 10% |

I deliberately **did not** score length directly.

This is important because the later Claude documents are longer partly because they have more elaboration-not necessarily because every elaboration is correct.

---

# 3. The four documents

## 3.1 GPT-5.6 Terra-medium v0.1

**Character:** architectural foundation.

This document establishes a particularly strong distinction between:

- locked design;
- proposed design;
- unresolved questions;
- prototype behaviour;
- deferred systems.

It also explicitly defines the GDD as establishing **"canonical design architecture, philosophy, system boundaries, initial rules direction, unresolved questions, and future design sequence."** `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

Its strongest contribution is therefore **design governance**.

It also contains unusually useful acceptance/invariant thinking. For example, it states that items cannot appear without valid sources, manufacturing cannot produce without valid inputs/capacity/permissions, ownership transfer must be auditable, and ordinary systems must not permit infinite risk-free currency or item generation. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

### Strengths

- Excellent canonical boundaries.
- Strong status discipline.
- Strong open-question register.
- Strong testing philosophy.
- Strong recognition of economy/exploit risks.
- Clear system dependencies.
- Very good distinction between game design and implementation.
- Particularly good foundation for future AI-assisted design.

### Weaknesses

Its main weakness is that it sometimes stops at the **architecture of a system** rather than developing the strongest proposed implementation.

That is not necessarily a flaw at this stage, but it means another document can provide more immediately actionable system design.

---

# 4. ChatGPT1 v0.1

**Character:** clean foundational specification.

The ChatGPT1 document is remarkably disciplined and readable.

It clearly establishes the foundational vision:

> living in a persistent world as an ordinary but consequential citizen

and defines player-driven economy, interdependence and emergence as locked principles. `Tiwakings Craftworld Industries — Master Game Design Document v0.1-chatgpt1.md`

Its manufacturing/economic sections are deliberately conservative. For example, it recognises automation while explicitly preventing automation from eliminating the social economy and favours many specialised participants over one player controlling everything. `Tiwakings Craftworld Industries — Master Game Design Document v0.1-chatgpt1.md`

### Strengths

- Excellent readability.
- Strong foundational philosophy.
- Very low speculative contamination.
- Good high-level economic model.
- Good profession/interdependence framing.
- Easy to use as an introductory reference.

### Weaknesses

It is now comparatively **underdeveloped**.

It identifies many things as TBD without giving enough architectural scaffolding for resolving them.

For example, it says resource attributes, crafting mathematics, economy, reputation, cities, and manufacturing limits require later design, but does not provide the same depth of testing/invariant framework found in the later documents.

### Verdict

**Excellent foundation, but no longer the strongest master document.**

I would retain it as the **conceptual baseline**, not the final authority.

---

# 5. Claude1 v0.1

**Character:** detailed first-pass systems architecture.

This is arguably the most interesting document to preserve *parts of*.

Its key strength is that it frequently answers:

> "What does this principle actually mean mechanically?"

For example, its resource system explicitly separates:

- discovery;
- extraction;
- processing;
- resource markets.

It proposes Surveying as a specialised information economy and distinguishes manual harvesting from installed harvesters. `TCIndustries_Master_GDD-claude1.md`

That is a meaningful advance over merely saying "resources should be dynamic."

It also includes explicit interaction and failure-mode analysis, which is excellent design practice.

Similarly, the retail system does not simply say "players can have shops." It connects provenance, reputation, location, commercial districts and discovery. `TCIndustries_Master_GDD-claude1.md`

### Strengths

- Best **mechanical elaboration** of the four.
- Excellent system interactions.
- Good failure-mode analysis.
- Strong anti-goal enforcement.
- Strong assumption register.
- Good content expansion architecture.
- Very useful AI-readable rule IDs.
- Particularly strong resource/crafting/manufacturing thinking.

### Weaknesses

The problem is precisely its strength:

> **It sometimes gets ahead of the project.**

Some mechanics are presented as very specific recommendations even though they are not actually locked project decisions.

The document itself acknowledges this, explicitly saying that taxonomies and abstractions are mostly proposed and that it intentionally avoids exact numbers. `TCIndustries_Master_GDD-claude1.md`

Its assumption register is excellent, but the number of assumptions demonstrates how much design has been introduced.

Examples include assumptions around:

- specialisation budgets;
- NPC vendor superiority/inferiority;
- city governance;
- full item tradeability;
- research;
- other systemic structures. `TCIndustries_Master_GDD-claude1.md`

### Verdict

**The best source of mechanical proposals, but not the best standalone canonical authority.**

---

# 6. Claude Opus 4.6 v1.0

**Character:** mature integrated master architecture.

This is the strongest document overall.

It takes the strengths of the earlier documents and combines them into a substantially more mature structure.

Its table of contents alone shows the difference in architectural maturity: it includes not only systems but **balance philosophy, exploit/fraud/economic stability, testing invariants, historical inspiration policy, prototype governance, deeper-design sequencing, assumptions, open questions, and glossary**. `TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

It also explicitly declares:

> the document is the authoritative design specification

while simultaneously preventing proposed or prototype material from silently becoming canonical. `TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

That is exactly the governance model this project needs.

### Strengths

#### 1. Best overall architecture

It connects:

**Vision → Pillars → Player loops → Systems → Balance → Exploits → Testing → Future specifications**

rather than treating each subsystem as an isolated chapter.

#### 2. Excellent interdependence enforcement

The document doesn't merely state that interdependence matters. It repeatedly asks whether a proposed system lets one player bypass another.

That is much stronger than a philosophical statement.

#### 3. Strong world architecture

It introduces useful distinctions between established settlements, player cities, wilderness, transportation and economic geography.

#### 4. Strong service design

It explicitly treats non-item services as first-class economic activity, including medicine, entertainment, repair, construction, research and transport. `TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

#### 5. Strong testing discipline

Its invariants are particularly valuable:

- no invalid item creation;
- provenance;
- valid schematic inputs;
- bounded extraction;
- valid manufacturing inputs;
- auditable ownership;
- currency conservation;
- anti-duplication;
- anti-infinite-resource/currency rules. `TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

#### 6. Strong prototype governance

It explicitly establishes:

**GDD → Prototype → Testing → Golden Master → Future UE5/C++**

and says the prototype is evidence, not authority. `TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

That is an extremely useful project-management distinction.

### Weaknesses

There are two.

#### Weakness 1 - Some proposals are too close to canon

Because this is called **v1.0 - Canonical Foundation**, there is a risk that readers will psychologically treat its proposed mechanics as approved.

The status system mitigates this, but not completely.

#### Weakness 2 - It inherits some speculative architecture

Some proposals originating in Claude1 have been incorporated into the larger architecture.

That is not necessarily wrong-but it means the synthesis should distinguish:

> **"This is a good design recommendation"**

from:

> **"The project has actually decided this."**

---

# 7. Comparative scorecard

My assessment is:

| GDD | Architecture | Vision | Governance | Coverage | Mechanics | Testing | AI usefulness | Overall |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| ChatGPT1 | 8 | 9 | 8 | 6 | 6 | 5 | 8 | **7.2** |
| Terra-medium | 9.5 | 9 | 9.5 | 8.5 | 8 | 9.5 | 9.5 | **9.0** |
| Claude1 | 9 | 9 | 9 | 9 | **9.5** | 8.5 | 9.5 | **9.1** |
| **Claude Opus v1.0** | **9.8** | **9.5** | **9.8** | **9.5** | **9.0** | **9.5** | **9.8** | **9.5** |

These numbers are not mathematical measurements; they are comparative editorial judgements.

The important result is the ranking:

### Overall

1. **Claude Opus 4.6 v1.0 - best overall**
2. **Claude1 - strongest mechanical source**
3. **Terra-medium - strongest governance/architecture source**
4. **ChatGPT1 - strongest clean foundation**

---

# 8. What each document contributes to the final GDD

This is where the synthesis becomes particularly useful.

| Source | Preserve | Do not blindly inherit |
|---|---|---|
| **ChatGPT1** | Vision, simplicity, anti-goals, economic philosophy | Sparse mechanics |
| **Terra** | Governance, invariants, open questions, system boundaries, testing | Some repetition / architectural breadth |
| **Claude1** | Resource architecture, system interactions, failure modes, assumptions, content architecture | Unapproved detailed assumptions |
| **Opus v1.0** | Overall structure, integration, services, world architecture, testing, prototype governance | Treating every proposal as de facto canonical |

---

# 9. The most important synthesis principle

I recommend the final document use **four layers of authority**, even though the status vocabulary remains the existing five-status system.

### Layer 1 - Project decisions

**LOCKED**

These are the actual design constitution.

Examples:

- living world;
- citizen fantasy;
- player-driven economy;
- interdependence;
- player-created society;
- non-combat viability;
- expandable architecture.

These are strongly convergent across the documents. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

### Layer 2 - Design recommendations

**PROPOSED**

These are good ideas that should be tested before becoming rules.

Examples:

- specific skill-domain architecture;
- resource pool algorithms;
- particular city governance models;
- specific transportation structures;
- specific reputation mechanisms.

### Layer 3 - unresolved decisions

**TBD**

The final GDD should be willing to say:

> "We don't know yet."

That is preferable to false precision.

### Layer 4 - future implementation/specification

**DEFERRED**

The GDD should not attempt to become the TDD.

---

# 10. Important contradiction/tension identified

The biggest issue across the four documents is **not a direct contradiction**.

It is a governance problem:

> **How much proposed design should be allowed into the Master GDD?**

Claude1 is strongest when answering "how could this work?"

Terra is strongest when answering "what status does this idea have?"

Opus is strongest when combining those two.

The final document should therefore retain detailed proposals **only where they clarify architecture without pretending that the mechanic is locked.**

---

# 11. Recommended canonical architecture

I recommend the final TCIndustries documentation hierarchy become:

```text
TCIndustries Master GDD
│
├── 1. Vision
├── 2. Design Constitution
│   ├── Locked pillars
│   ├── Anti-goals
│   └── Design invariants
│
├── 3. Player Experience
│
├── 4. World
│
├── 5. Character / Identity / Skills
│
├── 6. Resources
│
├── 7. Crafting
│
├── 8. Manufacturing / Automation
│
├── 9. Items / Provenance / Degradation
│
├── 10. Economy
│
├── 11. Commerce / Retail
│
├── 12. Services
│
├── 13. Buildings / Housing / Cities
│
├── 14. Organisations / Society
│
├── 15. Combat
│
├── 16. Emergence
│
├── 17. Content Expansion
│
├── 18. Balance Philosophy
│
├── 19. Exploit / Fraud / Economic Safety
│
├── 20. Testing / Simulation
│
├── 21. Historical SWG Relationship
│
├── 22. Prototype Relationship
│
├── 23. Future System Specifications
│
├── 24. Status Register
│
├── 25. Assumption Register
│
├── 26. Open Questions
│
└── 27. Glossary
```

This is essentially the **Opus architecture strengthened by Terra governance and Claude1 system-level detail**.

---

# 12. Synthesized Canonical Master GDD

Below is the resulting canonical synthesis.

---

# Tiwakings Craftworld Industries
## TCIndustries - Master Game Design Document

**Version:** 1.1 - Synthesized Canonical Foundation  
**Purpose:** Canonical game-design authority and high-level development blueprint  
**Audience:** AI design/development collaborators and project owner  
**Technical implementation:** Deferred to a separate Technical Design Document

---

## 1. Document Authority

This document is the current **canonical design authority** for TCIndustries.

The four previous GDDs are historical design sources.

Where they differ, this document supersedes them.

The prototype does not supersede this document.

### Status definitions

| Status | Meaning |
|---|---|
| **LOCKED** | Authoritative project decision |
| **PROPOSED** | Recommended design direction awaiting approval |
| **TBD** | Genuine unresolved design decision |
| **PROTOTYPE** | Experimental implementation/evidence, not canonical |
| **DEFERRED** | Recognised system intentionally postponed |

**No PROPOSED, TBD, PROTOTYPE, or DEFERRED material may silently become LOCKED.**

---

# 2. Executive Summary

TCIndustries is a persistent virtual-world sandbox MMORPG centred on **living as a consequential citizen in a player-driven society**.

The player is not required to be a hero.

They may become:

- a farmer;
- prospector;
- miner;
- artisan;
- manufacturer;
- architect;
- doctor;
- entertainer;
- merchant;
- trader;
- transporter;
- researcher;
- combatant;
- shopkeeper;
- city organiser;
- or a hybrid identity.

The central experience is:

> **Become somebody known in a persistent world.**

The game is inspired by the systemic philosophy of the original *Star Wars Galaxies*, particularly its dynamic resources, crafting, player economy, profession interdependence, services and player-created society.

TCIndustries is not Star Wars and must use original setting, names, lore and content.

The central loop is:

> **Discover → acquire → process → create/provide → specialise → establish reputation → trade → participate in society → create consequences.**

The game succeeds when these systems produce a world that feels alive.

---

# 3. Design Constitution

The following principles are **LOCKED**.

## 3.1 Living World

The world exists independently of any individual player's story.

Players participate in the world rather than being its predetermined protagonist.

## 3.2 Citizen Rather Than Chosen Hero

Players must be able to derive meaningful identity and prestige from ordinary social and economic activity.

## 3.3 Player-Driven Economy

Players must be meaningful producers, consumers, traders, service providers and economic actors.

## 3.4 Interdependence

Important high-level specialisations must depend upon other players.

Interdependence should create **opportunity**, not artificial inconvenience.

## 3.5 Player-Created Society

Players should be able to create:

- businesses;
- organisations;
- communities;
- commercial districts;
- settlements;
- cities;
- reputations;
- supply networks;
- institutions.

## 3.6 Emergence

The game should create conditions from which significant events emerge rather than scripting every important experience.

---

# 4. Anti-Goals

TCIndustries must not become:

- a conventional theme-park MMO;
- a rigid class MMO;
- a combat-only MMO;
- a quest-driven RPG;
- a single-player crafting game with multiplayer attached;
- an infinite personal-factory automation game;
- a Factorio/Satisfactory-style industrial optimisation game;
- a pure market spreadsheet simulator;
- an NPC-dominated economy;
- a game where veteran players can efficiently perform every important role themselves.

These are **design constraints**, not merely aesthetic preferences.

---

# 5. Core Player Experiences

The game should enable experiences such as:

### The Gold Rush

A player discovers an unusually valuable temporary resource.

The discovery causes:

> prospecting → extraction → transport → market activity → manufacturing → exceptional products → competition → regional change.

A resource discovery should therefore be an **event**, not merely an inventory update.

### Becoming Known

A player becomes known for:

- exceptional craftsmanship;
- reliable medical service;
- excellent resources;
- good prices;
- dependable transportation;
- innovative products.

### Genuine Dependency

Players should be able to say:

> "People need what I do."

### Ownership

Players should be able to say:

> "This is my shop."

> "I made this."

> "I discovered this."

> "This is our city."

---

# 6. Player Loops

The game should support multiple interconnected loops.

| Loop | Activities | Result |
|---|---|---|
| Discovery | Survey, explore | Information |
| Extraction | Harvest, mine, gather | Resources |
| Processing | Refine, convert | Materials |
| Crafting | Create, experiment | Products |
| Manufacturing | Repeat production | Scale |
| Retail | Sell, advertise | Commerce |
| Services | Heal, entertain, repair, transport | Social value |
| Combat | Fight, defend, explore | Demand/risk |
| Social | Organisations, cities, events | Society |

No single loop should become the mandatory definition of the game.

---

# 7. Non-Combat Viability

**LOCKED**

A player must be able to pursue a meaningful long-term career without making combat their primary activity.

Combat must not become the universal gateway to:

- wealth;
- progression;
- prestige;
- resources;
- social importance.

---

# 8. World

## 8.1 Persistent Geography

**PROPOSED**

The world should contain geographically differentiated regions.

Regional differences may include:

- resource availability;
- resource quality;
- terrain;
- climate;
- hazards;
- transport access;
- settlement potential;
- population;
- local demand.

## 8.2 Established Settlements

**PROPOSED**

Established settlements provide baseline infrastructure and services.

They should support the economy without replacing exceptional player goods and services.

## 8.3 Player Settlements

**PROPOSED**

Players should be able to create settlements and eventually cities containing:

- housing;
- shops;
- workshops;
- factories;
- civic structures;
- services;
- organisations.

Exact city formation and governance remain **TBD**.

## 8.4 Wilderness

**PROPOSED**

Wilderness should provide:

- exploration;
- resource discovery;
- harvesting;
- construction opportunities;
- hazards;
- points of interest.

It should not become primarily a quest-delivery environment.

---

# 9. Transportation

**PROPOSED**

Transportation should preserve meaningful geography.

Possible layers include:

- local movement;
- personal transport;
- public transport;
- player-operated transport;
- freight;
- constrained fast travel.

Exact implementation is **TBD**.

Fast travel must not completely eliminate the economic relevance of distance.

---

# 10. Character Identity

**LOCKED**

Player identity emerges from:

- skills;
- activities;
- possessions;
- businesses;
- organisations;
- reputation;
- relationships;
- products;
- services;
- locations.

Character identity must persist beyond individual transactions.

---

# 11. Skills and Professions

## 11.1 Skill-Based Model

**LOCKED**

TCIndustries uses flexible skills rather than rigid permanent classes.

Players may:

- learn;
- specialise;
- combine disciplines;
- change direction;
- respecialise.

## 11.2 Specialisation

**PROPOSED**

A meaningful capability limit should prevent one character from becoming simultaneously the best:

- crafter;
- doctor;
- merchant;
- architect;
- prospector;
- combatant;
- manufacturer;
- etc.

The exact specialisation budget is **TBD**.

## 11.3 Profession Families

**PROPOSED**

Initial families include:

- Resource;
- Craft;
- Manufacturing;
- Medical;
- Entertainment;
- Commerce;
- Transport;
- Research;
- Combat;
- Construction.

These are categories, not immutable classes.

---

# 12. Resources

## 12.1 Resource Philosophy

**LOCKED**

Resources are central economic entities rather than generic currencies.

They should have meaningful:

- type;
- location;
- quality;
- attributes;
- scarcity;
- availability;
- lifecycle.

## 12.2 Dynamic Resource Pools

**PROPOSED**

Resources should exist as temporary, geographically distributed resource occurrences.

Exceptional resources should be capable of generating economic events.

## 12.3 Resource Attributes

**PROPOSED**

Different resource occurrences should possess different measurable properties.

Crafting systems should value those properties differently.

Exact attribute definitions are **TBD**.

## 12.4 Discovery

**PROPOSED**

Surveying/prospecting should itself be a meaningful specialisation.

Information about resources may therefore become an economic product.

## 12.5 Extraction

**PROPOSED**

Extraction may support:

- manual harvesting;
- installed extraction equipment.

Exact throughput and infrastructure rules are **TBD**.

## 12.6 Processing

**PROPOSED**

Processing/refining should be a meaningful economic activity rather than necessarily being an automatic step.

This allows:

> harvester → refiner → crafter

to become a genuine social chain.

---

# 13. Crafting

## 13.1 Crafting Philosophy

**LOCKED**

Crafting must convert resource characteristics into meaningful product differentiation.

## 13.2 Schematics

**PROPOSED**

Schematics should define:

- permitted resources;
- required attributes;
- quantities;
- processing requirements;
- skill requirements;
- outputs.

## 13.3 Resource Dependence

**LOCKED**

The identity and quality of input resources must affect outputs.

Two players making the same nominal item should be capable of producing materially different products.

## 13.4 Experimentation

**PROPOSED**

Experimentation should allow skilled players to make meaningful trade-offs.

Exact mechanics are **TBD**.

## 13.5 Provenance

**PROPOSED/LOCKED PRINCIPLE**

Crafted goods should preserve meaningful association with their creator.

Exact provenance depth is **TBD**.

---

# 14. Manufacturing and Automation

## 14.1 Manufacturing Purpose

**LOCKED**

Manufacturing provides production scale.

It must not make individual players irrelevant.

## 14.2 Automation

**LOCKED**

Automation is permitted.

Automation must reinforce rather than destroy the player economy.

## 14.3 Scale Constraints

**PROPOSED**

Scale should encounter meaningful constraints involving combinations of:

- capital;
- land;
- maintenance;
- logistics;
- resources;
- labour;
- market demand;
- skill;
- social organisation.

## 14.4 Artisan versus Factory

**PROPOSED DESIGN PRINCIPLE**

Factories should be good at **repeatable scale**.

Specialist artisans should retain meaningful advantages in **exceptional/custom/high-value production**.

The exact boundary requires simulation.

---

# 15. Items

Items should potentially retain:

- creator;
- resource inputs;
- quality;
- modifications;
- ownership history;
- production history.

Exact provenance architecture is **TBD**.

## Degradation

**LOCKED PRINCIPLE**

Items should degrade, be consumed, require maintenance, or otherwise leave the economy.

This creates continuing demand for:

- resources;
- repair;
- replacement;
- crafting.

Exact degradation rules are **TBD**.

---

# 16. Economy

## 16.1 Player-Driven Economy

**LOCKED**

Players are major participants in production, trade and consumption.

## 16.2 Supply and Demand

**PROPOSED**

Prices should respond to combinations of:

- supply;
- demand;
- scarcity;
- quality;
- competition;
- location;
- transportation;
- player behaviour.

## 16.3 Economic Geography

**PROPOSED**

Location should matter economically without making travel unnecessarily tedious.

## 16.4 Currency

**PROPOSED/TBD**

A persistent economy requires explicit treatment of:

- currency sources;
- currency sinks;
- transaction flows;
- inflation;
- concentration.

Exact model is **TBD**.

---

# 17. Retail and Commerce

## 17.1 Player Shops

**LOCKED**

Player shops should be meaningful commercial identities rather than inventory containers.

## 17.2 Location

**PROPOSED**

Location should affect:

- visibility;
- convenience;
- foot traffic;
- transport;
- regional demand;
- commercial reputation.

## 17.3 Reputation

**PROPOSED**

Reputation should emerge primarily from observable behaviour and history.

Potential signals include:

- provenance;
- sales history;
- reliability;
- product quality;
- service history.

Formal ratings remain **TBD**.

---

# 18. Services

**LOCKED PRINCIPLE**

Non-item services are first-class economic activities.

Potential services include:

- medicine;
- entertainment;
- repair;
- transport;
- construction;
- research;
- appraisal;
- specialist consultation.

Services must create real player interaction without becoming artificial waiting mechanics.

---

# 19. Combat

## 19.1 Role

**LOCKED**

Combat is one important activity within the world.

It should create demand for:

- weapons;
- armour;
- medical services;
- repairs;
- consumables;
- transport;
- specialist equipment.

## 19.2 Detailed Combat

**DEFERRED**

Detailed combat systems require a dedicated specification.

## 19.3 PvP

**TBD**

The project has not yet resolved:

- open PvP;
- consensual PvP;
- territorial conflict;
- criminal systems;
- faction warfare;
- PvE-only alternatives.

This decision has significant dependencies and must be resolved before several downstream systems can be finalised.

---

# 20. Buildings, Ownership and Cities

## 20.1 Ownership

**LOCKED**

Players and organisations should own meaningful persistent assets.

## 20.2 Buildings

**PROPOSED**

Structures may include:

- homes;
- shops;
- workshops;
- factories;
- civic buildings;
- service facilities.

## 20.3 Cities

**PROPOSED**

Player cities should emerge from concentrations of player structures, organisations, services and population.

Exact:

- zoning;
- governance;
- taxation;
- maintenance;
- growth;
- dissolution

remain **TBD**.

---

# 21. Social Organisations

**PROPOSED**

Players should be able to form persistent organisations.

Organisations may coordinate:

- production;
- commerce;
- combat;
- transport;
- cities;
- resource gathering;
- social events.

Organisation permissions and asset management require dedicated design.

---

# 22. Emergent Gameplay

The desired systemic chain is:

```text
Exceptional resource appears
        ↓
Prospector discovers it
        ↓
Information spreads
        ↓
Gatherers respond
        ↓
Transport demand increases
        ↓
Resource enters markets
        ↓
Crafters seek it
        ↓
Exceptional products appear
        ↓
Retailers market them
        ↓
Reputation develops
        ↓
Consumers change behaviour
        ↓
Regional economy changes
```

This is the type of systemic event TCIndustries should produce.

---

# 23. Content Expansion

**LOCKED**

New content should primarily operate through existing generic systems.

For example:

- new resources should use the resource framework;
- new products should use schematics;
- new skills should use the skill architecture;
- new structures should use building rules;
- new regions should use existing world/resource rules.

The preferred expansion model is:

> **Add data and rules to existing systems rather than rewriting the engine for every new content type.**

---

# 24. Balance Philosophy

The central balance question is not:

> "How powerful is this system?"

It is:

> **"Does this system strengthen or weaken the living player society?"**

Every major proposed mechanic should therefore be evaluated against:

1. interdependence;
2. player identity;
3. economic viability;
4. non-combat viability;
5. emergence;
6. automation risk;
7. monopoly risk.

---

# 25. Economic and Exploit Safety

Future specifications must address:

- item duplication;
- currency duplication;
- invalid item creation;
- resource duplication;
- automated market manipulation;
- wash trading;
- vendor exploitation;
- abandoned property;
- land monopolisation;
- resource monopolisation;
- multi-account exploitation;
- organisation asset theft;
- contract fraud;
- governance griefing.

These are design dependencies, not merely technical problems.

---

# 26. Design Invariants

The following should become testing requirements.

1. Items cannot be created without valid sources.
2. Crafted provenance must identify the creator or permitted producer.
3. Schematics cannot accept invalid resources.
4. Resource identity remains stable during its lifecycle.
5. Extraction cannot exceed permitted resource availability.
6. Manufacturing cannot produce without valid inputs/capacity/permissions.
7. Ownership transfer must be auditable.
8. Currency must obey explicit source/sink rules.
9. Skill caps must enforce defined specialisation restrictions.
10. Building placement must obey ownership/land/zoning rules.
11. Repair/degradation cannot create material value.
12. No standard system may create infinite risk-free currency or items.

These are directly aligned with the strongest testing material in the Terra/Opus documents. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

# 27. Promotion Criteria

A proposed system should not become **LOCKED** until it has:

1. a defined player goal;
2. defined inputs;
3. defined state;
4. defined outputs;
5. failure cases;
6. dependencies;
7. exploit analysis;
8. player-visible consequences;
9. compatibility with the design pillars;
10. simulation/prototype evidence where appropriate.

This is one of the strongest principles inherited from the later GDDs. `TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

---

# 28. Historical SWG Relationship

TCIndustries may use historical SWG as a source of design inspiration.

Documentation must distinguish:

1. **Historical SWG behaviour**
2. **TCIndustries adaptation**
3. **New TCIndustries invention**

Historical similarity does not make a mechanic canonical.

This distinction is explicitly established in the Opus architecture. `TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

---

# 29. Prototype Relationship

The Seed TypeScript/Vite/React prototype is:

**PROTOTYPE**

It is evidence and experimentation, not authority.

The development hierarchy is:

```text
MASTER GDD
     ↓
SYSTEM SPECIFICATIONS
     ↓
PROTOTYPE
     ↓
TESTING / SIMULATION
     ↓
GOLDEN MASTER
     ↓
PRODUCTION IMPLEMENTATION
```

Prototype behaviour must never become canonical merely because code already implements it.

The Opus document is particularly strong on this governance point. `TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

---

# 30. Systems Requiring Dedicated Design

Priority order should be:

1. **Resource spawning and surveying**
2. **Resource quality/attribute mathematics**
3. **Skill budget and specialisation**
4. **Crafting quality and experimentation**
5. **Item durability and repair**
6. **Extraction and automation limits**
7. **Manufacturing**
8. **Currency sources and sinks**
9. **Retail and discovery**
10. **Buildings and land**
11. **Services**
12. **Cities and governance**
13. **Organisations**
14. **Transportation**
15. **Combat**
16. **Reputation**
17. **PvP**
18. **Advanced world simulation**

This ordering reflects the dependency relationships identified across the documents rather than simply following chapter order.

---

# 31. Open Questions

The most important unresolved questions include:

| Area | Question |
|---|---|
| Setting | What is the original setting/genre? |
| Resources | What exact attributes exist? |
| Resources | How do resource pools spawn and expire? |
| Surveying | How much information does a surveyor receive? |
| Skills | What is the specialisation-cap model? |
| Skills | How are skills acquired and improved? |
| Crafting | What is the quality formula? |
| Crafting | What exactly does experimentation do? |
| Manufacturing | Where is the artisan/factory boundary? |
| Economy | What are the currency faucets/sinks? |
| Retail | How are shops discovered? |
| Reputation | What measurable signals exist? |
| Cities | How are cities founded/governed? |
| Organisations | What powers and liabilities do organisations have? |
| Transport | What forms of transportation exist? |
| Combat | What is the combat system? |
| PvP | What PvP model exists? |
| Death | What happens when characters die or become injured? |
| Services | How are services valuable without forced downtime? |
| Accounts | What is the alt/multi-account policy? |

These questions should remain visible rather than being buried inside prose. The earlier documents, especially Terra and Claude1, are particularly strong in this regard. 

---

# 32. Assumption Policy

Any designer introducing an assumption must identify:

1. what is being assumed;
2. why;
3. what source supports it;
4. what happens if the assumption proves wrong.

Claude1's explicit assumption-register approach should be retained. Its register demonstrates the value of making assumptions visible rather than allowing them to become accidental canon. `TCIndustries_Master_GDD-claude1.md`

---

# 33. Recommended Design Workflow

The recommended future workflow is:

```text
MASTER GDD
     ↓
Identify unresolved dependency
     ↓
Create dedicated System Specification
     ↓
Define alternatives where necessary
     ↓
Prototype / simulate
     ↓
Evaluate against design invariants
     ↓
User/project-owner decision
     ↓
Promote selected rule to LOCKED
     ↓
Update Master GDD
     ↓
Update Golden Master
```

This prevents the Master GDD from becoming a giant document containing thousands of prematurely locked decisions.

---

# 34. Final Assessment

## Which GDD is best?

### **Claude Opus 4.6 v1.0**

It is the best existing **Master GDD** because it achieves the best balance between:

- vision;
- architecture;
- system coverage;
- design governance;
- economic thinking;
- world structure;
- testing;
- exploit awareness;
- prototype governance;
- future design sequencing.

Its strongest feature is that it behaves like a **design system**, rather than merely a description of a game.

---

## Which has the best mechanical material?

### **Claude1**

Its resource, crafting, retail and system-interaction thinking is exceptionally useful.

Its weakness is not poor design; it is that some of its good ideas are still **designer proposals rather than project decisions**.

That material should be mined selectively.

---

## Which has the best governance?

### **Terra-medium**

It is particularly strong at defining:

- what the GDD is;
- what is locked;
- what is proposed;
- what remains unresolved;
- what must be tested;
- what invariants should hold.

Its testing requirements and design-invariant approach should definitely survive. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

## Which is the cleanest foundation?

### **ChatGPT1**

It has the least architectural baggage and communicates the fundamental game very clearly.

It should effectively become the **philosophical baseline** underneath the more developed architecture.

---

# 35. Final recommendation

I **do not recommend simply renaming Claude Opus v1.0 to v1.1 and continuing from it.**

I recommend treating the four documents as follows:

```text
                    ┌──────────────────────┐
                    │    CHATGPT1          │
                    │ Vision / foundation  │
                    └──────────┬───────────┘
                               │
                    ┌──────────▼───────────┐
                    │ TERRA-MEDIUM         │
                    │ Governance / tests   │
                    └──────────┬───────────┘
                               │
             ┌─────────────────▼─────────────────┐
             │        CLAUDE OPUS v1.0           │
             │      PRIMARY ARCHITECTURE         │
             └─────────────────┬─────────────────┘
                               │
                    ┌──────────▼───────────┐
                    │ CLAUDE1             │
                    │ Mechanical proposals│
                    └──────────┬───────────┘
                               │
                               ▼
                  ┌─────────────────────────┐
                  │ SYNTHESIZED MASTER GDD  │
                  │                         │
                  │ Canonical authority     │
                  │ + design blueprint      │
                  └─────────────────────────┘
```

The result should therefore **not be considered "Claude's GDD."**

It should be considered a new **TCIndustries canonical GDD**, whose architecture is principally inherited from Opus but whose governance, testing and mechanical content have been deliberately reconciled from all four sources.

### Most importantly

I would preserve the following rule as a fundamental project principle:

> **A detailed design is not necessarily a decided design.**

That is the central lesson from comparing the four documents.

Claude1 demonstrates how to design systems in depth.  
Terra demonstrates how to control design authority.  
Opus demonstrates how to integrate the two.  
ChatGPT1 demonstrates how to keep the underlying game vision clear.

**That combination is stronger than any one of the four.**

And I would regard the synthesized document above as the **structural foundation**, not yet the final fully locked game design. The next step should be to resolve the highest-dependency systems-especially **Resources → Skills/Specialisation → Crafting → Manufacturing → Economy**-through dedicated specifications and simulation before their proposed mechanics are promoted to canonical rules. This is also consistent with the existing documents' recommendation to begin with the resource system and then move into skills and crafting. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

## Prompt:
20/08/2026, 22:37:51

based on your analysis, combine, synthesize, and output GDD in markdown.

## Response:
20/08/2026, 22:37:52

> **Remembering**
>
>
> ****
>
> 
>
> **Reviewed and preserved the canonical TCIndustries GDD synthesis**
>
> 
>
> Worked for 12s

# Tiwakings Craftworld Industries
# TCIndustries - Master Game Design Document

**Version:** 1.1 - Synthesized Canonical Foundation  
**Date:** 20 August 2026  
**Document Type:** Master Game Design Document  
**Purpose:** Canonical game-design authority and high-level development blueprint  
**Audience:** AI design/development collaborators and project owner  
**Technical Scope:** Game design only; technical implementation belongs in a separate TDD

---

> **Document Authority**
>
> This document synthesizes the four preceding TCIndustries GDDs:
>
> 1. *TCIndustries Master Game Design Document v0.1 - GPT-5.6 Terra-medium*
> 2. *Tiwakings Craftworld Industries - Master Game Design Document v0.1 - ChatGPT1*
> 3. *TCIndustries Master GDD v1.0 - Claude Opus 4.6 Thinking*
> 4. *TCIndustries Master GDD - Claude1*
>
> The synthesis uses the **Claude Opus document as the primary architectural spine**, the **Terra document as the principal governance/testing framework**, **Claude1 as the source of selected deeper mechanical proposals**, and **ChatGPT1 as a foundational reference for the project's simplest expression of its vision and anti-goals**.
>
> The source documents remain historical design records. This document is the current canonical synthesis.

---

# Table of Contents

1. [Executive Summary](#1-executive-summary)
2. [Design Authority and Status System](#2-design-authority-and-status-system)
3. [Vision](#3-vision)
4. [Design Pillars](#4-design-pillars)
5. [Player Experience and Core Loops](#5-player-experience-and-core-loops)
6. [Anti-Goals](#6-anti-goals)
7. [World Philosophy and Structure](#7-world-philosophy-and-structure)
8. [World Simulation](#8-world-simulation)
9. [Character, Identity, and Ownership](#9-character-identity-and-ownership)
10. [Skills, Professions, and Progression](#10-skills-professions-and-progression)
11. [Resource System](#11-resource-system)
12. [Crafting System](#12-crafting-system)
13. [Manufacturing and Automation](#13-manufacturing-and-automation)
14. [Items](#14-items)
15. [Economy](#15-economy)
16. [Retail and Commerce](#16-retail-and-commerce)
17. [Services](#17-services)
18. [Combat](#18-combat)
19. [Buildings, Housing, and Cities](#19-buildings-housing-and-cities)
20. [Social Systems and Organizations](#20-social-systems-and-organizations)
21. [Progression](#21-progression)
22. [Emergent Gameplay](#22-emergent-gameplay)
23. [Content Expansion Framework](#23-content-expansion-framework)
24. [Balance Philosophy](#24-balance-philosophy)
25. [Exploit, Fraud, and Economic Stability](#25-exploit-fraud-and-economic-stability)
26. [Testing Requirements and Design Invariants](#26-testing-requirements-and-design-invariants)
27. [Historical SWG Relationship](#27-historical-swg-relationship)
28. [Prototype Relationship](#28-prototype-relationship)
29. [Systems Requiring Deeper Design](#29-systems-requiring-deeper-design)
30. [Recommended Design Sequence](#30-recommended-design-sequence)
31. [Status Register](#31-status-register)
32. [Assumption Register](#32-assumption-register)
33. [Open-Question Register](#33-open-question-register)
34. [Glossary](#34-glossary)

---

# 1. Executive Summary

## 1.1 Game Concept

TCIndustries is a **persistent sandbox MMORPG and virtual-world simulation** centred on living within a player-driven society.

The player is not necessarily a hero, saviour, chosen one, or central protagonist.

The player may instead become:

- a farmer;
- miner;
- prospector;
- harvester;
- refiner;
- artisan;
- weaponsmith;
- armourer;
- tailor;
- architect;
- manufacturer;
- doctor;
- entertainer;
- merchant;
- trader;
- transporter;
- researcher;
- explorer;
- combatant;
- shopkeeper;
- business owner;
- organisation leader;
- city participant;
- or a hybrid identity created through the interaction of skills, possessions, reputation, relationships and player choices.

The core fantasy is:

> **Live in a persistent world, decide what kind of person you want to become, and become known for it.**

The game's primary content is not a predetermined sequence of quests.

It is the interaction between:

- players;
- resources;
- production;
- commerce;
- services;
- geography;
- ownership;
- reputation;
- organisations;
- cities;
- combat;
- and world systems.

The desired outcome is a world in which players can genuinely say:

> **"I live here. This is what I do. These are my people. This is my business. What I do matters."**

The four source GDDs converge strongly on this vision. 

---

## 1.2 Historical Inspiration

TCIndustries is heavily inspired by the **virtual-world philosophy of the original Star Wars Galaxies**, particularly:

- dynamic resources;
- resource quality;
- surveying;
- player crafting;
- profession interdependence;
- player commerce;
- recognisable crafters;
- services;
- housing;
- player settlements;
- social institutions;
- emergent economic activity.

TCIndustries is **not a reproduction of Star Wars**.

Historical SWG mechanics must always be distinguished from:

1. historical behaviour;
2. TCIndustries adaptations;
3. original TCIndustries inventions.

This distinction is a locked documentation principle. 

---

## 1.3 Central Game Loop

The broad systemic loop is:

```text
DISCOVER
    ↓
ACQUIRE
    ↓
PROCESS
    ↓
CRAFT / MANUFACTURE / PROVIDE SERVICE
    ↓
TRADE
    ↓
BUILD REPUTATION
    ↓
SPECIALISE
    ↓
FORM RELATIONSHIPS
    ↓
BUILD BUSINESSES / ORGANISATIONS / COMMUNITIES
    ↓
CHANGE THE WORLD
    ↓
CREATE NEW OPPORTUNITIES
```

This is not intended to be a mandatory linear progression path.

It is the interconnected economic and social structure from which individual player lives emerge.

---

## 1.4 Design Objective

The game's success should ultimately be measured by whether its systems produce a world that feels **alive**.

Maximum complexity is not the objective.

The objective is:

> **Meaningful player agency within an interconnected virtual society.**

---

# 2. Design Authority and Status System

## 2.1 Authority

This document is the canonical game-design authority.

Where prototype behaviour, historical mechanics, external suggestions, or earlier GDDs disagree with this document, this document takes precedence unless the design is formally revised.

---

## 2.2 Status Definitions

Every substantive design rule must use one of five statuses.

| Status | Meaning |
|---|---|
| **LOCKED** | Explicit project decision. Authoritative unless formally revised. |
| **PROPOSED** | Recommended design direction that has not yet been formally approved. |
| **TBD** | Genuine unresolved design decision. |
| **PROTOTYPE** | Behaviour explored in prototype code or experiments; not canonical. |
| **DEFERRED** | Recognised future system intentionally postponed. |

No status may be silently promoted.

In particular:

> **A detailed proposal is not automatically a decision.**

This is one of the most important governance principles carried forward from the Terra and Claude documents. 

---

## 2.3 Design Specification Hierarchy

Future design work should follow:

```text
MASTER GDD
     ↓
DEDICATED SYSTEM SPECIFICATION
     ↓
ALTERNATIVES / DESIGN OPTIONS
     ↓
PROTOTYPE / SIMULATION
     ↓
TESTING
     ↓
PROJECT-OWNER DECISION
     ↓
LOCKED RULE
     ↓
UPDATED MASTER GDD
```

The Master GDD should not become a dumping ground for premature numerical decisions.

---

# 3. Vision

## VIS-001 - Living Virtual World

**Status: LOCKED**

TCIndustries is a persistent multiplayer virtual world in which players create their own:

- economic;
- social;
- occupational;
- reputational;
- commercial;
- organisational

identities.

The world should feel inhabited not merely by NPCs, but by **player-created institutions**:

- shops;
- workshops;
- supply networks;
- social venues;
- cities;
- organisations;
- specialist communities;
- commercial districts;
- brands;
- reputations.

The world should continue existing independently of any particular player's story. 

---

## VIS-002 - Citizen Rather Than Chosen Hero

**Status: LOCKED**

Players should feel like citizens of a living society rather than protagonists of a predetermined story.

A player should be able to derive prestige from:

- discovering rare resources;
- producing respected goods;
- operating a successful shop;
- becoming a famous doctor;
- becoming a respected entertainer;
- coordinating a city;
- operating a trade network;
- designing desirable buildings;
- becoming an expert in a narrow specialisation;
- leading an organisation;
- becoming a respected combatant.

Combat must not be the only route to status.

---

## VIS-003 - Player-Driven Economy

**Status: LOCKED**

Players should be major participants in production and distribution.

The economy should be materially connected to:

- resources;
- labour;
- transportation;
- crafting;
- manufacturing;
- consumption;
- degradation;
- services;
- geography;
- scarcity;
- reputation.

The economy must generate meaningful opportunities for:

- production;
- trade;
- specialisation;
- commerce;
- entrepreneurship;
- services;
- information.

`TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

---

## VIS-004 - Interdependence

**Status: LOCKED**

Important professions should depend upon one another.

No major profession category should be fully self-sufficient at high effectiveness.

Interdependence should emerge from meaningful differences in:

- skill;
- resources;
- knowledge;
- capacity;
- geography;
- logistics;
- services;
- reputation;
- social organisation.

It must create **opportunity**, not artificial inconvenience.

---

## VIS-005 - Emergence

**Status: LOCKED**

The game should create conditions from which significant events emerge rather than scripting every important player experience.

The desired player reaction is:

> **"This world changed because of what players did."**

---

# 4. Design Pillars

## PIL-001 - Discovery and the Gold Rush

**Status: LOCKED**

Resources are central to the living-world experience.

Resource availability should be capable of being:

- temporary;
- geographically distributed;
- variable in quality;
- economically significant.

Exceptional discoveries should create events involving:

- prospecting;
- extraction;
- transportation;
- speculation;
- trade;
- manufacturing;
- competition;
- organisation;
- regional economic change.

A resource discovery should be an **event**, not merely an inventory update. `TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

### TCIndustries Adaptation

**Status: PROPOSED**

TCIndustries should use dynamic resource occurrences inspired by historical SWG resource behaviour, while adding explicit economic observability and anti-monopoly safeguards.

---

## PIL-002 - Identity and Reputation

**Status: LOCKED**

Players should become recognisable for what they do.

Examples:

> "That player makes the best armour."

> "That shop always has excellent medical supplies."

> "That prospector finds exceptional resources."

Reputation should emerge primarily from observable player behaviour and economic history rather than from one abstract reputation score. `TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

---

## PIL-003 - Interdependence

**Status: LOCKED**

The economy must create meaningful reasons for specialists to interact.

Examples include:

```text
Prospector
   ↓
Harvester
   ↓
Refiner
   ↓
Component Maker
   ↓
Crafter
   ↓
Merchant
   ↓
Customer
```

Other chains may involve:

```text
Combatant
   ↓
Equipment Demand
   ↓
Crafter
   ↓
Repair Specialist
   ↓
Medical Specialist
```

The exact chains are system-dependent.

---

## PIL-004 - Player-Created Society

**Status: LOCKED**

Players must be able to create persistent social and commercial structures.

Potential manifestations include:

- businesses;
- organisations;
- cities;
- trade networks;
- specialist communities;
- social venues;
- commercial districts;
- institutions.

Ownership is intended to have emotional significance:

> "This is my shop."

> "I made this."

> "I discovered this."

> "This is my business."

> "This is our city."

`TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

---

# 5. Player Experience and Core Loops

## LOOP-001 - Multiple Interconnected Loops

**Status: PROPOSED**

TCIndustries should support multiple legitimate player loops.

| Loop | Activities | Result |
|---|---|---|
| Discovery | Survey, explore, identify resources | Information and opportunity |
| Extraction | Gather, harvest, operate equipment | Resource supply |
| Processing | Refine, transform | Intermediate goods |
| Crafting | Create, experiment, specialise | Differentiated products |
| Manufacturing | Repeat production | Scale |
| Retail | Stock, price, advertise, sell | Commerce |
| Services | Heal, entertain, repair, transport | Social value |
| Combat | Fight, defend, explore | Demand and risk |
| Governance | Organisations, cities, events | Society |

`TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

## LOOP-002 - Non-Combat Viability

**Status: LOCKED**

A player must be able to pursue a meaningful long-term career without making combat their primary activity.

Combat may create:

- demand;
- risk;
- opportunity;
- materials;
- access;
- economic activity.

It must not be the universal gateway to:

- wealth;
- progression;
- prestige;
- social importance.

---

## LOOP-003 - Progression Expands Options

**Status: LOCKED**

Progression should increase:

- capability;
- specialisation;
- reputation;
- access;
- opportunities;
- complexity.

It should not force every player through one linear power ladder.

---

# 6. Anti-Goals

The following are **LOCKED** design constraints.

## AG-001 - Do Not Become a Spreadsheet Simulator

Economic depth is desirable.

Economic abstraction that makes the player experience primarily about optimising spreadsheets is not.

---

## AG-002 - Do Not Become a Single-Player Industrial Tycoon

Players must not simply accumulate unlimited factories and automate the entire economy personally.

Industrial scale must encounter meaningful constraints.

---

## AG-003 - Do Not Make Everyone Self-Sufficient

A single character should not efficiently perform every economically important role.

Universal self-sufficiency would destroy the player economy.

---

## AG-004 - Do Not Force Rigid Class Identity

Players should not be permanently locked into one class.

Skill-based identity should remain flexible.

---

## AG-005 - Do Not Make Content Completely Script-Dependent

The world should generate meaningful experiences through systems and player interaction.

---

## AG-006 - Do Not Create an NPC-Dominated Economy

NPCs may provide baseline services and economic support.

They must not replace meaningful player production and commerce.

---

## AG-007 - Do Not Become Combat-Only

Combat is a legitimate profession and economic driver, not the definition of the entire game.

These anti-goals are strongly convergent across the source documents. 

---

# 7. World Philosophy and Structure

## WRLD-001 - Persistent Shared World

**Status: LOCKED**

TCIndustries is a persistent shared world.

Player actions should have durable consequences where practical, including:

- structures;
- businesses;
- inventories;
- market conditions;
- resource availability;
- city development;
- organisational ownership;
- reputation;
- product provenance.

`TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

## WRLD-002 - Regionally Differentiated Geography

**Status: PROPOSED**

Regions should differ through combinations of:

- terrain;
- climate;
- resources;
- resource quality;
- settlement potential;
- dangers;
- transport access;
- local demand;
- environmental conditions.

Regions should not simply be level bands.

Their purpose is to generate meaningful geography and economic differentiation.

---

## WRLD-003 - World Structure

**Status: PROPOSED**

The world should contain:

### Established/NPC Settlements

Stable locations providing baseline services and economic infrastructure.

### Player Settlements and Cities

Player-created institutions with:

- buildings;
- governance;
- commerce;
- community identity;
- services.

### Wilderness

Areas supporting:

- exploration;
- resource discovery;
- harvesting;
- construction;
- combat;
- travel.

### Economic Corridors

Transportation routes connecting:

- resource regions;
- settlements;
- markets;
- industrial centres.

The exact city-formation model is **TBD**.

---

# 8. World Simulation

## WSIM-001 - Meaningful Simulation

**Status: PROPOSED**

TCIndustries should simulate systems that produce meaningful consequences.

It should **not** simulate everything merely because simulation is possible.

Potential simulation domains include:

- resource changes;
- production;
- consumption;
- population;
- markets;
- transportation;
- environmental conditions;
- settlement activity;
- NPC activity;
- player organisations.

---

## WSIM-002 - Persistent Time

**Status: PROPOSED**

A persistent world clock should support systems such as:

- resource lifecycles;
- production;
- travel;
- environmental change;
- events.

Exact time scale is **TBD**.

---

## WSIM-003 - Weather and Environment

**Status: DEFERRED / PROPOSED HOOK**

Weather and environmental systems may be added where they create meaningful decisions.

Examples could include:

- resource modifiers;
- travel conditions;
- ecological changes;
- temporary hazards.

Weather should not exist solely as visual decoration.

---

## WRLD-005 - Transportation

**Status: PROPOSED**

Transportation should preserve meaningful geography without making routine participation excessively tedious.

Potential levels include:

- local movement;
- personal transport;
- public transport;
- organisation-operated transport;
- freight;
- constrained fast travel.

Exact transportation rules are **TBD**.

Fast travel must not completely eliminate:

- regional markets;
- transport services;
- logistics;
- geographic advantage.

The source documents consistently identify transportation as an important unresolved economic dependency. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

# 9. Character, Identity, and Ownership

## PLR-001 - Persistent Identity

**Status: LOCKED**

A player's identity should emerge from:

- skills;
- activities;
- possessions;
- businesses;
- reputation;
- relationships;
- organisations;
- products;
- services;
- location.

---

## PLR-002 - Ownership

**Status: LOCKED**

Players and organisations should be able to own meaningful persistent assets.

Potential assets include:

- items;
- buildings;
- shops;
- factories;
- businesses;
- land rights;
- organisational property.

Ownership must be auditable and resistant to fraud.

---

## PLR-003 - Identity Persistence

**Status: PROPOSED**

Respecialisation or changes in occupation should not erase:

- historical reputation;
- ownership;
- item provenance;
- business history;
- social relationships.

A player should be able to change what they do without becoming an entirely new person.

---

# 10. Skills, Professions, and Progression

## PROF-001 - Flexible Skill-Based Professions

**Status: LOCKED**

TCIndustries uses flexible skill-based professions rather than rigid immutable classes.

A profession is an **emergent identity** created through:

- acquired skills;
- specialisation;
- equipment;
- reputation;
- behaviour;
- social relationships.

Players should be able to:

- learn;
- improve;
- specialise;
- combine disciplines;
- abandon skills;
- pursue new specialisations.

`TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

---

## PROF-002 - Skill Domains

**Status: PROPOSED**

Initial conceptual domains include:

- surveying/prospecting;
- extraction/harvesting;
- processing/refining;
- crafting;
- manufacturing;
- architecture/construction;
- commerce/retail;
- medical;
- entertainment;
- combat;
- survival/exploration;
- research/analysis;
- transport/logistics.

These are conceptual categories, not necessarily final skill trees.

---

## PROF-003 - Meaningful Specialisation

**Status: PROPOSED**

Players should be able to develop broad competence but should require meaningful investment to achieve high-tier expertise.

A master weaponsmith should not automatically be:

- a master doctor;
- architect;
- merchant;
- prospector;
- entertainer;
- combatant.

Specialisation should create markets for experts.

---

## PROF-004 - Respecialisation

**Status: LOCKED**

Players must be able to respecialise.

Permanent character traps are inconsistent with sandbox identity.

Respecialisation should **not** erase:

- character history;
- reputation;
- ownership;
- item provenance;
- business identity.

The precise respecialisation cost/mechanism is **TBD**. `TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

---

## PROF-005 - Capability Cap

**Status: PROPOSED**

A capability cap or specialisation budget should prevent one character from mastering every economically important role simultaneously.

The model must preserve:

- experimentation;
- hybrid characters;
- casual participation;
- autonomy;
- respecialisation.

The exact model is **TBD**.

---

# 11. Resource System

## RES-001 - Resource Design Goal

**Status: LOCKED**

Resources exist to create:

- discovery;
- scarcity;
- regional trade;
- product differentiation;
- temporary opportunity;
- economic events;
- specialisation.

Resources must not function solely as generic crafting currencies.

A resource should not merely be:

> `Iron × 1000`

It should possess meaningful identity.

`TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

---

## RES-002 - Resource Instances

**Status: PROPOSED**

Resources should exist as identifiable instances or pools with characteristics such as:

- resource class;
- location;
- attributes;
- quality;
- availability;
- lifecycle.

Example:

> **Kavren Ferritic Alloy - North Ember Plains - Cycle 184**

The purpose is to make resources economically meaningful objects rather than anonymous quantities. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

## RES-003 - Resource Categories and Attributes

**Status: PROPOSED**

Resource categories should be extensible.

Possible attributes include:

- purity;
- structural strength;
- conductivity;
- malleability;
- heat resistance;
- density;
- elasticity;
- chemical stability;
- energy yield;
- biological vitality;
- aesthetic quality;
- corrosion resistance.

This is an example taxonomy rather than a locked launch list.

The exact attribute library is **TBD**. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

## RES-004 - Resource Quality

**Status: LOCKED PRINCIPLE**

Resource quality must have gameplay consequences.

High-quality resources should potentially produce:

- better products;
- greater market value;
- specialised opportunities;
- reputation;
- competitive demand.

The exact quality formula is **TBD**.

---

## RES-005 - Dynamic Resource Occurrences

**Status: PROPOSED**

Resource availability should change over time and space.

A resource occurrence should have a lifecycle rather than being an indefinitely available static node.

Two separate concepts should be retained:

1. **Spawn lifecycle** - how long the resource occurrence exists.
2. **Local extraction state** - whether a specific location becomes exhausted, inefficient, congested or otherwise less productive.

Exact algorithms require simulation. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

## RES-006 - Surveying and Prospecting

**Status: PROPOSED**

Resource discovery should itself be a profession.

Surveying may provide progressively better information according to skill.

Lower capability might reveal:

- broad location;
- approximate resource type;
- rough quality.

Higher capability may reveal:

- more precise location;
- resource characteristics;
- quality;
- confidence.

Information may become an economic product.

Players may:

- sell information;
- sell surveying services;
- share discoveries;
- exploit discoveries themselves.

The system must prevent permanent monopolisation of useful information.

`TCIndustries_Master_GDD-claude1.md`

---

## RES-007 - Extraction

**Status: PROPOSED**

Extraction may support multiple scales:

### Manual

Hands-on gathering suitable for:

- flora;
- fauna;
- low-scale extraction;
- new players;
- opportunistic participation.

### Installed Equipment

Player-owned extraction equipment may:

- operate over time;
- require land;
- require maintenance;
- have throughput limits;
- require logistics.

Automation must not allow unlimited unattended individual scale.

Potential constraints include:

- installation cost;
- maintenance;
- fuel/power;
- land rights;
- output limits;
- competition;
- setup time;
- transport;
- vulnerability.

`TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

## RES-008 - Processing

**Status: PROPOSED**

Raw resources may be transformed into:

- refined resources;
- intermediate materials;
- components.

Processing should create genuine economic roles.

A processing stage is justified when it introduces at least one of:

- skill specialisation;
- location choice;
- quality trade-off;
- transformation;
- market opportunity;
- logistics.

It should not exist merely to create tedious clicks.

`TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

## RES-009 - Resource Markets

**Status: PROPOSED**

Resources and processed materials should be tradeable goods.

Relevant resource information should remain available to buyers so they can evaluate quality rather than purchasing anonymous commodities.

This permits reputations such as:

> "That supplier always has excellent ore."

---

# 12. Crafting System

## CRFT-001 - Crafting Philosophy

**Status: LOCKED**

Crafting must produce:

- meaningful product variation;
- player identity;
- market differentiation.

Crafting is not merely a recipe-unlock treadmill.

Product desirability may depend upon:

- crafter skill;
- specialisation;
- schematic;
- resource attributes;
- component quality;
- experimentation;
- customisation;
- provenance;
- reliability;
- reputation.

`TCIndustries_Master_GDD_v1.0-claude-opus-4-6-thiking.md`

---

## CRFT-002 - Resource Choice Affects Output

**Status: LOCKED**

The identity and quality of input resources must influence output.

Therefore:

> **Two players making the same nominal item should be capable of producing materially different products.**

---

## CRFT-003 - Schematics

**Status: PROPOSED**

A schematic may define:

- required skill;
- input categories;
- component slots;
- minimum resource attributes;
- process;
- experimentation options;
- output type;
- baseline statistics;
- quality limits;
- customisation;
- durability;
- manufacturing eligibility.

Schematics should reference resource **categories and attributes** rather than hard-coded individual resources wherever possible.

This supports the content-expansion architecture.

---

## CRFT-004 - Crafting Process

**Status: PROPOSED**

The conceptual process is:

1. Select schematic.
2. Select inputs/components.
3. Validate materials.
4. Assemble.
5. Experiment/refine where permitted.
6. Calculate outcome.
7. Create item.
8. Record provenance.

The exact player interaction model is **TBD**.

---

## CRFT-005 - Experimentation

**Status: PROPOSED**

Experimentation should allow skilled crafters to influence output properties.

Potential dimensions include:

- performance;
- durability;
- efficiency;
- weight;
- reliability;
- appearance;
- other schematic-defined attributes.

Experimentation should involve meaningful choices and trade-offs.

It should not simply provide:

> "Press button → maximise everything."

Historical SWG provides a useful inspiration for a skill-based experimentation model, but the precise TCIndustries mechanism remains a design decision. 

---

## CRFT-006 - Experimentation Risk

**Status: PROPOSED**

Experimentation may involve risk.

Potential consequences include:

- reduced output quality;
- failed experimentation;
- resource loss;
- unusable output;
- trade-offs between competing attributes.

Exact failure mechanics are **TBD**.

---

## CRFT-007 - Product Quality

**Status: PROPOSED**

Quality should result from some combination of:

- schematic baseline;
- crafter skill;
- resource quality;
- resource attributes;
- components;
- experimentation;
- production method.

The exact quality formula is **TBD**.

---

## CRFT-008 - Provenance

**Status: LOCKED**

Crafted items must retain meaningful provenance.

At minimum, provenance should support:

- original creator;
- producing organisation where applicable;
- creation time;
- item quality;
- relevant material identity;
- relevant production history.

The exact visibility, privacy, resale and modification rules remain **TBD**.

The important design principle is:

> **Trading an item must not erase who made it.**

This is central to crafter reputation.

---

## CRFT-009 - Signature Products

**Status: LOCKED**

The system should permit crafters to become known for exceptional products.

A player's output should potentially become identifiable as:

> "their work."

This is a major bridge between crafting mechanics and player identity.

---

# 13. Manufacturing and Automation

## MFG-001 - Manufacturing Purpose

**Status: LOCKED**

Manufacturing exists to support:

- scale;
- repeatable production;
- supply chains;
- businesses;
- organisations.

It must not replace player specialisation or social interaction.

---

## MFG-002 - Automation

**Status: LOCKED**

Automation is permitted.

Automation must not become the dominant path to economic success.

---

## MFG-003 - Factories

**Status: PROPOSED**

Factories should specialise in repeatable production.

They may provide:

- throughput;
- consistency;
- volume;
- reduced manual labour.

Factories should not necessarily provide the highest possible quality.

---

## MFG-004 - Factory Constraints

**Status: PROPOSED**

Industrial scale should encounter meaningful constraints involving combinations of:

- capital;
- land;
- maintenance;
- resources;
- logistics;
- energy/fuel where applicable;
- labour;
- skill;
- market demand;
- organisation;
- transport.

---

## MFG-005 - Scale Requires Coordination

**Status: LOCKED**

Large-scale production should increasingly require coordination.

The larger the operation becomes, the more it should interact with:

- organisations;
- logistics;
- supply networks;
- specialists;
- infrastructure;
- markets.

This is a direct defence against the single-player industrial empire anti-goal. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

## MFG-006 - Artisan/Factory Relationship

**Status: PROPOSED**

Factories should be strongest at:

> **repeatable scale**

while specialist artisans should retain meaningful advantages in:

> **exceptional, customised or high-value production.**

Whether factory production can ever reach exceptional quality is **TBD** and requires simulation.

---

# 14. Items

## ITEM-001 - Extensible Item Architecture

**Status: PROPOSED**

Items should support broad categories rather than requiring bespoke systems for every item type.

Potential categories include:

- equipment;
- clothing;
- tools;
- components;
- consumables;
- structures;
- vehicles;
- decorative objects;
- commercial goods.

The final item taxonomy is **TBD**.

---

## ITEM-002 - Item Quality

**Status: PROPOSED**

Items may have quality characteristics arising from:

- resource inputs;
- crafting;
- manufacturing;
- experimentation;
- modification;
- repair history.

---

## ITEM-003 - Degradation

**Status: LOCKED PRINCIPLE**

Items should degrade, be consumed, require maintenance, or otherwise leave the economy.

This creates continuing demand for:

- resources;
- repair;
- replacement;
- crafting;
- services.

The exact degradation model is **TBD**.

---

## ITEM-004 - Repair

**Status: PROPOSED**

Repair should be a meaningful economic activity.

Repair should not necessarily restore an item indefinitely at zero consequence.

Potential consequences may include:

- material cost;
- quality loss;
- maximum durability reduction;
- specialist skill;
- time;
- replacement trade-offs.

Exact rules are **TBD**.

---

# 15. Economy

## ECO-001 - Materially Connected Player Economy

**Status: LOCKED**

The economy must connect:

```text
Resources
   ↓
Processing
   ↓
Crafting
   ↓
Manufacturing
   ↓
Retail
   ↓
Consumption
   ↓
Degradation / Replacement
   ↓
New Demand
```

The economy must be connected to the physical and social world.

---

## ECO-002 - Currency

**Status: PROPOSED**

A primary currency is assumed for general transactions.

Its exact:

- name;
- source;
- sinks;
- supply;
- taxation;
- circulation

are **TBD**.

---

## ECO-003 - Supply and Demand

**Status: PROPOSED**

Prices should respond to combinations of:

- supply;
- demand;
- scarcity;
- quality;
- competition;
- location;
- transport;
- reputation;
- player behaviour.

---

## ECO-004 - Scarcity

**Status: LOCKED**

Scarcity must matter economically.

If everything is equally abundant and available, resource discovery, geography and specialisation lose much of their value.

---

## ECO-005 - Economic Sinks

**Status: PROPOSED**

The economy should contain recurring material and currency sinks.

Potential sinks include:

- maintenance;
- repairs;
- transport;
- construction;
- property;
- services;
- consumables;
- replacement.

---

## ECO-006 - NPC Baseline Economy

**Status: PROPOSED**

NPCs may provide baseline economic functions.

They should not compete with exceptional player production.

The NPC economy should act as a **floor**, not a ceiling.

---

## ECO-007 - Imperfect Information

**Status: PROPOSED**

Perfect information may eliminate much of the value of:

- merchants;
- surveyors;
- regional knowledge;
- reputation;
- logistics.

Some degree of information asymmetry may therefore be desirable.

The system must not, however, make essential information permanently monopolised.

---

# 16. Retail and Commerce

## RET-001 - Player Retail

**Status: LOCKED**

Player-run retail is a major legitimate playstyle.

---

## RET-002 - Shops and Vendors

**Status: PROPOSED**

Players should be able to establish commercial interfaces such as:

- shops;
- storefronts;
- market stalls;
- vendors;
- commercial buildings.

The exact implementation is **TBD**.

---

## RET-003 - Location Matters

**Status: PROPOSED**

Commercial location should affect:

- visibility;
- convenience;
- traffic;
- local demand;
- transportation;
- competition;
- regional access.

This gives buildings and cities genuine economic meaning.

---

## RET-004 - Merchant Specialisation

**Status: PROPOSED**

Commerce itself should be a legitimate specialisation.

A merchant may create value through:

- sourcing;
- logistics;
- market knowledge;
- pricing;
- presentation;
- reputation;
- location;
- customer relationships.

---

# 17. Services

## SERV-001 - Services Are First-Class Gameplay

**Status: LOCKED**

Services are legitimate player economic activities.

They must not be treated as secondary to item production.

---

## SERV-002 - Medical Services

**Status: PROPOSED**

Medical specialists may provide:

- treatment;
- recovery;
- preparation;
- combat support;
- other health-related services.

Exact mechanics depend upon combat and injury design.

---

## SERV-003 - Entertainment

**Status: PROPOSED**

Entertainment should provide social and potentially mechanical value.

It may help create:

- social hubs;
- events;
- venues;
- reputations;
- community identity.

---

## SERV-004 - Other Services

**Status: PROPOSED**

Potential service professions include:

- repair;
- construction;
- research;
- transport;
- appraisal;
- specialist consultation.

The exact service model is **TBD**.

A service should create value without simply imposing artificial waiting time.

---

# 18. Combat

## CMBT-001 - Combat's Role

**Status: LOCKED**

Combat is one legitimate profession and gameplay loop.

Combat should contribute to the economy through demand for:

- weapons;
- armour;
- medical services;
- repairs;
- consumables;
- transport;
- specialised equipment.

Combat must not dominate the entire world economy.

---

## CMBT-002 - Detailed Combat

**Status: DEFERRED**

Detailed combat rules require a dedicated system specification.

---

## CMBT-003 - PvP

**Status: TBD**

The project has not yet finalised:

- open PvP;
- consensual PvP;
- territorial conflict;
- crime;
- faction warfare;
- insurance;
- loss;
- player governance implications.

PvP is a high-dependency decision affecting:

- cities;
- transport;
- risk;
- ownership;
- economy;
- combat;
- insurance.

It should therefore not be prematurely locked. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

# 19. Buildings, Housing, and Cities

## BLD-001 - Meaningful Player Structures

**Status: LOCKED**

Players and organisations should be able to create meaningful persistent structures.

Potential structures include:

- homes;
- shops;
- workshops;
- factories;
- service facilities;
- civic buildings;
- entertainment venues.

---

## BLD-002 - Multi-Function Structures

**Status: PROPOSED**

A structure may potentially combine:

- residence;
- commerce;
- production;
- storage;
- services;
- social activity.

Exact structural rules are **TBD**.

---

## BLD-003 - Placement and Maintenance

**Status: PROPOSED**

Buildings should encounter meaningful constraints involving:

- land;
- location;
- ownership;
- maintenance;
- zoning;
- infrastructure;
- clutter.

These constraints should prevent uncontrolled world clutter and monopolisation.

---

## CITY-001 - Player Cities

**Status: PROPOSED**

Player cities should be institutions rather than merely collections of buildings.

A city may eventually contain:

- governance;
- zoning;
- commercial districts;
- services;
- organisations;
- infrastructure;
- taxation;
- civic identity.

The exact model is **TBD**.

---

# 20. Social Systems and Organizations

## SOC-001 - Persistent Organisations

**Status: LOCKED**

Players must be able to form persistent organisations.

Organisations may coordinate:

- production;
- commerce;
- resource extraction;
- logistics;
- combat;
- cities;
- events;
- social activities.

---

## SOC-002 - Organisation Roles and Assets

**Status: PROPOSED**

Organisations should eventually support:

- roles;
- permissions;
- assets;
- shared infrastructure;
- contracts;
- internal governance.

Fraud and asset-theft protections require dedicated design.

---

## SOC-003 - Evidence-Based Reputation

**Status: PROPOSED**

Reputation should not be reduced to a single meaningless number.

Relevant evidence may include:

- product provenance;
- quality history;
- sales history;
- service history;
- reliability;
- organisational history;
- customer behaviour.

The exact presentation and calculation model is **TBD**.

---

## SOC-004 - Interdependence Networks

**Status: LOCKED**

Organisations and social networks should be capable of coordinating chains such as:

```text
Resource Discovery
      ↓
Extraction
      ↓
Processing
      ↓
Manufacturing
      ↓
Retail
      ↓
Consumption
```

The goal is not to force guild membership.

The goal is to make cooperation economically meaningful.

---

# 21. Progression

## PROG-001 - Capability and Identity

**Status: LOCKED**

Progression should develop:

- capability;
- expertise;
- specialisation;
- reputation;
- access;
- social influence.

It should not simply create a universal linear power curve.

---

## PROG-002 - Skill Acquisition

**Status: TBD**

The project has not yet finalised whether skills are acquired through:

- use;
- training;
- skill points;
- mentorship;
- certifications;
- combinations of these.

This is a high-dependency decision.

---

## PROG-003 - Mastery

**Status: PROPOSED**

Mastery should represent genuine expertise rather than merely acting as a gatekeeping mechanism.

High-level mastery should create:

- better output;
- access to advanced work;
- reputation;
- economic demand;
- unique opportunities.

---

# 22. Emergent Gameplay

## EMRG-001 - Emergent Scenario Framework

**Status: LOCKED**

The game should be capable of producing systemic situations such as:

### 1. Gold Rush

An exceptional resource appears.

### 2. Regional Specialisation

A region becomes economically important because of its resources or geography.

### 3. Supply Shortage

A shortage causes merchants and manufacturers to form new relationships.

### 4. Specialist Reputation

A remote specialist gains customers because of reputation rather than location alone.

### 5. Industrial Organisation

A player organisation coordinates:

- extraction;
- processing;
- manufacturing;
- logistics;
- retail.

### 6. Economic Cascade

A shortage or discovery affects multiple industries.

### 7. Player-Created Event

Players create an event that generates:

- travel;
- trade;
- entertainment;
- services.

### 8. Social Institution

A shop, venue or organisation becomes a recognised community institution.

### 9. Non-Combat Prestige

A player becomes socially important without being a top combatant.

### 10. Historical Consequence

Player actions permanently alter a region, business, organisation or market.

These scenarios are evaluation targets for the entire system architecture.

---

# 23. Content Expansion Framework

## EXP-001 - Extensible Content

**Status: LOCKED**

Core systems must use extensible abstractions.

Future content should be able to introduce:

- resources;
- resource attributes;
- schematics;
- items;
- professions;
- buildings;
- services;
- regions;
- organisations;
- NPC roles;
- transportation;
- research;
- world events.

without requiring fundamental redesign of the underlying systems.

`TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

## EXP-002 - Content Declaration

**Status: PROPOSED**

New content should declare, where applicable:

- prerequisites;
- inputs;
- outputs;
- category;
- quality rules;
- rarity;
- progression impact;
- economic impact;
- social impact;
- balance risks.

Adding content solely to increase volume is insufficient.

New content should reinforce at least one design pillar.

---

# 24. Balance Philosophy

## BAL-001 - No Permanent Universal Optimum

**Status: PROPOSED**

The economy should not converge permanently on one:

- resource;
- profession;
- product;
- region;
- strategy.

Variation should arise from:

- temporary resource quality;
- availability;
- player skill;
- demand;
- logistics;
- customisation;
- reputation;
- situational requirements.

---

## BAL-002 - Accessibility Without Forced Dependency

**Status: PROPOSED**

Interdependence must not become hostage gameplay.

Players should have access to:

- basic participation;
- emergency support;
- basic functionality.

Player specialists should provide advantages in:

- quality;
- efficiency;
- convenience;
- expertise;
- exceptional performance.

This distinction is important:

> **Players should be valuable because they are better, not because the game arbitrarily prevents everyone else from doing anything.**

---

## BAL-003 - Casual and Veteran Participation

**Status: PROPOSED**

Casual players should be able to contribute meaningfully.

Veteran players should gain:

- expertise;
- reputation;
- specialisation;
- organisational influence;
- complex opportunities.

Veteran status should not automatically mean:

> "I can efficiently do everything myself."

---

# 25. Exploit, Fraud, and Economic Stability

## SAFE-001 - First-Class Concern

**Status: LOCKED**

Fraud and exploit resistance are design requirements, not merely implementation concerns.

This applies to:

- ownership;
- markets;
- organisations;
- automation;
- reputation;
- contracts;
- item transfer.

---

## SAFE-002 - Areas Requiring Dedicated Design

**Status: TBD**

Detailed systems are required for:

- item duplication;
- currency duplication;
- resource duplication;
- vendor arbitrage;
- land monopolisation;
- resource monopolisation;
- multi-account abuse;
- market manipulation;
- organisation asset theft;
- contract fraud;
- reputation manipulation;
- abandoned assets;
- governance griefing.

---

## SAFE-003 - Economic Monitoring

**Status: PROPOSED**

The eventual game should monitor:

- currency creation/destruction;
- resource extraction;
- item production/destruction;
- price distributions;
- market concentration;
- regional supply;
- factory utilisation;
- organisation asset concentration;
- new-player access;
- abandoned property.

This should inform future balancing rather than directly dictate player behaviour.

---

# 26. Testing Requirements and Design Invariants

## TEST-001 - Core Invariants

**Status: PROPOSED**

Future implementations should test at minimum:

1. Items cannot appear without valid creation sources.
2. Provenance cannot be falsely assigned.
3. Ownership transfers are auditable.
4. Invalid resources cannot satisfy schematics.
5. Resource extraction obeys resource rules.
6. Manufacturing cannot create invalid output.
7. Currency obeys explicit source/sink rules.
8. Skill caps enforce intended specialisation constraints.
9. Buildings obey ownership and placement rules.
10. Degradation/repair cannot become an unintended value generator.
11. Automation cannot create unlimited unattended production.
12. No ordinary system can generate infinite risk-free currency or item value.

The source documents identify these invariants as important foundations for later testing. `TCIndustries Master Game Design Document v0.1 - gpt-5.6-terra-medium.md`

---

## TEST-002 - Economic Simulation

**Status: PROPOSED**

Before important economic systems become LOCKED, they should be simulated against scenarios such as:

- resource abundance;
- resource scarcity;
- monopolisation;
- price collapse;
- price spikes;
- factory dominance;
- crafter dominance;
- new-player participation;
- veteran concentration;
- regional isolation;
- transportation disruption.

---

## TEST-003 - Promotion Criteria

**Status: PROPOSED**

A proposed system should not become LOCKED until it has:

1. defined purpose;
2. defined inputs;
3. defined outputs;
4. defined state;
5. failure cases;
6. dependencies;
7. exploit analysis;
8. player-visible consequences;
9. pillar compatibility;
10. prototype or simulation evidence where appropriate.

---

# 27. Historical SWG Relationship

## SWG-001 - Inspiration Without IP Reproduction

**Status: LOCKED**

Historical SWG is an important design reference.

TCIndustries must remain an original IP.

---

## SWG-002 - Three-Way Documentation Distinction

Every significant historically inspired mechanic should be identified as one of:

### Historical

What the original game did.

### Adaptation

How TCIndustries changes the concept.

### Invention

What TCIndustries introduces independently.

This prevents historical inspiration from being mistaken for canonical TCIndustries design.

---

# 28. Prototype Relationship

## PROTO-001 - Prototype Is Not Authority

**Status: LOCKED**

The TypeScript/Vite/React prototype is experimental.

It is:

> **evidence, not law.**

Prototype behaviour must not silently become canonical because it already exists in code.

---

## PROTO-002 - Development Hierarchy

The intended hierarchy is:

```text
GDD
 ↓
System Specification
 ↓
Prototype
 ↓
Testing / Simulation
 ↓
Golden Master
 ↓
Production Implementation
```

Technical decisions such as:

- engine;
- networking;
- database;
- rendering;
- UI architecture;
- deployment

belong in the TDD rather than this document.

The source GDDs consistently establish this separation. 

---

# 29. Systems Requiring Deeper Design

The Master GDD intentionally does not pretend that every system is finished.

The following require dedicated specifications.

## Highest Priority

1. Resource spawning
2. Resource attributes
3. Resource quality
4. Surveying
5. Extraction
6. Skill architecture
7. Specialisation budget
8. Crafting quality
9. Experimentation
10. Item degradation

## Next

11. Economy
12. Pricing
13. Player retail
14. Reputation
15. Manufacturing
16. Automation limits
17. Buildings
18. Ownership
19. Services
20. Organisations

## Later

21. Cities
22. Transportation
23. World simulation
24. NPC society
25. Combat
26. PvP
27. Housing
28. Governance

## Deferred Technical Work

29. Engine architecture
30. Networking
31. Database
32. Rendering
33. UI architecture
34. Deployment

This ordering is substantially consistent with the design sequences in the source documents. `Tiwakings Craftworld Industries — Master Game Design Document v0.1-chatgpt1.md`

---

# 30. Recommended Design Sequence

## Phase 1 - Resource Foundation

Define:

- resource taxonomy;
- attributes;
- quality;
- spawning;
- scarcity;
- regional variation;
- temporal variation;
- surveying;
- extraction;
- processing.

---

## Phase 2 - Crafting

Define:

- schematics;
- inputs;
- components;
- quality;
- experimentation;
- exceptional products;
- provenance.

---

## Phase 3 - Items

Define:

- item identity;
- quality;
- provenance;
- degradation;
- repair;
- modification;
- replacement.

---

## Phase 4 - Skills and Professions

Define:

- skill domains;
- learning;
- progression;
- mastery;
- specialisation;
- capability cap;
- respecialisation.

---

## Phase 5 - Economy

Define:

- currency;
- supply;
- demand;
- pricing;
- markets;
- trade;
- sinks;
- economic concentration.

---

## Phase 6 - Retail and Reputation

Define:

- shops;
- vendors;
- discovery;
- customer behaviour;
- reputation;
- brands;
- commercial location.

---

## Phase 7 - Manufacturing

Define:

- factories;
- production templates;
- automation;
- capacity;
- maintenance;
- logistics;
- scale;
- artisan/factory boundaries.

---

## Phase 8 - Social World

Define:

- organisations;
- permissions;
- shared assets;
- contracts;
- cities;
- governance;
- ownership.

---

## Phase 9 - Services and Combat

Integrate:

- medicine;
- entertainment;
- repair;
- construction;
- transport;
- combat;
- PvP.

---

## Phase 10 - World Simulation

Integrate the preceding systems into a coherent persistent simulation.

The source documents specifically recommend resolving the resource foundation before moving into crafting, skills, economy and later world systems. `Tiwakings Craftworld Industries — Master Game Design Document v0.1-chatgpt1.md`

---

# 31. Status Register

The following is the consolidated high-level status register.

| ID | Design Element | Status |
|---|---|---|
| VIS-001 | Living virtual world | **LOCKED** |
| VIS-002 | Citizen rather than chosen hero | **LOCKED** |
| VIS-003 | Player-driven economy | **LOCKED** |
| VIS-004 | Interdependence | **LOCKED** |
| VIS-005 | Emergent gameplay | **LOCKED** |
| PIL-001 | Discovery / Gold Rush | **LOCKED** |
| PIL-002 | Identity / Reputation | **LOCKED** |
| PIL-003 | Interdependence | **LOCKED** |
| PIL-004 | Player-created society | **LOCKED** |
| LOOP-002 | Non-combat viability | **LOCKED** |
| WRLD-001 | Persistent world | **LOCKED** |
| WRLD-002 | Regional differentiation | **PROPOSED** |
| WRLD-003 | NPC/player settlements/wilderness | **PROPOSED** |
| WRLD-004 | Persistent environmental simulation | **PROPOSED** |
| WRLD-005 | Constrained transportation | **PROPOSED** |
| PLR-001 | Persistent player identity | **LOCKED** |
| PLR-002 | Player/organisation ownership | **LOCKED** |
| PROF-001 | Flexible skills/professions | **LOCKED** |
| PROF-003 | Meaningful specialisation | **PROPOSED** |
| PROF-004 | Respecialisation | **LOCKED** |
| PROF-005 | Capability cap | **PROPOSED** |
| RES-001 | Resource discovery/scarcity | **LOCKED** |
| RES-002 | Resource instances | **PROPOSED** |
| RES-003 | Resource attributes | **PROPOSED** |
| RES-004 | Resource quality | **LOCKED PRINCIPLE** |
| RES-005 | Dynamic occurrences | **PROPOSED** |
| RES-006 | Surveying profession | **PROPOSED** |
| RES-007 | Extraction | **PROPOSED** |
| RES-008 | Processing | **PROPOSED** |
| CRFT-001 | Differentiated crafting | **LOCKED** |
| CRFT-002 | Resource choice affects output | **LOCKED** |
| CRFT-003 | Schematics | **PROPOSED** |
| CRFT-005 | Product quality | **PROPOSED** |
| CRFT-008 | Provenance | **LOCKED** |
| CRFT-009 | Signature products | **LOCKED** |
| MFG-001 | Manufacturing supports players | **LOCKED** |
| MFG-002 | Automation permitted | **LOCKED** |
| MFG-003 | Factories | **PROPOSED** |
| MFG-004 | Factory constraints | **PROPOSED** |
| MFG-005 | Scale requires coordination | **LOCKED** |
| ITEM-003 | Degradation | **LOCKED PRINCIPLE** |
| ITEM-004 | Repair | **PROPOSED** |
| ECO-001 | Materially connected economy | **LOCKED** |
| ECO-002 | Primary currency | **PROPOSED** |
| ECO-003 | Supply/demand pricing | **PROPOSED** |
| ECO-004 | Scarcity matters | **LOCKED** |
| ECO-005 | Economic sinks | **PROPOSED** |
| ECO-006 | Limited NPC economy | **PROPOSED** |
| RET-001 | Player retail | **LOCKED** |
| RET-003 | Location affects commerce | **PROPOSED** |
| SERV-001 | Services as first-class gameplay | **LOCKED** |
| CMBT-001 | Combat supports economy | **LOCKED** |
| CMBT-002 | Detailed combat | **DEFERRED** |
| CMBT-003 | PvP | **TBD** |
| BLD-001 | Meaningful player structures | **LOCKED** |
| CITY-001 | Player cities | **PROPOSED** |
| SOC-001 | Persistent organisations | **LOCKED** |
| SOC-003 | Evidence-based reputation | **PROPOSED** |
| SOC-004 | Interdependence chains | **LOCKED** |
| EXP-001 | Extensible content architecture | **LOCKED** |
| SAFE-001 | Fraud/exploit resistance | **LOCKED** |
| SAFE-002 | Detailed fraud rules | **TBD** |
| TEST-001 | Core invariants | **PROPOSED** |
| TEST-002 | Economic simulation | **PROPOSED** |
| TEST-003 | Promotion criteria | **PROPOSED** |
| SWG-001 | SWG inspiration without IP reproduction | **LOCKED** |
| SWG-002 | Historical/adaptation/invention distinction | **LOCKED** |
| PROTO-001 | Prototype non-authority | **LOCKED** |

---

# 32. Assumption Register

These assumptions have been carried forward because they are necessary to make the current architecture coherent, but they are **not automatically locked decisions**.

| ID | Assumption | Consequence if Incorrect |
|---|---|---|
| ASM-001 | TCIndustries is a persistent online multiplayer world | Core world/economy architecture changes |
| ASM-002 | Dynamic resources remain central | Resource/crafting/economy architecture changes |
| ASM-003 | Crafted products can meaningfully differ | Crafter identity/reputation weakens |
| ASM-004 | Player-owned structures exist | Retail/manufacturing/city design changes |
| ASM-005 | Original setting/IP | Content must remain distinct from SWG |
| ASM-006 | Simulation/testing will precede final production rules | Balance process changes |
| ASM-007 | Some currency model exists | Barter/multi-currency design required instead |
| ASM-008 | Items require degradation/consumption/maintenance | Alternative economic sinks required |
| ASM-009 | Skill specialisation requires some form of cap | Universal mastery risk increases |
| ASM-010 | Manufacturing and extraction automation exist | Industrial architecture changes |
| ASM-011 | Player organisations exist | Social/economic scale changes |
| ASM-012 | Player cities are desired | World/ownership architecture changes |

---

# 33. Open-Question Register

The following are intentionally unresolved.

| ID | Question | Importance |
|---|---|---|
| OQ-001 | What is the exact setting, genre, technology level and fiction? | Very High |
| OQ-002 | What exact resource attributes exist? | Very High |
| OQ-003 | How are resources spawned? | Very High |
| OQ-004 | How long do resource occurrences last? | Very High |
| OQ-005 | How does extraction affect local resource state? | High |
| OQ-006 | How much surveying information is revealed at each skill level? | High |
| OQ-007 | What is the exact crafting quality formula? | Very High |
| OQ-008 | What is the exact experimentation model? | Very High |
| OQ-009 | Does experimentation permit catastrophic failure? | High |
| OQ-010 | What exact skill architecture is used? | Very High |
| OQ-011 | What is the capability/specialisation cap? | Very High |
| OQ-012 | How are skills acquired and improved? | Very High |
| OQ-013 | What exactly does respecialisation cost? | High |
| OQ-014 | What is the exact item degradation model? | High |
| OQ-015 | What is the currency model? | Very High |
| OQ-016 | How are prices established? | Very High |
| OQ-017 | How do player shops attract customers? | High |
| OQ-018 | How is reputation represented? | High |
| OQ-019 | How are factories constrained? | Very High |
| OQ-020 | Can factories produce exceptional-quality goods? | Very High |
| OQ-021 | What transportation methods exist? | High |
| OQ-022 | How are cities founded? | High |
| OQ-023 | How are cities governed, taxed and dissolved? | High |
| OQ-024 | What organisation powers and liabilities exist? | High |
| OQ-025 | What is the combat system? | High |
| OQ-026 | What is the PvP model? | Very High |
| OQ-027 | What are injury/death/item-loss rules? | High |
| OQ-028 | How are services made valuable without forced downtime? | High |
| OQ-029 | What anti-multi-account rules are required? | High |
| OQ-030 | What contract/escrow/asset protection model exists? | High |
| OQ-031 | What systems are required for first playable prototype? | High |
| OQ-032 | What systems are required for persistent alpha? | High |

---

# 34. Glossary

| Term | Meaning |
|---|---|
| **Attribute** | A measurable property of a resource or item used by system rules. |
| **Brand** | Player or organisation identity associated with products, services or businesses. |
| **Crafting** | Player-directed creation of differentiated goods. |
| **Extraction** | Acquiring resources from the world. |
| **Factory** | A system or structure used for constrained repeatable production. |
| **Gold Rush** | A temporary player-driven economic event caused by an unusually valuable resource opportunity. |
| **Intermediate Good** | A processed material or component used to create another product. |
| **NPC Settlement** | A stable world-authored settlement providing baseline infrastructure and services. |
| **Organisation** | A persistent player-created social/economic group. |
| **Player City** | A player-created settlement functioning as a persistent institution. |
| **Provenance** | Persistent information identifying an item's creator and relevant production history. |
| **Resource Instance** | A specific resource occurrence with identity, attributes, location and lifecycle. |
| **Resource Pool** | A geographically bounded occurrence of a resource class with a defined lifecycle. |
| **Schematic** | A formal specification governing the creation of an item or product. |
| **Service Profession** | A profession whose primary economic output is a service rather than a physical item. |
| **Specialisation Budget** | Proposed mechanism limiting simultaneous high-tier expertise. |
| **Virtual Society** | The network of player-created relationships, institutions, businesses, organisations and communities. |

---

# Final Canonical Statement

TCIndustries is:

> **An expandable, systemic, player-driven virtual-world MMORPG in which resource discovery, crafting, specialisation, commerce, services, reputation, ownership and social interdependence combine to create emergent player lives.**

Its four foundational truths are:

1. **Discovery creates opportunity.**
2. **Identity creates meaning.**
3. **Interdependence creates society.**
4. **Player action creates the world.**

The game must therefore continually resist the forces that would turn it into:

- a theme park;
- a class-based RPG;
- a combat-only MMO;
- a spreadsheet;
- an automated factory simulator;
- or a collection of disconnected systems.

The intended experience is not:

> **"What does the game want me to do?"**

It is:

> **"What do I want to become in this world?"**

And the ultimate systemic test is:

> **Can the game produce a world where another player can answer, "I know who you are, what you do, what you make, and why it matters"?**

That is the canonical design target.



---
Powered by [ChatGPT Exporter](https://www.chatgptexporter.com)