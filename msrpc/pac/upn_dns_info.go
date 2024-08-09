package pac

import (
	"bytes"
	"context"

	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
	"github.com/oiweiwei/go-msrpc/text/encoding/utf16le"
)

func (o *UPNDNSInfo) Unmarshal(b []byte) error {

	r := ndr.NDR20(b, ndr.Opaque)

	if err := r.Unmarshal(context.Background(), o); err != nil {
		return err
	}

	if o.Flags&0x00000002 != 0 {

		r = ndr.NDR20(o.Raw, ndr.Opaque)

		if err := r.ReadData(&o.SAMNameLength); err != nil {
			return err
		}

		if err := r.ReadData(&o.SAMNameOffset); err != nil {
			return err
		}

		if err := r.ReadData(&o.SIDLength); err != nil {
			return err
		}

		if err := r.ReadData(&o.SIDOffset); err != nil {
			return err
		}

		o.Raw = r.Bytes()
	}

	var err error

	if o.UPNLength != 0 {
		if o.UPN, err = utf16le.Decode(b[o.UPNOffset : o.UPNOffset+o.UPNLength]); err != nil {
			return err
		}
	}

	if o.DNSDomainNameLength != 0 {
		if o.DNSDomainName, err = utf16le.Decode(b[o.DNSDomainNameOffset : o.DNSDomainNameOffset+o.DNSDomainNameLength]); err != nil {
			return err
		}
	}

	if o.SAMNameLength != 0 {
		if o.SAMName, err = utf16le.Decode(b[o.SAMNameOffset : o.SAMNameOffset+o.SAMNameLength]); err != nil {
			return err
		}
	}

	if o.SIDLength != 0 {
		o.SID = &dtyp.SID{}
		if err := o.SID.DecodeBinary(b[o.SIDOffset : o.SIDOffset+o.SIDLength]); err != nil {
			return err
		}
	}

	return nil
}

func (o *UPNDNSInfo) Marshal() ([]byte, error) {

	if o == nil {
		o = &UPNDNSInfo{}
	}

	buf, oft := new(bytes.Buffer), uint16(12) /* upn, dnsname, flags */

	if o.SAMName != "" || o.SID != nil {
		o.Flags |= 0x00000002
		oft += 8 /* sid and sam offset/length */
		buf.Write(make([]byte, 8))
	}

	if o.UPN != "" {
		b, err := utf16le.Encode(o.UPN)
		if err != nil {
			return nil, err
		}
		o.UPNLength, o.UPNOffset = uint16(len(b)), oft
		buf.Write(b)
	}

	oft += o.UPNLength

	if o.DNSDomainName != "" {
		b, err := utf16le.Encode(o.DNSDomainName)
		if err != nil {
			return nil, err
		}
		o.DNSDomainNameLength, o.DNSDomainNameOffset = uint16(len(b)), oft
		buf.Write(b)
	}

	oft += o.DNSDomainNameLength

	if o.SAMName != "" {
		b, err := utf16le.Encode(o.SAMName)
		if err != nil {
			return nil, err
		}
		o.SAMNameLength, o.SAMNameOffset = uint16(len(b)), oft
		buf.Write(b)
	}

	oft += o.SAMNameLength

	if o.SID != nil {
		b, err := o.SID.Bytes()
		if err != nil {
			return nil, err
		}
		o.SIDLength, o.SIDOffset = uint16(len(b)), oft
		buf.Write(b)
	}

	o.Raw = buf.Bytes()

	if o.Flags&0x00000002 != 0 {

		r := ndr.NDR20(nil, ndr.Opaque)

		if err := r.WriteData(o.SAMNameLength); err != nil {
			return nil, err
		}

		if err := r.WriteData(o.SAMNameOffset); err != nil {
			return nil, err
		}

		if err := r.WriteData(o.SIDLength); err != nil {
			return nil, err
		}

		if err := r.WriteData(o.SIDOffset); err != nil {
			return nil, err
		}

		// copy sam/sid data at the beginning of the raw.
		copy(o.Raw, r.Bytes())
	}

	b, err := ndr.Marshal(o, ndr.Opaque)
	if err != nil {
		return nil, err
	}

	return b, nil
}
