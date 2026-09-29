# TCIndustries — Design Invariant Register v0.1

**Status:** PROPOSED. No authority. No invariant below is binding; each is a
candidate for human approval. Approval would record it as a design invariant
(TEST-001); it would still set no numeric value.

**Filing note (2026-09-19):** per the owner's rulings HD-RET-01, HD-BLD-01, and
HD-SCOPE-01 (`governance/TCIndustries_Human_Rulings_Register.md`, Section C),
INV-016a, INV-016b, INV-007a, and INV-015 are APPROVED as design invariants
(TEST-001 class) and annotated in place below; every other invariant remains a
candidate (INV-007b explicitly so). Approval sets no numeric value.
Per-invariant granularity per dependency note §9 — no batch approval inferred.

**Filing note (2026-09-28):** per the owner's rulings HD-ITM-01, HD-PROF-01,
HD-BLD-02, HD-PROF-02, HD-ITM-02, and HD-BLD-03
(`governance/TCIndustries_Human_Rulings_Register.md`, Section D), **INV-010**
(conditional form) and **INV-011a** are APPROVED as design invariants (TEST-001
class) and annotated in place below; **INV-BLD-004b** is APPROVED and its
register entry is added below (text authored in the HD-BLD-02 proposal);
**INV-011b** is recorded as NOT approved in its cost form under HD-PROF-02.
INV-011b, INV-011c, and INV-BLD-004a remain candidates. HD-ITM-02 and HD-BLD-03
approved no new invariant. Per-invariant granularity again applies — no batch
approval inferred, and no candidate is promoted by this note.

**Author:** Buffy — Codebuff agent (Freebuff), T-01 deliverable, 2026-09-17.
**Task authorization:** owner approval of T-01 under
`proposals/TCIndustries_Gap_Analysis_SWG_Reuse_and_Test_Program_Proposal_v0.1.md`
§9 (recorded there 2026-09-17).
**Method (per task spec):** each open GDD area is expressed in the Master GDD
§34 form — **Principle → Constraint → Invariant → Failure Condition → Test** —
traced to its GDD anchors. Zero numeric values (BAL-001). Areas that cannot
yield an invariant are documented as such with reasons.
**Controlling references:** `canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`;
`proposals/TCIndustries_Open_Questions_Register.md`;
`proposals/TCIndustries_EIC_Program_State.md`.
**Context note:** the EIC programme already maintains invariant machinery for
interdependence (Architectures A–C falsified; HD-EIC-05/07 gate). This register
covers the *other* open GDD areas and marks EIC-coupled entries explicitly —
it opens no EIC gate.

**Legend:** [EIC-COUPLED] = inseparable from the interdependence core, whose
architecture work is gated by HD-EIC-07 — these invariants are stated so that
whatever boundaries the human later supplies can be tested against them.

---

## 1. Register (by Open Question)

### OQ-001 — Specialisation budget / skill-cap model (PROF-005) [EIC-COUPLED]

- **Principle (LOCKED):** "The game should have a capability cap or specialisation
  budget that prevents one character from mastering every economically important
  role at once." (PROF-005)
- **Constraint:** any cap model must preserve experimentation, casual
  participation, hybrid roles, respecialisation, and player autonomy (PROF-005);
  must support the respecialisation principle (PROF-004); must be validated at
  the independent-participant layer, not the single-character layer (HD-EIC-01);
  must not rely on a hard factory-quality ceiling (HD-EIC-02).
- **Invariant (candidate INV-001):** *For any cap model M and any single
  character c, there exists at least one economically important role r such that
  c cannot perform r at specialist effectiveness simultaneously with their
  achieved specialisations — and the same holds for every multi-account or
  organisational actor under the model.* (This is the invariant *form*; which
  roles count as "economically important" is a design decision, not fixed here.)
- **Failure condition:** a demonstrated strategy (single character, multi-account
  network, or organisation) achieving specialist effectiveness in every
  economically important role simultaneously under M.
- **Test (shape-level, runnable now):** adversarial walk-through per candidate
  model, actor-complete per HD-EIC-01 (the method of EIC Task T-03); numeric
  testing deferred until numerics are authorized.

### OQ-002 — Resource spawn rates, durations, depletion mathematics (RES-004/006)

