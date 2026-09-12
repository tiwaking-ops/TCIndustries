# TCIndustries — Economic Interdependence Core: Rigorous Design Investigation

**Document Type:** Design Investigation (Non-Canonical)  
**Version:** 1.0  
**Date:** 2026-08-25  
**Author Role:** Senior MMO Systems Architect / Virtual-Economy Designer  
**Author:** OpenAI GPT-5.6 Luna  
**Authority:** None. This document creates no new LOCKED design decisions. All recommendations are PROPOSED pending explicit human approval.  
**Controlling Canonical Reference:** `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`  
**Supporting References:**  
- `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md`  
- `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md`  
- `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`  
- `TCIndustries_Master_GDD_v1.0_Consolidated.md`  

**Status Discipline Applied:** CANONICAL | DERIVED | PROPOSED | EXTERNAL EVIDENCE | INFERENCE | HUMAN DECISION REQUIRED  

---

# 1. Executive Finding

**What the investigation reveals:**

Meaningful player interdependence in a persistent sandbox MMORPG cannot be reliably produced by any single mechanism (hard skill-cap, opportunity-cost alone, geographic friction, or reputation). It requires a *coupled* system in which at least three independent pressure axes operate simultaneously:

1. **Capability concentration** — a character cannot be elite at every economically important production role at once.
2. **Resource differentiation + temporality** — resources are not fungible commodities; quality, attributes, location, and lifecycle create temporary comparative advantages that reward specialised knowledge and networks.
3. **Quality / identity ceiling on automation** — factories and automated processes can scale volume of ordinary goods but cannot fully replicate the highest-value differentiated outputs or the social signalling of provenance and reputation.

When any one of these axes is removed or made trivial, the system collapses toward either universal self-sufficiency (one character or one organisation does everything) or pure spreadsheet optimisation (the only rational behaviour is mathematical minimisation of friction).

SWG demonstrated that dynamic resources + experimentation + factories + skill-box specialisation *can* produce famous specialists, resource rushes, and player-driven markets. It also demonstrated that the same systems can produce opaque spreadsheet optimisation, multi-account vertical integration, and eventual inflation/stagnation when sinks and monitoring fail. EVE demonstrates that complex multi-stage supply chains and organisational scale can sustain interdependence even with relatively weak per-character skill constraints, but at the cost of high barriers to entry and heavy reliance on player organisations.

**Critical design insight (INFERENCE):**  
The operational question is not “What should a character be forbidden from doing?” but “What combination of opportunity costs, quality ceilings, information asymmetries, capacity limits, and social/reputational payoffs makes voluntary specialisation the highest-value strategy for most rational players most of the time?”

A pure hard skill-cap is neither necessary nor sufficient. A pure opportunity-cost model without quality ceilings is fragile to multi-accounting and organisational scale. The viable region of the design space is hybrid.

This investigation develops three competing architectures that occupy different points in that space, stress-tests them, and recommends one candidate for further human consideration. Nothing in this document is LOCKED.

---

# 2. Canonical Boundaries

## 2.1 HUMAN-LOCKED Principles That Constrain the EIC

From `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` (human rulings applied 2026-08-25):

| ID | Statement | Status |
|---|---|---|
| VIS-001 | Living persistent player-driven virtual world | LOCKED |
| VIS-002 | Citizen rather than chosen hero | LOCKED |
| VIS-003 | Player-driven economy | LOCKED |
| VIS-004 | Interdependent professions | LOCKED |
| VIS-005 | Emergence over scripted experiences | LOCKED |
| VIS-006 | Anti-goals (no infinite personal-factory, no NPC-dominated economy, no efficient universal generalist, etc.) | LOCKED |
| PIL-001 | Discovery and the Gold Rush (resources temporary, geographically distributed, variable quality, economically meaningful) | LOCKED |
| PIL-002 | Identity and Reputation | LOCKED |
| PIL-003 | Interdependence (no major profession fully self-sufficient at high effectiveness; opportunity rather than coercive inconvenience) | LOCKED |
| PIL-004 | Player-Created Society | LOCKED |
| LOOP-002 | Meaningful long-term non-combat career | LOCKED |
| LOOP-003 | Progression expands choices, not one linear ladder | LOCKED |
| PROF-001 | Flexible skill-based profession architecture | LOCKED |
| PROF-004 | Respecialisation must be possible | LOCKED |
| RES-002 | Resources exist to create discovery/scarcity/economic events | LOCKED |
| CRFT-001 | Crafting must produce differentiated products (principle only) | HUMAN-LOCKED |
| MFG-001 | Automation must not eliminate player relevance | LOCKED |
| ECO-001 | Economy primarily player-driven | LOCKED |
| MFG-004 | No infinite self-sufficient personal-factory automation | DERIVED CONSTRAINT |
| ECO-003 | No pure spreadsheet optimisation | DERIVED CONSTRAINT |

**Explicitly confirmed HUMAN-LOCKED (principles only):** WRLD-001, PLR-001, PLR-003, CRFT-001, SAFE-001.

## 2.2 Explicitly Non-Canonical (Must Not Be Resurrected as Authority)

