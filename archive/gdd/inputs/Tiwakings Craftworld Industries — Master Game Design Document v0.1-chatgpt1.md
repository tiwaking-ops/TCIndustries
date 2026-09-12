# Tiwakings Craftworld Industries
## Master Game Design Document

**Short Title:** TCIndustries  
**Document:** Master Game Design Document  
**Version:** 0.1 — Foundation  
**Purpose:** Canonical game/rules specification  
**Technical implementation:** Deferred to separate Technical Design Document
**Author:** ChatGPT (per filename, unverified)

---

# 1. Executive Summary

TCIndustries is an expandable virtual-world simulation inspired by the design philosophy and systemic depth of the original *Star Wars Galaxies*.

It is not intended to reproduce Star Wars.

Its purpose is to reproduce and expand upon the underlying experience of:

> **living in a persistent world as an ordinary but consequential citizen.**

The player is not necessarily a hero.

They may become a farmer, miner, prospector, artisan, doctor, entertainer, merchant, architect, manufacturer, combatant, researcher, service provider, shopkeeper, or another identity created by the game's systems.

The game is fundamentally a **player-driven sandbox**.

Its central systems should cause players to depend upon one another for:

- resources
- products
- services
- knowledge
- equipment
- medical assistance
- entertainment
- commerce
- construction
- specialization
- social organization.

The game should create a world in which players can develop genuine:

- ownership
- identity
- reputation
- specialization
- relationships
- businesses
- communities
- economic influence.

The ultimate measure of success is not how many systems exist.

It is whether those systems combine to produce a world that feels **alive**.

---

# 2. Design Status

This document uses five design statuses.

| Status | Meaning |
|---|---|
| LOCKED | Authoritative project decision |
| PROPOSED | Recommended but requires eventual approval |
| TBD | Genuine unresolved design |
| PROTOTYPE | Exists in experimental prototype form |
| DEFERRED | Known but intentionally postponed |

The absence of a status on ordinary explanatory text does not override the status of individual mechanics.

---

# 3. Core Vision

## V-001 — Living World

**Status: LOCKED**

TCIndustries is a virtual world rather than simply a game containing missions.

The world should appear to continue existing independently of any particular player's story.

Players participate in the world rather than being the centre of it.

---

## V-002 — Citizen Rather Than Chosen Hero

**Status: LOCKED**

The player should be able to live an ordinary life and still have a meaningful experience.

The game should not require a heroic narrative to provide purpose.

A player should be able to say:

> "This is who I am in the world."

rather than:

> "This is the quest line I completed."

---

## V-003 — Player-Driven Economy

**Status: LOCKED**

Players should be major participants in the production and distribution of economic goods and services.

The economy should generate meaningful opportunities for:

- production
- trade
- specialization
- commerce
- scarcity
- reputation
- entrepreneurship.

---

## V-004 — Interdependence

**Status: LOCKED**

Important professions should depend upon one another.

A complete economy should therefore contain useful roles that cannot all be replaced by a single generalist player.

---

## V-005 — Emergence

**Status: LOCKED**

The game should create conditions from which interesting events emerge rather than attempting to script every important player experience.

---

# 4. Primary Player Experiences

The following experiences define the intended emotional targets.

## 4.1 The Gold Rush

**Status: LOCKED**

A player discovers an exceptional temporary resource.

The discovery may have:

- unusually high quality
- unusual attributes
- exceptional rarity
- exceptional economic value
- temporary availability.

The player should feel urgency.

The discovery should potentially cause:

- rapid harvesting
- market speculation
- production changes
- competition
- price changes
- product innovation.

A resource discovery should therefore be an **event**, not merely an inventory update.

---

## 4.2 Becoming Known

**Status: LOCKED**

Players should be capable of becoming recognized specialists.

Examples:

> "That player makes the best armour."

> "That shop always has good medical supplies."

> "That prospector finds exceptional resources."

Reputation should emerge from player behaviour and economic history rather than merely from an abstract reputation score.

---

## 4.3 Genuine Dependency

**Status: LOCKED**

Players should be able to become economically important to other players.

A specialist should potentially experience:

> "People need what I do."

This is a central source of player identity.

---

## 4.4 Building a Community

**Status: LOCKED**

Players should be able to establish persistent social and commercial communities.

A player's shop should potentially become:

