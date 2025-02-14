package ifsrmfilescreenexception

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
	// IFsrmFileScreenException interface identifier bee7ce02-df77-4515-9389-78f01c5afc1a
	FileScreenExceptionIID = &dcom.IID{Data1: 0xbee7ce02, Data2: 0xdf77, Data3: 0x4515, Data4: []byte{0x93, 0x89, 0x78, 0xf0, 0x1c, 0x5a, 0xfc, 0x1a}}
	// Syntax UUID
	FileScreenExceptionSyntaxUUID = &uuid.UUID{TimeLow: 0xbee7ce02, TimeMid: 0xdf77, TimeHiAndVersion: 0x4515, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x89, Node: [6]uint8{0x78, 0xf0, 0x1c, 0x5a, 0xfc, 0x1a}}
	// Syntax ID
	FileScreenExceptionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: FileScreenExceptionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmFileScreenException interface.
type FileScreenExceptionClient interface {

	// IFsrmObject retrieval method.
	Object() ifsrmobject.ObjectClient

	// Path operation.
	GetPath(context.Context, *GetPathRequest, ...dcerpc.CallOption) (*GetPathResponse, error)

	// AllowedFileGroups operation.
	GetAllowedFileGroups(context.Context, *GetAllowedFileGroupsRequest, ...dcerpc.CallOption) (*GetAllowedFileGroupsResponse, error)

	// AllowedFileGroups operation.
	SetAllowedFileGroups(context.Context, *SetAllowedFileGroupsRequest, ...dcerpc.CallOption) (*SetAllowedFileGroupsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) FileScreenExceptionClient
}

type xxx_DefaultFileScreenExceptionClient struct {
	ifsrmobject.ObjectClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultFileScreenExceptionClient) Object() ifsrmobject.ObjectClient {
	return o.ObjectClient
}

func (o *xxx_DefaultFileScreenExceptionClient) GetPath(ctx context.Context, in *GetPathRequest, opts ...dcerpc.CallOption) (*GetPathResponse, error) {
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
	out := &GetPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenExceptionClient) GetAllowedFileGroups(ctx context.Context, in *GetAllowedFileGroupsRequest, opts ...dcerpc.CallOption) (*GetAllowedFileGroupsResponse, error) {
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
	out := &GetAllowedFileGroupsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenExceptionClient) SetAllowedFileGroups(ctx context.Context, in *SetAllowedFileGroupsRequest, opts ...dcerpc.CallOption) (*SetAllowedFileGroupsResponse, error) {
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
	out := &SetAllowedFileGroupsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenExceptionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultFileScreenExceptionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultFileScreenExceptionClient) IPID(ctx context.Context, ipid *dcom.IPID) FileScreenExceptionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultFileScreenExceptionClient{
		ObjectClient: o.ObjectClient.IPID(ctx, ipid),
		cc:           o.cc,
		ipid:         ipid,
	}
}

func NewFileScreenExceptionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (FileScreenExceptionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(FileScreenExceptionSyntaxV0_0))...)
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
	return &xxx_DefaultFileScreenExceptionClient{
		ObjectClient: base,
		cc:           cc,
		ipid:         ipid,
	}, nil
}

// xxx_GetPathOperation structure represents the Path operation
type xxx_GetPathOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   *oaut.String   `idl:"name:path" json:"path"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPathOperation) OpNum() int { return 12 }

func (o *xxx_GetPathOperation) OpName() string { return "/IFsrmFileScreenException/v0/Path" }

func (o *xxx_GetPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// path {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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

func (o *xxx_GetPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// path {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
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

// GetPathRequest structure represents the Path operation request
type GetPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPathRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPathOperation) *xxx_GetPathOperation {
	if op == nil {
		op = &xxx_GetPathOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetPathRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPathResponse structure represents the Path operation response
type GetPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path *oaut.String   `idl:"name:path" json:"path"`
	// Return: The Path return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPathResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPathOperation) *xxx_GetPathOperation {
	if op == nil {
		op = &xxx_GetPathOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Path = op.Path
	o.Return = op.Return
	return op
}

func (o *GetPathResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Path = op.Path
	o.Return = op.Return
}
func (o *GetPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAllowedFileGroupsOperation structure represents the AllowedFileGroups operation
type xxx_GetAllowedFileGroupsOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	AllowList *fsrm.MutableCollection `idl:"name:allowList" json:"allow_list"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAllowedFileGroupsOperation) OpNum() int { return 13 }

func (o *xxx_GetAllowedFileGroupsOperation) OpName() string {
	return "/IFsrmFileScreenException/v0/AllowedFileGroups"
}