Per Authority & Provenance Reconciliation Matrix §5 and v1.1.1:

- Specific skill-point / skill-box architecture
- Specific two-to-three concurrent profession target
- Specific resource lifecycle model (time-limited vs extraction-limited)
- Specific experimentation-point model
- Factories permanently unable to produce Exceptional-quality output
- Particular currency model
- Particular service-consumption model
- Toxicity attribute as defined
- Use-based skill acquisition
- Vitality/Stamina/Focus pools

These remain historical or earlier proposals only.

## 2.3 Unresolved Questions Relevant to EIC

| ID | Question | EIC Relevance |
|---|---|---|
| OQ-001 | Specialisation budget / skill-cap model | Core |
| OQ-002 | Resource spawn/depletion/lifecycle | Core |
| OQ-003 | Resource attributes ↔ product outcomes | Core |
| OQ-004 | Crafting experimentation model | Core |
| OQ-005 | Manufacturing throughput, maintenance, ownership | Core |
| OQ-009 | Currency and stabilisation | Dependency (defer detailed design) |
| OQ-010 | Durability, decay, repair | Strong dependency (demand generation) |
| OQ-011 | Respecialisation | Core (interacts with specialisation) |
| OQ-012 | Multi-accounting / economic exploitation | Core (adversarial) |
| OQ-015 | Scope for early playable/persistent alpha | Boundary |
| OQ-016 | Vendor/retail mechanics | Strong dependency |
| OQ-017 | Provenance depth and visibility | Core (identity) |

**Not in scope for this investigation:** full combat design, city/governance rules, complete reputation system, setting/lore, numerical balance values.

---

# 3. Problem Definition

## 3.1 What “Economic Interdependence” Must Mean Operationally

**Definition (PROPOSED for this investigation):**  
Economic interdependence exists when, for a significant fraction of high-value economic activity, a rational player maximises expected utility by transacting with (or organising with) other specialised players rather than performing all stages themselves or relying solely on NPC substitutes.

This is **not** the same as:

| Concept | Distinction |
|---|---|
| Dependency | One player *cannot* complete an activity without another. Can be coercive. |
| Specialisation | Players choose different focus areas. Does not automatically produce trade. |
| Cooperation | Voluntary joint activity. May be temporary or altruistic. |
| Transaction | Exchange of goods/services for value. Can occur under monopoly or friction. |
| Inconvenience | Friction that makes self-sufficiency annoying but still optimal. |
| Coercion | Forced interaction (mandatory grouping, hard gates). Violates PIL-003 “opportunity rather than coercive inconvenience”. |
| Economic necessity | Structural conditions under which trade is the dominant strategy. |
| Voluntary economic advantage | Trade is preferred because it yields higher quality, higher volume, lower opportunity cost, better reputation, or better social outcomes. |

**Target state (derived from LOCKED pillars):**  
Players *prefer* to specialise and trade because the returns to specialisation (quality, volume, reputation, time efficiency, access to scarce resources) exceed the returns to generalism for most characters most of the time. Self-sufficiency remains possible at a meaningful but inferior quality/volume/identity tier.

## 3.2 The Critical Design Question

> What must a single character be structurally unable to do *simultaneously* (or only able to do at decisive disadvantage) for TCIndustries’ economy and society to remain meaningfully interdependent?

Possible axes (not mutually exclusive):

- Hard specialisation budget (skill points, concurrent mastery slots)
- Soft opportunity-cost / time-capacity constraints
- Knowledge and information asymmetry
- Resource access and geographic constraints
- Equipment / facility / throughput constraints
- Quality ceilings on automated or non-specialist production
- Reputation and provenance payoffs
- Organisational scale requirements
- Logistics and intermediate goods chains
- Service dependencies (repair, medical, entertainment, transport)

The investigation treats “hard skill-cap” as one candidate mechanism among several, not the default answer.

---

# 4. External Evidence

## 4.1 Star Wars Galaxies (Case Study — Not Authority)

**EXTERNAL EVIDENCE**

**What genuinely produced valuable interdependence:**
- Temporary, statistically variable resource spawns with unique named instances and attribute vectors (OQ, CD, SR, UT, MA, etc.). Exceptional resources were time-limited events.
- Experimentation system that allowed skilled crafters to push product attributes using resource quality + skill points. Hand-crafted items could exceed factory output in quality ceiling.
- Factories that scaled volume of ordinary goods but required manufacturing schematics derived from player-crafted prototypes; did not fully replace elite hand-crafting for top-tier items.
- Skill-box architecture with meaningful opportunity cost: a character could not master every high-value crafting profession simultaneously.
- Provenance (crafter name on items) + player vendors created recognisable specialist identities.
- Surveying created information asymmetry and economic value for prospectors.

**What created mere friction or failure modes:**
- Extreme opacity of optimal resource combinations → spreadsheet dominance and high knowledge barrier for new players.
- Multi-accounting and organisational vertical integration could approximate self-sufficiency.
- Factories + harvesters allowed significant automation once the player had capital and lots.
- Post-launch changes (NGE, reduced sinks, free character slots, reduced maintenance) destroyed economic pressure and led to inflation and city viability collapse.
- “Hologrind” and other singular goals collapsed lateral interdependence into a single optimisation target.

