package iapphostcollectionschema

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iisa "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa"
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
	_ = iisa.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/iisa"
)

var (
	// IAppHostCollectionSchema interface identifier de095db1-5368-4d11-81f6-efef619b7bcf
	AppHostCollectionSchemaIID = &dcom.IID{Data1: 0xde095db1, Data2: 0x5368, Data3: 0x4d11, Data4: []byte{0x81, 0xf6, 0xef, 0xef, 0x61, 0x9b, 0x7b, 0xcf}}
	// Syntax UUID
	AppHostCollectionSchemaSyntaxUUID = &uuid.UUID{TimeLow: 0xde095db1, TimeMid: 0x5368, TimeHiAndVersion: 0x4d11, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0xf6, Node: [6]uint8{0xef, 0xef, 0x61, 0x9b, 0x7b, 0xcf}}
	// Syntax ID
	AppHostCollectionSchemaSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostCollectionSchemaSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostCollectionSchema interface.
type AppHostCollectionSchemaClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// AddElementNames operation.
	GetAddElementNames(context.Context, *GetAddElementNamesRequest, ...dcerpc.CallOption) (*GetAddElementNamesResponse, error)

	// GetAddElementSchema operation.
	GetAddElementSchema(context.Context, *GetAddElementSchemaRequest, ...dcerpc.CallOption) (*GetAddElementSchemaResponse, error)

	// RemoveElementSchema operation.
	GetRemoveElementSchema(context.Context, *GetRemoveElementSchemaRequest, ...dcerpc.CallOption) (*GetRemoveElementSchemaResponse, error)

	// ClearElementSchema operation.
	GetClearElementSchema(context.Context, *GetClearElementSchemaRequest, ...dcerpc.CallOption) (*GetClearElementSchemaResponse, error)

	// IsMergeAppend operation.
	GetIsMergeAppend(context.Context, *GetIsMergeAppendRequest, ...dcerpc.CallOption) (*GetIsMergeAppendResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest, ...dcerpc.CallOption) (*GetMetadataResponse, error)

	// DoesAllowDuplicates operation.
	GetDoesAllowDuplicates(context.Context, *GetDoesAllowDuplicatesRequest, ...dcerpc.CallOption) (*GetDoesAllowDuplicatesResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostCollectionSchemaClient
}

type xxx_DefaultAppHostCollectionSchemaClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostCollectionSchemaClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAppHostCollectionSchemaClient) GetAddElementNames(ctx context.Context, in *GetAddElementNamesRequest, opts ...dcerpc.CallOption) (*GetAddElementNamesResponse, error) {
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
	out := &GetAddElementNamesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostCollectionSchemaClient) GetAddElementSchema(ctx context.Context, in *GetAddElementSchemaRequest, opts ...dcerpc.CallOption) (*GetAddElementSchemaResponse, error) {
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
	out := &GetAddElementSchemaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostCollectionSchemaClient) GetRemoveElementSchema(ctx context.Context, in *GetRemoveElementSchemaRequest, opts ...dcerpc.CallOption) (*GetRemoveElementSchemaResponse, error) {
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
	out := &GetRemoveElementSchemaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostCollectionSchemaClient) GetClearElementSchema(ctx context.Context, in *GetClearElementSchemaRequest, opts ...dcerpc.CallOption) (*GetClearElementSchemaResponse, error) {
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
	out := &GetClearElementSchemaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostCollectionSchemaClient) GetIsMergeAppend(ctx context.Context, in *GetIsMergeAppendRequest, opts ...dcerpc.CallOption) (*GetIsMergeAppendResponse, error) {
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
	out := &GetIsMergeAppendResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostCollectionSchemaClient) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...dcerpc.CallOption) (*GetMetadataResponse, error) {
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
	out := &GetMetadataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostCollectionSchemaClient) GetDoesAllowDuplicates(ctx context.Context, in *GetDoesAllowDuplicatesRequest, opts ...dcerpc.CallOption) (*GetDoesAllowDuplicatesResponse, error) {
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
	out := &GetDoesAllowDuplicatesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostCollectionSchemaClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostCollectionSchemaClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAppHostCollectionSchemaClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostCollectionSchemaClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostCollectionSchemaClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAppHostCollectionSchemaClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostCollectionSchemaClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostCollectionSchemaSyntaxV0_0))...)
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
	return &xxx_DefaultAppHostCollectionSchemaClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetAddElementNamesOperation structure represents the AddElementNames operation