- **Principle (LOCKED):** resources create discovery, scarcity, regional trade,
  product differentiation, temporary opportunity, and social/economic events
  (RES-002; PIL-001: "a resource discovery should be an **event**, not merely an
  inventory update").
- **Constraint:** resource instances are individual, temporary, statistically
  distinguishable, and geographically distributed (RES-004); availability is
  finite or cycling with discovery/depletion/rediscovery generating economic
  events (RES-006); numbers are TBD — none may be invented here.
- **Invariant (candidate INV-002):** *For any lifecycle parameterization P
  (whenever numerics are authorized), the spawn system exhibits all three regimes
  — discovery, scarcity pressure, and depletion/rediscovery — as observable
  economic events (price movement, regional traffic, information trade), and no
  parameterization within the tested envelope reduces resource change to
  inventory churn invisible to the market.*
- **Failure condition:** a parameterization where spawns deplete/refresh with no
  market-visible consequence, or where all resources are simultaneously abundant
  (no scarcity regime) or permanently absent (no economy).
- **Test:** structural simulation with NON-CANONICAL parameter envelopes
  (Gap-Analysis Task T-04 specifies exactly this); verdicts in the three-form
  taxonomy (PASS / FAIL-with-telemetry / NOT-MEASURABLE).

### OQ-003 — Resource attribute sets and attribute→outcome mapping (RES-005) [EIC-COUPLED]

- **Principle (PROPOSED):** resources possess measurable attributes relevant to
  recipes and product outcomes (RES-005).
- **Constraint:** exact attribute sets, ranges, and mappings are TBD; attributes
  must be capable of supporting product differentiation (CRFT-001 principle)
  and resource-quality-sensitive crafting; toxicity-as-defined is on the
  non-canonical list and may not be assumed.
- **Invariant (candidate INV-003):** *For any attribute set A and mapping R to
  product outcomes: for every attribute in A there exists at least one recipe
  class for which that attribute changes the outcome, and for every recipe class
  there exists at least one attribute in A that changes its outcome — i.e., no
  dead attributes and no recipe class indifferent to all attributes.*
- **Failure condition:** an attribute no recipe ever consults (dead attribute),
  or a product class whose outcome is identical regardless of input attributes
  (differentiation collapse, violating CRFT-001's principle).
- **Test:** attribute-set × recipe-class coverage matrix (structural; runnable
  as a checklist once candidate sets are proposed); numeric mapping curves
  deferred to the authorized-numerics stage.

### OQ-004 — Crafting experimentation model (CRFT-002/003) [EIC-COUPLED]

- **Principle (PROPOSED):** skilled crafters can influence product attributes
  within resource and schematic constraints (CRFT-003), creating differentiation
  and crafter reputation (PIL-002 support).
- **Constraint:** schematic structure is TBD (CRFT-002); experimentation
  point-allocation / success-critical-failure is on the **non-canonical list**
  (Open Questions Register §6) and may not be assumed as the model; a master's
  advantage must not derive from a hard factory-quality ceiling (HD-EIC-02).
- **Invariant (candidate INV-004a — skill sensitivity):** *Under equivalent
  resource quality, a higher-specialised crafter can produce outputs at least
  measurably different (better or differently distributed) than a generalist —
  the TEST-001 example invariant, in its general form.*
- **Invariant (candidate INV-004b — provenance visibility):** *Any
  differentiation produced by experimentation is attributable — the product
  retains creator/skill provenance visible enough for reputation to attach
  (PIL-002; OQ-017 overlap).*
- **Invariant (candidate INV-004c — resource gating):** *Experimentation cannot
  fully substitute for resource quality: there exist input-resource profiles
  that no amount of crafter skill converts into top-band outputs.* (Guards the
  RES-002 goal and keeps extraction/commerce roles relevant.)
- **Failure condition:** INV-004a fails when skill is outcome-irrelevant;
  INV-004b fails when identical products are indistinguishable by origin;
  INV-004c fails when skill compensates for any input, collapsing resource
  economics.
- **Test:** specification-level experiments over candidate experimentation
  shapes (Gap-Analysis Task T-05 derives the full brief); the fork's
  crafted-share and provenance evidence (Design-Evidence Register E-xx entries)
  evidences that the pattern is implementable, not which shape is right.

### OQ-005 — Manufacturing facility rules & automation constraints (MFG-002/003) [EIC-COUPLED]

- **Principle (LOCKED):** automation must not eliminate the economic or social
  relevance of other players (MFG-001); facilities convert inputs to outputs
  over time under defined rules (MFG-002).
- **Constraint:** input quality and scarcity must still matter; skilled setup
  must retain value; logistics and regional factors must remain relevant; no
  closed-loop self-sufficient industrial empire for one actor (MFG-003); the
  anti-Factorio drift constraint binds (MFG-004, DERIVED).
- **Invariant (candidate INV-005a — oversight value):** *For any automation
  configuration C, there exists a production stage in C whose throughput or
  outcome quality depends on specialised player input (setup, maintenance,
  recipe skill, or resource selection), such that removing the specialist
  strictly degrades C's output.*
- **Invariant (candidate INV-005b — input sensitivity):** *Automation output
  quality/quantity remains a function of input-resource profiles and logistics
  conditions; no configuration produces constant output independent of inputs.*
- **Failure condition (the MFG-004 test):** an actor configuration (single
  player, multi-account, or organisation) achieving closed-loop production where
  every remaining dependency on other players can be internalised or eliminated
  with no specialist-efficiency loss.
- **Test:** the falsifiable adversarial test specified in Gap-Analysis Task
  T-06 (actor-complete per HD-EIC-01; runnable when a simulation host or
  implementation exists — the brief states its own precondition).

### OQ-006 — Transportation model & fast-travel constraints (WRLD-005)

- **Principle (PROPOSED):** transportation creates meaningful geography without
  excessive friction; fast travel must not eliminate the economic relevance of
  location, local markets, or transport services (WRLD-005).
- **Constraint:** spectrum required (local movement → personal transit → public
  transit → org transport → freight → constrained fast travel); exact mechanics
  TBD; regional differentiation assumed desirable (AS-003 — flag: assumption).
- **Invariant (candidate INV-006):** *For any transport model T: there exist
  route pairs (origin, destination) whose price differential or availability
  difference sustains trade that would not exist under costless instant
  transfer — i.e., geography remains economically load-bearing under T.*
- **Failure condition:** a model where every good's price converges to a single
  global price with no regional character (ECO-002 collapse), or where
  travel/transport costs are so high that regional trade never emerges at all.
- **Test:** two-sided structural check: (a) do regional price spreads persist
  under simulated arbitrage pressure; (b) does baseline trade volume remain
  non-zero under friction. Simulation-class; depends on a regional-economy
  simulation host (candidate future task).

### OQ-007 — City formation, governance, taxation/maintenance (BLD-003, WRLD-003)

- **Principle (PROPOSED):** player cities support governance, zoning, taxation
  or maintenance models, civic structures, and community identity (BLD-003);
  settlements/cities/wilderness/corridors structure the world (WRLD-003).
- **Constraint:** exact formation/governance rules TBD; player-created society
  must remain emergent, not scripted (SOC-003; VIS-005); NPC settlements serve
  baseline needs without displacing player economy (ECO-001).
- **Invariant (design invariant INV-007a — APPROVED per HD-BLD-01, 2026-09-19):** *City governance mechanisms require
  ongoing player participation to persist — no city remains fully functional
  indefinitely with zero player involvement (guards "living society" vs.
  set-and-forget automation).*
- **Invariant (candidate INV-007b):** *Cities create differentiated economic
  effects: at least one locational advantage (zoning, tax, maintenance, civic
  structure access) is decision-relevant for commerce placement, and no
  placement is strictly dominant for all purposes.*
- **Failure condition:** INV-007a fails under an abandoned-city persistence
  strategy; INV-007b fails when either all locations are equivalent (no civic
  meaning) or one placement dominates everything (degenerate monoculture).
- **Test:** policy-shape walk-through against both failure conditions;
  full testing requires city simulation (candidate future task; T-08 brief
  structures the human decision in the meantime).

### BLD-004 — Building abandonment / placement & limits (added on filing, 2026-09-28)

*Entry text authored in `proposals/TCIndustries_BLD-004_Abandonment_Placement_Proposal_Drafter_v1_2026-09-19.md` and added here at the filing act reserved by decision record HD-BLD-02 §4(2). BLD-004 is a GDD TBD, not an OQ; this subsection is filed between the OQ-007 (city formation) and OQ-008 entries because it governs the same structures.*

- **Principle (HUMAN-LOCKED, HD-BLD-02/HD-BLD-03):** a placed structure
  requires ongoing maintenance credit; at zero credit the upkeep lapses and the
  structure is removed from the world, and everything inside it is removed with
  it — no salvage, no recovery window, no restitution (HD-BLD-03: 24-hour
  zero-credit upkeep period before lapse; owner-authorized value).
- **Constraint:** the credit balance is the recourse path (visible, declining,
  top-up before zero); no post-zero recourse window exists; removal is terminal
  at structure scale. Placement, lot, and density rules remain **TBD** (BLD-004
  remainder). All credit parameters are BAL-001-gated. Settlement-institution
  dissolution (backlog 6e(b)) remains open.
- **Invariant INV-BLD-004b — APPROVED as a design invariant (TEST-001 class),
  HD-BLD-02, 2026-09-19** (no ghost equilibrium; satisfied by construction under
  the ruled model): *Under any
  parameterisation, a settlement left with zero sustenance converges to
  non-functional state (INV-007a satisfied at the settlement scale); no
  configuration of placed structures persists indefinitely at full function
  with zero ongoing player involvement.*
- **Invariant (candidate INV-BLD-004a — function-before-title graduation) —
  remains a candidate; NOT approved (HD-BLD-02, 2026-09-19):** *For any vacancy
  state, functional degradation precedes any title
  transition, and every title transition requires a recourse window in which
  the owner (or their explicitly-designated successor) can restore the asset
  to full function by performing the ruled sustenance obligations.*
  Consistent with the ruled model: the credit balance is the recourse path, so
  no post-zero window is mandated.
- **Failure condition:** ghost-town equilibrium (vacant, fully-functional
  settlements persisting — INV-007a breach at scale); or confiscation drift
  (mechanisms that make ownership meaningfully temporary — PLR-003 breach).
- **Test (shape-level):** lifecycle walk-through — from active settlement
  through each vacancy stage, verify (a) function degrades before title, (b) a
  recourse path exists at every title-touching stage, (c) the terminal state is
  world-structure-consistent. Checklist-class, runnable at design time.

### OQ-008 — Combat scope, risk model, progression integration (CMBT-002)

- **Principle (LOCKED):** combat is one lifestyle among many (CMBT-001);
  no forced combat path (CMBT-003, DERIVED); non-combat viability binds (LOOP-002).
- **Constraint:** scope (PvE/PvP/territorial) and risk models are TBD; combat
  must generate demand for equipment, medicine, consumables (CMBT-001) without
  becoming the sole progression or status path.
- **Invariant (candidate INV-008a — demand generation):** *Combat activity
  creates net-positive demand flows into at least two non-combat economies
  (equipment, medicine, food/consumables, repair, transport) — combat is
  economically generative, not self-contained.*
- **Invariant (candidate INV-008b — non-combat parity of worth):** *There
  exists no combat-only acquisition path to (a) progression advancement, or
  (b) social status, that lacks a non-combat alternative route with comparable
  ceiling.* (Direct TEST-001-shaped expression of LOOP-002/CMBT-003.)
- **Failure condition:** INV-008a fails when combat consumes only self-produced
  goods; INV-008b fails when any essential advancement or status tier is
  combat-exclusive.
- **Test:** demand-flow analysis of any proposed combat design (checklist-class,
  runnable at design time); full testing requires the combat design to exist
  (T-08 brief structures that decision).

### OQ-009 — Currency system(s) & money sinks (ECO-004)

- **Principle (PROPOSED/TBD):** currency and sinks are unresolved; multiple or
  regional media remain possible (ECO-004); player-driven economy binds (ECO-001).
- **Constraint:** single unified currency as default is on the **non-canonical
  list** (may not be assumed); NPC vendors must not dominate production (ECO-001);
  economic observability is a PROPOSED design obligation (SAFE-003).
- **Invariant INV-009a — APPROVED as a design invariant (TEST-001 class),
  HD-ECO-01, 2026-09-28**, in the **faucet-exceeds-sink mandate** reading: *For any currency
  design, total faucet flow exceeds sink flow in every bounded observation window
  by a measurable, monitorable margin that the observability system can detect
  and attribute — i.e., inflation cannot proceed silently.* (Structure, not
  numbers: no rate is proposed; the invariant is that the *imbalance is
  observable and attributable*.) **As ruled, this is the stronger mandate
  reading:** faucets are required to outrun sinks **by design**, making net money
  growth a standing economic policy rather than a monitoring requirement alone.
  Owner's stated reason, verbatim: *"Historicity. SWG always had a problem with
  not enough money sinks. Players became money sinks by hoarding money."*
  Composing ruling: HD-RET-01 (player-only vendors, NPC commerce limited to
  baseline services) makes player creation the dominant faucet base — the two
  rulings are complementary, not in tension.
- **Invariant (candidate INV-009b — currency contestability) — NOT RULED; MOOT on
  the current design:** *If multiple
  currencies/media are adopted, each retains at least one use case in which it
  is strictly preferred, or it disappears through player choice — no
  forced-acceptance zombie currency is design-mandated.* It binds only if
  multiple currencies are adopted; HD-ECO-01 ruled a **single unified currency**
  (Structure A). Remains available as a standing constraint on any future
  multi-currency decision. Neither approved nor declined.
- **Failure condition:** unattributable aggregate money growth (no faucet/sink
  accounting), or a mandated medium with no preferred use.
- **Test:** ledger-accounting audit of any implemented/simulated economy (the
  fork's faucet/sink-tagged ledger *pattern* is the reference implementation
  shape — Design-Evidence Register); Task T-07 specifies the TCIndustries
  observability requirements this test presumes.

### OQ-010 — Item durability, decay, repair (ITM-002)

- **Principle (PROPOSED/TBD):** whether/how items decay is open; any such
  system must create demand for repair services and materials without becoming
  pure friction (ITM-002).
- **Constraint:** ownership safeguards bind (SAFE-001); decay must not make
  ownership meaningless (PLR-003 principle) nor make non-combat careers secondary
  (anti-goals #8).
- **Invariant INV-010 — APPROVED as a design invariant (TEST-001 class),
  HD-ITM-01, 2026-09-19** (conditional form; owner-approved reading per decision
  record §5a): *If decay exists, it is decoupled-enough
  from activity that all four hold: (a) decay creates repair demand flowing to
  other players (services/crafters), (b) item identity/provenance survives the
  repair cycle (PIL-002), (c) ownership remains meaningful across the item's
  life (no pre-ordained total loss), and (d) decay never destroys the only
  copy of an irreplaceable provenance-bearing asset without player-meaningful
  recourse.* Under the ruled no-decay model (HD-ITM-01 Option C) the conditional
  binds any future decay introduction; clauses (c)/(d) bind any destruction/loss
  system as ownership/recourse obligations, (a) attaches to whatever generates
  repair demand, (b) binds the provenance requirements of any repair/restoration
  mechanism. **HD-ITM-02 (2026-09-20) elaborates inside this surface and is
  consistent with the conditional; no new invariant was approved by it.**
- **Failure condition:** decay that is pure attrition (no service demand),
  erases provenance, or functions as ownership confiscation.
- **Test:** lifecycle walk-through of any proposed durability design against
  (a)–(d) (checklist-class); T-08 brief structures the durability decision.

### OQ-011 — Skill acquisition model & respecialisation rules (PROG-003, PROF-004)

- **Principle (LOCKED for respecialisation):** players must be able to
  respecialise without permanent traps; respecialisation must involve meaningful
  opportunity cost but never erase identity, provenance, business ownership, or
  reputation history (PROF-004 + its PROPOSED implementation principle).
- **Constraint:** acquisition model (points/XP/training/discovery) TBD; use-based
  skill-point acquisition is on the **non-canonical list** (may not be assumed);
  any model must support specialisation budgets (OQ-001) and respecialisation.
- **Invariant INV-011a — APPROVED as a design invariant (TEST-001 class),
  HD-PROF-01, 2026-09-19:** *Under any
  respecialisation mechanism, name, appearance, item provenance, business
  ownership, organisation membership, and reputation history survive intact.*
- **Invariant (candidate INV-011b — cost is real but bounded) — NOT APPROVED in
  its cost form under HD-PROF-02 (2026-09-20); remains a candidate:** *Respecialisation
  carries a cost the player cannot trivially nullify (time, economic, or
  capability opportunity cost) while remaining achievable — respecialisation is
  neither free nor punitive.* (Boundedness is qualitative; numbers deferred.)
  Recorded as not-approved-under-the-ruled-model, not as rejected in principle:
  the owner chose a free-drop respec (points returned, no fee/time/penalty), so a
  future respec ruling could adopt a different instrument and re-approve this.
- **Invariant (candidate INV-011c — acquisition-route equivalence) — NOT RULED;
  remains a candidate** (HD-PROF-03, 2026-09-28): *Whatever
  the acquisition model, alternative specialisations remain reachable from any
  starting state — no acquisition path locks a character out of any legitimate
  career permanently.* The owner addressed the reachability concern
  substantively when ruling trainer-gated acquisition: training is always
  available — for a fee from an NPC trainer, or free from a player who holds the
  class. No explicit approve/decline was given, and per-invariant granularity no
  approval is inferred from an answer to a question. Residual dependency recorded:
  a character needs either credits or a willing player trainer; because HD-ECO-01
  rules a faucet-exceeds-sink money-supply policy, the credits branch is not
  structurally scarce.
- **Failure condition:** identity erasure under respecialisation; zero-cost
  respec ping-ponging as a dominant strategy; or an unreachable-specialisation
  dead end.
- **Test:** state-machine walk-through of candidate models (checklist-class);
  the fork's skill-drop cascade is HISTORICAL evidence that drop-based
  respecialisation is implementable — it is not a TCIndustries rule (SWG-002).

### OQ-012 — Multi-accounting & exploit implementation (SAFE-002) [EIC-COUPLED]

- **Principle (settled stance):** legitimate multi-accounting is an in-scope
  adversarial condition with no account-level restriction (HD-EIC-03); exploit
  policies need explicit design (SAFE-002).
- **Constraint:** implementation is TBD (OQ-012); the validation layer is
  independent participants including orgs (HD-EIC-01); wash-trading and
  manipulation must be detectable (SAFE-003 observability).
- **Invariant (candidate INV-012):** *The observability surface can attribute
  economically coordinated behavior across nominal identities to reviewable
  signals (ledger patterns, anomaly lists, interdependence events) — without
  any account restriction mechanism existing.* (Detectability, not prevention:
  prevention policy is the human decision.)
- **Failure condition:** a coordination pattern (circular trades, pooled
  logistics, synthetic demand) producing no reviewable signal anywhere.
- **Test:** red-team exercise on the observability design (Task T-07 output);
  the fork's B5(i) anomaly-list *pattern* is the reference evidence shape.

### OQ-013 — Organisation rights, hierarchy, property (SOC-002)

- **Principle (PROPOSED):** organisations own property, operate businesses,
  manage membership, and participate in cities (SOC-002); player organisations
  are LOCKED (SOC-001).
- **Constraint:** collective ownership must satisfy the same safeguards as
  personal ownership (SAFE-001); ownership principle binds (PLR-003).
- **Invariant (candidate INV-013):** *For every organisation-owned asset class,
  ownership is well-defined under all membership events (join, leave, removal,
  dissolution): the asset never enters an undefined or permanently orphaned
  state, and transfer outcomes are determined by design rather than by exploit
  timing.*
- **Failure condition:** any membership transition that leaves an asset
  inaccessible, duplicable, or claimable through race/exploit behavior.
- **Test:** membership-event state-machine audit per asset class (checklist-class
  over any proposed org model); T-08-adjacent design work precedes full testing.

### OQ-014 — Species / ancestry / starting location — **no invariant (documented)**

**Not invariant-shaped.** Status DEFERRED pending setting development (GDD §30;
OQ-014). Any invariant here would presume unresolved setting decisions
(aesthetic identity, cultural mechanics). The only invariant-adjacent obligation
that *does* bind: character creation "must not trap players into inferior
long-term choices" (PLR-002) — that is captured as INV-011c (acquisition-route
equivalence) and needs no setting-specific duplicate. Recorded rather than
forced.

### OQ-015 — First playable vs persistent alpha scoping

- **Principle:** scope control for production planning (OQ-015).
- **Constraint:** which systems are required is TBD; the LOCKED vision layer
  (interdependence, player economy, non-combat viability) constrains any
  minimal slice.
- **Invariant (design invariant INV-015 — APPROVED per HD-SCOPE-01, 2026-09-19):** *Any "first playable" definition must
  include enough economy for at least one interdependence relation to be real —
  i.e., the minimal slice cannot consist of systems a single player exercises
  alone.* (Derives from VIS-004 + LOOP-002; prevents a degenerate
  single-player-crafting-game slice, which anti-goals #6/#7 forbid.)
- **Failure condition:** a proposed minimal scope whose loop closes within one
  character.
- **Test:** scope-proposal review against the invariant (checklist-class);
  the actual scoping decision remains human (T-08 brief covers it).

### OQ-016 — Vendor & retail mechanics (RET-002)

- **Principle (LOCKED):** players operate shops/vendors/commercial spaces
  (RET-001); vendors associate with player/org identity (RET-002 PROPOSED).
- **Constraint:** exact mechanics TBD; commercial districts and brands should
  be supportable (RET-003); NPC substitution limits bind (ECO-001; EIC-gated).
- **Invariant (design invariant INV-016a — identity persistence — APPROVED per HD-RET-01, 2026-09-19):** *A vendor's goods
  remain attributable to their supplier/owner through the full sale chain, and
  vendor identity (who stocks it) is visible to buyers.* (PIL-002 through
  retail; OQ-017 overlap.)
- **Invariant (design invariant INV-016b — player primacy — APPROVED per HD-RET-01, 2026-09-19):** *For every good class a
  player can produce, player-made supply is distinguishable from and not
  strictly dominated by NPC supply — NPC availability must not erase the market
  for the player version.* (ECO-001 expression at retail level.)
- **Failure condition:** anonymous commodities (no brand/identity attach), or
  NPC stock making any player production role pointless.
- **Test:** retail-flow audit of any vendor design (checklist-class); the
  fork's vendor/provenance patterns evidence implementability (Design-Evidence
  Register), not TCIndustries rules.

### OQ-017 — Provenance depth & visibility (CRFT-004, ITM-001)

- **Principle (PROPOSED):** crafted items retain meaningful provenance linking
  creator, organisation, and/or resource sources (CRFT-004); items support
  identity via provenance (ITM-001). Note: the Provenance & Reputation Engine
  itself is DEFERRED (§34 phase 2) — this invariant governs *depth/visibility
  requirements*, not engine design.
- **Constraint:** provenance tracking is assumed feasible (AS-004 — flag:
  assumption; consequence if wrong = alternative approaches needed); visibility
  rules (who sees what) are TBD; privacy/fairness considerations open.
- **Invariant (candidate INV-017):** *Provenance is present at three granularities
  — creator (always, for crafted goods), organisation (when applicable), and
  resource-origin (when distinguishable) — and visibility is decision-relevant:
  each granularity, when visible, changes at least one player decision (purchase,
  price, trust), and no granularity's visibility makes another impossible.*
- **Failure condition:** provenance recorded but invisible/unused (dead data),
  or visibility rules that leak more than decision-relevant information.
- **Test:** design review against the three-granularity requirement +
  decision-relevance check on any visibility proposal (checklist-class); the
  fork's `crafted_items` provenance evidence shows creator-granularity
  implementability (Design-Evidence Register).

---

## 2. Cross-cutting observation (recorded, not resolved)

Several invariants share one structural shape: **"no actor configuration
(single/multi-account/org) achieves X"** (INV-001, 004c, 005a/b, 012). This is
the HD-EIC-01 actor-layer lesson propagating beyond EIC — it suggests the
independent-participant validation layer should become a standing TCIndustries
testing convention. Recorded as an observation for the owner; adopting it as a
convention would be a human decision.

## 3. Acceptance self-audit (T-01 acceptance criteria)

| Criterion | Result |
|---|---|
| Every OQ-001…017 addressed (invariant proposed or documented why not) | ✅ 16 invariants proposed; OQ-014 documented as not-invariant-shaped (§1, OQ-014 entry) |
| Zero numeric values | ✅ no rate, cap, curve, or threshold proposed anywhere; "measurable/bounded" language used where shape requires |
| Every entry carries GDD citations | ✅ each entry cites rule IDs; non-canonical exclusions (toxicity, point-allocation, use-based SP, single-currency default) restated where relevant |
| Status discipline intact | ✅ document is PROPOSED; every invariant marked candidate; EIC-coupled entries flagged without opening HD-EIC-07; nothing promoted |

*Verification note: acceptance is auditable by reading this document against the
GDD §31 OQ table; no external claims are made.*

---

*End of Invariant Register v0.1. Each invariant awaits individual human approval
(TEST-001); approval records it as a design invariant without numeric content.
[Filing note 2026-09-19: INV-016a, INV-016b, INV-007a, INV-015 approved — see
head note and `governance/TCIndustries_Human_Rulings_Register.md` Section C.]*
