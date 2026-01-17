package imsmqapplication2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	imsmqapplication "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqapplication/v0"
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
	_ = dcom.GoPackage
	_ = imsmqapplication.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQApplication2 interface identifier 12a30900-7300-11d2-b0e6-00e02c074f6b
	Application2IID = &dcom.IID{Data1: 0x12a30900, Data2: 0x7300, Data3: 0x11d2, Data4: []byte{0xb0, 0xe6, 0x00, 0xe0, 0x2c, 0x07, 0x4f, 0x6b}}
	// Syntax UUID
	Application2SyntaxUUID = &uuid.UUID{TimeLow: 0x12a30900, TimeMid: 0x7300, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0xe6, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	// Syntax ID
	Application2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Application2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQApplication2 interface.
type Application2Client interface {

	// IMSMQApplication retrieval method.
	Application() imsmqapplication.ApplicationClient

	// The RegisterCertificate method MUST register an MQUSERSIGNCERT ([MS-MQMQ] section
	// 2.2.22) in User.Certificates. Implementations of this protocol use certificates to
	// verify the sender for messages that are requesting authentication and to ensure message
	// integrity.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<10>
	RegisterCertificate(context.Context, *RegisterCertificateRequest, ...dcerpc.CallOption) (*RegisterCertificateResponse, error)

	// The MachineNameOfMachineId method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return a string that contains the QueueManager.ComputerName
	// of the QueueManager that is identified by the bstrGuid input parameter.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<12>
	MachineNameOfMachineID(context.Context, *MachineNameOfMachineIDRequest, ...dcerpc.CallOption) (*MachineNameOfMachineIDResponse, error)

	// The MSMQVersionMajor method is received by the server in an RPC_REQUEST packet. In
	// response, the serverMUST return the major version number of the server.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetVersionMajor(context.Context, *GetVersionMajorRequest, ...dcerpc.CallOption) (*GetVersionMajorResponse, error)

	// The MSMQVersionMinor method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return the minor version number of the server.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the *ComputerName* instance variable is not NULL, return MQ_ERROR_INVALID_PARAMETER
	// (0xC00E0006) and take no further action.
	//
	// * Set the psMSMQVersionMinor output variable to the minor version number <14> ( 71c359c3-e9ec-4fe6-a101-aab1eabecdcf#Appendix_A_14
	// ) of the server.
	GetVersionMinor(context.Context, *GetVersionMinorRequest, ...dcerpc.CallOption) (*GetVersionMinorResponse, error)

	// The MSMQVersionBuild method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return the build version number of the server.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetVersionBuild(context.Context, *GetVersionBuildRequest, ...dcerpc.CallOption) (*GetVersionBuildResponse, error)

	// The IsDsEnabled method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return a BOOLEAN value that indicates whether the local QueueManager
	// is configured to use the directory.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetIsDSEnabled(context.Context, *GetIsDSEnabledRequest, ...dcerpc.CallOption) (*GetIsDSEnabledResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Application2Client
}

type xxx_DefaultApplication2Client struct {
	imsmqapplication.ApplicationClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultApplication2Client) Application() imsmqapplication.ApplicationClient {
	return o.ApplicationClient
}

func (o *xxx_DefaultApplication2Client) RegisterCertificate(ctx context.Context, in *RegisterCertificateRequest, opts ...dcerpc.CallOption) (*RegisterCertificateResponse, error) {
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
	out := &RegisterCertificateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultApplication2Client) MachineNameOfMachineID(ctx context.Context, in *MachineNameOfMachineIDRequest, opts ...dcerpc.CallOption) (*MachineNameOfMachineIDResponse, error) {
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
	out := &MachineNameOfMachineIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultApplication2Client) GetVersionMajor(ctx context.Context, in *GetVersionMajorRequest, opts ...dcerpc.CallOption) (*GetVersionMajorResponse, error) {
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
	out := &GetVersionMajorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultApplication2Client) GetVersionMinor(ctx context.Context, in *GetVersionMinorRequest, opts ...dcerpc.CallOption) (*GetVersionMinorResponse, error) {
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
	out := &GetVersionMinorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultApplication2Client) GetVersionBuild(ctx context.Context, in *GetVersionBuildRequest, opts ...dcerpc.CallOption) (*GetVersionBuildResponse, error) {
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
	out := &GetVersionBuildResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultApplication2Client) GetIsDSEnabled(ctx context.Context, in *GetIsDSEnabledRequest, opts ...dcerpc.CallOption) (*GetIsDSEnabledResponse, error) {
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
	out := &GetIsDSEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultApplication2Client) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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
	out := &GetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultApplication2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultApplication2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultApplication2Client) IPID(ctx context.Context, ipid *dcom.IPID) Application2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultApplication2Client{
		ApplicationClient: o.ApplicationClient.IPID(ctx, ipid),
		cc:                o.cc,
		ipid:              ipid,
	}
}

