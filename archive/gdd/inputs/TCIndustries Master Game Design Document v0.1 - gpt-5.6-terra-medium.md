# TCIndustries Master Game Design Document  
**Document version:** Initial Master GDD v0.1  
**Project:** *Tiwakings Craftworld Industries* (TCIndustries)  
**Document purpose:** Establish the canonical design architecture, philosophy, system boundaries, initial rules direction, unresolved questions, and future design sequence for TCIndustries.
**Author:** gpt-5.6-terra-medium (per filename, unverified)

---

# 1. Executive Summary

TCIndustries is a persistent sandbox MMORPG centered on living in a player-driven society. It is inspired heavily by the original *Star Wars Galaxies* virtual-world model: dynamic resources, interdependent professions, player-made goods, recognizable crafters, player commerce, housing, settlements, social services, and emergent economic life.

The intended player fantasy is not “become the chosen hero.” It is:

> Become somebody known in a persistent world.

A player may become a resource prospector, farmer, medical practitioner, entertainer, master craftsperson, merchant, trader, architect, combatant, researcher, city organizer, manufacturer, transport operator, or a hybrid of several identities.

The game’s central loop is:

> Discover resources → acquire and process materials → craft goods or provide services → specialize → establish reputation → participate in trade, society, and regional economies.

TCIndustries must avoid becoming an automation-first industrial game, a conventional class MMO, or a quest-driven theme park. Industry, production, combat, and economic simulation exist to create meaningful social roles and player interdependence.

This document is an initial canonical framework. It deliberately identifies many mechanics as **PROPOSED**, **TBD**, or **DEFERRED** rather than inventing premature false precision.

---

# 2. Status Definitions

| Status | Meaning |
|---|---|
| **LOCKED** | Explicitly decided by the project and authoritative unless formally revised. |
| **PROPOSED** | Recommended design direction pending approval and detailed specification. |
| **TBD** | Unresolved decision requiring future design work. |
| **PROTOTYPE** | Existing prototype behavior or experiment; not canonical. |
| **DEFERRED** | Recognized future system that should not block current design or implementation. |

---

# 3. Vision

## 3.1 Game Vision

**LOCKED — VIS-001**

TCIndustries is a persistent multiplayer virtual world in which players create their own economic, social, occupational, and reputational identities.

The world should feel inhabited not only by NPCs and systems, but by player-created institutions:

- shops,
- workshops,
- supply networks,
- social venues,
- cities,
- organizations,
- specialist communities,
- commercial districts,
- brands,
- player reputations.

The game’s primary content is the interaction between players, world systems, and player-created consequences.

## 3.2 Player Fantasy

**LOCKED — VIS-002**

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
- becoming an expert in a narrow specialization.

## 3.3 Target Experience

**LOCKED — VIS-003**

The intended experience is social, emergent, persistent, and interdependent.

The desired player thought is:

> “This world changed because of what players did, and other players know who matters here.”

## 3.4 Anti-Goals

**LOCKED — VIS-004**

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

# 4. Design Pillars

## 4.1 Discovery and the Gold Rush

**LOCKED — PIL-001**

Resources are temporary, geographically distributed, variable in quality, and economically meaningful.

Exceptional resource discoveries should create player-driven events:

- prospecting activity,
- mining camps,
- trade opportunities,
- speculation,
- guild coordination,
- regional traffic,
- competitive supply chains,
- demand for extractors, transport, processing, and manufacturing.

### Historical SWG Inspiration

Original SWG used temporary resource spawns with variable statistics, regional availability, and resource surveying. This created memorable periods where players sought unusually valuable resources.

### TCIndustries Adaptation

**PROPOSED — RES-001**

TCIndustries will retain the principle of dynamic, temporary, statistically variable resource occurrences, while improving information clarity, anti-monopoly controls, and long-term economic observability.

## 4.2 Identity and Reputation

**LOCKED — PIL-002**

Player identity must be meaningful, visible, and socially consequential.

Products, services, businesses, and buildings should preserve meaningful association with their creator, provider, organization, or brand.

## 4.3 Interdependence

**LOCKED — PIL-003**

No major profession category should be fully self-sufficient at high effectiveness.

Interdependence should arise through:

- distinct skills,
- specialized resources,
- intermediate components,
- service needs,
- regional logistics,
- knowledge gaps,
- reputation,
- production capacity limits,
- social organization.

Interdependence must create opportunity rather than coercive inconvenience.

## 4.4 Player-Created Society

**LOCKED — PIL-004**

Players must be able to form organizations, businesses, settlements, social hubs, trade networks, and reputational institutions with meaningful world presence.

---

# 5. Player Experience and Core Loops

## 5.1 Primary Player Loops

**PROPOSED — LOOP-001**

The game supports several interconnected loops rather than a single mandatory progression path.

| Loop | Core Activities | Social Outcome |
|---|---|---|
| Discovery | Survey, explore, identify resources | Information trade, rushes, regional knowledge |
| Extraction | Harvest, operate extraction equipment, gather | Resource supply and logistics |
| Processing | Refine, convert, prepare materials | Intermediate-goods markets |
| Crafting | Build items, experiment, specialize | Brand identity and product differentiation |
| Retail | Stock vendors, price goods, market services | Local commerce and commercial districts |
| Service | Heal, entertain, repair, transport, construct | Recurring player interaction |
| Combat | Fight threats, protect territory, acquire rewards | Demand for equipment, medicine, consumables |
| Social Governance | Form associations, cities, events, policies | Community identity and institutions |

## 5.2 Non-Combat Viability

**LOCKED — LOOP-002**

A player must be able to pursue a meaningful long-term career without engaging in combat as their primary activity.

