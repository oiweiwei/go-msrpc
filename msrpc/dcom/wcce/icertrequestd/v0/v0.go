package icertrequestd

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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = wcce.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wcce"
)

var (
	// ICertRequestD interface identifier d99e6e70-fc88-11d0-b498-00a0c90312f3
	CertRequestDIID = &dcom.IID{Data1: 0xd99e6e70, Data2: 0xfc88, Data3: 0x11d0, Data4: []byte{0xb4, 0x98, 0x00, 0xa0, 0xc9, 0x03, 0x12, 0xf3}}
	// Syntax UUID
	CertRequestDSyntaxUUID = &uuid.UUID{TimeLow: 0xd99e6e70, TimeMid: 0xfc88, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x98, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x3, 0x12, 0xf3}}
	// Syntax ID
	CertRequestDSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CertRequestDSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICertRequestD interface.
type CertRequestDClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The Request method initiates the certificate issuance process.
	//
	// Return Values:  The method MUST return zero unless otherwise explicitly stated in
	// this section. The server MUST return errors through the pdwDisposition parameter.
	//
	// This section, and the following sections, describe the processing rules for ICertRequestD::Request
	// and ICertRequestD2::Request2.
	Request(context.Context, *RequestRequest, ...dcerpc.CallOption) (*RequestResponse, error)

	// The GetCACert method returns property values on the CA. The main use of this method
	// is to enable clients to diagnose issues and the state of the server. In addition,
	// one of the properties returned by this method is required to support the advanced
	// CA functionality (GETCERT_CAXCHGCERT).
	//
	// Return Values: For a successful invocation, the CA MUST return 0; otherwise, the
	// CA MUST return a nonzero value.
	GetCACert(context.Context, *GetCACertRequest, ...dcerpc.CallOption) (*GetCACertResponse, error)

	// The Ping method performs a request response test (ping) to the CA.
	//
	// Return Values: For a successful invocation, the CA MUST return 0; otherwise, the
	// CA MUST return a nonzero value.
	Ping(context.Context, *PingRequest, ...dcerpc.CallOption) (*PingResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CertRequestDClient
}

type xxx_DefaultCertRequestDClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCertRequestDClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCertRequestDClient) Request(ctx context.Context, in *RequestRequest, opts ...dcerpc.CallOption) (*RequestResponse, error) {
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
	out := &RequestResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertRequestDClient) GetCACert(ctx context.Context, in *GetCACertRequest, opts ...dcerpc.CallOption) (*GetCACertResponse, error) {
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
	out := &GetCACertResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertRequestDClient) Ping(ctx context.Context, in *PingRequest, opts ...dcerpc.CallOption) (*PingResponse, error) {
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
	out := &PingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertRequestDClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCertRequestDClient) IPID(ctx context.Context, ipid *dcom.IPID) CertRequestDClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCertRequestDClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewCertRequestDClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CertRequestDClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CertRequestDSyntaxV0_0))...)
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
	return &xxx_DefaultCertRequestDClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_RequestOperation structure represents the Request operation
type xxx_RequestOperation struct {
	This               *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Flags              uint32                  `idl:"name:dwFlags" json:"flags"`
	Authority          string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	RequestID          uint32                  `idl:"name:pdwRequestId;pointer:ref" json:"request_id"`
	Disposition        uint32                  `idl:"name:pdwDisposition" json:"disposition"`
	Attributes         string                  `idl:"name:pwszAttributes;string;pointer:unique" json:"attributes"`
	Request            *wcce.CertTransportBlob `idl:"name:pctbRequest;pointer:ref" json:"request"`
	CertChain          *wcce.CertTransportBlob `idl:"name:pctbCertChain;pointer:ref" json:"cert_chain"`
	EncodedCert        *wcce.CertTransportBlob `idl:"name:pctbEncodedCert;pointer:ref" json:"encoded_cert"`
	DispositionMessage *wcce.CertTransportBlob `idl:"name:pctbDispositionMessage;pointer:ref" json:"disposition_message"`
	Return             int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_RequestOperation) OpNum() int { return 3 }

func (o *xxx_RequestOperation) OpName() string { return "/ICertRequestD/v0/Request" }

