# TCIndustries — Project Architecture Description for OpenCode

**Document Type:** Project Architecture Input / Development Governance Description  
**Project:** Tiwakings Craftworld Industries (TCIndustries)  
**Document Version:** 1.0  
**Status:** NON-CANONICAL — Architectural Planning Input  
**Author:** OpenAI GPT-5.6 Luna  
**Date:** 2026-09-12  
**Intended Consumer:** OpenCode / AI-assisted development environment  
**Purpose:** Provide OpenCode with a precise description of the TCIndustries project so that OpenCode can design a targeted repository, documentation, development, testing, and implementation architecture.

---

# 1. Document Authority

This document is **not a game-design authority document**.

It does not create, modify, promote, demote, or supersede any TCIndustries game-design rule.

Its purpose is to describe the project and its governance requirements so that OpenCode can design an appropriate technical/project architecture.

The current working canonical game-design reference is:

`TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`

The authority and provenance framework is defined by:

`TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`

The historical consolidation and audit documents are provenance and governance inputs, not substitutes for the current canonical baseline.

OpenCode must preserve this distinction.

---

# 2. Project Identity

## 2.1 Project Name

**Tiwakings Craftworld Industries**

**Short Name:** TCIndustries

TCIndustries is an original persistent sandbox MMORPG / virtual-world game project.

The project is heavily inspired by the systemic and social philosophy of early *Star Wars Galaxies*, particularly:

- dynamic resources,
- player-created economic activity,
- specialised professions,
- player crafting,
- meaningful item provenance,
- player commerce,
- settlements,
- social services,
- player organisations,
- persistent world consequences,
- emergent player-driven events.

However, TCIndustries is an **original intellectual property project**.

Historical SWG behaviour is reference material only. It must never automatically become TCIndustries design authority.

Where SWG-inspired concepts are researched or discussed, documentation must distinguish:

1. Historical SWG behaviour.
2. TCIndustries' proposed adaptation.
3. New TCIndustries invention.

---

# 3. Core Player Experience

The central player fantasy is:

> **Become somebody known in a persistent world.**

Players are intended to feel like citizens of a living society rather than protagonists of a predetermined heroic narrative.

A player may become, among other possibilities:

- a resource prospector,
- extractor,
- farmer,
- processor,
- craftsperson,
- manufacturer,
- merchant,
- trader,
- architect,
- builder,
- medical practitioner,
- entertainer,
- transport operator,
- researcher,
- combatant,
- city organiser,
- organisation leader,
- or hybrid specialist.

The project is therefore fundamentally a **systems-driven social sandbox**, not a conventional theme-park MMORPG.

The primary content is expected to emerge from interactions between:

**players + resources + production + services + geography + ownership + reputation + organisations + markets + world systems.**

---

# 4. Locked High-Level Design Principles

The following principles are part of the current TCIndustries design baseline and must be treated as high-authority project constraints when designing implementation architecture.

## 4.1 Living Persistent Virtual World

TCIndustries is a persistent multiplayer virtual world in which players create economic, social, occupational, and reputational identities.

The world should continue to exist independently of any particular player's personal story.

## 4.2 Citizen Rather Than Chosen Hero

The player is not necessarily the centre of the world's narrative.

Prestige should be achievable through non-combat accomplishments such as:

- discovering valuable resources,
- producing respected goods,
- operating successful businesses,
- providing important services,
- organising communities,
- creating desirable buildings,
- operating trade networks,
- becoming a specialist known by other players.

## 4.3 Player-Driven Economy

Players should be major participants in production and distribution.

Resources, labour, processing, transportation, crafting, services, consumption, ownership, and regional differences should have meaningful economic relationships.

## 4.4 Interdependence

Important professions should depend upon one another.

The architecture must therefore avoid making the entire economic ecosystem trivially self-sufficient for a single character, account, organisation, or automated system.

The precise mechanical implementation of interdependence remains an active design problem.

## 4.5 Emergence

The project favours systems that create conditions for unexpected player-driven outcomes rather than scripting every significant experience.

Architecture should therefore support systemic interaction and simulation rather than hard-coding every gameplay outcome into isolated bespoke content.