Combat may offer demand, risk, materials, access, and opportunities, but must not be the sole source of progression, wealth, or social status.

## 5.3 Personal Progression Philosophy

**LOCKED — LOOP-003**

Progression should expand choices, capability, specialization, reputation, and access to complex roles. It should not force all players through one linear power ladder.

---

# 6. World Philosophy and Structure

## 6.1 Persistent World

**LOCKED — WRLD-001**

TCIndustries is a persistent shared world. Player actions should have durable consequences where practical, including:

- built structures,
- business inventories,
- market prices,
- resource depletion,
- city development,
- organizational ownership,
- local reputation,
- product provenance.

## 6.2 Geographic Structure

**PROPOSED — WRLD-002**

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

## 6.3 Settlements, Cities, and Wilderness

**PROPOSED — WRLD-003**

The world should contain:

1. **NPC settlements**  
   Stable service, trade, tutorial, and baseline social locations.

2. **Player settlements/cities**  
   Player-founded areas with governance, zoning, commerce, civic structures, and community identity.

3. **Wilderness regions**  
   Exploration, harvesting, combat, resource discovery, and construction opportunities.

4. **Economic corridors**  
   Roads, transport routes, ports, stations, or equivalent systems connecting markets and resource regions.

Exact city-formation rules are **TBD**.

## 6.4 Time and Environmental Simulation

**PROPOSED — WRLD-004**

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

## 6.5 Transportation

**PROPOSED — WRLD-005**

Transportation should create meaningful geography without turning routine trade into excessive travel friction.

The system should support a spectrum:

- local movement,
- personal vehicles or mounts,
- public transit,
- organization-operated transport,
- freight logistics,
- fast travel under constraints.

Fast travel must not eliminate the economic relevance of location, local markets, or transport services.

Exact transportation mechanics are **TBD**.

---

# 7. Character, Identity, and Ownership

## 7.1 Character Identity

**LOCKED — PLR-001**

Each player character has a persistent identity that can accumulate:

- name,
- appearance,
- ownership,
- skill history,
- profession specializations,
- reputation,
- business affiliation,
- organization membership,
- product/service provenance,
- social relationships.

## 7.2 Character Creation

**PROPOSED — PLR-002**

Character creation should prioritize identity and aesthetic expression rather than granting permanent mechanical superiority.

Initial mechanical differences, if any, must not trap players into inferior long-term choices.

Species, ancestry, culture, and starting location design are **DEFERRED** pending setting development.

## 7.3 Ownership

**LOCKED — PLR-003**

Players and organizations must be able to own meaningful persistent assets, including:

- items,
- resources,
- structures,
- vendors,
- business inventory,
- manufacturing permissions,
- organization property,
- brands,
- contracts where implemented.

Ownership systems must include safeguards against fraud, abandoned assets, and exploitative transfer mechanics.

---

# 8. Skills, Professions, and Progression

## 8.1 Profession Architecture

**LOCKED — PROF-001**

TCIndustries uses a flexible skill-based profession architecture rather than rigid immutable classes.

A profession is an emergent identity resulting from acquired skills, specialization choices, equipment, social reputation, and player behavior.

## 8.2 Skill Domains

**PROPOSED — PROF-002**

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

## 8.3 Specialization

**PROPOSED — PROF-003**

Players may develop broad competency across domains but should require meaningful investment to reach high-tier specialization.

A master weaponsmith, for example, should not automatically also be a top medical practitioner, architect, merchant, and combatant.

Specialization creates markets for experts.

## 8.4 Respecialization

**LOCKED — PROF-004**

Players must be able to respecialize. Permanent character traps are contrary to sandbox identity development.

**PROPOSED implementation principle:** Respecialization should involve meaningful opportunity cost, cooldown, retraining time, loss of active specialization capacity, or economic cost—but should not erase player identity, item provenance, business ownership, or reputation history.

Exact respecialization rules are **TBD**.

## 8.5 Progression Caps

**PROPOSED — PROF-005**

The game should have a capability cap or specialization budget that prevents one character from mastering every economically important role at once.

The cap must preserve:

- experimentation,
- casual participation,
- hybrid roles,
- respecialization,
- player autonomy.

The exact cap model is **TBD**.

---

# 9. Resource System

## 9.1 Resource Design Goal

**LOCKED — RES-002**

The resource system exists to create discovery, scarcity, regional trade, product differentiation, temporary opportunity, and social/economic events.

Resources must not function solely as generic crafting currencies.

## 9.2 Resource Taxonomy

**PROPOSED — RES-003**

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
- fibers,
- woods,
- fuels,
- gases,
- biological materials,
- water and liquids,
- rare technical materials.

These are taxonomy examples, not a locked content list.

## 9.3 Resource Instances

**PROPOSED — RES-004**

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

> “Kavren Ferritic Alloy, discovered in the North Ember Plains, cycle 184.”

## 9.4 Resource Attributes

**PROPOSED — RES-005**

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

Exact initial attribute library is **TBD**.

## 9.5 Discovery and Surveying

**PROPOSED — RES-006**

Prospecting should be an active information-gathering profession.

Players may use surveying tools, local knowledge, maps, trade channels, and organizations to identify resource occurrences.

Discovery information may be:

- private,
- shared with an organization,
- sold,
- published,
- made public through market behavior,
- inferred from regional activity.

The game should permit information asymmetry, but must avoid allowing a small group to permanently monopolize all useful knowledge.

## 9.6 Temporal Availability and Depletion

**PROPOSED — RES-007**

Resource spawns should be temporary. Individual extraction sites may also be depleted or become inefficient over time.

Two separate concepts must be retained:

1. **Spawn lifecycle:** The resource exists in the world for a limited period.
2. **Local extraction state:** A specific location may become exhausted, congested, inaccessible, or less efficient.

