package ifsrmquotabase

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
	ifsrmobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmobject/v0"
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
	_ = ifsrmobject.GoPackage
	_ = oaut.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmQuotaBase interface identifier 1568a795-3924-4118-b74b-68d8f0fa5daf
	QuotaBaseIID = &dcom.IID{Data1: 0x1568a795, Data2: 0x3924, Data3: 0x4118, Data4: []byte{0xb7, 0x4b, 0x68, 0xd8, 0xf0, 0xfa, 0x5d, 0xaf}}
	// Syntax UUID
	QuotaBaseSyntaxUUID = &uuid.UUID{TimeLow: 0x1568a795, TimeMid: 0x3924, TimeHiAndVersion: 0x4118, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0x4b, Node: [6]uint8{0x68, 0xd8, 0xf0, 0xfa, 0x5d, 0xaf}}
	// Syntax ID
	QuotaBaseSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: QuotaBaseSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmQuotaBase interface.
type QuotaBaseClient interface {

	// IFsrmObject retrieval method.
	Object() ifsrmobject.ObjectClient

	// QuotaLimit operation.
	GetQuotaLimit(context.Context, *GetQuotaLimitRequest, ...dcerpc.CallOption) (*GetQuotaLimitResponse, error)

	// QuotaLimit operation.
	SetQuotaLimit(context.Context, *SetQuotaLimitRequest, ...dcerpc.CallOption) (*SetQuotaLimitResponse, error)

	// QuotaFlags operation.
	GetQuotaFlags(context.Context, *GetQuotaFlagsRequest, ...dcerpc.CallOption) (*GetQuotaFlagsResponse, error)

	// QuotaFlags operation.
	SetQuotaFlags(context.Context, *SetQuotaFlagsRequest, ...dcerpc.CallOption) (*SetQuotaFlagsResponse, error)

	// The Thresholds (get) method returns the thresholds for the quota object.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------------+
	//	|         RETURN          |                                   |
	//	|       VALUE/CODE        |            DESCRIPTION            |
	//	|                         |                                   |
	//	+-------------------------+-----------------------------------+
	//	+-------------------------+-----------------------------------+
	//	| 0x80070057 E_INVALIDARG | The thresholds parameter is NULL. |
	//	+-------------------------+-----------------------------------+
	GetThresholds(context.Context, *GetThresholdsRequest, ...dcerpc.CallOption) (*GetThresholdsResponse, error)

	// The AddThreshold method adds a threshold to the quota objects list of thresholds.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------+
	//	|              RETURN              |                                                                            |
	//	|            VALUE/CODE            |                                DESCRIPTION                                 |
	//	|                                  |                                                                            |
	//	+----------------------------------+----------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | The object being created already exists.                                   |
	//	+----------------------------------+----------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE   | The content of the threshold parameter is less than 1 or greater than 250. |
	//	+----------------------------------+----------------------------------------------------------------------------+
	AddThreshold(context.Context, *AddThresholdRequest, ...dcerpc.CallOption) (*AddThresholdResponse, error)

	// The DeleteThreshold method deletes a threshold from the quota objects list of thresholds.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND    | No threshold was found on the quota object equal to the value specified in the   |
	//	|                                | threshold parameter.                                                             |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE | The content of the threshold parameter is less than 1 or greater than 250.       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	DeleteThreshold(context.Context, *DeleteThresholdRequest, ...dcerpc.CallOption) (*DeleteThresholdResponse, error)

	// The ModifyThreshold method modifies the disk usage percentage of an existing threshold
	// in the quota object's list of thresholds.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND      | No threshold was found on the quota object equal to the value specified in the   |
	//	|                                  | threshold parameter.                                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | A threshold equal to the newThreshold parameter already exists for the quota     |
	//	|                                  | object.                                                                          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE   | The content of the threshold parameter is less than 1 or greater than 250.       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	ModifyThreshold(context.Context, *ModifyThresholdRequest, ...dcerpc.CallOption) (*ModifyThresholdResponse, error)

	// The CreateThresholdAction method creates an action and associates it with the specified
	// threshold.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND      | An object with the specified threshold could not be found.                       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | An action of the same type as the actionType parameter already exists for the    |
	//	|                                  | threshold.                                                                       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE   | The content of the threshold parameter is less than 1 or greater than 250.       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The action parameter is NULL.   |
	//	|                                  | The actionsType parameter is not a valid type.                                   |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	CreateThresholdAction(context.Context, *CreateThresholdActionRequest, ...dcerpc.CallOption) (*CreateThresholdActionResponse, error)

	// The EnumThresholdActions method enumerates all the actions for the specified threshold.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+----------------------------------------------------------------------------+
	//	|             RETURN             |                                                                            |
	//	|           VALUE/CODE           |                                DESCRIPTION                                 |
	//	|                                |                                                                            |
	//	+--------------------------------+----------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND    | An object with the specified threshold could not be found.                 |
	//	+--------------------------------+----------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE | The content of the threshold parameter is less than 1 or greater than 250. |
	//	+--------------------------------+----------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | The actions parameter is NULL.                                             |
	//	+--------------------------------+----------------------------------------------------------------------------+
	EnumThresholdActions(context.Context, *EnumThresholdActionsRequest, ...dcerpc.CallOption) (*EnumThresholdActionsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) QuotaBaseClient
}

