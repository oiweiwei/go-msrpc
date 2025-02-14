package icatalog64bitsupport

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// ICatalog64BitSupport interface identifier 1d118904-94b3-4a64-9fa6-ed432666a7b9
	Catalog64BitSupportIID = &dcom.IID{Data1: 0x1d118904, Data2: 0x94b3, Data3: 0x4a64, Data4: []byte{0x9f, 0xa6, 0xed, 0x43, 0x26, 0x66, 0xa7, 0xb9}}
	// Syntax UUID
	Catalog64BitSupportSyntaxUUID = &uuid.UUID{TimeLow: 0x1d118904, TimeMid: 0x94b3, TimeHiAndVersion: 0x4a64, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0xa6, Node: [6]uint8{0xed, 0x43, 0x26, 0x66, 0xa7, 0xb9}}
	// Syntax ID
	Catalog64BitSupportSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Catalog64BitSupportSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICatalog64BitSupport interface.
type Catalog64BitSupportClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to perform capability negotiation for the Multiple-Bitness
	// Capability (section 3.1.4.4).
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	SupportsMultipleBitness(context.Context, *SupportsMultipleBitnessRequest, ...dcerpc.CallOption) (*SupportsMultipleBitnessResponse, error)

	// This method is called by a client to perform capability negotiation for the 64-bit
	// QueryCell Marshaling Format Capability (section 3.1.4.2).
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	Initialize64BitQueryCellSupport(context.Context, *Initialize64BitQueryCellSupportRequest, ...dcerpc.CallOption) (*Initialize64BitQueryCellSupportResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Catalog64BitSupportClient
}

type xxx_DefaultCatalog64BitSupportClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCatalog64BitSupportClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCatalog64BitSupportClient) SupportsMultipleBitness(ctx context.Context, in *SupportsMultipleBitnessRequest, opts ...dcerpc.CallOption) (*SupportsMultipleBitnessResponse, error) {
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
	out := &SupportsMultipleBitnessResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalog64BitSupportClient) Initialize64BitQueryCellSupport(ctx context.Context, in *Initialize64BitQueryCellSupportRequest, opts ...dcerpc.CallOption) (*Initialize64BitQueryCellSupportResponse, error) {
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
	out := &Initialize64BitQueryCellSupportResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalog64BitSupportClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCatalog64BitSupportClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCatalog64BitSupportClient) IPID(ctx context.Context, ipid *dcom.IPID) Catalog64BitSupportClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCatalog64BitSupportClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewCatalog64BitSupportClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Catalog64BitSupportClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Catalog64BitSupportSyntaxV0_0))...)
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
	return &xxx_DefaultCatalog64BitSupportClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_SupportsMultipleBitnessOperation structure represents the SupportsMultipleBitness operation
