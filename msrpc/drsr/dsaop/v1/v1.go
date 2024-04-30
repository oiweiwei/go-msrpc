package dsaop

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
)

var (
	// import guard
	GoPackage = "drsr"
)

var (
	// Syntax UUID
	DsaopSyntaxUUID = &uuid.UUID{TimeLow: 0x7c44d7d4, TimeMid: 0x31d5, TimeHiAndVersion: 0x424c, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0x5e, Node: [6]uint8{0x2b, 0x3e, 0x1f, 0x32, 0x3d, 0x22}}
	// Syntax ID
	DsaopSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: DsaopSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// dsaop interface.
type DsaopClient interface {

	// The IDL_DSAPrepareScript method prepares the DC to run a maintenance script.
	//
	// Return Values: 0 if successful, or a Windows error code if a failure occurs.
	PrepareScript(context.Context, *PrepareScriptRequest, ...dcerpc.CallOption) (*PrepareScriptResponse, error)

	// The IDL_DSAExecuteScript method executes a maintenance script.
	//
	// Return Values: 0 if successful, or a Windows error code if a failure occurs.
	ExecuteScript(context.Context, *ExecuteScriptRequest, ...dcerpc.CallOption) (*ExecuteScriptResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error
}

// MessageExecuteScriptRequestV1 structure represents DSA_MSG_EXECUTE_SCRIPT_REQ_V1 RPC structure.
//
// The DSA_MSG_EXECUTE_SCRIPT_REQ_V1 structure defines a request message sent to the
// IDL_DSAExecuteScript method.
type MessageExecuteScriptRequestV1 struct {
	// Flags:  Unused. MUST be 0 and ignored.
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// cbPassword:  The count, in bytes, of the pbPassword array.
	PasswordLength uint32 `idl:"name:cbPassword" json:"password_length"`
	// pbPassword:  The password.
	Password []byte `idl:"name:pbPassword;size_is:(cbPassword)" json:"password"`
}

func (o *MessageExecuteScriptRequestV1) xxx_PreparePayload(ctx context.Context) error {
	if o.Password != nil && o.PasswordLength == 0 {
		o.PasswordLength = uint32(len(o.Password))
	}
	if o.PasswordLength < uint32(1) || o.PasswordLength > uint32(1024) {
		return fmt.Errorf("PasswordLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageExecuteScriptRequestV1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.PasswordLength); err != nil {
		return err
	}
	if o.Password != nil || o.PasswordLength > 0 {
		_ptr_pbPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.PasswordLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Password {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Password[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Password); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Password, _ptr_pbPassword); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageExecuteScriptRequestV1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordLength); err != nil {
		return err
	}
	_ptr_pbPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.PasswordLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.PasswordLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Password", sizeInfo[0])
		}
		o.Password = make([]byte, sizeInfo[0])
		for i1 := range o.Password {
			i1 := i1
			if err := w.ReadData(&o.Password[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbPassword := func(ptr interface{}) { o.Password = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Password, _s_pbPassword, _ptr_pbPassword); err != nil {
		return err
	}
	return nil
}

// MessageExecuteScriptRequest structure represents DSA_MSG_EXECUTE_SCRIPT_REQ RPC union.
//
// The DSA_MSG_EXECUTE_SCRIPT_REQ union defines the request messages sent to the IDL_DSAExecuteScript
// method.
type MessageExecuteScriptRequest struct {
	// Types that are assignable to Value
	//
	// *MessageExecuteScriptRequest_V1
	Value is_MessageExecuteScriptRequest `json:"value"`
}

func (o *MessageExecuteScriptRequest) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *MessageExecuteScriptRequest_V1:
		if value != nil {
			return value.V1
		}
	}
	return nil
}

type is_MessageExecuteScriptRequest interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_MessageExecuteScriptRequest()
}

func (o *MessageExecuteScriptRequest) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *MessageExecuteScriptRequest_V1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *MessageExecuteScriptRequest) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*MessageExecuteScriptRequest_V1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MessageExecuteScriptRequest_V1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *MessageExecuteScriptRequest) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &MessageExecuteScriptRequest_V1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// MessageExecuteScriptRequest_V1 structure represents DSA_MSG_EXECUTE_SCRIPT_REQ RPC union arm.
//
// It has following labels: 1
type MessageExecuteScriptRequest_V1 struct {
	// V1:  The version 1 request.
	V1 *MessageExecuteScriptRequestV1 `idl:"name:V1" json:"v1"`
}

