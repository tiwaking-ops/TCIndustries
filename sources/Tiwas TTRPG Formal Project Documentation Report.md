# Tiwas TTRPG Formal Project Documentation Report

## OpenCode Project Architecture Brief

### 1. Document Control

| Field | Value |
|---|---|
| Project | TCIndustries — Tiwakings Craftworld Industries |
| Artifact Type | OpenCode Project Architecture Description |
| Documentation System | Tiwas TTRPG Formal Project Documentation System |
| Report Status | PROPOSED — Implementation Planning Input |
| Report Version | 0.1 |
| Date | 2026-09-12 |
| Author | Perplexity — GPT-5.2 |
| Intended Consumer | OpenCode coding agent |
| Primary Purpose | Provide a controlled, auditable description of the project so OpenCode can generate a targeted software architecture |
| Governing Design Reference | `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` |
| Governing Authority Reference | `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` |
| Technical Design Status | Not yet canonical; implementation architecture must remain provisional |
| Approval Requirement | Human approval is required before any PROPOSED architecture becomes authoritative |

***

## 2. Instruction to OpenCode

Create a targeted project architecture for TCIndustries based on this report.

The architecture must support a prototype or staged implementation of a persistent, player-driven sandbox MMORPG focused on:

- Player-created economic and social identities.
- Resource discovery and regional scarcity.
- Interdependent professions.
- Player production, services, retail, and organisations.
- Differentiated crafted goods.
- Persistent ownership and provenance.
- Emergent gameplay rather than scripted hero progression.
- Controlled automation that does not make other players economically irrelevant.

Do not treat this report as permission to implement every design concept described below. The project is still at an early design stage. OpenCode must distinguish architecture required to support the project’s confirmed direction from mechanics that remain unresolved.

Where requirements are uncertain, implement extensible interfaces, data models, configuration boundaries, test seams, and placeholder services rather than inventing permanent mechanics.

***

## 3. Project Definition

### 3.1 Project Summary

TCIndustries is an original persistent sandbox MMORPG in which players participate in a living player-driven society.

The game is inspired by the systemic virtual-world philosophy of early massively multiplayer games, particularly dynamic resources, interdependent professions, player-made goods, commerce, housing, settlements, services, and emergent economic life. It must not reproduce protected intellectual property, names, factions, characters, lore, or setting material from external franchises.

The primary player fantasy is:

> Become somebody known in a persistent world.

Players may become resource prospectors, farmers, medical practitioners, entertainers, craft specialists, merchants, traders, architects, combatants, researchers, settlement organisers, manufacturers, transport operators, or combinations of these roles.

The game should make player identity meaningful through:

- Skills and specialisation.
- Product and service quality.
- Ownership.
- Business and organisation affiliation.
- Reputation.
- Geographic presence.
- Product provenance.
- Social relationships.
- Contributions to regional economies.

### 3.2 Core Gameplay Loop

The principal conceptual loop is:

> Discover resources → acquire and process materials → craft goods or provide services → specialise → establish reputation → participate in trade, society, and regional economies.

The software architecture must therefore be capable of representing relationships between:

- Characters.
- Skills.
- Resource instances.
- Resource attributes.
- Processing activities.
- Crafting activities.
- Items.
- Services.
- Businesses.
- Organisations.
- Markets.
- Locations.
- Ownership.
- Provenance.
- Reputation.
- Economic transactions.
- World-state changes.

### 3.3 Project Anti-Goals

The architecture must prevent or expose design drift toward:

- A conventional class-based MMORPG.
- A linear quest-driven theme park.
- A combat-only progression system.
- A single-player crafting game with multiplayer features.
- A pure spreadsheet economy detached from the world.
- An automation-first factory simulator.
- A world where NPCs replace meaningful player economic roles.
- A world where one character can efficiently perform every important role.
- A system that treats all goods as anonymous and interchangeable.
- A system that silently converts unresolved design proposals into fixed rules.

***

## 4. Authority and Status Model

All design and implementation artifacts must preserve explicit status and provenance.

### 4.1 Required Status Vocabulary