**Lessons TCIndustries should adopt (INFERENCE):**
- Dynamic, temporary, attribute-bearing resources are a powerful interdependence engine.
- Quality ceilings that differ between hand-crafted/specialist and automated production preserve specialist value.
- Provenance and recognisable specialists create identity payoffs that reinforce economic specialisation.
- Continuous economic monitoring and sinks are non-optional.

**Lessons TCIndustries should explicitly reject or improve:**
- Opacity that forces spreadsheet optimisation as the dominant strategy (conflicts with ECO-003).
- Permanent character traps (conflicts with PROF-004).
- Unmonitored inflation and removal of maintenance/decay pressure.
- Treating factories as pure volume tools without sufficient input, maintenance, or quality constraints.

## 4.2 EVE Online (Comparative)

**EXTERNAL EVIDENCE**

- Multi-stage production chains (raw → intermediate → component → finished) with planetary industry, reactions, manufacturing, invention.
- Weak per-character hard skill limits relative to organisational scale; alts and corporations routinely achieve vertical integration.
- Geographic and logistical friction (hauling, security space, structure vulnerability) create real costs.
- Destruction (PvP, structure loss) is a major demand generator and wealth sink.
- Markets are player-driven but fragmented; market power and manipulation are real and accepted as part of the sandbox.
- Specialisation is often organisational or alt-based rather than single-character.

**Transferability to TCIndustries:**
- Multi-stage intermediate goods are powerful for interdependence.
- Pure reliance on organisational scale and alts risks making the single-character experience secondary and raising barriers for new/solo players (conflicts with accessibility and identity pillars).
- Destruction as a primary sink is powerful but TCIndustries must not force combat as the only viable sink (LOOP-002, CMBT-003 derived).

## 4.3 Other Relevant Evidence

- Albion Online: player-crafted everything + full-loot risk zones create continuous demand; city markets create regional price differences.
- Ryzom: quality capped by lowest ingredient + material grades creates gatherer–crafter interdependence.
- General virtual-economy research: time is the fundamental scarce resource; opportunity cost is the correct valuation lens; MIMO (multiple independent money objects) and weak sinks produce inflation; social hierarchy and networking affect real-money valuation of virtual goods.
- Common failure pattern: systems that rely only on friction or inconvenience are eventually optimised around or tolerated; systems that rely only on hard gates feel coercive.

## 4.4 Economic Concepts Applied

- **Opportunity cost** is the primary rational driver of specialisation.
- **Comparative advantage** (not absolute advantage) is what makes trade rational even when one player is better at everything.
- **Transaction costs** (search, trust, logistics, information) can destroy gains from trade if too high.
- **Network effects** and reputation can lock in specialist advantages.
- **Market power** (temporary monopolies from rare resources or elite skills) can be healthy emergence; permanent structural monopolies that exclude new entrants are destructive.
- **Substitution** (NPC or generalist alternatives) sets a price ceiling on specialist services; too high a ceiling and specialists have no market; too low and specialists are irrelevant.

---

# 5. System Requirements

## 5.1 Canonical Requirements (must be satisfied)

From LOCKED principles:

1. Player-driven production and distribution of economically meaningful goods and services.
2. No major profession category fully self-sufficient at high effectiveness.
3. Resources create discovery, scarcity, regional trade, and events.
4. Crafted products are differentiated (not pure commodities).
5. Automation does not eliminate player relevance.
6. Meaningful non-combat careers are viable.
7. Interdependence creates opportunity, not coercive inconvenience.
8. Respecialisation remains possible.
9. One character cannot efficiently perform every economically important role.
10. No pure spreadsheet optimisation as the dominant strategy.
11. Emergence of player institutions, specialists, and economic events.

## 5.2 Derived Constraints

- MFG-004: No infinite self-sufficient personal-factory loop.
- ECO-003: Design must resist collapse into pure mathematical optimisation.

## 5.3 Proposed Requirements for the EIC (this investigation)

These are not LOCKED; they are the working requirements used to evaluate architectures:

R1. A rational solo generalist can achieve baseline competence and modest economic participation but cannot match specialist quality/volume/reputation at the high end.  
R2. Resource attributes and lifecycle create temporary comparative advantages that reward specialised knowledge and networks.  
R3. Experimentation / differentiation creates a meaningful quality spectrum that is not fully captured by a single “best” resource.  
R4. Manufacturing scales ordinary goods but cannot fully close the quality or identity gap to elite specialists.  
R5. NPC substitutes provide a floor (tutorial, convenience, crisis buffer) but not a ceiling that makes players irrelevant.  
R6. Demand is sustained by a combination of consumption, decay/repair, combat attrition, construction, services, and status signalling — not solely by arbitrary destruction.  
R7. Transaction costs are low enough that specialisation is attractive, high enough that random self-sufficiency is not free.  
R8. Multi-accounting and organisational vertical integration remain possible but face increasing coordination, capital, and quality ceilings that preserve space for independent specialists.  
R9. New players can enter the economy at multiple points (gathering, basic crafting, services, retail, logistics) without requiring mastery of the entire chain.  
R10. The system is simulatable with a modest agent set before full implementation.

