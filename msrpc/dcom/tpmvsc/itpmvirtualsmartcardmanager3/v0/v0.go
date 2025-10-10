package itpmvirtualsmartcardmanager3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	tpmvsc "github.com/oiweiwei/go-msrpc/msrpc/dcom/tpmvsc"
	itpmvirtualsmartcardmanager2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/tpmvsc/itpmvirtualsmartcardmanager2/v0"
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
	_ = itpmvirtualsmartcardmanager2.GoPackage
	_ = tpmvsc.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/tpmvsc"
)

var (
	// ITpmVirtualSmartCardManager3 interface identifier 3c745a97-f375-4150-be17-5950f694c699
	TpmVirtualSmartCardManager3IID = &dcom.IID{Data1: 0x3c745a97, Data2: 0xf375, Data3: 0x4150, Data4: []byte{0xbe, 0x17, 0x59, 0x50, 0xf6, 0x94, 0xc6, 0x99}}
	// Syntax UUID
	TpmVirtualSmartCardManager3SyntaxUUID = &uuid.UUID{TimeLow: 0x3c745a97, TimeMid: 0xf375, TimeHiAndVersion: 0x4150, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0x17, Node: [6]uint8{0x59, 0x50, 0xf6, 0x94, 0xc6, 0x99}}
	// Syntax ID
	TpmVirtualSmartCardManager3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: TpmVirtualSmartCardManager3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ITpmVirtualSmartCardManager3 interface.
type TpmVirtualSmartCardManager3Client interface {

	// ITpmVirtualSmartCardManager2 retrieval method.
	TpmVirtualSmartCardManager2() itpmvirtualsmartcardmanager2.TpmVirtualSmartCardManager2Client

	CreateVirtualSmartCardWithAttestation(context.Context, *CreateVirtualSmartCardWithAttestationRequest, ...dcerpc.CallOption) (*CreateVirtualSmartCardWithAttestationResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) TpmVirtualSmartCardManager3Client
}

type xxx_DefaultTpmVirtualSmartCardManager3Client struct {
	itpmvirtualsmartcardmanager2.TpmVirtualSmartCardManager2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultTpmVirtualSmartCardManager3Client) TpmVirtualSmartCardManager2() itpmvirtualsmartcardmanager2.TpmVirtualSmartCardManager2Client {
	return o.TpmVirtualSmartCardManager2Client
}

func (o *xxx_DefaultTpmVirtualSmartCardManager3Client) CreateVirtualSmartCardWithAttestation(ctx context.Context, in *CreateVirtualSmartCardWithAttestationRequest, opts ...dcerpc.CallOption) (*CreateVirtualSmartCardWithAttestationResponse, error) {
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
	out := &CreateVirtualSmartCardWithAttestationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTpmVirtualSmartCardManager3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTpmVirtualSmartCardManager3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultTpmVirtualSmartCardManager3Client) IPID(ctx context.Context, ipid *dcom.IPID) TpmVirtualSmartCardManager3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultTpmVirtualSmartCardManager3Client{
		TpmVirtualSmartCardManager2Client: o.TpmVirtualSmartCardManager2Client.IPID(ctx, ipid),
		cc:                                o.cc,
		ipid:                              ipid,
	}
}

func NewTpmVirtualSmartCardManager3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TpmVirtualSmartCardManager3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TpmVirtualSmartCardManager3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := itpmvirtualsmartcardmanager2.NewTpmVirtualSmartCardManager2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultTpmVirtualSmartCardManager3Client{
		TpmVirtualSmartCardManager2Client: base,
		cc:                                cc,
		ipid:                              ipid,
	}, nil
}

