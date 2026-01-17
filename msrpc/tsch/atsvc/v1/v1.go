package atsvc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	tsch "github.com/oiweiwei/go-msrpc/msrpc/tsch"
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
	_ = tsch.GoPackage
)

var (
	// import guard
	GoPackage = "tsch"
)

var (
	// Syntax UUID
	ATSvcSyntaxUUID = &uuid.UUID{TimeLow: 0x1ff70682, TimeMid: 0xa51, TimeHiAndVersion: 0x30e8, ClockSeqHiAndReserved: 0x7, ClockSeqLow: 0x6d, Node: [6]uint8{0x74, 0xb, 0xe8, 0xce, 0xe9, 0x8b}}
	// Syntax ID
	ATSvcSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: ATSvcSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// atsvc interface.
type ATSvcClient interface {

	// The NetrJobAdd method MUST add a single AT task to the server's task store.
	//
	// Return Values: For more information on return codes, see section 2.3.14 or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	JobAdd(context.Context, *JobAddRequest, ...dcerpc.CallOption) (*JobAddResponse, error)

	// The NetrJobDel method MUST delete a specified range of tasks from the task store.
	// The method is capable of deleting all AT tasks or just a subset of the tasks, as
	// determined by the values of the MinJobId and MaxJobId parameters.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	//
	// To delete all tasks, specify MinJobId as 0 and MaxJobId as 0xFFFFFFFF.
	JobDelete(context.Context, *JobDeleteRequest, ...dcerpc.CallOption) (*JobDeleteResponse, error)

	// The NetrJobEnum method MUST return an enumeration of all AT tasks on the specified
	// server.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	JobEnum(context.Context, *JobEnumRequest, ...dcerpc.CallOption) (*JobEnumResponse, error)

	// The NetrJobGetInfo method MUST return information for a specified ATSvc task. The
	// task identifier MUST be used to locate the task configuration.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	JobGetInfo(context.Context, *JobGetInfoRequest, ...dcerpc.CallOption) (*JobGetInfoResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// ATEnumContainer structure represents AT_ENUM_CONTAINER RPC structure.
//
// The ATSvc method NetrJobEnum uses the AT_ENUM_CONTAINER structure to return multiple
// AT_ENUM structures. The format of the AT_ENUM_CONTAINER structure is as follows:
type ATEnumContainer struct {
	// EntriesRead:  The number of entries in the Buffer array.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  Pointer to an array of AT_ENUM structures.
	Buffer []*tsch.ATEnum `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *ATEnumContainer) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ATEnumContainer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
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
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&tsch.ATEnum{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&tsch.ATEnum{}).MarshalNDR(ctx, w); err != nil {
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
func (o *ATEnumContainer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
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
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*tsch.ATEnum, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &tsch.ATEnum{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*tsch.ATEnum) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultATSvcClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultATSvcClient) JobAdd(ctx context.Context, in *JobAddRequest, opts ...dcerpc.CallOption) (*JobAddResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &JobAddResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultATSvcClient) JobDelete(ctx context.Context, in *JobDeleteRequest, opts ...dcerpc.CallOption) (*JobDeleteResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &JobDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultATSvcClient) JobEnum(ctx context.Context, in *JobEnumRequest, opts ...dcerpc.CallOption) (*JobEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &JobEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultATSvcClient) JobGetInfo(ctx context.Context, in *JobGetInfoRequest, opts ...dcerpc.CallOption) (*JobGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &JobGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultATSvcClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultATSvcClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewATSvcClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ATSvcClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ATSvcSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultATSvcClient{cc: cc}, nil
}

// xxx_JobAddOperation structure represents the NetrJobAdd operation
type xxx_JobAddOperation struct {
	ServerName string       `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	ATInfo     *tsch.ATInfo `idl:"name:pAtInfo" json:"at_info"`
	JobID      uint32       `idl:"name:pJobId" json:"job_id"`
	Return     uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_JobAddOperation) OpNum() int { return 0 }

func (o *xxx_JobAddOperation) OpName() string { return "/atsvc/v1/NetrJobAdd" }

func (o *xxx_JobAddOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobAddOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=ATSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// pAtInfo {in} (1:{alias=LPAT_INFO}*(1))(2:{alias=AT_INFO}(struct))
	{
		if o.ATInfo != nil {
			if err := o.ATInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsch.ATInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobAddOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=ATSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pAtInfo {in} (1:{alias=LPAT_INFO,pointer=ref}*(1))(2:{alias=AT_INFO}(struct))
	{
		if o.ATInfo == nil {
			o.ATInfo = &tsch.ATInfo{}
		}
		if err := o.ATInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobAddOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobAddOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pJobId {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.JobID); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobAddOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pJobId {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.JobID); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// JobAddRequest structure represents the NetrJobAdd operation request
type JobAddRequest struct {
	// ServerName: Pointer to a Unicode string that MUST specify the server. The client
	// MUST map this string to an RPC binding handle. The server MUST ignore this parameter.
	// For more information, see [C706] sections 4.3.5 and 5.1.5.2.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// pAtInfo:  Pointer to an AT_INFO structure (section 2.3.4) that MUST contain the
	// task configuration.
	ATInfo *tsch.ATInfo `idl:"name:pAtInfo" json:"at_info"`
}

func (o *JobAddRequest) xxx_ToOp(ctx context.Context, op *xxx_JobAddOperation) *xxx_JobAddOperation {
	if op == nil {
		op = &xxx_JobAddOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.ATInfo = o.ATInfo
	return op
}

func (o *JobAddRequest) xxx_FromOp(ctx context.Context, op *xxx_JobAddOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.ATInfo = op.ATInfo
}
func (o *JobAddRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *JobAddRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_JobAddOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// JobAddResponse structure represents the NetrJobAdd operation response
type JobAddResponse struct {
	// pJobId: MUST return a pointer to the task identifier when the NetrJobAdd method is
	// successful.
	JobID uint32 `idl:"name:pJobId" json:"job_id"`
	// Return: The NetrJobAdd return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *JobAddResponse) xxx_ToOp(ctx context.Context, op *xxx_JobAddOperation) *xxx_JobAddOperation {
	if op == nil {
		op = &xxx_JobAddOperation{}
	}
	if o == nil {
		return op
	}
	op.JobID = o.JobID
	op.Return = o.Return
	return op
}

func (o *JobAddResponse) xxx_FromOp(ctx context.Context, op *xxx_JobAddOperation) {
	if o == nil {
		return
	}
	o.JobID = op.JobID
	o.Return = op.Return
}
func (o *JobAddResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *JobAddResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_JobAddOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_JobDeleteOperation structure represents the NetrJobDel operation
type xxx_JobDeleteOperation struct {
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	MinJobID   uint32 `idl:"name:MinJobId" json:"min_job_id"`
	MaxJobID   uint32 `idl:"name:MaxJobId" json:"max_job_id"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_JobDeleteOperation) OpNum() int { return 1 }

func (o *xxx_JobDeleteOperation) OpName() string { return "/atsvc/v1/NetrJobDel" }

func (o *xxx_JobDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=ATSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// MinJobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MinJobID); err != nil {
			return err
		}
	}
	// MaxJobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxJobID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=ATSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MinJobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MinJobID); err != nil {
			return err
		}
	}
	// MaxJobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxJobID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// JobDeleteRequest structure represents the NetrJobDel operation request
type JobDeleteRequest struct {
	// ServerName:  Pointer to a Unicode string that MUST specify the server. The client
	// MUST map this string to an RPC binding handle. The server MUST ignore this parameter.
	// For more information, see [C706] sections 4.3.5 and 5.1.5.2.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// MinJobId: MUST specify the low end of the range of tasks to be deleted. This parameter
	// MUST NOT be greater than MaxJobId.
	MinJobID uint32 `idl:"name:MinJobId" json:"min_job_id"`
	// MaxJobId: MUST specify the high end of the range of tasks to be deleted. This parameter
	// MUST NOT be smaller than MinJobId.
	MaxJobID uint32 `idl:"name:MaxJobId" json:"max_job_id"`
}

func (o *JobDeleteRequest) xxx_ToOp(ctx context.Context, op *xxx_JobDeleteOperation) *xxx_JobDeleteOperation {
	if op == nil {
		op = &xxx_JobDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.MinJobID = o.MinJobID
	op.MaxJobID = o.MaxJobID
	return op
}

func (o *JobDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_JobDeleteOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.MinJobID = op.MinJobID
	o.MaxJobID = op.MaxJobID
}
func (o *JobDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *JobDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_JobDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// JobDeleteResponse structure represents the NetrJobDel operation response
type JobDeleteResponse struct {
	// Return: The NetrJobDel return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *JobDeleteResponse) xxx_ToOp(ctx context.Context, op *xxx_JobDeleteOperation) *xxx_JobDeleteOperation {
	if op == nil {
		op = &xxx_JobDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *JobDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_JobDeleteOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *JobDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *JobDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_JobDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_JobEnumOperation structure represents the NetrJobEnum operation
type xxx_JobEnumOperation struct {
	ServerName             string           `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	EnumContainer          *ATEnumContainer `idl:"name:pEnumContainer" json:"enum_container"`
	PreferredMaximumLength uint32           `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
	TotalEntries           uint32           `idl:"name:pTotalEntries" json:"total_entries"`
	Resume                 uint32           `idl:"name:pResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32           `idl:"name:Return" json:"return"`
}

func (o *xxx_JobEnumOperation) OpNum() int { return 2 }

func (o *xxx_JobEnumOperation) OpName() string { return "/atsvc/v1/NetrJobEnum" }

func (o *xxx_JobEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=ATSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// pEnumContainer {in, out} (1:{alias=LPAT_ENUM_CONTAINER}*(1))(2:{alias=AT_ENUM_CONTAINER}(struct))
	{
		if o.EnumContainer != nil {
			if err := o.EnumContainer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ATEnumContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// PreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// pResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_pResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_pResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=ATSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pEnumContainer {in, out} (1:{alias=LPAT_ENUM_CONTAINER,pointer=ref}*(1))(2:{alias=AT_ENUM_CONTAINER}(struct))
	{
		if o.EnumContainer == nil {
			o.EnumContainer = &ATEnumContainer{}
		}
		if err := o.EnumContainer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// pResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_pResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_pResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_pResumeHandle, _ptr_pResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pEnumContainer {in, out} (1:{alias=LPAT_ENUM_CONTAINER}*(1))(2:{alias=AT_ENUM_CONTAINER}(struct))
	{
		if o.EnumContainer != nil {
			if err := o.EnumContainer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ATEnumContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pTotalEntries {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// pResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_pResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_pResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pEnumContainer {in, out} (1:{alias=LPAT_ENUM_CONTAINER,pointer=ref}*(1))(2:{alias=AT_ENUM_CONTAINER}(struct))
	{
		if o.EnumContainer == nil {
			o.EnumContainer = &ATEnumContainer{}
		}
		if err := o.EnumContainer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pTotalEntries {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// pResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_pResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_pResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_pResumeHandle, _ptr_pResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// JobEnumRequest structure represents the NetrJobEnum operation request
type JobEnumRequest struct {
	// ServerName: Pointer to a Unicode string that MUST specify the server. The client
	// MUST map this string to an RPC binding handle. The server MUST ignore this parameter.
	// For more information, see [C706] sections 4.3.5 and 5.1.5.2.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// pEnumContainer:  Pointer to an AT_ENUM_CONTAINER (section 2.3.5) structure that
	// MUST contain a count of the number of entries returned and a buffer that contains
	// the entries. The client MUST send a pointer to this structure to the server with
	// the Buffer field set to NULL; upon return the Buffer field MUST contain a pointer
	// to an array of AT_ENUM structures.
	EnumContainer *ATEnumContainer `idl:"name:pEnumContainer" json:"enum_container"`
	// PreferedMaximumLength: MUST contain the preferred maximum length, in bytes, of data
	// to be returned. A value of 0xFFFFFFFF MUST indicate no preferred maximum length.
	PreferredMaximumLength uint32 `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
	// pResumeHandle: MUST be a pointer to an index into the list of tasks. This value indicates
	// the index number at which the enumeration MUST start.
	Resume uint32 `idl:"name:pResumeHandle;pointer:unique" json:"resume"`
}

func (o *JobEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_JobEnumOperation) *xxx_JobEnumOperation {
	if op == nil {
		op = &xxx_JobEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.EnumContainer = o.EnumContainer
	op.PreferredMaximumLength = o.PreferredMaximumLength
	op.Resume = o.Resume
	return op
}

func (o *JobEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_JobEnumOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.EnumContainer = op.EnumContainer
	o.PreferredMaximumLength = op.PreferredMaximumLength
	o.Resume = op.Resume
}
func (o *JobEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *JobEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_JobEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// JobEnumResponse structure represents the NetrJobEnum operation response
type JobEnumResponse struct {
	// pEnumContainer:  Pointer to an AT_ENUM_CONTAINER (section 2.3.5) structure that
	// MUST contain a count of the number of entries returned and a buffer that contains
	// the entries. The client MUST send a pointer to this structure to the server with
	// the Buffer field set to NULL; upon return the Buffer field MUST contain a pointer
	// to an array of AT_ENUM structures.
	EnumContainer *ATEnumContainer `idl:"name:pEnumContainer" json:"enum_container"`
	// pTotalEntries: Pointer to a value that MUST receive the total number of tasks in
	// the list, beyond the position specified by pResumeHandle.
	TotalEntries uint32 `idl:"name:pTotalEntries" json:"total_entries"`
	// pResumeHandle: MUST be a pointer to an index into the list of tasks. This value indicates
	// the index number at which the enumeration MUST start.
	Resume uint32 `idl:"name:pResumeHandle;pointer:unique" json:"resume"`
	// Return: The NetrJobEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *JobEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_JobEnumOperation) *xxx_JobEnumOperation {
	if op == nil {
		op = &xxx_JobEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.EnumContainer = o.EnumContainer
	op.TotalEntries = o.TotalEntries
	op.Resume = o.Resume
	op.Return = o.Return
	return op
}

func (o *JobEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_JobEnumOperation) {
	if o == nil {
		return
	}
	o.EnumContainer = op.EnumContainer
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *JobEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *JobEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_JobEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_JobGetInfoOperation structure represents the NetrJobGetInfo operation
type xxx_JobGetInfoOperation struct {
	ServerName string       `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	JobID      uint32       `idl:"name:JobId" json:"job_id"`
	ATInfo     *tsch.ATInfo `idl:"name:ppAtInfo" json:"at_info"`
	Return     uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_JobGetInfoOperation) OpNum() int { return 3 }

func (o *xxx_JobGetInfoOperation) OpName() string { return "/atsvc/v1/NetrJobGetInfo" }

func (o *xxx_JobGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=ATSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// JobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.JobID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=ATSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// JobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.JobID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppAtInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPAT_INFO}*(1))(3:{alias=AT_INFO}(struct))
	{
		if o.ATInfo != nil {
			_ptr_ppAtInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ATInfo != nil {
					if err := o.ATInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&tsch.ATInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ATInfo, _ptr_ppAtInfo); err != nil {
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
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JobGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppAtInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPAT_INFO,pointer=ref}*(1))(3:{alias=AT_INFO}(struct))
	{
		_ptr_ppAtInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ATInfo == nil {
				o.ATInfo = &tsch.ATInfo{}
			}
			if err := o.ATInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppAtInfo := func(ptr interface{}) { o.ATInfo = *ptr.(**tsch.ATInfo) }
		if err := w.ReadPointer(&o.ATInfo, _s_ppAtInfo, _ptr_ppAtInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// JobGetInfoRequest structure represents the NetrJobGetInfo operation request
type JobGetInfoRequest struct {
	// ServerName: Pointer to a Unicode string that MUST specify the server. The client
	// MUST map this string to an RPC binding handle. The server MUST ignore this parameter.
	// For more information, see [C706] sections 4.3.5 and 5.1.5.2.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// JobId: MUST contain a task identifier.
	JobID uint32 `idl:"name:JobId" json:"job_id"`
}

func (o *JobGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_JobGetInfoOperation) *xxx_JobGetInfoOperation {
	if op == nil {
		op = &xxx_JobGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.JobID = o.JobID
	return op
}

func (o *JobGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_JobGetInfoOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.JobID = op.JobID
}
func (o *JobGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *JobGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_JobGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// JobGetInfoResponse structure represents the NetrJobGetInfo operation response
type JobGetInfoResponse struct {
	// ppAtInfo:  MUST be a pointer to an AT_INFO structure as specified in section 2.3.4.
	ATInfo *tsch.ATInfo `idl:"name:ppAtInfo" json:"at_info"`
	// Return: The NetrJobGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *JobGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_JobGetInfoOperation) *xxx_JobGetInfoOperation {
	if op == nil {
		op = &xxx_JobGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.ATInfo = o.ATInfo
	op.Return = o.Return
	return op
}

func (o *JobGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_JobGetInfoOperation) {
	if o == nil {
		return
	}
	o.ATInfo = op.ATInfo
	o.Return = op.Return
}
func (o *JobGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *JobGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_JobGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
