package iserverhealthreport2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iserverhealthreport "github.com/oiweiwei/go-msrpc/msrpc/dcom/dfsrh/iserverhealthreport/v0"
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
	_ = iserverhealthreport.GoPackage
	_ = dtyp.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/dfsrh"
)

var (
	// IServerHealthReport2 interface identifier 20d15747-6c48-4254-a358-65039fd8c63c
	ServerHealthReport2IID = &dcom.IID{Data1: 0x20d15747, Data2: 0x6c48, Data3: 0x4254, Data4: []byte{0xa3, 0x58, 0x65, 0x03, 0x9f, 0xd8, 0xc6, 0x3c}}
	// Syntax UUID
	ServerHealthReport2SyntaxUUID = &uuid.UUID{TimeLow: 0x20d15747, TimeMid: 0x6c48, TimeHiAndVersion: 0x4254, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x58, Node: [6]uint8{0x65, 0x3, 0x9f, 0xd8, 0xc6, 0x3c}}
	// Syntax ID
	ServerHealthReport2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ServerHealthReport2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IServerHealthReport2 interface.
type ServerHealthReport2Client interface {

	// IServerHealthReport retrieval method.
	ServerHealthReport() iserverhealthreport.ServerHealthReportClient

	// GetReport2 operation.
	GetReport2(context.Context, *GetReport2Request, ...dcerpc.CallOption) (*GetReport2Response, error)

	// GetCompressedReport2 operation.
	GetCompressedReport2(context.Context, *GetCompressedReport2Request, ...dcerpc.CallOption) (*GetCompressedReport2Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ServerHealthReport2Client
}

type xxx_DefaultServerHealthReport2Client struct {
	iserverhealthreport.ServerHealthReportClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultServerHealthReport2Client) ServerHealthReport() iserverhealthreport.ServerHealthReportClient {
	return o.ServerHealthReportClient
}

func (o *xxx_DefaultServerHealthReport2Client) GetReport2(ctx context.Context, in *GetReport2Request, opts ...dcerpc.CallOption) (*GetReport2Response, error) {
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
	out := &GetReport2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServerHealthReport2Client) GetCompressedReport2(ctx context.Context, in *GetCompressedReport2Request, opts ...dcerpc.CallOption) (*GetCompressedReport2Response, error) {
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
	out := &GetCompressedReport2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServerHealthReport2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultServerHealthReport2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultServerHealthReport2Client) IPID(ctx context.Context, ipid *dcom.IPID) ServerHealthReport2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultServerHealthReport2Client{
		ServerHealthReportClient: o.ServerHealthReportClient.IPID(ctx, ipid),
		cc:                       o.cc,
		ipid:                     ipid,
	}
}

func NewServerHealthReport2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ServerHealthReport2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ServerHealthReport2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iserverhealthreport.NewServerHealthReportClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultServerHealthReport2Client{
		ServerHealthReportClient: base,
		cc:                       cc,
		ipid:                     ipid,
	}, nil
}

// xxx_GetReport2Operation structure represents the GetReport2 operation
type xxx_GetReport2Operation struct {
	This                    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ReplicationGroupGUID    *dtyp.GUID      `idl:"name:replicationGroupGuid" json:"replication_group_guid"`
	ReferenceMember         *oaut.String    `idl:"name:referenceMember" json:"reference_member"`
	ServerName              *oaut.String    `idl:"name:serverName" json:"server_name"`
	ReferenceVersionVectors *oaut.SafeArray `idl:"name:referenceVersionVectors" json:"reference_version_vectors"`
	Flags                   int32           `idl:"name:flags" json:"flags"`
	MemberVersionVectors    *oaut.SafeArray `idl:"name:memberVersionVectors" json:"member_version_vectors"`
	ReportXML               *oaut.String    `idl:"name:reportXML" json:"report_xml"`
	Return                  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReport2Operation) OpNum() int { return 9 }

func (o *xxx_GetReport2Operation) OpName() string { return "/IServerHealthReport2/v0/GetReport2" }

