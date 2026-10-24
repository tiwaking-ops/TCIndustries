# AGENTS.md — rules for tcgui

1. Work on ONE task card per session (docs/tasks/T-xx.md). Read the card, this file, docs/interfaces.md.
2. Run ./scripts/check.sh before finishing. Paste the summary in your report.
3. Do not change frozen interfaces without an ADR and an INTERFACE-CHANGE note.
4. Allowed crates are listed in docs/adr/0001-stack.md. Do not add others without an ADR.
5. No unsafe code without an ADR.
6. Never commit SWG-owned assets or code from community mods. extracted-swg/ and references/ are gitignored.
7. Never invent facts about Star Wars Galaxies. Mark unmeasured values "assumed".
8. If blocked: stop and write BLOCKED: <reason>, <tried>, <needed> in your task report.
9. Keep tasks small; split if the change exceeds about 600 lines.
10. Do not weaken tests to make them pass. Report the failure instead.
11. Report: what changed, commands run with results, assumptions, follow-ups.
