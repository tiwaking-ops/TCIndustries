package swg.ui;

import swg.model.Character;

import javax.swing.*;
import java.awt.*;

/**
 * Custom Swing component that renders the 250-point skill pool bar.
 *
 * Visual design: a segmented gauge inspired by pre-CU SWG's character
 * sheet, with amber fill that turns red as the pool nears exhaustion.
 */
public class PointPoolBar extends JComponent {

    private static final Color COL_FILL_SAFE      = new Color(0xC8860A);   // Amber
    private static final Color COL_FILL_WARN       = new Color(0xCC4400);  // Orange-red
    private static final Color COL_FILL_CRITICAL   = new Color(0xAA0000);  // Deep red
    private static final Color COL_BG              = new Color(0x0E1318);
    private static final Color COL_TRACK           = new Color(0x1A2530);
    private static final Color COL_BORDER          = new Color(0x2A3D50);
    private static final Color COL_SEGMENT_LINE    = new Color(0x0E1318);
    private static final Color COL_TEXT_PRIMARY    = new Color(0xEEDEBB);
    private static final Color COL_TEXT_SECONDARY  = new Color(0x7A8A9A);

    private static final Font FONT_VALUE = new Font("Serif", Font.BOLD, 18);
    private static final Font FONT_LABEL = new Font("SansSerif", Font.PLAIN, 11);
    private static final Font FONT_CAP   = new Font("SansSerif", Font.BOLD, 11);

    private static final int SEGMENT_COUNT = 25; // Each segment = 10 pts

    private final Character character;

    public PointPoolBar(Character character) {
        this.character = character;
        setPreferredSize(new Dimension(600, 64));
        setMinimumSize(new Dimension(300, 48));
    }

    @Override
    protected void paintComponent(Graphics g) {
        Graphics2D g2 = (Graphics2D) g.create();
        g2.setRenderingHint(RenderingHints.KEY_ANTIALIASING, RenderingHints.VALUE_ANTIALIAS_ON);
        g2.setRenderingHint(RenderingHints.KEY_TEXT_ANTIALIASING, RenderingHints.VALUE_TEXT_ANTIALIAS_ON);

        int w  = getWidth();
        int h  = getHeight();
        int remaining  = character.getRemainingPoints();
        int spent      = character.getSpentPoints();
        int max        = Character.MAX_SKILL_POINTS;
        double pctFill = (double) remaining / max;  // Remaining = filled

        // ── Background ────────────────────────────────────────────────────────
        g2.setColor(COL_BG);
        g2.fillRect(0, 0, w, h);

        // ── Labels section (left side) ────────────────────────────────────────
        int labelWidth = 140;
        g2.setFont(FONT_LABEL);
        g2.setColor(COL_TEXT_SECONDARY);
        g2.drawString("SKILL POINTS", 12, 20);

        g2.setFont(FONT_VALUE);
        g2.setColor(remaining <= 20 ? COL_FILL_CRITICAL : COL_TEXT_PRIMARY);
        g2.drawString(String.valueOf(remaining), 12, 44);

        g2.setFont(FONT_LABEL);
        g2.setColor(COL_TEXT_SECONDARY);
        g2.drawString("/ " + max + " remaining", 12 + g2.getFontMetrics(FONT_VALUE).stringWidth(String.valueOf(remaining)) + 4, 44);

        // Spent indicator
        g2.setFont(FONT_CAP);
        g2.setColor(new Color(0x88661A));
        g2.drawString("SPENT: " + spent, labelWidth - 10, 58);

        // ── Track ─────────────────────────────────────────────────────────────
        int trackX = labelWidth;
        int trackY = 12;
        int trackW = w - labelWidth - 16;
        int trackH = h - 24;

        // Track background
        g2.setColor(COL_TRACK);
        g2.fillRoundRect(trackX, trackY, trackW, trackH, 6, 6);
        g2.setColor(COL_BORDER);
        g2.drawRoundRect(trackX, trackY, trackW, trackH, 6, 6);

        // Fill (remaining points = filled portion, from left)
        int fillW = (int) (trackW * pctFill);
        if (fillW > 2) {
            Color fillColor;
            if (pctFill > 0.4)       fillColor = COL_FILL_SAFE;
            else if (pctFill > 0.15) fillColor = COL_FILL_WARN;
            else                     fillColor = COL_FILL_CRITICAL;

            // Gradient fill
            GradientPaint grad = new GradientPaint(
                    trackX, trackY, fillColor.brighter(),
                    trackX, trackY + trackH, fillColor.darker());
            g2.setPaint(grad);
            g2.fillRoundRect(trackX + 2, trackY + 2, fillW - 2, trackH - 4, 4, 4);
        }

        // Segment tick marks (one per 10 pts)
        g2.setColor(COL_SEGMENT_LINE);
        g2.setStroke(new BasicStroke(1.0f));
        double segW = (double) trackW / SEGMENT_COUNT;
        for (int i = 1; i < SEGMENT_COUNT; i++) {
            int sx = trackX + (int)(i * segW);
            g2.drawLine(sx, trackY + 4, sx, trackY + trackH - 4);
        }

        // Spent region overlay (right portion of the bar)
        if (spent > 0) {
            int spentX = trackX + fillW;
            int spentW = trackW - fillW;
            if (spentW > 0) {
                g2.setColor(new Color(0x1A0A0A, true)); // semi-transparent dark
                g2.fillRoundRect(spentX, trackY + 2, spentW - 2, trackH - 4, 4, 4);
            }
        }

        // Point labels at 0, 50, 100, 150, 200, 250
        g2.setFont(new Font("SansSerif", Font.PLAIN, 9));
        g2.setColor(new Color(0x445566));
        int[] milestones = {0, 50, 100, 150, 200, 250};
        for (int m : milestones) {
            int mx = trackX + (int)((double) m / max * trackW);
            g2.drawString(String.valueOf(m), mx - 4, trackY + trackH + 10);
        }

        g2.dispose();
    }

    /** Call this when the character's point pool changes to trigger a repaint. */
    public void refresh() {
        repaint();
    }
}