type xxx_GetAddElementNamesOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ElementName *oaut.String   `idl:"name:pbstrElementName" json:"element_name"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAddElementNamesOperation) OpNum() int { return 3 }

func (o *xxx_GetAddElementNamesOperation) OpName() string {
	return "/IAppHostCollectionSchema/v0/AddElementNames"
}

func (o *xxx_GetAddElementNamesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAddElementNamesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAddElementNamesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAddElementNamesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAddElementNamesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrElementName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ElementName != nil {
			_ptr_pbstrElementName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ElementName != nil {
					if err := o.ElementName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ElementName, _ptr_pbstrElementName); err != nil {
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

func (o *xxx_GetAddElementNamesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrElementName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrElementName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ElementName == nil {
				o.ElementName = &oaut.String{}
			}
			if err := o.ElementName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrElementName := func(ptr interface{}) { o.ElementName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ElementName, _s_pbstrElementName, _ptr_pbstrElementName); err != nil {
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

// GetAddElementNamesRequest structure represents the AddElementNames operation request
type GetAddElementNamesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAddElementNamesRequest) xxx_ToOp(ctx context.Context) *xxx_GetAddElementNamesOperation {
	if o == nil {
		return &xxx_GetAddElementNamesOperation{}
	}
	return &xxx_GetAddElementNamesOperation{
		This: o.This,
	}
}

func (o *GetAddElementNamesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAddElementNamesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAddElementNamesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAddElementNamesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAddElementNamesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAddElementNamesResponse structure represents the AddElementNames operation response
type GetAddElementNamesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ElementName *oaut.String   `idl:"name:pbstrElementName" json:"element_name"`
	// Return: The AddElementNames return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAddElementNamesResponse) xxx_ToOp(ctx context.Context) *xxx_GetAddElementNamesOperation {
	if o == nil {
		return &xxx_GetAddElementNamesOperation{}
	}
	return &xxx_GetAddElementNamesOperation{
		That:        o.That,
		ElementName: o.ElementName,
		Return:      o.Return,
	}
}

func (o *GetAddElementNamesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAddElementNamesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ElementName = op.ElementName
	o.Return = op.Return
}
func (o *GetAddElementNamesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAddElementNamesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAddElementNamesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAddElementSchemaOperation structure represents the GetAddElementSchema operation
type xxx_GetAddElementSchemaOperation struct {
	This        *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat             `idl:"name:That" json:"that"`
	ElementName *oaut.String               `idl:"name:bstrElementName" json:"element_name"`
	Schema      *iisa.AppHostElementSchema `idl:"name:ppSchema" json:"schema"`
	Return      int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAddElementSchemaOperation) OpNum() int { return 4 }

func (o *xxx_GetAddElementSchemaOperation) OpName() string {
	return "/IAppHostCollectionSchema/v0/GetAddElementSchema"
}

func (o *xxx_GetAddElementSchemaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAddElementSchemaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrElementName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ElementName != nil {
			_ptr_bstrElementName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ElementName != nil {
					if err := o.ElementName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ElementName, _ptr_bstrElementName); err != nil {
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

func (o *xxx_GetAddElementSchemaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrElementName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrElementName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ElementName == nil {
				o.ElementName = &oaut.String{}
			}
			if err := o.ElementName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrElementName := func(ptr interface{}) { o.ElementName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ElementName, _s_bstrElementName, _ptr_bstrElementName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAddElementSchemaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAddElementSchemaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppSchema {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElementSchema}(interface))
	{
		if o.Schema != nil {
			_ptr_ppSchema := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Schema != nil {
					if err := o.Schema.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElementSchema{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Schema, _ptr_ppSchema); err != nil {
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

func (o *xxx_GetAddElementSchemaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppSchema {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElementSchema}(interface))
	{
		_ptr_ppSchema := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Schema == nil {
				o.Schema = &iisa.AppHostElementSchema{}
			}
			if err := o.Schema.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppSchema := func(ptr interface{}) { o.Schema = *ptr.(**iisa.AppHostElementSchema) }
		if err := w.ReadPointer(&o.Schema, _s_ppSchema, _ptr_ppSchema); err != nil {
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

// GetAddElementSchemaRequest structure represents the GetAddElementSchema operation request
type GetAddElementSchemaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	ElementName *oaut.String   `idl:"name:bstrElementName" json:"element_name"`
}

func (o *GetAddElementSchemaRequest) xxx_ToOp(ctx context.Context) *xxx_GetAddElementSchemaOperation {
	if o == nil {
		return &xxx_GetAddElementSchemaOperation{}
	}
	return &xxx_GetAddElementSchemaOperation{
		This:        o.This,
		ElementName: o.ElementName,
	}
}

func (o *GetAddElementSchemaRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAddElementSchemaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ElementName = op.ElementName
}
func (o *GetAddElementSchemaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAddElementSchemaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAddElementSchemaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAddElementSchemaResponse structure represents the GetAddElementSchema operation response
type GetAddElementSchemaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Schema *iisa.AppHostElementSchema `idl:"name:ppSchema" json:"schema"`
	// Return: The GetAddElementSchema return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAddElementSchemaResponse) xxx_ToOp(ctx context.Context) *xxx_GetAddElementSchemaOperation {
	if o == nil {
		return &xxx_GetAddElementSchemaOperation{}
	}
	return &xxx_GetAddElementSchemaOperation{
		That:   o.That,
		Schema: o.Schema,
		Return: o.Return,
	}
}

func (o *GetAddElementSchemaResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAddElementSchemaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Schema = op.Schema
	o.Return = op.Return
}
func (o *GetAddElementSchemaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAddElementSchemaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAddElementSchemaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRemoveElementSchemaOperation structure represents the RemoveElementSchema operation
type xxx_GetRemoveElementSchemaOperation struct {
	This   *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Schema *iisa.AppHostElementSchema `idl:"name:ppSchema" json:"schema"`
	Return int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRemoveElementSchemaOperation) OpNum() int { return 5 }

func (o *xxx_GetRemoveElementSchemaOperation) OpName() string {
	return "/IAppHostCollectionSchema/v0/RemoveElementSchema"
}

func (o *xxx_GetRemoveElementSchemaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoveElementSchemaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRemoveElementSchemaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRemoveElementSchemaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoveElementSchemaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppSchema {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElementSchema}(interface))
	{
		if o.Schema != nil {
			_ptr_ppSchema := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Schema != nil {
					if err := o.Schema.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElementSchema{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Schema, _ptr_ppSchema); err != nil {
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

func (o *xxx_GetRemoveElementSchemaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppSchema {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElementSchema}(interface))
	{
		_ptr_ppSchema := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Schema == nil {
				o.Schema = &iisa.AppHostElementSchema{}
			}
			if err := o.Schema.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppSchema := func(ptr interface{}) { o.Schema = *ptr.(**iisa.AppHostElementSchema) }
		if err := w.ReadPointer(&o.Schema, _s_ppSchema, _ptr_ppSchema); err != nil {
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

// GetRemoveElementSchemaRequest structure represents the RemoveElementSchema operation request
type GetRemoveElementSchemaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRemoveElementSchemaRequest) xxx_ToOp(ctx context.Context) *xxx_GetRemoveElementSchemaOperation {
	if o == nil {
		return &xxx_GetRemoveElementSchemaOperation{}
	}
	return &xxx_GetRemoveElementSchemaOperation{
		This: o.This,
	}
}

func (o *GetRemoveElementSchemaRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRemoveElementSchemaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRemoveElementSchemaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetRemoveElementSchemaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRemoveElementSchemaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRemoveElementSchemaResponse structure represents the RemoveElementSchema operation response
type GetRemoveElementSchemaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Schema *iisa.AppHostElementSchema `idl:"name:ppSchema" json:"schema"`
	// Return: The RemoveElementSchema return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRemoveElementSchemaResponse) xxx_ToOp(ctx context.Context) *xxx_GetRemoveElementSchemaOperation {
	if o == nil {
		return &xxx_GetRemoveElementSchemaOperation{}
	}
	return &xxx_GetRemoveElementSchemaOperation{
		That:   o.That,
		Schema: o.Schema,
		Return: o.Return,
	}
}

func (o *GetRemoveElementSchemaResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRemoveElementSchemaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Schema = op.Schema
	o.Return = op.Return
}
func (o *GetRemoveElementSchemaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetRemoveElementSchemaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRemoveElementSchemaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClearElementSchemaOperation structure represents the ClearElementSchema operation
type xxx_GetClearElementSchemaOperation struct {
	This   *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Schema *iisa.AppHostElementSchema `idl:"name:ppSchema" json:"schema"`
	Return int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClearElementSchemaOperation) OpNum() int { return 6 }

func (o *xxx_GetClearElementSchemaOperation) OpName() string {
	return "/IAppHostCollectionSchema/v0/ClearElementSchema"
}

func (o *xxx_GetClearElementSchemaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClearElementSchemaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetClearElementSchemaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetClearElementSchemaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClearElementSchemaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppSchema {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElementSchema}(interface))
	{
		if o.Schema != nil {
			_ptr_ppSchema := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Schema != nil {
					if err := o.Schema.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElementSchema{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Schema, _ptr_ppSchema); err != nil {
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

func (o *xxx_GetClearElementSchemaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppSchema {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElementSchema}(interface))
	{
		_ptr_ppSchema := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Schema == nil {
				o.Schema = &iisa.AppHostElementSchema{}
			}
			if err := o.Schema.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppSchema := func(ptr interface{}) { o.Schema = *ptr.(**iisa.AppHostElementSchema) }
		if err := w.ReadPointer(&o.Schema, _s_ppSchema, _ptr_ppSchema); err != nil {
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

// GetClearElementSchemaRequest structure represents the ClearElementSchema operation request
type GetClearElementSchemaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClearElementSchemaRequest) xxx_ToOp(ctx context.Context) *xxx_GetClearElementSchemaOperation {
	if o == nil {
		return &xxx_GetClearElementSchemaOperation{}
	}
	return &xxx_GetClearElementSchemaOperation{
		This: o.This,
	}
}

func (o *GetClearElementSchemaRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClearElementSchemaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClearElementSchemaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClearElementSchemaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClearElementSchemaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClearElementSchemaResponse structure represents the ClearElementSchema operation response
type GetClearElementSchemaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Schema *iisa.AppHostElementSchema `idl:"name:ppSchema" json:"schema"`
	// Return: The ClearElementSchema return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClearElementSchemaResponse) xxx_ToOp(ctx context.Context) *xxx_GetClearElementSchemaOperation {
	if o == nil {
		return &xxx_GetClearElementSchemaOperation{}
	}
	return &xxx_GetClearElementSchemaOperation{
		That:   o.That,
		Schema: o.Schema,
		Return: o.Return,
	}
}

func (o *GetClearElementSchemaResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClearElementSchemaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Schema = op.Schema
	o.Return = op.Return
}
func (o *GetClearElementSchemaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClearElementSchemaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClearElementSchemaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsMergeAppendOperation structure represents the IsMergeAppend operation
type xxx_GetIsMergeAppendOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsMergeAppend int16          `idl:"name:pfIsMergeAppend" json:"is_merge_append"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsMergeAppendOperation) OpNum() int { return 7 }