---

# 6. Competing Architecture A — “Hard Concentration + Quality Ceiling”

**Core Principle**  
Interdependence is produced by a hard (or very firm) specialisation budget that prevents simultaneous elite mastery of multiple high-value production domains, combined with a quality ceiling on non-specialist and automated production that makes specialist output strictly preferable for high-value goods.

### Character Specialisation
- Characters have a limited “mastery capacity” (exact form TBD: concurrent elite slots, skill-point budget with diminishing returns past certain depth, or domain concentration).  
- Broad competence is possible; elite performance in 1–2 production domains is the practical maximum for most characters.  
- Attempting to generalise results in mediocre quality across the board and loss of reputation signalling.  
- Respecialisation is allowed at meaningful cost (time, economic, temporary capability loss) but does not erase provenance or identity history.

### Resource Economy
- Resources are temporary, named, attribute-bearing instances with regional and temporal variation (aligned with PIL-001).  
- Attribute vectors map non-linearly to product outcomes; “best overall” resource does not dominate every recipe.  
- Exceptional spawns create rushes and temporary local advantages.  
- Information (survey data) has economic value.

### Crafting
- Specialist skill depth + resource attributes + experimentation produce a quality spectrum.  
- Two specialists using different resource combinations and experimentation choices produce recognisably different products.  
- Opaque pure-math optimisation is resisted by multi-dimensional attributes and secondary properties (aesthetics, secondary effects, provenance premiums).

### Manufacturing
- Factories scale volume of ordinary and good-tier items.  
- Highest quality tier and certain identity-bearing properties require specialist hand-crafting or specialist oversight that cannot be fully automated.  
- Input quality still matters; factories do not magically upgrade poor resources.

### NPC Substitution
- NPCs provide basic tools, tutorial goods, emergency repair, and low-tier commodities at a price/quality floor.  
- They do not supply elite differentiated goods or high-reputation services.

### Demand
- Sustained by decay/repair, combat attrition, construction, consumables, status goods, and service consumption.  
- Exceptional goods command premiums via performance *and* identity/provenance.

### Interdependence Chain (example)
Prospector (info + access) → Extractor → Processor → Specialist Crafter (elite tier) → Manufacturer (volume of components or ordinary finished goods) → Retailer / Vendor → Consumer (combat, construction, status, service).  
Service providers (repair, medical, entertainment, transport) sit orthogonal to the goods chain and create recurring demand.

### Economic Value of Paying Another Player
Specialist quality and provenance cannot be matched by the generalist or by factories at the high end. Time spent becoming elite in one domain has higher return than spreading thin. Reputation compounds.

### Emergence Targets
Resource rushes, famous specialists, temporary local monopolies on exceptional resources, specialist recruitment by organisations, regional price differences, commercial hubs around high-quality production clusters.

### Failure Modes
- If the specialisation budget is too tight → coercive feeling, reduced hybrid viability.  
- If quality ceiling on factories is too low → factories become irrelevant; if too high → specialists become optional.  
- Spreadsheet optimisation of the attribute mapping.  
- Multi-accounting to circumvent the budget (one account = one elite domain).  
- New-player exclusion if elite tiers are the only economically relevant tier.

### Exploit Analysis
Multi-accounting is the primary attack: one player runs an elite weaponsmith, an elite armoursmith, an elite resource specialist, etc. Mitigation must come from account-level constraints, capital/facility limits, logistics, or quality ceilings that still require real attention. Organisational vertical integration is expected and healthy up to a point; the design must leave residual demand for independent specialists.

---

# 7. Competing Architecture B — “Opportunity-Cost + Capacity + Multi-Stage Chain”

**Core Principle**  
No hard per-character mastery limit. Interdependence arises from time/capacity constraints, multi-stage intermediate goods with specialised facilities, geographic and logistical friction, and quality that improves with focused practice and specialised tools/facilities. A generalist can do everything; they simply cannot do it *well or at scale* simultaneously.

### Character Specialisation
- Skills improve with use and investment; deep specialisation yields higher throughput, higher success rates, better experimentation outcomes, and access to advanced facilities.  
- No hard concurrent limit.  
- The constraint is real-time attention, facility ownership, and the opportunity cost of not specialising.  
- A generalist produces lower-quality, lower-volume output and has weaker reputation signals.

### Resource Economy
- Same dynamic resource philosophy as A, but with stronger emphasis on processing stages that require specialised stations or knowledge.  
- Intermediate goods (refined materials, components) become tradeable choke points.

### Crafting
- Differentiation still exists via resources + skill depth + experimentation.  
- Because there is no hard budget, the quality gap between specialist and generalist must be large enough that the opportunity cost of not specialising is decisive.

### Manufacturing
- Factories and processing plants require specialised setup, maintenance, and often ongoing specialist attention for optimal output.  
- Closed-loop self-sufficiency is possible but capital- and attention-intensive; most efficient path is still network specialisation.

### NPC Substitution
- Similar floor function as A. Slightly higher risk that NPCs + generalist play can cover more of the economy if the specialist premium is not large.

