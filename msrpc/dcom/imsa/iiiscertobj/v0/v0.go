package iiiscertobj

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
)

var (
	// import guard
	GoPackage = "dcom/imsa"
)

var (
	// IIISCertObj interface identifier bd0c73bc-805b-4043-9c30-9a28d64dd7d2
	IISCertObjectIID = &dcom.IID{Data1: 0xbd0c73bc, Data2: 0x805b, Data3: 0x4043, Data4: []byte{0x9c, 0x30, 0x9a, 0x28, 0xd6, 0x4d, 0xd7, 0xd2}}
	// Syntax UUID
	IISCertObjectSyntaxUUID = &uuid.UUID{TimeLow: 0xbd0c73bc, TimeMid: 0x805b, TimeHiAndVersion: 0x4043, ClockSeqHiAndReserved: 0x9c, ClockSeqLow: 0x30, Node: [6]uint8{0x9a, 0x28, 0xd6, 0x4d, 0xd7, 0xd2}}
	// Syntax ID
	IISCertObjectSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: IISCertObjectSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IIISCertObj interface.
type IISCertObjectClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// Opnum7NotUsedOnWire operation.
	// SetOpnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// SetOpnum8NotUsedOnWire

	// Opnum9NotUsedOnWire operation.
	// SetOpnum9NotUsedOnWire

	// The InstanceName method sets the web server instance to be used by subsequent method
	// calls.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [[MS-ERREF].
	//
	//	+----------------------------------+------------------------------------+
	//	|              RETURN              |                                    |
	//	|            VALUE/CODE            |            DESCRIPTION             |
	//	|                                  |                                    |
	//	+----------------------------------+------------------------------------+
	//	+----------------------------------+------------------------------------+
	//	| 0x00000000 S_OK                  | The call was successful.           |
	//	+----------------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | One or more arguments are invalid. |
	//	+----------------------------------+------------------------------------+
	//	| 0x000006cf RPC_S_STRING_TOO_LONG | The string is too long.            |
	//	+----------------------------------+------------------------------------+
	//
	// The opnum field value for this method is 10.
	SetInstanceName(context.Context, *SetInstanceNameRequest, ...dcerpc.CallOption) (*SetInstanceNameResponse, error)

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire

	// The IsInstalledRemote method determines if a certificate is associated with the specified
	// InstanceName.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.           |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// The opnum field value for this method is 12.
	IsInstalledRemote(context.Context, *IsInstalledRemoteRequest, ...dcerpc.CallOption) (*IsInstalledRemoteResponse, error)

	// Opnum13NotUsedOnWire operation.
	// Opnum13NotUsedOnWire

	// The IsExportableRemote method determines whether the server certificate associated
	// with InstanceName can be exported.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.           |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// The opnum field value for this method is 14.
	IsExportableRemote(context.Context, *IsExportableRemoteRequest, ...dcerpc.CallOption) (*IsExportableRemoteResponse, error)

	// Opnum15NotUsedOnWire operation.
	// Opnum15NotUsedOnWire

	// The GetCertInfoRemote method retrieves properties from a certificate associated with
	// the specified InstanceName.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------------+------------------------------------------------+
	//	|         RETURN          |                                                |
	//	|       VALUE/CODE        |                  DESCRIPTION                   |
	//	|                         |                                                |
	//	+-------------------------+------------------------------------------------+
	//	+-------------------------+------------------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.                       |
	//	+-------------------------+------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid.             |
	//	+-------------------------+------------------------------------------------+
	//	| 0x00000001 S_FALSE      | The call was successful. No data was returned. |
	//	+-------------------------+------------------------------------------------+
	//
	// The opnum field value for this method is 16.
	GetCertInfoRemote(context.Context, *GetCertInfoRemoteRequest, ...dcerpc.CallOption) (*GetCertInfoRemoteResponse, error)

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire

	// Opnum18NotUsedOnWire operation.
	// Opnum18NotUsedOnWire

	// Opnum19NotUsedOnWire operation.
	// Opnum19NotUsedOnWire

	// Opnum20NotUsedOnWire operation.
	// Opnum20NotUsedOnWire

	// Opnum21NotUsedOnWire operation.
	// Opnum21NotUsedOnWire

	// The ImportFromBlob method imports a previously exported certificate blob on the target
	// machine.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+----------------------------------+----------------------------------------+
	//	|              RETURN              |                                        |
	//	|            VALUE/CODE            |              DESCRIPTION               |
	//	|                                  |                                        |
	//	+----------------------------------+----------------------------------------+
	//	+----------------------------------+----------------------------------------+
	//	| 0x00000000 S_OK                  | The call was successful.               |
	//	+----------------------------------+----------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | One or more arguments are invalid.     |
	//	+----------------------------------+----------------------------------------+
	//	| 0x000006cf RPC_S_STRING_TOO_LONG | The string is too long.                |
	//	+----------------------------------+----------------------------------------+
	//	| 0x80092005 CRYPT_E_EXISTS        | The object or property already exists. |
	//	+----------------------------------+----------------------------------------+
	//
	// The opnum field value for this method is 22.
	ImportFromBlob(context.Context, *ImportFromBlobRequest, ...dcerpc.CallOption) (*ImportFromBlobResponse, error)

	// The ImportFromBlobGetHash method imports a previously exported certificate blob on
	// the target machine. In addition to data returned by method ImportFromBlob, this method
	// returns certificate hash and certificate hash buffer size in client-provided parameters
	// pcbCertHashSize and pCertHash. The server MUST allocate memory for the hash buffer
	// and assign this memory block to pCertHash. Size of required buffer is assigned to
	// pcbCertHashSize. If client will pass pCertHash equal to NULL, hash data will not
	// be returned.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+----------------------------------+----------------------------------------+
	//	|              RETURN              |                                        |
	//	|            VALUE/CODE            |              DESCRIPTION               |
	//	|                                  |                                        |
	//	+----------------------------------+----------------------------------------+
	//	+----------------------------------+----------------------------------------+
	//	| 0x00000000 S_OK                  | The call was successful.               |
	//	+----------------------------------+----------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | One or more arguments are invalid.     |
	//	+----------------------------------+----------------------------------------+
	//	| 0x000006cf RPC_S_STRING_TOO_LONG | The string is too long.                |
	//	+----------------------------------+----------------------------------------+
	//	| 0x80092005 CRYPT_E_EXISTS        | The object or property already exists. |
	//	+----------------------------------+----------------------------------------+
	//
	// The opnum field value for this method is 23.
	ImportFromBlobGetHash(context.Context, *ImportFromBlobGetHashRequest, ...dcerpc.CallOption) (*ImportFromBlobGetHashResponse, error)

	// Opnum24NotUsedOnWire operation.
	// Opnum24NotUsedOnWire

	// The ExportToBlob method exports the certificate referenced at InstanceName to a memory
	// buffer.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+------------------------------------+-------------------------------------------------------+
	//	|               RETURN               |                                                       |
	//	|             VALUE/CODE             |                      DESCRIPTION                      |
	//	|                                    |                                                       |
	//	+------------------------------------+-------------------------------------------------------+
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                              |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG            | One or more arguments are invalid.                    |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x000006cf RPC_S_STRING_TOO_LONG   | The string is too long.                               |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x800CC801 MD_ERROR_DATA_NOT_FOUND | The specified metadata was not found.                 |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80092004 CRYPT_E_NOT_FOUND       | Cannot find object or property.                       |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80090349 SEC_E_CERT_WRONG_USAGE  | The certificate is not valid for the requested usage. |
	//	+------------------------------------+-------------------------------------------------------+
	//
	// The opnum field value for this method is 25.
	ExportToBlob(context.Context, *ExportToBlobRequest, ...dcerpc.CallOption) (*ExportToBlobResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) IISCertObjectClient
}

