package imsadminbase2w

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	imsadminbasew "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/imsadminbasew/v0"
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
	_ = dcom.GoPackage
	_ = imsadminbasew.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/imsa"
)

var (
	// IMSAdminBase2W interface identifier 8298d101-f992-43b7-8eca-5052d885b995
	IMSAdminBase2WIID = &dcom.IID{Data1: 0x8298d101, Data2: 0xf992, Data3: 0x43b7, Data4: []byte{0x8e, 0xca, 0x50, 0x52, 0xd8, 0x85, 0xb9, 0x95}}
	// Syntax UUID
	IMSAdminBase2WSyntaxUUID = &uuid.UUID{TimeLow: 0x8298d101, TimeMid: 0xf992, TimeHiAndVersion: 0x43b7, ClockSeqHiAndReserved: 0x8e, ClockSeqLow: 0xca, Node: [6]uint8{0x50, 0x52, 0xd8, 0x85, 0xb9, 0x95}}
	// Syntax ID
	IMSAdminBase2WSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: IMSAdminBase2WSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSAdminBase2W interface.
type IMSAdminBase2WClient interface {

	// IMSAdminBaseW retrieval method.
	IMSAdminBaseW() imsadminbasew.IMSAdminBaseWClient

	// The BackupWithPasswd method backs up the metabase using a supplied password to encrypt
	// all secure data.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.           |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 34.
	BackupWithPassword(context.Context, *BackupWithPasswordRequest, ...dcerpc.CallOption) (*BackupWithPasswordResponse, error)

	// The RestoreWithPasswd method restores the metabase from a backup, using a supplied
	// password to decrypt the secure data.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.                                               |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007052B ERROR_WRONG_PASSWORD | Unable to update the password. The value provided as the current password is     |
	//	|                                 | incorrect.                                                                       |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 35.
	RestoreWithPassword(context.Context, *RestoreWithPasswordRequest, ...dcerpc.CallOption) (*RestoreWithPasswordResponse, error)

	// The Export method exports a section of the metabase to a file.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
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
	//	| 0x80070032 ERROR_NOT_SUPPORTED  | The request is not supported.              |
	//	+---------------------------------+--------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 36.
	Export(context.Context, *ExportRequest, ...dcerpc.CallOption) (*ExportResponse, error)

	// The Import method imports metabase data from an exported file into the metabase.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
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
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the file specified. |
	//	+---------------------------------+--------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 37.
	Import(context.Context, *ImportRequest, ...dcerpc.CallOption) (*ImportResponse, error)

	// The RestoreHistory method restores a metabase history entry for a specific history
	// version.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------------------------+---------------------------------------------------------------+
	//	|               RETURN                |                                                               |
	//	|             VALUE/CODE              |                          DESCRIPTION                          |
	//	|                                     |                                                               |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                                      |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND     | The system cannot find the file specified.                    |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND     | The system cannot find the path specified.                    |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x8007000E E_OUTOFMEMORY            | Ran out of memory.                                            |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY  | Not enough storage is available to process this command.      |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG             | One or more arguments are invalid.                            |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x80070013 ERROR_INVALID_DATA       | One or more arguments are invalid.                            |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x800703EC ERROR_INVALID_FLAGS      | Invalid flags were passed.                                    |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED      | Access is denied.                                             |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x800CC802 MD_ERROR_INVALID_VERSION | The version specified in metadata storage was not recognized. |
	//	+-------------------------------------+---------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 38.
	RestoreHistory(context.Context, *RestoreHistoryRequest, ...dcerpc.CallOption) (*RestoreHistoryResponse, error)

	// The EnumHistory method returns an enumerated history entry with a supplied index.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_PATH_NOT_FOUND      | The system cannot find the file specified.                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY   | Not enough storage is available to process this command.                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000012 ERROR_NO_MORE_ITEMS       | There are no more history versions.                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007A ERROR_INSUFFICIENT_BUFFER | The data area passed to a system call is too small. In this case the location    |
	//	|                                      | string does not have enough space to return the path to the history location.    |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED       | Access is denied.                                                                |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | One or more arguments are invalid.                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 39.
	EnumHistory(context.Context, *EnumHistoryRequest, ...dcerpc.CallOption) (*EnumHistoryResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) IMSAdminBase2WClient
}

type xxx_DefaultIMSAdminBase2WClient struct {
	imsadminbasew.IMSAdminBaseWClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultIMSAdminBase2WClient) IMSAdminBaseW() imsadminbasew.IMSAdminBaseWClient {
	return o.IMSAdminBaseWClient
}

func (o *xxx_DefaultIMSAdminBase2WClient) BackupWithPassword(ctx context.Context, in *BackupWithPasswordRequest, opts ...dcerpc.CallOption) (*BackupWithPasswordResponse, error) {
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
	out := &BackupWithPasswordResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBase2WClient) RestoreWithPassword(ctx context.Context, in *RestoreWithPasswordRequest, opts ...dcerpc.CallOption) (*RestoreWithPasswordResponse, error) {
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
	out := &RestoreWithPasswordResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBase2WClient) Export(ctx context.Context, in *ExportRequest, opts ...dcerpc.CallOption) (*ExportResponse, error) {
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
	out := &ExportResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBase2WClient) Import(ctx context.Context, in *ImportRequest, opts ...dcerpc.CallOption) (*ImportResponse, error) {
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
	out := &ImportResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBase2WClient) RestoreHistory(ctx context.Context, in *RestoreHistoryRequest, opts ...dcerpc.CallOption) (*RestoreHistoryResponse, error) {
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
	out := &RestoreHistoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBase2WClient) EnumHistory(ctx context.Context, in *EnumHistoryRequest, opts ...dcerpc.CallOption) (*EnumHistoryResponse, error) {
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
	out := &EnumHistoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBase2WClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIMSAdminBase2WClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultIMSAdminBase2WClient) IPID(ctx context.Context, ipid *dcom.IPID) IMSAdminBase2WClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultIMSAdminBase2WClient{
		IMSAdminBaseWClient: o.IMSAdminBaseWClient.IPID(ctx, ipid),
		cc:                  o.cc,
		ipid:                ipid,
	}
}

