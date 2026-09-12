# Star Wars Galaxies: Pre-CU — Sandbox Recreation
## Comprehensive Game Design Document for Autonomous AI Implementation

**Document Version:** 1.0
**Document Type:** Authoritative Technical & Systems Design Specification
**Status:** Complete — Ready for AI Agent Implementation
**Classification:** Fan-Made, Non-Commercial, Educational Preservation Project

> *Star Wars, Star Wars Galaxies, and all associated planet, species, creature, and vehicle names are trademarks of Lucasfilm Ltd. This document is an original systems-design work describing game mechanics, data structures, and software architecture. It reproduces no proprietary source code, art, audio, or text from the original game. It is written in the tradition of the long-running SWG emulation and preservation community, for hobbyist, educational, and non-commercial use. Game mechanics and systems are not subject to copyright protection; only specific creative expression is, and none is reproduced here.*

---

## How to Use This Document

This GDD is written to be consumed directly by an autonomous AI coding agent as a build specification, and secondarily by human developers validating that output. Every system section defines player experience goals, mechanics, rules, data entities, system interactions, edge cases, and implementation priorities, in that order where applicable, so that architecture, schema, and gameplay logic can be derived without additional clarification. See **Section 1.4 (Document Conventions)** for the tagging system used throughout, **Section 30** for the recommended build sequence, and **Section 31** for what must exist in a first playable build.

---

## Table of Contents

1. Executive Summary
2. Design Pillars
3. Historical Design Goals
4. Core Gameplay Loop
5. World Architecture
6. Character Architecture
7. Skill-Based Progression System
8. Profession System
9. Combat System
10. Crafting System
11. Resource System
12. Economy System
13. Housing System
14. City System
15. Faction System
16. Mission System
17. Creature System
18. Social Systems
19. Grouping & Community Systems
20. Travel & Transportation
21. Vendor System
22. Bazaar & Market System
23. Entertainer System
24. Medic System
25. Ranger System
26. Surveying & Harvesting
27. Player Lifecycle
28. Data Model Specifications
29. Server Architecture Requirements
30. AI Agent Implementation Roadmap
31. MVP Scope
32. Future Expansion Hooks
33. Game Balance Targets
34. Success Metrics

---

## 1. Executive Summary

### 1.1 Project Vision

This document specifies a ground-up, systems-faithful recreation of **Star Wars Galaxies as it existed prior to the Combat Upgrade (the "Pre-CU" era, launch through late 2004)**. Pre-CU SWG is remembered by its community as one of the deepest sandbox MMORPGs ever shipped: a game with no classes, no levels in the conventional sense, no linear quest content, and an economy in which nearly every usable item passed through player hands before reaching another player. Its defining quality was not any single system but the *density of interdependence* between systems — combat needed crafting, crafting needed harvesting, everyone needed entertainers and medics, and cities needed all of them to survive.

The purpose of this project is to rebuild that specific design identity, not to reinterpret it through the lens of modern MMO conventions. Where later SWG revisions (the Combat Upgrade and New Game Enhancements) simplified the profession system into hero classes and introduced quest-driven leveling, this document deliberately preserves the *original* horizontal, sandbox, interdependence-first design.

This document is written so that an autonomous AI development agent — with no additional design meetings, no access to the original game, and no ability to ask clarifying questions mid-build — can derive a complete, internally consistent software architecture and gameplay implementation directly from the text. Where the historical record is incomplete, contradictory, or simply lost to time, this document makes an explicit, flagged design ruling rather than leaving a gap.

### 1.2 Scope Definition

**In Scope:**

- Full ground-based gameplay across the core planetary set (Section 5)
- Skill-point-based horizontal progression across all Basic and Elite professions (Sections 7–8)
- Deterministic, HAM-pool-driven combat for PvE and factional PvP (Section 9)
- Full player crafting, resource, and experimentation systems (Sections 10–11)
- A 100% player-driven economy: bazaar, vendors, harvesting, factories (Sections 12, 21–22)
- Player housing, harvester/factory placement, and player-founded cities (Sections 13–14)
- The Rebel/Imperial Galactic Civil War faction system, ground-based only (Section 15)
- Procedural mission terminals (Section 16)
- Dynamic creature spawning, lairs, and taming (Section 17)
- Entertainer and Medic social-support gameplay loops (Sections 23–24)
- Scout/Ranger exploration, tracking, and harvesting gameplay (Sections 25–26)

**Out of Scope:**

- Space gameplay (Jump to Lightspeed expansion)
- Jedi unlock via Force-sensitive progression
- Voice chat or modern social features
- Quest-driven narrative content
- Instanced dungeons or raids
- Achievement systems or cosmetic unlocks
- Microtransactions or cash shop

Out-of-scope systems are not merely deferred features — several (quest chains, instancing, achievement/unlock economies) are in direct philosophical tension with the sandbox design pillars in Section 2 and must not be reintroduced through the back door during implementation (e.g., dressing a mission chain up as a "quest" is out of scope regardless of naming). See Section 32 for how the architecture should still leave room for Jedi and space as later, clearly-separated expansions.

### 1.3 Success Criteria

The implementation succeeds if:

- **Economic Interdependence Test:** A combat player requires services from crafters, entertainers, and medics during normal gameplay.
- **Horizontal Progression Test:** A 6-month veteran and 2-week newcomer can participate meaningfully in the same content.
- **Player-Driven Economy Test:** 95% of items in circulation are player-crafted.
- **Sandbox Validation Test:** Players create emergent gameplay not explicitly designed (player events, social structures, economic schemes).
- **Historical Accuracy Test:** Core gameplay loops match documented Pre-CU behavior.

Section 34 (Success Metrics) expands these five tests into measurable, trackable KPIs for ongoing use during development and after launch.

### 1.4 Document Conventions

Every requirement, mechanic, and design ruling in this document is tagged using the following convention. An AI implementation agent should treat these tags as machine-readable metadata, not decoration.

| Tag | Meaning |
|---|---|
| **[MVP]** | Required for the minimum viable playable build. See Section 31. |
| **[EXPANSION]** | Deferred past MVP, but the data model and architecture must not preclude it. See Section 32. |
| **[HISTORICAL]** | Documented or broadly corroborated Pre-CU behavior, recreated as-is. |
| **[ASSUMPTION]** | Historical documentation is incomplete, contradictory, or lost. A specific, authoritative design ruling is made here for this project. Treat as binding, not as uncertain guidance to be second-guessed during implementation. |
| **[IMPLEMENTATION RECOMMENDATION]** | A technical suggestion (algorithm, storage pattern, tech choice) that is *not* part of the historical game design. May be substituted with an equivalent approach at the implementing agent's discretion without violating this spec. |

Sections 2–12 below establish the foundational pillars, world, character, progression, profession, combat, crafting, and resource systems. Sections 13–34 build outward from that foundation into housing, civic, social, economic-infrastructure, and technical-architecture systems, and close with the implementation roadmap and success framework.

---

## 2. Design Pillars

These pillars are inviolable. When implementation decisions conflict with these principles, preserve the pillar.

### 2.1 Pillar 1: Player Interdependence Above Self-Sufficiency

**Principle:** No player can be fully self-sufficient. Meaningful progression requires cooperation with players in different professions.

**Manifestations:**

- Combat players require crafted weapons/armor from Weaponsmiths/Armorsmiths
- Crafters require resources from Scouts/Miners and components from other crafters
- All players accumulate Battle Fatigue requiring Entertainer healing
- All players accumulate Wounds requiring Medic healing
- Buffs from Doctors/Musicians are essential for difficult content
- Resource harvesters require maintenance from Architects
- Vehicle/droid crafting requires interdisciplinary component chains

**Implementation Requirements:**

- Systems must not allow bypassing interdependence through NPC services
- Loot drops cannot replace crafted equipment as best-in-slot
- Timers and decay must ensure repeated need for services
- Solo gameplay must accumulate penalties requiring social interaction

### 2.2 Pillar 2: Horizontal Progression & Accessible Content

**Principle:** Power curves are shallow. Content accessibility is wide. Veterans gain versatility and specialization, not exponential power.

**Manifestations:**

- Starting weapon deals 50–150 damage; high-end weapon deals 200–600 damage (not 10,000+)
- Veteran combat player with 6 months progression has ~2–3x effectiveness, not 50x
- Newbie Brawler can group with Master Pikeman effectively
- Most planets are accessible at any skill level (excluding small high-risk zones)
- PvP scaling prevents total dominance by veterans
- Master-level professions unlock capability diversity (weapon types, buff types) not just bigger numbers

**Implementation Requirements:**

- Damage/health scaling must remain within 3–5x range across full progression
- Content zones must not hard-gate by level
- Skill point allocation (250 cap) must force meaningful trade-offs
- Equipment decay and maintenance must create ongoing demand regardless of progression

### 2.3 Pillar 3: Player-Driven Economy as Primary Content

**Principle:** Economic gameplay is first-class content, not a support system. Crafting is a viable primary playstyle, not a side activity.

**Manifestations:**

- Resource quality varies from 1–1000, with distributions creating rarity
- Experimentation during crafting allows customization (damage vs. speed vs. accuracy)
- Resource spawn locations rotate every 7–14 days across entire planets
- Harvesters require significant capital investment and maintenance
- Factory runs and mass production enable merchant gameplay
- Bazaar serves as universal marketplace with search/filter functionality
- Vendor NPCs at player houses create distributed retail locations

**Implementation Requirements:**

- Loot economy must not compete with crafted economy (loot = resources/components only)
- Resource spawn algorithm must create scarcity and regional variation
- Crafting interface must allow meaningful experimentation choices
- Economic data (prices, resource stats) must be queryable by players
- Storage and logistics must be non-trivial (inventory weight, volume limits)

### 2.4 Pillar 4: Sandbox Simulation Over Scripted Content

**Principle:** The world is a simulation that responds to player actions, not a sequence of authored story beats.

**Manifestations:**

- Creatures spawn from lairs; destroy the lair, stop the spawns
- Resources spawn geologically based on planetary composition
- Player cities form where players choose to build, not in designated zones
- Faction bases can be placed and destroyed dynamically
- Mission terminals generate procedural objectives, not authored quests
- Entertainer performances are player-timed social events, not cutscenes
- Harvesters extract real resources from spawn pools

**Implementation Requirements:**

- World state must persist and respond to player actions
- Procedural generation systems must create repeatable variety
- Authored content must be minimal (basic profession trainers only)
- Player-created structures must integrate into world simulation (city services, factories)
- Territorial control must emerge from player actions, not scripted campaigns

### 2.5 Pillar 5: Meaningful Social Structures

**Principle:** The game encourages formation of persistent social groups with mechanical benefits.

**Manifestations:**

- Player cities require cooperation and governance (mayors, militias, city maintenance)
- Guilds (Player Associations) provide chat, shared structures, and factional alignment
- Entertainer/Medic gameplay creates social hubs (cantinas, medical centers)
- Group combat provides XP bonuses and tactical advantages
- Crafting specialization encourages workshop cooperatives
- Faction ranks unlock when sponsored by higher-ranked players

**Implementation Requirements:**

- Systems must mechanically reward persistent group membership
- Communication tools (spatial chat, group chat, guild chat) must be robust
- Shared ownership structures (guild halls, city structures) must be supported
- Reputation systems (vendor ratings, crafter reputation) must be trackable

---

## 3. Historical Design Goals

### 3.1 What SWG Pre-CU Achieved

**Primary Achievements:**

- **Living Economy:** Market prices determined by supply/demand, not vendor pricing
- **Meaningful Crafting Profession:** Crafters were not alts, but primary characters
- **Social Gameplay Loops:** Entertaining and healing were engaging activities, not time-sinks
- **Player Housing as Content:** Decorating, vendor management, and city-building were end-game activities
- **Horizontal Skill Progression:** Freedom to experiment with builds without rerolling characters

### 3.2 What SWG Pre-CU Avoided

**Deliberately Excluded Designs:**

- **Theme Park Questing:** No linear quest chains leading to max level
- **Class Lock-In:** No irreversible class choices at character creation
- **Bind-on-Equip Gear:** No soulbound equipment preventing trading
- **Instanced Content:** No private dungeons; all content in shared world
- **Exponential Power Curves:** No +1000% damage increases at max level
- **Self-Sufficient Characters:** No ability to craft, combat, heal, and buff on one character

### 3.3 Design Intent Clarifications

**Assumption Flags (Historical Behavior Where Documentation Is Limited):**

- **[ASSUMPTION] Resource Spawn Timing:** Historical sources suggest 7–14 day rotation; exact algorithm unknown. Implement 10-day baseline with variance.
- **[ASSUMPTION] Experimentation Success Rates:** Exact formulas lost; recreate based on observed ranges (critical success ~10%, critical failure ~5% with appropriate tools).
- **[ASSUMPTION] Combat XP Scaling:** Group XP bonus mechanics partially documented; implement 10% bonus per group member (max 50% at 5+ members).
- **[ASSUMPTION] Battle Fatigue Accumulation:** Rate varied by action; implement 1–3% per death, 0.1–0.5% per combat action with diminishing returns.

---

## 4. Core Gameplay Loop

### 4.1 Moment-to-Moment Gameplay (5–30 Minutes)

**For Combat-Focused Players:**

1. Accept mission from terminal (destroy lair, bounty hunt, patrol)
2. Travel to location (on foot, vehicle, or shuttleport)
3. Engage creatures/NPCs using action-based combat (HAM pool management)
4. Loot resources, credits, and mission tokens
5. Return to city/outpost for turn-in
6. Manage inventory, repair equipment wear
7. Repeat or switch activities

**For Crafter-Focused Players:**

1. Survey current resource spawns for quality materials
2. Place/manage harvesters on optimal spawn locations
3. Retrieve harvested resources from hoppers
4. Experiment with schematics to create items
5. Run factory production for mass manufacturing
6. List items on bazaar or stock personal vendor
7. Monitor sales, adjust prices, research market trends

**For Social/Support Players (Entertainer/Medic):**

1. Set up in cantina or medical center
2. Advertise services in spatial/planetary chat
3. Perform healing services (Battle Fatigue removal, Wound healing, buff application)
4. Socialize with clients, build reputation
5. Earn tips and credits for services
6. Monitor clientele needs, adjust service offerings

### 4.2 Session Gameplay Loop (1–4 Hours)

**Progression Loop:**

1. Identify skill to advance (e.g., Pistol Accuracy IV)
2. Engage in activities granting relevant XP (combat XP for weapon skills)
3. Accumulate XP toward next skill box
4. Purchase skill advancement at trainer
5. Repeat across profession tree branches

**Economic Loop:**

1. Identify market demand (check bazaar prices, talk to players)
2. Acquire resources (harvesting, trading, bazaar purchase)
3. Craft items with experimentation for quality
4. Market items (bazaar listings, vendor placement, chat advertising)
5. Reinvest profits in better tools, harvesters, or factories
6. Expand production capacity or diversify into component chains

**Social Loop:**

1. Join guild or city
2. Contribute to group goals (city maintenance, faction PvP, resource gathering)
3. Attend player events (cantina concerts, duels, auctions)
4. Build reputation through service quality or PvP performance
5. Unlock social advancement (city rank, guild officer, factional prestige)

### 4.3 Long-Term Gameplay Loop (Weeks to Months)

**Character Mastery:**

1. Master starter profession (4 profession trees, ~40–60 hours)
2. Explore hybrid builds (drop skills, learn new trees, respec)
3. Master advanced profession (Bounty Hunter, Creature Handler, etc., 80–120 hours)
4. Optimize template for specialization (PvP, PvE, crafting efficiency)
5. Maintain multiple templates by respeccing as needed

**Economic Empire Building:**

1. Establish harvester network across multiple planets
2. Build factory production lines for high-demand items
3. Recruit other crafters to supply component chains
4. Establish storefront in high-traffic player city
5. Build brand reputation for quality (high experimentation results)
6. Diversify into multiple crafting professions via alts or respec

**Community Leadership:**

1. Participate in city governance (vote for mayor, run for office)
2. Lead guild operations (organize PvP raids, manage guild hall)
3. Establish factional reputation (lead base assaults, command player militias)
4. Host community events (cantina concerts, swoop races, auctions)
5. Mentor new players through profession training or resource donations

### 4.4 System Interdependencies (The Virtuous Cycle)

```text
COMBAT PLAYER needs:
├─ Weapons/Armor → WEAPONSMITH/ARMORSMITH
│   └─ Resources → SCOUT/MINER
├─ Buffs → DOCTOR/MUSICIAN
├─ Wound Healing → COMBAT MEDIC
├─ Battle Fatigue Healing → ENTERTAINER/DANCER
└─ Missions → MISSION TERMINALS (system)

CRAFTER needs:
├─ Resources → SCOUT/MINER (self or market)
├─ Harvester Maintenance → ARCHITECT (for deed) + DROID ENGINEER (for repairs)
├─ Factory Production → ARCHITECT (for factory deed)
├─ Schematics → TRAINERS + XP GRINDING (often via combat)
└─ Market Demand → COMBAT PLAYERS + OTHER CRAFTERS

ENTERTAINER/MEDIC needs:
├─ Clients → COMBAT PLAYERS (Battle Fatigue, Wounds)
├─ Venue → ARCHITECTS (cantinas, hospitals in cities)
├─ Instruments/Stim Packs → CRAFTERS
└─ Income → TIPS + SERVICE FEES

SCOUT/MINER needs:
├─ Survey Tools → CRAFTERS (often Engineering)
├─ Harvesters → ARCHITECTS
├─ Protection in Dangerous Zones → COMBAT PLAYERS
└─ Market Buyers → CRAFTERS
```

This interdependence is **MANDATORY**. Systems must not allow bypassing these loops through NPC vendors or loot drops.

---

## 5. World Architecture

### 5.1 Planetary Design Philosophy

**Core Principle:** Each planet serves distinct ecological, economic, and social roles in the simulation. Planets are not level-gated zones, but differentiated by resource availability, creature danger, factional control, and social infrastructure.

**Planetary Roles:**

- **Starter Planets:** Low-danger zones near NPC cities (Tatooine, Corellia, Naboo)
- **Resource Planets:** High-value resource spawns, remote locations (Lok, Dathomir)
- **Factional Conflict Zones:** Contested territory with bases (Corellia, Talus)
- **Social Hubs:** High player density near shuttleports and cantinas (Coronet, Theed)
- **Wilderness Planets:** Low infrastructure, high resource value (Endor, Dathomir)

### 5.2 Planet List & Characteristics

#### 5.2.1 Tatooine
- **Terrain:** Desert, canyons, rocky outcrops
- **NPC Cities:** Mos Eisley (major hub), Mos Entha, Bestine, Anchorhead
- **Creature Difficulty:** Low to Medium (Worrt, Dewback, Tusken Raider)
- **Resources:** Ore, Radioactive, Fiberplast, low-quality organics
- **Factional Presence:** Neutral (both Imperial and Rebel bases allowed)
- **Role:** Primary starter planet, high social activity — **[MVP]**

#### 5.2.2 Corellia
- **Terrain:** Grasslands, forests, rivers
- **NPC Cities:** Coronet (largest city), Tyrena, Kor Vella, Doaba Guerfel
- **Creature Difficulty:** Low to Medium (Durni, Humbaba, Gronda)
- **Resources:** High-quality Aluminum, Copper, organics
- **Factional Presence:** Heavy Imperial presence, Rebel contested zones
- **Role:** Economic hub, balanced starter planet — **[MVP]**

#### 5.2.3 Naboo
- **Terrain:** Swamps, lakes, plains, forests
- **NPC Cities:** Theed (high-tier architecture), Keren, Moenia, Kaadara
- **Creature Difficulty:** Low to Medium (Gungan, Fambaa, Veermok)
- **Resources:** Organic resources (meat, hide, bone), water
- **Factional Presence:** Neutral, leaning Rebel-friendly
- **Role:** Aesthetic appeal, organic resource farming — **[MVP]**

#### 5.2.4 Talus
- **Terrain:** Forests, volcanic regions
- **NPC Cities:** Dearic, Nashal
- **Creature Difficulty:** Medium (Fynock, Kima, Kahmurra)
- **Resources:** Gemstones, Copper, Iron
- **Factional Presence:** Contested
- **Role:** Mid-tier resource planet — **[EXPANSION]**, see Section 31

#### 5.2.5 Rori
- **Terrain:** Swamps, wetlands (moon of Naboo)
- **NPC Cities:** Narmle, Restuss (later destroyed in quest event — omit for MVP)
- **Creature Difficulty:** Medium (Torton, Carrion Spat)
- **Resources:** Chemicals, radioactive materials
- **Factional Presence:** Neutral
- **Role:** Supplemental to Naboo — **[EXPANSION]**

#### 5.2.6 Lok
- **Terrain:** Volcanic, sulfur plains
- **NPC Cities:** None (pirate outposts only)
- **Creature Difficulty:** High (Kimogila, Canyon Krayt)
- **Resources:** Rare metals, high-value ores
- **Factional Presence:** Neutral (pirate faction NPCs)
- **Role:** High-risk, high-reward resource farming — **[EXPANSION]**

#### 5.2.7 Dantooine
- **Terrain:** Plains, grasslands
- **NPC Cities:** Pirate Outpost, Mining Outpost
- **Creature Difficulty:** Medium to High (Janta, Kunga, Thune)
- **Resources:** Wheat, organic crops, moderate metals
- **Factional Presence:** Neutral
- **Role:** Jedi progression planet (omit Jedi enclave for MVP) — **[EXPANSION]**

#### 5.2.8 Dathomir
- **Terrain:** Mountainous, rocky, hostile
- **NPC Cities:** None
- **Creature Difficulty:** Very High (Rancor, Nightsister, Singing Mountain Clan)
- **Resources:** Rare organics, radioactive, gemstones
- **Factional Presence:** Neutral (hostile NPC factions)
- **Role:** End-game resource farming, high-risk content — **[EXPANSION]**

#### 5.2.9 Yavin IV
- **Terrain:** Dense jungle, ruins
- **NPC Cities:** None
- **Creature Difficulty:** High (Tybis, Woolamander, Sith cultists)
- **Resources:** Gemstones, rare organics
- **Factional Presence:** Neutral
- **Role:** Exploration, high-difficulty creature hunting — **[EXPANSION]**

#### 5.2.10 Endor
- **Terrain:** Forest moon, redwood-like trees
- **NPC Cities:** Ewok villages (non-interactive), Research Outpost
- **Creature Difficulty:** Medium to High (Jinda, Marauder, Gorax)
- **Resources:** Wood, organic materials
- **Factional Presence:** Neutral
- **Role:** Organic resource farming, aesthetic appeal — **[EXPANSION]**

### 5.3 Zone Architecture

#### 5.3.1 NPC Cities (Static Structures)

**Purpose:** Provide starting points, trainers, and basic services.

**Components:**

- **Shuttleport:** Travel between cities and planets
- **Starport:** Planetary to planetary travel (requires shuttle ticket purchase)
- **Cantina:** Entertainer performances, social hub, mission terminals
- **Medical Center:** Medic services, clone facility (respawn point)
- **Cloning Facility:** Set as respawn location (death recovery point)
- **Trainers:** Profession skill trainers (Basic Professions only)
- **Misc Structures:** Banks (item storage), Bazaar terminals (market search)

**Design Rules:**

- NPC cities cannot be modified by players
- NPC vendors sell only basic supplies (food, water, basic tools)
- No high-quality equipment sold by NPCs
- NPC trainers only train basic profession skills (Novice to Master in starter trees)

#### 5.3.2 Player Cities (Dynamic Structures)

**Purpose:** Player-driven community hubs with civic structures and economic benefits. Full specification in Section 14.

**Formation Requirements:**

- 10+ player structures (houses, harvesters) within defined radius
- City Hall deed placed by founder
- Minimum maintenance pool funded by citizens

**Components:**

- **City Hall:** Governance, taxation, zoning controls
- **Cantina:** Player-built, supports Entertainer gameplay
- **Medical Center:** Player-built, supports Medic gameplay
- **Shuttleport:** Unlocked at City Rank 3, enables fast travel
- **Cloning Facility:** Unlocked at City Rank 4, respawn point
- **Gardens/Decorations:** Aesthetic improvements

**Design Rules:**

- Players vote for mayor in elections (weekly voting periods)
- Mayor sets tax rates, building restrictions, and militia permissions
- City maintenance costs scale with city rank and structure count
- Cities can lose rank if maintenance lapses

#### 5.3.3 Wilderness Zones

**Purpose:** Resource spawns, creature lairs, and open-world content.

**Characteristics:**

- No structure density limits (place harvesters/houses anywhere outside NPC cities)
- Creature lairs spawn dynamically based on biome
- Resource spawns rotate on timed cycles
- No invisible walls or hard boundaries (can traverse entire planet surface)

**Danger Zones:**

- **Low Danger:** Near NPC cities, starter zones (creature levels 1–20)
- **Medium Danger:** 5km+ from cities (creature levels 15–50)
- **High Danger:** Remote regions, specific biomes like Dathomir mountains (creature levels 50–90+)

### 5.4 World State Simulation

#### 5.4.1 Persistent World State

**What Persists:**

- Player structures (houses, harvesters, factories)
- Resource spawn locations and stats
- Creature lair locations (until destroyed)
- Faction base locations (until destroyed)
- Player city structures and rank

**What Resets:**

- Creature spawns (respawn from lairs on timers)
- Mission objectives (regenerate on terminal refresh)
- NPC vendor inventory (basic supplies unlimited)

#### 5.4.2 World Ticking Systems

**Resource Spawns:**

- Spawn new resource types on 10-day cycle (with variance)
- Despawn previous resource types (gradual concentration shift)
- Update spawns overnight during low-traffic hours

**Creature Lairs:**

- Spawn new lairs in valid biomes on 12–24 hour cycles
- Lairs destroyed by players do not respawn in same location
- Lair density capped per region to prevent overpopulation

**Harvesters:**

- Extract resources hourly (tick rate: 1 hour)
- Fill hoppers up to capacity (requires periodic player emptying)
- Decay maintenance pools (require weekly funding or shut down)

**City Maintenance:**

- Deduct city upkeep from treasury weekly
- Downgrade city rank if treasury insufficient
- Remove structures if city falls below minimum rank

