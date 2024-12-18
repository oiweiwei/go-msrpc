package idmremoteserver

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
	GoPackage = "dcom/dmrp"
)

var (
	// IDMRemoteServer interface identifier 3a410f21-553f-11d1-8e5e-00a0c92c9d5d
	IDMRemoteServerIID = &dcom.IID{Data1: 0x3a410f21, Data2: 0x553f, Data3: 0x11d1, Data4: []byte{0x8e, 0x5e, 0x00, 0xa0, 0xc9, 0x2c, 0x9d, 0x5d}}
	// Syntax UUID
	IDMRemoteServerSyntaxUUID = &uuid.UUID{TimeLow: 0x3a410f21, TimeMid: 0x553f, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8e, ClockSeqLow: 0x5e, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x2c, 0x9d, 0x5d}}
	// Syntax ID
	IDMRemoteServerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: IDMRemoteServerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IDMRemoteServer interface.
type IDMRemoteServerClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The CreateRemoteObject method creates a disk management server, on the remote machine
	// specified by RemoteComputerName, by invoking DCOM with the class GUID of Disk Management
	// server and the name of the remote machine, which starts the disk management server
	// on the remote machine. The method negotiates for the interface as described in section
	// 3.1.3, and as illustrated in section 4. The client holds a reference to the IDMRemoteServer
	// interface binding on the server, until the client has received an IVolumeClient,
	// or IVolumeClient3 interface binding to the remote server. The client MAY then release
	// the IDMRemoteServer interface on the server.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	CreateRemoteObject(context.Context, *CreateRemoteObjectRequest, ...dcerpc.CallOption) (*CreateRemoteObjectResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) IDMRemoteServerClient
}

type xxx_DefaultIDMRemoteServerClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultIDMRemoteServerClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultIDMRemoteServerClient) CreateRemoteObject(ctx context.Context, in *CreateRemoteObjectRequest, opts ...dcerpc.CallOption) (*CreateRemoteObjectResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &CreateRemoteObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIDMRemoteServerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIDMRemoteServerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultIDMRemoteServerClient) IPID(ctx context.Context, ipid *dcom.IPID) IDMRemoteServerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultIDMRemoteServerClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewIDMRemoteServerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IDMRemoteServerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IDMRemoteServerSyntaxV0_0))...)
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
	return &xxx_DefaultIDMRemoteServerClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CreateRemoteObjectOperation structure represents the CreateRemoteObject operation
type xxx_CreateRemoteObjectOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxCount           uint32         `idl:"name:cMax" json:"max_count"`
	RemoteComputerName string         `idl:"name:RemoteComputerName;max_is:(cMax)" json:"remote_computer_name"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateRemoteObjectOperation) OpNum() int { return 3 }

func (o *xxx_CreateRemoteObjectOperation) OpName() string {
	return "/IDMRemoteServer/v0/CreateRemoteObject"
}

func (o *xxx_CreateRemoteObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.RemoteComputerName != "" && o.MaxCount == 0 {
		o.MaxCount = uint32((len(o.RemoteComputerName) - 1))
		if len(o.RemoteComputerName) < 1 {
			o.MaxCount = uint32(0)
		}
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRemoteObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// cMax {in} (1:(uint32))
	{
		if err := w.WriteData(o.MaxCount); err != nil {
			return err
		}
	}
	// RemoteComputerName {in} (1:{pointer=ref}*(1)[dim:0,max_is=cMax,string](wchar))
	{
		dimSize1 := uint64((o.MaxCount + 1))
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_RemoteComputerName_buf := utf16.Encode([]rune(o.RemoteComputerName))
		if uint64(len(_RemoteComputerName_buf)) > sizeInfo[0] {
			_RemoteComputerName_buf = _RemoteComputerName_buf[:sizeInfo[0]]
		}
		for i1 := range _RemoteComputerName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_RemoteComputerName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_RemoteComputerName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreateRemoteObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// cMax {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MaxCount); err != nil {
			return err
		}
	}
	// RemoteComputerName {in} (1:{pointer=ref}*(1)[dim:0,max_is=cMax,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _RemoteComputerName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _RemoteComputerName_buf", sizeInfo[0])
		}
		_RemoteComputerName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _RemoteComputerName_buf {
			i1 := i1
			if err := w.ReadData(&_RemoteComputerName_buf[i1]); err != nil {
				return err
			}
		}
		o.RemoteComputerName = strings.TrimRight(string(utf16.Decode(_RemoteComputerName_buf)), ndr.ZeroString)
	}
	return nil
}

func (o *xxx_CreateRemoteObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRemoteObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateRemoteObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateRemoteObjectRequest structure represents the CreateRemoteObject operation request
type CreateRemoteObjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// cMax: Length of RemoteComputerName (in Unicode characters), including the terminating
	// null character.
	MaxCount uint32 `idl:"name:cMax" json:"max_count"`
	// RemoteComputerName: Null-terminated Unicode string that specifies the name of the
	// computer on which the server is to be activated. All UNC names ("\\server" or "server")
	// and DNS names ("domain.com", "example.microsoft.com", or "135.5.33.19") are allowed.
	RemoteComputerName string `idl:"name:RemoteComputerName;max_is:(cMax)" json:"remote_computer_name"`
}

func (o *CreateRemoteObjectRequest) xxx_ToOp(ctx context.Context) *xxx_CreateRemoteObjectOperation {
	if o == nil {
		return &xxx_CreateRemoteObjectOperation{}
	}
	return &xxx_CreateRemoteObjectOperation{
		This:               o.This,
		MaxCount:           o.MaxCount,
		RemoteComputerName: o.RemoteComputerName,
	}
}

func (o *CreateRemoteObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateRemoteObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MaxCount = op.MaxCount
	o.RemoteComputerName = op.RemoteComputerName
}
func (o *CreateRemoteObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateRemoteObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateRemoteObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateRemoteObjectResponse structure represents the CreateRemoteObject operation response
type CreateRemoteObjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateRemoteObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateRemoteObjectResponse) xxx_ToOp(ctx context.Context) *xxx_CreateRemoteObjectOperation {
	if o == nil {
		return &xxx_CreateRemoteObjectOperation{}
	}
	return &xxx_CreateRemoteObjectOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CreateRemoteObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateRemoteObjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateRemoteObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateRemoteObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateRemoteObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