| Status | Meaning |
|---|---|
| `HUMAN-LOCKED` | Explicitly approved by the project owner or authoritative project brief |
| `LOCKED` | Authoritative project requirement, subject to the project’s authority rules |
| `DERIVED CONSTRAINT` | A genuine logical consequence of an authoritative principle |
| `PROPOSED` | Recommended direction not yet explicitly approved |
| `TBD` | Unresolved design decision |
| `PROTOTYPE` | Behaviour demonstrated in an implementation or experiment |
| `EVIDENCE` | Observed result that informs design but does not create authority |
| `DEFERRED` | Recognised work intentionally postponed |
| `ASSUMPTION` | Temporary working assumption |
| `HISTORICAL` | External or historical reference material |
| `NON-CANONICAL` | Explicitly excluded from the current project baseline |

### 4.2 Authority Rules

OpenCode must follow these rules:

1. Canonical document inclusion does not automatically mean human approval.
2. Repeated appearance across LLM-generated documents does not establish authority.
3. Prototype behaviour is evidence, not automatically a design requirement.
4. Historical inspiration is not a TCIndustries rule.
5. A derived constraint must be a genuine logical consequence, not merely a desirable implementation.
6. Unresolved details must remain `PROPOSED`, `TBD`, or `DEFERRED`.
7. OpenCode must not silently promote a status.
8. Any implementation decision that narrows a `TBD` requirement must be recorded as a proposal.
9. Technical architecture may support unresolved mechanics without deciding them.
10. Human approval is required before changing design authority or canonical status.

***

## 5. Authoritative Design Requirements

The following requirements are the current high-level architecture drivers.

### 5.1 Living Persistent World

**Status:** `LOCKED`

The project is intended to be a persistent multiplayer virtual world. The architecture should allow durable world state, including where appropriate:

- Player-owned structures.
- Businesses.
- Organisation assets.
- Resource discoveries.
- Market conditions.
- Product provenance.
- Local reputation.
- Regional development.
- Player-created institutions.

The exact persistence, sharding, reset, and deployment model remains `TBD`.

### 5.2 Citizen Rather Than Chosen Hero

**Status:** `LOCKED`

The architecture must not assume that every player is the centre of a scripted heroic narrative.

Systems must support meaningful identities based on:

- Economic contribution.
- Crafting expertise.
- Resource knowledge.
- Medical or social services.
- Trade.
- Construction.
- Organisation leadership.
- Transport.
- Combat.
- Research.
- Local reputation.

### 5.3 Player-Driven Economy

**Status:** `LOCKED`

Players should be major participants in production, distribution, consumption, and services.

The architecture must keep the following concerns separable:

- Item creation.
- Item ownership.
- Item transfer.
- Pricing.
- Market listing.
- Service provision.
- Resource extraction.
- Processing.
- NPC participation.
- Currency or exchange media.

NPC systems must be configurable so that they do not automatically replace meaningful player economic roles.

### 5.4 Interdependent Professions

**Status:** `LOCKED`

Important professions should depend upon one another through combinations of:

- Specialised skills.
- Distinct resources.
- Intermediate goods.
- Tools and facilities.
- Logistics.
- Geographic distribution.
- Information asymmetry.
- Reputation.
- Capacity limits.
- Service requirements.
- Organisational coordination.

The architecture must avoid hard-coding one particular interdependence mechanism before testing confirms it.

### 5.5 Emergent Gameplay

**Status:** `LOCKED`

The architecture should create conditions from which player-driven events can emerge.

Examples include:

- A rare resource discovery.
- A temporary regional supply shortage.
- A famous crafter’s product line.
- A transport bottleneck.
- A player-founded commercial district.
- A settlement depending on a specialist service.
- A change in demand caused by player activity.
- A conflict over access to valuable resources.

The system should provide observable state transitions and event records so that emergent outcomes can be analysed after they occur.

### 5.6 Discovery and Resource Events

**Status:** `LOCKED` at principle level; implementation details `PROPOSED` or `TBD`

Resources should be temporary, geographically distributed, variable in quality, and economically meaningful.

A resource discovery should be capable of creating an event rather than being merely an inventory update.

The architecture should support:

- Resource families and classifications.
- Individual resource instances.
- Geographic distribution.
- Quantity or lifecycle state.
- Attribute values.
- Discoverability.
- Extraction requirements.
- Processing options.
- Spawn and depletion events.
- Provenance from discovery through production.

Exact resource tables, spawn rates, durations, depletion mathematics, and attribute ranges are `TBD`.

