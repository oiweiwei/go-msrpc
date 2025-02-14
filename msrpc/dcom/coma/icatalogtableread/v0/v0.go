package icatalogtableread

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
	// ICatalogTableRead interface identifier 0e3d6630-b46b-11d1-9d2d-006008b0e5ca
	CatalogTableReadIID = &dcom.IID{Data1: 0x0e3d6630, Data2: 0xb46b, Data3: 0x11d1, Data4: []byte{0x9d, 0x2d, 0x00, 0x60, 0x08, 0xb0, 0xe5, 0xca}}
	// Syntax UUID
	CatalogTableReadSyntaxUUID = &uuid.UUID{TimeLow: 0xe3d6630, TimeMid: 0xb46b, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x2d, Node: [6]uint8{0x0, 0x60, 0x8, 0xb0, 0xe5, 0xca}}
	// Syntax ID
	CatalogTableReadSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CatalogTableReadSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICatalogTableRead interface.
type CatalogTableReadClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to read entries from a catalog table according
	// to a query.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically, with the exception of E_DETAILEDERRORS (0x80110802).
	ReadTable(context.Context, *ReadTableRequest, ...dcerpc.CallOption) (*ReadTableResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CatalogTableReadClient
}

type xxx_DefaultCatalogTableReadClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCatalogTableReadClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCatalogTableReadClient) ReadTable(ctx context.Context, in *ReadTableRequest, opts ...dcerpc.CallOption) (*ReadTableResponse, error) {
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
	out := &ReadTableResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogTableReadClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCatalogTableReadClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCatalogTableReadClient) IPID(ctx context.Context, ipid *dcom.IPID) CatalogTableReadClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCatalogTableReadClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewCatalogTableReadClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CatalogTableReadClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CatalogTableReadSyntaxV0_0))...)
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
	return &xxx_DefaultCatalogTableReadClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ReadTableOperation structure represents the ReadTable operation
type xxx_ReadTableOperation struct {
	This                      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                      *dcom.ORPCThat `idl:"name:That" json:"that"`
	CatalogID                 *dtyp.GUID     `idl:"name:pCatalogIdentifier" json:"catalog_id"`
	TableID                   *dtyp.GUID     `idl:"name:pTableIdentifier" json:"table_id"`
	TableFlags                uint32         `idl:"name:tableFlags" json:"table_flags"`
	QueryCellArray            []byte         `idl:"name:pQueryCellArray;size_is:(cbQueryCellArray);pointer:unique" json:"query_cell_array"`
	QueryCellArrayLength      uint32         `idl:"name:cbQueryCellArray" json:"query_cell_array_length"`
	QueryComparison           []byte         `idl:"name:pQueryComparison;size_is:(cbQueryComparison);pointer:unique" json:"query_comparison"`
	QueryComparisonLength     uint32         `idl:"name:cbQueryComparison" json:"query_comparison_length"`
	QueryFormat               uint32         `idl:"name:eQueryFormat" json:"query_format"`
	TableDataFixed            []byte         `idl:"name:ppTableDataFixed;size_is:(, pcbTableDataFixed)" json:"table_data_fixed"`
	TableDataFixedLength      uint32         `idl:"name:pcbTableDataFixed" json:"table_data_fixed_length"`
	TableDataVariable         []byte         `idl:"name:ppTableDataVariable;size_is:(, pcbTableDataVariable)" json:"table_data_variable"`
	TableDataVariableLength   uint32         `idl:"name:pcbTableDataVariable" json:"table_data_variable_length"`
	TableDetailedErrors       []byte         `idl:"name:ppTableDetailedErrors;size_is:(, pcbTableDetailedErrors)" json:"table_detailed_errors"`
	TableDetailedErrorsLength uint32         `idl:"name:pcbTableDetailedErrors" json:"table_detailed_errors_length"`
	_                         []byte         `idl:"name:ppReserved1;size_is:(, pcbReserved1)"`
	_                         uint32         `idl:"name:pcbReserved1"`
	_                         []byte         `idl:"name:ppReserved2;size_is:(, pcbReserved2)"`
	_                         uint32         `idl:"name:pcbReserved2"`
	Return                    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReadTableOperation) OpNum() int { return 3 }

func (o *xxx_ReadTableOperation) OpName() string { return "/ICatalogTableRead/v0/ReadTable" }

