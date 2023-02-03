/*
	Copyright 2023 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

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
