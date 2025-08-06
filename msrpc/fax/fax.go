// The fax package implements the FAX client protocol.
//
// # Introduction
//
// The Fax Server and Client Remote Protocol Specification defines a protocol that is
// referred to as the Fax Server and Client Remote Protocol. This is a client/server
// protocol based on remote procedure call (RPC) that is used to send faxes and manage
// the fax server and its queues.
//
// # Overview
//
// The Fax Server and Client Remote Protocol manages and sends faxes, manages the fax
// server and its queues, and allows fax clients to act as RPC servers so that they
// can accept status notifications from fax servers acting as clients.
package fax

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
	GoPackage = "fax"
)

// MaxDevicesInGroup represents the FAX_MAX_DEVICES_IN_GROUP RPC constant
var MaxDevicesInGroup = 1000

// MaxRPCBuffer represents the FAX_MAX_RPC_BUFFER RPC constant
var MaxRPCBuffer = 1048576

// MaxRecipients represents the FAX_MAX_RECIPIENTS RPC constant
var MaxRecipients = 10000

// CopyBufferSize represents the RPC_COPY_BUFFER_SIZE RPC constant
var CopyBufferSize = 16384

// Fax structure represents RPC_FAX_HANDLE RPC structure.
type Fax dcetypes.ContextHandle

func (o *Fax) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Fax) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Fax) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Fax) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Port structure represents RPC_FAX_PORT_HANDLE RPC structure.
type Port dcetypes.ContextHandle

func (o *Port) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Port) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Port) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Port) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Service structure represents RPC_FAX_SVC_HANDLE RPC structure.
type Service dcetypes.ContextHandle

func (o *Service) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Service) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Service) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Service) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MessageEnum structure represents RPC_FAX_MSG_ENUM_HANDLE RPC structure.
type MessageEnum dcetypes.ContextHandle

func (o *MessageEnum) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *MessageEnum) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MessageEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MessageEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Copy structure represents RPC_FAX_COPY_HANDLE RPC structure.
type Copy dcetypes.ContextHandle

func (o *Copy) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Copy) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Copy) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Copy) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Event structure represents RPC_FAX_EVENT_HANDLE RPC structure.
type Event dcetypes.ContextHandle

func (o *Event) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Event) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Event) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Event) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EventEx structure represents RPC_FAX_EVENT_EX_HANDLE RPC structure.
type EventEx dcetypes.ContextHandle

func (o *EventEx) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *EventEx) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EventEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EventEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CoverPageInfoExW structure represents FAX_COVERPAGE_INFO_EXW RPC structure.
//
// The FAX_COVERPAGE_INFO_EXW structure is used as an argument for the FAX_SendDocumentEx
// (section 3.1.4.1.73) call that specifies information about the fax cover page used
// when sending a fax.
type CoverPageInfoExW struct {
	// dwSizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) value that holds the total size
	// of the structure, in bytes. This value MUST be 24 or 40 bytes. When filled in on
	// a 32-bit implementation, this value SHOULD be 24 bytes. When filled in on a 64-bit
	// implementation, this value SHOULD be 40 bytes.
	StructureSize uint32 `idl:"name:dwSizeOfStruct" json:"structure_size"`
	// dwCoverPageFormat: A DWORD that indicates the format of the cover page template.
	// This MUST be one of the values defined in FAX_ENUM_COVERPAGE_FORMATS (section 2.2.78)
	// enumeration. The required file format for the cover page template is described in
	// FAX_SendDocumentEx (section 3.1.4.1.73) method.
	CoverPageFormat uint32 `idl:"name:dwCoverPageFormat" json:"cover_page_format"`
	// lpwstrCoverPageFileName: A pointer to a null-terminated character string that holds
	// the file name of the cover page template. This file name SHOULD NOT include any path
	// separators. If bServerBased is FALSE, the file extension MUST be ".cov", and except
	// for the terminating null character, the character string MUST contain only characters
	// representing valid hexadecimal digits: "0123456789abcdefABCDEF". If bServerBased
	// is TRUE the file extension SHOULD be ".cov". The cover page file MUST be present
	// in the fax server queue directory when the FAX_SendDocumentEx call is made. If no
	// cover-page information is available, this pointer MUST be NULL.
	CoverPageFileName string `idl:"name:lpwstrCoverPageFileName;string" json:"cover_page_file_name"`
	// bServerBased: A Boolean that indicates whether the cover page template specified
	// by the lpwstrCoverPageFileName parameter is a new personal cover page template (when
	// set to FALSE)  or a server-based cover page template (when set to TRUE). For more
	// details on the semantics of TRUE and FALSE, see FAX_SendDocumentEx.
	ServerBased bool `idl:"name:bServerBased" json:"server_based"`
	// lpwstrNote: A pointer to a null-terminated character string that holds the content
	// for the note field of the cover page.
	Note string `idl:"name:lpwstrNote;string" json:"note"`
	// lpwstrSubject: A pointer to a null-terminated character string that holds the content
	// for the subject field.
	Subject string `idl:"name:lpwstrSubject;string" json:"subject"`
}

func (o *CoverPageInfoExW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *CoverPageInfoExW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if err := w.WriteData(o.CoverPageFormat); err != nil {
		return err
	}
	if o.CoverPageFileName != "" {
		_ptr_lpwstrCoverPageFileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.CoverPageFileName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.CoverPageFileName, _ptr_lpwstrCoverPageFileName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.ServerBased {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.Note != "" {
		_ptr_lpwstrNote := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Note); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Note, _ptr_lpwstrNote); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Subject != "" {
		_ptr_lpwstrSubject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Subject); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Subject, _ptr_lpwstrSubject); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CoverPageInfoExW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.CoverPageFormat); err != nil {
		return err
	}
	_ptr_lpwstrCoverPageFileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.CoverPageFileName); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrCoverPageFileName := func(ptr interface{}) { o.CoverPageFileName = *ptr.(*string) }
	if err := w.ReadPointer(&o.CoverPageFileName, _s_lpwstrCoverPageFileName, _ptr_lpwstrCoverPageFileName); err != nil {
		return err
	}
	var _bServerBased int32
	if err := w.ReadData(&_bServerBased); err != nil {
		return err
	}
	o.ServerBased = _bServerBased != 0
	_ptr_lpwstrNote := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Note); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrNote := func(ptr interface{}) { o.Note = *ptr.(*string) }
	if err := w.ReadPointer(&o.Note, _s_lpwstrNote, _ptr_lpwstrNote); err != nil {
		return err
	}
	_ptr_lpwstrSubject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Subject); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrSubject := func(ptr interface{}) { o.Subject = *ptr.(*string) }
	if err := w.ReadPointer(&o.Subject, _s_lpwstrSubject, _ptr_lpwstrSubject); err != nil {
		return err
	}
	return nil
}

// JobParamW structure represents FAX_JOB_PARAMW RPC structure.
//
// The FAX_JOB_PARAMW structure contains information about a fax job, including information
// about the personal profiles (section 3.1.1) for the sender and the recipient of the
// fax. This structure is used as an input argument for the FaxObs_SendDocument (section
// 3.1.4.2.7) method.
type JobParamW struct {
	// SizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) that contains the size, in bytes,
	// of this structure. This value MUST be 80 or 136 bytes. When filled in on a 32-bit
	// implementation, this value SHOULD be 80 bytes. When filled in on a 64-bit implementation,
	// this value SHOULD be 136 bytes.
	StructureSize uint32 `idl:"name:SizeOfStruct" json:"structure_size"`
	// RecipientNumber: A null-terminated character string that holds the fax number of
	// the fax transmission recipient.
	RecipientNumber string `idl:"name:RecipientNumber;string" json:"recipient_number"`
	// RecipientName: A null-terminated character string that holds the name of the fax
	// transmission recipient.
	RecipientName string `idl:"name:RecipientName;string" json:"recipient_name"`
	// Tsid: A null-terminated character string that holds the transmitting subscriber identifier
	// (TSID). The valid characters for a TSID string are the English letters, the numeric
	// symbols, and the punctuation marks (ASCII range 0x20 to 0x7F).
	TSID string `idl:"name:Tsid;string" json:"tsid"`
	// SenderName: A null-terminated character string that holds the name of the fax transmission
	// sender.
	SenderName string `idl:"name:SenderName;string" json:"sender_name"`
	// SenderCompany: A null-terminated character string that holds the name of the fax
	// transmission sender's company.
	SenderCompany string `idl:"name:SenderCompany;string" json:"sender_company"`
	// SenderDept: A null-terminated character string that holds the name of the fax transmission
	// sender's department.
	SenderDept string `idl:"name:SenderDept;string" json:"sender_dept"`
	// BillingCode: A null-terminated character string that holds an optional billing code
	// for the fax transmission.
	BillingCode string `idl:"name:BillingCode;string" json:"billing_code"`
	// ScheduleAction: A DWORD variable that indicates when the fax is to be sent. This
	// value can be one of the following values:
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|                                |                                                                                  |
	//	|           VALUE/CODE           |                                     MEANING                                      |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| JSA_NOW 0x00000000             | The fax is to be sent as soon as a fax device is available.                      |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| JSA_SPECIFIC_TIME 0x00000001   | The fax is to be sent at the time specified by the ScheduleTime member of this   |
	//	|                                | structure.                                                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| JSA_DISCOUNT_PERIOD 0x00000002 | The fax is to be sent during the discount rate period. The                       |
	//	|                                | FaxObs_GetConfiguration (section 3.1.4.2.24) method can be called to retrieve    |
	//	|                                | the discount period for the fax server.                                          |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	ScheduleAction uint32 `idl:"name:ScheduleAction" json:"schedule_action"`
	// ScheduleTime: A SYSTEMTIME ([MS-DTYP] section 2.3.13) structure indicating the local
	// date and time to send the fax, in UTC format. This member is used when the ScheduleAction
	// member is set to 0x00000001 (JSA_SPECIFIC_TIME), and is otherwise ignored.
	ScheduleTime *dtyp.SystemTime `idl:"name:ScheduleTime" json:"schedule_time"`
	// DeliveryReportType: A DWORD variable that indicates the fax delivery report type.
	// This value can be one of the FAX_ENUM_DELIVERY_REPORT_TYPES (section 2.2.76) enumeration
	// values. The DRT_ATTACH_FAX value can be combined with the DRT_EMAIL value in one
	// value by using an OR operation.
	DeliveryReportType uint32 `idl:"name:DeliveryReportType" json:"delivery_report_type"`
	// DeliveryReportAddress: A null-terminated character string. Contains the email address
	// for the delivery report when the DeliveryReportType member is set to 0x00000001 (DRT_E_MAIL).
	// Otherwise, this pointer value can be NULL.
	DeliveryReportAddress string `idl:"name:DeliveryReportAddress;string" json:"delivery_report_address"`
	// DocumentName: A null-terminated character string that holds the document name. A
	// NULL pointer value specifies that no document name is specified for this fax job.
	DocumentName string `idl:"name:DocumentName;string" json:"document_name"`
	// CallHandle: An unsigned 32-bit integer value containing an optional TAPI call handle.
	// For more information about TAPI, see [MSDN-TAPI2.2]. For more information about this
	// member, see FaxObs_SendDocument.
	CallHandle uint32 `idl:"name:CallHandle" json:"call_handle"`
	// Reserved: A table of three 32-bit unsigned integer fields (on 32-bit implementations),
	// or 64-bit unsigned integer fields (on 64-bit implementations). If the first value,
	// Reserved[0], is zero, then all values in this table SHOULD be ignored.
	//
	// If the fax job is a normal job sent to one fax device (port), the Reserved values
	// SHOULD be as follows:
	//
	// * *Reserved[0]* SHOULD be set to zero or to 0xFFFFFFFF (on 32-bit) or 0x00000000FFFFFFFF
	// (on 64-bit).
	//
	// * *Reserved[1]* SHOULD contain a device identifier such as the value contained by
	// the *DeviceId* member of a valid *FAX_PORT_INFO* (section 2.2.7 ( 2b46d16c-e74a-4e3b-b50c-0b94f78bebd0
	// ) ) or *_FAX_PORT_INFO* (section 2.2.8 ( ed920746-d222-4e0a-a75d-905f26cf1dfc ) )
	// structure, describing one fax port (device).
	//
	// * *Reserved[2]* SHOULD be ignored.
	//
	// * *Reserved[0]* SHOULD be set to 0xFFFFFFFE (on 32-bit) or 0x00000000FFFFFFFE (on
	// 64-bit).
	//
	// * *Reserved[1]* SHOULD be set to one of the following two values:
	//
	// * A value of 1 (0x00000001 on 32-bit or 0x0000000000000001 on 64-bit) for the first
	// *FaxObs_SendDocument* method call made by the client to start the broadcast sequence.
	//
	// * A value of 2 (0x00000002 on 32-bit or 0x0000000000000002 on 64-bit) for the second
	// and following *FaxObs_SendDocument* method calls made by the client to continue and
	// complete a started broadcast sequence.
	//
	// * *Reserved[2]* SHOULD be set to one of the following two values:
	//
	// * If *Reserved[1]* is set to a value of 1, *Reserved[2]* SHOULD be set to zero.
	//
	// * If *Reserved[1]* is set to a value of 2, *Reserved[2]* SHOULD contain the job identifier
	// returned by the *FaxObs_SendDocument* call that started the broadcast sequence.
	//
	// For more information about this member, see the FaxObs_SendDocument method.
	_ []uint64 `idl:"name:Reserved"`
}

