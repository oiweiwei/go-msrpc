package idatamanager

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
	pla "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla"
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
	_ = pla.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/pla"
)

var (
	// IDataManager interface identifier 03837541-098b-11d8-9414-505054503030
	DataManagerIID = &dcom.IID{Data1: 0x03837541, Data2: 0x098b, Data3: 0x11d8, Data4: []byte{0x94, 0x14, 0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax UUID
	DataManagerSyntaxUUID = &uuid.UUID{TimeLow: 0x3837541, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax ID
	DataManagerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DataManagerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IDataManager interface.
type DataManagerClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// Enabled operation.
	GetEnabled(context.Context, *GetEnabledRequest, ...dcerpc.CallOption) (*GetEnabledResponse, error)

	// Enabled operation.
	SetEnabled(context.Context, *SetEnabledRequest, ...dcerpc.CallOption) (*SetEnabledResponse, error)

	// CheckBeforeRunning operation.
	GetCheckBeforeRunning(context.Context, *GetCheckBeforeRunningRequest, ...dcerpc.CallOption) (*GetCheckBeforeRunningResponse, error)

	// CheckBeforeRunning operation.
	SetCheckBeforeRunning(context.Context, *SetCheckBeforeRunningRequest, ...dcerpc.CallOption) (*SetCheckBeforeRunningResponse, error)

	// MinFreeDisk operation.
	GetMinFreeDisk(context.Context, *GetMinFreeDiskRequest, ...dcerpc.CallOption) (*GetMinFreeDiskResponse, error)

	// MinFreeDisk operation.
	SetMinFreeDisk(context.Context, *SetMinFreeDiskRequest, ...dcerpc.CallOption) (*SetMinFreeDiskResponse, error)

	// MaxSize operation.
	GetMaxSize(context.Context, *GetMaxSizeRequest, ...dcerpc.CallOption) (*GetMaxSizeResponse, error)

	// MaxSize operation.
	SetMaxSize(context.Context, *SetMaxSizeRequest, ...dcerpc.CallOption) (*SetMaxSizeResponse, error)

	// MaxFolderCount operation.
	GetMaxFolderCount(context.Context, *GetMaxFolderCountRequest, ...dcerpc.CallOption) (*GetMaxFolderCountResponse, error)

	// MaxFolderCount operation.
	SetMaxFolderCount(context.Context, *SetMaxFolderCountRequest, ...dcerpc.CallOption) (*SetMaxFolderCountResponse, error)

	// The ResourcePolicy enumeration defines the order in which folders are deleted when
	// one of the disk resource limits is exceeded.
	GetResourcePolicy(context.Context, *GetResourcePolicyRequest, ...dcerpc.CallOption) (*GetResourcePolicyResponse, error)

	// The ResourcePolicy enumeration defines the order in which folders are deleted when
	// one of the disk resource limits is exceeded.
	SetResourcePolicy(context.Context, *SetResourcePolicyRequest, ...dcerpc.CallOption) (*SetResourcePolicyResponse, error)

	// The FolderActions (Get) method retrieves the FolderAction property, as specified
	// in the property table in section 3.2.4.2.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetFolderActions(context.Context, *GetFolderActionsRequest, ...dcerpc.CallOption) (*GetFolderActionsResponse, error)

	// ReportSchema operation.
	GetReportSchema(context.Context, *GetReportSchemaRequest, ...dcerpc.CallOption) (*GetReportSchemaResponse, error)

	// ReportSchema operation.
	SetReportSchema(context.Context, *SetReportSchemaRequest, ...dcerpc.CallOption) (*SetReportSchemaResponse, error)

	// ReportFileName operation.
	GetReportFileName(context.Context, *GetReportFileNameRequest, ...dcerpc.CallOption) (*GetReportFileNameResponse, error)

	// ReportFileName operation.
	SetReportFileName(context.Context, *SetReportFileNameRequest, ...dcerpc.CallOption) (*SetReportFileNameResponse, error)

	// RuleTargetFileName operation.
	GetRuleTargetFileName(context.Context, *GetRuleTargetFileNameRequest, ...dcerpc.CallOption) (*GetRuleTargetFileNameResponse, error)

	// RuleTargetFileName operation.
	SetRuleTargetFileName(context.Context, *SetRuleTargetFileNameRequest, ...dcerpc.CallOption) (*SetRuleTargetFileNameResponse, error)

	// EventsFileName operation.
	GetEventsFileName(context.Context, *GetEventsFileNameRequest, ...dcerpc.CallOption) (*GetEventsFileNameResponse, error)

	// EventsFileName operation.
	SetEventsFileName(context.Context, *SetEventsFileNameRequest, ...dcerpc.CallOption) (*SetEventsFileNameResponse, error)

	// Rules operation.
	GetRules(context.Context, *GetRulesRequest, ...dcerpc.CallOption) (*GetRulesResponse, error)

	// Rules operation.
	SetRules(context.Context, *SetRulesRequest, ...dcerpc.CallOption) (*SetRulesResponse, error)

	// The Run method is used to manually run the data manager. When the data manager is
	// run, the actions specified in the Steps parameter are executed on the data stored
	// in the folder specified by the bstrFolder. Actions taken can include creating an
	// XML report from binary performance files (.blg) or event trace files (.etl), running
	// XPath expressions against the report, transforming the report to HTML, cabbing the
	// report and sending it to a remote server, and deleting files in the directory specified
	// by bstrFolder.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	Run(context.Context, *RunRequest, ...dcerpc.CallOption) (*RunResponse, error)

	// The Extract method extracts the specified CAB file.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	Extract(context.Context, *ExtractRequest, ...dcerpc.CallOption) (*ExtractResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) DataManagerClient
}

type xxx_DefaultDataManagerClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDataManagerClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultDataManagerClient) GetEnabled(ctx context.Context, in *GetEnabledRequest, opts ...dcerpc.CallOption) (*GetEnabledResponse, error) {
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
	out := &GetEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) SetEnabled(ctx context.Context, in *SetEnabledRequest, opts ...dcerpc.CallOption) (*SetEnabledResponse, error) {
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
	out := &SetEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) GetCheckBeforeRunning(ctx context.Context, in *GetCheckBeforeRunningRequest, opts ...dcerpc.CallOption) (*GetCheckBeforeRunningResponse, error) {
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
	out := &GetCheckBeforeRunningResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) SetCheckBeforeRunning(ctx context.Context, in *SetCheckBeforeRunningRequest, opts ...dcerpc.CallOption) (*SetCheckBeforeRunningResponse, error) {
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
	out := &SetCheckBeforeRunningResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) GetMinFreeDisk(ctx context.Context, in *GetMinFreeDiskRequest, opts ...dcerpc.CallOption) (*GetMinFreeDiskResponse, error) {
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
	out := &GetMinFreeDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) SetMinFreeDisk(ctx context.Context, in *SetMinFreeDiskRequest, opts ...dcerpc.CallOption) (*SetMinFreeDiskResponse, error) {
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
	out := &SetMinFreeDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) GetMaxSize(ctx context.Context, in *GetMaxSizeRequest, opts ...dcerpc.CallOption) (*GetMaxSizeResponse, error) {
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
	out := &GetMaxSizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) SetMaxSize(ctx context.Context, in *SetMaxSizeRequest, opts ...dcerpc.CallOption) (*SetMaxSizeResponse, error) {
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
	out := &SetMaxSizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) GetMaxFolderCount(ctx context.Context, in *GetMaxFolderCountRequest, opts ...dcerpc.CallOption) (*GetMaxFolderCountResponse, error) {
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
	out := &GetMaxFolderCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) SetMaxFolderCount(ctx context.Context, in *SetMaxFolderCountRequest, opts ...dcerpc.CallOption) (*SetMaxFolderCountResponse, error) {
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
	out := &SetMaxFolderCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) GetResourcePolicy(ctx context.Context, in *GetResourcePolicyRequest, opts ...dcerpc.CallOption) (*GetResourcePolicyResponse, error) {
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
	out := &GetResourcePolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) SetResourcePolicy(ctx context.Context, in *SetResourcePolicyRequest, opts ...dcerpc.CallOption) (*SetResourcePolicyResponse, error) {
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
	out := &SetResourcePolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) GetFolderActions(ctx context.Context, in *GetFolderActionsRequest, opts ...dcerpc.CallOption) (*GetFolderActionsResponse, error) {
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
	out := &GetFolderActionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) GetReportSchema(ctx context.Context, in *GetReportSchemaRequest, opts ...dcerpc.CallOption) (*GetReportSchemaResponse, error) {
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
	out := &GetReportSchemaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) SetReportSchema(ctx context.Context, in *SetReportSchemaRequest, opts ...dcerpc.CallOption) (*SetReportSchemaResponse, error) {
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
	out := &SetReportSchemaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) GetReportFileName(ctx context.Context, in *GetReportFileNameRequest, opts ...dcerpc.CallOption) (*GetReportFileNameResponse, error) {
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
	out := &GetReportFileNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) SetReportFileName(ctx context.Context, in *SetReportFileNameRequest, opts ...dcerpc.CallOption) (*SetReportFileNameResponse, error) {
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
	out := &SetReportFileNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) GetRuleTargetFileName(ctx context.Context, in *GetRuleTargetFileNameRequest, opts ...dcerpc.CallOption) (*GetRuleTargetFileNameResponse, error) {
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
	out := &GetRuleTargetFileNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) SetRuleTargetFileName(ctx context.Context, in *SetRuleTargetFileNameRequest, opts ...dcerpc.CallOption) (*SetRuleTargetFileNameResponse, error) {
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
	out := &SetRuleTargetFileNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) GetEventsFileName(ctx context.Context, in *GetEventsFileNameRequest, opts ...dcerpc.CallOption) (*GetEventsFileNameResponse, error) {
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
	out := &GetEventsFileNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) SetEventsFileName(ctx context.Context, in *SetEventsFileNameRequest, opts ...dcerpc.CallOption) (*SetEventsFileNameResponse, error) {
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
	out := &SetEventsFileNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) GetRules(ctx context.Context, in *GetRulesRequest, opts ...dcerpc.CallOption) (*GetRulesResponse, error) {
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
	out := &GetRulesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) SetRules(ctx context.Context, in *SetRulesRequest, opts ...dcerpc.CallOption) (*SetRulesResponse, error) {
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
	out := &SetRulesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) Run(ctx context.Context, in *RunRequest, opts ...dcerpc.CallOption) (*RunResponse, error) {
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
	out := &RunResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) Extract(ctx context.Context, in *ExtractRequest, opts ...dcerpc.CallOption) (*ExtractResponse, error) {
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
	out := &ExtractResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDataManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) DataManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDataManagerClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}
func NewDataManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DataManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DataManagerSyntaxV0_0))...)
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
	return &xxx_DefaultDataManagerClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetEnabledOperation structure represents the Enabled operation
