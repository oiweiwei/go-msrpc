package ifsrmsetting

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
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
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmSetting interface identifier f411d4fd-14be-4260-8c40-03b7c95e608a
	SettingIID = &dcom.IID{Data1: 0xf411d4fd, Data2: 0x14be, Data3: 0x4260, Data4: []byte{0x8c, 0x40, 0x03, 0xb7, 0xc9, 0x5e, 0x60, 0x8a}}
	// Syntax UUID
	SettingSyntaxUUID = &uuid.UUID{TimeLow: 0xf411d4fd, TimeMid: 0x14be, TimeHiAndVersion: 0x4260, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x40, Node: [6]uint8{0x3, 0xb7, 0xc9, 0x5e, 0x60, 0x8a}}
	// Syntax ID
	SettingSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: SettingSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmSetting interface.
type SettingClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// SmtpServer operation.
	GetSMTPServer(context.Context, *GetSMTPServerRequest, ...dcerpc.CallOption) (*GetSMTPServerResponse, error)

	// SmtpServer operation.
	SetSMTPServer(context.Context, *SetSMTPServerRequest, ...dcerpc.CallOption) (*SetSMTPServerResponse, error)

	// MailFrom operation.
	GetMailFrom(context.Context, *GetMailFromRequest, ...dcerpc.CallOption) (*GetMailFromResponse, error)

	// MailFrom operation.
	SetMailFrom(context.Context, *SetMailFromRequest, ...dcerpc.CallOption) (*SetMailFromResponse, error)

	// AdminEmail operation.
	GetAdminEmail(context.Context, *GetAdminEmailRequest, ...dcerpc.CallOption) (*GetAdminEmailResponse, error)

	// AdminEmail operation.
	SetAdminEmail(context.Context, *SetAdminEmailRequest, ...dcerpc.CallOption) (*SetAdminEmailResponse, error)

	// DisableCommandLine operation.
	GetDisableCommandLine(context.Context, *GetDisableCommandLineRequest, ...dcerpc.CallOption) (*GetDisableCommandLineResponse, error)

	// DisableCommandLine operation.
	SetDisableCommandLine(context.Context, *SetDisableCommandLineRequest, ...dcerpc.CallOption) (*SetDisableCommandLineResponse, error)

	// EnableScreeningAudit operation.
	GetEnableScreeningAudit(context.Context, *GetEnableScreeningAuditRequest, ...dcerpc.CallOption) (*GetEnableScreeningAuditResponse, error)

	// EnableScreeningAudit operation.
	SetEnableScreeningAudit(context.Context, *SetEnableScreeningAuditRequest, ...dcerpc.CallOption) (*SetEnableScreeningAuditResponse, error)

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
	EmailTest(context.Context, *EmailTestRequest, ...dcerpc.CallOption) (*EmailTestResponse, error)

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
	SetActionRunLimitInterval(context.Context, *SetActionRunLimitIntervalRequest, ...dcerpc.CallOption) (*SetActionRunLimitIntervalResponse, error)

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
	GetActionRunLimitInterval(context.Context, *GetActionRunLimitIntervalRequest, ...dcerpc.CallOption) (*GetActionRunLimitIntervalResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) SettingClient
}