### 5.7 Identity and Reputation

**Status:** `LOCKED` at principle level

The architecture should preserve meaningful associations between:

- Characters.
- Organisations.
- Businesses.
- Items.
- Services.
- Buildings.
- Resource discoveries.
- Transactions.
- Public records.

Reputation should be capable of emerging from recorded behaviour and economic history rather than relying only on an abstract numerical score.

### 5.8 Player-Created Society

**Status:** `LOCKED`

The architecture must allow future support for:

- Player organisations.
- Businesses.
- Shops.
- Commercial spaces.
- Settlements.
- Cities.
- Social venues.
- Trade networks.
- Civic structures.
- Player-created institutions.

Exact governance, zoning, taxation, city formation, and organisation mechanics remain `TBD`.

### 5.9 Non-Combat Careers

**Status:** `LOCKED`

Players must be able to pursue meaningful long-term careers without combat as their primary activity.

Potential service and production domains include:

- Prospecting.
- Extraction.
- Processing.
- Crafting.
- Manufacturing.
- Architecture.
- Retail.
- Medicine.
- Entertainment.
- Transport.
- Research.
- Construction.
- Commerce.

The exact progression model and economic balance remain `TBD`.

### 5.10 Flexible Skill-Based Professions

**Status:** `LOCKED`

The project uses a flexible skill-based profession architecture rather than rigid immutable classes.

The architecture must therefore avoid assuming:

- One permanent class per character.
- One linear progression path.
- One universal combat-centric level.
- Permanent inability to respecialise.

Exact skill trees, ranks, prerequisites, costs, caps, and respecialisation rules remain `TBD`.

### 5.11 Differentiated Crafted Goods

**Status:** `LOCKED` at principle level; implementation details `PROPOSED`

Crafted goods should be capable of differing meaningfully based on:

- Input resources.
- Resource attributes.
- Crafter skill.
- Tools.
- Facilities.
- Experimentation.
- Processing history.
- Provenance.
- Condition.
- Organisation or brand association.

The architecture must support item attributes and provenance without prematurely deciding exact quality tiers or crafting formulas.

### 5.12 Controlled Automation

**Status:** `LOCKED`

Automation must not eliminate the economic or social relevance of other players.

Any manufacturing or automated production architecture must leave meaningful roles for:

- Resource discoverers.
- Extractors.
- Processors.
- Skilled operators.
- Maintainers.
- Transport providers.
- Retailers.
- Researchers.
- Service providers.
- Organisation managers.

The architecture should make production capacity, input quality, maintenance, logistics, ownership, and human oversight configurable rather than permanently fixed.

***

## 6. Proposed Technical Architecture

This section describes an implementation direction for OpenCode. It is not a final technical design decision.

### 6.1 Recommended Architecture Style

Use a modular, data-driven, server-authoritative architecture with clear separation between:

- Domain rules.
- Persistent state.
- Simulation services.
- Networking or API transport.
- User-interface clients.
- Administrative and diagnostic tooling.
- Test and simulation harnesses.

The system should be suitable for staged development, beginning with a deterministic single-process simulation or vertical slice and expanding toward multiplayer services.

Do not begin with a full production MMORPG infrastructure unless explicitly requested. The first implementation should prioritise rule validation, reproducibility, observability, and rapid iteration.

### 6.2 Suggested Architectural Layers

#### Domain Layer

Contains authoritative game concepts and rules:

- Character.
- Skill.
- Profession.
- Resource.
- Resource instance.
- Item.
- Recipe or schematic.
- Crafting operation.
- Processing operation.
- Service.
- Business.
- Organisation.
- Location.
- Ownership.
- Provenance.
- Market listing.
- Transaction.
- Reputation record.
- World event.

The domain layer must not depend directly on a user interface, database vendor, or transport protocol.

#### Application Layer

Coordinates use cases such as:

- Discovering a resource.
- Claiming or extracting a resource.
- Processing a material.
- Crafting an item.
- Listing an item for sale.
- Purchasing an item.
- Providing a service.
- Registering a business.
- Joining an organisation.
- Recording a reputation event.
- Advancing world time.
- Resolving resource lifecycle events.

Application services should validate commands, invoke domain rules, and emit auditable events.

#### Persistence Layer