### 5.5 Travel & Distance Design

#### 5.5.1 Travel Time Philosophy

**Principle:** Distance is meaningful. Fast travel is limited to create geographic value for cities and player placement.

**Travel Options:**

- **On Foot:** Slow (base speed 5–7 m/s), unlimited
- **Vehicles:** Medium (speeder bikes 15–25 m/s), crafted or looted
- **Mounts:** Medium (Dewbacks, Kaadu, 12–20 m/s), tamed by Bio-Engineer/Creature Handler
- **Shuttles (City to City):** Instant (5-minute wait at shuttle), costs credits
- **Starport (Planet to Planet):** Instant (loading screen), costs credits

**Design Rules:**

- No teleportation (except death recovery to clone facility)
- No instant group summons
- Vehicles require fuel (or maintenance in simplified model)
- Shuttles only connect established routes (NPC cities + Rank 3+ player cities)

#### 5.5.2 Distance as Economic Mechanic

**Implications:**

- Remote resource spawns require significant travel investment
- Player cities compete for convenient locations near resources/hunting grounds
- Harvesters in dangerous zones risk destruction for better resources
- Vendor location matters (cities near shuttles have higher traffic)

Full travel mechanics, vehicle data entities, and edge cases are specified in Section 20.

---

## 6. Character Architecture

### 6.1 Character Identity Philosophy

**Core Principle:** Characters are defined by skill allocation, not class selection. Identity is fluid, allowing respecialization without rerolling.

**Player Experience Goals:**

- Freedom to experiment with builds
- Ability to pivot professions based on evolving interests
- Persistent character identity despite profession changes
- Meaningful trade-offs (250 skill point cap enforces specialization)

### 6.2 Character Attributes (Species)

#### 6.2.1 Species List & Modifiers

**Available Species:**

- **Human:** Balanced, no penalties, no bonuses
- **Bothan:** +10 Agility, −10 Strength
- **Rodian:** +10 Agility, −10 Constitution
- **Trandoshan:** +20 Strength, −10 Quickness
- **Twi'lek:** +10 Charisma, −10 Stamina
- **Wookiee:** +50 Strength, −50 Quickness, +100 Health, cannot use certain armor
- **Zabrak:** +10 Stamina, −10 Willpower
- **Mon Calamari:** +10 Intelligence, −10 Constitution
- **Sullustan:** +10 Agility, −10 Stamina

**Design Rules:**

- Species choice is permanent
- Stat differences are minor (~5–10% variance in relevant attributes)
- No species gating for professions (Wookiees can be Entertainers, etc.)
- Wookiee armor restriction is cosmetic limitation (cannot wear certain slots)

#### 6.2.2 Base Attributes (HAM Pools)

**Health/Action/Mind (HAM) Pools:**

Each character has 9 attribute bars (3 pools × 3 bars each):

- **Health Pool:** Health, Strength, Constitution
- **Action Pool:** Action, Quickness, Stamina
- **Mind Pool:** Mind, Focus, Willpower

**Starting Values (Human baseline):**

| Attribute | Value |
|---|---|
| Health | 1000 |
| Strength | 500 |
| Constitution | 500 |
| Action | 1000 |
| Quickness | 500 |
| Stamina | 500 |
| Mind | 1000 |
| Focus | 500 |
| Willpower | 500 |

**Modifications:**

- Species apply static bonuses/penalties (see 6.2.1)
- Skill boxes provide +50 to +500 per box in relevant attributes
- Buffs from Doctors/Musicians add temporary bonuses (+500 to +3000 depending on buff power)
- Wounds reduce maximum pool size until healed by Medic
- Battle Fatigue reduces maximum pool size until healed by Entertainer

**Implementation Note:** HAM pools are the foundation of all combat and action resolution. Damage/costs always target specific bars (e.g., ranged attack damages Health, running costs Action).

### 6.3 Skill System Architecture

#### 6.3.1 Skill Point Economy

**Core Rules:**

- Every character has 250 skill points total
- Skills are organized into profession trees (see Section 8)
- Each skill box costs 1–10 skill points depending on tier
- Novice boxes cost 0 points (free entry into profession)
- Master boxes cost 10 points (capstone of profession tree)
- Dropping skills refunds skill points (instant respec allowed)

**Example Skill Box Costs:**

```text
Marksman Tree:
├─ Novice Marksman (0 points) [FREE ENTRY]
├─ Ranged Accuracy I (2 points)
├─ Ranged Accuracy II (3 points)
├─ Ranged Accuracy III (4 points)
├─ Ranged Accuracy IV (5 points)
├─ Ranged Speed I (2 points)
├─ Ranged Speed II (3 points)
├─ Ranged Speed III (4 points)
├─ Ranged Speed IV (5 points)
[...etc across 4 trees totaling ~80 points to Master Marksman]
```

#### 6.3.2 Skill Progression Mechanics

**Earning XP:**

1. Perform activities related to skill (combat for weapon skills, crafting for crafting skills)
2. Accumulate typed XP (e.g., Melee Combat XP, Scouting XP, Medical XP)
3. XP requirements scale per skill box (early boxes ~1k XP, master boxes ~100k XP)
4. Visit trainer NPC to purchase skill box (costs XP + credits)

**XP Types (Partial List):**

- Combat XP (general combat actions)
- Weapon-Specific XP (Rifle, Pistol, Melee, Unarmed)
- Crafting XP (divided by profession: Armor, Weapon, Engineering, etc.)
- Scouting XP (surveying, creature knowledge)
- Medical XP (healing wounds, applying medicine)
- Entertainer XP (performing, healing battle fatigue)
- Merchant XP (selling items via vendor/bazaar)

**Design Rules:**

- XP earned is capped per action (diminishing returns prevent low-level grinding)
- Group combat provides XP bonus (+10% per group member, max +50%)
- Crafting XP earned on successful experimentation or item creation
- Social XP (Entertainer, Medic) earned by servicing other players

#### 6.3.3 Skill Dependencies

**Prerequisite Chains:**

- Must purchase boxes in sequence within a tree (cannot skip tiers)
- Must complete Novice box before any tree boxes
- Must complete all 4 trees to unlock Master box
- Some professions require mastery of another (e.g., Bounty Hunter requires Master Marksman or Scout)

**Example:**

```text
To become Master Pistoleer:
1. Train Novice Marksman (free)
2. Train Pistol Accuracy I-IV (sequential)
3. Train Pistol Speed I-IV (sequential)
4. Train Pistol Support I-IV (sequential)
5. Train Pistol Finesse I-IV (sequential)
6. Train Master Pistoleer (requires all 4 trees complete)
```

#### 6.3.4 Skill Templates (Player Builds)

**Common Templates:**

- **Pure Combat:** Master Rifleman, Master Pistoleer, Novice Medic (250 points)
- **Crafter:** Master Armorsmith, Master Weaponsmith (160 points), Novice Merchant (10 points)
- **Hybrid:** Master Ranger, Novice Creature Handler, Novice Scout
- **Support:** Master Doctor, Master Combat Medic
- **Social:** Master Entertainer (Musician), Master Dancer, Master Image Designer

**Design Philosophy:**

- Mastering 2 elite professions fills ~200–220 skill points, leaving room for utility
- Hybrid builds sacrifice mastery for versatility
- Respeccing is encouraged (cost is time to re-earn XP, not permanence)

### 6.4 Character Customization (Appearance)

#### 6.4.1 Character Creation

**Customization Options:**

- Species selection (9 species)
- Gender (Male/Female)
- Facial features (sliders for nose, eyes, mouth, etc.)
- Hair style and color (species-dependent)
- Body height/weight (sliders)
- Starting clothing (cosmetic only)

**Implementation Note:** Character appearance is stored as mesh parameters + texture references. Image Designers can later modify most options (see Section 8.3.20).

#### 6.4.2 Ongoing Customization

- **Clothing/Armor:** Equipped items modify appearance
- **Image Designer Services:** Hair, face, body modifications post-creation
- **Wearables:** Backpacks, belts, jewelry (cosmetic or stat-modifying)

---

## 7. Skill-Based Progression System

### 7.1 Progression Philosophy

**Core Principle:** Progression is horizontal expansion of capabilities, not vertical power escalation. A Master Rifleman is not "level 80," they are specialized in rifle combat.

**Player Experience Goals:**

- Clear feedback on advancement (skill box unlocks)
- Freedom to change direction mid-progression
- Meaningful choices (skill point allocation trade-offs)
- Ongoing progression hooks (mastering new professions)

### 7.2 XP Earning Mechanics

#### 7.2.1 Combat XP

**Sources:**

- Killing creatures (scales with creature level)
- Destroying lairs (bonus XP)
- Completing combat missions (mission reward XP)
- PvP kills (factional XP + combat XP)

**Scaling:**

```text
Creature XP = Base XP × (Creature CL / Player Effective Level) × Group Bonus
```

- Diminishing returns if creature is too low level (>10 levels below player)
- Boss creatures grant 5–10x normal XP

**Group Bonus:**

- Solo: 1.0x XP
- 2 players: 1.1x XP each
- 3 players: 1.2x XP each
- 4 players: 1.3x XP each
- 5+ players: 1.5x XP each

#### 7.2.2 Crafting XP

**Sources:**

- Crafting items (XP per item crafted)
- Experimentation success (bonus XP for critical success)
- Factory runs (XP per batch, reduced per-item XP)
- Reverse engineering schematics (one-time XP)

**Scaling:**

- Item complexity determines XP (simple items 10–50 XP, complex items 500–2000 XP)
- Repeat crafting same item has diminishing returns (50% XP after 10th craft)
- Experimentation critical success grants 2x XP

#### 7.2.3 Social XP (Entertainer/Medic)

**Entertainer XP:**

- Healing Battle Fatigue (XP per tick of healing, scales with BF healed)
- Performing in groups (passive XP while performing in cantina)
- Tip bonuses (small XP bonus when tipped by audience)

**Medical XP:**

- Healing Wounds (XP per point healed)
- Applying enhancement buffs (XP per buff applied)
- Reviving incapacitated players (bonus XP)

**Scaling:**

- Higher-tier patients (more BF/Wounds) grant more XP
- Repeat healing same player has diminishing returns (cooldown period)

#### 7.2.4 Scout/Ranger XP

**Sources:**

- Surveying resources (XP per survey action)
- Tracking creatures (XP per tracking check)
- Harvesting resources (XP per sample)
- Camping (XP per camp use by group members)

#### 7.2.5 Merchant XP

**Sources:**

- Selling items via vendor (XP per transaction)
- Bazaar sales (XP per sale)
- Brokering trades (XP per player-to-player trade facilitated)

**Scaling:**

- Higher-value sales grant more XP (1 XP per 100 credits)
- Unique buyers grant more XP (repeat customers diminished)

### 7.3 Skill Unlocking Process

**Step-by-Step:**

1. Player earns XP through activities
2. XP accumulates in typed pool (e.g., Melee Combat XP: 5,340 / 10,000)
3. Player visits trainer NPC
4. Player selects skill box to purchase
5. Trainer checks:
   - Does player have enough XP?
   - Does player have enough credits?
   - Does player have prerequisite skills?
   - Does player have enough skill points available?
6. If all checks pass, skill box is granted
7. Player attributes/abilities update immediately

**Trainer Locations:**

- Basic profession trainers in all NPC cities
- Elite profession trainers in specific NPC cities (e.g., Bounty Hunter trainer only in select locations)
- Player cities can host trainers if city rank is sufficient

### 7.4 Skill Decay & Respec

#### 7.4.1 Dropping Skills

**Process:**

1. Player opens skill interface
2. Player selects skill box to drop
3. Confirmation prompt warns of dependent skills
4. Skill points refunded immediately
5. XP is NOT refunded (must re-earn if re-training)

**Design Rules:**

- Dropping a skill drops all dependent skills (e.g., dropping Ranged Accuracy II drops Accuracy III and IV)
- Dropping Novice box drops entire profession tree
- No cooldown on respec (instant reallocation)
- Equipped items requiring dropped skills become unusable (warning prompt)

#### 7.4.2 Skill Decay Absence

**Important:** SWG Pre-CU does NOT have skill decay from disuse. Once trained, skills persist until deliberately dropped.

---

## 8. Profession System

### 8.1 Profession Architecture

**Profession Categories:**

- **Basic/Starting Professions:** Available to all players, no prerequisites
- **Elite Professions:** Require mastery of one or more basic professions
- **Hybrid Professions:** Combine multiple basic professions

**Total Professions:** 32 (including hybrids and elite paths). This document enumerates 6 Basic, 23 numbered Elite (8.3.1–8.3.23), plus Tailor and Chef (crafting elites detailed fully in Section 10.9.5–10.9.6). Section 14 (City System) and Section 19 (Grouping & Community Systems) introduce two additional historically-documented elite professions — **Politician** and **Squad Leader** — that complete the roster. **[ASSUMPTION]** Exact historical profession counts varied by patch/era; this document's roster is the authoritative set for this project.

### 8.2 Basic Professions (Starting Professions)

#### 8.2.1 Artisan (Crafting Foundation)

**Description:** Entry-level crafting profession. Foundation for all advanced crafting paths.

**Skill Trees (4 trees):**

- Engineering I–IV: Survey tools, basic components
- Business I–IV: Resource efficiency, factory use
- Experimentation I–IV: Experimentation bonuses
- Complexity I–IV: Unlock complex schematics

**Benefits:**

- Unlock crafting interface
- Unlock surveying tools
- Unlock experimentation mechanics
- Access to basic resource schematics

**XP Type:** Crafting XP (General)
**Skill Point Cost:** ~40 points to Master Artisan

**Elite Paths:** Armorsmith (requires Master Artisan), Weaponsmith (requires Master Artisan), Architect (requires Master Artisan), Droid Engineer (requires Master Artisan), Tailor (requires Master Artisan), Chef (requires Master Artisan)

#### 8.2.2 Brawler (Unarmed Combat)

**Description:** Unarmed and melee combat specialist. Foundation for advanced melee professions.

**Skill Trees:**

- Unarmed Damage I–IV: Increase unarmed damage
- Unarmed Speed I–IV: Increase attack speed
- Melee Support I–IV: Defensive abilities, knockdowns
- Melee Finesse I–IV: Special attacks, posture changes

**Benefits:**

- Unlock unarmed combat abilities
- Unlock basic melee weapons (vibroblades, stun batons)
- Increased HAM pool bonuses (Health, Strength)

**XP Type:** Melee Combat XP
**Skill Point Cost:** ~40 points to Master Brawler

**Elite Paths:** Fencer (requires Master Brawler), Swordsman (requires Master Brawler), Pikeman (requires Master Brawler), Teras Kasi Artist (requires Master Brawler + other paths)

#### 8.2.3 Marksman (Ranged Combat Foundation)

**Description:** Ranged combat specialist. Foundation for pistol and rifle professions.

**Skill Trees:**

- Ranged Accuracy I–IV: Increase accuracy
- Ranged Speed I–IV: Increase attack speed
- Ranged Support I–IV: Special ammo, defensive postures
- Ranged Finesse I–IV: Critical hits, special attacks

**Benefits:**

- Unlock ranged weapons (pistols, rifles, carbines)
- Increased HAM pool bonuses (Action, Quickness)
- Unlock ranged combat stances

**XP Type:** Ranged Combat XP
**Skill Point Cost:** ~40 points to Master Marksman

**Elite Paths:** Pistoleer (requires Master Marksman), Rifleman (requires Master Marksman), Carbineer (requires Master Marksman), Commando (requires Master Marksman + other paths), Bounty Hunter (requires Master Marksman or Scout)

#### 8.2.4 Scout (Exploration & Tracking)

**Description:** Wilderness survival, tracking, and resource harvesting specialist.

**Skill Trees:**

- Exploration I–IV: Movement speed, terrain navigation
- Hunting I–IV: Creature tracking, harvesting bonuses
- Trapping I–IV: Trap placement, creature knowledge
- Survival I–IV: Camp placement, group buffs

**Benefits:**

- Unlock tracking abilities
- Unlock camp placement (group healing/buff zones)
- Harvest rare resources from creatures
- Movement speed bonuses

**XP Type:** Scouting XP
**Skill Point Cost:** ~40 points to Master Scout

**Elite Paths:** Ranger (requires Master Scout), Creature Handler (requires Master Scout), Bio-Engineer (requires Master Scout + other paths), Bounty Hunter (requires Master Marksman or Scout)

#### 8.2.5 Medic (Healing & Support Foundation)

**Description:** Wound healing and injury treatment specialist.

**Skill Trees:**

- Healing I–IV: Wound healing power
- Injury Treatment I–IV: Reviving incapacitated players
- Medicine Crafting I–IV: Craft stim packs, medical supplies
- Support I–IV: Diagnosis, disease cures

**Benefits:**

- Heal Wounds (permanent damage until healed)
- Revive incapacitated players
- Craft medical supplies (stim packs, enhancement packs)

**XP Type:** Medical XP
**Skill Point Cost:** ~40 points to Master Medic

**Elite Paths:** Doctor (requires Master Medic), Combat Medic (requires Master Medic)

#### 8.2.6 Entertainer (Social & Performance)

**Description:** Performance, image design, and Battle Fatigue healing specialist.

**Skill Trees:**

- Music I–IV: Instrument performance, BF healing
- Dance I–IV: Dance performance, BF healing
- Healing I–IV: Battle Fatigue healing rate
- Showmanship I–IV: Group performances, flourishes

**Benefits:**

- Heal Battle Fatigue (performance in cantinas)
- Perform for audiences (social gameplay)
- Unlock instrument/dance types

**XP Type:** Entertainer XP
**Skill Point Cost:** ~40 points to Master Entertainer

**Elite Paths:** Musician (requires Master Entertainer), Dancer (requires Master Entertainer), Image Designer (requires Master Entertainer)

### 8.3 Elite Professions (Advanced Specializations)

#### 8.3.1 Armorsmith
**Prerequisites:** Master Artisan
**Description:** Craft armor and protective gear. High demand in player economy.
**Skill Trees:** Personal Armor I–IV, Armor Assembly I–IV, Armor Experimentation I–IV, Armor Customization I–IV
**Benefits:** Craft armor (Composite, Bone, Chitin, Ubese, RIS); experimentation for damage reduction, encumbrance, effectiveness; access to rare armor schematics
**XP Type:** Armor Crafting XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** High demand from combat players

#### 8.3.2 Weaponsmith
**Prerequisites:** Master Artisan
**Description:** Craft melee and ranged weapons.
**Skill Trees:** Melee Weapons I–IV, Ranged Weapons I–IV, Weapon Experimentation I–IV, Weapon Assembly I–IV
**Benefits:** Craft weapons (power hammers, rifles, swords, pistols); experimentation for min/max damage, attack speed, accuracy; access to rare weapon schematics
**XP Type:** Weapon Crafting XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** Essential for combat player economy

#### 8.3.3 Architect
**Prerequisites:** Master Artisan
**Description:** Craft structures, furniture, and civic buildings.
**Skill Trees:** Structures I–IV, Furniture I–IV, Installation I–IV, City Planning I–IV
**Benefits:** Craft house deeds (small, medium, large); craft harvester deeds; craft factory deeds; craft civic structure deeds (cantinas, hospitals, shuttleports)
**XP Type:** Structure Crafting XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** High demand for houses, harvesters, and city structures

#### 8.3.4 Droid Engineer
**Prerequisites:** Master Artisan
**Description:** Craft droids for combat, crafting, and utility roles.
**Skill Trees:** Droid Assembly I–IV, Droid Customization I–IV, Advanced Droids I–IV, Droid Programming I–IV
**Benefits:** Craft droid chassis (R-series, protocol droids, combat droids); craft droid modules (combat, crafting, storage); program droid behavior (follow, defend, harvest)
**XP Type:** Droid Crafting XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** Moderate demand for combat and crafting droids

#### 8.3.5 Merchant
**Prerequisites:** Master Artisan
**Description:** Optimize vendor management, reduce vendor fees, and improve market presence.
**Skill Trees:** Vendor Management I–IV, Advertising I–IV, Wholesale I–IV, Appraisal I–IV
**Benefits:** Reduce vendor maintenance costs (up to 50% reduction); increase vendor inventory slots; enhanced bazaar listing features
**XP Type:** Merchant XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** Essential for high-volume crafters

#### 8.3.6 Bio-Engineer
**Prerequisites:** Master Scout + Novice Medic
**Description:** Create bio-engineered creatures and enhance creature stats.
**Skill Trees:** DNA Sampling I–IV, Creature Engineering I–IV, Cloning I–IV, Tissue Engineering I–IV
**Benefits:** Sample DNA from creatures; combine DNA to create enhanced pets; craft tissue for armor and enhancement items
**XP Type:** Bio-Engineering XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** Niche market for creature handlers and crafters

#### 8.3.7 Creature Handler
**Prerequisites:** Master Scout
**Description:** Tame, train, and command creatures in combat and utility roles.
**Skill Trees:** Taming I–IV, Training I–IV, Command I–IV, Beast Mastery I–IV
**Benefits:** Tame wild creatures (up to CL 20 at Master); train pets to learn abilities (attacks, buffs, tricks); command up to 3 pets simultaneously at Master
**XP Type:** Creature Handling XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** Combat and utility (pets for farming, tanking)

#### 8.3.8 Ranger
**Prerequisites:** Master Scout
**Description:** Advanced wilderness survival, tracking, and area-denial specialist. Full dedicated specification in Section 25.
**Skill Trees:** Tracking I–IV, Trapping I–IV, Camouflage I–IV, Wilderness Survival I–IV
**Benefits:** Track players (PvP functionality); place traps for area denial; enhanced camps provide stronger buffs; camouflage ability (stealth in wilderness)
**XP Type:** Ranger XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** PvP specialist, group support

#### 8.3.9 Pistoleer
**Prerequisites:** Master Marksman
**Description:** Dual-wield pistols, high burst damage specialist.
**Skill Trees:** Pistol Accuracy I–IV, Pistol Speed I–IV, Pistol Support I–IV, Pistol Finesse I–IV
**Benefits:** Dual-wield pistols; high attack speed (fastest ranged profession); special attacks (headshot, body shot)
**XP Type:** Pistol Combat XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** PvP and PvE combat

#### 8.3.10 Rifleman
**Prerequisites:** Master Marksman
**Description:** Long-range, high-damage specialist.
**Skill Trees:** Rifle Accuracy I–IV, Rifle Speed I–IV, Rifle Support I–IV, Rifle Finesse I–IV
**Benefits:** Highest single-target damage; long-range attacks (50m+ effective range); special attacks (sniper shot, overcharge)
**XP Type:** Rifle Combat XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** PvE and PvP DPS specialist

#### 8.3.11 Carbineer
**Prerequisites:** Master Marksman
**Description:** Balanced ranged combat, area-effect specialist.
**Skill Trees:** Carbine Accuracy I–IV, Carbine Speed I–IV, Carbine Support I–IV, Carbine Finesse I–IV
**Benefits:** Area-effect attacks (cone, line AoE); balanced damage and speed; suppression abilities (reduce enemy accuracy)
**XP Type:** Carbine Combat XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** PvE and PvP group combat

#### 8.3.12 Fencer
**Prerequisites:** Master Brawler
**Description:** One-handed melee, high finesse and defense.
**Skill Trees:** Fencing Accuracy I–IV, Fencing Speed I–IV, Fencing Defense I–IV, Fencing Finesse I–IV
**Benefits:** High defense (parry/dodge); fast attack speed; special attacks (lunge, riposte)
**XP Type:** Fencing Combat XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** PvP duelist, solo PvE

#### 8.3.13 Swordsman
**Prerequisites:** Master Brawler
**Description:** Two-handed melee, balanced offense and defense.
**Skill Trees:** Sword Accuracy I–IV, Sword Speed I–IV, Sword Defense I–IV, Sword Finesse I–IV
**Benefits:** Balanced damage and defense; special attacks (cleave, power strike); medium attack speed
**XP Type:** Sword Combat XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** PvE and PvP balanced melee

#### 8.3.14 Pikeman
**Prerequisites:** Master Brawler
**Description:** Polearm specialist, area-effect melee.
**Skill Trees:** Polearm Accuracy I–IV, Polearm Speed I–IV, Polearm Support I–IV, Polearm Finesse I–IV
**Benefits:** Area-effect melee attacks (cone attacks); knockdown abilities (crowd control); slower attack speed, higher damage
**XP Type:** Polearm Combat XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** PvE group combat, PvP crowd control

#### 8.3.15 Teras Kasi Artist (TKA)
**Prerequisites:** Master Brawler + Master in 3 other melee professions
**Description:** Elite unarmed combat specialist, highest melee damage.
**Skill Trees:** Meditative Techniques I–IV, Power Strikes I–IV, Defensive Techniques I–IV, Kata Mastery I–IV
**Benefits:** Highest unarmed damage; HAM regeneration abilities (meditation); complex attack chains (combos)
**XP Type:** Teras Kasi XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** Elite PvP and PvE melee

#### 8.3.16 Doctor
**Prerequisites:** Master Medic
**Description:** Advanced healing and buff specialist. Full dedicated specification in Section 24.
**Skill Trees:** Advanced Healing I–IV, Enhancement I–IV, Medicine Mastery I–IV, Diagnosis I–IV
**Benefits:** Craft enhancement buff packs (essential for difficult content); heal wounds efficiently; cure diseases and poisons; apply long-duration buffs
**XP Type:** Medical XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** Essential support role, high demand

#### 8.3.17 Combat Medic
**Prerequisites:** Master Medic
**Description:** Field healing and combat support specialist. Full dedicated specification in Section 24.
**Skill Trees:** Field Triage I–IV, Combat Medicine I–IV, Revivification I–IV, Battlefield Awareness I–IV
**Benefits:** Heal in combat without breaking stance; area-effect healing abilities; prevent player death (damage shields)
**XP Type:** Medical XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** Group PvE and PvP support

#### 8.3.18 Musician
**Prerequisites:** Master Entertainer
**Description:** Instrument performance specialist, Battle Fatigue healing. Full dedicated specification in Section 23.
**Skill Trees:** Instrumentation I–IV, Composition I–IV, Inspiration I–IV, Showmanship I–IV
**Benefits:** Heal Battle Fatigue faster than base Entertainer; unlock advanced instruments (Nalargon, Chidinkalu, etc.); provide group buffs during performances
**XP Type:** Entertainer XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** Social gameplay, essential for BF healing