func (o *xxx_GetAllowedFileGroupsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllowedFileGroupsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAllowedFileGroupsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAllowedFileGroupsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllowedFileGroupsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// allowList {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmMutableCollection}(interface))
	{
		if o.AllowList != nil {
			_ptr_allowList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AllowList != nil {
					if err := o.AllowList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.MutableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AllowList, _ptr_allowList); err != nil {
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

func (o *xxx_GetAllowedFileGroupsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// allowList {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmMutableCollection}(interface))
	{
		_ptr_allowList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AllowList == nil {
				o.AllowList = &fsrm.MutableCollection{}
			}
			if err := o.AllowList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_allowList := func(ptr interface{}) { o.AllowList = *ptr.(**fsrm.MutableCollection) }
		if err := w.ReadPointer(&o.AllowList, _s_allowList, _ptr_allowList); err != nil {
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

// GetAllowedFileGroupsRequest structure represents the AllowedFileGroups operation request
type GetAllowedFileGroupsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAllowedFileGroupsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAllowedFileGroupsOperation) *xxx_GetAllowedFileGroupsOperation {
	if op == nil {
		op = &xxx_GetAllowedFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetAllowedFileGroupsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAllowedFileGroupsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAllowedFileGroupsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAllowedFileGroupsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllowedFileGroupsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAllowedFileGroupsResponse structure represents the AllowedFileGroups operation response
type GetAllowedFileGroupsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	AllowList *fsrm.MutableCollection `idl:"name:allowList" json:"allow_list"`
	// Return: The AllowedFileGroups return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAllowedFileGroupsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAllowedFileGroupsOperation) *xxx_GetAllowedFileGroupsOperation {
	if op == nil {
		op = &xxx_GetAllowedFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.AllowList = op.AllowList
	o.Return = op.Return
	return op
}

func (o *GetAllowedFileGroupsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAllowedFileGroupsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AllowList = op.AllowList
	o.Return = op.Return
}
func (o *GetAllowedFileGroupsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAllowedFileGroupsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllowedFileGroupsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAllowedFileGroupsOperation structure represents the AllowedFileGroups operation
type xxx_SetAllowedFileGroupsOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	AllowList *fsrm.MutableCollection `idl:"name:allowList" json:"allow_list"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAllowedFileGroupsOperation) OpNum() int { return 14 }

func (o *xxx_SetAllowedFileGroupsOperation) OpName() string {
	return "/IFsrmFileScreenException/v0/AllowedFileGroups"
}

func (o *xxx_SetAllowedFileGroupsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAllowedFileGroupsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// allowList {in} (1:{pointer=ref}*(1))(2:{alias=IFsrmMutableCollection}(interface))
	{
		if o.AllowList != nil {
			_ptr_allowList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AllowList != nil {
					if err := o.AllowList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.MutableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AllowList, _ptr_allowList); err != nil {
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

func (o *xxx_SetAllowedFileGroupsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// allowList {in} (1:{pointer=ref}*(1))(2:{alias=IFsrmMutableCollection}(interface))
	{
		_ptr_allowList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AllowList == nil {
				o.AllowList = &fsrm.MutableCollection{}
			}
			if err := o.AllowList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_allowList := func(ptr interface{}) { o.AllowList = *ptr.(**fsrm.MutableCollection) }
		if err := w.ReadPointer(&o.AllowList, _s_allowList, _ptr_allowList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAllowedFileGroupsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAllowedFileGroupsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAllowedFileGroupsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAllowedFileGroupsRequest structure represents the AllowedFileGroups operation request
type SetAllowedFileGroupsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	AllowList *fsrm.MutableCollection `idl:"name:allowList" json:"allow_list"`
}

func (o *SetAllowedFileGroupsRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAllowedFileGroupsOperation) *xxx_SetAllowedFileGroupsOperation {
	if op == nil {
		op = &xxx_SetAllowedFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.AllowList = op.AllowList
	return op
}

func (o *SetAllowedFileGroupsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAllowedFileGroupsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AllowList = op.AllowList
}
func (o *SetAllowedFileGroupsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAllowedFileGroupsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAllowedFileGroupsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAllowedFileGroupsResponse structure represents the AllowedFileGroups operation response
type SetAllowedFileGroupsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AllowedFileGroups return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAllowedFileGroupsResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAllowedFileGroupsOperation) *xxx_SetAllowedFileGroupsOperation {
	if op == nil {
		op = &xxx_SetAllowedFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *SetAllowedFileGroupsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAllowedFileGroupsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAllowedFileGroupsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAllowedFileGroupsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAllowedFileGroupsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
