package ifsrmfilemanagementjobmanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
	_ = idispatch.GoPackage
)

// IFsrmFileManagementJobManager server interface.
type FileManagementJobManagerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// ActionVariables operation.
	GetActionVariables(context.Context, *GetActionVariablesRequest) (*GetActionVariablesResponse, error)

	// ActionVariableDescriptions operation.
	GetActionVariableDescriptions(context.Context, *GetActionVariableDescriptionsRequest) (*GetActionVariableDescriptionsResponse, error)

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
	EnumFileManagementJobs(context.Context, *EnumFileManagementJobsRequest) (*EnumFileManagementJobsResponse, error)

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
	CreateFileManagementJob(context.Context, *CreateFileManagementJobRequest) (*CreateFileManagementJobResponse, error)

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
	GetFileManagementJob(context.Context, *GetFileManagementJobRequest) (*GetFileManagementJobResponse, error)
}

func RegisterFileManagementJobManagerServer(conn dcerpc.Conn, o FileManagementJobManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFileManagementJobManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FileManagementJobManagerSyntaxV1_0))...)
}

func NewFileManagementJobManagerServerHandle(o FileManagementJobManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FileManagementJobManagerServerHandle(ctx, o, opNum, r)
	}
}

func FileManagementJobManagerServerHandle(ctx context.Context, o FileManagementJobManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // ActionVariables
		in := &GetActionVariablesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetActionVariables(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // ActionVariableDescriptions
		in := &GetActionVariableDescriptionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetActionVariableDescriptions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // EnumFileManagementJobs
		in := &EnumFileManagementJobsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumFileManagementJobs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // CreateFileManagementJob
		in := &CreateFileManagementJobRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateFileManagementJob(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // GetFileManagementJob
		in := &GetFileManagementJobRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileManagementJob(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
