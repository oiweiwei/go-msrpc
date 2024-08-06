package pac

import (
	"context"

	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
	"github.com/oiweiwei/go-msrpc/text/encoding/utf16le"
)

func (o *UPNDNSInfo) AfterUnmarshalNDR(ctx context.Context) error {

	oft := 12

	if o.Flags&0x00000002 != 0 {
		r := ndr.NDR20(o.Raw, ndr.Opaque)
		r.ReadData(&o.SAMNameLength)
		r.ReadData(&o.SAMNameOffset)
		r.ReadData(&o.SIDLength)
		r.ReadData(&o.SIDOffset)
		oft += 8
	}

	if o.UPNLength != 0 {
		oft := int(o.UPNOffset) - oft
		s, err := utf16le.Decode(o.Raw[oft : oft+int(o.UPNLength)])
		if err != nil {
			return err
		}
		o.UPN = s
	}

	if o.DNSDomainNameLength != 0 {
		oft := int(o.DNSDomainNameOffset) - oft
		s, err := utf16le.Decode(o.Raw[oft : oft+int(o.DNSDomainNameLength)])
		if err != nil {
			return err
		}
		o.DNSDomainName = s
	}

	if o.SAMNameLength != 0 {
		oft := int(o.SAMNameOffset) - oft
		s, err := utf16le.Decode(o.Raw[oft : oft+int(o.SAMNameLength)])
		if err != nil {
			return err
		}
		o.SAMName = s
	}

	if o.SIDLength != 0 {
		oft := int(o.SIDOffset) - oft
		o.SID = &dtyp.SID{}
		if err := o.SID.DecodeBinary(o.Raw[oft : oft+int(o.SIDLength)]); err != nil {
			return err
		}
	}

	return nil
}
