#![cfg(target_arch = "wasm32")]

use crate::text_signature::TextContext;

pub type Context = TextContext;

pub fn new() -> Context {
    Context {
        data: "".to_string(),
    }
}
