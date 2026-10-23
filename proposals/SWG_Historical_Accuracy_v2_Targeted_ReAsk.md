# SWG Historical Accuracy v2 — Targeted Re-Ask (Open-Question Instrument)

**Status:** PROPOSED candidate material. HUMAN-AUTHORISED for authoring (owner request, 2026-10-21,
following the Addendum A / Addendum B clarification round on
`investigations/TCIndustries_SWG_Historical_Recollection_Report_Tiwa_Pene_2026-10-21.md`).
**No design authority.** Answers recorded here are EVIDENCE / HISTORICAL input only. Nothing here promotes
any value to TCIndustries canon, closes any TBD, settles any open question, or authorises any tuning.
`TBD` / `unknown` / `I don't remember` / `never used it` are valid and **preferred over invention**.

**Instrument author (llm):** MiniMax M3.1 Flash Free (`opencode/minimax-m3.1-flash-free`) — via OpenCode, 2026-10-21
**Subject-matter source (human):** Tiwa Pene — Human Operator — the respondent whose recollection this
instrument interrogates, and the requester of the re-ask. Their own answers are the substance; the
questions are the assessor's.
**Assessor (llm):** — not yet assessed. Independent assessment is available on request, as with
`proposals/SWG_Historical_Accuracy-v1.md`.
**Version:** v1.0
**Date authored:** 2026-10-21

**Responds to:** the open-question surfaces named below — OQ-009 (currency & money sinks), OQ-010
(durability/decay/repair), OQ-011 (respecialisation cost), OQ-012 / SAFE-002 (multi-account & exploit
implementation policy), OQ-016 (vendor & retail structure), SERV-001 / SERV-002 (services and service
interdependence), and OQ-007 / BLD-004 (city formation, placement, abandonment). Each of these is a
**named-open** or **TBD** surface in the project's own registers. **Answering this instrument does not
answer, close, or prejudge any of them.** It supplies historical observation against a surface the project
has explicitly left open.

**Method:** SWG (Star Wars Galaxies, Pre-CU) is no longer playable. The only remaining record of its
systems is human recollected experience. This instrument is a second pass over the domains where the first
pass (223 questions, `proposals/SWG_Historical_Accuracy-v1.md`) produced **no usable evidence**, using
stems verified against SWG Pre-CU ground truth and the respondent's own vocabulary. Answers are recorded
verbatim with the responder's handle and date, then filed as EVIDENCE. They are never treated as measured
data, never as tuning targets, and never as rulings.

---

## 0. How this instrument differs from v1, and why

The v1 assessment produced four findings that bear directly on how this instrument is built. They are
stated here so that the questions are auditable against them, and so that a later reader can see that the
instrument's *shape* is a consequence of the assessment rather than a preference.

| v1 finding | Consequence for this instrument |
|---|---|
| **F-2 / F-2a** — figure-bearing stems anchor the respondent. Two of v1's three were later corroborated (240h elite grind; the 30–60s figure is a group-PvP figure), but the third (listing-fee percentage) remains uncorroborated. | **No stem in this instrument contains a number, a range, or a figure.** Every quantity is asked for, not suggested. If a quantity is needed, it is a question. |
| **F-3 / F-8** — answer density tracks the **instrument's fidelity to SWG Pre-CU**, not the respondent's expertise. The respondent holds mastery of every class in the game, yet the generic-MMO-stemmed categories came back empty. | **No stem may assert a mechanic that is not attested in SWG Pre-CU.** Stems are written from ground truth or phrased conditionally ("Was there…", "If there was…"). A stem the respondent must spend a word rejecting cannot also yield a description — so none is asked that way. |
| **F-3** — satisfaction-framed stems ("Did X feel right?") produced one-word answers; mechanism-framed stems ("How did X work?") produced prose. | **Mechanism framing is preferred throughout.** Satisfaction is asked only where satisfaction is the thing of interest, and even then the stem asks *what happened*, not *was it good*. |
| **F-9** — the respondent's vocabulary differs from the project's. The term "service XP" is the project's and appears nowhere in the SWG Pre-CU GDD; the respondent could not answer Q90 because of the label alone. | **The respondent's own words are used where known** — vendor, bazaar, Master Merchant, Vendor skill tree, junk merchant, cloning tank, surveying, harvesting, crafting. No project-internal coinage is used as a question label. |
| **F-1** — the confidence field was uniform `high` across all 223 answers and therefore carried no information. | **Confidence is required per item**, and `inference` is a first-class value. A stem may invite an answer the respondent does not actually remember; that is the point of the `inference` option. |
| **Q-C7** — the respondent could not answer a question about a mechanic they had never heard of, and said so. | **"Never used it" / "never saw that" is an explicit answer option on every question**, alongside `TBD` / `unknown`. Silence is a valid and informative answer. |

