package swg.manager;

import swg.model.Character;
import swg.model.Profession;
import swg.model.SkillBox;

import java.util.ArrayList;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/**
 * SkillManager acts as the central Model-Controller (in our MVC split).
 *
 * Responsibilities:
 *  1. Maintain the registry of all {@link Profession}s (and their {@link SkillBox}es).
 *  2. Enforce pre-CU rules when learning or surrendering a skill:
 *       - 250-point hard cap (delegated to {@link Character#spendPoints}).
 *       - Prerequisite chain satisfaction (all prereq boxes must be learned).
 *       - Dependent-guard on surrender (cannot surrender a box another
 *         learned box depends on).
 *  3. Fire {@link SkillChangeListener} events so the Swing UI can refresh
 *     without the model knowing anything about UI components (Observer pattern).
 *
 * Design patterns used:
 *  - Observer  : SkillChangeListener interface
 *  - Registry  : professions map acts as a simple service registry
 *  - Facade    : single entry-point for all skill operations
 */
public class SkillManager {

    // ── Observer interface ────────────────────────────────────────────────────

    /**
     * Implemented by any component that needs to react to skill-state changes.
     */
    public interface SkillChangeListener {
        /**
         * @param box    The SkillBox whose learned state changed.
         * @param character The character whose pool changed.
         */
        void onSkillChanged(SkillBox box, Character character);
    }

    // ── State ─────────────────────────────────────────────────────────────────

    private final Character character;
    private final Map<String, Profession> professions = new LinkedHashMap<>(); // ordered insertion
    private final Map<String, SkillBox>   allBoxes    = new LinkedHashMap<>(); // flattened index
    private final List<SkillChangeListener> listeners = new ArrayList<>();

    // ── Construction ──────────────────────────────────────────────────────────

    public SkillManager(Character character) {
        if (character == null) throw new IllegalArgumentException("Character required");
        this.character = character;
    }

    // ── Registry ──────────────────────────────────────────────────────────────

    /** Register a profession and index all its skill boxes. */
    public void registerProfession(Profession profession) {
        if (professions.containsKey(profession.getId())) {
            throw new IllegalArgumentException("Duplicate profession id: " + profession.getId());
        }
        professions.put(profession.getId(), profession);
        for (SkillBox box : profession.getSkillBoxList()) {
            if (allBoxes.containsKey(box.getId())) {
                throw new IllegalArgumentException("Duplicate SkillBox id across professions: " + box.getId());
            }
            allBoxes.put(box.getId(), box);
        }
    }

    public Map<String, Profession> getProfessions() {
        return Collections.unmodifiableMap(professions);
    }

    public List<Profession> getProfessionList() {
        return new ArrayList<>(professions.values());
    }

    public SkillBox getSkillBox(String id) {
        return allBoxes.get(id);
    }

    public Character getCharacter() { return character; }

    // ── Observer management ───────────────────────────────────────────────────

    public void addSkillChangeListener(SkillChangeListener l) {
        if (l != null) listeners.add(l);
    }

    public void removeSkillChangeListener(SkillChangeListener l) {
        listeners.remove(l);
    }

    private void fireSkillChanged(SkillBox box) {
        for (SkillChangeListener l : listeners) {
            l.onSkillChanged(box, character);
        }
    }

    // ── Core pre-CU skill logic ───────────────────────────────────────────────

    /**
     * Attempt to learn a skill box.
     *
     * Enforces:
     *  1. Not already learned.
     *  2. All prerequisites are learned.
     *  3. Enough points remain in the 250-pt pool.
     *
     * @return A human-readable result message (success or reason for failure).
     */
    public String learnSkill(String boxId) {
        SkillBox box = allBoxes.get(boxId);
        if (box == null)       return "Unknown skill: " + boxId;
        if (box.isLearned())   return "\"" + box.getDisplayName() + "\" is already learned.";

        // ── Prerequisite check ────────────────────────────────────────────────
        for (String prereqId : box.getPrerequisiteIds()) {
            SkillBox prereq = allBoxes.get(prereqId);
            if (prereq == null || !prereq.isLearned()) {
                String prereqName = (prereq != null) ? prereq.getDisplayName() : prereqId;
                return "Requires \"" + prereqName + "\" first.";
            }
        }

        // ── Point cap check ───────────────────────────────────────────────────
        if (!character.spendPoints(box.getPointCost())) {
            return "Not enough skill points. Need " + box.getPointCost()
                    + ", have " + character.getRemainingPoints() + ".";
        }

        box.setLearned(true);
        fireSkillChanged(box);
        return "Learned \"" + box.getDisplayName() + "\" (–" + box.getPointCost() + " pts).";
    }

    /**
     * Attempt to surrender (un-learn) a skill box.
     *
     * Enforces:
     *  1. Must be currently learned.
     *  2. No other learned box lists this box as a prerequisite.
     *
     * @return A human-readable result message.
     */
    public String surrenderSkill(String boxId) {
        SkillBox box = allBoxes.get(boxId);
        if (box == null)       return "Unknown skill: " + boxId;
        if (!box.isLearned())  return "\"" + box.getDisplayName() + "\" is not learned.";

        // ── Dependency guard ──────────────────────────────────────────────────
        for (SkillBox other : allBoxes.values()) {
            if (other.isLearned() && other.getPrerequisiteIds().contains(boxId)) {
                return "Cannot surrender: \"" + other.getDisplayName() + "\" depends on it.";
            }
        }

        box.setLearned(false);
        character.refundPoints(box.getPointCost());
        fireSkillChanged(box);
        return "Surrendered \"" + box.getDisplayName() + "\" (+" + box.getPointCost() + " pts refunded).";
    }

    /**
     * Convenience: can the given box be learned right now?
     * Used by the UI to enable/disable buttons.
     */
    public boolean canLearn(String boxId) {
        SkillBox box = allBoxes.get(boxId);
        if (box == null || box.isLearned()) return false;
        if (character.getRemainingPoints() < box.getPointCost()) return false;
        for (String prereqId : box.getPrerequisiteIds()) {
            SkillBox prereq = allBoxes.get(prereqId);
            if (prereq == null || !prereq.isLearned()) return false;
        }
        return true;
    }

    /**
     * Convenience: can the given box be surrendered right now?
     */
    public boolean canSurrender(String boxId) {
        SkillBox box = allBoxes.get(boxId);
        if (box == null || !box.isLearned()) return false;
        for (SkillBox other : allBoxes.values()) {
            if (other.isLearned() && other.getPrerequisiteIds().contains(boxId)) return false;
        }
        return true;
    }
}
