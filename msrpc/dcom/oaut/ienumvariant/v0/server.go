package ienumvariant

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

// IEnumVARIANT server interface.
type EnumVariantServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The IEnumVARIANT::Next method returns up to the number of requested items that occur
	// next in the enumeration sequence.
	//
	// Return Values: The method MUST return the information in an HRESULT data structure,
	// which is defined in [MS-ERREF] section 2.1. The severity bit in the structure identifies
	// the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// A static or semistatic IEnumVARIANT server MUST retrieve the next celt elements from
	// the sequence and fill in rgVar up to the celt elements or up to the remaining number
	// of elements that are not yet enumerated.
	//
	// A dynamic server MUST use its server-specific state to retrieve the next elements.
	//
	// In all cases, the server MUST:
	//
	// * Set pCeltFetched with the number retrieved.
	//
	// * Update the current position in the sequence.
	//
	// * Return a status of 1 (S_FALSE) if pCeltFetched is not equal to celt.
	Next(context.Context, *NextRequest) (*NextResponse, error)

	// The IEnumVARIANT::Skip method skips over the requested number of elements in the
	// enumeration sequence.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// A static or semi-static IEnumVARIANT server MUST update the current position in the
	// sequence by either celt elements, or the number of elements remaining, whichever
	// is smaller.
	//
	// A dynamic server MUST use its server-specific state to affect the dynamic collection
	// it manages, and MUST update the current position in the sequence by either celt elements,
	// or the number of elements remaining, whichever is smaller.
	//
	// In all cases, the server MUST return 1 (S_FALSE), if the current position was updated
	// by less than celt elements.
	Skip(context.Context, *SkipRequest) (*SkipResponse, error)

	// The IEnumVARIANT::Reset method resets the enumeration sequence to the beginning.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// A static server MUST reset its current position in the sequence.
	//
	// A semi-static or dynamic server MUST update the sequence of elements it maintains
	// and MUST reset the current position in the sequence to the beginning.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)

	// The IEnumVARIANT::Clone method creates a copy of the current state of the enumeration.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)
}

func RegisterEnumVariantServer(conn dcerpc.Conn, o EnumVariantServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEnumVariantServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EnumVariantSyntaxV0_0))...)
}

func NewEnumVariantServerHandle(o EnumVariantServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EnumVariantServerHandle(ctx, o, opNum, r)
	}
}

func EnumVariantServerHandle(ctx context.Context, o EnumVariantServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Next
		in := &NextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Next(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Skip
		in := &SkipRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Skip(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Reset
		in := &ResetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Reset(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // Clone
		in := &CloneRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Clone(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
