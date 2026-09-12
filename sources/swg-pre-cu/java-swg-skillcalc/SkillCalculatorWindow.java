package swg.ui;

import swg.manager.SkillManager;
import swg.model.Character;
import swg.model.Profession;
import swg.model.SkillBox;

import javax.swing.*;
import javax.swing.border.EmptyBorder;
import javax.swing.border.MatteBorder;
import java.awt.*;
import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.util.ArrayList;
import java.util.List;

/**
 * The main application window for the SWG Pre-CU Skill Calculator.
 *
 * Layout:
 *  ┌──────────────────────────────────────────────────────────┐
 *  │  TITLE BAR                                               │
 *  ├──────────────────────────────────────────────────────────┤
 *  │  POINT POOL BAR   (always visible at top)                │
 *  ├──────────────────────────────────────────────────────────┤
 *  │  TABBED PANE — one tab per Profession tree               │
 *  │    [Marksman] [Brawler] [Artisan]                        │
 *  │    ┌────────────────────────────────────────────────┐    │
 *  │    │  ProfessionTreePanel (grid of SkillBoxPanels) │    │
 *  │    └────────────────────────────────────────────────┘    │
 *  ├──────────────────────────────────────────────────────────┤
 *  │  LEGEND / KEY                                            │
 *  ├──────────────────────────────────────────────────────────┤
 *  │  STATUS BAR   (last action message)                      │
 *  └──────────────────────────────────────────────────────────┘
 */