- a business
- a meeting place
- a known brand
- part of a commercial district
- part of a city.

---

# 5. Anti-Goals

The following are design warnings.

## AG-001 — Do Not Become a Spreadsheet Simulator

**Status: LOCKED**

Economic depth is desirable.

Economic abstraction that removes the human meaning of the economy is not.

---

## AG-002 — Do Not Become a Single-Player Industrial Tycoon

**Status: LOCKED**

The player should not simply acquire unlimited factories and automate every stage personally.

Industrial scale should encounter meaningful constraints.

---

## AG-003 — Do Not Make Everyone Self-Sufficient

**Status: LOCKED**

Complete self-sufficiency would undermine interdependence.

---

## AG-004 — Do Not Force Class Identity

**Status: LOCKED**

The system should not permanently label players as one class.

---

## AG-005 — Do Not Make Content Completely Script-Dependent

**Status: LOCKED**

The world should generate meaningful experiences through systems.

---

# 6. Player Model

## 6.1 Player Identity

**Status: LOCKED**

The player's identity should emerge from:

- skills
- activities
- possessions
- businesses
- reputation
- relationships
- location
- organizations
- products
- services.

---

## 6.2 Skill-Based Development

**Status: LOCKED**

The progression model should support broad skill-based specialization rather than rigid classes.

A player should be able to:

- learn skills
- improve skills
- specialize
- combine disciplines
- abandon skills
- pursue another specialization.

The exact progression mathematics is:

**Status: TBD**

unless subsequently established by a canonical system specification.

---

# 7. Resource System

## 7.1 Resource Philosophy

**Status: LOCKED**

Resources are one of the central foundations of the economy.

A resource should not merely be:

> "Iron: quantity 1000."

Resources should potentially have a meaningful identity consisting of:

- type
- location
- quality
- attributes
- rarity
- availability
- temporal characteristics
- economic desirability.

---

## 7.2 Resource Spawns

**Status: PROPOSED**

Resource availability should change over time and space.

Different areas may contain different resources or resource characteristics.

Spawns may appear, change, deteriorate, or disappear.

The precise algorithm is **TBD**.

---

## 7.3 Resource Quality

**Status: LOCKED**

Resource quality must have meaningful gameplay consequences.

A high-quality resource should potentially produce superior crafted goods.

Quality should not simply be cosmetic.

---

## 7.4 Resource Attributes

**Status: PROPOSED**

Resources should possess multiple attributes.

Different recipes should value different attributes.

This creates opportunities for:

- resource specialization
- experimentation
- market differentiation
- product optimization.

The exact attribute model is **TBD**.

---

# 8. Crafting

## 8.1 Crafting as a Profession

**Status: LOCKED**

Crafting should be a meaningful profession rather than merely an item-production menu.

---

## 8.2 Schematics

**Status: PROPOSED**

Crafting should use schematics or equivalent recipe definitions.

A schematic should define the rules required to produce an item.

Exact schematic mechanics remain **TBD**.

---

## 8.3 Resource Choice

**Status: LOCKED**

The identity and quality of input resources should influence crafted output.

Therefore:

> Two players making the same nominal item should be capable of producing materially different products.

---

## 8.4 Experimentation

**Status: PROPOSED**

Crafting should provide opportunities for experimentation.

Experimentation should allow skilled players to make meaningful trade-offs rather than merely rolling for a random quality bonus.

The exact mathematical system is **TBD**.

---

## 8.5 Signature Products

**Status: LOCKED**

Crafting should support the emergence of exceptional products associated with particular players.

Product reputation should potentially become part of the economic identity of the crafter.

---

# 9. Manufacturing

## 9.1 Purpose

**Status: LOCKED**

Manufacturing exists to support larger-scale production while preserving the importance of individual players.

---

## 9.2 Automation

**Status: LOCKED**

Automation is permitted.

Automation must not eliminate the social economy.

---

## 9.3 Personal Scale

**Status: PROPOSED**

Individual players should face meaningful limits on:

- structures
- factories
- maintenance
- capital
- logistics
- labour
- time
- access to resources.

The precise limits are **TBD**.

---

## 9.4 Social Scale

**Status: LOCKED**

Large economic organizations should generally require cooperation among multiple players.

The game should favour:

> many specialized participants

over:

