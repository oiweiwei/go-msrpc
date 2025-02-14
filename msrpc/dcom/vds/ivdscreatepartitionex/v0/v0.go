package ivdscreatepartitionex

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
	vds "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds"
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
	_ = vds.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsCreatePartitionEx interface identifier 9882f547-cfc3-420b-9750-00dfbec50662
	CreatePartitionExIID = &dcom.IID{Data1: 0x9882f547, Data2: 0xcfc3, Data3: 0x420b, Data4: []byte{0x97, 0x50, 0x00, 0xdf, 0xbe, 0xc5, 0x06, 0x62}}
	// Syntax UUID
	CreatePartitionExSyntaxUUID = &uuid.UUID{TimeLow: 0x9882f547, TimeMid: 0xcfc3, TimeHiAndVersion: 0x420b, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0x50, Node: [6]uint8{0x0, 0xdf, 0xbe, 0xc5, 0x6, 0x62}}
	// Syntax ID
	CreatePartitionExSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CreatePartitionExSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsCreatePartitionEx interface.
type CreatePartitionExClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The CreatePartitionEx method creates a partition on a disk at a specified byte offset,
	// with an optional alignment parameter.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// Note  Creating or deleting partitions on dynamic disks is not supported.
	//
	// ERROR_SUCCESS (0x00000000)
	CreatePartitionEx(context.Context, *CreatePartitionExRequest, ...dcerpc.CallOption) (*CreatePartitionExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CreatePartitionExClient
}

type xxx_DefaultCreatePartitionExClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCreatePartitionExClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCreatePartitionExClient) CreatePartitionEx(ctx context.Context, in *CreatePartitionExRequest, opts ...dcerpc.CallOption) (*CreatePartitionExResponse, error) {
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
	out := &CreatePartitionExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCreatePartitionExClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCreatePartitionExClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCreatePartitionExClient) IPID(ctx context.Context, ipid *dcom.IPID) CreatePartitionExClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCreatePartitionExClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewCreatePartitionExClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CreatePartitionExClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CreatePartitionExSyntaxV0_0))...)
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
	return &xxx_DefaultCreatePartitionExClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CreatePartitionExOperation structure represents the CreatePartitionEx operation
type xxx_CreatePartitionExOperation struct {
	This       *dcom.ORPCThis                 `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat                 `idl:"name:That" json:"that"`
	Offset     uint64                         `idl:"name:ullOffset" json:"offset"`
	Size       uint64                         `idl:"name:ullSize" json:"size"`
	Align      uint32                         `idl:"name:ulAlign" json:"align"`
	Parameters *vds.CreatePartitionParameters `idl:"name:para" json:"parameters"`
	Async      *vds.Async                     `idl:"name:ppAsync" json:"async"`
	Return     int32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreatePartitionExOperation) OpNum() int { return 3 }

func (o *xxx_CreatePartitionExOperation) OpName() string {
	return "/IVdsCreatePartitionEx/v0/CreatePartitionEx"
}

func (o *xxx_CreatePartitionExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ullOffset {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Offset); err != nil {
			return err
		}
	}
	// ullSize {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	// ulAlign {in} (1:(uint32))
	{
		if err := w.WriteData(o.Align); err != nil {
			return err
		}
	}
	// para {in} (1:{pointer=ref}*(1))(2:{alias=CREATE_PARTITION_PARAMETERS}(struct))
	{
		if o.Parameters != nil {
			if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.CreatePartitionParameters{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreatePartitionExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ullOffset {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Offset); err != nil {
			return err
		}
	}
	// ullSize {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	// ulAlign {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Align); err != nil {
			return err
		}
	}
	// para {in} (1:{pointer=ref}*(1))(2:{alias=CREATE_PARTITION_PARAMETERS}(struct))
	{
		if o.Parameters == nil {
			o.Parameters = &vds.CreatePartitionParameters{}
		}
		if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppAsync {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		if o.Async != nil {
			_ptr_ppAsync := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Async != nil {
					if err := o.Async.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Async{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Async, _ptr_ppAsync); err != nil {
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

func (o *xxx_CreatePartitionExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppAsync {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		_ptr_ppAsync := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Async == nil {
				o.Async = &vds.Async{}
			}
			if err := o.Async.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppAsync := func(ptr interface{}) { o.Async = *ptr.(**vds.Async) }
		if err := w.ReadPointer(&o.Async, _s_ppAsync, _ptr_ppAsync); err != nil {
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

// CreatePartitionExRequest structure represents the CreatePartitionEx operation request
type CreatePartitionExRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset from the beginning of the disk where the new partition
	// will be created. If ulAlign is not zero, the offset MUST fall within the first cylinder
	// for an MBR disk (GPT disks do not have this restriction).
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// ullSize: The size of the new partition, in bytes.<105>
	Size uint64 `idl:"name:ullSize" json:"size"`
	// ulAlign: The number of bytes for volume alignment. The offset specified in ullOffset
	// will be rounded up or down to an alignment boundary. If zero is specified, the server
	// will base the alignment value on the size of the disk on which the volume is created.<106>
	Align uint32 `idl:"name:ulAlign" json:"align"`
	// para: A pointer to a CREATE_PARTITION_PARAMETERS structure that describes the new
	// partition to create.
	Parameters *vds.CreatePartitionParameters `idl:"name:para" json:"parameters"`
}

func (o *CreatePartitionExRequest) xxx_ToOp(ctx context.Context, op *xxx_CreatePartitionExOperation) *xxx_CreatePartitionExOperation {
	if op == nil {
		op = &xxx_CreatePartitionExOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Offset = op.Offset
	o.Size = op.Size
	o.Align = op.Align
	o.Parameters = op.Parameters
	return op
}

func (o *CreatePartitionExRequest) xxx_FromOp(ctx context.Context, op *xxx_CreatePartitionExOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
	o.Size = op.Size
	o.Align = op.Align
	o.Parameters = op.Parameters
}
func (o *CreatePartitionExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreatePartitionExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePartitionExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreatePartitionExResponse structure represents the CreatePartitionEx operation response
type CreatePartitionExResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it. If the IVdsAsync::Wait
	// method is called on the interface, the interfaces returned in the VDS_ASYNC_OUTPUT
	// structure MUST be released as well. For information on asynchronous tasks, see section
	// 3.4.5.1.9.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The CreatePartitionEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreatePartitionExResponse) xxx_ToOp(ctx context.Context, op *xxx_CreatePartitionExOperation) *xxx_CreatePartitionExOperation {
	if op == nil {
		op = &xxx_CreatePartitionExOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
	return op
}

func (o *CreatePartitionExResponse) xxx_FromOp(ctx context.Context, op *xxx_CreatePartitionExOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *CreatePartitionExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreatePartitionExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePartitionExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
