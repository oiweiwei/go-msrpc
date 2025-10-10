package imsmqapplication2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
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
	ImsmqApplication2IID = &dcom.IID{Data1: 0x12a30900, Data2: 0x7300, Data3: 0x11d2, Data4: []byte{0xb0, 0xe6, 0x00, 0xe0, 0x2c, 0x07, 0x4f, 0x6b}}
	// Syntax UUID
	ImsmqApplication2SyntaxUUID = &uuid.UUID{TimeLow: 0x12a30900, TimeMid: 0x7300, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0xe6, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	// Syntax ID
	ImsmqApplication2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ImsmqApplication2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQApplication2 interface.
type ImsmqApplication2Client interface {

	// IMSMQApplication retrieval method.
	ImsmqApplication() imsmqapplication.ImsmqApplicationClient

	// RegisterCertificate operation.
	RegisterCertificate(context.Context, *RegisterCertificateRequest, ...dcerpc.CallOption) (*RegisterCertificateResponse, error)

	// MachineNameOfMachineId operation.
	MachineNameOfMachineID(context.Context, *MachineNameOfMachineIDRequest, ...dcerpc.CallOption) (*MachineNameOfMachineIDResponse, error)

	// MSMQVersionMajor operation.
	GetMsmqVersionMajor(context.Context, *GetMsmqVersionMajorRequest, ...dcerpc.CallOption) (*GetMsmqVersionMajorResponse, error)

	// MSMQVersionMinor operation.
	GetMsmqVersionMinor(context.Context, *GetMsmqVersionMinorRequest, ...dcerpc.CallOption) (*GetMsmqVersionMinorResponse, error)

	// MSMQVersionBuild operation.
	GetMsmqVersionBuild(context.Context, *GetMsmqVersionBuildRequest, ...dcerpc.CallOption) (*GetMsmqVersionBuildResponse, error)

	// IsDsEnabled operation.
	GetIsDSEnabled(context.Context, *GetIsDSEnabledRequest, ...dcerpc.CallOption) (*GetIsDSEnabledResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ImsmqApplication2Client
}

type xxx_DefaultImsmqApplication2Client struct {
	imsmqapplication.ImsmqApplicationClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultImsmqApplication2Client) ImsmqApplication() imsmqapplication.ImsmqApplicationClient {
	return o.ImsmqApplicationClient
}

func (o *xxx_DefaultImsmqApplication2Client) RegisterCertificate(ctx context.Context, in *RegisterCertificateRequest, opts ...dcerpc.CallOption) (*RegisterCertificateResponse, error) {
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
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication2Client) MachineNameOfMachineID(ctx context.Context, in *MachineNameOfMachineIDRequest, opts ...dcerpc.CallOption) (*MachineNameOfMachineIDResponse, error) {
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
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication2Client) GetMsmqVersionMajor(ctx context.Context, in *GetMsmqVersionMajorRequest, opts ...dcerpc.CallOption) (*GetMsmqVersionMajorResponse, error) {
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
	out := &GetMsmqVersionMajorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication2Client) GetMsmqVersionMinor(ctx context.Context, in *GetMsmqVersionMinorRequest, opts ...dcerpc.CallOption) (*GetMsmqVersionMinorResponse, error) {
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
	out := &GetMsmqVersionMinorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication2Client) GetMsmqVersionBuild(ctx context.Context, in *GetMsmqVersionBuildRequest, opts ...dcerpc.CallOption) (*GetMsmqVersionBuildResponse, error) {
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
	out := &GetMsmqVersionBuildResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication2Client) GetIsDSEnabled(ctx context.Context, in *GetIsDSEnabledRequest, opts ...dcerpc.CallOption) (*GetIsDSEnabledResponse, error) {
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
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication2Client) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultImsmqApplication2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultImsmqApplication2Client) IPID(ctx context.Context, ipid *dcom.IPID) ImsmqApplication2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultImsmqApplication2Client{
		ImsmqApplicationClient: o.ImsmqApplicationClient.IPID(ctx, ipid),
		cc:                     o.cc,
		ipid:                   ipid,
	}
}

