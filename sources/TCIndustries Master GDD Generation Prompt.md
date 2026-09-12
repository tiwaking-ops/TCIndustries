# TCIndustries — Master GDD Generation Prompt

**Author:** Author LLM Unknown (process provenance; filed 2026-09-12)

## Role

Act as a **senior MMORPG game designer, systems designer, virtual-world architect, economic simulation designer, and technical-game-design specification writer**.

Your primary responsibility is to reason about TCIndustries as a coherent living-world simulation rather than as a collection of isolated game mechanics.

You must distinguish carefully between:

- established design decisions,
- proposed design decisions,
- unresolved decisions,
- prototype behaviour,
- deferred systems,
- and assumptions you have introduced yourself.

Do not silently convert an assumption into a rule.

---

# 1. Project

The game is called:

**Tiwakings Craftworld Industries**

Short title:

**TCIndustries**

TCIndustries is intended to be an **expandable replica/spiritual successor inspired heavily by the original Star Wars Galaxies (SWG)**, particularly its virtual-world philosophy, resource economy, crafting, professions, player-driven economy, social interdependence, and emergent gameplay.

The game is not intended to reproduce the Star Wars intellectual property.

The ultimate objective is a persistent virtual world in which players feel that they are **citizens of a living society**, rather than heroes following a predetermined narrative.

---

# 2. Core Design Philosophy

The central design goal is:

> **Create a living virtual world in which players develop their own identity, occupation, reputation, relationships, businesses, and place in society.**

The player should not be treated primarily as:

- the chosen hero,
- the saviour of the world,
- the protagonist of a fixed story,
- or a unit in an economic spreadsheet.

Instead, the player should feel that they **live in the world**.

Possible identities include:

- farmer
- miner
- resource prospector
- artisan
- weaponsmith
- armoursmith
- architect
- doctor
- entertainer
- merchant
- trader
- explorer
- combatant
- service provider
- researcher
- specialist crafter
- shop owner
- city participant
- guild/player-association member

The game should allow these identities to emerge from player choices rather than forcing every player through the same progression.

---

# 3. Four Pillars

The four primary design pillars are:

## 3.1 Discovery and the Gold Rush

The resource system should create genuine excitement around discovery.

Resources should not simply be infinite static commodities.

Resource deposits/spawns should have characteristics such as:

- quality
- rarity
- location
- availability
- temporal existence
- changing supply
- suitability for different recipes
- potentially different resource attributes

A player discovering an exceptional temporary resource spawn should experience the equivalent of a gold rush.

The system should encourage players to react to exceptional discoveries economically and socially.

---

## 3.2 Identity and Reputation

Players should be able to become known for what they do.

A player should be able to develop a reputation for:

- exceptional products
- particular product types
- reliable service
- medical expertise
- entertainment
- resource discovery
- trading
- architecture
- particular shops
- particular brands
- specialist knowledge

The game should support the emergence of recognizable player-created brands and reputations.

The ideal outcome is:

> Players do not merely buy "a product"; they seek out products made by a particular player or organization.

---

## 3.3 Interdependence

The economy and profession system should create genuine reasons for players to need one another.

Examples include:

- combatants needing equipment
- combatants needing medical services
- combatants needing buffs/services
- crafters needing resources
- manufacturers needing specialized components
- merchants needing products
- resource specialists needing customers
- cities needing services and infrastructure
- specialists needing other specialists

Avoid designing every profession as a self-contained mini-game that makes every other profession unnecessary.

Interdependence is a core feature.

---

## 3.4 Player-Created Society

Players should be able to create meaningful social and economic spaces.

Examples include:

- shops
- businesses
- player associations
- guilds
- cities
- commercial districts
- specialist communities
- supply networks
- brands
- reputational networks

The world should provide systems through which these structures can emerge organically.

---

# 4. Design Scale

TCIndustries should reproduce the **original SWG scale of personal participation and social interdependence**, rather than becoming a single-player industrial automation game.

The intended progression is broadly:

**resource discovery/gathering**
→ **processing**
→ **crafting**
→ **specialisation**
→ **retail/service**
→ **reputation**
→ **social/economic interdependence**

A player should NOT simply be able to scale indefinitely by purchasing more factories and personally managing a massive automated industrial empire.

Large-scale economic activity should increasingly require:

- other players,
- organizations,
- specialization,
- trade,
- social relationships,
- logistics,
- reputation,
- and coordination.

Automation may exist, but it must not destroy the social economy.

---

# 5. Sandbox Philosophy

TCIndustries should be fundamentally **sandbox-driven**.

The player should be able to choose their own path.

The game should not require a player to become a combatant.

A player should be able to pursue a meaningful life through non-combat activity.

Likewise, combat players should have reasons to interact with non-combat specialists.

The game should prioritize:

- freedom
- experimentation
- specialization
- respecialization
- social interaction
- economic choice
- emergent gameplay

over:

- rigid classes
- mandatory quest chains
- linear progression
- predetermined heroic narratives.

---

# 6. Expandability

The design must be intentionally expandable.

