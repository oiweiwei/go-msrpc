package icatalogtableinfo

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	coma "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = dtyp.GoPackage
	_ = coma.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// ICatalogTableInfo interface identifier a8927a41-d3ce-11d1-8472-006008b0e5ca
	CatalogTableInfoIID = &dcom.IID{Data1: 0xa8927a41, Data2: 0xd3ce, Data3: 0x11d1, Data4: []byte{0x84, 0x72, 0x00, 0x60, 0x08, 0xb0, 0xe5, 0xca}}
	// Syntax UUID
	CatalogTableInfoSyntaxUUID = &uuid.UUID{TimeLow: 0xa8927a41, TimeMid: 0xd3ce, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0x72, Node: [6]uint8{0x0, 0x60, 0x8, 0xb0, 0xe5, 0xca}}
	// Syntax ID
	CatalogTableInfoSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CatalogTableInfoSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICatalogTableInfo interface.
type CatalogTableInfoClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to retrieve the table metadata (section 3.1.1.2.1)
	// for a catalog table.
	//
	// Return Values:  This method MUST return S_OK (0x00000000) on success and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	GetClientTableInfo(context.Context, *GetClientTableInfoRequest, ...dcerpc.CallOption) (*GetClientTableInfoResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CatalogTableInfoClient
}

type xxx_DefaultCatalogTableInfoClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCatalogTableInfoClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCatalogTableInfoClient) GetClientTableInfo(ctx context.Context, in *GetClientTableInfoRequest, opts ...dcerpc.CallOption) (*GetClientTableInfoResponse, error) {
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
	out := &GetClientTableInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogTableInfoClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCatalogTableInfoClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCatalogTableInfoClient) IPID(ctx context.Context, ipid *dcom.IPID) CatalogTableInfoClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCatalogTableInfoClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewCatalogTableInfoClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CatalogTableInfoClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CatalogTableInfoSyntaxV0_0))...)
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
	return &xxx_DefaultCatalogTableInfoClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetClientTableInfoOperation structure represents the GetClientTableInfo operation
type xxx_GetClientTableInfoOperation struct {
	This                  *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat       `idl:"name:That" json:"that"`
	CatalogID             *dtyp.GUID           `idl:"name:pCatalogIdentifier" json:"catalog_id"`
	TableID               *dtyp.GUID           `idl:"name:pTableIdentifier" json:"table_id"`
	TableFlags            uint32               `idl:"name:tableFlags" json:"table_flags"`
	QueryCellArray        []byte               `idl:"name:pQueryCellArray;size_is:(cbQueryCellArray);pointer:unique" json:"query_cell_array"`
	QueryCellArrayLength  uint32               `idl:"name:cbQueryCellArray" json:"query_cell_array_length"`
	QueryComparison       []byte               `idl:"name:pQueryComparison;size_is:(cbQueryComparison);pointer:unique" json:"query_comparison"`
	QueryComparisonLength uint32               `idl:"name:cbQueryComparison" json:"query_comparison_length"`
	QueryFormat           uint32               `idl:"name:eQueryFormat" json:"query_format"`
	RequiredFixedGUID     *dtyp.GUID           `idl:"name:pRequiredFixedGuid" json:"required_fixed_guid"`
	_                     []byte               `idl:"name:ppReserved1;size_is:(, pcbReserved1)"`
	_                     uint32               `idl:"name:pcbReserved1"`
	AuxiliaryGUID         []*dtyp.GUID         `idl:"name:ppAuxiliaryGuid;size_is:(, pcAuxiliaryGuid)" json:"auxiliary_guid"`
	AuxiliaryGUIDCount    uint32               `idl:"name:pcAuxiliaryGuid" json:"auxiliary_guid_count"`
	PropertyMeta          []*coma.PropertyMeta `idl:"name:ppPropertyMeta;size_is:(, pcProperties)" json:"property_meta"`
	PropertiesCount       uint32               `idl:"name:pcProperties" json:"properties_count"`
	IIDs                  *dcom.IID            `idl:"name:piid" json:"iids"`
	Interface             []byte               `idl:"name:pItf" json:"interface"`
	_                     []byte               `idl:"name:ppReserved2;size_is:(, pcbReserved2)"`
	_                     uint32               `idl:"name:pcbReserved2"`
	Return                int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClientTableInfoOperation) OpNum() int { return 3 }

