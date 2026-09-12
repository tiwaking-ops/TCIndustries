package swg.manager;

import swg.model.Profession;
import swg.model.SkillBox;

import java.util.Arrays;
import java.util.Collections;

/**
 * Static factory that constructs the three starter profession trees
 * with authentic pre-CU structure and point costs.
 *
 * Grid layout convention (treeColumn / treeRow):
 *   Col 0 = Novice (entry box)
 *   Col 1-4 = Branch tiers (4 rows × 4 tiers)
 *   Col 5 = Master capstone
 *
 * Point costs follow pre-CU norms:
 *   Novice = 1 pt  |  Branch boxes = 1-4 pts  |  Master = 4 pts
 *
 * Branch rows:
 *   Row 0 = Weapons / Combat / Crafting I
 *   Row 1 = Wounds / Techniques / Crafting II
 *   Row 2 = Support / Titles / Surveying
 *   Row 3 = Speed / Tactics / Business
 */
public final class ProfessionFactory {

    private ProfessionFactory() {} // static utility class

    // ══════════════════════════════════════════════════════════════════════════
    //  MARKSMAN
    // ══════════════════════════════════════════════════════════════════════════
    public static Profession buildMarksman() {
        Profession p = new Profession(
                "marksman",
                "Marksman",
                "Masters of ranged combat, Marksmen train in pistols, carbines, rifles, " +
                "and advanced scouting techniques.");

        // ── Novice ────────────────────────────────────────────────────────────
        p.addSkillBox(new SkillBox(
                "marksman_novice", "Novice Marksman",
                "Entry into ranged combat. Grants basic pistol and rifle training.",
                1, "marksman", 0, 0));

        // ── Branch 0: Pistols (row 0) ─────────────────────────────────────────
        p.addSkillBox(new SkillBox(
                "marksman_pistols_1", "Pistols I",
                "+5% pistol accuracy, unlocks Point Blank Shot.",
                1, "marksman", 1, 0,
                Arrays.asList("marksman_novice")));
        p.addSkillBox(new SkillBox(
                "marksman_pistols_2", "Pistols II",
                "+10% pistol damage, unlocks Quick Draw.",
                2, "marksman", 2, 0,
                Arrays.asList("marksman_pistols_1")));
        p.addSkillBox(new SkillBox(
                "marksman_pistols_3", "Pistols III",
                "+15% pistol speed, unlocks Fan Fire.",
                2, "marksman", 3, 0,
                Arrays.asList("marksman_pistols_2")));
        p.addSkillBox(new SkillBox(
                "marksman_pistols_4", "Pistols IV",
                "Mastery of sidearms. Unlocks Pistol Spin.",
                3, "marksman", 4, 0,
                Arrays.asList("marksman_pistols_3")));

        // ── Branch 1: Carbines (row 1) ────────────────────────────────────────
        p.addSkillBox(new SkillBox(
                "marksman_carbines_1", "Carbines I",
                "+5% carbine accuracy, unlocks Burst Fire.",
                1, "marksman", 1, 1,
                Arrays.asList("marksman_novice")));
        p.addSkillBox(new SkillBox(
                "marksman_carbines_2", "Carbines II",
                "+10% carbine damage, unlocks Cone of Fire.",
                2, "marksman", 2, 1,
                Arrays.asList("marksman_carbines_1")));
        p.addSkillBox(new SkillBox(
                "marksman_carbines_3", "Carbines III",
                "+15% carbine speed, +5% dodge penalty to target.",
                2, "marksman", 3, 1,
                Arrays.asList("marksman_carbines_2")));
        p.addSkillBox(new SkillBox(
                "marksman_carbines_4", "Carbines IV",
                "Mastery of assault carbines. +25% damage.",
                3, "marksman", 4, 1,
                Arrays.asList("marksman_carbines_3")));

        // ── Branch 2: Rifles (row 2) ──────────────────────────────────────────
        p.addSkillBox(new SkillBox(
                "marksman_rifles_1", "Rifles I",
                "+5% rifle accuracy, unlocks Sniper Shot.",
                1, "marksman", 1, 2,
                Arrays.asList("marksman_novice")));
        p.addSkillBox(new SkillBox(
                "marksman_rifles_2", "Rifles II",
                "+10% rifle range, unlocks Leg Shot.",
                2, "marksman", 2, 2,
                Arrays.asList("marksman_rifles_1")));
        p.addSkillBox(new SkillBox(
                "marksman_rifles_3", "Rifles III",
                "+15% rifle damage, unlocks Head Shot.",
                2, "marksman", 3, 2,
                Arrays.asList("marksman_rifles_2")));
        p.addSkillBox(new SkillBox(
                "marksman_rifles_4", "Rifles IV",
                "Mastery of precision rifles. Unlocks Deadeye.",
                3, "marksman", 4, 2,
                Arrays.asList("marksman_rifles_3")));

        // ── Branch 3: Scout / Camouflage (row 3) ─────────────────────────────
        p.addSkillBox(new SkillBox(
                "marksman_scout_1", "Scout I",
                "Increases movement speed, reduces detection radius.",
                1, "marksman", 1, 3,
                Arrays.asList("marksman_novice")));
        p.addSkillBox(new SkillBox(
                "marksman_scout_2", "Scout II",
                "Unlocks Camouflage ability, +10% concealment.",
                2, "marksman", 2, 3,
                Arrays.asList("marksman_scout_1")));
        p.addSkillBox(new SkillBox(
                "marksman_scout_3", "Scout III",
                "Unlocks Lay Trap, +15% ambush damage.",
                2, "marksman", 3, 3,
                Arrays.asList("marksman_scout_2")));
        p.addSkillBox(new SkillBox(
                "marksman_scout_4", "Scout IV",
                "Master Scout. Unlocks Conceal, near-invisibility.",
                3, "marksman", 4, 3,
                Arrays.asList("marksman_scout_3")));

        // ── Master ────────────────────────────────────────────────────────────
        p.addSkillBox(new SkillBox(
                "marksman_master", "Master Marksman",
                "Supreme ranged mastery. Unlocks Overcharge Shot. All ranged +20%.",
                4, "marksman", 5, 0,
                Arrays.asList("marksman_pistols_4", "marksman_carbines_4",
                               "marksman_rifles_4", "marksman_scout_4")));

        return p;
    }

