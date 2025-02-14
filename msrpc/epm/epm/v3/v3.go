package epm

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
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
	_ = dcetypes.GoPackage
)

var (
	// import guard
	GoPackage = "epm"
)

var (
	// Syntax UUID
	EpmSyntaxUUID = &uuid.UUID{TimeLow: 0xe1af8308, TimeMid: 0x5d1f, TimeHiAndVersion: 0x11c9, ClockSeqHiAndReserved: 0x91, ClockSeqLow: 0xa4, Node: [6]uint8{0x8, 0x0, 0x2b, 0x14, 0xa0, 0xfa}}
	// Syntax ID
	EpmSyntaxV3_0 = &dcerpc.SyntaxID{IfUUID: EpmSyntaxUUID, IfVersionMajor: 3, IfVersionMinor: 0}
)

// epm interface.
type EpmClient interface {

	// ept_insert operation.
	Insert(context.Context, *InsertRequest, ...dcerpc.CallOption) (*InsertResponse, error)

	// ept_delete operation.
	Delete(context.Context, *DeleteRequest, ...dcerpc.CallOption) (*DeleteResponse, error)

	// ept_lookup operation.
	Lookup(context.Context, *LookupRequest, ...dcerpc.CallOption) (*LookupResponse, error)

	// ept_map operation.
	Map(context.Context, *MapRequest, ...dcerpc.CallOption) (*MapResponse, error)

	// ept_lookup_handle_free operation.
	LookupHandleFree(context.Context, *LookupHandleFreeRequest, ...dcerpc.CallOption) (*LookupHandleFreeResponse, error)

	// ept_inq_object operation.
	InquireObject(context.Context, *InquireObjectRequest, ...dcerpc.CallOption) (*InquireObjectResponse, error)

	// ept_mgmt_delete operation.
	ManagementDelete(context.Context, *ManagementDeleteRequest, ...dcerpc.CallOption) (*ManagementDeleteResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// MaxAnnotationSize represents the ept_max_annotation_size RPC constant
var MaxAnnotationSize = 64

// Entry structure represents ept_entry_t RPC structure.
type Entry struct {
	Object     *dtyp.GUID      `idl:"name:object" json:"object"`
	Tower      *dcetypes.Tower `idl:"name:tower" json:"tower"`
	Annotation string          `idl:"name:annotation;string" json:"annotation"`
}

func (o *Entry) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Entry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Object != nil {
		if err := o.Object.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Tower != nil {
		_ptr_tower := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Tower != nil {
				if err := o.Tower.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dcetypes.Tower{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Tower, _ptr_tower); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	sizeInfo := []uint64{
		0,
	}
	dimLength1 := ndr.CharNLen(o.Annotation)
	sizeInfo[0] = dimLength1
	if err := w.WriteSize(0); err != nil {
		return err
	}
	if err := w.WriteSize(dimLength1); err != nil {
		return err
	}
	_Annotation_buf := []byte(o.Annotation)
	if uint64(len(_Annotation_buf)) > 64-1 {
		_Annotation_buf = _Annotation_buf[:64-1]
	}
	if o.Annotation != ndr.ZeroString {
		_Annotation_buf = append(_Annotation_buf, byte(0))
	}
	for i1 := range _Annotation_buf {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(_Annotation_buf[i1]); err != nil {
			return err
		}
	}
	return nil
}
func (o *Entry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Object == nil {
		o.Object = &dtyp.GUID{}
	}
	if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_tower := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Tower == nil {
			o.Tower = &dcetypes.Tower{}
		}
		if err := o.Tower.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_tower := func(ptr interface{}) { o.Tower = *ptr.(**dcetypes.Tower) }
	if err := w.ReadPointer(&o.Tower, _s_tower, _ptr_tower); err != nil {
		return err
	}
	sizeInfo := []uint64{
		0,
	}
	for sz1 := range sizeInfo {
		if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
			return err
		}
		if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
			return err
		}
	}
	var _Annotation_buf []byte
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array _Annotation_buf", sizeInfo[0])
	}
	_Annotation_buf = make([]byte, sizeInfo[0])
	for i1 := range _Annotation_buf {
		i1 := i1
		if err := w.ReadData(&_Annotation_buf[i1]); err != nil {
			return err
		}
	}
	o.Annotation = strings.TrimRight(string(_Annotation_buf), ndr.ZeroString)
	return nil
}

// LookupHandle structure represents ept_lookup_handle_t RPC structure.
type LookupHandle dcetypes.ContextHandle

func (o *LookupHandle) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *LookupHandle) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LookupHandle) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *LookupHandle) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultEpmClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultEpmClient) Insert(ctx context.Context, in *InsertRequest, opts ...dcerpc.CallOption) (*InsertResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InsertResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultEpmClient) Delete(ctx context.Context, in *DeleteRequest, opts ...dcerpc.CallOption) (*DeleteResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultEpmClient) Lookup(ctx context.Context, in *LookupRequest, opts ...dcerpc.CallOption) (*LookupResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultEpmClient) Map(ctx context.Context, in *MapRequest, opts ...dcerpc.CallOption) (*MapResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MapResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultEpmClient) LookupHandleFree(ctx context.Context, in *LookupHandleFreeRequest, opts ...dcerpc.CallOption) (*LookupHandleFreeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupHandleFreeResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultEpmClient) InquireObject(ctx context.Context, in *InquireObjectRequest, opts ...dcerpc.CallOption) (*InquireObjectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InquireObjectResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultEpmClient) ManagementDelete(ctx context.Context, in *ManagementDeleteRequest, opts ...dcerpc.CallOption) (*ManagementDeleteResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ManagementDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultEpmClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEpmClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewEpmClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EpmClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EpmSyntaxV3_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultEpmClient{cc: cc}, nil
}

