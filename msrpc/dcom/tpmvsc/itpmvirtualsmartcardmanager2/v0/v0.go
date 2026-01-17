package itpmvirtualsmartcardmanager2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	tpmvsc "github.com/oiweiwei/go-msrpc/msrpc/dcom/tpmvsc"
	itpmvirtualsmartcardmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/tpmvsc/itpmvirtualsmartcardmanager/v0"
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
	_ = itpmvirtualsmartcardmanager.GoPackage
	_ = tpmvsc.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/tpmvsc"
)

var (
	// ITpmVirtualSmartCardManager2 interface identifier fdf8a2b9-02de-47f4-bc26-aa85ab5e5267
	VirtualSmartCardManager2IID = &dcom.IID{Data1: 0xfdf8a2b9, Data2: 0x02de, Data3: 0x47f4, Data4: []byte{0xbc, 0x26, 0xaa, 0x85, 0xab, 0x5e, 0x52, 0x67}}
	// Syntax UUID
	VirtualSmartCardManager2SyntaxUUID = &uuid.UUID{TimeLow: 0xfdf8a2b9, TimeMid: 0x2de, TimeHiAndVersion: 0x47f4, ClockSeqHiAndReserved: 0xbc, ClockSeqLow: 0x26, Node: [6]uint8{0xaa, 0x85, 0xab, 0x5e, 0x52, 0x67}}
	// Syntax ID
	VirtualSmartCardManager2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VirtualSmartCardManager2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ITpmVirtualSmartCardManager2 interface.
type VirtualSmartCardManager2Client interface {

	// ITpmVirtualSmartCardManager retrieval method.
	VirtualSmartCardManager() itpmvirtualsmartcardmanager.VirtualSmartCardManagerClient

	// This method is invoked by the requestor to create a VSC with the specified PIN policy
	// on the target.
	//
	// Return Values:  The server MUST return 0 if it successfully creates the new VSC,
	// and a nonzero value otherwise.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	CreateVirtualSmartCardWithPINPolicy(context.Context, *CreateVirtualSmartCardWithPINPolicyRequest, ...dcerpc.CallOption) (*CreateVirtualSmartCardWithPINPolicyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VirtualSmartCardManager2Client
}

type xxx_DefaultVirtualSmartCardManager2Client struct {
	itpmvirtualsmartcardmanager.VirtualSmartCardManagerClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVirtualSmartCardManager2Client) VirtualSmartCardManager() itpmvirtualsmartcardmanager.VirtualSmartCardManagerClient {
	return o.VirtualSmartCardManagerClient
}

func (o *xxx_DefaultVirtualSmartCardManager2Client) CreateVirtualSmartCardWithPINPolicy(ctx context.Context, in *CreateVirtualSmartCardWithPINPolicyRequest, opts ...dcerpc.CallOption) (*CreateVirtualSmartCardWithPINPolicyResponse, error) {
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
	out := &CreateVirtualSmartCardWithPINPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVirtualSmartCardManager2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVirtualSmartCardManager2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultVirtualSmartCardManager2Client) IPID(ctx context.Context, ipid *dcom.IPID) VirtualSmartCardManager2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVirtualSmartCardManager2Client{
		VirtualSmartCardManagerClient: o.VirtualSmartCardManagerClient.IPID(ctx, ipid),
		cc:                            o.cc,
		ipid:                          ipid,
	}
}

func NewVirtualSmartCardManager2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VirtualSmartCardManager2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VirtualSmartCardManager2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := itpmvirtualsmartcardmanager.NewVirtualSmartCardManagerClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultVirtualSmartCardManager2Client{
		VirtualSmartCardManagerClient: base,
		cc:                            cc,
		ipid:                          ipid,
	}, nil
}