Exact durations, depletion curves, and respawn rules are **TBD** and require economic simulation.

## 9.7 Extraction

**PROPOSED — RES-008**

Extraction should involve different scales:

- manual or personal gathering,
- small-scale equipment,
- stationary extraction installations,
- organization-supported extraction,
- specialized hazardous extraction.

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

## 9.8 Processing

**PROPOSED — RES-009**

Raw resources may be transformed into processed materials and intermediate goods.

Processing should create new specialization and trade opportunities but should not require excessive meaningless conversion steps.

A processing chain is justified only if it introduces at least one of:

- distinct skill role,
- location decision,
- quality tradeoff,
- resource transformation,
- economic market,
- logistical consequence.

---

# 10. Crafting System

## 10.1 Crafting Philosophy

**LOCKED — CRFT-001**

Crafting must produce meaningful product variation, player identity, and market differentiation.

Crafting is not merely a recipe-unlock treadmill.

A product’s desirability should depend on some combination of:

- crafter skill,
- specialization,
- schematic quality,
- resource attributes,
- component quality,
- experimentation choices,
- cosmetic customization,
- provenance,
- reliability,
- reputation.

## 10.2 Schematics

**PROPOSED — CRFT-002**

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
- optional customization slots,
- durability rules,
- manufacture eligibility.

Schematics should reference categories and attributes where possible, not hard-coded resource names.

## 10.3 Crafting Process

**PROPOSED — CRFT-003**

Crafting should include at least:

1. schematic selection;
2. input and component selection;
3. material validation;
4. assembly;
5. experimentation or refinement where allowed;
6. outcome calculation;
7. item creation and provenance recording.

The exact interaction model is **TBD**. It may involve a structured decision process rather than a timing minigame.

## 10.4 Experimentation

**PROPOSED — CRFT-004**

Experimentation allows skilled crafters to influence item properties within schematic-defined limits.

Potential outcomes include:

- higher performance,
- increased durability,
- lower weight,
- improved efficiency,
- improved appearance,
- increased reliability,
- reduced maintenance requirements.

Experimentation should include risk or constrained tradeoffs. It must not be a deterministic “always maximize every statistic” button.

Exact probability formulas are **TBD**.

## 10.5 Product Quality

**PROPOSED — CRFT-005**

Item quality should result from a combination of:

- schematic baseline,
- crafter relevant skill,
- resource quality and attributes,
- components,
- experimentation outcome,
- manufacturing quality where applicable,
- optional service or calibration stage.

Quality must be intelligible enough for players to make informed choices.

## 10.6 Provenance and Creator Identity

**LOCKED — CRFT-006**

Crafted items should retain provenance sufficient to support reputation and brand identity.

At minimum, item provenance should support:

- original creator,
- producing organization if applicable,
- schematic lineage/version,
- major material identity where appropriate,
- creation time,
- item quality,
- modification history where relevant.

Privacy, pseudonymous brands, and resale handling are **TBD**.

---

# 11. Manufacturing and Automation

## 11.1 Manufacturing Goal

**LOCKED — MFG-001**

Manufacturing exists to support production scale, supply chains, and businesses without replacing player specialization and social interaction.

## 11.2 Factory Role

**PROPOSED — MFG-002**

Factories may produce repeatable items from approved schematics and supplied inputs. They should be appropriate for:

- standard components,
- bulk consumables,
- common equipment,
- intermediate goods,
- established product lines.

Factories should not replace the value of master-crafted, experimental, limited, or custom goods.

## 11.3 Manufacturing Constraints

**PROPOSED — MFG-003**

Manufacturing should be constrained by meaningful requirements:

- schematic authorization,
- input materials,
- component availability,
- factory capacity,
- maintenance,
- location,
- fuel/power where appropriate,
- queue time,
- production quality limits,
- ownership limits,
- organization coordination.

Exact factory limits are **TBD**.

## 11.4 Social Scale Principle

**LOCKED — MFG-004**

Increasing economic scale must increasingly require organizations, trade, logistics, specialization, reputation, and coordination—not merely additional personal factories.

## 11.5 Manufacturing Quality Ceiling

**PROPOSED — MFG-005**

Manufactured products should generally have a lower quality ceiling or reduced customization compared with direct specialist crafting, unless a later high-tier industrial specialization is deliberately designed to alter this relationship.

This preserves the relevance of artisans while allowing practical volume production.

---

# 12. Items

## 12.1 Item Categories

**PROPOSED — ITEM-001**

Items may include:

- equipment,
- weapons,
- armor,
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

## 12.2 Durability and Degradation

**PROPOSED — ITEM-002**

Many functional items should have durability, maintenance needs, degradation, or finite consumable use.

This creates replacement demand and supports the economy. However, degradation must not feel like arbitrary punishment.

Durability should be appropriate to category:

- consumables: intentionally expended;
- tools and equipment: gradual wear;
- buildings/factories: maintenance;
- cosmetics: generally no functional decay;
- high-value legacy items: special rules **TBD**.

## 12.3 Repair

**PROPOSED — ITEM-003**

Repair should preserve the value of repair services and material demand.

Possible rules include:

- repair consumes materials or repair kits;
- repair requires appropriate skill for high-value items;
- repeated repair may reduce maximum durability;
- specialist repair may preserve more value than generic repair.

Exact repair model is **TBD**.

## 12.4 Customization

**PROPOSED — ITEM-004**

Items should support appropriate cosmetic and naming customization, subject to anti-abuse moderation rules.

Customization should help support identity, commerce, gifting, city culture, and player brands.

---

# 13. Economy

## 13.1 Economic Philosophy

**LOCKED — ECO-001**

