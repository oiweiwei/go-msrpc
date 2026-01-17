package iwrmremotesessionmgmt

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

var (
	// IWRMRemoteSessionMgmt interface identifier fc910418-55ca-45ef-b264-83d4ce7d30e0
	RemoteSessionManagementIID = &dcom.IID{Data1: 0xfc910418, Data2: 0x55ca, Data3: 0x45ef, Data4: []byte{0xb2, 0x64, 0x83, 0xd4, 0xce, 0x7d, 0x30, 0xe0}}
	// Syntax UUID
	RemoteSessionManagementSyntaxUUID = &uuid.UUID{TimeLow: 0xfc910418, TimeMid: 0x55ca, TimeHiAndVersion: 0x45ef, ClockSeqHiAndReserved: 0xb2, ClockSeqLow: 0x64, Node: [6]uint8{0x83, 0xd4, 0xce, 0x7d, 0x30, 0xe0}}
	// Syntax ID
	RemoteSessionManagementSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteSessionManagementSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWRMRemoteSessionMgmt interface.
type RemoteSessionManagementClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The GetRemoteUserCategories method retrieves user categories information from the
	// WSRM server.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IWRMRemoteSessionMgmt interface methods are specified in section 3.2.4.9.
	GetRemoteUserCategories(context.Context, *GetRemoteUserCategoriesRequest, ...dcerpc.CallOption) (*GetRemoteUserCategoriesResponse, error)

	// The SetRemoteUserCategories method sets user categories information on the WSRM server.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | Operation successful.                                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | One or more arguments are invalid.                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                      | be processed.<114>                                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMRemoteSessionMgmt interface methods are specified in section 3.2.4.9.
	SetRemoteUserCategories(context.Context, *SetRemoteUserCategoriesRequest, ...dcerpc.CallOption) (*SetRemoteUserCategoriesResponse, error)

	// The RefreshRemoteSessionWeights method forces reallocation of CPU quotas for the
	// sessions run by users according to the category type specified in bstrTargetUserSessions.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER | The XML data is invalid or cannot be processed.                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE          | The specified name contains characters that are invalid. The name cannot         |
	//	|                                      | start with a hyphen (-), cannot contain spaces, and cannot contain any of the    |
	//	|                                      | following characters: \ / ? * | : < > " , ;                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | Operation successful.                                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMRemoteSessionMgmt interface methods are specified in section 3.2.4.9.
	RefreshRemoteSessionWeights(context.Context, *RefreshRemoteSessionWeightsRequest, ...dcerpc.CallOption) (*RefreshRemoteSessionWeightsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteSessionManagementClient
}

type xxx_DefaultRemoteSessionManagementClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteSessionManagementClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultRemoteSessionManagementClient) GetRemoteUserCategories(ctx context.Context, in *GetRemoteUserCategoriesRequest, opts ...dcerpc.CallOption) (*GetRemoteUserCategoriesResponse, error) {
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
	out := &GetRemoteUserCategoriesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteSessionManagementClient) SetRemoteUserCategories(ctx context.Context, in *SetRemoteUserCategoriesRequest, opts ...dcerpc.CallOption) (*SetRemoteUserCategoriesResponse, error) {
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
	out := &SetRemoteUserCategoriesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteSessionManagementClient) RefreshRemoteSessionWeights(ctx context.Context, in *RefreshRemoteSessionWeightsRequest, opts ...dcerpc.CallOption) (*RefreshRemoteSessionWeightsResponse, error) {
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
	out := &RefreshRemoteSessionWeightsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteSessionManagementClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteSessionManagementClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRemoteSessionManagementClient) IPID(ctx context.Context, ipid *dcom.IPID) RemoteSessionManagementClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteSessionManagementClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewRemoteSessionManagementClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteSessionManagementClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteSessionManagementSyntaxV0_0))...)
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
	return &xxx_DefaultRemoteSessionManagementClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetRemoteUserCategoriesOperation structure represents the GetRemoteUserCategories operation
type xxx_GetRemoteUserCategoriesOperation struct {
	This                     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat `idl:"name:That" json:"that"`
	RemoteUserCategoriesInfo *oaut.String   `idl:"name:pbstrRemoteUserCategoriesInfo" json:"remote_user_categories_info"`
	Return                   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRemoteUserCategoriesOperation) OpNum() int { return 7 }

func (o *xxx_GetRemoteUserCategoriesOperation) OpName() string {
	return "/IWRMRemoteSessionMgmt/v0/GetRemoteUserCategories"
}

func (o *xxx_GetRemoteUserCategoriesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteUserCategoriesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRemoteUserCategoriesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRemoteUserCategoriesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteUserCategoriesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrRemoteUserCategoriesInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.RemoteUserCategoriesInfo != nil {
			_ptr_pbstrRemoteUserCategoriesInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RemoteUserCategoriesInfo != nil {
					if err := o.RemoteUserCategoriesInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RemoteUserCategoriesInfo, _ptr_pbstrRemoteUserCategoriesInfo); err != nil {
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

func (o *xxx_GetRemoteUserCategoriesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrRemoteUserCategoriesInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrRemoteUserCategoriesInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RemoteUserCategoriesInfo == nil {
				o.RemoteUserCategoriesInfo = &oaut.String{}
			}
			if err := o.RemoteUserCategoriesInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrRemoteUserCategoriesInfo := func(ptr interface{}) { o.RemoteUserCategoriesInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.RemoteUserCategoriesInfo, _s_pbstrRemoteUserCategoriesInfo, _ptr_pbstrRemoteUserCategoriesInfo); err != nil {
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

// GetRemoteUserCategoriesRequest structure represents the GetRemoteUserCategories operation request
type GetRemoteUserCategoriesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRemoteUserCategoriesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRemoteUserCategoriesOperation) *xxx_GetRemoteUserCategoriesOperation {
	if op == nil {
		op = &xxx_GetRemoteUserCategoriesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetRemoteUserCategoriesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRemoteUserCategoriesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRemoteUserCategoriesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRemoteUserCategoriesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRemoteUserCategoriesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRemoteUserCategoriesResponse structure represents the GetRemoteUserCategories operation response
type GetRemoteUserCategoriesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrRemoteUserCategoriesInfo: A pointer to a string that returns categories of remote
	// session users, in the format of a Users element (section 2.2.5.30).
	RemoteUserCategoriesInfo *oaut.String `idl:"name:pbstrRemoteUserCategoriesInfo" json:"remote_user_categories_info"`
	// Return: The GetRemoteUserCategories return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRemoteUserCategoriesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRemoteUserCategoriesOperation) *xxx_GetRemoteUserCategoriesOperation {
	if op == nil {
		op = &xxx_GetRemoteUserCategoriesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.RemoteUserCategoriesInfo = o.RemoteUserCategoriesInfo
	op.Return = o.Return
	return op
}

func (o *GetRemoteUserCategoriesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRemoteUserCategoriesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RemoteUserCategoriesInfo = op.RemoteUserCategoriesInfo
	o.Return = op.Return
}
func (o *GetRemoteUserCategoriesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRemoteUserCategoriesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRemoteUserCategoriesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetRemoteUserCategoriesOperation structure represents the SetRemoteUserCategories operation
type xxx_SetRemoteUserCategoriesOperation struct {
	This                     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat `idl:"name:That" json:"that"`
	RemoteUserCategoriesInfo *oaut.String   `idl:"name:bstrRemoteUserCategoriesInfo" json:"remote_user_categories_info"`
	Return                   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetRemoteUserCategoriesOperation) OpNum() int { return 8 }

func (o *xxx_SetRemoteUserCategoriesOperation) OpName() string {
	return "/IWRMRemoteSessionMgmt/v0/SetRemoteUserCategories"
}

func (o *xxx_SetRemoteUserCategoriesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRemoteUserCategoriesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrRemoteUserCategoriesInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.RemoteUserCategoriesInfo != nil {
			_ptr_bstrRemoteUserCategoriesInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RemoteUserCategoriesInfo != nil {
					if err := o.RemoteUserCategoriesInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RemoteUserCategoriesInfo, _ptr_bstrRemoteUserCategoriesInfo); err != nil {
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

func (o *xxx_SetRemoteUserCategoriesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrRemoteUserCategoriesInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrRemoteUserCategoriesInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RemoteUserCategoriesInfo == nil {
				o.RemoteUserCategoriesInfo = &oaut.String{}
			}
			if err := o.RemoteUserCategoriesInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrRemoteUserCategoriesInfo := func(ptr interface{}) { o.RemoteUserCategoriesInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.RemoteUserCategoriesInfo, _s_bstrRemoteUserCategoriesInfo, _ptr_bstrRemoteUserCategoriesInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRemoteUserCategoriesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRemoteUserCategoriesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetRemoteUserCategoriesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetRemoteUserCategoriesRequest structure represents the SetRemoteUserCategories operation request
type SetRemoteUserCategoriesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrRemoteUserCategoriesInfo: A string that specifies categories of remote session
	// users, in the format of a Users element (section 2.2.5.30).
	RemoteUserCategoriesInfo *oaut.String `idl:"name:bstrRemoteUserCategoriesInfo" json:"remote_user_categories_info"`
}

func (o *SetRemoteUserCategoriesRequest) xxx_ToOp(ctx context.Context, op *xxx_SetRemoteUserCategoriesOperation) *xxx_SetRemoteUserCategoriesOperation {
	if op == nil {
		op = &xxx_SetRemoteUserCategoriesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RemoteUserCategoriesInfo = o.RemoteUserCategoriesInfo
	return op
}

func (o *SetRemoteUserCategoriesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetRemoteUserCategoriesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RemoteUserCategoriesInfo = op.RemoteUserCategoriesInfo
}
func (o *SetRemoteUserCategoriesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetRemoteUserCategoriesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRemoteUserCategoriesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetRemoteUserCategoriesResponse structure represents the SetRemoteUserCategories operation response
type SetRemoteUserCategoriesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetRemoteUserCategories return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetRemoteUserCategoriesResponse) xxx_ToOp(ctx context.Context, op *xxx_SetRemoteUserCategoriesOperation) *xxx_SetRemoteUserCategoriesOperation {
	if op == nil {
		op = &xxx_SetRemoteUserCategoriesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetRemoteUserCategoriesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetRemoteUserCategoriesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetRemoteUserCategoriesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetRemoteUserCategoriesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRemoteUserCategoriesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RefreshRemoteSessionWeightsOperation structure represents the RefreshRemoteSessionWeights operation
type xxx_RefreshRemoteSessionWeightsOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaregetUserSessions *oaut.String   `idl:"name:bstrTaregetUserSessions" json:"tareget_user_sessions"`
	UpdateAll           bool           `idl:"name:bUpdateAll" json:"update_all"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RefreshRemoteSessionWeightsOperation) OpNum() int { return 9 }

func (o *xxx_RefreshRemoteSessionWeightsOperation) OpName() string {
	return "/IWRMRemoteSessionMgmt/v0/RefreshRemoteSessionWeights"
}

func (o *xxx_RefreshRemoteSessionWeightsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshRemoteSessionWeightsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrTaregetUserSessions {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.TaregetUserSessions != nil {
			_ptr_bstrTaregetUserSessions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TaregetUserSessions != nil {
					if err := o.TaregetUserSessions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TaregetUserSessions, _ptr_bstrTaregetUserSessions); err != nil {
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
	// bUpdateAll {in} (1:{alias=BOOL}(int32))
	{
		if !o.UpdateAll {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RefreshRemoteSessionWeightsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrTaregetUserSessions {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrTaregetUserSessions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TaregetUserSessions == nil {
				o.TaregetUserSessions = &oaut.String{}
			}
			if err := o.TaregetUserSessions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrTaregetUserSessions := func(ptr interface{}) { o.TaregetUserSessions = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.TaregetUserSessions, _s_bstrTaregetUserSessions, _ptr_bstrTaregetUserSessions); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bUpdateAll {in} (1:{alias=BOOL}(int32))
	{
		var _bUpdateAll int32
		if err := w.ReadData(&_bUpdateAll); err != nil {
			return err
		}
		o.UpdateAll = _bUpdateAll != 0
	}
	return nil
}

func (o *xxx_RefreshRemoteSessionWeightsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshRemoteSessionWeightsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RefreshRemoteSessionWeightsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RefreshRemoteSessionWeightsRequest structure represents the RefreshRemoteSessionWeights operation request
type RefreshRemoteSessionWeightsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	TaregetUserSessions *oaut.String   `idl:"name:bstrTaregetUserSessions" json:"tareget_user_sessions"`
	// bUpdateAll: A Boolean value that specifies whether to recursively delete all instances
	// of the specified machine group.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | Only the CPU quotas for users specified in bstrTargetUserSessions are            |
	//	|                  | reallocated according to their category type.                                    |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | The CPU quotas for all remote sessions are reallocated according to the category |
	//	|                  | type specified in bstrTargetUserSessions.                                        |
	//	+------------------+----------------------------------------------------------------------------------+
	UpdateAll bool `idl:"name:bUpdateAll" json:"update_all"`
}

func (o *RefreshRemoteSessionWeightsRequest) xxx_ToOp(ctx context.Context, op *xxx_RefreshRemoteSessionWeightsOperation) *xxx_RefreshRemoteSessionWeightsOperation {
	if op == nil {
		op = &xxx_RefreshRemoteSessionWeightsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.TaregetUserSessions = o.TaregetUserSessions
	op.UpdateAll = o.UpdateAll
	return op
}

func (o *RefreshRemoteSessionWeightsRequest) xxx_FromOp(ctx context.Context, op *xxx_RefreshRemoteSessionWeightsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TaregetUserSessions = op.TaregetUserSessions
	o.UpdateAll = op.UpdateAll
}
func (o *RefreshRemoteSessionWeightsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RefreshRemoteSessionWeightsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshRemoteSessionWeightsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RefreshRemoteSessionWeightsResponse structure represents the RefreshRemoteSessionWeights operation response
type RefreshRemoteSessionWeightsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RefreshRemoteSessionWeights return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RefreshRemoteSessionWeightsResponse) xxx_ToOp(ctx context.Context, op *xxx_RefreshRemoteSessionWeightsOperation) *xxx_RefreshRemoteSessionWeightsOperation {
	if op == nil {
		op = &xxx_RefreshRemoteSessionWeightsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RefreshRemoteSessionWeightsResponse) xxx_FromOp(ctx context.Context, op *xxx_RefreshRemoteSessionWeightsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RefreshRemoteSessionWeightsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RefreshRemoteSessionWeightsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshRemoteSessionWeightsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
