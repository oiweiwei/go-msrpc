package intmslibrarycontrol2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	intmslibrarycontrol1 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmslibrarycontrol1/v0"
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
	_ = intmslibrarycontrol1.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// INtmsLibraryControl2 interface identifier db90832f-6910-4d46-9f5e-9fd6bfa73903
	LibraryControl2IID = &dcom.IID{Data1: 0xdb90832f, Data2: 0x6910, Data3: 0x4d46, Data4: []byte{0x9f, 0x5e, 0x9f, 0xd6, 0xbf, 0xa7, 0x39, 0x03}}
	// Syntax UUID
	LibraryControl2SyntaxUUID = &uuid.UUID{TimeLow: 0xdb90832f, TimeMid: 0x6910, TimeHiAndVersion: 0x4d46, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0x5e, Node: [6]uint8{0x9f, 0xd6, 0xbf, 0xa7, 0x39, 0x3}}
	// Syntax ID
	LibraryControl2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: LibraryControl2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// INtmsLibraryControl2 interface.
type LibraryControl2Client interface {

	// INtmsLibraryControl1 retrieval method.
	LibraryControl1() intmslibrarycontrol1.LibraryControl1Client

	// The IdentifyNtmsSlot method identifies the media in a storage slot.
	IdentifyNTMSSlot(context.Context, *IdentifyNTMSSlotRequest, ...dcerpc.CallOption) (*IdentifyNTMSSlotResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) LibraryControl2Client
}

type xxx_DefaultLibraryControl2Client struct {
	intmslibrarycontrol1.LibraryControl1Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultLibraryControl2Client) LibraryControl1() intmslibrarycontrol1.LibraryControl1Client {
	return o.LibraryControl1Client
}

func (o *xxx_DefaultLibraryControl2Client) IdentifyNTMSSlot(ctx context.Context, in *IdentifyNTMSSlotRequest, opts ...dcerpc.CallOption) (*IdentifyNTMSSlotResponse, error) {
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
	out := &IdentifyNTMSSlotResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultLibraryControl2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultLibraryControl2Client) IPID(ctx context.Context, ipid *dcom.IPID) LibraryControl2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultLibraryControl2Client{
		LibraryControl1Client: o.LibraryControl1Client.IPID(ctx, ipid),
		cc:                    o.cc,
		ipid:                  ipid,
	}
}

func NewLibraryControl2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (LibraryControl2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(LibraryControl2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := intmslibrarycontrol1.NewLibraryControl1Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultLibraryControl2Client{
		LibraryControl1Client: base,
		cc:                    cc,
		ipid:                  ipid,
	}, nil
}

// xxx_IdentifyNTMSSlotOperation structure represents the IdentifyNtmsSlot operation
type xxx_IdentifyNTMSSlotOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	SlotID *dtyp.GUID     `idl:"name:lpSlotId" json:"slot_id"`
	Option uint32         `idl:"name:dwOption" json:"option"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IdentifyNTMSSlotOperation) OpNum() int { return 20 }

func (o *xxx_IdentifyNTMSSlotOperation) OpName() string {
	return "/INtmsLibraryControl2/v0/IdentifyNtmsSlot"
}

func (o *xxx_IdentifyNTMSSlotOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IdentifyNTMSSlotOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpSlotId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.SlotID != nil {
			if err := o.SlotID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwOption {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Option); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IdentifyNTMSSlotOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpSlotId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.SlotID == nil {
			o.SlotID = &dtyp.GUID{}
		}
		if err := o.SlotID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwOption {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Option); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IdentifyNTMSSlotOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IdentifyNTMSSlotOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IdentifyNTMSSlotOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// IdentifyNTMSSlotRequest structure represents the IdentifyNtmsSlot operation request
type IdentifyNTMSSlotRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpSlotId: A pointer to the identifier of the slot containing the media to identify
	SlotID *dtyp.GUID `idl:"name:lpSlotId" json:"slot_id"`
	// dwOption: A value from the NtmsDismountOptions (section 2.2.1.9) enumeration specifying
	// what to do with the media after identification is complete. If dwOption does not
	// contain a valid value, the default value NTMS_DISMOUNT_IMMEDIATE is taken.
	//
	//	+------------------------------------+------------------------------------------------------+
	//	|               RETURN               |                                                      |
	//	|             VALUE/CODE             |                     DESCRIPTION                      |
	//	|                                    |                                                      |
	//	+------------------------------------+------------------------------------------------------+
	//	+------------------------------------+------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                             |
	//	+------------------------------------+------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | Access to an object was denied.                      |
	//	+------------------------------------+------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | A parameter is not valid.                            |
	//	+------------------------------------+------------------------------------------------------+
	//	| 0x800710D1 ERROR_LIBRARY_OFFLINE   | The library identifier refers to an offline library. |
	//	+------------------------------------+------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | The object was not found.                            |
	//	+------------------------------------+------------------------------------------------------+
	Option uint32 `idl:"name:dwOption" json:"option"`
}

func (o *IdentifyNTMSSlotRequest) xxx_ToOp(ctx context.Context, op *xxx_IdentifyNTMSSlotOperation) *xxx_IdentifyNTMSSlotOperation {
	if op == nil {
		op = &xxx_IdentifyNTMSSlotOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SlotID = o.SlotID
	op.Option = o.Option
	return op
}

func (o *IdentifyNTMSSlotRequest) xxx_FromOp(ctx context.Context, op *xxx_IdentifyNTMSSlotOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SlotID = op.SlotID
	o.Option = op.Option
}
func (o *IdentifyNTMSSlotRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IdentifyNTMSSlotRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IdentifyNTMSSlotOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IdentifyNTMSSlotResponse structure represents the IdentifyNtmsSlot operation response
type IdentifyNTMSSlotResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The IdentifyNtmsSlot return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IdentifyNTMSSlotResponse) xxx_ToOp(ctx context.Context, op *xxx_IdentifyNTMSSlotOperation) *xxx_IdentifyNTMSSlotOperation {
	if op == nil {
		op = &xxx_IdentifyNTMSSlotOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *IdentifyNTMSSlotResponse) xxx_FromOp(ctx context.Context, op *xxx_IdentifyNTMSSlotOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *IdentifyNTMSSlotResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IdentifyNTMSSlotResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IdentifyNTMSSlotOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
