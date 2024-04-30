package errors

var (
	RPCVersionMismatch           = &RPCError{0x1c000008, "nca_rpc_version_mismatch", "The server does not support the RPC protocol version specified in the request PDU."}
	UnspecifiedReject            = &RPCError{0x1c000009, "nca_unspec_reject", "The request is being rejected for unspecified reasons."}
	ServerBadActivityID          = &RPCError{0x1c00000A, "nca_s_bad_actid", "The server has no state corresponding to the activity identifier in the message."}
	WhoAreYouFailed              = &RPCError{0x1c00000b, "nca_who_are_you_failed", "The Conversation Manager callback failed."}
	ManagerNotEntered            = &RPCError{0x1c00000c, "nca_manager_not_entered", "The server manager routine has not been entered and executed."}
	OperationRangeError          = &RPCError{0x1c010002, "nca_op_rng_error", "The operation number passed in the request PDU is greater than or equal to the number of operations in the interface."}
	UnknownInterface             = &RPCError{0x1c010003, "nca_unk_if", "The server does not export the requested interface."}
	WrongBootTime                = &RPCError{0x1c010006, "nca_wrong_boot_time", "The server boot time passed in the request PDU does not match the actual server boot time."}
	ServerYouCrashed             = &RPCError{0x1c010009, "nca_s_you_crashed", "A restarted server called back a client."}
	ProtocolError                = &RPCError{0x1c01000b, "nca_proto_error", "The RPC client or server protocol has been violated."}
	OutputArgsTooBig             = &RPCError{0x1c010013, "nca_out_args_too_big", "The output parameters of the operation exceed their declared maximum size."}
	ServerTooBusy                = &RPCError{0x1c010014, "nca_server_too_busy", "The server is too busy to handle the call."}
	UnsupportedType              = &RPCError{0x1c010017, "nca_unsupported_type", "The server does not implement the requested operation for the type of the requested object."}
	InvalidPresentationContextID = &RPCError{0x1c00001c, "nca_invalid_pres_context_id", "Invalid presentation context ID."}
	UnsupportedAuthnLevel        = &RPCError{0x1c00001d, "nca_unsupported_authn_level", "The server did not support the requested authentication level."}
	InvalidChecksum              = &RPCError{0x1c00001f, "nca_invalid_checksum", "Invalid checksum."}
	InvalidCRC                   = &RPCError{0x1c000020, "nca_invalid_crc", "Invalid CRC."}
)

func FromCode(code uint32) error {
	switch code {
	case 0x1c000008:
		return RPCVersionMismatch
	case 0x1c000009:
		return UnspecifiedReject
	case 0x1c00000A:
		return ServerBadActivityID
	case 0x1c00000b:
		return WhoAreYouFailed
	case 0x1c00000c:
		return ManagerNotEntered
	case 0x1c010002:
		return OperationRangeError
	case 0x1c010003:
		return UnknownInterface
	case 0x1c010006:
		return WrongBootTime
	case 0x1c010009:
		return ServerYouCrashed
	case 0x1c01000b:
		return ProtocolError
	case 0x1c010013:
		return OutputArgsTooBig
	case 0x1c010014:
		return ServerTooBusy
	case 0x1c010017:
		return UnsupportedType
	case 0x1c00001c:
		return InvalidPresentationContextID
	case 0x1c00001d:
		return UnsupportedAuthnLevel
	case 0x1c00001f:
		return InvalidChecksum
	case 0x1c000020:
		return InvalidCRC
	}

	return nil
}
