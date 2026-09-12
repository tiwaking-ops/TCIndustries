# TCIndustries — Master Game Design Document
## Version 1.1 — Canonical Baseline

**Project:** Tiwakings Craftworld Industries (TCIndustries)  
**Document Title:** Master Game Design Document  
**Version:** 1.1 — Canonical Baseline  
**Baseline Date:** 2026-08-25  
**Document Purpose:** Working canonical design reference for TCIndustries. This version applies authority/status corrections identified in the Canonical Audit of the v1.0 Consolidated Master GDD. No new game-design decisions are introduced.  
**Authority Statement:** This document is the current working canonical design reference for TCIndustries. It supersedes the v1.0 Consolidated Master GDD. Status tags are authoritative. No status is silently promoted. Derived constraints are explicitly distinguished from human-locked decisions.  

**Lineage:**
- Source consolidation: `TCIndustries_Master_GDD_v1.0_Consolidated.md`
- Controlling audit: `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md`
- Original first-pass sources remain as provenance only.

---

# 1. Document Control

| Field | Value |
|---|---|
| Project | TCIndustries (Tiwakings Craftworld Industries) |
| Document | Master Game Design Document |
| Version | 1.1 — Canonical Baseline |
| Status | Working Canonical Reference (status/authority pass complete) |
| Scope | Game design and rules only. Technical implementation belongs in a future Technical Design Document (TDD). |
| Change Control | Status changes require explicit project/human approval. Derived constraints may be refined when parent principles are refined. |
| Prior Version | 1.0 — Consolidated Master GDD |
| Controlling Audit | TCIndustries_Master_GDD_v1.0_Canonical_Audit.md |

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

| Status | Meaning |
|---|---|
| **LOCKED** | Explicit project/human decision; authoritative unless formally revised. Also called HUMAN-LOCKED where distinction from derived material is needed. |
| **DERIVED CONSTRAINT** | Binding consequence of one or more LOCKED principles. Not an independent human decision. May be refined when parent principles are refined. Must not be treated as LOCKED. |
| **PROPOSED** | Recommended design direction that has not been explicitly approved. |
| **TBD** | Genuine unresolved design decision. |
| **PROTOTYPE** / **EVIDENCE** | Behaviour demonstrated or explored by prototype implementation; not canonical merely because it exists. |
| **DEFERRED** | Recognized future system intentionally postponed. |
| **ASSUMPTION** | Working assumption used by one or more sources but not established as fact or decision. |
| **HISTORICAL** | Historical/reference information (particularly from SWG) that is not itself a TCIndustries rule. |

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
**Status: LOCKED**  
**Authority:** Explicit project brief / consistent across all sources as core objective.

TCIndustries is a persistent multiplayer virtual world in which players create their own economic, social, occupational, and reputational identities.

The world should feel inhabited not only by NPCs and systems, but by player-created institutions: shops, workshops, supply networks, social venues, cities, organisations, specialist communities, commercial districts, brands, and player reputations.

The world should appear to continue existing independently of any particular player’s story. Players participate in the world rather than being the centre of it.

The game’s primary content is the interaction between players, world systems, and player-created consequences.

## VIS-002 — Citizen Rather Than Chosen Hero
**Status: LOCKED**  
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
**Status: LOCKED**  
**Authority:** Explicit project direction across sources.

Players should be major participants in the production and distribution of economic goods and services. The economy must be materially connected to the world. Resources, labour, transport, crafting, consumption, degradation, services, and regional differences should matter.

## VIS-004 — Interdependence
**Status: LOCKED**  
**Authority:** Explicit project pillar.

Important professions should depend upon one another. A complete economy should contain useful roles that cannot all be replaced by a single generalist player.

## VIS-005 — Emergence
**Status: LOCKED**  
**Authority:** Explicit project direction.

The game should create conditions from which interesting events emerge rather than attempting to script every important player experience.

Desired player thought: “This world changed because of what players did, and other players know who matters here.”

## VIS-006 — Anti-Goals
**Status: LOCKED**  
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
**Status: LOCKED**  
**Authority:** Explicit project pillar.

Resources are temporary, geographically distributed, variable in quality, and economically meaningful.

Exceptional resource discoveries should create player-driven events: prospecting activity, mining camps, trade opportunities, speculation, guild coordination, regional traffic, competitive supply chains, and demand for extractors, transport, processing, and manufacturing.

A resource discovery should be an **event**, not merely an inventory update.

