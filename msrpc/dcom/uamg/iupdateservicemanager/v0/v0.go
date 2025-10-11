package iupdateservicemanager

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
	uamg "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg"
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
	_ = uamg.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IUpdateServiceManager interface identifier 23857e3c-02ba-44a3-9423-b1c900805f37
	UpdateServiceManagerIID = &dcom.IID{Data1: 0x23857e3c, Data2: 0x02ba, Data3: 0x44a3, Data4: []byte{0x94, 0x23, 0xb1, 0xc9, 0x00, 0x80, 0x5f, 0x37}}
	// Syntax UUID
	UpdateServiceManagerSyntaxUUID = &uuid.UUID{TimeLow: 0x23857e3c, TimeMid: 0x2ba, TimeHiAndVersion: 0x44a3, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x23, Node: [6]uint8{0xb1, 0xc9, 0x0, 0x80, 0x5f, 0x37}}
	// Syntax ID
	UpdateServiceManagerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: UpdateServiceManagerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUpdateServiceManager interface.
type UpdateServiceManagerClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The IUpdateServiceManager::Services (opnum 8) method retrieves a collection of update
	// services registered with the update agent.
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
	// This method SHOULD return the value of the Services ADM element.
	GetServices(context.Context, *GetServicesRequest, ...dcerpc.CallOption) (*GetServicesResponse, error)

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// The IUpdateServiceManager::RegisterServiceWithAU (opnum 10) method registers an update
	// service with the automatic update agent.
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
	// This method SHOULD register the given service with the automatic update agent through
	// an implementation-defined interface. This method SHOULD set the ServiceRegisteredWithAU
	// ADM element to the given update service identifier.
	RegisterServiceWithAU(context.Context, *RegisterServiceWithAURequest, ...dcerpc.CallOption) (*RegisterServiceWithAUResponse, error)

	// The IUpdateServiceManager::RemoveService (opnum 11) method removes an update service
	// from the update agent.
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
	// This method SHOULD trigger the update agent to remove the given service through an
	// implementation-defined interface. The method SHOULD remove the given service from
	// the Services ADM element.
	RemoveService(context.Context, *RemoveServiceRequest, ...dcerpc.CallOption) (*RemoveServiceResponse, error)

	// Opnum12NotUsedOnWire operation.
	// Opnum12NotUsedOnWire

	// The IUpdateServiceManager::AddScanPackageService (opnum 13) method registers an update
	// service based on a scan package.
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
	// This method SHOULD trigger the update agent to add a service based on this scan package
	// through an implementation-dependent<47> interface. This method SHOULD add the given
	// service to the Services ADM element and return an IUpdateService instance representing
	// the added service.
	AddScanPackageService(context.Context, *AddScanPackageServiceRequest, ...dcerpc.CallOption) (*AddScanPackageServiceResponse, error)

	// The IUpdateServiceManager::SetOption (opnum 14) method sets options on this interface.
	// The "AllowedServiceID" option restricts the IUpdateServiceManager::RegisterServiceWithAU
	// (opnum 10) (section 3.44.4.2) method to work only with the given service ID. The
	// "AllowWarningUI" option controls whether a warning UI is displayed when changing
	// the service registered with the automatic update agent.
	//
	// Return Values: The method MUST return information in an HRESULT  data structure.
	// The severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// If optionName is "AllowedServiceID", this method SHOULD set the value of the AllowedServiceID
	// ADM element to the string stored in optionValue. If optionName is "AllowWarningUI",
	// this method SHOULD set the value of the AllowWarningUI ADM element to the VARIANT_BOOL
	// stored in optionValue.
	SetOption(context.Context, *SetOptionRequest, ...dcerpc.CallOption) (*SetOptionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) UpdateServiceManagerClient
}

type xxx_DefaultUpdateServiceManagerClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUpdateServiceManagerClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultUpdateServiceManagerClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...dcerpc.CallOption) (*GetServicesResponse, error) {
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
	out := &GetServicesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateServiceManagerClient) RegisterServiceWithAU(ctx context.Context, in *RegisterServiceWithAURequest, opts ...dcerpc.CallOption) (*RegisterServiceWithAUResponse, error) {
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
	out := &RegisterServiceWithAUResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateServiceManagerClient) RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...dcerpc.CallOption) (*RemoveServiceResponse, error) {
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
	out := &RemoveServiceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateServiceManagerClient) AddScanPackageService(ctx context.Context, in *AddScanPackageServiceRequest, opts ...dcerpc.CallOption) (*AddScanPackageServiceResponse, error) {
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
	out := &AddScanPackageServiceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateServiceManagerClient) SetOption(ctx context.Context, in *SetOptionRequest, opts ...dcerpc.CallOption) (*SetOptionResponse, error) {
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
	out := &SetOptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateServiceManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUpdateServiceManagerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUpdateServiceManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) UpdateServiceManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUpdateServiceManagerClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewUpdateServiceManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (UpdateServiceManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(UpdateServiceManagerSyntaxV0_0))...)
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
	return &xxx_DefaultUpdateServiceManagerClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetServicesOperation structure represents the Services operation