func (*MessageExecuteScriptRequest_V1) is_MessageExecuteScriptRequest() {}

func (o *MessageExecuteScriptRequest_V1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.V1 != nil {
		if err := o.V1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MessageExecuteScriptRequestV1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageExecuteScriptRequest_V1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.V1 == nil {
		o.V1 = &MessageExecuteScriptRequestV1{}
	}
	if err := o.V1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MessageExecuteScriptReplyV1 structure represents DSA_MSG_EXECUTE_SCRIPT_REPLY_V1 RPC structure.
//
// The DSA_MSG_EXECUTE_SCRIPT_REPLY_V1 structure defines a response message received
// from the IDL_DSAExecuteScript method.
type MessageExecuteScriptReplyV1 struct {
	// dwOperationStatus:  0 if successful, or a Windows error code if a fatal error occurred.
	OperationStatus uint32 `idl:"name:dwOperationStatus" json:"operation_status"`
	// pwErrMessage:  Null if successful, or a description of the error if a fatal error
	// occurred.
	ErrorMessage string `idl:"name:pwErrMessage;string" json:"error_message"`
}

func (o *MessageExecuteScriptReplyV1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageExecuteScriptReplyV1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.OperationStatus); err != nil {
		return err
	}
	if o.ErrorMessage != "" {
		_ptr_pwErrMessage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ErrorMessage); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ErrorMessage, _ptr_pwErrMessage); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageExecuteScriptReplyV1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperationStatus); err != nil {
		return err
	}
	_ptr_pwErrMessage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ErrorMessage); err != nil {
			return err
		}
		return nil
	})
	_s_pwErrMessage := func(ptr interface{}) { o.ErrorMessage = *ptr.(*string) }
	if err := w.ReadPointer(&o.ErrorMessage, _s_pwErrMessage, _ptr_pwErrMessage); err != nil {
		return err
	}
	return nil
}

// MessageExecuteScriptReply structure represents DSA_MSG_EXECUTE_SCRIPT_REPLY RPC union.
//
// The DSA_MSG_EXECUTE_SCRIPT_REPLY union defines the response messages received from
// the IDL_DSAExecuteScript method.
type MessageExecuteScriptReply struct {
	// Types that are assignable to Value
	//
	// *MessageExecuteScriptReply_V1
	Value is_MessageExecuteScriptReply `json:"value"`
}

func (o *MessageExecuteScriptReply) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *MessageExecuteScriptReply_V1:
		if value != nil {
			return value.V1
		}
	}
	return nil
}

type is_MessageExecuteScriptReply interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_MessageExecuteScriptReply()
}

func (o *MessageExecuteScriptReply) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *MessageExecuteScriptReply_V1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *MessageExecuteScriptReply) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*MessageExecuteScriptReply_V1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MessageExecuteScriptReply_V1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *MessageExecuteScriptReply) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &MessageExecuteScriptReply_V1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// MessageExecuteScriptReply_V1 structure represents DSA_MSG_EXECUTE_SCRIPT_REPLY RPC union arm.
//
// It has following labels: 1
type MessageExecuteScriptReply_V1 struct {
	// V1:  The version 1 request.
	V1 *MessageExecuteScriptReplyV1 `idl:"name:V1" json:"v1"`
}

func (*MessageExecuteScriptReply_V1) is_MessageExecuteScriptReply() {}

