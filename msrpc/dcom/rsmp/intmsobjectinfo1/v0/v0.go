package intmsobjectinfo1

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
	rsmp "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
	_ = rsmp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// INtmsObjectInfo1 interface identifier 69ab7050-3059-11d1-8faf-00a024cb6019
	ObjectInfo1IID = &dcom.IID{Data1: 0x69ab7050, Data2: 0x3059, Data3: 0x11d1, Data4: []byte{0x8f, 0xaf, 0x00, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	// Syntax UUID
	ObjectInfo1SyntaxUUID = &uuid.UUID{TimeLow: 0x69ab7050, TimeMid: 0x3059, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xaf, Node: [6]uint8{0x0, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	// Syntax ID
	ObjectInfo1SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ObjectInfo1SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// INtmsObjectInfo1 interface.
type ObjectInfo1Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetNtmsServerObjectInformationA method retrieves information about an object,
	// as a sequence of ASCII characters.
	GetNTMSServerObjectInformationA(context.Context, *GetNTMSServerObjectInformationARequest, ...dcerpc.CallOption) (*GetNTMSServerObjectInformationAResponse, error)

	// The GetNtmsServerObjectInformationW method retrieves information about an object,
	// as a sequence of Unicode characters.
	GetNTMSServerObjectInformationW(context.Context, *GetNTMSServerObjectInformationWRequest, ...dcerpc.CallOption) (*GetNTMSServerObjectInformationWResponse, error)

	// The SetNtmsObjectInformationA method changes the information of an object, with strings
	// encoded using ASCII.
	SetNTMSObjectInformationA(context.Context, *SetNTMSObjectInformationARequest, ...dcerpc.CallOption) (*SetNTMSObjectInformationAResponse, error)

	// The SetNtmsObjectInformationW method changes the information of an object, with strings
	// encoded using Unicode.
	SetNTMSObjectInformationW(context.Context, *SetNTMSObjectInformationWRequest, ...dcerpc.CallOption) (*SetNTMSObjectInformationWResponse, error)

	// The CreateNtmsMediaA method creates a new offline medium for a media pool, with strings
	// encoded using ASCII.
	CreateNTMSMediaA(context.Context, *CreateNTMSMediaARequest, ...dcerpc.CallOption) (*CreateNTMSMediaAResponse, error)

	// The CreateNtmsMediaW method creates a new offline medium for a media pool, with strings
	// encoded using Unicode.
	CreateNTMSMediaW(context.Context, *CreateNTMSMediaWRequest, ...dcerpc.CallOption) (*CreateNTMSMediaWResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ObjectInfo1Client
}

type xxx_DefaultObjectInfo1Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultObjectInfo1Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultObjectInfo1Client) GetNTMSServerObjectInformationA(ctx context.Context, in *GetNTMSServerObjectInformationARequest, opts ...dcerpc.CallOption) (*GetNTMSServerObjectInformationAResponse, error) {
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
	out := &GetNTMSServerObjectInformationAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectInfo1Client) GetNTMSServerObjectInformationW(ctx context.Context, in *GetNTMSServerObjectInformationWRequest, opts ...dcerpc.CallOption) (*GetNTMSServerObjectInformationWResponse, error) {
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
	out := &GetNTMSServerObjectInformationWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectInfo1Client) SetNTMSObjectInformationA(ctx context.Context, in *SetNTMSObjectInformationARequest, opts ...dcerpc.CallOption) (*SetNTMSObjectInformationAResponse, error) {
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
	out := &SetNTMSObjectInformationAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectInfo1Client) SetNTMSObjectInformationW(ctx context.Context, in *SetNTMSObjectInformationWRequest, opts ...dcerpc.CallOption) (*SetNTMSObjectInformationWResponse, error) {
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
	out := &SetNTMSObjectInformationWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectInfo1Client) CreateNTMSMediaA(ctx context.Context, in *CreateNTMSMediaARequest, opts ...dcerpc.CallOption) (*CreateNTMSMediaAResponse, error) {
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
	out := &CreateNTMSMediaAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectInfo1Client) CreateNTMSMediaW(ctx context.Context, in *CreateNTMSMediaWRequest, opts ...dcerpc.CallOption) (*CreateNTMSMediaWResponse, error) {
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
	out := &CreateNTMSMediaWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectInfo1Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultObjectInfo1Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultObjectInfo1Client) IPID(ctx context.Context, ipid *dcom.IPID) ObjectInfo1Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultObjectInfo1Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewObjectInfo1Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ObjectInfo1Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ObjectInfo1SyntaxV0_0))...)
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
	return &xxx_DefaultObjectInfo1Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetNTMSServerObjectInformationAOperation structure represents the GetNtmsServerObjectInformationA operation
