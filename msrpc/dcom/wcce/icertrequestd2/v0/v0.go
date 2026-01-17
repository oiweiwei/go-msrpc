package icertrequestd2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	wcce "github.com/oiweiwei/go-msrpc/msrpc/dcom/wcce"
	icertrequestd "github.com/oiweiwei/go-msrpc/msrpc/dcom/wcce/icertrequestd/v0"
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
	_ = icertrequestd.GoPackage
	_ = wcce.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wcce"
)

var (
	// ICertRequestD2 interface identifier 5422fd3a-d4b8-4cef-a12e-e87d4ca22e90
	CertRequestD2IID = &dcom.IID{Data1: 0x5422fd3a, Data2: 0xd4b8, Data3: 0x4cef, Data4: []byte{0xa1, 0x2e, 0xe8, 0x7d, 0x4c, 0xa2, 0x2e, 0x90}}
	// Syntax UUID
	CertRequestD2SyntaxUUID = &uuid.UUID{TimeLow: 0x5422fd3a, TimeMid: 0xd4b8, TimeHiAndVersion: 0x4cef, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x2e, Node: [6]uint8{0xe8, 0x7d, 0x4c, 0xa2, 0x2e, 0x90}}
	// Syntax ID
	CertRequestD2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CertRequestD2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICertRequestD2 interface.
type CertRequestD2Client interface {

	// ICertRequestD retrieval method.
	CertRequestD() icertrequestd.CertRequestDClient

	// The Request2 method requests a certificate from the CA. It is similar to the ICertRequestD::Request
	// method, but it has an additional parameter, pwszSerialNumber, which is specified
	// as follows.
	//
	// Return Values: Identical to the return value of the ICertRequestD::Request method.
	Request2(context.Context, *Request2Request, ...dcerpc.CallOption) (*Request2Response, error)

	// The GetCAProperty method retrieves a property value from the CA.
	//
	// Return Values: For a successful invocation, the CA MUST return 0; otherwise, the
	// CA MUST return a nonzero value.
	GetCAProperty(context.Context, *GetCAPropertyRequest, ...dcerpc.CallOption) (*GetCAPropertyResponse, error)

	// The GetCAPropertyInfo method retrieves a set of property structures from the CA.
	// The list of properties is specified in section 3.2.1.4.3.2.
	//
	// Return Values: For a successful invocation, the CA MUST return 0. Otherwise, the
	// CA MUST return a nonzero value.
	GetCAPropertyInfo(context.Context, *GetCAPropertyInfoRequest, ...dcerpc.CallOption) (*GetCAPropertyInfoResponse, error)

	// The Ping2 method pings the CA.
	//
	// Return Values: For a successful invocation, the CA MUST return 0; otherwise, the
	// CA MUST return a nonzero value.
	Ping2(context.Context, *Ping2Request, ...dcerpc.CallOption) (*Ping2Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CertRequestD2Client
}

type xxx_DefaultCertRequestD2Client struct {
	icertrequestd.CertRequestDClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCertRequestD2Client) CertRequestD() icertrequestd.CertRequestDClient {
	return o.CertRequestDClient
}

func (o *xxx_DefaultCertRequestD2Client) Request2(ctx context.Context, in *Request2Request, opts ...dcerpc.CallOption) (*Request2Response, error) {
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
	out := &Request2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertRequestD2Client) GetCAProperty(ctx context.Context, in *GetCAPropertyRequest, opts ...dcerpc.CallOption) (*GetCAPropertyResponse, error) {
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
	out := &GetCAPropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertRequestD2Client) GetCAPropertyInfo(ctx context.Context, in *GetCAPropertyInfoRequest, opts ...dcerpc.CallOption) (*GetCAPropertyInfoResponse, error) {
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
	out := &GetCAPropertyInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertRequestD2Client) Ping2(ctx context.Context, in *Ping2Request, opts ...dcerpc.CallOption) (*Ping2Response, error) {
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
	out := &Ping2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertRequestD2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCertRequestD2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCertRequestD2Client) IPID(ctx context.Context, ipid *dcom.IPID) CertRequestD2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCertRequestD2Client{
		CertRequestDClient: o.CertRequestDClient.IPID(ctx, ipid),
		cc:                 o.cc,
		ipid:               ipid,
	}
}

func NewCertRequestD2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CertRequestD2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CertRequestD2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := icertrequestd.NewCertRequestDClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultCertRequestD2Client{
		CertRequestDClient: base,
		cc:                 cc,
		ipid:               ipid,
	}, nil
}