type xxx_GetServicesOperation struct {
	This        *dcom.ORPCThis                `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat                `idl:"name:That" json:"that"`
	ReturnValue *uamg.UpdateServiceCollection `idl:"name:retval" json:"return_value"`
	Return      int32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServicesOperation) OpNum() int { return 7 }

func (o *xxx_GetServicesOperation) OpName() string { return "/IUpdateServiceManager/v0/Services" }

func (o *xxx_GetServicesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServicesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetServicesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetServicesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServicesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateServiceCollection}(interface))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&uamg.UpdateServiceCollection{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetServicesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateServiceCollection}(interface))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &uamg.UpdateServiceCollection{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**uamg.UpdateServiceCollection) }
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

// GetServicesRequest structure represents the Services operation request
type GetServicesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetServicesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetServicesOperation) *xxx_GetServicesOperation {
	if op == nil {
		op = &xxx_GetServicesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetServicesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServicesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetServicesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetServicesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServicesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServicesResponse structure represents the Services operation response
type GetServicesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: A collection of update services registered with the update agent.
	ReturnValue *uamg.UpdateServiceCollection `idl:"name:retval" json:"return_value"`
	// Return: The Services return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetServicesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetServicesOperation) *xxx_GetServicesOperation {
	if op == nil {
		op = &xxx_GetServicesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetServicesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServicesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetServicesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetServicesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServicesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterServiceWithAUOperation structure represents the RegisterServiceWithAU operation
type xxx_RegisterServiceWithAUOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServiceID *oaut.String   `idl:"name:serviceID" json:"service_id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterServiceWithAUOperation) OpNum() int { return 9 }

func (o *xxx_RegisterServiceWithAUOperation) OpName() string {
	return "/IUpdateServiceManager/v0/RegisterServiceWithAU"
}

func (o *xxx_RegisterServiceWithAUOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterServiceWithAUOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RegisterServiceWithAUOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_RegisterServiceWithAUOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterServiceWithAUOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RegisterServiceWithAUOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RegisterServiceWithAURequest structure represents the RegisterServiceWithAU operation request
type RegisterServiceWithAURequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// serviceID: A string identifying the update service to register with the automatic
	// update agent. This service MUST already be registered with the update agent.
	ServiceID *oaut.String `idl:"name:serviceID" json:"service_id"`
}

func (o *RegisterServiceWithAURequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterServiceWithAUOperation) *xxx_RegisterServiceWithAUOperation {
	if op == nil {
		op = &xxx_RegisterServiceWithAUOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ServiceID = o.ServiceID
	return op
}

func (o *RegisterServiceWithAURequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterServiceWithAUOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ServiceID = op.ServiceID
}
func (o *RegisterServiceWithAURequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterServiceWithAURequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterServiceWithAUOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterServiceWithAUResponse structure represents the RegisterServiceWithAU operation response
type RegisterServiceWithAUResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RegisterServiceWithAU return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterServiceWithAUResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterServiceWithAUOperation) *xxx_RegisterServiceWithAUOperation {
	if op == nil {
		op = &xxx_RegisterServiceWithAUOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RegisterServiceWithAUResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterServiceWithAUOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RegisterServiceWithAUResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterServiceWithAUResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterServiceWithAUOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveServiceOperation structure represents the RemoveService operation
type xxx_RemoveServiceOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServiceID *oaut.String   `idl:"name:serviceID" json:"service_id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveServiceOperation) OpNum() int { return 10 }

func (o *xxx_RemoveServiceOperation) OpName() string {
	return "/IUpdateServiceManager/v0/RemoveService"
}

func (o *xxx_RemoveServiceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveServiceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoveServiceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_RemoveServiceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveServiceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoveServiceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RemoveServiceRequest structure represents the RemoveService operation request
type RemoveServiceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// serviceID: A string identifying the registered service to remove.
	ServiceID *oaut.String `idl:"name:serviceID" json:"service_id"`
}

func (o *RemoveServiceRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoveServiceOperation) *xxx_RemoveServiceOperation {
	if op == nil {
		op = &xxx_RemoveServiceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ServiceID = o.ServiceID
	return op
}

func (o *RemoveServiceRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveServiceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ServiceID = op.ServiceID
}
func (o *RemoveServiceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoveServiceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveServiceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveServiceResponse structure represents the RemoveService operation response
type RemoveServiceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RemoveService return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveServiceResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoveServiceOperation) *xxx_RemoveServiceOperation {
	if op == nil {
		op = &xxx_RemoveServiceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RemoveServiceResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveServiceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RemoveServiceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoveServiceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveServiceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddScanPackageServiceOperation structure represents the AddScanPackageService operation
type xxx_AddScanPackageServiceOperation struct {
	This             *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat      `idl:"name:That" json:"that"`
	ServiceName      *oaut.String        `idl:"name:serviceName" json:"service_name"`
	ScanFileLocation *oaut.String        `idl:"name:scanFileLocation" json:"scan_file_location"`
	Flags            int32               `idl:"name:flags" json:"flags"`
	Service          *uamg.UpdateService `idl:"name:ppService" json:"service"`
	Return           int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_AddScanPackageServiceOperation) OpNum() int { return 12 }

func (o *xxx_AddScanPackageServiceOperation) OpName() string {
	return "/IUpdateServiceManager/v0/AddScanPackageService"
}

func (o *xxx_AddScanPackageServiceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddScanPackageServiceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// serviceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ServiceName != nil {
			_ptr_serviceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServiceName != nil {
					if err := o.ServiceName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServiceName, _ptr_serviceName); err != nil {
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
	// scanFileLocation {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ScanFileLocation != nil {
			_ptr_scanFileLocation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ScanFileLocation != nil {
					if err := o.ScanFileLocation.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ScanFileLocation, _ptr_scanFileLocation); err != nil {
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
	// flags {in, default_value={0}} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddScanPackageServiceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// serviceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_serviceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServiceName == nil {
				o.ServiceName = &oaut.String{}
			}
			if err := o.ServiceName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_serviceName := func(ptr interface{}) { o.ServiceName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ServiceName, _s_serviceName, _ptr_serviceName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// scanFileLocation {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_scanFileLocation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ScanFileLocation == nil {
				o.ScanFileLocation = &oaut.String{}
			}
			if err := o.ScanFileLocation.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_scanFileLocation := func(ptr interface{}) { o.ScanFileLocation = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ScanFileLocation, _s_scanFileLocation, _ptr_scanFileLocation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in, default_value={0}} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddScanPackageServiceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddScanPackageServiceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppService {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateService}(interface))
	{
		if o.Service != nil {
			_ptr_ppService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Service != nil {
					if err := o.Service.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&uamg.UpdateService{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Service, _ptr_ppService); err != nil {
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

func (o *xxx_AddScanPackageServiceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppService {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateService}(interface))
	{
		_ptr_ppService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Service == nil {
				o.Service = &uamg.UpdateService{}
			}
			if err := o.Service.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppService := func(ptr interface{}) { o.Service = *ptr.(**uamg.UpdateService) }
		if err := w.ReadPointer(&o.Service, _s_ppService, _ptr_ppService); err != nil {
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

// AddScanPackageServiceRequest structure represents the AddScanPackageService operation request
type AddScanPackageServiceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// serviceName: The name of the service.
	ServiceName *oaut.String `idl:"name:serviceName" json:"service_name"`
	// scanFileLocation: The filesystem path of the scan package file.
	ScanFileLocation *oaut.String `idl:"name:scanFileLocation" json:"scan_file_location"`
	// flags: A bitmask produced from values from the UpdateServiceOption (section 2.2.14)
	// enumeration value.
	Flags int32 `idl:"name:flags" json:"flags"`
}

func (o *AddScanPackageServiceRequest) xxx_ToOp(ctx context.Context, op *xxx_AddScanPackageServiceOperation) *xxx_AddScanPackageServiceOperation {
	if op == nil {
		op = &xxx_AddScanPackageServiceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ServiceName = o.ServiceName
	op.ScanFileLocation = o.ScanFileLocation
	op.Flags = o.Flags
	return op
}

func (o *AddScanPackageServiceRequest) xxx_FromOp(ctx context.Context, op *xxx_AddScanPackageServiceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ServiceName = op.ServiceName
	o.ScanFileLocation = op.ScanFileLocation
	o.Flags = op.Flags
}
func (o *AddScanPackageServiceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddScanPackageServiceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddScanPackageServiceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddScanPackageServiceResponse structure represents the AddScanPackageService operation response
type AddScanPackageServiceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppService: An IUpdateService instance representing a service based on the scan package.
	Service *uamg.UpdateService `idl:"name:ppService" json:"service"`
	// Return: The AddScanPackageService return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddScanPackageServiceResponse) xxx_ToOp(ctx context.Context, op *xxx_AddScanPackageServiceOperation) *xxx_AddScanPackageServiceOperation {
	if op == nil {
		op = &xxx_AddScanPackageServiceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Service = o.Service
	op.Return = o.Return
	return op
}

func (o *AddScanPackageServiceResponse) xxx_FromOp(ctx context.Context, op *xxx_AddScanPackageServiceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Service = op.Service
	o.Return = op.Return
}
func (o *AddScanPackageServiceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddScanPackageServiceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddScanPackageServiceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOptionOperation structure represents the SetOption operation
type xxx_SetOptionOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	OptionName  *oaut.String   `idl:"name:optionName" json:"option_name"`
	OptionValue *oaut.Variant  `idl:"name:optionValue" json:"option_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOptionOperation) OpNum() int { return 13 }

func (o *xxx_SetOptionOperation) OpName() string { return "/IUpdateServiceManager/v0/SetOption" }

func (o *xxx_SetOptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// optionName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.OptionName != nil {
			_ptr_optionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OptionName != nil {
					if err := o.OptionName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OptionName, _ptr_optionName); err != nil {
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
	// optionValue {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.OptionValue != nil {
			if err := o.OptionValue.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// optionName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_optionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OptionName == nil {
				o.OptionName = &oaut.String{}
			}
			if err := o.OptionName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_optionName := func(ptr interface{}) { o.OptionName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.OptionName, _s_optionName, _ptr_optionName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// optionValue {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.OptionValue == nil {
			o.OptionValue = &oaut.Variant{}
		}
		if err := o.OptionValue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetOptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetOptionRequest structure represents the SetOption operation request
type SetOptionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// optionName: MUST be set either to "AllowedServiceID" or "AllowWarningUI".
	OptionName *oaut.String `idl:"name:optionName" json:"option_name"`
	// optionValue: If optionName is "AllowedServiceID", the vt member MUST be set to VT_BSTR,
	// as specified in [MS-OAUT] section 2.2.7, and the bstrVal member MUST contain the
	// identifier of the service to which to restrict the IUpdateServiceManager::RegisterServiceWithAU
	// (section 3.44.4.2) method.
	OptionValue *oaut.Variant `idl:"name:optionValue" json:"option_value"`
}

func (o *SetOptionRequest) xxx_ToOp(ctx context.Context, op *xxx_SetOptionOperation) *xxx_SetOptionOperation {
	if op == nil {
		op = &xxx_SetOptionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.OptionName = o.OptionName
	op.OptionValue = o.OptionValue
	return op
}

func (o *SetOptionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOptionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.OptionName = op.OptionName
	o.OptionValue = op.OptionValue
}
func (o *SetOptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetOptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOptionResponse structure represents the SetOption operation response
type SetOptionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetOption return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetOptionResponse) xxx_ToOp(ctx context.Context, op *xxx_SetOptionOperation) *xxx_SetOptionOperation {
	if op == nil {
		op = &xxx_SetOptionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetOptionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOptionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetOptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetOptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
