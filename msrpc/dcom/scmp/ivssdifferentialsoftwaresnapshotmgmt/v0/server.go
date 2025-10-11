package ivssdifferentialsoftwaresnapshotmgmt

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

// IVssDifferentialSoftwareSnapshotMgmt server interface.
type DifferentialSoftwareSnapshotManagementServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The AddDiffArea method creates a shadow copy storage association for a shadow copy.
	//
	// Return Values: The method MUST return the following error code for the specific conditions.
	//
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                         RETURN                          |                                                                                  |
	//	|                       VALUE/CODE                        |                                   DESCRIPTION                                    |
	//	|                                                         |                                                                                  |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004230d VSS_E_OBJECT_ALREADY_EXISTS                  | The object already exists on the server.                                         |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                                 | R returned when pwszVolumeName or pwszDiffAreaVolumeName is NULL, or if          |
	//	|                                                         | llMaximumDiffSpace is 0.                                                         |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004230c VSS_E_VOLUME_NOT_SUPPORTED                   | Returned when the pwszVolumeName does not support shadow copies, or              |
	//	|                                                         | pwszDiffAreaVolumeName does not support shadow copy storage.                     |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004231e VSS_E_MAXIMUM_DIFF-AREA_ASSOCIATIONS_REACHED | Returned when the maximum number of diff area associations for pwszVolumeName    |
	//	|                                                         | has been reached.                                                                |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80042306 VSS_E_PROVIDER_VETO                          | Returned when the snapshot provider receives an expected error and tries to veto |
	//	|                                                         | the impending operation.                                                         |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                               | Returned when the user making the request does not have sufficient privileges to |
	//	|                                                         | perform the operation.                                                           |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol
	// specified in [MS-RPCE].
	AddDiffArea(context.Context, *AddDiffAreaRequest) (*AddDiffAreaResponse, error)

	// The ChangeDiffAreaMaximumSize method changes the maximum size of a shadow copy storage
	// association on the server.
	//
	// Return Values: The method MUST return the following error code for the specific conditions.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80042308 VSS_E_OBJECT_NOT_FOUND     | The object does not exist on the server.                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG               | Returned when pwszVolumeName or pwszDiffAreaVolume is NULL.                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004231d VSS_E_VOLUME_IN_USE        | Returned when llMaximumDiffSpace is zero, and the diff area cannot be deleted    |
	//	|                                       | because shadow copies are still being stored.                                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004231f VSS_E_INSUFFICIENT_STORAGE | Returned if a nonzero size is specified in llMaximumDiffSpace that is smaller    |
	//	|                                       | than the size required for storing a single shadow copy.                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED             | Returned when the user making the request does not have sufficient privileges to |
	//	|                                       | perform the operation.                                                           |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol
	// [MS-RPCE].
	ChangeDiffAreaMaximumSize(context.Context, *ChangeDiffAreaMaximumSizeRequest) (*ChangeDiffAreaMaximumSizeResponse, error)

	// The QueryVolumesSupportedForDiffAreas method retrieves from the server the collection
	// of volumes that can be used as a shadow copy storage volume for a specified original
	// volume.
	//
	// Return Values: The method MUST return zero when it has succeeded or an implementation-specific
	// nonzero error code on failure.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|          RETURN           |                                                                                  |
	//	|        VALUE/CODE         |                                   DESCRIPTION                                    |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG   | Returned when pwszOriginalVolumeName or ppEnum is NULL.                          |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED | Returned when the user making the request does not have sufficient privileges to |
	//	|                           | perform the operation.                                                           |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol
	// [MS-RPCE].
	QueryVolumesSupportedForDiffAreas(context.Context, *QueryVolumesSupportedForDiffAreasRequest) (*QueryVolumesSupportedForDiffAreasResponse, error)

	// The QueryDiffAreasForVolume method retrieves from the server the collection of shadow
	// copy storage associations that are being used for shadow copy storage for a specified
	// original volume.
	//
	// Return Values: The method MUST return zero when it has succeeded or an implementation-specific
	// nonzero error code on failure.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|          RETURN           |                                                                                  |
	//	|        VALUE/CODE         |                                   DESCRIPTION                                    |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG   | Returned when pwszVolumeName or ppEnum is NULL.                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED | Returned when the user making the request does not have sufficient privileges to |
	//	|                           | perform the operation.                                                           |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol
	// [MS-RPCE].
	QueryDiffAreasForVolume(context.Context, *QueryDiffAreasForVolumeRequest) (*QueryDiffAreasForVolumeResponse, error)

	// The QueryDiffAreasOnVolume method retrieves from the server the collection of shadow
	// copy storage associations that are located on a specified volume.
	//
	// Return Values: The method MUST return zero when it has succeeded or an implementation-specific
	// nonzero error code on failure.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|          RETURN           |                                                                                  |
	//	|        VALUE/CODE         |                                   DESCRIPTION                                    |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG   | Returned when pwszVolumeName or ppEnum is NULL.                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED | Returned when the user making the request does not have sufficient privileges to |
	//	|                           | perform the operation.                                                           |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol
	// [MS-RPCE].
	QueryDiffAreasOnVolume(context.Context, *QueryDiffAreasOnVolumeRequest) (*QueryDiffAreasOnVolumeResponse, error)

	// Opnum08NotUsedOnWire operation.
	// Opnum08NotUsedOnWire
}

