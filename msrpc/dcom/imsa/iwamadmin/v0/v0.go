package iwamadmin

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
	GoPackage = "dcom/imsa"
)

var (
	// IWamAdmin interface identifier 29822ab7-f302-11d0-9953-00c04fd919c1
	WAMAdminIID = &dcom.IID{Data1: 0x29822ab7, Data2: 0xf302, Data3: 0x11d0, Data4: []byte{0x99, 0x53, 0x00, 0xc0, 0x4f, 0xd9, 0x19, 0xc1}}
	// Syntax UUID
	WAMAdminSyntaxUUID = &uuid.UUID{TimeLow: 0x29822ab7, TimeMid: 0xf302, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x53, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd9, 0x19, 0xc1}}
	// Syntax ID
	WAMAdminSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: WAMAdminSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWamAdmin interface.
type WAMAdminClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The AppCreate method creates a new application at the specified metabase path.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// The opnum field value for this method is 3.
	AppCreate(context.Context, *AppCreateRequest, ...dcerpc.CallOption) (*AppCreateResponse, error)

	// The AppDelete method deletes the application from the specified metabase path.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.         |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 4.
	AppDelete(context.Context, *AppDeleteRequest, ...dcerpc.CallOption) (*AppDeleteResponse, error)

	// The AppUnLoad method shuts down the specified application.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// The opnum field value for this method is 5.
	AppUnload(context.Context, *AppUnloadRequest, ...dcerpc.CallOption) (*AppUnloadResponse, error)

	// The AppGetStatus method retrieves the status of the application defined at the specified
	// metabase path.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 6.
	AppGetStatus(context.Context, *AppGetStatusRequest, ...dcerpc.CallOption) (*AppGetStatusResponse, error)

	// The AppDeleteRecoverable method deletes the application from the specified metabase
	// path and saves external state needed to recreate the application if it is recovered.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.         |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 7.
	AppDeleteRecoverable(context.Context, *AppDeleteRecoverableRequest, ...dcerpc.CallOption) (*AppDeleteRecoverableResponse, error)

	// The AppRecover method recreates an application that was deleted by the AppDeleteRecoverable
	// method.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 8.
	AppRecover(context.Context, *AppRecoverRequest, ...dcerpc.CallOption) (*AppRecoverResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) WAMAdminClient
}

type xxx_DefaultWAMAdminClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultWAMAdminClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultWAMAdminClient) AppCreate(ctx context.Context, in *AppCreateRequest, opts ...dcerpc.CallOption) (*AppCreateResponse, error) {
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
	out := &AppCreateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWAMAdminClient) AppDelete(ctx context.Context, in *AppDeleteRequest, opts ...dcerpc.CallOption) (*AppDeleteResponse, error) {
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
	out := &AppDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWAMAdminClient) AppUnload(ctx context.Context, in *AppUnloadRequest, opts ...dcerpc.CallOption) (*AppUnloadResponse, error) {
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
	out := &AppUnloadResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWAMAdminClient) AppGetStatus(ctx context.Context, in *AppGetStatusRequest, opts ...dcerpc.CallOption) (*AppGetStatusResponse, error) {
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
	out := &AppGetStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWAMAdminClient) AppDeleteRecoverable(ctx context.Context, in *AppDeleteRecoverableRequest, opts ...dcerpc.CallOption) (*AppDeleteRecoverableResponse, error) {
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
	out := &AppDeleteRecoverableResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWAMAdminClient) AppRecover(ctx context.Context, in *AppRecoverRequest, opts ...dcerpc.CallOption) (*AppRecoverResponse, error) {
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
	out := &AppRecoverResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWAMAdminClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWAMAdminClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultWAMAdminClient) IPID(ctx context.Context, ipid *dcom.IPID) WAMAdminClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultWAMAdminClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewWAMAdminClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WAMAdminClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WAMAdminSyntaxV0_0))...)
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
	return &xxx_DefaultWAMAdminClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_AppCreateOperation structure represents the AppCreate operation
type xxx_AppCreateOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   string         `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	InProc bool           `idl:"name:fInProc" json:"in_proc"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AppCreateOperation) OpNum() int { return 3 }

func (o *xxx_AppCreateOperation) OpName() string { return "/IWamAdmin/v0/AppCreate" }

