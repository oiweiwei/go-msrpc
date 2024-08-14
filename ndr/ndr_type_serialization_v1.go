package ndr

import (
	"context"
)

type CommonHeaderV1 struct {
	Version            uint8
	Endianness         uint8
	CommonHeaderLength uint16
	Filler             []byte
}

func (c *CommonHeaderV1) UnmarshalNDR(ctx context.Context, r Reader) error {
	r.ReadData(&c.Version)
	r.ReadData(&c.Endianness)
	r.ReadData(&c.CommonHeaderLength)

	c.Filler = make([]byte, 4)
	for i1 := 0; i1 < len(c.Filler); i1++ {
		r.ReadData(&c.Filler[i1])
	}

	return r.Err()
}

func (c *CommonHeaderV1) MarshalNDR(ctx context.Context, w Writer) error {
	w.WriteData(c.Version)
	w.WriteData(c.Endianness)
	w.WriteData(c.CommonHeaderLength)

	for i1 := 0; i1 < len(c.Filler); i1++ {
		if i1 >= 4 {
			break
		}
		w.WriteData(c.Filler[i1])
	}
	for i1 := len(c.Filler); i1 < 4; i1++ {
		w.WriteData(byte(0))
	}

	return w.Err()
}

type PrivateHeaderV1 struct {
	ObjectBufferLength uint32
	Filler             []byte
}

func (p *PrivateHeaderV1) UnmarshalNDR(ctx context.Context, r Reader) error {
	r.ReadData(&p.ObjectBufferLength)

	p.Filler = make([]byte, 4)
	for i1 := 0; i1 < len(p.Filler); i1++ {
		r.ReadData(&p.Filler[i1])
	}

	return r.Err()
}

func (p *PrivateHeaderV1) MarshalNDR(ctx context.Context, w Writer) error {
	w.WriteData(p.ObjectBufferLength)

	for i1 := 0; i1 < len(p.Filler); i1++ {
		if i1 >= 4 {
			break
		}
		w.WriteData(p.Filler[i1])
	}
	for i1 := len(p.Filler); i1 < 4; i1++ {
		w.WriteData(byte(0))
	}

	return w.Err()
}

func UnmarshalWithTypeSerializationV1(b []byte, d Unmarshaler, opts ...any) error {

	ctx := context.Background()

	r := NDR20(b, d, Opaque)

	var c CommonHeaderV1
	err := c.UnmarshalNDR(ctx, r)
	if err != nil {
		return err
	}

	var p PrivateHeaderV1
	err = p.UnmarshalNDR(ctx, r)
	if err != nil {
		return err
	}

	opts = append(opts, DefaultDataRepresentation|DataRepresentation(c.Endianness))

	return NDR20(r.Bytes()[:p.ObjectBufferLength], opts...).Unmarshal(ctx, d)
}

func MarshalWithTypeSerializationV1(d Marshaler, opts ...any) ([]byte, error) {

	ctx := context.Background()

	c := CommonHeaderV1{
		Version:            0x01,
		Endianness:         uint8(ByteOrderLittleEndian),
		CommonHeaderLength: 0x08,
		Filler:             []byte{0xCC, 0xCC, 0xCC, 0xCC},
	}

	for _, opt := range opts {
		switch v := opt.(type) {
		case DataRepresentation:
			c.Endianness = uint8(ByteOrderMask & uint32(v))
		}
	}

	opts = append(opts, DefaultDataRepresentation|DataRepresentation(c.Endianness))

	b, err := NDR20(nil, opts...).Marshal(ctx, d)
	if err != nil {
		return nil, err
	}

	// add padding to 8 bytes.
	if pad := len(b) % 8; pad != 0 {
		b = append(b, make([]byte, 8-pad)...)
	}

	p := PrivateHeaderV1{
		ObjectBufferLength: uint32(len(b)),
	}

	r := NDR20(nil, Opaque)

	if err := c.MarshalNDR(ctx, r); err != nil {
		return nil, err
	}

	if err := p.MarshalNDR(ctx, r); err != nil {
		return nil, err
	}

	return append(r.Bytes(), b...), nil
}

func UnmarshalerPointer(v Unmarshaler) Unmarshaler {
	return UnmarshalNDRFunc(func(ctx context.Context, r Reader) error {
		return r.ReadPointer(&v, nil, v)
	})
}

func MarshalerPointer(v Marshaler) Marshaler {
	return MarshalNDRFunc(func(ctx context.Context, w Writer) error {
		return w.WritePointer(&v, v)
	})
}
