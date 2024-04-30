package dscomm2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
)

// dscomm2 server interface.
type Dscomm2Server interface {

	// The S_DSGetComputerSites method returns the site identifier for every site of which
	// the specified computer is a member.
	//
	// Return Values:  If the method succeeds, the return value is MQ_OK (0x00000000).
	// If the method fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetComputerSites(context.Context, *GetComputerSitesRequest) (*GetComputerSitesResponse, error)

	// The S_DSGetPropsEx method returns the properties associated with the object specified
	// by a directory service pathname. This method differs from S_DSGetProps (section 3.1.4.7)
	// in that it supports a restricted set of properties that pertain only to queue or
	// machine object security.
	//
	// Return Values:  If the method succeeds, the return value is MQ_OK (0x00000000).
	// If the method fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetPropertiesEx(context.Context, *GetPropertiesExRequest) (*GetPropertiesExResponse, error)

	// This method returns the properties for the object specified by object identifier.
	// This method differs from S_DSGetPropsGuid in that it supports a restricted set of
	// properties that pertain only to queue or machine object security.
	//
	// Return Values:  If the method succeeds, the return value is MQ_OK (0x00000000).
	// If the method fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	GetPropertiesGUIDEx(context.Context, *GetPropertiesGUIDExRequest) (*GetPropertiesGUIDExResponse, error)

	// The S_DSBeginDeleteNotification method begins a delete notification and returns an
	// RPC context handle associated with the delete notification.
	//
	// Return Values:  If the method succeeds, the return value is MQ_OK (0x00000000).
	// If the method fails, the return value is an implementation-specific error code.<158>
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	BeginDeleteNotification(context.Context, *BeginDeleteNotificationRequest) (*BeginDeleteNotificationResponse, error)

	// This method instructs the server to notify the computer that owns the deleted object
	// about the deletion.
	//
	// Return Values:  If the method succeeds, the return value is MQ_OK (0x00000000).
	// If the method fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	NotifyDelete(context.Context, *NotifyDeleteRequest) (*NotifyDeleteResponse, error)

	// The S_DSEndDeleteNotification method closes the RPC context handle acquired from
	// a previous call to S_DSBeginDeleteNotification.
	//
	// Return Values:  None.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	EndDeleteNotification(context.Context, *EndDeleteNotificationRequest) (*EndDeleteNotificationResponse, error)

	// This method returns a value that indicates if that server is a Global Catalog Server.
	//
	// Return Values:  This method returns TRUE (0x00000001) if the Directory Service server
	// is also a Global Catalog Server; otherwise, it returns FALSE (0x00000000).
	//
	// TRUE (0x00000001)
	//
	// FALSE (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	IsServerGC(context.Context, *IsServerGCRequest) (*IsServerGCResponse, error)

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// The S_DSGetGCListInDomain method returns the list of Global Catalog Servers in the
	// specified domain.
	//
	// Return Values:  If the method succeeds, the return value is MQ_OK (0x00000000).
	// If the method fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, Remote Procedure Call Protocol Extensions, as specified in [MS-RPCE].
	GetGCListInDomain(context.Context, *GetGCListInDomainRequest) (*GetGCListInDomainResponse, error)
}

func RegisterDscomm2Server(conn dcerpc.Conn, o Dscomm2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDscomm2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Dscomm2SyntaxV1_0))...)
}

func NewDscomm2ServerHandle(o Dscomm2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Dscomm2ServerHandle(ctx, o, opNum, r)
	}
}

func Dscomm2ServerHandle(ctx context.Context, o Dscomm2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // S_DSGetComputerSites
		in := &GetComputerSitesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetComputerSites(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // S_DSGetPropsEx
		in := &GetPropertiesExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertiesEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // S_DSGetPropsGuidEx
		in := &GetPropertiesGUIDExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertiesGUIDEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // S_DSBeginDeleteNotification
		in := &BeginDeleteNotificationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BeginDeleteNotification(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // S_DSNotifyDelete
		in := &NotifyDeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NotifyDelete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // S_DSEndDeleteNotification
		in := &EndDeleteNotificationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EndDeleteNotification(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // S_DSIsServerGC
		in := &IsServerGCRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsServerGC(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	case 8: // S_DSGetGCListInDomain
		in := &GetGCListInDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetGCListInDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
