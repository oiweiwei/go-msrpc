package itransactionstream

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
	GoPackage = "dcom/com"
)

var (
	// ITransactionStream interface identifier 97199110-db2e-11d1-a251-0000f805ca53
	TransactionStreamIID = &dcom.IID{Data1: 0x97199110, Data2: 0xdb2e, Data3: 0x11d1, Data4: []byte{0xa2, 0x51, 0x00, 0x00, 0xf8, 0x05, 0xca, 0x53}}
	// Syntax UUID
	TransactionStreamSyntaxUUID = &uuid.UUID{TimeLow: 0x97199110, TimeMid: 0xdb2e, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x51, Node: [6]uint8{0x0, 0x0, 0xf8, 0x5, 0xca, 0x53}}
	// Syntax ID
	TransactionStreamSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: TransactionStreamSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ITransactionStream interface.
type TransactionStreamClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method returns the STxInfo ([MS-DTCO] section 2.2.5.10) of the currently active
	// transaction and the CurrentTSN value.
	//
	// Return Values: The method MUST return a positive value or zero, to indicate successful
	// completion, or a negative value to indicate failure. The client MUST treat any negative
	// return value as a fatal error.
	GetSeqAndTxViaExport(context.Context, *GetSeqAndTxViaExportRequest, ...dcerpc.CallOption) (*GetSeqAndTxViaExportResponse, error)

	// This method returns the Propagation_Token (as specified in [MS-DTCO] section 2.2.5.4)
	// of the currently active transaction and the CurrentTSN value.
	//
	// Return Values: The method MUST return a positive value or zero, to indicate successful
	// completion, or a negative value to indicate failure. The client MUST treat any negative
	// return value as a fatal error.
	GetSeqAndTxViaTransmitter(context.Context, *GetSeqAndTxViaTransmitterRequest, ...dcerpc.CallOption) (*GetSeqAndTxViaTransmitterResponse, error)

	// This method returns the STxInfo instance (as specified in [MS-DTCO] section 2.2.5.10)
	// of the currently active transaction or returns an error if the specified TSN is not
	// the same as the CurrentTSN value.
	//
	// Return Values: The method MUST return a positive value or zero to indicate successful
	// completion, or a negative value to indicate failure. The client MUST treat any negative
	// return value as a fatal error.
	GetTxViaExport(context.Context, *GetTxViaExportRequest, ...dcerpc.CallOption) (*GetTxViaExportResponse, error)

	// This method returns the Propagation_Token ([MS-DTCO] section 2.2.5.4) of the currently
	// active transaction, or returns an error if the specified TSN is not the same as the
	// CurrentTSN value.
	//
	// Return Values: The method MUST return a positive value or zero to indicate successful
	// completion, or a negative value to indicate failure. The client MUST treat any negative
	// return value as a fatal error.
	GetTxViaTransmitter(context.Context, *GetTxViaTransmitterRequest, ...dcerpc.CallOption) (*GetTxViaTransmitterResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) TransactionStreamClient
}

type xxx_DefaultTransactionStreamClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultTransactionStreamClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultTransactionStreamClient) GetSeqAndTxViaExport(ctx context.Context, in *GetSeqAndTxViaExportRequest, opts ...dcerpc.CallOption) (*GetSeqAndTxViaExportResponse, error) {
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
	out := &GetSeqAndTxViaExportResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransactionStreamClient) GetSeqAndTxViaTransmitter(ctx context.Context, in *GetSeqAndTxViaTransmitterRequest, opts ...dcerpc.CallOption) (*GetSeqAndTxViaTransmitterResponse, error) {
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
	out := &GetSeqAndTxViaTransmitterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransactionStreamClient) GetTxViaExport(ctx context.Context, in *GetTxViaExportRequest, opts ...dcerpc.CallOption) (*GetTxViaExportResponse, error) {
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
	out := &GetTxViaExportResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransactionStreamClient) GetTxViaTransmitter(ctx context.Context, in *GetTxViaTransmitterRequest, opts ...dcerpc.CallOption) (*GetTxViaTransmitterResponse, error) {
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
	out := &GetTxViaTransmitterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransactionStreamClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTransactionStreamClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultTransactionStreamClient) IPID(ctx context.Context, ipid *dcom.IPID) TransactionStreamClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultTransactionStreamClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewTransactionStreamClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TransactionStreamClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TransactionStreamSyntaxV0_0))...)
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
	return &xxx_DefaultTransactionStreamClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetSeqAndTxViaExportOperation structure represents the GetSeqAndTxViaExport operation
type xxx_GetSeqAndTxViaExportOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	KnownSeq           uint32         `idl:"name:ulKnownSeq" json:"known_seq"`
	WhereaboutsLength  uint32         `idl:"name:ulcbWhereabouts" json:"whereabouts_length"`
	Whereabouts        []byte         `idl:"name:rgbWhereabouts;size_is:(ulcbWhereabouts)" json:"whereabouts"`
	CurrentSeq         uint32         `idl:"name:pulCurrentSeq" json:"current_seq"`
	ExportCookieLength uint32         `idl:"name:pulcbExportCookie" json:"export_cookie_length"`
	ExportCookie       []byte         `idl:"name:prgbExportCookie;size_is:(, pulcbExportCookie)" json:"export_cookie"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSeqAndTxViaExportOperation) OpNum() int { return 3 }

func (o *xxx_GetSeqAndTxViaExportOperation) OpName() string {
	return "/ITransactionStream/v0/GetSeqAndTxViaExport"
}

func (o *xxx_GetSeqAndTxViaExportOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Whereabouts != nil && o.WhereaboutsLength == 0 {
		o.WhereaboutsLength = uint32(len(o.Whereabouts))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSeqAndTxViaExportOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulKnownSeq {in} (1:(uint32))
	{
		if err := w.WriteData(o.KnownSeq); err != nil {
			return err
		}
	}
	// ulcbWhereabouts {in} (1:(uint32))
	{
		if err := w.WriteData(o.WhereaboutsLength); err != nil {
			return err
		}
	}
	// rgbWhereabouts {in} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=ulcbWhereabouts](uchar))
	{
		dimSize1 := uint64(o.WhereaboutsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Whereabouts {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Whereabouts[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Whereabouts); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetSeqAndTxViaExportOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulKnownSeq {in} (1:(uint32))
	{
		if err := w.ReadData(&o.KnownSeq); err != nil {
			return err
		}
	}
	// ulcbWhereabouts {in} (1:(uint32))
	{
		if err := w.ReadData(&o.WhereaboutsLength); err != nil {
			return err
		}
	}
	// rgbWhereabouts {in} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=ulcbWhereabouts](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Whereabouts", sizeInfo[0])
		}
		o.Whereabouts = make([]byte, sizeInfo[0])
		for i1 := range o.Whereabouts {
			i1 := i1
			if err := w.ReadData(&o.Whereabouts[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetSeqAndTxViaExportOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ExportCookie != nil && o.ExportCookieLength == 0 {
		o.ExportCookieLength = uint32(len(o.ExportCookie))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSeqAndTxViaExportOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pulCurrentSeq {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.CurrentSeq); err != nil {
			return err
		}
	}
	// pulcbExportCookie {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ExportCookieLength); err != nil {
			return err
		}
	}
	// prgbExportCookie {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=pulcbExportCookie](uchar))
	{
		if o.ExportCookie != nil || o.ExportCookieLength > 0 {
			_ptr_prgbExportCookie := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ExportCookieLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ExportCookie {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ExportCookie[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ExportCookie); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExportCookie, _ptr_prgbExportCookie); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSeqAndTxViaExportOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pulCurrentSeq {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.CurrentSeq); err != nil {
			return err
		}
	}
	// pulcbExportCookie {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ExportCookieLength); err != nil {
			return err
		}
	}
	// prgbExportCookie {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=pulcbExportCookie](uchar))
	{
		_ptr_prgbExportCookie := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ExportCookie", sizeInfo[0])
			}
			o.ExportCookie = make([]byte, sizeInfo[0])
			for i1 := range o.ExportCookie {
				i1 := i1
				if err := w.ReadData(&o.ExportCookie[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_prgbExportCookie := func(ptr interface{}) { o.ExportCookie = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ExportCookie, _s_prgbExportCookie, _ptr_prgbExportCookie); err != nil {
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

// GetSeqAndTxViaExportRequest structure represents the GetSeqAndTxViaExport operation request
type GetSeqAndTxViaExportRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ulKnownSeq: The caller's CurrentTSN value of the currently active transaction known
	// by the client.
	KnownSeq uint32 `idl:"name:ulKnownSeq" json:"known_seq"`
	// ulcbWhereabouts: The unsigned size, in bytes, of rgbWhereabouts.
	WhereaboutsLength uint32 `idl:"name:ulcbWhereabouts" json:"whereabouts_length"`
	// rgbWhereabouts: The SWhereabouts instance ([MS-DTCO] section 2.2.5.11) of the caller's
	// local DTCO transaction manager implementation.
	Whereabouts []byte `idl:"name:rgbWhereabouts;size_is:(ulcbWhereabouts)" json:"whereabouts"`
}

func (o *GetSeqAndTxViaExportRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSeqAndTxViaExportOperation) *xxx_GetSeqAndTxViaExportOperation {
	if op == nil {
		op = &xxx_GetSeqAndTxViaExportOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.KnownSeq = o.KnownSeq
	op.WhereaboutsLength = o.WhereaboutsLength
	op.Whereabouts = o.Whereabouts
	return op
}

func (o *GetSeqAndTxViaExportRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSeqAndTxViaExportOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.KnownSeq = op.KnownSeq
	o.WhereaboutsLength = op.WhereaboutsLength
	o.Whereabouts = op.Whereabouts
}
func (o *GetSeqAndTxViaExportRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSeqAndTxViaExportRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSeqAndTxViaExportOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSeqAndTxViaExportResponse structure represents the GetSeqAndTxViaExport operation response
type GetSeqAndTxViaExportResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pulCurrentSeq: The TSN of the currently active transaction.
	CurrentSeq uint32 `idl:"name:pulCurrentSeq" json:"current_seq"`
	// pulcbExportCookie: The unsigned size, in bytes, of prgbExportCookie.
	ExportCookieLength uint32 `idl:"name:pulcbExportCookie" json:"export_cookie_length"`
	// prgbExportCookie: An STxInfo of the currently active transaction (as specified in
	// [MS-DTCO] section 2.2.5.10).
	ExportCookie []byte `idl:"name:prgbExportCookie;size_is:(, pulcbExportCookie)" json:"export_cookie"`
	// Return: The GetSeqAndTxViaExport return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSeqAndTxViaExportResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSeqAndTxViaExportOperation) *xxx_GetSeqAndTxViaExportOperation {
	if op == nil {
		op = &xxx_GetSeqAndTxViaExportOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.CurrentSeq = o.CurrentSeq
	op.ExportCookieLength = o.ExportCookieLength
	op.ExportCookie = o.ExportCookie
	op.Return = o.Return
	return op
}

func (o *GetSeqAndTxViaExportResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSeqAndTxViaExportOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CurrentSeq = op.CurrentSeq
	o.ExportCookieLength = op.ExportCookieLength
	o.ExportCookie = op.ExportCookie
	o.Return = op.Return
}
func (o *GetSeqAndTxViaExportResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSeqAndTxViaExportResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSeqAndTxViaExportOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSeqAndTxViaTransmitterOperation structure represents the GetSeqAndTxViaTransmitter operation
type xxx_GetSeqAndTxViaTransmitterOperation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	KnownSeq                uint32         `idl:"name:ulKnownSeq" json:"known_seq"`
	CurrentSeq              uint32         `idl:"name:pulCurrentSeq" json:"current_seq"`
	TransmitterBufferLength uint32         `idl:"name:pulcbTransmitterBuffer" json:"transmitter_buffer_length"`
	TransmitterBuffer       []byte         `idl:"name:prgbTransmitterBuffer;size_is:(, pulcbTransmitterBuffer)" json:"transmitter_buffer"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSeqAndTxViaTransmitterOperation) OpNum() int { return 4 }

