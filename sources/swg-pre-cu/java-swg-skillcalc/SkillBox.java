package swg.model;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/**
 * Represents a single skill box in a profession tree.
 * Pre-CU SWG skill boxes had a flat point cost (typically 1-4 pts)
 * and strict prerequisite chains within/across professions.
 */
public class SkillBox {

    private final String id;           // Unique identifier, e.g. "marksman_novice"
    private final String displayName;  // Human-readable name
    private final String description;  // Flavour/mechanic description
    private final int pointCost;       // Points consumed from the 250-pt pool
    private final List<String> prerequisiteIds; // IDs of boxes that must be learned first
    private final String professionId; // Parent profession
    private final int treeColumn;      // Column position in the UI grid (0=Novice, 1..4=branches, 5=Master)
    private final int treeRow;         // Row position in the UI grid

    // Runtime state — not part of the immutable definition
    private boolean learned = false;

    public SkillBox(String id, String displayName, String description,
                    int pointCost, String professionId,
                    int treeColumn, int treeRow,
                    List<String> prerequisiteIds) {
        if (id == null || id.isEmpty()) throw new IllegalArgumentException("SkillBox id must not be null/empty");
        if (pointCost < 0)             throw new IllegalArgumentException("Point cost cannot be negative");

        this.id              = id;
        this.displayName     = displayName;
        this.description     = description;
        this.pointCost       = pointCost;
        this.professionId    = professionId;
        this.treeColumn      = treeColumn;
        this.treeRow         = treeRow;
        this.prerequisiteIds = Collections.unmodifiableList(new ArrayList<>(prerequisiteIds));
    }

    // ── Convenience constructor with no prerequisites ─────────────────────────
    public SkillBox(String id, String displayName, String description,
                    int pointCost, String professionId,
                    int treeColumn, int treeRow) {
        this(id, displayName, description, pointCost, professionId, treeColumn, treeRow,
             Collections.emptyList());
    }

    // ── Accessors ─────────────────────────────────────────────────────────────

    public String getId()                   { return id; }
    public String getDisplayName()          { return displayName; }
    public String getDescription()          { return description; }
    public int    getPointCost()            { return pointCost; }
    public String getProfessionId()         { return professionId; }
    public int    getTreeColumn()           { return treeColumn; }
    public int    getTreeRow()              { return treeRow; }
    public List<String> getPrerequisiteIds(){ return prerequisiteIds; }

    public boolean isLearned()              { return learned; }
    public void    setLearned(boolean b)    { this.learned = b; }

    @Override
    public String toString() {
        return String.format("SkillBox[%s, cost=%d, learned=%b]", id, pointCost, learned);
    }
}