**Answering rules (restated for the respondent):**

- Answer **from recollection only**. If you are unsure, write **TBD** / **unknown** / **I don't remember** — never guess a number.
- Distinguish "I remember it worked like X" (recollection) from "I think it was probably X" (inference). Mark inference explicitly.
- If you never used a system, say **N/A — never used**. Do not describe a system you have no experience of.
- Numbers are **recollected magnitudes only**. They are never tuning targets and are never proposed as values.
- It is more useful to say *"I don't remember"* than to produce a plausible-sounding answer. A wrong number in this record is worse than a missing one.
- Anything you want to add that no question asked for goes in **Notes**.

**Answer record format** — one block per question, exactly as in v1:

```
### Q<number> — <short question restatement>
- Responder: <handle>
- Date: YYYY-MM-DD
- Answer: <text, or TBD / unknown / N/A>
- Confidence: high | medium | low | inference
- Notes: <optional>
```

---

## A. The one question still open from the report — service XP

*Carried from `investigations/TCIndustries_SWG_Historical_Recollection_Report_Tiwa_Pene_2026-10-21.md`
§8.2 (Q-C7b). Your answer to A1 closes the only open item in the v1 record. The record already holds your
Q91 answer — that service XP scaled with severity, and your stated preference was flat with a scaled
bonus — so A1–A4 are about the parts Q91 did not reach.*

1. You said the Vendor skill tree paid XP per day for each Vendor a merchant had. Did the other service
   professions — Medic, Doctor, Entertainer — earn XP only when they performed a service for someone, or
   did they also earn a standing amount each day for the people and equipment they had set up? In other
   words: was service XP in SWG earned by doing things, or by having things?
2. Was the Vendor tree's daily XP the same amount as the XP other service professions earned, or a
   different rate? And was it paid for every Vendor a merchant had, or only for Vendors that actually sold
   something?
3. Knowing that the Vendor skill tree "severely gates the time it takes to become a Master Merchant" — did
   service XP feel like a fair way to earn a living as a service profession? Could you make it as a
   career on its own, or did you need another way to earn while you levelled?
4. Was service XP the main way service professionals levelled, or was it something extra on top of XP
   from missions or from selling? For a merchant specifically, which was the bigger source of XP?

## B. Durability, decay and repair

*For OQ-010, which is a named-open structural question: universal decay with player repair (A), decay only
for identity-bearing crafted items (B), or no baseline decay with demand coming from destruction and
obsolescence instead (C). In the v1 record you said armour did not decay and that weapons needed repair.
These questions establish the shape of what was actually there, not which option the project should take.*

5. In SWG Pre-CU, did things wear out from use? Which kinds of items wore out, and which did not?
6. When a weapon or piece of equipment wore out, what happened to it? Did it stop working, did it get
   weaker, or did it just need to be fixed by someone?
7. Who fixed items, and what did that involve — skill, materials, credits, time, or a combination?
8. Was there a steady stream of people needing repairs? If so, who made a living from it, and was that
   a good living?
9. Was there ever anything in the game that was permanently destroyed or could no longer be used? If so,
   what was it and how did that happen?

