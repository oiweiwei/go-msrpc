package iupdateservicemanager2

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
	uamg "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg"
	iupdateservicemanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateservicemanager/v0"
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
	_ = iupdateservicemanager.GoPackage
	_ = oaut.GoPackage
	_ = uamg.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IUpdateServiceManager2 interface identifier 0bb8531d-7e8d-424f-986c-a0b8f60a3e7b
	UpdateServiceManager2IID = &dcom.IID{Data1: 0x0bb8531d, Data2: 0x7e8d, Data3: 0x424f, Data4: []byte{0x98, 0x6c, 0xa0, 0xb8, 0xf6, 0x0a, 0x3e, 0x7b}}
	// Syntax UUID
	UpdateServiceManager2SyntaxUUID = &uuid.UUID{TimeLow: 0xbb8531d, TimeMid: 0x7e8d, TimeHiAndVersion: 0x424f, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x6c, Node: [6]uint8{0xa0, 0xb8, 0xf6, 0xa, 0x3e, 0x7b}}
	// Syntax ID
	UpdateServiceManager2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: UpdateServiceManager2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUpdateServiceManager2 interface.
type UpdateServiceManager2Client interface {

	// IUpdateServiceManager retrieval method.
	UpdateServiceManager() iupdateservicemanager.UpdateServiceManagerClient

	// The IUpdateSearcher::ClientApplicationID (opnum 10) method retrieves the string used
	// to identify the current client application.
	//
	// The IUpdateSession::ClientApplicationID (opnum 9) method sets the identifier of the
	// calling application.
	//
	// The IUpdateHistoryEntry::ClientApplicationID (opnum 16) method retrieves the ID of
	// the application that initiated the operation.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 16) method sets a string that
	// identifies the client application that is using this interface.
	//
	// The IUpdateSession::ClientApplicationID (opnum 8) method retrieves the identifier
	// of the calling application.
	//
	// The IUpdateSearcher::ClientApplicationID (opnum 11) method sets the string used to
	// identify the current client application.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 15) method retrieves a string
	// that identifies the client application that is using this interface.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ClientApplicationID ADM element.
	GetClientApplicationID(context.Context, *GetClientApplicationIDRequest, ...dcerpc.CallOption) (*GetClientApplicationIDResponse, error)

	// The IUpdateSearcher::ClientApplicationID (opnum 10) method retrieves the string used
	// to identify the current client application.
	//
	// The IUpdateSession::ClientApplicationID (opnum 9) method sets the identifier of the
	// calling application.
	//
	// The IUpdateHistoryEntry::ClientApplicationID (opnum 16) method retrieves the ID of
	// the application that initiated the operation.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 16) method sets a string that
	// identifies the client application that is using this interface.
	//
	// The IUpdateSession::ClientApplicationID (opnum 8) method retrieves the identifier
	// of the calling application.
	//
	// The IUpdateSearcher::ClientApplicationID (opnum 11) method sets the string used to
	// identify the current client application.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 15) method retrieves a string
	// that identifies the client application that is using this interface.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ClientApplicationID ADM element.
	SetClientApplicationID(context.Context, *SetClientApplicationIDRequest, ...dcerpc.CallOption) (*SetClientApplicationIDResponse, error)

	// The IUpdateServiceManager2::QueryServiceRegistration (opnum 17) method retrieves
	// the update service registration record for a service.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the service registration record for the identified service
	// from the ServiceRegistrations ADM element.
	QueryServiceRegistration(context.Context, *QueryServiceRegistrationRequest, ...dcerpc.CallOption) (*QueryServiceRegistrationResponse, error)

	// The IUpdateServiceManager2::AddService2 (opnum 18) method registers an update service
	// with the update agent.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD trigger the update agent to add this update service through an
	// implementation-dependent<48> interface. This method SHOULD add the given service
	// to the Services ADM element maintained by the IUpdateServiceManager (section 3.44)
	// interface and return an IUpdateServiceRegistration instance that represents the service
	// registration.
	AddService2(context.Context, *AddService2Request, ...dcerpc.CallOption) (*AddService2Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) UpdateServiceManager2Client
}

