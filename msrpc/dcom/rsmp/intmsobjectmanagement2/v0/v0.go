package intmsobjectmanagement2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	intmsobjectmanagement1 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsobjectmanagement1/v0"
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
	_ = dcom.GoPackage
	_ = intmsobjectmanagement1.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// INtmsObjectManagement2 interface identifier 895a2c86-270d-489d-a6c0-dc2a9b35280e
	ObjectManagement2IID = &dcom.IID{Data1: 0x895a2c86, Data2: 0x270d, Data3: 0x489d, Data4: []byte{0xa6, 0xc0, 0xdc, 0x2a, 0x9b, 0x35, 0x28, 0x0e}}
	// Syntax UUID
	ObjectManagement2SyntaxUUID = &uuid.UUID{TimeLow: 0x895a2c86, TimeMid: 0x270d, TimeHiAndVersion: 0x489d, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0xc0, Node: [6]uint8{0xdc, 0x2a, 0x9b, 0x35, 0x28, 0xe}}
	// Syntax ID
	ObjectManagement2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ObjectManagement2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// INtmsObjectManagement2 interface.
type ObjectManagement2Client interface {

	// INtmsObjectManagement1 retrieval method.
	ObjectManagement1() intmsobjectmanagement1.ObjectManagement1Client

	EnumerateNTMSObjectR(context.Context, *EnumerateNTMSObjectRRequest, ...dcerpc.CallOption) (*EnumerateNTMSObjectRResponse, error)

	GetNTMSUIOptionsA(context.Context, *GetNTMSUIOptionsARequest, ...dcerpc.CallOption) (*GetNTMSUIOptionsAResponse, error)

	GetNTMSUIOptionsW(context.Context, *GetNTMSUIOptionsWRequest, ...dcerpc.CallOption) (*GetNTMSUIOptionsWResponse, error)

	SetNTMSUIOptionsA(context.Context, *SetNTMSUIOptionsARequest, ...dcerpc.CallOption) (*SetNTMSUIOptionsAResponse, error)

	SetNTMSUIOptionsW(context.Context, *SetNTMSUIOptionsWRequest, ...dcerpc.CallOption) (*SetNTMSUIOptionsWResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ObjectManagement2Client
}

type xxx_DefaultObjectManagement2Client struct {
	intmsobjectmanagement1.ObjectManagement1Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultObjectManagement2Client) ObjectManagement1() intmsobjectmanagement1.ObjectManagement1Client {
	return o.ObjectManagement1Client
}

func (o *xxx_DefaultObjectManagement2Client) EnumerateNTMSObjectR(ctx context.Context, in *EnumerateNTMSObjectRRequest, opts ...dcerpc.CallOption) (*EnumerateNTMSObjectRResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateNTMSObjectRResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement2Client) GetNTMSUIOptionsA(ctx context.Context, in *GetNTMSUIOptionsARequest, opts ...dcerpc.CallOption) (*GetNTMSUIOptionsAResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNTMSUIOptionsAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement2Client) GetNTMSUIOptionsW(ctx context.Context, in *GetNTMSUIOptionsWRequest, opts ...dcerpc.CallOption) (*GetNTMSUIOptionsWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNTMSUIOptionsWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement2Client) SetNTMSUIOptionsA(ctx context.Context, in *SetNTMSUIOptionsARequest, opts ...dcerpc.CallOption) (*SetNTMSUIOptionsAResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetNTMSUIOptionsAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement2Client) SetNTMSUIOptionsW(ctx context.Context, in *SetNTMSUIOptionsWRequest, opts ...dcerpc.CallOption) (*SetNTMSUIOptionsWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetNTMSUIOptionsWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultObjectManagement2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultObjectManagement2Client) IPID(ctx context.Context, ipid *dcom.IPID) ObjectManagement2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultObjectManagement2Client{
		ObjectManagement1Client: o.ObjectManagement1Client.IPID(ctx, ipid),
		cc:                      o.cc,
		ipid:                    ipid,
	}
}

func NewObjectManagement2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ObjectManagement2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ObjectManagement2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := intmsobjectmanagement1.NewObjectManagement1Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultObjectManagement2Client{
		ObjectManagement1Client: base,
		cc:                      cc,
		ipid:                    ipid,
	}, nil
}

