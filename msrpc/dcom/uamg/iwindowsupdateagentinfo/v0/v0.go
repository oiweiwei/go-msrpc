package iwindowsupdateagentinfo

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
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IWindowsUpdateAgentInfo interface identifier 85713fa1-7796-4fa2-be3b-e2d6124dd373
	WindowsUpdateAgentInfoIID = &dcom.IID{Data1: 0x85713fa1, Data2: 0x7796, Data3: 0x4fa2, Data4: []byte{0xbe, 0x3b, 0xe2, 0xd6, 0x12, 0x4d, 0xd3, 0x73}}
	// Syntax UUID
	WindowsUpdateAgentInfoSyntaxUUID = &uuid.UUID{TimeLow: 0x85713fa1, TimeMid: 0x7796, TimeHiAndVersion: 0x4fa2, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0x3b, Node: [6]uint8{0xe2, 0xd6, 0x12, 0x4d, 0xd3, 0x73}}
	// Syntax ID
	WindowsUpdateAgentInfoSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: WindowsUpdateAgentInfoSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWindowsUpdateAgentInfo interface.
type WindowsUpdateAgentInfoClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The IWindowsUpdateAgentInfo::GetInfo (opnum 8) method retrieves version information
	// for the server side of the protocol and the update agent.
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
	GetInfo(context.Context, *GetInfoRequest, ...dcerpc.CallOption) (*GetInfoResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) WindowsUpdateAgentInfoClient
}

type xxx_DefaultWindowsUpdateAgentInfoClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultWindowsUpdateAgentInfoClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultWindowsUpdateAgentInfoClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...dcerpc.CallOption) (*GetInfoResponse, error) {
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
	out := &GetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWindowsUpdateAgentInfoClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWindowsUpdateAgentInfoClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultWindowsUpdateAgentInfoClient) IPID(ctx context.Context, ipid *dcom.IPID) WindowsUpdateAgentInfoClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultWindowsUpdateAgentInfoClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewWindowsUpdateAgentInfoClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WindowsUpdateAgentInfoClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WindowsUpdateAgentInfoSyntaxV0_0))...)
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
	return &xxx_DefaultWindowsUpdateAgentInfoClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetInfoOperation structure represents the GetInfo operation
type xxx_GetInfoOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	VarInfoID   *oaut.Variant  `idl:"name:varInfoIdentifier" json:"var_info_id"`
	ReturnValue *oaut.Variant  `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInfoOperation) OpNum() int { return 7 }

func (o *xxx_GetInfoOperation) OpName() string { return "/IWindowsUpdateAgentInfo/v0/GetInfo" }

func (o *xxx_GetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// varInfoIdentifier {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.VarInfoID != nil {
			if err := o.VarInfoID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// varInfoIdentifier {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.VarInfoID == nil {
			o.VarInfoID = &oaut.Variant{}
		}
		if err := o.VarInfoID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &oaut.Variant{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**oaut.Variant) }
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

// GetInfoRequest structure represents the GetInfo operation request
type GetInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// varInfoIdentifier: A VARIANT ([MS-OAUT] section 2.2.29.2) containing a string specifying
	// the type of version information to retrieve. The vt member of varInfoIdentifier MUST
	// be set to VT_BSTR, and the BSTR ([MS-OAUT] section 2.2.23) stored in the bstrVal
	// member MUST case-insensitively compare equal to one of the following values.
	//
	//	+----------------------+--------------------------------------------------------------+
	//	|       BSTRVAL        |                      INFORMATION TO BE                       |
	//	|        VALUE         |                          RETRIEVED                           |
	//	+----------------------+--------------------------------------------------------------+
	//	+----------------------+--------------------------------------------------------------+
	//	| ApiMajorVersion      | The major version number of the server side of the protocol. |
	//	+----------------------+--------------------------------------------------------------+
	//	| ApiMinorVersion      | The minor version number of the server side of the protocol. |
	//	+----------------------+--------------------------------------------------------------+
	//	| ProductVersionString | The version of the update agent.                             |
	//	+----------------------+--------------------------------------------------------------+
	VarInfoID *oaut.Variant `idl:"name:varInfoIdentifier" json:"var_info_id"`
}

func (o *GetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetInfoOperation) *xxx_GetInfoOperation {
	if op == nil {
		op = &xxx_GetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.VarInfoID = o.VarInfoID
	return op
}

func (o *GetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VarInfoID = op.VarInfoID
}
func (o *GetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInfoResponse structure represents the GetInfo operation response
type GetInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: A VARIANT containing the requested version information, as specified by the
	// following table.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|    VARINFOIDENTIFIER BSTRVAL    |                                                                                  |
	//	|              VALUE              |                                      RETVAL                                      |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| ApiMajorVersion                 | The vt member MUST be set to VT_I4. The lVal member SHOULD<52> be set to the     |
	//	|                                 | server's major version.                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| ApiMinorVersion                 | The vt member MUST be set to VT_I4. The lVal member SHOULD<53> be set to the     |
	//	|                                 | server's minor version.                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| ProductVersionString            | The vt member MUST be set to VT_BSTR. The bstrVal member SHOULD<54> be set to    |
	//	|                                 | the version of the update agent.                                                 |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	ReturnValue *oaut.Variant `idl:"name:retval" json:"return_value"`
	// Return: The GetInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetInfoOperation) *xxx_GetInfoOperation {
	if op == nil {
		op = &xxx_GetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
