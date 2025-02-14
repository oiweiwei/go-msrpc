package ivdsserviceloader

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
	vds "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds"
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
	_ = vds.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsServiceLoader interface identifier e0393303-90d4-4a97-ab71-e9b671ee2729
	ServiceLoaderIID = &dcom.IID{Data1: 0xe0393303, Data2: 0x90d4, Data3: 0x4a97, Data4: []byte{0xab, 0x71, 0xe9, 0xb6, 0x71, 0xee, 0x27, 0x29}}
	// Syntax UUID
	ServiceLoaderSyntaxUUID = &uuid.UUID{TimeLow: 0xe0393303, TimeMid: 0x90d4, TimeHiAndVersion: 0x4a97, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x71, Node: [6]uint8{0xe9, 0xb6, 0x71, 0xee, 0x27, 0x29}}
	// Syntax ID
	ServiceLoaderSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ServiceLoaderSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsServiceLoader interface.
type ServiceLoaderClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The LoadService method is used by client applications to load the VDS service on
	// a remote machine.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	LoadService(context.Context, *LoadServiceRequest, ...dcerpc.CallOption) (*LoadServiceResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ServiceLoaderClient
}

type xxx_DefaultServiceLoaderClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultServiceLoaderClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultServiceLoaderClient) LoadService(ctx context.Context, in *LoadServiceRequest, opts ...dcerpc.CallOption) (*LoadServiceResponse, error) {
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
	out := &LoadServiceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceLoaderClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultServiceLoaderClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultServiceLoaderClient) IPID(ctx context.Context, ipid *dcom.IPID) ServiceLoaderClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultServiceLoaderClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewServiceLoaderClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ServiceLoaderClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ServiceLoaderSyntaxV0_0))...)
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
	return &xxx_DefaultServiceLoaderClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_LoadServiceOperation structure represents the LoadService operation
type xxx_LoadServiceOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MachineName string         `idl:"name:pwszMachineName;string;pointer:unique" json:"machine_name"`
	Service     *vds.Service   `idl:"name:ppService" json:"service"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_LoadServiceOperation) OpNum() int { return 3 }

func (o *xxx_LoadServiceOperation) OpName() string { return "/IVdsServiceLoader/v0/LoadService" }

func (o *xxx_LoadServiceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LoadServiceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszMachineName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.MachineName != "" {
			_ptr_pwszMachineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.MachineName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineName, _ptr_pwszMachineName); err != nil {
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

func (o *xxx_LoadServiceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszMachineName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszMachineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.MachineName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszMachineName := func(ptr interface{}) { o.MachineName = *ptr.(*string) }
		if err := w.ReadPointer(&o.MachineName, _s_pwszMachineName, _ptr_pwszMachineName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LoadServiceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LoadServiceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppService {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsService}(interface))
	{
		if o.Service != nil {
			_ptr_ppService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Service != nil {
					if err := o.Service.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Service{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Service, _ptr_ppService); err != nil {
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

func (o *xxx_LoadServiceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppService {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsService}(interface))
	{
		_ptr_ppService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Service == nil {
				o.Service = &vds.Service{}
			}
			if err := o.Service.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppService := func(ptr interface{}) { o.Service = *ptr.(**vds.Service) }
		if err := w.ReadPointer(&o.Service, _s_ppService, _ptr_ppService); err != nil {
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

// LoadServiceRequest structure represents the LoadService operation request
type LoadServiceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszMachineName: A pointer to a string that contains the name of the machine on which
	// the VDS service is loaded.
	MachineName string `idl:"name:pwszMachineName;string;pointer:unique" json:"machine_name"`
}

func (o *LoadServiceRequest) xxx_ToOp(ctx context.Context, op *xxx_LoadServiceOperation) *xxx_LoadServiceOperation {
	if op == nil {
		op = &xxx_LoadServiceOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.MachineName = op.MachineName
	return op
}

func (o *LoadServiceRequest) xxx_FromOp(ctx context.Context, op *xxx_LoadServiceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MachineName = op.MachineName
}
func (o *LoadServiceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LoadServiceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LoadServiceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LoadServiceResponse structure represents the LoadService operation response
type LoadServiceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppService: A pointer to the IVdsService interface that, if successfully completed,
	// returns the IVdsService interface to the VDS service that runs on the machine represented
	// by pwszMachineName.
	Service *vds.Service `idl:"name:ppService" json:"service"`
	// Return: The LoadService return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LoadServiceResponse) xxx_ToOp(ctx context.Context, op *xxx_LoadServiceOperation) *xxx_LoadServiceOperation {
	if op == nil {
		op = &xxx_LoadServiceOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Service = op.Service
	o.Return = op.Return
	return op
}

func (o *LoadServiceResponse) xxx_FromOp(ctx context.Context, op *xxx_LoadServiceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Service = op.Service
	o.Return = op.Return
}
func (o *LoadServiceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LoadServiceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LoadServiceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
