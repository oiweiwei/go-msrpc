package sasec

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

var (
	// import guard
	GoPackage = "tsch"
)

var (
	// Syntax UUID
	SasecSyntaxUUID = &uuid.UUID{TimeLow: 0x378e52b0, TimeMid: 0xc0a9, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x2d, Node: [6]uint8{0x0, 0xaa, 0x0, 0x51, 0xe4, 0xf}}
	// Syntax ID
	SasecSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: SasecSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// sasec interface.
type SasecClient interface {

	// The SASetAccountInformation method MUST set the credentials under which the task
	// MUST run.
	//
	// Return Values: For more information about return codes, see section 2.3.14 or Win32
	// Error Codes in [MS-ERREF] section 2.1.<52>
	SetAccountInformation(context.Context, *SetAccountInformationRequest, ...dcerpc.CallOption) (*SetAccountInformationResponse, error)

	// The SASetNSAccountInformation method MUST configure the credentials under which all
	// ATSvc tasks run.
	SetNSAccountInformation(context.Context, *SetNSAccountInformationRequest, ...dcerpc.CallOption) (*SetNSAccountInformationResponse, error)

	// The SAGetNSAccountInformation method MUST return the ATSvc account name.
	//
	// Return Values: For more information about return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.
	GetNSAccountInformation(context.Context, *GetNSAccountInformationRequest, ...dcerpc.CallOption) (*GetNSAccountInformationResponse, error)

	// The SAGetAccountInformation method MUST retrieve the account name for a specified
	// task.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	GetAccountInformation(context.Context, *GetAccountInformationRequest, ...dcerpc.CallOption) (*GetAccountInformationResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// Cnlen represents the CNLEN RPC constant
var Cnlen = 15

// Dnlen represents the DNLEN RPC constant
var Dnlen = 15

// Unlen represents the UNLEN RPC constant
var Unlen = 256

// MaxBufferSize represents the MAX_BUFFER_SIZE RPC constant
var MaxBufferSize = 273

type xxx_DefaultSasecClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultSasecClient) SetAccountInformation(ctx context.Context, in *SetAccountInformationRequest, opts ...dcerpc.CallOption) (*SetAccountInformationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetAccountInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSasecClient) SetNSAccountInformation(ctx context.Context, in *SetNSAccountInformationRequest, opts ...dcerpc.CallOption) (*SetNSAccountInformationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetNSAccountInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSasecClient) GetNSAccountInformation(ctx context.Context, in *GetNSAccountInformationRequest, opts ...dcerpc.CallOption) (*GetNSAccountInformationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNSAccountInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSasecClient) GetAccountInformation(ctx context.Context, in *GetAccountInformationRequest, opts ...dcerpc.CallOption) (*GetAccountInformationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetAccountInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSasecClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultSasecClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewSasecClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (SasecClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(SasecSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultSasecClient{cc: cc}, nil
}

// xxx_SetAccountInformationOperation structure represents the SASetAccountInformation operation
type xxx_SetAccountInformationOperation struct {
	Handle   string `idl:"name:Handle;string;pointer:unique" json:"handle"`
	JobName  string `idl:"name:pwszJobName;string" json:"job_name"`
	Account  string `idl:"name:pwszAccount;string" json:"account"`
	Password string `idl:"name:pwszPassword;string;pointer:unique" json:"password"`
	JobFlags uint32 `idl:"name:dwJobFlags" json:"job_flags"`
	Return   int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAccountInformationOperation) OpNum() int { return 0 }

func (o *xxx_SetAccountInformationOperation) OpName() string {
	return "/sasec/v1/SASetAccountInformation"
}

func (o *xxx_SetAccountInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Handle {in} (1:{handle, string, pointer=unique, alias=SASEC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.Handle != "" {
			_ptr_Handle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Handle); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Handle, _ptr_Handle); err != nil {
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
	// pwszJobName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.JobName); err != nil {
			return err
		}
	}
	// pwszAccount {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Account); err != nil {
			return err
		}
	}
	// pwszPassword {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Password != "" {
			_ptr_pwszPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_pwszPassword); err != nil {
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
	// dwJobFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.JobFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Handle {in} (1:{handle, string, pointer=unique, alias=SASEC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_Handle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Handle); err != nil {
				return err
			}
			return nil
		})
		_s_Handle := func(ptr interface{}) { o.Handle = *ptr.(*string) }
		if err := w.ReadPointer(&o.Handle, _s_Handle, _ptr_Handle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszJobName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.JobName); err != nil {
			return err
		}
	}
	// pwszAccount {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Account); err != nil {
			return err
		}
	}
	// pwszPassword {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
				return err
			}
			return nil
		})
		_s_pwszPassword := func(ptr interface{}) { o.Password = *ptr.(*string) }
		if err := w.ReadPointer(&o.Password, _s_pwszPassword, _ptr_pwszPassword); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwJobFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.JobFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetAccountInformationRequest structure represents the SASetAccountInformation operation request
