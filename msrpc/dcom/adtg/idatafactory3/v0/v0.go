package idatafactory3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	idatafactory2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/adtg/idatafactory2/v0"
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
	_ = idatafactory2.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/adtg"
)

var (
	// IDataFactory3 interface identifier 4639db2a-bfc5-11d2-9318-00c04fbbbfb3
	DataFactory3IID = &dcom.IID{Data1: 0x4639db2a, Data2: 0xbfc5, Data3: 0x11d2, Data4: []byte{0x93, 0x18, 0x00, 0xc0, 0x4f, 0xbb, 0xbf, 0xb3}}
	// Syntax UUID
	DataFactory3SyntaxUUID = &uuid.UUID{TimeLow: 0x4639db2a, TimeMid: 0xbfc5, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x18, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xbb, 0xbf, 0xb3}}
	// Syntax ID
	DataFactory3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DataFactory3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IDataFactory3 interface.
type DataFactory3Client interface {

	// IDataFactory2 retrieval method.
	DataFactory2() idatafactory2.DataFactory2Client

	// Execute operation.
	Execute(context.Context, *ExecuteRequest, ...dcerpc.CallOption) (*ExecuteResponse, error)

	// Synchronize operation.
	Synchronize(context.Context, *SynchronizeRequest, ...dcerpc.CallOption) (*SynchronizeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) DataFactory3Client
}

type xxx_DefaultDataFactory3Client struct {
	idatafactory2.DataFactory2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDataFactory3Client) DataFactory2() idatafactory2.DataFactory2Client {
	return o.DataFactory2Client
}

func (o *xxx_DefaultDataFactory3Client) Execute(ctx context.Context, in *ExecuteRequest, opts ...dcerpc.CallOption) (*ExecuteResponse, error) {
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
	out := &ExecuteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataFactory3Client) Synchronize(ctx context.Context, in *SynchronizeRequest, opts ...dcerpc.CallOption) (*SynchronizeResponse, error) {
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
	out := &SynchronizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataFactory3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDataFactory3Client) IPID(ctx context.Context, ipid *dcom.IPID) DataFactory3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDataFactory3Client{
		DataFactory2Client: o.DataFactory2Client.IPID(ctx, ipid),
		cc:                 o.cc,
		ipid:               ipid,
	}
}
func NewDataFactory3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DataFactory3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DataFactory3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idatafactory2.NewDataFactory2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultDataFactory3Client{
		DataFactory2Client: base,
		cc:                 cc,
		ipid:               ipid,
	}, nil
}

