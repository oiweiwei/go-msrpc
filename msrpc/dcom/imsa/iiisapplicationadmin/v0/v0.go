package iiisapplicationadmin

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
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/imsa"
)

var (
	// IIISApplicationAdmin interface identifier 7c4e1804-e342-483d-a43e-a850cfcc8d18
	IISApplicationAdminIID = &dcom.IID{Data1: 0x7c4e1804, Data2: 0xe342, Data3: 0x483d, Data4: []byte{0xa4, 0x3e, 0xa8, 0x50, 0xcf, 0xcc, 0x8d, 0x18}}
	// Syntax UUID
	IISApplicationAdminSyntaxUUID = &uuid.UUID{TimeLow: 0x7c4e1804, TimeMid: 0xe342, TimeHiAndVersion: 0x483d, ClockSeqHiAndReserved: 0xa4, ClockSeqLow: 0x3e, Node: [6]uint8{0xa8, 0x50, 0xcf, 0xcc, 0x8d, 0x18}}
	// Syntax ID
	IISApplicationAdminSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: IISApplicationAdminSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IIISApplicationAdmin interface.
type IISApplicationAdminClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The CreateApplication method creates a new application at the specified metabase
	// path.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+----------------------------+------------------------------------+
	//	|           RETURN           |                                    |
	//	|         VALUE/CODE         |            DESCRIPTION             |
	//	|                            |                                    |
	//	+----------------------------+------------------------------------+
	//	+----------------------------+------------------------------------+
	//	| 0x00000000 S_OK            | The call was successful.           |
	//	+----------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG    | One or more arguments are invalid. |
	//	+----------------------------+------------------------------------+
	//	| 0x80070490 ERROR_NOT_FOUND | Element not found.                 |
	//	+----------------------------+------------------------------------+
	//
	// The opnum field value for this method is 3.
	CreateApplication(context.Context, *CreateApplicationRequest, ...dcerpc.CallOption) (*CreateApplicationResponse, error)

	// The DeleteApplication method deletes the application from the specified metabase
	// path.
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
	DeleteApplication(context.Context, *DeleteApplicationRequest, ...dcerpc.CallOption) (*DeleteApplicationResponse, error)

	// The CreateApplicationPool method creates a new application pool.
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
	CreateApplicationPool(context.Context, *CreateApplicationPoolRequest, ...dcerpc.CallOption) (*CreateApplicationPoolResponse, error)

	// The DeleteApplicationPool method deletes an application pool.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+----------------------------+----------------------------------------------------------------------------+
	//	|           RETURN           |                                                                            |
	//	|         VALUE/CODE         |                                DESCRIPTION                                 |
	//	|                            |                                                                            |
	//	+----------------------------+----------------------------------------------------------------------------+
	//	+----------------------------+----------------------------------------------------------------------------+
	//	| 0x00000000 S_OK            | The call was successful.                                                   |
	//	+----------------------------+----------------------------------------------------------------------------+
	//	| 0x80070490 ERROR_NOT_FOUND | Element not found.                                                         |
	//	+----------------------------+----------------------------------------------------------------------------+
	//	| 0x800710D3 ERROR_NOT_EMPTY | The library, drive, or media pool must be empty to perform this operation. |
	//	+----------------------------+----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 6.
	DeleteApplicationPool(context.Context, *DeleteApplicationPoolRequest, ...dcerpc.CallOption) (*DeleteApplicationPoolResponse, error)

	// The EnumerateApplicationsInPool method returns the metabase paths for the applications
	// associated with the application pool.
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
	// The opnum field value for this method is 7.
	EnumerateApplicationsInPool(context.Context, *EnumerateApplicationsInPoolRequest, ...dcerpc.CallOption) (*EnumerateApplicationsInPoolResponse, error)

	// The RecycleApplicationPool method restarts an application pool.
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
	RecycleApplicationPool(context.Context, *RecycleApplicationPoolRequest, ...dcerpc.CallOption) (*RecycleApplicationPoolResponse, error)

	// The GetProcessMode method retrieves the application execution mode for the IIS server.
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
	// The opnum field value for this method is 9.
	GetProcessMode(context.Context, *GetProcessModeRequest, ...dcerpc.CallOption) (*GetProcessModeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) IISApplicationAdminClient
}

type xxx_DefaultIISApplicationAdminClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultIISApplicationAdminClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultIISApplicationAdminClient) CreateApplication(ctx context.Context, in *CreateApplicationRequest, opts ...dcerpc.CallOption) (*CreateApplicationResponse, error) {
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
	out := &CreateApplicationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISApplicationAdminClient) DeleteApplication(ctx context.Context, in *DeleteApplicationRequest, opts ...dcerpc.CallOption) (*DeleteApplicationResponse, error) {
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
	out := &DeleteApplicationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISApplicationAdminClient) CreateApplicationPool(ctx context.Context, in *CreateApplicationPoolRequest, opts ...dcerpc.CallOption) (*CreateApplicationPoolResponse, error) {
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
	out := &CreateApplicationPoolResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISApplicationAdminClient) DeleteApplicationPool(ctx context.Context, in *DeleteApplicationPoolRequest, opts ...dcerpc.CallOption) (*DeleteApplicationPoolResponse, error) {
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
	out := &DeleteApplicationPoolResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISApplicationAdminClient) EnumerateApplicationsInPool(ctx context.Context, in *EnumerateApplicationsInPoolRequest, opts ...dcerpc.CallOption) (*EnumerateApplicationsInPoolResponse, error) {
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
	out := &EnumerateApplicationsInPoolResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISApplicationAdminClient) RecycleApplicationPool(ctx context.Context, in *RecycleApplicationPoolRequest, opts ...dcerpc.CallOption) (*RecycleApplicationPoolResponse, error) {
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
	out := &RecycleApplicationPoolResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISApplicationAdminClient) GetProcessMode(ctx context.Context, in *GetProcessModeRequest, opts ...dcerpc.CallOption) (*GetProcessModeResponse, error) {
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
	out := &GetProcessModeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISApplicationAdminClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIISApplicationAdminClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultIISApplicationAdminClient) IPID(ctx context.Context, ipid *dcom.IPID) IISApplicationAdminClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultIISApplicationAdminClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewIISApplicationAdminClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IISApplicationAdminClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IISApplicationAdminSyntaxV0_0))...)
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
	return &xxx_DefaultIISApplicationAdminClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CreateApplicationOperation structure represents the CreateApplication operation
type xxx_CreateApplicationOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path       string         `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	AppMode    uint32         `idl:"name:dwAppMode" json:"app_mode"`
	AppPoolID  string         `idl:"name:szAppPoolId;string;pointer:unique" json:"app_pool_id"`
	CreatePool bool           `idl:"name:fCreatePool" json:"create_pool"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateApplicationOperation) OpNum() int { return 3 }

func (o *xxx_CreateApplicationOperation) OpName() string {
	return "/IIISApplicationAdmin/v0/CreateApplication"
}

func (o *xxx_CreateApplicationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateApplicationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwAppMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AppMode); err != nil {
			return err
		}
	}
	// szAppPoolId {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.AppPoolID != "" {
			_ptr_szAppPoolId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.AppPoolID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.AppPoolID, _ptr_szAppPoolId); err != nil {
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
	// fCreatePool {in} (1:{alias=BOOL}(int32))
	{
		if !o.CreatePool {
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

func (o *xxx_CreateApplicationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwAppMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AppMode); err != nil {
			return err
		}
	}
	// szAppPoolId {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_szAppPoolId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.AppPoolID); err != nil {
				return err
			}
			return nil
		})
		_s_szAppPoolId := func(ptr interface{}) { o.AppPoolID = *ptr.(*string) }
		if err := w.ReadPointer(&o.AppPoolID, _s_szAppPoolId, _ptr_szAppPoolId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fCreatePool {in} (1:{alias=BOOL}(int32))
	{
		var _bCreatePool int32
		if err := w.ReadData(&_bCreatePool); err != nil {
			return err
		}
		o.CreatePool = _bCreatePool != 0
	}
	return nil
}

