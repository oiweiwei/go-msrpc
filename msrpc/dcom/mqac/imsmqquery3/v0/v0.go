package imsmqquery3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	mqac "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac"
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
	_ = mqac.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQQuery3 interface identifier eba96b19-2168-11d3-898c-00e02c074f6b
	Query3IID = &dcom.IID{Data1: 0xeba96b19, Data2: 0x2168, Data3: 0x11d3, Data4: []byte{0x89, 0x8c, 0x00, 0xe0, 0x2c, 0x07, 0x4f, 0x6b}}
	// Syntax UUID
	Query3SyntaxUUID = &uuid.UUID{TimeLow: 0xeba96b19, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	// Syntax ID
	Query3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Query3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQQuery3 interface.
type Query3Client interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// LookupQueue_v2 operation.
	LookupQueueV2(context.Context, *LookupQueueV2Request, ...dcerpc.CallOption) (*LookupQueueV2Response, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// LookupQueue operation.
	LookupQueue(context.Context, *LookupQueueRequest, ...dcerpc.CallOption) (*LookupQueueResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Query3Client
}

type xxx_DefaultQuery3Client struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQuery3Client) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultQuery3Client) LookupQueueV2(ctx context.Context, in *LookupQueueV2Request, opts ...dcerpc.CallOption) (*LookupQueueV2Response, error) {
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
	out := &LookupQueueV2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuery3Client) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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

