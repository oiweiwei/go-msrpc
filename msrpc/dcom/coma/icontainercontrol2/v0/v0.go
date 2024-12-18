package icontainercontrol2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	coma "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = coma.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// IContainerControl2 interface identifier 6c935649-30a6-4211-8687-c4c83e5fe1c7
	ContainerControl2IID = &dcom.IID{Data1: 0x6c935649, Data2: 0x30a6, Data3: 0x4211, Data4: []byte{0x86, 0x87, 0xc4, 0xc8, 0x3e, 0x5f, 0xe1, 0xc7}}
	// Syntax UUID
	ContainerControl2SyntaxUUID = &uuid.UUID{TimeLow: 0x6c935649, TimeMid: 0x30a6, TimeHiAndVersion: 0x4211, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x87, Node: [6]uint8{0xc4, 0xc8, 0x3e, 0x5f, 0xe1, 0xc7}}
	// Syntax ID
	ContainerControl2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ContainerControl2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IContainerControl2 interface.
type ContainerControl2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to shut down an instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	ShutdownContainer(context.Context, *ShutdownContainerRequest, ...dcerpc.CallOption) (*ShutdownContainerResponse, error)

	// This method is called by a client to pause an instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	PauseContainer(context.Context, *PauseContainerRequest, ...dcerpc.CallOption) (*PauseContainerResponse, error)

	// This method is called by a client to resume a paused instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	ResumeContainer(context.Context, *ResumeContainerRequest, ...dcerpc.CallOption) (*ResumeContainerResponse, error)

	// This method is called by a client to determine if an instance container is paused.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	IsContainerPaused(context.Context, *IsContainerPausedRequest, ...dcerpc.CallOption) (*IsContainerPausedResponse, error)

	// This method is called by a client to get a list of instance containers for a conglomeration,
	// or to get a list of all running instance containers.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	GetRunningContainers(context.Context, *GetRunningContainersRequest, ...dcerpc.CallOption) (*GetRunningContainersResponse, error)

	// This method is called by a client to find the instance container for a process ID.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	GetContainerIDFromProcessID(context.Context, *GetContainerIDFromProcessIDRequest, ...dcerpc.CallOption) (*GetContainerIDFromProcessIDResponse, error)

	// This method is called by a client to recycle an instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	RecycleContainer(context.Context, *RecycleContainerRequest, ...dcerpc.CallOption) (*RecycleContainerResponse, error)

	// This method is called by a client to find the container identifier for the single
	// instance container for a conglomeration.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	GetContainerIDFromConglomerationID(context.Context, *GetContainerIDFromConglomerationIDRequest, ...dcerpc.CallOption) (*GetContainerIDFromConglomerationIDResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ContainerControl2Client
}

type xxx_DefaultContainerControl2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultContainerControl2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultContainerControl2Client) ShutdownContainer(ctx context.Context, in *ShutdownContainerRequest, opts ...dcerpc.CallOption) (*ShutdownContainerResponse, error) {
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
	out := &ShutdownContainerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultContainerControl2Client) PauseContainer(ctx context.Context, in *PauseContainerRequest, opts ...dcerpc.CallOption) (*PauseContainerResponse, error) {
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
	out := &PauseContainerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultContainerControl2Client) ResumeContainer(ctx context.Context, in *ResumeContainerRequest, opts ...dcerpc.CallOption) (*ResumeContainerResponse, error) {
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
	out := &ResumeContainerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultContainerControl2Client) IsContainerPaused(ctx context.Context, in *IsContainerPausedRequest, opts ...dcerpc.CallOption) (*IsContainerPausedResponse, error) {
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
	out := &IsContainerPausedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultContainerControl2Client) GetRunningContainers(ctx context.Context, in *GetRunningContainersRequest, opts ...dcerpc.CallOption) (*GetRunningContainersResponse, error) {
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
	out := &GetRunningContainersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultContainerControl2Client) GetContainerIDFromProcessID(ctx context.Context, in *GetContainerIDFromProcessIDRequest, opts ...dcerpc.CallOption) (*GetContainerIDFromProcessIDResponse, error) {
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
	out := &GetContainerIDFromProcessIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultContainerControl2Client) RecycleContainer(ctx context.Context, in *RecycleContainerRequest, opts ...dcerpc.CallOption) (*RecycleContainerResponse, error) {
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
	out := &RecycleContainerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultContainerControl2Client) GetContainerIDFromConglomerationID(ctx context.Context, in *GetContainerIDFromConglomerationIDRequest, opts ...dcerpc.CallOption) (*GetContainerIDFromConglomerationIDResponse, error) {
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
	out := &GetContainerIDFromConglomerationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultContainerControl2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultContainerControl2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultContainerControl2Client) IPID(ctx context.Context, ipid *dcom.IPID) ContainerControl2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultContainerControl2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewContainerControl2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ContainerControl2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ContainerControl2SyntaxV0_0))...)
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
	return &xxx_DefaultContainerControl2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ShutdownContainerOperation structure represents the ShutdownContainer operation
