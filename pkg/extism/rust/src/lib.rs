use std::io::Cursor;
use extism_pdk::*;
use lazy_static::lazy_static;
use regex::Regex;
use wee_alloc;
use text_signature::text_signature::{Decode, Encode, TextContext};


lazy_static! {
    static ref RE: Regex = Regex::new(r"\b\w{4}\b").unwrap();
}

#[global_allocator]
pub static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

#[plugin_fn]
pub unsafe fn match_regex(input: Vec<u8>) -> FnResult<Vec<u8>> {
    let mut raw = input;
    let mut cursor = Cursor::new(&mut raw);
    let result = TextContext::decode(&mut cursor);
    return match result {
        Ok(context) => {
            let mut s = context.unwrap();
            s.data = RE.replace_all(s.data.as_str(), "wasm").to_string();

            let mut cursor = Cursor::new(Vec::new());
            cursor = match TextContext::encode(s.clone(), &mut cursor) {
                Ok(_) => cursor,
                Err(err) => {
                    panic!("Error: {}", err);
                },
            };

            Ok(cursor.into_inner())
        }
        Err(e) => {
            panic!("Error: {}", e);
        },
    };

}
