package efsrpc

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

// efsrpc server interface.
type EfsrpcServer interface {
	OpenFileRaw(context.Context, *OpenFileRawRequest) (*OpenFileRawResponse, error)

	ReadFileRaw(context.Context, *ReadFileRawRequest) (*ReadFileRawResponse, error)

	WriteFileRaw(context.Context, *WriteFileRawRequest) (*WriteFileRawResponse, error)

	CloseRaw(context.Context, *CloseRawRequest) (*CloseRawResponse, error)

	// EfsRpcEncryptFileSrv operation.
	EncryptFileServer(context.Context, *EncryptFileServerRequest) (*EncryptFileServerResponse, error)

	DecryptFileServer(context.Context, *DecryptFileServerRequest) (*DecryptFileServerResponse, error)

	// EfsRpcQueryUsersOnFile operation.
	QueryUsersOnFile(context.Context, *QueryUsersOnFileRequest) (*QueryUsersOnFileResponse, error)

	QueryRecoveryAgents(context.Context, *QueryRecoveryAgentsRequest) (*QueryRecoveryAgentsResponse, error)

	RemoveUsersFromFile(context.Context, *RemoveUsersFromFileRequest) (*RemoveUsersFromFileResponse, error)

	AddUsersToFile(context.Context, *AddUsersToFileRequest) (*AddUsersToFileResponse, error)

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	NotSupported(context.Context, *NotSupportedRequest) (*NotSupportedResponse, error)

	FileKeyInfo(context.Context, *FileKeyInfoRequest) (*FileKeyInfoResponse, error)

	DuplicateEncryptionInfoFile(context.Context, *DuplicateEncryptionInfoFileRequest) (*DuplicateEncryptionInfoFileResponse, error)

	// Opnum14NotUsedOnWire operation.
	// Opnum14NotUsedOnWire

	AddUsersToFileEx(context.Context, *AddUsersToFileExRequest) (*AddUsersToFileExResponse, error)

	FileKeyInfoEx(context.Context, *FileKeyInfoExRequest) (*FileKeyInfoExResponse, error)

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire

	GetEncryptedFileMetadata(context.Context, *GetEncryptedFileMetadataRequest) (*GetEncryptedFileMetadataResponse, error)

	SetEncryptedFileMetadata(context.Context, *SetEncryptedFileMetadataRequest) (*SetEncryptedFileMetadataResponse, error)

	FlushEFSCache(context.Context, *FlushEFSCacheRequest) (*FlushEFSCacheResponse, error)

	// EfsRpcEncryptFileExSrv operation.
	EncryptFileExServer(context.Context, *EncryptFileExServerRequest) (*EncryptFileExServerResponse, error)

	// EfsRpcQueryProtectors operation.
	QueryProtectors(context.Context, *QueryProtectorsRequest) (*QueryProtectorsResponse, error)

	// Opnum23NotUsedOnWire operation.
	// Opnum23NotUsedOnWire

	// Opnum24NotUsedOnWire operation.
	// Opnum24NotUsedOnWire

	// Opnum25NotUsedOnWire operation.
	// Opnum25NotUsedOnWire

	// Opnum26NotUsedOnWire operation.
	// Opnum26NotUsedOnWire

	// Opnum27NotUsedOnWire operation.
	// Opnum27NotUsedOnWire

	// Opnum28NotUsedOnWire operation.
	// Opnum28NotUsedOnWire

	// Opnum29NotUsedOnWire operation.
	// Opnum29NotUsedOnWire

	// Opnum30NotUsedOnWire operation.
	// Opnum30NotUsedOnWire

	// Opnum31NotUsedOnWire operation.
	// Opnum31NotUsedOnWire

	// Opnum32NotUsedOnWire operation.
	// Opnum32NotUsedOnWire

	// Opnum33NotUsedOnWire operation.
	// Opnum33NotUsedOnWire

	// Opnum34NotUsedOnWire operation.
	// Opnum34NotUsedOnWire

	// Opnum35NotUsedOnWire operation.
	// Opnum35NotUsedOnWire

	// Opnum36NotUsedOnWire operation.
	// Opnum36NotUsedOnWire

	// Opnum37NotUsedOnWire operation.
	// Opnum37NotUsedOnWire

	// Opnum38NotUsedOnWire operation.
	// Opnum38NotUsedOnWire

	// Opnum39NotUsedOnWire operation.
	// Opnum39NotUsedOnWire

	// Opnum40NotUsedOnWire operation.
	// Opnum40NotUsedOnWire

	// Opnum41NotUsedOnWire operation.
	// Opnum41NotUsedOnWire

	// Opnum42NotUsedOnWire operation.
	// Opnum42NotUsedOnWire

	// Opnum43NotUsedOnWire operation.
	// Opnum43NotUsedOnWire

	// Opnum44NotUsedOnWire operation.
	// Opnum44NotUsedOnWire
}

