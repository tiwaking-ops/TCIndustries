# SWG Historical Accuracy — Human Recollection Questionnaire v1

**Status:** PROPOSED candidate material. HUMAN-APPROVED for authoring (owner decision 2026-10-18,
Item 8 of `proposals/SWG_Post_Phase10_Program_Proposal_v0.1.md`). No design authority: answers
recorded here are EVIDENCE/HISTORICAL input only; nothing here promotes any value to TCIndustries
canon or to the testbed's locked/provisional constants. `TBD` / `unknown` / `I don't remember`
are valid and preferred over invention.

**Author (llm):** Beam Free (`opencode/beam-free`) — via OpenCode
**Version:** v1.0
**Date authored:** 2026-10-18

**Assessor (llm):** HY4 Preview (`orcarouter/tencent/hy4-preview-free`) — via OpenCode, 2026-10-18.
Independent assessment per explicit owner authorization of 2026-10-18. Assessment confers no
authority and changes no status.

**Assessment outcome: SOUND, with one advisory finding.**

Verified by independent inspection:

1. **Author attribution — CORRECT.** Author field now names the true runtime identity. This
   corrects a mis-attribution present at first filing (see
   `investigations/Provenance_Incident_2026-10-18_Beam_Free.md`). No residual defect found.
2. **Authority basis — VERIFIED.** Item 8 and the 2026-10-18 owner split decision were confirmed
   present in the working tree of `proposals/SWG_Post_Phase10_Program_Proposal_v0.1.md` (§6, §9.1,
   and the decision table). The citation is genuine, not inferred.
3. **Status discipline — CORRECT.** PROPOSED; no claim to canon; answers explicitly designated
   EVIDENCE/HISTORICAL input only; `TBD` / `unknown` / `I don't remember` explicitly preferred over
   invention; no value is promoted to TCIndustries canon or to testbed constants. Compliant with
   project `AGENTS.md` and with the HISTORICAL status rules.
4. **Governance boundaries — RESPECTED.** No numeric tuning of economic or progression systems; no
   new residual channels or architectures invented; no EIC work drawn into scope; no automatic
   successor architecture proposed. Proxy 4 (Sandbox Validation) is honestly left DEFERRED with a
   stated reason rather than silently resolved.
5. **Scope — CORRECT.** Answers are framed as human recollection, to be appended rather than
   overwritten, with confidence recorded. Appropriate instrument for a human-pass-only proxy.

**Advisory finding (not a violation, no authority implication):** a number of question stems embed
specific figures or design claims drawn from the predecessor GDD — e.g. the 250 skill-point cap
(Q38), 30–60s PvP TTK (Q119), 150–250h elite grind (Q200), 4–6% listing fees (Q191), the 10×CLMax
drop cap (Q77), and "novice crafter ≈ novice combatant income" (Q49). These are leading questions:
they may anchor a responder's recollection toward the figure supplied rather than elicit it
independently. Because answers are recorded as EVIDENCE rather than targets, and because the
figure is visible in the stem, this does not contaminate the record — but a responder should be
aware the prompts are not neutral. Recommended (owner discretion): either retain as-is with the
prompts understood to be anchoring, or add a per-category note that figures in stems are
predecessor-GDD claims to be confirmed or contradicted, not assumed.

*Content is sound. No correction required. This assessment does not promote this document or any
answer recorded in it.*
**Responds to:** Item 8 "Historical Accuracy proxy" — the 5th of the five §1.3 success-criteria
tests (SWG Pre-CU GDD §34, HISTORICAL), which is human-pass-only and out of testbed scope
(proposal v0.2 §6.4). Proxy 4 (Sandbox Validation) is handled separately and is **DEFERRED**
(see §0).

