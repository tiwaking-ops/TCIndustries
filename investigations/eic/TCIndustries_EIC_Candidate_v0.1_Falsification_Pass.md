# TCIndustries — EIC-Candidate v0.1 Falsification Pass
## Adversarial Attack on Architecture C (Non-Canonical)

**Document Type:** Falsification / Adversarial Analysis  
**Version:** 1.0  
**Date:** 2026-08-25  
**Status:** Adversarial analysis only. Creates no design authority.  
**Controlling Candidate:** `TCIndustries_EIC_Candidate_v0.1.md`  
**Instruction obeyed:** Try to destroy Architecture C. Do not improve it. Do not elaborate it. Do not add mechanisms to save it. If it fails, say it fails.

**Method:**  
Assume a rational economic optimiser whose sole objective is maximising economic output (value produced + value captured) per unit of player attention and account infrastructure. Construct the strongest versions of self-sufficiency, multi-accounting, vertical integration, factory dominance, resource monopoly, and new-player exclusion. Test every invariant for PASS / FAIL / INCONCLUSIVE. Identify which mechanisms are actually load-bearing and which appear redundant or currently unfalsifiable.

---

# 1. Attack Stance

This document does not ask whether Architecture C is elegant.  
It asks whether Architecture C survives when players optimise against it.

The brutal test remains:

> Can the system produce meaningful player interdependence without making cooperation mandatory or self-sufficiency economically dominant?

Two opposite failure modes are treated as equally fatal:

**Failure A — Self-sufficiency dominates**  
Solo generalists, multi-accounts, or closed organisations make specialists optional. Transaction volume collapses toward self-supply.

**Failure B — Interdependence becomes coercive**  
A solo player or new entrant cannot meaningfully function without finding specialists. Transaction friction becomes mandatory. Social interaction becomes a tax rather than an opportunity.

Target state (from prior human direction): multiple viable strategies with different comparative advantages. No single strategy should dominate all high-value activity.

---

# 2. Rational Economic Optimiser (Meta-Agent)

**Objective function:**  
Maximise (economic value produced + value captured − costs) / (player attention hours + account infrastructure cost).

**Behaviours it actively searches for:**  
- Best generalist configuration on one character  
- Optimal alt combinations under one decision-maker  
- Vertical integration depth that internalises specialist functions  
- Resource capture and temporary monopoly extraction  
- Factory scaling that minimises ongoing specialist attention  
- Market concentration and arbitrage  
- Self-supply thresholds (when is buying worse than making?)  
- Organisational structures that convert many characters into one economic will  
- Any residual niche that can be automated or commoditised  

This agent is assumed smarter than the designer and better-informed than the average player. It does not role-play; it optimises.

---

# 3. Strongest Attack Strategies Constructed

## 3.1 Strongest Self-Sufficient Solo Strategy

**Configuration:**  
One character. Broad competence across extraction, processing, crafting, and basic manufacturing. Uses NPC floor for gaps. Focuses on mid-tier goods where differentiation premiums are small. Avoids elite-tier competition. Minimises decay exposure. Self-repairs. Self-transports within local region.

**How it attacks the candidate:**  
If mid-tier goods constitute the majority of economic volume, and elite/specialist premiums are confined to a thin high-end slice, the solo generalist can remain economically comfortable without ever transacting with a specialist. M1 (concentration pressure) is weak against “good enough” generalism. M3 and M5 only matter if consumers actually pay for differentiation at scale.

**Current vulnerability:** HIGH.  
The candidate has no demonstrated floor on how large the “good enough” self-supply band can be before specialist demand collapses.

## 3.2 Strongest Multi-Account Strategy

**Configuration:**  
N characters under one decision-maker. Each character takes one elite domain (or one stage of the chain). Shared capital, shared logistics, shared information. Attention is sequential or lightly parallelised. Facilities and inventory are treated as poolable. Reputation is managed on the most visible characters; production characters remain low-profile if beneficial.

**How it attacks the candidate:**  
M1’s possible implementations (formal budget, soft capacity, knowledge cost, facility concentration, opportunity cost) are all weaker against multi-accounting than against single-character generalism. A formal budget per character is almost definitionally circumvented by alts. Soft capacity and attention costs scale sub-linearly with organisation and practice. Knowledge can be externalised into shared documents and tools. Facility concentration can be distributed. The multi-account player approximates a small organisation without needing other humans.

**Current vulnerability:** CRITICAL.  
Unless multi-accounting carries material, hard-to-reduce costs (account policy, capital inefficiency, logistics friction that does not fully pool, or attention costs that remain high), M1 largely fails against this strategy. The candidate currently lists multi-accounting as a circumvention but does not demonstrate that residual mechanisms still produce interdependence once alts exist.

