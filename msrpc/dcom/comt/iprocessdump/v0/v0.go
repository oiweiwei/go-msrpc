package iprocessdump

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
	GoPackage = "dcom/comt"
)

var (
	// IProcessDump interface identifier 23c9dd26-2355-4fe2-84de-f779a238adbd
	ProcessDumpIID = &dcom.IID{Data1: 0x23c9dd26, Data2: 0x2355, Data3: 0x4fe2, Data4: []byte{0x84, 0xde, 0xf7, 0x79, 0xa2, 0x38, 0xad, 0xbd}}
	// Syntax UUID
	ProcessDumpSyntaxUUID = &uuid.UUID{TimeLow: 0x23c9dd26, TimeMid: 0x2355, TimeHiAndVersion: 0x4fe2, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0xde, Node: [6]uint8{0xf7, 0x79, 0xa2, 0x38, 0xad, 0xbd}}
	// Syntax ID
	ProcessDumpSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ProcessDumpSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IProcessDump interface.
type ProcessDumpClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// This method is called by a client to determine whether or not the COMT server supports
	// process dump.
	//
	// This method has no parameters.
	//
	// Return Values: This method returns S_OK (0x00000000) if the COMT server supports
	// process dump, and MUST return S_FALSE (0x00000001) if not.
	IsSupported(context.Context, *IsSupportedRequest, ...dcerpc.CallOption) (*IsSupportedResponse, error)

	// This method is called by a client to request a process dump for the process containing
	// a particular instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result (as specified in [MS-ERREF] section 2.1) on failure.
	DumpProcess(context.Context, *DumpProcessRequest, ...dcerpc.CallOption) (*DumpProcessResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ProcessDumpClient
}

type xxx_DefaultProcessDumpClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultProcessDumpClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultProcessDumpClient) IsSupported(ctx context.Context, in *IsSupportedRequest, opts ...dcerpc.CallOption) (*IsSupportedResponse, error) {
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
	out := &IsSupportedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultProcessDumpClient) DumpProcess(ctx context.Context, in *DumpProcessRequest, opts ...dcerpc.CallOption) (*DumpProcessResponse, error) {
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
	out := &DumpProcessResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultProcessDumpClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultProcessDumpClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultProcessDumpClient) IPID(ctx context.Context, ipid *dcom.IPID) ProcessDumpClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultProcessDumpClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewProcessDumpClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ProcessDumpClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ProcessDumpSyntaxV0_0))...)
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
	return &xxx_DefaultProcessDumpClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_IsSupportedOperation structure represents the IsSupported operation
type xxx_IsSupportedOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsSupportedOperation) OpNum() int { return 7 }

func (o *xxx_IsSupportedOperation) OpName() string { return "/IProcessDump/v0/IsSupported" }

