package ifsrmfilescreenbase

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
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmFileScreenBase interface identifier f3637e80-5b22-4a2b-a637-bbb642b41cfc
	FileScreenBaseIID = &dcom.IID{Data1: 0xf3637e80, Data2: 0x5b22, Data3: 0x4a2b, Data4: []byte{0xa6, 0x37, 0xbb, 0xb6, 0x42, 0xb4, 0x1c, 0xfc}}
	// Syntax UUID
	FileScreenBaseSyntaxUUID = &uuid.UUID{TimeLow: 0xf3637e80, TimeMid: 0x5b22, TimeHiAndVersion: 0x4a2b, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0x37, Node: [6]uint8{0xbb, 0xb6, 0x42, 0xb4, 0x1c, 0xfc}}
	// Syntax ID
	FileScreenBaseSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: FileScreenBaseSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmFileScreenBase interface.
type FileScreenBaseClient interface {

	// IFsrmObject retrieval method.
	Object() ifsrmobject.ObjectClient

	// BlockedFileGroups operation.
	GetBlockedFileGroups(context.Context, *GetBlockedFileGroupsRequest, ...dcerpc.CallOption) (*GetBlockedFileGroupsResponse, error)

	// BlockedFileGroups operation.
	SetBlockedFileGroups(context.Context, *SetBlockedFileGroupsRequest, ...dcerpc.CallOption) (*SetBlockedFileGroupsResponse, error)

	// FileScreenFlags operation.
	GetFileScreenFlags(context.Context, *GetFileScreenFlagsRequest, ...dcerpc.CallOption) (*GetFileScreenFlagsResponse, error)

	// FileScreenFlags operation.
	SetFileScreenFlags(context.Context, *SetFileScreenFlagsRequest, ...dcerpc.CallOption) (*SetFileScreenFlagsResponse, error)

	// The CreateAction method creates an action for this file screen object.
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
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | The object already exists.                                                       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The action parameter            |
	//	|                                  | is NULL. The actionType parameter is not a valid type. If actionType is          |
	//	|                                  | FsrmActionType_Unknown, the parameter MUST be considered an invalid value.       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	CreateAction(context.Context, *CreateActionRequest, ...dcerpc.CallOption) (*CreateActionResponse, error)

	// The EnumActions method enumerates all the actions for the file screen object.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+--------------------------------+
	//	|         RETURN          |                                |
	//	|       VALUE/CODE        |          DESCRIPTION           |
	//	|                         |                                |
	//	+-------------------------+--------------------------------+
	//	+-------------------------+--------------------------------+
	//	| 0x80070057 E_INVALIDARG | The actions parameter is NULL. |
	//	+-------------------------+--------------------------------+
	EnumActions(context.Context, *EnumActionsRequest, ...dcerpc.CallOption) (*EnumActionsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) FileScreenBaseClient
}

type xxx_DefaultFileScreenBaseClient struct {
	ifsrmobject.ObjectClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultFileScreenBaseClient) Object() ifsrmobject.ObjectClient {
	return o.ObjectClient
}

func (o *xxx_DefaultFileScreenBaseClient) GetBlockedFileGroups(ctx context.Context, in *GetBlockedFileGroupsRequest, opts ...dcerpc.CallOption) (*GetBlockedFileGroupsResponse, error) {
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
	out := &GetBlockedFileGroupsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenBaseClient) SetBlockedFileGroups(ctx context.Context, in *SetBlockedFileGroupsRequest, opts ...dcerpc.CallOption) (*SetBlockedFileGroupsResponse, error) {
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
	out := &SetBlockedFileGroupsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenBaseClient) GetFileScreenFlags(ctx context.Context, in *GetFileScreenFlagsRequest, opts ...dcerpc.CallOption) (*GetFileScreenFlagsResponse, error) {
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
	out := &GetFileScreenFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenBaseClient) SetFileScreenFlags(ctx context.Context, in *SetFileScreenFlagsRequest, opts ...dcerpc.CallOption) (*SetFileScreenFlagsResponse, error) {
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
	out := &SetFileScreenFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenBaseClient) CreateAction(ctx context.Context, in *CreateActionRequest, opts ...dcerpc.CallOption) (*CreateActionResponse, error) {
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
	out := &CreateActionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenBaseClient) EnumActions(ctx context.Context, in *EnumActionsRequest, opts ...dcerpc.CallOption) (*EnumActionsResponse, error) {
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
	out := &EnumActionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenBaseClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultFileScreenBaseClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultFileScreenBaseClient) IPID(ctx context.Context, ipid *dcom.IPID) FileScreenBaseClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultFileScreenBaseClient{
		ObjectClient: o.ObjectClient.IPID(ctx, ipid),
		cc:           o.cc,
		ipid:         ipid,
	}
}

func NewFileScreenBaseClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (FileScreenBaseClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(FileScreenBaseSyntaxV0_0))...)
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
	return &xxx_DefaultFileScreenBaseClient{
		ObjectClient: base,
		cc:           cc,
		ipid:         ipid,
	}, nil
}

// xxx_GetBlockedFileGroupsOperation structure represents the BlockedFileGroups operation
type xxx_GetBlockedFileGroupsOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	BlockList *fsrm.MutableCollection `idl:"name:blockList" json:"block_list"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBlockedFileGroupsOperation) OpNum() int { return 12 }

func (o *xxx_GetBlockedFileGroupsOperation) OpName() string {
	return "/IFsrmFileScreenBase/v0/BlockedFileGroups"
}

func (o *xxx_GetBlockedFileGroupsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBlockedFileGroupsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetBlockedFileGroupsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetBlockedFileGroupsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBlockedFileGroupsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// blockList {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmMutableCollection}(interface))
	{
		if o.BlockList != nil {
			_ptr_blockList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BlockList != nil {
					if err := o.BlockList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.MutableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BlockList, _ptr_blockList); err != nil {
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

func (o *xxx_GetBlockedFileGroupsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// blockList {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmMutableCollection}(interface))
	{
		_ptr_blockList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BlockList == nil {
				o.BlockList = &fsrm.MutableCollection{}
			}
			if err := o.BlockList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_blockList := func(ptr interface{}) { o.BlockList = *ptr.(**fsrm.MutableCollection) }
		if err := w.ReadPointer(&o.BlockList, _s_blockList, _ptr_blockList); err != nil {
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

// GetBlockedFileGroupsRequest structure represents the BlockedFileGroups operation request
type GetBlockedFileGroupsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetBlockedFileGroupsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetBlockedFileGroupsOperation) *xxx_GetBlockedFileGroupsOperation {
	if op == nil {
		op = &xxx_GetBlockedFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetBlockedFileGroupsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBlockedFileGroupsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetBlockedFileGroupsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetBlockedFileGroupsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBlockedFileGroupsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBlockedFileGroupsResponse structure represents the BlockedFileGroups operation response
type GetBlockedFileGroupsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	BlockList *fsrm.MutableCollection `idl:"name:blockList" json:"block_list"`
	// Return: The BlockedFileGroups return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetBlockedFileGroupsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetBlockedFileGroupsOperation) *xxx_GetBlockedFileGroupsOperation {
	if op == nil {
		op = &xxx_GetBlockedFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.BlockList = op.BlockList
	o.Return = op.Return
	return op
}

func (o *GetBlockedFileGroupsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBlockedFileGroupsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.BlockList = op.BlockList
	o.Return = op.Return
}
func (o *GetBlockedFileGroupsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetBlockedFileGroupsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBlockedFileGroupsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetBlockedFileGroupsOperation structure represents the BlockedFileGroups operation
type xxx_SetBlockedFileGroupsOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	BlockList *fsrm.MutableCollection `idl:"name:blockList" json:"block_list"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetBlockedFileGroupsOperation) OpNum() int { return 13 }

func (o *xxx_SetBlockedFileGroupsOperation) OpName() string {
	return "/IFsrmFileScreenBase/v0/BlockedFileGroups"
}

