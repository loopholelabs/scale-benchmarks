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

use std::io::Cursor;
use extism_pdk::*;
use lazy_static::lazy_static;
use regex::Regex;
use text_signature::text_signature::{Decode, Encode, TextContext};

lazy_static! {
    static ref RE: Regex = Regex::new(r"\b\w{4}\b").unwrap();
}

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