> one player controlling everything.

---

# 10. Economy

## 10.1 Player-Driven Economy

**Status: LOCKED**

Players should be significant economic actors.

---

## 10.2 Supply and Demand

**Status: PROPOSED**

Prices should respond meaningfully to:

- supply
- demand
- scarcity
- quality
- competition
- location
- transportation
- player behaviour.

Exact market mechanics are **TBD**.

---

## 10.3 Scarcity

**Status: LOCKED**

Scarcity must matter.

If an exceptional resource is genuinely scarce, players should have economic reasons to compete for it.

---

## 10.4 Item Degradation

**Status: LOCKED**

Items should not necessarily remain permanently useful.

Degradation creates:

- replacement demand
- repair professions
- resource consumption
- continuing markets.

Exact degradation rules are **TBD**.

---

# 11. Retail

## 11.1 Player Shops

**Status: LOCKED**

Players should be able to establish commercial identities.

A shop should be more than an inventory container.

It should potentially become:

- a destination
- a brand
- a social hub
- a source of reputation.

---

## 11.2 Location

**Status: PROPOSED**

Commercial location should influence commerce.

A shop in a major city should potentially experience different customer behaviour from an isolated rural shop.

The exact model is **TBD**.

---

# 12. Professions

## 12.1 Flexible Profession Model

**Status: LOCKED**

Professions should emerge from combinations of skills rather than rigid permanent classes.

---

## 12.2 Profession Categories

**Status: PROPOSED**

The initial conceptual profession space includes:

### Resource

- prospector
- miner
- harvester
- farmer
- specialist resource gatherer

### Craft

- artisan
- weaponsmith
- armourer
- architect
- specialist manufacturer

### Service

- doctor
- entertainer
- transporter
- repair specialist
- researcher

### Commerce

- merchant
- trader
- shopkeeper
- broker

### Combat

- combat specialists

These are categories, not necessarily final profession names.

---

# 13. Interdependence

**Status: LOCKED**

Interdependence is a systemic requirement.

The game should contain dependency chains such as:

**Resource Specialist**
→ raw material  
→ **Crafter**
→ equipment  
→ **Combatant**

and:

**Resource Specialist**
→ resource  
→ **Manufacturer**
→ component  
→ **Crafter**
→ finished product  
→ **Retailer**
→ customer.

Likewise:

**Combatant**
→ demand  
→ **Doctor / Entertainer / Crafter**
→ services/equipment  
→ Combatant.

No single chain should necessarily dominate the entire game.

---

# 14. Social Organizations

## 14.1 Player Associations

**Status: PROPOSED**

Players should be able to form persistent organizations.

Organizations may coordinate:

- production
- commerce
- combat
- resource gathering
- cities
- transportation
- social activity.

The exact organizational mechanics are **TBD**.

---

# 15. Cities

**Status: PROPOSED**

Player communities should eventually be able to create meaningful settlements.

A city should potentially emerge from:

- player structures
- population
- services
- commerce
- organizations
- infrastructure.

Cities should become social and economic ecosystems rather than merely cosmetic housing zones.

---

# 16. Services

**Status: LOCKED**

Not all valuable player output should be an item.

Services should constitute a meaningful portion of the economy.

Potential services include:

- medical treatment
- buffs
- entertainment
- repair
- transportation
- construction
- research
- specialist production.

---

# 17. Items

## 17.1 Item Identity

**Status: PROPOSED**

Items should retain meaningful information about their origin and characteristics.

Potential identity includes:

- creator
- resources used
- quality
- modifications
- history
- ownership.

The exact provenance model is **TBD**.

---

## 17.2 Durability

**Status: LOCKED**

Durability/degradation should support continuing demand.

---

# 18. Combat

**Status: PROPOSED**

Combat exists as one important part of the world, not as the mandatory centre of the game.

Combat should create demand for:

- weapons
- armour
- ammunition/consumables where appropriate
- medical services
- buffs
- repair
- specialized equipment.

Combat therefore participates in the broader economy.

Detailed combat mechanics are **DEFERRED** until the broader world/economic architecture is sufficiently established.

---

# 19. Progression

**Status: LOCKED**

Progression should primarily support:

- capability
- specialization
- mastery
- identity.

Progression should not exist solely to create an arbitrary treadmill.

The exact progression system is **TBD**.