**Authority basis:** `AGENTS.md`; `governance/TCIndustries_Testbed_Reuse_Human_Ruling_HD-TST-01_2026-09-14.md`
(patterns-only fork; fork behaviour is PROTOTYPE/EVIDENCE, never design authority);
`proposals/SWG_Phase10_Balance_Telemetry_Polish_Proposal_v0.2.md` §§2.2, 5.3, 6.4;
`sources/swg-pre-cu/SWG_PreCU_GDD.md` (HISTORICAL reference text only).
**Method:** SWG (Star Wars Galaxies, Pre-CU) is no longer playable; the only remaining record is
human recollected experience. This document is a structured interview instrument. Each question is
to be answered from recollection. Answers are recorded verbatim with the responder's handle and date,
then filed as EVIDENCE; they are never treated as measured data or as tuning targets.

---

## 0. Decision context (Item 8)

- **Proxy 4 — Sandbox Validation:** **DEFERRED.** Reason: emergent player-organized play (events,
  guilds, schemes not authored by design) requires actual players; the testbed has no live player
  cohort, and fabricating one would violate the "no simulated-player fabrication" rule (proposal
  v0.2 §5.2). The staged telemetry artifacts (anomaly lists, mission outcomes, wealth/Gini,
  volatility, profession distribution) remain on disk as EVIDENCE but receive no human annotation
  pass at this time. Re-open when a live cohort exists.
- **Historical Accuracy:** This document. The human has approved authoring it; the questions below
  are to be answered by the human from recollection. Each answer becomes a recorded EVIDENCE entry.

**Answering rules for the human:**
1. Answer from recollection only. If you are unsure, write `TBD` / `I don't remember` / `unknown` —
   never guess a number.
2. Distinguish **"I remember it worked like X"** (recollection) from **"I think it was probably X"**
   (inference). Mark inference explicitly.
3. Numbers are welcome as *recollected magnitudes* but are never promoted to targets or constants;
   they are evidence of memory, not measurement.
4. Where a system was never used by the responder, mark `N/A — never used`.
5. One answer per question is fine; multiple responders may append dated entries (do not overwrite
   prior answers — append).

---

## 1. Scope and coverage map

Categories covered (each a section below). A category marked **(fork-relevant)** is something the
testbed fork actually implements and can be compared against; others are context-only.

| # | Category | Fork-relevant? |
|---|---|---|
| A | Combat | Yes |
| B | Progression & skill trees | Yes |
| C | Economy & currency | Yes |
| D | Crafting | Yes |
| E | Resources & harvesting | Yes |
| F | Professions & service providers | Yes |
| G | Medical, wounds & incapacitation | Yes |
| H | Entertainment & buffs | Yes |
| I | Missions & terminals | Yes |
| J | Faction & PvP (GCW) | Yes |
| K | Creatures, spawns & ecology | Yes |
| L | Social, groups, guilds | Partial |
| M | Player cities & civic | Partial |
| N | Housing & structures | Partial |
| O | Death, cloning & respawn | Yes |
| P | Items, gear & equipment | Yes |
| Q | Character creation & appearance | Partial |
| R | Vendors, bazaars & trade | Partial |
| S | Travel & world | Partial |
| T | Elite professions & advanced systems | Partial |
| U | AI / NPC behaviour | Partial |
| V | Time, environment & world simulation | Partial |
| W | UI / client experience | No (context) |
| X | Cross-cutting "feel" | No (context) |

---

## A. Combat

### A.1 Melee / unarmed
1. How did melee combat *feel* in terms of speed — fast and twitchy, or slow and deliberate? Did it feel like a "dance" or a slugfest?
2. Was unarmed (brawler) viable as a primary combat path, or always a stepping stone? How did it compare to armed melee at the same skill tier?
3. Did melee weapons have meaningful reach/range differences that changed positioning, or did everyone fight at the same distance?
4. How important was facing/direction — could you "backstab" or did flanking matter?
5. Did melee combat require active management (stance, combos) or was it largely auto-attack with occasional specials?
6. What was the typical time-to-kill for a same-tier 1v1 melee duel? (recollected seconds)
7. Did damage feel "spiky" (big hits) or steady? Were there visible hit reactions?

