package iadproxy

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

// IADProxy server interface.
type IADProxyServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The CreateObject method MUST execute an LDAP command under machine security credentials,
	// or for a cluster, under the specified network name credentials in order to create
	// an Active Directory object that has a specific distinguished name and attributes.<52>
	//
	// Return Values: The method MUST return:
	//
	// * 0 on success.
	//
	// * For LDAP protocol failures:
	//
	// * If the LDAP error is LDAP_OPERATIONS_ERROR, dfsrHelperLdapErrorBase + the server-side
	// error code.
	//
	// * For all other LDAP errors, dfsrHelperLdapErrorBase + the LDAP return code. For
	// more information, see [LDAP-ERR] ( https://go.microsoft.com/fwlink/?LinkId=89933
	// ).
	//
	// * For all other failures, an implementation-specific nonzero HRESULT ( ../ms-dtyp/a9046ed2-bfb2-4d56-a719-2824afce59ac
	// ) error code, as specified in [MS-ERREF] ( ../ms-erref/1bc92ddf-b79e-413c-bbaa-99a5281a6c90
	// ) section 2.1 ( ../ms-erref/0642cb2f-2075-4469-918c-4441e69c548a ) , between 0x80000000
	// and 0xFFFFFFFF. For protocol purposes, all nonzero values MUST be treated as equivalent
	// failures.
	CreateObject(context.Context, *CreateObjectRequest) (*CreateObjectResponse, error)

	// The DeleteObject method MUST execute an LDAP command under machine security credentials
	// to delete an Active Directory object with a specified distinguished name.<46>
	//
	// The DeleteObject method executes an LDAP command to delete an Active Directory object
	// that has a specified distinguished name and attributes. The command MUST be executed
	// under the machine security credentials, or for a cluster, under the specified network
	// name credentials.<54>
	//
	// Return Values: The method MUST return:
	//
	// * A value of 0 when:
	//
	// * The method call is successful.
	//
	// * The LDAP error is LDAP_NO_SUCH_OBJECT.
	//
	// * For other LDAP protocol failures:
	//
	// * A value of dfsrHelperLdapErrorBase + the server-side error code  if the LDAP error
	// is LDAP_OPERATIONS_ERROR.
	//
	// * A value of dfsrHelperLdapErrorBase + the LDAP return code for all other LDAP errors.
	// For more information, see [LDAP-ERR] ( https://go.microsoft.com/fwlink/?LinkId=89933
	// ).
	//
	// * For all other failures, an implementation-specific nonzero HRESULT ( ../ms-dtyp/a9046ed2-bfb2-4d56-a719-2824afce59ac
	// ) error code, as specified in [MS-ERREF] ( ../ms-erref/1bc92ddf-b79e-413c-bbaa-99a5281a6c90
	// ) section 2.1 ( ../ms-erref/0642cb2f-2075-4469-918c-4441e69c548a ) , between 0x80000000
	// and 0xFFFFFFFF. For protocol purposes, all nonzero values MUST be treated as equivalent
	// failures.
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)

	// The ModifyObject method executes an LDAP command to add, delete, or modify attributes
	// of a specified Active Directory object. The command MUST be executed under machine
	// security credentials, or for a cluster, under the specified network name credentials
	// in order to modify an Active Directory object that has a specific distinguished name
	// and attributes.<56>
	//
	// The ModifyObject method MUST execute an LDAP command under machine security credentials
	// to add, delete, or modify attributes of an Active Directory object that has a specified
	// distinguished name.<48>
	//
	// Return Values: The method MUST return:
	//
	// * Zero on success.
	//
	// * For LDAP protocol failures:
	//
	// * If the LDAP error is LDAP_OPERATIONS_ERROR, dfsrHelperLdapErrorBase + the server-side
	// error code.
	//
	// * For all other LDAP errors, dfsrHelperLdapErrorBase + the LDAP return code. For
	// more information, see [LDAP-ERR] ( https://go.microsoft.com/fwlink/?LinkId=89933
	// ).
	//
	// * For all other failures, an implementation-specific nonzero HRESULT ( ../ms-dtyp/a9046ed2-bfb2-4d56-a719-2824afce59ac
	// ) error code, as specified in [MS-ERREF] ( ../ms-erref/1bc92ddf-b79e-413c-bbaa-99a5281a6c90
	// ) section 2.1 ( ../ms-erref/0642cb2f-2075-4469-918c-4441e69c548a ) , between 0x80000000
	// and 0xFFFFFFFF. For protocol purposes, all nonzero values MUST be treated as equivalent
	// failures.
	ModifyObject(context.Context, *ModifyObjectRequest) (*ModifyObjectResponse, error)
}

func RegisterIADProxyServer(conn dcerpc.Conn, o IADProxyServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIADProxyServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IADProxySyntaxV0_0))...)
}

func NewIADProxyServerHandle(o IADProxyServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IADProxyServerHandle(ctx, o, opNum, r)
	}
}

func IADProxyServerHandle(ctx context.Context, o IADProxyServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CreateObject
		op := &xxx_CreateObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // DeleteObject
		op := &xxx_DeleteObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // ModifyObject
		op := &xxx_ModifyObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IADProxy
type UnimplementedIADProxyServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedIADProxyServer) CreateObject(context.Context, *CreateObjectRequest) (*CreateObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIADProxyServer) DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIADProxyServer) ModifyObject(context.Context, *ModifyObjectRequest) (*ModifyObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ IADProxyServer = (*UnimplementedIADProxyServer)(nil)
