use text_signature::context::Context;

pub fn scale(ctx: &mut Context) -> Result<&mut Context, Box<dyn std::error::Error>> {
    ctx.data = "Hello from Rust!".to_string();

    Ok(ctx)
}