func RegisterEfsrpcServer(conn dcerpc.Conn, o EfsrpcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEfsrpcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EfsrpcSyntaxV1_0))...)
}

func NewEfsrpcServerHandle(o EfsrpcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EfsrpcServerHandle(ctx, o, opNum, r)
	}
}

func EfsrpcServerHandle(ctx context.Context, o EfsrpcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // EfsRpcOpenFileRaw
		op := &xxx_OpenFileRawOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenFileRawRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenFileRaw(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // EfsRpcReadFileRaw
		op := &xxx_ReadFileRawOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReadFileRawRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReadFileRaw(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // EfsRpcWriteFileRaw
		op := &xxx_WriteFileRawOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WriteFileRawRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WriteFileRaw(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // EfsRpcCloseRaw
		op := &xxx_CloseRawOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseRawRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseRaw(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // EfsRpcEncryptFileSrv
		op := &xxx_EncryptFileServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EncryptFileServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EncryptFileServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // EfsRpcDecryptFileSrv
		op := &xxx_DecryptFileServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DecryptFileServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DecryptFileServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // EfsRpcQueryUsersOnFile
		op := &xxx_QueryUsersOnFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryUsersOnFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryUsersOnFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // EfsRpcQueryRecoveryAgents
		op := &xxx_QueryRecoveryAgentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryRecoveryAgentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryRecoveryAgents(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // EfsRpcRemoveUsersFromFile
		op := &xxx_RemoveUsersFromFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveUsersFromFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveUsersFromFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // EfsRpcAddUsersToFile
		op := &xxx_AddUsersToFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddUsersToFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddUsersToFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Opnum10NotUsedOnWire
		// Opnum10NotUsedOnWire
		return nil, nil
	case 11: // EfsRpcNotSupported
		op := &xxx_NotSupportedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NotSupportedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NotSupported(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // EfsRpcFileKeyInfo
		op := &xxx_FileKeyInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FileKeyInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FileKeyInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // EfsRpcDuplicateEncryptionInfoFile
		op := &xxx_DuplicateEncryptionInfoFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DuplicateEncryptionInfoFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DuplicateEncryptionInfoFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Opnum14NotUsedOnWire
		// Opnum14NotUsedOnWire
		return nil, nil
	case 15: // EfsRpcAddUsersToFileEx
		op := &xxx_AddUsersToFileExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddUsersToFileExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddUsersToFileEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // EfsRpcFileKeyInfoEx
		op := &xxx_FileKeyInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FileKeyInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FileKeyInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Opnum17NotUsedOnWire
		// Opnum17NotUsedOnWire
		return nil, nil
	case 18: // EfsRpcGetEncryptedFileMetadata
		op := &xxx_GetEncryptedFileMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEncryptedFileMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEncryptedFileMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // EfsRpcSetEncryptedFileMetadata
		op := &xxx_SetEncryptedFileMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEncryptedFileMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEncryptedFileMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // EfsRpcFlushEfsCache
		op := &xxx_FlushEFSCacheOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FlushEFSCacheRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FlushEFSCache(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // EfsRpcEncryptFileExSrv
		op := &xxx_EncryptFileExServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EncryptFileExServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EncryptFileExServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // EfsRpcQueryProtectors
		op := &xxx_QueryProtectorsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryProtectorsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryProtectors(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // Opnum23NotUsedOnWire
		// Opnum23NotUsedOnWire
		return nil, nil
	case 24: // Opnum24NotUsedOnWire
		// Opnum24NotUsedOnWire
		return nil, nil
	case 25: // Opnum25NotUsedOnWire
		// Opnum25NotUsedOnWire
		return nil, nil
	case 26: // Opnum26NotUsedOnWire
		// Opnum26NotUsedOnWire
		return nil, nil
	case 27: // Opnum27NotUsedOnWire
		// Opnum27NotUsedOnWire
		return nil, nil
	case 28: // Opnum28NotUsedOnWire
		// Opnum28NotUsedOnWire
		return nil, nil
	case 29: // Opnum29NotUsedOnWire
		// Opnum29NotUsedOnWire
		return nil, nil
	case 30: // Opnum30NotUsedOnWire
		// Opnum30NotUsedOnWire
		return nil, nil
	case 31: // Opnum31NotUsedOnWire
		// Opnum31NotUsedOnWire
		return nil, nil
	case 32: // Opnum32NotUsedOnWire
		// Opnum32NotUsedOnWire
		return nil, nil
	case 33: // Opnum33NotUsedOnWire
		// Opnum33NotUsedOnWire
		return nil, nil
	case 34: // Opnum34NotUsedOnWire
		// Opnum34NotUsedOnWire
		return nil, nil
	case 35: // Opnum35NotUsedOnWire
		// Opnum35NotUsedOnWire
		return nil, nil
	case 36: // Opnum36NotUsedOnWire
		// Opnum36NotUsedOnWire
		return nil, nil
	case 37: // Opnum37NotUsedOnWire
		// Opnum37NotUsedOnWire
		return nil, nil
	case 38: // Opnum38NotUsedOnWire
		// Opnum38NotUsedOnWire
		return nil, nil
	case 39: // Opnum39NotUsedOnWire
		// Opnum39NotUsedOnWire
		return nil, nil
	case 40: // Opnum40NotUsedOnWire
		// Opnum40NotUsedOnWire
		return nil, nil
	case 41: // Opnum41NotUsedOnWire
		// Opnum41NotUsedOnWire
		return nil, nil
	case 42: // Opnum42NotUsedOnWire
		// Opnum42NotUsedOnWire
		return nil, nil
	case 43: // Opnum43NotUsedOnWire
		// Opnum43NotUsedOnWire
		return nil, nil
	case 44: // Opnum44NotUsedOnWire
		// Opnum44NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented efsrpc
type UnimplementedEfsrpcServer struct {
}

func (UnimplementedEfsrpcServer) OpenFileRaw(context.Context, *OpenFileRawRequest) (*OpenFileRawResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) ReadFileRaw(context.Context, *ReadFileRawRequest) (*ReadFileRawResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) WriteFileRaw(context.Context, *WriteFileRawRequest) (*WriteFileRawResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) CloseRaw(context.Context, *CloseRawRequest) (*CloseRawResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) EncryptFileServer(context.Context, *EncryptFileServerRequest) (*EncryptFileServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) DecryptFileServer(context.Context, *DecryptFileServerRequest) (*DecryptFileServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) QueryUsersOnFile(context.Context, *QueryUsersOnFileRequest) (*QueryUsersOnFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) QueryRecoveryAgents(context.Context, *QueryRecoveryAgentsRequest) (*QueryRecoveryAgentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) RemoveUsersFromFile(context.Context, *RemoveUsersFromFileRequest) (*RemoveUsersFromFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) AddUsersToFile(context.Context, *AddUsersToFileRequest) (*AddUsersToFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) NotSupported(context.Context, *NotSupportedRequest) (*NotSupportedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) FileKeyInfo(context.Context, *FileKeyInfoRequest) (*FileKeyInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) DuplicateEncryptionInfoFile(context.Context, *DuplicateEncryptionInfoFileRequest) (*DuplicateEncryptionInfoFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) AddUsersToFileEx(context.Context, *AddUsersToFileExRequest) (*AddUsersToFileExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) FileKeyInfoEx(context.Context, *FileKeyInfoExRequest) (*FileKeyInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) GetEncryptedFileMetadata(context.Context, *GetEncryptedFileMetadataRequest) (*GetEncryptedFileMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) SetEncryptedFileMetadata(context.Context, *SetEncryptedFileMetadataRequest) (*SetEncryptedFileMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) FlushEFSCache(context.Context, *FlushEFSCacheRequest) (*FlushEFSCacheResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) EncryptFileExServer(context.Context, *EncryptFileExServerRequest) (*EncryptFileExServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) QueryProtectors(context.Context, *QueryProtectorsRequest) (*QueryProtectorsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EfsrpcServer = (*UnimplementedEfsrpcServer)(nil)
