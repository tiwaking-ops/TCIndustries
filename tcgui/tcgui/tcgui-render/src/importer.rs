//! Optional SWG asset import surface — texture decoding (task T-81).
//!
//! Gated behind the `swg-import` Cargo feature. Decodes DDS and TGA images into
//! premultiplied RGBA, matching `tcgui_core::Color`'s premultiplied convention.
//!
//! Per ADR-0002 and `docs/importer/findings.md` (T-80), no extracted SWG assets
//! exist in this repository. This module therefore operates only on
//! locally generated fixtures and does not encode any SWG format assumptions
//! beyond the container formats themselves, which are industry-standard.

use tcgui_core::Color;

/// Errors produced while importing an image.
#[derive(Debug, thiserror::Error)]
pub enum ImportError {
    /// The underlying image bytes could not be decoded.
    #[error("failed to decode image: {0}")]
    Decode(String),
    /// Decoded image has zero pixels in one or both dimensions.
    #[error("image has zero extent: {width}x{height}")]
    EmptyImage { width: u32, height: u32 },
}

/// A decoded image in premultiplied RGBA8 form.
#[derive(Debug, Clone, PartialEq)]
pub struct ImportedTexture {
    /// Width in pixels.
    pub width: u32,
    /// Height in pixels.
    pub height: u32,
    /// Row-major premultiplied RGBA8 data, length = `width * height * 4`.
    pub data: Vec<u8>,
}

impl ImportedTexture {
    /// Number of pixels in the image.
    pub fn pixel_count(&self) -> usize {
        (self.width as usize) * (self.height as usize)
    }

    /// Fetch pixel `(x, y)` as a premultiplied [`Color`].
    pub fn pixel_color(&self, x: u32, y: u32) -> Option<Color> {
        if x >= self.width || y >= self.height {
            return None;
        }
        let i = ((y * self.width + x) as usize) * 4;
        // Channel values were premultiplied and rounded to u8 on import, so
        // converting back through /255.0 recovers the premultiplied float.
        Some(Color::premul(
            f32::from(self.data[i]) / 255.0,
            f32::from(self.data[i + 1]) / 255.0,
            f32::from(self.data[i + 2]) / 255.0,
            f32::from(self.data[i + 3]) / 255.0,
        ))
    }
}

/// Decode image bytes (DDS or TGA) into premultiplied RGBA8.
///
/// Accepts any format `image` recognizes given the enabled `dds`/`tga`
/// features. Non-RGBA inputs are converted to RGBA8; the resulting channel
/// values are then premultiplied against alpha.
///
/// Detection prefers `load_from_memory` and falls back to the explicit TGA
/// decoder. Some valid TGA files (notably minimal headers without a footer, and
/// RLE variants) are not recognised by `image`'s format sniffing even when the
/// `tga` codec is enabled, so the fallback keeps the common case working
/// without weakening the strictness of the primary path.
pub fn decode_image(bytes: &[u8]) -> Result<ImportedTexture, ImportError> {
    let rgba = match image::load_from_memory(bytes) {
        Ok(img) => img.to_rgba8(),
        Err(sniff_err) => decode_as_tga(bytes).ok_or_else(|| ImportError::Decode(sniff_err.to_string()))?,
    };

    let width = rgba.width();
    let height = rgba.height();
    if width == 0 || height == 0 {
        return Err(ImportError::EmptyImage { width, height });
    }

    let mut data = rgba.into_raw();
    premultiply_in_place(&mut data);

    Ok(ImportedTexture {
        width,
        height,
        data,
    })
}

/// Attempt to decode `bytes` using the TGA codec directly.
///
/// Returns `None` if the bytes are not decodable TGA rather than surfacing the
/// codec's own error text, letting callers report the original sniff failure.
fn decode_as_tga(bytes: &[u8]) -> Option<image::RgbaImage> {
    use image::ImageDecoder;

    let decoder = image::codecs::tga::TgaDecoder::new(bytes).ok()?;
    let (width, height) = decoder.dimensions();
    let mut buf = vec![0u8; decoder.total_bytes() as usize];
    decoder.read_image(&mut buf).ok()?;
    image::RgbaImage::from_raw(width, height, buf)
}