func (o *JobParamW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *JobParamW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if o.RecipientNumber != "" {
		_ptr_RecipientNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RecipientNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecipientNumber, _ptr_RecipientNumber); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RecipientName != "" {
		_ptr_RecipientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RecipientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecipientName, _ptr_RecipientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.TSID != "" {
		_ptr_Tsid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.TSID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TSID, _ptr_Tsid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SenderName != "" {
		_ptr_SenderName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SenderName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SenderName, _ptr_SenderName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SenderCompany != "" {
		_ptr_SenderCompany := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SenderCompany); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SenderCompany, _ptr_SenderCompany); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SenderDept != "" {
		_ptr_SenderDept := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SenderDept); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SenderDept, _ptr_SenderDept); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.BillingCode != "" {
		_ptr_BillingCode := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.BillingCode); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.BillingCode, _ptr_BillingCode); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ScheduleAction); err != nil {
		return err
	}
	if o.ScheduleTime != nil {
		if err := o.ScheduleTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DeliveryReportType); err != nil {
		return err
	}
	if o.DeliveryReportAddress != "" {
		_ptr_DeliveryReportAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DeliveryReportAddress); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DeliveryReportAddress, _ptr_DeliveryReportAddress); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DocumentName != "" {
		_ptr_DocumentName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DocumentName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DocumentName, _ptr_DocumentName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.CallHandle); err != nil {
		return err
	}
	// reserved Reserved
	for i1 := 0; uint64(i1) < 3; i1++ {
		if err := w.WriteData(uint64(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *JobParamW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	_ptr_RecipientNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RecipientNumber); err != nil {
			return err
		}
		return nil
	})
	_s_RecipientNumber := func(ptr interface{}) { o.RecipientNumber = *ptr.(*string) }
	if err := w.ReadPointer(&o.RecipientNumber, _s_RecipientNumber, _ptr_RecipientNumber); err != nil {
		return err
	}
	_ptr_RecipientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RecipientName); err != nil {
			return err
		}
		return nil
	})
	_s_RecipientName := func(ptr interface{}) { o.RecipientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.RecipientName, _s_RecipientName, _ptr_RecipientName); err != nil {
		return err
	}
	_ptr_Tsid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.TSID); err != nil {
			return err
		}
		return nil
	})
	_s_Tsid := func(ptr interface{}) { o.TSID = *ptr.(*string) }
	if err := w.ReadPointer(&o.TSID, _s_Tsid, _ptr_Tsid); err != nil {
		return err
	}
	_ptr_SenderName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SenderName); err != nil {
			return err
		}
		return nil
	})
	_s_SenderName := func(ptr interface{}) { o.SenderName = *ptr.(*string) }
	if err := w.ReadPointer(&o.SenderName, _s_SenderName, _ptr_SenderName); err != nil {
		return err
	}
	_ptr_SenderCompany := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SenderCompany); err != nil {
			return err
		}
		return nil
	})
	_s_SenderCompany := func(ptr interface{}) { o.SenderCompany = *ptr.(*string) }
	if err := w.ReadPointer(&o.SenderCompany, _s_SenderCompany, _ptr_SenderCompany); err != nil {
		return err
	}
	_ptr_SenderDept := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SenderDept); err != nil {
			return err
		}
		return nil
	})
	_s_SenderDept := func(ptr interface{}) { o.SenderDept = *ptr.(*string) }
	if err := w.ReadPointer(&o.SenderDept, _s_SenderDept, _ptr_SenderDept); err != nil {
		return err
	}
	_ptr_BillingCode := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.BillingCode); err != nil {
			return err
		}
		return nil
	})
	_s_BillingCode := func(ptr interface{}) { o.BillingCode = *ptr.(*string) }
	if err := w.ReadPointer(&o.BillingCode, _s_BillingCode, _ptr_BillingCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScheduleAction); err != nil {
		return err
	}
	if o.ScheduleTime == nil {
		o.ScheduleTime = &dtyp.SystemTime{}
	}
	if err := o.ScheduleTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeliveryReportType); err != nil {
		return err
	}
	_ptr_DeliveryReportAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DeliveryReportAddress); err != nil {
			return err
		}
		return nil
	})
	_s_DeliveryReportAddress := func(ptr interface{}) { o.DeliveryReportAddress = *ptr.(*string) }
	if err := w.ReadPointer(&o.DeliveryReportAddress, _s_DeliveryReportAddress, _ptr_DeliveryReportAddress); err != nil {
		return err
	}
	_ptr_DocumentName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DocumentName); err != nil {
			return err
		}
		return nil
	})
	_s_DocumentName := func(ptr interface{}) { o.DocumentName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DocumentName, _s_DocumentName, _ptr_DocumentName); err != nil {
		return err
	}
	if err := w.ReadData(&o.CallHandle); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved []uint64
	_Reserved = make([]uint64, 3)
	for i1 := range _Reserved {
		i1 := i1
		if err := w.ReadData((*ndr.Uint3264)(&_Reserved[i1])); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// DeviceReceiveMode type represents FAX_ENUM_DEVICE_RECEIVE_MODE RPC enumeration.
//
// The FAX_ENUM_DEVICE_RECEIVE_MODE enumeration constants describe the receive mode
// for a fax device.
type DeviceReceiveMode uint16

var (
	// FAX_DEVICE_RECEIVE_MODE_OFF: Do not answer incoming calls.
	DeviceReceiveModeOff DeviceReceiveMode = 0
	// FAX_DEVICE_RECEIVE_MODE_AUTO: Automatically answer incoming calls after the specified
	// number of rings.
	DeviceReceiveModeAuto DeviceReceiveMode = 1
	// FAX_DEVICE_RECEIVE_MODE_MANUAL: Manually answer incoming calls.
	DeviceReceiveModeManual DeviceReceiveMode = 2
)

func (o DeviceReceiveMode) String() string {
	switch o {
	case DeviceReceiveModeOff:
		return "DeviceReceiveModeOff"
	case DeviceReceiveModeAuto:
		return "DeviceReceiveModeAuto"
	case DeviceReceiveModeManual:
		return "DeviceReceiveModeManual"
	}
	return "Invalid"
}

// GroupStatus type represents FAX_ENUM_GROUP_STATUS RPC enumeration.
//
// The FAX_ENUM_GROUP_STATUS enumeration defines status types for outbound routing groups.
type GroupStatus uint16

var (
	// FAX_GROUP_STATUS_ALL_DEV_VALID: All the devices in the group are valid and available
	// for sending outgoing faxes.
	GroupStatusAllDevValid GroupStatus = 0
	// FAX_GROUP_STATUS_EMPTY : The group is empty (does not contain any devices), and does
	// not have any routing rules added.
	GroupStatusEmpty GroupStatus = 1
	// FAX_GROUP_STATUS_ALL_DEV_NOT_VALID : All the devices in the group are not available
	// for sending outgoing faxes. Devices could be unavailable if they do not exist or
	// are offline.
	GroupStatusAllDevNotValid GroupStatus = 2
	// FAX_GROUP_STATUS_SOME_DEV_NOT_VALID: Some (but not all) of the devices in the group
	// are not available for sending outgoing faxes. Devices could be unavailable if they
	// do not exist or are offline.
	GroupStatusSomeDevNotValid GroupStatus = 3
)

func (o GroupStatus) String() string {
	switch o {
	case GroupStatusAllDevValid:
		return "GroupStatusAllDevValid"
	case GroupStatusEmpty:
		return "GroupStatusEmpty"
	case GroupStatusAllDevNotValid:
		return "GroupStatusAllDevNotValid"
	case GroupStatusSomeDevNotValid:
		return "GroupStatusSomeDevNotValid"
	}
	return "Invalid"
}

// MessageFolder type represents FAX_ENUM_MESSAGE_FOLDER RPC enumeration.
//
// The FAX_ENUM_MESSAGE_FOLDER enumeration defines possible locations for a fax message.
type MessageFolder uint16

var (
	// FAX_MESSAGE_FOLDER_INBOX: The incoming fax transmission archive, defined in section
	// 3.1.1.
	MessageFolderInbox MessageFolder = 0
	// FAX_MESSAGE_FOLDER_SENTITEMS: The outgoing fax transmission archive, defined in section
	// 3.1.1.
	MessageFolderSentitems MessageFolder = 1
	// FAX_MESSAGE_FOLDER_QUEUE: The Outgoing and Incoming fax queue, defined in section
	// 3.1.1.
	MessageFolderQueue MessageFolder = 2
)

func (o MessageFolder) String() string {
	switch o {
	case MessageFolderInbox:
		return "MessageFolderInbox"
	case MessageFolderSentitems:
		return "MessageFolderSentitems"
	case MessageFolderQueue:
		return "MessageFolderQueue"
	}
	return "Invalid"
}

// PersonalProfileTypes type represents FAX_ENUM_PERSONAL_PROF_TYPES RPC enumeration.
//
// The FAX_ENUM_PERSONAL_PROF_TYPES enumeration defines values to indicate personal
// profile types.
type PersonalProfileTypes uint16

var (
	// RECIPIENT_PERSONAL_PROF: Indicates a recipient profile.
	PersonalProfileTypesRecipientPersonalProfile PersonalProfileTypes = 1
	// SENDER_PERSONAL_PROF: Indicates a sender profile.
	PersonalProfileTypesSenderPersonalProfile PersonalProfileTypes = 2
)

func (o PersonalProfileTypes) String() string {
	switch o {
	case PersonalProfileTypesRecipientPersonalProfile:
		return "PersonalProfileTypesRecipientPersonalProfile"
	case PersonalProfileTypesSenderPersonalProfile:
		return "PersonalProfileTypesSenderPersonalProfile"
	}
	return "Invalid"
}

// PriorityType type represents FAX_ENUM_PRIORITY_TYPE RPC enumeration.
//
// The FAX_ENUM_PRIORITY_TYPE enumeration defines types of priorities for outgoing faxes.
type PriorityType uint16

var (
	// FAX_PRIORITY_TYPE_LOW: The fax is sent with a low priority.
	PriorityTypeLow PriorityType = 0
	// FAX_PRIORITY_TYPE_NORMAL: The fax is sent with a normal priority.
	PriorityTypeNormal PriorityType = 1
	// FAX_PRIORITY_TYPE_HIGH: The fax is sent with a high priority.
	PriorityTypeHigh PriorityType = 2
)

func (o PriorityType) String() string {
	switch o {
	case PriorityTypeLow:
		return "PriorityTypeLow"
	case PriorityTypeNormal:
		return "PriorityTypeNormal"
	case PriorityTypeHigh:
		return "PriorityTypeHigh"
	}
	return "Invalid"
}

// SMTPAuthOptions type represents FAX_ENUM_SMTP_AUTH_OPTIONS RPC enumeration.
//
// The FAX_ENUM_SMTP_AUTH_OPTIONS enumeration defines the type of authentication used
// for SMTP connections.
type SMTPAuthOptions uint16

var (
	// FAX_SMTP_AUTH_ANONYMOUS: The server will send fax transmission receipts using a non-authenticated
	// SMTP server. The server's name and port are defined in the FAX_RECEIPTS_CONFIGW (section
	// 2.2.47) structure.
	SMTPAuthOptionsAnonymous SMTPAuthOptions = 0
	// FAX_SMTP_AUTH_BASIC: The server will send fax transmission receipts using a basic
	// (plain text) authenticated SMTP server. The server's name, port, user name, and password
	// are defined in the FAX_RECEIPTS_CONFIGW.
	SMTPAuthOptionsBasic SMTPAuthOptions = 1
	// FAX_SMTP_AUTH_NTLM: The server will send fax transmission receipts using an NTLM-authenticated
	// SMTP server. The server's name, port, user name, and password are defined in the
	// FAX_RECEIPTS_CONFIGW.
	SMTPAuthOptionsNTLM SMTPAuthOptions = 2
)

func (o SMTPAuthOptions) String() string {
	switch o {
	case SMTPAuthOptionsAnonymous:
		return "SMTPAuthOptionsAnonymous"
	case SMTPAuthOptionsBasic:
		return "SMTPAuthOptionsBasic"
	case SMTPAuthOptionsNTLM:
		return "SMTPAuthOptionsNTLM"
	}
	return "Invalid"
}

// ProductSKUType type represents PRODUCT_SKU_TYPE RPC enumeration.
//
// The PRODUCT_SKU_TYPE enumeration provides values that identify the different Stock
// Keeping Unit (SKU) versions of an operating system.<28>
type ProductSKUType uint16

var (
	// PRODUCT_SKU_UNKNOWN: The operating system is unknown.
	ProductSKUTypeUnknown ProductSKUType = 0
	// PRODUCT_SKU_PERSONAL: Client Personal Edition.
	ProductSKUTypePersonal ProductSKUType = 1
	// PRODUCT_SKU_PROFESSIONAL: Client Professional Edition.
	ProductSKUTypeProfessional ProductSKUType = 2
	// PRODUCT_SKU_SERVER: Server Standard Edition.
	ProductSKUTypeServer ProductSKUType = 4
	// PRODUCT_SKU_ADVANCED_SERVER: Server Advanced Edition
	ProductSKUTypeAdvancedServer ProductSKUType = 8
	// PRODUCT_SKU_DATA_CENTER: Server Datacenter Edition.
	ProductSKUTypeDataCenter ProductSKUType = 16
	// PRODUCT_SKU_DESKTOP_EMBEDDED: Client Embedded Edition.
	ProductSKUTypeDesktopEmbedded ProductSKUType = 32
	// PRODUCT_SKU_SERVER_EMBEDDED: Server Embedded Edition.
	ProductSKUTypeServerEmbedded ProductSKUType = 64
	// PRODUCT_SKU_WEB_SERVER: Server Web Server Edition.
	ProductSKUTypeWebServer ProductSKUType = 128
)

func (o ProductSKUType) String() string {
	switch o {
	case ProductSKUTypeUnknown:
		return "ProductSKUTypeUnknown"
	case ProductSKUTypePersonal:
		return "ProductSKUTypePersonal"
	case ProductSKUTypeProfessional:
		return "ProductSKUTypeProfessional"
	case ProductSKUTypeServer:
		return "ProductSKUTypeServer"
	case ProductSKUTypeAdvancedServer:
		return "ProductSKUTypeAdvancedServer"
	case ProductSKUTypeDataCenter:
		return "ProductSKUTypeDataCenter"
	case ProductSKUTypeDesktopEmbedded:
		return "ProductSKUTypeDesktopEmbedded"
	case ProductSKUTypeServerEmbedded:
		return "ProductSKUTypeServerEmbedded"
	case ProductSKUTypeWebServer:
		return "ProductSKUTypeWebServer"
	}
	return "Invalid"
}

// ConfigOption type represents FAX_ENUM_CONFIG_OPTION RPC enumeration.
//
// The FAX_ENUM_CONFIG_OPTION enumeration identifies the configuration option to be
// returned by the FAX_GetConfigOption (section 3.1.4.1.35) method.
type ConfigOption uint16

var (
	// FAX_CONFIG_OPTION_ALLOW_PERSONAL_CP: Represents whether the server allows personal
	// cover pages.
	ConfigOptionAllowPersonalCreatePartition ConfigOption = 0
	// FAX_CONFIG_OPTION_QUEUE_STATE: Corresponds to the state of the queue in Queue State.
	ConfigOptionQueueState ConfigOption = 1
	// FAX_CONFIG_OPTION_ALLOWED_RECEIPTS: Corresponds to the type of receipts the server
	// is configured to send.
	ConfigOptionAllowedReceipts ConfigOption = 2
	// FAX_CONFIG_OPTION_INCOMING_FAXES_PUBLIC: Corresponds to the viewing permissions of
	// incoming faxes.
	ConfigOptionIncomingFaxesPublic ConfigOption = 3
)

func (o ConfigOption) String() string {
	switch o {
	case ConfigOptionAllowPersonalCreatePartition:
		return "ConfigOptionAllowPersonalCreatePartition"
	case ConfigOptionQueueState:
		return "ConfigOptionQueueState"
	case ConfigOptionAllowedReceipts:
		return "ConfigOptionAllowedReceipts"
	case ConfigOptionIncomingFaxesPublic:
		return "ConfigOptionIncomingFaxesPublic"
	}
	return "Invalid"
}

// Time structure represents FAX_TIME RPC structure.
//
// The _FAX_TIME data type is the custom marshaled variant of the FAX_TIME (section
// 2.2.61) data structure. The _FAX_TIME is used in FAX_GENERAL_CONFIG (section 2.2.31)
// data type and the custom marshaled types _FAX_CONFIGURATIONW (section 2.2.29) and
// _FAX_OUTBOX_CONFIG (section 2.2.17).
//
// This data structure is custom marshaled as follows, and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion                                                                                                                 |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The FAX_TIME structure represents a time, using individual members for the current
// hour and minute. The time is expressed in Coordinated Universal Time (UTC). This
// structure is used in _FAX_CONFIGURATIONW (section 2.2.29), _FAX_OUTBOX_CONFIG (section
// 2.2.17), FAX_GENERAL_CONFIG (section 2.2.31) structures.
type Time struct {
	// Hour: A 16-bit unsigned integer that holds the current hour. This value MUST be between
	// 0 and 23 inclusive.
	Hour uint16 `idl:"name:Hour" json:"hour"`
	// Minute: A 16-bit unsigned integer that holds the current minute. This value MUST
	// be between 0 and 59 inclusive.
	Minute uint16 `idl:"name:Minute" json:"minute"`
}

func (o *Time) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Time) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.Hour); err != nil {
		return err
	}
	if err := w.WriteData(o.Minute); err != nil {
		return err
	}
	return nil
}
func (o *Time) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Hour); err != nil {
		return err
	}
	if err := w.ReadData(&o.Minute); err != nil {
		return err
	}
	return nil
}