func NewIMSAdminBase2WClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IMSAdminBase2WClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IMSAdminBase2WSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := imsadminbasew.NewIMSAdminBaseWClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultIMSAdminBase2WClient{
		IMSAdminBaseWClient: base,
		cc:                  cc,
		ipid:                ipid,
	}, nil
}

// xxx_BackupWithPasswordOperation structure represents the BackupWithPasswd operation
type xxx_BackupWithPasswordOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	BackupName string         `idl:"name:pszMDBackupName;string;pointer:unique" json:"backup_name"`
	Version    uint32         `idl:"name:dwMDVersion" json:"version"`
	Flags      uint32         `idl:"name:dwMDFlags" json:"flags"`
	Password   string         `idl:"name:pszPasswd;string;pointer:unique" json:"password"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupWithPasswordOperation) OpNum() int { return 34 }

func (o *xxx_BackupWithPasswordOperation) OpName() string {
	return "/IMSAdminBase2W/v0/BackupWithPasswd"
}

func (o *xxx_BackupWithPasswordOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupWithPasswordOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszMDBackupName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.BackupName != "" {
			_ptr_pszMDBackupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.BackupName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.BackupName, _ptr_pszMDBackupName); err != nil {
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
	// dwMDVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pszPasswd {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Password != "" {
			_ptr_pszPasswd := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_pszPasswd); err != nil {
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

func (o *xxx_BackupWithPasswordOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszMDBackupName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDBackupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.BackupName); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDBackupName := func(ptr interface{}) { o.BackupName = *ptr.(*string) }
		if err := w.ReadPointer(&o.BackupName, _s_pszMDBackupName, _ptr_pszMDBackupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pszPasswd {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszPasswd := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
				return err
			}
			return nil
		})
		_s_pszPasswd := func(ptr interface{}) { o.Password = *ptr.(*string) }
		if err := w.ReadPointer(&o.Password, _s_pszPasswd, _ptr_pszPasswd); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupWithPasswordOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupWithPasswordOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BackupWithPasswordOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// BackupWithPasswordRequest structure represents the BackupWithPasswd operation request
type BackupWithPasswordRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pszMDBackupName: The name of the backup that is being created.
	BackupName string `idl:"name:pszMDBackupName;string;pointer:unique" json:"backup_name"`
	// dwMDVersion: An integer value specifying either the specific version number to be
	// used for the backup or one of the following flag values. If the version number is
	// an explicit version number, it SHOULD be less than MD_BACKUP_MAX_VERSION (9999).
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|                VALUE                 |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MD_BACKUP_HIGHEST_VERSION 0xFFFFFFFE | Use the highest existing backup version for the backup name specified.           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MD_BACKUP_NEXT_VERSION 0xFFFFFFFF    | Use the highest existing backup version number plus one for the backup name      |
	//	|                                      | specified.                                                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Version uint32 `idl:"name:dwMDVersion" json:"version"`
	// dwMDFlags: An integer value containing the bit flags to alter backup functionality.
	// The flags can be zero or one or more of the following values.
	//
	//	+-------+---------+
	//	|       |         |
	//	| VALUE | MEANING |
	//	|       |         |
	//	+-------+---------+
	//	+-------+---------+
	//	|
	//	+-------+---------+
	//	| Back up even if a backup of the same name and version exists in the specified backup location, overwriting it if necessary.           |
	//	+-------+---------+
	//	| Perform a SaveData operation before the backup.           |
	//	+-------+---------+
	Flags uint32 `idl:"name:dwMDFlags" json:"flags"`
	// pszPasswd: A password string used to encrypt the secure properties in the metabase
	// backup. If a password is not supplied, this method functions exactly the same as
	// the Backup method.
	Password string `idl:"name:pszPasswd;string;pointer:unique" json:"password"`
}

func (o *BackupWithPasswordRequest) xxx_ToOp(ctx context.Context) *xxx_BackupWithPasswordOperation {
	if o == nil {
		return &xxx_BackupWithPasswordOperation{}
	}
	return &xxx_BackupWithPasswordOperation{
		This:       o.This,
		BackupName: o.BackupName,
		Version:    o.Version,
		Flags:      o.Flags,
		Password:   o.Password,
	}
}

func (o *BackupWithPasswordRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupWithPasswordOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BackupName = op.BackupName
	o.Version = op.Version
	o.Flags = op.Flags
	o.Password = op.Password
}
func (o *BackupWithPasswordRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BackupWithPasswordRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupWithPasswordOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupWithPasswordResponse structure represents the BackupWithPasswd operation response
type BackupWithPasswordResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The BackupWithPasswd return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupWithPasswordResponse) xxx_ToOp(ctx context.Context) *xxx_BackupWithPasswordOperation {
	if o == nil {
		return &xxx_BackupWithPasswordOperation{}
	}
	return &xxx_BackupWithPasswordOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *BackupWithPasswordResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupWithPasswordOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *BackupWithPasswordResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BackupWithPasswordResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupWithPasswordOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RestoreWithPasswordOperation structure represents the RestoreWithPasswd operation
type xxx_RestoreWithPasswordOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	BackupName string         `idl:"name:pszMDBackupName;string;pointer:unique" json:"backup_name"`
	Version    uint32         `idl:"name:dwMDVersion" json:"version"`
	Flags      uint32         `idl:"name:dwMDFlags" json:"flags"`
	Password   string         `idl:"name:pszPasswd;string;pointer:unique" json:"password"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RestoreWithPasswordOperation) OpNum() int { return 35 }

