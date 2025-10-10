package frstransport

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

// FrsTransport server interface.
type FrsTransportServer interface {
	CheckConnectivity(context.Context, *CheckConnectivityRequest) (*CheckConnectivityResponse, error)

	EstablishConnection(context.Context, *EstablishConnectionRequest) (*EstablishConnectionResponse, error)

	EstablishSession(context.Context, *EstablishSessionRequest) (*EstablishSessionResponse, error)

	RequestUpdates(context.Context, *RequestUpdatesRequest) (*RequestUpdatesResponse, error)

	RequestVersionVector(context.Context, *RequestVersionVectorRequest) (*RequestVersionVectorResponse, error)

	AsyncPoll(context.Context, *AsyncPollRequest) (*AsyncPollResponse, error)

	RequestRecords(context.Context, *RequestRecordsRequest) (*RequestRecordsResponse, error)

	UpdateCancel(context.Context, *UpdateCancelRequest) (*UpdateCancelResponse, error)

	RawGetFileData(context.Context, *RawGetFileDataRequest) (*RawGetFileDataResponse, error)

	RdcGetSignatures(context.Context, *RdcGetSignaturesRequest) (*RdcGetSignaturesResponse, error)

	RdcPushSourceNeeds(context.Context, *RdcPushSourceNeedsRequest) (*RdcPushSourceNeedsResponse, error)

	RdcGetFileData(context.Context, *RdcGetFileDataRequest) (*RdcGetFileDataResponse, error)

	RdcClose(context.Context, *RdcCloseRequest) (*RdcCloseResponse, error)

	InitializeFileTransferAsync(context.Context, *InitializeFileTransferAsyncRequest) (*InitializeFileTransferAsyncResponse, error)

	// Opnum14NotUsedOnWire operation.
	// Opnum14NotUsedOnWire

	RawGetFileDataAsync(context.Context, *RawGetFileDataAsyncRequest) (*RawGetFileDataAsyncResponse, error)

	RdcGetFileDataAsync(context.Context, *RdcGetFileDataAsyncRequest) (*RdcGetFileDataAsyncResponse, error)

	RdcFileDataTransferKeepAlive(context.Context, *RdcFileDataTransferKeepAliveRequest) (*RdcFileDataTransferKeepAliveResponse, error)
}

func RegisterFrsTransportServer(conn dcerpc.Conn, o FrsTransportServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFrsTransportServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FrsTransportSyntaxV1_0))...)
}

func NewFrsTransportServerHandle(o FrsTransportServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FrsTransportServerHandle(ctx, o, opNum, r)
	}
}

func FrsTransportServerHandle(ctx context.Context, o FrsTransportServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // CheckConnectivity
		op := &xxx_CheckConnectivityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CheckConnectivityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CheckConnectivity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // EstablishConnection
		op := &xxx_EstablishConnectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EstablishConnectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EstablishConnection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // EstablishSession
		op := &xxx_EstablishSessionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EstablishSessionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EstablishSession(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // RequestUpdates
		op := &xxx_RequestUpdatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RequestUpdatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RequestUpdates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RequestVersionVector
		op := &xxx_RequestVersionVectorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RequestVersionVectorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RequestVersionVector(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // AsyncPoll
		op := &xxx_AsyncPollOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AsyncPollRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AsyncPoll(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // RequestRecords
		op := &xxx_RequestRecordsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RequestRecordsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RequestRecords(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // UpdateCancel
		op := &xxx_UpdateCancelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UpdateCancelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UpdateCancel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // RawGetFileData
		op := &xxx_RawGetFileDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RawGetFileDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RawGetFileData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // RdcGetSignatures
		op := &xxx_RdcGetSignaturesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RdcGetSignaturesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RdcGetSignatures(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RdcPushSourceNeeds
		op := &xxx_RdcPushSourceNeedsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RdcPushSourceNeedsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RdcPushSourceNeeds(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RdcGetFileData
		op := &xxx_RdcGetFileDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RdcGetFileDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RdcGetFileData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // RdcClose
		op := &xxx_RdcCloseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RdcCloseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RdcClose(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // InitializeFileTransferAsync
		op := &xxx_InitializeFileTransferAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitializeFileTransferAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InitializeFileTransferAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Opnum14NotUsedOnWire
		// Opnum14NotUsedOnWire
		return nil, nil
	case 15: // RawGetFileDataAsync
		op := &xxx_RawGetFileDataAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RawGetFileDataAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RawGetFileDataAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // RdcGetFileDataAsync
		op := &xxx_RdcGetFileDataAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RdcGetFileDataAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RdcGetFileDataAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // RdcFileDataTransferKeepAlive
		op := &xxx_RdcFileDataTransferKeepAliveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RdcFileDataTransferKeepAliveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RdcFileDataTransferKeepAlive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented FrsTransport
type UnimplementedFrsTransportServer struct {
}

func (UnimplementedFrsTransportServer) CheckConnectivity(context.Context, *CheckConnectivityRequest) (*CheckConnectivityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) EstablishConnection(context.Context, *EstablishConnectionRequest) (*EstablishConnectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) EstablishSession(context.Context, *EstablishSessionRequest) (*EstablishSessionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) RequestUpdates(context.Context, *RequestUpdatesRequest) (*RequestUpdatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) RequestVersionVector(context.Context, *RequestVersionVectorRequest) (*RequestVersionVectorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) AsyncPoll(context.Context, *AsyncPollRequest) (*AsyncPollResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) RequestRecords(context.Context, *RequestRecordsRequest) (*RequestRecordsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) UpdateCancel(context.Context, *UpdateCancelRequest) (*UpdateCancelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) RawGetFileData(context.Context, *RawGetFileDataRequest) (*RawGetFileDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) RdcGetSignatures(context.Context, *RdcGetSignaturesRequest) (*RdcGetSignaturesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) RdcPushSourceNeeds(context.Context, *RdcPushSourceNeedsRequest) (*RdcPushSourceNeedsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) RdcGetFileData(context.Context, *RdcGetFileDataRequest) (*RdcGetFileDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) RdcClose(context.Context, *RdcCloseRequest) (*RdcCloseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) InitializeFileTransferAsync(context.Context, *InitializeFileTransferAsyncRequest) (*InitializeFileTransferAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) RawGetFileDataAsync(context.Context, *RawGetFileDataAsyncRequest) (*RawGetFileDataAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) RdcGetFileDataAsync(context.Context, *RdcGetFileDataAsyncRequest) (*RdcGetFileDataAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsTransportServer) RdcFileDataTransferKeepAlive(context.Context, *RdcFileDataTransferKeepAliveRequest) (*RdcFileDataTransferKeepAliveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FrsTransportServer = (*UnimplementedFrsTransportServer)(nil)