**Historical SWG Inspiration (HISTORICAL):** Original SWG used temporary resource spawns with variable statistics, regional availability, and resource surveying.

**TCIndustries Adaptation (PROPOSED — RES-001):** Retain the principle of dynamic, temporary, statistically variable resource occurrences, while improving information clarity, anti-monopoly controls, and long-term economic observability.

## PIL-002 — Identity and Reputation
**Status: LOCKED**  
**Authority:** Explicit project pillar.

Player identity must be meaningful, visible, and socially consequential.

Products, services, businesses, and buildings should preserve meaningful association with their creator, provider, organisation, or brand.

Reputation should emerge from player behaviour and economic history rather than merely from an abstract reputation score. Ideal outcome: players seek out products made by a particular player or organisation.

## PIL-003 — Interdependence
**Status: LOCKED**  
**Authority:** Explicit project pillar.

No major profession category should be fully self-sufficient at high effectiveness.

Interdependence should arise through: distinct skills, specialised resources, intermediate components, service needs, regional logistics, knowledge gaps, reputation, production capacity limits, and social organisation.

Interdependence must create opportunity rather than coercive inconvenience.

## PIL-004 — Player-Created Society
**Status: LOCKED**  
**Authority:** Explicit project pillar.

Players must be able to form organisations, businesses, settlements, social hubs, trade networks, and reputational institutions with meaningful world presence.

---

# 7. Player Experience and Core Gameplay Loops

## LOOP-001 — Interconnected Loops
**Status: PROPOSED**

The game supports several interconnected loops rather than a single mandatory progression path.

| Loop | Core Activities | Social Outcome |
|---|---|---|
| Discovery | Survey, explore, identify resources | Information trade, rushes, regional knowledge |
| Extraction | Harvest, operate extraction equipment, gather | Resource supply and logistics |
| Processing | Refine, convert, prepare materials | Intermediate-goods markets |
| Crafting | Build items, experiment, specialise | Brand identity and product differentiation |
| Retail | Stock vendors, price goods, market services | Local commerce and commercial districts |
| Service | Heal, entertain, repair, transport, construct | Recurring player interaction |
| Combat | Fight threats, protect territory, acquire rewards | Demand for equipment, medicine, consumables |
| Social Governance | Form associations, cities, events, policies | Community identity and institutions |

## LOOP-002 — Non-Combat Viability
**Status: LOCKED**  
**Authority:** Explicit project direction.

A player must be able to pursue a meaningful long-term career without engaging in combat as their primary activity.

Combat may offer demand, risk, materials, access, and opportunities, but must not be the sole source of progression, wealth, or social status.

## LOOP-003 — Progression Philosophy
**Status: LOCKED**  
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
**Status: LOCKED**  
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
**Status: DERIVED CONSTRAINT**  
**Derives from:** PIL-003 (Interdependence), PROF-001 / specialisation principles

Crafting depth must reinforce specialisation rather than allow one character to efficiently dominate all high-value production.

---

# 13. Manufacturing and Automation

## MFG-001 — Controlled Automation
**Status: LOCKED**  
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
**Status: DERIVED CONSTRAINT**  
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
**Status: DERIVED CONSTRAINT**  
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
**Status: DERIVED CONSTRAINT**  
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
**Status: DERIVED CONSTRAINT**  
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

| ID | Question | Affected Systems | Why It Matters | Current Status |
|---|---|---|---|---|
| OQ-001 | Exact specialisation budget / skill-cap model? | Skills, Professions, Interdependence | Prevents single-character mastery of all roles | TBD |
| OQ-002 | Resource spawn rates, durations, and depletion mathematics? | Resources, Economy, Discovery | Core scarcity and event generation | TBD |
| OQ-003 | Exact attribute sets for resources and mapping to product outcomes? | Resources, Crafting | Product differentiation and experimentation | TBD |
| OQ-004 | Crafting experimentation model details? | Crafting, Identity | Crafter skill expression and reputation | TBD |
| OQ-005 | Manufacturing/factory throughput, maintenance, and ownership rules? | Manufacturing, Automation | Anti-automation-drift enforcement | TBD |
| OQ-006 | Transportation model and fast-travel constraints? | World, Economy, Logistics | Geographic meaning vs. friction | TBD |
| OQ-007 | City formation, governance, taxation/maintenance models? | Cities, Society, Buildings | Player-created society pillar | TBD |
| OQ-008 | Combat scope (PvE/PvP), risk model, and progression integration? | Combat, Economy, Services | Demand generation without forced path | TBD |
| OQ-009 | Currency system(s) and money sinks? | Economy | Economic stability | TBD |
| OQ-010 | Item durability, decay, and repair model? | Items, Services, Economy | Service demand and item identity | TBD |
| OQ-011 | Respecialisation exact costs and rules? | Progression, Identity | Balance of flexibility vs. commitment | TBD |
| OQ-012 | Multi-accounting and economic exploit policies? | Economy, Safety | Fairness and stability | TBD |
| OQ-013 | Organisation rights, hierarchy, and property model? | Social Systems | Collective ownership | TBD |
| OQ-014 | Species/ancestry/starting location design? | Character | Setting and identity expression | DEFERRED |
| OQ-015 | Which systems are required for first playable prototype vs. persistent alpha? | Production planning | Scope control | TBD |
| OQ-016 | Exact vendor and retail mechanics? | Retail, Economy | Player commerce viability | TBD |
| OQ-017 | Provenance depth and visibility rules? | Identity, Crafting, Items | Reputation pillar support | TBD |

