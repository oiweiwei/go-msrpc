package iimport2

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
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// IImport2 interface identifier 1f7b1697-ecb2-4cbb-8a0e-75c427f4a6f0
	Import2IID = &dcom.IID{Data1: 0x1f7b1697, Data2: 0xecb2, Data3: 0x4cbb, Data4: []byte{0x8a, 0x0e, 0x75, 0xc4, 0x27, 0xf4, 0xa6, 0xf0}}
	// Syntax UUID
	Import2SyntaxUUID = &uuid.UUID{TimeLow: 0x1f7b1697, TimeMid: 0xecb2, TimeHiAndVersion: 0x4cbb, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0xe, Node: [6]uint8{0x75, 0xc4, 0x27, 0xf4, 0xa6, 0xf0}}
	// Syntax ID
	Import2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Import2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IImport2 interface.
type Import2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to set the import target partition (as specified
	// in section 3.1.1.5).
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result (as specified in [MS-ERREF], section 2.1) on failure. All failure results
	// MUST be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	SetPartition(context.Context, *SetPartitionRequest, ...dcerpc.CallOption) (*SetPartitionResponse, error)

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Import2Client
}

type xxx_DefaultImport2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultImport2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultImport2Client) SetPartition(ctx context.Context, in *SetPartitionRequest, opts ...dcerpc.CallOption) (*SetPartitionResponse, error) {
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
	out := &SetPartitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImport2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultImport2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultImport2Client) IPID(ctx context.Context, ipid *dcom.IPID) Import2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultImport2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewImport2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Import2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Import2SyntaxV0_0))...)
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
	return &xxx_DefaultImport2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_SetPartitionOperation structure represents the SetPartition operation
type xxx_SetPartitionOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	PartitionID *dtyp.GUID     `idl:"name:pPartitionIdentifier" json:"partition_id"`
	_           *dtyp.GUID     `idl:"name:pReserved"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPartitionOperation) OpNum() int { return 3 }

func (o *xxx_SetPartitionOperation) OpName() string { return "/IImport2/v0/SetPartition" }

func (o *xxx_SetPartitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPartitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pPartitionIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PartitionID != nil {
			if err := o.PartitionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetPartitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pPartitionIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PartitionID == nil {
			o.PartitionID = &dtyp.GUID{}
		}
		if err := o.PartitionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPartitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPartitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pReserved {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		// reserved pReserved
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetPartitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pReserved {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		// reserved pReserved
		var _pReserved *dtyp.GUID
		if _pReserved == nil {
			_pReserved = &dtyp.GUID{}
		}
		if err := _pReserved.UnmarshalNDR(ctx, w); err != nil {
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

// SetPartitionRequest structure represents the SetPartition operation request
type SetPartitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pPartitionIdentifier: A pointer to a variable containing a GUID to be used as the
	// partition identifier for the import target partition.
	PartitionID *dtyp.GUID `idl:"name:pPartitionIdentifier" json:"partition_id"`
}

func (o *SetPartitionRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPartitionOperation) *xxx_SetPartitionOperation {
	if op == nil {
		op = &xxx_SetPartitionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PartitionID = o.PartitionID
	return op
}

func (o *SetPartitionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPartitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PartitionID = op.PartitionID
}
func (o *SetPartitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPartitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPartitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPartitionResponse structure represents the SetPartition operation response
type SetPartitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetPartition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPartitionResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPartitionOperation) *xxx_SetPartitionOperation {
	if op == nil {
		op = &xxx_SetPartitionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetPartitionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPartitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPartitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPartitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPartitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