func (o *xxx_RestoreWithPasswordOperation) OpName() string {
	return "/IMSAdminBase2W/v0/RestoreWithPasswd"
}

func (o *xxx_RestoreWithPasswordOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreWithPasswordOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszMDBackupName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.BackupName != "" {
			_ptr_pszMDBackupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.BackupName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.BackupName, _ptr_pszMDBackupName); err != nil {
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
	// dwMDVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pszPasswd {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Password != "" {
			_ptr_pszPasswd := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_pszPasswd); err != nil {
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

func (o *xxx_RestoreWithPasswordOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszMDBackupName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDBackupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.BackupName); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDBackupName := func(ptr interface{}) { o.BackupName = *ptr.(*string) }
		if err := w.ReadPointer(&o.BackupName, _s_pszMDBackupName, _ptr_pszMDBackupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pszPasswd {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszPasswd := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
				return err
			}
			return nil
		})
		_s_pszPasswd := func(ptr interface{}) { o.Password = *ptr.(*string) }
		if err := w.ReadPointer(&o.Password, _s_pszPasswd, _ptr_pszPasswd); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreWithPasswordOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreWithPasswordOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RestoreWithPasswordOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RestoreWithPasswordRequest structure represents the RestoreWithPasswd operation request
type RestoreWithPasswordRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pszMDBackupName: A pointer to a Unicode string containing the name of the backup
	// to be restored.
	BackupName string `idl:"name:pszMDBackupName;string;pointer:unique" json:"backup_name"`
	// dwMDVersion: An integer value specifying the version number of the backup to be restored,
	// which MUST be less than or equal to MD_BACKUP_MAX_VERSION (9999) or the following
	// constant.
	//
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	|                                      |                                                                                |
	//	|                VALUE                 |                                    MEANING                                     |
	//	|                                      |                                                                                |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| MD_BACKUP_HIGHEST_VERSION 0xFFFFFFFE | Restore from the highest existing backup version in the specified backup name. |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	Version uint32 `idl:"name:dwMDVersion" json:"version"`
	// dwMDFlags: This parameter is reserved and MUST be set to zero.
	Flags uint32 `idl:"name:dwMDFlags" json:"flags"`
	// pszPasswd: A password string used to decrypt the secure properties in the metabase
	// backup. If the password is not correct, an error is returned. If a password is not
	// supplied, this method functions exactly the same as the Restore method.
	Password string `idl:"name:pszPasswd;string;pointer:unique" json:"password"`
}

