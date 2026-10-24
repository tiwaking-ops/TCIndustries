#![allow(missing_docs)]

pub mod font;
pub mod vertex;

#[cfg(feature = "swg-import")]
pub mod importer;

pub use font::{
    CachedGlyph, GlyphKey, LruGlyphCache, TextLayoutEngine, TextSpan,
};
pub use vertex::UiVertex;

#[cfg(feature = "swg-import")]
pub use importer::{ImportedTexture, ImportError, decode_image};

pub fn render_placeholder() {}
