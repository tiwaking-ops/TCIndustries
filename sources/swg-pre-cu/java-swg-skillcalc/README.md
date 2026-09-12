# SWG Pre-CU Skill Calculator
## Build & Run Instructions

### Requirements
- Java 8+ JDK (any distribution: OpenJDK, Oracle, Adoptium)

### Compile
```bash
# From the project root (where /src lives):
mkdir -p out
javac -source 8 -target 8 -d out $(find src -name "*.java")
```

### Run
```bash
java -cp out Main
```

### IDE Setup (IntelliJ IDEA / Eclipse / VS Code)
1. Create a new **Java project** (no Maven/Gradle needed)
2. Set the **source root** to `src/`
3. The run configuration should point to `Main` as the main class
4. No external dependencies — uses only `java.awt`, `javax.swing`, and the Java 8 standard library

---

## Architecture

### Pattern Summary

| Pattern | Where Used | Purpose |
|---------|-----------|---------|
| MVC | Entire app | `model` package = Model, `SkillManager` = Controller, `ui` package = View |
| Observer | `SkillManager.SkillChangeListener` | Model notifies UI of state changes without coupling |
| Factory | `ProfessionFactory` | Builds immutable profession trees; keeps model clean |
| Registry | `SkillManager.professions` / `allBoxes` maps | O(1) lookup of any skill box by ID |
| Facade | `SkillManager.learnSkill()` / `surrenderSkill()` | Single entry point for all game rule enforcement |

---

## Pre-CU Rules Implemented

1. **250-point hard cap** — shared pool across ALL professions via `Character.spendPoints()`
2. **Prerequisite chain** — each `SkillBox` lists prerequisite IDs; all must be `isLearned()` before spending
3. **Dependent guard on surrender** — cannot drop a skill if another learned skill lists it as a prereq
4. **Full refund** — surrendering a box restores its cost to the pool
5. **Master box lock** — requires all four branch-line Tier 4 boxes to be learned
6. **Cross-profession prereqs** — the `allBoxes` flat index allows any box to reference any other box by ID, enabling future cross-profession prerequisite chains (e.g., Ranger requires Marksman Scout IV)

---

## Profession Data (Starter Set)

| Profession | Boxes | Cost to Master |
|-----------|-------|---------------|
| Marksman  |  18   |    37 pts     |
| Brawler   |  18   |    37 pts     |
| Artisan   |  18   |    37 pts     |
| **All 3** |  54   |   111 pts     |

With 250 pts, a character can master all three starter professions and have **139 points** remaining for elite professions.

---

## File Structure

```
src/
├── Main.java                          ← Entry point, wires everything together
├── swg/
│   ├── model/
│   │   ├── SkillBox.java             ← Immutable skill node (id, cost, prereqs)
│   │   ├── Profession.java           ← Named collection of SkillBoxes
│   │   └── Character.java            ← Player character + 250-pt pool
│   ├── manager/
│   │   ├── SkillManager.java         ← All game logic + Observer events
│   │   └── ProfessionFactory.java    ← Authentic pre-CU skill tree data
│   └── ui/
│       ├── SkillCalculatorWindow.java ← Main JFrame host
│       ├── ProfessionTreePanel.java   ← One profession's grid + connector lines
│       ├── SkillBoxPanel.java         ← Individual skill cell with Learn/Drop
│       └── PointPoolBar.java          ← Custom 250-pt segmented gauge
```