func (o *RestoreWithPasswordRequest) xxx_ToOp(ctx context.Context) *xxx_RestoreWithPasswordOperation {
	if o == nil {
		return &xxx_RestoreWithPasswordOperation{}
	}
	return &xxx_RestoreWithPasswordOperation{
		This:       o.This,
		BackupName: o.BackupName,
		Version:    o.Version,
		Flags:      o.Flags,
		Password:   o.Password,
	}
}

func (o *RestoreWithPasswordRequest) xxx_FromOp(ctx context.Context, op *xxx_RestoreWithPasswordOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BackupName = op.BackupName
	o.Version = op.Version
	o.Flags = op.Flags
	o.Password = op.Password
}
func (o *RestoreWithPasswordRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RestoreWithPasswordRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreWithPasswordOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RestoreWithPasswordResponse structure represents the RestoreWithPasswd operation response
type RestoreWithPasswordResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RestoreWithPasswd return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RestoreWithPasswordResponse) xxx_ToOp(ctx context.Context) *xxx_RestoreWithPasswordOperation {
	if o == nil {
		return &xxx_RestoreWithPasswordOperation{}
	}
	return &xxx_RestoreWithPasswordOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RestoreWithPasswordResponse) xxx_FromOp(ctx context.Context, op *xxx_RestoreWithPasswordOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RestoreWithPasswordResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RestoreWithPasswordResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreWithPasswordOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExportOperation structure represents the Export operation
type xxx_ExportOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Password   string         `idl:"name:pszPasswd;string;pointer:unique" json:"password"`
	FileName   string         `idl:"name:pszFileName;string;pointer:unique" json:"file_name"`
	SourcePath string         `idl:"name:pszSourcePath;string;pointer:unique" json:"source_path"`
	Flags      uint32         `idl:"name:dwMDFlags" json:"flags"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ExportOperation) OpNum() int { return 36 }

func (o *xxx_ExportOperation) OpName() string { return "/IMSAdminBase2W/v0/Export" }

func (o *xxx_ExportOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszPasswd {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Password != "" {
			_ptr_pszPasswd := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_pszPasswd); err != nil {
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
	// pszFileName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.FileName != "" {
			_ptr_pszFileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FileName, _ptr_pszFileName); err != nil {
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
	// pszSourcePath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.SourcePath != "" {
			_ptr_pszSourcePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SourcePath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SourcePath, _ptr_pszSourcePath); err != nil {
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
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszPasswd {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszPasswd := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
				return err
			}
			return nil
		})
		_s_pszPasswd := func(ptr interface{}) { o.Password = *ptr.(*string) }
		if err := w.ReadPointer(&o.Password, _s_pszPasswd, _ptr_pszPasswd); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszFileName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszFileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
				return err
			}
			return nil
		})
		_s_pszFileName := func(ptr interface{}) { o.FileName = *ptr.(*string) }
		if err := w.ReadPointer(&o.FileName, _s_pszFileName, _ptr_pszFileName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszSourcePath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszSourcePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SourcePath); err != nil {
				return err
			}
			return nil
		})
		_s_pszSourcePath := func(ptr interface{}) { o.SourcePath = *ptr.(*string) }
		if err := w.ReadPointer(&o.SourcePath, _s_pszSourcePath, _ptr_pszSourcePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ExportOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ExportRequest structure represents the Export operation request
type ExportRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pszPasswd: A pointer to a Unicode string containing the password that will be used
	// to encrypt any secure properties being exported.
	Password string `idl:"name:pszPasswd;string;pointer:unique" json:"password"`
	// pszFileName: A pointer to a Unicode string containing the name of the file, including
	// the directory path, to which the data will be exported. The path MUST exist and be
	// local to the server.
	FileName string `idl:"name:pszFileName;string;pointer:unique" json:"file_name"`
	// pszSourcePath: A pointer to a Unicode string containing the path to the metabase
	// node to be exported.
	SourcePath string `idl:"name:pszSourcePath;string;pointer:unique" json:"source_path"`
	// dwMDFlags: A set of bit flags specifying the export operation to be performed. It
	// can be zero or one or more of the following values.
	//
	//	+--------------------------------+--------------------------------------------------------------------------+
	//	|                                |                                                                          |
	//	|             VALUE              |                                 MEANING                                  |
	//	|                                |                                                                          |
	//	+--------------------------------+--------------------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------------------+
	//	| MD_EXPORT_INHERITED 0x00000001 | Settings inherited from the parent nodes will be included in the export. |
	//	+--------------------------------+--------------------------------------------------------------------------+
	//	| MD_EXPORT_NODE_ONLY 0x00000002 | Child nodes will not be exported.                                        |
	//	+--------------------------------+--------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwMDFlags" json:"flags"`
}

