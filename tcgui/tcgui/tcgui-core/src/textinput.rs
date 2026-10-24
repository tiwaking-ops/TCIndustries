//! TextInput widget implementation with cursor, selection, copy/paste,
//! password mask, command history hook, and scroll clamping.

use crate::geometry::Rect;
use crate::widget::WidgetState;

/// Input action resulting from key or mouse events.
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TextInputAction {
    None,
    Submit,
    Changed,
    HistoryPrev,
    HistoryNext,
}

/// A text input widget supporting character insertion, deletion, caret positioning,
/// text selection, password masking, command history navigation, clipboard operations,
/// and horizontal scroll clamping.
#[derive(Debug, Clone)]
pub struct TextInput {
    /// The backing text string.
    pub text: String,
    /// Bounding rectangle in logical pixels.
    pub bounds: Rect,
    /// Current widget state (Normal, Hover, Pressed, Disabled).
    pub state: WidgetState,
    /// Current caret index (in characters, not bytes).
    pub caret: usize,
    /// Selection anchor index (in characters). If `anchor == caret`, there is no selection.
    pub selection_anchor: usize,
    /// Whether the text is masked (password mode, rendered/treated as `*`).
    pub masked: bool,
    /// Mask character for password mode (default `*`).
    pub mask_char: char,
    /// Command history entries for Up/Down arrow navigation.
    pub history: Vec<String>,
    /// Current navigation position in `history`, or `None` if editing new input.
    pub history_idx: Option<usize>,
    /// Temporary staging buffer while navigating history.
    pub history_staging: String,
    /// Horizontal scroll offset (in logical pixels) for overflowing text.
    pub scroll_offset: f32,
    /// Character width estimate (in pixels) used for scroll clamping and cursor placement.
    pub char_width: f32,
}

impl TextInput {
    pub fn new(bounds: Rect) -> Self {
        Self {
            text: String::new(),
            bounds,
            state: WidgetState::Normal,
            caret: 0,
            selection_anchor: 0,
            masked: false,
            mask_char: '*',
            history: Vec::new(),
            history_idx: None,
            history_staging: String::new(),
            scroll_offset: 0.0,
            char_width: 8.0, // standard approx font char width
        }
    }

    /// Returns the text as displayed (masked if `masked` is true).
    pub fn displayed_text(&self) -> String {
        if self.masked {
            self.mask_char.to_string().repeat(self.text.chars().count())
        } else {
            self.text.clone()
        }
    }

    /// Number of characters in the backing text.
    pub fn len_chars(&self) -> usize {
        self.text.chars().count()
    }

    /// Clears any current selection (sets anchor to caret).
    pub fn clear_selection(&mut self) {
        self.selection_anchor = self.caret;
    }

    /// Returns `(start, end)` character indices of the current selection.
    pub fn selection_range(&self) -> (usize, usize) {
        let min = self.caret.min(self.selection_anchor);
        let max = self.caret.max(self.selection_anchor);
        (min, max)
    }

    /// Returns true if there is an active text selection.
    pub fn has_selection(&self) -> bool {
        self.caret != self.selection_anchor
    }

    /// Deletes the currently selected text range, returning true if anything was deleted.
    pub fn delete_selection(&mut self) -> bool {
        if !self.has_selection() {
            return false;
        }
        let (start, end) = self.selection_range();
        let mut chars: Vec<char> = self.text.chars().collect();
        if start < chars.len() {
            let count = (end - start).min(chars.len() - start);
            chars.drain(start..start + count);
            self.text = chars.into_iter().collect();
            self.caret = start;
            self.selection_anchor = start;
            self.clamp_caret();
            true
        } else {
            false
        }
    }

    /// Inserts a character at the current caret position (deleting selection first).
    pub fn insert_char(&mut self, c: char) {
        self.delete_selection();
        let mut chars: Vec<char> = self.text.chars().collect();
        let idx = self.caret.min(chars.len());
        chars.insert(idx, c);
        self.text = chars.into_iter().collect();
        self.caret = idx + 1;
        self.selection_anchor = self.caret;
        self.clamp_caret();
        self.clamp_scroll();
    }