Stores durable state and event history.

The persistence design should support:

- Entity identifiers.
- Versioning.
- Timestamps.
- Ownership.
- Provenance.
- Status values.
- Relationships.
- Audit records.
- Reproducible simulation state.
- Migration of evolving schemas.

The first prototype may use a local relational database or structured files, but the persistence boundary must be replaceable.

#### Simulation Layer

Provides controlled world progression and systemic processing:

- Resource spawning.
- Resource depletion.
- Processing queues.
- Manufacturing queues.
- Market updates.
- Item condition changes, if later approved.
- Environmental or regional modifiers.
- NPC participation.
- Economic metrics.
- World event generation.

Simulation steps must be deterministic where practical and must support seeded test runs.

#### Interface Layer

The interface layer may initially expose:

- Command-line tools.
- REST or local HTTP endpoints.
- A lightweight administrative interface.
- Structured JSON inspection.
- Simulation reports.

A graphical game client should not be treated as a prerequisite for validating the core economy and interdependence model.

#### Analysis and Test Layer

Provide tools for:

- Scenario setup.
- Synthetic player populations.
- Economic simulation.
- Profession dependency testing.
- NPC substitution testing.
- Resource scarcity testing.
- Automation testing.
- Reproducibility.
- Metrics export.
- Failure-condition detection.

***

## 7. Suggested Domain Modules

OpenCode should consider the following modules. Module boundaries are proposals and may be revised after implementation experience.

| Module | Responsibility |
|---|---|
| `identity` | Characters, names, profiles, identity history |
| `skills` | Skills, domains, specialisation, training interfaces |
| `professions` | Derived profession identity and role classification |
| `resources` | Resource taxonomy, instances, attributes, lifecycle |
| `world` | Regions, locations, time, environmental state |
| `extraction` | Surveying, discovery, harvesting, depletion |
| `processing` | Conversion of resources into intermediate materials |
| `crafting` | Schematics, recipes, inputs, outputs, experimentation seams |
| `items` | Item identity, attributes, condition, stacking, provenance |
| `manufacturing` | Facilities, queues, capacity, automation constraints |
| `services` | Medical, entertainment, transport, construction, repair, and other services |
| `commerce` | Vendors, shops, listings, orders, pricing, transactions |
| `economy` | Currency abstractions, market metrics, supply and demand |
| `ownership` | Assets, transfer, permissions, custody, safeguards |
| `organisations` | Membership, collective identity, roles, assets |
| `reputation` | Behaviour records, provenance-linked reputation evidence |
| `events` | Domain events, world events, audit events |
| `simulation` | Time advancement, population simulation, scenario execution |
| `governance` | Status metadata, design decisions, change records |
| `observability` | Logs, metrics, traces, state inspection, reports |

OpenCode should not implement all modules as complete gameplay systems in the first milestone. It should establish clean boundaries and implement only the minimum required for the selected prototype slice.

***

## 8. Core Data and Domain Requirements

### 8.1 Entity Requirements

Every persistent entity should have, where applicable:

- Stable identifier.
- Entity type.
- Creation timestamp.
- Last-modified timestamp.
- Version or revision number.
- Current status.
- Ownership reference.
- Provenance reference.
- Source or creation event.
- Audit history.
- Extensible metadata.

### 8.2 Resource Instance Requirements

A resource instance should be capable of representing:

- Resource family.
- Resource class.
- Subtype.
- Geographic region.
- Discovery location.
- Discovery time.
- Lifecycle state.
- Remaining quantity or availability state.
- Attribute values.
- Extraction requirements.
- Processing compatibility.
- Rarity or scarcity metadata.
- Discovering character or organisation.
- Related world events.

Do not hard-code final resource names, attribute ranges, or numerical spawn rules.

### 8.3 Item Requirements

An item should be capable of representing:

- Item definition or type.
- Unique item identifier where required.
- Material inputs.
- Resource provenance.
- Crafter identity.
- Organisation or business association.
- Quality-related attributes.
- Condition or durability placeholder.
- Ownership.
- Transfer history.
- Creation event.
- Current location.
- Trade status.
- Stackability classification.

Generic commodity items may eventually support stacking, while distinctive crafted products should retain individual identity where useful for provenance and reputation.

### 8.4 Crafting Requirements