type xxx_GetNTMSServerObjectInformationAOperation struct {
	This     *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat           `idl:"name:That" json:"that"`
	ObjectID *dtyp.GUID               `idl:"name:lpObjectId;pointer:unique" json:"object_id"`
	Info     *rsmp.ObjectInformationA `idl:"name:lpInfo" json:"info"`
	Type     uint32                   `idl:"name:dwType" json:"type"`
	Size     uint32                   `idl:"name:dwSize" json:"size"`
	Return   int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSServerObjectInformationAOperation) OpNum() int { return 0 }

func (o *xxx_GetNTMSServerObjectInformationAOperation) OpName() string {
	return "/INtmsObjectInfo1/v0/GetNtmsServerObjectInformationA"
}

func (o *xxx_GetNTMSServerObjectInformationAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSServerObjectInformationAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			_ptr_lpObjectId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectID != nil {
					if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectID, _ptr_lpObjectId); err != nil {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// dwSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSServerObjectInformationAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpObjectId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectID == nil {
				o.ObjectID = &dtyp.GUID{}
			}
			if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpObjectId := func(ptr interface{}) { o.ObjectID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.ObjectID, _s_lpObjectId, _ptr_lpObjectId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// dwSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSServerObjectInformationAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSServerObjectInformationAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpInfo {out} (1:{alias=LPNTMS_OBJECTINFORMATIONA}*(1))(2:{alias=NTMS_OBJECTINFORMATIONA}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rsmp.ObjectInformationA{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_GetNTMSServerObjectInformationAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpInfo {out} (1:{alias=LPNTMS_OBJECTINFORMATIONA,pointer=ref}*(1))(2:{alias=NTMS_OBJECTINFORMATIONA}(struct))
	{
		if o.Info == nil {
			o.Info = &rsmp.ObjectInformationA{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// GetNTMSServerObjectInformationARequest structure represents the GetNtmsServerObjectInformationA operation request
type GetNTMSServerObjectInformationARequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpObjectId: A pointer to the identifier of the object for which to retrieve information.
	ObjectID *dtyp.GUID `idl:"name:lpObjectId;pointer:unique" json:"object_id"`
	// dwType: A value from the NtmsObjectsTypes enumeration defining the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// dwSize: The size, in bytes, of the appropriate structure for lpInfo.
	//
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	|               RETURN               |                                                                       |
	//	|             VALUE/CODE             |                              DESCRIPTION                              |
	//	|                                    |                                                                       |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                              |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | Access denied.                                                        |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY | An allocation failure occurred during processing.                     |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | A parameter is missing, or the dwType or dwSize parameter is invalid. |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | The lpObjectId parameter is invalid.                                  |
	//	+------------------------------------+-----------------------------------------------------------------------+
	Size uint32 `idl:"name:dwSize" json:"size"`
}

func (o *GetNTMSServerObjectInformationARequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSServerObjectInformationAOperation) *xxx_GetNTMSServerObjectInformationAOperation {
	if op == nil {
		op = &xxx_GetNTMSServerObjectInformationAOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.Size = o.Size
	return op
}

func (o *GetNTMSServerObjectInformationARequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSServerObjectInformationAOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.Size = op.Size
}
func (o *GetNTMSServerObjectInformationARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSServerObjectInformationARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSServerObjectInformationAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSServerObjectInformationAResponse structure represents the GetNtmsServerObjectInformationA operation response
type GetNTMSServerObjectInformationAResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpInfo: A pointer to an NTMS_OBJECTINFORMATIONA structure describing the properties
	// of the object.
	Info *rsmp.ObjectInformationA `idl:"name:lpInfo" json:"info"`
	// Return: The GetNtmsServerObjectInformationA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSServerObjectInformationAResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSServerObjectInformationAOperation) *xxx_GetNTMSServerObjectInformationAOperation {
	if op == nil {
		op = &xxx_GetNTMSServerObjectInformationAOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *GetNTMSServerObjectInformationAResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSServerObjectInformationAOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Info = op.Info
	o.Return = op.Return
}
func (o *GetNTMSServerObjectInformationAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSServerObjectInformationAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSServerObjectInformationAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNTMSServerObjectInformationWOperation structure represents the GetNtmsServerObjectInformationW operation
type xxx_GetNTMSServerObjectInformationWOperation struct {
	This     *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat           `idl:"name:That" json:"that"`
	ObjectID *dtyp.GUID               `idl:"name:lpObjectId" json:"object_id"`
	Info     *rsmp.ObjectInformationW `idl:"name:lpInfo" json:"info"`
	Type     uint32                   `idl:"name:dwType" json:"type"`
	Size     uint32                   `idl:"name:dwSize" json:"size"`
	Return   int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSServerObjectInformationWOperation) OpNum() int { return 1 }

func (o *xxx_GetNTMSServerObjectInformationWOperation) OpName() string {
	return "/INtmsObjectInfo1/v0/GetNtmsServerObjectInformationW"
}

func (o *xxx_GetNTMSServerObjectInformationWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSServerObjectInformationWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// dwSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSServerObjectInformationWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// dwSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSServerObjectInformationWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSServerObjectInformationWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpInfo {out} (1:{alias=LPNTMS_OBJECTINFORMATIONW}*(1))(2:{alias=NTMS_OBJECTINFORMATIONW}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rsmp.ObjectInformationW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_GetNTMSServerObjectInformationWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpInfo {out} (1:{alias=LPNTMS_OBJECTINFORMATIONW,pointer=ref}*(1))(2:{alias=NTMS_OBJECTINFORMATIONW}(struct))
	{
		if o.Info == nil {
			o.Info = &rsmp.ObjectInformationW{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// GetNTMSServerObjectInformationWRequest structure represents the GetNtmsServerObjectInformationW operation request
type GetNTMSServerObjectInformationWRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpObjectId: A pointer to the identifier of the object for which to retrieve information.
	ObjectID *dtyp.GUID `idl:"name:lpObjectId" json:"object_id"`
	// dwType: A value from the NtmsObjectsTypes enumeration defining the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// dwSize: The size, in bytes, of the appropriate structure for lpInfo.
	//
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	|               RETURN               |                                                                       |
	//	|             VALUE/CODE             |                              DESCRIPTION                              |
	//	|                                    |                                                                       |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                              |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | Access denied.                                                        |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY | An allocation failure occurred during processing.                     |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | A parameter is missing, or the dwType or dwSize parameter is invalid. |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | The lpObjectId parameter is invalid.                                  |
	//	+------------------------------------+-----------------------------------------------------------------------+
	Size uint32 `idl:"name:dwSize" json:"size"`
}

func (o *GetNTMSServerObjectInformationWRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSServerObjectInformationWOperation) *xxx_GetNTMSServerObjectInformationWOperation {
	if op == nil {
		op = &xxx_GetNTMSServerObjectInformationWOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.Size = o.Size
	return op
}

func (o *GetNTMSServerObjectInformationWRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSServerObjectInformationWOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.Size = op.Size
}
func (o *GetNTMSServerObjectInformationWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSServerObjectInformationWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSServerObjectInformationWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSServerObjectInformationWResponse structure represents the GetNtmsServerObjectInformationW operation response
type GetNTMSServerObjectInformationWResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpInfo: A pointer to an NTMS_OBJECTINFORMATIONW structure describing the properties
	// of the object.
	Info *rsmp.ObjectInformationW `idl:"name:lpInfo" json:"info"`
	// Return: The GetNtmsServerObjectInformationW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSServerObjectInformationWResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSServerObjectInformationWOperation) *xxx_GetNTMSServerObjectInformationWOperation {
	if op == nil {
		op = &xxx_GetNTMSServerObjectInformationWOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *GetNTMSServerObjectInformationWResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSServerObjectInformationWOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Info = op.Info
	o.Return = op.Return
}
func (o *GetNTMSServerObjectInformationWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSServerObjectInformationWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSServerObjectInformationWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNTMSObjectInformationAOperation structure represents the SetNtmsObjectInformationA operation
type xxx_SetNTMSObjectInformationAOperation struct {
	This     *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat           `idl:"name:That" json:"that"`
	ObjectID *dtyp.GUID               `idl:"name:lpObjectId" json:"object_id"`
	Info     *rsmp.ObjectInformationA `idl:"name:lpInfo" json:"info"`
	Return   int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNTMSObjectInformationAOperation) OpNum() int { return 2 }

func (o *xxx_SetNTMSObjectInformationAOperation) OpName() string {
	return "/INtmsObjectInfo1/v0/SetNtmsObjectInformationA"
}

func (o *xxx_SetNTMSObjectInformationAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectInformationAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpInfo {in} (1:{alias=LPNTMS_OBJECTINFORMATIONA}*(1))(2:{alias=NTMS_OBJECTINFORMATIONA}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rsmp.ObjectInformationA{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectInformationAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpInfo {in} (1:{alias=LPNTMS_OBJECTINFORMATIONA,pointer=ref}*(1))(2:{alias=NTMS_OBJECTINFORMATIONA}(struct))
	{
		if o.Info == nil {
			o.Info = &rsmp.ObjectInformationA{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectInformationAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectInformationAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNTMSObjectInformationAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNTMSObjectInformationARequest structure represents the SetNtmsObjectInformationA operation request
type SetNTMSObjectInformationARequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpObjectId: A pointer to the identifier of the object to change.
	ObjectID *dtyp.GUID `idl:"name:lpObjectId" json:"object_id"`
	// lpInfo:  A pointer to an NTMS_OBJECTINFORMATIONA structure describing the properties
	// of the object to change.
	//
	//	+------------------------------------+---------------------------------------------------+
	//	|               RETURN               |                                                   |
	//	|             VALUE/CODE             |                    DESCRIPTION                    |
	//	|                                    |                                                   |
	//	+------------------------------------+---------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                          |
	//	+------------------------------------+---------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | Access to an object was denied.                   |
	//	+------------------------------------+---------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY | An allocation failure occurred during processing. |
	//	+------------------------------------+---------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | A parameter is not valid.                         |
	//	+------------------------------------+---------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database query or update failed.              |
	//	+------------------------------------+---------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL     | The database is full.                             |
	//	+------------------------------------+---------------------------------------------------+
	Info *rsmp.ObjectInformationA `idl:"name:lpInfo" json:"info"`
}

func (o *SetNTMSObjectInformationARequest) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSObjectInformationAOperation) *xxx_SetNTMSObjectInformationAOperation {
	if op == nil {
		op = &xxx_SetNTMSObjectInformationAOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Info = o.Info
	return op
}

func (o *SetNTMSObjectInformationARequest) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSObjectInformationAOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Info = op.Info
}
func (o *SetNTMSObjectInformationARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNTMSObjectInformationARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSObjectInformationAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNTMSObjectInformationAResponse structure represents the SetNtmsObjectInformationA operation response
type SetNTMSObjectInformationAResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetNtmsObjectInformationA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNTMSObjectInformationAResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSObjectInformationAOperation) *xxx_SetNTMSObjectInformationAOperation {
	if op == nil {
		op = &xxx_SetNTMSObjectInformationAOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNTMSObjectInformationAResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSObjectInformationAOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNTMSObjectInformationAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNTMSObjectInformationAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSObjectInformationAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNTMSObjectInformationWOperation structure represents the SetNtmsObjectInformationW operation
type xxx_SetNTMSObjectInformationWOperation struct {
	This     *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat           `idl:"name:That" json:"that"`
	ObjectID *dtyp.GUID               `idl:"name:lpObjectId" json:"object_id"`
	Info     *rsmp.ObjectInformationW `idl:"name:lpInfo" json:"info"`
	Return   int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNTMSObjectInformationWOperation) OpNum() int { return 3 }

func (o *xxx_SetNTMSObjectInformationWOperation) OpName() string {
	return "/INtmsObjectInfo1/v0/SetNtmsObjectInformationW"
}

func (o *xxx_SetNTMSObjectInformationWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectInformationWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpInfo {in} (1:{alias=LPNTMS_OBJECTINFORMATIONW}*(1))(2:{alias=NTMS_OBJECTINFORMATIONW}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rsmp.ObjectInformationW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectInformationWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpInfo {in} (1:{alias=LPNTMS_OBJECTINFORMATIONW,pointer=ref}*(1))(2:{alias=NTMS_OBJECTINFORMATIONW}(struct))
	{
		if o.Info == nil {
			o.Info = &rsmp.ObjectInformationW{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectInformationWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectInformationWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNTMSObjectInformationWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNTMSObjectInformationWRequest structure represents the SetNtmsObjectInformationW operation request
type SetNTMSObjectInformationWRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpObjectId: A pointer to the identifier of the object to change.
	ObjectID *dtyp.GUID `idl:"name:lpObjectId" json:"object_id"`
	// lpInfo:  A pointer to an NTMS_OBJECTINFORMATIONW (section 2.2.4.21) structure describing
	// the properties of the object to change.
	//
	//	+------------------------------------+---------------------------------------------------+
	//	|               RETURN               |                                                   |
	//	|             VALUE/CODE             |                    DESCRIPTION                    |
	//	|                                    |                                                   |
	//	+------------------------------------+---------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                          |
	//	+------------------------------------+---------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | Access to an object is denied.                    |
	//	+------------------------------------+---------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY | An allocation failure occurred during processing. |
	//	+------------------------------------+---------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | A parameter is not valid.                         |
	//	+------------------------------------+---------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database query or update failed.              |
	//	+------------------------------------+---------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL     | The database is full.                             |
	//	+------------------------------------+---------------------------------------------------+
	Info *rsmp.ObjectInformationW `idl:"name:lpInfo" json:"info"`
}

func (o *SetNTMSObjectInformationWRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSObjectInformationWOperation) *xxx_SetNTMSObjectInformationWOperation {
	if op == nil {
		op = &xxx_SetNTMSObjectInformationWOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Info = o.Info
	return op
}

func (o *SetNTMSObjectInformationWRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSObjectInformationWOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Info = op.Info
}
func (o *SetNTMSObjectInformationWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNTMSObjectInformationWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSObjectInformationWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNTMSObjectInformationWResponse structure represents the SetNtmsObjectInformationW operation response
type SetNTMSObjectInformationWResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetNtmsObjectInformationW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNTMSObjectInformationWResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSObjectInformationWOperation) *xxx_SetNTMSObjectInformationWOperation {
	if op == nil {
		op = &xxx_SetNTMSObjectInformationWOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNTMSObjectInformationWResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSObjectInformationWOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNTMSObjectInformationWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNTMSObjectInformationWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSObjectInformationWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateNTMSMediaAOperation structure represents the CreateNtmsMediaA operation
type xxx_CreateNTMSMediaAOperation struct {
	This           *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Media          *rsmp.ObjectInformationA   `idl:"name:lpMedia" json:"media"`
	List           []*rsmp.ObjectInformationA `idl:"name:lpList;size_is:(lpdwListBufferSize);length_is:(dwListCount)" json:"list"`
	ListBufferSize uint32                     `idl:"name:lpdwListBufferSize" json:"list_buffer_size"`
	ListCount      uint32                     `idl:"name:dwListCount" json:"list_count"`
	Options        uint32                     `idl:"name:dwOptions" json:"options"`
	Return         int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateNTMSMediaAOperation) OpNum() int { return 4 }

func (o *xxx_CreateNTMSMediaAOperation) OpName() string {
	return "/INtmsObjectInfo1/v0/CreateNtmsMediaA"
}

func (o *xxx_CreateNTMSMediaAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.List != nil && o.ListBufferSize == 0 {
		o.ListBufferSize = uint32(len(o.List))
	}
	if o.List != nil && o.ListCount == 0 {
		o.ListCount = uint32(len(o.List))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMedia {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONA}*(1))(2:{alias=NTMS_OBJECTINFORMATIONA}(struct))
	{
		if o.Media != nil {
			if err := o.Media.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rsmp.ObjectInformationA{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpList {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONA}*(1))(2:{alias=NTMS_OBJECTINFORMATIONA}[dim:0,size_is=lpdwListBufferSize,length_is=dwListCount](struct))
	{
		dimSize1 := uint64(o.ListBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ListCount)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.List {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.List[i1] != nil {
				if err := o.List[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&rsmp.ObjectInformationA{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.List); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&rsmp.ObjectInformationA{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwListBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ListBufferSize); err != nil {
			return err
		}
	}
	// dwListCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ListCount); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMedia {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONA,pointer=ref}*(1))(2:{alias=NTMS_OBJECTINFORMATIONA}(struct))
	{
		if o.Media == nil {
			o.Media = &rsmp.ObjectInformationA{}
		}
		if err := o.Media.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpList {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONA,pointer=ref}*(1))(2:{alias=NTMS_OBJECTINFORMATIONA}[dim:0,size_is=lpdwListBufferSize,length_is=dwListCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.List", sizeInfo[0])
		}
		o.List = make([]*rsmp.ObjectInformationA, sizeInfo[0])
		for i1 := range o.List {
			i1 := i1
			if o.List[i1] == nil {
				o.List[i1] = &rsmp.ObjectInformationA{}
			}
			if err := o.List[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwListBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ListBufferSize); err != nil {
			return err
		}
	}
	// dwListCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ListCount); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpMedia {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONA}*(1))(2:{alias=NTMS_OBJECTINFORMATIONA}(struct))
	{
		if o.Media != nil {
			if err := o.Media.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rsmp.ObjectInformationA{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpList {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONA}*(1))(2:{alias=NTMS_OBJECTINFORMATIONA}[dim:0,size_is=lpdwListBufferSize,length_is=dwListCount](struct))
	{
		dimSize1 := uint64(o.ListBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ListCount)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.List {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.List[i1] != nil {
				if err := o.List[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&rsmp.ObjectInformationA{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.List); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&rsmp.ObjectInformationA{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_CreateNTMSMediaAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpMedia {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONA,pointer=ref}*(1))(2:{alias=NTMS_OBJECTINFORMATIONA}(struct))
	{
		if o.Media == nil {
			o.Media = &rsmp.ObjectInformationA{}
		}
		if err := o.Media.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpList {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONA,pointer=ref}*(1))(2:{alias=NTMS_OBJECTINFORMATIONA}[dim:0,size_is=lpdwListBufferSize,length_is=dwListCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.List", sizeInfo[0])
		}
		o.List = make([]*rsmp.ObjectInformationA, sizeInfo[0])
		for i1 := range o.List {
			i1 := i1
			if o.List[i1] == nil {
				o.List[i1] = &rsmp.ObjectInformationA{}
			}
			if err := o.List[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
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

// CreateNTMSMediaARequest structure represents the CreateNtmsMediaA operation request
type CreateNTMSMediaARequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMedia: A pointer to an NTMS_OBJECTINFORMATIONA (section 2.2.4.20) structure describing
	// the properties of the medium to create.
	Media *rsmp.ObjectInformationA `idl:"name:lpMedia" json:"media"`
	// lpList: An array of NTMS_OBJECTINFORMATIONA (section 2.2.4.20) structures specifying
	// the sides of the new medium.
	List []*rsmp.ObjectInformationA `idl:"name:lpList;size_is:(lpdwListBufferSize);length_is:(dwListCount)" json:"list"`
	// lpdwListBufferSize: A pointer to the size of lpList, in bytes.
	ListBufferSize uint32 `idl:"name:lpdwListBufferSize" json:"list_buffer_size"`
	// dwListCount: The number of elements in the lpList array.
	ListCount uint32 `idl:"name:dwListCount" json:"list_count"`
	// dwOptions: A bitmap of creation options.
	//
	// If a medium with the specified on-media identifier already exists in the system<65>
	// and the client does not want to duplicate the identifier, the client MUST set dwOptions
	// to NTMS_ERROR_ON_DUPLICATION (0x00000001) and the server MUST NOT create a medium
	// with the specified identifier.
	//
	// If a medium with the specified on-media identifier already exists in the system<66>
	// and the client wants to duplicate the identifier, the client MUST set dwOptions to
	// 0x00000000 and the server MUST create a medium with the specified identifier.
	//
	// If a medium with the specified on-media identifier does not exist in the system,<67>
	// there is no change in the server behavior due to this option.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED      | NTMS_USE_ACCESS to the media pool or offline media library is denied; other      |
	//	|                                     | security errors are possible but indicate a security subsystem error.            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY  | An allocation failure occurred during processing.                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER  | Invalid input parameter.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA      | The option NTMS_ERROR_ON_DUPLICATION was provided, and a medium with this        |
	//	|                                     | on-media identifier already exists.                                              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL | The specified media pool either does not exist, or is not a valid import or      |
	//	|                                     | application pool.                                                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE   | The database is inaccessible or damaged.                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL      | The database is full.                                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DB ERROR_MEDIA_INCOMPATIBLE | The number of specified sides does not match the number of sides associated with |
	//	|                                     | the media type of the media pool.                                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	Options uint32 `idl:"name:dwOptions" json:"options"`
}

func (o *CreateNTMSMediaARequest) xxx_ToOp(ctx context.Context, op *xxx_CreateNTMSMediaAOperation) *xxx_CreateNTMSMediaAOperation {
	if op == nil {
		op = &xxx_CreateNTMSMediaAOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Media = o.Media
	op.List = o.List
	op.ListBufferSize = o.ListBufferSize
	op.ListCount = o.ListCount
	op.Options = o.Options
	return op
}

func (o *CreateNTMSMediaARequest) xxx_FromOp(ctx context.Context, op *xxx_CreateNTMSMediaAOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Media = op.Media
	o.List = op.List
	o.ListBufferSize = op.ListBufferSize
	o.ListCount = op.ListCount
	o.Options = op.Options
}
func (o *CreateNTMSMediaARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateNTMSMediaARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNTMSMediaAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateNTMSMediaAResponse structure represents the CreateNtmsMediaA operation response
type CreateNTMSMediaAResponse struct {
	// XXX: lpdwListBufferSize is an implicit input depedency for output parameters
	ListBufferSize uint32 `idl:"name:lpdwListBufferSize" json:"list_buffer_size"`
	// XXX: dwListCount is an implicit input depedency for output parameters
	ListCount uint32 `idl:"name:dwListCount" json:"list_count"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpMedia: A pointer to an NTMS_OBJECTINFORMATIONA (section 2.2.4.20) structure describing
	// the properties of the medium to create.
	Media *rsmp.ObjectInformationA `idl:"name:lpMedia" json:"media"`
	// lpList: An array of NTMS_OBJECTINFORMATIONA (section 2.2.4.20) structures specifying
	// the sides of the new medium.
	List []*rsmp.ObjectInformationA `idl:"name:lpList;size_is:(lpdwListBufferSize);length_is:(dwListCount)" json:"list"`
	// Return: The CreateNtmsMediaA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateNTMSMediaAResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateNTMSMediaAOperation) *xxx_CreateNTMSMediaAOperation {
	if op == nil {
		op = &xxx_CreateNTMSMediaAOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.ListBufferSize == uint32(0) {
		op.ListBufferSize = o.ListBufferSize
	}
	if op.ListCount == uint32(0) {
		op.ListCount = o.ListCount
	}

	op.That = o.That
	op.Media = o.Media
	op.List = o.List
	op.Return = o.Return
	return op
}

func (o *CreateNTMSMediaAResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateNTMSMediaAOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.ListBufferSize = op.ListBufferSize
	o.ListCount = op.ListCount

	o.That = op.That
	o.Media = op.Media
	o.List = op.List
	o.Return = op.Return
}
func (o *CreateNTMSMediaAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateNTMSMediaAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNTMSMediaAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateNTMSMediaWOperation structure represents the CreateNtmsMediaW operation
type xxx_CreateNTMSMediaWOperation struct {
	This           *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Media          *rsmp.ObjectInformationW   `idl:"name:lpMedia" json:"media"`
	List           []*rsmp.ObjectInformationW `idl:"name:lpList;size_is:(lpdwListBufferSize);length_is:(dwListCount)" json:"list"`
	ListBufferSize uint32                     `idl:"name:lpdwListBufferSize" json:"list_buffer_size"`
	ListCount      uint32                     `idl:"name:dwListCount" json:"list_count"`
	Options        uint32                     `idl:"name:dwOptions" json:"options"`
	Return         int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateNTMSMediaWOperation) OpNum() int { return 5 }

func (o *xxx_CreateNTMSMediaWOperation) OpName() string {
	return "/INtmsObjectInfo1/v0/CreateNtmsMediaW"
}

func (o *xxx_CreateNTMSMediaWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.List != nil && o.ListBufferSize == 0 {
		o.ListBufferSize = uint32(len(o.List))
	}
	if o.List != nil && o.ListCount == 0 {
		o.ListCount = uint32(len(o.List))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMedia {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONW}*(1))(2:{alias=NTMS_OBJECTINFORMATIONW}(struct))
	{
		if o.Media != nil {
			if err := o.Media.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rsmp.ObjectInformationW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpList {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONW}*(1))(2:{alias=NTMS_OBJECTINFORMATIONW}[dim:0,size_is=lpdwListBufferSize,length_is=dwListCount](struct))
	{
		dimSize1 := uint64(o.ListBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ListCount)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.List {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.List[i1] != nil {
				if err := o.List[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&rsmp.ObjectInformationW{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.List); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&rsmp.ObjectInformationW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwListBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ListBufferSize); err != nil {
			return err
		}
	}
	// dwListCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ListCount); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMedia {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONW,pointer=ref}*(1))(2:{alias=NTMS_OBJECTINFORMATIONW}(struct))
	{
		if o.Media == nil {
			o.Media = &rsmp.ObjectInformationW{}
		}
		if err := o.Media.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpList {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONW,pointer=ref}*(1))(2:{alias=NTMS_OBJECTINFORMATIONW}[dim:0,size_is=lpdwListBufferSize,length_is=dwListCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.List", sizeInfo[0])
		}
		o.List = make([]*rsmp.ObjectInformationW, sizeInfo[0])
		for i1 := range o.List {
			i1 := i1
			if o.List[i1] == nil {
				o.List[i1] = &rsmp.ObjectInformationW{}
			}
			if err := o.List[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwListBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ListBufferSize); err != nil {
			return err
		}
	}
	// dwListCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ListCount); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpMedia {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONW}*(1))(2:{alias=NTMS_OBJECTINFORMATIONW}(struct))
	{
		if o.Media != nil {
			if err := o.Media.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rsmp.ObjectInformationW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpList {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONW}*(1))(2:{alias=NTMS_OBJECTINFORMATIONW}[dim:0,size_is=lpdwListBufferSize,length_is=dwListCount](struct))
	{
		dimSize1 := uint64(o.ListBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ListCount)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.List {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.List[i1] != nil {
				if err := o.List[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&rsmp.ObjectInformationW{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.List); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&rsmp.ObjectInformationW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_CreateNTMSMediaWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpMedia {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONW,pointer=ref}*(1))(2:{alias=NTMS_OBJECTINFORMATIONW}(struct))
	{
		if o.Media == nil {
			o.Media = &rsmp.ObjectInformationW{}
		}
		if err := o.Media.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpList {in, out} (1:{alias=LPNTMS_OBJECTINFORMATIONW,pointer=ref}*(1))(2:{alias=NTMS_OBJECTINFORMATIONW}[dim:0,size_is=lpdwListBufferSize,length_is=dwListCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.List", sizeInfo[0])
		}
		o.List = make([]*rsmp.ObjectInformationW, sizeInfo[0])
		for i1 := range o.List {
			i1 := i1
			if o.List[i1] == nil {
				o.List[i1] = &rsmp.ObjectInformationW{}
			}
			if err := o.List[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
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

// CreateNTMSMediaWRequest structure represents the CreateNtmsMediaW operation request
type CreateNTMSMediaWRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMedia: A pointer to an NTMS_OBJECTINFORMATIONW (section 2.2.4.21) structure describing
	// the properties of the medium to create.
	Media *rsmp.ObjectInformationW `idl:"name:lpMedia" json:"media"`
	// lpList: An array of NTMS_OBJECTINFORMATIONW (section 2.2.4.21) structures specifying
	// the sides of the new medium.
	List []*rsmp.ObjectInformationW `idl:"name:lpList;size_is:(lpdwListBufferSize);length_is:(dwListCount)" json:"list"`
	// lpdwListBufferSize: A pointer to the size of lpList, in bytes.
	ListBufferSize uint32 `idl:"name:lpdwListBufferSize" json:"list_buffer_size"`
	// dwListCount: The number of elements in the lpList array.
	ListCount uint32 `idl:"name:dwListCount" json:"list_count"`
	// dwOptions: A bitmap of creation options.
	//
	// If a medium with the specified on-media identifier already exists in the system<70>
	// and the client does not want to duplicate the identifier, the client MUST set dwOptions
	// to NTMS_ERROR_ON_DUPLICATION (0x00000001) and the server MUST NOT create a medium
	// with the specified identifier.
	//
	// If a medium with the specified on-media identifier already exists in the system<71>
	// and the client wants to duplicate the identifier, the client MUST set dwOptions to
	// 0x00000000 and the server MUST create a medium with the specified identifier.
	//
	// If a medium with the specified on-media identifier does not exist in the system,<72>there
	// is no change in the server behavior due to this option.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED      | NTMS_USE_ACCESS to the media pool or offline media library is denied; other      |
	//	|                                     | security errors are possible but indicate a security subsystem error.            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY  | An allocation failure occurred during processing.                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER  | Invalid input parameter.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA      | The option NTMS_ERROR_ON_DUPLICATION was provided, and a medium with this        |
	//	|                                     | on-media identifier already exists.                                              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL | The specified media pool either does not exist, or is not a valid import or      |
	//	|                                     | application pool.                                                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE   | The database is inaccessible or damaged.                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL      | The database is full.                                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DB ERROR_MEDIA_INCOMPATIBLE | The number of specified sides does not match the number of sides associated with |
	//	|                                     | the media type of the media pool.                                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	Options uint32 `idl:"name:dwOptions" json:"options"`
}

func (o *CreateNTMSMediaWRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateNTMSMediaWOperation) *xxx_CreateNTMSMediaWOperation {
	if op == nil {
		op = &xxx_CreateNTMSMediaWOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Media = o.Media
	op.List = o.List
	op.ListBufferSize = o.ListBufferSize
	op.ListCount = o.ListCount
	op.Options = o.Options
	return op
}

func (o *CreateNTMSMediaWRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateNTMSMediaWOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Media = op.Media
	o.List = op.List
	o.ListBufferSize = op.ListBufferSize
	o.ListCount = op.ListCount
	o.Options = op.Options
}
func (o *CreateNTMSMediaWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateNTMSMediaWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNTMSMediaWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateNTMSMediaWResponse structure represents the CreateNtmsMediaW operation response
type CreateNTMSMediaWResponse struct {
	// XXX: lpdwListBufferSize is an implicit input depedency for output parameters
	ListBufferSize uint32 `idl:"name:lpdwListBufferSize" json:"list_buffer_size"`
	// XXX: dwListCount is an implicit input depedency for output parameters
	ListCount uint32 `idl:"name:dwListCount" json:"list_count"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpMedia: A pointer to an NTMS_OBJECTINFORMATIONW (section 2.2.4.21) structure describing
	// the properties of the medium to create.
	Media *rsmp.ObjectInformationW `idl:"name:lpMedia" json:"media"`
	// lpList: An array of NTMS_OBJECTINFORMATIONW (section 2.2.4.21) structures specifying
	// the sides of the new medium.
	List []*rsmp.ObjectInformationW `idl:"name:lpList;size_is:(lpdwListBufferSize);length_is:(dwListCount)" json:"list"`
	// Return: The CreateNtmsMediaW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateNTMSMediaWResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateNTMSMediaWOperation) *xxx_CreateNTMSMediaWOperation {
	if op == nil {
		op = &xxx_CreateNTMSMediaWOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.ListBufferSize == uint32(0) {
		op.ListBufferSize = o.ListBufferSize
	}
	if op.ListCount == uint32(0) {
		op.ListCount = o.ListCount
	}

	op.That = o.That
	op.Media = o.Media
	op.List = o.List
	op.Return = o.Return
	return op
}

func (o *CreateNTMSMediaWResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateNTMSMediaWOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.ListBufferSize = op.ListBufferSize
	o.ListCount = op.ListCount

	o.That = op.That
	o.Media = op.Media
	o.List = op.List
	o.Return = op.Return
}
func (o *CreateNTMSMediaWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateNTMSMediaWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNTMSMediaWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