// ReceiptsConfigW structure represents FAX_RECEIPTS_CONFIGW RPC structure.
//
// The _FAX_RECEIPTS_CONFIGW data type is the custom-marshaled variant of the FAX_RECEIPTS_CONFIGW
// (section 2.2.47) structure. This data type is used by FAX_GetReceiptsConfiguration
// (section 3.1.4.1.54).
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (40 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Variable_Data (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The FAX_RECEIPTS_CONFIGW structure defines the format for the receipt settings of
// the fax server. This structure is used by FAX_SetReceiptsConfiguration (section 3.1.4.1.91).
// The information contained by this structure describes the delivery receipt support
// fax server configuration (section 3.1.1).
type ReceiptsConfigW struct {
	// dwSizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) value that holds the total size
	// of the structure, in bytes. This value MUST be 40 or 72 bytes. When filled in on
	// a 32-bit implementation, this value SHOULD be 40 bytes. When filled in on a 64-bit
	// implementation, this value SHOULD be 72 bytes.
	StructureSize uint32 `idl:"name:dwSizeOfStruct" json:"structure_size"`
	// dwAllowedReceipts: A DWORD that holds the type of receipts that the server supports.
	// This member MUST be one of the values defined in FAX_ENUM_DELIVERY_REPORT_TYPES (section
	// 2.2.76).
	AllowedReceipts uint32 `idl:"name:dwAllowedReceipts" json:"allowed_receipts"`
	// SMTPAuthOption: A type of SMTP authentication that the server will use for SMTP connections.
	// The options MUST be one of the enumerations defined in FAX_ENUM_SMTP_AUTH_OPTIONS
	// (section 2.2.56).
	SMTPAuthOption SMTPAuthOptions `idl:"name:SMTPAuthOption" json:"smtp_auth_option"`
	// lpwstrReserved: A reserved pointer, which MUST be set to NULL.
	_ string `idl:"name:lpwstrReserved;string"`
	// lpwstrSMTPServer: A null-terminated character string that holds the SMTP server name.
	SMTPServer string `idl:"name:lpwstrSMTPServer;string" json:"smtp_server"`
	// dwSMTPPort: A DWORD that holds the port number of the SMTP server.
	SMTPPort uint32 `idl:"name:dwSMTPPort" json:"smtp_port"`
	// lpwstrSMTPFrom: A null-terminated character string that holds the SMTP email address
	// of the sender of the fax receipt messages.
	SMTPFrom string `idl:"name:lpwstrSMTPFrom;string" json:"smtp_from"`
	// lpwstrSMTPUserName: A null-terminated character string that holds the user name to
	// use for Basic-authenticated SMTP connections.
	SMTPUserName string `idl:"name:lpwstrSMTPUserName;string" json:"smtp_user_name"`
	// lpwstrSMTPPassword: A null-terminated character string that holds the password to
	// use for Basic-authenticated SMTP connections. For anonymous access, no user name
	// and password is required. For Basic and Integrated authentication, a cleartext password
	// is sent over the wire. It is for the server to use the password that depends on the
	// authentication mode.
	SMTPPassword string `idl:"name:lpwstrSMTPPassword;string" json:"smtp_password"`
	// bIsToUseForMSRouteThroughEmailMethod: If set to TRUE, the routing extension MUST
	// use the DRT_EMAIL receipts settings to route incoming faxes by email.
	IsToUseForMsRouteThroughEmailMethod bool `idl:"name:bIsToUseForMSRouteThroughEmailMethod" json:"is_to_use_for_ms_route_through_email_method"`
}

func (o *ReceiptsConfigW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ReceiptsConfigW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowedReceipts); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.SMTPAuthOption)); err != nil {
		return err
	}
	// reserved lpwstrReserved
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	if o.SMTPServer != "" {
		_ptr_lpwstrSMTPServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SMTPServer); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SMTPServer, _ptr_lpwstrSMTPServer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SMTPPort); err != nil {
		return err
	}
	if o.SMTPFrom != "" {
		_ptr_lpwstrSMTPFrom := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SMTPFrom); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SMTPFrom, _ptr_lpwstrSMTPFrom); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SMTPUserName != "" {
		_ptr_lpwstrSMTPUserName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SMTPUserName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SMTPUserName, _ptr_lpwstrSMTPUserName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SMTPPassword != "" {
		_ptr_lpwstrSMTPPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SMTPPassword); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SMTPPassword, _ptr_lpwstrSMTPPassword); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.IsToUseForMsRouteThroughEmailMethod {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ReceiptsConfigW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowedReceipts); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.SMTPAuthOption)); err != nil {
		return err
	}
	// reserved lpwstrReserved
	var _lpwstrReserved string
	_ptr_lpwstrReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &_lpwstrReserved); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrReserved := func(ptr interface{}) { _lpwstrReserved = *ptr.(*string) }
	if err := w.ReadPointer(&_lpwstrReserved, _s_lpwstrReserved, _ptr_lpwstrReserved); err != nil {
		return err
	}
	_ptr_lpwstrSMTPServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SMTPServer); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrSMTPServer := func(ptr interface{}) { o.SMTPServer = *ptr.(*string) }
	if err := w.ReadPointer(&o.SMTPServer, _s_lpwstrSMTPServer, _ptr_lpwstrSMTPServer); err != nil {
		return err
	}
	if err := w.ReadData(&o.SMTPPort); err != nil {
		return err
	}
	_ptr_lpwstrSMTPFrom := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SMTPFrom); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrSMTPFrom := func(ptr interface{}) { o.SMTPFrom = *ptr.(*string) }
	if err := w.ReadPointer(&o.SMTPFrom, _s_lpwstrSMTPFrom, _ptr_lpwstrSMTPFrom); err != nil {
		return err
	}
	_ptr_lpwstrSMTPUserName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SMTPUserName); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrSMTPUserName := func(ptr interface{}) { o.SMTPUserName = *ptr.(*string) }
	if err := w.ReadPointer(&o.SMTPUserName, _s_lpwstrSMTPUserName, _ptr_lpwstrSMTPUserName); err != nil {
		return err
	}
	_ptr_lpwstrSMTPPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SMTPPassword); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrSMTPPassword := func(ptr interface{}) { o.SMTPPassword = *ptr.(*string) }
	if err := w.ReadPointer(&o.SMTPPassword, _s_lpwstrSMTPPassword, _ptr_lpwstrSMTPPassword); err != nil {
		return err
	}
	var _bIsToUseForMsRouteThroughEmailMethod int32
	if err := w.ReadData(&_bIsToUseForMsRouteThroughEmailMethod); err != nil {
		return err
	}
	o.IsToUseForMsRouteThroughEmailMethod = _bIsToUseForMsRouteThroughEmailMethod != 0
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ConfigW structure represents FAX_CONFIGURATIONW RPC structure.
//
// The _FAX_CONFIGURATIONW data type is the custom-marshaled variant of the FAX_CONFIGURATIONW
// (section 2.2.28) structure. This data type is used as an output parameter (as a byte
// array) for FAX_GetConfiguration (section 3.1.4.1.36) and FaxObs_GetConfiguration
// (section 3.1.4.2.24) to return the current fax server configuration settings. Along
// with the FAX_GENERAL_CONFIG (section 2.2.31) data structure, this data structure
// describes the general configuration of the fax server.
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (52 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Variable_Data (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The FAX_CONFIGURATIONW structure is used as an input parameter for the FAX_SetConfiguration
// (section 3.1.4.1.76) and FaxObs_SetConfiguration (section 3.1.4.2.25) methods to
// change the current fax server configuration settings. Along with the FAX_GENERAL_CONFIG
// (section 2.2.31) structure, this data structure describes the general configuration
// of the fax server.
type ConfigW struct {
	// SizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) value that holds the total size of
	// the structure, in bytes. This value MUST be either 52 or 64 bytes. When filled in
	// on a 32-bit implementation, this value SHOULD be 52 bytes. When filled in on a 64-bit
	// implementation, this value SHOULD be 64 bytes.
	StructureSize uint32 `idl:"name:SizeOfStruct" json:"structure_size"`
	// Retries: A DWORD variable that contains the value of the fax transmission retries
	// fax server configuration setting (section 3.1.1).
	Retries uint32 `idl:"name:Retries" json:"retries"`
	// RetryDelay: A DWORD variable that contains the value of the fax transmission retry
	// delay fax server configuration setting.
	RetryDelay uint32 `idl:"name:RetryDelay" json:"retry_delay"`
	// DirtyDays: A DWORD variable that contains the value of the dirty days fax server
	// configuration setting.
	DirtyDays uint32 `idl:"name:DirtyDays" json:"dirty_days"`
	// Branding: A Boolean flag that specifies whether the fax server generates a brand
	// (banner) at the top of outgoing fax transmissions. If this member is TRUE, the fax
	// server generates a brand that contains transmission-related information like the
	// transmitting subscriber identifier, date, time, and page count. This flag configures
	// the Branding fax server configuration setting.
	Branding bool `idl:"name:Branding" json:"branding"`
	// UseDeviceTsid: A Boolean flag that specifies whether the fax server uses the device's
	// transmitting subscriber identifier instead of the value specified in the Tsid member
	// of the FAX_JOB_PARAMW (section 2.2.13) structure. If this member is TRUE, the server
	// uses the device's transmitting subscriber identifier. This flag configures the "use
	// of the device's TSID" fax server configuration setting.
	UseDeviceTSID bool `idl:"name:UseDeviceTsid" json:"use_device_tsid"`
	// ServerCp: A Boolean flag that specifies whether fax client applications can include
	// a user-designed cover page template with the fax transmission. If this member is
	// TRUE, the client MUST use a common cover page template stored on the fax server.
	// If this member is FALSE, the client can use a personal cover page template. This
	// flag configures the personal cover page support fax server configuration setting.
	ServerCreatePartition bool `idl:"name:ServerCp" json:"server_create_partition"`
	// PauseServerQueue: A Boolean flag that specifies whether the fax server has paused
	// the outgoing fax queue. If this member is TRUE, the outgoing fax queue is paused
	// and the Queue State (section 3.1.1) is set to FAX_OUTBOX_PAUSED (0x00000004). If
	// this field is FALSE, the outgoing fax queue is not paused, and the Queue State is
	// either 0x00000000 or FAX_OUTBOX_BLOCKED (0x00000002).
	PauseServerQueue bool `idl:"name:PauseServerQueue" json:"pause_server_queue"`
	// StartCheapTime: Contains a FAX_TIME (section 2.2.61) structure that indicates the
	// hour and minute values of the current start cheap time fax server configuration setting.
	StartCheapTime *Time `idl:"name:StartCheapTime" json:"start_cheap_time"`
	// StopCheapTime: Contains a FAX_TIME that indicates the hour and minute values of the
	// current stop cheap time fax server configuration setting.
	StopCheapTime *Time `idl:"name:StopCheapTime" json:"stop_cheap_time"`
	// ArchiveOutgoingFaxes: A Boolean flag that specifies whether the fax server archives
	// fax transmissions. If this member is TRUE, the server archives transmissions in the
	// directory specified by the ArchiveDirectory member. This flag configures the Archive
	// Enabled fax server configuration setting.<6>
	ArchiveOutgoingFaxes bool `idl:"name:ArchiveOutgoingFaxes" json:"archive_outgoing_faxes"`
	// ArchiveDirectory: A pointer to a constant, null-terminated character string that
	// holds the fully qualified path of the Fax Archive Folder fax server configuration
	// setting. The path can be a UNC path or a path that begins with a drive letter. The
	// fax server ignores this member if the ArchiveOutgoingFaxes member is FALSE. This
	// member can be NULL if the ArchiveOutgoingFaxes member is FALSE.<7>
	ArchiveDirectory string `idl:"name:ArchiveDirectory;string" json:"archive_directory"`
	// ProfileName: Reserved (not used) when this structure is used for FAX_SetConfiguration
	// (section 3.1.4.1.76).
	ProfileName string `idl:"name:ProfileName;string" json:"profile_name"`
}

func (o *ConfigW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ConfigW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if err := w.WriteData(o.Retries); err != nil {
		return err
	}
	if err := w.WriteData(o.RetryDelay); err != nil {
		return err
	}
	if err := w.WriteData(o.DirtyDays); err != nil {
		return err
	}
	if !o.Branding {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.UseDeviceTSID {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.ServerCreatePartition {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.PauseServerQueue {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.StartCheapTime != nil {
		if err := o.StartCheapTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Time{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.StopCheapTime != nil {
		if err := o.StopCheapTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Time{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.ArchiveOutgoingFaxes {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.ArchiveDirectory != "" {
		_ptr_ArchiveDirectory := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ArchiveDirectory); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ArchiveDirectory, _ptr_ArchiveDirectory); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ProfileName != "" {
		_ptr_ProfileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ProfileName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ProfileName, _ptr_ProfileName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ConfigW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.Retries); err != nil {
		return err
	}
	if err := w.ReadData(&o.RetryDelay); err != nil {
		return err
	}
	if err := w.ReadData(&o.DirtyDays); err != nil {
		return err
	}
	var _bBranding int32
	if err := w.ReadData(&_bBranding); err != nil {
		return err
	}
	o.Branding = _bBranding != 0
	var _bUseDeviceTSID int32
	if err := w.ReadData(&_bUseDeviceTSID); err != nil {
		return err
	}
	o.UseDeviceTSID = _bUseDeviceTSID != 0
	var _bServerCreatePartition int32
	if err := w.ReadData(&_bServerCreatePartition); err != nil {
		return err
	}
	o.ServerCreatePartition = _bServerCreatePartition != 0
	var _bPauseServerQueue int32
	if err := w.ReadData(&_bPauseServerQueue); err != nil {
		return err
	}
	o.PauseServerQueue = _bPauseServerQueue != 0
	if o.StartCheapTime == nil {
		o.StartCheapTime = &Time{}
	}
	if err := o.StartCheapTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.StopCheapTime == nil {
		o.StopCheapTime = &Time{}
	}
	if err := o.StopCheapTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bArchiveOutgoingFaxes int32
	if err := w.ReadData(&_bArchiveOutgoingFaxes); err != nil {
		return err
	}
	o.ArchiveOutgoingFaxes = _bArchiveOutgoingFaxes != 0
	_ptr_ArchiveDirectory := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ArchiveDirectory); err != nil {
			return err
		}
		return nil
	})
	_s_ArchiveDirectory := func(ptr interface{}) { o.ArchiveDirectory = *ptr.(*string) }
	if err := w.ReadPointer(&o.ArchiveDirectory, _s_ArchiveDirectory, _ptr_ArchiveDirectory); err != nil {
		return err
	}
	_ptr_ProfileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ProfileName); err != nil {
			return err
		}
		return nil
	})
	_s_ProfileName := func(ptr interface{}) { o.ProfileName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ProfileName, _s_ProfileName, _ptr_ProfileName); err != nil {
		return err
	}
	return nil
}

// GlobalRoutingInfoW structure represents FAX_GLOBAL_ROUTING_INFOW RPC structure.
//
// An array of the FAX_GLOBAL_ROUTING_INFOW structure is used as an input parameter
// to FAX_SetGlobalRoutingInfo (section 3.1.4.1.81) and FaxObs_SetGlobalRoutingInfo
// (section 3.1.4.2.23).
//
// The _FAX_GLOBAL_ROUTING_INFOW structure is the custom-marshaled variant of the FAX_GLOBAL_ROUTING_INFOW
// data structure described in (section 2.2.32). A byte array of this structure is used
// as an output parameter in FAX_EnumGlobalRoutingInfo (section 3.1.4.1.20) and in FaxObs_EnumGlobalRoutingInfo
// (section 3.1.4.2.22).
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (28 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Variable_Data (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type GlobalRoutingInfoW struct {
	// SizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) that holds the total size of the
	// structure, in bytes. This value MUST be 28 bytes or 48 bytes. When filled in on a
	// 32-bit implementation, this value SHOULD be 28 bytes. When filled in on a 64-bit
	// implementation, this value SHOULD be 48 bytes.
	StructureSize uint32 `idl:"name:SizeOfStruct" json:"structure_size"`
	// Priority: A DWORD variable that holds the priority of the fax routing method. The
	// priority determines the relative order in which the fax service calls the fax routing
	// methods when the service receives a fax document. Values for this member MUST be
	// 1 through the maximum DWORD value (0xFFFFFFFF or 4,294,967,295), where 1 is the highest
	// priority.
	Priority uint32 `idl:"name:Priority" json:"priority"`
	// Guid: A pointer to a constant, null-terminated character string that holds the GUID
	// that uniquely identifies the fax routing method of interest.
	GUID string `idl:"name:Guid;string" json:"guid"`
	// FriendlyName: A pointer to a constant, null-terminated character string that holds
	// the user-friendly name to display for the fax routing method.
	FriendlyName string `idl:"name:FriendlyName;string" json:"friendly_name"`
	// FunctionName: A pointer to a null-terminated character string that holds the name
	// of the function that executes the specified fax routing method.
	FunctionName string `idl:"name:FunctionName;string" json:"function_name"`
	// ExtensionImageName: A pointer to a constant, null-terminated character string that
	// holds the name of the fax routing extensions that implements the fax routing method.
	ExtensionImageName string `idl:"name:ExtensionImageName;string" json:"extension_image_name"`
	// ExtensionFriendlyName: A pointer to a constant, null-terminated character string
	// that holds the user-friendly name to display for the fax routing extensions that
	// implement the fax routing method.
	ExtensionFriendlyName string `idl:"name:ExtensionFriendlyName;string" json:"extension_friendly_name"`
}

func (o *GlobalRoutingInfoW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *GlobalRoutingInfoW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if err := w.WriteData(o.Priority); err != nil {
		return err
	}
	if o.GUID != "" {
		_ptr_Guid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GUID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GUID, _ptr_Guid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.FriendlyName != "" {
		_ptr_FriendlyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.FriendlyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FriendlyName, _ptr_FriendlyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.FunctionName != "" {
		_ptr_FunctionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.FunctionName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FunctionName, _ptr_FunctionName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ExtensionImageName != "" {
		_ptr_ExtensionImageName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ExtensionImageName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ExtensionImageName, _ptr_ExtensionImageName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ExtensionFriendlyName != "" {
		_ptr_ExtensionFriendlyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ExtensionFriendlyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ExtensionFriendlyName, _ptr_ExtensionFriendlyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *GlobalRoutingInfoW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.Priority); err != nil {
		return err
	}
	_ptr_Guid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GUID); err != nil {
			return err
		}
		return nil
	})
	_s_Guid := func(ptr interface{}) { o.GUID = *ptr.(*string) }
	if err := w.ReadPointer(&o.GUID, _s_Guid, _ptr_Guid); err != nil {
		return err
	}
	_ptr_FriendlyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.FriendlyName); err != nil {
			return err
		}
		return nil
	})
	_s_FriendlyName := func(ptr interface{}) { o.FriendlyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FriendlyName, _s_FriendlyName, _ptr_FriendlyName); err != nil {
		return err
	}
	_ptr_FunctionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.FunctionName); err != nil {
			return err
		}
		return nil
	})
	_s_FunctionName := func(ptr interface{}) { o.FunctionName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FunctionName, _s_FunctionName, _ptr_FunctionName); err != nil {
		return err
	}
	_ptr_ExtensionImageName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ExtensionImageName); err != nil {
			return err
		}
		return nil
	})
	_s_ExtensionImageName := func(ptr interface{}) { o.ExtensionImageName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ExtensionImageName, _s_ExtensionImageName, _ptr_ExtensionImageName); err != nil {
		return err
	}
	_ptr_ExtensionFriendlyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ExtensionFriendlyName); err != nil {
			return err
		}
		return nil
	})
	_s_ExtensionFriendlyName := func(ptr interface{}) { o.ExtensionFriendlyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ExtensionFriendlyName, _s_ExtensionFriendlyName, _ptr_ExtensionFriendlyName); err != nil {
		return err
	}
	return nil
}