type xxx_DefaultSettingClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultSettingClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultSettingClient) GetSMTPServer(ctx context.Context, in *GetSMTPServerRequest, opts ...dcerpc.CallOption) (*GetSMTPServerResponse, error) {
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
	out := &GetSMTPServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) SetSMTPServer(ctx context.Context, in *SetSMTPServerRequest, opts ...dcerpc.CallOption) (*SetSMTPServerResponse, error) {
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
	out := &SetSMTPServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) GetMailFrom(ctx context.Context, in *GetMailFromRequest, opts ...dcerpc.CallOption) (*GetMailFromResponse, error) {
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
	out := &GetMailFromResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) SetMailFrom(ctx context.Context, in *SetMailFromRequest, opts ...dcerpc.CallOption) (*SetMailFromResponse, error) {
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
	out := &SetMailFromResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) GetAdminEmail(ctx context.Context, in *GetAdminEmailRequest, opts ...dcerpc.CallOption) (*GetAdminEmailResponse, error) {
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
	out := &GetAdminEmailResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) SetAdminEmail(ctx context.Context, in *SetAdminEmailRequest, opts ...dcerpc.CallOption) (*SetAdminEmailResponse, error) {
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
	out := &SetAdminEmailResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) GetDisableCommandLine(ctx context.Context, in *GetDisableCommandLineRequest, opts ...dcerpc.CallOption) (*GetDisableCommandLineResponse, error) {
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
	out := &GetDisableCommandLineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) SetDisableCommandLine(ctx context.Context, in *SetDisableCommandLineRequest, opts ...dcerpc.CallOption) (*SetDisableCommandLineResponse, error) {
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
	out := &SetDisableCommandLineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) GetEnableScreeningAudit(ctx context.Context, in *GetEnableScreeningAuditRequest, opts ...dcerpc.CallOption) (*GetEnableScreeningAuditResponse, error) {
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
	out := &GetEnableScreeningAuditResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) SetEnableScreeningAudit(ctx context.Context, in *SetEnableScreeningAuditRequest, opts ...dcerpc.CallOption) (*SetEnableScreeningAuditResponse, error) {
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
	out := &SetEnableScreeningAuditResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) EmailTest(ctx context.Context, in *EmailTestRequest, opts ...dcerpc.CallOption) (*EmailTestResponse, error) {
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
	out := &EmailTestResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) SetActionRunLimitInterval(ctx context.Context, in *SetActionRunLimitIntervalRequest, opts ...dcerpc.CallOption) (*SetActionRunLimitIntervalResponse, error) {
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
	out := &SetActionRunLimitIntervalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) GetActionRunLimitInterval(ctx context.Context, in *GetActionRunLimitIntervalRequest, opts ...dcerpc.CallOption) (*GetActionRunLimitIntervalResponse, error) {
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
	out := &GetActionRunLimitIntervalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSettingClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultSettingClient) IPID(ctx context.Context, ipid *dcom.IPID) SettingClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultSettingClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}
func NewSettingClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (SettingClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(SettingSyntaxV0_0))...)
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
	return &xxx_DefaultSettingClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetSMTPServerOperation structure represents the SmtpServer operation
type xxx_GetSMTPServerOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	SMTPServer *oaut.String   `idl:"name:smtpServer" json:"smtp_server"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSMTPServerOperation) OpNum() int { return 7 }

func (o *xxx_GetSMTPServerOperation) OpName() string { return "/IFsrmSetting/v0/SmtpServer" }

func (o *xxx_GetSMTPServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSMTPServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSMTPServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSMTPServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSMTPServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// smtpServer {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SMTPServer != nil {
			_ptr_smtpServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SMTPServer != nil {
					if err := o.SMTPServer.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SMTPServer, _ptr_smtpServer); err != nil {
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

func (o *xxx_GetSMTPServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// smtpServer {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_smtpServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SMTPServer == nil {
				o.SMTPServer = &oaut.String{}
			}
			if err := o.SMTPServer.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_smtpServer := func(ptr interface{}) { o.SMTPServer = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SMTPServer, _s_smtpServer, _ptr_smtpServer); err != nil {
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

// GetSMTPServerRequest structure represents the SmtpServer operation request
type GetSMTPServerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSMTPServerRequest) xxx_ToOp(ctx context.Context) *xxx_GetSMTPServerOperation {
	if o == nil {
		return &xxx_GetSMTPServerOperation{}
	}
	return &xxx_GetSMTPServerOperation{
		This: o.This,
	}
}

func (o *GetSMTPServerRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSMTPServerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSMTPServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSMTPServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSMTPServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSMTPServerResponse structure represents the SmtpServer operation response
type GetSMTPServerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	SMTPServer *oaut.String   `idl:"name:smtpServer" json:"smtp_server"`
	// Return: The SmtpServer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSMTPServerResponse) xxx_ToOp(ctx context.Context) *xxx_GetSMTPServerOperation {
	if o == nil {
		return &xxx_GetSMTPServerOperation{}
	}
	return &xxx_GetSMTPServerOperation{
		That:       o.That,
		SMTPServer: o.SMTPServer,
		Return:     o.Return,
	}
}

func (o *GetSMTPServerResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSMTPServerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SMTPServer = op.SMTPServer
	o.Return = op.Return
}
func (o *GetSMTPServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSMTPServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSMTPServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSMTPServerOperation structure represents the SmtpServer operation
type xxx_SetSMTPServerOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	SMTPServer *oaut.String   `idl:"name:smtpServer" json:"smtp_server"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSMTPServerOperation) OpNum() int { return 8 }

func (o *xxx_SetSMTPServerOperation) OpName() string { return "/IFsrmSetting/v0/SmtpServer" }

