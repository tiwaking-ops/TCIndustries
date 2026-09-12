# Tiwakings Craftworld Industries (TCIndustries)
## Master Game Design Document — First-Pass

**Version:** 0.1
**Date:** August 20, 2026
**Document type:** Game / rules specification. This is **not** a Technical Design Document — it intentionally contains no engine, networking, database, or UI implementation detail. That belongs in a future TDD once this GDD is stable.
**Authority:** This document is the canonical design reference. Where the existing TypeScript/Vite/React prototype disagrees with this document, this document wins — the prototype is evidence, not law.
**Author:** ChatGPT (per filename, unverified; content byte-identical to TCIndustries_Master_GDD-claude1.md — single authorship, duplicate filing)

---

### How to read this document

Every substantive design element below carries one of five status tags:

| Tag | Meaning |
|---|---|
| **[LOCKED]** | An explicit project decision. Treat as authoritative until formally revisited. |
| **[PROPOSED]** | A design recommendation that appears sound but has not been formally approved. |
| **[TBD]** | A genuine unresolved decision requiring future discussion. |
| **[PROTOTYPE]** | Behaviour explored in prototype code but not yet accepted as canonical. |
| **[DEFERRED]** | A known future system intentionally left undesigned for now, so it doesn't block current work. |

Nothing in this document silently promotes a [PROPOSED], [TBD], [PROTOTYPE], or [DEFERRED] item to [LOCKED]. Items are marked [LOCKED] only where the generation brief stated them as fixed decisions, or where noted as a direct restatement of one. Almost everything else — every taxonomy, formula shape, and structural rule invented in this pass — is [PROPOSED]: it is this designer's recommendation, not yet your approved design.

**On [PROTOTYPE]:** no prototype source was available in this session (checked — nothing was uploaded alongside the brief). Since no prototype behaviour was actually inspected, nothing in this pass is marked [PROTOTYPE]. This is logged as OQ-01 and AS-01 below. When the prototype is supplied, it should get its own reconciliation pass rather than being trusted by default.

Each major rule is tagged with a short ID (e.g. `RES-003`) so future conversations can reference it directly. All IDs are compiled in the Section 27 Status Register.

---

## Table of Contents