#### 8.3.19 Dancer
**Prerequisites:** Master Entertainer
**Description:** Dance performance specialist, Battle Fatigue healing. Full dedicated specification in Section 23.
**Skill Trees:** Dance Mastery I–IV, Choreography I–IV, Exotic Dances I–IV, Stage Presence I–IV
**Benefits:** Heal Battle Fatigue faster than base Entertainer; unlock advanced dances (Exotic Dance 1–4, Jedi Meditation, etc.); provide group buffs during performances
**XP Type:** Entertainer XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** Social gameplay, essential for BF healing

#### 8.3.20 Image Designer
**Prerequisites:** Master Entertainer
**Description:** Character customization specialist.
**Skill Trees:** Holoemote Design I–IV, Body Modification I–IV, Facial Modification I–IV, Hair Styling I–IV
**Benefits:** Modify player appearance (hair, face, body); create custom holoemotes; earn income from cosmetic services
**XP Type:** Image Designer XP | **Skill Point Cost:** ~80 points to Master
**Market Role:** Cosmetic services, niche market

#### 8.3.21 Bounty Hunter
**Prerequisites:** Master Marksman OR Master Scout + 4 other Master professions
**Description:** Player tracking, high single-target damage, PvP specialist.
**Skill Trees:** Investigation I–IV, Pursuit I–IV, Weapon Mastery I–IV, Capture I–IV
**Benefits:** Track players (PvP bounty missions); access to unique weapons (DL-44 metal, DXR-6, etc.); special attacks (wrist rocket, poison, knockdown); mission terminals for NPC and player bounties
**XP Type:** Bounty Hunter XP | **Skill Point Cost:** ~120 points to Master (elite profession)
**Market Role:** PvP specialist, rare profession

#### 8.3.22 Commando
**Prerequisites:** Master Marksman + 3 other Master professions
**Description:** Heavy weapons specialist, area-effect damage.
**Skill Trees:** Heavy Weapons I–IV, Demolitions I–IV, Weapon Engineering I–IV, Tactical Assault I–IV
**Benefits:** Use heavy weapons (flamethrower, rocket launcher, acid rifle, lightning cannon); area-effect attacks (high damage, slow attack speed); craft specialized ammo
**XP Type:** Commando XP | **Skill Point Cost:** ~120 points to Master (elite profession)
**Market Role:** PvE group farming, PvP siege

#### 8.3.23 Smuggler
**Prerequisites:** Master Pistoleer + Master Scout
**Description:** Spice crafting, slicing (item improvement), smuggling missions.
**Skill Trees:** Spice Production I–IV, Slicing I–IV, Underworld Contacts I–IV, Evasion I–IV
**Benefits:** Craft spices (temporary stat buffs, healing items); slice weapons/armor to improve stats (one-time permanent boost); access smuggling missions (high-risk, high-reward)
**XP Type:** Smuggler XP | **Skill Point Cost:** ~120 points to Master (elite profession)
**Market Role:** Niche market for slicing services, spice crafting

### 8.4 Profession Interdependencies

**Crafting Dependency Chain:**

```text
Architect needs:
├─ Metal (from Scout/Miner)
├─ Electronics (from Droid Engineer or Architect)
└─ Structure Maintenance (from self or other Architects)

Armorsmith needs:
├─ Metal/Hide/Bone (from Scout/Miner)
├─ Armor Segments (from Armorsmith or component crafter)
└─ Bio-Engineered Tissues (from Bio-Engineer, optional)

Weaponsmith needs:
├─ Metal/Polymer (from Scout/Miner)
├─ Power Cells (from Droid Engineer)
└─ Weapon Barrels/Scopes (from Weaponsmith component schematics)

Droid Engineer needs:
├─ Electronics (from Droid Engineer or Architect)
├─ Metal (from Scout/Miner)
└─ Droid Frames (from Droid Engineer component schematics)

Chef needs:
├─ Organic Resources (from Scout or Bio-Engineer)
└─ Additives (from Organic resources)

Tailor needs:
├─ Hide/Petrochem (from Scout/Miner)
└─ Synthetic Cloth (from Tailor component schematics)
```

**Combat Dependency Chain:**

```text
All Combat Professions need:
├─ Weapons (from Weaponsmith)
├─ Armor (from Armorsmith)
├─ Buffs (from Doctor)
├─ Wound Healing (from Combat Medic)
└─ Battle Fatigue Healing (from Entertainer)
```

**Support Dependency Chain:**

```text
Doctor needs:
├─ Medical Components (from Medic or Doctor crafting)
└─ Organic Resources (from Scout)

Entertainer needs:
├─ Instruments (from Architect or loot)
└─ Cantinas (from Architect)

Merchant needs:
├─ Goods to Sell (from all Crafters)
└─ Vendor Structures (from Architect)
```

This interdependency is **ESSENTIAL**. AI agents must ensure no single profession can bypass these chains.

---

## 9. Combat System

### 9.1 Combat Philosophy

**Core Principle:** Combat is tactical, resource-driven (HAM pools), and deterministic (stat-based outcomes, not twitch reflexes). Combat emphasizes preparation (buffs, equipment, positioning) over moment-to-moment skill.

**Player Experience Goals:**

- Strategic ability selection (manage HAM costs vs. damage output)
- Equipment choices matter (weapon type, armor, food buffs)
- Group dynamics encourage role specialization (tank, DPS, healer)
- Death has meaningful consequences (items decay, Battle Fatigue accumulates)

### 9.2 Combat Core Mechanics

#### 9.2.1 HAM Pool System (Health/Action/Mind)

**Foundation:** All combat actions (attacks, abilities) cost HAM and deal HAM damage. Characters are incapacitated when any one pool reaches zero.

**Primary Pools:**

- **Health Pool:** Damaged by most attacks, represents physical injury
- **Action Pool:** Consumed by abilities, damaged by some attacks (drains stamina)
- **Mind Pool:** Consumed by abilities, damaged by fear/stress attacks

**Secondary Attributes (within each pool):**

- Strength/Constitution (Health pool): Increase max health, reduce incoming damage
- Quickness/Stamina (Action pool): Increase attack speed, reduce action costs
- Focus/Willpower (Mind pool): Increase ability power, reduce mind costs

**Incapacitation:**

- When ANY pool reaches 0, character is incapacitated (cannot act, can be revived)
- Incapacitation timer: 5 minutes (can be revived by Combat Medic during this time)
- Death after timer: Clone at clone facility, lose 5% item condition, gain Battle Fatigue

**Wounds:**

- Wounds reduce maximum HAM pool size (permanent until healed by Medic)
- Wounds accumulate from incapacitation and certain attacks
- Example: 500 Health Wounds reduce max Health from 5000 to 4500

**Battle Fatigue:**

- Battle Fatigue reduces maximum HAM pool size (permanent until healed by Entertainer)
- Battle Fatigue accumulates from deaths and combat actions
- Example: 20% Battle Fatigue reduces max Health from 5000 to 4000

#### 9.2.2 Combat Resolution (Attack Flow)

**Step-by-Step Attack Resolution:**