    /// Inserts a string at the current caret position (paste/type).
    pub fn insert_str(&mut self, s: &str) {
        for c in s.chars() {
            if c != '\n' && c != '\r' {
                self.insert_char(c);
            }
        }
    }

    /// Backspace: deletes char before caret, or deletes selection.
    pub fn backspace(&mut self) {
        if self.delete_selection() {
            return;
        }
        if self.caret == 0 {
            return;
        }
        let mut chars: Vec<char> = self.text.chars().collect();
        let idx = self.caret - 1;
        if idx < chars.len() {
            chars.remove(idx);
            self.text = chars.into_iter().collect();
            self.caret = idx;
            self.selection_anchor = self.caret;
            self.clamp_caret();
            self.clamp_scroll();
        }
    }

    /// Delete forward: deletes char at caret, or deletes selection.
    pub fn delete_forward(&mut self) {
        if self.delete_selection() {
            return;
        }
        let mut chars: Vec<char> = self.text.chars().collect();
        if self.caret < chars.len() {
            chars.remove(self.caret);
            self.text = chars.into_iter().collect();
            self.clamp_caret();
            self.clamp_scroll();
        }
    }

    /// Move caret left, optionally expanding selection (Shift).
    pub fn move_left(&mut self, select: bool) {
        if self.caret > 0 {
            self.caret -= 1;
        }
        if !select {
            self.selection_anchor = self.caret;
        }
        self.clamp_scroll();
    }

    /// Move caret right, optionally expanding selection (Shift).
    pub fn move_right(&mut self, select: bool) {
        let len = self.len_chars();
        if self.caret < len {
            self.caret += 1;
        }
        if !select {
            self.selection_anchor = self.caret;
        }
        self.clamp_scroll();
    }

    /// Move caret to start of text (Home).
    pub fn move_home(&mut self, select: bool) {
        self.caret = 0;
        if !select {
            self.selection_anchor = self.caret;
        }
        self.scroll_offset = 0.0;
    }

    /// Move caret to end of text (End).
    pub fn move_end(&mut self, select: bool) {
        self.caret = self.len_chars();
        if !select {
            self.selection_anchor = self.caret;
        }
        self.clamp_scroll();
    }

    /// Clamps caret index within 0..=len_chars().
    pub fn clamp_caret(&mut self) {
        let len = self.len_chars();
        if self.caret > len {
            self.caret = len;
        }
        if self.selection_anchor > len {
            self.selection_anchor = len;
        }
    }

    /// Scroll clamping: adjusts `scroll_offset` so the caret is always visible within bounds.
    pub fn clamp_scroll(&mut self) {
        let visible_width = (self.bounds.width() - 16.0).max(20.0);
        let caret_x = self.caret as f32 * self.char_width;

        if caret_x < self.scroll_offset {
            self.scroll_offset = caret_x;
        } else if caret_x > self.scroll_offset + visible_width {
            self.scroll_offset = caret_x - visible_width;
        }
        if self.scroll_offset < 0.0 {
            self.scroll_offset = 0.0;
        }
    }

    /// Copy selected text (returns empty string if no selection).
    pub fn copy_selection(&self) -> String {
        if !self.has_selection() {
            return String::new();
        }
        let (start, end) = self.selection_range();
        let chars: Vec<char> = self.text.chars().collect();
        if start < chars.len() {
            let count = (end - start).min(chars.len() - start);
            chars[start..start + count].iter().collect()
        } else {
            String::new()
        }
    }

    /// Cut selected text (copies to return and deletes from text input).
    pub fn cut_selection(&mut self) -> String {
        let copied = self.copy_selection();
        self.delete_selection();
        copied
    }

    /// Command history navigation: previous history entry (Up arrow).
    pub fn history_prev(&mut self) -> bool {
        if self.history.is_empty() {
            return false;
        }
        let new_idx = match self.history_idx {
            None => {
                self.history_staging = self.text.clone();
                self.history.len() - 1
            }
            Some(idx) => {
                if idx > 0 {
                    idx - 1
                } else {
                    0
                }
            }
        };
        self.history_idx = Some(new_idx);
        self.text = self.history[new_idx].clone();
        self.move_end(false);
        true
    }