func NewImsmqApplication2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ImsmqApplication2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ImsmqApplication2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := imsmqapplication.NewImsmqApplicationClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultImsmqApplication2Client{
		ImsmqApplicationClient: base,
		cc:                     cc,
		ipid:                   ipid,
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
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	Flags               *oaut.Variant  `idl:"name:Flags" json:"flags"`
	ExternalCertificate *oaut.Variant  `idl:"name:ExternalCertificate" json:"external_certificate"`
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
	GUID *oaut.String   `idl:"name:bstrGuid" json:"guid"`
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
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MachineName *oaut.String   `idl:"name:pbstrMachineName" json:"machine_name"`
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

// xxx_GetMsmqVersionMajorOperation structure represents the MSMQVersionMajor operation
type xxx_GetMsmqVersionMajorOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	PsMsmqVersionMajor int16          `idl:"name:psMSMQVersionMajor" json:"ps_msmq_version_major"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMsmqVersionMajorOperation) OpNum() int { return 10 }

func (o *xxx_GetMsmqVersionMajorOperation) OpName() string {
	return "/IMSMQApplication2/v0/MSMQVersionMajor"
}

func (o *xxx_GetMsmqVersionMajorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMsmqVersionMajorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMsmqVersionMajorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMsmqVersionMajorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMsmqVersionMajorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
		if err := w.WriteData(o.PsMsmqVersionMajor); err != nil {
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

func (o *xxx_GetMsmqVersionMajorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
		if err := w.ReadData(&o.PsMsmqVersionMajor); err != nil {
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

// GetMsmqVersionMajorRequest structure represents the MSMQVersionMajor operation request
type GetMsmqVersionMajorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMsmqVersionMajorRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMsmqVersionMajorOperation) *xxx_GetMsmqVersionMajorOperation {
	if op == nil {
		op = &xxx_GetMsmqVersionMajorOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMsmqVersionMajorRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMsmqVersionMajorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMsmqVersionMajorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMsmqVersionMajorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMsmqVersionMajorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMsmqVersionMajorResponse structure represents the MSMQVersionMajor operation response
type GetMsmqVersionMajorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	PsMsmqVersionMajor int16          `idl:"name:psMSMQVersionMajor" json:"ps_msmq_version_major"`
	// Return: The MSMQVersionMajor return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMsmqVersionMajorResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMsmqVersionMajorOperation) *xxx_GetMsmqVersionMajorOperation {
	if op == nil {
		op = &xxx_GetMsmqVersionMajorOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PsMsmqVersionMajor = o.PsMsmqVersionMajor
	op.Return = o.Return
	return op
}

func (o *GetMsmqVersionMajorResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMsmqVersionMajorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PsMsmqVersionMajor = op.PsMsmqVersionMajor
	o.Return = op.Return
}
func (o *GetMsmqVersionMajorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMsmqVersionMajorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMsmqVersionMajorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMsmqVersionMinorOperation structure represents the MSMQVersionMinor operation
type xxx_GetMsmqVersionMinorOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	PsMsmqVersionMinor int16          `idl:"name:psMSMQVersionMinor" json:"ps_msmq_version_minor"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMsmqVersionMinorOperation) OpNum() int { return 11 }

func (o *xxx_GetMsmqVersionMinorOperation) OpName() string {
	return "/IMSMQApplication2/v0/MSMQVersionMinor"
}