type xxx_DefaultUpdateServiceManager2Client struct {
	iupdateservicemanager.UpdateServiceManagerClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUpdateServiceManager2Client) UpdateServiceManager() iupdateservicemanager.UpdateServiceManagerClient {
	return o.UpdateServiceManagerClient
}

func (o *xxx_DefaultUpdateServiceManager2Client) GetClientApplicationID(ctx context.Context, in *GetClientApplicationIDRequest, opts ...dcerpc.CallOption) (*GetClientApplicationIDResponse, error) {
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
	out := &GetClientApplicationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateServiceManager2Client) SetClientApplicationID(ctx context.Context, in *SetClientApplicationIDRequest, opts ...dcerpc.CallOption) (*SetClientApplicationIDResponse, error) {
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
	out := &SetClientApplicationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateServiceManager2Client) QueryServiceRegistration(ctx context.Context, in *QueryServiceRegistrationRequest, opts ...dcerpc.CallOption) (*QueryServiceRegistrationResponse, error) {
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
	out := &QueryServiceRegistrationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateServiceManager2Client) AddService2(ctx context.Context, in *AddService2Request, opts ...dcerpc.CallOption) (*AddService2Response, error) {
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
	out := &AddService2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateServiceManager2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUpdateServiceManager2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUpdateServiceManager2Client) IPID(ctx context.Context, ipid *dcom.IPID) UpdateServiceManager2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUpdateServiceManager2Client{
		UpdateServiceManagerClient: o.UpdateServiceManagerClient.IPID(ctx, ipid),
		cc:                         o.cc,
		ipid:                       ipid,
	}
}

func NewUpdateServiceManager2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (UpdateServiceManager2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(UpdateServiceManager2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iupdateservicemanager.NewUpdateServiceManagerClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultUpdateServiceManager2Client{
		UpdateServiceManagerClient: base,
		cc:                         cc,
		ipid:                       ipid,
	}, nil
}

// xxx_GetClientApplicationIDOperation structure represents the ClientApplicationID operation
type xxx_GetClientApplicationIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue *oaut.String   `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClientApplicationIDOperation) OpNum() int { return 14 }

func (o *xxx_GetClientApplicationIDOperation) OpName() string {
	return "/IUpdateServiceManager2/v0/ClientApplicationID"
}

func (o *xxx_GetClientApplicationIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientApplicationIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetClientApplicationIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetClientApplicationIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientApplicationIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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

func (o *xxx_GetClientApplicationIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &oaut.String{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
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

// GetClientApplicationIDRequest structure represents the ClientApplicationID operation request
type GetClientApplicationIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClientApplicationIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetClientApplicationIDOperation) *xxx_GetClientApplicationIDOperation {
	if op == nil {
		op = &xxx_GetClientApplicationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetClientApplicationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClientApplicationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClientApplicationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetClientApplicationIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientApplicationIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClientApplicationIDResponse structure represents the ClientApplicationID operation response
type GetClientApplicationIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: Returns a string identifying the client application using the interface.
	//
	// retval: A string identifying the client application that initiated the operation.
	//
	// retval: The identifier of the calling application previously set using the IUpdateSession::ClientApplicationID
	// (opnum 9) (section 3.16.4.2) method. If no identifier was previously set, this MUST
	// be NULL or the empty string.
	//
	// retval: A string identifying the client application that is using this interface.
	ReturnValue *oaut.String `idl:"name:retval" json:"return_value"`
	// Return: The ClientApplicationID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClientApplicationIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetClientApplicationIDOperation) *xxx_GetClientApplicationIDOperation {
	if op == nil {
		op = &xxx_GetClientApplicationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetClientApplicationIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClientApplicationIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetClientApplicationIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetClientApplicationIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientApplicationIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetClientApplicationIDOperation structure represents the ClientApplicationID operation
type xxx_SetClientApplicationIDOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value  *oaut.String   `idl:"name:value" json:"value"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetClientApplicationIDOperation) OpNum() int { return 15 }

func (o *xxx_SetClientApplicationIDOperation) OpName() string {
	return "/IUpdateServiceManager2/v0/ClientApplicationID"
}

