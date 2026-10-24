//! Resource counters and leak-soak verification for the tcgui workspace.
//!
//! Provides explicit, non-leaking soak test harness assertions.

use crate::arena::WidgetArena;

pub struct SoakTestHarness {
    iterations: usize,
}

impl SoakTestHarness {
    pub fn new(iterations: usize) -> Self {
        Self { iterations }
    }

    pub fn run_soak_test(&self) -> bool {
        let mut arena = WidgetArena::new();
        for _ in 0..self.iterations {
            let root = arena.create_widget(None, 0);
            let child1 = arena.create_widget(Some(root), 1);
            let child2 = arena.create_widget(Some(child1), 2);
            arena.destroy(root);

            // Stale IDs must return None
            assert!(arena.get(root).is_none());
            assert!(arena.get(child1).is_none());
            assert!(arena.get(child2).is_none());
        }

        let stats = arena.stats();
        stats.widgets == 0
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_soak_plateau_no_leaks() {
        let harness = SoakTestHarness::new(50_000);
        assert!(harness.run_soak_test(), "Memory leaked or widgets remained active after soak test");
    }
}