func (o *xxx_GetMsmqVersionMinorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMsmqVersionMinorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMsmqVersionMinorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMsmqVersionMinorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMsmqVersionMinorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
		if err := w.WriteData(o.PsMsmqVersionMinor); err != nil {
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

func (o *xxx_GetMsmqVersionMinorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
		if err := w.ReadData(&o.PsMsmqVersionMinor); err != nil {
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

// GetMsmqVersionMinorRequest structure represents the MSMQVersionMinor operation request
type GetMsmqVersionMinorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMsmqVersionMinorRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMsmqVersionMinorOperation) *xxx_GetMsmqVersionMinorOperation {
	if op == nil {
		op = &xxx_GetMsmqVersionMinorOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMsmqVersionMinorRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMsmqVersionMinorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMsmqVersionMinorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMsmqVersionMinorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMsmqVersionMinorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMsmqVersionMinorResponse structure represents the MSMQVersionMinor operation response
type GetMsmqVersionMinorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	PsMsmqVersionMinor int16          `idl:"name:psMSMQVersionMinor" json:"ps_msmq_version_minor"`
	// Return: The MSMQVersionMinor return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMsmqVersionMinorResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMsmqVersionMinorOperation) *xxx_GetMsmqVersionMinorOperation {
	if op == nil {
		op = &xxx_GetMsmqVersionMinorOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PsMsmqVersionMinor = o.PsMsmqVersionMinor
	op.Return = o.Return
	return op
}

func (o *GetMsmqVersionMinorResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMsmqVersionMinorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PsMsmqVersionMinor = op.PsMsmqVersionMinor
	o.Return = op.Return
}
func (o *GetMsmqVersionMinorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMsmqVersionMinorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMsmqVersionMinorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMsmqVersionBuildOperation structure represents the MSMQVersionBuild operation
type xxx_GetMsmqVersionBuildOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	PsMsmqVersionBuild int16          `idl:"name:psMSMQVersionBuild" json:"ps_msmq_version_build"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMsmqVersionBuildOperation) OpNum() int { return 12 }

func (o *xxx_GetMsmqVersionBuildOperation) OpName() string {
	return "/IMSMQApplication2/v0/MSMQVersionBuild"
}

func (o *xxx_GetMsmqVersionBuildOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMsmqVersionBuildOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMsmqVersionBuildOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMsmqVersionBuildOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMsmqVersionBuildOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
		if err := w.WriteData(o.PsMsmqVersionBuild); err != nil {
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

func (o *xxx_GetMsmqVersionBuildOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
		if err := w.ReadData(&o.PsMsmqVersionBuild); err != nil {
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

// GetMsmqVersionBuildRequest structure represents the MSMQVersionBuild operation request
type GetMsmqVersionBuildRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMsmqVersionBuildRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMsmqVersionBuildOperation) *xxx_GetMsmqVersionBuildOperation {
	if op == nil {
		op = &xxx_GetMsmqVersionBuildOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMsmqVersionBuildRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMsmqVersionBuildOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMsmqVersionBuildRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMsmqVersionBuildRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMsmqVersionBuildOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMsmqVersionBuildResponse structure represents the MSMQVersionBuild operation response
type GetMsmqVersionBuildResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	PsMsmqVersionBuild int16          `idl:"name:psMSMQVersionBuild" json:"ps_msmq_version_build"`
	// Return: The MSMQVersionBuild return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMsmqVersionBuildResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMsmqVersionBuildOperation) *xxx_GetMsmqVersionBuildOperation {
	if op == nil {
		op = &xxx_GetMsmqVersionBuildOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PsMsmqVersionBuild = o.PsMsmqVersionBuild
	op.Return = o.Return
	return op
}

func (o *GetMsmqVersionBuildResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMsmqVersionBuildOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PsMsmqVersionBuild = op.PsMsmqVersionBuild
	o.Return = op.Return
}
func (o *GetMsmqVersionBuildResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMsmqVersionBuildResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMsmqVersionBuildOperation{}
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
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsDSEnabled int16          `idl:"name:pfIsDsEnabled" json:"is_ds_enabled"`
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
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	PpcolProperties *oaut.Dispatch `idl:"name:ppcolProperties" json:"ppcol_properties"`
	Return          int32          `idl:"name:Return" json:"return"`
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
		if o.PpcolProperties != nil {
			_ptr_ppcolProperties := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PpcolProperties != nil {
					if err := o.PpcolProperties.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PpcolProperties, _ptr_ppcolProperties); err != nil {
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
			if o.PpcolProperties == nil {
				o.PpcolProperties = &oaut.Dispatch{}
			}
			if err := o.PpcolProperties.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppcolProperties := func(ptr interface{}) { o.PpcolProperties = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.PpcolProperties, _s_ppcolProperties, _ptr_ppcolProperties); err != nil {
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
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	PpcolProperties *oaut.Dispatch `idl:"name:ppcolProperties" json:"ppcol_properties"`
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
	op.PpcolProperties = o.PpcolProperties
	op.Return = o.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PpcolProperties = op.PpcolProperties
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
