package iresourcemanager2

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

// IResourceManager2 server interface.
type ResourceManager2Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The ExportObjects method creates XML for exporting objects.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | Operation successful.              |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// Additional IResourceManager2 interface methods are specified in section 3.2.4.2.
	//
	// The server SHOULD process this method call as follows.
	//
	// * If one or more ObjectIds XML elements included in the bstrObjectIds parameter cannot
	// be found, they MUST be ignored, and the call MUST proceed for other ObjectIds.
	//
	// * If a PMC ( e371cc74-bf4c-4870-8afe-5062cc628b4f#gt_65daee12-445f-41c0-8456-f728063bef24
	// ) object is specified in the enumObjectType parameter, this method MUST return an
	// XML string containing the requested PMCs in the configuration. The "Type" attributes
	// of individual objects in the ObjectIds XML element MUST be ignored.
	//
	// * If a resource policy object is specified in the enumObjectType parameter, this
	// method MUST return an XML string containing the requested RAPs ( e371cc74-bf4c-4870-8afe-5062cc628b4f#gt_6442eac9-3264-4d95-a435-fdc08e71603f
	// ) in the configuration. The "Type" attributes of individual objects in the ObjectIds
	// XML element MUST be ignored.
	//
	// * If a schedule object is specified in the enumObjectType parameter and the "Type"
	// attribute of the object requested in the ObjectIds XML element is "Calendar", this
	// method MUST return an XML string containing the requested calendar objects in the
	// configuration.
	//
	// * If a schedule object is specified in the enumObjectType parameter and the "Type"
	// attribute of the object requested in the ObjectIds XML element is not "Calendar",
	// this method MUST return an XML string containing the requested schedule objects in
	// the configuration.
	ExportObjects(context.Context, *ExportObjectsRequest) (*ExportObjectsResponse, error)

	// The GetImportConflicts method finds conflicts between import objects and existing
	// objects.
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
	//	|                                      | be processed.<36>                                                                |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IResourceManager2 interface methods are specified in section 3.2.4.2.
	GetImportConflicts(context.Context, *GetImportConflictsRequest) (*GetImportConflictsResponse, error)

	// The ImportXml method imports objects into the configuration.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                 |                                                                                  |
	//	|               VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                        | Operation successful.                                                            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                | One or more arguments are invalid.                                               |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0272 WRM_ERR_CAL_SCHEDULE_IN_USE | An existing schedule with the same name as the one in the supplied               |
	//	|                                        | CalendarsCollection is currently in use with an existing calendar. The complete  |
	//	|                                        | import process is canceled.<38>                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER   | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                        | be processed.<39>                                                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IResourceManager2 interface methods are specified in section 3.2.4.2.
	ImportXML(context.Context, *ImportXMLRequest) (*ImportXMLResponse, error)

	// The ExportXml method exports objects from the configuration.
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
	// Additional IResourceManager2 interface methods are specified in section 3.2.4.2.
	ExportXML(context.Context, *ExportXMLRequest) (*ExportXMLResponse, error)
}

func RegisterResourceManager2Server(conn dcerpc.Conn, o ResourceManager2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewResourceManager2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ResourceManager2SyntaxV0_0))...)
}

func NewResourceManager2ServerHandle(o ResourceManager2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ResourceManager2ServerHandle(ctx, o, opNum, r)
	}
}

func ResourceManager2ServerHandle(ctx context.Context, o ResourceManager2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // ExportObjects
		op := &xxx_ExportObjectsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExportObjectsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExportObjects(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetImportConflicts
		op := &xxx_GetImportConflictsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetImportConflictsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetImportConflicts(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ImportXml
		op := &xxx_ImportXMLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ImportXMLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ImportXML(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // ExportXml
		op := &xxx_ExportXMLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExportXMLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExportXML(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IResourceManager2
type UnimplementedResourceManager2Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedResourceManager2Server) ExportObjects(context.Context, *ExportObjectsRequest) (*ExportObjectsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManager2Server) GetImportConflicts(context.Context, *GetImportConflictsRequest) (*GetImportConflictsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManager2Server) ImportXML(context.Context, *ImportXMLRequest) (*ImportXMLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManager2Server) ExportXML(context.Context, *ExportXMLRequest) (*ExportXMLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ResourceManager2Server = (*UnimplementedResourceManager2Server)(nil)
