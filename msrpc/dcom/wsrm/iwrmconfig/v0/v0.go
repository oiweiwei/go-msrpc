package iwrmconfig

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
	wsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
	_ = wsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

var (
	// IWRMConfig interface identifier 21546ae8-4da5-445e-987f-627fea39c5e8
	ConfigIID = &dcom.IID{Data1: 0x21546ae8, Data2: 0x4da5, Data3: 0x445e, Data4: []byte{0x98, 0x7f, 0x62, 0x7f, 0xea, 0x39, 0xc5, 0xe8}}
	// Syntax UUID
	ConfigSyntaxUUID = &uuid.UUID{TimeLow: 0x21546ae8, TimeMid: 0x4da5, TimeHiAndVersion: 0x445e, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x7f, Node: [6]uint8{0x62, 0x7f, 0xea, 0x39, 0xc5, 0xe8}}
	// Syntax ID
	ConfigSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ConfigSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWRMConfig interface.
type ConfigClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

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
	GetConfig(context.Context, *GetConfigRequest, ...dcerpc.CallOption) (*GetConfigResponse, error)

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
	SetConfig(context.Context, *SetConfigRequest, ...dcerpc.CallOption) (*SetConfigResponse, error)

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
	IsEnabled(context.Context, *IsEnabledRequest, ...dcerpc.CallOption) (*IsEnabledResponse, error)

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
	EnableDisable(context.Context, *EnableDisableRequest, ...dcerpc.CallOption) (*EnableDisableResponse, error)

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
	GetExclusionList(context.Context, *GetExclusionListRequest, ...dcerpc.CallOption) (*GetExclusionListResponse, error)

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
	SetExclusionList(context.Context, *SetExclusionListRequest, ...dcerpc.CallOption) (*SetExclusionListResponse, error)

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
	WSRMActivate(context.Context, *WSRMActivateRequest, ...dcerpc.CallOption) (*WSRMActivateResponse, error)

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
	IsWSRMActivated(context.Context, *IsWSRMActivatedRequest, ...dcerpc.CallOption) (*IsWSRMActivatedResponse, error)

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
	RestoreExclusionList(context.Context, *RestoreExclusionListRequest, ...dcerpc.CallOption) (*RestoreExclusionListResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ConfigClient
}

type xxx_DefaultConfigClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultConfigClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultConfigClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...dcerpc.CallOption) (*GetConfigResponse, error) {
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
	out := &GetConfigResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigClient) SetConfig(ctx context.Context, in *SetConfigRequest, opts ...dcerpc.CallOption) (*SetConfigResponse, error) {
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
	out := &SetConfigResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigClient) IsEnabled(ctx context.Context, in *IsEnabledRequest, opts ...dcerpc.CallOption) (*IsEnabledResponse, error) {
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
	out := &IsEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigClient) EnableDisable(ctx context.Context, in *EnableDisableRequest, opts ...dcerpc.CallOption) (*EnableDisableResponse, error) {
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
	out := &EnableDisableResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigClient) GetExclusionList(ctx context.Context, in *GetExclusionListRequest, opts ...dcerpc.CallOption) (*GetExclusionListResponse, error) {
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
	out := &GetExclusionListResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigClient) SetExclusionList(ctx context.Context, in *SetExclusionListRequest, opts ...dcerpc.CallOption) (*SetExclusionListResponse, error) {
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
	out := &SetExclusionListResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigClient) WSRMActivate(ctx context.Context, in *WSRMActivateRequest, opts ...dcerpc.CallOption) (*WSRMActivateResponse, error) {
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
	out := &WSRMActivateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigClient) IsWSRMActivated(ctx context.Context, in *IsWSRMActivatedRequest, opts ...dcerpc.CallOption) (*IsWSRMActivatedResponse, error) {
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
	out := &IsWSRMActivatedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigClient) RestoreExclusionList(ctx context.Context, in *RestoreExclusionListRequest, opts ...dcerpc.CallOption) (*RestoreExclusionListResponse, error) {
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
	out := &RestoreExclusionListResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConfigClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultConfigClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultConfigClient) IPID(ctx context.Context, ipid *dcom.IPID) ConfigClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultConfigClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewConfigClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ConfigClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ConfigSyntaxV0_0))...)
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
	return &xxx_DefaultConfigClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetConfigOperation structure represents the GetConfig operation
