//! Widget arena: a generational-index tree of widget nodes.
//!
//! `WidgetId` is a generational handle that prevents stale references. When a
//! node is destroyed its generation is bumped (and the slot recycled via the
//! free-list), so any outstanding `WidgetId` pointing at that slot will be
//! rejected by `get` / `get_mut` / `destroy`.

/// Unique handle to a widget node in the arena.
#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub struct WidgetId {
    pub index: usize,
    pub generation: u32,
}

/// A single node in the widget tree.
#[derive(Debug, Clone)]
pub struct WidgetNode {
    /// This node's id.
    pub id: WidgetId,
    /// Parent widget id, or `None` for a root node.
    pub parent: Option<WidgetId>,
    /// Direct child widget ids.
    pub children: Vec<WidgetId>,
    /// Z-order (draw) priority. Higher values draw on top.
    pub z_order: i32,
    /// Whether the node is still active.
    pub active: bool,
}

/// An arena of widget nodes that manages lifetime via generational indices.
pub struct WidgetArena {
    nodes: Vec<Option<WidgetNode>>,
    generations: Vec<u32>,
    free_list: Vec<usize>,
    active_count: usize,
}

impl WidgetArena {
    pub fn new() -> Self {
        Self {
            nodes: Vec::new(),
            generations: Vec::new(),
            free_list: Vec::new(),
            active_count: 0,
        }
    }

    /// Creates a widget attached to `parent`, with the given `z_order`.
    pub fn create_widget(&mut self, parent: Option<WidgetId>, z_order: i32) -> WidgetId {
        let index = if let Some(idx) = self.free_list.pop() {
            idx
        } else {
            let idx = self.nodes.len();
            self.nodes.push(None);
            self.generations.push(0);
            idx
        };

        let generation = self.generations[index];
        let id = WidgetId { index, generation };

        let resolved_parent = if let Some(p_id) = parent {
            if let Some(Some(p_node)) = self.nodes.get_mut(p_id.index) {
                if p_node.id.generation == p_id.generation && p_node.active {
                    p_node.children.push(id);
                    Some(p_id)
                } else {
                    None
                }
            } else {
                None
            }
        } else {
            None
        };

        let node = WidgetNode {
            id,
            parent: resolved_parent,
            children: Vec::new(),
            z_order,
            active: true,
        };

        self.nodes[index] = Some(node);
        self.active_count += 1;

        id
    }

    /// Returns a reference to a live node, or `None` for a stale / destroyed id.
    pub fn get(&self, id: WidgetId) -> Option<&WidgetNode> {
        let node = self.nodes.get(id.index)?.as_ref()?;
        if node.id.generation == id.generation && node.active {
            Some(node)
        } else {
            None
        }
    }

    /// Returns a mutable reference to a live node, or `None` otherwise.
    pub fn get_mut(&mut self, id: WidgetId) -> Option<&mut WidgetNode> {
        let nodes = &mut self.nodes;
        let node = nodes.get_mut(id.index)?.as_mut()?;
        if node.id.generation == id.generation && node.active {
            Some(node)
        } else {
            None
        }
    }

    /// Removes the node with the given id, and all of its descendants.
    pub fn destroy(&mut self, id: WidgetId) {
        let mut to_destroy = vec![id];
        let mut queue = vec![id];

        while let Some(curr_id) = queue.pop() {
            if let Some(node) = self.get(curr_id) {
                let children = node.children.clone();
                for c in children {
                    to_destroy.push(c);
                    queue.push(c);
                }
            }
        }

        for d_id in to_destroy {
            if let Some(node_opt) = self.nodes.get_mut(d_id.index) {
                if let Some(node) = node_opt.take() {
                    if node.id.generation == d_id.generation {
                        if let Some(p_id) = node.parent {
                            if let Some(Some(p_node)) = self.nodes.get_mut(p_id.index) {
                                p_node.children.retain(|&c| c != d_id);
                            }
                        }
                        self.generations[d_id.index] += 1;
                        self.free_list.push(d_id.index);
                        self.active_count -= 1;
                    }
                }
            }
        }
    }

