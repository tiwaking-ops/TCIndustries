# TCIndustries — Tiwakings Craftworld Industries

**Tiwakings Craftworld Industries (TCIndustries)** is an original persistent-world sandbox MMORPG design project centred on:

- player identity and reputation;
- player interdependence;
- player-created society;
- discovery and resource scarcity;
- deep crafting and manufacturing;
- player-driven commerce;
- professions rather than rigid classes;
- meaningful non-combat careers;
- emergent gameplay.

The project is inspired by the systemic philosophy of early *Star Wars Galaxies*, but it is an **original intellectual property**. No external franchise lore, names, characters, factions, or protected content may enter canonical design.

## What this repository is

This repository is the **controlled documentation environment** for the project. It is a documentation-governance and design-systems repository, not a general game-development monorepo. The goal is to preserve design authority, provenance, status discipline, and an auditable decision history as the project evolves.

**Current working canonical design reference:** `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`

## Repository structure

| Directory | Purpose |
|---|---|
| `canonical/` | Current authoritative TCIndustries design (`canonical/gdd/` holds Master GDD versions) |
| `governance/` | Rules controlling documentation authority, status, and change control |
| `proposals/` | Ideas that have not become canonical |
| `investigations/` | Research, audits, comparisons, design investigations |
| `sources/` | Original/reference material and provenance |
| `archive/` | Superseded material retained for history |
| `AGENTS.md` | Operating instructions for AI agents working in this repository |

## Status discipline

Material in this repository carries status. Presence in `canonical/` **does not by itself** prove human approval. See `governance/` and `AGENTS.md` for the authoritative status and authority rules.

## Current governance state

- The Economy/Interdependence Core (EIC) Architecture C has been falsified and retired (HD-EIC-05).
- Any new EIC shaping requires human-supplied structural boundaries (HD-EIC-07).
- Provenance/heavy identity systems remain deferred until explicitly authorised.
- No numeric tuning of economic or progression systems without explicit human authorisation.

This repository was established 2026-09-12 with an intentionally minimal architecture. It will expand incrementally after a documentation audit identifies genuine needs.