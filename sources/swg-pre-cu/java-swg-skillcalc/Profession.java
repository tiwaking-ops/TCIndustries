package swg.model;

import java.util.ArrayList;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/**
 * Represents a single profession (e.g. Marksman, Brawler, Artisan).
 * Each profession holds an ordered map of SkillBoxes keyed by their ID.
 *
 * Pre-CU professions had:
 *   - One "Novice" entry box (no prerequisites outside the profession)
 *   - 4 branch tracks (rows) of 4 boxes each
 *   - One "Master" capstone box requiring all 4 branch tracks completed
 *
 * The treeColumn/treeRow fields on each SkillBox drive the GridLayout
 * rendering in the UI tier.
 */
public class Profession {

    private final String id;
    private final String displayName;
    private final String lore;            // Short flavour text
    private final Map<String, SkillBox> skillBoxes; // insertion-ordered

    public Profession(String id, String displayName, String lore) {
        this.id          = id;
        this.displayName = displayName;
        this.lore        = lore;
        this.skillBoxes  = new LinkedHashMap<>();
    }

    /** Register a SkillBox with this profession. Throws if the ID is already registered. */
    public void addSkillBox(SkillBox box) {
        if (skillBoxes.containsKey(box.getId())) {
            throw new IllegalArgumentException("Duplicate SkillBox id: " + box.getId());
        }
        skillBoxes.put(box.getId(), box);
    }

    /** Returns an unmodifiable view of this profession's skill boxes. */
    public Map<String, SkillBox> getSkillBoxes() {
        return Collections.unmodifiableMap(skillBoxes);
    }

    /** Convenience: return boxes as an ordered list. */
    public List<SkillBox> getSkillBoxList() {
        return new ArrayList<>(skillBoxes.values());
    }

    public String getId()          { return id; }
    public String getDisplayName() { return displayName; }
    public String getLore()        { return lore; }

    /** Count total skill points committed to this profession by the character. */
    public int pointsSpent() {
        int total = 0;
        for (SkillBox b : skillBoxes.values()) {
            if (b.isLearned()) total += b.getPointCost();
        }
        return total;
    }

    @Override
    public String toString() {
        return "Profession[" + id + "]";
    }
}