func (o *MessageExecuteScriptReply_V1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.V1 != nil {
		if err := o.V1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MessageExecuteScriptReplyV1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageExecuteScriptReply_V1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.V1 == nil {
		o.V1 = &MessageExecuteScriptReplyV1{}
	}
	if err := o.V1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MessagePrepareScriptRequestV1 structure represents DSA_MSG_PREPARE_SCRIPT_REQ_V1 RPC structure.
//
// The DSA_MSG_PREPARE_SCRIPT_REQ_V1 structure defines a request message sent to the
// IDL_DSAPrepareScript method.
type MessagePrepareScriptRequestV1 struct {
	// Reserved:  Unused. MUST be 0 and ignored.
	_ uint32 `idl:"name:Reserved"`
}

func (o *MessagePrepareScriptRequestV1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessagePrepareScriptRequestV1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	return nil
}
func (o *MessagePrepareScriptRequestV1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint32
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	return nil
}

// MessagePrepareScriptRequest structure represents DSA_MSG_PREPARE_SCRIPT_REQ RPC union.
//
// The DSA_MSG_PREPARE_SCRIPT_REQ union defines the request messages sent to the IDL_DSAPrepareScript
// method.
type MessagePrepareScriptRequest struct {
	// Types that are assignable to Value
	//
	// *MessagePrepareScriptRequest_V1
	Value is_MessagePrepareScriptRequest `json:"value"`
}

func (o *MessagePrepareScriptRequest) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *MessagePrepareScriptRequest_V1:
		if value != nil {
			return value.V1
		}
	}
	return nil
}

type is_MessagePrepareScriptRequest interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_MessagePrepareScriptRequest()
}

func (o *MessagePrepareScriptRequest) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *MessagePrepareScriptRequest_V1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *MessagePrepareScriptRequest) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*MessagePrepareScriptRequest_V1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MessagePrepareScriptRequest_V1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *MessagePrepareScriptRequest) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &MessagePrepareScriptRequest_V1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// MessagePrepareScriptRequest_V1 structure represents DSA_MSG_PREPARE_SCRIPT_REQ RPC union arm.
//
// It has following labels: 1
type MessagePrepareScriptRequest_V1 struct {
	// V1:  The version 1 request.
	V1 *MessagePrepareScriptRequestV1 `idl:"name:V1" json:"v1"`
}

func (*MessagePrepareScriptRequest_V1) is_MessagePrepareScriptRequest() {}

