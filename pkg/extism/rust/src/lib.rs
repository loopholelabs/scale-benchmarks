use extism_pdk::*;
use regex::Regex;
use serde::Serialize;
use lazy_static::lazy_static;

#[derive(Serialize)]
struct Output {
    pub matches: String,
}

lazy_static! {
    static ref RE: Regex = Regex::new(r"\b\w{4}\b").unwrap();
}

#[plugin_fn]
pub unsafe fn match_regex(input: String) -> FnResult<Json<Output>> {
    let output = Output {
        matches: RE.replace_all(input.as_str(), "wasm").to_string(),
    };

    Ok(Json(output))
}