    /// Command history navigation: next history entry (Down arrow).
    pub fn history_next(&mut self) -> bool {
        let Some(idx) = self.history_idx else {
            return false;
        };
        if idx + 1 < self.history.len() {
            let new_idx = idx + 1;
            self.history_idx = Some(new_idx);
            self.text = self.history[new_idx].clone();
            self.move_end(false);
        } else {
            self.history_idx = None;
            self.text = self.history_staging.clone();
            self.move_end(false);
        }
        true
    }

    /// Push current text to command history and reset staging.
    pub fn submit_to_history(&mut self) {
        if !self.text.is_empty() && self.history.last() != Some(&self.text) {
            self.history.push(self.text.clone());
        }
        self.history_idx = None;
        self.history_staging.clear();
    }

    pub fn bounds(&self) -> Rect {
        self.bounds
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::geometry::Vec2;

    #[test]
    fn test_textinput_caret_and_typing() {
        let mut ti = TextInput::new(Rect::from_origin_size(Vec2::ZERO, Vec2::new(200.0, 24.0)));
        ti.insert_char('H');
        ti.insert_char('i');
        assert_eq!(ti.text, "Hi");
        assert_eq!(ti.caret, 2);

        ti.move_left(false);
        assert_eq!(ti.caret, 1);
        ti.insert_char('a');
        assert_eq!(ti.text, "Hai");
        assert_eq!(ti.caret, 2);
    }

    #[test]
    fn test_textinput_selection_and_copy_cut() {
        let mut ti = TextInput::new(Rect::from_origin_size(Vec2::ZERO, Vec2::new(200.0, 24.0)));
        ti.insert_str("Hello World");
        // Select "World" (indices 6 to 11)
        ti.caret = 11;
        ti.selection_anchor = 6;
        assert!(ti.has_selection());
        assert_eq!(ti.copy_selection(), "World");

        let cut = ti.cut_selection();
        assert_eq!(cut, "World");
        assert_eq!(ti.text, "Hello ");
        assert!(!ti.has_selection());
    }

    #[test]
    fn test_textinput_password_mask() {
        let mut ti = TextInput::new(Rect::from_origin_size(Vec2::ZERO, Vec2::new(200.0, 24.0)));
        ti.masked = true;
        ti.insert_str("secret");
        assert_eq!(ti.text, "secret");
        assert_eq!(ti.displayed_text(), "******");
    }

    #[test]
    fn test_textinput_command_history() {
        let mut ti = TextInput::new(Rect::from_origin_size(Vec2::ZERO, Vec2::new(200.0, 24.0)));
        ti.insert_str("command 1");
        ti.submit_to_history();
        ti.text.clear();
        ti.insert_str("command 2");
        ti.submit_to_history();
        ti.text.clear();

        // Press Up (prev) -> gets "command 2"
        assert!(ti.history_prev());
        assert_eq!(ti.text, "command 2");

        // Press Up again -> gets "command 1"
        assert!(ti.history_prev());
        assert_eq!(ti.text, "command 1");

        // Press Down (next) -> gets "command 2"
        assert!(ti.history_next());
        assert_eq!(ti.text, "command 2");

        // Press Down again -> returns to staging (empty)
        assert!(ti.history_next());
        assert_eq!(ti.text, "");
    }

    #[test]
    fn test_textinput_scroll_clamping() {
        let mut ti = TextInput::new(Rect::from_origin_size(Vec2::ZERO, Vec2::new(100.0, 24.0)));
        ti.char_width = 10.0;
        // Insert 20 chars (200px width, visible width ~84px)
        for _ in 0..20 {
            ti.insert_char('a');
        }
        // Caret is at 20, caret_x = 200.0. Scroll offset should clamp.
        assert!(ti.scroll_offset > 0.0);
    }
}