func (o *ExportRequest) xxx_ToOp(ctx context.Context) *xxx_ExportOperation {
	if o == nil {
		return &xxx_ExportOperation{}
	}
	return &xxx_ExportOperation{
		This:       o.This,
		Password:   o.Password,
		FileName:   o.FileName,
		SourcePath: o.SourcePath,
		Flags:      o.Flags,
	}
}

func (o *ExportRequest) xxx_FromOp(ctx context.Context, op *xxx_ExportOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Password = op.Password
	o.FileName = op.FileName
	o.SourcePath = op.SourcePath
	o.Flags = op.Flags
}
func (o *ExportRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExportRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExportResponse structure represents the Export operation response
type ExportResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Export return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExportResponse) xxx_ToOp(ctx context.Context) *xxx_ExportOperation {
	if o == nil {
		return &xxx_ExportOperation{}
	}
	return &xxx_ExportOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ExportResponse) xxx_FromOp(ctx context.Context, op *xxx_ExportOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ExportResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExportResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ImportOperation structure represents the Import operation
type xxx_ImportOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	Password        string         `idl:"name:pszPasswd;string;pointer:unique" json:"password"`
	FileName        string         `idl:"name:pszFileName;string;pointer:unique" json:"file_name"`
	SourcePath      string         `idl:"name:pszSourcePath;string;pointer:unique" json:"source_path"`
	DestinationPath string         `idl:"name:pszDestPath;string;pointer:unique" json:"destination_path"`
	Flags           uint32         `idl:"name:dwMDFlags" json:"flags"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ImportOperation) OpNum() int { return 37 }

func (o *xxx_ImportOperation) OpName() string { return "/IMSAdminBase2W/v0/Import" }

func (o *xxx_ImportOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszPasswd {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Password != "" {
			_ptr_pszPasswd := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_pszPasswd); err != nil {
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
	// pszFileName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.FileName != "" {
			_ptr_pszFileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FileName, _ptr_pszFileName); err != nil {
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
	// pszSourcePath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.SourcePath != "" {
			_ptr_pszSourcePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SourcePath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SourcePath, _ptr_pszSourcePath); err != nil {
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
	// pszDestPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.DestinationPath != "" {
			_ptr_pszDestPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DestinationPath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DestinationPath, _ptr_pszDestPath); err != nil {
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
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszPasswd {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszPasswd := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
				return err
			}
			return nil
		})
		_s_pszPasswd := func(ptr interface{}) { o.Password = *ptr.(*string) }
		if err := w.ReadPointer(&o.Password, _s_pszPasswd, _ptr_pszPasswd); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszFileName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszFileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
				return err
			}
			return nil
		})
		_s_pszFileName := func(ptr interface{}) { o.FileName = *ptr.(*string) }
		if err := w.ReadPointer(&o.FileName, _s_pszFileName, _ptr_pszFileName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszSourcePath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszSourcePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SourcePath); err != nil {
				return err
			}
			return nil
		})
		_s_pszSourcePath := func(ptr interface{}) { o.SourcePath = *ptr.(*string) }
		if err := w.ReadPointer(&o.SourcePath, _s_pszSourcePath, _ptr_pszSourcePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszDestPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszDestPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DestinationPath); err != nil {
				return err
			}
			return nil
		})
		_s_pszDestPath := func(ptr interface{}) { o.DestinationPath = *ptr.(*string) }
		if err := w.ReadPointer(&o.DestinationPath, _s_pszDestPath, _ptr_pszDestPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ImportOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ImportRequest structure represents the Import operation request
type ImportRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pszPasswd: A pointer to a Unicode string containing the password that will be used
	// to decrypt the secure properties of the metabase data being imported.
	Password string `idl:"name:pszPasswd;string;pointer:unique" json:"password"`
	// pszFileName: A pointer to a Unicode string containing the name of the file, including
	// directory path, to import settings from. This file will have been created using the
	// Export function.
	FileName string `idl:"name:pszFileName;string;pointer:unique" json:"file_name"`
	// pszSourcePath: A pointer to a Unicode string containing the path to the metabase
	// node being imported from the file specified in pszFileName.
	SourcePath string `idl:"name:pszSourcePath;string;pointer:unique" json:"source_path"`
	// pszDestPath: A pointer to a Unicode string containing the path to the metabase node
	// into which the file data will be imported.
	DestinationPath string `idl:"name:pszDestPath;string;pointer:unique" json:"destination_path"`
	// dwMDFlags: A set of bit flags specifying the import operation to be performed. It
	// can be zero or one or more of the following values.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|                                |                                                                                  |
	//	|             VALUE              |                                     MEANING                                      |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| MD_IMPORT_INHERITED 0x00000001 | Inherited settings that were exported using the MD_EXPORT_INHERITED flag will be |
	//	|                                | imported.                                                                        |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| MD_IMPORT_NODE_ONLY 0x00000002 | Child nodes will not be imported.                                                |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| MD_IMPORT_MERGE 0x00000004     | Imported settings will be merged with any matching existing node settings.       |
	//	|                                | When a value for a setting is present in the data file and also in the current   |
	//	|                                | metabase, the data file setting will overwrite the existing metabase setting. If |
	//	|                                | this flag is not set and there is a current node in the metabase that conflicts  |
	//	|                                | with the node being imported, the imported node will replace the existing node.  |
	//	|                                | All settings from the existing node will be lost regardless of whether the       |
	//	|                                | imported node contains the setting or not.                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwMDFlags" json:"flags"`
}

