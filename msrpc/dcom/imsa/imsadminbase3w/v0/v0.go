package imsadminbase3w

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	imsadminbase2w "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/imsadminbase2w/v0"
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
	_ = imsadminbase2w.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/imsa"
)

var (
	// IMSAdminBase3W interface identifier f612954d-3b0b-4c56-9563-227b7be624b4
	IMSAdminBase3WIID = &dcom.IID{Data1: 0xf612954d, Data2: 0x3b0b, Data3: 0x4c56, Data4: []byte{0x95, 0x63, 0x22, 0x7b, 0x7b, 0xe6, 0x24, 0xb4}}
	// Syntax UUID
	IMSAdminBase3WSyntaxUUID = &uuid.UUID{TimeLow: 0xf612954d, TimeMid: 0x3b0b, TimeHiAndVersion: 0x4c56, ClockSeqHiAndReserved: 0x95, ClockSeqLow: 0x63, Node: [6]uint8{0x22, 0x7b, 0x7b, 0xe6, 0x24, 0xb4}}
	// Syntax ID
	IMSAdminBase3WSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: IMSAdminBase3WSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSAdminBase3W interface.
type IMSAdminBase3WClient interface {

	// IMSAdminBase2W retrieval method.
	IMSAdminBase2W() imsadminbase2w.IMSAdminBase2WClient

	// The GetChildPaths method returns all child nodes of a specified path from a supplied
	// metadata handle.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+--------------------------------------+----------------------------------------------------------+
	//	|                RETURN                |                                                          |
	//	|              VALUE/CODE              |                       DESCRIPTION                        |
	//	|                                      |                                                          |
	//	+--------------------------------------+----------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80070000 S_OK                      | The call was successful.                                 |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND      | The system cannot find the path specified.               |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | One or more arguments are invalid.                       |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED       | Access is denied.                                        |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY   | Not enough storage is available to process this command. |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x8007000E E_OUTOFMEMORY             | There was not enough memory to complete the method call. |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The data area passed to a system call is too small.      |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x800700A0 ERROR_BAD_ARGUMENTS       | One or more arguments are not correct.                   |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80004005 E_FAIL                    | An unspecified error occurred.                           |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80070006 E_HANDLE                  | An invalid handle was passed.                            |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x800CC800 MD_ERROR_NOT_INITIALIZED  | Metadata has not been initialized.                       |
	//	+--------------------------------------+----------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 40.
	GetChildPaths(context.Context, *GetChildPathsRequest, ...dcerpc.CallOption) (*GetChildPathsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) IMSAdminBase3WClient
}

type xxx_DefaultIMSAdminBase3WClient struct {
	imsadminbase2w.IMSAdminBase2WClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultIMSAdminBase3WClient) IMSAdminBase2W() imsadminbase2w.IMSAdminBase2WClient {
	return o.IMSAdminBase2WClient
}

func (o *xxx_DefaultIMSAdminBase3WClient) GetChildPaths(ctx context.Context, in *GetChildPathsRequest, opts ...dcerpc.CallOption) (*GetChildPathsResponse, error) {
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
	out := &GetChildPathsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBase3WClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIMSAdminBase3WClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultIMSAdminBase3WClient) IPID(ctx context.Context, ipid *dcom.IPID) IMSAdminBase3WClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultIMSAdminBase3WClient{
		IMSAdminBase2WClient: o.IMSAdminBase2WClient.IPID(ctx, ipid),
		cc:                   o.cc,
		ipid:                 ipid,
	}
}

func NewIMSAdminBase3WClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IMSAdminBase3WClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IMSAdminBase3WSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := imsadminbase2w.NewIMSAdminBase2WClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultIMSAdminBase3WClient{
		IMSAdminBase2WClient: base,
		cc:                   cc,
		ipid:                 ipid,
	}, nil
}

