package iupdatesession3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	uamg "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg"
	iupdatesession2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesession2/v0"
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
	_ = iupdatesession2.GoPackage
	_ = uamg.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IUpdateSession3 interface identifier 918efd1e-b5d8-4c90-8540-aeb9bdc56f9d
	UpdateSession3IID = &dcom.IID{Data1: 0x918efd1e, Data2: 0xb5d8, Data3: 0x4c90, Data4: []byte{0x85, 0x40, 0xae, 0xb9, 0xbd, 0xc5, 0x6f, 0x9d}}
	// Syntax UUID
	UpdateSession3SyntaxUUID = &uuid.UUID{TimeLow: 0x918efd1e, TimeMid: 0xb5d8, TimeHiAndVersion: 0x4c90, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x40, Node: [6]uint8{0xae, 0xb9, 0xbd, 0xc5, 0x6f, 0x9d}}
	// Syntax ID
	UpdateSession3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: UpdateSession3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUpdateSession3 interface.
type UpdateSession3Client interface {

	// IUpdateSession2 retrieval method.
	UpdateSession2() iupdatesession2.UpdateSession2Client

	// The IUpdateSession3::CreateUpdateServiceManager (Opnum 16) method retrieves an instance
	// of the IUpdateServiceManager2 interface.
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
	// This method SHOULD return a newly created object that implements IUpdateServiceManager2.
	CreateUpdateServiceManager(context.Context, *CreateUpdateServiceManagerRequest, ...dcerpc.CallOption) (*CreateUpdateServiceManagerResponse, error)

	// The IUpdateSearcher::QueryHistory (opnum 19) method retrieves a collection of history
	// events.
	//
	// The IUpdateSession3::QueryHistory (Opnum 17) method retrieves relevant update history
	// entries.
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
	QueryHistory(context.Context, *QueryHistoryRequest, ...dcerpc.CallOption) (*QueryHistoryResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) UpdateSession3Client
}

type xxx_DefaultUpdateSession3Client struct {
	iupdatesession2.UpdateSession2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUpdateSession3Client) UpdateSession2() iupdatesession2.UpdateSession2Client {
	return o.UpdateSession2Client
}

func (o *xxx_DefaultUpdateSession3Client) CreateUpdateServiceManager(ctx context.Context, in *CreateUpdateServiceManagerRequest, opts ...dcerpc.CallOption) (*CreateUpdateServiceManagerResponse, error) {
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
	out := &CreateUpdateServiceManagerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSession3Client) QueryHistory(ctx context.Context, in *QueryHistoryRequest, opts ...dcerpc.CallOption) (*QueryHistoryResponse, error) {
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
	out := &QueryHistoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSession3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUpdateSession3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUpdateSession3Client) IPID(ctx context.Context, ipid *dcom.IPID) UpdateSession3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUpdateSession3Client{
		UpdateSession2Client: o.UpdateSession2Client.IPID(ctx, ipid),
		cc:                   o.cc,
		ipid:                 ipid,
	}
}

func NewUpdateSession3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (UpdateSession3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(UpdateSession3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iupdatesession2.NewUpdateSession2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultUpdateSession3Client{
		UpdateSession2Client: base,
		cc:                   cc,
		ipid:                 ipid,
	}, nil
}

// xxx_CreateUpdateServiceManagerOperation structure represents the CreateUpdateServiceManager operation
type xxx_CreateUpdateServiceManagerOperation struct {
	This        *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat              `idl:"name:That" json:"that"`
	ReturnValue *uamg.UpdateServiceManager2 `idl:"name:retval" json:"return_value"`
	Return      int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateUpdateServiceManagerOperation) OpNum() int { return 17 }

func (o *xxx_CreateUpdateServiceManagerOperation) OpName() string {
	return "/IUpdateSession3/v0/CreateUpdateServiceManager"
}

func (o *xxx_CreateUpdateServiceManagerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateUpdateServiceManagerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateUpdateServiceManagerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CreateUpdateServiceManagerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateUpdateServiceManagerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateServiceManager2}(interface))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&uamg.UpdateServiceManager2{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CreateUpdateServiceManagerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateServiceManager2}(interface))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &uamg.UpdateServiceManager2{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**uamg.UpdateServiceManager2) }
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