The crafting subsystem should separate:

- Recipe or schematic definition.
- Input validation.
- Skill validation.
- Tool and facility validation.
- Resource attribute effects.
- Experimentation interface.
- Output generation.
- Provenance creation.
- Failure handling.
- Economic cost.
- Time or queue handling.

Exact experimentation mechanics are unresolved. Implement an extension point rather than committing to a final point-allocation or success-critical-failure model.

### 8.5 Ownership Requirements

Ownership should be represented independently from item or building definitions.

The design must leave room for:

- Character ownership.
- Organisation ownership.
- Business custody.
- Shared ownership.
- Permissions.
- Transfer restrictions.
- Abandoned asset handling.
- Fraud prevention.
- Administrative recovery.

Do not assume that every asset is freely tradeable or permanently bound. Those policies remain `TBD`.

### 8.6 Provenance Requirements

Provenance should be treated as a first-class cross-domain concern.

Potential provenance links include:

- Resource discovery.
- Resource extraction.
- Processing.
- Crafting.
- Manufacturing.
- Ownership.
- Business listing.
- Service delivery.
- Organisation participation.

Provenance records must not expose private information without an explicit visibility policy. Visibility and privacy rules remain `TBD`.

***

## 9. Event and Audit Model

Use domain events or an equivalent auditable mechanism for important state changes.

Examples:

- `ResourceDiscovered`
- `ResourceSpawned`
- `ResourceDepleted`
- `ResourceExtracted`
- `MaterialProcessed`
- `ItemCrafted`
- `ItemTransferred`
- `ItemListed`
- `ItemPurchased`
- `ServiceProvided`
- `BusinessRegistered`
- `OrganisationCreated`
- `OwnershipChanged`
- `ReputationEvidenceRecorded`
- `WorldTimeAdvanced`
- `ManufacturingStarted`
- `ManufacturingCompleted`
- `ManufacturingInterrupted`

Each event should include, where relevant:

- Event identifier.
- Event type.
- Aggregate identifier.
- Actor identifier.
- Location.
- Timestamp.
- Previous state reference.
- New state reference or change data.
- Provenance references.
- Correlation identifier.
- Causation identifier.
- Schema version.

The event model should support later analysis of:

- Who discovered a resource.
- How material moved through the economy.
- Which player crafted an item.
- How a price changed.
- Which services were used.
- Whether NPCs substituted for players.
- How an organisation affected regional activity.

***

## 10. Prototype Scope Recommendation

### 10.1 First Vertical Slice

OpenCode should initially target a small deterministic economic prototype rather than a complete MMORPG.

Recommended minimum slice:

1. One region.
2. A small resource taxonomy.
3. Temporary resource instances.
4. Surveying and discovery.
5. Extraction.
6. Basic processing.
7. A small number of schematics.
8. Crafting with differentiated resource attributes.
9. Item provenance.
10. Player ownership.
11. Basic listings and transactions.
12. A small set of professions.
13. Synthetic NPC and player agents.
14. Metrics for interdependence and substitution.
15. Reproducible simulation scenarios.

### 10.2 Prototype Objective

The first prototype should answer whether the architecture can support the Economic Interdependence Core:

- Can different roles contribute meaningfully?
- Can resource differences create economic value?
- Can crafting produce differentiated outcomes?
- Can players depend on one another without artificial inconvenience?
- Can NPCs be prevented from replacing all meaningful roles?
- Can automation remain bounded?
- Can the simulation expose failure conditions?

The prototype is evidence only. Its behaviour must not automatically become canonical game design.

***

## 11. Testing and Validation Requirements

### 11.1 Required Test Categories

OpenCode should create a testing structure for:

- Unit tests for domain rules.
- Integration tests for economic workflows.
- Persistence tests.
- Event serialization tests.
- Deterministic simulation tests.
- Property-based tests for resource and crafting invariants.
- Scenario tests for player interdependence.
- Regression tests for status and authority metadata.
- Failure-condition tests for automation and NPC substitution.

### 11.2 Suggested Invariants

The following are implementation-oriented candidates, not final game rules:

