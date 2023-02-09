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

use polyglot_rs::{Decoder, Encoder};
use std::io::Cursor;

pub trait Encode {
    fn encode<'a>(
        &'a self,
        b: &'a mut Cursor<Vec<u8>>,
    ) -> Result<&mut Cursor<Vec<u8>>, Box<dyn std::error::Error>>;
    fn internal_error<'a>(&'a self, b: &'a mut Cursor<Vec<u8>>, error: Box<dyn std::error::Error>);
}

pub trait Decode {
    fn decode(b: &mut Cursor<&mut Vec<u8>>) -> Result<Option<Self>, Box<dyn std::error::Error>>
    where
        Self: Sized;
}

#[derive(Clone)]
pub struct TextContext {
    pub data: String,
}

impl Encode for TextContext {
    fn encode<'a>(
        &'a self,
        b: &'a mut Cursor<Vec<u8>>,
    ) -> Result<&mut Cursor<Vec<u8>>, Box<dyn std::error::Error>> {
        b.encode_string(&self.data)?;
        Ok(b)
    }

    fn internal_error<'a>(&'a self, b: &'a mut Cursor<Vec<u8>>, error: Box<dyn std::error::Error>) {
        b.encode_error(error).unwrap();
    }
}

impl Decode for TextContext {
    fn decode(
        b: &mut Cursor<&mut Vec<u8>>,
    ) -> Result<Option<TextContext>, Box<dyn std::error::Error>> {
        if b.decode_none() {
            return Ok(None);
        }

        if let Ok(error) = b.decode_error() {
            return Err(error);
        }

        Ok(Some(TextContext {
            data: b.decode_string()?,
        }))
    }
}
