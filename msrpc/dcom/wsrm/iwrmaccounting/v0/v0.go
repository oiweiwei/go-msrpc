package iwrmaccounting

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

var (
	// IWRMAccounting interface identifier 4f7ca01c-a9e5-45b6-b142-2332a1339c1d
	AccountingIID = &dcom.IID{Data1: 0x4f7ca01c, Data2: 0xa9e5, Data3: 0x45b6, Data4: []byte{0xb1, 0x42, 0x23, 0x32, 0xa1, 0x33, 0x9c, 0x1d}}
	// Syntax UUID
	AccountingSyntaxUUID = &uuid.UUID{TimeLow: 0x4f7ca01c, TimeMid: 0xa9e5, TimeHiAndVersion: 0x45b6, ClockSeqHiAndReserved: 0xb1, ClockSeqLow: 0x42, Node: [6]uint8{0x23, 0x32, 0xa1, 0x33, 0x9c, 0x1d}}
	// Syntax ID
	AccountingSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AccountingSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWRMAccounting interface.
type AccountingClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	CreateAccountingDB(context.Context, *CreateAccountingDBRequest, ...dcerpc.CallOption) (*CreateAccountingDBResponse, error)

	GetAccountingMetadata(context.Context, *GetAccountingMetadataRequest, ...dcerpc.CallOption) (*GetAccountingMetadataResponse, error)

	ExecuteAccountingQuery(context.Context, *ExecuteAccountingQueryRequest, ...dcerpc.CallOption) (*ExecuteAccountingQueryResponse, error)

	GetRawAccountingData(context.Context, *GetRawAccountingDataRequest, ...dcerpc.CallOption) (*GetRawAccountingDataResponse, error)

	GetNextAccountingDataBatch(context.Context, *GetNextAccountingDataBatchRequest, ...dcerpc.CallOption) (*GetNextAccountingDataBatchResponse, error)

	DeleteAccountingData(context.Context, *DeleteAccountingDataRequest, ...dcerpc.CallOption) (*DeleteAccountingDataResponse, error)

	DefragmentDB(context.Context, *DefragmentDBRequest, ...dcerpc.CallOption) (*DefragmentDBResponse, error)

	CancelAccountingQuery(context.Context, *CancelAccountingQueryRequest, ...dcerpc.CallOption) (*CancelAccountingQueryResponse, error)

	RegisterAccountingClient(context.Context, *RegisterAccountingClientRequest, ...dcerpc.CallOption) (*RegisterAccountingClientResponse, error)

	DumpAccountingData(context.Context, *DumpAccountingDataRequest, ...dcerpc.CallOption) (*DumpAccountingDataResponse, error)

	GetAccountingClients(context.Context, *GetAccountingClientsRequest, ...dcerpc.CallOption) (*GetAccountingClientsResponse, error)

	SetAccountingClientStatus(context.Context, *SetAccountingClientStatusRequest, ...dcerpc.CallOption) (*SetAccountingClientStatusResponse, error)

	CheckAccountingConnection(context.Context, *CheckAccountingConnectionRequest, ...dcerpc.CallOption) (*CheckAccountingConnectionResponse, error)

	SetClientPermissions(context.Context, *SetClientPermissionsRequest, ...dcerpc.CallOption) (*SetClientPermissionsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AccountingClient
}

type xxx_DefaultAccountingClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAccountingClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultAccountingClient) CreateAccountingDB(ctx context.Context, in *CreateAccountingDBRequest, opts ...dcerpc.CallOption) (*CreateAccountingDBResponse, error) {
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
	out := &CreateAccountingDBResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) GetAccountingMetadata(ctx context.Context, in *GetAccountingMetadataRequest, opts ...dcerpc.CallOption) (*GetAccountingMetadataResponse, error) {
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
	out := &GetAccountingMetadataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) ExecuteAccountingQuery(ctx context.Context, in *ExecuteAccountingQueryRequest, opts ...dcerpc.CallOption) (*ExecuteAccountingQueryResponse, error) {
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
	out := &ExecuteAccountingQueryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) GetRawAccountingData(ctx context.Context, in *GetRawAccountingDataRequest, opts ...dcerpc.CallOption) (*GetRawAccountingDataResponse, error) {
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
	out := &GetRawAccountingDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) GetNextAccountingDataBatch(ctx context.Context, in *GetNextAccountingDataBatchRequest, opts ...dcerpc.CallOption) (*GetNextAccountingDataBatchResponse, error) {
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
	out := &GetNextAccountingDataBatchResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) DeleteAccountingData(ctx context.Context, in *DeleteAccountingDataRequest, opts ...dcerpc.CallOption) (*DeleteAccountingDataResponse, error) {
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
	out := &DeleteAccountingDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) DefragmentDB(ctx context.Context, in *DefragmentDBRequest, opts ...dcerpc.CallOption) (*DefragmentDBResponse, error) {
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
	out := &DefragmentDBResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) CancelAccountingQuery(ctx context.Context, in *CancelAccountingQueryRequest, opts ...dcerpc.CallOption) (*CancelAccountingQueryResponse, error) {
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
	out := &CancelAccountingQueryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) RegisterAccountingClient(ctx context.Context, in *RegisterAccountingClientRequest, opts ...dcerpc.CallOption) (*RegisterAccountingClientResponse, error) {
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
	out := &RegisterAccountingClientResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) DumpAccountingData(ctx context.Context, in *DumpAccountingDataRequest, opts ...dcerpc.CallOption) (*DumpAccountingDataResponse, error) {
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
	out := &DumpAccountingDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) GetAccountingClients(ctx context.Context, in *GetAccountingClientsRequest, opts ...dcerpc.CallOption) (*GetAccountingClientsResponse, error) {
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
	out := &GetAccountingClientsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) SetAccountingClientStatus(ctx context.Context, in *SetAccountingClientStatusRequest, opts ...dcerpc.CallOption) (*SetAccountingClientStatusResponse, error) {
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
	out := &SetAccountingClientStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) CheckAccountingConnection(ctx context.Context, in *CheckAccountingConnectionRequest, opts ...dcerpc.CallOption) (*CheckAccountingConnectionResponse, error) {
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
	out := &CheckAccountingConnectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) SetClientPermissions(ctx context.Context, in *SetClientPermissionsRequest, opts ...dcerpc.CallOption) (*SetClientPermissionsResponse, error) {
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
	out := &SetClientPermissionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAccountingClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAccountingClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAccountingClient) IPID(ctx context.Context, ipid *dcom.IPID) AccountingClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAccountingClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewAccountingClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AccountingClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AccountingSyntaxV0_0))...)
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
	return &xxx_DefaultAccountingClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_CreateAccountingDBOperation structure represents the CreateAccountingDb operation