func (o *xxx_SetBlockedFileGroupsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBlockedFileGroupsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// blockList {in} (1:{pointer=ref}*(1))(2:{alias=IFsrmMutableCollection}(interface))
	{
		if o.BlockList != nil {
			_ptr_blockList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BlockList != nil {
					if err := o.BlockList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.MutableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BlockList, _ptr_blockList); err != nil {
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

func (o *xxx_SetBlockedFileGroupsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// blockList {in} (1:{pointer=ref}*(1))(2:{alias=IFsrmMutableCollection}(interface))
	{
		_ptr_blockList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BlockList == nil {
				o.BlockList = &fsrm.MutableCollection{}
			}
			if err := o.BlockList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_blockList := func(ptr interface{}) { o.BlockList = *ptr.(**fsrm.MutableCollection) }
		if err := w.ReadPointer(&o.BlockList, _s_blockList, _ptr_blockList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBlockedFileGroupsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBlockedFileGroupsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetBlockedFileGroupsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetBlockedFileGroupsRequest structure represents the BlockedFileGroups operation request
type SetBlockedFileGroupsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	BlockList *fsrm.MutableCollection `idl:"name:blockList" json:"block_list"`
}

func (o *SetBlockedFileGroupsRequest) xxx_ToOp(ctx context.Context, op *xxx_SetBlockedFileGroupsOperation) *xxx_SetBlockedFileGroupsOperation {
	if op == nil {
		op = &xxx_SetBlockedFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.BlockList = op.BlockList
	return op
}

func (o *SetBlockedFileGroupsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetBlockedFileGroupsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BlockList = op.BlockList
}
func (o *SetBlockedFileGroupsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetBlockedFileGroupsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetBlockedFileGroupsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetBlockedFileGroupsResponse structure represents the BlockedFileGroups operation response
type SetBlockedFileGroupsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The BlockedFileGroups return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetBlockedFileGroupsResponse) xxx_ToOp(ctx context.Context, op *xxx_SetBlockedFileGroupsOperation) *xxx_SetBlockedFileGroupsOperation {
	if op == nil {
		op = &xxx_SetBlockedFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *SetBlockedFileGroupsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetBlockedFileGroupsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetBlockedFileGroupsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetBlockedFileGroupsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetBlockedFileGroupsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFileScreenFlagsOperation structure represents the FileScreenFlags operation
type xxx_GetFileScreenFlagsOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileScreenFlags int32          `idl:"name:fileScreenFlags" json:"file_screen_flags"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileScreenFlagsOperation) OpNum() int { return 14 }

func (o *xxx_GetFileScreenFlagsOperation) OpName() string {
	return "/IFsrmFileScreenBase/v0/FileScreenFlags"
}

func (o *xxx_GetFileScreenFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileScreenFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFileScreenFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFileScreenFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileScreenFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileScreenFlags {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.FileScreenFlags); err != nil {
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

func (o *xxx_GetFileScreenFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileScreenFlags {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.FileScreenFlags); err != nil {
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

// GetFileScreenFlagsRequest structure represents the FileScreenFlags operation request
type GetFileScreenFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFileScreenFlagsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFileScreenFlagsOperation) *xxx_GetFileScreenFlagsOperation {
	if op == nil {
		op = &xxx_GetFileScreenFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetFileScreenFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileScreenFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFileScreenFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFileScreenFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileScreenFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileScreenFlagsResponse structure represents the FileScreenFlags operation response
type GetFileScreenFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileScreenFlags int32          `idl:"name:fileScreenFlags" json:"file_screen_flags"`
	// Return: The FileScreenFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFileScreenFlagsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFileScreenFlagsOperation) *xxx_GetFileScreenFlagsOperation {
	if op == nil {
		op = &xxx_GetFileScreenFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.FileScreenFlags = op.FileScreenFlags
	o.Return = op.Return
	return op
}

func (o *GetFileScreenFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileScreenFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileScreenFlags = op.FileScreenFlags
	o.Return = op.Return
}
func (o *GetFileScreenFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFileScreenFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileScreenFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFileScreenFlagsOperation structure represents the FileScreenFlags operation
type xxx_SetFileScreenFlagsOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileScreenFlags int32          `idl:"name:fileScreenFlags" json:"file_screen_flags"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFileScreenFlagsOperation) OpNum() int { return 15 }

func (o *xxx_SetFileScreenFlagsOperation) OpName() string {
	return "/IFsrmFileScreenBase/v0/FileScreenFlags"
}

func (o *xxx_SetFileScreenFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileScreenFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fileScreenFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.FileScreenFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileScreenFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fileScreenFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.FileScreenFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileScreenFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileScreenFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFileScreenFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFileScreenFlagsRequest structure represents the FileScreenFlags operation request
type SetFileScreenFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	FileScreenFlags int32          `idl:"name:fileScreenFlags" json:"file_screen_flags"`
}

func (o *SetFileScreenFlagsRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFileScreenFlagsOperation) *xxx_SetFileScreenFlagsOperation {
	if op == nil {
		op = &xxx_SetFileScreenFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.FileScreenFlags = op.FileScreenFlags
	return op
}

func (o *SetFileScreenFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFileScreenFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FileScreenFlags = op.FileScreenFlags
}
func (o *SetFileScreenFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFileScreenFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFileScreenFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFileScreenFlagsResponse structure represents the FileScreenFlags operation response
type SetFileScreenFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FileScreenFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFileScreenFlagsResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFileScreenFlagsOperation) *xxx_SetFileScreenFlagsOperation {
	if op == nil {
		op = &xxx_SetFileScreenFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *SetFileScreenFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFileScreenFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFileScreenFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFileScreenFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFileScreenFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateActionOperation structure represents the CreateAction operation
type xxx_CreateActionOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ActionType fsrm.ActionType `idl:"name:actionType" json:"action_type"`
	Action     *fsrm.Action    `idl:"name:action" json:"action"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateActionOperation) OpNum() int { return 16 }

func (o *xxx_CreateActionOperation) OpName() string { return "/IFsrmFileScreenBase/v0/CreateAction" }

func (o *xxx_CreateActionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateActionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// actionType {in} (1:{alias=FsrmActionType}(enum))
	{
		if err := w.WriteEnum(uint16(o.ActionType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateActionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// actionType {in} (1:{alias=FsrmActionType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ActionType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateActionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateActionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateActionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateActionRequest structure represents the CreateAction operation request
type CreateActionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// actionType: Contains the type of action to be created.
	ActionType fsrm.ActionType `idl:"name:actionType" json:"action_type"`
}

func (o *CreateActionRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateActionOperation) *xxx_CreateActionOperation {
	if op == nil {
		op = &xxx_CreateActionOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.ActionType = op.ActionType
	return op
}

func (o *CreateActionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateActionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ActionType = op.ActionType
}
func (o *CreateActionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateActionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateActionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateActionResponse structure represents the CreateAction operation response
type CreateActionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// action: Pointer to an IFsrmAction interface pointer (section 3.2.4.2.4) that upon
	// completion points to the newly created action. A caller MUST release the object received
	// when the caller is done with it.
	Action *fsrm.Action `idl:"name:action" json:"action"`
	// Return: The CreateAction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateActionResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateActionOperation) *xxx_CreateActionOperation {
	if op == nil {
		op = &xxx_CreateActionOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Action = op.Action
	o.Return = op.Return
	return op
}

func (o *CreateActionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateActionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Action = op.Action
	o.Return = op.Return
}
func (o *CreateActionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateActionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateActionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumActionsOperation structure represents the EnumActions operation
type xxx_EnumActionsOperation struct {
	This    *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Actions *fsrm.Collection `idl:"name:actions" json:"actions"`
	Return  int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumActionsOperation) OpNum() int { return 17 }

func (o *xxx_EnumActionsOperation) OpName() string { return "/IFsrmFileScreenBase/v0/EnumActions" }

func (o *xxx_EnumActionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumActionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnumActionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_EnumActionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumActionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnumActionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// EnumActionsRequest structure represents the EnumActions operation request
type EnumActionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EnumActionsRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumActionsOperation) *xxx_EnumActionsOperation {
	if op == nil {
		op = &xxx_EnumActionsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *EnumActionsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumActionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EnumActionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumActionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumActionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumActionsResponse structure represents the EnumActions operation response
type EnumActionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// actions: Pointer to an IFsrmCollection interface pointer (section 3.2.4.2.1) that,
	// upon completion, contains IFsrmAction pointers of all the actions for the specified
	// action. A caller MUST release the collection received when the caller is done with
	// it. To get the specific action interface for the action, the caller MUST call QueryInterface
	// for the interface corresponding to the action type of the actions.
	Actions *fsrm.Collection `idl:"name:actions" json:"actions"`
	// Return: The EnumActions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumActionsResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumActionsOperation) *xxx_EnumActionsOperation {
	if op == nil {
		op = &xxx_EnumActionsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Actions = op.Actions
	o.Return = op.Return
	return op
}

func (o *EnumActionsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumActionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Actions = op.Actions
	o.Return = op.Return
}
func (o *EnumActionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumActionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumActionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
