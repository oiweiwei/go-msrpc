package errors

var (
	NCAAddrError        = &RPCError{0x1C000002, "nca_s_fault_addr_error", ""}
	NCACancel           = &RPCError{0x1C00000D, "nca_s_fault_cancel", ""}
	NCACodesetConvError = &RPCError{0x1C000023, "nca_s_fault_codeset_conv_error", ""}
	NCAContextMismatch  = &RPCError{0x1C00001A, "nca_s_fault_context_mismatch", "The context handle does not match any known context handles."}
	NCAFpDivZero        = &RPCError{0x1C000003, "nca_s_fault_fp_div_zero", ""}
	NCAFpError          = &RPCError{0x1C00000F, "nca_s_fault_fp_error", ""}
	NCAFpOverflow       = &RPCError{0x1C000005, "nca_s_fault_fp_overflow", ""}
	NCAFpUnderflow      = &RPCError{0x1C000004, "nca_s_fault_fp_underflow", ""}
	NCAIllInst          = &RPCError{0x1C00000E, "nca_s_fault_ill_inst", ""}
	NCAIntDivByZero     = &RPCError{0x1C000001, "nca_s_fault_int_div_by_zero", ""}
	NCAIntOverflow      = &RPCError{0x1C000010, "nca_s_fault_int_overflow", ""}
	NCAInvalidBound     = &RPCError{0x1C000007, "nca_s_fault_invalid_bound", ""}
	NCAInvalidTag       = &RPCError{0x1C000006, "nca_s_fault_invalid_tag", ""}
	NCANoClientStub     = &RPCError{0x1C000025, "nca_s_fault_no_client_stub", ""}
	NCAObjectNotFound   = &RPCError{0x1C000024, "nca_s_fault_object_not_found", ""}
	NCAPipeClosed       = &RPCError{0x1C000015, "nca_s_fault_pipe_closed", ""}
	NCAPipeCommError    = &RPCError{0x1C000018, "nca_s_fault_pipe_comm_error", ""}
	NCAPipeDiscipline   = &RPCError{0x1C000017, "nca_s_fault_pipe_discipline", ""}
	NCAPipeEmpty        = &RPCError{0x1C000014, "nca_s_fault_pipe_empty", ""}
	NCAPipeMemory       = &RPCError{0x1C000019, "nca_s_fault_pipe_memory", ""}
	NCAPipeOrder        = &RPCError{0x1C000016, "nca_s_fault_pipe_order", ""}
	NCARemoteNoMemory   = &RPCError{0x1C00001B, "nca_s_fault_remote_no_memory", ""}
	NCATxOpenFailed     = &RPCError{0x1C000022, "nca_s_fault_tx_open_failed", ""}
	NCSUserDefined      = &RPCError{0x1C000021, "ncs_s_fault_user_defined", ""}
	RPCCallCancelled    = &RPCError{0x16C9A031, "rpc_s_call_cancelled", ""}
	RPCAddrError        = &RPCError{0x16C9A074, "rpc_s_fault_addr_error", ""}
	RPCCodesetConvError = &RPCError{0x16C9A16E, "rpc_s_fault_codeset_conv_error", ""}
	RPCContextMismatch  = &RPCError{0x16C9A075, "rpc_s_fault_context_mismatch", "The context handle does not match any known context handles."}
	RPCFpDivByZero      = &RPCError{0x16C9A076, "rpc_s_fault_fp_div_by_zero", ""}
	RPCFpError          = &RPCError{0x16C9A077, "rpc_s_fault_fp_error", ""}
	RPCFpOverflow       = &RPCError{0x16C9A078, "rpc_s_fault_fp_overflow", ""}
	RPCFpUnderflow      = &RPCError{0x16C9A079, "rpc_s_fault_fp_underflow", ""}
	RPCIllInst          = &RPCError{0x16C9A07A, "rpc_s_fault_ill_inst", ""}
	RPCIntDivByZero     = &RPCError{0x16C9A07B, "rpc_s_fault_int_div_by_zero", ""}
	RPCIntOverflow      = &RPCError{0x16C9A07C, "rpc_s_fault_int_overflow", ""}
	RPCInvalidBound     = &RPCError{0x16C9A07D, "rpc_s_fault_invalid_bound", ""}
	RPCInvalidTag       = &RPCError{0x16C9A07E, "rpc_s_fault_invalid_tag", ""}
	RPCNoClientStub     = &RPCError{0x16C9A170, "rpc_s_fault_no_client_stub", ""}
	RPCObjectNotFound   = &RPCError{0x16C9A01B, "rpc_s_fault_object_not_found", ""}
	RPCPipeClosed       = &RPCError{0x16C9A07F, "rpc_s_fault_pipe_closed", ""}
	RPCPipeCommError    = &RPCError{0x16C9A080, "rpc_s_fault_pipe_comm_error", ""}
	RPCPipeDiscipline   = &RPCError{0x16C9A081, "rpc_s_fault_pipe_discipline", ""}
	RPCPipeEmpty        = &RPCError{0x16C9A082, "rpc_s_fault_pipe_empty", ""}
	RPCPipeMemory       = &RPCError{0x16C9A083, "rpc_s_fault_pipe_memory", ""}
	RPCPipeOrder        = &RPCError{0x16C9A084, "rpc_s_fault_pipe_order", ""}
	RPCRemoteNoMemory   = &RPCError{0x16C9A086, "rpc_s_fault_remote_no_memory", ""}
	RPCTxOpenFailed     = &RPCError{0x16C9A116, "rpc_s_fault_tx_open_failed", ""}
	RPCUnspec           = &RPCError{0x16C9A087, "rpc_s_fault_unspec", ""}
	RPCUserDefined      = &RPCError{0x16C9A113, "rpc_s_fault_user_defined", ""}
)

