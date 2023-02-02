use bad_signature::context::Context;

pub fn scale(ctx: &mut Context) -> Result<&mut Context, Box<dyn std::error::Error>> {
    ctx.data = 30;

    Ok(ctx)
}
