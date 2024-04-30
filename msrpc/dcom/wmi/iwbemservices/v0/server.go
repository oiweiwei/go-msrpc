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
		in := &OpenNamespaceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenNamespace(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // CancelAsyncCall
		in := &CancelAsyncCallRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CancelAsyncCall(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // QueryObjectSink
		in := &QueryObjectSinkRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryObjectSink(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // GetObject
		in := &GetObjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetObject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // GetObjectAsync
		in := &GetObjectAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetObjectAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // PutClass
		in := &PutClassRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PutClass(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // PutClassAsync
		in := &PutClassAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PutClassAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // DeleteClass
		in := &DeleteClassRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteClass(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // DeleteClassAsync
		in := &DeleteClassAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteClassAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // CreateClassEnum
		in := &CreateClassEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateClassEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // CreateClassEnumAsync
		in := &CreateClassEnumAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateClassEnumAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // PutInstance
		in := &PutInstanceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PutInstance(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // PutInstanceAsync
		in := &PutInstanceAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PutInstanceAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // DeleteInstance
		in := &DeleteInstanceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteInstance(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // DeleteInstanceAsync
		in := &DeleteInstanceAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteInstanceAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // CreateInstanceEnum
		in := &CreateInstanceEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateInstanceEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // CreateInstanceEnumAsync
		in := &CreateInstanceEnumAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateInstanceEnumAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // ExecQuery
		in := &ExecQueryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExecQuery(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // ExecQueryAsync
		in := &ExecQueryAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExecQueryAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // ExecNotificationQuery
		in := &ExecNotificationQueryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExecNotificationQuery(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // ExecNotificationQueryAsync
		in := &ExecNotificationQueryAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExecNotificationQueryAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // ExecMethod
		in := &ExecMethodRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExecMethod(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // ExecMethodAsync
		in := &ExecMethodAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExecMethodAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
