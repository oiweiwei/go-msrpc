package ireplicationutil

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// IReplicationUtil interface identifier 98315903-7be5-11d2-adc1-00a02463d6e7
	ReplicationUtilIID = &dcom.IID{Data1: 0x98315903, Data2: 0x7be5, Data3: 0x11d2, Data4: []byte{0xad, 0xc1, 0x00, 0xa0, 0x24, 0x63, 0xd6, 0xe7}}
	// Syntax UUID
	ReplicationUtilSyntaxUUID = &uuid.UUID{TimeLow: 0x98315903, TimeMid: 0x7be5, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xad, ClockSeqLow: 0xc1, Node: [6]uint8{0x0, 0xa0, 0x24, 0x63, 0xd6, 0xe7}}
	// Syntax ID
	ReplicationUtilSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ReplicationUtilSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IReplicationUtil interface.
type ReplicationUtilClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a replication client application to create a Common Internet
	// File System (CIFS) [MS-CIFS] file share for copying installer package files.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CreateShare(context.Context, *CreateShareRequest, ...dcerpc.CallOption) (*CreateShareResponse, error)

	// This method is called by a replication client application to create an empty directory
	// to back up a replication file share.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CreateEmptyDir(context.Context, *CreateEmptyDirRequest, ...dcerpc.CallOption) (*CreateEmptyDirResponse, error)

	// This method is called by a replication client application to remove a file share
	// that was used during replication and is no longer needed.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	RemoveShare(context.Context, *RemoveShareRequest, ...dcerpc.CallOption) (*RemoveShareResponse, error)

	// This method is called by a replication client application to request that a server
	// perform the actions necessary to begin a replication operation in which the server
	// is a replication target. This typically happens after the application has copied
	// import package files for the conglomerations to be copied to a replication target
	// file share on the server. As part of the handling of the method, the server sets
	// up a replication working directory and file share. The server's handling of the method
	// might also include management of replication history and backup state.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	BeginReplicationAsTarget(context.Context, *BeginReplicationAsTargetRequest, ...dcerpc.CallOption) (*BeginReplicationAsTargetResponse, error)

	// This method is called by a replication client application to obtain the value of
	// the Password property of a conglomeration.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	QueryConglomerationPassword(context.Context, *QueryConglomerationPasswordRequest, ...dcerpc.CallOption) (*QueryConglomerationPasswordResponse, error)

	// This method is called by a replication client application to ensure that the server's
	// base replication directory exists and to get its path.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CreateReplicationDir(context.Context, *CreateReplicationDirRequest, ...dcerpc.CallOption) (*CreateReplicationDirResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ReplicationUtilClient
}

type xxx_DefaultReplicationUtilClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultReplicationUtilClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultReplicationUtilClient) CreateShare(ctx context.Context, in *CreateShareRequest, opts ...dcerpc.CallOption) (*CreateShareResponse, error) {
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
	out := &CreateShareResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReplicationUtilClient) CreateEmptyDir(ctx context.Context, in *CreateEmptyDirRequest, opts ...dcerpc.CallOption) (*CreateEmptyDirResponse, error) {
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
	out := &CreateEmptyDirResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReplicationUtilClient) RemoveShare(ctx context.Context, in *RemoveShareRequest, opts ...dcerpc.CallOption) (*RemoveShareResponse, error) {
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
	out := &RemoveShareResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReplicationUtilClient) BeginReplicationAsTarget(ctx context.Context, in *BeginReplicationAsTargetRequest, opts ...dcerpc.CallOption) (*BeginReplicationAsTargetResponse, error) {
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
	out := &BeginReplicationAsTargetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReplicationUtilClient) QueryConglomerationPassword(ctx context.Context, in *QueryConglomerationPasswordRequest, opts ...dcerpc.CallOption) (*QueryConglomerationPasswordResponse, error) {
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
	out := &QueryConglomerationPasswordResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReplicationUtilClient) CreateReplicationDir(ctx context.Context, in *CreateReplicationDirRequest, opts ...dcerpc.CallOption) (*CreateReplicationDirResponse, error) {
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
	out := &CreateReplicationDirResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReplicationUtilClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultReplicationUtilClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultReplicationUtilClient) IPID(ctx context.Context, ipid *dcom.IPID) ReplicationUtilClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultReplicationUtilClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewReplicationUtilClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ReplicationUtilClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ReplicationUtilSyntaxV0_0))...)
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
	return &xxx_DefaultReplicationUtilClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CreateShareOperation structure represents the CreateShare operation