type xxx_DefaultQuotaBaseClient struct {
	ifsrmobject.ObjectClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQuotaBaseClient) Object() ifsrmobject.ObjectClient {
	return o.ObjectClient
}

func (o *xxx_DefaultQuotaBaseClient) GetQuotaLimit(ctx context.Context, in *GetQuotaLimitRequest, opts ...dcerpc.CallOption) (*GetQuotaLimitResponse, error) {
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
	out := &GetQuotaLimitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaBaseClient) SetQuotaLimit(ctx context.Context, in *SetQuotaLimitRequest, opts ...dcerpc.CallOption) (*SetQuotaLimitResponse, error) {
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
	out := &SetQuotaLimitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaBaseClient) GetQuotaFlags(ctx context.Context, in *GetQuotaFlagsRequest, opts ...dcerpc.CallOption) (*GetQuotaFlagsResponse, error) {
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
	out := &GetQuotaFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaBaseClient) SetQuotaFlags(ctx context.Context, in *SetQuotaFlagsRequest, opts ...dcerpc.CallOption) (*SetQuotaFlagsResponse, error) {
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
	out := &SetQuotaFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaBaseClient) GetThresholds(ctx context.Context, in *GetThresholdsRequest, opts ...dcerpc.CallOption) (*GetThresholdsResponse, error) {
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
	out := &GetThresholdsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaBaseClient) AddThreshold(ctx context.Context, in *AddThresholdRequest, opts ...dcerpc.CallOption) (*AddThresholdResponse, error) {
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
	out := &AddThresholdResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaBaseClient) DeleteThreshold(ctx context.Context, in *DeleteThresholdRequest, opts ...dcerpc.CallOption) (*DeleteThresholdResponse, error) {
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
	out := &DeleteThresholdResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaBaseClient) ModifyThreshold(ctx context.Context, in *ModifyThresholdRequest, opts ...dcerpc.CallOption) (*ModifyThresholdResponse, error) {
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
	out := &ModifyThresholdResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaBaseClient) CreateThresholdAction(ctx context.Context, in *CreateThresholdActionRequest, opts ...dcerpc.CallOption) (*CreateThresholdActionResponse, error) {
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
	out := &CreateThresholdActionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaBaseClient) EnumThresholdActions(ctx context.Context, in *EnumThresholdActionsRequest, opts ...dcerpc.CallOption) (*EnumThresholdActionsResponse, error) {
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
	out := &EnumThresholdActionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaBaseClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQuotaBaseClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQuotaBaseClient) IPID(ctx context.Context, ipid *dcom.IPID) QuotaBaseClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQuotaBaseClient{
		ObjectClient: o.ObjectClient.IPID(ctx, ipid),
		cc:           o.cc,
		ipid:         ipid,
	}
}

func NewQuotaBaseClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QuotaBaseClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QuotaBaseSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmobject.NewObjectClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultQuotaBaseClient{
		ObjectClient: base,
		cc:           cc,
		ipid:         ipid,
	}, nil
}

