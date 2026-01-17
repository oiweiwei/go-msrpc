package igettrackingdata

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = iunknown.GoPackage
)

// IGetTrackingData server interface.
type GetTrackingDataServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// A client calls this method to obtain tracking information for instance containers
	// across all conglomerations.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result (as specified in [MS-ERREF] section 2.1) on failure.
	GetContainerData(context.Context, *GetContainerDataRequest) (*GetContainerDataResponse, error)

	// A client calls this method to obtain tracking information for components that have
	// one or more component instances in a given instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result (as specified in [MS-ERREF] section 2.1) on failure.
	GetComponentDataByContainer(context.Context, *GetComponentDataByContainerRequest) (*GetComponentDataByContainerResponse, error)

	// A client calls this method to obtain tracking information for a single component
	// that has component instances in an instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result (as specified in [MS-ERREF] section 2.1) on failure.
	GetComponentDataByContainerAndClassID(context.Context, *GetComponentDataByContainerAndClassIDRequest) (*GetComponentDataByContainerAndClassIDResponse, error)

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire
}

func RegisterGetTrackingDataServer(conn dcerpc.Conn, o GetTrackingDataServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewGetTrackingDataServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(GetTrackingDataSyntaxV0_0))...)
}

func NewGetTrackingDataServerHandle(o GetTrackingDataServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return GetTrackingDataServerHandle(ctx, o, opNum, r)
	}
}

func GetTrackingDataServerHandle(ctx context.Context, o GetTrackingDataServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Opnum3NotUsedOnWire
		// Opnum3NotUsedOnWire
		return nil, nil
	case 4: // GetContainerData
		op := &xxx_GetContainerDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetContainerDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetContainerData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetComponentDataByContainer
		op := &xxx_GetComponentDataByContainerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetComponentDataByContainerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetComponentDataByContainer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // GetComponentDataByContainerAndCLSID
		op := &xxx_GetComponentDataByContainerAndClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetComponentDataByContainerAndClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetComponentDataByContainerAndClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IGetTrackingData
type UnimplementedGetTrackingDataServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedGetTrackingDataServer) GetContainerData(context.Context, *GetContainerDataRequest) (*GetContainerDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedGetTrackingDataServer) GetComponentDataByContainer(context.Context, *GetComponentDataByContainerRequest) (*GetComponentDataByContainerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedGetTrackingDataServer) GetComponentDataByContainerAndClassID(context.Context, *GetComponentDataByContainerAndClassIDRequest) (*GetComponentDataByContainerAndClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ GetTrackingDataServer = (*UnimplementedGetTrackingDataServer)(nil)