## 3.3 Strongest Vertically Integrated Organisation

**Configuration:**  
Large player organisation. Dedicated roles for prospecting, extraction, processing, component crafting, manufacturing, logistics, retail, and services. Internal markets or directed allocation. Capital concentration. Ability to absorb temporary resource shocks by switching regions or stockpiling. Recruits or trains specialists as employees rather than buying on the open market when profitable.

**How it attacks the candidate:**  
Organisations convert external interdependence into internal coordination. M2’s temporary resource advantages can be captured and held longer by coordinated multi-region presence. M3 differentiation can be internalised by employing the specialists. M4’s residual niches (setup, responsiveness, customisation) can be staffed internally. M5 provenance can be organisational brand rather than independent specialist brand. M7 reputation accrues to the organisation.

**Current vulnerability:** HIGH.  
The candidate expects residual demand for independent specialists. It has not shown that organisations will prefer open-market specialists over internalised ones once scale is reached. Healthy market power is acceptable; permanent closure that renders independents optional is not. The boundary is currently unproven.

## 3.4 Strongest Factory-Dominated Strategy (No Hard Quality Ceiling)

**Configuration:**  
Industrialist invests heavily in manufacturing capacity. Achieves high measured quality through good inputs (M2), careful initial setup, and periodic retuning. Uses automation to produce large volumes of products that match or closely approach specialist measured performance. Treats provenance as optional or organisational. Uses market power and volume to set prices that undercut low-volume specialists. Minimises ongoing specialist attention after initial configuration.

**How it attacks the candidate:**  
This is the central test of the re-framed M4.  
If factories can eventually match raw quality, the residual value of specialists must come from setup decay, responsiveness to new resources, customisation, provenance premiums, complementary services, or information advantages.  

Each residual channel can be attacked:  
- Setup knowledge → document, template, and share; or hire once and retain.  
- Responsiveness to new resources → organisational monitoring teams or multi-account sensors.  
- Customisation → limited market size; most demand may be for standardised high-quality goods.  
- Provenance → many consumers optimise pure performance; identity premiums may be thin.  
- Complementary services → can themselves be organised or automated over time.  
- Information → public or sold, not permanently scarce.

**Current vulnerability:** CRITICAL under the no-hard-ceiling rule.  
The candidate correctly refuses a hard quality lock. It has not yet demonstrated that the residual channels survive economic optimisation at scale. This is the single most important open failure mode in v0.1.

## 3.5 Strongest Resource Monopoly Strategy

**Configuration:**  
Network of characters or organisation focuses on survey, rapid deployment, and exclusive or preferential access to exceptional resource instances. Captures high-attribute spawns early. Sells at scarcity prices or feeds internal production. Uses information asymmetry. When a spawn ends, moves to the next.

**How it attacks the candidate:**  
M2 is designed to make monopolies temporary. The attack tests whether “temporary” is short enough and whether capture costs are high enough that monopoly rents do not become the dominant high-value strategy, crowding out broader specialist participation.

**Current vulnerability:** MEDIUM.  
Lifecycle temporality is a strong structural defence if actually implemented with meaningful duration limits and multi-region competition. Without concrete lifecycle shape, the attack cannot be fully resolved. Currently INCONCLUSIVE pending OQ-002.

## 3.6 Strongest New-Player Failure Case

**Configuration:**  
New character. Limited capital. Limited knowledge. Limited reputation. Faces established specialists, organisations, and multi-accounts. Attempts to enter via gathering, basic crafting, services, or retail.

**How it attacks the candidate:**  
If entry points exist only in theory, or if the economically relevant activity is concentrated in elite tiers that require long investment and existing networks, new players experience the system as closed. This produces Failure B (coercive interdependence) or simple exclusion.

**Current vulnerability:** MEDIUM–HIGH.  
The candidate asserts multiple entry points. It has not demonstrated that those entry points remain economically meaningful once optimisers and organisations exist. If mid-tier and service markets are thin, new-player paths become decorative.

---

# 4. Mechanism-by-Mechanism Attack Results

## M1 — Capability Concentration Pressure

**Attack result:** VULNERABLE / currently under-specified.

Formal budget: easily attacked by multi-accounting.  
Soft capacity / attention: scales with player skill at multitasking and tooling; multi-accounts and organisations reduce the effective cost.  
Knowledge maintenance: externalisable into shared tools and documents.  
Facility concentration: distributable across characters/accounts.  
Pure opportunity cost: only works if the quality/throughput gap is large *and* consumers actually pay for the gap at scale.