type xxx_DefaultIISCertObjectClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultIISCertObjectClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultIISCertObjectClient) SetInstanceName(ctx context.Context, in *SetInstanceNameRequest, opts ...dcerpc.CallOption) (*SetInstanceNameResponse, error) {
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
	out := &SetInstanceNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISCertObjectClient) IsInstalledRemote(ctx context.Context, in *IsInstalledRemoteRequest, opts ...dcerpc.CallOption) (*IsInstalledRemoteResponse, error) {
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
	out := &IsInstalledRemoteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISCertObjectClient) IsExportableRemote(ctx context.Context, in *IsExportableRemoteRequest, opts ...dcerpc.CallOption) (*IsExportableRemoteResponse, error) {
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
	out := &IsExportableRemoteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISCertObjectClient) GetCertInfoRemote(ctx context.Context, in *GetCertInfoRemoteRequest, opts ...dcerpc.CallOption) (*GetCertInfoRemoteResponse, error) {
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
	out := &GetCertInfoRemoteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISCertObjectClient) ImportFromBlob(ctx context.Context, in *ImportFromBlobRequest, opts ...dcerpc.CallOption) (*ImportFromBlobResponse, error) {
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
	out := &ImportFromBlobResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISCertObjectClient) ImportFromBlobGetHash(ctx context.Context, in *ImportFromBlobGetHashRequest, opts ...dcerpc.CallOption) (*ImportFromBlobGetHashResponse, error) {
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
	out := &ImportFromBlobGetHashResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISCertObjectClient) ExportToBlob(ctx context.Context, in *ExportToBlobRequest, opts ...dcerpc.CallOption) (*ExportToBlobResponse, error) {
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
	out := &ExportToBlobResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISCertObjectClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIISCertObjectClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultIISCertObjectClient) IPID(ctx context.Context, ipid *dcom.IPID) IISCertObjectClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultIISCertObjectClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewIISCertObjectClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IISCertObjectClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IISCertObjectSyntaxV0_0))...)
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
	return &xxx_DefaultIISCertObjectClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_SetInstanceNameOperation structure represents the InstanceName operation