// xxx_Request2Operation structure represents the Request2 operation
type xxx_Request2Operation struct {
	This               *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority          string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Flags              uint32                  `idl:"name:dwFlags" json:"flags"`
	SerialNumber       string                  `idl:"name:pwszSerialNumber;string;pointer:unique" json:"serial_number"`
	RequestID          uint32                  `idl:"name:pdwRequestId;pointer:ref" json:"request_id"`
	Disposition        uint32                  `idl:"name:pdwDisposition" json:"disposition"`
	Attributes         string                  `idl:"name:pwszAttributes;string;pointer:unique" json:"attributes"`
	Request            *wcce.CertTransportBlob `idl:"name:pctbRequest;pointer:ref" json:"request"`
	FullResponse       *wcce.CertTransportBlob `idl:"name:pctbFullResponse;pointer:ref" json:"full_response"`
	EncodedCert        *wcce.CertTransportBlob `idl:"name:pctbEncodedCert;pointer:ref" json:"encoded_cert"`
	DispositionMessage *wcce.CertTransportBlob `idl:"name:pctbDispositionMessage;pointer:ref" json:"disposition_message"`
	Return             int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_Request2Operation) OpNum() int { return 6 }

func (o *xxx_Request2Operation) OpName() string { return "/ICertRequestD2/v0/Request2" }

