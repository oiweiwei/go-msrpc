package crypto

import "context"

type Cipher interface {
	Size(context.Context, bool) int
	Wrap(context.Context, uint64, [][]byte, [][]byte) ([]byte, error)
	Unwrap(context.Context, uint64, [][]byte, [][]byte, []byte) ([]byte, error)
	MakeSignature(context.Context, uint64, [][]byte) ([]byte, error)
}

type etype int

var (
	EtypeAESSHA2 etype = 1
	EtypeRC4MD5  etype = 2
)

func NewCipher(ctx context.Context, etype etype, key []byte, isClient bool) Cipher {

	switch etype {
	case EtypeAESSHA2:
		return &AESSHA2{isClient: isClient, key: key}
	case EtypeRC4MD5:
		return &RC4MD5{isClient: isClient, key: key}
	}

	return nil
}

func MakeSeqNum(ctx context.Context, seqNum uint64, client bool) []byte {

	out := make([]byte, 8)

	out[0] = uint8(seqNum >> 24)
	out[1] = uint8(seqNum >> 16)
	out[2] = uint8(seqNum >> 8)
	out[3] = uint8(seqNum)
	out[4] = uint8(seqNum >> 56)
	out[5] = uint8(seqNum >> 48)
	out[6] = uint8(seqNum >> 40)
	out[7] = uint8(seqNum >> 32)

	if client {
		out[4] |= 0x80
	}

	return out
}