    // ══════════════════════════════════════════════════════════════════════════
    //  BRAWLER
    // ══════════════════════════════════════════════════════════════════════════
    public static Profession buildBrawler() {
        Profession p = new Profession(
                "brawler",
                "Brawler",
                "The foundation of melee combat. Brawlers excel in one-handed, " +
                "two-handed, unarmed, and polearm disciplines.");

        // Novice
        p.addSkillBox(new SkillBox(
                "brawler_novice", "Novice Brawler",
                "Entry into melee combat. Grants basic unarmed and blade training.",
                1, "brawler", 0, 0));

        // Branch 0: One-Hand (row 0)
        p.addSkillBox(new SkillBox(
                "brawler_1h_1", "One-Hand I",
                "+5% one-handed accuracy. Unlocks Slash.",
                1, "brawler", 1, 0,
                Arrays.asList("brawler_novice")));
        p.addSkillBox(new SkillBox(
                "brawler_1h_2", "One-Hand II",
                "+10% one-handed damage. Unlocks Lunge.",
                2, "brawler", 2, 0,
                Arrays.asList("brawler_1h_1")));
        p.addSkillBox(new SkillBox(
                "brawler_1h_3", "One-Hand III",
                "+15% attack speed. Unlocks Dizzy.",
                2, "brawler", 3, 0,
                Arrays.asList("brawler_1h_2")));
        p.addSkillBox(new SkillBox(
                "brawler_1h_4", "One-Hand IV",
                "Mastery of blades. Unlocks Death Blow.",
                3, "brawler", 4, 0,
                Arrays.asList("brawler_1h_3")));

        // Branch 1: Two-Hand (row 1)
        p.addSkillBox(new SkillBox(
                "brawler_2h_1", "Two-Hand I",
                "+5% two-handed accuracy. Unlocks Cleave.",
                1, "brawler", 1, 1,
                Arrays.asList("brawler_novice")));
        p.addSkillBox(new SkillBox(
                "brawler_2h_2", "Two-Hand II",
                "+10% two-handed damage. Unlocks Whirlwind.",
                2, "brawler", 2, 1,
                Arrays.asList("brawler_2h_1")));
        p.addSkillBox(new SkillBox(
                "brawler_2h_3", "Two-Hand III",
                "+15% damage on stunned targets.",
                2, "brawler", 3, 1,
                Arrays.asList("brawler_2h_2")));
        p.addSkillBox(new SkillBox(
                "brawler_2h_4", "Two-Hand IV",
                "Mastery of two-handed weapons. Unlocks Crushing Blow.",
                3, "brawler", 4, 1,
                Arrays.asList("brawler_2h_3")));

        // Branch 2: Unarmed (row 2)
        p.addSkillBox(new SkillBox(
                "brawler_unarmed_1", "Unarmed I",
                "+5% unarmed accuracy. Unlocks Kick.",
                1, "brawler", 1, 2,
                Arrays.asList("brawler_novice")));
        p.addSkillBox(new SkillBox(
                "brawler_unarmed_2", "Unarmed II",
                "+10% unarmed damage. Unlocks Jab.",
                2, "brawler", 2, 2,
                Arrays.asList("brawler_unarmed_1")));
        p.addSkillBox(new SkillBox(
                "brawler_unarmed_3", "Unarmed III",
                "+15% knockdown chance. Unlocks Pummel.",
                2, "brawler", 3, 2,
                Arrays.asList("brawler_unarmed_2")));
        p.addSkillBox(new SkillBox(
                "brawler_unarmed_4", "Unarmed IV",
                "Mastery of hand-to-hand. Unlocks Coup de Grâce.",
                3, "brawler", 4, 2,
                Arrays.asList("brawler_unarmed_3")));

        // Branch 3: Polearm (row 3)
        p.addSkillBox(new SkillBox(
                "brawler_polearm_1", "Polearm I",
                "+5% polearm accuracy. Unlocks Sweep.",
                1, "brawler", 1, 3,
                Arrays.asList("brawler_novice")));
        p.addSkillBox(new SkillBox(
                "brawler_polearm_2", "Polearm II",
                "+10% polearm damage. Unlocks Stun Strike.",
                2, "brawler", 2, 3,
                Arrays.asList("brawler_polearm_1")));
        p.addSkillBox(new SkillBox(
                "brawler_polearm_3", "Polearm III",
                "+20% reach advantage. Unlocks Knock Back.",
                2, "brawler", 3, 3,
                Arrays.asList("brawler_polearm_2")));
        p.addSkillBox(new SkillBox(
                "brawler_polearm_4", "Polearm IV",
                "Mastery of polearms. Unlocks Force Wave.",
                3, "brawler", 4, 3,
                Arrays.asList("brawler_polearm_3")));

        // Master
        p.addSkillBox(new SkillBox(
                "brawler_master", "Master Brawler",
                "The apex of melee combat. All melee damage +20%. Unlocks Battle Cry.",
                4, "brawler", 5, 0,
                Arrays.asList("brawler_1h_4", "brawler_2h_4",
                               "brawler_unarmed_4", "brawler_polearm_4")));

        return p;
    }

