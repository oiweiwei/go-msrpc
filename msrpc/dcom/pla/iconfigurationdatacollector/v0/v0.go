package iconfigurationdatacollector

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
	idatacollector "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/idatacollector/v0"
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
	_ = idatacollector.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/pla"
)

var (
	// IConfigurationDataCollector interface identifier 03837514-098b-11d8-9414-505054503030
	ConfigurationDataCollectorIID = &dcom.IID{Data1: 0x03837514, Data2: 0x098b, Data3: 0x11d8, Data4: []byte{0x94, 0x14, 0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax UUID
	ConfigurationDataCollectorSyntaxUUID = &uuid.UUID{TimeLow: 0x3837514, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax ID
	ConfigurationDataCollectorSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ConfigurationDataCollectorSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IConfigurationDataCollector interface.
type ConfigurationDataCollectorClient interface {

	// IDataCollector retrieval method.
	DataCollector() idatacollector.DataCollectorClient

	// FileMaxCount operation.
	GetFileMaxCount(context.Context, *GetFileMaxCountRequest, ...dcerpc.CallOption) (*GetFileMaxCountResponse, error)

	// FileMaxCount operation.
	SetFileMaxCount(context.Context, *SetFileMaxCountRequest, ...dcerpc.CallOption) (*SetFileMaxCountResponse, error)

	// FileMaxRecursiveDepth operation.
	GetFileMaxRecursiveDepth(context.Context, *GetFileMaxRecursiveDepthRequest, ...dcerpc.CallOption) (*GetFileMaxRecursiveDepthResponse, error)

	// FileMaxRecursiveDepth operation.
	SetFileMaxRecursiveDepth(context.Context, *SetFileMaxRecursiveDepthRequest, ...dcerpc.CallOption) (*SetFileMaxRecursiveDepthResponse, error)

	// FileMaxTotalSize operation.
	GetFileMaxTotalSize(context.Context, *GetFileMaxTotalSizeRequest, ...dcerpc.CallOption) (*GetFileMaxTotalSizeResponse, error)

	// FileMaxTotalSize operation.
	SetFileMaxTotalSize(context.Context, *SetFileMaxTotalSizeRequest, ...dcerpc.CallOption) (*SetFileMaxTotalSizeResponse, error)

	// Files operation.
	GetFiles(context.Context, *GetFilesRequest, ...dcerpc.CallOption) (*GetFilesResponse, error)

	// Files operation.
	SetFiles(context.Context, *SetFilesRequest, ...dcerpc.CallOption) (*SetFilesResponse, error)

	// ManagementQueries operation.
	GetManagementQueries(context.Context, *GetManagementQueriesRequest, ...dcerpc.CallOption) (*GetManagementQueriesResponse, error)

	// ManagementQueries operation.
	SetManagementQueries(context.Context, *SetManagementQueriesRequest, ...dcerpc.CallOption) (*SetManagementQueriesResponse, error)

	// QueryNetworkAdapters operation.
	GetQueryNetworkAdapters(context.Context, *GetQueryNetworkAdaptersRequest, ...dcerpc.CallOption) (*GetQueryNetworkAdaptersResponse, error)

	// QueryNetworkAdapters operation.
	SetQueryNetworkAdapters(context.Context, *SetQueryNetworkAdaptersRequest, ...dcerpc.CallOption) (*SetQueryNetworkAdaptersResponse, error)

	// RegistryKeys operation.
	GetRegistryKeys(context.Context, *GetRegistryKeysRequest, ...dcerpc.CallOption) (*GetRegistryKeysResponse, error)

	// RegistryKeys operation.
	SetRegistryKeys(context.Context, *SetRegistryKeysRequest, ...dcerpc.CallOption) (*SetRegistryKeysResponse, error)

	// RegistryMaxRecursiveDepth operation.
	GetRegistryMaxRecursiveDepth(context.Context, *GetRegistryMaxRecursiveDepthRequest, ...dcerpc.CallOption) (*GetRegistryMaxRecursiveDepthResponse, error)

	// RegistryMaxRecursiveDepth operation.
	SetRegistryMaxRecursiveDepth(context.Context, *SetRegistryMaxRecursiveDepthRequest, ...dcerpc.CallOption) (*SetRegistryMaxRecursiveDepthResponse, error)

	// SystemStateFile operation.
	GetSystemStateFile(context.Context, *GetSystemStateFileRequest, ...dcerpc.CallOption) (*GetSystemStateFileResponse, error)

	// SystemStateFile operation.
	SetSystemStateFile(context.Context, *SetSystemStateFileRequest, ...dcerpc.CallOption) (*SetSystemStateFileResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ConfigurationDataCollectorClient
}

type xxx_DefaultConfigurationDataCollectorClient struct {
	idatacollector.DataCollectorClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultConfigurationDataCollectorClient) DataCollector() idatacollector.DataCollectorClient {
	return o.DataCollectorClient
}

func (o *xxx_DefaultConfigurationDataCollectorClient) GetFileMaxCount(ctx context.Context, in *GetFileMaxCountRequest, opts ...dcerpc.CallOption) (*GetFileMaxCountResponse, error) {
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
	out := &GetFileMaxCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) SetFileMaxCount(ctx context.Context, in *SetFileMaxCountRequest, opts ...dcerpc.CallOption) (*SetFileMaxCountResponse, error) {
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
	out := &SetFileMaxCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) GetFileMaxRecursiveDepth(ctx context.Context, in *GetFileMaxRecursiveDepthRequest, opts ...dcerpc.CallOption) (*GetFileMaxRecursiveDepthResponse, error) {
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
	out := &GetFileMaxRecursiveDepthResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) SetFileMaxRecursiveDepth(ctx context.Context, in *SetFileMaxRecursiveDepthRequest, opts ...dcerpc.CallOption) (*SetFileMaxRecursiveDepthResponse, error) {
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
	out := &SetFileMaxRecursiveDepthResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) GetFileMaxTotalSize(ctx context.Context, in *GetFileMaxTotalSizeRequest, opts ...dcerpc.CallOption) (*GetFileMaxTotalSizeResponse, error) {
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
	out := &GetFileMaxTotalSizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) SetFileMaxTotalSize(ctx context.Context, in *SetFileMaxTotalSizeRequest, opts ...dcerpc.CallOption) (*SetFileMaxTotalSizeResponse, error) {
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
	out := &SetFileMaxTotalSizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) GetFiles(ctx context.Context, in *GetFilesRequest, opts ...dcerpc.CallOption) (*GetFilesResponse, error) {
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
	out := &GetFilesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) SetFiles(ctx context.Context, in *SetFilesRequest, opts ...dcerpc.CallOption) (*SetFilesResponse, error) {
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
	out := &SetFilesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) GetManagementQueries(ctx context.Context, in *GetManagementQueriesRequest, opts ...dcerpc.CallOption) (*GetManagementQueriesResponse, error) {
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
	out := &GetManagementQueriesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) SetManagementQueries(ctx context.Context, in *SetManagementQueriesRequest, opts ...dcerpc.CallOption) (*SetManagementQueriesResponse, error) {
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
	out := &SetManagementQueriesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) GetQueryNetworkAdapters(ctx context.Context, in *GetQueryNetworkAdaptersRequest, opts ...dcerpc.CallOption) (*GetQueryNetworkAdaptersResponse, error) {
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
	out := &GetQueryNetworkAdaptersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) SetQueryNetworkAdapters(ctx context.Context, in *SetQueryNetworkAdaptersRequest, opts ...dcerpc.CallOption) (*SetQueryNetworkAdaptersResponse, error) {
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
	out := &SetQueryNetworkAdaptersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) GetRegistryKeys(ctx context.Context, in *GetRegistryKeysRequest, opts ...dcerpc.CallOption) (*GetRegistryKeysResponse, error) {
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
	out := &GetRegistryKeysResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) SetRegistryKeys(ctx context.Context, in *SetRegistryKeysRequest, opts ...dcerpc.CallOption) (*SetRegistryKeysResponse, error) {
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
	out := &SetRegistryKeysResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) GetRegistryMaxRecursiveDepth(ctx context.Context, in *GetRegistryMaxRecursiveDepthRequest, opts ...dcerpc.CallOption) (*GetRegistryMaxRecursiveDepthResponse, error) {
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
	out := &GetRegistryMaxRecursiveDepthResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) SetRegistryMaxRecursiveDepth(ctx context.Context, in *SetRegistryMaxRecursiveDepthRequest, opts ...dcerpc.CallOption) (*SetRegistryMaxRecursiveDepthResponse, error) {
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
	out := &SetRegistryMaxRecursiveDepthResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) GetSystemStateFile(ctx context.Context, in *GetSystemStateFileRequest, opts ...dcerpc.CallOption) (*GetSystemStateFileResponse, error) {
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
	out := &GetSystemStateFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) SetSystemStateFile(ctx context.Context, in *SetSystemStateFileRequest, opts ...dcerpc.CallOption) (*SetSystemStateFileResponse, error) {
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
	out := &SetSystemStateFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigurationDataCollectorClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultConfigurationDataCollectorClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultConfigurationDataCollectorClient) IPID(ctx context.Context, ipid *dcom.IPID) ConfigurationDataCollectorClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultConfigurationDataCollectorClient{
		DataCollectorClient: o.DataCollectorClient.IPID(ctx, ipid),
		cc:                  o.cc,
		ipid:                ipid,
	}
}

func NewConfigurationDataCollectorClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ConfigurationDataCollectorClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ConfigurationDataCollectorSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idatacollector.NewDataCollectorClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultConfigurationDataCollectorClient{
		DataCollectorClient: base,
		cc:                  cc,
		ipid:                ipid,
	}, nil
}

// xxx_GetFileMaxCountOperation structure represents the FileMaxCount operation
type xxx_GetFileMaxCountOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Count  uint32         `idl:"name:count" json:"count"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileMaxCountOperation) OpNum() int { return 32 }

func (o *xxx_GetFileMaxCountOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/FileMaxCount"
}

func (o *xxx_GetFileMaxCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileMaxCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFileMaxCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFileMaxCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileMaxCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// count {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
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

func (o *xxx_GetFileMaxCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// count {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
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

// GetFileMaxCountRequest structure represents the FileMaxCount operation request
type GetFileMaxCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFileMaxCountRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFileMaxCountOperation) *xxx_GetFileMaxCountOperation {
	if op == nil {
		op = &xxx_GetFileMaxCountOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetFileMaxCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileMaxCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFileMaxCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFileMaxCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileMaxCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileMaxCountResponse structure represents the FileMaxCount operation response
type GetFileMaxCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Count uint32         `idl:"name:count" json:"count"`
	// Return: The FileMaxCount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFileMaxCountResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFileMaxCountOperation) *xxx_GetFileMaxCountOperation {
	if op == nil {
		op = &xxx_GetFileMaxCountOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Count = o.Count
	op.Return = o.Return
	return op
}

func (o *GetFileMaxCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileMaxCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Count = op.Count
	o.Return = op.Return
}
func (o *GetFileMaxCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFileMaxCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileMaxCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFileMaxCountOperation structure represents the FileMaxCount operation
type xxx_SetFileMaxCountOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Count  uint32         `idl:"name:count" json:"count"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFileMaxCountOperation) OpNum() int { return 33 }

func (o *xxx_SetFileMaxCountOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/FileMaxCount"
}

func (o *xxx_SetFileMaxCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileMaxCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// count {in} (1:(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileMaxCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// count {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileMaxCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileMaxCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFileMaxCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFileMaxCountRequest structure represents the FileMaxCount operation request
type SetFileMaxCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Count uint32         `idl:"name:count" json:"count"`
}

func (o *SetFileMaxCountRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFileMaxCountOperation) *xxx_SetFileMaxCountOperation {
	if op == nil {
		op = &xxx_SetFileMaxCountOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Count = o.Count
	return op
}

func (o *SetFileMaxCountRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFileMaxCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Count = op.Count
}
func (o *SetFileMaxCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFileMaxCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFileMaxCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFileMaxCountResponse structure represents the FileMaxCount operation response
type SetFileMaxCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FileMaxCount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFileMaxCountResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFileMaxCountOperation) *xxx_SetFileMaxCountOperation {
	if op == nil {
		op = &xxx_SetFileMaxCountOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetFileMaxCountResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFileMaxCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFileMaxCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFileMaxCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFileMaxCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFileMaxRecursiveDepthOperation structure represents the FileMaxRecursiveDepth operation
type xxx_GetFileMaxRecursiveDepthOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Depth  uint32         `idl:"name:depth" json:"depth"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileMaxRecursiveDepthOperation) OpNum() int { return 34 }

func (o *xxx_GetFileMaxRecursiveDepthOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/FileMaxRecursiveDepth"
}

func (o *xxx_GetFileMaxRecursiveDepthOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileMaxRecursiveDepthOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFileMaxRecursiveDepthOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFileMaxRecursiveDepthOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileMaxRecursiveDepthOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// depth {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Depth); err != nil {
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

func (o *xxx_GetFileMaxRecursiveDepthOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// depth {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Depth); err != nil {
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

// GetFileMaxRecursiveDepthRequest structure represents the FileMaxRecursiveDepth operation request
type GetFileMaxRecursiveDepthRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFileMaxRecursiveDepthRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFileMaxRecursiveDepthOperation) *xxx_GetFileMaxRecursiveDepthOperation {
	if op == nil {
		op = &xxx_GetFileMaxRecursiveDepthOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetFileMaxRecursiveDepthRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileMaxRecursiveDepthOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFileMaxRecursiveDepthRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFileMaxRecursiveDepthRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileMaxRecursiveDepthOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileMaxRecursiveDepthResponse structure represents the FileMaxRecursiveDepth operation response
type GetFileMaxRecursiveDepthResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Depth uint32         `idl:"name:depth" json:"depth"`
	// Return: The FileMaxRecursiveDepth return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFileMaxRecursiveDepthResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFileMaxRecursiveDepthOperation) *xxx_GetFileMaxRecursiveDepthOperation {
	if op == nil {
		op = &xxx_GetFileMaxRecursiveDepthOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Depth = o.Depth
	op.Return = o.Return
	return op
}

func (o *GetFileMaxRecursiveDepthResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileMaxRecursiveDepthOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Depth = op.Depth
	o.Return = op.Return
}
func (o *GetFileMaxRecursiveDepthResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFileMaxRecursiveDepthResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileMaxRecursiveDepthOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFileMaxRecursiveDepthOperation structure represents the FileMaxRecursiveDepth operation
type xxx_SetFileMaxRecursiveDepthOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Depth  uint32         `idl:"name:depth" json:"depth"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFileMaxRecursiveDepthOperation) OpNum() int { return 35 }

func (o *xxx_SetFileMaxRecursiveDepthOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/FileMaxRecursiveDepth"
}

func (o *xxx_SetFileMaxRecursiveDepthOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileMaxRecursiveDepthOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// depth {in} (1:(uint32))
	{
		if err := w.WriteData(o.Depth); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileMaxRecursiveDepthOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// depth {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Depth); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileMaxRecursiveDepthOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileMaxRecursiveDepthOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFileMaxRecursiveDepthOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFileMaxRecursiveDepthRequest structure represents the FileMaxRecursiveDepth operation request
type SetFileMaxRecursiveDepthRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Depth uint32         `idl:"name:depth" json:"depth"`
}

func (o *SetFileMaxRecursiveDepthRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFileMaxRecursiveDepthOperation) *xxx_SetFileMaxRecursiveDepthOperation {
	if op == nil {
		op = &xxx_SetFileMaxRecursiveDepthOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Depth = o.Depth
	return op
}

func (o *SetFileMaxRecursiveDepthRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFileMaxRecursiveDepthOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Depth = op.Depth
}
func (o *SetFileMaxRecursiveDepthRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFileMaxRecursiveDepthRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFileMaxRecursiveDepthOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFileMaxRecursiveDepthResponse structure represents the FileMaxRecursiveDepth operation response
type SetFileMaxRecursiveDepthResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FileMaxRecursiveDepth return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFileMaxRecursiveDepthResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFileMaxRecursiveDepthOperation) *xxx_SetFileMaxRecursiveDepthOperation {
	if op == nil {
		op = &xxx_SetFileMaxRecursiveDepthOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetFileMaxRecursiveDepthResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFileMaxRecursiveDepthOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFileMaxRecursiveDepthResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFileMaxRecursiveDepthResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFileMaxRecursiveDepthOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFileMaxTotalSizeOperation structure represents the FileMaxTotalSize operation
type xxx_GetFileMaxTotalSizeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Size   uint32         `idl:"name:size" json:"size"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileMaxTotalSizeOperation) OpNum() int { return 36 }

func (o *xxx_GetFileMaxTotalSizeOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/FileMaxTotalSize"
}

func (o *xxx_GetFileMaxTotalSizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileMaxTotalSizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFileMaxTotalSizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFileMaxTotalSizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileMaxTotalSizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// size {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Size); err != nil {
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

func (o *xxx_GetFileMaxTotalSizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// size {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Size); err != nil {
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

// GetFileMaxTotalSizeRequest structure represents the FileMaxTotalSize operation request
type GetFileMaxTotalSizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFileMaxTotalSizeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFileMaxTotalSizeOperation) *xxx_GetFileMaxTotalSizeOperation {
	if op == nil {
		op = &xxx_GetFileMaxTotalSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetFileMaxTotalSizeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileMaxTotalSizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFileMaxTotalSizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFileMaxTotalSizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileMaxTotalSizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileMaxTotalSizeResponse structure represents the FileMaxTotalSize operation response
type GetFileMaxTotalSizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Size uint32         `idl:"name:size" json:"size"`
	// Return: The FileMaxTotalSize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFileMaxTotalSizeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFileMaxTotalSizeOperation) *xxx_GetFileMaxTotalSizeOperation {
	if op == nil {
		op = &xxx_GetFileMaxTotalSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Size = o.Size
	op.Return = o.Return
	return op
}

func (o *GetFileMaxTotalSizeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileMaxTotalSizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Size = op.Size
	o.Return = op.Return
}
func (o *GetFileMaxTotalSizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFileMaxTotalSizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileMaxTotalSizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFileMaxTotalSizeOperation structure represents the FileMaxTotalSize operation
type xxx_SetFileMaxTotalSizeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Size   uint32         `idl:"name:size" json:"size"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFileMaxTotalSizeOperation) OpNum() int { return 37 }

func (o *xxx_SetFileMaxTotalSizeOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/FileMaxTotalSize"
}

func (o *xxx_SetFileMaxTotalSizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileMaxTotalSizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// size {in} (1:(uint32))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileMaxTotalSizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// size {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileMaxTotalSizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileMaxTotalSizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFileMaxTotalSizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFileMaxTotalSizeRequest structure represents the FileMaxTotalSize operation request
type SetFileMaxTotalSizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Size uint32         `idl:"name:size" json:"size"`
}

func (o *SetFileMaxTotalSizeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFileMaxTotalSizeOperation) *xxx_SetFileMaxTotalSizeOperation {
	if op == nil {
		op = &xxx_SetFileMaxTotalSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Size = o.Size
	return op
}

func (o *SetFileMaxTotalSizeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFileMaxTotalSizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Size = op.Size
}
func (o *SetFileMaxTotalSizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFileMaxTotalSizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFileMaxTotalSizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFileMaxTotalSizeResponse structure represents the FileMaxTotalSize operation response
type SetFileMaxTotalSizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FileMaxTotalSize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFileMaxTotalSizeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFileMaxTotalSizeOperation) *xxx_SetFileMaxTotalSizeOperation {
	if op == nil {
		op = &xxx_SetFileMaxTotalSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetFileMaxTotalSizeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFileMaxTotalSizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFileMaxTotalSizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFileMaxTotalSizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFileMaxTotalSizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFilesOperation structure represents the Files operation
type xxx_GetFilesOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Files  *oaut.SafeArray `idl:"name:Files" json:"files"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFilesOperation) OpNum() int { return 38 }

func (o *xxx_GetFilesOperation) OpName() string { return "/IConfigurationDataCollector/v0/Files" }

func (o *xxx_GetFilesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFilesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFilesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFilesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFilesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Files {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Files != nil {
			_ptr_Files := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Files != nil {
					if err := o.Files.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Files, _ptr_Files); err != nil {
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

func (o *xxx_GetFilesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Files {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_Files := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Files == nil {
				o.Files = &oaut.SafeArray{}
			}
			if err := o.Files.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Files := func(ptr interface{}) { o.Files = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Files, _s_Files, _ptr_Files); err != nil {
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

// GetFilesRequest structure represents the Files operation request
type GetFilesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFilesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFilesOperation) *xxx_GetFilesOperation {
	if op == nil {
		op = &xxx_GetFilesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetFilesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFilesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFilesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFilesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFilesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFilesResponse structure represents the Files operation response
type GetFilesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Files *oaut.SafeArray `idl:"name:Files" json:"files"`
	// Return: The Files return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFilesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFilesOperation) *xxx_GetFilesOperation {
	if op == nil {
		op = &xxx_GetFilesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Files = o.Files
	op.Return = o.Return
	return op
}

func (o *GetFilesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFilesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Files = op.Files
	o.Return = op.Return
}
func (o *GetFilesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFilesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFilesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFilesOperation structure represents the Files operation
type xxx_SetFilesOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Files  *oaut.SafeArray `idl:"name:Files" json:"files"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFilesOperation) OpNum() int { return 39 }

func (o *xxx_SetFilesOperation) OpName() string { return "/IConfigurationDataCollector/v0/Files" }

func (o *xxx_SetFilesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFilesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Files {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Files != nil {
			_ptr_Files := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Files != nil {
					if err := o.Files.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Files, _ptr_Files); err != nil {
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

func (o *xxx_SetFilesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Files {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_Files := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Files == nil {
				o.Files = &oaut.SafeArray{}
			}
			if err := o.Files.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Files := func(ptr interface{}) { o.Files = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Files, _s_Files, _ptr_Files); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFilesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFilesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFilesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFilesRequest structure represents the Files operation request
type SetFilesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Files *oaut.SafeArray `idl:"name:Files" json:"files"`
}

func (o *SetFilesRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFilesOperation) *xxx_SetFilesOperation {
	if op == nil {
		op = &xxx_SetFilesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Files = o.Files
	return op
}

func (o *SetFilesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFilesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Files = op.Files
}
func (o *SetFilesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFilesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFilesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFilesResponse structure represents the Files operation response
type SetFilesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Files return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFilesResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFilesOperation) *xxx_SetFilesOperation {
	if op == nil {
		op = &xxx_SetFilesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetFilesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFilesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFilesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFilesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFilesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetManagementQueriesOperation structure represents the ManagementQueries operation
type xxx_GetManagementQueriesOperation struct {
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Queries *oaut.SafeArray `idl:"name:Queries" json:"queries"`
	Return  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetManagementQueriesOperation) OpNum() int { return 40 }

func (o *xxx_GetManagementQueriesOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/ManagementQueries"
}

func (o *xxx_GetManagementQueriesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetManagementQueriesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetManagementQueriesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetManagementQueriesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetManagementQueriesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Queries {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Queries != nil {
			_ptr_Queries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Queries != nil {
					if err := o.Queries.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Queries, _ptr_Queries); err != nil {
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

func (o *xxx_GetManagementQueriesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Queries {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_Queries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Queries == nil {
				o.Queries = &oaut.SafeArray{}
			}
			if err := o.Queries.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Queries := func(ptr interface{}) { o.Queries = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Queries, _s_Queries, _ptr_Queries); err != nil {
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

// GetManagementQueriesRequest structure represents the ManagementQueries operation request
type GetManagementQueriesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetManagementQueriesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetManagementQueriesOperation) *xxx_GetManagementQueriesOperation {
	if op == nil {
		op = &xxx_GetManagementQueriesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetManagementQueriesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetManagementQueriesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetManagementQueriesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetManagementQueriesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetManagementQueriesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetManagementQueriesResponse structure represents the ManagementQueries operation response
type GetManagementQueriesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Queries *oaut.SafeArray `idl:"name:Queries" json:"queries"`
	// Return: The ManagementQueries return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetManagementQueriesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetManagementQueriesOperation) *xxx_GetManagementQueriesOperation {
	if op == nil {
		op = &xxx_GetManagementQueriesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Queries = o.Queries
	op.Return = o.Return
	return op
}

func (o *GetManagementQueriesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetManagementQueriesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Queries = op.Queries
	o.Return = op.Return
}
func (o *GetManagementQueriesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetManagementQueriesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetManagementQueriesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetManagementQueriesOperation structure represents the ManagementQueries operation
type xxx_SetManagementQueriesOperation struct {
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Queries *oaut.SafeArray `idl:"name:Queries" json:"queries"`
	Return  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetManagementQueriesOperation) OpNum() int { return 41 }

func (o *xxx_SetManagementQueriesOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/ManagementQueries"
}

func (o *xxx_SetManagementQueriesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetManagementQueriesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Queries {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Queries != nil {
			_ptr_Queries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Queries != nil {
					if err := o.Queries.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Queries, _ptr_Queries); err != nil {
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

func (o *xxx_SetManagementQueriesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Queries {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_Queries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Queries == nil {
				o.Queries = &oaut.SafeArray{}
			}
			if err := o.Queries.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Queries := func(ptr interface{}) { o.Queries = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Queries, _s_Queries, _ptr_Queries); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetManagementQueriesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetManagementQueriesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetManagementQueriesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetManagementQueriesRequest structure represents the ManagementQueries operation request
type SetManagementQueriesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Queries *oaut.SafeArray `idl:"name:Queries" json:"queries"`
}

func (o *SetManagementQueriesRequest) xxx_ToOp(ctx context.Context, op *xxx_SetManagementQueriesOperation) *xxx_SetManagementQueriesOperation {
	if op == nil {
		op = &xxx_SetManagementQueriesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Queries = o.Queries
	return op
}

func (o *SetManagementQueriesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetManagementQueriesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Queries = op.Queries
}
func (o *SetManagementQueriesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetManagementQueriesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetManagementQueriesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetManagementQueriesResponse structure represents the ManagementQueries operation response
type SetManagementQueriesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ManagementQueries return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetManagementQueriesResponse) xxx_ToOp(ctx context.Context, op *xxx_SetManagementQueriesOperation) *xxx_SetManagementQueriesOperation {
	if op == nil {
		op = &xxx_SetManagementQueriesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetManagementQueriesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetManagementQueriesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetManagementQueriesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetManagementQueriesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetManagementQueriesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetQueryNetworkAdaptersOperation structure represents the QueryNetworkAdapters operation
type xxx_GetQueryNetworkAdaptersOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Network int16          `idl:"name:network" json:"network"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQueryNetworkAdaptersOperation) OpNum() int { return 42 }

func (o *xxx_GetQueryNetworkAdaptersOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/QueryNetworkAdapters"
}

func (o *xxx_GetQueryNetworkAdaptersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueryNetworkAdaptersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetQueryNetworkAdaptersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetQueryNetworkAdaptersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueryNetworkAdaptersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// network {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Network); err != nil {
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

func (o *xxx_GetQueryNetworkAdaptersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// network {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Network); err != nil {
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

// GetQueryNetworkAdaptersRequest structure represents the QueryNetworkAdapters operation request
type GetQueryNetworkAdaptersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetQueryNetworkAdaptersRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQueryNetworkAdaptersOperation) *xxx_GetQueryNetworkAdaptersOperation {
	if op == nil {
		op = &xxx_GetQueryNetworkAdaptersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetQueryNetworkAdaptersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQueryNetworkAdaptersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetQueryNetworkAdaptersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQueryNetworkAdaptersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQueryNetworkAdaptersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQueryNetworkAdaptersResponse structure represents the QueryNetworkAdapters operation response
type GetQueryNetworkAdaptersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Network int16          `idl:"name:network" json:"network"`
	// Return: The QueryNetworkAdapters return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetQueryNetworkAdaptersResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQueryNetworkAdaptersOperation) *xxx_GetQueryNetworkAdaptersOperation {
	if op == nil {
		op = &xxx_GetQueryNetworkAdaptersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Network = o.Network
	op.Return = o.Return
	return op
}

func (o *GetQueryNetworkAdaptersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQueryNetworkAdaptersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Network = op.Network
	o.Return = op.Return
}
func (o *GetQueryNetworkAdaptersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQueryNetworkAdaptersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQueryNetworkAdaptersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetQueryNetworkAdaptersOperation structure represents the QueryNetworkAdapters operation
type xxx_SetQueryNetworkAdaptersOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Network int16          `idl:"name:network" json:"network"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetQueryNetworkAdaptersOperation) OpNum() int { return 43 }

func (o *xxx_SetQueryNetworkAdaptersOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/QueryNetworkAdapters"
}

func (o *xxx_SetQueryNetworkAdaptersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQueryNetworkAdaptersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// network {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Network); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQueryNetworkAdaptersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// network {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Network); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQueryNetworkAdaptersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQueryNetworkAdaptersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetQueryNetworkAdaptersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetQueryNetworkAdaptersRequest structure represents the QueryNetworkAdapters operation request
type SetQueryNetworkAdaptersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Network int16          `idl:"name:network" json:"network"`
}

func (o *SetQueryNetworkAdaptersRequest) xxx_ToOp(ctx context.Context, op *xxx_SetQueryNetworkAdaptersOperation) *xxx_SetQueryNetworkAdaptersOperation {
	if op == nil {
		op = &xxx_SetQueryNetworkAdaptersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Network = o.Network
	return op
}

func (o *SetQueryNetworkAdaptersRequest) xxx_FromOp(ctx context.Context, op *xxx_SetQueryNetworkAdaptersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Network = op.Network
}
func (o *SetQueryNetworkAdaptersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetQueryNetworkAdaptersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetQueryNetworkAdaptersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetQueryNetworkAdaptersResponse structure represents the QueryNetworkAdapters operation response
type SetQueryNetworkAdaptersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The QueryNetworkAdapters return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetQueryNetworkAdaptersResponse) xxx_ToOp(ctx context.Context, op *xxx_SetQueryNetworkAdaptersOperation) *xxx_SetQueryNetworkAdaptersOperation {
	if op == nil {
		op = &xxx_SetQueryNetworkAdaptersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetQueryNetworkAdaptersResponse) xxx_FromOp(ctx context.Context, op *xxx_SetQueryNetworkAdaptersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetQueryNetworkAdaptersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetQueryNetworkAdaptersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetQueryNetworkAdaptersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRegistryKeysOperation structure represents the RegistryKeys operation
type xxx_GetRegistryKeysOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Query  *oaut.SafeArray `idl:"name:query" json:"query"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRegistryKeysOperation) OpNum() int { return 44 }

func (o *xxx_GetRegistryKeysOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/RegistryKeys"
}

func (o *xxx_GetRegistryKeysOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRegistryKeysOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRegistryKeysOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRegistryKeysOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRegistryKeysOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// query {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Query != nil {
			_ptr_query := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Query != nil {
					if err := o.Query.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Query, _ptr_query); err != nil {
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

func (o *xxx_GetRegistryKeysOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// query {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_query := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Query == nil {
				o.Query = &oaut.SafeArray{}
			}
			if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_query := func(ptr interface{}) { o.Query = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Query, _s_query, _ptr_query); err != nil {
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

// GetRegistryKeysRequest structure represents the RegistryKeys operation request
type GetRegistryKeysRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRegistryKeysRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRegistryKeysOperation) *xxx_GetRegistryKeysOperation {
	if op == nil {
		op = &xxx_GetRegistryKeysOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetRegistryKeysRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRegistryKeysOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRegistryKeysRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRegistryKeysRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRegistryKeysOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRegistryKeysResponse structure represents the RegistryKeys operation response
type GetRegistryKeysResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Query *oaut.SafeArray `idl:"name:query" json:"query"`
	// Return: The RegistryKeys return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRegistryKeysResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRegistryKeysOperation) *xxx_GetRegistryKeysOperation {
	if op == nil {
		op = &xxx_GetRegistryKeysOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Query = o.Query
	op.Return = o.Return
	return op
}

func (o *GetRegistryKeysResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRegistryKeysOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Query = op.Query
	o.Return = op.Return
}
func (o *GetRegistryKeysResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRegistryKeysResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRegistryKeysOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetRegistryKeysOperation structure represents the RegistryKeys operation
type xxx_SetRegistryKeysOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Query  *oaut.SafeArray `idl:"name:query" json:"query"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetRegistryKeysOperation) OpNum() int { return 45 }

func (o *xxx_SetRegistryKeysOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/RegistryKeys"
}

func (o *xxx_SetRegistryKeysOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRegistryKeysOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// query {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Query != nil {
			_ptr_query := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Query != nil {
					if err := o.Query.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Query, _ptr_query); err != nil {
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

func (o *xxx_SetRegistryKeysOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// query {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_query := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Query == nil {
				o.Query = &oaut.SafeArray{}
			}
			if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_query := func(ptr interface{}) { o.Query = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Query, _s_query, _ptr_query); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRegistryKeysOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRegistryKeysOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetRegistryKeysOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetRegistryKeysRequest structure represents the RegistryKeys operation request
type SetRegistryKeysRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Query *oaut.SafeArray `idl:"name:query" json:"query"`
}

func (o *SetRegistryKeysRequest) xxx_ToOp(ctx context.Context, op *xxx_SetRegistryKeysOperation) *xxx_SetRegistryKeysOperation {
	if op == nil {
		op = &xxx_SetRegistryKeysOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Query = o.Query
	return op
}

func (o *SetRegistryKeysRequest) xxx_FromOp(ctx context.Context, op *xxx_SetRegistryKeysOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Query = op.Query
}
func (o *SetRegistryKeysRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetRegistryKeysRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRegistryKeysOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetRegistryKeysResponse structure represents the RegistryKeys operation response
type SetRegistryKeysResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RegistryKeys return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetRegistryKeysResponse) xxx_ToOp(ctx context.Context, op *xxx_SetRegistryKeysOperation) *xxx_SetRegistryKeysOperation {
	if op == nil {
		op = &xxx_SetRegistryKeysOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetRegistryKeysResponse) xxx_FromOp(ctx context.Context, op *xxx_SetRegistryKeysOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetRegistryKeysResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetRegistryKeysResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRegistryKeysOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRegistryMaxRecursiveDepthOperation structure represents the RegistryMaxRecursiveDepth operation
type xxx_GetRegistryMaxRecursiveDepthOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Depth  uint32         `idl:"name:depth" json:"depth"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRegistryMaxRecursiveDepthOperation) OpNum() int { return 46 }

func (o *xxx_GetRegistryMaxRecursiveDepthOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/RegistryMaxRecursiveDepth"
}

func (o *xxx_GetRegistryMaxRecursiveDepthOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRegistryMaxRecursiveDepthOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRegistryMaxRecursiveDepthOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRegistryMaxRecursiveDepthOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRegistryMaxRecursiveDepthOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// depth {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Depth); err != nil {
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

func (o *xxx_GetRegistryMaxRecursiveDepthOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// depth {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Depth); err != nil {
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

// GetRegistryMaxRecursiveDepthRequest structure represents the RegistryMaxRecursiveDepth operation request
type GetRegistryMaxRecursiveDepthRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRegistryMaxRecursiveDepthRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRegistryMaxRecursiveDepthOperation) *xxx_GetRegistryMaxRecursiveDepthOperation {
	if op == nil {
		op = &xxx_GetRegistryMaxRecursiveDepthOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetRegistryMaxRecursiveDepthRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRegistryMaxRecursiveDepthOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRegistryMaxRecursiveDepthRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRegistryMaxRecursiveDepthRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRegistryMaxRecursiveDepthOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRegistryMaxRecursiveDepthResponse structure represents the RegistryMaxRecursiveDepth operation response
type GetRegistryMaxRecursiveDepthResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Depth uint32         `idl:"name:depth" json:"depth"`
	// Return: The RegistryMaxRecursiveDepth return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRegistryMaxRecursiveDepthResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRegistryMaxRecursiveDepthOperation) *xxx_GetRegistryMaxRecursiveDepthOperation {
	if op == nil {
		op = &xxx_GetRegistryMaxRecursiveDepthOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Depth = o.Depth
	op.Return = o.Return
	return op
}

func (o *GetRegistryMaxRecursiveDepthResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRegistryMaxRecursiveDepthOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Depth = op.Depth
	o.Return = op.Return
}
func (o *GetRegistryMaxRecursiveDepthResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRegistryMaxRecursiveDepthResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRegistryMaxRecursiveDepthOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetRegistryMaxRecursiveDepthOperation structure represents the RegistryMaxRecursiveDepth operation
type xxx_SetRegistryMaxRecursiveDepthOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Depth  uint32         `idl:"name:depth" json:"depth"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetRegistryMaxRecursiveDepthOperation) OpNum() int { return 47 }

func (o *xxx_SetRegistryMaxRecursiveDepthOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/RegistryMaxRecursiveDepth"
}

func (o *xxx_SetRegistryMaxRecursiveDepthOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRegistryMaxRecursiveDepthOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// depth {in} (1:(uint32))
	{
		if err := w.WriteData(o.Depth); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRegistryMaxRecursiveDepthOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// depth {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Depth); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRegistryMaxRecursiveDepthOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRegistryMaxRecursiveDepthOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetRegistryMaxRecursiveDepthOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetRegistryMaxRecursiveDepthRequest structure represents the RegistryMaxRecursiveDepth operation request
type SetRegistryMaxRecursiveDepthRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Depth uint32         `idl:"name:depth" json:"depth"`
}

func (o *SetRegistryMaxRecursiveDepthRequest) xxx_ToOp(ctx context.Context, op *xxx_SetRegistryMaxRecursiveDepthOperation) *xxx_SetRegistryMaxRecursiveDepthOperation {
	if op == nil {
		op = &xxx_SetRegistryMaxRecursiveDepthOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Depth = o.Depth
	return op
}

func (o *SetRegistryMaxRecursiveDepthRequest) xxx_FromOp(ctx context.Context, op *xxx_SetRegistryMaxRecursiveDepthOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Depth = op.Depth
}
func (o *SetRegistryMaxRecursiveDepthRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetRegistryMaxRecursiveDepthRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRegistryMaxRecursiveDepthOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetRegistryMaxRecursiveDepthResponse structure represents the RegistryMaxRecursiveDepth operation response
type SetRegistryMaxRecursiveDepthResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RegistryMaxRecursiveDepth return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetRegistryMaxRecursiveDepthResponse) xxx_ToOp(ctx context.Context, op *xxx_SetRegistryMaxRecursiveDepthOperation) *xxx_SetRegistryMaxRecursiveDepthOperation {
	if op == nil {
		op = &xxx_SetRegistryMaxRecursiveDepthOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetRegistryMaxRecursiveDepthResponse) xxx_FromOp(ctx context.Context, op *xxx_SetRegistryMaxRecursiveDepthOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetRegistryMaxRecursiveDepthResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetRegistryMaxRecursiveDepthResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRegistryMaxRecursiveDepthOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSystemStateFileOperation structure represents the SystemStateFile operation
type xxx_GetSystemStateFileOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileName *oaut.String   `idl:"name:FileName" json:"file_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSystemStateFileOperation) OpNum() int { return 48 }

func (o *xxx_GetSystemStateFileOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/SystemStateFile"
}

func (o *xxx_GetSystemStateFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemStateFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSystemStateFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSystemStateFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemStateFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// FileName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FileName != nil {
			_ptr_FileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileName != nil {
					if err := o.FileName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileName, _ptr_FileName); err != nil {
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

func (o *xxx_GetSystemStateFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// FileName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_FileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileName == nil {
				o.FileName = &oaut.String{}
			}
			if err := o.FileName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_FileName := func(ptr interface{}) { o.FileName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FileName, _s_FileName, _ptr_FileName); err != nil {
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

// GetSystemStateFileRequest structure represents the SystemStateFile operation request
type GetSystemStateFileRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSystemStateFileRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSystemStateFileOperation) *xxx_GetSystemStateFileOperation {
	if op == nil {
		op = &xxx_GetSystemStateFileOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSystemStateFileRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSystemStateFileOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSystemStateFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSystemStateFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSystemStateFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSystemStateFileResponse structure represents the SystemStateFile operation response
type GetSystemStateFileResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileName *oaut.String   `idl:"name:FileName" json:"file_name"`
	// Return: The SystemStateFile return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSystemStateFileResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSystemStateFileOperation) *xxx_GetSystemStateFileOperation {
	if op == nil {
		op = &xxx_GetSystemStateFileOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FileName = o.FileName
	op.Return = o.Return
	return op
}

func (o *GetSystemStateFileResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSystemStateFileOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileName = op.FileName
	o.Return = op.Return
}
func (o *GetSystemStateFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSystemStateFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSystemStateFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSystemStateFileOperation structure represents the SystemStateFile operation
type xxx_SetSystemStateFileOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileName *oaut.String   `idl:"name:FileName" json:"file_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSystemStateFileOperation) OpNum() int { return 49 }

func (o *xxx_SetSystemStateFileOperation) OpName() string {
	return "/IConfigurationDataCollector/v0/SystemStateFile"
}

func (o *xxx_SetSystemStateFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSystemStateFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// FileName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FileName != nil {
			_ptr_FileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileName != nil {
					if err := o.FileName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileName, _ptr_FileName); err != nil {
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

func (o *xxx_SetSystemStateFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// FileName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_FileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileName == nil {
				o.FileName = &oaut.String{}
			}
			if err := o.FileName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_FileName := func(ptr interface{}) { o.FileName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FileName, _s_FileName, _ptr_FileName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSystemStateFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSystemStateFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSystemStateFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSystemStateFileRequest structure represents the SystemStateFile operation request
type SetSystemStateFileRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	FileName *oaut.String   `idl:"name:FileName" json:"file_name"`
}

func (o *SetSystemStateFileRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSystemStateFileOperation) *xxx_SetSystemStateFileOperation {
	if op == nil {
		op = &xxx_SetSystemStateFileOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.FileName = o.FileName
	return op
}

func (o *SetSystemStateFileRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSystemStateFileOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FileName = op.FileName
}
func (o *SetSystemStateFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSystemStateFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSystemStateFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSystemStateFileResponse structure represents the SystemStateFile operation response
type SetSystemStateFileResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SystemStateFile return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSystemStateFileResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSystemStateFileOperation) *xxx_SetSystemStateFileOperation {
	if op == nil {
		op = &xxx_SetSystemStateFileOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSystemStateFileResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSystemStateFileOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSystemStateFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSystemStateFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSystemStateFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
