// package ndr implements the NDR encoding.
package ndr

import (
	"context"
	"io"
)

// The NDR operation.
type Operation interface {
	// The operation number.
	OpNum() int
	// The operation name.
	OpName() string
	// The function to marshal the request parameters.
	MarshalNDRRequest(context.Context, Writer) error
	// The function to unmarshal the request parameters.
	UnmarshalNDRRequest(context.Context, Reader) error
	// The function to marshal the response fields.
	MarshalNDRResponse(context.Context, Writer) error
	// The function to unmarshal the response fields.
	UnmarshalNDRResponse(context.Context, Reader) error
}

// Marshaler interface ...
type Marshaler interface {
	MarshalNDR(context.Context, Writer) error
}

// Unmarshaler interface ...
type Unmarshaler interface {
	UnmarshalNDR(context.Context, Reader) error
}

type AfterUnmarshalNDR interface {
	AfterUnmarshalNDR(context.Context) error
}

// Error(er) interface.
type Error interface {
	// Err function must return the captured error.
	Err() error
	// SetError must set and returns the error.
	SetErr(error) error
}

// Buffer interface.
type Buffer interface {
	// EOF function must returns `true` if there is no more
	// data in the buffer.
	EOF() bool
	// Len function must returns the number of bytes available
	// for read or number of written bytes.
	Len() int
	// Offset function must return current buffer position.
	Offset() int
	// Bytes function must return the buffer bytes.
	Bytes() []byte
	// WithBytes function assigns the buffer to the parameter
	// bytes.
	WithBytes([]byte) NDR
}

// NDR interface ...
type NDR interface {
	// Writer.
	Writer
	// Reader.
	Reader
}

type Reader interface {

	// Reader is a generic I/O writer interface.
	io.Reader

	// Error to return an error.
	Error

	// Buffer.
	Buffer

	// Marshal function performs the NDR marshaling with pointer
	// resolution.
	Unmarshal(context.Context, Unmarshaler) error

	// ReadAlign function performs alignment of the output buffer.
	ReadAlign(int) error

	// ReadUnionAlign function performs alignment of the union
	// buffer.
	ReadUnionAlign(int) error

	// here go ReadX interface methods required to
	// fullfill unmarshaling requirements.

	// ReadData will align and read the type as-is.
	// If type has alignment, it should align the data in the
	// UnmarshalNDR method, if it's the primitive type,
	// primitive type alignment will be used.
	ReadData(any) error

	// ReadSize function reads the maximum count, offset,
	// and actual count for arrays, strings, pipes.
	ReadSize(*uint64) error

	// ReadSwitch function reads the switch value for the union.
	ReadSwitch(any) error

	// ReadEnum function reads the enumeration.
	ReadEnum(any) error

	// ReadPointer function reads the implementation-specific pointer.
	// It receives the setter function in case if pointer is not null.
	ReadPointer(Pointer, func(any), ...Unmarshaler) error

	// ReadDeferred function reads all the deferred pointers.
	ReadDeferred() error
}

// Writer interface ...
type Writer interface {

	// Writer is a generic I/O writer interface.
	io.Writer

	// Error to return an error.
	Error

	// Buffer.
	Buffer

	// Marshal function performs the NDR marshaling with pointer
	// resolution.
	Marshal(context.Context, Marshaler) ([]byte, error)

	// WriteAlign function performs alignment of the output buffer.
	WriteAlign(int) error

	// WriteUnionAlign function performs alignment of the union
	// output buffer.
	WriteUnionAlign(int) error

	// here go WriteX interface methods required to
	// fullfill marshaling requirements.

	// WriteData will write and align the type as-is.
	// If type implements align method, it will be used, if it's the
	// primitive type, primitive type alignment will be used.
	WriteData(any) error

	// WriteSize function writes the maximum count, offset,
	// and actual count for arrays, strings, pipes.
	WriteSize(uint64) error

	// WriteSwitch function writes the union switch.
	WriteSwitch(any) error

	// WriteEnum function writes the enumeration.
	WriteEnum(any) error

	// WritePointer function writes the implementation-specific pointer.
	// Optional data parameter will be used when the full pointer is used.
	WritePointer(Pointer, ...Marshaler) error

	// WriteDeferred function writes all the deferred pointers.
	WriteDeferred() error
}

// MarshalNDRFunc function ...
type MarshalNDRFunc func(context.Context, Writer) error

// MarshalNDR function implements the Marshaler interface.
func (f MarshalNDRFunc) MarshalNDR(ctx context.Context, w Writer) error {
	return f(ctx, w)
}

// UnmarshalNDRFunc function ...
type UnmarshalNDRFunc func(context.Context, Reader) error

// UnmarshalNDR function implements the Unmarshaler interface.
func (f UnmarshalNDRFunc) UnmarshalNDR(ctx context.Context, w Reader) error {
	return f(ctx, w)
}

var (
	// ZeroString represents the empty null-termination string.
	ZeroString = string([]byte{0})
)

// SizeInfoCtx is a context key type for the size information.
type SizeInfoCtx struct{}

var (
	// SizeInfoCtx is a context type for the size information.
	SizeInfo = SizeInfoCtx{}
)

// Pointer interface represents the generic pointer.
type Pointer interface{}

// Int3264 type ...
type Int3264 int64

// Uint3264 type ...
type Uint3264 uint64

type Enum any
