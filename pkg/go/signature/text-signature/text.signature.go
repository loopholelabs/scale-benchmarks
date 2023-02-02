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
