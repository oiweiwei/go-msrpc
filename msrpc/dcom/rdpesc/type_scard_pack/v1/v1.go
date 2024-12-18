package type_scard_pack

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

var (
	_ = context.Background
	_ = fmt.Errorf
	_ = utf16.Encode
	_ = strings.TrimPrefix
	_ = ndr.ZeroString
	_ = (*uuid.UUID)(nil)
	_ = (*dcerpc.SyntaxID)(nil)
	_ = (*errors.Error)(nil)
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rdpesc"
)

var (
	// Syntax UUID
	TypeSmartCardPackSyntaxUUID = &uuid.UUID{TimeLow: 0xa35af600, TimeMid: 0x9cf4, TimeHiAndVersion: 0x11cd, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x76, Node: [6]uint8{0x8, 0x0, 0x2b, 0x2b, 0xd7, 0x11}}
	// Syntax ID
	TypeSmartCardPackSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: TypeSmartCardPackSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// type_scard_pack interface.
type TypeSmartCardPackClient interface {

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// ReaderSmartCardContext structure represents REDIR_SCARDCONTEXT RPC structure.
//
// REDIR_SCARDCONTEXT represents a context to Smart Cards for Windows on the TS client.
type ReaderSmartCardContext struct {
	// cbContext:  The number of bytes in the pbContext field.
	ContextLength uint32 `idl:"name:cbContext" json:"context_length"`
	// pbContext:  An array of cbContext bytes that contains Smart Cards for Windows context.
	// The data is implementation-specific and MUST NOT be interpreted or changed on the
	// Protocol server.
	Context []byte `idl:"name:pbContext;size_is:(cbContext);pointer:unique" json:"context"`
}

func (o *ReaderSmartCardContext) xxx_PreparePayload(ctx context.Context) error {
	if o.Context != nil && o.ContextLength == 0 {
		o.ContextLength = uint32(len(o.Context))
	}
	if o.ContextLength > uint32(16) {
		return fmt.Errorf("ContextLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReaderSmartCardContext) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ContextLength); err != nil {
		return err
	}
	if o.Context != nil || o.ContextLength > 0 {
		_ptr_pbContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ContextLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Context {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Context[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Context); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Context, _ptr_pbContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReaderSmartCardContext) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ContextLength); err != nil {
		return err
	}
	_ptr_pbContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ContextLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ContextLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Context", sizeInfo[0])
		}
		o.Context = make([]byte, sizeInfo[0])
		for i1 := range o.Context {
			i1 := i1
			if err := w.ReadData(&o.Context[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbContext := func(ptr interface{}) { o.Context = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Context, _s_pbContext, _ptr_pbContext); err != nil {
		return err
	}
	return nil
}

// ReaderSmartCardHandle structure represents REDIR_SCARDHANDLE RPC structure.
//
// REDIR_SCARDHANDLE represents a smart card reader handle associated with Smart Cards
// for Windows context.
type ReaderSmartCardHandle struct {
	// Context:  A valid context, as specified in REDIR_SCARDCONTEXT.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// cbHandle:  The number of bytes in the pbHandle field.
	HandleLength uint32 `idl:"name:cbHandle" json:"handle_length"`
	// pbHandle:  An array of cbHandle bytes that corresponds to a smart card reader handle
	// on the TS client. The data is implementation-specific and MUST NOT be interpreted
	// or changed on the Protocol server.
	Handle []byte `idl:"name:pbHandle;size_is:(cbHandle)" json:"handle"`
}

func (o *ReaderSmartCardHandle) xxx_PreparePayload(ctx context.Context) error {
	if o.Handle != nil && o.HandleLength == 0 {
		o.HandleLength = uint32(len(o.Handle))
	}
	if o.HandleLength > uint32(16) {
		return fmt.Errorf("HandleLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReaderSmartCardHandle) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.HandleLength); err != nil {
		return err
	}
	if o.Handle != nil || o.HandleLength > 0 {
		_ptr_pbHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.HandleLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Handle {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Handle[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Handle); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Handle, _ptr_pbHandle); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReaderSmartCardHandle) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.HandleLength); err != nil {
		return err
	}
	_ptr_pbHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.HandleLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.HandleLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Handle", sizeInfo[0])
		}
		o.Handle = make([]byte, sizeInfo[0])
		for i1 := range o.Handle {
			i1 := i1
			if err := w.ReadData(&o.Handle[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbHandle := func(ptr interface{}) { o.Handle = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Handle, _s_pbHandle, _ptr_pbHandle); err != nil {
		return err
	}
	return nil
}

// LongReturn structure represents long_Return RPC structure.
type LongReturn struct {
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
}

func (o *LongReturn) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LongReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	return nil
}
func (o *LongReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	return nil
}

// ListReaderGroupsReturn structure represents ListReaderGroups_Return RPC structure.
//
// The ListReaderGroups_Return and ListReaders_Return structures are used to obtain
// results for those calls that return a multistring, in addition to a long return value.
// For more information, see sections 3.1.4.5, 3.1.4.6, 3.1.4.7, and 3.1.4.8.
type ListReaderGroupsReturn struct {
	// ReturnCode:  HRESULT or Win32 Error code. The value returned from the Smart Card
	// Redirection call.
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
	// cBytes:  The number of bytes in the msz array field.
	BytesCount uint32 `idl:"name:cBytes" json:"bytes_count"`
	// msz:  The meaning of this field is specific to the context (IOCTL) in which it is
	// used.
	//
	//	+------------------------------------------+-----------------------------------------------------+
	//	|                                          |                                                     |
	//	|                  VALUE                   |                       MEANING                       |
	//	|                                          |                                                     |
	//	+------------------------------------------+-----------------------------------------------------+
	//	+------------------------------------------+-----------------------------------------------------+
	//	| SCARD_IOCTL_LISTREADERSA 0x00090028      | ASCII multistring of readers on the system.         |
	//	+------------------------------------------+-----------------------------------------------------+
	//	| SCARD_IOCTL_LISTREADERSW 0x0009002C      | Unicode multistring of readers on the system.       |
	//	+------------------------------------------+-----------------------------------------------------+
	//	| SCARD_IOCTL_LISTREADERGROUPSA 0x00090020 | ASCII multistring of reader groups on the system.   |
	//	+------------------------------------------+-----------------------------------------------------+
	//	| SCARD_IOCTL_LISTREADERGROUPSW 0x00090024 | Unicode multistring of reader groups on the system. |
	//	+------------------------------------------+-----------------------------------------------------+
	Multistring []byte `idl:"name:msz;size_is:(cBytes);pointer:unique" json:"multistring"`
}

func (o *ListReaderGroupsReturn) xxx_PreparePayload(ctx context.Context) error {
	if o.Multistring != nil && o.BytesCount == 0 {
		o.BytesCount = uint32(len(o.Multistring))
	}
	if o.BytesCount > uint32(65536) {
		return fmt.Errorf("BytesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ListReaderGroupsReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesCount); err != nil {
		return err
	}
	if o.Multistring != nil || o.BytesCount > 0 {
		_ptr_msz := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BytesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Multistring {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Multistring[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Multistring); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Multistring, _ptr_msz); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ListReaderGroupsReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesCount); err != nil {
		return err
	}
	_ptr_msz := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BytesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BytesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Multistring", sizeInfo[0])
		}
		o.Multistring = make([]byte, sizeInfo[0])
		for i1 := range o.Multistring {
			i1 := i1
			if err := w.ReadData(&o.Multistring[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_msz := func(ptr interface{}) { o.Multistring = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Multistring, _s_msz, _ptr_msz); err != nil {
		return err
	}
	return nil
}

// ListReadersReturn structure represents ListReaders_Return RPC structure.
type ListReadersReturn struct {
	ReturnCode  int32  `idl:"name:ReturnCode" json:"return_code"`
	BytesCount  uint32 `idl:"name:cBytes" json:"bytes_count"`
	Multistring []byte `idl:"name:msz;size_is:(cBytes);pointer:unique" json:"multistring"`
}

func (o *ListReadersReturn) xxx_PreparePayload(ctx context.Context) error {
	if o.Multistring != nil && o.BytesCount == 0 {
		o.BytesCount = uint32(len(o.Multistring))
	}
	if o.BytesCount > uint32(65536) {
		return fmt.Errorf("BytesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ListReadersReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesCount); err != nil {
		return err
	}
	if o.Multistring != nil || o.BytesCount > 0 {
		_ptr_msz := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BytesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Multistring {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Multistring[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Multistring); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Multistring, _ptr_msz); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ListReadersReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesCount); err != nil {
		return err
	}
	_ptr_msz := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BytesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BytesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Multistring", sizeInfo[0])
		}
		o.Multistring = make([]byte, sizeInfo[0])
		for i1 := range o.Multistring {
			i1 := i1
			if err := w.ReadData(&o.Multistring[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_msz := func(ptr interface{}) { o.Multistring = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Multistring, _s_msz, _ptr_msz); err != nil {
		return err
	}
	return nil
}

// ContextCall structure represents Context_Call RPC structure.
//
// The Context_Call structure contains Smart Cards for Windows context.
type ContextCall struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
}

func (o *ContextCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ContextAndStringACall structure represents ContextAndStringA_Call RPC structure.
//
// The ContextAndStringA_Call structure contains information used in calls that only
// require a Smart Cards for Windows context and an ASCII string.
type ContextAndStringACall struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// sz:  The value of this string depends on the context (based on IOCTL) in which this
	// structure is used.
	//
	//	+----------------------------------------------+-------------------+
	//	|                                              |                   |
	//	|                    VALUE                     |      MEANING      |
	//	|                                              |                   |
	//	+----------------------------------------------+-------------------+
	//	+----------------------------------------------+-------------------+
	//	| SCARD_IOCTL_INTRODUCEREADERGROUPA 0x00090050 | Reader group name |
	//	+----------------------------------------------+-------------------+
	//	| SCARD_IOCTL_FORGETREADERGROUPA 0x00090058    | Reader group name |
	//	+----------------------------------------------+-------------------+
	//	| SCARD_IOCTL_FORGETREADERA 0x00090068         | Reader name       |
	//	+----------------------------------------------+-------------------+
	String string `idl:"name:sz;string" json:"string"`
}

func (o *ContextAndStringACall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextAndStringACall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.String != "" {
		_ptr_sz := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.String); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.String, _ptr_sz); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextAndStringACall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_sz := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.String); err != nil {
			return err
		}
		return nil
	})
	_s_sz := func(ptr interface{}) { o.String = *ptr.(*string) }
	if err := w.ReadPointer(&o.String, _s_sz, _ptr_sz); err != nil {
		return err
	}
	return nil
}

// ContextAndStringWCall structure represents ContextAndStringW_Call RPC structure.
//
// The ContextAndStringW_Call structure contains information used in calls that only
// require a Smart Cards for Windows context and a Unicode string.
type ContextAndStringWCall struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// sz:  The value of this Unicode string depends on the context (based on IOCTL) in
	// which this structure is used.
	//
	//	+----------------------------------------------+-------------------+
	//	|                                              |                   |
	//	|                    VALUE                     |      MEANING      |
	//	|                                              |                   |
	//	+----------------------------------------------+-------------------+
	//	+----------------------------------------------+-------------------+
	//	| SCARD_IOCTL_INTRODUCEREADERGROUPW 0x00090054 | Reader group name |
	//	+----------------------------------------------+-------------------+
	//	| SCARD_IOCTL_FORGETREADERGROUPW 0x0009005C    | Reader group name |
	//	+----------------------------------------------+-------------------+
	//	| SCARD_IOCTL_FORGETREADERW 0x0009006C         | Reader name       |
	//	+----------------------------------------------+-------------------+
	String string `idl:"name:sz;string" json:"string"`
}

func (o *ContextAndStringWCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextAndStringWCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.String != "" {
		_ptr_sz := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.String); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.String, _ptr_sz); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextAndStringWCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_sz := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.String); err != nil {
			return err
		}
		return nil
	})
	_s_sz := func(ptr interface{}) { o.String = *ptr.(*string) }
	if err := w.ReadPointer(&o.String, _s_sz, _ptr_sz); err != nil {
		return err
	}
	return nil
}

// ContextAndTwoStringACall structure represents ContextAndTwoStringA_Call RPC structure.
//
// The contents of the ContextAndTwoStringA_Call structure are used in those calls that
// require a valid Smart Cards for Windows context (as specified in section 3.2.5) and
// two strings (friendly names).
type ContextAndTwoStringACall struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// sz1:  The value of this ASCII string depends on the context (based on IOCTL) in which
	// it is used.
	//
	//	+-----------------------------------------------+-------------+
	//	|                                               |             |
	//	|                     VALUE                     |   MEANING   |
	//	|                                               |             |
	//	+-----------------------------------------------+-------------+
	//	+-----------------------------------------------+-------------+
	//	| SCARD_IOCTL_INTRODUCEREADERA 0x00090060       | Reader name |
	//	+-----------------------------------------------+-------------+
	//	| SCARD_IOCTL_ADDREADERTOGROUPA 0x00090070      | Reader name |
	//	+-----------------------------------------------+-------------+
	//	| SCARD_IOCTL_REMOVEREADERFROMGROUPA 0x00090078 | Reader name |
	//	+-----------------------------------------------+-------------+
	String1 string `idl:"name:sz1;string" json:"string1"`
	// sz2:  The value of this ASCII string depends on the context (based on IOCTL) in which
	// it is used.
	//
	//	+-----------------------------------------------+-------------------+
	//	|                                               |                   |
	//	|                     VALUE                     |      MEANING      |
	//	|                                               |                   |
	//	+-----------------------------------------------+-------------------+
	//	+-----------------------------------------------+-------------------+
	//	| SCARD_IOCTL_INTRODUCEREADERA 0x00090060       | Device name       |
	//	+-----------------------------------------------+-------------------+
	//	| SCARD_IOCTL_ADDREADERTOGROUPA 0x00090070      | Reader group name |
	//	+-----------------------------------------------+-------------------+
	//	| SCARD_IOCTL_REMOVEREADERFROMGROUPA 0x00090078 | Reader group name |
	//	+-----------------------------------------------+-------------------+
	String2 string `idl:"name:sz2;string" json:"string2"`
}

func (o *ContextAndTwoStringACall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextAndTwoStringACall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.String1 != "" {
		_ptr_sz1 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.String1); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.String1, _ptr_sz1); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.String2 != "" {
		_ptr_sz2 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.String2); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.String2, _ptr_sz2); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextAndTwoStringACall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_sz1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.String1); err != nil {
			return err
		}
		return nil
	})
	_s_sz1 := func(ptr interface{}) { o.String1 = *ptr.(*string) }
	if err := w.ReadPointer(&o.String1, _s_sz1, _ptr_sz1); err != nil {
		return err
	}
	_ptr_sz2 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.String2); err != nil {
			return err
		}
		return nil
	})
	_s_sz2 := func(ptr interface{}) { o.String2 = *ptr.(*string) }
	if err := w.ReadPointer(&o.String2, _s_sz2, _ptr_sz2); err != nil {
		return err
	}
	return nil
}

