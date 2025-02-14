package iimport

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
	// IImport interface identifier c2be6970-df9e-11d1-8b87-00c04fd7a924
	ImportIID = &dcom.IID{Data1: 0xc2be6970, Data2: 0xdf9e, Data3: 0x11d1, Data4: []byte{0x8b, 0x87, 0x00, 0xc0, 0x4f, 0xd7, 0xa9, 0x24}}
	// Syntax UUID
	ImportSyntaxUUID = &uuid.UUID{TimeLow: 0xc2be6970, TimeMid: 0xdf9e, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8b, ClockSeqLow: 0x87, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd7, 0xa9, 0x24}}
	// Syntax ID
	ImportSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ImportSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IImport interface.
type ImportClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to import one or more conglomerations from an installer
	// package file. Importing a conglomeration from an installer package file conceptually
	// consists of installation of modules, including the registration of the components
	// in those modules, and creating a conglomeration and component configurations equivalent
	// to the conglomeration and the component configurations that were exported to create
	// the installer package file. As a side effect, this method returns implementation-specific
	// detailed results of the registration operation for informational purposes.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	ImportFromFile(context.Context, *ImportFromFileRequest, ...dcerpc.CallOption) (*ImportFromFileResponse, error)

	// This method is called by a client to retrieve information about an installer package
	// file.
	//
	// Return Values:  This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	QueryFile(context.Context, *QueryFileRequest, ...dcerpc.CallOption) (*QueryFileResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ImportClient
}

type xxx_DefaultImportClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultImportClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultImportClient) ImportFromFile(ctx context.Context, in *ImportFromFileRequest, opts ...dcerpc.CallOption) (*ImportFromFileResponse, error) {
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
	out := &ImportFromFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImportClient) QueryFile(ctx context.Context, in *QueryFileRequest, opts ...dcerpc.CallOption) (*QueryFileResponse, error) {
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
	out := &QueryFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImportClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultImportClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultImportClient) IPID(ctx context.Context, ipid *dcom.IPID) ImportClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultImportClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewImportClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ImportClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ImportSyntaxV0_0))...)
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
	return &xxx_DefaultImportClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ImportFromFileOperation structure represents the ImportFromFile operation
