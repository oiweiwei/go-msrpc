package icomtrackinginfoevents

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comt"
)

var (
	// IComTrackingInfoEvents interface identifier 4e6cdcc9-fb25-4fd5-9cc5-c9f4b6559cec
	COMTrackingInfoEventsIID = &dcom.IID{Data1: 0x4e6cdcc9, Data2: 0xfb25, Data3: 0x4fd5, Data4: []byte{0x9c, 0xc5, 0xc9, 0xf4, 0xb6, 0x55, 0x9c, 0xec}}
	// Syntax UUID
	COMTrackingInfoEventsSyntaxUUID = &uuid.UUID{TimeLow: 0x4e6cdcc9, TimeMid: 0xfb25, TimeHiAndVersion: 0x4fd5, ClockSeqHiAndReserved: 0x9c, ClockSeqLow: 0xc5, Node: [6]uint8{0xc9, 0xf4, 0xb6, 0x55, 0x9c, 0xec}}
	// Syntax ID
	COMTrackingInfoEventsSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: COMTrackingInfoEventsSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IComTrackingInfoEvents interface.
type COMTrackingInfoEventsClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The OnNewTrackingInfo method handles a tracker event from the server.
	//
	// Return Values: The OnNewTrackingInfo method MUST return S_OK (0x00000000) on success
	// and a failure result (as specified in [MS-ERREF] section 2.1) on failure.
	OnNewTrackingInfo(context.Context, *OnNewTrackingInfoRequest, ...dcerpc.CallOption) (*OnNewTrackingInfoResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) COMTrackingInfoEventsClient
}

type xxx_DefaultCOMTrackingInfoEventsClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCOMTrackingInfoEventsClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCOMTrackingInfoEventsClient) OnNewTrackingInfo(ctx context.Context, in *OnNewTrackingInfoRequest, opts ...dcerpc.CallOption) (*OnNewTrackingInfoResponse, error) {
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
	out := &OnNewTrackingInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCOMTrackingInfoEventsClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCOMTrackingInfoEventsClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCOMTrackingInfoEventsClient) IPID(ctx context.Context, ipid *dcom.IPID) COMTrackingInfoEventsClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCOMTrackingInfoEventsClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewCOMTrackingInfoEventsClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (COMTrackingInfoEventsClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(COMTrackingInfoEventsSyntaxV0_0))...)
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
	return &xxx_DefaultCOMTrackingInfoEventsClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_OnNewTrackingInfoOperation structure represents the OnNewTrackingInfo operation