func (o *xxx_CreateApplicationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateApplicationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateApplicationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateApplicationRequest structure represents the CreateApplication operation request
type CreateApplicationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szMDPath: A pointer to a Unicode string that contains the metabase path of the application.
	Path string `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	// dwAppMode:  An unsigned 32-bit integer value indicating the process where the application
	// will run. This parameter MUST be set to one of the following values.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| eAppRunInProc 0x00000000               | The application runs in the IIS parent process.                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| eAppRunOutProcIsolated 0x00000001      | The application runs in its own process.                                         |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| eAppRunOutProcInDefaultPool 0x00000002 | The application runs in a shared process with other applications outside of the  |
	//	|                                        | IIS parent process.                                                              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	AppMode uint32 `idl:"name:dwAppMode" json:"app_mode"`
	// szAppPoolId: A pointer to a Unicode string that specifies the application pool name
	// with which to associate the new application.
	AppPoolID string `idl:"name:szAppPoolId;string;pointer:unique" json:"app_pool_id"`
	// fCreatePool: A flag indicating whether to create a new application pool if the pool
	// specified by the szAppPoolId parameter does not exist (TRUE) or to use an existing
	// application pool (FALSE).
	CreatePool bool `idl:"name:fCreatePool" json:"create_pool"`
}

func (o *CreateApplicationRequest) xxx_ToOp(ctx context.Context) *xxx_CreateApplicationOperation {
	if o == nil {
		return &xxx_CreateApplicationOperation{}
	}
	return &xxx_CreateApplicationOperation{
		This:       o.This,
		Path:       o.Path,
		AppMode:    o.AppMode,
		AppPoolID:  o.AppPoolID,
		CreatePool: o.CreatePool,
	}
}

func (o *CreateApplicationRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateApplicationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.AppMode = op.AppMode
	o.AppPoolID = op.AppPoolID
	o.CreatePool = op.CreatePool
}
func (o *CreateApplicationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateApplicationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateApplicationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateApplicationResponse structure represents the CreateApplication operation response
type CreateApplicationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateApplication return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateApplicationResponse) xxx_ToOp(ctx context.Context) *xxx_CreateApplicationOperation {
	if o == nil {
		return &xxx_CreateApplicationOperation{}
	}
	return &xxx_CreateApplicationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CreateApplicationResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateApplicationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateApplicationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateApplicationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateApplicationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteApplicationOperation structure represents the DeleteApplication operation
type xxx_DeleteApplicationOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path      string         `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	Recursive bool           `idl:"name:fRecursive" json:"recursive"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteApplicationOperation) OpNum() int { return 4 }

func (o *xxx_DeleteApplicationOperation) OpName() string {
	return "/IIISApplicationAdmin/v0/DeleteApplication"
}

func (o *xxx_DeleteApplicationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteApplicationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteApplicationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_DeleteApplicationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteApplicationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteApplicationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteApplicationRequest structure represents the DeleteApplication operation request
type DeleteApplicationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szMDPath:  A pointer to a Unicode string that contains the metabase path of the
	// application.
	Path string `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	// fRecursive: A flag indicating whether application definitions are also to be deleted
	// from all subkeys (TRUE) or just from the application at this key (FALSE).
	Recursive bool `idl:"name:fRecursive" json:"recursive"`
}