    /// Moves `child` to be a child of `new_parent`, preserving its subtree.
    pub fn reparent(&mut self, child: WidgetId, new_parent: Option<WidgetId>) {
        let (child_parent_old, _child_z) = match self.get(child) {
            Some(node) => (node.parent, node.z_order),
            None => return,
        };

        if let Some(old_p) = child_parent_old {
            if let Some(Some(p_node)) = self.nodes.get_mut(old_p.index) {
                if p_node.id == old_p {
                    p_node.children.retain(|&c| c != child);
                }
            }
        }

        let resolved_new_parent = if let Some(np) = new_parent {
            if let Some(Some(np_node)) = self.nodes.get_mut(np.index) {
                if np_node.id == np && np_node.active {
                    if self.is_descendant_of(np, child) {
                        return;
                    }
                    Some(np)
                } else {
                    None
                }
            } else {
                None
            }
        } else {
            None
        };

        if let Some(np) = resolved_new_parent {
            if let Some(Some(np_node)) = self.nodes.get_mut(np.index) {
                np_node.children.push(child);
            }
        }

        if let Some(Some(node)) = self.nodes.get_mut(child.index) {
            if node.id == child {
                node.parent = resolved_new_parent;
            }
        }
    }

    fn is_descendant_of(&self, id: WidgetId, ancestor: WidgetId) -> bool {
        let mut current = id;
        while let Some(node) = self.get(current) {
            match node.parent {
                Some(p) => {
                    if p == ancestor {
                        return true;
                    }
                    current = p;
                }
                None => return false,
            }
        }
        false
    }

    /// Sorts the child list of the node `id` by `z_order` (ascending).
    pub fn sort_children(&mut self, id: WidgetId) {
        let children_z = {
            let node = match self.get(id) {
                Some(n) => n,
                None => return,
            };
            node.children.iter().map(|&c| {
                let z = self.get(c).map(|n| n.z_order).unwrap_or(i32::MIN);
                (c, z)
            }).collect::<Vec<_>>()
        };

        if let Some(Some(node)) = self.nodes.get_mut(id.index) {
            if node.id == id {
                node.children.sort_by_key(|c| {
                    children_z.iter().find(|(id_val, _)| id_val == c).map(|&(_, z)| z).unwrap_or(i32::MIN)
                });
            }
        }
    }

    /// Raises the widget at `id` above all of its siblings (z-order raise-on-click).
    ///
    /// Sets the node's `z_order` to one greater than the maximum z-order among
    /// its siblings, then re-sorts the parent's child list. Stale ids and cycle
    /// attempts are no-ops.
    pub fn bring_to_front(&mut self, id: WidgetId) {
        let parent_id = match self.get(id) {
            Some(node) => node.parent,
            None => return,
        };

        let sibling_max = parent_id
            .and_then(|p| self.get(p))
            .map(|p_node| {
                p_node
                    .children
                    .iter()
                    .filter_map(|&c| self.get(c).map(|n| n.z_order))
                    .max()
                    .unwrap_or(0)
            })
            .unwrap_or(0);

        if let Some(Some(node)) = self.nodes.get_mut(id.index) {
            if node.id == id {
                node.z_order = sibling_max + 1;
            }
        }

        if let Some(p) = parent_id {
            self.sort_children(p);
        }
    }

    pub fn stats(&self) -> ArenaStats {
        ArenaStats {
            widgets: self.active_count,
            capacity: self.nodes.len(),
        }
    }
}

impl Default for WidgetArena {
    fn default() -> Self {
        Self::new()
    }
}

#[derive(Debug, Clone, Copy)]
pub struct ArenaStats {
    pub widgets: usize,
    pub capacity: usize,
}

#[cfg(test)]
mod tests {
    use super::*;

    // ------------------------------------------------------------------------
    // Lifecycle & generational ID reuse
    // ------------------------------------------------------------------------