type xxx_ImportFromFileOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ModuleDestination string         `idl:"name:pwszModuleDestination;string;pointer:unique" json:"module_destination"`
	InstallerPackage  string         `idl:"name:pwszInstallerPackage;string" json:"installer_package"`
	User              string         `idl:"name:pwszUser;string;pointer:unique" json:"user"`
	Password          string         `idl:"name:pwszPassword;string;pointer:unique" json:"password"`
	RemoteServerName  string         `idl:"name:pwszRemoteServerName;string;pointer:unique" json:"remote_server_name"`
	Flags             uint32         `idl:"name:dwFlags" json:"flags"`
	_                 *dtyp.GUID     `idl:"name:reserved1"`
	_                 uint32         `idl:"name:reserved2"`
	ModulesCount      uint32         `idl:"name:pcModules" json:"modules_count"`
	ModuleFlags       []uint32       `idl:"name:ppModuleFlags;size_is:(, pcModules)" json:"module_flags"`
	Modules           []string       `idl:"name:ppModules;size_is:(, pcModules);string" json:"modules"`
	ComponentsCount   uint32         `idl:"name:pcComponents" json:"components_count"`
	ResultClassIDs    []*dtyp.GUID   `idl:"name:ppResultCLSIDs;size_is:(, pcComponents)" json:"result_class_ids"`
	ResultNames       []string       `idl:"name:ppResultNames;size_is:(, pcComponents);string" json:"result_names"`
	ResultFlags       []uint32       `idl:"name:ppResultFlags;size_is:(, pcComponents)" json:"result_flags"`
	ResultHRs         []int32        `idl:"name:ppResultHRs;size_is:(, pcComponents)" json:"result_h_rs"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ImportFromFileOperation) OpNum() int { return 3 }

func (o *xxx_ImportFromFileOperation) OpName() string { return "/IImport/v0/ImportFromFile" }

func (o *xxx_ImportFromFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportFromFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszModuleDestination {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ModuleDestination != "" {
			_ptr_pwszModuleDestination := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ModuleDestination); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ModuleDestination, _ptr_pwszModuleDestination); err != nil {
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
	// pwszInstallerPackage {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.InstallerPackage); err != nil {
			return err
		}
	}
	// pwszUser {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.User != "" {
			_ptr_pwszUser := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.User); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.User, _ptr_pwszUser); err != nil {
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
	// pwszPassword {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.Password != "" {
			_ptr_pwszPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_pwszPassword); err != nil {
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
	// pwszRemoteServerName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.RemoteServerName != "" {
			_ptr_pwszRemoteServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.RemoteServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.RemoteServerName, _ptr_pwszRemoteServerName); err != nil {
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
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// reserved1 {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		// reserved reserved1
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved2 {in} (1:{alias=DWORD}(uint32))
	{
		// reserved reserved2
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportFromFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszModuleDestination {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszModuleDestination := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ModuleDestination); err != nil {
				return err
			}
			return nil
		})
		_s_pwszModuleDestination := func(ptr interface{}) { o.ModuleDestination = *ptr.(*string) }
		if err := w.ReadPointer(&o.ModuleDestination, _s_pwszModuleDestination, _ptr_pwszModuleDestination); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszInstallerPackage {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.InstallerPackage); err != nil {
			return err
		}
	}
	// pwszUser {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszUser := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.User); err != nil {
				return err
			}
			return nil
		})
		_s_pwszUser := func(ptr interface{}) { o.User = *ptr.(*string) }
		if err := w.ReadPointer(&o.User, _s_pwszUser, _ptr_pwszUser); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszPassword {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
				return err
			}
			return nil
		})
		_s_pwszPassword := func(ptr interface{}) { o.Password = *ptr.(*string) }
		if err := w.ReadPointer(&o.Password, _s_pwszPassword, _ptr_pwszPassword); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszRemoteServerName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszRemoteServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszRemoteServerName := func(ptr interface{}) { o.RemoteServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.RemoteServerName, _s_pwszRemoteServerName, _ptr_pwszRemoteServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// reserved1 {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		// reserved reserved1
		var _reserved1 *dtyp.GUID
		if _reserved1 == nil {
			_reserved1 = &dtyp.GUID{}
		}
		if err := _reserved1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved2 {in} (1:{alias=DWORD}(uint32))
	{
		// reserved reserved2
		var _reserved2 uint32
		if err := w.ReadData(&_reserved2); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportFromFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ModuleFlags != nil && o.ModulesCount == 0 {
		o.ModulesCount = uint32(len(o.ModuleFlags))
	}
	if o.Modules != nil && o.ModulesCount == 0 {
		o.ModulesCount = uint32(len(o.Modules))
	}
	if o.ResultClassIDs != nil && o.ComponentsCount == 0 {
		o.ComponentsCount = uint32(len(o.ResultClassIDs))
	}
	if o.ResultNames != nil && o.ComponentsCount == 0 {
		o.ComponentsCount = uint32(len(o.ResultNames))
	}
	if o.ResultFlags != nil && o.ComponentsCount == 0 {
		o.ComponentsCount = uint32(len(o.ResultFlags))
	}
	if o.ResultHRs != nil && o.ComponentsCount == 0 {
		o.ComponentsCount = uint32(len(o.ResultHRs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportFromFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcModules {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ModulesCount); err != nil {
			return err
		}
	}
	// ppModuleFlags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=pcModules](uint32))
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
	// ppModules {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcModules]*(1)[dim:0,string,null](wchar))
	{
		if o.Modules != nil || o.ModulesCount > 0 {
			_ptr_ppModules := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
				return nil
			})
			if err := w.WritePointer(&o.Modules, _ptr_ppModules); err != nil {
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
	// pcComponents {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ComponentsCount); err != nil {
			return err
		}
	}
	// ppResultCLSIDs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pcComponents](struct))
	{
		if o.ResultClassIDs != nil || o.ComponentsCount > 0 {
			_ptr_ppResultCLSIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ComponentsCount)
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
	// ppResultNames {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcComponents]*(1)[dim:0,string,null](wchar))
	{
		if o.ResultNames != nil || o.ComponentsCount > 0 {
			_ptr_ppResultNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ComponentsCount)
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
	// ppResultFlags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=pcComponents](uint32))
	{
		if o.ResultFlags != nil || o.ComponentsCount > 0 {
			_ptr_ppResultFlags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ComponentsCount)
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
	// ppResultHRs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LONG}[dim:0,size_is=pcComponents](int32))
	{
		if o.ResultHRs != nil || o.ComponentsCount > 0 {
			_ptr_ppResultHRs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ComponentsCount)
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

func (o *xxx_ImportFromFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcModules {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ModulesCount); err != nil {
			return err
		}
	}
	// ppModuleFlags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=pcModules](uint32))
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
	// ppModules {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcModules]*(1)[dim:0,string,null](wchar))
	{
		_ptr_ppModules := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			return nil
		})
		_s_ppModules := func(ptr interface{}) { o.Modules = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.Modules, _s_ppModules, _ptr_ppModules); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcComponents {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ComponentsCount); err != nil {
			return err
		}
	}
	// ppResultCLSIDs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pcComponents](struct))
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
	// ppResultNames {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcComponents]*(1)[dim:0,string,null](wchar))
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
	// ppResultFlags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=pcComponents](uint32))
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
	// ppResultHRs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LONG}[dim:0,size_is=pcComponents](int32))
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

// ImportFromFileRequest structure represents the ImportFromFile operation request
type ImportFromFileRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszModuleDestination:  Either a path in UNC to a directory that is to be used as
	// the installation target location for modules and other files, or NULL to indicate
	// that a directory is to be selected by the server.
	ModuleDestination string `idl:"name:pwszModuleDestination;string;pointer:unique" json:"module_destination"`
	// pwszInstallerPackage: A path in UNC to a file that the server will recognize as an
	// installer package file.
	InstallerPackage string `idl:"name:pwszInstallerPackage;string" json:"installer_package"`
	// pwszUser: Either a user account name to be used as the RunAsUser property (see section
	// 3.1.1.3.6) for the newly created conglomerations, or NULL to specify that the RunAsUser
	// property for each conglomeration represented in the installer package file is to
	// be used.
	User string `idl:"name:pwszUser;string;pointer:unique" json:"user"`
	// pwszPassword: Either a password to be used as the Password property (see section
	// 3.1.1.3.3) for the newly created conglomerations, or NULL to specify that the Password
	// property is to be set to NULL.
	Password string `idl:"name:pwszPassword;string;pointer:unique" json:"password"`
	// pwszRemoteServerName:  Either a remote server name to be used as the ServerName
	// property (see section 3.1.1.3.6) for the newly created conglomerations if the conglomerations
	// represented in the installer package file are proxy conglomerations; that is, if
	// they have the IsProxyApp property (see section 3.1.1.3.27) set to TRUE (0x00000001)),
	// or NULL for nonproxy conglomerations.
	RemoteServerName string `idl:"name:pwszRemoteServerName;string;pointer:unique" json:"remote_server_name"`
	// dwFlags: MUST be a combination of zero or more of the following flags.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|             FLAG             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| fIMPORT_OVERWRITE 0x00000001 | The server is requested to overwrite existing files when installing modules.     |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| fIMPORT_WITHUSERS 0x00000010 | The server is requested to create the role members represented in the installer  |
	//	|                              | package file.                                                                    |
	//	+------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
}

func (o *ImportFromFileRequest) xxx_ToOp(ctx context.Context, op *xxx_ImportFromFileOperation) *xxx_ImportFromFileOperation {
	if op == nil {
		op = &xxx_ImportFromFileOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ModuleDestination = o.ModuleDestination
	op.InstallerPackage = o.InstallerPackage
	op.User = o.User
	op.Password = o.Password
	op.RemoteServerName = o.RemoteServerName
	op.Flags = o.Flags
	return op
}

func (o *ImportFromFileRequest) xxx_FromOp(ctx context.Context, op *xxx_ImportFromFileOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ModuleDestination = op.ModuleDestination
	o.InstallerPackage = op.InstallerPackage
	o.User = op.User
	o.Password = op.Password
	o.RemoteServerName = op.RemoteServerName
	o.Flags = op.Flags
}
func (o *ImportFromFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ImportFromFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportFromFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ImportFromFileResponse structure represents the ImportFromFile operation response
type ImportFromFileResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pcModules:  A pointer to a variable that, upon successful completion, MUST be set
	// to the number of modules installed from the installer package file.
	ModulesCount uint32 `idl:"name:pcModules" json:"modules_count"`
	// ppModuleFlags:  A pointer to a variable that upon successful completion, MUST be
	// set to an array of fModuleStatus (section 2.2.3) values representing the detailed
	// results of registration for the modules in ppModules, in the same order.
	ModuleFlags []uint32 `idl:"name:ppModuleFlags;size_is:(, pcModules)" json:"module_flags"`
	// ppModules:  A pointer to a variable that, upon successful completion, MUST be set
	// to an array of strings representing file names of modules installed from the installer
	// package file.
	Modules []string `idl:"name:ppModules;size_is:(, pcModules);string" json:"modules"`
	// pcComponents:  A pointer to a variable that, upon successful completion, MUST be
	// set to the number of components that the server registered or attempted to register.
	ComponentsCount uint32 `idl:"name:pcComponents" json:"components_count"`
	// ppResultCLSIDs:  A pointer to a variable that, upon successful completion, MUST
	// be set to an array of GUID values, each being the CLSID of a component that the server
	// registered or attempted to register.
	ResultClassIDs []*dtyp.GUID `idl:"name:ppResultCLSIDs;size_is:(, pcComponents)" json:"result_class_ids"`
	// ppResultNames:  A pointer to a variable that, upon successful completion, SHOULD
	// be set to an array of string values, each being an implementation-specific name of
	// a component that the server registered or attempted to register, in the same order
	// as ppResultClsids.
	ResultNames []string `idl:"name:ppResultNames;size_is:(, pcComponents);string" json:"result_names"`
	// ppResultFlags:  A pointer to a variable that, upon successful completion, SHOULD
	// be set to an array of fComponentStatus (section 2.2.4) values, each representing
	// detailed results for a component that the server registered or attempted to register,
	// in the same order as ppResultClsids.
	ResultFlags []uint32 `idl:"name:ppResultFlags;size_is:(, pcComponents)" json:"result_flags"`
	// ppResultHRs:  A pointer to a variable that, upon successful completion, SHOULD be
	// set to an array of LONG values, each representing an HRESULT ([MS-ERREF] section
	// 2.1) for a component that the server registered or attempted to register, in the
	// same order as ppResultClsids.
	ResultHRs []int32 `idl:"name:ppResultHRs;size_is:(, pcComponents)" json:"result_h_rs"`
	// Return: The ImportFromFile return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ImportFromFileResponse) xxx_ToOp(ctx context.Context, op *xxx_ImportFromFileOperation) *xxx_ImportFromFileOperation {
	if op == nil {
		op = &xxx_ImportFromFileOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ModulesCount = o.ModulesCount
	op.ModuleFlags = o.ModuleFlags
	op.Modules = o.Modules
	op.ComponentsCount = o.ComponentsCount
	op.ResultClassIDs = o.ResultClassIDs
	op.ResultNames = o.ResultNames
	op.ResultFlags = o.ResultFlags
	op.ResultHRs = o.ResultHRs
	op.Return = o.Return
	return op
}

func (o *ImportFromFileResponse) xxx_FromOp(ctx context.Context, op *xxx_ImportFromFileOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ModulesCount = op.ModulesCount
	o.ModuleFlags = op.ModuleFlags
	o.Modules = op.Modules
	o.ComponentsCount = op.ComponentsCount
	o.ResultClassIDs = op.ResultClassIDs
	o.ResultNames = op.ResultNames
	o.ResultFlags = op.ResultFlags
	o.ResultHRs = op.ResultHRs
	o.Return = op.Return
}
func (o *ImportFromFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ImportFromFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportFromFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryFileOperation structure represents the QueryFile operation
type xxx_QueryFileOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	InstallerPackage string         `idl:"name:pwszInstallerPackage;string" json:"installer_package"`
	Conglomerations  uint32         `idl:"name:pdwConglomerations" json:"conglomerations"`
	Names            []string       `idl:"name:ppNames;size_is:(, pdwConglomerations);string" json:"names"`
	Descriptions     []string       `idl:"name:ppDescriptions;size_is:(, pdwConglomerations);string" json:"descriptions"`
	Users            uint32         `idl:"name:pdwUsers" json:"users"`
	IsProxy          uint32         `idl:"name:pdwIsProxy" json:"is_proxy"`
	ModulesCount     uint32         `idl:"name:pcModules" json:"modules_count"`
	Modules          []string       `idl:"name:ppModules;size_is:(, pcModules);string" json:"modules"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryFileOperation) OpNum() int { return 4 }