## 4.6 Discovery and the Gold Rush

Resources are intended to create discovery, scarcity, regional economic differences, temporary opportunities, and social events.

A major resource discovery should potentially become a player-driven event rather than merely an inventory transaction.

## 4.7 Identity and Reputation

Player identity should have persistent social meaning.

Products, services, businesses, organisations, and structures should be capable of maintaining meaningful associations with their creators, owners, providers, or brands.

## 4.8 Player-Created Society

Players should eventually be able to create meaningful:

- businesses,
- organisations,
- settlements,
- cities,
- social institutions,
- commercial districts,
- trade networks,
- and reputational institutions.

## 4.9 Non-Combat Viability

A player must be able to pursue a meaningful long-term career without making combat their primary activity.

Combat is one possible lifestyle rather than the mandatory progression path.

## 4.10 Flexible Skill-Based Professions

TCIndustries uses a flexible skill-based profession model rather than rigid immutable classes.

Profession identity emerges from combinations of:

- skills,
- specialisation,
- equipment,
- reputation,
- ownership,
- social relationships,
- and player behaviour.

Exact skill trees, costs, progression mathematics, and specialisation limits remain unresolved.

## 4.11 Persistent Identity and Ownership

Characters and players are intended to accumulate meaningful persistent identity and ownership.

Potential persistent entities include:

- characters,
- items,
- resources,
- structures,
- vendors,
- businesses,
- organisations,
- brands,
- production assets,
- contracts,
- reputational history,
- provenance.

The exact ownership, transfer, decay, taxation, and binding mechanics remain open unless explicitly specified elsewhere.

## 4.12 Controlled Automation

Automation and industrialisation are allowed as design concepts but must not eliminate meaningful player roles.

The project explicitly rejects drift toward an infinite, fully automated, self-sufficient industrial simulation.

---

# 5. Explicit Anti-Goals

The project must not silently evolve into:

1. A conventional quest-driven theme-park MMORPG.
2. A rigid class-based MMORPG.
3. A single-player crafting game with multiplayer chat.
4. An infinite personal-factory automation game.
5. A pure market spreadsheet simulator.
6. A combat-only MMORPG.
7. An economy in which NPCs perform all meaningful production.
8. A world where one player can efficiently perform every economically important role.
9. A game whose prototype implementation becomes its design authority merely because code already exists.

These anti-goals should influence architecture, testing, simulation, and development workflow.

---

# 6. Design Authority Model

This project requires unusually strict separation between **design authority**, **implementation**, **evidence**, and **historical reference**.

OpenCode must preserve the following conceptual hierarchy.

```text
HUMAN / PROJECT AUTHORITY
        │
        ▼
CANONICAL GAME DESIGN
        │
        ├── LOCKED
        ├── DERIVED CONSTRAINT
        ├── PROPOSED
        ├── TBD
        └── DEFERRED
        │
        ▼
IMPLEMENTATION SPECIFICATIONS
        │
        ▼
PROTOTYPES / SIMULATIONS / TESTS
        │
        ▼
VALIDATED IMPLEMENTATION
```

Separately:

```text
HISTORICAL RESEARCH
       │
       ├── SWG reference
       ├── external research
       └── comparative analysis
                │
                ▼
        DESIGN INPUT ONLY
```

Historical research and prototype behaviour do **not** automatically acquire design authority.

---

# 7. Required Status Model

The documentation architecture must support at least these statuses:

| Status | Meaning |
|---|---|
| **LOCKED** | Explicit project/human decision and authoritative unless formally revised. |
| **DERIVED CONSTRAINT** | Genuine logical consequence of LOCKED principles; binding but not an independent human decision. |
| **PROPOSED** | Recommended design direction awaiting approval. |
| **TBD** | Unresolved design decision. |
| **PROTOTYPE / EVIDENCE** | Behaviour demonstrated experimentally; not canonical merely because it exists. |
| **DEFERRED** | Recognised future work intentionally postponed. |
| **ASSUMPTION** | Working assumption not established as authoritative. |
| **HISTORICAL** | External or historical reference material, particularly SWG material. |

### Critical Governance Requirement

