package iclustercleanup

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/csvp"
)

var (
	// IClusterCleanup interface identifier d6105110-8917-41a5-aa32-8e0aa2933dc9
	ClusterCleanupIID = &dcom.IID{Data1: 0xd6105110, Data2: 0x8917, Data3: 0x41a5, Data4: []byte{0xaa, 0x32, 0x8e, 0x0a, 0xa2, 0x93, 0x3d, 0xc9}}
	// Syntax UUID
	ClusterCleanupSyntaxUUID = &uuid.UUID{TimeLow: 0xd6105110, TimeMid: 0x8917, TimeHiAndVersion: 0x41a5, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x32, Node: [6]uint8{0x8e, 0xa, 0xa2, 0x93, 0x3d, 0xc9}}
	// Syntax ID
	ClusterCleanupSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClusterCleanupSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IClusterCleanup interface.
type ClusterCleanupClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The CleanUpEvictedNode method removes all persistent artifacts that exist on the
	// server after it is evicted from a cluster.
	//
	// This method is idempotent. After it is invoked, the target server can no longer be
	// a server for the Failover Cluster: Cluster Management Remote Protocol (ClusAPI) ([MS-CMRP])
	// until the server is reconfigured as a member of a cluster by using implementation-specific
	// methods between servers.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------------+---------------------------------------------------------------------------+
	//	|         RETURN          |                                                                           |
	//	|       VALUE/CODE        |                                DESCRIPTION                                |
	//	|                         |                                                                           |
	//	+-------------------------+---------------------------------------------------------------------------+
	//	+-------------------------+---------------------------------------------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.                                                  |
	//	+-------------------------+---------------------------------------------------------------------------+
	//	| 0x80070102 WAIT_TIMEOUT | The Cleanup Timer (section 3.8.2.2) expired before cleanup was completed. |
	//	+-------------------------+---------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 3.
	CleanupEvictedNode(context.Context, *CleanupEvictedNodeRequest, ...dcerpc.CallOption) (*CleanupEvictedNodeResponse, error)

	// The ClearPR method performs a SCSI PERSISTENT RESERVE OUT command (see [SPC-3] section
	// 6.12) with a REGISTER AND IGNORE EXISTING KEY action, followed by a CLEAR action.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------+--------------------------+
	//	|             RETURN              |                          |
	//	|           VALUE/CODE            |       DESCRIPTION        |
	//	|                                 |                          |
	//	+---------------------------------+--------------------------+
	//	+---------------------------------+--------------------------+
	//	| 0x00000000 S_OK                 | The call was successful. |
	//	+---------------------------------+--------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND | The disk was not found.  |
	//	+---------------------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 4.
	ClearPR(context.Context, *ClearPRRequest, ...dcerpc.CallOption) (*ClearPRResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClusterCleanupClient
}

type xxx_DefaultClusterCleanupClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClusterCleanupClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultClusterCleanupClient) CleanupEvictedNode(ctx context.Context, in *CleanupEvictedNodeRequest, opts ...dcerpc.CallOption) (*CleanupEvictedNodeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CleanupEvictedNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterCleanupClient) ClearPR(ctx context.Context, in *ClearPRRequest, opts ...dcerpc.CallOption) (*ClearPRResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClearPRResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterCleanupClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClusterCleanupClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClusterCleanupClient) IPID(ctx context.Context, ipid *dcom.IPID) ClusterCleanupClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClusterCleanupClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewClusterCleanupClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClusterCleanupClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClusterCleanupSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultClusterCleanupClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CleanupEvictedNodeOperation structure represents the CleanUpEvictedNode operation
type xxx_CleanupEvictedNodeOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	DelayBeforeCleanup uint32         `idl:"name:DelayBeforeCleanup" json:"delay_before_cleanup"`
	Timeout            uint32         `idl:"name:TimeOut" json:"timeout"`
	Flags              uint32         `idl:"name:Flags" json:"flags"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CleanupEvictedNodeOperation) OpNum() int { return 3 }

func (o *xxx_CleanupEvictedNodeOperation) OpName() string {
	return "/IClusterCleanup/v0/CleanUpEvictedNode"
}

func (o *xxx_CleanupEvictedNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupEvictedNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DelayBeforeCleanup {in} (1:(uint32))
	{
		if err := w.WriteData(o.DelayBeforeCleanup); err != nil {
			return err
		}
	}
	// TimeOut {in} (1:(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// Flags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupEvictedNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DelayBeforeCleanup {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DelayBeforeCleanup); err != nil {
			return err
		}
	}
	// TimeOut {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// Flags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupEvictedNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupEvictedNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupEvictedNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CleanupEvictedNodeRequest structure represents the CleanUpEvictedNode operation request
type CleanupEvictedNodeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DelayBeforeCleanup: The number of milliseconds that the server MUST delay before
	// cleanup is started on the target server. If this value is zero, the server is cleaned
	// up immediately.
	DelayBeforeCleanup uint32 `idl:"name:DelayBeforeCleanup" json:"delay_before_cleanup"`
	// TimeOut: The number of milliseconds that the server MUST wait for cleanup to complete.
	// This time-out is independent of the preceding delay; therefore, if DelayBeforeCleanup
	// is greater than TimeOut, this method will time out. However, after cleanup is initiated,
	// cleanup will run to completion regardless of the method waiting.
	Timeout uint32 `idl:"name:TimeOut" json:"timeout"`
	// Flags: A set of bit flags specifying the requested actions to be taken during cleanup.
	// This parameter MUST be set to at least one of the following values.
	//
	//	+----------------------------------------------------------+----------------------------------------------------------------------+
	//	|                                                          |                                                                      |
	//	|                          VALUE                           |                               MEANING                                |
	//	|                                                          |                                                                      |
	//	+----------------------------------------------------------+----------------------------------------------------------------------+
	//	+----------------------------------------------------------+----------------------------------------------------------------------+
	//	| CLUSTERCLEANUP_STOP_CLUSTER_SERVICE 0x00000000           | Issue a stop command to the cluster service and wait for it to stop. |
	//	+----------------------------------------------------------+----------------------------------------------------------------------+
	//	| CLUSTERCLEANUP_DONT_STOP_CLUSTER_SERVICE 0x00000001      | Do not issue a stop command to the cluster service.                  |
	//	+----------------------------------------------------------+----------------------------------------------------------------------+
	//	| CLUSTERCLEANUP_DONT_WAIT_CLUSTER_SERVICE_STOP 0x00000002 | Do not wait for the cluster service to stop.                         |
	//	+----------------------------------------------------------+----------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *CleanupEvictedNodeRequest) xxx_ToOp(ctx context.Context, op *xxx_CleanupEvictedNodeOperation) *xxx_CleanupEvictedNodeOperation {
	if op == nil {
		op = &xxx_CleanupEvictedNodeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DelayBeforeCleanup = o.DelayBeforeCleanup
	op.Timeout = o.Timeout
	op.Flags = o.Flags
	return op
}

func (o *CleanupEvictedNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_CleanupEvictedNodeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DelayBeforeCleanup = op.DelayBeforeCleanup
	o.Timeout = op.Timeout
	o.Flags = op.Flags
}
func (o *CleanupEvictedNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CleanupEvictedNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CleanupEvictedNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CleanupEvictedNodeResponse structure represents the CleanUpEvictedNode operation response
type CleanupEvictedNodeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CleanUpEvictedNode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CleanupEvictedNodeResponse) xxx_ToOp(ctx context.Context, op *xxx_CleanupEvictedNodeOperation) *xxx_CleanupEvictedNodeOperation {
	if op == nil {
		op = &xxx_CleanupEvictedNodeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CleanupEvictedNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_CleanupEvictedNodeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CleanupEvictedNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CleanupEvictedNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CleanupEvictedNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClearPROperation structure represents the ClearPR operation
type xxx_ClearPROperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	DeviceNumber uint32         `idl:"name:DeviceNumber" json:"device_number"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearPROperation) OpNum() int { return 4 }