The game should support future addition of:

- resources
- resource types
- resource attributes
- recipes
- schematics
- products
- professions
- skills
- buildings
- businesses
- services
- NPC systems
- player organizations
- locations
- economic systems
- transportation
- technology
- social systems
- world regions
- gameplay professions

The GDD must therefore define **general rules and abstractions**, rather than hard-coding individual content wherever possible.

For example, define:

> "Resources have a taxonomy and a set of measurable attributes."

rather than merely defining:

> "Iron exists."

Likewise define:

> "Crafting uses schematics whose requirements reference resource attributes."

rather than hard-coding individual recipes.

---

# 7. Existing Prototype

A Seed-2.1-generated **TypeScript + Vite + React prototype** exists.

It may be used as:

- a source of prototype ideas,
- evidence of previously explored mechanics,
- an implementation reference,
- a starting point for experimentation,
- a test harness.

However:

> **The prototype is NOT authoritative.**

The GDD is authoritative.

Prototype behaviour must not be treated as a design decision merely because code already implements it.

The prototype may eventually be discarded and rebuilt.

---

# 8. Intended Development Architecture

The current intended development path is:

**Master GDD**
↓
**Canonical rules/system specifications**
↓
**TypeScript prototype**
↓
**simulation/testing**
↓
**validated/golden-master behaviour**
↓
**future Unreal Engine 5 / C++ implementation**

The prototype architecture currently envisioned is:

React UI
↓
Pure TypeScript API
↓
Domain/game systems
↓
Pure data
↓
JSON/data definitions

Potential domain areas include:

- Resource Engine
- Craft Engine
- Economy Engine
- progression/skill systems
- simulation
- world systems
- inventory
- production
- market systems

Do NOT make these technical architecture requirements part of the GDD. They belong primarily in a future Technical Design Document.

---

# 9. Status System

Every significant design element must receive one of these statuses:

### LOCKED

An explicit project decision that should be treated as authoritative.

### PROPOSED

A design recommendation that appears appropriate but has not yet been formally approved.

### TBD

A genuine unresolved design decision requiring future discussion.

### PROTOTYPE

A behaviour currently implemented or explored in prototype code but not yet accepted as canonical.

### DEFERRED

A known system that will be designed later and should not block current development.

Never silently promote PROPOSED, PROTOTYPE, TBD, or DEFERRED material into LOCKED rules.

---

# 10. GDD Objective

Create a **comprehensive Master Game Design Document** that can eventually be handed to another capable LLM and used as the canonical specification for implementing TCIndustries.

The document must describe the game sufficiently precisely that implementation decisions can be made without repeatedly rediscovering the game's fundamental design philosophy.

The GDD is a **game/rules specification**, NOT a Technical Design Document.

Do not fill it with:

- React implementation details
- TypeScript class structures
- Unreal Engine classes
- networking implementation
- database implementation
- rendering architecture
- UI component architecture

Those should be addressed in a later TDD.

---

# 11. Required GDD Areas

The Master GDD should progressively specify at least:

## A. Vision

- game vision
- player fantasy
- design pillars
- target experience
- design philosophy
- anti-goals

## B. World

- virtual-world philosophy
- world structure
- geography
- settlements
- cities
- wilderness
- resource regions
- population
- environmental systems
- world simulation
- time
- weather if appropriate
- transportation

## C. Player

- player identity
- character creation
- attributes
- skills
- professions
- specialization
- respecialization
- progression
- reputation
- ownership

## D. Resource System

- resource taxonomy
- resource discovery
- resource spawning
- resource quality
- resource attributes
- rarity
- regional variation
- temporal variation
- depletion
- extraction
- harvesting
- processing
- resource markets

## E. Crafting

- crafting philosophy
- schematics
- recipes
- experimentation
- resource quality effects
- product quality
- experimentation risk
- crafting skill
- specialization
- exceptional products
- product identity
- crafted-item provenance

## F. Manufacturing

- manufacturing
- factories
- production
- automation
- production limits
- maintenance
- inputs
- outputs
- intermediate goods
- production chains
- player limits
- social scale

Manufacturing must not automatically turn the game into an industrial tycoon simulation.

## G. Economy

- currency
- pricing
- supply/demand
- markets
- vendors
- shops
- player commerce
- trade
- scarcity
- resource economics
- item degradation
- replacement demand
- economic sinks
- economic exploits
- economic stability

## H. Retail

- player shops
- vendors
- storefronts
- product presentation
- reputation
- customer behaviour
- pricing
- location
- commercial districts

## I. Professions

Define a flexible skill-based profession architecture rather than rigid class trees.

Consider:

- combat
- medical
- entertainment
- crafting
- harvesting
- resource specialization
- architecture
- commerce
- research
- other service professions

## J. Social Systems

- player associations
- guilds
- organizations
- cooperation
- specialization
- reputation
- social dependency
- community structures
- cities
- player-created institutions

## K. Combat

Combat is important but should not consume the entire design.

Define how combat interacts with:

- equipment
- crafting
- medical services
- buffs
- consumables
- economy
- resource demand
- social interdependence.

