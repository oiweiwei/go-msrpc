// The even6 package implements the EVEN6 client protocol.
//
// # Introduction
//
// The EventLog Remoting Protocol Version 6.0, originally available in the Windows Vista
// operating system, is a remote procedure call (RPC)–based protocol that exposes
// RPC methods for reading events in both live event logs and backup event logs on remote
// computers. This protocol also specifies how to get general information for a log,
// such as number of records in the log, oldest records in the log, and if the log is
// full. It may also be used for clearing and backing up both types of event logs.
//
// # Overview
package even6

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
	GoPackage = "even6"
)

// MaxPayload represents the MAX_PAYLOAD RPC constant
var MaxPayload = 2097152

// MaxRPCQueryLength represents the MAX_RPC_QUERY_LENGTH RPC constant
var MaxRPCQueryLength = 1048576

// MaxRPCChannelNameLength represents the MAX_RPC_CHANNEL_NAME_LENGTH RPC constant
var MaxRPCChannelNameLength = 512

// MaxRPCQueryChannelSize represents the MAX_RPC_QUERY_CHANNEL_SIZE RPC constant
var MaxRPCQueryChannelSize = 512

// MaxRPCEventIDSize represents the MAX_RPC_EVENT_ID_SIZE RPC constant
var MaxRPCEventIDSize = 256

// MaxRPCFilePathLength represents the MAX_RPC_FILE_PATH_LENGTH RPC constant
var MaxRPCFilePathLength = 32768

// MaxRPCChannelPathLength represents the MAX_RPC_CHANNEL_PATH_LENGTH RPC constant
var MaxRPCChannelPathLength = 32768

// MaxRPCBookmarkLength represents the MAX_RPC_BOOKMARK_LENGTH RPC constant
var MaxRPCBookmarkLength = 1048576

// MaxRPCPublisherIDLength represents the MAX_RPC_PUBLISHER_ID_LENGTH RPC constant
var MaxRPCPublisherIDLength = 2048

// MaxRPCPropertyBufferSize represents the MAX_RPC_PROPERTY_BUFFER_SIZE RPC constant
var MaxRPCPropertyBufferSize = 2097152

// MaxRPCFilterLength represents the MAX_RPC_FILTER_LENGTH RPC constant
var MaxRPCFilterLength = 1048576

// MaxRPCRecordCount represents the MAX_RPC_RECORD_COUNT RPC constant
var MaxRPCRecordCount = 1024

// MaxRPCEventSize represents the MAX_RPC_EVENT_SIZE RPC constant
var MaxRPCEventSize = 2097152

// MaxRPCBatchSize represents the MAX_RPC_BATCH_SIZE RPC constant
var MaxRPCBatchSize = 2097152

// MaxRPCRenderedStringSize represents the MAX_RPC_RENDERED_STRING_SIZE RPC constant
var MaxRPCRenderedStringSize = 2097152

// MaxRPCChannelCount represents the MAX_RPC_CHANNEL_COUNT RPC constant
var MaxRPCChannelCount = 8192

// MaxRPCPublisherCount represents the MAX_RPC_PUBLISHER_COUNT RPC constant
var MaxRPCPublisherCount = 8192

// MaxRPCEventMetadataCount represents the MAX_RPC_EVENT_METADATA_COUNT RPC constant
var MaxRPCEventMetadataCount = 256

// MaxRPCVariantListCount represents the MAX_RPC_VARIANT_LIST_COUNT RPC constant
var MaxRPCVariantListCount = 256

// MaxRPCBoolArrayCount represents the MAX_RPC_BOOL_ARRAY_COUNT RPC constant
var MaxRPCBoolArrayCount = 524288

// MaxRPCUint32ArrayCount represents the MAX_RPC_UINT32_ARRAY_COUNT RPC constant
var MaxRPCUint32ArrayCount = 524288

// MaxRPCUint64ArrayCount represents the MAX_RPC_UINT64_ARRAY_COUNT RPC constant
var MaxRPCUint64ArrayCount = 262144

// MaxRPCStringArrayCount represents the MAX_RPC_STRING_ARRAY_COUNT RPC constant
var MaxRPCStringArrayCount = 4096

// MaxRPCGUIDArrayCount represents the MAX_RPC_GUID_ARRAY_COUNT RPC constant
var MaxRPCGUIDArrayCount = 131072

// MaxRPCStringLength represents the MAX_RPC_STRING_LENGTH RPC constant
var MaxRPCStringLength = 1048576
