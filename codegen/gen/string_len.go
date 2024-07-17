package gen

import (
	"context"

	"github.com/oiweiwei/go-msrpc/midl"
)

var (
	// UTF16 String.
	UTF16LenFunc = "ndr.UTF16Len"
	// Char String.
	ByteLenFunc = "ndr.CharLen"
	// UTF16 Null-Terminated String.
	UTF16NLenFunc = "ndr.UTF16NLen"
	// Char Null-Terminated String.
	ByteNLenFunc = "ndr.CharNLen"
	// MultiSzLen
	MultiSzLenFunc = "ndr.MultiSzLen"
)

// UTF16Len function returns the UTF16 length-count function call on variable `n`.
func (p *Generator) UTF16Len(ctx context.Context, n string, isNullTerminated bool) string {
	if isNullTerminated {
		return p.B(UTF16NLenFunc, n)
	}
	return p.B(UTF16LenFunc, n)
}

// ByteLen function returns the byte length-count function call on variable `n`.
func (p *Generator) ByteLen(ctx context.Context, n string, isNullTerminated bool) string {
	if isNullTerminated {
		return p.B(ByteNLenFunc, n)
	}
	return p.B(ByteLenFunc, n)
}

// DataLen function returns the proper length-count function call on variable `n`.
func (p *Generator) DataLen(ctx context.Context, field *midl.Field, scopes *Scopes, n string, defaultTypeSize ...string) string {

	if field.Attrs.Format.MultiSize {
		return p.B(MultiSzLenFunc, n)
	}

	if dim := scopes.Dim(); dim.IsString {

		switch scopes.Next().Kind() {
		case midl.TypeChar, midl.TypeUChar, midl.TypeUint8, midl.TypeInt8:
			return p.ByteLen(ctx, n, dim.IsNullTerminated || field.Attrs.Format.NullTerminated)
		case midl.TypeWChar, midl.TypeUint16, midl.TypeInt16:
			return p.UTF16Len(ctx, n, dim.IsNullTerminated || field.Attrs.Format.NullTerminated)
		}
	}

	if len(defaultTypeSize) > 0 {
		if defaultTypeSize[0] == "" {
			return p.Len(n)
		}
		return p.B(defaultTypeSize[0], p.Len(n))
	}

	return p.B("uint64", p.Len(n))
}