type xxx_OnNewTrackingInfoOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	ToplevelCollection *dcom.Unknown  `idl:"name:pToplevelCollection" json:"toplevel_collection"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_OnNewTrackingInfoOperation) OpNum() int { return 3 }

func (o *xxx_OnNewTrackingInfoOperation) OpName() string {
	return "/IComTrackingInfoEvents/v0/OnNewTrackingInfo"
}

func (o *xxx_OnNewTrackingInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnNewTrackingInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pToplevelCollection {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		if o.ToplevelCollection != nil {
			_ptr_pToplevelCollection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ToplevelCollection != nil {
					if err := o.ToplevelCollection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ToplevelCollection, _ptr_pToplevelCollection); err != nil {
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

func (o *xxx_OnNewTrackingInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pToplevelCollection {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_pToplevelCollection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ToplevelCollection == nil {
				o.ToplevelCollection = &dcom.Unknown{}
			}
			if err := o.ToplevelCollection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pToplevelCollection := func(ptr interface{}) { o.ToplevelCollection = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.ToplevelCollection, _s_pToplevelCollection, _ptr_pToplevelCollection); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnNewTrackingInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnNewTrackingInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OnNewTrackingInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// OnNewTrackingInfoRequest structure represents the OnNewTrackingInfo operation request
type OnNewTrackingInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pToplevelCollection: An interface pointer of a DCOM object. This MUST be a TrackingInfoCollection
	// OBJREF_CUSTOM (section 2.2.5.5). This collection SHOULD be of type TRKCOLL_PROCESSES
	// (as specified in section 2.2.5.5), and each TrackingInfoObject in the collection
	// SHOULD represent a process on the server. Each process TrackingInfoObject structure
	// SHOULD have the following properties.
	//
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	|   PROPERTY    |     VT     |                                                                                  |
	//	|     NAME      |   VALUE    |                                     MEANING                                      |
	//	|               |            |                                                                                  |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| ProcessID     | 0x00000013 | The process identifier.                                                          |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| ExeName       | 0x00000008 | Implementation-specific identifier of the type of process.<25>                   |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| Paused        | 0x00000013 | TRUE (0x00000001) if the distinguished container for the process is paused;      |
	//	|               |            | otherwise, FALSE (0x00000000).                                                   |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| Recycling     | 0x00000013 | TRUE (0x00000001) if the distinguished instance container for the process is     |
	//	|               |            | recycled; otherwise, FALSE (0x00000000).                                         |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| IsService     | 0x00000013 | TRUE (0x00000001) if the process is a system service; otherwise, FALSE           |
	//	|               |            | (0x00000000). The definition of system service is implementation-specific.<26>   |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| Applications  | 0x0000000D | A TrackingInfoCollection (section 2.2.5.5) of type TRKCOLL_CONTAINERS that       |
	//	|               |            | represents the instance containers in the process.                               |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//
	// Each instance container TrackingInfoObject SHOULD have the following properties.
	//
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	|    PROPERTY     |     VT     |                                                                                  |
	//	|      NAME       |   VALUE    |                                     MEANING                                      |
	//	|                 |            |                                                                                  |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	| ApplicationID   | 0x00000008 | The CurlyBraceGuidString (section 2.2.1) representation of the conglomeration    |
	//	|                 |            | identifier of the conglomeration that is associated with the instance container. |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	| ApplInstanceID  | 0x00000008 | The CurlyBraceGuidString (section 2.2.1) representation of the container         |
	//	|                 |            | identifier of the instance container.                                            |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	| ApplicationType | 0x00000013 | An implementation-specific<27> integer that identifies the type of instance      |
	//	|                 |            | container.                                                                       |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	| PartitionID     | 0x00000008 | The CurlyBraceGuidString (section 2.2.1) representation of the Partition ID of   |
	//	|                 |            | the conglomeration.                                                              |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	| Name            | 0x00000008 | An implementation-specific<28> Unicode string that provides a human-readable     |
	//	|                 |            | name for the conglomeration that is associated with the instance container.      |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	| Components      | 0x0000000D | A TrackingInfoCollection (section 2.2.5.5) of type TRKCOLL_COMPONENTS that       |
	//	|                 |            | represents the components instantiated in the instance container.                |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//
	// Each component TrackingInfoObject SHOULD have the following properties:
	//
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	|   PROPERTY    |     VT     |                                                                                  |
	//	|     NAME      |   VALUE    |                                     MEANING                                      |
	//	|               |            |                                                                                  |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| CLSID         | 0x00000008 | The CurlyBraceGuidString (section 2.2.1) representation of the CLSID of the      |
	//	|               |            | component.                                                                       |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| Objects       | 0x00000013 | The number of component instances for the component in an instance container.    |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| Activated     | 0x00000013 | The number of active component instances for the component in an instance        |
	//	|               |            | container.                                                                       |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| Pooled        | 0x00000013 | The number of pooled component instances for the component in an instance        |
	//	|               |            | container.                                                                       |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| InCall        | 0x00000013 | The number of component instances for the component in an instance container     |
	//	|               |            | that are currently performing a method call.                                     |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| CallTime      | 0x00000013 | A value that indicates the average amount of time, in milliseconds, that it      |
	//	|               |            | takes to complete method calls to component instances for the component. The     |
	//	|               |            | calculation of this value is implementation-specific.<29>                        |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| Name          | 0x00000008 | An implementation-specific Unicode string that provides a human-readable name    |
	//	|               |            | for the component.<30>                                                           |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	ToplevelCollection *dcom.Unknown `idl:"name:pToplevelCollection" json:"toplevel_collection"`
}

func (o *OnNewTrackingInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_OnNewTrackingInfoOperation) *xxx_OnNewTrackingInfoOperation {
	if op == nil {
		op = &xxx_OnNewTrackingInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ToplevelCollection = o.ToplevelCollection
	return op
}

func (o *OnNewTrackingInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_OnNewTrackingInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ToplevelCollection = op.ToplevelCollection
}
func (o *OnNewTrackingInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OnNewTrackingInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnNewTrackingInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OnNewTrackingInfoResponse structure represents the OnNewTrackingInfo operation response
type OnNewTrackingInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OnNewTrackingInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OnNewTrackingInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_OnNewTrackingInfoOperation) *xxx_OnNewTrackingInfoOperation {
	if op == nil {
		op = &xxx_OnNewTrackingInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *OnNewTrackingInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_OnNewTrackingInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *OnNewTrackingInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OnNewTrackingInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnNewTrackingInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
