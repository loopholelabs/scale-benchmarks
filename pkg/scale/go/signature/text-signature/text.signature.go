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

package text

import (
	"errors"

	"github.com/loopholelabs/polyglot-go"
)

var (
	NilDecode = errors.New("cannot decode into a nil root struct")
)

type TextContext struct {
	Data string
}

func NewTextContext() *TextContext {
	return &TextContext{}
}

func (x *TextContext) error(b *polyglot.Buffer, err error) {
	polyglot.Encoder(b).Error(err)
}

func (x *TextContext) internalEncode(b *polyglot.Buffer) {
	if x == nil {
		polyglot.Encoder(b).Nil()
	} else {
		polyglot.Encoder(b).String(x.Data)
	}
}

func (x *TextContext) internalDecode(b []byte) error {
	if x == nil {
		return NilDecode
	}
	d := polyglot.GetDecoder(b)
	defer d.Return()
	return x.decode(d)
}

func (x *TextContext) decode(d *polyglot.Decoder) error {
	if d.Nil() {
		return nil
	}

	err, _ := d.Error()
	if err != nil {
		return err
	}
	x.Data, err = d.String()
	if err != nil {
		return err
	}
	return nil
}