type xxx_GetEnabledOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled int16          `idl:"name:pfEnabled" json:"enabled"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEnabledOperation) OpNum() int { return 7 }

func (o *xxx_GetEnabledOperation) OpName() string { return "/IDataManager/v0/Enabled" }

func (o *xxx_GetEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfEnabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Enabled); err != nil {
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

func (o *xxx_GetEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfEnabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Enabled); err != nil {
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

// GetEnabledRequest structure represents the Enabled operation request
type GetEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEnabledRequest) xxx_ToOp(ctx context.Context) *xxx_GetEnabledOperation {
	if o == nil {
		return &xxx_GetEnabledOperation{}
	}
	return &xxx_GetEnabledOperation{
		This: o.This,
	}
}

func (o *GetEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEnabledResponse structure represents the Enabled operation response
type GetEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled int16          `idl:"name:pfEnabled" json:"enabled"`
	// Return: The Enabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEnabledResponse) xxx_ToOp(ctx context.Context) *xxx_GetEnabledOperation {
	if o == nil {
		return &xxx_GetEnabledOperation{}
	}
	return &xxx_GetEnabledOperation{
		That:    o.That,
		Enabled: o.Enabled,
		Return:  o.Return,
	}
}

func (o *GetEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enabled = op.Enabled
	o.Return = op.Return
}
func (o *GetEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEnabledOperation structure represents the Enabled operation
type xxx_SetEnabledOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled int16          `idl:"name:fEnabled" json:"enabled"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEnabledOperation) OpNum() int { return 8 }

func (o *xxx_SetEnabledOperation) OpName() string { return "/IDataManager/v0/Enabled" }

func (o *xxx_SetEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fEnabled {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Enabled); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fEnabled {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Enabled); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetEnabledRequest structure represents the Enabled operation request
type SetEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Enabled int16          `idl:"name:fEnabled" json:"enabled"`
}

func (o *SetEnabledRequest) xxx_ToOp(ctx context.Context) *xxx_SetEnabledOperation {
	if o == nil {
		return &xxx_SetEnabledOperation{}
	}
	return &xxx_SetEnabledOperation{
		This:    o.This,
		Enabled: o.Enabled,
	}
}

func (o *SetEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Enabled = op.Enabled
}
func (o *SetEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEnabledResponse structure represents the Enabled operation response
type SetEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Enabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEnabledResponse) xxx_ToOp(ctx context.Context) *xxx_SetEnabledOperation {
	if o == nil {
		return &xxx_SetEnabledOperation{}
	}
	return &xxx_SetEnabledOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCheckBeforeRunningOperation structure represents the CheckBeforeRunning operation
type xxx_GetCheckBeforeRunningOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Check  int16          `idl:"name:pfCheck" json:"check"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCheckBeforeRunningOperation) OpNum() int { return 9 }

func (o *xxx_GetCheckBeforeRunningOperation) OpName() string {
	return "/IDataManager/v0/CheckBeforeRunning"
}

func (o *xxx_GetCheckBeforeRunningOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCheckBeforeRunningOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCheckBeforeRunningOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCheckBeforeRunningOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCheckBeforeRunningOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfCheck {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Check); err != nil {
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

func (o *xxx_GetCheckBeforeRunningOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfCheck {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Check); err != nil {
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

// GetCheckBeforeRunningRequest structure represents the CheckBeforeRunning operation request
type GetCheckBeforeRunningRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCheckBeforeRunningRequest) xxx_ToOp(ctx context.Context) *xxx_GetCheckBeforeRunningOperation {
	if o == nil {
		return &xxx_GetCheckBeforeRunningOperation{}
	}
	return &xxx_GetCheckBeforeRunningOperation{
		This: o.This,
	}
}

func (o *GetCheckBeforeRunningRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCheckBeforeRunningOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCheckBeforeRunningRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCheckBeforeRunningRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCheckBeforeRunningOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCheckBeforeRunningResponse structure represents the CheckBeforeRunning operation response
type GetCheckBeforeRunningResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Check int16          `idl:"name:pfCheck" json:"check"`
	// Return: The CheckBeforeRunning return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCheckBeforeRunningResponse) xxx_ToOp(ctx context.Context) *xxx_GetCheckBeforeRunningOperation {
	if o == nil {
		return &xxx_GetCheckBeforeRunningOperation{}
	}
	return &xxx_GetCheckBeforeRunningOperation{
		That:   o.That,
		Check:  o.Check,
		Return: o.Return,
	}
}

func (o *GetCheckBeforeRunningResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCheckBeforeRunningOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Check = op.Check
	o.Return = op.Return
}
func (o *GetCheckBeforeRunningResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCheckBeforeRunningResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCheckBeforeRunningOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCheckBeforeRunningOperation structure represents the CheckBeforeRunning operation
type xxx_SetCheckBeforeRunningOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Check  int16          `idl:"name:fCheck" json:"check"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCheckBeforeRunningOperation) OpNum() int { return 10 }