func (o *xxx_QueryFileOperation) OpName() string { return "/IImport/v0/QueryFile" }

func (o *xxx_QueryFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszInstallerPackage {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.InstallerPackage); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszInstallerPackage {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.InstallerPackage); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Names != nil && o.Conglomerations == 0 {
		o.Conglomerations = uint32(len(o.Names))
	}
	if o.Descriptions != nil && o.Conglomerations == 0 {
		o.Conglomerations = uint32(len(o.Descriptions))
	}
	if o.Modules != nil && o.ModulesCount == 0 {
		o.ModulesCount = uint32(len(o.Modules))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwConglomerations {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Conglomerations); err != nil {
			return err
		}
	}
	// ppNames {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pdwConglomerations]*(1)[dim:0,string,null](wchar))
	{
		if o.Names != nil || o.Conglomerations > 0 {
			_ptr_ppNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Conglomerations)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Names {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Names[i1] != "" {
						_ptr_ppNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.Names[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.Names[i1], _ptr_ppNames); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Names, _ptr_ppNames); err != nil {
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
	// ppDescriptions {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pdwConglomerations]*(1)[dim:0,string,null](wchar))
	{
		if o.Descriptions != nil || o.Conglomerations > 0 {
			_ptr_ppDescriptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Conglomerations)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Descriptions {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Descriptions[i1] != "" {
						_ptr_ppDescriptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.Descriptions[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.Descriptions[i1], _ptr_ppDescriptions); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Descriptions); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Descriptions, _ptr_ppDescriptions); err != nil {
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
	// pdwUsers {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Users); err != nil {
			return err
		}
	}
	// pdwIsProxy {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.IsProxy); err != nil {
			return err
		}
	}
	// pcModules {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ModulesCount); err != nil {
			return err
		}
	}
	// ppModules {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcModules]*(1)[dim:0,string,null](wchar))
	{
		if o.Modules != nil || o.ModulesCount > 0 {
			_ptr_ppModules := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
				return nil
			})
			if err := w.WritePointer(&o.Modules, _ptr_ppModules); err != nil {
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

func (o *xxx_QueryFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwConglomerations {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Conglomerations); err != nil {
			return err
		}
	}
	// ppNames {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pdwConglomerations]*(1)[dim:0,string,null](wchar))
	{
		_ptr_ppNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
			}
			o.Names = make([]string, sizeInfo[0])
			for i1 := range o.Names {
				i1 := i1
				_ptr_ppNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.Names[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_ppNames := func(ptr interface{}) { o.Names[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.Names[i1], _s_ppNames, _ptr_ppNames); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppNames := func(ptr interface{}) { o.Names = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.Names, _s_ppNames, _ptr_ppNames); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppDescriptions {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pdwConglomerations]*(1)[dim:0,string,null](wchar))
	{
		_ptr_ppDescriptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Descriptions", sizeInfo[0])
			}
			o.Descriptions = make([]string, sizeInfo[0])
			for i1 := range o.Descriptions {
				i1 := i1
				_ptr_ppDescriptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.Descriptions[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_ppDescriptions := func(ptr interface{}) { o.Descriptions[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.Descriptions[i1], _s_ppDescriptions, _ptr_ppDescriptions); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppDescriptions := func(ptr interface{}) { o.Descriptions = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.Descriptions, _s_ppDescriptions, _ptr_ppDescriptions); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwUsers {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Users); err != nil {
			return err
		}
	}
	// pdwIsProxy {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.IsProxy); err != nil {
			return err
		}
	}
	// pcModules {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ModulesCount); err != nil {
			return err
		}
	}
	// ppModules {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcModules]*(1)[dim:0,string,null](wchar))
	{
		_ptr_ppModules := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			return nil
		})
		_s_ppModules := func(ptr interface{}) { o.Modules = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.Modules, _s_ppModules, _ptr_ppModules); err != nil {
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

// QueryFileRequest structure represents the QueryFile operation request
type QueryFileRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszInstallerPackage:  A path in UNC to a file that the server will recognize as
	// an installer package file.
	InstallerPackage string `idl:"name:pwszInstallerPackage;string" json:"installer_package"`
}

