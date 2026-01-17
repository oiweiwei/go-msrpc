package iremotesstpcertcheck

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	rrasm "github.com/oiweiwei/go-msrpc/msrpc/dcom/rrasm"
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
	_ = rrasm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rrasm"
)

var (
	// IRemoteSstpCertCheck interface identifier 5ff9bdf6-bd91-4d8b-a614-d6317acc8dd8
	RemoteSSTPCertCheckIID = &dcom.IID{Data1: 0x5ff9bdf6, Data2: 0xbd91, Data3: 0x4d8b, Data4: []byte{0xa6, 0x14, 0xd6, 0x31, 0x7a, 0xcc, 0x8d, 0xd8}}
	// Syntax UUID
	RemoteSSTPCertCheckSyntaxUUID = &uuid.UUID{TimeLow: 0x5ff9bdf6, TimeMid: 0xbd91, TimeHiAndVersion: 0x4d8b, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0x14, Node: [6]uint8{0xd6, 0x31, 0x7a, 0xcc, 0x8d, 0xd8}}
	// Syntax ID
	RemoteSSTPCertCheckSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteSSTPCertCheckSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRemoteSstpCertCheck interface.
type RemoteSSTPCertCheckClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// CheckIfCertificateAllowedRR operation.
	CheckInterfaceCertificateAllowedRR(context.Context, *CheckInterfaceCertificateAllowedRRRequest, ...dcerpc.CallOption) (*CheckInterfaceCertificateAllowedRRResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteSSTPCertCheckClient
}

type xxx_DefaultRemoteSSTPCertCheckClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteSSTPCertCheckClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRemoteSSTPCertCheckClient) CheckInterfaceCertificateAllowedRR(ctx context.Context, in *CheckInterfaceCertificateAllowedRRRequest, opts ...dcerpc.CallOption) (*CheckInterfaceCertificateAllowedRRResponse, error) {
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
	out := &CheckInterfaceCertificateAllowedRRResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteSSTPCertCheckClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteSSTPCertCheckClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRemoteSSTPCertCheckClient) IPID(ctx context.Context, ipid *dcom.IPID) RemoteSSTPCertCheckClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteSSTPCertCheckClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewRemoteSSTPCertCheckClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteSSTPCertCheckClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteSSTPCertCheckSyntaxV0_0))...)
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
	return &xxx_DefaultRemoteSSTPCertCheckClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CheckInterfaceCertificateAllowedRROperation structure represents the CheckIfCertificateAllowedRR operation
type xxx_CheckInterfaceCertificateAllowedRROperation struct {
	This          *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat       `idl:"name:That" json:"that"`
	AdminCertName string               `idl:"name:adminCertName" json:"admin_cert_name"`
	CertSHA1      *rrasm.SSTPCertInfo1 `idl:"name:certSha1" json:"cert_sha1"`
	CertSHA256    *rrasm.SSTPCertInfo1 `idl:"name:certSha256" json:"cert_sha256"`
	Return        int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_CheckInterfaceCertificateAllowedRROperation) OpNum() int { return 3 }

func (o *xxx_CheckInterfaceCertificateAllowedRROperation) OpName() string {
	return "/IRemoteSstpCertCheck/v0/CheckIfCertificateAllowedRR"
}

