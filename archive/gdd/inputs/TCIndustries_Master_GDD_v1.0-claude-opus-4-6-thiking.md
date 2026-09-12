# Tiwakings Craftworld Industries — Master Game Design Document

**Short Title:** TCIndustries
**Document:** Master Game Design Document
**Version:** 1.0 — Canonical Foundation
**Date:** 2026-08-20
**Purpose:** Canonical game-design and rules specification for TCIndustries.
**Scope:** Game design only. Technical implementation is deferred to a separate Technical Design Document (TDD).
**Author:** Claude Opus 4.6 (per filename, unverified)

---

> [!IMPORTANT]
> This document is the **authoritative design specification** for TCIndustries. Prototype behaviour, historical SWG mechanics, and designer assumptions are explicitly distinguished from locked project decisions. No status may be silently promoted.

---

# Table of Contents

1. [Executive Summary](#1-executive-summary)
2. [Design Status Definitions](#2-design-status-definitions)
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

TCIndustries is a persistent sandbox MMORPG centred on living in a player-driven society. It is inspired heavily by the original *Star Wars Galaxies* virtual-world model: dynamic resources, interdependent professions, player-made goods, recognisable crafters, player commerce, housing, settlements, social services, and emergent economic life.

It is not intended to reproduce Star Wars.

Its purpose is to reproduce and expand upon the underlying experience of:

> **Living in a persistent world as an ordinary but consequential citizen.**

The player is not necessarily a hero. They may become a farmer, miner, prospector, artisan, weaponsmith, armoursmith, architect, doctor, entertainer, merchant, trader, explorer, combatant, researcher, manufacturer, service provider, shop owner, city participant, or another identity created by the game's systems.

The game's central loop is:

> Discover resources → acquire and process materials → craft goods or provide services → specialise → establish reputation → participate in trade, society, and regional economies.

The game's central systems should cause players to depend upon one another for resources, products, services, knowledge, equipment, medical assistance, entertainment, commerce, construction, specialisation, and social organisation.

The ultimate measure of success is not how many systems exist. It is whether those systems combine to produce a world that feels **alive**.

TCIndustries must avoid becoming an automation-first industrial game, a conventional class MMO, a quest-driven theme park, or a spreadsheet economy simulator. Industry, production, combat, and economic simulation exist to create meaningful social roles and player interdependence.

This document is an initial canonical framework. It deliberately identifies many mechanics as **PROPOSED**, **TBD**, or **DEFERRED** rather than inventing premature false precision.

---

# 2. Design Status Definitions

Every significant design element in this document carries one of five statuses.

| Status | Meaning |
|---|---|
| **LOCKED** | An explicit project decision that is authoritative unless formally revised. |
| **PROPOSED** | A design recommendation that appears appropriate but has not been formally approved. |
| **TBD** | A genuine unresolved design decision requiring future discussion. |
| **PROTOTYPE** | A behaviour currently implemented or explored in prototype code but not accepted as canonical. |
| **DEFERRED** | A known system that will be designed later and should not block current development. |

> [!CAUTION]
> Never silently promote PROPOSED, PROTOTYPE, TBD, or DEFERRED material into LOCKED rules. Status changes require explicit documented approval.

The absence of a status tag on ordinary explanatory prose does not override the status of individual mechanics.

---

# 3. Vision

## VIS-001 — Living Virtual World

**Status: LOCKED**

TCIndustries is a persistent multiplayer virtual world in which players create their own economic, social, occupational, and reputational identities.

The world should feel inhabited not only by NPCs and systems, but by player-created institutions: shops, workshops, supply networks, social venues, cities, organisations, specialist communities, commercial districts, brands, and player reputations.

The world should appear to continue existing independently of any particular player's story. Players participate in the world rather than being the centre of it.

The game's primary content is the interaction between players, world systems, and player-created consequences.

---

## VIS-002 — Citizen Rather Than Chosen Hero

**Status: LOCKED**

Players should feel like citizens of a living society rather than protagonists of a predetermined story.

The player should be able to live an ordinary life and still have a meaningful experience. The game should not require a heroic narrative to provide purpose.

A player's prestige should be able to derive from non-combat achievements, including:

- discovering a rare resource,
- producing respected goods,
- operating a reliable shop,
- becoming a famous doctor or entertainer,
- coordinating a city,
- supplying a war effort,
- operating a trusted trade network,
- designing desirable buildings,
- becoming an expert in a narrow specialisation.

A player should be able to say:

> "This is who I am in the world."

rather than:

> "This is the quest line I completed."

---

## VIS-003 — Player-Driven Economy

**Status: LOCKED**

Players should be major participants in the production and distribution of economic goods and services.

The economy must be materially connected to the world. Resources, labour, transport, crafting, consumption, degradation, services, and regional differences should matter.

The economy should generate meaningful opportunities for production, trade, specialisation, commerce, scarcity, reputation, and entrepreneurship.

---

## VIS-004 — Interdependence

**Status: LOCKED**

Important professions should depend upon one another.

A complete economy should therefore contain useful roles that cannot all be replaced by a single generalist player.

---

## VIS-005 — Emergence

**Status: LOCKED**

The game should create conditions from which interesting events emerge rather than attempting to script every important player experience.

The desired player thought is:

> "This world changed because of what players did, and other players know who matters here."

---

# 4. Design Pillars

## PIL-001 — Discovery and the Gold Rush

**Status: LOCKED**

Resources are temporary, geographically distributed, variable in quality, and economically meaningful.

A player discovers an exceptional temporary resource. The discovery may have unusually high quality, unusual attributes, exceptional rarity, exceptional economic value, or temporary availability. The player should feel urgency.

Exceptional resource discoveries should create player-driven events:

- prospecting activity,
- mining camps,
- trade opportunities,
- speculation,
- guild coordination,
- regional traffic,
- competitive supply chains,
- demand for extractors, transport, processing, and manufacturing.

A resource discovery should be an **event**, not merely an inventory update.

### Historical SWG Inspiration

Original SWG used temporary resource spawns with variable statistics, regional availability, and resource surveying. This created memorable periods where players sought unusually valuable resources.

### TCIndustries Adaptation

**PROPOSED — RES-001**

TCIndustries will retain the principle of dynamic, temporary, statistically variable resource occurrences, while improving information clarity, anti-monopoly controls, and long-term economic observability.

---

## PIL-002 — Identity and Reputation

**Status: LOCKED**

Player identity must be meaningful, visible, and socially consequential.

Products, services, businesses, and buildings should preserve meaningful association with their creator, provider, organisation, or brand.

Players should be capable of becoming recognised specialists:

> "That player makes the best armour."

> "That shop always has good medical supplies."

> "That prospector finds exceptional resources."

Reputation should emerge from player behaviour and economic history rather than merely from an abstract reputation score.

The ideal outcome is:

> Players do not merely buy "a product"; they seek out products made by a particular player or organisation.

---

## PIL-003 — Interdependence

**Status: LOCKED**

No major profession category should be fully self-sufficient at high effectiveness.

Interdependence should arise through:

- distinct skills,
- specialised resources,
- intermediate components,
- service needs,
- regional logistics,
- knowledge gaps,
- reputation,
- production capacity limits,
- social organisation.

Interdependence must create opportunity rather than coercive inconvenience.

A specialist should potentially experience:

> "People need what I do."

This is a central source of player identity.

---

## PIL-004 — Player-Created Society

**Status: LOCKED**

Players must be able to form organisations, businesses, settlements, social hubs, trade networks, and reputational institutions with meaningful world presence.

A player's shop should potentially become:

- a business,
- a meeting place,
- a known brand,
- part of a commercial district,
- part of a city.

Players should be able to feel:

- "This is my shop."
- "I made this product."
- "I discovered this resource."
- "This is my business."
- "This is our city."

---

# 5. Player Experience and Core Loops

## LOOP-001 — Multi-Loop Player Experience

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

---

## LOOP-002 — Non-Combat Viability

**Status: LOCKED**

A player must be able to pursue a meaningful long-term career without engaging in combat as their primary activity.

Combat may offer demand, risk, materials, access, and opportunities, but must not be the sole source of progression, wealth, or social status.

---

## LOOP-003 — Progression Expands Options

**Status: LOCKED**

Progression should expand choices, capability, specialisation, reputation, and access to complex roles. It should not force all players through one linear power ladder.

---

# 6. Anti-Goals

## AG-001 — Do Not Become a Spreadsheet Simulator

**Status: LOCKED**

Economic depth is desirable. Economic abstraction that removes the human meaning of the economy is not. The game should not converge on an experience where optimal play is managing a spreadsheet.

---

## AG-002 — Do Not Become a Single-Player Industrial Tycoon

**Status: LOCKED**

The player should not simply acquire unlimited factories and automate every stage personally. Industrial scale should encounter meaningful constraints that drive social cooperation.

---

## AG-003 — Do Not Make Everyone Self-Sufficient

**Status: LOCKED**

Complete self-sufficiency would undermine interdependence. If a single player can efficiently perform every important economic role, the social economy collapses.

---

## AG-004 — Do Not Force Rigid Class Identity

**Status: LOCKED**

The system should not permanently label players as one class. Players should be able to develop, change, and combine their specialisations over time.

---

## AG-005 — Do Not Make Content Completely Script-Dependent

**Status: LOCKED**

The world should generate meaningful experiences through systems and player interaction, not solely through scripted quests and narrative content.

---

## AG-006 — Do Not Create NPC-Dominated Economy

**Status: LOCKED**

The economy must not be one where all meaningful production is performed by NPCs, rendering player economic activity irrelevant.

---

## AG-007 — Do Not Reduce to Combat-Only MMORPG

**Status: LOCKED**

Combat is important but must not be the sole or dominant measure of player value, progression, or social standing.

---

# 7. World Philosophy and Structure

## WRLD-001 — Persistent World

**Status: LOCKED**

TCIndustries is a persistent shared world. Player actions should have durable consequences where practical, including:

- built structures,
- business inventories,
- market prices,
- resource depletion,
- city development,
- organisational ownership,
- local reputation,
- product provenance.

---

## WRLD-002 — Geographic Structure

**Status: PROPOSED**

The world should consist of multiple regions with distinct combinations of:

- terrain,
- climate,
- resource pools,
- settlement potential,
- dangers,
- transportation constraints,
- local demand,
- aesthetic identity,
- environmental conditions.

Regions should not merely be level bands. They should create economic and social differentiation.

---

## WRLD-003 — Settlements, Cities, and Wilderness

**Status: PROPOSED**

The world should contain:

1. **NPC Settlements** — Stable service, trade, tutorial, and baseline social locations.
2. **Player Settlements/Cities** — Player-founded areas with governance, zoning, commerce, civic structures, and community identity.
3. **Wilderness Regions** — Exploration, harvesting, combat, resource discovery, and construction opportunities.
4. **Economic Corridors** — Roads, transport routes, ports, stations, or equivalent systems connecting markets and resource regions.

Exact city-formation rules are **TBD**.

---

## WRLD-004 — Time and Environmental Simulation

**Status: PROPOSED**

The world should use a persistent world clock and support at least basic environmental variation.

Potential environmental factors include:

- day/night cycle,
- weather,
- seasons,
- ecological changes,
- temporary hazards,
- travel conditions,
- resource extraction modifiers.

Weather and seasons must only be implemented where they create meaningful decisions rather than visual noise.

---

## WRLD-005 — Transportation

**Status: PROPOSED**

Transportation should create meaningful geography without turning routine trade into excessive travel friction.

The system should support a spectrum:

- local movement,
- personal vehicles or mounts,
- public transit,
- organisation-operated transport,
- freight logistics,
- fast travel under constraints.

Fast travel must not eliminate the economic relevance of location, local markets, or transport services.

Exact transportation mechanics are **TBD**.

---

# 8. World Simulation

## WSIM-001 — Simulation Philosophy

**Status: PROPOSED**

The world should simulate enough background activity to feel alive.

Potential systems include:

- resource changes,
- markets,
- NPC activity,
- transportation,
- population,
- environmental changes,
- production,
- consumption.

The design principle is:

> Simulate systems that produce meaningful consequences; avoid simulation for its own sake.

The exact simulation depth is **TBD**.

---

# 9. Character, Identity, and Ownership

## PLR-001 — Character Identity

**Status: LOCKED**

Each player character has a persistent identity that can accumulate:

- name,
- appearance,
- ownership,
- skill history,
- profession specialisations,
- reputation,
- business affiliation,
- organisation membership,
- product/service provenance,
- social relationships.

The player's identity should emerge from skills, activities, possessions, businesses, reputation, relationships, location, organisations, products, and services.

---

## PLR-002 — Character Creation

**Status: PROPOSED**

Character creation should prioritise identity and aesthetic expression rather than granting permanent mechanical superiority.

Initial mechanical differences, if any, must not trap players into inferior long-term choices.

Species, ancestry, culture, and starting location design are **DEFERRED** pending setting development.

---

## PLR-003 — Ownership

**Status: LOCKED**

Players and organisations must be able to own meaningful persistent assets, including:

- items,
- resources,
- structures,
- vendors,
- business inventory,
- manufacturing permissions,
- organisation property,
- brands,
- contracts where implemented.

Ownership is central to the emotional experience. Players should feel:

- "This is my shop."
- "I made this product."
- "I discovered this resource."
- "This is my business."
- "This is our city."

Ownership systems must include safeguards against fraud, abandoned assets, and exploitative transfer mechanics.

---

# 10. Skills, Professions, and Progression

## PROF-001 — Flexible Skill-Based Professions

**Status: LOCKED**

TCIndustries uses a flexible skill-based profession architecture rather than rigid immutable classes.

A profession is an emergent identity resulting from acquired skills, specialisation choices, equipment, social reputation, and player behaviour.

A player should be able to learn skills, improve skills, specialise, combine disciplines, abandon skills, and pursue another specialisation.

---

## PROF-002 — Skill Domains

**Status: PROPOSED**

Initial skill domains should include:

- resource surveying and prospecting,
- extraction and harvesting,
- processing and refining,
- crafting disciplines (weaponsmithing, armoursmithing, tailoring, etc.),
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

---

## PROF-003 — Specialisation

**Status: PROPOSED**

Players may develop broad competency across domains but should require meaningful investment to reach high-tier specialisation.

A master weaponsmith, for example, should not automatically also be a top medical practitioner, architect, merchant, and combatant.

Specialisation creates markets for experts.

---

## PROF-004 — Respecialisation

**Status: LOCKED**

Players must be able to respecialise. Permanent character traps are contrary to sandbox identity development.

**PROPOSED implementation principle:** Respecialisation should involve meaningful opportunity cost, cooldown, retraining time, loss of active specialisation capacity, or economic cost — but should not erase player identity, item provenance, business ownership, or reputation history.

Exact respecialisation rules are **TBD**.

---

## PROF-005 — Specialisation Budget / Capability Cap

**Status: PROPOSED**

The game should have a capability cap or specialisation budget that prevents one character from mastering every economically important role at once.

The cap must preserve:

- experimentation,
- casual participation,
- hybrid roles,
- respecialisation,
- player autonomy.

The exact cap model is **TBD**.

---

## PROF-006 — Profession Categories

**Status: PROPOSED**

The initial conceptual profession space includes:

| Category | Example Roles |
|---|---|
| Resource | Prospector, miner, harvester, farmer, specialist gatherer |
| Craft | Artisan, weaponsmith, armourer, architect, specialist manufacturer |
| Service | Doctor, entertainer, transporter, repair specialist, researcher |
| Commerce | Merchant, trader, shopkeeper, broker |
| Combat | Combat specialists (multiple disciplines) |

These are categories, not necessarily final profession names. Professions should emerge from combinations of skills rather than rigid permanent classes.

---

# 11. Resource System

## RES-002 — Resource Design Goal

**Status: LOCKED**

The resource system exists to create discovery, scarcity, regional trade, product differentiation, temporary opportunity, and social/economic events.

Resources must not function solely as generic crafting currencies.

A resource should not merely be:

> "Iron: quantity 1000."

Resources should potentially have a meaningful identity consisting of type, location, quality, attributes, rarity, availability, temporal characteristics, and economic desirability.

---

## RES-003 — Resource Taxonomy

**Status: PROPOSED**

Resources should be defined through extensible data categories:

- resource family,
- resource class,
- subtype,
- origin,
- material state,
- applicable uses,
- measurable attributes,
- rarity characteristics,
- spawn rules,
- extraction requirements,
- processing options.

Example high-level families:

- minerals,
- metals,
- stone,
- chemical compounds,
- organics,
- fibres,
- woods,
- fuels,
- gases,
- biological materials,
- water and liquids,
- rare technical materials.

These are taxonomy examples, not a locked content list. The taxonomy must be extensible to support future content additions without restructuring the resource engine.

---

## RES-004 — Resource Instances

**Status: PROPOSED**

A resource spawn is an individual temporary resource instance with:

- unique identifier,
- name,
- taxonomy classification,
- spawn time,
- expected end time or lifecycle state,
- geographic distribution,
- density,
- attribute values,
- rarity,
- extraction difficulty,
- discoverability,
- remaining accessible quantity where applicable.

Each instance should be distinguishable from generic material categories.

Example:

> "Kavren Ferritic Alloy, discovered in the North Ember Plains, cycle 184."

---

## RES-005 — Resource Attributes

**Status: PROPOSED**

Resources should possess measurable attributes relevant to recipes and product outcomes.

Possible attributes include:

- purity,
- structural strength,
- conductivity,
- malleability,
- heat resistance,
- density,
- elasticity,
- chemical stability,
- energy yield,
- biological vitality,
- aesthetic quality,
- corrosion resistance.

Recipes should reference attribute needs rather than only requiring a named material.

This creates opportunities for resource specialisation, experimentation, market differentiation, and product optimisation.

Exact initial attribute library is **TBD**.

---

## RES-006 — Discovery and Surveying

**Status: PROPOSED**

Prospecting should be an active information-gathering profession.

Players may use surveying tools, local knowledge, maps, trade channels, and organisations to identify resource occurrences.

Discovery information may be:

- private,
- shared with an organisation,
- sold,
- published,
- made public through market behaviour,
- inferred from regional activity.

The game should permit information asymmetry, but must avoid allowing a small group to permanently monopolise all useful knowledge.

---

## RES-007 — Resource Quality

**Status: LOCKED**

Resource quality must have meaningful gameplay consequences.

A high-quality resource should potentially produce superior crafted goods.

Quality should not simply be cosmetic.

---

## RES-008 — Temporal Availability and Depletion

**Status: PROPOSED**

Resource spawns should be temporary. Individual extraction sites may also be depleted or become inefficient over time.

Two separate concepts must be retained:

1. **Spawn lifecycle:** The resource exists in the world for a limited period.
2. **Local extraction state:** A specific location may become exhausted, congested, inaccessible, or less efficient.

Exact durations, depletion curves, and respawn rules are **TBD** and require economic simulation.

---

## RES-009 — Extraction

**Status: PROPOSED**

Extraction should involve different scales:

- manual or personal gathering,
- small-scale equipment,
- stationary extraction installations,
- organisation-supported extraction,
- specialised hazardous extraction.

Automation may assist extraction but must include limitations such as:

- installation costs,
- maintenance,
- fuel or power,
- land rights,
- output caps,
- local competition,
- setup time,
- transport needs,
- vulnerability or access constraints.

A player should not be able to scale indefinitely through unattended extraction infrastructure.

---

## RES-010 — Processing

**Status: PROPOSED**

Raw resources may be transformed into processed materials and intermediate goods.

Processing should create new specialisation and trade opportunities but should not require excessive meaningless conversion steps.

A processing chain is justified only if it introduces at least one of:

- distinct skill role,
- location decision,
- quality tradeoff,
- resource transformation,
- economic market,
- logistical consequence.

---

# 12. Crafting System

## CRFT-001 — Crafting Philosophy

**Status: LOCKED**

Crafting must produce meaningful product variation, player identity, and market differentiation.

Crafting is not merely a recipe-unlock treadmill. Crafting should be a meaningful profession rather than merely an item-production menu.

A product's desirability should depend on some combination of:

- crafter skill,
- specialisation,
- schematic quality,
- resource attributes,
- component quality,
- experimentation choices,
- cosmetic customisation,
- provenance,
- reliability,
- reputation.

---

## CRFT-002 — Resource Choice Affects Output

**Status: LOCKED**

The identity and quality of input resources should influence crafted output.

Therefore:

> Two players making the same nominal item should be capable of producing materially different products.

---

## CRFT-003 — Schematics

**Status: PROPOSED**

A schematic defines:

- required skill domains and tiers,
- input categories,
- component slots,
- minimum resource attributes,
- production process,
- possible experimentation properties,
- output type,
- base statistics,
- quality limits,
- optional customisation slots,
- durability rules,
- manufacture eligibility.

Schematics should reference categories and attributes where possible, not hard-coded resource names. This supports the extensibility principle (EXP-001).

---

## CRFT-004 — Crafting Process

**Status: PROPOSED**

Crafting should include at least:

1. Schematic selection.
2. Input and component selection.
3. Material validation.
4. Assembly.
5. Experimentation or refinement where allowed.
6. Outcome calculation.
7. Item creation and provenance recording.

The exact interaction model is **TBD**. It may involve a structured decision process rather than a timing minigame.

---

## CRFT-005 — Experimentation

**Status: PROPOSED**

Experimentation allows skilled crafters to influence item properties within schematic-defined limits.

Potential outcomes include:

- higher performance,
- increased durability,
- lower weight,
- improved efficiency,
- improved appearance,
- increased reliability,
- reduced maintenance requirements.

Experimentation should include risk or constrained tradeoffs. It must not be a deterministic "always maximise every statistic" button. Experimentation should allow skilled players to make meaningful trade-offs rather than merely rolling for a random quality bonus.

Exact probability formulas are **TBD**.

---

## CRFT-006 — Product Quality

**Status: PROPOSED**

Item quality should result from a combination of:

- schematic baseline,
- crafter relevant skill,
- resource quality and attributes,
- components,
- experimentation outcome,
- manufacturing quality where applicable,
- optional service or calibration stage.

Quality must be intelligible enough for players to make informed choices.

---

## CRFT-007 — Provenance and Creator Identity

**Status: LOCKED**

Crafted items should retain provenance sufficient to support reputation and brand identity.

At minimum, item provenance should support:

- original creator,
- producing organisation if applicable,
- schematic lineage/version,
- major material identity where appropriate,
- creation time,
- item quality,
- modification history where relevant.

Privacy, pseudonymous brands, and resale handling are **TBD**.

---

## CRFT-008 — Signature Products

**Status: LOCKED**

Crafting should support the emergence of exceptional products associated with particular players.

Product reputation should potentially become part of the economic identity of the crafter.

---

# 13. Manufacturing and Automation

## MFG-001 — Manufacturing Goal

**Status: LOCKED**

Manufacturing exists to support production scale, supply chains, and businesses without replacing player specialisation and social interaction.

---

## MFG-002 — Automation Permitted, Not Dominant

**Status: LOCKED**

Automation is permitted. Automation must not eliminate the social economy.

Production scale must not automatically equal player power. A player controlling 100 automated production units should not simply invalidate 100 other players.

---

## MFG-003 — Factory Role

**Status: PROPOSED**

Factories may produce repeatable items from approved schematics and supplied inputs. They should be appropriate for:

- standard components,
- bulk consumables,
- common equipment,
- intermediate goods,
- established product lines.

Factories should not replace the value of master-crafted, experimental, limited, or custom goods.

---

## MFG-004 — Manufacturing Constraints

**Status: PROPOSED**

Manufacturing should be constrained by meaningful requirements:

- schematic authorisation,
- input materials,
- component availability,
- factory capacity,
- maintenance,
- location,
- fuel/power where appropriate,
- queue time,
- production quality limits,
- ownership limits,
- organisation coordination.

Exact factory limits are **TBD**.

---

## MFG-005 — Social Scale Principle

**Status: LOCKED**

Increasing economic scale must increasingly require organisations, trade, logistics, specialisation, reputation, and coordination — not merely additional personal factories.

Large economic organisations should generally require cooperation among multiple players. The game should favour:

> Many specialised participants

over:

> One player controlling everything.

---

## MFG-006 — Manufacturing Quality Ceiling

**Status: PROPOSED**

Manufactured products should generally have a lower quality ceiling or reduced customisation compared with direct specialist crafting, unless a later high-tier industrial specialisation is deliberately designed to alter this relationship.

This preserves the relevance of artisans while allowing practical volume production.

---

# 14. Items

## ITEM-001 — Item Categories

**Status: PROPOSED**

Items may include:

- equipment,
- weapons,
- armour,
- tools,
- harvesting devices,
- consumables,
- medical supplies,
- entertainment instruments,
- construction materials,
- furniture,
- resource containers,
- components,
- schematics,
- vehicles,
- decorations,
- trade goods.

---

## ITEM-002 — Durability and Degradation

**Status: LOCKED**

Items should not necessarily remain permanently useful. Many functional items should have durability, maintenance needs, degradation, or finite consumable use.

This creates:

- replacement demand,
- repair professions,
- resource consumption,
- continuing markets.

However, degradation must not feel like arbitrary punishment.

Durability should be appropriate to category:

- consumables: intentionally expended;
- tools and equipment: gradual wear;
- buildings/factories: maintenance;
- cosmetics: generally no functional decay;
- high-value legacy items: special rules **TBD**.

---

## ITEM-003 — Repair

**Status: PROPOSED**

Repair should preserve the value of repair services and material demand.

Possible rules include:

- repair consumes materials or repair kits;
- repair requires appropriate skill for high-value items;
- repeated repair may reduce maximum durability;
- specialist repair may preserve more value than generic repair.

Exact repair model is **TBD**.

---

## ITEM-004 — Customisation

**Status: PROPOSED**

Items should support appropriate cosmetic and naming customisation, subject to anti-abuse moderation rules.

Customisation should help support identity, commerce, gifting, city culture, and player brands.

---

## ITEM-005 — Item Provenance

**Status: PROPOSED**

Items should retain meaningful information about their origin and characteristics.

Potential identity includes:

- creator,
- resources used,
- quality,
- modifications,
- history,
- ownership.

The exact provenance model is **TBD** (see also CRFT-007).

---

# 15. Economy

## ECO-001 — Economic Philosophy

**Status: LOCKED**

The economy must be materially connected to the world. Resources, labour, transport, crafting, consumption, degradation, services, and regional differences should matter.

The game should not rely solely on abstract currency generation and vendor sinks.

---

## ECO-002 — Currency

**Status: PROPOSED**

The game should use a primary currency for player trade, NPC transactions, taxes, maintenance, and selected sinks.

Currency naming, source rates, and sink design are **TBD**.

---

## ECO-003 — Supply and Demand

**Status: PROPOSED**

Player pricing should be broadly free within anti-exploit safeguards.

Supply and demand should be shaped by:

- temporary resources,
- item degradation,
- combat demand,
- city population,
- travel and location,
- industrial input demand,
- fashion and customisation,
- player events,
- new schematics and technology,
- regional specialisation.

Exact market mechanics are **TBD**.

---

## ECO-004 — Scarcity

**Status: LOCKED**

Scarcity must matter.

If an exceptional resource is genuinely scarce, players should have economic reasons to compete for it.

If every good is always abundant, the economy loses meaning and player specialisation loses value.

---

## ECO-005 — Economic Sinks

**Status: PROPOSED**

The economy requires recurring sinks for currency and materials.

Potential sinks include:

- building maintenance,
- factory operation,
- repair,
- transportation,
- taxes,
- NPC services,
- construction,
- consumables,
- item wear,
- cosmetic customisation,
- organisation upkeep.

Sinks must not disproportionately punish low-income or casual players.

---

## ECO-006 — NPC Economic Role

**Status: PROPOSED**

NPC vendors should provide baseline accessibility and emergency supply, but should not invalidate player commerce for economically meaningful goods.

NPC buy orders, if used, must be carefully limited to avoid creating infinite-risk-free farming loops.

---

## ECO-007 — Market Information

**Status: PROPOSED**

Market information should be useful but not perfectly frictionless.

Players should be able to discover prices through shops, market boards, trade networks, contracts, and reputation. Full global instant price transparency may reduce regional economies and merchant roles.

Exact market-information model is **TBD**.

---

# 16. Retail and Commerce

## RET-001 — Player Shops as Major Playstyle

**Status: LOCKED**

Player-run retail must be a major supported playstyle.

Players should be able to establish shops, stock goods, set prices, curate presentation, and develop repeat customers.

A shop should be more than an inventory container. It should potentially become a destination, a brand, a social hub, and a source of reputation.

---

## RET-002 — Vendors and Storefronts

**Status: PROPOSED**

Retail may use vendor systems located in player-owned or organisation-owned structures.

A shop should support:

- inventory listing,
- pricing,
- product descriptions,
- creator and brand information,
- shop identity,
- visitor access,
- optional sales history,
- customer discovery tools,
- permissions.

---

## RET-003 — Location and Commercial Districts

**Status: PROPOSED**

Location should influence commerce through visibility, convenience, city foot traffic, regional supply, transport access, events, and nearby services.

However, remote or specialist businesses must remain viable through directories, transport, reputation, delivery, or trade networks.

---

## RET-004 — Merchant Specialisation Value

**Status: PROPOSED**

Merchant-oriented skills should provide meaningful advantages beyond "placing a vendor," such as:

- improved market intelligence,
- better shop discovery,
- contract access,
- reduced transaction friction,
- product presentation,
- inventory management,
- trade-network tools,
- business reputation features.

Exact merchant mechanics are **TBD**.

---

# 17. Services

## SERV-001 — Service Philosophy

**Status: LOCKED**

Non-item player services are first-class economic and social gameplay, not secondary flavour.

Not all valuable player output should be an item. Services should constitute a meaningful portion of the economy.

---

## SERV-002 — Medical Services

**Status: PROPOSED**

Medical specialists may provide healing, recovery, treatment, stabilisation, rehabilitation, medical consumables, and possibly temporary performance support.

Medical gameplay must create real reasons to visit or hire players while avoiding mandatory downtime that prevents normal play.

---

## SERV-003 — Entertainment and Social Services

**Status: PROPOSED**

Entertainers may provide social gathering value, morale/recovery benefits, temporary buffs, event hosting, venue identity, and cultural expression.

The system must avoid reducing entertainers to passive buff dispensers.

---

## SERV-004 — Repair, Construction, Research, and Transport

**Status: PROPOSED**

Potential player service categories include:

- equipment repair,
- structure construction,
- maintenance,
- transport,
- freight movement,
- resource analysis,
- research,
- appraisal,
- design consultation,
- specialist customisation.

Each service must be assessed for whether it produces meaningful interaction rather than artificial gatekeeping.

---

# 18. Combat

## CMBT-001 — Combat Role in the World

**Status: LOCKED**

Combat is an important world activity but not the sole or dominant measure of player value.

Combat should generate demand for:

- weapons,
- armour,
- tools,
- medical services,
- consumables,
- repairs,
- transport,
- specialist equipment,
- crafted components.

Combat therefore participates in the broader economy and reinforces interdependence.

---

## CMBT-002 — Detailed Combat Design

**Status: DEFERRED**

Detailed combat mechanics, damage models, targeting, enemy AI, PvP policy, faction systems, and encounter design require a dedicated combat specification.

Combat design should proceed after the broader world/economic architecture is sufficiently established.

---

## CMBT-003 — PvP Policy

**Status: TBD**

The project has not yet established whether TCIndustries includes open PvP, consensual PvP, territorial conflict, criminal systems, faction warfare, or fully PvE gameplay.

> [!WARNING]
> This decision strongly affects cities, resource access, transport, item loss, insurance, organisations, and economic risk. It must be resolved before many dependent systems can be finalised.

---

# 19. Buildings, Housing, and Cities

## BLD-001 — Building Ownership

**Status: LOCKED**

Players and organisations should be able to own meaningful structures, including housing, shops, workshops, factories, and civic buildings.

---

## BLD-002 — Building Functions

**Status: PROPOSED**

Structures may provide:

- storage,
- habitation,
- social space,
- retail space,
- manufacturing capacity,
- workshop bonuses,
- city services,
- organisation administration,
- aesthetic identity,
- event hosting.

---

## BLD-003 — Maintenance and Limits

**Status: PROPOSED**

Buildings should require upkeep and obey ownership, zoning, placement, and population limits.

The purpose is to prevent abandoned-world clutter, land monopolisation, and unlimited asset accumulation.

Exact upkeep, decay, reclamation, and placement rules are **TBD**.

---

## CITY-001 — Player Cities

**Status: PROPOSED**

Player cities should be player-founded social and economic institutions, not merely housing clusters.

A city should potentially emerge from player structures, population, services, commerce, organisations, and infrastructure. Cities should become social and economic ecosystems rather than merely cosmetic housing zones.

Potential city systems include:

- civic roles,
- membership,
- zoning,
- taxes,
- city projects,
- local services,
- public buildings,
- event tools,
- commercial support,
- city reputation,
- infrastructure.

City governance model is **TBD**.

---

# 20. Social Systems and Organizations

## SOC-001 — Organisations

**Status: LOCKED**

Players must be able to form persistent organisations that support shared identity, ownership, cooperation, and economic coordination.

The exact terminology — guild, company, association, cooperative, syndicate, clan, or other — is setting-dependent and **TBD**.

---

## SOC-002 — Organisation Capabilities

**Status: PROPOSED**

Organisations may support:

- membership and roles,
- shared assets,
- shared storage,
- property ownership,
- manufacturing rights,
- city participation,
- shared brands,
- contracts,
- internal permissions,
- event coordination,
- public identity.

---

## SOC-003 — Reputation

**Status: PROPOSED**

Reputation should be primarily social and evidence-based, supported by system visibility rather than reduced to a universal numeric score.

Possible evidence includes:

- item provenance,
- customer reviews with safeguards,
- transaction history,
- city affiliation,
- organisation affiliation,
- service records,
- contracts completed,
- public achievements,
- local reputation.

A single global reputation number is discouraged because it can become gameable, reductive, and socially coercive.

---

## SOC-004 — Interdependence Chains

**Status: LOCKED**

The game should contain dependency chains such as:

**Resource Specialist** → raw material → **Crafter** → equipment → **Combatant**

and:

**Resource Specialist** → resource → **Manufacturer** → component → **Crafter** → finished product → **Retailer** → customer.

Likewise:

**Combatant** → demand → **Doctor / Entertainer / Crafter** → services/equipment → Combatant.

No single chain should necessarily dominate the entire game. Multiple interdependence paths should coexist.

---

# 21. Progression

## PROG-001 — Progression Philosophy

**Status: LOCKED**

Progression should primarily support:

- capability,
- specialisation,
- mastery,
- identity.

Progression should not exist solely to create an arbitrary treadmill. It should expand options and open doors to complex roles rather than merely increasing numbers.

---

## PROG-002 — Skill Acquisition Model

**Status: TBD**

The exact model for skill acquisition is unresolved. Possibilities include:

- use-based learning,
- training from NPCs or players,
- skill-point allocation,
- certifications,
- mentorship,
- quest/mission unlocks,
- hybrid approaches.

This decision affects progression feel, exploit risks, and how quickly specialisation develops.

---

## PROG-003 — Mastery

**Status: PROPOSED**

Mastery should represent deep expertise within a domain that confers meaningful advantages without creating insurmountable barriers for non-masters.

Mastery should create respect, demand, and pride — not gatekeeping.

---

# 22. Emergent Gameplay

## EMRG-001 — Emergent Gameplay Targets

**Status: LOCKED**

The following are desired emergent outcomes and should be used as design tests for all systems:

1. A rare, high-quality resource appears and creates a temporary prospecting rush.
2. A known crafter becomes sought after because their products consistently perform well.
3. A city develops an identity as a medical, industrial, cultural, or trade centre.
4. A combat conflict or PvE event increases demand for particular equipment and medical supplies.
5. A shortage of a component causes merchants and manufacturers to form new supply relationships.
6. A player organisation coordinates extraction, processing, production, retail, and logistics.
7. A remote specialist gains customers through reputation rather than only geographic luck.
8. Player-created events generate trade, travel, entertainment, and social interaction.
9. Resource geography causes regional economies without making participation inaccessible.
10. A player can become socially important without being a top combat player.

### Example: Exceptional Resource Chain

A rare high-quality resource appears.
→ Prospector discovers it.
→ Resource prices change.
→ Gatherers rush to the region.
→ Crafters purchase the resource.
→ Exceptional products appear.
→ Product reputation develops.
→ Retailers advertise the products.
→ Combat players seek them.
→ Demand increases.

This is the desired type of systemic chain.

---

# 23. Content Expansion Framework

## EXP-001 — Data-Driven Design Principle

**Status: LOCKED**

Core systems must be defined through extensible abstractions rather than hard-coded isolated content.

Future additions should be able to introduce:

- new resource categories and attributes,
- schematics,
- item types,
- professions,
- buildings,
- services,
- regions,
- organisations,
- NPC roles,
- transportation options,
- research paths,
- world events.

For example, adding a new resource should not require redesigning the resource engine. Adding a new product should not require redesigning crafting. Adding a new profession should not require rewriting the player model.

---

## EXP-002 — Expansion Content Requirements

**Status: PROPOSED**

New content should declare:

- prerequisites,
- inputs,
- outputs,
- category relationships,
- quality rules,
- rarity rules,
- economic impact,
- progression relevance,
- social implications,
- balance risks.

No content addition should be accepted solely because it adds volume; it should reinforce at least one design pillar.

---

# 24. Balance Philosophy

## BAL-001 — Meaningful Imperfection

**Status: PROPOSED**

The economy should not converge permanently on one universally optimal product, profession, resource, or region.

Variation should arise from:

- temporary resource quality,
- changing availability,
- different player skills,
- regional demand,
- item niches,
- customisation,
- logistics,
- reputation,
- situational combat needs.

---

## BAL-002 — Accessibility Versus Dependency

**Status: PROPOSED**

Interdependence must not become hostage gameplay.

Players should have access to basic self-sufficiency and emergency NPC support, while premium efficiency, quality, specialisation, and convenience come from player relationships.

---

## BAL-003 — Casual and Long-Term Players

**Status: PROPOSED**

Casual players should be able to participate meaningfully in discovery, gathering, local trade, services, and social life.

Long-term players should gain specialisation, reputation, organisational influence, and operational complexity — not unchecked combat or economic dominance.

---

# 25. Exploit, Fraud, and Economic Stability

## SAFE-001 — General Principle

**Status: LOCKED**

Systems involving trade, ownership, organisations, automation, and reputation must be designed with fraud and exploit resistance as first-class requirements.

---

## SAFE-002 — Risks Requiring Dedicated Design

**Status: TBD**

The following require detailed rule specifications:

- duplicate item prevention,
- currency generation limits,
- NPC vendor arbitrage,
- abandoned structure cleanup,
- land monopolisation,
- resource spawn monopolisation,
- multi-account extraction abuse,
- automated market manipulation,
- vendor bait-and-switch,
- organisation asset theft,
- contract fraud,
- price manipulation,
- griefing through city governance,
- player review abuse,
- service-payment disputes,
- inactive owner handling.

---

## SAFE-003 — Economic Monitoring

**Status: PROPOSED**

The future live game should monitor economic health through simulation and telemetry, including:

- currency creation and destruction,
- resource extraction volumes,
- item production and destruction,
- price distributions,
- market concentration,
- new-player access to essentials,
- regional supply imbalance,
- organisation asset concentration,
- factory utilisation,
- abandoned property rates.

---

# 26. Testing Requirements and Design Invariants

## TEST-001 — Core Invariants

**Status: PROPOSED**

Future implementations should validate the following invariants:

1. An item cannot be created without valid source inputs, approved system rewards, or explicit administrative creation.
2. Crafted provenance must identify the creator or permitted producing entity.
3. A schematic cannot accept materials outside its defined category and attribute requirements.
4. Resource spawn identity must remain stable for its lifecycle.
5. Extraction cannot produce more material than allowed by spawn and extraction rules.
6. Manufacturing cannot produce outputs without valid inputs, capacity, and permissions.
7. Ownership transfer must be atomic and auditable.
8. Currency transfers must conserve value except where explicit source/sink rules apply.
9. Skill caps must prevent forbidden simultaneous mastery combinations once defined.
10. Building placement must obey land, ownership, and zoning constraints.
11. Repair and degradation must not duplicate items or create material value.
12. No standard system should permit infinite risk-free currency or item generation.

---

## TEST-002 — Simulation Requirements

**Status: PROPOSED**

Before major system rules are locked, simulations should test:

- resource spawn frequency and gold-rush behaviour;
- high-quality resource concentration;
- new-player access to basic materials;
- factory output versus artisan relevance;
- price inflation and currency sink adequacy;
- organisation concentration and monopoly risk;
- durability/replacement demand;
- time-to-specialisation;
- viability of non-combat careers;
- viability of small independent shops;
- regional market behaviour;
- effect of travel costs and fast travel;
- multi-account and unattended automation incentives.

---

## TEST-003 — Promotion Criteria

**Status: PROPOSED**

A system should not be promoted to LOCKED until it has:

1. a defined player goal;
2. explicit inputs, state, outputs, and failure cases;
3. identified dependencies;
4. exploit considerations;
5. simulation or prototype validation where appropriate;
6. clear player-visible outcomes;
7. consistency with all four design pillars.

---

# 27. Historical SWG Relationship

## SWG-001 — Inspiration Policy

**Status: LOCKED**

TCIndustries is heavily inspired by the virtual-world philosophy and selected systemic lessons of original SWG, but it is not a Star Wars game and must not reproduce protected Star Wars intellectual property.

---

## SWG-002 — Adaptation Rule

**Status: LOCKED**

When adopting a historical SWG-inspired mechanic, design documentation must distinguish:

- **Historical SWG behaviour:** What SWG did, to the extent accurately known.
- **TCIndustries adaptation:** What TCIndustries chooses to retain or alter.
- **New TCIndustries invention:** A system created specifically for TCIndustries.

Historical similarity does not automatically make a mechanic canonical.

---

# 28. Prototype Relationship

## PROTO-001 — Prototype Status

**Status: LOCKED**

The Seed-2.1 TypeScript/Vite/React prototype is:

**Status: PROTOTYPE**

It is evidence and experimentation, not authority.

The prototype may be used as:

- a source of prototype ideas,
- evidence of previously explored mechanics,
- an implementation reference,
- a starting point for experimentation,
- a test harness.

However, the prototype is NOT authoritative. Prototype behaviour must not be treated as a design decision merely because code already implements it. The prototype may eventually be discarded and rebuilt.

The development hierarchy is:

**GDD** → canonical game rules
→ **Prototype** → experimental implementation
→ **Testing** → validation
→ **Golden Master** → validated behaviour
→ **Future UE5/C++** → production implementation.

---

# 29. Systems Requiring Deeper Design

The following systems are recognised but require dedicated canonical specifications before implementation can be considered authoritative.

| Priority | System | Reason |
|---|---|---|
| 1 | Resource spawning and surveying | Foundation of gold-rush economy and production chain |
| 2 | Resource quality/attribute mathematics | Determines product differentiation and resource value |
| 3 | Skill budget and specialisation | Determines interdependence and player freedom |
| 4 | Crafting quality and experimentation | Determines product identity and crafter reputation |
| 5 | Item durability, repair, and destruction | Determines replacement demand and economic health |
| 6 | Extraction and automation limits | Prevents industrial-simulation drift and monopoly |
| 7 | Manufacturing rules | Defines artisan versus factory relationship |
| 8 | Currency sources, sinks, and NPC market role | Required for economic stability |
| 9 | Shops, vendors, and discovery | Required for player commerce viability |
| 10 | Buildings, land rights, maintenance, and decay | Required for persistent player society |
| 11 | Cities and organisation governance | Required for player-created institutions |
| 12 | Combat and PvP policy | Affects risk, demand, territorial systems, and social rules |
| 13 | Medical and entertainment service loops | Required for non-combat service professions |
| 14 | Transportation and regional market design | Determines whether geography matters |
| 15 | Reputation and contract systems | Required for trust at economic scale |

---

# 30. Recommended Design Sequence

## Phase 1 — Resource Foundation

1. Lock resource taxonomy.
2. Lock resource spawn lifecycle model.
3. Define surveying and discovery rules.
4. Define extraction tiers and automation limits.
5. Define processing categories.

## Phase 2 — Crafting

1. Define core crafting schematic model.
2. Define quality calculation inputs and outputs.
3. Define experimentation mechanics.
4. Define item provenance rules.

## Phase 3 — Items and Durability

1. Define item identity model.
2. Define durability, degradation, and destruction.
3. Define repair model.
4. Define modification and customisation.

## Phase 4 — Skills and Professions

1. Define skill domains and specialisation budget.
2. Define skill acquisition model.
3. Define crafting profession families.
4. Define resource, processing, and manufacturing dependencies.
5. Define basic medical and entertainment service loops.

## Phase 5 — Economy

1. Define currency sources and sinks.
2. Define vendor, shop, and market-discovery rules.
3. Define NPC economic role.
4. Define pricing and supply/demand framework.

## Phase 6 — Retail and Reputation

1. Define vendor and storefront capabilities.
2. Define merchant specialisation.
3. Define customer discovery.
4. Define brand and reputation evidence systems.
5. Define commercial location rules.

## Phase 7 — Manufacturing

1. Define factory capabilities and constraints.
2. Define quality ceiling relationship with artisan crafting.
3. Define ownership limits and organisation requirements.
4. Define maintenance and logistics.

## Phase 8 — Social World

1. Define organisation structure and permissions.
2. Define building ownership and maintenance.
3. Define city foundation and governance.
4. Define shared assets and brand systems.
5. Define contracts and escrow where needed.

## Phase 9 — Services and Combat

1. Define combat loop and equipment demand.
2. Decide PvP policy.
3. Define medical service mechanics.
4. Define entertainment service mechanics.
5. Define repair, transport, and construction services.
6. Define death, injury, and recovery rules.

## Phase 10 — World Simulation and Integration

1. Define transportation and logistics.
2. Define world events and environmental systems.
3. Define NPC society and activity.
4. Define research, technology, and advanced content progression.
5. Integrate all systems into a coherent simulated world.

---

# 31. Status Register

| ID | System/Decision | Status | Rationale | Dependencies |
|---|---|---|---|---|
| VIS-001 | Persistent player-driven virtual society | LOCKED | Core project objective | All systems |
| VIS-002 | Citizen identity over hero narrative | LOCKED | Core player fantasy | Progression, social, economy |
| VIS-003 | Player-driven economy | LOCKED | Central world philosophy | All economic systems |
| VIS-004 | Interdependence | LOCKED | Core design requirement | Skills, economy, services |
| VIS-005 | Emergent gameplay | LOCKED | Central design philosophy | All systems |
| PIL-001 | Dynamic resource discovery / gold rush | LOCKED | Explicit project pillar | Resources, economy |
| PIL-002 | Identity and reputation pillar | LOCKED | Explicit project pillar | Crafting, retail, services |
| PIL-003 | Interdependence pillar | LOCKED | Explicit project pillar | Skills, economy, services |
| PIL-004 | Player-created society pillar | LOCKED | Explicit project pillar | Buildings, cities, organisations |
| AG-001 | Avoid spreadsheet simulation | LOCKED | Anti-drift requirement | Economy |
| AG-002 | Avoid single-player industrial empire | LOCKED | Anti-drift requirement | Manufacturing, automation |
| AG-003 | Avoid universal self-sufficiency | LOCKED | Anti-drift requirement | Skills, professions |
| AG-004 | Avoid rigid classes | LOCKED | Anti-drift requirement | Progression |
| AG-005 | Avoid scripted-only content | LOCKED | Anti-drift requirement | World design |
| AG-006 | Avoid NPC-dominated economy | LOCKED | Anti-drift requirement | Economy |
| AG-007 | Avoid combat-only MMORPG | LOCKED | Anti-drift requirement | Combat, services |
| LOOP-001 | Multi-loop player experience model | PROPOSED | Supports sandbox participation | All major systems |
| LOOP-002 | Non-combat long-term viability | LOCKED | Explicit project direction | Professions, economy |
| LOOP-003 | Progression expands options, not linear power | LOCKED | Core sandbox philosophy | Skills |
| WRLD-001 | Persistent shared world | LOCKED | Required for living-world premise | World, ownership |
| WRLD-002 | Regionally differentiated geography | PROPOSED | Supports trade and discovery | Resources, transport |
| WRLD-003 | NPC settlements, player cities, wilderness | PROPOSED | World structure recommendation | Buildings, cities |
| WRLD-004 | Persistent time/environment systems | PROPOSED | Supports world life if meaningful | World simulation |
| WRLD-005 | Constrained transportation spectrum | PROPOSED | Preserves geography and commerce | Economy, cities |
| WSIM-001 | Meaningful simulation, not simulation for its own sake | PROPOSED | Design philosophy for world systems | World, resources |
| PLR-001 | Persistent character identity | LOCKED | Supports reputation and ownership | Player systems |
| PLR-002 | Cosmetic-first character creation | PROPOSED | Avoids permanent optimisation traps | Character creation |
| PLR-003 | Persistent player/organisation ownership | LOCKED | Required for society and economy | Buildings, items |
| PROF-001 | Flexible skill-based professions | LOCKED | Explicit project requirement | Progression |
| PROF-002 | Initial skill-domain categories | PROPOSED | Initial profession architecture | Content design |
| PROF-003 | Meaningful high-tier specialisation | PROPOSED | Supports interdependence | Skill budget |
| PROF-004 | Respecialisation allowed | LOCKED | Explicit project philosophy | Progression |
| PROF-005 | Capability cap / specialisation budget | PROPOSED | Prevents universal mastery | Skill design |
| PROF-006 | Profession categories (resource, craft, service, commerce, combat) | PROPOSED | Initial profession space definition | Content design |
| RES-001 | Dynamic temporary resource occurrences (SWG adaptation) | PROPOSED | Pillar implementation | Economy, surveying |
| RES-002 | Resources create discovery and scarcity | LOCKED | Pillar implementation | Economy, crafting |
| RES-003 | Extensible resource taxonomy | PROPOSED | Supports content expansion | Data definitions |
| RES-004 | Temporary resource instances with identity | PROPOSED | Supports gold-rush behaviour | Surveying, extraction |
| RES-005 | Attribute-driven resources | PROPOSED | Supports product differentiation | Crafting |
| RES-006 | Surveying as information profession | PROPOSED | Supports discovery and trade | Resources, skills |
| RES-007 | Resource quality has gameplay consequences | LOCKED | Core crafting requirement | Crafting, economy |
| RES-008 | Spawn lifecycle and local depletion | PROPOSED | Creates temporal opportunity | Extraction, economy |
| RES-009 | Constrained extraction automation | PROPOSED | Prevents unlimited solo scale | Buildings, manufacturing |
| RES-010 | Meaningful processing chains | PROPOSED | Supports intermediate markets | Crafting, economy |
| CRFT-001 | Crafting creates differentiated products | LOCKED | Core identity/reputation requirement | Resources, skills |
| CRFT-002 | Resource choice affects output | LOCKED | Core crafting requirement | Resources |
| CRFT-003 | Extensible schematic structure | PROPOSED | Expansion-friendly crafting | Resources, items |
| CRFT-004 | Structured crafting process | PROPOSED | Implementation basis | Schematics, skills |
| CRFT-005 | Constrained experimentation | PROPOSED | Supports mastery and variation | Crafting quality |
| CRFT-006 | Multi-input quality calculation | PROPOSED | Supports resource/crafter importance | Resources, skills |
| CRFT-007 | Crafted item provenance | LOCKED | Supports reputation and brands | Items, retail |
| CRFT-008 | Signature products possible | LOCKED | Supports crafter identity | Reputation |
| MFG-001 | Manufacturing supports rather than replaces players | LOCKED | Anti-drift requirement | Economy, factories |
| MFG-002 | Automation permitted, not dominant | LOCKED | Explicit design boundary | Economy, social |
| MFG-003 | Factories for repeatable production | PROPOSED | Practical scale without full automation | Schematics, buildings |
| MFG-004 | Factory constraints | PROPOSED | Anti-monopoly and social scale | Economy |
| MFG-005 | Scale increasingly requires coordination | LOCKED | Explicit design-scale principle | Organisations, economy |
| MFG-006 | Artisan quality advantage over factories | PROPOSED | Protects specialist relevance | Crafting, factories |
| ITEM-001 | Broad extensible item categories | PROPOSED | Content framework | Crafting |
| ITEM-002 | Category-appropriate degradation | LOCKED | Replacement demand | Economy, repair |
| ITEM-003 | Repair as meaningful service/material sink | PROPOSED | Supports economy and services | Items, skills |
| ITEM-004 | Item customisation support | PROPOSED | Supports identity and commerce | Items, retail |
| ITEM-005 | Item provenance model | PROPOSED | Supports reputation | Crafting, retail |
| ECO-001 | Materially connected player economy | LOCKED | Core world philosophy | All economic systems |
| ECO-002 | Primary currency | PROPOSED | Necessary transaction medium | Economy design |
| ECO-003 | Player-led supply/demand pricing | PROPOSED | Supports trade and merchant play | Retail, resources |
| ECO-004 | Scarcity matters | LOCKED | Supports competition and value | Resources, economy |
| ECO-005 | Recurring material/currency sinks | PROPOSED | Inflation control | Buildings, items |
| ECO-006 | Limited NPC baseline economy | PROPOSED | Prevents NPC invalidation of players | Economy |
| ECO-007 | Imperfect market information | PROPOSED | Preserves regional/merchant value | Retail, transport |
| RET-001 | Player-run retail as major playstyle | LOCKED | Explicit player-society objective | Buildings, economy |
| RET-002 | Vendor/storefront capabilities | PROPOSED | Retail functionality | Buildings, items |
| RET-003 | Location affects commerce | PROPOSED | Supports geography and cities | Transport, world |
| RET-004 | Merchant specialisation value | PROPOSED | Makes commerce a profession | Skills, retail |
| SERV-001 | Services are first-class gameplay | LOCKED | Supports non-combat identities | Skills, economy |
| SERV-002 | Medical service profession | PROPOSED | Interdependence | Combat, items |
| SERV-003 | Entertainment service profession | PROPOSED | Social-hub creation | Social, buffs |
| SERV-004 | Repair/construction/research/transport services | PROPOSED | Service economy | Multiple systems |
| CMBT-001 | Combat supports, not dominates, economy | LOCKED | Core philosophy | Items, services |
| CMBT-002 | Detailed combat design | DEFERRED | Requires dedicated document | PvP policy, items |
| CMBT-003 | PvP policy | TBD | High-impact unresolved decision | Cities, risk, transport |
| BLD-001 | Meaningful player/organisation buildings | LOCKED | Player-created society | Ownership, cities |
| BLD-002 | Multi-function structures | PROPOSED | Supports housing/businesses | Buildings |
| BLD-003 | Maintenance and placement limits | PROPOSED | Prevents clutter/monopoly | Economy, land |
| CITY-001 | Player cities as institutions | PROPOSED | Supports society and governance | Buildings, organisations |
| SOC-001 | Persistent player organisations | LOCKED | Required for cooperation | Ownership, cities |
| SOC-002 | Organisation roles/assets/permissions | PROPOSED | Coordination tools | Ownership, contracts |
| SOC-003 | Evidence-based reputation (not single score) | PROPOSED | Supports identity without reductive score | Retail, services |
| SOC-004 | Interdependence chains | LOCKED | Core systemic requirement | All professions |
| EMRG-001 | Emergent gameplay targets (10 scenarios) | LOCKED | Evaluation framework | All systems |
| EXP-001 | Extensible data-driven content abstractions | LOCKED | Expandability requirement | All systems |
| EXP-002 | Content declaration requirements | PROPOSED | Prevents unbalanced feature additions | Content pipeline |
| BAL-001 | Avoid permanent universal optimum | PROPOSED | Protects economic diversity | Resources, crafting |
| BAL-002 | Accessibility without forced dependency | PROPOSED | Prevents coercive friction | Economy, services |
| BAL-003 | Casual and veteran participation | PROPOSED | Long-term population health | Progression, economy |
| SAFE-001 | Fraud/exploit resistance as first-class concern | LOCKED | Required for persistent economy | All transactional systems |
| SAFE-002 | Detailed fraud protection rules | TBD | Requires system-level design | Ownership, markets |
| SAFE-003 | Live economic monitoring | PROPOSED | Necessary for tuning | Analytics/simulation |
| PROG-001 | Progression supports capability and identity | LOCKED | Core philosophy | Skills |
| PROG-002 | Skill acquisition model | TBD | Affects progression feel and exploits | Skills, balance |
| PROG-003 | Mastery as expertise, not gatekeeping | PROPOSED | Design direction | Skills |
| TEST-001 | Core system invariants (12 rules) | PROPOSED | Supports testable rules | Implementation specs |
| TEST-002 | Economic simulation requirements | PROPOSED | Validates systemic behaviour | Simulation |
| TEST-003 | Promotion criteria for LOCKED systems | PROPOSED | Design governance | All systems |
| SWG-001 | SWG inspiration without IP reproduction | LOCKED | Project requirement | Content, setting |
| SWG-002 | Historical/adaptation/invention distinction | LOCKED | Documentation discipline | All specs |
| PROTO-001 | Prototype is not authoritative | LOCKED | Design governance | All systems |
| WSIM-001 | Meaningful simulation philosophy | PROPOSED | World design principle | World systems |

---

# 32. Assumption Register

Every substantive assumption introduced by the designer is recorded here. Assumptions must not be treated as locked decisions.

| ID | Assumption | Basis | Consequence if Incorrect |
|---|---|---|---|
| ASM-001 | TCIndustries is intended to be a persistent online multiplayer world rather than session-based gameplay. | Supported by project prompt | World persistence and economy design would require revision. |
| ASM-002 | Dynamic resources remain a desired central system rather than optional flavour. | Strongly supported by pillars | Resource, crafting, and economy architecture would change substantially. |
| ASM-003 | Crafted items can meaningfully differ in quality and desirability. | Strongly supported by identity pillar | Crafter reputation model would weaken if false. |
| ASM-004 | The game will support player-owned structures in some form. | Supported by required GDD areas | Cities, retail, and manufacturing must be redesigned if absent. |
| ASM-005 | The setting is original and must avoid Star Wars IP. | Explicitly stated in project prompt | Setting/content design must remain distinct from SWG. |
| ASM-006 | A future simulation/testing stage will exist before production implementation is treated as final. | Supported by intended development path | Balance-validation process would need revision. |
| ASM-007 | Some form of player currency will exist. | Introduced for economy coherence | Barter-only or multi-currency alternatives require separate design. |
| ASM-008 | Some items will degrade, be consumed, or require maintenance. | Introduced to support replacement demand | Alternative sink models would be needed if rejected. |
| ASM-009 | Player skill specialisation will have a cap or budget. | Introduced to preserve interdependence | Without it, universal mastery risk increases. |
| ASM-010 | Factories and extractors can exist but must be constrained. | Supported by design-scale statement | Manufacturing loop must be adjusted if factory gameplay is excluded. |

---

# 33. Open-Question Register

The most important unresolved decisions are listed here. Important unresolved decisions must not be buried inside prose.

| ID | Question | Why It Matters | Recommended Phase |
|---|---|---|---|
| OQ-001 | What is the setting's original genre, technology level, visual identity, and world fiction? | Determines resource names, professions, buildings, combat, transport | World foundation |
| OQ-002 | What is the exact skill-cap or specialisation-budget model? | Determines interdependence and hybrid viability | Phase 4 |
| OQ-003 | How are skills acquired: use-based, training, points, certifications, mentorship, quests, or hybrid? | Determines progression behaviour and exploit risks | Phase 4 |
| OQ-004 | What exact resource attributes exist at launch? | Determines crafting formulas and resource value | Phase 1 |
| OQ-005 | How long do resource spawns last, and how severe is depletion? | Determines rush behaviour and market stability | Phase 1 |
| OQ-006 | How much surveying information is private versus public? | Determines discovery value and monopoly risk | Phase 1 |
| OQ-007 | What are the exact crafting quality and experimentation formulas? | Determines product differentiation | Phase 2 |
| OQ-008 | Can crafted items fail, partially fail, or only vary in quality? | Determines risk and usability | Phase 2 |
| OQ-009 | What is the exact factory quality ceiling and ownership limit? | Determines artisan relevance and automation risk | Phase 7 |
| OQ-010 | What are currency sources, sinks, and NPC vendor policies? | Determines inflation and economic stability | Phase 5 |
| OQ-011 | Is there global market visibility, regional markets, or direct shop discovery only? | Determines merchant and location value | Phase 6 |
| OQ-012 | What are building placement, decay, land ownership, and maintenance rules? | Determines city viability and clutter control | Phase 8 |
| OQ-013 | How are cities founded, governed, taxed, and dissolved? | Determines player society structure | Phase 8 |
| OQ-014 | What is the PvP model? | Determines risk, crime, transport, cities, insurance, and conflict | Phase 9 |
| OQ-015 | What are death, injury, item-loss, and recovery rules? | Determines combat economy and frustration level | Phase 9 |
| OQ-016 | How are medical and entertainment services made valuable without becoming mandatory waiting systems? | Determines service-profession health | Phase 9 |
| OQ-017 | What transport methods exist, and what economic friction should they preserve? | Determines regional economics | Phase 10 |
| OQ-018 | Are alternate characters allowed, and what anti-multi-account protections are needed? | Determines specialisation and monopoly resistance | Account/economy policy |
| OQ-019 | What is the contract, escrow, and organisation-asset protection model? | Determines trust at social scale | Phase 8 |
| OQ-020 | Which systems are required for the first playable prototype versus the first persistent alpha? | Determines development scope | Production planning |

---

# 34. Glossary

| Term | Meaning |
|---|---|
| **Attribute** | A measurable material property of a resource, used by schematics and quality calculations. |
| **Brand** | A player or organisation identity associated with products, services, or businesses. |
| **Crafting** | Player-directed creation of items using schematics, materials, skill, and possible experimentation. |
| **Extraction** | Acquiring raw resources from the world through gathering, tools, equipment, or installations. |
| **Factory** | A structure or system that performs constrained repeatable production from approved schematics. |
| **Gold Rush** | A temporary player-driven economic event caused by discovery of an unusually valuable resource opportunity. |
| **Interdependence** | The design principle that important professions and activities should require cooperation with other players. |
| **Intermediate Good** | A processed material or component used as input to create another product. |
| **Provenance** | Persistent record of an item's creator, production history, materials, and modifications as applicable. |
| **Resource Instance** | A specific temporary resource spawn with individual identity, attributes, location, and lifecycle. |
| **Schematic** | A formal specification defining the rules for crafting or manufacturing an item. |
| **Service Profession** | A role providing value directly to other players rather than primarily producing tradable items. |
| **Specialisation Budget** | A proposed limit on how many high-tier skill paths a character can maintain simultaneously. |
| **Virtual Society** | The network of player-created relationships, businesses, institutions, reputations, and communities within the persistent world. |

---

# The Central Design Test

Every future system should be tested against this question:

> **Does this system increase meaningful player agency within a living, interconnected virtual society?**

A system should be questioned if it:

- makes other players unnecessary,
- makes specialisation meaningless,
- removes scarcity,
- removes identity,
- eliminates meaningful choice,
- turns the economy into a spreadsheet,
- allows unlimited individual automation,
- forces players into predetermined classes,
- or makes the world feel like a collection of menus rather than a place.

The game succeeds when the player stops thinking:

> "What does the game want me to do?"

and starts thinking:

> **"What do I want to become in this world?"**

---

*End of Master Game Design Document v1.0*
