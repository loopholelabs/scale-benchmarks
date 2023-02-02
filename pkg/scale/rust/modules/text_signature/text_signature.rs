use regex::Regex;
use lazy_static::lazy_static;
use text_signature::context::Context;

lazy_static! {
    static ref RE: Regex = Regex::new("p([a-z]+)ch").unwrap();
}

pub fn scale(ctx: &mut Context) -> Result<&mut Context, Box<dyn std::error::Error>> {
    ctx.data = RE.find(ctx.data.as_str()).unwrap().as_str().to_string();
    Ok(ctx)
}
