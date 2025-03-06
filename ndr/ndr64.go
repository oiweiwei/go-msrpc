package ndr

import (
	"context"
	"fmt"
)

// NDR64 function returns the NDR64 Marshaler/Unmarshaler.
func NDR64(buf []byte, opts ...any) NDR {
	return &ndr64{
		ndr20: NDR20(buf, opts...).(*ndr20),
	}
}

type ndr64 struct {
	*ndr20
}

func (w *ndr64) WithBytes(b []byte) NDR {
	return &ndr64{
		ndr20: w.ndr20.WithBytes(b).(*ndr20),
	}
}

// Marshal64 function marshals the data `d` using NDR64 format.
func Marshal64(d Marshaler, opts ...any) ([]byte, error) {
	return NDR64(nil, opts...).Marshal(context.Background(), d)
}

// Unmarshal64 function unmarshals the bytes `b` to the data `d`
// using NDR64 format.
func Unmarshal64(b []byte, d Unmarshaler, opts ...any) error {
	return NDR64(b, opts...).Unmarshal(context.Background(), d)
}

// WriteAlign function writes the alignment required for the data
// of size `sz`.
func (w *ndr64) WriteAlign(sz int) error {

	if w.err != nil || sz < 2 {
		return w.err
	}

	if sz = w.makeAlign(sz); sz == 5 {
		sz = 8
	}

	return w.buf.FillMod(sz)
}

func (w *ndr64) WriteTrailingGap(sz int) error {
	return w.WriteAlign(sz)
}

// WriteUnionAlign function writes the union alignment to the buffer.
func (w *ndr64) WriteUnionAlign(sz int) error {
	return w.WriteAlign(sz)
}

// ReadAlign function read the alignment required for the data
// of size `sz`.
func (w *ndr64) ReadAlign(sz int) error {

	if w.err != nil || sz < 2 {
		return w.err
	}

	if sz = w.makeAlign(sz); sz == 5 {
		sz = 8
	}

	return w.buf.SkipMod(sz)
}

func (w *ndr64) ReadTrailingGap(sz int) error {
	return w.ReadAlign(sz)
}

// ReadUnionAlign function reads the union alignment from the buffer.
func (w *ndr64) ReadUnionAlign(sz int) error {
	return w.ReadAlign(sz)
}

// ReadSize function reads the size value from the buffer.
func (w *ndr64) ReadSize(sz *uint64) error {

	if w.err != nil || w.opaque {
		*sz = 0 // clear size_info.
		return w.err
	}

	return w.ReadData(sz)
}

// ReadSwitch function reads the switch value from the buffer.
func (w *ndr64) ReadSwitch(sw any) error {

	if w.err != nil || w.opaque {
		return w.err
	}

	if enum, ok := sw.(EnumWrapper); ok {
		return w.ReadEnum(enum.Value())
	}

	return w.ReadData(sw)
}

// ReadEnum function reads the enumeration value from the buffer.
func (w *ndr64) ReadEnum(enum any) error {

	if w.err != nil {
		return w.err
	}

	if w.opaque {
		return w.ReadData(enum)
	}

	val := uint32(0)

	if err := w.ReadData(&val); err != nil {
		return err
	}

	switch enum := enum.(type) {
	case *uint32:
		*enum = val
	case *uint16:
		*enum = uint16(val)
	default:
		return w.SetErr(fmt.Errorf("unsupported type %T for enum", enum))
	}

	return nil
}

// WriteSize function writes the size value to the buffer.
func (w *ndr64) WriteSize(sz uint64) error {

	if w.err != nil || w.opaque {
		return w.err
	}

	return w.WriteData(sz)
}

// WriteSwitch function writes the switch value to the buffer.
func (w *ndr64) WriteSwitch(sw any) error {

	if w.err != nil || w.opaque {
		return w.err
	}

	if enum, ok := sw.(EnumWrapper); ok {
		return w.WriteEnum(enum.Value())
	}

	return w.WriteData(sw)
}

// WriteEnum function writes the enumeration value to the buffer.
func (w *ndr64) WriteEnum(enum any) error {

	if w.err != nil {
		return w.err
	}

	if w.opaque {
		return w.WriteData(enum)
	}

	var val uint32

	switch enum := enum.(type) {
	case uint32:
		val = enum
	case uint16:
		val = uint32(enum)
	default:
		return w.SetErr(fmt.Errorf("unsupported type %T for enum", enum))
	}

	return w.WriteData(val)
}

// ReadPointer function reads the pointer value and defers the actual data read
// until the ReadDeferred is called. The `setter` value is used in case of the
// pointer aliasing.
func (w *ndr64) ReadPointer(ptr Pointer, setter func(interface{}), mrs ...Unmarshaler) error {

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

	var pptr uint64

	if err := w.ReadData(&pptr); err != nil {
		return w.SetErr(err)
	}

	if pptr == 0 {
		return nil
	}

	/* NDR64 doesn't care about pointers.
	if ptr, ok := w.ptrs[uint64(pptr)]; ok {
		setter(interface{}(ptr))
		return nil
	}
	w.ptrs[pptr], w.rdeferred = ptr, append(w.rdeferred, mrs...)
	*/

	w.rdeferred = append(w.rdeferred, mrs...)

	return nil
}

// WritePointer function writes the pointer to the data and defers the
// actual data write until the WriteDeferred is called.
func (w *ndr64) WritePointer(ptr Pointer, mrs ...Marshaler) error {

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
		return w.SetErr(w.WriteData(uint64(0)))
	}

	if err := w.WriteData(uint64(w.buf.Pos() + 1)); err != nil {
		return w.SetErr(err)
	}

	w.wdeferred = append(w.wdeferred, mrs...)

	return nil
}

// WriteDeferred function writes the deferred pointer values.
func (w *ndr64) WriteDeferred() error {

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
func (w *ndr64) ReadDeferred() error {

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
func (w *ndr64) Marshal(ctx context.Context, mrs Marshaler) ([]byte, error) {

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
func (w *ndr64) Unmarshal(ctx context.Context, mrs Unmarshaler) error {

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