func (o *xxx_GetSeqAndTxViaTransmitterOperation) OpName() string {
	return "/ITransactionStream/v0/GetSeqAndTxViaTransmitter"
}

func (o *xxx_GetSeqAndTxViaTransmitterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSeqAndTxViaTransmitterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulKnownSeq {in} (1:(uint32))
	{
		if err := w.WriteData(o.KnownSeq); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSeqAndTxViaTransmitterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulKnownSeq {in} (1:(uint32))
	{
		if err := w.ReadData(&o.KnownSeq); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSeqAndTxViaTransmitterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.TransmitterBuffer != nil && o.TransmitterBufferLength == 0 {
		o.TransmitterBufferLength = uint32(len(o.TransmitterBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSeqAndTxViaTransmitterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pulCurrentSeq {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.CurrentSeq); err != nil {
			return err
		}
	}
	// pulcbTransmitterBuffer {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TransmitterBufferLength); err != nil {
			return err
		}
	}
	// prgbTransmitterBuffer {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=pulcbTransmitterBuffer](uchar))
	{
		if o.TransmitterBuffer != nil || o.TransmitterBufferLength > 0 {
			_ptr_prgbTransmitterBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.TransmitterBufferLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.TransmitterBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.TransmitterBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.TransmitterBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TransmitterBuffer, _ptr_prgbTransmitterBuffer); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSeqAndTxViaTransmitterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pulCurrentSeq {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.CurrentSeq); err != nil {
			return err
		}
	}
	// pulcbTransmitterBuffer {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TransmitterBufferLength); err != nil {
			return err
		}
	}
	// prgbTransmitterBuffer {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=pulcbTransmitterBuffer](uchar))
	{
		_ptr_prgbTransmitterBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.TransmitterBuffer", sizeInfo[0])
			}
			o.TransmitterBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.TransmitterBuffer {
				i1 := i1
				if err := w.ReadData(&o.TransmitterBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_prgbTransmitterBuffer := func(ptr interface{}) { o.TransmitterBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.TransmitterBuffer, _s_prgbTransmitterBuffer, _ptr_prgbTransmitterBuffer); err != nil {
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

// GetSeqAndTxViaTransmitterRequest structure represents the GetSeqAndTxViaTransmitter operation request
type GetSeqAndTxViaTransmitterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ulKnownSeq: The caller's CurrentTSN value of the currently active transaction.
	KnownSeq uint32 `idl:"name:ulKnownSeq" json:"known_seq"`
}

func (o *GetSeqAndTxViaTransmitterRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSeqAndTxViaTransmitterOperation) *xxx_GetSeqAndTxViaTransmitterOperation {
	if op == nil {
		op = &xxx_GetSeqAndTxViaTransmitterOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.KnownSeq = o.KnownSeq
	return op
}

func (o *GetSeqAndTxViaTransmitterRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSeqAndTxViaTransmitterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.KnownSeq = op.KnownSeq
}
func (o *GetSeqAndTxViaTransmitterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSeqAndTxViaTransmitterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSeqAndTxViaTransmitterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSeqAndTxViaTransmitterResponse structure represents the GetSeqAndTxViaTransmitter operation response
type GetSeqAndTxViaTransmitterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pulCurrentSeq: The TSN of the currently active transaction.
	CurrentSeq uint32 `idl:"name:pulCurrentSeq" json:"current_seq"`
	// pulcbTransmitterBuffer: The unsigned size, in bytes, of prgbTransmitterBuffer.
	TransmitterBufferLength uint32 `idl:"name:pulcbTransmitterBuffer" json:"transmitter_buffer_length"`
	// prgbTransmitterBuffer: A Propagation_Token of the currently active transaction.
	TransmitterBuffer []byte `idl:"name:prgbTransmitterBuffer;size_is:(, pulcbTransmitterBuffer)" json:"transmitter_buffer"`
	// Return: The GetSeqAndTxViaTransmitter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSeqAndTxViaTransmitterResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSeqAndTxViaTransmitterOperation) *xxx_GetSeqAndTxViaTransmitterOperation {
	if op == nil {
		op = &xxx_GetSeqAndTxViaTransmitterOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.CurrentSeq = o.CurrentSeq
	op.TransmitterBufferLength = o.TransmitterBufferLength
	op.TransmitterBuffer = o.TransmitterBuffer
	op.Return = o.Return
	return op
}

func (o *GetSeqAndTxViaTransmitterResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSeqAndTxViaTransmitterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CurrentSeq = op.CurrentSeq
	o.TransmitterBufferLength = op.TransmitterBufferLength
	o.TransmitterBuffer = op.TransmitterBuffer
	o.Return = op.Return
}
func (o *GetSeqAndTxViaTransmitterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSeqAndTxViaTransmitterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSeqAndTxViaTransmitterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTxViaExportOperation structure represents the GetTxViaExport operation
type xxx_GetTxViaExportOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	RequestSeq         uint32         `idl:"name:ulRequestSeq" json:"request_seq"`
	WhereaboutsLength  uint32         `idl:"name:ulcbWhereabouts" json:"whereabouts_length"`
	Whereabouts        []byte         `idl:"name:rgbWhereabouts;size_is:(ulcbWhereabouts)" json:"whereabouts"`
	ExportCookieLength uint32         `idl:"name:pulcbExportCookie" json:"export_cookie_length"`
	ExportCookie       []byte         `idl:"name:prgbExportCookie;size_is:(, pulcbExportCookie)" json:"export_cookie"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTxViaExportOperation) OpNum() int { return 5 }

func (o *xxx_GetTxViaExportOperation) OpName() string { return "/ITransactionStream/v0/GetTxViaExport" }

func (o *xxx_GetTxViaExportOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Whereabouts != nil && o.WhereaboutsLength == 0 {
		o.WhereaboutsLength = uint32(len(o.Whereabouts))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTxViaExportOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulRequestSeq {in} (1:(uint32))
	{
		if err := w.WriteData(o.RequestSeq); err != nil {
			return err
		}
	}
	// ulcbWhereabouts {in} (1:(uint32))
	{
		if err := w.WriteData(o.WhereaboutsLength); err != nil {
			return err
		}
	}
	// rgbWhereabouts {in} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=ulcbWhereabouts](uchar))
	{
		dimSize1 := uint64(o.WhereaboutsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Whereabouts {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Whereabouts[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Whereabouts); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetTxViaExportOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulRequestSeq {in} (1:(uint32))
	{
		if err := w.ReadData(&o.RequestSeq); err != nil {
			return err
		}
	}
	// ulcbWhereabouts {in} (1:(uint32))
	{
		if err := w.ReadData(&o.WhereaboutsLength); err != nil {
			return err
		}
	}
	// rgbWhereabouts {in} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=ulcbWhereabouts](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Whereabouts", sizeInfo[0])
		}
		o.Whereabouts = make([]byte, sizeInfo[0])
		for i1 := range o.Whereabouts {
			i1 := i1
			if err := w.ReadData(&o.Whereabouts[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetTxViaExportOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ExportCookie != nil && o.ExportCookieLength == 0 {
		o.ExportCookieLength = uint32(len(o.ExportCookie))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTxViaExportOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pulcbExportCookie {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ExportCookieLength); err != nil {
			return err
		}
	}
	// prgbExportCookie {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=pulcbExportCookie](uchar))
	{
		if o.ExportCookie != nil || o.ExportCookieLength > 0 {
			_ptr_prgbExportCookie := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ExportCookieLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ExportCookie {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ExportCookie[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ExportCookie); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExportCookie, _ptr_prgbExportCookie); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTxViaExportOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pulcbExportCookie {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ExportCookieLength); err != nil {
			return err
		}
	}
	// prgbExportCookie {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=pulcbExportCookie](uchar))
	{
		_ptr_prgbExportCookie := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ExportCookie", sizeInfo[0])
			}
			o.ExportCookie = make([]byte, sizeInfo[0])
			for i1 := range o.ExportCookie {
				i1 := i1
				if err := w.ReadData(&o.ExportCookie[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_prgbExportCookie := func(ptr interface{}) { o.ExportCookie = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ExportCookie, _s_prgbExportCookie, _ptr_prgbExportCookie); err != nil {
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

// GetTxViaExportRequest structure represents the GetTxViaExport operation request
type GetTxViaExportRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ulRequestSeq: The caller's CurrentTSN value of the currently active transaction.
	RequestSeq uint32 `idl:"name:ulRequestSeq" json:"request_seq"`
	// ulcbWhereabouts: The unsigned size, in bytes, of rgbWhereabouts.
	WhereaboutsLength uint32 `idl:"name:ulcbWhereabouts" json:"whereabouts_length"`
	// rgbWhereabouts: The SWhereabouts instance ([MS-DTCO] section 2.2.5.11) of the caller's
	// local DTCO transaction manager implementation.
	Whereabouts []byte `idl:"name:rgbWhereabouts;size_is:(ulcbWhereabouts)" json:"whereabouts"`
}

func (o *GetTxViaExportRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTxViaExportOperation) *xxx_GetTxViaExportOperation {
	if op == nil {
		op = &xxx_GetTxViaExportOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RequestSeq = o.RequestSeq
	op.WhereaboutsLength = o.WhereaboutsLength
	op.Whereabouts = o.Whereabouts
	return op
}

func (o *GetTxViaExportRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTxViaExportOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RequestSeq = op.RequestSeq
	o.WhereaboutsLength = op.WhereaboutsLength
	o.Whereabouts = op.Whereabouts
}
func (o *GetTxViaExportRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTxViaExportRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTxViaExportOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTxViaExportResponse structure represents the GetTxViaExport operation response
type GetTxViaExportResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pulcbExportCookie: The unsigned size, in bytes, of prgbExportCookie.
	ExportCookieLength uint32 `idl:"name:pulcbExportCookie" json:"export_cookie_length"`
	// prgbExportCookie: An STxInfo instance of the currently active transaction (as specified
	// in [MS-DTCO] section 2.2.5.10).
	ExportCookie []byte `idl:"name:prgbExportCookie;size_is:(, pulcbExportCookie)" json:"export_cookie"`
	// Return: The GetTxViaExport return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTxViaExportResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTxViaExportOperation) *xxx_GetTxViaExportOperation {
	if op == nil {
		op = &xxx_GetTxViaExportOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ExportCookieLength = o.ExportCookieLength
	op.ExportCookie = o.ExportCookie
	op.Return = o.Return
	return op
}

func (o *GetTxViaExportResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTxViaExportOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ExportCookieLength = op.ExportCookieLength
	o.ExportCookie = op.ExportCookie
	o.Return = op.Return
}
func (o *GetTxViaExportResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTxViaExportResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTxViaExportOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTxViaTransmitterOperation structure represents the GetTxViaTransmitter operation
type xxx_GetTxViaTransmitterOperation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	RequestSeq              uint32         `idl:"name:ulRequestSeq" json:"request_seq"`
	TransmitterBufferLength uint32         `idl:"name:pulcbTransmitterBuffer" json:"transmitter_buffer_length"`
	TransmitterBuffer       []byte         `idl:"name:prgbTransmitterBuffer;size_is:(, pulcbTransmitterBuffer)" json:"transmitter_buffer"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTxViaTransmitterOperation) OpNum() int { return 6 }

func (o *xxx_GetTxViaTransmitterOperation) OpName() string {
	return "/ITransactionStream/v0/GetTxViaTransmitter"
}

func (o *xxx_GetTxViaTransmitterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTxViaTransmitterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulRequestSeq {in} (1:(uint32))
	{
		if err := w.WriteData(o.RequestSeq); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTxViaTransmitterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulRequestSeq {in} (1:(uint32))
	{
		if err := w.ReadData(&o.RequestSeq); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTxViaTransmitterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.TransmitterBuffer != nil && o.TransmitterBufferLength == 0 {
		o.TransmitterBufferLength = uint32(len(o.TransmitterBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTxViaTransmitterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pulcbTransmitterBuffer {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TransmitterBufferLength); err != nil {
			return err
		}
	}
	// prgbTransmitterBuffer {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=pulcbTransmitterBuffer](uchar))
	{
		if o.TransmitterBuffer != nil || o.TransmitterBufferLength > 0 {
			_ptr_prgbTransmitterBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.TransmitterBufferLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.TransmitterBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.TransmitterBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.TransmitterBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TransmitterBuffer, _ptr_prgbTransmitterBuffer); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTxViaTransmitterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pulcbTransmitterBuffer {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TransmitterBufferLength); err != nil {
			return err
		}
	}
	// prgbTransmitterBuffer {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=pulcbTransmitterBuffer](uchar))
	{
		_ptr_prgbTransmitterBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.TransmitterBuffer", sizeInfo[0])
			}
			o.TransmitterBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.TransmitterBuffer {
				i1 := i1
				if err := w.ReadData(&o.TransmitterBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_prgbTransmitterBuffer := func(ptr interface{}) { o.TransmitterBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.TransmitterBuffer, _s_prgbTransmitterBuffer, _ptr_prgbTransmitterBuffer); err != nil {
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

// GetTxViaTransmitterRequest structure represents the GetTxViaTransmitter operation request
type GetTxViaTransmitterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ulRequestSeq: The caller's TSN of the currently active transaction.
	RequestSeq uint32 `idl:"name:ulRequestSeq" json:"request_seq"`
}

func (o *GetTxViaTransmitterRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTxViaTransmitterOperation) *xxx_GetTxViaTransmitterOperation {
	if op == nil {
		op = &xxx_GetTxViaTransmitterOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RequestSeq = o.RequestSeq
	return op
}

func (o *GetTxViaTransmitterRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTxViaTransmitterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RequestSeq = op.RequestSeq
}
func (o *GetTxViaTransmitterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTxViaTransmitterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTxViaTransmitterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTxViaTransmitterResponse structure represents the GetTxViaTransmitter operation response
type GetTxViaTransmitterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pulcbTransmitterBuffer: The unsigned size, in bytes, of prgbTransmitterBuffer.
	TransmitterBufferLength uint32 `idl:"name:pulcbTransmitterBuffer" json:"transmitter_buffer_length"`
	// prgbTransmitterBuffer: A Propagation_Token of the currently active transaction.
	TransmitterBuffer []byte `idl:"name:prgbTransmitterBuffer;size_is:(, pulcbTransmitterBuffer)" json:"transmitter_buffer"`
	// Return: The GetTxViaTransmitter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTxViaTransmitterResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTxViaTransmitterOperation) *xxx_GetTxViaTransmitterOperation {
	if op == nil {
		op = &xxx_GetTxViaTransmitterOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.TransmitterBufferLength = o.TransmitterBufferLength
	op.TransmitterBuffer = o.TransmitterBuffer
	op.Return = o.Return
	return op
}

func (o *GetTxViaTransmitterResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTxViaTransmitterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TransmitterBufferLength = op.TransmitterBufferLength
	o.TransmitterBuffer = op.TransmitterBuffer
	o.Return = op.Return
}
func (o *GetTxViaTransmitterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTxViaTransmitterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTxViaTransmitterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
