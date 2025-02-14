package ifsrmcommittablecollection

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
	ifsrmmutablecollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmmutablecollection/v0"
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
	_ = ifsrmmutablecollection.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmCommittableCollection interface identifier 96deb3b5-8b91-4a2a-9d93-80a35d8aa847
	CommittableCollectionIID = &dcom.IID{Data1: 0x96deb3b5, Data2: 0x8b91, Data3: 0x4a2a, Data4: []byte{0x9d, 0x93, 0x80, 0xa3, 0x5d, 0x8a, 0xa8, 0x47}}
	// Syntax UUID
	CommittableCollectionSyntaxUUID = &uuid.UUID{TimeLow: 0x96deb3b5, TimeMid: 0x8b91, TimeHiAndVersion: 0x4a2a, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x93, Node: [6]uint8{0x80, 0xa3, 0x5d, 0x8a, 0xa8, 0x47}}
	// Syntax ID
	CommittableCollectionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CommittableCollectionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmCommittableCollection interface.
type CommittableCollectionClient interface {

	// IFsrmMutableCollection retrieval method.
	MutableCollection() ifsrmmutablecollection.MutableCollectionClient

	// Commit operation.
	Commit(context.Context, *CommitRequest, ...dcerpc.CallOption) (*CommitResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CommittableCollectionClient
}

type xxx_DefaultCommittableCollectionClient struct {
	ifsrmmutablecollection.MutableCollectionClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCommittableCollectionClient) MutableCollection() ifsrmmutablecollection.MutableCollectionClient {
	return o.MutableCollectionClient
}

func (o *xxx_DefaultCommittableCollectionClient) Commit(ctx context.Context, in *CommitRequest, opts ...dcerpc.CallOption) (*CommitResponse, error) {
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
	out := &CommitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCommittableCollectionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCommittableCollectionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCommittableCollectionClient) IPID(ctx context.Context, ipid *dcom.IPID) CommittableCollectionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCommittableCollectionClient{
		MutableCollectionClient: o.MutableCollectionClient.IPID(ctx, ipid),
		cc:                      o.cc,
		ipid:                    ipid,
	}
}

func NewCommittableCollectionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CommittableCollectionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CommittableCollectionSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmmutablecollection.NewMutableCollectionClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultCommittableCollectionClient{
		MutableCollectionClient: base,
		cc:                      cc,
		ipid:                    ipid,
	}, nil
}

// xxx_CommitOperation structure represents the Commit operation
type xxx_CommitOperation struct {
	This    *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Options fsrm.CommitOptions `idl:"name:options" json:"options"`
	Results *fsrm.Collection   `idl:"name:results" json:"results"`
	Return  int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_CommitOperation) OpNum() int { return 18 }

func (o *xxx_CommitOperation) OpName() string { return "/IFsrmCommittableCollection/v0/Commit" }

func (o *xxx_CommitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// options {in} (1:{alias=FsrmCommitOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// options {in} (1:{alias=FsrmCommitOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// results {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.Results != nil {
			_ptr_results := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Results != nil {
					if err := o.Results.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Results, _ptr_results); err != nil {
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

func (o *xxx_CommitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// results {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_results := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Results == nil {
				o.Results = &fsrm.Collection{}
			}
			if err := o.Results.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_results := func(ptr interface{}) { o.Results = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.Results, _s_results, _ptr_results); err != nil {
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

// CommitRequest structure represents the Commit operation request
type CommitRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis     `idl:"name:This" json:"this"`
	Options fsrm.CommitOptions `idl:"name:options" json:"options"`
}

func (o *CommitRequest) xxx_ToOp(ctx context.Context, op *xxx_CommitOperation) *xxx_CommitOperation {
	if op == nil {
		op = &xxx_CommitOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Options = op.Options
	return op
}

func (o *CommitRequest) xxx_FromOp(ctx context.Context, op *xxx_CommitOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Options = op.Options
}
func (o *CommitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CommitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CommitResponse structure represents the Commit operation response
type CommitResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Results *fsrm.Collection `idl:"name:results" json:"results"`
	// Return: The Commit return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CommitResponse) xxx_ToOp(ctx context.Context, op *xxx_CommitOperation) *xxx_CommitOperation {
	if op == nil {
		op = &xxx_CommitOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Results = op.Results
	o.Return = op.Return
	return op
}

func (o *CommitResponse) xxx_FromOp(ctx context.Context, op *xxx_CommitOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Results = op.Results
	o.Return = op.Return
}
func (o *CommitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CommitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