func FaultFromCode(code uint32) error {

	switch code {
	case 0x1C000002:
		return NCAAddrError
	case 0x1C00000D:
		return NCACancel
	case 0x1C000023:
		return NCACodesetConvError
	case 0x1C00001A:
		return NCAContextMismatch
	case 0x1C000003:
		return NCAFpDivZero
	case 0x1C00000F:
		return NCAFpError
	case 0x1C000005:
		return NCAFpOverflow
	case 0x1C000004:
		return NCAFpUnderflow
	case 0x1C00000E:
		return NCAIllInst
	case 0x1C000001:
		return NCAIntDivByZero
	case 0x1C000010:
		return NCAIntOverflow
	case 0x1C000007:
		return NCAInvalidBound
	case 0x1C000006:
		return NCAInvalidTag
	case 0x1C000025:
		return NCANoClientStub
	case 0x1C000024:
		return NCAObjectNotFound
	case 0x1C000015:
		return NCAPipeClosed
	case 0x1C000018:
		return NCAPipeCommError
	case 0x1C000017:
		return NCAPipeDiscipline
	case 0x1C000014:
		return NCAPipeEmpty
	case 0x1C000019:
		return NCAPipeMemory
	case 0x1C000016:
		return NCAPipeOrder
	case 0x1C00001B:
		return NCARemoteNoMemory
	case 0x1C000022:
		return NCATxOpenFailed
	case 0x1C000021:
		return NCSUserDefined
	case 0x16C9A031:
		return RPCCallCancelled
	case 0x16C9A074:
		return RPCAddrError
	case 0x16C9A16E:
		return RPCCodesetConvError
	case 0x16C9A075:
		return RPCContextMismatch
	case 0x16C9A076:
		return RPCFpDivByZero
	case 0x16C9A077:
		return RPCFpError
	case 0x16C9A078:
		return RPCFpOverflow
	case 0x16C9A079:
		return RPCFpUnderflow
	case 0x16C9A07A:
		return RPCIllInst
	case 0x16C9A07B:
		return RPCIntDivByZero
	case 0x16C9A07C:
		return RPCIntOverflow
	case 0x16C9A07D:
		return RPCInvalidBound
	case 0x16C9A07E:
		return RPCInvalidTag
	case 0x16C9A170:
		return RPCNoClientStub
	case 0x16C9A01B:
		return RPCObjectNotFound
	case 0x16C9A07F:
		return RPCPipeClosed
	case 0x16C9A080:
		return RPCPipeCommError
	case 0x16C9A081:
		return RPCPipeDiscipline
	case 0x16C9A082:
		return RPCPipeEmpty
	case 0x16C9A083:
		return RPCPipeMemory
	case 0x16C9A084:
		return RPCPipeOrder
	case 0x16C9A086:
		return RPCRemoteNoMemory
	case 0x16C9A116:
		return RPCTxOpenFailed
	case 0x16C9A087:
		return RPCUnspec
	case 0x16C9A113:
		return RPCUserDefined
	}

	return nil
}