The economy must be materially connected to the world. Resources, labor, transport, crafting, consumption, degradation, services, and regional differences should matter.

The game should not rely solely on abstract currency generation and vendor sinks.

## 13.2 Currency

**PROPOSED — ECO-002**

The game should use a primary currency for player trade, NPC transactions, taxes, maintenance, and selected sinks.

Currency naming, source rates, and sink design are **TBD**.

## 13.3 Supply and Demand

**PROPOSED — ECO-003**

Player pricing should be broadly free within anti-exploit safeguards.

Supply and demand should be shaped by:

- temporary resources,
- item degradation,
- combat demand,
- city population,
- travel and location,
- industrial input demand,
- fashion and customization,
- player events,
- new schematics and technology,
- regional specialization.

## 13.4 Economic Sinks

**PROPOSED — ECO-004**

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
- cosmetic customization,
- organization upkeep.

Sinks must not disproportionately punish low-income or casual players.

## 13.5 NPC Economic Role

**PROPOSED — ECO-005**

NPC vendors should provide baseline accessibility and emergency supply, but should not invalidate player commerce for economically meaningful goods.

NPC buy orders, if used, must be carefully limited to avoid creating infinite-risk-free farming loops.

## 13.6 Market Information

**PROPOSED — ECO-006**

Market information should be useful but not perfectly frictionless.

Players should be able to discover prices through shops, market boards, trade networks, contracts, and reputation. Full global instant price transparency may reduce regional economies and merchant roles.

Exact market-information model is **TBD**.

---

# 14. Retail and Commerce

## 14.1 Player Shops

**LOCKED — RET-001**

Player-run retail must be a major supported playstyle.

Players should be able to establish shops, stock goods, set prices, curate presentation, and develop repeat customers.

## 14.2 Vendors and Storefronts

**PROPOSED — RET-002**

Retail may use vendor systems located in player-owned or organization-owned structures.

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

## 14.3 Location and Commercial Districts

**PROPOSED — RET-003**

Location should influence commerce through visibility, convenience, city foot traffic, regional supply, transport access, events, and nearby services.

However, remote or specialist businesses must remain viable through directories, transport, reputation, delivery, or trade networks.

## 14.4 Merchant Value

**PROPOSED — RET-004**

Merchant-oriented skills should provide meaningful advantages beyond “placing a vendor,” such as:

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

# 15. Services

## 15.1 Service Philosophy

**LOCKED — SERV-001**

Non-item player services are first-class economic and social gameplay, not secondary flavor.

## 15.2 Medical Services

**PROPOSED — SERV-002**

Medical specialists may provide healing, recovery, treatment, stabilization, rehabilitation, medical consumables, and possibly temporary performance support.

Medical gameplay must create real reasons to visit or hire players while avoiding mandatory downtime that prevents normal play.

## 15.3 Entertainment and Social Services

**PROPOSED — SERV-003**

Entertainers may provide social gathering value, morale/recovery benefits, temporary buffs, event hosting, venue identity, and cultural expression.

The system must avoid reducing entertainers to passive buff dispensers.

## 15.4 Repair, Construction, Research, and Transport

**PROPOSED — SERV-004**

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
- specialist customization.

Each service must be assessed for whether it produces meaningful interaction rather than artificial gatekeeping.

---

# 16. Combat

## 16.1 Combat Role

**LOCKED — CMBT-001**

Combat is an important world activity but not the sole or dominant measure of player value.

Combat should generate demand for:

- weapons,
- armor,
- tools,
- medical services,
- consumables,
- repairs,
- transport,
- specialist equipment,
- crafted components.

## 16.2 Combat Design Direction

**DEFERRED — CMBT-002**

Detailed combat mechanics, damage models, targeting, enemy AI, PvP policy, faction systems, and encounter design require a dedicated combat specification.

## 16.3 PvP

**TBD — CMBT-003**

The project has not yet established whether TCIndustries includes open PvP, consensual PvP, territorial conflict, criminal systems, faction warfare, or fully PvE gameplay.

This decision strongly affects cities, resource access, transport, item loss, insurance, organizations, and economic risk.

---

# 17. Buildings, Housing, and Cities

## 17.1 Building Ownership

**LOCKED — BLD-001**

Players and organizations should be able to own meaningful structures, including housing, shops, workshops, factories, and civic buildings.

## 17.2 Building Functions

**PROPOSED — BLD-002**

Structures may provide:

- storage,
- habitation,
- social space,
- retail space,
- manufacturing capacity,
- workshop bonuses,
- city services,
- organization administration,
- aesthetic identity,
- event hosting.

## 17.3 Maintenance and Limits

**PROPOSED — BLD-003**

Buildings should require upkeep and obey ownership, zoning, placement, and population limits.

The purpose is to prevent abandoned-world clutter, land monopolization, and unlimited asset accumulation.

Exact upkeep, decay, reclamation, and placement rules are **TBD**.

## 17.4 Player Cities

**PROPOSED — CITY-001**

Player cities should be player-founded social and economic institutions, not merely housing clusters.

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

# 18. Social Systems and Organizations

## 18.1 Organizations

**LOCKED — SOC-001**

Players must be able to form persistent organizations that support shared identity, ownership, cooperation, and economic coordination.

The exact terminology—guild, company, association, cooperative, syndicate, clan, or other—is setting-dependent and **TBD**.

## 18.2 Organization Capabilities

**PROPOSED — SOC-002**

Organizations may support:

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

## 18.3 Reputation

**PROPOSED — SOC-003**

Reputation should be primarily social and evidence-based, supported by system visibility rather than reduced to a universal numeric score.

Possible evidence includes:

- item provenance,
- customer reviews with safeguards,
- transaction history,
- city affiliation,
- organization affiliation,
- service records,
- contracts completed,
- public achievements,
- local reputation.