- A resource instance cannot be extracted after its lifecycle has ended.
- An item cannot be created without valid inputs unless an explicit exception is configured.
- Provenance must not be lost when an item changes ownership.
- A transaction cannot transfer an item without valid ownership or permission.
- A process cannot consume more input than is available.
- A completed operation must produce an auditable event.
- A simulation run with the same seed and inputs should produce the same result.
- An NPC configuration must not bypass all player-dependent systems by default.
- Automated production must remain bounded by configured inputs, capacity, skill, maintenance, logistics, or other explicit constraints.
- A `TBD` design field must not be treated as an implicit default without documentation.

### 11.3 Economic Test Metrics

The simulation should be able to report metrics such as:

- Resource discovery concentration.
- Resource availability by region.
- Profession participation.
- Number of unique actors in a production chain.
- Percentage of outputs produced by players versus NPCs.
- Degree of automation.
- Market concentration.
- Item provenance depth.
- Service utilisation.
- Average transaction path length.
- Role substitution rate.
- Regional price variation.
- Production bottlenecks.
- Abandoned or unused systems.
- Dependency failure points.

Metrics must be treated as evidence for design review, not as automatic proof that the design is successful.

***

## 12. Data-Driven Design Requirements

The architecture should use data-driven definitions for systems likely to change during design iteration.

Prefer configurable or data-defined representations for:

- Resources.
- Resource attributes.
- Recipes.
- Skills.
- Processing rules.
- Facility capabilities.
- Service types.
- Item attributes.
- Regional modifiers.
- NPC participation.
- Automation limits.
- Market rules.
- Reputation evidence types.
- Event schemas.
- Scenario definitions.

Avoid burying unresolved design values in code.

Every configurable rule should identify:

- Its purpose.
- Its current status.
- Its source or decision record.
- Its version.
- Its default behaviour.
- Its test coverage.
- Its change impact.

***

## 13. Documentation and Repository Requirements

OpenCode should create a repository structure that keeps design, technical architecture, implementation, tests, and evidence separate.

A proposed structure is:

```text
/
├── README.md
├── AGENTS.md
├── docs/
│   ├── governance/
│   │   ├── authority-model.md
│   │   ├── status-vocabulary.md
│   │   └── decision-records/
│   ├── design/
│   │   ├── canonical/
│   │   ├── proposed/
│   │   ├── tbd/
│   │   └── non-canonical/
│   ├── technical/
│   │   ├── architecture/
│   │   ├── domain-model/
│   │   ├── persistence/
│   │   ├── events/
│   │   └── simulation/
│   ├── testing/
│   └── evidence/
├── src/
│   ├── domain/
│   ├── application/
│   ├── infrastructure/
│   ├── simulation/
│   └── interfaces/
├── tests/
│   ├── unit/
│   ├── integration/
│   ├── simulation/
│   └── scenarios/
├── data/
│   ├── resources/
│   ├── recipes/
│   ├── skills/
│   ├── regions/
│   └── scenarios/
├── tools/
│   ├── simulation/
│   ├── analysis/
│   └── migration/
└── reports/
    ├── simulation-results/
    ├── architecture-reviews/
    └── audit-results/
```

OpenCode may adapt this structure to the selected programming language and framework, but it must preserve the separation of concerns.

***

## 14. OpenCode Operating Rules

OpenCode must:

1. Read and follow `AGENTS.md` before modifying the repository.
2. Treat the current Master GDD and authority matrix as design governance references.
3. Preserve the distinction between design authority and implementation evidence.
4. Mark new architecture decisions as `PROPOSED` unless explicitly approved.
5. Avoid inventing numerical values for unresolved mechanics.
6. Avoid implementing historical or non-canonical proposals as if they were current requirements.
7. Prefer interfaces and configuration boundaries for unresolved systems.
8. Keep domain logic independent from UI and infrastructure.
9. Add tests with each meaningful domain feature.
10. Record architectural decisions in versioned markdown documents.
11. Make simulation runs reproducible.
12. Provide clear migration paths when schemas or rules change.
13. Never silently rewrite canonical design documentation.
14. Report assumptions, unresolved questions, and risks after each substantial implementation task.
15. Explain which requirements are implemented, stubbed, deferred, or unsupported.

***

## 15. Explicit Non-Canonical Boundaries

The following concepts must not be implemented as confirmed TCIndustries requirements unless separately approved:

