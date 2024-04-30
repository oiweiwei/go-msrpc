package ifsrmactioncommand

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
	ifsrmaction "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmaction/v0"
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
	_ = ifsrmaction.GoPackage
	_ = oaut.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmActionCommand interface identifier 12937789-e247-4917-9c20-f3ee9c7ee783
	ActionCommandIID = &dcom.IID{Data1: 0x12937789, Data2: 0xe247, Data3: 0x4917, Data4: []byte{0x9c, 0x20, 0xf3, 0xee, 0x9c, 0x7e, 0xe7, 0x83}}
	// Syntax UUID
	ActionCommandSyntaxUUID = &uuid.UUID{TimeLow: 0x12937789, TimeMid: 0xe247, TimeHiAndVersion: 0x4917, ClockSeqHiAndReserved: 0x9c, ClockSeqLow: 0x20, Node: [6]uint8{0xf3, 0xee, 0x9c, 0x7e, 0xe7, 0x83}}
	// Syntax ID
	ActionCommandSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ActionCommandSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmActionCommand interface.
type ActionCommandClient interface {

	// IFsrmAction retrieval method.
	Action() ifsrmaction.ActionClient

	// ExecutablePath operation.
	GetExecutablePath(context.Context, *GetExecutablePathRequest, ...dcerpc.CallOption) (*GetExecutablePathResponse, error)

	// ExecutablePath operation.
	SetExecutablePath(context.Context, *SetExecutablePathRequest, ...dcerpc.CallOption) (*SetExecutablePathResponse, error)

	// Arguments operation.
	GetArguments(context.Context, *GetArgumentsRequest, ...dcerpc.CallOption) (*GetArgumentsResponse, error)

	// Arguments operation.
	SetArguments(context.Context, *SetArgumentsRequest, ...dcerpc.CallOption) (*SetArgumentsResponse, error)

	// Account operation.
	GetAccount(context.Context, *GetAccountRequest, ...dcerpc.CallOption) (*GetAccountResponse, error)

	// Account operation.
	SetAccount(context.Context, *SetAccountRequest, ...dcerpc.CallOption) (*SetAccountResponse, error)

	// WorkingDirectory operation.
	GetWorkingDirectory(context.Context, *GetWorkingDirectoryRequest, ...dcerpc.CallOption) (*GetWorkingDirectoryResponse, error)

	// WorkingDirectory operation.
	SetWorkingDirectory(context.Context, *SetWorkingDirectoryRequest, ...dcerpc.CallOption) (*SetWorkingDirectoryResponse, error)

	// MonitorCommand operation.
	GetMonitorCommand(context.Context, *GetMonitorCommandRequest, ...dcerpc.CallOption) (*GetMonitorCommandResponse, error)

	// MonitorCommand operation.
	SetMonitorCommand(context.Context, *SetMonitorCommandRequest, ...dcerpc.CallOption) (*SetMonitorCommandResponse, error)

	// KillTimeOut operation.
	GetKillTimeout(context.Context, *GetKillTimeoutRequest, ...dcerpc.CallOption) (*GetKillTimeoutResponse, error)

	// KillTimeOut operation.
	SetKillTimeout(context.Context, *SetKillTimeoutRequest, ...dcerpc.CallOption) (*SetKillTimeoutResponse, error)

	// LogResult operation.
	GetLogResult(context.Context, *GetLogResultRequest, ...dcerpc.CallOption) (*GetLogResultResponse, error)

	// LogResult operation.
	SetLogResult(context.Context, *SetLogResultRequest, ...dcerpc.CallOption) (*SetLogResultResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ActionCommandClient
}

type xxx_DefaultActionCommandClient struct {
	ifsrmaction.ActionClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultActionCommandClient) Action() ifsrmaction.ActionClient {
	return o.ActionClient
}

func (o *xxx_DefaultActionCommandClient) GetExecutablePath(ctx context.Context, in *GetExecutablePathRequest, opts ...dcerpc.CallOption) (*GetExecutablePathResponse, error) {
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
	out := &GetExecutablePathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) SetExecutablePath(ctx context.Context, in *SetExecutablePathRequest, opts ...dcerpc.CallOption) (*SetExecutablePathResponse, error) {
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
	out := &SetExecutablePathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) GetArguments(ctx context.Context, in *GetArgumentsRequest, opts ...dcerpc.CallOption) (*GetArgumentsResponse, error) {
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
	out := &GetArgumentsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) SetArguments(ctx context.Context, in *SetArgumentsRequest, opts ...dcerpc.CallOption) (*SetArgumentsResponse, error) {
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
	out := &SetArgumentsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...dcerpc.CallOption) (*GetAccountResponse, error) {
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
	out := &GetAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) SetAccount(ctx context.Context, in *SetAccountRequest, opts ...dcerpc.CallOption) (*SetAccountResponse, error) {
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
	out := &SetAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) GetWorkingDirectory(ctx context.Context, in *GetWorkingDirectoryRequest, opts ...dcerpc.CallOption) (*GetWorkingDirectoryResponse, error) {
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
	out := &GetWorkingDirectoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) SetWorkingDirectory(ctx context.Context, in *SetWorkingDirectoryRequest, opts ...dcerpc.CallOption) (*SetWorkingDirectoryResponse, error) {
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
	out := &SetWorkingDirectoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) GetMonitorCommand(ctx context.Context, in *GetMonitorCommandRequest, opts ...dcerpc.CallOption) (*GetMonitorCommandResponse, error) {
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
	out := &GetMonitorCommandResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) SetMonitorCommand(ctx context.Context, in *SetMonitorCommandRequest, opts ...dcerpc.CallOption) (*SetMonitorCommandResponse, error) {
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
	out := &SetMonitorCommandResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) GetKillTimeout(ctx context.Context, in *GetKillTimeoutRequest, opts ...dcerpc.CallOption) (*GetKillTimeoutResponse, error) {
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
	out := &GetKillTimeoutResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) SetKillTimeout(ctx context.Context, in *SetKillTimeoutRequest, opts ...dcerpc.CallOption) (*SetKillTimeoutResponse, error) {
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
	out := &SetKillTimeoutResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) GetLogResult(ctx context.Context, in *GetLogResultRequest, opts ...dcerpc.CallOption) (*GetLogResultResponse, error) {
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
	out := &GetLogResultResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) SetLogResult(ctx context.Context, in *SetLogResultRequest, opts ...dcerpc.CallOption) (*SetLogResultResponse, error) {
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
	out := &SetLogResultResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionCommandClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultActionCommandClient) IPID(ctx context.Context, ipid *dcom.IPID) ActionCommandClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultActionCommandClient{
		ActionClient: o.ActionClient.IPID(ctx, ipid),
		cc:           o.cc,
		ipid:         ipid,
	}
}
func NewActionCommandClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ActionCommandClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ActionCommandSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmaction.NewActionClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultActionCommandClient{
		ActionClient: base,
		cc:           cc,
		ipid:         ipid,
	}, nil
}

