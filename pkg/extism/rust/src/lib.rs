use polyglot_rs::{Decoder, Encoder};
use std::{io::Cursor, vec};

use extism_pdk::*;
use lazy_static::lazy_static;
use regex::Regex;

lazy_static! {
    static ref RE: Regex = Regex::new(r"\b\w{4}\b").unwrap();
}

#[plugin_fn]
pub unsafe fn match_regex(input: Vec<u8>) -> FnResult<Vec<u8>> {
    let mut raw = input;

    let mut decoder = Cursor::new(&mut raw);
    let i = decoder.decode_string().unwrap();

    let matches = RE.replace_all(i.as_str(), "wasm").to_string();

    let mut encoder = Cursor::new(vec![]);
    encoder.encode_string(&matches).unwrap();

    Ok(encoder.get_mut().to_vec())
}
