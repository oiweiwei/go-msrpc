package winreg

import (
	"context"
	"strings"

	"github.com/oiweiwei/go-msrpc/ndr"
)

func (o *UnicodeString) BeforePreparePayload(ctx context.Context) error {

	// XXX: skip over setting parameters if they are already set, skip empty strings.
	if o.Length != 0 || o.MaximumLength != 0 || o.Buffer == "" {
		return nil
	}

	if !strings.HasSuffix(o.Buffer, ndr.ZeroString) {
		o.Buffer += ndr.ZeroString
		o.Length = uint16(ndr.UTF16Len(o.Buffer)) * 2
		o.MaximumLength = o.Length
	}

	return nil
}