type xxx_ShutdownContainerOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ContainerID *dtyp.GUID     `idl:"name:ContainerIdentifier" json:"container_id"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ShutdownContainerOperation) OpNum() int { return 3 }

func (o *xxx_ShutdownContainerOperation) OpName() string {
	return "/IContainerControl2/v0/ShutdownContainer"
}

func (o *xxx_ShutdownContainerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShutdownContainerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ContainerIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID != nil {
			if err := o.ContainerID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ShutdownContainerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ContainerIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID == nil {
			o.ContainerID = &dtyp.GUID{}
		}
		if err := o.ContainerID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShutdownContainerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShutdownContainerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ShutdownContainerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ShutdownContainerRequest structure represents the ShutdownContainer operation request
type ShutdownContainerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ContainerIdentifier: The container identifier of an instance container.
	ContainerID *dtyp.GUID `idl:"name:ContainerIdentifier" json:"container_id"`
}

func (o *ShutdownContainerRequest) xxx_ToOp(ctx context.Context) *xxx_ShutdownContainerOperation {
	if o == nil {
		return &xxx_ShutdownContainerOperation{}
	}
	return &xxx_ShutdownContainerOperation{
		This:        o.This,
		ContainerID: o.ContainerID,
	}
}

func (o *ShutdownContainerRequest) xxx_FromOp(ctx context.Context, op *xxx_ShutdownContainerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ContainerID = op.ContainerID
}
func (o *ShutdownContainerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ShutdownContainerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ShutdownContainerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ShutdownContainerResponse structure represents the ShutdownContainer operation response
type ShutdownContainerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ShutdownContainer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ShutdownContainerResponse) xxx_ToOp(ctx context.Context) *xxx_ShutdownContainerOperation {
	if o == nil {
		return &xxx_ShutdownContainerOperation{}
	}
	return &xxx_ShutdownContainerOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ShutdownContainerResponse) xxx_FromOp(ctx context.Context, op *xxx_ShutdownContainerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ShutdownContainerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ShutdownContainerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ShutdownContainerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PauseContainerOperation structure represents the PauseContainer operation
type xxx_PauseContainerOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ContainerID *dtyp.GUID     `idl:"name:ContainerIdentifier" json:"container_id"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PauseContainerOperation) OpNum() int { return 4 }

func (o *xxx_PauseContainerOperation) OpName() string { return "/IContainerControl2/v0/PauseContainer" }

