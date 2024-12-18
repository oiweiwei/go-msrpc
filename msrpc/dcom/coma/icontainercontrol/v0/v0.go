package icontainercontrol

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// IContainerControl interface identifier 3f3b1b86-dbbe-11d1-9da6-00805f85cfe3
	ContainerControlIID = &dcom.IID{Data1: 0x3f3b1b86, Data2: 0xdbbe, Data3: 0x11d1, Data4: []byte{0x9d, 0xa6, 0x00, 0x80, 0x5f, 0x85, 0xcf, 0xe3}}
	// Syntax UUID
	ContainerControlSyntaxUUID = &uuid.UUID{TimeLow: 0x3f3b1b86, TimeMid: 0xdbbe, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0xa6, Node: [6]uint8{0x0, 0x80, 0x5f, 0x85, 0xcf, 0xe3}}
	// Syntax ID
	ContainerControlSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ContainerControlSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IContainerControl interface.
type ContainerControlClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to create an instance container for a conglomeration.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	CreateContainer(context.Context, *CreateContainerRequest, ...dcerpc.CallOption) (*CreateContainerResponse, error)

	// This method is called by a client to shut down all instance containers for a conglomeration.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	ShutdownContainers(context.Context, *ShutdownContainersRequest, ...dcerpc.CallOption) (*ShutdownContainersResponse, error)

	// This method is called by a client to perform implementation-specific repairs on the
	// server's catalog.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	RefreshComponents(context.Context, *RefreshComponentsRequest, ...dcerpc.CallOption) (*RefreshComponentsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ContainerControlClient
}

type xxx_DefaultContainerControlClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultContainerControlClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultContainerControlClient) CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...dcerpc.CallOption) (*CreateContainerResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &CreateContainerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultContainerControlClient) ShutdownContainers(ctx context.Context, in *ShutdownContainersRequest, opts ...dcerpc.CallOption) (*ShutdownContainersResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &ShutdownContainersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultContainerControlClient) RefreshComponents(ctx context.Context, in *RefreshComponentsRequest, opts ...dcerpc.CallOption) (*RefreshComponentsResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &RefreshComponentsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultContainerControlClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultContainerControlClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultContainerControlClient) IPID(ctx context.Context, ipid *dcom.IPID) ContainerControlClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultContainerControlClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewContainerControlClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ContainerControlClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ContainerControlSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultContainerControlClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CreateContainerOperation structure represents the CreateContainer operation
type xxx_CreateContainerOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConglomerationID *dtyp.GUID     `idl:"name:pConglomerationIdentifier" json:"conglomeration_id"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateContainerOperation) OpNum() int { return 3 }

func (o *xxx_CreateContainerOperation) OpName() string {
	return "/IContainerControl/v0/CreateContainer"
}

func (o *xxx_CreateContainerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateContainerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pConglomerationIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ConglomerationID != nil {
			if err := o.ConglomerationID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreateContainerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pConglomerationIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ConglomerationID == nil {
			o.ConglomerationID = &dtyp.GUID{}
		}
		if err := o.ConglomerationID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateContainerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateContainerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateContainerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateContainerRequest structure represents the CreateContainer operation request
type CreateContainerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pConglomerationIdentifier: The conglomeration identifier of a conglomeration.
	ConglomerationID *dtyp.GUID `idl:"name:pConglomerationIdentifier" json:"conglomeration_id"`
}

func (o *CreateContainerRequest) xxx_ToOp(ctx context.Context) *xxx_CreateContainerOperation {
	if o == nil {
		return &xxx_CreateContainerOperation{}
	}
	return &xxx_CreateContainerOperation{
		This:             o.This,
		ConglomerationID: o.ConglomerationID,
	}
}

func (o *CreateContainerRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateContainerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConglomerationID = op.ConglomerationID
}
func (o *CreateContainerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateContainerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateContainerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateContainerResponse structure represents the CreateContainer operation response
type CreateContainerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateContainer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateContainerResponse) xxx_ToOp(ctx context.Context) *xxx_CreateContainerOperation {
	if o == nil {
		return &xxx_CreateContainerOperation{}
	}
	return &xxx_CreateContainerOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CreateContainerResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateContainerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateContainerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateContainerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateContainerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ShutdownContainersOperation structure represents the ShutdownContainers operation
type xxx_ShutdownContainersOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConglomerationID *dtyp.GUID     `idl:"name:pConglomerationIdentifier" json:"conglomeration_id"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ShutdownContainersOperation) OpNum() int { return 4 }