public class SkillCalculatorWindow extends JFrame
        implements SkillManager.SkillChangeListener, PropertyChangeListener {

    // ── Palette ───────────────────────────────────────────────────────────────
    private static final Color COL_WINDOW_BG    = new Color(0x0A0F14);
    private static final Color COL_TITLEBAR_BG  = new Color(0x0D1920);
    private static final Color COL_STATUS_BG    = new Color(0x080D10);
    private static final Color COL_STATUS_OK    = new Color(0x66BB88);
    private static final Color COL_STATUS_ERR   = new Color(0xCC6644);
    private static final Color COL_GOLD         = new Color(0xC8A85A);
    private static final Color COL_TAB_BG       = new Color(0x12202C);
    private static final Color COL_DIVIDER      = new Color(0x1E3040);

    private static final Font FONT_TITLE  = new Font("Serif", Font.BOLD, 22);
    private static final Font FONT_SUB    = new Font("SansSerif", Font.ITALIC, 11);
    private static final Font FONT_STATUS = new Font("SansSerif", Font.PLAIN, 12);

    // ── State ─────────────────────────────────────────────────────────────────
    private final SkillManager manager;
    private final PointPoolBar poolBar;
    private final JLabel statusLabel;
    private final JTabbedPane tabbedPane;
    private final List<ProfessionTreePanel> treePanels = new ArrayList<>();

    // ── Constructor ───────────────────────────────────────────────────────────

    public SkillCalculatorWindow(SkillManager manager) {
        super("SWG Pre-CU Skill Calculator — " + manager.getCharacter().getName()
              + " (" + manager.getCharacter().getSpecies() + ")");
        this.manager = manager;

        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setMinimumSize(new Dimension(960, 680));
        setPreferredSize(new Dimension(1200, 780));

        getContentPane().setBackground(COL_WINDOW_BG);
        getContentPane().setLayout(new BorderLayout(0, 0));

        // Register this window as a listener for model changes
        manager.addSkillChangeListener(this);

        // ── Build sections in order ────────────────────────────────────────────
        add(buildTitleBar(), BorderLayout.NORTH);

        // Centre: pool bar + tabbed profession trees stacked in a container
        JPanel centrePanel = new JPanel(new BorderLayout(0, 0));
        centrePanel.setBackground(COL_WINDOW_BG);

        poolBar = new PointPoolBar(manager.getCharacter());
        poolBar.setBorder(new EmptyBorder(4, 12, 8, 12));
        poolBar.setBackground(COL_WINDOW_BG);
        centrePanel.add(poolBar, BorderLayout.NORTH);

        tabbedPane = buildTabbedPane();
        centrePanel.add(tabbedPane, BorderLayout.CENTER);

        add(centrePanel, BorderLayout.CENTER);

        JPanel bottom = new JPanel(new BorderLayout(0, 0));
        bottom.setBackground(COL_WINDOW_BG);
        bottom.add(buildLegend(), BorderLayout.NORTH);

        // Status bar
        statusLabel = new JLabel(" Ready — select a skill to begin.", SwingConstants.LEFT);
        statusLabel.setFont(FONT_STATUS);
        statusLabel.setForeground(COL_STATUS_OK);
        statusLabel.setBorder(new EmptyBorder(6, 14, 6, 14));
        statusLabel.setBackground(COL_STATUS_BG);
        statusLabel.setOpaque(true);
        JPanel statusBar = new JPanel(new BorderLayout());
        statusBar.setBackground(COL_STATUS_BG);
        statusBar.setBorder(new MatteBorder(1, 0, 0, 0, COL_DIVIDER));
        statusBar.add(statusLabel, BorderLayout.CENTER);

        // Character summary on the right of the status bar
        JLabel charSummary = new JLabel(
                manager.getCharacter().getName() + "  |  " + manager.getCharacter().getSpecies() + "   ");
        charSummary.setFont(new Font("SansSerif", Font.ITALIC, 11));
        charSummary.setForeground(new Color(0x445566));
        statusBar.add(charSummary, BorderLayout.EAST);

        bottom.add(statusBar, BorderLayout.SOUTH);
        add(bottom, BorderLayout.SOUTH);

        pack();
        setLocationRelativeTo(null);
    }

    // ── Title bar ─────────────────────────────────────────────────────────────

    private JPanel buildTitleBar() {
        JPanel bar = new JPanel(new BorderLayout(16, 0));
        bar.setBackground(COL_TITLEBAR_BG);
        bar.setBorder(new EmptyBorder(10, 18, 10, 18));

        JPanel left = new JPanel(new GridLayout(2, 1));
        left.setBackground(COL_TITLEBAR_BG);

        JLabel title = new JLabel("STAR WARS GALAXIES  ·  PRE-CU SKILL CALCULATOR");
        title.setFont(FONT_TITLE);
        title.setForeground(COL_GOLD);
        left.add(title);

        JLabel sub = new JLabel("250 Skill Point Hard Cap  ·  Profession Branching System  ·  Pre-Combat Upgrade Ruleset");
        sub.setFont(FONT_SUB);
        sub.setForeground(new Color(0x5A7A8A));
        left.add(sub);

        bar.add(left, BorderLayout.CENTER);
        bar.setBorder(BorderFactory.createCompoundBorder(
                new MatteBorder(0, 0, 2, 0, new Color(0xC8860A)),
                new EmptyBorder(10, 18, 10, 18)));
        return bar;
    }

    // ── Tabbed pane with one tab per profession ───────────────────────────────

    private JTabbedPane buildTabbedPane() {
        JTabbedPane tabs = new JTabbedPane(JTabbedPane.TOP);
        tabs.setBackground(COL_TAB_BG);
        tabs.setForeground(COL_GOLD);
        tabs.setFont(new Font("SansSerif", Font.BOLD, 12));

        // Custom UI colours for the tab bar
        UIManager.put("TabbedPane.selected",            new Color(0x1A2A38));
        UIManager.put("TabbedPane.background",          COL_TAB_BG);
        UIManager.put("TabbedPane.foreground",          COL_GOLD);
        UIManager.put("TabbedPane.tabAreaBackground",   COL_TAB_BG);
        UIManager.put("TabbedPane.contentAreaColor",    new Color(0x0E1620));
        tabs.updateUI();

        for (Profession prof : manager.getProfessionList()) {
            ProfessionTreePanel treePanel = new ProfessionTreePanel(prof, manager, this);
            treePanel.addPropertyChangeListener("skillActionMessage", this);
            treePanels.add(treePanel);
            tabs.addTab("  " + prof.getDisplayName() + "  ", treePanel);
        }

        return tabs;
    }

    // ── Legend ────────────────────────────────────────────────────────────────

    private JPanel buildLegend() {
        JPanel legend = new JPanel(new FlowLayout(FlowLayout.LEFT, 20, 6));
        legend.setBackground(new Color(0x0A1018));
        legend.setBorder(new MatteBorder(1, 0, 0, 0, COL_DIVIDER));

        legend.add(makeLegendItem(new Color(0x1A2530), new Color(0x88FFAA), "Available"));
        legend.add(makeLegendItem(new Color(0xC8860A), Color.WHITE,         "Learned"));
        legend.add(makeLegendItem(new Color(0x1A1F25), new Color(0x555E66), "Locked (prereq missing)"));
        legend.add(makeLegendItem(new Color(0x8B0000), new Color(0xFFCCCC), "Master Box"));

        JLabel hint = new JLabel("  Hover for skill details · Click Learn / Drop to modify");
        hint.setFont(new Font("SansSerif", Font.ITALIC, 10));
        hint.setForeground(new Color(0x3A4F62));
        legend.add(hint);

        return legend;
    }

    private JPanel makeLegendItem(Color bg, Color fg, String label) {
        JPanel item = new JPanel(new FlowLayout(FlowLayout.LEFT, 4, 0));
        item.setBackground(new Color(0x0A1018));

        JPanel swatch = new JPanel();
        swatch.setPreferredSize(new Dimension(14, 14));
        swatch.setBackground(bg);
        swatch.setBorder(BorderFactory.createLineBorder(fg.darker(), 1));
        item.add(swatch);

        JLabel lbl = new JLabel(label);
        lbl.setFont(new Font("SansSerif", Font.PLAIN, 10));
        lbl.setForeground(new Color(0x6A8499));
        item.add(lbl);

        return item;
    }

    // ── SkillChangeListener — refreshes the point pool bar ────────────────────

    @Override
    public void onSkillChanged(SkillBox box, Character character) {
        SwingUtilities.invokeLater(poolBar::refresh);
    }

    // ── PropertyChangeListener — receives status messages from tree panels ────

    @Override
    public void propertyChange(PropertyChangeEvent evt) {
        if ("skillActionMessage".equals(evt.getPropertyName())) {
            String msg = (String) evt.getNewValue();
            SwingUtilities.invokeLater(() -> {
                boolean isError = msg != null && (msg.startsWith("Not") || msg.startsWith("Requires")
                        || msg.startsWith("Cannot") || msg.startsWith("Unknown"));
                statusLabel.setForeground(isError ? COL_STATUS_ERR : COL_STATUS_OK);
                statusLabel.setText(" " + msg);
            });
        }
    }

    /** Clean up all listeners before closing. */
    @Override
    public void dispose() {
        for (ProfessionTreePanel tp : treePanels) tp.dispose();
        super.dispose();
    }
}
