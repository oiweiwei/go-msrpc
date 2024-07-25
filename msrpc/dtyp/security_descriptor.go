package dtyp

import (
	"context"

	"github.com/oiweiwei/go-msrpc/ndr"
)

func (o *SecurityDescriptor) Parse(b []byte) error {

	r := ndr.NDR20(b, ndr.Opaque)

	r.ReadData(&o.Revision)
	r.ReadData(&o.SBZ1)
	r.ReadData(&o.Control)

	if r.Err() != nil {
		return r.Err()
	}

	if o.Control&SelfRelative == 0 {
		return ndr.NDR20(b, ndr.Opaque).Unmarshal(context.Background(), o)
	}

	r.ReadData(&o.OffsetOwner)
	r.ReadData(&o.OffsetGroup)
	r.ReadData(&o.OffsetSACL)
	r.ReadData(&o.OffsetDACL)

	if r.Err() != nil {
		return r.Err()
	}

	if o.OffsetOwner != 0 {
		if o.Owner == nil {
			o.Owner = &SID{}
		}
		if err := r.WithBytes(b[o.OffsetOwner:]).Unmarshal(context.Background(), o.Owner); err != nil {
			return err
		}
	}

	if o.OffsetGroup != 0 {
		if o.Group == nil {
			o.Group = &SID{}
		}
		if err := r.WithBytes(b[o.OffsetGroup:]).Unmarshal(context.Background(), o.Group); err != nil {
			return err
		}
	}

	if o.OffsetSACL != 0 {
		if o.SACL == nil {
			o.SACL = &ACL{}
		}
		if err := r.WithBytes(b[o.OffsetSACL:]).Unmarshal(context.Background(), o.SACL); err != nil {
			return err
		}
	}

	if o.OffsetDACL != 0 {
		if o.DACL == nil {
			o.DACL = &ACL{}
		}
		if err := r.WithBytes(b[o.OffsetDACL:]).Unmarshal(context.Background(), o.DACL); err != nil {
			return err
		}
	}

	return nil
}
