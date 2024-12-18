package ifsrmreportscheduler

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmReportScheduler interface identifier 6879caf9-6617-4484-8719-71c3d8645f94
	ReportSchedulerIID = &dcom.IID{Data1: 0x6879caf9, Data2: 0x6617, Data3: 0x4484, Data4: []byte{0x87, 0x19, 0x71, 0xc3, 0xd8, 0x64, 0x5f, 0x94}}
	// Syntax UUID
	ReportSchedulerSyntaxUUID = &uuid.UUID{TimeLow: 0x6879caf9, TimeMid: 0x6617, TimeHiAndVersion: 0x4484, ClockSeqHiAndReserved: 0x87, ClockSeqLow: 0x19, Node: [6]uint8{0x71, 0xc3, 0xd8, 0x64, 0x5f, 0x94}}
	// Syntax ID
	ReportSchedulerSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: ReportSchedulerSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// IFsrmReportScheduler interface.
type ReportSchedulerClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The VerifyNamespaces method checks that all namespaces passed in exist and are valid
	// paths.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045304 FSRM_E_PATH_NOT_FOUND      | The specified path could not be found.                                           |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH        | The supplied path is not valid.                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045315 FSRM_E_VOLUME_UNSUPPORTED  | The volume is not an NTFS volume.                                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004531F FSRM_E_FILE_SYSTEM_CORRUPT | The file system is corrupt.                                                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG               | This code is returned for the following reasons: The namespacesSafeArray         |
	//	|                                       | parameter is NULL. The namespacesSafeArray parameter does not contain a variant  |
	//	|                                       | array of BSTRs.                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	VerifyNamespaces(context.Context, *VerifyNamespacesRequest, ...dcerpc.CallOption) (*VerifyNamespacesResponse, error)

	// The CreateScheduleTask method creates a task that is capable of triggering the generation
	// of a File Server Resource Manager Protocol report job.
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
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED  | This function is not supported for this object.                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045316 FSRM_E_UNEXPECTED     | An unexpected error occurred; check the application event log. You might get     |
	//	|                                  | this error if the XML is malformed.                                              |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The taskName parameter is NULL. |
	//	|                                  | The namespacesSafeArray parameter is NULL. The serializedTask parameter is NULL. |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	CreateScheduleTask(context.Context, *CreateScheduleTaskRequest, ...dcerpc.CallOption) (*CreateScheduleTaskResponse, error)

	// The ModifyScheduleTask method modifies the specified task that is capable of triggering
	// the generation of a File Server Resource Manager Protocol report job.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|           RETURN            |                                                                                  |
	//	|         VALUE/CODE          |                                   DESCRIPTION                                    |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | The specified task could not be found.                                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG     | This code is returned for the following reasons: The taskName parameter is NULL. |
	//	|                             | The serializedTask parameter is NULL. The taskName parameter is not a valid      |
	//	|                             | value.                                                                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	ModifyScheduleTask(context.Context, *ModifyScheduleTaskRequest, ...dcerpc.CallOption) (*ModifyScheduleTaskResponse, error)

	// The DeleteScheduleTask method deletes the specified task so that it no longer triggers
	// the generation of a File Server Resource Manager Protocol report job.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+----------------------------------------+
	//	|           RETURN            |                                        |
	//	|         VALUE/CODE          |              DESCRIPTION               |
	//	|                             |                                        |
	//	+-----------------------------+----------------------------------------+
	//	+-----------------------------+----------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | The specified task could not be found. |
	//	+-----------------------------+----------------------------------------+
	//	| 0x80070057 E_INVALIDARG     | The taskName parameter is NULL.        |
	//	+-----------------------------+----------------------------------------+
	DeleteScheduleTask(context.Context, *DeleteScheduleTaskRequest, ...dcerpc.CallOption) (*DeleteScheduleTaskResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ReportSchedulerClient
}

