package ifsrmstoragemoduledefinition

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
	ifsrmpipelinemoduledefinition "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmpipelinemoduledefinition/v0"
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
	_ = ifsrmpipelinemoduledefinition.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmStorageModuleDefinition interface identifier 15a81350-497d-4aba-80e9-d4dbcc5521fe
	StorageModuleDefinitionIID = &dcom.IID{Data1: 0x15a81350, Data2: 0x497d, Data3: 0x4aba, Data4: []byte{0x80, 0xe9, 0xd4, 0xdb, 0xcc, 0x55, 0x21, 0xfe}}
	// Syntax UUID
	StorageModuleDefinitionSyntaxUUID = &uuid.UUID{TimeLow: 0x15a81350, TimeMid: 0x497d, TimeHiAndVersion: 0x4aba, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0xe9, Node: [6]uint8{0xd4, 0xdb, 0xcc, 0x55, 0x21, 0xfe}}
	// Syntax ID
	StorageModuleDefinitionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: StorageModuleDefinitionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmStorageModuleDefinition interface.
type StorageModuleDefinitionClient interface {

	// IFsrmPipelineModuleDefinition retrieval method.
	PipelineModuleDefinition() ifsrmpipelinemoduledefinition.PipelineModuleDefinitionClient

	// Capabilities operation.
	GetCapabilities(context.Context, *GetCapabilitiesRequest, ...dcerpc.CallOption) (*GetCapabilitiesResponse, error)

	// Capabilities operation.
	SetCapabilities(context.Context, *SetCapabilitiesRequest, ...dcerpc.CallOption) (*SetCapabilitiesResponse, error)

	// StorageType operation.
	GetStorageType(context.Context, *GetStorageTypeRequest, ...dcerpc.CallOption) (*GetStorageTypeResponse, error)

	// StorageType operation.
	SetStorageType(context.Context, *SetStorageTypeRequest, ...dcerpc.CallOption) (*SetStorageTypeResponse, error)

	// UpdatesFileContent operation.
	GetUpdatesFileContent(context.Context, *GetUpdatesFileContentRequest, ...dcerpc.CallOption) (*GetUpdatesFileContentResponse, error)

	// UpdatesFileContent operation.
	SetUpdatesFileContent(context.Context, *SetUpdatesFileContentRequest, ...dcerpc.CallOption) (*SetUpdatesFileContentResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) StorageModuleDefinitionClient
}

type xxx_DefaultStorageModuleDefinitionClient struct {
	ifsrmpipelinemoduledefinition.PipelineModuleDefinitionClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultStorageModuleDefinitionClient) PipelineModuleDefinition() ifsrmpipelinemoduledefinition.PipelineModuleDefinitionClient {
	return o.PipelineModuleDefinitionClient
}

func (o *xxx_DefaultStorageModuleDefinitionClient) GetCapabilities(ctx context.Context, in *GetCapabilitiesRequest, opts ...dcerpc.CallOption) (*GetCapabilitiesResponse, error) {
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
	out := &GetCapabilitiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStorageModuleDefinitionClient) SetCapabilities(ctx context.Context, in *SetCapabilitiesRequest, opts ...dcerpc.CallOption) (*SetCapabilitiesResponse, error) {
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
	out := &SetCapabilitiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStorageModuleDefinitionClient) GetStorageType(ctx context.Context, in *GetStorageTypeRequest, opts ...dcerpc.CallOption) (*GetStorageTypeResponse, error) {
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
	out := &GetStorageTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStorageModuleDefinitionClient) SetStorageType(ctx context.Context, in *SetStorageTypeRequest, opts ...dcerpc.CallOption) (*SetStorageTypeResponse, error) {
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
	out := &SetStorageTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStorageModuleDefinitionClient) GetUpdatesFileContent(ctx context.Context, in *GetUpdatesFileContentRequest, opts ...dcerpc.CallOption) (*GetUpdatesFileContentResponse, error) {
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
	out := &GetUpdatesFileContentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStorageModuleDefinitionClient) SetUpdatesFileContent(ctx context.Context, in *SetUpdatesFileContentRequest, opts ...dcerpc.CallOption) (*SetUpdatesFileContentResponse, error) {
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
	out := &SetUpdatesFileContentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStorageModuleDefinitionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultStorageModuleDefinitionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultStorageModuleDefinitionClient) IPID(ctx context.Context, ipid *dcom.IPID) StorageModuleDefinitionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultStorageModuleDefinitionClient{
		PipelineModuleDefinitionClient: o.PipelineModuleDefinitionClient.IPID(ctx, ipid),
		cc:                             o.cc,
		ipid:                           ipid,
	}
}

func NewStorageModuleDefinitionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (StorageModuleDefinitionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(StorageModuleDefinitionSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmpipelinemoduledefinition.NewPipelineModuleDefinitionClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultStorageModuleDefinitionClient{
		PipelineModuleDefinitionClient: base,
		cc:                             cc,
		ipid:                           ipid,
	}, nil
}

// xxx_GetCapabilitiesOperation structure represents the Capabilities operation
type xxx_GetCapabilitiesOperation struct {
	This         *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat         `idl:"name:That" json:"that"`
	Capabilities fsrm.StorageModuleCaps `idl:"name:capabilities" json:"capabilities"`
	Return       int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCapabilitiesOperation) OpNum() int { return 31 }

func (o *xxx_GetCapabilitiesOperation) OpName() string {
	return "/IFsrmStorageModuleDefinition/v0/Capabilities"
}

func (o *xxx_GetCapabilitiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCapabilitiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCapabilitiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCapabilitiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCapabilitiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// capabilities {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmStorageModuleCaps}(enum))
	{
		if err := w.WriteEnum(uint16(o.Capabilities)); err != nil {
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

func (o *xxx_GetCapabilitiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// capabilities {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmStorageModuleCaps}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Capabilities)); err != nil {
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

// GetCapabilitiesRequest structure represents the Capabilities operation request
type GetCapabilitiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCapabilitiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCapabilitiesOperation) *xxx_GetCapabilitiesOperation {
	if op == nil {
		op = &xxx_GetCapabilitiesOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetCapabilitiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCapabilitiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCapabilitiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCapabilitiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCapabilitiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCapabilitiesResponse structure represents the Capabilities operation response
type GetCapabilitiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat         `idl:"name:That" json:"that"`
	Capabilities fsrm.StorageModuleCaps `idl:"name:capabilities" json:"capabilities"`
	// Return: The Capabilities return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCapabilitiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCapabilitiesOperation) *xxx_GetCapabilitiesOperation {
	if op == nil {
		op = &xxx_GetCapabilitiesOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Capabilities = op.Capabilities
	o.Return = op.Return
	return op
}

func (o *GetCapabilitiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCapabilitiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Capabilities = op.Capabilities
	o.Return = op.Return
}
func (o *GetCapabilitiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCapabilitiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCapabilitiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCapabilitiesOperation structure represents the Capabilities operation
type xxx_SetCapabilitiesOperation struct {
	This         *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat         `idl:"name:That" json:"that"`
	Capabilities fsrm.StorageModuleCaps `idl:"name:capabilities" json:"capabilities"`
	Return       int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCapabilitiesOperation) OpNum() int { return 32 }

func (o *xxx_SetCapabilitiesOperation) OpName() string {
	return "/IFsrmStorageModuleDefinition/v0/Capabilities"
}

func (o *xxx_SetCapabilitiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCapabilitiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// capabilities {in} (1:{alias=FsrmStorageModuleCaps}(enum))
	{
		if err := w.WriteEnum(uint16(o.Capabilities)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCapabilitiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// capabilities {in} (1:{alias=FsrmStorageModuleCaps}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Capabilities)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCapabilitiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCapabilitiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCapabilitiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetCapabilitiesRequest structure represents the Capabilities operation request
type SetCapabilitiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis         `idl:"name:This" json:"this"`
	Capabilities fsrm.StorageModuleCaps `idl:"name:capabilities" json:"capabilities"`
}

func (o *SetCapabilitiesRequest) xxx_ToOp(ctx context.Context, op *xxx_SetCapabilitiesOperation) *xxx_SetCapabilitiesOperation {
	if op == nil {
		op = &xxx_SetCapabilitiesOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Capabilities = op.Capabilities
	return op
}

func (o *SetCapabilitiesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCapabilitiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Capabilities = op.Capabilities
}
func (o *SetCapabilitiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetCapabilitiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCapabilitiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCapabilitiesResponse structure represents the Capabilities operation response
type SetCapabilitiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Capabilities return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCapabilitiesResponse) xxx_ToOp(ctx context.Context, op *xxx_SetCapabilitiesOperation) *xxx_SetCapabilitiesOperation {
	if op == nil {
		op = &xxx_SetCapabilitiesOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *SetCapabilitiesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCapabilitiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCapabilitiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetCapabilitiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCapabilitiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetStorageTypeOperation structure represents the StorageType operation
type xxx_GetStorageTypeOperation struct {
	This        *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat         `idl:"name:That" json:"that"`
	StorageType fsrm.StorageModuleType `idl:"name:storageType" json:"storage_type"`
	Return      int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetStorageTypeOperation) OpNum() int { return 33 }

func (o *xxx_GetStorageTypeOperation) OpName() string {
	return "/IFsrmStorageModuleDefinition/v0/StorageType"
}

func (o *xxx_GetStorageTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStorageTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetStorageTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetStorageTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStorageTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// storageType {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmStorageModuleType}(enum))
	{
		if err := w.WriteEnum(uint16(o.StorageType)); err != nil {
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

func (o *xxx_GetStorageTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// storageType {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmStorageModuleType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.StorageType)); err != nil {
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

// GetStorageTypeRequest structure represents the StorageType operation request
type GetStorageTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetStorageTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetStorageTypeOperation) *xxx_GetStorageTypeOperation {
	if op == nil {
		op = &xxx_GetStorageTypeOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetStorageTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetStorageTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetStorageTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetStorageTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStorageTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetStorageTypeResponse structure represents the StorageType operation response
type GetStorageTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat         `idl:"name:That" json:"that"`
	StorageType fsrm.StorageModuleType `idl:"name:storageType" json:"storage_type"`
	// Return: The StorageType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetStorageTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetStorageTypeOperation) *xxx_GetStorageTypeOperation {
	if op == nil {
		op = &xxx_GetStorageTypeOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.StorageType = op.StorageType
	o.Return = op.Return
	return op
}

func (o *GetStorageTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetStorageTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.StorageType = op.StorageType
	o.Return = op.Return
}
func (o *GetStorageTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetStorageTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStorageTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetStorageTypeOperation structure represents the StorageType operation
type xxx_SetStorageTypeOperation struct {
	This        *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat         `idl:"name:That" json:"that"`
	StorageType fsrm.StorageModuleType `idl:"name:storageType" json:"storage_type"`
	Return      int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetStorageTypeOperation) OpNum() int { return 34 }

func (o *xxx_SetStorageTypeOperation) OpName() string {
	return "/IFsrmStorageModuleDefinition/v0/StorageType"
}

func (o *xxx_SetStorageTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStorageTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// storageType {in} (1:{alias=FsrmStorageModuleType}(enum))
	{
		if err := w.WriteEnum(uint16(o.StorageType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStorageTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// storageType {in} (1:{alias=FsrmStorageModuleType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.StorageType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStorageTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStorageTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetStorageTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetStorageTypeRequest structure represents the StorageType operation request
type SetStorageTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis         `idl:"name:This" json:"this"`
	StorageType fsrm.StorageModuleType `idl:"name:storageType" json:"storage_type"`
}

func (o *SetStorageTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetStorageTypeOperation) *xxx_SetStorageTypeOperation {
	if op == nil {
		op = &xxx_SetStorageTypeOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.StorageType = op.StorageType
	return op
}

func (o *SetStorageTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetStorageTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.StorageType = op.StorageType
}
func (o *SetStorageTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetStorageTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetStorageTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetStorageTypeResponse structure represents the StorageType operation response
type SetStorageTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The StorageType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetStorageTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetStorageTypeOperation) *xxx_SetStorageTypeOperation {
	if op == nil {
		op = &xxx_SetStorageTypeOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *SetStorageTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetStorageTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetStorageTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetStorageTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetStorageTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetUpdatesFileContentOperation structure represents the UpdatesFileContent operation
type xxx_GetUpdatesFileContentOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	UpdatesFileContent int16          `idl:"name:updatesFileContent" json:"updates_file_content"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUpdatesFileContentOperation) OpNum() int { return 35 }

func (o *xxx_GetUpdatesFileContentOperation) OpName() string {
	return "/IFsrmStorageModuleDefinition/v0/UpdatesFileContent"
}

func (o *xxx_GetUpdatesFileContentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUpdatesFileContentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetUpdatesFileContentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetUpdatesFileContentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUpdatesFileContentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// updatesFileContent {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.UpdatesFileContent); err != nil {
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

func (o *xxx_GetUpdatesFileContentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// updatesFileContent {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.UpdatesFileContent); err != nil {
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

// GetUpdatesFileContentRequest structure represents the UpdatesFileContent operation request
type GetUpdatesFileContentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetUpdatesFileContentRequest) xxx_ToOp(ctx context.Context, op *xxx_GetUpdatesFileContentOperation) *xxx_GetUpdatesFileContentOperation {
	if op == nil {
		op = &xxx_GetUpdatesFileContentOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetUpdatesFileContentRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUpdatesFileContentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetUpdatesFileContentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetUpdatesFileContentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUpdatesFileContentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUpdatesFileContentResponse structure represents the UpdatesFileContent operation response
type GetUpdatesFileContentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	UpdatesFileContent int16          `idl:"name:updatesFileContent" json:"updates_file_content"`
	// Return: The UpdatesFileContent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUpdatesFileContentResponse) xxx_ToOp(ctx context.Context, op *xxx_GetUpdatesFileContentOperation) *xxx_GetUpdatesFileContentOperation {
	if op == nil {
		op = &xxx_GetUpdatesFileContentOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.UpdatesFileContent = op.UpdatesFileContent
	o.Return = op.Return
	return op
}

func (o *GetUpdatesFileContentResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUpdatesFileContentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.UpdatesFileContent = op.UpdatesFileContent
	o.Return = op.Return
}
func (o *GetUpdatesFileContentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetUpdatesFileContentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUpdatesFileContentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetUpdatesFileContentOperation structure represents the UpdatesFileContent operation
type xxx_SetUpdatesFileContentOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	UpdatesFileContent int16          `idl:"name:updatesFileContent" json:"updates_file_content"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetUpdatesFileContentOperation) OpNum() int { return 36 }

func (o *xxx_SetUpdatesFileContentOperation) OpName() string {
	return "/IFsrmStorageModuleDefinition/v0/UpdatesFileContent"
}

func (o *xxx_SetUpdatesFileContentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUpdatesFileContentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// updatesFileContent {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.UpdatesFileContent); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUpdatesFileContentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// updatesFileContent {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.UpdatesFileContent); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUpdatesFileContentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUpdatesFileContentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetUpdatesFileContentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetUpdatesFileContentRequest structure represents the UpdatesFileContent operation request
type SetUpdatesFileContentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	UpdatesFileContent int16          `idl:"name:updatesFileContent" json:"updates_file_content"`
}

func (o *SetUpdatesFileContentRequest) xxx_ToOp(ctx context.Context, op *xxx_SetUpdatesFileContentOperation) *xxx_SetUpdatesFileContentOperation {
	if op == nil {
		op = &xxx_SetUpdatesFileContentOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.UpdatesFileContent = op.UpdatesFileContent
	return op
}

func (o *SetUpdatesFileContentRequest) xxx_FromOp(ctx context.Context, op *xxx_SetUpdatesFileContentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.UpdatesFileContent = op.UpdatesFileContent
}
func (o *SetUpdatesFileContentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetUpdatesFileContentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetUpdatesFileContentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetUpdatesFileContentResponse structure represents the UpdatesFileContent operation response
type SetUpdatesFileContentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The UpdatesFileContent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetUpdatesFileContentResponse) xxx_ToOp(ctx context.Context, op *xxx_SetUpdatesFileContentOperation) *xxx_SetUpdatesFileContentOperation {
	if op == nil {
		op = &xxx_SetUpdatesFileContentOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *SetUpdatesFileContentResponse) xxx_FromOp(ctx context.Context, op *xxx_SetUpdatesFileContentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetUpdatesFileContentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetUpdatesFileContentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetUpdatesFileContentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