OpenCode must never silently convert:

- PROPOSED → LOCKED
- TBD → LOCKED
- PROTOTYPE → LOCKED
- ASSUMPTION → LOCKED
- HISTORICAL → TCIndustries rule
- DERIVED CONSTRAINT → independent human decision

A change of design authority must be explicitly documented.

---

# 8. Provenance Requirements

Every significant design decision should be traceable.

The project uses stable identifiers such as:

- `VIS-###`
- `PIL-###`
- `LOOP-###`
- `WRLD-###`
- `PLR-###`
- `PROF-###`
- `RES-###`
- `CRFT-###`
- `MFG-###`
- `ECO-###`
- `RET-###`
- `SERV-###`
- `CMBT-###`
- `BLD-###`
- `SOC-###`
- `SAFE-###`
- `TEST-###`
- `OQ-###`

OpenCode should design documentation architecture around stable identifiers rather than relying solely upon filenames or document chronology.

The project explicitly rejects the assumption that:

> newer file = more authoritative file

or:

> more detailed document = more authoritative document.

Authority must be determined by the project's governance rules.

---

# 9. Separation of Documentation Domains

The project should not be represented as one enormous undifferentiated document.

OpenCode should design an architecture capable of separating at least the following documentation domains.

## 9.1 Game Design Authority

Contains canonical game rules, principles, constraints, and approved system specifications.

Example:

`docs/design/`

or another architecture OpenCode determines is superior.

## 9.2 Governance and Authority

Contains:

- authority rules,
- status definitions,
- provenance rules,
- change-control rules,
- decision records,
- promotion/revision procedures.

## 9.3 System Specifications

Each sufficiently mature subsystem should eventually have its own focused specification.

Examples:

- resources,
- skills,
- crafting,
- manufacturing,
- economy,
- retail,
- services,
- combat,
- buildings,
- cities,
- organisations,
- transportation,
- reputation.

## 9.4 Research and Historical Material

Research must remain distinct from canonical design.

This includes:

- SWG research,
- competitor analysis,
- academic/economic research,
- technical research,
- external references.

## 9.5 Prototypes

Prototype source code and prototype documentation are evidence and experimentation.

They must not be allowed to silently redefine canonical rules.

## 9.6 Simulations and Experiments

Economic and systems simulations should have a distinct home.

They exist to answer questions such as:

- Does resource scarcity produce meaningful events?
- Can small players participate?
- Does manufacturing eliminate artisans?
- Does specialisation produce meaningful interdependence?
- Does travel create economic geography?
- Does the economy become monopolised?
- Are non-combat professions viable?

## 9.7 Testing and Validation

Testing should distinguish:

- software correctness,
- system behaviour,
- design-invariant validation,
- economic simulation,
- playtest evidence,
- balance evidence.

A test passing does not automatically make the tested behaviour canonical.

## 9.8 Project Management

Open questions, work items, implementation plans, milestones, and development tasks should remain distinct from design authority.

---

# 10. Canonical Design Versus Implementation

This is a critical architectural requirement.

The project expects a separation similar to:

```text
GAME DESIGN
    │
    │ defines intended rules
    ▼
SYSTEM SPECIFICATION
    │
    │ defines implementable behaviour
    ▼
IMPLEMENTATION
    │
    │ produces actual behaviour
    ▼
TEST / SIMULATION / PLAYTEST
    │
    │ produces evidence
    ▼
DESIGN REVIEW
```

The direction must **not** be:

```text
CODE EXISTS
    ↓
THEREFORE GAME RULE
```

Prototype behaviour may reveal:

- contradictions,
- usability problems,
- balance problems,
- emergent behaviour,
- implementation constraints,
- new design opportunities.

But those observations must return through the design-governance process before becoming canonical.

---

# 11. Expected Development Architecture

OpenCode should design an architecture suitable for a project that is expected to evolve through multiple implementation generations.

The current conceptual development hierarchy is:

```text
Canonical GDD
      ↓
System Specifications
      ↓
Prototype / Experimental Implementation
      ↓
Testing / Simulation / Playtesting
      ↓
Validated Behaviour
      ↓
Production Implementation
```

