package iresourcemanager

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

// IResourceManager server interface.
type ResourceManagerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The RetrieveEventList method is not implemented and always returns an error.
	//
	// Return Values: This method returns E_NOTIMPL, unless a parameter error is found.
	//
	//	+----------------------+-----------------------------------+
	//	|        RETURN        |                                   |
	//	|      VALUE/CODE      |            DESCRIPTION            |
	//	|                      |                                   |
	//	+----------------------+-----------------------------------+
	//	+----------------------+-----------------------------------+
	//	| 0x80004001 E_NOTIMPL | This function is not implemented. |
	//	+----------------------+-----------------------------------+
	//
	// Additional IResourceManager interface methods are specified in section 3.2.4.1.
	RetrieveEventList(context.Context, *RetrieveEventListRequest) (*RetrieveEventListResponse, error)

	// The GetSystemAffinity method gets the processor affinity of the system.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IResourceManager interface methods are specified in section 3.2.4.1.
	GetSystemAffinity(context.Context, *GetSystemAffinityRequest) (*GetSystemAffinityResponse, error)

	// The ImportXMLFiles method loads a specified WSRM configuration.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | Operation successful.                                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | One or more arguments are invalid.                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                      | be processed.<27>                                                                |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// The ImportXMLFiles method can be used to manage system resources by importing a valid
	// RAP.
	//
	// Additional IResourceManager interface methods are specified in section 3.2.4.1.
	ImportXMLFiles(context.Context, *ImportXMLFilesRequest) (*ImportXMLFilesResponse, error)

	// The ExportXMLFiles method saves the current WSRM configuration.<28>
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                      | be processed.<29>                                                                |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | Operation successful.                                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IResourceManager interface methods are specified in section 3.2.4.1.
	ExportXMLFiles(context.Context, *ExportXMLFilesRequest) (*ExportXMLFilesResponse, error)

	// The RestoreXMLFiles method restores the WSRM configuration to a specified state.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | Operation successful.                                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                      | be processed.<31>                                                                |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | One or more arguments are invalid.                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// This method can be used to reset the WSRM server state to a known value in case some
	// data becomes corrupted and the server cannot proceed.
	//
	// Additional IResourceManager interface methods are specified in section 3.2.4.1.
	RestoreXMLFiles(context.Context, *RestoreXMLFilesRequest) (*RestoreXMLFilesResponse, error)

	// The GetDependencies method returns a list of WSRM objects that are being used or
	// that depend on a specified object.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                              | Operation successful.                                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE                  | The specified name contains characters that are invalid. The name cannot         |
	//	|                                              | start with a hyphen ("-") and cannot contain spaces or any of the following      |
	//	|                                              | characters: \ / ? * | : < > " , ;.                                               |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80004005 E_FAIL                            | Either an object with the specified name and type was not found or its           |
	//	|                                              | dependency list could not be created.                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0271 WRM_ERR_CAL_UNKNOWN_SCHEDULE      | A calendar name or an invalid schedule name was passed in bstrObjectName for     |
	//	|                                              | OBJECT_SCHEDULE enumObject.                                                      |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF013A WRM_ERR_DEPENDENCIES_FOR_RESIDUAL | This is a residual PMC and is a part of all policies. All processes that do      |
	//	|                                              | not match any of the PMC specified by a user in a policy automatically match to  |
	//	|                                              | residual PMC.                                                                    |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The server SHOULD process this method call as follows:
	//
	// * If a PMC object is specified, this method MUST return a list of policies that make
	// use of the PMC.
	//
	// * If a resource policy object ( e371cc74-bf4c-4870-8afe-5062cc628b4f#gt_3731cbe9-91f2-49a8-b63f-6f5f7b69b05a
	// ) is specified, this method MUST return a list of calendar ( e371cc74-bf4c-4870-8afe-5062cc628b4f#gt_7204b2ed-dcef-4434-be15-6451f92d03fb
	// ) events, conditions, and schedules that make use of the resource policy ( e371cc74-bf4c-4870-8afe-5062cc628b4f#gt_559b0a4d-161b-4664-9c10-4fab98b97f1f
	// ).
	//
	// * If a schedule object is specified, this method MUST return a list of calendar events
	// making use of the schedule.
	//
	// * If a calendar object is specified, this method MUST return WRM_ERR_CAL_UNKNOWN_SCHEDULE.
	//
	// Additional IResourceManager interface methods are specified in section 3.2.4.1.
	GetDependencies(context.Context, *GetDependenciesRequest) (*GetDependenciesResponse, error)

	// The GetServiceList method provides a list of services that are registered with the
	// server and are not excluded by the exclusion list. This list of services is expected
	// to be used for defining process matching criteria (PMC).
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IResourceManager interface methods are specified in section 3.2.4.1.
	GetServiceList(context.Context, *GetServiceListRequest) (*GetServiceListResponse, error)

	// The GetIISAppPoolNames method returns a list of all the application pools on the
	// WSRM server that are defined by and known to the IIS server.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IResourceManager interface methods are specified in section 3.2.4.1.
	GetIISAppPoolNames(context.Context, *GetIISAppPoolNamesRequest) (*GetIISAppPoolNamesResponse, error)

	// The GetServerName method returns the server name.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.<33>
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IResourceManager interface methods are specified in section 3.2.4.1.
	GetServerName(context.Context, *GetServerNameRequest) (*GetServerNameResponse, error)

	// The GetCurrentMemory method determines the total amount of physical memory in the
	// system.
	//
	// Return Values: This method returns 0x00000001 for success or 0x00000000 if an error
	// occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000        | An error occurred.    |
	//	+-------------------+-----------------------+
	//	| 0x00000001 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IResourceManager interface methods are specified in section 3.2.4.1.
	GetCurrentMemory(context.Context, *GetCurrentMemoryRequest) (*GetCurrentMemoryResponse, error)
}