type SetAccountInformationRequest struct {
	// Handle: Pointer to a Unicode string that MUST specify the server. The client MUST
	// map this string to an RPC binding handle. The server MUST ignore this parameter.
	// For more information, see [C706] sections 4.3.5 and 5.1.5.2.
	Handle string `idl:"name:Handle;string;pointer:unique" json:"handle"`
	// pwszJobName: Pointer to a string that MUST specify a task name, such as "MyJob.job".
	JobName string `idl:"name:pwszJobName;string" json:"job_name"`
	// pwszAccount: Pointer to a string that MUST specify the account name. This string
	// MAY be expressed either as a UPN in the form user@domain or as a Security Account
	// Manager (SAM) name in the form domain\account.
	Account string `idl:"name:pwszAccount;string" json:"account"`
	// pwszPassword: Pointer to a string that MUST specify the password for the account.
	// See section 5.1.
	Password string `idl:"name:pwszPassword;string;pointer:unique" json:"password"`
	// dwJobFlags: The dwJobFlags field MUST contain individual bit flags that MUST have
	// one or more of the following values:
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |  8  | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |     |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | R L | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| RL TASK_FLAG_RUN_ONLY_IF_LOGGED_ON | When set, the task MUST run only if the user specified is logged on              |
	//	|                                    | interactively.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	JobFlags uint32 `idl:"name:dwJobFlags" json:"job_flags"`
}