// xxx_InsertOperation structure represents the ept_insert operation
type xxx_InsertOperation struct {
	EntriesLength uint32   `idl:"name:num_ents" json:"entries_length"`
	Entries       []*Entry `idl:"name:entries;size_is:(num_ents)" json:"entries"`
	Replace       bool     `idl:"name:replace" json:"replace"`
	Status        uint32   `idl:"name:status" json:"status"`
}

func (o *xxx_InsertOperation) OpNum() int { return 0 }

func (o *xxx_InsertOperation) OpName() string { return "/epm/v3/ept_insert" }

func (o *xxx_InsertOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Entries != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.Entries))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InsertOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// num_ents {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.EntriesLength); err != nil {
			return err
		}
	}
	// entries {in} (1:[dim:0,size_is=num_ents])(2:{alias=ept_entry_t}(struct))
	{
		dimSize1 := uint64(o.EntriesLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Entries {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Entries[i1] != nil {
				if err := o.Entries[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Entries); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// replace {in} (1:{alias=boolean32, names=unsigned32}(uint32))
	{
		if !o.Replace {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(uint32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_InsertOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// num_ents {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.EntriesLength); err != nil {
			return err
		}
	}
	// entries {in} (1:[dim:0,size_is=num_ents])(2:{alias=ept_entry_t}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Entries", sizeInfo[0])
		}
		o.Entries = make([]*Entry, sizeInfo[0])
		for i1 := range o.Entries {
			i1 := i1
			if o.Entries[i1] == nil {
				o.Entries[i1] = &Entry{}
			}
			if err := o.Entries[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// replace {in} (1:{alias=boolean32, names=unsigned32}(uint32))
	{
		var _bReplace uint32
		if err := w.ReadData(&_bReplace); err != nil {
			return err
		}
		o.Replace = _bReplace != 0
	}
	return nil
}

func (o *xxx_InsertOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InsertOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InsertOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// InsertRequest structure represents the ept_insert operation request
type InsertRequest struct {
	EntriesLength uint32   `idl:"name:num_ents" json:"entries_length"`
	Entries       []*Entry `idl:"name:entries;size_is:(num_ents)" json:"entries"`
	Replace       bool     `idl:"name:replace" json:"replace"`
}

func (o *InsertRequest) xxx_ToOp(ctx context.Context, op *xxx_InsertOperation) *xxx_InsertOperation {
	if op == nil {
		op = &xxx_InsertOperation{}
	}
	if o == nil {
		return op
	}
	op.EntriesLength = o.EntriesLength
	op.Entries = o.Entries
	op.Replace = o.Replace
	return op
}

func (o *InsertRequest) xxx_FromOp(ctx context.Context, op *xxx_InsertOperation) {
	if o == nil {
		return
	}
	o.EntriesLength = op.EntriesLength
	o.Entries = op.Entries
	o.Replace = op.Replace
}
func (o *InsertRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InsertRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InsertOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InsertResponse structure represents the ept_insert operation response
type InsertResponse struct {
	Status uint32 `idl:"name:status" json:"status"`
}

func (o *InsertResponse) xxx_ToOp(ctx context.Context, op *xxx_InsertOperation) *xxx_InsertOperation {
	if op == nil {
		op = &xxx_InsertOperation{}
	}
	if o == nil {
		return op
	}
	op.Status = o.Status
	return op
}

func (o *InsertResponse) xxx_FromOp(ctx context.Context, op *xxx_InsertOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
}
func (o *InsertResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InsertResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InsertOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteOperation structure represents the ept_delete operation
type xxx_DeleteOperation struct {
	EntriesLength uint32   `idl:"name:num_ents" json:"entries_length"`
	Entries       []*Entry `idl:"name:entries;size_is:(num_ents)" json:"entries"`
	Status        uint32   `idl:"name:status" json:"status"`
}

func (o *xxx_DeleteOperation) OpNum() int { return 1 }

func (o *xxx_DeleteOperation) OpName() string { return "/epm/v3/ept_delete" }

func (o *xxx_DeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Entries != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.Entries))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// num_ents {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.EntriesLength); err != nil {
			return err
		}
	}
	// entries {in} (1:[dim:0,size_is=num_ents])(2:{alias=ept_entry_t}(struct))
	{
		dimSize1 := uint64(o.EntriesLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Entries {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Entries[i1] != nil {
				if err := o.Entries[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Entries); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// num_ents {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.EntriesLength); err != nil {
			return err
		}
	}
	// entries {in} (1:[dim:0,size_is=num_ents])(2:{alias=ept_entry_t}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Entries", sizeInfo[0])
		}
		o.Entries = make([]*Entry, sizeInfo[0])
		for i1 := range o.Entries {
			i1 := i1
			if o.Entries[i1] == nil {
				o.Entries[i1] = &Entry{}
			}
			if err := o.Entries[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// DeleteRequest structure represents the ept_delete operation request
type DeleteRequest struct {
	EntriesLength uint32   `idl:"name:num_ents" json:"entries_length"`
	Entries       []*Entry `idl:"name:entries;size_is:(num_ents)" json:"entries"`
}

func (o *DeleteRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteOperation) *xxx_DeleteOperation {
	if op == nil {
		op = &xxx_DeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.EntriesLength = o.EntriesLength
	op.Entries = o.Entries
	return op
}

func (o *DeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteOperation) {
	if o == nil {
		return
	}
	o.EntriesLength = op.EntriesLength
	o.Entries = op.Entries
}
func (o *DeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteResponse structure represents the ept_delete operation response
type DeleteResponse struct {
	Status uint32 `idl:"name:status" json:"status"`
}

func (o *DeleteResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteOperation) *xxx_DeleteOperation {
	if op == nil {
		op = &xxx_DeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.Status = o.Status
	return op
}

func (o *DeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
}
func (o *DeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupOperation structure represents the ept_lookup operation
type xxx_LookupOperation struct {
	InquiryType   uint32                `idl:"name:inquiry_type" json:"inquiry_type"`
	Object        *dtyp.GUID            `idl:"name:object" json:"object"`
	InterfaceID   *dcetypes.InterfaceID `idl:"name:interface_id" json:"interface_id"`
	VersOption    uint32                `idl:"name:vers_option" json:"vers_option"`
	EntryHandle   *LookupHandle         `idl:"name:entry_handle" json:"entry_handle"`
	MaxEntries    uint32                `idl:"name:max_ents" json:"max_entries"`
	EntriesLength uint32                `idl:"name:num_ents" json:"entries_length"`
	Entries       []*Entry              `idl:"name:entries;size_is:(max_ents);length_is:(num_ents)" json:"entries"`
	Status        uint32                `idl:"name:status" json:"status"`
}

func (o *xxx_LookupOperation) OpNum() int { return 2 }

func (o *xxx_LookupOperation) OpName() string { return "/epm/v3/ept_lookup" }

func (o *xxx_LookupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// inquiry_type {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.InquiryType); err != nil {
			return err
		}
	}
	// object {in} (1:{pointer=ptr, alias=uuid_p_t}*(1))(2:{alias=GUID}(struct))
	{
		if o.Object != nil {
			_ptr_object := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Object != nil {
					if err := o.Object.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Object, _ptr_object); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// interface_id {in} (1:{pointer=unique, alias=rpc_if_id_p_t}*(1))(2:{alias=rpc_if_id_t}(struct))
	{
		if o.InterfaceID != nil {
			_ptr_interface_id := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InterfaceID != nil {
					if err := o.InterfaceID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcetypes.InterfaceID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InterfaceID, _ptr_interface_id); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// vers_option {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.VersOption); err != nil {
			return err
		}
	}
	// entry_handle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ept_lookup_handle_t, names=ndr_context_handle}(struct))
	{
		if o.EntryHandle != nil {
			if err := o.EntryHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&LookupHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// max_ents {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.MaxEntries); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// inquiry_type {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.InquiryType); err != nil {
			return err
		}
	}
	// object {in} (1:{pointer=ptr, alias=uuid_p_t}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_object := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Object == nil {
				o.Object = &dtyp.GUID{}
			}
			if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_object := func(ptr interface{}) { o.Object = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.Object, _s_object, _ptr_object); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// interface_id {in} (1:{pointer=unique, alias=rpc_if_id_p_t}*(1))(2:{alias=rpc_if_id_t}(struct))
	{
		_ptr_interface_id := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InterfaceID == nil {
				o.InterfaceID = &dcetypes.InterfaceID{}
			}
			if err := o.InterfaceID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_interface_id := func(ptr interface{}) { o.InterfaceID = *ptr.(**dcetypes.InterfaceID) }
		if err := w.ReadPointer(&o.InterfaceID, _s_interface_id, _ptr_interface_id); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// vers_option {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.VersOption); err != nil {
			return err
		}
	}
	// entry_handle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ept_lookup_handle_t, names=ndr_context_handle}(struct))
	{
		if o.EntryHandle == nil {
			o.EntryHandle = &LookupHandle{}
		}
		if err := o.EntryHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// max_ents {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.MaxEntries); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Entries != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.Entries))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// entry_handle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ept_lookup_handle_t, names=ndr_context_handle}(struct))
	{
		if o.EntryHandle != nil {
			if err := o.EntryHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&LookupHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// num_ents {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.EntriesLength); err != nil {
			return err
		}
	}
	// entries {out} (1:[dim:0,size_is=max_ents,length_is=num_ents])(2:{alias=ept_entry_t}(struct))
	{
		dimSize1 := uint64(o.MaxEntries)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.EntriesLength)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.Entries {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Entries[i1] != nil {
				if err := o.Entries[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Entries); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// entry_handle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ept_lookup_handle_t, names=ndr_context_handle}(struct))
	{
		if o.EntryHandle == nil {
			o.EntryHandle = &LookupHandle{}
		}
		if err := o.EntryHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// num_ents {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.EntriesLength); err != nil {
			return err
		}
	}
	// entries {out} (1:[dim:0,size_is=max_ents,length_is=num_ents])(2:{alias=ept_entry_t}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Entries", sizeInfo[0])
		}
		o.Entries = make([]*Entry, sizeInfo[0])
		for i1 := range o.Entries {
			i1 := i1
			if o.Entries[i1] == nil {
				o.Entries[i1] = &Entry{}
			}
			if err := o.Entries[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// LookupRequest structure represents the ept_lookup operation request
type LookupRequest struct {
	InquiryType uint32                `idl:"name:inquiry_type" json:"inquiry_type"`
	Object      *dtyp.GUID            `idl:"name:object" json:"object"`
	InterfaceID *dcetypes.InterfaceID `idl:"name:interface_id" json:"interface_id"`
	VersOption  uint32                `idl:"name:vers_option" json:"vers_option"`
	EntryHandle *LookupHandle         `idl:"name:entry_handle" json:"entry_handle"`
	MaxEntries  uint32                `idl:"name:max_ents" json:"max_entries"`
}

func (o *LookupRequest) xxx_ToOp(ctx context.Context, op *xxx_LookupOperation) *xxx_LookupOperation {
	if op == nil {
		op = &xxx_LookupOperation{}
	}
	if o == nil {
		return op
	}
	op.InquiryType = o.InquiryType
	op.Object = o.Object
	op.InterfaceID = o.InterfaceID
	op.VersOption = o.VersOption
	op.EntryHandle = o.EntryHandle
	op.MaxEntries = o.MaxEntries
	return op
}

func (o *LookupRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupOperation) {
	if o == nil {
		return
	}
	o.InquiryType = op.InquiryType
	o.Object = op.Object
	o.InterfaceID = op.InterfaceID
	o.VersOption = op.VersOption
	o.EntryHandle = op.EntryHandle
	o.MaxEntries = op.MaxEntries
}
func (o *LookupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupResponse structure represents the ept_lookup operation response
type LookupResponse struct {
	EntryHandle   *LookupHandle `idl:"name:entry_handle" json:"entry_handle"`
	EntriesLength uint32        `idl:"name:num_ents" json:"entries_length"`
	Entries       []*Entry      `idl:"name:entries;size_is:(max_ents);length_is:(num_ents)" json:"entries"`
	Status        uint32        `idl:"name:status" json:"status"`
}

func (o *LookupResponse) xxx_ToOp(ctx context.Context, op *xxx_LookupOperation) *xxx_LookupOperation {
	if op == nil {
		op = &xxx_LookupOperation{}
	}
	if o == nil {
		return op
	}
	op.EntryHandle = o.EntryHandle
	op.EntriesLength = o.EntriesLength
	op.Entries = o.Entries
	op.Status = o.Status
	return op
}

func (o *LookupResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupOperation) {
	if o == nil {
		return
	}
	o.EntryHandle = op.EntryHandle
	o.EntriesLength = op.EntriesLength
	o.Entries = op.Entries
	o.Status = op.Status
}
func (o *LookupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MapOperation structure represents the ept_map operation
type xxx_MapOperation struct {
	Object       *dtyp.GUID        `idl:"name:object" json:"object"`
	MapTower     *dcetypes.Tower   `idl:"name:map_tower" json:"map_tower"`
	EntryHandle  *LookupHandle     `idl:"name:entry_handle" json:"entry_handle"`
	MaxTowers    uint32            `idl:"name:max_towers" json:"max_towers"`
	TowersLength uint32            `idl:"name:num_towers" json:"towers_length"`
	Towers       []*dcetypes.Tower `idl:"name:towers;size_is:(max_towers);length_is:(num_towers)" json:"towers"`
	Status       uint32            `idl:"name:status" json:"status"`
}

func (o *xxx_MapOperation) OpNum() int { return 3 }

func (o *xxx_MapOperation) OpName() string { return "/epm/v3/ept_map" }

func (o *xxx_MapOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MapOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// object {in} (1:{pointer=ptr, alias=uuid_p_t}*(1))(2:{alias=GUID}(struct))
	{
		if o.Object != nil {
			_ptr_object := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Object != nil {
					if err := o.Object.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Object, _ptr_object); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// map_tower {in} (1:{pointer=ptr, alias=twr_p_t}*(1))(2:{alias=twr_t}(struct))
	{
		if o.MapTower != nil {
			_ptr_map_tower := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MapTower != nil {
					if err := o.MapTower.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcetypes.Tower{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MapTower, _ptr_map_tower); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// entry_handle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ept_lookup_handle_t, names=ndr_context_handle}(struct))
	{
		if o.EntryHandle != nil {
			if err := o.EntryHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&LookupHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// max_towers {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.MaxTowers); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MapOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// object {in} (1:{pointer=ptr, alias=uuid_p_t}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_object := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Object == nil {
				o.Object = &dtyp.GUID{}
			}
			if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_object := func(ptr interface{}) { o.Object = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.Object, _s_object, _ptr_object); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// map_tower {in} (1:{pointer=ptr, alias=twr_p_t}*(1))(2:{alias=twr_t}(struct))
	{
		_ptr_map_tower := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MapTower == nil {
				o.MapTower = &dcetypes.Tower{}
			}
			if err := o.MapTower.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_map_tower := func(ptr interface{}) { o.MapTower = *ptr.(**dcetypes.Tower) }
		if err := w.ReadPointer(&o.MapTower, _s_map_tower, _ptr_map_tower); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// entry_handle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ept_lookup_handle_t, names=ndr_context_handle}(struct))
	{
		if o.EntryHandle == nil {
			o.EntryHandle = &LookupHandle{}
		}
		if err := o.EntryHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// max_towers {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.MaxTowers); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MapOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Towers != nil && o.TowersLength == 0 {
		o.TowersLength = uint32(len(o.Towers))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MapOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// entry_handle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ept_lookup_handle_t, names=ndr_context_handle}(struct))
	{
		if o.EntryHandle != nil {
			if err := o.EntryHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&LookupHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// num_towers {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.TowersLength); err != nil {
			return err
		}
	}
	// towers {out} (1:{pointer=ref}*(1))(2:{pointer=ptr, alias=twr_p_t}[dim:0,size_is=max_towers,length_is=num_towers]*(1))(3:{alias=twr_t}(struct))
	{
		dimSize1 := uint64(o.MaxTowers)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.TowersLength)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.Towers {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Towers[i1] != nil {
				_ptr_towers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.Towers[i1] != nil {
						if err := o.Towers[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dcetypes.Tower{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Towers[i1], _ptr_towers); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Towers); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MapOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// entry_handle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ept_lookup_handle_t, names=ndr_context_handle}(struct))
	{
		if o.EntryHandle == nil {
			o.EntryHandle = &LookupHandle{}
		}
		if err := o.EntryHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// num_towers {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.TowersLength); err != nil {
			return err
		}
	}
	// towers {out} (1:{pointer=ref}*(1))(2:{pointer=ptr, alias=twr_p_t}[dim:0,size_is=max_towers,length_is=num_towers]*(1))(3:{alias=twr_t}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Towers", sizeInfo[0])
		}
		o.Towers = make([]*dcetypes.Tower, sizeInfo[0])
		for i1 := range o.Towers {
			i1 := i1
			_ptr_towers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Towers[i1] == nil {
					o.Towers[i1] = &dcetypes.Tower{}
				}
				if err := o.Towers[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_towers := func(ptr interface{}) { o.Towers[i1] = *ptr.(**dcetypes.Tower) }
			if err := w.ReadPointer(&o.Towers[i1], _s_towers, _ptr_towers); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// MapRequest structure represents the ept_map operation request
type MapRequest struct {
	Object      *dtyp.GUID      `idl:"name:object" json:"object"`
	MapTower    *dcetypes.Tower `idl:"name:map_tower" json:"map_tower"`
	EntryHandle *LookupHandle   `idl:"name:entry_handle" json:"entry_handle"`
	MaxTowers   uint32          `idl:"name:max_towers" json:"max_towers"`
}

func (o *MapRequest) xxx_ToOp(ctx context.Context, op *xxx_MapOperation) *xxx_MapOperation {
	if op == nil {
		op = &xxx_MapOperation{}
	}
	if o == nil {
		return op
	}
	op.Object = o.Object
	op.MapTower = o.MapTower
	op.EntryHandle = o.EntryHandle
	op.MaxTowers = o.MaxTowers
	return op
}

func (o *MapRequest) xxx_FromOp(ctx context.Context, op *xxx_MapOperation) {
	if o == nil {
		return
	}
	o.Object = op.Object
	o.MapTower = op.MapTower
	o.EntryHandle = op.EntryHandle
	o.MaxTowers = op.MaxTowers
}
func (o *MapRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MapRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MapOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MapResponse structure represents the ept_map operation response
type MapResponse struct {
	EntryHandle  *LookupHandle     `idl:"name:entry_handle" json:"entry_handle"`
	TowersLength uint32            `idl:"name:num_towers" json:"towers_length"`
	Towers       []*dcetypes.Tower `idl:"name:towers;size_is:(max_towers);length_is:(num_towers)" json:"towers"`
	Status       uint32            `idl:"name:status" json:"status"`
}

func (o *MapResponse) xxx_ToOp(ctx context.Context, op *xxx_MapOperation) *xxx_MapOperation {
	if op == nil {
		op = &xxx_MapOperation{}
	}
	if o == nil {
		return op
	}
	op.EntryHandle = o.EntryHandle
	op.TowersLength = o.TowersLength
	op.Towers = o.Towers
	op.Status = o.Status
	return op
}

func (o *MapResponse) xxx_FromOp(ctx context.Context, op *xxx_MapOperation) {
	if o == nil {
		return
	}
	o.EntryHandle = op.EntryHandle
	o.TowersLength = op.TowersLength
	o.Towers = op.Towers
	o.Status = op.Status
}
func (o *MapResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MapResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MapOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupHandleFreeOperation structure represents the ept_lookup_handle_free operation
type xxx_LookupHandleFreeOperation struct {
	EntryHandle *LookupHandle `idl:"name:entry_handle" json:"entry_handle"`
	Status      uint32        `idl:"name:status" json:"status"`
}

func (o *xxx_LookupHandleFreeOperation) OpNum() int { return 4 }

func (o *xxx_LookupHandleFreeOperation) OpName() string { return "/epm/v3/ept_lookup_handle_free" }

func (o *xxx_LookupHandleFreeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupHandleFreeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// entry_handle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ept_lookup_handle_t, names=ndr_context_handle}(struct))
	{
		if o.EntryHandle != nil {
			if err := o.EntryHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&LookupHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_LookupHandleFreeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// entry_handle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ept_lookup_handle_t, names=ndr_context_handle}(struct))
	{
		if o.EntryHandle == nil {
			o.EntryHandle = &LookupHandle{}
		}
		if err := o.EntryHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupHandleFreeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupHandleFreeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// entry_handle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ept_lookup_handle_t, names=ndr_context_handle}(struct))
	{
		if o.EntryHandle != nil {
			if err := o.EntryHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&LookupHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupHandleFreeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// entry_handle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ept_lookup_handle_t, names=ndr_context_handle}(struct))
	{
		if o.EntryHandle == nil {
			o.EntryHandle = &LookupHandle{}
		}
		if err := o.EntryHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// LookupHandleFreeRequest structure represents the ept_lookup_handle_free operation request
type LookupHandleFreeRequest struct {
	EntryHandle *LookupHandle `idl:"name:entry_handle" json:"entry_handle"`
}

func (o *LookupHandleFreeRequest) xxx_ToOp(ctx context.Context, op *xxx_LookupHandleFreeOperation) *xxx_LookupHandleFreeOperation {
	if op == nil {
		op = &xxx_LookupHandleFreeOperation{}
	}
	if o == nil {
		return op
	}
	op.EntryHandle = o.EntryHandle
	return op
}

func (o *LookupHandleFreeRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupHandleFreeOperation) {
	if o == nil {
		return
	}
	o.EntryHandle = op.EntryHandle
}
func (o *LookupHandleFreeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupHandleFreeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupHandleFreeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupHandleFreeResponse structure represents the ept_lookup_handle_free operation response
type LookupHandleFreeResponse struct {
	EntryHandle *LookupHandle `idl:"name:entry_handle" json:"entry_handle"`
	Status      uint32        `idl:"name:status" json:"status"`
}

func (o *LookupHandleFreeResponse) xxx_ToOp(ctx context.Context, op *xxx_LookupHandleFreeOperation) *xxx_LookupHandleFreeOperation {
	if op == nil {
		op = &xxx_LookupHandleFreeOperation{}
	}
	if o == nil {
		return op
	}
	op.EntryHandle = o.EntryHandle
	op.Status = o.Status
	return op
}

func (o *LookupHandleFreeResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupHandleFreeOperation) {
	if o == nil {
		return
	}
	o.EntryHandle = op.EntryHandle
	o.Status = op.Status
}
func (o *LookupHandleFreeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupHandleFreeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupHandleFreeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InquireObjectOperation structure represents the ept_inq_object operation
type xxx_InquireObjectOperation struct {
	Object *dtyp.GUID `idl:"name:ept_object" json:"object"`
	Status uint32     `idl:"name:status" json:"status"`
}

func (o *xxx_InquireObjectOperation) OpNum() int { return 5 }

func (o *xxx_InquireObjectOperation) OpName() string { return "/epm/v3/ept_inq_object" }

func (o *xxx_InquireObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquireObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_InquireObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_InquireObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquireObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ept_object {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.Object != nil {
			if err := o.Object.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquireObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ept_object {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.Object == nil {
			o.Object = &dtyp.GUID{}
		}
		if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// InquireObjectRequest structure represents the ept_inq_object operation request
type InquireObjectRequest struct {
}

func (o *InquireObjectRequest) xxx_ToOp(ctx context.Context, op *xxx_InquireObjectOperation) *xxx_InquireObjectOperation {
	if op == nil {
		op = &xxx_InquireObjectOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *InquireObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_InquireObjectOperation) {
	if o == nil {
		return
	}
}
func (o *InquireObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InquireObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InquireObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InquireObjectResponse structure represents the ept_inq_object operation response
type InquireObjectResponse struct {
	Object *dtyp.GUID `idl:"name:ept_object" json:"object"`
	Status uint32     `idl:"name:status" json:"status"`
}

func (o *InquireObjectResponse) xxx_ToOp(ctx context.Context, op *xxx_InquireObjectOperation) *xxx_InquireObjectOperation {
	if op == nil {
		op = &xxx_InquireObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.Object = o.Object
	op.Status = o.Status
	return op
}

func (o *InquireObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_InquireObjectOperation) {
	if o == nil {
		return
	}
	o.Object = op.Object
	o.Status = op.Status
}
func (o *InquireObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InquireObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InquireObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ManagementDeleteOperation structure represents the ept_mgmt_delete operation
type xxx_ManagementDeleteOperation struct {
	ObjectSpeced bool            `idl:"name:object_speced" json:"object_speced"`
	Object       *dtyp.GUID      `idl:"name:object" json:"object"`
	Tower        *dcetypes.Tower `idl:"name:tower" json:"tower"`
	Status       uint32          `idl:"name:status" json:"status"`
}

func (o *xxx_ManagementDeleteOperation) OpNum() int { return 6 }

func (o *xxx_ManagementDeleteOperation) OpName() string { return "/epm/v3/ept_mgmt_delete" }

func (o *xxx_ManagementDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// object_speced {in} (1:{alias=boolean32, names=unsigned32}(uint32))
	{
		if !o.ObjectSpeced {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(uint32(1)); err != nil {
				return err
			}
		}
	}
	// object {in} (1:{pointer=ptr, alias=uuid_p_t}*(1))(2:{alias=GUID}(struct))
	{
		if o.Object != nil {
			_ptr_object := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Object != nil {
					if err := o.Object.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Object, _ptr_object); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tower {in} (1:{pointer=ptr, alias=twr_p_t}*(1))(2:{alias=twr_t}(struct))
	{
		if o.Tower != nil {
			_ptr_tower := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Tower != nil {
					if err := o.Tower.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcetypes.Tower{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Tower, _ptr_tower); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// object_speced {in} (1:{alias=boolean32, names=unsigned32}(uint32))
	{
		var _bObjectSpeced uint32
		if err := w.ReadData(&_bObjectSpeced); err != nil {
			return err
		}
		o.ObjectSpeced = _bObjectSpeced != 0
	}
	// object {in} (1:{pointer=ptr, alias=uuid_p_t}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_object := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Object == nil {
				o.Object = &dtyp.GUID{}
			}
			if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_object := func(ptr interface{}) { o.Object = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.Object, _s_object, _ptr_object); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tower {in} (1:{pointer=ptr, alias=twr_p_t}*(1))(2:{alias=twr_t}(struct))
	{
		_ptr_tower := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Tower == nil {
				o.Tower = &dcetypes.Tower{}
			}
			if err := o.Tower.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_tower := func(ptr interface{}) { o.Tower = *ptr.(**dcetypes.Tower) }
		if err := w.ReadPointer(&o.Tower, _s_tower, _ptr_tower); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// ManagementDeleteRequest structure represents the ept_mgmt_delete operation request
type ManagementDeleteRequest struct {
	ObjectSpeced bool            `idl:"name:object_speced" json:"object_speced"`
	Object       *dtyp.GUID      `idl:"name:object" json:"object"`
	Tower        *dcetypes.Tower `idl:"name:tower" json:"tower"`
}

func (o *ManagementDeleteRequest) xxx_ToOp(ctx context.Context, op *xxx_ManagementDeleteOperation) *xxx_ManagementDeleteOperation {
	if op == nil {
		op = &xxx_ManagementDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectSpeced = o.ObjectSpeced
	op.Object = o.Object
	op.Tower = o.Tower
	return op
}

func (o *ManagementDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_ManagementDeleteOperation) {
	if o == nil {
		return
	}
	o.ObjectSpeced = op.ObjectSpeced
	o.Object = op.Object
	o.Tower = op.Tower
}
func (o *ManagementDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ManagementDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ManagementDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ManagementDeleteResponse structure represents the ept_mgmt_delete operation response
type ManagementDeleteResponse struct {
	Status uint32 `idl:"name:status" json:"status"`
}

func (o *ManagementDeleteResponse) xxx_ToOp(ctx context.Context, op *xxx_ManagementDeleteOperation) *xxx_ManagementDeleteOperation {
	if op == nil {
		op = &xxx_ManagementDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.Status = o.Status
	return op
}

func (o *ManagementDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_ManagementDeleteOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
}
func (o *ManagementDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ManagementDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ManagementDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
