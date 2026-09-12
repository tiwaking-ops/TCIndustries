package swg.ui;

import swg.manager.SkillManager;
import swg.model.Character;
import swg.model.SkillBox;

import javax.swing.*;
import javax.swing.border.EmptyBorder;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.MouseAdapter;
import java.awt.event.MouseEvent;

/**
 * A composite Swing panel representing one SkillBox in the tree grid.
 *
 * Visual states:
 *   LEARNED     – deep amber/gold, white text
 *   AVAILABLE   – dark grey, bright text
 *   LOCKED      – very dark, dim text (prerequisites not met)
 *   HOVER       – slightly brightened border on mouseover
 *
 * Each panel fires learn/surrender through {@link SkillManager} and then
 * refreshes itself; a {@link SkillManager.SkillChangeListener} on the
 * parent panel triggers broader UI refreshes.
 */
public class SkillBoxPanel extends JPanel implements SkillManager.SkillChangeListener {

    // ── Palette ───────────────────────────────────────────────────────────────
    private static final Color COL_LEARNED_BG     = new Color(0xC8860A); // Amber
    private static final Color COL_LEARNED_HOVER  = new Color(0xE09A10);
    private static final Color COL_AVAILABLE_BG   = new Color(0x2A3A4A);
    private static final Color COL_AVAILABLE_HOVER = new Color(0x3A4F62);
    private static final Color COL_LOCKED_BG      = new Color(0x1A1F25);
    private static final Color COL_MASTER_BG      = new Color(0x8B0000); // Deep crimson for master
    private static final Color COL_MASTER_LEARNED  = new Color(0xCC2200);
    private static final Color COL_TEXT_BRIGHT     = new Color(0xEEDEBB);
    private static final Color COL_TEXT_DIM        = new Color(0x555E66);
    private static final Color COL_BORDER_LEARNED  = new Color(0xFFD700);
    private static final Color COL_BORDER_NORMAL   = new Color(0x354555);
    private static final Color COL_BORDER_LOCKED   = new Color(0x222830);

    private static final Font FONT_NAME  = new Font("SansSerif", Font.BOLD, 11);
    private static final Font FONT_COST  = new Font("SansSerif", Font.PLAIN, 10);
    private static final Font FONT_NOVICE = new Font("SansSerif", Font.BOLD, 10);

    private final SkillBox box;
    private final SkillManager manager;
    private final boolean isMasterBox;

    private final JLabel nameLabel;
    private final JLabel costLabel;
    private final JButton actionButton;

    private boolean hovered = false;

    public SkillBoxPanel(SkillBox box, SkillManager manager) {
        this.box       = box;
        this.manager   = manager;
        this.isMasterBox = box.getId().endsWith("_master") || box.getId().endsWith("_novice");

        setLayout(new BorderLayout(2, 2));
        setBorder(BorderFactory.createCompoundBorder(
                BorderFactory.createLineBorder(COL_BORDER_NORMAL, 1),
                new EmptyBorder(4, 6, 4, 6)
        ));
        setPreferredSize(new Dimension(130, 72));
        setOpaque(true);

        // Name label
        nameLabel = new JLabel("<html><center>" + box.getDisplayName() + "</center></html>",
                               SwingConstants.CENTER);
        nameLabel.setFont(isMasterBox ? FONT_NOVICE : FONT_NAME);
        nameLabel.setForeground(COL_TEXT_BRIGHT);
        add(nameLabel, BorderLayout.CENTER);

        // Cost label
        costLabel = new JLabel(box.getPointCost() + " pt" + (box.getPointCost() != 1 ? "s" : ""),
                               SwingConstants.CENTER);
        costLabel.setFont(FONT_COST);
        add(costLabel, BorderLayout.SOUTH);

        // Action button (compact)
        actionButton = new JButton("Learn");
        actionButton.setFont(new Font("SansSerif", Font.BOLD, 9));
        actionButton.setMargin(new Insets(1, 4, 1, 4));
        actionButton.setFocusPainted(false);
        actionButton.setCursor(Cursor.getPredefinedCursor(Cursor.HAND_CURSOR));
        actionButton.addActionListener((ActionEvent e) -> handleAction());
        add(actionButton, BorderLayout.NORTH);

        // Hover effect
        MouseAdapter hoverAdapter = new MouseAdapter() {
            @Override public void mouseEntered(MouseEvent e) { hovered = true;  refreshVisuals(); }
            @Override public void mouseExited(MouseEvent e)  { hovered = false; refreshVisuals(); }
        };
        addMouseListener(hoverAdapter);

        // Tooltip with full description
        setToolTipText("<html><b>" + box.getDisplayName() + "</b> (" + box.getPointCost()
                + " pt)<br>" + box.getDescription() + "</html>");

        manager.addSkillChangeListener(this);
        refreshVisuals();
    }