func RegisterResourceManagerServer(conn dcerpc.Conn, o ResourceManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewResourceManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ResourceManagerSyntaxV0_0))...)
}

func NewResourceManagerServerHandle(o ResourceManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ResourceManagerServerHandle(ctx, o, opNum, r)
	}
}

func ResourceManagerServerHandle(ctx context.Context, o ResourceManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // RetrieveEventList
		op := &xxx_RetrieveEventListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RetrieveEventListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RetrieveEventList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetSystemAffinity
		op := &xxx_GetSystemAffinityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSystemAffinityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSystemAffinity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ImportXMLFiles
		op := &xxx_ImportXMLFilesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ImportXMLFilesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ImportXMLFiles(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // ExportXMLFiles
		op := &xxx_ExportXMLFilesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExportXMLFilesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExportXMLFiles(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RestoreXMLFiles
		op := &xxx_RestoreXMLFilesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RestoreXMLFilesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RestoreXMLFiles(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // GetDependencies
		op := &xxx_GetDependenciesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDependenciesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDependencies(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // GetServiceList
		op := &xxx_GetServiceListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServiceListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServiceList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // GetIISAppPoolNames
		op := &xxx_GetIISAppPoolNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIISAppPoolNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIISAppPoolNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // GetServerName
		op := &xxx_GetServerNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServerNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServerName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // GetCurrentMemory
		op := &xxx_GetCurrentMemoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCurrentMemoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCurrentMemory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IResourceManager
type UnimplementedResourceManagerServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedResourceManagerServer) RetrieveEventList(context.Context, *RetrieveEventListRequest) (*RetrieveEventListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) GetSystemAffinity(context.Context, *GetSystemAffinityRequest) (*GetSystemAffinityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) ImportXMLFiles(context.Context, *ImportXMLFilesRequest) (*ImportXMLFilesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) ExportXMLFiles(context.Context, *ExportXMLFilesRequest) (*ExportXMLFilesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) RestoreXMLFiles(context.Context, *RestoreXMLFilesRequest) (*RestoreXMLFilesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) GetDependencies(context.Context, *GetDependenciesRequest) (*GetDependenciesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) GetServiceList(context.Context, *GetServiceListRequest) (*GetServiceListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) GetIISAppPoolNames(context.Context, *GetIISAppPoolNamesRequest) (*GetIISAppPoolNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) GetServerName(context.Context, *GetServerNameRequest) (*GetServerNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) GetCurrentMemory(context.Context, *GetCurrentMemoryRequest) (*GetCurrentMemoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ResourceManagerServer = (*UnimplementedResourceManagerServer)(nil)
