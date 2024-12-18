package browser

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "brwsa"
)

var (
	// Syntax UUID
	BrowserSyntaxUUID = &uuid.UUID{TimeLow: 0x6bffd098, TimeMid: 0xa112, TimeHiAndVersion: 0x3610, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x33, Node: [6]uint8{0x1, 0x28, 0x92, 0x2, 0x1, 0x62}}
	// Syntax ID
	BrowserSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: BrowserSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// browser interface.
type BrowserClient interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// The I_BrowserrQueryOtherDomains method is received by the server in an RPC_REQUEST
	// packet. The client SHOULD NOT send this RPC request to a server that is not a primary
	// domain controller (PDC) acting as the domain master browser server.
	//
	// If this server is not a primary domain controller it MAY fail the request.<3>
	//
	// If the server is a primary domain controller, the server MUST update OtherDomains
	// as specified in [MS-WKST] section 3.2.6.1, WkstaQueryOtherDomains Event. The server
	// MUST construct a SERVER_ENUM structure as specified in 2.2.3.2, containing a SERVER_INFO_100
	// structure as specified in [MS-DTYP] section 2.3.11 for each name in OtherDomains,
	// and return this to the caller.
	//
	// Return Values: The method returns NERR_Success on success; otherwise, it returns
	// a nonzero error code, as specified in either Win32 Error Codes. The most common error
	// codes are listed in the following table.<5>
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The operation completed successfully.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | This value MUST be returned when the server could not allocate enough memory to  |
	//	|                                    | complete this operation.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | This value MUST be returned when a parameter is incorrect. For example, this     |
	//	|                                    | value is returned when the InfoStruct parameter is NULL or the Level100 member   |
	//	|                                    | in the structure pointed to by the InfoStruct parameter is NULL.                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | This value MUST be returned when the Level member is not 100.                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA         | The error ERROR_MORE_DATA indicates that not all available entries were          |
	//	|                                    | returned. Some more entries exist which were not returned in the response.       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	QueryOtherDomains(context.Context, *QueryOtherDomainsRequest, ...dcerpc.CallOption) (*QueryOtherDomainsResponse, error)

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// ServerInfo100Container structure represents SERVER_INFO_100_CONTAINER RPC structure.
//
// The SERVER_INFO_100_CONTAINER structure contains a count of the entries returned
// by the method and a pointer to a buffer.
type ServerInfo100Container struct {
	// EntriesRead:  The number of entries returned by the method call. This value MUST
	// be zero if no domains are configured in the primary domain controller or domain controller.
	// The client SHOULD set the EntriesRead field to 0, and the Buffer field to NULL, and
	// the server MUST ignore these fields.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  A pointer to an array of SERVER_INFO_100 data structures (as specified in
	// [MS-DTYP] section 2.3.11). If EntriesRead is zero, this field is undefined and MUST
	// NOT be considered a valid pointer.
	Buffer []*dtyp.ServerInfo100 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *ServerInfo100Container) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerInfo100Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.ServerInfo100{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&dtyp.ServerInfo100{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerInfo100Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*dtyp.ServerInfo100, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &dtyp.ServerInfo100{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*dtyp.ServerInfo100) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// ServerEnum structure represents SERVER_ENUM_STRUCT RPC structure.
//
// The SERVER_ENUM_STRUCT structure defines the layout for a structure with a value
// to indicate the information level submitted to the method and a pointer to a data
// structure that contains an array of data structures returned by the method. This
// structure is used by I_BrowserrQueryOtherDomains.
type ServerEnum struct {
	// Level:  The information level of the data. This member MUST be 100.
	Level uint32 `idl:"name:Level" json:"level"`
	// ServerInfo:  A structure that contains an array of data structures. The Level member
	// determines the data type of the members of this array.
	ServerInfo *ServerEnum_ServerInfo `idl:"name:ServerInfo;switch_is:Level" json:"server_info"`
}

func (o *ServerEnum) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	_swServerInfo := uint32(o.Level)
	if o.ServerInfo != nil {
		if err := o.ServerInfo.MarshalUnionNDR(ctx, w, _swServerInfo); err != nil {
			return err
		}
	} else {
		if err := (&ServerEnum_ServerInfo{}).MarshalUnionNDR(ctx, w, _swServerInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if o.ServerInfo == nil {
		o.ServerInfo = &ServerEnum_ServerInfo{}
	}
	_swServerInfo := uint32(o.Level)
	if err := o.ServerInfo.UnmarshalUnionNDR(ctx, w, _swServerInfo); err != nil {
		return err
	}
	return nil
}

// ServerEnum_ServerInfo structure represents SERVER_ENUM_STRUCT union anonymous member.
//
// The SERVER_ENUM_STRUCT structure defines the layout for a structure with a value
// to indicate the information level submitted to the method and a pointer to a data
// structure that contains an array of data structures returned by the method. This
// structure is used by I_BrowserrQueryOtherDomains.
type ServerEnum_ServerInfo struct {
	// Types that are assignable to Value
	//
	// *ServerInfo_Level100
	Value is_ServerEnum_ServerInfo `json:"value"`
}

func (o *ServerEnum_ServerInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ServerInfo_Level100:
		if value != nil {
			return value.Level100
		}
	}
	return nil
}

type is_ServerEnum_ServerInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ServerEnum_ServerInfo()
}

func (o *ServerEnum_ServerInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ServerInfo_Level100:
		return uint32(100)
	}
	return uint32(0)
}

func (o *ServerEnum_ServerInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(100):
		_o, _ := o.Value.(*ServerInfo_Level100)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerInfo_Level100{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *ServerEnum_ServerInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(100):
		o.Value = &ServerInfo_Level100{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// ServerInfo_Level100 structure represents ServerEnum_ServerInfo RPC union arm.
//
// It has following labels: 100
type ServerInfo_Level100 struct {
	// Level100:  A pointer to a SERVER_INFO_100_CONTAINER structure that contains the number
	// of entries returned by the method and a pointer to an array of SERVER_INFO_100 structures
	// (as specified in [MS-DTYP] section 2.3.11).
	Level100 *ServerInfo100Container `idl:"name:Level100" json:"level100"`
}

func (*ServerInfo_Level100) is_ServerEnum_ServerInfo() {}

func (o *ServerInfo_Level100) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Level100 != nil {
		_ptr_Level100 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Level100 != nil {
				if err := o.Level100.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ServerInfo100Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Level100, _ptr_Level100); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerInfo_Level100) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Level100 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Level100 == nil {
			o.Level100 = &ServerInfo100Container{}
		}
		if err := o.Level100.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Level100 := func(ptr interface{}) { o.Level100 = *ptr.(**ServerInfo100Container) }
	if err := w.ReadPointer(&o.Level100, _s_Level100, _ptr_Level100); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultBrowserClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultBrowserClient) QueryOtherDomains(ctx context.Context, in *QueryOtherDomainsRequest, opts ...dcerpc.CallOption) (*QueryOtherDomainsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryOtherDomainsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBrowserClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultBrowserClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewBrowserClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (BrowserClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(BrowserSyntaxV0_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultBrowserClient{cc: cc}, nil
}

// xxx_QueryOtherDomainsOperation structure represents the I_BrowserrQueryOtherDomains operation
type xxx_QueryOtherDomainsOperation struct {
	ServerName   string      `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	Info         *ServerEnum `idl:"name:InfoStruct" json:"info"`
	TotalEntries uint32      `idl:"name:TotalEntries" json:"total_entries"`
	Return       uint32      `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryOtherDomainsOperation) OpNum() int { return 2 }

func (o *xxx_QueryOtherDomainsOperation) OpName() string {
	return "/browser/v0/I_BrowserrQueryOtherDomains"
}

func (o *xxx_QueryOtherDomainsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOtherDomainsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=BROWSER_IDENTIFY_HANDLE, names=LPWSTR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// InfoStruct {in, out} (1:{alias=LPSERVER_ENUM_STRUCT}*(1))(2:{alias=SERVER_ENUM_STRUCT}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOtherDomainsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=BROWSER_IDENTIFY_HANDLE, names=LPWSTR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// InfoStruct {in, out} (1:{alias=LPSERVER_ENUM_STRUCT,pointer=ref}*(1))(2:{alias=SERVER_ENUM_STRUCT}(struct))
	{
		if o.Info == nil {
			o.Info = &ServerEnum{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOtherDomainsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOtherDomainsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// InfoStruct {in, out} (1:{alias=LPSERVER_ENUM_STRUCT}*(1))(2:{alias=SERVER_ENUM_STRUCT}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TotalEntries {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOtherDomainsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// InfoStruct {in, out} (1:{alias=LPSERVER_ENUM_STRUCT,pointer=ref}*(1))(2:{alias=SERVER_ENUM_STRUCT}(struct))
	{
		if o.Info == nil {
			o.Info = &ServerEnum{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TotalEntries {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryOtherDomainsRequest structure represents the I_BrowserrQueryOtherDomains operation request
type QueryOtherDomainsRequest struct {
	// ServerName: An optional BROWSER_IDENTIFY_HANDLE structure that specifies the name
	// of the server to execute the method. This value is ignored upon receipt.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// InfoStruct: A pointer to a SERVER_ENUM_STRUCT structure that contains the Level member
	// and a pointer to a SERVER_INFO_x structure, where <x> MUST be 100. The Level member
	// MUST be set to 100. If the Level member is set to any other value, the method MUST
	// return ERROR_INVALID_LEVEL.<4>
	Info *ServerEnum `idl:"name:InfoStruct" json:"info"`
}

func (o *QueryOtherDomainsRequest) xxx_ToOp(ctx context.Context) *xxx_QueryOtherDomainsOperation {
	if o == nil {
		return &xxx_QueryOtherDomainsOperation{}
	}
	return &xxx_QueryOtherDomainsOperation{
		ServerName: o.ServerName,
		Info:       o.Info,
	}
}

func (o *QueryOtherDomainsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryOtherDomainsOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Info = op.Info
}
func (o *QueryOtherDomainsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryOtherDomainsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryOtherDomainsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryOtherDomainsResponse structure represents the I_BrowserrQueryOtherDomains operation response
type QueryOtherDomainsResponse struct {
	// InfoStruct: A pointer to a SERVER_ENUM_STRUCT structure that contains the Level member
	// and a pointer to a SERVER_INFO_x structure, where <x> MUST be 100. The Level member
	// MUST be set to 100. If the Level member is set to any other value, the method MUST
	// return ERROR_INVALID_LEVEL.<4>
	Info *ServerEnum `idl:"name:InfoStruct" json:"info"`
	// TotalEntries: The number of entries returned by the method call. This parameter MUST
	// match the EntriesRead member of the SERVER_INFO_100_CONTAINER structure.
	TotalEntries uint32 `idl:"name:TotalEntries" json:"total_entries"`
	// Return: The I_BrowserrQueryOtherDomains return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryOtherDomainsResponse) xxx_ToOp(ctx context.Context) *xxx_QueryOtherDomainsOperation {
	if o == nil {
		return &xxx_QueryOtherDomainsOperation{}
	}
	return &xxx_QueryOtherDomainsOperation{
		Info:         o.Info,
		TotalEntries: o.TotalEntries,
		Return:       o.Return,
	}
}

func (o *QueryOtherDomainsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryOtherDomainsOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.TotalEntries = op.TotalEntries
	o.Return = op.Return
}
func (o *QueryOtherDomainsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryOtherDomainsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryOtherDomainsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