// xxx_CreateVirtualSmartCardWithPINPolicyOperation structure represents the CreateVirtualSmartCardWithPinPolicy operation
type xxx_CreateVirtualSmartCardWithPINPolicyOperation struct {
	This             *dcom.ORPCThis                                `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat                                `idl:"name:That" json:"that"`
	FriendlyName     string                                        `idl:"name:pszFriendlyName;string" json:"friendly_name"`
	AdminAlgorithmID uint8                                         `idl:"name:bAdminAlgId" json:"admin_algorithm_id"`
	AdminKey         []byte                                        `idl:"name:pbAdminKey;size_is:(cbAdminKey)" json:"admin_key"`
	AdminKeyLength   uint32                                        `idl:"name:cbAdminKey" json:"admin_key_length"`
	AdminKCV         []byte                                        `idl:"name:pbAdminKcv;size_is:(cbAdminKcv);pointer:unique" json:"admin_kcv"`
	AdminKCVLength   uint32                                        `idl:"name:cbAdminKcv" json:"admin_kcv_length"`
	PUK              []byte                                        `idl:"name:pbPuk;size_is:(cbPuk);pointer:unique" json:"puk"`
	PUKLength        uint32                                        `idl:"name:cbPuk" json:"puk_length"`
	PIN              []byte                                        `idl:"name:pbPin;size_is:(cbPin)" json:"pin"`
	PINLength        uint32                                        `idl:"name:cbPin" json:"pin_length"`
	PINPolicy        []byte                                        `idl:"name:pbPinPolicy;size_is:(cbPinPolicy);pointer:unique" json:"pin_policy"`
	PINPolicyLength  uint32                                        `idl:"name:cbPinPolicy" json:"pin_policy_length"`
	Generate         int32                                         `idl:"name:fGenerate" json:"generate"`
	StatusCallback   *tpmvsc.VirtualSmartCardManagerStatusCallback `idl:"name:pStatusCallback;pointer:unique" json:"status_callback"`
	InstanceID       string                                        `idl:"name:ppszInstanceId;string" json:"instance_id"`
	NeedReboot       int32                                         `idl:"name:pfNeedReboot" json:"need_reboot"`
	Return           int32                                         `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateVirtualSmartCardWithPINPolicyOperation) OpNum() int { return 5 }

func (o *xxx_CreateVirtualSmartCardWithPINPolicyOperation) OpName() string {
	return "/ITpmVirtualSmartCardManager2/v0/CreateVirtualSmartCardWithPinPolicy"
}