### Demand
- Same sources, with stronger reliance on intermediate-goods markets and logistics services.

### Interdependence Chain
More stages and more facility specialisation: Raw → Primary Processing (specialised) → Component Crafting → Assembly / Manufacturing → Distribution → Consumption. Logistics and facility maintenance create additional service demand.

### Economic Value of Paying Another Player
Time is the scarce resource. A specialist’s higher throughput and quality per unit time makes buying cheaper than self-producing at the same quality. Reputation and facility access reinforce this.

### Emergence Targets
Similar to A, plus stronger organisational production networks and logistics specialists. Higher risk of pure industrial organisations dominating.

### Failure Modes
- If the quality/throughput gap is too small, generalists and multi-account players dominate.  
- Vertical integration by wealthy organisations becomes the dominant strategy.  
- New players face high capital barriers to meaningful participation.  
- Without hard limits, the system can drift toward “everyone does a bit of everything and buys the rest,” reducing identity specialisation.

### Exploit Analysis
Multi-accounting and organisational scale are more powerful than in A because there is no hard character-level brake. The design relies entirely on attention costs, capital costs, and quality gradients. Botting and automation risk is higher.

---

# 8. Competing Architecture C — “Hybrid Concentration + Soft Capacity + Quality/Identity Ceilings”

**Core Principle**  
A moderate specialisation budget (firm but not extreme) is combined with capacity/facility constraints, multi-stage goods, dynamic resources, and strong quality + identity ceilings on automated and non-specialist production. The hard budget prevents the most extreme generalism; the other systems do the heavy lifting of making specialisation valuable and automation incomplete.

### Character Specialisation
- A moderate concurrent-mastery or concentration budget exists (exact form TBD; deliberately left as a tuning variable).  
- It is set so that a character can be elite in one primary production domain and competent in supporting domains, but cannot be elite across the full set of high-value economic roles.  
- Soft capacity limits (attention, facility slots, maintenance load) reinforce the budget.  
- Respecialisation remains possible with cost.

### Resource Economy
- Dynamic, temporary, attribute-bearing resources (PIL-001 aligned).  
- Attribute-to-outcome mapping is multi-dimensional and partially opaque by design (resists pure single-metric optimisation while remaining learnable).  
- Exceptional resources create time-limited comparative advantages.

### Crafting
- Specialist depth + resources + experimentation create a quality spectrum.  
- Provenance and secondary properties (aesthetic, narrative, set bonuses, service synergies) create non-purely-mathematical value.  
- Elite tier requires specialist skill; factories and generalists are capped below it.

### Manufacturing
- Factories scale ordinary and mid-tier goods efficiently.  
- Elite tier and certain identity-bearing properties require specialist intervention or cannot be factory-produced at all.  
- Maintenance, input quality, logistics, and ownership constraints prevent closed-loop personal empires from being the dominant strategy.

### NPC Substitution
- Clear floor: basic goods, tutorial support, emergency services.  
- Explicitly does not supply elite differentiated goods or high-reputation services. Price and quality are set so that they do not undercut the specialist market under normal conditions.

### Demand
- Composite: durability/decay/repair, combat attrition, construction and housing, consumables, status/identity goods, recurring services.  
- No single sink is load-bearing; the portfolio provides resilience.

### Interdependence Chain
Prospector / Surveyor → Extractor → Processor → Specialist Component Crafter → Manufacturer (volume) / Elite Finisher → Retailer → Consumer.  
Orthogonal service layer (repair, medical, entertainment, transport, architecture) creates continuous demand independent of goods volume.

### Economic Value of Paying Another Player
- Quality gap at the elite tier.  
- Opportunity cost of spreading concentration budget.  
- Reputation and provenance premiums.  
- Throughput and facility advantages of specialists.  
- Information and network advantages from focused play.

### Emergence Targets
Resource rushes and temporary camps, famous named specialists and brands, regional commercial specialisation, specialist recruitment, temporary resource monopolies that break when the spawn ends, industrial organisations that still buy elite finishing work, local price and style differences.

### Failure Modes
- Budget set too tight → feels coercive.  
- Budget set too loose → collapses toward Architecture B failure modes.  
- Quality ceiling on factories set incorrectly.  
- Attribute system becomes pure spreadsheet.  
- Multi-accounting still approximates broader coverage.  
- Transaction costs (search, trust, logistics) remain high enough that players prefer inferior self-supply.

### Exploit Analysis
Multi-accounting remains the leading threat but is partially blunted by the concentration budget per character plus facility/capital/logistics costs. Organisational vertical integration is expected; the residual elite-finishing and reputation markets preserve independent specialist value. Market manipulation and hoarding of exceptional resources are possible and should be treated as emergent content provided they are temporary (tied to resource lifecycle).

---

# 9. Comparative Evaluation