---

# 20. Automation and Scale

This is one of the most important TCIndustries design boundaries.

The game should support sophisticated production.

However:

> **Production scale must not automatically equal player power.**

A player controlling 100 automated production units should not simply invalidate 100 other players.

Large-scale production should encounter meaningful constraints such as:

- capital
- maintenance
- structures
- resource access
- logistics
- labour
- skill
- market demand
- social organization.

Exact implementation is **TBD**.

---

# 21. Emergent Gameplay

The game should deliberately support emergent situations.

### Example: Exceptional Resource

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

# 22. Reputation

**Status: PROPOSED**

Reputation should emerge from repeated player interaction.

Possible inputs include:

- product quality
- reliability
- service quality
- transaction history
- organization membership
- player recommendations.

A reputation system should avoid reducing reputation entirely to a single meaningless number.

The detailed model is **TBD**.

---

# 23. Ownership

**Status: LOCKED**

Ownership is central to the emotional experience.

Players should be able to feel:

- "this is my shop"
- "I made this product"
- "I discovered this resource"
- "this is my business"
- "this is our city."

Ownership should therefore be represented throughout the game.

---

# 24. World Simulation

The world should simulate enough background activity to feel alive.

Potential systems include:

- resource changes
- markets
- NPC activity
- transportation
- population
- environmental changes
- production
- consumption.

The exact simulation depth is **TBD**.

The design principle is:

> Simulate systems that produce meaningful consequences; avoid simulation for its own sake.

---

# 25. Content Architecture

**Status: LOCKED**

The game must be expandable.

New content should primarily be represented as data and rules operating on generic systems.

For example, adding a new resource should not require redesigning the resource engine.

Adding a new product should not require redesigning crafting.

Adding a new profession should not require rewriting the player model.

---

# 26. Prototype Relationship

The Seed-2.1 TypeScript prototype is:

**Status: PROTOTYPE**

It is evidence and experimentation, not authority.

The development hierarchy is:

**GDD**
→ canonical game rules

**Prototype**
→ experimental implementation

**Testing**
→ validation

**Golden Master**
→ validated behaviour

**Future UE5/C++**
→ production implementation.

---

# 27. Major Systems Still Requiring Design

The following should receive dedicated design specifications in future iterations.

### High Priority

1. Resource spawning
2. Resource quality/attribute mathematics
3. Skill system
4. Crafting mathematics
5. Experimentation
6. Item quality
7. Item degradation
8. Economy
9. Pricing
10. player shops
11. reputation
12. profession architecture
13. social organizations
14. building/ownership limits
15. manufacturing scale

### Medium Priority

16. transportation
17. cities
18. services
19. NPC society
20. world simulation
21. combat
22. housing
23. organizations
24. player governance

### Deferred

25. detailed UE5 architecture
26. networking
27. database architecture
28. rendering
29. UI architecture
30. deployment infrastructure.

---

# 28. Master Status Register

| ID | Decision | Status |
|---|---|---|
| V-001 | Living virtual world | LOCKED |
| V-002 | Citizen rather than chosen hero | LOCKED |
| V-003 | Player-driven economy | LOCKED |
| V-004 | Interdependence | LOCKED |
| V-005 | Emergent gameplay | LOCKED |
| AG-001 | Avoid spreadsheet simulation | LOCKED |
| AG-002 | Avoid single-player industrial empire | LOCKED |
| AG-003 | Avoid universal self-sufficiency | LOCKED |
| AG-004 | Avoid rigid classes | LOCKED |
| AG-005 | Avoid scripted-only content | LOCKED |
| R-001 | Dynamic resources | PROPOSED |
| R-002 | Resource quality | LOCKED |
| R-003 | Resource attributes | PROPOSED |
| C-001 | Resource-dependent crafting | LOCKED |
| C-002 | Schematics | PROPOSED |
| C-003 | Experimentation | PROPOSED |
| M-001 | Automation | LOCKED |
| M-002 | Individual production limits | PROPOSED |
| E-001 | Player-driven economy | LOCKED |
| E-002 | Supply/demand pricing | PROPOSED |
| E-003 | Item degradation | LOCKED |
| P-001 | Flexible skills/professions | LOCKED |
| S-001 | Player organizations | PROPOSED |
| S-002 | Player-created cities | PROPOSED |
| I-001 | Item provenance | PROPOSED |
| X-001 | Combat | PROPOSED |
| X-002 | Detailed combat rules | DEFERRED |
| T-001 | TypeScript prototype | PROTOTYPE |
| T-002 | UE5 technical implementation | DEFERRED |