func (o *xxx_SetClientApplicationIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientApplicationIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// value {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Value != nil {
			_ptr_value := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Value, _ptr_value); err != nil {
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

func (o *xxx_SetClientApplicationIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// value {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_value := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.String{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_value := func(ptr interface{}) { o.Value = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Value, _s_value, _ptr_value); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientApplicationIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientApplicationIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetClientApplicationIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetClientApplicationIDRequest structure represents the ClientApplicationID operation request
type SetClientApplicationIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// value: The identifier of the calling application.
	//
	// value: A string that identifies the client application that is using this interface.
	//
	// value: A string used to identify the client application using the interface.
	Value *oaut.String `idl:"name:value" json:"value"`
}

func (o *SetClientApplicationIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetClientApplicationIDOperation) *xxx_SetClientApplicationIDOperation {
	if op == nil {
		op = &xxx_SetClientApplicationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Value = o.Value
	return op
}

func (o *SetClientApplicationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetClientApplicationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Value = op.Value
}
func (o *SetClientApplicationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetClientApplicationIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientApplicationIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetClientApplicationIDResponse structure represents the ClientApplicationID operation response
type SetClientApplicationIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClientApplicationID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetClientApplicationIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetClientApplicationIDOperation) *xxx_SetClientApplicationIDOperation {
	if op == nil {
		op = &xxx_SetClientApplicationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetClientApplicationIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetClientApplicationIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetClientApplicationIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetClientApplicationIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientApplicationIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryServiceRegistrationOperation structure represents the QueryServiceRegistration operation
type xxx_QueryServiceRegistrationOperation struct {
	This        *dcom.ORPCThis                  `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat                  `idl:"name:That" json:"that"`
	ServiceID   *oaut.String                    `idl:"name:serviceID" json:"service_id"`
	ReturnValue *uamg.UpdateServiceRegistration `idl:"name:retval" json:"return_value"`
	Return      int32                           `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryServiceRegistrationOperation) OpNum() int { return 16 }

func (o *xxx_QueryServiceRegistrationOperation) OpName() string {
	return "/IUpdateServiceManager2/v0/QueryServiceRegistration"
}

func (o *xxx_QueryServiceRegistrationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryServiceRegistrationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// serviceID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ServiceID != nil {
			_ptr_serviceID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServiceID != nil {
					if err := o.ServiceID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServiceID, _ptr_serviceID); err != nil {
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

func (o *xxx_QueryServiceRegistrationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// serviceID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_serviceID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServiceID == nil {
				o.ServiceID = &oaut.String{}
			}
			if err := o.ServiceID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_serviceID := func(ptr interface{}) { o.ServiceID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ServiceID, _s_serviceID, _ptr_serviceID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryServiceRegistrationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryServiceRegistrationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateServiceRegistration}(interface))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&uamg.UpdateServiceRegistration{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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

func (o *xxx_QueryServiceRegistrationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateServiceRegistration}(interface))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &uamg.UpdateServiceRegistration{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**uamg.UpdateServiceRegistration) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
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

// QueryServiceRegistrationRequest structure represents the QueryServiceRegistration operation request
type QueryServiceRegistrationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// serviceID: A string identifying an update service.
	ServiceID *oaut.String `idl:"name:serviceID" json:"service_id"`
}

func (o *QueryServiceRegistrationRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryServiceRegistrationOperation) *xxx_QueryServiceRegistrationOperation {
	if op == nil {
		op = &xxx_QueryServiceRegistrationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ServiceID = o.ServiceID
	return op
}

func (o *QueryServiceRegistrationRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryServiceRegistrationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ServiceID = op.ServiceID
}
func (o *QueryServiceRegistrationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryServiceRegistrationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryServiceRegistrationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryServiceRegistrationResponse structure represents the QueryServiceRegistration operation response
type QueryServiceRegistrationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: An IUpdateServiceRegistration (section 3.46.4) instance for the given service.
	ReturnValue *uamg.UpdateServiceRegistration `idl:"name:retval" json:"return_value"`
	// Return: The QueryServiceRegistration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryServiceRegistrationResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryServiceRegistrationOperation) *xxx_QueryServiceRegistrationOperation {
	if op == nil {
		op = &xxx_QueryServiceRegistrationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *QueryServiceRegistrationResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryServiceRegistrationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *QueryServiceRegistrationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryServiceRegistrationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryServiceRegistrationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddService2Operation structure represents the AddService2 operation
type xxx_AddService2Operation struct {
	This                 *dcom.ORPCThis                  `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat                  `idl:"name:That" json:"that"`
	ServiceID            *oaut.String                    `idl:"name:serviceID" json:"service_id"`
	Flags                int32                           `idl:"name:flags" json:"flags"`
	AuthorizationCabPath *oaut.String                    `idl:"name:authorizationCabPath" json:"authorization_cab_path"`
	ReturnValue          *uamg.UpdateServiceRegistration `idl:"name:retval" json:"return_value"`
	Return               int32                           `idl:"name:Return" json:"return"`
}

func (o *xxx_AddService2Operation) OpNum() int { return 17 }

func (o *xxx_AddService2Operation) OpName() string { return "/IUpdateServiceManager2/v0/AddService2" }

func (o *xxx_AddService2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddService2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// serviceID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ServiceID != nil {
			_ptr_serviceID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServiceID != nil {
					if err := o.ServiceID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServiceID, _ptr_serviceID); err != nil {
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
	// flags {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// authorizationCabPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.AuthorizationCabPath != nil {
			_ptr_authorizationCabPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AuthorizationCabPath != nil {
					if err := o.AuthorizationCabPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AuthorizationCabPath, _ptr_authorizationCabPath); err != nil {
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

func (o *xxx_AddService2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// serviceID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_serviceID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServiceID == nil {
				o.ServiceID = &oaut.String{}
			}
			if err := o.ServiceID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_serviceID := func(ptr interface{}) { o.ServiceID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ServiceID, _s_serviceID, _ptr_serviceID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// authorizationCabPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_authorizationCabPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AuthorizationCabPath == nil {
				o.AuthorizationCabPath = &oaut.String{}
			}
			if err := o.AuthorizationCabPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_authorizationCabPath := func(ptr interface{}) { o.AuthorizationCabPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.AuthorizationCabPath, _s_authorizationCabPath, _ptr_authorizationCabPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddService2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddService2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateServiceRegistration}(interface))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&uamg.UpdateServiceRegistration{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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

func (o *xxx_AddService2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateServiceRegistration}(interface))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &uamg.UpdateServiceRegistration{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**uamg.UpdateServiceRegistration) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
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

// AddService2Request structure represents the AddService2 operation request
type AddService2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// serviceID: A string identifying the service.
	ServiceID *oaut.String `idl:"name:serviceID" json:"service_id"`
	// flags: A bitmask created with values from the AddServiceFlag (section 2.2.15) enumeration
	// specifying options.
	Flags int32 `idl:"name:flags" json:"flags"`
	// authorizationCabPath: The filesystem path of the authorization cabinet file containing
	// information required for a service registration. If this is NULL or the empty string,
	// the server SHOULD search for the authorization cabinet file online when a network
	// connection is available.
	AuthorizationCabPath *oaut.String `idl:"name:authorizationCabPath" json:"authorization_cab_path"`
}

func (o *AddService2Request) xxx_ToOp(ctx context.Context, op *xxx_AddService2Operation) *xxx_AddService2Operation {
	if op == nil {
		op = &xxx_AddService2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ServiceID = o.ServiceID
	op.Flags = o.Flags
	op.AuthorizationCabPath = o.AuthorizationCabPath
	return op
}

func (o *AddService2Request) xxx_FromOp(ctx context.Context, op *xxx_AddService2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ServiceID = op.ServiceID
	o.Flags = op.Flags
	o.AuthorizationCabPath = op.AuthorizationCabPath
}
func (o *AddService2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddService2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddService2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddService2Response structure represents the AddService2 operation response
type AddService2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: An IUpdateServiceRegistration instance.
	ReturnValue *uamg.UpdateServiceRegistration `idl:"name:retval" json:"return_value"`
	// Return: The AddService2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddService2Response) xxx_ToOp(ctx context.Context, op *xxx_AddService2Operation) *xxx_AddService2Operation {
	if op == nil {
		op = &xxx_AddService2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *AddService2Response) xxx_FromOp(ctx context.Context, op *xxx_AddService2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *AddService2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddService2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddService2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