The current prototype is not necessarily the final technology.

The project documentation has previously described a future transition from experimental TypeScript/Vite/React work toward a future UE5/C++ production implementation.

Therefore:

**Do not architect the documentation system around the assumption that the current prototype technology is the permanent production technology.**

The design/documentation layer must remain technology-independent wherever practical.

---

# 12. System-Centric Architecture

TCIndustries is a highly interconnected systems project.

OpenCode should therefore avoid an architecture in which gameplay systems are treated as isolated feature silos.

Important relationships include:

```text
Resources
   ↕
Extraction
   ↕
Processing
   ↕
Crafting
   ↕
Manufacturing
   ↕
Items
   ↕
Commerce
   ↕
Demand
   ↕
Services
   ↕
Players / Organisations
   ↕
Reputation
   ↕
World / Geography
```

The architecture should make dependencies explicit.

For example:

**Resource design** affects:

- crafting,
- product differentiation,
- economic scarcity,
- regional trade,
- discovery,
- manufacturing,
- reputation.

**Specialisation design** affects:

- profession viability,
- interdependence,
- crafting,
- services,
- player identity,
- economic opportunity.

**Manufacturing design** affects:

- automation,
- artisan relevance,
- organisations,
- economic concentration,
- resource demand.

Therefore system specifications should contain explicit dependency relationships rather than being treated as independent documents.

---

# 13. Economic Interdependence as a Major Architectural Concern

The project has identified an **Economic Interdependence Core** as a particularly important unresolved design cluster.

It connects:

```text
Specialisation Budget
        ↕
Resource Attributes
        ↕
Crafting Differentiation
        ↕
Manufacturing Constraints
        ↕
NPC Substitution Limits
        ↕
Player Demand
        ↕
Economic Value
```

This should be treated as a major cross-system dependency.

OpenCode should therefore design the repository so that these systems can be:

- specified independently,
- cross-referenced,
- tested together,
- simulated together,
- audited for contradictions,
- revised without duplicating rules across documents.

---

# 14. Testing Philosophy

Testing is not limited to checking whether software executes without errors.

The project expects design systems eventually to have explicit invariants and validation criteria.

Examples include:

- specialised roles remaining economically valuable;
- non-combat careers remaining viable;
- factories not eliminating all artisan relevance;
- resource discoveries producing meaningful opportunities;
- small independent businesses remaining viable;
- economic systems avoiding uncontrolled monopoly;
- geography remaining economically meaningful;
- automation not eliminating player participation;
- systems not permitting infinite risk-free generation of currency or items.

Where practical, major economic systems should be simulatable before full production implementation.

---

# 15. Documentation as Project Memory

The documentation architecture should assume that future development will involve:

- multiple LLMs,
- multiple OpenCode sessions,
- different models,
- long development periods,
- prototype rewrites,
- competing proposals,
- audits,
- playtests,
- historical research,
- human design rulings.

Therefore project state must be recoverable from repository artifacts.

OpenCode must **not** rely upon:

- conversation memory,
- one LLM's recollection,
- undocumented assumptions,
- Git chronology alone,
- code behaviour alone.

Important project state should exist in durable Markdown/project artifacts.

This is particularly important because AI-assisted development is vulnerable to context loss and design drift.

---

# 16. Decision-Making Workflow

The project should support a workflow approximately like:

```text
Question / Problem
       ↓
Research / Investigation
       ↓
Evidence
       ↓
Design Alternatives
       ↓
Design Analysis
       ↓
Human Decision
       ↓
Decision Record
       ↓
Canonical Specification
       ↓
Implementation Plan
       ↓
Implementation
       ↓
Testing / Simulation / Playtest
       ↓
Evidence Review
       ↓
Possible Design Revision
```

OpenCode should not collapse these stages into one action.

In particular:

**research ≠ decision**

**proposal ≠ decision**

**implementation ≠ decision**

**prototype success ≠ canonical approval**

**test result ≠ design ruling**

---

# 17. Human Authority

The human project owner remains the final authority over game design.

LLMs are advisors, researchers, auditors, implementers, critics, and documentation agents.