type xxx_DefaultReportSchedulerClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultReportSchedulerClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultReportSchedulerClient) VerifyNamespaces(ctx context.Context, in *VerifyNamespacesRequest, opts ...dcerpc.CallOption) (*VerifyNamespacesResponse, error) {
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
	out := &VerifyNamespacesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportSchedulerClient) CreateScheduleTask(ctx context.Context, in *CreateScheduleTaskRequest, opts ...dcerpc.CallOption) (*CreateScheduleTaskResponse, error) {
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
	out := &CreateScheduleTaskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportSchedulerClient) ModifyScheduleTask(ctx context.Context, in *ModifyScheduleTaskRequest, opts ...dcerpc.CallOption) (*ModifyScheduleTaskResponse, error) {
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
	out := &ModifyScheduleTaskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportSchedulerClient) DeleteScheduleTask(ctx context.Context, in *DeleteScheduleTaskRequest, opts ...dcerpc.CallOption) (*DeleteScheduleTaskResponse, error) {
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
	out := &DeleteScheduleTaskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportSchedulerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultReportSchedulerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultReportSchedulerClient) IPID(ctx context.Context, ipid *dcom.IPID) ReportSchedulerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultReportSchedulerClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewReportSchedulerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ReportSchedulerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ReportSchedulerSyntaxV1_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultReportSchedulerClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_VerifyNamespacesOperation structure represents the VerifyNamespaces operation
type xxx_VerifyNamespacesOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	NamespacesSafeArray *oaut.Variant  `idl:"name:namespacesSafeArray" json:"namespaces_safe_array"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_VerifyNamespacesOperation) OpNum() int { return 7 }

func (o *xxx_VerifyNamespacesOperation) OpName() string {
	return "/IFsrmReportScheduler/v1/VerifyNamespaces"
}

func (o *xxx_VerifyNamespacesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_VerifyNamespacesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// namespacesSafeArray {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.NamespacesSafeArray != nil {
			_ptr_namespacesSafeArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NamespacesSafeArray != nil {
					if err := o.NamespacesSafeArray.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NamespacesSafeArray, _ptr_namespacesSafeArray); err != nil {
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

func (o *xxx_VerifyNamespacesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// namespacesSafeArray {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_namespacesSafeArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NamespacesSafeArray == nil {
				o.NamespacesSafeArray = &oaut.Variant{}
			}
			if err := o.NamespacesSafeArray.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_namespacesSafeArray := func(ptr interface{}) { o.NamespacesSafeArray = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.NamespacesSafeArray, _s_namespacesSafeArray, _ptr_namespacesSafeArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_VerifyNamespacesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_VerifyNamespacesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_VerifyNamespacesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// VerifyNamespacesRequest structure represents the VerifyNamespaces operation request
type VerifyNamespacesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// namespacesSafeArray: Pointer to a variable that contains a VARIANT structure, which
	// contains a SAFEARRAY of VARIANT structures. The VARIANT structures contained in the
	// SAFEARRAY MUST be BSTR string values, each representing a local directory path that
	// needs to be verified as supported by the File Server Resource Manager Protocol storage
	// reports.
	NamespacesSafeArray *oaut.Variant `idl:"name:namespacesSafeArray" json:"namespaces_safe_array"`
}

func (o *VerifyNamespacesRequest) xxx_ToOp(ctx context.Context) *xxx_VerifyNamespacesOperation {
	if o == nil {
		return &xxx_VerifyNamespacesOperation{}
	}
	return &xxx_VerifyNamespacesOperation{
		This:                o.This,
		NamespacesSafeArray: o.NamespacesSafeArray,
	}
}

func (o *VerifyNamespacesRequest) xxx_FromOp(ctx context.Context, op *xxx_VerifyNamespacesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NamespacesSafeArray = op.NamespacesSafeArray
}
func (o *VerifyNamespacesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *VerifyNamespacesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_VerifyNamespacesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// VerifyNamespacesResponse structure represents the VerifyNamespaces operation response
type VerifyNamespacesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The VerifyNamespaces return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *VerifyNamespacesResponse) xxx_ToOp(ctx context.Context) *xxx_VerifyNamespacesOperation {
	if o == nil {
		return &xxx_VerifyNamespacesOperation{}
	}
	return &xxx_VerifyNamespacesOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *VerifyNamespacesResponse) xxx_FromOp(ctx context.Context, op *xxx_VerifyNamespacesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *VerifyNamespacesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *VerifyNamespacesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_VerifyNamespacesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateScheduleTaskOperation structure represents the CreateScheduleTask operation
type xxx_CreateScheduleTaskOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskName            *oaut.String   `idl:"name:taskName" json:"task_name"`
	NamespacesSafeArray *oaut.Variant  `idl:"name:namespacesSafeArray" json:"namespaces_safe_array"`
	SerializedTask      *oaut.String   `idl:"name:serializedTask" json:"serialized_task"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateScheduleTaskOperation) OpNum() int { return 8 }