A single global reputation number is discouraged because it can become gameable, reductive, and socially coercive.

---

# 19. Emergent Gameplay Targets

**LOCKED — EMRG-001**

The following are desired emergent outcomes and should be used as design tests:

1. A rare, high-quality resource appears and creates a temporary prospecting rush.
2. A known crafter becomes sought after because their products consistently perform well.
3. A city develops an identity as a medical, industrial, cultural, or trade center.
4. A combat conflict or PvE event increases demand for particular equipment and medical supplies.
5. A shortage of a component causes merchants and manufacturers to form new supply relationships.
6. A player organization coordinates extraction, processing, production, retail, and logistics.
7. A remote specialist gains customers through reputation rather than only geographic luck.
8. Player-created events generate trade, travel, entertainment, and social interaction.
9. Resource geography causes regional economies without making participation inaccessible.
10. A player can become socially important without being a top combat player.

---

# 20. Content Expansion Framework

## 20.1 Data-Driven Design Principle

**LOCKED — EXP-001**

Core systems must be defined through extensible abstractions rather than hard-coded isolated content.

Future additions should be able to introduce:

- new resource categories and attributes,
- schematics,
- item types,
- professions,
- buildings,
- services,
- regions,
- organizations,
- NPC roles,
- transportation options,
- research paths,
- world events.

## 20.2 Expansion Requirements

**PROPOSED — EXP-002**

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

# 21. Balance Philosophy

## 21.1 Meaningful Imperfection

**PROPOSED — BAL-001**

The economy should not converge permanently on one universally optimal product, profession, resource, or region.

Variation should arise from:

- temporary resource quality,
- changing availability,
- different player skills,
- regional demand,
- item niches,
- customization,
- logistics,
- reputation,
- situational combat needs.

## 21.2 Accessibility Versus Dependency

**PROPOSED — BAL-002**

Interdependence must not become hostage gameplay.

Players should have access to basic self-sufficiency and emergency NPC support, while premium efficiency, quality, specialization, and convenience come from player relationships.

## 21.3 Casual and Long-Term Players

**PROPOSED — BAL-003**

Casual players should be able to participate meaningfully in discovery, gathering, local trade, services, and social life.

Long-term players should gain specialization, reputation, organizational influence, and operational complexity—not unchecked combat or economic dominance.

---

# 22. Exploit, Fraud, and Economic Stability Considerations

## 22.1 General Principle

**LOCKED — SAFE-001**

Systems involving trade, ownership, organizations, automation, and reputation must be designed with fraud and exploit resistance as first-class requirements.

## 22.2 Risks Requiring Dedicated Design

**TBD — SAFE-002**

The following require detailed rule specifications:

- duplicate item prevention,
- currency generation limits,
- NPC vendor arbitrage,
- abandoned structure cleanup,
- land monopolization,
- resource spawn monopolization,
- multi-account extraction abuse,
- automated market manipulation,
- vendor bait-and-switch,
- organization asset theft,
- contract fraud,
- price manipulation,
- griefing through city governance,
- player review abuse,
- service-payment disputes,
- inactive owner handling.

## 22.3 Economic Monitoring

**PROPOSED — SAFE-003**

The future live game should monitor economic health through simulation and telemetry, including:

- currency creation and destruction,
- resource extraction volumes,
- item production and destruction,
- price distributions,
- market concentration,
- new-player access to essentials,
- regional supply imbalance,
- organization asset concentration,
- factory utilization,
- abandoned property rates.

---

# 23. Testing Requirements and Design Invariants

## 23.1 Core Invariants

**PROPOSED — TEST-001**

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

## 23.2 Simulation Requirements

**PROPOSED — TEST-002**

Before major system rules are locked, simulations should test:

- resource spawn frequency and gold-rush behavior;
- high-quality resource concentration;
- new-player access to basic materials;
- factory output versus artisan relevance;
- price inflation and currency sink adequacy;
- organization concentration and monopoly risk;
- durability/replacement demand;
- time-to-specialization;
- viability of non-combat careers;
- viability of small independent shops;
- regional market behavior;
- effect of travel costs and fast travel;
- multi-account and unattended automation incentives.

## 23.3 Acceptance Criteria

**PROPOSED — TEST-003**

A system should not be promoted to LOCKED until it has:

1. a defined player goal;
2. explicit inputs, state, outputs, and failure cases;
3. identified dependencies;
4. exploit considerations;
5. simulation or prototype validation where appropriate;
6. clear player-visible outcomes;
7. consistency with all four design pillars.

---

# 24. Historical SWG Relationship

## 24.1 Inspiration Policy

**LOCKED — SWG-001**

TCIndustries is heavily inspired by the virtual-world philosophy and selected systemic lessons of original SWG, but it is not a Star Wars game and must not reproduce protected Star Wars intellectual property.

## 24.2 Adaptation Rule

**LOCKED — SWG-002**

When adopting a historical SWG-inspired mechanic, design documentation must distinguish:

- **Historical SWG behavior:** What SWG did, to the extent accurately known.
- **TCIndustries adaptation:** What TCIndustries chooses to retain or alter.
- **New TCIndustries invention:** A system created specifically for TCIndustries.

Historical similarity does not automatically make a mechanic canonical.

---

# 25. Systems Requiring Deeper Design

The following systems are recognized but require dedicated canonical specifications before implementation can be considered authoritative.

