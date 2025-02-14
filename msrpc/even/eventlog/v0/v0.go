package eventlog

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
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
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "even"
)

var (
	// Syntax UUID
	EventlogSyntaxUUID = &uuid.UUID{TimeLow: 0x82273fdc, TimeMid: 0xe32a, TimeHiAndVersion: 0x18c3, ClockSeqHiAndReserved: 0x3f, ClockSeqLow: 0x78, Node: [6]uint8{0x82, 0x79, 0x29, 0xdc, 0x23, 0xea}}
	// Syntax ID
	EventlogSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: EventlogSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// eventlog interface.
type EventlogClient interface {

	// The ElfrClearELFW (Opnum 0) method instructs the server to clear an event log, and,
	// optionally, to back up the event log before the clear operation takes place.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	ClearEventLogW(context.Context, *ClearEventLogWRequest, ...dcerpc.CallOption) (*ClearEventLogWResponse, error)

	// The ElfrBackupELFW (Opnum 1) method instructs the server to back up the event log
	// to a specified file name.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].<25>
	BackupEventLogW(context.Context, *BackupEventLogWRequest, ...dcerpc.CallOption) (*BackupEventLogWResponse, error)

	// The ElfrCloseEL (Opnum 2) method instructs the server to close a handle to the event
	// log, freeing resources on the server that maintained an association between the handle
	// and the corresponding event log. This handle MUST have been obtained via the ElfrOpenELW
	// (section 3.1.4.3) method, the ElfrOpenELA (section 3.1.4.4) method, the ElfrOpenBELW
	// (section 3.1.4.1) method, or the ElfrOpenBELA (section 3.1.4.2) method.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	CloseEventLog(context.Context, *CloseEventLogRequest, ...dcerpc.CallOption) (*CloseEventLogResponse, error)

	// The ElfrDeregisterEventSource (Opnum 3) method instructs the server to close a handle
	// to the event log, freeing resources on the server that maintained an association
	// between the handle and the corresponding event log. This handle MUST have been obtained
	// via the ElfrRegisterEventSourceW (section 3.1.4.5) method or the ElfrRegisterEventSourceA
	// (section 3.1.4.6) method.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	DeregisterEventSource(context.Context, *DeregisterEventSourceRequest, ...dcerpc.CallOption) (*DeregisterEventSourceResponse, error)

	// The ElfrNumberOfRecords (Opnum 4) method instructs the server to report the number
	// of records currently in the event log.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	NumberOfRecords(context.Context, *NumberOfRecordsRequest, ...dcerpc.CallOption) (*NumberOfRecordsResponse, error)

	// The ElfrOldestRecord (Opnum 5) method instructs the server to report the record number
	// of the oldest record in the event log.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	OldestRecord(context.Context, *OldestRecordRequest, ...dcerpc.CallOption) (*OldestRecordResponse, error)

	// The ElfrChangeNotify (Opnum 6) method is intended for local use.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; the
	// method always returns STATUS_INVALID_HANDLE (0xC0000008) when called remotely. For
	// all other errors it MUST return an implementation-based, nonzero NTSTATUS value specified
	// in [MS-ERREF].
	ChangeNotify(context.Context, *ChangeNotifyRequest, ...dcerpc.CallOption) (*ChangeNotifyResponse, error)

	// The ElfrOpenELW method instructs the server to return a server context handle to
	// a live event log. The caller MUST<18> have permission to read the file that contains
	// the event log for this to succeed.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	OpenEventLogW(context.Context, *OpenEventLogWRequest, ...dcerpc.CallOption) (*OpenEventLogWResponse, error)

	// The ElfrRegisterEventSourceW (Opnum 8) method instructs the server to return a server
	// context handle to an event log for writing. The caller MUST have permission to write
	// to the file containing the event log for this to succeed. The module name argument
	// specifies the event source that is used to determine the relevant event log as specified
	// below.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	RegisterEventSourceW(context.Context, *RegisterEventSourceWRequest, ...dcerpc.CallOption) (*RegisterEventSourceWResponse, error)

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
	OpenBackupEventLogW(context.Context, *OpenBackupEventLogWRequest, ...dcerpc.CallOption) (*OpenBackupEventLogWResponse, error)

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
	ReadEventLogW(context.Context, *ReadEventLogWRequest, ...dcerpc.CallOption) (*ReadEventLogWResponse, error)

	// The ElfrReportEventW (Opnum 11) method writes events to the event log; the server
	// receives these events from the client.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	ReportEventW(context.Context, *ReportEventWRequest, ...dcerpc.CallOption) (*ReportEventWResponse, error)

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
	ClearEventLogA(context.Context, *ClearEventLogARequest, ...dcerpc.CallOption) (*ClearEventLogAResponse, error)

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
	BackupEventLogA(context.Context, *BackupEventLogARequest, ...dcerpc.CallOption) (*BackupEventLogAResponse, error)

	// The ElfrOpenELA (Opnum 14) method instructs the server to return a server context
	// handle to a live event log. For this to succeed, the caller MUST have permission
	// to read the file that contains the event log.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	//
	// This is identical to the ElfrOpenELW (section 3.1.4.3) method except that the ModuleName,
	// RegModuleName, and the UNCServerName are ANSI strings in this case.
	OpenEventLogA(context.Context, *OpenEventLogARequest, ...dcerpc.CallOption) (*OpenEventLogAResponse, error)

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
	RegisterEventSourceA(context.Context, *RegisterEventSourceARequest, ...dcerpc.CallOption) (*RegisterEventSourceAResponse, error)

	// The ElfrOpenBELA (Opnum 16) method instructs the server to return a handle to a backup
	// event log. The caller MUST have permission to read the file containing the backup
	// event log for this to succeed.
	//
	// Note  The server has an Access Control List (ACL) that is used to control access
	// to the log. The protocol has no methods for reading or setting that ACL.
	//
	// Return Values: The method returns STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	OpenBackupEventLogA(context.Context, *OpenBackupEventLogARequest, ...dcerpc.CallOption) (*OpenBackupEventLogAResponse, error)

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
	ReadEventLogA(context.Context, *ReadEventLogARequest, ...dcerpc.CallOption) (*ReadEventLogAResponse, error)

	// The ElfrReportEventA (Opnum 18) method writes events to the event log; the server
	// receives these events from the client.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	//
	// This method is identical to the ElfrReportEventW (section 3.1.4.13) method except
	// that the string arguments are ANSI strings in this case. Thus, the remarks in ElfrReportEventW
	// (section 3.1.4.13) apply to this method as well.
	ReportEventA(context.Context, *ReportEventARequest, ...dcerpc.CallOption) (*ReportEventAResponse, error)

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
	GetLogInformation(context.Context, *GetLogInformationRequest, ...dcerpc.CallOption) (*GetLogInformationResponse, error)

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
	ReportEventAndSourceW(context.Context, *ReportEventAndSourceWRequest, ...dcerpc.CallOption) (*ReportEventAndSourceWResponse, error)

	// The ElfrReportEventExW (Opnum 25) method<33> writes events to the event log; the
	// server receives these events from the client.
	//
	// Return Values: The method returns STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it returns an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	ReportEventExW(context.Context, *ReportEventExWRequest, ...dcerpc.CallOption) (*ReportEventExWResponse, error)

	// The ElfrReportEventExA (Opnum 26) method<36> writes events to the event log; the
	// server receives these events from the client.
	//
	// Return Values: The method MUST return STATUS_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-based, nonzero NTSTATUS value specified in [MS-ERREF].
	//
	// This method is identical to the ElfrReportEventExW (section 3.1.4.16) method except
	// that the string arguments are ANSI strings in this case. Thus, the remarks in ElfrReportEventExW
	// (section 3.1.4.16) apply to this method as well.
	ReportEventExA(context.Context, *ReportEventExARequest, ...dcerpc.CallOption) (*ReportEventExAResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// MaxStrings represents the MAX_STRINGS RPC constant
var MaxStrings = 256

// MaxSingleEvent represents the MAX_SINGLE_EVENT RPC constant
var MaxSingleEvent = 262143

// MaxBatchBuffer represents the MAX_BATCH_BUFF RPC constant
var MaxBatchBuffer = 524287

// String structure represents RPC_STRING RPC structure.
//
// EventLog Remoting Protocol APIs use the RPC_STRING to specify an ANSI string parameter.
type String struct {
	// Length:  The number of bytes (not the number of characters) in the string. This does
	// not include the null terminator.
	Length uint16 `idl:"name:Length" json:"length"`
	// MaximumLength:  If the string is the empty string, this MUST be set to zero. Otherwise,
	// it MUST be the number of bytes in the string, including the null terminator (that
	// is, it MUST be equal to the Length member plus 1).
	MaximumLength uint16 `idl:"name:MaximumLength" json:"maximum_length"`
	// Buffer:  Either a pointer to a buffer containing a null-terminated non-empty ANSI
	// string or NULL to indicate an empty string.
	Buffer []byte `idl:"name:Buffer;size_is:(MaximumLength)" json:"buffer"`
}

func (o *String) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.MaximumLength == 0 {
		o.MaximumLength = uint16(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *String) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.MaximumLength); err != nil {
		return err
	}
	if o.Buffer != nil || o.MaximumLength > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.MaximumLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Buffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *String) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumLength); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.MaximumLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.MaximumLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// ClientID structure represents RPC_CLIENT_ID RPC structure.
//
// The RPC_CLIENT_ID structure is used in the ElfrChangeNotify (section 3.1.4.23) method
// for local method invocations only.
type ClientID struct {
	// UniqueProcess:  A 32-bit unsigned integer. Ignored when ElfrChangeNotify (section
	// 3.1.4.23) is invoked remotely.
	UniqueProcess uint32 `idl:"name:UniqueProcess" json:"unique_process"`
	// UniqueThread:  A 32-bit unsigned integer. Ignored when ElfrChangeNotify (section
	// 3.1.4.23) is invoked remotely.
	UniqueThread uint32 `idl:"name:UniqueThread" json:"unique_thread"`
}

func (o *ClientID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.UniqueProcess); err != nil {
		return err
	}
	if err := w.WriteData(o.UniqueThread); err != nil {
		return err
	}
	return nil
}
func (o *ClientID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.UniqueProcess); err != nil {
		return err
	}
	if err := w.ReadData(&o.UniqueThread); err != nil {
		return err
	}
	return nil
}

// Handle structure represents IELF_HANDLE RPC structure.
type Handle dcetypes.ContextHandle

func (o *Handle) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Handle) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Handle) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Handle) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultEventlogClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultEventlogClient) ClearEventLogW(ctx context.Context, in *ClearEventLogWRequest, opts ...dcerpc.CallOption) (*ClearEventLogWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClearEventLogWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) BackupEventLogW(ctx context.Context, in *BackupEventLogWRequest, opts ...dcerpc.CallOption) (*BackupEventLogWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupEventLogWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) CloseEventLog(ctx context.Context, in *CloseEventLogRequest, opts ...dcerpc.CallOption) (*CloseEventLogResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseEventLogResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) DeregisterEventSource(ctx context.Context, in *DeregisterEventSourceRequest, opts ...dcerpc.CallOption) (*DeregisterEventSourceResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeregisterEventSourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) NumberOfRecords(ctx context.Context, in *NumberOfRecordsRequest, opts ...dcerpc.CallOption) (*NumberOfRecordsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NumberOfRecordsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) OldestRecord(ctx context.Context, in *OldestRecordRequest, opts ...dcerpc.CallOption) (*OldestRecordResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OldestRecordResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) ChangeNotify(ctx context.Context, in *ChangeNotifyRequest, opts ...dcerpc.CallOption) (*ChangeNotifyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ChangeNotifyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) OpenEventLogW(ctx context.Context, in *OpenEventLogWRequest, opts ...dcerpc.CallOption) (*OpenEventLogWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenEventLogWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) RegisterEventSourceW(ctx context.Context, in *RegisterEventSourceWRequest, opts ...dcerpc.CallOption) (*RegisterEventSourceWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RegisterEventSourceWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) OpenBackupEventLogW(ctx context.Context, in *OpenBackupEventLogWRequest, opts ...dcerpc.CallOption) (*OpenBackupEventLogWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenBackupEventLogWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) ReadEventLogW(ctx context.Context, in *ReadEventLogWRequest, opts ...dcerpc.CallOption) (*ReadEventLogWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReadEventLogWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) ReportEventW(ctx context.Context, in *ReportEventWRequest, opts ...dcerpc.CallOption) (*ReportEventWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReportEventWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) ClearEventLogA(ctx context.Context, in *ClearEventLogARequest, opts ...dcerpc.CallOption) (*ClearEventLogAResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClearEventLogAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) BackupEventLogA(ctx context.Context, in *BackupEventLogARequest, opts ...dcerpc.CallOption) (*BackupEventLogAResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupEventLogAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) OpenEventLogA(ctx context.Context, in *OpenEventLogARequest, opts ...dcerpc.CallOption) (*OpenEventLogAResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenEventLogAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) RegisterEventSourceA(ctx context.Context, in *RegisterEventSourceARequest, opts ...dcerpc.CallOption) (*RegisterEventSourceAResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RegisterEventSourceAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) OpenBackupEventLogA(ctx context.Context, in *OpenBackupEventLogARequest, opts ...dcerpc.CallOption) (*OpenBackupEventLogAResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenBackupEventLogAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) ReadEventLogA(ctx context.Context, in *ReadEventLogARequest, opts ...dcerpc.CallOption) (*ReadEventLogAResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReadEventLogAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) ReportEventA(ctx context.Context, in *ReportEventARequest, opts ...dcerpc.CallOption) (*ReportEventAResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReportEventAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) GetLogInformation(ctx context.Context, in *GetLogInformationRequest, opts ...dcerpc.CallOption) (*GetLogInformationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetLogInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) ReportEventAndSourceW(ctx context.Context, in *ReportEventAndSourceWRequest, opts ...dcerpc.CallOption) (*ReportEventAndSourceWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReportEventAndSourceWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) ReportEventExW(ctx context.Context, in *ReportEventExWRequest, opts ...dcerpc.CallOption) (*ReportEventExWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReportEventExWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) ReportEventExA(ctx context.Context, in *ReportEventExARequest, opts ...dcerpc.CallOption) (*ReportEventExAResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReportEventExAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventlogClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEventlogClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewEventlogClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EventlogClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EventlogSyntaxV0_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultEventlogClient{cc: cc}, nil
}

// xxx_ClearEventLogWOperation structure represents the ElfrClearELFW operation
type xxx_ClearEventLogWOperation struct {
	Log            *Handle             `idl:"name:LogHandle" json:"log"`
	BackupFileName *dtyp.UnicodeString `idl:"name:BackupFileName;pointer:unique" json:"backup_file_name"`
	Return         int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearEventLogWOperation) OpNum() int { return 0 }

func (o *xxx_ClearEventLogWOperation) OpName() string { return "/eventlog/v0/ElfrClearELFW" }

