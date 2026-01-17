package mgmt

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
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
	_ = dcetypes.GoPackage
)

var (
	// import guard
	GoPackage = "mgmt"
)

var (
	// Syntax UUID
	ManagementSyntaxUUID = &uuid.UUID{TimeLow: 0xafa8bd80, TimeMid: 0x7d8a, TimeHiAndVersion: 0x11c9, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0xf4, Node: [6]uint8{0x8, 0x0, 0x2b, 0x10, 0x29, 0x89}}
	// Syntax ID
	ManagementSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: ManagementSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// mgmt interface.
type ManagementClient interface {

	// rpc__mgmt_inq_if_ids operation.
	InquireInterfaceIDs(context.Context, *InquireInterfaceIDsRequest, ...dcerpc.CallOption) (*InquireInterfaceIDsResponse, error)

	// rpc__mgmt_inq_stats operation.
	InquireStats(context.Context, *InquireStatsRequest, ...dcerpc.CallOption) (*InquireStatsResponse, error)

	// rpc__mgmt_is_server_listening operation.
	IsServerListening(context.Context, *IsServerListeningRequest, ...dcerpc.CallOption) (*IsServerListeningResponse, error)

	// rpc__mgmt_stop_server_listening operation.
	StopServerListening(context.Context, *StopServerListeningRequest, ...dcerpc.CallOption) (*StopServerListeningResponse, error)

	// rpc__mgmt_inq_princ_name operation.
	InquirePrincName(context.Context, *InquirePrincNameRequest, ...dcerpc.CallOption) (*InquirePrincNameResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultManagementClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultManagementClient) InquireInterfaceIDs(ctx context.Context, in *InquireInterfaceIDsRequest, opts ...dcerpc.CallOption) (*InquireInterfaceIDsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InquireInterfaceIDsResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultManagementClient) InquireStats(ctx context.Context, in *InquireStatsRequest, opts ...dcerpc.CallOption) (*InquireStatsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InquireStatsResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultManagementClient) IsServerListening(ctx context.Context, in *IsServerListeningRequest, opts ...dcerpc.CallOption) (*IsServerListeningResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IsServerListeningResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManagementClient) StopServerListening(ctx context.Context, in *StopServerListeningRequest, opts ...dcerpc.CallOption) (*StopServerListeningResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StopServerListeningResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultManagementClient) InquirePrincName(ctx context.Context, in *InquirePrincNameRequest, opts ...dcerpc.CallOption) (*InquirePrincNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InquirePrincNameResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultManagementClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultManagementClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewManagementClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ManagementClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ManagementSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultManagementClient{cc: cc}, nil
}

// xxx_InquireInterfaceIDsOperation structure represents the rpc__mgmt_inq_if_ids operation
type xxx_InquireInterfaceIDsOperation struct {
	InterfaceIDVector *dcetypes.InterfaceIDVector `idl:"name:if_id_vector" json:"interface_id_vector"`
	Status            uint32                      `idl:"name:status" json:"status"`
}

func (o *xxx_InquireInterfaceIDsOperation) OpNum() int { return 0 }

func (o *xxx_InquireInterfaceIDsOperation) OpName() string { return "/mgmt/v1/rpc__mgmt_inq_if_ids" }

func (o *xxx_InquireInterfaceIDsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquireInterfaceIDsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_InquireInterfaceIDsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_InquireInterfaceIDsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquireInterfaceIDsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// if_id_vector {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=rpc_if_id_vector_p_t}*(1))(3:{alias=rpc_if_id_vector_t}(struct))
	{
		if o.InterfaceIDVector != nil {
			_ptr_if_id_vector := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InterfaceIDVector != nil {
					if err := o.InterfaceIDVector.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcetypes.InterfaceIDVector{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InterfaceIDVector, _ptr_if_id_vector); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquireInterfaceIDsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// if_id_vector {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=rpc_if_id_vector_p_t}*(1))(3:{alias=rpc_if_id_vector_t}(struct))
	{
		_ptr_if_id_vector := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InterfaceIDVector == nil {
				o.InterfaceIDVector = &dcetypes.InterfaceIDVector{}
			}
			if err := o.InterfaceIDVector.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_if_id_vector := func(ptr interface{}) { o.InterfaceIDVector = *ptr.(**dcetypes.InterfaceIDVector) }
		if err := w.ReadPointer(&o.InterfaceIDVector, _s_if_id_vector, _ptr_if_id_vector); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// InquireInterfaceIDsRequest structure represents the rpc__mgmt_inq_if_ids operation request
type InquireInterfaceIDsRequest struct {
}

func (o *InquireInterfaceIDsRequest) xxx_ToOp(ctx context.Context, op *xxx_InquireInterfaceIDsOperation) *xxx_InquireInterfaceIDsOperation {
	if op == nil {
		op = &xxx_InquireInterfaceIDsOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *InquireInterfaceIDsRequest) xxx_FromOp(ctx context.Context, op *xxx_InquireInterfaceIDsOperation) {
	if o == nil {
		return
	}
}
func (o *InquireInterfaceIDsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InquireInterfaceIDsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InquireInterfaceIDsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InquireInterfaceIDsResponse structure represents the rpc__mgmt_inq_if_ids operation response
type InquireInterfaceIDsResponse struct {
	InterfaceIDVector *dcetypes.InterfaceIDVector `idl:"name:if_id_vector" json:"interface_id_vector"`
	Status            uint32                      `idl:"name:status" json:"status"`
}

func (o *InquireInterfaceIDsResponse) xxx_ToOp(ctx context.Context, op *xxx_InquireInterfaceIDsOperation) *xxx_InquireInterfaceIDsOperation {
	if op == nil {
		op = &xxx_InquireInterfaceIDsOperation{}
	}
	if o == nil {
		return op
	}
	op.InterfaceIDVector = o.InterfaceIDVector
	op.Status = o.Status
	return op
}

func (o *InquireInterfaceIDsResponse) xxx_FromOp(ctx context.Context, op *xxx_InquireInterfaceIDsOperation) {
	if o == nil {
		return
	}
	o.InterfaceIDVector = op.InterfaceIDVector
	o.Status = op.Status
}
func (o *InquireInterfaceIDsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InquireInterfaceIDsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InquireInterfaceIDsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InquireStatsOperation structure represents the rpc__mgmt_inq_stats operation
type xxx_InquireStatsOperation struct {
	Count      uint32   `idl:"name:count" json:"count"`
	Statistics []uint32 `idl:"name:statistics;size_is:(count)" json:"statistics"`
	Status     uint32   `idl:"name:status" json:"status"`
}

func (o *xxx_InquireStatsOperation) OpNum() int { return 1 }

func (o *xxx_InquireStatsOperation) OpName() string { return "/mgmt/v1/rpc__mgmt_inq_stats" }

func (o *xxx_InquireStatsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquireStatsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// count {in, out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquireStatsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// count {in, out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquireStatsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Statistics != nil && o.Count == 0 {
		o.Count = uint32(len(o.Statistics))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquireStatsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// count {in, out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// statistics {out} (1:[dim:0,size_is=count])(2:{alias=unsigned32}(uint32))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Statistics {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Statistics[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Statistics); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquireStatsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// count {in, out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// statistics {out} (1:[dim:0,size_is=count])(2:{alias=unsigned32}(uint32))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Statistics", sizeInfo[0])
		}
		o.Statistics = make([]uint32, sizeInfo[0])
		for i1 := range o.Statistics {
			i1 := i1
			if err := w.ReadData(&o.Statistics[i1]); err != nil {
				return err
			}
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// InquireStatsRequest structure represents the rpc__mgmt_inq_stats operation request
type InquireStatsRequest struct {
	Count uint32 `idl:"name:count" json:"count"`
}

func (o *InquireStatsRequest) xxx_ToOp(ctx context.Context, op *xxx_InquireStatsOperation) *xxx_InquireStatsOperation {
	if op == nil {
		op = &xxx_InquireStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.Count = o.Count
	return op
}

func (o *InquireStatsRequest) xxx_FromOp(ctx context.Context, op *xxx_InquireStatsOperation) {
	if o == nil {
		return
	}
	o.Count = op.Count
}
func (o *InquireStatsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InquireStatsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InquireStatsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InquireStatsResponse structure represents the rpc__mgmt_inq_stats operation response
type InquireStatsResponse struct {
	Count      uint32   `idl:"name:count" json:"count"`
	Statistics []uint32 `idl:"name:statistics;size_is:(count)" json:"statistics"`
	Status     uint32   `idl:"name:status" json:"status"`
}

func (o *InquireStatsResponse) xxx_ToOp(ctx context.Context, op *xxx_InquireStatsOperation) *xxx_InquireStatsOperation {
	if op == nil {
		op = &xxx_InquireStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.Count = o.Count
	op.Statistics = o.Statistics
	op.Status = o.Status
	return op
}

func (o *InquireStatsResponse) xxx_FromOp(ctx context.Context, op *xxx_InquireStatsOperation) {
	if o == nil {
		return
	}
	o.Count = op.Count
	o.Statistics = op.Statistics
	o.Status = op.Status
}
func (o *InquireStatsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InquireStatsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InquireStatsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsServerListeningOperation structure represents the rpc__mgmt_is_server_listening operation
type xxx_IsServerListeningOperation struct {
	Status uint32 `idl:"name:status" json:"status"`
	Return bool   `idl:"name:Return" json:"return"`
}

func (o *xxx_IsServerListeningOperation) OpNum() int { return 2 }

func (o *xxx_IsServerListeningOperation) OpName() string {
	return "/mgmt/v1/rpc__mgmt_is_server_listening"
}

func (o *xxx_IsServerListeningOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsServerListeningOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_IsServerListeningOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_IsServerListeningOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsServerListeningOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=boolean32, names=unsigned32}(uint32))
	{
		if !o.Return {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(uint32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_IsServerListeningOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=boolean32, names=unsigned32}(uint32))
	{
		var _bReturn uint32
		if err := w.ReadData(&_bReturn); err != nil {
			return err
		}
		o.Return = _bReturn != 0
	}
	return nil
}

// IsServerListeningRequest structure represents the rpc__mgmt_is_server_listening operation request
type IsServerListeningRequest struct {
}

func (o *IsServerListeningRequest) xxx_ToOp(ctx context.Context, op *xxx_IsServerListeningOperation) *xxx_IsServerListeningOperation {
	if op == nil {
		op = &xxx_IsServerListeningOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *IsServerListeningRequest) xxx_FromOp(ctx context.Context, op *xxx_IsServerListeningOperation) {
	if o == nil {
		return
	}
}
func (o *IsServerListeningRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsServerListeningRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsServerListeningOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsServerListeningResponse structure represents the rpc__mgmt_is_server_listening operation response
type IsServerListeningResponse struct {
	Status uint32 `idl:"name:status" json:"status"`
	// Return: The rpc__mgmt_is_server_listening return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *IsServerListeningResponse) xxx_ToOp(ctx context.Context, op *xxx_IsServerListeningOperation) *xxx_IsServerListeningOperation {
	if op == nil {
		op = &xxx_IsServerListeningOperation{}
	}
	if o == nil {
		return op
	}
	op.Status = o.Status
	op.Return = o.Return
	return op
}

func (o *IsServerListeningResponse) xxx_FromOp(ctx context.Context, op *xxx_IsServerListeningOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
	o.Return = op.Return
}
func (o *IsServerListeningResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsServerListeningResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsServerListeningOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StopServerListeningOperation structure represents the rpc__mgmt_stop_server_listening operation
type xxx_StopServerListeningOperation struct {
	Status uint32 `idl:"name:status" json:"status"`
}

func (o *xxx_StopServerListeningOperation) OpNum() int { return 3 }

func (o *xxx_StopServerListeningOperation) OpName() string {
	return "/mgmt/v1/rpc__mgmt_stop_server_listening"
}

func (o *xxx_StopServerListeningOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopServerListeningOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_StopServerListeningOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_StopServerListeningOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopServerListeningOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopServerListeningOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// StopServerListeningRequest structure represents the rpc__mgmt_stop_server_listening operation request
type StopServerListeningRequest struct {
}

func (o *StopServerListeningRequest) xxx_ToOp(ctx context.Context, op *xxx_StopServerListeningOperation) *xxx_StopServerListeningOperation {
	if op == nil {
		op = &xxx_StopServerListeningOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *StopServerListeningRequest) xxx_FromOp(ctx context.Context, op *xxx_StopServerListeningOperation) {
	if o == nil {
		return
	}
}
func (o *StopServerListeningRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StopServerListeningRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopServerListeningOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StopServerListeningResponse structure represents the rpc__mgmt_stop_server_listening operation response
type StopServerListeningResponse struct {
	Status uint32 `idl:"name:status" json:"status"`
}

func (o *StopServerListeningResponse) xxx_ToOp(ctx context.Context, op *xxx_StopServerListeningOperation) *xxx_StopServerListeningOperation {
	if op == nil {
		op = &xxx_StopServerListeningOperation{}
	}
	if o == nil {
		return op
	}
	op.Status = o.Status
	return op
}

func (o *StopServerListeningResponse) xxx_FromOp(ctx context.Context, op *xxx_StopServerListeningOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
}
func (o *StopServerListeningResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StopServerListeningResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopServerListeningOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InquirePrincNameOperation structure represents the rpc__mgmt_inq_princ_name operation
type xxx_InquirePrincNameOperation struct {
	AuthnProto    uint32 `idl:"name:authn_proto" json:"authn_proto"`
	PrincNameSize uint32 `idl:"name:princ_name_size" json:"princ_name_size"`
	PrincName     string `idl:"name:princ_name;size_is:(princ_name_size);string" json:"princ_name"`
	Status        uint32 `idl:"name:status" json:"status"`
}

func (o *xxx_InquirePrincNameOperation) OpNum() int { return 4 }

func (o *xxx_InquirePrincNameOperation) OpName() string { return "/mgmt/v1/rpc__mgmt_inq_princ_name" }

func (o *xxx_InquirePrincNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquirePrincNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// authn_proto {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.AuthnProto); err != nil {
			return err
		}
	}
	// princ_name_size {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.PrincNameSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquirePrincNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// authn_proto {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.AuthnProto); err != nil {
			return err
		}
	}
	// princ_name_size {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.PrincNameSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquirePrincNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquirePrincNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// princ_name {out} (1:{string}[dim:0,size_is=princ_name_size,string,null](char))
	{
		dimSize1 := uint64(o.PrincNameSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.CharNLen(o.PrincName)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		_PrincName_buf := []byte(o.PrincName)
		if uint64(len(_PrincName_buf)) > sizeInfo[0]-1 {
			_PrincName_buf = _PrincName_buf[:sizeInfo[0]-1]
		}
		if o.PrincName != ndr.ZeroString {
			_PrincName_buf = append(_PrincName_buf, byte(0))
		}
		for i1 := range _PrincName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_PrincName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_PrincName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InquirePrincNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// princ_name {out} (1:{string}[dim:0,size_is=princ_name_size,string,null](char))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _PrincName_buf []byte
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _PrincName_buf", sizeInfo[0])
		}
		_PrincName_buf = make([]byte, sizeInfo[0])
		for i1 := range _PrincName_buf {
			i1 := i1
			if err := w.ReadData(&_PrincName_buf[i1]); err != nil {
				return err
			}
		}
		o.PrincName = strings.TrimRight(string(_PrincName_buf), ndr.ZeroString)
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// InquirePrincNameRequest structure represents the rpc__mgmt_inq_princ_name operation request
type InquirePrincNameRequest struct {
	AuthnProto    uint32 `idl:"name:authn_proto" json:"authn_proto"`
	PrincNameSize uint32 `idl:"name:princ_name_size" json:"princ_name_size"`
}

func (o *InquirePrincNameRequest) xxx_ToOp(ctx context.Context, op *xxx_InquirePrincNameOperation) *xxx_InquirePrincNameOperation {
	if op == nil {
		op = &xxx_InquirePrincNameOperation{}
	}
	if o == nil {
		return op
	}
	op.AuthnProto = o.AuthnProto
	op.PrincNameSize = o.PrincNameSize
	return op
}

func (o *InquirePrincNameRequest) xxx_FromOp(ctx context.Context, op *xxx_InquirePrincNameOperation) {
	if o == nil {
		return
	}
	o.AuthnProto = op.AuthnProto
	o.PrincNameSize = op.PrincNameSize
}
func (o *InquirePrincNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InquirePrincNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InquirePrincNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InquirePrincNameResponse structure represents the rpc__mgmt_inq_princ_name operation response
type InquirePrincNameResponse struct {
	// XXX: princ_name_size is an implicit input depedency for output parameters
	PrincNameSize uint32 `idl:"name:princ_name_size" json:"princ_name_size"`

	PrincName string `idl:"name:princ_name;size_is:(princ_name_size);string" json:"princ_name"`
	Status    uint32 `idl:"name:status" json:"status"`
}

func (o *InquirePrincNameResponse) xxx_ToOp(ctx context.Context, op *xxx_InquirePrincNameOperation) *xxx_InquirePrincNameOperation {
	if op == nil {
		op = &xxx_InquirePrincNameOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.PrincNameSize == uint32(0) {
		op.PrincNameSize = o.PrincNameSize
	}

	op.PrincName = o.PrincName
	op.Status = o.Status
	return op
}

func (o *InquirePrincNameResponse) xxx_FromOp(ctx context.Context, op *xxx_InquirePrincNameOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.PrincNameSize = op.PrincNameSize

	o.PrincName = op.PrincName
	o.Status = op.Status
}
func (o *InquirePrincNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InquirePrincNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InquirePrincNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