func (o *xxx_SetCheckBeforeRunningOperation) OpName() string {
	return "/IDataManager/v0/CheckBeforeRunning"
}

func (o *xxx_SetCheckBeforeRunningOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCheckBeforeRunningOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fCheck {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Check); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCheckBeforeRunningOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fCheck {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Check); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCheckBeforeRunningOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCheckBeforeRunningOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCheckBeforeRunningOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetCheckBeforeRunningRequest structure represents the CheckBeforeRunning operation request
type SetCheckBeforeRunningRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Check int16          `idl:"name:fCheck" json:"check"`
}

func (o *SetCheckBeforeRunningRequest) xxx_ToOp(ctx context.Context) *xxx_SetCheckBeforeRunningOperation {
	if o == nil {
		return &xxx_SetCheckBeforeRunningOperation{}
	}
	return &xxx_SetCheckBeforeRunningOperation{
		This:  o.This,
		Check: o.Check,
	}
}

func (o *SetCheckBeforeRunningRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCheckBeforeRunningOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Check = op.Check
}
func (o *SetCheckBeforeRunningRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetCheckBeforeRunningRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCheckBeforeRunningOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCheckBeforeRunningResponse structure represents the CheckBeforeRunning operation response
type SetCheckBeforeRunningResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CheckBeforeRunning return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCheckBeforeRunningResponse) xxx_ToOp(ctx context.Context) *xxx_SetCheckBeforeRunningOperation {
	if o == nil {
		return &xxx_SetCheckBeforeRunningOperation{}
	}
	return &xxx_SetCheckBeforeRunningOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetCheckBeforeRunningResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCheckBeforeRunningOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCheckBeforeRunningResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetCheckBeforeRunningResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCheckBeforeRunningOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMinFreeDiskOperation structure represents the MinFreeDisk operation
type xxx_GetMinFreeDiskOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MinFreeDisk uint32         `idl:"name:MinFreeDisk" json:"min_free_disk"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMinFreeDiskOperation) OpNum() int { return 11 }

func (o *xxx_GetMinFreeDiskOperation) OpName() string { return "/IDataManager/v0/MinFreeDisk" }

func (o *xxx_GetMinFreeDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMinFreeDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMinFreeDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMinFreeDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMinFreeDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// MinFreeDisk {out, retval} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.MinFreeDisk); err != nil {
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

func (o *xxx_GetMinFreeDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// MinFreeDisk {out, retval} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.MinFreeDisk); err != nil {
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

// GetMinFreeDiskRequest structure represents the MinFreeDisk operation request
type GetMinFreeDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMinFreeDiskRequest) xxx_ToOp(ctx context.Context) *xxx_GetMinFreeDiskOperation {
	if o == nil {
		return &xxx_GetMinFreeDiskOperation{}
	}
	return &xxx_GetMinFreeDiskOperation{
		This: o.This,
	}
}

func (o *GetMinFreeDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMinFreeDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMinFreeDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMinFreeDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMinFreeDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMinFreeDiskResponse structure represents the MinFreeDisk operation response
type GetMinFreeDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MinFreeDisk uint32         `idl:"name:MinFreeDisk" json:"min_free_disk"`
	// Return: The MinFreeDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMinFreeDiskResponse) xxx_ToOp(ctx context.Context) *xxx_GetMinFreeDiskOperation {
	if o == nil {
		return &xxx_GetMinFreeDiskOperation{}
	}
	return &xxx_GetMinFreeDiskOperation{
		That:        o.That,
		MinFreeDisk: o.MinFreeDisk,
		Return:      o.Return,
	}
}

func (o *GetMinFreeDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMinFreeDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MinFreeDisk = op.MinFreeDisk
	o.Return = op.Return
}
func (o *GetMinFreeDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMinFreeDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMinFreeDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMinFreeDiskOperation structure represents the MinFreeDisk operation
type xxx_SetMinFreeDiskOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MinFreeDisk uint32         `idl:"name:MinFreeDisk" json:"min_free_disk"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMinFreeDiskOperation) OpNum() int { return 12 }

func (o *xxx_SetMinFreeDiskOperation) OpName() string { return "/IDataManager/v0/MinFreeDisk" }

func (o *xxx_SetMinFreeDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMinFreeDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// MinFreeDisk {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.MinFreeDisk); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMinFreeDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// MinFreeDisk {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.MinFreeDisk); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMinFreeDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMinFreeDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMinFreeDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMinFreeDiskRequest structure represents the MinFreeDisk operation request
type SetMinFreeDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	MinFreeDisk uint32         `idl:"name:MinFreeDisk" json:"min_free_disk"`
}

func (o *SetMinFreeDiskRequest) xxx_ToOp(ctx context.Context) *xxx_SetMinFreeDiskOperation {
	if o == nil {
		return &xxx_SetMinFreeDiskOperation{}
	}
	return &xxx_SetMinFreeDiskOperation{
		This:        o.This,
		MinFreeDisk: o.MinFreeDisk,
	}
}

func (o *SetMinFreeDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMinFreeDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MinFreeDisk = op.MinFreeDisk
}
func (o *SetMinFreeDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetMinFreeDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMinFreeDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMinFreeDiskResponse structure represents the MinFreeDisk operation response
type SetMinFreeDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MinFreeDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMinFreeDiskResponse) xxx_ToOp(ctx context.Context) *xxx_SetMinFreeDiskOperation {
	if o == nil {
		return &xxx_SetMinFreeDiskOperation{}
	}
	return &xxx_SetMinFreeDiskOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetMinFreeDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMinFreeDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMinFreeDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetMinFreeDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMinFreeDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMaxSizeOperation structure represents the MaxSize operation
type xxx_GetMaxSizeOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxSize uint32         `idl:"name:pulMaxSize" json:"max_size"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMaxSizeOperation) OpNum() int { return 13 }

func (o *xxx_GetMaxSizeOperation) OpName() string { return "/IDataManager/v0/MaxSize" }

func (o *xxx_GetMaxSizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaxSizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMaxSizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMaxSizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaxSizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pulMaxSize {out, retval} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.MaxSize); err != nil {
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

func (o *xxx_GetMaxSizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pulMaxSize {out, retval} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.MaxSize); err != nil {
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

// GetMaxSizeRequest structure represents the MaxSize operation request
type GetMaxSizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMaxSizeRequest) xxx_ToOp(ctx context.Context) *xxx_GetMaxSizeOperation {
	if o == nil {
		return &xxx_GetMaxSizeOperation{}
	}
	return &xxx_GetMaxSizeOperation{
		This: o.This,
	}
}

func (o *GetMaxSizeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMaxSizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMaxSizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMaxSizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMaxSizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMaxSizeResponse structure represents the MaxSize operation response
type GetMaxSizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxSize uint32         `idl:"name:pulMaxSize" json:"max_size"`
	// Return: The MaxSize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMaxSizeResponse) xxx_ToOp(ctx context.Context) *xxx_GetMaxSizeOperation {
	if o == nil {
		return &xxx_GetMaxSizeOperation{}
	}
	return &xxx_GetMaxSizeOperation{
		That:    o.That,
		MaxSize: o.MaxSize,
		Return:  o.Return,
	}
}

func (o *GetMaxSizeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMaxSizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MaxSize = op.MaxSize
	o.Return = op.Return
}
func (o *GetMaxSizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMaxSizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMaxSizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMaxSizeOperation structure represents the MaxSize operation
type xxx_SetMaxSizeOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxSize uint32         `idl:"name:ulMaxSize" json:"max_size"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMaxSizeOperation) OpNum() int { return 14 }