    // ══════════════════════════════════════════════════════════════════════════
    //  ARTISAN
    // ══════════════════════════════════════════════════════════════════════════
    public static Profession buildArtisan() {
        Profession p = new Profession(
                "artisan",
                "Artisan",
                "The crafter's foundation. Artisans are masters of surveying resources, " +
                "engineering goods, and business acumen.");

        // Novice
        p.addSkillBox(new SkillBox(
                "artisan_novice", "Novice Artisan",
                "Entry into crafting. Grants basic tool use and survey ability.",
                1, "artisan", 0, 0));

        // Branch 0: Engineering (row 0)
        p.addSkillBox(new SkillBox(
                "artisan_eng_1", "Engineering I",
                "+5% crafting speed. Enables Tool Upgrade I.",
                1, "artisan", 1, 0,
                Arrays.asList("artisan_novice")));
        p.addSkillBox(new SkillBox(
                "artisan_eng_2", "Engineering II",
                "+10% crafting quality. Unlocks Advanced Components.",
                2, "artisan", 2, 0,
                Arrays.asList("artisan_eng_1")));
        p.addSkillBox(new SkillBox(
                "artisan_eng_3", "Engineering III",
                "+15% experimentation points per assembly.",
                2, "artisan", 3, 0,
                Arrays.asList("artisan_eng_2")));
        p.addSkillBox(new SkillBox(
                "artisan_eng_4", "Engineering IV",
                "Master tier crafting. Enables Prototype Station use.",
                3, "artisan", 4, 0,
                Arrays.asList("artisan_eng_3")));

        // Branch 1: Domestic Arts (row 1)
        p.addSkillBox(new SkillBox(
                "artisan_domestic_1", "Domestic Arts I",
                "Unlocks food and clothing crafting schematics.",
                1, "artisan", 1, 1,
                Arrays.asList("artisan_novice")));
        p.addSkillBox(new SkillBox(
                "artisan_domestic_2", "Domestic Arts II",
                "Unlocks furniture and décor schematics. +10% quality.",
                2, "artisan", 2, 1,
                Arrays.asList("artisan_domestic_1")));
        p.addSkillBox(new SkillBox(
                "artisan_domestic_3", "Domestic Arts III",
                "Unlocks advanced fabrics and exotic ingredient use.",
                2, "artisan", 3, 1,
                Arrays.asList("artisan_domestic_2")));
        p.addSkillBox(new SkillBox(
                "artisan_domestic_4", "Domestic Arts IV",
                "Mastery of domestic crafting. Unlocks Luxury Goods.",
                3, "artisan", 4, 1,
                Arrays.asList("artisan_domestic_3")));

        // Branch 2: Surveying (row 2)
        p.addSkillBox(new SkillBox(
                "artisan_survey_1", "Surveying I",
                "Increases survey tool range by 15%.",
                1, "artisan", 1, 2,
                Arrays.asList("artisan_novice")));
        p.addSkillBox(new SkillBox(
                "artisan_survey_2", "Surveying II",
                "Unlocks Mineral Surveying. +10% resource concentration reading.",
                2, "artisan", 2, 2,
                Arrays.asList("artisan_survey_1")));
        p.addSkillBox(new SkillBox(
                "artisan_survey_3", "Surveying III",
                "Unlocks Remote Sampling. +20% extraction efficiency.",
                2, "artisan", 3, 2,
                Arrays.asList("artisan_survey_2")));
        p.addSkillBox(new SkillBox(
                "artisan_survey_4", "Surveying IV",
                "Master Surveyor. Pinpoint rare resource nodes.",
                3, "artisan", 4, 2,
                Arrays.asList("artisan_survey_3")));

        // Branch 3: Business (row 3)
        p.addSkillBox(new SkillBox(
                "artisan_business_1", "Business I",
                "Vendor terminal access. Reduced broker fees.",
                1, "artisan", 1, 3,
                Arrays.asList("artisan_novice")));
        p.addSkillBox(new SkillBox(
                "artisan_business_2", "Business II",
                "Unlocks Bazaar Terminal price alerts. 2 vendor slots.",
                2, "artisan", 2, 3,
                Arrays.asList("artisan_business_1")));
        p.addSkillBox(new SkillBox(
                "artisan_business_3", "Business III",
                "Unlocks Merchant Certification. 3 vendor slots.",
                2, "artisan", 3, 3,
                Arrays.asList("artisan_business_2")));
        p.addSkillBox(new SkillBox(
                "artisan_business_4", "Business IV",
                "Master Trader. Maximum vendor capacity, zero broker fees.",
                3, "artisan", 4, 3,
                Arrays.asList("artisan_business_3")));

        // Master
        p.addSkillBox(new SkillBox(
                "artisan_master", "Master Artisan",
                "The pinnacle of crafting. All quality +25%. Enables Elite Profession access.",
                4, "artisan", 5, 0,
                Arrays.asList("artisan_eng_4", "artisan_domestic_4",
                               "artisan_survey_4", "artisan_business_4")));

        return p;
    }
}