**Verdict on necessity:**  
M1 in some form appears necessary to prevent single-character universal elite generalism.  
M1 as currently specified is **not sufficient** against multi-accounting and organisational internalisation.  
Whether a *formal* concentration budget is required is **unproven**. The candidate correctly leaves this open; the falsification pass confirms it must remain open until comparative simulation exists.

**Invariant status:** INCONCLUSIVE (depends on implementation choice and multi-account costs).

## M2 — Dynamic Attribute-Bearing Resources with Temporality

**Attack result:** STRUCTURALLY STRONG if lifecycle is real and multi-dimensional attributes resist single ranking.

Primary residual risks: information monopoly, spreadsheet collapse of attributes, and multi-account harvesting networks that capture most exceptional spawns. These are real but do not fully destroy the mechanism if spawns are sufficiently temporary and geographically dispersed.

**Verdict on necessity:** HIGH. Removing M2 collapses resource differentiation and removes a major source of temporary comparative advantage. Appears load-bearing.

**Invariant status:** PASS (conditional on concrete lifecycle and attribute design that has not yet been specified). Currently the invariant is well-formed but not yet empirically tested.

## M3 — Crafting Differentiation

**Attack result:** VULNERABLE to convergence.

If experimentation/knowledge space is fully solvable, public guides will collapse differentiation into a solved optimal configuration. If consumers primarily optimise measured performance, secondary properties and “style” command thin premiums. Automation that can reproduce the same configurations further erodes specialist uniqueness.

**Verdict on necessity:** HIGH in principle (supports CRFT-001 and identity).  
Current form is **not yet shown to survive** public optimisation and consumer pure-performance preference.

**Invariant status:** INCONCLUSIVE / at risk of FAIL if differentiation is primarily performance-based and solvable.

## M4 — Automation Relevance Limits (No Hard Quality Ceiling)

**Attack result:** CRITICAL VULNERABILITY.

This is the hardest and most important attack surface.  
Under the explicit rule that factories may match or exceed raw quality, every residual specialist channel can be attacked by scale, documentation, organisation, or consumer indifference to provenance.

The candidate’s residual list (setup decay, responsiveness, customisation, provenance, services, information) is plausible but currently unproven. No evidence yet shows that these channels remain economically significant after optimisers have had time to close them.

**Verdict on necessity:** The *problem* M4 addresses is load-bearing (MFG-001 / MFG-004).  
The *current non-hard-ceiling solution set* is **not yet demonstrated to work**. This is the single largest open risk in the candidate.

**Invariant status:** INCONCLUSIVE → leaning FAIL until residual channels are shown to survive optimisation.  
A hard quality ceiling would make the invariant easier to satisfy but is explicitly rejected as non-canonical. The candidate is therefore taking the harder, correct path — and has not yet shown it can walk it.

## M5 — Provenance and Identity Visibility

**Attack result:** WEAK unless consumers actually value it.

Provenance only matters if a meaningful fraction of high-value demand pays for it. Historical MMO evidence is mixed; many players optimise pure stats. If provenance premiums are thin, M5 becomes decorative.

**Verdict on necessity:** SUPPORTING, not load-bearing by itself. Useful amplifier of M3 and residual channel for M4. Insufficient alone.

**Invariant status:** INCONCLUSIVE (depends on actual consumer preference weight, which is an empirical / design choice question).

## M6 — Multi-Source Demand Generation

**Attack result:** NECESSARY but currently underspecified.

Without sustained demand, specialisation has nothing to sell into. Portfolio of sinks is the correct structural idea. Exact weighting is open. Risk of over-production outrunning sinks, or of players minimising exposure to decay, is real but standard and manageable.

**Verdict on necessity:** HIGH. Load-bearing for any interdependence model.

**Invariant status:** PASS (structural). Implementation details remain open and do not currently falsify the candidate.

## M7 — Reputation and Economic Feedback

**Attack result:** SUPPORTING / slow.

Reputation compounds focused specialisation over time. It is attackable by pure mechanical performance matching, alt cycling, and organisational brand substitution. Valuable but not fast enough or strong enough to carry interdependence by itself.

**Verdict on necessity:** SUPPORTING. Reinforces M1 and M3. Not independently sufficient.

**Invariant status:** INCONCLUSIVE (depends on observability and resistance to gaming).

## M8 — NPC Substitution Floor Only

**Attack result:** NECESSARY boundary condition.