    /// Verify arena stats return to baseline after 10,000 create/destroy cycles
    /// and that stale (destroyed) IDs resolve to None.
    #[test]
    fn test_arena_returns_to_baseline_after_10k_cycles() {
        let mut arena = WidgetArena::new();

        // Initial state: zero widgets, zero capacity.
        let base = arena.stats();
        assert_eq!(base.widgets, 0);
        assert_eq!(base.capacity, 0);

        // Run 10,000 create/destroy cycles.
        for _ in 0..10_000 {
            let id = arena.create_widget(None, 0);
            assert!(arena.get(id).is_some());
            arena.destroy(id);
            assert!(arena.get(id).is_none(), "stale ID must return None");
        }

        // Verify stats are exactly back at baseline (active widgets == 0).
        let stats = arena.stats();
        assert_eq!(
            stats.widgets,
            0,
            "widgets should be 0 after all cycles destroyed"
        );
    }

    /// Verify stale IDs return None after destroy.
    #[test]
    fn test_stale_id_returns_none() {
        let mut arena = WidgetArena::new();
        let id = arena.create_widget(None, 0);
        arena.destroy(id);
        assert!(arena.get(id).is_none());
        assert!(arena.get_mut(id).is_none());
    }

    /// Verify that generation increments on reuse so old IDs stay invalid.
    #[test]
    fn test_generation_bumps_on_destroy() {
        let mut arena = WidgetArena::new();
        let id0 = arena.create_widget(None, 0);
        let gen0 = id0.generation;
        arena.destroy(id0);

        // Reuse the slot.
        let id1 = arena.create_widget(None, 0);
        assert_eq!(
            id1.generation,
            gen0 + 1,
            "generation should increment on slot reuse"
        );
        // Old ID must still be stale.
        assert!(arena.get(id0).is_none());
    }

    // ------------------------------------------------------------------------
    // Tree structure: parent / children
    // ------------------------------------------------------------------------

    #[test]
    fn test_create_widget_with_valid_parent() {
        let mut arena = WidgetArena::new();
        let parent = arena.create_widget(None, 0);
        let child = arena.create_widget(Some(parent), 1);

        let parent_node = arena.get(parent).unwrap();
        assert_eq!(parent_node.children, vec![child]);
        assert_eq!(arena.get(child).unwrap().parent, Some(parent));
    }

    #[test]
    fn test_create_widget_with_stale_parent_becomes_root() {
        let mut arena = WidgetArena::new();
        let parent = arena.create_widget(None, 0);
        arena.destroy(parent);

        let child = arena.create_widget(Some(parent), 0);
        assert!(
            arena.get(child).unwrap().parent.is_none(),
            "stale parent should be treated as root"
        );
    }

    #[test]
    fn test_destroy_subtree_removes_descendants() {
        let mut arena = WidgetArena::new();
        let root = arena.create_widget(None, 0);
        let a = arena.create_widget(Some(root), 1);
        let b = arena.create_widget(Some(root), 2);
        let c = arena.create_widget(Some(a), 3);

        arena.destroy(root);

        assert!(arena.get(root).is_none());
        assert!(arena.get(a).is_none());
        assert!(arena.get(b).is_none());
        assert!(arena.get(c).is_none());
        assert_eq!(arena.stats().widgets, 0);
    }

    #[test]
    fn test_destroy_leaf_does_not_affect_siblings() {
        let mut arena = WidgetArena::new();
        let root = arena.create_widget(None, 0);
        let a = arena.create_widget(Some(root), 1);
        let b = arena.create_widget(Some(root), 2);

        arena.destroy(a);

        assert!(arena.get(root).is_some());
        assert!(arena.get(b).is_some());
        assert!(arena.get(a).is_none());
        assert_eq!(arena.stats().widgets, 2);
    }

    // ------------------------------------------------------------------------
    // Reparenting
    // ------------------------------------------------------------------------

    #[test]
    fn test_reparent_to_new_parent() {
        let mut arena = WidgetArena::new();
        let root = arena.create_widget(None, 0);
        let a = arena.create_widget(Some(root), 1);
        let b = arena.create_widget(None, 0);

        arena.reparent(b, Some(root));

        let root_node = arena.get(root).unwrap();
        assert!(root_node.children.contains(&a));
        assert!(root_node.children.contains(&b));
        assert_eq!(arena.get(b).unwrap().parent, Some(root));
    }

