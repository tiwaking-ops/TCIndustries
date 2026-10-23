# TCIndustries GUI Public Interfaces (`docs/interfaces.md`)

**Status:** FROZEN  
**Scope:** Public API definitions for `tcgui-core`.

## 1. Geometry
- `Vec2 { x: f32, y: f32 }`
- `Rect { min: Vec2, max: Vec2 }`
- `Color { r: f32, g: f32, b: f32, a: f32 }` (premultiplied)

### Vec2
- `Vec2::ZERO` — zero vector constant.
- `Vec2::new(x, y)` — construct.
- `Vec2::splat(v)` — both components set to `v`.
- `Vec2::map(f)` — apply unary f32 function to each component.
- `Vec2::default()` == `ZERO`.
- Operators: `Add (Vec2)`, `Sub (Vec2)`, `Mul (f32)`.

### Rect
- `Rect::new(min, max)` — from two corners.
- `Rect::from_origin_size(origin, size)` — from origin + size.
- `Rect::zero()` — zero-area rect at origin.
- `Rect::default()` == `zero()`.
- `Rect::width()`, `Rect::height()`, `Rect::size()`, `Rect::contains(p)`.

### Color (premultiplied alpha)
- `Color::WHITE`, `Color::BLACK`, `Color::TRANSPARENT` — constants.
- `Color::from_rgba(r, g, b, a)` — construct from straight (non-premultiplied) RGBA.
- `Color::premul(r, g, b, a)` — construct from already-premultiplied values.
- `Color::r_straight()`, `Color::g_straight()`, `Color::b_straight()` — recover non-premultiplied channels (0.0 if alpha == 0).
- `Color::default()` == `TRANSPARENT`.
- Free function `blend_colors(base, over) -> Color` — Porter-Duff "over" compositing of two premultiplied colors.

## 2. Arena Tree
- `WidgetId { index: usize, generation: u32 }` — generational arena handle; stale IDs return `None`.
- `WidgetNode` fields: `id`, `parent: Option<WidgetId>`, `children: Vec<WidgetId>`, `z_order: i32`, `active: bool`.
- `WidgetArena` methods:
  - `new()` / `default()`
  - `create_widget(parent: Option<WidgetId>, z_order: i32) -> WidgetId` — creates a node; a stale `parent` is silently treated as a root.
  - `get(id) -> Option<&WidgetNode>` — returns `None` for stale / destroyed ids.
  - `get_mut(id) -> Option<&mut WidgetNode>` — mutable borrow, `None` if stale.
  - `destroy(id)` — destroys the node **and all its descendants** (subtree). Stale ids are ignored.
  - `reparent(child: WidgetId, new_parent: Option<WidgetId>)` — moves `child` (and its subtree) under `new_parent`; `None` makes it a root; reparenting a node under its own descendant is rejected (cycle guard). Stale child or stale `new_parent` is handled gracefully (treated as root).
  - `sort_children(id)` — stable-reorders the node's child list ascending by `z_order`; stale id is a no-op.
  - `stats() -> ArenaStats`.

### ArenaStats
- `widgets: usize` — number of live widgets.
- `capacity: usize` — number of allocated slots (live + recycled free-list).

## 3. Data Binding
- `UiModel` key-value store with typed binding subscriptions and automatic teardown on destroy.