// xxx_ExecuteOperation structure represents the Execute operation
type xxx_ExecuteOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConnectionString *oaut.String   `idl:"name:ConnectionString" json:"connection_string"`
	HandlerString    *oaut.String   `idl:"name:HandlerString" json:"handler_string"`
	QueryString      *oaut.String   `idl:"name:QueryString" json:"query_string"`
	MarshalOptions   int32          `idl:"name:lMarshalOptions" json:"marshal_options"`
	Properties       *oaut.Variant  `idl:"name:Properties" json:"properties"`
	TableID          *oaut.Variant  `idl:"name:TableId" json:"table_id"`
	ExecuteOptions   int32          `idl:"name:lExecuteOptions" json:"execute_options"`
	Parameters       *oaut.Variant  `idl:"name:pParameters" json:"parameters"`
	LocaleID         int32          `idl:"name:lcid" json:"locale_id"`
	Information      *oaut.Variant  `idl:"name:pInformation" json:"information"`
	RecordSet        *oaut.Dispatch `idl:"name:ppRecordset" json:"record_set"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ExecuteOperation) OpNum() int { return 9 }

func (o *xxx_ExecuteOperation) OpName() string { return "/IDataFactory3/v0/Execute" }

func (o *xxx_ExecuteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecuteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ConnectionString {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConnectionString != nil {
			_ptr_ConnectionString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConnectionString != nil {
					if err := o.ConnectionString.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConnectionString, _ptr_ConnectionString); err != nil {
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
	// HandlerString {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.HandlerString != nil {
			_ptr_HandlerString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.HandlerString != nil {
					if err := o.HandlerString.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.HandlerString, _ptr_HandlerString); err != nil {
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
	// QueryString {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.QueryString != nil {
			_ptr_QueryString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueryString != nil {
					if err := o.QueryString.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryString, _ptr_QueryString); err != nil {
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
	// lMarshalOptions {in} (1:(int32))
	{
		if err := w.WriteData(o.MarshalOptions); err != nil {
			return err
		}
	}
	// Properties {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Properties != nil {
			if err := o.Properties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TableId {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.TableID != nil {
			if err := o.TableID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lExecuteOptions {in} (1:(int32))
	{
		if err := w.WriteData(o.ExecuteOptions); err != nil {
			return err
		}
	}
	// pParameters {in, out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Parameters != nil {
			_ptr_pParameters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Parameters != nil {
					if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Parameters, _ptr_pParameters); err != nil {
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
	// lcid {in, optional} (1:(int32))
	{
		if err := w.WriteData(o.LocaleID); err != nil {
			return err
		}
	}
	// pInformation {in, out, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Information != nil {
			_ptr_pInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Information != nil {
					if err := o.Information.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Information, _ptr_pInformation); err != nil {
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

func (o *xxx_ExecuteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ConnectionString {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_ConnectionString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConnectionString == nil {
				o.ConnectionString = &oaut.String{}
			}
			if err := o.ConnectionString.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ConnectionString := func(ptr interface{}) { o.ConnectionString = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConnectionString, _s_ConnectionString, _ptr_ConnectionString); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// HandlerString {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_HandlerString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.HandlerString == nil {
				o.HandlerString = &oaut.String{}
			}
			if err := o.HandlerString.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_HandlerString := func(ptr interface{}) { o.HandlerString = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.HandlerString, _s_HandlerString, _ptr_HandlerString); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// QueryString {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_QueryString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueryString == nil {
				o.QueryString = &oaut.String{}
			}
			if err := o.QueryString.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_QueryString := func(ptr interface{}) { o.QueryString = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.QueryString, _s_QueryString, _ptr_QueryString); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lMarshalOptions {in} (1:(int32))
	{
		if err := w.ReadData(&o.MarshalOptions); err != nil {
			return err
		}
	}
	// Properties {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Properties == nil {
			o.Properties = &oaut.Variant{}
		}
		if err := o.Properties.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TableId {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.TableID == nil {
			o.TableID = &oaut.Variant{}
		}
		if err := o.TableID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lExecuteOptions {in} (1:(int32))
	{
		if err := w.ReadData(&o.ExecuteOptions); err != nil {
			return err
		}
	}
	// pParameters {in, out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pParameters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Parameters == nil {
				o.Parameters = &oaut.Variant{}
			}
			if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pParameters := func(ptr interface{}) { o.Parameters = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Parameters, _s_pParameters, _ptr_pParameters); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lcid {in, optional} (1:(int32))
	{
		if err := w.ReadData(&o.LocaleID); err != nil {
			return err
		}
	}
	// pInformation {in, out, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Information == nil {
				o.Information = &oaut.Variant{}
			}
			if err := o.Information.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pInformation := func(ptr interface{}) { o.Information = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Information, _s_pInformation, _ptr_pInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecuteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecuteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pParameters {in, out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Parameters != nil {
			_ptr_pParameters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Parameters != nil {
					if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Parameters, _ptr_pParameters); err != nil {
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
	// pInformation {in, out, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Information != nil {
			_ptr_pInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Information != nil {
					if err := o.Information.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Information, _ptr_pInformation); err != nil {
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
	// ppRecordset {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		if o.RecordSet != nil {
			_ptr_ppRecordset := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RecordSet != nil {
					if err := o.RecordSet.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RecordSet, _ptr_ppRecordset); err != nil {
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

func (o *xxx_ExecuteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pParameters {in, out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pParameters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Parameters == nil {
				o.Parameters = &oaut.Variant{}
			}
			if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pParameters := func(ptr interface{}) { o.Parameters = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Parameters, _s_pParameters, _ptr_pParameters); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pInformation {in, out, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Information == nil {
				o.Information = &oaut.Variant{}
			}
			if err := o.Information.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pInformation := func(ptr interface{}) { o.Information = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Information, _s_pInformation, _ptr_pInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppRecordset {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_ppRecordset := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RecordSet == nil {
				o.RecordSet = &oaut.Dispatch{}
			}
			if err := o.RecordSet.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRecordset := func(ptr interface{}) { o.RecordSet = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.RecordSet, _s_ppRecordset, _ptr_ppRecordset); err != nil {
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

// ExecuteRequest structure represents the Execute operation request
type ExecuteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	ConnectionString *oaut.String   `idl:"name:ConnectionString" json:"connection_string"`
	HandlerString    *oaut.String   `idl:"name:HandlerString" json:"handler_string"`
	QueryString      *oaut.String   `idl:"name:QueryString" json:"query_string"`
	MarshalOptions   int32          `idl:"name:lMarshalOptions" json:"marshal_options"`
	Properties       *oaut.Variant  `idl:"name:Properties" json:"properties"`
	TableID          *oaut.Variant  `idl:"name:TableId" json:"table_id"`
	ExecuteOptions   int32          `idl:"name:lExecuteOptions" json:"execute_options"`
	Parameters       *oaut.Variant  `idl:"name:pParameters" json:"parameters"`
	LocaleID         int32          `idl:"name:lcid" json:"locale_id"`
	Information      *oaut.Variant  `idl:"name:pInformation" json:"information"`
}

func (o *ExecuteRequest) xxx_ToOp(ctx context.Context) *xxx_ExecuteOperation {
	if o == nil {
		return &xxx_ExecuteOperation{}
	}
	return &xxx_ExecuteOperation{
		This:             o.This,
		ConnectionString: o.ConnectionString,
		HandlerString:    o.HandlerString,
		QueryString:      o.QueryString,
		MarshalOptions:   o.MarshalOptions,
		Properties:       o.Properties,
		TableID:          o.TableID,
		ExecuteOptions:   o.ExecuteOptions,
		Parameters:       o.Parameters,
		LocaleID:         o.LocaleID,
		Information:      o.Information,
	}
}

func (o *ExecuteRequest) xxx_FromOp(ctx context.Context, op *xxx_ExecuteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConnectionString = op.ConnectionString
	o.HandlerString = op.HandlerString
	o.QueryString = op.QueryString
	o.MarshalOptions = op.MarshalOptions
	o.Properties = op.Properties
	o.TableID = op.TableID
	o.ExecuteOptions = op.ExecuteOptions
	o.Parameters = op.Parameters
	o.LocaleID = op.LocaleID
	o.Information = op.Information
}
func (o *ExecuteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExecuteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecuteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExecuteResponse structure represents the Execute operation response
type ExecuteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Parameters  *oaut.Variant  `idl:"name:pParameters" json:"parameters"`
	Information *oaut.Variant  `idl:"name:pInformation" json:"information"`
	RecordSet   *oaut.Dispatch `idl:"name:ppRecordset" json:"record_set"`
	// Return: The Execute return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExecuteResponse) xxx_ToOp(ctx context.Context) *xxx_ExecuteOperation {
	if o == nil {
		return &xxx_ExecuteOperation{}
	}
	return &xxx_ExecuteOperation{
		That:        o.That,
		Parameters:  o.Parameters,
		Information: o.Information,
		RecordSet:   o.RecordSet,
		Return:      o.Return,
	}
}

func (o *ExecuteResponse) xxx_FromOp(ctx context.Context, op *xxx_ExecuteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Parameters = op.Parameters
	o.Information = op.Information
	o.RecordSet = op.RecordSet
	o.Return = op.Return
}
func (o *ExecuteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExecuteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecuteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SynchronizeOperation structure represents the Synchronize operation
type xxx_SynchronizeOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConnectionString   *oaut.String   `idl:"name:ConnectionString" json:"connection_string"`
	HandlerString      *oaut.String   `idl:"name:HandlerString" json:"handler_string"`
	SynchronizeOptions int32          `idl:"name:lSynchronizeOptions" json:"synchronize_options"`
	RecordSet          *oaut.Dispatch `idl:"name:ppRecordset" json:"record_set"`
	StatusArray        *oaut.Variant  `idl:"name:pStatusArray" json:"status_array"`
	LocaleID           int32          `idl:"name:lcid" json:"locale_id"`
	Information        *oaut.Variant  `idl:"name:pInformation" json:"information"`
	Result             *oaut.Variant  `idl:"name:pResult" json:"result"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SynchronizeOperation) OpNum() int { return 10 }