// JobParamExW structure represents FAX_JOB_PARAM_EXW RPC structure.
//
// The FAX_JOB_PARAM_EXW structure defines information about the new job to create when
// sending a fax message.
type JobParamExW struct {
	// dwSizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) that contains the size, in bytes,
	// of this structure. MUST be set to 44 bytes on 32-bit implementations, and MUST be
	// set to 64 bytes on 64-bit implementations.
	StructureSize uint32 `idl:"name:dwSizeOfStruct" json:"structure_size"`
	// dwScheduleAction: A DWORD that MUST specify when to send the fax. This member MUST
	// be one of the following values.
	//
	//	+-----------------------+----------------------------------------------------------------------------------+
	//	|                       |                                                                                  |
	//	|      VALUE/CODE       |                                     MEANING                                      |
	//	|                       |                                                                                  |
	//	+-----------------------+----------------------------------------------------------------------------------+
	//	+-----------------------+----------------------------------------------------------------------------------+
	//	| JSA_NOW 0             | Send the fax as soon as a device is available.                                   |
	//	+-----------------------+----------------------------------------------------------------------------------+
	//	| JSA_SPECIFIC_TIME 1   | Send the fax at the time specified by the tmSchedule member.                     |
	//	+-----------------------+----------------------------------------------------------------------------------+
	//	| JSA_DISCOUNT_PERIOD 2 | Send the fax during the discount rate period. Call the FAX_GetConfiguration      |
	//	|                       | (section 3.1.4.1.36) function to retrieve the discount period for the fax        |
	//	|                       | server.                                                                          |
	//	+-----------------------+----------------------------------------------------------------------------------+
	ScheduleAction uint32 `idl:"name:dwScheduleAction" json:"schedule_action"`
	// tmSchedule: A SYSTEMTIME ([MS-DTYP] section 2.3.13) structure that contains the date
	// and time to send the fax. The time MUST be specified in UTC. This parameter SHOULD
	// be ignored unless dwScheduleAction is set to 1 (JSA_SPECIFIC_TIME). If the time specified
	// has already passed, the method behaves as if 0 (JSA_NOW) was specified.
	Schedule *dtyp.SystemTime `idl:"name:tmSchedule" json:"schedule"`
	// dwReceiptDeliveryType: A DWORD that holds the type of receipt delivered to the sender
	// when the fax is successfully sent and when the fax transmission fails. It can also
	// specify if a receipt will be sent for each recipient or for all the recipients together.
	// The value of this parameter MUST be a logical combination of one of the delivery
	// method flags and optionally one of the delivery grouping flags as specified in FAX_ENUM_DELIVERY_REPORT_TYPES
	// (section 2.2.76). The fax client MUST NOT use the DRT_INBOX value if the protocol
	// version reported by the server is FAX_API_VERSION_2 (0x00020000) or FAX_API_VERSION_3
	// (0x00030000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	ReceiptDeliveryType uint32 `idl:"name:dwReceiptDeliveryType" json:"receipt_delivery_type"`
	// lpwstrReceiptDeliveryAddress: A pointer to a constant, null-terminated character
	// string. If the dwReceiptDeliveryType member contains the DRT_EMAIL or DRT_ATTACH_FAX
	// flag, the string SHOULD be the address to which the delivery receipt (DR) or non-delivery
	// receipt (NDR) SHOULD be sent. If the dwReceiptDeliveryType member is equal to DRT_INBOX,
	// the string SHOULD be the name of the MAPI profile to which the DR or NDR SHOULD be
	// sent. For more information about MAPI, refer to [MSDN-MAPIPRF]. If the dwReceiptDeliveryType
	// member is equal to DRT_MSGBOX, the string SHOULD be the computer name to send the
	// receipt to as a text message containing a character string, as described in Messenger
	// Service Remote Protocol Specification [MS-MSRP] section 3.2.4.1. If the dwReceiptDeliveryType
	// member is set to DRT_NONE, the pointer SHOULD be NULL.
	ReceiptDeliveryAddress string `idl:"name:lpwstrReceiptDeliveryAddress;string" json:"receipt_delivery_address"`
	// Priority: A value specifying the priority level of the outgoing fax.
	Priority PriorityType `idl:"name:Priority" json:"priority"`
	// hCall: Reserved.
	HCall uint32 `idl:"name:hCall" json:"h_call"`
	// dwReserved: This field SHOULD be set to zero.
	_ []uint64 `idl:"name:dwReserved"`
	// lpwstrDocumentName: A null-terminated character string that holds the document name.
	// A NULL pointer value specifies that no document name is specified for this fax job.
	DocumentName string `idl:"name:lpwstrDocumentName;string" json:"document_name"`
	// dwPageCount: A DWORD value that holds the number of pages in the fax document pointed
	// to by the lpcwstrFileName parameter of the FAX_SendDocumentEx (section 3.1.4.1.73)
	// method. This value MUST be used only for fax documents in TIFF, which is the only
	// supported format.
	PageCount uint32 `idl:"name:dwPageCount" json:"page_count"`
}

func (o *JobParamExW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *JobParamExW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if err := w.WriteData(o.ScheduleAction); err != nil {
		return err
	}
	if o.Schedule != nil {
		if err := o.Schedule.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ReceiptDeliveryType); err != nil {
		return err
	}
	if o.ReceiptDeliveryAddress != "" {
		_ptr_lpwstrReceiptDeliveryAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ReceiptDeliveryAddress); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ReceiptDeliveryAddress, _ptr_lpwstrReceiptDeliveryAddress); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.Priority)); err != nil {
		return err
	}
	if err := w.WriteData(o.HCall); err != nil {
		return err
	}
	// reserved dwReserved
	for i1 := 0; uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint64(0)); err != nil {
			return err
		}
	}
	if o.DocumentName != "" {
		_ptr_lpwstrDocumentName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DocumentName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DocumentName, _ptr_lpwstrDocumentName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PageCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *JobParamExW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScheduleAction); err != nil {
		return err
	}
	if o.Schedule == nil {
		o.Schedule = &dtyp.SystemTime{}
	}
	if err := o.Schedule.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReceiptDeliveryType); err != nil {
		return err
	}
	_ptr_lpwstrReceiptDeliveryAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ReceiptDeliveryAddress); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrReceiptDeliveryAddress := func(ptr interface{}) { o.ReceiptDeliveryAddress = *ptr.(*string) }
	if err := w.ReadPointer(&o.ReceiptDeliveryAddress, _s_lpwstrReceiptDeliveryAddress, _ptr_lpwstrReceiptDeliveryAddress); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Priority)); err != nil {
		return err
	}
	if err := w.ReadData(&o.HCall); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved []uint64
	_dwReserved = make([]uint64, 4)
	for i1 := range _dwReserved {
		i1 := i1
		if err := w.ReadData((*ndr.Uint3264)(&_dwReserved[i1])); err != nil {
			return err
		}
	}
	_ptr_lpwstrDocumentName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DocumentName); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrDocumentName := func(ptr interface{}) { o.DocumentName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DocumentName, _s_lpwstrDocumentName, _ptr_lpwstrDocumentName); err != nil {
		return err
	}
	if err := w.ReadData(&o.PageCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// OutboundRoutingGroupw structure represents RPC_FAX_OUTBOUND_ROUTING_GROUPW RPC structure.
//
// The _RPC_FAX_OUTBOUND_ROUTING_GROUPW data type is used as an array of structures
// passed as an output byte-array argument for FAX_EnumOutboundGroups (section 3.1.4.1.26).
// The group name contained by this structure describes one Routing Group (section 3.1.1).
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (16 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Variable_Data (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The RPC_FAX_OUTBOUND_ROUTING_GROUPW data type is used as an input argument for FAX_SetOutboundGroup
// (section 3.1.4.1.85). The group name contained by this structure describes one Routing
// Group (section 3.1.1).
//
// For reference, the data type definition is shown as follows.
type OutboundRoutingGroupw struct {
	// dwSizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) value that holds the total size
	// of the structure, in bytes. This value MUST be 20 or 40 bytes. When filled in on
	// a 32-bit implementation, this value SHOULD be 20 bytes. When filled in on a 64-bit
	// implementation, this value SHOULD be 40 bytes.
	StructureSize uint32 `idl:"name:dwSizeOfStruct" json:"structure_size"`
	// lpwstrGroupName: A null-terminated character string containing the group name. The
	// length of this string MUST be between 1 and 128 characters, excluding the length
	// of the terminating null character. The group name is case-insensitive.
	GroupName string `idl:"name:lpwstrGroupName;string" json:"group_name"`
	// dwNumDevices: A DWORD value that holds the number of devices in the group. The value
	// MUST be in the range between 0 and 1,000. The dwNumDevices parameter also indicates
	// the length of the lpdwDevices array, which is not larger than the actual number of
	// devices.
	DevicesLength uint32 `idl:"name:dwNumDevices" json:"devices_length"`
	// lpdwDevices: A pointer to a DWORD array which contains dwNumDevices entries. Each
	// DWORD value specifies one device identifier in the group. A device MUST only appear
	// once in a group's device list. A single device can belong to one or more groups.
	// Groups are not set to store invalid devices. The order of the devices in the array
	// determines the order the devices are to be used to send faxes, when the group is
	// selected.
	Devices []uint32 `idl:"name:lpdwDevices;size_is:(dwNumDevices);pointer:unique" json:"devices"`
	// Status: Current status of the group from the enumeration FAX_ENUM_GROUP_STATUS (section
	// 2.2.59).
	Status GroupStatus `idl:"name:Status" json:"status"`
}

func (o *OutboundRoutingGroupw) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Devices != nil && o.DevicesLength == 0 {
		o.DevicesLength = uint32(len(o.Devices))
	}
	if o.DevicesLength > uint32(1000) {
		return fmt.Errorf("DevicesLength is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OutboundRoutingGroupw) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if o.GroupName != "" {
		_ptr_lpwstrGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GroupName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GroupName, _ptr_lpwstrGroupName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DevicesLength); err != nil {
		return err
	}
	if o.Devices != nil || o.DevicesLength > 0 {
		_ptr_lpdwDevices := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DevicesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Devices {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Devices[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Devices); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Devices, _ptr_lpdwDevices); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *OutboundRoutingGroupw) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	_ptr_lpwstrGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GroupName); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrGroupName := func(ptr interface{}) { o.GroupName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GroupName, _s_lpwstrGroupName, _ptr_lpwstrGroupName); err != nil {
		return err
	}
	if err := w.ReadData(&o.DevicesLength); err != nil {
		return err
	}
	_ptr_lpdwDevices := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DevicesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DevicesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Devices", sizeInfo[0])
		}
		o.Devices = make([]uint32, sizeInfo[0])
		for i1 := range o.Devices {
			i1 := i1
			if err := w.ReadData(&o.Devices[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpdwDevices := func(ptr interface{}) { o.Devices = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.Devices, _s_lpdwDevices, _ptr_lpdwDevices); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// PortInfo structure represents FAX_PORT_INFO RPC structure.
//
// The FAX_PORT_INFO structure describes one fax port. The data includes, among other
// items, a device identifier, the port's name and current status, and subscriber identifiers.
//
// A fax client application passes the FAX_PORT_INFO in a call to the FAX_SetPort (section
// 3.1.4.1.88) function to modify the configuration of the fax port of interest.
//
// This structure is also used as an input argument for the FaxObs_SetPort (section
// 3.1.4.2.17) method.
//
// The _FAX_PORT_INFO data structure is the custom-marshaled variant of the FAX_PORT_INFO
// (section 2.2.7) data structure. This structure describes one fax port. The data includes,
// among other items, a device identifier, the port's name and current status, and subscriber
// identifiers.
//
// If an application calls the FAX_EnumPorts (section 3.1.4.1.28) function to enumerate
// all the fax devices currently attached to a fax server, the function returns a byte
// array of _FAX_PORT_INFO structures. Each structure describes one device in detail.
//
// If an application calls the FAX_GetPort (section 3.1.4.1.51) function to query one
// device, that function returns information about the device in one _FAX_PORT_INFO.
//
// This structure is also returned as a single structure by the FaxObs_GetPort (section
// 3.1.4.2.16) method and as an array of structures by the FaxObs_EnumPorts (section
// 3.1.4.2.15) method.
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (36 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Variable_Data (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The _FAX_PORT_INFO data structure is the custom-marshaled variant of the FAX_PORT_INFO
// (section 2.2.7) data structure. This structure describes one fax port. The data includes,
// among other items, a device identifier, the port's name and current status, and subscriber
// identifiers.
//
// If an application calls the FAX_EnumPorts (section 3.1.4.1.28) function to enumerate
// all the fax devices currently attached to a fax server, the function returns a byte
// array of _FAX_PORT_INFO structures. Each structure describes one device in detail.
//
// If an application calls the FAX_GetPort (section 3.1.4.1.51) function to query one
// device, that function returns information about the device in one _FAX_PORT_INFO.
//
// This structure is also returned as a single structure by the FaxObs_GetPort (section
// 3.1.4.2.16) method and as an array of structures by the FaxObs_EnumPorts (section
// 3.1.4.2.15) method.
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (36 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Variable_Data (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type PortInfo struct {
	// SizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) that holds the size of the structure,
	// in bytes. This value MUST be 36 bytes or 48 bytes. When filled in on a 32-bit implementation,
	// this value SHOULD be 36 bytes. When filled in on a 64-bit implementation, this value
	// SHOULD be 48 bytes.
	StructureSize uint32 `idl:"name:SizeOfStruct" json:"structure_size"`
	// DeviceId: A DWORD variable that holds the line identifier for the fax device (port)
	// of interest.
	DeviceID uint32 `idl:"name:DeviceId" json:"device_id"`
	// State: A DWORD variable that holds a fax device status code or value. This member
	// can be one of the following predefined device status codes.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|                                 |                                                                                  |
	//	|           VALUE/CODE            |                                     MEANING                                      |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_DIALING 0x20000001          | The device is dialing a fax number.                                              |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_SENDING 0x20000002          | The device is sending a fax document.                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_RECEIVING 0x20000004        | The device is receiving a fax document.                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_COMPLETED 0x20000008        | The device completed sending or receiving a fax transmission.                    |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_HANDLED 0x20000010          | The fax service processed the outbound fax document; the fax service provider    |
	//	|                                 | (FSP) will transmit the fax document.                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_UNAVAILABLE 0x20000020      | The device is not available because it is in use by another application.         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_BUSY 0x20000040             | The device encountered a busy signal.                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_NO_ANSWER 0x20000080        | The receiving device did not answer the call.                                    |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_BAD_ADDRESS 0x20000100      | The device dialed an invalid fax number.                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_NO_DIAL_TONE 0x20000200     | The sending device cannot complete the call because it does not detect a dial    |
	//	|                                 | tone.                                                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_DISCONNECTED 0x20000400     | The fax call was disconnected by the sender or the caller.                       |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_FATAL_ERROR 0x20000800      | The device has encountered a fatal protocol error.                               |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_NOT_FAX_CALL 0x20001000     | The device received a call that was a data call or a voice call.                 |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_CALL_DELAYED 0x20002000     | The device delayed a fax call because the sending device received a busy signal  |
	//	|                                 | multiple times. The device cannot retry the call because dialing restrictions    |
	//	|                                 | exist (some countries and regions restrict the number of retry attempts when a   |
	//	|                                 | number is busy).                                                                 |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_CALL_BLACKLISTED 0x20004000 | The device could not complete a call because the telephone number was blocked or |
	//	|                                 | reserved; emergency numbers such as 911 are blocked.                             |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_INITIALIZING 0x20008000     | The device is initializing a call.                                               |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_OFFLINE 0x20010000          | The device is offline and unavailable.                                           |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_RINGING 0x20020000          | The device is ringing.                                                           |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_AVAILABLE 0x20100000        | The device is available.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_ABORTING 0x20200000         | The device is aborting a fax job.                                                |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_ROUTING 0x20400000          | The device is routing a received fax document.                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_ANSWERED 0x20800000         | The device answered a new call.                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	State uint32 `idl:"name:State" json:"state"`
	// Flags: A DWORD variable that holds a set of bit flags that specify the capability
	// of the fax port. This member can be a bitwise OR combination of the following flag
	// values.
	//
	//	+------------------------+----------------------------------------------------------------------------------+
	//	|                        |                                                                                  |
	//	|       VALUE/CODE       |                                     MEANING                                      |
	//	|                        |                                                                                  |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| FPF_RECEIVE 0x00000001 | The device can receive faxes.                                                    |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| FPF_SEND 0x00000002    | The device can send faxes.                                                       |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| FPF_VIRTUAL 0x00000004 | The device is a virtual fax device. Note that the implementer cannot set a       |
	//	|                        | device to be virtual. When FAX_GetPort (section 3.1.4.1.51) is called, the       |
	//	|                        | FAX_PORT_INFO flag's FPF_VIRTUAL value indicates whether the device is virtual.  |
	//	|                        | When FAX_SetPort (section 3.1.4.1.88) is called, the service will only relate to |
	//	|                        | the FPF_RECEIVE and FPF_SEND values.                                             |
	//	+------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// Rings: A DWORD variable that holds the number of times an incoming fax call rings
	// before the specified device answers the call. Values can be from 0 to 99 inclusive.
	// This value SHOULD be ignored unless the FPF_RECEIVE port capability bit flag is set.
	Rings uint32 `idl:"name:Rings" json:"rings"`
	// Priority: A DWORD variable that holds the priority that determines the relative order
	// in which available fax devices send outgoing transmissions. Values for this member
	// can be 1 through n, where n is the value of the PortsReturned parameter returned
	// by a call to the FAX_EnumPorts (section 3.1.4.1.28) function. When the fax server
	// initiates an outgoing fax transmission, it attempts to select the device with the
	// highest priority and FPF_SEND port capability. If that device is not available, the
	// server selects the next available device that follows in rank order, and so on. The
	// value of the Priority member has no effect on incoming transmissions.
	Priority uint32 `idl:"name:Priority" json:"priority"`
	// DeviceName: A pointer to a constant null-terminated character string that holds the
	// name of the fax device of interest.
	DeviceName string `idl:"name:DeviceName;string" json:"device_name"`
	// Tsid: A pointer to a constant null-terminated character string that holds the transmitting
	// subscriber identifier (TSID). This identifier is usually a telephone number. Only
	// English letters, numeric symbols, and punctuation marks (ASCII range 0x20 to 0x7F)
	// can be used in a TSID.
	TSID string `idl:"name:Tsid;string" json:"tsid"`
	// Csid: A pointer to a constant null-terminated character string that holds the called
	// subscriber identifier (CSID). This identifier is usually a telephone number. Only
	// English letters, numeric symbols, and punctuation marks (ASCII range 0x20 to 0x7F)
	// can be used in a CSID.
	CSID string `idl:"name:Csid;string" json:"csid"`
}

func (o *PortInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PortInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceID); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Rings); err != nil {
		return err
	}
	if err := w.WriteData(o.Priority); err != nil {
		return err
	}
	if o.DeviceName != "" {
		_ptr_DeviceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DeviceName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DeviceName, _ptr_DeviceName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.TSID != "" {
		_ptr_Tsid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.TSID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TSID, _ptr_Tsid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.CSID != "" {
		_ptr_Csid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.CSID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.CSID, _ptr_Csid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PortInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceID); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Rings); err != nil {
		return err
	}
	if err := w.ReadData(&o.Priority); err != nil {
		return err
	}
	_ptr_DeviceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DeviceName); err != nil {
			return err
		}
		return nil
	})
	_s_DeviceName := func(ptr interface{}) { o.DeviceName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DeviceName, _s_DeviceName, _ptr_DeviceName); err != nil {
		return err
	}
	_ptr_Tsid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.TSID); err != nil {
			return err
		}
		return nil
	})
	_s_Tsid := func(ptr interface{}) { o.TSID = *ptr.(*string) }
	if err := w.ReadPointer(&o.TSID, _s_Tsid, _ptr_Tsid); err != nil {
		return err
	}
	_ptr_Csid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.CSID); err != nil {
			return err
		}
		return nil
	})
	_s_Csid := func(ptr interface{}) { o.CSID = *ptr.(*string) }
	if err := w.ReadPointer(&o.CSID, _s_Csid, _ptr_Csid); err != nil {
		return err
	}
	return nil
}

