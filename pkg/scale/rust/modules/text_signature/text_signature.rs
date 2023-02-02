use regex::Regex;
use lazy_static::lazy_static;
use text_signature::context::Context;

lazy_static! {
    static ref RE: Regex = Regex::new(r"\b\w{4}\b").unwrap();
}

pub fn scale(ctx: &mut Context) -> Result<&mut Context, Box<dyn std::error::Error>> {
    ctx.data = RE.replace_all(ctx.data.as_str(), "wasm").to_string();
    Ok(ctx)
}
