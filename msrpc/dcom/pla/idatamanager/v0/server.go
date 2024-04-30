package idatamanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
	_ = idispatch.GoPackage
)

// IDataManager server interface.
type DataManagerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Enabled operation.
	GetEnabled(context.Context, *GetEnabledRequest) (*GetEnabledResponse, error)

	// Enabled operation.
	SetEnabled(context.Context, *SetEnabledRequest) (*SetEnabledResponse, error)

	// CheckBeforeRunning operation.
	GetCheckBeforeRunning(context.Context, *GetCheckBeforeRunningRequest) (*GetCheckBeforeRunningResponse, error)

	// CheckBeforeRunning operation.
	SetCheckBeforeRunning(context.Context, *SetCheckBeforeRunningRequest) (*SetCheckBeforeRunningResponse, error)

	// MinFreeDisk operation.
	GetMinFreeDisk(context.Context, *GetMinFreeDiskRequest) (*GetMinFreeDiskResponse, error)

	// MinFreeDisk operation.
	SetMinFreeDisk(context.Context, *SetMinFreeDiskRequest) (*SetMinFreeDiskResponse, error)

	// MaxSize operation.
	GetMaxSize(context.Context, *GetMaxSizeRequest) (*GetMaxSizeResponse, error)

	// MaxSize operation.
	SetMaxSize(context.Context, *SetMaxSizeRequest) (*SetMaxSizeResponse, error)

	// MaxFolderCount operation.
	GetMaxFolderCount(context.Context, *GetMaxFolderCountRequest) (*GetMaxFolderCountResponse, error)

	// MaxFolderCount operation.
	SetMaxFolderCount(context.Context, *SetMaxFolderCountRequest) (*SetMaxFolderCountResponse, error)

	// The ResourcePolicy enumeration defines the order in which folders are deleted when
	// one of the disk resource limits is exceeded.
	GetResourcePolicy(context.Context, *GetResourcePolicyRequest) (*GetResourcePolicyResponse, error)

	// The ResourcePolicy enumeration defines the order in which folders are deleted when
	// one of the disk resource limits is exceeded.
	SetResourcePolicy(context.Context, *SetResourcePolicyRequest) (*SetResourcePolicyResponse, error)

	// The FolderActions (Get) method retrieves the FolderAction property, as specified
	// in the property table in section 3.2.4.2.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetFolderActions(context.Context, *GetFolderActionsRequest) (*GetFolderActionsResponse, error)

	// ReportSchema operation.
	GetReportSchema(context.Context, *GetReportSchemaRequest) (*GetReportSchemaResponse, error)

	// ReportSchema operation.
	SetReportSchema(context.Context, *SetReportSchemaRequest) (*SetReportSchemaResponse, error)

	// ReportFileName operation.
	GetReportFileName(context.Context, *GetReportFileNameRequest) (*GetReportFileNameResponse, error)

	// ReportFileName operation.
	SetReportFileName(context.Context, *SetReportFileNameRequest) (*SetReportFileNameResponse, error)

	// RuleTargetFileName operation.
	GetRuleTargetFileName(context.Context, *GetRuleTargetFileNameRequest) (*GetRuleTargetFileNameResponse, error)

	// RuleTargetFileName operation.
	SetRuleTargetFileName(context.Context, *SetRuleTargetFileNameRequest) (*SetRuleTargetFileNameResponse, error)

	// EventsFileName operation.
	GetEventsFileName(context.Context, *GetEventsFileNameRequest) (*GetEventsFileNameResponse, error)

	// EventsFileName operation.
	SetEventsFileName(context.Context, *SetEventsFileNameRequest) (*SetEventsFileNameResponse, error)

	// Rules operation.
	GetRules(context.Context, *GetRulesRequest) (*GetRulesResponse, error)

	// Rules operation.
	SetRules(context.Context, *SetRulesRequest) (*SetRulesResponse, error)

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
	Run(context.Context, *RunRequest) (*RunResponse, error)

	// The Extract method extracts the specified CAB file.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	Extract(context.Context, *ExtractRequest) (*ExtractResponse, error)
}

func RegisterDataManagerServer(conn dcerpc.Conn, o DataManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDataManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DataManagerSyntaxV0_0))...)
}

func NewDataManagerServerHandle(o DataManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DataManagerServerHandle(ctx, o, opNum, r)
	}
}

func DataManagerServerHandle(ctx context.Context, o DataManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Enabled
		in := &GetEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Enabled
		in := &SetEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // CheckBeforeRunning
		in := &GetCheckBeforeRunningRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCheckBeforeRunning(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // CheckBeforeRunning
		in := &SetCheckBeforeRunningRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetCheckBeforeRunning(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // MinFreeDisk
		in := &GetMinFreeDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMinFreeDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // MinFreeDisk
		in := &SetMinFreeDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMinFreeDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // MaxSize
		in := &GetMaxSizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMaxSize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // MaxSize
		in := &SetMaxSizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMaxSize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // MaxFolderCount
		in := &GetMaxFolderCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMaxFolderCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // MaxFolderCount
		in := &SetMaxFolderCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMaxFolderCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // ResourcePolicy
		in := &GetResourcePolicyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetResourcePolicy(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // ResourcePolicy
		in := &SetResourcePolicyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetResourcePolicy(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // FolderActions
		in := &GetFolderActionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFolderActions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // ReportSchema
		in := &GetReportSchemaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetReportSchema(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // ReportSchema
		in := &SetReportSchemaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetReportSchema(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // ReportFileName
		in := &GetReportFileNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetReportFileName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // ReportFileName
		in := &SetReportFileNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetReportFileName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // RuleTargetFileName
		in := &GetRuleTargetFileNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRuleTargetFileName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // RuleTargetFileName
		in := &SetRuleTargetFileNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetRuleTargetFileName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // EventsFileName
		in := &GetEventsFileNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventsFileName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // EventsFileName
		in := &SetEventsFileNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEventsFileName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // Rules
		in := &GetRulesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRules(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // Rules
		in := &SetRulesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetRules(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // Run
		in := &RunRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Run(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // Extract
		in := &ExtractRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Extract(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