| Priority | System | Reason |
|---|---|---|
| 1 | Resource spawning and surveying | Foundation of gold-rush economy and production chain |
| 2 | Skill budget and specialization | Determines interdependence and player freedom |
| 3 | Crafting quality and experimentation | Determines product identity and crafter reputation |
| 4 | Item durability, repair, and destruction | Determines replacement demand and economic health |
| 5 | Extraction and automation limits | Prevents industrial-simulation drift and monopoly |
| 6 | Manufacturing rules | Defines artisan versus factory relationship |
| 7 | Currency sources, sinks, and NPC market role | Required for economic stability |
| 8 | Shops, vendors, and discovery | Required for player commerce viability |
| 9 | Buildings, land rights, maintenance, and decay | Required for persistent player society |
| 10 | Cities and organization governance | Required for player-created institutions |
| 11 | Combat and PvP policy | Affects risk, demand, territorial systems, and social rules |
| 12 | Medical and entertainment service loops | Required for non-combat service professions |
| 13 | Transportation and regional market design | Determines whether geography matters |
| 14 | Reputation and contract systems | Required for trust at economic scale |

---

# 26. Recommended Design Sequence

## Iteration 1: Economic Foundations

1. Lock resource taxonomy.
2. Lock resource spawn lifecycle model.
3. Define surveying and discovery rules.
4. Define extraction tiers and automation limits.
5. Define processing categories.
6. Define core crafting schematic model.
7. Define quality calculation inputs and outputs.
8. Define item provenance rules.

## Iteration 2: Player Interdependence

1. Define skill domains and specialization budget.
2. Define crafting profession families.
3. Define resource, processing, and manufacturing dependencies.
4. Define durability, repair, consumables, and replacement demand.
5. Define basic medical and entertainment service loops.

## Iteration 3: Commerce and Society

1. Define currency sources and sinks.
2. Define vendor, shop, and market-discovery rules.
3. Define building ownership and maintenance.
4. Define organizations and permission systems.
5. Define city foundation and governance.

## Iteration 4: World Risk and Expansion

1. Define combat loop.
2. Decide PvP policy.
3. Define transport and logistics.
4. Define world events and environmental systems.
5. Define research, technology, and advanced content progression.

---

# 27. Status Register