1. [Executive Summary](#1-executive-summary)
2. [Vision](#2-vision)
3. [Design Pillars](#3-design-pillars)
4. [Player Experience](#4-player-experience)
5. [World Philosophy](#5-world-philosophy)
6. [World Simulation](#6-world-simulation)
7. [Character and Skill System](#7-character-and-skill-system)
8. [Resource System](#8-resource-system)
9. [Crafting System](#9-crafting-system)
10. [Manufacturing](#10-manufacturing)
11. [Economy](#11-economy)
12. [Retail](#12-retail)
13. [Professions](#13-professions)
14. [Social Systems](#14-social-systems)
15. [Services](#15-services)
16. [Combat](#16-combat)
17. [Items](#17-items)
18. [Buildings](#18-buildings)
19. [Cities](#19-cities)
20. [Progression](#20-progression)
21. [Emergent Gameplay](#21-emergent-gameplay)
22. [Content Expansion](#22-content-expansion)
23. [Anti-Goals](#23-anti-goals)
24. [Balance Philosophy](#24-balance-philosophy)
25. [Exploit / Fraud Considerations](#25-exploit--fraud-considerations)
26. [Testing Requirements](#26-testing-requirements)
27. [Status Register](#27-status-register)
28. [Open Questions](#28-open-questions)
29. [Assumption Register](#29-assumption-register)
30. [Systems Requiring Deeper Design](#30-systems-requiring-deeper-design)
31. [Recommended Design Sequence](#31-recommended-design-sequence)
32. [Glossary](#32-glossary)

---

## 1. Executive Summary

TCIndustries is a persistent-world MMORPG built as a spiritual successor to the *virtual-world philosophy* of early Star Wars Galaxies — its resource economy, crafting depth, profession interdependence, and player-driven society — under entirely original names, systems, and content. It is not a reproduction of that game's intellectual property, and nothing in this document borrows Star Wars-specific lore, characters, or proper nouns.

The design goal is not "an MMO with a crafting system bolted on." It is a world where a player's occupation, reputation, relationships, and possessions *are* the game. Combat is one lifestyle among many available ones, not the default one, and no single system — least of all automation — is allowed to make other players optional to a fulfilling economic or social life.

This first-pass GDD establishes:

- The philosophical foundation (four pillars, anti-goals, sandbox commitment) — mostly **[LOCKED]**, because the brief specified these directly.
- The taxonomies and abstractions each major system will run on (resource attributes, schematics, skill disciplines, structure categories, profession families) — mostly **[PROPOSED]**, because they are this pass's recommendation, not yet approved canon.
- Explicit mechanical guardrails against the specific failure modes the brief named (Factorio-drift, EVE-spreadsheet-drift, theme-park-drift, single-player industrial empires).
- A Status Register, Open Question Register, and Assumption Register, so nothing here is quietly treated as settled that isn't.
- A recommended sequence of design passes to follow this one.

**What this document deliberately does not do:** invent exact numbers. Skill point costs, spawn rates, prices, decay curves, and formula constants are marked **[TBD]** throughout rather than guessed at, because false precision this early would look like a decision and get built against as if it were one. Section 24 explains this reasoning in more depth. This document defines *shapes* (what inputs a formula takes, what it must guarantee, what would break if it were wrong) so that a later numeric-tuning pass has something correct to tune.

---

## 2. Vision

### 2.1 Game Vision — `VIS-001` **[LOCKED]**

> Create a living virtual world in which players develop their own identity, occupation, reputation, relationships, businesses, and place in society.

This is the brief's own core design goal, restated verbatim as the document's north star. Every system below should be traceable back to it.

### 2.2 Player Fantasy — `VIS-002` **[LOCKED]**

The player is **not**: the chosen hero, the saviour of the world, the protagonist of a fixed story, or a unit in an economic spreadsheet.

The player **is**: a citizen. Someone who wakes up in this world and has to decide, largely for themselves, what they're going to be — and who then gets to *become* known for being that.

### 2.3 Target Experience — `VIS-003` **[PROPOSED]**

A representative session should be recognisable as "a day in the life," not "the next quest in the chain." For example: a shop owner logs in, checks overnight vendor sales, restocks from a harvester that finished its cycle, fills two standing commissions for regular customers, hears a rumour about an exceptional resource spawn two zones over and decides whether to chase it or stay put, and spends twenty minutes actually crafting something using saved-up high-quality resources because a customer specifically asked for their best work. None of that requires combat, a quest-giver, or a level-up screen. A combat-focused session should look structurally similar but mirrored: checking gear condition, visiting an armourer and a medic before heading out, and coming back to sell salvage to specialists rather than a generic vendor.

### 2.4 Anti-Goals (summary)

TCIndustries should not gradually become Factorio, Satisfactory, a Capitalism-style business sim, EVE-style spreadsheet gameplay, a conventional theme-park MMO, a quest-driven RPG, a conventional class-based MMO, or a single-player crafting game. Full list and the specific mechanical guardrails against each drift risk are in Section 23.

---

## 3. Design Pillars

These four pillars are **[LOCKED]** as stated in the brief. What follows under each is this pass's **[PROPOSED]** translation of the pillar into mechanics — the pillar itself is fixed, the specific mechanics serving it are not.

### 3.1 Discovery and the Gold Rush — `VIS-004`

**Design goal:** resource discovery should feel like an event, not a database lookup.

Mechanically, this pillar is served primarily by the Resource System (Section 8): resources are not static infinite commodities but time-limited **Resource Pools** with rolled quality, discoverable through a skill-gated Surveying activity, rare in their best rolls, and finite in duration. An exceptional spawn should be the kind of thing a player tells other players about — either because they're recruiting help to exploit it fast, or because they're trying to keep it secret long enough to corner it. Both reactions are the pillar working correctly.

### 3.2 Identity and Reputation — `VIS-005`

**Design goal:** players should seek out a *particular maker*, not "a product."

Mechanically served primarily by Crafting's provenance system (Section 9.5) — every crafted item permanently records who made it — and by Retail's shop branding (Section 12). Reputation should be something that accrues to a name over time through visible, unforgeable signals (what you made, how much of it, how well), not a self-reported title.

### 3.3 Interdependence — `VIS-006`

**Design goal:** professions should need each other, not merely coexist.

This is the pillar most systems above are load-bearing for, and it is also the pillar most vulnerable to being quietly eroded by any single "convenient" feature (an NPC vendor that's just as good as a player crafter, a solo-craftable top-tier item, a factory that never needs a skilled operator). Every system section below includes an explicit **Interdependence Check** — the question "does this let a player skip needing another player?" — because that check is this pillar's actual enforcement mechanism, not a slogan.

### 3.4 Player-Created Society — `VIS-007`

**Design goal:** players should be able to build structures — literal and social — that outlast a single session.

Mechanically served primarily by Buildings, Cities, and Social Systems (Sections 14, 18, 19): shops, workshops, player associations, and player-founded cities with their own governance and identity. The test for this pillar is whether the *place* becomes meaningful — whether "the forge district in [player city]" becomes a phrase players actually use.

---

## 4. Player Experience — **[PROPOSED]**

### 4.1 Onboarding intent

A new player should be able to answer "what am I doing right now" within their first session without having chosen combat. The very first hour should expose, at minimum: one gathering activity, one crafting attempt, and one moment of trading with another real player (even if it's just a starter item). Exact tutorial structure is **[DEFERRED]** — see Section 30 — but this ordering constraint (economy-and-people before combat-and-power) is proposed as a hard requirement on whatever onboarding is eventually built, because an onboarding flow that teaches combat-first will silently make combat the game's default fantasy regardless of what this document says.

### 4.2 No forced path — `VIS-008` **[LOCKED]**

Restates Section 5 of the brief directly: the game must not require a player to become a combatant, and a player must be able to pursue a full, meaningful life through non-combat activity alone. This is a hard constraint on every other system: if any later system can only be reasonably engaged with by fighting things, that system violates this rule.

### 4.3 Respecialization as safety net

Because the game asks players to choose an identity early, it needs a credible, non-punitive (but not free) way to change that choice later without abandoning the character. See `CHR-004` in Section 7.


---

## 5. World Philosophy — **[PROPOSED]**

The world exists to make Section 3's four pillars physically real. Concretely, that means the world must have:

- **Uneven geography.** If every zone has equivalent resources and equivalent settlement access, there is no reason to travel, trade between regions, or specialise regionally. Regional inequality is a feature, not a gap to be patched.
- **Room for players to build in it.** A world philosophy that treats the map as finished, authored content is incompatible with Pillar 3.4 (Player-Created Society). Meaningful stretches of land must be claimable and buildable.
- **A reason to keep discovering it.** Because resources migrate (Section 8), the "interesting" parts of the map should shift over a long timescale, so veteran players don't end up with a fully solved mental map on day 30.

`WLD-001` **[PROPOSED]** — World structure is hierarchical: **Region** (a large, thematically or geographically coherent area) → **Zone** (a subdivision of a Region, the unit at which resource pools and points of interest actually exist) → individual **Resource Pools**, **Points of Interest**, and **claimable land parcels** within a Zone. This hierarchy exists purely as an organisational abstraction for content and expansion (Section 22), not as a hard technical statement about how the world is streamed or rendered — that belongs in a TDD.

### Interdependence Check
A world philosophy that makes every zone self-sufficient defeats Pillar 3.3 before any other system gets a chance to. Regional variation in Section 8's resource spawning is the mechanism that prevents this; it is treated as a requirement of the world philosophy, not just a resource-system nicety.

---

## 6. World Simulation — **[PROPOSED]**

### 6.1 Settlements — `WLD-002`

Two settlement types exist:

- **Established Settlements**: world-authored, always present, never destructible or foundable by players. They provide baseline civic services every player needs regardless of profession or reputation — training access, banking, a transport hub, and a baseline market presence. They exist so a brand-new player, or a player far from any player city, is never fully cut off from the game's core loops.
- **Player Cities**: player-founded, grow from clusters of player-owned structures, and are the primary mechanical expression of Pillar 3.4. Full treatment in Section 19.

Established Settlements must deliberately **not** out-compete player commerce — their vendors sell only baseline, non-exceptional goods (see `ECN-002` in Section 11) so they act as a floor, not a ceiling, on the player economy.

### 6.2 Wilderness — `WLD-003`

Wilderness is unclaimed land where resource extraction and exploration happen. It should carry a **low density of world-authored content** — enough to make travel feel purposeful (points of interest, environmental hazards, opportunities) — without becoming a quest-delivery system. Wilderness's primary job is to host Resource Pools and available land for future player structures, not to tell a story.

### 6.3 Time — `WLD-004` **[TBD]**

A single persistent world clock is assumed, since resource spawn cycles (Section 8) need a shared clock to key off. Whether the world has a visible day/night cycle, and whether day/night affects anything mechanically, is genuinely undecided and left **[TBD]** — it's a presentation and pacing question best answered once the resource spawn cadence (also TBD) is defined, since the two would likely want to relate to each other.

### 6.4 Weather — `WLD-005` **[DEFERRED]**

The brief itself lists weather as conditional ("weather if appropriate"). Deferred entirely for this pass. Worth flagging one integration hook for later: weather is a very natural way to add temporal variation to certain Resource Domains (Section 8.1) without new systems — e.g., an Aquatic resource pool behaving differently during a storm — but this is a future idea, not a current commitment.

### 6.5 Transportation — `WLD-006`

`WLD-006` **[PROPOSED]** — Transportation is defined at the abstraction level only: Zones and Settlements are connected by transportation nodes; both world-authored transport (baseline routes between Established Settlements, so no player is ever stranded) and player-operable transport exist. Critically, transportation should be designable as a **player service and business**, not merely a menu-driven fast-travel button — see `SVC-001` in Section 15. Specific vehicle types, mounts, or vehicle crafting are genre-dependent choices and are marked **[TBD]** pending a setting decision (see `AS-02` in the Assumption Register — this GDD is deliberately genre/setting-neutral).

---

## 7. Character and Skill System — **[PROPOSED]**

### 7.1 Attribute pools — `CHR-001`

Characters have a small set of resource pools that gate sustained activity across *both* combat and non-combat play, so that "getting tired" is a shared mechanic rather than something combat has and crafting doesn't. Proposed pools: **Vitality** (physical health), **Stamina** (sustained physical exertion — harvesting, construction, sustained combat), and **Focus** (mental exertion — crafting experimentation, entertaining, research-type activity). Tying crafting and entertaining to a depletable pool is a deliberate, direct mechanical hook for Pillar 3.3: a crafter who runs out of Focus has a real reason to want a service from a Medical or Entertainment specialist, the same way a fighter running low on Vitality does. Exact pool sizes, regeneration rates, and formulas are **[TBD]**.

### 7.2 Skill Disciplines and Skill Boxes — `CHR-002`

Historical SWG behaviour, adapted: rather than classes, characters invest in **Skill Disciplines** (e.g., a crafting discipline, a combat discipline, a medical discipline), each broken into tiers of **Skill Boxes** that grant abilities, passive bonuses, or unlock specific schematics/recipes. Boxes are purchased with **Skill Points**. This structure is the direct mechanical foundation of Section 5's sandbox requirement — it replaces "pick a class" with "spend a shared budget across whatever boxes you want," which is what makes hybrid characters (part-crafter, part-medic) possible at all.

### 7.3 Skill Point Cap — `CHR-003`

A capped Skill Point budget is assumed to exist — otherwise every character eventually masters everything and Pillar 3.3 collapses once the playerbase is sufficiently veteran. The cap should be generous enough to support roughly two to three concurrent, meaningfully-developed professions, but not so generous that mastering a large fraction of all disciplines is realistic. Exact totals and per-box costs are **[TBD]** — see Section 20 for the philosophy this cap needs to satisfy, and Section 26 for a testable invariant derived from it.

### 7.4 Skill acquisition — `CHR-005`

Skill Points are earned primarily through **use-based experience** ("learn by doing" — historical SWG behaviour, adapted) rather than a generic level-up pool, so that a character's earned points reflect what they actually did. Whether a secondary, smaller pool of points is awarded for other activity (exploration, quests-if-any, social participation) is **[TBD]**.

### 7.5 Respecialization — `CHR-004`

Skill Boxes can be unlearned through a **Respecialization** action, returning the spent Skill Points for reallocation. This exists specifically to satisfy Section 4.3's requirement for a safety net against a badly-chosen initial identity. It must have real friction — cost in time, currency, and/or a consumable item — so that specialization still feels like a meaningful choice and identity doesn't become trivially fluid; a same-session, free respec would undercut Pillar 3.2 (Identity and Reputation) by making "being a blacksmith" reversible on a whim. Exact costs are **[TBD]**.

### Interdependence Check
The single biggest risk in this section is `CHR-003`: if the Skill Point cap is set too high, nothing above matters — a sufficiently patient player becomes self-sufficient and Pillar 3.3 quietly dies. This is flagged again in Section 24 (Balance Philosophy) and Section 26 (Testing Requirements) as the one number in this whole document most worth getting right before anything else is tuned.

---

## 8. Resource System — **[PROPOSED]**

This is the mechanical heart of Pillar 3.1 (Discovery and the Gold Rush) and deserves the deepest treatment in this pass.

**Design goal:** resources must be interesting to find, not just interesting to have. A resource system where "Iron" is a static, always-available, always-identical commodity produces a crafting system where the only variable is time spent — no discovery, no urgency, no reputation for having sourced something good. Everything below exists to prevent that.

**SWG inspiration:** historical SWG resources spawned procedurally across the world with randomised quality attributes drawn from a shared attribute pool (only a relevant subset applied per resource type), existed for a limited time before despawning and being replaced, and were the direct input to a crafting system where output quality depended on input quality. That structure — taxonomy, rolled attributes, temporal spawning, quality-propagation into crafting — is adapted here under original names. What's proposed as new: an explicit, documented Toxicity attribute, and a more deliberate two-tier extraction split (manual vs. installed Harvesters) than this designer can confidently attribute to any single historical implementation.

### 8.1 Resource Taxonomy — `RES-001`

Three-level hierarchy:

- **Domain** — the broadest grouping. Proposed domains: **Mineral**, **Chemical**, **Gaseous**, **Aquatic**, **Flora**, **Fauna-derived**.
- **Category** — a subdivision within a Domain that determines *which* attributes (from the universal pool below) are relevant and what schematics accept it (e.g., within Mineral: Ferrous Metal, Non-Ferrous Metal, Ore, Gemstone).
- **Specific Resource Class** — a concrete, game-generated instance within a Category (e.g., a particular named alloy) with attribute values rolled within ranges the Category defines. This is the actual unit that spawns, gets harvested, and gets traded.

This three-level split is what makes the system expandable per Section 22: adding a new Specific Resource Class never requires new code, only new data referencing an existing Category.

### 8.2 Universal Attribute Pool — `RES-002`

| Attribute | Abbrev. | Typically relevant to |
|---|---|---|
| Overall Quality | OQ | All — derived/composite, always present |
| Structural Integrity | SI | Mineral, Fauna-derived |
| Malleability | ML | Mineral |
| Conductivity | CD | Mineral, Chemical |
| Heat Resistance | HR | Mineral, Chemical |
| Cold Resistance | CR | Mineral, Flora |
| Decay Resistance | DX | Flora, Fauna-derived, Chemical |
| Purity | PF | Chemical, Aquatic, Flora (consumable-relevant) |
| Potential Energy | PE | Chemical, Gaseous |
| Density | DC | Mineral, Aquatic |
| Elemental Resistance | ER | Mineral, Chemical |
| Toxicity | TX | Chemical, Flora, Fauna-derived *(new for TCIndustries — not a direct SWG adaptation)* |

Each Category defines which subset of this pool applies and the roll range for each. A Specific Resource Class never rolls all twelve — most Categories should use four to seven. Exact ranges and roll distributions are **[TBD]**.

### 8.3 Spawning and Resource Pools — `RES-003`, `RES-004`

Resources spawn as **Resource Pools**: a Specific Resource Class instance, rolled attribute values, a geographic footprint within a Zone, and a lifespan. Proposed lifecycle (`RES-004`): pools are **time-limited, not extraction-limited** — a pool despawns on a timer regardless of how much has been harvested from it, rather than running out from overharvesting. This is a deliberate choice over a hard-depletion model, for two reasons: it avoids the "someone strip-mined my spot" frustration of a hard-cap system, and it guarantees rotation — no player or pool can be permanently camped, because it will disappear on schedule and something new will appear elsewhere. Whether overharvesting should *accelerate* despawn as a secondary effect is left **[TBD]** as an optional refinement, not a base assumption.

Pool lifespans should vary (short/medium/long-lived tiers) so that not every discovery is equally urgent — a long-lived pool rewards patient exploitation, a short-lived one rewards the literal gold-rush reaction the pillar is named for. Exact durations: **[TBD]**.

### 8.4 Quality and Rarity — `RES-005`

Overall Quality (OQ) is derived from how favourably a pool's attributes rolled relative to the Category's possible range. Exceptional spawns — rare, high-OQ rolls — are the pillar's core moment and should be rare enough that players genuinely react to them (chase them, hoard them, fight over access to them) rather than shrug at them. Target rarity is **[TBD]**, to be set once simulation tooling exists (Section 26).

### 8.5 Discovery — `RES-006`

Discovery is itself a skill-gated activity (a Surveying/Prospecting-type Skill Discipline, Section 7), not a free map overlay. Lower-skill surveying should return vague information (a general area, a rough quality band); higher skill should return precision (near-exact location, a confident quality read). This makes "I know where the good stuff is" a genuine, saleable specialisation in its own right, and creates an information economy layered on top of the physical resource economy — a specialist surveyor can sell tips, sell their services outright, or exploit the find themselves.

### 8.6 Extraction — `RES-007`

Two extraction modes:

- **Manual harvesting** — direct, hands-on gathering. No infrastructure required. Appropriate for Flora/Fauna-derived resources and for low-commitment or early-game participation.
- **Installed Harvesters** — player-placed structures set on a Resource Pool's location, requiring claimed land and ongoing maintenance (see `MFG-003`, Section 10), that extract passively over time up to a throughput cap. This is the connective tissue between the Resource System and Manufacturing.

### 8.7 Processing — `RES-008`

Many Categories require a **processing/refining** step (raw → refined) before becoming eligible for higher-tier schematics. Processing is its own activity/specialisation, not an automatic side effect of harvesting — this exists specifically so that "I harvest," "I refine," and "I craft" can be three different players' jobs rather than one player's trivial pipeline, directly serving Pillar 3.3.

### 8.8 Resource Markets — `RES-009`

Harvested and processed resources are tradeable goods (via Section 11's Economy and Section 12's Retail systems). A traded resource lot carries its rolled attributes as visible metadata, so buyers can genuinely shop for quality rather than trusting a flat commodity price. This visibility is what lets a harvester or refiner build a reputation for "always has good ore" — without it, Pillar 3.2 has nothing to attach to on the sourcing side of the economy.

### Interactions
Resources feed Crafting (Section 9) directly through schematic thresholds; feed Manufacturing (Section 10) through Harvester structures; feed the Economy (Section 11) as a tradeable good category in their own right, independent of anything crafted from them.

### Exploit / Failure Modes
- A pool with too-generous a lifespan effectively becomes a permanent, campable resource — undermining the rotation the timer is supposed to guarantee. Needs simulation-backed tuning, not a guess.
- If Surveying skill investment isn't meaningfully better than random wandering, `RES-006`'s entire information-economy hook collapses.
- If processing (`RES-008`) is trivially fast/cheap, it stops being a real specialisation and just becomes a tax on harvesters' time.

### Open Questions
See `OQ-03` through `OQ-06` in Section 28 (exact attribute ranges, spawn algorithm, overharvesting effects, pool lifespan tiers).

---

## 9. Crafting System — **[PROPOSED]**

**Design goal:** crafting must be the mechanism that turns resource quality into product identity — and it must be the *only* reliable way to reach the top of the quality curve, because that exclusivity is what protects Pillar 3.2 from being undercut by Manufacturing (Section 10).

**SWG inspiration:** historical SWG crafting combined resource-attribute thresholds, crafter skill, and a manual experimentation minigame (spending experimentation points across sub-attributes with success/critical-success/failure odds) to produce variable-quality output, including rare "exceptional" results. That structure is adapted directly below. The permanent, always-visible maker's mark on exceptional output is treated here as the single most important thread connecting this section back to Pillar 3.2, and is called out explicitly as a rule this designer recommends locking early (see the flag under 9.5).

### 9.1 Schematics — `CRF-001`

A **Schematic** defines: required input Resource Categories (not specific instances — any Class within the Category qualifies) with minimum and ideal attribute thresholds per relevant attribute; a required Tool/Station tier; a Skill Discipline/Box prerequisite; and, for more complex items, required **sub-component schematics** — crafted parts that must themselves be produced first. Nested schematics deepen production chains deliberately: a complex item shouldn't be one crafter's one-step action, it should touch several specialists' output.

### 9.2 Quality Propagation — `CRF-002`

Output quality is a function of: (a) how favourably the actual input resources' attributes compare to the schematic's thresholds, (b) the crafter's relevant skill/mastery level, (c) the Tool/Station's quality tier, and (d) the Experimentation outcome (below). The exact formula combining these is **[TBD]** — this document commits to the four *inputs* the formula must use, not the formula itself, to avoid the false-precision problem described in Section 24.

### 9.3 Experimentation — `CRF-003`

After assembling a schematic's base components, the crafter allocates an **Experimentation Point** pool (sized by skill) across the item's defined sub-attributes — e.g., Durability, Efficiency, Quality, and an Appearance/customization slot — with a probability curve per point spent that can succeed, critically succeed, or fail. This point-allocation minigame is the "soul" of the crafting system, carried over from SWG largely intact because it is the specific mechanic that makes two crafters' output of the *same* schematic meaningfully different — without it, skill differentiation between crafters collapses to a single scalar and Pillar 3.2 has much less to work with.

### 9.4 Risk — `CRF-005`

Experimentation carries genuine failure risk — a bad roll can consume input resources without producing a usable item, or lock in a worse-than-intended result. This risk is what gives a rare, high-quality resource lot real stakes: a low-skill crafter attempting an ambitious item with a scarce resource is making a real gamble, not a formality. Removing this risk (e.g., a "safe mode") would flatten the gold-rush pillar's payoff, since the value of an exceptional resource partly comes from the fact that using it well isn't guaranteed.

### 9.5 Exceptional Items and Provenance — `CRF-004`

Rare critical-success outcomes produce **Exceptional** items. Every crafted item — exceptional or not — permanently stores its maker's identity as visible metadata (**provenance**), inspectable by anyone who examines it. This is the direct, load-bearing mechanism for Pillar 3.2 ("players seek out products made by a particular player"), and this designer recommends it be an early candidate for promotion to **[LOCKED]**, since a large amount of downstream design (retail branding, reputation, crafter renown, the entire "famous crafter" emergent-gameplay case in Section 21) assumes provenance is permanent and unforgeable. As proposed, provenance cannot be stripped, transferred, or faked — it is tied to the crafting character's verified identity, which also matters for the fraud considerations in Section 25.

### 9.6 Specialization — `CRF-006`

Schematics are gated behind specific Skill Disciplines. Given the Skill Point cap (`CHR-003`), no single character can hold every crafting discipline at a competent level — this is what keeps crafting specialisation real rather than a matter of "eventually learn everything."

### Interactions
Consumes Resources (Section 8) directly via schematic thresholds; feeds Items (Section 17) with the quality/durability/provenance data those items carry; feeds Manufacturing (Section 10) via the Production Template concept (a manually-crafted item becomes a factory's template); feeds Retail (Section 12) and the Economy (Section 11) as the source of tradeable goods; feeds Combat (Section 16) as the source of equipment.

### Exploit / Failure Modes
- If Experimentation risk is too low, quality differentiation between skilled and unskilled crafters shrinks and Pillar 3.2 weakens.
- If a non-manual path (see Manufacturing, `MFG-002`) can ever reach Exceptional quality, this entire section's connection to Pillar 3.2 breaks — flagged again below as the single most important cross-system rule in this document.
- Provenance data must be resistant to laundering (trading an item doesn't erase who made it) or reputation becomes gameable.

### Open Questions
See `OQ-07`, `OQ-08` (exact quality formula, experimentation probability curve).

---

## 10. Manufacturing — **[PROPOSED]**

**Design goal:** allow bulk production of *already-solved* designs without ever letting automation replace the interesting, identity-bearing part of crafting. This section exists specifically to satisfy Section 4's Design Scale principle and is the single biggest lever against Factorio/Satisfactory-style drift (Section 23).

### 10.1 The Automation Boundary — `MFG-006`

Stated as plainly as possible, because this sentence should anchor every other rule in this section:

> **Automation may exist for repetition. It must never exist for discovery, for breaking a quality ceiling, or for identity.**

Everything below is an attempt to make that sentence enforceable rather than aspirational.

### 10.2 Factories and Production Templates — `MFG-001`, `MFG-002`

A **Factory** can mass-produce a schematic only after a crafter has manually crafted and "locked in" a specific configured version of it — a **Production Template** (SWG-adapted from the historical "factory crate" concept). The Factory then repeats that exact configuration at volume.

`MFG-002` **[PROPOSED — flagged for early lock]**: **Factories can never produce Exceptional-tier output.** A Production Template locks in a fixed, non-exceptional quality level; the Factory reproduces that level exactly, at volume, and nothing a Factory produces can ever match what a skilled crafter's manual Experimentation can achieve. This is, alongside `CRF-004`, the other single most load-bearing rule in this document — it's the concrete mechanism that keeps "I want the *best* version of this item" pointed at a specific player rather than at whoever owns the biggest factory. This designer strongly recommends this rule be promoted to **[LOCKED]** before any numeric tuning work begins on Manufacturing, because a huge amount of downstream economic and reputation design assumes it holds.

### 10.3 Throughput and Maintenance — `MFG-003`

Factories have a per-cycle throughput cap and require ongoing maintenance (power/currency/resource upkeep) to remain active. An unmaintained Factory should stop producing, not merely slow down — maintenance needs to be a real, recurring decision, not a one-time setup cost, because recurring cost is what keeps Manufacturing from being "buy once, print forever."

### 10.4 Ownership Limits — `MFG-004`, `MFG-005`

Two soft caps, both economic rather than hard denials, working together to enforce Section 4's Design Scale principle:

- **Rising marginal upkeep** (`MFG-004`): each additional Factory/Harvester a single character owns costs proportionally more to maintain than the last, so scaling up single-handedly gets progressively less efficient rather than linearly easier.
- **Land scarcity** (`MFG-005`): Factories and Harvesters require claimed land, which is a finite, shared world resource (Section 18). A player cannot simply buy their way to infinite production capacity if there's nowhere left to put the structures.

Together, these mean large-scale production realistically requires either multiple specialised characters, trusted collaborators, an organisation, or straightforward reliance on trading with other specialists — not a purely solo route. Exact upkeep curves and land allotments are **[TBD]**.

### Interactions
Consumes Resources (Section 8) at defined rates via Harvesters; consumes Crafting output (Section 9) as its Production Template source; interacts with the Economy (Section 11) as a maintenance-cost sink; interacts with Buildings (Section 18) for placement and land.

### Exploit / Failure Modes
- If maintenance costs don't scale (`MFG-004` absent or too flat), a single well-capitalised player can trivially out-produce a whole specialist community — the exact single-player-industrial-empire outcome Section 4 explicitly forbids.
- If land is too abundant, `MFG-005`'s cap does nothing; if too scarce, new players may never get a foothold. Needs simulation-backed tuning.
- Structures abandoned-and-reclaimed to dodge upkeep cycles is a plausible exploit pattern; flagged again in Section 25.

### Open Questions
See `OQ-09`, `OQ-10` (upkeep curve shape, land allotment per character/account).

---

## 11. Economy — **[PROPOSED]**

**Design goal:** currency should flow through player hands, not accumulate frictionlessly, and the game's own systems (not moderation after the fact) should be what keeps the economy from either collapsing into worthlessness or freezing up from hoarding.

### 11.1 Currency — `ECN-001`

A single unified currency is assumed for this pass, for simplicity. A dual-currency model (e.g., separating a common trade currency from a rarer prestige currency) is a plausible **[TBD]** alternative worth a dedicated future pass, not adopted here by default.

### 11.2 Pricing — `ECN-002`

Pricing on player-made goods is entirely player-set — no algorithmic price floors or ceilings on the player market. World-provided NPC vendors at Established Settlements (Section 6.1) sell only **baseline, non-exceptional goods**, deliberately kept below competitive quality with anything a real crafter would sell, so they function as a safety-net floor for new players rather than a competing seller. This is the direct mechanism that stops Established Settlements from quietly replacing player commerce.

### 11.3 Degradation and Replacement Demand — `ECN-003`

Equipment loses condition through use and requires repair (a baseline self-repair option, plus a better, profession-gated repair service — see `SVC-001`), eventually reaching a worn-out state. This is the economy's primary long-term demand engine: a permanently-durable item, once crafted, would eventually saturate demand and collapse the market for that item category entirely — this is called out explicitly as a failure mode to guard against, not just a flavour detail.

### 11.4 Economic Sinks — `ECN-004`

Enumerated sinks proposed so far, consolidated from earlier sections: Factory/Harvester maintenance (`MFG-003`), Respecialization cost (`CHR-004`), item repair cost (`ECN-003`), city taxation (`SOC-004`, Section 19), experimentation resource loss on failure (`CRF-005`), and consumable services — medicine, buffs, transport fares (Section 15). Whether these sinks, in aggregate, roughly balance whatever currency faucets exist is a numeric question this pass cannot answer without simulation tooling — see Section 26.

### 11.5 Market Infrastructure — `ECN-006`

A searchable **Commerce Directory** is proposed: a cross-town/region listing search (SWG-adapted from its Bazaar/vendor-search concept) that lets buyers find available listings without physically walking every shop. Deliberately proposed as a *search layer over physical vendors*, not a disembodied NPC-run auction house — every listing still resolves to a real player shop with a real owner, so this feature helps discovery without disintermediating Retail's location- and reputation-based design (Section 12).

### 11.6 Exploit Risks — `ECN-005`

Flagged here, expanded in Section 25: resource/currency duplication (a technical/TDD concern, noted here because it's an economy-design dependency), multi-boxing or alt-account use to fake self-sufficiency and quietly defeat Pillar 3.3, vendor price manipulation and wash trading, and automated "vendor-camping" bots. Mitigations for the account-policy-adjacent ones (multi-boxing especially) are **[TBD]** and arguably outside pure game design — flagged as needing a policy decision alongside the design one.

### Interactions
Every other system feeds the Economy: Resources and Crafting as goods, Manufacturing and Buildings as sinks, Services as recurring sinks, Cities as a taxation layer, Retail as the transaction surface.

### Open Questions
See `OQ-11` through `OQ-13` (currency model, sink/faucet numeric targets, multi-boxing mitigation policy).

---

## 12. Retail — **[PROPOSED]**

**Design goal:** give reputation (Pillar 3.2) and place (Pillar 3.4) a physical, visitable form.

### 12.1 Shops and Vendors — `RTL-001`

Player shops are owner-branded structures (Section 18) where goods are listed for sale, either attended (a player-run stall) or unattended (a vendor proxy that sells while the owner is offline). Every listing visibly carries its maker's provenance (`CRF-004`) where applicable, so browsing a shop is also browsing a reputation.

### 12.2 Reputation Display — `RTL-002`

Reputation display is proposed conservatively for launch: visible **sales volume** and **maker tags on individual items** are enough signal to start with. A formal star-rating or review system is explicitly **[DEFERRED]**, not because it wouldn't be valuable, but because rating systems are unusually easy to abuse (fake reviews, review-bombing, brigading) and deserve a dedicated anti-abuse design pass rather than being bolted on as an afterthought here.

### 12.3 Commercial Districts — `RTL-003`

Commercial districts should emerge organically around Settlements and Player Cities rather than being fully pre-planned. Location — proximity to transport nodes and foot traffic — becomes a scarce, tradeable form of prestige in its own right, reinforcing both Pillar 3.2 (a shop's location can itself be part of its reputation) and Pillar 3.4.

### Interactions
Surfaces Crafting (Section 9) and Resource (Section 8) output for sale; is the physical target that the Commerce Directory (`ECN-006`) search layer points back to; anchors Cities' economic identity (Section 19).

### Open Questions
See `OQ-14` (what reputation signal, if any, should exist beyond sales volume and maker tags before a full rating system is designed).

---

## 13. Professions — **[PROPOSED]**

**Design goal:** define profession as a flexible, skill-based architecture — not a rigid class list — while still making mastery of any one thing feel rare and impressive.

### 13.1 Profession Families — `PRF-001`

Professions are organised into families, each containing multiple Skill Disciplines (Section 7) a character can invest in independently:

| Family | Representative disciplines |
|---|---|
| Combat | Ranged, melee, defensive/tank-leaning, support-in-combat |
| Medical | Field treatment, advanced/injury-specific treatment |
| Entertainment | Performance-based buffs, venue operation |
| Crafting | One or more disciplines per major item category (weapons, armour, structures, consumables, etc.) |
| Harvesting / Resource Specialist | Surveying, manual harvesting, Harvester operation |
| Architecture / Construction | Structure design and placement, city infrastructure |
| Commerce / Trade | Shop operation, negotiation-relevant bonuses, logistics |
| Research | **[DEFERRED]** — reserved family name for a future science/discovery-adjacent profession line; not designed in this pass |
| Transportation / Logistics | Player-operated transport services (ties to `WLD-006`) |

This table is intentionally a skeleton, not a final list — new disciplines and even new families should be addable as data under Section 22's expansion rules without restructuring this taxonomy.

### 13.2 Hybridization — `PRF-002` **[LOCKED — restates the brief's Section 5 requirement concretely]**

A character may invest Skill Points across any combination of families, subject only to the Skill Point cap (`CHR-003`). There is no profession-exclusivity rule preventing, say, a character from being part-Medical and part-Crafting.

### 13.3 Mastery Scarcity — `PRF-003`

Because the Skill Point cap is finite, full mastery of a top-tier discipline should cost enough of that budget that a character can realistically master very few of them — this is what makes "the Master Armoursmith" a meaningful, rare thing to be, rather than an inevitable end-state everyone reaches eventually.

### Interactions
Directly gates Crafting schematics (`CRF-006`), Combat capability (Section 16), Service provision (Section 15), and Resource discovery (`RES-006`).

### Open Questions
See `OQ-15` (final discipline list per family — deliberately incomplete here).

---

## 14. Social Systems — **[PROPOSED]**

**Design goal:** give players structures to organise *around*, at more than one scale.

### 14.1 Two-Tier Structure — `SOC-001`

- **Player Associations**: small-scale, informal-to-semi-formal groupings — a shop's staff, a small crafting collective, a standing adventuring party. Low overhead to form.
- **Player Cities**: larger, place-based, governance-bearing settlements. Full treatment in Section 19, since Cities also function as Buildings/world-structure content.

### 14.2 Reputation as a Social Layer — `SOC-005`

Reputation (Pillar 3.2) is not purely an individual-character stat — it should be able to attach to an Association or City's *name* as well (a district known for good armour, a trading company known for reliable supply), giving groups, not just individuals, something to build and protect. Mechanically **[TBD]** pending Retail's `RTL-002` reputation-display decision, since a group-level reputation system likely wants to reuse the same underlying signal (sales volume, maker tags) rather than invent a separate one.

### Interactions
Cities (Section 19) are this section's largest concrete expression; Retail (Section 12) and Crafting (Section 9) provide the reputation substrate; Services (Section 15) are frequently organised through Associations (e.g., a standing medical-service collective).

### Open Questions
See `OQ-16` (Association tooling — permissions, shared storage, formal roles — none of which is designed in this pass).

---

## 15. Services — **[PROPOSED]**

**Design goal:** create economic relationships that are *recurring*, not one-time — because a durable good satisfies demand once, but a service has to be bought again.

### 15.1 Service Archetypes — `SVC-001`

Proposed initial list, each requiring a provider-side Skill Discipline and each with a real consumer-side reason to exist: **Medical** (injury/wound treatment beyond passive regeneration — see `CBT-002`), **Entertainment** (performance-based buffs, morale/fatigue recovery), **Transportation** (player-operated routes and charters, per `WLD-006`), **Construction** (structure placement/upgrade assistance), **Repair** (better-than-baseline item restoration, per `ECN-003`), **Research** (**[DEFERRED]**, reserved alongside the Research profession family), and **Commissioned Crafting** (a specialist producing to a customer's specific order rather than to open-market stock).

### 15.2 Why Recurrence Matters — `SVC-002` **[flagged for early lock]**

This designer recommends explicitly locking the principle (not necessarily every archetype above) that **services must be consumable or time-limited by design** — a buff that expires, medicine that gets used up, a repair that wears off again with use. A permanent, one-time buff or a repair that never needs repeating would quietly convert a Service into a disguised Item, losing the recurring economic relationship that makes Services a distinct and important category for Pillar 3.3. Without recurrence, "I need a doctor" stops being true after the first visit, and the interdependence pillar loses one of its steadiest engines.

### Interactions
Combat (Section 16) is the heaviest consumer of Medical and Entertainment services; Manufacturing/Buildings (Sections 10, 18) consume Construction; the whole game consumes Repair and Transportation.

### Exploit / Failure Modes
A service that can be trivially self-provided (e.g., a combat character who can also fully self-heal without meaningful cost) undercuts this section's entire purpose — flagged as an Interdependence Check against Combat design specifically (see Section 16).

### Open Questions
See `OQ-17` (exact service list is not claimed to be final or exhaustive).

---

## 16. Combat — **[PROPOSED / PARTIALLY DEFERRED]**

Per the brief's own framing (Section 11.K): combat is important but must not consume the entire design. This section defines combat's *interdependence hooks* only — how it connects to the rest of the world. Combat's internal design (abilities, classes-if-any, PvP ruleset, damage formulas) is explicitly **[DEFERRED]** — see `CBT-003` and Section 30.

### 16.1 Combat as a Consumer, Not a Centre — `CBT-001`

Combat should be designed as a heavy *consumer* of the crafting, medical, and entertainment economy rather than a self-contained system:

- Weapons and armour come from Crafting (Section 9) and degrade per `ECN-003`, creating steady replacement and repair demand.
- Recovery beyond passive regeneration requires Medical services (`SVC-001`).
- Pre-engagement buffs come from Entertainment/support-adjacent services.
- Post-engagement recovery venues (social spaces analogous to a tavern) are a natural hook into Social Systems (Section 14) and Retail-adjacent commercial districts (Section 12).

### 16.2 Wound System — `CBT-002` **[PROPOSED]**

A wound/injury layer beyond a simple hit-point pool is proposed (SWG-adapted, historical behaviour), specifically because a pure HP-and-potions model tends to make Medical services skippable once a player stockpiles enough consumables. An injury system that resists simple item-spam and requires an actual Medical specialist's skill to treat properly keeps `SVC-001`'s Medical archetype genuinely necessary rather than a convenience. Exact mechanics are **[TBD]**; only the design requirement ("healing beyond a baseline should require another player's skill, not just an inventory of potions") is proposed here.

### 16.3 Deferred Scope — `CBT-003` **[DEFERRED]**

Not addressed in this pass: specific abilities/ability trees, whether PvP exists and under what ruleset, encounter design, enemy/creature design, group role structure. All of these are legitimate, large design efforts that deserve their own pass — see Section 31's recommended sequence.

### Interdependence Check
The single biggest risk to this section is a self-sufficient combatant: any ability, consumable, or mechanic that lets a combat character fully sustain themselves (full self-repair, full self-healing, no consumable dependency) without ever touching another profession's output quietly defeats Pillar 3.3 from the combat side. This check should be applied to every concrete combat ability once they're designed, not just to this section's abstractions.

---

## 17. Items — **[PROPOSED]**

**Design goal:** items should be able to carry a maker's identity and should not be permanent — see `ECN-003`'s replacement-demand argument.

### 17.1 Quality and Durability — `ITM-001`

Items carry a quality tier (set at crafting time, per Section 9) and a durability/condition stat that depletes with use and is restored via the Repair service (`SVC-001`) or a lesser baseline self-repair option.

### 17.2 Tradeability — `ITM-002` **[PROPOSED — flagged for early lock]**

Items are proposed to be **fully tradeable by default, with no account/character soul-binding.** Soul-binding (an item becoming permanently non-tradeable once acquired or equipped) is a common MMO pattern, but it is flagged here as actively hostile to this design's Pillar 3.4/economy goals: it removes items from the player market permanently, undercutting the exact secondary economy the Retail and Economy sections depend on. This designer recommends the no-soul-binding rule be locked early, as an explicit rejection of a very common genre default rather than an oversight.

### 17.3 Provenance and Customization — `ITM-003`

Every crafted item stores maker provenance (`CRF-004`) permanently. A visual customization/appearance layer, separate from the mechanical Experimentation outcome, is proposed as a stretch addition — letting a maker's items be recognisable on sight, not just on inspection — but is not designed in depth here.

### Interactions
Produced by Crafting (Section 9) and Manufacturing (Section 10); consumed by Combat (Section 16) and everyday activity; traded via Retail (Section 12) and the Economy (Section 11).

---

## 18. Buildings — **[PROPOSED]**

**Design goal:** give the "Player-Created Society" pillar somewhere to physically live, while keeping land a genuinely scarce resource (per `MFG-005`).

### 18.1 Structure Categories — `BLD-001`

**Personal** (housing), **Commercial** (shops, per Section 12), **Production** (Harvesters and Factories, per Sections 8 and 10), and **Civic** (City Hall-equivalent, banks, transport hubs — placed by Cities, Section 19).

### 18.2 Placement Rules — `BLD-002` **[TBD]**

Structures require claimed land, are subject to zoning and density rules (to prevent unbounded visual/world clutter), and should have some distance/spacing requirement relative to other structures. Exact numbers are undesigned — this is flagged as a placeholder abstraction only.

### 18.3 Maintenance and Abandonment — `BLD-003`

Structures require ongoing maintenance (mirroring `MFG-003`'s logic for Production structures, extended to all structure types). An unmaintained structure should decay and eventually become reclaimable, returning its land to the available pool — this is the mechanism that keeps land scarcity (`MFG-005`) meaningful over the long run rather than gradually accreting into a fully-claimed, permanently-locked map.

### Interactions
Hosts Production (`MFG-001` onward), Retail (Section 12), and Civic/City (Section 19) structures; interacts with the Economy as a maintenance sink (`ECN-004`).

### Open Questions
See `OQ-18` (zoning/density specifics, maintenance decay curve).

---

## 19. Cities — **[PROPOSED]**

**Design goal:** the most concrete, ambitious expression of Pillar 3.4 — a city should be a place players actually think of as theirs.

### 19.1 Founding — `SOC-002`

A Player City is founded once a cluster of player-owned structures in proximity crosses a threshold (structure count and/or nearby-citizen count), triggering a founding action that establishes a City Hall-equivalent Civic structure. Historical SWG behaviour, adapted.

### 19.2 Ranks and Services — `SOC-005` (city-specific extension)

Cities should unlock Rank/Tier progression as they grow (by population, structure count, or tax revenue), granting city-wide services and bonuses — banking, a transport hub connection, production bonuses for a chosen specialisation, and similar — chosen by city leadership. Specific service list and thresholds: **[TBD]**.

### 19.3 Governance — `SOC-003` **[TBD]**

Whether a Mayor role is filled by election or by founder-appointment is genuinely undecided. This designer's inclination is toward election, since it more directly serves the "living society" fantasy (Section 2.2) than simple appointment — but election systems carry real griefing risk (vote manipulation, a hostile or absentee mayor damaging a city with no recourse) that needs its own design pass, including a recall/impeachment mechanism, before this can be responsibly locked. Left **[TBD]** rather than guessed at.

### 19.4 Taxation — `SOC-004` **[TBD]**

That city taxation exists as an economic sink (`ECN-004`) is assumed; what it taxes (vendor sales? structure maintenance? a flat citizen fee?) and who sets the rate are undecided.

### Interactions
Built from Buildings (Section 18); hosts Retail (Section 12) commercial districts; is the primary target of Social Systems' larger-scale reputation (`SOC-005`); interacts with the Economy as both a sink (taxation) and a market hub.

### Exploit / Failure Modes
A hostile or negligent mayor with no recall mechanism could unilaterally damage a city's economy or membership — the single biggest reason `SOC-003` is left TBD rather than locked in this pass.

### Open Questions
See `OQ-19` through `OQ-21` (governance model, taxation base, rank thresholds and rewards).

---

## 20. Progression — **[PROPOSED]**

**Design goal:** progression should support freedom of identity (Section 5's sandbox philosophy), not funnel every player down one measured path.

### 20.1 Philosophy

Restates and consolidates `CHR-002` through `CHR-005`: use-based Skill Point acquisition, spent across Skill Boxes within Skill Disciplines, subject to a capped total budget, with a friction-bearing (not free) Respecialization option. Progression should feel like **accumulating a résumé**, not climbing a single ladder — there is no unified "character level" implied anywhere in this design, deliberately, since a single level number would quietly re-introduce the "who's ahead" framing this design is trying to avoid.

### 20.2 What the Skill Point Cap Must Guarantee — `CHR-003` (restated with a concrete target)

The cap should be set so that a dedicated player can develop roughly two to three professions to a meaningful working level, or one profession to full top-tier Mastery plus a modest secondary capability — but not both a top-tier Mastery *and* two other full professions simultaneously. This is stated here as a design *requirement* the eventual numbers must satisfy, not as the numbers themselves (which remain **[TBD]**).

### Interactions
This section is a restatement/consolidation layer over Section 7 (Character and Skill System) rather than a new system — kept separate here only because the brief's own Section 16 structure calls for a distinct Progression section.

### Open Questions
See `OQ-22` (whether any secondary, non-use-based Skill Point source should exist).

---

## 21. Emergent Gameplay — **[PROPOSED]**

Explicitly identifying the emergent events this design is *for*, and which system(s) each depends on, so future tuning work can check "is this actually happening" against a real list rather than a vague hope.

| Emergent event | Generated by |
|---|---|
| Resource rushes | `RES-003`/`RES-004` spawn/despawn timing + `RES-005` rarity |
| Market shortages/spikes | Resource despawn timing interacting with Economy (`ECN` family) |
| Famous crafters | `CRF-004` provenance + Pillar 3.2 |
| Regional economies | Resource regional variation (`WLD-001`) + transportation cost/time (`WLD-006`) |
| Player-created supply chains | Manufacturing (`MFG` family) + Interdependence (Pillar 3.3) |
| Economic crises (inflation/deflation swings) | Economy sinks/faucets (`ECN-004`) — needs simulation testing, Section 26 |
| Social hubs | Cities (`SOC-002`) + Entertainment venues (`SVC-001`) |
| Specialist dependencies | Profession gating (`PRF-001`) + Services' recurrence (`SVC-002`) |

If, after implementation and testing, any row in this table isn't actually happening, that's a signal the underlying system needs revisiting — this table is meant to function as a checklist, not just documentation.

---

## 22. Content Expansion — **[LOCKED — restates the brief's Section 6 requirement]**

New content must be addable as **data referencing existing abstractions**, not as new hard-coded systems. Concretely, per system:

- New Resource Classes reference an existing Category and roll within it — no new code (`RES-001`).
- New Schematics reference existing Resource Categories, Attributes, and Skill Disciplines (`CRF-001`).
- New Skill Boxes slot into existing (or cleanly new) Skill Disciplines/Families (`CHR-002`, `PRF-001`).
- New Structure types reference existing zoning/maintenance rules (`BLD-001`–`BLD-003`).
- New Zones/Regions reference existing Resource Pool spawn rules (`WLD-001`).

The practical test for any future system proposal: **can it be expressed as a new row of data against the taxonomies already defined in this document, or does it require a new taxonomy?** The former is cheap and expected; the latter is a real design decision and should go through the same status-tagging discipline as everything in this document, not be quietly added to game data.

---

## 23. Anti-Goals — **[LOCKED — restates the brief's Section 13 directly]**

TCIndustries must not gradually become:

| Anti-goal | Primary guardrail(s) already in this document |
|---|---|
| Factorio / Satisfactory | `MFG-002` (no factory-produced Exceptional items), `MFG-004`/`MFG-005` (scaling cost + land scarcity) |
| Capitalism-style business sim | Pillar 3.2/3.3 mechanics throughout — value comes from *who* made something, not just throughput |
| EVE-style spreadsheet gameplay | `RTL-002`'s conservative reputation approach; provenance (`CRF-004`) keeps identity visible and human, not purely numeric |
| Conventional theme-park MMO | `WLD-003` (low-authored-content wilderness), Section 5's sandbox commitment |
| Quest-driven RPG | No unified quest-chain structure proposed anywhere in this document |
| Conventional class-based MMO | `CHR-002`/`PRF-002` (Skill Discipline hybridization, no rigid classes) |
| Single-player crafting game | Pillar 3.3 Interdependence Checks applied throughout; `MFG` family scaling limits |

Industrial depth, economic simulation, and automation are all explicitly welcome per the brief — the constraint is that they must serve Pillar 3.4's core fantasy (living as a citizen in a player-driven society), not replace it. Where a future system increases efficiency but makes other players less relevant, forces every player down one path, reduces the economy to a spreadsheet, or lets a player become powerful without interacting with anyone — it should be treated as a drift risk and checked against this table before being accepted, per the brief's own Section 21 "most important principle."

---

## 24. Balance Philosophy — **[PROPOSED]**

Principles, deliberately not numbers, for four reasons:

1. **Verticality should stay shallow relative to horizontal breadth.** A modestly-skilled crafter with a good reputation should remain economically relevant, not be flatly obsoleted by a maxed-out crafter the way pure numeric power-creep systems tend to work. Identity and reputation (Pillar 3.2) should matter as much as raw stat ceiling.
2. **Scarcity should be real but not punishing.** Despawn timers (`RES-004`) create urgency, not permanent loss — a missed exceptional spawn is a "next time," not a "the game screwed me."
3. **Every strong solo-power mechanic gets checked against Pillar 3.3 before it ships**, using the Interdependence Check pattern used throughout this document: does this let a player skip needing another player?
4. **Numbers are explicitly out of scope for this document.** Skill costs, spawn rates, prices, decay curves, and formula constants are marked `[TBD]` throughout rather than invented, because a guessed number written into a GDD tends to get treated as a decision the moment someone builds against it — and an early wrong number is far more expensive to unwind than an honestly-marked gap. Section 8 of the development pipeline (simulation/testing, referenced in the brief) is where real numbers belong; this document defines what those numbers need to guarantee.

---

## 25. Exploit / Fraud Considerations — **[PROPOSED]**

Consolidated from flags raised throughout the document:

| Risk | Affected system(s) | Mitigation status |
|---|---|---|
| Resource/currency duplication | Economy, Resources | Technical/TDD concern — noted as a dependency, not solved here |
| Multi-boxing/alt abuse to fake self-sufficiency and bypass Pillar 3.3 | Economy, Professions | **[TBD]** — partly a policy question, not purely a design one |
| Vendor price manipulation / wash trading | Economy, Retail | **[TBD]** — transaction-history visibility proposed as a partial mitigation |
| Automated vendor-camping / market bots | Economy, Retail | Technical/TDD concern |
| Structure abandon-and-reclaim cycling to dodge maintenance | Manufacturing, Buildings | **[TBD]** — likely needs a reclamation cooldown |
| Reputation gaming (flooding fake "exceptional" claims, alt-account brand puffery) | Crafting, Retail | Mitigated in design by `CRF-004`'s permanent, verified-identity provenance — this property should be locked alongside `CRF-004` itself |
| City governance griefing (absentee/hostile mayor) | Cities | **[TBD]** — needs a recall/impeachment mechanism before `SOC-003` can be locked |

---

## 26. Testing Requirements — **[PROPOSED]**

Invariants worth testing once an implementation exists, framed as checkable statements rather than vague quality goals:

1. **No profession should be capable of full solo self-sufficiency** across gather → process → craft → sell → combat-support at competitive efficiency. Testable via simulation: run a single simulated character through the full loop and compare its efficiency against a specialised multi-character group.
2. **Currency sinks should approximately balance currency faucets** over a simulated economy run, with no runaway inflation or deflation. Exact target ratio is `[TBD]`, to be set once a faucet list exists (currently only sinks are enumerated in `ECN-004`).
3. **Exceptional resource spawns should remain rare** enough to stay exciting without becoming a hard bottleneck that stalls crafting entirely. Target rarity `[TBD]`; testable by measuring time-to-next-exceptional-spawn against typical crafting demand.
4. **Factory output must never statistically match manual-Exceptional output** — a hard rule (`MFG-002`) that should be directly testable by comparing output-quality distributions between the two production paths.
5. **The Skill Point cap must arithmetically prevent** simultaneous full Mastery of more than the intended number of top-tier disciplines — a straightforward arithmetic check once discipline costs and the cap are both set (`CHR-003`, `PRF-003`).

These are acceptance criteria for a *later* implementation and testing phase, not claims that any of this has been tested yet.

---

## 27. Status Register

Master table of every tagged decision in this document. IDs are stable and intended for future conversations to reference directly.

| ID | System / Decision | Status | Rationale | Dependencies |
|---|---|---|---|---|
| `VIS-001` | Core game vision statement | LOCKED | Brief's own stated core goal | — |
| `VIS-002` | Player fantasy ("citizen, not hero") | LOCKED | Brief's own stated fantasy | — |
| `VIS-003` | Target session experience | PROPOSED | This pass's illustrative interpretation | VIS-001, VIS-002 |
| `VIS-004` | Pillar 1: Discovery and the Gold Rush | LOCKED | Brief's stated pillar | RES family |
| `VIS-005` | Pillar 2: Identity and Reputation | LOCKED | Brief's stated pillar | CRF-004, RTL-002 |
| `VIS-006` | Pillar 3: Interdependence | LOCKED | Brief's stated pillar | Applied throughout |
| `VIS-007` | Pillar 4: Player-Created Society | LOCKED | Brief's stated pillar | BLD, SOC families |
| `VIS-008` | No forced combat path | LOCKED | Brief Section 5, restated directly | Applies to all professions |
| `WLD-001` | Region > Zone > Pool hierarchy | PROPOSED | Organisational abstraction for expansion | RES-003, EXP-001 |
| `WLD-002` | Settlement types (Established vs Player City) | PROPOSED | Gives new players a floor without capping player commerce | ECN-002, SOC-002 |
| `WLD-003` | Wilderness as low-authored-content space | PROPOSED | Anti-theme-park guardrail | Section 23 |
| `WLD-004` | World time / day-night cycle | TBD | Presentation choice, wants to pair with spawn cadence | RES-004 |
| `WLD-005` | Weather | DEFERRED | Brief itself lists as conditional | — |
| `WLD-006` | Transportation abstraction | PROPOSED | Enables transport-as-service | SVC-001 |
| `CHR-001` | Vitality / Stamina / Focus pools | PROPOSED | Ties crafting fatigue to services, not just combat | SVC-001 |
| `CHR-002` | Skill Discipline / Skill Box structure | PROPOSED | SWG-adapted; enables hybrid characters | PRF-001 |
| `CHR-003` | Skill Point Cap | PROPOSED (numbers TBD) | Single most load-bearing number for Pillar 3 | Section 24, 26 |
| `CHR-004` | Respecialization mechanism | PROPOSED | Safety net without trivialising identity | Section 4.3 |
| `CHR-005` | Use-based ("learn by doing") XP | PROPOSED | SWG-adapted sandbox progression | CHR-002 |
| `RES-001` | Resource taxonomy (Domain > Category > Class) | PROPOSED | SWG-adapted; enables data-driven expansion | EXP-001 |
| `RES-002` | Universal Attribute Pool (incl. new Toxicity) | PROPOSED | SWG-adapted + one new attribute | CRF-001 |
| `RES-003` | Resource Pool spawning | PROPOSED | Core gold-rush mechanism | WLD-001 |
| `RES-004` | Time-limited (not extraction-limited) despawn | PROPOSED | Avoids strip-mining frustration while guaranteeing rotation | RES-003 |
| `RES-005` | Overall Quality / rarity of exceptional spawns | PROPOSED | Rarity target needs simulation | Section 26 |
| `RES-006` | Skill-gated Surveying / discovery | PROPOSED | Makes "knowing where" a real specialisation | CHR-002 |
| `RES-007` | Manual harvesting vs. installed Harvesters | PROPOSED | Splits casual and infrastructure play | MFG-001 |
| `RES-008` | Processing/refining as a separate activity | PROPOSED | Deepens interdependence chain | CRF-001 |
| `RES-009` | Resource-lot quality metadata, tradeable | PROPOSED | Lets harvesters/refiners build reputation | RTL-002 |
| `CRF-001` | Schematic structure | PROPOSED | SWG-adapted | RES-001, CHR-002 |
| `CRF-002` | Quality propagation (4 inputs, formula TBD) | PROPOSED | Commits to inputs, not formula | Section 24 |
| `CRF-003` | Experimentation point-allocation | PROPOSED | SWG-adapted "soul" of crafting | CRF-002 |
| `CRF-004` | Exceptional items + permanent provenance | **PROPOSED — recommended for early LOCK** | Load-bearing for Pillar 3.2 | RTL-001, ITM-003 |
| `CRF-005` | Experimentation risk (resource-consuming failure) | PROPOSED | Gives rare resources real stakes | RES-005 |
| `CRF-006` | Schematic Skill Discipline gating | PROPOSED | Keeps crafting specialisation real | CHR-003 |
| `MFG-001` | Factories require a manual Production Template | PROPOSED | SWG-adapted | CRF-001 |
| `MFG-002` | Factories never produce Exceptional output | **PROPOSED — recommended for early LOCK** | The core Factorio-drift guardrail | CRF-004, MFG-006 |
| `MFG-003` | Factory/Harvester throughput cap + maintenance | PROPOSED | Prevents "buy once, print forever" | ECN-004 |
| `MFG-004` | Rising marginal upkeep per structure owned | PROPOSED | Soft cap on solo scaling | Section 4 |
| `MFG-005` | Land scarcity as a structural constraint | PROPOSED | Forces collaboration to scale | BLD-002 |
| `MFG-006` | Automation Boundary principle | **PROPOSED — anchor rule** | "Repetition yes, discovery/identity no" | MFG-002, CRF-004 |
| `ECN-001` | Single unified currency | PROPOSED (alt TBD) | Simplicity default, not a rejection of dual-currency | — |
| `ECN-002` | Player-set pricing; NPC vendors baseline-only | PROPOSED | Stops NPCs undercutting player economy | WLD-002 |
| `ECN-003` | Item degradation and replacement demand | PROPOSED | Long-term crafting demand engine | ITM-001 |
| `ECN-004` | Economic sinks (consolidated list) | PROPOSED | Inflation-control backbone | Section 26 |
| `ECN-005` | Exploit risk flags | PROPOSED (mitigations mostly TBD) | See Section 25 | Section 25 |
| `ECN-006` | Commerce Directory (search over physical vendors) | PROPOSED | Discovery without disintermediating Retail | RTL-001 |
| `RTL-001` | Player shops / vendors, owner-branded | PROPOSED | Physical home for reputation | CRF-004 |
| `RTL-002` | Conservative reputation display; ratings DEFERRED | PROPOSED / DEFERRED (ratings) | Rating systems need dedicated anti-abuse design | Section 25 |
| `RTL-003` | Commercial districts, location as prestige | PROPOSED | Reinforces Pillar 3.2 and 3.4 | SOC-005 |
| `PRF-001` | Profession Family taxonomy (skeleton) | PROPOSED | Organises Skill Disciplines, not final list | CHR-002 |
| `PRF-002` | Hybridization across families | LOCKED | Restates brief Section 5 directly | CHR-003 |
| `PRF-003` | Mastery scarcity via Skill Point cost | PROPOSED | Makes "the Master X" meaningful | CHR-003 |
| `SOC-001` | Player Associations vs. Player Cities (two-tier) | PROPOSED | Two scales of player organisation | SOC-002 |
| `SOC-002` | City founding trigger | PROPOSED | SWG-adapted | BLD-001 |
| `SOC-003` | City governance (elected vs. appointed Mayor) | TBD | Griefing risk needs its own pass | Section 25 |
| `SOC-004` | City taxation base and authority | TBD | Undesigned | ECN-004 |
| `SOC-005` | City ranks/services; group-level reputation | PROPOSED / TBD mix | Ranks proposed skeletally; reputation tie TBD | RTL-002 |
| `SVC-001` | Service archetypes (Medical, Entertainment, etc.) | PROPOSED | Not claimed exhaustive | CHR-002 |
| `SVC-002` | Services must be recurring/consumable by design | **PROPOSED — recommended for early LOCK** | Prevents Services collapsing into disguised Items | ECN-003 |
| `CBT-001` | Combat as economic consumer, not centre | PROPOSED | Keeps combat design scoped per brief | SVC-001, ECN-003 |
| `CBT-002` | Wound/injury system beyond simple HP | PROPOSED (mechanics TBD) | Keeps Medical service genuinely necessary | SVC-001 |
| `CBT-003` | Full combat ability/class/PvP design | DEFERRED | Explicitly out of scope this pass, per brief | Section 30 |
| `ITM-001` | Quality and durability/condition stat | PROPOSED | Standard item-state model | ECN-003 |
| `ITM-002` | No soul-binding; full tradeability by default | **PROPOSED — recommended for early LOCK** | Protects secondary economy from genre-default erosion | ECN-006 |
| `ITM-003` | Provenance + optional customization layer | PROPOSED | Extends CRF-004 to the item itself | CRF-004 |
| `BLD-001` | Structure categories (Personal/Commercial/Production/Civic) | PROPOSED | Organising taxonomy | EXP-001 |
| `BLD-002` | Placement / zoning / density rules | TBD | Numbers undesigned | MFG-005 |
| `BLD-003` | Maintenance decay to abandonment/reclamation | PROPOSED | Keeps land scarcity meaningful long-term | MFG-005 |
| `EXP-001` | Data-driven content expansion principle | LOCKED | Restates brief Section 6 directly | Applies to all taxonomies |
| `AG-001` | Anti-Goals list | LOCKED | Restates brief Section 13 directly | Section 23 table |
| `BAL-001` | No invented numeric precision in this pass | LOCKED | Restates brief Section 14 directly | Section 24 |

---

## 28. Open Questions

Consolidated from every `[TBD]` and inline flag above. Nothing here should be treated as answered by omission elsewhere in this document.

1. **OQ-01** — No prototype source was available this session. Should a reconciliation pass (comparing prototype behaviour against this GDD) happen before or after the next systems-design pass?
2. **OQ-02** — Should genre/setting/tone be formally decided before the next systems pass, or can systems design continue setting-agnostic for now? *(see AS-02)*
3. **OQ-03** — Exact per-Category resource attribute ranges and roll distributions.
4. **OQ-04** — Exact resource spawn algorithm and density per Zone.
5. **OQ-05** — Whether overharvesting should accelerate a Resource Pool's despawn, as an optional refinement to the base timer model.
6. **OQ-06** — Exact Resource Pool lifespan tiers/durations.
7. **OQ-07** — Exact quality-propagation formula combining resource, skill, tool, and experimentation inputs.
8. **OQ-08** — Exact Experimentation success/critical/failure probability curve.
9. **OQ-09** — Exact Factory/Harvester maintenance/upkeep cost curve.
10. **OQ-10** — Exact land allotment per character or account.
11. **OQ-11** — Single vs. dual currency model.
12. **OQ-12** — Numeric sink/faucet balance targets for the economy.
13. **OQ-13** — Multi-boxing / alt-account mitigation — policy question as much as a design one.
14. **OQ-14** — What reputation signal, if any, should exist beyond sales volume and maker tags before a formal rating system is designed.
15. **OQ-15** — Final Skill Discipline list per Profession Family (this pass's list is a skeleton).
16. **OQ-16** — Player Association tooling: permissions, shared storage, formal roles.
17. **OQ-17** — Final Service archetype list (this pass's list is not claimed exhaustive).
18. **OQ-18** — Structure zoning/density specifics and maintenance decay curve.
19. **OQ-19** — City governance model: elected vs. appointed Mayor, and what recall/impeachment mechanism exists.
20. **OQ-20** — City taxation base and who sets the rate.
21. **OQ-21** — City rank thresholds and the specific rewards each rank unlocks.
22. **OQ-22** — Whether a secondary, non-use-based Skill Point source should exist alongside use-based gain.
23. **OQ-23** — Exact wound/injury system mechanics for Combat.

---

## 29. Assumption Register

Every substantive assumption this pass introduced beyond what the brief explicitly stated.

1. **AS-01** — No prototype code was available or inspected this session. Nothing in this document has been marked `[PROTOTYPE]` as a result — this is an absence of information, not a design finding. *(Related: OQ-01)*
2. **AS-02** — This GDD is written genre/setting-neutral. No fictional world, lore, races, or genre commitment (sci-fi, fantasy, or otherwise) is assumed beyond "not the Star Wars IP." Generic terminology is used throughout deliberately. *(Related: OQ-02, WLD-006)*
3. **AS-03** — A single persistent world clock is assumed to exist, since resource spawn timing (`RES-004`) needs a shared clock, even though visible day/night presentation itself is left TBD. *(Related: WLD-004)*
4. **AS-04** — A single unified currency is assumed as the simplest default starting point for this pass. This is not a rejection of a dual-currency model on its merits — that alternative simply wasn't designed out. *(Related: ECN-001)*
5. **AS-05** — "Two to three concurrent professions" is assumed as the right target breadth for the Skill Point cap, based on the brief's general hybrid/sandbox framing. The brief does not state this multiplier explicitly — it is this designer's judgment call and should be treated as a starting hypothesis, not a derived requirement. *(Related: CHR-003, PRF-003)*
6. **AS-06** — Established Settlement NPC vendors are assumed to need to be *strictly inferior* to player-crafted goods, not merely different in kind. This is a stronger reading of "shouldn't out-compete the player economy" than the brief states in so many words. *(Related: ECN-002)*
7. **AS-07** — Elected (rather than appointed) City Mayors are assumed to be the more thematically appropriate default direction, given the "citizen of a living society" fantasy — while the actual governance mechanic is left TBD given real griefing risk. *(Related: SOC-003)*
8. **AS-08** — Full item tradeability with no soul-binding is assumed as the right default. The brief doesn't address soul-binding either way; this is this designer's recommendation against a common genre convention, flagged explicitly rather than adopted silently. *(Related: ITM-002)*
9. **AS-09** — "Research" is assumed to warrant a reserved-but-undesigned Profession Family and Service archetype slot, since the brief itself mentions research under both Professions (Section I) and Services (Section L) without elaborating — this pass preserves the placeholder rather than either fleshing it out or dropping it. *(Related: PRF-001, SVC-001)*

---

## 30. Systems Requiring Deeper Design

Consolidated "not designed yet, and shouldn't be mistaken for designed" list, per the brief's Section 20 requirement.

- Full combat ability/class/PvP-ruleset design (`CBT-003`)
- All numeric tuning across every system: Skill Point costs and cap, resource attribute ranges and spawn algorithm, Experimentation probability curves, Factory/Harvester upkeep curves, land allotments, currency sink/faucet targets
- Weather (`WLD-005`)
- Day/night cycle decision (`WLD-004`)
- A formal reputation/rating system, designed with anti-abuse as a first-class concern rather than an afterthought (extends `RTL-002`)
- The Research profession/service family (currently a reserved name only)
- Transportation vehicle specifics, which are genre-dependent (`WLD-006`)
- Full City governance mechanics: election vs. appointment, recall/impeachment, taxation base and authority (`SOC-003`, `SOC-004`)
- Exact scope of Established Settlement NPC vendor goods (`ECN-002`)
- New-player onboarding/tutorial flow (Section 4.1)
- Player Association formal tooling: permissions, shared storage, roles (`SOC-001`)
- Housing/structure customization and decoration depth (extends `BLD-001`)
- Genre, setting, tone, and lore — deliberately kept separate from systems design in this pass (`AS-02`)
- Multi-boxing / alt-account policy, which sits partly outside pure game design (`ECN-005`)
- Exact wound/injury mechanics (`CBT-002`)

---

## 31. Recommended Design Sequence

A proposed order for the design passes that should follow this one, sequenced by dependency (each step mostly depends on the ones before it):

1. **Genre, setting, and tone.** Colours the vocabulary of every system below it; cheap to decide now, expensive to retrofit once schematics and resource names exist.
2. **Resource System deep-dive.** Finalise the Category list, exact attribute ranges, and the spawn algorithm. Crafting and Economy both depend on this being real before they can be tuned.
3. **Crafting and Schematic deep-dive.** Experimentation math, the quality-propagation formula, and an initial schematic list large enough to support a first playable slice.
4. **Skill Point economy.** Exact costs, the Skill Point cap, and the full Skill Discipline list per Profession Family — the single number most of the rest of the game's balance depends on (Section 24, Section 26).
5. **Manufacturing numeric tuning.** Throughput caps, upkeep curves, land allotments.
6. **Economy simulation pass.** Sink/faucet modelling against real numbers from steps 2–5, target inflation/deflation curve.
7. **City governance detailed design.** Election/appointment mechanic, recall, taxation.
8. **Combat system design.** Abilities, the wound/injury model, PvP ruleset if any — deliberately sequenced after crafting/economy so combat's equipment and service dependencies (Section 16) have something real to depend on.
9. **Services deep-dive**, particularly Medical and Entertainment, since Combat depends on both.
10. **Buildings/structures numeric specification.** Footprints, costs, zoning, maintenance decay curves.
11. **Testing/simulation framework build-out**, validating Section 26's invariants against the real numbers set in steps 2–10.
12. **Content-authoring pipeline design**, formalising Section 22's data-driven expansion model into an actual content-creation workflow.

---

## 32. Glossary

| Term | Definition |
|---|---|
| **Region** | A large, thematically/geographically coherent area of the world; contains Zones. |
| **Zone** | A subdivision of a Region; the level at which Resource Pools and points of interest exist. |
| **Resource Pool** | A single spawned instance of a Specific Resource Class, with rolled attributes, a geographic footprint, and a limited lifespan. |
| **Resource Domain** | The broadest resource grouping (Mineral, Chemical, Gaseous, Aquatic, Flora, Fauna-derived). |
| **Resource Category** | A subdivision within a Domain that determines which attributes apply and which schematics accept it. |
| **Specific Resource Class** | A concrete, game-generated resource instance within a Category; the actual tradeable/harvestable unit. |
| **Attribute** | A measurable resource property (e.g. Conductivity, Malleability) drawn from the Universal Attribute Pool. |
| **Overall Quality (OQ)** | A derived, composite quality score for a resource, always present regardless of Category. |
| **Schematic** | The definition of a craftable item: required Resource Categories and thresholds, required Tool/Station, required Skill, and any sub-component Schematics. |
| **Experimentation** | The point-allocation minigame a crafter performs after assembly to push an item's sub-attributes toward higher quality, with success/failure risk. |
| **Exceptional Item** | A rare, critical-success crafting outcome; the top of the quality curve, reachable only through manual crafting. |
| **Provenance** | The permanent, unforgeable record of who crafted an item, visible on inspection. |
| **Production Template** | A manually-crafted, "locked in" configuration of a Schematic that a Factory can then reproduce at volume. |
| **Factory** | A player-placed structure that mass-produces a Production Template at fixed (non-Exceptional) quality. |
| **Harvester** | A player-placed structure that passively extracts resources from a Resource Pool over time. |
| **Skill Discipline** | A named branch of character development (e.g. a crafting or combat specialisation) made up of tiered Skill Boxes. |
| **Skill Box** | A single purchasable unit within a Skill Discipline, granting abilities, passive bonuses, or unlocks. |
| **Skill Point** | The spendable currency for Skill Boxes, earned primarily through use-based experience. |
| **Respecialization** | The (friction-bearing, not free) act of unlearning Skill Boxes to reallocate their Skill Points. |
| **Profession Family** | An organisational grouping of related Skill Disciplines (e.g. Medical, Crafting, Combat). |
| **Player Association** | A small-scale, low-overhead player organisation (a shop's staff, a small collective, a standing group). |
| **Player City** | A larger, place-based, governance-bearing settlement founded by players. |
| **Established Settlement** | A world-authored, always-present settlement offering baseline civic services; not player-foundable or destructible. |
| **Commerce Directory** | A searchable cross-town/region listing of player vendor stock, resolving back to physical shops rather than a disembodied auction house. |
| **Vitality / Stamina / Focus** | The three proposed character resource pools gating physical and mental sustained activity, spanning both combat and non-combat play. |
| **Wound** | The proposed injury-layer concept sitting beneath/beyond simple hit points, intended to keep Medical services necessary. |

---

*End of first-pass Master GDD. See Section 31 for what should happen next.*
