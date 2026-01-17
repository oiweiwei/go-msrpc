package idatamanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_GetEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Enabled
		op := &xxx_SetEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // CheckBeforeRunning
		op := &xxx_GetCheckBeforeRunningOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCheckBeforeRunningRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCheckBeforeRunning(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // CheckBeforeRunning
		op := &xxx_SetCheckBeforeRunningOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCheckBeforeRunningRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCheckBeforeRunning(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // MinFreeDisk
		op := &xxx_GetMinFreeDiskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMinFreeDiskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMinFreeDisk(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // MinFreeDisk
		op := &xxx_SetMinFreeDiskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMinFreeDiskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMinFreeDisk(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // MaxSize
		op := &xxx_GetMaxSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMaxSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMaxSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // MaxSize
		op := &xxx_SetMaxSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMaxSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMaxSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // MaxFolderCount
		op := &xxx_GetMaxFolderCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMaxFolderCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMaxFolderCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // MaxFolderCount
		op := &xxx_SetMaxFolderCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMaxFolderCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMaxFolderCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // ResourcePolicy
		op := &xxx_GetResourcePolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResourcePolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResourcePolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // ResourcePolicy
		op := &xxx_SetResourcePolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetResourcePolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetResourcePolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // FolderActions
		op := &xxx_GetFolderActionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFolderActionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFolderActions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // ReportSchema
		op := &xxx_GetReportSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReportSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReportSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // ReportSchema
		op := &xxx_SetReportSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetReportSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetReportSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // ReportFileName
		op := &xxx_GetReportFileNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReportFileNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReportFileName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // ReportFileName
		op := &xxx_SetReportFileNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetReportFileNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetReportFileName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // RuleTargetFileName
		op := &xxx_GetRuleTargetFileNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRuleTargetFileNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRuleTargetFileName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // RuleTargetFileName
		op := &xxx_SetRuleTargetFileNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetRuleTargetFileNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetRuleTargetFileName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // EventsFileName
		op := &xxx_GetEventsFileNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEventsFileNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEventsFileName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // EventsFileName
		op := &xxx_SetEventsFileNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEventsFileNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEventsFileName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // Rules
		op := &xxx_GetRulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // Rules
		op := &xxx_SetRulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetRulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetRules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // Run
		op := &xxx_RunOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RunRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Run(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // Extract
		op := &xxx_ExtractOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExtractRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Extract(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IDataManager
type UnimplementedDataManagerServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedDataManagerServer) GetEnabled(context.Context, *GetEnabledRequest) (*GetEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) SetEnabled(context.Context, *SetEnabledRequest) (*SetEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) GetCheckBeforeRunning(context.Context, *GetCheckBeforeRunningRequest) (*GetCheckBeforeRunningResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) SetCheckBeforeRunning(context.Context, *SetCheckBeforeRunningRequest) (*SetCheckBeforeRunningResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) GetMinFreeDisk(context.Context, *GetMinFreeDiskRequest) (*GetMinFreeDiskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) SetMinFreeDisk(context.Context, *SetMinFreeDiskRequest) (*SetMinFreeDiskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) GetMaxSize(context.Context, *GetMaxSizeRequest) (*GetMaxSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) SetMaxSize(context.Context, *SetMaxSizeRequest) (*SetMaxSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) GetMaxFolderCount(context.Context, *GetMaxFolderCountRequest) (*GetMaxFolderCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) SetMaxFolderCount(context.Context, *SetMaxFolderCountRequest) (*SetMaxFolderCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) GetResourcePolicy(context.Context, *GetResourcePolicyRequest) (*GetResourcePolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) SetResourcePolicy(context.Context, *SetResourcePolicyRequest) (*SetResourcePolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) GetFolderActions(context.Context, *GetFolderActionsRequest) (*GetFolderActionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) GetReportSchema(context.Context, *GetReportSchemaRequest) (*GetReportSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) SetReportSchema(context.Context, *SetReportSchemaRequest) (*SetReportSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) GetReportFileName(context.Context, *GetReportFileNameRequest) (*GetReportFileNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) SetReportFileName(context.Context, *SetReportFileNameRequest) (*SetReportFileNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) GetRuleTargetFileName(context.Context, *GetRuleTargetFileNameRequest) (*GetRuleTargetFileNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) SetRuleTargetFileName(context.Context, *SetRuleTargetFileNameRequest) (*SetRuleTargetFileNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) GetEventsFileName(context.Context, *GetEventsFileNameRequest) (*GetEventsFileNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) SetEventsFileName(context.Context, *SetEventsFileNameRequest) (*SetEventsFileNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) GetRules(context.Context, *GetRulesRequest) (*GetRulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) SetRules(context.Context, *SetRulesRequest) (*SetRulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) Run(context.Context, *RunRequest) (*RunResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataManagerServer) Extract(context.Context, *ExtractRequest) (*ExtractResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DataManagerServer = (*UnimplementedDataManagerServer)(nil)