type xxx_SetInstanceNameOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	NewValue *oaut.String   `idl:"name:newVal" json:"new_value"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInstanceNameOperation) OpNum() int { return 10 }

func (o *xxx_SetInstanceNameOperation) OpName() string { return "/IIISCertObj/v0/InstanceName" }

func (o *xxx_SetInstanceNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInstanceNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// newVal {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NewValue != nil {
			_ptr_newVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewValue != nil {
					if err := o.NewValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewValue, _ptr_newVal); err != nil {
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

func (o *xxx_SetInstanceNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// newVal {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_newVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewValue == nil {
				o.NewValue = &oaut.String{}
			}
			if err := o.NewValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_newVal := func(ptr interface{}) { o.NewValue = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NewValue, _s_newVal, _ptr_newVal); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInstanceNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInstanceNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetInstanceNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetInstanceNameRequest structure represents the InstanceName operation request
type SetInstanceNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// newVal: A string that specifies the web server instance.<33>
	NewValue *oaut.String `idl:"name:newVal" json:"new_value"`
}

func (o *SetInstanceNameRequest) xxx_ToOp(ctx context.Context, op *xxx_SetInstanceNameOperation) *xxx_SetInstanceNameOperation {
	if op == nil {
		op = &xxx_SetInstanceNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.NewValue = o.NewValue
	return op
}

func (o *SetInstanceNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetInstanceNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NewValue = op.NewValue
}
func (o *SetInstanceNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetInstanceNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInstanceNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInstanceNameResponse structure represents the InstanceName operation response
type SetInstanceNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The InstanceName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetInstanceNameResponse) xxx_ToOp(ctx context.Context, op *xxx_SetInstanceNameOperation) *xxx_SetInstanceNameOperation {
	if op == nil {
		op = &xxx_SetInstanceNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetInstanceNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetInstanceNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetInstanceNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetInstanceNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInstanceNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsInstalledRemoteOperation structure represents the IsInstalledRemote operation
type xxx_IsInstalledRemoteOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsInstalledRemoteOperation) OpNum() int { return 12 }

func (o *xxx_IsInstalledRemoteOperation) OpName() string { return "/IIISCertObj/v0/IsInstalledRemote" }

func (o *xxx_IsInstalledRemoteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsInstalledRemoteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsInstalledRemoteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_IsInstalledRemoteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsInstalledRemoteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReturnValue); err != nil {
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

func (o *xxx_IsInstalledRemoteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReturnValue); err != nil {
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

// IsInstalledRemoteRequest structure represents the IsInstalledRemote operation request
type IsInstalledRemoteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *IsInstalledRemoteRequest) xxx_ToOp(ctx context.Context, op *xxx_IsInstalledRemoteOperation) *xxx_IsInstalledRemoteOperation {
	if op == nil {
		op = &xxx_IsInstalledRemoteOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *IsInstalledRemoteRequest) xxx_FromOp(ctx context.Context, op *xxx_IsInstalledRemoteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *IsInstalledRemoteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsInstalledRemoteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsInstalledRemoteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsInstalledRemoteResponse structure represents the IsInstalledRemote operation response
type IsInstalledRemoteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: A pointer to a VARIANT_BOOL.
	ReturnValue int16 `idl:"name:retval" json:"return_value"`
	// Return: The IsInstalledRemote return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsInstalledRemoteResponse) xxx_ToOp(ctx context.Context, op *xxx_IsInstalledRemoteOperation) *xxx_IsInstalledRemoteOperation {
	if op == nil {
		op = &xxx_IsInstalledRemoteOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *IsInstalledRemoteResponse) xxx_FromOp(ctx context.Context, op *xxx_IsInstalledRemoteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *IsInstalledRemoteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsInstalledRemoteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsInstalledRemoteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsExportableRemoteOperation structure represents the IsExportableRemote operation
type xxx_IsExportableRemoteOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsExportableRemoteOperation) OpNum() int { return 14 }

func (o *xxx_IsExportableRemoteOperation) OpName() string {
	return "/IIISCertObj/v0/IsExportableRemote"
}

func (o *xxx_IsExportableRemoteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsExportableRemoteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsExportableRemoteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_IsExportableRemoteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsExportableRemoteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReturnValue); err != nil {
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

func (o *xxx_IsExportableRemoteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReturnValue); err != nil {
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

// IsExportableRemoteRequest structure represents the IsExportableRemote operation request
type IsExportableRemoteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *IsExportableRemoteRequest) xxx_ToOp(ctx context.Context, op *xxx_IsExportableRemoteOperation) *xxx_IsExportableRemoteOperation {
	if op == nil {
		op = &xxx_IsExportableRemoteOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *IsExportableRemoteRequest) xxx_FromOp(ctx context.Context, op *xxx_IsExportableRemoteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *IsExportableRemoteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsExportableRemoteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsExportableRemoteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsExportableRemoteResponse structure represents the IsExportableRemote operation response
type IsExportableRemoteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: A pointer to a VARIANT_BOOL.
	ReturnValue int16 `idl:"name:retval" json:"return_value"`
	// Return: The IsExportableRemote return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsExportableRemoteResponse) xxx_ToOp(ctx context.Context, op *xxx_IsExportableRemoteOperation) *xxx_IsExportableRemoteOperation {
	if op == nil {
		op = &xxx_IsExportableRemoteOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *IsExportableRemoteResponse) xxx_FromOp(ctx context.Context, op *xxx_IsExportableRemoteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *IsExportableRemoteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsExportableRemoteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsExportableRemoteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCertInfoRemoteOperation structure represents the GetCertInfoRemote operation
type xxx_GetCertInfoRemoteOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	BinaryVariant *oaut.Variant  `idl:"name:BinaryVariant" json:"binary_variant"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCertInfoRemoteOperation) OpNum() int { return 16 }

