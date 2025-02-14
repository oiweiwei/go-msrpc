package iwbemservices

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
	_ = (*errors.Error)(nil)
	_ = iunknown.GoPackage
)

// IWbemServices server interface.
type ServicesServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The IWbemServices::OpenNamespace method provides the client with an IWbemServices
	// interface pointer that is scoped to the requested namespace. The specified namespace
	// MUST be a child namespace of the current namespace through which this method is called.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR, as specified in section
	// 2.2.11, to indicate the successful completion of the method.
	//
	// Requirements described in the parameter definitions are checked, and if the requirements
	// are not met, the server returns WBEM_E_INVALID_PARAMETER.
	OpenNamespace(context.Context, *OpenNamespaceRequest) (*OpenNamespaceResponse, error)

	// The IWbemServices::CancelAsyncCall method cancels a currently pending asynchronous
	// method call identified by the IWbemObjectSink pointer passed to the initial asynchronous
	// method.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	//	+----------------------+-------------------------------------------------------+
	//	|        RETURN        |                                                       |
	//	|      VALUE/CODE      |                      DESCRIPTION                      |
	//	|                      |                                                       |
	//	+----------------------+-------------------------------------------------------+
	//	+----------------------+-------------------------------------------------------+
	//	| 0x00 WBEM_S_NO_ERROR | Indicates a successful completion to the method call. |
	//	+----------------------+-------------------------------------------------------+
	CancelAsyncCall(context.Context, *CancelAsyncCallRequest) (*CancelAsyncCallResponse, error)

	// The QueryObjectSink method obtains a notification handler that allows the client
	// to send events directly to the server.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	QueryObjectSink(context.Context, *QueryObjectSinkRequest) (*QueryObjectSinkResponse, error)

	// The IWbemServices::GetObject method retrieves a CIM class or a CIM instance. This
	// method MUST retrieve CIM objects from the namespace that is associated with the current
	// IWbemServices interface.
	//
	// Return Values: This method MUST return an HRESULT that MUST indicate the status of
	// the method call. The HRESULT MUST have the type and values as specified in section
	// 2.2.11. The server MUST return WBEM_S_NO_ERROR (as specified in section 2.2.11) to
	// indicate the successful completion of the method.
	GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error)

	// The IWbemServices::GetObjectAsync method is the asynchronous version of the IWbemServices::GetObject
	// method.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	GetObjectAsync(context.Context, *GetObjectAsyncRequest) (*GetObjectAsyncResponse, error)

	// The IWbemServices::PutClass method creates a new class or updates an existing class
	// in the namespace that is associated with the current IWbemServices interface. The
	// server MUST NOT allow the creation of classes that have names that begin or end with
	// an underscore because those names are reserved for system classes. If the class name
	// does not conform to the CLASS-NAME element defined in WQL, the server MUST return
	// WBEM_E_INVALID_PARAMETER.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	PutClass(context.Context, *PutClassRequest) (*PutClassResponse, error)

	// The IWbemServices::PutClassAsync method is the asynchronous version of the IWbemServices::PutClass
	// method. The PutClassAsync method creates a new class or updates an existing class.
	// The server MUST NOT allow the creation of classes that have names that begin or end
	// with an underscore because those names are reserved for system classes. If the class
	// name does not conform to the CLASS-NAME element defined in WQL, the server MUST return
	// WBEM_E_INVALID_PARAMETER.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	PutClassAsync(context.Context, *PutClassAsyncRequest) (*PutClassAsyncResponse, error)

	// The IWbemServices::DeleteClass method MUST delete a specified class from the namespace
	// that is associated with the current IWbemServices interface.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	DeleteClass(context.Context, *DeleteClassRequest) (*DeleteClassResponse, error)

	// The IWbemServices::DeleteClassAsync method is the asynchronous version of the IWbemServices::DeleteClass
	// method. The DeleteClassAsync method MUST delete a specified class from the namespace.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	DeleteClassAsync(context.Context, *DeleteClassAsyncRequest) (*DeleteClassAsyncResponse, error)

	// The IWbemServices::CreateClassEnum method provides a class enumeration. When this
	// method is invoked, the server MUST return all classes that satisfy the selection
	// criteria from the namespace that is associated with the current IWbemServices interface.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	CreateClassEnum(context.Context, *CreateClassEnumRequest) (*CreateClassEnumResponse, error)

	// The IWbemServices::CreateClassEnumAsync method provides an asynchronous class enumeration.
	// When this method is invoked, the server MUST return all classes that satisfy the
	// selection criteria.
	//
	// Return Values: This method MUST return an HRESULT, which MUST indicate the status
	// of the method call. The HRESULT MUST have the type and values as specified in section
	// 2.2.11. The server MUST return WBEM_S_NO_ERROR (specified in section 2.2.11) to indicate
	// the successful completion of the method.
	CreateClassEnumAsync(context.Context, *CreateClassEnumAsyncRequest) (*CreateClassEnumAsyncResponse, error)

	// The IWbemServices::PutInstance method creates or updates an instance of an existing
	// class.
	//
	// The PutInstance method opnum equals 14.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	PutInstance(context.Context, *PutInstanceRequest) (*PutInstanceResponse, error)

	// The IWbemServices::PutInstanceAsync method is the asynchronous version of the PutInstance
	// method. The PutInstanceAsync method creates or updates an instance of an existing
	// class.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	PutInstanceAsync(context.Context, *PutInstanceAsyncRequest) (*PutInstanceAsyncResponse, error)

	// The IWbemServices::DeleteInstance method deletes an instance of an existing class
	// from the namespace that is pointed to by the IWbemServices interface object that
	// is used to call the method.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	DeleteInstance(context.Context, *DeleteInstanceRequest) (*DeleteInstanceResponse, error)

	// The IWbemServices::DeleteInstanceAsync method is the asynchronous version of the
	// IWbemServices::DeleteInstance method. The IWbemServices::DeleteInstanceAsync method
	// deletes an instance of an existing class from the namespace that is pointed to by
	// the IWbemServices interface that is used to call the method.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	DeleteInstanceAsync(context.Context, *DeleteInstanceAsyncRequest) (*DeleteInstanceAsyncResponse, error)

	// The IWbemServices::CreateInstanceEnum method provides an instance enumeration. When
	// this method is invoked, the server MUST return all instances for the specific class
	// in the current namespace.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return the following value (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	CreateInstanceEnum(context.Context, *CreateInstanceEnumRequest) (*CreateInstanceEnumResponse, error)

	// The IWbemServices::CreateInstanceEnumAsync method provides an asynchronous instance
	// enumeration. When this method is invoked, the server MUST return all instances for
	// the specific class in the current namespace.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	CreateInstanceEnumAsync(context.Context, *CreateInstanceEnumAsyncRequest) (*CreateInstanceEnumAsyncResponse, error)

	// The IWbemServices::ExecQuery method returns an enumerable collection of IWbemClassObject
	// interface objects based on a query.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	ExecQuery(context.Context, *ExecQueryRequest) (*ExecQueryResponse, error)

	// The IWbemServices::ExecQueryAsync method is the asynchronous version of the IWbemServices::ExecQuery
	// method. The IWbemServices::ExecQueryAsync method returns an enumerable collection
	// of IWbemClassObject interface objects based on a query.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	ExecQueryAsync(context.Context, *ExecQueryAsyncRequest) (*ExecQueryAsyncResponse, error)

	// The IWbemServices::ExecNotificationQuery method provides a subscription for event
	// notifications. When this method is invoked, the server runs a query to deliver events
	// matching the query. The call is executed semisynchronously and MUST follow the rules
	// that are specified in section 3.1.1.1.2. The WMI client can poll the returned enumerator
	// for events as they arrive. Releasing the returned enumerator cancels the query.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	ExecNotificationQuery(context.Context, *ExecNotificationQueryRequest) (*ExecNotificationQueryResponse, error)

	// The IWbemServices::ExecNotificationQueryAsync method is the asynchronous version
	// of the IWbemServices::ExecNotificationQuery method. The IWbemServices::ExecNotificationQueryAsync
	// method provides subscription for asynchronous event notifications. When this method
	// is invoked, the server performs the same task as the IWbemServices::ExecNotificationQuery
	// method, except that events are supplied to the specified response handler (pResponseHandler)
	// until the IWbemServices::CancelAsyncCall method is called.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR, as specified in section
	// 2.2.11, to indicate the successful completion of the method.
	ExecNotificationQueryAsync(context.Context, *ExecNotificationQueryAsyncRequest) (*ExecNotificationQueryAsyncResponse, error)

	// The IWbemServices::ExecMethod method executes a CIM method that is implemented by
	// a CIM class or a CIM instance that is retrieved from the IWbemServices interface.
	//
	// Return Values: This method MUST return an HRESULT, which MUST indicate the status
	// of the method call. HRESULT MUST have the type and values as specified in section
	// 2.2.11. The server MUST return WBEM_S_NO_ERROR (specified in section 2.2.11) to indicate
	// the successful completion of the method.
	ExecMethod(context.Context, *ExecMethodRequest) (*ExecMethodResponse, error)

	// The IWbemServices::ExecMethodAsync method asynchronously executes a CIM method that
	// is implemented by a CIM class or a CIM instance that is retrieved from the IWbemServices
	// interface.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	ExecMethodAsync(context.Context, *ExecMethodAsyncRequest) (*ExecMethodAsyncResponse, error)
}