func (o *xxx_GetIsMergeAppendOperation) OpName() string {
	return "/IAppHostCollectionSchema/v0/IsMergeAppend"
}

func (o *xxx_GetIsMergeAppendOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsMergeAppendOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsMergeAppendOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsMergeAppendOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsMergeAppendOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfIsMergeAppend {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.IsMergeAppend); err != nil {
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

func (o *xxx_GetIsMergeAppendOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfIsMergeAppend {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.IsMergeAppend); err != nil {
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

// GetIsMergeAppendRequest structure represents the IsMergeAppend operation request
type GetIsMergeAppendRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsMergeAppendRequest) xxx_ToOp(ctx context.Context) *xxx_GetIsMergeAppendOperation {
	if o == nil {
		return &xxx_GetIsMergeAppendOperation{}
	}
	return &xxx_GetIsMergeAppendOperation{
		This: o.This,
	}
}

func (o *GetIsMergeAppendRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsMergeAppendOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsMergeAppendRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetIsMergeAppendRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsMergeAppendOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsMergeAppendResponse structure represents the IsMergeAppend operation response
type GetIsMergeAppendResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsMergeAppend int16          `idl:"name:pfIsMergeAppend" json:"is_merge_append"`
	// Return: The IsMergeAppend return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsMergeAppendResponse) xxx_ToOp(ctx context.Context) *xxx_GetIsMergeAppendOperation {
	if o == nil {
		return &xxx_GetIsMergeAppendOperation{}
	}
	return &xxx_GetIsMergeAppendOperation{
		That:          o.That,
		IsMergeAppend: o.IsMergeAppend,
		Return:        o.Return,
	}
}

func (o *GetIsMergeAppendResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsMergeAppendOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsMergeAppend = op.IsMergeAppend
	o.Return = op.Return
}
func (o *GetIsMergeAppendResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetIsMergeAppendResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsMergeAppendOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMetadataOperation structure represents the GetMetadata operation
type xxx_GetMetadataOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	MetadataType *oaut.String   `idl:"name:bstrMetadataType" json:"metadata_type"`
	Value        *oaut.Variant  `idl:"name:pValue" json:"value"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMetadataOperation) OpNum() int { return 8 }

func (o *xxx_GetMetadataOperation) OpName() string { return "/IAppHostCollectionSchema/v0/GetMetadata" }

func (o *xxx_GetMetadataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMetadataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrMetadataType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MetadataType != nil {
			_ptr_bstrMetadataType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MetadataType != nil {
					if err := o.MetadataType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MetadataType, _ptr_bstrMetadataType); err != nil {
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

func (o *xxx_GetMetadataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrMetadataType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMetadataType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MetadataType == nil {
				o.MetadataType = &oaut.String{}
			}
			if err := o.MetadataType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMetadataType := func(ptr interface{}) { o.MetadataType = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MetadataType, _s_bstrMetadataType, _ptr_bstrMetadataType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMetadataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMetadataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Value != nil {
			_ptr_pValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Value, _ptr_pValue); err != nil {
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

func (o *xxx_GetMetadataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.Variant{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pValue := func(ptr interface{}) { o.Value = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Value, _s_pValue, _ptr_pValue); err != nil {
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

// GetMetadataRequest structure represents the GetMetadata operation request
type GetMetadataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	MetadataType *oaut.String   `idl:"name:bstrMetadataType" json:"metadata_type"`
}

func (o *GetMetadataRequest) xxx_ToOp(ctx context.Context) *xxx_GetMetadataOperation {
	if o == nil {
		return &xxx_GetMetadataOperation{}
	}
	return &xxx_GetMetadataOperation{
		This:         o.This,
		MetadataType: o.MetadataType,
	}
}

func (o *GetMetadataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMetadataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MetadataType = op.MetadataType
}
func (o *GetMetadataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMetadataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMetadataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMetadataResponse structure represents the GetMetadata operation response
type GetMetadataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value *oaut.Variant  `idl:"name:pValue" json:"value"`
	// Return: The GetMetadata return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMetadataResponse) xxx_ToOp(ctx context.Context) *xxx_GetMetadataOperation {
	if o == nil {
		return &xxx_GetMetadataOperation{}
	}
	return &xxx_GetMetadataOperation{
		That:   o.That,
		Value:  o.Value,
		Return: o.Return,
	}
}

func (o *GetMetadataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMetadataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Value = op.Value
	o.Return = op.Return
}
func (o *GetMetadataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMetadataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMetadataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDoesAllowDuplicatesOperation structure represents the DoesAllowDuplicates operation
type xxx_GetDoesAllowDuplicatesOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	AllowDuplicates int16          `idl:"name:pfAllowDuplicates" json:"allow_duplicates"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDoesAllowDuplicatesOperation) OpNum() int { return 9 }

func (o *xxx_GetDoesAllowDuplicatesOperation) OpName() string {
	return "/IAppHostCollectionSchema/v0/DoesAllowDuplicates"
}

func (o *xxx_GetDoesAllowDuplicatesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDoesAllowDuplicatesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDoesAllowDuplicatesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDoesAllowDuplicatesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDoesAllowDuplicatesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfAllowDuplicates {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.AllowDuplicates); err != nil {
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

func (o *xxx_GetDoesAllowDuplicatesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfAllowDuplicates {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.AllowDuplicates); err != nil {
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

// GetDoesAllowDuplicatesRequest structure represents the DoesAllowDuplicates operation request
type GetDoesAllowDuplicatesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDoesAllowDuplicatesRequest) xxx_ToOp(ctx context.Context) *xxx_GetDoesAllowDuplicatesOperation {
	if o == nil {
		return &xxx_GetDoesAllowDuplicatesOperation{}
	}
	return &xxx_GetDoesAllowDuplicatesOperation{
		This: o.This,
	}
}

func (o *GetDoesAllowDuplicatesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDoesAllowDuplicatesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDoesAllowDuplicatesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDoesAllowDuplicatesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDoesAllowDuplicatesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDoesAllowDuplicatesResponse structure represents the DoesAllowDuplicates operation response
type GetDoesAllowDuplicatesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	AllowDuplicates int16          `idl:"name:pfAllowDuplicates" json:"allow_duplicates"`
	// Return: The DoesAllowDuplicates return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDoesAllowDuplicatesResponse) xxx_ToOp(ctx context.Context) *xxx_GetDoesAllowDuplicatesOperation {
	if o == nil {
		return &xxx_GetDoesAllowDuplicatesOperation{}
	}
	return &xxx_GetDoesAllowDuplicatesOperation{
		That:            o.That,
		AllowDuplicates: o.AllowDuplicates,
		Return:          o.Return,
	}
}

func (o *GetDoesAllowDuplicatesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDoesAllowDuplicatesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AllowDuplicates = op.AllowDuplicates
	o.Return = op.Return
}
func (o *GetDoesAllowDuplicatesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDoesAllowDuplicatesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDoesAllowDuplicatesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