func (o *xxx_GetCertInfoRemoteOperation) OpName() string { return "/IIISCertObj/v0/GetCertInfoRemote" }

func (o *xxx_GetCertInfoRemoteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCertInfoRemoteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCertInfoRemoteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCertInfoRemoteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCertInfoRemoteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// BinaryVariant {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.BinaryVariant != nil {
			_ptr_BinaryVariant := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BinaryVariant != nil {
					if err := o.BinaryVariant.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BinaryVariant, _ptr_BinaryVariant); err != nil {
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

func (o *xxx_GetCertInfoRemoteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// BinaryVariant {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_BinaryVariant := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BinaryVariant == nil {
				o.BinaryVariant = &oaut.Variant{}
			}
			if err := o.BinaryVariant.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_BinaryVariant := func(ptr interface{}) { o.BinaryVariant = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.BinaryVariant, _s_BinaryVariant, _ptr_BinaryVariant); err != nil {
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

// GetCertInfoRemoteRequest structure represents the GetCertInfoRemote operation request
type GetCertInfoRemoteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCertInfoRemoteRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCertInfoRemoteOperation) *xxx_GetCertInfoRemoteOperation {
	if op == nil {
		op = &xxx_GetCertInfoRemoteOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCertInfoRemoteRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCertInfoRemoteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCertInfoRemoteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCertInfoRemoteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCertInfoRemoteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCertInfoRemoteResponse structure represents the GetCertInfoRemote operation response
type GetCertInfoRemoteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// BinaryVariant:  A pointer to a VARIANT that will contain the certificate data. If
	// the method returns S_OK, BinaryVariant contains a single dimension SAFEARRAY of VT_UI1
	// elements as defined in [MS-OAUT]. The data contained in the array is a null-terminated
	// Unicode string containing attribute data from the certificate. The format and contents
	// are described further in the method details.
	BinaryVariant *oaut.Variant `idl:"name:BinaryVariant" json:"binary_variant"`
	// Return: The GetCertInfoRemote return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCertInfoRemoteResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCertInfoRemoteOperation) *xxx_GetCertInfoRemoteOperation {
	if op == nil {
		op = &xxx_GetCertInfoRemoteOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.BinaryVariant = o.BinaryVariant
	op.Return = o.Return
	return op
}

func (o *GetCertInfoRemoteResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCertInfoRemoteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.BinaryVariant = op.BinaryVariant
	o.Return = op.Return
}
func (o *GetCertInfoRemoteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCertInfoRemoteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCertInfoRemoteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ImportFromBlobOperation structure represents the ImportFromBlob operation
type xxx_ImportFromBlobOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	InstanceName      *oaut.String   `idl:"name:InstanceName" json:"instance_name"`
	Password          *oaut.String   `idl:"name:Password" json:"password"`
	InstallToMetabase int16          `idl:"name:bInstallToMetabase" json:"install_to_metabase"`
	AllowExport       int16          `idl:"name:bAllowExport" json:"allow_export"`
	OverwriteExisting int16          `idl:"name:bOverWriteExisting" json:"overwrite_existing"`
	Length            uint32         `idl:"name:cbSize" json:"length"`
	BlobBinary        string         `idl:"name:pBlobBinary;size_is:(cbSize);string" json:"blob_binary"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ImportFromBlobOperation) OpNum() int { return 22 }

func (o *xxx_ImportFromBlobOperation) OpName() string { return "/IIISCertObj/v0/ImportFromBlob" }

func (o *xxx_ImportFromBlobOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.BlobBinary != "" && o.Length == 0 {
		o.Length = uint32(len(o.BlobBinary))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportFromBlobOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// InstanceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.InstanceName != nil {
			_ptr_InstanceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InstanceName != nil {
					if err := o.InstanceName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InstanceName, _ptr_InstanceName); err != nil {
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
	// Password {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Password != nil {
			_ptr_Password := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Password != nil {
					if err := o.Password.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_Password); err != nil {
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
	// bInstallToMetabase {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.InstallToMetabase); err != nil {
			return err
		}
	}
	// bAllowExport {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.AllowExport); err != nil {
			return err
		}
	}
	// bOverWriteExisting {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.OverwriteExisting); err != nil {
			return err
		}
	}
	// cbSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	// pBlobBinary {in} (1:{string, pointer=ref}*(1))(2:{alias=CHAR}[dim:0,size_is=cbSize,string,null](char))
	{
		dimSize1 := uint64(o.Length)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.CharNLen(o.BlobBinary)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		_BlobBinary_buf := []byte(o.BlobBinary)
		if uint64(len(_BlobBinary_buf)) > sizeInfo[0]-1 {
			_BlobBinary_buf = _BlobBinary_buf[:sizeInfo[0]-1]
		}
		if o.BlobBinary != ndr.ZeroString {
			_BlobBinary_buf = append(_BlobBinary_buf, byte(0))
		}
		for i1 := range _BlobBinary_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_BlobBinary_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_BlobBinary_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ImportFromBlobOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// InstanceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_InstanceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InstanceName == nil {
				o.InstanceName = &oaut.String{}
			}
			if err := o.InstanceName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_InstanceName := func(ptr interface{}) { o.InstanceName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.InstanceName, _s_InstanceName, _ptr_InstanceName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Password {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_Password := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Password == nil {
				o.Password = &oaut.String{}
			}
			if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Password := func(ptr interface{}) { o.Password = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Password, _s_Password, _ptr_Password); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bInstallToMetabase {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.InstallToMetabase); err != nil {
			return err
		}
	}
	// bAllowExport {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.AllowExport); err != nil {
			return err
		}
	}
	// bOverWriteExisting {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.OverwriteExisting); err != nil {
			return err
		}
	}
	// cbSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	// pBlobBinary {in} (1:{string, pointer=ref}*(1))(2:{alias=CHAR}[dim:0,size_is=cbSize,string,null](char))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _BlobBinary_buf []byte
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _BlobBinary_buf", sizeInfo[0])
		}
		_BlobBinary_buf = make([]byte, sizeInfo[0])
		for i1 := range _BlobBinary_buf {
			i1 := i1
			if err := w.ReadData(&_BlobBinary_buf[i1]); err != nil {
				return err
			}
		}
		o.BlobBinary = strings.TrimRight(string(_BlobBinary_buf), ndr.ZeroString)
	}
	return nil
}

func (o *xxx_ImportFromBlobOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportFromBlobOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ImportFromBlobOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ImportFromBlobRequest structure represents the ImportFromBlob operation request
type ImportFromBlobRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// InstanceName: A string that specifies a web server instance.
	InstanceName *oaut.String `idl:"name:InstanceName" json:"instance_name"`
	// Password:  A password used to decrypt the imported certificate data.
	Password *oaut.String `idl:"name:Password" json:"password"`
	// bInstallToMetabase: If set to 1 or VARIANT_TRUE, indicates that the imported certificate
	// is associated with the web server instance specified by InstanceName.
	InstallToMetabase int16 `idl:"name:bInstallToMetabase" json:"install_to_metabase"`
	// bAllowExport: If set to 1 or VARIANT_TRUE, indicates that the newly imported certificate
	// is made exportable.
	AllowExport int16 `idl:"name:bAllowExport" json:"allow_export"`
	// bOverWriteExisting: If set to 1 or VARIANT_TRUE, indicates that importing a duplicate
	// certificate will not generate an error.
	OverwriteExisting int16 `idl:"name:bOverWriteExisting" json:"overwrite_existing"`
	// cbSize: Contains the number of bytes in the pBlobBinary buffer including the terminating
	// null character.
	Length uint32 `idl:"name:cbSize" json:"length"`
	// pBlobBinary: A buffer containing an exported, base64-encoded certificate to be imported
	// on the target machine. This buffer is a null-terminated array of bytes.
	BlobBinary string `idl:"name:pBlobBinary;size_is:(cbSize);string" json:"blob_binary"`
}

func (o *ImportFromBlobRequest) xxx_ToOp(ctx context.Context, op *xxx_ImportFromBlobOperation) *xxx_ImportFromBlobOperation {
	if op == nil {
		op = &xxx_ImportFromBlobOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.InstanceName = o.InstanceName
	op.Password = o.Password
	op.InstallToMetabase = o.InstallToMetabase
	op.AllowExport = o.AllowExport
	op.OverwriteExisting = o.OverwriteExisting
	op.Length = o.Length
	op.BlobBinary = o.BlobBinary
	return op
}

func (o *ImportFromBlobRequest) xxx_FromOp(ctx context.Context, op *xxx_ImportFromBlobOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.InstanceName = op.InstanceName
	o.Password = op.Password
	o.InstallToMetabase = op.InstallToMetabase
	o.AllowExport = op.AllowExport
	o.OverwriteExisting = op.OverwriteExisting
	o.Length = op.Length
	o.BlobBinary = op.BlobBinary
}
func (o *ImportFromBlobRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ImportFromBlobRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportFromBlobOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ImportFromBlobResponse structure represents the ImportFromBlob operation response
type ImportFromBlobResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ImportFromBlob return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ImportFromBlobResponse) xxx_ToOp(ctx context.Context, op *xxx_ImportFromBlobOperation) *xxx_ImportFromBlobOperation {
	if op == nil {
		op = &xxx_ImportFromBlobOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ImportFromBlobResponse) xxx_FromOp(ctx context.Context, op *xxx_ImportFromBlobOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ImportFromBlobResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ImportFromBlobResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportFromBlobOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ImportFromBlobGetHashOperation structure represents the ImportFromBlobGetHash operation
type xxx_ImportFromBlobGetHashOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	InstanceName      *oaut.String   `idl:"name:InstanceName" json:"instance_name"`
	Password          *oaut.String   `idl:"name:Password" json:"password"`
	InstallToMetabase int16          `idl:"name:bInstallToMetabase" json:"install_to_metabase"`
	AllowExport       int16          `idl:"name:bAllowExport" json:"allow_export"`
	OverwriteExisting int16          `idl:"name:bOverWriteExisting" json:"overwrite_existing"`
	Length            uint32         `idl:"name:cbSize" json:"length"`
	BlobBinary        string         `idl:"name:pBlobBinary;size_is:(cbSize);string" json:"blob_binary"`
	CertHashLength    uint32         `idl:"name:pcbCertHashSize" json:"cert_hash_length"`
	CertHash          []byte         `idl:"name:pCertHash;size_is:(, pcbCertHashSize)" json:"cert_hash"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ImportFromBlobGetHashOperation) OpNum() int { return 23 }

func (o *xxx_ImportFromBlobGetHashOperation) OpName() string {
	return "/IIISCertObj/v0/ImportFromBlobGetHash"
}

func (o *xxx_ImportFromBlobGetHashOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.BlobBinary != "" && o.Length == 0 {
		o.Length = uint32(len(o.BlobBinary))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportFromBlobGetHashOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// InstanceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.InstanceName != nil {
			_ptr_InstanceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InstanceName != nil {
					if err := o.InstanceName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InstanceName, _ptr_InstanceName); err != nil {
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
	// Password {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Password != nil {
			_ptr_Password := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Password != nil {
					if err := o.Password.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_Password); err != nil {
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
	// bInstallToMetabase {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.InstallToMetabase); err != nil {
			return err
		}
	}
	// bAllowExport {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.AllowExport); err != nil {
			return err
		}
	}
	// bOverWriteExisting {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.OverwriteExisting); err != nil {
			return err
		}
	}
	// cbSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	// pBlobBinary {in} (1:{string, pointer=ref}*(1))(2:{alias=CHAR}[dim:0,size_is=cbSize,string,null](char))
	{
		dimSize1 := uint64(o.Length)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.CharNLen(o.BlobBinary)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		_BlobBinary_buf := []byte(o.BlobBinary)
		if uint64(len(_BlobBinary_buf)) > sizeInfo[0]-1 {
			_BlobBinary_buf = _BlobBinary_buf[:sizeInfo[0]-1]
		}
		if o.BlobBinary != ndr.ZeroString {
			_BlobBinary_buf = append(_BlobBinary_buf, byte(0))
		}
		for i1 := range _BlobBinary_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_BlobBinary_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_BlobBinary_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ImportFromBlobGetHashOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// InstanceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_InstanceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InstanceName == nil {
				o.InstanceName = &oaut.String{}
			}
			if err := o.InstanceName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_InstanceName := func(ptr interface{}) { o.InstanceName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.InstanceName, _s_InstanceName, _ptr_InstanceName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Password {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_Password := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Password == nil {
				o.Password = &oaut.String{}
			}
			if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Password := func(ptr interface{}) { o.Password = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Password, _s_Password, _ptr_Password); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bInstallToMetabase {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.InstallToMetabase); err != nil {
			return err
		}
	}
	// bAllowExport {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.AllowExport); err != nil {
			return err
		}
	}
	// bOverWriteExisting {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.OverwriteExisting); err != nil {
			return err
		}
	}
	// cbSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	// pBlobBinary {in} (1:{string, pointer=ref}*(1))(2:{alias=CHAR}[dim:0,size_is=cbSize,string,null](char))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _BlobBinary_buf []byte
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _BlobBinary_buf", sizeInfo[0])
		}
		_BlobBinary_buf = make([]byte, sizeInfo[0])
		for i1 := range _BlobBinary_buf {
			i1 := i1
			if err := w.ReadData(&_BlobBinary_buf[i1]); err != nil {
				return err
			}
		}
		o.BlobBinary = strings.TrimRight(string(_BlobBinary_buf), ndr.ZeroString)
	}
	return nil
}

func (o *xxx_ImportFromBlobGetHashOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.CertHash != nil && o.CertHashLength == 0 {
		o.CertHashLength = uint32(len(o.CertHash))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportFromBlobGetHashOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcbCertHashSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CertHashLength); err != nil {
			return err
		}
	}
	// pCertHash {out} (1:{pointer=ref}*(2)*(1))(2:{alias=CHAR}[dim:0,size_is=pcbCertHashSize](char))
	{
		if o.CertHash != nil || o.CertHashLength > 0 {
			_ptr_pCertHash := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.CertHashLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.CertHash {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.CertHash[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.CertHash); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CertHash, _ptr_pCertHash); err != nil {
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

func (o *xxx_ImportFromBlobGetHashOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcbCertHashSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CertHashLength); err != nil {
			return err
		}
	}
	// pCertHash {out} (1:{pointer=ref}*(2)*(1))(2:{alias=CHAR}[dim:0,size_is=pcbCertHashSize](char))
	{
		_ptr_pCertHash := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.CertHash", sizeInfo[0])
			}
			o.CertHash = make([]byte, sizeInfo[0])
			for i1 := range o.CertHash {
				i1 := i1
				if err := w.ReadData(&o.CertHash[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pCertHash := func(ptr interface{}) { o.CertHash = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.CertHash, _s_pCertHash, _ptr_pCertHash); err != nil {
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

// ImportFromBlobGetHashRequest structure represents the ImportFromBlobGetHash operation request
type ImportFromBlobGetHashRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// InstanceName: A string that specifies a web server instance.
	InstanceName *oaut.String `idl:"name:InstanceName" json:"instance_name"`
	// Password: A password used to decrypt the imported certificate data.
	Password *oaut.String `idl:"name:Password" json:"password"`
	// bInstallToMetabase: If set to VARIANT_TRUE, indicates that the imported certificate
	// is associated with the web server instance specified by InstanceName.
	InstallToMetabase int16 `idl:"name:bInstallToMetabase" json:"install_to_metabase"`
	// bAllowExport: If set to VARIANT_TRUE, indicates that the newly imported certificate
	// is to be made exportable.
	AllowExport int16 `idl:"name:bAllowExport" json:"allow_export"`
	// bOverWriteExisting: If set to VARIANT_TRUE, indicates that importing a duplicate
	// certificate will not generate an error.
	OverwriteExisting int16 `idl:"name:bOverWriteExisting" json:"overwrite_existing"`
	// cbSize: Contains the number of bytes in the pBlobBinary buffer including the terminating
	// null character.
	Length uint32 `idl:"name:cbSize" json:"length"`
	// pBlobBinary: A buffer containing an exported, base64-encoded certificate to be imported
	// on the target machine. This buffer is a null-terminated array of bytes.
	BlobBinary string `idl:"name:pBlobBinary;size_is:(cbSize);string" json:"blob_binary"`
}

func (o *ImportFromBlobGetHashRequest) xxx_ToOp(ctx context.Context, op *xxx_ImportFromBlobGetHashOperation) *xxx_ImportFromBlobGetHashOperation {
	if op == nil {
		op = &xxx_ImportFromBlobGetHashOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.InstanceName = o.InstanceName
	op.Password = o.Password
	op.InstallToMetabase = o.InstallToMetabase
	op.AllowExport = o.AllowExport
	op.OverwriteExisting = o.OverwriteExisting
	op.Length = o.Length
	op.BlobBinary = o.BlobBinary
	return op
}

func (o *ImportFromBlobGetHashRequest) xxx_FromOp(ctx context.Context, op *xxx_ImportFromBlobGetHashOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.InstanceName = op.InstanceName
	o.Password = op.Password
	o.InstallToMetabase = op.InstallToMetabase
	o.AllowExport = op.AllowExport
	o.OverwriteExisting = op.OverwriteExisting
	o.Length = op.Length
	o.BlobBinary = op.BlobBinary
}
func (o *ImportFromBlobGetHashRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ImportFromBlobGetHashRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportFromBlobGetHashOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ImportFromBlobGetHashResponse structure represents the ImportFromBlobGetHash operation response
type ImportFromBlobGetHashResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pcbCertHashSize: If the method succeeds, returns the number of bytes returned in
	// the pCertHash buffer.
	CertHashLength uint32 `idl:"name:pcbCertHashSize" json:"cert_hash_length"`
	// pCertHash: If the method succeeds, returns a pointer to a memory buffer containing
	// the certificate signature hash. The client MUST free the pointer returned in pCertHash
	// using the appropriate memory allocator as specified by the DCOM implementation.<39>
	CertHash []byte `idl:"name:pCertHash;size_is:(, pcbCertHashSize)" json:"cert_hash"`
	// Return: The ImportFromBlobGetHash return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ImportFromBlobGetHashResponse) xxx_ToOp(ctx context.Context, op *xxx_ImportFromBlobGetHashOperation) *xxx_ImportFromBlobGetHashOperation {
	if op == nil {
		op = &xxx_ImportFromBlobGetHashOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.CertHashLength = o.CertHashLength
	op.CertHash = o.CertHash
	op.Return = o.Return
	return op
}

func (o *ImportFromBlobGetHashResponse) xxx_FromOp(ctx context.Context, op *xxx_ImportFromBlobGetHashOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CertHashLength = op.CertHashLength
	o.CertHash = op.CertHash
	o.Return = op.Return
}
func (o *ImportFromBlobGetHashResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ImportFromBlobGetHashResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportFromBlobGetHashOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExportToBlobOperation structure represents the ExportToBlob operation
type xxx_ExportToBlobOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	InstanceName *oaut.String   `idl:"name:InstanceName" json:"instance_name"`
	Password     *oaut.String   `idl:"name:Password" json:"password"`
	PrivateKey   int16          `idl:"name:bPrivateKey" json:"private_key"`
	CertChain    int16          `idl:"name:bCertChain" json:"cert_chain"`
	Length       uint32         `idl:"name:pcbSize" json:"length"`
	BlobBinary   string         `idl:"name:pBlobBinary;string" json:"blob_binary"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ExportToBlobOperation) OpNum() int { return 25 }

func (o *xxx_ExportToBlobOperation) OpName() string { return "/IIISCertObj/v0/ExportToBlob" }

func (o *xxx_ExportToBlobOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportToBlobOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// InstanceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.InstanceName != nil {
			_ptr_InstanceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InstanceName != nil {
					if err := o.InstanceName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InstanceName, _ptr_InstanceName); err != nil {
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
	// Password {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Password != nil {
			_ptr_Password := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Password != nil {
					if err := o.Password.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_Password); err != nil {
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
	// bPrivateKey {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.PrivateKey); err != nil {
			return err
		}
	}
	// bCertChain {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.CertChain); err != nil {
			return err
		}
	}
	// pcbSize {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	// pBlobBinary {in, out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=CHAR}[dim:0,string,null](char))
	{
		if o.BlobBinary != "" {
			_ptr_pBlobBinary := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.BlobBinary); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.BlobBinary, _ptr_pBlobBinary); err != nil {
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

func (o *xxx_ExportToBlobOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// InstanceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_InstanceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InstanceName == nil {
				o.InstanceName = &oaut.String{}
			}
			if err := o.InstanceName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_InstanceName := func(ptr interface{}) { o.InstanceName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.InstanceName, _s_InstanceName, _ptr_InstanceName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Password {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_Password := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Password == nil {
				o.Password = &oaut.String{}
			}
			if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Password := func(ptr interface{}) { o.Password = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Password, _s_Password, _ptr_Password); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bPrivateKey {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.PrivateKey); err != nil {
			return err
		}
	}
	// bCertChain {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.CertChain); err != nil {
			return err
		}
	}
	// pcbSize {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	// pBlobBinary {in, out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=CHAR}[dim:0,string,null](char))
	{
		_ptr_pBlobBinary := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.BlobBinary); err != nil {
				return err
			}
			return nil
		})
		_s_pBlobBinary := func(ptr interface{}) { o.BlobBinary = *ptr.(*string) }
		if err := w.ReadPointer(&o.BlobBinary, _s_pBlobBinary, _ptr_pBlobBinary); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportToBlobOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportToBlobOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcbSize {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	// pBlobBinary {in, out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=CHAR}[dim:0,string,null](char))
	{
		if o.BlobBinary != "" {
			_ptr_pBlobBinary := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.BlobBinary); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.BlobBinary, _ptr_pBlobBinary); err != nil {
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

func (o *xxx_ExportToBlobOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcbSize {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	// pBlobBinary {in, out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=CHAR}[dim:0,string,null](char))
	{
		_ptr_pBlobBinary := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.BlobBinary); err != nil {
				return err
			}
			return nil
		})
		_s_pBlobBinary := func(ptr interface{}) { o.BlobBinary = *ptr.(*string) }
		if err := w.ReadPointer(&o.BlobBinary, _s_pBlobBinary, _ptr_pBlobBinary); err != nil {
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

// ExportToBlobRequest structure represents the ExportToBlob operation request
type ExportToBlobRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// InstanceName: A string that specifies a web server instance.
	InstanceName *oaut.String `idl:"name:InstanceName" json:"instance_name"`
	// Password: A password used to encrypt the exported certificate data.
	Password *oaut.String `idl:"name:Password" json:"password"`
	// bPrivateKey: If set to VARIANT_TRUE, indicates that the private key of the certificate
	// is to be exported.
	PrivateKey int16 `idl:"name:bPrivateKey" json:"private_key"`
	// bCertChain:  If set to VARIANT_TRUE, indicates that the certificate chain of the
	// certificate referenced by InstanceName is to be exported.
	CertChain int16 `idl:"name:bCertChain" json:"cert_chain"`
	// pcbSize: If the method succeeds, returns the number of valid bytes returned in pBlobBinary.
	Length uint32 `idl:"name:pcbSize" json:"length"`
	// pBlobBinary: If the method succeeds, returns a pointer to a memory buffer containing
	// the exported certificate data. The buffer contains a null-terminated, base64-encoded
	// array of bytes. The client MUST free the pointer returned in pBlobBinary using the
	// appropriate memory allocator as specified for the DCOM implementation.<43>
	BlobBinary string `idl:"name:pBlobBinary;string" json:"blob_binary"`
}

func (o *ExportToBlobRequest) xxx_ToOp(ctx context.Context, op *xxx_ExportToBlobOperation) *xxx_ExportToBlobOperation {
	if op == nil {
		op = &xxx_ExportToBlobOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.InstanceName = o.InstanceName
	op.Password = o.Password
	op.PrivateKey = o.PrivateKey
	op.CertChain = o.CertChain
	op.Length = o.Length
	op.BlobBinary = o.BlobBinary
	return op
}

func (o *ExportToBlobRequest) xxx_FromOp(ctx context.Context, op *xxx_ExportToBlobOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.InstanceName = op.InstanceName
	o.Password = op.Password
	o.PrivateKey = op.PrivateKey
	o.CertChain = op.CertChain
	o.Length = op.Length
	o.BlobBinary = op.BlobBinary
}
func (o *ExportToBlobRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ExportToBlobRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportToBlobOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExportToBlobResponse structure represents the ExportToBlob operation response
type ExportToBlobResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pcbSize: If the method succeeds, returns the number of valid bytes returned in pBlobBinary.
	Length uint32 `idl:"name:pcbSize" json:"length"`
	// pBlobBinary: If the method succeeds, returns a pointer to a memory buffer containing
	// the exported certificate data. The buffer contains a null-terminated, base64-encoded
	// array of bytes. The client MUST free the pointer returned in pBlobBinary using the
	// appropriate memory allocator as specified for the DCOM implementation.<43>
	BlobBinary string `idl:"name:pBlobBinary;string" json:"blob_binary"`
	// Return: The ExportToBlob return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExportToBlobResponse) xxx_ToOp(ctx context.Context, op *xxx_ExportToBlobOperation) *xxx_ExportToBlobOperation {
	if op == nil {
		op = &xxx_ExportToBlobOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Length = o.Length
	op.BlobBinary = o.BlobBinary
	op.Return = o.Return
	return op
}

func (o *ExportToBlobResponse) xxx_FromOp(ctx context.Context, op *xxx_ExportToBlobOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Length = op.Length
	o.BlobBinary = op.BlobBinary
	o.Return = op.Return
}
func (o *ExportToBlobResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ExportToBlobResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportToBlobOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