## C. Respecialisation and unlearning skills

*For OQ-011, which is a named-open structural question: economic cost (A), retraining time (B), or
capability decay / surrender (C). In the v1 record you said players could remove skills from the skill
tree. These questions establish what that actually cost and what it enabled.*

10. You said players could remove a skill from the skill tree. Was doing so free, immediate, and repeatable
    as many times as you liked — or did it cost something, such as credits, materials, or time?
11. If a player had already put experience into a skill and then removed it, did they get that experience
    back, or was it gone?
12. Were there limits — a limit per day, a limit in total, a limit on which skills, a cooldown, or a cost
    per removal?
13. How often did people actually do this, and what were the reasons? Was it mostly changing career,
    filling a temporary need, correcting a mistake, or something else?

## D. Vendors and how things were sold

*For OQ-016 (vendor and retail structure) and OQ-009 (pricing mechanics, RET-002). In the v1 record you
described two tiers — Master Merchants could place vendors without fees or caps, while ordinary traders
could only use the bazaar, capped in price and in how much they could list, and were charged fees.*

14. Who could place a vendor, and was that tied to a skill or a profession? Were there different levels of
    vendor, and who could get which?
15. What did it cost to keep a vendor running, and what happened if you stopped paying for it?
16. Was there a limit on how many vendors a single person could run, or on how much they could stock or
    sell through them? What determined the limit?
17. What were the ways you could sell something, and which did you use and why?
18. Were there price limits, quantity limits, or fees on some ways of selling but not others? Which ways,
    and who was exempt from them?
19. When you dealt in the market, what made it work as a real market rather than a dump? What broke it,
    if anything?

## E. Money sinks, money faucets, and the duplication episode

*For OQ-009 — whose sink instrument classes are named-open, each to be adopted by its own future ruling —
and for OQ-012 / SAFE-002 (the multi-account and exploit implementation policy). In the v1 record you said
the only real money sinks were housing and buildings, that harvesters and factories earned more than they
consumed, that travel and training were cheap, that prices were pinned to mission income, and that a
duplication glitch caused inflation which hoarding behaviour then absorbed.*

20. Which things in the game regularly took money out of a player's hands — for an ordinary player, not
    for the wealthy?
21. Which things took money out of a player's hands only if they were already established or wealthy?
22. You said a duplication glitch that some people exploited caused inflation. What happened to the economy
    afterwards? Did prices rise, and if they rose, did they ever come back down?
23. Did the developers do anything about it? Did the player base notice, and did people care?
24. You said players hoarded money and that merchants acted as money sinks. In practice, what did hoarding
    look like — what did people hold on to, and why?
25. What was the most reliable way for a new player to earn money, and what was the most reliable way for
    an established player? Were those the same income?

## F. Creatures and the world they lived in

*For the creature and ecology domain, which the v1 instrument failed to cover because its stems were
generic rather than SWG-specific. You mastered the Bio Engineer career, and creature handling was a
notable class.*

26. How did creatures behave around each other? Did they hunt as a group, fight one at a time, or run away?
27. Before you committed to a fight, how did you know whether a creature was dangerous for you? What
    information was available, and what was not?
28. Tell me about a creature you remember vividly. What made it stick in your memory — and what did it
    actually do when you met it?
29. What did you do with creatures you did not want to fight? Was there a way to tame them, harvest from
    them, avoid them, or drive them off?
30. In a place you knew well, how did the creatures there change over time? Did they come back, move away,
    or thin out?

## G. Items and gear

*For the item and equipment domain, again uncovered by v1 because its stems assumed mechanics SWG did not
have. You mastered Weaponsmithing and Armoursmithing, so this is asking the person who made the things.*

31. As someone who made weapons and armour, what made an item you made worth having to someone else, rather
    than just being another item?
32. Did you ever make something for another player? How did that work — whose materials did you use, and
    how were you paid?
33. When you crafted, what part of the result could you not control? How did you feel about that, and did
    players talk about it the way people talk about a machine's dice?
