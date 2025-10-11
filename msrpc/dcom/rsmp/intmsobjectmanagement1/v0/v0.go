package intmsobjectmanagement1

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
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// INtmsObjectManagement1 interface identifier b057dc50-3059-11d1-8faf-00a024cb6019
	ObjectManagement1IID = &dcom.IID{Data1: 0xb057dc50, Data2: 0x3059, Data3: 0x11d1, Data4: []byte{0x8f, 0xaf, 0x00, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	// Syntax UUID
	ObjectManagement1SyntaxUUID = &uuid.UUID{TimeLow: 0xb057dc50, TimeMid: 0x3059, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xaf, Node: [6]uint8{0x0, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	// Syntax ID
	ObjectManagement1SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ObjectManagement1SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// INtmsObjectManagement1 interface.
type ObjectManagement1Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetNtmsObjectSecurity method retrieves the security descriptor of an object.
	GetNTMSObjectSecurity(context.Context, *GetNTMSObjectSecurityRequest, ...dcerpc.CallOption) (*GetNTMSObjectSecurityResponse, error)

	// The SetNtmsObjectSecurity method changes the security descriptor of an object.
	SetNTMSObjectSecurity(context.Context, *SetNTMSObjectSecurityRequest, ...dcerpc.CallOption) (*SetNTMSObjectSecurityResponse, error)

	// The GetNtmsObjectAttributeA method retrieves private data of an object, with strings
	// encoded using ASCII.
	GetNTMSObjectAttributeA(context.Context, *GetNTMSObjectAttributeARequest, ...dcerpc.CallOption) (*GetNTMSObjectAttributeAResponse, error)

	// The GetNtmsObjectAttributeW method retrieves private data from an object, with strings
	// encoded using Unicode.
	GetNTMSObjectAttributeW(context.Context, *GetNTMSObjectAttributeWRequest, ...dcerpc.CallOption) (*GetNTMSObjectAttributeWResponse, error)

	// The SetNtmsObjectAttributeA method changes the private data of an object, with strings
	// encoded using ASCII.
	SetNTMSObjectAttributeA(context.Context, *SetNTMSObjectAttributeARequest, ...dcerpc.CallOption) (*SetNTMSObjectAttributeAResponse, error)

	// The SetNtmsObjectAttributeW method changes the private data of an object, with strings
	// encoded using Unicode.
	SetNTMSObjectAttributeW(context.Context, *SetNTMSObjectAttributeWRequest, ...dcerpc.CallOption) (*SetNTMSObjectAttributeWResponse, error)

	// The EnumerateNtmsObject method enumerates the objects of the container specified
	// by the lpContainerId parameter.
	EnumerateNTMSObject(context.Context, *EnumerateNTMSObjectRequest, ...dcerpc.CallOption) (*EnumerateNTMSObjectResponse, error)

	// The DisableNtmsObject method disables an object.
	DisableNTMSObject(context.Context, *DisableNTMSObjectRequest, ...dcerpc.CallOption) (*DisableNTMSObjectResponse, error)

	// The EnableNtmsObject method enables an object.
	EnableNTMSObject(context.Context, *EnableNTMSObjectRequest, ...dcerpc.CallOption) (*EnableNTMSObjectResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ObjectManagement1Client
}

type xxx_DefaultObjectManagement1Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultObjectManagement1Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultObjectManagement1Client) GetNTMSObjectSecurity(ctx context.Context, in *GetNTMSObjectSecurityRequest, opts ...dcerpc.CallOption) (*GetNTMSObjectSecurityResponse, error) {
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
	out := &GetNTMSObjectSecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement1Client) SetNTMSObjectSecurity(ctx context.Context, in *SetNTMSObjectSecurityRequest, opts ...dcerpc.CallOption) (*SetNTMSObjectSecurityResponse, error) {
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
	out := &SetNTMSObjectSecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement1Client) GetNTMSObjectAttributeA(ctx context.Context, in *GetNTMSObjectAttributeARequest, opts ...dcerpc.CallOption) (*GetNTMSObjectAttributeAResponse, error) {
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
	out := &GetNTMSObjectAttributeAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement1Client) GetNTMSObjectAttributeW(ctx context.Context, in *GetNTMSObjectAttributeWRequest, opts ...dcerpc.CallOption) (*GetNTMSObjectAttributeWResponse, error) {
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
	out := &GetNTMSObjectAttributeWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement1Client) SetNTMSObjectAttributeA(ctx context.Context, in *SetNTMSObjectAttributeARequest, opts ...dcerpc.CallOption) (*SetNTMSObjectAttributeAResponse, error) {
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
	out := &SetNTMSObjectAttributeAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement1Client) SetNTMSObjectAttributeW(ctx context.Context, in *SetNTMSObjectAttributeWRequest, opts ...dcerpc.CallOption) (*SetNTMSObjectAttributeWResponse, error) {
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
	out := &SetNTMSObjectAttributeWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement1Client) EnumerateNTMSObject(ctx context.Context, in *EnumerateNTMSObjectRequest, opts ...dcerpc.CallOption) (*EnumerateNTMSObjectResponse, error) {
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
	out := &EnumerateNTMSObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement1Client) DisableNTMSObject(ctx context.Context, in *DisableNTMSObjectRequest, opts ...dcerpc.CallOption) (*DisableNTMSObjectResponse, error) {
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
	out := &DisableNTMSObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement1Client) EnableNTMSObject(ctx context.Context, in *EnableNTMSObjectRequest, opts ...dcerpc.CallOption) (*EnableNTMSObjectResponse, error) {
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
	out := &EnableNTMSObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement1Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultObjectManagement1Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultObjectManagement1Client) IPID(ctx context.Context, ipid *dcom.IPID) ObjectManagement1Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultObjectManagement1Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewObjectManagement1Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ObjectManagement1Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ObjectManagement1SyntaxV0_0))...)
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
	return &xxx_DefaultObjectManagement1Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetNTMSObjectSecurityOperation structure represents the GetNtmsObjectSecurity operation
type xxx_GetNTMSObjectSecurityOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID            *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Type                uint32         `idl:"name:dwType" json:"type"`
	SecurityInformation uint32         `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  []byte         `idl:"name:lpSecurityDescriptor;size_is:(nLength)" json:"security_descriptor"`
	Length              uint32         `idl:"name:nLength" json:"length"`
	LengthNeeded        uint32         `idl:"name:lpnLengthNeeded" json:"length_needed"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSObjectSecurityOperation) OpNum() int { return 0 }

func (o *xxx_GetNTMSObjectSecurityOperation) OpName() string {
	return "/INtmsObjectManagement1/v0/GetNtmsObjectSecurity"
}

func (o *xxx_GetNTMSObjectSecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectSecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// nLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectSecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// nLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectSecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectSecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpSecurityDescriptor {out} (1:{alias=PSECURITY_DESCRIPTOR_NTMS}*(1)[dim:0,size_is=nLength](uint8))
	{
		dimSize1 := uint64(o.Length)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.SecurityDescriptor {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.SecurityDescriptor[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.SecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpnLengthNeeded {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LengthNeeded); err != nil {
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

func (o *xxx_GetNTMSObjectSecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpSecurityDescriptor {out} (1:{alias=PSECURITY_DESCRIPTOR_NTMS,pointer=ref}*(1)[dim:0,size_is=nLength](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SecurityDescriptor", sizeInfo[0])
		}
		o.SecurityDescriptor = make([]byte, sizeInfo[0])
		for i1 := range o.SecurityDescriptor {
			i1 := i1
			if err := w.ReadData(&o.SecurityDescriptor[i1]); err != nil {
				return err
			}
		}
	}
	// lpnLengthNeeded {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LengthNeeded); err != nil {
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

// GetNTMSObjectSecurityRequest structure represents the GetNtmsObjectSecurity operation request
type GetNTMSObjectSecurityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpObjectId: A pointer to the identifier of the object for which to retrieve information.
	ObjectID *dtyp.GUID `idl:"name:lpObjectId" json:"object_id"`
	// dwType: A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// SecurityInformation: A SECURITY_INFORMATION structure specifying the security data
	// to retrieve.
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	// nLength: The size, in bytes, of the client buffer for lpSecurityDescriptor.
	Length uint32 `idl:"name:nLength" json:"length"`
}

func (o *GetNTMSObjectSecurityRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectSecurityOperation) *xxx_GetNTMSObjectSecurityOperation {
	if op == nil {
		op = &xxx_GetNTMSObjectSecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.SecurityInformation = o.SecurityInformation
	op.Length = o.Length
	return op
}

func (o *GetNTMSObjectSecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectSecurityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.SecurityInformation = op.SecurityInformation
	o.Length = op.Length
}
func (o *GetNTMSObjectSecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSObjectSecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectSecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSObjectSecurityResponse structure represents the GetNtmsObjectSecurity operation response
type GetNTMSObjectSecurityResponse struct {
	// XXX: nLength is an implicit input depedency for output parameters
	Length uint32 `idl:"name:nLength" json:"length"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpSecurityDescriptor: A pointer to a SECURITY_DESCRIPTOR structure that describes
	// the security of the object.
	SecurityDescriptor []byte `idl:"name:lpSecurityDescriptor;size_is:(nLength)" json:"security_descriptor"`
	// lpnLengthNeeded: A pointer to the required size of lpSecurityDescriptor if the specified
	// client buffer was not large enough.
	//
	//	+----------------------------------------+------------------------------------------------+
	//	|                 RETURN                 |                                                |
	//	|               VALUE/CODE               |                  DESCRIPTION                   |
	//	|                                        |                                                |
	//	+----------------------------------------+------------------------------------------------+
	//	+----------------------------------------+------------------------------------------------+
	//	| 0x00000000 S_OK                        | The call was successful.                       |
	//	+----------------------------------------+------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED         | Access to an object was denied.                |
	//	+----------------------------------------+------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER     | A parameter is not valid.                      |
	//	+----------------------------------------+------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER   | The specified buffer size is not large enough. |
	//	+----------------------------------------+------------------------------------------------+
	//	| 0x80070546 ERROR_NO_SECURITY_ON_OBJECT | The object has no security information.        |
	//	+----------------------------------------+------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND      | The object was not found.                      |
	//	+----------------------------------------+------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE      | The database query or update failed.           |
	//	+----------------------------------------+------------------------------------------------+
	LengthNeeded uint32 `idl:"name:lpnLengthNeeded" json:"length_needed"`
	// Return: The GetNtmsObjectSecurity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSObjectSecurityResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectSecurityOperation) *xxx_GetNTMSObjectSecurityOperation {
	if op == nil {
		op = &xxx_GetNTMSObjectSecurityOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Length == uint32(0) {
		op.Length = o.Length
	}

	op.That = o.That
	op.SecurityDescriptor = o.SecurityDescriptor
	op.LengthNeeded = o.LengthNeeded
	op.Return = o.Return
	return op
}

func (o *GetNTMSObjectSecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectSecurityOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Length = op.Length

	o.That = op.That
	o.SecurityDescriptor = op.SecurityDescriptor
	o.LengthNeeded = op.LengthNeeded
	o.Return = op.Return
}
func (o *GetNTMSObjectSecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSObjectSecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectSecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNTMSObjectSecurityOperation structure represents the SetNtmsObjectSecurity operation
type xxx_SetNTMSObjectSecurityOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID            *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Type                uint32         `idl:"name:dwType" json:"type"`
	SecurityInformation uint32         `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  []byte         `idl:"name:lpSecurityDescriptor;size_is:(nLength)" json:"security_descriptor"`
	Length              uint32         `idl:"name:nLength" json:"length"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNTMSObjectSecurityOperation) OpNum() int { return 1 }

func (o *xxx_SetNTMSObjectSecurityOperation) OpName() string {
	return "/INtmsObjectManagement1/v0/SetNtmsObjectSecurity"
}

func (o *xxx_SetNTMSObjectSecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.Length == 0 {
		o.Length = uint32(len(o.SecurityDescriptor))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectSecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// lpSecurityDescriptor {in} (1:{alias=PSECURITY_DESCRIPTOR_NTMS}*(1)[dim:0,size_is=nLength](uint8))
	{
		dimSize1 := uint64(o.Length)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.SecurityDescriptor {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.SecurityDescriptor[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.SecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// nLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectSecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// lpSecurityDescriptor {in} (1:{alias=PSECURITY_DESCRIPTOR_NTMS,pointer=ref}*(1)[dim:0,size_is=nLength](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SecurityDescriptor", sizeInfo[0])
		}
		o.SecurityDescriptor = make([]byte, sizeInfo[0])
		for i1 := range o.SecurityDescriptor {
			i1 := i1
			if err := w.ReadData(&o.SecurityDescriptor[i1]); err != nil {
				return err
			}
		}
	}
	// nLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectSecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectSecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNTMSObjectSecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNTMSObjectSecurityRequest structure represents the SetNtmsObjectSecurity operation request
type SetNTMSObjectSecurityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpObjectId: A pointer to the identifier of the object for which to change security
	// information.
	ObjectID *dtyp.GUID `idl:"name:lpObjectId" json:"object_id"`
	// dwType: A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// SecurityInformation: A SECURITY_INFORMATION structure specifying the security data
	// to change.
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	// lpSecurityDescriptor: A pointer to a SECURITY_DESCRIPTOR structure that describes
	// the security descriptor to write to the object.
	SecurityDescriptor []byte `idl:"name:lpSecurityDescriptor;size_is:(nLength)" json:"security_descriptor"`
	// nLength: The length, in bytes, of lpSecurityDescriptor.
	//
	//	+------------------------------------+-------------------------------------------------------------------+
	//	|               RETURN               |                                                                   |
	//	|             VALUE/CODE             |                            DESCRIPTION                            |
	//	|                                    |                                                                   |
	//	+------------------------------------+-------------------------------------------------------------------+
	//	+------------------------------------+-------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                          |
	//	+------------------------------------+-------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | Privileges required to modify the security descriptor are denied. |
	//	+------------------------------------+-------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | A parameter is not valid.                                         |
	//	+------------------------------------+-------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | The object was not found.                                         |
	//	+------------------------------------+-------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database query or update failed.                              |
	//	+------------------------------------+-------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL     | The database is full.                                             |
	//	+------------------------------------+-------------------------------------------------------------------+
	Length uint32 `idl:"name:nLength" json:"length"`
}

func (o *SetNTMSObjectSecurityRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSObjectSecurityOperation) *xxx_SetNTMSObjectSecurityOperation {
	if op == nil {
		op = &xxx_SetNTMSObjectSecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.SecurityInformation = o.SecurityInformation
	op.SecurityDescriptor = o.SecurityDescriptor
	op.Length = o.Length
	return op
}

func (o *SetNTMSObjectSecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSObjectSecurityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.SecurityInformation = op.SecurityInformation
	o.SecurityDescriptor = op.SecurityDescriptor
	o.Length = op.Length
}
func (o *SetNTMSObjectSecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNTMSObjectSecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSObjectSecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNTMSObjectSecurityResponse structure represents the SetNtmsObjectSecurity operation response
type SetNTMSObjectSecurityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetNtmsObjectSecurity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNTMSObjectSecurityResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSObjectSecurityOperation) *xxx_SetNTMSObjectSecurityOperation {
	if op == nil {
		op = &xxx_SetNTMSObjectSecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNTMSObjectSecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSObjectSecurityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNTMSObjectSecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNTMSObjectSecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSObjectSecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNTMSObjectAttributeAOperation structure represents the GetNtmsObjectAttributeA operation
type xxx_GetNTMSObjectAttributeAOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID            *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Type                uint32         `idl:"name:dwType" json:"type"`
	AttributeName       uint8          `idl:"name:lpAttributeName" json:"attribute_name"`
	AttributeData       []byte         `idl:"name:lpAttributeData;size_is:(lpdwAttributeBufferSize);length_is:(lpAttributeSize)" json:"attribute_data"`
	AttributeBufferSize uint32         `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`
	AttributeSize       uint32         `idl:"name:lpAttributeSize" json:"attribute_size"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSObjectAttributeAOperation) OpNum() int { return 2 }

func (o *xxx_GetNTMSObjectAttributeAOperation) OpName() string {
	return "/INtmsObjectManagement1/v0/GetNtmsObjectAttributeA"
}

func (o *xxx_GetNTMSObjectAttributeAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{pointer=ref}*(1)(char))
	{
		if err := w.WriteData(o.AttributeName); err != nil {
			return err
		}
	}
	// lpdwAttributeBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{pointer=ref}*(1)(char))
	{
		if err := w.ReadData(&o.AttributeName); err != nil {
			return err
		}
	}
	// lpdwAttributeBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.AttributeData != nil && o.AttributeSize == 0 {
		o.AttributeSize = uint32(len(o.AttributeData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpAttributeData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwAttributeBufferSize,length_is=lpAttributeSize](uint8))
	{
		dimSize1 := uint64(o.AttributeBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.AttributeSize)
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
		for i1 := range o.AttributeData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AttributeData[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AttributeData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeSize); err != nil {
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

func (o *xxx_GetNTMSObjectAttributeAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpAttributeData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwAttributeBufferSize,length_is=lpAttributeSize](uint8))
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AttributeData", sizeInfo[0])
		}
		o.AttributeData = make([]byte, sizeInfo[0])
		for i1 := range o.AttributeData {
			i1 := i1
			if err := w.ReadData(&o.AttributeData[i1]); err != nil {
				return err
			}
		}
	}
	// lpAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeSize); err != nil {
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

// GetNTMSObjectAttributeARequest structure represents the GetNtmsObjectAttributeA operation request
type GetNTMSObjectAttributeARequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpObjectId: A pointer to the identifier of the object for which to retrieve private
	// data.
	ObjectID *dtyp.GUID `idl:"name:lpObjectId" json:"object_id"`
	// dwType: A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// lpAttributeName: A null-terminated sequence of ASCII characters specifying the name
	// of the extended attribute to retrieve. The attribute name MUST be identical to that
	// specified when creating this attribute using SetNtmsObjectAttributeA.
	AttributeName uint8 `idl:"name:lpAttributeName" json:"attribute_name"`
	// lpdwAttributeBufferSize: A pointer to the size, in bytes, of the client buffer for
	// lpAttributeData.
	AttributeBufferSize uint32 `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`
}

func (o *GetNTMSObjectAttributeARequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeAOperation) *xxx_GetNTMSObjectAttributeAOperation {
	if op == nil {
		op = &xxx_GetNTMSObjectAttributeAOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.AttributeName = o.AttributeName
	op.AttributeBufferSize = o.AttributeBufferSize
	return op
}

func (o *GetNTMSObjectAttributeARequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeAOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.AttributeName = op.AttributeName
	o.AttributeBufferSize = op.AttributeBufferSize
}
func (o *GetNTMSObjectAttributeARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSObjectAttributeARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectAttributeAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSObjectAttributeAResponse structure represents the GetNtmsObjectAttributeA operation response
type GetNTMSObjectAttributeAResponse struct {
	// XXX: lpdwAttributeBufferSize is an implicit input depedency for output parameters
	AttributeBufferSize uint32 `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpAttributeData: A buffer containing the attribute.
	AttributeData []byte `idl:"name:lpAttributeData;size_is:(lpdwAttributeBufferSize);length_is:(lpAttributeSize)" json:"attribute_data"`
	// lpAttributeSize: The size of lpAttributeData. If the specified client buffer was
	// not large enough, lpAttributeSize MUST point to the required size of lpAttributeData;
	// otherwise, it MUST point to the number of bytes that are returned by the server in
	// the buffer lpAttributeData.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED       | Access to the object is denied; other security errors are possible but indicate  |
	//	|                                      | a security subsystem error.                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER   | The parameter is not valid.                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The specified buffer size is not large enough.                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800700E8 ERROR_NO_DATA             | The specified attribute is greater than or equal to NTMS_MAXATTR_LENGTH, defined |
	//	|                                      | in the Platform SDK file NTMSApi.h.                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800706C6 ERROR_RPC_S_INVALID_BOUND | The array bounds are invalid.                                                    |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND    | The specified attribute was not found.                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE    | The database query or update failed.                                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	AttributeSize uint32 `idl:"name:lpAttributeSize" json:"attribute_size"`
	// Return: The GetNtmsObjectAttributeA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSObjectAttributeAResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeAOperation) *xxx_GetNTMSObjectAttributeAOperation {
	if op == nil {
		op = &xxx_GetNTMSObjectAttributeAOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.AttributeBufferSize == uint32(0) {
		op.AttributeBufferSize = o.AttributeBufferSize
	}

	op.That = o.That
	op.AttributeData = o.AttributeData
	op.AttributeSize = o.AttributeSize
	op.Return = o.Return
	return op
}

func (o *GetNTMSObjectAttributeAResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeAOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.AttributeBufferSize = op.AttributeBufferSize

	o.That = op.That
	o.AttributeData = op.AttributeData
	o.AttributeSize = op.AttributeSize
	o.Return = op.Return
}
func (o *GetNTMSObjectAttributeAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSObjectAttributeAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectAttributeAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNTMSObjectAttributeWOperation structure represents the GetNtmsObjectAttributeW operation
type xxx_GetNTMSObjectAttributeWOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID            *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Type                uint32         `idl:"name:dwType" json:"type"`
	AttributeName       string         `idl:"name:lpAttributeName;string" json:"attribute_name"`
	AttributeData       []byte         `idl:"name:lpAttributeData;size_is:(lpdwAttributeBufferSize);length_is:(lpAttributeSize)" json:"attribute_data"`
	AttributeBufferSize uint32         `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`
	AttributeSize       uint32         `idl:"name:lpAttributeSize" json:"attribute_size"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSObjectAttributeWOperation) OpNum() int { return 3 }

func (o *xxx_GetNTMSObjectAttributeWOperation) OpName() string {
	return "/INtmsObjectManagement1/v0/GetNtmsObjectAttributeW"
}

func (o *xxx_GetNTMSObjectAttributeWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.AttributeName); err != nil {
			return err
		}
	}
	// lpdwAttributeBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.AttributeName); err != nil {
			return err
		}
	}
	// lpdwAttributeBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.AttributeData != nil && o.AttributeSize == 0 {
		o.AttributeSize = uint32(len(o.AttributeData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpAttributeData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwAttributeBufferSize,length_is=lpAttributeSize](uint8))
	{
		dimSize1 := uint64(o.AttributeBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.AttributeSize)
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
		for i1 := range o.AttributeData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AttributeData[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AttributeData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeSize); err != nil {
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

func (o *xxx_GetNTMSObjectAttributeWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpAttributeData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwAttributeBufferSize,length_is=lpAttributeSize](uint8))
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AttributeData", sizeInfo[0])
		}
		o.AttributeData = make([]byte, sizeInfo[0])
		for i1 := range o.AttributeData {
			i1 := i1
			if err := w.ReadData(&o.AttributeData[i1]); err != nil {
				return err
			}
		}
	}
	// lpAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeSize); err != nil {
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

// GetNTMSObjectAttributeWRequest structure represents the GetNtmsObjectAttributeW operation request
type GetNTMSObjectAttributeWRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpObjectId: A pointer to the identifier of the object for which to retrieve private
	// data.
	ObjectID *dtyp.GUID `idl:"name:lpObjectId" json:"object_id"`
	// dwType: A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// lpAttributeName: A null-terminated sequence of Unicode characters specifying the
	// name of the extended attribute to retrieve. The attribute name MUST be identical
	// to that specified when creating this attribute using SetNtmsObjectAttributeW.
	AttributeName string `idl:"name:lpAttributeName;string" json:"attribute_name"`
	// lpdwAttributeBufferSize: A pointer to the size, in bytes, of the client buffer for
	// lpAttributeData.
	AttributeBufferSize uint32 `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`
}

func (o *GetNTMSObjectAttributeWRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeWOperation) *xxx_GetNTMSObjectAttributeWOperation {
	if op == nil {
		op = &xxx_GetNTMSObjectAttributeWOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.AttributeName = o.AttributeName
	op.AttributeBufferSize = o.AttributeBufferSize
	return op
}

func (o *GetNTMSObjectAttributeWRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeWOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.AttributeName = op.AttributeName
	o.AttributeBufferSize = op.AttributeBufferSize
}
func (o *GetNTMSObjectAttributeWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSObjectAttributeWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectAttributeWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSObjectAttributeWResponse structure represents the GetNtmsObjectAttributeW operation response
type GetNTMSObjectAttributeWResponse struct {
	// XXX: lpdwAttributeBufferSize is an implicit input depedency for output parameters
	AttributeBufferSize uint32 `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpAttributeData: A buffer containing the attribute.
	AttributeData []byte `idl:"name:lpAttributeData;size_is:(lpdwAttributeBufferSize);length_is:(lpAttributeSize)" json:"attribute_data"`
	// lpAttributeSize: The size of lpAttributeData. If the specified client buffer was
	// not large enough, lpAttributeSize MUST point to the required size of lpAttributeData;
	// otherwise, it MUST point to the number of bytes that are returned by the server in
	// the buffer lpAttributeData.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED       | Access to the object is denied; other security errors are possible but indicate  |
	//	|                                      | a security subsystem error.                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER   | The parameter is not valid.                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The specified buffer size is not large enough.                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800700E8 ERROR_NO_DATA             | The specified attribute is greater than or equal to NTMS_MAXATTR_LENGTH, defined |
	//	|                                      | in the Platform SDK file NTMSApi.h.                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND    | The specified attribute was not found.                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE    | The database query or update failed.                                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	AttributeSize uint32 `idl:"name:lpAttributeSize" json:"attribute_size"`
	// Return: The GetNtmsObjectAttributeW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSObjectAttributeWResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeWOperation) *xxx_GetNTMSObjectAttributeWOperation {
	if op == nil {
		op = &xxx_GetNTMSObjectAttributeWOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.AttributeBufferSize == uint32(0) {
		op.AttributeBufferSize = o.AttributeBufferSize
	}

	op.That = o.That
	op.AttributeData = o.AttributeData
	op.AttributeSize = o.AttributeSize
	op.Return = o.Return
	return op
}

func (o *GetNTMSObjectAttributeWResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeWOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.AttributeBufferSize = op.AttributeBufferSize

	o.That = op.That
	o.AttributeData = op.AttributeData
	o.AttributeSize = op.AttributeSize
	o.Return = op.Return
}
func (o *GetNTMSObjectAttributeWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSObjectAttributeWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectAttributeWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNTMSObjectAttributeAOperation structure represents the SetNtmsObjectAttributeA operation
type xxx_SetNTMSObjectAttributeAOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID      *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Type          uint32         `idl:"name:dwType" json:"type"`
	AttributeName uint8          `idl:"name:lpAttributeName" json:"attribute_name"`
	AttributeData []byte         `idl:"name:lpAttributeData;size_is:(AttributeSize)" json:"attribute_data"`
	AttributeSize uint32         `idl:"name:AttributeSize" json:"attribute_size"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNTMSObjectAttributeAOperation) OpNum() int { return 4 }

func (o *xxx_SetNTMSObjectAttributeAOperation) OpName() string {
	return "/INtmsObjectManagement1/v0/SetNtmsObjectAttributeA"
}

func (o *xxx_SetNTMSObjectAttributeAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.AttributeData != nil && o.AttributeSize == 0 {
		o.AttributeSize = uint32(len(o.AttributeData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectAttributeAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{pointer=ref}*(1)(char))
	{
		if err := w.WriteData(o.AttributeName); err != nil {
			return err
		}
	}
	// lpAttributeData {in} (1:{pointer=ref}*(1)[dim:0,size_is=AttributeSize](uint8))
	{
		dimSize1 := uint64(o.AttributeSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.AttributeData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AttributeData[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AttributeData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// AttributeSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectAttributeAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{pointer=ref}*(1)(char))
	{
		if err := w.ReadData(&o.AttributeName); err != nil {
			return err
		}
	}
	// lpAttributeData {in} (1:{pointer=ref}*(1)[dim:0,size_is=AttributeSize](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AttributeData", sizeInfo[0])
		}
		o.AttributeData = make([]byte, sizeInfo[0])
		for i1 := range o.AttributeData {
			i1 := i1
			if err := w.ReadData(&o.AttributeData[i1]); err != nil {
				return err
			}
		}
	}
	// AttributeSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectAttributeAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectAttributeAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNTMSObjectAttributeAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNTMSObjectAttributeARequest structure represents the SetNtmsObjectAttributeA operation request
type SetNTMSObjectAttributeARequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpObjectId: A pointer to the identifier of the object for which to set private data.
	ObjectID *dtyp.GUID `idl:"name:lpObjectId" json:"object_id"`
	// dwType: A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// lpAttributeName: A null-terminated sequence of ASCII characters specifying the name
	// of the extended attribute to set. The client can give any name to the extended attribute
	// and MUST use the same name in the GetNtmsObjectAttributeA method.
	AttributeName uint8 `idl:"name:lpAttributeName" json:"attribute_name"`
	// lpAttributeData: A buffer containing the attribute.
	AttributeData []byte `idl:"name:lpAttributeData;size_is:(AttributeSize)" json:"attribute_data"`
	// AttributeSize: The size of lpAttributeData.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | Access to the object is denied; other security errors are possible but indicate  |
	//	|                                    | a security subsystem error.                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY | An allocation failure occurred during processing.                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | The parameter is not valid.                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007007B ERROR_INVALID_NAME      | The attribute name is invalid or too long. The NTMS_MAXATTR_NAMELEN              |
	//	|                                    | value, defined in the Platform SDK file NTMSApi.h, specifies the maximum         |
	//	|                                    | null-terminated attribute name length.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | The object was not found.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800708CA ERROR_NOT_CONNECTED     | Unable to connect to the RSM service.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	AttributeSize uint32 `idl:"name:AttributeSize" json:"attribute_size"`
}

func (o *SetNTMSObjectAttributeARequest) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSObjectAttributeAOperation) *xxx_SetNTMSObjectAttributeAOperation {
	if op == nil {
		op = &xxx_SetNTMSObjectAttributeAOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.AttributeName = o.AttributeName
	op.AttributeData = o.AttributeData
	op.AttributeSize = o.AttributeSize
	return op
}

func (o *SetNTMSObjectAttributeARequest) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSObjectAttributeAOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.AttributeName = op.AttributeName
	o.AttributeData = op.AttributeData
	o.AttributeSize = op.AttributeSize
}
func (o *SetNTMSObjectAttributeARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNTMSObjectAttributeARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSObjectAttributeAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNTMSObjectAttributeAResponse structure represents the SetNtmsObjectAttributeA operation response
type SetNTMSObjectAttributeAResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetNtmsObjectAttributeA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNTMSObjectAttributeAResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSObjectAttributeAOperation) *xxx_SetNTMSObjectAttributeAOperation {
	if op == nil {
		op = &xxx_SetNTMSObjectAttributeAOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNTMSObjectAttributeAResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSObjectAttributeAOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNTMSObjectAttributeAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNTMSObjectAttributeAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSObjectAttributeAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNTMSObjectAttributeWOperation structure represents the SetNtmsObjectAttributeW operation
type xxx_SetNTMSObjectAttributeWOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID      *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Type          uint32         `idl:"name:dwType" json:"type"`
	AttributeName string         `idl:"name:lpAttributeName;string" json:"attribute_name"`
	AttributeData []byte         `idl:"name:lpAttributeData;size_is:(AttributeSize)" json:"attribute_data"`
	AttributeSize uint32         `idl:"name:AttributeSize" json:"attribute_size"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNTMSObjectAttributeWOperation) OpNum() int { return 5 }

func (o *xxx_SetNTMSObjectAttributeWOperation) OpName() string {
	return "/INtmsObjectManagement1/v0/SetNtmsObjectAttributeW"
}

func (o *xxx_SetNTMSObjectAttributeWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.AttributeData != nil && o.AttributeSize == 0 {
		o.AttributeSize = uint32(len(o.AttributeData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectAttributeWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.AttributeName); err != nil {
			return err
		}
	}
	// lpAttributeData {in} (1:{pointer=ref}*(1)[dim:0,size_is=AttributeSize](uint8))
	{
		dimSize1 := uint64(o.AttributeSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.AttributeData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AttributeData[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AttributeData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// AttributeSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectAttributeWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.AttributeName); err != nil {
			return err
		}
	}
	// lpAttributeData {in} (1:{pointer=ref}*(1)[dim:0,size_is=AttributeSize](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AttributeData", sizeInfo[0])
		}
		o.AttributeData = make([]byte, sizeInfo[0])
		for i1 := range o.AttributeData {
			i1 := i1
			if err := w.ReadData(&o.AttributeData[i1]); err != nil {
				return err
			}
		}
	}
	// AttributeSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectAttributeWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSObjectAttributeWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNTMSObjectAttributeWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNTMSObjectAttributeWRequest structure represents the SetNtmsObjectAttributeW operation request
type SetNTMSObjectAttributeWRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpObjectId: A pointer to the identifier of the object for which to set private data.
	ObjectID *dtyp.GUID `idl:"name:lpObjectId" json:"object_id"`
	// dwType: A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// lpAttributeName: A null-terminated sequence of Unicode characters specifying the
	// name of the extended attribute to set. The client can give any name to the extended
	// attribute and MUST use the same name in the GetNtmsObjectAttributeW method.
	AttributeName string `idl:"name:lpAttributeName;string" json:"attribute_name"`
	// lpAttributeData: The buffer containing the attribute.
	AttributeData []byte `idl:"name:lpAttributeData;size_is:(AttributeSize)" json:"attribute_data"`
	// AttributeSize: The size of lpAttributeData.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | Access to the object is denied; other security errors are possible but indicate  |
	//	|                                    | a security subsystem error.                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY | An allocation failure occurred during processing.                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | The parameter is not valid.                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007007B ERROR_INVALID_NAME      | The attribute name is invalid or too long. The NTMS_MAXATTR_NAMELEN              |
	//	|                                    | value, defined in the Platform SDK file NTMSApi.h, specifies the maximum         |
	//	|                                    | null-terminated attribute name length.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | The object was not found.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800708CA ERROR_NOT_CONNECTED     | Unable to connect to the RSM service.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	AttributeSize uint32 `idl:"name:AttributeSize" json:"attribute_size"`
}

func (o *SetNTMSObjectAttributeWRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSObjectAttributeWOperation) *xxx_SetNTMSObjectAttributeWOperation {
	if op == nil {
		op = &xxx_SetNTMSObjectAttributeWOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.AttributeName = o.AttributeName
	op.AttributeData = o.AttributeData
	op.AttributeSize = o.AttributeSize
	return op
}

func (o *SetNTMSObjectAttributeWRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSObjectAttributeWOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.AttributeName = op.AttributeName
	o.AttributeData = op.AttributeData
	o.AttributeSize = op.AttributeSize
}
func (o *SetNTMSObjectAttributeWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNTMSObjectAttributeWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSObjectAttributeWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNTMSObjectAttributeWResponse structure represents the SetNtmsObjectAttributeW operation response
type SetNTMSObjectAttributeWResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetNtmsObjectAttributeW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNTMSObjectAttributeWResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSObjectAttributeWOperation) *xxx_SetNTMSObjectAttributeWOperation {
	if op == nil {
		op = &xxx_SetNTMSObjectAttributeWOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNTMSObjectAttributeWResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSObjectAttributeWOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNTMSObjectAttributeWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNTMSObjectAttributeWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSObjectAttributeWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateNTMSObjectOperation structure represents the EnumerateNtmsObject operation
type xxx_EnumerateNTMSObjectOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	ContainerID    *dtyp.GUID     `idl:"name:lpContainerId;pointer:unique" json:"container_id"`
	List           []*dtyp.GUID   `idl:"name:lpList;size_is:(lpdwListBufferSize);length_is:(lpdwListBufferSize)" json:"list"`
	ListBufferSize uint32         `idl:"name:lpdwListBufferSize" json:"list_buffer_size"`
	ListSize       uint32         `idl:"name:lpdwListSize" json:"list_size"`
	Type           uint32         `idl:"name:dwType" json:"type"`
	Options        uint32         `idl:"name:dwOptions" json:"options"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateNTMSObjectOperation) OpNum() int { return 6 }

func (o *xxx_EnumerateNTMSObjectOperation) OpName() string {
	return "/INtmsObjectManagement1/v0/EnumerateNtmsObject"
}

func (o *xxx_EnumerateNTMSObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateNTMSObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpContainerId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ContainerID != nil {
			_ptr_lpContainerId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ContainerID != nil {
					if err := o.ContainerID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ContainerID, _ptr_lpContainerId); err != nil {
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
	// lpdwListBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ListBufferSize); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateNTMSObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpContainerId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpContainerId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ContainerID == nil {
				o.ContainerID = &dtyp.GUID{}
			}
			if err := o.ContainerID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpContainerId := func(ptr interface{}) { o.ContainerID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.ContainerID, _s_lpContainerId, _ptr_lpContainerId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwListBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ListBufferSize); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateNTMSObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateNTMSObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpList {out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}[dim:0,size_is=lpdwListBufferSize,length_is=lpdwListBufferSize](struct))
	{
		dimSize1 := uint64(o.ListBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ListBufferSize)
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
		for i1 := range o.List {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.List[i1] != nil {
				if err := o.List[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.List); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwListSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ListSize); err != nil {
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

func (o *xxx_EnumerateNTMSObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpList {out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}[dim:0,size_is=lpdwListBufferSize,length_is=lpdwListBufferSize](struct))
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.List", sizeInfo[0])
		}
		o.List = make([]*dtyp.GUID, sizeInfo[0])
		for i1 := range o.List {
			i1 := i1
			if o.List[i1] == nil {
				o.List[i1] = &dtyp.GUID{}
			}
			if err := o.List[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwListSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ListSize); err != nil {
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

// EnumerateNTMSObjectRequest structure represents the EnumerateNtmsObject operation request
type EnumerateNTMSObjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpContainerId: A pointer to the GUID of the container for which to enumerate objects;
	// can be set to NULL to enumerate all objects of type dwType.
	ContainerID *dtyp.GUID `idl:"name:lpContainerId;pointer:unique" json:"container_id"`
	// lpdwListBufferSize: A pointer to the size, in bytes, of lpList.
	ListBufferSize uint32 `idl:"name:lpdwListBufferSize" json:"list_buffer_size"`
	// dwType: A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// dwOptions: This parameter is unused. It MUST be 0 and MUST be ignored on receipt.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED       | Access to an object was denied.                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY   | An allocation failure occurred during processing.                                |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER   | The lpdwListSize pointer is NULL, or lpContainerId is not of the object type     |
	//	|                                      | specified by dwType.                                                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The specified buffer size is too small. The required size is returned in the     |
	//	|                                      | lpdwListSize parameter.                                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND    | The lpContainerId is not the identifier of any container in the database.        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Options uint32 `idl:"name:dwOptions" json:"options"`
}

func (o *EnumerateNTMSObjectRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumerateNTMSObjectOperation) *xxx_EnumerateNTMSObjectOperation {
	if op == nil {
		op = &xxx_EnumerateNTMSObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ContainerID = o.ContainerID
	op.ListBufferSize = o.ListBufferSize
	op.Type = o.Type
	op.Options = o.Options
	return op
}

func (o *EnumerateNTMSObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateNTMSObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ContainerID = op.ContainerID
	o.ListBufferSize = op.ListBufferSize
	o.Type = op.Type
	o.Options = op.Options
}
func (o *EnumerateNTMSObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateNTMSObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateNTMSObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateNTMSObjectResponse structure represents the EnumerateNtmsObject operation response
type EnumerateNTMSObjectResponse struct {
	// XXX: lpdwListBufferSize is an implicit input depedency for output parameters
	ListBufferSize uint32 `idl:"name:lpdwListBufferSize" json:"list_buffer_size"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpList: An array of object identifiers.
	List []*dtyp.GUID `idl:"name:lpList;size_is:(lpdwListBufferSize);length_is:(lpdwListBufferSize)" json:"list"`
	// lpdwListSize: A pointer to the number of elements in lpList.
	ListSize uint32 `idl:"name:lpdwListSize" json:"list_size"`
	// Return: The EnumerateNtmsObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateNTMSObjectResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumerateNTMSObjectOperation) *xxx_EnumerateNTMSObjectOperation {
	if op == nil {
		op = &xxx_EnumerateNTMSObjectOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.ListBufferSize == uint32(0) {
		op.ListBufferSize = o.ListBufferSize
	}

	op.That = o.That
	op.List = o.List
	op.ListSize = o.ListSize
	op.Return = o.Return
	return op
}

func (o *EnumerateNTMSObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateNTMSObjectOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.ListBufferSize = op.ListBufferSize

	o.That = op.That
	o.List = op.List
	o.ListSize = op.ListSize
	o.Return = op.Return
}
func (o *EnumerateNTMSObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateNTMSObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateNTMSObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DisableNTMSObjectOperation structure represents the DisableNtmsObject operation
type xxx_DisableNTMSObjectOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Type     uint32         `idl:"name:dwType" json:"type"`
	ObjectID *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DisableNTMSObjectOperation) OpNum() int { return 7 }

func (o *xxx_DisableNTMSObjectOperation) OpName() string {
	return "/INtmsObjectManagement1/v0/DisableNtmsObject"
}

func (o *xxx_DisableNTMSObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DisableNTMSObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DisableNTMSObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DisableNTMSObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DisableNTMSObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DisableNTMSObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DisableNTMSObjectRequest structure represents the DisableNtmsObject operation request
type DisableNTMSObjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwType: One of the NTMS_DRIVE, NTMS_LIBRARY, or NTMS_PHYSICAL_MEDIA values from the
	// NtmsObjectsTypes (section 2.2.1.6) enumeration, specifying the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// lpObjectId: A pointer to the identifier of the object to disable.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | NTMS_CONTROL_ACCESS to the library containing the object is denied.              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | The parameter is NULL or invalid.                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D1 ERROR_LIBRARY_OFFLINE   | The lpObjectId parameter refers to an offline library that cannot be enabled or  |
	//	|                                    | disabled.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | The lpObjectId parameter is invalid.                                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database is inaccessible or damaged.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE     | The object is already disabled.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	ObjectID *dtyp.GUID `idl:"name:lpObjectId" json:"object_id"`
}

func (o *DisableNTMSObjectRequest) xxx_ToOp(ctx context.Context, op *xxx_DisableNTMSObjectOperation) *xxx_DisableNTMSObjectOperation {
	if op == nil {
		op = &xxx_DisableNTMSObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Type = o.Type
	op.ObjectID = o.ObjectID
	return op
}

func (o *DisableNTMSObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_DisableNTMSObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Type = op.Type
	o.ObjectID = op.ObjectID
}
func (o *DisableNTMSObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DisableNTMSObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DisableNTMSObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DisableNTMSObjectResponse structure represents the DisableNtmsObject operation response
type DisableNTMSObjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DisableNtmsObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DisableNTMSObjectResponse) xxx_ToOp(ctx context.Context, op *xxx_DisableNTMSObjectOperation) *xxx_DisableNTMSObjectOperation {
	if op == nil {
		op = &xxx_DisableNTMSObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DisableNTMSObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_DisableNTMSObjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DisableNTMSObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DisableNTMSObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DisableNTMSObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnableNTMSObjectOperation structure represents the EnableNtmsObject operation
type xxx_EnableNTMSObjectOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Type     uint32         `idl:"name:dwType" json:"type"`
	ObjectID *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EnableNTMSObjectOperation) OpNum() int { return 8 }

func (o *xxx_EnableNTMSObjectOperation) OpName() string {
	return "/INtmsObjectManagement1/v0/EnableNtmsObject"
}

func (o *xxx_EnableNTMSObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableNTMSObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EnableNTMSObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableNTMSObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableNTMSObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnableNTMSObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// EnableNTMSObjectRequest structure represents the EnableNtmsObject operation request
type EnableNTMSObjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwType: One of the NTMS_DRIVE, NTMS_LIBRARY, or NTMS_PHYSICAL_MEDIA values from the
	// NtmsObjectsTypes (section 2.2.1.6) enumeration, specifying the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// lpObjectId: A pointer to the identifier of the object to enable.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | NTMS_CONTROL_ACCESS to the library containing the object is denied.              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | The parameter is NULL or invalid.                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D1 ERROR_LIBRARY_OFFLINE   | The lpObjectId parameter refers to an offline library that cannot be enabled or  |
	//	|                                    | disabled.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | The lpObjectId parameter is invalid.                                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database is inaccessible or damaged.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE     | The object is already enabled.                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	ObjectID *dtyp.GUID `idl:"name:lpObjectId" json:"object_id"`
}

func (o *EnableNTMSObjectRequest) xxx_ToOp(ctx context.Context, op *xxx_EnableNTMSObjectOperation) *xxx_EnableNTMSObjectOperation {
	if op == nil {
		op = &xxx_EnableNTMSObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Type = o.Type
	op.ObjectID = o.ObjectID
	return op
}

func (o *EnableNTMSObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_EnableNTMSObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Type = op.Type
	o.ObjectID = op.ObjectID
}
func (o *EnableNTMSObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnableNTMSObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnableNTMSObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnableNTMSObjectResponse structure represents the EnableNtmsObject operation response
type EnableNTMSObjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EnableNtmsObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnableNTMSObjectResponse) xxx_ToOp(ctx context.Context, op *xxx_EnableNTMSObjectOperation) *xxx_EnableNTMSObjectOperation {
	if op == nil {
		op = &xxx_EnableNTMSObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *EnableNTMSObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_EnableNTMSObjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *EnableNTMSObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnableNTMSObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnableNTMSObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