If the NPC floor is too high, specialists are undercut. If too low or absent, new-player and crisis paths fail. The mechanism itself is sound; the positioning is a tuning problem.

**Verdict on necessity:** HIGH as a boundary. Not a primary interdependence engine.

**Invariant status:** PASS (structural), pending price/quality positioning.

---

# 5. Formal Concentration Budget vs Non-Budget Alternatives

**Comparative question (as directed):**  
Does a formal concentration budget materially outperform a pure soft-pressure set (capacity + knowledge + facilities + opportunity cost + resources + geography + reputation + production constraints) on the interdependence invariants?

**Current answer from adversarial analysis:**

- Against **single-character** universal elite generalism: a formal budget is the cleanest and most reliable instrument. Soft pressures can work but require large, stable quality/throughput gaps and consumer willingness to pay for them.
- Against **multi-accounting**: a per-character formal budget is largely neutralised. Soft pressures are also weakened but may retain more residual force if attention, capital inefficiency, and logistics do not fully pool.
- Against **organisations**: both formal budgets and soft pressures are partially internalised. Residual specialist demand must come from elsewhere (M2 temporality, M4 residual niches, M5/M7 identity).

**Conclusion:**  
A formal concentration budget has **not earned its existence** yet.  
It may still be justified, but only if comparative simulation shows material improvement on the invariants that soft pressures cannot match, *especially* under multi-account and organisational attack.  

Until that comparison exists, HD-1 remains correctly open. Choosing a budget now would be design preference, not evidence.

---

# 6. Invariant Scorecard

| Mechanism | Invariant | Status | Notes |
|---|---|---|---|
| M1 | Marginal return to deepening primary > adding another elite domain | INCONCLUSIVE | Depends on implementation; weak against multi-account |
| M2 | Exceptional resources create time-limited shifts; attributes do not collapse to one ranking | PASS (conditional) | Strong structural claim; needs concrete lifecycle |
| M3 | Differentiated products treated as non-identical by consumers | INCONCLUSIVE / at risk | Vulnerable to solved optimisation and pure-performance preference |
| M4 | Residual specialist demand remains even when automation matches raw quality | INCONCLUSIVE → leaning FAIL | Central open risk; residual channels unproven |
| M5 | Provenance/identity commands measurable premium | INCONCLUSIVE | Empirical / preference dependent |
| M6 | Aggregate demand remains positive; no single sink is load-bearing | PASS (structural) | Implementation open but direction sound |
| M7 | Established specialists hold durable advantage over pure mechanical newcomers | INCONCLUSIVE | Slow and gameable |
| M8 | Specialist high-value goods command premium over NPC floor | PASS (structural) | Positioning is tuning |
| **Whole EIC** | High-value activity preferentially involves specialised players/orgs rather than pure self-supply; cooperation is not mandatory | **INCONCLUSIVE → at risk of FAIL under optimiser attack** | Self-sufficiency and factory-dominance strategies are not yet defeated |

---

# 7. Which Mechanisms Appear Necessary vs Redundant

**Apparently load-bearing (remove → interdependence collapses):**  
- M2 (dynamic resources)  
- M3 (differentiation) — in principle  
- M4 (automation relevance limits) — the *problem* is load-bearing; the *solution* is unproven  
- M6 (multi-source demand)  

**Supporting / amplifying:**  
- M1 (concentration pressure) — useful, form open, insufficient alone  
- M5 (provenance)  
- M7 (reputation)  
- M8 (NPC floor)  

**Currently redundant or non-discriminating:**  
None are clearly redundant. Several are under-powered or unproven. No mechanism can yet be removed without increasing risk.

**Currently unfalsifiable (lack concrete form):**  
- Exact shape of M1  
- Exact residual channels under M4 that survive optimisation  
- Concrete resource lifecycle (M2)  
- Concrete experimentation/differentiation model (M3)  
- Consumer weight on provenance vs performance (M5)

---

# 8. Direct Answers to the Required Questions

1. **10-agent matrix:** Run conceptually. Solo generalist and multi-account strategies are the most dangerous to interdependence. Large organisation and factory-dominated strategies are the next most dangerous. New-entrant path is at risk of becoming thin.

2. **Rational economic optimiser:** Explicitly introduced and used as the attack stance.

3. **Strongest self-sufficient strategy:** Constructed (§3.1). Currently not defeated.

4. **Strongest multi-account strategy:** Constructed (§3.2). Currently the most effective attack on M1.

5. **Strongest vertically integrated organisation:** Constructed (§3.3). Residual independent specialist demand is asserted but not demonstrated.

