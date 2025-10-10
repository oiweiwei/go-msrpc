package iresourcemanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
	wsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
	_ = wsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

var (
	// IResourceManager interface identifier c5cebee2-9df5-4cdd-a08c-c2471bc144b4
	ResourceManagerIID = &dcom.IID{Data1: 0xc5cebee2, Data2: 0x9df5, Data3: 0x4cdd, Data4: []byte{0xa0, 0x8c, 0xc2, 0x47, 0x1b, 0xc1, 0x44, 0xb4}}
	// Syntax UUID
	ResourceManagerSyntaxUUID = &uuid.UUID{TimeLow: 0xc5cebee2, TimeMid: 0x9df5, TimeHiAndVersion: 0x4cdd, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x8c, Node: [6]uint8{0xc2, 0x47, 0x1b, 0xc1, 0x44, 0xb4}}
	// Syntax ID
	ResourceManagerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ResourceManagerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IResourceManager interface.
type ResourceManagerClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	RetrieveEventList(context.Context, *RetrieveEventListRequest, ...dcerpc.CallOption) (*RetrieveEventListResponse, error)

	GetSystemAffinity(context.Context, *GetSystemAffinityRequest, ...dcerpc.CallOption) (*GetSystemAffinityResponse, error)

	ImportXMLFiles(context.Context, *ImportXMLFilesRequest, ...dcerpc.CallOption) (*ImportXMLFilesResponse, error)

	ExportXMLFiles(context.Context, *ExportXMLFilesRequest, ...dcerpc.CallOption) (*ExportXMLFilesResponse, error)

	RestoreXMLFiles(context.Context, *RestoreXMLFilesRequest, ...dcerpc.CallOption) (*RestoreXMLFilesResponse, error)

	GetDependencies(context.Context, *GetDependenciesRequest, ...dcerpc.CallOption) (*GetDependenciesResponse, error)

	GetServiceList(context.Context, *GetServiceListRequest, ...dcerpc.CallOption) (*GetServiceListResponse, error)

	GetIISAppPoolNames(context.Context, *GetIISAppPoolNamesRequest, ...dcerpc.CallOption) (*GetIISAppPoolNamesResponse, error)

	GetServerName(context.Context, *GetServerNameRequest, ...dcerpc.CallOption) (*GetServerNameResponse, error)

	GetCurrentMemory(context.Context, *GetCurrentMemoryRequest, ...dcerpc.CallOption) (*GetCurrentMemoryResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ResourceManagerClient
}

type xxx_DefaultResourceManagerClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultResourceManagerClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultResourceManagerClient) RetrieveEventList(ctx context.Context, in *RetrieveEventListRequest, opts ...dcerpc.CallOption) (*RetrieveEventListResponse, error) {
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
	out := &RetrieveEventListResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManagerClient) GetSystemAffinity(ctx context.Context, in *GetSystemAffinityRequest, opts ...dcerpc.CallOption) (*GetSystemAffinityResponse, error) {
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
	out := &GetSystemAffinityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManagerClient) ImportXMLFiles(ctx context.Context, in *ImportXMLFilesRequest, opts ...dcerpc.CallOption) (*ImportXMLFilesResponse, error) {
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
	out := &ImportXMLFilesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManagerClient) ExportXMLFiles(ctx context.Context, in *ExportXMLFilesRequest, opts ...dcerpc.CallOption) (*ExportXMLFilesResponse, error) {
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
	out := &ExportXMLFilesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManagerClient) RestoreXMLFiles(ctx context.Context, in *RestoreXMLFilesRequest, opts ...dcerpc.CallOption) (*RestoreXMLFilesResponse, error) {
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
	out := &RestoreXMLFilesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManagerClient) GetDependencies(ctx context.Context, in *GetDependenciesRequest, opts ...dcerpc.CallOption) (*GetDependenciesResponse, error) {
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
	out := &GetDependenciesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManagerClient) GetServiceList(ctx context.Context, in *GetServiceListRequest, opts ...dcerpc.CallOption) (*GetServiceListResponse, error) {
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
	out := &GetServiceListResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManagerClient) GetIISAppPoolNames(ctx context.Context, in *GetIISAppPoolNamesRequest, opts ...dcerpc.CallOption) (*GetIISAppPoolNamesResponse, error) {
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
	out := &GetIISAppPoolNamesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManagerClient) GetServerName(ctx context.Context, in *GetServerNameRequest, opts ...dcerpc.CallOption) (*GetServerNameResponse, error) {
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
	out := &GetServerNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManagerClient) GetCurrentMemory(ctx context.Context, in *GetCurrentMemoryRequest, opts ...dcerpc.CallOption) (*GetCurrentMemoryResponse, error) {
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
	out := &GetCurrentMemoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultResourceManagerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultResourceManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) ResourceManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultResourceManagerClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewResourceManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ResourceManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ResourceManagerSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultResourceManagerClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_RetrieveEventListOperation structure represents the RetrieveEventList operation