func (o *xxx_PauseContainerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseContainerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ContainerIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID != nil {
			if err := o.ContainerID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_PauseContainerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ContainerIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID == nil {
			o.ContainerID = &dtyp.GUID{}
		}
		if err := o.ContainerID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseContainerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseContainerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PauseContainerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PauseContainerRequest structure represents the PauseContainer operation request
type PauseContainerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ContainerIdentifier: The container identifier of an instance container.
	ContainerID *dtyp.GUID `idl:"name:ContainerIdentifier" json:"container_id"`
}

func (o *PauseContainerRequest) xxx_ToOp(ctx context.Context) *xxx_PauseContainerOperation {
	if o == nil {
		return &xxx_PauseContainerOperation{}
	}
	return &xxx_PauseContainerOperation{
		This:        o.This,
		ContainerID: o.ContainerID,
	}
}

func (o *PauseContainerRequest) xxx_FromOp(ctx context.Context, op *xxx_PauseContainerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ContainerID = op.ContainerID
}
func (o *PauseContainerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PauseContainerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PauseContainerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PauseContainerResponse structure represents the PauseContainer operation response
type PauseContainerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PauseContainer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PauseContainerResponse) xxx_ToOp(ctx context.Context) *xxx_PauseContainerOperation {
	if o == nil {
		return &xxx_PauseContainerOperation{}
	}
	return &xxx_PauseContainerOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PauseContainerResponse) xxx_FromOp(ctx context.Context, op *xxx_PauseContainerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PauseContainerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PauseContainerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PauseContainerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResumeContainerOperation structure represents the ResumeContainer operation
type xxx_ResumeContainerOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ContainerID *dtyp.GUID     `idl:"name:ContainerIdentifier" json:"container_id"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ResumeContainerOperation) OpNum() int { return 5 }

func (o *xxx_ResumeContainerOperation) OpName() string {
	return "/IContainerControl2/v0/ResumeContainer"
}

func (o *xxx_ResumeContainerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeContainerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ContainerIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID != nil {
			if err := o.ContainerID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ResumeContainerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ContainerIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID == nil {
			o.ContainerID = &dtyp.GUID{}
		}
		if err := o.ContainerID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeContainerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeContainerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResumeContainerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ResumeContainerRequest structure represents the ResumeContainer operation request
type ResumeContainerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ContainerIdentifier: The container identifier of an instance container.
	ContainerID *dtyp.GUID `idl:"name:ContainerIdentifier" json:"container_id"`
}

func (o *ResumeContainerRequest) xxx_ToOp(ctx context.Context) *xxx_ResumeContainerOperation {
	if o == nil {
		return &xxx_ResumeContainerOperation{}
	}
	return &xxx_ResumeContainerOperation{
		This:        o.This,
		ContainerID: o.ContainerID,
	}
}

func (o *ResumeContainerRequest) xxx_FromOp(ctx context.Context, op *xxx_ResumeContainerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ContainerID = op.ContainerID
}
func (o *ResumeContainerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ResumeContainerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResumeContainerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResumeContainerResponse structure represents the ResumeContainer operation response
type ResumeContainerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ResumeContainer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResumeContainerResponse) xxx_ToOp(ctx context.Context) *xxx_ResumeContainerOperation {
	if o == nil {
		return &xxx_ResumeContainerOperation{}
	}
	return &xxx_ResumeContainerOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ResumeContainerResponse) xxx_FromOp(ctx context.Context, op *xxx_ResumeContainerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ResumeContainerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ResumeContainerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResumeContainerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsContainerPausedOperation structure represents the IsContainerPaused operation
type xxx_IsContainerPausedOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ContainerID *dtyp.GUID     `idl:"name:ContainerIdentifier" json:"container_id"`
	Paused      bool           `idl:"name:bPaused" json:"paused"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsContainerPausedOperation) OpNum() int { return 6 }

func (o *xxx_IsContainerPausedOperation) OpName() string {
	return "/IContainerControl2/v0/IsContainerPaused"
}

func (o *xxx_IsContainerPausedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsContainerPausedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ContainerIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID != nil {
			if err := o.ContainerID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_IsContainerPausedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ContainerIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID == nil {
			o.ContainerID = &dtyp.GUID{}
		}
		if err := o.ContainerID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsContainerPausedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsContainerPausedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// bPaused {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.Paused {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
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

func (o *xxx_IsContainerPausedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// bPaused {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bPaused int32
		if err := w.ReadData(&_bPaused); err != nil {
			return err
		}
		o.Paused = _bPaused != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsContainerPausedRequest structure represents the IsContainerPaused operation request
type IsContainerPausedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ContainerIdentifier: The container identifier of an instance container.
	ContainerID *dtyp.GUID `idl:"name:ContainerIdentifier" json:"container_id"`
}

func (o *IsContainerPausedRequest) xxx_ToOp(ctx context.Context) *xxx_IsContainerPausedOperation {
	if o == nil {
		return &xxx_IsContainerPausedOperation{}
	}
	return &xxx_IsContainerPausedOperation{
		This:        o.This,
		ContainerID: o.ContainerID,
	}
}

func (o *IsContainerPausedRequest) xxx_FromOp(ctx context.Context, op *xxx_IsContainerPausedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ContainerID = op.ContainerID
}
func (o *IsContainerPausedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsContainerPausedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsContainerPausedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsContainerPausedResponse structure represents the IsContainerPaused operation response
type IsContainerPausedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// bPaused: A pointer to a variable that, upon successful completion, MUST be set to
	// indicate whether or not the instance container is paused.
	Paused bool `idl:"name:bPaused" json:"paused"`
	// Return: The IsContainerPaused return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsContainerPausedResponse) xxx_ToOp(ctx context.Context) *xxx_IsContainerPausedOperation {
	if o == nil {
		return &xxx_IsContainerPausedOperation{}
	}
	return &xxx_IsContainerPausedOperation{
		That:   o.That,
		Paused: o.Paused,
		Return: o.Return,
	}
}

func (o *IsContainerPausedResponse) xxx_FromOp(ctx context.Context, op *xxx_IsContainerPausedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Paused = op.Paused
	o.Return = op.Return
}
func (o *IsContainerPausedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsContainerPausedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsContainerPausedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRunningContainersOperation structure represents the GetRunningContainers operation
type xxx_GetRunningContainersOperation struct {
	This             *dcom.ORPCThis            `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat            `idl:"name:That" json:"that"`
	PartitionID      *dtyp.GUID                `idl:"name:PartitionId" json:"partition_id"`
	ConglomerationID *dtyp.GUID                `idl:"name:ConglomerationId" json:"conglomeration_id"`
	ContainersLength uint32                    `idl:"name:pdwNumContainers" json:"containers_length"`
	Containers       []*coma.InstanceContainer `idl:"name:ppContainers;size_is:(, pdwNumContainers)" json:"containers"`
	Return           int32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRunningContainersOperation) OpNum() int { return 7 }

func (o *xxx_GetRunningContainersOperation) OpName() string {
	return "/IContainerControl2/v0/GetRunningContainers"
}

func (o *xxx_GetRunningContainersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRunningContainersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// PartitionId {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PartitionID != nil {
			if err := o.PartitionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ConglomerationId {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
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

func (o *xxx_GetRunningContainersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// PartitionId {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PartitionID == nil {
			o.PartitionID = &dtyp.GUID{}
		}
		if err := o.PartitionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ConglomerationId {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
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

func (o *xxx_GetRunningContainersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Containers != nil && o.ContainersLength == 0 {
		o.ContainersLength = uint32(len(o.Containers))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRunningContainersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwNumContainers {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ContainersLength); err != nil {
			return err
		}
	}
	// ppContainers {out} (1:{pointer=ref}*(2)*(1))(2:{alias=InstanceContainer}[dim:0,size_is=pdwNumContainers](struct))
	{
		if o.Containers != nil || o.ContainersLength > 0 {
			_ptr_ppContainers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ContainersLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Containers {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Containers[i1] != nil {
						if err := o.Containers[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&coma.InstanceContainer{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Containers); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&coma.InstanceContainer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Containers, _ptr_ppContainers); err != nil {
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

func (o *xxx_GetRunningContainersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwNumContainers {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ContainersLength); err != nil {
			return err
		}
	}
	// ppContainers {out} (1:{pointer=ref}*(2)*(1))(2:{alias=InstanceContainer}[dim:0,size_is=pdwNumContainers](struct))
	{
		_ptr_ppContainers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Containers", sizeInfo[0])
			}
			o.Containers = make([]*coma.InstanceContainer, sizeInfo[0])
			for i1 := range o.Containers {
				i1 := i1
				if o.Containers[i1] == nil {
					o.Containers[i1] = &coma.InstanceContainer{}
				}
				if err := o.Containers[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppContainers := func(ptr interface{}) { o.Containers = *ptr.(*[]*coma.InstanceContainer) }
		if err := w.ReadPointer(&o.Containers, _s_ppContainers, _ptr_ppContainers); err != nil {
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

// GetRunningContainersRequest structure represents the GetRunningContainers operation request
type GetRunningContainersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// PartitionId: The partition identifier of a partition.
	PartitionID *dtyp.GUID `idl:"name:PartitionId" json:"partition_id"`
	// ConglomerationId: The conglomeration identifier of a conglomeration, or GUID_NULL
	// for all instance containers.
	ConglomerationID *dtyp.GUID `idl:"name:ConglomerationId" json:"conglomeration_id"`
}

func (o *GetRunningContainersRequest) xxx_ToOp(ctx context.Context) *xxx_GetRunningContainersOperation {
	if o == nil {
		return &xxx_GetRunningContainersOperation{}
	}
	return &xxx_GetRunningContainersOperation{
		This:             o.This,
		PartitionID:      o.PartitionID,
		ConglomerationID: o.ConglomerationID,
	}
}

func (o *GetRunningContainersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRunningContainersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PartitionID = op.PartitionID
	o.ConglomerationID = op.ConglomerationID
}
func (o *GetRunningContainersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetRunningContainersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRunningContainersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRunningContainersResponse structure represents the GetRunningContainers operation response
type GetRunningContainersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwNumContainers: A pointer to a variable that, upon successful completion, MUST
	// be set to the number of elements in ppContainers.
	ContainersLength uint32 `idl:"name:pdwNumContainers" json:"containers_length"`
	// ppContainers: An array of InstanceContainer (section 2.2.9) structures, each of which
	// represents an instance container for the conglomeration specified in ConglomerationId.
	Containers []*coma.InstanceContainer `idl:"name:ppContainers;size_is:(, pdwNumContainers)" json:"containers"`
	// Return: The GetRunningContainers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRunningContainersResponse) xxx_ToOp(ctx context.Context) *xxx_GetRunningContainersOperation {
	if o == nil {
		return &xxx_GetRunningContainersOperation{}
	}
	return &xxx_GetRunningContainersOperation{
		That:             o.That,
		ContainersLength: o.ContainersLength,
		Containers:       o.Containers,
		Return:           o.Return,
	}
}

func (o *GetRunningContainersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRunningContainersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ContainersLength = op.ContainersLength
	o.Containers = op.Containers
	o.Return = op.Return
}
func (o *GetRunningContainersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetRunningContainersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRunningContainersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetContainerIDFromProcessIDOperation structure represents the GetContainerIDFromProcessID operation
type xxx_GetContainerIDFromProcessIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	PID         uint32         `idl:"name:dwPID" json:"pid"`
	ContainerID *oaut.String   `idl:"name:pbstrContainerID" json:"container_id"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetContainerIDFromProcessIDOperation) OpNum() int { return 8 }

func (o *xxx_GetContainerIDFromProcessIDOperation) OpName() string {
	return "/IContainerControl2/v0/GetContainerIDFromProcessID"
}

func (o *xxx_GetContainerIDFromProcessIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainerIDFromProcessIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwPID {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainerIDFromProcessIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwPID {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainerIDFromProcessIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainerIDFromProcessIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrContainerID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ContainerID != nil {
			_ptr_pbstrContainerID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ContainerID != nil {
					if err := o.ContainerID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ContainerID, _ptr_pbstrContainerID); err != nil {
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

func (o *xxx_GetContainerIDFromProcessIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrContainerID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrContainerID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ContainerID == nil {
				o.ContainerID = &oaut.String{}
			}
			if err := o.ContainerID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrContainerID := func(ptr interface{}) { o.ContainerID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ContainerID, _s_pbstrContainerID, _ptr_pbstrContainerID); err != nil {
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

// GetContainerIDFromProcessIDRequest structure represents the GetContainerIDFromProcessID operation request
type GetContainerIDFromProcessIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwPID: The ProcessIdentifier (see section 3.1.1.3.21) of an instance container.
	PID uint32 `idl:"name:dwPID" json:"pid"`
}

func (o *GetContainerIDFromProcessIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetContainerIDFromProcessIDOperation {
	if o == nil {
		return &xxx_GetContainerIDFromProcessIDOperation{}
	}
	return &xxx_GetContainerIDFromProcessIDOperation{
		This: o.This,
		PID:  o.PID,
	}
}

func (o *GetContainerIDFromProcessIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetContainerIDFromProcessIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PID = op.PID
}
func (o *GetContainerIDFromProcessIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetContainerIDFromProcessIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetContainerIDFromProcessIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetContainerIDFromProcessIDResponse structure represents the GetContainerIDFromProcessID operation response
type GetContainerIDFromProcessIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrContainerID: A pointer to a value that, upon successful completion, MUST be
	// set to the Curly Braced GUID String Syntax ([MS-DTYP] section 2.3.4.3) representation
	// of the container identifier of an instance container.
	ContainerID *oaut.String `idl:"name:pbstrContainerID" json:"container_id"`
	// Return: The GetContainerIDFromProcessID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetContainerIDFromProcessIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetContainerIDFromProcessIDOperation {
	if o == nil {
		return &xxx_GetContainerIDFromProcessIDOperation{}
	}
	return &xxx_GetContainerIDFromProcessIDOperation{
		That:        o.That,
		ContainerID: o.ContainerID,
		Return:      o.Return,
	}
}

func (o *GetContainerIDFromProcessIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetContainerIDFromProcessIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ContainerID = op.ContainerID
	o.Return = op.Return
}
func (o *GetContainerIDFromProcessIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetContainerIDFromProcessIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetContainerIDFromProcessIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RecycleContainerOperation structure represents the RecycleContainer operation
type xxx_RecycleContainerOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ContainerID *dtyp.GUID     `idl:"name:ContainerIdentifier" json:"container_id"`
	ReasonCode  int32          `idl:"name:lReasonCode" json:"reason_code"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RecycleContainerOperation) OpNum() int { return 9 }

func (o *xxx_RecycleContainerOperation) OpName() string {
	return "/IContainerControl2/v0/RecycleContainer"
}

func (o *xxx_RecycleContainerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecycleContainerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ContainerIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID != nil {
			if err := o.ContainerID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lReasonCode {in} (1:(int32))
	{
		if err := w.WriteData(o.ReasonCode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecycleContainerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ContainerIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID == nil {
			o.ContainerID = &dtyp.GUID{}
		}
		if err := o.ContainerID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lReasonCode {in} (1:(int32))
	{
		if err := w.ReadData(&o.ReasonCode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecycleContainerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecycleContainerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RecycleContainerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RecycleContainerRequest structure represents the RecycleContainer operation request
type RecycleContainerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ContainerIdentifier: The container identifier of an instance container.
	ContainerID *dtyp.GUID `idl:"name:ContainerIdentifier" json:"container_id"`
	// lReasonCode: A value representing an application-specific informational reason code
	// for why the instance container is being recycled.
	ReasonCode int32 `idl:"name:lReasonCode" json:"reason_code"`
}

func (o *RecycleContainerRequest) xxx_ToOp(ctx context.Context) *xxx_RecycleContainerOperation {
	if o == nil {
		return &xxx_RecycleContainerOperation{}
	}
	return &xxx_RecycleContainerOperation{
		This:        o.This,
		ContainerID: o.ContainerID,
		ReasonCode:  o.ReasonCode,
	}
}

func (o *RecycleContainerRequest) xxx_FromOp(ctx context.Context, op *xxx_RecycleContainerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ContainerID = op.ContainerID
	o.ReasonCode = op.ReasonCode
}
func (o *RecycleContainerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RecycleContainerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RecycleContainerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RecycleContainerResponse structure represents the RecycleContainer operation response
type RecycleContainerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RecycleContainer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RecycleContainerResponse) xxx_ToOp(ctx context.Context) *xxx_RecycleContainerOperation {
	if o == nil {
		return &xxx_RecycleContainerOperation{}
	}
	return &xxx_RecycleContainerOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RecycleContainerResponse) xxx_FromOp(ctx context.Context, op *xxx_RecycleContainerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RecycleContainerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RecycleContainerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RecycleContainerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetContainerIDFromConglomerationIDOperation structure represents the GetContainerIDFromConglomerationID operation
type xxx_GetContainerIDFromConglomerationIDOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConglomerationID *dtyp.GUID     `idl:"name:ConglomerationIdentifier" json:"conglomeration_id"`
	ContainerID      *dtyp.GUID     `idl:"name:ContainerIdentifier" json:"container_id"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetContainerIDFromConglomerationIDOperation) OpNum() int { return 10 }

func (o *xxx_GetContainerIDFromConglomerationIDOperation) OpName() string {
	return "/IContainerControl2/v0/GetContainerIDFromConglomerationID"
}

func (o *xxx_GetContainerIDFromConglomerationIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainerIDFromConglomerationIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ConglomerationIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
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

func (o *xxx_GetContainerIDFromConglomerationIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ConglomerationIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
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

func (o *xxx_GetContainerIDFromConglomerationIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainerIDFromConglomerationIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ContainerIdentifier {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID != nil {
			if err := o.ContainerID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_GetContainerIDFromConglomerationIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ContainerIdentifier {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID == nil {
			o.ContainerID = &dtyp.GUID{}
		}
		if err := o.ContainerID.UnmarshalNDR(ctx, w); err != nil {
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

// GetContainerIDFromConglomerationIDRequest structure represents the GetContainerIDFromConglomerationID operation request
type GetContainerIDFromConglomerationIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ConglomerationIdentifier: The conglomeration identifier of a conglomeration.
	ConglomerationID *dtyp.GUID `idl:"name:ConglomerationIdentifier" json:"conglomeration_id"`
}

func (o *GetContainerIDFromConglomerationIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetContainerIDFromConglomerationIDOperation {
	if o == nil {
		return &xxx_GetContainerIDFromConglomerationIDOperation{}
	}
	return &xxx_GetContainerIDFromConglomerationIDOperation{
		This:             o.This,
		ConglomerationID: o.ConglomerationID,
	}
}

func (o *GetContainerIDFromConglomerationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetContainerIDFromConglomerationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConglomerationID = op.ConglomerationID
}
func (o *GetContainerIDFromConglomerationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetContainerIDFromConglomerationIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetContainerIDFromConglomerationIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetContainerIDFromConglomerationIDResponse structure represents the GetContainerIDFromConglomerationID operation response
type GetContainerIDFromConglomerationIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ContainerIdentifier: A pointer to a variable that, upon successful completion, MUST
	// be set to the container identifier of the single instance container for the conglomeration
	// specified in ConglomerationIdentifier.
	ContainerID *dtyp.GUID `idl:"name:ContainerIdentifier" json:"container_id"`
	// Return: The GetContainerIDFromConglomerationID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetContainerIDFromConglomerationIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetContainerIDFromConglomerationIDOperation {
	if o == nil {
		return &xxx_GetContainerIDFromConglomerationIDOperation{}
	}
	return &xxx_GetContainerIDFromConglomerationIDOperation{
		That:        o.That,
		ContainerID: o.ContainerID,
		Return:      o.Return,
	}
}

func (o *GetContainerIDFromConglomerationIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetContainerIDFromConglomerationIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ContainerID = op.ContainerID
	o.Return = op.Return
}
func (o *GetContainerIDFromConglomerationIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetContainerIDFromConglomerationIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetContainerIDFromConglomerationIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