func (o *xxx_ClearPROperation) OpName() string { return "/IClusterCleanup/v0/ClearPR" }

func (o *xxx_ClearPROperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearPROperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DeviceNumber {in} (1:(uint32))
	{
		if err := w.WriteData(o.DeviceNumber); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearPROperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DeviceNumber {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DeviceNumber); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearPROperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearPROperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearPROperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ClearPRRequest structure represents the ClearPR operation request
type ClearPRRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DeviceNumber:  The number of the disk to act on.
	DeviceNumber uint32 `idl:"name:DeviceNumber" json:"device_number"`
}

func (o *ClearPRRequest) xxx_ToOp(ctx context.Context, op *xxx_ClearPROperation) *xxx_ClearPROperation {
	if op == nil {
		op = &xxx_ClearPROperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DeviceNumber = o.DeviceNumber
	return op
}

func (o *ClearPRRequest) xxx_FromOp(ctx context.Context, op *xxx_ClearPROperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DeviceNumber = op.DeviceNumber
}
func (o *ClearPRRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ClearPRRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearPROperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearPRResponse structure represents the ClearPR operation response
type ClearPRResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClearPR return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ClearPRResponse) xxx_ToOp(ctx context.Context, op *xxx_ClearPROperation) *xxx_ClearPROperation {
	if op == nil {
		op = &xxx_ClearPROperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ClearPRResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearPROperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ClearPRResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ClearPRResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearPROperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