### A.2 Ranged
8. How did ranged combat handle movement — could you shoot while moving, or did you have to stop?
9. Did range bands (point-blank / optimal / max) matter in practice, or did everyone fight at one distance?
10. How did cover/line-of-sight work — was it meaningful or ignored?
11. Did ranged feel dominant over melee, balanced, or weaker? Did the meta shift over time?
12. What was typical TTK for ranged 1v1 at equivalent gear?
13. Did ammunition matter (consumable, weight, crafting) or was it effectively free?

### A.3 Specials / abilities
14. How were combat specials triggered — hotbar, radial menu, macros? How many were practical to use in a fight?
15. Did specials have cooldowns, resource costs (HAM), or both? Which gated more?
16. Were there "combos" or chains that required sequencing, or were specials independent?
17. Did specials feel impactful (visible, decisive) or like minor modifiers?
18. Were there defensive specials (dodge, block, parry) that players actually used?
19. How did crowd control (stun, blind, dizzy, knockdown) feel — frequent and decisive, or rare?

### A.4 HAM (Health/Action/Mind)
20. Did the three pools feel distinct and meaningful, or did everyone just watch health?
21. Which pool was most commonly targeted by attacks? Did "mind" attacks (intimidate, etc.) ever matter?
22. Did HAM damage feel like a resource to manage during a fight, or just a death timer?
23. How did buffs interact with HAM — did buffs feel like a meaningful combat lever?
24. Did "wounds" (post-combat HAM loss) feel like a real penalty or a minor inconvenience?

### A.5 Incapacitation & death in combat
25. What happened when you were incapacitated — how long, what could you do, could allies help?
26. Did incapacitation feel fair or frustrating? Was there a "bleed out" timer?
27. Could you be finished off while incapacitated, and did that feel common?
28. How did the incap→revive→death loop feel in group combat vs solo?

---

## B. Progression & skill trees

### B.1 XP and skill acquisition
29. How did XP gain feel — fast and rewarding early, slow later? Did it feel grindy?
30. Were there "XP debt" or "XP loss on death" mechanics, and how did they feel?
31. Did training a skill feel like a meaningful choice or a formality once XP was earned?
32. Were trainers located sensibly (in cities) and did finding them feel like part of the game?
33. Could you "unlearn" or respec skills, and did players want to?
34. How long did basic mastery feel (recollected hours)? Elite mastery? Full cap?

### B.2 Skill trees / boxes
35. Did the skill tree feel deep and branching, or linear?
36. Were there "must-have" boxes everyone took, or genuine build diversity?
37. Did hybrid builds feel viable, or did specialization dominate?
38. How did the 250 skill-point cap feel — restrictive, generous, or just right?
39. Were there "dead" skills nobody used? Which ones?

### B.3 Certification / equipment gating
40. Did certification gating (weapon/armor use requires skill) feel like a meaningful progression gate or an annoyance?
41. Did you ever feel "locked out" of content by certification requirements?

---

## C. Economy & currency

### C.1 Credits & money flow
42. Did credits feel scarce or abundant at novice level? At master level?
43. Did the economy feel inflationary over time? Did prices rise?
44. Were there meaningful "money sinks" (maintenance, travel, training) that shaped play?
45. Did starting credits feel sufficient to begin?
46. How did tithing/taxes (city, guild) feel?

### C.2 Income sources
47. What were the main income sources for a novice? For a master?
48. Did combat pay (loot, mission rewards) feel comparable to crafting income?
49. Did the "novice crafter ≈ novice combatant income" balance feel true in practice?
50. Were there "get rich quick" paths that distorted the economy?

### C.3 Market dynamics
51. Did prices feel stable or volatile week to week?
52. Did bazaar/auction feel like a real market or a dumping ground?
53. Did " undercutting" and price wars feel common? How did sellers feel about it?
54. Were there cartels/monopolies, and how did the community view them?

---

## D. Crafting

### D.1 Crafting loop
55. Did crafting feel engaging (mini-game, choices) or tedious (click-through)?
56. Did experimentation feel meaningful — did tweaking properties produce visibly different items?
57. Did assembly feel like a real process or a formality?
58. How long did a typical craft session feel?
59. Did failed crafts feel punishing or acceptable?

