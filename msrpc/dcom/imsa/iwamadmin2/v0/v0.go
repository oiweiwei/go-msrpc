package iwamadmin2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iwamadmin "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/iwamadmin/v0"
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
	_ = iwamadmin.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/imsa"
)

var (
	// IWamAdmin2 interface identifier 29822ab8-f302-11d0-9953-00c04fd919c1
	WAMAdmin2IID = &dcom.IID{Data1: 0x29822ab8, Data2: 0xf302, Data3: 0x11d0, Data4: []byte{0x99, 0x53, 0x00, 0xc0, 0x4f, 0xd9, 0x19, 0xc1}}
	// Syntax UUID
	WAMAdmin2SyntaxUUID = &uuid.UUID{TimeLow: 0x29822ab8, TimeMid: 0xf302, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x53, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd9, 0x19, 0xc1}}
	// Syntax ID
	WAMAdmin2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: WAMAdmin2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWamAdmin2 interface.
type WAMAdmin2Client interface {

	// IWamAdmin retrieval method.
	WAMAdmin() iwamadmin.WAMAdminClient

	// The AppCreate2 method creates a new application at the specified metabase path.
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
	AppCreate2(context.Context, *AppCreate2Request, ...dcerpc.CallOption) (*AppCreate2Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) WAMAdmin2Client
}

type xxx_DefaultWAMAdmin2Client struct {
	iwamadmin.WAMAdminClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultWAMAdmin2Client) WAMAdmin() iwamadmin.WAMAdminClient {
	return o.WAMAdminClient
}

func (o *xxx_DefaultWAMAdmin2Client) AppCreate2(ctx context.Context, in *AppCreate2Request, opts ...dcerpc.CallOption) (*AppCreate2Response, error) {
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
	out := &AppCreate2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWAMAdmin2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWAMAdmin2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultWAMAdmin2Client) IPID(ctx context.Context, ipid *dcom.IPID) WAMAdmin2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultWAMAdmin2Client{
		WAMAdminClient: o.WAMAdminClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewWAMAdmin2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WAMAdmin2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WAMAdmin2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iwamadmin.NewWAMAdminClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultWAMAdmin2Client{
		WAMAdminClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_AppCreate2Operation structure represents the AppCreate2 operation
type xxx_AppCreate2Operation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path    string         `idl:"name:szMDPath;string;pointer:unique" json:"path"`
	AppMode uint32         `idl:"name:dwAppMode" json:"app_mode"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AppCreate2Operation) OpNum() int { return 9 }

func (o *xxx_AppCreate2Operation) OpName() string { return "/IWamAdmin2/v0/AppCreate2" }

func (o *xxx_AppCreate2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppCreate2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_AppCreate2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_AppCreate2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AppCreate2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AppCreate2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AppCreate2Request structure represents the AppCreate2 operation request
type AppCreate2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szMDPath:  A pointer to a Unicode string that contains the metabase path of the
	// application.
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
}

func (o *AppCreate2Request) xxx_ToOp(ctx context.Context) *xxx_AppCreate2Operation {
	if o == nil {
		return &xxx_AppCreate2Operation{}
	}
	return &xxx_AppCreate2Operation{
		This:    o.This,
		Path:    o.Path,
		AppMode: o.AppMode,
	}
}

func (o *AppCreate2Request) xxx_FromOp(ctx context.Context, op *xxx_AppCreate2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.AppMode = op.AppMode
}
func (o *AppCreate2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AppCreate2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppCreate2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AppCreate2Response structure represents the AppCreate2 operation response
type AppCreate2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AppCreate2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AppCreate2Response) xxx_ToOp(ctx context.Context) *xxx_AppCreate2Operation {
	if o == nil {
		return &xxx_AppCreate2Operation{}
	}
	return &xxx_AppCreate2Operation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AppCreate2Response) xxx_FromOp(ctx context.Context, op *xxx_AppCreate2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AppCreate2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AppCreate2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AppCreate2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
