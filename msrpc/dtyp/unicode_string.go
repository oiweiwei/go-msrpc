package dtyp

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

	// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp/94a16bb6-c610-4cb9-8db6-26f15f560061
	// The string pointed to by the buffer member MUST NOT include a terminating null character.
	if strings.HasSuffix(o.Buffer, ndr.ZeroString) {
		o.Buffer = strings.TrimRight(o.Buffer, ndr.ZeroString)
		o.Length = uint16(ndr.UTF16Len(o.Buffer)) * 2
		o.MaximumLength = o.Length + 2
	}

	return nil
}