func (o *xxx_IsSupportedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsSupportedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsSupportedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_IsSupportedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsSupportedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsSupportedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// IsSupportedRequest structure represents the IsSupported operation request
type IsSupportedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *IsSupportedRequest) xxx_ToOp(ctx context.Context, op *xxx_IsSupportedOperation) *xxx_IsSupportedOperation {
	if op == nil {
		op = &xxx_IsSupportedOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *IsSupportedRequest) xxx_FromOp(ctx context.Context, op *xxx_IsSupportedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *IsSupportedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsSupportedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsSupportedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsSupportedResponse structure represents the IsSupported operation response
type IsSupportedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The IsSupported return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsSupportedResponse) xxx_ToOp(ctx context.Context, op *xxx_IsSupportedOperation) *xxx_IsSupportedOperation {
	if op == nil {
		op = &xxx_IsSupportedOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *IsSupportedResponse) xxx_FromOp(ctx context.Context, op *xxx_IsSupportedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *IsSupportedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsSupportedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsSupportedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DumpProcessOperation structure represents the DumpProcess operation
type xxx_DumpProcessOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ContainerID *oaut.String   `idl:"name:bstrContainerID" json:"container_id"`
	Directory   *oaut.String   `idl:"name:bstrDirectory" json:"directory"`
	MaxFiles    uint32         `idl:"name:dwMaxFiles" json:"max_files"`
	DumpFile    *oaut.String   `idl:"name:pbstrDumpFile" json:"dump_file"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DumpProcessOperation) OpNum() int { return 8 }

func (o *xxx_DumpProcessOperation) OpName() string { return "/IProcessDump/v0/DumpProcess" }

func (o *xxx_DumpProcessOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DumpProcessOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrContainerID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ContainerID != nil {
			_ptr_bstrContainerID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ContainerID != nil {
					if err := o.ContainerID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ContainerID, _ptr_bstrContainerID); err != nil {
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
	// bstrDirectory {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Directory != nil {
			_ptr_bstrDirectory := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Directory != nil {
					if err := o.Directory.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Directory, _ptr_bstrDirectory); err != nil {
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
	// dwMaxFiles {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxFiles); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DumpProcessOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrContainerID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrContainerID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ContainerID == nil {
				o.ContainerID = &oaut.String{}
			}
			if err := o.ContainerID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrContainerID := func(ptr interface{}) { o.ContainerID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ContainerID, _s_bstrContainerID, _ptr_bstrContainerID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrDirectory {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrDirectory := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Directory == nil {
				o.Directory = &oaut.String{}
			}
			if err := o.Directory.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrDirectory := func(ptr interface{}) { o.Directory = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Directory, _s_bstrDirectory, _ptr_bstrDirectory); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMaxFiles {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxFiles); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DumpProcessOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DumpProcessOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrDumpFile {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DumpFile != nil {
			_ptr_pbstrDumpFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DumpFile != nil {
					if err := o.DumpFile.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DumpFile, _ptr_pbstrDumpFile); err != nil {
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

func (o *xxx_DumpProcessOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrDumpFile {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrDumpFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DumpFile == nil {
				o.DumpFile = &oaut.String{}
			}
			if err := o.DumpFile.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrDumpFile := func(ptr interface{}) { o.DumpFile = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DumpFile, _s_pbstrDumpFile, _ptr_pbstrDumpFile); err != nil {
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

// DumpProcessRequest structure represents the DumpProcess operation request
type DumpProcessRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrContainerID: The CurlyBraceGuidString (section 2.2.1) representation of a container
	// identifier for a distinguished container.
	ContainerID *oaut.String `idl:"name:bstrContainerID" json:"container_id"`
	// bstrDirectory: Either a path, in the convention of the server's file system, to a
	// location in which the file produced by process dump is to be written, or NULL to
	// indicate that the client wants the COMT server to write the file to an implementation-specific
	// default location.
	Directory *oaut.String `idl:"name:bstrDirectory" json:"directory"`
	// dwMaxFiles: The maximum number of process dump files associated with the conglomeration
	// of the instance container identified by the bstrContainerID parameter that the client
	// requests the COMT server to leave in the location specified by the bstrDirectory
	// parameter before the server begins deleting previously written files. A value of
	// 0x00000000 indicates that the COMT server is to use an implementation-specific default
	// limit.
	MaxFiles uint32 `idl:"name:dwMaxFiles" json:"max_files"`
}

func (o *DumpProcessRequest) xxx_ToOp(ctx context.Context, op *xxx_DumpProcessOperation) *xxx_DumpProcessOperation {
	if op == nil {
		op = &xxx_DumpProcessOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.ContainerID = op.ContainerID
	o.Directory = op.Directory
	o.MaxFiles = op.MaxFiles
	return op
}

func (o *DumpProcessRequest) xxx_FromOp(ctx context.Context, op *xxx_DumpProcessOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ContainerID = op.ContainerID
	o.Directory = op.Directory
	o.MaxFiles = op.MaxFiles
}
func (o *DumpProcessRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DumpProcessRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DumpProcessOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DumpProcessResponse structure represents the DumpProcess operation response
type DumpProcessResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrDumpFile: A pointer to a variable that, upon successful completion, contains
	// a fully qualified path, in the convention of the server's file system, to the process
	// dump file written.
	DumpFile *oaut.String `idl:"name:pbstrDumpFile" json:"dump_file"`
	// Return: The DumpProcess return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DumpProcessResponse) xxx_ToOp(ctx context.Context, op *xxx_DumpProcessOperation) *xxx_DumpProcessOperation {
	if op == nil {
		op = &xxx_DumpProcessOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.DumpFile = op.DumpFile
	o.Return = op.Return
	return op
}

func (o *DumpProcessResponse) xxx_FromOp(ctx context.Context, op *xxx_DumpProcessOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DumpFile = op.DumpFile
	o.Return = op.Return
}
func (o *DumpProcessResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DumpProcessResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DumpProcessOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
