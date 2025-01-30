package ifsrmfilemanagementjobmanager

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
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmFileManagementJobManager interface identifier ee321ecb-d95e-48e9-907c-c7685a013235
	FileManagementJobManagerIID = &dcom.IID{Data1: 0xee321ecb, Data2: 0xd95e, Data3: 0x48e9, Data4: []byte{0x90, 0x7c, 0xc7, 0x68, 0x5a, 0x01, 0x32, 0x35}}
	// Syntax UUID
	FileManagementJobManagerSyntaxUUID = &uuid.UUID{TimeLow: 0xee321ecb, TimeMid: 0xd95e, TimeHiAndVersion: 0x48e9, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0x7c, Node: [6]uint8{0xc7, 0x68, 0x5a, 0x1, 0x32, 0x35}}
	// Syntax ID
	FileManagementJobManagerSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: FileManagementJobManagerSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// IFsrmFileManagementJobManager interface.
type FileManagementJobManagerClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// ActionVariables operation.
	GetActionVariables(context.Context, *GetActionVariablesRequest, ...dcerpc.CallOption) (*GetActionVariablesResponse, error)

	// ActionVariableDescriptions operation.
	GetActionVariableDescriptions(context.Context, *GetActionVariableDescriptionsRequest, ...dcerpc.CallOption) (*GetActionVariableDescriptionsResponse, error)

	// The EnumFileManagementJobs method returns all the fileManagementJobs from the List
	// of Persisted File Management Jobs (section 3.2.1.7) on the server.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | This code is returned for the following reasons: The fileManagementJobs          |
	//	|                                 | parameter is NULL.                                                               |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | The options parameter does not contain a valid FsrmEnumOptions (section          |
	//	|                                 | 2.2.1.2.5) value.                                                                |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumFileManagementJobs(context.Context, *EnumFileManagementJobsRequest, ...dcerpc.CallOption) (*EnumFileManagementJobsResponse, error)

	// The CreateFileManagementJob method creates a blank Non-Persisted File Management
	// Job Instance (section 3.2.1.7.1.2) and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+------------------------------------------+
	//	|         RETURN          |                                          |
	//	|       VALUE/CODE        |               DESCRIPTION                |
	//	|                         |                                          |
	//	+-------------------------+------------------------------------------+
	//	+-------------------------+------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The fileManagementJob parameter is NULL. |
	//	+-------------------------+------------------------------------------+
	CreateFileManagementJob(context.Context, *CreateFileManagementJobRequest, ...dcerpc.CallOption) (*CreateFileManagementJobResponse, error)

	// The GetFileManagementJob method returns a pointer to the fileManagementJob with the
	// specified name from the List of Persisted File Management Jobs (section 3.2.1.7)
	// and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+-------------------------------------------------------+
	//	|           RETURN            |                                                       |
	//	|         VALUE/CODE          |                      DESCRIPTION                      |
	//	|                             |                                                       |
	//	+-----------------------------+-------------------------------------------------------+
	//	+-----------------------------+-------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | The specified file management job could not be found. |
	//	+-----------------------------+-------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG     | The fileManagementJob parameter is NULL.              |
	//	+-----------------------------+-------------------------------------------------------+
	GetFileManagementJob(context.Context, *GetFileManagementJobRequest, ...dcerpc.CallOption) (*GetFileManagementJobResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) FileManagementJobManagerClient
}

type xxx_DefaultFileManagementJobManagerClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultFileManagementJobManagerClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultFileManagementJobManagerClient) GetActionVariables(ctx context.Context, in *GetActionVariablesRequest, opts ...dcerpc.CallOption) (*GetActionVariablesResponse, error) {
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
	out := &GetActionVariablesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobManagerClient) GetActionVariableDescriptions(ctx context.Context, in *GetActionVariableDescriptionsRequest, opts ...dcerpc.CallOption) (*GetActionVariableDescriptionsResponse, error) {
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
	out := &GetActionVariableDescriptionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobManagerClient) EnumFileManagementJobs(ctx context.Context, in *EnumFileManagementJobsRequest, opts ...dcerpc.CallOption) (*EnumFileManagementJobsResponse, error) {
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
	out := &EnumFileManagementJobsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobManagerClient) CreateFileManagementJob(ctx context.Context, in *CreateFileManagementJobRequest, opts ...dcerpc.CallOption) (*CreateFileManagementJobResponse, error) {
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
	out := &CreateFileManagementJobResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobManagerClient) GetFileManagementJob(ctx context.Context, in *GetFileManagementJobRequest, opts ...dcerpc.CallOption) (*GetFileManagementJobResponse, error) {
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
	out := &GetFileManagementJobResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileManagementJobManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultFileManagementJobManagerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultFileManagementJobManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) FileManagementJobManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultFileManagementJobManagerClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewFileManagementJobManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (FileManagementJobManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(FileManagementJobManagerSyntaxV1_0))...)
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
	return &xxx_DefaultFileManagementJobManagerClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetActionVariablesOperation structure represents the ActionVariables operation
type xxx_GetActionVariablesOperation struct {
	This      *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Variables *oaut.SafeArray `idl:"name:variables" json:"variables"`
	Return    int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetActionVariablesOperation) OpNum() int { return 7 }

func (o *xxx_GetActionVariablesOperation) OpName() string {
	return "/IFsrmFileManagementJobManager/v1/ActionVariables"
}