func (o *xxx_ShutdownContainersOperation) OpName() string {
	return "/IContainerControl/v0/ShutdownContainers"
}

func (o *xxx_ShutdownContainersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShutdownContainersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pConglomerationIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ConglomerationID != nil {
			if err := o.ConglomerationID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ShutdownContainersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pConglomerationIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ConglomerationID == nil {
			o.ConglomerationID = &dtyp.GUID{}
		}
		if err := o.ConglomerationID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShutdownContainersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShutdownContainersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ShutdownContainersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ShutdownContainersRequest structure represents the ShutdownContainers operation request
type ShutdownContainersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pConglomerationIdentifier: The conglomeration identifier of a conglomeration.
	ConglomerationID *dtyp.GUID `idl:"name:pConglomerationIdentifier" json:"conglomeration_id"`
}

func (o *ShutdownContainersRequest) xxx_ToOp(ctx context.Context) *xxx_ShutdownContainersOperation {
	if o == nil {
		return &xxx_ShutdownContainersOperation{}
	}
	return &xxx_ShutdownContainersOperation{
		This:             o.This,
		ConglomerationID: o.ConglomerationID,
	}
}

func (o *ShutdownContainersRequest) xxx_FromOp(ctx context.Context, op *xxx_ShutdownContainersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConglomerationID = op.ConglomerationID
}
func (o *ShutdownContainersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ShutdownContainersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ShutdownContainersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ShutdownContainersResponse structure represents the ShutdownContainers operation response
type ShutdownContainersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ShutdownContainers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ShutdownContainersResponse) xxx_ToOp(ctx context.Context) *xxx_ShutdownContainersOperation {
	if o == nil {
		return &xxx_ShutdownContainersOperation{}
	}
	return &xxx_ShutdownContainersOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ShutdownContainersResponse) xxx_FromOp(ctx context.Context, op *xxx_ShutdownContainersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ShutdownContainersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ShutdownContainersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ShutdownContainersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RefreshComponentsOperation structure represents the RefreshComponents operation
type xxx_RefreshComponentsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RefreshComponentsOperation) OpNum() int { return 5 }

func (o *xxx_RefreshComponentsOperation) OpName() string {
	return "/IContainerControl/v0/RefreshComponents"
}

func (o *xxx_RefreshComponentsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshComponentsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RefreshComponentsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_RefreshComponentsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshComponentsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RefreshComponentsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RefreshComponentsRequest structure represents the RefreshComponents operation request
type RefreshComponentsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RefreshComponentsRequest) xxx_ToOp(ctx context.Context) *xxx_RefreshComponentsOperation {
	if o == nil {
		return &xxx_RefreshComponentsOperation{}
	}
	return &xxx_RefreshComponentsOperation{
		This: o.This,
	}
}

func (o *RefreshComponentsRequest) xxx_FromOp(ctx context.Context, op *xxx_RefreshComponentsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RefreshComponentsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RefreshComponentsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshComponentsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RefreshComponentsResponse structure represents the RefreshComponents operation response
type RefreshComponentsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RefreshComponents return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RefreshComponentsResponse) xxx_ToOp(ctx context.Context) *xxx_RefreshComponentsOperation {
	if o == nil {
		return &xxx_RefreshComponentsOperation{}
	}
	return &xxx_RefreshComponentsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RefreshComponentsResponse) xxx_FromOp(ctx context.Context, op *xxx_RefreshComponentsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RefreshComponentsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RefreshComponentsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshComponentsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
