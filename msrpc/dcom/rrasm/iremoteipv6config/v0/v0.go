package iremoteipv6config

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
	GoPackage = "dcom/rrasm"
)

var (
	// IRemoteIPV6Config interface identifier 6139d8a4-e508-4ebb-bac7-d7f275145897
	RemoteIPv6ConfigIID = &dcom.IID{Data1: 0x6139d8a4, Data2: 0xe508, Data3: 0x4ebb, Data4: []byte{0xba, 0xc7, 0xd7, 0xf2, 0x75, 0x14, 0x58, 0x97}}
	// Syntax UUID
	RemoteIPv6ConfigSyntaxUUID = &uuid.UUID{TimeLow: 0x6139d8a4, TimeMid: 0xe508, TimeHiAndVersion: 0x4ebb, ClockSeqHiAndReserved: 0xba, ClockSeqLow: 0xc7, Node: [6]uint8{0xd7, 0xf2, 0x75, 0x14, 0x58, 0x97}}
	// Syntax ID
	RemoteIPv6ConfigSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteIPv6ConfigSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRemoteIPV6Config interface.
type RemoteIPv6ConfigClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetAddressList operation.
	GetAddressList(context.Context, *GetAddressListRequest, ...dcerpc.CallOption) (*GetAddressListResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteIPv6ConfigClient
}

// IPv6Address structure represents IPV6Address RPC structure.
type IPv6Address struct {
	Bytes []byte `idl:"name:bytes" json:"bytes"`
}

func (o *IPv6Address) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPv6Address) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Bytes {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Bytes[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Bytes); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv6Address) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Bytes = make([]byte, 16)
	for i1 := range o.Bytes {
		i1 := i1
		if err := w.ReadData(&o.Bytes[i1]); err != nil {
			return err
		}
	}
	return nil
}

type xxx_DefaultRemoteIPv6ConfigClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteIPv6ConfigClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRemoteIPv6ConfigClient) GetAddressList(ctx context.Context, in *GetAddressListRequest, opts ...dcerpc.CallOption) (*GetAddressListResponse, error) {
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
	out := &GetAddressListResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteIPv6ConfigClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteIPv6ConfigClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRemoteIPv6ConfigClient) IPID(ctx context.Context, ipid *dcom.IPID) RemoteIPv6ConfigClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteIPv6ConfigClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewRemoteIPv6ConfigClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteIPv6ConfigClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteIPv6ConfigSyntaxV0_0))...)
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
	return &xxx_DefaultRemoteIPv6ConfigClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetAddressListOperation structure represents the GetAddressList operation
type xxx_GetAddressListOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	InterfaceName   string         `idl:"name:pszInterfaceName;string" json:"interface_name"`
	AddressesLength uint32         `idl:"name:pdwNumAddresses" json:"addresses_length"`
	IPv6AddressList []*IPv6Address `idl:"name:ppIPV6AddressList;size_is:(, pdwNumAddresses)" json:"ipv6_address_list"`
	InterfaceIndex  uint32         `idl:"name:dwIfIndex" json:"interface_index"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAddressListOperation) OpNum() int { return 3 }

func (o *xxx_GetAddressListOperation) OpName() string { return "/IRemoteIPV6Config/v0/GetAddressList" }

func (o *xxx_GetAddressListOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAddressListOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszInterfaceName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.InterfaceName); err != nil {
			return err
		}
	}
	// dwIfIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InterfaceIndex); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAddressListOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszInterfaceName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.InterfaceName); err != nil {
			return err
		}
	}
	// dwIfIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InterfaceIndex); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAddressListOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.IPv6AddressList != nil && o.AddressesLength == 0 {
		o.AddressesLength = uint32(len(o.IPv6AddressList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAddressListOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwNumAddresses {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AddressesLength); err != nil {
			return err
		}
	}
	// ppIPV6AddressList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IPV6Address}[dim:0,size_is=pdwNumAddresses](struct))
	{
		if o.IPv6AddressList != nil || o.AddressesLength > 0 {
			_ptr_ppIPV6AddressList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.AddressesLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.IPv6AddressList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.IPv6AddressList[i1] != nil {
						if err := o.IPv6AddressList[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.IPv6AddressList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.IPv6AddressList, _ptr_ppIPV6AddressList); err != nil {
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

func (o *xxx_GetAddressListOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwNumAddresses {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AddressesLength); err != nil {
			return err
		}
	}
	// ppIPV6AddressList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IPV6Address}[dim:0,size_is=pdwNumAddresses](struct))
	{
		_ptr_ppIPV6AddressList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.IPv6AddressList", sizeInfo[0])
			}
			o.IPv6AddressList = make([]*IPv6Address, sizeInfo[0])
			for i1 := range o.IPv6AddressList {
				i1 := i1
				if o.IPv6AddressList[i1] == nil {
					o.IPv6AddressList[i1] = &IPv6Address{}
				}
				if err := o.IPv6AddressList[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppIPV6AddressList := func(ptr interface{}) { o.IPv6AddressList = *ptr.(*[]*IPv6Address) }
		if err := w.ReadPointer(&o.IPv6AddressList, _s_ppIPV6AddressList, _ptr_ppIPV6AddressList); err != nil {
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

// GetAddressListRequest structure represents the GetAddressList operation request
type GetAddressListRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	InterfaceName  string         `idl:"name:pszInterfaceName;string" json:"interface_name"`
	InterfaceIndex uint32         `idl:"name:dwIfIndex" json:"interface_index"`
}

func (o *GetAddressListRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAddressListOperation) *xxx_GetAddressListOperation {
	if op == nil {
		op = &xxx_GetAddressListOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.InterfaceName = o.InterfaceName
	op.InterfaceIndex = o.InterfaceIndex
	return op
}

func (o *GetAddressListRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAddressListOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.InterfaceName = op.InterfaceName
	o.InterfaceIndex = op.InterfaceIndex
}
func (o *GetAddressListRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAddressListRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAddressListOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAddressListResponse structure represents the GetAddressList operation response
type GetAddressListResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	AddressesLength uint32         `idl:"name:pdwNumAddresses" json:"addresses_length"`
	IPv6AddressList []*IPv6Address `idl:"name:ppIPV6AddressList;size_is:(, pdwNumAddresses)" json:"ipv6_address_list"`
	// Return: The GetAddressList return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAddressListResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAddressListOperation) *xxx_GetAddressListOperation {
	if op == nil {
		op = &xxx_GetAddressListOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.AddressesLength = o.AddressesLength
	op.IPv6AddressList = o.IPv6AddressList
	op.Return = o.Return
	return op
}

func (o *GetAddressListResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAddressListOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AddressesLength = op.AddressesLength
	o.IPv6AddressList = op.IPv6AddressList
	o.Return = op.Return
}
func (o *GetAddressListResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAddressListResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAddressListOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
