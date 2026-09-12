**Project Description for OpenCode — TCIndustries Documentation Control Architecture**

Use the text below as the primary project description / `AGENTS.md` seed (or equivalent project brief) when initialising OpenCode on this repository. It is written so OpenCode can generate a targeted, disciplined project architecture focused on **controlling game development documentation**, not on building the game client/server itself.

---

### Project Identity

**Name:** TCIndustries (full formal name: Tiwakings Craftworld Industries)  
**Type:** Persistent-world sandbox MMORPG design project  
**Primary artefact:** Master Game Design Document (GDD) and supporting design packages  
**Current phase:** Design integrity, authority control, and evidence-based systems architecture (pre-implementation)

TCIndustries is a persistent sandbox MMORPG centred on:

- Player identity and reputation
- Player interdependence
- Player-created society
- Discovery and resource scarcity
- Deep crafting and manufacturing
- Player-driven commerce
- Professions rather than rigid classes
- Meaningful non-combat careers
- Emergent gameplay

It draws philosophical inspiration from the systemic depth of early *Star Wars Galaxies* but is an **original IP**. No Star Wars lore, names, characters, factions, or protected content may ever enter canonical design.

### Core Objective for OpenCode

Create and maintain a **formal documentation control system** that treats the game design as a governed engineering artefact. The system must enforce:

1. Strict separation of **LOCKED** (human-approved) decisions from **PROPOSED**, **TBD**, **DERIVED CONSTRAINT**, **PROTOTYPE/EVIDENCE**, **DEFERRED**, and **ASSUMPTION** material.
2. Explicit authority and provenance tracking (which human ruling or source document owns each statement).
3. No silent invention or promotion of canon.
4. Traceability of every significant rule via stable identifiers.
5. Cross-system dependency awareness and falsifiability where systems claims are made.
6. Protection of the four design pillars as binding constraints on all downstream work.

This is **not** a general game-development monorepo. It is a documentation-governance and design-systems repository. Implementation code, engine work, and technical design documents are out of scope until explicitly authorised.

### Design Pillars (Binding Constraints)

Treat these as non-negotiable system constraints:

1. Discovery and the Gold Rush
2. Identity and Reputation
3. Interdependence
4. Player-Created Society

Supporting commitments: sandbox gameplay, meaningful non-combat careers, no forced combat path, player-driven economy, player-created identity.

### Documentation Status Discipline (Mandatory)

Every significant design element must carry an explicit status. Status meanings are authoritative:

- **LOCKED / HUMAN-LOCKED** — Explicit human/project decision. Treat as law. Do not reinterpret or casually modify.
- **DERIVED CONSTRAINT** — Logical consequence of one or more LOCKED principles. Binding but not an independent human decision. May be refined only when parent principles change.
- **PROPOSED** — Candidate design. May be analysed, challenged, improved, rejected, or replaced. Never treat as canon.
- **TBD** — Genuine unresolved decision. Do not invent an answer to fill gaps.
- **PROTOTYPE / EVIDENCE** — Observed behaviour from prototypes. Evidence only; not design authority.
- **DEFERRED** — Intentionally postponed. Do not pull into scope without explicit authorisation.
- **ASSUMPTION** — Working assumption that has not been elevated to decision.
- **HISTORICAL** — Reference material (including SWG-inspired notes). Never automatically TCIndustries rules.

**Critical rule:** Never silently promote any non-LOCKED status into LOCKED. Status changes require explicit, recorded human ruling.

### Authority Model

1. Explicit human rulings have highest authority.
2. The current Master GDD (latest status-patched version) is the working canonical design reference unless a newer document has formally superseded it.
3. When documents conflict, surface the conflict; do not silently reconcile.
4. Prototype behaviour is evidence, not authority. GDD rules govern unless a later human ruling changes the relationship.
5. Historical inspiration is never automatic canon.
6. Older first-pass mechanical proposals are non-canonical unless re-confirmed.

### Governance Rules Already in Force

- No numeric tuning of economic or progression systems without explicit authorisation.
- No invention of new residual channels or architectures without human-defined structural boundaries.
- No automatic creation of “Architecture D” or equivalent after falsification of a candidate.
- Provenance / identity systems remain deferred until authorised.
- Economic Interdependence Core (EIC) work is currently gated: Architecture C has been falsified and retired; any new shaping requires human-supplied structural boundaries first.

### Expected Project Architecture OpenCode Should Produce

OpenCode should design a clean, minimal repository / folder architecture optimised for **documentation control**, not for game engine code. Recommended high-level structure (refine as needed):

```
/
├── AGENTS.md                          # Project instructions for AI agents (this brief + operating rules)
├── docs/
│   ├── canonical/                     # Only documents with current working-canonical status
│   │   └── Master_GDD/                # Versioned Master GDD + status patches
│   ├── authority/                     # Human rulings, provenance matrices, decision gates
│   ├── investigations/                # Design investigations, candidates, falsification passes
│   ├── simulations/                   # Simulation specifications and results
│   ├── provisional/                   # Functional shapes, candidates under test
│   └── archive/                       # Superseded or historical documents
├── rulings/                           # Chronological human decision records
├── templates/                         # Document templates that enforce status tags, identifiers, authority statements
├── scripts/                           # Lightweight validation / consistency checkers (status hygiene, identifier uniqueness, etc.)
└── README.md                          # Human-facing orientation
```

Key architectural requirements:

- Clear separation of **canonical** vs **working/proposed** material.
- Every design document must declare: Project, Title, Version, Status, Authority Statement, Lineage.
- Stable rule identifiers (e.g. CRFT-001, PIL-003) preserved across versions.
- Human Decision records (HD-*) treated as first-class artefacts.
- Support for progressive refinement and falsification workflows (candidate → adversarial test → human gate).
- Minimal tooling surface. Prefer plain Markdown + disciplined conventions over heavy frameworks unless a clear need is demonstrated.

### Operating Constraints for Any AI Agent (including OpenCode)

- Act as senior Game Systems Architect + Design Auditor + Documentation Architect.
- Preserve design integrity over volume of output.
- Surface missing information rather than inventing it.
- Label proposals, assumptions, and recommendations explicitly.
- Never convert suggestions into decisions.
- When a design choice is required to proceed, mark it PROPOSED or ASSUMPTION.
- Respect current human gates (especially post-Architecture-C EIC boundaries).
- Optimise for correct, useful, auditable design — not for the largest possible documentation set.

### Immediate Next Human Gate

Any further Economic Interdependence Core architectural work is blocked until the human supplies structural boundaries for bounded new shaping (per HD-EIC-07). Do not invent those boundaries.

---

This description is deliberately self-contained so OpenCode can bootstrap a targeted documentation-control architecture without needing the full conversation history. It encodes the existing formal system already operating on the project (status discipline, authority provenance, human ruling gates, pillar constraints, and anti-canon-invention rules) and asks OpenCode to operationalise it as repository structure, templates, and agent behaviour.