func (o *SetAccountInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAccountInformationOperation) *xxx_SetAccountInformationOperation {
	if op == nil {
		op = &xxx_SetAccountInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.JobName = o.JobName
	op.Account = o.Account
	op.Password = o.Password
	op.JobFlags = o.JobFlags
	return op
}

func (o *SetAccountInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAccountInformationOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.JobName = op.JobName
	o.Account = op.Account
	o.Password = op.Password
	o.JobFlags = op.JobFlags
}
func (o *SetAccountInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAccountInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAccountInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAccountInformationResponse structure represents the SASetAccountInformation operation response
type SetAccountInformationResponse struct {
	// Return: The SASetAccountInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAccountInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAccountInformationOperation) *xxx_SetAccountInformationOperation {
	if op == nil {
		op = &xxx_SetAccountInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetAccountInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAccountInformationOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetAccountInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAccountInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAccountInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNSAccountInformationOperation structure represents the SASetNSAccountInformation operation
type xxx_SetNSAccountInformationOperation struct {
	Handle   string `idl:"name:Handle;string;pointer:unique" json:"handle"`
	Account  string `idl:"name:pwszAccount;string;pointer:unique" json:"account"`
	Password string `idl:"name:pwszPassword;string;pointer:unique" json:"password"`
	Return   int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNSAccountInformationOperation) OpNum() int { return 1 }

func (o *xxx_SetNSAccountInformationOperation) OpName() string {
	return "/sasec/v1/SASetNSAccountInformation"
}

func (o *xxx_SetNSAccountInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNSAccountInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Handle {in} (1:{handle, string, pointer=unique, alias=SASEC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.Handle != "" {
			_ptr_Handle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Handle); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Handle, _ptr_Handle); err != nil {
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
	// pwszAccount {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Account != "" {
			_ptr_pwszAccount := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Account); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Account, _ptr_pwszAccount); err != nil {
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
	// pwszPassword {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Password != "" {
			_ptr_pwszPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_pwszPassword); err != nil {
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

func (o *xxx_SetNSAccountInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Handle {in} (1:{handle, string, pointer=unique, alias=SASEC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_Handle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Handle); err != nil {
				return err
			}
			return nil
		})
		_s_Handle := func(ptr interface{}) { o.Handle = *ptr.(*string) }
		if err := w.ReadPointer(&o.Handle, _s_Handle, _ptr_Handle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAccount {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAccount := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Account); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAccount := func(ptr interface{}) { o.Account = *ptr.(*string) }
		if err := w.ReadPointer(&o.Account, _s_pwszAccount, _ptr_pwszAccount); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszPassword {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
				return err
			}
			return nil
		})
		_s_pwszPassword := func(ptr interface{}) { o.Password = *ptr.(*string) }
		if err := w.ReadPointer(&o.Password, _s_pwszPassword, _ptr_pwszPassword); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNSAccountInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNSAccountInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNSAccountInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetNSAccountInformationRequest structure represents the SASetNSAccountInformation operation request
type SetNSAccountInformationRequest struct {
	// Handle: Pointer to a Unicode string that MUST specify the server. The client MUST
	// map this string to an RPC binding handle. The server MUST ignore this parameter.
	// For more details, see [C706] sections 4.3.5 and 5.1.5.2.
	Handle string `idl:"name:Handle;string;pointer:unique" json:"handle"`
	// pwszAccount: MUST be a pointer to a string that specifies the account name.
	Account string `idl:"name:pwszAccount;string;pointer:unique" json:"account"`
	// pwszPassword: MUST be a pointer to a string that specifies the password for the named
	// account. See section 5.1 for security considerations.
	Password string `idl:"name:pwszPassword;string;pointer:unique" json:"password"`
}

func (o *SetNSAccountInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNSAccountInformationOperation) *xxx_SetNSAccountInformationOperation {
	if op == nil {
		op = &xxx_SetNSAccountInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Account = o.Account
	op.Password = o.Password
	return op
}

func (o *SetNSAccountInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNSAccountInformationOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Account = op.Account
	o.Password = op.Password
}
func (o *SetNSAccountInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNSAccountInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNSAccountInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNSAccountInformationResponse structure represents the SASetNSAccountInformation operation response
type SetNSAccountInformationResponse struct {
	// Return: The SASetNSAccountInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNSAccountInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNSAccountInformationOperation) *xxx_SetNSAccountInformationOperation {
	if op == nil {
		op = &xxx_SetNSAccountInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetNSAccountInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNSAccountInformationOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetNSAccountInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNSAccountInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNSAccountInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNSAccountInformationOperation structure represents the SAGetNSAccountInformation operation
type xxx_GetNSAccountInformationOperation struct {
	Handle     string `idl:"name:Handle;string;pointer:unique" json:"handle"`
	BufferSize uint32 `idl:"name:ccBufferSize" json:"buffer_size"`
	Buffer     string `idl:"name:wszBuffer;size_is:(ccBufferSize)" json:"buffer"`
	Return     int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNSAccountInformationOperation) OpNum() int { return 2 }

func (o *xxx_GetNSAccountInformationOperation) OpName() string {
	return "/sasec/v1/SAGetNSAccountInformation"
}

func (o *xxx_GetNSAccountInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != "" && o.BufferSize == 0 {
		o.BufferSize = uint32(ndr.UTF16Len(o.Buffer))
	}
	if o.BufferSize > uint32(273) {
		return fmt.Errorf("BufferSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNSAccountInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Handle {in} (1:{handle, string, pointer=unique, alias=SASEC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.Handle != "" {
			_ptr_Handle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Handle); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Handle, _ptr_Handle); err != nil {
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
	// ccBufferSize {in} (1:{range=(0,273), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// wszBuffer {in, out} (1:[dim:0,size_is=ccBufferSize,string](wchar))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_Buffer_buf := utf16.Encode([]rune(o.Buffer))
		if uint64(len(_Buffer_buf)) > sizeInfo[0] {
			_Buffer_buf = _Buffer_buf[:sizeInfo[0]]
		}
		for i1 := range _Buffer_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetNSAccountInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Handle {in} (1:{handle, string, pointer=unique, alias=SASEC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_Handle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Handle); err != nil {
				return err
			}
			return nil
		})
		_s_Handle := func(ptr interface{}) { o.Handle = *ptr.(*string) }
		if err := w.ReadPointer(&o.Handle, _s_Handle, _ptr_Handle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ccBufferSize {in} (1:{range=(0,273), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// wszBuffer {in, out} (1:[dim:0,size_is=ccBufferSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Buffer_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
		}
		_Buffer_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Buffer_buf {
			i1 := i1
			if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
	}
	return nil
}

func (o *xxx_GetNSAccountInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNSAccountInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// wszBuffer {in, out} (1:[dim:0,size_is=ccBufferSize,string](wchar))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_Buffer_buf := utf16.Encode([]rune(o.Buffer))
		if uint64(len(_Buffer_buf)) > sizeInfo[0] {
			_Buffer_buf = _Buffer_buf[:sizeInfo[0]]
		}
		for i1 := range _Buffer_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
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

func (o *xxx_GetNSAccountInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// wszBuffer {in, out} (1:[dim:0,size_is=ccBufferSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Buffer_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
		}
		_Buffer_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Buffer_buf {
			i1 := i1
			if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNSAccountInformationRequest structure represents the SAGetNSAccountInformation operation request
type GetNSAccountInformationRequest struct {
	// Handle: Pointer to a Unicode string that MUST specify the server. The client MUST
	// map this string to an RPC Binding handle. The server MUST ignore this parameter.
	// For more details, see [C706] sections 4.3.5 and 5.1.5.2.
	Handle string `idl:"name:Handle;string;pointer:unique" json:"handle"`
	// ccBufferSize: MUST contain the number of characters in the array supplied by the
	// client and filled by the server. This value MUST be the size of the wszBuffer parameter.
	// MAX_BUFFER_SIZE is equal to 273. For more information on MAX_BUFFER_SIZE, see the
	// SaSec interface IDL (section 6.2).
	BufferSize uint32 `idl:"name:ccBufferSize" json:"buffer_size"`
	// wszBuffer: Upon input, MUST be an empty array of size equal to the ccBufferSize parameter.
	// The client SHOULD initialize the array to contain zeroes. Upon return, the array
	// MUST contain the ATSvc account name.
	Buffer string `idl:"name:wszBuffer;size_is:(ccBufferSize)" json:"buffer"`
}

func (o *GetNSAccountInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNSAccountInformationOperation) *xxx_GetNSAccountInformationOperation {
	if op == nil {
		op = &xxx_GetNSAccountInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.BufferSize = o.BufferSize
	op.Buffer = o.Buffer
	return op
}

func (o *GetNSAccountInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNSAccountInformationOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.BufferSize = op.BufferSize
	o.Buffer = op.Buffer
}
func (o *GetNSAccountInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNSAccountInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNSAccountInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNSAccountInformationResponse structure represents the SAGetNSAccountInformation operation response
type GetNSAccountInformationResponse struct {
	// XXX: ccBufferSize is an implicit input depedency for output parameters
	BufferSize uint32 `idl:"name:ccBufferSize" json:"buffer_size"`

	// wszBuffer: Upon input, MUST be an empty array of size equal to the ccBufferSize parameter.
	// The client SHOULD initialize the array to contain zeroes. Upon return, the array
	// MUST contain the ATSvc account name.
	Buffer string `idl:"name:wszBuffer;size_is:(ccBufferSize)" json:"buffer"`
	// Return: The SAGetNSAccountInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNSAccountInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNSAccountInformationOperation) *xxx_GetNSAccountInformationOperation {
	if op == nil {
		op = &xxx_GetNSAccountInformationOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferSize == uint32(0) {
		op.BufferSize = o.BufferSize
	}

	op.Buffer = o.Buffer
	op.Return = o.Return
	return op
}

func (o *GetNSAccountInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNSAccountInformationOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferSize = op.BufferSize

	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *GetNSAccountInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNSAccountInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNSAccountInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAccountInformationOperation structure represents the SAGetAccountInformation operation
type xxx_GetAccountInformationOperation struct {
	Handle     string `idl:"name:Handle;string;pointer:unique" json:"handle"`
	JobName    string `idl:"name:pwszJobName;string" json:"job_name"`
	BufferSize uint32 `idl:"name:ccBufferSize" json:"buffer_size"`
	Buffer     string `idl:"name:wszBuffer;size_is:(ccBufferSize)" json:"buffer"`
	Return     int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAccountInformationOperation) OpNum() int { return 3 }

func (o *xxx_GetAccountInformationOperation) OpName() string {
	return "/sasec/v1/SAGetAccountInformation"
}

func (o *xxx_GetAccountInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != "" && o.BufferSize == 0 {
		o.BufferSize = uint32(ndr.UTF16Len(o.Buffer))
	}
	if o.BufferSize > uint32(273) {
		return fmt.Errorf("BufferSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccountInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Handle {in} (1:{handle, string, pointer=unique, alias=SASEC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.Handle != "" {
			_ptr_Handle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Handle); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Handle, _ptr_Handle); err != nil {
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
	// pwszJobName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.JobName); err != nil {
			return err
		}
	}
	// ccBufferSize {in} (1:{range=(0,273), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// wszBuffer {in, out} (1:[dim:0,size_is=ccBufferSize,string](wchar))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_Buffer_buf := utf16.Encode([]rune(o.Buffer))
		if uint64(len(_Buffer_buf)) > sizeInfo[0] {
			_Buffer_buf = _Buffer_buf[:sizeInfo[0]]
		}
		for i1 := range _Buffer_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetAccountInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Handle {in} (1:{handle, string, pointer=unique, alias=SASEC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_Handle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Handle); err != nil {
				return err
			}
			return nil
		})
		_s_Handle := func(ptr interface{}) { o.Handle = *ptr.(*string) }
		if err := w.ReadPointer(&o.Handle, _s_Handle, _ptr_Handle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszJobName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.JobName); err != nil {
			return err
		}
	}
	// ccBufferSize {in} (1:{range=(0,273), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// wszBuffer {in, out} (1:[dim:0,size_is=ccBufferSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Buffer_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
		}
		_Buffer_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Buffer_buf {
			i1 := i1
			if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
	}
	return nil
}

func (o *xxx_GetAccountInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccountInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// wszBuffer {in, out} (1:[dim:0,size_is=ccBufferSize,string](wchar))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_Buffer_buf := utf16.Encode([]rune(o.Buffer))
		if uint64(len(_Buffer_buf)) > sizeInfo[0] {
			_Buffer_buf = _Buffer_buf[:sizeInfo[0]]
		}
		for i1 := range _Buffer_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
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

func (o *xxx_GetAccountInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// wszBuffer {in, out} (1:[dim:0,size_is=ccBufferSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Buffer_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
		}
		_Buffer_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Buffer_buf {
			i1 := i1
			if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetAccountInformationRequest structure represents the SAGetAccountInformation operation request
type GetAccountInformationRequest struct {
	// Handle: Pointer to a Unicode string that MUST specify the server. The client MUST
	// map this string to an RPC binding handle. The server MUST ignore this parameter.
	// For more details, see [C706] sections 4.3.5 and 5.1.5.2.
	Handle string `idl:"name:Handle;string;pointer:unique" json:"handle"`
	// pwszJobName: MUST be a pointer to a string that specifies a task name, such as "MyJob.job".
	JobName string `idl:"name:pwszJobName;string" json:"job_name"`
	// ccBufferSize: MUST contain the number of characters in the array supplied by the
	// client and filled by the server. This value MUST be the size of the wszBuffer parameter.
	// MAX_BUFFER_SIZE is equal to 273. For more information on MAX_BUFFER_SIZE, see the
	// SaSec interface IDL (section 6.2).
	BufferSize uint32 `idl:"name:ccBufferSize" json:"buffer_size"`
	// wszBuffer: Upon input, MUST be an empty array of size equal to the ccBufferSize parameter.
	// The client SHOULD initialize the array to contain zeroes. Upon return, the array
	// MUST contain the name of the account to be used as the context the task runs under.
	Buffer string `idl:"name:wszBuffer;size_is:(ccBufferSize)" json:"buffer"`
}

func (o *GetAccountInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAccountInformationOperation) *xxx_GetAccountInformationOperation {
	if op == nil {
		op = &xxx_GetAccountInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.JobName = o.JobName
	op.BufferSize = o.BufferSize
	op.Buffer = o.Buffer
	return op
}

func (o *GetAccountInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAccountInformationOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.JobName = op.JobName
	o.BufferSize = op.BufferSize
	o.Buffer = op.Buffer
}
func (o *GetAccountInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAccountInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccountInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAccountInformationResponse structure represents the SAGetAccountInformation operation response
type GetAccountInformationResponse struct {
	// XXX: ccBufferSize is an implicit input depedency for output parameters
	BufferSize uint32 `idl:"name:ccBufferSize" json:"buffer_size"`

	// wszBuffer: Upon input, MUST be an empty array of size equal to the ccBufferSize parameter.
	// The client SHOULD initialize the array to contain zeroes. Upon return, the array
	// MUST contain the name of the account to be used as the context the task runs under.
	Buffer string `idl:"name:wszBuffer;size_is:(ccBufferSize)" json:"buffer"`
	// Return: The SAGetAccountInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAccountInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAccountInformationOperation) *xxx_GetAccountInformationOperation {
	if op == nil {
		op = &xxx_GetAccountInformationOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferSize == uint32(0) {
		op.BufferSize = o.BufferSize
	}

	op.Buffer = o.Buffer
	op.Return = o.Return
	return op
}

func (o *GetAccountInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAccountInformationOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferSize = op.BufferSize

	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *GetAccountInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAccountInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccountInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
