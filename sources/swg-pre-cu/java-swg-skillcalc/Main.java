import swg.manager.ProfessionFactory;
import swg.manager.SkillManager;
import swg.model.Character;
import swg.ui.SkillCalculatorWindow;

import javax.swing.*;
import java.awt.*;

/**
 * ════════════════════════════════════════════════════════════════════════════
 *  SWG Pre-CU Skill Calculator — Main Entry Point
 * ════════════════════════════════════════════════════════════════════════════
 *
 * Architecture Overview:
 * ──────────────────────
 * The application follows a layered MVC architecture:
 *
 *   MODEL   : swg.model.Character       — character identity + point pool
 *             swg.model.Profession       — named collection of SkillBoxes
 *             swg.model.SkillBox         — individual skill node (cost, prereqs)
 *
 *   CONTROLLER : swg.manager.SkillManager — enforces all pre-CU rules:
 *                  • 250-point hard cap (Character.spendPoints / refundPoints)
 *                  • Prerequisite chain validation (must learn lower tiers first)
 *                  • Dependent-guard on surrender (cannot drop a prereq of a
 *                    learned skill)
 *                  • Observer pattern via SkillManager.SkillChangeListener
 *                    (model fires events; UI components update themselves)
 *
 *   VIEW    : swg.ui.SkillCalculatorWindow  — main JFrame, tab host
 *             swg.ui.ProfessionTreePanel    — one profession's grid
 *             swg.ui.SkillBoxPanel          — single learnable skill cell
 *             swg.ui.PointPoolBar           — custom 250-pt gauge component
 *
 * Design Patterns:
 * ──────────────────────
 *   • MVC          — model, controller, and view layers are strictly separated.
 *   • Observer     — SkillManager.SkillChangeListener interface; all UI cells
 *                    register themselves and repaint on any skill state change.
 *   • Factory      — ProfessionFactory builds immutable profession trees with
 *                    accurate pre-CU skill data without polluting model classes.
 *   • Registry     — SkillManager maintains an ordered map of professions and
 *                    a flat indexed map of all SkillBoxes for O(1) lookup.
 *   • Facade       — SkillManager exposes a single API (learnSkill, surrenderSkill,
 *                    canLearn, canSurrender) so the UI never manipulates model
 *                    objects directly.
 *
 * Pre-CU Rules Implemented:
 * ──────────────────────
 *   1. Strict 250-point cap shared across all professions.
 *   2. Prerequisite chain: each SkillBox lists the IDs of boxes that must be
 *      learned before it. Novice boxes have no prerequisites.
 *   3. Prevent surrendering a skill if another learned skill depends on it.
 *   4. Full refund on surrender.
 *   5. Master box requires ALL four branch-line capstone boxes (tier 4).
 *
 * ════════════════════════════════════════════════════════════════════════════
 */
public final class Main {

    public static void main(String[] args) {
        // ── Apply native look-and-feel or fall back to Nimbus for dark theming ──
        applyLookAndFeel();

        // ── Swing UI must be built on the Event Dispatch Thread ──────────────
        SwingUtilities.invokeLater(() -> {
            // 1. Create the character (the 250-pt pool holder)
            Character character = new Character("Arik Draysen", "Human");

            // 2. Wire up the SkillManager (controller / model facade)
            SkillManager manager = new SkillManager(character);

            // 3. Register the three starter profession trees via the factory
            manager.registerProfession(ProfessionFactory.buildMarksman());
            manager.registerProfession(ProfessionFactory.buildBrawler());
            manager.registerProfession(ProfessionFactory.buildArtisan());

            // 4. Build and show the main window (the View)
            SkillCalculatorWindow window = new SkillCalculatorWindow(manager);
            window.setVisible(true);
        });
    }

    // ── Look & Feel ───────────────────────────────────────────────────────────

    private static void applyLookAndFeel() {
        // Try Nimbus first for better dark-colour support
        try {
            for (UIManager.LookAndFeelInfo info : UIManager.getInstalledLookAndFeels()) {
                if ("Nimbus".equals(info.getName())) {
                    UIManager.setLookAndFeel(info.getClassName());
                    tweakNimbusDefaults();
                    return;
                }
            }
        } catch (Exception ignored) {}

        // Fall back to system LAF
        try {
            UIManager.setLookAndFeel(UIManager.getSystemLookAndFeelClassName());
        } catch (Exception ignored) {}
    }

    private static void tweakNimbusDefaults() {
        // Push Nimbus toward a dark base so our custom colours read correctly
        UIManager.put("control",        new Color(0x1A2530));
        UIManager.put("info",           new Color(0x1A2530));
        UIManager.put("nimbusBase",     new Color(0x0E1620));
        UIManager.put("nimbusBlueGrey", new Color(0x1E2E3A));
        UIManager.put("nimbusLightBackground", new Color(0x0A1218));
        UIManager.put("text",           new Color(0xCCDDEE));
        UIManager.put("TabbedPane.contentAreaColor", new Color(0x0E1620));
        UIManager.put("TabbedPane.selected",         new Color(0x1A2A38));
        UIManager.put("TabbedPane.background",       new Color(0x12202C));
    }
}