34. How did quality actually work in practice — what were the ranges you saw, and what did players chase?

## H. Player cities and civic life

*For OQ-007 and BLD-004. The v1 instrument asked this category in generic terms and got one-word answers
from a respondent who has played the system.*

35. What did founding a player city actually involve, step by step?
36. Who paid for the city, and who was allowed to spend its money?
37. What happened to a city when the people running it stopped playing?
38. What did a city need in order to be worth living in, and what made a good one?

## I. Death, cloning, and playing more than one character

*For the death and cloning domain, and for OQ-012 / SAFE-002. The v1 stems for this category asserted
mechanics SWG did not have, so they produced rejections rather than descriptions.*

39. What actually happened to your character when you died in ordinary PvE, and how was that different from
    dying in PvP?
40. Was there a cost to dying at all? If there was, what form did that cost take?
41. Did you ever play more than one character at once? If so, did you use them together or separately, and
    what for?
42. Were there rules about multiple characters, and did people follow them?

## J. Elite professions

*For the elite-profession domain, and because a respondent who mastered every class has a comparative view
of which elites worked and why. In the v1 record you said some elites were "completely broken and never
worked correctly".*

43. Of the elite careers, which one do you think was genuinely well designed, and what made it work?
44. Which one was broken, and what specifically was the failure? Was the problem in the numbers, in the
    skills, in the group requirement, or in the tools?
45. What did an elite career need that an ordinary career did not?
46. Did people play elites differently, or did most people just level them because they were on the way?

## K. Social structure and guilds

*For the social domain, which produced no usable evidence in v1.*

47. Why did people join a guild? What was actually in it for them?
48. Did a guild do anything besides give people a chat channel? Did it enforce anything, and did that
    work?
49. Was there a strong culture of helping new players, and what did that look like in practice?

## L. The client and how you played

*For the client domain, which produced no usable evidence in v1.*

50. How did the client handle the amount of information you had to keep track of during a fight?
51. What was hard to use, and what did you work around instead?
52. Did you ever use add-ons, macros, or third-party tools? What did you use them for?

## M. Design judgement

*Open-ended. These are not instrument questions with a right answer; they are the respondent's own
judgement, recorded as evidence of opinion and explicitly not as design instruction.*

53. If you could change one thing about SWG to make it better today, what would it be?
54. If you were building a new MMO and wanted SWG's strengths, what would you copy — and what would you
    deliberately leave behind?
55. What surprised you most about how the game actually played, compared to how you expected it to?
56. What do you wish the developers had understood about how players actually used the game?

---

## Status of this document

- Content status: **PROPOSED** candidate material. Answers are EVIDENCE (human recollection), not design
  authority and not measured data.
- This document does not change the status of any other document and does not authorize any code, tuning, or
  design decision. It does not answer OQ-009, OQ-010, OQ-011, OQ-012, OQ-016, SERV-002, OQ-007, or BLD-004,
  and an answer in it cannot close one. Those surfaces remain exactly as open as they were.
- The answers are filed as EVIDENCE against `investigations/`, and are cross-referenced from
  `investigations/TCIndustries_SWG_Historical_Recollection_Report_Tiwa_Pene_2026-10-21.md` as Addendum C,
  which is also where the Q-C7b answer from Part A is filed.
- Proxy 4 (Sandbox Validation) remains **DEFERRED** and is not covered here.
- **Transport note.** The local fill-and-export form for this instrument,
  `proposals/SWG_Historical_Answers_v2.html`, was rebuilt before any answer was collected, after two defects
  were found in it: truncated question text, and an export path that failed at run time. Both are recorded at
  **D-3** in the report named above. **This instrument document itself was never affected** — the defects
  were in the rendering step that turns this file into the form, not in this file. The rebuilt form renders
  all 56 questions in full and its export was exercised and verified end to end.

*End of SWG_Historical_Accuracy_v2_Targeted_ReAsk.md (v1.0).*
