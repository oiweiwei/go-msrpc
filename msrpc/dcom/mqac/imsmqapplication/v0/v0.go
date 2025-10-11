package imsmqapplication

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQApplication interface identifier d7d6e085-dccd-11d0-aa4b-0060970debae
	ApplicationIID = &dcom.IID{Data1: 0xd7d6e085, Data2: 0xdccd, Data3: 0x11d0, Data4: []byte{0xaa, 0x4b, 0x00, 0x60, 0x97, 0x0d, 0xeb, 0xae}}
	// Syntax UUID
	ApplicationSyntaxUUID = &uuid.UUID{TimeLow: 0xd7d6e085, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	// Syntax ID
	ApplicationSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ApplicationSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQApplication interface.
type ApplicationClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The MachineIdOfMachineName method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return a string that contains the QueueManager.Identifier
	// for the computer name that was passed as the input parameter.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<9>
	MachineIDOfMachineName(context.Context, *MachineIDOfMachineNameRequest, ...dcerpc.CallOption) (*MachineIDOfMachineNameResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ApplicationClient
}

type xxx_DefaultApplicationClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultApplicationClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultApplicationClient) MachineIDOfMachineName(ctx context.Context, in *MachineIDOfMachineNameRequest, opts ...dcerpc.CallOption) (*MachineIDOfMachineNameResponse, error) {
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
	out := &MachineIDOfMachineNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultApplicationClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultApplicationClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultApplicationClient) IPID(ctx context.Context, ipid *dcom.IPID) ApplicationClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultApplicationClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewApplicationClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ApplicationClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ApplicationSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultApplicationClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_MachineIDOfMachineNameOperation structure represents the MachineIdOfMachineName operation
type xxx_MachineIDOfMachineNameOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MachineName *oaut.String   `idl:"name:MachineName" json:"machine_name"`
	GUID        *oaut.String   `idl:"name:pbstrGuid" json:"guid"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_MachineIDOfMachineNameOperation) OpNum() int { return 7 }

func (o *xxx_MachineIDOfMachineNameOperation) OpName() string {
	return "/IMSMQApplication/v0/MachineIdOfMachineName"
}

func (o *xxx_MachineIDOfMachineNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MachineIDOfMachineNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// MachineName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineName != nil {
			_ptr_MachineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.MachineName, _ptr_MachineName); err != nil {
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

func (o *xxx_MachineIDOfMachineNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// MachineName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_MachineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineName == nil {
				o.MachineName = &oaut.String{}
			}
			if err := o.MachineName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_MachineName := func(ptr interface{}) { o.MachineName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineName, _s_MachineName, _ptr_MachineName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MachineIDOfMachineNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MachineIDOfMachineNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrGuid {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.GUID != nil {
			_ptr_pbstrGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.GUID, _ptr_pbstrGuid); err != nil {
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

func (o *xxx_MachineIDOfMachineNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrGuid {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.GUID == nil {
				o.GUID = &oaut.String{}
			}
			if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrGuid := func(ptr interface{}) { o.GUID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.GUID, _s_pbstrGuid, _ptr_pbstrGuid); err != nil {
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

// MachineIDOfMachineNameRequest structure represents the MachineIdOfMachineName operation request
type MachineIDOfMachineNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// MachineName: A BSTR that specifies the NETBIOS or DNS computer name for which a GUID
	// is to be retrieved.
	MachineName *oaut.String `idl:"name:MachineName" json:"machine_name"`
}

func (o *MachineIDOfMachineNameRequest) xxx_ToOp(ctx context.Context, op *xxx_MachineIDOfMachineNameOperation) *xxx_MachineIDOfMachineNameOperation {
	if op == nil {
		op = &xxx_MachineIDOfMachineNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MachineName = o.MachineName
	return op
}

func (o *MachineIDOfMachineNameRequest) xxx_FromOp(ctx context.Context, op *xxx_MachineIDOfMachineNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MachineName = op.MachineName
}
func (o *MachineIDOfMachineNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MachineIDOfMachineNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MachineIDOfMachineNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MachineIDOfMachineNameResponse structure represents the MachineIdOfMachineName operation response
type MachineIDOfMachineNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrGuid: A pointer to a BSTR that upon successful completion, contains the string
	// representation of the specified machine GUID. The server MUST NOT include the surrounding
	// curly braces ({ }) with the returned GUID. The string MUST use the following format,
	// which is specified in ABNF.
	GUID *oaut.String `idl:"name:pbstrGuid" json:"guid"`
	// Return: The MachineIdOfMachineName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MachineIDOfMachineNameResponse) xxx_ToOp(ctx context.Context, op *xxx_MachineIDOfMachineNameOperation) *xxx_MachineIDOfMachineNameOperation {
	if op == nil {
		op = &xxx_MachineIDOfMachineNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.GUID = o.GUID
	op.Return = o.Return
	return op
}

func (o *MachineIDOfMachineNameResponse) xxx_FromOp(ctx context.Context, op *xxx_MachineIDOfMachineNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.GUID = op.GUID
	o.Return = op.Return
}
func (o *MachineIDOfMachineNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MachineIDOfMachineNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MachineIDOfMachineNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
