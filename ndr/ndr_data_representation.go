package ndr

import (
	"encoding/binary"

	"github.com/oiweiwei/go-msrpc/ndr/math"
)

// The NDR data representation.
type DataRepresentation uint32

var (
	// The default data representation (ASCII, FloatIEEE, LittleEndian).
	DefaultDataRepresentation = DataRepresentation(CharASCII | FloatingPointIEEE | ByteOrderLittleEndian)
)

// ByteOrder function returns the byte order for the data
// representation.
func (d DataRepresentation) ByteOrder() binary.ByteOrder {
	switch (uint32)(d) & ByteOrderMask {
	case ByteOrderLittleEndian:
		return binary.LittleEndian
	}
	return binary.BigEndian
}

// FloatFormat function returns the floating-point format for
// the data representation.
func (d DataRepresentation) FloatFormat() math.FloatFormat {
	switch (uint32)(d) & FloatingPointMask {
	case FloatingPointVAX:
		return math.Vax
	case FloatingPointCray:
		return math.Cray
	case FloatingPointIBM:
		return math.IBMHex
	}
	return math.IEEE
}

var (
	// IEEE floating point representation.
	FloatingPointIEEE uint32 = 0x00000000
	// VAX floating point representation.
	FloatingPointVAX uint32 = 0x00000100
	// Cray floating point representation.
	FloatingPointCray uint32 = 0x00000200
	// IBM floating point representation.
	FloatingPointIBM uint32 = 0x00000300
	// Mask.
	FloatingPointMask uint32 = 0x0000FF00
)

var (
	// ASCII character representation.
	CharASCII uint32 = 0x00000000
	// EBCDIC character representation.
	CharEBCDIC uint32 = 0x00000001
	// Mask.
	CharMask uint32 = 0x0000000F
)

var (
	// Big-endian data representation.
	ByteOrderBigEndian uint32 = 0x00000000
	// Little-endian data representation.
	ByteOrderLittleEndian uint32 = 0x00000010
	// Mask.
	ByteOrderMask uint32 = 0x000000F0
)
