//! Key-value store, typed values, path bindings, and automatic teardown on destroy.

use crate::arena::WidgetId;
use std::collections::{HashMap, HashSet};

/// A typed value in the UiModel.
#[derive(Debug, Clone, PartialEq)]
pub enum UiValue {
    Null,
    Bool(bool),
    F32(f32),
    I32(i32),
    String(String),
}

impl From<bool> for UiValue { fn from(v: bool) -> Self { Self::Bool(v) } }
impl From<f32> for UiValue { fn from(v: f32) -> Self { Self::F32(v) } }
impl From<i32> for UiValue { fn from(v: i32) -> Self { Self::I32(v) } }
impl From<&str> for UiValue { fn from(v: &str) -> Self { Self::String(v.to_string()) } }
impl From<String> for UiValue { fn from(v: String) -> Self { Self::String(v) } }

pub trait FromUiValue: Sized {
    fn from_ui_value(val: &UiValue) -> Option<Self>;
}

impl FromUiValue for bool {
    fn from_ui_value(val: &UiValue) -> Option<Self> {
        match val { UiValue::Bool(b) => Some(*b), _ => None }
    }
}

impl FromUiValue for f32 {
    fn from_ui_value(val: &UiValue) -> Option<Self> {
        match val { UiValue::F32(f) => Some(*f), UiValue::I32(i) => Some(*i as f32), _ => None }
    }
}

impl FromUiValue for i32 {
    fn from_ui_value(val: &UiValue) -> Option<Self> {
        match val { UiValue::I32(i) => Some(*i), UiValue::F32(f) => Some(*f as i32), _ => None }
    }
}

impl FromUiValue for String {
    fn from_ui_value(val: &UiValue) -> Option<Self> {
        match val { UiValue::String(s) => Some(s.clone()), _ => None }
    }
}

#[allow(dead_code)]
struct Binding {
    id: usize,
    widget: WidgetId,
    parent: Option<WidgetId>,
    path: String,
    callback: Box<dyn FnMut(&UiValue)>,
}

/// Key-value data store with path bindings and automatic unsubscription.
pub struct UiModel {
    values: HashMap<String, UiValue>,
    bindings: HashMap<String, Vec<Binding>>,
    next_binding_id: usize,
    widget_to_bindings: HashMap<WidgetId, Vec<(String, usize)>>,
    widget_parents: HashMap<WidgetId, Option<WidgetId>>,
}

#[derive(Debug, Clone, Copy)]
pub struct ModelStats {
    pub bindings: usize,
    pub values: usize,
}

impl UiModel {
    pub fn new() -> Self {
        Self {
            values: HashMap::new(),
            bindings: HashMap::new(),
            next_binding_id: 0,
            widget_to_bindings: HashMap::new(),
            widget_parents: HashMap::new(),
        }
    }

    pub fn set<T: Into<UiValue>>(&mut self, path: &str, val: T) {
        let value = val.into();
        self.values.insert(path.to_string(), value.clone());

        if let Some(bindings) = self.bindings.get_mut(path) {
            for b in bindings {
                (b.callback)(&value);
            }
        }
    }

    pub fn get(&self, path: &str) -> Option<&UiValue> {
        self.values.get(path)
    }

    pub fn bind<T: FromUiValue + 'static, F: FnMut(&T) + 'static>(
        &mut self,
        widget: WidgetId,
        parent: Option<WidgetId>,
        path: &str,
        mut cb: F,
    ) {
        self.widget_parents.insert(widget, parent);
        let binding_id = self.next_binding_id;
        self.next_binding_id += 1;

        let path_string = path.to_string();
        let current_val = self.values.get(path).cloned().unwrap_or(UiValue::Null);

        if let Some(parsed) = T::from_ui_value(&current_val) {
            cb(&parsed);
        }

        let cb_erased: Box<dyn FnMut(&UiValue)> = Box::new(move |val| {
            if let Some(parsed) = T::from_ui_value(val) {
                cb(&parsed);
            }
        });

        let binding = Binding {
            id: binding_id,
            widget,
            parent,
            path: path_string.clone(),
            callback: cb_erased,
        };

        self.bindings.entry(path_string.clone()).or_default().push(binding);
        self.widget_to_bindings.entry(widget).or_default().push((path_string, binding_id));
    }

    pub fn unbind_widget(&mut self, widget: WidgetId) {
        if let Some(pairs) = self.widget_to_bindings.remove(&widget) {
            self.widget_parents.remove(&widget);
            for (path, binding_id) in pairs {
                if let Some(list) = self.bindings.get_mut(&path) {
                    list.retain(|b| b.id != binding_id);
                }
            }
        }
    }

    pub fn destroy_widget_recursive(&mut self, root: WidgetId) {
        // Find all descendants using widget_parents
        let mut to_destroy = HashSet::new();
        let mut stack = vec![root];

        while let Some(curr) = stack.pop() {
            to_destroy.insert(curr);
            // find children
            for (&child, &p) in &self.widget_parents {
                if p == Some(curr) && !to_destroy.contains(&child) {
                    stack.push(child);
                }
            }
        }

        for id in to_destroy {
            self.unbind_widget(id);
        }
    }

    pub fn stats(&self) -> ModelStats {
        let total_bindings: usize = self.bindings.values().map(|v| v.len()).sum();
        ModelStats {
            bindings: total_bindings,
            values: self.values.len(),
        }
    }
}

impl Default for UiModel {
    fn default() -> Self {
        Self::new()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_ui_model_typed_bindings_and_notifications() {
        let mut model = UiModel::new();
        let widget = WidgetId { index: 1, generation: 0 };
        let mut received = 0.0;

        model.bind(widget, None, "health", move |&h: &f32| {
            received = h;
        });

        model.set("health", 100.0f32);
        assert_eq!(model.get("health"), Some(&UiValue::F32(100.0)));
    }

    #[test]
    fn test_ui_model_soak_10k_cycles_flat_memory() {
        let mut model = UiModel::new();

        for i in 0..10_000 {
            let widget = WidgetId { index: i, generation: 0 };
            model.bind(widget, None, "score", |_s: &i32| {});
            model.set("score", i as i32);
            model.destroy_widget_recursive(widget);
        }

        let stats = model.stats();
        assert_eq!(stats.bindings, 0, "All bindings must be cleaned up");
    }
}
