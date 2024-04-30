//go:generate go run ./maketables.go -o tables.go

package oem

// package oem privdes the set of original equipment manufacturer (OEM)
// encoders/decoders.

import (
	"context"
	"sync"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"

	"github.com/indece-official/go-ebcdic"
)

// The OEM encoding.
type Encoding interface {
	// Encode function encodes the string to the OEM encoding.
	Encode(string) ([]byte, error)
	// Decode function decodes the bytes from the OEM encoding.
	Decode([]byte) (string, error)
}

var (
	ASCII                    = func() Encoding { return ascii{} }
	CP437_DOSLatinUS         = func() Encoding { return Charmap{charmap.CodePage437} }
	CP500_IBMInternational   = func() Encoding { return EBCDIC{ebcdic.EBCDIC500} }
	CP737_DOSGreek           = func() Encoding { return NewCharmap2(cp737_DOSGreek) }
	CP775_DOSBaltRim         = func() Encoding { return NewCharmap2(cp775_DOSBaltRim) }
	CP850_DOSLatin1          = func() Encoding { return Charmap{charmap.CodePage850} }
	CP852_DOSLatin2          = func() Encoding { return Charmap{charmap.CodePage852} }
	CP855_DOSCyrillic        = func() Encoding { return Charmap{charmap.CodePage855} }
	CP857_DOSTurkish         = func() Encoding { return NewCharmap2(cp857_DOSTurkish) }
	CP860_DOSPortuguese      = func() Encoding { return Charmap{charmap.CodePage860} }
	CP861_DOSIcelandic       = func() Encoding { return NewCharmap2(cp861_DOSIcelandic) }
	CP862_DOSHebrew          = func() Encoding { return Charmap{charmap.CodePage862} }
	CP863_DOSCanadaF         = func() Encoding { return Charmap{charmap.CodePage863} }
	CP864_DOSArabic          = func() Encoding { return NewCharmap2(cp864_DOSArabic) }
	CP865_DOSNordic          = func() Encoding { return Charmap{charmap.CodePage863} }
	CP866_DOSCyrillicRussian = func() Encoding { return Charmap{charmap.CodePage866} }
	CP869_DOSGreek2          = func() Encoding { return NewCharmap2(cp869_DOSGreek2) }
	CP874_DOSThai            = func() Encoding { return NewCharmap2(cp874) }
	CP875_IBMGreek           = func() Encoding { return NewCharmap2(cp875_IBMGreek) }
	CP932                    = func() Encoding { return Charmap{japanese.ShiftJIS} }
	CP936                    = func() Encoding { return Charmap{simplifiedchinese.GBK} }
	CP949                    = func() Encoding { return Charmap{korean.EUCKR} }
	CP950                    = func() Encoding { return Charmap{traditionalchinese.Big5} }
	CP1026_IBMLatin5Turkish  = func() Encoding { return NewCharmap2(cp1026_IBMLatin5Turkish) }
	CP1250                   = func() Encoding { return Charmap{charmap.Windows1250} }
	CP1251                   = func() Encoding { return Charmap{charmap.Windows1251} }
	CP1252                   = func() Encoding { return Charmap{charmap.Windows1252} }
	EBCDIC037                = func() Encoding { return EBCDIC{ebcdic.EBCDIC037} }
	EBCDIC1140               = func() Encoding { return EBCDIC{ebcdic.EBCDIC1140} }
	EBCDIC1141               = func() Encoding { return EBCDIC{ebcdic.EBCDIC1141} }
	EBCDIC1148               = func() Encoding { return EBCDIC{ebcdic.EBCDIC1148} }
	EBCDIC273                = func() Encoding { return EBCDIC{ebcdic.EBCDIC273} }
	EBCDIC500                = func() Encoding { return EBCDIC{ebcdic.EBCDIC500} }
)

var (
	defaultEncodingMu = new(sync.Mutex)
	defaultEncoding   = CP437_DOSLatinUS()
)

// DefaultEncoding function returns the default OEM encoding.
func DefaultEncoding() Encoding {
	return defaultEncoding
}

// WithDefaultEncoding function sets the default OEM encoding for
// all top-level encode/decode operations.
func WithDefaultEncoding(e Encoding) {
	defaultEncodingMu.Lock()
	defaultEncoding = e
	defaultEncodingMu.Unlock()
}

// Encode function encodes the string to the default OEM encoding.
func Encode(s string) ([]byte, error) {
	return defaultEncoding.Encode(s)
}

// Decode function decodes the string from the default OEM encoding.
func Decode(b []byte) (string, error) {
	return defaultEncoding.Decode(b)
}

type ctxKey struct{}

func WithContext(ctx context.Context, e Encoding) context.Context {
	return context.WithValue(ctx, ctxKey{}, e)
}

func FromContext(ctx context.Context) Encoding {
	e, ok := ctx.Value(ctxKey{}).(Encoding)
	if !ok {
		return defaultEncoding
	}
	return e
}