// xxx_GetExecutablePathOperation structure represents the ExecutablePath operation
type xxx_GetExecutablePathOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	ExecutablePath *oaut.String   `idl:"name:executablePath" json:"executable_path"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetExecutablePathOperation) OpNum() int { return 12 }

func (o *xxx_GetExecutablePathOperation) OpName() string {
	return "/IFsrmActionCommand/v0/ExecutablePath"
}

func (o *xxx_GetExecutablePathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExecutablePathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetExecutablePathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetExecutablePathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExecutablePathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// executablePath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ExecutablePath != nil {
			_ptr_executablePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ExecutablePath != nil {
					if err := o.ExecutablePath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExecutablePath, _ptr_executablePath); err != nil {
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

func (o *xxx_GetExecutablePathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// executablePath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_executablePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ExecutablePath == nil {
				o.ExecutablePath = &oaut.String{}
			}
			if err := o.ExecutablePath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_executablePath := func(ptr interface{}) { o.ExecutablePath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ExecutablePath, _s_executablePath, _ptr_executablePath); err != nil {
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

// GetExecutablePathRequest structure represents the ExecutablePath operation request
type GetExecutablePathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetExecutablePathRequest) xxx_ToOp(ctx context.Context) *xxx_GetExecutablePathOperation {
	if o == nil {
		return &xxx_GetExecutablePathOperation{}
	}
	return &xxx_GetExecutablePathOperation{
		This: o.This,
	}
}

func (o *GetExecutablePathRequest) xxx_FromOp(ctx context.Context, op *xxx_GetExecutablePathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetExecutablePathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetExecutablePathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExecutablePathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetExecutablePathResponse structure represents the ExecutablePath operation response
type GetExecutablePathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	ExecutablePath *oaut.String   `idl:"name:executablePath" json:"executable_path"`
	// Return: The ExecutablePath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetExecutablePathResponse) xxx_ToOp(ctx context.Context) *xxx_GetExecutablePathOperation {
	if o == nil {
		return &xxx_GetExecutablePathOperation{}
	}
	return &xxx_GetExecutablePathOperation{
		That:           o.That,
		ExecutablePath: o.ExecutablePath,
		Return:         o.Return,
	}
}

func (o *GetExecutablePathResponse) xxx_FromOp(ctx context.Context, op *xxx_GetExecutablePathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ExecutablePath = op.ExecutablePath
	o.Return = op.Return
}
func (o *GetExecutablePathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetExecutablePathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExecutablePathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetExecutablePathOperation structure represents the ExecutablePath operation
type xxx_SetExecutablePathOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	ExecutablePath *oaut.String   `idl:"name:executablePath" json:"executable_path"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetExecutablePathOperation) OpNum() int { return 13 }