// RuleDestination structure represents FAX_RULE_DESTINATION RPC union.
type RuleDestination struct {
	// Types that are assignable to Value
	//
	// *RuleDestination_DeviceID
	// *RuleDestination_GroupName
	Value is_RuleDestination `json:"value"`
}

func (o *RuleDestination) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *RuleDestination_DeviceID:
		if value != nil {
			return value.DeviceID
		}
	case *RuleDestination_GroupName:
		if value != nil {
			return value.GroupName
		}
	}
	return nil
}

type is_RuleDestination interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_RuleDestination()
}

func (o *RuleDestination) NDRSwitchValue(sw int32) int32 {
	if o == nil {
		return int32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *RuleDestination_DeviceID:
		return int32(0)
	}
	return int32(0)
}

func (o *RuleDestination) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw int32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(int32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case int32(0):
		_o, _ := o.Value.(*RuleDestination_DeviceID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RuleDestination_DeviceID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		_o, _ := o.Value.(*RuleDestination_GroupName)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RuleDestination_GroupName{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *RuleDestination) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw int32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*int32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case int32(0):
		o.Value = &RuleDestination_DeviceID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &RuleDestination_GroupName{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// RuleDestination_DeviceID structure represents FAX_RULE_DESTINATION RPC union arm.
//
// It has following labels: 0
type RuleDestination_DeviceID struct {
	DeviceID uint32 `idl:"name:dwDeviceId" json:"device_id"`
}

func (*RuleDestination_DeviceID) is_RuleDestination() {}

func (o *RuleDestination_DeviceID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.DeviceID); err != nil {
		return err
	}
	return nil
}
func (o *RuleDestination_DeviceID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.DeviceID); err != nil {
		return err
	}
	return nil
}

// RuleDestination_GroupName structure represents FAX_RULE_DESTINATION RPC default union arm.
type RuleDestination_GroupName struct {
	GroupName string `idl:"name:lpwstrGroupName;string" json:"group_name"`
}

func (*RuleDestination_GroupName) is_RuleDestination() {}

func (o *RuleDestination_GroupName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GroupName != "" {
		_ptr_lpwstrGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GroupName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GroupName, _ptr_lpwstrGroupName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RuleDestination_GroupName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_lpwstrGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GroupName); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrGroupName := func(ptr interface{}) { o.GroupName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GroupName, _s_lpwstrGroupName, _ptr_lpwstrGroupName); err != nil {
		return err
	}
	return nil
}

// RuleStatus type represents FAX_ENUM_RULE_STATUS RPC enumeration.
//
// The FAX_ENUM_RULE_STATUS enumeration defines the possible status values for an outbound
// routing rule.
type RuleStatus uint16

var (
	// FAX_RULE_STATUS_VALID: Indicates this outbound routing rule is valid.
	RuleStatusValid RuleStatus = 0
	// FAX_RULE_STATUS_EMPTY_GROUP: Indicates the rule's destination group has no devices.
	RuleStatusEmptyGroup RuleStatus = 1
	// FAX_RULE_STATUS_ALL_GROUP_DEV_NOT_VALID: Indicates that all devices in the rule's
	// destination group are invalid.
	RuleStatusAllGroupDevNotValid RuleStatus = 2
	// FAX_RULE_STATUS_SOME_GROUP_DEV_NOT_VALID: Indicates the rule's destination group
	// has some invalid devices.
	RuleStatusSomeGroupDevNotValid RuleStatus = 3
	// FAX_RULE_STATUS_BAD_DEVICE: Indicates the rule's destination device is not valid.
	RuleStatusBadDevice RuleStatus = 4
)

func (o RuleStatus) String() string {
	switch o {
	case RuleStatusValid:
		return "RuleStatusValid"
	case RuleStatusEmptyGroup:
		return "RuleStatusEmptyGroup"
	case RuleStatusAllGroupDevNotValid:
		return "RuleStatusAllGroupDevNotValid"
	case RuleStatusSomeGroupDevNotValid:
		return "RuleStatusSomeGroupDevNotValid"
	case RuleStatusBadDevice:
		return "RuleStatusBadDevice"
	}
	return "Invalid"
}

// OutboundRoutingRuleW structure represents RPC_FAX_OUTBOUND_ROUTING_RULEW RPC structure.
//
// The RPC_FAX_OUTBOUND_ROUTING_RULEW data type is used as an input argument for FAX_SetOutboundRule
// (section 3.1.4.1.86). The information contained in this structure describes one routing
// rule in the Configuration of the Routing Rules (section 3.1.1).
//
// For reference, the data type definition is as follows.
//
// The _RPC_FAX_OUTBOUND_ROUTING_RULEW data type is used as an array of structures passed
// as an output byte-array argument for FAX_EnumOutboundRules (section 3.1.4.1.27).
// The information contained in each _RPC_FAX_OUTBOUND_ROUTING_RULEW structure describes
// one routing rule in the Configuration of the Routing Rules (section 3.1.1).
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (24 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Variable_Data (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type OutboundRoutingRuleW struct {
	// dwSizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) value that holds the total size
	// of the structure, in bytes. This value MUST be 24 or 40 bytes. When filled in on
	// a 32-bit implementation, this value SHOULD be 24 bytes. When filled in on a 64-bit
	// implementation this value SHOULD be 40 bytes.
	StructureSize uint32 `idl:"name:dwSizeOfStruct" json:"structure_size"`
	// dwAreaCode: A DWORD that contains the area code of the rule. A value of zero indicates
	// the special any-area value ROUTING_RULE_AREA_CODE_ANY. The dwAreaCode and dwCountryCode
	// members MUST form a unique key. This value is unrestricted.
	AreaCode uint32 `idl:"name:dwAreaCode" json:"area_code"`
	// dwCountryCode: A DWORD that contains the country/region code of the rule. A value
	// of zero indicates the special any-country, any-region value ROUTING_RULE_COUNTRY_CODE_ANY.
	// The dwAreaCode and dwCountryCode numeric values MUST form a unique key. Clients can
	// obtain valid values with the FAX_GetCountryList (section 3.1.4.1.37) method.
	CountryCode uint32 `idl:"name:dwCountryCode" json:"country_code"`
	// lpwstrCountryName: A pointer to a null-terminated string that contains the country/region
	// name indicated by the dwCountryCode parameter if known; otherwise, a NULL pointer.
	// If dwCountryCode is zero, this pointer MUST be NULL.
	CountryName string `idl:"name:lpwstrCountryName;string" json:"country_name"`
	// Destination: A FAX_RULE_DESTINATION (section 2.2.81) union that contains the destination
	// of the rule. When bUseGroup is FALSE, this union value MUST hold a device identifier;
	// otherwise, this union value MUST represent the name of an outbound routing group
	// of devices.
	Destination *RuleDestination `idl:"name:Destination;switch_is:bUseGroup" json:"destination"`
	// bUseGroup: A Boolean value that indicates whether the group is used in the destination.
	// If TRUE, the group MUST be used as the rule's destination. If FALSE, the device MUST
	// be used as the rule's destination.
	UseGroup bool `idl:"name:bUseGroup" json:"use_group"`
}

func (o *OutboundRoutingRuleW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OutboundRoutingRuleW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if err := w.WriteData(o.AreaCode); err != nil {
		return err
	}
	if err := w.WriteData(o.CountryCode); err != nil {
		return err
	}
	if o.CountryName != "" {
		_ptr_lpwstrCountryName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.CountryName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.CountryName, _ptr_lpwstrCountryName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	_exprbUseGroup := int32(0)
	if o.UseGroup {
		_exprbUseGroup = int32(1)
	} else {
		_exprbUseGroup = int32(0)
	}
	_swDestination := int32(_exprbUseGroup)
	if o.Destination != nil {
		if err := o.Destination.MarshalUnionNDR(ctx, w, _swDestination); err != nil {
			return err
		}
	} else {
		if err := (&RuleDestination{}).MarshalUnionNDR(ctx, w, _swDestination); err != nil {
			return err
		}
	}
	if !o.UseGroup {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *OutboundRoutingRuleW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.AreaCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.CountryCode); err != nil {
		return err
	}
	_ptr_lpwstrCountryName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.CountryName); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrCountryName := func(ptr interface{}) { o.CountryName = *ptr.(*string) }
	if err := w.ReadPointer(&o.CountryName, _s_lpwstrCountryName, _ptr_lpwstrCountryName); err != nil {
		return err
	}
	if o.Destination == nil {
		o.Destination = &RuleDestination{}
	}
	_exprbUseGroup := int32(0)
	if o.UseGroup {
		_exprbUseGroup = int32(1)
	} else {
		_exprbUseGroup = int32(0)
	}
	_swDestination := int32(_exprbUseGroup)
	if err := o.Destination.UnmarshalUnionNDR(ctx, w, _swDestination); err != nil {
		return err
	}
	var _bUseGroup int32
	if err := w.ReadData(&_bUseGroup); err != nil {
		return err
	}
	o.UseGroup = _bUseGroup != 0
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// Version structure represents FAX_VERSION RPC structure.
//
// The _FAX_VERSION structure is the custom-marshaled variant of the FAX_VERSION (section
// 2.2.22) structure. The _FAX_VERSION contains the same information about the version
// of the fax server components as contained in the FAX_VERSION. The _FAX_VERSION is
// embedded in the FAX_ROUTING_EXTENSION_INFO (section 2.2.49) and FAX_DEVICE_PROVIDER_INFO
// (section 2.2.30) structures.
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (20 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The FAX_VERSION structure contains information about the version of the fax server
// components. This structure is used by FAX_GetVersion (section 3.1.4.1.64).
type Version struct {
	// dwSizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) value that holds the total size
	// of the structure, in bytes. This value MUST be 20 bytes.
	StructureSize uint32 `idl:"name:dwSizeOfStruct" json:"structure_size"`
	// bValid: A Boolean value indicating the validity of the version information returned.
	Valid bool `idl:"name:bValid" json:"valid"`
	// wMajorVersion: A WORD containing the major version number of the fax server component.
	MajorVersion uint16 `idl:"name:wMajorVersion" json:"major_version"`
	// wMinorVersion: A WORD containing the minor version number of the fax server component.
	MinorVersion uint16 `idl:"name:wMinorVersion" json:"minor_version"`
	// wMajorBuildNumber: A WORD containing the major build number of the fax server component.
	MajorBuildNumber uint16 `idl:"name:wMajorBuildNumber" json:"major_build_number"`
	// wMinorBuildNumber: A WORD containing the minor build number of the fax server component.
	MinorBuildNumber uint16 `idl:"name:wMinorBuildNumber" json:"minor_build_number"`
	// dwFlags: A DWORD that MUST contain one of the following values.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|             VALUE/CODE             |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                         | Indicates that the server component was built in release mode. Note  If built in |
	//	|                                    | release mode, this value MUST be zero, which is the default.                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_VER_FLAG_CHECKED 0x00000001    | Indicates that the server component was built in debug mode.                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_VER_FLAG_EVALUATION 0x00000002 | Indicates that the server component was built for evaluation purposes. Reserved  |
	//	|                                    | for future use.                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
}

func (o *Version) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Version) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if !o.Valid {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MajorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.MinorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.MajorBuildNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.MinorBuildNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *Version) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	var _bValid int32
	if err := w.ReadData(&_bValid); err != nil {
		return err
	}
	o.Valid = _bValid != 0
	if err := w.ReadData(&o.MajorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.MajorBuildNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinorBuildNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// OutboxConfig structure represents FAX_OUTBOX_CONFIG RPC structure.
//
// The FAX_OUTBOX_CONFIG structure defines information about outbox settings of the
// fax server. This data structure is used as a parameter to the FAX_SetOutboxConfiguration
// (section 3.1.4.1.87) method.
//
// The _FAX_OUTBOX_CONFIG data type is the custom-marshaled variant of the FAX_OUTBOX_CONFIG
// (section 2.2.16) data structure. The _FAX_OUTBOX_CONFIG data type is returned from
// the FAX_GetOutboxConfiguration (section 3.1.4.1.47) method.
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (36 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type OutboxConfig struct {
	// dwSizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) value that holds the total size
	// of the structure, in bytes. This value MUST be 36 bytes.
	StructureSize uint32 `idl:"name:dwSizeOfStruct" json:"structure_size"`
	// bAllowPersonalCP: A Boolean that indicates whether fax client applications can include
	// a user-designed cover page template with the fax transmission. If this member is
	// TRUE, the client can provide a personal cover page template. If this member is FALSE,
	// the client MUST use a common cover page stored on the fax server.
	AllowPersonalCreatePartition bool `idl:"name:bAllowPersonalCP" json:"allow_personal_create_partition"`
	// bUseDeviceTSID: A Boolean variable that indicates whether the fax server MAY use
	// the devices transmitting subscriber identifier instead of the value specified when
	// submitting a new job. If this member is TRUE, the server SHOULD use the devices transmitting
	// subscriber identifier.
	UseDeviceTSID bool `idl:"name:bUseDeviceTSID" json:"use_device_tsid"`
	// dwRetries: A DWORD that holds the number of times the fax server will attempt to
	// retransmit an outgoing fax if the initial transmission fails.
	Retries uint32 `idl:"name:dwRetries" json:"retries"`
	// dwRetryDelay: A DWORD that holds the minimum number of minutes that will elapse between
	// retransmission attempts by the fax server.
	RetryDelay uint32 `idl:"name:dwRetryDelay" json:"retry_delay"`
	// dtDiscountStart: A FAX_TIME (section 2.2.61) structure that MUST specify the hour
	// and minute at which the discount period begins. The discount period applies only
	// to outgoing transmissions.
	DiscountStart *Time `idl:"name:dtDiscountStart" json:"discount_start"`
	// dtDiscountEnd: A FAX_TIME structure that holds the hour and minute at which the discount
	// period ends. The discount period applies only to outgoing transmissions.
	DiscountEnd *Time `idl:"name:dtDiscountEnd" json:"discount_end"`
	// dwAgeLimit: A DWORD variable that holds the number of days the fax server will keep
	// unsuccessful fax messages in its outbox queue. If a fax message stays in the outbox
	// queue longer than the value specified, it MAY be automatically deleted. If this value
	// is zero, the time limit MUST NOT be used.
	AgeLimit uint32 `idl:"name:dwAgeLimit" json:"age_limit"`
	// bBranding: A Boolean that indicates whether the fax server generates a brand (banner)
	// at the top of outgoing fax transmissions. If this member is TRUE, the fax server
	// SHOULD generate a brand that contains transmission-related information such as the
	// transmitting subscriber identifier, date, time, and page count.
	Branding bool `idl:"name:bBranding" json:"branding"`
}

func (o *OutboxConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OutboxConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if !o.AllowPersonalCreatePartition {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.UseDeviceTSID {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Retries); err != nil {
		return err
	}
	if err := w.WriteData(o.RetryDelay); err != nil {
		return err
	}
	if o.DiscountStart != nil {
		if err := o.DiscountStart.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Time{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DiscountEnd != nil {
		if err := o.DiscountEnd.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Time{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AgeLimit); err != nil {
		return err
	}
	if !o.Branding {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	return nil
}
func (o *OutboxConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	var _bAllowPersonalCreatePartition int32
	if err := w.ReadData(&_bAllowPersonalCreatePartition); err != nil {
		return err
	}
	o.AllowPersonalCreatePartition = _bAllowPersonalCreatePartition != 0
	var _bUseDeviceTSID int32
	if err := w.ReadData(&_bUseDeviceTSID); err != nil {
		return err
	}
	o.UseDeviceTSID = _bUseDeviceTSID != 0
	if err := w.ReadData(&o.Retries); err != nil {
		return err
	}
	if err := w.ReadData(&o.RetryDelay); err != nil {
		return err
	}
	if o.DiscountStart == nil {
		o.DiscountStart = &Time{}
	}
	if err := o.DiscountStart.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DiscountEnd == nil {
		o.DiscountEnd = &Time{}
	}
	if err := o.DiscountEnd.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AgeLimit); err != nil {
		return err
	}
	var _bBranding int32
	if err := w.ReadData(&_bBranding); err != nil {
		return err
	}
	o.Branding = _bBranding != 0
	return nil
}

// ActivityLoggingConfigW structure represents FAX_ACTIVITY_LOGGING_CONFIGW RPC structure.
//
// The FAX_ACTIVITY_LOGGING_CONFIGW structure is used as an input parameter for the
// FAX_SetActivityLoggingConfiguration (section 3.1.4.1.74) call.
type ActivityLoggingConfigW struct {
	// dwSizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) value that holds the size of this
	// structure, in bytes. This value MUST be 16 bytes or 28 bytes. When filled in on a
	// 32-bit implementation, this value SHOULD be 16 bytes. When filled in on a 64-bit
	// implementation, this value SHOULD be 28 bytes.
	StructureSize uint32 `idl:"name:dwSizeOfStruct" json:"structure_size"`
	// bLogIncoming: A Boolean flag that indicates whether incoming fax activities are logged.
	LogIncoming bool `idl:"name:bLogIncoming" json:"log_incoming"`
	// bLogOutgoing: A Boolean flag that indicates whether outgoing fax activities are logged.
	LogOutgoing bool `idl:"name:bLogOutgoing" json:"log_outgoing"`
	// lpwstrDBPath: A pointer to a null-terminated character string that holds the directory
	// on the server where the activity logging database files reside. <5>
	DBPath string `idl:"name:lpwstrDBPath;string" json:"db_path"`
}

func (o *ActivityLoggingConfigW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ActivityLoggingConfigW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if !o.LogIncoming {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.LogOutgoing {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.DBPath != "" {
		_ptr_lpwstrDBPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DBPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DBPath, _ptr_lpwstrDBPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ActivityLoggingConfigW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	var _bLogIncoming int32
	if err := w.ReadData(&_bLogIncoming); err != nil {
		return err
	}
	o.LogIncoming = _bLogIncoming != 0
	var _bLogOutgoing int32
	if err := w.ReadData(&_bLogOutgoing); err != nil {
		return err
	}
	o.LogOutgoing = _bLogOutgoing != 0
	_ptr_lpwstrDBPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DBPath); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrDBPath := func(ptr interface{}) { o.DBPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.DBPath, _s_lpwstrDBPath, _ptr_lpwstrDBPath); err != nil {
		return err
	}
	return nil
}

// PortInfoExW structure represents FAX_PORT_INFO_EXW RPC structure.
//
// The FAX_PORT_INFO_EXW structure defines information about a single fax device, known
// as a port. This structure is used for FAX_SetPortEx (section 3.1.4.1.89).
//
// The _FAX_PORT_INFO_EXW data type is the custom-marshaled variant of the FAX_PORT_INFO_EXW
// (section 2.2.45) structure. This data type is used for FAX_EnumPortsEx (section 3.1.4.1.29)
// and FAX_GetPortEx (section 3.1.4.1.52).
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (48 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Variable_Data (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The _FAX_PORT_INFO_EXW data type is the custom-marshaled variant of the FAX_PORT_INFO_EXW
// (section 2.2.45) structure. This data type is used for FAX_EnumPortsEx (section 3.1.4.1.29)
// and FAX_GetPortEx (section 3.1.4.1.52).
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (48 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Variable_Data (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type PortInfoExW struct {
	// dwSizeOfStruct: DWORD ([MS-DTYP] section 2.2.9) value that holds the total size of
	// the structure, in bytes. This value MUST be 48 or 72 bytes. When filled in on a 32-bit
	// implementation, this value SHOULD be 48 bytes. When filled in on a 64-bit implementation,
	// this value SHOULD be 72 bytes.
	StructureSize uint32 `idl:"name:dwSizeOfStruct" json:"structure_size"`
	// dwDeviceID: A DWORD that holds the line identifier for the specified fax device (port).
	DeviceID uint32 `idl:"name:dwDeviceID" json:"device_id"`
	// lpcwstrDeviceName: A null-terminated character string that holds the name of the
	// fax device.
	DeviceName string `idl:"name:lpcwstrDeviceName;string" json:"device_name"`
	// lpwstrDescription: A null-terminated character string that holds the description
	// of the fax device. The length of this string MUST NOT exceed MAX_FAX_STRING_LEN (section
	// 2.2.86) characters, including the length of the terminating null character.
	Description string `idl:"name:lpwstrDescription;string" json:"description"`
	// lpcwstrProviderName: A null-terminated character string that holds the name of the
	// fax device provider.
	ProviderName string `idl:"name:lpcwstrProviderName;string" json:"provider_name"`
	// lpcwstrProviderGUID: A null-terminated character string that holds the GUID of the
	// fax device provider.
	ProviderGUID string `idl:"name:lpcwstrProviderGUID;string" json:"provider_guid"`
	// bSend: A Boolean value that indicates whether the fax device is enabled to send faxes.
	Send bool `idl:"name:bSend" json:"send"`
	// ReceiveMode: An FAX_ENUM_DEVICE_RECEIVE_MODE (section 2.2.55) enumeration value that
	// indicates whether the fax device is enabled to receive faxes and whether the calls
	// are manually or automatically answered.
	ReceiveMode DeviceReceiveMode `idl:"name:ReceiveMode" json:"receive_mode"`
	// dwStatus: A DWORD that holds the current status of the device. It SHOULD contain
	// any combination of values from the FAX_ENUM_DEVICE_STATUS (section 2.2.64) enumeration
	// or 0 (meaning: status unknown).
	Status uint32 `idl:"name:dwStatus" json:"status"`
	// dwRings: A DWORD that holds the number of times an incoming fax call rings before
	// the specified device answers the call.
	Rings uint32 `idl:"name:dwRings" json:"rings"`
	// lpwstrCsid: A null-terminated character string that holds the called subscriber identifier
	// for faxes sent using this device. This identifier can be a telephone number.
	CSID string `idl:"name:lpwstrCsid;string" json:"csid"`
	// lpwstrTsid: A null-terminated character string that holds the transmitting subscriber
	// identifier for faxes sent using this device. This identifier can be a telephone number.
	TSID string `idl:"name:lpwstrTsid;string" json:"tsid"`
}

func (o *PortInfoExW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PortInfoExW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceID); err != nil {
		return err
	}
	if o.DeviceName != "" {
		_ptr_lpcwstrDeviceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DeviceName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DeviceName, _ptr_lpcwstrDeviceName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_lpwstrDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_lpwstrDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ProviderName != "" {
		_ptr_lpcwstrProviderName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ProviderName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ProviderName, _ptr_lpcwstrProviderName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ProviderGUID != "" {
		_ptr_lpcwstrProviderGUID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ProviderGUID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ProviderGUID, _ptr_lpcwstrProviderGUID); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.Send {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.ReceiveMode)); err != nil {
		return err
	}
	if err := w.WriteData(o.Status); err != nil {
		return err
	}
	if err := w.WriteData(o.Rings); err != nil {
		return err
	}
	if o.CSID != "" {
		_ptr_lpwstrCsid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.CSID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.CSID, _ptr_lpwstrCsid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.TSID != "" {
		_ptr_lpwstrTsid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.TSID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TSID, _ptr_lpwstrTsid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PortInfoExW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceID); err != nil {
		return err
	}
	_ptr_lpcwstrDeviceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DeviceName); err != nil {
			return err
		}
		return nil
	})
	_s_lpcwstrDeviceName := func(ptr interface{}) { o.DeviceName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DeviceName, _s_lpcwstrDeviceName, _ptr_lpcwstrDeviceName); err != nil {
		return err
	}
	_ptr_lpwstrDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_lpwstrDescription, _ptr_lpwstrDescription); err != nil {
		return err
	}
	_ptr_lpcwstrProviderName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ProviderName); err != nil {
			return err
		}
		return nil
	})
	_s_lpcwstrProviderName := func(ptr interface{}) { o.ProviderName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ProviderName, _s_lpcwstrProviderName, _ptr_lpcwstrProviderName); err != nil {
		return err
	}
	_ptr_lpcwstrProviderGUID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ProviderGUID); err != nil {
			return err
		}
		return nil
	})
	_s_lpcwstrProviderGUID := func(ptr interface{}) { o.ProviderGUID = *ptr.(*string) }
	if err := w.ReadPointer(&o.ProviderGUID, _s_lpcwstrProviderGUID, _ptr_lpcwstrProviderGUID); err != nil {
		return err
	}
	var _bSend int32
	if err := w.ReadData(&_bSend); err != nil {
		return err
	}
	o.Send = _bSend != 0
	if err := w.ReadEnum((*uint16)(&o.ReceiveMode)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Status); err != nil {
		return err
	}
	if err := w.ReadData(&o.Rings); err != nil {
		return err
	}
	_ptr_lpwstrCsid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.CSID); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrCsid := func(ptr interface{}) { o.CSID = *ptr.(*string) }
	if err := w.ReadPointer(&o.CSID, _s_lpwstrCsid, _ptr_lpwstrCsid); err != nil {
		return err
	}
	_ptr_lpwstrTsid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.TSID); err != nil {
			return err
		}
		return nil
	})
	_s_lpwstrTsid := func(ptr interface{}) { o.TSID = *ptr.(*string) }
	if err := w.ReadPointer(&o.TSID, _s_lpwstrTsid, _ptr_lpwstrTsid); err != nil {
		return err
	}
	return nil
}

// ServerActivity structure represents FAX_SERVER_ACTIVITY RPC structure.
//
// The _FAX_SERVER_ACTIVITY data type is the custom-marshaled variant of the FAX_SERVER_ACTIVITY
// data structure described in section 2.2.19. The _FAX_SERVER_ACTIVITY defines information
// about the server's fax queue activity and the events reported by the fax server.
// This structure is used as a union field of the FAX_EVENT_EX (section 2.2.67) and
// FAX_EVENT_EX_1 (section 2.2.68) structures.
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (36 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The FAX_SERVER_ACTIVITY structure defines information about the server's fax queue
// activity and the events reported by the fax server. This structure is used as an
// argument for FAX_GetServerActivity (section 3.1.4.1.61).
type ServerActivity struct {
	// dwSizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) value that holds the total size
	// of the structure, in bytes. This value MUST be 36 bytes.
	StructureSize uint32 `idl:"name:dwSizeOfStruct" json:"structure_size"`
	// dwIncomingMessages: A DWORD that indicates the number of messages currently being
	// received by the fax server. This variable MAY also be set to the count of the number
	// of incoming messages that were successfully received and are currently being routed
	// using an inbound routing method. If the routing fails, the incoming job SHOULD be
	// marked for a routing retry and the dwRoutingMessages member used to count this job
	// when the routing restarts. If this value is nonzero, stopping the server MAY result
	// in the loss of incoming messages.
	IncomingMessages uint32 `idl:"name:dwIncomingMessages" json:"incoming_messages"`
	// dwRoutingMessages: A DWORD that indicates the number of incoming messages being rerouted
	// after a routing failure.
	RoutingMessages uint32 `idl:"name:dwRoutingMessages" json:"routing_messages"`
	// dwOutgoingMessages: A DWORD that indicates the number of messages currently being
	// sent by the fax server. If this value is nonzero, stopping the server MAY result
	// in the loss of outgoing messages.
	OutgoingMessages uint32 `idl:"name:dwOutgoingMessages" json:"outgoing_messages"`
	// dwDelegatedOutgoingMessages: A DWORD that indicates the number of messages currently
	// being sent by a Fax Service Provider on behalf of the fax server. The fax server
	// is not currently sending these messages.
	DelegatedOutgoingMessages uint32 `idl:"name:dwDelegatedOutgoingMessages" json:"delegated_outgoing_messages"`
	// dwQueuedMessages: A DWORD that indicates the number of outgoing messages waiting
	// to be processed in the server's fax queue.
	QueuedMessages uint32 `idl:"name:dwQueuedMessages" json:"queued_messages"`
	// dwErrorEvents: A DWORD that indicates the number of error entries added to the system
	// event log since the last time the fax server was started.
	ErrorEvents uint32 `idl:"name:dwErrorEvents" json:"error_events"`
	// dwWarningEvents: A DWORD that indicates the number of warning entries added to the
	// system event log since the last time the fax server was started.
	WarningEvents uint32 `idl:"name:dwWarningEvents" json:"warning_events"`
	// dwInformationEvents: A DWORD that indicates the number of information entries added
	// to the system event log since the last time the fax server was started.
	InformationEvents uint32 `idl:"name:dwInformationEvents" json:"information_events"`
}

func (o *ServerActivity) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerActivity) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if err := w.WriteData(o.IncomingMessages); err != nil {
		return err
	}
	if err := w.WriteData(o.RoutingMessages); err != nil {
		return err
	}
	if err := w.WriteData(o.OutgoingMessages); err != nil {
		return err
	}
	if err := w.WriteData(o.DelegatedOutgoingMessages); err != nil {
		return err
	}
	if err := w.WriteData(o.QueuedMessages); err != nil {
		return err
	}
	if err := w.WriteData(o.ErrorEvents); err != nil {
		return err
	}
	if err := w.WriteData(o.WarningEvents); err != nil {
		return err
	}
	if err := w.WriteData(o.InformationEvents); err != nil {
		return err
	}
	return nil
}
func (o *ServerActivity) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.IncomingMessages); err != nil {
		return err
	}
	if err := w.ReadData(&o.RoutingMessages); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutgoingMessages); err != nil {
		return err
	}
	if err := w.ReadData(&o.DelegatedOutgoingMessages); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueuedMessages); err != nil {
		return err
	}
	if err := w.ReadData(&o.ErrorEvents); err != nil {
		return err
	}
	if err := w.ReadData(&o.WarningEvents); err != nil {
		return err
	}
	if err := w.ReadData(&o.InformationEvents); err != nil {
		return err
	}
	return nil
}

