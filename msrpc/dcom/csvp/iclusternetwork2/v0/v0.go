package iclusternetwork2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	csvp "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = oaut.GoPackage
	_ = csvp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/csvp"
)

var (
	// IClusterNetwork2 interface identifier 2931c32c-f731-4c56-9feb-3d5f1c5e72bf
	ClusterNetwork2IID = &dcom.IID{Data1: 0x2931c32c, Data2: 0xf731, Data3: 0x4c56, Data4: []byte{0x9f, 0xeb, 0x3d, 0x5f, 0x1c, 0x5e, 0x72, 0xbf}}
	// Syntax UUID
	ClusterNetwork2SyntaxUUID = &uuid.UUID{TimeLow: 0x2931c32c, TimeMid: 0xf731, TimeHiAndVersion: 0x4c56, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0xeb, Node: [6]uint8{0x3d, 0x5f, 0x1c, 0x5e, 0x72, 0xbf}}
	// Syntax ID
	ClusterNetwork2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClusterNetwork2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IClusterNetwork2 interface.
type ClusterNetwork2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The SendRTMessage method determines whether roundtrip communication works between
	// two network addresses.
	//
	// The server SHOULD fail this method if the server Initialization State is FALSE.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 3.
	SendRTMessage(context.Context, *SendRTMessageRequest, ...dcerpc.CallOption) (*SendRTMessageResponse, error)

	// The InitializeNode method prepares the server in an implementation-specific way to
	// execute the other methods in the interface. It also informs the client about what
	// port will be used and version information.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 4.
	InitializeNode(context.Context, *InitializeNodeRequest, ...dcerpc.CallOption) (*InitializeNodeResponse, error)

	// The GetIpConfigSerialized method queries the network adapter configuration and returns
	// select information about the adapters.
	//
	// The server SHOULD support this method even if the server Initialization State is
	// FALSE.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 5.
	GetIPConfigSerialized(context.Context, *GetIPConfigSerializedRequest, ...dcerpc.CallOption) (*GetIPConfigSerializedResponse, error)

	// The CleanupNode method cleans up any state initialized by InitializeNode.
	//
	// The server SHOULD fail this method if the server Initialization State is FALSE.
	//
	// This method has no parameters.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 6.
	CleanupNode(context.Context, *CleanupNodeRequest, ...dcerpc.CallOption) (*CleanupNodeResponse, error)

	// The QueryFirewallConfiguration method determines whether the firewall state of the
	// server is compatible with use in a failover cluster. The firewall settings that constitute
	// failover cluster compatibility are implementation-specific. When the server firewall
	// enforces policies specified in [MS-FASP], the server SHOULD determine the firewall
	// state according to how the group of rules is enabled, as specified later in this
	// section.
	//
	// The server SHOULD support this method even if the server Initialization State is
	// FALSE.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 7.
	QueryFirewallConfiguration(context.Context, *QueryFirewallConfigurationRequest, ...dcerpc.CallOption) (*QueryFirewallConfigurationResponse, error)

	// The ProcessAddRoutes method<30> adds Route elements to a Route Collection and initiates
	// monitoring of these routes for packet loss and status data.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.1 and 2.2.
	//
	//	+-------------------+---------------------------------+
	//	|      RETURN       |                                 |
	//	|    VALUE/CODE     |           DESCRIPTION           |
	//	|                   |                                 |
	//	+-------------------+---------------------------------+
	//	+-------------------+---------------------------------+
	//	| 0x00000000 S_OK   | The call was successful.        |
	//	+-------------------+---------------------------------+
	//	| 0x80004005 E_FAIL | Route Monitoring State is TRUE. |
	//	+-------------------+---------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	//
	// The opnum field value for this method is 8.
	ProcessAddRoutes(context.Context, *ProcessAddRoutesRequest, ...dcerpc.CallOption) (*ProcessAddRoutesResponse, error)

	// The GetAddRoutesStatus method<31> retrieves packet loss information and status for
	// the Route elements in the Route Collection previously added with the ProcessAddRoutes
	// method.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.1 and 2.2.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	//
	// The opnum field value for this method is 9.
	GetAddRoutesStatus(context.Context, *GetAddRoutesStatusRequest, ...dcerpc.CallOption) (*GetAddRoutesStatusResponse, error)

	// Opnum10Reserved operation.
	// Opnum10Reserved

	// The CancelAddRoutesRequest method<32> stops packet loss and status monitoring for
	// Route elements previously added in a ProcessAddRoutes (section 3.6.4.6) invocation
	// and removes these routes from Route Collection.
	//
	// This method has no parameters.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.1 and 2.2.
	//
	//	+-------------------+----------------------------------+
	//	|      RETURN       |                                  |
	//	|    VALUE/CODE     |           DESCRIPTION            |
	//	|                   |                                  |
	//	+-------------------+----------------------------------+
	//	+-------------------+----------------------------------+
	//	| 0x00000000 S_OK   | The call was successful.         |
	//	+-------------------+----------------------------------+
	//	| 0x80004005 E_FAIL | Route Monitoring State is FALSE. |
	//	+-------------------+----------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	//
	// The opnum field value for this method is 11.
	CancelAddRoutesRequest(context.Context, *CancelAddRoutesRequestRequest, ...dcerpc.CallOption) (*CancelAddRoutesRequestResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClusterNetwork2Client
}

type xxx_DefaultClusterNetwork2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClusterNetwork2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultClusterNetwork2Client) SendRTMessage(ctx context.Context, in *SendRTMessageRequest, opts ...dcerpc.CallOption) (*SendRTMessageResponse, error) {
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
	out := &SendRTMessageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterNetwork2Client) InitializeNode(ctx context.Context, in *InitializeNodeRequest, opts ...dcerpc.CallOption) (*InitializeNodeResponse, error) {
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
	out := &InitializeNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterNetwork2Client) GetIPConfigSerialized(ctx context.Context, in *GetIPConfigSerializedRequest, opts ...dcerpc.CallOption) (*GetIPConfigSerializedResponse, error) {
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
	out := &GetIPConfigSerializedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterNetwork2Client) CleanupNode(ctx context.Context, in *CleanupNodeRequest, opts ...dcerpc.CallOption) (*CleanupNodeResponse, error) {
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
	out := &CleanupNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterNetwork2Client) QueryFirewallConfiguration(ctx context.Context, in *QueryFirewallConfigurationRequest, opts ...dcerpc.CallOption) (*QueryFirewallConfigurationResponse, error) {
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
	out := &QueryFirewallConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterNetwork2Client) ProcessAddRoutes(ctx context.Context, in *ProcessAddRoutesRequest, opts ...dcerpc.CallOption) (*ProcessAddRoutesResponse, error) {
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
	out := &ProcessAddRoutesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterNetwork2Client) GetAddRoutesStatus(ctx context.Context, in *GetAddRoutesStatusRequest, opts ...dcerpc.CallOption) (*GetAddRoutesStatusResponse, error) {
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
	out := &GetAddRoutesStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterNetwork2Client) CancelAddRoutesRequest(ctx context.Context, in *CancelAddRoutesRequestRequest, opts ...dcerpc.CallOption) (*CancelAddRoutesRequestResponse, error) {
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
	out := &CancelAddRoutesRequestResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterNetwork2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClusterNetwork2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClusterNetwork2Client) IPID(ctx context.Context, ipid *dcom.IPID) ClusterNetwork2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClusterNetwork2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewClusterNetwork2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClusterNetwork2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClusterNetwork2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultClusterNetwork2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_SendRTMessageOperation structure represents the SendRTMessage operation