func (o *xxx_DefaultQuery3Client) LookupQueue(ctx context.Context, in *LookupQueueRequest, opts ...dcerpc.CallOption) (*LookupQueueResponse, error) {
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
	out := &LookupQueueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuery3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQuery3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQuery3Client) IPID(ctx context.Context, ipid *dcom.IPID) Query3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQuery3Client{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewQuery3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Query3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Query3SyntaxV0_0))...)
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
	return &xxx_DefaultQuery3Client{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_LookupQueueV2Operation structure represents the LookupQueue_v2 operation
type xxx_LookupQueueV2Operation struct {
	This                *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat    `idl:"name:That" json:"that"`
	QueueGUID           *oaut.Variant     `idl:"name:QueueGuid" json:"queue_guid"`
	ServiceTypeGUID     *oaut.Variant     `idl:"name:ServiceTypeGuid" json:"service_type_guid"`
	Label               *oaut.Variant     `idl:"name:Label" json:"label"`
	CreateTime          *oaut.Variant     `idl:"name:CreateTime" json:"create_time"`
	ModifyTime          *oaut.Variant     `idl:"name:ModifyTime" json:"modify_time"`
	RelationServiceType *oaut.Variant     `idl:"name:RelServiceType" json:"relation_service_type"`
	RelationLabel       *oaut.Variant     `idl:"name:RelLabel" json:"relation_label"`
	RelationCreateTime  *oaut.Variant     `idl:"name:RelCreateTime" json:"relation_create_time"`
	RelationModifyTime  *oaut.Variant     `idl:"name:RelModifyTime" json:"relation_modify_time"`
	S                   *mqac.QueueInfos3 `idl:"name:ppqinfos" json:"s"`
	Return              int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupQueueV2Operation) OpNum() int { return 7 }

func (o *xxx_LookupQueueV2Operation) OpName() string { return "/IMSMQQuery3/v0/LookupQueue_v2" }

func (o *xxx_LookupQueueV2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupQueueV2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// QueueGuid {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.QueueGUID != nil {
			_ptr_QueueGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueueGUID != nil {
					if err := o.QueueGUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueueGUID, _ptr_QueueGuid); err != nil {
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
	// ServiceTypeGuid {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ServiceTypeGUID != nil {
			_ptr_ServiceTypeGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServiceTypeGUID != nil {
					if err := o.ServiceTypeGUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServiceTypeGUID, _ptr_ServiceTypeGuid); err != nil {
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
	// Label {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Label != nil {
			_ptr_Label := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Label != nil {
					if err := o.Label.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Label, _ptr_Label); err != nil {
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
	// CreateTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.CreateTime != nil {
			_ptr_CreateTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CreateTime != nil {
					if err := o.CreateTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CreateTime, _ptr_CreateTime); err != nil {
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
	// ModifyTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ModifyTime != nil {
			_ptr_ModifyTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ModifyTime != nil {
					if err := o.ModifyTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModifyTime, _ptr_ModifyTime); err != nil {
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
	// RelServiceType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.RelationServiceType != nil {
			_ptr_RelServiceType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RelationServiceType != nil {
					if err := o.RelationServiceType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RelationServiceType, _ptr_RelServiceType); err != nil {
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
	// RelLabel {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.RelationLabel != nil {
			_ptr_RelLabel := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RelationLabel != nil {
					if err := o.RelationLabel.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RelationLabel, _ptr_RelLabel); err != nil {
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
	// RelCreateTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.RelationCreateTime != nil {
			_ptr_RelCreateTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RelationCreateTime != nil {
					if err := o.RelationCreateTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RelationCreateTime, _ptr_RelCreateTime); err != nil {
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
	// RelModifyTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.RelationModifyTime != nil {
			_ptr_RelModifyTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RelationModifyTime != nil {
					if err := o.RelationModifyTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RelationModifyTime, _ptr_RelModifyTime); err != nil {
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

func (o *xxx_LookupQueueV2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// QueueGuid {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_QueueGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueueGUID == nil {
				o.QueueGUID = &oaut.Variant{}
			}
			if err := o.QueueGUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_QueueGuid := func(ptr interface{}) { o.QueueGUID = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.QueueGUID, _s_QueueGuid, _ptr_QueueGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ServiceTypeGuid {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ServiceTypeGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServiceTypeGUID == nil {
				o.ServiceTypeGUID = &oaut.Variant{}
			}
			if err := o.ServiceTypeGUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ServiceTypeGuid := func(ptr interface{}) { o.ServiceTypeGUID = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ServiceTypeGUID, _s_ServiceTypeGuid, _ptr_ServiceTypeGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Label {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Label := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Label == nil {
				o.Label = &oaut.Variant{}
			}
			if err := o.Label.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Label := func(ptr interface{}) { o.Label = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Label, _s_Label, _ptr_Label); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// CreateTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_CreateTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CreateTime == nil {
				o.CreateTime = &oaut.Variant{}
			}
			if err := o.CreateTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_CreateTime := func(ptr interface{}) { o.CreateTime = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.CreateTime, _s_CreateTime, _ptr_CreateTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ModifyTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ModifyTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ModifyTime == nil {
				o.ModifyTime = &oaut.Variant{}
			}
			if err := o.ModifyTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ModifyTime := func(ptr interface{}) { o.ModifyTime = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ModifyTime, _s_ModifyTime, _ptr_ModifyTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RelServiceType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_RelServiceType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RelationServiceType == nil {
				o.RelationServiceType = &oaut.Variant{}
			}
			if err := o.RelationServiceType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_RelServiceType := func(ptr interface{}) { o.RelationServiceType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.RelationServiceType, _s_RelServiceType, _ptr_RelServiceType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RelLabel {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_RelLabel := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RelationLabel == nil {
				o.RelationLabel = &oaut.Variant{}
			}
			if err := o.RelationLabel.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_RelLabel := func(ptr interface{}) { o.RelationLabel = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.RelationLabel, _s_RelLabel, _ptr_RelLabel); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RelCreateTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_RelCreateTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RelationCreateTime == nil {
				o.RelationCreateTime = &oaut.Variant{}
			}
			if err := o.RelationCreateTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_RelCreateTime := func(ptr interface{}) { o.RelationCreateTime = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.RelationCreateTime, _s_RelCreateTime, _ptr_RelCreateTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RelModifyTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_RelModifyTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RelationModifyTime == nil {
				o.RelationModifyTime = &oaut.Variant{}
			}
			if err := o.RelationModifyTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_RelModifyTime := func(ptr interface{}) { o.RelationModifyTime = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.RelationModifyTime, _s_RelModifyTime, _ptr_RelModifyTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupQueueV2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupQueueV2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppqinfos {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfos3}(interface))
	{
		if o.S != nil {
			_ptr_ppqinfos := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.S != nil {
					if err := o.S.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfos3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.S, _ptr_ppqinfos); err != nil {
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

func (o *xxx_LookupQueueV2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppqinfos {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfos3}(interface))
	{
		_ptr_ppqinfos := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.S == nil {
				o.S = &mqac.QueueInfos3{}
			}
			if err := o.S.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfos := func(ptr interface{}) { o.S = *ptr.(**mqac.QueueInfos3) }
		if err := w.ReadPointer(&o.S, _s_ppqinfos, _ptr_ppqinfos); err != nil {
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

// LookupQueueV2Request structure represents the LookupQueue_v2 operation request
type LookupQueueV2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	QueueGUID           *oaut.Variant  `idl:"name:QueueGuid" json:"queue_guid"`
	ServiceTypeGUID     *oaut.Variant  `idl:"name:ServiceTypeGuid" json:"service_type_guid"`
	Label               *oaut.Variant  `idl:"name:Label" json:"label"`
	CreateTime          *oaut.Variant  `idl:"name:CreateTime" json:"create_time"`
	ModifyTime          *oaut.Variant  `idl:"name:ModifyTime" json:"modify_time"`
	RelationServiceType *oaut.Variant  `idl:"name:RelServiceType" json:"relation_service_type"`
	RelationLabel       *oaut.Variant  `idl:"name:RelLabel" json:"relation_label"`
	RelationCreateTime  *oaut.Variant  `idl:"name:RelCreateTime" json:"relation_create_time"`
	RelationModifyTime  *oaut.Variant  `idl:"name:RelModifyTime" json:"relation_modify_time"`
}

func (o *LookupQueueV2Request) xxx_ToOp(ctx context.Context, op *xxx_LookupQueueV2Operation) *xxx_LookupQueueV2Operation {
	if op == nil {
		op = &xxx_LookupQueueV2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.QueueGUID = o.QueueGUID
	op.ServiceTypeGUID = o.ServiceTypeGUID
	op.Label = o.Label
	op.CreateTime = o.CreateTime
	op.ModifyTime = o.ModifyTime
	op.RelationServiceType = o.RelationServiceType
	op.RelationLabel = o.RelationLabel
	op.RelationCreateTime = o.RelationCreateTime
	op.RelationModifyTime = o.RelationModifyTime
	return op
}

func (o *LookupQueueV2Request) xxx_FromOp(ctx context.Context, op *xxx_LookupQueueV2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QueueGUID = op.QueueGUID
	o.ServiceTypeGUID = op.ServiceTypeGUID
	o.Label = op.Label
	o.CreateTime = op.CreateTime
	o.ModifyTime = op.ModifyTime
	o.RelationServiceType = op.RelationServiceType
	o.RelationLabel = op.RelationLabel
	o.RelationCreateTime = op.RelationCreateTime
	o.RelationModifyTime = op.RelationModifyTime
}
func (o *LookupQueueV2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupQueueV2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupQueueV2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupQueueV2Response structure represents the LookupQueue_v2 operation response
type LookupQueueV2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat    `idl:"name:That" json:"that"`
	S    *mqac.QueueInfos3 `idl:"name:ppqinfos" json:"s"`
	// Return: The LookupQueue_v2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupQueueV2Response) xxx_ToOp(ctx context.Context, op *xxx_LookupQueueV2Operation) *xxx_LookupQueueV2Operation {
	if op == nil {
		op = &xxx_LookupQueueV2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.S = o.S
	op.Return = o.Return
	return op
}

func (o *LookupQueueV2Response) xxx_FromOp(ctx context.Context, op *xxx_LookupQueueV2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.S = op.S
	o.Return = op.Return
}
func (o *LookupQueueV2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupQueueV2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupQueueV2Operation{}
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

func (o *xxx_GetPropertiesOperation) OpNum() int { return 8 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/IMSMQQuery3/v0/Properties" }

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

// xxx_LookupQueueOperation structure represents the LookupQueue operation
type xxx_LookupQueueOperation struct {
	This                     *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat    `idl:"name:That" json:"that"`
	QueueGUID                *oaut.Variant     `idl:"name:QueueGuid" json:"queue_guid"`
	ServiceTypeGUID          *oaut.Variant     `idl:"name:ServiceTypeGuid" json:"service_type_guid"`
	Label                    *oaut.Variant     `idl:"name:Label" json:"label"`
	CreateTime               *oaut.Variant     `idl:"name:CreateTime" json:"create_time"`
	ModifyTime               *oaut.Variant     `idl:"name:ModifyTime" json:"modify_time"`
	RelationServiceType      *oaut.Variant     `idl:"name:RelServiceType" json:"relation_service_type"`
	RelationLabel            *oaut.Variant     `idl:"name:RelLabel" json:"relation_label"`
	RelationCreateTime       *oaut.Variant     `idl:"name:RelCreateTime" json:"relation_create_time"`
	RelationModifyTime       *oaut.Variant     `idl:"name:RelModifyTime" json:"relation_modify_time"`
	MulticastAddress         *oaut.Variant     `idl:"name:MulticastAddress" json:"multicast_address"`
	RelationMulticastAddress *oaut.Variant     `idl:"name:RelMulticastAddress" json:"relation_multicast_address"`
	S                        *mqac.QueueInfos3 `idl:"name:ppqinfos" json:"s"`
	Return                   int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupQueueOperation) OpNum() int { return 9 }

func (o *xxx_LookupQueueOperation) OpName() string { return "/IMSMQQuery3/v0/LookupQueue" }

func (o *xxx_LookupQueueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupQueueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// QueueGuid {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.QueueGUID != nil {
			_ptr_QueueGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueueGUID != nil {
					if err := o.QueueGUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueueGUID, _ptr_QueueGuid); err != nil {
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
	// ServiceTypeGuid {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ServiceTypeGUID != nil {
			_ptr_ServiceTypeGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServiceTypeGUID != nil {
					if err := o.ServiceTypeGUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServiceTypeGUID, _ptr_ServiceTypeGuid); err != nil {
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
	// Label {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Label != nil {
			_ptr_Label := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Label != nil {
					if err := o.Label.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Label, _ptr_Label); err != nil {
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
	// CreateTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.CreateTime != nil {
			_ptr_CreateTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CreateTime != nil {
					if err := o.CreateTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CreateTime, _ptr_CreateTime); err != nil {
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
	// ModifyTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ModifyTime != nil {
			_ptr_ModifyTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ModifyTime != nil {
					if err := o.ModifyTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModifyTime, _ptr_ModifyTime); err != nil {
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
	// RelServiceType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.RelationServiceType != nil {
			_ptr_RelServiceType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RelationServiceType != nil {
					if err := o.RelationServiceType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RelationServiceType, _ptr_RelServiceType); err != nil {
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
	// RelLabel {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.RelationLabel != nil {
			_ptr_RelLabel := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RelationLabel != nil {
					if err := o.RelationLabel.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RelationLabel, _ptr_RelLabel); err != nil {
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
	// RelCreateTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.RelationCreateTime != nil {
			_ptr_RelCreateTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RelationCreateTime != nil {
					if err := o.RelationCreateTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RelationCreateTime, _ptr_RelCreateTime); err != nil {
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
	// RelModifyTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.RelationModifyTime != nil {
			_ptr_RelModifyTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RelationModifyTime != nil {
					if err := o.RelationModifyTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RelationModifyTime, _ptr_RelModifyTime); err != nil {
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
	// MulticastAddress {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.MulticastAddress != nil {
			_ptr_MulticastAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MulticastAddress != nil {
					if err := o.MulticastAddress.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MulticastAddress, _ptr_MulticastAddress); err != nil {
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
	// RelMulticastAddress {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.RelationMulticastAddress != nil {
			_ptr_RelMulticastAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RelationMulticastAddress != nil {
					if err := o.RelationMulticastAddress.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RelationMulticastAddress, _ptr_RelMulticastAddress); err != nil {
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

func (o *xxx_LookupQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// QueueGuid {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_QueueGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueueGUID == nil {
				o.QueueGUID = &oaut.Variant{}
			}
			if err := o.QueueGUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_QueueGuid := func(ptr interface{}) { o.QueueGUID = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.QueueGUID, _s_QueueGuid, _ptr_QueueGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ServiceTypeGuid {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ServiceTypeGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServiceTypeGUID == nil {
				o.ServiceTypeGUID = &oaut.Variant{}
			}
			if err := o.ServiceTypeGUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ServiceTypeGuid := func(ptr interface{}) { o.ServiceTypeGUID = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ServiceTypeGUID, _s_ServiceTypeGuid, _ptr_ServiceTypeGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Label {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Label := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Label == nil {
				o.Label = &oaut.Variant{}
			}
			if err := o.Label.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Label := func(ptr interface{}) { o.Label = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Label, _s_Label, _ptr_Label); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// CreateTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_CreateTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CreateTime == nil {
				o.CreateTime = &oaut.Variant{}
			}
			if err := o.CreateTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_CreateTime := func(ptr interface{}) { o.CreateTime = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.CreateTime, _s_CreateTime, _ptr_CreateTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ModifyTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ModifyTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ModifyTime == nil {
				o.ModifyTime = &oaut.Variant{}
			}
			if err := o.ModifyTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ModifyTime := func(ptr interface{}) { o.ModifyTime = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ModifyTime, _s_ModifyTime, _ptr_ModifyTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RelServiceType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_RelServiceType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RelationServiceType == nil {
				o.RelationServiceType = &oaut.Variant{}
			}
			if err := o.RelationServiceType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_RelServiceType := func(ptr interface{}) { o.RelationServiceType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.RelationServiceType, _s_RelServiceType, _ptr_RelServiceType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RelLabel {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_RelLabel := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RelationLabel == nil {
				o.RelationLabel = &oaut.Variant{}
			}
			if err := o.RelationLabel.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_RelLabel := func(ptr interface{}) { o.RelationLabel = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.RelationLabel, _s_RelLabel, _ptr_RelLabel); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RelCreateTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_RelCreateTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RelationCreateTime == nil {
				o.RelationCreateTime = &oaut.Variant{}
			}
			if err := o.RelationCreateTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_RelCreateTime := func(ptr interface{}) { o.RelationCreateTime = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.RelationCreateTime, _s_RelCreateTime, _ptr_RelCreateTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RelModifyTime {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_RelModifyTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RelationModifyTime == nil {
				o.RelationModifyTime = &oaut.Variant{}
			}
			if err := o.RelationModifyTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_RelModifyTime := func(ptr interface{}) { o.RelationModifyTime = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.RelationModifyTime, _s_RelModifyTime, _ptr_RelModifyTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MulticastAddress {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_MulticastAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MulticastAddress == nil {
				o.MulticastAddress = &oaut.Variant{}
			}
			if err := o.MulticastAddress.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_MulticastAddress := func(ptr interface{}) { o.MulticastAddress = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.MulticastAddress, _s_MulticastAddress, _ptr_MulticastAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RelMulticastAddress {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_RelMulticastAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RelationMulticastAddress == nil {
				o.RelationMulticastAddress = &oaut.Variant{}
			}
			if err := o.RelationMulticastAddress.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_RelMulticastAddress := func(ptr interface{}) { o.RelationMulticastAddress = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.RelationMulticastAddress, _s_RelMulticastAddress, _ptr_RelMulticastAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupQueueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupQueueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppqinfos {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfos3}(interface))
	{
		if o.S != nil {
			_ptr_ppqinfos := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.S != nil {
					if err := o.S.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfos3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.S, _ptr_ppqinfos); err != nil {
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

func (o *xxx_LookupQueueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppqinfos {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfos3}(interface))
	{
		_ptr_ppqinfos := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.S == nil {
				o.S = &mqac.QueueInfos3{}
			}
			if err := o.S.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfos := func(ptr interface{}) { o.S = *ptr.(**mqac.QueueInfos3) }
		if err := w.ReadPointer(&o.S, _s_ppqinfos, _ptr_ppqinfos); err != nil {
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

// LookupQueueRequest structure represents the LookupQueue operation request
type LookupQueueRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                     *dcom.ORPCThis `idl:"name:This" json:"this"`
	QueueGUID                *oaut.Variant  `idl:"name:QueueGuid" json:"queue_guid"`
	ServiceTypeGUID          *oaut.Variant  `idl:"name:ServiceTypeGuid" json:"service_type_guid"`
	Label                    *oaut.Variant  `idl:"name:Label" json:"label"`
	CreateTime               *oaut.Variant  `idl:"name:CreateTime" json:"create_time"`
	ModifyTime               *oaut.Variant  `idl:"name:ModifyTime" json:"modify_time"`
	RelationServiceType      *oaut.Variant  `idl:"name:RelServiceType" json:"relation_service_type"`
	RelationLabel            *oaut.Variant  `idl:"name:RelLabel" json:"relation_label"`
	RelationCreateTime       *oaut.Variant  `idl:"name:RelCreateTime" json:"relation_create_time"`
	RelationModifyTime       *oaut.Variant  `idl:"name:RelModifyTime" json:"relation_modify_time"`
	MulticastAddress         *oaut.Variant  `idl:"name:MulticastAddress" json:"multicast_address"`
	RelationMulticastAddress *oaut.Variant  `idl:"name:RelMulticastAddress" json:"relation_multicast_address"`
}

func (o *LookupQueueRequest) xxx_ToOp(ctx context.Context, op *xxx_LookupQueueOperation) *xxx_LookupQueueOperation {
	if op == nil {
		op = &xxx_LookupQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.QueueGUID = o.QueueGUID
	op.ServiceTypeGUID = o.ServiceTypeGUID
	op.Label = o.Label
	op.CreateTime = o.CreateTime
	op.ModifyTime = o.ModifyTime
	op.RelationServiceType = o.RelationServiceType
	op.RelationLabel = o.RelationLabel
	op.RelationCreateTime = o.RelationCreateTime
	op.RelationModifyTime = o.RelationModifyTime
	op.MulticastAddress = o.MulticastAddress
	op.RelationMulticastAddress = o.RelationMulticastAddress
	return op
}

func (o *LookupQueueRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupQueueOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QueueGUID = op.QueueGUID
	o.ServiceTypeGUID = op.ServiceTypeGUID
	o.Label = op.Label
	o.CreateTime = op.CreateTime
	o.ModifyTime = op.ModifyTime
	o.RelationServiceType = op.RelationServiceType
	o.RelationLabel = op.RelationLabel
	o.RelationCreateTime = op.RelationCreateTime
	o.RelationModifyTime = op.RelationModifyTime
	o.MulticastAddress = op.MulticastAddress
	o.RelationMulticastAddress = op.RelationMulticastAddress
}
func (o *LookupQueueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupQueueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupQueueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupQueueResponse structure represents the LookupQueue operation response
type LookupQueueResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat    `idl:"name:That" json:"that"`
	S    *mqac.QueueInfos3 `idl:"name:ppqinfos" json:"s"`
	// Return: The LookupQueue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupQueueResponse) xxx_ToOp(ctx context.Context, op *xxx_LookupQueueOperation) *xxx_LookupQueueOperation {
	if op == nil {
		op = &xxx_LookupQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.S = o.S
	op.Return = o.Return
	return op
}

func (o *LookupQueueResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupQueueOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.S = op.S
	o.Return = op.Return
}
func (o *LookupQueueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupQueueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupQueueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