1. Attacker selects target and ability
2. Check ability HAM costs (deduct from attacker's pools)
3. Check range (weapon range vs. distance to target)
4. Check line of sight (no obstructions)
5. Calculate hit chance:
   - Base accuracy (weapon + skill bonuses)
   - Target defense (armor, posture, buffs)
   - Environmental modifiers (terrain, weather)
6. Roll 1d100 vs. hit chance
7. If hit:
   - Calculate damage (weapon damage + skill modifiers)
   - Apply armor mitigation (damage reduction)
   - Apply resistance modifiers (kinetic, energy, elemental)
   - Deduct damage from target's appropriate HAM pool
   - Check for special effects (knockdown, stun, DoT)
8. Apply weapon condition decay (small % per attack)

**Implementation Formula (Simplified):**

```text
Hit Chance = Base Weapon Accuracy + Attacker Accuracy Skill - Target Defense Skill - Target Posture Bonus
Damage = (Weapon Min-Max Damage + Skill Damage Bonus) × (1 - Armor Mitigation %) × Resistance Multiplier
```

#### 9.2.3 Postures & Stances

**Postures (affects defense and movement):**

- **Standing:** Default, balanced defense and movement
- **Kneeling:** +10% ranged accuracy, −20% movement speed, −10% melee defense
- **Prone:** +20% ranged accuracy, −50% movement speed, −30% melee defense, harder to hit at range
- **Crouched:** +5% melee defense, −10% movement speed

**Stances (affects playstyle, cannot switch during combat):**

- **Normal:** Balanced offense and defense
- **Aggressive:** +10% damage, −10% defense
- **Defensive:** −10% damage, +10% defense
- **Berserk:** +20% damage, −20% defense, +10% attack speed

**Design Rules:**

- Posture changes take 1–2 seconds (animation lock)
- Stance changes require 5-second channel (out of combat only)
- Interrupting posture/stance change resets action

#### 9.2.4 Special Attacks & Abilities

**Ability Categories:**

- **Basic Attacks:** Auto-attack, low HAM cost, low damage
- **Special Attacks:** High HAM cost, high damage, cooldowns
- **Crowd Control:** Knockdown, stun, snare (disable enemy actions)
- **Buffs/Debuffs:** Increase ally stats or reduce enemy stats
- **Area-of-Effect (AoE):** Damage or debuff multiple targets
- **Damage-Over-Time (DoT):** Apply bleeding, poison, burning effects

**Example Abilities:**

*Rifleman – Overcharge Shot:* Cost: 500 Action, 200 Mind | Damage: 300–600 (high burst) | Cooldown: 30 seconds | Effect: Single-target, high damage, high accuracy

*Pikeman – Sweep Attack:* Cost: 400 Action, 100 Health | Damage: 150–300 (AoE cone, 90 degrees) | Cooldown: 20 seconds | Effect: Knockdown on hit (3-second disable)

*Combat Medic – Bacta Injection:* Cost: 300 Mind | Healing: 500–1000 Health | Cooldown: 15 seconds | Effect: Instant heal, single-target

**Design Rules:**

- Cooldowns prevent spamming (high-damage abilities on 20–60 second cooldowns)
- HAM costs balance power (expensive abilities drain resources)
- No global cooldowns (can queue abilities, limited by HAM pools)

#### 9.2.5 Weapon Types & Damage Types

**Weapon Categories:**

*Melee:* Unarmed (fast, low damage, no decay) · One-Handed/Fencing (fast, medium damage, high finesse) · Two-Handed/Swords (medium speed, high damage, balanced) · Polearms/Pikes (slow, very high damage, AoE potential)

*Ranged:* Pistols (fast, medium damage, dual-wield capable) · Rifles (slow, very high damage, long range) · Carbines (medium speed, medium damage, AoE attacks)

*Heavy Weapons (Commando-only):* Flamethrower (short range, AoE, fire DoT) · Rocket Launcher (long range, AoE, knockdown) · Acid Rifle (medium range, acid DoT, armor debuff) · Lightning Cannon (medium range, chain lightning, mind damage)

**Damage Types:**

- **Kinetic:** Physical damage (blasters, bullets, melee)
- **Energy:** Blaster energy damage (most blasters)
- **Elemental (Fire/Acid/Electric):** DoT damage, special effects
- **Stun:** Non-lethal, drains Action/Mind pools

**Armor Resistances:** Each armor type has different resistances — Composite Armor: balanced kinetic/energy (50%/50%) · Bone Armor: high kinetic, low energy (70%/30%) · Chitin Armor: high energy, low kinetic (30%/70%) · RIS Armor (high-end): customizable resistances via experimentation

#### 9.2.6 Combat States & Conditions

**States:** Incapacitated (HAM pool at 0, cannot act, can be revived) · Dead (incapacitation timer expired, respawn at clone facility) · Stunned (cannot act 2–5 seconds) · Knocked Down (cannot act 3 seconds, must stand up) · Snared (movement −50–90%) · Blinded (accuracy −50%) · Dizzy (cannot use abilities 5 seconds)

**Damage-Over-Time (DoT):** Bleeding (20–50 dmg/tick, 3s ticks, 15–30s) · Poison (30–80 dmg/tick, 20s, reduces healing received) · Fire (40–100 dmg/tick, 10s, spreads to nearby targets) · Disease (10–30 dmg/tick, 60s, lowers max HAM by 5%)

**Cure Conditions:** Combat Medic abilities can cure DoTs and conditions; stim packs can cure bleeding/poison; time-based expiration (DoTs expire naturally)

### 9.3 Group Combat Mechanics

#### 9.3.1 Grouping Benefits

- **XP Bonus:** +10% XP per group member (max +50% at 5+ members)
- **Loot Sharing:** Group loot distribution (round-robin, master looter, or need/greed)
- **Threat Management:** Tanks generate threat to protect squishier members
- **Buff Sharing:** Group buffs (Doctor, Combat Medic, Ranger camps)

#### 9.3.2 Threat/Aggro System

**Threat Mechanics:** Creatures attack highest-threat target; damage generates threat (1 damage = 1 threat); healing generates threat (1 healing = 0.5 threat); taunts generate fixed threat (500–2000 threat)

**Tank Role:** Fencer/Swordsman taunts pull aggro; high health/defense survive focused damage; positioning (body-blocking) protects ranged DPS

#### 9.3.3 Healer Role

Combat Medic heals wounds mid-combat; Doctor buffs prevent incapacitation (larger HAM pools); Entertainer pre-combat (heal Battle Fatigue before engagement)

#### 9.3.4 DPS Role

Rifleman/Pistoleer/Carbineer maximize damage output; manage HAM costs to sustain damage; focus fire on called targets

### 9.4 PvP Combat

#### 9.4.1 Factional PvP (Rebel vs. Imperial)

**Flagging:** Players declare factional allegiance (Rebel or Imperial); "Overt" status enables PvP (attackable by opposing faction); "Covert" status disables PvP (safe from attack, cannot attack enemies)

**Switching Overt/Covert:** Toggle via command (`/pvp`); 5-minute timer to switch from Overt to Covert (prevents combat logging); attacking enemy immediately flags Overt

**Factional Points:** Earn points for killing enemy players; points unlock factional ranks (recruit → sergeant → major → colonel → general); ranks unlock base building permissions, factional rewards

#### 9.4.2 Open-World PvP

**Rules:** PvP only enabled if both players are Overt (same faction = safe); no PvP in NPC cities (safe zones); PvP unrestricted in wilderness and player cities (if city allows)

**Death Penalties:** Lose 10% item condition (double PvE death penalty); accumulate Battle Fatigue (requires Entertainer healing); drop small amount of credits

#### 9.4.3 Base Destruction (Factional Warfare)

**Mechanics:** Faction bases can be placed in wilderness; opposing faction can attack and destroy bases; destroying base grants factional points to attackers; base defenders earn points for defending (killing attackers)

**Base Vulnerability:** Bases vulnerable during declared war times (set by owning guild); bases have structure HP (destroyed when HP reaches 0); turrets defend base (NPC automated defenses)

### 9.5 Creature Combat (PvE)

#### 9.5.1 Creature Difficulty (Combat Level)

**CL Rating:** Creatures rated CL 1 to CL 90+; CL determines health, damage, defense, and XP rewards; players engage creatures near their effective CL (sum of skill modifiers)

**Difficulty Scaling:** Green/trivial (creature 10+ CL below player) · Blue/easy (5–10 CL below) · White/moderate (±5 CL) · Yellow/challenging (5–10 CL above) · Red/dangerous (10+ CL above)

#### 9.5.2 Creature Types

**Categories:** Herbivores — low aggression, flee when attacked (Banthas, Dewbacks) · Carnivores — aggressive, attack on sight (Rancors, Krayts) · Scavengers — attack if low health or in packs (Worrt, Squall) · Humanoids — NPCs (Tusken Raiders, Stormtroopers), use weapons and tactics

**Special Creature Abilities:** Knockdown (creatures knock players prone) · Fear (reduces player accuracy/defense) · Poison (applies DoT) · Pack Tactics (bonus damage when multiple creatures attack same target)

#### 9.5.3 Lairs & Spawning

**Lair Mechanics:** Creatures spawn from lairs (nest, den, warren); destroying lair stops spawns in that location; lairs have HP (destroyed by dealing damage); destroying lair grants bonus XP

**Spawn Behavior:** Lairs spawn 3–10 creatures based on lair type; creatures patrol radius around lair; aggro range varies by species (5m to 80m)

### 9.6 Death & Respawn

#### 9.6.1 Incapacitation & Revive

**Incapacitation:** When any HAM pool reaches 0, player is incapacitated; incapacitation lasts 5 minutes (or until revived); other players can revive via Combat Medic abilities

**Revive Process:** Combat Medic uses revive ability; player returns to standing with 10% HAM pools; player gains Wounds and Battle Fatigue

#### 9.6.2 Death

**Death Trigger:** Incapacitation timer expires without revive; player killed by environmental damage (fall damage, lava)

**Death Consequences:** Respawn at designated clone facility (player chooses location); lose 5% condition on all equipped items; gain 1–3% Battle Fatigue (requires Entertainer healing); gain 50–200 Wounds (requires Medic healing); drop small amount of credits (PvP: 10% of carried credits)

#### 9.6.3 Clone Facilities

**Clone Binding:** Players bind to clone facility (designate respawn point); facilities in NPC cities (free); facilities in player cities (require city rank 4+)

**Implementation:** Player opens clone facility interface, selects "bind here"; on death, player respawns at bound location; can change binding at any clone facility

### 9.7 Combat Balance Targets

#### 9.7.1 Damage Scaling

**Guidelines:** Starting weapon: 50–150 damage per hit · Mid-tier weapon: 150–400 damage per hit · End-game weapon: 300–700 damage per hit · Maximum crit: 1000–1500 damage (rare, requires buffs + high-end weapon)

**HAM Pool Scaling:** Unbuffed player: 3000–6000 Health (depending on profession) · Buffed player: 6000–12000 Health (Doctor buffs + food) · Boss creatures: 20,000–100,000 Health

**Time-to-Kill (TTK):** Solo player vs. white-con creature: 20–40 seconds · Group (5 players) vs. elite creature: 60–120 seconds · PvP (1v1, equal skill/gear): 30–60 seconds

#### 9.7.2 Ability Balance

**Cooldown Guidelines:** Basic attacks — no cooldown (HAM-limited) · Special attacks — 10–20 second cooldown · Elite attacks — 30–60 second cooldown · Ultimate abilities — 2–5 minute cooldown

**HAM Cost Guidelines:** Basic attack — 10–50 Action · Special attack — 100–500 Action · Expensive abilities — 500–1000 Action + 200–500 Mind/Health

---

## 10. Crafting System

### 10.1 Crafting Philosophy

**Core Principle:** Crafting is the PRIMARY source of useful equipment in the game. Crafted items are superior to looted items in all cases. Crafting involves resource quality, experimentation, and player skill, creating a competitive economy.

**Player Experience Goals:**

- Crafting is engaging gameplay (experimentation creates variety)
- Resource scarcity creates market value (rare resources = valuable items)
- Crafters build reputations for quality (high-experimentation results sought after)
- Economic gameplay is viable as primary activity (crafters earn income competitive with combat)

### 10.2 Crafting Core Loop

**Step-by-Step Crafting Process:**

1. Acquire schematic (from trainer or looted)
2. Gather required resources (survey/harvest or bazaar purchase)
3. Open crafting tool (generic tool or specialized tool for profession)
4. Select schematic
5. Allocate resources to schematic slots
6. Customize item (name, description)
7. Experimentation phase (allocate experimentation points to attributes)
8. Craft item (roll results based on experimentation)
9. Repeat via prototype or factory production

### 10.3 Schematics

#### 10.3.1 Schematic Sources

**Acquisition Methods:**

- **Trainer Schematics:** Learned from profession trainers (most common)
- **Looted Schematics:** Rare drops from creatures (unique items)
- **Quest Schematics:** Rewards from mission terminals (rare)
- **Reverse Engineering:** Deconstruct items to learn schematics (one-time use)

#### 10.3.2 Schematic Attributes

**Each schematic defines:** required resources (type and quantity); complexity (skill requirement); experimentation attributes (what can be customized); item category (weapon, armor, food, component, etc.); profession restriction (who can craft it)

**Example Schematic (DL-44 Blaster Pistol):**

```text
Schematic: DL-44 Metal Blaster Pistol
Profession: Weaponsmith (Pistol Specialization IV required)
Complexity: 25

Required Resources:
├─ Metal (Weapon Grade): 60 units
├─ Polymer (High-Grade): 20 units
├─ Blaster Power Handler (component): 1 unit
└─ Weapon Barrel (component): 1 unit

Experimentation Attributes:
├─ Damage (Min/Max): 10 experimentation points available
├─ Attack Speed: 8 experimentation points available
├─ Accuracy: 6 experimentation points available
└─ Weapon Effectiveness: 4 experimentation points available

Output: 1 DL-44 Blaster Pistol
```

### 10.4 Resources

#### 10.4.1 Resource Categories

**Primary Resource Types:** Ore & Metal (Iron, Aluminum, Copper, Steel — weapons, armor, structures) · Radioactive (Plutonium, Uranium — energy weapons, high-tech components) · Fiberplast (Petrochem polymers — armor, clothing, electronics) · Organic (Meat, Hide, Bone — food, clothing, Bio-Engineering) · Chemical (Acid, Liquid Petrochem — consumables, advanced components) · Gemstones (Crystals — decorations, lightsaber components — not MVP) · Water (Moisture, liquid — cooking, Bio-Engineering) · Flora (Wood, Plant Fiber — structures, furniture)

#### 10.4.2 Resource Attributes

**Every resource spawn has 10 stats (0–1000 scale):** Overall Quality (OQ) — average of all stats; Stat 1–9 — specific to resource type (e.g., Malleability for metals, Flavor for organics)

**Example Resource:**

```text
Resource Name: Lokian Aluminum (spawn ID: 482910)
Planet: Lok
Concentration: High (40% spawn density in region)

Stats:
├─ Overall Quality: 847
├─ Malleability: 923
├─ Conductivity: 812
├─ Unit Toughness: 789
├─ Shock Resistance: 901
├─ [... additional stats]

Spawned: Day 1 of cycle
Despawn: Day 10 of cycle
```

**Implementation:** Resource stats are RANDOMLY generated at spawn time; high-stat resources (OQ 900+) are rare (~5% of spawns); low-stat resources (OQ <500) are common (~40% of spawns)

#### 10.4.3 Resource Spawning System

**Spawn Cycle:** Resources spawn in geographic regions (5–20km radius circles); each spawn lasts 7–14 days (randomized per spawn); when spawn despawns, new spawn appears elsewhere on planet; multiple resource types can spawn simultaneously on same planet

**Spawn Density:** Concentration ranges from 1% to 90% (affects harvesting rate); high-concentration spawns (70%+) are rare and valuable; spawns can overlap (same location has multiple resource types)

**Design Rules:** No two identical spawns (stats are unique each time); resource names are procedurally generated (planet name + resource type + ID); spawn locations are not telegraphed (players must survey to find)

### 10.5 Surveying & Sampling

*Full player-facing workflow, tool progression, and Scout/Ranger bonuses are specified in Section 26. This subsection defines the underlying mechanic.*

#### 10.5.1 Survey Tools

**Tool Types:** Mineral Survey Device (locate ore/metal spawns) · Chemical Survey Device (locate radioactive/chemical spawns) · Flora Survey Device (locate wood/plant spawns) · Organic Survey Device (locate meat/hide/bone spawns) · Water Survey Device (locate moisture/water spawns)

**Usage:** Equip survey tool; use tool (generates waypoint to nearest spawn); tool displays concentration % and distance; repeat survey to triangulate exact location; use sampling tool at location to extract sample

#### 10.5.2 Sampling

**Sample Tool:** Extracts 1 unit of resource per use (2–5 second animation); resource stats revealed after sampling; inventory storage (resources stack by spawn ID)

**Sample Efficiency:** Basic sampling — 1 unit per use; improved tools (crafted) — 2–3 units per use; diminishing returns (same location) — concentration decreases after repeated sampling

#### 10.5.3 Resource Data Sharing

**Implementation:** Players can share waypoints (spawn locations); bazaar listings display resource stats; third-party tools can track spawn history (historical accuracy)

### 10.6 Harvesting Structures

#### 10.6.1 Harvesters

**Harvester Types:** Personal Harvester (small, low yield, 50 units/hour) · Medium Harvester (moderate yield, 100 units/hour) · Heavy Harvester (high yield, 200 units/hour)

**Placement:** Place harvester deed at surveyed location; harvester extracts resources automatically (hourly tick); hopper stores extracted resources (max capacity: 50,000 units for heavy)

**Maintenance:** Weekly maintenance fee (paid from harvester maintenance pool); if maintenance lapses, harvester shuts down; harvesters can be destroyed (PvP or decay)

#### 10.6.2 Harvester Management

**Operations:** Empty hopper (transfer resources to inventory); add maintenance credits; redeed harvester (pack up and move); set permissions (allow guild members to access)

**Design Rules:** Harvesters only extract from ONE spawn (must move when spawn despawns); harvesters operate 24/7 (passive income); harvesters are visible to all players (can be located and attacked in PvP zones)

### 10.7 Experimentation

#### 10.7.1 Experimentation Mechanics

**Process:**

1. After resource allocation, enter experimentation phase
2. Schematic lists experimentation attributes (e.g., Damage, Speed, Accuracy)
3. Player allocates experimentation points (limited pool, e.g., 10 points total)
4. Each attribute can receive multiple points; each point increases chance of success on that attribute
5. Click "Experiment" button
6. Game rolls for each attribute:
   - Critical Success (10%): +5–10% improvement
   - Success (50%): +2–5% improvement
   - Normal (30%): +0–2% improvement
   - Failure (8%): No change
   - Critical Failure (2%): −2–5% penalty
7. Results displayed (e.g., "Damage experimentation: Success! +4%")
8. Repeat experimentation (up to 3–5 rounds depending on schematic)

#### 10.7.2 Experimentation Point Allocation Strategy

**Trade-offs:** Allocate all points to one attribute (maximize one stat); spread points across attributes (balanced item); focus on market demand (players want damage > accuracy for rifles)

**Tools Modify Success:** Basic crafting tool — standard success rates; advanced crafting tool — +10% success chance; specialized station (structure) — +20% success chance

#### 10.7.3 Experimentation Results Impact

**Example Results:**

```text
Base DL-44 Stats (before experimentation):
├─ Min Damage: 200
├─ Max Damage: 600
├─ Attack Speed: 2.5 seconds
├─ Accuracy: +15

After Experimentation (focused on Damage):
├─ Min Damage: 220 (+10% from critical success)
├─ Max Damage: 660 (+10% from critical success)
├─ Attack Speed: 2.5 seconds (no points allocated)
├─ Accuracy: +16 (+6% from minor success)

Market Value: High (damage-focused meta)
```

#### 10.7.4 Resource Quality Impact on Experimentation

**High-Quality Resources:** Resources with high relevant stats improve base item stats BEFORE experimentation; example: High Malleability metal increases base weapon damage; experimentation THEN applies % boosts on top of improved base

**Formula (Simplified):**

```text
Base Stat = Schematic Base + (Resource Quality × Stat Multiplier)
Final Stat = Base Stat × (1 + Experimentation Bonus %)
```

**Example:**

```text
DL-44 Damage Calculation:
├─ Schematic Base Damage: 200-600
├─ Resource Metal OQ: 900 (high quality)
├─ Damage Multiplier: 0.5 (50% contribution from resource)
├─ Improved Base: 200 + (900 × 0.5) = 650 Min, 600 + (900 × 0.5) = 1050 Max
├─ Experimentation Bonus: +10%
└─ Final Damage: 715-1155
```

This creates MASSIVE value variance: Low-quality resource (OQ 300) → 200–600 damage (no improvement); High-quality resource (OQ 900) → 715–1155 damage (nearly 2x better)

### 10.8 Factory Production

#### 10.8.1 Factory Mechanics

**Purpose:** Mass-produce items using schematic and resources without manual crafting.

**Setup:** Place factory deed (structure, requires power/maintenance); load schematic into factory; load resources into factory hopper; set run quantity (e.g., 100 units); start production

**Production Rate:** 1 item per hour per factory run slot (factories have 5–10 slots); factories can run 24/7 (passive production)

**Factory Output:** All items in a factory run have IDENTICAL stats (no per-item experimentation variance); experimentation happens ONCE on prototype (results apply to all factory items)

#### 10.8.2 Prototype System

**Prototype Workflow:** Craft prototype manually (with full experimentation); evaluate prototype results; if satisfied, create factory schematic from prototype; load factory schematic into factory for mass production

**Design Philosophy:** Prototype experimentation determines quality of entire batch; bad prototype = bad batch (incentive to craft multiple prototypes and pick best); factory production is time-efficient but resource-intensive

#### 10.8.3 Factory Maintenance

**Operating Costs:** Weekly maintenance fee (similar to harvesters); power costs (require power generators nearby); hopper capacity limits (requires periodic resource refilling)

**Factory Management:** Check production status (remaining items in run); empty output hopper (collect finished items); load new runs (queue production)

### 10.9 Crafting Professions Detail

#### 10.9.1 Armorsmith (Full Detail)

**Schematic Categories:** Personal Armor (wearable armor: chest, legs, arms, helmet) · Layered Armor (advanced armor with sockets for enhancements) · Armor Segments (crafted components used in advanced armor)

**Armor Types:** Bone Armor (high kinetic resist, low energy resist, lightweight) · Composite Armor (balanced resists, medium weight) · Chitin Armor (high energy resist, low kinetic resist) · Ubese Armor (balanced, high-end, expensive) · RIS Armor (highest-tier, modular, requires advanced components)

**Experimentation Attributes:** Armor Rating (damage reduction %, primary) · Encumbrance (weight penalty, affects HAM pool modifiers) · Durability (condition decay rate) · Special Resistances (elemental resist: fire, acid, electric)

**Resource Dependencies:** Metal (iron, aluminum, steel) for plates; Hide/Bone for organic armor; Armor Segments (crafted components from self or others)

**Market Demand:** All combat players need armor (consumable due to decay); high-quality armor (high armor rating, low encumbrance) commands premium prices; specialty armor (high elemental resist) for specific content (e.g., Dathomir fire-resist armor)

#### 10.9.2 Weaponsmith (Full Detail)

**Schematic Categories:** Melee Weapons (swords, vibroblades, axes, polearms) · Ranged Weapons (pistols, rifles, carbines) · Heavy Weapons (limited — components for Commando weapons) · Weapon Components (barrels, scopes, power handlers)

**Weapon Types:** Pistols (DL-44, CDEF, Scout Blaster, etc.) · Rifles (DXR-6, T-21, Tusken Rifle, etc.) · Carbines (E-11, DH-17, etc.) · Melee (Vibroblade, Power Hammer, Lance, etc.)

**Experimentation Attributes:** Damage (Min/Max) — primary damage output · Attack Speed — seconds between attacks (lower = better) · Accuracy — hit chance modifier · Weapon Effectiveness — condition decay rate (durability)

**Resource Dependencies:** Metal (weapon-grade) for weapon bodies; Polymer for grips, stocks; Power Handlers (component) for energy weapons; Barrels (component) for customization

**Market Demand:** Weapons decay with use (constant demand); high-damage weapons command premium prices; specialized weapons (high accuracy for PvP, high damage for PvE) serve different markets

#### 10.9.3 Architect (Full Detail)

**Schematic Categories:** Structures (houses, guild halls, cantinas, hospitals) · Furniture (chairs, tables, containers) · Harvesters (resource extraction structures) · Factories (crafting automation structures) · City Structures (civic buildings: city halls, shuttleports)

**Structure Types:** Small House (1-room, low storage, cheap) · Medium House (3-rooms, moderate storage) · Large House (5+ rooms, high storage, expensive) · Guild Hall (large, multi-purpose, high capacity)

**Experimentation Attributes:** Structure Capacity (storage space, decoration limits) · Maintenance Cost (weekly upkeep reduction) · Durability (condition decay rate)

**Resource Dependencies:** Metal (structural iron, aluminum); Wood (flora resources); Electronics (crafted components)

**Market Demand:** Housing is one-time purchase (low repeat demand); harvesters/factories are high-value, recurring purchases (destroyed/redeeded); city structures are guild/community purchases (high-value, low-volume)

#### 10.9.4 Droid Engineer (Full Detail)

**Schematic Categories:** Droid Chassis (R-series, protocol, combat droids) · Droid Modules (combat modules, crafting modules, storage modules) · Droid Components (power cells, armor, motivators)

**Droid Types:** R-series/Utility (storage, harvesting assistance) · Protocol Droids (merchant assistance, translation) · Combat Droids (pet-like combat assistance, lower power than Creature Handler pets)

**Experimentation Attributes:** Module Capacity (how many modules droid can equip) · Droid Speed (movement speed) · Droid HAM (durability in combat)

**Resource Dependencies:** Electronics (high-quality for modules); Metal (chassis construction); Droid Components (self-crafted or from other Engineers)

**Market Demand:** Moderate demand (droids are long-term investments); combat droids less popular than Creature Handler pets (balance issue)

#### 10.9.5 Chef (Full Detail)

**Schematic Categories:** Food (stat-boosting consumables) · Drink (stat-boosting consumables) · Spices, via Smuggler (temporary buffs)

**Food/Drink Benefits:** Provide HAM pool buffs (e.g., +500 Health for 30 minutes); provide stat buffs (e.g., +50 Strength); stackable with Doctor buffs

**Experimentation Attributes:** Buff Duration (how long buff lasts) · Buff Potency (strength of buff) · Filling (how much it satiates, affects consumption rate)

**Resource Dependencies:** Organic resources (meat, wheat, vegetables); Additives (crafted or harvested)

**Market Demand:** Moderate demand (buffs consumed frequently by combat players); high-quality food (long duration, high potency) in demand for difficult content

#### 10.9.6 Tailor (Full Detail)

**Schematic Categories:** Clothing (shirts, pants, robes, jackets) · Wearables (backpacks, belts, bandoliers) · Accessories (jewelry, cosmetics)

**Clothing Benefits:** Primarily cosmetic (appearance customization); some clothing provides minor stat buffs (e.g., +10 camouflage)

**Experimentation Attributes:** Appearance (color customization) · Stat Buffs, minor (small bonuses to social stats or utility)

**Resource Dependencies:** Hide (organic); Synthetic cloth (crafted from petrochem)

**Market Demand:** Low demand (cosmetic-focused, one-time purchases); niche market for roleplayers and social players

---

## 11. Resource System

### 11.1 Resource Philosophy

**Core Principle:** Resources are the FOUNDATION of the economy. Resource quality and scarcity determine item value. Resources are DYNAMIC (spawn locations and stats change on timers), forcing adaptation and creating geographic economic value.

**Player Experience Goals:**

- Resource hunting is engaging exploration gameplay
- Resource quality variation creates competition (finding high-OQ spawns)
- Resource market is volatile and interesting (prices fluctuate based on spawns)
- Geographic knowledge has value (knowing good spawn locations)

### 11.2 Resource Spawn Algorithm

#### 11.2.1 Spawn Generation (Technical Specification for AI Agent)

**Spawn Timing:** Each planet has 20–50 active resource spawns at any time; each spawn lasts 7–14 days (randomize per spawn: 7 + rand(0–7) days); when spawn despawns, new spawn of SAME TYPE appears elsewhere on planet; spawn/despawn occurs during daily server maintenance (minimize player disruption)

**Spawn Location:** Spawn center is random point on planet surface (exclude NPC city zones); spawn radius is 5–20km (randomize per spawn); concentration gradient — 90% at center, tapering to 10% at edge (Gaussian distribution); multiple spawns can overlap (additive concentration)

**Spawn Stats Generation — [IMPLEMENTATION RECOMMENDATION]:**

```python
# Pseudocode for AI agent implementation
def generate_resource_spawn(resource_type, planet):
    spawn = {}
    spawn['type'] = resource_type
    spawn['planet'] = planet
    spawn['id'] = generate_unique_id()
    spawn['name'] = f"{planet.name} {resource_type} {spawn['id']}"

    # Generate 10 stats (0-1000 scale)
    spawn['stats'] = {}
    for stat in resource_type.stat_list:
        # Weighted toward mid-range (500), with rare high/low outliers
        base = random.normal(mean=500, stddev=150)
        spawn['stats'][stat] = clamp(base, 0, 1000)

    # Calculate Overall Quality (average of all stats)
    spawn['stats']['OQ'] = average(spawn['stats'].values())

    # Generate location and concentration
    spawn['center'] = random_point_on_planet(planet, exclude_npc_cities=True)
    spawn['radius_km'] = random.randint(5, 20)
    spawn['peak_concentration'] = random.randint(50, 90)

    # Set spawn/despawn times
    spawn['spawn_date'] = current_date()
    spawn['despawn_date'] = current_date() + days(7 + random.randint(0, 7))

    return spawn
```

#### 11.2.2 Spawn Rarity Curves

**OQ Distribution (target probabilities):** OQ 0–300, trash — 10% · OQ 300–500, low — 30% · OQ 500–700, average — 40% · OQ 700–900, good — 15% · OQ 900–1000, excellent — 5%

**Concentration Distribution:** 1–30%, scattered — 30% · 30–60%, moderate — 50% · 60–90%, dense — 20%

**Design Intent:** Excellent resources (OQ 900+, concentration 70%+) are RARE events (1–2 per planet per month); when excellent spawn appears, creates economic "gold rush" (competition for harvester placement)

#### 11.2.3 Resource Type Distribution

**Per Planet:** Each planet has 5–10 resource types that spawn (based on planetary composition) — Tatooine: Ore (iron, copper), Radioactive, Fiberplast (low organics) · Corellia: Ore (aluminum, copper), Organics (meat, hide), Flora · Naboo: Organics (high quality), Water, Flora · Lok: Rare metals (high OQ potential), Radioactive · Dathomir: Rare organics, Gemstones, Radioactive

**Design Rule:** Not all resource types spawn on all planets. Players must travel for specific resources (creates interplanetary trade).

### 11.3 Surveying Mechanics (Full Specification)

#### 11.3.1 Survey Tool Usage

**Process:**

1. Equip survey tool (specific to resource category: Mineral, Chemical, Organic, etc.)
2. Use tool (command: `/survey` or toolbar button)
3. Tool displays UI: Resource Type Dropdown (select specific resource, e.g., "Iron" within Mineral category); Range Slider (set search range, 5km to 100km); Survey Button (execute survey)
4. On survey: server checks for spawns of selected type within range; returns closest spawn with distance and concentration; creates waypoint to spawn center (approximate, ±500m variance)
5. Player travels to waypoint
6. Repeat survey at new location (triangulate exact high-concentration point)

**Survey Results Display:**

```text
Survey Result:
Resource: Tatooinian Iron (ID: 482910)
Distance: 2.3 km (Northwest)
Concentration: 67% (Good)
[Waypoint created]
```

**Cooldown:** 10-second cooldown between surveys (prevent spam)

#### 11.3.2 Sampling

**Process:**

1. Equip sample tool (generic or specialized)
2. Stand at location (survey first to confirm spawn presence)
3. Use sample tool (5-second animation)
4. Receive 1–3 units of resource (based on tool quality and concentration)
5. Resource added to inventory with full stat display

**Sample Display:**

```text
Sampled: Tatooinian Iron (ID: 482910)
Quantity: 2 units
Stats:
├─ Overall Quality: 734
├─ Malleability: 812
├─ Conductivity: 698
├─ Unit Toughness: 721
└─ [... additional stats]
```

**Diminishing Returns:** Sampling same location repeatedly decreases concentration over time (local depletion); concentration recovers slowly (1% per hour)

### 11.4 Harvester Mechanics (Full Specification)

#### 11.4.1 Harvester Placement

**Process:** Survey to find high-concentration spawn; acquire harvester deed (purchased from Architect or bazaar); place deed at desired location (must be within spawn radius); harvester structure appears (takes 1 minute to "build"); harvester begins operation (hourly resource extraction)

**Placement Restrictions:** Cannot place within 100m of another harvester (prevent clustering); cannot place in NPC cities or restricted zones; can be destroyed by other players in PvP zones

#### 11.4.2 Harvester Operation

**Extraction Rate:** Base rate — 50 units/hour (Personal), 100 units/hour (Medium), 200 units/hour (Heavy); multiplied by concentration at harvester location (e.g., 70% concentration = 140 units/hour for Heavy)

**Hopper Capacity:** Personal — 10,000 units · Medium — 25,000 units · Heavy — 50,000 units

**Hopper Management:** Players must manually empty hopper (transfer to inventory); if hopper fills, harvester stops extracting (no overflow)

#### 11.4.3 Harvester Maintenance

**Maintenance Pool:** Harvesters have credit pool (player deposits credits); weekly maintenance cost deducted from pool (500–5000 credits based on harvester type); if pool empty, harvester shuts down (no extraction)

**Power Requirements:** Harvesters require nearby power generator (separate structure); generator also has maintenance pool

**Lifespan:** Harvesters tied to spawn (when spawn despawns, harvester extracts nothing); players must redeed and move harvester to new spawn

#### 11.4.4 Harvester Permissions

**Access Control:** Owner sets permissions (personal, guild, public); guild members can empty hopper if permissions allow; public access allows anyone to empty (rare, used for community projects)

### 11.5 Resource Storage & Transportation

#### 11.5.1 Inventory Management

**Storage Limits:** Player inventory — 80 item slots; resources stack up to 100,000 units per stack (by spawn ID); different spawn IDs do NOT stack (Tatooinian Iron #482910 ≠ Tatooinian Iron #482911)

**Weight/Volume:** Resources have weight (affects movement speed if overburdened); vehicles/mounts increase carry capacity

#### 11.5.2 Storage Structures

**House Storage:** Houses have storage chests (capacity based on house size) — Small house: 200 slots · Medium house: 400 slots · Large house: 800 slots

**Storage Access:** Owner and designated friends can access house storage; resources safe from decay/theft in storage

#### 11.5.3 Resource Trading

**Bazaar Sales:** Resources listed on bazaar with stats displayed; buyers filter by resource type, OQ, specific stats; prices player-set (market determines value)

**Direct Trade:** Players can trade resources directly (face-to-face); no bazaar fee (encourages community interaction)

### 11.6 Resource Market Dynamics

#### 11.6.1 Price Drivers

**Supply Factors:** Current spawn quality (high OQ spawns increase supply of good resources); harvester placement competition (limited high-concentration spots); spawn despawn cycles (scarcity when good spawn despawns)

**Demand Factors:** Crafter demand (Weaponsmiths drive metal prices, Architects drive wood prices); market trends (flavor-of-the-month armor types shift demand)

**Historical Pricing (guideline for AI agent balance):** Trash resources (OQ <300) — 1–5 credits/unit · Average resources (OQ 500–700) — 10–30 credits/unit · Good resources (OQ 700–900) — 50–150 credits/unit · Excellent resources (OQ 900+) — 200–1000+ credits/unit

#### 11.6.2 Economic Emergent Behaviors (Historical)

**Resource Cartels:** Guilds monopolize high-OQ spawn locations (place multiple harvesters); control supply to manipulate prices

**Spawn Location Trading:** Players sell waypoints to excellent spawns (information economy)

**Resource Speculation:** Crafters stockpile resources during good spawns; sell when spawns despawn and prices spike

**Implementation Note:** AI agents should NOT hardcode prices. Prices emerge from player trading. Bazaar system only tracks transaction history.

---

## 12. Economy System

### 12.1 Economic Philosophy

**Core Principle:** The economy is PLAYER-DRIVEN. NPCs do not compete with players. Money sinks and faucets are balanced to create sustainable inflation (~5–10% annually). Economic gameplay (crafting, trading, merchanting) is viable primary content.

**Player Experience Goals:**

- Crafters earn competitive income with combat players
- Market dynamics are interesting (prices fluctuate based on supply/demand)
- Economic mastery is a skill (market knowledge, resource timing, vendor placement)
- New players can enter economy (low barriers to entry, mid-tier profits)

### 12.2 Currency & Money Faucets

#### 12.2.1 Credit Sources (Money Faucets)

**Primary Faucets:**

- **Mission Rewards:** 500–10,000 credits per mission (scales with difficulty)
- **NPC Vendor Sales:** Selling looted items to NPCs (10–100 credits per item)
- **Factional Payouts:** Killing enemy faction NPCs (50–500 credits)
- **Quest Rewards (limited):** One-time rewards from rare quests (1,000–50,000 credits)

**Design Rules:**

- Mission rewards are the PRIMARY faucet (incentivize mission gameplay over passive idling)
- NPC vendor sales exist as a low-value floor for loot (junk disposal), never a competitor to the player crafting economy — items sold to NPC vendors are destroyed, not resold, so this faucet cannot leak crafted supply back into circulation
- Faucets scale with player engagement, not with time played passively (no AFK income sources)
- Total faucet inflow across the population is monitored against total sink outflow (Section 12.4) to hold the target inflation band

#### 12.2.2 Secondary & Indirect Faucets

- **Bounty Payouts:** Bounty Hunter contracts pay a bounty fee funded by the target's faction or a contracting player, not printed from nothing — treated as a peer-to-peer transfer for inflation accounting, except the NPC-funded default bounty floor (small, ~200–1000 credits)
- **Treasure/Loot Credits:** Small credit drops from creature kills (10–200 credits per kill, CL-scaled), representing "found" wealth rather than salvage
- **City Salaries (optional, mayor-configurable):** Player cities may allocate treasury funds to salaried civic roles (militia, city workers) — this recirculates existing city tax revenue rather than creating new credits, so it is a transfer, not a faucet

### 12.3 Money Sinks

**Player Experience Goal:** Spending money should feel like investment (into gear, structures, mobility, or services) rather than a punitive tax, while still removing enough currency from circulation to counteract faucets and crafted-item deflation over time.

**Primary Sinks:**

- **Structure Maintenance:** Weekly upkeep on harvesters, factories, vendors, houses, and city structures (Sections 11.4.3, 13, 21) — the single largest sustained sink, since it scales with the size of a player's economic footprint
- **Skill Training Costs:** Credits paid to trainers alongside XP when purchasing a skill box (Section 7.3) — scales with skill tier (10–50 credits for early boxes, 5,000–20,000 credits for Master boxes)
- **Travel Costs:** Shuttle and starport tickets (Section 20) — small, frequent sink that scales with playtime
- **Bazaar/Vendor Fees:** Listing fees and sales commission (Section 22.2) — scales with transaction volume, acting as a proportional tax on the crafted economy itself
- **Item Repair & Reload:** Weapon/armor condition restoration, ammunition purchase (component-based, drawing on crafted goods more than raw credits, but crafters' raw-resource costs are themselves a credit sink one layer upstream)
- **City Founding & Structure Deed Costs:** One-time credit costs to purchase house/harvester/factory/city-structure deeds from crafters (a transfer to crafters, but crafters in turn spend on resources and maintenance, so a portion recirculates out of the credit pool via their own sinks)
- **Faction Base Costs:** Base deeds and turret upkeep (Section 15)

**Design Rules:**

- Sinks must be recurring (subscriptions on maintenance), not one-time, so that wealth cannot be permanently "parked" outside the economy's velocity
- No sink may be mandatory to the point of blocking play (a player who cannot afford maintenance loses the structure, not their ability to continue playing)
- Sink costs scale with the wealth tier of the activity (a Master crafter's Heavy Harvester costs more to maintain than a Novice's Personal Harvester), producing a soft progressive cost curve

### 12.4 Inflation Control & Economic Balance

**Target:** ~5–10% credit-supply growth annually, tracked as (total faucet inflow − total sink outflow) / total credits in circulation, measured server-wide on a rolling 30-day window.

**Mechanics:**

- **[IMPLEMENTATION RECOMMENDATION]** A server-side economic telemetry job aggregates all faucet and sink transactions nightly and computes net inflation. This feeds Section 34's economic health dashboard.
- If measured inflation exceeds the target band for a sustained period, the design lever is to increase sink costs (maintenance fees, training costs) rather than reduce faucets, since faucets are tied to gameplay incentives (missions, combat) that should not be nerfed reactively
- Player-to-player transfers (bazaar sales, direct trades, guild dues) are excluded from faucet/sink accounting — they do not create or destroy credits, only move them, but they are tracked separately for wealth-distribution metrics (Section 34)

**Design Rules:**

- Credit duplication of any kind is a critical-severity bug class, not a balance issue — see Edge Cases below
- NPC vendor buy-back prices must remain low enough that "vendor-scumming" looted junk is never competitive with crafted-item income (reinforces Pillar 3)

### 12.5 Wealth Distribution & Economic Roles

**Player Experience Goal:** Multiple viable economic roles should coexist without one dominating: harvester-barons (resource supply), master crafters (production), merchants (distribution/logistics), and combat players (resource demand + mission faucet income) should each be able to sustain a character economically.

**Roles:**

- **Producer (Scout/Miner):** Converts time and travel into raw resources; income scales with willingness to explore dangerous/remote spawns
- **Manufacturer (Crafter):** Converts resources + skill + experimentation into finished goods; income scales with reputation and schematic access
- **Distributor (Merchant):** Converts logistics and market knowledge into margin; income scales with vendor network size and market information
- **Consumer (Combat/Social players):** Converts mission/service income into demand for crafted goods; funds the entire chain from the top of Section 4.4's cycle

**Implementation Requirements:**

- No profession's income ceiling should exceed roughly 2–3x another's at equivalent time investment (Pillar 2's horizontal philosophy applied to economics, not just combat)
- Market data (Section 22.4) must be visible enough that new players can identify which economic role is currently under-supplied and fill it

### 12.6 Data Entities

```text
Entity: CreditLedgerEntry
- entry_id: UUID
- account_id: UUID (character or structure owner)
- amount: int (positive = credit, negative = debit)
- category: enum [mission_reward, npc_sale, faction_payout, maintenance_fee,
                   training_cost, travel_cost, bazaar_fee, vendor_sale,
                   direct_trade, structure_purchase, faction_base_cost, other]
- is_faucet_or_sink: enum [faucet, sink, transfer]
- counterparty_id: UUID (nullable; other character/structure/NPC in transfer)
- timestamp: datetime
- location: {planet, x, y, z}

Entity: EconomicSnapshot (server telemetry, generated nightly)
- snapshot_date: date
- total_credits_in_circulation: bigint
- faucet_total_30d: bigint
- sink_total_30d: bigint
- net_inflation_pct_30d: float
- gini_coefficient: float (wealth distribution, see Section 34)
- top_resource_prices: [{resource_type, avg_price_per_unit}]
```

### 12.7 System Interactions

- Section 9 (Combat) feeds the mission-reward faucet
- Section 10–11 (Crafting/Resources) generate the goods that give credits somewhere to go
- Section 13–14 (Housing/City) and Section 21 (Vendor) generate the largest recurring sinks
- Section 22 (Bazaar) is the primary transfer mechanism and the best source of live price data
- Section 15 (Faction) contributes both a faucet (payouts) and a sink (base costs)

### 12.8 Edge Cases

- **Credit/item duplication exploits:** All currency and item state changes must be atomic, server-authoritative transactions (see Section 29.5); a crafting, trading, or mailing action that can be interrupted mid-transaction (client crash, disconnect) must roll back completely rather than partially apply
- **Abandoned structures with full maintenance pools:** If an owner never returns, credits sit locked in a maintenance pool indefinitely — **[ASSUMPTION]** after 180 days of owner inactivity, structure and pool are reclaimed and the structure is removed, returning the lot to the world (not a credit sink event, since those credits were already removed from "active" circulation)
- **New player economic bootstrapping:** A brand-new character with no crafted gear and no credits must still be able to earn starting capital via NPC-sale loot and starter missions without needing another player's charity — starter missions and starting equipment are tuned accordingly (cross-ref Section 27.2)
- **Wealth concentration by guilds/cartels:** Section 11.6.2's resource cartels are historically accurate emergent behavior and should not be mechanically prevented, only monitored via the Gini coefficient metric (Section 34); this is a sandbox outcome, not a bug

### 12.9 Implementation Priorities

- **[MVP]** Credit ledger, all primary faucets and sinks, basic inflation telemetry
- **[MVP]** Structure maintenance sink (harvester, factory, vendor, house)
- **[EXPANSION]** Full economic dashboard with Gini coefficient and market-trend visualizations for players (a simplified bazaar price-history view is MVP; the full analytics suite is not)

---

## 13. Housing System

### 13.1 Player Experience Goals

Housing in Pre-CU SWG is not a cosmetic side feature — it is a primary end-game activity. Players should feel that decorating, provisioning, and maintaining a house is as legitimate a way to spend a play session as combat or crafting. A house is simultaneously a storage solution, a vendor location, a status symbol, and a piece of the player-driven world simulation (Pillar 4). Placing a house should feel like a meaningful claim on the world, not a menu transaction.

### 13.2 Mechanics

#### 13.2.1 Structure Placement

1. Player purchases a house deed (crafted by an Architect, or bought on the bazaar)
2. Player selects placement mode and finds a valid, unobstructed patch of terrain outside NPC city limits and outside another structure's no-build radius
3. Client renders a placement ghost (green = valid, red = invalid) based on terrain slope, structure collision, and zoning rules
4. Player confirms placement; server validates and instantiates the structure; a short "construction" timer (60–120 seconds, **[ASSUMPTION]**) plays before the structure becomes enterable
5. Placing player is automatically the structure's Owner

#### 13.2.2 House Tiers **[HISTORICAL]**

| Tier | Footprint | Storage Slots | Rooms | Relative Cost |
|---|---|---|---|---|
| Small House | Small | 200 | 1 | Low |
| Medium House | Medium | 400 | 3 | Moderate |
| Large House | Large | 800 | 5+ | High |
| Guild Hall | Very Large | 1000+ | Multi-purpose | Very High, guild-funded |

#### 13.2.3 Interior & Decoration

- Houses have an explorable interior (cell-based, see Section 29.2) separate from the exterior world coordinate space
- Furniture (crafted by Architects) can be placed with free-form position/rotation within interior bounds
- Decoration count is capped per house tier to bound rendering/storage load (**[IMPLEMENTATION RECOMMENDATION]**: cap similar to historical ~200–300 placed items for a Large house)
- Houses support "public" or "private" entry flags; public houses can be entered by any player (common for player-run shops or museums), private houses require an admin/friend permission check

#### 13.2.4 Permissions

- **Owner:** full control (decorate, admit/ban, sell items from house vendor, redeed structure)
- **Co-Owner/Admin:** can decorate and manage storage, cannot redeed or transfer ownership
- **Friends List (Entry Only):** can enter and use storage if granted, cannot decorate
- **Banned List:** explicitly denied entry even if the house is public

#### 13.2.5 Maintenance & Decay

- Houses have a maintenance pool funded like harvesters/factories (Section 11.4.3); weekly upkeep is small relative to income (housing should never be the dominant sink — Section 12.3)
- If maintenance lapses, the house enters a "condemned" grace period (14 days, **[ASSUMPTION]**) during which the owner is warned on login; if unpaid past the grace period, the structure and its full contents are destroyed
- This is a deliberately harsh, historically accurate consequence and must be prominently warned about in UI, not softened in mechanics

### 13.3 Data Entities

```text
Entity: Structure (base type, extended by House, Harvester, Factory, Vendor, CivicStructure)
- structure_id: UUID
- owner_character_id: UUID
- deed_type: string (schematic reference)
- location: {planet, x, y, z, heading}
- cell_map_id: UUID (interior layout reference)
- condition_pct: float (0-100)
- maintenance_pool_credits: int
- maintenance_cost_weekly: int
- last_maintenance_tick: datetime
- status: enum [active, condemned_grace_period, destroyed]
- admin_list: [character_id]
- entry_permission: enum [public, friends_only, private]
- friends_list: [character_id]
- banned_list: [character_id]

Entity: House (extends Structure)
- tier: enum [small, medium, large, guild_hall]
- storage_capacity: int
- placed_items: [PlacedItem]
- vendor_id: UUID (nullable, see Section 21)

Entity: PlacedItem
- item_id: UUID
- structure_id: UUID
- position: {x, y, z, rotation}
- item_reference: UUID (links to crafted Item entity, Section 28)
```

### 13.4 System Interactions

- Section 8.3.3 (Architect) is the sole source of house, furniture, and structure deeds — reinforcing Pillar 1
- Section 21 (Vendor) structures are frequently placed adjacent to or inside houses
- Section 14 (City) structure density and zoning rules constrain where houses may be placed within city limits
- Section 11.5.2 already defines house storage tiers; this section is the authoritative source, Section 11 cross-references it

### 13.5 Edge Cases

- **Placement conflicts:** two players attempting to place a structure in the same location within the same server tick — server resolves by transaction order (first validated placement wins; the second receives an "invalid placement" rejection, not a duplicate structure)
- **Owner account deletion/ban:** structure enters condemned state immediately (skips warning grace period) and is reclaimed after a shorter fixed window (30 days, **[ASSUMPTION]**) to avoid permanently squatting valuable terrain
- **House full of items during redeed:** redeeding (packing up) a house with items in storage is blocked; owner must clear storage first, preventing accidental item loss
- **Guild Hall ownership transfer:** must transfer to the guild entity, not an individual officer, so leadership turnover doesn't strand the structure (see Section 19.4)

### 13.6 Implementation Priorities

- **[MVP]** Small/Medium/Large house placement, storage, maintenance/decay, permissions, basic decoration
- **[MVP]** House-embedded vendor support (Section 21)
- **[EXPANSION]** Guild Halls, advanced decoration item variety, house-as-museum public browsing tools

---

## 14. City System

### 14.1 Player Experience Goals

Player cities are the clearest expression of Pillar 5 (Meaningful Social Structures): a city is a shared, persistent investment that only exists because a community chose to build it, fund it, and govern it. Founding, growing, and governing a city should feel like guild-level community achievement, not an individual real-estate purchase. A thriving city should visibly outperform the NPC starter cities in convenience (shuttle access, trainers, cantina activity) as a direct reward for community cohesion.

### 14.2 Mechanics

#### 14.2.1 City Founding

1. A player (the "Mayor-Founder") purchases a City Hall deed from an Architect
2. At least 10 player structures (any mix of houses/harvesters/factories) must exist within the city's defined radius **[HISTORICAL]**, either pre-existing or placed after founding within a grace period
3. Founder places the City Hall deed; the game instantiates the city as Rank 1 ("Outpost")
4. The city automatically enrolls all citizens whose structures fall within its radius; citizens may opt out

#### 14.2.2 City Ranks & Unlocks **[ASSUMPTION — exact historical thresholds undocumented; the following are authoritative for this project]**

| Rank | Name | Citizen/Upkeep Threshold | Unlocks |
|---|---|---|---|
| 1 | Outpost | Founding minimum | City Hall, basic zoning |
| 2 | Township | 20+ structures, treasury funded 30 days | Player-built Cantina, Medical Center |
| 3 | City | 40+ structures, active mayor | Shuttleport (fast travel), trainer hosting |
| 4 | Metropolis | 75+ structures, sustained treasury surplus | Cloning Facility (respawn point), city specialization slot |

#### 14.2.3 Governance

- **Elections:** weekly voting period; any citizen may run for mayor; citizens vote once per election; highest vote count wins a 7-day term (**[ASSUMPTION]** on exact term length)
- **Mayor Powers:** set tax rate (flat fee or % of local vendor sales), approve/deny new structure placement within city limits, appoint militia (city-sanctioned PvP defenders), set city specialization (Rank 4+), evict citizens for non-payment of city tax
- **Politician profession [HISTORICAL — added here to complete the profession roster referenced in Section 8.1]:** an elite profession (prerequisite: Master Entertainer or Master Merchant, **[ASSUMPTION]**) that grants bonus city management tools: reduced city maintenance costs, expanded zoning tools, and the exclusive ability to unlock city specialization bonuses at Metropolis rank. A city is not required to have a Politician mayor, but gains no specialization bonus without one.
- **City Specialization (Rank 4 only):** the mayor (with a Politician) selects one specialization (Military, Bio-Tech, Diplomacy, Outfitting, Municipal) granting a passive bonus to matching activity within city limits (e.g., Outfitting reduces local crafting station fees) — **[EXPANSION]**, not required for MVP

#### 14.2.4 City Maintenance

- Weekly upkeep is deducted from the city treasury (funded by citizen taxes and structure fees)
- If treasury cannot cover upkeep, city rank is downgraded one tier; if a city at Rank 1 cannot pay, it is dissolved and all citizen structures revert to standalone (still owned, just no longer part of a city)
- Downgrading removes the rank-gated structure (e.g., losing Rank 3 removes shuttleport access) but does not destroy citizen-owned houses/harvesters

### 14.3 Data Entities

```text
Entity: City
- city_id: UUID
- name: string
- planet: string
- center: {x, y, z}
- radius_m: int
- rank: enum [outpost, township, city, metropolis]
- treasury_credits: int
- upkeep_cost_weekly: int
- tax_rate: float
- specialization: enum [none, military, bio_tech, diplomacy, outfitting, municipal]
- mayor_character_id: UUID (nullable)
- mayor_term_end: datetime
- founding_date: datetime
- civic_structures: [structure_id]  (references Structure entities, Section 13.3)

Entity: CityCitizen
- city_id: UUID
- character_id: UUID
- joined_date: datetime
- tax_paid_to_date: int
- is_militia: bool

Entity: CityElectionBallot
- election_id: UUID
- city_id: UUID
- candidate_character_id: UUID
- voter_character_id: UUID
- cast_at: datetime
```

### 14.4 System Interactions

- Depends on Section 13 (Housing) for the structure-count founding threshold
- Unlocks Section 20 (Travel) fast-travel nodes at Rank 3+
- Unlocks Section 9.6.3 (Clone Facility) respawn binding at Rank 4
- Interacts with Section 15 (Faction) — a city may lean Rebel, Imperial, or Neutral based on citizen faction composition, affecting which faction's base-building is tolerated nearby (**[ASSUMPTION]**)
- Politician profession cross-links to Section 8 (Profession System)

### 14.5 Edge Cases

- **Overlapping city radii:** a new City Hall deed cannot be placed such that its radius overlaps an existing city's radius, preventing annexation disputes by placement race
- **Mayor goes inactive mid-term:** if a mayor does not log in for 14 days (**[ASSUMPTION]**), an emergency election is triggered early rather than waiting for term end
- **Structure count drops below threshold after founding** (players abandon/redeed houses): city does not auto-dissolve from structure count alone post-founding, only from treasury failure — this prevents cascading collapse from a few departures
- **Tie vote in election:** resolved by earliest-registered candidacy (**[ASSUMPTION]**), not randomly, for reproducibility/audit

### 14.6 Implementation Priorities

- **[MVP]** City founding, Rank 1–3, elections, tax collection, treasury/upkeep, shuttleport unlock
- **[MVP]** Politician profession only as a lightweight tax-reduction perk; full toolset can trail
- **[EXPANSION]** Rank 4 Metropolis, city specialization bonuses, militia sanctioning tools

---

## 15. Faction System

### 15.1 Player Experience Goals

The Galactic Civil War should feel like an optional, opt-in layer of conflict layered on top of the sandbox, not a mandatory endgame. A player should be able to spend a hundred hours entirely neutral, or dive into factional warfare from day one. Faction gameplay rewards commitment (rank, base access) but never locks a neutral player out of core systems (crafting, housing, entertaining all remain fully available to Neutrals).

### 15.2 Mechanics

#### 15.2.1 Faction Alignment

- On character creation or at any time thereafter via a faction recruiter NPC, a player may declare Rebel, Imperial, or remain Neutral
- Switching factions is allowed but resets faction rank progress and imposes a cooldown (**[ASSUMPTION]** 30 days) to prevent rapid switching for tactical advantage
- Neutral players cannot be flagged Overt and cannot participate in factional PvP, base assaults, or faction-rank rewards, but suffer no other restriction

#### 15.2.2 Overt/Covert Status

Full flagging rules are specified in Section 9.4.1; this section covers the faction-standing layer built on top of that flag state.

#### 15.2.3 Faction Points & Ranks **[HISTORICAL, exact thresholds ASSUMED]**

| Rank | Title | Points Required (cumulative) | Unlocks |
|---|---|---|---|
| 0 | Recruit | 0 | Faction chat, base blueprints (basic) |
| 1 | Sergeant | 2,500 | Faction vendor discount (10%) |
| 2 | Major | 10,000 | Faction-exclusive weapon schematics |
| 3 | Colonel | 30,000 | Base turret upgrades, militia command |
| 4 | General | 75,000, and sponsorship by an existing General | Faction leadership tools, city specialization influence |

- Faction points are earned from enemy player kills (while both parties Overt), enemy NPC kills, and successful base assaults/defenses
- Rank 4 (General) requires **sponsorship**: an existing General must vouch for the candidate, reinforcing Pillar 5's "ranks unlock when sponsored by higher-ranked players"
- Faction points decay slowly if a player is fully inactive from factional content for 60+ days (**[ASSUMPTION]**), preventing rank stagnation from one-time point farming

#### 15.2.4 Base Building & Destruction

Base structure/HP/turret mechanics are specified in Section 9.4.3. This section adds the faction-standing gate: base deeds require the owning guild's founder to hold at least Major rank, and base placement is subject to the same overlap/zoning validation as player cities (Section 14.5).

#### 15.2.5 Faction Perks

- Vendor discounts at faction-aligned NPC vendors (rank-scaled, 10–30%)
- Access to faction-exclusive weapon/armor schematics at Major+ (cosmetically distinct, not a Pillar 2 violation — power stays within the 3–5x band, exclusivity is aesthetic/flavor, not a stat outlier)
- Faction chat channel (Section 18.2) for coordinating with aligned players galaxy-wide

### 15.3 Data Entities

```text
Entity: FactionStanding
- character_id: UUID
- faction: enum [rebel, imperial, neutral]
- points: int
- rank: enum [recruit, sergeant, major, colonel, general]
- overt_status: bool
- overt_expires_at: datetime (nullable; the 5-minute covert-switch timer from Section 9.4.1)
- last_switch_date: datetime
- sponsor_character_id: UUID (nullable; required for General rank)

Entity: FactionBase
- base_id: UUID
- owning_guild_id: UUID
- faction: enum [rebel, imperial]
- location: {planet, x, y, z}
- structure_hp: int
- structure_hp_max: int
- turret_ids: [structure_id]
- war_declared_windows: [{start, end}]
```

### 15.4 System Interactions

- Builds directly on Section 9.4 (PvP Combat) flagging and base mechanics
- Section 14.4 notes city-level faction leaning influenced by citizen composition
- Section 16 (Mission System) includes faction-specific bounty missions available only to aligned players
- Section 8.3.21 (Bounty Hunter) missions frequently target Overt enemy-faction players

### 15.5 Edge Cases

- **Neutral player attacked in a contested zone:** never possible — Overt combat requires both participants to be Overt and opposing-faction; a Neutral player is never a valid PvP target regardless of location
- **Faction switch mid-base-siege:** disallowed; a character with a pending war-window base assault cannot switch factions until the assault window closes
- **Sponsor goes inactive/leaves faction before sponsorship completes:** sponsorship request expires after 14 days and candidate must find a new sponsor
- **Guild with mixed-faction membership owning a base:** disallowed at base-deed placement time — a base's owning guild must be single-faction-aligned at the officer level (**[ASSUMPTION]**)

### 15.6 Implementation Priorities

- **[MVP]** Faction alignment, Overt/Covert flagging (shared with Section 9.4), faction points, ranks through Colonel, vendor discounts
- **[EXPANSION]** Full base building/sieging, General rank + sponsorship, faction-exclusive schematics, city specialization influence

---

## 16. Mission System

### 16.1 Player Experience Goals

Mission terminals exist to give players a reason to travel, fight, or serve without ever becoming an authored quest chain (explicitly out of scope, Section 1.2). A mission should feel like a procedurally-posted job board: generic, repeatable, and disposable, never narratively significant. The player experience goal is "always something to do near the terminal" rather than "follow the story."

### 16.2 Mechanics

#### 16.2.1 Mission Terminal Types

- **Combat Terminal:** generates Destroy Lair, Assault (NPC camp), Recon (scout a location, low risk), and Bounty (NPC target) missions
- **Crafting Terminal (Artisan trainer-adjacent):** generates Delivery (craft and deliver N items) and Sample Collection missions
- **Bounty Hunter Terminal:** generates player-target bounty missions, visible only to Bounty Hunters (Section 8.3.21), sourced from faction bounty pools and player-posted contracts (Section 15)
- **Entertainer/Medic Terminal — [ASSUMPTION, low historical documentation]:** minor NPC-audience missions granting small bonus XP for performing/healing; intentionally low-impact so real player demand (Section 4.1) remains the primary loop for these professions

#### 16.2.2 Mission Generation Algorithm

```text
1. On terminal refresh (every 5-10 minutes, or on request up to a rate limit):
2. Select mission type weighted by terminal category
3. Select a target (lair / NPC camp / delivery destination) from nearby
   valid POIs within a travel-appropriate radius of the terminal
   (avoid missions requiring impractical cross-planet travel for a
   low-tier mission)
4. Scale difficulty to requesting player's effective combat level
   or crafting skill tier
5. Calculate reward: credits (Section 12.2.1 faucet), faction points
   (if applicable), and typed XP
6. Post mission to terminal list (typically 8-15 concurrent listings)
7. Mission expires after 2 hours if unclaimed, is regenerated
```

#### 16.2.3 Mission Lifecycle

1. Player reviews terminal listing (objective, location, reward, difficulty flag)
2. Player accepts mission (added to player's mission log, waypoint created)
3. Player completes objective in the open world (no instancing — Pillar 4/3.2)
4. Player returns to terminal (or a designated turn-in NPC) to collect reward, or reward auto-grants on objective completion for Recon/Sample types
5. Mission log slot frees up (cap of 3–5 concurrent missions, **[ASSUMPTION]**)

#### 16.2.4 Group Missions

- Combat missions above a difficulty threshold are flagged "Group Recommended"; any group member may accept on behalf of the group, and completion credit/reward is shared per Section 9.3.1's group bonus rules
- Group missions do not use instancing; the target lair/camp is the same shared-world object regardless of group size (Pillar 4)

### 16.3 Data Entities

```text
Entity: MissionTerminal
- terminal_id: UUID
- location: {planet, x, y, z}
- category: enum [combat, crafting, bounty_hunter, entertainer_medic]
- refresh_interval_sec: int

Entity: Mission
- mission_id: UUID
- terminal_id: UUID
- type: enum [destroy_lair, assault, recon, bounty_npc, bounty_player,
               delivery, sample_collection, performance, healing]
- target_reference: UUID (lair_id, npc_camp_id, or delivery target)
- difficulty_cl: int (nullable for non-combat)
- reward_credits: int
- reward_faction_points: int
- reward_xp: {xp_type: amount}
- posted_at: datetime
- expires_at: datetime
- accepted_by_character_id: UUID (nullable)
- accepted_by_group_id: UUID (nullable)
- status: enum [available, accepted, completed, expired, abandoned]
```

### 16.4 System Interactions

- Combat missions are the primary source of the Section 12.2.1 mission-reward faucet
- Destroy Lair missions interact directly with Section 17 (Creature System) lair mechanics
- Crafting/Delivery missions create artificial short-term demand that supplements (never replaces) the organic player-to-player demand of Section 12
- Bounty missions interact with Section 15 (Faction) standing and Section 8.3.21 (Bounty Hunter)

### 16.5 Edge Cases

- **Mission target destroyed by another player before accepter arrives:** Destroy Lair and Assault missions check target validity on turn-in; if the target no longer exists, the mission auto-completes at reduced reward (50%, **[ASSUMPTION]**) rather than becoming unsolvable
- **Player logs out mid-mission:** mission persists in the log across sessions until its expiry timer elapses; expiry is wall-clock, not play-time, to prevent indefinite mission hoarding
- **Group disbands mid-mission:** individual members retain credit for objectives already completed; the group-bonus reward multiplier is recalculated at turn-in based on who is still grouped
- **Delivery mission target player is offline/doesn't exist (crafting terminal delivery-to-NPC variant only):** all Delivery missions in this design target NPCs, never other players, to avoid griefing/holding missions hostage

### 16.6 Implementation Priorities

- **[MVP]** Combat Terminal (Destroy Lair, Assault, Recon), Crafting Terminal (Delivery)
- **[EXPANSION]** Bounty Hunter Terminal (requires Section 15 faction depth), Entertainer/Medic Terminal (low priority, minor system)

---

## 17. Creature System

### 17.1 Player Experience Goals

Creatures should feel like a living ecosystem, not a spawn table. A player exploring Dathomir should sense escalating danger through creature density and type before ever seeing a UI level indicator. Creatures should also be a resource: killing and harvesting a Bantha should feel materially different from destroying a droid, because one yields hide/meat/bone for Chefs, Tailors, and Bio-Engineers, and the other does not.

### 17.2 Mechanics

#### 17.2.1 Creature Template & Taxonomy

Every creature in the world is an instance of a **Creature Template** (data entity below), which defines its combat stats, AI behavior class, harvestable yield, and taming eligibility. Section 9.5 defines combat-facing behavior (CL, difficulty color, special abilities); this section defines the simulation and harvesting layer underneath it.

**AI Behavior Classes [HISTORICAL]:**

- **Passive/Herbivore:** does not initiate combat; flees when attacked or approached too closely
- **Aggressive/Carnivore:** attacks any player within aggro radius on sight
- **Scavenger:** aggressive only when in a pack or when the player is below a health threshold
- **Humanoid NPC:** uses profession-like combat abilities and may flee at low health rather than fight to the death

#### 17.2.2 Creature Harvesting **[HISTORICAL]**

- On death, a creature corpse can be harvested by any player (not just the killer, unless the spawning group's loot rule restricts it — Section 9.3.1) using the same sample-tool interaction as resource harvesting (Section 11.3.2)
- Yield type is fixed per creature template (a Bantha always yields Meat + Hide; a battle droid yields no organic resources, only occasional salvaged components)
- Yield quality (an OQ-like stat) is generated per-corpse at time of death, using the creature's template stat range, not the regional resource-spawn system — this is a separate, creature-specific quality roll
- Harvesting a corpse has a short window (10 minutes, **[ASSUMPTION]**) before the corpse despawns

#### 17.2.3 Lairs & Spawning

Full lair destruction/respawn mechanics are specified in Section 9.5.3 and Section 5.4.2. This section adds:

- Each lair template defines a **spawn table** (which creature templates, weighted) and a **max active population** (3–10)
- Lairs regenerate population over time up to the max if players kill individual creatures without destroying the lair itself, keeping "farming" viable without requiring full lair destruction every time
- Destroying the lair structure itself removes it permanently from that location; a new lair (of a valid type for the biome) spawns elsewhere on the planet within the 12–24 hour world-tick window (Section 5.4.2)

#### 17.2.4 Taming Integration

- Creature Handlers (Section 8.3.7) may attempt to tame any creature template flagged `tamable = true` and at or below their Command skill tier's CL cap
- A successful taming interaction converts the wild creature instance into an owned Pet entity bound to the taming character, removing it from the wild population count for that lair
- Bio-Engineers (Section 8.3.6) may sample DNA from a tamed or wild creature (harvestable window only) to unlock enhanced-stat pet variants — this is a distinct system layered on top of basic taming, not a prerequisite for it

### 17.3 Data Entities

```text
Entity: CreatureTemplate
- template_id: UUID
- name: string
- species_family: string (e.g., "Bantha", "Womp Rat", "Tusken Raider")
- ai_class: enum [passive, aggressive, scavenger, humanoid]
- cl_range: {min, max}
- ham_pools: {health, action, mind}
- damage_range: {min, max}
- special_abilities: [string]
- harvest_yields: [{resource_category, base_quantity_range, quality_range}]
- tamable: bool
- aggro_radius_m: int
- pack_size_range: {min, max}

Entity: CreatureInstance (a live, spawned creature)
- instance_id: UUID
- template_id: UUID
- lair_id: UUID (nullable, for free-roaming or event spawns)
- current_ham: {health, action, mind}
- location: {planet, x, y, z}
- state: enum [idle, alert, aggro, fleeing, dead]
- corpse_expires_at: datetime (nullable, set on death)

Entity: Lair
- lair_id: UUID
- planet: string
- location: {x, y, z}
- lair_hp: int
- lair_hp_max: int
- spawn_table: [{template_id, weight}]
- max_population: int
- current_population: int
- destroyed_at: datetime (nullable)

Entity: Pet (extends CreatureInstance ownership)
- pet_id: UUID
- owner_character_id: UUID
- name: string (player-assigned)
- trained_abilities: [string]
- loyalty_or_bond_stat: int (nullable, EXPANSION-tier depth)
```

### 17.4 System Interactions

- Section 9.5 (Creature Combat) is the combat-resolution layer that operates on CreatureInstance
- Section 10.4/10.9.5/10.9.6 (Chef, Tailor) depend on creature harvest yields for organic resources
- Section 8.3.6/8.3.7 (Bio-Engineer, Creature Handler) depend on the taming/DNA sampling hooks here
- Section 16 (Mission System) Destroy Lair missions reference Lair entities directly

### 17.5 Edge Cases

- **Multiple players harvesting one corpse:** first-harvest takes the full yield unless the group loot rule (Section 9.3.1) specifies split/round-robin; corpse is consumed on first successful harvest, not partially harvestable by multiple players
- **Lair destroyed while creatures from it are mid-fight elsewhere:** already-spawned CreatureInstances complete their current combat/flee behavior independently; only future spawning stops
- **Taming attempt on a creature already engaged in combat with another player:** taming interaction is blocked while the target is in an active aggro state toward a third party, preventing taming-as-combat-interruption abuse
- **Pack creature whose packmates are all killed:** last surviving pack member reverts to base aggressive/scavenger behavior rather than a permanently boosted "last stand" state, to avoid unintended difficulty spikes

### 17.6 Implementation Priorities

- **[MVP]** Creature templates, AI behavior classes, lair spawn/destroy cycle, basic corpse harvesting
- **[MVP]** Basic taming (Creature Handler)
- **[EXPANSION]** Bio-Engineer DNA sampling and enhanced pet variants, pack tactic depth, named/boss creature uniques

---

## 18. Social Systems

### 18.1 Player Experience Goals

Communication tools are load-bearing infrastructure for every other system in this document — the economy (Section 12), grouping (Section 19), and social professions (Sections 23–24) all depend on players being able to reliably reach each other. The goal is a chat and identity system invisible enough that players never think about it, freeing their attention for the systems it supports.

### 18.2 Mechanics

#### 18.2.1 Chat Channels **[HISTORICAL]**

| Channel | Range/Scope | Use Case |
|---|---|---|
| Spatial (`/say`) | ~20m radius | Local conversation, vendor advertising |
| Shout | ~50m radius | Wider local broadcast |
| Planet | Entire current planet | Planet-wide advertising, LFG |
| Group | Current group members, any location | Coordination |
| Guild | Current guild members, any location | Coordination, governance |
| Faction | Same-faction players, any location | Section 15 coordination |
| Tell (`/tell name`) | Single named player, any location | Private messaging |

#### 18.2.2 Mail System

- Asynchronous in-game mail between characters, deliverable whether the recipient is online or not
- Supports text body and limited item/credit attachments (used heavily by the crafted economy for remote transactions and gifting)
- Mailbox has a capacity cap (**[ASSUMPTION]** 50 messages); oldest read mail is flagged for cleanup, never auto-deleted with attachments still unclaimed

#### 18.2.3 Emotes & Holoemotes

- A standard emote library (wave, bow, dance, sit, etc.) available to all characters for baseline social expression
- Holoemotes (custom, player-designed) require Image Designer services (Section 8.3.20) and are a premium social/cosmetic good

#### 18.2.4 Waypoints & Friends

- Players may set personal waypoints (from survey results, mission accept, or manual map placement) that render as an on-screen/map directional marker
- Waypoints may be shared to another player or group via chat link, converting a spoken location into a clickable, importable marker (supports Section 11.6.2's information economy)
- A Friends List tracks online/offline status of favorited players across the galaxy, independent of guild membership

#### 18.2.5 Reputation Signals

- Vendor and crafter reputation (Pillar 5) is not a numeric score but an emergent signal: bazaar purchase history, vendor "visited" counts, and player word-of-mouth in chat channels — **[ASSUMPTION]** no formal star-rating system existed historically; this project deliberately does not add one, to avoid gamifying reputation in a way the original did not

### 18.3 Data Entities

```text
Entity: ChatMessage (ephemeral, not persisted long-term except Tell/Mail logs)
- channel: enum [spatial, shout, planet, group, guild, faction, tell]
- sender_character_id: UUID
- recipient_character_id: UUID (nullable, Tell only)
- body: string
- sent_at: datetime

Entity: MailMessage
- mail_id: UUID
- sender_character_id: UUID
- recipient_character_id: UUID
- subject: string
- body: string
- attached_credits: int
- attached_items: [item_id]
- sent_at: datetime
- read: bool
- claimed_attachments: bool

Entity: Waypoint
- waypoint_id: UUID
- owner_character_id: UUID
- label: string
- location: {planet, x, y, z}
- source: enum [survey, mission, manual, shared]

Entity: FriendsListEntry
- character_id: UUID
- friend_character_id: UUID
- added_at: datetime
```

### 18.4 System Interactions

- Underpins Section 12/22 (Economy/Bazaar) vendor advertising and remote-sale mail delivery
- Underpins Section 19 (Grouping) coordination
- Underpins Section 23/24 (Entertainer/Medic) service advertising in spatial/planet chat (explicitly named in Section 4.1)

### 18.5 Edge Cases

- **Mail attachment claim after item/structure referenced no longer exists** (e.g., sender's crafted item was in a house that got condemned before mail was claimed — not possible, since mailing an item transfers ownership out of the house at send time, so this cannot occur by construction)
- **Tell to an offline or nonexistent character name:** returns a clear client-side error, never silently fails
- **Waypoint sharing a location the recipient's client hasn't loaded (different planet):** waypoint still imports; distance/direction UI simply shows "different planet" until the recipient travels there

### 18.6 Implementation Priorities

- **[MVP]** All chat channels, mail with item/credit attachments, waypoints, friends list
- **[EXPANSION]** Holoemote creation tooling (depends on Image Designer, itself EXPANSION-tier per Section 8)

---

## 19. Grouping & Community Systems

### 19.1 Player Experience Goals

Grouping should be low-friction and immediately rewarding (Section 7.2.1's XP bonus), while guild membership represents a deeper, longer-term social commitment with real mechanical stakes (shared structures, factional identity, city citizenship overlap). The two systems are deliberately distinct: a group is a session-scale tool, a guild is a persistent institution.

### 19.2 Mechanics

#### 19.2.1 Group Formation

1. Any player may invite another player (or a group leader may invite on behalf of the group) via `/invite name` or a target-and-click UI action
2. Invitee accepts/declines; on accept, both are placed in a shared Group entity
3. Group leader (defaults to founder, transferable) controls loot rules and can disband or remove members
4. **[ASSUMPTION]** Group size cap: 20 players, matching the group-content-scaling assumptions of Section 9.3 and 7.2.1 (whose bonus formula caps its stated bonus at 5+ members, but larger raid-style groups for base assaults and boss lairs remain useful for coordination even past the XP-bonus cap)

#### 19.2.2 Loot Rules

- **Round-Robin:** loot rotates through group member order automatically
- **Master Looter:** group leader (or designated member) manually distributes all loot
- **Need/Greed:** each eligible member rolls or declares need/greed on drops; highest need wins, ties broken by greed roll

#### 19.2.3 Squad Leader Profession **[HISTORICAL — added here to complete the profession roster referenced in Section 8.1]**

- Elite profession, prerequisite Master Brawler or Master Marksman (**[ASSUMPTION]**)
- Grants group-wide command buffs (accuracy, defense, or HAM regeneration bonuses applied to all group members while the Squad Leader is present and in a specific "command stance")
- Does not replace Doctor/Musician buffs (Pillar 1) — Squad Leader buffs are smaller in magnitude and combat-stance-gated, layering on top of, not substituting for, crafted/social-profession support
- **[EXPANSION]** — valuable for large-group PvP/base content, not required for MVP's core loops

#### 19.2.4 Guilds (Player Associations)

- Any player may found a guild via an NPC guild registrar for a credit cost (a money sink, Section 12.3) and a minimum founding member count (**[ASSUMPTION]** 3)
- **Roles:** Guild Leader (full control, including disbanding the guild and transferring leadership), Officer (invite/kick members, manage guild hall permissions, cannot disband), Member (guild chat, guild hall access per permission)
- Guild membership is independent of, but commonly correlated with, City citizenship and Faction alignment
- Guild Halls (Section 13.2.2) are owned by the Guild entity itself, not an individual, so leadership turnover cannot strand the structure (Section 13.5)
- Guild "tax" — an optional, leader-configured percentage automatically diverted from member bazaar sales into a guild treasury (**[ASSUMPTION]**), used to fund the guild hall and faction base maintenance

#### 19.2.5 Mentorship **[ASSUMPTION — later-era feature; included as a light-touch, MVP-safe onboarding aid]**

An experienced player may declare a newer player (under a playtime/skill-point threshold) as a protégé, granting the mentor a small XP bonus and the protégé a temporary (7-day) small XP boost, encouraging veterans to actively play alongside newcomers rather than simply donating credits. This directly supports the Horizontal Progression Test (Section 1.3).

### 19.3 Data Entities

```text
Entity: Group
- group_id: UUID
- leader_character_id: UUID
- member_character_ids: [UUID] (max 20)
- loot_rule: enum [round_robin, master_looter, need_greed]
- created_at: datetime

Entity: Guild
- guild_id: UUID
- name: string
- tag: string (short chat prefix)
- leader_character_id: UUID
- officer_character_ids: [UUID]
- member_character_ids: [UUID]
- treasury_credits: int
- tax_rate_pct: float
- faction_alignment: enum [rebel, imperial, neutral, mixed]
- guild_hall_structure_id: UUID (nullable)
- founded_at: datetime

Entity: MentorshipBond
- mentor_character_id: UUID
- protege_character_id: UUID
- started_at: datetime
- expires_at: datetime
```

### 19.4 System Interactions

- Group XP bonus formula defined in Section 7.2.1; threat/tank/heal roles defined in Section 9.3
- Guild ownership underlies Guild Hall structures (Section 13) and Faction Bases (Section 15)
- City citizenship (Section 14.3) is tracked separately from guild membership but both display in social UI together

### 19.5 Edge Cases

- **Group leader disconnects without transferring leadership:** leadership auto-transfers to the next-longest-tenured group member after a short timeout (2 minutes), so loot/invite control is never orphaned
- **Guild leader account banned/deleted:** leadership auto-transfers to the highest-tenured Officer; if no officers exist, the guild enters a 30-day dissolution countdown during which any member may petition to become leader
- **Mentor/protégé relationship where mentor's skill points are lower than protégé's (protégé outleveled mentor):** bond remains valid until its 7-day expiry; the system does not re-validate the threshold mid-bond, since the intent is encouraging sustained co-play, not gatekeeping by skill snapshot
- **Group exceeding 20 during a merge (two groups try to combine):** merge is rejected with a clear error if it would exceed cap; leader must split before merging

### 19.6 Implementation Priorities

- **[MVP]** Group formation, all three loot rules, Guild founding/roles/chat/treasury
- **[MVP]** Mentorship (cheap to implement, high onboarding value)
- **[EXPANSION]** Squad Leader profession, guild-to-guild alliance tooling

---

## 20. Travel & Transportation

### 20.1 Player Experience Goals

Distance must remain meaningful (Section 5.5) throughout the game's systems, not just in the World Architecture overview. This section is the authoritative, implementation-level specification for every mode of movement, so that travel cost, speed, and access restrictions are consistent everywhere they are referenced (housing placement value, harvester logistics, mission travel, faction-restricted routes).

### 20.2 Mechanics

#### 20.2.1 Movement Modes

| Mode | Speed | Source | Notes |
|---|---|---|---|
| On Foot | 5–7 m/s | Always available | No cost, no cargo bonus |
| Speeder Bike / Swoop | 15–25 m/s | Crafted (Weaponsmith/Artisan-adjacent schematic, **[ASSUMPTION]** on exact crafting profession) or looted | Requires fuel or periodic maintenance (**[ASSUMPTION]**: simplified condition-decay model matching weapon/armor decay rather than a separate fuel-item economy) |
| Tamed Mount (Dewback, Kaadu) | 12–20 m/s | Creature Handler taming (Section 8.3.7) or Bio-Engineer | No decay, but requires the mount's own HAM upkeep like a pet |
| Landspeeder / Groundcar | 20–30 m/s | Crafted (Droid Engineer-adjacent, **[ASSUMPTION]**) | Higher cargo capacity than bikes |

#### 20.2.2 Fast Travel Network

- **Shuttleport (intra-planet, city-to-city):** instant travel between any two unlocked shuttleport nodes on the same planet; 5-minute real-time wait for next shuttle departure (**[HISTORICAL]**); costs a small credit fee (Section 12.3 sink)
- **Starport (inter-planet):** instant travel (loading transition) between planets; higher credit cost than shuttle; available at all NPC cities and Rank 3+ player cities (Section 14.2.2)
- No player-to-player teleportation or group summon exists at any point in this design (Pillar 4 — distance stays meaningful)

#### 20.2.3 Faction-Restricted Routes **[ASSUMPTION — historical specifics uncertain]**

An Overt player faces elevated NPC guard aggression when using a shuttleport/starport in a city with heavy opposing-faction NPC presence (e.g., an Overt Rebel using a heavily Imperial-garrisoned starport). This is modeled as increased hostile-NPC encounter risk near the terminal, not a hard travel block — Covert or Neutral players are never restricted.

### 20.3 Data Entities

```text
Entity: Vehicle
- vehicle_id: UUID
- owner_character_id: UUID
- vehicle_type: enum [speeder_bike, swoop, landspeeder, groundcar]
- condition_pct: float
- speed_mps: float
- cargo_capacity_bonus: int
- schematic_reference: UUID

Entity: TravelNode (shuttleport / starport)
- node_id: UUID
- node_type: enum [shuttleport, starport]
- location: {planet, x, y, z}
- connected_node_ids: [UUID]
- owning_city_id: UUID (nullable, null for NPC city nodes)
- fee_credits: int
- departure_interval_sec: int
```

### 20.4 System Interactions

- Section 14.2.2 gates shuttleport/starport hosting by city rank
- Section 12.3 counts every fare as a recurring credit sink
- Section 11.4/13 harvester and house placement value is directly a function of proximity to a TravelNode
- Section 9.6.3 clone-facility binding interacts with travel since players choose respawn points strategically near travel infrastructure

### 20.5 Edge Cases

- **Vehicle summoned inside a building/cell interior:** vehicles can only be summoned/mounted in exterior world space, never inside house or civic-structure interiors, avoiding physics/collision edge cases
- **Mount owner incapacitated while mounted:** player is automatically dismounted on incapacitation (Section 9.2.1) so revival mechanics function normally
- **Shuttleport destination city dissolved since last use:** stale destination is removed from the node list at the moment a city's rank drops below 3 (Section 14.6), not lazily at next use, so players are never shown an invalid destination
- **Two vehicles/mounts summoned simultaneously by one player:** only one active vehicle/mount per character is allowed; summoning a second automatically stores the first

### 20.6 Implementation Priorities

- **[MVP]** On-foot movement, crafted speeder bikes, shuttleport/starport network at NPC cities
- **[MVP]** Tamed mounts (depends on Creature Handler, itself MVP per Section 8)
- **[EXPANSION]** Faction-restricted route risk modeling, landspeeders/groundcars as a distinct higher-capacity tier

---

## 21. Vendor System

### 21.1 Player Experience Goals

A vendor is a crafter's storefront: a persistent, physical presence in the world that sells goods even while the owner is offline. Placing and stocking a vendor should feel like opening a shop, not filling out a listing form — the vendor is a structure players walk up to, browse, and buy from, reinforcing Pillar 3's "Vendor NPCs at player houses create distributed retail locations."

### 21.2 Mechanics

#### 21.2.1 Vendor Placement & Setup

1. Player purchases a Vendor deed (crafted by Architect or Merchant-adjacent schematic, **[ASSUMPTION]**) and places it as a structure (Section 13.2.1 placement flow), typically inside or adjacent to a house
2. Owner accesses the vendor's stocking interface to list items: sets price, quantity, and an optional item description
3. Vendor displays a "shopkeeper" NPC-like avatar (owner-customizable appearance, **[ASSUMPTION]**) that other players interact with to browse/buy
4. Any player walking up to the vendor can browse the full listing and purchase instantly (credits deducted, item transferred, no owner interaction required)

#### 21.2.2 Vendor Maintenance

- Vendors have a maintenance pool identical in structure to harvesters/houses (Section 11.4.3/13.2.5); weekly upkeep scales with number of active listings
- Merchant profession (Section 8.3.5) reduces vendor maintenance cost up to 50% at Master, directly rewarding specialization in the distribution economic role (Section 12.5)
- If maintenance lapses, the vendor stops accepting purchases (goes "closed") but does not immediately destroy stocked inventory, giving the owner a grace period (matching Section 13.2.5's condemned-state pattern) to refund the pool before item loss

#### 21.2.3 Vendor Inventory Management

- Base inventory slot cap (**[ASSUMPTION]** 50 for a Novice-placed vendor), increased by Merchant skill tiers
- Owner can restock, reprice, or pull items at any time from the management interface (in person or remotely if the owner has appropriate Merchant skill — **[ASSUMPTION]**: remote management requires at least mid-tier Merchant skill, reinforcing profession value)
- Sales proceeds accumulate in the vendor's till and must be collected by the owner (in person) — this, combined with maintenance, is why a Merchant profession's "reduce vendor fees / increase inventory" bonuses matter economically

### 21.3 Data Entities

```text
Entity: Vendor (extends Structure, Section 13.3)
- vendor_id: UUID
- structure_id: UUID
- inventory_slot_cap: int
- listings: [VendorListing]
- till_credits: int (accumulated, uncollected sales revenue)
- status: enum [open, closed_maintenance_lapsed]

Entity: VendorListing
- listing_id: UUID
- vendor_id: UUID
- item_reference: UUID
- price_credits: int
- quantity: int
- description: string
- listed_at: datetime
```

### 21.4 System Interactions

- Vendor structures are a Structure subtype (Section 13.3) and inherit placement/permission/maintenance rules
- Section 22 (Bazaar) search indexes every active VendorListing galaxy-wide — the Bazaar is not a separate inventory, it is a search layer over Vendor + Bazaar-stall listings (see Section 22.2)
- Section 8.3.5 (Merchant) profession directly modifies vendor economics
- Section 12.3 lists vendor maintenance as a primary sink

### 21.5 Edge Cases

- **Buyer purchases the last unit of a listing at the same instant as another buyer:** server-authoritative transaction lock on the listing row ensures only one purchase succeeds; the second buyer receives an immediate "sold out" response and is never charged
- **Owner deletes a listing while a buyer's purchase transaction is in flight:** the in-flight transaction either completes atomically against the pre-deletion state or is rejected outright — partial states (credits deducted, no item received) must be architecturally impossible (Section 29.5)
- **Vendor placed then immediately abandoned with full inventory:** covered by the same 180-day inactive-structure reclamation as Section 12.8, with proceeds/unsold items returned to an owner-accessible "reclaimed goods" mail if the owner ever returns within a longer secondary window (**[ASSUMPTION]**)

### 21.6 Implementation Priorities

- **[MVP]** Vendor placement, stocking, in-person purchase, till collection, maintenance
- **[EXPANSION]** Remote inventory management (Merchant skill-gated), vendor appearance customization

---

## 22. Bazaar & Market System

### 22.1 Player Experience Goals

The Bazaar is the galaxy-wide nervous system of the player economy: a player should be able to ask "who is selling Chitin armor under 500 credits anywhere in the galaxy" and get a real answer, without that convenience undermining the value of a well-placed physical vendor (Section 21) or the meaningfulness of distance (Pillar 3, Section 5.5.2). It resolves the tension between "the economy must be searchable" and "geography must still matter."

### 22.2 Mechanics

#### 22.2.1 Bazaar Terminals

- Physical terminal structures located in every NPC city and every Rank 2+ player city (cantina-adjacent, **[ASSUMPTION]**)
- Provide a searchable index of **every active VendorListing galaxy-wide**, plus **Bazaar Stall listings** (a lighter-weight sell option described below)
- Search supports filters: item category, resource type + stat-range filters (e.g., "Metal, Malleability > 800"), price range, planet/city, and seller name

#### 22.2.2 Bazaar Stalls (Consignment Selling Without a Full Vendor)

- A player without a placed Vendor structure may pay a smaller listing fee to consign an item directly through any Bazaar Terminal
- **[ASSUMPTION — historical implementation detail uncertain]** Item is held in escrow by the terminal system (not physically visible in the world) until sold; on sale, proceeds are mailed to the seller (Section 18.2.2) and the item is transferred to the buyer's inventory directly, without requiring either party to travel
- This distinguishes Stalls (convenience, escrow-based, small fee, capped listing duration e.g. 7 days) from Vendors (physical presence, unlimited duration while maintained, larger potential inventory, no escrow — buyer must travel to the structure)

#### 22.2.3 Fees

- **Listing Fee:** flat or item-value-scaled fee charged at listing time for Bazaar Stalls (Section 12.3 sink)
- **Sales Commission:** percentage (historically small, **[ASSUMPTION]** 4–6%) deducted from the sale price on a successful Stall sale; Vendor sales (Section 21) do not carry a commission beyond ordinary vendor maintenance, since the goods never enter Bazaar escrow

#### 22.2.4 Market Data Tools

- Every Bazaar Terminal query optionally returns recent transaction history (last-sold price, volume) for the searched item/resource type, satisfying Section 11.5.3 and Section 12.5's requirement that market data be queryable
- **[EXPANSION]** Full historical price-trend charts; MVP only requires last-N-sales visibility

### 22.3 Data Entities

```text
Entity: BazaarStallListing
- listing_id: UUID
- seller_character_id: UUID
- item_reference: UUID (item held in escrow)
- price_credits: int
- listing_fee_paid: int
- listed_at_terminal_id: UUID
- listed_at: datetime
- expires_at: datetime
- status: enum [active, sold, expired_returned]

Entity: MarketTransactionRecord (feeds Section 12.4 telemetry and 22.2.4 price history)
- transaction_id: UUID
- item_category_or_resource_type: string
- price_credits: int
- quantity: int
- source: enum [vendor_sale, bazaar_stall_sale, direct_trade]
- planet: string
- timestamp: datetime
```

### 22.4 System Interactions

- Indexes Section 21 (Vendor) listings galaxy-wide
- Feeds Section 12.4/12.6 economic telemetry and Section 34 market-health metrics
- Section 11.5.3 (Resource Trading) and Section 11.6 (Resource Market Dynamics) rely on Bazaar search/filter for resource-stat-based shopping specifically

### 22.5 Edge Cases

- **Seller goes offline/inactive with items in escrow and no sale:** Stall listings expire automatically (7 days) and unsold items are returned to the seller's mailbox (Section 18.2.2), never lost
- **Search query returning results across many planets:** results always display planet/city/distance so the search convenience never implies teleportation is available — buyer must still use Section 20 travel to physically collect a Vendor purchase (Stall purchases deliver directly, see 22.2.2)
- **Price manipulation via wash trading (self-buying to inflate transaction history):** flagged for server-side economic-anomaly detection (Section 29.6) but not preemptively blocked, since legitimate low-liquidity resource types can look similar; this is a monitoring, not a hard-block, edge case

### 22.6 Implementation Priorities

- **[MVP]** Bazaar Terminal search/filter across Vendor + Stall listings, Stall consignment selling with escrow, basic last-sold price display
- **[EXPANSION]** Full price-history charting, economic-anomaly detection dashboards

---

## 23. Entertainer System

### 23.1 Player Experience Goals

Entertaining is a complete, viable primary playstyle built around being present and performing for other people — not a passive buff dispenser. The Battle Fatigue healing loop (Pillar 1) should pull combat players physically into cantinas and hold them there for a social experience, not just a menu interaction. This is one of Pre-CU's most historically distinctive systems and must be implemented with real fidelity: audiences must actually watch, performers must actually perform.

### 23.2 Mechanics

#### 23.2.1 Starting & Maintaining a Performance

1. Entertainer selects a performance type (a learned Dance or Music routine, per their skill tree — Section 8.2.6/8.3.18/8.3.19) and enters "Perform" mode in a valid venue (Cantina, or any location if solo-busking at reduced effectiveness, **[ASSUMPTION]**)
2. Performance plays a looping animation/audio cue; the performer is stationary and vulnerable (cannot fight while performing) — this is a deliberate trade-off, not an oversight
3. Nearby players (within venue "watch radius") who target the performer and select "Watch" begin receiving the performance's effects on a tick basis
4. Performance continues until manually stopped, the performer runs out of Mind pool (performing costs Mind, matching the resource-driven combat philosophy of Section 9.1), or the performer is interrupted (combat, incapacitation)

#### 23.2.2 Battle Fatigue Healing & Buffs

- Watching a performance heals Battle Fatigue on a per-tick basis (Section 9.2.1 defines Battle Fatigue's HAM-pool-reduction mechanic); heal rate scales with the performer's relevant skill tier (Music/Dance Mastery) and performance quality roll
- Musicians and Dancers (Section 8.3.18/8.3.19) heal faster than a base Entertainer and unlock Mind-pool regeneration buffs applicable during the performance, stackable with Doctor buffs (Section 24) but not with another Entertainer's simultaneous buff of the same type
- Performance "flourishes" (skill-gated bonus animations) grant a temporary quality boost to the current performance tick rate, rewarding active play over passive idling even within the Entertainer's own performance loop

#### 23.2.3 Tips & Reputation

- Watching players may tip credits directly to the performer at any time; tipping grants the performer a small XP bonus (Section 7.2.3)
- No formal reputation score exists (Section 18.2.5) — a performer's draw is built through actual social reputation and location (a well-known Entertainer in a high-traffic city cantina earns more than an identical-skill Entertainer performing alone in the wilderness)

#### 23.2.4 Known Historical Edge Case: Passive "Watching"

**[HISTORICAL, flagged for design awareness]** In the original game, players could set "Watch" and then go AFK to passively heal Battle Fatigue, which the community both relied upon and criticized. This design preserves the mechanic as historically accurate (the watch/heal loop is not artificially blocked by anti-AFK detection), since removing it would be a philosophy change, not a bug fix — but idle-watching still requires physical presence at the venue and does not heal faster or slower than active engagement, so it creates no exploit relative to intended design, only a known play-pattern.

### 23.3 Data Entities

```text
Entity: PerformanceSession
- session_id: UUID
- performer_character_id: UUID
- performance_type_id: string (specific dance/song from skill tree)
- venue_structure_id: UUID (nullable, null = busking/no venue bonus)
- started_at: datetime
- ended_at: datetime (nullable, active if null)
- current_mind_cost_rate: int (per tick)

Entity: WatchSession
- watch_id: UUID
- watcher_character_id: UUID
- session_id: UUID
- started_watching_at: datetime
- battle_fatigue_healed_total: int
- buffs_received: [BuffInstance]

Entity: BuffInstance (shared with Section 24 Doctor buffs)
- buff_id: UUID
- target_character_id: UUID
- source_character_id: UUID
- buff_type: string
- ham_pool_modifiers: {health, action, mind}
- applied_at: datetime
- expires_at: datetime
- source_profession: enum [musician, dancer, doctor, squad_leader, food]
```

### 23.4 System Interactions

- Directly resolves the Section 9.2.1 Battle Fatigue mechanic — no other system heals Battle Fatigue
- Depends on Section 13/14 (Cantina structures, whether NPC-city or player-city-built)
- Interacts with Section 24 (Medic/Doctor) buff-stacking rules (same-type buffs from different sources do not stack; different-type buffs do)
- Section 8.3.20 (Image Designer) is a separate Entertainer elite path sharing the Master Entertainer prerequisite but no gameplay overlap with performance

### 23.5 Edge Cases

- **Performer incapacitated mid-performance (rare, environmental damage only, since performing is otherwise safe from direct attack in the sense that performers cannot be targeted while in NPC-city cantinas — Section 9.4.2 safe zones):** performance ends immediately for all watchers, no partial-tick heal is lost, effects already applied persist normally
- **Watcher leaves watch radius mid-tick:** watch session ends cleanly; no penalty, no exploit (can re-enter and resume watching)
- **Two Entertainers performing simultaneously in the same venue:** watchers can only actively watch one performer at a time (single Watch target), preventing double-dipping the BF-heal rate, though both performers still function and can each have their own audience
- **Performance in an active PvP zone:** allowed if the venue is not in a designated safe zone; the performer's stationary, defenseless state is a genuine tactical vulnerability, matching Section 9's tactical/preparation-over-twitch philosophy

### 23.6 Implementation Priorities

- **[MVP]** Perform/Watch loop, Battle Fatigue healing, tipping, basic buffs
- **[MVP]** Musician and Dancer elite paths (they are the primary reason to play Entertainer at all)
- **[EXPANSION]** Full flourish/quality-roll depth, Image Designer holoemote tooling

---

## 24. Medic System

### 24.1 Player Experience Goals

Where Entertainers heal the mind, Medics heal the body — the two professions are deliberately parallel in structure (venue-based, buff-granting, essential-but-non-combat) but distinct in mechanics (active tool-use and combat-adjacent play, versus performance). A Medic should feel indispensable to any group pushing into dangerous content, and a solo Doctor should be able to run a thriving practice out of a city medical center.

### 24.2 Mechanics

#### 24.2.1 Wound Healing

1. Injured player (carrying Wounds, Section 9.2.1) seeks out a Medic, or a grouped Combat Medic heals in the field
2. Medic selects a healing tool (stim pack, medical droid module, or bare-hands Medic ability at lower efficiency) and targets the wounded player
3. Healing action has a short cast/animation time (2–5 seconds, **[ASSUMPTION]**) during which a Combat Medic (but not a base Medic or Doctor) may remain mobile/combat-capable per their Field Triage tree (Section 8.3.17)
4. Wounds are reduced by a flat-plus-skill-scaled amount per application; repeated healing of the same player has a short cooldown/diminishing-returns window (matching Section 7.2.3's stated diminishing-returns design) to prevent trivial infinite spam-healing

#### 24.2.2 Incapacitation Revival

Full incapacitation state machine is defined in Section 9.6.1. This section specifies the Medic-side interaction: a Combat Medic (Section 8.3.17) targets an incapacitated player and performs a Revive action (cast time, Mind-pool cost); on success the target stands with 10% HAM pools and immediately begins carrying the Wounds/Battle Fatigue penalties from Section 9.6.1, which the reviving Medic (or another) can then treat.

#### 24.2.3 Doctor Buffs

- Doctors (Section 8.3.16) craft and apply Enhancement buff packs, granting large (+500 to +3000, per Section 6.2.2) temporary HAM pool increases essential for high-CL content (Pillar 2's "Buffs from Doctors/Musicians are essential for difficult content")
- Buffs are applied out of combat, require the target to stand still for a short application window, and stack with Entertainer/Chef buffs of different types but not with a second Doctor buff of the identical type (same rule as Section 23.3's BuffInstance)
- Buff duration is long (30–60+ minutes, **[ASSUMPTION]**) relative to a single content session, reinforcing the historical "buff before you go out" pre-combat ritual named in Section 4.1's Session Gameplay Loop

#### 24.2.4 Disease & Poison Cures

- Doctors and Combat Medics can cure the Disease and Poison DoT conditions defined in Section 9.2.6, in addition to natural time-based expiration
- Cure actions consume a crafted medical component (matching Pillar 1 — even healing itself has a crafted-goods dependency, since stim packs and cure kits are Medic-crafted consumables per Section 8.2.5)

### 24.3 Data Entities

```text
Entity: HealingAction
- action_id: UUID
- healer_character_id: UUID
- target_character_id: UUID
- action_type: enum [wound_heal, revive, buff_apply, cure_disease, cure_poison]
- tool_or_item_used: UUID (nullable, e.g., stim pack item reference)
- amount_healed: int (nullable, wound_heal only)
- performed_at: datetime
- cooldown_expires_at: datetime (per healer-target pair, diminishing returns)

Entity: MedicalCenter (extends CivicStructure, Section 13/14)
- structure_id: UUID
- city_id: UUID (nullable, null = NPC city)
- clone_facility_linked: bool
```

### 24.4 System Interactions

- Resolves Section 9.2.1 Wounds and Section 9.6.1 incapacitation/revive
- Buff mechanics share the BuffInstance entity with Section 23 (Entertainer)
- Depends on Section 8.2.5/8.3.16/8.3.17 (Medic, Doctor, Combat Medic) skill trees for effectiveness scaling
- Depends on Section 10.9 crafted stim packs/medical components (Medic Medicine Crafting tree, Section 8.2.5)

### 24.5 Edge Cases

- **Reviving a player whose incapacitation timer is about to expire:** revive must complete before the 5-minute timer (Section 9.6.1) elapses server-side; a revive action started just before expiry but completing just after fails and the target dies normally — no grace extension, to keep the timer authoritative and exploit-free
- **Multiple Medics attempting to heal the same wound simultaneously:** first successful action applies; subsequent concurrent actions recalculate against the now-reduced wound total rather than stacking full heals, preventing double-healing from a race condition
- **Buff applied, then target immediately logs out:** buff persists across the logout/login boundary using its absolute expiry timestamp, consistent with wall-clock (not play-time) expiry used elsewhere (Section 16.5)
- **Disease cure item used on a target with no active disease:** action fails gracefully (item not consumed, clear feedback), preventing wasted crafted consumables from misclicks

### 24.6 Implementation Priorities

- **[MVP]** Wound healing, incapacitation/revive, basic Doctor buffs, stim pack crafting/use
- **[EXPANSION]** Full disease/poison cure depth, Combat Medic mobile field-triage nuance beyond baseline healing

---

## 25. Ranger System

### 25.1 Player Experience Goals

The Ranger is the game's expert wilderness specialist: part tracker, part area-denial tactician, part traveling support unit. Playing a Ranger should feel meaningfully different from playing a Scout (its prerequisite) — less about finding resources and more about controlling territory and reading the world (creatures, players, terrain) better than anyone else.

### 25.2 Mechanics

#### 25.2.1 Tracking

- Rangers (Section 8.3.8) can track both creatures and, with sufficient skill, players — generating a directional trail/waypoint toward the tracked target similar in UI presentation to Section 11.3.1's survey waypoints, but based on footprint/scent simulation rather than resource geology
- Player tracking is a PvP-support tool: an Overt Ranger can track an Overt enemy-faction player who has recently passed through the area, supporting Section 15's factional gameplay and Section 8.3.21's Bounty Hunter loop without duplicating Bounty Hunter's dedicated Investigation tree

#### 25.2.2 Camps

- A Ranger (or base Scout, at reduced effectiveness) can deploy a temporary Camp: a small, timed structure (not a persistent Structure entity like housing) that provides a group buff radius (HAM regeneration boost, matching Pillar 5's "Scout require... Survey Tools" interdependence framing) for a limited duration (**[ASSUMPTION]** 30–60 minutes) before automatically despawning
- Camps are visually a campfire/tent cluster and cannot be placed inside NPC or player city limits (wilderness-only, reinforcing the distinction from permanent housing)

#### 25.2.3 Traps

- Rangers can place area-denial traps (snare, damage, or alert-type) at a target location; traps are triggered by proximity from non-owning, non-grouped characters or hostile creatures
- Traps have a limited number of charges/uses and a cooldown between placements per Ranger, preventing trap-field saturation of an area

#### 25.2.4 Camouflage

- A Ranger skill line grants reduced creature aggro radius and reduced player-tracking visibility while stationary or moving slowly in wilderness terrain, supporting both PvE scouting-ahead play and PvP ambush tactics

### 25.3 Data Entities

```text
Entity: TrackingAttempt
- attempt_id: UUID
- ranger_character_id: UUID
- target_type: enum [creature, player]
- target_reference_id: UUID
- success: bool
- waypoint_generated: UUID (nullable, references Waypoint, Section 18.3)

Entity: Camp
- camp_id: UUID
- deployed_by_character_id: UUID
- location: {planet, x, y, z}
- buff_radius_m: int
- buff_type: string
- deployed_at: datetime
- expires_at: datetime

Entity: Trap
- trap_id: UUID
- owner_character_id: UUID
- location: {planet, x, y, z}
- trap_type: enum [snare, damage, alert]
- charges_remaining: int
- placed_at: datetime
```

### 25.4 System Interactions

- Extends Section 8.3.8's Ranger profession summary with full operational detail
- Camps support Section 9.3 (Group Combat) as a pre-fight or mid-fight buff source, alongside Doctor/Musician buffs
- Player tracking supports Section 15 (Faction) PvP and Section 8.3.21 (Bounty Hunter) without being a strict duplicate
- Depends on Section 8.2.4 (Scout) as prerequisite, per Section 6.3.3's dependency chains

### 25.5 Edge Cases

- **Tracking a player who has since logged out:** tracking attempts against offline characters fail immediately with "trail has gone cold" feedback rather than resolving to a stale last-known location, preventing offline-player location leaks
- **Camp deployed just outside city radius, buffing players who then walk into the city:** buff persists per its normal duration once applied regardless of the buffed player's subsequent location, consistent with other buff-persistence rules (Section 24.5)
- **Trap triggered by the owner's own grouped ally:** grouped/owned characters never trigger a Ranger's own traps, avoiding friendly-fire frustration in group PvE content
- **Multiple Rangers layering camps in the same area:** camp buffs follow the same same-type-does-not-stack rule as other buffs (Section 23.3/24.2.3)

### 25.6 Implementation Priorities

- **[MVP]** Camps (direct group-support value, low implementation complexity)
- **[EXPANSION]** Player tracking, traps, camouflage (higher PvP-system complexity, lower priority than core PvE loops)

---

## 26. Surveying & Harvesting

### 26.1 Player Experience Goals

This section defines the moment-to-moment player workflow and profession-specific bonuses for finding and extracting resources. The underlying spawn algorithm and harvester structure mechanics are fully specified in Section 11.2–11.4; this section focuses on the tool progression, Scout/Ranger role, and the player-facing "hunt for the good spot" experience that makes surveying engaging rather than a chore.

### 26.2 Mechanics

#### 26.2.1 Tool Progression

| Tier | Source | Sample Yield | Survey Precision |
|---|---|---|---|
| Basic Survey/Sample Tool | Starting equipment, NPC vendor | 1 unit/sample | ±500m waypoint (Section 11.3.1) |
| Crafted Survey/Sample Tool | Artisan/Engineering schematic | 2–3 units/sample | ±250m waypoint |
| Master-tier Tool (Scout-only bonus applied) | Crafted + Scout skill modifier | 2–3 units/sample, faster animation | ±100m waypoint |

- Tool tier affects the base sample yield and animation speed; **Scout and Ranger skill tiers apply a multiplicative bonus on top of tool tier**, so a Master Scout with a basic tool still outperforms a Novice with a crafted tool, preserving skill investment as meaningful independent of gear (Pillar 2)

#### 26.2.2 Scout/Ranger Surveying Bonuses

- Scout's Hunting tree (Section 8.2.4) grants reduced survey cooldown and improved concentration-detection radius
- Ranger's Wilderness Survival tree (Section 8.3.8) further improves sample yield-per-action and grants a passive chance to detect nearby lairs/creature resources (linking Section 17.2.2's creature harvesting into the same surveying gameplay loop)

#### 26.2.3 Hot-Spot Competition Dynamics

- Because excellent-quality spawns are rare and time-limited (Section 11.2.2), the moment a high-OQ spawn is discovered, it becomes a competitive race to place a harvester (Section 11.4.1's 100m no-clustering rule directly creates this competition)
- This is intentional emergent sandbox gameplay (Pillar 4/Section 11.6.2) — the design must not soften this into a queue or reservation system, as that would remove the exploration-and-speed skill expression that makes surveying compelling beyond a menu action

#### 26.2.4 Player Workflow Summary

1. Equip category-appropriate survey tool
2. Survey → get waypoint + concentration estimate
3. Travel to waypoint, survey again to refine (triangulate)
4. Sample at refined location to confirm exact stats before committing a harvester
5. If stats justify investment, place harvester deed (Section 11.4.1) at the sampled spot
6. Return periodically to empty hopper and fund maintenance

### 26.3 Data Entities

Surveying and sampling reuse the core Section 11.2/11.3 Resource Spawn entities; this section adds only the tool-tier modifier data:

```text
Entity: SurveyTool (extends generic Item, Section 28)
- tool_id: UUID
- category: enum [mineral, chemical, flora, organic, water]
- tier: enum [basic, crafted, master]
- sample_yield_bonus: float
- survey_precision_bonus: float
```

### 26.4 System Interactions

- Directly extends Section 11.3 (Surveying Mechanics) and Section 11.4 (Harvester Mechanics)
- Scout (Section 8.2.4) and Ranger (Section 8.3.8) skill trees apply modifiers here
- Feeds the resource-acquisition side of Section 10.2's crafting core loop

### 26.5 Edge Cases

- Identical to Section 11's edge cases (concentration depletion, spawn despawn mid-harvest) — not duplicated here to avoid drift between the two sections; see Section 11.3.2 and 11.4.3

### 26.6 Implementation Priorities

- **[MVP]** Basic and Crafted tool tiers, Scout bonuses
- **[EXPANSION]** Master-tier tool bonuses, Ranger passive lair-detection bonus

---

## 27. Player Lifecycle

### 27.1 Player Experience Goals

A new player's first hour must establish the sandbox's core loop (Section 4.1) without resorting to the authored questing this design explicitly rejects (Section 1.2). A returning veteran after months away must be able to re-orient in a world whose resource spawns, market prices, and city landscape have all changed — this is a feature (Section 5.4's living world), not a problem to be smoothed away, but onboarding tooling should help both audiences find their footing quickly.

### 27.2 Mechanics

#### 27.2.1 Account & Character Creation

1. Account registration (out of scope for gameplay design, standard auth — see Section 29)
2. Character slots: **[ASSUMPTION]** 2–4 characters per account (later Pre-CU-era accounts commonly supported multiple characters, enabling the "alts for hybrid crafting" pattern referenced in Section 4.3's Economic Empire Building loop)
3. Character creation flow: Species → Gender → Appearance sliders → Name selection (server-unique, profanity/impersonation-filtered) → Starting City selection (any Section 5.2 MVP starter planet's NPC city)
4. Starting character spawns with: a small credit stipend, a basic weapon or tool appropriate to a suggested-but-not-mandatory starting path, and default clothing (Section 6.4.1)

#### 27.2.2 New Player Onboarding

- A brief, skippable orientation (NPC-guided, not a quest chain — no XP/item reward gate behind completing it) explains: the trainer/skill-box flow (Section 7.3), how to check the mission terminal (Section 16), and how to reach the cantina/bazaar
- The player's first meaningful choice is **which Basic Profession to start training** (Section 8.2) — the game does not force this choice before allowing movement/exploration, preserving sandbox freedom from minute one
- Starter missions (low-CL combat, simple delivery) exist specifically to bootstrap new-player credits without requiring veteran charity (cross-ref Section 12.8's edge case)

#### 27.2.3 Mid-Game: Specialization & Economic Participation

- Once a Basic Profession reaches Novice-complete, the player is prompted (non-mandatory) toward Elite Profession options with a brief in-fiction description of each (matching Section 8.3's descriptions), not a numeric power comparison, preserving Pillar 2's "versatility, not bigger numbers" framing
- Respec (Section 7.4) is surfaced clearly in the skill UI at all times, not hidden, since encouraged experimentation is core to the horizontal-progression philosophy

#### 27.2.4 Late Game: Mastery & Leadership

- Master-tier players are the primary suppliers of Section 4.3's Community Leadership loop (city government, guild leadership, mentorship per Section 19.2.5)
- No "max level" celebratory gate exists (no level 90 fanfare) — mastery is per-profession and ongoing, with the natural next step always being "master another profession" or "go deeper into economic/social/community systems," matching Pillar 2's rejection of a terminal endgame

#### 27.2.5 Session Boundaries: Logout & Safe Camping

- **[ASSUMPTION — exact historical implementation uncertain]** Logging out in a non-combat state (no recent damage taken/dealt, not incapacitated) is immediate. Logging out during or immediately after combat requires a short "camping" channel (holding a stationary logout timer, ~10-20 seconds, interruptible by taking damage) before the character safely leaves the world, preventing combat-logout abuse without introducing a punitive universal delay for ordinary logouts.
- Characters remain in the world (vulnerable to PvP if Overt in a contestable zone, or simply idle/AFK-visible in PvE) for a brief period after logout is initiated but before world removal, rather than vanishing instantly

#### 27.2.6 Returning Player Re-Onboarding

- On login after a long absence (**[ASSUMPTION]** 30+ days), a lightweight summary surfaces: current market price trends for the player's known crafting categories (Section 22.2.4), any structures that have entered a maintenance grace period (Section 12.8/13.2.5/21.2.2) needing attention, and guild/city status changes
- This is presented as an information panel, never a forced tutorial replay

### 27.3 Data Entities

```text
Entity: Account
- account_id: UUID
- character_slots_max: int
- created_at: datetime

Entity: Character (core identity record; extended by species/HAM/skills entities
  defined across Sections 6, 7, 28)
- character_id: UUID
- account_id: UUID
- name: string (server-unique)
- species: enum (Section 6.2.1)
- gender: enum
- starting_city: string
- created_at: datetime
- last_login_at: datetime
- logout_state: enum [logged_out, camping, in_world]
- camping_started_at: datetime (nullable)
```

### 27.4 System Interactions

- Character creation feeds directly into Section 6 (Character Architecture) and Section 7.3 (first skill training)
- Onboarding mission flow uses Section 16 (Mission System) starter missions
- Absence/return handling surfaces data from Section 12.8, 13.2.5, 21.2.2 (maintenance grace periods) and Section 22 (market data)

### 27.5 Edge Cases

- **Player disconnects (not graceful logout) mid-combat:** treated identically to a combat-state logout attempt — character remains in world, vulnerable, until either the client reconnects or a server-side timeout removes the character to a safe recovery state after an extended period (**[ASSUMPTION]** 10 minutes), preventing disconnect from being a combat-escape exploit while not punishing genuine connection loss indefinitely
- **All character slots full and player wants a new character:** must delete or the account must purchase additional slots (if that monetization existed historically) — **[ASSUMPTION]** for this non-commercial project, slot count is a fixed account attribute, not a purchasable one, consistent with Section 1.2's "no microtransactions"
- **Character name released after deletion:** name becomes available for reuse after a cooldown (**[ASSUMPTION]** 30 days) to prevent immediate impersonation of a just-deleted, possibly-known character

### 27.6 Implementation Priorities

- **[MVP]** Account/character creation, starting city selection, basic onboarding panel, starter missions, camping-based logout
- **[EXPANSION]** Returning-player re-onboarding summary panel (valuable but not blocking for a first playable build)

---

## 28. Data Model Specifications

### 28.1 Purpose

This section consolidates the authoritative entity schemas referenced throughout Sections 1–27 into a single reference an AI implementation agent can use to derive database migrations and API contracts directly. Where an entity was already fully defined in its owning section, it is referenced here rather than repeated verbatim; only entities not yet formally schematized, plus the top-level relationships between all of them, are defined in full below. **[IMPLEMENTATION RECOMMENDATION]** Field types are given in a language-agnostic pseudo-schema notation; map to the implementing agent's chosen ORM/database conventions.

### 28.2 Entity Relationship Overview

```text
Account 1──* Character
Character 1──1 HAMPoolState
Character 1──* SkillBoxOwnership ──* SkillBox ──* SkillTree ──1 Profession
Character 1──* Item (inventory)
Character 1──* Structure (as owner)
Character *──1 FactionStanding
Character *──0..1 Group
Character *──0..1 Guild (via GuildMembership)
Character *──0..1 City (via CityCitizen)
Structure ──┬─ House
            ├─ Harvester
            ├─ Factory
            ├─ Vendor
            └─ CivicStructure ──┬─ Cantina
                                 ├─ MedicalCenter
                                 ├─ CityHall
                                 └─ TravelNode
ResourceSpawn 1──* ResourceStack (owned by Character/Structure inventories)
Schematic *──1 Profession (crafting restriction)
Item *──0..1 Schematic (crafted_from reference)
CreatureTemplate 1──* CreatureInstance ──0..1 Lair
Mission *──1 MissionTerminal
BazaarStallListing *──1 Character (seller)
VendorListing *──1 Vendor
```

### 28.3 Core Identity Entities

```text
Entity: Item (base type for all equippable/tradeable objects)
- item_id: UUID
- item_category: enum [weapon, armor, clothing, structure_deed, resource_stack,
                        consumable, schematic, vehicle_deed, component, decoration,
                        instrument, medical_supply, container]
- display_name: string (player-assigned or schematic default)
- crafted_from_schematic_id: UUID (nullable; null for loot/NPC-sourced/starting items)
- crafter_character_id: UUID (nullable; preserves crafter attribution/reputation)
- condition_pct: float (0-100, decays with use; see Section 9.2.2 weapon decay,
                          10.9.1/10.9.2 armor/weapon experimentation attributes)
- stats: map<string, float> (attribute:value pairs, shape defined by item_category
                              and schematic — e.g., weapon.min_damage, armor.rating)
- owner_character_id: UUID (nullable; null if in a Structure's storage/vendor rather
                              than a personal inventory slot)
- container_location: enum [inventory, bank, house_storage, vendor_listing,
                             bazaar_escrow, equipped, mail_attachment]
- stack_quantity: int (1 for unique items; >1 for stackable resources/consumables)

Entity: Schematic
- schematic_id: UUID
- name: string
- profession_restriction: UUID (references Profession)
- complexity: int
- required_resources: [{resource_category, quantity, min_stat_requirements}]
- experimentation_attributes: [{attribute_name, points_available}]
- output_item_category: string
- source: enum [trainer, loot, quest_reward, reverse_engineered]

Entity: Profession (reference/lookup table, not per-character)
- profession_id: UUID
- name: string
- category: enum [basic, elite, hybrid]
- prerequisite_profession_ids: [UUID]
- prerequisite_master_count: int (for multi-prerequisite elites like TKA, Commando)
- skill_trees: [SkillTree]
- total_skill_point_cost: int

Entity: SkillTree
- tree_id: UUID
- profession_id: UUID
- name: string
- boxes: [SkillBox] (ordered, sequential prerequisite within tree)

Entity: SkillBox
- box_id: UUID
- tree_id: UUID
- tier: int (sequence within tree)
- skill_point_cost: int
- xp_cost: int
- xp_type: string
- credit_cost: int
- attribute_modifiers: map<string, float>
- granted_abilities: [string]

Entity: SkillBoxOwnership
- character_id: UUID
- box_id: UUID
- acquired_at: datetime
```

### 28.4 Character State Entities

```text
Entity: HAMPoolState
- character_id: UUID
- health_current: int
- health_max: int
- action_current: int
- action_max: int
- mind_current: int
- mind_max: int
- strength: int
- constitution: int
- quickness: int
- stamina: int
- focus: int
- willpower: int
- wounds: {health: int, action: int, mind: int}  (reduces respective max)
- battle_fatigue_pct: float (reduces all max pools)
- active_buffs: [BuffInstance]  (Section 23.3/24.2.3)
- posture: enum [standing, kneeling, prone, crouched]  (Section 9.2.3)
- stance: enum [normal, aggressive, defensive, berserk]
- incapacitated_at: datetime (nullable)
- last_combat_action_at: datetime (drives the logout-camping rule, Section 27.2.5)
```

*Character, Account, FactionStanding, Group, Guild, City, CityCitizen, Structure and its
subtypes, ResourceSpawn/ResourceStack, CreatureTemplate/CreatureInstance/Lair, Mission,
Vendor/VendorListing, BazaarStallListing, ChatMessage/MailMessage/Waypoint, and
CreditLedgerEntry are fully defined in their owning sections (27.3, 27.3, 15.3, 19.3,
19.3, 14.3, 14.3, 13.3, 11.2/10.4, 17.3, 16.3, 21.3, 22.3, 18.3, 12.6 respectively) and
are not repeated here to avoid definition drift between two copies of the same schema.*

### 28.5 Implementation Priorities

- **[MVP]** Item, Schematic, Profession, SkillTree, SkillBox, SkillBoxOwnership, HAMPoolState, and every entity marked MVP within its owning section
- **[IMPLEMENTATION RECOMMENDATION]** Given the transaction-integrity requirements of Section 12.8/21.5/22.5, Item ownership and CreditLedgerEntry writes should use the same relational database with ACID transaction support (Section 29.4) rather than an eventually-consistent store, even if other subsystems (chat, telemetry) use different storage.

---

## 29. Server Architecture Requirements

### 29.1 Player Experience Goals (Technical)

None of this section is player-visible directly, but every player-visible promise in this document — a harvester that's still running when you log back in, a bazaar search that's actually accurate, a PvP fight that can't be won by exploiting lag — is a server-architecture guarantee. The technical goal is a simulation that is authoritative, persistent, and honest.

### 29.2 Core Architecture Mechanics

#### 29.2.1 Authoritative Server Model

**[IMPLEMENTATION RECOMMENDATION]** All game state (position, HAM pools, inventory, structures, market) is authoritative on the server. Clients render predicted local movement for responsiveness but every combat roll, crafting result, and economic transaction is resolved server-side and pushed to clients — never client-computed and merely reported. This is *lower*-risk than a typical action MMO because Section 9.1 establishes combat as deterministic and stat-based rather than twitch-timing-based, meaning latency-sensitive client prediction is far less load-bearing here than in a reflex shooter.

#### 29.2.2 Spatial Partitioning

- World space is partitioned per-planet, then into a grid of zones/cells for interest-management (only sync entities near a player to that player's client)
- Structure interiors (Section 13.2.3) use a separate cell-graph addressing space from exterior world coordinates, consistent with the `cell_map_id` field on the House entity (Section 13.3)
- **[IMPLEMENTATION RECOMMENDATION]** Each planet can be run as an independently scalable process/shard, since cross-planet interaction is limited to travel-node transitions (Section 20) and galaxy-wide systems (Bazaar search, mail, guild/faction chat) that are naturally suited to a separate, planet-agnostic service layer

#### 29.2.3 World Tick Systems

Multiple independent tick rates are required, matching the cadences already specified throughout this document — this section is the consolidated technical reference:

| System | Tick Rate | Reference |
|---|---|---|
| Combat resolution | Real-time, per-action (sub-second) | Section 9.2.2 |
| HAM regeneration (out of combat) | ~1-5 second tick | Section 9.2.1 |
| Harvester extraction | 1 hour | Section 5.4.2, 11.4.2 |
| Lair respawn cycle | 12–24 hours | Section 9.5.3 |
| Resource spawn rotation | ~10 days (7-14 variance) | Section 11.2.1 |
| Structure maintenance deduction | 1 week | Section 12.3 |
| City treasury/rank evaluation | 1 week | Section 14.2.4 |
| Economic telemetry snapshot | 1 day (nightly) | Section 12.4 |
| Mission terminal refresh | 5–10 minutes | Section 16.2.2 |

**[IMPLEMENTATION RECOMMENDATION]** Implement as a scheduled job system (cron-like or a durable task queue) separate from the real-time game loop, so a slow economic-snapshot job can never stall combat tick processing.

#### 29.2.4 Client-Server Communication

- Real-time state sync (position, combat, chat) over a low-latency persistent connection (**[IMPLEMENTATION RECOMMENDATION]**: WebSocket or a UDP-based protocol depending on target platform constraints)
- Non-real-time operations (bazaar search, mail, skill training, structure management) are well-suited to conventional request/response API calls layered alongside the real-time connection

### 29.3 Data Entities

Server architecture does not introduce new gameplay entities; it defines operational infrastructure around the entities in Section 28:

```text
Entity: WorldTickJob (operational/infra record, not player-facing)
- job_id: UUID
- job_type: enum [harvester_extraction, lair_respawn, resource_rotation,
                   structure_maintenance, city_evaluation, economic_snapshot,
                   mission_refresh]
- scheduled_for: datetime
- executed_at: datetime (nullable)
- status: enum [pending, running, completed, failed]
- affected_entity_ids: [UUID]
```

### 29.4 Persistence Requirements

- **[IMPLEMENTATION RECOMMENDATION]** A relational database (ACID-compliant) for all economically or ownership-sensitive state: credits (Section 12.6), items (Section 28.3), structures (Section 13.3), bazaar/vendor listings (Sections 21–22). Exploitable duplication (Section 12.8) is a direct consequence of non-atomic writes to this data, so transactional integrity here is non-negotiable, not a nice-to-have.
- **[IMPLEMENTATION RECOMMENDATION]** A separate, higher-throughput/lower-durability store (in-memory or a fast key-value store) is appropriate for ephemeral real-time state: live positions, HAM pool ticking, chat (Section 18.3 notes chat is not persisted long-term except Tell/Mail)
- Regular backups/snapshots of the durable store, with point-in-time recovery, given the severity of any economic data loss to player trust in a persistent-world game

### 29.5 Anti-Exploit & Integrity Requirements

- Every transaction that moves credits or items (crafting output, trades, bazaar sales, mail attachments, maintenance payments) must be wrapped in a single atomic database transaction — partial application (e.g., credits deducted but item not granted) must be architecturally impossible, not merely "unlikely" (directly required by Sections 12.8, 21.5, 22.5)
- All combat, crafting, and experimentation rolls (Sections 9.2.2, 10.7.1) are computed server-side using server-seeded randomness; clients never submit a roll result, only an action request
- Rate limiting on high-frequency player-triggered systems (surveying, Section 11.3.1's 10-second cooldown; mission acceptance) is enforced server-side, not merely suggested client-side

### 29.6 Economic & Behavioral Monitoring

- **[IMPLEMENTATION RECOMMENDATION]** A lightweight anomaly-detection job (referenced in Section 22.5) flags statistically unusual transaction patterns (rapid self-trading, sudden price outliers, structure placement/destruction bursts) for human review, without automatically actioning accounts — this project has no dedicated trust & safety team, so tooling should surface signal, not auto-ban
- Logging/telemetry events should be structured to directly populate the Section 34 success-metrics dashboard (economic Gini coefficient, mission completion rates, profession population distribution, average session interdependence events)

### 29.7 System Interactions

- Every gameplay section in this document depends on the tick-rate table in 29.2.3 being correctly implemented for its stated cadence
- Section 12 (Economy) and Section 28 (Data Models) are the most architecturally load-bearing sections for the persistence-layer decisions here

### 29.8 Edge Cases

- **World tick job overlapping a player action mid-transition** (e.g., resource spawn rotation firing while a player is mid-harvest): the in-progress harvest completes against the pre-rotation spawn state; the harvester's *next* tick after rotation reflects the new spawn (or stops extracting if its bound spawn despawned, per Section 11.4.3)
- **Server crash/restart mid-tick:** WorldTickJob status tracking (29.3) ensures idempotent resumption — a job marked `running` but not `completed` at restart is safely re-evaluated rather than blindly re-run (which could double-apply harvester extraction, for example)

### 29.9 Implementation Priorities

- **[MVP]** Authoritative server loop, ACID persistence for economic/ownership data, all tick systems at [MVP]-tagged cadences, basic rate limiting
- **[EXPANSION]** Full anomaly-detection tooling, multi-shard planet scaling (a single-process world is adequate for an indie MVP's expected concurrency)

---

## 30. AI Agent Implementation Roadmap

### 30.1 Purpose

This section sequences the preceding 29 sections into a dependency-ordered build plan. Because nearly every system in this design depends on another (Pillar 1, made literal), build order is not arbitrary — attempting Entertainer (Section 23) before HAM pools (Section 9.2.1) exist, for example, is not possible. Each phase below lists its entry criteria (what must already exist) and exit criteria (what a working build can demonstrate) so an autonomous agent can self-verify progress without human sign-off at every step.

### 30.2 Build Phases

#### Phase 0 — Foundation
**Implements:** Section 29 (server skeleton, auth, persistence layer), Section 27.3 (Account/Character schema), Section 6 (species, HAM baseline, appearance)
**Entry criteria:** None (starting point)
**Exit criteria:** A player can create an account, create a character with species/appearance, and see it standing in a bare world with correct starting HAM values.

#### Phase 1 — Core World & Movement
**Implements:** Section 5 (planet/zone loading, at minimum one MVP planet), Section 20.2.1 (on-foot movement), Section 18.2.1 (spatial chat)
**Entry criteria:** Phase 0 complete
**Exit criteria:** A player can walk around a persistent zone on one planet and see/chat with another connected player in real time.

#### Phase 2 — Skills & Professions (Data Layer)
**Implements:** Section 7 (XP/skill-point system), Section 8.2 (Basic Profession trees as data), Section 28.3 (Profession/SkillTree/SkillBox schema)
**Entry criteria:** Phase 0 complete
**Exit criteria:** A player can earn XP from a placeholder action, spend skill points at a trainer NPC, and see the skill persist across logout/login.

#### Phase 3 — Combat Core
**Implements:** Section 9.2 (HAM-driven combat resolution, postures/stances), Section 9.5 (basic creature AI + lairs), Section 17.2.1–17.2.3 (creature templates/spawning)
**Entry criteria:** Phases 1 and 2 complete (needs movement + at least Marksman/Brawler skill trees)
**Exit criteria:** A player can fight and defeat a spawned creature using a skill-gated ability, take damage against the correct HAM pool, and be incapacitated/revived-or-cloned correctly (Section 9.6).

#### Phase 4 — Resources & Crafting
**Implements:** Section 11 (resource spawn algorithm, surveying, harvesters), Section 10 (schematics, experimentation, factories), Section 8.2.1/8.3.1–8.3.5 (Artisan + core crafting elites)
**Entry criteria:** Phase 3 complete (crafted gear needs the combat/item system to matter)
**Exit criteria:** A player can survey, harvest a resource, craft an item from a schematic with experimentation, and equip a crafted weapon/armor with correctly-computed stats.

#### Phase 5 — Economy Infrastructure
**Implements:** Section 12 (credit ledger, faucets/sinks), Section 21 (Vendor), Section 22 (Bazaar), Section 13 (Housing, since vendors are commonly house-attached)
**Entry criteria:** Phase 4 complete (needs crafted items to sell)
**Exit criteria:** A player can place a house, place a vendor, list a crafted item, and have a second player find and buy it via Bazaar search — with correct atomic credit/item transfer (Section 29.5).

#### Phase 6 — Social Support Professions
**Implements:** Section 23 (Entertainer), Section 24 (Medic/Doctor), Section 9.2.1's Battle Fatigue/Wounds full loop
**Entry criteria:** Phase 3 (combat generates BF/Wounds) and Phase 5 (venues are typically city/house-adjacent) complete
**Exit criteria:** A combat player can accumulate Battle Fatigue and Wounds through Phase 3's combat loop and have them healed by a second player performing Section 23/24's mechanics — demonstrating Pillar 1's core interdependence claim end-to-end.

#### Phase 7 — Civic Systems
**Implements:** Section 14 (City founding/governance), Section 19.2.4 (Guilds), Section 18.2.2 (Mail), full Section 13 permission/maintenance depth
**Entry criteria:** Phase 5 complete
**Exit criteria:** A group of players can found a city, elect a mayor, and see rank-gated unlocks (shuttleport) appear.

#### Phase 8 — Faction & Advanced PvP
**Implements:** Section 15 (Faction standing/ranks), Section 9.4 (PvP flagging/combat), Section 9.4.3 (base building, EXPANSION-tier depth optional)
**Entry criteria:** Phase 3 and Phase 7 complete
**Exit criteria:** Two Overt, opposing-faction players can engage in valid PvP combat with correct death penalties (Section 9.4.2), and faction points/ranks accrue correctly.

#### Phase 9 — Remaining Elite Professions & Missions
**Implements:** Remaining Section 8.3 elite professions not yet built (Bounty Hunter, Commando, Smuggler, Creature Handler, Bio-Engineer, Image Designer, TKA), Section 16 (Mission terminals), Section 25 (Ranger), Section 26 (Surveying tool-tier depth)
**Entry criteria:** Phases 3–8 complete (these professions build on combat, crafting, and faction systems already live)
**Exit criteria:** Full Section 8 profession roster is playable; mission terminals generate and reward correctly across all MVP-tagged mission types.

#### Phase 10 — Balance, Telemetry & Polish
**Implements:** Section 33 (balance target validation), Section 34 (success metrics instrumentation), Section 29.6 (monitoring)
**Entry criteria:** Phase 9 complete
**Exit criteria:** Section 1.3's five success-criteria tests can be run and pass (or produce actionable telemetry showing where they don't).

### 30.3 Cross-Cutting Implementation Notes

- Every phase should ship with the [MVP]-tagged subset of its sections first; [EXPANSION]-tagged features within a phase's sections can trail into a later pass without blocking phase completion (see Section 31 for the consolidated MVP list)
- Sections 28 (Data Models) and 29 (Server Architecture) are not a "Phase" of their own because they are cross-cutting — their content should be substantially designed before Phase 0 begins and revised incrementally as each phase's entity needs become concrete

### 30.4 Implementation Priorities

This roadmap *is* the implementation-priority mechanism for sequencing; Section 31 provides the complementary "what," this section provides the "in what order."

---

## 31. MVP Scope

### 31.1 Purpose

This section consolidates every **[MVP]** tag scattered through Sections 5–30 into a single checklist, so "is this playable yet" has one authoritative answer. MVP is defined as: *the smallest build that can pass all five Section 1.3 success-criteria tests*, not the smallest build that merely runs.

### 31.2 MVP World Scope

- **Planets:** Tatooine, Corellia, Naboo only (Section 5.2.1–5.2.3). All other planets (Talus, Rori, Lok, Dantooine, Dathomir, Yavin IV, Endor) are Phase-2 ground-game content — still fully in-scope for the *design*, deferred only for initial *build* effort. This is distinct from Section 32's true expansions.
- **Cities:** NPC starter cities on the three MVP planets, plus player-city founding up to Rank 3 (Section 14.2.2)
- **Zones:** Wilderness danger tiers (low/medium/high) on all three MVP planets

### 31.3 MVP Character & Progression Scope

- All 9 species (Section 6.2.1) — low implementation cost, high player-expression value, no reason to cut
- Full HAM/skill-point/respec system (Section 6.3, 7)
- All 6 Basic Professions (Section 8.2)
- **MVP Elite Professions:** Armorsmith, Weaponsmith, Architect, Merchant, Pistoleer, Rifleman, Carbineer, Fencer, Swordsman, Doctor, Combat Medic, Musician, Dancer, Creature Handler — the professions directly required to demonstrate Pillar 1's interdependence loop end-to-end (weapons/armor, structures, distribution, all core combat styles, both healing types, both social-buff types, basic taming)
- **Deferred Elite Professions [EXPANSION]:** Droid Engineer, Bio-Engineer, Ranger, Pikeman, Teras Kasi Artist, Image Designer, Bounty Hunter, Commando, Smuggler, Chef, Tailor, Politician (beyond its lightweight MVP perk, Section 14.2.3), Squad Leader

### 31.4 MVP Systems Scope

| System | MVP Scope |
|---|---|
| Combat (Section 9) | Full HAM/posture/stance/ability resolution, PvE, basic PvP flagging |
| Crafting (Section 10) | Full schematic/experimentation/factory loop |
| Resources (Section 11) | Full spawn algorithm, surveying, harvesters |
| Economy (Section 12) | Full ledger, primary faucets/sinks, basic telemetry |
| Housing (Section 13) | Small/Medium/Large houses, storage, maintenance |
| City (Section 14) | Founding through Rank 3, elections, tax |
| Faction (Section 15) | Alignment, Overt/Covert, points/ranks through Colonel |
| Missions (Section 16) | Combat Terminal, Crafting Delivery Terminal |
| Creatures (Section 17) | Full template/spawn/lair/harvest/basic taming |
| Social (Section 18) | All chat channels, mail, waypoints, friends |
| Grouping (Section 19) | Groups, loot rules, Guilds, Mentorship |
| Travel (Section 20) | On-foot, speeder bikes, shuttle/starport network |
| Vendor (Section 21) | Full placement/stocking/purchase/maintenance |
| Bazaar (Section 22) | Search/filter, Stall consignment, basic price history |
| Entertainer (Section 23) | Full Perform/Watch loop, Musician, Dancer |
| Medic (Section 24) | Wound healing, revive, Doctor buffs, stim packs |
| Ranger (Section 25) | Camps only |
| Surveying (Section 26) | Basic + Crafted tool tiers |
| Player Lifecycle (Section 27) | Full account/character creation, onboarding, camping logout |

### 31.5 Explicit Non-Goals for MVP (and Beyond — See Section 1.2)

Space (Jump to Lightspeed), Jedi, voice chat, quest chains, instancing, achievements/cosmetic unlocks, and microtransactions are excluded at every phase, not just MVP — see Section 32 for how the architecture should still accommodate them later without a rewrite.

### 31.6 MVP Definition of Done

MVP is complete when Phase 9 of Section 30's roadmap is reached and Section 1.3's five tests can be actively run against the live build (instrumented via Section 34).

---

## 32. Future Expansion Hooks

### 32.1 Purpose

The systems explicitly excluded in Section 1.2 are excluded from *this build*, not from the *architecture's imagination*. This section exists so the AI implementation agent does not make Phase 0–10 decisions that would require a rewrite to later add these systems — it should leave the door open without walking through it.

### 32.2 Space Gameplay (Jump to Lightspeed)

- **Hook:** Section 20's TravelNode entity already models Starport-to-Starport travel as a discrete, non-simulated transition. A future space layer would insert a simulated flight phase between planetary Starport departure and destination arrival, without needing to change the ground-side TravelNode contract.
- **Do not build:** ship combat, ship crafting/schematics, space stations, space-specific professions (Pilot).

### 32.3 Jedi / Force-Sensitive Progression

- **Hook:** Section 5.2.7 (Dantooine) and Section 5.2.5 (Rori) are already flagged as historically Jedi-adjacent planets with their village/enclave content explicitly omitted. A future unlock system (historically, a hidden, randomized "Holocron" village-visit chain) could be layered onto the existing Section 7 skill-point system as simply another Profession entity (Section 28.3) with an unusual, hidden unlock condition rather than a normal trainer purchase — the Profession/SkillTree/SkillBox schema does not need to change.
- **Do not build:** Jedi combat trees, Force powers, lightsaber crafting, or any Jedi-visible UI in MVP.

### 32.4 Quest-Driven Narrative Content

- Deliberately excluded from design philosophy (Section 2.4, Pillar 4), not merely unbuilt. If ever added, it should be modeled as strictly optional, clearly-labeled flavor content layered on top of the Mission Terminal system (Section 16), never replacing procedural missions as the primary loop.

### 32.5 Instanced Content

- Deliberately excluded (Section 3.2). The Lair (Section 17.3) and FactionBase (Section 15.3) entities are intentionally shared-world objects with no instance-id concept. Adding instancing later would require a genuinely new spatial model, not a hook — this is flagged as a hard architectural boundary, not a soft one.

### 32.6 Achievement Systems & Cosmetic Unlocks

- Section 18.2.5 already explains why no reputation-score system exists. Any future achievement layer should be additive telemetry (reading Section 29.6/34 event streams) rather than a new gameplay-gating system, to avoid reintroducing the vertical-unlock philosophy Pillar 2 rejects.

### 32.7 Other Natural Post-Launch Additions (Non-Canonical Exclusions, Just Deferred)

- **Additional planets** (Section 5.2.4–5.2.10) — architecturally identical to MVP planets, pure content addition
- **City Specialization** (Section 14.2.3) and full **Politician** toolkit
- **Squad Leader** profession (Section 19.2.3)
- **Bio-Engineer** DNA/enhanced-pet depth (Section 17.2.4)
- **Veteran rewards** — **[ASSUMPTION]** not part of the documented Pre-CU design pillars; if added, must be cosmetic/flavor only per Pillar 2, never a power increase
- **Server-side economic dashboards for players** (Section 22.2.4's full price-history charting)

### 32.8 Implementation Priorities

All content in this section is **[EXPANSION]** or explicitly out-of-scope per Section 1.2; nothing here should consume Phase 0–10 (Section 30) effort. Its only implementation obligation on the current build is: do not close doors that a hook in this section would need open (e.g., do not hardcode "exactly 3 planets" anywhere the planet count should be a config value).

---

## 33. Game Balance Targets

### 33.1 Purpose

This section consolidates every numeric balance target given throughout the document into one reference table, so playtesting and tuning have a single source of truth. Values already fully specified in their owning section are referenced, not restated in full; only cross-system targets not yet consolidated anywhere are added new here.

### 33.2 Combat Balance (full detail: Section 9.7)

Damage scaling (50–150 starting, 300–700 end-game, 1000–1500 max crit), HAM pool scaling (3,000–12,000 Health depending on buffs), Time-to-Kill (20–120 seconds depending on content), ability cooldowns and HAM costs — all fully specified in Section 9.7 and not repeated here.

### 33.3 Progression Time Investment (full detail: Section 4.3, 8.2–8.3)

| Milestone | Target Time |
|---|---|
| Master a Basic Profession | 40–60 hours (Section 4.3) |
| Master an Elite Profession | 80–120 hours (Section 4.3, matches ~80 skill-point professions in Section 8.3) |
| Master a multi-prerequisite Elite (Bounty Hunter, Commando, TKA) | 150–250 hours (120-point cost, Section 8.3.21/22, plus prerequisite mastery) |
| Fill the full 250-point skill cap (2 Elites + utility) | 200–300 hours cumulative |

### 33.4 Economic Balance (full detail: Section 11.6.1, 12)

Resource pricing bands (1–1000+ credits/unit by OQ tier, Section 11.6.1), target inflation (5–10% annually, Section 12.4). New cross-system targets:

| Metric | Target |
|---|---|
| Novice crafter hourly income (casual play) | Roughly comparable to a Novice combat player's mission income — neither role should be a strictly dominant strategy for a new character (Pillar 2 applied to economics) |
| Master crafter hourly income (focused production) | 2–3x a Novice's, matching the Section 2.2 "2-3x effectiveness, not 50x" progression band applied to economic power |
| Structure maintenance as % of owner's typical weekly income | 5–15%, enough to matter, not enough to be oppressive (Section 12.3's "never the dominant sink" rule) |
| Bazaar Stall commission | 4–6% (Section 22.2.3) |

### 33.5 Housing & Civic Cost Targets

| Item | Target Cost Band |
|---|---|
| Small House deed | Affordable to a Novice crafter within a few play sessions |
| Medium/Large House deed | Requires either a Master crafter's output or several sessions of saving |
| Harvester deed (any tier) | Comparable to a Medium House; the real cost is maintenance, not acquisition |
| City founding (City Hall deed) | Expensive enough to require guild/community pooling, not a solo weekend project |

### 33.6 Faction Rank Time Investment (full detail: Section 15.2.3)

Faction point thresholds (2,500 / 10,000 / 30,000 / 75,000 cumulative) are specified in Section 15.2.3; at an estimated 50–150 points per meaningful PvP engagement, Sergeant is reachable in single-digit hours of active PvP, General requires sustained, likely multi-week commitment plus sponsorship — intentionally the slowest-earned rank in the entire progression system, reflecting its status as a social/leadership capstone rather than a solo grind.

### 33.7 Implementation Priorities

- **[MVP]** Combat and progression-time targets (Sections 9.7, 33.3) must be validated first, since they gate whether Section 1.3's Horizontal Progression Test can even be evaluated
- **[MVP]** Economic balance targets (33.4) directly gate the Player-Driven Economy Test
- **[EXPANSION]** Fine-tuned housing/civic/faction cost bands can be iterated post-MVP using the Section 29.6 telemetry

---

## 34. Success Metrics

### 34.1 Purpose

Section 1.3 defines five pass/fail tests. This section converts them into ongoing, measurable KPIs — the instrumentation an AI agent (or human team) should build and monitor continuously, both during development and after the game is live, rather than checking once and moving on.

### 34.2 Metrics Mapped to Section 1.3 Tests

| Section 1.3 Test | Measurable KPI | Target | Data Source |
|---|---|---|---|
| Economic Interdependence Test | % of active combat players who received a crafted item, an Entertainer heal, or a Medic heal within the last 7 days | ≥ 90% | Section 29.6 telemetry across Sections 9, 23, 24 |
| Horizontal Progression Test | Ratio of a 6-month character's average combat effectiveness (damage output, survivability) to a 2-week character's, in equivalent gear tier | 2–3x (matches Section 2.2/9.7) | Combat telemetry |
| Player-Driven Economy Test | % of items in active player inventories/vendors that are `crafted_from_schematic_id != null` (Section 28.3) vs. NPC/loot-sourced | ≥ 95% | Item entity query |
| Sandbox Validation Test | Qualitative: presence of player-organized events, guild structures, and economic schemes not authored by the design (tracked via community/GM observation, not automatable) | Ongoing positive signal, not a single number | Community observation, chat/event logs |
| Historical Accuracy Test | Structured playtest comparisons of core loops (Section 4) against documented Pre-CU community accounts | Design review checklist, not telemetry | Human validation pass |

### 34.3 Supporting Economic Health Metrics (full detail: Section 12.4, 12.6)

- **Net credit inflation (30-day rolling):** target 5–10% annualized (Section 12.4)
- **Wealth Gini coefficient:** tracked, not target-capped — Section 11.6.2 notes cartel/wealth-concentration behavior is historically accurate emergent gameplay, so this metric is diagnostic (are we seeing healthy economic differentiation or runaway monopolization) rather than a hard pass/fail gate
- **Resource price volatility:** average % price swing per resource category per spawn-rotation cycle (10 days, Section 11.2.1) — should be nonzero (proves the spawn algorithm is actually creating scarcity-driven price movement) but not so extreme that new players cannot predict basic costs

### 34.4 Supporting Engagement & Social Metrics

- **Profession population distribution:** % of active characters with at least one Elite Profession box in each category (combat / crafting / social-support); a healthy sandbox should not see any category collapse toward zero, since Pillar 1 requires all three to be populated for interdependence to function
- **Average "interdependence events" per active player per week:** count of times a player received a service from another player's profession (crafted item purchase, Entertainer heal, Medic heal, harvested resource purchase) — this is the single most direct proxy for whether Pillar 1 is being lived out in practice, not just theoretically possible
- **City survival rate:** % of founded cities still at Rank 1+ after 90 days — a proxy for whether Section 14's governance/maintenance balance is sustainable rather than causing community projects to collapse

### 34.5 Playtesting Protocol Recommendations

- **[IMPLEMENTATION RECOMMENDATION]** Run structured playtests with cohorts of new (Section 27.2.2) and veteran (Section 27.2.4) players simultaneously to directly observe the Horizontal Progression Test in practice, not just measure it after the fact
- **[IMPLEMENTATION RECOMMENDATION]** Seed at least one full economic cycle (a 10-day resource rotation, Section 11.2.1) before evaluating economic KPIs, since Day 1 of a fresh server has no meaningful price history

### 34.6 Implementation Priorities

- **[MVP]** All five Section 1.3 test proxies must be instrumented before Phase 10 (Section 30.2) is considered complete
- **[MVP]** Economic health metrics (34.3), since they gate the Player-Driven Economy Test
- **[EXPANSION]** Full engagement/social metrics dashboard (34.4) can mature iteratively post-launch

---

## Closing Note

This document is intended to be read start-to-finish once by an implementing agent, then used as a searchable reference throughout the build described in Section 30. Every numbered cross-reference in this text (e.g., "Section 9.2.1") is a literal pointer for that purpose. Where this document makes an **[ASSUMPTION]**, that ruling is binding for this project regardless of whether it matches every player's personal memory of the original game — the goal, per Section 1.1, is a coherent, buildable, faithful-in-spirit sandbox, not a forensic reconstruction of every undocumented server-side detail of a game that shipped in 2003.

*End of document.*