type xxx_RetrieveEventListOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventList *oaut.String   `idl:"name:pbstrEventList" json:"event_list"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RetrieveEventListOperation) OpNum() int { return 7 }

func (o *xxx_RetrieveEventListOperation) OpName() string {
	return "/IResourceManager/v0/RetrieveEventList"
}

func (o *xxx_RetrieveEventListOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveEventListOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_RetrieveEventListOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_RetrieveEventListOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveEventListOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrEventList {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EventList != nil {
			_ptr_pbstrEventList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EventList != nil {
					if err := o.EventList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventList, _ptr_pbstrEventList); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveEventListOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrEventList {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrEventList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EventList == nil {
				o.EventList = &oaut.String{}
			}
			if err := o.EventList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrEventList := func(ptr interface{}) { o.EventList = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EventList, _s_pbstrEventList, _ptr_pbstrEventList); err != nil {
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

// RetrieveEventListRequest structure represents the RetrieveEventList operation request
type RetrieveEventListRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RetrieveEventListRequest) xxx_ToOp(ctx context.Context, op *xxx_RetrieveEventListOperation) *xxx_RetrieveEventListOperation {
	if op == nil {
		op = &xxx_RetrieveEventListOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *RetrieveEventListRequest) xxx_FromOp(ctx context.Context, op *xxx_RetrieveEventListOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RetrieveEventListRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RetrieveEventListRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrieveEventListOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RetrieveEventListResponse structure represents the RetrieveEventList operation response
type RetrieveEventListResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventList *oaut.String   `idl:"name:pbstrEventList" json:"event_list"`
	// Return: The RetrieveEventList return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RetrieveEventListResponse) xxx_ToOp(ctx context.Context, op *xxx_RetrieveEventListOperation) *xxx_RetrieveEventListOperation {
	if op == nil {
		op = &xxx_RetrieveEventListOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.EventList = o.EventList
	op.Return = o.Return
	return op
}

func (o *RetrieveEventListResponse) xxx_FromOp(ctx context.Context, op *xxx_RetrieveEventListOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EventList = op.EventList
	o.Return = op.Return
}
func (o *RetrieveEventListResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RetrieveEventListResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrieveEventListOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSystemAffinityOperation structure represents the GetSystemAffinity operation
type xxx_GetSystemAffinityOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	SystemAffinity uint64         `idl:"name:pdwSysAffinity" json:"system_affinity"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSystemAffinityOperation) OpNum() int { return 8 }

func (o *xxx_GetSystemAffinityOperation) OpName() string {
	return "/IResourceManager/v0/GetSystemAffinity"
}

