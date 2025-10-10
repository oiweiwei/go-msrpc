package intmslibrarycontrol1

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

// INtmsLibraryControl1 server interface.
type NTMSLibraryControl1Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	EjectNTMSMedia(context.Context, *EjectNTMSMediaRequest) (*EjectNTMSMediaResponse, error)

	InjectNTMSMedia(context.Context, *InjectNTMSMediaRequest) (*InjectNTMSMediaResponse, error)

	AccessNTMSLibraryDoor(context.Context, *AccessNTMSLibraryDoorRequest) (*AccessNTMSLibraryDoorResponse, error)

	CleanNTMSDrive(context.Context, *CleanNTMSDriveRequest) (*CleanNTMSDriveResponse, error)

	DismountNTMSDrive(context.Context, *DismountNTMSDriveRequest) (*DismountNTMSDriveResponse, error)

	InventoryNTMSLibrary(context.Context, *InventoryNTMSLibraryRequest) (*InventoryNTMSLibraryResponse, error)

	// INtmsLibraryControl1_LocalOnlyOpnum09 operation.
	NTMSLibraryControl1LocalOnlyOpnum09(context.Context, *NTMSLibraryControl1LocalOnlyOpnum09Request) (*NTMSLibraryControl1LocalOnlyOpnum09Response, error)

	CancelNTMSLibraryRequest(context.Context, *CancelNTMSLibraryRequestRequest) (*CancelNTMSLibraryRequestResponse, error)

	ReserveNTMSCleanerSlot(context.Context, *ReserveNTMSCleanerSlotRequest) (*ReserveNTMSCleanerSlotResponse, error)

	ReleaseNTMSCleanerSlot(context.Context, *ReleaseNTMSCleanerSlotRequest) (*ReleaseNTMSCleanerSlotResponse, error)

	InjectNTMSCleaner(context.Context, *InjectNTMSCleanerRequest) (*InjectNTMSCleanerResponse, error)

	EjectNTMSCleaner(context.Context, *EjectNTMSCleanerRequest) (*EjectNTMSCleanerResponse, error)

	DeleteNTMSLibrary(context.Context, *DeleteNTMSLibraryRequest) (*DeleteNTMSLibraryResponse, error)

	DeleteNTMSDrive(context.Context, *DeleteNTMSDriveRequest) (*DeleteNTMSDriveResponse, error)

	GetNTMSRequestOrder(context.Context, *GetNTMSRequestOrderRequest) (*GetNTMSRequestOrderResponse, error)

	SetNTMSRequestOrder(context.Context, *SetNTMSRequestOrderRequest) (*SetNTMSRequestOrderResponse, error)

	DeleteNTMSRequests(context.Context, *DeleteNTMSRequestsRequest) (*DeleteNTMSRequestsResponse, error)

	BeginNTMSDeviceChangeDetection(context.Context, *BeginNTMSDeviceChangeDetectionRequest) (*BeginNTMSDeviceChangeDetectionResponse, error)

	SetNTMSDeviceChangeDetection(context.Context, *SetNTMSDeviceChangeDetectionRequest) (*SetNTMSDeviceChangeDetectionResponse, error)

	EndNTMSDeviceChangeDetection(context.Context, *EndNTMSDeviceChangeDetectionRequest) (*EndNTMSDeviceChangeDetectionResponse, error)
}

func RegisterNTMSLibraryControl1Server(conn dcerpc.Conn, o NTMSLibraryControl1Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewNTMSLibraryControl1ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(NTMSLibraryControl1SyntaxV0_0))...)
}

func NewNTMSLibraryControl1ServerHandle(o NTMSLibraryControl1Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return NTMSLibraryControl1ServerHandle(ctx, o, opNum, r)
	}
}

