package iexport

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// IExport interface identifier cfadac84-e12c-11d1-b34c-00c04f990d54
	ExportIID = &dcom.IID{Data1: 0xcfadac84, Data2: 0xe12c, Data3: 0x11d1, Data4: []byte{0xb3, 0x4c, 0x00, 0xc0, 0x4f, 0x99, 0x0d, 0x54}}
	// Syntax UUID
	ExportSyntaxUUID = &uuid.UUID{TimeLow: 0xcfadac84, TimeMid: 0xe12c, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x4c, Node: [6]uint8{0x0, 0xc0, 0x4f, 0x99, 0xd, 0x54}}
	// Syntax ID
	ExportSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ExportSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IExport interface.
type ExportClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to export a conglomeration to an installer package
	// file.
	//
	// pConglomerationIdentifier: The conglomeration identifier of a conglomeration on the
	// server.
	//
	// pwszInstallerPackage: A path in UNC that is to be used as the location for the server
	// to create an installer package file.
	//
	// pwszReserved:  MUST be an empty (zero-length) string.
	//
	// dwFlags:  MUST be a combination of zero or more of the following flags.
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
	//	|                              | setting the IsProxyApp property to TRUE (0x00000001).                            |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| fEXPORT_CATVER300 0x00000080 | The server SHOULD only include configuration that is defined in catalog version  |
	//	|                              | 3.00.                                                                            |
	//	+------------------------------+----------------------------------------------------------------------------------+
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
	ExportConglomeration(context.Context, *ExportConglomerationRequest, ...dcerpc.CallOption) (*ExportConglomerationResponse, error)

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ExportClient
}

type xxx_DefaultExportClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultExportClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultExportClient) ExportConglomeration(ctx context.Context, in *ExportConglomerationRequest, opts ...dcerpc.CallOption) (*ExportConglomerationResponse, error) {
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
	out := &ExportConglomerationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultExportClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultExportClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultExportClient) IPID(ctx context.Context, ipid *dcom.IPID) ExportClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultExportClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewExportClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ExportClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ExportSyntaxV0_0))...)
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
	return &xxx_DefaultExportClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ExportConglomerationOperation structure represents the ExportConglomeration operation
type xxx_ExportConglomerationOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConglomerationID *dtyp.GUID     `idl:"name:pConglomerationIdentifier" json:"conglomeration_id"`
	InstallerPackage string         `idl:"name:pwszInstallerPackage" json:"installer_package"`
	_                string         `idl:"name:pwszReserved"`
	Flags            uint32         `idl:"name:dwFlags" json:"flags"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ExportConglomerationOperation) OpNum() int { return 3 }

func (o *xxx_ExportConglomerationOperation) OpName() string {
	return "/IExport/v0/ExportConglomeration"
}

func (o *xxx_ExportConglomerationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportConglomerationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pConglomerationIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ConglomerationID != nil {
			if err := o.ConglomerationID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ExportConglomerationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pConglomerationIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ConglomerationID == nil {
			o.ConglomerationID = &dtyp.GUID{}
		}
		if err := o.ConglomerationID.UnmarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ExportConglomerationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportConglomerationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ExportConglomerationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ExportConglomerationRequest structure represents the ExportConglomeration operation request
type ExportConglomerationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	ConglomerationID *dtyp.GUID     `idl:"name:pConglomerationIdentifier" json:"conglomeration_id"`
	InstallerPackage string         `idl:"name:pwszInstallerPackage" json:"installer_package"`
	Flags            uint32         `idl:"name:dwFlags" json:"flags"`
}

func (o *ExportConglomerationRequest) xxx_ToOp(ctx context.Context, op *xxx_ExportConglomerationOperation) *xxx_ExportConglomerationOperation {
	if op == nil {
		op = &xxx_ExportConglomerationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ConglomerationID = o.ConglomerationID
	op.InstallerPackage = o.InstallerPackage
	op.Flags = o.Flags
	return op
}

func (o *ExportConglomerationRequest) xxx_FromOp(ctx context.Context, op *xxx_ExportConglomerationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConglomerationID = op.ConglomerationID
	o.InstallerPackage = op.InstallerPackage
	o.Flags = op.Flags
}
func (o *ExportConglomerationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ExportConglomerationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportConglomerationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExportConglomerationResponse structure represents the ExportConglomeration operation response
type ExportConglomerationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExportConglomeration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExportConglomerationResponse) xxx_ToOp(ctx context.Context, op *xxx_ExportConglomerationOperation) *xxx_ExportConglomerationOperation {
	if op == nil {
		op = &xxx_ExportConglomerationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ExportConglomerationResponse) xxx_FromOp(ctx context.Context, op *xxx_ExportConglomerationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ExportConglomerationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ExportConglomerationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportConglomerationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
