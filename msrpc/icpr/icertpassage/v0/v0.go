package icertpassage

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	wcce "github.com/oiweiwei/go-msrpc/msrpc/dcom/wcce"
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
	_ = wcce.GoPackage
)

var (
	// import guard
	GoPackage = "icpr"
)

var (
	// Syntax UUID
	CertPassageSyntaxUUID = &uuid.UUID{TimeLow: 0x91ae6020, TimeMid: 0x9e3c, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x7c, Node: [6]uint8{0x0, 0xaa, 0x0, 0xc0, 0x91, 0xbe}}
	// Syntax ID
	CertPassageSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CertPassageSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICertPassage interface.
type CertPassageClient interface {

	// The CertServerRequest method processes a certificate enrollment request from the
	// client.<6>
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success. This
	// method's return values MUST have identical syntax and semantics to the return values
	// specified in [MS-WCCE] section 3.2.1.4.2.1.
	//
	// If the ADM element Config.CA.Interface.Flags contains the value IF_NORPCICERTREQUEST,
	// the server SHOULD return an error.<7>
	//
	// If the ADM element Config.CA.Interface.Flags contains the value IF_ENFORCEENCRYPTICERTREQUEST
	// and the RPC_C_AUTHN_LEVEL_PKT_PRIVACY authentication level ([MS-RPCE] section 2.2.1.1.8)
	// is not specified on the RPC connection from the client, the CA MUST refuse to establish
	// a connection with the client by returning E_ACCESSDENIED (0x80000009).
	//
	// Otherwise, the processing rules for the ICertRequestD::Request method ([MS-WCCE]
	// section 3.2.2.6.2.1) apply, except that if the ADM element Config.CA.Interface.Flags
	// contains the value IF_NOREMOTEICERTREQUEST, these values are ignored and the request
	// is processed as though the values were absent.
	CertServerRequest(context.Context, *CertServerRequestRequest, ...dcerpc.CallOption) (*CertServerRequestResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultCertPassageClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultCertPassageClient) CertServerRequest(ctx context.Context, in *CertServerRequestRequest, opts ...dcerpc.CallOption) (*CertServerRequestResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CertServerRequestResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertPassageClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCertPassageClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewCertPassageClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CertPassageClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CertPassageSyntaxV0_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultCertPassageClient{cc: cc}, nil
}

// xxx_CertServerRequestOperation structure represents the CertServerRequest operation
type xxx_CertServerRequestOperation struct {
	Flags              uint32                  `idl:"name:dwFlags" json:"flags"`
	Authority          string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	RequestID          uint32                  `idl:"name:pdwRequestId;pointer:ref" json:"request_id"`
	Disposition        uint32                  `idl:"name:pdwDisposition" json:"disposition"`
	Attributes         *wcce.CertTransportBlob `idl:"name:pctbAttribs;pointer:ref" json:"attributes"`
	Request            *wcce.CertTransportBlob `idl:"name:pctbRequest;pointer:ref" json:"request"`
	Cert               *wcce.CertTransportBlob `idl:"name:pctbCert;pointer:ref" json:"cert"`
	EncodedCert        *wcce.CertTransportBlob `idl:"name:pctbEncodedCert;pointer:ref" json:"encoded_cert"`
	DispositionMessage *wcce.CertTransportBlob `idl:"name:pctbDispositionMessage;pointer:ref" json:"disposition_message"`
	Return             uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_CertServerRequestOperation) OpNum() int { return 0 }

func (o *xxx_CertServerRequestOperation) OpName() string { return "/ICertPassage/v0/CertServerRequest" }

func (o *xxx_CertServerRequestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CertServerRequestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
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
	// pdwRequestId {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	// pctbAttribs {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Attributes != nil {
			if err := o.Attributes.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wcce.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pctbRequest {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Request != nil {
			if err := o.Request.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wcce.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CertServerRequestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwRequestId {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestID); err != nil {
			return err
		}
	}
	// pctbAttribs {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Attributes == nil {
			o.Attributes = &wcce.CertTransportBlob{}
		}
		if err := o.Attributes.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pctbRequest {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Request == nil {
			o.Request = &wcce.CertTransportBlob{}
		}
		if err := o.Request.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CertServerRequestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CertServerRequestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwRequestId {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	// pdwDisposition {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Disposition); err != nil {
			return err
		}
	}
	// pctbCert {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Cert != nil {
			if err := o.Cert.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wcce.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pctbEncodedCert {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.EncodedCert != nil {
			if err := o.EncodedCert.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wcce.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pctbDispositionMessage {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.DispositionMessage != nil {
			if err := o.DispositionMessage.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wcce.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CertServerRequestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwRequestId {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestID); err != nil {
			return err
		}
	}
	// pdwDisposition {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Disposition); err != nil {
			return err
		}
	}
	// pctbCert {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Cert == nil {
			o.Cert = &wcce.CertTransportBlob{}
		}
		if err := o.Cert.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pctbEncodedCert {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.EncodedCert == nil {
			o.EncodedCert = &wcce.CertTransportBlob{}
		}
		if err := o.EncodedCert.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pctbDispositionMessage {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.DispositionMessage == nil {
			o.DispositionMessage = &wcce.CertTransportBlob{}
		}
		if err := o.DispositionMessage.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CertServerRequestRequest structure represents the CertServerRequest operation request
type CertServerRequestRequest struct {
	// dwFlags: The dwFlags parameter has identical syntax and semantics to the dwFlags
	// parameter specified in [MS-WCCE] section 3.2.1.4.2.1.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pwszAuthority: The pwszAuthority parameter has identical syntax and semantics to
	// the pwszAuthority parameter specified in [MS-WCCE] section 3.2.1.4.2.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// pdwRequestId: The pdwRequestId parameter has identical syntax and semantics to the
	// pdwRequestId parameter specified in [MS-WCCE] section 3.2.1.4.2.1.
	RequestID uint32 `idl:"name:pdwRequestId;pointer:ref" json:"request_id"`
	// pctbAttribs: A pointer to a CERTTRANSBLOB structure, as specified in [MS-WCCE] section
	// 2.2.2.2, where the pb field of this structure points to a Unicode (as specified in
	// [UNICODE4.0]) null-terminated string and the cb field contains the length of the
	// string, including the NULL-terminated character (in bytes). If the value of the cb
	// field does not match the length, in bytes, of the string (including the terminating
	// null character), the CA MUST return the E_INVALIDARG error (0x80070057) to the client.
	// Otherwise, the semantics of the string pointed to by the pb field are identical to
	// the pwszAttributes parameter specified in [MS-WCCE] section 3.2.1.4.2.1.
	Attributes *wcce.CertTransportBlob `idl:"name:pctbAttribs;pointer:ref" json:"attributes"`
	// pctbRequest: The pctbRequest parameter has identical syntax and semantics to the
	// pctbRequest parameter, as specified in [MS-WCCE] section 3.2.1.4.2.1.
	Request *wcce.CertTransportBlob `idl:"name:pctbRequest;pointer:ref" json:"request"`
}

func (o *CertServerRequestRequest) xxx_ToOp(ctx context.Context, op *xxx_CertServerRequestOperation) *xxx_CertServerRequestOperation {
	if op == nil {
		op = &xxx_CertServerRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.Flags = o.Flags
	op.Authority = o.Authority
	op.RequestID = o.RequestID
	op.Attributes = o.Attributes
	op.Request = o.Request
	return op
}

func (o *CertServerRequestRequest) xxx_FromOp(ctx context.Context, op *xxx_CertServerRequestOperation) {
	if o == nil {
		return
	}
	o.Flags = op.Flags
	o.Authority = op.Authority
	o.RequestID = op.RequestID
	o.Attributes = op.Attributes
	o.Request = op.Request
}
func (o *CertServerRequestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CertServerRequestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CertServerRequestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CertServerRequestResponse structure represents the CertServerRequest operation response
type CertServerRequestResponse struct {
	// pdwRequestId: The pdwRequestId parameter has identical syntax and semantics to the
	// pdwRequestId parameter specified in [MS-WCCE] section 3.2.1.4.2.1.
	RequestID uint32 `idl:"name:pdwRequestId;pointer:ref" json:"request_id"`
	// pdwDisposition: The pdwDisposition parameter has identical syntax and semantics to
	// the pdwDisposition parameter specified in [MS-WCCE] section 3.2.1.4.2.1.
	Disposition uint32 `idl:"name:pdwDisposition" json:"disposition"`
	// pctbCert: The pctbCert parameter has identical syntax and semantics to the pctbCertChain
	// parameter, as specified in [MS-WCCE] section 3.2.1.4.2.1.
	Cert *wcce.CertTransportBlob `idl:"name:pctbCert;pointer:ref" json:"cert"`
	// pctbEncodedCert: The pctbEncodedCert parameter has identical syntax and semantics
	// to the pctbEncodedCert parameter, as specified in [MS-WCCE] section 3.2.1.4.2.1.
	EncodedCert *wcce.CertTransportBlob `idl:"name:pctbEncodedCert;pointer:ref" json:"encoded_cert"`
	// pctbDispositionMessage: The pctbDispositionMessage parameter has identical syntax
	// and semantics to the pctbDispositionMessage parameter, as specified in [MS-WCCE]
	// section 3.2.1.4.2.1.
	DispositionMessage *wcce.CertTransportBlob `idl:"name:pctbDispositionMessage;pointer:ref" json:"disposition_message"`
	// Return: The CertServerRequest return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CertServerRequestResponse) xxx_ToOp(ctx context.Context, op *xxx_CertServerRequestOperation) *xxx_CertServerRequestOperation {
	if op == nil {
		op = &xxx_CertServerRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.RequestID = o.RequestID
	op.Disposition = o.Disposition
	op.Cert = o.Cert
	op.EncodedCert = o.EncodedCert
	op.DispositionMessage = o.DispositionMessage
	op.Return = o.Return
	return op
}

func (o *CertServerRequestResponse) xxx_FromOp(ctx context.Context, op *xxx_CertServerRequestOperation) {
	if o == nil {
		return
	}
	o.RequestID = op.RequestID
	o.Disposition = op.Disposition
	o.Cert = op.Cert
	o.EncodedCert = op.EncodedCert
	o.DispositionMessage = op.DispositionMessage
	o.Return = op.Return
}
func (o *CertServerRequestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CertServerRequestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CertServerRequestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
