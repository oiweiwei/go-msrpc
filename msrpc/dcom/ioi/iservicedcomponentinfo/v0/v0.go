package iservicedcomponentinfo

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
	_ = iunknown.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/ioi"
)

var (
	// IServicedComponentInfo interface identifier 8165b19e-8d3a-4d0b-80c8-97de310db583
	ServicedComponentInfoIID = &dcom.IID{Data1: 0x8165b19e, Data2: 0x8d3a, Data3: 0x4d0b, Data4: []byte{0x80, 0xc8, 0x97, 0xde, 0x31, 0x0d, 0xb5, 0x83}}
	// Syntax UUID
	ServicedComponentInfoSyntaxUUID = &uuid.UUID{TimeLow: 0x8165b19e, TimeMid: 0x8d3a, TimeHiAndVersion: 0x4d0b, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0xc8, Node: [6]uint8{0x97, 0xde, 0x31, 0xd, 0xb5, 0x83}}
	// Syntax ID
	ServicedComponentInfoSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ServicedComponentInfoSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IServicedComponentInfo interface.
type ServicedComponentInfoClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetComponentInfo method is used to determine the environment of the server object.
	//
	// Return Values: An HRESULT that specifies success or failure. All success HRESULT
	// values MUST be treated as success and all failure HRESULT values MUST be treated
	// as failure.
	GetComponentInfo(context.Context, *GetComponentInfoRequest, ...dcerpc.CallOption) (*GetComponentInfoResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ServicedComponentInfoClient
}

type xxx_DefaultServicedComponentInfoClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultServicedComponentInfoClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultServicedComponentInfoClient) GetComponentInfo(ctx context.Context, in *GetComponentInfoRequest, opts ...dcerpc.CallOption) (*GetComponentInfoResponse, error) {
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
	out := &GetComponentInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicedComponentInfoClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultServicedComponentInfoClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultServicedComponentInfoClient) IPID(ctx context.Context, ipid *dcom.IPID) ServicedComponentInfoClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultServicedComponentInfoClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewServicedComponentInfoClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ServicedComponentInfoClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ServicedComponentInfoSyntaxV0_0))...)
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
	return &xxx_DefaultServicedComponentInfoClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetComponentInfoOperation structure represents the GetComponentInfo operation
type xxx_GetComponentInfoOperation struct {
	This      *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	InfoMask  int32           `idl:"name:infoMask" json:"info_mask"`
	InfoArray *oaut.SafeArray `idl:"name:infoArray" json:"info_array"`
	Return    int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetComponentInfoOperation) OpNum() int { return 3 }

func (o *xxx_GetComponentInfoOperation) OpName() string {
	return "/IServicedComponentInfo/v0/GetComponentInfo"
}

func (o *xxx_GetComponentInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// infoMask {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.InfoMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// infoMask {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.InfoMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// infoMask {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.InfoMask); err != nil {
			return err
		}
	}
	// infoArray {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.InfoArray != nil {
			_ptr_infoArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InfoArray != nil {
					if err := o.InfoArray.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InfoArray, _ptr_infoArray); err != nil {
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

func (o *xxx_GetComponentInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// infoMask {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.InfoMask); err != nil {
			return err
		}
	}
	// infoArray {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_infoArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InfoArray == nil {
				o.InfoArray = &oaut.SafeArray{}
			}
			if err := o.InfoArray.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_infoArray := func(ptr interface{}) { o.InfoArray = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.InfoArray, _s_infoArray, _ptr_infoArray); err != nil {
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

// GetComponentInfoRequest structure represents the GetComponentInfo operation request
type GetComponentInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// infoMask: A bitwise OR of zero of more of the following values:
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The serviced component's process identifier (PID).                               |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | The serviced component's application domain identifier (ID).                     |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000004 | The serviced component's remote URI [RFC3986], which represents the server       |
	//	|            | object identity.                                                                 |
	//	+------------+----------------------------------------------------------------------------------+
	InfoMask int32 `idl:"name:infoMask" json:"info_mask"`
}

func (o *GetComponentInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetComponentInfoOperation) *xxx_GetComponentInfoOperation {
	if op == nil {
		op = &xxx_GetComponentInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.InfoMask = o.InfoMask
	return op
}

func (o *GetComponentInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetComponentInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.InfoMask = op.InfoMask
}
func (o *GetComponentInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetComponentInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetComponentInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetComponentInfoResponse structure represents the GetComponentInfo operation response
type GetComponentInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// infoMask: A bitwise OR of zero of more of the following values:
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The serviced component's process identifier (PID).                               |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | The serviced component's application domain identifier (ID).                     |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000004 | The serviced component's remote URI [RFC3986], which represents the server       |
	//	|            | object identity.                                                                 |
	//	+------------+----------------------------------------------------------------------------------+
	InfoMask int32 `idl:"name:infoMask" json:"info_mask"`
	// infoArray: An array that contains a set of values returned by the server corresponding
	// to the bits set in infoMask.
	InfoArray *oaut.SafeArray `idl:"name:infoArray" json:"info_array"`
	// Return: The GetComponentInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetComponentInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetComponentInfoOperation) *xxx_GetComponentInfoOperation {
	if op == nil {
		op = &xxx_GetComponentInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.InfoMask = o.InfoMask
	op.InfoArray = o.InfoArray
	op.Return = o.Return
	return op
}

func (o *GetComponentInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetComponentInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.InfoMask = op.InfoMask
	o.InfoArray = op.InfoArray
	o.Return = op.Return
}
func (o *GetComponentInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetComponentInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetComponentInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