type xxx_GetConfigOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ConfigInfo     *oaut.String    `idl:"name:pbstrConfigInfo" json:"config_info"`
	EnumConfigType wsrm.ConfigType `idl:"name:enumConfigType" json:"enum_config_type"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetConfigOperation) OpNum() int { return 7 }

func (o *xxx_GetConfigOperation) OpName() string { return "/IWRMConfig/v0/GetConfig" }

func (o *xxx_GetConfigOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// enumConfigType {in} (1:{v1_enum, alias=CONFIGTYPE}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumConfigType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// enumConfigType {in} (1:{v1_enum, alias=CONFIGTYPE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumConfigType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrConfigInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConfigInfo != nil {
			_ptr_pbstrConfigInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigInfo != nil {
					if err := o.ConfigInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigInfo, _ptr_pbstrConfigInfo); err != nil {
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

func (o *xxx_GetConfigOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrConfigInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrConfigInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigInfo == nil {
				o.ConfigInfo = &oaut.String{}
			}
			if err := o.ConfigInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrConfigInfo := func(ptr interface{}) { o.ConfigInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConfigInfo, _s_pbstrConfigInfo, _ptr_pbstrConfigInfo); err != nil {
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

// GetConfigRequest structure represents the GetConfig operation request
type GetConfigRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// enumConfigType: A CONFIGTYPE enumeration (section 2.2.3.1) value that specifies the
	// type of WSRM configuration information to get.
	//
	// Obtaining calendar information is not supported. If this parameter is set to CONFIGTYPE_CALENDARING,
	// ERROR_NOT_SUPPORTED SHOULD be returned.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|                           |                                                                                  |
	//	|           VALUE           |                                     MEANING                                      |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| CONFIGTYPE_ACCOUNTING 1   | The WSRM configuration information is in the form of an AccountingConfigInfo     |
	//	|                           | element (section 2.2.5.2).                                                       |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| CONFIGTYPE_NOTIFICATION 2 | The WSRM configuration information is in the form of a NotificationConfigInfo    |
	//	|                           | element (section 2.2.5.19).                                                      |
	//	+---------------------------+----------------------------------------------------------------------------------+
	EnumConfigType wsrm.ConfigType `idl:"name:enumConfigType" json:"enum_config_type"`
}

func (o *GetConfigRequest) xxx_ToOp(ctx context.Context, op *xxx_GetConfigOperation) *xxx_GetConfigOperation {
	if op == nil {
		op = &xxx_GetConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.EnumConfigType = o.EnumConfigType
	return op
}

func (o *GetConfigRequest) xxx_FromOp(ctx context.Context, op *xxx_GetConfigOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EnumConfigType = op.EnumConfigType
}
func (o *GetConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetConfigRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetConfigResponse structure represents the GetConfig operation response
type GetConfigResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrConfigInfo: A pointer to a string that returns WSRM configuration information.
	// The type of information is determined by the value of the enumConfigType parameter.
	ConfigInfo *oaut.String `idl:"name:pbstrConfigInfo" json:"config_info"`
	// Return: The GetConfig return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetConfigResponse) xxx_ToOp(ctx context.Context, op *xxx_GetConfigOperation) *xxx_GetConfigOperation {
	if op == nil {
		op = &xxx_GetConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ConfigInfo = o.ConfigInfo
	op.Return = o.Return
	return op
}

func (o *GetConfigResponse) xxx_FromOp(ctx context.Context, op *xxx_GetConfigOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ConfigInfo = op.ConfigInfo
	o.Return = op.Return
}
func (o *GetConfigResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetConfigResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetConfigOperation structure represents the SetConfig operation
type xxx_SetConfigOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ConfigInfo     *oaut.String    `idl:"name:bstrConfigInfo" json:"config_info"`
	EnumConfigType wsrm.ConfigType `idl:"name:enumConfigType" json:"enum_config_type"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetConfigOperation) OpNum() int { return 8 }

