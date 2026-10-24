//! Table-driven unit tests for the layout engine (≥ 30 test cases, errors, idempotency).

use super::*;
use crate::geometry::Vec2;

#[test]
fn test_layout_basic_stack_and_idempotency() {
    let mut tree = LayoutTree::new();
    let child1 = tree.add(LayoutNode::leaf(1));
    let child2 = tree.add(LayoutNode::leaf(2));
    let mut root_node = LayoutNode::row(0, Axis::Horizontal)
        .with_child((SizeValue::Fixed(50.0), SizeValue::Fixed(100.0)), Constraints::default(), child1)
        .with_child((SizeValue::Fixed(50.0), SizeValue::Fixed(100.0)), Constraints::default(), child2);
    root_node.size = (SizeValue::Fixed(200.0), SizeValue::Fixed(100.0));
    let root = tree.add(root_node);

    let engine = LayoutEngine::new(1.0);
    let result1 = engine.layout(&tree, root, Vec2::new(200.0, 100.0));
    let result2 = engine.layout(&tree, root, Vec2::new(200.0, 100.0));

    assert!(!result1.has_errors());
    assert_eq!(result1.rects, result2.rects, "layout must be idempotent");

    let r1 = result1.get(child1).unwrap();
    assert_eq!(r1.min.x, 0.0);
    assert_eq!(r1.width(), 50.0);

    let r2 = result1.get(child2).unwrap();
    assert_eq!(r2.min.x, 50.0);
    assert_eq!(r2.width(), 50.0);
}

#[test]
fn test_layout_ui_scale() {
    let mut tree = LayoutTree::new();
    let child = tree.add(LayoutNode::leaf(1));
    let mut root_node = LayoutNode::column(0)
        .with_child((SizeValue::Fixed(10.0), SizeValue::Fixed(20.0)), Constraints::default(), child);
    root_node.size = (SizeValue::Fixed(100.0), SizeValue::Fixed(100.0));
    let root = tree.add(root_node);

    let engine = LayoutEngine::new(2.0); // 2x scale
    let result = engine.layout(&tree, root, Vec2::new(100.0, 100.0));
    assert!(!result.has_errors());

    let r = result.get(child).unwrap();
    assert_eq!(r.width(), 20.0); // 10 * 2.0
    assert_eq!(r.height(), 40.0); // 20 * 2.0
}

#[test]
fn test_layout_errors_table() {
    // 1. Invalid scale: zero
    {
        let engine = LayoutEngine::new(0.0);
        let mut t = LayoutTree::new();
        let r = t.add(LayoutNode::leaf(0));
        let res = engine.layout(&t, r, Vec2::new(100.0, 100.0));
        assert!(res.has_errors());
    }

    // 2. Invalid scale: negative
    {
        let engine = LayoutEngine::new(-1.5);
        let mut t = LayoutTree::new();
        let r = t.add(LayoutNode::leaf(0));
        let res = engine.layout(&t, r, Vec2::new(100.0, 100.0));
        assert!(res.has_errors());
    }

    // 3. Invalid scale: NaN
    {
        let engine = LayoutEngine::new(f32::NAN);
        let mut t = LayoutTree::new();
        let r = t.add(LayoutNode::leaf(0));
        let res = engine.layout(&t, r, Vec2::new(100.0, 100.0));
        assert!(res.has_errors());
    }

    // 4. Contradictory constraints (min > max)
    {
        let engine = LayoutEngine::new(1.0);
        let mut t = LayoutTree::new();
        let mut n = LayoutNode::leaf(0);
        n.constraints = Constraints { min_width: 100.0, max_width: 50.0, min_height: 0.0, max_height: 0.0 };
        let r = t.add(n);
        let res = engine.layout(&t, r, Vec2::new(200.0, 200.0));
        assert!(res.has_errors());
    }

    // 5. Invalid percent (> 1.0)
    {
        let engine = LayoutEngine::new(1.0);
        let mut t = LayoutTree::new();
        let mut n = LayoutNode::leaf(0);
        n.size = (SizeValue::Percent(1.5), SizeValue::Fixed(10.0));
        let r = t.add(n);
        let res = engine.layout(&t, r, Vec2::new(100.0, 100.0));
        assert!(res.has_errors());
    }

    // 6. Invalid percent (< 0.0)
    {
        let engine = LayoutEngine::new(1.0);
        let mut t = LayoutTree::new();
        let mut n = LayoutNode::leaf(0);
        n.size = (SizeValue::Percent(-0.1), SizeValue::Fixed(10.0));
        let r = t.add(n);
        let res = engine.layout(&t, r, Vec2::new(100.0, 100.0));
        assert!(res.has_errors());
    }

    // 7. Empty container stack
    {
        let engine = LayoutEngine::new(1.0);
        let mut t = LayoutTree::new();
        let r = t.add(LayoutNode::row(0, Axis::Horizontal));
        let res = engine.layout(&t, r, Vec2::new(100.0, 100.0));
        assert!(res.has_errors());
    }

    // 8. Grid mismatch (expects 4, has 2)
    {
        let engine = LayoutEngine::new(1.0);
        let mut t = LayoutTree::new();
        let c1 = t.add(LayoutNode::leaf(1));
        let c2 = t.add(LayoutNode::leaf(2));
        let r = t.add(
            LayoutNode::grid(0, 2, 2, 0.0)
                .with_child((SizeValue::Fixed(10.0), SizeValue::Fixed(10.0)), Constraints::default(), c1)
                .with_child((SizeValue::Fixed(10.0), SizeValue::Fixed(10.0)), Constraints::default(), c2),
        );
        let res = engine.layout(&t, r, Vec2::new(100.0, 100.0));
        assert!(res.has_errors());
    }
}

#[test]
fn test_layout_30_table_cases() {
    let mut engine = LayoutEngine::new(1.0);

    for i in 0..35 {
        let mut tree = LayoutTree::new();
        let scale = 1.0 + (i as f32 * 0.05);
        engine.ui_scale = scale;

        let leaf = tree.add(LayoutNode::leaf(1));
        let width = 10.0 + (i as f32 * 5.0);
        let height = 20.0 + (i as f32 * 2.0);
        let root = tree.add(
            LayoutNode::column(0)
                .with_child((SizeValue::Fixed(width), SizeValue::Fixed(height)), Constraints::default(), leaf),
        );

        let res = engine.layout(&tree, root, Vec2::new(500.0, 500.0));
        assert!(!res.has_errors(), "Case {i} produced unexpected error: {:?}", res.errors);
        let rect = res.get(leaf).unwrap();
        assert!((rect.width() - width * scale).abs() < 1e-4);
        assert!((rect.height() - height * scale).abs() < 1e-4);
    }
}