| ID | System/Decision | Status | Rationale | Dependencies |
|---|---|---|---|---|
| VIS-001 | Persistent player-driven virtual society | LOCKED | Core project objective | All systems |
| VIS-002 | Citizen identity over hero narrative | LOCKED | Core player fantasy | Progression, social, economy |
| VIS-003 | Emergent and interdependent target experience | LOCKED | Central world philosophy | All systems |
| VIS-004 | Anti-goals: no theme park, class MMO, industrial tycoon drift | LOCKED | Prevents design drift | All systems |
| PIL-001 | Dynamic resource discovery/gold rush pillar | LOCKED | Explicit project pillar | Resources, economy |
| PIL-002 | Identity and reputation pillar | LOCKED | Explicit project pillar | Crafting, retail, services |
| PIL-003 | Interdependence pillar | LOCKED | Explicit project pillar | Skills, economy, services |
| PIL-004 | Player-created society pillar | LOCKED | Explicit project pillar | Buildings, cities, organizations |
| LOOP-001 | Multi-loop player experience model | PROPOSED | Supports sandbox participation | All major systems |
| LOOP-002 | Non-combat long-term viability | LOCKED | Explicit project direction | Professions, economy |
| LOOP-003 | Progression expands options, not linear power | LOCKED | Core sandbox philosophy | Skills |
| WRLD-001 | Persistent shared world | LOCKED | Required for living-world premise | World, ownership |
| WRLD-002 | Regionally differentiated geography | PROPOSED | Supports trade and discovery | Resources, transport |
| WRLD-003 | NPC settlements, player cities, wilderness | PROPOSED | World structure recommendation | Buildings, cities |
| WRLD-004 | Persistent time/environment systems | PROPOSED | Supports world life if meaningful | World simulation |
| WRLD-005 | Constrained transportation spectrum | PROPOSED | Preserves geography and commerce | Economy, cities |
| PLR-001 | Persistent character identity | LOCKED | Supports reputation and ownership | Player systems |
| PLR-002 | Cosmetic-first character creation | PROPOSED | Avoids permanent optimization traps | Character creation |
| PLR-003 | Persistent player/organization ownership | LOCKED | Required for society and economy | Buildings, items |
| PROF-001 | Flexible skill-based professions | LOCKED | Explicit project requirement | Progression |
| PROF-002 | Initial skill-domain categories | PROPOSED | Initial profession architecture | Content design |
| PROF-003 | Meaningful high-tier specialization | PROPOSED | Supports interdependence | Skill budget |
| PROF-004 | Respecialization allowed | LOCKED | Explicit project philosophy | Progression |
| PROF-005 | Capability cap/specialization budget | PROPOSED | Prevents universal mastery | Skill design |
| RES-002 | Resources create discovery and scarcity | LOCKED | Pillar implementation | Economy, crafting |
| RES-003 | Extensible resource taxonomy | PROPOSED | Supports content expansion | Data definitions |
| RES-004 | Temporary resource instances | PROPOSED | Supports gold-rush behavior | Surveying, extraction |
| RES-005 | Attribute-driven resources | PROPOSED | Supports product differentiation | Crafting |
| RES-006 | Surveying as information profession | PROPOSED | Supports discovery and trade | Resources, skills |
| RES-007 | Spawn lifecycle and local depletion | PROPOSED | Creates temporal opportunity | Extraction, economy |
| RES-008 | Constrained extraction automation | PROPOSED | Prevents unlimited solo scale | Buildings, manufacturing |
| RES-009 | Meaningful processing chains | PROPOSED | Supports intermediate markets | Crafting, economy |
| CRFT-001 | Crafting creates differentiated products | LOCKED | Core identity/reputation requirement | Resources, skills |
| CRFT-002 | Extensible schematic structure | PROPOSED | Expansion-friendly crafting | Resources, items |
| CRFT-003 | Structured crafting process | PROPOSED | Implementation basis | Schematics, skills |
| CRFT-004 | Constrained experimentation | PROPOSED | Supports mastery and variation | Crafting quality |
| CRFT-005 | Multi-input quality calculation | PROPOSED | Supports resource/crafter importance | Resources, skills |
| CRFT-006 | Crafted item provenance | LOCKED | Supports reputation and brands | Items, retail |
| MFG-001 | Manufacturing supports rather than replaces players | LOCKED | Anti-drift requirement | Economy, factories |
| MFG-002 | Factories for repeatable production | PROPOSED | Practical scale without full automation | Schematics, buildings |
| MFG-003 | Factory constraints | PROPOSED | Anti-monopoly and social scale | Economy |
| MFG-004 | Scale increasingly requires coordination | LOCKED | Explicit design-scale principle | Organizations, economy |
| MFG-005 | Artisan quality advantage | PROPOSED | Protects specialist relevance | Crafting, factories |
| ITEM-001 | Broad extensible item categories | PROPOSED | Content framework | Crafting |
| ITEM-002 | Category-appropriate degradation | PROPOSED | Replacement demand | Economy, repair |
| ITEM-003 | Repair as meaningful service/material sink | PROPOSED | Supports economy and services | Items, skills |
| ECO-001 | Materially connected player economy | LOCKED | Core world philosophy | All economic systems |
| ECO-002 | Primary currency | PROPOSED | Necessary transaction medium | Economy design |
| ECO-003 | Player-led supply/demand pricing | PROPOSED | Supports trade and merchant play | Retail, resources |
| ECO-004 | Recurring material/currency sinks | PROPOSED | Inflation control | Buildings, items |
| ECO-005 | Limited NPC baseline economy | PROPOSED | Prevents NPC invalidation of players | Economy |
| ECO-006 | Imperfect market information | PROPOSED | Preserves regional/merchant value | Retail, transport |
| RET-001 | Player-run retail as major playstyle | LOCKED | Explicit player-society objective | Buildings, economy |
| RET-002 | Vendor/storefront capabilities | PROPOSED | Retail functionality | Buildings, items |
| RET-003 | Location affects commerce | PROPOSED | Supports geography and cities | Transport, world |
| RET-004 | Merchant specialization value | PROPOSED | Makes commerce a profession | Skills, retail |
| SERV-001 | Services are first-class gameplay | LOCKED | Supports non-combat identities | Skills, economy |
| SERV-002 | Medical service profession | PROPOSED | Interdependence | Combat, items |
| SERV-003 | Entertainment service profession | PROPOSED | Social-hub creation | Social, buffs |
| SERV-004 | Repair/construction/research/transport services | PROPOSED | Service economy | Multiple systems |
| CMBT-001 | Combat supports, not dominates, economy | LOCKED | Core philosophy | Items, services |
| CMBT-002 | Detailed combat design | DEFERRED | Requires dedicated document | PvP policy, items |
| CMBT-003 | PvP policy | TBD | High-impact unresolved decision | Cities, risk, transport |
| BLD-001 | Meaningful player/organization buildings | LOCKED | Player-created society | Ownership, cities |
| BLD-002 | Multi-function structures | PROPOSED | Supports housing/businesses | Buildings |
| BLD-003 | Maintenance and placement limits | PROPOSED | Prevents clutter/monopoly | Economy, land |
| CITY-001 | Player cities as institutions | PROPOSED | Supports society and governance | Buildings, organizations |
| SOC-001 | Persistent player organizations | LOCKED | Required for cooperation | Ownership, cities |
| SOC-002 | Organization roles/assets/permissions | PROPOSED | Coordination tools | Ownership, contracts |
| SOC-003 | Evidence-based reputation | PROPOSED | Supports identity without reductive score | Retail, services |
| EMRG-001 | Emergent gameplay targets | LOCKED | Evaluation framework | All systems |
| EXP-001 | Extensible data-driven content abstractions | LOCKED | Expandability requirement | All systems |
| EXP-002 | Content declaration requirements | PROPOSED | Prevents unbalanced feature additions | Content pipeline |
| BAL-001 | Avoid permanent universal optimum | PROPOSED | Protects economic diversity | Resources, crafting |
| BAL-002 | Accessibility without forced dependency | PROPOSED | Prevents coercive friction | Economy, services |
| BAL-003 | Casual and veteran participation | PROPOSED | Long-term population health | Progression, economy |
| SAFE-001 | Fraud/exploit resistance as first-class concern | LOCKED | Required for persistent economy | All transactional systems |
| SAFE-002 | Detailed fraud protection rules | TBD | Requires system-level design | Ownership, markets |
| SAFE-003 | Live economic monitoring | PROPOSED | Necessary for tuning | Analytics/simulation |
| TEST-001 | Core system invariants | PROPOSED | Supports testable rules | Implementation specs |
| TEST-002 | Economic simulation requirements | PROPOSED | Validates systemic behavior | Simulation |
| TEST-003 | Promotion criteria for LOCKED systems | PROPOSED | Design governance | All systems |
| SWG-001 | SWG inspiration without IP reproduction | LOCKED | Project requirement | Content, setting |
| SWG-002 | Historical/adaptation/invention distinction | LOCKED | Documentation discipline | All specs |

---

# 28. Assumption Register