## L. Services

Define non-item services that players can provide to one another.

Examples:

- medical
- entertainment
- buffs
- transportation
- construction
- repair
- research
- specialist crafting

## M. Items

- item creation
- quality
- durability
- degradation
- repair
- modification
- customization
- ownership
- provenance
- replacement demand

## N. Progression

Progression should support freedom rather than rigid classes.

Define:

- skill acquisition
- skill improvement
- specialization
- respecialization
- costs
- caps
- dependencies
- mastery.

## O. Player Housing / Buildings

- personal structures
- commercial structures
- workshops
- factories
- shops
- cities
- ownership
- maintenance
- location
- limits
- customization

## P. Emergent Gameplay

Explicitly identify systems that should create emergent events such as:

- resource rushes
- market shortages
- famous crafters
- regional economies
- player-created supply chains
- economic crises
- social hubs
- specialist dependencies.

## Q. Content Expansion

Define how future content can be added without rewriting core systems.

---

# 12. Design Method

Do not attempt to invent everything arbitrarily.

For every major system:

1. Identify the underlying design goal.
2. Identify relevant SWG inspiration.
3. Determine what should be replicated.
4. Determine what should be improved or expanded.
5. Identify interactions with other systems.
6. Identify possible exploits or failure modes.
7. Define the proposed rules.
8. Assign a status.
9. Identify unresolved questions.
10. Explain the consequences of the design.

Where an SWG mechanic is known, distinguish:

**historical SWG behaviour**

from

**TCIndustries adaptation**

from

**new TCIndustries invention**.

---

# 13. Anti-Drift Rules

The following are critical.

Do NOT gradually turn TCIndustries into:

- Factorio
- Satisfactory
- Capitalism-style business simulation
- EVE-style spreadsheet gameplay
- a conventional theme-park MMORPG
- a quest-driven RPG
- a conventional class-based MMO
- a single-player crafting game.

Industrial depth is welcome.

Economic simulation is welcome.

Automation is welcome.

But these must serve the primary fantasy of:

> **living as a citizen in a player-driven virtual society.**

---

# 14. LLM Implementation Readiness

The GDD must ultimately be written so that another LLM can use it to implement the game.

Therefore avoid vague statements such as:

> "Crafting should feel deep."

Instead eventually define:

- inputs
- outputs
- state
- rules
- calculations
- constraints
- edge cases
- progression
- dependencies
- failure conditions
- player-visible outcomes.

Where exact values are not yet known, mark them:

**TBD**

rather than inventing false precision.

---

# 15. Testing Philosophy

The eventual implementation should be testable against the GDD.

Identify:

- invariants
- balance requirements
- economic constraints
- simulation requirements
- expected emergent behaviours
- exploit resistance
- edge cases
- acceptance criteria.

Where appropriate, define future simulation tests.

---

# 16. Output Structure

Produce the Master GDD using a hierarchical structure such as:

1. Executive Summary
2. Vision
3. Design Pillars
4. Player Experience
5. World Philosophy
6. World Simulation
7. Character and Skill System
8. Resource System
9. Crafting System
10. Manufacturing
11. Economy
12. Retail
13. Professions
14. Social Systems
15. Services
16. Combat
17. Items
18. Buildings
19. Cities
20. Progression
21. Emergent Gameplay
22. Content Expansion
23. Anti-Goals
24. Balance Philosophy
25. Exploit/Fraud Considerations
26. Testing Requirements
27. Status Register
28. Open Questions
29. Future Systems
30. Glossary

Add sections where necessary.

---

# 17. Status Register

At the end maintain a master table containing:

| ID | System/Decision | Status | Rationale | Dependencies |
|---|---|---|---|---|

Use unique IDs so later conversations can refer to individual decisions.

---

# 18. Open Questions

Maintain a separate list of unresolved questions.

Do not bury important unresolved decisions inside prose.

---

# 19. Assumptions

Maintain an explicit assumption register.

Every substantive assumption introduced by the designer must be recorded.

---

# 20. First Deliverable

Do NOT pretend the game is fully designed after one pass.

Produce:

1. A coherent initial Master GDD.
2. A status register.
3. An assumption register.
4. An open-question register.
5. A list of systems requiring deeper design.
6. A recommended design sequence for subsequent iterations.

The objective is to establish the canonical architecture and philosophy first, then progressively lock the actual mechanics.

---

# 21. Most Important Principle

The most important design test for any proposed system is:

> **Does this make the virtual world more alive and create meaningful opportunities for players to develop identity, specialization, relationships, reputation, ownership, and economic/social impact?**

If a mechanic makes the game more efficient but makes other players less relevant, question it.

If automation eliminates meaningful specialization, question it.

If progression forces everyone down the same path, question it.

If the economy becomes merely a spreadsheet, question it.

If players can become economically powerful without interacting with anyone else, question it.

The goal is not maximum simulation complexity.

The goal is:

> **maximum meaningful player agency within a living, interconnected virtual society.**

Now produce the Master GDD according to this specification.