---

# 29. Assumption Register

The following are assumptions made during this first pass and must not be treated as locked decisions.

| ID | Assumption |
|---|---|
| A-001 | The game will ultimately be persistent/multiplayer rather than permanently single-player. |
| A-002 | The world will contain player-owned structures. |
| A-003 | Resource spawning will be dynamic rather than completely static. |
| A-004 | Crafting will use a schematic/recipe abstraction. |
| A-005 | Item provenance may become an important part of reputation. |
| A-006 | Player organizations will eventually exist. |
| A-007 | Player-created cities are desired but their exact mechanics are unresolved. |
| A-008 | Manufacturing will exist but be constrained relative to social scale. |

---

# 30. Open Questions

The most important unresolved questions are:

1. What is the exact resource attribute model?
2. How are resource spawns generated?
3. How long do resource spawns last?
4. How is resource quality calculated?
5. What exactly determines crafted item quality?
6. How does experimentation work?
7. How are skills learned and improved?
8. How freely can players change specialization?
9. How are professions represented mechanically?
10. What limits individual industrial automation?
11. How does the market establish prices?
12. How are NPCs incorporated into the economy?
13. How is reputation calculated and represented?
14. How do player shops attract customers?
15. How do player cities form?
16. How does transportation affect economics?
17. What role does combat play in the overall economy?
18. What degree of persistence is required?
19. What systems are simulated continuously?
20. What systems are intentionally abstracted?

---

# 31. Recommended Design Sequence

The next design work should NOT attempt to design everything simultaneously.

Recommended order:

### Phase 1 — Resource Foundation

Define:

- resource taxonomy
- attributes
- quality
- spawning
- scarcity
- extraction
- regional variation
- temporal variation.

### Phase 2 — Crafting

Define:

- schematics
- inputs
- experimentation
- skill
- quality
- output attributes
- exceptional products.

### Phase 3 — Items

Define:

- item identity
- provenance
- durability
- degradation
- modification
- repair.

### Phase 4 — Skills and Professions

Define:

- skills
- learning
- progression
- specialization
- respecialization
- profession emergence.

### Phase 5 — Economy

Define:

- currency
- markets
- supply
- demand
- pricing
- trade
- scarcity
- economic sinks.

### Phase 6 — Retail and Reputation

Define:

- vendors
- shops
- brands
- customer behaviour
- reputation
- commercial locations.

### Phase 7 — Manufacturing

Define:

- factories
- automation
- production
- limits
- maintenance
- logistics
- social scale.

### Phase 8 — Social World

Define:

- organizations
- guilds
- communities
- cities
- ownership
- governance.

### Phase 9 — Services and Combat

Integrate:

- medical
- entertainment
- buffs
- combat
- repair
- transportation.

### Phase 10 — World Simulation

Finally integrate the systems into a coherent simulated world.

---

# 32. The Central Design Test

Every future system should be tested against this question:

> **Does this system increase meaningful player agency within a living, interconnected virtual society?**

A system should be questioned if it:

- makes other players unnecessary;
- makes specialization meaningless;
- removes scarcity;
- removes identity;
- eliminates meaningful choice;
- turns the economy into a spreadsheet;
- allows unlimited individual automation;
- forces players into predetermined classes;
- or makes the world feel like a collection of menus rather than a place.

The ultimate goal is not maximum complexity.

It is:

> **a world in which players can genuinely say "I live here, this is what I do, these are my people, this is my business, and what I do matters."**

# 33. Current Design Conclusion

TCIndustries is therefore best understood as:

> **An expandable, systemic, player-driven virtual-world simulation inspired by the original Star Wars Galaxies, in which resource discovery, crafting, specialization, commerce, services, reputation, ownership and social interdependence combine to create emergent player lives.**

The TypeScript prototype is the laboratory.

The GDD is the authority.

The eventual UE5/C++ implementation is the production realization.

The game succeeds when the player stops thinking:

> "What does the game want me to do?"

and starts thinking:

> **"What do I want to become in this world?"**