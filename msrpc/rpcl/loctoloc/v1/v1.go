package loctoloc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	rpcl "github.com/oiweiwei/go-msrpc/msrpc/rpcl"
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
	_ = rpcl.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "rpcl"
)

var (
	// Syntax UUID
	LocToLocSyntaxUUID = &uuid.UUID{TimeLow: 0xe33c0cc4, TimeMid: 0x482, TimeHiAndVersion: 0x101a, ClockSeqHiAndReserved: 0xbc, ClockSeqLow: 0xc, Node: [6]uint8{0x2, 0x60, 0x8c, 0x6b, 0xa2, 0x18}}
	// Syntax ID
	LocToLocSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: LocToLocSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// LocToLoc interface.
type LocToLocClient interface {

	// The I_nsi_lookup_begin method is invoked by a client locator to enumerate the binding
	// information for a set of RPC servers that satisfy a given set of criteria. The Microsoft
	// Interface Definition Language (MIDL) syntax of the method is specified as follows.
	//
	// Return Values: This method does not return any values. RPC exceptions might be thrown
	// from this method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	LookupBegin(context.Context, *LookupBeginRequest, ...dcerpc.CallOption) (*LookupBeginResponse, error)

	// The I_nsi_lookup_done method is invoked to free any resources associated with the
	// context handle returned by a preceding call to the I_nsi_lookup_begin method. The
	// MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method does not return any values.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	LookupDone(context.Context, *LookupDoneRequest, ...dcerpc.CallOption) (*LookupDoneResponse, error)

	// The I_nsi_lookup_next method is invoked to continue an enumeration of binding vectors
	// that satisfy the criteria specified in a call to the I_nsi_lookup_begin method. The
	// number of bindings in the binding_vector is limited by the parameter binding_max_count
	// specified in the call to the I_nsi_lookup_begin method. The MIDL syntax of this method
	// is specified as follows.
	//
	// Return Values: This method does not return any values.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	LookupNext(context.Context, *LookupNextRequest, ...dcerpc.CallOption) (*LookupNextResponse, error)

	// The I_nsi_entry_object_inq_next method is invoked to continue an enumeration initiated
	// by a previous call to the I_nsi_entry_object_inq_begin method. The MIDL syntax of
	// the method is specified as follows.
	//
	// Return Values: This method does not return any values. RPC exceptions can be thrown
	// from this method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	EntryObjectInquireNext(context.Context, *EntryObjectInquireNextRequest, ...dcerpc.CallOption) (*EntryObjectInquireNextResponse, error)

	// The I_nsi_ping_locator method is invoked by the client to determine if the target
	// computer is available as a master locator. The MIDL syntax of the method is specified
	// as follows.
	//
	// Return Values: This method does not return any values.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	PingLocator(context.Context, *PingLocatorRequest, ...dcerpc.CallOption) (*PingLocatorResponse, error)

	// The I_nsi_entry_object_inq_done method is invoked to free any resources associated
	// with the context handle returned by a preceding call to the I_nsi_entry_object_inq_begin
	// method. The MIDL syntax of the method is specified as follows.
	//
	// Return Values: This method does not return any values. RPC exceptions can be thrown
	// from this method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	EntryObjectInquireDone(context.Context, *EntryObjectInquireDoneRequest, ...dcerpc.CallOption) (*EntryObjectInquireDoneResponse, error)

	// The I_nsi_entry_object_inq_begin method is invoked to enumerate the object UUIDs
	// on a name service entry. The MIDL syntax of the method is specified as follows.
	//
	// Return Values: This method does not return any values. RPC exceptions can be thrown
	// from this method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	EntryObjectInquireBegin(context.Context, *EntryObjectInquireBeginRequest, ...dcerpc.CallOption) (*EntryObjectInquireBeginResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultLocToLocClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultLocToLocClient) LookupBegin(ctx context.Context, in *LookupBeginRequest, opts ...dcerpc.CallOption) (*LookupBeginResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupBeginResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultLocToLocClient) LookupDone(ctx context.Context, in *LookupDoneRequest, opts ...dcerpc.CallOption) (*LookupDoneResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupDoneResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultLocToLocClient) LookupNext(ctx context.Context, in *LookupNextRequest, opts ...dcerpc.CallOption) (*LookupNextResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupNextResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultLocToLocClient) EntryObjectInquireNext(ctx context.Context, in *EntryObjectInquireNextRequest, opts ...dcerpc.CallOption) (*EntryObjectInquireNextResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EntryObjectInquireNextResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultLocToLocClient) PingLocator(ctx context.Context, in *PingLocatorRequest, opts ...dcerpc.CallOption) (*PingLocatorResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PingLocatorResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultLocToLocClient) EntryObjectInquireDone(ctx context.Context, in *EntryObjectInquireDoneRequest, opts ...dcerpc.CallOption) (*EntryObjectInquireDoneResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EntryObjectInquireDoneResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultLocToLocClient) EntryObjectInquireBegin(ctx context.Context, in *EntryObjectInquireBeginRequest, opts ...dcerpc.CallOption) (*EntryObjectInquireBeginResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EntryObjectInquireBeginResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultLocToLocClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultLocToLocClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewLocToLocClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (LocToLocClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(LocToLocSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultLocToLocClient{cc: cc}, nil
}

// xxx_LookupBeginOperation structure represents the I_nsi_lookup_begin operation
type xxx_LookupBeginOperation struct {
	EntryNameSyntax uint32         `idl:"name:entry_name_syntax" json:"entry_name_syntax"`
	EntryName       string         `idl:"name:entry_name" json:"entry_name"`
	InterfaceID     *rpcl.SyntaxID `idl:"name:interfaceid;pointer:unique" json:"interface_id"`
	TransferSyntax  *rpcl.SyntaxID `idl:"name:xfersyntax;pointer:unique" json:"transfer_syntax"`
	ObjectUUID      *dtyp.GUID     `idl:"name:obj_uuid" json:"object_uuid"`
	BindingMaxCount uint32         `idl:"name:binding_max_count" json:"binding_max_count"`
	MaxCacheAge     uint32         `idl:"name:MaxCacheAge" json:"max_cache_age"`
	ImportContext   *rpcl.Context  `idl:"name:import_context" json:"import_context"`
	Status          uint16         `idl:"name:status" json:"status"`
}

func (o *xxx_LookupBeginOperation) OpNum() int { return 0 }

func (o *xxx_LookupBeginOperation) OpName() string { return "/LocToLoc/v1/I_nsi_lookup_begin" }

func (o *xxx_LookupBeginOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupBeginOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// entry_name_syntax {in} (1:(uint32))
	{
		if err := w.WriteData(o.EntryNameSyntax); err != nil {
			return err
		}
	}
	// entry_name {in} (1:{string, pointer=unique, alias=STRING_T}*(1)[dim:0,string,null](wchar))
	{
		if o.EntryName != "" {
			_ptr_entry_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.EntryName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.EntryName, _ptr_entry_name); err != nil {
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
	// interfaceid {in} (1:{pointer=unique}*(1))(2:{alias=RPC_SYNTAX_IDENTIFIER}(struct))
	{
		if o.InterfaceID != nil {
			_ptr_interfaceid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InterfaceID != nil {
					if err := o.InterfaceID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rpcl.SyntaxID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InterfaceID, _ptr_interfaceid); err != nil {
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
	// xfersyntax {in} (1:{pointer=unique}*(1))(2:{alias=RPC_SYNTAX_IDENTIFIER}(struct))
	{
		if o.TransferSyntax != nil {
			_ptr_xfersyntax := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TransferSyntax != nil {
					if err := o.TransferSyntax.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rpcl.SyntaxID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TransferSyntax, _ptr_xfersyntax); err != nil {
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
	// obj_uuid {in} (1:{pointer=unique, alias=NSI_UUID_P_T}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectUUID != nil {
			_ptr_obj_uuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectUUID != nil {
					if err := o.ObjectUUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectUUID, _ptr_obj_uuid); err != nil {
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
	// binding_max_count {in} (1:(uint32))
	{
		if err := w.WriteData(o.BindingMaxCount); err != nil {
			return err
		}
	}
	// MaxCacheAge {in} (1:(uint32))
	{
		if err := w.WriteData(o.MaxCacheAge); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupBeginOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// entry_name_syntax {in} (1:(uint32))
	{
		if err := w.ReadData(&o.EntryNameSyntax); err != nil {
			return err
		}
	}
	// entry_name {in} (1:{string, pointer=unique, alias=STRING_T}*(1)[dim:0,string,null](wchar))
	{
		_ptr_entry_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.EntryName); err != nil {
				return err
			}
			return nil
		})
		_s_entry_name := func(ptr interface{}) { o.EntryName = *ptr.(*string) }
		if err := w.ReadPointer(&o.EntryName, _s_entry_name, _ptr_entry_name); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// interfaceid {in} (1:{pointer=unique}*(1))(2:{alias=RPC_SYNTAX_IDENTIFIER}(struct))
	{
		_ptr_interfaceid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InterfaceID == nil {
				o.InterfaceID = &rpcl.SyntaxID{}
			}
			if err := o.InterfaceID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_interfaceid := func(ptr interface{}) { o.InterfaceID = *ptr.(**rpcl.SyntaxID) }
		if err := w.ReadPointer(&o.InterfaceID, _s_interfaceid, _ptr_interfaceid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// xfersyntax {in} (1:{pointer=unique}*(1))(2:{alias=RPC_SYNTAX_IDENTIFIER}(struct))
	{
		_ptr_xfersyntax := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TransferSyntax == nil {
				o.TransferSyntax = &rpcl.SyntaxID{}
			}
			if err := o.TransferSyntax.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_xfersyntax := func(ptr interface{}) { o.TransferSyntax = *ptr.(**rpcl.SyntaxID) }
		if err := w.ReadPointer(&o.TransferSyntax, _s_xfersyntax, _ptr_xfersyntax); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// obj_uuid {in} (1:{pointer=unique, alias=NSI_UUID_P_T}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_obj_uuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectUUID == nil {
				o.ObjectUUID = &dtyp.GUID{}
			}
			if err := o.ObjectUUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_obj_uuid := func(ptr interface{}) { o.ObjectUUID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.ObjectUUID, _s_obj_uuid, _ptr_obj_uuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// binding_max_count {in} (1:(uint32))
	{
		if err := w.ReadData(&o.BindingMaxCount); err != nil {
			return err
		}
	}
	// MaxCacheAge {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MaxCacheAge); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupBeginOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupBeginOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// import_context {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.ImportContext != nil {
			if err := o.ImportContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rpcl.Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// status {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupBeginOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// import_context {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.ImportContext == nil {
			o.ImportContext = &rpcl.Context{}
		}
		if err := o.ImportContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// LookupBeginRequest structure represents the I_nsi_lookup_begin operation request
type LookupBeginRequest struct {
	// entry_name_syntax: An identifier that represents the syntax used for entry_name.
	// The value MUST be RPC_C_NS_SYNTAX_DCE.<8>
	EntryNameSyntax uint32 `idl:"name:entry_name_syntax" json:"entry_name_syntax"`
	// entry_name: A Unicode [UNICODE] string optionally specifying the entry name of the
	// name service entry, using the syntax identified by the entry_name_syntax parameter,
	// as specified in section 2.2.2. This parameter can optionally be null or an empty
	// string.
	EntryName string `idl:"name:entry_name" json:"entry_name"`
	// interfaceid: An optional interface specification. Specified to request only bindings
	// for server entries that have advertised interfaces compatible with this parameter.
	// The client sets interfaceid to NULL to indicate that this parameter is not specified.
	// Interface compatibility is specified in section 3.4.1.5.1.
	InterfaceID *rpcl.SyntaxID `idl:"name:interfaceid;pointer:unique" json:"interface_id"`
	// xfersyntax: An optional transfer syntax specification. Specified to request only
	// bindings for server entries that have advertised interfaces compatible with this
	// parameter. The client sets xfersyntax to NULL to indicate that this parameter is
	// not specified. Interface compatibility is specified in section 3.4.1.5.1.
	TransferSyntax *rpcl.SyntaxID `idl:"name:xfersyntax;pointer:unique" json:"transfer_syntax"`
	// obj_uuid: An optional pointer to an object UUID specification. Specified to request
	// only bindings for the server entries that export this object UUID. If the parameter
	// is NULL or if it contains a null GUID, the parameter is ignored.
	ObjectUUID *dtyp.GUID `idl:"name:obj_uuid" json:"object_uuid"`
	// binding_max_count: The maximum number of elements allowed in the binding vector returned
	// from the I_nsi_lookup_next method. If 0 is specified, an appropriate implementation-specific
	// default maximum MUST be used.<9>
	BindingMaxCount uint32 `idl:"name:binding_max_count" json:"binding_max_count"`
	// MaxCacheAge: Specifies the maximum number of seconds that any results returned from
	// a cache might have been present in the cache without being refreshed. This information
	// is as specified in [C706] Part 2, Name Service Caching.
	MaxCacheAge uint32 `idl:"name:MaxCacheAge" json:"max_cache_age"`
}

func (o *LookupBeginRequest) xxx_ToOp(ctx context.Context) *xxx_LookupBeginOperation {
	if o == nil {
		return &xxx_LookupBeginOperation{}
	}
	return &xxx_LookupBeginOperation{
		EntryNameSyntax: o.EntryNameSyntax,
		EntryName:       o.EntryName,
		InterfaceID:     o.InterfaceID,
		TransferSyntax:  o.TransferSyntax,
		ObjectUUID:      o.ObjectUUID,
		BindingMaxCount: o.BindingMaxCount,
		MaxCacheAge:     o.MaxCacheAge,
	}
}

func (o *LookupBeginRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupBeginOperation) {
	if o == nil {
		return
	}
	o.EntryNameSyntax = op.EntryNameSyntax
	o.EntryName = op.EntryName
	o.InterfaceID = op.InterfaceID
	o.TransferSyntax = op.TransferSyntax
	o.ObjectUUID = op.ObjectUUID
	o.BindingMaxCount = op.BindingMaxCount
	o.MaxCacheAge = op.MaxCacheAge
}
func (o *LookupBeginRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *LookupBeginRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupBeginOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupBeginResponse structure represents the I_nsi_lookup_begin operation response
type LookupBeginResponse struct {
	// import_context: On successful completion of this method, returns a context handle
	// for enumerating binding vectors by using the I_nsi_lookup_next method. This context
	// handle MUST be closed by using the I_nsi_lookup_done method.
	ImportContext *rpcl.Context `idl:"name:import_context" json:"import_context"`
	// status: A 16-bit value that indicates the results of the method call. In case of
	// success, the value MUST be NSI_S_OK. The value MUST be a nonzero value on failure.
	// All failures MUST be treated identically as failure of the whole enumeration process.
	Status uint16 `idl:"name:status" json:"status"`
}

func (o *LookupBeginResponse) xxx_ToOp(ctx context.Context) *xxx_LookupBeginOperation {
	if o == nil {
		return &xxx_LookupBeginOperation{}
	}
	return &xxx_LookupBeginOperation{
		ImportContext: o.ImportContext,
		Status:        o.Status,
	}
}

func (o *LookupBeginResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupBeginOperation) {
	if o == nil {
		return
	}
	o.ImportContext = op.ImportContext
	o.Status = op.Status
}
func (o *LookupBeginResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *LookupBeginResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupBeginOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupDoneOperation structure represents the I_nsi_lookup_done operation
type xxx_LookupDoneOperation struct {
	ImportContext *rpcl.Context `idl:"name:import_context" json:"import_context"`
	Status        uint16        `idl:"name:status" json:"status"`
}

func (o *xxx_LookupDoneOperation) OpNum() int { return 1 }

func (o *xxx_LookupDoneOperation) OpName() string { return "/LocToLoc/v1/I_nsi_lookup_done" }

func (o *xxx_LookupDoneOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupDoneOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// import_context {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.ImportContext != nil {
			if err := o.ImportContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rpcl.Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_LookupDoneOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// import_context {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.ImportContext == nil {
			o.ImportContext = &rpcl.Context{}
		}
		if err := o.ImportContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupDoneOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupDoneOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// import_context {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.ImportContext != nil {
			if err := o.ImportContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rpcl.Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// status {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupDoneOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// import_context {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.ImportContext == nil {
			o.ImportContext = &rpcl.Context{}
		}
		if err := o.ImportContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// LookupDoneRequest structure represents the I_nsi_lookup_done operation request
type LookupDoneRequest struct {
	// import_context: A context handle returned by the server from a preceding call to
	// the I_nsi_lookup_begin method. On successful completion, this parameter MUST be set
	// to NULL by the server and MUST NOT be modified on failure.
	ImportContext *rpcl.Context `idl:"name:import_context" json:"import_context"`
}

func (o *LookupDoneRequest) xxx_ToOp(ctx context.Context) *xxx_LookupDoneOperation {
	if o == nil {
		return &xxx_LookupDoneOperation{}
	}
	return &xxx_LookupDoneOperation{
		ImportContext: o.ImportContext,
	}
}

func (o *LookupDoneRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupDoneOperation) {
	if o == nil {
		return
	}
	o.ImportContext = op.ImportContext
}
func (o *LookupDoneRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *LookupDoneRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupDoneOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupDoneResponse structure represents the I_nsi_lookup_done operation response
type LookupDoneResponse struct {
	// import_context: A context handle returned by the server from a preceding call to
	// the I_nsi_lookup_begin method. On successful completion, this parameter MUST be set
	// to NULL by the server and MUST NOT be modified on failure.
	ImportContext *rpcl.Context `idl:"name:import_context" json:"import_context"`
	// status: A 16-bit value that indicates the results of the method call. In case of
	// success, the value will contain NSI_S_OK, or a nonzero value on failure. All failures
	// MUST be treated identically as a failure of the freeing process initiated by this
	// method, but no further action is required by the caller.
	Status uint16 `idl:"name:status" json:"status"`
}

func (o *LookupDoneResponse) xxx_ToOp(ctx context.Context) *xxx_LookupDoneOperation {
	if o == nil {
		return &xxx_LookupDoneOperation{}
	}
	return &xxx_LookupDoneOperation{
		ImportContext: o.ImportContext,
		Status:        o.Status,
	}
}

func (o *LookupDoneResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupDoneOperation) {
	if o == nil {
		return
	}
	o.ImportContext = op.ImportContext
	o.Status = op.Status
}
func (o *LookupDoneResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *LookupDoneResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupDoneOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupNextOperation structure represents the I_nsi_lookup_next operation
type xxx_LookupNextOperation struct {
	ImportContext *rpcl.Context       `idl:"name:import_context" json:"import_context"`
	BindingVector *rpcl.BindingVector `idl:"name:binding_vector" json:"binding_vector"`
	Status        uint16              `idl:"name:status" json:"status"`
}

func (o *xxx_LookupNextOperation) OpNum() int { return 2 }

func (o *xxx_LookupNextOperation) OpName() string { return "/LocToLoc/v1/I_nsi_lookup_next" }

func (o *xxx_LookupNextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// import_context {in} (1:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.ImportContext != nil {
			if err := o.ImportContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rpcl.Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_LookupNextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// import_context {in} (1:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.ImportContext == nil {
			o.ImportContext = &rpcl.Context{}
		}
		if err := o.ImportContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// binding_vector {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=NSI_BINDING_VECTOR_P_T}*(1))(3:{alias=NSI_BINDING_VECTOR_T}(struct))
	{
		if o.BindingVector != nil {
			_ptr_binding_vector := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BindingVector != nil {
					if err := o.BindingVector.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rpcl.BindingVector{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BindingVector, _ptr_binding_vector); err != nil {
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
	// status {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// binding_vector {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=NSI_BINDING_VECTOR_P_T}*(1))(3:{alias=NSI_BINDING_VECTOR_T}(struct))
	{
		_ptr_binding_vector := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BindingVector == nil {
				o.BindingVector = &rpcl.BindingVector{}
			}
			if err := o.BindingVector.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_binding_vector := func(ptr interface{}) { o.BindingVector = *ptr.(**rpcl.BindingVector) }
		if err := w.ReadPointer(&o.BindingVector, _s_binding_vector, _ptr_binding_vector); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// LookupNextRequest structure represents the I_nsi_lookup_next operation request
type LookupNextRequest struct {
	// import_context: A context handle returned by a preceding call to the I_nsi_lookup_begin
	// method.
	ImportContext *rpcl.Context `idl:"name:import_context" json:"import_context"`
}

func (o *LookupNextRequest) xxx_ToOp(ctx context.Context) *xxx_LookupNextOperation {
	if o == nil {
		return &xxx_LookupNextOperation{}
	}
	return &xxx_LookupNextOperation{
		ImportContext: o.ImportContext,
	}
}

func (o *LookupNextRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupNextOperation) {
	if o == nil {
		return
	}
	o.ImportContext = op.ImportContext
}
func (o *LookupNextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *LookupNextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupNextResponse structure represents the I_nsi_lookup_next operation response
type LookupNextResponse struct {
	// binding_vector: On successful completion, returns a vector containing bindings that
	// satisfy the criteria defined in the preceding call to the I_nsi_lookup_begin method.
	// The caller MUST not assume that the bindings are ordered. The client is responsible
	// for freeing the memory allocated for the binding_vector. The memory allocated for
	// the binding_vector does not need to be freed before subsequent calls to I_nsi_lookup_next.
	BindingVector *rpcl.BindingVector `idl:"name:binding_vector" json:"binding_vector"`
	// status: A 16-bit value that indicates the result of the method call. Any other values,
	// except those listed as follows, MUST be treated as failures and MUST be treated identically.
	// Failure is typically a serious condition (e.g., host out of memory) and SHOULD abort
	// the current operation and then propagated to the higher-layer caller. In the event
	// of failure, the caller SHOULD invoke I_nsi_lookup_done immediately, although it might
	// fail as well.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| NSI_S_OK 0x00000000               | The call returned successfully and binding vector contains at least one binding. |
	//	|                                   | There can be additional bindings that satisfy the criteria.                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| NSI_S_NO_MORE_BINDINGS 0x00000001 | There are no more bindings that satisfy the criteria and binding vector contains |
	//	|                                   | no bindings.                                                                     |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	Status uint16 `idl:"name:status" json:"status"`
}

func (o *LookupNextResponse) xxx_ToOp(ctx context.Context) *xxx_LookupNextOperation {
	if o == nil {
		return &xxx_LookupNextOperation{}
	}
	return &xxx_LookupNextOperation{
		BindingVector: o.BindingVector,
		Status:        o.Status,
	}
}

func (o *LookupNextResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupNextOperation) {
	if o == nil {
		return
	}
	o.BindingVector = op.BindingVector
	o.Status = op.Status
}
func (o *LookupNextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *LookupNextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EntryObjectInquireNextOperation structure represents the I_nsi_entry_object_inq_next operation
type xxx_EntryObjectInquireNextOperation struct {
	InquireContext *rpcl.Context    `idl:"name:InqContext" json:"inquire_context"`
	Vector         *rpcl.UUIDVector `idl:"name:uuid_vec" json:"vector"`
	Status         uint16           `idl:"name:status" json:"status"`
}

func (o *xxx_EntryObjectInquireNextOperation) OpNum() int { return 3 }

func (o *xxx_EntryObjectInquireNextOperation) OpName() string {
	return "/LocToLoc/v1/I_nsi_entry_object_inq_next"
}

func (o *xxx_EntryObjectInquireNextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireNextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// InqContext {in} (1:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.InquireContext != nil {
			if err := o.InquireContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rpcl.Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireNextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// InqContext {in} (1:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.InquireContext == nil {
			o.InquireContext = &rpcl.Context{}
		}
		if err := o.InquireContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireNextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireNextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// uuid_vec {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=NSI_UUID_VECTOR_P_T}*(1))(3:{alias=NSI_UUID_VECTOR_T}(struct))
	{
		if o.Vector != nil {
			_ptr_uuid_vec := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Vector != nil {
					if err := o.Vector.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rpcl.UUIDVector{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Vector, _ptr_uuid_vec); err != nil {
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
	// status {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireNextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// uuid_vec {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=NSI_UUID_VECTOR_P_T}*(1))(3:{alias=NSI_UUID_VECTOR_T}(struct))
	{
		_ptr_uuid_vec := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Vector == nil {
				o.Vector = &rpcl.UUIDVector{}
			}
			if err := o.Vector.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_uuid_vec := func(ptr interface{}) { o.Vector = *ptr.(**rpcl.UUIDVector) }
		if err := w.ReadPointer(&o.Vector, _s_uuid_vec, _ptr_uuid_vec); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// EntryObjectInquireNextRequest structure represents the I_nsi_entry_object_inq_next operation request
type EntryObjectInquireNextRequest struct {
	// InqContext: A context handle returned by the server from a preceding call to the
	// I_nsi_entry_object_inq_begin method.
	InquireContext *rpcl.Context `idl:"name:InqContext" json:"inquire_context"`
}

func (o *EntryObjectInquireNextRequest) xxx_ToOp(ctx context.Context) *xxx_EntryObjectInquireNextOperation {
	if o == nil {
		return &xxx_EntryObjectInquireNextOperation{}
	}
	return &xxx_EntryObjectInquireNextOperation{
		InquireContext: o.InquireContext,
	}
}

func (o *EntryObjectInquireNextRequest) xxx_FromOp(ctx context.Context, op *xxx_EntryObjectInquireNextOperation) {
	if o == nil {
		return
	}
	o.InquireContext = op.InquireContext
}
func (o *EntryObjectInquireNextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EntryObjectInquireNextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EntryObjectInquireNextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EntryObjectInquireNextResponse structure represents the I_nsi_entry_object_inq_next operation response
type EntryObjectInquireNextResponse struct {
	// uuid_vec: On successful completion, returns a vector of object UUIDs for the name
	// service entry. The caller of this method is responsible for freeing any memory allocated
	// for this parameter.
	Vector *rpcl.UUIDVector `idl:"name:uuid_vec" json:"vector"`
	// status: A 16-bit value that indicates the results of the method call. In case of
	// success, the value will contain NSI_S_OK, or a nonzero value on failure. All failures
	// MUST be treated identically as a failure of the continuation of the enumeration process.
	Status uint16 `idl:"name:status" json:"status"`
}

func (o *EntryObjectInquireNextResponse) xxx_ToOp(ctx context.Context) *xxx_EntryObjectInquireNextOperation {
	if o == nil {
		return &xxx_EntryObjectInquireNextOperation{}
	}
	return &xxx_EntryObjectInquireNextOperation{
		Vector: o.Vector,
		Status: o.Status,
	}
}

func (o *EntryObjectInquireNextResponse) xxx_FromOp(ctx context.Context, op *xxx_EntryObjectInquireNextOperation) {
	if o == nil {
		return
	}
	o.Vector = op.Vector
	o.Status = op.Status
}
func (o *EntryObjectInquireNextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EntryObjectInquireNextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EntryObjectInquireNextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PingLocatorOperation structure represents the I_nsi_ping_locator operation
type xxx_PingLocatorOperation struct {
	Status uint32 `idl:"name:status" json:"status"`
}

func (o *xxx_PingLocatorOperation) OpNum() int { return 4 }

func (o *xxx_PingLocatorOperation) OpName() string { return "/LocToLoc/v1/I_nsi_ping_locator" }

func (o *xxx_PingLocatorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PingLocatorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_PingLocatorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_PingLocatorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PingLocatorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PingLocatorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// PingLocatorRequest structure represents the I_nsi_ping_locator operation request
type PingLocatorRequest struct {
}

func (o *PingLocatorRequest) xxx_ToOp(ctx context.Context) *xxx_PingLocatorOperation {
	if o == nil {
		return &xxx_PingLocatorOperation{}
	}
	return &xxx_PingLocatorOperation{}
}

func (o *PingLocatorRequest) xxx_FromOp(ctx context.Context, op *xxx_PingLocatorOperation) {
	if o == nil {
		return
	}
}
func (o *PingLocatorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PingLocatorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PingLocatorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PingLocatorResponse structure represents the I_nsi_ping_locator operation response
type PingLocatorResponse struct {
	// status: A 32-bit value that indicates the results of the method call. In case of
	// success, the value will contain NSI_S_OK, or a nonzero value on failure. All failures
	// MUST be treated identically as a failure of the pinging process initiated by this
	// method, and the target computer SHOULD be treated as unavailable as a master locator.<10>
	Status uint32 `idl:"name:status" json:"status"`
}

func (o *PingLocatorResponse) xxx_ToOp(ctx context.Context) *xxx_PingLocatorOperation {
	if o == nil {
		return &xxx_PingLocatorOperation{}
	}
	return &xxx_PingLocatorOperation{
		Status: o.Status,
	}
}

func (o *PingLocatorResponse) xxx_FromOp(ctx context.Context, op *xxx_PingLocatorOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
}
func (o *PingLocatorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PingLocatorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PingLocatorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EntryObjectInquireDoneOperation structure represents the I_nsi_entry_object_inq_done operation
type xxx_EntryObjectInquireDoneOperation struct {
	InquireContext *rpcl.Context `idl:"name:InqContext" json:"inquire_context"`
	Status         uint16        `idl:"name:status" json:"status"`
}

func (o *xxx_EntryObjectInquireDoneOperation) OpNum() int { return 5 }

func (o *xxx_EntryObjectInquireDoneOperation) OpName() string {
	return "/LocToLoc/v1/I_nsi_entry_object_inq_done"
}

func (o *xxx_EntryObjectInquireDoneOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireDoneOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// InqContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.InquireContext != nil {
			if err := o.InquireContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rpcl.Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireDoneOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// InqContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.InquireContext == nil {
			o.InquireContext = &rpcl.Context{}
		}
		if err := o.InquireContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireDoneOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireDoneOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// InqContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.InquireContext != nil {
			if err := o.InquireContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rpcl.Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// status {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireDoneOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// InqContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.InquireContext == nil {
			o.InquireContext = &rpcl.Context{}
		}
		if err := o.InquireContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// EntryObjectInquireDoneRequest structure represents the I_nsi_entry_object_inq_done operation request
type EntryObjectInquireDoneRequest struct {
	// InqContext: A context handle returned by the server from a preceding I_nsi_entry_object_inq_begin
	// call. On successful completion, this parameter MUST be set to NULL by the server
	// and MUST NOT be modified on failure.
	InquireContext *rpcl.Context `idl:"name:InqContext" json:"inquire_context"`
}

func (o *EntryObjectInquireDoneRequest) xxx_ToOp(ctx context.Context) *xxx_EntryObjectInquireDoneOperation {
	if o == nil {
		return &xxx_EntryObjectInquireDoneOperation{}
	}
	return &xxx_EntryObjectInquireDoneOperation{
		InquireContext: o.InquireContext,
	}
}

func (o *EntryObjectInquireDoneRequest) xxx_FromOp(ctx context.Context, op *xxx_EntryObjectInquireDoneOperation) {
	if o == nil {
		return
	}
	o.InquireContext = op.InquireContext
}
func (o *EntryObjectInquireDoneRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EntryObjectInquireDoneRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EntryObjectInquireDoneOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EntryObjectInquireDoneResponse structure represents the I_nsi_entry_object_inq_done operation response
type EntryObjectInquireDoneResponse struct {
	// InqContext: A context handle returned by the server from a preceding I_nsi_entry_object_inq_begin
	// call. On successful completion, this parameter MUST be set to NULL by the server
	// and MUST NOT be modified on failure.
	InquireContext *rpcl.Context `idl:"name:InqContext" json:"inquire_context"`
	// status: A 16-bit value that indicates the results of the method call. In case of
	// success the value will contain NSI_S_OK, or a nonzero value on failure. All failures
	// MUST be treated identically as a failure of the freeing of resources initiated by
	// this method, but no further action is required by the caller.
	Status uint16 `idl:"name:status" json:"status"`
}

func (o *EntryObjectInquireDoneResponse) xxx_ToOp(ctx context.Context) *xxx_EntryObjectInquireDoneOperation {
	if o == nil {
		return &xxx_EntryObjectInquireDoneOperation{}
	}
	return &xxx_EntryObjectInquireDoneOperation{
		InquireContext: o.InquireContext,
		Status:         o.Status,
	}
}

func (o *EntryObjectInquireDoneResponse) xxx_FromOp(ctx context.Context, op *xxx_EntryObjectInquireDoneOperation) {
	if o == nil {
		return
	}
	o.InquireContext = op.InquireContext
	o.Status = op.Status
}
func (o *EntryObjectInquireDoneResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EntryObjectInquireDoneResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EntryObjectInquireDoneOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EntryObjectInquireBeginOperation structure represents the I_nsi_entry_object_inq_begin operation
type xxx_EntryObjectInquireBeginOperation struct {
	EntryNameSyntax uint32        `idl:"name:EntryNameSyntax" json:"entry_name_syntax"`
	EntryName       string        `idl:"name:EntryName" json:"entry_name"`
	InquireContext  *rpcl.Context `idl:"name:InqContext" json:"inquire_context"`
	Status          uint16        `idl:"name:status" json:"status"`
}

func (o *xxx_EntryObjectInquireBeginOperation) OpNum() int { return 6 }

func (o *xxx_EntryObjectInquireBeginOperation) OpName() string {
	return "/LocToLoc/v1/I_nsi_entry_object_inq_begin"
}

func (o *xxx_EntryObjectInquireBeginOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireBeginOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// EntryNameSyntax {in} (1:(uint32))
	{
		if err := w.WriteData(o.EntryNameSyntax); err != nil {
			return err
		}
	}
	// EntryName {in} (1:{string, pointer=unique, alias=STRING_T}*(1)[dim:0,string,null](wchar))
	{
		if o.EntryName != "" {
			_ptr_EntryName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.EntryName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.EntryName, _ptr_EntryName); err != nil {
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

func (o *xxx_EntryObjectInquireBeginOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// EntryNameSyntax {in} (1:(uint32))
	{
		if err := w.ReadData(&o.EntryNameSyntax); err != nil {
			return err
		}
	}
	// EntryName {in} (1:{string, pointer=unique, alias=STRING_T}*(1)[dim:0,string,null](wchar))
	{
		_ptr_EntryName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.EntryName); err != nil {
				return err
			}
			return nil
		})
		_s_EntryName := func(ptr interface{}) { o.EntryName = *ptr.(*string) }
		if err := w.ReadPointer(&o.EntryName, _s_EntryName, _ptr_EntryName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireBeginOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireBeginOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// InqContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.InquireContext != nil {
			if err := o.InquireContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rpcl.Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// status {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EntryObjectInquireBeginOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// InqContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSI_NS_HANDLE_T, names=ndr_context_handle}(struct))
	{
		if o.InquireContext == nil {
			o.InquireContext = &rpcl.Context{}
		}
		if err := o.InquireContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// status {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// EntryObjectInquireBeginRequest structure represents the I_nsi_entry_object_inq_begin operation request
type EntryObjectInquireBeginRequest struct {
	// EntryNameSyntax: An identifier that represents the syntax used for the entry_name
	// parameter. The value MUST be RPC_C_NS_SYNTAX_DCE.
	EntryNameSyntax uint32 `idl:"name:EntryNameSyntax" json:"entry_name_syntax"`
	// EntryName: A Unicode [UNICODE] string specifying the entry name of the name service
	// entry, using the syntax identified by the entry_name_syntax parameter, as specified
	// in section 2.2.2.
	EntryName string `idl:"name:EntryName" json:"entry_name"`
}

func (o *EntryObjectInquireBeginRequest) xxx_ToOp(ctx context.Context) *xxx_EntryObjectInquireBeginOperation {
	if o == nil {
		return &xxx_EntryObjectInquireBeginOperation{}
	}
	return &xxx_EntryObjectInquireBeginOperation{
		EntryNameSyntax: o.EntryNameSyntax,
		EntryName:       o.EntryName,
	}
}

func (o *EntryObjectInquireBeginRequest) xxx_FromOp(ctx context.Context, op *xxx_EntryObjectInquireBeginOperation) {
	if o == nil {
		return
	}
	o.EntryNameSyntax = op.EntryNameSyntax
	o.EntryName = op.EntryName
}
func (o *EntryObjectInquireBeginRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EntryObjectInquireBeginRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EntryObjectInquireBeginOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EntryObjectInquireBeginResponse structure represents the I_nsi_entry_object_inq_begin operation response
type EntryObjectInquireBeginResponse struct {
	// InqContext: On successful completion, returns a context handle for enumerating object
	// UUID vectors by using the I_nsi_entry_object_inq_next method. This context handle
	// MUST be closed by using the I_nsi_entry_object_inq_done method.
	InquireContext *rpcl.Context `idl:"name:InqContext" json:"inquire_context"`
	// status: A 16-bit value that indicates the results of the method call. In case of
	// success, the value will contain NSI_S_OK, or a nonzero value on failure. All failures
	// MUST be treated identically as a failure of the whole enumeration process.
	Status uint16 `idl:"name:status" json:"status"`
}

func (o *EntryObjectInquireBeginResponse) xxx_ToOp(ctx context.Context) *xxx_EntryObjectInquireBeginOperation {
	if o == nil {
		return &xxx_EntryObjectInquireBeginOperation{}
	}
	return &xxx_EntryObjectInquireBeginOperation{
		InquireContext: o.InquireContext,
		Status:         o.Status,
	}
}

func (o *EntryObjectInquireBeginResponse) xxx_FromOp(ctx context.Context, op *xxx_EntryObjectInquireBeginOperation) {
	if o == nil {
		return
	}
	o.InquireContext = op.InquireContext
	o.Status = op.Status
}
func (o *EntryObjectInquireBeginResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EntryObjectInquireBeginResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EntryObjectInquireBeginOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