func NewApplication2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Application2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Application2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := imsmqapplication.NewApplicationClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultApplication2Client{
		ApplicationClient: base,
		cc:                cc,
		ipid:              ipid,
	}, nil
}

// xxx_RegisterCertificateOperation structure represents the RegisterCertificate operation
type xxx_RegisterCertificateOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	Flags               *oaut.Variant  `idl:"name:Flags" json:"flags"`
	ExternalCertificate *oaut.Variant  `idl:"name:ExternalCertificate" json:"external_certificate"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterCertificateOperation) OpNum() int { return 8 }

func (o *xxx_RegisterCertificateOperation) OpName() string {
	return "/IMSMQApplication2/v0/RegisterCertificate"
}

func (o *xxx_RegisterCertificateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterCertificateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Flags {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Flags != nil {
			_ptr_Flags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Flags != nil {
					if err := o.Flags.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Flags, _ptr_Flags); err != nil {
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
	// ExternalCertificate {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ExternalCertificate != nil {
			_ptr_ExternalCertificate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ExternalCertificate != nil {
					if err := o.ExternalCertificate.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExternalCertificate, _ptr_ExternalCertificate); err != nil {
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

func (o *xxx_RegisterCertificateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Flags {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Flags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Flags == nil {
				o.Flags = &oaut.Variant{}
			}
			if err := o.Flags.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Flags := func(ptr interface{}) { o.Flags = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Flags, _s_Flags, _ptr_Flags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ExternalCertificate {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ExternalCertificate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ExternalCertificate == nil {
				o.ExternalCertificate = &oaut.Variant{}
			}
			if err := o.ExternalCertificate.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ExternalCertificate := func(ptr interface{}) { o.ExternalCertificate = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ExternalCertificate, _s_ExternalCertificate, _ptr_ExternalCertificate); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterCertificateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterCertificateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RegisterCertificateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RegisterCertificateRequest structure represents the RegisterCertificate operation request
type RegisterCertificateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Flags: A pointer to a VARIANT that contains a VT_I4 integer that corresponds to the
	// MQCERT_REGISTER enumeration as defined in the following table.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| MQCERT_REGISTER_ALWAYS 0x00000001       | Register an MQUSERSIGNCERT ([MS-MQMQ] section 2.2.22) in User.Certificates. If   |
	//	|                                         | the ExternalCertificate input parameter is not specified or is NULL, the server  |
	//	|                                         | MUST delete the certificate from the internal store and delete any existing      |
	//	|                                         | MQUSERSIGNCERT ([MS-MQMQ] section 2.2.22) with a matching Digest property from   |
	//	|                                         | User.Certificates. The server MUST then add a newly created MQUSERSIGNCERT       |
	//	|                                         | ([MS-MQMQ] section 2.2.22) to User.Certificates. If the ExternalCertificate is   |
	//	|                                         | not NULL, the server MUST add an MQUSERSIGNCERT ([MS-MQMQ] section 2.2.22), as   |
	//	|                                         | specified by the ExternalCertificate input parameter, to User.Certificates.      |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| MQCERT_REGISTER_IF_NOT_EXIST 0x00000002 | Register an MQUSERSIGNCERT ([MS-MQMQ] section 2.2.22) in User.Certificates only  |
	//	|                                         | if no certificate is registered in the internal store. This option cannot be     |
	//	|                                         | used with ExternalCertificate.                                                   |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	Flags *oaut.Variant `idl:"name:Flags" json:"flags"`
	// ExternalCertificate: A pointer to a VARIANT that contains a byte array (VT_ARRAY|VT_UI1) or a pointer (VT_BYREF) to a byte array that specifies the binary representation of the MQUSERSIGNCERT ([MS-MQMQ] section 2.2.22) that is to be registered. The MQUSERSIGNCERT MUST contain an X.509-encoded certificate, as specified in [RFC3280].
	ExternalCertificate *oaut.Variant `idl:"name:ExternalCertificate" json:"external_certificate"`
}

func (o *RegisterCertificateRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterCertificateOperation) *xxx_RegisterCertificateOperation {
	if op == nil {
		op = &xxx_RegisterCertificateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Flags = o.Flags
	op.ExternalCertificate = o.ExternalCertificate
	return op
}

func (o *RegisterCertificateRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterCertificateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
	o.ExternalCertificate = op.ExternalCertificate
}
func (o *RegisterCertificateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterCertificateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterCertificateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterCertificateResponse structure represents the RegisterCertificate operation response
type RegisterCertificateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RegisterCertificate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterCertificateResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterCertificateOperation) *xxx_RegisterCertificateOperation {
	if op == nil {
		op = &xxx_RegisterCertificateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RegisterCertificateResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterCertificateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RegisterCertificateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterCertificateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterCertificateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MachineNameOfMachineIDOperation structure represents the MachineNameOfMachineId operation
type xxx_MachineNameOfMachineIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	GUID        *oaut.String   `idl:"name:bstrGuid" json:"guid"`
	MachineName *oaut.String   `idl:"name:pbstrMachineName" json:"machine_name"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_MachineNameOfMachineIDOperation) OpNum() int { return 9 }