// xxx_CreateVirtualSmartCardWithAttestationOperation structure represents the CreateVirtualSmartCardWithAttestation operation
type xxx_CreateVirtualSmartCardWithAttestationOperation struct {
	This             *dcom.ORPCThis                                   `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat                                   `idl:"name:That" json:"that"`
	FriendlyName     string                                           `idl:"name:pszFriendlyName;string" json:"friendly_name"`
	AdminAlgorithmID uint8                                            `idl:"name:bAdminAlgId" json:"admin_algorithm_id"`
	AdminKey         []byte                                           `idl:"name:pbAdminKey;size_is:(cbAdminKey)" json:"admin_key"`
	AdminKeyLength   uint32                                           `idl:"name:cbAdminKey" json:"admin_key_length"`
	AdminKcv         []byte                                           `idl:"name:pbAdminKcv;size_is:(cbAdminKcv);pointer:unique" json:"admin_kcv"`
	AdminKcvLength   uint32                                           `idl:"name:cbAdminKcv" json:"admin_kcv_length"`
	Puk              []byte                                           `idl:"name:pbPuk;size_is:(cbPuk);pointer:unique" json:"puk"`
	PukLength        uint32                                           `idl:"name:cbPuk" json:"puk_length"`
	Pin              []byte                                           `idl:"name:pbPin;size_is:(cbPin)" json:"pin"`
	PinLength        uint32                                           `idl:"name:cbPin" json:"pin_length"`
	PinPolicy        []byte                                           `idl:"name:pbPinPolicy;size_is:(cbPinPolicy);pointer:unique" json:"pin_policy"`
	PinPolicyLength  uint32                                           `idl:"name:cbPinPolicy" json:"pin_policy_length"`
	AttestationType  tpmvsc.TpmvscAttestationType                     `idl:"name:attestationType" json:"attestation_type"`
	Generate         int32                                            `idl:"name:fGenerate" json:"generate"`
	StatusCallback   *tpmvsc.TpmVirtualSmartCardManagerStatusCallback `idl:"name:pStatusCallback;pointer:unique" json:"status_callback"`
	InstanceID       string                                           `idl:"name:ppszInstanceId;string" json:"instance_id"`
	Return           int32                                            `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateVirtualSmartCardWithAttestationOperation) OpNum() int { return 6 }

func (o *xxx_CreateVirtualSmartCardWithAttestationOperation) OpName() string {
	return "/ITpmVirtualSmartCardManager3/v0/CreateVirtualSmartCardWithAttestation"
}