func (o *xxx_CheckInterfaceCertificateAllowedRROperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckInterfaceCertificateAllowedRROperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// adminCertName {in} (1:{string, alias=PCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.AdminCertName); err != nil {
			return err
		}
	}
	// certSha1 {in, out} (1:{alias=PSSTP_CERT_INFO_1}*(1))(2:{alias=SSTP_CERT_INFO_1}(struct))
	{
		if o.CertSHA1 != nil {
			if err := o.CertSHA1.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.SSTPCertInfo1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// certSha256 {in, out} (1:{alias=PSSTP_CERT_INFO_1}*(1))(2:{alias=SSTP_CERT_INFO_1}(struct))
	{
		if o.CertSHA256 != nil {
			if err := o.CertSHA256.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.SSTPCertInfo1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckInterfaceCertificateAllowedRROperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// adminCertName {in} (1:{string, alias=PCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.AdminCertName); err != nil {
			return err
		}
	}
	// certSha1 {in, out} (1:{alias=PSSTP_CERT_INFO_1,pointer=ref}*(1))(2:{alias=SSTP_CERT_INFO_1}(struct))
	{
		if o.CertSHA1 == nil {
			o.CertSHA1 = &rrasm.SSTPCertInfo1{}
		}
		if err := o.CertSHA1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// certSha256 {in, out} (1:{alias=PSSTP_CERT_INFO_1,pointer=ref}*(1))(2:{alias=SSTP_CERT_INFO_1}(struct))
	{
		if o.CertSHA256 == nil {
			o.CertSHA256 = &rrasm.SSTPCertInfo1{}
		}
		if err := o.CertSHA256.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckInterfaceCertificateAllowedRROperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckInterfaceCertificateAllowedRROperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// certSha1 {in, out} (1:{alias=PSSTP_CERT_INFO_1}*(1))(2:{alias=SSTP_CERT_INFO_1}(struct))
	{
		if o.CertSHA1 != nil {
			if err := o.CertSHA1.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.SSTPCertInfo1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// certSha256 {in, out} (1:{alias=PSSTP_CERT_INFO_1}*(1))(2:{alias=SSTP_CERT_INFO_1}(struct))
	{
		if o.CertSHA256 != nil {
			if err := o.CertSHA256.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.SSTPCertInfo1{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CheckInterfaceCertificateAllowedRROperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// certSha1 {in, out} (1:{alias=PSSTP_CERT_INFO_1,pointer=ref}*(1))(2:{alias=SSTP_CERT_INFO_1}(struct))
	{
		if o.CertSHA1 == nil {
			o.CertSHA1 = &rrasm.SSTPCertInfo1{}
		}
		if err := o.CertSHA1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// certSha256 {in, out} (1:{alias=PSSTP_CERT_INFO_1,pointer=ref}*(1))(2:{alias=SSTP_CERT_INFO_1}(struct))
	{
		if o.CertSHA256 == nil {
			o.CertSHA256 = &rrasm.SSTPCertInfo1{}
		}
		if err := o.CertSHA256.UnmarshalNDR(ctx, w); err != nil {
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

// CheckInterfaceCertificateAllowedRRRequest structure represents the CheckIfCertificateAllowedRR operation request
type CheckInterfaceCertificateAllowedRRRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis       `idl:"name:This" json:"this"`
	AdminCertName string               `idl:"name:adminCertName" json:"admin_cert_name"`
	CertSHA1      *rrasm.SSTPCertInfo1 `idl:"name:certSha1" json:"cert_sha1"`
	CertSHA256    *rrasm.SSTPCertInfo1 `idl:"name:certSha256" json:"cert_sha256"`
}

func (o *CheckInterfaceCertificateAllowedRRRequest) xxx_ToOp(ctx context.Context, op *xxx_CheckInterfaceCertificateAllowedRROperation) *xxx_CheckInterfaceCertificateAllowedRROperation {
	if op == nil {
		op = &xxx_CheckInterfaceCertificateAllowedRROperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.AdminCertName = o.AdminCertName
	op.CertSHA1 = o.CertSHA1
	op.CertSHA256 = o.CertSHA256
	return op
}

func (o *CheckInterfaceCertificateAllowedRRRequest) xxx_FromOp(ctx context.Context, op *xxx_CheckInterfaceCertificateAllowedRROperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AdminCertName = op.AdminCertName
	o.CertSHA1 = op.CertSHA1
	o.CertSHA256 = op.CertSHA256
}
func (o *CheckInterfaceCertificateAllowedRRRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CheckInterfaceCertificateAllowedRRRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CheckInterfaceCertificateAllowedRROperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CheckInterfaceCertificateAllowedRRResponse structure represents the CheckIfCertificateAllowedRR operation response
type CheckInterfaceCertificateAllowedRRResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat       `idl:"name:That" json:"that"`
	CertSHA1   *rrasm.SSTPCertInfo1 `idl:"name:certSha1" json:"cert_sha1"`
	CertSHA256 *rrasm.SSTPCertInfo1 `idl:"name:certSha256" json:"cert_sha256"`
	// Return: The CheckIfCertificateAllowedRR return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CheckInterfaceCertificateAllowedRRResponse) xxx_ToOp(ctx context.Context, op *xxx_CheckInterfaceCertificateAllowedRROperation) *xxx_CheckInterfaceCertificateAllowedRROperation {
	if op == nil {
		op = &xxx_CheckInterfaceCertificateAllowedRROperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.CertSHA1 = o.CertSHA1
	op.CertSHA256 = o.CertSHA256
	op.Return = o.Return
	return op
}

func (o *CheckInterfaceCertificateAllowedRRResponse) xxx_FromOp(ctx context.Context, op *xxx_CheckInterfaceCertificateAllowedRROperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CertSHA1 = op.CertSHA1
	o.CertSHA256 = op.CertSHA256
	o.Return = op.Return
}
func (o *CheckInterfaceCertificateAllowedRRResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CheckInterfaceCertificateAllowedRRResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CheckInterfaceCertificateAllowedRROperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