type xxx_CreateAccountingDBOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServerName  *oaut.String   `idl:"name:bstrServerName" json:"server_name"`
	WindowsAuth bool           `idl:"name:bWindowsAuth" json:"windows_auth"`
	UserName    *oaut.String   `idl:"name:bstrUserName" json:"user_name"`
	Password    *oaut.String   `idl:"name:bstrPasswd" json:"password"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateAccountingDBOperation) OpNum() int { return 7 }

func (o *xxx_CreateAccountingDBOperation) OpName() string {
	return "/IWRMAccounting/v0/CreateAccountingDb"
}

func (o *xxx_CreateAccountingDBOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAccountingDBOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrServerName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ServerName != nil {
			_ptr_bstrServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.ServerName, _ptr_bstrServerName); err != nil {
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
	// bWindowsAuth {in} (1:{alias=BOOL}(int32))
	{
		if !o.WindowsAuth {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// bstrUserName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.UserName != nil {
			_ptr_bstrUserName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserName != nil {
					if err := o.UserName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserName, _ptr_bstrUserName); err != nil {
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
	// bstrPasswd {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Password != nil {
			_ptr_bstrPasswd := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Password != nil {
					if err := o.Password.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_bstrPasswd); err != nil {
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

func (o *xxx_CreateAccountingDBOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrServerName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServerName == nil {
				o.ServerName = &oaut.String{}
			}
			if err := o.ServerName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrServerName := func(ptr interface{}) { o.ServerName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ServerName, _s_bstrServerName, _ptr_bstrServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bWindowsAuth {in} (1:{alias=BOOL}(int32))
	{
		var _bWindowsAuth int32
		if err := w.ReadData(&_bWindowsAuth); err != nil {
			return err
		}
		o.WindowsAuth = _bWindowsAuth != 0
	}
	// bstrUserName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrUserName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserName == nil {
				o.UserName = &oaut.String{}
			}
			if err := o.UserName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrUserName := func(ptr interface{}) { o.UserName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.UserName, _s_bstrUserName, _ptr_bstrUserName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPasswd {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPasswd := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Password == nil {
				o.Password = &oaut.String{}
			}
			if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPasswd := func(ptr interface{}) { o.Password = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Password, _s_bstrPasswd, _ptr_bstrPasswd); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAccountingDBOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAccountingDBOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateAccountingDBOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateAccountingDBRequest structure represents the CreateAccountingDb operation request
type CreateAccountingDBRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	ServerName  *oaut.String   `idl:"name:bstrServerName" json:"server_name"`
	WindowsAuth bool           `idl:"name:bWindowsAuth" json:"windows_auth"`
	UserName    *oaut.String   `idl:"name:bstrUserName" json:"user_name"`
	Password    *oaut.String   `idl:"name:bstrPasswd" json:"password"`
}

func (o *CreateAccountingDBRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateAccountingDBOperation) *xxx_CreateAccountingDBOperation {
	if op == nil {
		op = &xxx_CreateAccountingDBOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ServerName = o.ServerName
	op.WindowsAuth = o.WindowsAuth
	op.UserName = o.UserName
	op.Password = o.Password
	return op
}

func (o *CreateAccountingDBRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateAccountingDBOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ServerName = op.ServerName
	o.WindowsAuth = op.WindowsAuth
	o.UserName = op.UserName
	o.Password = op.Password
}
func (o *CreateAccountingDBRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateAccountingDBRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateAccountingDBOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateAccountingDBResponse structure represents the CreateAccountingDb operation response
type CreateAccountingDBResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateAccountingDb return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateAccountingDBResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateAccountingDBOperation) *xxx_CreateAccountingDBOperation {
	if op == nil {
		op = &xxx_CreateAccountingDBOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CreateAccountingDBResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateAccountingDBOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateAccountingDBResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateAccountingDBResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateAccountingDBOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAccountingMetadataOperation structure represents the GetAccountingMetadata operation
type xxx_GetAccountingMetadataOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Metadata *oaut.String   `idl:"name:pbstrMetaData" json:"metadata"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAccountingMetadataOperation) OpNum() int { return 8 }

func (o *xxx_GetAccountingMetadataOperation) OpName() string {
	return "/IWRMAccounting/v0/GetAccountingMetadata"
}