### D.2 Schematics & progression
60. Did learning schematics feel like progression?
61. Were schematics gated by skill level sensibly?
62. Did the "generic taxonomy" (frame/grip/etc.) make sense to players?

### D.3 Quality & item variance
63. Did item quality vary enough that "a good roll" mattered?
64. Did players chase perfect rolls, and how did that feel (fun grind vs toxic)?
65. Did crafted items feel meaningfully better than looted/NPC items?

### D.4 Factories & mass production
66. Did factories feel like a meaningful endgame for crafters?
67. Did factory runs feel satisfying or like AFK time?
68. Were there "factory spam" problems (too many identical items)?

---

## E. Resources & harvesting

### E.1 Surveying
69. Did surveying feel like exploration or a chore?
70. Did resource spawns feel dynamic and interesting, or predictable?
71. Did "shift" changes (resource stats changing) create meaningful gameplay?

### E.2 Sampling
72. Did sampling feel engaging or repetitive?
73. Did resource quality (stats) matter enough to seek specific spawns?
74. How did resource depletion feel — fair or frustrating?

### E.3 Resource economy
75. Did "good" resources feel scarce and valuable?
76. Did players hoard resources, and how did that affect markets?
77. Did the 10×CLMax drop cap (creature loot) feel right?

---

## F. Professions & service providers

### F.1 Medic
78. Did healing feel responsive and meaningful in combat?
79. Did wound treatment feel like a real service or a formality?
80. Did medics feel valued and needed, or underused?
81. Did stimpacks feel like a crutch that reduced medic value?

### F.2 Entertainer / buffs
82. Did buffs feel impactful (noticeable stat changes)?
83. Did watching entertainers feel like a social activity or a tax?
84. Did the "buff cycle" (travel to cantina, wait, get buffed) feel good or tedious?
85. Did Mind-buffing / action-buffing feel as important as health?

### F.3 Other service professions
86. Did image designers feel used? Did appearance changes matter to players?
87. Did smugglers feel like a real profession with a role?
88. Did scouts/rangers feel useful (camping, tracking)?
89. Did merchants feel distinct from regular traders?

### F.4 Service XP
90. Did service XP gain feel fair for the time invested?
91. Did flat XP awards (per heal/buff) feel right, or should they scale with amount healed/buffed?
92. Did "buff begging" (players asking for free buffs) feel like a social problem?

---

## G. Medical, wounds & incapacitation

### G.1 Wounds
93. Did wounds feel like a meaningful post-combat consequence?
94. Did wound healing feel accessible (enough medics) or scarce?
95. Did "wound rot" or neglect feel like a real risk?

### G.2 States (stun, blind, dizzy, knockdown, etc.)
96. Which states felt most impactful? Which were ignored?
97. Did state-stacking feel fair or frustrating?
98. Did states create memorable moments or just annoyance?

---

## H. Entertainment & buffs (extended)

99. Did the entertainer profession feel like a "real" profession or a buff vending machine?
100. Did music/dance feel like gameplay or decoration?
101. Did group buffs (camp buffs) feel meaningful?
102. Did "buff duration" feel right — too short, too long?

---

## I. Missions & terminals

### I.1 Mission terminals
103. Did mission terminals feel like a good "go do something" entry point?
104. Did mission variety feel sufficient or repetitive?
105. Did mission difficulty scaling feel right?
106. Did mission rewards feel worth the time?

### I.2 Mission types
107. Which mission types were most/least popular? Why?
108. Did delivery/escort missions feel fun or tedious?
109. Did bounty missions feel fair to targets?
110. Did crafting missions feel meaningful?

### I.3 Mission economy
111. Did missions dominate the economy (faucet), and did that feel inflationary?
112. Did mission-running feel like "work" or "play"?

---

## J. Faction & PvP (GCW)

### J.1 Faction choice
113. Did choosing Imperial/Rebel feel meaningful and permanent?
114. Did faction identity matter in daily play, or only in PvP zones?