// ReassignInfo structure represents FAX_REASSIGN_INFO RPC structure.
//
// The FAX_REASSIGN_INFO structure contains information about the reassignment of a
// fax.
type ReassignInfo struct {
	// lpcwstrRecipients: A pointer to a constant, null-terminated character string that
	// holds an array of intended recipients to which the fax message can be assigned. The
	// recipients are separated by a semicolon. Each recipient refers to a Fax User Account.
	Recipients string `idl:"name:lpcwstrRecipients;string" json:"recipients"`
	// lpcwstrSenderName: A pointer to a constant null-terminated character string that
	// describes the sender name for the received fax.
	SenderName string `idl:"name:lpcwstrSenderName;string" json:"sender_name"`
	// lpcwstrSenderFaxNumber: A pointer to a constant null-terminated character string
	// that describes the sender fax number for the received fax.
	SenderFaxNumber string `idl:"name:lpcwstrSenderFaxNumber;string" json:"sender_fax_number"`
	// lpcwstrSubject: A pointer to a constant, null-terminated character string that describes
	// the subject of the received fax.
	Subject string `idl:"name:lpcwstrSubject;string" json:"subject"`
	// bHasCoverPage: Boolean value that indicates whether the fax includes a cover page.
	// If this member is TRUE, the fax SHOULD include a cover page.
	HasCoverPage bool `idl:"name:bHasCoverPage" json:"has_cover_page"`
}

