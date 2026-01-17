package iwbembackuprestore

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IWbemBackupRestore interface identifier c49e32c7-bc8b-11d2-85d4-00105a1f8304
	BackupRestoreIID = &dcom.IID{Data1: 0xc49e32c7, Data2: 0xbc8b, Data3: 0x11d2, Data4: []byte{0x85, 0xd4, 0x00, 0x10, 0x5a, 0x1f, 0x83, 0x04}}
	// Syntax UUID
	BackupRestoreSyntaxUUID = &uuid.UUID{TimeLow: 0xc49e32c7, TimeMid: 0xbc8b, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0xd4, Node: [6]uint8{0x0, 0x10, 0x5a, 0x1f, 0x83, 0x4}}
	// Syntax ID
	BackupRestoreSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: BackupRestoreSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemBackupRestore interface.
type BackupRestoreClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// On the IWbemBackupRestore::Backup method invocation, the server MUST back up the
	// contents of the CIM database.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// In case of failure, the server MUST return an HRESULT whose S (severity) bit is set
	// as specified in [MS-ERREF] section 2.1. The actual HRESULT value is implementation
	// dependent.
	//
	// The IWbemBackupRestore::Backup method MUST be called on the interface that is obtained
	// from the DCOM Remote Protocol activation of a CLSID_WbemBackupRestore interface,
	// as specified in this section.
	Backup(context.Context, *BackupRequest, ...dcerpc.CallOption) (*BackupResponse, error)

	// On the IWbemBackupRestore::Restore method invocation, the server MUST restore the
	// contents of the CIM database.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// If the WBEM_FLAG_BACKUP_RESTORE_FORCE_SHUTDOWN flag is not set, the server MUST return
	// WBEM_E_INVALID_PARAMETER.
	//
	// In case of failure, the server MUST return an HRESULT whose S (severity) bit is set
	// as specified in [MS-ERREF] section 2.1. The actual HRESULT value is implementation
	// dependent.
	Restore(context.Context, *RestoreRequest, ...dcerpc.CallOption) (*RestoreResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) BackupRestoreClient
}

type xxx_DefaultBackupRestoreClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultBackupRestoreClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultBackupRestoreClient) Backup(ctx context.Context, in *BackupRequest, opts ...dcerpc.CallOption) (*BackupResponse, error) {
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
	out := &BackupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBackupRestoreClient) Restore(ctx context.Context, in *RestoreRequest, opts ...dcerpc.CallOption) (*RestoreResponse, error) {
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
	out := &RestoreResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBackupRestoreClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultBackupRestoreClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultBackupRestoreClient) IPID(ctx context.Context, ipid *dcom.IPID) BackupRestoreClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultBackupRestoreClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewBackupRestoreClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (BackupRestoreClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(BackupRestoreSyntaxV0_0))...)
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
	return &xxx_DefaultBackupRestoreClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_BackupOperation structure represents the Backup operation
type xxx_BackupOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	BackupToFile string         `idl:"name:strBackupToFile;string" json:"backup_to_file"`
	Flags        int32          `idl:"name:lFlags" json:"flags"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupOperation) OpNum() int { return 3 }

func (o *xxx_BackupOperation) OpName() string { return "/IWbemBackupRestore/v0/Backup" }

func (o *xxx_BackupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strBackupToFile {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.BackupToFile); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strBackupToFile {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.BackupToFile); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BackupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// BackupRequest structure represents the Backup operation request
type BackupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strBackupToFile: MUST be a UTF-16 string, which MUST contain the name of the file
	// to which the CIM database is backed up. This parameter MUST NOT<58> be NULL.
	BackupToFile string `idl:"name:strBackupToFile;string" json:"backup_to_file"`
	// lFlags:  This parameter is not used, and its value MUST be 0x0.
	Flags int32 `idl:"name:lFlags" json:"flags"`
}

func (o *BackupRequest) xxx_ToOp(ctx context.Context, op *xxx_BackupOperation) *xxx_BackupOperation {
	if op == nil {
		op = &xxx_BackupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BackupToFile = o.BackupToFile
	op.Flags = o.Flags
	return op
}

func (o *BackupRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BackupToFile = op.BackupToFile
	o.Flags = op.Flags
}
func (o *BackupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BackupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupResponse structure represents the Backup operation response
type BackupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Backup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupResponse) xxx_ToOp(ctx context.Context, op *xxx_BackupOperation) *xxx_BackupOperation {
	if op == nil {
		op = &xxx_BackupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *BackupResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *BackupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BackupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RestoreOperation structure represents the Restore operation
type xxx_RestoreOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	RestoreFromFile string         `idl:"name:strRestoreFromFile;string" json:"restore_from_file"`
	Flags           int32          `idl:"name:lFlags" json:"flags"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RestoreOperation) OpNum() int { return 4 }

func (o *xxx_RestoreOperation) OpName() string { return "/IWbemBackupRestore/v0/Restore" }

func (o *xxx_RestoreOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strRestoreFromFile {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.RestoreFromFile); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strRestoreFromFile {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.RestoreFromFile); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RestoreOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RestoreRequest structure represents the Restore operation request
type RestoreRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strRestoreFromFile: MUST be a UTF-16 string that MUST contain the name of the file
	// from which to restore the CIM database. This parameter MUST NOT<60> be NULL.
	RestoreFromFile string `idl:"name:strRestoreFromFile;string" json:"restore_from_file"`
	// lFlags: Flags that affect the behavior of the Restore method. The flags' behavior
	// MUST be interpreted as specified in the following table.
	//
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                    |                                                                                  |
	//	|                       VALUE                        |                                     MEANING                                      |
	//	|                                                    |                                                                                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_BACKUP_RESTORE_FORCE_SHUTDOWN 0x00000001 | If the bit is not set and if there are any active clients, the server MUST NOT   |
	//	|                                                    | perform the restore. If the bit is set, the server MUST shut down any active     |
	//	|                                                    | clients before performing the restore operation.                                 |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
}

func (o *RestoreRequest) xxx_ToOp(ctx context.Context, op *xxx_RestoreOperation) *xxx_RestoreOperation {
	if op == nil {
		op = &xxx_RestoreOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RestoreFromFile = o.RestoreFromFile
	op.Flags = o.Flags
	return op
}

func (o *RestoreRequest) xxx_FromOp(ctx context.Context, op *xxx_RestoreOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RestoreFromFile = op.RestoreFromFile
	o.Flags = op.Flags
}
func (o *RestoreRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RestoreRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RestoreResponse structure represents the Restore operation response
type RestoreResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Restore return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RestoreResponse) xxx_ToOp(ctx context.Context, op *xxx_RestoreOperation) *xxx_RestoreOperation {
	if op == nil {
		op = &xxx_RestoreOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RestoreResponse) xxx_FromOp(ctx context.Context, op *xxx_RestoreOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RestoreResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RestoreResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