| Criterion | Model A (Hard Concentration + Quality Ceiling) | Model B (Opportunity-Cost + Capacity + Chain) | Model C (Hybrid) |
|---|---|---|---|
| Real interdependence | High (structural) | Medium–High (depends on gap size) | High |
| Player freedom | Lower (budget constraint) | Highest | Medium–High |
| Solo viability | Medium (can be competent generalist) | High | Medium–High |
| Specialist value | High | Depends on gap | High |
| Anti-automation resilience | High (quality ceiling) | Medium | High |
| Alt-account resilience | Medium (alts each take one elite slot) | Low | Medium–High |
| Resource economy quality | High | High | High |
| Crafting depth | High | High | High |
| Manufacturing viability | High (volume role clear) | High | High |
| New-player accessibility | Medium (risk of elite-only relevance) | Medium (capital barriers) | Higher (multiple entry points) |
| Economic emergence | High | High (organisational) | High |
| Risk of coercive inconvenience | Medium–High | Low | Medium |
| Risk of spreadsheet optimisation | Medium | Medium–High | Medium (mitigated by multi-dimensional value) |
| Economic stability | Medium | Medium | Higher (portfolio of mechanisms) |
| Design complexity | Medium | Medium–High | Higher |
| Simulation feasibility | Good | Good | Good |
| Fit with TCIndustries pillars | Strong on interdependence & identity; weaker on freedom | Strong on freedom & emergence; weaker on structural interdependence | Best overall fit |

**Judgements (INFERENCE):**

- **Strongest overall:** Model C (Hybrid). Best balance of structural pressure and freedom; multiple reinforcing mechanisms reduce single-point failure.  
- **Weakest for TCIndustries goals:** Model B. Too dependent on gap size and attention costs; vulnerable to multi-accounting and organisational closure.  
- **Most elegant:** Model A. Cleanest single-principle story.  
- **Most robust:** Model C.  
- **Highest-risk:** Model B (collapse to self-sufficiency or org dominance).  
- **Easiest to prototype:** Model A (fewer interacting soft systems).  
- **Most likely to produce the intended social economy:** Model C.

Trade-off summary: A sacrifices some freedom and hybrid viability for clarity and structural force. B maximises freedom but places almost all burden on economic incentives and quality gradients. C accepts higher design and tuning complexity in exchange for resilience and better alignment with the full set of LOCKED pillars.

---

# 10. Adversarial Stress Test

Leading architectures (A and C) under expert optimisation:

| Attack | Model A Response | Model C Response | Residual Risk |
|---|---|---|---|
| Multi-accounting (one elite domain per character) | Budget applies per character; capital, facilities, logistics still scale with number of alts | Same + soft capacity and facility limits add friction | Still possible; cost rises with scale |
| Organisational vertical integration | Elite finishing still prefers specialists; reputation markets remain | Same; multi-stage + quality ceilings leave residual demand | Healthy market power possible; permanent exclusion of independents is the failure mode to watch |
| Resource monopolisation | Temporary by design (lifecycle); information markets and multiple regions mitigate | Same | Acceptable emergence if temporary |
| Best-resource monoculture | Multi-dimensional attributes + secondary properties resist single ranking | Same + provenance/identity value | Requires careful attribute design |
| Factory dominance | Quality ceiling prevents factories from supplying elite tier | Same | Critical tuning parameter |
| Spreadsheet optimisation | Multi-dimensional + partial opacity + identity value | Same | Never fully eliminable; goal is to keep it from being the *only* viable strategy |
| New-player exclusion | Risk if only elite tier matters economically | Multiple entry points (gathering, basic crafting, services, retail) | Explicit design requirement |
| Transaction-cost avoidance (self-supply of inferior goods) | Quality gap must be large enough that inferior self-supply is not preferred for important uses | Same | Core test of the quality ceiling |
| Botting / AFK automation | Maintenance, attention, and quality ceilings limit pure AFK empires | Same | Continuous monitoring required |

**Conclusion of stress test:** Neither A nor C is immune. C’s multiple overlapping mechanisms give it more residual resilience when any single mechanism is partially circumvented. Both require ongoing economic observation and the willingness to adjust ceilings and budgets post-launch.

---

# 11. Recommended Candidate

**Recommendation: Architecture C — Hybrid Concentration + Soft Capacity + Quality/Identity Ceilings**

**Why it wins:**
- Satisfies the structural requirement that one character cannot efficiently master every high-value role (addresses VIS-006 and PIL-003) without relying on an extreme hard cap that would feel coercive.
- Preserves meaningful freedom and hybrid identities (PROF-001, LOOP-003, player fantasy).
- Uses dynamic resources, multi-stage goods, and quality ceilings as co-equal engines rather than treating specialisation budget as the sole load-bearing wall.
- Creates natural space for both independent specialists (identity/reputation) and industrial organisations (scale).
- Aligns with the explicit project preference for opportunity over coercive inconvenience.
- Multiple mechanisms reduce the chance that a single design error or player optimisation destroys interdependence.

**What it sacrifices:**
- Greater design and tuning complexity than a pure hard-cap model.
- Less pure elegance than Model A.
- Requires careful setting of the concentration budget (too tight → coercion; too loose → Model B failures).