func (o *xxx_ReadTableOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
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

func (o *xxx_ReadTableOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ReadTableOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ReadTableOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.TableDataFixed != nil && o.TableDataFixedLength == 0 {
		o.TableDataFixedLength = uint32(len(o.TableDataFixed))
	}
	if o.TableDataVariable != nil && o.TableDataVariableLength == 0 {
		o.TableDataVariableLength = uint32(len(o.TableDataVariable))
	}
	if o.TableDetailedErrors != nil && o.TableDetailedErrorsLength == 0 {
		o.TableDetailedErrorsLength = uint32(len(o.TableDetailedErrors))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadTableOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppTableDataFixed {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbTableDataFixed](char))
	{
		if o.TableDataFixed != nil || o.TableDataFixedLength > 0 {
			_ptr_ppTableDataFixed := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.TableDataFixedLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.TableDataFixed {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.TableDataFixed[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.TableDataFixed); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TableDataFixed, _ptr_ppTableDataFixed); err != nil {
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
	// pcbTableDataFixed {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.TableDataFixedLength); err != nil {
			return err
		}
	}
	// ppTableDataVariable {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbTableDataVariable](char))
	{
		if o.TableDataVariable != nil || o.TableDataVariableLength > 0 {
			_ptr_ppTableDataVariable := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.TableDataVariableLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.TableDataVariable {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.TableDataVariable[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.TableDataVariable); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TableDataVariable, _ptr_ppTableDataVariable); err != nil {
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
	// pcbTableDataVariable {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.TableDataVariableLength); err != nil {
			return err
		}
	}
	// ppTableDetailedErrors {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbTableDetailedErrors](char))
	{
		if o.TableDetailedErrors != nil || o.TableDetailedErrorsLength > 0 {
			_ptr_ppTableDetailedErrors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.TableDetailedErrorsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.TableDetailedErrors {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.TableDetailedErrors[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.TableDetailedErrors); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TableDetailedErrors, _ptr_ppTableDetailedErrors); err != nil {
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
	// pcbTableDetailedErrors {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.TableDetailedErrorsLength); err != nil {
			return err
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

func (o *xxx_ReadTableOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppTableDataFixed {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbTableDataFixed](char))
	{
		_ptr_ppTableDataFixed := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.TableDataFixed", sizeInfo[0])
			}
			o.TableDataFixed = make([]byte, sizeInfo[0])
			for i1 := range o.TableDataFixed {
				i1 := i1
				if err := w.ReadData(&o.TableDataFixed[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppTableDataFixed := func(ptr interface{}) { o.TableDataFixed = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.TableDataFixed, _s_ppTableDataFixed, _ptr_ppTableDataFixed); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbTableDataFixed {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.TableDataFixedLength); err != nil {
			return err
		}
	}
	// ppTableDataVariable {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbTableDataVariable](char))
	{
		_ptr_ppTableDataVariable := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.TableDataVariable", sizeInfo[0])
			}
			o.TableDataVariable = make([]byte, sizeInfo[0])
			for i1 := range o.TableDataVariable {
				i1 := i1
				if err := w.ReadData(&o.TableDataVariable[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppTableDataVariable := func(ptr interface{}) { o.TableDataVariable = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.TableDataVariable, _s_ppTableDataVariable, _ptr_ppTableDataVariable); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbTableDataVariable {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.TableDataVariableLength); err != nil {
			return err
		}
	}
	// ppTableDetailedErrors {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbTableDetailedErrors](char))
	{
		_ptr_ppTableDetailedErrors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.TableDetailedErrors", sizeInfo[0])
			}
			o.TableDetailedErrors = make([]byte, sizeInfo[0])
			for i1 := range o.TableDetailedErrors {
				i1 := i1
				if err := w.ReadData(&o.TableDetailedErrors[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppTableDetailedErrors := func(ptr interface{}) { o.TableDetailedErrors = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.TableDetailedErrors, _s_ppTableDetailedErrors, _ptr_ppTableDetailedErrors); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbTableDetailedErrors {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.TableDetailedErrorsLength); err != nil {
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

// ReadTableRequest structure represents the ReadTable operation request
type ReadTableRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pCatalogIdentifier:  The catalog identifier of the COMA catalog. MUST be set to
	// {6E38D3C4-C2A7-11D1-8DEC-00C04FC2E0C7}.
	CatalogID *dtyp.GUID `idl:"name:pCatalogIdentifier" json:"catalog_id"`
	// pTableIdentifier:  The table identifier for one of the tables defined in section
	// 3.1.1.3 for the negotiated catalog version.
	TableID *dtyp.GUID `idl:"name:pTableIdentifier" json:"table_id"`
	// tableFlags:  An fTableFlags (section 2.2.1.1) value supported (see section 3.1.1.2.3)
	// by the table identified by pTableIdentifier.
	TableFlags uint32 `idl:"name:tableFlags" json:"table_flags"`
	// pQueryCellArray:  A QueryCellArray structure, marshaled in the negotiated format
	// as specified in section 2.2.1.5, for a supported query (see section 3.1.1.2.2) on
	// the table identified by pTableIdentifier.
	QueryCellArray []byte `idl:"name:pQueryCellArray;size_is:(cbQueryCellArray);pointer:unique" json:"query_cell_array"`
	// cbQueryCellArray:  The size in bytes of pQueryCellArray.
	QueryCellArrayLength uint32 `idl:"name:cbQueryCellArray" json:"query_cell_array_length"`
	// pQueryComparison:  A QueryComparisonData (section 2.2.1.6) structure for a supported
	// query (see section 3.1.1.2.2) on the table identified by pTableIdentifier.
	QueryComparison []byte `idl:"name:pQueryComparison;size_is:(cbQueryComparison);pointer:unique" json:"query_comparison"`
	// cbQueryComparison:  The size in bytes of pQueryComparison.
	QueryComparisonLength uint32 `idl:"name:cbQueryComparison" json:"query_comparison_length"`
	// eQueryFormat:  MUST be set to eQUERYFORMAT_1 (0x00000001).
	QueryFormat uint32 `idl:"name:eQueryFormat" json:"query_format"`
}

func (o *ReadTableRequest) xxx_ToOp(ctx context.Context, op *xxx_ReadTableOperation) *xxx_ReadTableOperation {
	if op == nil {
		op = &xxx_ReadTableOperation{}
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

func (o *ReadTableRequest) xxx_FromOp(ctx context.Context, op *xxx_ReadTableOperation) {
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
func (o *ReadTableRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReadTableRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReadTableOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReadTableResponse structure represents the ReadTable operation response
type ReadTableResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTableDataFixed:  A pointer to a variable that, upon successful completion, MUST
	// be set to a TableDataFixed structure, marshaled as specified in section 2.2.1.10.
	TableDataFixed []byte `idl:"name:ppTableDataFixed;size_is:(, pcbTableDataFixed)" json:"table_data_fixed"`
	// pcbTableDataFixed:  A pointer to a value that, upon successful completion, MUST
	// be set to the length in bytes of the TableDataFixed structure returned in ppTableDataFixed.
	TableDataFixedLength uint32 `idl:"name:pcbTableDataFixed" json:"table_data_fixed_length"`
	// ppTableDataVariable:  A pointer to a pointer variable that, upon successful completion,
	// MUST be set to a TableDataVariable structure, marshaled as specified in section 2.2.1.15.
	TableDataVariable []byte `idl:"name:ppTableDataVariable;size_is:(, pcbTableDataVariable)" json:"table_data_variable"`
	// pcbTableDataVariable: A pointer to a value that, upon successful completion, MUST
	// be the length in bytes of the TableDataVariable structure returned in ppTableDataVariable.
	TableDataVariableLength uint32 `idl:"name:pcbTableDataVariable" json:"table_data_variable_length"`
	// ppTableDetailedErrors: A pointer to a variable that, upon successful completion,
	// MUST be set to NULL, and that upon partial failure MAY<291> be set to a TableDetailedErrorArray
	// structure, marshaled as specified in section 2.2.1.17.
	TableDetailedErrors []byte `idl:"name:ppTableDetailedErrors;size_is:(, pcbTableDetailedErrors)" json:"table_detailed_errors"`
	// pcbTableDetailedErrors: A pointer to a pointer variable that, upon completion, MUST
	// be set to the length in bytes of the TableDetailedErrorArray structure returned in
	// ppTableDetailedErrors if such a structure was returned, and MUST be set to zero otherwise.
	TableDetailedErrorsLength uint32 `idl:"name:pcbTableDetailedErrors" json:"table_detailed_errors_length"`
	// Return: The ReadTable return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReadTableResponse) xxx_ToOp(ctx context.Context, op *xxx_ReadTableOperation) *xxx_ReadTableOperation {
	if op == nil {
		op = &xxx_ReadTableOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.TableDataFixed = o.TableDataFixed
	op.TableDataFixedLength = o.TableDataFixedLength
	op.TableDataVariable = o.TableDataVariable
	op.TableDataVariableLength = o.TableDataVariableLength
	op.TableDetailedErrors = o.TableDetailedErrors
	op.TableDetailedErrorsLength = o.TableDetailedErrorsLength
	op.Return = o.Return
	return op
}

func (o *ReadTableResponse) xxx_FromOp(ctx context.Context, op *xxx_ReadTableOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TableDataFixed = op.TableDataFixed
	o.TableDataFixedLength = op.TableDataFixedLength
	o.TableDataVariable = op.TableDataVariable
	o.TableDataVariableLength = op.TableDataVariableLength
	o.TableDetailedErrors = op.TableDetailedErrors
	o.TableDetailedErrorsLength = op.TableDetailedErrorsLength
	o.Return = op.Return
}
func (o *ReadTableResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReadTableResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReadTableOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
