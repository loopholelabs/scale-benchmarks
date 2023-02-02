use extism_pdk::*;
use regex::Regex;
use serde::Serialize;

#[derive(Serialize)]
struct Output {
    pub matches: String,
}

#[plugin_fn]
pub unsafe fn match_regex(input: String) -> FnResult<Json<Output>> {
    let r = Regex::new("p([a-z]+)ch").unwrap();

    let output = Output {
        matches: r.find(input.as_str()).unwrap().as_str().to_string(),
    };

    Ok(Json(output))
}
