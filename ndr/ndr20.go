package ndr

// ndr20.go module contains the implementation of the Network
// Data Representation 2.0 (NDR2.0).

import (
	"context"
	"fmt"
)

type opaque struct{}

// Opaque is an NDR option that is used to indicate that
// implicit NDR data, like size labels, non-encapsulated switches,
// pointer values must be omitted. Any pointer deferred values
// must be encoded immediately.
var Opaque opaque

type debug struct{}

var Debug debug

type noLayout struct{}

var NoLayout noLayout

// NDR20 function returns the NDR2.0 Marshaler/Unmarshaler.
func NDR20(buf []byte, opts ...any) NDR {

	ndr := &ndr20{
		ptrs: make(map[uint64]Pointer),
		drep: DefaultDataRepresentation,
	}

	var chnk ChunkedBuffer

	for i := range opts {
		switch o := opts[i].(type) {
		case DataRepresentation:
			ndr.drep = o
		case opaque:
			ndr.opaque = true
		case noLayout:
			ndr.noLayout = true
		case debug:
			ndr.debug = true
		case ChunkedBuffer:
			chnk = o
		}
	}
	if chnk == nil {
		chnk = NewChunk(buf, ndr.drep)
	}
	ndr.buf = NewAlignBuffer(chnk)
	return ndr
}

// WithBytes function sets the current buffer bytes to value `b`.
func (w *ndr20) WithBytes(b []byte) NDR {
	return &ndr20{
		drep:     w.drep,
		buf:      NewAlignBuffer(NewChunk(b, w.drep)),
		ptrs:     make(map[uint64]Pointer),
		opaque:   w.opaque,
		noLayout: w.noLayout,
		noop:     w.noLayout,
		err:      w.err,
	}
}

// Marshal function marshals the data `d` using NDR2.0 format.
func Marshal(d Marshaler, opts ...any) ([]byte, error) {
	return NDR20(nil, opts...).Marshal(context.Background(), d)
}