### J.2 Overt/covert flagging
115. Did the flagging system (overt/covert) feel intuitive?
116. Did "going overt" feel like a real commitment with real risk?
117. Did covert players feel like they had a distinct role?

### J.3 PvP combat
118. Did PvP feel balanced between professions, or were there dominant builds?
119. Did TTK in PvP feel right (30–60s)?
120. Did "zerg" (group) PvP dominate over 1v1?
121. Did PvP feel fun or frustrating overall?

### J.4 Death penalties in PvP
122. Did losing to another player feel fair? What was lost?
123. Did "deathblows" (finishing incapacitated players) feel like a problem?
124. Did insurance/cloning mitigate loss enough?

### J.5 Bases & GCW
125. Did faction bases feel like meaningful objectives?
126. Did base assaults feel epic or tedious?
127. Did GCW feel like a living war or a background system?

---

## K. Creatures, spawns & ecology

### K.1 Creatures
128. Did creatures feel dangerous enough at appropriate levels?
129. Did creature variety feel sufficient?
130. Did "con" (difficulty) colors feel accurate?
131. Did named/boss creatures feel special?

### K.2 Lairs & spawns
132. Did lairs feel like a good PvE activity?
133. Did spawn density feel right — too many, too few?
134. Did creatures feel like they "lived" in the world (ecology) or just popped in?

### K.3 Tamable pets (creature handler)
135. Did taming feel engaging?
136. Did pets feel useful in combat?
137. Did the pet ownership/cap feel right?

---

## L. Social, groups, guilds

### L.1 Grouping
138. Did grouping feel necessary or optional?
139. Did group XP sharing feel fair?
140. Did "group roles" (tank/healer/dps) exist, or was everyone self-sufficient?
141. Did groups feel cohesive or just a party of solo players?

### L.2 Guilds
142. Did guilds feel meaningful beyond a chat channel?
143. Did guild size limits feel right?
144. Did guild politics feel like emergent gameplay?

### L.3 Chat & communication
145. Did chat channels (spatial, guild, group, OOC) feel sufficient?
146. Did the community feel healthy/toxic over time?
147. Did spatial chat (range-limited) create immersion?

---

## M. Player cities & civic

### M.1 City founding
148. Did founding a city feel achievable and meaningful?
149. Did city placement feel strategic?
150. Did the "mayor" role feel like real governance?

### M.2 City management
151. Did maintenance costs feel sustainable?
152. Did city structures (walls, defenses) matter?
153. Did cities feel like living communities or ghost towns?
154. Did city specialization (crafting city, combat city) emerge?

### M.3 Civic structures
155. Did civic structures (cloning, shuttles, vendors) feel like meaningful infrastructure?

---

## N. Housing & structures

### N.1 Placement & decoration
156. Did housing feel like a meaningful personal space?
157. Did decoration feel deep or limited?
158. Did structure placement feel strategic (buffs, maintenance)?

### N.2 Maintenance & upkeep
159. Did maintenance feel like a fair ongoing cost or a punishment?
160. Did structures decay/destroy feel fair?

---

## O. Death, cloning & respawn

### O.1 Cloning
161. Did cloning feel like a meaningful preparation (choosing a clone point)?
162. Did "no clone nearby" feel like a real risk?
163. Did cloning costs feel fair?

### O.2 Death penalties
164. Did death feel punishing enough to matter but not rage-inducing?
165. Did "deathblow" / item loss feel fair?
166. Did insurance (item insurance) feel worth it?
167. Did the "corpse run" feel like an adventure or a chore?

### O.3 Respawn experience
168. Did respawning feel disorienting or smooth?
169. Did you lose "progress" (buffs, states) on death in a way that felt right?

---

## P. Items, gear & equipment

### P.1 Weapons
170. Did weapon variety feel sufficient? Were there "best" weapons everyone used?
171. Did weapon stats (speed, damage, range) create meaningful trade-offs?
172. Did looted weapons ever compete with crafted?