// ContextAndTwoStringWCall structure represents ContextAndTwoStringW_Call RPC structure.
//
// The contents of the ContextAndTwoStringW_Call structure is used in those calls that
// require a valid Smart Cards for Windows context (as specified in section 3.2.5) and
// two strings (friendly names).
type ContextAndTwoStringWCall struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// sz1:  The value of this Unicode string depends on the context (based on IOCTL) in
	// which it is used.
	//
	//	+-----------------------------------------------+-------------+
	//	|                                               |             |
	//	|                     VALUE                     |   MEANING   |
	//	|                                               |             |
	//	+-----------------------------------------------+-------------+
	//	+-----------------------------------------------+-------------+
	//	| SCARD_IOCTL_INTRODUCEREADERW 0x00090064       | Reader name |
	//	+-----------------------------------------------+-------------+
	//	| SCARD_IOCTL_ADDREADERTOGROUPW 0x00090074      | Reader name |
	//	+-----------------------------------------------+-------------+
	//	| SCARD_IOCTL_REMOVEREADERFROMGROUPW 0x0009007C | Reader name |
	//	+-----------------------------------------------+-------------+
	String1 string `idl:"name:sz1;string" json:"string1"`
	// sz2:  The value of this Unicode string depends on the context (based on IOCTL) in
	// which it is used.
	//
	//	+-----------------------------------------------+-------------------+
	//	|                                               |                   |
	//	|                     VALUE                     |      MEANING      |
	//	|                                               |                   |
	//	+-----------------------------------------------+-------------------+
	//	+-----------------------------------------------+-------------------+
	//	| SCARD_IOCTL_INTRODUCEREADERW 0x00090064       | Device name       |
	//	+-----------------------------------------------+-------------------+
	//	| SCARD_IOCTL_ADDREADERTOGROUPW 0x00090074      | Reader group name |
	//	+-----------------------------------------------+-------------------+
	//	| SCARD_IOCTL_REMOVEREADERFROMGROUPW 0x0009007C | Reader group name |
	//	+-----------------------------------------------+-------------------+
	String2 string `idl:"name:sz2;string" json:"string2"`
}

func (o *ContextAndTwoStringWCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextAndTwoStringWCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.String1 != "" {
		_ptr_sz1 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.String1); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.String1, _ptr_sz1); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.String2 != "" {
		_ptr_sz2 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.String2); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.String2, _ptr_sz2); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextAndTwoStringWCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_sz1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.String1); err != nil {
			return err
		}
		return nil
	})
	_s_sz1 := func(ptr interface{}) { o.String1 = *ptr.(*string) }
	if err := w.ReadPointer(&o.String1, _s_sz1, _ptr_sz1); err != nil {
		return err
	}
	_ptr_sz2 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.String2); err != nil {
			return err
		}
		return nil
	})
	_s_sz2 := func(ptr interface{}) { o.String2 = *ptr.(*string) }
	if err := w.ReadPointer(&o.String2, _s_sz2, _ptr_sz2); err != nil {
		return err
	}
	return nil
}

// EstablishContextCall structure represents EstablishContext_Call RPC structure.
//
// The EstablishContext_Call structure is used to specify the scope of Smart Cards for
// Windows context to be created (for more information, see section 3.1.4.1).
type EstablishContextCall struct {
	// dwScope:  The scope of the context that will be established. The following table
	// shows valid values of this field.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|                                 |                                                                                  |
	//	|              VALUE              |                                     MEANING                                      |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| SCARD_SCOPE_USER 0x00000000     | The context is a user context; any database operations MUST be performed with    |
	//	|                                 | the domain of the user.                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| SCARD_SCOPE_TERMINAL 0x00000001 | The context is a terminal context; any database operations MUST be performed     |
	//	|                                 | with the domain of the terminal. This flag is currently unused; it is here for   |
	//	|                                 | compatibility with [PCSC5] section 3.1.3.                                        |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| SCARD_SCOPE_SYSTEM 0x00000002   | The context is the system context; any database operations MUST be performed     |
	//	|                                 | within the domain of the system.                                                 |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	Scope uint32 `idl:"name:dwScope" json:"scope"`
}

func (o *EstablishContextCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EstablishContextCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Scope); err != nil {
		return err
	}
	return nil
}
func (o *EstablishContextCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Scope); err != nil {
		return err
	}
	return nil
}

// EstablishContextReturn structure represents EstablishContext_Return RPC structure.
//
// The EstablishContext_Return structure is used to provide a response to an Establish
// Context call (for more information, see section 3.1.4.1.)
type EstablishContextReturn struct {
	// ReturnCode:  HRESULT or Win32 Error code. Zero indicates success; any other value
	// indicates failure.
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
}

func (o *EstablishContextReturn) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EstablishContextReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *EstablishContextReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ListReaderGroupsCall structure represents ListReaderGroups_Call RPC structure.
//
// The ListReaderGroups_Call structure contains the parameters for the List Readers
// Groups call (for more information, see sections 3.1.4.5 and 3.1.4.6).
type ListReaderGroupsCall struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// fmszGroupsIsNULL:  A Boolean value specifying whether the caller wants to retrieve
	// just the length of the data. Set to FALSE (0x00000000) in order to allow the data
	// to be returned. Set to TRUE (0x00000001) and only the length of the data will be
	// returned.
	GroupsIsNull int32 `idl:"name:fmszGroupsIsNULL" json:"groups_is_null"`
	// cchGroups:  The length of the string buffer specified by the caller. If cchGroups
	// is set to SCARD_AUTOALLOCATE with a value of 0xFFFFFFFF, a string of any length can
	// be returned. Otherwise, the returned string MUST NOT exceed cchGroups characters
	// in length, including any null characters. When the string to be returned exceeds
	// cchGroups characters in length, including any null characters, ListReaderGroups_Return.ReturnCode
	// MUST be set to SCARD_E_INSUFFICIENT_BUFFER (0x80100008). The cchGroups field MUST
	// be ignored if fmszGroupsIsNULL is set to TRUE (0x00000001). Also, if fmszGroupsIsNULL
	// is set to FALSE (0x00000000) but cchGroups is set to 0x00000000, then the call MUST
	// succeed, ListReaderGroups_Return.cBytes MUST be set to the length of the data, in
	// bytes, and ListReaderGroups_Return.msz MUST be set to NULL.
	GroupsLength uint32 `idl:"name:cchGroups" json:"groups_length"`
}

func (o *ListReaderGroupsCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ListReaderGroupsCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.GroupsIsNull); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupsLength); err != nil {
		return err
	}
	return nil
}
func (o *ListReaderGroupsCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupsIsNull); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupsLength); err != nil {
		return err
	}
	return nil
}

// ListReadersCall structure represents ListReaders_Call RPC structure.
//
// The ListReaders_Call structure contains the parameters for the List Readers call
// (for more information, see sections 3.1.4.7 and 3.1.4.8).
type ListReadersCall struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// cBytes:  The length, in bytes, of reader groups specified in mszGroups.
	BytesCount uint32 `idl:"name:cBytes" json:"bytes_count"`
	// mszGroups:  The names of the reader groups defined in the system. Reader groups not
	// present on the protocol server MUST be ignored. The value of this is dependent on
	// the context (IOCTL) that it is used.
	//
	//	+-------------------------------------+---------------------+
	//	|                                     |                     |
	//	|                VALUE                |       MEANING       |
	//	|                                     |                     |
	//	+-------------------------------------+---------------------+
	//	+-------------------------------------+---------------------+
	//	| SCARD_IOCTL_LISTREADERSA 0x00090028 | ASCII multistring   |
	//	+-------------------------------------+---------------------+
	//	| SCARD_IOCTL_LISTREADERSW 0x0009002C | Unicode multistring |
	//	+-------------------------------------+---------------------+
	Groups []byte `idl:"name:mszGroups;size_is:(cBytes);pointer:unique" json:"groups"`
	// fmszReadersIsNULL:  A Boolean value specifying whether the caller wants to retrieve
	// the length of the data. Set to FALSE (0x00000000) to allow the data to be returned.
	// Set to TRUE (0x00000001), and only the length of the data will be returned.
	ReadersIsNull int32 `idl:"name:fmszReadersIsNULL" json:"readers_is_null"`
	// cchReaders:  The length of the string buffer specified by the caller. If cchReaders
	// is set to SCARD_AUTOALLOCATE with a value of 0xFFFFFFFF, a string of any length can
	// be returned. Otherwise, the returned string MUST NOT exceed cchReaders characters
	// in length, including any NULL characters. When the string to be returned exceeds
	// cchReaders characters in length, including any null characters, ListReaders_Return.ReturnCode
	// MUST be set to SCARD_E_INSUFFICIENT_BUFFER (0x80100008). The cchReaders field MUST
	// be ignored if fmszReadersIsNULL is set to TRUE (0x00000001). Also, if fmszReadersIsNULL
	// is set to FALSE (0x00000000) but cchReaders is set to 0x00000000, then the call MUST
	// succeed, ListReaders_Return.cBytes MUST be set to the length of the data in bytes,
	// and ListReaders_Return.msz MUST be set to NULL.
	ReadersLength uint32 `idl:"name:cchReaders" json:"readers_length"`
}

func (o *ListReadersCall) xxx_PreparePayload(ctx context.Context) error {
	if o.Groups != nil && o.BytesCount == 0 {
		o.BytesCount = uint32(len(o.Groups))
	}
	if o.BytesCount > uint32(65536) {
		return fmt.Errorf("BytesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ListReadersCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BytesCount); err != nil {
		return err
	}
	if o.Groups != nil || o.BytesCount > 0 {
		_ptr_mszGroups := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BytesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Groups {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Groups[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Groups); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Groups, _ptr_mszGroups); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ReadersIsNull); err != nil {
		return err
	}
	if err := w.WriteData(o.ReadersLength); err != nil {
		return err
	}
	return nil
}
func (o *ListReadersCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesCount); err != nil {
		return err
	}
	_ptr_mszGroups := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BytesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BytesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Groups", sizeInfo[0])
		}
		o.Groups = make([]byte, sizeInfo[0])
		for i1 := range o.Groups {
			i1 := i1
			if err := w.ReadData(&o.Groups[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_mszGroups := func(ptr interface{}) { o.Groups = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Groups, _s_mszGroups, _ptr_mszGroups); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadersIsNull); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadersLength); err != nil {
		return err
	}
	return nil
}

// ReaderStateCommonCall structure represents ReaderState_Common_Call RPC structure.
//
// The ReaderState_Common_Call structure contains the state of the reader at the time
// of the call as seen by the caller.
type ReaderStateCommonCall struct {
	// dwCurrentState:  A bitmap that specifies the current reader state according to the
	// TS client. Possible values are specified in section 2.2.7.
	CurrentState uint32 `idl:"name:dwCurrentState" json:"current_state"`
	// dwEventState:  A bitmap that defines the state of the reader after a state change.
	// Possible values are specified in section 2.2.7.
	EventState uint32 `idl:"name:dwEventState" json:"event_state"`
	// cbAtr:  The number of bytes used in the ATR string.
	AttributeLength uint32 `idl:"name:cbAtr" json:"attribute_length"`
	// rgbAtr:  The value for the card's ATR string. If cbAtr is NOT zero, this value MUST
	// be formatted in accordance to [ISO/IEC-7816-3] section 8. Unused bytes MUST be set
	// to 0 and MUST be ignored.
	Attribute []byte `idl:"name:rgbAtr" json:"attribute"`
}

func (o *ReaderStateCommonCall) xxx_PreparePayload(ctx context.Context) error {
	if o.AttributeLength > uint32(36) {
		return fmt.Errorf("AttributeLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReaderStateCommonCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentState); err != nil {
		return err
	}
	if err := w.WriteData(o.EventState); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeLength); err != nil {
		return err
	}
	for i1 := range o.Attribute {
		i1 := i1
		if uint64(i1) >= 36 {
			break
		}
		if err := w.WriteData(o.Attribute[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Attribute); uint64(i1) < 36; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReaderStateCommonCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentState); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventState); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeLength); err != nil {
		return err
	}
	o.Attribute = make([]byte, 36)
	for i1 := range o.Attribute {
		i1 := i1
		if err := w.ReadData(&o.Attribute[i1]); err != nil {
			return err
		}
	}
	return nil
}

// ReaderStateA structure represents ReaderStateA RPC structure.
//
// The ReaderStateA structure contains information used in calls that only require Smart
// Cards for Windows context and an ASCII string.
type ReaderStateA struct {
	// szReader:  An ASCII string specifying the reader name.
	Reader string `idl:"name:szReader;string" json:"reader"`
	// Common:  A packet that specifies the state of the reader at the time of the call.
	// For information about this packet, see section 2.2.1.5.
	Common *ReaderStateCommonCall `idl:"name:Common" json:"common"`
}