func (o *xxx_GetAccountingMetadataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccountingMetadataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAccountingMetadataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAccountingMetadataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccountingMetadataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrMetaData {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Metadata != nil {
			_ptr_pbstrMetaData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Metadata != nil {
					if err := o.Metadata.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Metadata, _ptr_pbstrMetaData); err != nil {
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

func (o *xxx_GetAccountingMetadataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrMetaData {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrMetaData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Metadata == nil {
				o.Metadata = &oaut.String{}
			}
			if err := o.Metadata.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrMetaData := func(ptr interface{}) { o.Metadata = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Metadata, _s_pbstrMetaData, _ptr_pbstrMetaData); err != nil {
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

// GetAccountingMetadataRequest structure represents the GetAccountingMetadata operation request
type GetAccountingMetadataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAccountingMetadataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAccountingMetadataOperation) *xxx_GetAccountingMetadataOperation {
	if op == nil {
		op = &xxx_GetAccountingMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAccountingMetadataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAccountingMetadataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAccountingMetadataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAccountingMetadataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccountingMetadataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAccountingMetadataResponse structure represents the GetAccountingMetadata operation response
type GetAccountingMetadataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Metadata *oaut.String   `idl:"name:pbstrMetaData" json:"metadata"`
	// Return: The GetAccountingMetadata return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAccountingMetadataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAccountingMetadataOperation) *xxx_GetAccountingMetadataOperation {
	if op == nil {
		op = &xxx_GetAccountingMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Metadata = o.Metadata
	op.Return = o.Return
	return op
}

func (o *GetAccountingMetadataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAccountingMetadataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Metadata = op.Metadata
	o.Return = op.Return
}
func (o *GetAccountingMetadataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAccountingMetadataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccountingMetadataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExecuteAccountingQueryOperation structure represents the ExecuteAccountingQuery operation
type xxx_ExecuteAccountingQueryOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	AccountingQuery *oaut.String   `idl:"name:bstrAccountingQuery" json:"accounting_query"`
	StartingDate    *oaut.String   `idl:"name:bstrStartingDate" json:"starting_date"`
	EndingDate      *oaut.String   `idl:"name:bstrEndingDate" json:"ending_date"`
	Result          *oaut.String   `idl:"name:pbstrResult" json:"result"`
	IsThereMoreData bool           `idl:"name:pbIsThereMoreData" json:"is_there_more_data"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ExecuteAccountingQueryOperation) OpNum() int { return 9 }

func (o *xxx_ExecuteAccountingQueryOperation) OpName() string {
	return "/IWRMAccounting/v0/ExecuteAccountingQuery"
}

func (o *xxx_ExecuteAccountingQueryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecuteAccountingQueryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrAccountingQuery {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.AccountingQuery != nil {
			_ptr_bstrAccountingQuery := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AccountingQuery != nil {
					if err := o.AccountingQuery.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AccountingQuery, _ptr_bstrAccountingQuery); err != nil {
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
	// bstrStartingDate {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.StartingDate != nil {
			_ptr_bstrStartingDate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.StartingDate != nil {
					if err := o.StartingDate.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.StartingDate, _ptr_bstrStartingDate); err != nil {
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
	// bstrEndingDate {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EndingDate != nil {
			_ptr_bstrEndingDate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EndingDate != nil {
					if err := o.EndingDate.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EndingDate, _ptr_bstrEndingDate); err != nil {
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

func (o *xxx_ExecuteAccountingQueryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrAccountingQuery {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrAccountingQuery := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AccountingQuery == nil {
				o.AccountingQuery = &oaut.String{}
			}
			if err := o.AccountingQuery.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrAccountingQuery := func(ptr interface{}) { o.AccountingQuery = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.AccountingQuery, _s_bstrAccountingQuery, _ptr_bstrAccountingQuery); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrStartingDate {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrStartingDate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.StartingDate == nil {
				o.StartingDate = &oaut.String{}
			}
			if err := o.StartingDate.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrStartingDate := func(ptr interface{}) { o.StartingDate = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.StartingDate, _s_bstrStartingDate, _ptr_bstrStartingDate); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrEndingDate {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrEndingDate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EndingDate == nil {
				o.EndingDate = &oaut.String{}
			}
			if err := o.EndingDate.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrEndingDate := func(ptr interface{}) { o.EndingDate = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EndingDate, _s_bstrEndingDate, _ptr_bstrEndingDate); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecuteAccountingQueryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecuteAccountingQueryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrResult {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Result != nil {
			_ptr_pbstrResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Result != nil {
					if err := o.Result.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Result, _ptr_pbstrResult); err != nil {
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
	// pbIsThereMoreData {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.IsThereMoreData {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
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

func (o *xxx_ExecuteAccountingQueryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrResult {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Result == nil {
				o.Result = &oaut.String{}
			}
			if err := o.Result.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrResult := func(ptr interface{}) { o.Result = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Result, _s_pbstrResult, _ptr_pbstrResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbIsThereMoreData {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bIsThereMoreData int32
		if err := w.ReadData(&_bIsThereMoreData); err != nil {
			return err
		}
		o.IsThereMoreData = _bIsThereMoreData != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ExecuteAccountingQueryRequest structure represents the ExecuteAccountingQuery operation request
type ExecuteAccountingQueryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	AccountingQuery *oaut.String   `idl:"name:bstrAccountingQuery" json:"accounting_query"`
	StartingDate    *oaut.String   `idl:"name:bstrStartingDate" json:"starting_date"`
	EndingDate      *oaut.String   `idl:"name:bstrEndingDate" json:"ending_date"`
}

func (o *ExecuteAccountingQueryRequest) xxx_ToOp(ctx context.Context, op *xxx_ExecuteAccountingQueryOperation) *xxx_ExecuteAccountingQueryOperation {
	if op == nil {
		op = &xxx_ExecuteAccountingQueryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.AccountingQuery = o.AccountingQuery
	op.StartingDate = o.StartingDate
	op.EndingDate = o.EndingDate
	return op
}

func (o *ExecuteAccountingQueryRequest) xxx_FromOp(ctx context.Context, op *xxx_ExecuteAccountingQueryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AccountingQuery = op.AccountingQuery
	o.StartingDate = op.StartingDate
	o.EndingDate = op.EndingDate
}
func (o *ExecuteAccountingQueryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ExecuteAccountingQueryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecuteAccountingQueryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExecuteAccountingQueryResponse structure represents the ExecuteAccountingQuery operation response
type ExecuteAccountingQueryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	Result          *oaut.String   `idl:"name:pbstrResult" json:"result"`
	IsThereMoreData bool           `idl:"name:pbIsThereMoreData" json:"is_there_more_data"`
	// Return: The ExecuteAccountingQuery return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExecuteAccountingQueryResponse) xxx_ToOp(ctx context.Context, op *xxx_ExecuteAccountingQueryOperation) *xxx_ExecuteAccountingQueryOperation {
	if op == nil {
		op = &xxx_ExecuteAccountingQueryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Result = o.Result
	op.IsThereMoreData = o.IsThereMoreData
	op.Return = o.Return
	return op
}

func (o *ExecuteAccountingQueryResponse) xxx_FromOp(ctx context.Context, op *xxx_ExecuteAccountingQueryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Result = op.Result
	o.IsThereMoreData = op.IsThereMoreData
	o.Return = op.Return
}
func (o *ExecuteAccountingQueryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ExecuteAccountingQueryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecuteAccountingQueryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRawAccountingDataOperation structure represents the GetRawAccountingData operation
type xxx_GetRawAccountingDataOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	StartingDate    *oaut.String   `idl:"name:bstrStartingDate" json:"starting_date"`
	EndingDate      *oaut.String   `idl:"name:bstrEndingDate" json:"ending_date"`
	MachineName     *oaut.String   `idl:"name:bstrMachineName" json:"machine_name"`
	Result          *oaut.String   `idl:"name:pbstrResult" json:"result"`
	IsThereMoreData bool           `idl:"name:pbIsThereMoreData" json:"is_there_more_data"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRawAccountingDataOperation) OpNum() int { return 10 }

func (o *xxx_GetRawAccountingDataOperation) OpName() string {
	return "/IWRMAccounting/v0/GetRawAccountingData"
}

func (o *xxx_GetRawAccountingDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRawAccountingDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrStartingDate {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.StartingDate != nil {
			_ptr_bstrStartingDate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.StartingDate != nil {
					if err := o.StartingDate.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.StartingDate, _ptr_bstrStartingDate); err != nil {
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
	// bstrEndingDate {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EndingDate != nil {
			_ptr_bstrEndingDate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EndingDate != nil {
					if err := o.EndingDate.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EndingDate, _ptr_bstrEndingDate); err != nil {
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
	// bstrMachineName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineName != nil {
			_ptr_bstrMachineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineName != nil {
					if err := o.MachineName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineName, _ptr_bstrMachineName); err != nil {
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

func (o *xxx_GetRawAccountingDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrStartingDate {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrStartingDate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.StartingDate == nil {
				o.StartingDate = &oaut.String{}
			}
			if err := o.StartingDate.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrStartingDate := func(ptr interface{}) { o.StartingDate = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.StartingDate, _s_bstrStartingDate, _ptr_bstrStartingDate); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrEndingDate {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrEndingDate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EndingDate == nil {
				o.EndingDate = &oaut.String{}
			}
			if err := o.EndingDate.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrEndingDate := func(ptr interface{}) { o.EndingDate = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EndingDate, _s_bstrEndingDate, _ptr_bstrEndingDate); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMachineName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineName == nil {
				o.MachineName = &oaut.String{}
			}
			if err := o.MachineName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineName := func(ptr interface{}) { o.MachineName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineName, _s_bstrMachineName, _ptr_bstrMachineName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRawAccountingDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRawAccountingDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrResult {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Result != nil {
			_ptr_pbstrResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Result != nil {
					if err := o.Result.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Result, _ptr_pbstrResult); err != nil {
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
	// pbIsThereMoreData {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.IsThereMoreData {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
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

func (o *xxx_GetRawAccountingDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrResult {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Result == nil {
				o.Result = &oaut.String{}
			}
			if err := o.Result.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrResult := func(ptr interface{}) { o.Result = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Result, _s_pbstrResult, _ptr_pbstrResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbIsThereMoreData {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bIsThereMoreData int32
		if err := w.ReadData(&_bIsThereMoreData); err != nil {
			return err
		}
		o.IsThereMoreData = _bIsThereMoreData != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetRawAccountingDataRequest structure represents the GetRawAccountingData operation request
type GetRawAccountingDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	StartingDate *oaut.String   `idl:"name:bstrStartingDate" json:"starting_date"`
	EndingDate   *oaut.String   `idl:"name:bstrEndingDate" json:"ending_date"`
	MachineName  *oaut.String   `idl:"name:bstrMachineName" json:"machine_name"`
}

func (o *GetRawAccountingDataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRawAccountingDataOperation) *xxx_GetRawAccountingDataOperation {
	if op == nil {
		op = &xxx_GetRawAccountingDataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.StartingDate = o.StartingDate
	op.EndingDate = o.EndingDate
	op.MachineName = o.MachineName
	return op
}

func (o *GetRawAccountingDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRawAccountingDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.StartingDate = op.StartingDate
	o.EndingDate = op.EndingDate
	o.MachineName = op.MachineName
}
func (o *GetRawAccountingDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRawAccountingDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRawAccountingDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRawAccountingDataResponse structure represents the GetRawAccountingData operation response
type GetRawAccountingDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	Result          *oaut.String   `idl:"name:pbstrResult" json:"result"`
	IsThereMoreData bool           `idl:"name:pbIsThereMoreData" json:"is_there_more_data"`
	// Return: The GetRawAccountingData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRawAccountingDataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRawAccountingDataOperation) *xxx_GetRawAccountingDataOperation {
	if op == nil {
		op = &xxx_GetRawAccountingDataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Result = o.Result
	op.IsThereMoreData = o.IsThereMoreData
	op.Return = o.Return
	return op
}

func (o *GetRawAccountingDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRawAccountingDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Result = op.Result
	o.IsThereMoreData = op.IsThereMoreData
	o.Return = op.Return
}
func (o *GetRawAccountingDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRawAccountingDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRawAccountingDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNextAccountingDataBatchOperation structure represents the GetNextAccountingDataBatch operation
type xxx_GetNextAccountingDataBatchOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	Result          *oaut.String   `idl:"name:pbstrResult" json:"result"`
	IsThereMoreData bool           `idl:"name:pbIsThereMoreData" json:"is_there_more_data"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNextAccountingDataBatchOperation) OpNum() int { return 11 }

func (o *xxx_GetNextAccountingDataBatchOperation) OpName() string {
	return "/IWRMAccounting/v0/GetNextAccountingDataBatch"
}

func (o *xxx_GetNextAccountingDataBatchOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextAccountingDataBatchOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetNextAccountingDataBatchOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetNextAccountingDataBatchOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextAccountingDataBatchOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrResult {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Result != nil {
			_ptr_pbstrResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Result != nil {
					if err := o.Result.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Result, _ptr_pbstrResult); err != nil {
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
	// pbIsThereMoreData {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.IsThereMoreData {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
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

func (o *xxx_GetNextAccountingDataBatchOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrResult {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Result == nil {
				o.Result = &oaut.String{}
			}
			if err := o.Result.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrResult := func(ptr interface{}) { o.Result = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Result, _s_pbstrResult, _ptr_pbstrResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbIsThereMoreData {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bIsThereMoreData int32
		if err := w.ReadData(&_bIsThereMoreData); err != nil {
			return err
		}
		o.IsThereMoreData = _bIsThereMoreData != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNextAccountingDataBatchRequest structure represents the GetNextAccountingDataBatch operation request
type GetNextAccountingDataBatchRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetNextAccountingDataBatchRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNextAccountingDataBatchOperation) *xxx_GetNextAccountingDataBatchOperation {
	if op == nil {
		op = &xxx_GetNextAccountingDataBatchOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetNextAccountingDataBatchRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNextAccountingDataBatchOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNextAccountingDataBatchRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNextAccountingDataBatchRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNextAccountingDataBatchOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNextAccountingDataBatchResponse structure represents the GetNextAccountingDataBatch operation response
type GetNextAccountingDataBatchResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	Result          *oaut.String   `idl:"name:pbstrResult" json:"result"`
	IsThereMoreData bool           `idl:"name:pbIsThereMoreData" json:"is_there_more_data"`
	// Return: The GetNextAccountingDataBatch return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNextAccountingDataBatchResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNextAccountingDataBatchOperation) *xxx_GetNextAccountingDataBatchOperation {
	if op == nil {
		op = &xxx_GetNextAccountingDataBatchOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Result = o.Result
	op.IsThereMoreData = o.IsThereMoreData
	op.Return = o.Return
	return op
}

func (o *GetNextAccountingDataBatchResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNextAccountingDataBatchOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Result = op.Result
	o.IsThereMoreData = op.IsThereMoreData
	o.Return = op.Return
}
func (o *GetNextAccountingDataBatchResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNextAccountingDataBatchResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNextAccountingDataBatchOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteAccountingDataOperation structure represents the DeleteAccountingData operation
type xxx_DeleteAccountingDataOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	StartingDate *oaut.String   `idl:"name:bstrStartingDate" json:"starting_date"`
	EndingDate   *oaut.String   `idl:"name:bstrEndingDate" json:"ending_date"`
	MachineName  *oaut.String   `idl:"name:bstrMachineName" json:"machine_name"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteAccountingDataOperation) OpNum() int { return 12 }

func (o *xxx_DeleteAccountingDataOperation) OpName() string {
	return "/IWRMAccounting/v0/DeleteAccountingData"
}

func (o *xxx_DeleteAccountingDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAccountingDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrStartingDate {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.StartingDate != nil {
			_ptr_bstrStartingDate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.StartingDate != nil {
					if err := o.StartingDate.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.StartingDate, _ptr_bstrStartingDate); err != nil {
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
	// bstrEndingDate {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EndingDate != nil {
			_ptr_bstrEndingDate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EndingDate != nil {
					if err := o.EndingDate.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EndingDate, _ptr_bstrEndingDate); err != nil {
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
	// bstrMachineName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineName != nil {
			_ptr_bstrMachineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineName != nil {
					if err := o.MachineName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineName, _ptr_bstrMachineName); err != nil {
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

func (o *xxx_DeleteAccountingDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrStartingDate {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrStartingDate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.StartingDate == nil {
				o.StartingDate = &oaut.String{}
			}
			if err := o.StartingDate.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrStartingDate := func(ptr interface{}) { o.StartingDate = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.StartingDate, _s_bstrStartingDate, _ptr_bstrStartingDate); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrEndingDate {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrEndingDate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EndingDate == nil {
				o.EndingDate = &oaut.String{}
			}
			if err := o.EndingDate.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrEndingDate := func(ptr interface{}) { o.EndingDate = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EndingDate, _s_bstrEndingDate, _ptr_bstrEndingDate); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMachineName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineName == nil {
				o.MachineName = &oaut.String{}
			}
			if err := o.MachineName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineName := func(ptr interface{}) { o.MachineName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineName, _s_bstrMachineName, _ptr_bstrMachineName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAccountingDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAccountingDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteAccountingDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteAccountingDataRequest structure represents the DeleteAccountingData operation request
type DeleteAccountingDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	StartingDate *oaut.String   `idl:"name:bstrStartingDate" json:"starting_date"`
	EndingDate   *oaut.String   `idl:"name:bstrEndingDate" json:"ending_date"`
	MachineName  *oaut.String   `idl:"name:bstrMachineName" json:"machine_name"`
}

func (o *DeleteAccountingDataRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteAccountingDataOperation) *xxx_DeleteAccountingDataOperation {
	if op == nil {
		op = &xxx_DeleteAccountingDataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.StartingDate = o.StartingDate
	op.EndingDate = o.EndingDate
	op.MachineName = o.MachineName
	return op
}

func (o *DeleteAccountingDataRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteAccountingDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.StartingDate = op.StartingDate
	o.EndingDate = op.EndingDate
	o.MachineName = op.MachineName
}
func (o *DeleteAccountingDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteAccountingDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteAccountingDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteAccountingDataResponse structure represents the DeleteAccountingData operation response
type DeleteAccountingDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteAccountingData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteAccountingDataResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteAccountingDataOperation) *xxx_DeleteAccountingDataOperation {
	if op == nil {
		op = &xxx_DeleteAccountingDataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteAccountingDataResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteAccountingDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteAccountingDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteAccountingDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteAccountingDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DefragmentDBOperation structure represents the DefragmentDB operation
type xxx_DefragmentDBOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DefragmentDBOperation) OpNum() int { return 13 }

func (o *xxx_DefragmentDBOperation) OpName() string { return "/IWRMAccounting/v0/DefragmentDB" }

func (o *xxx_DefragmentDBOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DefragmentDBOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DefragmentDBOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_DefragmentDBOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DefragmentDBOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DefragmentDBOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DefragmentDBRequest structure represents the DefragmentDB operation request
type DefragmentDBRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *DefragmentDBRequest) xxx_ToOp(ctx context.Context, op *xxx_DefragmentDBOperation) *xxx_DefragmentDBOperation {
	if op == nil {
		op = &xxx_DefragmentDBOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *DefragmentDBRequest) xxx_FromOp(ctx context.Context, op *xxx_DefragmentDBOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *DefragmentDBRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DefragmentDBRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DefragmentDBOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DefragmentDBResponse structure represents the DefragmentDB operation response
type DefragmentDBResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DefragmentDB return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DefragmentDBResponse) xxx_ToOp(ctx context.Context, op *xxx_DefragmentDBOperation) *xxx_DefragmentDBOperation {
	if op == nil {
		op = &xxx_DefragmentDBOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DefragmentDBResponse) xxx_FromOp(ctx context.Context, op *xxx_DefragmentDBOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DefragmentDBResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DefragmentDBResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DefragmentDBOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CancelAccountingQueryOperation structure represents the CancelAccountingQuery operation
type xxx_CancelAccountingQueryOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Flag   bool           `idl:"name:flag" json:"flag"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CancelAccountingQueryOperation) OpNum() int { return 14 }

func (o *xxx_CancelAccountingQueryOperation) OpName() string {
	return "/IWRMAccounting/v0/CancelAccountingQuery"
}

func (o *xxx_CancelAccountingQueryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelAccountingQueryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// flag {in} (1:{alias=BOOL}(int32))
	{
		if !o.Flag {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CancelAccountingQueryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// flag {in} (1:{alias=BOOL}(int32))
	{
		var _bFlag int32
		if err := w.ReadData(&_bFlag); err != nil {
			return err
		}
		o.Flag = _bFlag != 0
	}
	return nil
}

func (o *xxx_CancelAccountingQueryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelAccountingQueryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CancelAccountingQueryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CancelAccountingQueryRequest structure represents the CancelAccountingQuery operation request
type CancelAccountingQueryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Flag bool           `idl:"name:flag" json:"flag"`
}

func (o *CancelAccountingQueryRequest) xxx_ToOp(ctx context.Context, op *xxx_CancelAccountingQueryOperation) *xxx_CancelAccountingQueryOperation {
	if op == nil {
		op = &xxx_CancelAccountingQueryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Flag = o.Flag
	return op
}

func (o *CancelAccountingQueryRequest) xxx_FromOp(ctx context.Context, op *xxx_CancelAccountingQueryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flag = op.Flag
}
func (o *CancelAccountingQueryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CancelAccountingQueryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelAccountingQueryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CancelAccountingQueryResponse structure represents the CancelAccountingQuery operation response
type CancelAccountingQueryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CancelAccountingQuery return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CancelAccountingQueryResponse) xxx_ToOp(ctx context.Context, op *xxx_CancelAccountingQueryOperation) *xxx_CancelAccountingQueryOperation {
	if op == nil {
		op = &xxx_CancelAccountingQueryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CancelAccountingQueryResponse) xxx_FromOp(ctx context.Context, op *xxx_CancelAccountingQueryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CancelAccountingQueryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CancelAccountingQueryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelAccountingQueryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterAccountingClientOperation structure represents the RegisterAccountingClient operation
type xxx_RegisterAccountingClientOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	ClientID *oaut.String   `idl:"name:bstrClientId" json:"client_id"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterAccountingClientOperation) OpNum() int { return 15 }

func (o *xxx_RegisterAccountingClientOperation) OpName() string {
	return "/IWRMAccounting/v0/RegisterAccountingClient"
}

func (o *xxx_RegisterAccountingClientOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterAccountingClientOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrClientId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ClientID != nil {
			_ptr_bstrClientId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientID != nil {
					if err := o.ClientID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientID, _ptr_bstrClientId); err != nil {
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

func (o *xxx_RegisterAccountingClientOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrClientId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrClientId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientID == nil {
				o.ClientID = &oaut.String{}
			}
			if err := o.ClientID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrClientId := func(ptr interface{}) { o.ClientID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ClientID, _s_bstrClientId, _ptr_bstrClientId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterAccountingClientOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterAccountingClientOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RegisterAccountingClientOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RegisterAccountingClientRequest structure represents the RegisterAccountingClient operation request
type RegisterAccountingClientRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	ClientID *oaut.String   `idl:"name:bstrClientId" json:"client_id"`
}

func (o *RegisterAccountingClientRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterAccountingClientOperation) *xxx_RegisterAccountingClientOperation {
	if op == nil {
		op = &xxx_RegisterAccountingClientOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ClientID = o.ClientID
	return op
}

func (o *RegisterAccountingClientRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterAccountingClientOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClientID = op.ClientID
}
func (o *RegisterAccountingClientRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterAccountingClientRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterAccountingClientOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterAccountingClientResponse structure represents the RegisterAccountingClient operation response
type RegisterAccountingClientResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RegisterAccountingClient return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterAccountingClientResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterAccountingClientOperation) *xxx_RegisterAccountingClientOperation {
	if op == nil {
		op = &xxx_RegisterAccountingClientOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RegisterAccountingClientResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterAccountingClientOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RegisterAccountingClientResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterAccountingClientResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterAccountingClientOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DumpAccountingDataOperation structure represents the DumpAccountingData operation
type xxx_DumpAccountingDataOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	AccountingData *oaut.String   `idl:"name:bstrAccountingData" json:"accounting_data"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DumpAccountingDataOperation) OpNum() int { return 16 }

func (o *xxx_DumpAccountingDataOperation) OpName() string {
	return "/IWRMAccounting/v0/DumpAccountingData"
}

func (o *xxx_DumpAccountingDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DumpAccountingDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrAccountingData {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.AccountingData != nil {
			_ptr_bstrAccountingData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AccountingData != nil {
					if err := o.AccountingData.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AccountingData, _ptr_bstrAccountingData); err != nil {
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

func (o *xxx_DumpAccountingDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrAccountingData {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrAccountingData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AccountingData == nil {
				o.AccountingData = &oaut.String{}
			}
			if err := o.AccountingData.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrAccountingData := func(ptr interface{}) { o.AccountingData = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.AccountingData, _s_bstrAccountingData, _ptr_bstrAccountingData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DumpAccountingDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DumpAccountingDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DumpAccountingDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DumpAccountingDataRequest structure represents the DumpAccountingData operation request
type DumpAccountingDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	AccountingData *oaut.String   `idl:"name:bstrAccountingData" json:"accounting_data"`
}

func (o *DumpAccountingDataRequest) xxx_ToOp(ctx context.Context, op *xxx_DumpAccountingDataOperation) *xxx_DumpAccountingDataOperation {
	if op == nil {
		op = &xxx_DumpAccountingDataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.AccountingData = o.AccountingData
	return op
}

func (o *DumpAccountingDataRequest) xxx_FromOp(ctx context.Context, op *xxx_DumpAccountingDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AccountingData = op.AccountingData
}
func (o *DumpAccountingDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DumpAccountingDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DumpAccountingDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DumpAccountingDataResponse structure represents the DumpAccountingData operation response
type DumpAccountingDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DumpAccountingData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DumpAccountingDataResponse) xxx_ToOp(ctx context.Context, op *xxx_DumpAccountingDataOperation) *xxx_DumpAccountingDataOperation {
	if op == nil {
		op = &xxx_DumpAccountingDataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DumpAccountingDataResponse) xxx_FromOp(ctx context.Context, op *xxx_DumpAccountingDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DumpAccountingDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DumpAccountingDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DumpAccountingDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAccountingClientsOperation structure represents the GetAccountingClients operation
type xxx_GetAccountingClientsOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ClientIDs *oaut.String   `idl:"name:pbstrClientIds" json:"client_ids"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAccountingClientsOperation) OpNum() int { return 17 }

func (o *xxx_GetAccountingClientsOperation) OpName() string {
	return "/IWRMAccounting/v0/GetAccountingClients"
}

func (o *xxx_GetAccountingClientsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccountingClientsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAccountingClientsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAccountingClientsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccountingClientsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrClientIds {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ClientIDs != nil {
			_ptr_pbstrClientIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientIDs != nil {
					if err := o.ClientIDs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientIDs, _ptr_pbstrClientIds); err != nil {
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

func (o *xxx_GetAccountingClientsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrClientIds {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrClientIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientIDs == nil {
				o.ClientIDs = &oaut.String{}
			}
			if err := o.ClientIDs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrClientIds := func(ptr interface{}) { o.ClientIDs = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ClientIDs, _s_pbstrClientIds, _ptr_pbstrClientIds); err != nil {
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

// GetAccountingClientsRequest structure represents the GetAccountingClients operation request
type GetAccountingClientsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAccountingClientsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAccountingClientsOperation) *xxx_GetAccountingClientsOperation {
	if op == nil {
		op = &xxx_GetAccountingClientsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAccountingClientsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAccountingClientsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAccountingClientsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAccountingClientsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccountingClientsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAccountingClientsResponse structure represents the GetAccountingClients operation response
type GetAccountingClientsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ClientIDs *oaut.String   `idl:"name:pbstrClientIds" json:"client_ids"`
	// Return: The GetAccountingClients return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAccountingClientsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAccountingClientsOperation) *xxx_GetAccountingClientsOperation {
	if op == nil {
		op = &xxx_GetAccountingClientsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ClientIDs = o.ClientIDs
	op.Return = o.Return
	return op
}

func (o *GetAccountingClientsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAccountingClientsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ClientIDs = op.ClientIDs
	o.Return = op.Return
}
func (o *GetAccountingClientsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAccountingClientsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccountingClientsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAccountingClientStatusOperation structure represents the SetAccountingClientStatus operation
type xxx_SetAccountingClientStatusOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ClientIDs *oaut.String   `idl:"name:bstrClientIds" json:"client_ids"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAccountingClientStatusOperation) OpNum() int { return 18 }

func (o *xxx_SetAccountingClientStatusOperation) OpName() string {
	return "/IWRMAccounting/v0/SetAccountingClientStatus"
}

func (o *xxx_SetAccountingClientStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountingClientStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrClientIds {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ClientIDs != nil {
			_ptr_bstrClientIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientIDs != nil {
					if err := o.ClientIDs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientIDs, _ptr_bstrClientIds); err != nil {
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

func (o *xxx_SetAccountingClientStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrClientIds {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrClientIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientIDs == nil {
				o.ClientIDs = &oaut.String{}
			}
			if err := o.ClientIDs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrClientIds := func(ptr interface{}) { o.ClientIDs = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ClientIDs, _s_bstrClientIds, _ptr_bstrClientIds); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountingClientStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountingClientStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAccountingClientStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAccountingClientStatusRequest structure represents the SetAccountingClientStatus operation request
type SetAccountingClientStatusRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	ClientIDs *oaut.String   `idl:"name:bstrClientIds" json:"client_ids"`
}

func (o *SetAccountingClientStatusRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAccountingClientStatusOperation) *xxx_SetAccountingClientStatusOperation {
	if op == nil {
		op = &xxx_SetAccountingClientStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ClientIDs = o.ClientIDs
	return op
}

func (o *SetAccountingClientStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAccountingClientStatusOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClientIDs = op.ClientIDs
}
func (o *SetAccountingClientStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAccountingClientStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAccountingClientStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAccountingClientStatusResponse structure represents the SetAccountingClientStatus operation response
type SetAccountingClientStatusResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetAccountingClientStatus return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAccountingClientStatusResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAccountingClientStatusOperation) *xxx_SetAccountingClientStatusOperation {
	if op == nil {
		op = &xxx_SetAccountingClientStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAccountingClientStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAccountingClientStatusOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAccountingClientStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAccountingClientStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAccountingClientStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CheckAccountingConnectionOperation structure represents the CheckAccountingConnection operation
type xxx_CheckAccountingConnectionOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CheckAccountingConnectionOperation) OpNum() int { return 19 }

func (o *xxx_CheckAccountingConnectionOperation) OpName() string {
	return "/IWRMAccounting/v0/CheckAccountingConnection"
}

func (o *xxx_CheckAccountingConnectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckAccountingConnectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CheckAccountingConnectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CheckAccountingConnectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckAccountingConnectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CheckAccountingConnectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CheckAccountingConnectionRequest structure represents the CheckAccountingConnection operation request
type CheckAccountingConnectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CheckAccountingConnectionRequest) xxx_ToOp(ctx context.Context, op *xxx_CheckAccountingConnectionOperation) *xxx_CheckAccountingConnectionOperation {
	if op == nil {
		op = &xxx_CheckAccountingConnectionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CheckAccountingConnectionRequest) xxx_FromOp(ctx context.Context, op *xxx_CheckAccountingConnectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CheckAccountingConnectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CheckAccountingConnectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CheckAccountingConnectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CheckAccountingConnectionResponse structure represents the CheckAccountingConnection operation response
type CheckAccountingConnectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CheckAccountingConnection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CheckAccountingConnectionResponse) xxx_ToOp(ctx context.Context, op *xxx_CheckAccountingConnectionOperation) *xxx_CheckAccountingConnectionOperation {
	if op == nil {
		op = &xxx_CheckAccountingConnectionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CheckAccountingConnectionResponse) xxx_FromOp(ctx context.Context, op *xxx_CheckAccountingConnectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CheckAccountingConnectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CheckAccountingConnectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CheckAccountingConnectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetClientPermissionsOperation structure represents the SetClientPermissions operation
type xxx_SetClientPermissionsOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	ClientID       *oaut.String   `idl:"name:bstrClientId" json:"client_id"`
	AddPermissions bool           `idl:"name:fAddPermissions" json:"add_permissions"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetClientPermissionsOperation) OpNum() int { return 20 }

func (o *xxx_SetClientPermissionsOperation) OpName() string {
	return "/IWRMAccounting/v0/SetClientPermissions"
}

func (o *xxx_SetClientPermissionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientPermissionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrClientId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ClientID != nil {
			_ptr_bstrClientId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientID != nil {
					if err := o.ClientID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientID, _ptr_bstrClientId); err != nil {
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
	// fAddPermissions {in} (1:{alias=BOOL}(int32))
	{
		if !o.AddPermissions {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetClientPermissionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrClientId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrClientId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientID == nil {
				o.ClientID = &oaut.String{}
			}
			if err := o.ClientID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrClientId := func(ptr interface{}) { o.ClientID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ClientID, _s_bstrClientId, _ptr_bstrClientId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fAddPermissions {in} (1:{alias=BOOL}(int32))
	{
		var _bAddPermissions int32
		if err := w.ReadData(&_bAddPermissions); err != nil {
			return err
		}
		o.AddPermissions = _bAddPermissions != 0
	}
	return nil
}

func (o *xxx_SetClientPermissionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientPermissionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetClientPermissionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetClientPermissionsRequest structure represents the SetClientPermissions operation request
type SetClientPermissionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	ClientID       *oaut.String   `idl:"name:bstrClientId" json:"client_id"`
	AddPermissions bool           `idl:"name:fAddPermissions" json:"add_permissions"`
}

func (o *SetClientPermissionsRequest) xxx_ToOp(ctx context.Context, op *xxx_SetClientPermissionsOperation) *xxx_SetClientPermissionsOperation {
	if op == nil {
		op = &xxx_SetClientPermissionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ClientID = o.ClientID
	op.AddPermissions = o.AddPermissions
	return op
}

func (o *SetClientPermissionsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetClientPermissionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClientID = op.ClientID
	o.AddPermissions = op.AddPermissions
}
func (o *SetClientPermissionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetClientPermissionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientPermissionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetClientPermissionsResponse structure represents the SetClientPermissions operation response
type SetClientPermissionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetClientPermissions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetClientPermissionsResponse) xxx_ToOp(ctx context.Context, op *xxx_SetClientPermissionsOperation) *xxx_SetClientPermissionsOperation {
	if op == nil {
		op = &xxx_SetClientPermissionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetClientPermissionsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetClientPermissionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetClientPermissionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetClientPermissionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientPermissionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
