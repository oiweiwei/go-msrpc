package witness

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// Witness server interface.
type WitnessServer interface {

	// The WitnessrGetInterfaceList method returns information about the interfaces to which
	// witness client connections can be made.
	//
	// Return Values: Returns 0x00000000 (ERROR_SUCCESS) on success or a nonzero error code,
	// as specified in [MS-ERREF] section 2.2. The most common error codes are listed in
	// the following table.
	//
	//	+------------------------------------+--------------------------------------------------------------+
	//	|               RETURN               |                                                              |
	//	|             VALUE/CODE             |                         DESCRIPTION                          |
	//	|                                    |                                                              |
	//	+------------------------------------+--------------------------------------------------------------+
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The operation completed successfully.                        |
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                            |
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The parameter is incorrect.                                  |
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS     | No more data is available.                                   |
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0x0000000E ERROR_OUTOFMEMORY       | There is not enough storage space to complete the operation. |
	//	+------------------------------------+--------------------------------------------------------------+
	//
	// If there are no entries in the InterfaceList, the server MUST fail the request and
	// return the error code ERROR_NO_MORE_ITEMS.
	//
	// If no entry in the InterfaceList has a State of AVAILABLE, the server MUST wait until
	// at least one entry enters that State, as specified in section 3.1.6.1.
	//
	// For each Interface in the InterfaceList, the server MUST construct a WITNESS_INTERFACE_INFO
	// structure as follows:
	//
	// * The *InterfaceGroupName* field of the WITNESS_INTERFACE_INFO structure MUST be
	// set to *Interface.InterfaceGroupName*.
	//
	// * The *State* field MUST be set to *Interface.State*.
	//
	// * The *Version* field MUST be set to *WitnessServiceVersion*.
	//
	// * If *Interface.IPv4Address* is not empty, the *IPV4* field MUST be set to *Interface.IPv4Address*
	// , and IPv4 flag MUST be set in the *Flags* field.
	//
	// * If *Interface.IPv6Address* is not empty, the *IPV6* field MUST be set to *Interface.IPv6Address*
	// , and IPv6 flag MUST be set in the *Flags* field.
	//
	// * In an implementation-dependent manner, the server MUST determine if the *IPv4Address*
	// or *IPv6Address* match any interface which is hosted on the server and the server
	// is also running this Witness Service instance. If the address is not hosted on the
	// local server, the INTERFACE_WITNESS flag MUST be set in the *Flags* field. Otherwise,
	// the flag MUST NOT be set.
	GetInterfaceList(context.Context, *GetInterfaceListRequest) (*GetInterfaceListResponse, error)

	// The WitnessrRegister method allows the witness client to register for resource state
	// change notifications of a NetName and IPAddress. The client can subsequently call
	// the WitnessrAsyncNotify method to receive notifications when there is a state change
	// on any of these resources.
	//
	// Return Values: Returns 0x00000000 (ERROR_SUCCESS) on success or a nonzero error code,
	// as specified in [MS-ERREF] section 2.2. The most common error codes are listed in
	// the following table.
	//
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	|                RETURN                |                                                                        |
	//	|              VALUE/CODE              |                              DESCRIPTION                               |
	//	|                                      |                                                                        |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS             | The operation completed successfully.                                  |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED       | Access is denied.                                                      |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x000005AA ERROR_NO_SYSTEM_RESOURCES | Insufficient system resources exist to complete the requested service. |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER   | The parameter is incorrect.                                            |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x0000139F ERROR_INVALID_STATE       | The specified resource state is invalid.                               |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x0000051A ERROR_REVISION_MISMATCH   | The client request contains an invalid Witness protocol version.       |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//
	// If the Version field of the request is not 0x00010001, the server MUST stop processing
	// the request and return the error code ERROR_REVISION_MISMATCH.
	//
	// If NetName, IpAddress or ClientComputerName is NULL, the server MUST fail the request
	// and return the error code ERROR_INVALID_PARAMETER.
	//
	// If the NetName parameter is not equal to ServerGlobalName, the server MUST fail the
	// request and return the error code ERROR_INVALID_PARAMETER.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)

	// The WitnessrUnRegister method allows the client to unregister for notifications from
	// the server. The Witness Service removes its internal state of the registration and
	// no longer notifies the client in the event of any resource state changes.
	//
	// Return Values: Returns 0x00000000 (ERROR_SUCCESS) on success or a nonzero error code,
	// as specified in [MS-ERREF] section 2.2. The most common error codes are listed in
	// the following table.
	//
	//	+--------------------------------+--------------------------------------------+
	//	|             RETURN             |                                            |
	//	|           VALUE/CODE           |                DESCRIPTION                 |
	//	|                                |                                            |
	//	+--------------------------------+--------------------------------------------+
	//	+--------------------------------+--------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The operation completed successfully.      |
	//	+--------------------------------+--------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied.                          |
	//	+--------------------------------+--------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND     | The specified CONTEXT_HANDLE is not found. |
	//	+--------------------------------+--------------------------------------------+
	Unregister(context.Context, *UnregisterRequest) (*UnregisterResponse, error)

	// The WitnessrAsyncNotify method is used by the client to request notification of registered
	// resource changes from the server.
	//
	// Return Values: Returns 0x00000000 (ERROR_SUCCESS) on success or a nonzero error code,
	// as specified in [MS-ERREF] section 2.2. The most common error codes are listed in
	// the following table.
	//
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	|                RETURN                |                                                                        |
	//	|              VALUE/CODE              |                              DESCRIPTION                               |
	//	|                                      |                                                                        |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS             | The operation completed successfully.                                  |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED       | Access is denied.                                                      |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x000005AA ERROR_NO_SYSTEM_RESOURCES | Insufficient system resources exist to complete the requested service. |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND           | The specified resource name is not found.                              |
	//	+--------------------------------------+------------------------------------------------------------------------+
	AsyncNotify(context.Context, *AsyncNotifyRequest) (*AsyncNotifyResponse, error)

	// The WitnessrRegisterEx method allows the witness client to register for resource
	// state change notifications of a NetName, ShareName and multiple IPAddresses. The
	// client can subsequently call the WitnessrAsyncNotify method to receive notifications
	// when there is a state change on any of these resources.
	//
	// Return Values: Returns 0x00000000 (ERROR_SUCCESS) on success or a nonzero error code,
	// as specified in [MS-ERREF] section 2.2. The most common error codes are listed in
	// the following table.
	//
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	|                RETURN                |                                                                        |
	//	|              VALUE/CODE              |                              DESCRIPTION                               |
	//	|                                      |                                                                        |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS             | The operation completed successfully.                                  |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED       | Access is denied.                                                      |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x000005AA ERROR_NO_SYSTEM_RESOURCES | Insufficient system resources exist to complete the requested service. |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER   | The parameter is incorrect.                                            |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x0000139F ERROR_INVALID_STATE       | The specified resource state is invalid.                               |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x0000051A ERROR_REVISION_MISMATCH   | The client request contains an invalid Witness protocol version.       |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//
	// This opnum is applicable only to servers that implement Witness protocol version
	// 2.
	//
	// If the Version field of the request is not 0x00020000, the server MUST stop processing
	// the request and return the error code ERROR_REVISION_MISMATCH.
	//
	// If NetName, IpAddress, or ClientComputerName is NULL, the server MUST fail the request
	// and return the error code ERROR_INVALID_PARAMETER.
	//
	// If the NetName parameter is not equal to ServerGlobalName, the server MUST fail the
	// request and return the error code ERROR_INVALID_PARAMETER.
	//
	// If ShareName is not NULL, the server MUST enumerate the shares by calling NetrShareEnum
	// as specified in [MS-SRVS] section 3.1.4.8. If the enumeration fails or if no shares
	// are returned, the server MUST return the error code ERROR_INVALID_STATE.
	//
	// If none of the shares in the list has shi*_type set to STYPE_CLUSTER_SOFS as specified
	// in [MS-SRVS] section 3.1.4.8, the server MUST ignore ShareName.
	//
	// Otherwise, the server MUST fail the request with the error code ERROR_INVALID_STATE
	// for the following:
	//
	// * *ShareName* does not exist in the enumerated list.
	//
	// * The server MUST search for an *Interface* in *InterfaceList* , where *Interface.IPv4Address*
	// or *Interface.IPv6Address* matches the IpAddress parameter based on its format. If
	// no matching entry is found and *ShareName* has shi*_type set to STYPE_CLUSTER_SOFS,
	// as specified in [MS-SRVS] section 2.2.2.4 ( ../ms-srvs/6069f8c0-c93f-43a0-a5b4-7ed447eb4b84
	// ) , the server MUST fail the request with ERROR_INVALID_STATE. **
	RegisterEx(context.Context, *RegisterExRequest) (*RegisterExResponse, error)
}

