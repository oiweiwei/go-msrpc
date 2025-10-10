package iwrmaccounting

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

// IWRMAccounting server interface.
type AccountingServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	CreateAccountingDB(context.Context, *CreateAccountingDBRequest) (*CreateAccountingDBResponse, error)

	GetAccountingMetadata(context.Context, *GetAccountingMetadataRequest) (*GetAccountingMetadataResponse, error)

	ExecuteAccountingQuery(context.Context, *ExecuteAccountingQueryRequest) (*ExecuteAccountingQueryResponse, error)

	GetRawAccountingData(context.Context, *GetRawAccountingDataRequest) (*GetRawAccountingDataResponse, error)

	GetNextAccountingDataBatch(context.Context, *GetNextAccountingDataBatchRequest) (*GetNextAccountingDataBatchResponse, error)

	DeleteAccountingData(context.Context, *DeleteAccountingDataRequest) (*DeleteAccountingDataResponse, error)

	DefragmentDB(context.Context, *DefragmentDBRequest) (*DefragmentDBResponse, error)

	CancelAccountingQuery(context.Context, *CancelAccountingQueryRequest) (*CancelAccountingQueryResponse, error)

	RegisterAccountingClient(context.Context, *RegisterAccountingClientRequest) (*RegisterAccountingClientResponse, error)

	DumpAccountingData(context.Context, *DumpAccountingDataRequest) (*DumpAccountingDataResponse, error)

	GetAccountingClients(context.Context, *GetAccountingClientsRequest) (*GetAccountingClientsResponse, error)

	SetAccountingClientStatus(context.Context, *SetAccountingClientStatusRequest) (*SetAccountingClientStatusResponse, error)

	CheckAccountingConnection(context.Context, *CheckAccountingConnectionRequest) (*CheckAccountingConnectionResponse, error)

	SetClientPermissions(context.Context, *SetClientPermissionsRequest) (*SetClientPermissionsResponse, error)
}

func RegisterAccountingServer(conn dcerpc.Conn, o AccountingServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAccountingServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AccountingSyntaxV0_0))...)
}

func NewAccountingServerHandle(o AccountingServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AccountingServerHandle(ctx, o, opNum, r)
	}
}

func AccountingServerHandle(ctx context.Context, o AccountingServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CreateAccountingDb
		op := &xxx_CreateAccountingDBOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateAccountingDBRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateAccountingDB(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetAccountingMetadata
		op := &xxx_GetAccountingMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAccountingMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAccountingMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ExecuteAccountingQuery
		op := &xxx_ExecuteAccountingQueryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecuteAccountingQueryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExecuteAccountingQuery(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetRawAccountingData
		op := &xxx_GetRawAccountingDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRawAccountingDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRawAccountingData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // GetNextAccountingDataBatch
		op := &xxx_GetNextAccountingDataBatchOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNextAccountingDataBatchRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNextAccountingDataBatch(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // DeleteAccountingData
		op := &xxx_DeleteAccountingDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAccountingDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAccountingData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // DefragmentDB
		op := &xxx_DefragmentDBOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DefragmentDBRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DefragmentDB(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // CancelAccountingQuery
		op := &xxx_CancelAccountingQueryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CancelAccountingQueryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CancelAccountingQuery(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // RegisterAccountingClient
		op := &xxx_RegisterAccountingClientOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterAccountingClientRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterAccountingClient(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // DumpAccountingData
		op := &xxx_DumpAccountingDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DumpAccountingDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DumpAccountingData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // GetAccountingClients
		op := &xxx_GetAccountingClientsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAccountingClientsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAccountingClients(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // SetAccountingClientStatus
		op := &xxx_SetAccountingClientStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAccountingClientStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAccountingClientStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // CheckAccountingConnection
		op := &xxx_CheckAccountingConnectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CheckAccountingConnectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CheckAccountingConnection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // SetClientPermissions
		op := &xxx_SetClientPermissionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClientPermissionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClientPermissions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMAccounting
type UnimplementedAccountingServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedAccountingServer) CreateAccountingDB(context.Context, *CreateAccountingDBRequest) (*CreateAccountingDBResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) GetAccountingMetadata(context.Context, *GetAccountingMetadataRequest) (*GetAccountingMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) ExecuteAccountingQuery(context.Context, *ExecuteAccountingQueryRequest) (*ExecuteAccountingQueryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) GetRawAccountingData(context.Context, *GetRawAccountingDataRequest) (*GetRawAccountingDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) GetNextAccountingDataBatch(context.Context, *GetNextAccountingDataBatchRequest) (*GetNextAccountingDataBatchResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) DeleteAccountingData(context.Context, *DeleteAccountingDataRequest) (*DeleteAccountingDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) DefragmentDB(context.Context, *DefragmentDBRequest) (*DefragmentDBResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) CancelAccountingQuery(context.Context, *CancelAccountingQueryRequest) (*CancelAccountingQueryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) RegisterAccountingClient(context.Context, *RegisterAccountingClientRequest) (*RegisterAccountingClientResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) DumpAccountingData(context.Context, *DumpAccountingDataRequest) (*DumpAccountingDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) GetAccountingClients(context.Context, *GetAccountingClientsRequest) (*GetAccountingClientsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) SetAccountingClientStatus(context.Context, *SetAccountingClientStatusRequest) (*SetAccountingClientStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) CheckAccountingConnection(context.Context, *CheckAccountingConnectionRequest) (*CheckAccountingConnectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) SetClientPermissions(context.Context, *SetClientPermissionsRequest) (*SetClientPermissionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AccountingServer = (*UnimplementedAccountingServer)(nil)
