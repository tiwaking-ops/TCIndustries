//! Core geometry types for the TCIndustries GUI system.
//!
//! All color channels use premultiplied alpha representation (r, g, b are
//! pre-scaled by the alpha component). This avoids repeated multiplication
//! during compositing and prevents edge artifacts in partial-alpha blending.

/// A 2-dimensional vector using floating-point coordinates.
#[derive(Debug, Clone, Copy, PartialEq)]
pub struct Vec2 {
    /// X coordinate.
    pub x: f32,
    /// Y coordinate.
    pub y: f32,
}

impl Vec2 {
    /// The zero vector.
    pub const ZERO: Self = Self { x: 0.0, y: 0.0 };

    pub const fn new(x: f32, y: f32) -> Self {
        Self { x, y }
    }

    pub fn splat(v: f32) -> Self {
        Self { x: v, y: v }
    }

    pub fn map<F: FnMut(f32) -> f32>(self, mut f: F) -> Self {
        Self {
            x: f(self.x),
            y: f(self.y),
        }
    }
}

impl Default for Vec2 {
    fn default() -> Self {
        Self::ZERO
    }
}

impl std::ops::Add for Vec2 {
    type Output = Self;
    fn add(self, rhs: Self) -> Self {
        Self::new(self.x + rhs.x, self.y + rhs.y)
    }
}

impl std::ops::Sub for Vec2 {
    type Output = Self;
    fn sub(self, rhs: Self) -> Self {
        Self::new(self.x - rhs.x, self.y - rhs.y)
    }
}

impl std::ops::Mul<f32> for Vec2 {
    type Output = Self;
    fn mul(self, rhs: f32) -> Self {
        Self::new(self.x * rhs, self.y * rhs)
    }
}

/// An axis-aligned rectangle defined by its minimum and maximum corner positions.
#[derive(Debug, Clone, Copy, PartialEq)]
pub struct Rect {
    /// Minimum (lower-left) corner.
    pub min: Vec2,
    /// Maximum (upper-right) corner.
    pub max: Vec2,
}

impl Rect {
    /// Constructs a rect from two corners.
    pub fn new(min: Vec2, max: Vec2) -> Self {
        Self { min, max }
    }

    /// Constructs a rect from an origin (min corner) and a size.
    pub fn from_origin_size(origin: Vec2, size: Vec2) -> Self {
        Self {
            min: origin,
            max: Vec2::new(origin.x + size.x, origin.y + size.y),
        }
    }

    /// Width of the rectangle (max.x - min.x).
    #[inline]
    pub fn width(&self) -> f32 {
        self.max.x - self.min.x
    }

    /// Height of the rectangle (max.y - min.y).
    #[inline]
    pub fn height(&self) -> f32 {
        self.max.y - self.min.y
    }

    /// Size of the rectangle as a Vec2.
    #[inline]
    pub fn size(&self) -> Vec2 {
        Vec2::new(self.width(), self.height())
    }

    /// Checks if a point lies inside this rectangle (inclusive on all edges).
    #[inline]
    pub fn contains(&self, p: Vec2) -> bool {
        p.x >= self.min.x && p.x <= self.max.x && p.y >= self.min.y && p.y <= self.max.y
    }

    /// Returns a zero-area rect at the origin.
    #[inline]
    pub fn zero() -> Self {
        Self {
            min: Vec2::ZERO,
            max: Vec2::ZERO,
        }
    }
}

impl Default for Rect {
    fn default() -> Self {
        Self::zero()
    }
}

/// A color with premultiplied alpha channels.
///
/// In premultiplied alpha, the R, G, and B channels are pre-scaled by the
/// alpha value. This eliminates the need for per-pixel multiplication during
/// compositing and eliminates alpha-distribution artifacts at geometry edges.
///
/// To convert from standard (straight) RGBA:
///   premul.r = r * a
///   premul.g = g * a
///   premul.b = b * a
///   premul.a = a
#[derive(Debug, Clone, Copy, PartialEq)]
pub struct Color {
    /// Red channel, premultiplied by alpha.
    pub r: f32,
    /// Green channel, premultiplied by alpha.
    pub g: f32,
    /// Blue channel, premultiplied by alpha.
    pub b: f32,
    /// Alpha channel (0.0 = fully transparent, 1.0 = fully opaque).
    pub a: f32,
}