// xxx_GetQuotaLimitOperation structure represents the QuotaLimit operation
type xxx_GetQuotaLimitOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	QuotaLimit *oaut.Variant  `idl:"name:quotaLimit" json:"quota_limit"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQuotaLimitOperation) OpNum() int { return 12 }

func (o *xxx_GetQuotaLimitOperation) OpName() string { return "/IFsrmQuotaBase/v0/QuotaLimit" }

func (o *xxx_GetQuotaLimitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaLimitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetQuotaLimitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetQuotaLimitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaLimitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quotaLimit {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.QuotaLimit != nil {
			_ptr_quotaLimit := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QuotaLimit != nil {
					if err := o.QuotaLimit.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QuotaLimit, _ptr_quotaLimit); err != nil {
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

func (o *xxx_GetQuotaLimitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quotaLimit {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_quotaLimit := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QuotaLimit == nil {
				o.QuotaLimit = &oaut.Variant{}
			}
			if err := o.QuotaLimit.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quotaLimit := func(ptr interface{}) { o.QuotaLimit = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.QuotaLimit, _s_quotaLimit, _ptr_quotaLimit); err != nil {
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

// GetQuotaLimitRequest structure represents the QuotaLimit operation request
type GetQuotaLimitRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetQuotaLimitRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaLimitOperation) *xxx_GetQuotaLimitOperation {
	if op == nil {
		op = &xxx_GetQuotaLimitOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetQuotaLimitRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaLimitOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetQuotaLimitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQuotaLimitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaLimitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQuotaLimitResponse structure represents the QuotaLimit operation response
type GetQuotaLimitResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	QuotaLimit *oaut.Variant  `idl:"name:quotaLimit" json:"quota_limit"`
	// Return: The QuotaLimit return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetQuotaLimitResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaLimitOperation) *xxx_GetQuotaLimitOperation {
	if op == nil {
		op = &xxx_GetQuotaLimitOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.QuotaLimit = op.QuotaLimit
	o.Return = op.Return
	return op
}

func (o *GetQuotaLimitResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaLimitOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.QuotaLimit = op.QuotaLimit
	o.Return = op.Return
}
func (o *GetQuotaLimitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQuotaLimitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaLimitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetQuotaLimitOperation structure represents the QuotaLimit operation
type xxx_SetQuotaLimitOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	QuotaLimit *oaut.Variant  `idl:"name:quotaLimit" json:"quota_limit"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetQuotaLimitOperation) OpNum() int { return 13 }

func (o *xxx_SetQuotaLimitOperation) OpName() string { return "/IFsrmQuotaBase/v0/QuotaLimit" }

func (o *xxx_SetQuotaLimitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuotaLimitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// quotaLimit {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.QuotaLimit != nil {
			if err := o.QuotaLimit.MarshalNDR(ctx, w); err != nil {
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
	return nil
}

func (o *xxx_SetQuotaLimitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// quotaLimit {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.QuotaLimit == nil {
			o.QuotaLimit = &oaut.Variant{}
		}
		if err := o.QuotaLimit.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuotaLimitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuotaLimitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetQuotaLimitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetQuotaLimitRequest structure represents the QuotaLimit operation request
type SetQuotaLimitRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	QuotaLimit *oaut.Variant  `idl:"name:quotaLimit" json:"quota_limit"`
}