func (o *xxx_ClearEventLogWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearEventLogWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// BackupFileName {in} (1:{pointer=unique, alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.BackupFileName != nil {
			_ptr_BackupFileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BackupFileName != nil {
					if err := o.BackupFileName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BackupFileName, _ptr_BackupFileName); err != nil {
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

func (o *xxx_ClearEventLogWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// BackupFileName {in} (1:{pointer=unique, alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		_ptr_BackupFileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BackupFileName == nil {
				o.BackupFileName = &dtyp.UnicodeString{}
			}
			if err := o.BackupFileName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_BackupFileName := func(ptr interface{}) { o.BackupFileName = *ptr.(**dtyp.UnicodeString) }
		if err := w.ReadPointer(&o.BackupFileName, _s_BackupFileName, _ptr_BackupFileName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearEventLogWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearEventLogWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearEventLogWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ClearEventLogWRequest structure represents the ElfrClearELFW operation request
type ClearEventLogWRequest struct {
	// LogHandle: Handle to the event log to be cleared. This parameter is a server context
	// handle, as specified in section 2.2.6. This handle MUST NOT be one obtained via the
	// ElfrOpenBELA (section 3.1.4.2) method or the ElfrOpenBELW (section 3.1.4.1) method.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// BackupFileName: Provides a Unicode string (as specified in section 2.2.11) that points
	// to an NT Object Path of a file in which a current copy of the event log is to be
	// placed. If this is NULL or empty, no backup is to be created. The path is relative
	// to the server rather than the client.
	BackupFileName *dtyp.UnicodeString `idl:"name:BackupFileName;pointer:unique" json:"backup_file_name"`
}

func (o *ClearEventLogWRequest) xxx_ToOp(ctx context.Context, op *xxx_ClearEventLogWOperation) *xxx_ClearEventLogWOperation {
	if op == nil {
		op = &xxx_ClearEventLogWOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.BackupFileName = o.BackupFileName
	return op
}

func (o *ClearEventLogWRequest) xxx_FromOp(ctx context.Context, op *xxx_ClearEventLogWOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.BackupFileName = op.BackupFileName
}
func (o *ClearEventLogWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ClearEventLogWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearEventLogWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearEventLogWResponse structure represents the ElfrClearELFW operation response
type ClearEventLogWResponse struct {
	// Return: The ElfrClearELFW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ClearEventLogWResponse) xxx_ToOp(ctx context.Context, op *xxx_ClearEventLogWOperation) *xxx_ClearEventLogWOperation {
	if op == nil {
		op = &xxx_ClearEventLogWOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ClearEventLogWResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearEventLogWOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ClearEventLogWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ClearEventLogWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearEventLogWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupEventLogWOperation structure represents the ElfrBackupELFW operation
type xxx_BackupEventLogWOperation struct {
	Log            *Handle             `idl:"name:LogHandle" json:"log"`
	BackupFileName *dtyp.UnicodeString `idl:"name:BackupFileName" json:"backup_file_name"`
	Return         int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupEventLogWOperation) OpNum() int { return 1 }

func (o *xxx_BackupEventLogWOperation) OpName() string { return "/eventlog/v0/ElfrBackupELFW" }

func (o *xxx_BackupEventLogWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEventLogWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// BackupFileName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.BackupFileName != nil {
			if err := o.BackupFileName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEventLogWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// BackupFileName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.BackupFileName == nil {
			o.BackupFileName = &dtyp.UnicodeString{}
		}
		if err := o.BackupFileName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEventLogWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEventLogWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEventLogWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupEventLogWRequest structure represents the ElfrBackupELFW operation request
type BackupEventLogWRequest struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6. This handle MUST NOT be obtained via the ElfrOpenBELA
	// (section 3.1.4.2) method or the ElfrOpenBELW (section 3.1.4.1) method.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// BackupFileName: Provides a Unicode string (as specified in section 2.2.11) that points
	// to an NT Object Path of a file, (as specified in section 2.2.4.1), in which a current
	// copy of the event log is to be placed. This MUST NOT be NULL or empty. The path is
	// evalutated relative to the server.
	BackupFileName *dtyp.UnicodeString `idl:"name:BackupFileName" json:"backup_file_name"`
}

func (o *BackupEventLogWRequest) xxx_ToOp(ctx context.Context, op *xxx_BackupEventLogWOperation) *xxx_BackupEventLogWOperation {
	if op == nil {
		op = &xxx_BackupEventLogWOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.BackupFileName = o.BackupFileName
	return op
}

func (o *BackupEventLogWRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupEventLogWOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.BackupFileName = op.BackupFileName
}
func (o *BackupEventLogWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BackupEventLogWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupEventLogWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupEventLogWResponse structure represents the ElfrBackupELFW operation response
type BackupEventLogWResponse struct {
	// Return: The ElfrBackupELFW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupEventLogWResponse) xxx_ToOp(ctx context.Context, op *xxx_BackupEventLogWOperation) *xxx_BackupEventLogWOperation {
	if op == nil {
		op = &xxx_BackupEventLogWOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BackupEventLogWResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupEventLogWOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BackupEventLogWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BackupEventLogWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupEventLogWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseEventLogOperation structure represents the ElfrCloseEL operation
type xxx_CloseEventLogOperation struct {
	Log    *Handle `idl:"name:LogHandle" json:"log"`
	Return int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseEventLogOperation) OpNum() int { return 2 }

func (o *xxx_CloseEventLogOperation) OpName() string { return "/eventlog/v0/ElfrCloseEL" }

func (o *xxx_CloseEventLogOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseEventLogOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseEventLogOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseEventLogOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseEventLogOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// LogHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseEventLogOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseEventLogRequest structure represents the ElfrCloseEL operation request
type CloseEventLogRequest struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6.
	Log *Handle `idl:"name:LogHandle" json:"log"`
}

func (o *CloseEventLogRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseEventLogOperation) *xxx_CloseEventLogOperation {
	if op == nil {
		op = &xxx_CloseEventLogOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	return op
}

func (o *CloseEventLogRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseEventLogOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
}
func (o *CloseEventLogRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseEventLogRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseEventLogOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseEventLogResponse structure represents the ElfrCloseEL operation response
type CloseEventLogResponse struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// Return: The ElfrCloseEL return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseEventLogResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseEventLogOperation) *xxx_CloseEventLogOperation {
	if op == nil {
		op = &xxx_CloseEventLogOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.Return = o.Return
	return op
}

func (o *CloseEventLogResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseEventLogOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.Return = op.Return
}
func (o *CloseEventLogResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseEventLogResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseEventLogOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeregisterEventSourceOperation structure represents the ElfrDeregisterEventSource operation
type xxx_DeregisterEventSourceOperation struct {
	Log    *Handle `idl:"name:LogHandle" json:"log"`
	Return int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_DeregisterEventSourceOperation) OpNum() int { return 3 }

func (o *xxx_DeregisterEventSourceOperation) OpName() string {
	return "/eventlog/v0/ElfrDeregisterEventSource"
}

func (o *xxx_DeregisterEventSourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeregisterEventSourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeregisterEventSourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeregisterEventSourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeregisterEventSourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// LogHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeregisterEventSourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeregisterEventSourceRequest structure represents the ElfrDeregisterEventSource operation request
type DeregisterEventSourceRequest struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6.
	Log *Handle `idl:"name:LogHandle" json:"log"`
}

func (o *DeregisterEventSourceRequest) xxx_ToOp(ctx context.Context, op *xxx_DeregisterEventSourceOperation) *xxx_DeregisterEventSourceOperation {
	if op == nil {
		op = &xxx_DeregisterEventSourceOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	return op
}

func (o *DeregisterEventSourceRequest) xxx_FromOp(ctx context.Context, op *xxx_DeregisterEventSourceOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
}
func (o *DeregisterEventSourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeregisterEventSourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeregisterEventSourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeregisterEventSourceResponse structure represents the ElfrDeregisterEventSource operation response
type DeregisterEventSourceResponse struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// Return: The ElfrDeregisterEventSource return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeregisterEventSourceResponse) xxx_ToOp(ctx context.Context, op *xxx_DeregisterEventSourceOperation) *xxx_DeregisterEventSourceOperation {
	if op == nil {
		op = &xxx_DeregisterEventSourceOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.Return = o.Return
	return op
}

func (o *DeregisterEventSourceResponse) xxx_FromOp(ctx context.Context, op *xxx_DeregisterEventSourceOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.Return = op.Return
}
func (o *DeregisterEventSourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeregisterEventSourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeregisterEventSourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NumberOfRecordsOperation structure represents the ElfrNumberOfRecords operation
type xxx_NumberOfRecordsOperation struct {
	Log             *Handle `idl:"name:LogHandle" json:"log"`
	NumberOfRecords uint32  `idl:"name:NumberOfRecords" json:"number_of_records"`
	Return          int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_NumberOfRecordsOperation) OpNum() int { return 4 }

func (o *xxx_NumberOfRecordsOperation) OpName() string { return "/eventlog/v0/ElfrNumberOfRecords" }

func (o *xxx_NumberOfRecordsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NumberOfRecordsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_NumberOfRecordsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NumberOfRecordsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NumberOfRecordsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// NumberOfRecords {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.NumberOfRecords); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NumberOfRecordsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// NumberOfRecords {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.NumberOfRecords); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NumberOfRecordsRequest structure represents the ElfrNumberOfRecords operation request
type NumberOfRecordsRequest struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6.
	Log *Handle `idl:"name:LogHandle" json:"log"`
}

func (o *NumberOfRecordsRequest) xxx_ToOp(ctx context.Context, op *xxx_NumberOfRecordsOperation) *xxx_NumberOfRecordsOperation {
	if op == nil {
		op = &xxx_NumberOfRecordsOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	return op
}

func (o *NumberOfRecordsRequest) xxx_FromOp(ctx context.Context, op *xxx_NumberOfRecordsOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
}
func (o *NumberOfRecordsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *NumberOfRecordsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NumberOfRecordsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NumberOfRecordsResponse structure represents the ElfrNumberOfRecords operation response
type NumberOfRecordsResponse struct {
	// NumberOfRecords: Total number of records in the specified event log.
	NumberOfRecords uint32 `idl:"name:NumberOfRecords" json:"number_of_records"`
	// Return: The ElfrNumberOfRecords return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *NumberOfRecordsResponse) xxx_ToOp(ctx context.Context, op *xxx_NumberOfRecordsOperation) *xxx_NumberOfRecordsOperation {
	if op == nil {
		op = &xxx_NumberOfRecordsOperation{}
	}
	if o == nil {
		return op
	}
	op.NumberOfRecords = o.NumberOfRecords
	op.Return = o.Return
	return op
}

func (o *NumberOfRecordsResponse) xxx_FromOp(ctx context.Context, op *xxx_NumberOfRecordsOperation) {
	if o == nil {
		return
	}
	o.NumberOfRecords = op.NumberOfRecords
	o.Return = op.Return
}
func (o *NumberOfRecordsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *NumberOfRecordsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NumberOfRecordsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OldestRecordOperation structure represents the ElfrOldestRecord operation
type xxx_OldestRecordOperation struct {
	Log                *Handle `idl:"name:LogHandle" json:"log"`
	OldestRecordNumber uint32  `idl:"name:OldestRecordNumber" json:"oldest_record_number"`
	Return             int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_OldestRecordOperation) OpNum() int { return 5 }

func (o *xxx_OldestRecordOperation) OpName() string { return "/eventlog/v0/ElfrOldestRecord" }

func (o *xxx_OldestRecordOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OldestRecordOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OldestRecordOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OldestRecordOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OldestRecordOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// OldestRecordNumber {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.OldestRecordNumber); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OldestRecordOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// OldestRecordNumber {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.OldestRecordNumber); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OldestRecordRequest structure represents the ElfrOldestRecord operation request
type OldestRecordRequest struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6.
	Log *Handle `idl:"name:LogHandle" json:"log"`
}

func (o *OldestRecordRequest) xxx_ToOp(ctx context.Context, op *xxx_OldestRecordOperation) *xxx_OldestRecordOperation {
	if op == nil {
		op = &xxx_OldestRecordOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	return op
}

func (o *OldestRecordRequest) xxx_FromOp(ctx context.Context, op *xxx_OldestRecordOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
}
func (o *OldestRecordRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OldestRecordRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OldestRecordOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OldestRecordResponse structure represents the ElfrOldestRecord operation response
type OldestRecordResponse struct {
	// OldestRecordNumber: The number of the oldest record in the specified event log. The
	// chronology is based on the time that records are written (not the record generation
	// time specified by the event source).
	OldestRecordNumber uint32 `idl:"name:OldestRecordNumber" json:"oldest_record_number"`
	// Return: The ElfrOldestRecord return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OldestRecordResponse) xxx_ToOp(ctx context.Context, op *xxx_OldestRecordOperation) *xxx_OldestRecordOperation {
	if op == nil {
		op = &xxx_OldestRecordOperation{}
	}
	if o == nil {
		return op
	}
	op.OldestRecordNumber = o.OldestRecordNumber
	op.Return = o.Return
	return op
}

func (o *OldestRecordResponse) xxx_FromOp(ctx context.Context, op *xxx_OldestRecordOperation) {
	if o == nil {
		return
	}
	o.OldestRecordNumber = op.OldestRecordNumber
	o.Return = op.Return
}
func (o *OldestRecordResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OldestRecordResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OldestRecordOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ChangeNotifyOperation structure represents the ElfrChangeNotify operation
type xxx_ChangeNotifyOperation struct {
	Log      *Handle   `idl:"name:LogHandle" json:"log"`
	ClientID *ClientID `idl:"name:ClientId" json:"client_id"`
	Event    uint32    `idl:"name:Event" json:"event"`
	Return   int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_ChangeNotifyOperation) OpNum() int { return 6 }

func (o *xxx_ChangeNotifyOperation) OpName() string { return "/eventlog/v0/ElfrChangeNotify" }

func (o *xxx_ChangeNotifyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeNotifyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ClientId {in} (1:{alias=RPC_CLIENT_ID}(struct))
	{
		if o.ClientID != nil {
			if err := o.ClientID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ClientID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Event {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Event); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeNotifyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ClientId {in} (1:{alias=RPC_CLIENT_ID}(struct))
	{
		if o.ClientID == nil {
			o.ClientID = &ClientID{}
		}
		if err := o.ClientID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Event {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Event); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeNotifyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeNotifyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeNotifyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ChangeNotifyRequest structure represents the ElfrChangeNotify operation request
type ChangeNotifyRequest struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// ClientId: Ignored when the method is called remotely.
	ClientID *ClientID `idl:"name:ClientId" json:"client_id"`
	// Event: Ignored when the method is called remotely.
	Event uint32 `idl:"name:Event" json:"event"`
}

func (o *ChangeNotifyRequest) xxx_ToOp(ctx context.Context, op *xxx_ChangeNotifyOperation) *xxx_ChangeNotifyOperation {
	if op == nil {
		op = &xxx_ChangeNotifyOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.ClientID = o.ClientID
	op.Event = o.Event
	return op
}

func (o *ChangeNotifyRequest) xxx_FromOp(ctx context.Context, op *xxx_ChangeNotifyOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.ClientID = op.ClientID
	o.Event = op.Event
}
func (o *ChangeNotifyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ChangeNotifyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangeNotifyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ChangeNotifyResponse structure represents the ElfrChangeNotify operation response
type ChangeNotifyResponse struct {
	// Return: The ElfrChangeNotify return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ChangeNotifyResponse) xxx_ToOp(ctx context.Context, op *xxx_ChangeNotifyOperation) *xxx_ChangeNotifyOperation {
	if op == nil {
		op = &xxx_ChangeNotifyOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ChangeNotifyResponse) xxx_FromOp(ctx context.Context, op *xxx_ChangeNotifyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ChangeNotifyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ChangeNotifyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangeNotifyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenEventLogWOperation structure represents the ElfrOpenELW operation
type xxx_OpenEventLogWOperation struct {
	UNCServerName string              `idl:"name:UNCServerName" json:"unc_server_name"`
	ModuleName    *dtyp.UnicodeString `idl:"name:ModuleName" json:"module_name"`
	RegModuleName *dtyp.UnicodeString `idl:"name:RegModuleName" json:"reg_module_name"`
	MajorVersion  uint32              `idl:"name:MajorVersion" json:"major_version"`
	MinorVersion  uint32              `idl:"name:MinorVersion" json:"minor_version"`
	Log           *Handle             `idl:"name:LogHandle" json:"log"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenEventLogWOperation) OpNum() int { return 7 }

func (o *xxx_OpenEventLogWOperation) OpName() string { return "/eventlog/v0/ElfrOpenELW" }

func (o *xxx_OpenEventLogWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEventLogWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UNCServerName {in} (1:{handle, pointer=unique, alias=EVENTLOG_HANDLE_W}*(1)[dim:0,string](wchar))
	{
		if o.UNCServerName != "" {
			_ptr_UNCServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.UNCServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.UNCServerName, _ptr_UNCServerName); err != nil {
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
	// ModuleName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ModuleName != nil {
			if err := o.ModuleName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// RegModuleName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.RegModuleName != nil {
			if err := o.RegModuleName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.MinorVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEventLogWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UNCServerName {in} (1:{handle, pointer=unique, alias=EVENTLOG_HANDLE_W}*(1)[dim:0,string](wchar))
	{
		_ptr_UNCServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.UNCServerName); err != nil {
				return err
			}
			return nil
		})
		_s_UNCServerName := func(ptr interface{}) { o.UNCServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.UNCServerName, _s_UNCServerName, _ptr_UNCServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ModuleName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ModuleName == nil {
			o.ModuleName = &dtyp.UnicodeString{}
		}
		if err := o.ModuleName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RegModuleName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.RegModuleName == nil {
			o.RegModuleName = &dtyp.UnicodeString{}
		}
		if err := o.RegModuleName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MinorVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEventLogWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEventLogWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// LogHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEventLogWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// LogHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenEventLogWRequest structure represents the ElfrOpenELW operation request
type OpenEventLogWRequest struct {
	// UNCServerName: A server interface handle. A pointer to a Unicode string specifying
	// the server, as specified in section 2.2.7. The client MUST map this string to an
	// RPC binding handle, and the server MUST ignore this argument, as specified in [C706]
	// sections 4.3.5 and 5.1.5.2.
	UNCServerName string `idl:"name:UNCServerName" json:"unc_server_name"`
	// ModuleName: Specifies the event log name, as defined in section 1.8.2 and specified
	// in section 2.2.11, for which a handle is needed.
	ModuleName *dtyp.UnicodeString `idl:"name:ModuleName" json:"module_name"`
	// RegModuleName: This parameter MUST be ignored by the server. Clients MUST specify
	// an empty string.
	RegModuleName *dtyp.UnicodeString `idl:"name:RegModuleName" json:"reg_module_name"`
	// MajorVersion: Major version of the client. This value MUST be set to 1.
	MajorVersion uint32 `idl:"name:MajorVersion" json:"major_version"`
	// MinorVersion: Minor version of the client. This value MUST be set to 1.
	MinorVersion uint32 `idl:"name:MinorVersion" json:"minor_version"`
}

func (o *OpenEventLogWRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenEventLogWOperation) *xxx_OpenEventLogWOperation {
	if op == nil {
		op = &xxx_OpenEventLogWOperation{}
	}
	if o == nil {
		return op
	}
	op.UNCServerName = o.UNCServerName
	op.ModuleName = o.ModuleName
	op.RegModuleName = o.RegModuleName
	op.MajorVersion = o.MajorVersion
	op.MinorVersion = o.MinorVersion
	return op
}

func (o *OpenEventLogWRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenEventLogWOperation) {
	if o == nil {
		return
	}
	o.UNCServerName = op.UNCServerName
	o.ModuleName = op.ModuleName
	o.RegModuleName = op.RegModuleName
	o.MajorVersion = op.MajorVersion
	o.MinorVersion = op.MinorVersion
}
func (o *OpenEventLogWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenEventLogWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenEventLogWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenEventLogWResponse structure represents the ElfrOpenELW operation response
type OpenEventLogWResponse struct {
	// LogHandle: Pointer to an event log handle. This parameter is a server context handle,
	// as specified in section 2.2.6. This handle MUST be closed by using the ElfrCloseEL
	// (section 3.1.4.21) method once the handle is no longer needed. In the case when the
	// client cannot call the ElfrCloseEL function such as the abnormal termination of the
	// client, this context handle will be revoked by the server so that there will not
	// be any resource leaks.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// Return: The ElfrOpenELW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenEventLogWResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenEventLogWOperation) *xxx_OpenEventLogWOperation {
	if op == nil {
		op = &xxx_OpenEventLogWOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.Return = o.Return
	return op
}

func (o *OpenEventLogWResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenEventLogWOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.Return = op.Return
}
func (o *OpenEventLogWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenEventLogWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenEventLogWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterEventSourceWOperation structure represents the ElfrRegisterEventSourceW operation
type xxx_RegisterEventSourceWOperation struct {
	UNCServerName string              `idl:"name:UNCServerName" json:"unc_server_name"`
	ModuleName    *dtyp.UnicodeString `idl:"name:ModuleName" json:"module_name"`
	RegModuleName *dtyp.UnicodeString `idl:"name:RegModuleName" json:"reg_module_name"`
	MajorVersion  uint32              `idl:"name:MajorVersion" json:"major_version"`
	MinorVersion  uint32              `idl:"name:MinorVersion" json:"minor_version"`
	Log           *Handle             `idl:"name:LogHandle" json:"log"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterEventSourceWOperation) OpNum() int { return 8 }

func (o *xxx_RegisterEventSourceWOperation) OpName() string {
	return "/eventlog/v0/ElfrRegisterEventSourceW"
}

func (o *xxx_RegisterEventSourceWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterEventSourceWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UNCServerName {in} (1:{handle, pointer=unique, alias=EVENTLOG_HANDLE_W}*(1)[dim:0,string](wchar))
	{
		if o.UNCServerName != "" {
			_ptr_UNCServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.UNCServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.UNCServerName, _ptr_UNCServerName); err != nil {
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
	// ModuleName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ModuleName != nil {
			if err := o.ModuleName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// RegModuleName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.RegModuleName != nil {
			if err := o.RegModuleName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.MinorVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterEventSourceWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UNCServerName {in} (1:{handle, pointer=unique, alias=EVENTLOG_HANDLE_W}*(1)[dim:0,string](wchar))
	{
		_ptr_UNCServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.UNCServerName); err != nil {
				return err
			}
			return nil
		})
		_s_UNCServerName := func(ptr interface{}) { o.UNCServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.UNCServerName, _s_UNCServerName, _ptr_UNCServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ModuleName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ModuleName == nil {
			o.ModuleName = &dtyp.UnicodeString{}
		}
		if err := o.ModuleName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RegModuleName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.RegModuleName == nil {
			o.RegModuleName = &dtyp.UnicodeString{}
		}
		if err := o.RegModuleName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MinorVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterEventSourceWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterEventSourceWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// LogHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterEventSourceWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// LogHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RegisterEventSourceWRequest structure represents the ElfrRegisterEventSourceW operation request
type RegisterEventSourceWRequest struct {
	// UNCServerName: A server interface handle. A pointer to a Unicode (as specified in
	// [MS-DTYP]) string specifying the server, as specified in section 2.2.7. The client
	// MUST map this string to an RPC binding handle, and the server MUST ignore this argument,
	// as specified in [C706] sections 4.3.5 and 5.1.5.2.
	UNCServerName string `idl:"name:UNCServerName" json:"unc_server_name"`
	// ModuleName: Specifies the event source, as defined in section 1.8.3 and specified
	// in section 2.2.11, for which a handle is needed.
	ModuleName *dtyp.UnicodeString `idl:"name:ModuleName" json:"module_name"`
	// RegModuleName: This parameter MUST be ignored by the server. Clients MUST specify
	// an empty string.
	RegModuleName *dtyp.UnicodeString `idl:"name:RegModuleName" json:"reg_module_name"`
	// MajorVersion: Major version of the client. This value MUST be set to 1.
	MajorVersion uint32 `idl:"name:MajorVersion" json:"major_version"`
	// MinorVersion: Minor version of the client. This value MUST be set to 1.
	MinorVersion uint32 `idl:"name:MinorVersion" json:"minor_version"`
}

func (o *RegisterEventSourceWRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterEventSourceWOperation) *xxx_RegisterEventSourceWOperation {
	if op == nil {
		op = &xxx_RegisterEventSourceWOperation{}
	}
	if o == nil {
		return op
	}
	op.UNCServerName = o.UNCServerName
	op.ModuleName = o.ModuleName
	op.RegModuleName = o.RegModuleName
	op.MajorVersion = o.MajorVersion
	op.MinorVersion = o.MinorVersion
	return op
}

func (o *RegisterEventSourceWRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterEventSourceWOperation) {
	if o == nil {
		return
	}
	o.UNCServerName = op.UNCServerName
	o.ModuleName = op.ModuleName
	o.RegModuleName = op.RegModuleName
	o.MajorVersion = op.MajorVersion
	o.MinorVersion = op.MinorVersion
}
func (o *RegisterEventSourceWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterEventSourceWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterEventSourceWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterEventSourceWResponse structure represents the ElfrRegisterEventSourceW operation response
type RegisterEventSourceWResponse struct {
	// LogHandle: Pointer to an event log handle. This parameter is a server context handle,
	// as specified in section 2.2.6.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// Return: The ElfrRegisterEventSourceW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterEventSourceWResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterEventSourceWOperation) *xxx_RegisterEventSourceWOperation {
	if op == nil {
		op = &xxx_RegisterEventSourceWOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.Return = o.Return
	return op
}

func (o *RegisterEventSourceWResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterEventSourceWOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.Return = op.Return
}
func (o *RegisterEventSourceWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterEventSourceWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterEventSourceWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenBackupEventLogWOperation structure represents the ElfrOpenBELW operation
type xxx_OpenBackupEventLogWOperation struct {
	UNCServerName  string              `idl:"name:UNCServerName" json:"unc_server_name"`
	BackupFileName *dtyp.UnicodeString `idl:"name:BackupFileName" json:"backup_file_name"`
	MajorVersion   uint32              `idl:"name:MajorVersion" json:"major_version"`
	MinorVersion   uint32              `idl:"name:MinorVersion" json:"minor_version"`
	Log            *Handle             `idl:"name:LogHandle" json:"log"`
	Return         int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenBackupEventLogWOperation) OpNum() int { return 9 }

func (o *xxx_OpenBackupEventLogWOperation) OpName() string { return "/eventlog/v0/ElfrOpenBELW" }

func (o *xxx_OpenBackupEventLogWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenBackupEventLogWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UNCServerName {in} (1:{handle, pointer=unique, alias=EVENTLOG_HANDLE_W}*(1)[dim:0,string](wchar))
	{
		if o.UNCServerName != "" {
			_ptr_UNCServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.UNCServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.UNCServerName, _ptr_UNCServerName); err != nil {
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
	// BackupFileName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.BackupFileName != nil {
			if err := o.BackupFileName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.MinorVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenBackupEventLogWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UNCServerName {in} (1:{handle, pointer=unique, alias=EVENTLOG_HANDLE_W}*(1)[dim:0,string](wchar))
	{
		_ptr_UNCServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.UNCServerName); err != nil {
				return err
			}
			return nil
		})
		_s_UNCServerName := func(ptr interface{}) { o.UNCServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.UNCServerName, _s_UNCServerName, _ptr_UNCServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BackupFileName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.BackupFileName == nil {
			o.BackupFileName = &dtyp.UnicodeString{}
		}
		if err := o.BackupFileName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MinorVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenBackupEventLogWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenBackupEventLogWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// LogHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenBackupEventLogWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// LogHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenBackupEventLogWRequest structure represents the ElfrOpenBELW operation request
type OpenBackupEventLogWRequest struct {
	// UNCServerName: A server interface handle. A pointer to a Unicode string specifying
	// the server, as specified in section 2.2.7. The client MUST map this string to an
	// RPC binding handle, and the server MUST ignore this argument. See [C706] sections
	// 4.3.5 and 5.1.5.2.
	UNCServerName string `idl:"name:UNCServerName" json:"unc_server_name"`
	// BackupFileName: Provides a Unicode string (as specified in section 2.2.11) that points
	// to an NT Object Path of the file where the backup event log is located, as specified
	// in section 2.2.4.1.
	BackupFileName *dtyp.UnicodeString `idl:"name:BackupFileName" json:"backup_file_name"`
	// MajorVersion: Major version of the client. This value MUST be set to 1.
	MajorVersion uint32 `idl:"name:MajorVersion" json:"major_version"`
	// MinorVersion: Minor version of the client. This value MUST be set to 1.
	MinorVersion uint32 `idl:"name:MinorVersion" json:"minor_version"`
}

func (o *OpenBackupEventLogWRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenBackupEventLogWOperation) *xxx_OpenBackupEventLogWOperation {
	if op == nil {
		op = &xxx_OpenBackupEventLogWOperation{}
	}
	if o == nil {
		return op
	}
	op.UNCServerName = o.UNCServerName
	op.BackupFileName = o.BackupFileName
	op.MajorVersion = o.MajorVersion
	op.MinorVersion = o.MinorVersion
	return op
}

func (o *OpenBackupEventLogWRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenBackupEventLogWOperation) {
	if o == nil {
		return
	}
	o.UNCServerName = op.UNCServerName
	o.BackupFileName = op.BackupFileName
	o.MajorVersion = op.MajorVersion
	o.MinorVersion = op.MinorVersion
}
func (o *OpenBackupEventLogWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenBackupEventLogWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenBackupEventLogWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenBackupEventLogWResponse structure represents the ElfrOpenBELW operation response
type OpenBackupEventLogWResponse struct {
	// LogHandle: Pointer to an event log handle. This parameter is a server context handle,
	// as specified in section 2.2.6. This handle MUST be closed using the ElfrCloseEL (Opnum
	// 2) (section 3.1.4.21) method once the handle is no longer needed. In the case when
	// the client cannot call the ElfrCloseEL function, such as the abnormal termination
	// of the client, this context handle will be revoked by the server so that there will
	// not be any resource leaks. The processing rule to revoke a context handle that has
	// been terminated abnormally is defined in [MS-RPCE] section 3.3.3.2.1.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// Return: The ElfrOpenBELW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenBackupEventLogWResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenBackupEventLogWOperation) *xxx_OpenBackupEventLogWOperation {
	if op == nil {
		op = &xxx_OpenBackupEventLogWOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.Return = o.Return
	return op
}

func (o *OpenBackupEventLogWResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenBackupEventLogWOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.Return = op.Return
}
func (o *OpenBackupEventLogWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenBackupEventLogWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenBackupEventLogWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReadEventLogWOperation structure represents the ElfrReadELW operation
type xxx_ReadEventLogWOperation struct {
	Log                    *Handle `idl:"name:LogHandle" json:"log"`
	ReadFlags              uint32  `idl:"name:ReadFlags" json:"read_flags"`
	RecordOffset           uint32  `idl:"name:RecordOffset" json:"record_offset"`
	NumberOfBytesToRead    uint32  `idl:"name:NumberOfBytesToRead" json:"number_of_bytes_to_read"`
	Buffer                 []byte  `idl:"name:Buffer;size_is:(NumberOfBytesToRead)" json:"buffer"`
	NumberOfBytesRead      uint32  `idl:"name:NumberOfBytesRead" json:"number_of_bytes_read"`
	MinNumberOfBytesNeeded uint32  `idl:"name:MinNumberOfBytesNeeded" json:"min_number_of_bytes_needed"`
	Return                 int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_ReadEventLogWOperation) OpNum() int { return 10 }

func (o *xxx_ReadEventLogWOperation) OpName() string { return "/eventlog/v0/ElfrReadELW" }

func (o *xxx_ReadEventLogWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadEventLogWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ReadFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.ReadFlags); err != nil {
			return err
		}
	}
	// RecordOffset {in} (1:(uint32))
	{
		if err := w.WriteData(o.RecordOffset); err != nil {
			return err
		}
	}
	// NumberOfBytesToRead {in} (1:{range=(0,524287), alias=RULONG}(uint32))
	{
		if err := w.WriteData(o.NumberOfBytesToRead); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadEventLogWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ReadFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ReadFlags); err != nil {
			return err
		}
	}
	// RecordOffset {in} (1:(uint32))
	{
		if err := w.ReadData(&o.RecordOffset); err != nil {
			return err
		}
	}
	// NumberOfBytesToRead {in} (1:{range=(0,524287), alias=RULONG}(uint32))
	{
		if err := w.ReadData(&o.NumberOfBytesToRead); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadEventLogWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadEventLogWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=NumberOfBytesToRead](uchar))
	{
		dimSize1 := uint64(o.NumberOfBytesToRead)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Buffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Buffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// NumberOfBytesRead {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.NumberOfBytesRead); err != nil {
			return err
		}
	}
	// MinNumberOfBytesNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MinNumberOfBytesNeeded); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadEventLogWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=NumberOfBytesToRead](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
	}
	// NumberOfBytesRead {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.NumberOfBytesRead); err != nil {
			return err
		}
	}
	// MinNumberOfBytesNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MinNumberOfBytesNeeded); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReadEventLogWRequest structure represents the ElfrReadELW operation request
type ReadEventLogWRequest struct {
	// LogHandle: Handle to an event log to read. This parameter is a server context handle,
	// as specified in section 2.2.6.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// ReadFlags: The caller MUST specify whether the read is to start at a specific record
	// or is to proceed from the last record read. The value MUST include one and only one
	// of the following flags.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EVENTLOG_SEQUENTIAL_READ 0x00000001 | Read operation proceeds sequentially from the last call to the ElfrReadELW       |
	//	|                                     | (section 3.1.4.7) method or the ElfrReadELA (section 3.1.4.8) method, using this |
	//	|                                     | handle. This flag MUST NOT be used with EVENTLOG_SEEK_READ.                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EVENTLOG_SEEK_READ 0x00000002       | Read operation proceeds from the record specified by the RecordOffset parameter. |
	//	|                                     | This flag MUST NOT be used with EVENTLOG_SEQUENTIAL_READ.                        |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// Because the method reads as many records as can fit in the buffer, the caller MUST
	// also set one and only one of the following flags to indicate the direction for successive
	// read operations.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| EVENTLOG_FORWARDS_READ 0x00000004  | Log is read in chronological order. This flag MUST NOT be used with              |
	//	|                                    | EVENTLOG_BACKWARDS_READ.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| EVENTLOG_BACKWARDS_READ 0x00000008 | Log is read in reverse chronological order. This flag MUST NOT be used with      |
	//	|                                    | EVENTLOG_FORWARDS_READ.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	ReadFlags uint32 `idl:"name:ReadFlags" json:"read_flags"`
	// RecordOffset: Log entry record number from which the read operation starts (this
	// is not a byte offset but a number). This parameter MUST be ignored unless the EVENTLOG_SEEK_READ
	// bit is set in the ReadFlags parameter.
	RecordOffset uint32 `idl:"name:RecordOffset" json:"record_offset"`
	// NumberOfBytesToRead: Size of the Buffer in bytes. This is the maximum amount of data
	// that can be read.
	NumberOfBytesToRead uint32 `idl:"name:NumberOfBytesToRead" json:"number_of_bytes_to_read"`
}

func (o *ReadEventLogWRequest) xxx_ToOp(ctx context.Context, op *xxx_ReadEventLogWOperation) *xxx_ReadEventLogWOperation {
	if op == nil {
		op = &xxx_ReadEventLogWOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.ReadFlags = o.ReadFlags
	op.RecordOffset = o.RecordOffset
	op.NumberOfBytesToRead = o.NumberOfBytesToRead
	return op
}

func (o *ReadEventLogWRequest) xxx_FromOp(ctx context.Context, op *xxx_ReadEventLogWOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.ReadFlags = op.ReadFlags
	o.RecordOffset = op.RecordOffset
	o.NumberOfBytesToRead = op.NumberOfBytesToRead
}
func (o *ReadEventLogWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReadEventLogWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReadEventLogWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReadEventLogWResponse structure represents the ElfrReadELW operation response
type ReadEventLogWResponse struct {
	// XXX: NumberOfBytesToRead is an implicit input depedency for output parameters
	NumberOfBytesToRead uint32 `idl:"name:NumberOfBytesToRead" json:"number_of_bytes_to_read"`

	// Buffer: The buffer in which to place data read from the event log.
	Buffer []byte `idl:"name:Buffer;size_is:(NumberOfBytesToRead)" json:"buffer"`
	// NumberOfBytesRead: Pointer to a variable that receives the number of bytes actually
	// read by the method.
	NumberOfBytesRead uint32 `idl:"name:NumberOfBytesRead" json:"number_of_bytes_read"`
	// MinNumberOfBytesNeeded: If the method fails because the buffer is too small to fit
	// even a single record, this MUST be set to the minimum number of bytes needed to fit
	// the next record. Otherwise, this MUST NOT be set, and MUST be ignored by the caller.
	MinNumberOfBytesNeeded uint32 `idl:"name:MinNumberOfBytesNeeded" json:"min_number_of_bytes_needed"`
	// Return: The ElfrReadELW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReadEventLogWResponse) xxx_ToOp(ctx context.Context, op *xxx_ReadEventLogWOperation) *xxx_ReadEventLogWOperation {
	if op == nil {
		op = &xxx_ReadEventLogWOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.NumberOfBytesToRead == uint32(0) {
		op.NumberOfBytesToRead = o.NumberOfBytesToRead
	}

	op.Buffer = o.Buffer
	op.NumberOfBytesRead = o.NumberOfBytesRead
	op.MinNumberOfBytesNeeded = o.MinNumberOfBytesNeeded
	op.Return = o.Return
	return op
}

func (o *ReadEventLogWResponse) xxx_FromOp(ctx context.Context, op *xxx_ReadEventLogWOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.NumberOfBytesToRead = op.NumberOfBytesToRead

	o.Buffer = op.Buffer
	o.NumberOfBytesRead = op.NumberOfBytesRead
	o.MinNumberOfBytesNeeded = op.MinNumberOfBytesNeeded
	o.Return = op.Return
}
func (o *ReadEventLogWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReadEventLogWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReadEventLogWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReportEventWOperation structure represents the ElfrReportEventW operation
type xxx_ReportEventWOperation struct {
	Log           *Handle               `idl:"name:LogHandle" json:"log"`
	Time          uint32                `idl:"name:Time" json:"time"`
	EventType     uint16                `idl:"name:EventType" json:"event_type"`
	EventCategory uint16                `idl:"name:EventCategory" json:"event_category"`
	EventID       uint32                `idl:"name:EventID" json:"event_id"`
	StringsLength uint16                `idl:"name:NumStrings" json:"strings_length"`
	DataSize      uint32                `idl:"name:DataSize" json:"data_size"`
	ComputerName  *dtyp.UnicodeString   `idl:"name:ComputerName" json:"computer_name"`
	UserSID       *dtyp.SID             `idl:"name:UserSID;pointer:unique" json:"user_sid"`
	Strings       []*dtyp.UnicodeString `idl:"name:Strings;size_is:(NumStrings);pointer:unique" json:"strings"`
	Data          []byte                `idl:"name:Data;size_is:(DataSize);pointer:unique" json:"data"`
	Flags         uint16                `idl:"name:Flags" json:"flags"`
	RecordNumber  uint32                `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	TimeWritten   uint32                `idl:"name:TimeWritten;pointer:unique" json:"time_written"`
	Return        int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_ReportEventWOperation) OpNum() int { return 11 }

func (o *xxx_ReportEventWOperation) OpName() string { return "/eventlog/v0/ElfrReportEventW" }

func (o *xxx_ReportEventWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Strings != nil && o.StringsLength == 0 {
		o.StringsLength = uint16(len(o.Strings))
	}
	if o.Data != nil && o.DataSize == 0 {
		o.DataSize = uint32(len(o.Data))
	}
	if o.StringsLength > uint16(256) {
		return fmt.Errorf("StringsLength is out of range")
	}
	if o.DataSize > uint32(61440) {
		return fmt.Errorf("DataSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Time {in} (1:(uint32))
	{
		if err := w.WriteData(o.Time); err != nil {
			return err
		}
	}
	// EventType {in} (1:(uint16))
	{
		if err := w.WriteData(o.EventType); err != nil {
			return err
		}
	}
	// EventCategory {in} (1:(uint16))
	{
		if err := w.WriteData(o.EventCategory); err != nil {
			return err
		}
	}
	// EventID {in} (1:(uint32))
	{
		if err := w.WriteData(o.EventID); err != nil {
			return err
		}
	}
	// NumStrings {in} (1:{range=(0,256)}(uint16))
	{
		if err := w.WriteData(o.StringsLength); err != nil {
			return err
		}
	}
	// DataSize {in} (1:{range=(0,61440)}(uint32))
	{
		if err := w.WriteData(o.DataSize); err != nil {
			return err
		}
	}
	// ComputerName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ComputerName != nil {
			if err := o.ComputerName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// UserSID {in} (1:{pointer=unique, alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.UserSID != nil {
			_ptr_UserSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserSID != nil {
					if err := o.UserSID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserSID, _ptr_UserSID); err != nil {
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
	// Strings {in} (1:{pointer=unique}[dim:0,size_is=NumStrings])(2:{alias=PRPC_UNICODE_STRING}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		dimSize1 := uint64(o.StringsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Strings {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Strings[i1] != nil {
				_ptr_Strings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.Strings[i1] != nil {
						if err := o.Strings[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Strings[i1], _ptr_Strings); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Strings); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Data {in} (1:{pointer=unique}*(1)[dim:0,size_is=DataSize](uchar))
	{
		if o.Data != nil || o.DataSize > 0 {
			_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DataSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Data {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Data[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data, _ptr_Data); err != nil {
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
	// Flags {in} (1:(uint16))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TimeWritten {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_TimeWritten := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.TimeWritten); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TimeWritten, _ptr_TimeWritten); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Time {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Time); err != nil {
			return err
		}
	}
	// EventType {in} (1:(uint16))
	{
		if err := w.ReadData(&o.EventType); err != nil {
			return err
		}
	}
	// EventCategory {in} (1:(uint16))
	{
		if err := w.ReadData(&o.EventCategory); err != nil {
			return err
		}
	}
	// EventID {in} (1:(uint32))
	{
		if err := w.ReadData(&o.EventID); err != nil {
			return err
		}
	}
	// NumStrings {in} (1:{range=(0,256)}(uint16))
	{
		if err := w.ReadData(&o.StringsLength); err != nil {
			return err
		}
	}
	// DataSize {in} (1:{range=(0,61440)}(uint32))
	{
		if err := w.ReadData(&o.DataSize); err != nil {
			return err
		}
	}
	// ComputerName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ComputerName == nil {
			o.ComputerName = &dtyp.UnicodeString{}
		}
		if err := o.ComputerName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UserSID {in} (1:{pointer=unique, alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		_ptr_UserSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserSID == nil {
				o.UserSID = &dtyp.SID{}
			}
			if err := o.UserSID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_UserSID := func(ptr interface{}) { o.UserSID = *ptr.(**dtyp.SID) }
		if err := w.ReadPointer(&o.UserSID, _s_UserSID, _ptr_UserSID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Strings {in} (1:{pointer=unique}[dim:0,size_is=NumStrings])(2:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Strings", sizeInfo[0])
		}
		o.Strings = make([]*dtyp.UnicodeString, sizeInfo[0])
		for i1 := range o.Strings {
			i1 := i1
			_ptr_Strings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Strings[i1] == nil {
					o.Strings[i1] = &dtyp.UnicodeString{}
				}
				if err := o.Strings[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Strings := func(ptr interface{}) { o.Strings[i1] = *ptr.(**dtyp.UnicodeString) }
			if err := w.ReadPointer(&o.Strings[i1], _s_Strings, _ptr_Strings); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Data {in} (1:{pointer=unique}*(1)[dim:0,size_is=DataSize](uchar))
	{
		_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
			}
			o.Data = make([]byte, sizeInfo[0])
			for i1 := range o.Data {
				i1 := i1
				if err := w.ReadData(&o.Data[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_Data := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Data, _s_Data, _ptr_Data); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Flags {in} (1:(uint16))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		_s_RecordNumber := func(ptr interface{}) { o.RecordNumber = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.RecordNumber, _s_RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TimeWritten {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_TimeWritten := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.TimeWritten); err != nil {
				return err
			}
			return nil
		})
		_s_TimeWritten := func(ptr interface{}) { o.TimeWritten = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.TimeWritten, _s_TimeWritten, _ptr_TimeWritten); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TimeWritten {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_TimeWritten := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.TimeWritten); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TimeWritten, _ptr_TimeWritten); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		_s_RecordNumber := func(ptr interface{}) { o.RecordNumber = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.RecordNumber, _s_RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TimeWritten {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_TimeWritten := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.TimeWritten); err != nil {
				return err
			}
			return nil
		})
		_s_TimeWritten := func(ptr interface{}) { o.TimeWritten = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.TimeWritten, _s_TimeWritten, _ptr_TimeWritten); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReportEventWRequest structure represents the ElfrReportEventW operation request
type ReportEventWRequest struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6. This handle MUST NOT be obtained via the ElfrOpenBELA
	// (section 3.1.4.2) method or the ElfrOpenBELW (section 3.1.4.1) method. A handle received
	// from either of those two methods will have the backup flag set, so the server checks
	// this flag before calling this method.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// Time: Time at which the event was generated by the event source (not the time at
	// which the event was logged). The time MUST be expressed as the number of seconds
	// since 00:00:00 on January 1, 1970 (UTC).
	Time uint32 `idl:"name:Time" json:"time"`
	// EventType: Type of the event, as specified in section 2.2.2.
	EventType uint16 `idl:"name:EventType" json:"event_type"`
	// EventCategory: Event category, as specified in section 1.8.5.
	EventCategory uint16 `idl:"name:EventCategory" json:"event_category"`
	// EventID: EventID, as specified in section 3.1.1.4.
	EventID uint32 `idl:"name:EventID" json:"event_id"`
	// NumStrings: Number of strings in the array pointed to by the Strings parameter. A
	// value of zero indicates that no strings are present.
	StringsLength uint16 `idl:"name:NumStrings" json:"strings_length"`
	// DataSize: Number of bytes of event-specific raw binary data to write to the log.
	// This binary data is passed in the Data parameter. If the DataSize parameter is zero,
	// event-specific data MUST NOT be present.
	DataSize uint32 `idl:"name:DataSize" json:"data_size"`
	// ComputerName: A string to assist in identifying the machine that generated the event.
	// In practice, the name of the computer. There are no character restrictions on this
	// field's content (for example, a FQDN can be used). The API is not intended to support
	// dynamically changing computer names. The ComputerName parameter SHOULD<27> be cached
	// the first time a client calls the API, and SHOULD use that name on subsequent calls
	// until the machine is rebooted.
	ComputerName *dtyp.UnicodeString `idl:"name:ComputerName" json:"computer_name"`
	// UserSID: Either NULL or a user SID. If this is NULL, the event is to have a zero
	// length UserSid field.
	UserSID *dtyp.SID `idl:"name:UserSID;pointer:unique" json:"user_sid"`
	// Strings: Specifies strings containing information specific to the event. This parameter
	// MUST be a valid pointer. If the NumStrings parameter is zero, this parameter MUST
	// be NULL. For example, an event relating to file deletion could use a string to specify
	// the path of the file being deleted.
	Strings []*dtyp.UnicodeString `idl:"name:Strings;size_is:(NumStrings);pointer:unique" json:"strings"`
	// Data: Pointer to the buffer that contains the event-specific binary data. This parameter
	// MUST be a valid pointer (or NULL), even if the DataSize parameter is zero.
	Data []byte `idl:"name:Data;size_is:(DataSize);pointer:unique" json:"data"`
	// Flags: Unused. MUST be set to zero when sent and MUST be ignored on receipt.
	Flags uint16 `idl:"name:Flags" json:"flags"`
	// RecordNumber: Unused. Can be set to any arbitrary value when sent, and any value
	// sent by the client MUST be ignored on receipt by the server.
	RecordNumber uint32 `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	// TimeWritten: Unused. Can be set to any arbitrary value when sent, and any value sent
	// by the client MUST be ignored on receipt by the server.
	TimeWritten uint32 `idl:"name:TimeWritten;pointer:unique" json:"time_written"`
}

func (o *ReportEventWRequest) xxx_ToOp(ctx context.Context, op *xxx_ReportEventWOperation) *xxx_ReportEventWOperation {
	if op == nil {
		op = &xxx_ReportEventWOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.Time = o.Time
	op.EventType = o.EventType
	op.EventCategory = o.EventCategory
	op.EventID = o.EventID
	op.StringsLength = o.StringsLength
	op.DataSize = o.DataSize
	op.ComputerName = o.ComputerName
	op.UserSID = o.UserSID
	op.Strings = o.Strings
	op.Data = o.Data
	op.Flags = o.Flags
	op.RecordNumber = o.RecordNumber
	op.TimeWritten = o.TimeWritten
	return op
}

func (o *ReportEventWRequest) xxx_FromOp(ctx context.Context, op *xxx_ReportEventWOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.Time = op.Time
	o.EventType = op.EventType
	o.EventCategory = op.EventCategory
	o.EventID = op.EventID
	o.StringsLength = op.StringsLength
	o.DataSize = op.DataSize
	o.ComputerName = op.ComputerName
	o.UserSID = op.UserSID
	o.Strings = op.Strings
	o.Data = op.Data
	o.Flags = op.Flags
	o.RecordNumber = op.RecordNumber
	o.TimeWritten = op.TimeWritten
}
func (o *ReportEventWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReportEventWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportEventWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReportEventWResponse structure represents the ElfrReportEventW operation response
type ReportEventWResponse struct {
	// RecordNumber: Unused. Can be set to any arbitrary value when sent, and any value
	// sent by the client MUST be ignored on receipt by the server.
	RecordNumber uint32 `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	// TimeWritten: Unused. Can be set to any arbitrary value when sent, and any value sent
	// by the client MUST be ignored on receipt by the server.
	TimeWritten uint32 `idl:"name:TimeWritten;pointer:unique" json:"time_written"`
	// Return: The ElfrReportEventW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReportEventWResponse) xxx_ToOp(ctx context.Context, op *xxx_ReportEventWOperation) *xxx_ReportEventWOperation {
	if op == nil {
		op = &xxx_ReportEventWOperation{}
	}
	if o == nil {
		return op
	}
	op.RecordNumber = o.RecordNumber
	op.TimeWritten = o.TimeWritten
	op.Return = o.Return
	return op
}

func (o *ReportEventWResponse) xxx_FromOp(ctx context.Context, op *xxx_ReportEventWOperation) {
	if o == nil {
		return
	}
	o.RecordNumber = op.RecordNumber
	o.TimeWritten = op.TimeWritten
	o.Return = op.Return
}
func (o *ReportEventWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReportEventWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportEventWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClearEventLogAOperation structure represents the ElfrClearELFA operation
type xxx_ClearEventLogAOperation struct {
	Log            *Handle `idl:"name:LogHandle" json:"log"`
	BackupFileName *String `idl:"name:BackupFileName;pointer:unique" json:"backup_file_name"`
	Return         int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearEventLogAOperation) OpNum() int { return 12 }

func (o *xxx_ClearEventLogAOperation) OpName() string { return "/eventlog/v0/ElfrClearELFA" }

func (o *xxx_ClearEventLogAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearEventLogAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// BackupFileName {in} (1:{pointer=unique, alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.BackupFileName != nil {
			_ptr_BackupFileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BackupFileName != nil {
					if err := o.BackupFileName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BackupFileName, _ptr_BackupFileName); err != nil {
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

func (o *xxx_ClearEventLogAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// BackupFileName {in} (1:{pointer=unique, alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		_ptr_BackupFileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BackupFileName == nil {
				o.BackupFileName = &String{}
			}
			if err := o.BackupFileName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_BackupFileName := func(ptr interface{}) { o.BackupFileName = *ptr.(**String) }
		if err := w.ReadPointer(&o.BackupFileName, _s_BackupFileName, _ptr_BackupFileName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearEventLogAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearEventLogAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearEventLogAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ClearEventLogARequest structure represents the ElfrClearELFA operation request
type ClearEventLogARequest struct {
	// LogHandle: Handle to the event log to be cleared. This parameter is a server context
	// handle, as specified in section 2.2.6. This handle MUST NOT be one obtained via the
	// ElfrOpenBELA (section 3.1.4.2) method or the ElfrOpenBELW (section 3.1.4.1) method.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// BackupFileName: Provides an ANSI string (as specified in section 2.2.12) that points
	// to an NT Object Path of a file (as specified in section 2.2.4.1), in which a current
	// copy of the event log is to be placed. If this is NULL or empty, the server MUST
	// NOT create a backup as part of this method.
	BackupFileName *String `idl:"name:BackupFileName;pointer:unique" json:"backup_file_name"`
}

func (o *ClearEventLogARequest) xxx_ToOp(ctx context.Context, op *xxx_ClearEventLogAOperation) *xxx_ClearEventLogAOperation {
	if op == nil {
		op = &xxx_ClearEventLogAOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.BackupFileName = o.BackupFileName
	return op
}

func (o *ClearEventLogARequest) xxx_FromOp(ctx context.Context, op *xxx_ClearEventLogAOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.BackupFileName = op.BackupFileName
}
func (o *ClearEventLogARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ClearEventLogARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearEventLogAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearEventLogAResponse structure represents the ElfrClearELFA operation response
type ClearEventLogAResponse struct {
	// Return: The ElfrClearELFA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ClearEventLogAResponse) xxx_ToOp(ctx context.Context, op *xxx_ClearEventLogAOperation) *xxx_ClearEventLogAOperation {
	if op == nil {
		op = &xxx_ClearEventLogAOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ClearEventLogAResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearEventLogAOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ClearEventLogAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ClearEventLogAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearEventLogAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupEventLogAOperation structure represents the ElfrBackupELFA operation
type xxx_BackupEventLogAOperation struct {
	Log            *Handle `idl:"name:LogHandle" json:"log"`
	BackupFileName *String `idl:"name:BackupFileName" json:"backup_file_name"`
	Return         int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupEventLogAOperation) OpNum() int { return 13 }

func (o *xxx_BackupEventLogAOperation) OpName() string { return "/eventlog/v0/ElfrBackupELFA" }

func (o *xxx_BackupEventLogAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEventLogAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// BackupFileName {in} (1:{alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.BackupFileName != nil {
			if err := o.BackupFileName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEventLogAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// BackupFileName {in} (1:{alias=PRPC_STRING,pointer=ref}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.BackupFileName == nil {
			o.BackupFileName = &String{}
		}
		if err := o.BackupFileName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEventLogAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEventLogAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEventLogAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupEventLogARequest structure represents the ElfrBackupELFA operation request
type BackupEventLogARequest struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6. This handle MUST NOT be obtained via the ElfrOpenBELA
	// (section 3.1.4.2) method or the ElfrOpenBELW (section 3.1.4.1) method.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// BackupFileName: Provides an ANSI string (as specified in section 2.2.12) that points
	// to an NT Object Path of a file (as specified in section 2.2.4.1), in which a current
	// copy of the event log is to be placed. This MUST NOT be NULL or empty.
	BackupFileName *String `idl:"name:BackupFileName" json:"backup_file_name"`
}

func (o *BackupEventLogARequest) xxx_ToOp(ctx context.Context, op *xxx_BackupEventLogAOperation) *xxx_BackupEventLogAOperation {
	if op == nil {
		op = &xxx_BackupEventLogAOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.BackupFileName = o.BackupFileName
	return op
}

func (o *BackupEventLogARequest) xxx_FromOp(ctx context.Context, op *xxx_BackupEventLogAOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.BackupFileName = op.BackupFileName
}
func (o *BackupEventLogARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BackupEventLogARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupEventLogAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupEventLogAResponse structure represents the ElfrBackupELFA operation response
type BackupEventLogAResponse struct {
	// Return: The ElfrBackupELFA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupEventLogAResponse) xxx_ToOp(ctx context.Context, op *xxx_BackupEventLogAOperation) *xxx_BackupEventLogAOperation {
	if op == nil {
		op = &xxx_BackupEventLogAOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BackupEventLogAResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupEventLogAOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BackupEventLogAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BackupEventLogAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupEventLogAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenEventLogAOperation structure represents the ElfrOpenELA operation
type xxx_OpenEventLogAOperation struct {
	UNCServerName string  `idl:"name:UNCServerName" json:"unc_server_name"`
	ModuleName    *String `idl:"name:ModuleName" json:"module_name"`
	RegModuleName *String `idl:"name:RegModuleName" json:"reg_module_name"`
	MajorVersion  uint32  `idl:"name:MajorVersion" json:"major_version"`
	MinorVersion  uint32  `idl:"name:MinorVersion" json:"minor_version"`
	Log           *Handle `idl:"name:LogHandle" json:"log"`
	Return        int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenEventLogAOperation) OpNum() int { return 14 }

func (o *xxx_OpenEventLogAOperation) OpName() string { return "/eventlog/v0/ElfrOpenELA" }

func (o *xxx_OpenEventLogAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEventLogAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UNCServerName {in} (1:{handle, pointer=unique, alias=EVENTLOG_HANDLE_A}*(1)[dim:0,string](char))
	{
		if o.UNCServerName != "" {
			_ptr_UNCServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharString(ctx, w, o.UNCServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.UNCServerName, _ptr_UNCServerName); err != nil {
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
	// ModuleName {in} (1:{alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.ModuleName != nil {
			if err := o.ModuleName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// RegModuleName {in} (1:{alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.RegModuleName != nil {
			if err := o.RegModuleName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.MinorVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEventLogAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UNCServerName {in} (1:{handle, pointer=unique, alias=EVENTLOG_HANDLE_A}*(1)[dim:0,string](char))
	{
		_ptr_UNCServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharString(ctx, w, &o.UNCServerName); err != nil {
				return err
			}
			return nil
		})
		_s_UNCServerName := func(ptr interface{}) { o.UNCServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.UNCServerName, _s_UNCServerName, _ptr_UNCServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ModuleName {in} (1:{alias=PRPC_STRING,pointer=ref}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.ModuleName == nil {
			o.ModuleName = &String{}
		}
		if err := o.ModuleName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RegModuleName {in} (1:{alias=PRPC_STRING,pointer=ref}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.RegModuleName == nil {
			o.RegModuleName = &String{}
		}
		if err := o.RegModuleName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MinorVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEventLogAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEventLogAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// LogHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEventLogAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// LogHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenEventLogARequest structure represents the ElfrOpenELA operation request
type OpenEventLogARequest struct {
	// UNCServerName: A server interface handle. A pointer to an ANSI string (see [MSDN-ANSI])
	// specifying the server, as specified in section 2.2.7. The client MUST map this string
	// to an RPC binding handle, and the server MUST ignore this argument, as specified
	// in [C706] sections 4.3.5 and 5.1.5.2.
	UNCServerName string `idl:"name:UNCServerName" json:"unc_server_name"`
	// ModuleName: Specifies the event log name, as defined in section 1.8.2 and specified
	// in section 2.2.12, for which a handle is needed.
	ModuleName *String `idl:"name:ModuleName" json:"module_name"`
	// RegModuleName: This parameter MUST be ignored by the server. Clients MUST specify
	// an empty string.
	RegModuleName *String `idl:"name:RegModuleName" json:"reg_module_name"`
	// MajorVersion: Major version of the client. This value MUST be set to 1.
	MajorVersion uint32 `idl:"name:MajorVersion" json:"major_version"`
	// MinorVersion: Minor version of the client. This value MUST be set to 1.
	MinorVersion uint32 `idl:"name:MinorVersion" json:"minor_version"`
}

func (o *OpenEventLogARequest) xxx_ToOp(ctx context.Context, op *xxx_OpenEventLogAOperation) *xxx_OpenEventLogAOperation {
	if op == nil {
		op = &xxx_OpenEventLogAOperation{}
	}
	if o == nil {
		return op
	}
	op.UNCServerName = o.UNCServerName
	op.ModuleName = o.ModuleName
	op.RegModuleName = o.RegModuleName
	op.MajorVersion = o.MajorVersion
	op.MinorVersion = o.MinorVersion
	return op
}

func (o *OpenEventLogARequest) xxx_FromOp(ctx context.Context, op *xxx_OpenEventLogAOperation) {
	if o == nil {
		return
	}
	o.UNCServerName = op.UNCServerName
	o.ModuleName = op.ModuleName
	o.RegModuleName = op.RegModuleName
	o.MajorVersion = op.MajorVersion
	o.MinorVersion = op.MinorVersion
}
func (o *OpenEventLogARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenEventLogARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenEventLogAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenEventLogAResponse structure represents the ElfrOpenELA operation response
type OpenEventLogAResponse struct {
	// LogHandle: Pointer to an event log handle. This parameter is a server context handle,
	// as specified in section 2.2.6. This handle MUST be closed by using the ElfrCloseEL
	// (section 3.1.4.21) method once the handle is no longer needed. In the case when the
	// client cannot call the ElfrCloseEL function, such as the abnormal termination of
	// the client, this context handle will be revoked by the server so that there will
	// not be any resource leaks.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// Return: The ElfrOpenELA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenEventLogAResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenEventLogAOperation) *xxx_OpenEventLogAOperation {
	if op == nil {
		op = &xxx_OpenEventLogAOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.Return = o.Return
	return op
}

func (o *OpenEventLogAResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenEventLogAOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.Return = op.Return
}
func (o *OpenEventLogAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenEventLogAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenEventLogAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterEventSourceAOperation structure represents the ElfrRegisterEventSourceA operation
type xxx_RegisterEventSourceAOperation struct {
	UNCServerName string  `idl:"name:UNCServerName" json:"unc_server_name"`
	ModuleName    *String `idl:"name:ModuleName" json:"module_name"`
	RegModuleName *String `idl:"name:RegModuleName" json:"reg_module_name"`
	MajorVersion  uint32  `idl:"name:MajorVersion" json:"major_version"`
	MinorVersion  uint32  `idl:"name:MinorVersion" json:"minor_version"`
	Log           *Handle `idl:"name:LogHandle" json:"log"`
	Return        int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterEventSourceAOperation) OpNum() int { return 15 }

func (o *xxx_RegisterEventSourceAOperation) OpName() string {
	return "/eventlog/v0/ElfrRegisterEventSourceA"
}

func (o *xxx_RegisterEventSourceAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterEventSourceAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UNCServerName {in} (1:{handle, pointer=unique, alias=EVENTLOG_HANDLE_A}*(1)[dim:0,string](char))
	{
		if o.UNCServerName != "" {
			_ptr_UNCServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharString(ctx, w, o.UNCServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.UNCServerName, _ptr_UNCServerName); err != nil {
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
	// ModuleName {in} (1:{alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.ModuleName != nil {
			if err := o.ModuleName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// RegModuleName {in} (1:{alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.RegModuleName != nil {
			if err := o.RegModuleName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.MinorVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterEventSourceAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UNCServerName {in} (1:{handle, pointer=unique, alias=EVENTLOG_HANDLE_A}*(1)[dim:0,string](char))
	{
		_ptr_UNCServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharString(ctx, w, &o.UNCServerName); err != nil {
				return err
			}
			return nil
		})
		_s_UNCServerName := func(ptr interface{}) { o.UNCServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.UNCServerName, _s_UNCServerName, _ptr_UNCServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ModuleName {in} (1:{alias=PRPC_STRING,pointer=ref}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.ModuleName == nil {
			o.ModuleName = &String{}
		}
		if err := o.ModuleName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RegModuleName {in} (1:{alias=PRPC_STRING,pointer=ref}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.RegModuleName == nil {
			o.RegModuleName = &String{}
		}
		if err := o.RegModuleName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MinorVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterEventSourceAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterEventSourceAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// LogHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterEventSourceAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// LogHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RegisterEventSourceARequest structure represents the ElfrRegisterEventSourceA operation request
type RegisterEventSourceARequest struct {
	// UNCServerName: A server interface handle. A pointer to an ANSI string (see [MSDN-ANSI])
	// specifying the server, as specified in section 2.2.7. The client MUST map this string
	// to an RPC binding handle, and the server MUST ignore this argument, as specified
	// in [C706] sections 4.3.5 and 5.1.5.2.
	UNCServerName string `idl:"name:UNCServerName" json:"unc_server_name"`
	// ModuleName: Specifies the event source, as defined in section 1.8.3 and specified
	// in section 2.2.12, for which a handle is needed.
	ModuleName *String `idl:"name:ModuleName" json:"module_name"`
	// RegModuleName: This parameter MUST be ignored by the server. Clients MUST specify
	// an empty string.
	RegModuleName *String `idl:"name:RegModuleName" json:"reg_module_name"`
	// MajorVersion: Major version of the client. This value MUST be set to 1.
	MajorVersion uint32 `idl:"name:MajorVersion" json:"major_version"`
	// MinorVersion: Minor version of the client. This value MUST be set to 1.
	MinorVersion uint32 `idl:"name:MinorVersion" json:"minor_version"`
}

func (o *RegisterEventSourceARequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterEventSourceAOperation) *xxx_RegisterEventSourceAOperation {
	if op == nil {
		op = &xxx_RegisterEventSourceAOperation{}
	}
	if o == nil {
		return op
	}
	op.UNCServerName = o.UNCServerName
	op.ModuleName = o.ModuleName
	op.RegModuleName = o.RegModuleName
	op.MajorVersion = o.MajorVersion
	op.MinorVersion = o.MinorVersion
	return op
}

func (o *RegisterEventSourceARequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterEventSourceAOperation) {
	if o == nil {
		return
	}
	o.UNCServerName = op.UNCServerName
	o.ModuleName = op.ModuleName
	o.RegModuleName = op.RegModuleName
	o.MajorVersion = op.MajorVersion
	o.MinorVersion = op.MinorVersion
}
func (o *RegisterEventSourceARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterEventSourceARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterEventSourceAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterEventSourceAResponse structure represents the ElfrRegisterEventSourceA operation response
type RegisterEventSourceAResponse struct {
	// LogHandle: Pointer to an event log handle. This parameter is a server context handle,
	// as specified in section 2.2.6.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// Return: The ElfrRegisterEventSourceA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterEventSourceAResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterEventSourceAOperation) *xxx_RegisterEventSourceAOperation {
	if op == nil {
		op = &xxx_RegisterEventSourceAOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.Return = o.Return
	return op
}

func (o *RegisterEventSourceAResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterEventSourceAOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.Return = op.Return
}
func (o *RegisterEventSourceAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterEventSourceAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterEventSourceAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenBackupEventLogAOperation structure represents the ElfrOpenBELA operation
type xxx_OpenBackupEventLogAOperation struct {
	UNCServerName  string  `idl:"name:UNCServerName" json:"unc_server_name"`
	BackupFileName *String `idl:"name:BackupFileName" json:"backup_file_name"`
	MajorVersion   uint32  `idl:"name:MajorVersion" json:"major_version"`
	MinorVersion   uint32  `idl:"name:MinorVersion" json:"minor_version"`
	Log            *Handle `idl:"name:LogHandle" json:"log"`
	Return         int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenBackupEventLogAOperation) OpNum() int { return 16 }

func (o *xxx_OpenBackupEventLogAOperation) OpName() string { return "/eventlog/v0/ElfrOpenBELA" }

func (o *xxx_OpenBackupEventLogAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenBackupEventLogAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UNCServerName {in} (1:{handle, pointer=unique, alias=EVENTLOG_HANDLE_A}*(1)[dim:0,string](char))
	{
		if o.UNCServerName != "" {
			_ptr_UNCServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharString(ctx, w, o.UNCServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.UNCServerName, _ptr_UNCServerName); err != nil {
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
	// BackupFileName {in} (1:{alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.BackupFileName != nil {
			if err := o.BackupFileName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.MinorVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenBackupEventLogAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UNCServerName {in} (1:{handle, pointer=unique, alias=EVENTLOG_HANDLE_A}*(1)[dim:0,string](char))
	{
		_ptr_UNCServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharString(ctx, w, &o.UNCServerName); err != nil {
				return err
			}
			return nil
		})
		_s_UNCServerName := func(ptr interface{}) { o.UNCServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.UNCServerName, _s_UNCServerName, _ptr_UNCServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BackupFileName {in} (1:{alias=PRPC_STRING,pointer=ref}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.BackupFileName == nil {
			o.BackupFileName = &String{}
		}
		if err := o.BackupFileName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MinorVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenBackupEventLogAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenBackupEventLogAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// LogHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenBackupEventLogAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// LogHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenBackupEventLogARequest structure represents the ElfrOpenBELA operation request
type OpenBackupEventLogARequest struct {
	// UNCServerName: A server interface handle. A pointer to an ANSI string (see [MSDN-ANSI])
	// specifying the server, as specified in section 2.2.7. The client MUST map this string
	// to an RPC binding handle, and the server MUST ignore this argument, as specified
	// in [C706] sections 4.3.5 and 5.1.5.2.
	UNCServerName string `idl:"name:UNCServerName" json:"unc_server_name"`
	// BackupFileName: Provides an ANSI string (as specified in section 2.2.12) that points
	// to an NT Object Path of the file where the backup event log is located, as specified
	// in section 2.2.4.1.
	BackupFileName *String `idl:"name:BackupFileName" json:"backup_file_name"`
	// MajorVersion: Major version of the client. This value MUST be set to 1.
	MajorVersion uint32 `idl:"name:MajorVersion" json:"major_version"`
	// MinorVersion: Minor version of the client. This value MUST be set to 1.
	MinorVersion uint32 `idl:"name:MinorVersion" json:"minor_version"`
}

func (o *OpenBackupEventLogARequest) xxx_ToOp(ctx context.Context, op *xxx_OpenBackupEventLogAOperation) *xxx_OpenBackupEventLogAOperation {
	if op == nil {
		op = &xxx_OpenBackupEventLogAOperation{}
	}
	if o == nil {
		return op
	}
	op.UNCServerName = o.UNCServerName
	op.BackupFileName = o.BackupFileName
	op.MajorVersion = o.MajorVersion
	op.MinorVersion = o.MinorVersion
	return op
}

func (o *OpenBackupEventLogARequest) xxx_FromOp(ctx context.Context, op *xxx_OpenBackupEventLogAOperation) {
	if o == nil {
		return
	}
	o.UNCServerName = op.UNCServerName
	o.BackupFileName = op.BackupFileName
	o.MajorVersion = op.MajorVersion
	o.MinorVersion = op.MinorVersion
}
func (o *OpenBackupEventLogARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenBackupEventLogARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenBackupEventLogAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenBackupEventLogAResponse structure represents the ElfrOpenBELA operation response
type OpenBackupEventLogAResponse struct {
	// LogHandle: Pointer to an event log handle. This parameter is a server context handle,
	// as specified in section 2.2.6. This handle MUST be closed by using the ElfrCloseEL
	// (section 3.1.4.21) method once the handle is no longer needed. In the case when the
	// client cannot call the ElfrCloseEL function, such as the abnormal termination of
	// the client, this context handle will be revoked by the server so that there will
	// not be any resource leaks.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// Return: The ElfrOpenBELA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenBackupEventLogAResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenBackupEventLogAOperation) *xxx_OpenBackupEventLogAOperation {
	if op == nil {
		op = &xxx_OpenBackupEventLogAOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.Return = o.Return
	return op
}

func (o *OpenBackupEventLogAResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenBackupEventLogAOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.Return = op.Return
}
func (o *OpenBackupEventLogAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenBackupEventLogAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenBackupEventLogAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReadEventLogAOperation structure represents the ElfrReadELA operation
type xxx_ReadEventLogAOperation struct {
	Log                    *Handle `idl:"name:LogHandle" json:"log"`
	ReadFlags              uint32  `idl:"name:ReadFlags" json:"read_flags"`
	RecordOffset           uint32  `idl:"name:RecordOffset" json:"record_offset"`
	NumberOfBytesToRead    uint32  `idl:"name:NumberOfBytesToRead" json:"number_of_bytes_to_read"`
	Buffer                 []byte  `idl:"name:Buffer;size_is:(NumberOfBytesToRead)" json:"buffer"`
	NumberOfBytesRead      uint32  `idl:"name:NumberOfBytesRead" json:"number_of_bytes_read"`
	MinNumberOfBytesNeeded uint32  `idl:"name:MinNumberOfBytesNeeded" json:"min_number_of_bytes_needed"`
	Return                 int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_ReadEventLogAOperation) OpNum() int { return 17 }

func (o *xxx_ReadEventLogAOperation) OpName() string { return "/eventlog/v0/ElfrReadELA" }

func (o *xxx_ReadEventLogAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadEventLogAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ReadFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.ReadFlags); err != nil {
			return err
		}
	}
	// RecordOffset {in} (1:(uint32))
	{
		if err := w.WriteData(o.RecordOffset); err != nil {
			return err
		}
	}
	// NumberOfBytesToRead {in} (1:{range=(0,524287), alias=RULONG}(uint32))
	{
		if err := w.WriteData(o.NumberOfBytesToRead); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadEventLogAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ReadFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ReadFlags); err != nil {
			return err
		}
	}
	// RecordOffset {in} (1:(uint32))
	{
		if err := w.ReadData(&o.RecordOffset); err != nil {
			return err
		}
	}
	// NumberOfBytesToRead {in} (1:{range=(0,524287), alias=RULONG}(uint32))
	{
		if err := w.ReadData(&o.NumberOfBytesToRead); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadEventLogAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadEventLogAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=NumberOfBytesToRead](uchar))
	{
		dimSize1 := uint64(o.NumberOfBytesToRead)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Buffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Buffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// NumberOfBytesRead {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.NumberOfBytesRead); err != nil {
			return err
		}
	}
	// MinNumberOfBytesNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MinNumberOfBytesNeeded); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadEventLogAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=NumberOfBytesToRead](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
	}
	// NumberOfBytesRead {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.NumberOfBytesRead); err != nil {
			return err
		}
	}
	// MinNumberOfBytesNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MinNumberOfBytesNeeded); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReadEventLogARequest structure represents the ElfrReadELA operation request
type ReadEventLogARequest struct {
	// LogHandle: Handle to an event log to read. This parameter is a server context handle,
	// as specified in section 2.2.6.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// ReadFlags: The caller MUST specify if the read is to start at a specific record,
	// or is to proceed from the last record read. The value MUST be one and only one of
	// the following flags.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EVENTLOG_SEQUENTIAL_READ 0x00000001 | Read operation proceeds sequentially from the last call to the ElfrReadELA       |
	//	|                                     | (section 3.1.4.8) method or the ElfrReadELW (section 3.1.4.7) method, using this |
	//	|                                     | handle. This flag cannot be used with EVENTLOG_SEEK_READ.                        |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EVENTLOG_SEEK_READ 0x00000002       | Read operation proceeds from the record specified by the RecordOffset parameter. |
	//	|                                     | This flag cannot be used with EVENTLOG_SEQUENTIAL_READ.                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// Because the method reads as many records as can fit in the buffer, the caller MUST
	// also set one and only one of the following flags to indicate the direction for successive
	// read operations.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| EVENTLOG_FORWARDS_READ 0x00000004  | Log is read in chronological order. This flag cannot be used with                |
	//	|                                    | EVENTLOG_BACKWARDS_READ.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| EVENTLOG_BACKWARDS_READ 0x00000008 | Log is read in reverse chronological order. This flag cannot be used with        |
	//	|                                    | EVENTLOG_FORWARDS_READ.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	ReadFlags uint32 `idl:"name:ReadFlags" json:"read_flags"`
	// RecordOffset: Log entry record number at which the read operation is to start. Each
	// event in a log has a record number. This parameter MUST be ignored unless the EVENTLOG_SEEK_READ
	// bit is set in the ReadFlags parameter.
	RecordOffset uint32 `idl:"name:RecordOffset" json:"record_offset"`
	// NumberOfBytesToRead: Size of the buffer in bytes. This is the maximum amount of data
	// that can be read.
	NumberOfBytesToRead uint32 `idl:"name:NumberOfBytesToRead" json:"number_of_bytes_to_read"`
}

func (o *ReadEventLogARequest) xxx_ToOp(ctx context.Context, op *xxx_ReadEventLogAOperation) *xxx_ReadEventLogAOperation {
	if op == nil {
		op = &xxx_ReadEventLogAOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.ReadFlags = o.ReadFlags
	op.RecordOffset = o.RecordOffset
	op.NumberOfBytesToRead = o.NumberOfBytesToRead
	return op
}

func (o *ReadEventLogARequest) xxx_FromOp(ctx context.Context, op *xxx_ReadEventLogAOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.ReadFlags = op.ReadFlags
	o.RecordOffset = op.RecordOffset
	o.NumberOfBytesToRead = op.NumberOfBytesToRead
}
func (o *ReadEventLogARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReadEventLogARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReadEventLogAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReadEventLogAResponse structure represents the ElfrReadELA operation response
type ReadEventLogAResponse struct {
	// XXX: NumberOfBytesToRead is an implicit input depedency for output parameters
	NumberOfBytesToRead uint32 `idl:"name:NumberOfBytesToRead" json:"number_of_bytes_to_read"`

	// Buffer: Data read from the event log.
	Buffer []byte `idl:"name:Buffer;size_is:(NumberOfBytesToRead)" json:"buffer"`
	// NumberOfBytesRead: Number of bytes read by the method.
	NumberOfBytesRead uint32 `idl:"name:NumberOfBytesRead" json:"number_of_bytes_read"`
	// MinNumberOfBytesNeeded: If the method fails because the buffer is too small to fit
	// even a single record, this MUST be set to the minimum number of bytes needed to fit
	// the next record. Otherwise, this MUST NOT be set, and MUST be ignored by the caller.
	MinNumberOfBytesNeeded uint32 `idl:"name:MinNumberOfBytesNeeded" json:"min_number_of_bytes_needed"`
	// Return: The ElfrReadELA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReadEventLogAResponse) xxx_ToOp(ctx context.Context, op *xxx_ReadEventLogAOperation) *xxx_ReadEventLogAOperation {
	if op == nil {
		op = &xxx_ReadEventLogAOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.NumberOfBytesToRead == uint32(0) {
		op.NumberOfBytesToRead = o.NumberOfBytesToRead
	}

	op.Buffer = o.Buffer
	op.NumberOfBytesRead = o.NumberOfBytesRead
	op.MinNumberOfBytesNeeded = o.MinNumberOfBytesNeeded
	op.Return = o.Return
	return op
}

func (o *ReadEventLogAResponse) xxx_FromOp(ctx context.Context, op *xxx_ReadEventLogAOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.NumberOfBytesToRead = op.NumberOfBytesToRead

	o.Buffer = op.Buffer
	o.NumberOfBytesRead = op.NumberOfBytesRead
	o.MinNumberOfBytesNeeded = op.MinNumberOfBytesNeeded
	o.Return = op.Return
}
func (o *ReadEventLogAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReadEventLogAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReadEventLogAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReportEventAOperation structure represents the ElfrReportEventA operation
type xxx_ReportEventAOperation struct {
	Log           *Handle   `idl:"name:LogHandle" json:"log"`
	Time          uint32    `idl:"name:Time" json:"time"`
	EventType     uint16    `idl:"name:EventType" json:"event_type"`
	EventCategory uint16    `idl:"name:EventCategory" json:"event_category"`
	EventID       uint32    `idl:"name:EventID" json:"event_id"`
	StringsLength uint16    `idl:"name:NumStrings" json:"strings_length"`
	DataSize      uint32    `idl:"name:DataSize" json:"data_size"`
	ComputerName  *String   `idl:"name:ComputerName" json:"computer_name"`
	UserSID       *dtyp.SID `idl:"name:UserSID;pointer:unique" json:"user_sid"`
	Strings       []*String `idl:"name:Strings;size_is:(NumStrings);pointer:unique" json:"strings"`
	Data          []byte    `idl:"name:Data;size_is:(DataSize);pointer:unique" json:"data"`
	Flags         uint16    `idl:"name:Flags" json:"flags"`
	RecordNumber  uint32    `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	TimeWritten   uint32    `idl:"name:TimeWritten;pointer:unique" json:"time_written"`
	Return        int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_ReportEventAOperation) OpNum() int { return 18 }

func (o *xxx_ReportEventAOperation) OpName() string { return "/eventlog/v0/ElfrReportEventA" }

func (o *xxx_ReportEventAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Strings != nil && o.StringsLength == 0 {
		o.StringsLength = uint16(len(o.Strings))
	}
	if o.Data != nil && o.DataSize == 0 {
		o.DataSize = uint32(len(o.Data))
	}
	if o.StringsLength > uint16(256) {
		return fmt.Errorf("StringsLength is out of range")
	}
	if o.DataSize > uint32(61440) {
		return fmt.Errorf("DataSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Time {in} (1:(uint32))
	{
		if err := w.WriteData(o.Time); err != nil {
			return err
		}
	}
	// EventType {in} (1:(uint16))
	{
		if err := w.WriteData(o.EventType); err != nil {
			return err
		}
	}
	// EventCategory {in} (1:(uint16))
	{
		if err := w.WriteData(o.EventCategory); err != nil {
			return err
		}
	}
	// EventID {in} (1:(uint32))
	{
		if err := w.WriteData(o.EventID); err != nil {
			return err
		}
	}
	// NumStrings {in} (1:{range=(0,256)}(uint16))
	{
		if err := w.WriteData(o.StringsLength); err != nil {
			return err
		}
	}
	// DataSize {in} (1:{range=(0,61440)}(uint32))
	{
		if err := w.WriteData(o.DataSize); err != nil {
			return err
		}
	}
	// ComputerName {in} (1:{alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.ComputerName != nil {
			if err := o.ComputerName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// UserSID {in} (1:{pointer=unique, alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.UserSID != nil {
			_ptr_UserSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserSID != nil {
					if err := o.UserSID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserSID, _ptr_UserSID); err != nil {
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
	// Strings {in} (1:{pointer=unique}[dim:0,size_is=NumStrings])(2:{alias=PRPC_STRING}*(1))(3:{alias=RPC_STRING}(struct))
	{
		dimSize1 := uint64(o.StringsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Strings {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Strings[i1] != nil {
				_ptr_Strings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.Strings[i1] != nil {
						if err := o.Strings[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&String{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Strings[i1], _ptr_Strings); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Strings); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Data {in} (1:{pointer=unique}*(1)[dim:0,size_is=DataSize](uchar))
	{
		if o.Data != nil || o.DataSize > 0 {
			_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DataSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Data {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Data[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data, _ptr_Data); err != nil {
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
	// Flags {in} (1:(uint16))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TimeWritten {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_TimeWritten := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.TimeWritten); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TimeWritten, _ptr_TimeWritten); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Time {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Time); err != nil {
			return err
		}
	}
	// EventType {in} (1:(uint16))
	{
		if err := w.ReadData(&o.EventType); err != nil {
			return err
		}
	}
	// EventCategory {in} (1:(uint16))
	{
		if err := w.ReadData(&o.EventCategory); err != nil {
			return err
		}
	}
	// EventID {in} (1:(uint32))
	{
		if err := w.ReadData(&o.EventID); err != nil {
			return err
		}
	}
	// NumStrings {in} (1:{range=(0,256)}(uint16))
	{
		if err := w.ReadData(&o.StringsLength); err != nil {
			return err
		}
	}
	// DataSize {in} (1:{range=(0,61440)}(uint32))
	{
		if err := w.ReadData(&o.DataSize); err != nil {
			return err
		}
	}
	// ComputerName {in} (1:{alias=PRPC_STRING,pointer=ref}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.ComputerName == nil {
			o.ComputerName = &String{}
		}
		if err := o.ComputerName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UserSID {in} (1:{pointer=unique, alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		_ptr_UserSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserSID == nil {
				o.UserSID = &dtyp.SID{}
			}
			if err := o.UserSID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_UserSID := func(ptr interface{}) { o.UserSID = *ptr.(**dtyp.SID) }
		if err := w.ReadPointer(&o.UserSID, _s_UserSID, _ptr_UserSID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Strings {in} (1:{pointer=unique}[dim:0,size_is=NumStrings])(2:{alias=PRPC_STRING,pointer=ref}*(1))(3:{alias=RPC_STRING}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Strings", sizeInfo[0])
		}
		o.Strings = make([]*String, sizeInfo[0])
		for i1 := range o.Strings {
			i1 := i1
			_ptr_Strings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Strings[i1] == nil {
					o.Strings[i1] = &String{}
				}
				if err := o.Strings[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Strings := func(ptr interface{}) { o.Strings[i1] = *ptr.(**String) }
			if err := w.ReadPointer(&o.Strings[i1], _s_Strings, _ptr_Strings); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Data {in} (1:{pointer=unique}*(1)[dim:0,size_is=DataSize](uchar))
	{
		_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
			}
			o.Data = make([]byte, sizeInfo[0])
			for i1 := range o.Data {
				i1 := i1
				if err := w.ReadData(&o.Data[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_Data := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Data, _s_Data, _ptr_Data); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Flags {in} (1:(uint16))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		_s_RecordNumber := func(ptr interface{}) { o.RecordNumber = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.RecordNumber, _s_RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TimeWritten {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_TimeWritten := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.TimeWritten); err != nil {
				return err
			}
			return nil
		})
		_s_TimeWritten := func(ptr interface{}) { o.TimeWritten = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.TimeWritten, _s_TimeWritten, _ptr_TimeWritten); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TimeWritten {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_TimeWritten := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.TimeWritten); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TimeWritten, _ptr_TimeWritten); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		_s_RecordNumber := func(ptr interface{}) { o.RecordNumber = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.RecordNumber, _s_RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TimeWritten {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_TimeWritten := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.TimeWritten); err != nil {
				return err
			}
			return nil
		})
		_s_TimeWritten := func(ptr interface{}) { o.TimeWritten = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.TimeWritten, _s_TimeWritten, _ptr_TimeWritten); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReportEventARequest structure represents the ElfrReportEventA operation request
type ReportEventARequest struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6. This handle MUST NOT be obtained via the ElfrOpenBELA
	// (section 3.1.4.2) method or the ElfrOpenBELW (section 3.1.4.1) method.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// Time: Time at which the event was generated by the event source (not the time at
	// which the event was). The time MUST be expressed as the number of seconds since 00:00:00
	// on January 1, 1970 (UTC).
	Time uint32 `idl:"name:Time" json:"time"`
	// EventType: Type of the event, as specified in section 2.2.2.
	EventType uint16 `idl:"name:EventType" json:"event_type"`
	// EventCategory: Event category, as specified in section 1.8.5.
	EventCategory uint16 `idl:"name:EventCategory" json:"event_category"`
	// EventID: EventID, as specified in section 3.1.1.4.
	EventID uint32 `idl:"name:EventID" json:"event_id"`
	// NumStrings: Number of strings in the array pointed to by the Strings parameter. A
	// value of zero indicates that no strings are present.
	StringsLength uint16 `idl:"name:NumStrings" json:"strings_length"`
	// DataSize: Number of bytes of event-specific raw binary data to write to the log.
	// This binary data is passed in the Data parameter. If no event-specific data is present,
	// this parameter MUST be set to zero.
	DataSize uint32 `idl:"name:DataSize" json:"data_size"`
	// ComputerName: A string to assist in identifying the machine that generated the event.
	// In practice, the name of the computer. There are no character restrictions on this
	// field's content (for example, a FQDN can be used).<31>
	ComputerName *String `idl:"name:ComputerName" json:"computer_name"`
	// UserSID: Either NULL or a user SID. If this is NULL, the event is to have a zero
	// length UserSid field.
	UserSID *dtyp.SID `idl:"name:UserSID;pointer:unique" json:"user_sid"`
	// Strings: Specifies strings containing information specific to the event. This parameter
	// MUST be a valid pointer. If the NumStrings parameter is zero, this parameter MUST
	// be NULL. For example, an event relating to file deletion could use a string to specify
	// the path of the file being deleted.
	Strings []*String `idl:"name:Strings;size_is:(NumStrings);pointer:unique" json:"strings"`
	// Data: Pointer to the buffer that contains the event-specific binary data. This parameter
	// MUST be a valid pointer (or NULL), even if the DataSize parameter is 0.
	Data []byte `idl:"name:Data;size_is:(DataSize);pointer:unique" json:"data"`
	// Flags: Unused. MUST be set to zero when sent and MUST be ignored on receipt.
	Flags uint16 `idl:"name:Flags" json:"flags"`
	// RecordNumber: Unused. Can be set to any arbitrary value when sent, and any value
	// sent by the client MUST be ignored on receipt by the server.
	RecordNumber uint32 `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	// TimeWritten: Unused. Can be set to any arbitrary value when sent, and any value sent
	// by the client MUST be ignored on receipt by the server.
	TimeWritten uint32 `idl:"name:TimeWritten;pointer:unique" json:"time_written"`
}

func (o *ReportEventARequest) xxx_ToOp(ctx context.Context, op *xxx_ReportEventAOperation) *xxx_ReportEventAOperation {
	if op == nil {
		op = &xxx_ReportEventAOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.Time = o.Time
	op.EventType = o.EventType
	op.EventCategory = o.EventCategory
	op.EventID = o.EventID
	op.StringsLength = o.StringsLength
	op.DataSize = o.DataSize
	op.ComputerName = o.ComputerName
	op.UserSID = o.UserSID
	op.Strings = o.Strings
	op.Data = o.Data
	op.Flags = o.Flags
	op.RecordNumber = o.RecordNumber
	op.TimeWritten = o.TimeWritten
	return op
}

func (o *ReportEventARequest) xxx_FromOp(ctx context.Context, op *xxx_ReportEventAOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.Time = op.Time
	o.EventType = op.EventType
	o.EventCategory = op.EventCategory
	o.EventID = op.EventID
	o.StringsLength = op.StringsLength
	o.DataSize = op.DataSize
	o.ComputerName = op.ComputerName
	o.UserSID = op.UserSID
	o.Strings = op.Strings
	o.Data = op.Data
	o.Flags = op.Flags
	o.RecordNumber = op.RecordNumber
	o.TimeWritten = op.TimeWritten
}
func (o *ReportEventARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReportEventARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportEventAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReportEventAResponse structure represents the ElfrReportEventA operation response
type ReportEventAResponse struct {
	// RecordNumber: Unused. Can be set to any arbitrary value when sent, and any value
	// sent by the client MUST be ignored on receipt by the server.
	RecordNumber uint32 `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	// TimeWritten: Unused. Can be set to any arbitrary value when sent, and any value sent
	// by the client MUST be ignored on receipt by the server.
	TimeWritten uint32 `idl:"name:TimeWritten;pointer:unique" json:"time_written"`
	// Return: The ElfrReportEventA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReportEventAResponse) xxx_ToOp(ctx context.Context, op *xxx_ReportEventAOperation) *xxx_ReportEventAOperation {
	if op == nil {
		op = &xxx_ReportEventAOperation{}
	}
	if o == nil {
		return op
	}
	op.RecordNumber = o.RecordNumber
	op.TimeWritten = o.TimeWritten
	op.Return = o.Return
	return op
}

func (o *ReportEventAResponse) xxx_FromOp(ctx context.Context, op *xxx_ReportEventAOperation) {
	if o == nil {
		return
	}
	o.RecordNumber = op.RecordNumber
	o.TimeWritten = op.TimeWritten
	o.Return = op.Return
}
func (o *ReportEventAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReportEventAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportEventAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLogInformationOperation structure represents the ElfrGetLogInformation operation
type xxx_GetLogInformationOperation struct {
	Log               *Handle `idl:"name:LogHandle" json:"log"`
	InfoLevel         uint32  `idl:"name:InfoLevel" json:"info_level"`
	Buffer            []byte  `idl:"name:lpBuffer;size_is:(cbBufSize)" json:"buffer"`
	BufferLength      uint32  `idl:"name:cbBufSize" json:"buffer_length"`
	BytesNeededLength uint32  `idl:"name:pcbBytesNeeded" json:"bytes_needed_length"`
	Return            int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLogInformationOperation) OpNum() int { return 22 }

func (o *xxx_GetLogInformationOperation) OpName() string { return "/eventlog/v0/ElfrGetLogInformation" }

func (o *xxx_GetLogInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.BufferLength > uint32(1024) {
		return fmt.Errorf("BufferLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InfoLevel {in} (1:(uint32))
	{
		if err := w.WriteData(o.InfoLevel); err != nil {
			return err
		}
	}
	// cbBufSize {in} (1:{range=(0,1024)}(uint32))
	{
		if err := w.WriteData(o.BufferLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InfoLevel {in} (1:(uint32))
	{
		if err := w.ReadData(&o.InfoLevel); err != nil {
			return err
		}
	}
	// cbBufSize {in} (1:{range=(0,1024)}(uint32))
	{
		if err := w.ReadData(&o.BufferLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpBuffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=cbBufSize](uchar))
	{
		dimSize1 := uint64(o.BufferLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Buffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Buffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pcbBytesNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.BytesNeededLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpBuffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=cbBufSize](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
	}
	// pcbBytesNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.BytesNeededLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetLogInformationRequest structure represents the ElfrGetLogInformation operation request
type GetLogInformationRequest struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// InfoLevel: The level of event log information to return. This MUST be set to zero.
	InfoLevel uint32 `idl:"name:InfoLevel" json:"info_level"`
	// cbBufSize: The size in bytes of the buffer pointed to by the lpBuffer parameter.
	BufferLength uint32 `idl:"name:cbBufSize" json:"buffer_length"`
}

func (o *GetLogInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLogInformationOperation) *xxx_GetLogInformationOperation {
	if op == nil {
		op = &xxx_GetLogInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.InfoLevel = o.InfoLevel
	op.BufferLength = o.BufferLength
	return op
}

func (o *GetLogInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLogInformationOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.InfoLevel = op.InfoLevel
	o.BufferLength = op.BufferLength
}
func (o *GetLogInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLogInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLogInformationResponse structure represents the ElfrGetLogInformation operation response
type GetLogInformationResponse struct {
	// XXX: cbBufSize is an implicit input depedency for output parameters
	BufferLength uint32 `idl:"name:cbBufSize" json:"buffer_length"`

	// lpBuffer: The event log information. This MUST point to either an EVENTLOG_FULL_INFORMATION
	// (section 2.2.4) structure or be NULL.
	Buffer []byte `idl:"name:lpBuffer;size_is:(cbBufSize)" json:"buffer"`
	// pcbBytesNeeded: Number of bytes required for the requested information, regardless
	// of if the function succeeds. This parameter MUST NOT be NULL.
	BytesNeededLength uint32 `idl:"name:pcbBytesNeeded" json:"bytes_needed_length"`
	// Return: The ElfrGetLogInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLogInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLogInformationOperation) *xxx_GetLogInformationOperation {
	if op == nil {
		op = &xxx_GetLogInformationOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferLength == uint32(0) {
		op.BufferLength = o.BufferLength
	}

	op.Buffer = o.Buffer
	op.BytesNeededLength = o.BytesNeededLength
	op.Return = o.Return
	return op
}

func (o *GetLogInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLogInformationOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferLength = op.BufferLength

	o.Buffer = op.Buffer
	o.BytesNeededLength = op.BytesNeededLength
	o.Return = op.Return
}
func (o *GetLogInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLogInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReportEventAndSourceWOperation structure represents the ElfrReportEventAndSourceW operation
type xxx_ReportEventAndSourceWOperation struct {
	Log           *Handle               `idl:"name:LogHandle" json:"log"`
	Time          uint32                `idl:"name:Time" json:"time"`
	EventType     uint16                `idl:"name:EventType" json:"event_type"`
	EventCategory uint16                `idl:"name:EventCategory" json:"event_category"`
	EventID       uint32                `idl:"name:EventID" json:"event_id"`
	SourceName    *dtyp.UnicodeString   `idl:"name:SourceName" json:"source_name"`
	StringsLength uint16                `idl:"name:NumStrings" json:"strings_length"`
	DataSize      uint32                `idl:"name:DataSize" json:"data_size"`
	ComputerName  *dtyp.UnicodeString   `idl:"name:ComputerName" json:"computer_name"`
	UserSID       *dtyp.SID             `idl:"name:UserSID;pointer:unique" json:"user_sid"`
	Strings       []*dtyp.UnicodeString `idl:"name:Strings;size_is:(NumStrings);pointer:unique" json:"strings"`
	Data          []byte                `idl:"name:Data;size_is:(DataSize);pointer:unique" json:"data"`
	Flags         uint16                `idl:"name:Flags" json:"flags"`
	RecordNumber  uint32                `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	TimeWritten   uint32                `idl:"name:TimeWritten;pointer:unique" json:"time_written"`
	Return        int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_ReportEventAndSourceWOperation) OpNum() int { return 24 }

func (o *xxx_ReportEventAndSourceWOperation) OpName() string {
	return "/eventlog/v0/ElfrReportEventAndSourceW"
}

func (o *xxx_ReportEventAndSourceWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Strings != nil && o.StringsLength == 0 {
		o.StringsLength = uint16(len(o.Strings))
	}
	if o.Data != nil && o.DataSize == 0 {
		o.DataSize = uint32(len(o.Data))
	}
	if o.StringsLength > uint16(256) {
		return fmt.Errorf("StringsLength is out of range")
	}
	if o.DataSize > uint32(61440) {
		return fmt.Errorf("DataSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventAndSourceWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Time {in} (1:(uint32))
	{
		if err := w.WriteData(o.Time); err != nil {
			return err
		}
	}
	// EventType {in} (1:(uint16))
	{
		if err := w.WriteData(o.EventType); err != nil {
			return err
		}
	}
	// EventCategory {in} (1:(uint16))
	{
		if err := w.WriteData(o.EventCategory); err != nil {
			return err
		}
	}
	// EventID {in} (1:(uint32))
	{
		if err := w.WriteData(o.EventID); err != nil {
			return err
		}
	}
	// SourceName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.SourceName != nil {
			if err := o.SourceName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// NumStrings {in} (1:{range=(0,256)}(uint16))
	{
		if err := w.WriteData(o.StringsLength); err != nil {
			return err
		}
	}
	// DataSize {in} (1:{range=(0,61440)}(uint32))
	{
		if err := w.WriteData(o.DataSize); err != nil {
			return err
		}
	}
	// ComputerName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ComputerName != nil {
			if err := o.ComputerName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// UserSID {in} (1:{pointer=unique, alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.UserSID != nil {
			_ptr_UserSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserSID != nil {
					if err := o.UserSID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserSID, _ptr_UserSID); err != nil {
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
	// Strings {in} (1:{pointer=unique}[dim:0,size_is=NumStrings])(2:{alias=PRPC_UNICODE_STRING}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		dimSize1 := uint64(o.StringsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Strings {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Strings[i1] != nil {
				_ptr_Strings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.Strings[i1] != nil {
						if err := o.Strings[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Strings[i1], _ptr_Strings); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Strings); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Data {in} (1:{pointer=unique}*(1)[dim:0,size_is=DataSize](uchar))
	{
		if o.Data != nil || o.DataSize > 0 {
			_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DataSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Data {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Data[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data, _ptr_Data); err != nil {
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
	// Flags {in} (1:(uint16))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TimeWritten {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_TimeWritten := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.TimeWritten); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TimeWritten, _ptr_TimeWritten); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventAndSourceWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Time {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Time); err != nil {
			return err
		}
	}
	// EventType {in} (1:(uint16))
	{
		if err := w.ReadData(&o.EventType); err != nil {
			return err
		}
	}
	// EventCategory {in} (1:(uint16))
	{
		if err := w.ReadData(&o.EventCategory); err != nil {
			return err
		}
	}
	// EventID {in} (1:(uint32))
	{
		if err := w.ReadData(&o.EventID); err != nil {
			return err
		}
	}
	// SourceName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.SourceName == nil {
			o.SourceName = &dtyp.UnicodeString{}
		}
		if err := o.SourceName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NumStrings {in} (1:{range=(0,256)}(uint16))
	{
		if err := w.ReadData(&o.StringsLength); err != nil {
			return err
		}
	}
	// DataSize {in} (1:{range=(0,61440)}(uint32))
	{
		if err := w.ReadData(&o.DataSize); err != nil {
			return err
		}
	}
	// ComputerName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ComputerName == nil {
			o.ComputerName = &dtyp.UnicodeString{}
		}
		if err := o.ComputerName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UserSID {in} (1:{pointer=unique, alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		_ptr_UserSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserSID == nil {
				o.UserSID = &dtyp.SID{}
			}
			if err := o.UserSID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_UserSID := func(ptr interface{}) { o.UserSID = *ptr.(**dtyp.SID) }
		if err := w.ReadPointer(&o.UserSID, _s_UserSID, _ptr_UserSID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Strings {in} (1:{pointer=unique}[dim:0,size_is=NumStrings])(2:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Strings", sizeInfo[0])
		}
		o.Strings = make([]*dtyp.UnicodeString, sizeInfo[0])
		for i1 := range o.Strings {
			i1 := i1
			_ptr_Strings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Strings[i1] == nil {
					o.Strings[i1] = &dtyp.UnicodeString{}
				}
				if err := o.Strings[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Strings := func(ptr interface{}) { o.Strings[i1] = *ptr.(**dtyp.UnicodeString) }
			if err := w.ReadPointer(&o.Strings[i1], _s_Strings, _ptr_Strings); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Data {in} (1:{pointer=unique}*(1)[dim:0,size_is=DataSize](uchar))
	{
		_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
			}
			o.Data = make([]byte, sizeInfo[0])
			for i1 := range o.Data {
				i1 := i1
				if err := w.ReadData(&o.Data[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_Data := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Data, _s_Data, _ptr_Data); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Flags {in} (1:(uint16))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		_s_RecordNumber := func(ptr interface{}) { o.RecordNumber = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.RecordNumber, _s_RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TimeWritten {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_TimeWritten := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.TimeWritten); err != nil {
				return err
			}
			return nil
		})
		_s_TimeWritten := func(ptr interface{}) { o.TimeWritten = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.TimeWritten, _s_TimeWritten, _ptr_TimeWritten); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventAndSourceWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventAndSourceWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TimeWritten {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_TimeWritten := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.TimeWritten); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TimeWritten, _ptr_TimeWritten); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventAndSourceWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		_s_RecordNumber := func(ptr interface{}) { o.RecordNumber = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.RecordNumber, _s_RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TimeWritten {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_TimeWritten := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.TimeWritten); err != nil {
				return err
			}
			return nil
		})
		_s_TimeWritten := func(ptr interface{}) { o.TimeWritten = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.TimeWritten, _s_TimeWritten, _ptr_TimeWritten); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReportEventAndSourceWRequest structure represents the ElfrReportEventAndSourceW operation request
type ReportEventAndSourceWRequest struct {
	// LogHandle: Handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6. This handle MUST NOT be obtained via the ElfrOpenBELA
	// (section 3.1.4.2) method or the ElfrOpenBELW (section 3.1.4.1) method.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// Time: Time at which the event was generated by the event source (not the time at
	// which the event was logged). The time MUST be expressed as the number of seconds
	// since 00:00:00 on January 1, 1970 (UTC).
	Time uint32 `idl:"name:Time" json:"time"`
	// EventType: Type of the event, as specified in section 2.2.2.
	EventType uint16 `idl:"name:EventType" json:"event_type"`
	// EventCategory: Event category, as specified in section 1.8.5.
	EventCategory uint16 `idl:"name:EventCategory" json:"event_category"`
	// EventID: EventID, as specified in section 3.1.1.4.
	EventID uint32 `idl:"name:EventID" json:"event_id"`
	// SourceName: Specifies the name of the event source.
	SourceName *dtyp.UnicodeString `idl:"name:SourceName" json:"source_name"`
	// NumStrings: Number of strings in the array pointed to by the Strings parameter. If
	// no strings are present, this value MUST be set to zero.
	StringsLength uint16 `idl:"name:NumStrings" json:"strings_length"`
	// DataSize: Number of bytes of event-specific raw binary data to write to the log.
	// This binary data is passed in the Data parameter. If no event-specific data is present,
	// this parameter MUST be set to zero.
	DataSize uint32 `idl:"name:DataSize" json:"data_size"`
	// ComputerName: A string to assist in identifying the machine that generated the event.
	// In practice, the name of the computer. There are no character restrictions on this
	// field's content (for example, a FQDN can be used).
	ComputerName *dtyp.UnicodeString `idl:"name:ComputerName" json:"computer_name"`
	// UserSID: Either NULL or a user SID. If this is NULL, the event is to have a zero
	// length UserSid field.
	UserSID *dtyp.SID `idl:"name:UserSID;pointer:unique" json:"user_sid"`
	// Strings: Strings containing text information specific to the event. This parameter
	// MUST be a valid pointer. If the NumStrings parameter is zero, this parameter MUST
	// be NULL. For example, an event relating to file deletion could use a string to specify
	// the path of the file being deleted.
	Strings []*dtyp.UnicodeString `idl:"name:Strings;size_is:(NumStrings);pointer:unique" json:"strings"`
	// Data: Pointer to a buffer that contains binary information specific to the event.
	// This parameter MUST be a valid pointer (or NULL), even if the DataSize parameter
	// is zero.
	Data []byte `idl:"name:Data;size_is:(DataSize);pointer:unique" json:"data"`
	// Flags: Unused. MUST be set to zero when sent and MUST be ignored on receipt.
	Flags uint16 `idl:"name:Flags" json:"flags"`
	// RecordNumber: Unused. Can be set to any arbitrary value when sent, and any value
	// sent by the client MUST be ignored on receipt by the server.
	RecordNumber uint32 `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	// TimeWritten: Unused. Can be set to any arbitrary value when sent, and any value sent
	// by the client MUST be ignored on receipt by the server.
	TimeWritten uint32 `idl:"name:TimeWritten;pointer:unique" json:"time_written"`
}

func (o *ReportEventAndSourceWRequest) xxx_ToOp(ctx context.Context, op *xxx_ReportEventAndSourceWOperation) *xxx_ReportEventAndSourceWOperation {
	if op == nil {
		op = &xxx_ReportEventAndSourceWOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.Time = o.Time
	op.EventType = o.EventType
	op.EventCategory = o.EventCategory
	op.EventID = o.EventID
	op.SourceName = o.SourceName
	op.StringsLength = o.StringsLength
	op.DataSize = o.DataSize
	op.ComputerName = o.ComputerName
	op.UserSID = o.UserSID
	op.Strings = o.Strings
	op.Data = o.Data
	op.Flags = o.Flags
	op.RecordNumber = o.RecordNumber
	op.TimeWritten = o.TimeWritten
	return op
}

func (o *ReportEventAndSourceWRequest) xxx_FromOp(ctx context.Context, op *xxx_ReportEventAndSourceWOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.Time = op.Time
	o.EventType = op.EventType
	o.EventCategory = op.EventCategory
	o.EventID = op.EventID
	o.SourceName = op.SourceName
	o.StringsLength = op.StringsLength
	o.DataSize = op.DataSize
	o.ComputerName = op.ComputerName
	o.UserSID = op.UserSID
	o.Strings = op.Strings
	o.Data = op.Data
	o.Flags = op.Flags
	o.RecordNumber = op.RecordNumber
	o.TimeWritten = op.TimeWritten
}
func (o *ReportEventAndSourceWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReportEventAndSourceWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportEventAndSourceWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReportEventAndSourceWResponse structure represents the ElfrReportEventAndSourceW operation response
type ReportEventAndSourceWResponse struct {
	// RecordNumber: Unused. Can be set to any arbitrary value when sent, and any value
	// sent by the client MUST be ignored on receipt by the server.
	RecordNumber uint32 `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	// TimeWritten: Unused. Can be set to any arbitrary value when sent, and any value sent
	// by the client MUST be ignored on receipt by the server.
	TimeWritten uint32 `idl:"name:TimeWritten;pointer:unique" json:"time_written"`
	// Return: The ElfrReportEventAndSourceW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReportEventAndSourceWResponse) xxx_ToOp(ctx context.Context, op *xxx_ReportEventAndSourceWOperation) *xxx_ReportEventAndSourceWOperation {
	if op == nil {
		op = &xxx_ReportEventAndSourceWOperation{}
	}
	if o == nil {
		return op
	}
	op.RecordNumber = o.RecordNumber
	op.TimeWritten = o.TimeWritten
	op.Return = o.Return
	return op
}

func (o *ReportEventAndSourceWResponse) xxx_FromOp(ctx context.Context, op *xxx_ReportEventAndSourceWOperation) {
	if o == nil {
		return
	}
	o.RecordNumber = op.RecordNumber
	o.TimeWritten = op.TimeWritten
	o.Return = op.Return
}
func (o *ReportEventAndSourceWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReportEventAndSourceWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportEventAndSourceWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReportEventExWOperation structure represents the ElfrReportEventExW operation
type xxx_ReportEventExWOperation struct {
	Log           *Handle               `idl:"name:LogHandle" json:"log"`
	TimeGenerated *dtyp.Filetime        `idl:"name:TimeGenerated" json:"time_generated"`
	EventType     uint16                `idl:"name:EventType" json:"event_type"`
	EventCategory uint16                `idl:"name:EventCategory" json:"event_category"`
	EventID       uint32                `idl:"name:EventID" json:"event_id"`
	StringsLength uint16                `idl:"name:NumStrings" json:"strings_length"`
	DataSize      uint32                `idl:"name:DataSize" json:"data_size"`
	ComputerName  *dtyp.UnicodeString   `idl:"name:ComputerName" json:"computer_name"`
	UserSID       *dtyp.SID             `idl:"name:UserSID;pointer:unique" json:"user_sid"`
	Strings       []*dtyp.UnicodeString `idl:"name:Strings;size_is:(NumStrings);pointer:unique" json:"strings"`
	Data          []byte                `idl:"name:Data;size_is:(DataSize);pointer:unique" json:"data"`
	Flags         uint16                `idl:"name:Flags" json:"flags"`
	RecordNumber  uint32                `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	Return        int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_ReportEventExWOperation) OpNum() int { return 25 }

func (o *xxx_ReportEventExWOperation) OpName() string { return "/eventlog/v0/ElfrReportEventExW" }

func (o *xxx_ReportEventExWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Strings != nil && o.StringsLength == 0 {
		o.StringsLength = uint16(len(o.Strings))
	}
	if o.Data != nil && o.DataSize == 0 {
		o.DataSize = uint32(len(o.Data))
	}
	if o.StringsLength > uint16(256) {
		return fmt.Errorf("StringsLength is out of range")
	}
	if o.DataSize > uint32(61440) {
		return fmt.Errorf("DataSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventExWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TimeGenerated {in} (1:{alias=PFILETIME}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.TimeGenerated != nil {
			if err := o.TimeGenerated.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// EventType {in} (1:(uint16))
	{
		if err := w.WriteData(o.EventType); err != nil {
			return err
		}
	}
	// EventCategory {in} (1:(uint16))
	{
		if err := w.WriteData(o.EventCategory); err != nil {
			return err
		}
	}
	// EventID {in} (1:(uint32))
	{
		if err := w.WriteData(o.EventID); err != nil {
			return err
		}
	}
	// NumStrings {in} (1:{range=(0,256)}(uint16))
	{
		if err := w.WriteData(o.StringsLength); err != nil {
			return err
		}
	}
	// DataSize {in} (1:{range=(0,61440)}(uint32))
	{
		if err := w.WriteData(o.DataSize); err != nil {
			return err
		}
	}
	// ComputerName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ComputerName != nil {
			if err := o.ComputerName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// UserSID {in} (1:{pointer=unique, alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.UserSID != nil {
			_ptr_UserSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserSID != nil {
					if err := o.UserSID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserSID, _ptr_UserSID); err != nil {
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
	// Strings {in} (1:{pointer=unique}[dim:0,size_is=NumStrings])(2:{alias=PRPC_UNICODE_STRING}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		dimSize1 := uint64(o.StringsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Strings {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Strings[i1] != nil {
				_ptr_Strings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.Strings[i1] != nil {
						if err := o.Strings[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Strings[i1], _ptr_Strings); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Strings); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Data {in} (1:{pointer=unique}*(1)[dim:0,size_is=DataSize](uchar))
	{
		if o.Data != nil || o.DataSize > 0 {
			_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DataSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Data {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Data[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data, _ptr_Data); err != nil {
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
	// Flags {in} (1:(uint16))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventExWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TimeGenerated {in} (1:{alias=PFILETIME,pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.TimeGenerated == nil {
			o.TimeGenerated = &dtyp.Filetime{}
		}
		if err := o.TimeGenerated.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// EventType {in} (1:(uint16))
	{
		if err := w.ReadData(&o.EventType); err != nil {
			return err
		}
	}
	// EventCategory {in} (1:(uint16))
	{
		if err := w.ReadData(&o.EventCategory); err != nil {
			return err
		}
	}
	// EventID {in} (1:(uint32))
	{
		if err := w.ReadData(&o.EventID); err != nil {
			return err
		}
	}
	// NumStrings {in} (1:{range=(0,256)}(uint16))
	{
		if err := w.ReadData(&o.StringsLength); err != nil {
			return err
		}
	}
	// DataSize {in} (1:{range=(0,61440)}(uint32))
	{
		if err := w.ReadData(&o.DataSize); err != nil {
			return err
		}
	}
	// ComputerName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ComputerName == nil {
			o.ComputerName = &dtyp.UnicodeString{}
		}
		if err := o.ComputerName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UserSID {in} (1:{pointer=unique, alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		_ptr_UserSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserSID == nil {
				o.UserSID = &dtyp.SID{}
			}
			if err := o.UserSID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_UserSID := func(ptr interface{}) { o.UserSID = *ptr.(**dtyp.SID) }
		if err := w.ReadPointer(&o.UserSID, _s_UserSID, _ptr_UserSID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Strings {in} (1:{pointer=unique}[dim:0,size_is=NumStrings])(2:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Strings", sizeInfo[0])
		}
		o.Strings = make([]*dtyp.UnicodeString, sizeInfo[0])
		for i1 := range o.Strings {
			i1 := i1
			_ptr_Strings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Strings[i1] == nil {
					o.Strings[i1] = &dtyp.UnicodeString{}
				}
				if err := o.Strings[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Strings := func(ptr interface{}) { o.Strings[i1] = *ptr.(**dtyp.UnicodeString) }
			if err := w.ReadPointer(&o.Strings[i1], _s_Strings, _ptr_Strings); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Data {in} (1:{pointer=unique}*(1)[dim:0,size_is=DataSize](uchar))
	{
		_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
			}
			o.Data = make([]byte, sizeInfo[0])
			for i1 := range o.Data {
				i1 := i1
				if err := w.ReadData(&o.Data[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_Data := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Data, _s_Data, _ptr_Data); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Flags {in} (1:(uint16))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		_s_RecordNumber := func(ptr interface{}) { o.RecordNumber = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.RecordNumber, _s_RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventExWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventExWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventExWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		_s_RecordNumber := func(ptr interface{}) { o.RecordNumber = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.RecordNumber, _s_RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReportEventExWRequest structure represents the ElfrReportEventExW operation request
type ReportEventExWRequest struct {
	// LogHandle: A handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6. This handle MUST NOT be obtained via the ElfrOpenBELA
	// (section 3.1.4.2) method or the ElfrOpenBELW (section 3.1.4.1) method. A handle received
	// from either of those two methods will have the backup flag set, so the server checks
	// this flag before calling this method.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// TimeGenerated: The time at which the event was generated by the event source. This
	// time is represented as a pointer to FILETIME as defined in [MS-DTYP] section 2.3.3.<34>
	TimeGenerated *dtyp.Filetime `idl:"name:TimeGenerated" json:"time_generated"`
	// EventType: The type of event, as specified in section 2.2.2.
	EventType uint16 `idl:"name:EventType" json:"event_type"`
	// EventCategory: The event category, as specified in section 1.8.5.
	EventCategory uint16 `idl:"name:EventCategory" json:"event_category"`
	// EventID: The EventID, as specified in section 3.1.1.4.
	EventID uint32 `idl:"name:EventID" json:"event_id"`
	// NumStrings: The number of strings in the array pointed to by the Strings parameter.
	// A value of zero indicates that no strings are present.
	StringsLength uint16 `idl:"name:NumStrings" json:"strings_length"`
	// DataSize: The number of bytes of event-specific raw binary data to write to the log.
	// This binary data is passed in the Data parameter. If the DataSize parameter is zero,
	// event-specific data MUST NOT be present.
	DataSize uint32 `idl:"name:DataSize" json:"data_size"`
	// ComputerName: A string to assist in identifying the machine that generated the event;
	// for example, the name of the computer. There are no character restrictions on this
	// field's content (for example, a FQDN can be used). The API is not intended to support
	// dynamically changing computer names. The ComputerName parameter is cached the first
	// time a client calls the API, and that name used on subsequent calls until the machine
	// is rebooted.
	ComputerName *dtyp.UnicodeString `idl:"name:ComputerName" json:"computer_name"`
	// UserSID: Either NULL or a user SID. If this is NULL, the event is to have a zero
	// length UserSid field.
	UserSID *dtyp.SID `idl:"name:UserSID;pointer:unique" json:"user_sid"`
	// Strings: Specifies the strings containing information specific to the event. This
	// parameter MUST be a valid pointer. If the NumStrings parameter is zero, this parameter
	// MUST be NULL. For example, an event relating to file deletion could use a string
	// to specify the path of the file being deleted.
	Strings []*dtyp.UnicodeString `idl:"name:Strings;size_is:(NumStrings);pointer:unique" json:"strings"`
	// Data: A pointer to the buffer that contains the event-specific binary data. This
	// parameter MUST be a valid pointer (or NULL), even if the DataSize parameter is zero.
	Data []byte `idl:"name:Data;size_is:(DataSize);pointer:unique" json:"data"`
	// Flags: Unused. MUST be set to zero when sent and MUST be ignored on receipt.
	Flags uint16 `idl:"name:Flags" json:"flags"`
	// RecordNumber: Unused. Can be set to any arbitrary value when sent, and any value
	// sent by the client MUST be ignored on receipt by the server.
	RecordNumber uint32 `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
}

func (o *ReportEventExWRequest) xxx_ToOp(ctx context.Context, op *xxx_ReportEventExWOperation) *xxx_ReportEventExWOperation {
	if op == nil {
		op = &xxx_ReportEventExWOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.TimeGenerated = o.TimeGenerated
	op.EventType = o.EventType
	op.EventCategory = o.EventCategory
	op.EventID = o.EventID
	op.StringsLength = o.StringsLength
	op.DataSize = o.DataSize
	op.ComputerName = o.ComputerName
	op.UserSID = o.UserSID
	op.Strings = o.Strings
	op.Data = o.Data
	op.Flags = o.Flags
	op.RecordNumber = o.RecordNumber
	return op
}

func (o *ReportEventExWRequest) xxx_FromOp(ctx context.Context, op *xxx_ReportEventExWOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.TimeGenerated = op.TimeGenerated
	o.EventType = op.EventType
	o.EventCategory = op.EventCategory
	o.EventID = op.EventID
	o.StringsLength = op.StringsLength
	o.DataSize = op.DataSize
	o.ComputerName = op.ComputerName
	o.UserSID = op.UserSID
	o.Strings = op.Strings
	o.Data = op.Data
	o.Flags = op.Flags
	o.RecordNumber = op.RecordNumber
}
func (o *ReportEventExWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReportEventExWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportEventExWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReportEventExWResponse structure represents the ElfrReportEventExW operation response
type ReportEventExWResponse struct {
	// RecordNumber: Unused. Can be set to any arbitrary value when sent, and any value
	// sent by the client MUST be ignored on receipt by the server.
	RecordNumber uint32 `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	// Return: The ElfrReportEventExW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReportEventExWResponse) xxx_ToOp(ctx context.Context, op *xxx_ReportEventExWOperation) *xxx_ReportEventExWOperation {
	if op == nil {
		op = &xxx_ReportEventExWOperation{}
	}
	if o == nil {
		return op
	}
	op.RecordNumber = o.RecordNumber
	op.Return = o.Return
	return op
}

func (o *ReportEventExWResponse) xxx_FromOp(ctx context.Context, op *xxx_ReportEventExWOperation) {
	if o == nil {
		return
	}
	o.RecordNumber = op.RecordNumber
	o.Return = op.Return
}
func (o *ReportEventExWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReportEventExWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportEventExWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReportEventExAOperation structure represents the ElfrReportEventExA operation
type xxx_ReportEventExAOperation struct {
	Log           *Handle        `idl:"name:LogHandle" json:"log"`
	TimeGenerated *dtyp.Filetime `idl:"name:TimeGenerated" json:"time_generated"`
	EventType     uint16         `idl:"name:EventType" json:"event_type"`
	EventCategory uint16         `idl:"name:EventCategory" json:"event_category"`
	EventID       uint32         `idl:"name:EventID" json:"event_id"`
	StringsLength uint16         `idl:"name:NumStrings" json:"strings_length"`
	DataSize      uint32         `idl:"name:DataSize" json:"data_size"`
	ComputerName  *String        `idl:"name:ComputerName" json:"computer_name"`
	UserSID       *dtyp.SID      `idl:"name:UserSID;pointer:unique" json:"user_sid"`
	Strings       []*String      `idl:"name:Strings;size_is:(NumStrings);pointer:unique" json:"strings"`
	Data          []byte         `idl:"name:Data;size_is:(DataSize);pointer:unique" json:"data"`
	Flags         uint16         `idl:"name:Flags" json:"flags"`
	RecordNumber  uint32         `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReportEventExAOperation) OpNum() int { return 26 }

func (o *xxx_ReportEventExAOperation) OpName() string { return "/eventlog/v0/ElfrReportEventExA" }

func (o *xxx_ReportEventExAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Strings != nil && o.StringsLength == 0 {
		o.StringsLength = uint16(len(o.Strings))
	}
	if o.Data != nil && o.DataSize == 0 {
		o.DataSize = uint32(len(o.Data))
	}
	if o.StringsLength > uint16(256) {
		return fmt.Errorf("StringsLength is out of range")
	}
	if o.DataSize > uint32(61440) {
		return fmt.Errorf("DataSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventExAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TimeGenerated {in} (1:{alias=PFILETIME}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.TimeGenerated != nil {
			if err := o.TimeGenerated.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// EventType {in} (1:(uint16))
	{
		if err := w.WriteData(o.EventType); err != nil {
			return err
		}
	}
	// EventCategory {in} (1:(uint16))
	{
		if err := w.WriteData(o.EventCategory); err != nil {
			return err
		}
	}
	// EventID {in} (1:(uint32))
	{
		if err := w.WriteData(o.EventID); err != nil {
			return err
		}
	}
	// NumStrings {in} (1:{range=(0,256)}(uint16))
	{
		if err := w.WriteData(o.StringsLength); err != nil {
			return err
		}
	}
	// DataSize {in} (1:{range=(0,61440)}(uint32))
	{
		if err := w.WriteData(o.DataSize); err != nil {
			return err
		}
	}
	// ComputerName {in} (1:{alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.ComputerName != nil {
			if err := o.ComputerName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// UserSID {in} (1:{pointer=unique, alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.UserSID != nil {
			_ptr_UserSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserSID != nil {
					if err := o.UserSID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserSID, _ptr_UserSID); err != nil {
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
	// Strings {in} (1:{pointer=unique}[dim:0,size_is=NumStrings])(2:{alias=PRPC_STRING}*(1))(3:{alias=RPC_STRING}(struct))
	{
		dimSize1 := uint64(o.StringsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Strings {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Strings[i1] != nil {
				_ptr_Strings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.Strings[i1] != nil {
						if err := o.Strings[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&String{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Strings[i1], _ptr_Strings); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Strings); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Data {in} (1:{pointer=unique}*(1)[dim:0,size_is=DataSize](uchar))
	{
		if o.Data != nil || o.DataSize > 0 {
			_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DataSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Data {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Data[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data, _ptr_Data); err != nil {
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
	// Flags {in} (1:(uint16))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventExAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LogHandle {in} (1:{context_handle, alias=IELF_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Handle{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TimeGenerated {in} (1:{alias=PFILETIME,pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.TimeGenerated == nil {
			o.TimeGenerated = &dtyp.Filetime{}
		}
		if err := o.TimeGenerated.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// EventType {in} (1:(uint16))
	{
		if err := w.ReadData(&o.EventType); err != nil {
			return err
		}
	}
	// EventCategory {in} (1:(uint16))
	{
		if err := w.ReadData(&o.EventCategory); err != nil {
			return err
		}
	}
	// EventID {in} (1:(uint32))
	{
		if err := w.ReadData(&o.EventID); err != nil {
			return err
		}
	}
	// NumStrings {in} (1:{range=(0,256)}(uint16))
	{
		if err := w.ReadData(&o.StringsLength); err != nil {
			return err
		}
	}
	// DataSize {in} (1:{range=(0,61440)}(uint32))
	{
		if err := w.ReadData(&o.DataSize); err != nil {
			return err
		}
	}
	// ComputerName {in} (1:{alias=PRPC_STRING,pointer=ref}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.ComputerName == nil {
			o.ComputerName = &String{}
		}
		if err := o.ComputerName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UserSID {in} (1:{pointer=unique, alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		_ptr_UserSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserSID == nil {
				o.UserSID = &dtyp.SID{}
			}
			if err := o.UserSID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_UserSID := func(ptr interface{}) { o.UserSID = *ptr.(**dtyp.SID) }
		if err := w.ReadPointer(&o.UserSID, _s_UserSID, _ptr_UserSID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Strings {in} (1:{pointer=unique}[dim:0,size_is=NumStrings])(2:{alias=PRPC_STRING,pointer=ref}*(1))(3:{alias=RPC_STRING}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Strings", sizeInfo[0])
		}
		o.Strings = make([]*String, sizeInfo[0])
		for i1 := range o.Strings {
			i1 := i1
			_ptr_Strings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Strings[i1] == nil {
					o.Strings[i1] = &String{}
				}
				if err := o.Strings[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Strings := func(ptr interface{}) { o.Strings[i1] = *ptr.(**String) }
			if err := w.ReadPointer(&o.Strings[i1], _s_Strings, _ptr_Strings); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Data {in} (1:{pointer=unique}*(1)[dim:0,size_is=DataSize](uchar))
	{
		_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
			}
			o.Data = make([]byte, sizeInfo[0])
			for i1 := range o.Data {
				i1 := i1
				if err := w.ReadData(&o.Data[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_Data := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Data, _s_Data, _ptr_Data); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Flags {in} (1:(uint16))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		_s_RecordNumber := func(ptr interface{}) { o.RecordNumber = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.RecordNumber, _s_RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventExAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventExAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportEventExAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// RecordNumber {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_RecordNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.RecordNumber); err != nil {
				return err
			}
			return nil
		})
		_s_RecordNumber := func(ptr interface{}) { o.RecordNumber = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.RecordNumber, _s_RecordNumber, _ptr_RecordNumber); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReportEventExARequest structure represents the ElfrReportEventExA operation request
type ReportEventExARequest struct {
	// LogHandle: A handle to an event log. This parameter is a server context handle, as
	// specified in section 2.2.6. This handle MUST NOT be obtained via the ElfrOpenBELA
	// (section 3.1.4.2) method or the ElfrOpenBELW (section 3.1.4.1) method. A handle received
	// from either of those two methods will have the backup flag set, so the server checks
	// this flag before calling this method.
	Log *Handle `idl:"name:LogHandle" json:"log"`
	// TimeGenerated: The time at which the event was generated by the event source. This
	// time is represented as a pointer to FILETIME as defined in [MS-DTYP] section 2.3.3.
	TimeGenerated *dtyp.Filetime `idl:"name:TimeGenerated" json:"time_generated"`
	// EventType: The type of the event, as specified in section 2.2.2.
	EventType uint16 `idl:"name:EventType" json:"event_type"`
	// EventCategory: The event category, as specified in section 1.8.5.
	EventCategory uint16 `idl:"name:EventCategory" json:"event_category"`
	// EventID: The EventID, as specified in section 3.1.1.4.
	EventID uint32 `idl:"name:EventID" json:"event_id"`
	// NumStrings: The number of strings in the array pointed to by the Strings parameter.
	// A value of zero indicates that no strings are present.
	StringsLength uint16 `idl:"name:NumStrings" json:"strings_length"`
	// DataSize: The number of bytes of event-specific raw binary data to write to the log.
	// This binary data is passed in the Data parameter. If the DataSize parameter is zero,
	// event-specific data MUST NOT be present.
	DataSize uint32 `idl:"name:DataSize" json:"data_size"`
	// ComputerName: A string to assist in identifying the machine that generated the event.
	// In practice, the name of the computer. There are no character restrictions on this
	// field's content (for example, a FQDN can be used). The API is not intended to support
	// dynamically changing computer names. The ComputerName parameter is cached the first
	// time a client calls the API, and SHOULD use that name on subsequent calls until the
	// machine is rebooted.
	ComputerName *String `idl:"name:ComputerName" json:"computer_name"`
	// UserSID: Either NULL or a user SID. If this is NULL, the event is to have a zero
	// length UserSid field.
	UserSID *dtyp.SID `idl:"name:UserSID;pointer:unique" json:"user_sid"`
	// Strings: Specifies strings containing information specific to the event. This parameter
	// MUST be a valid pointer. If the NumStrings parameter is zero, this parameter MUST
	// be NULL. For example, an event relating to file deletion could use a string to specify
	// the path of the file being deleted.
	Strings []*String `idl:"name:Strings;size_is:(NumStrings);pointer:unique" json:"strings"`
	// Data: A pointer to the buffer that contains the event-specific binary data. This
	// parameter MUST be a valid pointer (or NULL), even if the DataSize parameter is zero.
	Data []byte `idl:"name:Data;size_is:(DataSize);pointer:unique" json:"data"`
	// Flags: Unused. MUST be set to zero when sent and MUST be ignored on receipt.
	Flags uint16 `idl:"name:Flags" json:"flags"`
	// RecordNumber: Unused. This can be set to any arbitrary value when sent. Any value
	// sent by the client MUST be ignored on receipt by the server.
	RecordNumber uint32 `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
}

func (o *ReportEventExARequest) xxx_ToOp(ctx context.Context, op *xxx_ReportEventExAOperation) *xxx_ReportEventExAOperation {
	if op == nil {
		op = &xxx_ReportEventExAOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.TimeGenerated = o.TimeGenerated
	op.EventType = o.EventType
	op.EventCategory = o.EventCategory
	op.EventID = o.EventID
	op.StringsLength = o.StringsLength
	op.DataSize = o.DataSize
	op.ComputerName = o.ComputerName
	op.UserSID = o.UserSID
	op.Strings = o.Strings
	op.Data = o.Data
	op.Flags = o.Flags
	op.RecordNumber = o.RecordNumber
	return op
}

func (o *ReportEventExARequest) xxx_FromOp(ctx context.Context, op *xxx_ReportEventExAOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.TimeGenerated = op.TimeGenerated
	o.EventType = op.EventType
	o.EventCategory = op.EventCategory
	o.EventID = op.EventID
	o.StringsLength = op.StringsLength
	o.DataSize = op.DataSize
	o.ComputerName = op.ComputerName
	o.UserSID = op.UserSID
	o.Strings = op.Strings
	o.Data = op.Data
	o.Flags = op.Flags
	o.RecordNumber = op.RecordNumber
}
func (o *ReportEventExARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReportEventExARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportEventExAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReportEventExAResponse structure represents the ElfrReportEventExA operation response
type ReportEventExAResponse struct {
	// RecordNumber: Unused. This can be set to any arbitrary value when sent. Any value
	// sent by the client MUST be ignored on receipt by the server.
	RecordNumber uint32 `idl:"name:RecordNumber;pointer:unique" json:"record_number"`
	// Return: The ElfrReportEventExA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReportEventExAResponse) xxx_ToOp(ctx context.Context, op *xxx_ReportEventExAOperation) *xxx_ReportEventExAOperation {
	if op == nil {
		op = &xxx_ReportEventExAOperation{}
	}
	if o == nil {
		return op
	}
	op.RecordNumber = o.RecordNumber
	op.Return = o.Return
	return op
}

func (o *ReportEventExAResponse) xxx_FromOp(ctx context.Context, op *xxx_ReportEventExAOperation) {
	if o == nil {
		return
	}
	o.RecordNumber = op.RecordNumber
	o.Return = op.Return
}
func (o *ReportEventExAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReportEventExAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportEventExAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