impl Color {
    /// Fully opaque white (1, 1, 1, 1).
    pub const WHITE: Self = Self {
        r: 1.0,
        g: 1.0,
        b: 1.0,
        a: 1.0,
    };

    /// Fully opaque black (0, 0, 0, 1).
    pub const BLACK: Self = Self {
        r: 0.0,
        g: 0.0,
        b: 0.0,
        a: 1.0,
    };

    /// Fully transparent black (0, 0, 0, 0).
    pub const TRANSPARENT: Self = Self {
        r: 0.0,
        g: 0.0,
        b: 0.0,
        a: 0.0,
    };

    /// Constructs a premultiplied Color from standard (straight) RGBA components.
    ///
    /// The r, g, b values are assumed to be in [0, 1] and will be
    /// pre-multiplied by the alpha value.
    #[inline]
    pub fn from_rgba(r: f32, g: f32, b: f32, a: f32) -> Self {
        Self {
            r: r * a,
            g: g * a,
            b: b * a,
            a,
        }
    }

    /// Constructs a premultiplied Color from already-premultiplied components.
    /// Use this when you already have premultiplied values.
    #[inline]
    pub fn premul(r: f32, g: f32, b: f32, a: f32) -> Self {
        Self { r, g, b, a }
    }

    /// Returns the straight (non-premultiplied) red component.
    ///
    /// Returns 0.0 if alpha is zero to avoid division by zero.
    #[inline]
    pub fn r_straight(&self) -> f32 {
        if self.a == 0.0 {
            0.0
        } else {
            self.r / self.a
        }
    }

    /// Returns the straight (non-premultiplied) green component.
    ///
    /// Returns 0.0 if alpha is zero to avoid division by zero.
    #[inline]
    pub fn g_straight(&self) -> f32 {
        if self.a == 0.0 {
            0.0
        } else {
            self.g / self.a
        }
    }

    /// Returns the straight (non-premultiplied) blue component.
    ///
    /// Returns 0.0 if alpha is zero to avoid division by zero.
    #[inline]
    pub fn b_straight(&self) -> f32 {
        if self.a == 0.0 {
            0.0
        } else {
            self.b / self.a
        }
    }
}

impl Default for Color {
    fn default() -> Self {
        Self::TRANSPARENT
    }
}