func (o *SetQuotaLimitRequest) xxx_ToOp(ctx context.Context, op *xxx_SetQuotaLimitOperation) *xxx_SetQuotaLimitOperation {
	if op == nil {
		op = &xxx_SetQuotaLimitOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.QuotaLimit = op.QuotaLimit
	return op
}

func (o *SetQuotaLimitRequest) xxx_FromOp(ctx context.Context, op *xxx_SetQuotaLimitOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QuotaLimit = op.QuotaLimit
}
func (o *SetQuotaLimitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetQuotaLimitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetQuotaLimitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetQuotaLimitResponse structure represents the QuotaLimit operation response
type SetQuotaLimitResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The QuotaLimit return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetQuotaLimitResponse) xxx_ToOp(ctx context.Context, op *xxx_SetQuotaLimitOperation) *xxx_SetQuotaLimitOperation {
	if op == nil {
		op = &xxx_SetQuotaLimitOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *SetQuotaLimitResponse) xxx_FromOp(ctx context.Context, op *xxx_SetQuotaLimitOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetQuotaLimitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetQuotaLimitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetQuotaLimitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetQuotaFlagsOperation structure represents the QuotaFlags operation
type xxx_GetQuotaFlagsOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	QuotaFlags int32          `idl:"name:quotaFlags" json:"quota_flags"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQuotaFlagsOperation) OpNum() int { return 14 }

func (o *xxx_GetQuotaFlagsOperation) OpName() string { return "/IFsrmQuotaBase/v0/QuotaFlags" }

func (o *xxx_GetQuotaFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetQuotaFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetQuotaFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quotaFlags {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.QuotaFlags); err != nil {
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

func (o *xxx_GetQuotaFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quotaFlags {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.QuotaFlags); err != nil {
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

// GetQuotaFlagsRequest structure represents the QuotaFlags operation request
type GetQuotaFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetQuotaFlagsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaFlagsOperation) *xxx_GetQuotaFlagsOperation {
	if op == nil {
		op = &xxx_GetQuotaFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetQuotaFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetQuotaFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQuotaFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQuotaFlagsResponse structure represents the QuotaFlags operation response
type GetQuotaFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	QuotaFlags int32          `idl:"name:quotaFlags" json:"quota_flags"`
	// Return: The QuotaFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetQuotaFlagsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaFlagsOperation) *xxx_GetQuotaFlagsOperation {
	if op == nil {
		op = &xxx_GetQuotaFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.QuotaFlags = op.QuotaFlags
	o.Return = op.Return
	return op
}

func (o *GetQuotaFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.QuotaFlags = op.QuotaFlags
	o.Return = op.Return
}
func (o *GetQuotaFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQuotaFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetQuotaFlagsOperation structure represents the QuotaFlags operation
type xxx_SetQuotaFlagsOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	QuotaFlags int32          `idl:"name:quotaFlags" json:"quota_flags"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetQuotaFlagsOperation) OpNum() int { return 15 }

func (o *xxx_SetQuotaFlagsOperation) OpName() string { return "/IFsrmQuotaBase/v0/QuotaFlags" }

func (o *xxx_SetQuotaFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuotaFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// quotaFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.QuotaFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuotaFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// quotaFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.QuotaFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuotaFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuotaFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetQuotaFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetQuotaFlagsRequest structure represents the QuotaFlags operation request
type SetQuotaFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	QuotaFlags int32          `idl:"name:quotaFlags" json:"quota_flags"`
}

func (o *SetQuotaFlagsRequest) xxx_ToOp(ctx context.Context, op *xxx_SetQuotaFlagsOperation) *xxx_SetQuotaFlagsOperation {
	if op == nil {
		op = &xxx_SetQuotaFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.QuotaFlags = op.QuotaFlags
	return op
}

func (o *SetQuotaFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetQuotaFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QuotaFlags = op.QuotaFlags
}
func (o *SetQuotaFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetQuotaFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetQuotaFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetQuotaFlagsResponse structure represents the QuotaFlags operation response
type SetQuotaFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The QuotaFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetQuotaFlagsResponse) xxx_ToOp(ctx context.Context, op *xxx_SetQuotaFlagsOperation) *xxx_SetQuotaFlagsOperation {
	if op == nil {
		op = &xxx_SetQuotaFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *SetQuotaFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetQuotaFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetQuotaFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetQuotaFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetQuotaFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetThresholdsOperation structure represents the Thresholds operation
type xxx_GetThresholdsOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Thresholds *oaut.SafeArray `idl:"name:thresholds" json:"thresholds"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetThresholdsOperation) OpNum() int { return 16 }

func (o *xxx_GetThresholdsOperation) OpName() string { return "/IFsrmQuotaBase/v0/Thresholds" }

func (o *xxx_GetThresholdsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetThresholdsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetThresholdsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetThresholdsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetThresholdsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// thresholds {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Thresholds != nil {
			_ptr_thresholds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Thresholds != nil {
					if err := o.Thresholds.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Thresholds, _ptr_thresholds); err != nil {
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

func (o *xxx_GetThresholdsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// thresholds {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_thresholds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Thresholds == nil {
				o.Thresholds = &oaut.SafeArray{}
			}
			if err := o.Thresholds.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_thresholds := func(ptr interface{}) { o.Thresholds = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Thresholds, _s_thresholds, _ptr_thresholds); err != nil {
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

// GetThresholdsRequest structure represents the Thresholds operation request
type GetThresholdsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetThresholdsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetThresholdsOperation) *xxx_GetThresholdsOperation {
	if op == nil {
		op = &xxx_GetThresholdsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetThresholdsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetThresholdsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetThresholdsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetThresholdsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetThresholdsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetThresholdsResponse structure represents the Thresholds operation response
type GetThresholdsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// thresholds: Pointer to a SAFEARRAY that upon completion contains all the threshold
	// values for the quota object. The caller MUST release the SAFEARRAY when it is done
	// with it.
	Thresholds *oaut.SafeArray `idl:"name:thresholds" json:"thresholds"`
	// Return: The Thresholds return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetThresholdsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetThresholdsOperation) *xxx_GetThresholdsOperation {
	if op == nil {
		op = &xxx_GetThresholdsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Thresholds = op.Thresholds
	o.Return = op.Return
	return op
}

func (o *GetThresholdsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetThresholdsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Thresholds = op.Thresholds
	o.Return = op.Return
}
func (o *GetThresholdsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetThresholdsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetThresholdsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddThresholdOperation structure represents the AddThreshold operation
type xxx_AddThresholdOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Threshold int32          `idl:"name:threshold" json:"threshold"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AddThresholdOperation) OpNum() int { return 17 }

func (o *xxx_AddThresholdOperation) OpName() string { return "/IFsrmQuotaBase/v0/AddThreshold" }

func (o *xxx_AddThresholdOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddThresholdOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// threshold {in} (1:{alias=FSRM_QUOTA_THRESHOLD}(int32))
	{
		if err := w.WriteData(o.Threshold); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddThresholdOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// threshold {in} (1:{alias=FSRM_QUOTA_THRESHOLD}(int32))
	{
		if err := w.ReadData(&o.Threshold); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddThresholdOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddThresholdOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddThresholdOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddThresholdRequest structure represents the AddThreshold operation request
type AddThresholdRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// threshold: Contains the value of the threshold to add to the quota object.
	Threshold int32 `idl:"name:threshold" json:"threshold"`
}

func (o *AddThresholdRequest) xxx_ToOp(ctx context.Context, op *xxx_AddThresholdOperation) *xxx_AddThresholdOperation {
	if op == nil {
		op = &xxx_AddThresholdOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Threshold = op.Threshold
	return op
}

func (o *AddThresholdRequest) xxx_FromOp(ctx context.Context, op *xxx_AddThresholdOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Threshold = op.Threshold
}
func (o *AddThresholdRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddThresholdRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddThresholdOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddThresholdResponse structure represents the AddThreshold operation response
type AddThresholdResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AddThreshold return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddThresholdResponse) xxx_ToOp(ctx context.Context, op *xxx_AddThresholdOperation) *xxx_AddThresholdOperation {
	if op == nil {
		op = &xxx_AddThresholdOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *AddThresholdResponse) xxx_FromOp(ctx context.Context, op *xxx_AddThresholdOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddThresholdResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddThresholdResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddThresholdOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteThresholdOperation structure represents the DeleteThreshold operation
type xxx_DeleteThresholdOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Threshold int32          `idl:"name:threshold" json:"threshold"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteThresholdOperation) OpNum() int { return 18 }

func (o *xxx_DeleteThresholdOperation) OpName() string { return "/IFsrmQuotaBase/v0/DeleteThreshold" }

func (o *xxx_DeleteThresholdOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteThresholdOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// threshold {in} (1:{alias=FSRM_QUOTA_THRESHOLD}(int32))
	{
		if err := w.WriteData(o.Threshold); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteThresholdOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// threshold {in} (1:{alias=FSRM_QUOTA_THRESHOLD}(int32))
	{
		if err := w.ReadData(&o.Threshold); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteThresholdOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteThresholdOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteThresholdOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteThresholdRequest structure represents the DeleteThreshold operation request
type DeleteThresholdRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// threshold: Contains the value of the threshold to delete from the quota object.
	Threshold int32 `idl:"name:threshold" json:"threshold"`
}

func (o *DeleteThresholdRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteThresholdOperation) *xxx_DeleteThresholdOperation {
	if op == nil {
		op = &xxx_DeleteThresholdOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Threshold = op.Threshold
	return op
}

func (o *DeleteThresholdRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteThresholdOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Threshold = op.Threshold
}
func (o *DeleteThresholdRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteThresholdRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteThresholdOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteThresholdResponse structure represents the DeleteThreshold operation response
type DeleteThresholdResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteThreshold return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteThresholdResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteThresholdOperation) *xxx_DeleteThresholdOperation {
	if op == nil {
		op = &xxx_DeleteThresholdOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *DeleteThresholdResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteThresholdOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteThresholdResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteThresholdResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteThresholdOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyThresholdOperation structure represents the ModifyThreshold operation
type xxx_ModifyThresholdOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	Threshold    int32          `idl:"name:threshold" json:"threshold"`
	NewThreshold int32          `idl:"name:newThreshold" json:"new_threshold"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyThresholdOperation) OpNum() int { return 19 }

func (o *xxx_ModifyThresholdOperation) OpName() string { return "/IFsrmQuotaBase/v0/ModifyThreshold" }

func (o *xxx_ModifyThresholdOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyThresholdOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// threshold {in} (1:{alias=FSRM_QUOTA_THRESHOLD}(int32))
	{
		if err := w.WriteData(o.Threshold); err != nil {
			return err
		}
	}
	// newThreshold {in} (1:{alias=FSRM_QUOTA_THRESHOLD}(int32))
	{
		if err := w.WriteData(o.NewThreshold); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyThresholdOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// threshold {in} (1:{alias=FSRM_QUOTA_THRESHOLD}(int32))
	{
		if err := w.ReadData(&o.Threshold); err != nil {
			return err
		}
	}
	// newThreshold {in} (1:{alias=FSRM_QUOTA_THRESHOLD}(int32))
	{
		if err := w.ReadData(&o.NewThreshold); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyThresholdOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyThresholdOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ModifyThresholdOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ModifyThresholdRequest structure represents the ModifyThreshold operation request
type ModifyThresholdRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// threshold: Contains the value of the threshold to modify.
	Threshold int32 `idl:"name:threshold" json:"threshold"`
	// newThreshold: Contains the new value of the threshold.
	NewThreshold int32 `idl:"name:newThreshold" json:"new_threshold"`
}

func (o *ModifyThresholdRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifyThresholdOperation) *xxx_ModifyThresholdOperation {
	if op == nil {
		op = &xxx_ModifyThresholdOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Threshold = op.Threshold
	o.NewThreshold = op.NewThreshold
	return op
}

func (o *ModifyThresholdRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyThresholdOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Threshold = op.Threshold
	o.NewThreshold = op.NewThreshold
}
func (o *ModifyThresholdRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyThresholdRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyThresholdOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyThresholdResponse structure represents the ModifyThreshold operation response
type ModifyThresholdResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModifyThreshold return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyThresholdResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifyThresholdOperation) *xxx_ModifyThresholdOperation {
	if op == nil {
		op = &xxx_ModifyThresholdOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *ModifyThresholdResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyThresholdOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ModifyThresholdResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyThresholdResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyThresholdOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateThresholdActionOperation structure represents the CreateThresholdAction operation
type xxx_CreateThresholdActionOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Threshold  int32           `idl:"name:threshold" json:"threshold"`
	ActionType fsrm.ActionType `idl:"name:actionType" json:"action_type"`
	Action     *fsrm.Action    `idl:"name:action" json:"action"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateThresholdActionOperation) OpNum() int { return 20 }

func (o *xxx_CreateThresholdActionOperation) OpName() string {
	return "/IFsrmQuotaBase/v0/CreateThresholdAction"
}

func (o *xxx_CreateThresholdActionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateThresholdActionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// threshold {in} (1:{alias=FSRM_QUOTA_THRESHOLD}(int32))
	{
		if err := w.WriteData(o.Threshold); err != nil {
			return err
		}
	}
	// actionType {in} (1:{alias=FsrmActionType}(enum))
	{
		if err := w.WriteEnum(uint16(o.ActionType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateThresholdActionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// threshold {in} (1:{alias=FSRM_QUOTA_THRESHOLD}(int32))
	{
		if err := w.ReadData(&o.Threshold); err != nil {
			return err
		}
	}
	// actionType {in} (1:{alias=FsrmActionType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ActionType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateThresholdActionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateThresholdActionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// action {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmAction}(interface))
	{
		if o.Action != nil {
			_ptr_action := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Action != nil {
					if err := o.Action.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Action{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Action, _ptr_action); err != nil {
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

func (o *xxx_CreateThresholdActionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// action {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmAction}(interface))
	{
		_ptr_action := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Action == nil {
				o.Action = &fsrm.Action{}
			}
			if err := o.Action.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_action := func(ptr interface{}) { o.Action = *ptr.(**fsrm.Action) }
		if err := w.ReadPointer(&o.Action, _s_action, _ptr_action); err != nil {
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

// CreateThresholdActionRequest structure represents the CreateThresholdAction operation request
type CreateThresholdActionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// threshold: Contains the threshold to create the action for.
	Threshold int32 `idl:"name:threshold" json:"threshold"`
	// actionType: Contains the type of action to be created.
	ActionType fsrm.ActionType `idl:"name:actionType" json:"action_type"`
}

func (o *CreateThresholdActionRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateThresholdActionOperation) *xxx_CreateThresholdActionOperation {
	if op == nil {
		op = &xxx_CreateThresholdActionOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Threshold = op.Threshold
	o.ActionType = op.ActionType
	return op
}

func (o *CreateThresholdActionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateThresholdActionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Threshold = op.Threshold
	o.ActionType = op.ActionType
}
func (o *CreateThresholdActionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateThresholdActionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateThresholdActionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateThresholdActionResponse structure represents the CreateThresholdAction operation response
type CreateThresholdActionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// action: Pointer to an IFsrmAction interface pointer (section 3.2.4.2.4) that upon
	// completion points to the newly created action. The caller MUST release the object
	// when it is done with it.
	Action *fsrm.Action `idl:"name:action" json:"action"`
	// Return: The CreateThresholdAction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateThresholdActionResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateThresholdActionOperation) *xxx_CreateThresholdActionOperation {
	if op == nil {
		op = &xxx_CreateThresholdActionOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Action = op.Action
	o.Return = op.Return
	return op
}

func (o *CreateThresholdActionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateThresholdActionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Action = op.Action
	o.Return = op.Return
}
func (o *CreateThresholdActionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateThresholdActionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateThresholdActionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumThresholdActionsOperation structure represents the EnumThresholdActions operation
type xxx_EnumThresholdActionsOperation struct {
	This      *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Threshold int32            `idl:"name:threshold" json:"threshold"`
	Actions   *fsrm.Collection `idl:"name:actions" json:"actions"`
	Return    int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumThresholdActionsOperation) OpNum() int { return 21 }

func (o *xxx_EnumThresholdActionsOperation) OpName() string {
	return "/IFsrmQuotaBase/v0/EnumThresholdActions"
}

func (o *xxx_EnumThresholdActionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumThresholdActionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// threshold {in} (1:{alias=FSRM_QUOTA_THRESHOLD}(int32))
	{
		if err := w.WriteData(o.Threshold); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumThresholdActionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// threshold {in} (1:{alias=FSRM_QUOTA_THRESHOLD}(int32))
	{
		if err := w.ReadData(&o.Threshold); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumThresholdActionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumThresholdActionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// actions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.Actions != nil {
			_ptr_actions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Actions != nil {
					if err := o.Actions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Actions, _ptr_actions); err != nil {
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

func (o *xxx_EnumThresholdActionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// actions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_actions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Actions == nil {
				o.Actions = &fsrm.Collection{}
			}
			if err := o.Actions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_actions := func(ptr interface{}) { o.Actions = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.Actions, _s_actions, _ptr_actions); err != nil {
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

// EnumThresholdActionsRequest structure represents the EnumThresholdActions operation request
type EnumThresholdActionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// threshold: The threshold for which the associated actions will be enumerated.
	Threshold int32 `idl:"name:threshold" json:"threshold"`
}

func (o *EnumThresholdActionsRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumThresholdActionsOperation) *xxx_EnumThresholdActionsOperation {
	if op == nil {
		op = &xxx_EnumThresholdActionsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Threshold = op.Threshold
	return op
}

func (o *EnumThresholdActionsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumThresholdActionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Threshold = op.Threshold
}
func (o *EnumThresholdActionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumThresholdActionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumThresholdActionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumThresholdActionsResponse structure represents the EnumThresholdActions operation response
type EnumThresholdActionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// actions: Pointer to an IFsrmCollection interface pointer (section 3.2.4.2.1) that
	// upon completion contains IFsrmAction interface pointers (section 3.2.4.2.4) of all
	// the actions for the specified action. The caller MUST release the collection when
	// it is done with it.
	Actions *fsrm.Collection `idl:"name:actions" json:"actions"`
	// Return: The EnumThresholdActions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumThresholdActionsResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumThresholdActionsOperation) *xxx_EnumThresholdActionsOperation {
	if op == nil {
		op = &xxx_EnumThresholdActionsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Actions = op.Actions
	o.Return = op.Return
	return op
}

func (o *EnumThresholdActionsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumThresholdActionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Actions = op.Actions
	o.Return = op.Return
}
func (o *EnumThresholdActionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumThresholdActionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumThresholdActionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