func RegisterServicesServer(conn dcerpc.Conn, o ServicesServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewServicesServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ServicesSyntaxV0_0))...)
}

func NewServicesServerHandle(o ServicesServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ServicesServerHandle(ctx, o, opNum, r)
	}
}

func ServicesServerHandle(ctx context.Context, o ServicesServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // OpenNamespace
		op := &xxx_OpenNamespaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenNamespaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenNamespace(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // CancelAsyncCall
		op := &xxx_CancelAsyncCallOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CancelAsyncCallRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CancelAsyncCall(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // QueryObjectSink
		op := &xxx_QueryObjectSinkOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryObjectSinkRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryObjectSink(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // GetObject
		op := &xxx_GetObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // GetObjectAsync
		op := &xxx_GetObjectAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetObjectAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetObjectAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // PutClass
		op := &xxx_PutClassOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PutClassRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PutClass(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // PutClassAsync
		op := &xxx_PutClassAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PutClassAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PutClassAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // DeleteClass
		op := &xxx_DeleteClassOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteClassRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteClass(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // DeleteClassAsync
		op := &xxx_DeleteClassAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteClassAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteClassAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // CreateClassEnum
		op := &xxx_CreateClassEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateClassEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateClassEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // CreateClassEnumAsync
		op := &xxx_CreateClassEnumAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateClassEnumAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateClassEnumAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // PutInstance
		op := &xxx_PutInstanceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PutInstanceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PutInstance(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // PutInstanceAsync
		op := &xxx_PutInstanceAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PutInstanceAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PutInstanceAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // DeleteInstance
		op := &xxx_DeleteInstanceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteInstanceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteInstance(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // DeleteInstanceAsync
		op := &xxx_DeleteInstanceAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteInstanceAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteInstanceAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // CreateInstanceEnum
		op := &xxx_CreateInstanceEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateInstanceEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateInstanceEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // CreateInstanceEnumAsync
		op := &xxx_CreateInstanceEnumAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateInstanceEnumAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateInstanceEnumAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // ExecQuery
		op := &xxx_ExecQueryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecQueryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExecQuery(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // ExecQueryAsync
		op := &xxx_ExecQueryAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecQueryAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExecQueryAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // ExecNotificationQuery
		op := &xxx_ExecNotificationQueryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecNotificationQueryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExecNotificationQuery(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // ExecNotificationQueryAsync
		op := &xxx_ExecNotificationQueryAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecNotificationQueryAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExecNotificationQueryAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // ExecMethod
		op := &xxx_ExecMethodOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecMethodRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExecMethod(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // ExecMethodAsync
		op := &xxx_ExecMethodAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecMethodAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExecMethodAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWbemServices
type UnimplementedServicesServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedServicesServer) OpenNamespace(context.Context, *OpenNamespaceRequest) (*OpenNamespaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) CancelAsyncCall(context.Context, *CancelAsyncCallRequest) (*CancelAsyncCallResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) QueryObjectSink(context.Context, *QueryObjectSinkRequest) (*QueryObjectSinkResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) GetObjectAsync(context.Context, *GetObjectAsyncRequest) (*GetObjectAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) PutClass(context.Context, *PutClassRequest) (*PutClassResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) PutClassAsync(context.Context, *PutClassAsyncRequest) (*PutClassAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) DeleteClass(context.Context, *DeleteClassRequest) (*DeleteClassResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) DeleteClassAsync(context.Context, *DeleteClassAsyncRequest) (*DeleteClassAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) CreateClassEnum(context.Context, *CreateClassEnumRequest) (*CreateClassEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) CreateClassEnumAsync(context.Context, *CreateClassEnumAsyncRequest) (*CreateClassEnumAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) PutInstance(context.Context, *PutInstanceRequest) (*PutInstanceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) PutInstanceAsync(context.Context, *PutInstanceAsyncRequest) (*PutInstanceAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) DeleteInstance(context.Context, *DeleteInstanceRequest) (*DeleteInstanceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) DeleteInstanceAsync(context.Context, *DeleteInstanceAsyncRequest) (*DeleteInstanceAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) CreateInstanceEnum(context.Context, *CreateInstanceEnumRequest) (*CreateInstanceEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) CreateInstanceEnumAsync(context.Context, *CreateInstanceEnumAsyncRequest) (*CreateInstanceEnumAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) ExecQuery(context.Context, *ExecQueryRequest) (*ExecQueryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) ExecQueryAsync(context.Context, *ExecQueryAsyncRequest) (*ExecQueryAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) ExecNotificationQuery(context.Context, *ExecNotificationQueryRequest) (*ExecNotificationQueryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) ExecNotificationQueryAsync(context.Context, *ExecNotificationQueryAsyncRequest) (*ExecNotificationQueryAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) ExecMethod(context.Context, *ExecMethodRequest) (*ExecMethodResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServicesServer) ExecMethodAsync(context.Context, *ExecMethodAsyncRequest) (*ExecMethodAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ServicesServer = (*UnimplementedServicesServer)(nil)
