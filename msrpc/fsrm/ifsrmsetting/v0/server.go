package ifsrmsetting

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

// IFsrmSetting server interface.
type SettingServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// SmtpServer operation.
	GetSMTPServer(context.Context, *GetSMTPServerRequest) (*GetSMTPServerResponse, error)

	// SmtpServer operation.
	SetSMTPServer(context.Context, *SetSMTPServerRequest) (*SetSMTPServerResponse, error)

	// MailFrom operation.
	GetMailFrom(context.Context, *GetMailFromRequest) (*GetMailFromResponse, error)

	// MailFrom operation.
	SetMailFrom(context.Context, *SetMailFromRequest) (*SetMailFromResponse, error)

	// AdminEmail operation.
	GetAdminEmail(context.Context, *GetAdminEmailRequest) (*GetAdminEmailResponse, error)

	// AdminEmail operation.
	SetAdminEmail(context.Context, *SetAdminEmailRequest) (*SetAdminEmailResponse, error)

	// DisableCommandLine operation.
	GetDisableCommandLine(context.Context, *GetDisableCommandLineRequest) (*GetDisableCommandLineResponse, error)

	// DisableCommandLine operation.
	SetDisableCommandLine(context.Context, *SetDisableCommandLineRequest) (*SetDisableCommandLineResponse, error)

	// EnableScreeningAudit operation.
	GetEnableScreeningAudit(context.Context, *GetEnableScreeningAuditRequest) (*GetEnableScreeningAuditResponse, error)

	// EnableScreeningAudit operation.
	SetEnableScreeningAudit(context.Context, *SetEnableScreeningAuditRequest) (*SetEnableScreeningAuditResponse, error)

	// The EmailTest method sends an email message to the specified email address using
	// the settings that the File Server Resource Manager Protocol is configured to use.
	// The settings include SMTP server name and Mail from email address. The format of
	// the email address has to be as specified in [RFC5322].
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE        | The content of the mailTo parameter exceeds the maximum length of 4,000          |
	//	|                                       | characters.                                                                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045318 FSRM_E_INVALID_SMTP_SERVER | The SmtpServer property is not set.                                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004531C FSRM_E_EMAIL_NOT_SENT      | An email message could not be sent.                                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	EmailTest(context.Context, *EmailTestRequest) (*EmailTestResponse, error)

	// The SetActionRunLimitInterval method sets run limit intervals for actions that are
	// configured to use the general setting's run limit interval.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The actionType parameter is not a valid type. If actionType is                   |
	//	|                         | FsrmActionType_Unknown, the parameter MUST be considered an invalid value.       |
	//	+-------------------------+----------------------------------------------------------------------------------+
	SetActionRunLimitInterval(context.Context, *SetActionRunLimitIntervalRequest) (*SetActionRunLimitIntervalResponse, error)

	// The GetActionRunLimitInterval method returns the Run limit interval for actions that
	// are configured to use the general setting's Run limit interval.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The actionType parameter is     |
	//	|                         | not a valid type. If actionType is FsrmActionType_Unknown, the parameter MUST be |
	//	|                         | considered an invalid value. The delayTimeMinutes parameter is NULL.             |
	//	+-------------------------+----------------------------------------------------------------------------------+
	GetActionRunLimitInterval(context.Context, *GetActionRunLimitIntervalRequest) (*GetActionRunLimitIntervalResponse, error)
}

func RegisterSettingServer(conn dcerpc.Conn, o SettingServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewSettingServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(SettingSyntaxV0_0))...)
}

func NewSettingServerHandle(o SettingServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return SettingServerHandle(ctx, o, opNum, r)
	}
}

func SettingServerHandle(ctx context.Context, o SettingServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // SmtpServer
		op := &xxx_GetSMTPServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSMTPServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSMTPServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // SmtpServer
		op := &xxx_SetSMTPServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSMTPServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSMTPServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // MailFrom
		op := &xxx_GetMailFromOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMailFromRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMailFrom(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // MailFrom
		op := &xxx_SetMailFromOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMailFromRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMailFrom(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // AdminEmail
		op := &xxx_GetAdminEmailOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAdminEmailRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAdminEmail(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // AdminEmail
		op := &xxx_SetAdminEmailOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAdminEmailRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAdminEmail(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // DisableCommandLine
		op := &xxx_GetDisableCommandLineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDisableCommandLineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDisableCommandLine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // DisableCommandLine
		op := &xxx_SetDisableCommandLineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDisableCommandLineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDisableCommandLine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // EnableScreeningAudit
		op := &xxx_GetEnableScreeningAuditOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEnableScreeningAuditRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEnableScreeningAudit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // EnableScreeningAudit
		op := &xxx_SetEnableScreeningAuditOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEnableScreeningAuditRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEnableScreeningAudit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // EmailTest
		op := &xxx_EmailTestOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EmailTestRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EmailTest(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // SetActionRunLimitInterval
		op := &xxx_SetActionRunLimitIntervalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetActionRunLimitIntervalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetActionRunLimitInterval(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // GetActionRunLimitInterval
		op := &xxx_GetActionRunLimitIntervalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetActionRunLimitIntervalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetActionRunLimitInterval(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmSetting
type UnimplementedSettingServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedSettingServer) GetSMTPServer(context.Context, *GetSMTPServerRequest) (*GetSMTPServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSettingServer) SetSMTPServer(context.Context, *SetSMTPServerRequest) (*SetSMTPServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSettingServer) GetMailFrom(context.Context, *GetMailFromRequest) (*GetMailFromResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSettingServer) SetMailFrom(context.Context, *SetMailFromRequest) (*SetMailFromResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSettingServer) GetAdminEmail(context.Context, *GetAdminEmailRequest) (*GetAdminEmailResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSettingServer) SetAdminEmail(context.Context, *SetAdminEmailRequest) (*SetAdminEmailResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSettingServer) GetDisableCommandLine(context.Context, *GetDisableCommandLineRequest) (*GetDisableCommandLineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSettingServer) SetDisableCommandLine(context.Context, *SetDisableCommandLineRequest) (*SetDisableCommandLineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSettingServer) GetEnableScreeningAudit(context.Context, *GetEnableScreeningAuditRequest) (*GetEnableScreeningAuditResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSettingServer) SetEnableScreeningAudit(context.Context, *SetEnableScreeningAuditRequest) (*SetEnableScreeningAuditResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSettingServer) EmailTest(context.Context, *EmailTestRequest) (*EmailTestResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSettingServer) SetActionRunLimitInterval(context.Context, *SetActionRunLimitIntervalRequest) (*SetActionRunLimitIntervalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSettingServer) GetActionRunLimitInterval(context.Context, *GetActionRunLimitIntervalRequest) (*GetActionRunLimitIntervalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ SettingServer = (*UnimplementedSettingServer)(nil)