func (o *xxx_RequestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.Authority) > int(1536) {
		return fmt.Errorf("Authority is out of range")
	}
	if len(o.Attributes) > int(1536) {
		return fmt.Errorf("Attributes is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique, range=(1,1536)}*(1)[dim:0,string,null](wchar))
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
	// pwszAttributes {in} (1:{string, pointer=unique, range=(1,1536)}*(1)[dim:0,string,null](wchar))
	{
		if o.Attributes != "" {
			_ptr_pwszAttributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Attributes); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Attributes, _ptr_pwszAttributes); err != nil {
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

func (o *xxx_RequestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique, range=(1,1536)}*(1)[dim:0,string,null](wchar))
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
	// pwszAttributes {in} (1:{string, pointer=unique, range=(1,1536)}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAttributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Attributes); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAttributes := func(ptr interface{}) { o.Attributes = *ptr.(*string) }
		if err := w.ReadPointer(&o.Attributes, _s_pwszAttributes, _ptr_pwszAttributes); err != nil {
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

func (o *xxx_RequestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pctbCertChain {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.CertChain != nil {
			if err := o.CertChain.MarshalNDR(ctx, w); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pctbCertChain {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.CertChain == nil {
			o.CertChain = &wcce.CertTransportBlob{}
		}
		if err := o.CertChain.UnmarshalNDR(ctx, w); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RequestRequest structure represents the Request operation request
type RequestRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwFlags: This field MUST contain packed data as specified in section 3.2.1.4.3.1.1.
	// The data in this field MUST define the structure of the pctbRequest parameter and
	// the expected content in pctbCertChain.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pwszAuthority: A null-terminated [UNICODE] string that contains the name of the CA.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// pdwRequestId: A 32-bit integer value that contains the identifier used to identify
	// the request. Details about processing information are specified in section 3.1.1.4.3.
	RequestID uint32 `idl:"name:pdwRequestId;pointer:ref" json:"request_id"`
	// pwszAttributes: A null-terminated [UNICODE] string that contains a set of request
	// attributes. The parameter contains zero or more request attributes, which MUST be
	// empty or take the form of name/value pairs. The name/value pairs MUST be formatted
	// as "Name:Value". A colon MUST be the separator, and a new line ('\n') MUST separate
	// name/value pairs.
	Attributes string `idl:"name:pwszAttributes;string;pointer:unique" json:"attributes"`
	// pctbRequest: A CERTTRANSBLOB structure that contains a certificate request as a raw
	// binary object. This request binary object can be in one of a number of formats. The
	// format used is specified in the dwFlags parameter. The syntax of that structure is
	// provided in section 2.2.2.8.
	Request *wcce.CertTransportBlob `idl:"name:pctbRequest;pointer:ref" json:"request"`
}

func (o *RequestRequest) xxx_ToOp(ctx context.Context) *xxx_RequestOperation {
	if o == nil {
		return &xxx_RequestOperation{}
	}
	return &xxx_RequestOperation{
		This:       o.This,
		Flags:      o.Flags,
		Authority:  o.Authority,
		RequestID:  o.RequestID,
		Attributes: o.Attributes,
		Request:    o.Request,
	}
}

func (o *RequestRequest) xxx_FromOp(ctx context.Context, op *xxx_RequestOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
	o.Authority = op.Authority
	o.RequestID = op.RequestID
	o.Attributes = op.Attributes
	o.Request = op.Request
}
func (o *RequestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RequestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RequestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RequestResponse structure represents the Request operation response
type RequestResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwRequestId: A 32-bit integer value that contains the identifier used to identify
	// the request. Details about processing information are specified in section 3.1.1.4.3.
	RequestID uint32 `idl:"name:pdwRequestId;pointer:ref" json:"request_id"`
	// pdwDisposition: An unsigned integer that identifies the request status for this invocation.
	// The value MUST be one of the following:
	//
	// * CR_DISP_ISSUED, 0x00000003: The requested certificate ( 719b890d-62e6-4322-b9b1-1f34d11535b4#gt_7a0f4b71-23ba-434f-b781-28053ed64879
	// ) was issued.
	//
	// * CR_DISP_UNDER_SUBMISSION, 0x00000005: The requested certificate was not issued
	// and is now in a pending state waiting for additional processing before it can be
	// issued.
	//
	// * A nonzero value, excluding 0x00000003 and 0x00000005, indicating an error.
	Disposition uint32 `idl:"name:pdwDisposition" json:"disposition"`
	// pctbCertChain: A CERTTRANSBLOB structure that is empty or contains a simple CMS or
	// a CMC full PKI response for the certificate chain issued by the CA based on the request
	// (in the pctbRequest parameter) supplied by the caller. The parameter format is as
	// requested by the client in the dwFlags parameter. Message syntax MUST be as specified
	// in section 2.2.2.2.
	CertChain *wcce.CertTransportBlob `idl:"name:pctbCertChain;pointer:ref" json:"cert_chain"`
	// pctbEncodedCert: A CERTTRANSBLOB structure that is empty or contains the issued certificate.
	// The returned value MUST be an X509 cert encoded by using DER, as specified in [X660].
	// Message syntax MUST be as specified in section 2.2.2.2.
	EncodedCert *wcce.CertTransportBlob `idl:"name:pctbEncodedCert;pointer:ref" json:"encoded_cert"`
	// pctbDispositionMessage:  A CERTTRANSBLOB structure that contains a null-terminated
	// [UNICODE] string with a message that identifies the status of the request. Message
	// syntax MUST be as specified in section 2.2.2.2.
	DispositionMessage *wcce.CertTransportBlob `idl:"name:pctbDispositionMessage;pointer:ref" json:"disposition_message"`
	// Return: The Request return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RequestResponse) xxx_ToOp(ctx context.Context) *xxx_RequestOperation {
	if o == nil {
		return &xxx_RequestOperation{}
	}
	return &xxx_RequestOperation{
		That:               o.That,
		RequestID:          o.RequestID,
		Disposition:        o.Disposition,
		CertChain:          o.CertChain,
		EncodedCert:        o.EncodedCert,
		DispositionMessage: o.DispositionMessage,
		Return:             o.Return,
	}
}

func (o *RequestResponse) xxx_FromOp(ctx context.Context, op *xxx_RequestOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RequestID = op.RequestID
	o.Disposition = op.Disposition
	o.CertChain = op.CertChain
	o.EncodedCert = op.EncodedCert
	o.DispositionMessage = op.DispositionMessage
	o.Return = op.Return
}
func (o *RequestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RequestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RequestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCACertOperation structure represents the GetCACert operation
type xxx_GetCACertOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Chain     uint32                  `idl:"name:fchain" json:"chain"`
	Authority string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Out       *wcce.CertTransportBlob `idl:"name:pctbOut;pointer:ref" json:"out"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCACertOperation) OpNum() int { return 4 }

func (o *xxx_GetCACertOperation) OpName() string { return "/ICertRequestD/v0/GetCACert" }

func (o *xxx_GetCACertOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.Authority) > int(1536) {
		return fmt.Errorf("Authority is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCACertOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fchain {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Chain); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique, range=(1,1536)}*(1)[dim:0,string,null](wchar))
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
	return nil
}

func (o *xxx_GetCACertOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fchain {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Chain); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique, range=(1,1536)}*(1)[dim:0,string,null](wchar))
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
	return nil
}

func (o *xxx_GetCACertOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCACertOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pctbOut {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Out != nil {
			if err := o.Out.MarshalNDR(ctx, w); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCACertOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pctbOut {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Out == nil {
			o.Out = &wcce.CertTransportBlob{}
		}
		if err := o.Out.UnmarshalNDR(ctx, w); err != nil {
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

// GetCACertRequest structure represents the GetCACert operation request
type GetCACertRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// fchain: Specifies the type of information to include in the output parameter.
	Chain uint32 `idl:"name:fchain" json:"chain"`
	// pwszAuthority: Contains the name of the CA.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *GetCACertRequest) xxx_ToOp(ctx context.Context) *xxx_GetCACertOperation {
	if o == nil {
		return &xxx_GetCACertOperation{}
	}
	return &xxx_GetCACertOperation{
		This:      o.This,
		Chain:     o.Chain,
		Authority: o.Authority,
	}
}

func (o *GetCACertRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCACertOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Chain = op.Chain
	o.Authority = op.Authority
}
func (o *GetCACertRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCACertRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCACertOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCACertResponse structure represents the GetCACert operation response
type GetCACertResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pctbOut:  If the function returns success (0) this parameter is a pointer to a CERTTRANSBLOB
	// structure containing the returned value.
	Out *wcce.CertTransportBlob `idl:"name:pctbOut;pointer:ref" json:"out"`
	// Return: The GetCACert return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCACertResponse) xxx_ToOp(ctx context.Context) *xxx_GetCACertOperation {
	if o == nil {
		return &xxx_GetCACertOperation{}
	}
	return &xxx_GetCACertOperation{
		That:   o.That,
		Out:    o.Out,
		Return: o.Return,
	}
}

func (o *GetCACertResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCACertOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Out = op.Out
	o.Return = op.Return
}
func (o *GetCACertResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCACertResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCACertOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PingOperation structure represents the Ping operation
type xxx_PingOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PingOperation) OpNum() int { return 5 }

func (o *xxx_PingOperation) OpName() string { return "/ICertRequestD/v0/Ping" }

func (o *xxx_PingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.Authority) > int(1536) {
		return fmt.Errorf("Authority is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszAuthority {in} (1:{string, pointer=unique, range=(1,1536)}*(1)[dim:0,string,null](wchar))
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
	return nil
}

func (o *xxx_PingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszAuthority {in} (1:{string, pointer=unique, range=(1,1536)}*(1)[dim:0,string,null](wchar))
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
	return nil
}

func (o *xxx_PingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PingRequest structure represents the Ping operation request
type PingRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: A null-terminated [UNICODE] string that MUST contain the name of the
	// CA. The CA name MUST be the CN value in the Subject field of the CA signing certificates
	// or its sanitized name. The sanitized names algorithm is specified in section 3.1.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *PingRequest) xxx_ToOp(ctx context.Context) *xxx_PingOperation {
	if o == nil {
		return &xxx_PingOperation{}
	}
	return &xxx_PingOperation{
		This:      o.This,
		Authority: o.Authority,
	}
}

func (o *PingRequest) xxx_FromOp(ctx context.Context, op *xxx_PingOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *PingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PingResponse structure represents the Ping operation response
type PingResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Ping return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PingResponse) xxx_ToOp(ctx context.Context) *xxx_PingOperation {
	if o == nil {
		return &xxx_PingOperation{}
	}
	return &xxx_PingOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PingResponse) xxx_FromOp(ctx context.Context, op *xxx_PingOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