**Key assumptions it relies on:**
- A quality/identity ceiling on factories and non-specialist production can be made meaningful and stable.
- Resource attribute space can be designed to resist pure single-metric optimisation while remaining learnable.
- Decay/repair + other sinks can sustain demand without forcing combat.
- Transaction costs can be kept in a productive band.
- Multi-accounting can be made expensive enough (capital, logistics, attention, account policy) that it does not fully neutralise the concentration budget.

**What remains uncertain / must not yet be locked:**
- Exact form of the specialisation budget (points, slots, concentration curves, diminishing returns).
- Exact quality ceiling rules for manufacturing.
- Exact resource lifecycle mathematics.
- Exact experimentation model.
- Exact attribute set and mapping functions.
- Numerical values of any kind.
- Full multi-accounting policy.
- Detailed currency and stabilisation design.

**Status of recommendation:** PROPOSED. Requires explicit human approval before becoming design baseline.

---

# 12. Candidate EIC Architecture (Recommended)

**Cross-system flow (PROPOSED):**

```
Character Concentration Budget
        ↕
Dynamic Attribute-Bearing Resources (temporary, regional, discoverable)
        ↕
Processing / Intermediate Goods (specialised facilities & knowledge)
        ↕
Crafting Differentiation (skill depth + resources + experimentation)
        ↕
Manufacturing (volume of ordinary/mid-tier; quality/identity ceiling)
        ↕
Distribution & Retail (player vendors, provenance visibility)
        ↕
Demand Generation (decay/repair, attrition, construction, consumables, status, services)
        ↕
Reputation & Identity Feedback (famous specialists, brands, provenance premiums)
        ↕
Economic Value Signals → further specialisation and organisational formation
```

**NPC role:** Floor provider only.  
**Automation role:** Scale ordinary production; never fully close the elite quality or identity gap.  
**Specialist role:** Elite quality, experimentation outcomes, provenance, and reputation-bearing production and services.  
**Organisation role:** Capital, logistics, multi-stage coordination, volume — still purchases or recruits elite finishing and specialist services.

This architecture is the minimum coupled system judged necessary. Removing any major node (concentration pressure, dynamic resources, quality ceiling, or sustained demand) re-opens a path to self-sufficiency or automation dominance.

---

# 13. Design Invariants

For each major subsystem and for the EIC as a whole. Thresholds are future tuning variables, not invented numbers.

## 13.1 Specialisation / Concentration

- **Principle:** One character cannot be simultaneously elite across the full set of high-value economic production roles.  
- **Constraint:** Must not create permanent character traps or eliminate meaningful hybrid identities.  
- **Invariant:** Under normal play, the marginal return to deepening a primary specialisation exceeds the return to adding a new elite domain beyond the budget.  
- **Failure Condition:** A significant population of characters efficiently supplies elite output across three or more previously distinct high-value domains without meaningful quality loss.  
- **Test:** Agent simulation of generalist vs specialist returns; multi-account optimisation runs.

## 13.2 Resource System

- **Principle:** Resources create discovery events, temporary scarcity, regional differentiation, and non-fungible quality.  
- **Constraint:** Must not reduce to a single “best resource” ranking that dominates all recipes.  
- **Invariant:** Exceptional resource instances create measurable, time-limited shifts in local production advantages and trade flows.  
- **Failure Condition:** Players ignore resource attributes and treat all resources of a class as interchangeable, or a single static ranking dictates all high-value production.  
- **Test:** Resource lifecycle + attribute mapping simulation; observe whether attribute diversity produces differentiated product outcomes and trade.

## 13.3 Crafting Differentiation

- **Principle:** Products of the same schematic are meaningfully differentiated by resources, skill, and experimentation.  
- **Constraint:** Differentiation must not be so opaque that spreadsheet optimisation is the only viable high-end strategy.  
- **Invariant:** Two skilled crafters using different valid resource sets and experimentation choices produce products that rational consumers treat as non-identical (performance, secondary properties, or provenance).  
- **Failure Condition:** All high-end products of a schematic converge to a single optimal configuration that every rational crafter reproduces.  
- **Test:** Controlled crafting trials with varied inputs; consumer preference / price observation in simulation.

## 13.4 Manufacturing / Automation

- **Principle:** Automation scales volume of ordinary and mid-tier goods without eliminating specialist relevance.  
- **Constraint:** No infinite self-sufficient personal-factory closed loop (MFG-004).  
- **Invariant:** Elite-tier or identity-bearing output retains a material quality or provenance advantage that factories cannot fully replicate under the rules.  
- **Failure Condition:** Optimal strategy for high-value goods becomes pure factory production with minimal specialist involvement.  
- **Test:** Compare factory-only vs specialist-involved production paths for elite goods under capital and attention constraints.

## 13.5 NPC Substitution

- **Principle:** NPCs provide a floor, never a ceiling that makes players irrelevant.  
- **Constraint:** NPC goods/services must not undercut the specialist market under normal economic conditions.  
- **Invariant:** Specialist player-produced elite goods and high-reputation services command a sustained premium over NPC alternatives.  
- **Failure Condition:** Rational players prefer NPC sources for the majority of high-value economic activity.  
- **Test:** Price and volume comparison in simulation with and without NPC price/quality floors.

## 13.6 Demand