    #[test]
    fn test_reparent_to_none_makes_root() {
        let mut arena = WidgetArena::new();
        let root = arena.create_widget(None, 0);
        let child = arena.create_widget(Some(root), 1);

        arena.reparent(child, None);

        assert!(arena.get(child).unwrap().parent.is_none());
        assert!(!arena.get(root).unwrap().children.contains(&child));
    }

    #[test]
    fn test_reparent_stale_child_is_noop() {
        let mut arena = WidgetArena::new();
        let root = arena.create_widget(None, 0);
        let child = arena.create_widget(Some(root), 1);
        arena.destroy(child);

        arena.reparent(child, Some(root));

        assert!(arena.get(child).is_none());
    }

    #[test]
    fn test_reparent_to_descendant_is_noop() {
        let mut arena = WidgetArena::new();
        let root = arena.create_widget(None, 0);
        let child = arena.create_widget(Some(root), 1);
        let grandchild = arena.create_widget(Some(child), 2);

        // Attempt to reparent root under its grandchild (cycle).
        arena.reparent(root, Some(grandchild));

        // root should still be a root.
        assert!(arena.get(root).unwrap().parent.is_none());
        // grandchild's parent should still be child.
        assert_eq!(arena.get(grandchild).unwrap().parent, Some(child));
    }

    // ------------------------------------------------------------------------
    // Z-order sorting
    // ------------------------------------------------------------------------

    #[test]
    fn test_sort_children_z_order() {
        let mut arena = WidgetArena::new();
        let root = arena.create_widget(None, 0);

        // Insert in arbitrary order.
        let low = arena.create_widget(Some(root), 0);
        let mid = arena.create_widget(Some(root), 10);
        let high = arena.create_widget(Some(root), 20);

        arena.sort_children(root);

        let children = &arena.get(root).unwrap().children;
        assert_eq!(children[0], low);
        assert_eq!(children[1], mid);
        assert_eq!(children[2], high);
    }

    #[test]
    fn test_sort_children_stale_id_is_noop() {
        let mut arena = WidgetArena::new();
        let root = arena.create_widget(None, 0);
        arena.create_widget(Some(root), 5);
        arena.create_widget(Some(root), 3);

        let root_copy = arena.create_widget(None, 0);
        arena.destroy(root);

        // Must not panic.
        arena.sort_children(root);
        arena.sort_children(root_copy);
    }

    // ------------------------------------------------------------------------
    // ArenaStats
    // ------------------------------------------------------------------------

    #[test]
    fn test_stats_reflect_active_count() {
        let mut arena = WidgetArena::new();
        assert_eq!(arena.stats().widgets, 0);

        let root = arena.create_widget(None, 0);
        let a = arena.create_widget(Some(root), 0);
        let _b = arena.create_widget(Some(a), 0);

        assert_eq!(arena.stats().widgets, 3);

        arena.destroy(a);
        assert_eq!(arena.stats().widgets, 1); // only root survives

        arena.destroy(root);
        assert_eq!(arena.stats().widgets, 0);
    }

    #[test]
    fn test_stats_capacity_grows_and_stays() {
        let mut arena = WidgetArena::new();

        for i in 0..50 {
            let id = arena.create_widget(None, i as i32);
            arena.destroy(id);
        }

        let stats = arena.stats();
        assert!(stats.capacity > 0, "capacity should reflect allocated slots");
        assert_eq!(stats.widgets, 0, "no live widgets should remain");
    }

    // ------------------------------------------------------------------------
    // Mixed stress: 10k create/destroy with subtrees
    // ------------------------------------------------------------------------

    #[test]
    fn test_10k_subtree_cycles_return_to_baseline() {
        let mut arena = WidgetArena::new();

        for _ in 0..10_000 {
            let root = arena.create_widget(None, 0);
            let _a = arena.create_widget(Some(root), 1);
            let _b = arena.create_widget(Some(root), 2);
            let _c = arena.create_widget(Some(root), 3);
            arena.destroy(root);
        }

        let stats = arena.stats();
        assert_eq!(
            stats.widgets,
            0,
            "all subtrees destroyed, widgets must be 0"
        );
    }
}