- Factories being unable to produce exceptional-tier output.
- Specific vitality, stamina, or focus pools.
- A specific Skill Discipline or Skill Box architecture.
- Use-based skill-point acquisition.
- A fixed two-to-three-profession limit.
- Toxicity as a mandatory resource attribute.
- A selected time-limited or extraction-limited resource model.
- Harvesters as a mandatory installed-structure implementation.
- A specific experimentation point-allocation model.
- A specific experimentation success or failure model.
- A single unified currency as the default.
- A mandatory Commerce Directory.
- Full tradeability by default.
- Mandatory account or character soul-binding.
- Services being universally recurring or consumable.
- Any exact skill tree, level curve, combat system, city-governance system, or numerical economy model not separately approved.

***

## 16. Current Risks and Unresolved Questions

### 16.1 Design Risks

- The economy may become a spreadsheet optimisation exercise.
- NPCs may replace player services and production.
- Automation may make other players optional.
- Resource scarcity may create unhealthy monopolies.
- Interdependence may become artificial inconvenience.
- Provenance may create excessive tracking or privacy problems.
- Specialisation may become a permanent character trap.
- Regional economies may become tedious if transport friction is excessive.
- A single player may dominate too many economically important roles.
- The project may prematurely freeze numerical rules before evidence exists.

### 16.2 Technical Questions

The following questions remain unresolved and must be treated as `TBD`:

- Which programming language and runtime should be used?
- Which game engine, if any, will host the eventual client?
- Will the first prototype be a command-line simulation, web service, or engine project?
- What persistence technology is appropriate?
- What multiplayer networking model is required?
- What is the target concurrency?
- What is the world partitioning or sharding model?
- Which systems require transactional consistency?
- How will offline simulation and world time operate?
- What data must be private, public, or selectively visible?
- What is the final economic exchange model?
- How are organisations governed?
- How are cities formed and maintained?
- How are abandoned assets handled?
- How are exploits, fraud, duplication, and market manipulation mitigated?

OpenCode must not resolve these questions silently. It should identify the minimum decision needed for the current milestone and document alternatives.

***

## 17. Required OpenCode Deliverables

After analysing this report, OpenCode should produce:

1. A proposed repository architecture.
2. An `AGENTS.md` file containing the project-specific implementation rules.
3. A technical architecture document.
4. A domain model document.
5. A first-pass entity and event catalogue.
6. A dependency graph of proposed modules.
7. A minimal prototype implementation plan.
8. A test strategy.
9. A reproducible simulation strategy.
10. An assumptions and unresolved-decisions register.
11. A list of architecture decisions requiring human approval.
12. A milestone plan that begins with the smallest useful economic vertical slice.

Each deliverable must include:

- Document title.
- Document type.
- Version.
- Status.
- Author.
- Date.
- Scope.
- Authority or source references.
- Assumptions.
- Open questions.
- Change history.

***

## 18. Acceptance Criteria

The proposed architecture is acceptable for review only if it:

- Clearly separates canonical requirements from proposals.
- Supports resource discovery, processing, crafting, ownership, provenance, and commerce.
- Can represent multiple player roles without requiring rigid classes.
- Can support non-combat careers.
- Can record player and organisation identity.
- Can test economic interdependence.
- Can measure NPC substitution and automation.
- Supports deterministic simulation.
- Keeps unresolved rules configurable.
- Provides auditable state changes.
- Avoids premature numerical commitment.
- Does not introduce external intellectual property.
- Includes automated tests for core domain invariants.
- Documents all assumptions and architecture trade-offs.
- Does not claim that prototype behaviour is canonical design.

***

## 19. Human Approval Gate

Before the architecture is treated as an approved technical baseline, the project owner must review:

- The selected programming language and runtime.
- The proposed repository structure.
- The first prototype scope.
- The persistence approach.
- The event and provenance model.
- The simulation model.
- The treatment of NPCs and automation.
- The proposed domain boundaries.
- Any implementation choice that narrows a `TBD` design decision.
- Any change to the authority or status model.

Until that review occurs, all architecture decisions in this report remain:

> **PROPOSED — suitable for OpenCode planning and prototype construction, but not yet canonical.**

***

## 20. Source Control Note

This report is derived from the current TCIndustries project documentation, particularly:

- `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`
- `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`

