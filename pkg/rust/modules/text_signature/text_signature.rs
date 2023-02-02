use regex::Regex;
use text_signature::context::Context;

pub fn scale(ctx: &mut Context) -> Result<&mut Context, Box<dyn std::error::Error>> {
    let r = Regex::new("p([a-z]+)ch").unwrap();

    ctx.data = r.find("peach").unwrap().as_str().to_string();

    Ok(ctx)
}