- **Principle:** Demand is sustained by multiple independent sinks so that no single activity is load-bearing.  
- **Constraint:** Must not require forced combat as the only viable sink.  
- **Invariant:** Aggregate demand remains positive and responsive to supply shocks without continuous external intervention.  
- **Failure Condition:** Persistent over-supply with no recovery mechanism, or demand collapses when any single sink is removed.  
- **Test:** Multi-sink simulation with shock scenarios (resource boom, combat reduction, construction boom).

## 13.7 EIC as a Whole

- **Principle:** Rational players preferentially specialise and transact rather than pursue universal self-sufficiency.  
- **Constraint:** Interdependence must create opportunity, not coercive inconvenience.  
- **Invariant:** In a mixed population of generalists, specialists, small networks, and organisations, the majority of high-value transactions involve specialised players or organisations rather than pure self-supply.  
- **Failure Condition:** Dominant strategy for most players becomes self-sufficient generalism or closed organisational loops that render independent specialists optional.  
- **Test:** Full multi-agent simulation (see §14).

---

# 14. Simulation Specification

**Minimum simulation required to falsify the candidate.**

### Agent Types
1. Solo generalist  
2. Specialist gatherer / prospector  
3. Processor  
4. Specialist crafter (one primary domain)  
5. Manufacturer / industrialist  
6. Retailer  
7. Service provider (repair / medical / etc.)  
8. Combat / construction consumer  
9. Small specialist partnership (2–3 agents)  
10. Large organisation (vertically oriented)  
11. Multi-account player (N characters under one decision-maker)  
12. New entrant  
13. High-skill established specialist  
14. NPC supplier (floor only)

### Core Loops to Simulate
- Resource spawn → discovery → extraction → processing → crafting / manufacturing → distribution → consumption / decay  
- Specialisation investment and respecialisation  
- Reputation / provenance premium formation  
- Factory operation under input, maintenance, and quality rules  
- Multi-account coordination costs  

### Falsification Conditions
The design is falsified if, under reasonable parameter ranges:
- Solo generalists or multi-account generalists consistently match or exceed specialist returns at the high end.  
- Factory-only production dominates elite goods.  
- Independent specialists become economically irrelevant once organisations exist.  
- New entrants have no viable entry points.  
- Attribute optimisation collapses to a single dominant strategy with no residual differentiation.  
- Transaction volume collapses because self-supply is preferred.

### Purpose
Not numerical balance. Structural validation that interdependence is an equilibrium outcome rather than an imposed inconvenience.

---

# 15. Open Decisions for Human Ruling

These decisions belong exclusively to the project owner. This investigation does not decide them.

1. **Accept or reject Architecture C as the working EIC candidate** for the next design baseline.  
2. **Form of the specialisation / concentration budget** (hard slots, points with diminishing returns, soft concentration curves, hybrid). Even the existence of a firm budget is a human decision; the investigation only shows it is useful.  
3. **Strength of the quality / identity ceiling on manufacturing** (what factories can and cannot produce at the elite tier).  
4. **Resource lifecycle model family** (time-limited, extraction-limited, hybrid, or other) — shape only.  
5. **Experimentation model family** (points, success/fail, continuous influence, other).  
6. **Attitude toward multi-accounting** (technical limits, policy, economic friction only, or accept as normal).  
7. **Relative weight of identity/provenance premiums versus pure performance** in consumer value.  
8. **Minimum viable scope for a persistent alpha** that still tests the EIC invariants (OQ-015).  
9. **Whether any currently PROPOSED elements from older documents should be explicitly resurrected or permanently retired.**

Until these are ruled on, the candidate remains PROPOSED and non-canonical.

---

# 16. Recommended Next Experiment

**Smallest useful prototype / simulation that distinguishes the leading model (C) from its strongest competitor (A):**

Build a minimal multi-agent economic simulation containing:
- A moderate concentration budget (C) vs a hard concurrent-mastery limit (A).  
- A simplified dynamic resource system with 2–3 attributes and temporary spawns.  
- A quality ceiling on “factory” agents.  
- Specialist vs generalist vs multi-account agents.  
- A simple demand sink (decay + consumption).  

**Success metric for the experiment:**  
Under both budget regimes, measure (1) specialist premium, (2) multi-account advantage, (3) factory vs specialist market share at the high end, and (4) new-entrant viability.  

If the hybrid budget produces comparable or better interdependence metrics with higher measured player-freedom / hybrid-identity scores, C is further supported. If the hard budget is clearly superior on interdependence with acceptable freedom cost, A should be reconsidered.

This experiment can be run before any full resource-attribute design or full crafting system is locked.

---

# Document Control

**Status of this document:** Design Investigation only.  
**No LOCKED decisions are created or modified by this document.**  
**Next required human action:** Ruling on the Open Decisions in §15, particularly acceptance or rejection of Architecture C as the working candidate for the Economic Interdependence Core package.

**Lineage:** Produced under the authority and status discipline of Master GDD v1.1.1 Status Patch and the Authority & Provenance Reconciliation Matrix. External evidence cited from public sources on SWG, EVE Online, and virtual-economy research; no historical game is treated as TCIndustries authority.

---

*End of Investigation*