func (o *ReaderStateA) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReaderStateA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Reader != "" {
		_ptr_szReader := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.Reader); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Reader, _ptr_szReader); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Common != nil {
		if err := o.Common.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderStateCommonCall{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReaderStateA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_szReader := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.Reader); err != nil {
			return err
		}
		return nil
	})
	_s_szReader := func(ptr interface{}) { o.Reader = *ptr.(*string) }
	if err := w.ReadPointer(&o.Reader, _s_szReader, _ptr_szReader); err != nil {
		return err
	}
	if o.Common == nil {
		o.Common = &ReaderStateCommonCall{}
	}
	if err := o.Common.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ReaderStateW structure represents ReaderStateW RPC structure.
//
// The ReaderStateW structure is a Unicode representation of the state of a smart card
// reader.
type ReaderStateW struct {
	// szReader:  A Unicode string specifying the reader name.
	Reader string `idl:"name:szReader;string" json:"reader"`
	// Common:  A packet that specifies the state of the reader at the time of the call.
	// For information about this packet, see section 2.2.1.5.
	Common *ReaderStateCommonCall `idl:"name:Common" json:"common"`
}

func (o *ReaderStateW) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReaderStateW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Reader != "" {
		_ptr_szReader := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Reader); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Reader, _ptr_szReader); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Common != nil {
		if err := o.Common.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderStateCommonCall{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReaderStateW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_szReader := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Reader); err != nil {
			return err
		}
		return nil
	})
	_s_szReader := func(ptr interface{}) { o.Reader = *ptr.(*string) }
	if err := w.ReadPointer(&o.Reader, _s_szReader, _ptr_szReader); err != nil {
		return err
	}
	if o.Common == nil {
		o.Common = &ReaderStateCommonCall{}
	}
	if err := o.Common.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ReaderStateReturn structure represents ReaderState_Return RPC structure.
//
// The ReaderState_Return structure specifies state information returned from Smart
// Cards for Windows.
type ReaderStateReturn struct {
	// dwCurrentState:  A bitmap that defines the current state of the reader at the time
	// of the call. Possible values are specified in section 2.2.7.
	CurrentState uint32 `idl:"name:dwCurrentState" json:"current_state"`
	// dwEventState:  A bitmap that defines the state of the reader after a state change
	// as seen by Smart Cards for Windows. Possible values are specified in section 2.2.7.
	EventState uint32 `idl:"name:dwEventState" json:"event_state"`
	// cbAtr:  The number of used bytes in rgbAtr.
	AttributeLength uint32 `idl:"name:cbAtr" json:"attribute_length"`
	// rgbAtr:  The values for the card's ATR string. Unused bytes MUST be set to zero and
	// MUST be ignored on receipt.
	Attribute []byte `idl:"name:rgbAtr" json:"attribute"`
}

func (o *ReaderStateReturn) xxx_PreparePayload(ctx context.Context) error {
	if o.AttributeLength > uint32(36) {
		return fmt.Errorf("AttributeLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReaderStateReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentState); err != nil {
		return err
	}
	if err := w.WriteData(o.EventState); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeLength); err != nil {
		return err
	}
	for i1 := range o.Attribute {
		i1 := i1
		if uint64(i1) >= 36 {
			break
		}
		if err := w.WriteData(o.Attribute[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Attribute); uint64(i1) < 36; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReaderStateReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentState); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventState); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeLength); err != nil {
		return err
	}
	o.Attribute = make([]byte, 36)
	for i1 := range o.Attribute {
		i1 := i1
		if err := w.ReadData(&o.Attribute[i1]); err != nil {
			return err
		}
	}
	return nil
}

// GetStatusChangeACall structure represents GetStatusChangeA_Call RPC structure.
//
// The GetStatusChangeA_Call structure provides the state change in the reader as specified
// in section 3.1.4.23.
type GetStatusChangeACall struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// dwTimeOut:  The maximum amount of time, in milliseconds, to wait for an action. If
	// this member is set to 0xFFFFFFFF (INFINITE), the caller MUST wait until an action
	// occurs.
	Timeout uint32 `idl:"name:dwTimeOut" json:"timeout"`
	// cReaders:  The number of ReaderStates to track.
	ReadersCount uint32 `idl:"name:cReaders" json:"readers_count"`
	// rgReaderStates:  Smart card readers that the caller is tracking.
	ReaderStates []*ReaderStateA `idl:"name:rgReaderStates;size_is:(cReaders)" json:"reader_states"`
}