// xxx_EnumerateNTMSObjectROperation structure represents the EnumerateNtmsObjectR operation
type xxx_EnumerateNTMSObjectROperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	ContainerID    *dtyp.GUID     `idl:"name:lpContainerId;pointer:unique" json:"container_id"`
	List           []*dtyp.GUID   `idl:"name:lpList;size_is:(lpdwListBufferSize);length_is:(lpdwListSize)" json:"list"`
	ListBufferSize uint32         `idl:"name:lpdwListBufferSize" json:"list_buffer_size"`
	ListSize       uint32         `idl:"name:lpdwListSize" json:"list_size"`
	OutputSize     uint32         `idl:"name:lpdwOutputSize" json:"output_size"`
	Type           uint32         `idl:"name:dwType" json:"type"`
	Options        uint32         `idl:"name:dwOptions" json:"options"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateNTMSObjectROperation) OpNum() int { return 9 }

func (o *xxx_EnumerateNTMSObjectROperation) OpName() string {
	return "/INtmsObjectManagement2/v0/EnumerateNtmsObjectR"
}

func (o *xxx_EnumerateNTMSObjectROperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateNTMSObjectROperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpContainerId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID != nil {
			_ptr_lpContainerId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ContainerID != nil {
					if err := o.ContainerID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ContainerID, _ptr_lpContainerId); err != nil {
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
	// lpdwListBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ListBufferSize); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateNTMSObjectROperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpContainerId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpContainerId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ContainerID == nil {
				o.ContainerID = &dtyp.GUID{}
			}
			if err := o.ContainerID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpContainerId := func(ptr interface{}) { o.ContainerID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.ContainerID, _s_lpContainerId, _ptr_lpContainerId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwListBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ListBufferSize); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateNTMSObjectROperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.List != nil && o.ListSize == 0 {
		o.ListSize = uint32(len(o.List))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateNTMSObjectROperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpList {out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}[dim:0,size_is=lpdwListBufferSize,length_is=lpdwListSize](struct))
	{
		dimSize1 := uint64(o.ListBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ListSize)
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
		for i1 := range o.List {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.List[i1] != nil {
				if err := o.List[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.List); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwListSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ListSize); err != nil {
			return err
		}
	}
	// lpdwOutputSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutputSize); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateNTMSObjectROperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpList {out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}[dim:0,size_is=lpdwListBufferSize,length_is=lpdwListSize](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.List", sizeInfo[0])
		}
		o.List = make([]*dtyp.GUID, sizeInfo[0])
		for i1 := range o.List {
			i1 := i1
			if o.List[i1] == nil {
				o.List[i1] = &dtyp.GUID{}
			}
			if err := o.List[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwListSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ListSize); err != nil {
			return err
		}
	}
	// lpdwOutputSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutputSize); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumerateNTMSObjectRRequest structure represents the EnumerateNtmsObjectR operation request
type EnumerateNTMSObjectRRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	ContainerID    *dtyp.GUID     `idl:"name:lpContainerId;pointer:unique" json:"container_id"`
	ListBufferSize uint32         `idl:"name:lpdwListBufferSize" json:"list_buffer_size"`
	Type           uint32         `idl:"name:dwType" json:"type"`
	Options        uint32         `idl:"name:dwOptions" json:"options"`
}

func (o *EnumerateNTMSObjectRRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumerateNTMSObjectROperation) *xxx_EnumerateNTMSObjectROperation {
	if op == nil {
		op = &xxx_EnumerateNTMSObjectROperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ContainerID = o.ContainerID
	op.ListBufferSize = o.ListBufferSize
	op.Type = o.Type
	op.Options = o.Options
	return op
}

func (o *EnumerateNTMSObjectRRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateNTMSObjectROperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ContainerID = op.ContainerID
	o.ListBufferSize = op.ListBufferSize
	o.Type = op.Type
	o.Options = op.Options
}
func (o *EnumerateNTMSObjectRRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateNTMSObjectRRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateNTMSObjectROperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateNTMSObjectRResponse structure represents the EnumerateNtmsObjectR operation response
type EnumerateNTMSObjectRResponse struct {
	// XXX: lpdwListBufferSize is an implicit input depedency for output parameters
	ListBufferSize uint32 `idl:"name:lpdwListBufferSize" json:"list_buffer_size"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	List       []*dtyp.GUID   `idl:"name:lpList;size_is:(lpdwListBufferSize);length_is:(lpdwListSize)" json:"list"`
	ListSize   uint32         `idl:"name:lpdwListSize" json:"list_size"`
	OutputSize uint32         `idl:"name:lpdwOutputSize" json:"output_size"`
	// Return: The EnumerateNtmsObjectR return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateNTMSObjectRResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumerateNTMSObjectROperation) *xxx_EnumerateNTMSObjectROperation {
	if op == nil {
		op = &xxx_EnumerateNTMSObjectROperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.ListBufferSize == uint32(0) {
		op.ListBufferSize = o.ListBufferSize
	}

	op.That = o.That
	op.List = o.List
	op.ListSize = o.ListSize
	op.OutputSize = o.OutputSize
	op.Return = o.Return
	return op
}

func (o *EnumerateNTMSObjectRResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateNTMSObjectROperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.ListBufferSize = op.ListBufferSize

	o.That = op.That
	o.List = op.List
	o.ListSize = op.ListSize
	o.OutputSize = op.OutputSize
	o.Return = op.Return
}
func (o *EnumerateNTMSObjectRResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateNTMSObjectRResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateNTMSObjectROperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNTMSUIOptionsAOperation structure represents the GetNtmsUIOptionsA operation
type xxx_GetNTMSUIOptionsAOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID    *dtyp.GUID     `idl:"name:lpObjectId;pointer:unique" json:"object_id"`
	Type        uint32         `idl:"name:dwType" json:"type"`
	Destination []byte         `idl:"name:lpszDestination;size_is:(lpdwBufSize);length_is:(lpdwDataSize)" json:"destination"`
	BufferSize  uint32         `idl:"name:lpdwBufSize" json:"buffer_size"`
	DataSize    uint32         `idl:"name:lpdwDataSize" json:"data_size"`
	OutSize     uint32         `idl:"name:lpdwOutSize" json:"out_size"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSUIOptionsAOperation) OpNum() int { return 10 }

func (o *xxx_GetNTMSUIOptionsAOperation) OpName() string {
	return "/INtmsObjectManagement2/v0/GetNtmsUIOptionsA"
}

func (o *xxx_GetNTMSUIOptionsAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSUIOptionsAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			_ptr_lpObjectId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectID != nil {
					if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectID, _ptr_lpObjectId); err != nil {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpdwBufSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSUIOptionsAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpObjectId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectID == nil {
				o.ObjectID = &dtyp.GUID{}
			}
			if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpObjectId := func(ptr interface{}) { o.ObjectID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.ObjectID, _s_lpObjectId, _ptr_lpObjectId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpdwBufSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSUIOptionsAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Destination != nil && o.DataSize == 0 {
		o.DataSize = uint32(len(o.Destination))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSUIOptionsAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpszDestination {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwBufSize,length_is=lpdwDataSize](uchar))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.DataSize)
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
		for i1 := range o.Destination {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Destination[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Destination); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpdwDataSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataSize); err != nil {
			return err
		}
	}
	// lpdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutSize); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSUIOptionsAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpszDestination {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwBufSize,length_is=lpdwDataSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Destination", sizeInfo[0])
		}
		o.Destination = make([]byte, sizeInfo[0])
		for i1 := range o.Destination {
			i1 := i1
			if err := w.ReadData(&o.Destination[i1]); err != nil {
				return err
			}
		}
	}
	// lpdwDataSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataSize); err != nil {
			return err
		}
	}
	// lpdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutSize); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNTMSUIOptionsARequest structure represents the GetNtmsUIOptionsA operation request
type GetNTMSUIOptionsARequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	ObjectID   *dtyp.GUID     `idl:"name:lpObjectId;pointer:unique" json:"object_id"`
	Type       uint32         `idl:"name:dwType" json:"type"`
	BufferSize uint32         `idl:"name:lpdwBufSize" json:"buffer_size"`
}

func (o *GetNTMSUIOptionsARequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSUIOptionsAOperation) *xxx_GetNTMSUIOptionsAOperation {
	if op == nil {
		op = &xxx_GetNTMSUIOptionsAOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.BufferSize = o.BufferSize
	return op
}

func (o *GetNTMSUIOptionsARequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSUIOptionsAOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.BufferSize = op.BufferSize
}
func (o *GetNTMSUIOptionsARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSUIOptionsARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSUIOptionsAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSUIOptionsAResponse structure represents the GetNtmsUIOptionsA operation response
type GetNTMSUIOptionsAResponse struct {
	// XXX: lpdwBufSize is an implicit input depedency for output parameters
	BufferSize uint32 `idl:"name:lpdwBufSize" json:"buffer_size"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Destination []byte         `idl:"name:lpszDestination;size_is:(lpdwBufSize);length_is:(lpdwDataSize)" json:"destination"`
	DataSize    uint32         `idl:"name:lpdwDataSize" json:"data_size"`
	OutSize     uint32         `idl:"name:lpdwOutSize" json:"out_size"`
	// Return: The GetNtmsUIOptionsA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSUIOptionsAResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSUIOptionsAOperation) *xxx_GetNTMSUIOptionsAOperation {
	if op == nil {
		op = &xxx_GetNTMSUIOptionsAOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferSize == uint32(0) {
		op.BufferSize = o.BufferSize
	}

	op.That = o.That
	op.Destination = o.Destination
	op.DataSize = o.DataSize
	op.OutSize = o.OutSize
	op.Return = o.Return
	return op
}

func (o *GetNTMSUIOptionsAResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSUIOptionsAOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferSize = op.BufferSize

	o.That = op.That
	o.Destination = op.Destination
	o.DataSize = op.DataSize
	o.OutSize = op.OutSize
	o.Return = op.Return
}
func (o *GetNTMSUIOptionsAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSUIOptionsAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSUIOptionsAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNTMSUIOptionsWOperation structure represents the GetNtmsUIOptionsW operation
type xxx_GetNTMSUIOptionsWOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID    *dtyp.GUID     `idl:"name:lpObjectId;pointer:unique" json:"object_id"`
	Type        uint32         `idl:"name:dwType" json:"type"`
	Destination string         `idl:"name:lpszDestination;size_is:(lpdwBufSize);length_is:(lpdwDataSize)" json:"destination"`
	BufferSize  uint32         `idl:"name:lpdwBufSize" json:"buffer_size"`
	DataSize    uint32         `idl:"name:lpdwDataSize" json:"data_size"`
	OutSize     uint32         `idl:"name:lpdwOutSize" json:"out_size"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSUIOptionsWOperation) OpNum() int { return 11 }

func (o *xxx_GetNTMSUIOptionsWOperation) OpName() string {
	return "/INtmsObjectManagement2/v0/GetNtmsUIOptionsW"
}

func (o *xxx_GetNTMSUIOptionsWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSUIOptionsWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			_ptr_lpObjectId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectID != nil {
					if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectID, _ptr_lpObjectId); err != nil {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpdwBufSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSUIOptionsWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpObjectId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectID == nil {
				o.ObjectID = &dtyp.GUID{}
			}
			if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpObjectId := func(ptr interface{}) { o.ObjectID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.ObjectID, _s_lpObjectId, _ptr_lpObjectId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpdwBufSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSUIOptionsWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Destination != "" && o.DataSize == 0 {
		o.DataSize = uint32(ndr.UTF16Len(o.Destination))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSUIOptionsWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpszDestination {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwBufSize,length_is=lpdwDataSize,string](wchar))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.DataSize)
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
		_Destination_buf := utf16.Encode([]rune(o.Destination))
		if uint64(len(_Destination_buf)) > sizeInfo[0] {
			_Destination_buf = _Destination_buf[:sizeInfo[0]]
		}
		for i1 := range _Destination_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Destination_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Destination_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// lpdwDataSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataSize); err != nil {
			return err
		}
	}
	// lpdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutSize); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSUIOptionsWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpszDestination {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwBufSize,length_is=lpdwDataSize,string](wchar))
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
		var _Destination_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Destination_buf", sizeInfo[0])
		}
		_Destination_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Destination_buf {
			i1 := i1
			if err := w.ReadData(&_Destination_buf[i1]); err != nil {
				return err
			}
		}
		o.Destination = strings.TrimRight(string(utf16.Decode(_Destination_buf)), ndr.ZeroString)
	}
	// lpdwDataSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataSize); err != nil {
			return err
		}
	}
	// lpdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutSize); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNTMSUIOptionsWRequest structure represents the GetNtmsUIOptionsW operation request
type GetNTMSUIOptionsWRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	ObjectID   *dtyp.GUID     `idl:"name:lpObjectId;pointer:unique" json:"object_id"`
	Type       uint32         `idl:"name:dwType" json:"type"`
	BufferSize uint32         `idl:"name:lpdwBufSize" json:"buffer_size"`
}

func (o *GetNTMSUIOptionsWRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSUIOptionsWOperation) *xxx_GetNTMSUIOptionsWOperation {
	if op == nil {
		op = &xxx_GetNTMSUIOptionsWOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.BufferSize = o.BufferSize
	return op
}

func (o *GetNTMSUIOptionsWRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSUIOptionsWOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.BufferSize = op.BufferSize
}
func (o *GetNTMSUIOptionsWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSUIOptionsWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSUIOptionsWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSUIOptionsWResponse structure represents the GetNtmsUIOptionsW operation response
type GetNTMSUIOptionsWResponse struct {
	// XXX: lpdwBufSize is an implicit input depedency for output parameters
	BufferSize uint32 `idl:"name:lpdwBufSize" json:"buffer_size"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Destination string         `idl:"name:lpszDestination;size_is:(lpdwBufSize);length_is:(lpdwDataSize)" json:"destination"`
	DataSize    uint32         `idl:"name:lpdwDataSize" json:"data_size"`
	OutSize     uint32         `idl:"name:lpdwOutSize" json:"out_size"`
	// Return: The GetNtmsUIOptionsW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSUIOptionsWResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSUIOptionsWOperation) *xxx_GetNTMSUIOptionsWOperation {
	if op == nil {
		op = &xxx_GetNTMSUIOptionsWOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferSize == uint32(0) {
		op.BufferSize = o.BufferSize
	}

	op.That = o.That
	op.Destination = o.Destination
	op.DataSize = o.DataSize
	op.OutSize = o.OutSize
	op.Return = o.Return
	return op
}

func (o *GetNTMSUIOptionsWResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSUIOptionsWOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferSize = op.BufferSize

	o.That = op.That
	o.Destination = op.Destination
	o.DataSize = op.DataSize
	o.OutSize = op.OutSize
	o.Return = op.Return
}
func (o *GetNTMSUIOptionsWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSUIOptionsWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSUIOptionsWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNTMSUIOptionsAOperation structure represents the SetNtmsUIOptionsA operation
type xxx_SetNTMSUIOptionsAOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID    *dtyp.GUID     `idl:"name:lpObjectId;pointer:unique" json:"object_id"`
	Type        uint32         `idl:"name:dwType" json:"type"`
	Operation   uint32         `idl:"name:dwOperation" json:"operation"`
	Destination string         `idl:"name:lpszDestination;string" json:"destination"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNTMSUIOptionsAOperation) OpNum() int { return 12 }

func (o *xxx_SetNTMSUIOptionsAOperation) OpName() string {
	return "/INtmsObjectManagement2/v0/SetNtmsUIOptionsA"
}

func (o *xxx_SetNTMSUIOptionsAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSUIOptionsAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			_ptr_lpObjectId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectID != nil {
					if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectID, _ptr_lpObjectId); err != nil {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// dwOperation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Operation); err != nil {
			return err
		}
	}
	// lpszDestination {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](char))
	{
		if err := ndr.WriteCharNString(ctx, w, o.Destination); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSUIOptionsAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpObjectId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectID == nil {
				o.ObjectID = &dtyp.GUID{}
			}
			if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpObjectId := func(ptr interface{}) { o.ObjectID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.ObjectID, _s_lpObjectId, _ptr_lpObjectId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// dwOperation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Operation); err != nil {
			return err
		}
	}
	// lpszDestination {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](char))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.Destination); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSUIOptionsAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSUIOptionsAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSUIOptionsAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetNTMSUIOptionsARequest structure represents the SetNtmsUIOptionsA operation request
type SetNTMSUIOptionsARequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	ObjectID    *dtyp.GUID     `idl:"name:lpObjectId;pointer:unique" json:"object_id"`
	Type        uint32         `idl:"name:dwType" json:"type"`
	Operation   uint32         `idl:"name:dwOperation" json:"operation"`
	Destination string         `idl:"name:lpszDestination;string" json:"destination"`
}

func (o *SetNTMSUIOptionsARequest) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSUIOptionsAOperation) *xxx_SetNTMSUIOptionsAOperation {
	if op == nil {
		op = &xxx_SetNTMSUIOptionsAOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.Operation = o.Operation
	op.Destination = o.Destination
	return op
}

func (o *SetNTMSUIOptionsARequest) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSUIOptionsAOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.Operation = op.Operation
	o.Destination = op.Destination
}
func (o *SetNTMSUIOptionsARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNTMSUIOptionsARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSUIOptionsAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNTMSUIOptionsAResponse structure represents the SetNtmsUIOptionsA operation response
type SetNTMSUIOptionsAResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetNtmsUIOptionsA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNTMSUIOptionsAResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSUIOptionsAOperation) *xxx_SetNTMSUIOptionsAOperation {
	if op == nil {
		op = &xxx_SetNTMSUIOptionsAOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNTMSUIOptionsAResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSUIOptionsAOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNTMSUIOptionsAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNTMSUIOptionsAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSUIOptionsAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNTMSUIOptionsWOperation structure represents the SetNtmsUIOptionsW operation
type xxx_SetNTMSUIOptionsWOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID    *dtyp.GUID     `idl:"name:lpObjectId;pointer:unique" json:"object_id"`
	Type        uint32         `idl:"name:dwType" json:"type"`
	Operation   uint32         `idl:"name:dwOperation" json:"operation"`
	Destination string         `idl:"name:lpszDestination;string" json:"destination"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNTMSUIOptionsWOperation) OpNum() int { return 13 }

func (o *xxx_SetNTMSUIOptionsWOperation) OpName() string {
	return "/INtmsObjectManagement2/v0/SetNtmsUIOptionsW"
}

func (o *xxx_SetNTMSUIOptionsWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSUIOptionsWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			_ptr_lpObjectId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectID != nil {
					if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectID, _ptr_lpObjectId); err != nil {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// dwOperation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Operation); err != nil {
			return err
		}
	}
	// lpszDestination {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Destination); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSUIOptionsWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpObjectId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectID == nil {
				o.ObjectID = &dtyp.GUID{}
			}
			if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpObjectId := func(ptr interface{}) { o.ObjectID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.ObjectID, _s_lpObjectId, _ptr_lpObjectId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// dwOperation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Operation); err != nil {
			return err
		}
	}
	// lpszDestination {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Destination); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSUIOptionsWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSUIOptionsWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSUIOptionsWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetNTMSUIOptionsWRequest structure represents the SetNtmsUIOptionsW operation request
type SetNTMSUIOptionsWRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	ObjectID    *dtyp.GUID     `idl:"name:lpObjectId;pointer:unique" json:"object_id"`
	Type        uint32         `idl:"name:dwType" json:"type"`
	Operation   uint32         `idl:"name:dwOperation" json:"operation"`
	Destination string         `idl:"name:lpszDestination;string" json:"destination"`
}

func (o *SetNTMSUIOptionsWRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSUIOptionsWOperation) *xxx_SetNTMSUIOptionsWOperation {
	if op == nil {
		op = &xxx_SetNTMSUIOptionsWOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.Operation = o.Operation
	op.Destination = o.Destination
	return op
}

func (o *SetNTMSUIOptionsWRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSUIOptionsWOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.Operation = op.Operation
	o.Destination = op.Destination
}
func (o *SetNTMSUIOptionsWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNTMSUIOptionsWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSUIOptionsWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNTMSUIOptionsWResponse structure represents the SetNtmsUIOptionsW operation response
type SetNTMSUIOptionsWResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetNtmsUIOptionsW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNTMSUIOptionsWResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSUIOptionsWOperation) *xxx_SetNTMSUIOptionsWOperation {
	if op == nil {
		op = &xxx_SetNTMSUIOptionsWOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNTMSUIOptionsWResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSUIOptionsWOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNTMSUIOptionsWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNTMSUIOptionsWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSUIOptionsWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