type xxx_CreateShareOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ShareName string         `idl:"name:pwszShareName" json:"share_name"`
	Path      string         `idl:"name:pwszPath" json:"path"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateShareOperation) OpNum() int { return 3 }

func (o *xxx_CreateShareOperation) OpName() string { return "/IReplicationUtil/v0/CreateShare" }

func (o *xxx_CreateShareOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateShareOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszShareName {in} (1:{alias=LPCWSTR}*(1)[dim:0,string](wchar))
	{
		if err := ndr.WriteUTF16String(ctx, w, o.ShareName); err != nil {
			return err
		}
	}
	// pwszPath {in} (1:{alias=LPCWSTR}*(1)[dim:0,string](wchar))
	{
		if err := ndr.WriteUTF16String(ctx, w, o.Path); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateShareOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszShareName {in} (1:{alias=LPCWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		if err := ndr.ReadUTF16String(ctx, w, &o.ShareName); err != nil {
			return err
		}
	}
	// pwszPath {in} (1:{alias=LPCWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		if err := ndr.ReadUTF16String(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateShareOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateShareOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateShareOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateShareRequest structure represents the CreateShare operation request
type CreateShareRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszShareName: The share name (as specified for a path in UNC syntax) of the file
	// share to create. This MUST be a SourceShareName or TargetNewShareName, according
	// to the following ABNF syntax, as specified in [RFC4234].
	//
	// SourceShareName = "ReplicaSource" CurlyBracedGuidString TargetNewShareName = "ReplicaTargetNew"
	// CurlyBracedGuidString
	//
	// Where CurlyBracedGuidString is in Curly Braced GUID String Syntax ([MS-DTYP] section
	// 2.3.4.3).
	//
	// These formats have the following usage.
	//
	//	+--------------------+-----------------------------------------------+
	//	|                    |                                               |
	//	|       FORMAT       |                     USAGE                     |
	//	|                    |                                               |
	//	+--------------------+-----------------------------------------------+
	//	+--------------------+-----------------------------------------------+
	//	| SourceShareName    | Used when the server is a replication source. |
	//	+--------------------+-----------------------------------------------+
	//	| TargetNewShareName | Used when the server is a replication target. |
	//	+--------------------+-----------------------------------------------+
	ShareName string `idl:"name:pwszShareName" json:"share_name"`
	// pwszPath: An ImplementationSpecificPathProperty (section 2.2.2.2) representing the
	// path to the directory that is to back the file share. This MUST be derived from the
	// server's base replication directory path by appending one of the following strings.
	//
	//	+------------------+-------------------------------------+
	//	|                  |                                     |
	//	|      VALUE       |               MEANING               |
	//	|                  |                                     |
	//	+------------------+-------------------------------------+
	//	+------------------+-------------------------------------+
	//	| "\ReplicaSource" | The server is a replication source. |
	//	+------------------+-------------------------------------+
	//	| "\ReplicaNew"    | The server is a replication target. |
	//	+------------------+-------------------------------------+
	Path string `idl:"name:pwszPath" json:"path"`
}

func (o *CreateShareRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateShareOperation) *xxx_CreateShareOperation {
	if op == nil {
		op = &xxx_CreateShareOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ShareName = o.ShareName
	op.Path = o.Path
	return op
}

func (o *CreateShareRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateShareOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ShareName = op.ShareName
	o.Path = op.Path
}
func (o *CreateShareRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateShareRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateShareOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateShareResponse structure represents the CreateShare operation response
type CreateShareResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateShare return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateShareResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateShareOperation) *xxx_CreateShareOperation {
	if op == nil {
		op = &xxx_CreateShareOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CreateShareResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateShareOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateShareResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateShareResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateShareOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateEmptyDirOperation structure represents the CreateEmptyDir operation
type xxx_CreateEmptyDirOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   string         `idl:"name:pwszPath" json:"path"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateEmptyDirOperation) OpNum() int { return 4 }

func (o *xxx_CreateEmptyDirOperation) OpName() string { return "/IReplicationUtil/v0/CreateEmptyDir" }

