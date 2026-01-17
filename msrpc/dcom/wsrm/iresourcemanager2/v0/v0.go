package iresourcemanager2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
	wsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
	_ = wsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

var (
	// IResourceManager2 interface identifier 2a3eb639-d134-422d-90d8-aaa1b5216202
	ResourceManager2IID = &dcom.IID{Data1: 0x2a3eb639, Data2: 0xd134, Data3: 0x422d, Data4: []byte{0x90, 0xd8, 0xaa, 0xa1, 0xb5, 0x21, 0x62, 0x02}}
	// Syntax UUID
	ResourceManager2SyntaxUUID = &uuid.UUID{TimeLow: 0x2a3eb639, TimeMid: 0xd134, TimeHiAndVersion: 0x422d, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0xd8, Node: [6]uint8{0xaa, 0xa1, 0xb5, 0x21, 0x62, 0x2}}
	// Syntax ID
	ResourceManager2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ResourceManager2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IResourceManager2 interface.
type ResourceManager2Client interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The ExportObjects method creates XML for exporting objects.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | Operation successful.              |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// Additional IResourceManager2 interface methods are specified in section 3.2.4.2.
	//
	// The server SHOULD process this method call as follows.
	//
	// * If one or more ObjectIds XML elements included in the bstrObjectIds parameter cannot
	// be found, they MUST be ignored, and the call MUST proceed for other ObjectIds.
	//
	// * If a PMC ( e371cc74-bf4c-4870-8afe-5062cc628b4f#gt_65daee12-445f-41c0-8456-f728063bef24
	// ) object is specified in the enumObjectType parameter, this method MUST return an
	// XML string containing the requested PMCs in the configuration. The "Type" attributes
	// of individual objects in the ObjectIds XML element MUST be ignored.
	//
	// * If a resource policy object is specified in the enumObjectType parameter, this
	// method MUST return an XML string containing the requested RAPs ( e371cc74-bf4c-4870-8afe-5062cc628b4f#gt_6442eac9-3264-4d95-a435-fdc08e71603f
	// ) in the configuration. The "Type" attributes of individual objects in the ObjectIds
	// XML element MUST be ignored.
	//
	// * If a schedule object is specified in the enumObjectType parameter and the "Type"
	// attribute of the object requested in the ObjectIds XML element is "Calendar", this
	// method MUST return an XML string containing the requested calendar objects in the
	// configuration.
	//
	// * If a schedule object is specified in the enumObjectType parameter and the "Type"
	// attribute of the object requested in the ObjectIds XML element is not "Calendar",
	// this method MUST return an XML string containing the requested schedule objects in
	// the configuration.
	ExportObjects(context.Context, *ExportObjectsRequest, ...dcerpc.CallOption) (*ExportObjectsResponse, error)

	// The GetImportConflicts method finds conflicts between import objects and existing
	// objects.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | Operation successful.                                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | One or more arguments are invalid.                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                      | be processed.<36>                                                                |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IResourceManager2 interface methods are specified in section 3.2.4.2.
	GetImportConflicts(context.Context, *GetImportConflictsRequest, ...dcerpc.CallOption) (*GetImportConflictsResponse, error)

	// The ImportXml method imports objects into the configuration.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                 |                                                                                  |
	//	|               VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                        | Operation successful.                                                            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                | One or more arguments are invalid.                                               |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0272 WRM_ERR_CAL_SCHEDULE_IN_USE | An existing schedule with the same name as the one in the supplied               |
	//	|                                        | CalendarsCollection is currently in use with an existing calendar. The complete  |
	//	|                                        | import process is canceled.<38>                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER   | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                        | be processed.<39>                                                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IResourceManager2 interface methods are specified in section 3.2.4.2.
	ImportXML(context.Context, *ImportXMLRequest, ...dcerpc.CallOption) (*ImportXMLResponse, error)

	// The ExportXml method exports objects from the configuration.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IResourceManager2 interface methods are specified in section 3.2.4.2.
	ExportXML(context.Context, *ExportXMLRequest, ...dcerpc.CallOption) (*ExportXMLResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ResourceManager2Client
}

type xxx_DefaultResourceManager2Client struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultResourceManager2Client) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultResourceManager2Client) ExportObjects(ctx context.Context, in *ExportObjectsRequest, opts ...dcerpc.CallOption) (*ExportObjectsResponse, error) {
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
	out := &ExportObjectsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManager2Client) GetImportConflicts(ctx context.Context, in *GetImportConflictsRequest, opts ...dcerpc.CallOption) (*GetImportConflictsResponse, error) {
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
	out := &GetImportConflictsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManager2Client) ImportXML(ctx context.Context, in *ImportXMLRequest, opts ...dcerpc.CallOption) (*ImportXMLResponse, error) {
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
	out := &ImportXMLResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManager2Client) ExportXML(ctx context.Context, in *ExportXMLRequest, opts ...dcerpc.CallOption) (*ExportXMLResponse, error) {
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
	out := &ExportXMLResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceManager2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultResourceManager2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultResourceManager2Client) IPID(ctx context.Context, ipid *dcom.IPID) ResourceManager2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultResourceManager2Client{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewResourceManager2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ResourceManager2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ResourceManager2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultResourceManager2Client{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_ExportObjectsOperation structure represents the ExportObjects operation
type xxx_ExportObjectsOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ObjectIDs      *oaut.String    `idl:"name:bstrObjectIds" json:"object_ids"`
	EnumObjectType wsrm.ObjectType `idl:"name:enumObjectType" json:"enum_object_type"`
	ObjectXML      *oaut.String    `idl:"name:pbstrObjectXml" json:"object_xml"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_ExportObjectsOperation) OpNum() int { return 7 }

func (o *xxx_ExportObjectsOperation) OpName() string { return "/IResourceManager2/v0/ExportObjects" }

func (o *xxx_ExportObjectsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportObjectsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrObjectIds {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ObjectIDs != nil {
			_ptr_bstrObjectIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectIDs != nil {
					if err := o.ObjectIDs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectIDs, _ptr_bstrObjectIds); err != nil {
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
	// enumObjectType {in} (1:{v1_enum, alias=OBJECT_TYPE}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumObjectType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportObjectsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrObjectIds {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrObjectIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectIDs == nil {
				o.ObjectIDs = &oaut.String{}
			}
			if err := o.ObjectIDs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrObjectIds := func(ptr interface{}) { o.ObjectIDs = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ObjectIDs, _s_bstrObjectIds, _ptr_bstrObjectIds); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// enumObjectType {in} (1:{v1_enum, alias=OBJECT_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumObjectType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportObjectsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportObjectsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrObjectXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ObjectXML != nil {
			_ptr_pbstrObjectXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectXML != nil {
					if err := o.ObjectXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectXML, _ptr_pbstrObjectXml); err != nil {
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

func (o *xxx_ExportObjectsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrObjectXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrObjectXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectXML == nil {
				o.ObjectXML = &oaut.String{}
			}
			if err := o.ObjectXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrObjectXml := func(ptr interface{}) { o.ObjectXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ObjectXML, _s_pbstrObjectXml, _ptr_pbstrObjectXml); err != nil {
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

// ExportObjectsRequest structure represents the ExportObjects operation request
type ExportObjectsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrObjectIds: A string that identifies the objects to be exported, in the format
	// of an ObjectIds element (section 2.2.5.20).
	ObjectIDs *oaut.String `idl:"name:bstrObjectIds" json:"object_ids"`
	// enumObjectType: An OBJECT_TYPE enumeration value (section 2.2.3.6) that specifies
	// the type of the objects to be exported. This determines the format of XML object
	// returned in the pbstrObjectXml parameter, as follows.
	//
	//	+---------------------------+---------------------------------------------------------------------------+
	//	|          OBJECT           |                                    XML                                    |
	//	|           TYPE            |                                  OBJECT                                   |
	//	+---------------------------+---------------------------------------------------------------------------+
	//	+---------------------------+---------------------------------------------------------------------------+
	//	| OBJECT_SELECTION_CRITERIA | ProcessMatchingCriteria element (section 2.2.5.24)                        |
	//	+---------------------------+---------------------------------------------------------------------------+
	//	| OBJECT_POLICY             | Policy element (section 2.2.5.21)                                         |
	//	+---------------------------+---------------------------------------------------------------------------+
	//	| OBJECT_SCHEDULE           | Calendar element (section 2.2.5.7) or Schedule element (section 2.2.5.26) |
	//	+---------------------------+---------------------------------------------------------------------------+
	EnumObjectType wsrm.ObjectType `idl:"name:enumObjectType" json:"enum_object_type"`
}

func (o *ExportObjectsRequest) xxx_ToOp(ctx context.Context, op *xxx_ExportObjectsOperation) *xxx_ExportObjectsOperation {
	if op == nil {
		op = &xxx_ExportObjectsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectIDs = o.ObjectIDs
	op.EnumObjectType = o.EnumObjectType
	return op
}

func (o *ExportObjectsRequest) xxx_FromOp(ctx context.Context, op *xxx_ExportObjectsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectIDs = op.ObjectIDs
	o.EnumObjectType = op.EnumObjectType
}
func (o *ExportObjectsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ExportObjectsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportObjectsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExportObjectsResponse structure represents the ExportObjects operation response
type ExportObjectsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrObjectXml: A pointer to a string that returns the XML for the objects to be
	// exported. The format of the XML depends on the type of exported objects specified
	// by the enumObjectType parameter.<35>
	ObjectXML *oaut.String `idl:"name:pbstrObjectXml" json:"object_xml"`
	// Return: The ExportObjects return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExportObjectsResponse) xxx_ToOp(ctx context.Context, op *xxx_ExportObjectsOperation) *xxx_ExportObjectsOperation {
	if op == nil {
		op = &xxx_ExportObjectsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ObjectXML = o.ObjectXML
	op.Return = o.Return
	return op
}

func (o *ExportObjectsResponse) xxx_FromOp(ctx context.Context, op *xxx_ExportObjectsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ObjectXML = op.ObjectXML
	o.Return = op.Return
}
func (o *ExportObjectsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ExportObjectsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportObjectsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetImportConflictsOperation structure represents the GetImportConflicts operation
type xxx_GetImportConflictsOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	PMCXML             *oaut.String   `idl:"name:bstrPMCXml" json:"pmc_xml"`
	PolicyXML          *oaut.String   `idl:"name:bstrPolicyXml" json:"policy_xml"`
	CalendarXML        *oaut.String   `idl:"name:bstrCalendarXml" json:"calendar_xml"`
	ConditionalXML     *oaut.String   `idl:"name:bstrConditionalXml" json:"conditional_xml"`
	MachineGroupXML    *oaut.String   `idl:"name:bstrMachineGroupXml" json:"machine_group_xml"`
	ConfigurationXMLs  *oaut.String   `idl:"name:bstrConfigurationXmls" json:"configuration_xmls"`
	ConflictingObjects *oaut.String   `idl:"name:pbstrConflictingObjects" json:"conflicting_objects"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetImportConflictsOperation) OpNum() int { return 8 }

func (o *xxx_GetImportConflictsOperation) OpName() string {
	return "/IResourceManager2/v0/GetImportConflicts"
}

func (o *xxx_GetImportConflictsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImportConflictsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPMCXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PMCXML != nil {
			_ptr_bstrPMCXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PMCXML != nil {
					if err := o.PMCXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PMCXML, _ptr_bstrPMCXml); err != nil {
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
	// bstrPolicyXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyXML != nil {
			_ptr_bstrPolicyXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyXML != nil {
					if err := o.PolicyXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyXML, _ptr_bstrPolicyXml); err != nil {
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
	// bstrCalendarXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CalendarXML != nil {
			_ptr_bstrCalendarXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CalendarXML != nil {
					if err := o.CalendarXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CalendarXML, _ptr_bstrCalendarXml); err != nil {
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
	// bstrConditionalXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConditionalXML != nil {
			_ptr_bstrConditionalXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConditionalXML != nil {
					if err := o.ConditionalXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConditionalXML, _ptr_bstrConditionalXml); err != nil {
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
	// bstrMachineGroupXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineGroupXML != nil {
			_ptr_bstrMachineGroupXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineGroupXML != nil {
					if err := o.MachineGroupXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineGroupXML, _ptr_bstrMachineGroupXml); err != nil {
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
	// bstrConfigurationXmls {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConfigurationXMLs != nil {
			_ptr_bstrConfigurationXmls := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigurationXMLs != nil {
					if err := o.ConfigurationXMLs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigurationXMLs, _ptr_bstrConfigurationXmls); err != nil {
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

func (o *xxx_GetImportConflictsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPMCXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPMCXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PMCXML == nil {
				o.PMCXML = &oaut.String{}
			}
			if err := o.PMCXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPMCXml := func(ptr interface{}) { o.PMCXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PMCXML, _s_bstrPMCXml, _ptr_bstrPMCXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPolicyXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyXML == nil {
				o.PolicyXML = &oaut.String{}
			}
			if err := o.PolicyXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyXml := func(ptr interface{}) { o.PolicyXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyXML, _s_bstrPolicyXml, _ptr_bstrPolicyXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrCalendarXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrCalendarXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CalendarXML == nil {
				o.CalendarXML = &oaut.String{}
			}
			if err := o.CalendarXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrCalendarXml := func(ptr interface{}) { o.CalendarXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CalendarXML, _s_bstrCalendarXml, _ptr_bstrCalendarXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrConditionalXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConditionalXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConditionalXML == nil {
				o.ConditionalXML = &oaut.String{}
			}
			if err := o.ConditionalXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConditionalXml := func(ptr interface{}) { o.ConditionalXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConditionalXML, _s_bstrConditionalXml, _ptr_bstrConditionalXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMachineGroupXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineGroupXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineGroupXML == nil {
				o.MachineGroupXML = &oaut.String{}
			}
			if err := o.MachineGroupXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineGroupXml := func(ptr interface{}) { o.MachineGroupXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineGroupXML, _s_bstrMachineGroupXml, _ptr_bstrMachineGroupXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrConfigurationXmls {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConfigurationXmls := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigurationXMLs == nil {
				o.ConfigurationXMLs = &oaut.String{}
			}
			if err := o.ConfigurationXMLs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConfigurationXmls := func(ptr interface{}) { o.ConfigurationXMLs = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConfigurationXMLs, _s_bstrConfigurationXmls, _ptr_bstrConfigurationXmls); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImportConflictsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImportConflictsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrConflictingObjects {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConflictingObjects != nil {
			_ptr_pbstrConflictingObjects := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConflictingObjects != nil {
					if err := o.ConflictingObjects.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConflictingObjects, _ptr_pbstrConflictingObjects); err != nil {
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

func (o *xxx_GetImportConflictsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrConflictingObjects {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrConflictingObjects := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConflictingObjects == nil {
				o.ConflictingObjects = &oaut.String{}
			}
			if err := o.ConflictingObjects.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrConflictingObjects := func(ptr interface{}) { o.ConflictingObjects = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConflictingObjects, _s_pbstrConflictingObjects, _ptr_pbstrConflictingObjects); err != nil {
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

// GetImportConflictsRequest structure represents the GetImportConflicts operation request
type GetImportConflictsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrPMCXml: An optional string that specifies process matching criteria (PMC), in
	// the form of a ProcessMatchingCriteriaCollection element (section 2.2.5.25).
	PMCXML *oaut.String `idl:"name:bstrPMCXml" json:"pmc_xml"`
	// bstrPolicyXml: An optional string that specifies resource policies, in the form of
	// a PolicyCollection element (section 2.2.5.22).
	PolicyXML *oaut.String `idl:"name:bstrPolicyXml" json:"policy_xml"`
	// bstrCalendarXml: An optional string that specifies calendar elements, in the form
	// of a CalendarsCollection element (section 2.2.5.11).
	CalendarXML *oaut.String `idl:"name:bstrCalendarXml" json:"calendar_xml"`
	// bstrConditionalXml: An optional string that specifies a conditional policy, in the
	// form of a ConditionalPolicy element (section 2.2.5.12).
	ConditionalXML *oaut.String `idl:"name:bstrConditionalXml" json:"conditional_xml"`
	// bstrMachineGroupXml: An optional string that specifies a machine group, in the form
	// of a MachineGroup element (section 2.2.5.18).
	MachineGroupXML *oaut.String `idl:"name:bstrMachineGroupXml" json:"machine_group_xml"`
	// bstrConfigurationXmls: Not used.
	ConfigurationXMLs *oaut.String `idl:"name:bstrConfigurationXmls" json:"configuration_xmls"`
}

func (o *GetImportConflictsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetImportConflictsOperation) *xxx_GetImportConflictsOperation {
	if op == nil {
		op = &xxx_GetImportConflictsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PMCXML = o.PMCXML
	op.PolicyXML = o.PolicyXML
	op.CalendarXML = o.CalendarXML
	op.ConditionalXML = o.ConditionalXML
	op.MachineGroupXML = o.MachineGroupXML
	op.ConfigurationXMLs = o.ConfigurationXMLs
	return op
}

func (o *GetImportConflictsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetImportConflictsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PMCXML = op.PMCXML
	o.PolicyXML = op.PolicyXML
	o.CalendarXML = op.CalendarXML
	o.ConditionalXML = op.ConditionalXML
	o.MachineGroupXML = op.MachineGroupXML
	o.ConfigurationXMLs = op.ConfigurationXMLs
}
func (o *GetImportConflictsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetImportConflictsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetImportConflictsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetImportConflictsResponse structure represents the GetImportConflicts operation response
type GetImportConflictsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrConflictingObjects: A pointer to a string that SHOULD identify the conflicting
	// objects, in the form of an ObjectIds element (section 2.2.5.20).
	ConflictingObjects *oaut.String `idl:"name:pbstrConflictingObjects" json:"conflicting_objects"`
	// Return: The GetImportConflicts return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetImportConflictsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetImportConflictsOperation) *xxx_GetImportConflictsOperation {
	if op == nil {
		op = &xxx_GetImportConflictsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ConflictingObjects = o.ConflictingObjects
	op.Return = o.Return
	return op
}

func (o *GetImportConflictsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetImportConflictsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ConflictingObjects = op.ConflictingObjects
	o.Return = op.Return
}
func (o *GetImportConflictsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetImportConflictsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetImportConflictsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ImportXMLOperation structure represents the ImportXml operation
type xxx_ImportXMLOperation struct {
	This              *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat  `idl:"name:That" json:"that"`
	PMCXML            *oaut.String    `idl:"name:bstrPMCXml" json:"pmc_xml"`
	PolicyXML         *oaut.String    `idl:"name:bstrPolicyXml" json:"policy_xml"`
	CalendarXML       *oaut.String    `idl:"name:bstrCalendarXml" json:"calendar_xml"`
	ConditionalXML    *oaut.String    `idl:"name:bstrConditionalXml" json:"conditional_xml"`
	MachineGroupXML   *oaut.String    `idl:"name:bstrMachineGroupXml" json:"machine_group_xml"`
	ConfigurationXMLs *oaut.String    `idl:"name:bstrConfigurationXmls" json:"configuration_xmls"`
	EnumImportType    wsrm.ImportType `idl:"name:enumImportType" json:"enum_import_type"`
	Return            int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_ImportXMLOperation) OpNum() int { return 9 }

func (o *xxx_ImportXMLOperation) OpName() string { return "/IResourceManager2/v0/ImportXml" }

func (o *xxx_ImportXMLOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportXMLOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPMCXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PMCXML != nil {
			_ptr_bstrPMCXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PMCXML != nil {
					if err := o.PMCXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PMCXML, _ptr_bstrPMCXml); err != nil {
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
	// bstrPolicyXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyXML != nil {
			_ptr_bstrPolicyXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyXML != nil {
					if err := o.PolicyXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyXML, _ptr_bstrPolicyXml); err != nil {
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
	// bstrCalendarXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CalendarXML != nil {
			_ptr_bstrCalendarXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CalendarXML != nil {
					if err := o.CalendarXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CalendarXML, _ptr_bstrCalendarXml); err != nil {
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
	// bstrConditionalXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConditionalXML != nil {
			_ptr_bstrConditionalXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConditionalXML != nil {
					if err := o.ConditionalXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConditionalXML, _ptr_bstrConditionalXml); err != nil {
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
	// bstrMachineGroupXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineGroupXML != nil {
			_ptr_bstrMachineGroupXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineGroupXML != nil {
					if err := o.MachineGroupXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineGroupXML, _ptr_bstrMachineGroupXml); err != nil {
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
	// bstrConfigurationXmls {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConfigurationXMLs != nil {
			_ptr_bstrConfigurationXmls := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigurationXMLs != nil {
					if err := o.ConfigurationXMLs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigurationXMLs, _ptr_bstrConfigurationXmls); err != nil {
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
	// enumImportType {in} (1:{v1_enum, alias=IMPORT_TYPE}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumImportType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportXMLOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPMCXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPMCXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PMCXML == nil {
				o.PMCXML = &oaut.String{}
			}
			if err := o.PMCXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPMCXml := func(ptr interface{}) { o.PMCXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PMCXML, _s_bstrPMCXml, _ptr_bstrPMCXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPolicyXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyXML == nil {
				o.PolicyXML = &oaut.String{}
			}
			if err := o.PolicyXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyXml := func(ptr interface{}) { o.PolicyXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyXML, _s_bstrPolicyXml, _ptr_bstrPolicyXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrCalendarXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrCalendarXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CalendarXML == nil {
				o.CalendarXML = &oaut.String{}
			}
			if err := o.CalendarXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrCalendarXml := func(ptr interface{}) { o.CalendarXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CalendarXML, _s_bstrCalendarXml, _ptr_bstrCalendarXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrConditionalXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConditionalXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConditionalXML == nil {
				o.ConditionalXML = &oaut.String{}
			}
			if err := o.ConditionalXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConditionalXml := func(ptr interface{}) { o.ConditionalXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConditionalXML, _s_bstrConditionalXml, _ptr_bstrConditionalXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMachineGroupXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineGroupXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineGroupXML == nil {
				o.MachineGroupXML = &oaut.String{}
			}
			if err := o.MachineGroupXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineGroupXml := func(ptr interface{}) { o.MachineGroupXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineGroupXML, _s_bstrMachineGroupXml, _ptr_bstrMachineGroupXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrConfigurationXmls {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConfigurationXmls := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigurationXMLs == nil {
				o.ConfigurationXMLs = &oaut.String{}
			}
			if err := o.ConfigurationXMLs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConfigurationXmls := func(ptr interface{}) { o.ConfigurationXMLs = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConfigurationXMLs, _s_bstrConfigurationXmls, _ptr_bstrConfigurationXmls); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// enumImportType {in} (1:{v1_enum, alias=IMPORT_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumImportType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportXMLOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportXMLOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportXMLOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ImportXMLRequest structure represents the ImportXml operation request
type ImportXMLRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrPMCXml: A string that specifies process matching criteria (PMC), in the form
	// of a ProcessMatchingCriteriaCollection element (section 2.2.5.25). For an example,
	// see ProcessMatchingCriteriaCollection example (section 4.2.21).
	PMCXML *oaut.String `idl:"name:bstrPMCXml" json:"pmc_xml"`
	// bstrPolicyXml: A string that specifies a resource policy, in the form of a PolicyCollection
	// element (section 2.2.5.22).
	PolicyXML *oaut.String `idl:"name:bstrPolicyXml" json:"policy_xml"`
	// bstrCalendarXml: A string that specifies a calendar, in the form of a CalendarsCollection
	// element (section 2.2.5.11). For an example, see CalendarsCollection example (section
	// 4.2.8).<37>
	CalendarXML *oaut.String `idl:"name:bstrCalendarXml" json:"calendar_xml"`
	// bstrConditionalXml: A string that specifies a conditional policy, in the form of
	// a ConditionalPolicy element (section 2.2.5.12).
	ConditionalXML *oaut.String `idl:"name:bstrConditionalXml" json:"conditional_xml"`
	// bstrMachineGroupXml: A string that specifies a machine group, in the form of a MachineGroup
	// element (section 2.2.5.17).
	MachineGroupXML *oaut.String `idl:"name:bstrMachineGroupXml" json:"machine_group_xml"`
	// bstrConfigurationXmls: A string that specifies a configuration to be loaded by the
	// WSRM server, in the form of a ConfigurationFiles element (section 2.2.5.13).
	ConfigurationXMLs *oaut.String `idl:"name:bstrConfigurationXmls" json:"configuration_xmls"`
	// enumImportType: An IMPORT_TYPE enumeration value (section 2.2.3.6) that specifies
	// the mode in which to handle conflicting objects.
	EnumImportType wsrm.ImportType `idl:"name:enumImportType" json:"enum_import_type"`
}

func (o *ImportXMLRequest) xxx_ToOp(ctx context.Context, op *xxx_ImportXMLOperation) *xxx_ImportXMLOperation {
	if op == nil {
		op = &xxx_ImportXMLOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PMCXML = o.PMCXML
	op.PolicyXML = o.PolicyXML
	op.CalendarXML = o.CalendarXML
	op.ConditionalXML = o.ConditionalXML
	op.MachineGroupXML = o.MachineGroupXML
	op.ConfigurationXMLs = o.ConfigurationXMLs
	op.EnumImportType = o.EnumImportType
	return op
}

func (o *ImportXMLRequest) xxx_FromOp(ctx context.Context, op *xxx_ImportXMLOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PMCXML = op.PMCXML
	o.PolicyXML = op.PolicyXML
	o.CalendarXML = op.CalendarXML
	o.ConditionalXML = op.ConditionalXML
	o.MachineGroupXML = op.MachineGroupXML
	o.ConfigurationXMLs = op.ConfigurationXMLs
	o.EnumImportType = op.EnumImportType
}
func (o *ImportXMLRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ImportXMLRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportXMLOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ImportXMLResponse structure represents the ImportXml operation response
type ImportXMLResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ImportXml return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ImportXMLResponse) xxx_ToOp(ctx context.Context, op *xxx_ImportXMLOperation) *xxx_ImportXMLOperation {
	if op == nil {
		op = &xxx_ImportXMLOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ImportXMLResponse) xxx_FromOp(ctx context.Context, op *xxx_ImportXMLOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ImportXMLResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ImportXMLResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportXMLOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExportXMLOperation structure represents the ExportXml operation
type xxx_ExportXMLOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	PMCXML            *oaut.String   `idl:"name:pbstrPMCXml" json:"pmc_xml"`
	PolicyXML         *oaut.String   `idl:"name:pbstrPolicyXml" json:"policy_xml"`
	CalendarXML       *oaut.String   `idl:"name:pbstrCalendarXml" json:"calendar_xml"`
	ConditionalXML    *oaut.String   `idl:"name:pbstrConditionalXml" json:"conditional_xml"`
	MachineGroupXML   *oaut.String   `idl:"name:pbstrMachineGroupXml" json:"machine_group_xml"`
	ConfigurationXMLs *oaut.String   `idl:"name:pbstrConfigurationXmls" json:"configuration_xmls"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ExportXMLOperation) OpNum() int { return 10 }

func (o *xxx_ExportXMLOperation) OpName() string { return "/IResourceManager2/v0/ExportXml" }

func (o *xxx_ExportXMLOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportXMLOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_ExportXMLOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_ExportXMLOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportXMLOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrPMCXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PMCXML != nil {
			_ptr_pbstrPMCXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PMCXML != nil {
					if err := o.PMCXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PMCXML, _ptr_pbstrPMCXml); err != nil {
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
	// pbstrPolicyXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyXML != nil {
			_ptr_pbstrPolicyXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyXML != nil {
					if err := o.PolicyXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyXML, _ptr_pbstrPolicyXml); err != nil {
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
	// pbstrCalendarXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CalendarXML != nil {
			_ptr_pbstrCalendarXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CalendarXML != nil {
					if err := o.CalendarXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CalendarXML, _ptr_pbstrCalendarXml); err != nil {
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
	// pbstrConditionalXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConditionalXML != nil {
			_ptr_pbstrConditionalXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConditionalXML != nil {
					if err := o.ConditionalXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConditionalXML, _ptr_pbstrConditionalXml); err != nil {
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
	// pbstrMachineGroupXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineGroupXML != nil {
			_ptr_pbstrMachineGroupXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineGroupXML != nil {
					if err := o.MachineGroupXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineGroupXML, _ptr_pbstrMachineGroupXml); err != nil {
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
	// pbstrConfigurationXmls {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConfigurationXMLs != nil {
			_ptr_pbstrConfigurationXmls := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigurationXMLs != nil {
					if err := o.ConfigurationXMLs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigurationXMLs, _ptr_pbstrConfigurationXmls); err != nil {
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

func (o *xxx_ExportXMLOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrPMCXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrPMCXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PMCXML == nil {
				o.PMCXML = &oaut.String{}
			}
			if err := o.PMCXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrPMCXml := func(ptr interface{}) { o.PMCXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PMCXML, _s_pbstrPMCXml, _ptr_pbstrPMCXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrPolicyXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrPolicyXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyXML == nil {
				o.PolicyXML = &oaut.String{}
			}
			if err := o.PolicyXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrPolicyXml := func(ptr interface{}) { o.PolicyXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyXML, _s_pbstrPolicyXml, _ptr_pbstrPolicyXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrCalendarXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrCalendarXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CalendarXML == nil {
				o.CalendarXML = &oaut.String{}
			}
			if err := o.CalendarXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrCalendarXml := func(ptr interface{}) { o.CalendarXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CalendarXML, _s_pbstrCalendarXml, _ptr_pbstrCalendarXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrConditionalXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrConditionalXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConditionalXML == nil {
				o.ConditionalXML = &oaut.String{}
			}
			if err := o.ConditionalXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrConditionalXml := func(ptr interface{}) { o.ConditionalXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConditionalXML, _s_pbstrConditionalXml, _ptr_pbstrConditionalXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrMachineGroupXml {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrMachineGroupXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineGroupXML == nil {
				o.MachineGroupXML = &oaut.String{}
			}
			if err := o.MachineGroupXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrMachineGroupXml := func(ptr interface{}) { o.MachineGroupXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineGroupXML, _s_pbstrMachineGroupXml, _ptr_pbstrMachineGroupXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrConfigurationXmls {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrConfigurationXmls := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigurationXMLs == nil {
				o.ConfigurationXMLs = &oaut.String{}
			}
			if err := o.ConfigurationXMLs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrConfigurationXmls := func(ptr interface{}) { o.ConfigurationXMLs = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConfigurationXMLs, _s_pbstrConfigurationXmls, _ptr_pbstrConfigurationXmls); err != nil {
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

// ExportXMLRequest structure represents the ExportXml operation request
type ExportXMLRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ExportXMLRequest) xxx_ToOp(ctx context.Context, op *xxx_ExportXMLOperation) *xxx_ExportXMLOperation {
	if op == nil {
		op = &xxx_ExportXMLOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ExportXMLRequest) xxx_FromOp(ctx context.Context, op *xxx_ExportXMLOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ExportXMLRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ExportXMLRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportXMLOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExportXMLResponse structure represents the ExportXml operation response
type ExportXMLResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrPMCXml: A pointer to a string that returns process matching criteria (PMC),
	// in the form of a ProcessMatchingCriteriaCollection element (section 2.2.5.25). For
	// an example, see ProcessMatchingCriteriaCollection example (section 4.2.21).
	PMCXML *oaut.String `idl:"name:pbstrPMCXml" json:"pmc_xml"`
	// pbstrPolicyXml: A pointer to a string that returns a resource policy, in the form
	// of a PolicyCollection element (section 2.2.5.22).
	PolicyXML *oaut.String `idl:"name:pbstrPolicyXml" json:"policy_xml"`
	// pbstrCalendarXml: A pointer to a string that returns a calendar, in the form of a
	// CalendarsCollection element (section 2.2.5.11). For an example, see CalendarsCollection
	// example (section 4.2.8).
	CalendarXML *oaut.String `idl:"name:pbstrCalendarXml" json:"calendar_xml"`
	// pbstrConditionalXml: A pointer to a string that SHOULD return a conditional policy,
	// in the format of a ConditionalPolicy element (section 2.2.5.12).
	ConditionalXML *oaut.String `idl:"name:pbstrConditionalXml" json:"conditional_xml"`
	// pbstrMachineGroupXml: A pointer to a string that SHOULD return a machine group, in
	// the format of a MachineGroup element (section 2.2.5.17).
	MachineGroupXML *oaut.String `idl:"name:pbstrMachineGroupXml" json:"machine_group_xml"`
	// pbstrConfigurationXmls: A pointer to a string that SHOULD return a configuration,
	// in the format of a ConfigurationFiles element (section 2.2.5.13).
	ConfigurationXMLs *oaut.String `idl:"name:pbstrConfigurationXmls" json:"configuration_xmls"`
	// Return: The ExportXml return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExportXMLResponse) xxx_ToOp(ctx context.Context, op *xxx_ExportXMLOperation) *xxx_ExportXMLOperation {
	if op == nil {
		op = &xxx_ExportXMLOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PMCXML = o.PMCXML
	op.PolicyXML = o.PolicyXML
	op.CalendarXML = o.CalendarXML
	op.ConditionalXML = o.ConditionalXML
	op.MachineGroupXML = o.MachineGroupXML
	op.ConfigurationXMLs = o.ConfigurationXMLs
	op.Return = o.Return
	return op
}

func (o *ExportXMLResponse) xxx_FromOp(ctx context.Context, op *xxx_ExportXMLOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PMCXML = op.PMCXML
	o.PolicyXML = op.PolicyXML
	o.CalendarXML = op.CalendarXML
	o.ConditionalXML = op.ConditionalXML
	o.MachineGroupXML = op.MachineGroupXML
	o.ConfigurationXMLs = op.ConfigurationXMLs
	o.Return = op.Return
}
func (o *ExportXMLResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ExportXMLResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportXMLOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