func (o *DeleteApplicationRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteApplicationOperation {
	if o == nil {
		return &xxx_DeleteApplicationOperation{}
	}
	return &xxx_DeleteApplicationOperation{
		This:      o.This,
		Path:      o.Path,
		Recursive: o.Recursive,
	}
}

func (o *DeleteApplicationRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteApplicationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.Recursive = op.Recursive
}
func (o *DeleteApplicationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteApplicationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteApplicationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteApplicationResponse structure represents the DeleteApplication operation response
type DeleteApplicationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteApplication return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteApplicationResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteApplicationOperation {
	if o == nil {
		return &xxx_DeleteApplicationOperation{}
	}
	return &xxx_DeleteApplicationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeleteApplicationResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteApplicationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteApplicationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteApplicationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteApplicationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateApplicationPoolOperation structure represents the CreateApplicationPool operation
type xxx_CreateApplicationPoolOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pool   string         `idl:"name:szPool;string;pointer:unique" json:"pool"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateApplicationPoolOperation) OpNum() int { return 5 }

func (o *xxx_CreateApplicationPoolOperation) OpName() string {
	return "/IIISApplicationAdmin/v0/CreateApplicationPool"
}

func (o *xxx_CreateApplicationPoolOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateApplicationPoolOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szPool {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Pool != "" {
			_ptr_szPool := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Pool); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Pool, _ptr_szPool); err != nil {
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

func (o *xxx_CreateApplicationPoolOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szPool {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_szPool := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Pool); err != nil {
				return err
			}
			return nil
		})
		_s_szPool := func(ptr interface{}) { o.Pool = *ptr.(*string) }
		if err := w.ReadPointer(&o.Pool, _s_szPool, _ptr_szPool); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateApplicationPoolOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateApplicationPoolOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateApplicationPoolOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateApplicationPoolRequest structure represents the CreateApplicationPool operation request
type CreateApplicationPoolRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szPool: A pointer to a Unicode string that contains the name of the new application
	// pool.
	Pool string `idl:"name:szPool;string;pointer:unique" json:"pool"`
}

func (o *CreateApplicationPoolRequest) xxx_ToOp(ctx context.Context) *xxx_CreateApplicationPoolOperation {
	if o == nil {
		return &xxx_CreateApplicationPoolOperation{}
	}
	return &xxx_CreateApplicationPoolOperation{
		This: o.This,
		Pool: o.Pool,
	}
}

func (o *CreateApplicationPoolRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateApplicationPoolOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Pool = op.Pool
}
func (o *CreateApplicationPoolRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateApplicationPoolRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateApplicationPoolOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateApplicationPoolResponse structure represents the CreateApplicationPool operation response
type CreateApplicationPoolResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateApplicationPool return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateApplicationPoolResponse) xxx_ToOp(ctx context.Context) *xxx_CreateApplicationPoolOperation {
	if o == nil {
		return &xxx_CreateApplicationPoolOperation{}
	}
	return &xxx_CreateApplicationPoolOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CreateApplicationPoolResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateApplicationPoolOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateApplicationPoolResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateApplicationPoolResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateApplicationPoolOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteApplicationPoolOperation structure represents the DeleteApplicationPool operation
type xxx_DeleteApplicationPoolOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pool   string         `idl:"name:szPool;string;pointer:unique" json:"pool"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteApplicationPoolOperation) OpNum() int { return 6 }

func (o *xxx_DeleteApplicationPoolOperation) OpName() string {
	return "/IIISApplicationAdmin/v0/DeleteApplicationPool"
}

func (o *xxx_DeleteApplicationPoolOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteApplicationPoolOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szPool {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Pool != "" {
			_ptr_szPool := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Pool); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Pool, _ptr_szPool); err != nil {
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

func (o *xxx_DeleteApplicationPoolOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szPool {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_szPool := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Pool); err != nil {
				return err
			}
			return nil
		})
		_s_szPool := func(ptr interface{}) { o.Pool = *ptr.(*string) }
		if err := w.ReadPointer(&o.Pool, _s_szPool, _ptr_szPool); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteApplicationPoolOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteApplicationPoolOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteApplicationPoolOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteApplicationPoolRequest structure represents the DeleteApplicationPool operation request
type DeleteApplicationPoolRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szPool: A pointer to a Unicode string that contains the name of the application pool
	// to delete.
	Pool string `idl:"name:szPool;string;pointer:unique" json:"pool"`
}

func (o *DeleteApplicationPoolRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteApplicationPoolOperation {
	if o == nil {
		return &xxx_DeleteApplicationPoolOperation{}
	}
	return &xxx_DeleteApplicationPoolOperation{
		This: o.This,
		Pool: o.Pool,
	}
}

func (o *DeleteApplicationPoolRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteApplicationPoolOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Pool = op.Pool
}
func (o *DeleteApplicationPoolRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteApplicationPoolRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteApplicationPoolOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteApplicationPoolResponse structure represents the DeleteApplicationPool operation response
type DeleteApplicationPoolResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteApplicationPool return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteApplicationPoolResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteApplicationPoolOperation {
	if o == nil {
		return &xxx_DeleteApplicationPoolOperation{}
	}
	return &xxx_DeleteApplicationPoolOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeleteApplicationPoolResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteApplicationPoolOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteApplicationPoolResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteApplicationPoolResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteApplicationPoolOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateApplicationsInPoolOperation structure represents the EnumerateApplicationsInPool operation
type xxx_EnumerateApplicationsInPoolOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pool   string         `idl:"name:szPool;string;pointer:unique" json:"pool"`
	Buffer *oaut.String   `idl:"name:bstrBuffer" json:"buffer"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateApplicationsInPoolOperation) OpNum() int { return 7 }

func (o *xxx_EnumerateApplicationsInPoolOperation) OpName() string {
	return "/IIISApplicationAdmin/v0/EnumerateApplicationsInPool"
}

func (o *xxx_EnumerateApplicationsInPoolOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateApplicationsInPoolOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szPool {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Pool != "" {
			_ptr_szPool := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Pool); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Pool, _ptr_szPool); err != nil {
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

func (o *xxx_EnumerateApplicationsInPoolOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szPool {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_szPool := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Pool); err != nil {
				return err
			}
			return nil
		})
		_s_szPool := func(ptr interface{}) { o.Pool = *ptr.(*string) }
		if err := w.ReadPointer(&o.Pool, _s_szPool, _ptr_szPool); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateApplicationsInPoolOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateApplicationsInPoolOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// bstrBuffer {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Buffer != nil {
			_ptr_bstrBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil {
					if err := o.Buffer.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_bstrBuffer); err != nil {
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

func (o *xxx_EnumerateApplicationsInPoolOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// bstrBuffer {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Buffer == nil {
				o.Buffer = &oaut.String{}
			}
			if err := o.Buffer.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrBuffer := func(ptr interface{}) { o.Buffer = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Buffer, _s_bstrBuffer, _ptr_bstrBuffer); err != nil {
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

// EnumerateApplicationsInPoolRequest structure represents the EnumerateApplicationsInPool operation request
type EnumerateApplicationsInPoolRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szPool: A pointer to a Unicode string that contains the name of the application pool
	// to enumerate.
	Pool string `idl:"name:szPool;string;pointer:unique" json:"pool"`
}

func (o *EnumerateApplicationsInPoolRequest) xxx_ToOp(ctx context.Context) *xxx_EnumerateApplicationsInPoolOperation {
	if o == nil {
		return &xxx_EnumerateApplicationsInPoolOperation{}
	}
	return &xxx_EnumerateApplicationsInPoolOperation{
		This: o.This,
		Pool: o.Pool,
	}
}

func (o *EnumerateApplicationsInPoolRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateApplicationsInPoolOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Pool = op.Pool
}
func (o *EnumerateApplicationsInPoolRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateApplicationsInPoolRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateApplicationsInPoolOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateApplicationsInPoolResponse structure represents the EnumerateApplicationsInPool operation response
type EnumerateApplicationsInPoolResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// bstrBuffer: A pointer to a BSTR that receives the application metabase paths. The
	// BSTR contains a sequence of contiguous null-terminated strings. The buffer is terminated
	// by another null character. The server allocates storage, and the client is responsible
	// for freeing the storage with SysFreeString; see [MS-OAUT].
	Buffer *oaut.String `idl:"name:bstrBuffer" json:"buffer"`
	// Return: The EnumerateApplicationsInPool return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateApplicationsInPoolResponse) xxx_ToOp(ctx context.Context) *xxx_EnumerateApplicationsInPoolOperation {
	if o == nil {
		return &xxx_EnumerateApplicationsInPoolOperation{}
	}
	return &xxx_EnumerateApplicationsInPoolOperation{
		That:   o.That,
		Buffer: o.Buffer,
		Return: o.Return,
	}
}

func (o *EnumerateApplicationsInPoolResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateApplicationsInPoolOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *EnumerateApplicationsInPoolResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateApplicationsInPoolResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateApplicationsInPoolOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RecycleApplicationPoolOperation structure represents the RecycleApplicationPool operation
type xxx_RecycleApplicationPoolOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pool   string         `idl:"name:szPool;string;pointer:unique" json:"pool"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RecycleApplicationPoolOperation) OpNum() int { return 8 }

func (o *xxx_RecycleApplicationPoolOperation) OpName() string {
	return "/IIISApplicationAdmin/v0/RecycleApplicationPool"
}

func (o *xxx_RecycleApplicationPoolOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecycleApplicationPoolOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szPool {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Pool != "" {
			_ptr_szPool := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Pool); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Pool, _ptr_szPool); err != nil {
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

func (o *xxx_RecycleApplicationPoolOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szPool {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_szPool := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Pool); err != nil {
				return err
			}
			return nil
		})
		_s_szPool := func(ptr interface{}) { o.Pool = *ptr.(*string) }
		if err := w.ReadPointer(&o.Pool, _s_szPool, _ptr_szPool); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecycleApplicationPoolOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecycleApplicationPoolOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RecycleApplicationPoolOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RecycleApplicationPoolRequest structure represents the RecycleApplicationPool operation request
type RecycleApplicationPoolRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szPool: A pointer to a Unicode string that contains the name of the application pool
	// to restart.
	Pool string `idl:"name:szPool;string;pointer:unique" json:"pool"`
}

func (o *RecycleApplicationPoolRequest) xxx_ToOp(ctx context.Context) *xxx_RecycleApplicationPoolOperation {
	if o == nil {
		return &xxx_RecycleApplicationPoolOperation{}
	}
	return &xxx_RecycleApplicationPoolOperation{
		This: o.This,
		Pool: o.Pool,
	}
}

func (o *RecycleApplicationPoolRequest) xxx_FromOp(ctx context.Context, op *xxx_RecycleApplicationPoolOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Pool = op.Pool
}
func (o *RecycleApplicationPoolRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RecycleApplicationPoolRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RecycleApplicationPoolOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RecycleApplicationPoolResponse structure represents the RecycleApplicationPool operation response
type RecycleApplicationPoolResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RecycleApplicationPool return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RecycleApplicationPoolResponse) xxx_ToOp(ctx context.Context) *xxx_RecycleApplicationPoolOperation {
	if o == nil {
		return &xxx_RecycleApplicationPoolOperation{}
	}
	return &xxx_RecycleApplicationPoolOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RecycleApplicationPoolResponse) xxx_FromOp(ctx context.Context, op *xxx_RecycleApplicationPoolOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RecycleApplicationPoolResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RecycleApplicationPoolResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RecycleApplicationPoolOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetProcessModeOperation structure represents the GetProcessMode operation
type xxx_GetProcessModeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Mode   uint32         `idl:"name:pdwMode" json:"mode"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetProcessModeOperation) OpNum() int { return 9 }

func (o *xxx_GetProcessModeOperation) OpName() string {
	return "/IIISApplicationAdmin/v0/GetProcessMode"
}

func (o *xxx_GetProcessModeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProcessModeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetProcessModeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetProcessModeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProcessModeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwMode {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Mode); err != nil {
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

func (o *xxx_GetProcessModeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwMode {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Mode); err != nil {
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

// GetProcessModeRequest structure represents the GetProcessMode operation request
type GetProcessModeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetProcessModeRequest) xxx_ToOp(ctx context.Context) *xxx_GetProcessModeOperation {
	if o == nil {
		return &xxx_GetProcessModeOperation{}
	}
	return &xxx_GetProcessModeOperation{
		This: o.This,
	}
}

func (o *GetProcessModeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetProcessModeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetProcessModeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetProcessModeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProcessModeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetProcessModeResponse structure represents the GetProcessMode operation response
type GetProcessModeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwMode: A pointer to an unsigned 32-bit integer that receives the server's application
	// execution mode. This parameter MUST be set to one of the following values.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The server is hosting applications in application pools.                         |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 | The server is hosting applications in the IIS server process and child           |
	//	|            | processes.                                                                       |
	//	+------------+----------------------------------------------------------------------------------+
	Mode uint32 `idl:"name:pdwMode" json:"mode"`
	// Return: The GetProcessMode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetProcessModeResponse) xxx_ToOp(ctx context.Context) *xxx_GetProcessModeOperation {
	if o == nil {
		return &xxx_GetProcessModeOperation{}
	}
	return &xxx_GetProcessModeOperation{
		That:   o.That,
		Mode:   o.Mode,
		Return: o.Return,
	}
}

func (o *GetProcessModeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetProcessModeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Mode = op.Mode
	o.Return = op.Return
}
func (o *GetProcessModeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetProcessModeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProcessModeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