type xxx_SendRTMessageOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	SourceIPAddress      *oaut.String   `idl:"name:SourceIPAddress" json:"source_ip_address"`
	DestinationIPAddress *oaut.String   `idl:"name:DestIPAddress" json:"destination_ip_address"`
	DestinationPort      uint16         `idl:"name:DestPort" json:"destination_port"`
	AddressFamily        uint16         `idl:"name:AddressFamily" json:"address_family"`
	MessageSize          uint32         `idl:"name:MessageSize" json:"message_size"`
	Timeout              uint32         `idl:"name:Timeout" json:"timeout"`
	RTElapsedTime        uint32         `idl:"name:RTElapsedTime" json:"rt_elapsed_time"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SendRTMessageOperation) OpNum() int { return 3 }

func (o *xxx_SendRTMessageOperation) OpName() string { return "/IClusterNetwork2/v0/SendRTMessage" }

func (o *xxx_SendRTMessageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendRTMessageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// SourceIPAddress {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SourceIPAddress != nil {
			_ptr_SourceIPAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SourceIPAddress != nil {
					if err := o.SourceIPAddress.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SourceIPAddress, _ptr_SourceIPAddress); err != nil {
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
	// DestIPAddress {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DestinationIPAddress != nil {
			_ptr_DestIPAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DestinationIPAddress != nil {
					if err := o.DestinationIPAddress.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DestinationIPAddress, _ptr_DestIPAddress); err != nil {
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
	// DestPort {in} (1:(uint16))
	{
		if err := w.WriteData(o.DestinationPort); err != nil {
			return err
		}
	}
	// AddressFamily {in} (1:(uint16))
	{
		if err := w.WriteData(o.AddressFamily); err != nil {
			return err
		}
	}
	// MessageSize {in} (1:(uint32))
	{
		if err := w.WriteData(o.MessageSize); err != nil {
			return err
		}
	}
	// Timeout {in} (1:(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendRTMessageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// SourceIPAddress {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_SourceIPAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SourceIPAddress == nil {
				o.SourceIPAddress = &oaut.String{}
			}
			if err := o.SourceIPAddress.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_SourceIPAddress := func(ptr interface{}) { o.SourceIPAddress = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SourceIPAddress, _s_SourceIPAddress, _ptr_SourceIPAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DestIPAddress {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_DestIPAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DestinationIPAddress == nil {
				o.DestinationIPAddress = &oaut.String{}
			}
			if err := o.DestinationIPAddress.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DestIPAddress := func(ptr interface{}) { o.DestinationIPAddress = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DestinationIPAddress, _s_DestIPAddress, _ptr_DestIPAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DestPort {in} (1:(uint16))
	{
		if err := w.ReadData(&o.DestinationPort); err != nil {
			return err
		}
	}
	// AddressFamily {in} (1:(uint16))
	{
		if err := w.ReadData(&o.AddressFamily); err != nil {
			return err
		}
	}
	// MessageSize {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MessageSize); err != nil {
			return err
		}
	}
	// Timeout {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendRTMessageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendRTMessageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// RTElapsedTime {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.RTElapsedTime); err != nil {
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

func (o *xxx_SendRTMessageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// RTElapsedTime {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.RTElapsedTime); err != nil {
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

// SendRTMessageRequest structure represents the SendRTMessage operation request
type SendRTMessageRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// SourceIPAddress: The address from which to send the network request. IPv4 addresses
	// MUST be represented in dotted decimal notation. IPv6 addresses MUST be represented
	// in the form specified by [RFC1924].<25>
	SourceIPAddress *oaut.String `idl:"name:SourceIPAddress" json:"source_ip_address"`
	// DestIPAddress: The address to which to send the network request. The address is in
	// the same representation as SourceIPAddress.
	DestinationIPAddress *oaut.String `idl:"name:DestIPAddress" json:"destination_ip_address"`
	// DestPort: This server MUST ignore this value.
	DestinationPort uint16 `idl:"name:DestPort" json:"destination_port"`
	// AddressFamily: The address type of the SourceIPAddress and DestIPAddress parameters.
	//
	//	+-----------------+-----------------------------------+
	//	|                 |                                   |
	//	|      VALUE      |              MEANING              |
	//	|                 |                                   |
	//	+-----------------+-----------------------------------+
	//	+-----------------+-----------------------------------+
	//	| AF_INET 0x0002  | The addresses are in IPv4 format. |
	//	+-----------------+-----------------------------------+
	//	| AF_INET6 0x0017 | The addresses are in IPv6 format. |
	//	+-----------------+-----------------------------------+
	AddressFamily uint16 `idl:"name:AddressFamily" json:"address_family"`
	// MessageSize: The server MUST ignore this value.
	MessageSize uint32 `idl:"name:MessageSize" json:"message_size"`
	// Timeout: An implementation-specific value<26> indicating the maximum amount of time
	// to wait for a response from the destination address.
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
}

func (o *SendRTMessageRequest) xxx_ToOp(ctx context.Context, op *xxx_SendRTMessageOperation) *xxx_SendRTMessageOperation {
	if op == nil {
		op = &xxx_SendRTMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SourceIPAddress = o.SourceIPAddress
	op.DestinationIPAddress = o.DestinationIPAddress
	op.DestinationPort = o.DestinationPort
	op.AddressFamily = o.AddressFamily
	op.MessageSize = o.MessageSize
	op.Timeout = o.Timeout
	return op
}

func (o *SendRTMessageRequest) xxx_FromOp(ctx context.Context, op *xxx_SendRTMessageOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SourceIPAddress = op.SourceIPAddress
	o.DestinationIPAddress = op.DestinationIPAddress
	o.DestinationPort = op.DestinationPort
	o.AddressFamily = op.AddressFamily
	o.MessageSize = op.MessageSize
	o.Timeout = op.Timeout
}
func (o *SendRTMessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SendRTMessageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendRTMessageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SendRTMessageResponse structure represents the SendRTMessage operation response
type SendRTMessageResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// RTElapsedTime: The elapsed time (in milliseconds) between when the server sends the
	// message from the SourceIPAddress to DestIPAddress and when it receives a reply from
	// the destination address.
	RTElapsedTime uint32 `idl:"name:RTElapsedTime" json:"rt_elapsed_time"`
	// Return: The SendRTMessage return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SendRTMessageResponse) xxx_ToOp(ctx context.Context, op *xxx_SendRTMessageOperation) *xxx_SendRTMessageOperation {
	if op == nil {
		op = &xxx_SendRTMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.RTElapsedTime = o.RTElapsedTime
	op.Return = o.Return
	return op
}

func (o *SendRTMessageResponse) xxx_FromOp(ctx context.Context, op *xxx_SendRTMessageOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RTElapsedTime = op.RTElapsedTime
	o.Return = op.Return
}
func (o *SendRTMessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SendRTMessageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendRTMessageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InitializeNodeOperation structure represents the InitializeNode operation
type xxx_InitializeNodeOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	RequestUDPPort        uint16         `idl:"name:RequestUDPPort" json:"request_udp_port"`
	BoundUDPPort          uint16         `idl:"name:BoundUDPPort" json:"bound_udp_port"`
	NodeMajorVersion      uint32         `idl:"name:NodeMajorVersion" json:"node_major_version"`
	NodeMinorVersion      uint32         `idl:"name:NodeMinorVersion" json:"node_minor_version"`
	ClusterPrepareVersion uint32         `idl:"name:ClusprepVersion" json:"cluster_prepare_version"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_InitializeNodeOperation) OpNum() int { return 4 }