func (o *QueryFileRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryFileOperation) *xxx_QueryFileOperation {
	if op == nil {
		op = &xxx_QueryFileOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.InstallerPackage = o.InstallerPackage
	return op
}

func (o *QueryFileRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryFileOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.InstallerPackage = op.InstallerPackage
}
func (o *QueryFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryFileResponse structure represents the QueryFile operation response
type QueryFileResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwConglomerations:  A pointer to a variable that, upon successful completion, MUST
	// be set to the number of conglomerations represented in the installer package file.
	Conglomerations uint32 `idl:"name:pdwConglomerations" json:"conglomerations"`
	// ppNames:  A pointer to a variable that, upon successful completion, MUST be set
	// to an array of string values, each of which is the value of the Name property (see
	// section 3.1.1.3.3) of a conglomeration represented in the installer package file.
	Names []string `idl:"name:ppNames;size_is:(, pdwConglomerations);string" json:"names"`
	// ppDescriptions:  A pointer to a variable that, upon successful completion, MUST
	// be set to an array of string values, each of which is the value of the Description
	// property (see section 3.1.1.3.1 ) of a conglomeration represented in the installer
	// package file.
	Descriptions []string `idl:"name:ppDescriptions;size_is:(, pdwConglomerations);string" json:"descriptions"`
	// pdwUsers:  A pointer to a variable that, upon successful completion, MUST be set
	// to the value TRUE (0x00000001) if the installer package file contains configuration
	// for user accounts, and FALSE (0x00000000) otherwise.
	Users uint32 `idl:"name:pdwUsers" json:"users"`
	// pdwIsProxy:  A pointer to a variable that, upon successful completion, MUST be set
	// to the value TRUE (0x00000001) if the installer package file contains a proxy conglomeration,
	// and FALSE (0x00000000) otherwise.
	IsProxy uint32 `idl:"name:pdwIsProxy" json:"is_proxy"`
	// pcModules:  A pointer to a variable that, upon successful completion, MUST be set
	// to the number of module contained in the installer package file.
	ModulesCount uint32 `idl:"name:pcModules" json:"modules_count"`
	// ppModules:  A pointer to a variable that, upon successful completion, MUST be set
	// to an array of strings, one for each module contained in the installer package file,
	// each representing a file name for the module.
	Modules []string `idl:"name:ppModules;size_is:(, pcModules);string" json:"modules"`
	// Return: The QueryFile return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryFileResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryFileOperation) *xxx_QueryFileOperation {
	if op == nil {
		op = &xxx_QueryFileOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Conglomerations = o.Conglomerations
	op.Names = o.Names
	op.Descriptions = o.Descriptions
	op.Users = o.Users
	op.IsProxy = o.IsProxy
	op.ModulesCount = o.ModulesCount
	op.Modules = o.Modules
	op.Return = o.Return
	return op
}

func (o *QueryFileResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryFileOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Conglomerations = op.Conglomerations
	o.Names = op.Names
	o.Descriptions = op.Descriptions
	o.Users = op.Users
	o.IsProxy = op.IsProxy
	o.ModulesCount = op.ModulesCount
	o.Modules = op.Modules
	o.Return = op.Return
}
func (o *QueryFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