func (o *xxx_GetClientTableInfoOperation) OpName() string {
	return "/ICatalogTableInfo/v0/GetClientTableInfo"
}

func (o *xxx_GetClientTableInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.QueryCellArray != nil && o.QueryCellArrayLength == 0 {
		o.QueryCellArrayLength = uint32(len(o.QueryCellArray))
	}
	if o.QueryComparison != nil && o.QueryComparisonLength == 0 {
		o.QueryComparisonLength = uint32(len(o.QueryComparison))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientTableInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pCatalogIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.CatalogID != nil {
			if err := o.CatalogID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pTableIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.TableID != nil {
			if err := o.TableID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// tableFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TableFlags); err != nil {
			return err
		}
	}
	// pQueryCellArray {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbQueryCellArray](char))
	{
		if o.QueryCellArray != nil || o.QueryCellArrayLength > 0 {
			_ptr_pQueryCellArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.QueryCellArrayLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.QueryCellArray {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.QueryCellArray[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.QueryCellArray); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryCellArray, _ptr_pQueryCellArray); err != nil {
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
	// cbQueryCellArray {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.QueryCellArrayLength); err != nil {
			return err
		}
	}
	// pQueryComparison {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbQueryComparison](char))
	{
		if o.QueryComparison != nil || o.QueryComparisonLength > 0 {
			_ptr_pQueryComparison := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.QueryComparisonLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.QueryComparison {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.QueryComparison[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.QueryComparison); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryComparison, _ptr_pQueryComparison); err != nil {
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
	// cbQueryComparison {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.QueryComparisonLength); err != nil {
			return err
		}
	}
	// eQueryFormat {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.QueryFormat); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientTableInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pCatalogIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.CatalogID == nil {
			o.CatalogID = &dtyp.GUID{}
		}
		if err := o.CatalogID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pTableIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.TableID == nil {
			o.TableID = &dtyp.GUID{}
		}
		if err := o.TableID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// tableFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TableFlags); err != nil {
			return err
		}
	}
	// pQueryCellArray {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbQueryCellArray](char))
	{
		_ptr_pQueryCellArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.QueryCellArray", sizeInfo[0])
			}
			o.QueryCellArray = make([]byte, sizeInfo[0])
			for i1 := range o.QueryCellArray {
				i1 := i1
				if err := w.ReadData(&o.QueryCellArray[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pQueryCellArray := func(ptr interface{}) { o.QueryCellArray = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.QueryCellArray, _s_pQueryCellArray, _ptr_pQueryCellArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbQueryCellArray {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.QueryCellArrayLength); err != nil {
			return err
		}
	}
	// pQueryComparison {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbQueryComparison](char))
	{
		_ptr_pQueryComparison := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.QueryComparison", sizeInfo[0])
			}
			o.QueryComparison = make([]byte, sizeInfo[0])
			for i1 := range o.QueryComparison {
				i1 := i1
				if err := w.ReadData(&o.QueryComparison[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pQueryComparison := func(ptr interface{}) { o.QueryComparison = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.QueryComparison, _s_pQueryComparison, _ptr_pQueryComparison); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbQueryComparison {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.QueryComparisonLength); err != nil {
			return err
		}
	}
	// eQueryFormat {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.QueryFormat); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientTableInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.AuxiliaryGUID != nil && o.AuxiliaryGUIDCount == 0 {
		o.AuxiliaryGUIDCount = uint32(len(o.AuxiliaryGUID))
	}
	if o.PropertyMeta != nil && o.PropertiesCount == 0 {
		o.PropertiesCount = uint32(len(o.PropertyMeta))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientTableInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pRequiredFixedGuid {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequiredFixedGUID != nil {
			if err := o.RequiredFixedGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ppReserved1 {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbReserved1](char))
	{
		// reserved ppReserved1
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pcbReserved1 {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		// reserved pcbReserved1
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// ppAuxiliaryGuid {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pcAuxiliaryGuid](struct))
	{
		if o.AuxiliaryGUID != nil || o.AuxiliaryGUIDCount > 0 {
			_ptr_ppAuxiliaryGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.AuxiliaryGUIDCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.AuxiliaryGUID {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.AuxiliaryGUID[i1] != nil {
						if err := o.AuxiliaryGUID[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.AuxiliaryGUID); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AuxiliaryGUID, _ptr_ppAuxiliaryGuid); err != nil {
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
	// pcAuxiliaryGuid {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.AuxiliaryGUIDCount); err != nil {
			return err
		}
	}
	// ppPropertyMeta {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyMeta}[dim:0,size_is=pcProperties](struct))
	{
		if o.PropertyMeta != nil || o.PropertiesCount > 0 {
			_ptr_ppPropertyMeta := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.PropertiesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PropertyMeta {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.PropertyMeta[i1] != nil {
						if err := o.PropertyMeta[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&coma.PropertyMeta{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.PropertyMeta); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&coma.PropertyMeta{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyMeta, _ptr_ppPropertyMeta); err != nil {
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
	// pcProperties {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.PropertiesCount); err != nil {
			return err
		}
	}
	// piid {out} (1:{pointer=ref}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IIDs != nil {
			if err := o.IIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pItf {out, iid_is={piid}} (1:{pointer=ref}*(2)*(1)(void))
	{
		if o.Interface != nil {
			_ptr_pItf := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				// FIXME unknown type pItf
				return nil
			})
			if err := w.WritePointer(&o.Interface, _ptr_pItf); err != nil {
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
	// ppReserved2 {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbReserved2](char))
	{
		// reserved ppReserved2
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pcbReserved2 {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		// reserved pcbReserved2
		if err := w.WriteData(uint32(0)); err != nil {
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

func (o *xxx_GetClientTableInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pRequiredFixedGuid {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequiredFixedGUID == nil {
			o.RequiredFixedGUID = &dtyp.GUID{}
		}
		if err := o.RequiredFixedGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ppReserved1 {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbReserved1](char))
	{
		// reserved ppReserved1
		var _ppReserved1 []byte
		_ptr_ppReserved1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _ppReserved1", sizeInfo[0])
			}
			_ppReserved1 = make([]byte, sizeInfo[0])
			for i1 := range _ppReserved1 {
				i1 := i1
				if err := w.ReadData(&_ppReserved1[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppReserved1 := func(ptr interface{}) { _ppReserved1 = *ptr.(*[]byte) }
		if err := w.ReadPointer(&_ppReserved1, _s_ppReserved1, _ptr_ppReserved1); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbReserved1 {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		// reserved pcbReserved1
		var _pcbReserved1 uint32
		if err := w.ReadData(&_pcbReserved1); err != nil {
			return err
		}
	}
	// ppAuxiliaryGuid {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pcAuxiliaryGuid](struct))
	{
		_ptr_ppAuxiliaryGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.AuxiliaryGUID", sizeInfo[0])
			}
			o.AuxiliaryGUID = make([]*dtyp.GUID, sizeInfo[0])
			for i1 := range o.AuxiliaryGUID {
				i1 := i1
				if o.AuxiliaryGUID[i1] == nil {
					o.AuxiliaryGUID[i1] = &dtyp.GUID{}
				}
				if err := o.AuxiliaryGUID[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppAuxiliaryGuid := func(ptr interface{}) { o.AuxiliaryGUID = *ptr.(*[]*dtyp.GUID) }
		if err := w.ReadPointer(&o.AuxiliaryGUID, _s_ppAuxiliaryGuid, _ptr_ppAuxiliaryGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcAuxiliaryGuid {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.AuxiliaryGUIDCount); err != nil {
			return err
		}
	}
	// ppPropertyMeta {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyMeta}[dim:0,size_is=pcProperties](struct))
	{
		_ptr_ppPropertyMeta := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.PropertyMeta", sizeInfo[0])
			}
			o.PropertyMeta = make([]*coma.PropertyMeta, sizeInfo[0])
			for i1 := range o.PropertyMeta {
				i1 := i1
				if o.PropertyMeta[i1] == nil {
					o.PropertyMeta[i1] = &coma.PropertyMeta{}
				}
				if err := o.PropertyMeta[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppPropertyMeta := func(ptr interface{}) { o.PropertyMeta = *ptr.(*[]*coma.PropertyMeta) }
		if err := w.ReadPointer(&o.PropertyMeta, _s_ppPropertyMeta, _ptr_ppPropertyMeta); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcProperties {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.PropertiesCount); err != nil {
			return err
		}
	}
	// piid {out} (1:{pointer=ref}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IIDs == nil {
			o.IIDs = &dcom.IID{}
		}
		if err := o.IIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pItf {out, iid_is={piid}} (1:{pointer=ref}*(2)*(1)(void))
	{
		_ptr_pItf := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			// FIXME: unknown type pItf
			return nil
		})
		_s_pItf := func(ptr interface{}) { o.Interface = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Interface, _s_pItf, _ptr_pItf); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppReserved2 {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbReserved2](char))
	{
		// reserved ppReserved2
		var _ppReserved2 []byte
		_ptr_ppReserved2 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _ppReserved2", sizeInfo[0])
			}
			_ppReserved2 = make([]byte, sizeInfo[0])
			for i1 := range _ppReserved2 {
				i1 := i1
				if err := w.ReadData(&_ppReserved2[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppReserved2 := func(ptr interface{}) { _ppReserved2 = *ptr.(*[]byte) }
		if err := w.ReadPointer(&_ppReserved2, _s_ppReserved2, _ptr_ppReserved2); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbReserved2 {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		// reserved pcbReserved2
		var _pcbReserved2 uint32
		if err := w.ReadData(&_pcbReserved2); err != nil {
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

// GetClientTableInfoRequest structure represents the GetClientTableInfo operation request
type GetClientTableInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pCatalogIdentifier:  The catalog identifier of the COMA catalog. MUST be set to
	// {6E38D3C4-C2A7-11D1-8DEC-00C04FC2E0C7}.
	CatalogID *dtyp.GUID `idl:"name:pCatalogIdentifier" json:"catalog_id"`
	// pTableIdentifier: The table identifier for one of the tables defined in section 3.1.1.3
	// for the negotiated catalog version.
	TableID *dtyp.GUID `idl:"name:pTableIdentifier" json:"table_id"`
	// tableFlags: An fTableFlags (section 2.2.1.1) value supported by the table identified
	// by pTableIdentifier.
	TableFlags uint32 `idl:"name:tableFlags" json:"table_flags"`
	// pQueryCellArray: A QueryCellArray (section 2.2.1.5) structure, marshaled in the negotiated
	// format, as specified in section 2.2.1.5, for a supported query (see section 3.1.1.2)
	// on the table identified by pTableIdentifier.
	QueryCellArray []byte `idl:"name:pQueryCellArray;size_is:(cbQueryCellArray);pointer:unique" json:"query_cell_array"`
	// cbQueryCellArray: The size in bytes of pQueryCellArray.
	QueryCellArrayLength uint32 `idl:"name:cbQueryCellArray" json:"query_cell_array_length"`
	// pQueryComparison: A QueryComparisonData (section 2.2.1.6) structure for a supported
	// query (see section 3.1.1.2) on the table identified by pTableIdentifier.
	QueryComparison []byte `idl:"name:pQueryComparison;size_is:(cbQueryComparison);pointer:unique" json:"query_comparison"`
	// cbQueryComparison: The size in bytes of pQueryComparison.
	QueryComparisonLength uint32 `idl:"name:cbQueryComparison" json:"query_comparison_length"`
	// eQueryFormat: MUST be set to eQUERYFORMAT_1 (0x00000001).
	QueryFormat uint32 `idl:"name:eQueryFormat" json:"query_format"`
}

func (o *GetClientTableInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetClientTableInfoOperation) *xxx_GetClientTableInfoOperation {
	if op == nil {
		op = &xxx_GetClientTableInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.CatalogID = o.CatalogID
	op.TableID = o.TableID
	op.TableFlags = o.TableFlags
	op.QueryCellArray = o.QueryCellArray
	op.QueryCellArrayLength = o.QueryCellArrayLength
	op.QueryComparison = o.QueryComparison
	op.QueryComparisonLength = o.QueryComparisonLength
	op.QueryFormat = o.QueryFormat
	return op
}

func (o *GetClientTableInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClientTableInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CatalogID = op.CatalogID
	o.TableID = op.TableID
	o.TableFlags = op.TableFlags
	o.QueryCellArray = op.QueryCellArray
	o.QueryCellArrayLength = op.QueryCellArrayLength
	o.QueryComparison = op.QueryComparison
	o.QueryComparisonLength = op.QueryComparisonLength
	o.QueryFormat = op.QueryFormat
}
func (o *GetClientTableInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetClientTableInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientTableInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClientTableInfoResponse structure represents the GetClientTableInfo operation response
type GetClientTableInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pRequiredFixedGuid:  A pointer to a variable that, upon successful completion, MUST
	// be set to {92AD68AB-17E0-11D1-B230-00C04FB9473F} and SHOULD<285> be ignored on receipt.
	RequiredFixedGUID *dtyp.GUID `idl:"name:pRequiredFixedGuid" json:"required_fixed_guid"`
	// ppAuxiliaryGuid: A pointer to a variable that, upon successful completion, MUST be
	// set to the AuxiliaryGuid value specified in the definition of the table identified
	// by pTableIdentifier, and NULL if the table definition specifies no such value. This
	// value SHOULD<286> be ignored on receipt.
	AuxiliaryGUID []*dtyp.GUID `idl:"name:ppAuxiliaryGuid;size_is:(, pcAuxiliaryGuid)" json:"auxiliary_guid"`
	// pcAuxiliaryGuid: A pointer to a variable that, upon successful completion, MUST be
	// the number of elements in ppAuxiliaryGuids (zero or one).
	AuxiliaryGUIDCount uint32 `idl:"name:pcAuxiliaryGuid" json:"auxiliary_guid_count"`
	// ppPropertyMeta: A pointer to a variable that, upon successful completion, MUST be
	// set to an array of PropertyMeta (section 2.2.1.7) structures representing the schema
	// (see section 3.1.1.1) for the table identified by pTableIdentifier in the negotiated
	// catalog version.
	PropertyMeta []*coma.PropertyMeta `idl:"name:ppPropertyMeta;size_is:(, pcProperties)" json:"property_meta"`
	// pcProperties: A pointer to a variable that, upon successful completion, MUST be set
	// to the length of the array returned in ppPropertyMeta.
	PropertiesCount uint32 `idl:"name:pcProperties" json:"properties_count"`
	// piid: A pointer to a variable that, upon successful completion, MUST be set to IID_ICatalogTableRead
	// (see section 1.9).
	IIDs *dcom.IID `idl:"name:piid" json:"iids"`
	// pItf: A pointer to a variable that, upon successful completion, MUST be set to the
	// ICatalogTableRead (section 3.1.4.8) interface of the server.
	Interface []byte `idl:"name:pItf" json:"interface"`
	// Return: The GetClientTableInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClientTableInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetClientTableInfoOperation) *xxx_GetClientTableInfoOperation {
	if op == nil {
		op = &xxx_GetClientTableInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.RequiredFixedGUID = o.RequiredFixedGUID
	op.AuxiliaryGUID = o.AuxiliaryGUID
	op.AuxiliaryGUIDCount = o.AuxiliaryGUIDCount
	op.PropertyMeta = o.PropertyMeta
	op.PropertiesCount = o.PropertiesCount
	op.IIDs = o.IIDs
	op.Interface = o.Interface
	op.Return = o.Return
	return op
}

func (o *GetClientTableInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClientTableInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RequiredFixedGUID = op.RequiredFixedGUID
	o.AuxiliaryGUID = op.AuxiliaryGUID
	o.AuxiliaryGUIDCount = op.AuxiliaryGUIDCount
	o.PropertyMeta = op.PropertyMeta
	o.PropertiesCount = op.PropertiesCount
	o.IIDs = op.IIDs
	o.Interface = op.Interface
	o.Return = op.Return
}
func (o *GetClientTableInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetClientTableInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientTableInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
