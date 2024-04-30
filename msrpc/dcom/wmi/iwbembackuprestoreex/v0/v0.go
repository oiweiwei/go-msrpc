package iwbembackuprestoreex

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iwbembackuprestore "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbembackuprestore/v0"
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
	_ = iwbembackuprestore.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IWbemBackupRestoreEx interface identifier a359dec5-e813-4834-8a2a-ba7f1d777d76
	BackupRestoreExIID = &dcom.IID{Data1: 0xa359dec5, Data2: 0xe813, Data3: 0x4834, Data4: []byte{0x8a, 0x2a, 0xba, 0x7f, 0x1d, 0x77, 0x7d, 0x76}}
	// Syntax UUID
	BackupRestoreExSyntaxUUID = &uuid.UUID{TimeLow: 0xa359dec5, TimeMid: 0xe813, TimeHiAndVersion: 0x4834, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x2a, Node: [6]uint8{0xba, 0x7f, 0x1d, 0x77, 0x7d, 0x76}}
	// Syntax ID
	BackupRestoreExSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: BackupRestoreExSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemBackupRestoreEx interface.
type BackupRestoreExClient interface {

	// IWbemBackupRestore retrieval method.
	BackupRestore() iwbembackuprestore.BackupRestoreClient

	// On the IWbemBackupRestoreEx::Pause method invocation, the server MUST set the IsServerPaused
	// flag to True and MUST persist the CIM database in a consistent state.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// If Pause is called and the IsServerPaused flag is set to True, the server MUST return
	// WBEM_E_INVALID_OPERATION. In case of any other failure, the server MUST return an
	// HRESULT whose S (severity) bit is set as specified in [MS-ERREF] section 2.1. The
	// actual HRESULT value is implementation dependent.
	//
	// The IWbemBackupRestoreEx::Pause method MUST be called on the interface that is obtained
	// from the DCOM Remote Protocol activation of the CLSID_WbemBackupRestore interface,
	// as specified in this section.
	Pause(context.Context, *PauseRequest, ...dcerpc.CallOption) (*PauseResponse, error)

	// On the IWbemBackupRestoreEx::Resume method invocation, the server MUST set the IsServerPaused
	// flag to False.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return a WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// If Resume is called and the IsServerPaused flag is set to False, the server MUST
	// return WBEM_E_INVALID_OPERATION.
	//
	// In case of any other failure, the server MUST return an HRESULT whose S (severity)
	// bit is set as specified in [MS-ERREF] section 2.1. The actual HRESULT value is implementation
	// dependent.
	Resume(context.Context, *ResumeRequest, ...dcerpc.CallOption) (*ResumeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) BackupRestoreExClient
}

type xxx_DefaultBackupRestoreExClient struct {
	iwbembackuprestore.BackupRestoreClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultBackupRestoreExClient) BackupRestore() iwbembackuprestore.BackupRestoreClient {
	return o.BackupRestoreClient
}

func (o *xxx_DefaultBackupRestoreExClient) Pause(ctx context.Context, in *PauseRequest, opts ...dcerpc.CallOption) (*PauseResponse, error) {
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
	out := &PauseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBackupRestoreExClient) Resume(ctx context.Context, in *ResumeRequest, opts ...dcerpc.CallOption) (*ResumeResponse, error) {
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
	out := &ResumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBackupRestoreExClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultBackupRestoreExClient) IPID(ctx context.Context, ipid *dcom.IPID) BackupRestoreExClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultBackupRestoreExClient{
		BackupRestoreClient: o.BackupRestoreClient.IPID(ctx, ipid),
		cc:                  o.cc,
		ipid:                ipid,
	}
}
func NewBackupRestoreExClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (BackupRestoreExClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(BackupRestoreExSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iwbembackuprestore.NewBackupRestoreClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultBackupRestoreExClient{
		BackupRestoreClient: base,
		cc:                  cc,
		ipid:                ipid,
	}, nil
}

// xxx_PauseOperation structure represents the Pause operation
type xxx_PauseOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PauseOperation) OpNum() int { return 5 }

func (o *xxx_PauseOperation) OpName() string { return "/IWbemBackupRestoreEx/v0/Pause" }

func (o *xxx_PauseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PauseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_PauseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PauseRequest structure represents the Pause operation request
type PauseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *PauseRequest) xxx_ToOp(ctx context.Context) *xxx_PauseOperation {
	if o == nil {
		return &xxx_PauseOperation{}
	}
	return &xxx_PauseOperation{
		This: o.This,
	}
}

func (o *PauseRequest) xxx_FromOp(ctx context.Context, op *xxx_PauseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *PauseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PauseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PauseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PauseResponse structure represents the Pause operation response
type PauseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Pause return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PauseResponse) xxx_ToOp(ctx context.Context) *xxx_PauseOperation {
	if o == nil {
		return &xxx_PauseOperation{}
	}
	return &xxx_PauseOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PauseResponse) xxx_FromOp(ctx context.Context, op *xxx_PauseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PauseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PauseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PauseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResumeOperation structure represents the Resume operation
type xxx_ResumeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ResumeOperation) OpNum() int { return 6 }

func (o *xxx_ResumeOperation) OpName() string { return "/IWbemBackupRestoreEx/v0/Resume" }

func (o *xxx_ResumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ResumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ResumeRequest structure represents the Resume operation request
type ResumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ResumeRequest) xxx_ToOp(ctx context.Context) *xxx_ResumeOperation {
	if o == nil {
		return &xxx_ResumeOperation{}
	}
	return &xxx_ResumeOperation{
		This: o.This,
	}
}

func (o *ResumeRequest) xxx_FromOp(ctx context.Context, op *xxx_ResumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ResumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ResumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResumeResponse structure represents the Resume operation response
type ResumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Resume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResumeResponse) xxx_ToOp(ctx context.Context) *xxx_ResumeOperation {
	if o == nil {
		return &xxx_ResumeOperation{}
	}
	return &xxx_ResumeOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ResumeResponse) xxx_FromOp(ctx context.Context, op *xxx_ResumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ResumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ResumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