func RegisterWitnessServer(conn dcerpc.Conn, o WitnessServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWitnessServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WitnessSyntaxV1_1))...)
}

func NewWitnessServerHandle(o WitnessServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WitnessServerHandle(ctx, o, opNum, r)
	}
}

func WitnessServerHandle(ctx context.Context, o WitnessServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // WitnessrGetInterfaceList
		op := &xxx_GetInterfaceListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInterfaceListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInterfaceList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // WitnessrRegister
		op := &xxx_RegisterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Register(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // WitnessrUnRegister
		op := &xxx_UnregisterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UnregisterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Unregister(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // WitnessrAsyncNotify
		op := &xxx_AsyncNotifyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AsyncNotifyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AsyncNotify(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // WitnessrRegisterEx
		op := &xxx_RegisterExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented Witness
type UnimplementedWitnessServer struct {
}

func (UnimplementedWitnessServer) GetInterfaceList(context.Context, *GetInterfaceListRequest) (*GetInterfaceListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWitnessServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWitnessServer) Unregister(context.Context, *UnregisterRequest) (*UnregisterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWitnessServer) AsyncNotify(context.Context, *AsyncNotifyRequest) (*AsyncNotifyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWitnessServer) RegisterEx(context.Context, *RegisterExRequest) (*RegisterExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WitnessServer = (*UnimplementedWitnessServer)(nil)
