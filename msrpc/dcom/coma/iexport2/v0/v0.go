package iexport2

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
	_ = iunknown.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// IExport2 interface identifier f131ea3e-b7be-480e-a60d-51cb2785779e
	Export2IID = &dcom.IID{Data1: 0xf131ea3e, Data2: 0xb7be, Data3: 0x480e, Data4: []byte{0xa6, 0x0d, 0x51, 0xcb, 0x27, 0x85, 0x77, 0x9e}}
	// Syntax UUID
	Export2SyntaxUUID = &uuid.UUID{TimeLow: 0xf131ea3e, TimeMid: 0xb7be, TimeHiAndVersion: 0x480e, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0xd, Node: [6]uint8{0x51, 0xcb, 0x27, 0x85, 0x77, 0x9e}}
	// Syntax ID
	Export2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Export2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IExport2 interface.
type Export2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to export all conglomerations in a partition at
	// once to an installer package file.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF], section 2.1, on failure. All failure results
	// MUST be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	ExportPartition(context.Context, *ExportPartitionRequest, ...dcerpc.CallOption) (*ExportPartitionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Export2Client
}

type xxx_DefaultExport2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultExport2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultExport2Client) ExportPartition(ctx context.Context, in *ExportPartitionRequest, opts ...dcerpc.CallOption) (*ExportPartitionResponse, error) {
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
	out := &ExportPartitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultExport2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultExport2Client) IPID(ctx context.Context, ipid *dcom.IPID) Export2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultExport2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewExport2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Export2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Export2SyntaxV0_0))...)
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
	return &xxx_DefaultExport2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ExportPartitionOperation structure represents the ExportPartition operation
type xxx_ExportPartitionOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	PartitionID      *dtyp.GUID     `idl:"name:pPartitionIdentifier" json:"partition_id"`
	InstallerPackage string         `idl:"name:pwszInstallerPackage" json:"installer_package"`
	_                string         `idl:"name:pwszReserved"`
	Flags            uint32         `idl:"name:dwFlags" json:"flags"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ExportPartitionOperation) OpNum() int { return 3 }

func (o *xxx_ExportPartitionOperation) OpName() string { return "/IExport2/v0/ExportPartition" }

func (o *xxx_ExportPartitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportPartitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pPartitionIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PartitionID != nil {
			if err := o.PartitionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pwszInstallerPackage {in} (1:{alias=LPCWSTR}*(1)[dim:0,string](wchar))
	{
		if err := ndr.WriteUTF16String(ctx, w, o.InstallerPackage); err != nil {
			return err
		}
	}
	// pwszReserved {in} (1:{alias=LPCWSTR}*(1)[dim:0,string](wchar))
	{
		// reserved pwszReserved
		if err := ndr.WriteUTF16String(ctx, w, ""); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportPartitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pPartitionIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PartitionID == nil {
			o.PartitionID = &dtyp.GUID{}
		}
		if err := o.PartitionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pwszInstallerPackage {in} (1:{alias=LPCWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		if err := ndr.ReadUTF16String(ctx, w, &o.InstallerPackage); err != nil {
			return err
		}
	}
	// pwszReserved {in} (1:{alias=LPCWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		// reserved pwszReserved
		var _pwszReserved string
		if err := ndr.ReadUTF16String(ctx, w, &_pwszReserved); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportPartitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportPartitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ExportPartitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ExportPartitionRequest structure represents the ExportPartition operation request
type ExportPartitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pPartitionIdentifier: The partition identifier of a partition other than the global
	// partition on the server.
	PartitionID *dtyp.GUID `idl:"name:pPartitionIdentifier" json:"partition_id"`
	// pwszInstallerPackage:  A path in UNC that is to be used as the location for the
	// server to create an installer package file.
	InstallerPackage string `idl:"name:pwszInstallerPackage" json:"installer_package"`
	// dwFlags:  MUST be a combination of zero or more of the following flags.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|             FLAG             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| fEXPORT_OVERWRITE 0x00000001 | The server SHOULD mark the installer package file with a directive that existing |
	//	|                              | files be overwritten on import (section 3.1.4.12.1).                             |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| fEXPORT_WITHUSERS 0x00000010 | The server SHOULD include user account information in the installer package      |
	//	|                              | file.                                                                            |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| fEXPORT_PROXY 0x00000020     | The server SHOULD mark the exported conglomeration as a proxy conglomeration by  |
	//	|                              | setting the IsProxyApp property is set to TRUE (0x00000001).                     |
	//	+------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
}

func (o *ExportPartitionRequest) xxx_ToOp(ctx context.Context) *xxx_ExportPartitionOperation {
	if o == nil {
		return &xxx_ExportPartitionOperation{}
	}
	return &xxx_ExportPartitionOperation{
		This:             o.This,
		PartitionID:      o.PartitionID,
		InstallerPackage: o.InstallerPackage,
		Flags:            o.Flags,
	}
}

func (o *ExportPartitionRequest) xxx_FromOp(ctx context.Context, op *xxx_ExportPartitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PartitionID = op.PartitionID
	o.InstallerPackage = op.InstallerPackage
	o.Flags = op.Flags
}
func (o *ExportPartitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExportPartitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportPartitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExportPartitionResponse structure represents the ExportPartition operation response
type ExportPartitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExportPartition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExportPartitionResponse) xxx_ToOp(ctx context.Context) *xxx_ExportPartitionOperation {
	if o == nil {
		return &xxx_ExportPartitionOperation{}
	}
	return &xxx_ExportPartitionOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ExportPartitionResponse) xxx_FromOp(ctx context.Context, op *xxx_ExportPartitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ExportPartitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExportPartitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportPartitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
