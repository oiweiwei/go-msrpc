package itransactionstream

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// ITransactionStream server interface.
type TransactionStreamServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method returns the STxInfo ([MS-DTCO] section 2.2.5.10) of the currently active
	// transaction and the CurrentTSN value.
	//
	// Return Values: The method MUST return a positive value or zero, to indicate successful
	// completion, or a negative value to indicate failure. The client MUST treat any negative
	// return value as a fatal error.
	GetSeqAndTxViaExport(context.Context, *GetSeqAndTxViaExportRequest) (*GetSeqAndTxViaExportResponse, error)

	// This method returns the Propagation_Token (as specified in [MS-DTCO] section 2.2.5.4)
	// of the currently active transaction and the CurrentTSN value.
	//
	// Return Values: The method MUST return a positive value or zero, to indicate successful
	// completion, or a negative value to indicate failure. The client MUST treat any negative
	// return value as a fatal error.
	GetSeqAndTxViaTransmitter(context.Context, *GetSeqAndTxViaTransmitterRequest) (*GetSeqAndTxViaTransmitterResponse, error)

	// This method returns the STxInfo instance (as specified in [MS-DTCO] section 2.2.5.10)
	// of the currently active transaction or returns an error if the specified TSN is not
	// the same as the CurrentTSN value.
	//
	// Return Values: The method MUST return a positive value or zero to indicate successful
	// completion, or a negative value to indicate failure. The client MUST treat any negative
	// return value as a fatal error.
	GetTxViaExport(context.Context, *GetTxViaExportRequest) (*GetTxViaExportResponse, error)

	// This method returns the Propagation_Token ([MS-DTCO] section 2.2.5.4) of the currently
	// active transaction, or returns an error if the specified TSN is not the same as the
	// CurrentTSN value.
	//
	// Return Values: The method MUST return a positive value or zero to indicate successful
	// completion, or a negative value to indicate failure. The client MUST treat any negative
	// return value as a fatal error.
	GetTxViaTransmitter(context.Context, *GetTxViaTransmitterRequest) (*GetTxViaTransmitterResponse, error)
}

func RegisterTransactionStreamServer(conn dcerpc.Conn, o TransactionStreamServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTransactionStreamServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TransactionStreamSyntaxV0_0))...)
}

func NewTransactionStreamServerHandle(o TransactionStreamServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TransactionStreamServerHandle(ctx, o, opNum, r)
	}
}

func TransactionStreamServerHandle(ctx context.Context, o TransactionStreamServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetSeqAndTxViaExport
		in := &GetSeqAndTxViaExportRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSeqAndTxViaExport(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // GetSeqAndTxViaTransmitter
		in := &GetSeqAndTxViaTransmitterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSeqAndTxViaTransmitter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // GetTxViaExport
		in := &GetTxViaExportRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTxViaExport(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // GetTxViaTransmitter
		in := &GetTxViaTransmitterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTxViaTransmitter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