// CreateUpdateServiceManagerRequest structure represents the CreateUpdateServiceManager operation request
type CreateUpdateServiceManagerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CreateUpdateServiceManagerRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateUpdateServiceManagerOperation) *xxx_CreateUpdateServiceManagerOperation {
	if op == nil {
		op = &xxx_CreateUpdateServiceManagerOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CreateUpdateServiceManagerRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateUpdateServiceManagerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CreateUpdateServiceManagerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateUpdateServiceManagerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateUpdateServiceManagerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateUpdateServiceManagerResponse structure represents the CreateUpdateServiceManager operation response
type CreateUpdateServiceManagerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: An IUpdateServiceManager2 instance.
	ReturnValue *uamg.UpdateServiceManager2 `idl:"name:retval" json:"return_value"`
	// Return: The CreateUpdateServiceManager return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateUpdateServiceManagerResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateUpdateServiceManagerOperation) *xxx_CreateUpdateServiceManagerOperation {
	if op == nil {
		op = &xxx_CreateUpdateServiceManagerOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *CreateUpdateServiceManagerResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateUpdateServiceManagerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *CreateUpdateServiceManagerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateUpdateServiceManagerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateUpdateServiceManagerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryHistoryOperation structure represents the QueryHistory operation
type xxx_QueryHistoryOperation struct {
	This        *dcom.ORPCThis                     `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat                     `idl:"name:That" json:"that"`
	Criteria    *oaut.String                       `idl:"name:criteria" json:"criteria"`
	StartIndex  int32                              `idl:"name:startIndex" json:"start_index"`
	Count       int32                              `idl:"name:count" json:"count"`
	ReturnValue *uamg.UpdateHistoryEntryCollection `idl:"name:retval" json:"return_value"`
	Return      int32                              `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryHistoryOperation) OpNum() int { return 18 }

func (o *xxx_QueryHistoryOperation) OpName() string { return "/IUpdateSession3/v0/QueryHistory" }

func (o *xxx_QueryHistoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryHistoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// criteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Criteria != nil {
			_ptr_criteria := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Criteria != nil {
					if err := o.Criteria.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Criteria, _ptr_criteria); err != nil {
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
	// startIndex {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.StartIndex); err != nil {
			return err
		}
	}
	// count {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryHistoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// criteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_criteria := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Criteria == nil {
				o.Criteria = &oaut.String{}
			}
			if err := o.Criteria.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_criteria := func(ptr interface{}) { o.Criteria = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Criteria, _s_criteria, _ptr_criteria); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// startIndex {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.StartIndex); err != nil {
			return err
		}
	}
	// count {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryHistoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryHistoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateHistoryEntryCollection}(interface))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&uamg.UpdateHistoryEntryCollection{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_QueryHistoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateHistoryEntryCollection}(interface))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &uamg.UpdateHistoryEntryCollection{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**uamg.UpdateHistoryEntryCollection) }
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

// QueryHistoryRequest structure represents the QueryHistory operation request
type QueryHistoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// criteria: A string specifying criteria that select history entries. The string MUST
	// be one of the following:
	//
	// * NULL or empty, in which case all history entries are selected.
	//
	// * "UpdateID='<id>'", where <id> is an update identifier. In this case, only history
	// entries pertaining to this update are selected.
	Criteria *oaut.String `idl:"name:criteria" json:"criteria"`
	// startIndex: The zero-based index of the first history entry to retrieve.
	StartIndex int32 `idl:"name:startIndex" json:"start_index"`
	// count: The number of entries to retrieve.
	Count int32 `idl:"name:count" json:"count"`
}

func (o *QueryHistoryRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryHistoryOperation) *xxx_QueryHistoryOperation {
	if op == nil {
		op = &xxx_QueryHistoryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Criteria = o.Criteria
	op.StartIndex = o.StartIndex
	op.Count = o.Count
	return op
}

func (o *QueryHistoryRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryHistoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Criteria = op.Criteria
	o.StartIndex = op.StartIndex
	o.Count = op.Count
}
func (o *QueryHistoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryHistoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryHistoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryHistoryResponse structure represents the QueryHistory operation response
type QueryHistoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: An IUpdateHistoryEntryCollection containing the history entries requested.
	// If fewer entries are available than requested in the count parameter, only the entries
	// available are retrieved.
	ReturnValue *uamg.UpdateHistoryEntryCollection `idl:"name:retval" json:"return_value"`
	// Return: The QueryHistory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryHistoryResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryHistoryOperation) *xxx_QueryHistoryOperation {
	if op == nil {
		op = &xxx_QueryHistoryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *QueryHistoryResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryHistoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *QueryHistoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryHistoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryHistoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
