package iwbemservices

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	wmi "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi"
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
	_ = iunknown.GoPackage
	_ = oaut.GoPackage
	_ = wmi.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IWbemServices interface identifier 9556dc99-828c-11cf-a37e-00aa003240c7
	ServicesIID = &dcom.IID{Data1: 0x9556dc99, Data2: 0x828c, Data3: 0x11cf, Data4: []byte{0xa3, 0x7e, 0x00, 0xaa, 0x00, 0x32, 0x40, 0xc7}}
	// Syntax UUID
	ServicesSyntaxUUID = &uuid.UUID{TimeLow: 0x9556dc99, TimeMid: 0x828c, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x7e, Node: [6]uint8{0x0, 0xaa, 0x0, 0x32, 0x40, 0xc7}}
	// Syntax ID
	ServicesSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ServicesSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemServices interface.
type ServicesClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

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
	OpenNamespace(context.Context, *OpenNamespaceRequest, ...dcerpc.CallOption) (*OpenNamespaceResponse, error)

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
	CancelAsyncCall(context.Context, *CancelAsyncCallRequest, ...dcerpc.CallOption) (*CancelAsyncCallResponse, error)

	// The QueryObjectSink method obtains a notification handler that allows the client
	// to send events directly to the server.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	QueryObjectSink(context.Context, *QueryObjectSinkRequest, ...dcerpc.CallOption) (*QueryObjectSinkResponse, error)

	// The IWbemServices::GetObject method retrieves a CIM class or a CIM instance. This
	// method MUST retrieve CIM objects from the namespace that is associated with the current
	// IWbemServices interface.
	//
	// Return Values: This method MUST return an HRESULT that MUST indicate the status of
	// the method call. The HRESULT MUST have the type and values as specified in section
	// 2.2.11. The server MUST return WBEM_S_NO_ERROR (as specified in section 2.2.11) to
	// indicate the successful completion of the method.
	GetObject(context.Context, *GetObjectRequest, ...dcerpc.CallOption) (*GetObjectResponse, error)

	// The IWbemServices::GetObjectAsync method is the asynchronous version of the IWbemServices::GetObject
	// method.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	GetObjectAsync(context.Context, *GetObjectAsyncRequest, ...dcerpc.CallOption) (*GetObjectAsyncResponse, error)

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
	PutClass(context.Context, *PutClassRequest, ...dcerpc.CallOption) (*PutClassResponse, error)

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
	PutClassAsync(context.Context, *PutClassAsyncRequest, ...dcerpc.CallOption) (*PutClassAsyncResponse, error)

	// The IWbemServices::DeleteClass method MUST delete a specified class from the namespace
	// that is associated with the current IWbemServices interface.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	DeleteClass(context.Context, *DeleteClassRequest, ...dcerpc.CallOption) (*DeleteClassResponse, error)

	// The IWbemServices::DeleteClassAsync method is the asynchronous version of the IWbemServices::DeleteClass
	// method. The DeleteClassAsync method MUST delete a specified class from the namespace.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	DeleteClassAsync(context.Context, *DeleteClassAsyncRequest, ...dcerpc.CallOption) (*DeleteClassAsyncResponse, error)

	// The IWbemServices::CreateClassEnum method provides a class enumeration. When this
	// method is invoked, the server MUST return all classes that satisfy the selection
	// criteria from the namespace that is associated with the current IWbemServices interface.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	CreateClassEnum(context.Context, *CreateClassEnumRequest, ...dcerpc.CallOption) (*CreateClassEnumResponse, error)

	// The IWbemServices::CreateClassEnumAsync method provides an asynchronous class enumeration.
	// When this method is invoked, the server MUST return all classes that satisfy the
	// selection criteria.
	//
	// Return Values: This method MUST return an HRESULT, which MUST indicate the status
	// of the method call. The HRESULT MUST have the type and values as specified in section
	// 2.2.11. The server MUST return WBEM_S_NO_ERROR (specified in section 2.2.11) to indicate
	// the successful completion of the method.
	CreateClassEnumAsync(context.Context, *CreateClassEnumAsyncRequest, ...dcerpc.CallOption) (*CreateClassEnumAsyncResponse, error)

	// The IWbemServices::PutInstance method creates or updates an instance of an existing
	// class.
	//
	// The PutInstance method opnum equals 14.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	PutInstance(context.Context, *PutInstanceRequest, ...dcerpc.CallOption) (*PutInstanceResponse, error)

	// The IWbemServices::PutInstanceAsync method is the asynchronous version of the PutInstance
	// method. The PutInstanceAsync method creates or updates an instance of an existing
	// class.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	PutInstanceAsync(context.Context, *PutInstanceAsyncRequest, ...dcerpc.CallOption) (*PutInstanceAsyncResponse, error)

	// The IWbemServices::DeleteInstance method deletes an instance of an existing class
	// from the namespace that is pointed to by the IWbemServices interface object that
	// is used to call the method.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	DeleteInstance(context.Context, *DeleteInstanceRequest, ...dcerpc.CallOption) (*DeleteInstanceResponse, error)

	// The IWbemServices::DeleteInstanceAsync method is the asynchronous version of the
	// IWbemServices::DeleteInstance method. The IWbemServices::DeleteInstanceAsync method
	// deletes an instance of an existing class from the namespace that is pointed to by
	// the IWbemServices interface that is used to call the method.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	DeleteInstanceAsync(context.Context, *DeleteInstanceAsyncRequest, ...dcerpc.CallOption) (*DeleteInstanceAsyncResponse, error)

	// The IWbemServices::CreateInstanceEnum method provides an instance enumeration. When
	// this method is invoked, the server MUST return all instances for the specific class
	// in the current namespace.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return the following value (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	CreateInstanceEnum(context.Context, *CreateInstanceEnumRequest, ...dcerpc.CallOption) (*CreateInstanceEnumResponse, error)

	// The IWbemServices::CreateInstanceEnumAsync method provides an asynchronous instance
	// enumeration. When this method is invoked, the server MUST return all instances for
	// the specific class in the current namespace.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	CreateInstanceEnumAsync(context.Context, *CreateInstanceEnumAsyncRequest, ...dcerpc.CallOption) (*CreateInstanceEnumAsyncResponse, error)

	// The IWbemServices::ExecQuery method returns an enumerable collection of IWbemClassObject
	// interface objects based on a query.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	ExecQuery(context.Context, *ExecQueryRequest, ...dcerpc.CallOption) (*ExecQueryResponse, error)

	// The IWbemServices::ExecQueryAsync method is the asynchronous version of the IWbemServices::ExecQuery
	// method. The IWbemServices::ExecQueryAsync method returns an enumerable collection
	// of IWbemClassObject interface objects based on a query.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	ExecQueryAsync(context.Context, *ExecQueryAsyncRequest, ...dcerpc.CallOption) (*ExecQueryAsyncResponse, error)

	// The IWbemServices::ExecNotificationQuery method provides a subscription for event
	// notifications. When this method is invoked, the server runs a query to deliver events
	// matching the query. The call is executed semisynchronously and MUST follow the rules
	// that are specified in section 3.1.1.1.2. The WMI client can poll the returned enumerator
	// for events as they arrive. Releasing the returned enumerator cancels the query.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	ExecNotificationQuery(context.Context, *ExecNotificationQueryRequest, ...dcerpc.CallOption) (*ExecNotificationQueryResponse, error)

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
	ExecNotificationQueryAsync(context.Context, *ExecNotificationQueryAsyncRequest, ...dcerpc.CallOption) (*ExecNotificationQueryAsyncResponse, error)

	// The IWbemServices::ExecMethod method executes a CIM method that is implemented by
	// a CIM class or a CIM instance that is retrieved from the IWbemServices interface.
	//
	// Return Values: This method MUST return an HRESULT, which MUST indicate the status
	// of the method call. HRESULT MUST have the type and values as specified in section
	// 2.2.11. The server MUST return WBEM_S_NO_ERROR (specified in section 2.2.11) to indicate
	// the successful completion of the method.
	ExecMethod(context.Context, *ExecMethodRequest, ...dcerpc.CallOption) (*ExecMethodResponse, error)

	// The IWbemServices::ExecMethodAsync method asynchronously executes a CIM method that
	// is implemented by a CIM class or a CIM instance that is retrieved from the IWbemServices
	// interface.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (as specified in section
	// 2.2.11) to indicate the successful completion of the method.
	ExecMethodAsync(context.Context, *ExecMethodAsyncRequest, ...dcerpc.CallOption) (*ExecMethodAsyncResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ServicesClient
}

type xxx_DefaultServicesClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultServicesClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultServicesClient) OpenNamespace(ctx context.Context, in *OpenNamespaceRequest, opts ...dcerpc.CallOption) (*OpenNamespaceResponse, error) {
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
	out := &OpenNamespaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) CancelAsyncCall(ctx context.Context, in *CancelAsyncCallRequest, opts ...dcerpc.CallOption) (*CancelAsyncCallResponse, error) {
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
	out := &CancelAsyncCallResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) QueryObjectSink(ctx context.Context, in *QueryObjectSinkRequest, opts ...dcerpc.CallOption) (*QueryObjectSinkResponse, error) {
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
	out := &QueryObjectSinkResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...dcerpc.CallOption) (*GetObjectResponse, error) {
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
	out := &GetObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) GetObjectAsync(ctx context.Context, in *GetObjectAsyncRequest, opts ...dcerpc.CallOption) (*GetObjectAsyncResponse, error) {
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
	out := &GetObjectAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) PutClass(ctx context.Context, in *PutClassRequest, opts ...dcerpc.CallOption) (*PutClassResponse, error) {
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
	out := &PutClassResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) PutClassAsync(ctx context.Context, in *PutClassAsyncRequest, opts ...dcerpc.CallOption) (*PutClassAsyncResponse, error) {
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
	out := &PutClassAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) DeleteClass(ctx context.Context, in *DeleteClassRequest, opts ...dcerpc.CallOption) (*DeleteClassResponse, error) {
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
	out := &DeleteClassResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) DeleteClassAsync(ctx context.Context, in *DeleteClassAsyncRequest, opts ...dcerpc.CallOption) (*DeleteClassAsyncResponse, error) {
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
	out := &DeleteClassAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) CreateClassEnum(ctx context.Context, in *CreateClassEnumRequest, opts ...dcerpc.CallOption) (*CreateClassEnumResponse, error) {
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
	out := &CreateClassEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) CreateClassEnumAsync(ctx context.Context, in *CreateClassEnumAsyncRequest, opts ...dcerpc.CallOption) (*CreateClassEnumAsyncResponse, error) {
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
	out := &CreateClassEnumAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) PutInstance(ctx context.Context, in *PutInstanceRequest, opts ...dcerpc.CallOption) (*PutInstanceResponse, error) {
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
	out := &PutInstanceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) PutInstanceAsync(ctx context.Context, in *PutInstanceAsyncRequest, opts ...dcerpc.CallOption) (*PutInstanceAsyncResponse, error) {
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
	out := &PutInstanceAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) DeleteInstance(ctx context.Context, in *DeleteInstanceRequest, opts ...dcerpc.CallOption) (*DeleteInstanceResponse, error) {
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
	out := &DeleteInstanceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) DeleteInstanceAsync(ctx context.Context, in *DeleteInstanceAsyncRequest, opts ...dcerpc.CallOption) (*DeleteInstanceAsyncResponse, error) {
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
	out := &DeleteInstanceAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) CreateInstanceEnum(ctx context.Context, in *CreateInstanceEnumRequest, opts ...dcerpc.CallOption) (*CreateInstanceEnumResponse, error) {
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
	out := &CreateInstanceEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) CreateInstanceEnumAsync(ctx context.Context, in *CreateInstanceEnumAsyncRequest, opts ...dcerpc.CallOption) (*CreateInstanceEnumAsyncResponse, error) {
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
	out := &CreateInstanceEnumAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) ExecQuery(ctx context.Context, in *ExecQueryRequest, opts ...dcerpc.CallOption) (*ExecQueryResponse, error) {
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
	out := &ExecQueryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) ExecQueryAsync(ctx context.Context, in *ExecQueryAsyncRequest, opts ...dcerpc.CallOption) (*ExecQueryAsyncResponse, error) {
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
	out := &ExecQueryAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) ExecNotificationQuery(ctx context.Context, in *ExecNotificationQueryRequest, opts ...dcerpc.CallOption) (*ExecNotificationQueryResponse, error) {
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
	out := &ExecNotificationQueryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) ExecNotificationQueryAsync(ctx context.Context, in *ExecNotificationQueryAsyncRequest, opts ...dcerpc.CallOption) (*ExecNotificationQueryAsyncResponse, error) {
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
	out := &ExecNotificationQueryAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) ExecMethod(ctx context.Context, in *ExecMethodRequest, opts ...dcerpc.CallOption) (*ExecMethodResponse, error) {
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
	out := &ExecMethodResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) ExecMethodAsync(ctx context.Context, in *ExecMethodAsyncRequest, opts ...dcerpc.CallOption) (*ExecMethodAsyncResponse, error) {
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
	out := &ExecMethodAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServicesClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultServicesClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultServicesClient) IPID(ctx context.Context, ipid *dcom.IPID) ServicesClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultServicesClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewServicesClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ServicesClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ServicesSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultServicesClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_OpenNamespaceOperation structure represents the OpenNamespace operation
type xxx_OpenNamespaceOperation struct {
	This             *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Namespace        *oaut.String    `idl:"name:strNamespace" json:"namespace"`
	Flags            int32           `idl:"name:lFlags" json:"flags"`
	Context          *wmi.Context    `idl:"name:pCtx" json:"context"`
	WorkingNamespace *wmi.Services   `idl:"name:ppWorkingNamespace;pointer:unique" json:"working_namespace"`
	Result           *wmi.CallResult `idl:"name:ppResult;pointer:unique" json:"result"`
	Return           int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenNamespaceOperation) OpNum() int { return 3 }

func (o *xxx_OpenNamespaceOperation) OpName() string { return "/IWbemServices/v0/OpenNamespace" }