The Master GDD identifies TCIndustries as a persistent player-driven sandbox MMORPG with dynamic resources, interdependent professions, player production, services, commerce, ownership, reputation, and emergent society. It also explicitly separates `LOCKED`, `DERIVED CONSTRAINT`, `PROPOSED`, `TBD`, `PROTOTYPE`, `DEFERRED`, `ASSUMPTION`, and `HISTORICAL` material. [pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws](https://pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws.com/projectfiles-prod/gateway-object-store/repos/files-86a01c8bc03d47db8256f30f54010452/downloads/6b4c61eb64e7fb5abfaa31067bf7c9a776bc6806/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Checksum-Mode=ENABLED&X-Amz-Credential=ASIA2F3EMEYE5VS6YDSL%2F20260912%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260912T065121Z&X-Amz-Expires=600&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEOL%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJGMEQCIHokNDC0Ikcp2z2dT9M%2BeO5OllFCFFkYphEuIGtmCSNkAiA9MjLwcv2Fc5QBm9g9213uqmbzOEJPHRPhRBoVYwRUCyqvBQiq%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAEaDDY5OTc1MzMwOTcwNSIMg9tlSEXUHeIxXqbjKoMF%2B5%2FUob0SYaYWY1R3DhMKhY6ffqKvaACzxO6MRunZdZA5bn6Ozw5KAGTub7Nwaa0i5QedFbmAmkKiKoPzFjDuhMjtJKCvbOz51wTq3u4V%2By%2F5nGbI%2F3FZvg9ffVpe0BPOgYWQX2x83L92KmNn3b%2BEp%2BSen5NYJgpH0dwV4oOF8oL2NSOENFjQI0XmsqR1zgKIjX8%2BnJ3NPnQz1P9CY8ltn9zfb0X76MzK9%2FnubgG4n65YO%2FwkA7Ow7CvUFNQtMZeY%2FP8i7zYefE227w8Mzz8FZOHmKavDhEZzfD0LRuoUVohlzOhMh%2FRyBPR7IwNoy3nr0Rpt%2F3w25yPAcpn9Oj6kLI5qTSo%2Fy6ogqH8Tax%2BYU077qlHcTbEtDiiQfOEpJ5WzR05ZlMgqI%2FsQdgyA2DWUZ5O2vkS4FTBDvCy2ynSs0uRtt5qXv%2F9ilDrtNP4hSStcnTTC%2FfTbypMdgUZpE8Fmo2c8bRDucvNKidYS59adKOh5Sxs1zFlRgrsEFgeDQ5%2BPpUWGzDn%2BmRvbdqE2iefiIEOWd%2BR5Y0mkRkH7Umxzw5kyUlOZ6ZSlofnTE61qTHATKmFW%2FPJeruWqQCQpHvh5ROygvUwRVRst2qgMpcQTuo4ei0RwEoYu659XXoKPiSr7A1gO5vnVc9KFpi%2FTdX%2FFfEdB9es5RwfPpPz2ekA0ELTN%2FPzaCLc%2BPjIwPmWr2cg3iqG2rbgTfaH0XIbQhbYoZEGmjif5ll56rv%2BFlDdPnPWnnEUGuebeRFmtERPn%2BPNlesozHBAK%2BZwzV9LuewlXmtFpc2DkaJ4WA1Ph%2BpvsYwvVCVgfp45SmRKZAjrQyEh%2F2CFosNp6v3yEvaUyosnl3TPoyDD81pLVBjqRAaxDP2zEpjQZWxGiw9E1%2FITwvMD0wakutt8sq4QQcYBYeI7npIxUHMpPm0Xl8HmUybIIcTlAo1%2BU8t5jweGsJ2DhLEhOVT8Xm7pmM7%2F79dFpMHRYF5BMk%2F%2BBIXp7N14rzycqMUNCxQRmXNVlorJRUEaX%2BDoI55TjYNV9UXvZTEApXWYq6IZR5iM2SydyWSe1Cxo%3D&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=576a6f3ab5d8ffdac0577027ff105be817cb9b4ecfd0a79b9ca76642d0a3b504)

The authority matrix establishes that canonical presence, repeated LLM proposals, prototype behaviour, and historical inspiration do not independently constitute human approval. It also identifies the Economic Interdependence Core as the next major validation stage before deeper mechanical specification.[2]