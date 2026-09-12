package swg.model;

/**
 * Represents the player character.
 * The Character is a thin value-object: it stores identity data and
 * delegates all skill-progression logic to {@link swg.manager.SkillManager}.
 *
 * Pre-CU hard cap: 250 skill points total, shared across ALL professions.
 */
public class Character {

    public static final int MAX_SKILL_POINTS = 250;

    private final String name;
    private final String species;
    private int spentPoints = 0;       // Points currently committed

    public Character(String name, String species) {
        if (name == null || name.isEmpty())   throw new IllegalArgumentException("Character name required");
        if (species == null || species.isEmpty()) throw new IllegalArgumentException("Species required");
        this.name    = name;
        this.species = species;
    }

    // ── Point pool management ─────────────────────────────────────────────────

    /** Points remaining in the 250-point pool. */
    public int getRemainingPoints() {
        return MAX_SKILL_POINTS - spentPoints;
    }

    public int getSpentPoints() { return spentPoints; }

    /**
     * Attempt to spend {@code cost} points.
     * @return true if successful; false if the pool would be exceeded.
     */
    public boolean spendPoints(int cost) {
        if (cost < 0) throw new IllegalArgumentException("Cannot spend a negative cost");
        if (spentPoints + cost > MAX_SKILL_POINTS) return false;
        spentPoints += cost;
        return true;
    }

    /**
     * Refund {@code cost} points back to the pool.
     * @throws IllegalStateException if refund would push spent below zero.
     */
    public void refundPoints(int cost) {
        if (cost < 0) throw new IllegalArgumentException("Cannot refund a negative cost");
        if (spentPoints - cost < 0) throw new IllegalStateException("Refund exceeds spent points");
        spentPoints -= cost;
    }

    // ── Identity ──────────────────────────────────────────────────────────────

    public String getName()    { return name; }
    public String getSpecies() { return species; }

    @Override
    public String toString() {
        return String.format("Character[%s (%s) – %d/%d pts used]",
                name, species, spentPoints, MAX_SKILL_POINTS);
    }
}