func (o *xxx_OpenNamespaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNamespaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strNamespace {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Namespace != nil {
			_ptr_strNamespace := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Namespace != nil {
					if err := o.Namespace.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Namespace, _ptr_strNamespace); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// ppWorkingNamespace {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemServices}(interface))
	{
		if o.WorkingNamespace != nil {
			_ptr_ppWorkingNamespace := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WorkingNamespace != nil {
					_ptr_ppWorkingNamespace := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.WorkingNamespace != nil {
							if err := o.WorkingNamespace.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.Services{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.WorkingNamespace, _ptr_ppWorkingNamespace); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WorkingNamespace, _ptr_ppWorkingNamespace); err != nil {
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
	// ppResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.Result != nil {
			_ptr_ppResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Result != nil {
					_ptr_ppResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Result != nil {
							if err := o.Result.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Result, _ptr_ppResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Result, _ptr_ppResult); err != nil {
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

func (o *xxx_OpenNamespaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strNamespace {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strNamespace := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Namespace == nil {
				o.Namespace = &oaut.String{}
			}
			if err := o.Namespace.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strNamespace := func(ptr interface{}) { o.Namespace = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Namespace, _s_strNamespace, _ptr_strNamespace); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppWorkingNamespace {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemServices}(interface))
	{
		_ptr_ppWorkingNamespace := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppWorkingNamespace := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.WorkingNamespace == nil {
					o.WorkingNamespace = &wmi.Services{}
				}
				if err := o.WorkingNamespace.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppWorkingNamespace := func(ptr interface{}) { o.WorkingNamespace = *ptr.(**wmi.Services) }
			if err := w.ReadPointer(&o.WorkingNamespace, _s_ppWorkingNamespace, _ptr_ppWorkingNamespace); err != nil {
				return err
			}
			return nil
		})
		_s_ppWorkingNamespace := func(ptr interface{}) { o.WorkingNamespace = *ptr.(**wmi.Services) }
		if err := w.ReadPointer(&o.WorkingNamespace, _s_ppWorkingNamespace, _ptr_ppWorkingNamespace); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Result == nil {
					o.Result = &wmi.CallResult{}
				}
				if err := o.Result.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppResult := func(ptr interface{}) { o.Result = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.Result, _s_ppResult, _ptr_ppResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppResult := func(ptr interface{}) { o.Result = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.Result, _s_ppResult, _ptr_ppResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNamespaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNamespaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppWorkingNamespace {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemServices}(interface))
	{
		if o.WorkingNamespace != nil {
			_ptr_ppWorkingNamespace := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WorkingNamespace != nil {
					_ptr_ppWorkingNamespace := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.WorkingNamespace != nil {
							if err := o.WorkingNamespace.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.Services{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.WorkingNamespace, _ptr_ppWorkingNamespace); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WorkingNamespace, _ptr_ppWorkingNamespace); err != nil {
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
	// ppResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.Result != nil {
			_ptr_ppResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Result != nil {
					_ptr_ppResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Result != nil {
							if err := o.Result.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Result, _ptr_ppResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Result, _ptr_ppResult); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNamespaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppWorkingNamespace {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemServices}(interface))
	{
		_ptr_ppWorkingNamespace := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppWorkingNamespace := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.WorkingNamespace == nil {
					o.WorkingNamespace = &wmi.Services{}
				}
				if err := o.WorkingNamespace.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppWorkingNamespace := func(ptr interface{}) { o.WorkingNamespace = *ptr.(**wmi.Services) }
			if err := w.ReadPointer(&o.WorkingNamespace, _s_ppWorkingNamespace, _ptr_ppWorkingNamespace); err != nil {
				return err
			}
			return nil
		})
		_s_ppWorkingNamespace := func(ptr interface{}) { o.WorkingNamespace = *ptr.(**wmi.Services) }
		if err := w.ReadPointer(&o.WorkingNamespace, _s_ppWorkingNamespace, _ptr_ppWorkingNamespace); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Result == nil {
					o.Result = &wmi.CallResult{}
				}
				if err := o.Result.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppResult := func(ptr interface{}) { o.Result = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.Result, _s_ppResult, _ptr_ppResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppResult := func(ptr interface{}) { o.Result = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.Result, _s_ppResult, _ptr_ppResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenNamespaceRequest structure represents the OpenNamespace operation request
type OpenNamespaceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strNamespace: MUST be the CIM path to the target namespace. This parameter MUST NOT
	// be NULL.
	Namespace *oaut.String `idl:"name:strNamespace" json:"namespace"`
	// lFlags: Flags that affect the behavior of the OpenNamespace method. The behavior
	// of each flag MUST be interpreted as follows:
	//
	// * If this bit is not set, the server MUST make the method call synchronous.
	//
	// * If this bit is set, the server MUST make the method call semisynchronously.
	//
	// # Name
	//
	// # Value
	//
	// WBEM_FLAG_RETURN_IMMEDIATELY
	//
	// 0x00000010
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: This parameter MUST be NULL.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// ppWorkingNamespace: This parameter MUST NOT be NULL on input when WBEM_FLAG_RETURN_IMMEDIATELY
	// is not set. If the method returns WBEM_S_NO_ERROR, ppWorkingNamespace MUST receive
	// a pointer to an IWbemServices interface pointer to the requested namespace.
	//
	// The output parameter MUST be based on the state of the lFlags parameter (whether
	// WBEM_FLAG_RETURN_IMMEDIATELY is set) as listed in the following table.
	//
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                         SUCCESS                         |                         FAILURE                         |
	//	|                  STATE                   |                        OPERATION                        |                        OPERATION                        |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to the requested IWbemServices interface.   | MUST be set to NULL if the input parameter is not-NULL. |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | MUST be set to NULL if the input parameter is not-NULL. | MUST be set to NULL if the input parameter is not-NULL. |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	WorkingNamespace *wmi.Services `idl:"name:ppWorkingNamespace;pointer:unique" json:"working_namespace"`
	// ppResult: The output parameter MUST be filled according to the state of the lFlags
	// parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY is set) as listed in the following
	// table.
	//
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                         SUCCESS                         |                         FAILURE                         |
	//	|                  STATE                   |                        OPERATION                        |                        OPERATION                        |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to NULL if the input parameter is not-NULL. | MUST be set to NULL if the input parameter is not-NULL. |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | MUST be set to the requested IWbemCallResult interface. | MUST be set to NULL if the input parameter is not-NULL. |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	Result *wmi.CallResult `idl:"name:ppResult;pointer:unique" json:"result"`
}

func (o *OpenNamespaceRequest) xxx_ToOp(ctx context.Context) *xxx_OpenNamespaceOperation {
	if o == nil {
		return &xxx_OpenNamespaceOperation{}
	}
	return &xxx_OpenNamespaceOperation{
		This:             o.This,
		Namespace:        o.Namespace,
		Flags:            o.Flags,
		Context:          o.Context,
		WorkingNamespace: o.WorkingNamespace,
		Result:           o.Result,
	}
}

func (o *OpenNamespaceRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenNamespaceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Namespace = op.Namespace
	o.Flags = op.Flags
	o.Context = op.Context
	o.WorkingNamespace = op.WorkingNamespace
	o.Result = op.Result
}
func (o *OpenNamespaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenNamespaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenNamespaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenNamespaceResponse structure represents the OpenNamespace operation response
type OpenNamespaceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppWorkingNamespace: This parameter MUST NOT be NULL on input when WBEM_FLAG_RETURN_IMMEDIATELY
	// is not set. If the method returns WBEM_S_NO_ERROR, ppWorkingNamespace MUST receive
	// a pointer to an IWbemServices interface pointer to the requested namespace.
	//
	// The output parameter MUST be based on the state of the lFlags parameter (whether
	// WBEM_FLAG_RETURN_IMMEDIATELY is set) as listed in the following table.
	//
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                         SUCCESS                         |                         FAILURE                         |
	//	|                  STATE                   |                        OPERATION                        |                        OPERATION                        |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to the requested IWbemServices interface.   | MUST be set to NULL if the input parameter is not-NULL. |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | MUST be set to NULL if the input parameter is not-NULL. | MUST be set to NULL if the input parameter is not-NULL. |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	WorkingNamespace *wmi.Services `idl:"name:ppWorkingNamespace;pointer:unique" json:"working_namespace"`
	// ppResult: The output parameter MUST be filled according to the state of the lFlags
	// parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY is set) as listed in the following
	// table.
	//
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                         SUCCESS                         |                         FAILURE                         |
	//	|                  STATE                   |                        OPERATION                        |                        OPERATION                        |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to NULL if the input parameter is not-NULL. | MUST be set to NULL if the input parameter is not-NULL. |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | MUST be set to the requested IWbemCallResult interface. | MUST be set to NULL if the input parameter is not-NULL. |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	Result *wmi.CallResult `idl:"name:ppResult;pointer:unique" json:"result"`
	// Return: The OpenNamespace return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenNamespaceResponse) xxx_ToOp(ctx context.Context) *xxx_OpenNamespaceOperation {
	if o == nil {
		return &xxx_OpenNamespaceOperation{}
	}
	return &xxx_OpenNamespaceOperation{
		That:             o.That,
		WorkingNamespace: o.WorkingNamespace,
		Result:           o.Result,
		Return:           o.Return,
	}
}

func (o *OpenNamespaceResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenNamespaceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.WorkingNamespace = op.WorkingNamespace
	o.Result = op.Result
	o.Return = op.Return
}
func (o *OpenNamespaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenNamespaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenNamespaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CancelAsyncCallOperation structure represents the CancelAsyncCall operation
type xxx_CancelAsyncCallOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Sink   *wmi.ObjectSink `idl:"name:pSink" json:"sink"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CancelAsyncCallOperation) OpNum() int { return 4 }

func (o *xxx_CancelAsyncCallOperation) OpName() string { return "/IWbemServices/v0/CancelAsyncCall" }

func (o *xxx_CancelAsyncCallOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelAsyncCallOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pSink {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.Sink != nil {
			_ptr_pSink := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Sink != nil {
					if err := o.Sink.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Sink, _ptr_pSink); err != nil {
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

func (o *xxx_CancelAsyncCallOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pSink {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_pSink := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Sink == nil {
				o.Sink = &wmi.ObjectSink{}
			}
			if err := o.Sink.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pSink := func(ptr interface{}) { o.Sink = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.Sink, _s_pSink, _ptr_pSink); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelAsyncCallOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelAsyncCallOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelAsyncCallOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CancelAsyncCallRequest structure represents the CancelAsyncCall operation request
type CancelAsyncCallRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pSink: MUST be a pointer to the IWbemObjectSink interface object that was passed
	// to the asynchronous method that the client wants to cancel. This parameter MUST NOT
	// be NULL.
	Sink *wmi.ObjectSink `idl:"name:pSink" json:"sink"`
}

func (o *CancelAsyncCallRequest) xxx_ToOp(ctx context.Context) *xxx_CancelAsyncCallOperation {
	if o == nil {
		return &xxx_CancelAsyncCallOperation{}
	}
	return &xxx_CancelAsyncCallOperation{
		This: o.This,
		Sink: o.Sink,
	}
}

func (o *CancelAsyncCallRequest) xxx_FromOp(ctx context.Context, op *xxx_CancelAsyncCallOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Sink = op.Sink
}
func (o *CancelAsyncCallRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CancelAsyncCallRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelAsyncCallOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CancelAsyncCallResponse structure represents the CancelAsyncCall operation response
type CancelAsyncCallResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CancelAsyncCall return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CancelAsyncCallResponse) xxx_ToOp(ctx context.Context) *xxx_CancelAsyncCallOperation {
	if o == nil {
		return &xxx_CancelAsyncCallOperation{}
	}
	return &xxx_CancelAsyncCallOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CancelAsyncCallResponse) xxx_FromOp(ctx context.Context, op *xxx_CancelAsyncCallOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CancelAsyncCallResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CancelAsyncCallResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelAsyncCallOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryObjectSinkOperation structure represents the QueryObjectSink operation
type xxx_QueryObjectSinkOperation struct {
	This            *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Flags           int32           `idl:"name:lFlags" json:"flags"`
	ResponseHandler *wmi.ObjectSink `idl:"name:ppResponseHandler" json:"response_handler"`
	Return          int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryObjectSinkOperation) OpNum() int { return 5 }

func (o *xxx_QueryObjectSinkOperation) OpName() string { return "/IWbemServices/v0/QueryObjectSink" }

func (o *xxx_QueryObjectSinkOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryObjectSinkOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryObjectSinkOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryObjectSinkOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryObjectSinkOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppResponseHandler {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.ResponseHandler != nil {
			_ptr_ppResponseHandler := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResponseHandler != nil {
					if err := o.ResponseHandler.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResponseHandler, _ptr_ppResponseHandler); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryObjectSinkOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppResponseHandler {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_ppResponseHandler := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResponseHandler == nil {
				o.ResponseHandler = &wmi.ObjectSink{}
			}
			if err := o.ResponseHandler.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppResponseHandler := func(ptr interface{}) { o.ResponseHandler = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.ResponseHandler, _s_ppResponseHandler, _ptr_ppResponseHandler); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryObjectSinkRequest structure represents the QueryObjectSink operation request
type QueryObjectSinkRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lFlags: This parameter is not used and its value MUST be 0x0.
	Flags int32 `idl:"name:lFlags" json:"flags"`
}

func (o *QueryObjectSinkRequest) xxx_ToOp(ctx context.Context) *xxx_QueryObjectSinkOperation {
	if o == nil {
		return &xxx_QueryObjectSinkOperation{}
	}
	return &xxx_QueryObjectSinkOperation{
		This:  o.This,
		Flags: o.Flags,
	}
}

func (o *QueryObjectSinkRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryObjectSinkOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
}
func (o *QueryObjectSinkRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryObjectSinkRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryObjectSinkOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryObjectSinkResponse structure represents the QueryObjectSink operation response
type QueryObjectSinkResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppResponseHandler: MUST be a pointer to a QueryObjectSink interface pointer to the
	// notification handler that allows the client to send events directly to the server.
	// This parameter MUST be set to NULL on error.
	ResponseHandler *wmi.ObjectSink `idl:"name:ppResponseHandler" json:"response_handler"`
	// Return: The QueryObjectSink return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryObjectSinkResponse) xxx_ToOp(ctx context.Context) *xxx_QueryObjectSinkOperation {
	if o == nil {
		return &xxx_QueryObjectSinkOperation{}
	}
	return &xxx_QueryObjectSinkOperation{
		That:            o.That,
		ResponseHandler: o.ResponseHandler,
		Return:          o.Return,
	}
}

func (o *QueryObjectSinkResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryObjectSinkOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ResponseHandler = op.ResponseHandler
	o.Return = op.Return
}
func (o *QueryObjectSinkResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryObjectSinkResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryObjectSinkOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetObjectOperation structure represents the GetObject operation
type xxx_GetObjectOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	ObjectPath *oaut.String     `idl:"name:strObjectPath" json:"object_path"`
	Flags      int32            `idl:"name:lFlags" json:"flags"`
	Context    *wmi.Context     `idl:"name:pCtx" json:"context"`
	Object     *wmi.ClassObject `idl:"name:ppObject;pointer:unique" json:"object"`
	CallResult *wmi.CallResult  `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetObjectOperation) OpNum() int { return 6 }

func (o *xxx_GetObjectOperation) OpName() string { return "/IWbemServices/v0/GetObject" }

func (o *xxx_GetObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strObjectPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ObjectPath != nil {
			_ptr_strObjectPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectPath != nil {
					if err := o.ObjectPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectPath, _ptr_strObjectPath); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// ppObject {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.Object != nil {
			_ptr_ppObject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Object != nil {
					_ptr_ppObject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Object != nil {
							if err := o.Object.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Object, _ptr_ppObject); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Object, _ptr_ppObject); err != nil {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.CallResult != nil {
			_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CallResult != nil {
					_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.CallResult != nil {
							if err := o.CallResult.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
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

func (o *xxx_GetObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strObjectPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strObjectPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectPath == nil {
				o.ObjectPath = &oaut.String{}
			}
			if err := o.ObjectPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strObjectPath := func(ptr interface{}) { o.ObjectPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ObjectPath, _s_strObjectPath, _ptr_strObjectPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppObject {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_ppObject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppObject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Object == nil {
					o.Object = &wmi.ClassObject{}
				}
				if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppObject := func(ptr interface{}) { o.Object = *ptr.(**wmi.ClassObject) }
			if err := w.ReadPointer(&o.Object, _s_ppObject, _ptr_ppObject); err != nil {
				return err
			}
			return nil
		})
		_s_ppObject := func(ptr interface{}) { o.Object = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.Object, _s_ppObject, _ptr_ppObject); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.CallResult == nil {
					o.CallResult = &wmi.CallResult{}
				}
				if err := o.CallResult.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppObject {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.Object != nil {
			_ptr_ppObject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Object != nil {
					_ptr_ppObject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Object != nil {
							if err := o.Object.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Object, _ptr_ppObject); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Object, _ptr_ppObject); err != nil {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.CallResult != nil {
			_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CallResult != nil {
					_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.CallResult != nil {
							if err := o.CallResult.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppObject {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_ppObject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppObject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Object == nil {
					o.Object = &wmi.ClassObject{}
				}
				if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppObject := func(ptr interface{}) { o.Object = *ptr.(**wmi.ClassObject) }
			if err := w.ReadPointer(&o.Object, _s_ppObject, _ptr_ppObject); err != nil {
				return err
			}
			return nil
		})
		_s_ppObject := func(ptr interface{}) { o.Object = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.Object, _s_ppObject, _ptr_ppObject); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.CallResult == nil {
					o.CallResult = &wmi.CallResult{}
				}
				if err := o.CallResult.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetObjectRequest structure represents the GetObject operation request
type GetObjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strObjectPath: MUST be the CIM path of the CIM object to be retrieved. If the parameter
	// is NULL, the server MUST return an empty CIM Object.
	ObjectPath *oaut.String `idl:"name:strObjectPath" json:"object_path"`
	// lFlags: Specifies the behavior of the IWbemServices::GetObject method. Flag behavior
	// MUST be interpreted as specified in the following table.
	//
	// The server MUST allow any combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is not set, the server SHOULD return no CIM localizable information. |
	//	|                                             | If this bit is set, the server SHOULD return CIM localizable information for the |
	//	|                                             | CIM object, as specified in section 2.2.6.                                       |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY 0x00000010     | If this bit is not set, the server MUST make the method call synchronously. If   |
	//	|                                             | this bit is set, the server MUST make the method call semisynchronously.         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_DIRECT_READ 0x00000200            | If this bit is set, the server MUST disregard any derived class when it searches |
	//	|                                             | the result. If this bit is not set, the server MUST consider the entire class    |
	//	|                                             | hierarchy when it returns the result.                                            |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass for processing to the implementer of the
	// CIM object that is referred to by strObjectPath. If the parameter is set to NULL,
	// the server MUST ignore it.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// ppObject: If the parameter is set to NULL, the server MUST ignore it. In this case,
	// the output parameter MUST be filled according to the state of the lFlags parameter
	// (whether WBEM_FLAG_RETURN_IMMEDIATELY is set) as listed in the following table.
	//
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                         SUCCESS                         |                         FAILURE                         |
	//	|                  STATE                   |                        OPERATION                        |                        OPERATION                        |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST contain an IWbemClassObject interface pointer.     | MUST be set to NULL if the input parameter is non-NULL. |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | MUST be set to NULL if the input parameter is non-NULL. | MUST be set to NULL if the input parameter is non-NULL. |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	Object *wmi.ClassObject `idl:"name:ppObject;pointer:unique" json:"object"`
	// ppCallResult: The output parameter MUST be filled according to the state of the lFlags
	// parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY is set) as listed in the following
	// table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	|                   FLAG                   |                                     SUCCESS                                      |                               FAILURE                                |
	//	|                  STATE                   |                                    OPERATION                                     |                              OPERATION                               |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to NULL if the ppCallResult input parameter is non-NULL.             | MUST be set to NULL if the ppCallResult input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | The ppCallResult parameter MUST NOT be NULL upon input. If NULL, the server MUST | MUST be set to NULL if the ppCallResult input parameter is non-NULL. |
	//	|                                          | return WBEM_E_INVALID_PARAMETER. Upon output, the parameter MUST contain the     |                                                                      |
	//	|                                          | IWbemCallResult interface pointer.                                               |                                                                      |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
}

func (o *GetObjectRequest) xxx_ToOp(ctx context.Context) *xxx_GetObjectOperation {
	if o == nil {
		return &xxx_GetObjectOperation{}
	}
	return &xxx_GetObjectOperation{
		This:       o.This,
		ObjectPath: o.ObjectPath,
		Flags:      o.Flags,
		Context:    o.Context,
		Object:     o.Object,
		CallResult: o.CallResult,
	}
}

func (o *GetObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_GetObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectPath = op.ObjectPath
	o.Flags = op.Flags
	o.Context = op.Context
	o.Object = op.Object
	o.CallResult = op.CallResult
}
func (o *GetObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetObjectResponse structure represents the GetObject operation response
type GetObjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppObject: If the parameter is set to NULL, the server MUST ignore it. In this case,
	// the output parameter MUST be filled according to the state of the lFlags parameter
	// (whether WBEM_FLAG_RETURN_IMMEDIATELY is set) as listed in the following table.
	//
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                         SUCCESS                         |                         FAILURE                         |
	//	|                  STATE                   |                        OPERATION                        |                        OPERATION                        |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST contain an IWbemClassObject interface pointer.     | MUST be set to NULL if the input parameter is non-NULL. |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | MUST be set to NULL if the input parameter is non-NULL. | MUST be set to NULL if the input parameter is non-NULL. |
	//	+------------------------------------------+---------------------------------------------------------+---------------------------------------------------------+
	Object *wmi.ClassObject `idl:"name:ppObject;pointer:unique" json:"object"`
	// ppCallResult: The output parameter MUST be filled according to the state of the lFlags
	// parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY is set) as listed in the following
	// table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	|                   FLAG                   |                                     SUCCESS                                      |                               FAILURE                                |
	//	|                  STATE                   |                                    OPERATION                                     |                              OPERATION                               |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to NULL if the ppCallResult input parameter is non-NULL.             | MUST be set to NULL if the ppCallResult input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | The ppCallResult parameter MUST NOT be NULL upon input. If NULL, the server MUST | MUST be set to NULL if the ppCallResult input parameter is non-NULL. |
	//	|                                          | return WBEM_E_INVALID_PARAMETER. Upon output, the parameter MUST contain the     |                                                                      |
	//	|                                          | IWbemCallResult interface pointer.                                               |                                                                      |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
	// Return: The GetObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetObjectResponse) xxx_ToOp(ctx context.Context) *xxx_GetObjectOperation {
	if o == nil {
		return &xxx_GetObjectOperation{}
	}
	return &xxx_GetObjectOperation{
		That:       o.That,
		Object:     o.Object,
		CallResult: o.CallResult,
		Return:     o.Return,
	}
}

func (o *GetObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_GetObjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Object = op.Object
	o.CallResult = op.CallResult
	o.Return = op.Return
}
func (o *GetObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetObjectAsyncOperation structure represents the GetObjectAsync operation
type xxx_GetObjectAsyncOperation struct {
	This            *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ObjectPath      *oaut.String    `idl:"name:strObjectPath" json:"object_path"`
	Flags           int32           `idl:"name:lFlags" json:"flags"`
	Context         *wmi.Context    `idl:"name:pCtx" json:"context"`
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
	Return          int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetObjectAsyncOperation) OpNum() int { return 7 }

func (o *xxx_GetObjectAsyncOperation) OpName() string { return "/IWbemServices/v0/GetObjectAsync" }

func (o *xxx_GetObjectAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strObjectPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ObjectPath != nil {
			_ptr_strObjectPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectPath != nil {
					if err := o.ObjectPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectPath, _ptr_strObjectPath); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.ResponseHandler != nil {
			_ptr_pResponseHandler := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResponseHandler != nil {
					if err := o.ResponseHandler.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResponseHandler, _ptr_pResponseHandler); err != nil {
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

func (o *xxx_GetObjectAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strObjectPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strObjectPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectPath == nil {
				o.ObjectPath = &oaut.String{}
			}
			if err := o.ObjectPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strObjectPath := func(ptr interface{}) { o.ObjectPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ObjectPath, _s_strObjectPath, _ptr_strObjectPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_pResponseHandler := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResponseHandler == nil {
				o.ResponseHandler = &wmi.ObjectSink{}
			}
			if err := o.ResponseHandler.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pResponseHandler := func(ptr interface{}) { o.ResponseHandler = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.ResponseHandler, _s_pResponseHandler, _ptr_pResponseHandler); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetObjectAsyncRequest structure represents the GetObjectAsync operation request
type GetObjectAsyncRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strObjectPath: MUST be the CIM path of the CIM object to be retrieved. If this parameter
	// is set to NULL, the server MUST return an empty CIM object.
	ObjectPath *oaut.String `idl:"name:strObjectPath" json:"object_path"`
	// lFlags: Specifies the behavior of the GetObjectAsync method. Flag behavior MUST be
	// interpreted as specified in the following table.
	//
	// The server MUST accept a combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is not set, the server SHOULD return no CIM localizable information. |
	//	|                                             | If this bit is set, the server SHOULD return CIM localizable information for the |
	//	|                                             | CIM object, as specified in section 2.2.6.                                       |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SEND_STATUS 0x00000080            | If this bit is not set, the server MUST make one final                           |
	//	|                                             | IWbemObjectSink::SetStatus call on the interface pointer that is provided in the |
	//	|                                             | pResponseHandler parameter. If this bit is set, the server MAY make intermediate |
	//	|                                             | IWbemObjectSink::SetStatus calls on the interface pointer prior to call          |
	//	|                                             | completion.                                                                      |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_DIRECT_READ 0x00000200            | If this bit is not set, the implementer MUST consider the entire class hierarchy |
	//	|                                             | when it returns the result. If this bit is set, the server MUST disregard any    |
	//	|                                             | derived class when it searches the result.                                       |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to provide to the server about the CIM object referred
	// to by strObjectPath. If pCtx is NULL, the parameter MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// pResponseHandler: MUST be a pointer to the IWbemObjectSink interface that is implemented
	// by the caller, where enumeration results are delivered. The parameter MUST NOT be
	// NULL. If the parameter is NULL, the server MUST return WBEM_E_INVALID_PARAMETER.
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
}

func (o *GetObjectAsyncRequest) xxx_ToOp(ctx context.Context) *xxx_GetObjectAsyncOperation {
	if o == nil {
		return &xxx_GetObjectAsyncOperation{}
	}
	return &xxx_GetObjectAsyncOperation{
		This:            o.This,
		ObjectPath:      o.ObjectPath,
		Flags:           o.Flags,
		Context:         o.Context,
		ResponseHandler: o.ResponseHandler,
	}
}

func (o *GetObjectAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_GetObjectAsyncOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectPath = op.ObjectPath
	o.Flags = op.Flags
	o.Context = op.Context
	o.ResponseHandler = op.ResponseHandler
}
func (o *GetObjectAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetObjectAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetObjectAsyncResponse structure represents the GetObjectAsync operation response
type GetObjectAsyncResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The GetObjectAsync return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetObjectAsyncResponse) xxx_ToOp(ctx context.Context) *xxx_GetObjectAsyncOperation {
	if o == nil {
		return &xxx_GetObjectAsyncOperation{}
	}
	return &xxx_GetObjectAsyncOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *GetObjectAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_GetObjectAsyncOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *GetObjectAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetObjectAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PutClassOperation structure represents the PutClass operation
type xxx_PutClassOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Object     *wmi.ClassObject `idl:"name:pObject" json:"object"`
	Flags      int32            `idl:"name:lFlags" json:"flags"`
	Context    *wmi.Context     `idl:"name:pCtx" json:"context"`
	CallResult *wmi.CallResult  `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_PutClassOperation) OpNum() int { return 8 }

func (o *xxx_PutClassOperation) OpName() string { return "/IWbemServices/v0/PutClass" }

func (o *xxx_PutClassOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutClassOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pObject {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.Object != nil {
			_ptr_pObject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Object != nil {
					if err := o.Object.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Object, _ptr_pObject); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.CallResult != nil {
			_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CallResult != nil {
					_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.CallResult != nil {
							if err := o.CallResult.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
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

func (o *xxx_PutClassOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pObject {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_pObject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Object == nil {
				o.Object = &wmi.ClassObject{}
			}
			if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pObject := func(ptr interface{}) { o.Object = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.Object, _s_pObject, _ptr_pObject); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.CallResult == nil {
					o.CallResult = &wmi.CallResult{}
				}
				if err := o.CallResult.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutClassOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutClassOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.CallResult != nil {
			_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CallResult != nil {
					_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.CallResult != nil {
							if err := o.CallResult.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutClassOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.CallResult == nil {
					o.CallResult = &wmi.CallResult{}
				}
				if err := o.CallResult.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PutClassRequest structure represents the PutClass operation request
type PutClassRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pObject: MUST be a pointer to an IWbemClassObject interface pointer that MUST contain
	// the class information to create a new class or update an existing class. This parameter
	// MUST NOT be NULL.
	Object *wmi.ClassObject `idl:"name:pObject" json:"object"`
	// lFlags: Specifies the behavior of the PutClass method. Flag behavior MUST be interpreted
	// as specified in the following table.
	//
	// The server MUST accept a combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is set, the server SHOULD ignore all the amended qualifiers while it |
	//	|                                             | creates or updates the CIM class.<34> If this bit is not set, the server SHOULD  |
	//	|                                             | include all the qualifiers, including amended qualifiers, while it updates or    |
	//	|                                             | creates the CIM class.                                                           |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY 0x00000010     | If this bit is not set, the server MUST make the method call synchronously. If   |
	//	|                                             | this bit is set, the server MUST make the method call semisynchronously.         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_UPDATE_ONLY 0x00000001            | The server MUST update a CIM class pObject if the CIM class is present. This     |
	//	|                                             | flag is mutually exclusive with WBEM_FLAG_CREATE_ONLY. If none of these flags    |
	//	|                                             | are set, the server MUST create or update a CIM class pObject.                   |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_CREATE_ONLY 0x00000002            | The server MUST create a CIM class pObject if the CIM class is not already       |
	//	|                                             | present.                                                                         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_UPDATE_FORCE_MODE 0x00000040      | The server MUST update the class even if conflicting child classes exist.        |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_UPDATE_SAFE_MODE 0x00000020       | The server MUST update the class as long as the change does not cause any        |
	//	|                                             | conflicts with existing child classes or instances. This flag is mutually        |
	//	|                                             | exclusive with WBEM_FLAG_UPDATE_FORCE_MODE. If none of these flags are set,      |
	//	|                                             | the server MUST update the class if there is no derived class, if there is no    |
	//	|                                             | instance for that class, or if the class is unchanged.                           |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to provide to the server about the CIM class that
	// is referred to by the pObject parameter. If the pCtx parameter is NULL, it MUST be
	// ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// ppCallResult: If the input parameter is non-NULL, the server MUST return WBEM_S_NO_ERROR
	// and IWbemCallResult MUST deliver the result of the requested operation (regardless
	// whether WBEM_FLAG_RETURN_IMMEDIATELY is set). The output parameter MUST be filled
	// according to the state of the lFlags parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY
	// is set) as listed in the following table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                                OPERATION STARTED                                 |                   OPERATION FAILED TO                   |
	//	|                  STATE                   |                                   SUCCESSFULLY                                   |                          START                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to IWbemCallResult if the input parameter is non-NULL.               | MUST be set to NULL if the input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | This parameter MUST NOT be NULL upon input. If NULL, the server MUST             | MUST be set to NULL if the input parameter is non-NULL. |
	//	|                                          | return WBEM_E_INVALID_PARAMETER. On output, the parameter MUST contain the       |                                                         |
	//	|                                          | IWbemCallResult interface pointer.                                               |                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
}

func (o *PutClassRequest) xxx_ToOp(ctx context.Context) *xxx_PutClassOperation {
	if o == nil {
		return &xxx_PutClassOperation{}
	}
	return &xxx_PutClassOperation{
		This:       o.This,
		Object:     o.Object,
		Flags:      o.Flags,
		Context:    o.Context,
		CallResult: o.CallResult,
	}
}

func (o *PutClassRequest) xxx_FromOp(ctx context.Context, op *xxx_PutClassOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Object = op.Object
	o.Flags = op.Flags
	o.Context = op.Context
	o.CallResult = op.CallResult
}
func (o *PutClassRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PutClassRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutClassOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PutClassResponse structure represents the PutClass operation response
type PutClassResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppCallResult: If the input parameter is non-NULL, the server MUST return WBEM_S_NO_ERROR
	// and IWbemCallResult MUST deliver the result of the requested operation (regardless
	// whether WBEM_FLAG_RETURN_IMMEDIATELY is set). The output parameter MUST be filled
	// according to the state of the lFlags parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY
	// is set) as listed in the following table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                                OPERATION STARTED                                 |                   OPERATION FAILED TO                   |
	//	|                  STATE                   |                                   SUCCESSFULLY                                   |                          START                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to IWbemCallResult if the input parameter is non-NULL.               | MUST be set to NULL if the input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | This parameter MUST NOT be NULL upon input. If NULL, the server MUST             | MUST be set to NULL if the input parameter is non-NULL. |
	//	|                                          | return WBEM_E_INVALID_PARAMETER. On output, the parameter MUST contain the       |                                                         |
	//	|                                          | IWbemCallResult interface pointer.                                               |                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
	// Return: The PutClass return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PutClassResponse) xxx_ToOp(ctx context.Context) *xxx_PutClassOperation {
	if o == nil {
		return &xxx_PutClassOperation{}
	}
	return &xxx_PutClassOperation{
		That:       o.That,
		CallResult: o.CallResult,
		Return:     o.Return,
	}
}

func (o *PutClassResponse) xxx_FromOp(ctx context.Context, op *xxx_PutClassOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CallResult = op.CallResult
	o.Return = op.Return
}
func (o *PutClassResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PutClassResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutClassOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PutClassAsyncOperation structure represents the PutClassAsync operation
type xxx_PutClassAsyncOperation struct {
	This            *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Object          *wmi.ClassObject `idl:"name:pObject" json:"object"`
	Flags           int32            `idl:"name:lFlags" json:"flags"`
	Context         *wmi.Context     `idl:"name:pCtx" json:"context"`
	ResponseHandler *wmi.ObjectSink  `idl:"name:pResponseHandler" json:"response_handler"`
	Return          int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_PutClassAsyncOperation) OpNum() int { return 9 }

func (o *xxx_PutClassAsyncOperation) OpName() string { return "/IWbemServices/v0/PutClassAsync" }

func (o *xxx_PutClassAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutClassAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pObject {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.Object != nil {
			_ptr_pObject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Object != nil {
					if err := o.Object.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Object, _ptr_pObject); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.ResponseHandler != nil {
			_ptr_pResponseHandler := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResponseHandler != nil {
					if err := o.ResponseHandler.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResponseHandler, _ptr_pResponseHandler); err != nil {
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

func (o *xxx_PutClassAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pObject {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_pObject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Object == nil {
				o.Object = &wmi.ClassObject{}
			}
			if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pObject := func(ptr interface{}) { o.Object = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.Object, _s_pObject, _ptr_pObject); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_pResponseHandler := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResponseHandler == nil {
				o.ResponseHandler = &wmi.ObjectSink{}
			}
			if err := o.ResponseHandler.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pResponseHandler := func(ptr interface{}) { o.ResponseHandler = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.ResponseHandler, _s_pResponseHandler, _ptr_pResponseHandler); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutClassAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutClassAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutClassAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PutClassAsyncRequest structure represents the PutClassAsync operation request
type PutClassAsyncRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pObject: MUST be a pointer to an IWbemClassObject interface pointer that MUST contain
	// the class information to create a new class or update an existing class. The class
	// that is specified by the parameter MUST have been correctly initialized with all
	// the required property values. This parameter MUST NOT be NULL.
	Object *wmi.ClassObject `idl:"name:pObject" json:"object"`
	// lFlags: Specifies the behavior of the PutClassAsync method. Flag behavior MUST be
	// interpreted as specified in the following table.
	//
	// The server MUST accept a combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is set, the server SHOULD ignore all the amended qualifiers while    |
	//	|                                             | it creates or updates a CIM class.<36> If this bit is not set, the server SHOULD |
	//	|                                             | include all the qualifiers, including amended qualifiers, while it updates or    |
	//	|                                             | creates a CIM class.                                                             |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_UPDATE_ONLY 0x00000001            | The server MUST update a CIM class pObject if the CIM class is present. This     |
	//	|                                             | flag is mutually exclusive with WBEM_FLAG_CREATE_ONLY. If none of these flags    |
	//	|                                             | are set, the server MUST create or update a CIM class pObject.                   |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_CREATE_ONLY 0x00000002            | The server MUST create a CIM class pObject if the CIM class is not already       |
	//	|                                             | present.                                                                         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_UPDATE_FORCE_MODE 0x00000040      | The server MUST forcefully update the class even when conflicting child classes  |
	//	|                                             | exist.                                                                           |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_UPDATE_SAFE_MODE 0x00000020       | The server MUST update the class as long as the change does not cause any        |
	//	|                                             | conflicts with existing child classes or instances. This flag is mutually        |
	//	|                                             | exclusive with WBEM_FLAG_UPDATE_FORCE_MODE. If none of these flags are set,      |
	//	|                                             | the server MUST update the class if there is no derived class, if there is no    |
	//	|                                             | instance for that class, or if the class is unchanged.                           |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SEND_STATUS 0x00000080            | If this bit is not set, the server MUST make one final                           |
	//	|                                             | IWbemObjectSink::SetStatus method call on the interface pointer that is provided |
	//	|                                             | in the pResponseHandler parameter. If this bit is set, the server MAY make       |
	//	|                                             | intermediate IWbemObjectSink::SetStatus calls on the interface pointer prior to  |
	//	|                                             | call completion.                                                                 |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass to the server. If the pCtx parameter is
	// NULL, the parameter MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// pResponseHandler: MUST be a pointer to an IWbemObjectSink interface object that is
	// implemented by the client of this method. The parameter MUST NOT be NULL.
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
}

func (o *PutClassAsyncRequest) xxx_ToOp(ctx context.Context) *xxx_PutClassAsyncOperation {
	if o == nil {
		return &xxx_PutClassAsyncOperation{}
	}
	return &xxx_PutClassAsyncOperation{
		This:            o.This,
		Object:          o.Object,
		Flags:           o.Flags,
		Context:         o.Context,
		ResponseHandler: o.ResponseHandler,
	}
}

func (o *PutClassAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_PutClassAsyncOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Object = op.Object
	o.Flags = op.Flags
	o.Context = op.Context
	o.ResponseHandler = op.ResponseHandler
}
func (o *PutClassAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PutClassAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutClassAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PutClassAsyncResponse structure represents the PutClassAsync operation response
type PutClassAsyncResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PutClassAsync return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PutClassAsyncResponse) xxx_ToOp(ctx context.Context) *xxx_PutClassAsyncOperation {
	if o == nil {
		return &xxx_PutClassAsyncOperation{}
	}
	return &xxx_PutClassAsyncOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PutClassAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_PutClassAsyncOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PutClassAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PutClassAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutClassAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteClassOperation structure represents the DeleteClass operation
type xxx_DeleteClassOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Class      *oaut.String    `idl:"name:strClass" json:"class"`
	Flags      int32           `idl:"name:lFlags" json:"flags"`
	Context    *wmi.Context    `idl:"name:pCtx" json:"context"`
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteClassOperation) OpNum() int { return 10 }

func (o *xxx_DeleteClassOperation) OpName() string { return "/IWbemServices/v0/DeleteClass" }

func (o *xxx_DeleteClassOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteClassOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strClass {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Class != nil {
			_ptr_strClass := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Class != nil {
					if err := o.Class.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Class, _ptr_strClass); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.CallResult != nil {
			_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CallResult != nil {
					_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.CallResult != nil {
							if err := o.CallResult.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
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

func (o *xxx_DeleteClassOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strClass {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strClass := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Class == nil {
				o.Class = &oaut.String{}
			}
			if err := o.Class.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strClass := func(ptr interface{}) { o.Class = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Class, _s_strClass, _ptr_strClass); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.CallResult == nil {
					o.CallResult = &wmi.CallResult{}
				}
				if err := o.CallResult.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteClassOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteClassOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.CallResult != nil {
			_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CallResult != nil {
					_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.CallResult != nil {
							if err := o.CallResult.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteClassOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.CallResult == nil {
					o.CallResult = &wmi.CallResult{}
				}
				if err := o.CallResult.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteClassRequest structure represents the DeleteClass operation request
type DeleteClassRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strClass: MUST be the name of the class to delete. This parameter MUST NOT be NULL.
	Class *oaut.String `idl:"name:strClass" json:"class"`
	// lFlags: Specifies the behavior of the DeleteClass method. Flag behavior MUST be interpreted
	// as specified in the following table.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY 0x00000010 | If this bit is set, the server MUST make the method call semisynchronously. If   |
	//	|                                         | this bit is not set, the server MUST make the method call synchronously.         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass to the server. If pCtx is NULL, the parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// ppCallResult: The output parameter MUST be filled according to the state of the lFlags
	// parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY is set) as listed in the following
	// table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	|                   FLAG                   |                                OPERATION STARTED                                 |                         OPERATION FAILED TO                          |
	//	|                  STATE                   |                                   SUCCESSFULLY                                   |                                START                                 |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to IWbemCallResult if the ppCallResult input parameter is non-NULL.  | MUST be set to NULL if the ppCallResult input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | The ppCallResult parameter MUST NOT be NULL upon input. If NULL, the server      | MUST be set to NULL if the ppCallResult input parameter is non-NULL. |
	//	|                                          | MUST return WBEM_E_INVALID_PARAMETER. On output, the parameter MUST contain the  |                                                                      |
	//	|                                          | IWbemCallResult interface pointer.                                               |                                                                      |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
}

func (o *DeleteClassRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteClassOperation {
	if o == nil {
		return &xxx_DeleteClassOperation{}
	}
	return &xxx_DeleteClassOperation{
		This:       o.This,
		Class:      o.Class,
		Flags:      o.Flags,
		Context:    o.Context,
		CallResult: o.CallResult,
	}
}

func (o *DeleteClassRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteClassOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Class = op.Class
	o.Flags = op.Flags
	o.Context = op.Context
	o.CallResult = op.CallResult
}
func (o *DeleteClassRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteClassRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteClassOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteClassResponse structure represents the DeleteClass operation response
type DeleteClassResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppCallResult: The output parameter MUST be filled according to the state of the lFlags
	// parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY is set) as listed in the following
	// table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	|                   FLAG                   |                                OPERATION STARTED                                 |                         OPERATION FAILED TO                          |
	//	|                  STATE                   |                                   SUCCESSFULLY                                   |                                START                                 |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to IWbemCallResult if the ppCallResult input parameter is non-NULL.  | MUST be set to NULL if the ppCallResult input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | The ppCallResult parameter MUST NOT be NULL upon input. If NULL, the server      | MUST be set to NULL if the ppCallResult input parameter is non-NULL. |
	//	|                                          | MUST return WBEM_E_INVALID_PARAMETER. On output, the parameter MUST contain the  |                                                                      |
	//	|                                          | IWbemCallResult interface pointer.                                               |                                                                      |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
	// Return: The DeleteClass return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteClassResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteClassOperation {
	if o == nil {
		return &xxx_DeleteClassOperation{}
	}
	return &xxx_DeleteClassOperation{
		That:       o.That,
		CallResult: o.CallResult,
		Return:     o.Return,
	}
}

func (o *DeleteClassResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteClassOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CallResult = op.CallResult
	o.Return = op.Return
}
func (o *DeleteClassResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteClassResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteClassOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteClassAsyncOperation structure represents the DeleteClassAsync operation
type xxx_DeleteClassAsyncOperation struct {
	This            *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Class           *oaut.String    `idl:"name:strClass" json:"class"`
	Flags           int32           `idl:"name:lFlags" json:"flags"`
	Context         *wmi.Context    `idl:"name:pCtx" json:"context"`
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
	Return          int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteClassAsyncOperation) OpNum() int { return 11 }

func (o *xxx_DeleteClassAsyncOperation) OpName() string { return "/IWbemServices/v0/DeleteClassAsync" }

func (o *xxx_DeleteClassAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteClassAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strClass {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Class != nil {
			_ptr_strClass := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Class != nil {
					if err := o.Class.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Class, _ptr_strClass); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.ResponseHandler != nil {
			_ptr_pResponseHandler := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResponseHandler != nil {
					if err := o.ResponseHandler.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResponseHandler, _ptr_pResponseHandler); err != nil {
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

func (o *xxx_DeleteClassAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strClass {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strClass := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Class == nil {
				o.Class = &oaut.String{}
			}
			if err := o.Class.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strClass := func(ptr interface{}) { o.Class = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Class, _s_strClass, _ptr_strClass); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_pResponseHandler := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResponseHandler == nil {
				o.ResponseHandler = &wmi.ObjectSink{}
			}
			if err := o.ResponseHandler.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pResponseHandler := func(ptr interface{}) { o.ResponseHandler = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.ResponseHandler, _s_pResponseHandler, _ptr_pResponseHandler); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteClassAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteClassAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteClassAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteClassAsyncRequest structure represents the DeleteClassAsync operation request
type DeleteClassAsyncRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strClass: MUST be the name of the class to delete. This parameter MUST NOT be NULL.
	Class *oaut.String `idl:"name:strClass" json:"class"`
	// lFlags: Specifies the behavior of the DeleteClassAsync method. Flag behavior MUST
	// be interpreted as specified in the following table.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|                                  |                                                                                  |
	//	|              VALUE               |                                     MEANING                                      |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SEND_STATUS 0x00000080 | If this bit is not set, the server MUST make one final                           |
	//	|                                  | IWbemObjectSink::SetStatus call on the interface pointer that is provided in the |
	//	|                                  | pResponseHandler parameter. If this bit is set, the server MAY make intermediate |
	//	|                                  | IWbemObjectSink::SetStatus calls on the interface pointer prior to call          |
	//	|                                  | completion.                                                                      |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass to the server. If pCtx is NULL, the parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// pResponseHandler: MUST be a pointer to an IWbemObjectSink interface object that is
	// implemented by the client of this method. This parameter MUST NOT be NULL.
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
}

func (o *DeleteClassAsyncRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteClassAsyncOperation {
	if o == nil {
		return &xxx_DeleteClassAsyncOperation{}
	}
	return &xxx_DeleteClassAsyncOperation{
		This:            o.This,
		Class:           o.Class,
		Flags:           o.Flags,
		Context:         o.Context,
		ResponseHandler: o.ResponseHandler,
	}
}

func (o *DeleteClassAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteClassAsyncOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Class = op.Class
	o.Flags = op.Flags
	o.Context = op.Context
	o.ResponseHandler = op.ResponseHandler
}
func (o *DeleteClassAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteClassAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteClassAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteClassAsyncResponse structure represents the DeleteClassAsync operation response
type DeleteClassAsyncResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteClassAsync return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteClassAsyncResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteClassAsyncOperation {
	if o == nil {
		return &xxx_DeleteClassAsyncOperation{}
	}
	return &xxx_DeleteClassAsyncOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeleteClassAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteClassAsyncOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteClassAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteClassAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteClassAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateClassEnumOperation structure represents the CreateClassEnum operation
type xxx_CreateClassEnumOperation struct {
	This       *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Superclass *oaut.String         `idl:"name:strSuperclass" json:"superclass"`
	Flags      int32                `idl:"name:lFlags" json:"flags"`
	Context    *wmi.Context         `idl:"name:pCtx" json:"context"`
	Enum       *wmi.EnumClassObject `idl:"name:ppEnum" json:"enum"`
	Return     int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateClassEnumOperation) OpNum() int { return 12 }

func (o *xxx_CreateClassEnumOperation) OpName() string { return "/IWbemServices/v0/CreateClassEnum" }

func (o *xxx_CreateClassEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClassEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strSuperclass {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Superclass != nil {
			_ptr_strSuperclass := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Superclass != nil {
					if err := o.Superclass.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Superclass, _ptr_strSuperclass); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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

func (o *xxx_CreateClassEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strSuperclass {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strSuperclass := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Superclass == nil {
				o.Superclass = &oaut.String{}
			}
			if err := o.Superclass.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strSuperclass := func(ptr interface{}) { o.Superclass = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Superclass, _s_strSuperclass, _ptr_strSuperclass); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClassEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClassEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumWbemClassObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.EnumClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClassEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumWbemClassObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &wmi.EnumClassObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**wmi.EnumClassObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateClassEnumRequest structure represents the CreateClassEnum operation request
type CreateClassEnumRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	Superclass *oaut.String   `idl:"name:strSuperclass" json:"superclass"`
	// lFlags: Flags affect the behavior of the CreateClassEnum method. Flag behavior MUST
	// be interpreted as specified in the following table.
	//
	// The server MUST allow any combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is not set, the server SHOULD return no CIM localizable information. |
	//	|                                             | If this bit is set, the server SHOULD return CIM localizable information for the |
	//	|                                             | CIM object, as specified in section 2.2.6.                                       |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY 0x00000010     | If this bit is not set, the server MUST make the method call synchronously. If   |
	//	|                                             | this bit is set, the server MUST make the method call semisynchronously.         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SHALLOW 0x00000001                | When this bit is not set, the server MUST return all classes that are derived    |
	//	|                                             | from the requested class and all its subclasses. When this bit is set, the       |
	//	|                                             | server MUST return only the classes that are directly derived from the requested |
	//	|                                             | class.                                                                           |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_FORWARD_ONLY 0x00000020           | When this bit is not set, the server MUST return an enumerator that has reset    |
	//	|                                             | capability. When this bit is set, the server MUST return an enumerator that does |
	//	|                                             | not have reset capability, as specified in section 3.1.4.4.                      |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface that MUST contain additional
	// information that the client wants to pass to the server. If the pCtx parameter is
	// NULL, it MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
}

func (o *CreateClassEnumRequest) xxx_ToOp(ctx context.Context) *xxx_CreateClassEnumOperation {
	if o == nil {
		return &xxx_CreateClassEnumOperation{}
	}
	return &xxx_CreateClassEnumOperation{
		This:       o.This,
		Superclass: o.Superclass,
		Flags:      o.Flags,
		Context:    o.Context,
	}
}

func (o *CreateClassEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateClassEnumOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Superclass = op.Superclass
	o.Flags = op.Flags
	o.Context = op.Context
}
func (o *CreateClassEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateClassEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateClassEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateClassEnumResponse structure represents the CreateClassEnum operation response
type CreateClassEnumResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: MUST receive the pointer to the enumerator that implements the IEnumWbemClassObject
	// interface. This parameter MUST NOT be NULL.
	Enum *wmi.EnumClassObject `idl:"name:ppEnum" json:"enum"`
	// Return: The CreateClassEnum return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateClassEnumResponse) xxx_ToOp(ctx context.Context) *xxx_CreateClassEnumOperation {
	if o == nil {
		return &xxx_CreateClassEnumOperation{}
	}
	return &xxx_CreateClassEnumOperation{
		That:   o.That,
		Enum:   o.Enum,
		Return: o.Return,
	}
}

func (o *CreateClassEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateClassEnumOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *CreateClassEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateClassEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateClassEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateClassEnumAsyncOperation structure represents the CreateClassEnumAsync operation
type xxx_CreateClassEnumAsyncOperation struct {
	This            *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Superclass      *oaut.String    `idl:"name:strSuperclass" json:"superclass"`
	Flags           int32           `idl:"name:lFlags" json:"flags"`
	Context         *wmi.Context    `idl:"name:pCtx" json:"context"`
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
	Return          int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateClassEnumAsyncOperation) OpNum() int { return 13 }

func (o *xxx_CreateClassEnumAsyncOperation) OpName() string {
	return "/IWbemServices/v0/CreateClassEnumAsync"
}

func (o *xxx_CreateClassEnumAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClassEnumAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strSuperclass {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Superclass != nil {
			_ptr_strSuperclass := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Superclass != nil {
					if err := o.Superclass.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Superclass, _ptr_strSuperclass); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.ResponseHandler != nil {
			_ptr_pResponseHandler := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResponseHandler != nil {
					if err := o.ResponseHandler.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResponseHandler, _ptr_pResponseHandler); err != nil {
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

func (o *xxx_CreateClassEnumAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strSuperclass {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strSuperclass := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Superclass == nil {
				o.Superclass = &oaut.String{}
			}
			if err := o.Superclass.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strSuperclass := func(ptr interface{}) { o.Superclass = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Superclass, _s_strSuperclass, _ptr_strSuperclass); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_pResponseHandler := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResponseHandler == nil {
				o.ResponseHandler = &wmi.ObjectSink{}
			}
			if err := o.ResponseHandler.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pResponseHandler := func(ptr interface{}) { o.ResponseHandler = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.ResponseHandler, _s_pResponseHandler, _ptr_pResponseHandler); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClassEnumAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClassEnumAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClassEnumAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateClassEnumAsyncRequest structure represents the CreateClassEnumAsync operation request
type CreateClassEnumAsyncRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	Superclass *oaut.String   `idl:"name:strSuperclass" json:"superclass"`
	// lFlags: Flags that affect the behavior of the CreateClassEnum method. Flag behavior
	// MUST be interpreted as specified in the following table.
	//
	// The server MUST allow any combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is not set, the server SHOULD return no CIM localizable information. |
	//	|                                             | If this bit is set, the server SHOULD return CIM localizable information for the |
	//	|                                             | CIM object as specified in section 2.2.6.                                        |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SEND_STATUS 0x00000080            | If this bit is not set, the server MUST make one final                           |
	//	|                                             | IWbemObjectSink::SetStatus call on the interface pointer that is provided in the |
	//	|                                             | pResponseHandler parameter. If this bit is set, the server MAY make intermediate |
	//	|                                             | IWbemObjectSink::SetStatus calls on the interface pointer prior to call          |
	//	|                                             | completion.                                                                      |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SHALLOW 0x00000001                | When this bit is not set, the server MUST return all classes that are derived    |
	//	|                                             | from the requested class and all its subclasses. When this bit is set, the       |
	//	|                                             | server MUST only return the classes that are directly derived from the requested |
	//	|                                             | class.                                                                           |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass to the server. If pCtx is NULL, the parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// pResponseHandler: MUST be a pointer to the IWbemObjectSink that is implemented by
	// the caller, where enumeration results are delivered. The parameter MUST NOT be NULL.
	// In error cases, indicated by the return value, the supplied IWbemObjectSink interface
	// pointer MUST NOT be used. If WBEM_S_NO_ERROR is returned, the user IWbemObjectSink
	// interface pointer MUST be called to indicate the results of the CreateClassEnumAsync
	// operation, as specified later in this section.
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
}

func (o *CreateClassEnumAsyncRequest) xxx_ToOp(ctx context.Context) *xxx_CreateClassEnumAsyncOperation {
	if o == nil {
		return &xxx_CreateClassEnumAsyncOperation{}
	}
	return &xxx_CreateClassEnumAsyncOperation{
		This:            o.This,
		Superclass:      o.Superclass,
		Flags:           o.Flags,
		Context:         o.Context,
		ResponseHandler: o.ResponseHandler,
	}
}

func (o *CreateClassEnumAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateClassEnumAsyncOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Superclass = op.Superclass
	o.Flags = op.Flags
	o.Context = op.Context
	o.ResponseHandler = op.ResponseHandler
}
func (o *CreateClassEnumAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateClassEnumAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateClassEnumAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateClassEnumAsyncResponse structure represents the CreateClassEnumAsync operation response
type CreateClassEnumAsyncResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateClassEnumAsync return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateClassEnumAsyncResponse) xxx_ToOp(ctx context.Context) *xxx_CreateClassEnumAsyncOperation {
	if o == nil {
		return &xxx_CreateClassEnumAsyncOperation{}
	}
	return &xxx_CreateClassEnumAsyncOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CreateClassEnumAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateClassEnumAsyncOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateClassEnumAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateClassEnumAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateClassEnumAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PutInstanceOperation structure represents the PutInstance operation
type xxx_PutInstanceOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Instance   *wmi.ClassObject `idl:"name:pInst" json:"instance"`
	Flags      int32            `idl:"name:lFlags" json:"flags"`
	Context    *wmi.Context     `idl:"name:pCtx" json:"context"`
	CallResult *wmi.CallResult  `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_PutInstanceOperation) OpNum() int { return 14 }

func (o *xxx_PutInstanceOperation) OpName() string { return "/IWbemServices/v0/PutInstance" }

func (o *xxx_PutInstanceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutInstanceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pInst {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.Instance != nil {
			_ptr_pInst := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Instance != nil {
					if err := o.Instance.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Instance, _ptr_pInst); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.CallResult != nil {
			_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CallResult != nil {
					_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.CallResult != nil {
							if err := o.CallResult.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
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

func (o *xxx_PutInstanceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pInst {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_pInst := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Instance == nil {
				o.Instance = &wmi.ClassObject{}
			}
			if err := o.Instance.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pInst := func(ptr interface{}) { o.Instance = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.Instance, _s_pInst, _ptr_pInst); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.CallResult == nil {
					o.CallResult = &wmi.CallResult{}
				}
				if err := o.CallResult.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutInstanceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutInstanceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.CallResult != nil {
			_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CallResult != nil {
					_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.CallResult != nil {
							if err := o.CallResult.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutInstanceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.CallResult == nil {
					o.CallResult = &wmi.CallResult{}
				}
				if err := o.CallResult.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PutInstanceRequest structure represents the PutInstance operation request
type PutInstanceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pInst: MUST be a pointer to an IWbemClassObject interface object that MUST contain
	// the class instance that the client wants to create or update. This parameter MUST
	// NOT be NULL.
	Instance *wmi.ClassObject `idl:"name:pInst" json:"instance"`
	// lFlags: Flags that affect the behavior of the PutInstance method. Flag behavior MUST
	// be interpreted as specified in the following table.
	//
	// The server MUST accept a combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is set, the server SHOULD ignore all the amended qualifiers while    |
	//	|                                             | this method creates or updates a CIM instance. If this bit is not set, the       |
	//	|                                             | server SHOULD include all the qualifiers, including amended qualifiers, while    |
	//	|                                             | this method creates or updates a CIM instance.                                   |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY 0x00000010     | If this bit is not set, the server MUST make the method call synchronously. If   |
	//	|                                             | this bit is set, the server MUST make the method call semisynchronously.         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_UPDATE_ONLY 0x00000001            | The server MUST update a CIM instance pObject if the CIM instance is present.    |
	//	|                                             | This flag is mutually exclusive with WBEM_FLAG_CREATE_ONLY. If none of these     |
	//	|                                             | flags are set, the server MUST create or update a CIM instance pObject.          |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_CREATE_ONLY 0x00000002            | The server MUST create a CIM instance pObject if the CIM instance is not already |
	//	|                                             | present.                                                                         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: This parameter is optional. The pCtx parameter MUST be a pointer to an IWbemContext
	// interface object. The pCtx parameter indicates whether the client is requesting a
	// partial-instance update or a full-instance update. A partial-instance update modifies
	// a subset of the CIM instance properties. In contrast, a full-instance update modifies
	// all the properties. If NULL, this parameter indicates that the client application
	// is requesting a full-instance update. When pCtx is used to perform a partial-instance
	// update, the IWbemContext interface object MUST be filled in with the properties that
	// are specified in the following table. If the IWbemContext interface object does not
	// contain the properties in the table, the method MUST return WBEM_E_INVALID_CONTEXT.
	//
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	//	|        PROPERTY        |                    |                                                                                  |
	//	|          NAME          |        TYPE        |                                   DESCRIPTION                                    |
	//	|                        |                    |                                                                                  |
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	//	| __PUT_EXTENSIONS       | VT_BOOL            | If this property is set to TRUE, one or more of the other IWbemContext values    |
	//	|                        |                    | have been specified. To perform a partial instance update, this property MUST    |
	//	|                        |                    | be set to TRUE and the properties that follow MUST be set as specified in their  |
	//	|                        |                    | respective descriptions.                                                         |
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	//	| __PUT_EXT_STRICT_NULLS | VT_BOOL            | If this property is set to TRUE, the server MUST force the setting of properties |
	//	|                        |                    | to NULL. This parameter is optional.                                             |
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	//	| __PUT_EXT_PROPERTIES   | VT_ARRAY | VT_BSTR | Contains a CIM property list to update. The server MUST ignore the properties    |
	//	|                        |                    | that are not listed. To perform a partial instance update, the list of           |
	//	|                        |                    | properties MUST be specified.                                                    |
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	//	| __PUT_EXT_ATOMIC       | VT_BOOL            | If the return code indicates success, all CIM property updates MUST have been    |
	//	|                        |                    | successful. On failure, the server MUST revert any changes to the original state |
	//	|                        |                    | for all CIM property that was updated. On failure, not a single change MUST      |
	//	|                        |                    | remain. The operation is successful when all properties are updated.             |
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// ppCallResult: If the input parameter is non-NULL, the server MUST return WBEM_S_NO_ERROR
	// and IWbemCallResult MUST deliver the result of the requested operation (regardless
	// whether WBEM_FLAG_RETURN_IMMEDIATELY is set). The output parameter MUST be filled
	// according to the state of the lFlags parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY
	// is set) as listed in the following table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                                OPERATION STARTED                                 |                   OPERATION FAILED TO                   |
	//	|                  STATE                   |                                   SUCCESSFULLY                                   |                          START                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to IWbemCallResult if the input parameter is non-NULL.               | MUST be set to NULL if the input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | This parameter MUST NOT be NULL upon input. If NULL, the server MUST             | MUST be set to NULL if the input parameter is non-NULL. |
	//	|                                          | return WBEM_E_INVALID_PARAMETER. On output, the parameter MUST contain the       |                                                         |
	//	|                                          | IWbemCallResult interface pointer.                                               |                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
}

func (o *PutInstanceRequest) xxx_ToOp(ctx context.Context) *xxx_PutInstanceOperation {
	if o == nil {
		return &xxx_PutInstanceOperation{}
	}
	return &xxx_PutInstanceOperation{
		This:       o.This,
		Instance:   o.Instance,
		Flags:      o.Flags,
		Context:    o.Context,
		CallResult: o.CallResult,
	}
}

func (o *PutInstanceRequest) xxx_FromOp(ctx context.Context, op *xxx_PutInstanceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Instance = op.Instance
	o.Flags = op.Flags
	o.Context = op.Context
	o.CallResult = op.CallResult
}
func (o *PutInstanceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PutInstanceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutInstanceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PutInstanceResponse structure represents the PutInstance operation response
type PutInstanceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppCallResult: If the input parameter is non-NULL, the server MUST return WBEM_S_NO_ERROR
	// and IWbemCallResult MUST deliver the result of the requested operation (regardless
	// whether WBEM_FLAG_RETURN_IMMEDIATELY is set). The output parameter MUST be filled
	// according to the state of the lFlags parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY
	// is set) as listed in the following table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                                OPERATION STARTED                                 |                   OPERATION FAILED TO                   |
	//	|                  STATE                   |                                   SUCCESSFULLY                                   |                          START                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to IWbemCallResult if the input parameter is non-NULL.               | MUST be set to NULL if the input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | This parameter MUST NOT be NULL upon input. If NULL, the server MUST             | MUST be set to NULL if the input parameter is non-NULL. |
	//	|                                          | return WBEM_E_INVALID_PARAMETER. On output, the parameter MUST contain the       |                                                         |
	//	|                                          | IWbemCallResult interface pointer.                                               |                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
	// Return: The PutInstance return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PutInstanceResponse) xxx_ToOp(ctx context.Context) *xxx_PutInstanceOperation {
	if o == nil {
		return &xxx_PutInstanceOperation{}
	}
	return &xxx_PutInstanceOperation{
		That:       o.That,
		CallResult: o.CallResult,
		Return:     o.Return,
	}
}

func (o *PutInstanceResponse) xxx_FromOp(ctx context.Context, op *xxx_PutInstanceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CallResult = op.CallResult
	o.Return = op.Return
}
func (o *PutInstanceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PutInstanceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutInstanceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PutInstanceAsyncOperation structure represents the PutInstanceAsync operation
type xxx_PutInstanceAsyncOperation struct {
	This            *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Instance        *wmi.ClassObject `idl:"name:pInst" json:"instance"`
	Flags           int32            `idl:"name:lFlags" json:"flags"`
	Context         *wmi.Context     `idl:"name:pCtx" json:"context"`
	ResponseHandler *wmi.ObjectSink  `idl:"name:pResponseHandler" json:"response_handler"`
	Return          int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_PutInstanceAsyncOperation) OpNum() int { return 15 }

func (o *xxx_PutInstanceAsyncOperation) OpName() string { return "/IWbemServices/v0/PutInstanceAsync" }

func (o *xxx_PutInstanceAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutInstanceAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pInst {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.Instance != nil {
			_ptr_pInst := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Instance != nil {
					if err := o.Instance.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Instance, _ptr_pInst); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.ResponseHandler != nil {
			_ptr_pResponseHandler := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResponseHandler != nil {
					if err := o.ResponseHandler.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResponseHandler, _ptr_pResponseHandler); err != nil {
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

func (o *xxx_PutInstanceAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pInst {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_pInst := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Instance == nil {
				o.Instance = &wmi.ClassObject{}
			}
			if err := o.Instance.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pInst := func(ptr interface{}) { o.Instance = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.Instance, _s_pInst, _ptr_pInst); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_pResponseHandler := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResponseHandler == nil {
				o.ResponseHandler = &wmi.ObjectSink{}
			}
			if err := o.ResponseHandler.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pResponseHandler := func(ptr interface{}) { o.ResponseHandler = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.ResponseHandler, _s_pResponseHandler, _ptr_pResponseHandler); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutInstanceAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutInstanceAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutInstanceAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PutInstanceAsyncRequest structure represents the PutInstanceAsync operation request
type PutInstanceAsyncRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pInst: MUST be a pointer to an IWbemClassObject interface object that MUST contain
	// the class instance that the client wants to create or update. This parameter MUST
	// NOT be NULL.
	Instance *wmi.ClassObject `idl:"name:pInst" json:"instance"`
	// lFlags: Flags that affect the behavior of the PutInstanceAsync method. Flag behavior
	// MUST be interpreted as specified in the following table.
	//
	// The server MUST accept a combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not comply with this condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is set, the server SHOULD ignore all the amended qualifiers while    |
	//	|                                             | this method creates or updates a CIM instance. If this bit is not set, the       |
	//	|                                             | server SHOULD include all the qualifiers, including amended qualifiers, while    |
	//	|                                             | this method creates or updates a CIM instance.                                   |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_UPDATE_ONLY 0x00000001            | The server MUST update a CIM instance pObject if the CIM instance is present.    |
	//	|                                             | This flag is mutually exclusive with WBEM_FLAG_CREATE_ONLY. If none of these     |
	//	|                                             | flags are set, the server MUST create or update a CIM instance pObject.          |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_CREATE_ONLY 0x00000002            | The server MUST create a CIM instance pObject if the CIM instance is not already |
	//	|                                             | present.                                                                         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SEND_STATUS 0x00000080            | If this bit is not set, the server MUST make one final                           |
	//	|                                             | IWbemObjectSink::SetStatus call on the interface pointer that is provided in the |
	//	|                                             | pResponseHandler parameter. If this bit is set, the server MAY make intermediate |
	//	|                                             | IWbemObjectSink::SetStatus calls on the interface pointer prior to call          |
	//	|                                             | completion.                                                                      |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: This parameter is optional. The pCtx parameter MUST be a pointer to an IWbemContext
	// (section 2.2.13) interface object. The pCtx parameter indicates whether the client
	// is requesting a partial-instance update or full-instance update. A partial-instance
	// update modifies a subset of CIM instance properties; a full-instance update modifies
	// all the properties. If NULL, this parameter indicates that the client application
	// is requesting a full-instance update. When pCtx is used to perform a partial-instance
	// update, the IWbemContext interface MUST be completed with the properties that are
	// specified in the following table. If the IWbemContext interface object does not contain
	// the properties in the table, the method MUST return WBEM_E_INVALID_CONTEXT.
	//
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	//	|        PROPERTY        |                    |                                                                                  |
	//	|          NAME          |        TYPE        |                                   DESCRIPTION                                    |
	//	|                        |                    |                                                                                  |
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	//	| __PUT_EXTENSIONS       | VT_BOOL            | If this property is set to TRUE, one or more of the other IWbemContext values    |
	//	|                        |                    | have been specified. To perform a partial-instance update, this property MUST be |
	//	|                        |                    | set to TRUE.                                                                     |
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	//	| __PUT_EXT_STRICT_NULLS | VT_BOOL            | If this property is set to TRUE, the server MUST force the setting of properties |
	//	|                        |                    | to NULL. This parameter is optional.                                             |
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	//	| __PUT_EXT_PROPERTIES   | VT_ARRAY | VT_BSTR | Contains a CIM property list to update. The server MUST ignore properties that   |
	//	|                        |                    | are not listed. To perform a partial-instance update, the list of properties     |
	//	|                        |                    | MUST be specified.                                                               |
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	//	| __PUT_EXT_ATOMIC       | VT_BOOL            | If the return code indicates success, all CIM property updates MUST have been    |
	//	|                        |                    | successful. On failure, the server MUST revert any changes to the original       |
	//	|                        |                    | state for all CIM property updates. On failure, any changes MUST NOT remain. The |
	//	|                        |                    | operation is successful when all properties are updated.                         |
	//	+------------------------+--------------------+----------------------------------------------------------------------------------+
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// pResponseHandler: MUST be a pointer to an IWbemObjectSink interface object that is
	// implemented by the client of this method. This parameter MUST NOT be NULL.
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
}

func (o *PutInstanceAsyncRequest) xxx_ToOp(ctx context.Context) *xxx_PutInstanceAsyncOperation {
	if o == nil {
		return &xxx_PutInstanceAsyncOperation{}
	}
	return &xxx_PutInstanceAsyncOperation{
		This:            o.This,
		Instance:        o.Instance,
		Flags:           o.Flags,
		Context:         o.Context,
		ResponseHandler: o.ResponseHandler,
	}
}

func (o *PutInstanceAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_PutInstanceAsyncOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Instance = op.Instance
	o.Flags = op.Flags
	o.Context = op.Context
	o.ResponseHandler = op.ResponseHandler
}
func (o *PutInstanceAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PutInstanceAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutInstanceAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PutInstanceAsyncResponse structure represents the PutInstanceAsync operation response
type PutInstanceAsyncResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PutInstanceAsync return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PutInstanceAsyncResponse) xxx_ToOp(ctx context.Context) *xxx_PutInstanceAsyncOperation {
	if o == nil {
		return &xxx_PutInstanceAsyncOperation{}
	}
	return &xxx_PutInstanceAsyncOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PutInstanceAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_PutInstanceAsyncOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PutInstanceAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PutInstanceAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutInstanceAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteInstanceOperation structure represents the DeleteInstance operation
type xxx_DeleteInstanceOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ObjectPath *oaut.String    `idl:"name:strObjectPath" json:"object_path"`
	Flags      int32           `idl:"name:lFlags" json:"flags"`
	Context    *wmi.Context    `idl:"name:pCtx" json:"context"`
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteInstanceOperation) OpNum() int { return 16 }

func (o *xxx_DeleteInstanceOperation) OpName() string { return "/IWbemServices/v0/DeleteInstance" }

func (o *xxx_DeleteInstanceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteInstanceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strObjectPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ObjectPath != nil {
			_ptr_strObjectPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectPath != nil {
					if err := o.ObjectPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectPath, _ptr_strObjectPath); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.CallResult != nil {
			_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CallResult != nil {
					_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.CallResult != nil {
							if err := o.CallResult.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
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

func (o *xxx_DeleteInstanceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strObjectPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strObjectPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectPath == nil {
				o.ObjectPath = &oaut.String{}
			}
			if err := o.ObjectPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strObjectPath := func(ptr interface{}) { o.ObjectPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ObjectPath, _s_strObjectPath, _ptr_strObjectPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.CallResult == nil {
					o.CallResult = &wmi.CallResult{}
				}
				if err := o.CallResult.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteInstanceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteInstanceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.CallResult != nil {
			_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CallResult != nil {
					_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.CallResult != nil {
							if err := o.CallResult.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteInstanceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.CallResult == nil {
					o.CallResult = &wmi.CallResult{}
				}
				if err := o.CallResult.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteInstanceRequest structure represents the DeleteInstance operation request
type DeleteInstanceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strObjectPath: MUST be the CIM path to the class instance that the client wants to
	// delete. This parameter MUST NOT be NULL. The CIM path MUST contain the class name
	// and the value of the key properties.
	ObjectPath *oaut.String `idl:"name:strObjectPath" json:"object_path"`
	// lFlags: Flags that affect the behavior of the IWbemServices::DeleteInstance method.
	// Flag behavior MUST be interpreted as specified in the following table.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY 0x00000010 | If this bit is not set, the server MUST make the method call synchronously. If   |
	//	|                                         | this bit is set, the server MUST make the method call semisynchronously.         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Any other DWORD value that does not match the preceding condition MUST be treated
	// as invalid.
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass to the server. If pCtx is NULL, the parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// ppCallResult: If the input parameter is non-NULL, the server MUST return WBEM_S_NO_ERROR
	// and IWbemCallResult MUST deliver the result of the requested operation (regardless
	// whether WBEM_FLAG_RETURN_IMMEDIATELY is set). The output parameter MUST be filled
	// according to the state of the lFlags parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY
	// is set) as listed in the following table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                                OPERATION STARTED                                 |                   OPERATION FAILED TO                   |
	//	|                  STATE                   |                                   SUCCESSFULLY                                   |                          START                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to IWbemCallResult if the input parameter is non-NULL.               | MUST be set to NULL if the input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | This parameter MUST NOT be NULL upon input. If NULL, the server MUST             | MUST be set to NULL if the input parameter is non-NULL. |
	//	|                                          | return WBEM_E_INVALID_PARAMETER. On output, the parameter MUST contain the       |                                                         |
	//	|                                          | IWbemCallResult interface pointer.                                               |                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
}

func (o *DeleteInstanceRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteInstanceOperation {
	if o == nil {
		return &xxx_DeleteInstanceOperation{}
	}
	return &xxx_DeleteInstanceOperation{
		This:       o.This,
		ObjectPath: o.ObjectPath,
		Flags:      o.Flags,
		Context:    o.Context,
		CallResult: o.CallResult,
	}
}

func (o *DeleteInstanceRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteInstanceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectPath = op.ObjectPath
	o.Flags = op.Flags
	o.Context = op.Context
	o.CallResult = op.CallResult
}
func (o *DeleteInstanceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteInstanceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteInstanceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteInstanceResponse structure represents the DeleteInstance operation response
type DeleteInstanceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppCallResult: If the input parameter is non-NULL, the server MUST return WBEM_S_NO_ERROR
	// and IWbemCallResult MUST deliver the result of the requested operation (regardless
	// whether WBEM_FLAG_RETURN_IMMEDIATELY is set). The output parameter MUST be filled
	// according to the state of the lFlags parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY
	// is set) as listed in the following table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                                OPERATION STARTED                                 |                   OPERATION FAILED TO                   |
	//	|                  STATE                   |                                   SUCCESSFULLY                                   |                          START                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to IWbemCallResult if the input parameter is non-NULL.               | MUST be set to NULL if the input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | This parameter MUST NOT be NULL upon input. If NULL, the server MUST             | MUST be set to NULL if the input parameter is non-NULL. |
	//	|                                          | return WBEM_E_INVALID_PARAMETER. On output, the parameter MUST contain the       |                                                         |
	//	|                                          | IWbemCallResult interface pointer.                                               |                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
	// Return: The DeleteInstance return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteInstanceResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteInstanceOperation {
	if o == nil {
		return &xxx_DeleteInstanceOperation{}
	}
	return &xxx_DeleteInstanceOperation{
		That:       o.That,
		CallResult: o.CallResult,
		Return:     o.Return,
	}
}

func (o *DeleteInstanceResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteInstanceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CallResult = op.CallResult
	o.Return = op.Return
}
func (o *DeleteInstanceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteInstanceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteInstanceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteInstanceAsyncOperation structure represents the DeleteInstanceAsync operation
type xxx_DeleteInstanceAsyncOperation struct {
	This            *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ObjectPath      *oaut.String    `idl:"name:strObjectPath" json:"object_path"`
	Flags           int32           `idl:"name:lFlags" json:"flags"`
	Context         *wmi.Context    `idl:"name:pCtx" json:"context"`
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
	Return          int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteInstanceAsyncOperation) OpNum() int { return 17 }

func (o *xxx_DeleteInstanceAsyncOperation) OpName() string {
	return "/IWbemServices/v0/DeleteInstanceAsync"
}

func (o *xxx_DeleteInstanceAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteInstanceAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strObjectPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ObjectPath != nil {
			_ptr_strObjectPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectPath != nil {
					if err := o.ObjectPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectPath, _ptr_strObjectPath); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.ResponseHandler != nil {
			_ptr_pResponseHandler := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResponseHandler != nil {
					if err := o.ResponseHandler.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResponseHandler, _ptr_pResponseHandler); err != nil {
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

func (o *xxx_DeleteInstanceAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strObjectPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strObjectPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectPath == nil {
				o.ObjectPath = &oaut.String{}
			}
			if err := o.ObjectPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strObjectPath := func(ptr interface{}) { o.ObjectPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ObjectPath, _s_strObjectPath, _ptr_strObjectPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_pResponseHandler := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResponseHandler == nil {
				o.ResponseHandler = &wmi.ObjectSink{}
			}
			if err := o.ResponseHandler.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pResponseHandler := func(ptr interface{}) { o.ResponseHandler = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.ResponseHandler, _s_pResponseHandler, _ptr_pResponseHandler); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteInstanceAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteInstanceAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteInstanceAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteInstanceAsyncRequest structure represents the DeleteInstanceAsync operation request
type DeleteInstanceAsyncRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strObjectPath: MUST be the CIM path to the class instance that the client wants to
	// delete. This parameter MUST NOT be NULL. The CIM path MUST contain the class name
	// and the value of the key properties.
	ObjectPath *oaut.String `idl:"name:strObjectPath" json:"object_path"`
	// lFlags: Flags that affect the behavior of the IWbemServices::DeleteInstanceAsync
	// method. Flag behavior MUST be interpreted as specified in the following table.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|                                  |                                                                                  |
	//	|              VALUE               |                                     MEANING                                      |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SEND_STATUS 0x00000080 | If this bit is not set, the server MUST make one final                           |
	//	|                                  | IWbemObjectSink::SetStatus call on the interface pointer that is provided in the |
	//	|                                  | pResponseHandler parameter. If this bit is set, the server MAY make intermediate |
	//	|                                  | IWbemObjectSink::SetStatus calls on the interface pointer prior to call          |
	//	|                                  | completion.                                                                      |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which contains additional information
	// that the client wants to pass to the server. If pCtx is NULL, the parameter MUST
	// be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// pResponseHandler: MUST be a pointer to an IWbemObjectSink interface object that is
	// implemented by the client of this method. This parameter MUST NOT be NULL.
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
}

func (o *DeleteInstanceAsyncRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteInstanceAsyncOperation {
	if o == nil {
		return &xxx_DeleteInstanceAsyncOperation{}
	}
	return &xxx_DeleteInstanceAsyncOperation{
		This:            o.This,
		ObjectPath:      o.ObjectPath,
		Flags:           o.Flags,
		Context:         o.Context,
		ResponseHandler: o.ResponseHandler,
	}
}

func (o *DeleteInstanceAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteInstanceAsyncOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectPath = op.ObjectPath
	o.Flags = op.Flags
	o.Context = op.Context
	o.ResponseHandler = op.ResponseHandler
}
func (o *DeleteInstanceAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteInstanceAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteInstanceAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteInstanceAsyncResponse structure represents the DeleteInstanceAsync operation response
type DeleteInstanceAsyncResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteInstanceAsync return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteInstanceAsyncResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteInstanceAsyncOperation {
	if o == nil {
		return &xxx_DeleteInstanceAsyncOperation{}
	}
	return &xxx_DeleteInstanceAsyncOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeleteInstanceAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteInstanceAsyncOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteInstanceAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteInstanceAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteInstanceAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateInstanceEnumOperation structure represents the CreateInstanceEnum operation
type xxx_CreateInstanceEnumOperation struct {
	This       *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Superclass *oaut.String         `idl:"name:strSuperClass" json:"superclass"`
	Flags      int32                `idl:"name:lFlags" json:"flags"`
	Context    *wmi.Context         `idl:"name:pCtx" json:"context"`
	Enum       *wmi.EnumClassObject `idl:"name:ppEnum" json:"enum"`
	Return     int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateInstanceEnumOperation) OpNum() int { return 18 }

func (o *xxx_CreateInstanceEnumOperation) OpName() string {
	return "/IWbemServices/v0/CreateInstanceEnum"
}

func (o *xxx_CreateInstanceEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateInstanceEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strSuperClass {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Superclass != nil {
			_ptr_strSuperClass := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Superclass != nil {
					if err := o.Superclass.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Superclass, _ptr_strSuperClass); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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

func (o *xxx_CreateInstanceEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strSuperClass {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strSuperClass := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Superclass == nil {
				o.Superclass = &oaut.String{}
			}
			if err := o.Superclass.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strSuperClass := func(ptr interface{}) { o.Superclass = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Superclass, _s_strSuperClass, _ptr_strSuperClass); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateInstanceEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateInstanceEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumWbemClassObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.EnumClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateInstanceEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumWbemClassObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &wmi.EnumClassObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**wmi.EnumClassObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateInstanceEnumRequest structure represents the CreateInstanceEnum operation request
type CreateInstanceEnumRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strSuperClass: MUST contain the name of the CIM class for which the client wants
	// instances. This parameter MUST NOT be NULL.
	Superclass *oaut.String `idl:"name:strSuperClass" json:"superclass"`
	// lFlags: Flags that affect the behavior of the CreateInstanceEnum method. Flag behavior
	// MUST be interpreted as specified in the following table.
	//
	// The server MUST allow any combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is not set, the server SHOULD return no CIM localizable information. |
	//	|                                             | If this bit is set, the server SHOULD return CIM localizable information for the |
	//	|                                             | CIM object, as specified in section 2.2.6.                                       |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY 0x00000010     | If this bit is not set, the server MUST make the method call synchronously. If   |
	//	|                                             | this bit is set, the server MUST make the method call semisynchronously.         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_DIRECT_READ 0x00000200            | If this bit is not set, the server MUST consider the entire class hierarchy when |
	//	|                                             | it returns the result. If this bit is set, the server MUST disregard any derived |
	//	|                                             | class when it searches the result.                                               |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SHALLOW 0x00000001                | If this bit is set, the server MUST return instances of the requested class only |
	//	|                                             | and MUST exclude instances of classes that are derived from the requested class. |
	//	|                                             | If this bit is not set, the server MUST return all instances of the requested    |
	//	|                                             | class as well as instances of classes that are derived from the requested class. |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_FORWARD_ONLY 0x00000020           | If this bit is not set, the server MUST return an enumerator that has reset      |
	//	|                                             | capability. If this bit is set, the server MUST return an enumerator that does   |
	//	|                                             | not have reset capability, as specified in section 3.1.4.4.                      |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which contains additional information
	// that the client wants to pass to the server. If pCtx is NULL, the parameter MUST
	// be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
}

func (o *CreateInstanceEnumRequest) xxx_ToOp(ctx context.Context) *xxx_CreateInstanceEnumOperation {
	if o == nil {
		return &xxx_CreateInstanceEnumOperation{}
	}
	return &xxx_CreateInstanceEnumOperation{
		This:       o.This,
		Superclass: o.Superclass,
		Flags:      o.Flags,
		Context:    o.Context,
	}
}

func (o *CreateInstanceEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateInstanceEnumOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Superclass = op.Superclass
	o.Flags = op.Flags
	o.Context = op.Context
}
func (o *CreateInstanceEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateInstanceEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateInstanceEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateInstanceEnumResponse structure represents the CreateInstanceEnum operation response
type CreateInstanceEnumResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: MUST receive the pointer to the enumerator that is used to enumerate through
	// the returned class instances, which implements the IEnumWbemClassObject interface.
	// This parameter MUST NOT be NULL.
	Enum *wmi.EnumClassObject `idl:"name:ppEnum" json:"enum"`
	// Return: The CreateInstanceEnum return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateInstanceEnumResponse) xxx_ToOp(ctx context.Context) *xxx_CreateInstanceEnumOperation {
	if o == nil {
		return &xxx_CreateInstanceEnumOperation{}
	}
	return &xxx_CreateInstanceEnumOperation{
		That:   o.That,
		Enum:   o.Enum,
		Return: o.Return,
	}
}

func (o *CreateInstanceEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateInstanceEnumOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *CreateInstanceEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateInstanceEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateInstanceEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateInstanceEnumAsyncOperation structure represents the CreateInstanceEnumAsync operation
type xxx_CreateInstanceEnumAsyncOperation struct {
	This            *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Superclass      *oaut.String    `idl:"name:strSuperClass" json:"superclass"`
	Flags           int32           `idl:"name:lFlags" json:"flags"`
	Context         *wmi.Context    `idl:"name:pCtx" json:"context"`
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
	Return          int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateInstanceEnumAsyncOperation) OpNum() int { return 19 }

func (o *xxx_CreateInstanceEnumAsyncOperation) OpName() string {
	return "/IWbemServices/v0/CreateInstanceEnumAsync"
}

func (o *xxx_CreateInstanceEnumAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateInstanceEnumAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strSuperClass {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Superclass != nil {
			_ptr_strSuperClass := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Superclass != nil {
					if err := o.Superclass.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Superclass, _ptr_strSuperClass); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.ResponseHandler != nil {
			_ptr_pResponseHandler := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResponseHandler != nil {
					if err := o.ResponseHandler.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResponseHandler, _ptr_pResponseHandler); err != nil {
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

func (o *xxx_CreateInstanceEnumAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strSuperClass {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strSuperClass := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Superclass == nil {
				o.Superclass = &oaut.String{}
			}
			if err := o.Superclass.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strSuperClass := func(ptr interface{}) { o.Superclass = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Superclass, _s_strSuperClass, _ptr_strSuperClass); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_pResponseHandler := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResponseHandler == nil {
				o.ResponseHandler = &wmi.ObjectSink{}
			}
			if err := o.ResponseHandler.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pResponseHandler := func(ptr interface{}) { o.ResponseHandler = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.ResponseHandler, _s_pResponseHandler, _ptr_pResponseHandler); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateInstanceEnumAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateInstanceEnumAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateInstanceEnumAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateInstanceEnumAsyncRequest structure represents the CreateInstanceEnumAsync operation request
type CreateInstanceEnumAsyncRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strSuperClass: MUST contain the name of the CIM class for which the client wants
	// instances. This parameter MUST NOT be NULL.
	Superclass *oaut.String `idl:"name:strSuperClass" json:"superclass"`
	// lFlags: Flags that affect the behavior of the IWbemServices::CreateInstanceEnumAsync
	// method. Flag behavior MUST be interpreted as specified in the following table.
	//
	// The server MUST allow any combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is not set, the server SHOULD return no CIM localizable information. |
	//	|                                             | If this bit is set, the server SHOULD return CIM localizable information for the |
	//	|                                             | CIM object, as specified in section 2.2.6.                                       |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SEND_STATUS 0x00000080            | If this bit is not set the server MUST make one final IWbemObjectSink::SetStatus |
	//	|                                             | call on the interface pointer that is provided in the pResponseHandler           |
	//	|                                             | parameter. If this bit is set, the server MAY make intermediate                  |
	//	|                                             | IWbemObjectSink::SetStatus calls on the interface pointer prior to call          |
	//	|                                             | completion.                                                                      |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_DIRECT_READ 0x00000200            | If this bit is not set, the server MUST consider the entire class hierarchy when |
	//	|                                             | it returns the result. If this bit is set, the server MUST disregard any derived |
	//	|                                             | class when it searches the result.                                               |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SHALLOW 0x00000001                | If this bit is set, the server MUST return instances of the requested class only |
	//	|                                             | and MUST exclude instances of classes that are derived from the requested class. |
	//	|                                             | If this bit is not set, the server MUST return all instances of the requested    |
	//	|                                             | class as well as instances of classes that are derived from the requested class. |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass to the server. If pCtx is NULL, the parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// pResponseHandler: MUST be a pointer to the IWbemObjectSink interface that is implemented
	// by the caller and where enumeration results are delivered. The parameter MUST NOT
	// be NULL.
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
}

func (o *CreateInstanceEnumAsyncRequest) xxx_ToOp(ctx context.Context) *xxx_CreateInstanceEnumAsyncOperation {
	if o == nil {
		return &xxx_CreateInstanceEnumAsyncOperation{}
	}
	return &xxx_CreateInstanceEnumAsyncOperation{
		This:            o.This,
		Superclass:      o.Superclass,
		Flags:           o.Flags,
		Context:         o.Context,
		ResponseHandler: o.ResponseHandler,
	}
}

func (o *CreateInstanceEnumAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateInstanceEnumAsyncOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Superclass = op.Superclass
	o.Flags = op.Flags
	o.Context = op.Context
	o.ResponseHandler = op.ResponseHandler
}
func (o *CreateInstanceEnumAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateInstanceEnumAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateInstanceEnumAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateInstanceEnumAsyncResponse structure represents the CreateInstanceEnumAsync operation response
type CreateInstanceEnumAsyncResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateInstanceEnumAsync return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateInstanceEnumAsyncResponse) xxx_ToOp(ctx context.Context) *xxx_CreateInstanceEnumAsyncOperation {
	if o == nil {
		return &xxx_CreateInstanceEnumAsyncOperation{}
	}
	return &xxx_CreateInstanceEnumAsyncOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CreateInstanceEnumAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateInstanceEnumAsyncOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateInstanceEnumAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateInstanceEnumAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateInstanceEnumAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExecQueryOperation structure represents the ExecQuery operation
type xxx_ExecQueryOperation struct {
	This          *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat       `idl:"name:That" json:"that"`
	QueryLanguage *oaut.String         `idl:"name:strQueryLanguage" json:"query_language"`
	Query         *oaut.String         `idl:"name:strQuery" json:"query"`
	Flags         int32                `idl:"name:lFlags" json:"flags"`
	Context       *wmi.Context         `idl:"name:pCtx" json:"context"`
	Enum          *wmi.EnumClassObject `idl:"name:ppEnum" json:"enum"`
	Return        int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_ExecQueryOperation) OpNum() int { return 20 }

func (o *xxx_ExecQueryOperation) OpName() string { return "/IWbemServices/v0/ExecQuery" }

func (o *xxx_ExecQueryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecQueryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strQueryLanguage {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.QueryLanguage != nil {
			_ptr_strQueryLanguage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueryLanguage != nil {
					if err := o.QueryLanguage.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryLanguage, _ptr_strQueryLanguage); err != nil {
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
	// strQuery {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Query != nil {
			_ptr_strQuery := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Query != nil {
					if err := o.Query.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Query, _ptr_strQuery); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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

func (o *xxx_ExecQueryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strQueryLanguage {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strQueryLanguage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueryLanguage == nil {
				o.QueryLanguage = &oaut.String{}
			}
			if err := o.QueryLanguage.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strQueryLanguage := func(ptr interface{}) { o.QueryLanguage = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.QueryLanguage, _s_strQueryLanguage, _ptr_strQueryLanguage); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// strQuery {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strQuery := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Query == nil {
				o.Query = &oaut.String{}
			}
			if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strQuery := func(ptr interface{}) { o.Query = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Query, _s_strQuery, _ptr_strQuery); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecQueryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecQueryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumWbemClassObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.EnumClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecQueryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumWbemClassObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &wmi.EnumClassObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**wmi.EnumClassObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ExecQueryRequest structure represents the ExecQuery operation request
type ExecQueryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strQueryLanguage: MUST be set to "WQL".
	QueryLanguage *oaut.String `idl:"name:strQueryLanguage" json:"query_language"`
	// strQuery: MUST contain the "WQL" query text as specified in [UNICODE] (UTF-16) and
	// in section 2.2.1. This parameter MUST NOT be NULL.
	Query *oaut.String `idl:"name:strQuery" json:"query"`
	// lFlags: Specifies the behavior of the IWbemServices::ExecQuery method. Flag behavior
	// MUST be interpreted as specified in the following table.
	//
	// The server MUST allow any combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is not set, the server SHOULD not return CIM localizable             |
	//	|                                             | information. If this bit is set, the server SHOULD return CIM localizable        |
	//	|                                             | information for the CIM object, as specified in section 2.2.6.                   |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY 0x00000010     | If this bit is not set, the server MUST make the method call synchronously. If   |
	//	|                                             | this bit is set, the server MUST make the method call semisynchronously.         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_DIRECT_READ 0x00000200            | If this bit is not set, the server MUST consider the entire class hierarchy when |
	//	|                                             | it returns the result. If this bit is set, the server MUST disregard any derived |
	//	|                                             | class when it searches the result.                                               |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_PROTOTYPE 0x00000002              | If this bit is not set, the server MUST run the query. If this bit is set, the   |
	//	|                                             | server MUST only return the class schema of the resulting objects.               |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_FORWARD_ONLY 0x00000020           | If this bit is not set, the server MUST return an enumerator that has reset      |
	//	|                                             | capability. If this bit is set, the server MUST return an enumerator without     |
	//	|                                             | reset capability, as specified in section 3.1.4.4.                               |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass to the server. If pCtx is NULL, the parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
}

func (o *ExecQueryRequest) xxx_ToOp(ctx context.Context) *xxx_ExecQueryOperation {
	if o == nil {
		return &xxx_ExecQueryOperation{}
	}
	return &xxx_ExecQueryOperation{
		This:          o.This,
		QueryLanguage: o.QueryLanguage,
		Query:         o.Query,
		Flags:         o.Flags,
		Context:       o.Context,
	}
}

func (o *ExecQueryRequest) xxx_FromOp(ctx context.Context, op *xxx_ExecQueryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QueryLanguage = op.QueryLanguage
	o.Query = op.Query
	o.Flags = op.Flags
	o.Context = op.Context
}
func (o *ExecQueryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExecQueryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecQueryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExecQueryResponse structure represents the ExecQuery operation response
type ExecQueryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: MUST receive the pointer to the IEnumWbemClassObject that is used to enumerate
	// through the CIM objects that are returned for the query result set. This parameter
	// MUST NOT be NULL.
	Enum *wmi.EnumClassObject `idl:"name:ppEnum" json:"enum"`
	// Return: The ExecQuery return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExecQueryResponse) xxx_ToOp(ctx context.Context) *xxx_ExecQueryOperation {
	if o == nil {
		return &xxx_ExecQueryOperation{}
	}
	return &xxx_ExecQueryOperation{
		That:   o.That,
		Enum:   o.Enum,
		Return: o.Return,
	}
}

func (o *ExecQueryResponse) xxx_FromOp(ctx context.Context, op *xxx_ExecQueryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *ExecQueryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExecQueryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecQueryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExecQueryAsyncOperation structure represents the ExecQueryAsync operation
type xxx_ExecQueryAsyncOperation struct {
	This            *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat  `idl:"name:That" json:"that"`
	QueryLanguage   *oaut.String    `idl:"name:strQueryLanguage" json:"query_language"`
	Query           *oaut.String    `idl:"name:strQuery" json:"query"`
	Flags           int32           `idl:"name:lFlags" json:"flags"`
	Context         *wmi.Context    `idl:"name:pCtx" json:"context"`
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
	Return          int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_ExecQueryAsyncOperation) OpNum() int { return 21 }

func (o *xxx_ExecQueryAsyncOperation) OpName() string { return "/IWbemServices/v0/ExecQueryAsync" }

func (o *xxx_ExecQueryAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecQueryAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strQueryLanguage {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.QueryLanguage != nil {
			_ptr_strQueryLanguage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueryLanguage != nil {
					if err := o.QueryLanguage.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryLanguage, _ptr_strQueryLanguage); err != nil {
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
	// strQuery {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Query != nil {
			_ptr_strQuery := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Query != nil {
					if err := o.Query.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Query, _ptr_strQuery); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.ResponseHandler != nil {
			_ptr_pResponseHandler := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResponseHandler != nil {
					if err := o.ResponseHandler.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResponseHandler, _ptr_pResponseHandler); err != nil {
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

func (o *xxx_ExecQueryAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strQueryLanguage {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strQueryLanguage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueryLanguage == nil {
				o.QueryLanguage = &oaut.String{}
			}
			if err := o.QueryLanguage.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strQueryLanguage := func(ptr interface{}) { o.QueryLanguage = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.QueryLanguage, _s_strQueryLanguage, _ptr_strQueryLanguage); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// strQuery {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strQuery := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Query == nil {
				o.Query = &oaut.String{}
			}
			if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strQuery := func(ptr interface{}) { o.Query = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Query, _s_strQuery, _ptr_strQuery); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_pResponseHandler := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResponseHandler == nil {
				o.ResponseHandler = &wmi.ObjectSink{}
			}
			if err := o.ResponseHandler.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pResponseHandler := func(ptr interface{}) { o.ResponseHandler = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.ResponseHandler, _s_pResponseHandler, _ptr_pResponseHandler); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecQueryAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecQueryAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecQueryAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ExecQueryAsyncRequest structure represents the ExecQueryAsync operation request
type ExecQueryAsyncRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strQueryLanguage: MUST be set to "WQL".
	QueryLanguage *oaut.String `idl:"name:strQueryLanguage" json:"query_language"`
	// strQuery: MUST contain the WQL query text as specified in section 2.2.1. This parameter
	// MUST NOT be NULL.
	Query *oaut.String `idl:"name:strQuery" json:"query"`
	// lFlags: Specifies the behavior of the IWbemServices::ExecQueryAsync method. Flag
	// behavior MUST be interpreted as specified in the following table.
	//
	// The server MUST allow any combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is not set, the server SHOULD not return CIM localizable             |
	//	|                                             | information. If this bit is set, the server SHOULD return CIM localizable        |
	//	|                                             | information for the CIM object, as specified in section 2.2.6.                   |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SEND_STATUS 0x00000080            | If this bit is not set the server MUST make one final IWbemObjectSink::SetStatus |
	//	|                                             | call on the interface pointer that is provided in the pResponseHandler           |
	//	|                                             | parameter. If this bit is set, the server MAY make intermediate                  |
	//	|                                             | IWbemObjectSink::SetStatus calls on the interface pointer prior to call          |
	//	|                                             | completion.                                                                      |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_PROTOTYPE 0x00000002              | If this bit is not set, the server MUST run the query. If this bit is set, the   |
	//	|                                             | server MUST only return the class schema of the resulting objects.               |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_DIRECT_READ 0x00000200            | If this bit is not set, the server MUST consider the entire class hierarchy when |
	//	|                                             | it returns the result. If this bit is set, the server MUST disregard any derived |
	//	|                                             | class when it searches the result.                                               |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass to the server. If pCtx is NULL, the parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// pResponseHandler: MUST be a pointer to the IWbemObjectSink interface that is implemented
	// by the caller, where enumeration results are delivered. The parameter MUST NOT be
	// NULL.
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
}

func (o *ExecQueryAsyncRequest) xxx_ToOp(ctx context.Context) *xxx_ExecQueryAsyncOperation {
	if o == nil {
		return &xxx_ExecQueryAsyncOperation{}
	}
	return &xxx_ExecQueryAsyncOperation{
		This:            o.This,
		QueryLanguage:   o.QueryLanguage,
		Query:           o.Query,
		Flags:           o.Flags,
		Context:         o.Context,
		ResponseHandler: o.ResponseHandler,
	}
}

func (o *ExecQueryAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_ExecQueryAsyncOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QueryLanguage = op.QueryLanguage
	o.Query = op.Query
	o.Flags = op.Flags
	o.Context = op.Context
	o.ResponseHandler = op.ResponseHandler
}
func (o *ExecQueryAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExecQueryAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecQueryAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExecQueryAsyncResponse structure represents the ExecQueryAsync operation response
type ExecQueryAsyncResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExecQueryAsync return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExecQueryAsyncResponse) xxx_ToOp(ctx context.Context) *xxx_ExecQueryAsyncOperation {
	if o == nil {
		return &xxx_ExecQueryAsyncOperation{}
	}
	return &xxx_ExecQueryAsyncOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ExecQueryAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_ExecQueryAsyncOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ExecQueryAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExecQueryAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecQueryAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExecNotificationQueryOperation structure represents the ExecNotificationQuery operation
type xxx_ExecNotificationQueryOperation struct {
	This          *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat       `idl:"name:That" json:"that"`
	QueryLanguage *oaut.String         `idl:"name:strQueryLanguage" json:"query_language"`
	Query         *oaut.String         `idl:"name:strQuery" json:"query"`
	Flags         int32                `idl:"name:lFlags" json:"flags"`
	Context       *wmi.Context         `idl:"name:pCtx" json:"context"`
	Enum          *wmi.EnumClassObject `idl:"name:ppEnum" json:"enum"`
	Return        int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_ExecNotificationQueryOperation) OpNum() int { return 22 }

func (o *xxx_ExecNotificationQueryOperation) OpName() string {
	return "/IWbemServices/v0/ExecNotificationQuery"
}

func (o *xxx_ExecNotificationQueryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecNotificationQueryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strQueryLanguage {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.QueryLanguage != nil {
			_ptr_strQueryLanguage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueryLanguage != nil {
					if err := o.QueryLanguage.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryLanguage, _ptr_strQueryLanguage); err != nil {
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
	// strQuery {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Query != nil {
			_ptr_strQuery := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Query != nil {
					if err := o.Query.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Query, _ptr_strQuery); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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

func (o *xxx_ExecNotificationQueryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strQueryLanguage {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strQueryLanguage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueryLanguage == nil {
				o.QueryLanguage = &oaut.String{}
			}
			if err := o.QueryLanguage.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strQueryLanguage := func(ptr interface{}) { o.QueryLanguage = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.QueryLanguage, _s_strQueryLanguage, _ptr_strQueryLanguage); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// strQuery {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strQuery := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Query == nil {
				o.Query = &oaut.String{}
			}
			if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strQuery := func(ptr interface{}) { o.Query = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Query, _s_strQuery, _ptr_strQuery); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecNotificationQueryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecNotificationQueryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumWbemClassObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.EnumClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecNotificationQueryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumWbemClassObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &wmi.EnumClassObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**wmi.EnumClassObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ExecNotificationQueryRequest structure represents the ExecNotificationQuery operation request
type ExecNotificationQueryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strQueryLanguage: MUST be set to "WQL".
	QueryLanguage *oaut.String `idl:"name:strQueryLanguage" json:"query_language"`
	// strQuery: MUST contain the WQL event-related query text as specified in section 2.2.1.
	// This parameter MUST NOT be NULL.
	Query *oaut.String `idl:"name:strQuery" json:"query"`
	// lFlags: Specifies the behavior of the IWbemServices::ExecNotificationQuery method.
	// Flag behavior MUST be interpreted as specified in the following table.
	//
	// The server MUST allow any combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is not set, the server SHOULD return no CIM localizable information. |
	//	|                                             | If this bit is set, the server SHOULD return CIM localizable information for the |
	//	|                                             | CIM object, as specified in section 2.2.6.                                       |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY 0x00000010     | If this bit is set, the server MUST make the method call semisynchronously. This |
	//	|                                             | flag MUST always be set.                                                         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_FORWARD_ONLY 0x00000020           | If this bit is set, the server MUST return an enumerator that does not have      |
	//	|                                             | reset capability, as specified in section 3.1.4.4. This flag MUST always be set. |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass to the server. If pCtx is NULL, the parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
}

func (o *ExecNotificationQueryRequest) xxx_ToOp(ctx context.Context) *xxx_ExecNotificationQueryOperation {
	if o == nil {
		return &xxx_ExecNotificationQueryOperation{}
	}
	return &xxx_ExecNotificationQueryOperation{
		This:          o.This,
		QueryLanguage: o.QueryLanguage,
		Query:         o.Query,
		Flags:         o.Flags,
		Context:       o.Context,
	}
}

func (o *ExecNotificationQueryRequest) xxx_FromOp(ctx context.Context, op *xxx_ExecNotificationQueryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QueryLanguage = op.QueryLanguage
	o.Query = op.Query
	o.Flags = op.Flags
	o.Context = op.Context
}
func (o *ExecNotificationQueryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExecNotificationQueryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecNotificationQueryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExecNotificationQueryResponse structure represents the ExecNotificationQuery operation response
type ExecNotificationQueryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: MUST receive the pointer to the IEnumWbemClassObject that is used to enumerate
	// through the CIM objects that are returned for the query result set. This parameter
	// MUST NOT be NULL.
	Enum *wmi.EnumClassObject `idl:"name:ppEnum" json:"enum"`
	// Return: The ExecNotificationQuery return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExecNotificationQueryResponse) xxx_ToOp(ctx context.Context) *xxx_ExecNotificationQueryOperation {
	if o == nil {
		return &xxx_ExecNotificationQueryOperation{}
	}
	return &xxx_ExecNotificationQueryOperation{
		That:   o.That,
		Enum:   o.Enum,
		Return: o.Return,
	}
}

func (o *ExecNotificationQueryResponse) xxx_FromOp(ctx context.Context, op *xxx_ExecNotificationQueryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *ExecNotificationQueryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExecNotificationQueryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecNotificationQueryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExecNotificationQueryAsyncOperation structure represents the ExecNotificationQueryAsync operation
type xxx_ExecNotificationQueryAsyncOperation struct {
	This            *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat  `idl:"name:That" json:"that"`
	QueryLanguage   *oaut.String    `idl:"name:strQueryLanguage" json:"query_language"`
	Query           *oaut.String    `idl:"name:strQuery" json:"query"`
	Flags           int32           `idl:"name:lFlags" json:"flags"`
	Context         *wmi.Context    `idl:"name:pCtx" json:"context"`
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
	Return          int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_ExecNotificationQueryAsyncOperation) OpNum() int { return 23 }

func (o *xxx_ExecNotificationQueryAsyncOperation) OpName() string {
	return "/IWbemServices/v0/ExecNotificationQueryAsync"
}

func (o *xxx_ExecNotificationQueryAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecNotificationQueryAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strQueryLanguage {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.QueryLanguage != nil {
			_ptr_strQueryLanguage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueryLanguage != nil {
					if err := o.QueryLanguage.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryLanguage, _ptr_strQueryLanguage); err != nil {
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
	// strQuery {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Query != nil {
			_ptr_strQuery := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Query != nil {
					if err := o.Query.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Query, _ptr_strQuery); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.ResponseHandler != nil {
			_ptr_pResponseHandler := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResponseHandler != nil {
					if err := o.ResponseHandler.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResponseHandler, _ptr_pResponseHandler); err != nil {
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

func (o *xxx_ExecNotificationQueryAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strQueryLanguage {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strQueryLanguage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueryLanguage == nil {
				o.QueryLanguage = &oaut.String{}
			}
			if err := o.QueryLanguage.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strQueryLanguage := func(ptr interface{}) { o.QueryLanguage = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.QueryLanguage, _s_strQueryLanguage, _ptr_strQueryLanguage); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// strQuery {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strQuery := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Query == nil {
				o.Query = &oaut.String{}
			}
			if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strQuery := func(ptr interface{}) { o.Query = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Query, _s_strQuery, _ptr_strQuery); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_pResponseHandler := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResponseHandler == nil {
				o.ResponseHandler = &wmi.ObjectSink{}
			}
			if err := o.ResponseHandler.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pResponseHandler := func(ptr interface{}) { o.ResponseHandler = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.ResponseHandler, _s_pResponseHandler, _ptr_pResponseHandler); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecNotificationQueryAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecNotificationQueryAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecNotificationQueryAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ExecNotificationQueryAsyncRequest structure represents the ExecNotificationQueryAsync operation request
type ExecNotificationQueryAsyncRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strQueryLanguage: MUST be set to "WQL".
	QueryLanguage *oaut.String `idl:"name:strQueryLanguage" json:"query_language"`
	// strQuery: MUST contain the WQL event-related query text as specified in section 2.2.1.
	// This parameter MUST NOT be NULL.
	Query *oaut.String `idl:"name:strQuery" json:"query"`
	// lFlags: Specifies the behavior of the IWbemServices::ExecNotificationQueryAsync method.
	// Flag behavior MUST be interpreted as specified in the following table.
	//
	// The server MUST allow any combination of zero or more flags from the following table
	// and MUST comply with all the restrictions in a flag description. Any other DWORD
	// value that does not match a flag condition MUST be treated as not valid.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_USE_AMENDED_QUALIFIERS 0x00020000 | If this bit is not set, the server SHOULD return no CIM localizable information. |
	//	|                                             | If this bit is set, the server SHOULD return CIM localizable information.        |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SEND_STATUS 0x00000080            | This flag is ignored.                                                            |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass to the server. If pCtx is NULL, this parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// pResponseHandler: MUST be a pointer to the IWbemObjectSink interface that is implemented
	// by the caller, where enumeration results are delivered. This parameter MUST NOT be
	// NULL.
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
}

func (o *ExecNotificationQueryAsyncRequest) xxx_ToOp(ctx context.Context) *xxx_ExecNotificationQueryAsyncOperation {
	if o == nil {
		return &xxx_ExecNotificationQueryAsyncOperation{}
	}
	return &xxx_ExecNotificationQueryAsyncOperation{
		This:            o.This,
		QueryLanguage:   o.QueryLanguage,
		Query:           o.Query,
		Flags:           o.Flags,
		Context:         o.Context,
		ResponseHandler: o.ResponseHandler,
	}
}

func (o *ExecNotificationQueryAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_ExecNotificationQueryAsyncOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QueryLanguage = op.QueryLanguage
	o.Query = op.Query
	o.Flags = op.Flags
	o.Context = op.Context
	o.ResponseHandler = op.ResponseHandler
}
func (o *ExecNotificationQueryAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExecNotificationQueryAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecNotificationQueryAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExecNotificationQueryAsyncResponse structure represents the ExecNotificationQueryAsync operation response
type ExecNotificationQueryAsyncResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExecNotificationQueryAsync return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExecNotificationQueryAsyncResponse) xxx_ToOp(ctx context.Context) *xxx_ExecNotificationQueryAsyncOperation {
	if o == nil {
		return &xxx_ExecNotificationQueryAsyncOperation{}
	}
	return &xxx_ExecNotificationQueryAsyncOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ExecNotificationQueryAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_ExecNotificationQueryAsyncOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ExecNotificationQueryAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExecNotificationQueryAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecNotificationQueryAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExecMethodOperation structure represents the ExecMethod operation
type xxx_ExecMethodOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	ObjectPath *oaut.String     `idl:"name:strObjectPath" json:"object_path"`
	MethodName *oaut.String     `idl:"name:strMethodName" json:"method_name"`
	Flags      int32            `idl:"name:lFlags" json:"flags"`
	Context    *wmi.Context     `idl:"name:pCtx" json:"context"`
	InParams   *wmi.ClassObject `idl:"name:pInParams" json:"in_params"`
	OutParams  *wmi.ClassObject `idl:"name:ppOutParams;pointer:unique" json:"out_params"`
	CallResult *wmi.CallResult  `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_ExecMethodOperation) OpNum() int { return 24 }

func (o *xxx_ExecMethodOperation) OpName() string { return "/IWbemServices/v0/ExecMethod" }

func (o *xxx_ExecMethodOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecMethodOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strObjectPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ObjectPath != nil {
			_ptr_strObjectPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectPath != nil {
					if err := o.ObjectPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectPath, _ptr_strObjectPath); err != nil {
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
	// strMethodName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MethodName != nil {
			_ptr_strMethodName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MethodName != nil {
					if err := o.MethodName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MethodName, _ptr_strMethodName); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// pInParams {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.InParams != nil {
			_ptr_pInParams := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InParams != nil {
					if err := o.InParams.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InParams, _ptr_pInParams); err != nil {
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
	// ppOutParams {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.OutParams != nil {
			_ptr_ppOutParams := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OutParams != nil {
					_ptr_ppOutParams := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.OutParams != nil {
							if err := o.OutParams.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.OutParams, _ptr_ppOutParams); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OutParams, _ptr_ppOutParams); err != nil {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.CallResult != nil {
			_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CallResult != nil {
					_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.CallResult != nil {
							if err := o.CallResult.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
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

func (o *xxx_ExecMethodOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strObjectPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strObjectPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectPath == nil {
				o.ObjectPath = &oaut.String{}
			}
			if err := o.ObjectPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strObjectPath := func(ptr interface{}) { o.ObjectPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ObjectPath, _s_strObjectPath, _ptr_strObjectPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// strMethodName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strMethodName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MethodName == nil {
				o.MethodName = &oaut.String{}
			}
			if err := o.MethodName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strMethodName := func(ptr interface{}) { o.MethodName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MethodName, _s_strMethodName, _ptr_strMethodName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pInParams {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_pInParams := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InParams == nil {
				o.InParams = &wmi.ClassObject{}
			}
			if err := o.InParams.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pInParams := func(ptr interface{}) { o.InParams = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.InParams, _s_pInParams, _ptr_pInParams); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppOutParams {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_ppOutParams := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppOutParams := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.OutParams == nil {
					o.OutParams = &wmi.ClassObject{}
				}
				if err := o.OutParams.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppOutParams := func(ptr interface{}) { o.OutParams = *ptr.(**wmi.ClassObject) }
			if err := w.ReadPointer(&o.OutParams, _s_ppOutParams, _ptr_ppOutParams); err != nil {
				return err
			}
			return nil
		})
		_s_ppOutParams := func(ptr interface{}) { o.OutParams = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.OutParams, _s_ppOutParams, _ptr_ppOutParams); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.CallResult == nil {
					o.CallResult = &wmi.CallResult{}
				}
				if err := o.CallResult.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecMethodOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecMethodOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppOutParams {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.OutParams != nil {
			_ptr_ppOutParams := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OutParams != nil {
					_ptr_ppOutParams := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.OutParams != nil {
							if err := o.OutParams.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.OutParams, _ptr_ppOutParams); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OutParams, _ptr_ppOutParams); err != nil {
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
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		if o.CallResult != nil {
			_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CallResult != nil {
					_ptr_ppCallResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.CallResult != nil {
							if err := o.CallResult.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&wmi.CallResult{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CallResult, _ptr_ppCallResult); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecMethodOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppOutParams {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_ppOutParams := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppOutParams := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.OutParams == nil {
					o.OutParams = &wmi.ClassObject{}
				}
				if err := o.OutParams.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppOutParams := func(ptr interface{}) { o.OutParams = *ptr.(**wmi.ClassObject) }
			if err := w.ReadPointer(&o.OutParams, _s_ppOutParams, _ptr_ppOutParams); err != nil {
				return err
			}
			return nil
		})
		_s_ppOutParams := func(ptr interface{}) { o.OutParams = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.OutParams, _s_ppOutParams, _ptr_ppOutParams); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppCallResult {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IWbemCallResult}(interface))
	{
		_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppCallResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.CallResult == nil {
					o.CallResult = &wmi.CallResult{}
				}
				if err := o.CallResult.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
			if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
				return err
			}
			return nil
		})
		_s_ppCallResult := func(ptr interface{}) { o.CallResult = *ptr.(**wmi.CallResult) }
		if err := w.ReadPointer(&o.CallResult, _s_ppCallResult, _ptr_ppCallResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ExecMethodRequest structure represents the ExecMethod operation request
type ExecMethodRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strObjectPath: MUST be the CIM path to the class or instance that implements the
	// method. This parameter MUST NOT be NULL. The CIM path MUST contain the class name
	// and the value of the key properties.
	ObjectPath *oaut.String `idl:"name:strObjectPath" json:"object_path"`
	// strMethodName: MUST be the name of the method to be executed. This parameter MUST
	// NOT be NULL.
	MethodName *oaut.String `idl:"name:strMethodName" json:"method_name"`
	// lFlags: Specifies the behavior of the IWbemServices::ExecMethod method. Flag behavior
	// MUST be interpreted as specified in the following table.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY 0x00000010 | If this bit is not set, the server MUST make the method call synchronously. If   |
	//	|                                         | this bit is set, the server MUST make the method call semisynchronously.         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass to the server. If pCtx is NULL, the parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// pInParams: MUST be a pointer to an IWbemClassObject interface pointer, which MUST
	// contain an instance of input parameter CIM class as defined in [MS-WMIO] (section
	// 2.3.3), with method parameter values set as properties. This parameter MUST be NULL
	// when the method has no input parameters.
	InParams *wmi.ClassObject `idl:"name:pInParams" json:"in_params"`
	// ppOutParams: The output parameter MUST be filled according to the state of the lFlags
	// parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY is set) as listed in the following
	// table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                                     SUCCESS                                      |                         FAILURE                         |
	//	|                  STATE                   |                                    OPERATION                                     |                        OPERATION                        |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | This parameter MUST NOT be NULL upon input. If NULL, the server MUST             | MUST be set to NULL if the input parameter is non-NULL. |
	//	|                                          | return WBEM_E_INVALID_PARAMETER. Upon output, the parameter MUST contain an      |                                                         |
	//	|                                          | IWbemClassObject interface pointer.                                              |                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | MUST return NULL.                                                                | MUST be set to NULL if the input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	OutParams *wmi.ClassObject `idl:"name:ppOutParams;pointer:unique" json:"out_params"`
	// ppCallResult: In this situation, the output parameter MUST be filled according to
	// the state of the lFlags parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY is set) as
	// listed in the following table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	|                                          |                                     SUCCESS                                      |                               FAILURE                                |
	//	|                CONDITION                 |                                    OPERATION                                     |                              OPERATION                               |
	//	|                                          |                                                                                  |                                                                      |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to IWbemCallResult if the ppCallResult input parameter is non-NULL.  | MUST be set to NULL if the ppCallResult input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | The ppCallResult parameter MUST NOT be NULL upon input. If NULL, the server MUST | MUST be set to NULL if the ppCallResult input parameter is non-NULL. |
	//	|                                          | return WBEM_E_INVALID_PARAMETER. Upon output, the parameter MUST contain the     |                                                                      |
	//	|                                          | IWbemCallResult interface pointer.                                               |                                                                      |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
}

func (o *ExecMethodRequest) xxx_ToOp(ctx context.Context) *xxx_ExecMethodOperation {
	if o == nil {
		return &xxx_ExecMethodOperation{}
	}
	return &xxx_ExecMethodOperation{
		This:       o.This,
		ObjectPath: o.ObjectPath,
		MethodName: o.MethodName,
		Flags:      o.Flags,
		Context:    o.Context,
		InParams:   o.InParams,
		OutParams:  o.OutParams,
		CallResult: o.CallResult,
	}
}

func (o *ExecMethodRequest) xxx_FromOp(ctx context.Context, op *xxx_ExecMethodOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectPath = op.ObjectPath
	o.MethodName = op.MethodName
	o.Flags = op.Flags
	o.Context = op.Context
	o.InParams = op.InParams
	o.OutParams = op.OutParams
	o.CallResult = op.CallResult
}
func (o *ExecMethodRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExecMethodRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecMethodOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExecMethodResponse structure represents the ExecMethod operation response
type ExecMethodResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppOutParams: The output parameter MUST be filled according to the state of the lFlags
	// parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY is set) as listed in the following
	// table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	|                   FLAG                   |                                     SUCCESS                                      |                         FAILURE                         |
	//	|                  STATE                   |                                    OPERATION                                     |                        OPERATION                        |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | This parameter MUST NOT be NULL upon input. If NULL, the server MUST             | MUST be set to NULL if the input parameter is non-NULL. |
	//	|                                          | return WBEM_E_INVALID_PARAMETER. Upon output, the parameter MUST contain an      |                                                         |
	//	|                                          | IWbemClassObject interface pointer.                                              |                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | MUST return NULL.                                                                | MUST be set to NULL if the input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------------+
	OutParams *wmi.ClassObject `idl:"name:ppOutParams;pointer:unique" json:"out_params"`
	// ppCallResult: In this situation, the output parameter MUST be filled according to
	// the state of the lFlags parameter (whether WBEM_FLAG_RETURN_IMMEDIATELY is set) as
	// listed in the following table.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	|                                          |                                     SUCCESS                                      |                               FAILURE                                |
	//	|                CONDITION                 |                                    OPERATION                                     |                              OPERATION                               |
	//	|                                          |                                                                                  |                                                                      |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is not set. | MUST be set to IWbemCallResult if the ppCallResult input parameter is non-NULL.  | MUST be set to NULL if the ppCallResult input parameter is non-NULL. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	//	| WBEM_FLAG_RETURN_IMMEDIATELY is set.     | The ppCallResult parameter MUST NOT be NULL upon input. If NULL, the server MUST | MUST be set to NULL if the ppCallResult input parameter is non-NULL. |
	//	|                                          | return WBEM_E_INVALID_PARAMETER. Upon output, the parameter MUST contain the     |                                                                      |
	//	|                                          | IWbemCallResult interface pointer.                                               |                                                                      |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------+
	CallResult *wmi.CallResult `idl:"name:ppCallResult;pointer:unique" json:"call_result"`
	// Return: The ExecMethod return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExecMethodResponse) xxx_ToOp(ctx context.Context) *xxx_ExecMethodOperation {
	if o == nil {
		return &xxx_ExecMethodOperation{}
	}
	return &xxx_ExecMethodOperation{
		That:       o.That,
		OutParams:  o.OutParams,
		CallResult: o.CallResult,
		Return:     o.Return,
	}
}

func (o *ExecMethodResponse) xxx_FromOp(ctx context.Context, op *xxx_ExecMethodOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.OutParams = op.OutParams
	o.CallResult = op.CallResult
	o.Return = op.Return
}
func (o *ExecMethodResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExecMethodResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecMethodOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExecMethodAsyncOperation structure represents the ExecMethodAsync operation
type xxx_ExecMethodAsyncOperation struct {
	This            *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat   `idl:"name:That" json:"that"`
	ObjectPath      *oaut.String     `idl:"name:strObjectPath" json:"object_path"`
	MethodName      *oaut.String     `idl:"name:strMethodName" json:"method_name"`
	Flags           int32            `idl:"name:lFlags" json:"flags"`
	Context         *wmi.Context     `idl:"name:pCtx" json:"context"`
	InParams        *wmi.ClassObject `idl:"name:pInParams" json:"in_params"`
	ResponseHandler *wmi.ObjectSink  `idl:"name:pResponseHandler" json:"response_handler"`
	Return          int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_ExecMethodAsyncOperation) OpNum() int { return 25 }

func (o *xxx_ExecMethodAsyncOperation) OpName() string { return "/IWbemServices/v0/ExecMethodAsync" }

func (o *xxx_ExecMethodAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecMethodAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strObjectPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ObjectPath != nil {
			_ptr_strObjectPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectPath != nil {
					if err := o.ObjectPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectPath, _ptr_strObjectPath); err != nil {
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
	// strMethodName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MethodName != nil {
			_ptr_strMethodName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MethodName != nil {
					if err := o.MethodName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MethodName, _ptr_strMethodName); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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
	// pInParams {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.InParams != nil {
			_ptr_pInParams := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InParams != nil {
					if err := o.InParams.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InParams, _ptr_pInParams); err != nil {
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
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.ResponseHandler != nil {
			_ptr_pResponseHandler := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResponseHandler != nil {
					if err := o.ResponseHandler.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResponseHandler, _ptr_pResponseHandler); err != nil {
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

func (o *xxx_ExecMethodAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strObjectPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strObjectPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectPath == nil {
				o.ObjectPath = &oaut.String{}
			}
			if err := o.ObjectPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strObjectPath := func(ptr interface{}) { o.ObjectPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ObjectPath, _s_strObjectPath, _ptr_strObjectPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// strMethodName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strMethodName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MethodName == nil {
				o.MethodName = &oaut.String{}
			}
			if err := o.MethodName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strMethodName := func(ptr interface{}) { o.MethodName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MethodName, _s_strMethodName, _ptr_strMethodName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pInParams {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_pInParams := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InParams == nil {
				o.InParams = &wmi.ClassObject{}
			}
			if err := o.InParams.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pInParams := func(ptr interface{}) { o.InParams = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.InParams, _s_pInParams, _ptr_pInParams); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResponseHandler {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_pResponseHandler := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResponseHandler == nil {
				o.ResponseHandler = &wmi.ObjectSink{}
			}
			if err := o.ResponseHandler.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pResponseHandler := func(ptr interface{}) { o.ResponseHandler = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.ResponseHandler, _s_pResponseHandler, _ptr_pResponseHandler); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecMethodAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecMethodAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecMethodAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ExecMethodAsyncRequest structure represents the ExecMethodAsync operation request
type ExecMethodAsyncRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strObjectPath: MUST be the CIM path to the class or instance that implements the
	// method. This parameter MUST NOT be NULL. The CIM path MUST contain the class name
	// and the value of the key properties.
	ObjectPath *oaut.String `idl:"name:strObjectPath" json:"object_path"`
	// strMethodName: MUST be the name of the method to be executed. This parameter MUST
	// NOT be NULL.
	MethodName *oaut.String `idl:"name:strMethodName" json:"method_name"`
	// lFlags: Specifies the behavior of the ExecMethodAsync method. Flag behavior MUST
	// be interpreted as specified in the following table.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|                                  |                                                                                  |
	//	|              VALUE               |                                     MEANING                                      |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| WBEM_FLAG_SEND_STATUS 0x00000080 | If this bit is not set, the server MUST make just one final                      |
	//	|                                  | IWbemObjectSink::SetStatus call on the interface pointer that is provided in the |
	//	|                                  | pResponseHandler parameter. If this bit is set, the server MAY make intermediate |
	//	|                                  | IWbemObjectSink::SetStatus calls on the interface pointer prior to call          |
	//	|                                  | completion.                                                                      |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information that the client wants to pass to the server. If pCtx is NULL, the parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
	// pInParams: MUST be a pointer to an IWbemClassObject interface pointer, which MUST
	// contain an instance of input parameter CIM class as defined in [MS-WMIO] (section
	// 2.3.3), with method parameter values set as properties. This parameter MUST be NULL
	// when the method has no input parameters.
	InParams *wmi.ClassObject `idl:"name:pInParams" json:"in_params"`
	// pResponseHandler: MUST be a pointer to an IWbemObjectSink interface object that is
	// implemented by the client of this method. This parameter MUST NOT be NULL.
	ResponseHandler *wmi.ObjectSink `idl:"name:pResponseHandler" json:"response_handler"`
}

func (o *ExecMethodAsyncRequest) xxx_ToOp(ctx context.Context) *xxx_ExecMethodAsyncOperation {
	if o == nil {
		return &xxx_ExecMethodAsyncOperation{}
	}
	return &xxx_ExecMethodAsyncOperation{
		This:            o.This,
		ObjectPath:      o.ObjectPath,
		MethodName:      o.MethodName,
		Flags:           o.Flags,
		Context:         o.Context,
		InParams:        o.InParams,
		ResponseHandler: o.ResponseHandler,
	}
}

func (o *ExecMethodAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_ExecMethodAsyncOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectPath = op.ObjectPath
	o.MethodName = op.MethodName
	o.Flags = op.Flags
	o.Context = op.Context
	o.InParams = op.InParams
	o.ResponseHandler = op.ResponseHandler
}
func (o *ExecMethodAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExecMethodAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecMethodAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExecMethodAsyncResponse structure represents the ExecMethodAsync operation response
type ExecMethodAsyncResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExecMethodAsync return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExecMethodAsyncResponse) xxx_ToOp(ctx context.Context) *xxx_ExecMethodAsyncOperation {
	if o == nil {
		return &xxx_ExecMethodAsyncOperation{}
	}
	return &xxx_ExecMethodAsyncOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ExecMethodAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_ExecMethodAsyncOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ExecMethodAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExecMethodAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecMethodAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
