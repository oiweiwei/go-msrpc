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
type TransportServer interface {
	CheckConnectivity(context.Context, *CheckConnectivityRequest) (*CheckConnectivityResponse, error)

	EstablishConnection(context.Context, *EstablishConnectionRequest) (*EstablishConnectionResponse, error)

	EstablishSession(context.Context, *EstablishSessionRequest) (*EstablishSessionResponse, error)

	RequestUpdates(context.Context, *RequestUpdatesRequest) (*RequestUpdatesResponse, error)

	RequestVersionVector(context.Context, *RequestVersionVectorRequest) (*RequestVersionVectorResponse, error)

	AsyncPoll(context.Context, *AsyncPollRequest) (*AsyncPollResponse, error)

	RequestRecords(context.Context, *RequestRecordsRequest) (*RequestRecordsResponse, error)

	UpdateCancel(context.Context, *UpdateCancelRequest) (*UpdateCancelResponse, error)

	RawGetFileData(context.Context, *RawGetFileDataRequest) (*RawGetFileDataResponse, error)

	GetSignatures(context.Context, *GetSignaturesRequest) (*GetSignaturesResponse, error)

	PushSourceNeeds(context.Context, *PushSourceNeedsRequest) (*PushSourceNeedsResponse, error)

	GetFileData(context.Context, *GetFileDataRequest) (*GetFileDataResponse, error)

	Close(context.Context, *CloseRequest) (*CloseResponse, error)

	InitializeFileTransferAsync(context.Context, *InitializeFileTransferAsyncRequest) (*InitializeFileTransferAsyncResponse, error)

	// Opnum14NotUsedOnWire operation.
	// Opnum14NotUsedOnWire

	RawGetFileDataAsync(context.Context, *RawGetFileDataAsyncRequest) (*RawGetFileDataAsyncResponse, error)

	GetFileDataAsync(context.Context, *GetFileDataAsyncRequest) (*GetFileDataAsyncResponse, error)

	FileDataTransferKeepAlive(context.Context, *FileDataTransferKeepAliveRequest) (*FileDataTransferKeepAliveResponse, error)
}

func RegisterTransportServer(conn dcerpc.Conn, o TransportServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTransportServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TransportSyntaxV1_0))...)
}

func NewTransportServerHandle(o TransportServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TransportServerHandle(ctx, o, opNum, r)
	}
}

func TransportServerHandle(ctx context.Context, o TransportServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
		op := &xxx_GetSignaturesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSignaturesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSignatures(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RdcPushSourceNeeds
		op := &xxx_PushSourceNeedsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PushSourceNeedsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PushSourceNeeds(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RdcGetFileData
		op := &xxx_GetFileDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // RdcClose
		op := &xxx_CloseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Close(ctx, req)
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
		op := &xxx_GetFileDataAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileDataAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileDataAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // RdcFileDataTransferKeepAlive
		op := &xxx_FileDataTransferKeepAliveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FileDataTransferKeepAliveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FileDataTransferKeepAlive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented FrsTransport
type UnimplementedTransportServer struct {
}

func (UnimplementedTransportServer) CheckConnectivity(context.Context, *CheckConnectivityRequest) (*CheckConnectivityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) EstablishConnection(context.Context, *EstablishConnectionRequest) (*EstablishConnectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) EstablishSession(context.Context, *EstablishSessionRequest) (*EstablishSessionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) RequestUpdates(context.Context, *RequestUpdatesRequest) (*RequestUpdatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) RequestVersionVector(context.Context, *RequestVersionVectorRequest) (*RequestVersionVectorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) AsyncPoll(context.Context, *AsyncPollRequest) (*AsyncPollResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) RequestRecords(context.Context, *RequestRecordsRequest) (*RequestRecordsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) UpdateCancel(context.Context, *UpdateCancelRequest) (*UpdateCancelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) RawGetFileData(context.Context, *RawGetFileDataRequest) (*RawGetFileDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) GetSignatures(context.Context, *GetSignaturesRequest) (*GetSignaturesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) PushSourceNeeds(context.Context, *PushSourceNeedsRequest) (*PushSourceNeedsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) GetFileData(context.Context, *GetFileDataRequest) (*GetFileDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) Close(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) InitializeFileTransferAsync(context.Context, *InitializeFileTransferAsyncRequest) (*InitializeFileTransferAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) RawGetFileDataAsync(context.Context, *RawGetFileDataAsyncRequest) (*RawGetFileDataAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) GetFileDataAsync(context.Context, *GetFileDataAsyncRequest) (*GetFileDataAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) FileDataTransferKeepAlive(context.Context, *FileDataTransferKeepAliveRequest) (*FileDataTransferKeepAliveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TransportServer = (*UnimplementedTransportServer)(nil)