func (o *xxx_CreateVirtualSmartCardWithAttestationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.AdminKey != nil && o.AdminKeyLength == 0 {
		o.AdminKeyLength = uint32(len(o.AdminKey))
	}
	if o.AdminKcv != nil && o.AdminKcvLength == 0 {
		o.AdminKcvLength = uint32(len(o.AdminKcv))
	}
	if o.Puk != nil && o.PukLength == 0 {
		o.PukLength = uint32(len(o.Puk))
	}
	if o.Pin != nil && o.PinLength == 0 {
		o.PinLength = uint32(len(o.Pin))
	}
	if o.PinPolicy != nil && o.PinPolicyLength == 0 {
		o.PinPolicyLength = uint32(len(o.PinPolicy))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVirtualSmartCardWithAttestationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if o.AdminKcv != nil || o.AdminKcvLength > 0 {
			_ptr_pbAdminKcv := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.AdminKcvLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.AdminKcv {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.AdminKcv[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.AdminKcv); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AdminKcv, _ptr_pbAdminKcv); err != nil {
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
		if err := w.WriteData(o.AdminKcvLength); err != nil {
			return err
		}
	}
	// pbPuk {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbPuk](uchar))
	{
		if o.Puk != nil || o.PukLength > 0 {
			_ptr_pbPuk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.PukLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Puk {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Puk[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Puk); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Puk, _ptr_pbPuk); err != nil {
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
		if err := w.WriteData(o.PukLength); err != nil {
			return err
		}
	}
	// pbPin {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbPin](uchar))
	{
		dimSize1 := uint64(o.PinLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Pin {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Pin[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Pin); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbPin {in} (1:(uint32))
	{
		if err := w.WriteData(o.PinLength); err != nil {
			return err
		}
	}
	// pbPinPolicy {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbPinPolicy](uchar))
	{
		if o.PinPolicy != nil || o.PinPolicyLength > 0 {
			_ptr_pbPinPolicy := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.PinPolicyLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PinPolicy {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.PinPolicy[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.PinPolicy); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PinPolicy, _ptr_pbPinPolicy); err != nil {
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
		if err := w.WriteData(o.PinPolicyLength); err != nil {
			return err
		}
	}
	// attestationType {in} (1:{v1_enum, alias=TPMVSC_ATTESTATION_TYPE}(enum))
	{
		if err := w.WriteEnum(uint32(o.AttestationType)); err != nil {
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
					if err := (&tpmvsc.TpmVirtualSmartCardManagerStatusCallback{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CreateVirtualSmartCardWithAttestationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
				return fmt.Errorf("buffer overflow for size %d of array o.AdminKcv", sizeInfo[0])
			}
			o.AdminKcv = make([]byte, sizeInfo[0])
			for i1 := range o.AdminKcv {
				i1 := i1
				if err := w.ReadData(&o.AdminKcv[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pbAdminKcv := func(ptr interface{}) { o.AdminKcv = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.AdminKcv, _s_pbAdminKcv, _ptr_pbAdminKcv); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbAdminKcv {in} (1:(uint32))
	{
		if err := w.ReadData(&o.AdminKcvLength); err != nil {
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
				return fmt.Errorf("buffer overflow for size %d of array o.Puk", sizeInfo[0])
			}
			o.Puk = make([]byte, sizeInfo[0])
			for i1 := range o.Puk {
				i1 := i1
				if err := w.ReadData(&o.Puk[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pbPuk := func(ptr interface{}) { o.Puk = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Puk, _s_pbPuk, _ptr_pbPuk); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbPuk {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PukLength); err != nil {
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
			return fmt.Errorf("buffer overflow for size %d of array o.Pin", sizeInfo[0])
		}
		o.Pin = make([]byte, sizeInfo[0])
		for i1 := range o.Pin {
			i1 := i1
			if err := w.ReadData(&o.Pin[i1]); err != nil {
				return err
			}
		}
	}
	// cbPin {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PinLength); err != nil {
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
				return fmt.Errorf("buffer overflow for size %d of array o.PinPolicy", sizeInfo[0])
			}
			o.PinPolicy = make([]byte, sizeInfo[0])
			for i1 := range o.PinPolicy {
				i1 := i1
				if err := w.ReadData(&o.PinPolicy[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pbPinPolicy := func(ptr interface{}) { o.PinPolicy = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.PinPolicy, _s_pbPinPolicy, _ptr_pbPinPolicy); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbPinPolicy {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PinPolicyLength); err != nil {
			return err
		}
	}
	// attestationType {in} (1:{v1_enum, alias=TPMVSC_ATTESTATION_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.AttestationType)); err != nil {
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
				o.StatusCallback = &tpmvsc.TpmVirtualSmartCardManagerStatusCallback{}
			}
			if err := o.StatusCallback.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pStatusCallback := func(ptr interface{}) { o.StatusCallback = *ptr.(**tpmvsc.TpmVirtualSmartCardManagerStatusCallback) }
		if err := w.ReadPointer(&o.StatusCallback, _s_pStatusCallback, _ptr_pStatusCallback); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVirtualSmartCardWithAttestationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVirtualSmartCardWithAttestationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVirtualSmartCardWithAttestationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateVirtualSmartCardWithAttestationRequest structure represents the CreateVirtualSmartCardWithAttestation operation request
type CreateVirtualSmartCardWithAttestationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This             *dcom.ORPCThis                                   `idl:"name:This" json:"this"`
	FriendlyName     string                                           `idl:"name:pszFriendlyName;string" json:"friendly_name"`
	AdminAlgorithmID uint8                                            `idl:"name:bAdminAlgId" json:"admin_algorithm_id"`
	AdminKey         []byte                                           `idl:"name:pbAdminKey;size_is:(cbAdminKey)" json:"admin_key"`
	AdminKeyLength   uint32                                           `idl:"name:cbAdminKey" json:"admin_key_length"`
	AdminKcv         []byte                                           `idl:"name:pbAdminKcv;size_is:(cbAdminKcv);pointer:unique" json:"admin_kcv"`
	AdminKcvLength   uint32                                           `idl:"name:cbAdminKcv" json:"admin_kcv_length"`
	Puk              []byte                                           `idl:"name:pbPuk;size_is:(cbPuk);pointer:unique" json:"puk"`
	PukLength        uint32                                           `idl:"name:cbPuk" json:"puk_length"`
	Pin              []byte                                           `idl:"name:pbPin;size_is:(cbPin)" json:"pin"`
	PinLength        uint32                                           `idl:"name:cbPin" json:"pin_length"`
	PinPolicy        []byte                                           `idl:"name:pbPinPolicy;size_is:(cbPinPolicy);pointer:unique" json:"pin_policy"`
	PinPolicyLength  uint32                                           `idl:"name:cbPinPolicy" json:"pin_policy_length"`
	AttestationType  tpmvsc.TpmvscAttestationType                     `idl:"name:attestationType" json:"attestation_type"`
	Generate         int32                                            `idl:"name:fGenerate" json:"generate"`
	StatusCallback   *tpmvsc.TpmVirtualSmartCardManagerStatusCallback `idl:"name:pStatusCallback;pointer:unique" json:"status_callback"`
}

func (o *CreateVirtualSmartCardWithAttestationRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateVirtualSmartCardWithAttestationOperation) *xxx_CreateVirtualSmartCardWithAttestationOperation {
	if op == nil {
		op = &xxx_CreateVirtualSmartCardWithAttestationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.FriendlyName = o.FriendlyName
	op.AdminAlgorithmID = o.AdminAlgorithmID
	op.AdminKey = o.AdminKey
	op.AdminKeyLength = o.AdminKeyLength
	op.AdminKcv = o.AdminKcv
	op.AdminKcvLength = o.AdminKcvLength
	op.Puk = o.Puk
	op.PukLength = o.PukLength
	op.Pin = o.Pin
	op.PinLength = o.PinLength
	op.PinPolicy = o.PinPolicy
	op.PinPolicyLength = o.PinPolicyLength
	op.AttestationType = o.AttestationType
	op.Generate = o.Generate
	op.StatusCallback = o.StatusCallback
	return op
}

func (o *CreateVirtualSmartCardWithAttestationRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateVirtualSmartCardWithAttestationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FriendlyName = op.FriendlyName
	o.AdminAlgorithmID = op.AdminAlgorithmID
	o.AdminKey = op.AdminKey
	o.AdminKeyLength = op.AdminKeyLength
	o.AdminKcv = op.AdminKcv
	o.AdminKcvLength = op.AdminKcvLength
	o.Puk = op.Puk
	o.PukLength = op.PukLength
	o.Pin = op.Pin
	o.PinLength = op.PinLength
	o.PinPolicy = op.PinPolicy
	o.PinPolicyLength = op.PinPolicyLength
	o.AttestationType = op.AttestationType
	o.Generate = op.Generate
	o.StatusCallback = op.StatusCallback
}
func (o *CreateVirtualSmartCardWithAttestationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateVirtualSmartCardWithAttestationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVirtualSmartCardWithAttestationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateVirtualSmartCardWithAttestationResponse structure represents the CreateVirtualSmartCardWithAttestation operation response
type CreateVirtualSmartCardWithAttestationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	InstanceID string         `idl:"name:ppszInstanceId;string" json:"instance_id"`
	// Return: The CreateVirtualSmartCardWithAttestation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateVirtualSmartCardWithAttestationResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateVirtualSmartCardWithAttestationOperation) *xxx_CreateVirtualSmartCardWithAttestationOperation {
	if op == nil {
		op = &xxx_CreateVirtualSmartCardWithAttestationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.InstanceID = o.InstanceID
	op.Return = o.Return
	return op
}

func (o *CreateVirtualSmartCardWithAttestationResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateVirtualSmartCardWithAttestationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.InstanceID = op.InstanceID
	o.Return = op.Return
}
func (o *CreateVirtualSmartCardWithAttestationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateVirtualSmartCardWithAttestationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVirtualSmartCardWithAttestationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