func (o *xxx_Request2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.Authority) > int(1536) {
		return fmt.Errorf("Authority is out of range")
	}
	if len(o.SerialNumber) > int(64) {
		return fmt.Errorf("SerialNumber is out of range")
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

func (o *xxx_Request2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pwszSerialNumber {in} (1:{string, pointer=unique, range=(1,64)}*(1)[dim:0,string,null](wchar))
	{
		if o.SerialNumber != "" {
			_ptr_pwszSerialNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SerialNumber); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SerialNumber, _ptr_pwszSerialNumber); err != nil {
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

func (o *xxx_Request2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pwszSerialNumber {in} (1:{string, pointer=unique, range=(1,64)}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszSerialNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SerialNumber); err != nil {
				return err
			}
			return nil
		})
		_s_pwszSerialNumber := func(ptr interface{}) { o.SerialNumber = *ptr.(*string) }
		if err := w.ReadPointer(&o.SerialNumber, _s_pwszSerialNumber, _ptr_pwszSerialNumber); err != nil {
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

func (o *xxx_Request2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Request2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pctbFullResponse {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.FullResponse != nil {
			if err := o.FullResponse.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_Request2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pctbFullResponse {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.FullResponse == nil {
			o.FullResponse = &wcce.CertTransportBlob{}
		}
		if err := o.FullResponse.UnmarshalNDR(ctx, w); err != nil {
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

// Request2Request structure represents the Request2 operation request
type Request2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: Identical to the pwszAuthority parameter in the ICertRequestD::Request
	// method.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// dwFlags: Identical to the dwFlags parameter in the ICertRequestD::Request method.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pwszSerialNumber: A null-terminated [UNICODE] string that specifies a serial number
	// that identifies a certificate. The string MUST specify the serial number as an even
	// number of hexadecimal digits. If necessary, a zero can be prefixed to the number
	// to produce an even number of digits. The string MUST NOT contain more than one leading
	// zero. Information on the serial number is specified in [RFC3280] section 4.1.2.2.
	SerialNumber string `idl:"name:pwszSerialNumber;string;pointer:unique" json:"serial_number"`
	// pdwRequestId: Identical to the pdwRequestId parameter in the ICertRequestD::Request
	// method.
	RequestID uint32 `idl:"name:pdwRequestId;pointer:ref" json:"request_id"`
	// pwszAttributes: Identical to the pwszAttributes parameter in the ICertRequestD::Request
	// method.
	Attributes string `idl:"name:pwszAttributes;string;pointer:unique" json:"attributes"`
	// pctbRequest: Identical to the pctbRequest parameter in the ICertRequestD::Request
	// method.
	Request *wcce.CertTransportBlob `idl:"name:pctbRequest;pointer:ref" json:"request"`
}

func (o *Request2Request) xxx_ToOp(ctx context.Context, op *xxx_Request2Operation) *xxx_Request2Operation {
	if op == nil {
		op = &xxx_Request2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Authority = o.Authority
	op.Flags = o.Flags
	op.SerialNumber = o.SerialNumber
	op.RequestID = o.RequestID
	op.Attributes = o.Attributes
	op.Request = o.Request
	return op
}

func (o *Request2Request) xxx_FromOp(ctx context.Context, op *xxx_Request2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.Flags = op.Flags
	o.SerialNumber = op.SerialNumber
	o.RequestID = op.RequestID
	o.Attributes = op.Attributes
	o.Request = op.Request
}
func (o *Request2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *Request2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Request2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Request2Response structure represents the Request2 operation response
type Request2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwRequestId: Identical to the pdwRequestId parameter in the ICertRequestD::Request
	// method.
	RequestID uint32 `idl:"name:pdwRequestId;pointer:ref" json:"request_id"`
	// pdwDisposition:  Identical to the pdwDisposition parameter in the ICertRequestD::Request
	// method.
	Disposition uint32 `idl:"name:pdwDisposition" json:"disposition"`
	// pctbFullResponse: Identical to the pctbCertChain parameter in the ICertRequestD::Request
	// method.
	FullResponse *wcce.CertTransportBlob `idl:"name:pctbFullResponse;pointer:ref" json:"full_response"`
	// pctbEncodedCert: Identical to the pctbEncodedCert parameter in the ICertRequestD::Request
	// method.
	EncodedCert *wcce.CertTransportBlob `idl:"name:pctbEncodedCert;pointer:ref" json:"encoded_cert"`
	// pctbDispositionMessage: Identical to the pctbDispositionMessage parameter in the
	// ICertRequestD::Request method.
	DispositionMessage *wcce.CertTransportBlob `idl:"name:pctbDispositionMessage;pointer:ref" json:"disposition_message"`
	// Return: The Request2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Request2Response) xxx_ToOp(ctx context.Context, op *xxx_Request2Operation) *xxx_Request2Operation {
	if op == nil {
		op = &xxx_Request2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.RequestID = o.RequestID
	op.Disposition = o.Disposition
	op.FullResponse = o.FullResponse
	op.EncodedCert = o.EncodedCert
	op.DispositionMessage = o.DispositionMessage
	op.Return = o.Return
	return op
}

func (o *Request2Response) xxx_FromOp(ctx context.Context, op *xxx_Request2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RequestID = op.RequestID
	o.Disposition = op.Disposition
	o.FullResponse = op.FullResponse
	o.EncodedCert = op.EncodedCert
	o.DispositionMessage = op.DispositionMessage
	o.Return = op.Return
}
func (o *Request2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *Request2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Request2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCAPropertyOperation structure represents the GetCAProperty operation
type xxx_GetCAPropertyOperation struct {
	This          *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority     string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	PropertyID    int32                   `idl:"name:PropID" json:"property_id"`
	PropertyIndex int32                   `idl:"name:PropIndex" json:"property_index"`
	PropertyType  int32                   `idl:"name:PropType" json:"property_type"`
	PropertyValue *wcce.CertTransportBlob `idl:"name:pctbPropertyValue;pointer:ref" json:"property_value"`
	Return        int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCAPropertyOperation) OpNum() int { return 7 }

func (o *xxx_GetCAPropertyOperation) OpName() string { return "/ICertRequestD2/v0/GetCAProperty" }

func (o *xxx_GetCAPropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
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

func (o *xxx_GetCAPropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// PropID {in} (1:(int32))
	{
		if err := w.WriteData(o.PropertyID); err != nil {
			return err
		}
	}
	// PropIndex {in} (1:(int32))
	{
		if err := w.WriteData(o.PropertyIndex); err != nil {
			return err
		}
	}
	// PropType {in} (1:(int32))
	{
		if err := w.WriteData(o.PropertyType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAPropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// PropID {in} (1:(int32))
	{
		if err := w.ReadData(&o.PropertyID); err != nil {
			return err
		}
	}
	// PropIndex {in} (1:(int32))
	{
		if err := w.ReadData(&o.PropertyIndex); err != nil {
			return err
		}
	}
	// PropType {in} (1:(int32))
	{
		if err := w.ReadData(&o.PropertyType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAPropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAPropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pctbPropertyValue {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.PropertyValue != nil {
			if err := o.PropertyValue.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetCAPropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pctbPropertyValue {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.PropertyValue == nil {
			o.PropertyValue = &wcce.CertTransportBlob{}
		}
		if err := o.PropertyValue.UnmarshalNDR(ctx, w); err != nil {
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

// GetCAPropertyRequest structure represents the GetCAProperty operation request
type GetCAPropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: Contains the name of the CA.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// PropID:  An integer value that specifies the property to be returned.
	//
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	|             PROPERTY             |    NUMERICAL    |                 |                                                                                  |
	//	|               NAME               |      VALUE      |   TYPE/INDEX    |                                     MEANING                                      |
	//	|                                  |                 |                 |                                                                                  |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_FILEVERSION              | 0x00000001      | String          | A string that MUST contain the CA version information.                           |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_PRODUCTVERSION           | 0x00000002      | String          | A string that MUST contain the build number of the CA.                           |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_EXITCOUNT                | 0x00000003      | Long            | MUST be the number of exit algorithms registered on the CA.                      |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_EXITDESCRIPTION          | 0x00000004      | String indexed  | A string that MUST contain the name of the exit algorithm identified by the      |
	//	|                                  |                 |                 | PropIndex parameter.                                                             |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_POLICYDESCRIPTION        | 0x00000005      | String          | A string that MUST contain the description of the policy algorithm on the CA.    |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CANAME                   | 0x00000006      | String          | A string that MUST contain the CN, as specified in [RFC3280], of a CA.           |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_SANITIZEDCANAME          | 0x00000007      | String          | A string that MUST contain the sanitized name of the CA. More information about  |
	//	|                                  |                 |                 | sanitized name is specified in section 3.1.1.4.1.1.                              |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_SHAREDFOLDER             | 0x00000008      | String          | A string that MUST contain the UNC path of a folder that contains the CA         |
	//	|                                  |                 |                 | information and signature certificates.                                          |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_PARENTCA                 | 0x00000009      | String          | A string that MUST contain the name of the parent CA to the current CA.          |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CATYPE                   | 0x0000000A      | Long            | MUST be a CAINFO structure that MUST contain the CA type. More information is    |
	//	|                                  |                 |                 | specified in section 3.2.1.4.3.2.10.                                             |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CASIGCERTCOUNT           | 0x0000000B      | Long            | MUST be the number of signing certificates on the CA.                            |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CASIGCERT                | 0x0000000C      | Binary, indexed | MUST be a binary object that contains a signing certificate identified by the    |
	//	|                                  |                 |                 | PropIndex parameter.                                                             |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CASIGCERTCHAIN           | 0x0000000D      | Binary, indexed | MUST be a binary object that contains the certificate chain for a signing        |
	//	|                                  |                 |                 | certificate identified by the PropIndex parameter.                               |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CAXCHGCERTCOUNT          | 0x0000000E      | Long            | MUST be 0x1.                                                                     |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CAXCHGCERT               | 0x0000000F      | Binary, indexed | MUST be a binary object that contains the CA's current exchange certificate      |
	//	|                                  |                 |                 | from the Current_CA_Exchange_Cert datum. The PropIndex parameter MUST be 0x0 or  |
	//	|                                  |                 |                 | 0xFFFFFFFF.                                                                      |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CAXCHGCERTCHAIN          | 0x00000010      | Binary, indexed | MUST be a binary object that contains the certificate chain for the CA's current |
	//	|                                  |                 |                 | exchange certificate from the Current_CA_Exchange_Cert datum. The PropIndex      |
	//	|                                  |                 |                 | parameter MUST be 0x0 or 0xFFFFFFFF.                                             |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_BASECRL                  | 0x00000011      | Binary, indexed | MUST be a CRL, for a CA signing certificate identified by the PropIndex          |
	//	|                                  |                 |                 | parameter.                                                                       |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_DELTACRL                 | 0x00000012      | Binary, indexed | MUST be a delta CRL, for a CA signing certificate identified by the PropIndex    |
	//	|                                  |                 |                 | parameter. For more information about delta CRLs, see [MSFT-CRL]. Additional     |
	//	|                                  |                 |                 | information is specified in [RFC3280] section 5.2.                               |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CACERTSTATE              | 0x00000013      | Long indexed    | MUST be a byte array that contains the disposition status of all CA signing      |
	//	|                                  |                 |                 | certificates. Disposition status is specified in section 3.2.1.4.3.2.19.         |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CRLSTATE                 | 0x00000014      | Long indexed    | MUST be a byte array that contains the status for all the CRLs of the CA.        |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CAPROPIDMAX              | 0x00000015      | Long            | MUST be the maximum property identifier supported by the CA.                     |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_DNSNAME                  | 0x00000016      | String          | MUST be the fully qualified domain name (FQDN) of the computer on which the CA   |
	//	|                                  |                 |                 | is installed.                                                                    |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_ROLESEPARATIONENABLED    | 0x00000017      | Long            | Indicates whether administrative role separation has been enabled on the CA.     |
	//	|                                  |                 |                 | A nonzero return value means that role separation has been enabled. Zero means   |
	//	|                                  |                 |                 | that role separation has not been enabled.                                       |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_KRACERTUSEDCOUNT         | 0x00000018      | Long            | MUST be the minimum number of KRAs to use when archiving a private key. For more |
	//	|                                  |                 |                 | information about KRA usage, see [MSFT-ARCHIVE].                                 |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_KRACERTCOUNT             | 0x00000019      | Long            | MUST be the maximum number of KRA certificates available on the CA.              |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_KRACERT                  | 0x0000001A      | Binary, indexed | A KRA certificate identified by the PropIndex parameter.                         |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_KRACERTSTATE             | 0x0000001B      | Long, indexed   | MUST be a byte array that contains the status of the KRA certificates registered |
	//	|                                  |                 |                 | with the CA.                                                                     |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_ADVANCEDSERVER           | 0x0000001C      | Long            | MUST identify whether the CA operating system is an advanced server platform.    |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_TEMPLATES                | 0x0000001D      | String          | MUST be a collection of name and OID pairs that identify the templates supported |
	//	|                                  |                 |                 | by a CA.                                                                         |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_BASECRLPUBLISHSTATUS     | 0x0000001E      | Long, indexed   | MUST be the publishing status of a signing certificate base CRL identified by    |
	//	|                                  |                 |                 | the PropIndex parameter.                                                         |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_DELTACRLPUBLISHSTATUS    | 0x0000001F      | Long, indexed   | MUST be the publishing status of a signing certificate delta CRL identified by   |
	//	|                                  |                 |                 | the PropIndex parameter.                                                         |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CASIGCERTCRLCHAIN        | 0x00000020      | Binary, indexed | MUST be a binary object that contains the certificate chain for a signing        |
	//	|                                  |                 |                 | certificate and the CRL for the certificates in the chain identified by the      |
	//	|                                  |                 |                 | PropIndex parameter.                                                             |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CAXCHGCERTCRLCHAIN       | 0x00000021      | Binary, indexed | MUST be a binary object for a chain containing CRLs for the CA's current         |
	//	|                                  |                 |                 | exchange certificate from the Current_CA_Exchange_Cert datum. The PropIndex      |
	//	|                                  |                 |                 | parameter MUST be 0x00000000 or 0xFFFFFFFF.                                      |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CACERTSTATUSCODE         | 0x00000022      | Long, indexed   | MUST be an HRESULT that identifies the result of certificate validation, as      |
	//	|                                  |                 |                 | specified in [RFC3280], by the CA for the CA signing certificates identified by  |
	//	|                                  |                 |                 | the PropIndex parameter.                                                         |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CAFORWARDCROSSCERT       | 0x00000023      | Binary, indexed | MUST be a forward cross certificate, by index, from a CA. For more information   |
	//	|                                  |                 |                 | about cross certificates, see [MSFT-CROSSCERT].                                  |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CABACKWARDCROSSCERT      | 0x00000024      | Binary, indexed | MUST be a backward cross certificate, by index, from a CA. For more information  |
	//	|                                  |                 |                 | about cross certificates.                                                        |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CAFORWARDCROSSCERTSTATE  | 0x00000025      | Long, indexed   | MUST be a byte array that identifies the status of all backward cross            |
	//	|                                  |                 |                 | certificates for a CA.                                                           |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CABACKWARDCROSSCERTSTATE | 0x00000026      | Long, indexed   | MUST be a byte array that identifies the disposition status of all forward cross |
	//	|                                  |                 |                 | certificates for a CA.                                                           |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CACERTVERSION            | 0x00000027      | Long, indexed   | MUST be an indexed 32-bit integer that contains the version number of a CA       |
	//	|                                  |                 |                 | signing certificate.                                                             |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_SANITIZEDCASHORTNAME     | 0x00000028      | String          | The property MUST return the sanitized shortened name of the CA. More            |
	//	|                                  |                 |                 | information about the sanitized name is specified in section 3.1.1.4.1.1.        |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CERTCDPURLS              | 0x00000029      | String, indexed | MUST be a null-terminated [UNICODE] string of the format "String1\nString2\n",   |
	//	|                                  |                 |                 | where each string (separated by '\n') MUST represent a URI to be part of a CRL   |
	//	|                                  |                 |                 | Distribution Point (CDP) extension, as specified in [RFC3280] section 4.2.1.14.  |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CERTAIAURLS              | 0x0000002A      | String, indexed | MUST be a null-terminated [UNICODE] string of the format "String1\nString2\n",   |
	//	|                                  |                 |                 | where each string (separated by '\n') MUST represent a URI to be part of         |
	//	|                                  |                 |                 | Authority Information Access extension, as specified in [RFC3280] section        |
	//	|                                  |                 |                 | 4.2.2.1.                                                                         |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_CERTAIAOCSPRLS           | 0x0000002B      | String, indexed | MUST be a null-terminated [UNICODE] string of the format "String1\nString2\n",   |
	//	|                                  |                 |                 | where each string (separated by '\n') MUST represent the OCSP URLs configured on |
	//	|                                  |                 |                 | the CA, as specified in [RFC3280] section 4.2.2.1.                               |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_LOCALENAME               | 0x0000002C      | String          | MUST be a null-terminated [UNICODE] string in the 'Language-Region' format (as   |
	//	|                                  |                 |                 | specified in [RFC4646]) that represents the locale of the CA.                    |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_SUBJECTTEMPLATE_OIDS     | 0x0000002D      | String          | MUST be a null-terminated [UNICODE] string of the format "OID1\nOID2\n", where   |
	//	|                                  |                 |                 | each OID (separated by '\n') MUST represent a Relative Distinguished Name that   |
	//	|                                  |                 |                 | is in a certificate Subject Distinguished Name.                                  |
	//	+----------------------------------+-----------------+-----------------+----------------------------------------------------------------------------------+
	PropertyID int32 `idl:"name:PropID" json:"property_id"`
	// PropIndex: This parameter is used as the index to a property that can contain multiple
	// values.
	PropertyIndex int32 `idl:"name:PropIndex" json:"property_index"`
	// PropType: An integer value that specifies the property data type.
	//
	//	+----------------------------+-------------------------------------------------------------+
	//	|                            |                                                             |
	//	|           VALUE            |                           MEANING                           |
	//	|                            |                                                             |
	//	+----------------------------+-------------------------------------------------------------+
	//	+----------------------------+-------------------------------------------------------------+
	//	| PROPTYPE_LONG 0x00000001   | The property type is a signed long integer or a byte array. |
	//	+----------------------------+-------------------------------------------------------------+
	//	| PROPTYPE_BINARY 0x00000003 | The property type is binary data.                           |
	//	+----------------------------+-------------------------------------------------------------+
	//	| PROPTYPE_STRING 0x00000004 | The property type is a string.                              |
	//	+----------------------------+-------------------------------------------------------------+
	PropertyType int32 `idl:"name:PropType" json:"property_type"`
}

func (o *GetCAPropertyRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCAPropertyOperation) *xxx_GetCAPropertyOperation {
	if op == nil {
		op = &xxx_GetCAPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Authority = o.Authority
	op.PropertyID = o.PropertyID
	op.PropertyIndex = o.PropertyIndex
	op.PropertyType = o.PropertyType
	return op
}

func (o *GetCAPropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCAPropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.PropertyID = op.PropertyID
	o.PropertyIndex = op.PropertyIndex
	o.PropertyType = op.PropertyType
}
func (o *GetCAPropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCAPropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCAPropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCAPropertyResponse structure represents the GetCAProperty operation response
type GetCAPropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pctbPropertyValue: If the function succeeds, this method returns a CERTTRANSBLOB
	// structure in this parameter that contains the property value. If the function fails,
	// the content of this parameter is undefined.
	PropertyValue *wcce.CertTransportBlob `idl:"name:pctbPropertyValue;pointer:ref" json:"property_value"`
	// Return: The GetCAProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCAPropertyResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCAPropertyOperation) *xxx_GetCAPropertyOperation {
	if op == nil {
		op = &xxx_GetCAPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PropertyValue = o.PropertyValue
	op.Return = o.Return
	return op
}

func (o *GetCAPropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCAPropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertyValue = op.PropertyValue
	o.Return = op.Return
}
func (o *GetCAPropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCAPropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCAPropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCAPropertyInfoOperation structure represents the GetCAPropertyInfo operation
type xxx_GetCAPropertyInfoOperation struct {
	This          *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority     string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	PropertyCount int32                   `idl:"name:pcProperty" json:"property_count"`
	PropertyInfo  *wcce.CertTransportBlob `idl:"name:pctbPropInfo;pointer:ref" json:"property_info"`
	Return        int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCAPropertyInfoOperation) OpNum() int { return 8 }

func (o *xxx_GetCAPropertyInfoOperation) OpName() string {
	return "/ICertRequestD2/v0/GetCAPropertyInfo"
}

func (o *xxx_GetCAPropertyInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
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

func (o *xxx_GetCAPropertyInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCAPropertyInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCAPropertyInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAPropertyInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcProperty {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.PropertyCount); err != nil {
			return err
		}
	}
	// pctbPropInfo {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.PropertyInfo != nil {
			if err := o.PropertyInfo.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetCAPropertyInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcProperty {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.PropertyCount); err != nil {
			return err
		}
	}
	// pctbPropInfo {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.PropertyInfo == nil {
			o.PropertyInfo = &wcce.CertTransportBlob{}
		}
		if err := o.PropertyInfo.UnmarshalNDR(ctx, w); err != nil {
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

// GetCAPropertyInfoRequest structure represents the GetCAPropertyInfo operation request
type GetCAPropertyInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: Contains the name of the CA.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *GetCAPropertyInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCAPropertyInfoOperation) *xxx_GetCAPropertyInfoOperation {
	if op == nil {
		op = &xxx_GetCAPropertyInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Authority = o.Authority
	return op
}

func (o *GetCAPropertyInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCAPropertyInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *GetCAPropertyInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCAPropertyInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCAPropertyInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCAPropertyInfoResponse structure represents the GetCAPropertyInfo operation response
type GetCAPropertyInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pcProperty: An integer value that contains the number of property structures returned.
	PropertyCount int32 `idl:"name:pcProperty" json:"property_count"`
	// pctbPropInfo: A CERTTRANSBLOB structure that contains zero or more CATRANSPROP structures.
	// For more information about the CERTTRANSBLOB and CATRANSPROP structures, see Common
	// Structures.
	PropertyInfo *wcce.CertTransportBlob `idl:"name:pctbPropInfo;pointer:ref" json:"property_info"`
	// Return: The GetCAPropertyInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCAPropertyInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCAPropertyInfoOperation) *xxx_GetCAPropertyInfoOperation {
	if op == nil {
		op = &xxx_GetCAPropertyInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PropertyCount = o.PropertyCount
	op.PropertyInfo = o.PropertyInfo
	op.Return = o.Return
	return op
}

func (o *GetCAPropertyInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCAPropertyInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertyCount = op.PropertyCount
	o.PropertyInfo = op.PropertyInfo
	o.Return = op.Return
}
func (o *GetCAPropertyInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCAPropertyInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCAPropertyInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Ping2Operation structure represents the Ping2 operation
type xxx_Ping2Operation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_Ping2Operation) OpNum() int { return 9 }

func (o *xxx_Ping2Operation) OpName() string { return "/ICertRequestD2/v0/Ping2" }

func (o *xxx_Ping2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
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

func (o *xxx_Ping2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_Ping2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_Ping2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Ping2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_Ping2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// Ping2Request structure represents the Ping2 operation request
type Ping2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: Contains the name of the CA.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *Ping2Request) xxx_ToOp(ctx context.Context, op *xxx_Ping2Operation) *xxx_Ping2Operation {
	if op == nil {
		op = &xxx_Ping2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Authority = o.Authority
	return op
}

func (o *Ping2Request) xxx_FromOp(ctx context.Context, op *xxx_Ping2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *Ping2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *Ping2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Ping2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Ping2Response structure represents the Ping2 operation response
type Ping2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Ping2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Ping2Response) xxx_ToOp(ctx context.Context, op *xxx_Ping2Operation) *xxx_Ping2Operation {
	if op == nil {
		op = &xxx_Ping2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *Ping2Response) xxx_FromOp(ctx context.Context, op *xxx_Ping2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *Ping2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *Ping2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Ping2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
