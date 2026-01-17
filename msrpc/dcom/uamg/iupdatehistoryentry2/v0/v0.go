package iupdatehistoryentry2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	uamg "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg"
	iupdatehistoryentry "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatehistoryentry/v0"
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
	_ = iupdatehistoryentry.GoPackage
	_ = uamg.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IUpdateHistoryEntry2 interface identifier c2bfb780-4539-4132-ab8c-0a8772013ab6
	UpdateHistoryEntry2IID = &dcom.IID{Data1: 0xc2bfb780, Data2: 0x4539, Data3: 0x4132, Data4: []byte{0xab, 0x8c, 0x0a, 0x87, 0x72, 0x01, 0x3a, 0xb6}}
	// Syntax UUID
	UpdateHistoryEntry2SyntaxUUID = &uuid.UUID{TimeLow: 0xc2bfb780, TimeMid: 0x4539, TimeHiAndVersion: 0x4132, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x8c, Node: [6]uint8{0xa, 0x87, 0x72, 0x1, 0x3a, 0xb6}}
	// Syntax ID
	UpdateHistoryEntry2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: UpdateHistoryEntry2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUpdateHistoryEntry2 interface.
type UpdateHistoryEntry2Client interface {

	// IUpdateHistoryEntry retrieval method.
	UpdateHistoryEntry() iupdatehistoryentry.UpdateHistoryEntryClient

	// The IUpdate::Categories (opnum 12) method retrieves a collection of the categories
	// to which the update belongs.
	//
	// The IUpdateHistoryEntry2::Categories (opnum 22) method retrieves a collection of
	// the update categories to which an update belongs.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The server SHOULD return the value of the Categories ADM element.
	GetCategories(context.Context, *GetCategoriesRequest, ...dcerpc.CallOption) (*GetCategoriesResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) UpdateHistoryEntry2Client
}

type xxx_DefaultUpdateHistoryEntry2Client struct {
	iupdatehistoryentry.UpdateHistoryEntryClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUpdateHistoryEntry2Client) UpdateHistoryEntry() iupdatehistoryentry.UpdateHistoryEntryClient {
	return o.UpdateHistoryEntryClient
}

func (o *xxx_DefaultUpdateHistoryEntry2Client) GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...dcerpc.CallOption) (*GetCategoriesResponse, error) {
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
	out := &GetCategoriesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateHistoryEntry2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUpdateHistoryEntry2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUpdateHistoryEntry2Client) IPID(ctx context.Context, ipid *dcom.IPID) UpdateHistoryEntry2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUpdateHistoryEntry2Client{
		UpdateHistoryEntryClient: o.UpdateHistoryEntryClient.IPID(ctx, ipid),
		cc:                       o.cc,
		ipid:                     ipid,
	}
}

func NewUpdateHistoryEntry2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (UpdateHistoryEntry2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(UpdateHistoryEntry2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iupdatehistoryentry.NewUpdateHistoryEntryClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultUpdateHistoryEntry2Client{
		UpdateHistoryEntryClient: base,
		cc:                       cc,
		ipid:                     ipid,
	}, nil
}

// xxx_GetCategoriesOperation structure represents the Categories operation
type xxx_GetCategoriesOperation struct {
	This        *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat           `idl:"name:That" json:"that"`
	ReturnValue *uamg.CategoryCollection `idl:"name:retval" json:"return_value"`
	Return      int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCategoriesOperation) OpNum() int { return 21 }

func (o *xxx_GetCategoriesOperation) OpName() string { return "/IUpdateHistoryEntry2/v0/Categories" }

func (o *xxx_GetCategoriesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCategoriesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCategoriesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCategoriesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCategoriesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=ICategoryCollection}(interface))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&uamg.CategoryCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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

func (o *xxx_GetCategoriesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=ICategoryCollection}(interface))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &uamg.CategoryCollection{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**uamg.CategoryCollection) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
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

// GetCategoriesRequest structure represents the Categories operation request
type GetCategoriesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCategoriesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCategoriesOperation) *xxx_GetCategoriesOperation {
	if op == nil {
		op = &xxx_GetCategoriesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCategoriesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCategoriesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCategoriesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCategoriesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCategoriesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCategoriesResponse structure represents the Categories operation response
type GetCategoriesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: An ICategoryCollection instance containing the categories to which this update
	// belongs.
	//
	// retval: A collection of the update categories to which an update belongs.
	ReturnValue *uamg.CategoryCollection `idl:"name:retval" json:"return_value"`
	// Return: The Categories return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCategoriesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCategoriesOperation) *xxx_GetCategoriesOperation {
	if op == nil {
		op = &xxx_GetCategoriesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetCategoriesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCategoriesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetCategoriesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCategoriesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCategoriesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