---

# 32. Assumption Register

| ID | Assumption | Source | Status | Consequence if Wrong |
|---|---|---|---|---|
| AS-001 | A functional prototype (Seed-2.1 or similar) exists or will exist for testing | Multiple sources | ASSUMPTION | Testing approach and reconciliation process must be redesigned |
| AS-002 | Data-driven content expansion is preferred over hard-coded systems | Consistent across sources | ASSUMPTION / PROPOSED | Content pipeline and technical architecture change |
| AS-003 | Regional economic differentiation is desirable and feasible | Consistent proposals | ASSUMPTION | World design and transport systems simplify or change |
| AS-004 | Meaningful provenance tracking is technically and design-feasible at scale | Identity pillar proposals | ASSUMPTION | Reputation systems may need alternative approaches |
| AS-005 | Non-combat careers can generate sufficient player demand without combat gating | Core vision | ASSUMPTION | Demand generation systems may need redesign |

---

# 33. Status / Decision Register (Selected Key Rules)

| ID | Rule Summary | Status | Authority / Evidence Class | Dependencies | Notes |
|---|---|---|---|---|---|
| VIS-001 | Persistent player-driven virtual society | LOCKED | Explicit project objective | All systems | Core |
| VIS-002 | Citizen identity over hero narrative | LOCKED | Explicit player fantasy | Progression, social, economy | Core |
| VIS-003 | Player-driven economy | LOCKED | Explicit project direction | Economy, production | Core |
| VIS-004 | Interdependence of professions | LOCKED | Explicit pillar | Skills, economy, services | Core |
| VIS-005 | Emergence over scripting | LOCKED | Explicit project direction | All systems | Core |
| VIS-006 | Anti-goals (theme-park, class, automation, etc.) | LOCKED | Explicit anti-drift | All systems | Core |
| PIL-001 | Discovery / Gold Rush | LOCKED | Explicit pillar | Resources, economy | Core |
| PIL-002 | Identity and Reputation | LOCKED | Explicit pillar | Crafting, retail, services | Core |
| PIL-003 | Interdependence | LOCKED | Explicit pillar | Skills, economy | Core |
| PIL-004 | Player-Created Society | LOCKED | Explicit pillar | Buildings, cities, orgs | Core |
| LOOP-002 | Non-combat long-term viability | LOCKED | Explicit direction | Professions, economy | Core |
| LOOP-003 | Progression expands options | LOCKED | Explicit philosophy | Skills | Core |
| WRLD-001 | Persistent shared world | LOCKED | Required by vision | World, ownership | Core |
| PLR-001 | Persistent character identity | LOCKED | Required by pillars | Player systems | Core |
| PLR-003 | Persistent ownership | LOCKED | Required by society/economy | Buildings, items | Core |
| PROF-001 | Flexible skill-based professions | LOCKED | Explicit requirement | Progression | Core |
| PROF-004 | Respecialisation allowed | LOCKED | Explicit philosophy | Progression | Core |
| RES-002 | Resources create discovery/scarcity | LOCKED | Pillar implementation | Economy, crafting | Core |
| CRFT-001 | Differentiated products | LOCKED | Identity/reputation requirement | Resources, skills | Core |
| MFG-001 | Controlled automation | LOCKED | Anti-drift | Manufacturing | Core |
| MFG-004 | Anti-Factorio drift | DERIVED CONSTRAINT | Derives from VIS-006, MFG-001 | Manufacturing | Reclassified in v1.1 |
| ECO-001 | Player-driven economy | LOCKED | Explicit | Production, retail | Core |
| ECO-003 | No pure spreadsheet optimisation | DERIVED CONSTRAINT | Derives from VIS-006, VIS-003 | Economy | Reclassified in v1.1 |
| RET-001 | Player retail | LOCKED | Explicit | Economy, identity | Core |
| SERV-001 | Meaningful non-combat services | LOCKED | Explicit | Interdependence | Core |
| CMBT-001 | Combat as one lifestyle among many | LOCKED | Explicit | Economy, services | Core |
| CMBT-003 | No Forced Combat Path | DERIVED CONSTRAINT | Derives from LOOP-002, VIS-006 | Combat | Reclassified in v1.1 |
| BLD-001 | Player-owned structures | LOCKED | Explicit | Society, economy | Core |
| SOC-001 | Player organisations | LOCKED | Explicit | Society | Core |
| PROG-002 | Multiple viable paths | DERIVED CONSTRAINT | Derives from LOOP-002 | Progression | Reclassified in v1.1 |
| CRFT-006 | No Instant Mastery of Everything | DERIVED CONSTRAINT | Derives from PIL-003, specialisation principles | Crafting | Reclassified in v1.1 |
| EMRG-001 | Conditions for emergence | LOCKED | Explicit | All systems | Core |
| SAFE-001 | Ownership/transfer safeguards | LOCKED | Explicit | Ownership | Core |
| SWG-001/002 | Inspiration only; no SW IP | LOCKED | Explicit | Setting | Core |
| PROT-001 | GDD over prototype | LOCKED | Explicit principle | All | Core |
| Most detailed taxonomies, formulas, trees, rates | PROPOSED or TBD | LLM recommendations | Various | Not promoted |

