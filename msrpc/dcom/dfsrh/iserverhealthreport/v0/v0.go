package iserverhealthreport

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
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/dfsrh"
)

var (
	// IServerHealthReport interface identifier e65e8028-83e8-491b-9af7-aaf6bd51a0ce
	ServerHealthReportIID = &dcom.IID{Data1: 0xe65e8028, Data2: 0x83e8, Data3: 0x491b, Data4: []byte{0x9a, 0xf7, 0xaa, 0xf6, 0xbd, 0x51, 0xa0, 0xce}}
	// Syntax UUID
	ServerHealthReportSyntaxUUID = &uuid.UUID{TimeLow: 0xe65e8028, TimeMid: 0x83e8, TimeHiAndVersion: 0x491b, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0xf7, Node: [6]uint8{0xaa, 0xf6, 0xbd, 0x51, 0xa0, 0xce}}
	// Syntax ID
	ServerHealthReportSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ServerHealthReportSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IServerHealthReport interface.
type ServerHealthReportClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetReport method retrieves health information for the specified replication group
	// that is hosted on the server in addition to the global health data of the DFS-R service
	// on the server.
	//
	// Return Values: The method MUST return 0 on success; or return an implementation-specific
	// nonzero HRESULT error code, as specified in [MS-ERREF] section 2.1, between 0x80000000
	// and 0xFFFFFFFF on failure. For protocol purposes, all nonzero values MUST be treated
	// as equivalent failures.
	GetReport(context.Context, *GetReportRequest, ...dcerpc.CallOption) (*GetReportResponse, error)

	// The GetCompressedReport method gets the health information for the specified replication
	// group and the global health data of the DFS-R service on the server. The server MUST
	// encode the report as a field in the format that is specified by the DFS-R compression
	// algorithm in [MS-FRS2] section 3.1.1.1.
	//
	// Return Values: The method MUST return zero on success, or an implementation-specific
	// nonzero HRESULT error code, as specified in [MS-ERREF] section 2.1, between 0x80000000
	// and 0xFFFFFFFF on failure. For protocol purposes, all nonzero values MUST be treated
	// as equivalent failures.
	GetCompressedReport(context.Context, *GetCompressedReportRequest, ...dcerpc.CallOption) (*GetCompressedReportResponse, error)

	// The GetRawReportEx method is not currently in use and has never been implemented
	// in any version of the DFS-R Helper Protocol. It is reserved for future use.
	//
	// Return Values: The server MUST return the E_NOTIMPL error code (numeric value 0x80004001)
	// and take no action.<60>
	//
	//	+----------------------+------------------+
	//	|        RETURN        |                  |
	//	|      VALUE/CODE      |   DESCRIPTION    |
	//	|                      |                  |
	//	+----------------------+------------------+
	//	+----------------------+------------------+
	//	| 0x80004001 E_NOTIMPL | Not implemented. |
	//	+----------------------+------------------+
	GetRawReportEx(context.Context, *GetRawReportExRequest, ...dcerpc.CallOption) (*GetRawReportExResponse, error)

	// The GetReferenceVersionVectors method gets the version vectors for all replicated
	// folders in the specified replication group.
	//
	// Return Values: The method MUST return zero on success or return an implementation-specific
	// nonzero HRESULT error code, as specified in [MS-ERREF] section 2.1, between 0x80000000
	// and 0xFFFFFFFF on failure. For protocol purposes, all nonzero values MUST be treated
	// as equivalent failures.
	GetReferenceVersionVectors(context.Context, *GetReferenceVersionVectorsRequest, ...dcerpc.CallOption) (*GetReferenceVersionVectorsResponse, error)

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// The GetReferenceBacklogCounts method gets the outbound backlog for a replicated folder
	// on the member, relative to specific version vectors.
	//
	// Return Values: The method MUST return 0 on success or return an implementation-specific
	// nonzero HRESULT error code, as specified in [MS-ERREF] section 2.1, between 0x80000000
	// and 0xFFFFFFFF on failure. For protocol purposes, all nonzero values MUST be treated
	// as equivalent failures.
	//
	// After the server receives this message, it MUST get the backlog count for each version
	// vector that is supplied in the message parameters. If the server fails to retrieve
	// a backlog count, it returns a special value in the backlogCounts array at an index
	// that corresponds to the index in the flatMemberVersionVectors for the entry that
	// was used as input. The overall method MAY still return success (S_OK). <62> These
	// special values are as follows:
	//
	// * BACKLOG_CONTENT_SET_NOT_PRESENT (0xffffffff): The content set is not present in
	// DFS-R.
	//
	// * BACKLOG_ERROR_GET_BACKLOG_FAILED (0xfffffffe): Describes one or more of the following
	// conditions:
	//
	// * Run-time errors or implementation-specific errors that prevent the calculation
	// of the backlog count.
	//
	// * The flat member version vector could not be decompressed by using DFS-R. The DFS-R
	// decompression algorithm is specified in [MS-FRS2] ( ../ms-frs2/9914bdd4-9579-43a7-9f2d-9efe2e162944
	// ) section 3.1.1.2 ( ../ms-frs2/8cb5bae9-edf3-4833-9f0a-9d7e24218d3d ).
	//
	// * The version vector is empty (has a length of zero).
	//
	// The backlog counts MUST be saved in the backlogCounts output parameter.
	GetReferenceBacklogCounts(context.Context, *GetReferenceBacklogCountsRequest, ...dcerpc.CallOption) (*GetReferenceBacklogCountsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ServerHealthReportClient
}

type xxx_DefaultServerHealthReportClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultServerHealthReportClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultServerHealthReportClient) GetReport(ctx context.Context, in *GetReportRequest, opts ...dcerpc.CallOption) (*GetReportResponse, error) {
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
	out := &GetReportResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServerHealthReportClient) GetCompressedReport(ctx context.Context, in *GetCompressedReportRequest, opts ...dcerpc.CallOption) (*GetCompressedReportResponse, error) {
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
	out := &GetCompressedReportResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServerHealthReportClient) GetRawReportEx(ctx context.Context, in *GetRawReportExRequest, opts ...dcerpc.CallOption) (*GetRawReportExResponse, error) {
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
	out := &GetRawReportExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServerHealthReportClient) GetReferenceVersionVectors(ctx context.Context, in *GetReferenceVersionVectorsRequest, opts ...dcerpc.CallOption) (*GetReferenceVersionVectorsResponse, error) {
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
	out := &GetReferenceVersionVectorsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServerHealthReportClient) GetReferenceBacklogCounts(ctx context.Context, in *GetReferenceBacklogCountsRequest, opts ...dcerpc.CallOption) (*GetReferenceBacklogCountsResponse, error) {
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
	out := &GetReferenceBacklogCountsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServerHealthReportClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultServerHealthReportClient) IPID(ctx context.Context, ipid *dcom.IPID) ServerHealthReportClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultServerHealthReportClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewServerHealthReportClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ServerHealthReportClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ServerHealthReportSyntaxV0_0))...)
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
	return &xxx_DefaultServerHealthReportClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetReportOperation structure represents the GetReport operation