func (o *xxx_CreateEmptyDirOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateEmptyDirOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszPath {in} (1:{alias=LPCWSTR}*(1)[dim:0,string](wchar))
	{
		if err := ndr.WriteUTF16String(ctx, w, o.Path); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateEmptyDirOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszPath {in} (1:{alias=LPCWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		if err := ndr.ReadUTF16String(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateEmptyDirOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateEmptyDirOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateEmptyDirOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateEmptyDirRequest structure represents the CreateEmptyDir operation request
type CreateEmptyDirRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszPath: An ImplementationSpecificPathProperty (section 2.2.2.2) representing the
	// path to the directory that is to be created by the server. This MUST be derived from
	// the server’s base replication directory path by appending one of the following
	// strings.
	//
	//	+------------------+-------------------------------------+
	//	|                  |                                     |
	//	|      VALUE       |               MEANING               |
	//	|                  |                                     |
	//	+------------------+-------------------------------------+
	//	+------------------+-------------------------------------+
	//	| "\ReplicaSource" | The server is a replication source. |
	//	+------------------+-------------------------------------+
	//	| "\ReplicaNew"    | The server is a replication target. |
	//	+------------------+-------------------------------------+
	Path string `idl:"name:pwszPath" json:"path"`
}

func (o *CreateEmptyDirRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateEmptyDirOperation) *xxx_CreateEmptyDirOperation {
	if op == nil {
		op = &xxx_CreateEmptyDirOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Path = o.Path
	return op
}

func (o *CreateEmptyDirRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateEmptyDirOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *CreateEmptyDirRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateEmptyDirRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateEmptyDirOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateEmptyDirResponse structure represents the CreateEmptyDir operation response
type CreateEmptyDirResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateEmptyDir return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateEmptyDirResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateEmptyDirOperation) *xxx_CreateEmptyDirOperation {
	if op == nil {
		op = &xxx_CreateEmptyDirOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CreateEmptyDirResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateEmptyDirOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateEmptyDirResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateEmptyDirResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateEmptyDirOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveShareOperation structure represents the RemoveShare operation
type xxx_RemoveShareOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ShareName string         `idl:"name:pwszShareName" json:"share_name"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveShareOperation) OpNum() int { return 5 }

func (o *xxx_RemoveShareOperation) OpName() string { return "/IReplicationUtil/v0/RemoveShare" }

func (o *xxx_RemoveShareOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveShareOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszShareName {in} (1:{alias=LPCWSTR}*(1)[dim:0,string](wchar))
	{
		if err := ndr.WriteUTF16String(ctx, w, o.ShareName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveShareOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszShareName {in} (1:{alias=LPCWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		if err := ndr.ReadUTF16String(ctx, w, &o.ShareName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveShareOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveShareOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoveShareOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RemoveShareRequest structure represents the RemoveShare operation request
type RemoveShareRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszShareName: The share name (as specified for a path in UNC syntax) of the file
	// share to remove. This MUST be a SourceShareName, TargetNewShareName, or TargetCurrentShareName,
	// according to the following Augmented Backus-Naur Form (ABNF) syntax, as specified
	// in [RFC4234].
	//
	// SourceShareName = "ReplicaSource" CurlyBracedGuidString TargetNewShareName = "ReplicaTargetNew"
	// CurlyBracedGuidString TargetCurrentShareName = "ReplicaTargetCurrent"
	//
	// Where CurlyBracedGuidString is in Curly Braced GUID String Syntax ([MS-DTYP] section
	// 2.3.4.3).
	//
	// These formats have the following usage.
	//
	//	+------------------------+----------------------------------------------------------------------------------+
	//	|                        |                                                                                  |
	//	|         FORMAT         |                                      USAGE                                       |
	//	|                        |                                                                                  |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| SourceShareName        | Used when the server is a replication source and the replication client          |
	//	|                        | application is finished copying files from the server.                           |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| TargetNewShareName     | Used when the server is a replication target and the replication client          |
	//	|                        | application is finished copying files to the server.                             |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| TargetCurrentShareName | Used when the server is a replication target and the replication client          |
	//	|                        | application is finished importing conglomerations from the share.                |
	//	+------------------------+----------------------------------------------------------------------------------+
	ShareName string `idl:"name:pwszShareName" json:"share_name"`
}

func (o *RemoveShareRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoveShareOperation) *xxx_RemoveShareOperation {
	if op == nil {
		op = &xxx_RemoveShareOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ShareName = o.ShareName
	return op
}

func (o *RemoveShareRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveShareOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ShareName = op.ShareName
}
func (o *RemoveShareRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoveShareRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveShareOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveShareResponse structure represents the RemoveShare operation response
type RemoveShareResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RemoveShare return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveShareResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoveShareOperation) *xxx_RemoveShareOperation {
	if op == nil {
		op = &xxx_RemoveShareOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RemoveShareResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveShareOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RemoveShareResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoveShareResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveShareOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BeginReplicationAsTargetOperation structure represents the BeginReplicationAsTarget operation
type xxx_BeginReplicationAsTargetOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	BaseReplicationDir string         `idl:"name:pwszBaseReplicationDir" json:"base_replication_dir"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BeginReplicationAsTargetOperation) OpNum() int { return 6 }

func (o *xxx_BeginReplicationAsTargetOperation) OpName() string {
	return "/IReplicationUtil/v0/BeginReplicationAsTarget"
}

func (o *xxx_BeginReplicationAsTargetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginReplicationAsTargetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszBaseReplicationDir {in} (1:{alias=LPCWSTR}*(1)[dim:0,string](wchar))
	{
		if err := ndr.WriteUTF16String(ctx, w, o.BaseReplicationDir); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginReplicationAsTargetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszBaseReplicationDir {in} (1:{alias=LPCWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		if err := ndr.ReadUTF16String(ctx, w, &o.BaseReplicationDir); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginReplicationAsTargetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginReplicationAsTargetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BeginReplicationAsTargetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// BeginReplicationAsTargetRequest structure represents the BeginReplicationAsTarget operation request
type BeginReplicationAsTargetRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszBaseReplicationDir: The server's base replication directory path.
	BaseReplicationDir string `idl:"name:pwszBaseReplicationDir" json:"base_replication_dir"`
}

func (o *BeginReplicationAsTargetRequest) xxx_ToOp(ctx context.Context, op *xxx_BeginReplicationAsTargetOperation) *xxx_BeginReplicationAsTargetOperation {
	if op == nil {
		op = &xxx_BeginReplicationAsTargetOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BaseReplicationDir = o.BaseReplicationDir
	return op
}

func (o *BeginReplicationAsTargetRequest) xxx_FromOp(ctx context.Context, op *xxx_BeginReplicationAsTargetOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BaseReplicationDir = op.BaseReplicationDir
}
func (o *BeginReplicationAsTargetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BeginReplicationAsTargetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BeginReplicationAsTargetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BeginReplicationAsTargetResponse structure represents the BeginReplicationAsTarget operation response
type BeginReplicationAsTargetResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The BeginReplicationAsTarget return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BeginReplicationAsTargetResponse) xxx_ToOp(ctx context.Context, op *xxx_BeginReplicationAsTargetOperation) *xxx_BeginReplicationAsTargetOperation {
	if op == nil {
		op = &xxx_BeginReplicationAsTargetOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *BeginReplicationAsTargetResponse) xxx_FromOp(ctx context.Context, op *xxx_BeginReplicationAsTargetOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *BeginReplicationAsTargetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BeginReplicationAsTargetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BeginReplicationAsTargetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryConglomerationPasswordOperation structure represents the QueryConglomerationPassword operation
type xxx_QueryConglomerationPasswordOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConglomerationID *dtyp.GUID     `idl:"name:ConglomerationId" json:"conglomeration_id"`
	Password         []byte         `idl:"name:ppvPassword;size_is:(, pcbPassword)" json:"password"`
	PasswordLength   uint32         `idl:"name:pcbPassword" json:"password_length"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryConglomerationPasswordOperation) OpNum() int { return 7 }

func (o *xxx_QueryConglomerationPasswordOperation) OpName() string {
	return "/IReplicationUtil/v0/QueryConglomerationPassword"
}

func (o *xxx_QueryConglomerationPasswordOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryConglomerationPasswordOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ConglomerationId {in} (1:{alias=REFGUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ConglomerationID != nil {
			if err := o.ConglomerationID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_QueryConglomerationPasswordOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ConglomerationId {in} (1:{alias=REFGUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ConglomerationID == nil {
			o.ConglomerationID = &dtyp.GUID{}
		}
		if err := o.ConglomerationID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryConglomerationPasswordOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Password != nil && o.PasswordLength == 0 {
		o.PasswordLength = uint32(len(o.Password))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryConglomerationPasswordOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppvPassword {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbPassword](char))
	{
		if o.Password != nil || o.PasswordLength > 0 {
			_ptr_ppvPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.PasswordLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Password {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Password[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Password); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_ppvPassword); err != nil {
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
	// pcbPassword {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.PasswordLength); err != nil {
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

func (o *xxx_QueryConglomerationPasswordOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppvPassword {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbPassword](char))
	{
		_ptr_ppvPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Password", sizeInfo[0])
			}
			o.Password = make([]byte, sizeInfo[0])
			for i1 := range o.Password {
				i1 := i1
				if err := w.ReadData(&o.Password[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppvPassword := func(ptr interface{}) { o.Password = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Password, _s_ppvPassword, _ptr_ppvPassword); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbPassword {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.PasswordLength); err != nil {
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

// QueryConglomerationPasswordRequest structure represents the QueryConglomerationPassword operation request
type QueryConglomerationPasswordRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ConglomerationId: The conglomeration identifier of the conglomeration whose Password
	// property is requested.
	ConglomerationID *dtyp.GUID `idl:"name:ConglomerationId" json:"conglomeration_id"`
}

func (o *QueryConglomerationPasswordRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryConglomerationPasswordOperation) *xxx_QueryConglomerationPasswordOperation {
	if op == nil {
		op = &xxx_QueryConglomerationPasswordOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ConglomerationID = o.ConglomerationID
	return op
}

func (o *QueryConglomerationPasswordRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryConglomerationPasswordOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConglomerationID = op.ConglomerationID
}
func (o *QueryConglomerationPasswordRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryConglomerationPasswordRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryConglomerationPasswordOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryConglomerationPasswordResponse structure represents the QueryConglomerationPassword operation response
type QueryConglomerationPasswordResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppvPassword: A pointer to a variable that, upon successful completion, MUST be set
	// to an array of bytes containing a null-terminated array of wchar_t in little-endian
	// byte order.
	Password []byte `idl:"name:ppvPassword;size_is:(, pcbPassword)" json:"password"`
	// pcbPassword: A pointer to a variable that, upon successful completion, MUST be set
	// to the length of ppvPassword.
	PasswordLength uint32 `idl:"name:pcbPassword" json:"password_length"`
	// Return: The QueryConglomerationPassword return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryConglomerationPasswordResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryConglomerationPasswordOperation) *xxx_QueryConglomerationPasswordOperation {
	if op == nil {
		op = &xxx_QueryConglomerationPasswordOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Password = o.Password
	op.PasswordLength = o.PasswordLength
	op.Return = o.Return
	return op
}

func (o *QueryConglomerationPasswordResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryConglomerationPasswordOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Password = op.Password
	o.PasswordLength = op.PasswordLength
	o.Return = op.Return
}
func (o *QueryConglomerationPasswordResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryConglomerationPasswordResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryConglomerationPasswordOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateReplicationDirOperation structure represents the CreateReplicationDir operation
type xxx_CreateReplicationDirOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	BaseReplicationDir string         `idl:"name:ppwszBaseReplicationDir" json:"base_replication_dir"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateReplicationDirOperation) OpNum() int { return 8 }

func (o *xxx_CreateReplicationDirOperation) OpName() string {
	return "/IReplicationUtil/v0/CreateReplicationDir"
}

func (o *xxx_CreateReplicationDirOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateReplicationDirOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateReplicationDirOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CreateReplicationDirOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateReplicationDirOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppwszBaseReplicationDir {out} (1:{pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string](wchar))
	{
		if o.BaseReplicationDir != "" {
			_ptr_ppwszBaseReplicationDir := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.BaseReplicationDir); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.BaseReplicationDir, _ptr_ppwszBaseReplicationDir); err != nil {
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

func (o *xxx_CreateReplicationDirOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppwszBaseReplicationDir {out} (1:{pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		_ptr_ppwszBaseReplicationDir := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.BaseReplicationDir); err != nil {
				return err
			}
			return nil
		})
		_s_ppwszBaseReplicationDir := func(ptr interface{}) { o.BaseReplicationDir = *ptr.(*string) }
		if err := w.ReadPointer(&o.BaseReplicationDir, _s_ppwszBaseReplicationDir, _ptr_ppwszBaseReplicationDir); err != nil {
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

// CreateReplicationDirRequest structure represents the CreateReplicationDir operation request
type CreateReplicationDirRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CreateReplicationDirRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateReplicationDirOperation) *xxx_CreateReplicationDirOperation {
	if op == nil {
		op = &xxx_CreateReplicationDirOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CreateReplicationDirRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateReplicationDirOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CreateReplicationDirRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateReplicationDirRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateReplicationDirOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateReplicationDirResponse structure represents the CreateReplicationDir operation response
type CreateReplicationDirResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppwszBaseReplicationDir: A pointer to a variable that, upon successful completion,
	// MUST contain the server's base replication directory path.
	BaseReplicationDir string `idl:"name:ppwszBaseReplicationDir" json:"base_replication_dir"`
	// Return: The CreateReplicationDir return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateReplicationDirResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateReplicationDirOperation) *xxx_CreateReplicationDirOperation {
	if op == nil {
		op = &xxx_CreateReplicationDirOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.BaseReplicationDir = o.BaseReplicationDir
	op.Return = o.Return
	return op
}

func (o *CreateReplicationDirResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateReplicationDirOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.BaseReplicationDir = op.BaseReplicationDir
	o.Return = op.Return
}
func (o *CreateReplicationDirResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateReplicationDirResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateReplicationDirOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