/// Premultiply straight RGBA8 bytes in place: `rgb *= a / 255`.
///
/// Rounds to nearest to keep the round trip through u8 as faithful as possible.
fn premultiply_in_place(data: &mut [u8]) {
    let (chunks, _) = data.as_chunks_mut::<4>();
    for px in chunks {
        let a = u32::from(px[3]);
        if a == 255 {
            continue; // fully opaque: premultiplying is the identity
        }
        for c in px.iter_mut().take(3) {
            let v = u32::from(*c) * a; // max 255*255 = 65025
            // round(v / 255) using integer math only.
            let premul = ((v + 128) + ((v + 128) >> 8)) >> 8;
            *c = premul.min(255) as u8;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    /// Build an **uncompressed** 32-bit TGA (image type 0x02) around `pixels`.
    ///
    /// Hand-rolling the fixture avoids `TgaEncoder`, which emits RLE-compressed
    /// TGA (type 0x0A) whose header `image` cannot always auto-detect. Writing
    /// the header explicitly keeps the fixtures deterministic and makes it
    /// obvious they are locally generated — T-81's acceptance requirement.
    fn tga_rgba_uncompressed(width: u16, height: u16, pixels: &[u8]) -> Vec<u8> {
        assert_eq!(
            pixels.len(),
            (width as usize) * (height as usize) * 4,
            "fixture pixel data length must match extent"
        );
        // TGA stores rows bottom-up and channels as BGRA.
        let mut out = Vec::with_capacity(18 + pixels.len());

        out.extend_from_slice(&[
            0x00, // id length: no ID field
            0x00, // colour map type: no colour map
            0x02, // image type: uncompressed true-colour
            0x00, 0x00, // colour map specification: unused
            0x00, 0x00,
            0x00, // colour map entry size: unused
            0x00, 0x00, // x origin
            0x00, 0x00, // y origin
        ]);
        out.extend_from_slice(&width.to_le_bytes());
        out.extend_from_slice(&height.to_le_bytes());
        out.extend_from_slice(&[
            0x20, // pixel depth: 32 bits
            0x28, // image descriptor: 32bpp, 8-bit alpha, origin bottom-left
        ]);

        let row_bytes = (width as usize) * 4;
        for y in (0..height as usize).rev() {
            let row = &pixels[y * row_bytes..(y + 1) * row_bytes];
        let (px_chunks, _) = row.as_chunks::<4>();
        for px in px_chunks {
            let [r, g, b, a] = [px[0], px[1], px[2], px[3]];
            out.extend_from_slice(&[b, g, r, a]);
        }
        }
        out
    }

    fn encode_tga_rgba(width: u32, height: u32, pixels: &[u8]) -> Vec<u8> {
        tga_rgba_uncompressed(width as u16, height as u16, pixels)
    }


    #[test]
    fn test_decode_tga_opaque_pixels_round_trip() {
        // 2x1: one opaque red, one opaque blue. Alpha 255 => premultiply identity.
        let pixels = [255u8, 0, 0, 255, 0, 0, 255, 255];
        let bytes = encode_tga_rgba(2, 1, &pixels);
        let tex = decode_image(&bytes).unwrap();

        assert_eq!(tex.width, 2);
        assert_eq!(tex.height, 1);
        assert_eq!(tex.pixel_count(), 2);
        assert_eq!(tex.data.len(), 8);

        let a = tex.pixel_color(0, 0).unwrap();
        assert!((a.r - 1.0).abs() < 0.01, "red channel: {}", a.r);
        assert!((a.g - 0.0).abs() < 0.01);
        assert!((a.a - 1.0).abs() < 0.01, "alpha: {}", a.a);

        let b = tex.pixel_color(1, 0).unwrap();
        assert!((b.b - 1.0).abs() < 0.01, "blue channel: {}", b.b);
    }

    #[test]
    fn test_decode_premultiplies_partial_alpha() {
        // Fully saturated colour at 50% alpha must lose about half its energy.
        let pixels = [255u8, 255, 255, 128];
        let bytes = encode_tga_rgba(1, 1, &pixels);
        let tex = decode_image(&bytes).unwrap();

        let c = tex.pixel_color(0, 0).unwrap();
        assert!((c.a - 128.0 / 255.0).abs() < 0.01, "alpha {}", c.a);
        for ch in [c.r, c.g, c.b] {
            assert!(
                (ch - 128.0 / 255.0).abs() < 0.02,
                "expected ~0.502 premultiplied, got {ch}"
            );
        }
    }

    #[test]
    fn test_decode_zero_alpha_fully_transparent() {
        // Any colour at zero alpha collapses to zero — premultiplied black.
        let pixels = [255u8, 255, 255, 0];
        let bytes = encode_tga_rgba(1, 1, &pixels);
        let tex = decode_image(&bytes).unwrap();

        let c = tex.pixel_color(0, 0).unwrap();
        assert_eq!(c.a, 0.0);
        assert!(c.r.abs() < 1e-6 && c.g.abs() < 1e-6 && c.b.abs() < 1e-6);
    }

    #[test]
    fn test_decode_rejects_garbage() {
        let err = decode_image(&[0u8, 1, 2, 3, 4, 5, 6, 7]).unwrap_err();
        assert!(matches!(err, ImportError::Decode(_)), "got {err:?}");
    }

    #[test]
    fn test_pixel_color_out_of_bounds_is_none() {
        let pixels = [10u8, 20, 30, 255];
        let bytes = encode_tga_rgba(1, 1, &pixels);
        let tex = decode_image(&bytes).unwrap();
        assert!(tex.pixel_color(1, 0).is_none());
        assert!(tex.pixel_color(0, 1).is_none());
        assert!(tex.pixel_color(0, 0).is_some());
    }

    #[test]
    fn test_premultiply_matches_float_reference() {
        // Cross-check the integer fast path against straightforward float math.
        let mut cases = Vec::new();
        for a in (0u8..=255).step_by(17) {
            for c in (0u8..=255).step_by(17) {
                cases.extend_from_slice(&[c, c, c, a]);
            }
        }
        let mut got = cases.clone();
        premultiply_in_place(&mut got);

        let (got_chunks, _) = got.as_chunks::<4>();
        for (i, px) in got_chunks.iter().enumerate() {
            let src = &cases[i * 4..i * 4 + 4];
            let alpha = f32::from(src[3]) / 255.0;
            let expected = f32::from(src[0]) / 255.0 * alpha;
            let actual = f32::from(px[0]) / 255.0;
            assert!(
                (expected - actual).abs() <= 0.004,
                "channel mismatch at alpha {}: expected {expected}, got {actual}",
                src[3]
            );
        }
    }
}