/// Blends two premultiplied colors together using a Porter-Duff "over"
/// compositing.
///
/// Both inputs must use premultiplied alpha. The result is also premultiplied.
#[inline]
pub fn blend_colors(base: Color, over: Color) -> Color {
    let out_a = over.a + base.a * (1.0 - over.a);
    if out_a == 0.0 {
        return Color::TRANSPARENT;
    }
    let out_r = over.r + base.r * (1.0 - over.a);
    let out_g = over.g + base.g * (1.0 - over.a);
    let out_b = over.b + base.b * (1.0 - over.a);
    Color {
        r: out_r,
        g: out_g,
        b: out_b,
        a: out_a,
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_vec2_equality_and_zero() {
        assert_eq!(Vec2::new(1.0, 2.0), Vec2::new(1.0, 2.0));
        assert_ne!(Vec2::new(1.0, 2.0), Vec2::new(1.0, 3.0));
        assert_eq!(Vec2::ZERO, Vec2::new(0.0, 0.0));
    }

    #[test]
    fn test_vec2_arithmetic() {
        let a = Vec2::new(1.0, 2.0);
        let b = Vec2::new(3.0, 4.0);
        assert_eq!(a + b, Vec2::new(4.0, 6.0));
        assert_eq!(b - a, Vec2::new(2.0, 2.0));
        assert_eq!(a * 2.0, Vec2::new(2.0, 4.0));
    }

    #[test]
    fn test_vec2_splat_and_map() {
        assert_eq!(Vec2::splat(5.0), Vec2::new(5.0, 5.0));
        let v = Vec2::new(1.0, 2.0);
        assert_eq!(v.map(|x| x * 10.0), Vec2::new(10.0, 20.0));
    }

    #[test]
    fn test_rect_dimensions() {
        let r = Rect::new(Vec2::new(0.0, 0.0), Vec2::new(4.0, 5.0));
        assert_eq!(r.width(), 4.0);
        assert_eq!(r.height(), 5.0);
        assert_eq!(r.size(), Vec2::new(4.0, 5.0));
    }

    #[test]
    fn test_rect_contains() {
        let r = Rect::new(Vec2::new(0.0, 0.0), Vec2::new(10.0, 10.0));
        assert!(r.contains(Vec2::new(5.0, 5.0)));
        assert!(r.contains(Vec2::new(0.0, 0.0)));
        assert!(r.contains(Vec2::new(10.0, 10.0)));
        assert!(!r.contains(Vec2::new(-1.0, 5.0)));
        assert!(!r.contains(Vec2::new(11.0, 5.0)));
    }

    #[test]
    fn test_rect_from_origin_size() {
        let r = Rect::from_origin_size(Vec2::new(1.0, 2.0), Vec2::new(3.0, 4.0));
        assert_eq!(r.min, Vec2::new(1.0, 2.0));
        assert_eq!(r.max, Vec2::new(4.0, 6.0));
        assert_eq!(r.width(), 3.0);
        assert_eq!(r.height(), 4.0);
    }

    #[test]
    fn test_color_from_rgba_is_premultiplied() {
        // straight (0.5, 0.4, 0.3, 0.5) -> premultiplied (0.25, 0.2, 0.15, 0.5)
        let c = Color::from_rgba(0.5, 0.4, 0.3, 0.5);
        assert!((c.r - 0.25).abs() < 1e-6);
        assert!((c.g - 0.2).abs() < 1e-6);
        assert!((c.b - 0.15).abs() < 1e-6);
        assert!((c.a - 0.5).abs() < 1e-6);
    }

    #[test]
    fn test_color_straight_roundtrip() {
        let c = Color::from_rgba(0.8, 0.6, 0.4, 0.5);
        assert!((c.r_straight() - 0.8).abs() < 1e-6);
        assert!((c.g_straight() - 0.6).abs() < 1e-6);
        assert!((c.b_straight() - 0.4).abs() < 1e-6);
        assert!((c.a - 0.5).abs() < 1e-6);
    }

    #[test]
    fn test_color_straight_alpha_zero() {
        let c = Color::from_rgba(1.0, 0.5, 0.25, 0.0);
        assert_eq!(c.r_straight(), 0.0);
        assert_eq!(c.g_straight(), 0.0);
        assert_eq!(c.b_straight(), 0.0);
    }

    #[test]
    fn test_color_constants() {
        assert_eq!(Color::WHITE, Color::from_rgba(1.0, 1.0, 1.0, 1.0));
        assert_eq!(Color::BLACK, Color::from_rgba(0.0, 0.0, 0.0, 1.0));
        assert_eq!(Color::TRANSPARENT, Color::from_rgba(0.0, 0.0, 0.0, 0.0));
    }

    #[test]
    fn test_blend_colors_premultiplied() {
        // Opaque red over opaque blue -> result is red (red fully on top).
        let base = Color::from_rgba(0.0, 0.0, 1.0, 1.0);
        let over = Color::from_rgba(1.0, 0.0, 0.0, 1.0);
        let result = blend_colors(base, over);
        assert!((result.r - 1.0).abs() < 1e-6);
        assert!((result.a - 1.0).abs() < 1e-6);
    }

    #[test]
    fn test_blend_colors_transparent_over_opaque() {
        let base = Color::from_rgba(1.0, 0.0, 0.0, 1.0);
        let over = Color::from_rgba(0.0, 1.0, 0.0, 0.0);
        let result = blend_colors(base, over);
        assert!((result.a - 1.0).abs() < 1e-6);
        assert!((result.r - 1.0).abs() < 1e-6);
        assert!((result.g - 0.0).abs() < 1e-6);
    }

    #[test]
    fn test_color_default_is_transparent() {
        assert_eq!(Color::default(), Color::TRANSPARENT);
    }

    #[test]
    fn test_rect_default_is_zero_area() {
        let r = Rect::default();
        assert_eq!(r.min, Vec2::ZERO);
        assert_eq!(r.max, Vec2::ZERO);
        assert_eq!(r.width(), 0.0);
        assert_eq!(r.height(), 0.0);
    }
}