type xxx_SupportsMultipleBitnessOperation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	SupportsMultipleBitness bool           `idl:"name:pbSupportsMultipleBitness" json:"supports_multiple_bitness"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SupportsMultipleBitnessOperation) OpNum() int { return 3 }

func (o *xxx_SupportsMultipleBitnessOperation) OpName() string {
	return "/ICatalog64BitSupport/v0/SupportsMultipleBitness"
}

func (o *xxx_SupportsMultipleBitnessOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SupportsMultipleBitnessOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_SupportsMultipleBitnessOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_SupportsMultipleBitnessOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SupportsMultipleBitnessOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbSupportsMultipleBitness {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.SupportsMultipleBitness {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
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

func (o *xxx_SupportsMultipleBitnessOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbSupportsMultipleBitness {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bSupportsMultipleBitness int32
		if err := w.ReadData(&_bSupportsMultipleBitness); err != nil {
			return err
		}
		o.SupportsMultipleBitness = _bSupportsMultipleBitness != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SupportsMultipleBitnessRequest structure represents the SupportsMultipleBitness operation request
type SupportsMultipleBitnessRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *SupportsMultipleBitnessRequest) xxx_ToOp(ctx context.Context, op *xxx_SupportsMultipleBitnessOperation) *xxx_SupportsMultipleBitnessOperation {
	if op == nil {
		op = &xxx_SupportsMultipleBitnessOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *SupportsMultipleBitnessRequest) xxx_FromOp(ctx context.Context, op *xxx_SupportsMultipleBitnessOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *SupportsMultipleBitnessRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SupportsMultipleBitnessRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SupportsMultipleBitnessOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SupportsMultipleBitnessResponse structure represents the SupportsMultipleBitness operation response
type SupportsMultipleBitnessResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbSupportsMultipleBitness: A pointer to a value that, upon successful completion,
	// indicates whether the server supports the multiple-bitness capability.
	SupportsMultipleBitness bool `idl:"name:pbSupportsMultipleBitness" json:"supports_multiple_bitness"`
	// Return: The SupportsMultipleBitness return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SupportsMultipleBitnessResponse) xxx_ToOp(ctx context.Context, op *xxx_SupportsMultipleBitnessOperation) *xxx_SupportsMultipleBitnessOperation {
	if op == nil {
		op = &xxx_SupportsMultipleBitnessOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SupportsMultipleBitness = o.SupportsMultipleBitness
	op.Return = o.Return
	return op
}

func (o *SupportsMultipleBitnessResponse) xxx_FromOp(ctx context.Context, op *xxx_SupportsMultipleBitnessOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SupportsMultipleBitness = op.SupportsMultipleBitness
	o.Return = op.Return
}
func (o *SupportsMultipleBitnessResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SupportsMultipleBitnessResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SupportsMultipleBitnessOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Initialize64BitQueryCellSupportOperation structure represents the Initialize64BitQueryCellSupport operation
type xxx_Initialize64BitQueryCellSupportOperation struct {
	This                          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ClientSupports64BitQueryCells bool           `idl:"name:bClientSupports64BitQueryCells" json:"client_supports64_bit_query_cells"`
	ServerSupports64BitQueryCells bool           `idl:"name:pbServerSupports64BitQueryCells" json:"server_supports64_bit_query_cells"`
	Return                        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_Initialize64BitQueryCellSupportOperation) OpNum() int { return 4 }

func (o *xxx_Initialize64BitQueryCellSupportOperation) OpName() string {
	return "/ICatalog64BitSupport/v0/Initialize64BitQueryCellSupport"
}

func (o *xxx_Initialize64BitQueryCellSupportOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Initialize64BitQueryCellSupportOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bClientSupports64BitQueryCells {in} (1:{alias=BOOL}(int32))
	{
		if !o.ClientSupports64BitQueryCells {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_Initialize64BitQueryCellSupportOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bClientSupports64BitQueryCells {in} (1:{alias=BOOL}(int32))
	{
		var _bClientSupports64BitQueryCells int32
		if err := w.ReadData(&_bClientSupports64BitQueryCells); err != nil {
			return err
		}
		o.ClientSupports64BitQueryCells = _bClientSupports64BitQueryCells != 0
	}
	return nil
}

func (o *xxx_Initialize64BitQueryCellSupportOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Initialize64BitQueryCellSupportOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbServerSupports64BitQueryCells {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.ServerSupports64BitQueryCells {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
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

func (o *xxx_Initialize64BitQueryCellSupportOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbServerSupports64BitQueryCells {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bServerSupports64BitQueryCells int32
		if err := w.ReadData(&_bServerSupports64BitQueryCells); err != nil {
			return err
		}
		o.ServerSupports64BitQueryCells = _bServerSupports64BitQueryCells != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// Initialize64BitQueryCellSupportRequest structure represents the Initialize64BitQueryCellSupport operation request
type Initialize64BitQueryCellSupportRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bClientSupports64BitQueryCells: A BOOL value that indicates whether the client supports
	// the 64-bit QueryCell Marshaling Format.
	ClientSupports64BitQueryCells bool `idl:"name:bClientSupports64BitQueryCells" json:"client_supports64_bit_query_cells"`
}

func (o *Initialize64BitQueryCellSupportRequest) xxx_ToOp(ctx context.Context, op *xxx_Initialize64BitQueryCellSupportOperation) *xxx_Initialize64BitQueryCellSupportOperation {
	if op == nil {
		op = &xxx_Initialize64BitQueryCellSupportOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ClientSupports64BitQueryCells = o.ClientSupports64BitQueryCells
	return op
}

func (o *Initialize64BitQueryCellSupportRequest) xxx_FromOp(ctx context.Context, op *xxx_Initialize64BitQueryCellSupportOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClientSupports64BitQueryCells = op.ClientSupports64BitQueryCells
}
func (o *Initialize64BitQueryCellSupportRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *Initialize64BitQueryCellSupportRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Initialize64BitQueryCellSupportOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Initialize64BitQueryCellSupportResponse structure represents the Initialize64BitQueryCellSupport operation response
type Initialize64BitQueryCellSupportResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbServerSupports64BitQueryCells: A pointer to a BOOL value that, upon successful
	// completion, indicates whether the server supports the 64-bit QueryCell Marshaling
	// Format.
	ServerSupports64BitQueryCells bool `idl:"name:pbServerSupports64BitQueryCells" json:"server_supports64_bit_query_cells"`
	// Return: The Initialize64BitQueryCellSupport return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Initialize64BitQueryCellSupportResponse) xxx_ToOp(ctx context.Context, op *xxx_Initialize64BitQueryCellSupportOperation) *xxx_Initialize64BitQueryCellSupportOperation {
	if op == nil {
		op = &xxx_Initialize64BitQueryCellSupportOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ServerSupports64BitQueryCells = o.ServerSupports64BitQueryCells
	op.Return = o.Return
	return op
}

func (o *Initialize64BitQueryCellSupportResponse) xxx_FromOp(ctx context.Context, op *xxx_Initialize64BitQueryCellSupportOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ServerSupports64BitQueryCells = op.ServerSupports64BitQueryCells
	o.Return = op.Return
}
func (o *Initialize64BitQueryCellSupportResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *Initialize64BitQueryCellSupportResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Initialize64BitQueryCellSupportOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
