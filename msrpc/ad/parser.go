package ad

import (
	"encoding/binary"
	"time"

	"github.com/oiweiwei/go-msrpc/msrpc/drsr/drsuapi/v4"
	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	"github.com/oiweiwei/go-msrpc/ndr"
	"github.com/oiweiwei/go-msrpc/text/encoding/utf16le"
)

type StringUnicode struct{}

func (StringUnicode) Size() int { return -1 }

func (StringUnicode) Convert(b []byte, _ ...any) (any, error) {
	ret, err := utf16le.Decode(b)
	return ret, err
}

type Enumeration struct{}

func (Enumeration) Size() int { return 4 }

func (Enumeration) Convert(b []byte, _ ...any) (any, error) {
	return binary.LittleEndian.Uint32(b), nil
}

type GeneralizedTime struct{}

func (GeneralizedTime) Size() int { return 8 }

func (GeneralizedTime) Convert(b []byte, _ ...any) (any, error) {
	seconds := binary.LittleEndian.Uint64(b)
	return time.Unix(int64(seconds)-11644473600, 0).UTC(), nil
}

type Interval struct{}

func (Interval) Size() int { return 8 }

func (Interval) Convert(b []byte, _ ...any) (any, error) {
	return binary.LittleEndian.Uint64(b), nil
}

type Boolean struct{}

func (Boolean) Size() int { return 4 }

func (Boolean) Convert(b []byte, _ ...any) (any, error) {
	return binary.LittleEndian.Uint32(b) != 0, nil
}

type DSDN struct{}

func (DSDN) Size() int { return -1 }

func (DSDN) Convert(b []byte, _ ...any) (any, error) {
	dn := drsuapi.DSName{}
	if err := ndr.Unmarshal(b, &dn, ndr.Opaque); err != nil {
		return nil, err
	}
	return &dn, nil
}

type ObjectIdentifier struct{}

func (ObjectIdentifier) Size() int { return 4 }

func (ObjectIdentifier) Convert(b []byte, opts ...any) (any, error) {

	var prefixes drsuapi.PrefixTable

	for _, opt := range opts {
		switch opt := opt.(type) {
		case drsuapi.PrefixTable:
			prefixes = opt
		}
	}

	ret, err := prefixes.AttributeToOID(binary.LittleEndian.Uint32(b))
	return ret, err
}

type SID struct{}

func (SID) Size() int { return 28 }

func (SID) Convert(b []byte, _ ...any) (any, error) {
	sid := dtyp.SID{}
	if err := ndr.Unmarshal(b, &sid, ndr.Opaque); err != nil {
		return nil, err
	}
	return &sid, nil
}

type SecurityDescriptor struct{}

func (SecurityDescriptor) Size() int { return -1 }

func (SecurityDescriptor) Convert(b []byte, _ ...any) (any, error) {
	sd := &dtyp.SecurityDescriptor{}

	if err := sd.Parse(b); err != nil {
		return nil, err
	}

	return sd, nil
}

type Raw struct{}

func (Raw) Size() int { return -1 }

func (Raw) Convert(b []byte, _ ...any) (any, error) {
	return b, nil
}