func (o *xxx_MachineNameOfMachineIDOperation) OpName() string {
	return "/IMSMQApplication2/v0/MachineNameOfMachineId"
}

func (o *xxx_MachineNameOfMachineIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MachineNameOfMachineIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrGuid {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.GUID != nil {
			_ptr_bstrGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.GUID != nil {
					if err := o.GUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.GUID, _ptr_bstrGuid); err != nil {
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

func (o *xxx_MachineNameOfMachineIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrGuid {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.GUID == nil {
				o.GUID = &oaut.String{}
			}
			if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrGuid := func(ptr interface{}) { o.GUID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.GUID, _s_bstrGuid, _ptr_bstrGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MachineNameOfMachineIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MachineNameOfMachineIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrMachineName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineName != nil {
			_ptr_pbstrMachineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineName != nil {
					if err := o.MachineName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineName, _ptr_pbstrMachineName); err != nil {
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

func (o *xxx_MachineNameOfMachineIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrMachineName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrMachineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineName == nil {
				o.MachineName = &oaut.String{}
			}
			if err := o.MachineName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrMachineName := func(ptr interface{}) { o.MachineName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineName, _s_pbstrMachineName, _ptr_pbstrMachineName); err != nil {
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

// MachineNameOfMachineIDRequest structure represents the MachineNameOfMachineId operation request
type MachineNameOfMachineIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrGuid: A BSTR that specifies an identifier for the computer name, in the GUID
	// format. The server MUST accept GUIDs with or without surrounding curly braces ({
	// }).
	GUID *oaut.String `idl:"name:bstrGuid" json:"guid"`
}

func (o *MachineNameOfMachineIDRequest) xxx_ToOp(ctx context.Context, op *xxx_MachineNameOfMachineIDOperation) *xxx_MachineNameOfMachineIDOperation {
	if op == nil {
		op = &xxx_MachineNameOfMachineIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.GUID = o.GUID
	return op
}

func (o *MachineNameOfMachineIDRequest) xxx_FromOp(ctx context.Context, op *xxx_MachineNameOfMachineIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.GUID = op.GUID
}
func (o *MachineNameOfMachineIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MachineNameOfMachineIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MachineNameOfMachineIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MachineNameOfMachineIDResponse structure represents the MachineNameOfMachineId operation response
type MachineNameOfMachineIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrMachineName: A pointer to a BSTR that after successful completion, contains
	// the computer name in the DNS or Universal Naming Convention (UNC) format.
	MachineName *oaut.String `idl:"name:pbstrMachineName" json:"machine_name"`
	// Return: The MachineNameOfMachineId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MachineNameOfMachineIDResponse) xxx_ToOp(ctx context.Context, op *xxx_MachineNameOfMachineIDOperation) *xxx_MachineNameOfMachineIDOperation {
	if op == nil {
		op = &xxx_MachineNameOfMachineIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MachineName = o.MachineName
	op.Return = o.Return
	return op
}

func (o *MachineNameOfMachineIDResponse) xxx_FromOp(ctx context.Context, op *xxx_MachineNameOfMachineIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MachineName = op.MachineName
	o.Return = op.Return
}
func (o *MachineNameOfMachineIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MachineNameOfMachineIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MachineNameOfMachineIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVersionMajorOperation structure represents the MSMQVersionMajor operation
type xxx_GetVersionMajorOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	VersionMajor int16          `idl:"name:psMSMQVersionMajor" json:"version_major"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVersionMajorOperation) OpNum() int { return 10 }

func (o *xxx_GetVersionMajorOperation) OpName() string {
	return "/IMSMQApplication2/v0/MSMQVersionMajor"
}

func (o *xxx_GetVersionMajorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionMajorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetVersionMajorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetVersionMajorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionMajorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// psMSMQVersionMajor {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.WriteData(o.VersionMajor); err != nil {
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

func (o *xxx_GetVersionMajorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// psMSMQVersionMajor {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.ReadData(&o.VersionMajor); err != nil {
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

// GetVersionMajorRequest structure represents the MSMQVersionMajor operation request
type GetVersionMajorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetVersionMajorRequest) xxx_ToOp(ctx context.Context, op *xxx_GetVersionMajorOperation) *xxx_GetVersionMajorOperation {
	if op == nil {
		op = &xxx_GetVersionMajorOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetVersionMajorRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVersionMajorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetVersionMajorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetVersionMajorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionMajorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVersionMajorResponse structure represents the MSMQVersionMajor operation response
type GetVersionMajorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// psMSMQVersionMajor: A pointer to a short that when successfully completed, contains
	// the major version number of the server.
	VersionMajor int16 `idl:"name:psMSMQVersionMajor" json:"version_major"`
	// Return: The MSMQVersionMajor return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetVersionMajorResponse) xxx_ToOp(ctx context.Context, op *xxx_GetVersionMajorOperation) *xxx_GetVersionMajorOperation {
	if op == nil {
		op = &xxx_GetVersionMajorOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.VersionMajor = o.VersionMajor
	op.Return = o.Return
	return op
}

func (o *GetVersionMajorResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVersionMajorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VersionMajor = op.VersionMajor
	o.Return = op.Return
}
func (o *GetVersionMajorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetVersionMajorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionMajorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVersionMinorOperation structure represents the MSMQVersionMinor operation
type xxx_GetVersionMinorOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	VersionMinor int16          `idl:"name:psMSMQVersionMinor" json:"version_minor"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVersionMinorOperation) OpNum() int { return 11 }

func (o *xxx_GetVersionMinorOperation) OpName() string {
	return "/IMSMQApplication2/v0/MSMQVersionMinor"
}

func (o *xxx_GetVersionMinorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionMinorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetVersionMinorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetVersionMinorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionMinorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// psMSMQVersionMinor {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.WriteData(o.VersionMinor); err != nil {
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

func (o *xxx_GetVersionMinorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// psMSMQVersionMinor {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.ReadData(&o.VersionMinor); err != nil {
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

// GetVersionMinorRequest structure represents the MSMQVersionMinor operation request
type GetVersionMinorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetVersionMinorRequest) xxx_ToOp(ctx context.Context, op *xxx_GetVersionMinorOperation) *xxx_GetVersionMinorOperation {
	if op == nil {
		op = &xxx_GetVersionMinorOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetVersionMinorRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVersionMinorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetVersionMinorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetVersionMinorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionMinorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVersionMinorResponse structure represents the MSMQVersionMinor operation response
type GetVersionMinorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// psMSMQVersionMinor: A pointer to a short that when successfully completed, contains
	// the minor version number of the server.
	VersionMinor int16 `idl:"name:psMSMQVersionMinor" json:"version_minor"`
	// Return: The MSMQVersionMinor return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetVersionMinorResponse) xxx_ToOp(ctx context.Context, op *xxx_GetVersionMinorOperation) *xxx_GetVersionMinorOperation {
	if op == nil {
		op = &xxx_GetVersionMinorOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.VersionMinor = o.VersionMinor
	op.Return = o.Return
	return op
}

func (o *GetVersionMinorResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVersionMinorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VersionMinor = op.VersionMinor
	o.Return = op.Return
}
func (o *GetVersionMinorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetVersionMinorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionMinorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVersionBuildOperation structure represents the MSMQVersionBuild operation
type xxx_GetVersionBuildOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	VersionBuild int16          `idl:"name:psMSMQVersionBuild" json:"version_build"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVersionBuildOperation) OpNum() int { return 12 }

func (o *xxx_GetVersionBuildOperation) OpName() string {
	return "/IMSMQApplication2/v0/MSMQVersionBuild"
}

func (o *xxx_GetVersionBuildOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionBuildOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetVersionBuildOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetVersionBuildOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionBuildOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// psMSMQVersionBuild {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.WriteData(o.VersionBuild); err != nil {
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

func (o *xxx_GetVersionBuildOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// psMSMQVersionBuild {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.ReadData(&o.VersionBuild); err != nil {
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

// GetVersionBuildRequest structure represents the MSMQVersionBuild operation request
type GetVersionBuildRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetVersionBuildRequest) xxx_ToOp(ctx context.Context, op *xxx_GetVersionBuildOperation) *xxx_GetVersionBuildOperation {
	if op == nil {
		op = &xxx_GetVersionBuildOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetVersionBuildRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVersionBuildOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetVersionBuildRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetVersionBuildRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionBuildOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVersionBuildResponse structure represents the MSMQVersionBuild operation response
type GetVersionBuildResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// psMSMQVersionBuild: A pointer to a short that when successfully completed, contains
	// the build version number of the server.
	VersionBuild int16 `idl:"name:psMSMQVersionBuild" json:"version_build"`
	// Return: The MSMQVersionBuild return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetVersionBuildResponse) xxx_ToOp(ctx context.Context, op *xxx_GetVersionBuildOperation) *xxx_GetVersionBuildOperation {
	if op == nil {
		op = &xxx_GetVersionBuildOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.VersionBuild = o.VersionBuild
	op.Return = o.Return
	return op
}

func (o *GetVersionBuildResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVersionBuildOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VersionBuild = op.VersionBuild
	o.Return = op.Return
}
func (o *GetVersionBuildResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetVersionBuildResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionBuildOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsDSEnabledOperation structure represents the IsDsEnabled operation
type xxx_GetIsDSEnabledOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsDSEnabled int16          `idl:"name:pfIsDsEnabled" json:"is_ds_enabled"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsDSEnabledOperation) OpNum() int { return 13 }

func (o *xxx_GetIsDSEnabledOperation) OpName() string { return "/IMSMQApplication2/v0/IsDsEnabled" }

func (o *xxx_GetIsDSEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsDSEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsDSEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsDSEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsDSEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfIsDsEnabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.IsDSEnabled); err != nil {
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

func (o *xxx_GetIsDSEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfIsDsEnabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.IsDSEnabled); err != nil {
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

// GetIsDSEnabledRequest structure represents the IsDsEnabled operation request
type GetIsDSEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsDSEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsDSEnabledOperation) *xxx_GetIsDSEnabledOperation {
	if op == nil {
		op = &xxx_GetIsDSEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsDSEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsDSEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsDSEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsDSEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsDSEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsDSEnabledResponse structure represents the IsDsEnabled operation response
type GetIsDSEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pfIsDsEnabled: A pointer to a VARIANT_BOOL that when successfully completed, contains
	// one of the following values.
	//
	//	+----------------------+----------------------------------------------------------------+
	//	|                      |                                                                |
	//	|        VALUE         |                            MEANING                             |
	//	|                      |                                                                |
	//	+----------------------+----------------------------------------------------------------+
	//	+----------------------+----------------------------------------------------------------+
	//	| VARIANT_TRUE 0xffff  | The local QueueManager is configured to use the directory.     |
	//	+----------------------+----------------------------------------------------------------+
	//	| VARIANT_FALSE 0x0000 | The local QueueManager is not configured to use the directory. |
	//	+----------------------+----------------------------------------------------------------+
	IsDSEnabled int16 `idl:"name:pfIsDsEnabled" json:"is_ds_enabled"`
	// Return: The IsDsEnabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsDSEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsDSEnabledOperation) *xxx_GetIsDSEnabledOperation {
	if op == nil {
		op = &xxx_GetIsDSEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsDSEnabled = o.IsDSEnabled
	op.Return = o.Return
	return op
}

func (o *GetIsDSEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsDSEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsDSEnabled = op.IsDSEnabled
	o.Return = op.Return
}
func (o *GetIsDSEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsDSEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsDSEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertiesOperation structure represents the Properties operation
type xxx_GetPropertiesOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Properties *oaut.Dispatch `idl:"name:ppcolProperties" json:"properties"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 14 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/IMSMQApplication2/v0/Properties" }

func (o *xxx_GetPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppcolProperties {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		if o.Properties != nil {
			_ptr_ppcolProperties := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Properties != nil {
					if err := o.Properties.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Properties, _ptr_ppcolProperties); err != nil {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppcolProperties {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_ppcolProperties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Properties == nil {
				o.Properties = &oaut.Dispatch{}
			}
			if err := o.Properties.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppcolProperties := func(ptr interface{}) { o.Properties = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.Properties, _s_ppcolProperties, _ptr_ppcolProperties); err != nil {
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

// GetPropertiesRequest structure represents the Properties operation request
type GetPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesResponse structure represents the Properties operation response
type GetPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Properties *oaut.Dispatch `idl:"name:ppcolProperties" json:"properties"`
	// Return: The Properties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Properties = o.Properties
	op.Return = o.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Properties = op.Properties
	o.Return = op.Return
}
func (o *GetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