func (o *xxx_GetReport2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReport2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// serverName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ServerName != nil {
			_ptr_serverName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServerName != nil {
					if err := o.ServerName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_serverName); err != nil {
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
	// flags {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReport2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// serverName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_serverName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServerName == nil {
				o.ServerName = &oaut.String{}
			}
			if err := o.ServerName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_serverName := func(ptr interface{}) { o.ServerName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ServerName, _s_serverName, _ptr_serverName); err != nil {
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
	// flags {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReport2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReport2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetReport2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetReport2Request structure represents the GetReport2 operation request
type GetReport2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	ReplicationGroupGUID    *dtyp.GUID      `idl:"name:replicationGroupGuid" json:"replication_group_guid"`
	ReferenceMember         *oaut.String    `idl:"name:referenceMember" json:"reference_member"`
	ServerName              *oaut.String    `idl:"name:serverName" json:"server_name"`
	ReferenceVersionVectors *oaut.SafeArray `idl:"name:referenceVersionVectors" json:"reference_version_vectors"`
	Flags                   int32           `idl:"name:flags" json:"flags"`
}

func (o *GetReport2Request) xxx_ToOp(ctx context.Context, op *xxx_GetReport2Operation) *xxx_GetReport2Operation {
	if op == nil {
		op = &xxx_GetReport2Operation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.ReplicationGroupGUID = op.ReplicationGroupGUID
	o.ReferenceMember = op.ReferenceMember
	o.ServerName = op.ServerName
	o.ReferenceVersionVectors = op.ReferenceVersionVectors
	o.Flags = op.Flags
	return op
}

func (o *GetReport2Request) xxx_FromOp(ctx context.Context, op *xxx_GetReport2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReplicationGroupGUID = op.ReplicationGroupGUID
	o.ReferenceMember = op.ReferenceMember
	o.ServerName = op.ServerName
	o.ReferenceVersionVectors = op.ReferenceVersionVectors
	o.Flags = op.Flags
}
func (o *GetReport2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetReport2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReport2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReport2Response structure represents the GetReport2 operation response
type GetReport2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                 *dcom.ORPCThat  `idl:"name:That" json:"that"`
	MemberVersionVectors *oaut.SafeArray `idl:"name:memberVersionVectors" json:"member_version_vectors"`
	ReportXML            *oaut.String    `idl:"name:reportXML" json:"report_xml"`
	// Return: The GetReport2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReport2Response) xxx_ToOp(ctx context.Context, op *xxx_GetReport2Operation) *xxx_GetReport2Operation {
	if op == nil {
		op = &xxx_GetReport2Operation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.MemberVersionVectors = op.MemberVersionVectors
	o.ReportXML = op.ReportXML
	o.Return = op.Return
	return op
}

func (o *GetReport2Response) xxx_FromOp(ctx context.Context, op *xxx_GetReport2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MemberVersionVectors = op.MemberVersionVectors
	o.ReportXML = op.ReportXML
	o.Return = op.Return
}
func (o *GetReport2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetReport2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReport2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCompressedReport2Operation structure represents the GetCompressedReport2 operation
type xxx_GetCompressedReport2Operation struct {
	This                    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ReplicationGroupGUID    *dtyp.GUID      `idl:"name:replicationGroupGuid" json:"replication_group_guid"`
	ReferenceMember         *oaut.String    `idl:"name:referenceMember" json:"reference_member"`
	ServerName              *oaut.String    `idl:"name:serverName" json:"server_name"`
	ReferenceVersionVectors *oaut.SafeArray `idl:"name:referenceVersionVectors" json:"reference_version_vectors"`
	Flags                   int32           `idl:"name:flags" json:"flags"`
	MemberVersionVectors    *oaut.SafeArray `idl:"name:memberVersionVectors" json:"member_version_vectors"`
	ReportCompressed        *oaut.String    `idl:"name:reportCompressed" json:"report_compressed"`
	UncompressedReportSize  int32           `idl:"name:uncompressedReportSize" json:"uncompressed_report_size"`
	Return                  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCompressedReport2Operation) OpNum() int { return 10 }

func (o *xxx_GetCompressedReport2Operation) OpName() string {
	return "/IServerHealthReport2/v0/GetCompressedReport2"
}

func (o *xxx_GetCompressedReport2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCompressedReport2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// serverName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ServerName != nil {
			_ptr_serverName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServerName != nil {
					if err := o.ServerName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_serverName); err != nil {
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
	// flags {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCompressedReport2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// serverName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_serverName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServerName == nil {
				o.ServerName = &oaut.String{}
			}
			if err := o.ServerName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_serverName := func(ptr interface{}) { o.ServerName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ServerName, _s_serverName, _ptr_serverName); err != nil {
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
	// flags {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCompressedReport2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCompressedReport2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// uncompressedReportSize {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
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

func (o *xxx_GetCompressedReport2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// uncompressedReportSize {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
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

// GetCompressedReport2Request structure represents the GetCompressedReport2 operation request
type GetCompressedReport2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	ReplicationGroupGUID    *dtyp.GUID      `idl:"name:replicationGroupGuid" json:"replication_group_guid"`
	ReferenceMember         *oaut.String    `idl:"name:referenceMember" json:"reference_member"`
	ServerName              *oaut.String    `idl:"name:serverName" json:"server_name"`
	ReferenceVersionVectors *oaut.SafeArray `idl:"name:referenceVersionVectors" json:"reference_version_vectors"`
	Flags                   int32           `idl:"name:flags" json:"flags"`
}

func (o *GetCompressedReport2Request) xxx_ToOp(ctx context.Context, op *xxx_GetCompressedReport2Operation) *xxx_GetCompressedReport2Operation {
	if op == nil {
		op = &xxx_GetCompressedReport2Operation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.ReplicationGroupGUID = op.ReplicationGroupGUID
	o.ReferenceMember = op.ReferenceMember
	o.ServerName = op.ServerName
	o.ReferenceVersionVectors = op.ReferenceVersionVectors
	o.Flags = op.Flags
	return op
}

func (o *GetCompressedReport2Request) xxx_FromOp(ctx context.Context, op *xxx_GetCompressedReport2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReplicationGroupGUID = op.ReplicationGroupGUID
	o.ReferenceMember = op.ReferenceMember
	o.ServerName = op.ServerName
	o.ReferenceVersionVectors = op.ReferenceVersionVectors
	o.Flags = op.Flags
}
func (o *GetCompressedReport2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCompressedReport2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCompressedReport2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCompressedReport2Response structure represents the GetCompressedReport2 operation response
type GetCompressedReport2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	MemberVersionVectors   *oaut.SafeArray `idl:"name:memberVersionVectors" json:"member_version_vectors"`
	ReportCompressed       *oaut.String    `idl:"name:reportCompressed" json:"report_compressed"`
	UncompressedReportSize int32           `idl:"name:uncompressedReportSize" json:"uncompressed_report_size"`
	// Return: The GetCompressedReport2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCompressedReport2Response) xxx_ToOp(ctx context.Context, op *xxx_GetCompressedReport2Operation) *xxx_GetCompressedReport2Operation {
	if op == nil {
		op = &xxx_GetCompressedReport2Operation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.MemberVersionVectors = op.MemberVersionVectors
	o.ReportCompressed = op.ReportCompressed
	o.UncompressedReportSize = op.UncompressedReportSize
	o.Return = op.Return
	return op
}

func (o *GetCompressedReport2Response) xxx_FromOp(ctx context.Context, op *xxx_GetCompressedReport2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MemberVersionVectors = op.MemberVersionVectors
	o.ReportCompressed = op.ReportCompressed
	o.UncompressedReportSize = op.UncompressedReportSize
	o.Return = op.Return
}
func (o *GetCompressedReport2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCompressedReport2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCompressedReport2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