func (o *xxx_InitializeNodeOperation) OpName() string { return "/IClusterNetwork2/v0/InitializeNode" }

func (o *xxx_InitializeNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// RequestUDPPort {in} (1:(uint16))
	{
		if err := w.WriteData(o.RequestUDPPort); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// RequestUDPPort {in} (1:(uint16))
	{
		if err := w.ReadData(&o.RequestUDPPort); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// BoundUDPPort {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.WriteData(o.BoundUDPPort); err != nil {
			return err
		}
	}
	// NodeMajorVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.NodeMajorVersion); err != nil {
			return err
		}
	}
	// NodeMinorVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.NodeMinorVersion); err != nil {
			return err
		}
	}
	// ClusprepVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ClusterPrepareVersion); err != nil {
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

func (o *xxx_InitializeNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// BoundUDPPort {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.ReadData(&o.BoundUDPPort); err != nil {
			return err
		}
	}
	// NodeMajorVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.NodeMajorVersion); err != nil {
			return err
		}
	}
	// NodeMinorVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.NodeMinorVersion); err != nil {
			return err
		}
	}
	// ClusprepVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ClusterPrepareVersion); err != nil {
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

// InitializeNodeRequest structure represents the InitializeNode operation request
type InitializeNodeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// RequestUDPPort: A value that the client provides that affects the value of BoundUDPPort.
	RequestUDPPort uint16 `idl:"name:RequestUDPPort" json:"request_udp_port"`
}