func (o *xxx_SetExecutablePathOperation) OpName() string {
	return "/IFsrmActionCommand/v0/ExecutablePath"
}

func (o *xxx_SetExecutablePathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExecutablePathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// executablePath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ExecutablePath != nil {
			_ptr_executablePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ExecutablePath != nil {
					if err := o.ExecutablePath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExecutablePath, _ptr_executablePath); err != nil {
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

func (o *xxx_SetExecutablePathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// executablePath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_executablePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ExecutablePath == nil {
				o.ExecutablePath = &oaut.String{}
			}
			if err := o.ExecutablePath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_executablePath := func(ptr interface{}) { o.ExecutablePath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ExecutablePath, _s_executablePath, _ptr_executablePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExecutablePathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExecutablePathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetExecutablePathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetExecutablePathRequest structure represents the ExecutablePath operation request
type SetExecutablePathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	ExecutablePath *oaut.String   `idl:"name:executablePath" json:"executable_path"`
}

func (o *SetExecutablePathRequest) xxx_ToOp(ctx context.Context) *xxx_SetExecutablePathOperation {
	if o == nil {
		return &xxx_SetExecutablePathOperation{}
	}
	return &xxx_SetExecutablePathOperation{
		This:           o.This,
		ExecutablePath: o.ExecutablePath,
	}
}

func (o *SetExecutablePathRequest) xxx_FromOp(ctx context.Context, op *xxx_SetExecutablePathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ExecutablePath = op.ExecutablePath
}
func (o *SetExecutablePathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetExecutablePathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExecutablePathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetExecutablePathResponse structure represents the ExecutablePath operation response
type SetExecutablePathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExecutablePath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetExecutablePathResponse) xxx_ToOp(ctx context.Context) *xxx_SetExecutablePathOperation {
	if o == nil {
		return &xxx_SetExecutablePathOperation{}
	}
	return &xxx_SetExecutablePathOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetExecutablePathResponse) xxx_FromOp(ctx context.Context, op *xxx_SetExecutablePathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetExecutablePathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetExecutablePathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExecutablePathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetArgumentsOperation structure represents the Arguments operation
type xxx_GetArgumentsOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Arguments *oaut.String   `idl:"name:arguments" json:"arguments"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetArgumentsOperation) OpNum() int { return 14 }

func (o *xxx_GetArgumentsOperation) OpName() string { return "/IFsrmActionCommand/v0/Arguments" }

func (o *xxx_GetArgumentsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetArgumentsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetArgumentsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetArgumentsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetArgumentsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// arguments {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Arguments != nil {
			_ptr_arguments := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Arguments != nil {
					if err := o.Arguments.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Arguments, _ptr_arguments); err != nil {
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

func (o *xxx_GetArgumentsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// arguments {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_arguments := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Arguments == nil {
				o.Arguments = &oaut.String{}
			}
			if err := o.Arguments.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_arguments := func(ptr interface{}) { o.Arguments = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Arguments, _s_arguments, _ptr_arguments); err != nil {
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

// GetArgumentsRequest structure represents the Arguments operation request
type GetArgumentsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetArgumentsRequest) xxx_ToOp(ctx context.Context) *xxx_GetArgumentsOperation {
	if o == nil {
		return &xxx_GetArgumentsOperation{}
	}
	return &xxx_GetArgumentsOperation{
		This: o.This,
	}
}

func (o *GetArgumentsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetArgumentsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetArgumentsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetArgumentsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetArgumentsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetArgumentsResponse structure represents the Arguments operation response
type GetArgumentsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Arguments *oaut.String   `idl:"name:arguments" json:"arguments"`
	// Return: The Arguments return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetArgumentsResponse) xxx_ToOp(ctx context.Context) *xxx_GetArgumentsOperation {
	if o == nil {
		return &xxx_GetArgumentsOperation{}
	}
	return &xxx_GetArgumentsOperation{
		That:      o.That,
		Arguments: o.Arguments,
		Return:    o.Return,
	}
}

func (o *GetArgumentsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetArgumentsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Arguments = op.Arguments
	o.Return = op.Return
}
func (o *GetArgumentsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetArgumentsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetArgumentsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetArgumentsOperation structure represents the Arguments operation
type xxx_SetArgumentsOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Arguments *oaut.String   `idl:"name:arguments" json:"arguments"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetArgumentsOperation) OpNum() int { return 15 }

func (o *xxx_SetArgumentsOperation) OpName() string { return "/IFsrmActionCommand/v0/Arguments" }

func (o *xxx_SetArgumentsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetArgumentsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// arguments {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Arguments != nil {
			_ptr_arguments := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Arguments != nil {
					if err := o.Arguments.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Arguments, _ptr_arguments); err != nil {
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

func (o *xxx_SetArgumentsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// arguments {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_arguments := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Arguments == nil {
				o.Arguments = &oaut.String{}
			}
			if err := o.Arguments.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_arguments := func(ptr interface{}) { o.Arguments = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Arguments, _s_arguments, _ptr_arguments); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetArgumentsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetArgumentsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetArgumentsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetArgumentsRequest structure represents the Arguments operation request
type SetArgumentsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	Arguments *oaut.String   `idl:"name:arguments" json:"arguments"`
}

func (o *SetArgumentsRequest) xxx_ToOp(ctx context.Context) *xxx_SetArgumentsOperation {
	if o == nil {
		return &xxx_SetArgumentsOperation{}
	}
	return &xxx_SetArgumentsOperation{
		This:      o.This,
		Arguments: o.Arguments,
	}
}

func (o *SetArgumentsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetArgumentsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Arguments = op.Arguments
}
func (o *SetArgumentsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetArgumentsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetArgumentsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetArgumentsResponse structure represents the Arguments operation response
type SetArgumentsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Arguments return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetArgumentsResponse) xxx_ToOp(ctx context.Context) *xxx_SetArgumentsOperation {
	if o == nil {
		return &xxx_SetArgumentsOperation{}
	}
	return &xxx_SetArgumentsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetArgumentsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetArgumentsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetArgumentsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetArgumentsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetArgumentsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAccountOperation structure represents the Account operation
type xxx_GetAccountOperation struct {
	This    *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Account fsrm.AccountType `idl:"name:account" json:"account"`
	Return  int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAccountOperation) OpNum() int { return 16 }

func (o *xxx_GetAccountOperation) OpName() string { return "/IFsrmActionCommand/v0/Account" }

func (o *xxx_GetAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// account {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmAccountType}(enum))
	{
		if err := w.WriteData(uint16(o.Account)); err != nil {
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

func (o *xxx_GetAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// account {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmAccountType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Account)); err != nil {
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

// GetAccountRequest structure represents the Account operation request
type GetAccountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAccountRequest) xxx_ToOp(ctx context.Context) *xxx_GetAccountOperation {
	if o == nil {
		return &xxx_GetAccountOperation{}
	}
	return &xxx_GetAccountOperation{
		This: o.This,
	}
}

func (o *GetAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAccountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAccountResponse structure represents the Account operation response
type GetAccountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Account fsrm.AccountType `idl:"name:account" json:"account"`
	// Return: The Account return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAccountResponse) xxx_ToOp(ctx context.Context) *xxx_GetAccountOperation {
	if o == nil {
		return &xxx_GetAccountOperation{}
	}
	return &xxx_GetAccountOperation{
		That:    o.That,
		Account: o.Account,
		Return:  o.Return,
	}
}

func (o *GetAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAccountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Account = op.Account
	o.Return = op.Return
}
func (o *GetAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAccountOperation structure represents the Account operation
type xxx_SetAccountOperation struct {
	This    *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Account fsrm.AccountType `idl:"name:account" json:"account"`
	Return  int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAccountOperation) OpNum() int { return 17 }

func (o *xxx_SetAccountOperation) OpName() string { return "/IFsrmActionCommand/v0/Account" }

func (o *xxx_SetAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// account {in} (1:{alias=FsrmAccountType}(enum))
	{
		if err := w.WriteData(uint16(o.Account)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// account {in} (1:{alias=FsrmAccountType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Account)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAccountRequest structure represents the Account operation request
type SetAccountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis   `idl:"name:This" json:"this"`
	Account fsrm.AccountType `idl:"name:account" json:"account"`
}

func (o *SetAccountRequest) xxx_ToOp(ctx context.Context) *xxx_SetAccountOperation {
	if o == nil {
		return &xxx_SetAccountOperation{}
	}
	return &xxx_SetAccountOperation{
		This:    o.This,
		Account: o.Account,
	}
}

func (o *SetAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAccountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Account = op.Account
}
func (o *SetAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAccountResponse structure represents the Account operation response
type SetAccountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Account return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAccountResponse) xxx_ToOp(ctx context.Context) *xxx_SetAccountOperation {
	if o == nil {
		return &xxx_SetAccountOperation{}
	}
	return &xxx_SetAccountOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAccountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetWorkingDirectoryOperation structure represents the WorkingDirectory operation
type xxx_GetWorkingDirectoryOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	WorkingDirectory *oaut.String   `idl:"name:workingDirectory" json:"working_directory"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetWorkingDirectoryOperation) OpNum() int { return 18 }

func (o *xxx_GetWorkingDirectoryOperation) OpName() string {
	return "/IFsrmActionCommand/v0/WorkingDirectory"
}

func (o *xxx_GetWorkingDirectoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetWorkingDirectoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetWorkingDirectoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetWorkingDirectoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetWorkingDirectoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// workingDirectory {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.WorkingDirectory != nil {
			_ptr_workingDirectory := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WorkingDirectory != nil {
					if err := o.WorkingDirectory.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WorkingDirectory, _ptr_workingDirectory); err != nil {
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

func (o *xxx_GetWorkingDirectoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// workingDirectory {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_workingDirectory := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WorkingDirectory == nil {
				o.WorkingDirectory = &oaut.String{}
			}
			if err := o.WorkingDirectory.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_workingDirectory := func(ptr interface{}) { o.WorkingDirectory = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.WorkingDirectory, _s_workingDirectory, _ptr_workingDirectory); err != nil {
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

// GetWorkingDirectoryRequest structure represents the WorkingDirectory operation request
type GetWorkingDirectoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetWorkingDirectoryRequest) xxx_ToOp(ctx context.Context) *xxx_GetWorkingDirectoryOperation {
	if o == nil {
		return &xxx_GetWorkingDirectoryOperation{}
	}
	return &xxx_GetWorkingDirectoryOperation{
		This: o.This,
	}
}

func (o *GetWorkingDirectoryRequest) xxx_FromOp(ctx context.Context, op *xxx_GetWorkingDirectoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetWorkingDirectoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetWorkingDirectoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetWorkingDirectoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetWorkingDirectoryResponse structure represents the WorkingDirectory operation response
type GetWorkingDirectoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	WorkingDirectory *oaut.String   `idl:"name:workingDirectory" json:"working_directory"`
	// Return: The WorkingDirectory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetWorkingDirectoryResponse) xxx_ToOp(ctx context.Context) *xxx_GetWorkingDirectoryOperation {
	if o == nil {
		return &xxx_GetWorkingDirectoryOperation{}
	}
	return &xxx_GetWorkingDirectoryOperation{
		That:             o.That,
		WorkingDirectory: o.WorkingDirectory,
		Return:           o.Return,
	}
}

func (o *GetWorkingDirectoryResponse) xxx_FromOp(ctx context.Context, op *xxx_GetWorkingDirectoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.WorkingDirectory = op.WorkingDirectory
	o.Return = op.Return
}
func (o *GetWorkingDirectoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetWorkingDirectoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetWorkingDirectoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetWorkingDirectoryOperation structure represents the WorkingDirectory operation
type xxx_SetWorkingDirectoryOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	WorkingDirectory *oaut.String   `idl:"name:workingDirectory" json:"working_directory"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetWorkingDirectoryOperation) OpNum() int { return 19 }

func (o *xxx_SetWorkingDirectoryOperation) OpName() string {
	return "/IFsrmActionCommand/v0/WorkingDirectory"
}

func (o *xxx_SetWorkingDirectoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetWorkingDirectoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// workingDirectory {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.WorkingDirectory != nil {
			_ptr_workingDirectory := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WorkingDirectory != nil {
					if err := o.WorkingDirectory.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WorkingDirectory, _ptr_workingDirectory); err != nil {
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

func (o *xxx_SetWorkingDirectoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// workingDirectory {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_workingDirectory := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WorkingDirectory == nil {
				o.WorkingDirectory = &oaut.String{}
			}
			if err := o.WorkingDirectory.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_workingDirectory := func(ptr interface{}) { o.WorkingDirectory = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.WorkingDirectory, _s_workingDirectory, _ptr_workingDirectory); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetWorkingDirectoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetWorkingDirectoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetWorkingDirectoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetWorkingDirectoryRequest structure represents the WorkingDirectory operation request
type SetWorkingDirectoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	WorkingDirectory *oaut.String   `idl:"name:workingDirectory" json:"working_directory"`
}

func (o *SetWorkingDirectoryRequest) xxx_ToOp(ctx context.Context) *xxx_SetWorkingDirectoryOperation {
	if o == nil {
		return &xxx_SetWorkingDirectoryOperation{}
	}
	return &xxx_SetWorkingDirectoryOperation{
		This:             o.This,
		WorkingDirectory: o.WorkingDirectory,
	}
}

func (o *SetWorkingDirectoryRequest) xxx_FromOp(ctx context.Context, op *xxx_SetWorkingDirectoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WorkingDirectory = op.WorkingDirectory
}
func (o *SetWorkingDirectoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetWorkingDirectoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetWorkingDirectoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetWorkingDirectoryResponse structure represents the WorkingDirectory operation response
type SetWorkingDirectoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The WorkingDirectory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetWorkingDirectoryResponse) xxx_ToOp(ctx context.Context) *xxx_SetWorkingDirectoryOperation {
	if o == nil {
		return &xxx_SetWorkingDirectoryOperation{}
	}
	return &xxx_SetWorkingDirectoryOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetWorkingDirectoryResponse) xxx_FromOp(ctx context.Context, op *xxx_SetWorkingDirectoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetWorkingDirectoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetWorkingDirectoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetWorkingDirectoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMonitorCommandOperation structure represents the MonitorCommand operation
type xxx_GetMonitorCommandOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	MonitorCommand int16          `idl:"name:monitorCommand" json:"monitor_command"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMonitorCommandOperation) OpNum() int { return 20 }

func (o *xxx_GetMonitorCommandOperation) OpName() string {
	return "/IFsrmActionCommand/v0/MonitorCommand"
}

func (o *xxx_GetMonitorCommandOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMonitorCommandOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMonitorCommandOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMonitorCommandOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMonitorCommandOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// monitorCommand {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.MonitorCommand); err != nil {
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

func (o *xxx_GetMonitorCommandOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// monitorCommand {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.MonitorCommand); err != nil {
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

// GetMonitorCommandRequest structure represents the MonitorCommand operation request
type GetMonitorCommandRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMonitorCommandRequest) xxx_ToOp(ctx context.Context) *xxx_GetMonitorCommandOperation {
	if o == nil {
		return &xxx_GetMonitorCommandOperation{}
	}
	return &xxx_GetMonitorCommandOperation{
		This: o.This,
	}
}

func (o *GetMonitorCommandRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMonitorCommandOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMonitorCommandRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMonitorCommandRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMonitorCommandOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMonitorCommandResponse structure represents the MonitorCommand operation response
type GetMonitorCommandResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	MonitorCommand int16          `idl:"name:monitorCommand" json:"monitor_command"`
	// Return: The MonitorCommand return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMonitorCommandResponse) xxx_ToOp(ctx context.Context) *xxx_GetMonitorCommandOperation {
	if o == nil {
		return &xxx_GetMonitorCommandOperation{}
	}
	return &xxx_GetMonitorCommandOperation{
		That:           o.That,
		MonitorCommand: o.MonitorCommand,
		Return:         o.Return,
	}
}

func (o *GetMonitorCommandResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMonitorCommandOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MonitorCommand = op.MonitorCommand
	o.Return = op.Return
}
func (o *GetMonitorCommandResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMonitorCommandResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMonitorCommandOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMonitorCommandOperation structure represents the MonitorCommand operation
type xxx_SetMonitorCommandOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	MonitorCommand int16          `idl:"name:monitorCommand" json:"monitor_command"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMonitorCommandOperation) OpNum() int { return 21 }

func (o *xxx_SetMonitorCommandOperation) OpName() string {
	return "/IFsrmActionCommand/v0/MonitorCommand"
}

func (o *xxx_SetMonitorCommandOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMonitorCommandOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// monitorCommand {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.MonitorCommand); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMonitorCommandOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// monitorCommand {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.MonitorCommand); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMonitorCommandOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMonitorCommandOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMonitorCommandOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMonitorCommandRequest structure represents the MonitorCommand operation request
type SetMonitorCommandRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	MonitorCommand int16          `idl:"name:monitorCommand" json:"monitor_command"`
}

func (o *SetMonitorCommandRequest) xxx_ToOp(ctx context.Context) *xxx_SetMonitorCommandOperation {
	if o == nil {
		return &xxx_SetMonitorCommandOperation{}
	}
	return &xxx_SetMonitorCommandOperation{
		This:           o.This,
		MonitorCommand: o.MonitorCommand,
	}
}

func (o *SetMonitorCommandRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMonitorCommandOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MonitorCommand = op.MonitorCommand
}
func (o *SetMonitorCommandRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetMonitorCommandRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMonitorCommandOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMonitorCommandResponse structure represents the MonitorCommand operation response
type SetMonitorCommandResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MonitorCommand return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMonitorCommandResponse) xxx_ToOp(ctx context.Context) *xxx_SetMonitorCommandOperation {
	if o == nil {
		return &xxx_SetMonitorCommandOperation{}
	}
	return &xxx_SetMonitorCommandOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetMonitorCommandResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMonitorCommandOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMonitorCommandResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetMonitorCommandResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMonitorCommandOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetKillTimeoutOperation structure represents the KillTimeOut operation
type xxx_GetKillTimeoutOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Minutes int32          `idl:"name:minutes" json:"minutes"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetKillTimeoutOperation) OpNum() int { return 22 }

func (o *xxx_GetKillTimeoutOperation) OpName() string { return "/IFsrmActionCommand/v0/KillTimeOut" }

func (o *xxx_GetKillTimeoutOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKillTimeoutOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetKillTimeoutOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetKillTimeoutOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKillTimeoutOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// minutes {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Minutes); err != nil {
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

func (o *xxx_GetKillTimeoutOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// minutes {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Minutes); err != nil {
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

// GetKillTimeoutRequest structure represents the KillTimeOut operation request
type GetKillTimeoutRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetKillTimeoutRequest) xxx_ToOp(ctx context.Context) *xxx_GetKillTimeoutOperation {
	if o == nil {
		return &xxx_GetKillTimeoutOperation{}
	}
	return &xxx_GetKillTimeoutOperation{
		This: o.This,
	}
}

func (o *GetKillTimeoutRequest) xxx_FromOp(ctx context.Context, op *xxx_GetKillTimeoutOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetKillTimeoutRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetKillTimeoutRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetKillTimeoutOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetKillTimeoutResponse structure represents the KillTimeOut operation response
type GetKillTimeoutResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Minutes int32          `idl:"name:minutes" json:"minutes"`
	// Return: The KillTimeOut return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetKillTimeoutResponse) xxx_ToOp(ctx context.Context) *xxx_GetKillTimeoutOperation {
	if o == nil {
		return &xxx_GetKillTimeoutOperation{}
	}
	return &xxx_GetKillTimeoutOperation{
		That:    o.That,
		Minutes: o.Minutes,
		Return:  o.Return,
	}
}

func (o *GetKillTimeoutResponse) xxx_FromOp(ctx context.Context, op *xxx_GetKillTimeoutOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Minutes = op.Minutes
	o.Return = op.Return
}
func (o *GetKillTimeoutResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetKillTimeoutResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetKillTimeoutOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetKillTimeoutOperation structure represents the KillTimeOut operation
type xxx_SetKillTimeoutOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Minutes int32          `idl:"name:minutes" json:"minutes"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetKillTimeoutOperation) OpNum() int { return 23 }

func (o *xxx_SetKillTimeoutOperation) OpName() string { return "/IFsrmActionCommand/v0/KillTimeOut" }

func (o *xxx_SetKillTimeoutOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetKillTimeoutOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// minutes {in} (1:(int32))
	{
		if err := w.WriteData(o.Minutes); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetKillTimeoutOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// minutes {in} (1:(int32))
	{
		if err := w.ReadData(&o.Minutes); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetKillTimeoutOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetKillTimeoutOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetKillTimeoutOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetKillTimeoutRequest structure represents the KillTimeOut operation request
type SetKillTimeoutRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Minutes int32          `idl:"name:minutes" json:"minutes"`
}

func (o *SetKillTimeoutRequest) xxx_ToOp(ctx context.Context) *xxx_SetKillTimeoutOperation {
	if o == nil {
		return &xxx_SetKillTimeoutOperation{}
	}
	return &xxx_SetKillTimeoutOperation{
		This:    o.This,
		Minutes: o.Minutes,
	}
}

func (o *SetKillTimeoutRequest) xxx_FromOp(ctx context.Context, op *xxx_SetKillTimeoutOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Minutes = op.Minutes
}
func (o *SetKillTimeoutRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetKillTimeoutRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetKillTimeoutOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetKillTimeoutResponse structure represents the KillTimeOut operation response
type SetKillTimeoutResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The KillTimeOut return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetKillTimeoutResponse) xxx_ToOp(ctx context.Context) *xxx_SetKillTimeoutOperation {
	if o == nil {
		return &xxx_SetKillTimeoutOperation{}
	}
	return &xxx_SetKillTimeoutOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetKillTimeoutResponse) xxx_FromOp(ctx context.Context, op *xxx_SetKillTimeoutOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetKillTimeoutResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetKillTimeoutResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetKillTimeoutOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLogResultOperation structure represents the LogResult operation
type xxx_GetLogResultOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	LogResults int16          `idl:"name:logResults" json:"log_results"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLogResultOperation) OpNum() int { return 24 }

func (o *xxx_GetLogResultOperation) OpName() string { return "/IFsrmActionCommand/v0/LogResult" }

func (o *xxx_GetLogResultOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogResultOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLogResultOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLogResultOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogResultOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// logResults {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.LogResults); err != nil {
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

func (o *xxx_GetLogResultOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// logResults {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.LogResults); err != nil {
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

// GetLogResultRequest structure represents the LogResult operation request
type GetLogResultRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLogResultRequest) xxx_ToOp(ctx context.Context) *xxx_GetLogResultOperation {
	if o == nil {
		return &xxx_GetLogResultOperation{}
	}
	return &xxx_GetLogResultOperation{
		This: o.This,
	}
}

func (o *GetLogResultRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLogResultOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLogResultRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetLogResultRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogResultOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLogResultResponse structure represents the LogResult operation response
type GetLogResultResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	LogResults int16          `idl:"name:logResults" json:"log_results"`
	// Return: The LogResult return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLogResultResponse) xxx_ToOp(ctx context.Context) *xxx_GetLogResultOperation {
	if o == nil {
		return &xxx_GetLogResultOperation{}
	}
	return &xxx_GetLogResultOperation{
		That:       o.That,
		LogResults: o.LogResults,
		Return:     o.Return,
	}
}

func (o *GetLogResultResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLogResultOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LogResults = op.LogResults
	o.Return = op.Return
}
func (o *GetLogResultResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetLogResultResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogResultOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetLogResultOperation structure represents the LogResult operation
type xxx_SetLogResultOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	LogResults int16          `idl:"name:logResults" json:"log_results"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetLogResultOperation) OpNum() int { return 25 }

func (o *xxx_SetLogResultOperation) OpName() string { return "/IFsrmActionCommand/v0/LogResult" }

func (o *xxx_SetLogResultOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogResultOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// logResults {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.LogResults); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogResultOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// logResults {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.LogResults); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogResultOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLogResultOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetLogResultOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetLogResultRequest structure represents the LogResult operation request
type SetLogResultRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	LogResults int16          `idl:"name:logResults" json:"log_results"`
}

func (o *SetLogResultRequest) xxx_ToOp(ctx context.Context) *xxx_SetLogResultOperation {
	if o == nil {
		return &xxx_SetLogResultOperation{}
	}
	return &xxx_SetLogResultOperation{
		This:       o.This,
		LogResults: o.LogResults,
	}
}

func (o *SetLogResultRequest) xxx_FromOp(ctx context.Context, op *xxx_SetLogResultOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LogResults = op.LogResults
}
func (o *SetLogResultRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetLogResultRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLogResultOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetLogResultResponse structure represents the LogResult operation response
type SetLogResultResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The LogResult return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetLogResultResponse) xxx_ToOp(ctx context.Context) *xxx_SetLogResultOperation {
	if o == nil {
		return &xxx_SetLogResultOperation{}
	}
	return &xxx_SetLogResultOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetLogResultResponse) xxx_FromOp(ctx context.Context, op *xxx_SetLogResultOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetLogResultResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetLogResultResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLogResultOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