func NTMSLibraryControl1ServerHandle(ctx context.Context, o NTMSLibraryControl1Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // EjectNtmsMedia
		op := &xxx_EjectNTMSMediaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EjectNTMSMediaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EjectNTMSMedia(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // InjectNtmsMedia
		op := &xxx_InjectNTMSMediaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InjectNTMSMediaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InjectNTMSMedia(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // AccessNtmsLibraryDoor
		op := &xxx_AccessNTMSLibraryDoorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AccessNTMSLibraryDoorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AccessNTMSLibraryDoor(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // CleanNtmsDrive
		op := &xxx_CleanNTMSDriveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CleanNTMSDriveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CleanNTMSDrive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // DismountNtmsDrive
		op := &xxx_DismountNTMSDriveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DismountNTMSDriveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DismountNTMSDrive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // InventoryNtmsLibrary
		op := &xxx_InventoryNTMSLibraryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InventoryNTMSLibraryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InventoryNTMSLibrary(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // INtmsLibraryControl1_LocalOnlyOpnum09
		op := &xxx_NTMSLibraryControl1LocalOnlyOpnum09Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NTMSLibraryControl1LocalOnlyOpnum09Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NTMSLibraryControl1LocalOnlyOpnum09(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // CancelNtmsLibraryRequest
		op := &xxx_CancelNTMSLibraryRequestOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CancelNTMSLibraryRequestRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CancelNTMSLibraryRequest(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ReserveNtmsCleanerSlot
		op := &xxx_ReserveNTMSCleanerSlotOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReserveNTMSCleanerSlotRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReserveNTMSCleanerSlot(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ReleaseNtmsCleanerSlot
		op := &xxx_ReleaseNTMSCleanerSlotOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReleaseNTMSCleanerSlotRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReleaseNTMSCleanerSlot(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // InjectNtmsCleaner
		op := &xxx_InjectNTMSCleanerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InjectNTMSCleanerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InjectNTMSCleaner(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // EjectNtmsCleaner
		op := &xxx_EjectNTMSCleanerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EjectNTMSCleanerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EjectNTMSCleaner(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // DeleteNtmsLibrary
		op := &xxx_DeleteNTMSLibraryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteNTMSLibraryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteNTMSLibrary(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // DeleteNtmsDrive
		op := &xxx_DeleteNTMSDriveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteNTMSDriveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteNTMSDrive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // GetNtmsRequestOrder
		op := &xxx_GetNTMSRequestOrderOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSRequestOrderRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSRequestOrder(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // SetNtmsRequestOrder
		op := &xxx_SetNTMSRequestOrderOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNTMSRequestOrderRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNTMSRequestOrder(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // DeleteNtmsRequests
		op := &xxx_DeleteNTMSRequestsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteNTMSRequestsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteNTMSRequests(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // BeginNtmsDeviceChangeDetection
		op := &xxx_BeginNTMSDeviceChangeDetectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BeginNTMSDeviceChangeDetectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BeginNTMSDeviceChangeDetection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // SetNtmsDeviceChangeDetection
		op := &xxx_SetNTMSDeviceChangeDetectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNTMSDeviceChangeDetectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNTMSDeviceChangeDetection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // EndNtmsDeviceChangeDetection
		op := &xxx_EndNTMSDeviceChangeDetectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EndNTMSDeviceChangeDetectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EndNTMSDeviceChangeDetection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented INtmsLibraryControl1
type UnimplementedNTMSLibraryControl1Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedNTMSLibraryControl1Server) EjectNTMSMedia(context.Context, *EjectNTMSMediaRequest) (*EjectNTMSMediaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) InjectNTMSMedia(context.Context, *InjectNTMSMediaRequest) (*InjectNTMSMediaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) AccessNTMSLibraryDoor(context.Context, *AccessNTMSLibraryDoorRequest) (*AccessNTMSLibraryDoorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) CleanNTMSDrive(context.Context, *CleanNTMSDriveRequest) (*CleanNTMSDriveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) DismountNTMSDrive(context.Context, *DismountNTMSDriveRequest) (*DismountNTMSDriveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) InventoryNTMSLibrary(context.Context, *InventoryNTMSLibraryRequest) (*InventoryNTMSLibraryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) NTMSLibraryControl1LocalOnlyOpnum09(context.Context, *NTMSLibraryControl1LocalOnlyOpnum09Request) (*NTMSLibraryControl1LocalOnlyOpnum09Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) CancelNTMSLibraryRequest(context.Context, *CancelNTMSLibraryRequestRequest) (*CancelNTMSLibraryRequestResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) ReserveNTMSCleanerSlot(context.Context, *ReserveNTMSCleanerSlotRequest) (*ReserveNTMSCleanerSlotResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) ReleaseNTMSCleanerSlot(context.Context, *ReleaseNTMSCleanerSlotRequest) (*ReleaseNTMSCleanerSlotResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) InjectNTMSCleaner(context.Context, *InjectNTMSCleanerRequest) (*InjectNTMSCleanerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) EjectNTMSCleaner(context.Context, *EjectNTMSCleanerRequest) (*EjectNTMSCleanerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) DeleteNTMSLibrary(context.Context, *DeleteNTMSLibraryRequest) (*DeleteNTMSLibraryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) DeleteNTMSDrive(context.Context, *DeleteNTMSDriveRequest) (*DeleteNTMSDriveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) GetNTMSRequestOrder(context.Context, *GetNTMSRequestOrderRequest) (*GetNTMSRequestOrderResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) SetNTMSRequestOrder(context.Context, *SetNTMSRequestOrderRequest) (*SetNTMSRequestOrderResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) DeleteNTMSRequests(context.Context, *DeleteNTMSRequestsRequest) (*DeleteNTMSRequestsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) BeginNTMSDeviceChangeDetection(context.Context, *BeginNTMSDeviceChangeDetectionRequest) (*BeginNTMSDeviceChangeDetectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) SetNTMSDeviceChangeDetection(context.Context, *SetNTMSDeviceChangeDetectionRequest) (*SetNTMSDeviceChangeDetectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSLibraryControl1Server) EndNTMSDeviceChangeDetection(context.Context, *EndNTMSDeviceChangeDetectionRequest) (*EndNTMSDeviceChangeDetectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ NTMSLibraryControl1Server = (*UnimplementedNTMSLibraryControl1Server)(nil)