func (o *xxx_SynchronizeOperation) OpName() string { return "/IDataFactory3/v0/Synchronize" }

func (o *xxx_SynchronizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SynchronizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ConnectionString {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConnectionString != nil {
			_ptr_ConnectionString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConnectionString != nil {
					if err := o.ConnectionString.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConnectionString, _ptr_ConnectionString); err != nil {
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
	// HandlerString {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.HandlerString != nil {
			_ptr_HandlerString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.HandlerString != nil {
					if err := o.HandlerString.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.HandlerString, _ptr_HandlerString); err != nil {
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
	// lSynchronizeOptions {in} (1:(int32))
	{
		if err := w.WriteData(o.SynchronizeOptions); err != nil {
			return err
		}
	}
	// ppRecordset {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		if o.RecordSet != nil {
			_ptr_ppRecordset := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RecordSet != nil {
					if err := o.RecordSet.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RecordSet, _ptr_ppRecordset); err != nil {
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
	// pStatusArray {in, out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.StatusArray != nil {
			_ptr_pStatusArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.StatusArray != nil {
					if err := o.StatusArray.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.StatusArray, _ptr_pStatusArray); err != nil {
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
	// lcid {in, optional} (1:(int32))
	{
		if err := w.WriteData(o.LocaleID); err != nil {
			return err
		}
	}
	// pInformation {in, out, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Information != nil {
			_ptr_pInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Information != nil {
					if err := o.Information.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Information, _ptr_pInformation); err != nil {
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

func (o *xxx_SynchronizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ConnectionString {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_ConnectionString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConnectionString == nil {
				o.ConnectionString = &oaut.String{}
			}
			if err := o.ConnectionString.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ConnectionString := func(ptr interface{}) { o.ConnectionString = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConnectionString, _s_ConnectionString, _ptr_ConnectionString); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// HandlerString {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_HandlerString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.HandlerString == nil {
				o.HandlerString = &oaut.String{}
			}
			if err := o.HandlerString.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_HandlerString := func(ptr interface{}) { o.HandlerString = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.HandlerString, _s_HandlerString, _ptr_HandlerString); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lSynchronizeOptions {in} (1:(int32))
	{
		if err := w.ReadData(&o.SynchronizeOptions); err != nil {
			return err
		}
	}
	// ppRecordset {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_ppRecordset := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RecordSet == nil {
				o.RecordSet = &oaut.Dispatch{}
			}
			if err := o.RecordSet.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRecordset := func(ptr interface{}) { o.RecordSet = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.RecordSet, _s_ppRecordset, _ptr_ppRecordset); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pStatusArray {in, out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pStatusArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.StatusArray == nil {
				o.StatusArray = &oaut.Variant{}
			}
			if err := o.StatusArray.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pStatusArray := func(ptr interface{}) { o.StatusArray = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.StatusArray, _s_pStatusArray, _ptr_pStatusArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lcid {in, optional} (1:(int32))
	{
		if err := w.ReadData(&o.LocaleID); err != nil {
			return err
		}
	}
	// pInformation {in, out, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Information == nil {
				o.Information = &oaut.Variant{}
			}
			if err := o.Information.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pInformation := func(ptr interface{}) { o.Information = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Information, _s_pInformation, _ptr_pInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SynchronizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SynchronizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppRecordset {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		if o.RecordSet != nil {
			_ptr_ppRecordset := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RecordSet != nil {
					if err := o.RecordSet.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RecordSet, _ptr_ppRecordset); err != nil {
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
	// pStatusArray {in, out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.StatusArray != nil {
			_ptr_pStatusArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.StatusArray != nil {
					if err := o.StatusArray.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.StatusArray, _ptr_pStatusArray); err != nil {
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
	// pInformation {in, out, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Information != nil {
			_ptr_pInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Information != nil {
					if err := o.Information.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Information, _ptr_pInformation); err != nil {
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
	// pResult {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Result != nil {
			_ptr_pResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Result != nil {
					if err := o.Result.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Result, _ptr_pResult); err != nil {
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

func (o *xxx_SynchronizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppRecordset {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_ppRecordset := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RecordSet == nil {
				o.RecordSet = &oaut.Dispatch{}
			}
			if err := o.RecordSet.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRecordset := func(ptr interface{}) { o.RecordSet = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.RecordSet, _s_ppRecordset, _ptr_ppRecordset); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pStatusArray {in, out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pStatusArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.StatusArray == nil {
				o.StatusArray = &oaut.Variant{}
			}
			if err := o.StatusArray.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pStatusArray := func(ptr interface{}) { o.StatusArray = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.StatusArray, _s_pStatusArray, _ptr_pStatusArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pInformation {in, out, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Information == nil {
				o.Information = &oaut.Variant{}
			}
			if err := o.Information.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pInformation := func(ptr interface{}) { o.Information = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Information, _s_pInformation, _ptr_pInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResult {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Result == nil {
				o.Result = &oaut.Variant{}
			}
			if err := o.Result.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pResult := func(ptr interface{}) { o.Result = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Result, _s_pResult, _ptr_pResult); err != nil {
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

// SynchronizeRequest structure represents the Synchronize operation request
type SynchronizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	ConnectionString   *oaut.String   `idl:"name:ConnectionString" json:"connection_string"`
	HandlerString      *oaut.String   `idl:"name:HandlerString" json:"handler_string"`
	SynchronizeOptions int32          `idl:"name:lSynchronizeOptions" json:"synchronize_options"`
	RecordSet          *oaut.Dispatch `idl:"name:ppRecordset" json:"record_set"`
	StatusArray        *oaut.Variant  `idl:"name:pStatusArray" json:"status_array"`
	LocaleID           int32          `idl:"name:lcid" json:"locale_id"`
	Information        *oaut.Variant  `idl:"name:pInformation" json:"information"`
}

func (o *SynchronizeRequest) xxx_ToOp(ctx context.Context) *xxx_SynchronizeOperation {
	if o == nil {
		return &xxx_SynchronizeOperation{}
	}
	return &xxx_SynchronizeOperation{
		This:               o.This,
		ConnectionString:   o.ConnectionString,
		HandlerString:      o.HandlerString,
		SynchronizeOptions: o.SynchronizeOptions,
		RecordSet:          o.RecordSet,
		StatusArray:        o.StatusArray,
		LocaleID:           o.LocaleID,
		Information:        o.Information,
	}
}

func (o *SynchronizeRequest) xxx_FromOp(ctx context.Context, op *xxx_SynchronizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConnectionString = op.ConnectionString
	o.HandlerString = op.HandlerString
	o.SynchronizeOptions = op.SynchronizeOptions
	o.RecordSet = op.RecordSet
	o.StatusArray = op.StatusArray
	o.LocaleID = op.LocaleID
	o.Information = op.Information
}
func (o *SynchronizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SynchronizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SynchronizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SynchronizeResponse structure represents the Synchronize operation response
type SynchronizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	RecordSet   *oaut.Dispatch `idl:"name:ppRecordset" json:"record_set"`
	StatusArray *oaut.Variant  `idl:"name:pStatusArray" json:"status_array"`
	Information *oaut.Variant  `idl:"name:pInformation" json:"information"`
	Result      *oaut.Variant  `idl:"name:pResult" json:"result"`
	// Return: The Synchronize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SynchronizeResponse) xxx_ToOp(ctx context.Context) *xxx_SynchronizeOperation {
	if o == nil {
		return &xxx_SynchronizeOperation{}
	}
	return &xxx_SynchronizeOperation{
		That:        o.That,
		RecordSet:   o.RecordSet,
		StatusArray: o.StatusArray,
		Information: o.Information,
		Result:      o.Result,
		Return:      o.Return,
	}
}

func (o *SynchronizeResponse) xxx_FromOp(ctx context.Context, op *xxx_SynchronizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RecordSet = op.RecordSet
	o.StatusArray = op.StatusArray
	o.Information = op.Information
	o.Result = op.Result
	o.Return = op.Return
}
func (o *SynchronizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SynchronizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SynchronizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