func (o *xxx_CreateVirtualSmartCardWithPINPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.AdminKey != nil && o.AdminKeyLength == 0 {
		o.AdminKeyLength = uint32(len(o.AdminKey))
	}
	if o.AdminKCV != nil && o.AdminKCVLength == 0 {
		o.AdminKCVLength = uint32(len(o.AdminKCV))
	}
	if o.PUK != nil && o.PUKLength == 0 {
		o.PUKLength = uint32(len(o.PUK))
	}
	if o.PIN != nil && o.PINLength == 0 {
		o.PINLength = uint32(len(o.PIN))
	}
	if o.PINPolicy != nil && o.PINPolicyLength == 0 {
		o.PINPolicyLength = uint32(len(o.PINPolicy))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVirtualSmartCardWithPINPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszFriendlyName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FriendlyName); err != nil {
			return err
		}
	}
	// bAdminAlgId {in} (1:(uchar))
	{
		if err := w.WriteData(o.AdminAlgorithmID); err != nil {
			return err
		}
	}
	// pbAdminKey {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbAdminKey](uchar))
	{
		dimSize1 := uint64(o.AdminKeyLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.AdminKey {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AdminKey[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AdminKey); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbAdminKey {in} (1:(uint32))
	{
		if err := w.WriteData(o.AdminKeyLength); err != nil {
			return err
		}
	}
	// pbAdminKcv {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbAdminKcv](uchar))
	{
		if o.AdminKCV != nil || o.AdminKCVLength > 0 {
			_ptr_pbAdminKcv := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.AdminKCVLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.AdminKCV {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.AdminKCV[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.AdminKCV); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AdminKCV, _ptr_pbAdminKcv); err != nil {
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
	// cbAdminKcv {in} (1:(uint32))
	{
		if err := w.WriteData(o.AdminKCVLength); err != nil {
			return err
		}
	}
	// pbPuk {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbPuk](uchar))
	{
		if o.PUK != nil || o.PUKLength > 0 {
			_ptr_pbPuk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.PUKLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PUK {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.PUK[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.PUK); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PUK, _ptr_pbPuk); err != nil {
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
	// cbPuk {in} (1:(uint32))
	{
		if err := w.WriteData(o.PUKLength); err != nil {
			return err
		}
	}
	// pbPin {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbPin](uchar))
	{
		dimSize1 := uint64(o.PINLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.PIN {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.PIN[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.PIN); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbPin {in} (1:(uint32))
	{
		if err := w.WriteData(o.PINLength); err != nil {
			return err
		}
	}
	// pbPinPolicy {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbPinPolicy](uchar))
	{
		if o.PINPolicy != nil || o.PINPolicyLength > 0 {
			_ptr_pbPinPolicy := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.PINPolicyLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PINPolicy {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.PINPolicy[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.PINPolicy); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PINPolicy, _ptr_pbPinPolicy); err != nil {
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
	// cbPinPolicy {in} (1:(uint32))
	{
		if err := w.WriteData(o.PINPolicyLength); err != nil {
			return err
		}
	}
	// fGenerate {in} (1:(int32))
	{
		if err := w.WriteData(o.Generate); err != nil {
			return err
		}
	}
	// pStatusCallback {in} (1:{pointer=unique}*(1))(2:{alias=ITpmVirtualSmartCardManagerStatusCallback}(interface))
	{
		if o.StatusCallback != nil {
			_ptr_pStatusCallback := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.StatusCallback != nil {
					if err := o.StatusCallback.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&tpmvsc.VirtualSmartCardManagerStatusCallback{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.StatusCallback, _ptr_pStatusCallback); err != nil {
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
	return nil
}

func (o *xxx_CreateVirtualSmartCardWithPINPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszFriendlyName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FriendlyName); err != nil {
			return err
		}
	}
	// bAdminAlgId {in} (1:(uchar))
	{
		if err := w.ReadData(&o.AdminAlgorithmID); err != nil {
			return err
		}
	}
	// pbAdminKey {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbAdminKey](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.AdminKey", sizeInfo[0])
		}
		o.AdminKey = make([]byte, sizeInfo[0])
		for i1 := range o.AdminKey {
			i1 := i1
			if err := w.ReadData(&o.AdminKey[i1]); err != nil {
				return err
			}
		}
	}
	// cbAdminKey {in} (1:(uint32))
	{
		if err := w.ReadData(&o.AdminKeyLength); err != nil {
			return err
		}
	}
	// pbAdminKcv {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbAdminKcv](uchar))
	{
		_ptr_pbAdminKcv := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.AdminKCV", sizeInfo[0])
			}
			o.AdminKCV = make([]byte, sizeInfo[0])
			for i1 := range o.AdminKCV {
				i1 := i1
				if err := w.ReadData(&o.AdminKCV[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pbAdminKcv := func(ptr interface{}) { o.AdminKCV = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.AdminKCV, _s_pbAdminKcv, _ptr_pbAdminKcv); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbAdminKcv {in} (1:(uint32))
	{
		if err := w.ReadData(&o.AdminKCVLength); err != nil {
			return err
		}
	}
	// pbPuk {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbPuk](uchar))
	{
		_ptr_pbPuk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.PUK", sizeInfo[0])
			}
			o.PUK = make([]byte, sizeInfo[0])
			for i1 := range o.PUK {
				i1 := i1
				if err := w.ReadData(&o.PUK[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pbPuk := func(ptr interface{}) { o.PUK = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.PUK, _s_pbPuk, _ptr_pbPuk); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbPuk {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PUKLength); err != nil {
			return err
		}
	}
	// pbPin {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbPin](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.PIN", sizeInfo[0])
		}
		o.PIN = make([]byte, sizeInfo[0])
		for i1 := range o.PIN {
			i1 := i1
			if err := w.ReadData(&o.PIN[i1]); err != nil {
				return err
			}
		}
	}
	// cbPin {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PINLength); err != nil {
			return err
		}
	}
	// pbPinPolicy {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbPinPolicy](uchar))
	{
		_ptr_pbPinPolicy := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.PINPolicy", sizeInfo[0])
			}
			o.PINPolicy = make([]byte, sizeInfo[0])
			for i1 := range o.PINPolicy {
				i1 := i1
				if err := w.ReadData(&o.PINPolicy[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pbPinPolicy := func(ptr interface{}) { o.PINPolicy = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.PINPolicy, _s_pbPinPolicy, _ptr_pbPinPolicy); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbPinPolicy {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PINPolicyLength); err != nil {
			return err
		}
	}
	// fGenerate {in} (1:(int32))
	{
		if err := w.ReadData(&o.Generate); err != nil {
			return err
		}
	}
	// pStatusCallback {in} (1:{pointer=unique}*(1))(2:{alias=ITpmVirtualSmartCardManagerStatusCallback}(interface))
	{
		_ptr_pStatusCallback := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.StatusCallback == nil {
				o.StatusCallback = &tpmvsc.VirtualSmartCardManagerStatusCallback{}
			}
			if err := o.StatusCallback.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pStatusCallback := func(ptr interface{}) { o.StatusCallback = *ptr.(**tpmvsc.VirtualSmartCardManagerStatusCallback) }
		if err := w.ReadPointer(&o.StatusCallback, _s_pStatusCallback, _ptr_pStatusCallback); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVirtualSmartCardWithPINPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVirtualSmartCardWithPINPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppszInstanceId {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.InstanceID != "" {
			_ptr_ppszInstanceId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.InstanceID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.InstanceID, _ptr_ppszInstanceId); err != nil {
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
	// pfNeedReboot {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.NeedReboot); err != nil {
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

func (o *xxx_CreateVirtualSmartCardWithPINPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppszInstanceId {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_ppszInstanceId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.InstanceID); err != nil {
				return err
			}
			return nil
		})
		_s_ppszInstanceId := func(ptr interface{}) { o.InstanceID = *ptr.(*string) }
		if err := w.ReadPointer(&o.InstanceID, _s_ppszInstanceId, _ptr_ppszInstanceId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pfNeedReboot {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.NeedReboot); err != nil {
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

// CreateVirtualSmartCardWithPINPolicyRequest structure represents the CreateVirtualSmartCardWithPinPolicy operation request
type CreateVirtualSmartCardWithPINPolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pszFriendlyName:  A Unicode string for use in any user interface messages relating
	// to this VSC.
	FriendlyName string `idl:"name:pszFriendlyName;string" json:"friendly_name"`
	// bAdminAlgId:  An unsigned byte value. This parameter MUST be set to 0x82.
	AdminAlgorithmID uint8 `idl:"name:bAdminAlgId" json:"admin_algorithm_id"`
	// pbAdminKey:  An array of 24 bytes containing a TDEA [SP800-67] key intended to be
	// used as the administrative key for the new VSC.
	AdminKey []byte `idl:"name:pbAdminKey;size_is:(cbAdminKey)" json:"admin_key"`
	// cbAdminKey:  A 32-bit unsigned integer value. It MUST be set to 24.
	AdminKeyLength uint32 `idl:"name:cbAdminKey" json:"admin_key_length"`
	// pbAdminKcv:  An array of bytes containing the Key Check Value (KCV) for the administrative
	// key contained in the pbAdminKey parameter. This parameter is optional and MUST be
	// set to NULL if absent. If present, it MUST be computed by encrypting eight zero bytes
	// using the TDEA [SP800-67] block cipher and taking the first three bytes.
	AdminKCV []byte `idl:"name:pbAdminKcv;size_is:(cbAdminKcv);pointer:unique" json:"admin_kcv"`
	// cbAdminKcv:  A 32-bit unsigned integer value. It MUST be set to 0 if the pbAdmin
	// parameter is NULL, and MUST be set to 3 otherwise.
	AdminKCVLength uint32 `idl:"name:cbAdminKcv" json:"admin_kcv_length"`
	// pbPuk:  An array of bytes containing the desired PUK for the new VSC. This parameter
	// is optional and MUST be set to NULL if absent. If present, its length MUST be between
	// 8 and 127 bytes, inclusive.
	PUK []byte `idl:"name:pbPuk;size_is:(cbPuk);pointer:unique" json:"puk"`
	// cbPuk:  A 32-bit unsigned integer value. It MUST be equal to the length of the pbPuk
	// parameter in bytes. If pbPuk is NULL, this parameter MUST be set to 0.
	PUKLength uint32 `idl:"name:cbPuk" json:"puk_length"`
	// pbPin:  An array of bytes containing the desired PIN for the new VSC. Its length
	// MUST be between 4 and 127 bytes, inclusive.
	PIN []byte `idl:"name:pbPin;size_is:(cbPin)" json:"pin"`
	// cbPin: A 32-bit unsigned integer value. It MUST be equal to the length of the pbPin
	// parameter in bytes.
	PINLength uint32 `idl:"name:cbPin" json:"pin_length"`
	// pbPinPolicy:  A PinPolicySerialization structure specifying the PIN policy for the
	// new VSC, as described in section 2.2.2.1.
	PINPolicy []byte `idl:"name:pbPinPolicy;size_is:(cbPinPolicy);pointer:unique" json:"pin_policy"`
	// cbPinPolicy:  A 32-bit unsigned integer value. It MUST be equal to the length in
	// bytes of the pbPinPolicy parameter.
	PINPolicyLength uint32 `idl:"name:cbPinPolicy" json:"pin_policy_length"`
	// fGenerate:  A Boolean value that indicates whether a file system is to be generated
	// on the new VSC.
	Generate int32 `idl:"name:fGenerate" json:"generate"`
	// pStatusCallback:  A reference to an instance of the ITpmVirtualSmartCardManagerStatusCallback
	// DCOM interface on the requestor. The server uses this interface to provide feedback
	// on progress and errors. This parameter is optional and MUST be set to NULL if absent.
	StatusCallback *tpmvsc.VirtualSmartCardManagerStatusCallback `idl:"name:pStatusCallback;pointer:unique" json:"status_callback"`
}

func (o *CreateVirtualSmartCardWithPINPolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateVirtualSmartCardWithPINPolicyOperation) *xxx_CreateVirtualSmartCardWithPINPolicyOperation {
	if op == nil {
		op = &xxx_CreateVirtualSmartCardWithPINPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.FriendlyName = o.FriendlyName
	op.AdminAlgorithmID = o.AdminAlgorithmID
	op.AdminKey = o.AdminKey
	op.AdminKeyLength = o.AdminKeyLength
	op.AdminKCV = o.AdminKCV
	op.AdminKCVLength = o.AdminKCVLength
	op.PUK = o.PUK
	op.PUKLength = o.PUKLength
	op.PIN = o.PIN
	op.PINLength = o.PINLength
	op.PINPolicy = o.PINPolicy
	op.PINPolicyLength = o.PINPolicyLength
	op.Generate = o.Generate
	op.StatusCallback = o.StatusCallback
	return op
}

func (o *CreateVirtualSmartCardWithPINPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateVirtualSmartCardWithPINPolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FriendlyName = op.FriendlyName
	o.AdminAlgorithmID = op.AdminAlgorithmID
	o.AdminKey = op.AdminKey
	o.AdminKeyLength = op.AdminKeyLength
	o.AdminKCV = op.AdminKCV
	o.AdminKCVLength = op.AdminKCVLength
	o.PUK = op.PUK
	o.PUKLength = op.PUKLength
	o.PIN = op.PIN
	o.PINLength = op.PINLength
	o.PINPolicy = op.PINPolicy
	o.PINPolicyLength = op.PINPolicyLength
	o.Generate = op.Generate
	o.StatusCallback = op.StatusCallback
}
func (o *CreateVirtualSmartCardWithPINPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateVirtualSmartCardWithPINPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVirtualSmartCardWithPINPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateVirtualSmartCardWithPINPolicyResponse structure represents the CreateVirtualSmartCardWithPinPolicy operation response
type CreateVirtualSmartCardWithPINPolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppszInstanceId:  A Unicode string containing a unique instance identifier for the
	// VSC created by this operation.
	InstanceID string `idl:"name:ppszInstanceId;string" json:"instance_id"`
	// pfNeedReboot:  A Boolean value that indicates whether or not a reboot is required
	// on the server before the newly-created VSC is made available to applications.
	NeedReboot int32 `idl:"name:pfNeedReboot" json:"need_reboot"`
	// Return: The CreateVirtualSmartCardWithPinPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateVirtualSmartCardWithPINPolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateVirtualSmartCardWithPINPolicyOperation) *xxx_CreateVirtualSmartCardWithPINPolicyOperation {
	if op == nil {
		op = &xxx_CreateVirtualSmartCardWithPINPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.InstanceID = o.InstanceID
	op.NeedReboot = o.NeedReboot
	op.Return = o.Return
	return op
}

func (o *CreateVirtualSmartCardWithPINPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateVirtualSmartCardWithPINPolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.InstanceID = op.InstanceID
	o.NeedReboot = op.NeedReboot
	o.Return = op.Return
}
func (o *CreateVirtualSmartCardWithPINPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateVirtualSmartCardWithPINPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVirtualSmartCardWithPINPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