func (o *xxx_GetActionVariablesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionVariablesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetActionVariablesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetActionVariablesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionVariablesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// variables {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Variables != nil {
			_ptr_variables := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Variables != nil {
					if err := o.Variables.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Variables, _ptr_variables); err != nil {
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

func (o *xxx_GetActionVariablesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// variables {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_variables := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Variables == nil {
				o.Variables = &oaut.SafeArray{}
			}
			if err := o.Variables.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_variables := func(ptr interface{}) { o.Variables = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Variables, _s_variables, _ptr_variables); err != nil {
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

// GetActionVariablesRequest structure represents the ActionVariables operation request
type GetActionVariablesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetActionVariablesRequest) xxx_ToOp(ctx context.Context) *xxx_GetActionVariablesOperation {
	if o == nil {
		return &xxx_GetActionVariablesOperation{}
	}
	return &xxx_GetActionVariablesOperation{
		This: o.This,
	}
}

func (o *GetActionVariablesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetActionVariablesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetActionVariablesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetActionVariablesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionVariablesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetActionVariablesResponse structure represents the ActionVariables operation response
type GetActionVariablesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Variables *oaut.SafeArray `idl:"name:variables" json:"variables"`
	// Return: The ActionVariables return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetActionVariablesResponse) xxx_ToOp(ctx context.Context) *xxx_GetActionVariablesOperation {
	if o == nil {
		return &xxx_GetActionVariablesOperation{}
	}
	return &xxx_GetActionVariablesOperation{
		That:      o.That,
		Variables: o.Variables,
		Return:    o.Return,
	}
}

func (o *GetActionVariablesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetActionVariablesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Variables = op.Variables
	o.Return = op.Return
}
func (o *GetActionVariablesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetActionVariablesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionVariablesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetActionVariableDescriptionsOperation structure represents the ActionVariableDescriptions operation
type xxx_GetActionVariableDescriptionsOperation struct {
	This         *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Descriptions *oaut.SafeArray `idl:"name:descriptions" json:"descriptions"`
	Return       int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetActionVariableDescriptionsOperation) OpNum() int { return 8 }

func (o *xxx_GetActionVariableDescriptionsOperation) OpName() string {
	return "/IFsrmFileManagementJobManager/v1/ActionVariableDescriptions"
}

func (o *xxx_GetActionVariableDescriptionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionVariableDescriptionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetActionVariableDescriptionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetActionVariableDescriptionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionVariableDescriptionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// descriptions {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Descriptions != nil {
			_ptr_descriptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Descriptions != nil {
					if err := o.Descriptions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Descriptions, _ptr_descriptions); err != nil {
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

func (o *xxx_GetActionVariableDescriptionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// descriptions {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_descriptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Descriptions == nil {
				o.Descriptions = &oaut.SafeArray{}
			}
			if err := o.Descriptions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_descriptions := func(ptr interface{}) { o.Descriptions = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Descriptions, _s_descriptions, _ptr_descriptions); err != nil {
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

// GetActionVariableDescriptionsRequest structure represents the ActionVariableDescriptions operation request
type GetActionVariableDescriptionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetActionVariableDescriptionsRequest) xxx_ToOp(ctx context.Context) *xxx_GetActionVariableDescriptionsOperation {
	if o == nil {
		return &xxx_GetActionVariableDescriptionsOperation{}
	}
	return &xxx_GetActionVariableDescriptionsOperation{
		This: o.This,
	}
}

func (o *GetActionVariableDescriptionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetActionVariableDescriptionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetActionVariableDescriptionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetActionVariableDescriptionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionVariableDescriptionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetActionVariableDescriptionsResponse structure represents the ActionVariableDescriptions operation response
type GetActionVariableDescriptionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Descriptions *oaut.SafeArray `idl:"name:descriptions" json:"descriptions"`
	// Return: The ActionVariableDescriptions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetActionVariableDescriptionsResponse) xxx_ToOp(ctx context.Context) *xxx_GetActionVariableDescriptionsOperation {
	if o == nil {
		return &xxx_GetActionVariableDescriptionsOperation{}
	}
	return &xxx_GetActionVariableDescriptionsOperation{
		That:         o.That,
		Descriptions: o.Descriptions,
		Return:       o.Return,
	}
}

func (o *GetActionVariableDescriptionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetActionVariableDescriptionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Descriptions = op.Descriptions
	o.Return = op.Return
}
func (o *GetActionVariableDescriptionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetActionVariableDescriptionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionVariableDescriptionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumFileManagementJobsOperation structure represents the EnumFileManagementJobs operation
type xxx_EnumFileManagementJobsOperation struct {
	This               *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Options            fsrm.EnumOptions `idl:"name:options" json:"options"`
	FileManagementJobs *fsrm.Collection `idl:"name:fileManagementJobs" json:"file_management_jobs"`
	Return             int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumFileManagementJobsOperation) OpNum() int { return 9 }

func (o *xxx_EnumFileManagementJobsOperation) OpName() string {
	return "/IFsrmFileManagementJobManager/v1/EnumFileManagementJobs"
}

func (o *xxx_EnumFileManagementJobsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileManagementJobsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileManagementJobsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileManagementJobsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileManagementJobsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileManagementJobs {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.FileManagementJobs != nil {
			_ptr_fileManagementJobs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileManagementJobs != nil {
					if err := o.FileManagementJobs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileManagementJobs, _ptr_fileManagementJobs); err != nil {
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

func (o *xxx_EnumFileManagementJobsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileManagementJobs {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_fileManagementJobs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileManagementJobs == nil {
				o.FileManagementJobs = &fsrm.Collection{}
			}
			if err := o.FileManagementJobs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileManagementJobs := func(ptr interface{}) { o.FileManagementJobs = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.FileManagementJobs, _s_fileManagementJobs, _ptr_fileManagementJobs); err != nil {
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

// EnumFileManagementJobsRequest structure represents the EnumFileManagementJobs operation request
type EnumFileManagementJobsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// options: Contains the FsrmEnumOptions (section 2.2.1.2.5) to use when enumerating
	// the fileManagementJobs.
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *EnumFileManagementJobsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumFileManagementJobsOperation {
	if o == nil {
		return &xxx_EnumFileManagementJobsOperation{}
	}
	return &xxx_EnumFileManagementJobsOperation{
		This:    o.This,
		Options: o.Options,
	}
}

func (o *EnumFileManagementJobsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumFileManagementJobsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Options = op.Options
}
func (o *EnumFileManagementJobsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumFileManagementJobsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumFileManagementJobsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumFileManagementJobsResponse structure represents the EnumFileManagementJobs operation response
type EnumFileManagementJobsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileManagementJobs: Pointer to an IFsrmCollection interface pointer (section 3.2.4.2.1)
	// that upon completion contains pointers to every file management job on the server.
	// A caller MUST release the collection received when the caller is done with it.
	FileManagementJobs *fsrm.Collection `idl:"name:fileManagementJobs" json:"file_management_jobs"`
	// Return: The EnumFileManagementJobs return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumFileManagementJobsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumFileManagementJobsOperation {
	if o == nil {
		return &xxx_EnumFileManagementJobsOperation{}
	}
	return &xxx_EnumFileManagementJobsOperation{
		That:               o.That,
		FileManagementJobs: o.FileManagementJobs,
		Return:             o.Return,
	}
}

func (o *EnumFileManagementJobsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumFileManagementJobsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileManagementJobs = op.FileManagementJobs
	o.Return = op.Return
}
func (o *EnumFileManagementJobsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumFileManagementJobsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumFileManagementJobsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateFileManagementJobOperation structure represents the CreateFileManagementJob operation
type xxx_CreateFileManagementJobOperation struct {
	This              *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat          `idl:"name:That" json:"that"`
	FileManagementJob *fsrm.FileManagementJob `idl:"name:fileManagementJob" json:"file_management_job"`
	Return            int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateFileManagementJobOperation) OpNum() int { return 10 }

func (o *xxx_CreateFileManagementJobOperation) OpName() string {
	return "/IFsrmFileManagementJobManager/v1/CreateFileManagementJob"
}

func (o *xxx_CreateFileManagementJobOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFileManagementJobOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateFileManagementJobOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CreateFileManagementJobOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFileManagementJobOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileManagementJob {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileManagementJob}(interface))
	{
		if o.FileManagementJob != nil {
			_ptr_fileManagementJob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileManagementJob != nil {
					if err := o.FileManagementJob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.FileManagementJob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileManagementJob, _ptr_fileManagementJob); err != nil {
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

func (o *xxx_CreateFileManagementJobOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileManagementJob {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileManagementJob}(interface))
	{
		_ptr_fileManagementJob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileManagementJob == nil {
				o.FileManagementJob = &fsrm.FileManagementJob{}
			}
			if err := o.FileManagementJob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileManagementJob := func(ptr interface{}) { o.FileManagementJob = *ptr.(**fsrm.FileManagementJob) }
		if err := w.ReadPointer(&o.FileManagementJob, _s_fileManagementJob, _ptr_fileManagementJob); err != nil {
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

// CreateFileManagementJobRequest structure represents the CreateFileManagementJob operation request
type CreateFileManagementJobRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CreateFileManagementJobRequest) xxx_ToOp(ctx context.Context) *xxx_CreateFileManagementJobOperation {
	if o == nil {
		return &xxx_CreateFileManagementJobOperation{}
	}
	return &xxx_CreateFileManagementJobOperation{
		This: o.This,
	}
}

func (o *CreateFileManagementJobRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateFileManagementJobOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CreateFileManagementJobRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateFileManagementJobRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFileManagementJobOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateFileManagementJobResponse structure represents the CreateFileManagementJob operation response
type CreateFileManagementJobResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileManagementJob: Pointer to an IFsrmFileManagementJob interface pointer (section
	// 3.2.4.2.48) that upon completion points to a blank fileManagementJob. A caller MUST
	// release the fileManagementJob received when the caller is done with it. To have the
	// fileManagementJob added to the server's List of Persisted File Management Jobs (section
	// 3.2.1.7), the client MUST call Commit (section 3.2.4.2.48.1).
	FileManagementJob *fsrm.FileManagementJob `idl:"name:fileManagementJob" json:"file_management_job"`
	// Return: The CreateFileManagementJob return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateFileManagementJobResponse) xxx_ToOp(ctx context.Context) *xxx_CreateFileManagementJobOperation {
	if o == nil {
		return &xxx_CreateFileManagementJobOperation{}
	}
	return &xxx_CreateFileManagementJobOperation{
		That:              o.That,
		FileManagementJob: o.FileManagementJob,
		Return:            o.Return,
	}
}

func (o *CreateFileManagementJobResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateFileManagementJobOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileManagementJob = op.FileManagementJob
	o.Return = op.Return
}
func (o *CreateFileManagementJobResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateFileManagementJobResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFileManagementJobOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFileManagementJobOperation structure represents the GetFileManagementJob operation
type xxx_GetFileManagementJobOperation struct {
	This              *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Name              *oaut.String            `idl:"name:name" json:"name"`
	FileManagementJob *fsrm.FileManagementJob `idl:"name:fileManagementJob" json:"file_management_job"`
	Return            int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileManagementJobOperation) OpNum() int { return 11 }

func (o *xxx_GetFileManagementJobOperation) OpName() string {
	return "/IFsrmFileManagementJobManager/v1/GetFileManagementJob"
}

func (o *xxx_GetFileManagementJobOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileManagementJobOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// name {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
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

func (o *xxx_GetFileManagementJobOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// name {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_name := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileManagementJobOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileManagementJobOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileManagementJob {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileManagementJob}(interface))
	{
		if o.FileManagementJob != nil {
			_ptr_fileManagementJob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileManagementJob != nil {
					if err := o.FileManagementJob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.FileManagementJob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileManagementJob, _ptr_fileManagementJob); err != nil {
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

func (o *xxx_GetFileManagementJobOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileManagementJob {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileManagementJob}(interface))
	{
		_ptr_fileManagementJob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileManagementJob == nil {
				o.FileManagementJob = &fsrm.FileManagementJob{}
			}
			if err := o.FileManagementJob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileManagementJob := func(ptr interface{}) { o.FileManagementJob = *ptr.(**fsrm.FileManagementJob) }
		if err := w.ReadPointer(&o.FileManagementJob, _s_fileManagementJob, _ptr_fileManagementJob); err != nil {
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

// GetFileManagementJobRequest structure represents the GetFileManagementJob operation request
type GetFileManagementJobRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// name: Contains the name of the fileManagementJob to return.
	Name *oaut.String `idl:"name:name" json:"name"`
}

func (o *GetFileManagementJobRequest) xxx_ToOp(ctx context.Context) *xxx_GetFileManagementJobOperation {
	if o == nil {
		return &xxx_GetFileManagementJobOperation{}
	}
	return &xxx_GetFileManagementJobOperation{
		This: o.This,
		Name: o.Name,
	}
}

func (o *GetFileManagementJobRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileManagementJobOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
}
func (o *GetFileManagementJobRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetFileManagementJobRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileManagementJobOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileManagementJobResponse structure represents the GetFileManagementJob operation response
type GetFileManagementJobResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileManagementJob: Pointer to an IFsrmFileManagementJob interface pointer (section
	// 3.2.4.2.48) that upon completion points to the fileManagementJob with the specified
	// name. A caller MUST release the fileManagementJob received when the caller is done
	// with it.
	FileManagementJob *fsrm.FileManagementJob `idl:"name:fileManagementJob" json:"file_management_job"`
	// Return: The GetFileManagementJob return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFileManagementJobResponse) xxx_ToOp(ctx context.Context) *xxx_GetFileManagementJobOperation {
	if o == nil {
		return &xxx_GetFileManagementJobOperation{}
	}
	return &xxx_GetFileManagementJobOperation{
		That:              o.That,
		FileManagementJob: o.FileManagementJob,
		Return:            o.Return,
	}
}

func (o *GetFileManagementJobResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileManagementJobOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileManagementJob = op.FileManagementJob
	o.Return = op.Return
}
func (o *GetFileManagementJobResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetFileManagementJobResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileManagementJobOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
