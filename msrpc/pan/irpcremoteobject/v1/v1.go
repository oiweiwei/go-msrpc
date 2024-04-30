package irpcremoteobject

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
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "pan"
)

var (
	// Syntax UUID
	IrpcRemoteObjectSyntaxUUID = &uuid.UUID{TimeLow: 0xae33069b, TimeMid: 0xa2a8, TimeHiAndVersion: 0x46ee, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x35, Node: [6]uint8{0xdd, 0xfd, 0x33, 0x9b, 0xe2, 0x81}}
	// Syntax ID
	IrpcRemoteObjectSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: IrpcRemoteObjectSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// IRPCRemoteObject interface.
type IrpcRemoteObjectClient interface {

	// The IRPCRemoteObject_Create method creates a remote object on a server and returns
	// it to the client.
	//
	// Return Values: This method MUST return zero to indicate success, or an HRESULT error
	// value ([MS-ERREF] section 2.1.1) to indicate failure. The client MUST consider all
	// error return values fatal and report them to the higher-level caller.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	Create(context.Context, *CreateRequest, ...dcerpc.CallOption) (*CreateResponse, error)

	// The IRPCRemoteObject_Delete method destroys the specified remote object.
	//
	// Return Values: This method has no return values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	Delete(context.Context, *DeleteRequest, ...dcerpc.CallOption) (*DeleteResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error
}

// RemoteObject structure represents PRPCREMOTEOBJECT RPC structure.
type RemoteObject dcetypes.ContextHandle

func (o *RemoteObject) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *RemoteObject) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RemoteObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

type xxx_DefaultIrpcRemoteObjectClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultIrpcRemoteObjectClient) Create(ctx context.Context, in *CreateRequest, opts ...dcerpc.CallOption) (*CreateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIrpcRemoteObjectClient) Delete(ctx context.Context, in *DeleteRequest, opts ...dcerpc.CallOption) (*DeleteResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultIrpcRemoteObjectClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}
func NewIrpcRemoteObjectClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IrpcRemoteObjectClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IrpcRemoteObjectSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultIrpcRemoteObjectClient{cc: cc}, nil
}

// xxx_CreateOperation structure represents the IRPCRemoteObject_Create operation
type xxx_CreateOperation struct {
	RemoteObject *RemoteObject `idl:"name:ppRemoteObj" json:"remote_object"`
	Return       int32         `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateOperation) OpNum() int { return 0 }

func (o *xxx_CreateOperation) OpName() string { return "/IRPCRemoteObject/v1/IRPCRemoteObject_Create" }

func (o *xxx_CreateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_CreateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_CreateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppRemoteObj {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RemoteObject != nil {
			if err := o.RemoteObject.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteObject{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CreateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppRemoteObj {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RemoteObject == nil {
			o.RemoteObject = &RemoteObject{}
		}
		if err := o.RemoteObject.UnmarshalNDR(ctx, w); err != nil {
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

// CreateRequest structure represents the IRPCRemoteObject_Create operation request
type CreateRequest struct {
}

func (o *CreateRequest) xxx_ToOp(ctx context.Context) *xxx_CreateOperation {
	if o == nil {
		return &xxx_CreateOperation{}
	}
	return &xxx_CreateOperation{}
}

func (o *CreateRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateOperation) {
	if o == nil {
		return
	}
}
func (o *CreateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateResponse structure represents the IRPCRemoteObject_Create operation response
type CreateResponse struct {
	// ppRemoteObj: MUST be a remote object context handle returned by the server. It MUST
	// be a non-NULL value.
	RemoteObject *RemoteObject `idl:"name:ppRemoteObj" json:"remote_object"`
	// Return: The IRPCRemoteObject_Create return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateResponse) xxx_ToOp(ctx context.Context) *xxx_CreateOperation {
	if o == nil {
		return &xxx_CreateOperation{}
	}
	return &xxx_CreateOperation{
		RemoteObject: o.RemoteObject,
		Return:       o.Return,
	}
}

func (o *CreateResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateOperation) {
	if o == nil {
		return
	}
	o.RemoteObject = op.RemoteObject
	o.Return = op.Return
}
func (o *CreateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteOperation structure represents the IRPCRemoteObject_Delete operation
type xxx_DeleteOperation struct {
	RemoteObject *RemoteObject `idl:"name:ppRemoteObj" json:"remote_object"`
}

func (o *xxx_DeleteOperation) OpNum() int { return 1 }

func (o *xxx_DeleteOperation) OpName() string { return "/IRPCRemoteObject/v1/IRPCRemoteObject_Delete" }

func (o *xxx_DeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
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
	// ppRemoteObj {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RemoteObject != nil {
			if err := o.RemoteObject.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteObject{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ppRemoteObj {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RemoteObject == nil {
			o.RemoteObject = &RemoteObject{}
		}
		if err := o.RemoteObject.UnmarshalNDR(ctx, w); err != nil {
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
	// ppRemoteObj {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RemoteObject != nil {
			if err := o.RemoteObject.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteObject{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppRemoteObj {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RemoteObject == nil {
			o.RemoteObject = &RemoteObject{}
		}
		if err := o.RemoteObject.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// DeleteRequest structure represents the IRPCRemoteObject_Delete operation request
type DeleteRequest struct {
	// ppRemoteObj: MUST be the remote object to delete. The handle MUST have been returned
	// by the server in the ppRemoteObj output parameter of a prior call to IRPCRemoteObject_Create
	// and MUST NOT have been previously deleted. If this handle were previously registered
	// by a successful call to IRPCAsyncNotify_RegisterClient, then it MUST have been subsequently
	// unregistered by a call to IRPCAsyncNotify_UnregisterClient. It MUST NOT be NULL.
	RemoteObject *RemoteObject `idl:"name:ppRemoteObj" json:"remote_object"`
}

func (o *DeleteRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteOperation {
	if o == nil {
		return &xxx_DeleteOperation{}
	}
	return &xxx_DeleteOperation{
		RemoteObject: o.RemoteObject,
	}
}

func (o *DeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteOperation) {
	if o == nil {
		return
	}
	o.RemoteObject = op.RemoteObject
}
func (o *DeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteResponse structure represents the IRPCRemoteObject_Delete operation response
type DeleteResponse struct {
	// ppRemoteObj: MUST be the remote object to delete. The handle MUST have been returned
	// by the server in the ppRemoteObj output parameter of a prior call to IRPCRemoteObject_Create
	// and MUST NOT have been previously deleted. If this handle were previously registered
	// by a successful call to IRPCAsyncNotify_RegisterClient, then it MUST have been subsequently
	// unregistered by a call to IRPCAsyncNotify_UnregisterClient. It MUST NOT be NULL.
	RemoteObject *RemoteObject `idl:"name:ppRemoteObj" json:"remote_object"`
}

func (o *DeleteResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteOperation {
	if o == nil {
		return &xxx_DeleteOperation{}
	}
	return &xxx_DeleteOperation{
		RemoteObject: o.RemoteObject,
	}
}

func (o *DeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteOperation) {
	if o == nil {
		return
	}
	o.RemoteObject = op.RemoteObject
}
func (o *DeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