---

# 34. Recommended Design Sequence

This sequence reflects the controlling Canonical Audit. It replaces the linear list from v1.0.

**Completed in this version (v1.1):**
- Authority/status hygiene pass (HUMAN-LOCKED vs DERIVED CONSTRAINT distinction).

**Next required work (process, not design expansion):**
0. **Evidence Reconciliation Pass** — Inspect any existing prototype; produce GDD ↔ prototype ↔ decision evidence matrix. (Open process requirement.)

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
2. Provenance & Reputation Engine
3. Commerce & Services
4. Society (organisations, cities, governance, institutions)
5. World & Logistics
6. Combat (designed into the existing economy)
7. Simulation and numeric tuning only after structural shapes exist

No further LLM consolidation of the Master GDD should occur. The next operations are controlled correction (this document), evidence reconciliation, and then actual systems design.

---

# 35. Glossary

| Term | Definition |
|---|---|
| Citizen | Player fantasy of being an ordinary but consequential participant in a living society rather than a chosen hero. |
| Gold Rush | Player-driven economic and social event triggered by exceptional temporary resource discoveries. |
| Interdependence | Design condition in which high-effectiveness roles require other specialised players. |
| Provenance | Association of an item, product, or service with its creator, organisation, or resource origins. |
| Profession | Emergent identity arising from skills, specialisation, equipment, reputation, and behaviour — not a rigid class. |
| Specialisation Budget | Constraint (skill points, capacity, or equivalent) preventing one character from mastering all economically important roles. |
| Status Tag | Explicit marker (LOCKED, DERIVED CONSTRAINT, PROPOSED, TBD, etc.) controlling the authority of a design statement. |
| DERIVED CONSTRAINT | Binding consequence of LOCKED principle(s); not an independent human decision. |

---

# 36. Provenance and Change Notes

## v1.1 Canonical Baseline — Change Scope

This version applies **only** the authority/status corrections approved in the Canonical Audit of v1.0. No new game-design decisions, mechanics, numbers, or system expansions were introduced.

### Status reclassifications (LOCKED → DERIVED CONSTRAINT)

| ID | Rule | Parent Principle(s) |
|---|---|---|
| CRFT-006 | No Instant Mastery of Everything | PIL-003, specialisation principles |
| MFG-004 | Anti-Factorio Drift | VIS-006, MFG-001 |
| ECO-003 | No Pure Spreadsheet Optimisation | VIS-006, VIS-003 |
| CMBT-003 | No Forced Combat Path | LOOP-002, VIS-006 |
| PROG-002 | Multiple Viable Paths | LOOP-002 |

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