func RegisterDifferentialSoftwareSnapshotManagementServer(conn dcerpc.Conn, o DifferentialSoftwareSnapshotManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDifferentialSoftwareSnapshotManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DifferentialSoftwareSnapshotManagementSyntaxV0_0))...)
}

func NewDifferentialSoftwareSnapshotManagementServerHandle(o DifferentialSoftwareSnapshotManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DifferentialSoftwareSnapshotManagementServerHandle(ctx, o, opNum, r)
	}
}

func DifferentialSoftwareSnapshotManagementServerHandle(ctx context.Context, o DifferentialSoftwareSnapshotManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // AddDiffArea
		op := &xxx_AddDiffAreaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddDiffAreaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddDiffArea(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // ChangeDiffAreaMaximumSize
		op := &xxx_ChangeDiffAreaMaximumSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ChangeDiffAreaMaximumSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ChangeDiffAreaMaximumSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // QueryVolumesSupportedForDiffAreas
		op := &xxx_QueryVolumesSupportedForDiffAreasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryVolumesSupportedForDiffAreasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryVolumesSupportedForDiffAreas(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // QueryDiffAreasForVolume
		op := &xxx_QueryDiffAreasForVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryDiffAreasForVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryDiffAreasForVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // QueryDiffAreasOnVolume
		op := &xxx_QueryDiffAreasOnVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryDiffAreasOnVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryDiffAreasOnVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Opnum08NotUsedOnWire
		// Opnum08NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IVssDifferentialSoftwareSnapshotMgmt
type UnimplementedDifferentialSoftwareSnapshotManagementServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedDifferentialSoftwareSnapshotManagementServer) AddDiffArea(context.Context, *AddDiffAreaRequest) (*AddDiffAreaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDifferentialSoftwareSnapshotManagementServer) ChangeDiffAreaMaximumSize(context.Context, *ChangeDiffAreaMaximumSizeRequest) (*ChangeDiffAreaMaximumSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDifferentialSoftwareSnapshotManagementServer) QueryVolumesSupportedForDiffAreas(context.Context, *QueryVolumesSupportedForDiffAreasRequest) (*QueryVolumesSupportedForDiffAreasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDifferentialSoftwareSnapshotManagementServer) QueryDiffAreasForVolume(context.Context, *QueryDiffAreasForVolumeRequest) (*QueryDiffAreasForVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDifferentialSoftwareSnapshotManagementServer) QueryDiffAreasOnVolume(context.Context, *QueryDiffAreasOnVolumeRequest) (*QueryDiffAreasOnVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DifferentialSoftwareSnapshotManagementServer = (*UnimplementedDifferentialSoftwareSnapshotManagementServer)(nil)
