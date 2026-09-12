package swg.ui;

import swg.manager.SkillManager;
import swg.model.Profession;
import swg.model.SkillBox;

import javax.swing.*;
import javax.swing.border.EmptyBorder;
import java.awt.*;
import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.util.ArrayList;
import java.util.List;

/**
 * Renders the skill-tree grid for one {@link Profession}.
 *
 * Layout: the tree is a 4-row × 6-column grid:
 *
 *   Col 0: Novice (spans all 4 rows visually, placed at row 0)
 *   Col 1-4: Branch tiers (row 0-3)
 *   Col 5: Master (spans all 4 rows, placed at row 0)
 *
 * Connector arrows between columns are painted by the custom paintComponent.
 */
public class ProfessionTreePanel extends JPanel implements PropertyChangeListener {

    private static final Color COL_PANEL_BG    = new Color(0x0E1318);
    private static final Color COL_HEADER_BG   = new Color(0x14202C);
    private static final Color COL_HEADER_TEXT  = new Color(0xC8A85A);
    private static final Color COL_LORE_TEXT    = new Color(0x556677);
    private static final Color COL_CONNECTOR    = new Color(0x2A4055);
    private static final Color COL_CONNECTOR_ON = new Color(0xC8860A);

    private static final Font FONT_HEADER = new Font("Serif", Font.BOLD, 16);
    private static final Font FONT_LORE   = new Font("SansSerif", Font.ITALIC, 11);

    private final Profession profession;
    private final SkillManager manager;
    private final List<SkillBoxPanel> boxPanels = new ArrayList<>();

    // Grid: [col][row] → SkillBoxPanel (some cells empty)
    // col: 0=Novice, 1-4=branch tiers, 5=Master
    // row: 0-3
    private final SkillBoxPanel[][] grid = new SkillBoxPanel[6][4];

    // For painting connectors we need the grid panel reference
    private JPanel gridPanel;

    public ProfessionTreePanel(Profession profession, SkillManager manager,
                               SkillManager.SkillChangeListener externalListener) {
        this.profession = profession;
        this.manager    = manager;

        setLayout(new BorderLayout(0, 8));
        setBackground(COL_PANEL_BG);
        setBorder(new EmptyBorder(12, 12, 12, 12));

        // ── Header ────────────────────────────────────────────────────────────
        JPanel header = new JPanel(new BorderLayout(0, 2));
        header.setBackground(COL_HEADER_BG);
        header.setBorder(BorderFactory.createCompoundBorder(
                BorderFactory.createMatteBorder(0, 0, 1, 0, COL_CONNECTOR),
                new EmptyBorder(8, 12, 8, 12)));

        JLabel title = new JLabel(profession.getDisplayName().toUpperCase());
        title.setFont(FONT_HEADER);
        title.setForeground(COL_HEADER_TEXT);
        header.add(title, BorderLayout.NORTH);

        JLabel lore = new JLabel("<html><i>" + profession.getLore() + "</i></html>");
        lore.setFont(FONT_LORE);
        lore.setForeground(COL_LORE_TEXT);
        header.add(lore, BorderLayout.SOUTH);

        add(header, BorderLayout.NORTH);

        // ── Grid panel ────────────────────────────────────────────────────────
        // We use a custom connector-drawing panel, then overlay the skill boxes
        // via GridBagLayout.
        gridPanel = new ConnectorGridPanel();
        gridPanel.setLayout(new GridBagLayout());
        gridPanel.setBackground(COL_PANEL_BG);

        // Column labels
        String[] colLabels = {"Novice", "Tier I", "Tier II", "Tier III", "Tier IV", "Master"};
        for (int col = 0; col < 6; col++) {
            JLabel lbl = new JLabel(colLabels[col], SwingConstants.CENTER);
            lbl.setFont(new Font("SansSerif", Font.BOLD, 9));
            lbl.setForeground(new Color(0x445566));
            GridBagConstraints gbc = new GridBagConstraints();
            gbc.gridx    = col;
            gbc.gridy    = 0;
            gbc.insets   = new Insets(0, 2, 2, 2);
            gbc.anchor   = GridBagConstraints.CENTER;
            gridPanel.add(lbl, gbc);
        }

        // Place SkillBox panels from the profession definition
        for (SkillBox box : profession.getSkillBoxList()) {
            int col = box.getTreeColumn();
            int row = box.getTreeRow();

            SkillBoxPanel panel = new SkillBoxPanel(box, manager);
            panel.addPropertyChangeListener("skillActionMessage", this);

            // External listener is registered on the box panel to handle status bar
            manager.addSkillChangeListener(externalListener);

            boxPanels.add(panel);
            grid[col][row] = panel;

            GridBagConstraints gbc = new GridBagConstraints();
            gbc.gridx   = col;
            gbc.gridy   = row + 1; // +1 for column labels row
            gbc.insets  = new Insets(4, 4, 4, 4);
            gbc.fill    = GridBagConstraints.BOTH;
            gbc.weightx = 1.0;

            // Novice and Master span all 4 rows
            if (col == 0 || col == 5) {
                gbc.gridheight = 4;
                gbc.weighty    = 1.0;
            }

            gridPanel.add(panel, gbc);
        }

        // Wrap in a scroll pane so very wide trees still work
        JScrollPane scrollPane = new JScrollPane(gridPanel);
        scrollPane.setBorder(BorderFactory.createEmptyBorder());
        scrollPane.getViewport().setBackground(COL_PANEL_BG);
        scrollPane.setHorizontalScrollBarPolicy(JScrollPane.HORIZONTAL_SCROLLBAR_AS_NEEDED);
        scrollPane.setVerticalScrollBarPolicy(JScrollPane.VERTICAL_SCROLLBAR_AS_NEEDED);
        add(scrollPane, BorderLayout.CENTER);
    }

