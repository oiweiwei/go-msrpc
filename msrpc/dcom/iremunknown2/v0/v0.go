package iremunknown2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iremunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown/v0"
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
	_ = iremunknown.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom"
)

var (
	// IRemUnknown2 interface identifier 00000143-0000-0000-c000-000000000046
	RemoteUnknown2IID = &dcom.IID{Data1: 0x00000143, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	RemoteUnknown2SyntaxUUID = &uuid.UUID{TimeLow: 0x143, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	RemoteUnknown2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteUnknown2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRemUnknown2 interface.
type RemoteUnknown2Client interface {

	// IRemUnknown retrieval method.
	RemoteUnknown() iremunknown.RemoteUnknownClient

	// The RemQueryInterface2 (Opnum 6) method acquires standard object references (see
	// section 2.2.18.1) to additional interfaces on the object, marshaled as an MInterfacePointer
	// structure.
	RemoteQueryInterface2(context.Context, *RemoteQueryInterface2Request, ...dcerpc.CallOption) (*RemoteQueryInterface2Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteUnknown2Client
}

type xxx_DefaultRemoteUnknown2Client struct {
	iremunknown.RemoteUnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteUnknown2Client) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.RemoteUnknownClient
}

func (o *xxx_DefaultRemoteUnknown2Client) RemoteQueryInterface2(ctx context.Context, in *RemoteQueryInterface2Request, opts ...dcerpc.CallOption) (*RemoteQueryInterface2Response, error) {
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
	out := &RemoteQueryInterface2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteUnknown2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteUnknown2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRemoteUnknown2Client) IPID(ctx context.Context, ipid *dcom.IPID) RemoteUnknown2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteUnknown2Client{
		RemoteUnknownClient: o.RemoteUnknownClient.IPID(ctx, ipid),
		cc:                  o.cc,
		ipid:                ipid,
	}
}

func NewRemoteUnknown2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteUnknown2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteUnknown2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iremunknown.NewRemoteUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultRemoteUnknown2Client{
		RemoteUnknownClient: base,
		cc:                  cc,
		ipid:                ipid,
	}, nil
}

// xxx_RemoteQueryInterface2Operation structure represents the RemQueryInterface2 operation
type xxx_RemoteQueryInterface2Operation struct {
	This      *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat           `idl:"name:That" json:"that"`
	IPID      *dtyp.GUID               `idl:"name:ripid" json:"ipid"`
	IIDsCount uint16                   `idl:"name:cIids" json:"iids_count"`
	IIDs      []*dcom.IID              `idl:"name:iids;size_is:(cIids)" json:"iids"`
	HResult   []int32                  `idl:"name:phr;size_is:(cIids)" json:"hresult"`
	Interface []*dcom.InterfacePointer `idl:"name:ppMIF;size_is:(cIids)" json:"interface"`
	Return    int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteQueryInterface2Operation) OpNum() int { return 6 }

func (o *xxx_RemoteQueryInterface2Operation) OpName() string {
	return "/IRemUnknown2/v0/RemQueryInterface2"
}

func (o *xxx_RemoteQueryInterface2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.IIDs != nil && o.IIDsCount == 0 {
		o.IIDsCount = uint16(len(o.IIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteQueryInterface2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ripid {in} (1:{alias=REFIPID, names=REFGUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.IPID != nil {
			if err := o.IPID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// cIids {in} (1:(uint16))
	{
		if err := w.WriteData(o.IIDsCount); err != nil {
			return err
		}
	}
	// iids {in} (1:{pointer=ref}*(1))(2:{alias=IID, names=GUID}[dim:0,size_is=cIids](struct))
	{
		dimSize1 := uint64(o.IIDsCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.IIDs {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.IIDs[i1] != nil {
				if err := o.IIDs[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.IIDs); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RemoteQueryInterface2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ripid {in} (1:{alias=REFIPID, names=REFGUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.IPID == nil {
			o.IPID = &dtyp.GUID{}
		}
		if err := o.IPID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// cIids {in} (1:(uint16))
	{
		if err := w.ReadData(&o.IIDsCount); err != nil {
			return err
		}
	}
	// iids {in} (1:{pointer=ref}*(1))(2:{alias=IID, names=GUID}[dim:0,size_is=cIids](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.IIDs", sizeInfo[0])
		}
		o.IIDs = make([]*dcom.IID, sizeInfo[0])
		for i1 := range o.IIDs {
			i1 := i1
			if o.IIDs[i1] == nil {
				o.IIDs[i1] = &dcom.IID{}
			}
			if err := o.IIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RemoteQueryInterface2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteQueryInterface2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// phr {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}[dim:0,size_is=cIids](int32))
	{
		dimSize1 := uint64(o.IIDsCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.HResult {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.HResult[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.HResult); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		}
	}
	// ppMIF {out} (1:{pointer=ref}*(1))(2:{disable_consistency_check, alias=PMInterfacePointerInternal}[dim:0,size_is=cIids]*(1))(3:{alias=MInterfacePointer}(struct))
	{
		dimSize1 := uint64(o.IIDsCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Interface {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Interface[i1] != nil {
				_ptr_ppMIF := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.Interface[i1] != nil {
						if err := o.Interface[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dcom.InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Interface[i1], _ptr_ppMIF); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Interface); uint64(i1) < sizeInfo[0]; i1++ {
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

func (o *xxx_RemoteQueryInterface2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// phr {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}[dim:0,size_is=cIids](int32))
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
			return fmt.Errorf("buffer overflow for size %d of array o.HResult", sizeInfo[0])
		}
		o.HResult = make([]int32, sizeInfo[0])
		for i1 := range o.HResult {
			i1 := i1
			if err := w.ReadData(&o.HResult[i1]); err != nil {
				return err
			}
		}
	}
	// ppMIF {out} (1:{pointer=ref}*(1))(2:{disable_consistency_check, alias=PMInterfacePointerInternal}[dim:0,size_is=cIids]*(1))(3:{alias=MInterfacePointer}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Interface", sizeInfo[0])
		}
		o.Interface = make([]*dcom.InterfacePointer, sizeInfo[0])
		for i1 := range o.Interface {
			i1 := i1
			_ptr_ppMIF := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Interface[i1] == nil {
					o.Interface[i1] = &dcom.InterfacePointer{}
				}
				if err := o.Interface[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppMIF := func(ptr interface{}) { o.Interface[i1] = *ptr.(**dcom.InterfacePointer) }
			if err := w.ReadPointer(&o.Interface[i1], _s_ppMIF, _ptr_ppMIF); err != nil {
				return err
			}
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

// RemoteQueryInterface2Request structure represents the RemQueryInterface2 operation request
type RemoteQueryInterface2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ripid:  This MUST specify an IPID that identifies the interface on the object in
	// the object exporter that is queried for more interfaces.
	IPID *dtyp.GUID `idl:"name:ripid" json:"ipid"`
	// cIids:  This MUST specify the number of elements in the iids, phr, and ppMIF parameters.
	IIDsCount uint16 `idl:"name:cIids" json:"iids_count"`
	// iids:  This MUST specify an array of IIDs for which the client requests object references.
	IIDs []*dcom.IID `idl:"name:iids;size_is:(cIids)" json:"iids"`
}

func (o *RemoteQueryInterface2Request) xxx_ToOp(ctx context.Context) *xxx_RemoteQueryInterface2Operation {
	if o == nil {
		return &xxx_RemoteQueryInterface2Operation{}
	}
	return &xxx_RemoteQueryInterface2Operation{
		This:      o.This,
		IPID:      o.IPID,
		IIDsCount: o.IIDsCount,
		IIDs:      o.IIDs,
	}
}

func (o *RemoteQueryInterface2Request) xxx_FromOp(ctx context.Context, op *xxx_RemoteQueryInterface2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IPID = op.IPID
	o.IIDsCount = op.IIDsCount
	o.IIDs = op.IIDs
}
func (o *RemoteQueryInterface2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoteQueryInterface2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteQueryInterface2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteQueryInterface2Response structure represents the RemQueryInterface2 operation response
type RemoteQueryInterface2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// phr:  This MUST contain an array of HRESULTs specifying the respective success or
	// failure of each query operation.
	HResult []int32 `idl:"name:phr;size_is:(cIids)" json:"hresult"`
	// ppMIF:  This MUST contain an array of MInterfacePointer structures containing the
	// results of each query operation.
	Interface []*dcom.InterfacePointer `idl:"name:ppMIF;size_is:(cIids)" json:"interface"`
	// Return: The RemQueryInterface2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoteQueryInterface2Response) xxx_ToOp(ctx context.Context) *xxx_RemoteQueryInterface2Operation {
	if o == nil {
		return &xxx_RemoteQueryInterface2Operation{}
	}
	return &xxx_RemoteQueryInterface2Operation{
		That:      o.That,
		HResult:   o.HResult,
		Interface: o.Interface,
		Return:    o.Return,
	}
}

func (o *RemoteQueryInterface2Response) xxx_FromOp(ctx context.Context, op *xxx_RemoteQueryInterface2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.HResult = op.HResult
	o.Interface = op.Interface
	o.Return = op.Return
}
func (o *RemoteQueryInterface2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoteQueryInterface2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteQueryInterface2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
