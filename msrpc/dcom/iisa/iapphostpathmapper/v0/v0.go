package iapphostpathmapper

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
	GoPackage = "dcom/iisa"
)

var (
	// IAppHostPathMapper interface identifier e7927575-5cc3-403b-822e-328a6b904bee
	AppHostPathMapperIID = &dcom.IID{Data1: 0xe7927575, Data2: 0x5cc3, Data3: 0x403b, Data4: []byte{0x82, 0x2e, 0x32, 0x8a, 0x6b, 0x90, 0x4b, 0xee}}
	// Syntax UUID
	AppHostPathMapperSyntaxUUID = &uuid.UUID{TimeLow: 0xe7927575, TimeMid: 0x5cc3, TimeHiAndVersion: 0x403b, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x2e, Node: [6]uint8{0x32, 0x8a, 0x6b, 0x90, 0x4b, 0xee}}
	// Syntax ID
	AppHostPathMapperSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostPathMapperSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostPathMapper interface.
type AppHostPathMapperClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// MapPath operation.
	MapPath(context.Context, *MapPathRequest, ...dcerpc.CallOption) (*MapPathResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostPathMapperClient
}

type xxx_DefaultAppHostPathMapperClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostPathMapperClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAppHostPathMapperClient) MapPath(ctx context.Context, in *MapPathRequest, opts ...dcerpc.CallOption) (*MapPathResponse, error) {
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
	out := &MapPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostPathMapperClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostPathMapperClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAppHostPathMapperClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostPathMapperClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostPathMapperClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAppHostPathMapperClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostPathMapperClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostPathMapperSyntaxV0_0))...)
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
	return &xxx_DefaultAppHostPathMapperClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_MapPathOperation structure represents the MapPath operation
type xxx_MapPathOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConfigPath         *oaut.String   `idl:"name:bstrConfigPath" json:"config_path"`
	MappedPhysicalPath *oaut.String   `idl:"name:bstrMappedPhysicalPath" json:"mapped_physical_path"`
	NewPhysicalPath    *oaut.String   `idl:"name:pbstrNewPhysicalPath" json:"new_physical_path"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_MapPathOperation) OpNum() int { return 3 }

func (o *xxx_MapPathOperation) OpName() string { return "/IAppHostPathMapper/v0/MapPath" }

func (o *xxx_MapPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MapPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrConfigPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConfigPath != nil {
			_ptr_bstrConfigPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigPath != nil {
					if err := o.ConfigPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigPath, _ptr_bstrConfigPath); err != nil {
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
	// bstrMappedPhysicalPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MappedPhysicalPath != nil {
			_ptr_bstrMappedPhysicalPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MappedPhysicalPath != nil {
					if err := o.MappedPhysicalPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MappedPhysicalPath, _ptr_bstrMappedPhysicalPath); err != nil {
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

func (o *xxx_MapPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrConfigPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConfigPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigPath == nil {
				o.ConfigPath = &oaut.String{}
			}
			if err := o.ConfigPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConfigPath := func(ptr interface{}) { o.ConfigPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConfigPath, _s_bstrConfigPath, _ptr_bstrConfigPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMappedPhysicalPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMappedPhysicalPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MappedPhysicalPath == nil {
				o.MappedPhysicalPath = &oaut.String{}
			}
			if err := o.MappedPhysicalPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMappedPhysicalPath := func(ptr interface{}) { o.MappedPhysicalPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MappedPhysicalPath, _s_bstrMappedPhysicalPath, _ptr_bstrMappedPhysicalPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MapPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MapPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrNewPhysicalPath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NewPhysicalPath != nil {
			_ptr_pbstrNewPhysicalPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewPhysicalPath != nil {
					if err := o.NewPhysicalPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewPhysicalPath, _ptr_pbstrNewPhysicalPath); err != nil {
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

func (o *xxx_MapPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrNewPhysicalPath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrNewPhysicalPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewPhysicalPath == nil {
				o.NewPhysicalPath = &oaut.String{}
			}
			if err := o.NewPhysicalPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrNewPhysicalPath := func(ptr interface{}) { o.NewPhysicalPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NewPhysicalPath, _s_pbstrNewPhysicalPath, _ptr_pbstrNewPhysicalPath); err != nil {
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

// MapPathRequest structure represents the MapPath operation request
type MapPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	ConfigPath         *oaut.String   `idl:"name:bstrConfigPath" json:"config_path"`
	MappedPhysicalPath *oaut.String   `idl:"name:bstrMappedPhysicalPath" json:"mapped_physical_path"`
}

func (o *MapPathRequest) xxx_ToOp(ctx context.Context, op *xxx_MapPathOperation) *xxx_MapPathOperation {
	if op == nil {
		op = &xxx_MapPathOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.ConfigPath = op.ConfigPath
	o.MappedPhysicalPath = op.MappedPhysicalPath
	return op
}

func (o *MapPathRequest) xxx_FromOp(ctx context.Context, op *xxx_MapPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConfigPath = op.ConfigPath
	o.MappedPhysicalPath = op.MappedPhysicalPath
}
func (o *MapPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MapPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MapPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MapPathResponse structure represents the MapPath operation response
type MapPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	NewPhysicalPath *oaut.String   `idl:"name:pbstrNewPhysicalPath" json:"new_physical_path"`
	// Return: The MapPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MapPathResponse) xxx_ToOp(ctx context.Context, op *xxx_MapPathOperation) *xxx_MapPathOperation {
	if op == nil {
		op = &xxx_MapPathOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.NewPhysicalPath = op.NewPhysicalPath
	o.Return = op.Return
	return op
}

func (o *MapPathResponse) xxx_FromOp(ctx context.Context, op *xxx_MapPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.NewPhysicalPath = op.NewPhysicalPath
	o.Return = op.Return
}
func (o *MapPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MapPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MapPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