// MarshalResponse function marshal the operation `d` response using NDR2.0 format.
func MarshalResponse(d Operation, opts ...any) ([]byte, error) {
	w := NDR20(nil, opts...)
	if err := d.MarshalNDRResponse(context.Background(), w); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

// Unmarshal function unmarshals the bytes `b` to the data `d`
// using NDR2.0 format.
func Unmarshal(b []byte, d Unmarshaler, opts ...any) error {
	return NDR20(b, opts...).Unmarshal(context.Background(), d)
}

// The NDR2.0 implementation.
type ndr20 struct {
	// The data representation.
	drep DataRepresentation
	// The first captured error.
	err error
	// The buffer to read/write.
	buf AlignBuffer
	// Temporary location to hold the data for read/write.
	put [8]byte
	// The list of deferred write pointer.
	wdeferred []Marshaler
	// The list of deferred read pointers.
	rdeferred []Unmarshaler
	// The pointers map.
	ptrs map[uint64]Pointer
	// The flag that indicates whether to include NDR-related
	// labels into the marshaled/unmarshaled output.
	opaque, debug, noLayout, noop bool
}

// Err function returns the NDR error.
func (w *ndr20) Err() error {
	return w.err
}

// EOF function indicates the end of current buffer.
func (w *ndr20) EOF() bool { return w.buf.EOF() }

// Offset function returns the buffer offset.
func (w *ndr20) Offset() int { return w.buf.Pos() }

func (w *ndr20) wDeferred() []Marshaler {
	var wdeferred []Marshaler
	wdeferred, w.wdeferred = w.wdeferred, nil
	return wdeferred
}

func (w *ndr20) rDeferred() []Unmarshaler {
	var rdeferred []Unmarshaler
	rdeferred, w.rdeferred = w.rdeferred, nil
	return rdeferred
}

// makeAlign function computes the alignment for the size.
// 6 (1+5), 7 (2+5), 9 (4+5), 13 (8+5) values are used to
// recover the size of the data if it was not an NDR pointer/array.
// (see Opaque).
func (w *ndr20) makeAlign(sz int) int {
	switch sz {
	case 5, 6, 7, 9, 13:
		if w.opaque {
			return sz - 5
		}
		return 5
	}
	return sz
}

// WriteAlign function writes the alignment required for the data
// of size `sz`.
func (w *ndr20) WriteAlign(sz int) error {

	if w.err != nil || sz < 2 {
		return w.err
	}

	if sz = w.makeAlign(sz); sz == 5 {
		sz = 4
	}

	return w.buf.FillMod(sz)
}

func (w *ndr20) WriteUnionAlign(sz int) error {
	// not used.
	return w.err
}

// ReadAlign function read the alignment required for the data
// of size `sz`.
func (w *ndr20) ReadAlign(sz int) error {

	if w.err != nil || sz < 2 {
		return w.err
	}

	if sz = w.makeAlign(sz); sz == 5 {
		sz = 4
	}

	return w.buf.SkipMod(sz)
}

func (w *ndr20) ReadUnionAlign(sz int) error {
	// not used.
	return w.err
}

// ReadSize function reads the size label information (length, size, offset)
// from the buffer.
func (w *ndr20) ReadSize(sz *uint64) error {

	if w.err != nil || w.opaque {
		*sz = 0 // clear size_info.
		return w.err
	}

	sz20 := uint32(0)

	if err := w.ReadData(&sz20); err != nil {
		return err
	}

	*sz = uint64(sz20)
	return nil
}

// ReadSwitch function reads the non-encapsulated NDR switch
// value from the buffer.
func (w *ndr20) ReadSwitch(sw any) error {

	if w.err != nil || w.opaque {
		return w.err
	}

	if enum, ok := sw.(EnumWrapper); ok {
		return w.ReadEnum(enum.Value())
	}

	return w.ReadData(sw)
}

// ReadEnum function reads the enumeration value from the buffer.
func (w *ndr20) ReadEnum(enum any) error {

	if w.err != nil {
		return w.err
	}

	return w.ReadData(enum)
}

// Len function returns the current buffer length.
func (w *ndr20) Len() int { return w.buf.Len() }

// Bytes function returns the available bytes in the buffer.
func (w *ndr20) Bytes() []byte { return w.buf.Bytes() }

// WriteSize function writes the NDR size information (length, size, offset)
// to the buffer.
func (w *ndr20) WriteSize(sz uint64) error {

	if w.err != nil || w.opaque {
		return w.err
	}

	return w.WriteData(uint32(sz))
}

// WriteSwitch function writes the non-encapsulated NDR
// switch value to the buffer.
func (w *ndr20) WriteSwitch(sw any) error {

	if w.err != nil || w.opaque {
		return w.err
	}

	if enum, ok := sw.(EnumWrapper); ok {
		return w.WriteEnum(enum.Value())
	}

	return w.WriteData(sw)
}

// WriteEnum function writes the enumeration value to the buffer.
func (w *ndr20) WriteEnum(enum any) error {

	if w.err != nil {
		return w.err
	}

	return w.WriteData(enum)
}

// SetErr function sets and returns the error `err`. Any method call
// will not be executed and the first error captured will be returned.
func (w *ndr20) SetErr(err error) error {

	if w.err != nil {
		return w.err
	}

	if err != nil {
		w.err = err
	}

	return w.err
}

// DataSize function returns the data size for the given value `d`.
func DataSize(d any) int {

	switch d.(type) {
	case bool, *bool, uint8, *uint8, int8, *int8:
		return 1
	case uint16, *uint16, int16, *int16:
		return 2
	case uint32, *uint32, int32, *int32, float32, *float32, Uint3264, *Uint3264, Int3264, *Int3264:
		return 4
	case uint64, *uint64, int64, *int64, float64, *float64:
		return 8
	case DataRepresentation, *DataRepresentation:
		return 0
	}

	panic(fmt.Sprintf("unsupported type %T", d))
}

// ReadData function reads the data `d` from the buffer.
func (w *ndr20) ReadData(d any) error {

	if w.err != nil {
		return w.err
	}

	sz, buf := DataSize(d), w.put[:]
	if sz > 0 {
		if err := w.buf.SkipMod(sz); err != nil {
			return w.SetErr(err)
		}
		buf = buf[:sz]
	}

	if drep, ok := d.(*DataRepresentation); ok {
		return w.SetErr(w.buf.ReadRepresentation(drep))
	}

	if _, err := w.buf.Read(buf); err != nil {
		return w.SetErr(err)
	}

	order, float := w.buf.Order(), w.buf.Float()

	switch d := d.(type) {
	case *bool:
		*d = buf[0] > 0
	case *uint8:
		*d = buf[0]
	case *int8:
		*d = int8(buf[0])
	case *uint16:
		*d = order.Uint16(buf)
	case *int16:
		*d = int16(order.Uint16(buf))
	case *uint32:
		*d = order.Uint32(buf)
	case *int32:
		*d = int32(order.Uint32(buf))
	case *uint64:
		*d = order.Uint64(buf)
	case *int64:
		*d = int64(order.Uint64(buf))
	case *float32:
		*d = float.Float32frombits(order.Uint32(buf))
	case *float64:
		*d = float.Float64frombits(order.Uint64(buf))
	case *Int3264:
		*(*int64)(d) = int64(order.Uint32(buf))
	case *Uint3264:
		*(*uint64)(d) = uint64(order.Uint32(buf))
	default:
		return w.SetErr(fmt.Errorf("unsupported type %T", d))
	}

	return nil
}

// Write function implements the io.Writer interface.
func (w *ndr20) Write(b []byte) (n int, err error) {

	if w.err != nil {
		return 0, w.err
	}

	if n, err = w.buf.Write(b); err != nil {
		return n, w.SetErr(err)
	}

	return n, nil
}

// Read function implements the io.Reader interface.
func (w *ndr20) Read(p []byte) (n int, err error) {

	if w.err != nil {
		return 0, w.err
	}

	if n, err = w.buf.Read(p); err != nil {
		return n, w.SetErr(err)
	}

	return n, nil
}

// WriteData function writes the data type into buffer.
func (w *ndr20) WriteData(d any) error {

	if w.err != nil {
		return w.err
	}

	sz, buf := DataSize(d), w.put[:]
	if sz > 0 {
		if err := w.buf.FillMod(sz); err != nil {
			return w.SetErr(err)
		}
		buf = buf[:sz]
	}

	if drep, ok := d.(DataRepresentation); ok {
		return w.SetErr(w.buf.WriteRepresentation(drep))
	}

	order, float := w.buf.Order(), w.buf.Float()

	switch d := d.(type) {
	case bool:
		if d {
			buf[0] = uint8(1)
		} else {
			buf[0] = uint8(0)
		}
	case uint8:
		buf[0] = d
	case int8:
		buf[0] = uint8(d)
	case uint16:
		order.PutUint16(buf, d)
	case int16:
		order.PutUint16(buf, uint16(d))
	case uint32:
		order.PutUint32(buf, d)
	case int32:
		order.PutUint32(buf, uint32(d))
	case uint64:
		order.PutUint64(buf, d)
	case int64:
		order.PutUint64(buf, uint64(d))
	case float32:
		order.PutUint32(buf, float.Float32bits(d))
	case float64:
		order.PutUint64(buf, float.Float64bits(d))
	case Int3264:
		order.PutUint32(buf, uint32(d))
	case Uint3264:
		order.PutUint32(buf, uint32(d))
	default:
		return w.SetErr(fmt.Errorf("unsupported type %T", d))
	}

	if _, err := w.buf.Write(buf); err != nil {
		return w.SetErr(err)
	}

	return nil
}

// ReadPointer function reads the pointer value and defers the actual data read
// until the ReadDeferred is called. The `setter` value is used in case of the
// pointer aliasing.
func (w *ndr20) ReadPointer(ptr Pointer, setter func(interface{}), mrs ...Unmarshaler) error {

	if w.err != nil {
		return w.err
	}

	if w.opaque {
		for _, mr := range mrs {
			// start new execution context for the marshaler.
			if err := mr.UnmarshalNDR(context.Background(), w); err != nil {
				return w.SetErr(err)
			}
		}
		return nil
	}

	var pptr uint32

	if err := w.ReadData(&pptr); err != nil {
		return w.SetErr(err)
	}

	if pptr == 0 {
		return nil
	}

	if ptr, ok := w.ptrs[uint64(pptr)]; ok {
		setter(interface{}(ptr))
		return nil
	}

	w.ptrs[uint64(pptr)], w.rdeferred = ptr, append(w.rdeferred, mrs...)
	return nil
}

// WritePointer function writes the pointer to the data and defers the
// actual data write until the WriteDeferred is called.
func (w *ndr20) WritePointer(ptr Pointer, mrs ...Marshaler) error {

	if w.err != nil {
		return w.err
	}

	if w.opaque {
		for _, mr := range mrs {
			// start new execution context for the marshaler.
			if _, err := w.Marshal(context.Background(), mr); err != nil {
				return w.SetErr(err)
			}
		}
		return nil
	}

	if ptr == nil {
		return w.SetErr(w.WriteData(uint32(0)))
	}

	if err := w.WriteData(uint32(w.buf.Pos() + 1)); err != nil {
		return w.SetErr(err)
	}

	w.wdeferred = append(w.wdeferred, mrs...)

	return nil
}

// WriteDeferred function writes the deferred pointer values.
func (w *ndr20) WriteDeferred() error {

	if w.err != nil {
		return w.err
	}

	for _, deferred := range w.wDeferred() {
		// start new execution context for the marshaler.
		if _, err := w.Marshal(context.Background(), deferred); err != nil {
			return w.SetErr(err)
		}
	}

	return nil
}

// ReadDeferred function reads and decodes pointer values.
func (w *ndr20) ReadDeferred() error {

	if w.err != nil {
		return w.err
	}

	for _, deferred := range w.rDeferred() {
		// start new execution context for the marshaler.
		if err := w.Unmarshal(context.Background(), deferred); err != nil {
			return w.SetErr(err)
		}
	}

	return nil
}

// Marshal function marshals the `mrs` into the buffer and returns the
// array of bytes.
func (w *ndr20) Marshal(ctx context.Context, mrs Marshaler) ([]byte, error) {

	if w.err != nil || w.noop {
		return nil, w.err
	}

	if err := mrs.MarshalNDR(ctx, w); err != nil {
		return nil, w.SetErr(err)
	}

	if err := w.WriteDeferred(); err != nil {
		return nil, w.SetErr(err)
	}

	return w.buf.Bytes(), nil
}

// Unmarshal function unmarshals the `mrs` from the buffer.
func (w *ndr20) Unmarshal(ctx context.Context, mrs Unmarshaler) error {

	if w.err != nil || w.noop {
		return w.err
	}

	if err := mrs.UnmarshalNDR(ctx, w); err != nil {
		return w.SetErr(err)
	}

	if err := w.ReadDeferred(); err != nil {
		return w.SetErr(err)
	}

	if hook, ok := (any)(mrs).(AfterUnmarshalNDR); ok && hook != nil {
		if err := hook.AfterUnmarshalNDR(ctx); err != nil {
			return w.SetErr(err)
		}
	}

	return nil
}