### P.2 Armor
173. Did armor feel like a real protection system?
174. Did armor sets/bonuses exist and matter?
175. Did "armor decay" feel fair?

### P.3 Consumables
176. Did food/drink buffs feel meaningful?
177. Did stims feel like a combat necessity or an optional boost?
178. Did consumable crafting feel integrated with the economy?

### P.4 Itemization overall
179. Did the "no levels, gear-driven" philosophy feel right compared to level-based MMOs?
180. Did item binding/trading feel right (tradeable vs no-trade)?

---

## Q. Character creation & appearance

181. Did species choice feel meaningful (stats, abilities) or cosmetic?
182. Did appearance customization feel deep enough?
183. Did "image designer" respecs feel like a real service?
184. Did starting gear feel appropriate?

---

## R. Vendors, bazaars & trade

### R.1 Vendors
185. Did placing a vendor feel like a real business?
186. Did vendor maintenance (funding) feel fair?
187. Did "vendor search" (bazaar) feel adequate?

### R.2 Direct trade
188. Did direct player trading feel safe and convenient?
189. Were there scam risks, and how did the community handle them?

### R.3 Marketplaces
190. Did the bazaar/auction feel like the "real" market?
191. Did listing fees/commissions feel right (4–6%)?

---

## S. Travel & world

### S.1 Planets & zones
192. Did the galaxy feel vast and explorable?
193. Did travel time feel like a feature (immersion) or a tax?
194. Did different planets feel distinct?

### S.2 Shuttles & ports
195. Did shuttle schedules feel immersive or annoying?
196. Did shuttle costs feel fair?

### S.3 Vehicles & mounts
197. Did vehicles feel necessary and fun?
198. Did mount speed feel right?

---

## T. Elite professions & advanced systems

199. Did elite professions (Commando, Bounty Hunter, etc.) feel worth the grind?
200. Did "multi-prereq elite" feel like a real achievement (150–250h)?
201. Did elite abilities feel powerful and special?
202. Did the "no Jedi at launch" (then Jedi unlock) situation feel right?
203. Did Bio-engineering / DNA / pets feel like a deep system?

---

## U. AI / NPC behaviour

204. Did NPCs (vendors, trainers, quest givers) feel alive or static?
205. Did creature AI feel believable (flee, aggro, pack behavior)?
206. Did "con" (difficulty) and aggro ranges feel right?

---

## V. Time, environment & world simulation

207. Did day/night cycle matter for gameplay?
208. Did weather matter?
209. Did the world feel "alive" (dynamic spawns, events)?
210. Did "fast travel" options undermine exploration?

---

## W. UI / client experience (context only)

211. Did the radial menu feel intuitive or clunky?
212. Did chat/UI feel adequate for an MMO?
213. Did the map/navigation feel helpful?
214. Did performance (client) feel acceptable?
215. Did the UI get in the way of immersion?

---

## X. Cross-cutting "feel"

216. Overall, did SWG Pre-CU feel like a "world" or a "game"?
217. What was the single best-feeling system, and why?
218. What was the single worst-feeling system, and why?
219. Did the game reward "mastery" (skill) or "time" (grind) more?
220. Did the community's emergent play (player cities, events, markets) feel like the real game?
221. Did the "sandbox" promise feel delivered?
222. What do you miss most that no other MMO replicated?
223. What would you *not* want to replicate?

---

## Answer record template

For each answered question, append an entry in this form (do not overwrite prior entries):

```
### Q<number> — <short question restatement>
- Responder: <handle>
- Date: YYYY-MM-DD
- Answer: <text, or TBD / unknown / N/A>
- Confidence: high | medium | low | inference
- Notes: <optional>
```

## Status of this document

- Content status: PROPOSED candidate material. Answers are EVIDENCE (human recollection), not design
  authority and not measured data.
- This document does not change any status of any other document and does not authorize any code,
  tuning, or design decision. It is an interview instrument + answer record.
- Proxy 4 (Sandbox Validation) remains **DEFERRED** (see §0) and is not covered here.

*End of SWG_Historical_Accuracy-v1.md (v1.0).*