type xxx_GetReportOperation struct {
	This                    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ReplicationGroupGUID    *dtyp.GUID      `idl:"name:replicationGroupGuid" json:"replication_group_guid"`
	ReferenceMember         *oaut.String    `idl:"name:referenceMember" json:"reference_member"`
	ReferenceVersionVectors *oaut.SafeArray `idl:"name:referenceVersionVectors" json:"reference_version_vectors"`
	Flags                   int32           `idl:"name:flags" json:"flags"`
	MemberVersionVectors    *oaut.SafeArray `idl:"name:memberVersionVectors" json:"member_version_vectors"`
	ReportXML               *oaut.String    `idl:"name:reportXML" json:"report_xml"`
	Return                  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReportOperation) OpNum() int { return 3 }

func (o *xxx_GetReportOperation) OpName() string { return "/IServerHealthReport/v0/GetReport" }

func (o *xxx_GetReportOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// replicationGroupGuid {in} (1:{alias=GUID}(struct))
	{
		if o.ReplicationGroupGUID != nil {
			if err := o.ReplicationGroupGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// referenceMember {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReferenceMember != nil {
			_ptr_referenceMember := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferenceMember != nil {
					if err := o.ReferenceMember.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferenceMember, _ptr_referenceMember); err != nil {
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
	// referenceVersionVectors {in} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.ReferenceVersionVectors != nil {
			_ptr_referenceVersionVectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferenceVersionVectors != nil {
					_ptr_referenceVersionVectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.ReferenceVersionVectors != nil {
							if err := o.ReferenceVersionVectors.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.ReferenceVersionVectors, _ptr_referenceVersionVectors); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferenceVersionVectors, _ptr_referenceVersionVectors); err != nil {
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
	// flags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// replicationGroupGuid {in} (1:{alias=GUID}(struct))
	{
		if o.ReplicationGroupGUID == nil {
			o.ReplicationGroupGUID = &dtyp.GUID{}
		}
		if err := o.ReplicationGroupGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// referenceMember {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_referenceMember := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferenceMember == nil {
				o.ReferenceMember = &oaut.String{}
			}
			if err := o.ReferenceMember.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_referenceMember := func(ptr interface{}) { o.ReferenceMember = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReferenceMember, _s_referenceMember, _ptr_referenceMember); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// referenceVersionVectors {in} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_referenceVersionVectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_referenceVersionVectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.ReferenceVersionVectors == nil {
					o.ReferenceVersionVectors = &oaut.SafeArray{}
				}
				if err := o.ReferenceVersionVectors.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_referenceVersionVectors := func(ptr interface{}) { o.ReferenceVersionVectors = *ptr.(**oaut.SafeArray) }
			if err := w.ReadPointer(&o.ReferenceVersionVectors, _s_referenceVersionVectors, _ptr_referenceVersionVectors); err != nil {
				return err
			}
			return nil
		})
		_s_referenceVersionVectors := func(ptr interface{}) { o.ReferenceVersionVectors = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.ReferenceVersionVectors, _s_referenceVersionVectors, _ptr_referenceVersionVectors); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// memberVersionVectors {out} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.MemberVersionVectors != nil {
			_ptr_memberVersionVectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MemberVersionVectors != nil {
					_ptr_memberVersionVectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.MemberVersionVectors != nil {
							if err := o.MemberVersionVectors.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.MemberVersionVectors, _ptr_memberVersionVectors); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MemberVersionVectors, _ptr_memberVersionVectors); err != nil {
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
	// reportXML {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReportXML != nil {
			_ptr_reportXML := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReportXML != nil {
					if err := o.ReportXML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReportXML, _ptr_reportXML); err != nil {
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

func (o *xxx_GetReportOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// memberVersionVectors {out} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_memberVersionVectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_memberVersionVectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.MemberVersionVectors == nil {
					o.MemberVersionVectors = &oaut.SafeArray{}
				}
				if err := o.MemberVersionVectors.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_memberVersionVectors := func(ptr interface{}) { o.MemberVersionVectors = *ptr.(**oaut.SafeArray) }
			if err := w.ReadPointer(&o.MemberVersionVectors, _s_memberVersionVectors, _ptr_memberVersionVectors); err != nil {
				return err
			}
			return nil
		})
		_s_memberVersionVectors := func(ptr interface{}) { o.MemberVersionVectors = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.MemberVersionVectors, _s_memberVersionVectors, _ptr_memberVersionVectors); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reportXML {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_reportXML := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReportXML == nil {
				o.ReportXML = &oaut.String{}
			}
			if err := o.ReportXML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_reportXML := func(ptr interface{}) { o.ReportXML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReportXML, _s_reportXML, _ptr_reportXML); err != nil {
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

// GetReportRequest structure represents the GetReport operation request
type GetReportRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// replicationGroupGuid: MUST be the identifier of the replication group for which the
	// server gets health information. This field corresponds to the objectGUID field of
	// the msDFSR-ReplicationGroup configuration object in Active Directory. The msDFSR-ReplicationGroup
	// is specified in [MS-FRS2] section 2.3.5.
	ReplicationGroupGUID *dtyp.GUID `idl:"name:replicationGroupGuid" json:"replication_group_guid"`
	// referenceMember: MUST be set to NULL.
	ReferenceMember *oaut.String `idl:"name:referenceMember" json:"reference_member"`
	// referenceVersionVectors: If the flags parameter has REPORTING_FLAGS_BACKLOG set,
	// the set of version vectors for replicated folders on the reference member MUST be
	// passed by using this parameter. Otherwise, this parameter MUST be set to NULL. The
	// VersionVectorData structure is specified in section 2.2.1.4.
	ReferenceVersionVectors *oaut.SafeArray `idl:"name:referenceVersionVectors" json:"reference_version_vectors"`
	// flags: Any values of the DfsrReportingFlags enumeration MUST be combined together
	// by using a bitwise OR operation. For more information about DfsrReportingFlags, see
	// section 2.2.1.2.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|                         |                                                                                  |
	//	|          VALUE          |                                     MEANING                                      |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| REPORTING_FLAGS_NONE    | When this value is set, the server MUST NOT return any optional information.     |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| REPORTING_FLAGS_BACKLOG | In addition to the default reporting information, when this value is set, the    |
	//	|                         | server MUST return the count of backlogged transactions.                         |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| REPORTING_FLAGS_FILES   | In addition to the default reporting information, when this value is set, the    |
	//	|                         | server MUST return the information about the count and cumulative size of files  |
	//	|                         | in the replicated folders.                                                       |
	//	+-------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:flags" json:"flags"`
}

func (o *GetReportRequest) xxx_ToOp(ctx context.Context) *xxx_GetReportOperation {
	if o == nil {
		return &xxx_GetReportOperation{}
	}
	return &xxx_GetReportOperation{
		This:                    o.This,
		ReplicationGroupGUID:    o.ReplicationGroupGUID,
		ReferenceMember:         o.ReferenceMember,
		ReferenceVersionVectors: o.ReferenceVersionVectors,
		Flags:                   o.Flags,
	}
}

func (o *GetReportRequest) xxx_FromOp(ctx context.Context, op *xxx_GetReportOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReplicationGroupGUID = op.ReplicationGroupGUID
	o.ReferenceMember = op.ReferenceMember
	o.ReferenceVersionVectors = op.ReferenceVersionVectors
	o.Flags = op.Flags
}
func (o *GetReportRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetReportRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReportResponse structure represents the GetReport operation response
type GetReportResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// memberVersionVectors: If the flags parameter has REPORTING_FLAGS_BACKLOG set, the
	// set of version vectors for the replicated folders on the server MUST be returned
	// in this output parameter. The VersionVectorData structure is specified in section
	// 2.2.1.4.
	MemberVersionVectors *oaut.SafeArray `idl:"name:memberVersionVectors" json:"member_version_vectors"`
	// reportXML: The report body in the XML format MUST be returned in this output parameter.
	// The report body MUST follow the XML format described in section 2.2.1.5.
	ReportXML *oaut.String `idl:"name:reportXML" json:"report_xml"`
	// Return: The GetReport return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReportResponse) xxx_ToOp(ctx context.Context) *xxx_GetReportOperation {
	if o == nil {
		return &xxx_GetReportOperation{}
	}
	return &xxx_GetReportOperation{
		That:                 o.That,
		MemberVersionVectors: o.MemberVersionVectors,
		ReportXML:            o.ReportXML,
		Return:               o.Return,
	}
}

func (o *GetReportResponse) xxx_FromOp(ctx context.Context, op *xxx_GetReportOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MemberVersionVectors = op.MemberVersionVectors
	o.ReportXML = op.ReportXML
	o.Return = op.Return
}
func (o *GetReportResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetReportResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCompressedReportOperation structure represents the GetCompressedReport operation
type xxx_GetCompressedReportOperation struct {
	This                    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ReplicationGroupGUID    *dtyp.GUID      `idl:"name:replicationGroupGuid" json:"replication_group_guid"`
	ReferenceMember         *oaut.String    `idl:"name:referenceMember" json:"reference_member"`
	ReferenceVersionVectors *oaut.SafeArray `idl:"name:referenceVersionVectors" json:"reference_version_vectors"`
	Flags                   int32           `idl:"name:flags" json:"flags"`
	MemberVersionVectors    *oaut.SafeArray `idl:"name:memberVersionVectors" json:"member_version_vectors"`
	ReportCompressed        *oaut.String    `idl:"name:reportCompressed" json:"report_compressed"`
	UncompressedReportSize  int32           `idl:"name:uncompressedReportSize" json:"uncompressed_report_size"`
	Return                  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCompressedReportOperation) OpNum() int { return 4 }

func (o *xxx_GetCompressedReportOperation) OpName() string {
	return "/IServerHealthReport/v0/GetCompressedReport"
}

func (o *xxx_GetCompressedReportOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCompressedReportOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// replicationGroupGuid {in} (1:{alias=GUID}(struct))
	{
		if o.ReplicationGroupGUID != nil {
			if err := o.ReplicationGroupGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// referenceMember {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReferenceMember != nil {
			_ptr_referenceMember := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferenceMember != nil {
					if err := o.ReferenceMember.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferenceMember, _ptr_referenceMember); err != nil {
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
	// referenceVersionVectors {in} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.ReferenceVersionVectors != nil {
			_ptr_referenceVersionVectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferenceVersionVectors != nil {
					_ptr_referenceVersionVectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.ReferenceVersionVectors != nil {
							if err := o.ReferenceVersionVectors.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.ReferenceVersionVectors, _ptr_referenceVersionVectors); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferenceVersionVectors, _ptr_referenceVersionVectors); err != nil {
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
	// flags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCompressedReportOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// replicationGroupGuid {in} (1:{alias=GUID}(struct))
	{
		if o.ReplicationGroupGUID == nil {
			o.ReplicationGroupGUID = &dtyp.GUID{}
		}
		if err := o.ReplicationGroupGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// referenceMember {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_referenceMember := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferenceMember == nil {
				o.ReferenceMember = &oaut.String{}
			}
			if err := o.ReferenceMember.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_referenceMember := func(ptr interface{}) { o.ReferenceMember = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReferenceMember, _s_referenceMember, _ptr_referenceMember); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// referenceVersionVectors {in} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_referenceVersionVectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_referenceVersionVectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.ReferenceVersionVectors == nil {
					o.ReferenceVersionVectors = &oaut.SafeArray{}
				}
				if err := o.ReferenceVersionVectors.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_referenceVersionVectors := func(ptr interface{}) { o.ReferenceVersionVectors = *ptr.(**oaut.SafeArray) }
			if err := w.ReadPointer(&o.ReferenceVersionVectors, _s_referenceVersionVectors, _ptr_referenceVersionVectors); err != nil {
				return err
			}
			return nil
		})
		_s_referenceVersionVectors := func(ptr interface{}) { o.ReferenceVersionVectors = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.ReferenceVersionVectors, _s_referenceVersionVectors, _ptr_referenceVersionVectors); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCompressedReportOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCompressedReportOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// memberVersionVectors {out} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.MemberVersionVectors != nil {
			_ptr_memberVersionVectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MemberVersionVectors != nil {
					_ptr_memberVersionVectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.MemberVersionVectors != nil {
							if err := o.MemberVersionVectors.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.MemberVersionVectors, _ptr_memberVersionVectors); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MemberVersionVectors, _ptr_memberVersionVectors); err != nil {
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
	// reportCompressed {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReportCompressed != nil {
			_ptr_reportCompressed := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReportCompressed != nil {
					if err := o.ReportCompressed.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReportCompressed, _ptr_reportCompressed); err != nil {
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
	// uncompressedReportSize {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.UncompressedReportSize); err != nil {
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

func (o *xxx_GetCompressedReportOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// memberVersionVectors {out} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_memberVersionVectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_memberVersionVectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.MemberVersionVectors == nil {
					o.MemberVersionVectors = &oaut.SafeArray{}
				}
				if err := o.MemberVersionVectors.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_memberVersionVectors := func(ptr interface{}) { o.MemberVersionVectors = *ptr.(**oaut.SafeArray) }
			if err := w.ReadPointer(&o.MemberVersionVectors, _s_memberVersionVectors, _ptr_memberVersionVectors); err != nil {
				return err
			}
			return nil
		})
		_s_memberVersionVectors := func(ptr interface{}) { o.MemberVersionVectors = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.MemberVersionVectors, _s_memberVersionVectors, _ptr_memberVersionVectors); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reportCompressed {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_reportCompressed := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReportCompressed == nil {
				o.ReportCompressed = &oaut.String{}
			}
			if err := o.ReportCompressed.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_reportCompressed := func(ptr interface{}) { o.ReportCompressed = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReportCompressed, _s_reportCompressed, _ptr_reportCompressed); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// uncompressedReportSize {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.UncompressedReportSize); err != nil {
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

// GetCompressedReportRequest structure represents the GetCompressedReport operation request
type GetCompressedReportRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// replicationGroupGuid: MUST be the identifier of the replication group for which the
	// server gets health information. This field corresponds to the objectGUID field of
	// the msDFSR-ReplicationGroup configuration object in Active Directory. The msDFSR-ReplicationGroup
	// is specified in [MS-FRS2] section 2.3.5.
	ReplicationGroupGUID *dtyp.GUID `idl:"name:replicationGroupGuid" json:"replication_group_guid"`
	// referenceMember: MUST be set to NULL.
	ReferenceMember *oaut.String `idl:"name:referenceMember" json:"reference_member"`
	// referenceVersionVectors: If the flags parameter has REPORTING_FLAGS_BACKLOG set,
	// the set of version vectors for replicated folders on the reference member MUST be
	// passed by using this parameter. Otherwise, this parameter MUST be set to NULL. The
	// VersionVectorData structure is specified in section 2.2.1.4.
	ReferenceVersionVectors *oaut.SafeArray `idl:"name:referenceVersionVectors" json:"reference_version_vectors"`
	// flags: MUST be zero or more combinations of values of DfsrReportingFlags enumeration.
	// The DfsrReportingFlags enumeration is specified in section 2.2.1.2.
	Flags int32 `idl:"name:flags" json:"flags"`
}

func (o *GetCompressedReportRequest) xxx_ToOp(ctx context.Context) *xxx_GetCompressedReportOperation {
	if o == nil {
		return &xxx_GetCompressedReportOperation{}
	}
	return &xxx_GetCompressedReportOperation{
		This:                    o.This,
		ReplicationGroupGUID:    o.ReplicationGroupGUID,
		ReferenceMember:         o.ReferenceMember,
		ReferenceVersionVectors: o.ReferenceVersionVectors,
		Flags:                   o.Flags,
	}
}

func (o *GetCompressedReportRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCompressedReportOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReplicationGroupGUID = op.ReplicationGroupGUID
	o.ReferenceMember = op.ReferenceMember
	o.ReferenceVersionVectors = op.ReferenceVersionVectors
	o.Flags = op.Flags
}
func (o *GetCompressedReportRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCompressedReportRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCompressedReportOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCompressedReportResponse structure represents the GetCompressedReport operation response
type GetCompressedReportResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// memberVersionVectors: If the flags parameter has REPORTING_FLAGS_BACKLOG set, the
	// set of version vectors for the replicated folders on the server MUST be returned
	// in this output parameter. The VersionVectorData structure is specified in section
	// 2.2.1.4.
	MemberVersionVectors *oaut.SafeArray `idl:"name:memberVersionVectors" json:"member_version_vectors"`
	// reportCompressed: The compressed report body in the XML format MUST be returned in
	// this output parameter. The format of the XML MUST be the same as for the reportXML
	// member of the GetReport method. This MUST be an encoded field whose format is specified
	// by the DFS-R compression algorithm (as specified in [MS-FRS2] section 3.1.1.1).
	ReportCompressed *oaut.String `idl:"name:reportCompressed" json:"report_compressed"`
	// uncompressedReportSize: The size, in bytes, of the uncompressed data returned in
	// the reportCompressed parameter.
	UncompressedReportSize int32 `idl:"name:uncompressedReportSize" json:"uncompressed_report_size"`
	// Return: The GetCompressedReport return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCompressedReportResponse) xxx_ToOp(ctx context.Context) *xxx_GetCompressedReportOperation {
	if o == nil {
		return &xxx_GetCompressedReportOperation{}
	}
	return &xxx_GetCompressedReportOperation{
		That:                   o.That,
		MemberVersionVectors:   o.MemberVersionVectors,
		ReportCompressed:       o.ReportCompressed,
		UncompressedReportSize: o.UncompressedReportSize,
		Return:                 o.Return,
	}
}

func (o *GetCompressedReportResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCompressedReportOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MemberVersionVectors = op.MemberVersionVectors
	o.ReportCompressed = op.ReportCompressed
	o.UncompressedReportSize = op.UncompressedReportSize
	o.Return = op.Return
}
func (o *GetCompressedReportResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCompressedReportResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCompressedReportOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRawReportExOperation structure represents the GetRawReportEx operation
type xxx_GetRawReportExOperation struct {
	This                 *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ReplicationGroupGUID *dtyp.GUID      `idl:"name:replicationGroupGuid" json:"replication_group_guid"`
	ReportOptions        *oaut.SafeArray `idl:"name:reportOptions" json:"report_options"`
	Report               *oaut.SafeArray `idl:"name:report" json:"report"`
	Return               int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRawReportExOperation) OpNum() int { return 5 }

func (o *xxx_GetRawReportExOperation) OpName() string {
	return "/IServerHealthReport/v0/GetRawReportEx"
}

func (o *xxx_GetRawReportExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRawReportExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// replicationGroupGuid {in} (1:{alias=GUID}(struct))
	{
		if o.ReplicationGroupGUID != nil {
			if err := o.ReplicationGroupGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// reportOptions {in} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.ReportOptions != nil {
			_ptr_reportOptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReportOptions != nil {
					_ptr_reportOptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.ReportOptions != nil {
							if err := o.ReportOptions.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.ReportOptions, _ptr_reportOptions); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReportOptions, _ptr_reportOptions); err != nil {
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

func (o *xxx_GetRawReportExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// replicationGroupGuid {in} (1:{alias=GUID}(struct))
	{
		if o.ReplicationGroupGUID == nil {
			o.ReplicationGroupGUID = &dtyp.GUID{}
		}
		if err := o.ReplicationGroupGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reportOptions {in} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_reportOptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_reportOptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.ReportOptions == nil {
					o.ReportOptions = &oaut.SafeArray{}
				}
				if err := o.ReportOptions.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_reportOptions := func(ptr interface{}) { o.ReportOptions = *ptr.(**oaut.SafeArray) }
			if err := w.ReadPointer(&o.ReportOptions, _s_reportOptions, _ptr_reportOptions); err != nil {
				return err
			}
			return nil
		})
		_s_reportOptions := func(ptr interface{}) { o.ReportOptions = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.ReportOptions, _s_reportOptions, _ptr_reportOptions); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRawReportExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRawReportExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// report {out} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Report != nil {
			_ptr_report := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Report != nil {
					_ptr_report := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Report != nil {
							if err := o.Report.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Report, _ptr_report); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Report, _ptr_report); err != nil {
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

func (o *xxx_GetRawReportExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// report {out} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_report := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_report := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Report == nil {
					o.Report = &oaut.SafeArray{}
				}
				if err := o.Report.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_report := func(ptr interface{}) { o.Report = *ptr.(**oaut.SafeArray) }
			if err := w.ReadPointer(&o.Report, _s_report, _ptr_report); err != nil {
				return err
			}
			return nil
		})
		_s_report := func(ptr interface{}) { o.Report = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Report, _s_report, _ptr_report); err != nil {
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

// GetRawReportExRequest structure represents the GetRawReportEx operation request
type GetRawReportExRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// replicationGroupGuid: Not implemented.
	ReplicationGroupGUID *dtyp.GUID `idl:"name:replicationGroupGuid" json:"replication_group_guid"`
	// reportOptions: Not implemented.
	ReportOptions *oaut.SafeArray `idl:"name:reportOptions" json:"report_options"`
}

func (o *GetRawReportExRequest) xxx_ToOp(ctx context.Context) *xxx_GetRawReportExOperation {
	if o == nil {
		return &xxx_GetRawReportExOperation{}
	}
	return &xxx_GetRawReportExOperation{
		This:                 o.This,
		ReplicationGroupGUID: o.ReplicationGroupGUID,
		ReportOptions:        o.ReportOptions,
	}
}

func (o *GetRawReportExRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRawReportExOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReplicationGroupGUID = op.ReplicationGroupGUID
	o.ReportOptions = op.ReportOptions
}
func (o *GetRawReportExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetRawReportExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRawReportExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRawReportExResponse structure represents the GetRawReportEx operation response
type GetRawReportExResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// report: Not implemented.
	Report *oaut.SafeArray `idl:"name:report" json:"report"`
	// Return: The GetRawReportEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRawReportExResponse) xxx_ToOp(ctx context.Context) *xxx_GetRawReportExOperation {
	if o == nil {
		return &xxx_GetRawReportExOperation{}
	}
	return &xxx_GetRawReportExOperation{
		That:   o.That,
		Report: o.Report,
		Return: o.Return,
	}
}

func (o *GetRawReportExResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRawReportExOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Report = op.Report
	o.Return = op.Return
}
func (o *GetRawReportExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetRawReportExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRawReportExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetReferenceVersionVectorsOperation structure represents the GetReferenceVersionVectors operation
type xxx_GetReferenceVersionVectorsOperation struct {
	This                    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ReplicationGroupGUID    *dtyp.GUID      `idl:"name:replicationGroupGuid" json:"replication_group_guid"`
	ReferenceVersionVectors *oaut.SafeArray `idl:"name:referenceVersionVectors" json:"reference_version_vectors"`
	Return                  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReferenceVersionVectorsOperation) OpNum() int { return 6 }

func (o *xxx_GetReferenceVersionVectorsOperation) OpName() string {
	return "/IServerHealthReport/v0/GetReferenceVersionVectors"
}

func (o *xxx_GetReferenceVersionVectorsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceVersionVectorsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// replicationGroupGuid {in} (1:{alias=GUID}(struct))
	{
		if o.ReplicationGroupGUID != nil {
			if err := o.ReplicationGroupGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetReferenceVersionVectorsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// replicationGroupGuid {in} (1:{alias=GUID}(struct))
	{
		if o.ReplicationGroupGUID == nil {
			o.ReplicationGroupGUID = &dtyp.GUID{}
		}
		if err := o.ReplicationGroupGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceVersionVectorsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceVersionVectorsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// referenceVersionVectors {out} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.ReferenceVersionVectors != nil {
			_ptr_referenceVersionVectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferenceVersionVectors != nil {
					_ptr_referenceVersionVectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.ReferenceVersionVectors != nil {
							if err := o.ReferenceVersionVectors.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.ReferenceVersionVectors, _ptr_referenceVersionVectors); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferenceVersionVectors, _ptr_referenceVersionVectors); err != nil {
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

func (o *xxx_GetReferenceVersionVectorsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// referenceVersionVectors {out} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_referenceVersionVectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_referenceVersionVectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.ReferenceVersionVectors == nil {
					o.ReferenceVersionVectors = &oaut.SafeArray{}
				}
				if err := o.ReferenceVersionVectors.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_referenceVersionVectors := func(ptr interface{}) { o.ReferenceVersionVectors = *ptr.(**oaut.SafeArray) }
			if err := w.ReadPointer(&o.ReferenceVersionVectors, _s_referenceVersionVectors, _ptr_referenceVersionVectors); err != nil {
				return err
			}
			return nil
		})
		_s_referenceVersionVectors := func(ptr interface{}) { o.ReferenceVersionVectors = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.ReferenceVersionVectors, _s_referenceVersionVectors, _ptr_referenceVersionVectors); err != nil {
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

// GetReferenceVersionVectorsRequest structure represents the GetReferenceVersionVectors operation request
type GetReferenceVersionVectorsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// replicationGroupGuid: MUST be the identifier of the replication group for which the
	// server gets version vectors.
	ReplicationGroupGUID *dtyp.GUID `idl:"name:replicationGroupGuid" json:"replication_group_guid"`
}

func (o *GetReferenceVersionVectorsRequest) xxx_ToOp(ctx context.Context) *xxx_GetReferenceVersionVectorsOperation {
	if o == nil {
		return &xxx_GetReferenceVersionVectorsOperation{}
	}
	return &xxx_GetReferenceVersionVectorsOperation{
		This:                 o.This,
		ReplicationGroupGUID: o.ReplicationGroupGUID,
	}
}

func (o *GetReferenceVersionVectorsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetReferenceVersionVectorsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReplicationGroupGUID = op.ReplicationGroupGUID
}
func (o *GetReferenceVersionVectorsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetReferenceVersionVectorsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReferenceVersionVectorsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReferenceVersionVectorsResponse structure represents the GetReferenceVersionVectors operation response
type GetReferenceVersionVectorsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// referenceVersionVectors: MUST be the array of version vectors for the replicated
	// folders on the server. MUST be returned in this output parameter. For more information,
	// see section 2.2.1.4.
	ReferenceVersionVectors *oaut.SafeArray `idl:"name:referenceVersionVectors" json:"reference_version_vectors"`
	// Return: The GetReferenceVersionVectors return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReferenceVersionVectorsResponse) xxx_ToOp(ctx context.Context) *xxx_GetReferenceVersionVectorsOperation {
	if o == nil {
		return &xxx_GetReferenceVersionVectorsOperation{}
	}
	return &xxx_GetReferenceVersionVectorsOperation{
		That:                    o.That,
		ReferenceVersionVectors: o.ReferenceVersionVectors,
		Return:                  o.Return,
	}
}

func (o *GetReferenceVersionVectorsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetReferenceVersionVectorsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReferenceVersionVectors = op.ReferenceVersionVectors
	o.Return = op.Return
}
func (o *GetReferenceVersionVectorsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetReferenceVersionVectorsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReferenceVersionVectorsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetReferenceBacklogCountsOperation structure represents the GetReferenceBacklogCounts operation
type xxx_GetReferenceBacklogCountsOperation struct {
	This                     *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat  `idl:"name:That" json:"that"`
	FlatMemberVersionVectors *oaut.SafeArray `idl:"name:flatMemberVersionVectors" json:"flat_member_version_vectors"`
	BacklogCounts            *oaut.SafeArray `idl:"name:backlogCounts" json:"backlog_counts"`
	Return                   int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReferenceBacklogCountsOperation) OpNum() int { return 8 }

func (o *xxx_GetReferenceBacklogCountsOperation) OpName() string {
	return "/IServerHealthReport/v0/GetReferenceBacklogCounts"
}

func (o *xxx_GetReferenceBacklogCountsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceBacklogCountsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// flatMemberVersionVectors {in} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.FlatMemberVersionVectors != nil {
			_ptr_flatMemberVersionVectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FlatMemberVersionVectors != nil {
					_ptr_flatMemberVersionVectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.FlatMemberVersionVectors != nil {
							if err := o.FlatMemberVersionVectors.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.FlatMemberVersionVectors, _ptr_flatMemberVersionVectors); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FlatMemberVersionVectors, _ptr_flatMemberVersionVectors); err != nil {
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

func (o *xxx_GetReferenceBacklogCountsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// flatMemberVersionVectors {in} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_flatMemberVersionVectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_flatMemberVersionVectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.FlatMemberVersionVectors == nil {
					o.FlatMemberVersionVectors = &oaut.SafeArray{}
				}
				if err := o.FlatMemberVersionVectors.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_flatMemberVersionVectors := func(ptr interface{}) { o.FlatMemberVersionVectors = *ptr.(**oaut.SafeArray) }
			if err := w.ReadPointer(&o.FlatMemberVersionVectors, _s_flatMemberVersionVectors, _ptr_flatMemberVersionVectors); err != nil {
				return err
			}
			return nil
		})
		_s_flatMemberVersionVectors := func(ptr interface{}) { o.FlatMemberVersionVectors = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.FlatMemberVersionVectors, _s_flatMemberVersionVectors, _ptr_flatMemberVersionVectors); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceBacklogCountsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceBacklogCountsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// backlogCounts {out} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.BacklogCounts != nil {
			_ptr_backlogCounts := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BacklogCounts != nil {
					_ptr_backlogCounts := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.BacklogCounts != nil {
							if err := o.BacklogCounts.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.BacklogCounts, _ptr_backlogCounts); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BacklogCounts, _ptr_backlogCounts); err != nil {
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

func (o *xxx_GetReferenceBacklogCountsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// backlogCounts {out} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_backlogCounts := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_backlogCounts := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.BacklogCounts == nil {
					o.BacklogCounts = &oaut.SafeArray{}
				}
				if err := o.BacklogCounts.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_backlogCounts := func(ptr interface{}) { o.BacklogCounts = *ptr.(**oaut.SafeArray) }
			if err := w.ReadPointer(&o.BacklogCounts, _s_backlogCounts, _ptr_backlogCounts); err != nil {
				return err
			}
			return nil
		})
		_s_backlogCounts := func(ptr interface{}) { o.BacklogCounts = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.BacklogCounts, _s_backlogCounts, _ptr_backlogCounts); err != nil {
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

// GetReferenceBacklogCountsRequest structure represents the GetReferenceBacklogCounts operation request
type GetReferenceBacklogCountsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// flatMemberVersionVectors: MUST be the version vector of the replication group on
	// another member that is participating in the same replication group. If multiple members
	// are specified in the flatMemberVersionVectors array, the backlogCounts array contains
	// the backlog counts for each reference vector specified.
	FlatMemberVersionVectors *oaut.SafeArray `idl:"name:flatMemberVersionVectors" json:"flat_member_version_vectors"`
}

func (o *GetReferenceBacklogCountsRequest) xxx_ToOp(ctx context.Context) *xxx_GetReferenceBacklogCountsOperation {
	if o == nil {
		return &xxx_GetReferenceBacklogCountsOperation{}
	}
	return &xxx_GetReferenceBacklogCountsOperation{
		This:                     o.This,
		FlatMemberVersionVectors: o.FlatMemberVersionVectors,
	}
}

func (o *GetReferenceBacklogCountsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetReferenceBacklogCountsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FlatMemberVersionVectors = op.FlatMemberVersionVectors
}
func (o *GetReferenceBacklogCountsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetReferenceBacklogCountsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReferenceBacklogCountsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReferenceBacklogCountsResponse structure represents the GetReferenceBacklogCounts operation response
type GetReferenceBacklogCountsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// backlogCounts: The relative backlog for each reference vector in flatMemberVersionVectors
	// MUST be returned in this output parameter. The length of the backlogCounts array
	// MUST be the same as the length of flatMemberVersionVectors. The value on each position
	// in the returned array MUST correspond to the version vector on the same position
	// in the flatMemberVersionVectors array.
	BacklogCounts *oaut.SafeArray `idl:"name:backlogCounts" json:"backlog_counts"`
	// Return: The GetReferenceBacklogCounts return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReferenceBacklogCountsResponse) xxx_ToOp(ctx context.Context) *xxx_GetReferenceBacklogCountsOperation {
	if o == nil {
		return &xxx_GetReferenceBacklogCountsOperation{}
	}
	return &xxx_GetReferenceBacklogCountsOperation{
		That:          o.That,
		BacklogCounts: o.BacklogCounts,
		Return:        o.Return,
	}
}

func (o *GetReferenceBacklogCountsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetReferenceBacklogCountsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.BacklogCounts = op.BacklogCounts
	o.Return = op.Return
}
func (o *GetReferenceBacklogCountsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetReferenceBacklogCountsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReferenceBacklogCountsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
