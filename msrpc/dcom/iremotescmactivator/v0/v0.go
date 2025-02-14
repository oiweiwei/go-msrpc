package iremotescmactivator

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
)

var (
	// import guard
	GoPackage = "dcom"
)

var (
	// Syntax UUID
	RemoteSCMActivatorSyntaxUUID = &uuid.UUID{TimeLow: 0x1a0, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	RemoteSCMActivatorSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteSCMActivatorSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRemoteSCMActivator interface.
type RemoteSCMActivatorClient interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

	// The RemoteGetClassObject (Opnum 3) method is used by clients to create an object
	// reference for the class factory object.
	RemoteGetClassObject(context.Context, *RemoteGetClassObjectRequest, ...dcerpc.CallOption) (*RemoteGetClassObjectResponse, error)

	// The RemoteCreateInstance (Opnum 4) method is used by clients to create an object
	// reference for the actual object.
	RemoteCreateInstance(context.Context, *RemoteCreateInstanceRequest, ...dcerpc.CallOption) (*RemoteCreateInstanceResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultRemoteSCMActivatorClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultRemoteSCMActivatorClient) RemoteGetClassObject(ctx context.Context, in *RemoteGetClassObjectRequest, opts ...dcerpc.CallOption) (*RemoteGetClassObjectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoteGetClassObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteSCMActivatorClient) RemoteCreateInstance(ctx context.Context, in *RemoteCreateInstanceRequest, opts ...dcerpc.CallOption) (*RemoteCreateInstanceResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoteCreateInstanceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteSCMActivatorClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteSCMActivatorClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewRemoteSCMActivatorClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteSCMActivatorClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteSCMActivatorSyntaxV0_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultRemoteSCMActivatorClient{cc: cc}, nil
}

// xxx_RemoteGetClassObjectOperation structure represents the RemoteGetClassObject operation
type xxx_RemoteGetClassObjectOperation struct {
	ORPCThis         *dcom.ORPCThis         `idl:"name:orpcthis" json:"orpc_this"`
	ORPCThat         *dcom.ORPCThat         `idl:"name:orpcthat" json:"orpc_that"`
	ActPropertiesIn  *dcom.InterfacePointer `idl:"name:pActProperties;pointer:unique" json:"act_properties_in"`
	ActPropertiesOut *dcom.InterfacePointer `idl:"name:ppActProperties" json:"act_properties_out"`
	Return           int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteGetClassObjectOperation) OpNum() int { return 3 }

func (o *xxx_RemoteGetClassObjectOperation) OpName() string {
	return "/IRemoteSCMActivator/v0/RemoteGetClassObject"
}

func (o *xxx_RemoteGetClassObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteGetClassObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// orpcthis {in} (1:{pointer=ref}*(1))(2:{alias=ORPCTHIS}(struct))
	{
		if o.ORPCThis != nil {
			if err := o.ORPCThis.MarshalNDR(ctx, w); err != nil {
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
	// pActProperties {in} (1:{pointer=unique}*(1))(2:{alias=MInterfacePointer}(struct))
	{
		if o.ActPropertiesIn != nil {
			_ptr_pActProperties := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ActPropertiesIn != nil {
					if err := o.ActPropertiesIn.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ActPropertiesIn, _ptr_pActProperties); err != nil {
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

func (o *xxx_RemoteGetClassObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// orpcthis {in} (1:{pointer=ref}*(1))(2:{alias=ORPCTHIS}(struct))
	{
		if o.ORPCThis == nil {
			o.ORPCThis = &dcom.ORPCThis{}
		}
		if err := o.ORPCThis.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pActProperties {in} (1:{pointer=unique}*(1))(2:{alias=MInterfacePointer}(struct))
	{
		_ptr_pActProperties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ActPropertiesIn == nil {
				o.ActPropertiesIn = &dcom.InterfacePointer{}
			}
			if err := o.ActPropertiesIn.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pActProperties := func(ptr interface{}) { o.ActPropertiesIn = *ptr.(**dcom.InterfacePointer) }
		if err := w.ReadPointer(&o.ActPropertiesIn, _s_pActProperties, _ptr_pActProperties); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteGetClassObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteGetClassObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// orpcthat {out} (1:{pointer=ref}*(1))(2:{alias=ORPCTHAT}(struct))
	{
		if o.ORPCThat != nil {
			if err := o.ORPCThat.MarshalNDR(ctx, w); err != nil {
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
	// ppActProperties {out} (1:{pointer=ref}*(2)*(1))(2:{alias=MInterfacePointer}(struct))
	{
		if o.ActPropertiesOut != nil {
			_ptr_ppActProperties := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ActPropertiesOut != nil {
					if err := o.ActPropertiesOut.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ActPropertiesOut, _ptr_ppActProperties); err != nil {
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

func (o *xxx_RemoteGetClassObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// orpcthat {out} (1:{pointer=ref}*(1))(2:{alias=ORPCTHAT}(struct))
	{
		if o.ORPCThat == nil {
			o.ORPCThat = &dcom.ORPCThat{}
		}
		if err := o.ORPCThat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppActProperties {out} (1:{pointer=ref}*(2)*(1))(2:{alias=MInterfacePointer}(struct))
	{
		_ptr_ppActProperties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ActPropertiesOut == nil {
				o.ActPropertiesOut = &dcom.InterfacePointer{}
			}
			if err := o.ActPropertiesOut.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppActProperties := func(ptr interface{}) { o.ActPropertiesOut = *ptr.(**dcom.InterfacePointer) }
		if err := w.ReadPointer(&o.ActPropertiesOut, _s_ppActProperties, _ptr_ppActProperties); err != nil {
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

// RemoteGetClassObjectRequest structure represents the RemoteGetClassObject operation request
type RemoteGetClassObjectRequest struct {
	// orpcthis:  This MUST specify an ORPCTHIS. The COMVERSION field SHOULD contain the
	// negotiated version as specified in section 2.2.11. The extensions field MUST be set
	// to NULL.
	ORPCThis *dcom.ORPCThis `idl:"name:orpcthis" json:"orpc_this"`
	// pActProperties: This MUST specify an MInterfacePointer that MUST contain an OBJREF_CUSTOM
	// with a CLSID field set to CLSID_ActivationPropertiesIn (section 1.9) and a pObjectData
	// field that MUST contain an activation properties BLOB (section 2.2.22). The iid field
	// of the OBJREF portion of the structure MUST be set to IID_IActivationPropertiesIn
	// (see section 1.9). An implementation MAY use this value as the IID of an interface
	// with the local IDL attribute (section 2.2.27).<73>
	//
	// The activation properties BLOB MUST contain properties marked Required in the following
	// table and MAY contain properties marked Optional.
	//
	//	+---------------------------+------------+----------------------+
	//	|         PROPERTY          |            |     REQUIRED OR      |
	//	|           NAME            |  SECTION   |       OPTIONAL       |
	//	|                           |            |                      |
	//	+---------------------------+------------+----------------------+
	//	+---------------------------+------------+----------------------+
	//	| InstantiationInfoData     | 2.2.22.2.1 | Required             |
	//	+---------------------------+------------+----------------------+
	//	| ScmRequestInfoData        | 2.2.22.2.4 | Required             |
	//	+---------------------------+------------+----------------------+
	//	| LocationInfoData          | 2.2.22.2.6 | Required             |
	//	+---------------------------+------------+----------------------+
	//	| SecurityInfoData          | 2.2.22.2.7 | Optional             |
	//	+---------------------------+------------+----------------------+
	//	| ActivationContextInfoData | 2.2.22.2.5 | Optional             |
	//	+---------------------------+------------+----------------------+
	//	| InstanceInfoData          | 2.2.22.2.3 | Optional             |
	//	+---------------------------+------------+----------------------+
	//	| SpecialPropertiesData     | 2.2.22.2.2 | Optional             |
	//	+---------------------------+------------+----------------------+
	ActPropertiesIn *dcom.InterfacePointer `idl:"name:pActProperties;pointer:unique" json:"act_properties_in"`
}

func (o *RemoteGetClassObjectRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoteGetClassObjectOperation) *xxx_RemoteGetClassObjectOperation {
	if op == nil {
		op = &xxx_RemoteGetClassObjectOperation{}
	}
	if o == nil {
		return op
	}
	o.ORPCThis = op.ORPCThis
	o.ActPropertiesIn = op.ActPropertiesIn
	return op
}

func (o *RemoteGetClassObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteGetClassObjectOperation) {
	if o == nil {
		return
	}
	o.ORPCThis = op.ORPCThis
	o.ActPropertiesIn = op.ActPropertiesIn
}
func (o *RemoteGetClassObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoteGetClassObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteGetClassObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteGetClassObjectResponse structure represents the RemoteGetClassObject operation response
type RemoteGetClassObjectResponse struct {
	// orpcthat: This MUST contain an ORPCTHAT. The extensions field MUST be set to NULL.
	ORPCThat *dcom.ORPCThat `idl:"name:orpcthat" json:"orpc_that"`
	// ppActProperties:  This MUST contain an MInterfacePointer that MUST contain an OBJREF_CUSTOM
	// with a CLSID field set to CLSID_ActivationPropertiesOut (section 1.9) and a pObjectData
	// field that MUST contain an activation properties BLOB (section 2.2.22). The iid field
	// of the OBJREF portion of the structure MUST be set to IID_ IActivationPropertiesOut
	// (see section 1.9). An implementation MAY use this value as the IID of an interface
	// with the local IDL attribute (section 2.2.27).<75>
	//
	// The activation properties BLOB MUST contain all properties listed in the following
	// table. Clients SHOULD ignore properties that they do not recognize.
	//
	//	+------------------+------------+----------------------+
	//	|     PROPERTY     |            |     REQUIRED OR      |
	//	|       NAME       |  SECTION   |       OPTIONAL       |
	//	|                  |            |                      |
	//	+------------------+------------+----------------------+
	//	+------------------+------------+----------------------+
	//	| ScmReplyInfoData | 2.2.22.2.8 | Required             |
	//	+------------------+------------+----------------------+
	//	| PropsOutInfo     | 2.2.22.2.9 | Required             |
	//	+------------------+------------+----------------------+
	ActPropertiesOut *dcom.InterfacePointer `idl:"name:ppActProperties" json:"act_properties_out"`
	// Return: The RemoteGetClassObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoteGetClassObjectResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoteGetClassObjectOperation) *xxx_RemoteGetClassObjectOperation {
	if op == nil {
		op = &xxx_RemoteGetClassObjectOperation{}
	}
	if o == nil {
		return op
	}
	o.ORPCThat = op.ORPCThat
	o.ActPropertiesOut = op.ActPropertiesOut
	o.Return = op.Return
	return op
}

func (o *RemoteGetClassObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteGetClassObjectOperation) {
	if o == nil {
		return
	}
	o.ORPCThat = op.ORPCThat
	o.ActPropertiesOut = op.ActPropertiesOut
	o.Return = op.Return
}
func (o *RemoteGetClassObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoteGetClassObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteGetClassObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoteCreateInstanceOperation structure represents the RemoteCreateInstance operation
type xxx_RemoteCreateInstanceOperation struct {
	ORPCThis         *dcom.ORPCThis         `idl:"name:orpcthis" json:"orpc_this"`
	ORPCThat         *dcom.ORPCThat         `idl:"name:orpcthat" json:"orpc_that"`
	UnknownOuter     *dcom.InterfacePointer `idl:"name:pUnkOuter;pointer:unique" json:"unknown_outer"`
	ActPropertiesIn  *dcom.InterfacePointer `idl:"name:pActProperties;pointer:unique" json:"act_properties_in"`
	ActPropertiesOut *dcom.InterfacePointer `idl:"name:ppActProperties" json:"act_properties_out"`
	Return           int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteCreateInstanceOperation) OpNum() int { return 4 }

func (o *xxx_RemoteCreateInstanceOperation) OpName() string {
	return "/IRemoteSCMActivator/v0/RemoteCreateInstance"
}

func (o *xxx_RemoteCreateInstanceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteCreateInstanceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// orpcthis {in} (1:{pointer=ref}*(1))(2:{alias=ORPCTHIS}(struct))
	{
		if o.ORPCThis != nil {
			if err := o.ORPCThis.MarshalNDR(ctx, w); err != nil {
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
	// pUnkOuter {in} (1:{pointer=unique}*(1))(2:{alias=MInterfacePointer}(struct))
	{
		if o.UnknownOuter != nil {
			_ptr_pUnkOuter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UnknownOuter != nil {
					if err := o.UnknownOuter.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UnknownOuter, _ptr_pUnkOuter); err != nil {
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
	// pActProperties {in} (1:{pointer=unique}*(1))(2:{alias=MInterfacePointer}(struct))
	{
		if o.ActPropertiesIn != nil {
			_ptr_pActProperties := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ActPropertiesIn != nil {
					if err := o.ActPropertiesIn.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ActPropertiesIn, _ptr_pActProperties); err != nil {
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

func (o *xxx_RemoteCreateInstanceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// orpcthis {in} (1:{pointer=ref}*(1))(2:{alias=ORPCTHIS}(struct))
	{
		if o.ORPCThis == nil {
			o.ORPCThis = &dcom.ORPCThis{}
		}
		if err := o.ORPCThis.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pUnkOuter {in} (1:{pointer=unique}*(1))(2:{alias=MInterfacePointer}(struct))
	{
		_ptr_pUnkOuter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UnknownOuter == nil {
				o.UnknownOuter = &dcom.InterfacePointer{}
			}
			if err := o.UnknownOuter.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pUnkOuter := func(ptr interface{}) { o.UnknownOuter = *ptr.(**dcom.InterfacePointer) }
		if err := w.ReadPointer(&o.UnknownOuter, _s_pUnkOuter, _ptr_pUnkOuter); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pActProperties {in} (1:{pointer=unique}*(1))(2:{alias=MInterfacePointer}(struct))
	{
		_ptr_pActProperties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ActPropertiesIn == nil {
				o.ActPropertiesIn = &dcom.InterfacePointer{}
			}
			if err := o.ActPropertiesIn.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pActProperties := func(ptr interface{}) { o.ActPropertiesIn = *ptr.(**dcom.InterfacePointer) }
		if err := w.ReadPointer(&o.ActPropertiesIn, _s_pActProperties, _ptr_pActProperties); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteCreateInstanceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteCreateInstanceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// orpcthat {out} (1:{pointer=ref}*(1))(2:{alias=ORPCTHAT}(struct))
	{
		if o.ORPCThat != nil {
			if err := o.ORPCThat.MarshalNDR(ctx, w); err != nil {
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
	// ppActProperties {out} (1:{pointer=ref}*(2)*(1))(2:{alias=MInterfacePointer}(struct))
	{
		if o.ActPropertiesOut != nil {
			_ptr_ppActProperties := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ActPropertiesOut != nil {
					if err := o.ActPropertiesOut.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ActPropertiesOut, _ptr_ppActProperties); err != nil {
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

func (o *xxx_RemoteCreateInstanceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// orpcthat {out} (1:{pointer=ref}*(1))(2:{alias=ORPCTHAT}(struct))
	{
		if o.ORPCThat == nil {
			o.ORPCThat = &dcom.ORPCThat{}
		}
		if err := o.ORPCThat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppActProperties {out} (1:{pointer=ref}*(2)*(1))(2:{alias=MInterfacePointer}(struct))
	{
		_ptr_ppActProperties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ActPropertiesOut == nil {
				o.ActPropertiesOut = &dcom.InterfacePointer{}
			}
			if err := o.ActPropertiesOut.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppActProperties := func(ptr interface{}) { o.ActPropertiesOut = *ptr.(**dcom.InterfacePointer) }
		if err := w.ReadPointer(&o.ActPropertiesOut, _s_ppActProperties, _ptr_ppActProperties); err != nil {
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

// RemoteCreateInstanceRequest structure represents the RemoteCreateInstance operation request
type RemoteCreateInstanceRequest struct {
	// orpcthis: This MUST specify an ORPCTHIS. The COMVERSION field SHOULD contain the
	// negotiated version as defined in section 1.7. The extensions field MUST be set to
	// NULL.
	ORPCThis *dcom.ORPCThis `idl:"name:orpcthis" json:"orpc_this"`
	// pUnkOuter:  This MUST be NULL and MUST be ignored by the recipient.
	UnknownOuter *dcom.InterfacePointer `idl:"name:pUnkOuter;pointer:unique" json:"unknown_outer"`
	// pActProperties: This MUST specify an MInterfacePointer that MUST contain an OBJREF_CUSTOM
	// with a CLSID field set to CLSID_ActivationPropertiesIn (see section 1.9) and a pObjectData
	// field that MUST contain an activation properties BLOB (see section 2.2.22). The iid
	// field of the OBJREF portion of the structure MUST be set to IID_IActivationPropertiesIn
	// (see section 1.9). An implementation MAY use this value as the IID of an interface
	// with the local IDL attribute (section 2.2.27).<76>
	//
	// The activation properties BLOB MUST contain properties that are marked Required in
	// the following table and MAY contain properties that are marked Optional.
	//
	//	+---------------------------+------------+----------------------+
	//	|         PROPERTY          |            |     REQUIRED OR      |
	//	|           NAME            |  SECTION   |       OPTIONAL       |
	//	|                           |            |                      |
	//	+---------------------------+------------+----------------------+
	//	+---------------------------+------------+----------------------+
	//	| InstantiationInfoData     | 2.2.22.2.1 | Required             |
	//	+---------------------------+------------+----------------------+
	//	| ScmRequestInfoData        | 2.2.22.2.4 | Required             |
	//	+---------------------------+------------+----------------------+
	//	| LocationInfoData          | 2.2.22.2.6 | Required             |
	//	+---------------------------+------------+----------------------+
	//	| SecurityInfoData          | 2.2.22.2.7 | Optional             |
	//	+---------------------------+------------+----------------------+
	//	| ActivationContextInfoData | 2.2.22.2.5 | Optional             |
	//	+---------------------------+------------+----------------------+
	//	| InstanceInfoData          | 2.2.22.2.3 | Optional             |
	//	+---------------------------+------------+----------------------+
	//	| SpecialPropertiesData     | 2.2.22.2.2 | Optional             |
	//	+---------------------------+------------+----------------------+
	ActPropertiesIn *dcom.InterfacePointer `idl:"name:pActProperties;pointer:unique" json:"act_properties_in"`
}

func (o *RemoteCreateInstanceRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoteCreateInstanceOperation) *xxx_RemoteCreateInstanceOperation {
	if op == nil {
		op = &xxx_RemoteCreateInstanceOperation{}
	}
	if o == nil {
		return op
	}
	o.ORPCThis = op.ORPCThis
	o.UnknownOuter = op.UnknownOuter
	o.ActPropertiesIn = op.ActPropertiesIn
	return op
}

func (o *RemoteCreateInstanceRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteCreateInstanceOperation) {
	if o == nil {
		return
	}
	o.ORPCThis = op.ORPCThis
	o.UnknownOuter = op.UnknownOuter
	o.ActPropertiesIn = op.ActPropertiesIn
}
func (o *RemoteCreateInstanceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoteCreateInstanceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteCreateInstanceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteCreateInstanceResponse structure represents the RemoteCreateInstance operation response
type RemoteCreateInstanceResponse struct {
	// orpcthat: This MUST contain an ORPCTHAT. The extensions field MUST be set to NULL.
	ORPCThat *dcom.ORPCThat `idl:"name:orpcthat" json:"orpc_that"`
	// ppActProperties:  This MUST contain an MInterfacePointer that MUST contain an OBJREF_CUSTOM
	// with a CLSID field set to CLSID_ActivationPropertiesOut (see section 1.9) and a pObjectData
	// field that MUST contain an activation properties BLOB (see section 2.2.22). The iid
	// field of the OBJREF portion of the structure MUST be set to IID_IActivationPropertiesOut
	// (see section 1.9). An implementation MAY use this value as the IID of an interface
	// with the local IDL attribute (section 2.2.27).<78>
	//
	// The activation properties BLOB MUST contain all properties listed in the following
	// table. Clients SHOULD ignore properties that they do not recognize.
	//
	//	+------------------+------------+----------------------+
	//	|     PROPERTY     |            |     REQUIRED OR      |
	//	|       NAME       |  SECTION   |       OPTIONAL       |
	//	|                  |            |                      |
	//	+------------------+------------+----------------------+
	//	+------------------+------------+----------------------+
	//	| ScmReplyInfoData | 2.2.22.2.8 | Required             |
	//	+------------------+------------+----------------------+
	//	| PropsOutInfo     | 2.2.22.2.9 | Required             |
	//	+------------------+------------+----------------------+
	ActPropertiesOut *dcom.InterfacePointer `idl:"name:ppActProperties" json:"act_properties_out"`
	// Return: The RemoteCreateInstance return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoteCreateInstanceResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoteCreateInstanceOperation) *xxx_RemoteCreateInstanceOperation {
	if op == nil {
		op = &xxx_RemoteCreateInstanceOperation{}
	}
	if o == nil {
		return op
	}
	o.ORPCThat = op.ORPCThat
	o.ActPropertiesOut = op.ActPropertiesOut
	o.Return = op.Return
	return op
}

func (o *RemoteCreateInstanceResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteCreateInstanceOperation) {
	if o == nil {
		return
	}
	o.ORPCThat = op.ORPCThat
	o.ActPropertiesOut = op.ActPropertiesOut
	o.Return = op.Return
}
func (o *RemoteCreateInstanceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoteCreateInstanceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteCreateInstanceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
