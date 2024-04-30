package ntlm

import (
	"context"

	"github.com/oiweiwei/go-msrpc/ssp/ntlm/internal"
)

type Signature struct {
	Version  uint32
	Checksum []byte
	SeqNum   uint32
}

func (s *Signature) Marshal(ctx context.Context) ([]byte, error) {

	e := internal.NewCodec(ctx, nil)

	// version.
	e.WriteData(ctx, s.Version)
	// checksum.
	e.WriteBytes(ctx, s.Checksum, 8)
	// seq_num.
	e.WriteData(ctx, s.SeqNum)

	return e.Bytes(ctx)
}

func (s *Signature) Unmarshal(ctx context.Context, b []byte) error {

	e := internal.NewCodec(ctx, b)

	// version.
	e.ReadData(ctx, &s.Version)
	// checksum.
	e.ReadBytes(ctx, &s.Checksum, 8)
	// seq_num.
	e.ReadData(ctx, &s.SeqNum)

	return e.ReadAll(ctx)
}
