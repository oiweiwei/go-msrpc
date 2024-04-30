package icatalogutils

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
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// ICatalogUtils interface identifier 456129e2-1078-11d2-b0f9-00805fc73204
	CatalogUtilsIID = &dcom.IID{Data1: 0x456129e2, Data2: 0x1078, Data3: 0x11d2, Data4: []byte{0xb0, 0xf9, 0x00, 0x80, 0x5f, 0xc7, 0x32, 0x04}}
	// Syntax UUID
	CatalogUtilsSyntaxUUID = &uuid.UUID{TimeLow: 0x456129e2, TimeMid: 0x1078, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0xf9, Node: [6]uint8{0x0, 0x80, 0x5f, 0xc7, 0x32, 0x4}}
	// Syntax ID
	CatalogUtilsSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CatalogUtilsSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICatalogUtils interface.
type CatalogUtilsClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to verify that a user account and password are
	// valid.
	//
	// Return Values: This method MUST return a value in the table below on success, or
	// a failure result, as specified in [MS-ERREF] section 2.1, on failure. All failure
	// results MUST be treated identically.
	//
	//	+--------------------+-------------------------------------------------------------+
	//	|       RETURN       |                                                             |
	//	|     VALUE/CODE     |                         DESCRIPTION                         |
	//	|                    |                                                             |
	//	+--------------------+-------------------------------------------------------------+
	//	+--------------------+-------------------------------------------------------------+
	//	| 0x00000000 S_OK    | The user account and password are valid.                    |
	//	+--------------------+-------------------------------------------------------------+
	//	| 0x00000001 S_FALSE | The user account was not found or the password was invalid. |
	//	+--------------------+-------------------------------------------------------------+
	ValidateUser(context.Context, *ValidateUserRequest, ...dcerpc.CallOption) (*ValidateUserResponse, error)

	// This method is called by a COMA client to synchronize with the server after performing
	// a write operation.
	//
	// This method has no parameters.
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
	WaitForEndWrites(context.Context, *WaitForEndWritesRequest, ...dcerpc.CallOption) (*WaitForEndWritesResponse, error)

	// This method is called by a client to get information about the event classes associated
	// with an IID that are configured in the Global Partition.
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
	GetEventClassesForIID(context.Context, *GetEventClassesForIIDRequest, ...dcerpc.CallOption) (*GetEventClassesForIIDResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CatalogUtilsClient
}

type xxx_DefaultCatalogUtilsClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCatalogUtilsClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCatalogUtilsClient) ValidateUser(ctx context.Context, in *ValidateUserRequest, opts ...dcerpc.CallOption) (*ValidateUserResponse, error) {
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
	out := &ValidateUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogUtilsClient) WaitForEndWrites(ctx context.Context, in *WaitForEndWritesRequest, opts ...dcerpc.CallOption) (*WaitForEndWritesResponse, error) {
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
	out := &WaitForEndWritesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogUtilsClient) GetEventClassesForIID(ctx context.Context, in *GetEventClassesForIIDRequest, opts ...dcerpc.CallOption) (*GetEventClassesForIIDResponse, error) {
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
	out := &GetEventClassesForIIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogUtilsClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCatalogUtilsClient) IPID(ctx context.Context, ipid *dcom.IPID) CatalogUtilsClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCatalogUtilsClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewCatalogUtilsClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CatalogUtilsClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CatalogUtilsSyntaxV0_0))...)
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
	return &xxx_DefaultCatalogUtilsClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ValidateUserOperation structure represents the ValidateUser operation