func (o *GetStatusChangeACall) xxx_PreparePayload(ctx context.Context) error {
	if o.ReaderStates != nil && o.ReadersCount == 0 {
		o.ReadersCount = uint32(len(o.ReaderStates))
	}
	if o.ReadersCount > uint32(11) {
		return fmt.Errorf("ReadersCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetStatusChangeACall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if err := w.WriteData(o.ReadersCount); err != nil {
		return err
	}
	if o.ReaderStates != nil || o.ReadersCount > 0 {
		_ptr_rgReaderStates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ReadersCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ReaderStates {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ReaderStates[i1] != nil {
					if err := o.ReaderStates[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReaderStateA{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ReaderStates); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ReaderStateA{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReaderStates, _ptr_rgReaderStates); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetStatusChangeACall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadersCount); err != nil {
		return err
	}
	_ptr_rgReaderStates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ReadersCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ReadersCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ReaderStates", sizeInfo[0])
		}
		o.ReaderStates = make([]*ReaderStateA, sizeInfo[0])
		for i1 := range o.ReaderStates {
			i1 := i1
			if o.ReaderStates[i1] == nil {
				o.ReaderStates[i1] = &ReaderStateA{}
			}
			if err := o.ReaderStates[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgReaderStates := func(ptr interface{}) { o.ReaderStates = *ptr.(*[]*ReaderStateA) }
	if err := w.ReadPointer(&o.ReaderStates, _s_rgReaderStates, _ptr_rgReaderStates); err != nil {
		return err
	}
	return nil
}

// LocateCardsACall structure represents LocateCardsA_Call RPC structure.
//
// The parameters of the LocateCardsA_Call structure specify the list of smart card
// readers to search for the specified card types. For call information, see section
// 3.1.4.21.
type LocateCardsACall struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// cBytes:  The number of bytes in the mszCards field.
	BytesCount uint32 `idl:"name:cBytes" json:"bytes_count"`
	// mszCards:  An ASCII multistring of card names to locate. Card names MUST be registered
	// in Smart Cards for Windows. Unknown card types MUST be ignored.
	Cards []byte `idl:"name:mszCards;size_is:(cBytes)" json:"cards"`
	// cReaders:  The number of reader state structures.
	ReadersCount uint32 `idl:"name:cReaders" json:"readers_count"`
	// rgReaderStates:  The reader state information specifying which readers are searched
	// for the cards listed in mszCards.
	ReaderStates []*ReaderStateA `idl:"name:rgReaderStates;size_is:(cReaders)" json:"reader_states"`
}

func (o *LocateCardsACall) xxx_PreparePayload(ctx context.Context) error {
	if o.Cards != nil && o.BytesCount == 0 {
		o.BytesCount = uint32(len(o.Cards))
	}
	if o.ReaderStates != nil && o.ReadersCount == 0 {
		o.ReadersCount = uint32(len(o.ReaderStates))
	}
	if o.BytesCount > uint32(65536) {
		return fmt.Errorf("BytesCount is out of range")
	}
	if o.ReadersCount > uint32(10) {
		return fmt.Errorf("ReadersCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LocateCardsACall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BytesCount); err != nil {
		return err
	}
	if o.Cards != nil || o.BytesCount > 0 {
		_ptr_mszCards := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BytesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Cards {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Cards[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Cards); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Cards, _ptr_mszCards); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ReadersCount); err != nil {
		return err
	}
	if o.ReaderStates != nil || o.ReadersCount > 0 {
		_ptr_rgReaderStates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ReadersCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ReaderStates {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ReaderStates[i1] != nil {
					if err := o.ReaderStates[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReaderStateA{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ReaderStates); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ReaderStateA{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReaderStates, _ptr_rgReaderStates); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *LocateCardsACall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesCount); err != nil {
		return err
	}
	_ptr_mszCards := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BytesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BytesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Cards", sizeInfo[0])
		}
		o.Cards = make([]byte, sizeInfo[0])
		for i1 := range o.Cards {
			i1 := i1
			if err := w.ReadData(&o.Cards[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_mszCards := func(ptr interface{}) { o.Cards = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Cards, _s_mszCards, _ptr_mszCards); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadersCount); err != nil {
		return err
	}
	_ptr_rgReaderStates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ReadersCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ReadersCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ReaderStates", sizeInfo[0])
		}
		o.ReaderStates = make([]*ReaderStateA, sizeInfo[0])
		for i1 := range o.ReaderStates {
			i1 := i1
			if o.ReaderStates[i1] == nil {
				o.ReaderStates[i1] = &ReaderStateA{}
			}
			if err := o.ReaderStates[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgReaderStates := func(ptr interface{}) { o.ReaderStates = *ptr.(*[]*ReaderStateA) }
	if err := w.ReadPointer(&o.ReaderStates, _s_rgReaderStates, _ptr_rgReaderStates); err != nil {
		return err
	}
	return nil
}

// LocateCardsWCall structure represents LocateCardsW_Call RPC structure.
//
// The parameters of the LocateCardsW_Call structure specify the list of smart card
// readers to search for the specified card types. For more information, see section
// 3.1.4.22.
type LocateCardsWCall struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// cBytes:  The number of bytes in the mszCards field.
	BytesCount uint32 `idl:"name:cBytes" json:"bytes_count"`
	// mszCards:  A Unicode multistring of card names to locate. Card names MUST be registered
	// in Smart Cards for Windows. Unknown card types MUST be ignored.
	Cards []byte `idl:"name:mszCards;size_is:(cBytes)" json:"cards"`
	// cReaders:  The number of reader state structures.
	ReadersCount uint32 `idl:"name:cReaders" json:"readers_count"`
	// rgReaderStates:  The reader state information used to locate the cards listed in
	// mszCards.
	ReaderStates []*ReaderStateW `idl:"name:rgReaderStates;size_is:(cReaders)" json:"reader_states"`
}

func (o *LocateCardsWCall) xxx_PreparePayload(ctx context.Context) error {
	if o.Cards != nil && o.BytesCount == 0 {
		o.BytesCount = uint32(len(o.Cards))
	}
	if o.ReaderStates != nil && o.ReadersCount == 0 {
		o.ReadersCount = uint32(len(o.ReaderStates))
	}
	if o.BytesCount > uint32(65536) {
		return fmt.Errorf("BytesCount is out of range")
	}
	if o.ReadersCount > uint32(10) {
		return fmt.Errorf("ReadersCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LocateCardsWCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BytesCount); err != nil {
		return err
	}
	if o.Cards != nil || o.BytesCount > 0 {
		_ptr_mszCards := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BytesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Cards {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Cards[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Cards); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Cards, _ptr_mszCards); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ReadersCount); err != nil {
		return err
	}
	if o.ReaderStates != nil || o.ReadersCount > 0 {
		_ptr_rgReaderStates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ReadersCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ReaderStates {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ReaderStates[i1] != nil {
					if err := o.ReaderStates[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReaderStateW{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ReaderStates); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ReaderStateW{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReaderStates, _ptr_rgReaderStates); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *LocateCardsWCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesCount); err != nil {
		return err
	}
	_ptr_mszCards := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BytesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BytesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Cards", sizeInfo[0])
		}
		o.Cards = make([]byte, sizeInfo[0])
		for i1 := range o.Cards {
			i1 := i1
			if err := w.ReadData(&o.Cards[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_mszCards := func(ptr interface{}) { o.Cards = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Cards, _s_mszCards, _ptr_mszCards); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadersCount); err != nil {
		return err
	}
	_ptr_rgReaderStates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ReadersCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ReadersCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ReaderStates", sizeInfo[0])
		}
		o.ReaderStates = make([]*ReaderStateW, sizeInfo[0])
		for i1 := range o.ReaderStates {
			i1 := i1
			if o.ReaderStates[i1] == nil {
				o.ReaderStates[i1] = &ReaderStateW{}
			}
			if err := o.ReaderStates[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgReaderStates := func(ptr interface{}) { o.ReaderStates = *ptr.(*[]*ReaderStateW) }
	if err := w.ReadPointer(&o.ReaderStates, _s_rgReaderStates, _ptr_rgReaderStates); err != nil {
		return err
	}
	return nil
}

// LocateCardsAttributeMask structure represents LocateCards_ATRMask RPC structure.
//
// The LocateCards_ATRMask structure contains the information to identify a card type.
type LocateCardsAttributeMask struct {
	// cbAtr:  The number of bytes used in the rgbAtr and rgbMask fields.
	AttributeLength uint32 `idl:"name:cbAtr" json:"attribute_length"`
	// rgbAtr:  Values for the card's Answer To Reset (ATR) string. This value MUST be formatted
	// as specified in [ISO/IEC-7816-3] section 8. Unused bytes MUST be set to 0 and MUST
	// be ignored.
	Attribute []byte `idl:"name:rgbAtr" json:"attribute"`
	// rgbMask:  Values for the mask for the card's ATR string. Each bit that cannot vary
	// between cards of the same type MUST be set to 1. Unused bytes MUST be set to 0 and
	// MUST be ignored.
	Mask []byte `idl:"name:rgbMask" json:"mask"`
}

func (o *LocateCardsAttributeMask) xxx_PreparePayload(ctx context.Context) error {
	if o.AttributeLength > uint32(36) {
		return fmt.Errorf("AttributeLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LocateCardsAttributeMask) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeLength); err != nil {
		return err
	}
	for i1 := range o.Attribute {
		i1 := i1
		if uint64(i1) >= 36 {
			break
		}
		if err := w.WriteData(o.Attribute[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Attribute); uint64(i1) < 36; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Mask {
		i1 := i1
		if uint64(i1) >= 36 {
			break
		}
		if err := w.WriteData(o.Mask[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Mask); uint64(i1) < 36; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *LocateCardsAttributeMask) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeLength); err != nil {
		return err
	}
	o.Attribute = make([]byte, 36)
	for i1 := range o.Attribute {
		i1 := i1
		if err := w.ReadData(&o.Attribute[i1]); err != nil {
			return err
		}
	}
	o.Mask = make([]byte, 36)
	for i1 := range o.Mask {
		i1 := i1
		if err := w.ReadData(&o.Mask[i1]); err != nil {
			return err
		}
	}
	return nil
}

// LocateCardsByAtraCall structure represents LocateCardsByATRA_Call RPC structure.
//
// The LocateCardsByATRA_Call structure returns information concerning the status of
// the smart card of interest (ATR).
type LocateCardsByAtraCall struct {
	// Context:  A valid context, as specified in section 2.2.2.13.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// cAtrs:  The number of bytes in the rgAtrMasks field.
	AttributesCount uint32 `idl:"name:cAtrs" json:"attributes_count"`
	// rgAtrMasks:  An array of ATRs to match against currently inserted cards.
	AttributeMasks []*LocateCardsAttributeMask `idl:"name:rgAtrMasks;size_is:(cAtrs)" json:"attribute_masks"`
	// cReaders:  The number of elements in the rgReaderStates field.
	ReadersCount uint32 `idl:"name:cReaders" json:"readers_count"`
	// rgReaderStates:  The states of the readers that the application is monitoring. The
	// states reflect what the application determines to be the current states of the readers
	// and that might differ from the actual states.
	ReaderStates []*ReaderStateA `idl:"name:rgReaderStates;size_is:(cReaders)" json:"reader_states"`
}

func (o *LocateCardsByAtraCall) xxx_PreparePayload(ctx context.Context) error {
	if o.AttributeMasks != nil && o.AttributesCount == 0 {
		o.AttributesCount = uint32(len(o.AttributeMasks))
	}
	if o.ReaderStates != nil && o.ReadersCount == 0 {
		o.ReadersCount = uint32(len(o.ReaderStates))
	}
	if o.AttributesCount > uint32(1000) {
		return fmt.Errorf("AttributesCount is out of range")
	}
	if o.ReadersCount > uint32(10) {
		return fmt.Errorf("ReadersCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LocateCardsByAtraCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AttributesCount); err != nil {
		return err
	}
	if o.AttributeMasks != nil || o.AttributesCount > 0 {
		_ptr_rgAtrMasks := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AttributesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.AttributeMasks {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.AttributeMasks[i1] != nil {
					if err := o.AttributeMasks[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&LocateCardsAttributeMask{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.AttributeMasks); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&LocateCardsAttributeMask{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AttributeMasks, _ptr_rgAtrMasks); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ReadersCount); err != nil {
		return err
	}
	if o.ReaderStates != nil || o.ReadersCount > 0 {
		_ptr_rgReaderStates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ReadersCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ReaderStates {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ReaderStates[i1] != nil {
					if err := o.ReaderStates[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReaderStateA{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ReaderStates); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ReaderStateA{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReaderStates, _ptr_rgReaderStates); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *LocateCardsByAtraCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributesCount); err != nil {
		return err
	}
	_ptr_rgAtrMasks := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AttributesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AttributesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AttributeMasks", sizeInfo[0])
		}
		o.AttributeMasks = make([]*LocateCardsAttributeMask, sizeInfo[0])
		for i1 := range o.AttributeMasks {
			i1 := i1
			if o.AttributeMasks[i1] == nil {
				o.AttributeMasks[i1] = &LocateCardsAttributeMask{}
			}
			if err := o.AttributeMasks[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgAtrMasks := func(ptr interface{}) { o.AttributeMasks = *ptr.(*[]*LocateCardsAttributeMask) }
	if err := w.ReadPointer(&o.AttributeMasks, _s_rgAtrMasks, _ptr_rgAtrMasks); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadersCount); err != nil {
		return err
	}
	_ptr_rgReaderStates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ReadersCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ReadersCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ReaderStates", sizeInfo[0])
		}
		o.ReaderStates = make([]*ReaderStateA, sizeInfo[0])
		for i1 := range o.ReaderStates {
			i1 := i1
			if o.ReaderStates[i1] == nil {
				o.ReaderStates[i1] = &ReaderStateA{}
			}
			if err := o.ReaderStates[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgReaderStates := func(ptr interface{}) { o.ReaderStates = *ptr.(*[]*ReaderStateA) }
	if err := w.ReadPointer(&o.ReaderStates, _s_rgReaderStates, _ptr_rgReaderStates); err != nil {
		return err
	}
	return nil
}

// LocateCardsByAtrwCall structure represents LocateCardsByATRW_Call RPC structure.
//
// The LocateCardsByATRW_Call structure returns information concerning the status of
// the smart card of interest (ATR).
type LocateCardsByAtrwCall struct {
	// Context:  A valid context, as specified in section 2.2.2.14.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// cAtrs:  The number of bytes in the rgAtrMasks field.
	AttributesCount uint32 `idl:"name:cAtrs" json:"attributes_count"`
	// rgAtrMasks:  An array of ATRs to match against currently inserted cards.
	AttributeMasks []*LocateCardsAttributeMask `idl:"name:rgAtrMasks;size_is:(cAtrs)" json:"attribute_masks"`
	// cReaders:  The number of elements in the rgReaderStates field.
	ReadersCount uint32 `idl:"name:cReaders" json:"readers_count"`
	// rgReaderStates:  The states of the readers that the application is monitoring. The
	// states reflects what the application believes is the current states of the readers
	// and might differ from the actual states.
	ReaderStates []*ReaderStateW `idl:"name:rgReaderStates;size_is:(cReaders)" json:"reader_states"`
}

func (o *LocateCardsByAtrwCall) xxx_PreparePayload(ctx context.Context) error {
	if o.AttributeMasks != nil && o.AttributesCount == 0 {
		o.AttributesCount = uint32(len(o.AttributeMasks))
	}
	if o.ReaderStates != nil && o.ReadersCount == 0 {
		o.ReadersCount = uint32(len(o.ReaderStates))
	}
	if o.AttributesCount > uint32(1000) {
		return fmt.Errorf("AttributesCount is out of range")
	}
	if o.ReadersCount > uint32(10) {
		return fmt.Errorf("ReadersCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LocateCardsByAtrwCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AttributesCount); err != nil {
		return err
	}
	if o.AttributeMasks != nil || o.AttributesCount > 0 {
		_ptr_rgAtrMasks := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AttributesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.AttributeMasks {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.AttributeMasks[i1] != nil {
					if err := o.AttributeMasks[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&LocateCardsAttributeMask{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.AttributeMasks); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&LocateCardsAttributeMask{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AttributeMasks, _ptr_rgAtrMasks); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ReadersCount); err != nil {
		return err
	}
	if o.ReaderStates != nil || o.ReadersCount > 0 {
		_ptr_rgReaderStates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ReadersCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ReaderStates {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ReaderStates[i1] != nil {
					if err := o.ReaderStates[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReaderStateW{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ReaderStates); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ReaderStateW{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReaderStates, _ptr_rgReaderStates); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *LocateCardsByAtrwCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributesCount); err != nil {
		return err
	}
	_ptr_rgAtrMasks := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AttributesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AttributesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AttributeMasks", sizeInfo[0])
		}
		o.AttributeMasks = make([]*LocateCardsAttributeMask, sizeInfo[0])
		for i1 := range o.AttributeMasks {
			i1 := i1
			if o.AttributeMasks[i1] == nil {
				o.AttributeMasks[i1] = &LocateCardsAttributeMask{}
			}
			if err := o.AttributeMasks[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgAtrMasks := func(ptr interface{}) { o.AttributeMasks = *ptr.(*[]*LocateCardsAttributeMask) }
	if err := w.ReadPointer(&o.AttributeMasks, _s_rgAtrMasks, _ptr_rgAtrMasks); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadersCount); err != nil {
		return err
	}
	_ptr_rgReaderStates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ReadersCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ReadersCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ReaderStates", sizeInfo[0])
		}
		o.ReaderStates = make([]*ReaderStateW, sizeInfo[0])
		for i1 := range o.ReaderStates {
			i1 := i1
			if o.ReaderStates[i1] == nil {
				o.ReaderStates[i1] = &ReaderStateW{}
			}
			if err := o.ReaderStates[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgReaderStates := func(ptr interface{}) { o.ReaderStates = *ptr.(*[]*ReaderStateW) }
	if err := w.ReadPointer(&o.ReaderStates, _s_rgReaderStates, _ptr_rgReaderStates); err != nil {
		return err
	}
	return nil
}

// LocateCardsReturn structure represents LocateCards_Return RPC structure.
//
// The LocateCards_Return and GetStatusChange_Return structures are used to obtain the
// results on those calls that return updated reader state information. (for more information,
// see sections 3.1.4.21, 3.1.4.22, 3.1.4.23, 3.1.4.24, 3.1.4.25, and 3.1.4.26).
type LocateCardsReturn struct {
	// ReturnCode:  HRESULT or Win32 Error code. Zero indicates success; any other value
	// indicates failure.
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
	// cReaders:  The number of elements in the rgReaderStates field.
	ReadersCount uint32 `idl:"name:cReaders" json:"readers_count"`
	// rgReaderStates:  The current states of the readers being watched.
	ReaderStates []*ReaderStateReturn `idl:"name:rgReaderStates;size_is:(cReaders)" json:"reader_states"`
}

func (o *LocateCardsReturn) xxx_PreparePayload(ctx context.Context) error {
	if o.ReaderStates != nil && o.ReadersCount == 0 {
		o.ReadersCount = uint32(len(o.ReaderStates))
	}
	if o.ReadersCount > uint32(10) {
		return fmt.Errorf("ReadersCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LocateCardsReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.ReadersCount); err != nil {
		return err
	}
	if o.ReaderStates != nil || o.ReadersCount > 0 {
		_ptr_rgReaderStates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ReadersCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ReaderStates {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ReaderStates[i1] != nil {
					if err := o.ReaderStates[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReaderStateReturn{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ReaderStates); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ReaderStateReturn{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReaderStates, _ptr_rgReaderStates); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *LocateCardsReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadersCount); err != nil {
		return err
	}
	_ptr_rgReaderStates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ReadersCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ReadersCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ReaderStates", sizeInfo[0])
		}
		o.ReaderStates = make([]*ReaderStateReturn, sizeInfo[0])
		for i1 := range o.ReaderStates {
			i1 := i1
			if o.ReaderStates[i1] == nil {
				o.ReaderStates[i1] = &ReaderStateReturn{}
			}
			if err := o.ReaderStates[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgReaderStates := func(ptr interface{}) { o.ReaderStates = *ptr.(*[]*ReaderStateReturn) }
	if err := w.ReadPointer(&o.ReaderStates, _s_rgReaderStates, _ptr_rgReaderStates); err != nil {
		return err
	}
	return nil
}

// GetStatusChangeReturn structure represents GetStatusChange_Return RPC structure.
type GetStatusChangeReturn struct {
	ReturnCode   int32                `idl:"name:ReturnCode" json:"return_code"`
	ReadersCount uint32               `idl:"name:cReaders" json:"readers_count"`
	ReaderStates []*ReaderStateReturn `idl:"name:rgReaderStates;size_is:(cReaders)" json:"reader_states"`
}

func (o *GetStatusChangeReturn) xxx_PreparePayload(ctx context.Context) error {
	if o.ReaderStates != nil && o.ReadersCount == 0 {
		o.ReadersCount = uint32(len(o.ReaderStates))
	}
	if o.ReadersCount > uint32(10) {
		return fmt.Errorf("ReadersCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetStatusChangeReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.ReadersCount); err != nil {
		return err
	}
	if o.ReaderStates != nil || o.ReadersCount > 0 {
		_ptr_rgReaderStates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ReadersCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ReaderStates {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ReaderStates[i1] != nil {
					if err := o.ReaderStates[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReaderStateReturn{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ReaderStates); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ReaderStateReturn{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReaderStates, _ptr_rgReaderStates); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetStatusChangeReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadersCount); err != nil {
		return err
	}
	_ptr_rgReaderStates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ReadersCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ReadersCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ReaderStates", sizeInfo[0])
		}
		o.ReaderStates = make([]*ReaderStateReturn, sizeInfo[0])
		for i1 := range o.ReaderStates {
			i1 := i1
			if o.ReaderStates[i1] == nil {
				o.ReaderStates[i1] = &ReaderStateReturn{}
			}
			if err := o.ReaderStates[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgReaderStates := func(ptr interface{}) { o.ReaderStates = *ptr.(*[]*ReaderStateReturn) }
	if err := w.ReadPointer(&o.ReaderStates, _s_rgReaderStates, _ptr_rgReaderStates); err != nil {
		return err
	}
	return nil
}

// GetStatusChangeWCall structure represents GetStatusChangeW_Call RPC structure.
//
// The GetStatusChangeW_Call structure provides the state change in the Reader as specified
// in section 3.1.4.24.
type GetStatusChangeWCall struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// dwTimeOut:  Maximum amount of time, in milliseconds, to wait for an action. If set
	// to 0xFFFFFFFF (INFINITE), the caller MUST wait until an action occurs.
	Timeout uint32 `idl:"name:dwTimeOut" json:"timeout"`
	// cReaders:  The number of ReaderStates to track.
	ReadersCount uint32 `idl:"name:cReaders" json:"readers_count"`
	// rgReaderStates:  Smart card readers that the caller is tracking.
	ReaderStates []*ReaderStateW `idl:"name:rgReaderStates;size_is:(cReaders)" json:"reader_states"`
}

func (o *GetStatusChangeWCall) xxx_PreparePayload(ctx context.Context) error {
	if o.ReaderStates != nil && o.ReadersCount == 0 {
		o.ReadersCount = uint32(len(o.ReaderStates))
	}
	if o.ReadersCount > uint32(11) {
		return fmt.Errorf("ReadersCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetStatusChangeWCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if err := w.WriteData(o.ReadersCount); err != nil {
		return err
	}
	if o.ReaderStates != nil || o.ReadersCount > 0 {
		_ptr_rgReaderStates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ReadersCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ReaderStates {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ReaderStates[i1] != nil {
					if err := o.ReaderStates[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReaderStateW{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ReaderStates); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ReaderStateW{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReaderStates, _ptr_rgReaderStates); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetStatusChangeWCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadersCount); err != nil {
		return err
	}
	_ptr_rgReaderStates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ReadersCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ReadersCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ReaderStates", sizeInfo[0])
		}
		o.ReaderStates = make([]*ReaderStateW, sizeInfo[0])
		for i1 := range o.ReaderStates {
			i1 := i1
			if o.ReaderStates[i1] == nil {
				o.ReaderStates[i1] = &ReaderStateW{}
			}
			if err := o.ReaderStates[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgReaderStates := func(ptr interface{}) { o.ReaderStates = *ptr.(*[]*ReaderStateW) }
	if err := w.ReadPointer(&o.ReaderStates, _s_rgReaderStates, _ptr_rgReaderStates); err != nil {
		return err
	}
	return nil
}

// ConnectCommon structure represents Connect_Common RPC structure.
//
// The Connect_Common structure contains information common to both versions of the
// Connect function (for more information, see sections 2.2.2.13 and 2.2.2.14).
type ConnectCommon struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// dwShareMode:  A flag that indicates whether other applications are allowed to form
	// connections to the card. Possible values of this field are specified in section 2.2.6.
	ShareMode uint32 `idl:"name:dwShareMode" json:"share_mode"`
	// dwPreferredProtocols:  A bitmask of acceptable protocols for the connection, as specified
	// in section 2.2.5.
	PreferredProtocols uint32 `idl:"name:dwPreferredProtocols" json:"preferred_protocols"`
}

func (o *ConnectCommon) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ConnectCommon) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ShareMode); err != nil {
		return err
	}
	if err := w.WriteData(o.PreferredProtocols); err != nil {
		return err
	}
	return nil
}
func (o *ConnectCommon) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ShareMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.PreferredProtocols); err != nil {
		return err
	}
	return nil
}

// ConnectACall structure represents ConnectA_Call RPC structure.
//
// ConnectA_Call opens a connection to the smart card located in the reader identified
// by a reader name.
type ConnectACall struct {
	// szReader:  An ASCII string specifying the reader name to connect to.
	Reader string `idl:"name:szReader;string" json:"reader"`
	// Common:  Additional parameters that are required for the Connect call are specified
	// in section 3.1.4.28. For more information, see section 2.2.1.3.
	Common *ConnectCommon `idl:"name:Common" json:"common"`
}

func (o *ConnectACall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ConnectACall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Reader != "" {
		_ptr_szReader := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.Reader); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Reader, _ptr_szReader); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Common != nil {
		if err := o.Common.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ConnectCommon{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ConnectACall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_szReader := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.Reader); err != nil {
			return err
		}
		return nil
	})
	_s_szReader := func(ptr interface{}) { o.Reader = *ptr.(*string) }
	if err := w.ReadPointer(&o.Reader, _s_szReader, _ptr_szReader); err != nil {
		return err
	}
	if o.Common == nil {
		o.Common = &ConnectCommon{}
	}
	if err := o.Common.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ConnectWCall structure represents ConnectW_Call RPC structure.
//
// The ConnectW_Call structure is used to open a connection to the smart card located
// in the reader identified by a reader name.
type ConnectWCall struct {
	// szReader:  A Unicode string specifying the reader name to connect to.
	Reader string `idl:"name:szReader;string" json:"reader"`
	// Common:  Additional parameters that are required for the Connect call. For more information,
	// see sections 3.1.4.29 and 2.2.1.3.
	Common *ConnectCommon `idl:"name:Common" json:"common"`
}

func (o *ConnectWCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ConnectWCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Reader != "" {
		_ptr_szReader := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Reader); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Reader, _ptr_szReader); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Common != nil {
		if err := o.Common.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ConnectCommon{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ConnectWCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_szReader := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Reader); err != nil {
			return err
		}
		return nil
	})
	_s_szReader := func(ptr interface{}) { o.Reader = *ptr.(*string) }
	if err := w.ReadPointer(&o.Reader, _s_szReader, _ptr_szReader); err != nil {
		return err
	}
	if o.Common == nil {
		o.Common = &ConnectCommon{}
	}
	if err := o.Common.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ConnectReturn structure represents Connect_Return RPC structure.
//
// The Connect_Return structure is used to obtain return information from a Connect
// call (for more information, see sections 3.1.4.28 and 3.1.4.29).
type ConnectReturn struct {
	// ReturnCode:  HRESULT or Win32 Error code. Zero indicates success; any other value
	// indicates failure.
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
	// hCard:  A handle, as specified in section 2.2.1.2.
	Card *ReaderSmartCardHandle `idl:"name:hCard" json:"card"`
	// dwActiveProtocol:  A value that indicates the active smart card transmission protocol.
	// Possible values are specified in section 2.2.5.
	ActiveProtocol uint32 `idl:"name:dwActiveProtocol" json:"active_protocol"`
}

func (o *ConnectReturn) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ConnectReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if o.Card != nil {
		if err := o.Card.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardHandle{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ActiveProtocol); err != nil {
		return err
	}
	return nil
}
func (o *ConnectReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if o.Card == nil {
		o.Card = &ReaderSmartCardHandle{}
	}
	if err := o.Card.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ActiveProtocol); err != nil {
		return err
	}
	return nil
}

// ReconnectCall structure represents Reconnect_Call RPC structure.
//
// The Reconnect_Call structure is used to reopen a connection to the smart card associated
// with a valid context. For more information, see section 3.1.4.36.
type ReconnectCall struct {
	// hCard:  A handle, as specified in section 2.2.1.2.
	Card *ReaderSmartCardHandle `idl:"name:hCard" json:"card"`
	// dwShareMode:  A flag that indicates whether other applications can form connections
	// to this card. For acceptable values of this field, see section 2.2.6.
	ShareMode uint32 `idl:"name:dwShareMode" json:"share_mode"`
	// dwPreferredProtocols:  A bit mask of acceptable protocols for this connection. For
	// specifics on possible values, see section 2.2.5.
	PreferredProtocols uint32 `idl:"name:dwPreferredProtocols" json:"preferred_protocols"`
	// dwInitialization:  A type of initialization that SHOULD be performed on the card.
	//
	//	+-------------------------------+------------------------------------+
	//	|                               |                                    |
	//	|             VALUE             |              MEANING               |
	//	|                               |                                    |
	//	+-------------------------------+------------------------------------+
	//	+-------------------------------+------------------------------------+
	//	| SCARD_LEAVE_CARD 0x00000000   | Do not do anything.                |
	//	+-------------------------------+------------------------------------+
	//	| SCARD_RESET_CARD 0x00000001   | Reset the smart card.              |
	//	+-------------------------------+------------------------------------+
	//	| SCARD_UNPOWER_CARD 0x00000002 | Turn off and reset the smart card. |
	//	+-------------------------------+------------------------------------+
	Initialization uint32 `idl:"name:dwInitialization" json:"initialization"`
}

func (o *ReconnectCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReconnectCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Card != nil {
		if err := o.Card.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardHandle{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ShareMode); err != nil {
		return err
	}
	if err := w.WriteData(o.PreferredProtocols); err != nil {
		return err
	}
	if err := w.WriteData(o.Initialization); err != nil {
		return err
	}
	return nil
}
func (o *ReconnectCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Card == nil {
		o.Card = &ReaderSmartCardHandle{}
	}
	if err := o.Card.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ShareMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.PreferredProtocols); err != nil {
		return err
	}
	if err := w.ReadData(&o.Initialization); err != nil {
		return err
	}
	return nil
}

// ReconnectReturn structure represents Reconnect_Return RPC structure.
//
// The Reconnect_Return structure is used to obtain return information from a Reconnect
// call (for more information, see section 3.1.4.36).
type ReconnectReturn struct {
	// ReturnCode:  HRESULT or Win32 Error code. Zero indicates success; any other value
	// indicates failure.
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
	// dwActiveProtocol:  A flag that indicates the established active protocol. For more
	// information on acceptable values, see section 2.2.5 .
	ActiveProtocol uint32 `idl:"name:dwActiveProtocol" json:"active_protocol"`
}

func (o *ReconnectReturn) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReconnectReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.ActiveProtocol); err != nil {
		return err
	}
	return nil
}
func (o *ReconnectReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ActiveProtocol); err != nil {
		return err
	}
	return nil
}

// CardAndDispositionCall structure represents HCardAndDisposition_Call RPC structure.
//
// The HCardAndDisposition_Call structure defines the action taken on the disposition
// of a smart card associated with a valid context when a connection is terminated.
type CardAndDispositionCall struct {
	// hCard:  A handle, as specified in section 2.2.1.2.
	Card *ReaderSmartCardHandle `idl:"name:hCard" json:"card"`
	// dwDisposition:  The action to take on the card in the connected reader upon close.
	// This value is ignored on a BeginTransaction message call, as specified in section
	// 3.2.5.3.61.
	//
	//	+-------------------------------+------------------------------------+
	//	|                               |                                    |
	//	|             VALUE             |              MEANING               |
	//	|                               |                                    |
	//	+-------------------------------+------------------------------------+
	//	+-------------------------------+------------------------------------+
	//	| SCARD_LEAVE_CARD 0x00000000   | Do not do anything.                |
	//	+-------------------------------+------------------------------------+
	//	| SCARD_RESET_CARD 0x00000001   | Reset the smart card.              |
	//	+-------------------------------+------------------------------------+
	//	| SCARD_UNPOWER_CARD 0x00000002 | Turn off and reset the smart card. |
	//	+-------------------------------+------------------------------------+
	//	| SCARD_EJECT_CARD 0x00000003   | Eject the smart card.              |
	//	+-------------------------------+------------------------------------+
	Disposition uint32 `idl:"name:dwDisposition" json:"disposition"`
}

func (o *CardAndDispositionCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CardAndDispositionCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Card != nil {
		if err := o.Card.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardHandle{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Disposition); err != nil {
		return err
	}
	return nil
}
func (o *CardAndDispositionCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Card == nil {
		o.Card = &ReaderSmartCardHandle{}
	}
	if err := o.Card.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Disposition); err != nil {
		return err
	}
	return nil
}

// StateCall structure represents State_Call RPC structure.
//
// The State_Call structure defines parameters to the State call (as specified in section
// 3.1.4.40) for querying the contents of a smart card reader.
type StateCall struct {
	// hCard:  A handle, as specified in section 2.2.1.2.
	Card *ReaderSmartCardHandle `idl:"name:hCard" json:"card"`
	// fpbAtrIsNULL:  A Boolean value specifying whether the caller wants to retrieve the
	// length of the data. Set to FALSE (0x00000000) to allow the data to be returned. Set
	// to TRUE (0x00000001), and only the length of the data will be returned. SHOULD be
	// set to TRUE if cbAtrLen is set to SCARD_AUTOALLOCATE (0xFFFFFFFF).
	//
	//	+-------+------------+
	//	|       |            |
	//	| NAME  |   VALUE    |
	//	|       |            |
	//	+-------+------------+
	//	+-------+------------+
	//	| FALSE | 0x00000000 |
	//	+-------+------------+
	//	| TRUE  | 0x00000001 |
	//	+-------+------------+
	AttributeIsNull int32 `idl:"name:fpbAtrIsNULL" json:"attribute_is_null"`
	// cbAtrLen:  The length of the buffer specified on the TS server side. If cbAtrLen
	// is set to SCARD_AUTOALLOCATE with a value of 0xFFFFFFFF, an array of any length can
	// be returned. Otherwise, the returned array MUST NOT exceed cbAtrLen bytes in length.
	// When the array to be returned exceeds cbAtrLen bytes in length, State_Return.ReturnCode
	// MUST be set to SCARD_E_INSUFFICIENT_BUFFER (0x80100008). Also, cbAtrLen is ignored
	// if fpbAtrIsNULL is set to TRUE (0x00000001). If fpbAtrIsNULL is set to FALSE (0x00000000)
	// but cbAtrLen is set to 0x00000000, then the call MUST succeed, State_Return.cbAtrLen
	// MUST be set to the length of the data in bytes, and State_Return.rgAtr MUST be set
	// to NULL.
	AttributeLength uint32 `idl:"name:cbAtrLen" json:"attribute_length"`
}

func (o *StateCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StateCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Card != nil {
		if err := o.Card.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardHandle{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AttributeIsNull); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeLength); err != nil {
		return err
	}
	return nil
}
func (o *StateCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Card == nil {
		o.Card = &ReaderSmartCardHandle{}
	}
	if err := o.Card.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeIsNull); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeLength); err != nil {
		return err
	}
	return nil
}

// StateReturn structure represents State_Return RPC structure.
//
// The State_Return structure defines return information about the state of the smart
// card reader (for more information, see section 3.1.4.40).
type StateReturn struct {
	// ReturnCode:  HRESULT or Win32 Error code. Zero indicates success; any other value
	// indicates failure.
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
	// dwState:  The current state of the smart card in the Reader. Possible values are
	// specified in section 2.2.4.
	State uint32 `idl:"name:dwState" json:"state"`
	// dwProtocol:  The current protocol, if any. Possible values are specified in section
	// 2.2.5.
	Protocol uint32 `idl:"name:dwProtocol" json:"protocol"`
	// cbAtrLen:  The number of bytes in the rgAtr field.
	AttributeLength uint32 `idl:"name:cbAtrLen" json:"attribute_length"`
	// rgAtr:  A pointer to a buffer that receives the ATR string from the currently inserted
	// card, if available.
	Attribute []byte `idl:"name:rgAtr;size_is:(cbAtrLen);pointer:unique" json:"attribute"`
}

func (o *StateReturn) xxx_PreparePayload(ctx context.Context) error {
	if o.Attribute != nil && o.AttributeLength == 0 {
		o.AttributeLength = uint32(len(o.Attribute))
	}
	if o.AttributeLength > uint32(36) {
		return fmt.Errorf("AttributeLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StateReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeLength); err != nil {
		return err
	}
	if o.Attribute != nil || o.AttributeLength > 0 {
		_ptr_rgAtr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AttributeLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Attribute {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Attribute[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Attribute); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Attribute, _ptr_rgAtr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *StateReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeLength); err != nil {
		return err
	}
	_ptr_rgAtr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AttributeLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AttributeLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Attribute", sizeInfo[0])
		}
		o.Attribute = make([]byte, sizeInfo[0])
		for i1 := range o.Attribute {
			i1 := i1
			if err := w.ReadData(&o.Attribute[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgAtr := func(ptr interface{}) { o.Attribute = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Attribute, _s_rgAtr, _ptr_rgAtr); err != nil {
		return err
	}
	return nil
}

// StatusCall structure represents Status_Call RPC structure.
//
// Status_Call obtains the status of a connection for a valid smart card reader handle.
type StatusCall struct {
	// hCard:  A handle, as specified in section 2.2.1.2.
	Card *ReaderSmartCardHandle `idl:"name:hCard" json:"card"`
	// fmszReaderNamesIsNULL:  A Boolean value specifying whether the caller wants to retrieve
	// the length of the data. Set to FALSE (0x00000000) to allow the data to be returned.
	// Set to TRUE (0x00000001), and only the length of the data will be returned. Also,
	// cchReaderLen is ignored if this value is TRUE (0x00000001).
	//
	//	+-------+------------+
	//	|       |            |
	//	| NAME  |   VALUE    |
	//	|       |            |
	//	+-------+------------+
	//	+-------+------------+
	//	| FALSE | 0x00000000 |
	//	+-------+------------+
	//	| TRUE  | 0x00000001 |
	//	+-------+------------+
	ReaderNamesIsNull int32 `idl:"name:fmszReaderNamesIsNULL" json:"reader_names_is_null"`
	// cchReaderLen:  The length of the string buffer specified on the TS server side. If
	// cchReaderLen is set to SCARD_AUTOALLOCATE with a value of 0xFFFFFFFF, a string of
	// any length can be returned. Otherwise, the returned string MUST NOT exceed cchReaderLen
	// characters in length, including any null characters. When the string to be returned
	// exceeds cchReaderLen characters in length, including any null characters, Status_Return.ReturnCode
	// MUST be set to SCARD_E_INSUFFICIENT_BUFFER (0x80100008). The cchReaderLen field MUST
	// be ignored if fmszReaderNamesIsNULL is TRUE (0x00000001). Also, if fmszReaderNamesIsNULL
	// is set to FALSE (0x00000000) but cchReaderLen is set to 0x00000000, then the call
	// MUST succeed, Status_Return.cbAtrLen MUST be set to the length of the data in bytes,
	// and Status_Return.pbAtr MUST be set to NULL.
	ReaderLength uint32 `idl:"name:cchReaderLen" json:"reader_length"`
	// cbAtrLen:  Unused. MUST be ignored upon receipt.
	AttributeLength uint32 `idl:"name:cbAtrLen" json:"attribute_length"`
}

func (o *StatusCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StatusCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Card != nil {
		if err := o.Card.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardHandle{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ReaderNamesIsNull); err != nil {
		return err
	}
	if err := w.WriteData(o.ReaderLength); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeLength); err != nil {
		return err
	}
	return nil
}
func (o *StatusCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Card == nil {
		o.Card = &ReaderSmartCardHandle{}
	}
	if err := o.Card.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReaderNamesIsNull); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReaderLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeLength); err != nil {
		return err
	}
	return nil
}

// StatusReturn structure represents Status_Return RPC structure.
//
// The Status_Return structure defines return information about the status of the smart
// card reader (for more information, see sections 3.1.4.33 and 3.1.4.34).
type StatusReturn struct {
	// ReturnCode:  HRESULT or Win32 Error code. Zero indicates success; any other value
	// indicates failure.
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
	// cBytes:  The number of bytes in the mszReaderNames field.
	BytesCount uint32 `idl:"name:cBytes" json:"bytes_count"`
	// mszReaderNames:  A multistring containing the names that the reader is known by.
	// The value of this is dependent on the context (IOCTL) that it is used.
	//
	//	+--------------------------------+---------------------+
	//	|                                |                     |
	//	|             VALUE              |       MEANING       |
	//	|                                |                     |
	//	+--------------------------------+---------------------+
	//	+--------------------------------+---------------------+
	//	| SCARD_IOCTL_STATUSA 0x000900C8 | ASCII multistring   |
	//	+--------------------------------+---------------------+
	//	| SCARD_IOCTL_STATUSW 0x000900CC | Unicode multistring |
	//	+--------------------------------+---------------------+
	ReaderNames []byte `idl:"name:mszReaderNames;size_is:(cBytes);pointer:unique" json:"reader_names"`
	// dwState:  The current state of the smart card in the reader. Possible values are
	// specified in section 2.2.4.
	State uint32 `idl:"name:dwState" json:"state"`
	// dwProtocol:  The current protocol, if any. Possible values are specified in section
	// 2.2.5.
	Protocol uint32 `idl:"name:dwProtocol" json:"protocol"`
	// pbAtr:  A pointer to a buffer that receives the ATR string from the currently inserted
	// card, if available.
	Attribute []byte `idl:"name:pbAtr" json:"attribute"`
	// cbAtrLen:  The number of bytes in the ATR string.
	AttributeLength uint32 `idl:"name:cbAtrLen" json:"attribute_length"`
}

func (o *StatusReturn) xxx_PreparePayload(ctx context.Context) error {
	if o.ReaderNames != nil && o.BytesCount == 0 {
		o.BytesCount = uint32(len(o.ReaderNames))
	}
	if o.BytesCount > uint32(65536) {
		return fmt.Errorf("BytesCount is out of range")
	}
	if o.AttributeLength > uint32(32) {
		return fmt.Errorf("AttributeLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StatusReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesCount); err != nil {
		return err
	}
	if o.ReaderNames != nil || o.BytesCount > 0 {
		_ptr_mszReaderNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BytesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ReaderNames {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.ReaderNames[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.ReaderNames); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReaderNames, _ptr_mszReaderNames); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	for i1 := range o.Attribute {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(o.Attribute[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Attribute); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AttributeLength); err != nil {
		return err
	}
	return nil
}
func (o *StatusReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesCount); err != nil {
		return err
	}
	_ptr_mszReaderNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BytesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BytesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ReaderNames", sizeInfo[0])
		}
		o.ReaderNames = make([]byte, sizeInfo[0])
		for i1 := range o.ReaderNames {
			i1 := i1
			if err := w.ReadData(&o.ReaderNames[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_mszReaderNames := func(ptr interface{}) { o.ReaderNames = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ReaderNames, _s_mszReaderNames, _ptr_mszReaderNames); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	o.Attribute = make([]byte, 32)
	for i1 := range o.Attribute {
		i1 := i1
		if err := w.ReadData(&o.Attribute[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.AttributeLength); err != nil {
		return err
	}
	return nil
}

// SmartCardIORequest structure represents SCardIO_Request RPC structure.
//
// The SCardIO_Request structure represents the data to be prepended to a Transmit command
// (for more information, see section 3.1.4.35).
type SmartCardIORequest struct {
	// dwProtocol:  The protocol in use. Possible values are specified in section 2.2.5.
	Protocol uint32 `idl:"name:dwProtocol" json:"protocol"`
	// cbExtraBytes:  The number of bytes in the pbExtraBytes field.
	ExtraBytesLength uint32 `idl:"name:cbExtraBytes" json:"extra_bytes_length"`
	// pbExtraBytes:  Request data.
	ExtraBytes []byte `idl:"name:pbExtraBytes;size_is:(cbExtraBytes);pointer:unique" json:"extra_bytes"`
}

func (o *SmartCardIORequest) xxx_PreparePayload(ctx context.Context) error {
	if o.ExtraBytes != nil && o.ExtraBytesLength == 0 {
		o.ExtraBytesLength = uint32(len(o.ExtraBytes))
	}
	if o.ExtraBytesLength > uint32(1024) {
		return fmt.Errorf("ExtraBytesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SmartCardIORequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	if err := w.WriteData(o.ExtraBytesLength); err != nil {
		return err
	}
	if o.ExtraBytes != nil || o.ExtraBytesLength > 0 {
		_ptr_pbExtraBytes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ExtraBytesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ExtraBytes {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.ExtraBytes[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.ExtraBytes); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ExtraBytes, _ptr_pbExtraBytes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SmartCardIORequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExtraBytesLength); err != nil {
		return err
	}
	_ptr_pbExtraBytes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ExtraBytesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ExtraBytesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ExtraBytes", sizeInfo[0])
		}
		o.ExtraBytes = make([]byte, sizeInfo[0])
		for i1 := range o.ExtraBytes {
			i1 := i1
			if err := w.ReadData(&o.ExtraBytes[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbExtraBytes := func(ptr interface{}) { o.ExtraBytes = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ExtraBytes, _s_pbExtraBytes, _ptr_pbExtraBytes); err != nil {
		return err
	}
	return nil
}

// TransmitCall structure represents Transmit_Call RPC structure.
//
// The Transmit_Call structure is used to send data to the smart card associated with
// a valid context.
type TransmitCall struct {
	// hCard:  A handle, as specified in section 2.2.1.2.
	Card *ReaderSmartCardHandle `idl:"name:hCard" json:"card"`
	// ioSendPci:  A packet specifying input header information as specified in section
	// 2.2.1.8.
	SendPCI *SmartCardIORequest `idl:"name:ioSendPci" json:"send_pci"`
	// cbSendLength:  The length, in bytes, of the pbSendBuffer field.
	SendLength uint32 `idl:"name:cbSendLength" json:"send_length"`
	// pbSendBuffer:  The data to be written to the card. The format of the data is specific
	// to an individual card. For more information about data formats, see [ISO/IEC-7816-4]
	// sections 5 through 7.
	SendBuffer []byte `idl:"name:pbSendBuffer;size_is:(cbSendLength)" json:"send_buffer"`
	// pioRecvPci:  If non-NULL, this field is an SCardIO_Request packet that is set up
	// in the same way as the ioSendPci field and passed as the pioRecvPci parameter of
	// the Transmit call. If the value of this is NULL, the caller is not requesting the
	// pioRecvPci value to be returned.
	RecvPCI *SmartCardIORequest `idl:"name:pioRecvPci;pointer:unique" json:"recv_pci"`
	// fpbRecvBufferIsNULL:  A Boolean value specifying whether the caller wants to retrieve
	// the length of the data. MUST be set to TRUE (0x00000001) if the caller wants only
	// to retrieve the length of the data; otherwise, it MUST be set to FALSE (0x00000000).
	//
	//	+-------+------------+
	//	|       |            |
	//	| NAME  |   VALUE    |
	//	|       |            |
	//	+-------+------------+
	//	+-------+------------+
	//	| FALSE | 0x00000000 |
	//	+-------+------------+
	//	| TRUE  | 0x00000001 |
	//	+-------+------------+
	RecvBufferIsNull int32 `idl:"name:fpbRecvBufferIsNULL" json:"recv_buffer_is_null"`
	// cbRecvLength:  The maximum size of the buffer to be returned. MUST be ignored if
	// fpbRecvBufferIsNULL is set to TRUE (0x00000001).
	RecvLength uint32 `idl:"name:cbRecvLength" json:"recv_length"`
}

func (o *TransmitCall) xxx_PreparePayload(ctx context.Context) error {
	if o.SendBuffer != nil && o.SendLength == 0 {
		o.SendLength = uint32(len(o.SendBuffer))
	}
	if o.SendLength > uint32(66560) {
		return fmt.Errorf("SendLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransmitCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Card != nil {
		if err := o.Card.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardHandle{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SendPCI != nil {
		if err := o.SendPCI.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SmartCardIORequest{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SendLength); err != nil {
		return err
	}
	if o.SendBuffer != nil || o.SendLength > 0 {
		_ptr_pbSendBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SendLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SendBuffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.SendBuffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.SendBuffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SendBuffer, _ptr_pbSendBuffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RecvPCI != nil {
		_ptr_pioRecvPci := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecvPCI != nil {
				if err := o.RecvPCI.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SmartCardIORequest{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecvPCI, _ptr_pioRecvPci); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RecvBufferIsNull); err != nil {
		return err
	}
	if err := w.WriteData(o.RecvLength); err != nil {
		return err
	}
	return nil
}
func (o *TransmitCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Card == nil {
		o.Card = &ReaderSmartCardHandle{}
	}
	if err := o.Card.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SendPCI == nil {
		o.SendPCI = &SmartCardIORequest{}
	}
	if err := o.SendPCI.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.SendLength); err != nil {
		return err
	}
	_ptr_pbSendBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SendLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SendLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SendBuffer", sizeInfo[0])
		}
		o.SendBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.SendBuffer {
			i1 := i1
			if err := w.ReadData(&o.SendBuffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbSendBuffer := func(ptr interface{}) { o.SendBuffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SendBuffer, _s_pbSendBuffer, _ptr_pbSendBuffer); err != nil {
		return err
	}
	_ptr_pioRecvPci := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecvPCI == nil {
			o.RecvPCI = &SmartCardIORequest{}
		}
		if err := o.RecvPCI.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pioRecvPci := func(ptr interface{}) { o.RecvPCI = *ptr.(**SmartCardIORequest) }
	if err := w.ReadPointer(&o.RecvPCI, _s_pioRecvPci, _ptr_pioRecvPci); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecvBufferIsNull); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecvLength); err != nil {
		return err
	}
	return nil
}

// TransmitReturn structure represents Transmit_Return RPC structure.
//
// The Transmit_Return structure defines return information from a smart card after
// a Transmit call (for more information, see section 3.1.4.35).
type TransmitReturn struct {
	// ReturnCode:  HRESULT or Win32 Error code. Zero indicates success; any other value
	// indicates failure.
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
	// pioRecvPci:  The protocol header structure for the instruction, followed by a buffer
	// in which to receive any returned protocol control information (PCI) that is specific
	// to the protocol in use. If this field is NULL, a protocol header MUST NOT be returned.
	RecvPCI *SmartCardIORequest `idl:"name:pioRecvPci;pointer:unique" json:"recv_pci"`
	// cbRecvLength:  The size, in bytes, of the pbRecvBuffer field.
	RecvLength uint32 `idl:"name:cbRecvLength" json:"recv_length"`
	// pbRecvBuffer:  The data returned from the card.
	RecvBuffer []byte `idl:"name:pbRecvBuffer;size_is:(cbRecvLength);pointer:unique" json:"recv_buffer"`
}

func (o *TransmitReturn) xxx_PreparePayload(ctx context.Context) error {
	if o.RecvBuffer != nil && o.RecvLength == 0 {
		o.RecvLength = uint32(len(o.RecvBuffer))
	}
	if o.RecvLength > uint32(66560) {
		return fmt.Errorf("RecvLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransmitReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if o.RecvPCI != nil {
		_ptr_pioRecvPci := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecvPCI != nil {
				if err := o.RecvPCI.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SmartCardIORequest{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecvPCI, _ptr_pioRecvPci); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RecvLength); err != nil {
		return err
	}
	if o.RecvBuffer != nil || o.RecvLength > 0 {
		_ptr_pbRecvBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.RecvLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.RecvBuffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.RecvBuffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.RecvBuffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecvBuffer, _ptr_pbRecvBuffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransmitReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	_ptr_pioRecvPci := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecvPCI == nil {
			o.RecvPCI = &SmartCardIORequest{}
		}
		if err := o.RecvPCI.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pioRecvPci := func(ptr interface{}) { o.RecvPCI = *ptr.(**SmartCardIORequest) }
	if err := w.ReadPointer(&o.RecvPCI, _s_pioRecvPci, _ptr_pioRecvPci); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecvLength); err != nil {
		return err
	}
	_ptr_pbRecvBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.RecvLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.RecvLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.RecvBuffer", sizeInfo[0])
		}
		o.RecvBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.RecvBuffer {
			i1 := i1
			if err := w.ReadData(&o.RecvBuffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbRecvBuffer := func(ptr interface{}) { o.RecvBuffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.RecvBuffer, _s_pbRecvBuffer, _ptr_pbRecvBuffer); err != nil {
		return err
	}
	return nil
}

// GetTransmitCountCall structure represents GetTransmitCount_Call RPC structure.
//
// The GetTransmitCount_Call structure is used to obtain the number of transmit calls
// sent to the card since the reader was introduced.
type GetTransmitCountCall struct {
	// hCard:  A handle, as specified in section 2.2.1.2.
	Card *ReaderSmartCardHandle `idl:"name:hCard" json:"card"`
}

func (o *GetTransmitCountCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetTransmitCountCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Card != nil {
		if err := o.Card.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardHandle{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetTransmitCountCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Card == nil {
		o.Card = &ReaderSmartCardHandle{}
	}
	if err := o.Card.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// GetTransmitCountReturn structure represents GetTransmitCount_Return RPC structure.
//
// The GetTransmitCount_Return structure defines the number of transmit calls that were
// performed on the smart card reader (for more information, see section 3.1.4.41).
type GetTransmitCountReturn struct {
	// ReturnCode:  HRESULT or Win32 Error code. Zero indicates success; any other value
	// indicates failure.
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
	// cTransmitCount:  The field specifies the number of successful Transmit calls (for
	// more information, see section 3.1.4.35) performed on the reader since it was introduced
	// to the system.
	TransmitCount uint32 `idl:"name:cTransmitCount" json:"transmit_count"`
}

func (o *GetTransmitCountReturn) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetTransmitCountReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.TransmitCount); err != nil {
		return err
	}
	return nil
}
func (o *GetTransmitCountReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.TransmitCount); err != nil {
		return err
	}
	return nil
}

// ControlCall structure represents Control_Call RPC structure.
//
// Normally, communication is to the smart card via the reader. However, in some cases,
// the ability to communicate directly with the smart card reader is requested. The
// Control_Call structure provides the ability to talk to the reader.
type ControlCall struct {
	// hCard:  A handle, as specified in section 2.2.1.2.
	Card *ReaderSmartCardHandle `idl:"name:hCard" json:"card"`
	// dwControlCode:  The control code for the operation. These values are specific to
	// the hardware device. This protocol MUST NOT restrict or define any values for this
	// control codes.
	ControlCode uint32 `idl:"name:dwControlCode" json:"control_code"`
	// cbInBufferSize:  The size in bytes of the pvInBuffer field.
	InBufferLength uint32 `idl:"name:cbInBufferSize" json:"in_buffer_length"`
	// pvInBuffer:  A buffer that contains the data required to perform the operation. This
	// field SHOULD be NULL if the dwControlCode field specifies an operation that does
	// not require input data. Otherwise, this data is specific to the function being performed.
	InBuffer []byte `idl:"name:pvInBuffer;size_is:(cbInBufferSize);pointer:unique" json:"in_buffer"`
	// fpvOutBufferIsNULL:  A Boolean value specifying whether the caller wants to retrieve
	// the length of the data. MUST be set to TRUE (0x00000001) if the caller wants only
	// to retrieve the length of the data; otherwise, it MUST be set to FALSE (0x00000000).
	//
	//	+-------+------------+
	//	|       |            |
	//	| NAME  |   VALUE    |
	//	|       |            |
	//	+-------+------------+
	//	+-------+------------+
	//	| FALSE | 0x00000000 |
	//	+-------+------------+
	//	| TRUE  | 0x00000001 |
	//	+-------+------------+
	OutBufferIsNull int32 `idl:"name:fpvOutBufferIsNULL" json:"out_buffer_is_null"`
	// cbOutBufferSize:  The maximum size of the buffer to be returned. This field MUST
	// be ignored if fpvOutBufferIsNULL is set to TRUE (0x00000001).
	OutBufferLength uint32 `idl:"name:cbOutBufferSize" json:"out_buffer_length"`
}

func (o *ControlCall) xxx_PreparePayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferLength == 0 {
		o.InBufferLength = uint32(len(o.InBuffer))
	}
	if o.InBufferLength > uint32(66560) {
		return fmt.Errorf("InBufferLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ControlCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Card != nil {
		if err := o.Card.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardHandle{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ControlCode); err != nil {
		return err
	}
	if err := w.WriteData(o.InBufferLength); err != nil {
		return err
	}
	if o.InBuffer != nil || o.InBufferLength > 0 {
		_ptr_pvInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InBufferLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.InBuffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.InBuffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InBuffer, _ptr_pvInBuffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.OutBufferIsNull); err != nil {
		return err
	}
	if err := w.WriteData(o.OutBufferLength); err != nil {
		return err
	}
	return nil
}
func (o *ControlCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Card == nil {
		o.Card = &ReaderSmartCardHandle{}
	}
	if err := o.Card.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ControlCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.InBufferLength); err != nil {
		return err
	}
	_ptr_pvInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.InBufferLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.InBufferLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
		}
		o.InBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.InBuffer {
			i1 := i1
			if err := w.ReadData(&o.InBuffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pvInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.InBuffer, _s_pvInBuffer, _ptr_pvInBuffer); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutBufferIsNull); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutBufferLength); err != nil {
		return err
	}
	return nil
}

// ControlReturn structure represents Control_Return RPC structure.
//
// The Control_Return structure is used to obtain information from a Control_Call (for
// more information, see section 3.1.4.37).
type ControlReturn struct {
	// ReturnCode:  HRESULT or Win32 Error code. Zero indicates success; any other value
	// indicates failure.
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
	// cbOutBufferSize:  The number of bytes in the pvOutBuffer field.
	OutBufferLength uint32 `idl:"name:cbOutBufferSize" json:"out_buffer_length"`
	// pvOutBuffer:  Contains the return data specific to the value of the Control_Call
	// structure.
	OutBuffer []byte `idl:"name:pvOutBuffer;size_is:(cbOutBufferSize);pointer:unique" json:"out_buffer"`
}

func (o *ControlReturn) xxx_PreparePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.OutBufferLength == 0 {
		o.OutBufferLength = uint32(len(o.OutBuffer))
	}
	if o.OutBufferLength > uint32(66560) {
		return fmt.Errorf("OutBufferLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ControlReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.OutBufferLength); err != nil {
		return err
	}
	if o.OutBuffer != nil || o.OutBufferLength > 0 {
		_ptr_pvOutBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.OutBufferLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.OutBuffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.OutBuffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.OutBuffer, _ptr_pvOutBuffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ControlReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutBufferLength); err != nil {
		return err
	}
	_ptr_pvOutBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.OutBufferLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.OutBufferLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pvOutBuffer := func(ptr interface{}) { o.OutBuffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.OutBuffer, _s_pvOutBuffer, _ptr_pvOutBuffer); err != nil {
		return err
	}
	return nil
}

// GetAttributeCall structure represents GetAttrib_Call RPC structure.
//
// The GetAttrib_Call structure is used to read smart card reader attributes.
type GetAttributeCall struct {
	// hCard:  A handle, as specified in section 2.2.1.2.
	Card *ReaderSmartCardHandle `idl:"name:hCard" json:"card"`
	// dwAttrId:  An identifier for the attribute to get. For more information on defined
	// attributes, see [PCSC3] section 3.1.2.
	AttributeID uint32 `idl:"name:dwAttrId" json:"attribute_id"`
	// fpbAttrIsNULL:  A Boolean value specifying whether the caller wants to retrieve the
	// length of the data. Set to FALSE (0x00000000) in order to allow the data to be returned.
	// Set to TRUE (0x00000001) and only the length of the data will be returned.
	//
	//	+-------+------------+
	//	|       |            |
	//	| NAME  |   VALUE    |
	//	|       |            |
	//	+-------+------------+
	//	+-------+------------+
	//	| FALSE | 0x00000000 |
	//	+-------+------------+
	//	| TRUE  | 0x00000001 |
	//	+-------+------------+
	AttributeIsNull int32 `idl:"name:fpbAttrIsNULL" json:"attribute_is_null"`
	// cbAttrLen:  The length of the buffer specified on the TS Server side. If cbAttrLen
	// is set to SCARD_AUTOALLOCATE with a value of 0xFFFFFFFF then any buffer length can
	// be returned. Otherwise, the returned buffer MUST NOT exceed cbAttrLen bytes in length.
	// When the buffer to be returned exceeds cbAttrLen bytes in length, GetAttrib_Return.ReturnCode
	// MUST be set to SCARD_E_INSUFFICIENT_BUFFER (0x80100008). The cbAttrLen field MUST
	// be ignored if fpbAttrIsNULL is set to TRUE (0x00000001). Also, if fpbAttrIsNULL is
	// set to FALSE (0x00000000) but cbAttrLen is set to 0x00000000, then the call MUST
	// succeed, GetAttrib_Return.cbAttrLen MUST be set to the length of the data, in bytes,
	// and GetAttrib_Return.pbAttr MUST be set to NULL.
	AttributeLength uint32 `idl:"name:cbAttrLen" json:"attribute_length"`
}

func (o *GetAttributeCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetAttributeCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Card != nil {
		if err := o.Card.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardHandle{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AttributeID); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeIsNull); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeLength); err != nil {
		return err
	}
	return nil
}
func (o *GetAttributeCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Card == nil {
		o.Card = &ReaderSmartCardHandle{}
	}
	if err := o.Card.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeIsNull); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeLength); err != nil {
		return err
	}
	return nil
}

// GetAttributeReturn structure represents GetAttrib_Return RPC structure.
//
// The GetAttrib_Return structure defines attribute information from a smart card reader
// (for more information, see section 3.1.4.38).
type GetAttributeReturn struct {
	// ReturnCode:  HRESULT or Win32 Error code. Zero indicates success; any other value
	// indicates failure.
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
	// cbAttrLen:  The number of bytes in the pbAttr field.
	AttributeLength uint32 `idl:"name:cbAttrLen" json:"attribute_length"`
	// pbAttr:  A pointer to an array that contains any values returned from the corresponding
	// call.
	Attribute []byte `idl:"name:pbAttr;size_is:(cbAttrLen);pointer:unique" json:"attribute"`
}

func (o *GetAttributeReturn) xxx_PreparePayload(ctx context.Context) error {
	if o.Attribute != nil && o.AttributeLength == 0 {
		o.AttributeLength = uint32(len(o.Attribute))
	}
	if o.AttributeLength > uint32(65536) {
		return fmt.Errorf("AttributeLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetAttributeReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeLength); err != nil {
		return err
	}
	if o.Attribute != nil || o.AttributeLength > 0 {
		_ptr_pbAttr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AttributeLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Attribute {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Attribute[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Attribute); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Attribute, _ptr_pbAttr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetAttributeReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeLength); err != nil {
		return err
	}
	_ptr_pbAttr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AttributeLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AttributeLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Attribute", sizeInfo[0])
		}
		o.Attribute = make([]byte, sizeInfo[0])
		for i1 := range o.Attribute {
			i1 := i1
			if err := w.ReadData(&o.Attribute[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbAttr := func(ptr interface{}) { o.Attribute = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Attribute, _s_pbAttr, _ptr_pbAttr); err != nil {
		return err
	}
	return nil
}

// SetAttributeCall structure represents SetAttrib_Call RPC structure.
//
// The SetAttrib_Call structure allows users to set smart card reader attributes.
type SetAttributeCall struct {
	// hCard:  A handle, as specified in section 2.2.1.2.
	Card *ReaderSmartCardHandle `idl:"name:hCard" json:"card"`
	// dwAttrId:  The identifier of the attribute to set. The values are write-only. For
	// more information on possible values, see [PCSC3] section 3.1.2.
	AttributeID uint32 `idl:"name:dwAttrId" json:"attribute_id"`
	// cbAttrLen:  The size, in bytes, of the data corresponding to the pbAttr field.
	AttributeLength uint32 `idl:"name:cbAttrLen" json:"attribute_length"`
	// pbAttr:  A buffer that contains the attribute whose identifier is supplied in the
	// dwAttrId field. The format is specific to the value being set.
	Attribute []byte `idl:"name:pbAttr;size_is:(cbAttrLen)" json:"attribute"`
}

func (o *SetAttributeCall) xxx_PreparePayload(ctx context.Context) error {
	if o.Attribute != nil && o.AttributeLength == 0 {
		o.AttributeLength = uint32(len(o.Attribute))
	}
	if o.AttributeLength > uint32(65536) {
		return fmt.Errorf("AttributeLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SetAttributeCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Card != nil {
		if err := o.Card.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardHandle{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AttributeID); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeLength); err != nil {
		return err
	}
	if o.Attribute != nil || o.AttributeLength > 0 {
		_ptr_pbAttr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AttributeLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Attribute {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Attribute[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Attribute); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Attribute, _ptr_pbAttr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SetAttributeCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Card == nil {
		o.Card = &ReaderSmartCardHandle{}
	}
	if err := o.Card.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeLength); err != nil {
		return err
	}
	_ptr_pbAttr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AttributeLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AttributeLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Attribute", sizeInfo[0])
		}
		o.Attribute = make([]byte, sizeInfo[0])
		for i1 := range o.Attribute {
			i1 := i1
			if err := w.ReadData(&o.Attribute[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbAttr := func(ptr interface{}) { o.Attribute = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Attribute, _s_pbAttr, _ptr_pbAttr); err != nil {
		return err
	}
	return nil
}

// ReadCacheCommon structure represents ReadCache_Common RPC structure.
//
// The ReadCache_Common structure contains information common to both the ReadCacheA_Call
// and ReadCacheW_Call structures.
type ReadCacheCommon struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// CardIdentifier:  A UUID that specifies the name of the smart card with which the
	// name-value pair is associated.
	CardID *dtyp.UUID `idl:"name:CardIdentifier" json:"card_id"`
	// FreshnessCounter:  A value specifying the current revision of the data.
	FreshnessCounter uint32 `idl:"name:FreshnessCounter" json:"freshness_counter"`
	// fPbDataIsNULL:  A Boolean value specifying whether the caller wants to retrieve the
	// length of the data. It MUST be set to TRUE (0x00000001) if the caller wants only
	// to retrieve the length of the data; otherwise, it MUST be set to FALSE (0x00000000).
	DataIsNull int32 `idl:"name:fPbDataIsNULL" json:"data_is_null"`
	// cbDataLen:  The length of the buffer specified on the server side. If cbDataLen is
	// set to SCARD_AUTOALLOCATE with a value of 0xFFFFFFFF, a buffer of any length can
	// be returned. Otherwise, the returned buffer MUST NOT exceed cbDataLen bytes. This
	// field MUST be ignored if fPbDataIsNULL is set to TRUE (0x00000001).
	DataLength uint32 `idl:"name:cbDataLen" json:"data_length"`
}

func (o *ReadCacheCommon) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReadCacheCommon) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CardID != nil {
		_ptr_CardIdentifier := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CardID != nil {
				if err := o.CardID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.UUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CardID, _ptr_CardIdentifier); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.FreshnessCounter); err != nil {
		return err
	}
	if err := w.WriteData(o.DataIsNull); err != nil {
		return err
	}
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	return nil
}
func (o *ReadCacheCommon) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_CardIdentifier := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.CardID == nil {
			o.CardID = &dtyp.UUID{}
		}
		if err := o.CardID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_CardIdentifier := func(ptr interface{}) { o.CardID = *ptr.(**dtyp.UUID) }
	if err := w.ReadPointer(&o.CardID, _s_CardIdentifier, _ptr_CardIdentifier); err != nil {
		return err
	}
	if err := w.ReadData(&o.FreshnessCounter); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataIsNull); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	return nil
}

// ReadCacheACall structure represents ReadCacheA_Call RPC structure.
//
// The ReadCacheA_Call structure is used to obtain the card and reader information from
// the cache.
type ReadCacheACall struct {
	// szLookupName:  An ASCII string containing the lookup name.
	LookupName string `idl:"name:szLookupName;string" json:"lookup_name"`
	// Common:  Additional parameters for the Read Cache call (for additional information,
	// see section 3.1.4.42), as specified in section 2.2.1.9.
	Common *ReadCacheCommon `idl:"name:Common" json:"common"`
}

func (o *ReadCacheACall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReadCacheACall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LookupName != "" {
		_ptr_szLookupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.LookupName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LookupName, _ptr_szLookupName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Common != nil {
		if err := o.Common.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReadCacheCommon{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReadCacheACall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_szLookupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.LookupName); err != nil {
			return err
		}
		return nil
	})
	_s_szLookupName := func(ptr interface{}) { o.LookupName = *ptr.(*string) }
	if err := w.ReadPointer(&o.LookupName, _s_szLookupName, _ptr_szLookupName); err != nil {
		return err
	}
	if o.Common == nil {
		o.Common = &ReadCacheCommon{}
	}
	if err := o.Common.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ReadCacheWCall structure represents ReadCacheW_Call RPC structure.
//
// The ReadCacheW_Call structure is used to obtain the card and reader information from
// the cache.
type ReadCacheWCall struct {
	// szLookupName:  A Unicode string containing the lookup name.
	LookupName string `idl:"name:szLookupName;string" json:"lookup_name"`
	// Common:  Additional parameters for the Read Cache call (for additional information,
	// see section 3.1.4.43), as specified in section 2.2.1.9.
	Common *ReadCacheCommon `idl:"name:Common" json:"common"`
}

func (o *ReadCacheWCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReadCacheWCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LookupName != "" {
		_ptr_szLookupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LookupName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LookupName, _ptr_szLookupName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Common != nil {
		if err := o.Common.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReadCacheCommon{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReadCacheWCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_szLookupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LookupName); err != nil {
			return err
		}
		return nil
	})
	_s_szLookupName := func(ptr interface{}) { o.LookupName = *ptr.(*string) }
	if err := w.ReadPointer(&o.LookupName, _s_szLookupName, _ptr_szLookupName); err != nil {
		return err
	}
	if o.Common == nil {
		o.Common = &ReadCacheCommon{}
	}
	if err := o.Common.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ReadCacheReturn structure represents ReadCache_Return RPC structure.
//
// The ReadCache_Return structure is used to obtain the data that corresponds to the
// lookup item requested in ReadCacheA_Call as specified in section 2.2.2.25, or ReadCacheW_Call
// as specified in section 2.2.2.26. For more call information, see sections 3.1.4.42
// and 3.1.4.43.
type ReadCacheReturn struct {
	// ReturnCode:  HRESULT or Win32 Error codes. Zero indicates success; any other value
	// indicates failure.
	ReturnCode int32 `idl:"name:ReturnCode" json:"return_code"`
	// cbDataLen:  The number of bytes in the pbData field.
	DataLength uint32 `idl:"name:cbDataLen" json:"data_length"`
	// pbData:  The value of the look up item.
	Data []byte `idl:"name:pbData;size_is:(cbDataLen);pointer:unique" json:"data"`
}

func (o *ReadCacheReturn) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if o.DataLength > uint32(65536) {
		return fmt.Errorf("DataLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReadCacheReturn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	if o.Data != nil || o.DataLength > 0 {
		_ptr_pbData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DataLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Data {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Data[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Data, _ptr_pbData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReadCacheReturn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	_ptr_pbData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DataLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DataLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbData := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_pbData, _ptr_pbData); err != nil {
		return err
	}
	return nil
}

// WriteCacheCommon structure represents WriteCache_Common RPC structure.
//
// The WriteCache_Common structure contains information common between the WriteCacheA_Call
// and WriteCacheW_Call structures.
type WriteCacheCommon struct {
	// Context:  A valid context, as specified in section 2.2.1.1.
	Context *ReaderSmartCardContext `idl:"name:Context" json:"context"`
	// CardIdentifier:  A UUID that identifies the smart card with which the data SHOULD
	// be stored. CardIdentifier MUST be a unique value per the smart card.
	CardID *dtyp.UUID `idl:"name:CardIdentifier" json:"card_id"`
	// FreshnessCounter:  A value specifying the current revision of the data.
	FreshnessCounter uint32 `idl:"name:FreshnessCounter" json:"freshness_counter"`
	// cbDataLen:  The number of bytes in the pbData field.
	DataLength uint32 `idl:"name:cbDataLen" json:"data_length"`
	// pbData:  cbDataLen bytes of data to be stored.
	Data []byte `idl:"name:pbData;size_is:(cbDataLen);pointer:unique" json:"data"`
}

func (o *WriteCacheCommon) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if o.DataLength > uint32(65536) {
		return fmt.Errorf("DataLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WriteCacheCommon) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Context != nil {
		if err := o.Context.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReaderSmartCardContext{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CardID != nil {
		_ptr_CardIdentifier := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CardID != nil {
				if err := o.CardID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.UUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CardID, _ptr_CardIdentifier); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.FreshnessCounter); err != nil {
		return err
	}
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	if o.Data != nil || o.DataLength > 0 {
		_ptr_pbData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DataLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Data {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Data[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Data, _ptr_pbData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WriteCacheCommon) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Context == nil {
		o.Context = &ReaderSmartCardContext{}
	}
	if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_CardIdentifier := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.CardID == nil {
			o.CardID = &dtyp.UUID{}
		}
		if err := o.CardID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_CardIdentifier := func(ptr interface{}) { o.CardID = *ptr.(**dtyp.UUID) }
	if err := w.ReadPointer(&o.CardID, _s_CardIdentifier, _ptr_CardIdentifier); err != nil {
		return err
	}
	if err := w.ReadData(&o.FreshnessCounter); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	_ptr_pbData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DataLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DataLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbData := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_pbData, _ptr_pbData); err != nil {
		return err
	}
	return nil
}

// WriteCacheACall structure represents WriteCacheA_Call RPC structure.
//
// The WriteCacheA_Call structure is used to write the card and reader information to
// the cache.
type WriteCacheACall struct {
	// szLookupName:  An ASCII string containing the lookup name.
	LookupName string `idl:"name:szLookupName;string" json:"lookup_name"`
	// Common:  Additional parameters for the Write Cache call (for more information, see
	// section 3.1.4.44), as specified in section 2.2.1.10.
	Common *WriteCacheCommon `idl:"name:Common" json:"common"`
}

func (o *WriteCacheACall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WriteCacheACall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LookupName != "" {
		_ptr_szLookupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.LookupName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LookupName, _ptr_szLookupName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Common != nil {
		if err := o.Common.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&WriteCacheCommon{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *WriteCacheACall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_szLookupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.LookupName); err != nil {
			return err
		}
		return nil
	})
	_s_szLookupName := func(ptr interface{}) { o.LookupName = *ptr.(*string) }
	if err := w.ReadPointer(&o.LookupName, _s_szLookupName, _ptr_szLookupName); err != nil {
		return err
	}
	if o.Common == nil {
		o.Common = &WriteCacheCommon{}
	}
	if err := o.Common.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// WriteCacheWCall structure represents WriteCacheW_Call RPC structure.
//
// The WriteCacheW_Call structure is used to write the card and reader information to
// the cache.
type WriteCacheWCall struct {
	// szLookupName:  An Unicode string containing the lookup name.
	LookupName string `idl:"name:szLookupName;string" json:"lookup_name"`
	// Common:  Additional parameters for the Write Cache call (for more information, see
	// section 2.2.1.10.
	Common *WriteCacheCommon `idl:"name:Common" json:"common"`
}

func (o *WriteCacheWCall) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WriteCacheWCall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LookupName != "" {
		_ptr_szLookupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LookupName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LookupName, _ptr_szLookupName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Common != nil {
		if err := o.Common.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&WriteCacheCommon{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *WriteCacheWCall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_szLookupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LookupName); err != nil {
			return err
		}
		return nil
	})
	_s_szLookupName := func(ptr interface{}) { o.LookupName = *ptr.(*string) }
	if err := w.ReadPointer(&o.LookupName, _s_szLookupName, _ptr_szLookupName); err != nil {
		return err
	}
	if o.Common == nil {
		o.Common = &WriteCacheCommon{}
	}
	if err := o.Common.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultTypeSmartCardPackClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultTypeSmartCardPackClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTypeSmartCardPackClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewTypeSmartCardPackClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TypeSmartCardPackClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TypeSmartCardPackSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultTypeSmartCardPackClient{cc: cc}, nil
}
