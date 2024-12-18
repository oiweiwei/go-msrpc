package iwbemwcosmartenum

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
	GoPackage = "dcom/wmi"
)

var (
	// IWbemWCOSmartEnum interface identifier 423ec01e-2e35-11d2-b604-00104b703efd
	WCOSmartEnumIID = &dcom.IID{Data1: 0x423ec01e, Data2: 0x2e35, Data3: 0x11d2, Data4: []byte{0xb6, 0x04, 0x00, 0x10, 0x4b, 0x70, 0x3e, 0xfd}}
	// Syntax UUID
	WCOSmartEnumSyntaxUUID = &uuid.UUID{TimeLow: 0x423ec01e, TimeMid: 0x2e35, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0x4, Node: [6]uint8{0x0, 0x10, 0x4b, 0x70, 0x3e, 0xfd}}
	// Syntax ID
	WCOSmartEnumSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: WCOSmartEnumSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemWCOSmartEnum interface.
type WCOSmartEnumClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// Next operation.
	Next(context.Context, *NextRequest, ...dcerpc.CallOption) (*NextResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) WCOSmartEnumClient
}

type xxx_DefaultWCOSmartEnumClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultWCOSmartEnumClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultWCOSmartEnumClient) Next(ctx context.Context, in *NextRequest, opts ...dcerpc.CallOption) (*NextResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &NextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWCOSmartEnumClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWCOSmartEnumClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultWCOSmartEnumClient) IPID(ctx context.Context, ipid *dcom.IPID) WCOSmartEnumClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultWCOSmartEnumClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewWCOSmartEnumClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WCOSmartEnumClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WCOSmartEnumSyntaxV0_0))...)
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
	return &xxx_DefaultWCOSmartEnumClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_NextOperation structure represents the Next operation
type xxx_NextOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	ProxyGUID  *dtyp.GUID     `idl:"name:proxyGUID" json:"proxy_guid"`
	Timeout    int32          `idl:"name:lTimeout" json:"timeout"`
	Count      uint32         `idl:"name:uCount" json:"count"`
	Returned   uint32         `idl:"name:puReturned" json:"returned"`
	BufferSize uint32         `idl:"name:pdwBuffSize" json:"buffer_size"`
	Buffer     []byte         `idl:"name:pBuffer;size_is:(, pdwBuffSize)" json:"buffer"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_NextOperation) OpNum() int { return 3 }

func (o *xxx_NextOperation) OpName() string { return "/IWbemWCOSmartEnum/v0/Next" }

func (o *xxx_NextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// proxyGUID {in} (1:{alias=REFGUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ProxyGUID != nil {
			if err := o.ProxyGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lTimeout {in} (1:(int32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// uCount {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// proxyGUID {in} (1:{alias=REFGUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ProxyGUID == nil {
			o.ProxyGUID = &dtyp.GUID{}
		}
		if err := o.ProxyGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lTimeout {in} (1:(int32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// uCount {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// puReturned {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Returned); err != nil {
			return err
		}
	}
	// pdwBuffSize {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// pBuffer {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pdwBuffSize](uint8))
	{
		if o.Buffer != nil || o.BufferSize > 0 {
			_ptr_pBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Buffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Buffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_pBuffer); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// puReturned {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Returned); err != nil {
			return err
		}
	}
	// pdwBuffSize {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// pBuffer {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pdwBuffSize](uint8))
	{
		_ptr_pBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
			}
			o.Buffer = make([]byte, sizeInfo[0])
			for i1 := range o.Buffer {
				i1 := i1
				if err := w.ReadData(&o.Buffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pBuffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_pBuffer, _ptr_pBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NextRequest structure represents the Next operation request
type NextRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	ProxyGUID *dtyp.GUID     `idl:"name:proxyGUID" json:"proxy_guid"`
	Timeout   int32          `idl:"name:lTimeout" json:"timeout"`
	Count     uint32         `idl:"name:uCount" json:"count"`
}

func (o *NextRequest) xxx_ToOp(ctx context.Context) *xxx_NextOperation {
	if o == nil {
		return &xxx_NextOperation{}
	}
	return &xxx_NextOperation{
		This:      o.This,
		ProxyGUID: o.ProxyGUID,
		Timeout:   o.Timeout,
		Count:     o.Count,
	}
}

func (o *NextRequest) xxx_FromOp(ctx context.Context, op *xxx_NextOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ProxyGUID = op.ProxyGUID
	o.Timeout = op.Timeout
	o.Count = op.Count
}
func (o *NextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *NextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NextResponse structure represents the Next operation response
type NextResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Returned   uint32         `idl:"name:puReturned" json:"returned"`
	BufferSize uint32         `idl:"name:pdwBuffSize" json:"buffer_size"`
	Buffer     []byte         `idl:"name:pBuffer;size_is:(, pdwBuffSize)" json:"buffer"`
	// Return: The Next return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *NextResponse) xxx_ToOp(ctx context.Context) *xxx_NextOperation {
	if o == nil {
		return &xxx_NextOperation{}
	}
	return &xxx_NextOperation{
		That:       o.That,
		Returned:   o.Returned,
		BufferSize: o.BufferSize,
		Buffer:     o.Buffer,
		Return:     o.Return,
	}
}

func (o *NextResponse) xxx_FromOp(ctx context.Context, op *xxx_NextOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Returned = op.Returned
	o.BufferSize = op.BufferSize
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *NextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *NextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
