package iclusterupdate

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
	GoPackage = "dcom/csvp"
)

var (
	// IClusterUpdate interface identifier e3c9b851-c442-432b-8fc6-a7faafc09d3b
	ClusterUpdateIID = &dcom.IID{Data1: 0xe3c9b851, Data2: 0xc442, Data3: 0x432b, Data4: []byte{0x8f, 0xc6, 0xa7, 0xfa, 0xaf, 0xc0, 0x9d, 0x3b}}
	// Syntax UUID
	ClusterUpdateSyntaxUUID = &uuid.UUID{TimeLow: 0xe3c9b851, TimeMid: 0xc442, TimeHiAndVersion: 0x432b, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xc6, Node: [6]uint8{0xa7, 0xfa, 0xaf, 0xc0, 0x9d, 0x3b}}
	// Syntax ID
	ClusterUpdateSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClusterUpdateSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IClusterUpdate interface.
type ClusterUpdateClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetUpdates method queries the local server for all of the updates that are installed
	// on the local server.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 3.
	GetUpdates(context.Context, *GetUpdatesRequest, ...dcerpc.CallOption) (*GetUpdatesResponse, error)

	// The Count method returns the number of updates that are installed on the local server.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 4.
	Count(context.Context, *CountRequest, ...dcerpc.CallOption) (*CountResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClusterUpdateClient
}

type xxx_DefaultClusterUpdateClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClusterUpdateClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultClusterUpdateClient) GetUpdates(ctx context.Context, in *GetUpdatesRequest, opts ...dcerpc.CallOption) (*GetUpdatesResponse, error) {
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
	out := &GetUpdatesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterUpdateClient) Count(ctx context.Context, in *CountRequest, opts ...dcerpc.CallOption) (*CountResponse, error) {
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
	out := &CountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterUpdateClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClusterUpdateClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClusterUpdateClient) IPID(ctx context.Context, ipid *dcom.IPID) ClusterUpdateClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClusterUpdateClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewClusterUpdateClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClusterUpdateClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClusterUpdateSyntaxV0_0))...)
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
	return &xxx_DefaultClusterUpdateClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetUpdatesOperation structure represents the GetUpdates operation