func (o *ReassignInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ReassignInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Recipients != "" {
		_ptr_lpcwstrRecipients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Recipients); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Recipients, _ptr_lpcwstrRecipients); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SenderName != "" {
		_ptr_lpcwstrSenderName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SenderName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SenderName, _ptr_lpcwstrSenderName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SenderFaxNumber != "" {
		_ptr_lpcwstrSenderFaxNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SenderFaxNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SenderFaxNumber, _ptr_lpcwstrSenderFaxNumber); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Subject != "" {
		_ptr_lpcwstrSubject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Subject); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Subject, _ptr_lpcwstrSubject); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.HasCoverPage {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ReassignInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_lpcwstrRecipients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Recipients); err != nil {
			return err
		}
		return nil
	})
	_s_lpcwstrRecipients := func(ptr interface{}) { o.Recipients = *ptr.(*string) }
	if err := w.ReadPointer(&o.Recipients, _s_lpcwstrRecipients, _ptr_lpcwstrRecipients); err != nil {
		return err
	}
	_ptr_lpcwstrSenderName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SenderName); err != nil {
			return err
		}
		return nil
	})
	_s_lpcwstrSenderName := func(ptr interface{}) { o.SenderName = *ptr.(*string) }
	if err := w.ReadPointer(&o.SenderName, _s_lpcwstrSenderName, _ptr_lpcwstrSenderName); err != nil {
		return err
	}
	_ptr_lpcwstrSenderFaxNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SenderFaxNumber); err != nil {
			return err
		}
		return nil
	})
	_s_lpcwstrSenderFaxNumber := func(ptr interface{}) { o.SenderFaxNumber = *ptr.(*string) }
	if err := w.ReadPointer(&o.SenderFaxNumber, _s_lpcwstrSenderFaxNumber, _ptr_lpcwstrSenderFaxNumber); err != nil {
		return err
	}
	_ptr_lpcwstrSubject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Subject); err != nil {
			return err
		}
		return nil
	})
	_s_lpcwstrSubject := func(ptr interface{}) { o.Subject = *ptr.(*string) }
	if err := w.ReadPointer(&o.Subject, _s_lpcwstrSubject, _ptr_lpcwstrSubject); err != nil {
		return err
	}
	var _bHasCoverPage int32
	if err := w.ReadData(&_bHasCoverPage); err != nil {
		return err
	}
	o.HasCoverPage = _bHasCoverPage != 0
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// MessageProperties structure represents FAX_MESSAGE_PROPS RPC structure.
//
// The FAX_MESSAGE_PROPS structure defines the properties of a fax message that can
// be set.
type MessageProperties struct {
	// dwValidityMask: A DWORD ([MS-DTYP] section 2.2.9) value that defines a bitwise combination
	// of valid fields in the structure.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|             VALUE/CODE              |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_MSG_PROP_FIELD_MSG_FLAGS 0x0001 | Indicates whether the value in dwMsgFlags is valid. If this bit is set, the      |
	//	|                                     | value in dwMsgFlags is valid.                                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	ValidityMask uint32 `idl:"name:dwValidityMask" json:"validity_mask"`
	// dwMsgFlags: A DWORD bitmask that specifies the state to which the message flags are
	// set.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|          VALUE/CODE          |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_MSG_FLAG_READ 0x00000001 | Determines whether this fax message is marked as read. If this bit is set,       |
	//	|                              | the message is marked as read. If this bit is reset, the message is marked as    |
	//	|                              | unread.                                                                          |
	//	+------------------------------+----------------------------------------------------------------------------------+
	MessageFlags uint32 `idl:"name:dwMsgFlags" json:"message_flags"`
}

func (o *MessageProperties) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MessageProperties) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ValidityMask); err != nil {
		return err
	}
	if err := w.WriteData(o.MessageFlags); err != nil {
		return err
	}
	return nil
}
func (o *MessageProperties) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValidityMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.MessageFlags); err != nil {
		return err
	}
	return nil
}

// JobEntry structure represents FAX_JOB_ENTRY RPC structure.
//
// The FAX_JOB_ENTRY structure describes one fax job. The structure includes information
// about the job type and status, the recipient's and the sender's personal profiles
// (section 3.1.1), scheduling and delivery settings, and the page count.
//
// This structure is used as an input argument for the FaxObs_SetJob (section 3.1.4.2.11)
// method.
//
// The _FAX_JOB_ENTRY structure is the custom-marshaled variant of the FAX_JOB_ENTRY
// (section 2.2.5) structure and describes one fax job. The structure includes information
// about the job type and status, the personal profiles (section 3.1.1) of the recipient
// and sender, scheduling and delivery settings, and the page count. The SizeOfStruct,
// RecipientNumber, and QueueStatus fields in the Fixed_Portion block MUST NOT be 0.
// Except for these fields and the JobId field, all fields of this structure are optional,
// and if the respective information is not available, the fields in the Fixed_Portion
// block MUST be zero.
//
// An application can call the FAX_GetJob (section 3.1.4.1.41) method to retrieve information
// about a specified job at the server, information which is returned in a _FAX_JOB_ENTRY
// structure.
//
// An application can call the FAX_EnumJobs function (section 3.1.4.1.21) to enumerate
// all queued and active fax jobs (see Fax Queue in section 3.1.1 for details) on the
// fax server of interest. The FAX_EnumJobs function returns an array of _FAX_JOB_ENTRY
// structures. Each structure describes one fax job in detail.
//
// This structure is also returned as a single structure by the FaxObs_GetJob (section
// 3.1.4.2.10) method and as an array of structures by the FaxObs_EnumJobs (section
// 3.1.4.2.9) method.
//
// This data structure is custom marshaled as follows and uses the custom-marshaling
// rules defined in section 2.2.1.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Fixed_Portion (92 bytes)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Variable_Data (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type JobEntry struct {
	// SizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) that indicates the size, in bytes,
	// of the FAX_JOB_ENTRY structure. This value MUST be 92 or 136 bytes. When filled in
	// on a 32-bit implementation, this value SHOULD be 92 bytes. When filled in on a 64-bit
	// implementation, this value SHOULD be 136 bytes.
	StructureSize uint32 `idl:"name:SizeOfStruct" json:"structure_size"`
	// JobId: A DWORD that indicates a unique number that identifies the fax jobs of interest.
	// This is the same kind of job identifier number as the JobId parameter for the FAX_SetJob
	// (section 3.1.4.1.82) function.
	JobID uint32 `idl:"name:JobId" json:"job_id"`
	// UserName: A null-terminated character string that contains the name of the fax user
	// account that submitted the fax job, if known; otherwise, a NULL pointer.
	UserName string `idl:"name:UserName" json:"user_name"`
	// JobType: A DWORD that indicates the type of the fax job of interest. This field is
	// one of the following values, which are defined in section 3.1.1.
	//
	//	+------------+----------------------------------+
	//	|            |                                  |
	//	| VALUE/CODE |             MEANING              |
	//	|            |                                  |
	//	+------------+----------------------------------+
	//	+------------+----------------------------------+
	//	| 0x00000000 | The job type is JT_UNKNOWN.      |
	//	+------------+----------------------------------+
	//	| 0x00000001 | The job type is JT_SEND.         |
	//	+------------+----------------------------------+
	//	| 0x00000002 | The job type is JT_RECEIVE.      |
	//	+------------+----------------------------------+
	//	| 0x00000003 | The job type is JT_ROUTING.      |
	//	+------------+----------------------------------+
	//	| 0x00000004 | The job type is JT_FAIL_RECEIVE. |
	//	+------------+----------------------------------+
	JobType uint32 `idl:"name:JobType" json:"job_type"`
	// QueueStatus: A DWORD variable containing a set of bit flags indicating the job status
	// of the fax job identified by the JobId field. This value MUST be a bitwise OR combination
	// of one or more of the Job Status values listed in section 3.1.1.
	QueueStatus uint32 `idl:"name:QueueStatus" json:"queue_status"`
	// Status: A DWORD that specifies the status of the fax device (or port) that received
	// or sent the fax job described by this structure, captured at the time the job information
	// was recorded. This member SHOULD be ignored when this structure is used as an input
	// argument for the FaxObs_SetJob method. This value MUST be one of the following predefined
	// device status codes.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|                                 |                                                                                  |
	//	|           VALUE/CODE            |                                     MEANING                                      |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_UNKNOWN 0x00000000          | The status of the device is unknown.                                             |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_DIALING 0x20000001          | The device is dialing a fax number.                                              |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_SENDING 0x20000002          | The device is sending a fax document.                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_RECEIVING 0x20000004        | The device is receiving a fax document.                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_COMPLETED 0x20000008        | The device completed sending or receiving a fax transmission.                    |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_HANDLED 0x20000010          | The fax service processed the outbound fax document; the fax service provider    |
	//	|                                 | (FSP) will transmit the fax document.                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_UNAVAILABLE 0x20000020      | The device is not available because it is in use by another application.         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_BUSY 0x20000040             | The device encountered a busy signal.                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_NO_ANSWER 0x20000080        | The receiving device did not answer the call.                                    |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_BAD_ADDRESS 0x20000100      | The device dialed an invalid fax number.                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_NO_DIAL_TONE 0x20000200     | The sending device cannot complete the call because it does not detect a dial    |
	//	|                                 | tone.                                                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_DISCONNECTED 0x20000400     | The fax call was disconnected by the sender or the caller.                       |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_FATAL_ERROR 0x20000800      | The device has encountered a fatal protocol error.                               |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_NOT_FAX_CALL 0x20001000     | The device received a call that was a data call or a voice call.                 |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_CALL_DELAYED 0x20002000     | The device delayed a fax call because the sending device received a busy signal  |
	//	|                                 | multiple times. The device cannot retry the call because dialing restrictions    |
	//	|                                 | exist (some countries and regions restrict the number of retry attempts when a   |
	//	|                                 | number is busy).                                                                 |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_CALL_BLACKLISTED 0x20004000 | The device could not complete a call because the telephone number was blocked or |
	//	|                                 | reserved; emergency numbers such as 911 are blocked.                             |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_INITIALIZING 0x20008000     | The device is initializing a call.                                               |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_OFFLINE 0x20010000          | The device is offline and unavailable.                                           |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_RINGING 0x20020000          | The device is ringing.                                                           |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_AVAILABLE 0x20100000        | The device is available.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_ABORTING 0x20200000         | The device is aborting a fax job.                                                |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_ROUTING 0x20400000          | The device is routing a received fax document.                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FPS_ANSWERED 0x20800000         | The device answered a new call.                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	Status uint32 `idl:"name:Status" json:"status"`
	// Size: A DWORD variable that indicates the total size, in bytes, of the fax document
	// to transmit, if known, or zero otherwise.  The size, if known, includes the size
	// of the cover page, if a cover page is present, and the size of the fax body, if a
	// fax body is present. The size MUST NOT exceed 4 gigabytes.
	Size uint32 `idl:"name:Size" json:"size"`
	// PageCount: A DWORD that indicates the total number of pages in the fax transmission,
	// including the cover page, if any, and the fax body, if any, of the fax submitted
	// with this fax job. If the fax is sent to multiple recipients, this total number of
	// pages is the number of fax pages sent to each individual recipient (not the sum of
	// the fax pages sent to all recipients).
	PageCount uint32 `idl:"name:PageCount" json:"page_count"`
	// RecipientNumber: A null-terminated character string that contains the fax number
	// of the recipient of the fax transmission, if known, or a NULL pointer otherwise.
	// This information comes from the recipient's personal profile.
	RecipientNumber string `idl:"name:RecipientNumber" json:"recipient_number"`
	// RecipientName: A null-terminated character string that contains the name of the recipient
	// of the fax, if known, or a NULL pointer otherwise. This information comes from the
	// recipient's personal profile.
	RecipientName string `idl:"name:RecipientName" json:"recipient_name"`
	// Tsid: A null-terminated character string that contains the transmitting subscriber
	// identifier (TSID), if known, or a NULL pointer otherwise. This information comes
	// from the sender's personal profile.
	TSID string `idl:"name:Tsid" json:"tsid"`
	// SenderName: A null-terminated character string that contains the fax sender name,
	// if known, or a NULL pointer otherwise. This information comes from the sender's personal
	// profile.
	SenderName string `idl:"name:SenderName" json:"sender_name"`
	// SenderCompany: A null-terminated character string that contains the fax sender company,
	// if known, or a NULL pointer otherwise. This information comes from the sender's personal
	// profile.
	SenderCompany string `idl:"name:SenderCompany" json:"sender_company"`
	// SenderDept: A null-terminated character string that contains the fax sender department,
	// if known, or a NULL pointer otherwise. This information comes from the sender's personal
	// profile.
	SenderDept string `idl:"name:SenderDept" json:"sender_dept"`
	// BillingCode: A null-terminated character string that contains the fax billing code,
	// if known, or a NULL pointer otherwise.
	BillingCode string `idl:"name:BillingCode" json:"billing_code"`
	// ScheduleAction: A DWORD that indicates when the fax is to be sent. This can be one
	// of the following values:
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|                                |                                                                                  |
	//	|           VALUE/CODE           |                                     MEANING                                      |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| JSA_NOW 0x00000000             | The fax is to be sent as soon as a fax device is available.                      |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| JSA_SPECIFIC_TIME 0x00000001   | The fax is to be sent at the time specified by the ScheduleTime field of this    |
	//	|                                | FAX_JOB_ENTRY structure.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| JSA_DISCOUNT_PERIOD 0x00000002 | The fax is to be sent during the discount rate period. The                       |
	//	|                                | FaxObs_GetConfiguration (section 3.1.4.2.24) method can be called to retrieve    |
	//	|                                | the discount period for the fax server.                                          |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	ScheduleAction uint32 `idl:"name:ScheduleAction" json:"schedule_action"`
	// ScheduleTime: A SYSTEMTIME ([MS-DTYP] section 2.3.13) structure indicating the local
	// date and time to send the fax, in Coordinated Universal Time (UTC) format. This parameter
	// MUST be ignored unless the ScheduleAction parameter is set to 1 (JSA_SPECIFIC_TIME).
	ScheduleTime *dtyp.SystemTime `idl:"name:ScheduleTime" json:"schedule_time"`
	// DeliveryReportType: A DWORD variable that indicates the fax delivery report type.
	// This value MUST be one of the FAX_ENUM_DELIVERY_REPORT_TYPES (section 2.2.76) enumeration
	// values. The DRT_ATTACH_FAX value can be combined with the DRT_EMAIL value by an OR
	// operation.
	DeliveryReportType uint32 `idl:"name:DeliveryReportType" json:"delivery_report_type"`
	// DeliveryReportAddress: A null-terminated character string that contains the email
	// address for the delivery report, if known, or a NULL pointer otherwise.
	DeliveryReportAddress string `idl:"name:DeliveryReportAddress" json:"delivery_report_address"`
	// DocumentName: A null-terminated character string that contains the document name,
	// if known, or a NULL pointer otherwise.
	DocumentName string `idl:"name:DocumentName" json:"document_name"`
}

