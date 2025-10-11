package iwrmconfig

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

// IWRMConfig server interface.
type ConfigServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The GetConfig method gets WSRM configuration information concerning accounting and
	// notifications.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------+----------------------------------------+
	//	|             RETURN             |                                        |
	//	|           VALUE/CODE           |              DESCRIPTION               |
	//	|                                |                                        |
	//	+--------------------------------+----------------------------------------+
	//	+--------------------------------+----------------------------------------+
	//	| 0x00000000 S_OK                | Operation successful.                  |
	//	+--------------------------------+----------------------------------------+
	//	| 0x80070032 ERROR_NOT_SUPPORTED | The requested action is not supported. |
	//	+--------------------------------+----------------------------------------+
	//
	// Additional IWRMConfig interface methods are specified in section 3.2.4.5.
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)

	// The SetConfig method sets WSRM configuration information concerning accounting and
	// notifications.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                     RETURN                      |                                                                                  |
	//	|                   VALUE/CODE                    |                                   DESCRIPTION                                    |
	//	|                                                 |                                                                                  |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                 | Operation successful.                                                            |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED                  | Access is denied.                                                                |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070032 ERROR_NOT_SUPPORTED                  | The requested action is not supported.                                           |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                         | One or more arguments are invalid.                                               |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006E WRM_ERR_TOO_LONG_CONFIG_VALUE        | One or more specified values have exceeded an implementation-defined limit.<72>  |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER            | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                                 | be processed.<73>                                                                |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0190 WRM_ERR_INVALID_NTFY_ENABLE          | The notification-enabled value MUST be Boolean (section 2.2.1.2).                |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0194 WRM_ERR_INVALID_EVENTLIST            | The notification event list format is invalid.                                   |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF01F6 WRM_ERR_INVALID_ACC_ENABLE           | The accounting-enabled value MUST be Boolean (section 2.2.1.2).                  |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF01F9 WRM_ERR_ACC_INVALID_DUMPING_INTERVAL | The logging interval for accounting is invalid. The interval MUST be between 2   |
	//	|                                                 | and 60,000 minutes, inclusive.                                                   |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Note  When the CONFIGTYPE_ACCOUNTING option is used, an accounting client SHOULD
	// call SetClientPermissions (section 3.2.4.3.14) prior to SetConfig in order to obtain
	// authorization to modify an accounting database on a remote WSRM server. If a client
	// does not have permission, ERROR_ACCESS_DENIED SHOULD be returned.<74>
	//
	// Additional IWRMConfig interface methods are specified in section 3.2.4.5.
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)

	// The IsEnabled method returns whether a type of WSRM configuration object is enabled
	// or disabled.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | Operation successful.              |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// Additional IWRMConfig interface methods are specified in section 3.2.4.5.
	IsEnabled(context.Context, *IsEnabledRequest) (*IsEnabledResponse, error)

	// The EnableDisable method enables or disables the management of a specified type of
	// WSRM configuration.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                        RETURN                         |                                                                                  |
	//	|                      VALUE/CODE                       |                                   DESCRIPTION                                    |
	//	|                                                       |                                                                                  |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                       | Operation successful.                                                            |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070032 ERROR_NOT_SUPPORTED                        | The requested action is not supported.<75>                                       |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                               | One or more arguments are invalid.                                               |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x81FF006D WRM_WARNING_CALENDARING_ALREADY_ENABLED    | The management service is already in calendar mode.                              |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x81FF0196 WRM_WARNING_NOTIFICATION_ALREADY_DISABLED  | Notification is already disabled.                                                |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x81FF0197 WRM_WARNING_NOTIFICATION_ALREADY_ENABLED   | Notification is already enabled.                                                 |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x81FF01F4 WRM_WARNING_ACCOUNTING_ALREADY_DISABLED    | Accounting is already disabled.                                                  |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x81FF01F5 WRM_WARNING_ACCOUNTING_ALREADY_ENABLED     | Accounting is already enabled.                                                   |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF007E WRM_ERR_ACCENABLED_UNDER_GROUPPOLICY       | The accounting setting is currently configured under Group Policy and cannot be  |
	//	|                                                       | modified.<76>                                                                    |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0191 WRM_ERR_SMTPSERVERNAME_CANNOT_BE_NULL      | Both the Simple Mail Transfer Protocol (SMTP) server and email name are required |
	//	|                                                       | in order to enable notification.                                                 |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0193 WRM_ERR_EVENTLIST_CANNOT_BE_NULL           | Notification requires at least one event.                                        |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF01F8 WRM_ERR_ACCOUNTING_NOT_CONFIGURED_PROPERLY | Accounting has not been configured properly. The accounting server might not be  |
	//	|                                                       | present or access might be denied.<77>                                           |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF020B WRM_ERR_REMOTE_SERVICE_NOT_STARTED         | Connection to the remote server could not be established because the WSRM        |
	//	|                                                       | management service has not been started on the remote server.<78>                |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMConfig interface methods are specified in section 3.2.4.5.
	EnableDisable(context.Context, *EnableDisableRequest) (*EnableDisableResponse, error)

	// The GetExclusionList method gets the current exclusion list of processes not managed
	// by the WSRM management service.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                      RETURN                      |                                                                                  |
	//	|                    VALUE/CODE                    |                                   DESCRIPTION                                    |
	//	|                                                  |                                                                                  |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                  | Operation successful.                                                            |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                          | One or more arguments are invalid.                                               |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006B WRM_ERR_TOO_LONG_PROCESS_NAME         | The exclusion list could not be returned because the number of characters in a   |
	//	|                                                  | process name has exceeded an implementation-defined limit.<79>                   |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006C WRM_ERR_EXCLUSION_LIST_LIMIT_EXCEEDED | The exclusion list could not be returned because the number of processes that    |
	//	|                                                  | can be excluded has exceeded an implementation-defined limit.<80>                |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMConfig interface methods are specified in section 3.2.4.5.
	GetExclusionList(context.Context, *GetExclusionListRequest) (*GetExclusionListResponse, error)

	// The SetExclusionList method updates the contents of the exclusion list which is of
	// type USER_EXCLUSION_LIST.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                      RETURN                      |                                                                                  |
	//	|                    VALUE/CODE                    |                                   DESCRIPTION                                    |
	//	|                                                  |                                                                                  |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                  | Operation successful.                                                            |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006B WRM_ERR_TOO_LONG_PROCESS_NAME         | The exclusion list could not be returned because the number of characters in a   |
	//	|                                                  | process name has exceeded an implementation-defined limit.<81>.                  |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006C WRM_ERR_EXCLUSION_LIST_LIMIT_EXCEEDED | The exclusion list could not be returned because the number of processes that    |
	//	|                                                  | can be excluded has exceeded an implementation-defined limit.<82>                |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMConfig interface methods are specified in section 3.2.4.5.
	SetExclusionList(context.Context, *SetExclusionListRequest) (*SetExclusionListResponse, error)

	// The WSRMActivate method activates or deactivates WSRM management state.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+---------------------------------------------+------------------------------------+
	//	|                   RETURN                    |                                    |
	//	|                 VALUE/CODE                  |            DESCRIPTION             |
	//	|                                             |                                    |
	//	+---------------------------------------------+------------------------------------+
	//	+---------------------------------------------+------------------------------------+
	//	| 0x00000000 S_OK                             | Operation successful.              |
	//	+---------------------------------------------+------------------------------------+
	//	| 0x81FF0075 WRM_WRN_WSRM_ALREADY_ACTIVATED   | WSRM has already been activated.   |
	//	+---------------------------------------------+------------------------------------+
	//	| 0x81FF0074 WRM_WRN_WSRM_ALREADY_DEACTIVATED | WSRM has already been deactivated. |
	//	+---------------------------------------------+------------------------------------+
	//
	// This method MUST update the registry according to the management state of WSRM. If
	// the WSRM management service later restarts, for example, after a reboot of the operating
	// system, the WSRM management service MUST attempt to restore its prior management
	// state.<83>
	//
	// Additional IWRMConfig interface methods are specified in section 3.2.4.5.
	WSRMActivate(context.Context, *WSRMActivateRequest) (*WSRMActivateResponse, error)

	// The IsWSRMActivated method returns whether WSRM is active.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | Operation successful.              |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// Additional IWRMConfig interface methods are specified in section 3.2.4.5.
	IsWSRMActivated(context.Context, *IsWSRMActivatedRequest) (*IsWSRMActivatedResponse, error)

	// The RestoreExclusionList method restores the exclusion list, which is of type USER_EXCLUSION_LIST,
	// to default values.
	//
	// This method has no parameters.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IWRMConfig interface methods are specified in section 3.2.4.5.
	RestoreExclusionList(context.Context, *RestoreExclusionListRequest) (*RestoreExclusionListResponse, error)
}

