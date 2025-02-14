package eventlog

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// eventlog server interface.
type EventlogServer interface {

	// The ElfrClearELFW (Opnum 0) method instructs the server to clear an event log, and,
	// optionally, to back up the event log before the clear operation takes place.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	ClearEventLogW(context.Context, *ClearEventLogWRequest) (*ClearEventLogWResponse, error)

	// The ElfrBackupELFW (Opnum 1) method instructs the server to back up the event log
	// to a specified file name.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].<25>
	BackupEventLogW(context.Context, *BackupEventLogWRequest) (*BackupEventLogWResponse, error)

	// The ElfrCloseEL (Opnum 2) method instructs the server to close a handle to the event
	// log, freeing resources on the server that maintained an association between the handle
	// and the corresponding event log. This handle MUST have been obtained via the ElfrOpenELW
	// (section 3.1.4.3) method, the ElfrOpenELA (section 3.1.4.4) method, the ElfrOpenBELW
	// (section 3.1.4.1) method, or the ElfrOpenBELA (section 3.1.4.2) method.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	CloseEventLog(context.Context, *CloseEventLogRequest) (*CloseEventLogResponse, error)

	// The ElfrDeregisterEventSource (Opnum 3) method instructs the server to close a handle
	// to the event log, freeing resources on the server that maintained an association
	// between the handle and the corresponding event log. This handle MUST have been obtained
	// via the ElfrRegisterEventSourceW (section 3.1.4.5) method or the ElfrRegisterEventSourceA
	// (section 3.1.4.6) method.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	DeregisterEventSource(context.Context, *DeregisterEventSourceRequest) (*DeregisterEventSourceResponse, error)

	// The ElfrNumberOfRecords (Opnum 4) method instructs the server to report the number
	// of records currently in the event log.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	NumberOfRecords(context.Context, *NumberOfRecordsRequest) (*NumberOfRecordsResponse, error)

	// The ElfrOldestRecord (Opnum 5) method instructs the server to report the record number
	// of the oldest record in the event log.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	OldestRecord(context.Context, *OldestRecordRequest) (*OldestRecordResponse, error)

	// The ElfrChangeNotify (Opnum 6) method is intended for local use.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; the
	// method always returns STATUS_INVALID_HANDLE (0xC0000008) when called remotely. For
	// all other errors it MUST return an implementation-based, nonzero NTSTATUS value specified
	// in [MS-ERREF].
	ChangeNotify(context.Context, *ChangeNotifyRequest) (*ChangeNotifyResponse, error)

	// The ElfrOpenELW method instructs the server to return a server context handle to
	// a live event log. The caller MUST<18> have permission to read the file that contains
	// the event log for this to succeed.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	OpenEventLogW(context.Context, *OpenEventLogWRequest) (*OpenEventLogWResponse, error)

	// The ElfrRegisterEventSourceW (Opnum 8) method instructs the server to return a server
	// context handle to an event log for writing. The caller MUST have permission to write
	// to the file containing the event log for this to succeed. The module name argument
	// specifies the event source that is used to determine the relevant event log as specified
	// below.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	RegisterEventSourceW(context.Context, *RegisterEventSourceWRequest) (*RegisterEventSourceWResponse, error)

	// The ElfrOpenBELW (Opnum 9) method instructs the server to return a handle to a backup
	// event log. The caller MUST have permission to read the file containing the backup
	// event log for this to succeed.
	//
	// Note  The server has an Access Control List (ACL) that is used to control access
	// to the log. The protocol has no methods for reading or setting that ACL.
	//
	// Return Values: The method MUST return STATUS_SUCCESS on success; otherwise, it MUST
	// return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	//
	// Note  The value of STATUS_SUCCESS is 0x00000000.
	OpenBackupEventLogW(context.Context, *OpenBackupEventLogWRequest) (*OpenBackupEventLogWResponse, error)

	// The ElfrReadELW (Opnum 10) method reads events from the event log; the server transmits
	// these events to the client and advances the reader's position within the event log
	// associated with the server context handle that is passed in the LogHandle parameter.
	// The strings in the returned event MUST be in [UNICODE].
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success. If
	// the method is successful, the read position MUST be adjusted by NumberOfBytesRead.
	// The method MUST return STATUS_BUFFER_TOO_SMALL (0xC0000023) if the buffer is too
	// small to fit even one record. Otherwise, it MUST return any other implementation-based,
	// nonzero NTSTATUS value specified in [MS-ERREF].
	ReadEventLogW(context.Context, *ReadEventLogWRequest) (*ReadEventLogWResponse, error)

	// The ElfrReportEventW (Opnum 11) method writes events to the event log; the server
	// receives these events from the client.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	ReportEventW(context.Context, *ReportEventWRequest) (*ReportEventWResponse, error)

	// The ElfrClearELFA (Opnum 12) method instructs the server to clear an event log, and,
	// optionally, to back up the event log before the clear operation takes place.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	//
	// The ElfrClearELFA and ElfrClearELFW (section 3.1.4.9) methods are identical in functionality.
	// The difference between the two methods is that the ElfrClearELFA method specifies
	// BackupFileName as an ANSI string. The ElfrClearELFW method specifies BackupFileName
	// as a UNICODE string.
	ClearEventLogA(context.Context, *ClearEventLogARequest) (*ClearEventLogAResponse, error)

	// The ElfrBackupELFA (Opnum 13) method instructs the server to back up the event log
	// to a specified file name.<26>
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	//
	// ElfrBackupELFA is identical to the ElfrBackupELFW (section 3.1.4.11) method except
	// in the following case:
	//
	// * When BackUpFileName is an ANSI string.
	BackupEventLogA(context.Context, *BackupEventLogARequest) (*BackupEventLogAResponse, error)

	// The ElfrOpenELA (Opnum 14) method instructs the server to return a server context
	// handle to a live event log. For this to succeed, the caller MUST have permission
	// to read the file that contains the event log.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	//
	// This is identical to the ElfrOpenELW (section 3.1.4.3) method except that the ModuleName,
	// RegModuleName, and the UNCServerName are ANSI strings in this case.
	OpenEventLogA(context.Context, *OpenEventLogARequest) (*OpenEventLogAResponse, error)

	// The ElfrRegisterEventSourceA (Opnum 15) method instructs the server to return a server
	// context handle to an event log for writing. The caller MUST have permission to write
	// to the file containing the event log for this to succeed. The module name argument
	// specifies the event source, which is used to determine the relevant event log, as
	// specified in the following sections.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	//
	// This is identical to the ElfrRegisterEventSourceW (section 3.1.4.5) method except
	// that the ModuleName, RegModuleName, and the UNCServerName parameters are ANSI strings
	// in this case.
	RegisterEventSourceA(context.Context, *RegisterEventSourceARequest) (*RegisterEventSourceAResponse, error)

	// The ElfrOpenBELA (Opnum 16) method instructs the server to return a handle to a backup
	// event log. The caller MUST have permission to read the file containing the backup
	// event log for this to succeed.
	//
	// Note  The server has an Access Control List (ACL) that is used to control access
	// to the log. The protocol has no methods for reading or setting that ACL.
	//
	// Return Values: The method returns STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	OpenBackupEventLogA(context.Context, *OpenBackupEventLogARequest) (*OpenBackupEventLogAResponse, error)

	// The ElfrReadELA (Opnum 17) method reads events from the event log; the server transmits
	// these events to the client and advances the reader's position within the event log
	// associated with the server context handle that is passed in the LogHandle parameter.
	// The strings in the returned events MUST be ANSI.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success. The
	// method MUST return STATUS_BUFFER_TOO_SMALL (0xC0000023) if the buffer is too small
	// to fit even one record. Otherwise, it MUST return any other implementation-based,
	// nonzero NTSTATUS value specified in [MS-ERREF].
	//
	// This is identical to the ElfrReadELW (section 3.1.4.7) method except that the events
	// placed in the buffer MUST be ANSI strings rather than [UNICODE] strings.
	ReadEventLogA(context.Context, *ReadEventLogARequest) (*ReadEventLogAResponse, error)

	// The ElfrReportEventA (Opnum 18) method writes events to the event log; the server
	// receives these events from the client.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	//
	// This method is identical to the ElfrReportEventW (section 3.1.4.13) method except
	// that the string arguments are ANSI strings in this case. Thus, the remarks in ElfrReportEventW
	// (section 3.1.4.13) apply to this method as well.
	ReportEventA(context.Context, *ReportEventARequest) (*ReportEventAResponse, error)

	// Opnum19NotUsedOnWire operation.
	// Opnum19NotUsedOnWire

	// Opnum20NotUsedOnWire operation.
	// Opnum20NotUsedOnWire

	// Opnum21NotUsedOnWire operation.
	// Opnum21NotUsedOnWire

	// The ElfrGetLogInformation (Opnum 22) method instructs the server to return information
	// on an event log.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success. The
	// method MUST return STATUS_BUFFER_TOO_SMALL (0xC0000023) if the buffer is too small
	// to fit even one record. Otherwise, it MUST return an implementation-based, nonzero
	// NTSTATUS value specified in [MS-ERREF].
	GetLogInformation(context.Context, *GetLogInformationRequest) (*GetLogInformationResponse, error)

	// Opnum23NotUsedOnWire operation.
	// Opnum23NotUsedOnWire

	// This method<32> instructs the server to write an event to an event log. It differs
	// from the other methods for writing an event by specifying the event source at the
	// time of the write. The other methods for writing an event required the event source
	// to be specified when the handle was opened for write.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	//
	// Note  If the method is not supported, the RPC transport itself (as opposed to this
	// protocol) returns RPC_S_PROCNUM_OUT_OF_RANGE (0x6D1).
	//
	// This method is almost identical to the ElfrReportEventW (section 3.1.4.13) method
	// except that it has a SourceName parameter. The server uses this SourceName parameter
	// to log the source into the event instead of retrieving the source name from the LogHandle
	// parameter.
	ReportEventAndSourceW(context.Context, *ReportEventAndSourceWRequest) (*ReportEventAndSourceWResponse, error)

	// The ElfrReportEventExW (Opnum 25) method<33> writes events to the event log; the
	// server receives these events from the client.
	//
	// Return Values: The method returns STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it returns an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	ReportEventExW(context.Context, *ReportEventExWRequest) (*ReportEventExWResponse, error)

	// The ElfrReportEventExA (Opnum 26) method<36> writes events to the event log; the
	// server receives these events from the client.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	//
	// This method is identical to the ElfrReportEventExW (section 3.1.4.16) method except
	// that the string arguments are ANSI strings in this case. Thus, the remarks in ElfrReportEventExW
	// (section 3.1.4.16) apply to this method as well.
	ReportEventExA(context.Context, *ReportEventExARequest) (*ReportEventExAResponse, error)
}