    // ── Action handler ────────────────────────────────────────────────────────

    private void handleAction() {
        String msg;
        if (box.isLearned()) {
            msg = manager.surrenderSkill(box.getId());
        } else {
            msg = manager.learnSkill(box.getId());
        }
        // Result is surfaced by the parent window's status bar via the listener.
        // Broadcast directly here so the status bar can pick it up via a secondary
        // listener registered on the manager.
        manager.addSkillChangeListener((b, c) -> {}); // no-op, just ensures the parent fires
        // The parent SkillCalculatorWindow also registers as a listener and handles status.
        firePropertyChange("skillActionMessage", null, msg);
    }

    // ── SkillChangeListener ───────────────────────────────────────────────────

    @Override
    public void onSkillChanged(SkillBox changedBox, Character character) {
        // Re-render whenever ANY skill changes (prerequisites for this box may now be met).
        refreshVisuals();
    }

    // ── Visual update ─────────────────────────────────────────────────────────

    private void refreshVisuals() {
        boolean learned   = box.isLearned();
        boolean canLearn  = manager.canLearn(box.getId());
        boolean canSurr   = manager.canSurrender(box.getId());

        Color bg;
        Color textColor;
        Color borderColor;

        if (learned) {
            bg = (isMasterBox ? COL_MASTER_LEARNED : (hovered ? COL_LEARNED_HOVER : COL_LEARNED_BG));
            textColor   = Color.WHITE;
            borderColor = COL_BORDER_LEARNED;
            actionButton.setText("Drop");
            actionButton.setBackground(new Color(0x7A1010));
            actionButton.setForeground(new Color(0xFFCCCC));
            actionButton.setEnabled(canSurr);
        } else if (canLearn) {
            bg = (hovered ? COL_AVAILABLE_HOVER : COL_AVAILABLE_BG);
            textColor   = COL_TEXT_BRIGHT;
            borderColor = COL_BORDER_NORMAL;
            actionButton.setText("Learn");
            actionButton.setBackground(new Color(0x1A4A2A));
            actionButton.setForeground(new Color(0x88FFAA));
            actionButton.setEnabled(true);
        } else {
            bg = (isMasterBox ? new Color(0x3A1A1A) : COL_LOCKED_BG);
            textColor   = COL_TEXT_DIM;
            borderColor = COL_BORDER_LOCKED;
            actionButton.setText("Learn");
            actionButton.setBackground(new Color(0x202428));
            actionButton.setForeground(COL_TEXT_DIM);
            actionButton.setEnabled(false);
        }

        setBackground(bg);
        nameLabel.setForeground(textColor);
        costLabel.setForeground(textColor.darker().darker().brighter()); // subtle cost

        setBorder(BorderFactory.createCompoundBorder(
                BorderFactory.createLineBorder(borderColor, learned ? 2 : 1),
                new EmptyBorder(4, 6, 4, 6)
        ));

        // Cost label colour
        costLabel.setForeground(learned ? new Color(0xFFE680) : (canLearn ? new Color(0x8899AA) : COL_TEXT_DIM));

        repaint();
    }

    /** Allow external components to remove the listener when this panel is discarded. */
    public void dispose() {
        manager.removeSkillChangeListener(this);
    }
}