func (o *xxx_AppCreateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppCreateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_szMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_szMDPath); err != nil {
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
	// fInProc {in} (1:{alias=BOOL}(int32))
	{
		if !o.InProc {
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

func (o *xxx_AppCreateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_szMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_szMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_szMDPath, _ptr_szMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fInProc {in} (1:{alias=BOOL}(int32))
	{
		var _bInProc int32
		if err := w.ReadData(&_bInProc); err != nil {
			return err
		}
		o.InProc = _bInProc != 0
	}
	return nil
}

func (o *xxx_AppCreateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppCreateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AppCreateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AppCreateRequest structure represents the AppCreate operation request
type AppCreateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szMDPath: A pointer to a Unicode string that contains the metabase path of the application.
	Path string `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	// fInProc:  A flag indicating whether the application runs in the parent IIS server
	// process or in its own process.
	InProc bool `idl:"name:fInProc" json:"in_proc"`
}

func (o *AppCreateRequest) xxx_ToOp(ctx context.Context) *xxx_AppCreateOperation {
	if o == nil {
		return &xxx_AppCreateOperation{}
	}
	return &xxx_AppCreateOperation{
		This:   o.This,
		Path:   o.Path,
		InProc: o.InProc,
	}
}

func (o *AppCreateRequest) xxx_FromOp(ctx context.Context, op *xxx_AppCreateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.InProc = op.InProc
}
func (o *AppCreateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AppCreateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppCreateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AppCreateResponse structure represents the AppCreate operation response
type AppCreateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AppCreate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AppCreateResponse) xxx_ToOp(ctx context.Context) *xxx_AppCreateOperation {
	if o == nil {
		return &xxx_AppCreateOperation{}
	}
	return &xxx_AppCreateOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AppCreateResponse) xxx_FromOp(ctx context.Context, op *xxx_AppCreateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AppCreateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AppCreateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppCreateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AppDeleteOperation structure represents the AppDelete operation
type xxx_AppDeleteOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path      string         `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	Recursive bool           `idl:"name:fRecursive" json:"recursive"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AppDeleteOperation) OpNum() int { return 4 }

func (o *xxx_AppDeleteOperation) OpName() string { return "/IWamAdmin/v0/AppDelete" }

func (o *xxx_AppDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_szMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_szMDPath); err != nil {
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
	// fRecursive {in} (1:{alias=BOOL}(int32))
	{
		if !o.Recursive {
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

func (o *xxx_AppDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_szMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_szMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_szMDPath, _ptr_szMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fRecursive {in} (1:{alias=BOOL}(int32))
	{
		var _bRecursive int32
		if err := w.ReadData(&_bRecursive); err != nil {
			return err
		}
		o.Recursive = _bRecursive != 0
	}
	return nil
}

func (o *xxx_AppDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AppDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AppDeleteRequest structure represents the AppDelete operation request
type AppDeleteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szMDPath: A pointer to a Unicode string that contains the metabase path of the application.
	Path string `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	// fRecursive: A flag indicating whether application definitions are also to be deleted
	// from all subkeys (TRUE) or just from the application at this key (FALSE).
	Recursive bool `idl:"name:fRecursive" json:"recursive"`
}

func (o *AppDeleteRequest) xxx_ToOp(ctx context.Context) *xxx_AppDeleteOperation {
	if o == nil {
		return &xxx_AppDeleteOperation{}
	}
	return &xxx_AppDeleteOperation{
		This:      o.This,
		Path:      o.Path,
		Recursive: o.Recursive,
	}
}

func (o *AppDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_AppDeleteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.Recursive = op.Recursive
}
func (o *AppDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AppDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AppDeleteResponse structure represents the AppDelete operation response
type AppDeleteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AppDelete return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AppDeleteResponse) xxx_ToOp(ctx context.Context) *xxx_AppDeleteOperation {
	if o == nil {
		return &xxx_AppDeleteOperation{}
	}
	return &xxx_AppDeleteOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AppDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_AppDeleteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AppDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AppDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AppUnloadOperation structure represents the AppUnLoad operation
type xxx_AppUnloadOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path      string         `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	Recursive bool           `idl:"name:fRecursive" json:"recursive"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AppUnloadOperation) OpNum() int { return 5 }

func (o *xxx_AppUnloadOperation) OpName() string { return "/IWamAdmin/v0/AppUnLoad" }

func (o *xxx_AppUnloadOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppUnloadOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_szMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_szMDPath); err != nil {
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
	// fRecursive {in} (1:{alias=BOOL}(int32))
	{
		if !o.Recursive {
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

func (o *xxx_AppUnloadOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_szMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_szMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_szMDPath, _ptr_szMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fRecursive {in} (1:{alias=BOOL}(int32))
	{
		var _bRecursive int32
		if err := w.ReadData(&_bRecursive); err != nil {
			return err
		}
		o.Recursive = _bRecursive != 0
	}
	return nil
}

func (o *xxx_AppUnloadOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppUnloadOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AppUnloadOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AppUnloadRequest structure represents the AppUnLoad operation request
type AppUnloadRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szMDPath:  A pointer to a Unicode string that contains the metabase path of the
	// application.
	Path string `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	// fRecursive: A flag indicating whether applications are also unloaded from all subkeys
	// (TRUE) or just from the application at this key (FALSE).
	Recursive bool `idl:"name:fRecursive" json:"recursive"`
}

func (o *AppUnloadRequest) xxx_ToOp(ctx context.Context) *xxx_AppUnloadOperation {
	if o == nil {
		return &xxx_AppUnloadOperation{}
	}
	return &xxx_AppUnloadOperation{
		This:      o.This,
		Path:      o.Path,
		Recursive: o.Recursive,
	}
}

func (o *AppUnloadRequest) xxx_FromOp(ctx context.Context, op *xxx_AppUnloadOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.Recursive = op.Recursive
}
func (o *AppUnloadRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AppUnloadRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppUnloadOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AppUnloadResponse structure represents the AppUnLoad operation response
type AppUnloadResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AppUnLoad return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AppUnloadResponse) xxx_ToOp(ctx context.Context) *xxx_AppUnloadOperation {
	if o == nil {
		return &xxx_AppUnloadOperation{}
	}
	return &xxx_AppUnloadOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AppUnloadResponse) xxx_FromOp(ctx context.Context, op *xxx_AppUnloadOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AppUnloadResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AppUnloadResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppUnloadOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AppGetStatusOperation structure represents the AppGetStatus operation
type xxx_AppGetStatusOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path      string         `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	AppStatus uint32         `idl:"name:pdwAppStatus" json:"app_status"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AppGetStatusOperation) OpNum() int { return 6 }

func (o *xxx_AppGetStatusOperation) OpName() string { return "/IWamAdmin/v0/AppGetStatus" }

func (o *xxx_AppGetStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppGetStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_szMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_szMDPath); err != nil {
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

func (o *xxx_AppGetStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_szMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_szMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_szMDPath, _ptr_szMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppGetStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppGetStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwAppStatus {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AppStatus); err != nil {
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

func (o *xxx_AppGetStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwAppStatus {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AppStatus); err != nil {
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

// AppGetStatusRequest structure represents the AppGetStatus operation request
type AppGetStatusRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szMDPath:  A pointer to a Unicode string that contains the metabase path of the
	// application.
	Path string `idl:"name:szMDPath;string;pointer:unique" json:"path"`
}

func (o *AppGetStatusRequest) xxx_ToOp(ctx context.Context) *xxx_AppGetStatusOperation {
	if o == nil {
		return &xxx_AppGetStatusOperation{}
	}
	return &xxx_AppGetStatusOperation{
		This: o.This,
		Path: o.Path,
	}
}

func (o *AppGetStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_AppGetStatusOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *AppGetStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AppGetStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppGetStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AppGetStatusResponse structure represents the AppGetStatus operation response
type AppGetStatusResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwAppStatus: A pointer to a 32-bit unsigned integer that receives the value indicating
	// the status of the application. This field MUST be set to one of the following values.
	//
	//	+---------------------------------+-----------------------------------------------------------+
	//	|                                 |                                                           |
	//	|              VALUE              |                          MEANING                          |
	//	|                                 |                                                           |
	//	+---------------------------------+-----------------------------------------------------------+
	//	+---------------------------------+-----------------------------------------------------------+
	//	| APPSTATUS_STOPPED 0x00000000    | The application is defined but is not currently running.  |
	//	+---------------------------------+-----------------------------------------------------------+
	//	| APPSTATUS_RUNNING 0x00000001    | The application is defined and is currently running.      |
	//	+---------------------------------+-----------------------------------------------------------+
	//	| APPSTATUS_NOTDEFINED 0x00000002 | No application is defined at the specified metabase path. |
	//	+---------------------------------+-----------------------------------------------------------+
	AppStatus uint32 `idl:"name:pdwAppStatus" json:"app_status"`
	// Return: The AppGetStatus return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AppGetStatusResponse) xxx_ToOp(ctx context.Context) *xxx_AppGetStatusOperation {
	if o == nil {
		return &xxx_AppGetStatusOperation{}
	}
	return &xxx_AppGetStatusOperation{
		That:      o.That,
		AppStatus: o.AppStatus,
		Return:    o.Return,
	}
}

func (o *AppGetStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_AppGetStatusOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AppStatus = op.AppStatus
	o.Return = op.Return
}
func (o *AppGetStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AppGetStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppGetStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AppDeleteRecoverableOperation structure represents the AppDeleteRecoverable operation
type xxx_AppDeleteRecoverableOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path      string         `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	Recursive bool           `idl:"name:fRecursive" json:"recursive"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AppDeleteRecoverableOperation) OpNum() int { return 7 }

func (o *xxx_AppDeleteRecoverableOperation) OpName() string {
	return "/IWamAdmin/v0/AppDeleteRecoverable"
}

func (o *xxx_AppDeleteRecoverableOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppDeleteRecoverableOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_szMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_szMDPath); err != nil {
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
	// fRecursive {in} (1:{alias=BOOL}(int32))
	{
		if !o.Recursive {
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

func (o *xxx_AppDeleteRecoverableOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_szMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_szMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_szMDPath, _ptr_szMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fRecursive {in} (1:{alias=BOOL}(int32))
	{
		var _bRecursive int32
		if err := w.ReadData(&_bRecursive); err != nil {
			return err
		}
		o.Recursive = _bRecursive != 0
	}
	return nil
}

func (o *xxx_AppDeleteRecoverableOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppDeleteRecoverableOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AppDeleteRecoverableOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AppDeleteRecoverableRequest structure represents the AppDeleteRecoverable operation request
type AppDeleteRecoverableRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szMDPath:  A pointer to a Unicode string that contains the metabase path of the
	// application.
	Path string `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	// fRecursive:  A flag indicating whether application definitions are also to be deleted
	// from all subkeys (TRUE) or just from the application at this key (FALSE).
	Recursive bool `idl:"name:fRecursive" json:"recursive"`
}

func (o *AppDeleteRecoverableRequest) xxx_ToOp(ctx context.Context) *xxx_AppDeleteRecoverableOperation {
	if o == nil {
		return &xxx_AppDeleteRecoverableOperation{}
	}
	return &xxx_AppDeleteRecoverableOperation{
		This:      o.This,
		Path:      o.Path,
		Recursive: o.Recursive,
	}
}

func (o *AppDeleteRecoverableRequest) xxx_FromOp(ctx context.Context, op *xxx_AppDeleteRecoverableOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.Recursive = op.Recursive
}
func (o *AppDeleteRecoverableRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AppDeleteRecoverableRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppDeleteRecoverableOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AppDeleteRecoverableResponse structure represents the AppDeleteRecoverable operation response
type AppDeleteRecoverableResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AppDeleteRecoverable return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AppDeleteRecoverableResponse) xxx_ToOp(ctx context.Context) *xxx_AppDeleteRecoverableOperation {
	if o == nil {
		return &xxx_AppDeleteRecoverableOperation{}
	}
	return &xxx_AppDeleteRecoverableOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AppDeleteRecoverableResponse) xxx_FromOp(ctx context.Context, op *xxx_AppDeleteRecoverableOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AppDeleteRecoverableResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AppDeleteRecoverableResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppDeleteRecoverableOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AppRecoverOperation structure represents the AppRecover operation
type xxx_AppRecoverOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path      string         `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	Recursive bool           `idl:"name:fRecursive" json:"recursive"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AppRecoverOperation) OpNum() int { return 8 }

func (o *xxx_AppRecoverOperation) OpName() string { return "/IWamAdmin/v0/AppRecover" }

func (o *xxx_AppRecoverOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppRecoverOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_szMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_szMDPath); err != nil {
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
	// fRecursive {in} (1:{alias=BOOL}(int32))
	{
		if !o.Recursive {
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

func (o *xxx_AppRecoverOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_szMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_szMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_szMDPath, _ptr_szMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fRecursive {in} (1:{alias=BOOL}(int32))
	{
		var _bRecursive int32
		if err := w.ReadData(&_bRecursive); err != nil {
			return err
		}
		o.Recursive = _bRecursive != 0
	}
	return nil
}

func (o *xxx_AppRecoverOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppRecoverOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AppRecoverOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AppRecoverRequest structure represents the AppRecover operation request
type AppRecoverRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szMDPath: A pointer to a Unicode string that contains the metabase path of the application.
	Path string `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	// fRecursive:  A flag indicating whether application definitions are also to be recovered
	// from all subkeys (TRUE) or just from the application at this key (FALSE).
	Recursive bool `idl:"name:fRecursive" json:"recursive"`
}

func (o *AppRecoverRequest) xxx_ToOp(ctx context.Context) *xxx_AppRecoverOperation {
	if o == nil {
		return &xxx_AppRecoverOperation{}
	}
	return &xxx_AppRecoverOperation{
		This:      o.This,
		Path:      o.Path,
		Recursive: o.Recursive,
	}
}

func (o *AppRecoverRequest) xxx_FromOp(ctx context.Context, op *xxx_AppRecoverOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.Recursive = op.Recursive
}
func (o *AppRecoverRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AppRecoverRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppRecoverOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AppRecoverResponse structure represents the AppRecover operation response
type AppRecoverResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AppRecover return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AppRecoverResponse) xxx_ToOp(ctx context.Context) *xxx_AppRecoverOperation {
	if o == nil {
		return &xxx_AppRecoverOperation{}
	}
	return &xxx_AppRecoverOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AppRecoverResponse) xxx_FromOp(ctx context.Context, op *xxx_AppRecoverOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AppRecoverResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AppRecoverResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppRecoverOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