type xxx_GetUpdatesOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	UpdateCount uint32         `idl:"name:UpdateCount" json:"update_count"`
	Updates     *oaut.String   `idl:"name:updates" json:"updates"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUpdatesOperation) OpNum() int { return 3 }

func (o *xxx_GetUpdatesOperation) OpName() string { return "/IClusterUpdate/v0/GetUpdates" }

func (o *xxx_GetUpdatesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUpdatesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetUpdatesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetUpdatesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUpdatesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// UpdateCount {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.UpdateCount); err != nil {
			return err
		}
	}
	// updates {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Updates != nil {
			_ptr_updates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Updates != nil {
					if err := o.Updates.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Updates, _ptr_updates); err != nil {
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

func (o *xxx_GetUpdatesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// UpdateCount {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.UpdateCount); err != nil {
			return err
		}
	}
	// updates {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_updates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Updates == nil {
				o.Updates = &oaut.String{}
			}
			if err := o.Updates.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_updates := func(ptr interface{}) { o.Updates = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Updates, _s_updates, _ptr_updates); err != nil {
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

// GetUpdatesRequest structure represents the GetUpdates operation request
type GetUpdatesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetUpdatesRequest) xxx_ToOp(ctx context.Context) *xxx_GetUpdatesOperation {
	if o == nil {
		return &xxx_GetUpdatesOperation{}
	}
	return &xxx_GetUpdatesOperation{
		This: o.This,
	}
}

func (o *GetUpdatesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUpdatesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetUpdatesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetUpdatesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUpdatesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUpdatesResponse structure represents the GetUpdates operation response
type GetUpdatesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// UpdateCount: Upon successful completion of the method, the server MUST set this parameter
	// to the number of updates in the ClusterUpdates collection.
	UpdateCount uint32 `idl:"name:UpdateCount" json:"update_count"`
	// updates: Upon successful completion of this method, the server MUST set this parameter
	// to a null-terminated Unicode string containing XML representing the contents of the
	// ClusterUpdates collection. The XML is formatted as follows:
	//
	// The XML string starts with an XML tag with the name "updates" that has an attribute
	// with the name "version" with a value set to 1.0.
	//
	// <updates version="1.0">
	//
	// * "id", with text containing the value of *ClusterUpdate.Id*.
	//
	// * "title", with text containing the value of *ClusterUpdate.Title*.
	//
	// * "description", with text containing the value of *ClusterUpdate.Description*.
	//
	// * "supportUrl", with text containing the value of *ClusterUpdate.SupportUrl*.
	//
	// * "knowledgebase", with a child "articleId" XML element for each entry in the *ClusterUpdate.ArticleIds*
	// collection. If the *ClusterUpdate.ArticleIds* collection is empty, then the "knowledgebase"
	// element MUST be an empty XML element. Otherwise, the child XML elements are as follows:
	//
	// * "articleId", with text containing the value of one entry from the *ClusterUpdate.ArticleIds*
	// collection.
	//
	// * "securityBulletin", with a child XML element for each entry in the *ClusterUpdate.SecurityBulletins*
	// collection. If the *ClusterUpdate.SecurityBulletins* collection is empty, then the
	// "securityBulletin" element MUST be an empty XML element. Otherwise, the child XML
	// elements are as follows:
	//
	// * "bulletinId", with text containing the value of one entry from the *ClusterUpdate.SecurityBulletins*
	// collection.
	//
	// * "superseded", with a child XML element for each entry in the *ClusterUpdate.UpdateIds*
	// collection. If the *ClusterUpdate.UpdateIds* collection is empty, then the "superseded"
	// element MUST be an empty XML element. Otherwise, the child XML elements are as follows:
	//
	// * "updateId", with text containing the value of one entry from the *ClusterUpdate.UpdateIds*
	// collection.
	//
	// The XML string concludes with an XML close tag with the name "updates".
	Updates *oaut.String `idl:"name:updates" json:"updates"`
	// Return: The GetUpdates return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUpdatesResponse) xxx_ToOp(ctx context.Context) *xxx_GetUpdatesOperation {
	if o == nil {
		return &xxx_GetUpdatesOperation{}
	}
	return &xxx_GetUpdatesOperation{
		That:        o.That,
		UpdateCount: o.UpdateCount,
		Updates:     o.Updates,
		Return:      o.Return,
	}
}

func (o *GetUpdatesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUpdatesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.UpdateCount = op.UpdateCount
	o.Updates = op.Updates
	o.Return = op.Return
}
func (o *GetUpdatesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetUpdatesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUpdatesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CountOperation structure represents the Count operation
type xxx_CountOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Count  int32          `idl:"name:Count" json:"count"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CountOperation) OpNum() int { return 4 }

func (o *xxx_CountOperation) OpName() string { return "/IClusterUpdate/v0/Count" }

func (o *xxx_CountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Count {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Count); err != nil {
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

func (o *xxx_CountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Count {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Count); err != nil {
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

// CountRequest structure represents the Count operation request
type CountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CountRequest) xxx_ToOp(ctx context.Context) *xxx_CountOperation {
	if o == nil {
		return &xxx_CountOperation{}
	}
	return &xxx_CountOperation{
		This: o.This,
	}
}

func (o *CountRequest) xxx_FromOp(ctx context.Context, op *xxx_CountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CountResponse structure represents the Count operation response
type CountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Count: A value indicating the number of updates installed on the local server.
	Count int32 `idl:"name:Count" json:"count"`
	// Return: The Count return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CountResponse) xxx_ToOp(ctx context.Context) *xxx_CountOperation {
	if o == nil {
		return &xxx_CountOperation{}
	}
	return &xxx_CountOperation{
		That:   o.That,
		Count:  o.Count,
		Return: o.Return,
	}
}

func (o *CountResponse) xxx_FromOp(ctx context.Context, op *xxx_CountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Count = op.Count
	o.Return = op.Return
}
func (o *CountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