type xxx_ValidateUserOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	PrincipalName string         `idl:"name:pwszPrincipalName;string;pointer:unique" json:"principal_name"`
	Password      string         `idl:"name:pwszPassword;string;pointer:unique" json:"password"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ValidateUserOperation) OpNum() int { return 3 }

func (o *xxx_ValidateUserOperation) OpName() string { return "/ICatalogUtils/v0/ValidateUser" }

func (o *xxx_ValidateUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszPrincipalName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.PrincipalName != "" {
			_ptr_pwszPrincipalName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.PrincipalName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.PrincipalName, _ptr_pwszPrincipalName); err != nil {
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
	// pwszPassword {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Password != "" {
			_ptr_pwszPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_pwszPassword); err != nil {
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

func (o *xxx_ValidateUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszPrincipalName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszPrincipalName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.PrincipalName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszPrincipalName := func(ptr interface{}) { o.PrincipalName = *ptr.(*string) }
		if err := w.ReadPointer(&o.PrincipalName, _s_pwszPrincipalName, _ptr_pwszPrincipalName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszPassword {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
				return err
			}
			return nil
		})
		_s_pwszPassword := func(ptr interface{}) { o.Password = *ptr.(*string) }
		if err := w.ReadPointer(&o.Password, _s_pwszPassword, _ptr_pwszPassword); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ValidateUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ValidateUserRequest structure represents the ValidateUser operation request
type ValidateUserRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszPrincipalName:  The principal name of the user account.
	PrincipalName string `idl:"name:pwszPrincipalName;string;pointer:unique" json:"principal_name"`
	// pwszPassword:  A password for the user account.
	Password string `idl:"name:pwszPassword;string;pointer:unique" json:"password"`
}

func (o *ValidateUserRequest) xxx_ToOp(ctx context.Context) *xxx_ValidateUserOperation {
	if o == nil {
		return &xxx_ValidateUserOperation{}
	}
	return &xxx_ValidateUserOperation{
		This:          o.This,
		PrincipalName: o.PrincipalName,
		Password:      o.Password,
	}
}

func (o *ValidateUserRequest) xxx_FromOp(ctx context.Context, op *xxx_ValidateUserOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PrincipalName = op.PrincipalName
	o.Password = op.Password
}
func (o *ValidateUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ValidateUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ValidateUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ValidateUserResponse structure represents the ValidateUser operation response
type ValidateUserResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ValidateUser return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ValidateUserResponse) xxx_ToOp(ctx context.Context) *xxx_ValidateUserOperation {
	if o == nil {
		return &xxx_ValidateUserOperation{}
	}
	return &xxx_ValidateUserOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ValidateUserResponse) xxx_FromOp(ctx context.Context, op *xxx_ValidateUserOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ValidateUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ValidateUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ValidateUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WaitForEndWritesOperation structure represents the WaitForEndWrites operation
type xxx_WaitForEndWritesOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_WaitForEndWritesOperation) OpNum() int { return 4 }

func (o *xxx_WaitForEndWritesOperation) OpName() string { return "/ICatalogUtils/v0/WaitForEndWrites" }

func (o *xxx_WaitForEndWritesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForEndWritesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_WaitForEndWritesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_WaitForEndWritesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForEndWritesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_WaitForEndWritesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// WaitForEndWritesRequest structure represents the WaitForEndWrites operation request
type WaitForEndWritesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *WaitForEndWritesRequest) xxx_ToOp(ctx context.Context) *xxx_WaitForEndWritesOperation {
	if o == nil {
		return &xxx_WaitForEndWritesOperation{}
	}
	return &xxx_WaitForEndWritesOperation{
		This: o.This,
	}
}

func (o *WaitForEndWritesRequest) xxx_FromOp(ctx context.Context, op *xxx_WaitForEndWritesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *WaitForEndWritesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *WaitForEndWritesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForEndWritesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WaitForEndWritesResponse structure represents the WaitForEndWrites operation response
type WaitForEndWritesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The WaitForEndWrites return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WaitForEndWritesResponse) xxx_ToOp(ctx context.Context) *xxx_WaitForEndWritesOperation {
	if o == nil {
		return &xxx_WaitForEndWritesOperation{}
	}
	return &xxx_WaitForEndWritesOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *WaitForEndWritesResponse) xxx_FromOp(ctx context.Context, op *xxx_WaitForEndWritesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *WaitForEndWritesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *WaitForEndWritesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForEndWritesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEventClassesForIIDOperation structure represents the GetEventClassesForIID operation
type xxx_GetEventClassesForIIDOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	IID          string         `idl:"name:wszIID;string;pointer:unique" json:"iid"`
	ClassesCount uint32         `idl:"name:pcClasses" json:"classes_count"`
	ClassIDs     []string       `idl:"name:pawszCLSIDs;size_is:(, pcClasses);string" json:"class_ids"`
	ProgIDs      []string       `idl:"name:pawszProgIDs;size_is:(, pcClasses);string" json:"prog_ids"`
	Descriptions []string       `idl:"name:pawszDescriptions;size_is:(, pcClasses);string" json:"descriptions"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventClassesForIIDOperation) OpNum() int { return 5 }

