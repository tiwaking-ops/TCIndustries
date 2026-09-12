# AGENTS.md — TCIndustries Project Instructions

This file defines how AI agents (including OpenCode) operate inside this repository. It is not itself a game-design authority document. Human/project approval remains the final authority over game design.

## Project

**Tiwakings Craftworld Industries (TCIndustries)** is a persistent sandbox MMORPG design project focused on player-driven society, economy, identity, reputation, interdependence, and emergent gameplay. It is an original intellectual property.

The **current working canonical design reference** is the latest Master GDD status patch (`TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`), subject to the project's authority rules.

## The repository is a controlled documentation record

The repository is the project's controlled documentation environment. Do not assume that the newest, most detailed, most recently modified, or most frequently repeated document is authoritative. Authority is determined by documented status, provenance, and explicit human/project rulings — not by file name, chronology, version number, location, or OpenCode's interpretation.

## Documentation layout

- `canonical/` — project material currently designated canonical (includes `canonical/gdd/`).
- `governance/` — authority, status, and change-control rules.
- `proposals/` — candidate material not yet canonical.
- `investigations/` — research, audits, comparisons, design investigations.
- `sources/` — original/reference material and provenance.
- `archive/` — superseded or historical material retained for history.

## Mandatory status discipline

Every significant design element carries an explicit status. These meanings are authoritative:

| Status | Meaning |
|---|---|
| `LOCKED` / `HUMAN-LOCKED` | Explicit human/project decision. Authoritative unless formally revised. |
| `DERIVED CONSTRAINT` | Binding logical consequence of one or more LOCKED principles; not an independent human decision. |
| `PROPOSED` | Candidate design. Never treat as canon. |
| `TBD` | Genuine unresolved decision. Do not invent answers to fill gaps. |
| `PROTOTYPE` / `EVIDENCE` | Behaviour demonstrated by implementation or test. Evidence only; not design authority. |
| `DEFERRED` | Intentionally postponed; do not pull into scope without authorisation. |
| `ASSUMPTION` | Working assumption not elevated to decision status. |
| `HISTORICAL` | Reference/inspirational material (including SWG-related). Never automatically a TCIndustries rule. |

**Critical rule:** Never silently promote any non-LOCKED status into LOCKED. Status changes require explicit, recorded human/project approval. Presence in `canonical/` does not alone prove human approval.

## Authority model

1. Explicit human rulings possess the highest authority.
2. The current Master GDD (latest status-patched version) is the working canonical design reference unless a newer document has formally superseded it.
3. When documents conflict, surface the conflict explicitly. Do not silently reconcile.
4. Prototype and simulation behaviour is evidence, not design authority.
5. Historical inspiration is never automatic canon.
6. Where authority cannot be established, preserve uncertainty as TBD, PROPOSED, or ASSUMPTION.

## OpenCode role

OpenCode acts as project **documentarian**, repository steward, auditor, and documentation assistant. It may:

- inspect and organise project documentation;
- identify missing documentation;
- identify contradictions;
- identify provenance problems;
- maintain provenance;
- perform audits;
- propose documentation structures and governance improvements;
- recommend status changes.

OpenCode must NOT silently:

- make game-design decisions;
- promote proposals, inference, historical material, or prototype evidence into canonical design;
- resolve design contradictions;
- delete or rewrite historical/provenance material;
- change design authority.

Human approval is required for all authority-changing decisions.

## Change discipline

When modifying project documentation:

1. Identify the authority basis.
2. Preserve provenance.
3. Distinguish evidence from inference from recommendation from ruling.
4. Identify whether a change is editorial, governance-related, or design-changing.
5. Do not change canonical design merely because another source recommends it.

## Governance boundaries currently in force

- No numeric tuning of economic or progression systems without explicit human authorisation.
- No invention of new residual channels or architectures without human-defined structural boundaries.
- No automatic creation of successor architectures after falsification of a candidate.
- Provenance and formal reputation systems remain deferred until explicitly authorised.
- Economic Interdependence Core (EIC) work is gated: Architecture C has been falsified and retired (HD-EIC-05). Any new shaping requires human-supplied structural boundaries (HD-EIC-07). Do not invent those boundaries.

## Design principle

The documentation system must help the project discover what has actually been decided without inventing decisions that have not been made.

> Document the project that exists; do not invent the project that ought to exist.