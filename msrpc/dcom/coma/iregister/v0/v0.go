package iregister

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
	// IRegister interface identifier 8db2180e-bd29-11d1-8b7e-00c04fd7a924
	RegisterIID = &dcom.IID{Data1: 0x8db2180e, Data2: 0xbd29, Data3: 0x11d1, Data4: []byte{0x8b, 0x7e, 0x00, 0xc0, 0x4f, 0xd7, 0xa9, 0x24}}
	// Syntax UUID
	RegisterSyntaxUUID = &uuid.UUID{TimeLow: 0x8db2180e, TimeMid: 0xbd29, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8b, ClockSeqLow: 0x7e, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd7, 0xa9, 0x24}}
	// Syntax ID
	RegisterSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RegisterSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRegister interface.
type RegisterClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to register the components in one or more modules
	// and to create component full configurations for those modules in an existing conglomeration.
	// This method supports conglomerations in the global partition only.
	//
	// Alternatively, this method can be called to verify modules without actually registering
	// the components. As a side effect, this method returns implementation-specific detailed
	// results of the registration or verification operation for informational purposes.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	RegisterModule(context.Context, *RegisterModuleRequest, ...dcerpc.CallOption) (*RegisterModuleResponse, error)

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RegisterClient
}

type xxx_DefaultRegisterClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRegisterClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRegisterClient) RegisterModule(ctx context.Context, in *RegisterModuleRequest, opts ...dcerpc.CallOption) (*RegisterModuleResponse, error) {
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
	out := &RegisterModuleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRegisterClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRegisterClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRegisterClient) IPID(ctx context.Context, ipid *dcom.IPID) RegisterClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRegisterClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewRegisterClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RegisterClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RegisterSyntaxV0_0))...)
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
	return &xxx_DefaultRegisterClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_RegisterModuleOperation structure represents the RegisterModule operation