func (o *InitializeNodeRequest) xxx_ToOp(ctx context.Context, op *xxx_InitializeNodeOperation) *xxx_InitializeNodeOperation {
	if op == nil {
		op = &xxx_InitializeNodeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RequestUDPPort = o.RequestUDPPort
	return op
}

func (o *InitializeNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_InitializeNodeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RequestUDPPort = op.RequestUDPPort
}
func (o *InitializeNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InitializeNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InitializeNodeResponse structure represents the InitializeNode operation response
type InitializeNodeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// BoundUDPPort:  This parameter is currently not used by the protocol.
	BoundUDPPort uint16 `idl:"name:BoundUDPPort" json:"bound_udp_port"`
	// NodeMajorVersion: The server’s operating system major version value.
	NodeMajorVersion uint32 `idl:"name:NodeMajorVersion" json:"node_major_version"`
	// NodeMinorVersion: The server’s operating system minor version value.
	NodeMinorVersion uint32 `idl:"name:NodeMinorVersion" json:"node_minor_version"`
	// ClusprepVersion: The cluster prepare version.
	ClusterPrepareVersion uint32 `idl:"name:ClusprepVersion" json:"cluster_prepare_version"`
	// Return: The InitializeNode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *InitializeNodeResponse) xxx_ToOp(ctx context.Context, op *xxx_InitializeNodeOperation) *xxx_InitializeNodeOperation {
	if op == nil {
		op = &xxx_InitializeNodeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.BoundUDPPort = o.BoundUDPPort
	op.NodeMajorVersion = o.NodeMajorVersion
	op.NodeMinorVersion = o.NodeMinorVersion
	op.ClusterPrepareVersion = o.ClusterPrepareVersion
	op.Return = o.Return
	return op
}

func (o *InitializeNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_InitializeNodeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.BoundUDPPort = op.BoundUDPPort
	o.NodeMajorVersion = op.NodeMajorVersion
	o.NodeMinorVersion = op.NodeMinorVersion
	o.ClusterPrepareVersion = op.ClusterPrepareVersion
	o.Return = op.Return
}
func (o *InitializeNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InitializeNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIPConfigSerializedOperation structure represents the GetIpConfigSerialized operation
type xxx_GetIPConfigSerializedOperation struct {
	This               *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ApplyClusterFilter bool            `idl:"name:ApplyClusterFilter" json:"apply_cluster_filter"`
	Data               *oaut.SafeArray `idl:"name:Data" json:"data"`
	OutLength          int32           `idl:"name:pcbOut" json:"out_length"`
	Return             int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIPConfigSerializedOperation) OpNum() int { return 5 }

func (o *xxx_GetIPConfigSerializedOperation) OpName() string {
	return "/IClusterNetwork2/v0/GetIpConfigSerialized"
}

func (o *xxx_GetIPConfigSerializedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIPConfigSerializedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ApplyClusterFilter {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.ApplyClusterFilter); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIPConfigSerializedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ApplyClusterFilter {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.ApplyClusterFilter); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIPConfigSerializedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIPConfigSerializedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Data {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Data != nil {
			_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Data != nil {
					if err := o.Data.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data, _ptr_Data); err != nil {
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
	// pcbOut {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.OutLength); err != nil {
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

func (o *xxx_GetIPConfigSerializedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Data {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Data == nil {
				o.Data = &oaut.SafeArray{}
			}
			if err := o.Data.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Data := func(ptr interface{}) { o.Data = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Data, _s_Data, _ptr_Data); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbOut {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.OutLength); err != nil {
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

// GetIPConfigSerializedRequest structure represents the GetIpConfigSerialized operation request
type GetIPConfigSerializedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ApplyClusterFilter: A flag that indicates which adapters to return. If FALSE, then
	// all adapters MUST be returned. If TRUE, then all nonfiltered adapters MUST be returned.
	// Adapters that MUST be filtered are cluster adapters (as specified in the ClusterAdapter
	// field of the ADAPTER2 <27> structure), loopback adapters, and tunnel adapters.
	//
	//	+----------------+----------------------------------+
	//	|                |                                  |
	//	|     VALUE      |             MEANING              |
	//	|                |                                  |
	//	+----------------+----------------------------------+
	//	+----------------+----------------------------------+
	//	| TRUE -128 — -1 | Return all nonfiltered adapters. |
	//	+----------------+----------------------------------+
	//	| FALSE 0        | Return all adapters.             |
	//	+----------------+----------------------------------+
	//	| TRUE 1 — 128   | Return all nonfiltered adapters. |
	//	+----------------+----------------------------------+
	ApplyClusterFilter bool `idl:"name:ApplyClusterFilter" json:"apply_cluster_filter"`
}

func (o *GetIPConfigSerializedRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIPConfigSerializedOperation) *xxx_GetIPConfigSerializedOperation {
	if op == nil {
		op = &xxx_GetIPConfigSerializedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ApplyClusterFilter = o.ApplyClusterFilter
	return op
}

func (o *GetIPConfigSerializedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIPConfigSerializedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ApplyClusterFilter = op.ApplyClusterFilter
}
func (o *GetIPConfigSerializedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIPConfigSerializedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIPConfigSerializedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIPConfigSerializedResponse structure represents the GetIpConfigSerialized operation response
type GetIPConfigSerializedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Data: A buffer that, on success, SHOULD<28> contain a valid ADAPTERLIST2 structure.
	// The client MUST ignore all Guid items in the ADAPTERLIST2 structure except for those
	// Guid items ranging from the first item through the count of 2 multiplied by the value
	// of NumberOfAdapters.
	Data *oaut.SafeArray `idl:"name:Data" json:"data"`
	// pcbOut: MUST be the size of the Data buffer, in bytes.
	OutLength int32 `idl:"name:pcbOut" json:"out_length"`
	// Return: The GetIpConfigSerialized return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIPConfigSerializedResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIPConfigSerializedOperation) *xxx_GetIPConfigSerializedOperation {
	if op == nil {
		op = &xxx_GetIPConfigSerializedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Data = o.Data
	op.OutLength = o.OutLength
	op.Return = o.Return
	return op
}

func (o *GetIPConfigSerializedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIPConfigSerializedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Data = op.Data
	o.OutLength = op.OutLength
	o.Return = op.Return
}
func (o *GetIPConfigSerializedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIPConfigSerializedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIPConfigSerializedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CleanupNodeOperation structure represents the CleanupNode operation
type xxx_CleanupNodeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CleanupNodeOperation) OpNum() int { return 6 }

func (o *xxx_CleanupNodeOperation) OpName() string { return "/IClusterNetwork2/v0/CleanupNode" }

func (o *xxx_CleanupNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CleanupNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CleanupNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CleanupNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CleanupNodeRequest structure represents the CleanupNode operation request
type CleanupNodeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CleanupNodeRequest) xxx_ToOp(ctx context.Context, op *xxx_CleanupNodeOperation) *xxx_CleanupNodeOperation {
	if op == nil {
		op = &xxx_CleanupNodeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CleanupNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_CleanupNodeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CleanupNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CleanupNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CleanupNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CleanupNodeResponse structure represents the CleanupNode operation response
type CleanupNodeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CleanupNode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CleanupNodeResponse) xxx_ToOp(ctx context.Context, op *xxx_CleanupNodeOperation) *xxx_CleanupNodeOperation {
	if op == nil {
		op = &xxx_CleanupNodeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CleanupNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_CleanupNodeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CleanupNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CleanupNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CleanupNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryFirewallConfigurationOperation structure represents the QueryFirewallConfiguration operation
type xxx_QueryFirewallConfigurationOperation struct {
	This                   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServerRulesEnabled     bool           `idl:"name:serverRulesEnabled" json:"server_rules_enabled"`
	ManagementRulesEnabled bool           `idl:"name:mgmtRulesEnabled" json:"management_rules_enabled"`
	Return                 int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryFirewallConfigurationOperation) OpNum() int { return 7 }

func (o *xxx_QueryFirewallConfigurationOperation) OpName() string {
	return "/IClusterNetwork2/v0/QueryFirewallConfiguration"
}

func (o *xxx_QueryFirewallConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFirewallConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryFirewallConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryFirewallConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFirewallConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// serverRulesEnabled {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.ServerRulesEnabled); err != nil {
			return err
		}
	}
	// mgmtRulesEnabled {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.ManagementRulesEnabled); err != nil {
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

func (o *xxx_QueryFirewallConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// serverRulesEnabled {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.ServerRulesEnabled); err != nil {
			return err
		}
	}
	// mgmtRulesEnabled {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.ManagementRulesEnabled); err != nil {
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

// QueryFirewallConfigurationRequest structure represents the QueryFirewallConfiguration operation request
type QueryFirewallConfigurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryFirewallConfigurationRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryFirewallConfigurationOperation) *xxx_QueryFirewallConfigurationOperation {
	if op == nil {
		op = &xxx_QueryFirewallConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *QueryFirewallConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryFirewallConfigurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryFirewallConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryFirewallConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryFirewallConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryFirewallConfigurationResponse structure represents the QueryFirewallConfiguration operation response
type QueryFirewallConfigurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// serverRulesEnabled: An output parameter that MUST be set on a successful return.
	// The value MUST be TRUE if firewall settings are compatible with server-to-server
	// communication in a failover cluster. When the server firewall enforces policies specified
	// in [MS-FASP], the server SHOULD set this value to TRUE if the group of rules with
	// the localized name "Failover Clusters" is enabled.
	//
	//	+----------------+------------------------------------------------------------------+
	//	|                |                                                                  |
	//	|     VALUE      |                             MEANING                              |
	//	|                |                                                                  |
	//	+----------------+------------------------------------------------------------------+
	//	+----------------+------------------------------------------------------------------+
	//	| TRUE -128 — -1 | Firewall settings allow the traffic specified previously.        |
	//	+----------------+------------------------------------------------------------------+
	//	| FALSE 0        | Firewall settings do not allow the traffic specified previously. |
	//	+----------------+------------------------------------------------------------------+
	//	| TRUE 1 — 128   | Firewall settings allow the traffic specified previously.        |
	//	+----------------+------------------------------------------------------------------+
	ServerRulesEnabled bool `idl:"name:serverRulesEnabled" json:"server_rules_enabled"`
	// mgmtRulesEnabled: An output parameter that MUST be set on a successful return. The
	// value MUST be TRUE if firewall settings are compatible with failover cluster management
	// components. When the server firewall enforces policies specified in [MS-FASP], the
	// server SHOULD set this value to TRUE if the group of rules with the localized name
	// "Failover Cluster Manager"<29> is enabled.
	//
	//	+----------------+------------------------------------------------------------------+
	//	|                |                                                                  |
	//	|     VALUE      |                             MEANING                              |
	//	|                |                                                                  |
	//	+----------------+------------------------------------------------------------------+
	//	+----------------+------------------------------------------------------------------+
	//	| TRUE -128 — -1 | Firewall settings allow the traffic specified previously.        |
	//	+----------------+------------------------------------------------------------------+
	//	| FALSE 0        | Firewall settings do not allow the traffic specified previously. |
	//	+----------------+------------------------------------------------------------------+
	//	| TRUE 1 — 128   | Firewall settings allow the traffic specified previously.        |
	//	+----------------+------------------------------------------------------------------+
	ManagementRulesEnabled bool `idl:"name:mgmtRulesEnabled" json:"management_rules_enabled"`
	// Return: The QueryFirewallConfiguration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryFirewallConfigurationResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryFirewallConfigurationOperation) *xxx_QueryFirewallConfigurationOperation {
	if op == nil {
		op = &xxx_QueryFirewallConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ServerRulesEnabled = o.ServerRulesEnabled
	op.ManagementRulesEnabled = o.ManagementRulesEnabled
	op.Return = o.Return
	return op
}

func (o *QueryFirewallConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryFirewallConfigurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ServerRulesEnabled = op.ServerRulesEnabled
	o.ManagementRulesEnabled = op.ManagementRulesEnabled
	o.Return = op.Return
}
func (o *QueryFirewallConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryFirewallConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryFirewallConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ProcessAddRoutesOperation structure represents the ProcessAddRoutes operation
type xxx_ProcessAddRoutesOperation struct {
	This    *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat         `idl:"name:That" json:"that"`
	Request *csvp.AddRoutesRequest `idl:"name:request" json:"request"`
	Return  int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_ProcessAddRoutesOperation) OpNum() int { return 8 }

func (o *xxx_ProcessAddRoutesOperation) OpName() string {
	return "/IClusterNetwork2/v0/ProcessAddRoutes"
}

func (o *xxx_ProcessAddRoutesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ProcessAddRoutesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// request {in} (1:{pointer=ref}*(1))(2:{alias=ADD_ROUTES_REQUEST}(struct))
	{
		if o.Request != nil {
			if err := o.Request.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.AddRoutesRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ProcessAddRoutesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// request {in} (1:{pointer=ref}*(1))(2:{alias=ADD_ROUTES_REQUEST}(struct))
	{
		if o.Request == nil {
			o.Request = &csvp.AddRoutesRequest{}
		}
		if err := o.Request.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ProcessAddRoutesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ProcessAddRoutesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ProcessAddRoutesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ProcessAddRoutesRequest structure represents the ProcessAddRoutes operation request
type ProcessAddRoutesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// request: Designates Route elements to be added to Route Collection.
	Request *csvp.AddRoutesRequest `idl:"name:request" json:"request"`
}

func (o *ProcessAddRoutesRequest) xxx_ToOp(ctx context.Context, op *xxx_ProcessAddRoutesOperation) *xxx_ProcessAddRoutesOperation {
	if op == nil {
		op = &xxx_ProcessAddRoutesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Request = o.Request
	return op
}

func (o *ProcessAddRoutesRequest) xxx_FromOp(ctx context.Context, op *xxx_ProcessAddRoutesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Request = op.Request
}
func (o *ProcessAddRoutesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ProcessAddRoutesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ProcessAddRoutesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ProcessAddRoutesResponse structure represents the ProcessAddRoutes operation response
type ProcessAddRoutesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ProcessAddRoutes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ProcessAddRoutesResponse) xxx_ToOp(ctx context.Context, op *xxx_ProcessAddRoutesOperation) *xxx_ProcessAddRoutesOperation {
	if op == nil {
		op = &xxx_ProcessAddRoutesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ProcessAddRoutesResponse) xxx_FromOp(ctx context.Context, op *xxx_ProcessAddRoutesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ProcessAddRoutesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ProcessAddRoutesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ProcessAddRoutesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAddRoutesStatusOperation structure represents the GetAddRoutesStatus operation
type xxx_GetAddRoutesStatusOperation struct {
	This   *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Reply  *csvp.AddRoutesReply `idl:"name:reply" json:"reply"`
	Return int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAddRoutesStatusOperation) OpNum() int { return 9 }

func (o *xxx_GetAddRoutesStatusOperation) OpName() string {
	return "/IClusterNetwork2/v0/GetAddRoutesStatus"
}

func (o *xxx_GetAddRoutesStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAddRoutesStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAddRoutesStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAddRoutesStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAddRoutesStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// reply {out} (1:{pointer=ref}*(1))(2:{alias=ADD_ROUTES_REPLY}(struct))
	{
		if o.Reply != nil {
			if err := o.Reply.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.AddRoutesReply{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetAddRoutesStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// reply {out} (1:{pointer=ref}*(1))(2:{alias=ADD_ROUTES_REPLY}(struct))
	{
		if o.Reply == nil {
			o.Reply = &csvp.AddRoutesReply{}
		}
		if err := o.Reply.UnmarshalNDR(ctx, w); err != nil {
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

// GetAddRoutesStatusRequest structure represents the GetAddRoutesStatus operation request
type GetAddRoutesStatusRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAddRoutesStatusRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAddRoutesStatusOperation) *xxx_GetAddRoutesStatusOperation {
	if op == nil {
		op = &xxx_GetAddRoutesStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAddRoutesStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAddRoutesStatusOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAddRoutesStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAddRoutesStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAddRoutesStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAddRoutesStatusResponse structure represents the GetAddRoutesStatus operation response
type GetAddRoutesStatusResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// reply: Contains packet loss information and status for Route elements.
	Reply *csvp.AddRoutesReply `idl:"name:reply" json:"reply"`
	// Return: The GetAddRoutesStatus return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAddRoutesStatusResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAddRoutesStatusOperation) *xxx_GetAddRoutesStatusOperation {
	if op == nil {
		op = &xxx_GetAddRoutesStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Reply = o.Reply
	op.Return = o.Return
	return op
}

func (o *GetAddRoutesStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAddRoutesStatusOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Reply = op.Reply
	o.Return = op.Return
}
func (o *GetAddRoutesStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAddRoutesStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAddRoutesStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CancelAddRoutesRequestOperation structure represents the CancelAddRoutesRequest operation
type xxx_CancelAddRoutesRequestOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CancelAddRoutesRequestOperation) OpNum() int { return 11 }

func (o *xxx_CancelAddRoutesRequestOperation) OpName() string {
	return "/IClusterNetwork2/v0/CancelAddRoutesRequest"
}

func (o *xxx_CancelAddRoutesRequestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelAddRoutesRequestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CancelAddRoutesRequestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CancelAddRoutesRequestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelAddRoutesRequestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CancelAddRoutesRequestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CancelAddRoutesRequestRequest structure represents the CancelAddRoutesRequest operation request
type CancelAddRoutesRequestRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CancelAddRoutesRequestRequest) xxx_ToOp(ctx context.Context, op *xxx_CancelAddRoutesRequestOperation) *xxx_CancelAddRoutesRequestOperation {
	if op == nil {
		op = &xxx_CancelAddRoutesRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CancelAddRoutesRequestRequest) xxx_FromOp(ctx context.Context, op *xxx_CancelAddRoutesRequestOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CancelAddRoutesRequestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CancelAddRoutesRequestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelAddRoutesRequestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CancelAddRoutesRequestResponse structure represents the CancelAddRoutesRequest operation response
type CancelAddRoutesRequestResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CancelAddRoutesRequest return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CancelAddRoutesRequestResponse) xxx_ToOp(ctx context.Context, op *xxx_CancelAddRoutesRequestOperation) *xxx_CancelAddRoutesRequestOperation {
	if op == nil {
		op = &xxx_CancelAddRoutesRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CancelAddRoutesRequestResponse) xxx_FromOp(ctx context.Context, op *xxx_CancelAddRoutesRequestOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CancelAddRoutesRequestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CancelAddRoutesRequestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelAddRoutesRequestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