func (o *xxx_SetConfigOperation) OpName() string { return "/IWRMConfig/v0/SetConfig" }

func (o *xxx_SetConfigOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConfigOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrConfigInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConfigInfo != nil {
			_ptr_bstrConfigInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigInfo != nil {
					if err := o.ConfigInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigInfo, _ptr_bstrConfigInfo); err != nil {
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
	// enumConfigType {in} (1:{v1_enum, alias=CONFIGTYPE}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumConfigType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConfigOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrConfigInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConfigInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigInfo == nil {
				o.ConfigInfo = &oaut.String{}
			}
			if err := o.ConfigInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConfigInfo := func(ptr interface{}) { o.ConfigInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConfigInfo, _s_bstrConfigInfo, _ptr_bstrConfigInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// enumConfigType {in} (1:{v1_enum, alias=CONFIGTYPE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumConfigType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConfigOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConfigOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetConfigOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetConfigRequest structure represents the SetConfig operation request
type SetConfigRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrConfigInfo: A string that contains WSRM configuration information. The type of
	// information is determined by the value of the enumConfigType parameter.
	ConfigInfo *oaut.String `idl:"name:bstrConfigInfo" json:"config_info"`
	// enumConfigType: A CONFIGTYPE enumeration (section 2.2.3.1) value that specifies the
	// type of WSRM configuration information to set.
	//
	// Setting calendar information is not supported. If this parameter is set to CONFIGTYPE_CALENDARING,
	// ERROR_NOT_SUPPORTED SHOULD be returned.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|                           |                                                                                  |
	//	|           VALUE           |                                     MEANING                                      |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| CONFIGTYPE_ACCOUNTING 1   | The WSRM configuration information is in the form of an AccountingConfigInfo     |
	//	|                           | element (section 2.2.5.2).                                                       |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| CONFIGTYPE_NOTIFICATION 2 | The WSRM configuration information is in the form of an NotificationConfigInfo   |
	//	|                           | element (section 2.2.5.19).                                                      |
	//	+---------------------------+----------------------------------------------------------------------------------+
	EnumConfigType wsrm.ConfigType `idl:"name:enumConfigType" json:"enum_config_type"`
}

func (o *SetConfigRequest) xxx_ToOp(ctx context.Context, op *xxx_SetConfigOperation) *xxx_SetConfigOperation {
	if op == nil {
		op = &xxx_SetConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ConfigInfo = o.ConfigInfo
	op.EnumConfigType = o.EnumConfigType
	return op
}

func (o *SetConfigRequest) xxx_FromOp(ctx context.Context, op *xxx_SetConfigOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConfigInfo = op.ConfigInfo
	o.EnumConfigType = op.EnumConfigType
}
func (o *SetConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetConfigRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetConfigOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetConfigResponse structure represents the SetConfig operation response
type SetConfigResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetConfig return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetConfigResponse) xxx_ToOp(ctx context.Context, op *xxx_SetConfigOperation) *xxx_SetConfigOperation {
	if op == nil {
		op = &xxx_SetConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetConfigResponse) xxx_FromOp(ctx context.Context, op *xxx_SetConfigOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetConfigResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetConfigResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetConfigOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsEnabledOperation structure represents the IsEnabled operation
type xxx_IsEnabledOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Enable         bool            `idl:"name:pbEnable" json:"enable"`
	EnumConfigType wsrm.ConfigType `idl:"name:enumConfigType" json:"enum_config_type"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_IsEnabledOperation) OpNum() int { return 9 }

func (o *xxx_IsEnabledOperation) OpName() string { return "/IWRMConfig/v0/IsEnabled" }

func (o *xxx_IsEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// enumConfigType {in} (1:{v1_enum, alias=CONFIGTYPE}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumConfigType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// enumConfigType {in} (1:{v1_enum, alias=CONFIGTYPE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumConfigType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbEnable {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.Enable {
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

func (o *xxx_IsEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbEnable {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bEnable int32
		if err := w.ReadData(&_bEnable); err != nil {
			return err
		}
		o.Enable = _bEnable != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsEnabledRequest structure represents the IsEnabled operation request
type IsEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// enumConfigType: A CONFIGTYPE enumeration (section 2.2.3.1) value that specifies the
	// type of WSRM configuration object.
	EnumConfigType wsrm.ConfigType `idl:"name:enumConfigType" json:"enum_config_type"`
}

func (o *IsEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_IsEnabledOperation) *xxx_IsEnabledOperation {
	if op == nil {
		op = &xxx_IsEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.EnumConfigType = o.EnumConfigType
	return op
}

func (o *IsEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_IsEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EnumConfigType = op.EnumConfigType
}
func (o *IsEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsEnabledResponse structure represents the IsEnabled operation response
type IsEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbEnable: A pointer to a Boolean value that indicates whether the WSRM configuration
	// object is enabled or disabled.
	Enable bool `idl:"name:pbEnable" json:"enable"`
	// Return: The IsEnabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_IsEnabledOperation) *xxx_IsEnabledOperation {
	if op == nil {
		op = &xxx_IsEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enable = o.Enable
	op.Return = o.Return
	return op
}

func (o *IsEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_IsEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enable = op.Enable
	o.Return = op.Return
}
func (o *IsEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnableDisableOperation structure represents the EnableDisable operation
type xxx_EnableDisableOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	EnableDisable  bool            `idl:"name:bEnableDisable" json:"enable_disable"`
	EnumConfigType wsrm.ConfigType `idl:"name:enumConfigType" json:"enum_config_type"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_EnableDisableOperation) OpNum() int { return 10 }

func (o *xxx_EnableDisableOperation) OpName() string { return "/IWRMConfig/v0/EnableDisable" }

func (o *xxx_EnableDisableOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableDisableOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bEnableDisable {in} (1:{alias=BOOL}(int32))
	{
		if !o.EnableDisable {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// enumConfigType {in} (1:{v1_enum, alias=CONFIGTYPE}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumConfigType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableDisableOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bEnableDisable {in} (1:{alias=BOOL}(int32))
	{
		var _bEnableDisable int32
		if err := w.ReadData(&_bEnableDisable); err != nil {
			return err
		}
		o.EnableDisable = _bEnableDisable != 0
	}
	// enumConfigType {in} (1:{v1_enum, alias=CONFIGTYPE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumConfigType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableDisableOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableDisableOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnableDisableOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// EnableDisableRequest structure represents the EnableDisable operation request
type EnableDisableRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bEnableDisable: A Boolean value that specifies whether to enable or disable the management
	// of the WSRM configuration. The type of configuration is determined by the value of
	// the enumConfigType parameter.
	//
	//	+------------------+--------------------------------------------------------+
	//	|                  |                                                        |
	//	|      VALUE       |                        MEANING                         |
	//	|                  |                                                        |
	//	+------------------+--------------------------------------------------------+
	//	+------------------+--------------------------------------------------------+
	//	| FALSE 0x00000000 | Management of the WSRM configuration MUST be disabled. |
	//	+------------------+--------------------------------------------------------+
	//	| TRUE 0x00000001  | Management of the WSRM configuration MUST be enabled.  |
	//	+------------------+--------------------------------------------------------+
	EnableDisable bool `idl:"name:bEnableDisable" json:"enable_disable"`
	// enumConfigType: A CONFIGTYPE enumeration (section 2.2.3.1) value that specifies the
	// type of WSRM configuration to enable or disable.
	EnumConfigType wsrm.ConfigType `idl:"name:enumConfigType" json:"enum_config_type"`
}

func (o *EnableDisableRequest) xxx_ToOp(ctx context.Context, op *xxx_EnableDisableOperation) *xxx_EnableDisableOperation {
	if op == nil {
		op = &xxx_EnableDisableOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.EnableDisable = o.EnableDisable
	op.EnumConfigType = o.EnumConfigType
	return op
}

func (o *EnableDisableRequest) xxx_FromOp(ctx context.Context, op *xxx_EnableDisableOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EnableDisable = op.EnableDisable
	o.EnumConfigType = op.EnumConfigType
}
func (o *EnableDisableRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnableDisableRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnableDisableOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnableDisableResponse structure represents the EnableDisable operation response
type EnableDisableResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EnableDisable return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnableDisableResponse) xxx_ToOp(ctx context.Context, op *xxx_EnableDisableOperation) *xxx_EnableDisableOperation {
	if op == nil {
		op = &xxx_EnableDisableOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *EnableDisableResponse) xxx_FromOp(ctx context.Context, op *xxx_EnableDisableOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *EnableDisableResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnableDisableResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnableDisableOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetExclusionListOperation structure represents the GetExclusionList operation
type xxx_GetExclusionListOperation struct {
	This          *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat         `idl:"name:That" json:"that"`
	ExclusionList *oaut.String           `idl:"name:pbstrExclusionList" json:"exclusion_list"`
	EnumListType  wsrm.ExclusionlistType `idl:"name:enumListType" json:"enum_list_type"`
	Return        int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetExclusionListOperation) OpNum() int { return 11 }

func (o *xxx_GetExclusionListOperation) OpName() string { return "/IWRMConfig/v0/GetExclusionList" }

func (o *xxx_GetExclusionListOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExclusionListOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// enumListType {in} (1:{v1_enum, alias=EXCLUSIONLIST_TYPE}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumListType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExclusionListOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// enumListType {in} (1:{v1_enum, alias=EXCLUSIONLIST_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumListType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExclusionListOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExclusionListOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrExclusionList {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ExclusionList != nil {
			_ptr_pbstrExclusionList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ExclusionList != nil {
					if err := o.ExclusionList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExclusionList, _ptr_pbstrExclusionList); err != nil {
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

func (o *xxx_GetExclusionListOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrExclusionList {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrExclusionList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ExclusionList == nil {
				o.ExclusionList = &oaut.String{}
			}
			if err := o.ExclusionList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrExclusionList := func(ptr interface{}) { o.ExclusionList = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ExclusionList, _s_pbstrExclusionList, _ptr_pbstrExclusionList); err != nil {
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

// GetExclusionListRequest structure represents the GetExclusionList operation request
type GetExclusionListRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// enumListType: An EXCLUSIONLIST_TYPE enumeration (section 2.2.3.2) value that specifies
	// whether the list is system-defined or user-defined.
	EnumListType wsrm.ExclusionlistType `idl:"name:enumListType" json:"enum_list_type"`
}

func (o *GetExclusionListRequest) xxx_ToOp(ctx context.Context, op *xxx_GetExclusionListOperation) *xxx_GetExclusionListOperation {
	if op == nil {
		op = &xxx_GetExclusionListOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.EnumListType = o.EnumListType
	return op
}

func (o *GetExclusionListRequest) xxx_FromOp(ctx context.Context, op *xxx_GetExclusionListOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EnumListType = op.EnumListType
}
func (o *GetExclusionListRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetExclusionListRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExclusionListOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetExclusionListResponse structure represents the GetExclusionList operation response
type GetExclusionListResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrExclusionList: A pointer to a string that returns the exclusion list, in the
	// form of an ExclusionList element (section 2.2.5.16). For an example, see ExclusionList
	// Example (section 4.2.13).
	ExclusionList *oaut.String `idl:"name:pbstrExclusionList" json:"exclusion_list"`
	// Return: The GetExclusionList return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetExclusionListResponse) xxx_ToOp(ctx context.Context, op *xxx_GetExclusionListOperation) *xxx_GetExclusionListOperation {
	if op == nil {
		op = &xxx_GetExclusionListOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ExclusionList = o.ExclusionList
	op.Return = o.Return
	return op
}

func (o *GetExclusionListResponse) xxx_FromOp(ctx context.Context, op *xxx_GetExclusionListOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ExclusionList = op.ExclusionList
	o.Return = op.Return
}
func (o *GetExclusionListResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetExclusionListResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExclusionListOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetExclusionListOperation structure represents the SetExclusionList operation
type xxx_SetExclusionListOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ExclusionList *oaut.String   `idl:"name:bstrExclusionList" json:"exclusion_list"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetExclusionListOperation) OpNum() int { return 12 }

func (o *xxx_SetExclusionListOperation) OpName() string { return "/IWRMConfig/v0/SetExclusionList" }

func (o *xxx_SetExclusionListOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExclusionListOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrExclusionList {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ExclusionList != nil {
			_ptr_bstrExclusionList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ExclusionList != nil {
					if err := o.ExclusionList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExclusionList, _ptr_bstrExclusionList); err != nil {
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

func (o *xxx_SetExclusionListOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrExclusionList {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrExclusionList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ExclusionList == nil {
				o.ExclusionList = &oaut.String{}
			}
			if err := o.ExclusionList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrExclusionList := func(ptr interface{}) { o.ExclusionList = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ExclusionList, _s_bstrExclusionList, _ptr_bstrExclusionList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExclusionListOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExclusionListOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetExclusionListOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetExclusionListRequest structure represents the SetExclusionList operation request
type SetExclusionListRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrExclusionList: A string that specifies a list of processes, in the form of an
	// ExclusionList element (section 2.2.5.16). For an example, see (section 4.2.13).
	ExclusionList *oaut.String `idl:"name:bstrExclusionList" json:"exclusion_list"`
}

func (o *SetExclusionListRequest) xxx_ToOp(ctx context.Context, op *xxx_SetExclusionListOperation) *xxx_SetExclusionListOperation {
	if op == nil {
		op = &xxx_SetExclusionListOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ExclusionList = o.ExclusionList
	return op
}

func (o *SetExclusionListRequest) xxx_FromOp(ctx context.Context, op *xxx_SetExclusionListOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ExclusionList = op.ExclusionList
}
func (o *SetExclusionListRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetExclusionListRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExclusionListOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetExclusionListResponse structure represents the SetExclusionList operation response
type SetExclusionListResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetExclusionList return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetExclusionListResponse) xxx_ToOp(ctx context.Context, op *xxx_SetExclusionListOperation) *xxx_SetExclusionListOperation {
	if op == nil {
		op = &xxx_SetExclusionListOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetExclusionListResponse) xxx_FromOp(ctx context.Context, op *xxx_SetExclusionListOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetExclusionListResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetExclusionListResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExclusionListOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WSRMActivateOperation structure represents the WSRMActivate operation
type xxx_WSRMActivateOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Activate bool           `idl:"name:bActivate" json:"activate"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_WSRMActivateOperation) OpNum() int { return 13 }

func (o *xxx_WSRMActivateOperation) OpName() string { return "/IWRMConfig/v0/WSRMActivate" }

func (o *xxx_WSRMActivateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WSRMActivateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bActivate {in} (1:{alias=BOOL}(int32))
	{
		if !o.Activate {
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

func (o *xxx_WSRMActivateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bActivate {in} (1:{alias=BOOL}(int32))
	{
		var _bActivate int32
		if err := w.ReadData(&_bActivate); err != nil {
			return err
		}
		o.Activate = _bActivate != 0
	}
	return nil
}

func (o *xxx_WSRMActivateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WSRMActivateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_WSRMActivateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// WSRMActivateRequest structure represents the WSRMActivate operation request
type WSRMActivateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bActivate: A Boolean value that specifies whether to activate or deactivate WSRM.
	//
	//	+------------------+---------------------------------------------------------------------------------+
	//	|                  |                                                                                 |
	//	|      VALUE       |                                     MEANING                                     |
	//	|                  |                                                                                 |
	//	+------------------+---------------------------------------------------------------------------------+
	//	+------------------+---------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | Management state of WSRM management service MUST be set to inactive or stopped. |
	//	+------------------+---------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | Management state of WSRM management service MUST be set to active or running.   |
	//	+------------------+---------------------------------------------------------------------------------+
	Activate bool `idl:"name:bActivate" json:"activate"`
}

func (o *WSRMActivateRequest) xxx_ToOp(ctx context.Context, op *xxx_WSRMActivateOperation) *xxx_WSRMActivateOperation {
	if op == nil {
		op = &xxx_WSRMActivateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Activate = o.Activate
	return op
}

func (o *WSRMActivateRequest) xxx_FromOp(ctx context.Context, op *xxx_WSRMActivateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Activate = op.Activate
}
func (o *WSRMActivateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WSRMActivateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WSRMActivateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WSRMActivateResponse structure represents the WSRMActivate operation response
type WSRMActivateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The WSRMActivate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WSRMActivateResponse) xxx_ToOp(ctx context.Context, op *xxx_WSRMActivateOperation) *xxx_WSRMActivateOperation {
	if op == nil {
		op = &xxx_WSRMActivateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *WSRMActivateResponse) xxx_FromOp(ctx context.Context, op *xxx_WSRMActivateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *WSRMActivateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WSRMActivateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WSRMActivateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsWSRMActivatedOperation structure represents the IsWSRMActivated operation
type xxx_IsWSRMActivatedOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Activated bool           `idl:"name:pbActivated" json:"activated"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsWSRMActivatedOperation) OpNum() int { return 14 }

func (o *xxx_IsWSRMActivatedOperation) OpName() string { return "/IWRMConfig/v0/IsWSRMActivated" }

func (o *xxx_IsWSRMActivatedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsWSRMActivatedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsWSRMActivatedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_IsWSRMActivatedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsWSRMActivatedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbActivated {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.Activated {
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

func (o *xxx_IsWSRMActivatedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbActivated {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bActivated int32
		if err := w.ReadData(&_bActivated); err != nil {
			return err
		}
		o.Activated = _bActivated != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsWSRMActivatedRequest structure represents the IsWSRMActivated operation request
type IsWSRMActivatedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *IsWSRMActivatedRequest) xxx_ToOp(ctx context.Context, op *xxx_IsWSRMActivatedOperation) *xxx_IsWSRMActivatedOperation {
	if op == nil {
		op = &xxx_IsWSRMActivatedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *IsWSRMActivatedRequest) xxx_FromOp(ctx context.Context, op *xxx_IsWSRMActivatedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *IsWSRMActivatedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsWSRMActivatedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsWSRMActivatedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsWSRMActivatedResponse structure represents the IsWSRMActivated operation response
type IsWSRMActivatedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbActivated: A pointer to a Boolean value that returns whether the WSRM server is
	// in the active management state. If TRUE, WSRM is active; otherwise, it is inactive.
	Activated bool `idl:"name:pbActivated" json:"activated"`
	// Return: The IsWSRMActivated return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsWSRMActivatedResponse) xxx_ToOp(ctx context.Context, op *xxx_IsWSRMActivatedOperation) *xxx_IsWSRMActivatedOperation {
	if op == nil {
		op = &xxx_IsWSRMActivatedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Activated = o.Activated
	op.Return = o.Return
	return op
}

func (o *IsWSRMActivatedResponse) xxx_FromOp(ctx context.Context, op *xxx_IsWSRMActivatedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Activated = op.Activated
	o.Return = op.Return
}
func (o *IsWSRMActivatedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsWSRMActivatedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsWSRMActivatedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RestoreExclusionListOperation structure represents the RestoreExclusionList operation
type xxx_RestoreExclusionListOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RestoreExclusionListOperation) OpNum() int { return 15 }

func (o *xxx_RestoreExclusionListOperation) OpName() string {
	return "/IWRMConfig/v0/RestoreExclusionList"
}

func (o *xxx_RestoreExclusionListOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreExclusionListOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RestoreExclusionListOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_RestoreExclusionListOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreExclusionListOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RestoreExclusionListOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RestoreExclusionListRequest structure represents the RestoreExclusionList operation request
type RestoreExclusionListRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RestoreExclusionListRequest) xxx_ToOp(ctx context.Context, op *xxx_RestoreExclusionListOperation) *xxx_RestoreExclusionListOperation {
	if op == nil {
		op = &xxx_RestoreExclusionListOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *RestoreExclusionListRequest) xxx_FromOp(ctx context.Context, op *xxx_RestoreExclusionListOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RestoreExclusionListRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RestoreExclusionListRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreExclusionListOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RestoreExclusionListResponse structure represents the RestoreExclusionList operation response
type RestoreExclusionListResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RestoreExclusionList return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RestoreExclusionListResponse) xxx_ToOp(ctx context.Context, op *xxx_RestoreExclusionListOperation) *xxx_RestoreExclusionListOperation {
	if op == nil {
		op = &xxx_RestoreExclusionListOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RestoreExclusionListResponse) xxx_FromOp(ctx context.Context, op *xxx_RestoreExclusionListOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RestoreExclusionListResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RestoreExclusionListResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreExclusionListOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