type xxx_RegisterModuleOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConglomerationID  *dtyp.GUID     `idl:"name:ConglomerationIdentifier" json:"conglomeration_id"`
	Modules           []string       `idl:"name:ppModules;size_is:(cModules, );string" json:"modules"`
	ModulesCount      uint32         `idl:"name:cModules" json:"modules_count"`
	Flags             uint32         `idl:"name:dwFlags" json:"flags"`
	RequestedClassIDs []*dtyp.GUID   `idl:"name:pRequestedCLSIDs;size_is:(cRequested);pointer:unique" json:"requested_class_ids"`
	RequestedCount    uint32         `idl:"name:cRequested" json:"requested_count"`
	ModuleFlags       []uint32       `idl:"name:ppModuleFlags;size_is:(, cModules)" json:"module_flags"`
	ResultsCount      uint32         `idl:"name:pcResults" json:"results_count"`
	ResultClassIDs    []*dtyp.GUID   `idl:"name:ppResultCLSIDs;size_is:(, pcResults)" json:"result_class_ids"`
	ResultNames       []string       `idl:"name:ppResultNames;size_is:(, pcResults);string" json:"result_names"`
	ResultFlags       []uint32       `idl:"name:ppResultFlags;size_is:(, pcResults)" json:"result_flags"`
	ResultHRs         []int32        `idl:"name:ppResultHRs;size_is:(, pcResults)" json:"result_h_rs"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterModuleOperation) OpNum() int { return 3 }

func (o *xxx_RegisterModuleOperation) OpName() string { return "/IRegister/v0/RegisterModule" }

func (o *xxx_RegisterModuleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Modules != nil && o.ModulesCount == 0 {
		o.ModulesCount = uint32(len(o.Modules))
	}
	if o.RequestedClassIDs != nil && o.RequestedCount == 0 {
		o.RequestedCount = uint32(len(o.RequestedClassIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterModuleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ConglomerationIdentifier {in} (1:{alias=GUID}(struct))
	{
		if o.ConglomerationID != nil {
			if err := o.ConglomerationID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ppModules {in} (1:{string, pointer=ref}*(1))(2:{alias=LPWSTR}[dim:0,size_is=cModules]*(1)[dim:0,string,null](wchar))
	{
		dimSize1 := uint64(o.ModulesCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Modules {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Modules[i1] != "" {
				_ptr_ppModules := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if err := ndr.WriteUTF16NString(ctx, w, o.Modules[i1]); err != nil {
						return err
					}
					return nil
				})
				if err := w.WritePointer(&o.Modules[i1], _ptr_ppModules); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Modules); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cModules {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ModulesCount); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pRequestedCLSIDs {in} (1:{pointer=unique}*(1))(2:{alias=GUID}[dim:0,size_is=cRequested](struct))
	{
		if o.RequestedClassIDs != nil || o.RequestedCount > 0 {
			_ptr_pRequestedCLSIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.RequestedCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.RequestedClassIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.RequestedClassIDs[i1] != nil {
						if err := o.RequestedClassIDs[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.RequestedClassIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RequestedClassIDs, _ptr_pRequestedCLSIDs); err != nil {
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
	// cRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterModuleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ConglomerationIdentifier {in} (1:{alias=GUID}(struct))
	{
		if o.ConglomerationID == nil {
			o.ConglomerationID = &dtyp.GUID{}
		}
		if err := o.ConglomerationID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ppModules {in} (1:{string, pointer=ref}*(1))(2:{alias=LPWSTR}[dim:0,size_is=cModules]*(1)[dim:0,string,null](wchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Modules", sizeInfo[0])
		}
		o.Modules = make([]string, sizeInfo[0])
		for i1 := range o.Modules {
			i1 := i1
			_ptr_ppModules := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.Modules[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_ppModules := func(ptr interface{}) { o.Modules[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.Modules[i1], _s_ppModules, _ptr_ppModules); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cModules {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ModulesCount); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pRequestedCLSIDs {in} (1:{pointer=unique}*(1))(2:{alias=GUID}[dim:0,size_is=cRequested](struct))
	{
		_ptr_pRequestedCLSIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.RequestedClassIDs", sizeInfo[0])
			}
			o.RequestedClassIDs = make([]*dtyp.GUID, sizeInfo[0])
			for i1 := range o.RequestedClassIDs {
				i1 := i1
				if o.RequestedClassIDs[i1] == nil {
					o.RequestedClassIDs[i1] = &dtyp.GUID{}
				}
				if err := o.RequestedClassIDs[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pRequestedCLSIDs := func(ptr interface{}) { o.RequestedClassIDs = *ptr.(*[]*dtyp.GUID) }
		if err := w.ReadPointer(&o.RequestedClassIDs, _s_pRequestedCLSIDs, _ptr_pRequestedCLSIDs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterModuleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ResultClassIDs != nil && o.ResultsCount == 0 {
		o.ResultsCount = uint32(len(o.ResultClassIDs))
	}
	if o.ResultNames != nil && o.ResultsCount == 0 {
		o.ResultsCount = uint32(len(o.ResultNames))
	}
	if o.ResultFlags != nil && o.ResultsCount == 0 {
		o.ResultsCount = uint32(len(o.ResultFlags))
	}
	if o.ResultHRs != nil && o.ResultsCount == 0 {
		o.ResultsCount = uint32(len(o.ResultHRs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterModuleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppModuleFlags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=cModules](uint32))
	{
		if o.ModuleFlags != nil || o.ModulesCount > 0 {
			_ptr_ppModuleFlags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ModulesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ModuleFlags {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ModuleFlags[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ModuleFlags); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModuleFlags, _ptr_ppModuleFlags); err != nil {
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
	// pcResults {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ResultsCount); err != nil {
			return err
		}
	}
	// ppResultCLSIDs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pcResults](struct))
	{
		if o.ResultClassIDs != nil || o.ResultsCount > 0 {
			_ptr_ppResultCLSIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ResultsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ResultClassIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ResultClassIDs[i1] != nil {
						if err := o.ResultClassIDs[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ResultClassIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultClassIDs, _ptr_ppResultCLSIDs); err != nil {
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
	// ppResultNames {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcResults]*(1)[dim:0,string,null](wchar))
	{
		if o.ResultNames != nil || o.ResultsCount > 0 {
			_ptr_ppResultNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ResultsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ResultNames {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ResultNames[i1] != "" {
						_ptr_ppResultNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.ResultNames[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.ResultNames[i1], _ptr_ppResultNames); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ResultNames); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultNames, _ptr_ppResultNames); err != nil {
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
	// ppResultFlags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=pcResults](uint32))
	{
		if o.ResultFlags != nil || o.ResultsCount > 0 {
			_ptr_ppResultFlags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ResultsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ResultFlags {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ResultFlags[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ResultFlags); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultFlags, _ptr_ppResultFlags); err != nil {
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
	// ppResultHRs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LONG}[dim:0,size_is=pcResults](int32))
	{
		if o.ResultHRs != nil || o.ResultsCount > 0 {
			_ptr_ppResultHRs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ResultsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ResultHRs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ResultHRs[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ResultHRs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(int32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultHRs, _ptr_ppResultHRs); err != nil {
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

func (o *xxx_RegisterModuleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppModuleFlags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=cModules](uint32))
	{
		_ptr_ppModuleFlags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ModuleFlags", sizeInfo[0])
			}
			o.ModuleFlags = make([]uint32, sizeInfo[0])
			for i1 := range o.ModuleFlags {
				i1 := i1
				if err := w.ReadData(&o.ModuleFlags[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppModuleFlags := func(ptr interface{}) { o.ModuleFlags = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.ModuleFlags, _s_ppModuleFlags, _ptr_ppModuleFlags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcResults {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ResultsCount); err != nil {
			return err
		}
	}
	// ppResultCLSIDs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pcResults](struct))
	{
		_ptr_ppResultCLSIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ResultClassIDs", sizeInfo[0])
			}
			o.ResultClassIDs = make([]*dtyp.GUID, sizeInfo[0])
			for i1 := range o.ResultClassIDs {
				i1 := i1
				if o.ResultClassIDs[i1] == nil {
					o.ResultClassIDs[i1] = &dtyp.GUID{}
				}
				if err := o.ResultClassIDs[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppResultCLSIDs := func(ptr interface{}) { o.ResultClassIDs = *ptr.(*[]*dtyp.GUID) }
		if err := w.ReadPointer(&o.ResultClassIDs, _s_ppResultCLSIDs, _ptr_ppResultCLSIDs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppResultNames {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcResults]*(1)[dim:0,string,null](wchar))
	{
		_ptr_ppResultNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ResultNames", sizeInfo[0])
			}
			o.ResultNames = make([]string, sizeInfo[0])
			for i1 := range o.ResultNames {
				i1 := i1
				_ptr_ppResultNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.ResultNames[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_ppResultNames := func(ptr interface{}) { o.ResultNames[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.ResultNames[i1], _s_ppResultNames, _ptr_ppResultNames); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppResultNames := func(ptr interface{}) { o.ResultNames = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.ResultNames, _s_ppResultNames, _ptr_ppResultNames); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppResultFlags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=pcResults](uint32))
	{
		_ptr_ppResultFlags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ResultFlags", sizeInfo[0])
			}
			o.ResultFlags = make([]uint32, sizeInfo[0])
			for i1 := range o.ResultFlags {
				i1 := i1
				if err := w.ReadData(&o.ResultFlags[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppResultFlags := func(ptr interface{}) { o.ResultFlags = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.ResultFlags, _s_ppResultFlags, _ptr_ppResultFlags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppResultHRs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LONG}[dim:0,size_is=pcResults](int32))
	{
		_ptr_ppResultHRs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ResultHRs", sizeInfo[0])
			}
			o.ResultHRs = make([]int32, sizeInfo[0])
			for i1 := range o.ResultHRs {
				i1 := i1
				if err := w.ReadData(&o.ResultHRs[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppResultHRs := func(ptr interface{}) { o.ResultHRs = *ptr.(*[]int32) }
		if err := w.ReadPointer(&o.ResultHRs, _s_ppResultHRs, _ptr_ppResultHRs); err != nil {
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

// RegisterModuleRequest structure represents the RegisterModule operation request
type RegisterModuleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ConglomerationIdentifier:  The conglomeration identifier of an existing conglomeration
	// on the server, in which the component full configurations are to be created or against
	// which the modules are to be verified (as specified later).
	ConglomerationID *dtyp.GUID `idl:"name:ConglomerationIdentifier" json:"conglomeration_id"`
	// ppModules:  An array of one or more strings, each of which is a path in UNC to a
	// file that the server will recognize as a module.
	Modules []string `idl:"name:ppModules;size_is:(cModules, );string" json:"modules"`
	// cModules:  The number of elements in ppModules.
	ModulesCount uint32 `idl:"name:cModules" json:"modules_count"`
	// dwFlags:  A combination of zero or more of the following flags.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| fREGISTER_VERIFYONLY 0x00000020   | The server SHOULD verify the modules but MUST NOT actually register any          |
	//	|                                   | components.                                                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| fREGISTER_EVENTCLASSES 0x00000400 | The server MUST configure the components registered by this operation as event   |
	//	|                                   | classes.                                                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pRequestedCLSIDs:  Either an array of one or more CLSIDs of components to be registered
	// (or verified), or NULL to specify that all components in the modules are to be registered
	// (or verified).
	RequestedClassIDs []*dtyp.GUID `idl:"name:pRequestedCLSIDs;size_is:(cRequested);pointer:unique" json:"requested_class_ids"`
	// cRequested:  The number of elements in pRequestedCLSIDs.
	RequestedCount uint32 `idl:"name:cRequested" json:"requested_count"`
}

func (o *RegisterModuleRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterModuleOperation) *xxx_RegisterModuleOperation {
	if op == nil {
		op = &xxx_RegisterModuleOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ConglomerationID = o.ConglomerationID
	op.Modules = o.Modules
	op.ModulesCount = o.ModulesCount
	op.Flags = o.Flags
	op.RequestedClassIDs = o.RequestedClassIDs
	op.RequestedCount = o.RequestedCount
	return op
}

func (o *RegisterModuleRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterModuleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConglomerationID = op.ConglomerationID
	o.Modules = op.Modules
	o.ModulesCount = op.ModulesCount
	o.Flags = op.Flags
	o.RequestedClassIDs = op.RequestedClassIDs
	o.RequestedCount = op.RequestedCount
}
func (o *RegisterModuleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterModuleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterModuleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterModuleResponse structure represents the RegisterModule operation response
type RegisterModuleResponse struct {
	// XXX: cModules is an implicit input depedency for output parameters
	ModulesCount uint32 `idl:"name:cModules" json:"modules_count"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppModuleFlags:  A pointer to a variable that, upon successful completion, SHOULD
	// be set to an array of fModuleStatus (section 2.2.3) values that represent the detailed
	// results of registration for the modules located by the paths in ppModules, in the
	// same order.
	ModuleFlags []uint32 `idl:"name:ppModuleFlags;size_is:(, cModules)" json:"module_flags"`
	// pcResults:  A pointer to a variable that, upon successful completion, MUST be set
	// to the number of result items, as specified later.
	ResultsCount uint32 `idl:"name:pcResults" json:"results_count"`
	// ppResultCLSIDs:  A pointer to a variable that, upon successful completion, MUST
	// be set to an array of GUID values, each being the CLSID of a result item, as specified
	// later.
	ResultClassIDs []*dtyp.GUID `idl:"name:ppResultCLSIDs;size_is:(, pcResults)" json:"result_class_ids"`
	// ppResultNames:  A pointer to a variable that, upon successful completion, MUST be
	// set to an array of string values, each being an implementation-specific<317> name
	// of a result item, as specified later, in the same order as ppResultClsids.
	ResultNames []string `idl:"name:ppResultNames;size_is:(, pcResults);string" json:"result_names"`
	// ppResultFlags:  A pointer to a variable that upon successful completion, MUST be
	// set to an array of fComponentStatus (section 2.2.4) values, each representing implementation-specific
	// detailed results for a result item, as specified later, in the same order as ppResultClsids.
	ResultFlags []uint32 `idl:"name:ppResultFlags;size_is:(, pcResults)" json:"result_flags"`
	// ppResultHRs:  A pointer to a variable that, upon successful completion, MUST be
	// set to an array of LONG values, each representing an HRESULT ([MS-ERREF] section
	// 2.1) for a result item, as specified later, in the same order as ppResultClsids.
	ResultHRs []int32 `idl:"name:ppResultHRs;size_is:(, pcResults)" json:"result_h_rs"`
	// Return: The RegisterModule return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterModuleResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterModuleOperation) *xxx_RegisterModuleOperation {
	if op == nil {
		op = &xxx_RegisterModuleOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.ModulesCount == uint32(0) {
		op.ModulesCount = o.ModulesCount
	}

	op.That = o.That
	op.ModuleFlags = o.ModuleFlags
	op.ResultsCount = o.ResultsCount
	op.ResultClassIDs = o.ResultClassIDs
	op.ResultNames = o.ResultNames
	op.ResultFlags = o.ResultFlags
	op.ResultHRs = o.ResultHRs
	op.Return = o.Return
	return op
}

func (o *RegisterModuleResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterModuleOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.ModulesCount = op.ModulesCount

	o.That = op.That
	o.ModuleFlags = op.ModuleFlags
	o.ResultsCount = op.ResultsCount
	o.ResultClassIDs = op.ResultClassIDs
	o.ResultNames = op.ResultNames
	o.ResultFlags = op.ResultFlags
	o.ResultHRs = op.ResultHRs
	o.Return = op.Return
}
func (o *RegisterModuleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterModuleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterModuleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