func (o *xxx_GetSystemAffinityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemAffinityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_GetSystemAffinityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_GetSystemAffinityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemAffinityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwSysAffinity {out} (1:{pointer=ref}*(1))(2:{alias=DWORD64}(uint64))
	{
		if err := w.WriteData(o.SystemAffinity); err != nil {
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

func (o *xxx_GetSystemAffinityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwSysAffinity {out} (1:{pointer=ref}*(1))(2:{alias=DWORD64}(uint64))
	{
		if err := w.ReadData(&o.SystemAffinity); err != nil {
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

// GetSystemAffinityRequest structure represents the GetSystemAffinity operation request
type GetSystemAffinityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSystemAffinityRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSystemAffinityOperation) *xxx_GetSystemAffinityOperation {
	if op == nil {
		op = &xxx_GetSystemAffinityOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSystemAffinityRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSystemAffinityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSystemAffinityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSystemAffinityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSystemAffinityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSystemAffinityResponse structure represents the GetSystemAffinity operation response
type GetSystemAffinityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	SystemAffinity uint64         `idl:"name:pdwSysAffinity" json:"system_affinity"`
	// Return: The GetSystemAffinity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSystemAffinityResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSystemAffinityOperation) *xxx_GetSystemAffinityOperation {
	if op == nil {
		op = &xxx_GetSystemAffinityOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SystemAffinity = o.SystemAffinity
	op.Return = o.Return
	return op
}

func (o *GetSystemAffinityResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSystemAffinityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SystemAffinity = op.SystemAffinity
	o.Return = op.Return
}
func (o *GetSystemAffinityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSystemAffinityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSystemAffinityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ImportXMLFilesOperation structure represents the ImportXMLFiles operation
type xxx_ImportXMLFilesOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	PMCXML         *oaut.String   `idl:"name:bstrPMCXml" json:"pmc_xml"`
	PolicyXML      *oaut.String   `idl:"name:bstrPolicyXml" json:"policy_xml"`
	CalendarXML    *oaut.String   `idl:"name:bstrCalendarXml" json:"calendar_xml"`
	ConditionalXML *oaut.String   `idl:"name:bstrConditionalXml" json:"conditional_xml"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ImportXMLFilesOperation) OpNum() int { return 9 }

func (o *xxx_ImportXMLFilesOperation) OpName() string { return "/IResourceManager/v0/ImportXMLFiles" }

func (o *xxx_ImportXMLFilesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportXMLFilesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPMCXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PMCXML != nil {
			_ptr_bstrPMCXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PMCXML != nil {
					if err := o.PMCXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PMCXML, _ptr_bstrPMCXml); err != nil {
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
	// bstrPolicyXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyXML != nil {
			_ptr_bstrPolicyXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyXML != nil {
					if err := o.PolicyXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyXML, _ptr_bstrPolicyXml); err != nil {
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
	// bstrCalendarXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CalendarXML != nil {
			_ptr_bstrCalendarXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CalendarXML != nil {
					if err := o.CalendarXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CalendarXML, _ptr_bstrCalendarXml); err != nil {
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
	// bstrConditionalXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConditionalXML != nil {
			_ptr_bstrConditionalXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConditionalXML != nil {
					if err := o.ConditionalXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConditionalXML, _ptr_bstrConditionalXml); err != nil {
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

func (o *xxx_ImportXMLFilesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPMCXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPMCXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PMCXML == nil {
				o.PMCXML = &oaut.String{}
			}
			if err := o.PMCXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPMCXml := func(ptr interface{}) { o.PMCXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PMCXML, _s_bstrPMCXml, _ptr_bstrPMCXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPolicyXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyXML == nil {
				o.PolicyXML = &oaut.String{}
			}
			if err := o.PolicyXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyXml := func(ptr interface{}) { o.PolicyXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyXML, _s_bstrPolicyXml, _ptr_bstrPolicyXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrCalendarXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrCalendarXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CalendarXML == nil {
				o.CalendarXML = &oaut.String{}
			}
			if err := o.CalendarXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrCalendarXml := func(ptr interface{}) { o.CalendarXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CalendarXML, _s_bstrCalendarXml, _ptr_bstrCalendarXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrConditionalXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConditionalXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConditionalXML == nil {
				o.ConditionalXML = &oaut.String{}
			}
			if err := o.ConditionalXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConditionalXml := func(ptr interface{}) { o.ConditionalXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConditionalXML, _s_bstrConditionalXml, _ptr_bstrConditionalXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportXMLFilesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportXMLFilesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ImportXMLFilesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ImportXMLFilesRequest structure represents the ImportXMLFiles operation request
type ImportXMLFilesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	PMCXML         *oaut.String   `idl:"name:bstrPMCXml" json:"pmc_xml"`
	PolicyXML      *oaut.String   `idl:"name:bstrPolicyXml" json:"policy_xml"`
	CalendarXML    *oaut.String   `idl:"name:bstrCalendarXml" json:"calendar_xml"`
	ConditionalXML *oaut.String   `idl:"name:bstrConditionalXml" json:"conditional_xml"`
}

func (o *ImportXMLFilesRequest) xxx_ToOp(ctx context.Context, op *xxx_ImportXMLFilesOperation) *xxx_ImportXMLFilesOperation {
	if op == nil {
		op = &xxx_ImportXMLFilesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PMCXML = o.PMCXML
	op.PolicyXML = o.PolicyXML
	op.CalendarXML = o.CalendarXML
	op.ConditionalXML = o.ConditionalXML
	return op
}

func (o *ImportXMLFilesRequest) xxx_FromOp(ctx context.Context, op *xxx_ImportXMLFilesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PMCXML = op.PMCXML
	o.PolicyXML = op.PolicyXML
	o.CalendarXML = op.CalendarXML
	o.ConditionalXML = op.ConditionalXML
}
func (o *ImportXMLFilesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ImportXMLFilesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportXMLFilesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ImportXMLFilesResponse structure represents the ImportXMLFiles operation response
type ImportXMLFilesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ImportXMLFiles return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ImportXMLFilesResponse) xxx_ToOp(ctx context.Context, op *xxx_ImportXMLFilesOperation) *xxx_ImportXMLFilesOperation {
	if op == nil {
		op = &xxx_ImportXMLFilesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ImportXMLFilesResponse) xxx_FromOp(ctx context.Context, op *xxx_ImportXMLFilesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ImportXMLFilesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ImportXMLFilesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportXMLFilesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExportXMLFilesOperation structure represents the ExportXMLFiles operation
type xxx_ExportXMLFilesOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	PMCXML         *oaut.String   `idl:"name:pbstrPMCXml" json:"pmc_xml"`
	PolicyXML      *oaut.String   `idl:"name:pbstrPolicyXml" json:"policy_xml"`
	CalendarXML    *oaut.String   `idl:"name:pbstrCalendarXml" json:"calendar_xml"`
	ConditionalXML *oaut.String   `idl:"name:pbstrConditionalXml" json:"conditional_xml"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ExportXMLFilesOperation) OpNum() int { return 10 }

func (o *xxx_ExportXMLFilesOperation) OpName() string { return "/IResourceManager/v0/ExportXMLFiles" }

func (o *xxx_ExportXMLFilesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportXMLFilesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_ExportXMLFilesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_ExportXMLFilesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportXMLFilesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrPMCXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PMCXML != nil {
			_ptr_pbstrPMCXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PMCXML != nil {
					if err := o.PMCXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PMCXML, _ptr_pbstrPMCXml); err != nil {
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
	// pbstrPolicyXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyXML != nil {
			_ptr_pbstrPolicyXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyXML != nil {
					if err := o.PolicyXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyXML, _ptr_pbstrPolicyXml); err != nil {
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
	// pbstrCalendarXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CalendarXML != nil {
			_ptr_pbstrCalendarXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CalendarXML != nil {
					if err := o.CalendarXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CalendarXML, _ptr_pbstrCalendarXml); err != nil {
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
	// pbstrConditionalXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConditionalXML != nil {
			_ptr_pbstrConditionalXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConditionalXML != nil {
					if err := o.ConditionalXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConditionalXML, _ptr_pbstrConditionalXml); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportXMLFilesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrPMCXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrPMCXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PMCXML == nil {
				o.PMCXML = &oaut.String{}
			}
			if err := o.PMCXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrPMCXml := func(ptr interface{}) { o.PMCXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PMCXML, _s_pbstrPMCXml, _ptr_pbstrPMCXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrPolicyXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrPolicyXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyXML == nil {
				o.PolicyXML = &oaut.String{}
			}
			if err := o.PolicyXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrPolicyXml := func(ptr interface{}) { o.PolicyXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyXML, _s_pbstrPolicyXml, _ptr_pbstrPolicyXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrCalendarXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrCalendarXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CalendarXML == nil {
				o.CalendarXML = &oaut.String{}
			}
			if err := o.CalendarXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrCalendarXml := func(ptr interface{}) { o.CalendarXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CalendarXML, _s_pbstrCalendarXml, _ptr_pbstrCalendarXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrConditionalXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrConditionalXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConditionalXML == nil {
				o.ConditionalXML = &oaut.String{}
			}
			if err := o.ConditionalXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrConditionalXml := func(ptr interface{}) { o.ConditionalXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConditionalXML, _s_pbstrConditionalXml, _ptr_pbstrConditionalXml); err != nil {
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

// ExportXMLFilesRequest structure represents the ExportXMLFiles operation request
type ExportXMLFilesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ExportXMLFilesRequest) xxx_ToOp(ctx context.Context, op *xxx_ExportXMLFilesOperation) *xxx_ExportXMLFilesOperation {
	if op == nil {
		op = &xxx_ExportXMLFilesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ExportXMLFilesRequest) xxx_FromOp(ctx context.Context, op *xxx_ExportXMLFilesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ExportXMLFilesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ExportXMLFilesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportXMLFilesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExportXMLFilesResponse structure represents the ExportXMLFiles operation response
type ExportXMLFilesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	PMCXML         *oaut.String   `idl:"name:pbstrPMCXml" json:"pmc_xml"`
	PolicyXML      *oaut.String   `idl:"name:pbstrPolicyXml" json:"policy_xml"`
	CalendarXML    *oaut.String   `idl:"name:pbstrCalendarXml" json:"calendar_xml"`
	ConditionalXML *oaut.String   `idl:"name:pbstrConditionalXml" json:"conditional_xml"`
	// Return: The ExportXMLFiles return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExportXMLFilesResponse) xxx_ToOp(ctx context.Context, op *xxx_ExportXMLFilesOperation) *xxx_ExportXMLFilesOperation {
	if op == nil {
		op = &xxx_ExportXMLFilesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PMCXML = o.PMCXML
	op.PolicyXML = o.PolicyXML
	op.CalendarXML = o.CalendarXML
	op.ConditionalXML = o.ConditionalXML
	op.Return = o.Return
	return op
}

func (o *ExportXMLFilesResponse) xxx_FromOp(ctx context.Context, op *xxx_ExportXMLFilesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PMCXML = op.PMCXML
	o.PolicyXML = op.PolicyXML
	o.CalendarXML = op.CalendarXML
	o.ConditionalXML = op.ConditionalXML
	o.Return = op.Return
}
func (o *ExportXMLFilesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ExportXMLFilesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportXMLFilesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RestoreXMLFilesOperation structure represents the RestoreXMLFiles operation
type xxx_RestoreXMLFilesOperation struct {
	This        *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat   `idl:"name:That" json:"that"`
	EnumRestore wsrm.RestoreMode `idl:"name:enumRestore" json:"enum_restore"`
	Return      int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_RestoreXMLFilesOperation) OpNum() int { return 11 }

func (o *xxx_RestoreXMLFilesOperation) OpName() string { return "/IResourceManager/v0/RestoreXMLFiles" }

func (o *xxx_RestoreXMLFilesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreXMLFilesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// enumRestore {in} (1:{v1_enum, alias=RESTORE_MODE}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumRestore)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreXMLFilesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// enumRestore {in} (1:{v1_enum, alias=RESTORE_MODE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumRestore)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreXMLFilesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreXMLFilesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RestoreXMLFilesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RestoreXMLFilesRequest structure represents the RestoreXMLFiles operation request
type RestoreXMLFilesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis   `idl:"name:This" json:"this"`
	EnumRestore wsrm.RestoreMode `idl:"name:enumRestore" json:"enum_restore"`
}

func (o *RestoreXMLFilesRequest) xxx_ToOp(ctx context.Context, op *xxx_RestoreXMLFilesOperation) *xxx_RestoreXMLFilesOperation {
	if op == nil {
		op = &xxx_RestoreXMLFilesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.EnumRestore = o.EnumRestore
	return op
}

func (o *RestoreXMLFilesRequest) xxx_FromOp(ctx context.Context, op *xxx_RestoreXMLFilesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EnumRestore = op.EnumRestore
}
func (o *RestoreXMLFilesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RestoreXMLFilesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreXMLFilesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RestoreXMLFilesResponse structure represents the RestoreXMLFiles operation response
type RestoreXMLFilesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RestoreXMLFiles return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RestoreXMLFilesResponse) xxx_ToOp(ctx context.Context, op *xxx_RestoreXMLFilesOperation) *xxx_RestoreXMLFilesOperation {
	if op == nil {
		op = &xxx_RestoreXMLFilesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RestoreXMLFilesResponse) xxx_FromOp(ctx context.Context, op *xxx_RestoreXMLFilesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RestoreXMLFilesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RestoreXMLFilesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreXMLFilesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDependenciesOperation structure represents the GetDependencies operation
type xxx_GetDependenciesOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ObjectName     *oaut.String    `idl:"name:bstrObjectName" json:"object_name"`
	EnumObject     wsrm.ObjectType `idl:"name:enumObject" json:"enum_object"`
	DependencyList *oaut.String    `idl:"name:pbstrDependencyList" json:"dependency_list"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDependenciesOperation) OpNum() int { return 12 }

func (o *xxx_GetDependenciesOperation) OpName() string { return "/IResourceManager/v0/GetDependencies" }

func (o *xxx_GetDependenciesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDependenciesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrObjectName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ObjectName != nil {
			_ptr_bstrObjectName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectName != nil {
					if err := o.ObjectName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectName, _ptr_bstrObjectName); err != nil {
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
	// enumObject {in} (1:{v1_enum, alias=OBJECT_TYPE}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumObject)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDependenciesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrObjectName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrObjectName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectName == nil {
				o.ObjectName = &oaut.String{}
			}
			if err := o.ObjectName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrObjectName := func(ptr interface{}) { o.ObjectName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ObjectName, _s_bstrObjectName, _ptr_bstrObjectName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// enumObject {in} (1:{v1_enum, alias=OBJECT_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumObject)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDependenciesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDependenciesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrDependencyList {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DependencyList != nil {
			_ptr_pbstrDependencyList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DependencyList != nil {
					if err := o.DependencyList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DependencyList, _ptr_pbstrDependencyList); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDependenciesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrDependencyList {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrDependencyList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DependencyList == nil {
				o.DependencyList = &oaut.String{}
			}
			if err := o.DependencyList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrDependencyList := func(ptr interface{}) { o.DependencyList = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DependencyList, _s_pbstrDependencyList, _ptr_pbstrDependencyList); err != nil {
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

// GetDependenciesRequest structure represents the GetDependencies operation request
type GetDependenciesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	ObjectName *oaut.String    `idl:"name:bstrObjectName" json:"object_name"`
	EnumObject wsrm.ObjectType `idl:"name:enumObject" json:"enum_object"`
}

func (o *GetDependenciesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDependenciesOperation) *xxx_GetDependenciesOperation {
	if op == nil {
		op = &xxx_GetDependenciesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectName = o.ObjectName
	op.EnumObject = o.EnumObject
	return op
}

func (o *GetDependenciesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDependenciesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectName = op.ObjectName
	o.EnumObject = op.EnumObject
}
func (o *GetDependenciesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDependenciesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDependenciesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDependenciesResponse structure represents the GetDependencies operation response
type GetDependenciesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	DependencyList *oaut.String   `idl:"name:pbstrDependencyList" json:"dependency_list"`
	// Return: The GetDependencies return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDependenciesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDependenciesOperation) *xxx_GetDependenciesOperation {
	if op == nil {
		op = &xxx_GetDependenciesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DependencyList = o.DependencyList
	op.Return = o.Return
	return op
}

func (o *GetDependenciesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDependenciesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DependencyList = op.DependencyList
	o.Return = op.Return
}
func (o *GetDependenciesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDependenciesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDependenciesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetServiceListOperation structure represents the GetServiceList operation
type xxx_GetServiceListOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServiceList *oaut.String   `idl:"name:pbstrServiceList" json:"service_list"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServiceListOperation) OpNum() int { return 13 }

func (o *xxx_GetServiceListOperation) OpName() string { return "/IResourceManager/v0/GetServiceList" }

func (o *xxx_GetServiceListOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServiceListOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_GetServiceListOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_GetServiceListOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServiceListOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrServiceList {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ServiceList != nil {
			_ptr_pbstrServiceList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServiceList != nil {
					if err := o.ServiceList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServiceList, _ptr_pbstrServiceList); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServiceListOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrServiceList {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrServiceList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServiceList == nil {
				o.ServiceList = &oaut.String{}
			}
			if err := o.ServiceList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrServiceList := func(ptr interface{}) { o.ServiceList = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ServiceList, _s_pbstrServiceList, _ptr_pbstrServiceList); err != nil {
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

// GetServiceListRequest structure represents the GetServiceList operation request
type GetServiceListRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetServiceListRequest) xxx_ToOp(ctx context.Context, op *xxx_GetServiceListOperation) *xxx_GetServiceListOperation {
	if op == nil {
		op = &xxx_GetServiceListOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetServiceListRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServiceListOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetServiceListRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetServiceListRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServiceListOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServiceListResponse structure represents the GetServiceList operation response
type GetServiceListResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServiceList *oaut.String   `idl:"name:pbstrServiceList" json:"service_list"`
	// Return: The GetServiceList return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetServiceListResponse) xxx_ToOp(ctx context.Context, op *xxx_GetServiceListOperation) *xxx_GetServiceListOperation {
	if op == nil {
		op = &xxx_GetServiceListOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ServiceList = o.ServiceList
	op.Return = o.Return
	return op
}

func (o *GetServiceListResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServiceListOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ServiceList = op.ServiceList
	o.Return = op.Return
}
func (o *GetServiceListResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetServiceListResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServiceListOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIISAppPoolNamesOperation structure represents the GetIISAppPoolNames operation
type xxx_GetIISAppPoolNamesOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	IISAppPoolList  *oaut.String   `idl:"name:pbstrIISAppPoolList" json:"iis_app_pool_list"`
	SystemDirectory *oaut.String   `idl:"name:pbstrSystemDirectory" json:"system_directory"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIISAppPoolNamesOperation) OpNum() int { return 14 }

func (o *xxx_GetIISAppPoolNamesOperation) OpName() string {
	return "/IResourceManager/v0/GetIISAppPoolNames"
}

func (o *xxx_GetIISAppPoolNamesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIISAppPoolNamesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_GetIISAppPoolNamesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_GetIISAppPoolNamesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIISAppPoolNamesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrIISAppPoolList {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.IISAppPoolList != nil {
			_ptr_pbstrIISAppPoolList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.IISAppPoolList != nil {
					if err := o.IISAppPoolList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.IISAppPoolList, _ptr_pbstrIISAppPoolList); err != nil {
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
	// pbstrSystemDirectory {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SystemDirectory != nil {
			_ptr_pbstrSystemDirectory := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SystemDirectory != nil {
					if err := o.SystemDirectory.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SystemDirectory, _ptr_pbstrSystemDirectory); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIISAppPoolNamesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrIISAppPoolList {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrIISAppPoolList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.IISAppPoolList == nil {
				o.IISAppPoolList = &oaut.String{}
			}
			if err := o.IISAppPoolList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrIISAppPoolList := func(ptr interface{}) { o.IISAppPoolList = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.IISAppPoolList, _s_pbstrIISAppPoolList, _ptr_pbstrIISAppPoolList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrSystemDirectory {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrSystemDirectory := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SystemDirectory == nil {
				o.SystemDirectory = &oaut.String{}
			}
			if err := o.SystemDirectory.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrSystemDirectory := func(ptr interface{}) { o.SystemDirectory = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SystemDirectory, _s_pbstrSystemDirectory, _ptr_pbstrSystemDirectory); err != nil {
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

// GetIISAppPoolNamesRequest structure represents the GetIISAppPoolNames operation request
type GetIISAppPoolNamesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIISAppPoolNamesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIISAppPoolNamesOperation) *xxx_GetIISAppPoolNamesOperation {
	if op == nil {
		op = &xxx_GetIISAppPoolNamesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIISAppPoolNamesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIISAppPoolNamesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIISAppPoolNamesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIISAppPoolNamesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIISAppPoolNamesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIISAppPoolNamesResponse structure represents the GetIISAppPoolNames operation response
type GetIISAppPoolNamesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	IISAppPoolList  *oaut.String   `idl:"name:pbstrIISAppPoolList" json:"iis_app_pool_list"`
	SystemDirectory *oaut.String   `idl:"name:pbstrSystemDirectory" json:"system_directory"`
	// Return: The GetIISAppPoolNames return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIISAppPoolNamesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIISAppPoolNamesOperation) *xxx_GetIISAppPoolNamesOperation {
	if op == nil {
		op = &xxx_GetIISAppPoolNamesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IISAppPoolList = o.IISAppPoolList
	op.SystemDirectory = o.SystemDirectory
	op.Return = o.Return
	return op
}

func (o *GetIISAppPoolNamesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIISAppPoolNamesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IISAppPoolList = op.IISAppPoolList
	o.SystemDirectory = op.SystemDirectory
	o.Return = op.Return
}
func (o *GetIISAppPoolNamesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIISAppPoolNamesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIISAppPoolNamesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetServerNameOperation structure represents the GetServerName operation
type xxx_GetServerNameOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServerName *oaut.String   `idl:"name:pbstrServerName" json:"server_name"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServerNameOperation) OpNum() int { return 15 }

func (o *xxx_GetServerNameOperation) OpName() string { return "/IResourceManager/v0/GetServerName" }

func (o *xxx_GetServerNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_GetServerNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_GetServerNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrServerName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ServerName != nil {
			_ptr_pbstrServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServerName != nil {
					if err := o.ServerName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pbstrServerName); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrServerName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServerName == nil {
				o.ServerName = &oaut.String{}
			}
			if err := o.ServerName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrServerName := func(ptr interface{}) { o.ServerName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ServerName, _s_pbstrServerName, _ptr_pbstrServerName); err != nil {
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

// GetServerNameRequest structure represents the GetServerName operation request
type GetServerNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetServerNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetServerNameOperation) *xxx_GetServerNameOperation {
	if op == nil {
		op = &xxx_GetServerNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetServerNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServerNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetServerNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetServerNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServerNameResponse structure represents the GetServerName operation response
type GetServerNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServerName *oaut.String   `idl:"name:pbstrServerName" json:"server_name"`
	// Return: The GetServerName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetServerNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetServerNameOperation) *xxx_GetServerNameOperation {
	if op == nil {
		op = &xxx_GetServerNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ServerName = o.ServerName
	op.Return = o.Return
	return op
}

func (o *GetServerNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServerNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ServerName = op.ServerName
	o.Return = op.Return
}
func (o *GetServerNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetServerNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCurrentMemoryOperation structure represents the GetCurrentMemory operation
type xxx_GetCurrentMemoryOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	CurrentMemory uint64         `idl:"name:pdwCurrMemory" json:"current_memory"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCurrentMemoryOperation) OpNum() int { return 16 }

func (o *xxx_GetCurrentMemoryOperation) OpName() string {
	return "/IResourceManager/v0/GetCurrentMemory"
}

func (o *xxx_GetCurrentMemoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCurrentMemoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_GetCurrentMemoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_GetCurrentMemoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCurrentMemoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwCurrMemory {out} (1:{pointer=ref}*(1))(2:{alias=DWORD64}(uint64))
	{
		if err := w.WriteData(o.CurrentMemory); err != nil {
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

func (o *xxx_GetCurrentMemoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwCurrMemory {out} (1:{pointer=ref}*(1))(2:{alias=DWORD64}(uint64))
	{
		if err := w.ReadData(&o.CurrentMemory); err != nil {
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

// GetCurrentMemoryRequest structure represents the GetCurrentMemory operation request
type GetCurrentMemoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCurrentMemoryRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCurrentMemoryOperation) *xxx_GetCurrentMemoryOperation {
	if op == nil {
		op = &xxx_GetCurrentMemoryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCurrentMemoryRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCurrentMemoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCurrentMemoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCurrentMemoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCurrentMemoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCurrentMemoryResponse structure represents the GetCurrentMemory operation response
type GetCurrentMemoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	CurrentMemory uint64         `idl:"name:pdwCurrMemory" json:"current_memory"`
	// Return: The GetCurrentMemory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCurrentMemoryResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCurrentMemoryOperation) *xxx_GetCurrentMemoryOperation {
	if op == nil {
		op = &xxx_GetCurrentMemoryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.CurrentMemory = o.CurrentMemory
	op.Return = o.Return
	return op
}

func (o *GetCurrentMemoryResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCurrentMemoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CurrentMemory = op.CurrentMemory
	o.Return = op.Return
}
func (o *GetCurrentMemoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCurrentMemoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCurrentMemoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