They are not autonomous design authorities.

When evidence conflicts with a canonical rule, the correct response is to surface the conflict rather than silently rewriting the rule.

When two design documents conflict, OpenCode should identify:

1. the conflicting statements;
2. their identifiers;
3. their status;
4. their provenance;
5. their authority;
6. the dependency consequences;
7. the unresolved human decision.

It should not silently choose whichever document appears newer or more detailed.

---

# 18. Change Control

Design-changing work should be distinguishable from implementation-only work.

A design-changing change should normally produce or update an appropriate formal project artifact containing:

- unique identifier,
- title,
- author,
- date,
- status,
- affected systems,
- authority/provenance,
- decision or finding,
- rationale,
- dependencies,
- consequences,
- validation requirements,
- affected canonical documents.

OpenCode should architect the repository so that this information can be found reliably.

---

# 19. Auditability Requirements

A future developer or LLM should be able to answer:

- What is currently canonical?
- Why is it canonical?
- Who approved it?
- Which document established it?
- What earlier proposals existed?
- Which proposals were rejected?
- Which parts are still TBD?
- Which behaviour exists only in the prototype?
- Which behaviour has been validated?
- Which systems depend upon this rule?
- What would break if this rule changed?
- What tests validate the rule?

The architecture should optimise for these questions.

---

# 20. Repository Architecture Objective

OpenCode's task is **not** simply to create a conventional game-development folder structure.

It must design a repository architecture that simultaneously supports:

1. **Game design authority**
2. **Formal documentation**
3. **Governance**
4. **Provenance**
5. **System specifications**
6. **Research**
7. **Prototype development**
8. **Simulation**
9. **Testing**
10. **Playtesting**
11. **Implementation**
12. **Audit**
13. **Decision history**
14. **Future technology migration**

The resulting architecture should make it difficult for development activity to accidentally bypass the documentation/governance system.

---

# 21. Important Architectural Principle: Documentation Must Control Development

The project specifically wants the **Tiwas TTRPG formal documentation philosophy** adapted to TCIndustries.

The intended principle is:

> **Development is controlled by documented design authority rather than documentation merely describing development after the fact.**

This means the repository should make a distinction between:

```text
AUTHORITATIVE DESIGN
```

and:

```text
IMPLEMENTATION / EVIDENCE
```

Implementation should trace back to design.

Evidence should trace forward to validation.

Design changes should be traceable through formal decisions.

---

# 22. Recommended Architectural Properties

OpenCode should investigate and propose an architecture possessing the following properties.

### 22.1 Single authoritative location for each rule

Avoid duplicated canonical rules across multiple documents.

### 22.2 Stable identifiers

Requirements, rules, decisions, systems, tests, and constraints should have stable identifiers where appropriate.

### 22.3 Explicit status

Documents and individual design statements should make authority visible.

### 22.4 Explicit provenance

Important decisions should identify their origin.

### 22.5 Cross-referencing

Related systems should reference canonical identifiers rather than copying entire rules.

### 22.6 Historical preservation

Superseded material should remain recoverable without becoming current authority.

### 22.7 Prototype separation

Prototype code and experimental behaviour must remain visibly non-canonical.

### 22.8 Validation linkage

Major rules should eventually be connected to tests, simulations, or playtest evidence.

### 22.9 Human approval gates

Certain transitions should require explicit human confirmation.

### 22.10 Technology independence

Game-design authority should survive changes of engine, programming language, prototype, or production technology.

---

# 23. Current Design Maturity

TCIndustries is **not yet a fully specified game**.

The project has a relatively mature high-level vision and governance model but intentionally incomplete mechanical specifications.

Important unresolved systems include, among others:

- skill trees and specialisation budgets;
- resource spawning and lifecycle mathematics;
- resource attribute mathematics;
- crafting experimentation;
- manufacturing and automation;
- transportation and logistics;
- city formation and governance;
- combat;
- currency and economic sinks;
- organisation governance;
- item durability and repair;
- service professions;
- reputation and contracts;
- first-playable scope.

These unresolved areas must remain visibly unresolved.

OpenCode should not fill these gaps merely because an architecture template expects them to contain values.

