# TCIndustries — OpenCode Project Architecture Brief
## Formal Project Documentation Report

**Project:** Tiwakings Craftworld Industries (TCIndustries)  
**Document Title:** OpenCode Project Architecture Brief  
**Document Type:** Formal Project Documentation Report (Non-Canonical)  
**Version:** 0.1  
**Date:** 2026-09-12  
**Author:** Grok 4.5 (xAI)  
**Status:** PROPOSED — Working brief for OpenCode initialisation only. Creates no design authority.  
**Authority Statement:** This document creates no game-design decisions and no canonical rules. It exists solely to provide OpenCode with a controlled, governance-aware project description so that any generated project architecture remains subordinate to existing human rulings, status discipline, and design pillars.  
**Scope:** Documentation control architecture and agent operating constraints for the TCIndustries design repository. Implementation code, engine work, and Technical Design Documents are out of scope.  

**Lineage / Controlling References:**  
- `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`  
- `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`  
- `TCIndustries_EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md`  
- `AGENTS.md` (project instructions)  
- Prior OpenCode project description (conversational, non-formal)  

**Purpose:**  
Provide OpenCode with a single, formally structured project description that enables it to generate a targeted repository architecture optimised for rigorous game-design documentation control, authority tracking, and status hygiene. This brief encodes the existing TCIndustries formal documentation standards so that any architecture produced by OpenCode remains consistent with project governance.

---

# 1. Document Control

| Field | Value |
|---|---|
| Project | TCIndustries (Tiwakings Craftworld Industries) |
| Document | OpenCode Project Architecture Brief |
| Version | 0.1 |
| Status | PROPOSED |
| Author | Grok 4.5 (xAI) |
| Date | 2026-09-12 |
| Change Control | Any elevation of status or content into canonical use requires explicit human approval |
| Scope Restriction | Documentation governance and agent behaviour only |

---

# 2. Project Identity

**Full Formal Name:** Tiwakings Craftworld Industries  
**Short Name:** TCIndustries  

TCIndustries is a persistent-world sandbox MMORPG centred on:

- player identity and reputation;
- player interdependence;
- player-created society;
- discovery and resource scarcity;
- deep crafting and manufacturing;
- player-driven commerce;
- professions rather than rigid classes;
- meaningful non-combat careers;
- emergent gameplay.

The project is inspired by the philosophy and systemic depth of early *Star Wars Galaxies*, but TCIndustries is an original intellectual property. No Star Wars-specific lore, names, characters, factions, terminology, or protected content may enter canonical design.

---

# 3. Primary Objective for OpenCode

OpenCode is to generate and maintain a **formal documentation control architecture** for the TCIndustries design repository.

The architecture must treat game design as a governed engineering artefact. It must enforce:

1. Strict separation of LOCKED / HUMAN-LOCKED decisions from PROPOSED, TBD, DERIVED CONSTRAINT, PROTOTYPE/EVIDENCE, DEFERRED, and ASSUMPTION material.
2. Explicit authority and provenance tracking for every significant statement.
3. Prohibition on silent invention or promotion of canon.
4. Stable rule identifiers and full traceability.
5. Cross-system dependency awareness and falsifiability where systemic claims are advanced.
6. Treatment of the four design pillars as binding constraints on all downstream work.

This is not a general game-development monorepo. It is a documentation-governance and design-systems repository.

---

# 4. Design Pillars (Binding Constraints)

The following pillars are treated as non-negotiable system constraints on all work:

1. Discovery and the Gold Rush  
2. Identity and Reputation  
3. Interdependence  
4. Player-Created Society  

Supporting commitments that must also be preserved:

- sandbox gameplay;
- meaningful non-combat careers;
- no forced combat path;
- player-driven economy;
- player-created identity.

---

# 5. Mandatory Status Discipline

Every significant design element must carry an explicit status. The following meanings are authoritative and must not be collapsed:

| Status | Meaning |
|---|---|
| **LOCKED / HUMAN-LOCKED** | Explicit human or project decision. Authoritative. Must not be casually modified or reinterpreted. |
| **DERIVED CONSTRAINT** | Binding logical consequence of one or more LOCKED principles. Not an independent human decision. May be refined only when parent principles are refined. |
| **PROPOSED** | Candidate design. Not canonical. May be analysed, challenged, improved, rejected, or replaced. |
| **TBD** | Genuine unresolved design decision. Do not invent an answer merely to achieve apparent completeness. |
| **PROTOTYPE / EVIDENCE** | Behaviour demonstrated by prototype or test. Evidence only; does not confer design authority. |
| **DEFERRED** | Intentionally postponed. Must not be pulled into current scope without explicit authorisation. |
| **ASSUMPTION** | Working assumption that has not been elevated to decision status. |
| **HISTORICAL** | Reference or inspirational material. Never automatically a TCIndustries rule. |

**Critical Rule:** Never silently promote any non-LOCKED status into LOCKED. Status changes require explicit, recorded human or project approval.

---

# 6. Authority Model

1. Explicit human rulings possess the highest authority.  
2. The current Master GDD (latest status-patched version) is the working canonical design reference unless a newer document has formally superseded it.  
3. When documents conflict, the conflict must be identified explicitly; silent reconciliation is prohibited.  
4. Prototype behaviour is evidence, not design authority.  
5. Historical inspiration (including early *Star Wars Galaxies*) is never automatic canon.  
6. Where authority cannot be established, uncertainty is preserved as TBD, PROPOSED, or ASSUMPTION.

---

# 7. Active Governance Boundaries

The following constraints are currently in force and must be respected by any architecture or agent behaviour:

- No numeric tuning of economic or progression systems without explicit human authorisation.  
- No invention of new residual channels or architectures without prior human-defined structural boundaries.  
- No automatic creation of successor architectures after falsification of a candidate.  
- Provenance and formal reputation systems remain deferred until explicitly authorised.  
- Economic Interdependence Core (EIC) work is gated: Architecture C has been falsified and retired (HD-EIC-05). Any new shaping requires human-supplied structural boundaries (HD-EIC-07) before further architectural work is authorised.

---

# 8. Required Architectural Characteristics

Any project architecture generated by OpenCode must satisfy the following characteristics:

- Clear separation of canonical material from working, provisional, and archival material.  
- Every design document must declare: Project, Document Title, Version, Status, Authority Statement, and Lineage.  
- Stable rule identifiers retained across versions.  
- Human Decision records (HD-*) treated as first-class artefacts.  
- Support for progressive refinement and falsification workflows (candidate → adversarial test → human gate).  
- Minimal tooling surface. Prefer plain Markdown and disciplined conventions over heavy frameworks unless a demonstrated need exists.  
- Templates that enforce status tags, authority statements, and identifier discipline.

---

# 9. Agent Operating Constraints

Any AI agent operating inside this architecture, including OpenCode itself, must:

- Act as senior Game Systems Architect, Design Auditor, and Documentation Architect.  
- Prioritise design integrity over volume of output.  
- Surface missing information rather than invent answers.  
- Label proposals, assumptions, and recommendations explicitly.  
- Never convert suggestions into decisions.  
- Respect all current human gates and governance boundaries.  
- Optimise for correct, useful, auditable design.

---

# 10. Immediate Human Gate

Further Economic Interdependence Core architectural work remains blocked pending human supply of structural boundaries for any bounded new shaping (HD-EIC-07). No architecture or agent may invent those boundaries.

---

# 11. Document Status Declaration

This document is issued as **PROPOSED**. It possesses no design authority. It may be used by OpenCode to initialise a documentation-control architecture. Any subsequent elevation of its content, structure, or recommendations into project practice requires explicit human confirmation.

**End of Formal Project Documentation Report**