| ID | Assumption | Status | Consequence if Incorrect |
|---|---|---|---|
| ASM-001 | TCIndustries is intended to be a persistent online multiplayer world rather than session-based gameplay. | Assumption supported by prompt | World persistence and economy design would require revision. |
| ASM-002 | Dynamic resources remain a desired central system rather than optional flavor. | Assumption strongly supported by pillars | Resource, crafting, and economy architecture would change substantially. |
| ASM-003 | Crafted items can meaningfully differ in quality and desirability. | Assumption strongly supported by identity pillar | Crafter reputation model would weaken if false. |
| ASM-004 | The game will support player-owned structures in some form. | Assumption supported by required GDD areas | Cities, retail, and manufacturing must be redesigned if absent. |
| ASM-005 | The setting is original and must avoid Star Wars IP. | Assumption explicitly supported by project statement | Setting/content design must remain distinct from SWG. |
| ASM-006 | A future simulation/testing stage will exist before production implementation is treated as final. | Assumption supported by intended development path | Balance-validation process would need revision. |
| ASM-007 | Some form of player currency will exist. | Assumption introduced for economy coherence | Barter-only or multi-currency alternatives require separate design. |
| ASM-008 | Some items will degrade, be consumed, or require maintenance. | Assumption introduced to support replacement demand | Alternative sink models would be needed if rejected. |
| ASM-009 | Player skill specialization will have a cap or budget. | Assumption introduced to preserve interdependence | Without it, universal mastery risk increases. |
| ASM-010 | Factories and extractors can exist but must be constrained. | Assumption supported by design-scale statement | Manufacturing loop must be adjusted if factory gameplay is excluded. |

---

# 29. Open-Question Register

| ID | Question | Why It Matters | Recommended Design Phase |
|---|---|---|---|
| OQ-001 | What is the setting’s original genre, technology level, visual identity, and world fiction? | Determines resource names, professions, buildings, combat, transport | World foundation |
| OQ-002 | What is the exact skill-cap or specialization-budget model? | Determines interdependence and hybrid viability | Progression |
| OQ-003 | How are skills acquired: use-based, training, points, certifications, mentorship, quests, or hybrid? | Determines progression behavior and exploit risks | Progression |
| OQ-004 | What exact resource attributes exist at launch? | Determines crafting formulas and resource value | Resource system |
| OQ-005 | How long do resource spawns last, and how severe is depletion? | Determines rush behavior and market stability | Resource simulation |
| OQ-006 | How much surveying information is private versus public? | Determines discovery value and monopoly risk | Resource system |
| OQ-007 | What are the exact crafting quality and experimentation formulas? | Determines product differentiation | Crafting |
| OQ-008 | Can crafted items fail, partially fail, or only vary in quality? | Determines risk and usability | Crafting |
| OQ-009 | What is the exact factory quality ceiling and ownership limit? | Determines artisan relevance and automation risk | Manufacturing |
| OQ-010 | What are currency sources, sinks, and NPC vendor policies? | Determines inflation and economic stability | Economy |
| OQ-011 | Is there global market visibility, regional markets, or direct shop discovery only? | Determines merchant and location value | Retail |
| OQ-012 | What are building placement, decay, land ownership, and maintenance rules? | Determines city viability and clutter control | Buildings |
| OQ-013 | How are cities founded, governed, taxed, and dissolved? | Determines player society structure | Cities |
| OQ-014 | What is the PvP model? | Determines risk, crime, transport, cities, insurance, and conflict | Combat |
| OQ-015 | What are death, injury, item-loss, and recovery rules? | Determines combat economy and frustration level | Combat |
| OQ-016 | How are medical and entertainment services made valuable without becoming mandatory waiting systems? | Determines service-profession health | Services |
| OQ-017 | What transport methods exist, and what economic friction should they preserve? | Determines regional economics | World/transport |
| OQ-018 | Are alternate characters allowed, and what anti-multi-account protections are needed? | Determines specialization and monopoly resistance | Account/economy policy |
| OQ-019 | What is the contract, escrow, and organization-asset protection model? | Determines trust at social scale | Economy/social |
| OQ-020 | Which systems are required for the first playable prototype versus the first persistent alpha? | Determines development scope | Production planning |

---

# 30. Glossary

| Term | Meaning |
|---|---|
| **Attribute** | A measurable material property used by schematics and quality calculations. |
| **Brand** | A player or organization identity associated with products, services, or businesses. |
| **Crafting** | Player-directed creation of items using schematics, materials, skill, and possible experimentation. |
| **Extraction** | Acquiring raw resources from the world through gathering, tools, equipment, or installations. |
| **Factory** | A structure or system that performs constrained repeatable production. |
| **Gold Rush** | A temporary player-driven economic event caused by discovery of an unusually valuable resource opportunity. |
| **Intermediate Good** | A processed material or component used to create another product. |
| **Provenance** | Persistent record of an item’s creator, production history, materials, and modifications as applicable. |
| **Resource Instance** | A specific temporary resource spawn with individual identity, attributes, location, and lifecycle. |
| **Schematic** | A formal specification defining the rules for crafting or manufacturing an item. |
| **Service Profession** | A role providing value directly to other players rather than primarily producing tradable items. |
| **Specialization Budget** | A future proposed limit on how many high-tier skill paths a character can maintain simultaneously. |
| **Virtual Society** | The network of player-created relationships, businesses, institutions, reputations, and communities within the persistent world. |

---

## Initial Master GDD Conclusion

The current canonical direction is clear:

- TCIndustries is a living-world sandbox MMORPG.
- Dynamic resource discovery, player-made goods, reputation, services, and player-created society are foundational.
- The game must preserve the social scale of SWG-inspired interdependence.
- Automation and industrial systems are permitted only where they reinforce, rather than replace, player relationships and specialization.
- Detailed mechanics remain intentionally unresolved until validated through focused system specifications and simulation.

The recommended next deliverable is a dedicated **Resource System Canonical Specification**, followed by the **Skill/Specialization Specification** and **Crafting Quality Specification**.