func RegisterEventlogServer(conn dcerpc.Conn, o EventlogServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventlogServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventlogSyntaxV0_0))...)
}

func NewEventlogServerHandle(o EventlogServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventlogServerHandle(ctx, o, opNum, r)
	}
}

func EventlogServerHandle(ctx context.Context, o EventlogServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // ElfrClearELFW
		op := &xxx_ClearEventLogWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClearEventLogWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClearEventLogW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // ElfrBackupELFW
		op := &xxx_BackupEventLogWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupEventLogWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupEventLogW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // ElfrCloseEL
		op := &xxx_CloseEventLogOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseEventLogRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseEventLog(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // ElfrDeregisterEventSource
		op := &xxx_DeregisterEventSourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeregisterEventSourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeregisterEventSource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // ElfrNumberOfRecords
		op := &xxx_NumberOfRecordsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NumberOfRecordsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NumberOfRecords(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // ElfrOldestRecord
		op := &xxx_OldestRecordOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OldestRecordRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OldestRecord(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // ElfrChangeNotify
		op := &xxx_ChangeNotifyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ChangeNotifyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ChangeNotify(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // ElfrOpenELW
		op := &xxx_OpenEventLogWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenEventLogWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenEventLogW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ElfrRegisterEventSourceW
		op := &xxx_RegisterEventSourceWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterEventSourceWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterEventSourceW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ElfrOpenBELW
		op := &xxx_OpenBackupEventLogWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenBackupEventLogWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenBackupEventLogW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // ElfrReadELW
		op := &xxx_ReadEventLogWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReadEventLogWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReadEventLogW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // ElfrReportEventW
		op := &xxx_ReportEventWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReportEventWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReportEventW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // ElfrClearELFA
		op := &xxx_ClearEventLogAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClearEventLogARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClearEventLogA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // ElfrBackupELFA
		op := &xxx_BackupEventLogAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupEventLogARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupEventLogA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // ElfrOpenELA
		op := &xxx_OpenEventLogAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenEventLogARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenEventLogA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // ElfrRegisterEventSourceA
		op := &xxx_RegisterEventSourceAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterEventSourceARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterEventSourceA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // ElfrOpenBELA
		op := &xxx_OpenBackupEventLogAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenBackupEventLogARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenBackupEventLogA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // ElfrReadELA
		op := &xxx_ReadEventLogAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReadEventLogARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReadEventLogA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // ElfrReportEventA
		op := &xxx_ReportEventAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReportEventARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReportEventA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // Opnum19NotUsedOnWire
		// Opnum19NotUsedOnWire
		return nil, nil
	case 20: // Opnum20NotUsedOnWire
		// Opnum20NotUsedOnWire
		return nil, nil
	case 21: // Opnum21NotUsedOnWire
		// Opnum21NotUsedOnWire
		return nil, nil
	case 22: // ElfrGetLogInformation
		op := &xxx_GetLogInformationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLogInformationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLogInformation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // Opnum23NotUsedOnWire
		// Opnum23NotUsedOnWire
		return nil, nil
	case 24: // ElfrReportEventAndSourceW
		op := &xxx_ReportEventAndSourceWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReportEventAndSourceWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReportEventAndSourceW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // ElfrReportEventExW
		op := &xxx_ReportEventExWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReportEventExWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReportEventExW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // ElfrReportEventExA
		op := &xxx_ReportEventExAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReportEventExARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReportEventExA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented eventlog
type UnimplementedEventlogServer struct {
}

func (UnimplementedEventlogServer) ClearEventLogW(context.Context, *ClearEventLogWRequest) (*ClearEventLogWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) BackupEventLogW(context.Context, *BackupEventLogWRequest) (*BackupEventLogWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) CloseEventLog(context.Context, *CloseEventLogRequest) (*CloseEventLogResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) DeregisterEventSource(context.Context, *DeregisterEventSourceRequest) (*DeregisterEventSourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) NumberOfRecords(context.Context, *NumberOfRecordsRequest) (*NumberOfRecordsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) OldestRecord(context.Context, *OldestRecordRequest) (*OldestRecordResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) ChangeNotify(context.Context, *ChangeNotifyRequest) (*ChangeNotifyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) OpenEventLogW(context.Context, *OpenEventLogWRequest) (*OpenEventLogWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) RegisterEventSourceW(context.Context, *RegisterEventSourceWRequest) (*RegisterEventSourceWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) OpenBackupEventLogW(context.Context, *OpenBackupEventLogWRequest) (*OpenBackupEventLogWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) ReadEventLogW(context.Context, *ReadEventLogWRequest) (*ReadEventLogWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) ReportEventW(context.Context, *ReportEventWRequest) (*ReportEventWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) ClearEventLogA(context.Context, *ClearEventLogARequest) (*ClearEventLogAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) BackupEventLogA(context.Context, *BackupEventLogARequest) (*BackupEventLogAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) OpenEventLogA(context.Context, *OpenEventLogARequest) (*OpenEventLogAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) RegisterEventSourceA(context.Context, *RegisterEventSourceARequest) (*RegisterEventSourceAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) OpenBackupEventLogA(context.Context, *OpenBackupEventLogARequest) (*OpenBackupEventLogAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) ReadEventLogA(context.Context, *ReadEventLogARequest) (*ReadEventLogAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) ReportEventA(context.Context, *ReportEventARequest) (*ReportEventAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) GetLogInformation(context.Context, *GetLogInformationRequest) (*GetLogInformationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) ReportEventAndSourceW(context.Context, *ReportEventAndSourceWRequest) (*ReportEventAndSourceWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) ReportEventExW(context.Context, *ReportEventExWRequest) (*ReportEventExWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventlogServer) ReportEventExA(context.Context, *ReportEventExARequest) (*ReportEventExAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EventlogServer = (*UnimplementedEventlogServer)(nil)