    // ── PropertyChangeListener – forwards skill action messages upward ────────

    @Override
    public void propertyChange(PropertyChangeEvent evt) {
        if ("skillActionMessage".equals(evt.getPropertyName())) {
            firePropertyChange("skillActionMessage", null, evt.getNewValue());
        }
    }

    public void dispose() {
        for (SkillBoxPanel p : boxPanels) p.dispose();
    }

    // ══════════════════════════════════════════════════════════════════════════
    //  Inner class: draws connector lines between columns
    // ══════════════════════════════════════════════════════════════════════════
    private class ConnectorGridPanel extends JPanel {

        @Override
        protected void paintComponent(Graphics g) {
            super.paintComponent(g);
            Graphics2D g2 = (Graphics2D) g.create();
            g2.setRenderingHint(RenderingHints.KEY_ANTIALIASING, RenderingHints.VALUE_ANTIALIAS_ON);

            // Draw horizontal connectors between adjacent non-null grid cells
            // We draw these BEHIND the components, so component opaqueness handles overlap.
            for (int col = 0; col < 5; col++) {
                for (int row = 0; row < 4; row++) {
                    SkillBoxPanel src  = grid[col][row];
                    SkillBoxPanel dest = grid[col + 1][row];
                    if (src == null || dest == null) continue;

                    // Midpoint right-edge of src → midpoint left-edge of dest
                    Rectangle srcBounds  = src.getBounds();
                    Rectangle destBounds = dest.getBounds();

                    if (srcBounds.width == 0) continue; // not yet laid out

                    int x1 = srcBounds.x + srcBounds.width;
                    int y1 = srcBounds.y + srcBounds.height / 2;
                    int x2 = destBounds.x;
                    int y2 = destBounds.y + destBounds.height / 2;

                    boolean lit = src.getClientProperty("learned") != null
                            || (grid[col][row] != null && manager.getSkillBox(
                               profession.getSkillBoxList()
                                         .stream()
                                         .filter(b -> b.getTreeColumn() == col && b.getTreeRow() == row)
                                         .findFirst().map(b -> b.getId()).orElse(""))
                               .isLearned());

                    g2.setStroke(new BasicStroke(lit ? 2f : 1f, BasicStroke.CAP_ROUND, BasicStroke.JOIN_ROUND));
                    g2.setColor(lit ? COL_CONNECTOR_ON : COL_CONNECTOR);
                    g2.drawLine(x1, y1, x2, y2);

                    // Arrow tip
                    int ax = x2 - 5;
                    g2.fillPolygon(new int[]{x2, ax, ax}, new int[]{y2, y2 - 4, y2 + 4}, 3);
                }
            }
            g2.dispose();
        }
    }
}