func (o *xxx_SetSMTPServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSMTPServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// smtpServer {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SMTPServer != nil {
			_ptr_smtpServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SMTPServer != nil {
					if err := o.SMTPServer.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SMTPServer, _ptr_smtpServer); err != nil {
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

func (o *xxx_SetSMTPServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// smtpServer {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_smtpServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SMTPServer == nil {
				o.SMTPServer = &oaut.String{}
			}
			if err := o.SMTPServer.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_smtpServer := func(ptr interface{}) { o.SMTPServer = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SMTPServer, _s_smtpServer, _ptr_smtpServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSMTPServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSMTPServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSMTPServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSMTPServerRequest structure represents the SmtpServer operation request
type SetSMTPServerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	SMTPServer *oaut.String   `idl:"name:smtpServer" json:"smtp_server"`
}

func (o *SetSMTPServerRequest) xxx_ToOp(ctx context.Context) *xxx_SetSMTPServerOperation {
	if o == nil {
		return &xxx_SetSMTPServerOperation{}
	}
	return &xxx_SetSMTPServerOperation{
		This:       o.This,
		SMTPServer: o.SMTPServer,
	}
}

func (o *SetSMTPServerRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSMTPServerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SMTPServer = op.SMTPServer
}
func (o *SetSMTPServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSMTPServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSMTPServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSMTPServerResponse structure represents the SmtpServer operation response
type SetSMTPServerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SmtpServer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSMTPServerResponse) xxx_ToOp(ctx context.Context) *xxx_SetSMTPServerOperation {
	if o == nil {
		return &xxx_SetSMTPServerOperation{}
	}
	return &xxx_SetSMTPServerOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSMTPServerResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSMTPServerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSMTPServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSMTPServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSMTPServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMailFromOperation structure represents the MailFrom operation
type xxx_GetMailFromOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailFrom *oaut.String   `idl:"name:mailFrom" json:"mail_from"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMailFromOperation) OpNum() int { return 9 }

func (o *xxx_GetMailFromOperation) OpName() string { return "/IFsrmSetting/v0/MailFrom" }

func (o *xxx_GetMailFromOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailFromOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMailFromOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMailFromOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailFromOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// mailFrom {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailFrom != nil {
			_ptr_mailFrom := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailFrom != nil {
					if err := o.MailFrom.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailFrom, _ptr_mailFrom); err != nil {
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

func (o *xxx_GetMailFromOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// mailFrom {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailFrom := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailFrom == nil {
				o.MailFrom = &oaut.String{}
			}
			if err := o.MailFrom.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailFrom := func(ptr interface{}) { o.MailFrom = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailFrom, _s_mailFrom, _ptr_mailFrom); err != nil {
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

// GetMailFromRequest structure represents the MailFrom operation request
type GetMailFromRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMailFromRequest) xxx_ToOp(ctx context.Context) *xxx_GetMailFromOperation {
	if o == nil {
		return &xxx_GetMailFromOperation{}
	}
	return &xxx_GetMailFromOperation{
		This: o.This,
	}
}

func (o *GetMailFromRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMailFromOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMailFromRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMailFromRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailFromOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMailFromResponse structure represents the MailFrom operation response
type GetMailFromResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailFrom *oaut.String   `idl:"name:mailFrom" json:"mail_from"`
	// Return: The MailFrom return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMailFromResponse) xxx_ToOp(ctx context.Context) *xxx_GetMailFromOperation {
	if o == nil {
		return &xxx_GetMailFromOperation{}
	}
	return &xxx_GetMailFromOperation{
		That:     o.That,
		MailFrom: o.MailFrom,
		Return:   o.Return,
	}
}

func (o *GetMailFromResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMailFromOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MailFrom = op.MailFrom
	o.Return = op.Return
}
func (o *GetMailFromResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMailFromResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailFromOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMailFromOperation structure represents the MailFrom operation
type xxx_SetMailFromOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailFrom *oaut.String   `idl:"name:mailFrom" json:"mail_from"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMailFromOperation) OpNum() int { return 10 }

func (o *xxx_SetMailFromOperation) OpName() string { return "/IFsrmSetting/v0/MailFrom" }

func (o *xxx_SetMailFromOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailFromOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// mailFrom {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailFrom != nil {
			_ptr_mailFrom := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailFrom != nil {
					if err := o.MailFrom.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailFrom, _ptr_mailFrom); err != nil {
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

func (o *xxx_SetMailFromOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// mailFrom {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailFrom := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailFrom == nil {
				o.MailFrom = &oaut.String{}
			}
			if err := o.MailFrom.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailFrom := func(ptr interface{}) { o.MailFrom = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailFrom, _s_mailFrom, _ptr_mailFrom); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailFromOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailFromOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMailFromOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMailFromRequest structure represents the MailFrom operation request
type SetMailFromRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	MailFrom *oaut.String   `idl:"name:mailFrom" json:"mail_from"`
}

func (o *SetMailFromRequest) xxx_ToOp(ctx context.Context) *xxx_SetMailFromOperation {
	if o == nil {
		return &xxx_SetMailFromOperation{}
	}
	return &xxx_SetMailFromOperation{
		This:     o.This,
		MailFrom: o.MailFrom,
	}
}

func (o *SetMailFromRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMailFromOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MailFrom = op.MailFrom
}
func (o *SetMailFromRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetMailFromRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailFromOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMailFromResponse structure represents the MailFrom operation response
type SetMailFromResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MailFrom return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMailFromResponse) xxx_ToOp(ctx context.Context) *xxx_SetMailFromOperation {
	if o == nil {
		return &xxx_SetMailFromOperation{}
	}
	return &xxx_SetMailFromOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetMailFromResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMailFromOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMailFromResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetMailFromResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailFromOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAdminEmailOperation structure represents the AdminEmail operation
type xxx_GetAdminEmailOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	AdminEmail *oaut.String   `idl:"name:adminEmail" json:"admin_email"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAdminEmailOperation) OpNum() int { return 11 }

func (o *xxx_GetAdminEmailOperation) OpName() string { return "/IFsrmSetting/v0/AdminEmail" }

func (o *xxx_GetAdminEmailOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminEmailOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAdminEmailOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAdminEmailOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminEmailOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// adminEmail {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.AdminEmail != nil {
			_ptr_adminEmail := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AdminEmail != nil {
					if err := o.AdminEmail.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AdminEmail, _ptr_adminEmail); err != nil {
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

func (o *xxx_GetAdminEmailOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// adminEmail {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_adminEmail := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AdminEmail == nil {
				o.AdminEmail = &oaut.String{}
			}
			if err := o.AdminEmail.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_adminEmail := func(ptr interface{}) { o.AdminEmail = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.AdminEmail, _s_adminEmail, _ptr_adminEmail); err != nil {
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

// GetAdminEmailRequest structure represents the AdminEmail operation request
type GetAdminEmailRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAdminEmailRequest) xxx_ToOp(ctx context.Context) *xxx_GetAdminEmailOperation {
	if o == nil {
		return &xxx_GetAdminEmailOperation{}
	}
	return &xxx_GetAdminEmailOperation{
		This: o.This,
	}
}

func (o *GetAdminEmailRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAdminEmailOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAdminEmailRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAdminEmailRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminEmailOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAdminEmailResponse structure represents the AdminEmail operation response
type GetAdminEmailResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	AdminEmail *oaut.String   `idl:"name:adminEmail" json:"admin_email"`
	// Return: The AdminEmail return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAdminEmailResponse) xxx_ToOp(ctx context.Context) *xxx_GetAdminEmailOperation {
	if o == nil {
		return &xxx_GetAdminEmailOperation{}
	}
	return &xxx_GetAdminEmailOperation{
		That:       o.That,
		AdminEmail: o.AdminEmail,
		Return:     o.Return,
	}
}

func (o *GetAdminEmailResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAdminEmailOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AdminEmail = op.AdminEmail
	o.Return = op.Return
}
func (o *GetAdminEmailResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAdminEmailResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminEmailOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAdminEmailOperation structure represents the AdminEmail operation
type xxx_SetAdminEmailOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	AdminEmail *oaut.String   `idl:"name:adminEmail" json:"admin_email"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAdminEmailOperation) OpNum() int { return 12 }

func (o *xxx_SetAdminEmailOperation) OpName() string { return "/IFsrmSetting/v0/AdminEmail" }

func (o *xxx_SetAdminEmailOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAdminEmailOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// adminEmail {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.AdminEmail != nil {
			_ptr_adminEmail := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AdminEmail != nil {
					if err := o.AdminEmail.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AdminEmail, _ptr_adminEmail); err != nil {
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

func (o *xxx_SetAdminEmailOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// adminEmail {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_adminEmail := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AdminEmail == nil {
				o.AdminEmail = &oaut.String{}
			}
			if err := o.AdminEmail.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_adminEmail := func(ptr interface{}) { o.AdminEmail = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.AdminEmail, _s_adminEmail, _ptr_adminEmail); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAdminEmailOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAdminEmailOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAdminEmailOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAdminEmailRequest structure represents the AdminEmail operation request
type SetAdminEmailRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	AdminEmail *oaut.String   `idl:"name:adminEmail" json:"admin_email"`
}

func (o *SetAdminEmailRequest) xxx_ToOp(ctx context.Context) *xxx_SetAdminEmailOperation {
	if o == nil {
		return &xxx_SetAdminEmailOperation{}
	}
	return &xxx_SetAdminEmailOperation{
		This:       o.This,
		AdminEmail: o.AdminEmail,
	}
}

func (o *SetAdminEmailRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAdminEmailOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AdminEmail = op.AdminEmail
}
func (o *SetAdminEmailRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetAdminEmailRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAdminEmailOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAdminEmailResponse structure represents the AdminEmail operation response
type SetAdminEmailResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AdminEmail return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAdminEmailResponse) xxx_ToOp(ctx context.Context) *xxx_SetAdminEmailOperation {
	if o == nil {
		return &xxx_SetAdminEmailOperation{}
	}
	return &xxx_SetAdminEmailOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetAdminEmailResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAdminEmailOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAdminEmailResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetAdminEmailResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAdminEmailOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDisableCommandLineOperation structure represents the DisableCommandLine operation
type xxx_GetDisableCommandLineOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	DisableCommandLine int16          `idl:"name:disableCommandLine" json:"disable_command_line"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDisableCommandLineOperation) OpNum() int { return 13 }

func (o *xxx_GetDisableCommandLineOperation) OpName() string {
	return "/IFsrmSetting/v0/DisableCommandLine"
}

func (o *xxx_GetDisableCommandLineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisableCommandLineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDisableCommandLineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDisableCommandLineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisableCommandLineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// disableCommandLine {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.DisableCommandLine); err != nil {
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

func (o *xxx_GetDisableCommandLineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// disableCommandLine {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.DisableCommandLine); err != nil {
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

// GetDisableCommandLineRequest structure represents the DisableCommandLine operation request
type GetDisableCommandLineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDisableCommandLineRequest) xxx_ToOp(ctx context.Context) *xxx_GetDisableCommandLineOperation {
	if o == nil {
		return &xxx_GetDisableCommandLineOperation{}
	}
	return &xxx_GetDisableCommandLineOperation{
		This: o.This,
	}
}

func (o *GetDisableCommandLineRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDisableCommandLineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDisableCommandLineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDisableCommandLineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisableCommandLineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDisableCommandLineResponse structure represents the DisableCommandLine operation response
type GetDisableCommandLineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	DisableCommandLine int16          `idl:"name:disableCommandLine" json:"disable_command_line"`
	// Return: The DisableCommandLine return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDisableCommandLineResponse) xxx_ToOp(ctx context.Context) *xxx_GetDisableCommandLineOperation {
	if o == nil {
		return &xxx_GetDisableCommandLineOperation{}
	}
	return &xxx_GetDisableCommandLineOperation{
		That:               o.That,
		DisableCommandLine: o.DisableCommandLine,
		Return:             o.Return,
	}
}

func (o *GetDisableCommandLineResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDisableCommandLineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DisableCommandLine = op.DisableCommandLine
	o.Return = op.Return
}
func (o *GetDisableCommandLineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDisableCommandLineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisableCommandLineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDisableCommandLineOperation structure represents the DisableCommandLine operation
type xxx_SetDisableCommandLineOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	DisableCommandLine int16          `idl:"name:disableCommandLine" json:"disable_command_line"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDisableCommandLineOperation) OpNum() int { return 14 }

func (o *xxx_SetDisableCommandLineOperation) OpName() string {
	return "/IFsrmSetting/v0/DisableCommandLine"
}

func (o *xxx_SetDisableCommandLineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDisableCommandLineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// disableCommandLine {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.DisableCommandLine); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDisableCommandLineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// disableCommandLine {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.DisableCommandLine); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDisableCommandLineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDisableCommandLineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDisableCommandLineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDisableCommandLineRequest structure represents the DisableCommandLine operation request
type SetDisableCommandLineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	DisableCommandLine int16          `idl:"name:disableCommandLine" json:"disable_command_line"`
}

func (o *SetDisableCommandLineRequest) xxx_ToOp(ctx context.Context) *xxx_SetDisableCommandLineOperation {
	if o == nil {
		return &xxx_SetDisableCommandLineOperation{}
	}
	return &xxx_SetDisableCommandLineOperation{
		This:               o.This,
		DisableCommandLine: o.DisableCommandLine,
	}
}

func (o *SetDisableCommandLineRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDisableCommandLineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DisableCommandLine = op.DisableCommandLine
}
func (o *SetDisableCommandLineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetDisableCommandLineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDisableCommandLineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDisableCommandLineResponse structure represents the DisableCommandLine operation response
type SetDisableCommandLineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DisableCommandLine return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDisableCommandLineResponse) xxx_ToOp(ctx context.Context) *xxx_SetDisableCommandLineOperation {
	if o == nil {
		return &xxx_SetDisableCommandLineOperation{}
	}
	return &xxx_SetDisableCommandLineOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetDisableCommandLineResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDisableCommandLineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDisableCommandLineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetDisableCommandLineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDisableCommandLineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEnableScreeningAuditOperation structure represents the EnableScreeningAudit operation
type xxx_GetEnableScreeningAuditOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	EnableScreeningAudit int16          `idl:"name:enableScreeningAudit" json:"enable_screening_audit"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEnableScreeningAuditOperation) OpNum() int { return 15 }

func (o *xxx_GetEnableScreeningAuditOperation) OpName() string {
	return "/IFsrmSetting/v0/EnableScreeningAudit"
}

func (o *xxx_GetEnableScreeningAuditOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnableScreeningAuditOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetEnableScreeningAuditOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetEnableScreeningAuditOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnableScreeningAuditOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// enableScreeningAudit {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.EnableScreeningAudit); err != nil {
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

func (o *xxx_GetEnableScreeningAuditOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// enableScreeningAudit {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.EnableScreeningAudit); err != nil {
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

// GetEnableScreeningAuditRequest structure represents the EnableScreeningAudit operation request
type GetEnableScreeningAuditRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEnableScreeningAuditRequest) xxx_ToOp(ctx context.Context) *xxx_GetEnableScreeningAuditOperation {
	if o == nil {
		return &xxx_GetEnableScreeningAuditOperation{}
	}
	return &xxx_GetEnableScreeningAuditOperation{
		This: o.This,
	}
}

func (o *GetEnableScreeningAuditRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEnableScreeningAuditOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEnableScreeningAuditRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetEnableScreeningAuditRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnableScreeningAuditOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEnableScreeningAuditResponse structure represents the EnableScreeningAudit operation response
type GetEnableScreeningAuditResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	EnableScreeningAudit int16          `idl:"name:enableScreeningAudit" json:"enable_screening_audit"`
	// Return: The EnableScreeningAudit return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEnableScreeningAuditResponse) xxx_ToOp(ctx context.Context) *xxx_GetEnableScreeningAuditOperation {
	if o == nil {
		return &xxx_GetEnableScreeningAuditOperation{}
	}
	return &xxx_GetEnableScreeningAuditOperation{
		That:                 o.That,
		EnableScreeningAudit: o.EnableScreeningAudit,
		Return:               o.Return,
	}
}

func (o *GetEnableScreeningAuditResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEnableScreeningAuditOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EnableScreeningAudit = op.EnableScreeningAudit
	o.Return = op.Return
}
func (o *GetEnableScreeningAuditResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetEnableScreeningAuditResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnableScreeningAuditOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEnableScreeningAuditOperation structure represents the EnableScreeningAudit operation
type xxx_SetEnableScreeningAuditOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	EnableScreeningAudit int16          `idl:"name:enableScreeningAudit" json:"enable_screening_audit"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEnableScreeningAuditOperation) OpNum() int { return 16 }

func (o *xxx_SetEnableScreeningAuditOperation) OpName() string {
	return "/IFsrmSetting/v0/EnableScreeningAudit"
}

func (o *xxx_SetEnableScreeningAuditOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnableScreeningAuditOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// enableScreeningAudit {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.EnableScreeningAudit); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnableScreeningAuditOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// enableScreeningAudit {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.EnableScreeningAudit); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnableScreeningAuditOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnableScreeningAuditOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEnableScreeningAuditOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetEnableScreeningAuditRequest structure represents the EnableScreeningAudit operation request
type SetEnableScreeningAuditRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	EnableScreeningAudit int16          `idl:"name:enableScreeningAudit" json:"enable_screening_audit"`
}

func (o *SetEnableScreeningAuditRequest) xxx_ToOp(ctx context.Context) *xxx_SetEnableScreeningAuditOperation {
	if o == nil {
		return &xxx_SetEnableScreeningAuditOperation{}
	}
	return &xxx_SetEnableScreeningAuditOperation{
		This:                 o.This,
		EnableScreeningAudit: o.EnableScreeningAudit,
	}
}

func (o *SetEnableScreeningAuditRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEnableScreeningAuditOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EnableScreeningAudit = op.EnableScreeningAudit
}
func (o *SetEnableScreeningAuditRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetEnableScreeningAuditRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEnableScreeningAuditOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEnableScreeningAuditResponse structure represents the EnableScreeningAudit operation response
type SetEnableScreeningAuditResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EnableScreeningAudit return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEnableScreeningAuditResponse) xxx_ToOp(ctx context.Context) *xxx_SetEnableScreeningAuditOperation {
	if o == nil {
		return &xxx_SetEnableScreeningAuditOperation{}
	}
	return &xxx_SetEnableScreeningAuditOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetEnableScreeningAuditResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEnableScreeningAuditOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEnableScreeningAuditResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetEnableScreeningAuditResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEnableScreeningAuditOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EmailTestOperation structure represents the EmailTest operation
type xxx_EmailTestOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EmailTestOperation) OpNum() int { return 17 }

func (o *xxx_EmailTestOperation) OpName() string { return "/IFsrmSetting/v0/EmailTest" }

func (o *xxx_EmailTestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EmailTestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// mailTo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailTo != nil {
			_ptr_mailTo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailTo != nil {
					if err := o.MailTo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailTo, _ptr_mailTo); err != nil {
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

func (o *xxx_EmailTestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// mailTo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailTo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailTo == nil {
				o.MailTo = &oaut.String{}
			}
			if err := o.MailTo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailTo := func(ptr interface{}) { o.MailTo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailTo, _s_mailTo, _ptr_mailTo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EmailTestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EmailTestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EmailTestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// EmailTestRequest structure represents the EmailTest operation request
type EmailTestRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// mailTo: Contains the email address for the File Server Resource Manager Protocol
	// to send the test email message to.
	MailTo *oaut.String `idl:"name:mailTo" json:"mail_to"`
}

func (o *EmailTestRequest) xxx_ToOp(ctx context.Context) *xxx_EmailTestOperation {
	if o == nil {
		return &xxx_EmailTestOperation{}
	}
	return &xxx_EmailTestOperation{
		This:   o.This,
		MailTo: o.MailTo,
	}
}

func (o *EmailTestRequest) xxx_FromOp(ctx context.Context, op *xxx_EmailTestOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MailTo = op.MailTo
}
func (o *EmailTestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EmailTestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EmailTestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EmailTestResponse structure represents the EmailTest operation response
type EmailTestResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EmailTest return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EmailTestResponse) xxx_ToOp(ctx context.Context) *xxx_EmailTestOperation {
	if o == nil {
		return &xxx_EmailTestOperation{}
	}
	return &xxx_EmailTestOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *EmailTestResponse) xxx_FromOp(ctx context.Context, op *xxx_EmailTestOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *EmailTestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EmailTestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EmailTestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetActionRunLimitIntervalOperation structure represents the SetActionRunLimitInterval operation
type xxx_SetActionRunLimitIntervalOperation struct {
	This             *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ActionType       fsrm.ActionType `idl:"name:actionType" json:"action_type"`
	DelayTimeMinutes int32           `idl:"name:delayTimeMinutes" json:"delay_time_minutes"`
	Return           int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetActionRunLimitIntervalOperation) OpNum() int { return 18 }

func (o *xxx_SetActionRunLimitIntervalOperation) OpName() string {
	return "/IFsrmSetting/v0/SetActionRunLimitInterval"
}

func (o *xxx_SetActionRunLimitIntervalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetActionRunLimitIntervalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// actionType {in} (1:{alias=FsrmActionType}(enum))
	{
		if err := w.WriteData(uint16(o.ActionType)); err != nil {
			return err
		}
	}
	// delayTimeMinutes {in} (1:(int32))
	{
		if err := w.WriteData(o.DelayTimeMinutes); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetActionRunLimitIntervalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// actionType {in} (1:{alias=FsrmActionType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.ActionType)); err != nil {
			return err
		}
	}
	// delayTimeMinutes {in} (1:(int32))
	{
		if err := w.ReadData(&o.DelayTimeMinutes); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetActionRunLimitIntervalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetActionRunLimitIntervalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetActionRunLimitIntervalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetActionRunLimitIntervalRequest structure represents the SetActionRunLimitInterval operation request
type SetActionRunLimitIntervalRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// actionType: Contains the action type that this run limit interval applies to.
	ActionType fsrm.ActionType `idl:"name:actionType" json:"action_type"`
	// delayTimeMinutes: Contains the run limit interval for this action type.
	DelayTimeMinutes int32 `idl:"name:delayTimeMinutes" json:"delay_time_minutes"`
}

func (o *SetActionRunLimitIntervalRequest) xxx_ToOp(ctx context.Context) *xxx_SetActionRunLimitIntervalOperation {
	if o == nil {
		return &xxx_SetActionRunLimitIntervalOperation{}
	}
	return &xxx_SetActionRunLimitIntervalOperation{
		This:             o.This,
		ActionType:       o.ActionType,
		DelayTimeMinutes: o.DelayTimeMinutes,
	}
}

func (o *SetActionRunLimitIntervalRequest) xxx_FromOp(ctx context.Context, op *xxx_SetActionRunLimitIntervalOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ActionType = op.ActionType
	o.DelayTimeMinutes = op.DelayTimeMinutes
}
func (o *SetActionRunLimitIntervalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetActionRunLimitIntervalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetActionRunLimitIntervalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetActionRunLimitIntervalResponse structure represents the SetActionRunLimitInterval operation response
type SetActionRunLimitIntervalResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetActionRunLimitInterval return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetActionRunLimitIntervalResponse) xxx_ToOp(ctx context.Context) *xxx_SetActionRunLimitIntervalOperation {
	if o == nil {
		return &xxx_SetActionRunLimitIntervalOperation{}
	}
	return &xxx_SetActionRunLimitIntervalOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetActionRunLimitIntervalResponse) xxx_FromOp(ctx context.Context, op *xxx_SetActionRunLimitIntervalOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetActionRunLimitIntervalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetActionRunLimitIntervalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetActionRunLimitIntervalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetActionRunLimitIntervalOperation structure represents the GetActionRunLimitInterval operation
type xxx_GetActionRunLimitIntervalOperation struct {
	This             *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ActionType       fsrm.ActionType `idl:"name:actionType" json:"action_type"`
	DelayTimeMinutes int32           `idl:"name:delayTimeMinutes" json:"delay_time_minutes"`
	Return           int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetActionRunLimitIntervalOperation) OpNum() int { return 19 }

func (o *xxx_GetActionRunLimitIntervalOperation) OpName() string {
	return "/IFsrmSetting/v0/GetActionRunLimitInterval"
}

func (o *xxx_GetActionRunLimitIntervalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionRunLimitIntervalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// actionType {in} (1:{alias=FsrmActionType}(enum))
	{
		if err := w.WriteData(uint16(o.ActionType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionRunLimitIntervalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// actionType {in} (1:{alias=FsrmActionType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.ActionType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionRunLimitIntervalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionRunLimitIntervalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// delayTimeMinutes {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.DelayTimeMinutes); err != nil {
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

func (o *xxx_GetActionRunLimitIntervalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// delayTimeMinutes {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.DelayTimeMinutes); err != nil {
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

// GetActionRunLimitIntervalRequest structure represents the GetActionRunLimitInterval operation request
type GetActionRunLimitIntervalRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// actionType: Contains the action type to return the Run limit interval for.
	ActionType fsrm.ActionType `idl:"name:actionType" json:"action_type"`
}

func (o *GetActionRunLimitIntervalRequest) xxx_ToOp(ctx context.Context) *xxx_GetActionRunLimitIntervalOperation {
	if o == nil {
		return &xxx_GetActionRunLimitIntervalOperation{}
	}
	return &xxx_GetActionRunLimitIntervalOperation{
		This:       o.This,
		ActionType: o.ActionType,
	}
}

func (o *GetActionRunLimitIntervalRequest) xxx_FromOp(ctx context.Context, op *xxx_GetActionRunLimitIntervalOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ActionType = op.ActionType
}
func (o *GetActionRunLimitIntervalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetActionRunLimitIntervalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionRunLimitIntervalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetActionRunLimitIntervalResponse structure represents the GetActionRunLimitInterval operation response
type GetActionRunLimitIntervalResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// delayTimeMinutes: Pointer to a variable that upon completion contains the Run limit
	// interval for the specified action type.
	DelayTimeMinutes int32 `idl:"name:delayTimeMinutes" json:"delay_time_minutes"`
	// Return: The GetActionRunLimitInterval return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetActionRunLimitIntervalResponse) xxx_ToOp(ctx context.Context) *xxx_GetActionRunLimitIntervalOperation {
	if o == nil {
		return &xxx_GetActionRunLimitIntervalOperation{}
	}
	return &xxx_GetActionRunLimitIntervalOperation{
		That:             o.That,
		DelayTimeMinutes: o.DelayTimeMinutes,
		Return:           o.Return,
	}
}

func (o *GetActionRunLimitIntervalResponse) xxx_FromOp(ctx context.Context, op *xxx_GetActionRunLimitIntervalOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DelayTimeMinutes = op.DelayTimeMinutes
	o.Return = op.Return
}
func (o *GetActionRunLimitIntervalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetActionRunLimitIntervalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionRunLimitIntervalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