---

# 24. Explicit Instruction to OpenCode

Use this document as **project-context input**, not as an instruction to invent game mechanics.

Your immediate task is to analyse the existing TCIndustries repository and create a **targeted project architecture** appropriate to the project's actual nature.

The architecture should answer:

1. What documentation domains should exist?
2. Which documents are authoritative?
3. How should governance control design changes?
4. How should system specifications relate to the Master GDD?
5. How should prototypes be separated from canonical design?
6. How should research and historical SWG material be isolated?
7. How should simulations and playtests connect to design validation?
8. How should implementation trace back to approved design?
9. How should future LLMs recover project state?
10. How should the repository support eventual transition from prototype technology to production technology?
11. What should belong in `AGENTS.md`?
12. What should belong in dedicated governance/design documents instead?
13. What automated checks or conventions could prevent authority/status drift?
14. What architecture best supports long-term AI-assisted development without allowing LLM-generated content to silently become canonical?

Do **not** assume that the existing repository structure is correct.

Do **not** redesign the game.

Do **not** resolve TBD game-design questions.

Do **not** promote proposals into canonical rules.

Do **not** treat prototype implementation as authoritative.

Do **not** discard historical/provenance material merely because it is not canonical.

Instead, produce an architecture that allows the project to develop while preserving these distinctions.

---

# 25. Expected Architectural Deliverable

The resulting OpenCode architecture proposal should identify at minimum:

### A. Repository Structure

A proposed directory/file structure with the purpose of each major area.

### B. Documentation Authority Model

A clear hierarchy showing which artifacts control which kinds of decisions.

### C. OpenCode Instruction Model

Determine what belongs in:

- root `AGENTS.md`;
- subordinate `AGENTS.md` files;
- `opencode.json` instructions, if appropriate;
- formal project documentation.

### D. Design Workflow

Define the lifecycle from:

**question → research → proposal → decision → specification → implementation → validation → audit**

### E. Traceability Model

Define how design identifiers, decisions, implementation, tests, simulations, and playtests should reference one another.

### F. Governance Gates

Identify where human approval is required.

### G. Prototype / Production Boundary

Define how experimental implementations can inform the project without becoming accidental design authority.

### H. Documentation Maintenance

Define how OpenCode should detect:

- stale documents,
- contradictory rules,
- duplicate authority,
- orphaned decisions,
- missing provenance,
- status drift,
- implementation without design authority,
- canonical rules lacking validation.

### I. Migration Strategy

Ensure that the architecture remains useful if the project moves between:

- prototype technologies,
- engines,
- languages,
- frameworks,
- or implementation architectures.

---

# 26. Final Architectural Principle

The fundamental architectural requirement for TCIndustries is:

> **The repository must function as a controlled development system in which game design authority, project governance, implementation, evidence, and historical knowledge are explicitly separated but systematically connected.**

The goal is not merely to build a game.

The goal is to build the game **without losing control of what the game is supposed to be** as development, experimentation, AI-assisted work, and multiple implementation generations occur.

OpenCode should therefore optimise the project architecture for:

**authority → traceability → controlled change → implementation → evidence → validation → human decision.**

That principle should take precedence over convenience when the two conflict.

---

# 27. Source Basis

This description is based primarily on the current TCIndustries documentation baseline:

- `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`
- `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md`
- `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md`
- `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`
- `TCIndustries_Master_GDD_v1.0_Consolidated.md`

The v1.1.1 patch explicitly establishes itself as the current working canonical design reference and records human authority rulings.

The Authority & Provenance Reconciliation Matrix establishes the distinction between explicit human/project authority, derived constraints, proposals, historical references, and prototype evidence.

The project documentation also explicitly identifies the prototype as evidence rather than design authority and calls for dedicated future specifications for major systems.

---

# 28. Status of This Document

**Status:** NON-CANONICAL  
**Function:** Architectural planning input  
**Game-design authority:** None  
**Implementation authority:** None  
**Human approval required:** Yes, if any statement in this document is subsequently adopted as a project governance rule.

**Author:** OpenAI GPT-5.6 Luna  
**Date:** 2026-09-12