//! Font rasterization via `ab_glyph`, LRU glyph cache, text measurement, wrapping, and color runs.

use ab_glyph::{point, Font, FontArc, GlyphId, OutlinedGlyph};
use std::collections::{HashMap, VecDeque};
use tcgui_core::geometry::{Color, Rect, Vec2};

/// Key identifying a rasterized glyph in the cache.
#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub struct GlyphKey {
    pub glyph_id: u16,
    pub font_size_scaled: u32,
}

/// Cached metrics and atlas coordinates for a rasterized glyph.
#[derive(Debug, Clone, Copy)]
pub struct CachedGlyph {
    pub uv_rect: Rect,
    pub width: u32,
    pub height: u32,
    pub bearing_x: f32,
    pub bearing_y: f32,
}

/// Fixed-capacity LRU glyph cache ensuring memory and cache size never grow unbounded.
pub struct LruGlyphCache {
    capacity: usize,
    cache: HashMap<GlyphKey, CachedGlyph>,
    lru_list: VecDeque<GlyphKey>,
}

impl LruGlyphCache {
    /// Creates a new LRU glyph cache with the specified maximum capacity.
    pub fn new(capacity: usize) -> Self {
        Self {
            capacity,
            cache: HashMap::with_capacity(capacity),
            lru_list: VecDeque::new(),
        }
    }

    /// Returns the number of cached glyphs.
    pub fn len(&self) -> usize {
        self.cache.len()
    }

    /// Returns true if the cache is empty.
    pub fn is_empty(&self) -> bool {
        self.cache.is_empty()
    }

    /// Returns the configured maximum capacity.
    pub fn capacity(&self) -> usize {
        self.capacity
    }

    /// Retrieves a cached glyph by key, updating its recency.
    pub fn get(&mut self, key: &GlyphKey) -> Option<&CachedGlyph> {
        if self.cache.contains_key(key) {
            self.lru_list.retain(|k| k != key);
            self.lru_list.push_back(*key);
            self.cache.get(key)
        } else {
            None
        }
    }

    /// Inserts a glyph into the cache, evicting the least recently used entry if at capacity.
    pub fn insert(&mut self, key: GlyphKey, glyph: CachedGlyph) {
        if self.cache.contains_key(&key) {
            self.lru_list.retain(|k| k != &key);
        } else if self.cache.len() >= self.capacity {
            if let Some(oldest) = self.lru_list.pop_front() {
                self.cache.remove(&oldest);
            }
        }
        self.cache.insert(key, glyph);
        self.lru_list.push_back(key);
    }
}

/// Font atlas manager wrapping `ab_glyph::FontArc` and LRU glyph caching.
pub struct FontAtlas {
    font: FontArc,
    pub cache: LruGlyphCache,
}

impl FontAtlas {
    /// Loads a font from binary data (e.g. TTF/OTF bytes).
    pub fn try_from_bytes(bytes: &[u8], cache_capacity: usize) -> Result<Self, ab_glyph::InvalidFont> {
        let font = FontArc::try_from_vec(bytes.to_vec())?;
        Ok(Self {
            font,
            cache: LruGlyphCache::new(cache_capacity),
        })
    }

    /// Gets glyph ID for a character.
    pub fn glyph_id(&self, c: char) -> Option<GlyphId> {
        Some(self.font.glyph_id(c))
    }

    /// Rasterizes a glyph outline using ab_glyph if available.
    pub fn rasterize_glyph(&self, glyph_id: GlyphId, font_size: f32) -> Option<OutlinedGlyph> {
        let glyph = glyph_id.with_scale_and_position(font_size, point(0.0, 0.0));
        self.font.outline_glyph(glyph)
    }
}

/// A styled text span with font size and color (color run).
#[derive(Debug, Clone)]
pub struct TextSpan {
    pub text: String,
    pub color: Color,
    pub font_size: f32,
}

/// Measures the total advance width of a sequence of styled text runs.
pub fn measure_color_runs(spans: &[TextSpan]) -> Vec2 {
    let mut width = 0.0;
    let mut max_height = 0.0;
    for span in spans {
        let line = TextLayoutEngine::measure_text(&span.text, span.font_size);
        width += line.x;
        if line.y > max_height {
            max_height = line.y;
        }
    }
    Vec2::new(width, max_height)
}

/// Text layout engine supporting measurement, word wrapping, and color runs.
pub struct TextLayoutEngine;

impl TextLayoutEngine {
    /// Measures the bounding size of a text string at a given font size.
    pub fn measure_text(text: &str, font_size: f32) -> Vec2 {
        let char_w = font_size * 0.6;
        let lines = text.lines().count().max(1);
        let max_line_chars = text
            .lines()
            .map(|l| l.chars().count())
            .max()
            .unwrap_or(text.chars().count());
        let width = max_line_chars as f32 * char_w;
        let height = lines as f32 * (font_size * 1.2);
        Vec2::new(width, height)
    }

    /// Wraps text into lines that fit within `max_width`.
    pub fn wrap_text(text: &str, max_width: f32, font_size: f32) -> Vec<String> {
        let mut result_lines = Vec::new();
        let char_w = font_size * 0.6;
        let max_chars = (max_width / char_w).max(1.0) as usize;

        for paragraph in text.lines() {
            if paragraph.is_empty() {
                result_lines.push(String::new());
                continue;
            }
            let words: Vec<&str> = paragraph.split_whitespace().collect();
            if words.is_empty() {
                result_lines.push(paragraph.to_string());
                continue;
            }

            let mut current_line = String::new();
            for word in words {
                if current_line.is_empty() {
                    current_line = word.to_string();
                } else if current_line.len() + 1 + word.len() <= max_chars {
                    current_line.push(' ');
                    current_line.push_str(word);
                } else {
                    result_lines.push(current_line);
                    current_line = word.to_string();
                }
            }
            if !current_line.is_empty() {
                result_lines.push(current_line);
            }
        }

        if result_lines.is_empty() {
            result_lines.push(text.to_string());
        }
        result_lines
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_lru_glyph_cache_capacity_limit() {
        let capacity = 100;
        let mut cache = LruGlyphCache::new(capacity);

        // Stress test with 1,000,000 insertions
        for i in 0..1_000_000 {
            let key = GlyphKey {
                glyph_id: (i % 5000) as u16,
                font_size_scaled: (i % 3) as u32,
            };
            let glyph = CachedGlyph {
                uv_rect: Rect::zero(),
                width: 10,
                height: 10,
                bearing_x: 0.0,
                bearing_y: 0.0,
            };
            cache.insert(key, glyph);
            assert!(
                cache.len() <= capacity,
                "Cache size {} exceeded capacity {}",
                cache.len(),
                capacity
            );
        }

        assert_eq!(cache.len(), capacity);
    }

    #[test]
    fn test_text_measurement() {
        let size = TextLayoutEngine::measure_text("Hello", 12.0);
        assert!(size.x > 0.0);
        assert!(size.y > 0.0);
    }

    #[test]
    fn test_text_wrapping() {
        let text = "The quick brown fox jumps over the lazy dog";
        let lines = TextLayoutEngine::wrap_text(text, 100.0, 12.0);
        assert!(!lines.is_empty());
        for line in &lines {
            let sz = TextLayoutEngine::measure_text(line, 12.0);
            assert!(sz.x <= 150.0);
        }
    }
}