func (o *JobEntry) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *JobEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if err := w.WriteData(o.JobID); err != nil {
		return err
	}
	if o.UserName != "" {
		_ptr_UserName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.UserName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.UserName, _ptr_UserName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.JobType); err != nil {
		return err
	}
	if err := w.WriteData(o.QueueStatus); err != nil {
		return err
	}
	if err := w.WriteData(o.Status); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.PageCount); err != nil {
		return err
	}
	if o.RecipientNumber != "" {
		_ptr_RecipientNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.RecipientNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecipientNumber, _ptr_RecipientNumber); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RecipientName != "" {
		_ptr_RecipientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.RecipientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RecipientName, _ptr_RecipientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.TSID != "" {
		_ptr_Tsid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.TSID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TSID, _ptr_Tsid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SenderName != "" {
		_ptr_SenderName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.SenderName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SenderName, _ptr_SenderName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SenderCompany != "" {
		_ptr_SenderCompany := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.SenderCompany); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SenderCompany, _ptr_SenderCompany); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SenderDept != "" {
		_ptr_SenderDept := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.SenderDept); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SenderDept, _ptr_SenderDept); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.BillingCode != "" {
		_ptr_BillingCode := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.BillingCode); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.BillingCode, _ptr_BillingCode); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ScheduleAction); err != nil {
		return err
	}
	if o.ScheduleTime != nil {
		if err := o.ScheduleTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DeliveryReportType); err != nil {
		return err
	}
	if o.DeliveryReportAddress != "" {
		_ptr_DeliveryReportAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.DeliveryReportAddress); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DeliveryReportAddress, _ptr_DeliveryReportAddress); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DocumentName != "" {
		_ptr_DocumentName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.DocumentName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DocumentName, _ptr_DocumentName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *JobEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.JobID); err != nil {
		return err
	}
	_ptr_UserName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.UserName); err != nil {
			return err
		}
		return nil
	})
	_s_UserName := func(ptr interface{}) { o.UserName = *ptr.(*string) }
	if err := w.ReadPointer(&o.UserName, _s_UserName, _ptr_UserName); err != nil {
		return err
	}
	if err := w.ReadData(&o.JobType); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueueStatus); err != nil {
		return err
	}
	if err := w.ReadData(&o.Status); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.PageCount); err != nil {
		return err
	}
	_ptr_RecipientNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.RecipientNumber); err != nil {
			return err
		}
		return nil
	})
	_s_RecipientNumber := func(ptr interface{}) { o.RecipientNumber = *ptr.(*string) }
	if err := w.ReadPointer(&o.RecipientNumber, _s_RecipientNumber, _ptr_RecipientNumber); err != nil {
		return err
	}
	_ptr_RecipientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.RecipientName); err != nil {
			return err
		}
		return nil
	})
	_s_RecipientName := func(ptr interface{}) { o.RecipientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.RecipientName, _s_RecipientName, _ptr_RecipientName); err != nil {
		return err
	}
	_ptr_Tsid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.TSID); err != nil {
			return err
		}
		return nil
	})
	_s_Tsid := func(ptr interface{}) { o.TSID = *ptr.(*string) }
	if err := w.ReadPointer(&o.TSID, _s_Tsid, _ptr_Tsid); err != nil {
		return err
	}
	_ptr_SenderName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.SenderName); err != nil {
			return err
		}
		return nil
	})
	_s_SenderName := func(ptr interface{}) { o.SenderName = *ptr.(*string) }
	if err := w.ReadPointer(&o.SenderName, _s_SenderName, _ptr_SenderName); err != nil {
		return err
	}
	_ptr_SenderCompany := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.SenderCompany); err != nil {
			return err
		}
		return nil
	})
	_s_SenderCompany := func(ptr interface{}) { o.SenderCompany = *ptr.(*string) }
	if err := w.ReadPointer(&o.SenderCompany, _s_SenderCompany, _ptr_SenderCompany); err != nil {
		return err
	}
	_ptr_SenderDept := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.SenderDept); err != nil {
			return err
		}
		return nil
	})
	_s_SenderDept := func(ptr interface{}) { o.SenderDept = *ptr.(*string) }
	if err := w.ReadPointer(&o.SenderDept, _s_SenderDept, _ptr_SenderDept); err != nil {
		return err
	}
	_ptr_BillingCode := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.BillingCode); err != nil {
			return err
		}
		return nil
	})
	_s_BillingCode := func(ptr interface{}) { o.BillingCode = *ptr.(*string) }
	if err := w.ReadPointer(&o.BillingCode, _s_BillingCode, _ptr_BillingCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScheduleAction); err != nil {
		return err
	}
	if o.ScheduleTime == nil {
		o.ScheduleTime = &dtyp.SystemTime{}
	}
	if err := o.ScheduleTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeliveryReportType); err != nil {
		return err
	}
	_ptr_DeliveryReportAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.DeliveryReportAddress); err != nil {
			return err
		}
		return nil
	})
	_s_DeliveryReportAddress := func(ptr interface{}) { o.DeliveryReportAddress = *ptr.(*string) }
	if err := w.ReadPointer(&o.DeliveryReportAddress, _s_DeliveryReportAddress, _ptr_DeliveryReportAddress); err != nil {
		return err
	}
	_ptr_DocumentName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.DocumentName); err != nil {
			return err
		}
		return nil
	})
	_s_DocumentName := func(ptr interface{}) { o.DocumentName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DocumentName, _s_DocumentName, _ptr_DocumentName); err != nil {
		return err
	}
	return nil
}

// CompletionEvent structure represents FAX_EVENT RPC structure.
//
// The FAX_EVENT structure represents the contents of an input/output (I/O) completion
// packet. The fax server sends the completion packet to notify a fax client application
// about an asynchronous fax server event.
type CompletionEvent struct {
	// SizeOfStruct: A DWORD ([MS-DTYP] section 2.2.9) value that holds the total size of
	// the structure, in bytes. This value MUST be 24 bytes.
	StructureSize uint32 `idl:"name:SizeOfStruct" json:"structure_size"`
	// TimeStamp: Specifies a FILETIME ([MS-DTYP] section 2.3.3) structure that contains
	// the time at which the fax server generated the event.
	Timestamp *dtyp.Filetime `idl:"name:TimeStamp" json:"timestamp"`
	// DeviceId: Specifies a DWORD variable that indicates the line identifier for the fax
	// device (port) of interest.
	DeviceID uint32 `idl:"name:DeviceId" json:"device_id"`
	// EventId: Specifies a DWORD variable that identifies the current asynchronous event
	// that occurred within the fax server. The following table lists the possible events
	// and their meanings.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|            VALUE/CODE             |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_DIALING 0x00000001            | The sending device is dialing a fax number.                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_SENDING 0x00000002            | The sending device is transmitting a page of fax data.                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_RECEIVING 0x00000003          | The receiving device is receiving a page of fax data.                            |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_COMPLETED 0x00000004          | The device has completed a fax transmission call.                                |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_BUSY 0x00000005               | The sending device has encountered a busy signal.                                |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_NO_ANSWER 0x00000006          | The receiving device does not answer.                                            |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_BAD_ADDRESS 0x00000007        | The sending device cannot complete the call because the fax number is invalid.   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_NO_DIAL_TONE 0x00000008       | The sending device cannot complete the call because it does not detect a dial    |
	//	|                                   | tone.                                                                            |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_DISCONNECTED 0x00000009       | The device cannot complete the call because a fax device was disconnected or     |
	//	|                                   | because the fax call itself was disconnected.                                    |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_FATAL_ERROR 0x0000000A        | The device encountered a fatal protocol error.                                   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_NOT_FAX_CALL 0x0000000B       | The modem device received a data call or a voice call.                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_CALL_DELAYED 0x0000000C       | The sending device received a busy signal multiple times. The device cannot      |
	//	|                                   | retry the call because dialing restrictions exist (some countries and regions    |
	//	|                                   | restrict the number of retry attempts when a number is busy).                    |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_CALL_BLACKLISTED 0x0000000D   | The device cannot complete the call because the telephone number is blocked or   |
	//	|                                   | reserved; numbers such as 911 are blocked.                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_RINGING 0x0000000E            | The receiving device is ringing.                                                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_ABORTING 0x0000000F           | The device is aborting a fax job.                                                |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_ROUTING 0x00000010            | The receiving device is routing a received fax document.                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_MODEM_POWERED_ON 0x000000011  | The modem device was turned on.                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_MODEM_POWERED_OFF 0x000000012 | The modem device was turned off.                                                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_IDLE 0x000000013              | The device is idle.                                                              |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_FAXSVC_ENDED 0x000000014      | The fax service has terminated. For more information, see the following Remarks  |
	//	|                                   | section.                                                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_ANSWERED 0x000000015          | The receiving device answered a new call.                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_JOB_QUEUED 0x000000016        | The fax job has been queued.                                                     |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_DELETED 0x00000017            | The fax job has been processed. The job identifier for the job is no longer      |
	//	|                                   | valid.                                                                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_INITIALIZING 0x00000018       | The modem device is being initialized.                                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_LINE_UNAVAILABLE 0x00000019   | The device cannot complete the call because the requested line is unavailable.   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_HANDLED 0x0000001A            | The fax job has been processed.                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_FAXSVC_STARTED 0x0000001B     | The fax service has started. For more information, see the following Remarks     |
	//	|                                   | section. Interchangeable with FEI_NEVENTS.                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| FEI_NEVENTS 0x0000001B            | The total number of fax events received. For more information, see the following |
	//	|                                   | Remarks section. Interchangeable with FEI_FAXSVC_STARTED.                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	EventID uint32 `idl:"name:EventId" json:"event_id"`
	// JobId: Specifies a unique number that identifies the fax job of interest. If this
	// member is equal to the value 0xffffffff, it indicates an inactive fax job. Note that
	// this number is not a print spooler identification number.
	//
	// After a fax client application receives the FEI_FAXSVC_ENDED message from the fax
	// service, it will no longer receive fax events. To resume receiving fax events, the
	// application MUST call the FaxInitializeEventQueue function again when the fax service
	// restarts. The application can determine whether the fax service is running by using
	// the service control manager.
	//
	// If the application receives events by means of notification messages, it can use
	// the FEI_NEVENTS event. If the message is between the application's base window message
	// and the base window message + FEI_NEVENTS, then the application can process the message
	// as a fax window message. An application specifies the base window message by using
	// the MessageStart parameter in the FaxInitializeEventQueue function; the base window
	// message MUST be greater than the WM_USER message.
	JobID uint32 `idl:"name:JobId" json:"job_id"`
}

func (o *CompletionEvent) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *CompletionEvent) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.StructureSize); err != nil {
		return err
	}
	if o.Timestamp != nil {
		if err := o.Timestamp.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DeviceID); err != nil {
		return err
	}
	if err := w.WriteData(o.EventID); err != nil {
		return err
	}
	if err := w.WriteData(o.JobID); err != nil {
		return err
	}
	return nil
}
func (o *CompletionEvent) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.StructureSize); err != nil {
		return err
	}
	if o.Timestamp == nil {
		o.Timestamp = &dtyp.Filetime{}
	}
	if err := o.Timestamp.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceID); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventID); err != nil {
		return err
	}
	if err := w.ReadData(&o.JobID); err != nil {
		return err
	}
	return nil
}