func (o *ImportRequest) xxx_ToOp(ctx context.Context) *xxx_ImportOperation {
	if o == nil {
		return &xxx_ImportOperation{}
	}
	return &xxx_ImportOperation{
		This:            o.This,
		Password:        o.Password,
		FileName:        o.FileName,
		SourcePath:      o.SourcePath,
		DestinationPath: o.DestinationPath,
		Flags:           o.Flags,
	}
}

func (o *ImportRequest) xxx_FromOp(ctx context.Context, op *xxx_ImportOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Password = op.Password
	o.FileName = op.FileName
	o.SourcePath = op.SourcePath
	o.DestinationPath = op.DestinationPath
	o.Flags = op.Flags
}
func (o *ImportRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ImportRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ImportResponse structure represents the Import operation response
type ImportResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Import return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ImportResponse) xxx_ToOp(ctx context.Context) *xxx_ImportOperation {
	if o == nil {
		return &xxx_ImportOperation{}
	}
	return &xxx_ImportOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ImportResponse) xxx_FromOp(ctx context.Context, op *xxx_ImportOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ImportResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ImportResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RestoreHistoryOperation structure represents the RestoreHistory operation
type xxx_RestoreHistoryOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	HistoryLocation string         `idl:"name:pszMDHistoryLocation;string;pointer:unique" json:"history_location"`
	MajorVersion    uint32         `idl:"name:dwMDMajorVersion" json:"major_version"`
	MinorVersion    uint32         `idl:"name:dwMDMinorVersion" json:"minor_version"`
	Flags           uint32         `idl:"name:dwMDFlags" json:"flags"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RestoreHistoryOperation) OpNum() int { return 38 }

func (o *xxx_RestoreHistoryOperation) OpName() string { return "/IMSAdminBase2W/v0/RestoreHistory" }

func (o *xxx_RestoreHistoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreHistoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszMDHistoryLocation {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.HistoryLocation != "" {
			_ptr_pszMDHistoryLocation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.HistoryLocation); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.HistoryLocation, _ptr_pszMDHistoryLocation); err != nil {
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
	// dwMDMajorVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// dwMDMinorVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MinorVersion); err != nil {
			return err
		}
	}
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreHistoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszMDHistoryLocation {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDHistoryLocation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.HistoryLocation); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDHistoryLocation := func(ptr interface{}) { o.HistoryLocation = *ptr.(*string) }
		if err := w.ReadPointer(&o.HistoryLocation, _s_pszMDHistoryLocation, _ptr_pszMDHistoryLocation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDMajorVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// dwMDMinorVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MinorVersion); err != nil {
			return err
		}
	}
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreHistoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreHistoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RestoreHistoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RestoreHistoryRequest structure represents the RestoreHistory operation request
type RestoreHistoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pszMDHistoryLocation: A pointer to a Unicode string containing the absolute path
	// to the location of the history files for the metabase. If an empty string is passed
	// to this function, the server SHOULD use the default history path.<25>
	HistoryLocation string `idl:"name:pszMDHistoryLocation;string;pointer:unique" json:"history_location"`
	// dwMDMajorVersion: An integer value containing the predecimal version value of the
	// history entry to restore from. If the dwMDFlags parameter contains the MD_HISTORY_LATEST
	// flag, this value MUST be set to zero.
	MajorVersion uint32 `idl:"name:dwMDMajorVersion" json:"major_version"`
	// dwMDMinorVersion: An integer value containing the postdecimal version value of the
	// history entry to restore from. If the dwMDFlags parameter contains the MD_HISTORY_LATEST
	// flag, this value MUST be set to zero.
	MinorVersion uint32 `idl:"name:dwMDMinorVersion" json:"minor_version"`
	// dwMDFlags: A set of bit flags specifying the options to be executed during the RestoreHistory
	// call.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MD_HISTORY_LATEST 0x00000001 | Restore to the most recent history file. If this is set, the dwMDMajorVersion    |
	//	|                              | and dwMDMinorVersion parameters must be set to zero.                             |
	//	+------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwMDFlags" json:"flags"`
}