func (o *xxx_SetMaxSizeOperation) OpName() string { return "/IDataManager/v0/MaxSize" }

func (o *xxx_SetMaxSizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxSizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulMaxSize {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.MaxSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxSizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulMaxSize {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.MaxSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxSizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxSizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMaxSizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMaxSizeRequest structure represents the MaxSize operation request
type SetMaxSizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	MaxSize uint32         `idl:"name:ulMaxSize" json:"max_size"`
}

func (o *SetMaxSizeRequest) xxx_ToOp(ctx context.Context) *xxx_SetMaxSizeOperation {
	if o == nil {
		return &xxx_SetMaxSizeOperation{}
	}
	return &xxx_SetMaxSizeOperation{
		This:    o.This,
		MaxSize: o.MaxSize,
	}
}

func (o *SetMaxSizeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMaxSizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MaxSize = op.MaxSize
}
func (o *SetMaxSizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetMaxSizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMaxSizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMaxSizeResponse structure represents the MaxSize operation response
type SetMaxSizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MaxSize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMaxSizeResponse) xxx_ToOp(ctx context.Context) *xxx_SetMaxSizeOperation {
	if o == nil {
		return &xxx_SetMaxSizeOperation{}
	}
	return &xxx_SetMaxSizeOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetMaxSizeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMaxSizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMaxSizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetMaxSizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMaxSizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMaxFolderCountOperation structure represents the MaxFolderCount operation
type xxx_GetMaxFolderCountOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxFolderCount uint32         `idl:"name:pulMaxFolderCount" json:"max_folder_count"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMaxFolderCountOperation) OpNum() int { return 15 }

func (o *xxx_GetMaxFolderCountOperation) OpName() string { return "/IDataManager/v0/MaxFolderCount" }

func (o *xxx_GetMaxFolderCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaxFolderCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMaxFolderCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMaxFolderCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaxFolderCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pulMaxFolderCount {out, retval} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.MaxFolderCount); err != nil {
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

func (o *xxx_GetMaxFolderCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pulMaxFolderCount {out, retval} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.MaxFolderCount); err != nil {
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

// GetMaxFolderCountRequest structure represents the MaxFolderCount operation request
type GetMaxFolderCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMaxFolderCountRequest) xxx_ToOp(ctx context.Context) *xxx_GetMaxFolderCountOperation {
	if o == nil {
		return &xxx_GetMaxFolderCountOperation{}
	}
	return &xxx_GetMaxFolderCountOperation{
		This: o.This,
	}
}

func (o *GetMaxFolderCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMaxFolderCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMaxFolderCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMaxFolderCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMaxFolderCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMaxFolderCountResponse structure represents the MaxFolderCount operation response
type GetMaxFolderCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxFolderCount uint32         `idl:"name:pulMaxFolderCount" json:"max_folder_count"`
	// Return: The MaxFolderCount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMaxFolderCountResponse) xxx_ToOp(ctx context.Context) *xxx_GetMaxFolderCountOperation {
	if o == nil {
		return &xxx_GetMaxFolderCountOperation{}
	}
	return &xxx_GetMaxFolderCountOperation{
		That:           o.That,
		MaxFolderCount: o.MaxFolderCount,
		Return:         o.Return,
	}
}

func (o *GetMaxFolderCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMaxFolderCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MaxFolderCount = op.MaxFolderCount
	o.Return = op.Return
}
func (o *GetMaxFolderCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMaxFolderCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMaxFolderCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMaxFolderCountOperation structure represents the MaxFolderCount operation
type xxx_SetMaxFolderCountOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxFolderCount uint32         `idl:"name:ulMaxFolderCount" json:"max_folder_count"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMaxFolderCountOperation) OpNum() int { return 16 }

func (o *xxx_SetMaxFolderCountOperation) OpName() string { return "/IDataManager/v0/MaxFolderCount" }

func (o *xxx_SetMaxFolderCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxFolderCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulMaxFolderCount {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.MaxFolderCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxFolderCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulMaxFolderCount {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.MaxFolderCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxFolderCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxFolderCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMaxFolderCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMaxFolderCountRequest structure represents the MaxFolderCount operation request
type SetMaxFolderCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	MaxFolderCount uint32         `idl:"name:ulMaxFolderCount" json:"max_folder_count"`
}

func (o *SetMaxFolderCountRequest) xxx_ToOp(ctx context.Context) *xxx_SetMaxFolderCountOperation {
	if o == nil {
		return &xxx_SetMaxFolderCountOperation{}
	}
	return &xxx_SetMaxFolderCountOperation{
		This:           o.This,
		MaxFolderCount: o.MaxFolderCount,
	}
}

func (o *SetMaxFolderCountRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMaxFolderCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MaxFolderCount = op.MaxFolderCount
}
func (o *SetMaxFolderCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetMaxFolderCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMaxFolderCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMaxFolderCountResponse structure represents the MaxFolderCount operation response
type SetMaxFolderCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MaxFolderCount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMaxFolderCountResponse) xxx_ToOp(ctx context.Context) *xxx_SetMaxFolderCountOperation {
	if o == nil {
		return &xxx_SetMaxFolderCountOperation{}
	}
	return &xxx_SetMaxFolderCountOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetMaxFolderCountResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMaxFolderCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMaxFolderCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetMaxFolderCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMaxFolderCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetResourcePolicyOperation structure represents the ResourcePolicy operation
type xxx_GetResourcePolicyOperation struct {
	This   *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Policy pla.ResourcePolicy `idl:"name:pPolicy" json:"policy"`
	Return int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResourcePolicyOperation) OpNum() int { return 17 }

func (o *xxx_GetResourcePolicyOperation) OpName() string { return "/IDataManager/v0/ResourcePolicy" }

func (o *xxx_GetResourcePolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourcePolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetResourcePolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetResourcePolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourcePolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pPolicy {out, retval} (1:{pointer=ref}*(1))(2:{alias=ResourcePolicy}(enum))
	{
		if err := w.WriteData(uint16(o.Policy)); err != nil {
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

func (o *xxx_GetResourcePolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pPolicy {out, retval} (1:{pointer=ref}*(1))(2:{alias=ResourcePolicy}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Policy)); err != nil {
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

// GetResourcePolicyRequest structure represents the ResourcePolicy operation request
type GetResourcePolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetResourcePolicyRequest) xxx_ToOp(ctx context.Context) *xxx_GetResourcePolicyOperation {
	if o == nil {
		return &xxx_GetResourcePolicyOperation{}
	}
	return &xxx_GetResourcePolicyOperation{
		This: o.This,
	}
}

func (o *GetResourcePolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResourcePolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetResourcePolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetResourcePolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResourcePolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResourcePolicyResponse structure represents the ResourcePolicy operation response
type GetResourcePolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Policy pla.ResourcePolicy `idl:"name:pPolicy" json:"policy"`
	// Return: The ResourcePolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetResourcePolicyResponse) xxx_ToOp(ctx context.Context) *xxx_GetResourcePolicyOperation {
	if o == nil {
		return &xxx_GetResourcePolicyOperation{}
	}
	return &xxx_GetResourcePolicyOperation{
		That:   o.That,
		Policy: o.Policy,
		Return: o.Return,
	}
}

func (o *GetResourcePolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResourcePolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Policy = op.Policy
	o.Return = op.Return
}
func (o *GetResourcePolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetResourcePolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResourcePolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetResourcePolicyOperation structure represents the ResourcePolicy operation
type xxx_SetResourcePolicyOperation struct {
	This   *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Policy pla.ResourcePolicy `idl:"name:Policy" json:"policy"`
	Return int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_SetResourcePolicyOperation) OpNum() int { return 18 }

func (o *xxx_SetResourcePolicyOperation) OpName() string { return "/IDataManager/v0/ResourcePolicy" }

func (o *xxx_SetResourcePolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResourcePolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Policy {in} (1:{alias=ResourcePolicy}(enum))
	{
		if err := w.WriteData(uint16(o.Policy)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResourcePolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Policy {in} (1:{alias=ResourcePolicy}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Policy)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResourcePolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResourcePolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetResourcePolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetResourcePolicyRequest structure represents the ResourcePolicy operation request
type SetResourcePolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis     `idl:"name:This" json:"this"`
	Policy pla.ResourcePolicy `idl:"name:Policy" json:"policy"`
}

func (o *SetResourcePolicyRequest) xxx_ToOp(ctx context.Context) *xxx_SetResourcePolicyOperation {
	if o == nil {
		return &xxx_SetResourcePolicyOperation{}
	}
	return &xxx_SetResourcePolicyOperation{
		This:   o.This,
		Policy: o.Policy,
	}
}

func (o *SetResourcePolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetResourcePolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Policy = op.Policy
}
func (o *SetResourcePolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetResourcePolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetResourcePolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetResourcePolicyResponse structure represents the ResourcePolicy operation response
type SetResourcePolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ResourcePolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetResourcePolicyResponse) xxx_ToOp(ctx context.Context) *xxx_SetResourcePolicyOperation {
	if o == nil {
		return &xxx_SetResourcePolicyOperation{}
	}
	return &xxx_SetResourcePolicyOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetResourcePolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetResourcePolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetResourcePolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetResourcePolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetResourcePolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFolderActionsOperation structure represents the FolderActions operation
type xxx_GetFolderActionsOperation struct {
	This    *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Actions *pla.FolderActionCollection `idl:"name:Actions" json:"actions"`
	Return  int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFolderActionsOperation) OpNum() int { return 19 }

func (o *xxx_GetFolderActionsOperation) OpName() string { return "/IDataManager/v0/FolderActions" }

func (o *xxx_GetFolderActionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFolderActionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFolderActionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFolderActionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFolderActionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Actions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFolderActionCollection}(interface))
	{
		if o.Actions != nil {
			_ptr_Actions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Actions != nil {
					if err := o.Actions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.FolderActionCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Actions, _ptr_Actions); err != nil {
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

func (o *xxx_GetFolderActionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Actions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFolderActionCollection}(interface))
	{
		_ptr_Actions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Actions == nil {
				o.Actions = &pla.FolderActionCollection{}
			}
			if err := o.Actions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Actions := func(ptr interface{}) { o.Actions = *ptr.(**pla.FolderActionCollection) }
		if err := w.ReadPointer(&o.Actions, _s_Actions, _ptr_Actions); err != nil {
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

// GetFolderActionsRequest structure represents the FolderActions operation request
type GetFolderActionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFolderActionsRequest) xxx_ToOp(ctx context.Context) *xxx_GetFolderActionsOperation {
	if o == nil {
		return &xxx_GetFolderActionsOperation{}
	}
	return &xxx_GetFolderActionsOperation{
		This: o.This,
	}
}

func (o *GetFolderActionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFolderActionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFolderActionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetFolderActionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFolderActionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFolderActionsResponse structure represents the FolderActions operation response
type GetFolderActionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Actions: Receives a pointer to a folder action collection object.
	Actions *pla.FolderActionCollection `idl:"name:Actions" json:"actions"`
	// Return: The FolderActions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFolderActionsResponse) xxx_ToOp(ctx context.Context) *xxx_GetFolderActionsOperation {
	if o == nil {
		return &xxx_GetFolderActionsOperation{}
	}
	return &xxx_GetFolderActionsOperation{
		That:    o.That,
		Actions: o.Actions,
		Return:  o.Return,
	}
}

func (o *GetFolderActionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFolderActionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Actions = op.Actions
	o.Return = op.Return
}
func (o *GetFolderActionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetFolderActionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFolderActionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetReportSchemaOperation structure represents the ReportSchema operation
type xxx_GetReportSchemaOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReportSchema *oaut.String   `idl:"name:ReportSchema" json:"report_schema"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReportSchemaOperation) OpNum() int { return 20 }

func (o *xxx_GetReportSchemaOperation) OpName() string { return "/IDataManager/v0/ReportSchema" }

func (o *xxx_GetReportSchemaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportSchemaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetReportSchemaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetReportSchemaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportSchemaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ReportSchema {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReportSchema != nil {
			_ptr_ReportSchema := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReportSchema != nil {
					if err := o.ReportSchema.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReportSchema, _ptr_ReportSchema); err != nil {
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

func (o *xxx_GetReportSchemaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ReportSchema {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_ReportSchema := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReportSchema == nil {
				o.ReportSchema = &oaut.String{}
			}
			if err := o.ReportSchema.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReportSchema := func(ptr interface{}) { o.ReportSchema = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReportSchema, _s_ReportSchema, _ptr_ReportSchema); err != nil {
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

// GetReportSchemaRequest structure represents the ReportSchema operation request
type GetReportSchemaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetReportSchemaRequest) xxx_ToOp(ctx context.Context) *xxx_GetReportSchemaOperation {
	if o == nil {
		return &xxx_GetReportSchemaOperation{}
	}
	return &xxx_GetReportSchemaOperation{
		This: o.This,
	}
}

func (o *GetReportSchemaRequest) xxx_FromOp(ctx context.Context, op *xxx_GetReportSchemaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetReportSchemaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetReportSchemaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportSchemaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReportSchemaResponse structure represents the ReportSchema operation response
type GetReportSchemaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReportSchema *oaut.String   `idl:"name:ReportSchema" json:"report_schema"`
	// Return: The ReportSchema return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReportSchemaResponse) xxx_ToOp(ctx context.Context) *xxx_GetReportSchemaOperation {
	if o == nil {
		return &xxx_GetReportSchemaOperation{}
	}
	return &xxx_GetReportSchemaOperation{
		That:         o.That,
		ReportSchema: o.ReportSchema,
		Return:       o.Return,
	}
}

func (o *GetReportSchemaResponse) xxx_FromOp(ctx context.Context, op *xxx_GetReportSchemaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReportSchema = op.ReportSchema
	o.Return = op.Return
}
func (o *GetReportSchemaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetReportSchemaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportSchemaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetReportSchemaOperation structure represents the ReportSchema operation
type xxx_SetReportSchemaOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReportSchema *oaut.String   `idl:"name:ReportSchema" json:"report_schema"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetReportSchemaOperation) OpNum() int { return 21 }

func (o *xxx_SetReportSchemaOperation) OpName() string { return "/IDataManager/v0/ReportSchema" }

func (o *xxx_SetReportSchemaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportSchemaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ReportSchema {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReportSchema != nil {
			_ptr_ReportSchema := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReportSchema != nil {
					if err := o.ReportSchema.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReportSchema, _ptr_ReportSchema); err != nil {
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

func (o *xxx_SetReportSchemaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ReportSchema {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_ReportSchema := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReportSchema == nil {
				o.ReportSchema = &oaut.String{}
			}
			if err := o.ReportSchema.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReportSchema := func(ptr interface{}) { o.ReportSchema = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReportSchema, _s_ReportSchema, _ptr_ReportSchema); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportSchemaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportSchemaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetReportSchemaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetReportSchemaRequest structure represents the ReportSchema operation request
type SetReportSchemaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	ReportSchema *oaut.String   `idl:"name:ReportSchema" json:"report_schema"`
}

func (o *SetReportSchemaRequest) xxx_ToOp(ctx context.Context) *xxx_SetReportSchemaOperation {
	if o == nil {
		return &xxx_SetReportSchemaOperation{}
	}
	return &xxx_SetReportSchemaOperation{
		This:         o.This,
		ReportSchema: o.ReportSchema,
	}
}

func (o *SetReportSchemaRequest) xxx_FromOp(ctx context.Context, op *xxx_SetReportSchemaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReportSchema = op.ReportSchema
}
func (o *SetReportSchemaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetReportSchemaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetReportSchemaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetReportSchemaResponse structure represents the ReportSchema operation response
type SetReportSchemaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ReportSchema return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetReportSchemaResponse) xxx_ToOp(ctx context.Context) *xxx_SetReportSchemaOperation {
	if o == nil {
		return &xxx_SetReportSchemaOperation{}
	}
	return &xxx_SetReportSchemaOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetReportSchemaResponse) xxx_FromOp(ctx context.Context, op *xxx_SetReportSchemaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetReportSchemaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetReportSchemaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetReportSchemaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetReportFileNameOperation structure represents the ReportFileName operation
type xxx_GetReportFileNameOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileName *oaut.String   `idl:"name:pbstrFilename" json:"file_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReportFileNameOperation) OpNum() int { return 22 }

func (o *xxx_GetReportFileNameOperation) OpName() string { return "/IDataManager/v0/ReportFileName" }

func (o *xxx_GetReportFileNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportFileNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetReportFileNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetReportFileNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportFileNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrFilename {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FileName != nil {
			_ptr_pbstrFilename := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.FileName, _ptr_pbstrFilename); err != nil {
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

func (o *xxx_GetReportFileNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrFilename {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrFilename := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileName == nil {
				o.FileName = &oaut.String{}
			}
			if err := o.FileName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrFilename := func(ptr interface{}) { o.FileName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FileName, _s_pbstrFilename, _ptr_pbstrFilename); err != nil {
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

// GetReportFileNameRequest structure represents the ReportFileName operation request
type GetReportFileNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetReportFileNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetReportFileNameOperation {
	if o == nil {
		return &xxx_GetReportFileNameOperation{}
	}
	return &xxx_GetReportFileNameOperation{
		This: o.This,
	}
}

func (o *GetReportFileNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetReportFileNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetReportFileNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetReportFileNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportFileNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReportFileNameResponse structure represents the ReportFileName operation response
type GetReportFileNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileName *oaut.String   `idl:"name:pbstrFilename" json:"file_name"`
	// Return: The ReportFileName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReportFileNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetReportFileNameOperation {
	if o == nil {
		return &xxx_GetReportFileNameOperation{}
	}
	return &xxx_GetReportFileNameOperation{
		That:     o.That,
		FileName: o.FileName,
		Return:   o.Return,
	}
}

func (o *GetReportFileNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetReportFileNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileName = op.FileName
	o.Return = op.Return
}
func (o *GetReportFileNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetReportFileNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportFileNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetReportFileNameOperation structure represents the ReportFileName operation
type xxx_SetReportFileNameOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileName *oaut.String   `idl:"name:pbstrFilename" json:"file_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetReportFileNameOperation) OpNum() int { return 23 }

func (o *xxx_SetReportFileNameOperation) OpName() string { return "/IDataManager/v0/ReportFileName" }

func (o *xxx_SetReportFileNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportFileNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pbstrFilename {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FileName != nil {
			_ptr_pbstrFilename := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.FileName, _ptr_pbstrFilename); err != nil {
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

func (o *xxx_SetReportFileNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pbstrFilename {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrFilename := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileName == nil {
				o.FileName = &oaut.String{}
			}
			if err := o.FileName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrFilename := func(ptr interface{}) { o.FileName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FileName, _s_pbstrFilename, _ptr_pbstrFilename); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportFileNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportFileNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetReportFileNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetReportFileNameRequest structure represents the ReportFileName operation request
type SetReportFileNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	FileName *oaut.String   `idl:"name:pbstrFilename" json:"file_name"`
}

func (o *SetReportFileNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetReportFileNameOperation {
	if o == nil {
		return &xxx_SetReportFileNameOperation{}
	}
	return &xxx_SetReportFileNameOperation{
		This:     o.This,
		FileName: o.FileName,
	}
}

func (o *SetReportFileNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetReportFileNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FileName = op.FileName
}
func (o *SetReportFileNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetReportFileNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetReportFileNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetReportFileNameResponse structure represents the ReportFileName operation response
type SetReportFileNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ReportFileName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetReportFileNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetReportFileNameOperation {
	if o == nil {
		return &xxx_SetReportFileNameOperation{}
	}
	return &xxx_SetReportFileNameOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetReportFileNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetReportFileNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetReportFileNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetReportFileNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetReportFileNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRuleTargetFileNameOperation structure represents the RuleTargetFileName operation
type xxx_GetRuleTargetFileNameOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileName *oaut.String   `idl:"name:Filename" json:"file_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRuleTargetFileNameOperation) OpNum() int { return 24 }

func (o *xxx_GetRuleTargetFileNameOperation) OpName() string {
	return "/IDataManager/v0/RuleTargetFileName"
}

func (o *xxx_GetRuleTargetFileNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRuleTargetFileNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRuleTargetFileNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRuleTargetFileNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRuleTargetFileNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Filename {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FileName != nil {
			_ptr_Filename := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.FileName, _ptr_Filename); err != nil {
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

func (o *xxx_GetRuleTargetFileNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Filename {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_Filename := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileName == nil {
				o.FileName = &oaut.String{}
			}
			if err := o.FileName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Filename := func(ptr interface{}) { o.FileName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FileName, _s_Filename, _ptr_Filename); err != nil {
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

// GetRuleTargetFileNameRequest structure represents the RuleTargetFileName operation request
type GetRuleTargetFileNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRuleTargetFileNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetRuleTargetFileNameOperation {
	if o == nil {
		return &xxx_GetRuleTargetFileNameOperation{}
	}
	return &xxx_GetRuleTargetFileNameOperation{
		This: o.This,
	}
}

func (o *GetRuleTargetFileNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRuleTargetFileNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRuleTargetFileNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetRuleTargetFileNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRuleTargetFileNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRuleTargetFileNameResponse structure represents the RuleTargetFileName operation response
type GetRuleTargetFileNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileName *oaut.String   `idl:"name:Filename" json:"file_name"`
	// Return: The RuleTargetFileName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRuleTargetFileNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetRuleTargetFileNameOperation {
	if o == nil {
		return &xxx_GetRuleTargetFileNameOperation{}
	}
	return &xxx_GetRuleTargetFileNameOperation{
		That:     o.That,
		FileName: o.FileName,
		Return:   o.Return,
	}
}

func (o *GetRuleTargetFileNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRuleTargetFileNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileName = op.FileName
	o.Return = op.Return
}
func (o *GetRuleTargetFileNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetRuleTargetFileNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRuleTargetFileNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetRuleTargetFileNameOperation structure represents the RuleTargetFileName operation
type xxx_SetRuleTargetFileNameOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileName *oaut.String   `idl:"name:Filename" json:"file_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetRuleTargetFileNameOperation) OpNum() int { return 25 }

func (o *xxx_SetRuleTargetFileNameOperation) OpName() string {
	return "/IDataManager/v0/RuleTargetFileName"
}

func (o *xxx_SetRuleTargetFileNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRuleTargetFileNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Filename {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FileName != nil {
			_ptr_Filename := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.FileName, _ptr_Filename); err != nil {
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

func (o *xxx_SetRuleTargetFileNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Filename {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_Filename := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileName == nil {
				o.FileName = &oaut.String{}
			}
			if err := o.FileName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Filename := func(ptr interface{}) { o.FileName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FileName, _s_Filename, _ptr_Filename); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRuleTargetFileNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRuleTargetFileNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetRuleTargetFileNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetRuleTargetFileNameRequest structure represents the RuleTargetFileName operation request
type SetRuleTargetFileNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	FileName *oaut.String   `idl:"name:Filename" json:"file_name"`
}

func (o *SetRuleTargetFileNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetRuleTargetFileNameOperation {
	if o == nil {
		return &xxx_SetRuleTargetFileNameOperation{}
	}
	return &xxx_SetRuleTargetFileNameOperation{
		This:     o.This,
		FileName: o.FileName,
	}
}

func (o *SetRuleTargetFileNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetRuleTargetFileNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FileName = op.FileName
}
func (o *SetRuleTargetFileNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetRuleTargetFileNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRuleTargetFileNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetRuleTargetFileNameResponse structure represents the RuleTargetFileName operation response
type SetRuleTargetFileNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RuleTargetFileName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetRuleTargetFileNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetRuleTargetFileNameOperation {
	if o == nil {
		return &xxx_SetRuleTargetFileNameOperation{}
	}
	return &xxx_SetRuleTargetFileNameOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetRuleTargetFileNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetRuleTargetFileNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetRuleTargetFileNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetRuleTargetFileNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRuleTargetFileNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEventsFileNameOperation structure represents the EventsFileName operation
type xxx_GetEventsFileNameOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileName *oaut.String   `idl:"name:pbstrFilename" json:"file_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventsFileNameOperation) OpNum() int { return 26 }

func (o *xxx_GetEventsFileNameOperation) OpName() string { return "/IDataManager/v0/EventsFileName" }

func (o *xxx_GetEventsFileNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventsFileNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetEventsFileNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetEventsFileNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventsFileNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrFilename {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FileName != nil {
			_ptr_pbstrFilename := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.FileName, _ptr_pbstrFilename); err != nil {
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

func (o *xxx_GetEventsFileNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrFilename {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrFilename := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileName == nil {
				o.FileName = &oaut.String{}
			}
			if err := o.FileName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrFilename := func(ptr interface{}) { o.FileName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FileName, _s_pbstrFilename, _ptr_pbstrFilename); err != nil {
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

// GetEventsFileNameRequest structure represents the EventsFileName operation request
type GetEventsFileNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEventsFileNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetEventsFileNameOperation {
	if o == nil {
		return &xxx_GetEventsFileNameOperation{}
	}
	return &xxx_GetEventsFileNameOperation{
		This: o.This,
	}
}

func (o *GetEventsFileNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventsFileNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEventsFileNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetEventsFileNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventsFileNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEventsFileNameResponse structure represents the EventsFileName operation response
type GetEventsFileNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileName *oaut.String   `idl:"name:pbstrFilename" json:"file_name"`
	// Return: The EventsFileName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEventsFileNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetEventsFileNameOperation {
	if o == nil {
		return &xxx_GetEventsFileNameOperation{}
	}
	return &xxx_GetEventsFileNameOperation{
		That:     o.That,
		FileName: o.FileName,
		Return:   o.Return,
	}
}

func (o *GetEventsFileNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEventsFileNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileName = op.FileName
	o.Return = op.Return
}
func (o *GetEventsFileNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetEventsFileNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventsFileNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEventsFileNameOperation structure represents the EventsFileName operation
type xxx_SetEventsFileNameOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileName *oaut.String   `idl:"name:pbstrFilename" json:"file_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEventsFileNameOperation) OpNum() int { return 27 }

func (o *xxx_SetEventsFileNameOperation) OpName() string { return "/IDataManager/v0/EventsFileName" }

func (o *xxx_SetEventsFileNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventsFileNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pbstrFilename {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FileName != nil {
			_ptr_pbstrFilename := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.FileName, _ptr_pbstrFilename); err != nil {
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

func (o *xxx_SetEventsFileNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pbstrFilename {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrFilename := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileName == nil {
				o.FileName = &oaut.String{}
			}
			if err := o.FileName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrFilename := func(ptr interface{}) { o.FileName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FileName, _s_pbstrFilename, _ptr_pbstrFilename); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventsFileNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventsFileNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEventsFileNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetEventsFileNameRequest structure represents the EventsFileName operation request
type SetEventsFileNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	FileName *oaut.String   `idl:"name:pbstrFilename" json:"file_name"`
}

func (o *SetEventsFileNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetEventsFileNameOperation {
	if o == nil {
		return &xxx_SetEventsFileNameOperation{}
	}
	return &xxx_SetEventsFileNameOperation{
		This:     o.This,
		FileName: o.FileName,
	}
}

func (o *SetEventsFileNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEventsFileNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FileName = op.FileName
}
func (o *SetEventsFileNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetEventsFileNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventsFileNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEventsFileNameResponse structure represents the EventsFileName operation response
type SetEventsFileNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EventsFileName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEventsFileNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetEventsFileNameOperation {
	if o == nil {
		return &xxx_SetEventsFileNameOperation{}
	}
	return &xxx_SetEventsFileNameOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetEventsFileNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEventsFileNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEventsFileNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetEventsFileNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventsFileNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRulesOperation structure represents the Rules operation
type xxx_GetRulesOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	XML    *oaut.String   `idl:"name:pbstrXml" json:"xml"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRulesOperation) OpNum() int { return 28 }

func (o *xxx_GetRulesOperation) OpName() string { return "/IDataManager/v0/Rules" }

func (o *xxx_GetRulesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRulesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRulesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRulesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRulesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrXml {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.XML != nil {
			_ptr_pbstrXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.XML != nil {
					if err := o.XML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.XML, _ptr_pbstrXml); err != nil {
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

func (o *xxx_GetRulesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrXml {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.XML == nil {
				o.XML = &oaut.String{}
			}
			if err := o.XML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrXml := func(ptr interface{}) { o.XML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.XML, _s_pbstrXml, _ptr_pbstrXml); err != nil {
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

// GetRulesRequest structure represents the Rules operation request
type GetRulesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRulesRequest) xxx_ToOp(ctx context.Context) *xxx_GetRulesOperation {
	if o == nil {
		return &xxx_GetRulesOperation{}
	}
	return &xxx_GetRulesOperation{
		This: o.This,
	}
}

func (o *GetRulesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRulesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRulesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetRulesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRulesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRulesResponse structure represents the Rules operation response
type GetRulesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	XML  *oaut.String   `idl:"name:pbstrXml" json:"xml"`
	// Return: The Rules return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRulesResponse) xxx_ToOp(ctx context.Context) *xxx_GetRulesOperation {
	if o == nil {
		return &xxx_GetRulesOperation{}
	}
	return &xxx_GetRulesOperation{
		That:   o.That,
		XML:    o.XML,
		Return: o.Return,
	}
}

func (o *GetRulesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRulesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.XML = op.XML
	o.Return = op.Return
}
func (o *GetRulesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetRulesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRulesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetRulesOperation structure represents the Rules operation
type xxx_SetRulesOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	XML    *oaut.String   `idl:"name:bstrXml" json:"xml"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetRulesOperation) OpNum() int { return 29 }

func (o *xxx_SetRulesOperation) OpName() string { return "/IDataManager/v0/Rules" }

func (o *xxx_SetRulesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRulesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.XML != nil {
			_ptr_bstrXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.XML != nil {
					if err := o.XML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.XML, _ptr_bstrXml); err != nil {
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

func (o *xxx_SetRulesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.XML == nil {
				o.XML = &oaut.String{}
			}
			if err := o.XML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrXml := func(ptr interface{}) { o.XML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.XML, _s_bstrXml, _ptr_bstrXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRulesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRulesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetRulesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetRulesRequest structure represents the Rules operation request
type SetRulesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	XML  *oaut.String   `idl:"name:bstrXml" json:"xml"`
}

func (o *SetRulesRequest) xxx_ToOp(ctx context.Context) *xxx_SetRulesOperation {
	if o == nil {
		return &xxx_SetRulesOperation{}
	}
	return &xxx_SetRulesOperation{
		This: o.This,
		XML:  o.XML,
	}
}

func (o *SetRulesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetRulesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.XML = op.XML
}
func (o *SetRulesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetRulesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRulesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetRulesResponse structure represents the Rules operation response
type SetRulesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Rules return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetRulesResponse) xxx_ToOp(ctx context.Context) *xxx_SetRulesOperation {
	if o == nil {
		return &xxx_SetRulesOperation{}
	}
	return &xxx_SetRulesOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetRulesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetRulesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetRulesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetRulesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRulesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RunOperation structure represents the Run operation
type xxx_RunOperation struct {
	This   *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Steps  pla.DataManagerSteps `idl:"name:Steps" json:"steps"`
	Folder *oaut.String         `idl:"name:bstrFolder" json:"folder"`
	Errors *pla.ValueMap        `idl:"name:Errors" json:"errors"`
	Return int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_RunOperation) OpNum() int { return 30 }

func (o *xxx_RunOperation) OpName() string { return "/IDataManager/v0/Run" }

func (o *xxx_RunOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Steps {in} (1:{alias=DataManagerSteps}(enum))
	{
		if err := w.WriteData(uint16(o.Steps)); err != nil {
			return err
		}
	}
	// bstrFolder {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Folder != nil {
			_ptr_bstrFolder := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Folder != nil {
					if err := o.Folder.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Folder, _ptr_bstrFolder); err != nil {
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

func (o *xxx_RunOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Steps {in} (1:{alias=DataManagerSteps}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Steps)); err != nil {
			return err
		}
	}
	// bstrFolder {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrFolder := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Folder == nil {
				o.Folder = &oaut.String{}
			}
			if err := o.Folder.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrFolder := func(ptr interface{}) { o.Folder = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Folder, _s_bstrFolder, _ptr_bstrFolder); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Errors {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IValueMap}(interface))
	{
		if o.Errors != nil {
			_ptr_Errors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Errors != nil {
					if err := o.Errors.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.ValueMap{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Errors, _ptr_Errors); err != nil {
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

func (o *xxx_RunOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Errors {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IValueMap}(interface))
	{
		_ptr_Errors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Errors == nil {
				o.Errors = &pla.ValueMap{}
			}
			if err := o.Errors.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Errors := func(ptr interface{}) { o.Errors = *ptr.(**pla.ValueMap) }
		if err := w.ReadPointer(&o.Errors, _s_Errors, _ptr_Errors); err != nil {
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

// RunRequest structure represents the Run operation request
type RunRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Steps: Supplies the actions to be performed by the data manager. For more information,
	// see section 2.2.2.6.
	Steps pla.DataManagerSteps `idl:"name:Steps" json:"steps"`
	// bstrFolder: Supplies the name of the subfolder where the report will be generated,
	// rules applied, and/or HTML created.
	Folder *oaut.String `idl:"name:bstrFolder" json:"folder"`
}

func (o *RunRequest) xxx_ToOp(ctx context.Context) *xxx_RunOperation {
	if o == nil {
		return &xxx_RunOperation{}
	}
	return &xxx_RunOperation{
		This:   o.This,
		Steps:  o.Steps,
		Folder: o.Folder,
	}
}

func (o *RunRequest) xxx_FromOp(ctx context.Context, op *xxx_RunOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Steps = op.Steps
	o.Folder = op.Folder
}
func (o *RunRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RunRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RunOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RunResponse structure represents the Run operation response
type RunResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Errors: Receives a validation value map, stored as an IValueMap, containing a list
	// of subfolders on which the data manager did not successfully run; each subfolder
	// is represented as an IValueMapItem. The Names property of the IValueMapItem represents
	// the path to a subfolder, while the Values property of the IValueMapItem represents
	// the HRESULT describing the specific problem with that subfolder. The ValueMapType
	// of the IValueMap is plaValidation; more information can be found in section 2.2.11.
	Errors *pla.ValueMap `idl:"name:Errors" json:"errors"`
	// Return: The Run return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RunResponse) xxx_ToOp(ctx context.Context) *xxx_RunOperation {
	if o == nil {
		return &xxx_RunOperation{}
	}
	return &xxx_RunOperation{
		That:   o.That,
		Errors: o.Errors,
		Return: o.Return,
	}
}

func (o *RunResponse) xxx_FromOp(ctx context.Context, op *xxx_RunOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Errors = op.Errors
	o.Return = op.Return
}
func (o *RunResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RunResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RunOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExtractOperation structure represents the Extract operation
type xxx_ExtractOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	CabFileName     *oaut.String   `idl:"name:CabFilename" json:"cab_file_name"`
	DestinationPath *oaut.String   `idl:"name:DestinationPath" json:"destination_path"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ExtractOperation) OpNum() int { return 31 }

func (o *xxx_ExtractOperation) OpName() string { return "/IDataManager/v0/Extract" }

func (o *xxx_ExtractOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExtractOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// CabFilename {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CabFileName != nil {
			_ptr_CabFilename := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CabFileName != nil {
					if err := o.CabFileName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CabFileName, _ptr_CabFilename); err != nil {
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
	// DestinationPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DestinationPath != nil {
			_ptr_DestinationPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DestinationPath != nil {
					if err := o.DestinationPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DestinationPath, _ptr_DestinationPath); err != nil {
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

func (o *xxx_ExtractOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// CabFilename {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_CabFilename := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CabFileName == nil {
				o.CabFileName = &oaut.String{}
			}
			if err := o.CabFileName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_CabFilename := func(ptr interface{}) { o.CabFileName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CabFileName, _s_CabFilename, _ptr_CabFilename); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DestinationPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_DestinationPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DestinationPath == nil {
				o.DestinationPath = &oaut.String{}
			}
			if err := o.DestinationPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DestinationPath := func(ptr interface{}) { o.DestinationPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DestinationPath, _s_DestinationPath, _ptr_DestinationPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExtractOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExtractOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ExtractOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ExtractRequest structure represents the Extract operation request
type ExtractRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// CabFilename: The name of the CAB file to extract.
	CabFileName *oaut.String `idl:"name:CabFilename" json:"cab_file_name"`
	// DestinationPath: The destination path for where to place the CAB file.
	DestinationPath *oaut.String `idl:"name:DestinationPath" json:"destination_path"`
}

func (o *ExtractRequest) xxx_ToOp(ctx context.Context) *xxx_ExtractOperation {
	if o == nil {
		return &xxx_ExtractOperation{}
	}
	return &xxx_ExtractOperation{
		This:            o.This,
		CabFileName:     o.CabFileName,
		DestinationPath: o.DestinationPath,
	}
}

func (o *ExtractRequest) xxx_FromOp(ctx context.Context, op *xxx_ExtractOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CabFileName = op.CabFileName
	o.DestinationPath = op.DestinationPath
}
func (o *ExtractRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExtractRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExtractOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExtractResponse structure represents the Extract operation response
type ExtractResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Extract return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExtractResponse) xxx_ToOp(ctx context.Context) *xxx_ExtractOperation {
	if o == nil {
		return &xxx_ExtractOperation{}
	}
	return &xxx_ExtractOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ExtractResponse) xxx_FromOp(ctx context.Context, op *xxx_ExtractOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ExtractResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExtractResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExtractOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