// xxx_GetChildPathsOperation structure represents the GetChildPaths operation
type xxx_GetChildPathsOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle             uint32         `idl:"name:hMDHandle" json:"handle"`
	Path               string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	BufferSize         uint32         `idl:"name:cchMDBufferSize" json:"buffer_size"`
	Buffer             string         `idl:"name:pszBuffer;size_is:(cchMDBufferSize);pointer:unique" json:"buffer"`
	RequiredBufferSize uint32         `idl:"name:pcchMDRequiredBufferSize;pointer:unique" json:"required_buffer_size"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetChildPathsOperation) OpNum() int { return 40 }

func (o *xxx_GetChildPathsOperation) OpName() string { return "/IMSAdminBase3W/v0/GetChildPaths" }

func (o *xxx_GetChildPathsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != "" && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChildPathsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// cchMDBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// pszBuffer {in, out} (1:{pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,size_is=cchMDBufferSize,string](wchar))
	{
		if o.Buffer != "" || o.BufferSize > 0 {
			_ptr_pszBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				_Buffer_buf := utf16.Encode([]rune(o.Buffer))
				if uint64(len(_Buffer_buf)) > sizeInfo[0] {
					_Buffer_buf = _Buffer_buf[:sizeInfo[0]]
				}
				for i1 := range _Buffer_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_Buffer_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_pszBuffer); err != nil {
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
	// pcchMDRequiredBufferSize {in, out} (1:{pointer=unique}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_pcchMDRequiredBufferSize := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.RequiredBufferSize); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RequiredBufferSize, _ptr_pcchMDRequiredBufferSize); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChildPathsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cchMDBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// pszBuffer {in, out} (1:{pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,size_is=cchMDBufferSize,string](wchar))
	{
		_ptr_pszBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			var _Buffer_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
			}
			_Buffer_buf = make([]uint16, sizeInfo[0])
			for i1 := range _Buffer_buf {
				i1 := i1
				if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
					return err
				}
			}
			o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
			return nil
		})
		_s_pszBuffer := func(ptr interface{}) { o.Buffer = *ptr.(*string) }
		if err := w.ReadPointer(&o.Buffer, _s_pszBuffer, _ptr_pszBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcchMDRequiredBufferSize {in, out} (1:{pointer=unique}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_pcchMDRequiredBufferSize := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.RequiredBufferSize); err != nil {
				return err
			}
			return nil
		})
		_s_pcchMDRequiredBufferSize := func(ptr interface{}) { o.RequiredBufferSize = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.RequiredBufferSize, _s_pcchMDRequiredBufferSize, _ptr_pcchMDRequiredBufferSize); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChildPathsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChildPathsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pszBuffer {in, out} (1:{pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,size_is=cchMDBufferSize,string](wchar))
	{
		if o.Buffer != "" || o.BufferSize > 0 {
			_ptr_pszBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				_Buffer_buf := utf16.Encode([]rune(o.Buffer))
				if uint64(len(_Buffer_buf)) > sizeInfo[0] {
					_Buffer_buf = _Buffer_buf[:sizeInfo[0]]
				}
				for i1 := range _Buffer_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_Buffer_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_pszBuffer); err != nil {
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
	// pcchMDRequiredBufferSize {in, out} (1:{pointer=unique}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_pcchMDRequiredBufferSize := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.RequiredBufferSize); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RequiredBufferSize, _ptr_pcchMDRequiredBufferSize); err != nil {
			return err
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

func (o *xxx_GetChildPathsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pszBuffer {in, out} (1:{pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,size_is=cchMDBufferSize,string](wchar))
	{
		_ptr_pszBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			var _Buffer_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
			}
			_Buffer_buf = make([]uint16, sizeInfo[0])
			for i1 := range _Buffer_buf {
				i1 := i1
				if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
					return err
				}
			}
			o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
			return nil
		})
		_s_pszBuffer := func(ptr interface{}) { o.Buffer = *ptr.(*string) }
		if err := w.ReadPointer(&o.Buffer, _s_pszBuffer, _ptr_pszBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcchMDRequiredBufferSize {in, out} (1:{pointer=unique}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_pcchMDRequiredBufferSize := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.RequiredBufferSize); err != nil {
				return err
			}
			return nil
		})
		_s_pcchMDRequiredBufferSize := func(ptr interface{}) { o.RequiredBufferSize = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.RequiredBufferSize, _s_pcchMDRequiredBufferSize, _ptr_pcchMDRequiredBufferSize); err != nil {
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

// GetChildPathsRequest structure represents the GetChildPaths operation request
type GetChildPathsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	Handle uint32         `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node to be
	// opened, relative to the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// cchMDBufferSize: The size, in WCHAR, of the pszBuffer buffer to hold the paths for
	// all child nodes under the path specified.
	BufferSize uint32 `idl:"name:cchMDBufferSize" json:"buffer_size"`
	// pszBuffer: A pointer to a Unicode character buffer passed in by the caller to store
	// the retrieved child paths. The return data will be a set of WCHAR strings, where
	// each includes two terminating null characters.
	Buffer string `idl:"name:pszBuffer;size_is:(cchMDBufferSize);pointer:unique" json:"buffer"`
	// pcchMDRequiredBufferSize: An integer value indicating the required size of the buffer
	// if the supplied buffer proves to be insufficient. If the supplied buffer is sufficient,
	// this value will not be adjusted.
	RequiredBufferSize uint32 `idl:"name:pcchMDRequiredBufferSize;pointer:unique" json:"required_buffer_size"`
}

func (o *GetChildPathsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetChildPathsOperation) *xxx_GetChildPathsOperation {
	if op == nil {
		op = &xxx_GetChildPathsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.BufferSize = o.BufferSize
	op.Buffer = o.Buffer
	op.RequiredBufferSize = o.RequiredBufferSize
	return op
}

func (o *GetChildPathsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetChildPathsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.BufferSize = op.BufferSize
	o.Buffer = op.Buffer
	o.RequiredBufferSize = op.RequiredBufferSize
}
func (o *GetChildPathsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetChildPathsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetChildPathsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetChildPathsResponse structure represents the GetChildPaths operation response
type GetChildPathsResponse struct {
	// XXX: cchMDBufferSize is an implicit input depedency for output parameters
	BufferSize uint32 `idl:"name:cchMDBufferSize" json:"buffer_size"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pszBuffer: A pointer to a Unicode character buffer passed in by the caller to store
	// the retrieved child paths. The return data will be a set of WCHAR strings, where
	// each includes two terminating null characters.
	Buffer string `idl:"name:pszBuffer;size_is:(cchMDBufferSize);pointer:unique" json:"buffer"`
	// pcchMDRequiredBufferSize: An integer value indicating the required size of the buffer
	// if the supplied buffer proves to be insufficient. If the supplied buffer is sufficient,
	// this value will not be adjusted.
	RequiredBufferSize uint32 `idl:"name:pcchMDRequiredBufferSize;pointer:unique" json:"required_buffer_size"`
	// Return: The GetChildPaths return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetChildPathsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetChildPathsOperation) *xxx_GetChildPathsOperation {
	if op == nil {
		op = &xxx_GetChildPathsOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferSize == uint32(0) {
		op.BufferSize = o.BufferSize
	}

	op.That = o.That
	op.Buffer = o.Buffer
	op.RequiredBufferSize = o.RequiredBufferSize
	op.Return = o.Return
	return op
}

func (o *GetChildPathsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetChildPathsOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferSize = op.BufferSize

	o.That = op.That
	o.Buffer = op.Buffer
	o.RequiredBufferSize = op.RequiredBufferSize
	o.Return = op.Return
}
func (o *GetChildPathsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetChildPathsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetChildPathsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