func (o *RestoreHistoryRequest) xxx_ToOp(ctx context.Context) *xxx_RestoreHistoryOperation {
	if o == nil {
		return &xxx_RestoreHistoryOperation{}
	}
	return &xxx_RestoreHistoryOperation{
		This:            o.This,
		HistoryLocation: o.HistoryLocation,
		MajorVersion:    o.MajorVersion,
		MinorVersion:    o.MinorVersion,
		Flags:           o.Flags,
	}
}

func (o *RestoreHistoryRequest) xxx_FromOp(ctx context.Context, op *xxx_RestoreHistoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.HistoryLocation = op.HistoryLocation
	o.MajorVersion = op.MajorVersion
	o.MinorVersion = op.MinorVersion
	o.Flags = op.Flags
}
func (o *RestoreHistoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RestoreHistoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreHistoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RestoreHistoryResponse structure represents the RestoreHistory operation response
type RestoreHistoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RestoreHistory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RestoreHistoryResponse) xxx_ToOp(ctx context.Context) *xxx_RestoreHistoryOperation {
	if o == nil {
		return &xxx_RestoreHistoryOperation{}
	}
	return &xxx_RestoreHistoryOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RestoreHistoryResponse) xxx_FromOp(ctx context.Context, op *xxx_RestoreHistoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RestoreHistoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RestoreHistoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreHistoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumHistoryOperation structure represents the EnumHistory operation
type xxx_EnumHistoryOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	HistoryLocation string         `idl:"name:pszMDHistoryLocation;size_is:(100)" json:"history_location"`
	MajorVersion    uint32         `idl:"name:pdwMDMajorVersion" json:"major_version"`
	MinorVersion    uint32         `idl:"name:pdwMDMinorVersion" json:"minor_version"`
	HistoryTime     *dtyp.Filetime `idl:"name:pftMDHistoryTime" json:"history_time"`
	EnumIndex       uint32         `idl:"name:dwMDEnumIndex" json:"enum_index"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumHistoryOperation) OpNum() int { return 39 }

func (o *xxx_EnumHistoryOperation) OpName() string { return "/IMSAdminBase2W/v0/EnumHistory" }

func (o *xxx_EnumHistoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	// cannot evaluate expression 100
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumHistoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszMDHistoryLocation {in, out} (1:{alias=LPWSTR}*(1)[dim:0,size_is=100,string](wchar))
	{
		dimSize1 := uint64(100)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_HistoryLocation_buf := utf16.Encode([]rune(o.HistoryLocation))
		if uint64(len(_HistoryLocation_buf)) > sizeInfo[0] {
			_HistoryLocation_buf = _HistoryLocation_buf[:sizeInfo[0]]
		}
		for i1 := range _HistoryLocation_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_HistoryLocation_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_HistoryLocation_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// dwMDEnumIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EnumIndex); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumHistoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszMDHistoryLocation {in, out} (1:{alias=LPWSTR,pointer=ref}*(1)[dim:0,size_is=100,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _HistoryLocation_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _HistoryLocation_buf", sizeInfo[0])
		}
		_HistoryLocation_buf = make([]uint16, sizeInfo[0])
		for i1 := range _HistoryLocation_buf {
			i1 := i1
			if err := w.ReadData(&_HistoryLocation_buf[i1]); err != nil {
				return err
			}
		}
		o.HistoryLocation = strings.TrimRight(string(utf16.Decode(_HistoryLocation_buf)), ndr.ZeroString)
	}
	// dwMDEnumIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EnumIndex); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumHistoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	// cannot evaluate expression 100
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumHistoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pszMDHistoryLocation {in, out} (1:{alias=LPWSTR}*(1)[dim:0,size_is=100,string](wchar))
	{
		dimSize1 := uint64(100)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_HistoryLocation_buf := utf16.Encode([]rune(o.HistoryLocation))
		if uint64(len(_HistoryLocation_buf)) > sizeInfo[0] {
			_HistoryLocation_buf = _HistoryLocation_buf[:sizeInfo[0]]
		}
		for i1 := range _HistoryLocation_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_HistoryLocation_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_HistoryLocation_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// pdwMDMajorVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// pdwMDMinorVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MinorVersion); err != nil {
			return err
		}
	}
	// pftMDHistoryTime {out} (1:{alias=PFILETIME}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.HistoryTime != nil {
			if err := o.HistoryTime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EnumHistoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pszMDHistoryLocation {in, out} (1:{alias=LPWSTR,pointer=ref}*(1)[dim:0,size_is=100,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _HistoryLocation_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _HistoryLocation_buf", sizeInfo[0])
		}
		_HistoryLocation_buf = make([]uint16, sizeInfo[0])
		for i1 := range _HistoryLocation_buf {
			i1 := i1
			if err := w.ReadData(&_HistoryLocation_buf[i1]); err != nil {
				return err
			}
		}
		o.HistoryLocation = strings.TrimRight(string(utf16.Decode(_HistoryLocation_buf)), ndr.ZeroString)
	}
	// pdwMDMajorVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// pdwMDMinorVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MinorVersion); err != nil {
			return err
		}
	}
	// pftMDHistoryTime {out} (1:{alias=PFILETIME,pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.HistoryTime == nil {
			o.HistoryTime = &dtyp.Filetime{}
		}
		if err := o.HistoryTime.UnmarshalNDR(ctx, w); err != nil {
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

// EnumHistoryRequest structure represents the EnumHistory operation request
type EnumHistoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pszMDHistoryLocation: A pointer to a Unicode string that on input contains the path
	// to the history files being enumerated. If this is an empty string, the server SHOULD
	// use a default path. If an empty string is passed in, the default history path will
	// be written to the buffer.<22>
	HistoryLocation string `idl:"name:pszMDHistoryLocation;size_is:(100)" json:"history_location"`
	// dwMDEnumIndex: An integer value containing the current index of the history entry
	// to be enumerated. This value SHOULD start at zero on the first call and SHOULD be
	// increased by one on subsequent calls until the last entry in the history is reached.
	// This indexing is controlled by the client, so the client is responsible for selecting
	// the next history file to be enumerated.
	EnumIndex uint32 `idl:"name:dwMDEnumIndex" json:"enum_index"`
}

func (o *EnumHistoryRequest) xxx_ToOp(ctx context.Context) *xxx_EnumHistoryOperation {
	if o == nil {
		return &xxx_EnumHistoryOperation{}
	}
	return &xxx_EnumHistoryOperation{
		This:            o.This,
		HistoryLocation: o.HistoryLocation,
		EnumIndex:       o.EnumIndex,
	}
}

func (o *EnumHistoryRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumHistoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.HistoryLocation = op.HistoryLocation
	o.EnumIndex = op.EnumIndex
}
func (o *EnumHistoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumHistoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumHistoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumHistoryResponse structure represents the EnumHistory operation response
type EnumHistoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pszMDHistoryLocation: A pointer to a Unicode string that on input contains the path
	// to the history files being enumerated. If this is an empty string, the server SHOULD
	// use a default path. If an empty string is passed in, the default history path will
	// be written to the buffer.<22>
	HistoryLocation string         `idl:"name:pszMDHistoryLocation;size_is:(100)" json:"history_location"`
	MajorVersion    uint32         `idl:"name:pdwMDMajorVersion" json:"major_version"`
	MinorVersion    uint32         `idl:"name:pdwMDMinorVersion" json:"minor_version"`
	HistoryTime     *dtyp.Filetime `idl:"name:pftMDHistoryTime" json:"history_time"`
	// Return: The EnumHistory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumHistoryResponse) xxx_ToOp(ctx context.Context) *xxx_EnumHistoryOperation {
	if o == nil {
		return &xxx_EnumHistoryOperation{}
	}
	return &xxx_EnumHistoryOperation{
		That:            o.That,
		HistoryLocation: o.HistoryLocation,
		MajorVersion:    o.MajorVersion,
		MinorVersion:    o.MinorVersion,
		HistoryTime:     o.HistoryTime,
		Return:          o.Return,
	}
}

func (o *EnumHistoryResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumHistoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.HistoryLocation = op.HistoryLocation
	o.MajorVersion = op.MajorVersion
	o.MinorVersion = op.MinorVersion
	o.HistoryTime = op.HistoryTime
	o.Return = op.Return
}
func (o *EnumHistoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumHistoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumHistoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