func (o *xxx_CreateScheduleTaskOperation) OpName() string {
	return "/IFsrmReportScheduler/v1/CreateScheduleTask"
}

func (o *xxx_CreateScheduleTaskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateScheduleTaskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// taskName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.TaskName != nil {
			_ptr_taskName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TaskName != nil {
					if err := o.TaskName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TaskName, _ptr_taskName); err != nil {
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
	// namespacesSafeArray {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.NamespacesSafeArray != nil {
			_ptr_namespacesSafeArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NamespacesSafeArray != nil {
					if err := o.NamespacesSafeArray.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NamespacesSafeArray, _ptr_namespacesSafeArray); err != nil {
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
	// serializedTask {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SerializedTask != nil {
			_ptr_serializedTask := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SerializedTask != nil {
					if err := o.SerializedTask.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SerializedTask, _ptr_serializedTask); err != nil {
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

func (o *xxx_CreateScheduleTaskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// taskName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_taskName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TaskName == nil {
				o.TaskName = &oaut.String{}
			}
			if err := o.TaskName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_taskName := func(ptr interface{}) { o.TaskName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.TaskName, _s_taskName, _ptr_taskName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// namespacesSafeArray {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_namespacesSafeArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NamespacesSafeArray == nil {
				o.NamespacesSafeArray = &oaut.Variant{}
			}
			if err := o.NamespacesSafeArray.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_namespacesSafeArray := func(ptr interface{}) { o.NamespacesSafeArray = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.NamespacesSafeArray, _s_namespacesSafeArray, _ptr_namespacesSafeArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// serializedTask {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_serializedTask := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SerializedTask == nil {
				o.SerializedTask = &oaut.String{}
			}
			if err := o.SerializedTask.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_serializedTask := func(ptr interface{}) { o.SerializedTask = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SerializedTask, _s_serializedTask, _ptr_serializedTask); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateScheduleTaskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateScheduleTaskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateScheduleTaskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateScheduleTaskRequest structure represents the CreateScheduleTask operation request
type CreateScheduleTaskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// taskName: Contains the name of the Task Scheduler task to create.
	TaskName *oaut.String `idl:"name:taskName" json:"task_name"`
	// namespacesSafeArray: Pointer to a variable that contains a VARIANT structure, which
	// contains a SAFEARRAY of VARIANT structures. The VARIANT structures contained in the
	// SAFEARRAY MUST be BSTR string values, each representing a local directory path that
	// needs to be verified as supported by the File Server Resource Manager Protocol storage.
	NamespacesSafeArray *oaut.Variant `idl:"name:namespacesSafeArray" json:"namespaces_safe_array"`
	// serializedTask: Contains the XML representation of the task to be created.<63>
	SerializedTask *oaut.String `idl:"name:serializedTask" json:"serialized_task"`
}

func (o *CreateScheduleTaskRequest) xxx_ToOp(ctx context.Context) *xxx_CreateScheduleTaskOperation {
	if o == nil {
		return &xxx_CreateScheduleTaskOperation{}
	}
	return &xxx_CreateScheduleTaskOperation{
		This:                o.This,
		TaskName:            o.TaskName,
		NamespacesSafeArray: o.NamespacesSafeArray,
		SerializedTask:      o.SerializedTask,
	}
}

func (o *CreateScheduleTaskRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateScheduleTaskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TaskName = op.TaskName
	o.NamespacesSafeArray = op.NamespacesSafeArray
	o.SerializedTask = op.SerializedTask
}
func (o *CreateScheduleTaskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateScheduleTaskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateScheduleTaskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateScheduleTaskResponse structure represents the CreateScheduleTask operation response
type CreateScheduleTaskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateScheduleTask return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateScheduleTaskResponse) xxx_ToOp(ctx context.Context) *xxx_CreateScheduleTaskOperation {
	if o == nil {
		return &xxx_CreateScheduleTaskOperation{}
	}
	return &xxx_CreateScheduleTaskOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CreateScheduleTaskResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateScheduleTaskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateScheduleTaskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateScheduleTaskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateScheduleTaskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyScheduleTaskOperation structure represents the ModifyScheduleTask operation
type xxx_ModifyScheduleTaskOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskName            *oaut.String   `idl:"name:taskName" json:"task_name"`
	NamespacesSafeArray *oaut.Variant  `idl:"name:namespacesSafeArray" json:"namespaces_safe_array"`
	SerializedTask      *oaut.String   `idl:"name:serializedTask" json:"serialized_task"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyScheduleTaskOperation) OpNum() int { return 9 }

func (o *xxx_ModifyScheduleTaskOperation) OpName() string {
	return "/IFsrmReportScheduler/v1/ModifyScheduleTask"
}

func (o *xxx_ModifyScheduleTaskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyScheduleTaskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// taskName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.TaskName != nil {
			_ptr_taskName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TaskName != nil {
					if err := o.TaskName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TaskName, _ptr_taskName); err != nil {
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
	// namespacesSafeArray {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.NamespacesSafeArray != nil {
			_ptr_namespacesSafeArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NamespacesSafeArray != nil {
					if err := o.NamespacesSafeArray.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NamespacesSafeArray, _ptr_namespacesSafeArray); err != nil {
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
	// serializedTask {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SerializedTask != nil {
			_ptr_serializedTask := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SerializedTask != nil {
					if err := o.SerializedTask.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SerializedTask, _ptr_serializedTask); err != nil {
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

func (o *xxx_ModifyScheduleTaskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// taskName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_taskName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TaskName == nil {
				o.TaskName = &oaut.String{}
			}
			if err := o.TaskName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_taskName := func(ptr interface{}) { o.TaskName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.TaskName, _s_taskName, _ptr_taskName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// namespacesSafeArray {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_namespacesSafeArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NamespacesSafeArray == nil {
				o.NamespacesSafeArray = &oaut.Variant{}
			}
			if err := o.NamespacesSafeArray.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_namespacesSafeArray := func(ptr interface{}) { o.NamespacesSafeArray = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.NamespacesSafeArray, _s_namespacesSafeArray, _ptr_namespacesSafeArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// serializedTask {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_serializedTask := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SerializedTask == nil {
				o.SerializedTask = &oaut.String{}
			}
			if err := o.SerializedTask.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_serializedTask := func(ptr interface{}) { o.SerializedTask = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SerializedTask, _s_serializedTask, _ptr_serializedTask); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyScheduleTaskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyScheduleTaskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ModifyScheduleTaskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ModifyScheduleTaskRequest structure represents the ModifyScheduleTask operation request
type ModifyScheduleTaskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// taskName: Contains the name of the Task Scheduler task to modify.
	TaskName *oaut.String `idl:"name:taskName" json:"task_name"`
	// namespacesSafeArray: Pointer to a variable that contains a VARIANT structure, which
	// contains a SAFEARRAY of VARIANT structures. The VARIANT structures contained in the
	// SAFEARRAY MUST be BSTR string values, each representing a local directory path that
	// needs to be verified as supported by the File Server Resource Manager Protocol storage.
	NamespacesSafeArray *oaut.Variant `idl:"name:namespacesSafeArray" json:"namespaces_safe_array"`
	// serializedTask: Contains the XML representation of the task to be modified.<64>
	SerializedTask *oaut.String `idl:"name:serializedTask" json:"serialized_task"`
}

func (o *ModifyScheduleTaskRequest) xxx_ToOp(ctx context.Context) *xxx_ModifyScheduleTaskOperation {
	if o == nil {
		return &xxx_ModifyScheduleTaskOperation{}
	}
	return &xxx_ModifyScheduleTaskOperation{
		This:                o.This,
		TaskName:            o.TaskName,
		NamespacesSafeArray: o.NamespacesSafeArray,
		SerializedTask:      o.SerializedTask,
	}
}

func (o *ModifyScheduleTaskRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyScheduleTaskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TaskName = op.TaskName
	o.NamespacesSafeArray = op.NamespacesSafeArray
	o.SerializedTask = op.SerializedTask
}
func (o *ModifyScheduleTaskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ModifyScheduleTaskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyScheduleTaskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyScheduleTaskResponse structure represents the ModifyScheduleTask operation response
type ModifyScheduleTaskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModifyScheduleTask return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyScheduleTaskResponse) xxx_ToOp(ctx context.Context) *xxx_ModifyScheduleTaskOperation {
	if o == nil {
		return &xxx_ModifyScheduleTaskOperation{}
	}
	return &xxx_ModifyScheduleTaskOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ModifyScheduleTaskResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyScheduleTaskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ModifyScheduleTaskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ModifyScheduleTaskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyScheduleTaskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteScheduleTaskOperation structure represents the DeleteScheduleTask operation
type xxx_DeleteScheduleTaskOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskName *oaut.String   `idl:"name:taskName" json:"task_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteScheduleTaskOperation) OpNum() int { return 10 }

func (o *xxx_DeleteScheduleTaskOperation) OpName() string {
	return "/IFsrmReportScheduler/v1/DeleteScheduleTask"
}

func (o *xxx_DeleteScheduleTaskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteScheduleTaskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// taskName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.TaskName != nil {
			_ptr_taskName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TaskName != nil {
					if err := o.TaskName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TaskName, _ptr_taskName); err != nil {
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

func (o *xxx_DeleteScheduleTaskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// taskName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_taskName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TaskName == nil {
				o.TaskName = &oaut.String{}
			}
			if err := o.TaskName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_taskName := func(ptr interface{}) { o.TaskName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.TaskName, _s_taskName, _ptr_taskName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteScheduleTaskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteScheduleTaskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteScheduleTaskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteScheduleTaskRequest structure represents the DeleteScheduleTask operation request
type DeleteScheduleTaskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// taskName: Contains the name of the task to delete.
	TaskName *oaut.String `idl:"name:taskName" json:"task_name"`
}

func (o *DeleteScheduleTaskRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteScheduleTaskOperation {
	if o == nil {
		return &xxx_DeleteScheduleTaskOperation{}
	}
	return &xxx_DeleteScheduleTaskOperation{
		This:     o.This,
		TaskName: o.TaskName,
	}
}

func (o *DeleteScheduleTaskRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteScheduleTaskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TaskName = op.TaskName
}
func (o *DeleteScheduleTaskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteScheduleTaskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteScheduleTaskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteScheduleTaskResponse structure represents the DeleteScheduleTask operation response
type DeleteScheduleTaskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteScheduleTask return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteScheduleTaskResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteScheduleTaskOperation {
	if o == nil {
		return &xxx_DeleteScheduleTaskOperation{}
	}
	return &xxx_DeleteScheduleTaskOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeleteScheduleTaskResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteScheduleTaskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteScheduleTaskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteScheduleTaskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteScheduleTaskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