func RegisterConfigServer(conn dcerpc.Conn, o ConfigServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewConfigServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ConfigSyntaxV0_0))...)
}

func NewConfigServerHandle(o ConfigServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ConfigServerHandle(ctx, o, opNum, r)
	}
}

func ConfigServerHandle(ctx context.Context, o ConfigServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetConfig
		op := &xxx_GetConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // SetConfig
		op := &xxx_SetConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // IsEnabled
		op := &xxx_IsEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // EnableDisable
		op := &xxx_EnableDisableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnableDisableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnableDisable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // GetExclusionList
		op := &xxx_GetExclusionListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExclusionListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetExclusionList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // SetExclusionList
		op := &xxx_SetExclusionListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetExclusionListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetExclusionList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // WSRMActivate
		op := &xxx_WSRMActivateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WSRMActivateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WSRMActivate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // IsWSRMActivated
		op := &xxx_IsWSRMActivatedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsWSRMActivatedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsWSRMActivated(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // RestoreExclusionList
		op := &xxx_RestoreExclusionListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RestoreExclusionListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RestoreExclusionList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMConfig
type UnimplementedConfigServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedConfigServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) IsEnabled(context.Context, *IsEnabledRequest) (*IsEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) EnableDisable(context.Context, *EnableDisableRequest) (*EnableDisableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) GetExclusionList(context.Context, *GetExclusionListRequest) (*GetExclusionListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) SetExclusionList(context.Context, *SetExclusionListRequest) (*SetExclusionListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) WSRMActivate(context.Context, *WSRMActivateRequest) (*WSRMActivateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) IsWSRMActivated(context.Context, *IsWSRMActivatedRequest) (*IsWSRMActivatedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) RestoreExclusionList(context.Context, *RestoreExclusionListRequest) (*RestoreExclusionListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ConfigServer = (*UnimplementedConfigServer)(nil)