6. **Strongest factory-dominated strategy:** Constructed (§3.4). Central test of no-hard-ceiling M4. Currently the candidate’s largest vulnerability.

7. **Strongest resource-monopoly strategy:** Constructed (§3.5). Temporality is the main defence; currently INCONCLUSIVE pending lifecycle design.

8. **Strongest new-player failure case:** Constructed (§3.6). Entry points exist in theory; economic meaningfulness under optimisation is unproven.

9. **No-hard-quality-ceiling manufacturing model:** Explicitly stress-tested. Residual specialist value channels are plausible but unproven. Leaning toward failure until shown otherwise.

10. **Formal concentration budget vs non-budget alternatives:** Compared (§5). Budget has not earned its existence. Soft pressures may suffice against single characters but are also weak against multi-accounts. Comparative simulation required.

11. **Which M1–M8 are actually necessary:** M2, M3 (principle), M4 (problem), M6 appear load-bearing. Others support or amplify.

12. **Redundant mechanisms:** None clearly redundant. Several are currently too weak or unproven.

13. **Invariants that cannot currently be falsified:** Those depending on unspecified forms (exact M1, exact M4 residuals, concrete lifecycle, concrete experimentation model, real consumer provenance weight).

14. **PASS / FAIL / INCONCLUSIVE for every invariant:** See scorecard in §6. Whole-candidate interdependence invariant is **INCONCLUSIVE → at risk of FAIL**.

15. **No redesign to save the candidate:** Obeyed. No new mechanisms added. No rescue patches proposed.

---

# 9. Overall Verdict on Architecture C / EIC-Candidate v0.1

**Architecture C is not destroyed.**  
It is also **not validated**.

**Status after adversarial pass:**

- The coupled-system insight remains sound.  
- Several mechanisms (especially M2 and M6) are structurally strong.  
- The refusal of a hard factory quality ceiling is correct on governance grounds and forces the harder, better question.  
- The openness of the concentration budget is correct; it has not earned implementation.  
- The candidate currently **fails to demonstrate** that residual specialist value survives factory quality parity and multi-account optimisation.  
- The candidate currently **fails to demonstrate** that self-sufficiency is not the dominant comfortable strategy for a large fraction of economic activity.  
- The candidate currently **fails to demonstrate** that new-entrant paths remain economically meaningful under optimisation.

**Therefore:**

EIC-Candidate v0.1 survives as a **leading structural hypothesis**.  
It does **not** survive as a design that has been shown to produce the required interdependence.

The correct scientific status is:

> **NOT YET FALSIFIED, BUT UNDER SERIOUS UNRESOLVED ATTACK.**  
> Primary attack surfaces: multi-accounting vs M1, and factory quality parity vs M4 residual channels.

---

# 10. What This Pass Does *Not* Authorise

- It does not promote any mechanism to LOCKED or baseline.  
- It does not resolve HD-1 through HD-9.  
- It does not add new mechanics.  
- It does not redesign Architecture C.  
- It does not claim simulation results; all analysis is structural and adversarial reasoning only.

---

# 11. Required Next Step (No Expansion of Design)

The next required activity is **not** another design document.

It is a **minimum comparative simulation** that can produce evidence on the two critical unresolved attacks:

1. **Factory quality parity test**  
   Can residual specialist channels (setup, responsiveness, customisation, provenance, services) retain economically significant demand when automation matches measured quality? Measure specialist premium and market share under optimiser pressure.

2. **Concentration pressure comparison**  
   Formal budget vs pure soft-pressure set, under single-character, multi-account, and small-organisation agents. Measure effect on specialist premium, self-supply rate, and multi-account advantage.

Only after these two tests produce evidence should any human ruling on HD-1, HD-3, or HD-4 be considered.

If the simulation shows that residual channels collapse and soft pressures are insufficient, Architecture C (as currently formulated) fails and a different structural approach is required.  
If the simulation shows residual channels and/or a formal budget produce material interdependence gains, the candidate advances with evidence.

Until then, the honest status remains:

**EIC-Candidate v0.1 — Leading hypothesis under unresolved adversarial pressure. Not validated. Not rejected. Not canonical.**

---

**Document Control**

**Status:** Adversarial falsification analysis only.  
**No design authority created.**  
**No mechanisms added or removed.**  
**No candidate redesign performed.**  

**Lineage:** Produced in direct response to the instruction to try to destroy Architecture C rather than improve it. Respects all status and authority rules of Master GDD v1.1.1 and the Authority & Provenance Reconciliation Matrix.

---

*End of Falsification Pass*