func (o *MessagePrepareScriptRequest_V1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.V1 != nil {
		if err := o.V1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MessagePrepareScriptRequestV1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessagePrepareScriptRequest_V1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.V1 == nil {
		o.V1 = &MessagePrepareScriptRequestV1{}
	}
	if err := o.V1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MessagePrepareScriptReplyV1 structure represents DSA_MSG_PREPARE_SCRIPT_REPLY_V1 RPC structure.
//
// The DSA_MSG_PREPARE_SCRIPT_REPLY_V1 structure defines a response message received
// from the IDL_DSAPrepareScript method.
type MessagePrepareScriptReplyV1 struct {
	// dwOperationStatus:  0 if successful, or a Windows error code if a fatal error occurred.
	OperationStatus uint32 `idl:"name:dwOperationStatus" json:"operation_status"`
	// pwErrMessage:  Null if successful, or a description of an error if a fatal error
	// occurred.
	ErrorMessage string `idl:"name:pwErrMessage;string" json:"error_message"`
	// cbPassword:  The count, in bytes, of the pbPassword array.
	PasswordLength uint32 `idl:"name:cbPassword" json:"password_length"`
	// pbPassword:  The password.
	Password []byte `idl:"name:pbPassword;size_is:(cbPassword)" json:"password"`
	// cbHashBody:  The count, in bytes, of the pbHashBody array.
	HashBodyLength uint32 `idl:"name:cbHashBody" json:"hash_body_length"`
	// pbHashBody:  The hash of the script value.
	HashBody []byte `idl:"name:pbHashBody;size_is:(cbHashBody)" json:"hash_body"`
	// cbHashSignature:  The count, in bytes, of the pbHashSignature array.
	HashSignatureLength uint32 `idl:"name:cbHashSignature" json:"hash_signature_length"`
	// pbHashSignature:  The script signature.
	HashSignature []byte `idl:"name:pbHashSignature;size_is:(cbHashSignature)" json:"hash_signature"`
}

func (o *MessagePrepareScriptReplyV1) xxx_PreparePayload(ctx context.Context) error {
	if o.Password != nil && o.PasswordLength == 0 {
		o.PasswordLength = uint32(len(o.Password))
	}
	if o.HashBody != nil && o.HashBodyLength == 0 {
		o.HashBodyLength = uint32(len(o.HashBody))
	}
	if o.HashSignature != nil && o.HashSignatureLength == 0 {
		o.HashSignatureLength = uint32(len(o.HashSignature))
	}
	if o.PasswordLength > uint32(1024) {
		return fmt.Errorf("PasswordLength is out of range")
	}
	if o.HashBodyLength > uint32(10485760) {
		return fmt.Errorf("HashBodyLength is out of range")
	}
	if o.HashSignatureLength > uint32(10485760) {
		return fmt.Errorf("HashSignatureLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessagePrepareScriptReplyV1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.OperationStatus); err != nil {
		return err
	}
	if o.ErrorMessage != "" {
		_ptr_pwErrMessage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ErrorMessage); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ErrorMessage, _ptr_pwErrMessage); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PasswordLength); err != nil {
		return err
	}
	if o.Password != nil || o.PasswordLength > 0 {
		_ptr_pbPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.PasswordLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Password {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Password[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Password); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Password, _ptr_pbPassword); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.HashBodyLength); err != nil {
		return err
	}
	if o.HashBody != nil || o.HashBodyLength > 0 {
		_ptr_pbHashBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.HashBodyLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.HashBody {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.HashBody[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.HashBody); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.HashBody, _ptr_pbHashBody); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.HashSignatureLength); err != nil {
		return err
	}
	if o.HashSignature != nil || o.HashSignatureLength > 0 {
		_ptr_pbHashSignature := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.HashSignatureLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.HashSignature {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.HashSignature[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.HashSignature); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.HashSignature, _ptr_pbHashSignature); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessagePrepareScriptReplyV1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperationStatus); err != nil {
		return err
	}
	_ptr_pwErrMessage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ErrorMessage); err != nil {
			return err
		}
		return nil
	})
	_s_pwErrMessage := func(ptr interface{}) { o.ErrorMessage = *ptr.(*string) }
	if err := w.ReadPointer(&o.ErrorMessage, _s_pwErrMessage, _ptr_pwErrMessage); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordLength); err != nil {
		return err
	}
	_ptr_pbPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.PasswordLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.PasswordLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Password", sizeInfo[0])
		}
		o.Password = make([]byte, sizeInfo[0])
		for i1 := range o.Password {
			i1 := i1
			if err := w.ReadData(&o.Password[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbPassword := func(ptr interface{}) { o.Password = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Password, _s_pbPassword, _ptr_pbPassword); err != nil {
		return err
	}
	if err := w.ReadData(&o.HashBodyLength); err != nil {
		return err
	}
	_ptr_pbHashBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.HashBodyLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.HashBodyLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.HashBody", sizeInfo[0])
		}
		o.HashBody = make([]byte, sizeInfo[0])
		for i1 := range o.HashBody {
			i1 := i1
			if err := w.ReadData(&o.HashBody[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbHashBody := func(ptr interface{}) { o.HashBody = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.HashBody, _s_pbHashBody, _ptr_pbHashBody); err != nil {
		return err
	}
	if err := w.ReadData(&o.HashSignatureLength); err != nil {
		return err
	}
	_ptr_pbHashSignature := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.HashSignatureLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.HashSignatureLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.HashSignature", sizeInfo[0])
		}
		o.HashSignature = make([]byte, sizeInfo[0])
		for i1 := range o.HashSignature {
			i1 := i1
			if err := w.ReadData(&o.HashSignature[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbHashSignature := func(ptr interface{}) { o.HashSignature = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.HashSignature, _s_pbHashSignature, _ptr_pbHashSignature); err != nil {
		return err
	}
	return nil
}

// MessagePrepareScriptReply structure represents DSA_MSG_PREPARE_SCRIPT_REPLY RPC union.
//
// The DSA_MSG_PREPARE_SCRIPT_REPLY union defines the response messages received from
// the IDL_DSAPrepareScript method.
type MessagePrepareScriptReply struct {
	// Types that are assignable to Value
	//
	// *MessagePrepareScriptReply_V1
	Value is_MessagePrepareScriptReply `json:"value"`
}

func (o *MessagePrepareScriptReply) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *MessagePrepareScriptReply_V1:
		if value != nil {
			return value.V1
		}
	}
	return nil
}

type is_MessagePrepareScriptReply interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_MessagePrepareScriptReply()
}

func (o *MessagePrepareScriptReply) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *MessagePrepareScriptReply_V1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *MessagePrepareScriptReply) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*MessagePrepareScriptReply_V1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MessagePrepareScriptReply_V1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *MessagePrepareScriptReply) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &MessagePrepareScriptReply_V1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// MessagePrepareScriptReply_V1 structure represents DSA_MSG_PREPARE_SCRIPT_REPLY RPC union arm.
//
// It has following labels: 1
type MessagePrepareScriptReply_V1 struct {
	// V1:  The version 1 response.
	V1 *MessagePrepareScriptReplyV1 `idl:"name:V1" json:"v1"`
}

func (*MessagePrepareScriptReply_V1) is_MessagePrepareScriptReply() {}

func (o *MessagePrepareScriptReply_V1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.V1 != nil {
		if err := o.V1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MessagePrepareScriptReplyV1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessagePrepareScriptReply_V1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.V1 == nil {
		o.V1 = &MessagePrepareScriptReplyV1{}
	}
	if err := o.V1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultDsaopClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultDsaopClient) PrepareScript(ctx context.Context, in *PrepareScriptRequest, opts ...dcerpc.CallOption) (*PrepareScriptResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PrepareScriptResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDsaopClient) ExecuteScript(ctx context.Context, in *ExecuteScriptRequest, opts ...dcerpc.CallOption) (*ExecuteScriptResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ExecuteScriptResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDsaopClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}
func NewDsaopClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DsaopClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DsaopSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultDsaopClient{cc: cc}, nil
}

// xxx_PrepareScriptOperation structure represents the IDL_DSAPrepareScript operation
type xxx_PrepareScriptOperation struct {
	InVersion  uint32                       `idl:"name:dwInVersion" json:"in_version"`
	In         *MessagePrepareScriptRequest `idl:"name:pmsgIn;switch_is:dwInVersion;pointer:ref" json:"in"`
	OutVersion uint32                       `idl:"name:pdwOutVersion;pointer:ref" json:"out_version"`
	Out        *MessagePrepareScriptReply   `idl:"name:pmsgOut;switch_is:*pdwOutVersion;pointer:ref" json:"out"`
	Return     uint32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_PrepareScriptOperation) OpNum() int { return 0 }

func (o *xxx_PrepareScriptOperation) OpName() string { return "/dsaop/v1/IDL_DSAPrepareScript" }

func (o *xxx_PrepareScriptOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareScriptOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwInVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InVersion); err != nil {
			return err
		}
	}
	// pmsgIn {in} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DSA_MSG_PREPARE_SCRIPT_REQ}(union))
	{
		_swIn := uint32(o.InVersion)
		if o.In != nil {
			if err := o.In.MarshalUnionNDR(ctx, w, _swIn); err != nil {
				return err
			}
		} else {
			if err := (&MessagePrepareScriptRequest{}).MarshalUnionNDR(ctx, w, _swIn); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_PrepareScriptOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwInVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InVersion); err != nil {
			return err
		}
	}
	// pmsgIn {in} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DSA_MSG_PREPARE_SCRIPT_REQ}(union))
	{
		if o.In == nil {
			o.In = &MessagePrepareScriptRequest{}
		}
		_swIn := uint32(o.InVersion)
		if err := o.In.UnmarshalUnionNDR(ctx, w, _swIn); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareScriptOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareScriptOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwOutVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutVersion); err != nil {
			return err
		}
	}
	// pmsgOut {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DSA_MSG_PREPARE_SCRIPT_REPLY}(union))
	{
		_swOut := uint32(o.OutVersion)
		if o.Out != nil {
			if err := o.Out.MarshalUnionNDR(ctx, w, _swOut); err != nil {
				return err
			}
		} else {
			if err := (&MessagePrepareScriptReply{}).MarshalUnionNDR(ctx, w, _swOut); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareScriptOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwOutVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutVersion); err != nil {
			return err
		}
	}
	// pmsgOut {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DSA_MSG_PREPARE_SCRIPT_REPLY}(union))
	{
		if o.Out == nil {
			o.Out = &MessagePrepareScriptReply{}
		}
		_swOut := uint32(o.OutVersion)
		if err := o.Out.UnmarshalUnionNDR(ctx, w, _swOut); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PrepareScriptRequest structure represents the IDL_DSAPrepareScript operation request
type PrepareScriptRequest struct {
	// dwInVersion: The version of the request message.
	InVersion uint32 `idl:"name:dwInVersion" json:"in_version"`
	// pmsgIn: A pointer to the request message.
	In *MessagePrepareScriptRequest `idl:"name:pmsgIn;switch_is:dwInVersion;pointer:ref" json:"in"`
}

func (o *PrepareScriptRequest) xxx_ToOp(ctx context.Context) *xxx_PrepareScriptOperation {
	if o == nil {
		return &xxx_PrepareScriptOperation{}
	}
	return &xxx_PrepareScriptOperation{
		InVersion: o.InVersion,
		In:        o.In,
	}
}

func (o *PrepareScriptRequest) xxx_FromOp(ctx context.Context, op *xxx_PrepareScriptOperation) {
	if o == nil {
		return
	}
	o.InVersion = op.InVersion
	o.In = op.In
}
func (o *PrepareScriptRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PrepareScriptRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PrepareScriptOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PrepareScriptResponse structure represents the IDL_DSAPrepareScript operation response
type PrepareScriptResponse struct {
	// pdwOutVersion: A pointer to the version of the response message.
	OutVersion uint32 `idl:"name:pdwOutVersion;pointer:ref" json:"out_version"`
	// pmsgOut: A pointer to the response message.
	Out *MessagePrepareScriptReply `idl:"name:pmsgOut;switch_is:*pdwOutVersion;pointer:ref" json:"out"`
	// Return: The IDL_DSAPrepareScript return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PrepareScriptResponse) xxx_ToOp(ctx context.Context) *xxx_PrepareScriptOperation {
	if o == nil {
		return &xxx_PrepareScriptOperation{}
	}
	return &xxx_PrepareScriptOperation{
		OutVersion: o.OutVersion,
		Out:        o.Out,
		Return:     o.Return,
	}
}

func (o *PrepareScriptResponse) xxx_FromOp(ctx context.Context, op *xxx_PrepareScriptOperation) {
	if o == nil {
		return
	}
	o.OutVersion = op.OutVersion
	o.Out = op.Out
	o.Return = op.Return
}
func (o *PrepareScriptResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PrepareScriptResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PrepareScriptOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExecuteScriptOperation structure represents the IDL_DSAExecuteScript operation
type xxx_ExecuteScriptOperation struct {
	InVersion  uint32                       `idl:"name:dwInVersion" json:"in_version"`
	In         *MessageExecuteScriptRequest `idl:"name:pmsgIn;switch_is:dwInVersion;pointer:ref" json:"in"`
	OutVersion uint32                       `idl:"name:pdwOutVersion;pointer:ref" json:"out_version"`
	Out        *MessageExecuteScriptReply   `idl:"name:pmsgOut;switch_is:*pdwOutVersion;pointer:ref" json:"out"`
	Return     uint32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_ExecuteScriptOperation) OpNum() int { return 1 }

func (o *xxx_ExecuteScriptOperation) OpName() string { return "/dsaop/v1/IDL_DSAExecuteScript" }

func (o *xxx_ExecuteScriptOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecuteScriptOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwInVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InVersion); err != nil {
			return err
		}
	}
	// pmsgIn {in} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DSA_MSG_EXECUTE_SCRIPT_REQ}(union))
	{
		_swIn := uint32(o.InVersion)
		if o.In != nil {
			if err := o.In.MarshalUnionNDR(ctx, w, _swIn); err != nil {
				return err
			}
		} else {
			if err := (&MessageExecuteScriptRequest{}).MarshalUnionNDR(ctx, w, _swIn); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecuteScriptOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwInVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InVersion); err != nil {
			return err
		}
	}
	// pmsgIn {in} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DSA_MSG_EXECUTE_SCRIPT_REQ}(union))
	{
		if o.In == nil {
			o.In = &MessageExecuteScriptRequest{}
		}
		_swIn := uint32(o.InVersion)
		if err := o.In.UnmarshalUnionNDR(ctx, w, _swIn); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecuteScriptOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecuteScriptOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwOutVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutVersion); err != nil {
			return err
		}
	}
	// pmsgOut {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DSA_MSG_EXECUTE_SCRIPT_REPLY}(union))
	{
		_swOut := uint32(o.OutVersion)
		if o.Out != nil {
			if err := o.Out.MarshalUnionNDR(ctx, w, _swOut); err != nil {
				return err
			}
		} else {
			if err := (&MessageExecuteScriptReply{}).MarshalUnionNDR(ctx, w, _swOut); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExecuteScriptOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwOutVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutVersion); err != nil {
			return err
		}
	}
	// pmsgOut {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DSA_MSG_EXECUTE_SCRIPT_REPLY}(union))
	{
		if o.Out == nil {
			o.Out = &MessageExecuteScriptReply{}
		}
		_swOut := uint32(o.OutVersion)
		if err := o.Out.UnmarshalUnionNDR(ctx, w, _swOut); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ExecuteScriptRequest structure represents the IDL_DSAExecuteScript operation request
type ExecuteScriptRequest struct {
	// dwInVersion: The version of the request message.
	InVersion uint32 `idl:"name:dwInVersion" json:"in_version"`
	// pmsgIn: A pointer to the request message.
	In *MessageExecuteScriptRequest `idl:"name:pmsgIn;switch_is:dwInVersion;pointer:ref" json:"in"`
}

func (o *ExecuteScriptRequest) xxx_ToOp(ctx context.Context) *xxx_ExecuteScriptOperation {
	if o == nil {
		return &xxx_ExecuteScriptOperation{}
	}
	return &xxx_ExecuteScriptOperation{
		InVersion: o.InVersion,
		In:        o.In,
	}
}

func (o *ExecuteScriptRequest) xxx_FromOp(ctx context.Context, op *xxx_ExecuteScriptOperation) {
	if o == nil {
		return
	}
	o.InVersion = op.InVersion
	o.In = op.In
}
func (o *ExecuteScriptRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExecuteScriptRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecuteScriptOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExecuteScriptResponse structure represents the IDL_DSAExecuteScript operation response
type ExecuteScriptResponse struct {
	// pdwOutVersion: A pointer to the version of the response message.
	OutVersion uint32 `idl:"name:pdwOutVersion;pointer:ref" json:"out_version"`
	// pmsgOut: A pointer to the response message.
	Out *MessageExecuteScriptReply `idl:"name:pmsgOut;switch_is:*pdwOutVersion;pointer:ref" json:"out"`
	// Return: The IDL_DSAExecuteScript return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ExecuteScriptResponse) xxx_ToOp(ctx context.Context) *xxx_ExecuteScriptOperation {
	if o == nil {
		return &xxx_ExecuteScriptOperation{}
	}
	return &xxx_ExecuteScriptOperation{
		OutVersion: o.OutVersion,
		Out:        o.Out,
		Return:     o.Return,
	}
}

func (o *ExecuteScriptResponse) xxx_FromOp(ctx context.Context, op *xxx_ExecuteScriptOperation) {
	if o == nil {
		return
	}
	o.OutVersion = op.OutVersion
	o.Out = op.Out
	o.Return = op.Return
}
func (o *ExecuteScriptResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExecuteScriptResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExecuteScriptOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