func (o *xxx_GetEventClassesForIIDOperation) OpName() string {
	return "/ICatalogUtils/v0/GetEventClassesForIID"
}

func (o *xxx_GetEventClassesForIIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassesForIIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// wszIID {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.IID != "" {
			_ptr_wszIID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.IID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.IID, _ptr_wszIID); err != nil {
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

func (o *xxx_GetEventClassesForIIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// wszIID {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_wszIID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.IID); err != nil {
				return err
			}
			return nil
		})
		_s_wszIID := func(ptr interface{}) { o.IID = *ptr.(*string) }
		if err := w.ReadPointer(&o.IID, _s_wszIID, _ptr_wszIID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassesForIIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ClassIDs != nil && o.ClassesCount == 0 {
		o.ClassesCount = uint32(len(o.ClassIDs))
	}
	if o.ProgIDs != nil && o.ClassesCount == 0 {
		o.ClassesCount = uint32(len(o.ProgIDs))
	}
	if o.Descriptions != nil && o.ClassesCount == 0 {
		o.ClassesCount = uint32(len(o.Descriptions))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassesForIIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcClasses {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClassesCount); err != nil {
			return err
		}
	}
	// pawszCLSIDs {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		if o.ClassIDs != nil || o.ClassesCount > 0 {
			_ptr_pawszCLSIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ClassesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ClassIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ClassIDs[i1] != "" {
						_ptr_pawszCLSIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.ClassIDs[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.ClassIDs[i1], _ptr_pawszCLSIDs); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ClassIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClassIDs, _ptr_pawszCLSIDs); err != nil {
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
	// pawszProgIDs {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		if o.ProgIDs != nil || o.ClassesCount > 0 {
			_ptr_pawszProgIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ClassesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ProgIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ProgIDs[i1] != "" {
						_ptr_pawszProgIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.ProgIDs[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.ProgIDs[i1], _ptr_pawszProgIDs); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ProgIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProgIDs, _ptr_pawszProgIDs); err != nil {
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
	// pawszDescriptions {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		if o.Descriptions != nil || o.ClassesCount > 0 {
			_ptr_pawszDescriptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ClassesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Descriptions {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Descriptions[i1] != "" {
						_ptr_pawszDescriptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.Descriptions[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.Descriptions[i1], _ptr_pawszDescriptions); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Descriptions); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Descriptions, _ptr_pawszDescriptions); err != nil {
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

func (o *xxx_GetEventClassesForIIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcClasses {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClassesCount); err != nil {
			return err
		}
	}
	// pawszCLSIDs {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		_ptr_pawszCLSIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ClassIDs", sizeInfo[0])
			}
			o.ClassIDs = make([]string, sizeInfo[0])
			for i1 := range o.ClassIDs {
				i1 := i1
				_ptr_pawszCLSIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.ClassIDs[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_pawszCLSIDs := func(ptr interface{}) { o.ClassIDs[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.ClassIDs[i1], _s_pawszCLSIDs, _ptr_pawszCLSIDs); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pawszCLSIDs := func(ptr interface{}) { o.ClassIDs = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.ClassIDs, _s_pawszCLSIDs, _ptr_pawszCLSIDs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pawszProgIDs {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		_ptr_pawszProgIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ProgIDs", sizeInfo[0])
			}
			o.ProgIDs = make([]string, sizeInfo[0])
			for i1 := range o.ProgIDs {
				i1 := i1
				_ptr_pawszProgIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.ProgIDs[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_pawszProgIDs := func(ptr interface{}) { o.ProgIDs[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.ProgIDs[i1], _s_pawszProgIDs, _ptr_pawszProgIDs); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pawszProgIDs := func(ptr interface{}) { o.ProgIDs = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.ProgIDs, _s_pawszProgIDs, _ptr_pawszProgIDs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pawszDescriptions {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		_ptr_pawszDescriptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Descriptions", sizeInfo[0])
			}
			o.Descriptions = make([]string, sizeInfo[0])
			for i1 := range o.Descriptions {
				i1 := i1
				_ptr_pawszDescriptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.Descriptions[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_pawszDescriptions := func(ptr interface{}) { o.Descriptions[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.Descriptions[i1], _s_pawszDescriptions, _ptr_pawszDescriptions); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pawszDescriptions := func(ptr interface{}) { o.Descriptions = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.Descriptions, _s_pawszDescriptions, _ptr_pawszDescriptions); err != nil {
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

// GetEventClassesForIIDRequest structure represents the GetEventClassesForIID operation request
type GetEventClassesForIIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// wszIID: The Curly Braced GUID String Syntax ([MS-DTYP] section 2.3.4.3) representation
	// of an IID for which event classes will be retrieved, or NULL for an empty (zero-length)
	// string to indicate all event classes.
	IID string `idl:"name:wszIID;string;pointer:unique" json:"iid"`
}

func (o *GetEventClassesForIIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetEventClassesForIIDOperation {
	if o == nil {
		return &xxx_GetEventClassesForIIDOperation{}
	}
	return &xxx_GetEventClassesForIIDOperation{
		This: o.This,
		IID:  o.IID,
	}
}

func (o *GetEventClassesForIIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassesForIIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IID = op.IID
}
func (o *GetEventClassesForIIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetEventClassesForIIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassesForIIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEventClassesForIIDResponse structure represents the GetEventClassesForIID operation response
type GetEventClassesForIIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pcClasses: A pointer to a value that, upon successful completion, MUST be set to
	// the number of event classes for which information was returned.
	ClassesCount uint32 `idl:"name:pcClasses" json:"classes_count"`
	// pawszCLSIDs: A pointer to a value that, upon successful completion, MUST be set to
	// an array of Curly Braced GUID String Syntax ([MS-DTYP] section 2.3.4.3) representations
	// of CLSIDs of event classes.
	ClassIDs []string `idl:"name:pawszCLSIDs;size_is:(, pcClasses);string" json:"class_ids"`
	// pawszProgIDs: A pointer to a value that, upon successful completion, MUST be set
	// to an array of ProgIDs of event classes, in the same order as pawszCLSIDs.
	ProgIDs []string `idl:"name:pawszProgIDs;size_is:(, pcClasses);string" json:"prog_ids"`
	// pawszDescriptions: A pointer to a value that, upon successful completion, MUST be
	// set to an array of descriptions of event classes, in the same order as pawszCLSIDs.
	Descriptions []string `idl:"name:pawszDescriptions;size_is:(, pcClasses);string" json:"descriptions"`
	// Return: The GetEventClassesForIID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEventClassesForIIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetEventClassesForIIDOperation {
	if o == nil {
		return &xxx_GetEventClassesForIIDOperation{}
	}
	return &xxx_GetEventClassesForIIDOperation{
		That:         o.That,
		ClassesCount: o.ClassesCount,
		ClassIDs:     o.ClassIDs,
		ProgIDs:      o.ProgIDs,
		Descriptions: o.Descriptions,
		Return:       o.Return,
	}
}

func (o *GetEventClassesForIIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassesForIIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ClassesCount = op.ClassesCount
	o.ClassIDs = op.ClassIDs
	o.ProgIDs = op.ProgIDs
	o.Descriptions = op.Descriptions
	o.Return = op.Return
}
func (o *GetEventClassesForIIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetEventClassesForIIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassesForIIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
