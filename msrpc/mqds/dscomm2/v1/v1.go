package dscomm2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	dscomm "github.com/oiweiwei/go-msrpc/msrpc/mqds/dscomm/v1"
	mqmq "github.com/oiweiwei/go-msrpc/msrpc/mqmq"
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
	_ = dscomm.GoPackage
	_ = mqmq.GoPackage
)

var (
	// import guard
	GoPackage = "mqds"
)

var (
	// Syntax UUID
	Dscomm2SyntaxUUID = &uuid.UUID{TimeLow: 0x708cca10, TimeMid: 0x9569, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xb2, ClockSeqLow: 0xa5, Node: [6]uint8{0x0, 0x60, 0x97, 0x7d, 0x81, 0x18}}
	// Syntax ID
	Dscomm2SyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: Dscomm2SyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// dscomm2 interface.
type Dscomm2Client interface {

	// The S_DSGetComputerSites method returns the site identifier for every site of which
	// the specified computer is a member.
	//
	// Return Values:  If the method succeeds, the return value is MQ_OK (0x00000000).
	// If the method fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetComputerSites(context.Context, *GetComputerSitesRequest, ...dcerpc.CallOption) (*GetComputerSitesResponse, error)

	// The S_DSGetPropsEx method returns the properties associated with the object specified
	// by a directory service pathname. This method differs from S_DSGetProps (section 3.1.4.7)
	// in that it supports a restricted set of properties that pertain only to queue or
	// machine object security.
	//
	// Return Values:  If the method succeeds, the return value is MQ_OK (0x00000000).
	// If the method fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetPropertiesEx(context.Context, *GetPropertiesExRequest, ...dcerpc.CallOption) (*GetPropertiesExResponse, error)

	// This method returns the properties for the object specified by object identifier.
	// This method differs from S_DSGetPropsGuid in that it supports a restricted set of
	// properties that pertain only to queue or machine object security.
	//
	// Return Values:  If the method succeeds, the return value is MQ_OK (0x00000000).
	// If the method fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	GetPropertiesGUIDEx(context.Context, *GetPropertiesGUIDExRequest, ...dcerpc.CallOption) (*GetPropertiesGUIDExResponse, error)

	// The S_DSBeginDeleteNotification method begins a delete notification and returns an
	// RPC context handle associated with the delete notification.
	//
	// Return Values:  If the method succeeds, the return value is MQ_OK (0x00000000).
	// If the method fails, the return value is an implementation-specific error code.<158>
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	BeginDeleteNotification(context.Context, *BeginDeleteNotificationRequest, ...dcerpc.CallOption) (*BeginDeleteNotificationResponse, error)

	// This method instructs the server to notify the computer that owns the deleted object
	// about the deletion.
	//
	// Return Values:  If the method succeeds, the return value is MQ_OK (0x00000000).
	// If the method fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	NotifyDelete(context.Context, *NotifyDeleteRequest, ...dcerpc.CallOption) (*NotifyDeleteResponse, error)

	// The S_DSEndDeleteNotification method closes the RPC context handle acquired from
	// a previous call to S_DSBeginDeleteNotification.
	//
	// Return Values:  None.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	EndDeleteNotification(context.Context, *EndDeleteNotificationRequest, ...dcerpc.CallOption) (*EndDeleteNotificationResponse, error)

	// This method returns a value that indicates if that server is a Global Catalog Server.
	//
	// Return Values:  This method returns TRUE (0x00000001) if the Directory Service server
	// is also a Global Catalog Server; otherwise, it returns FALSE (0x00000000).
	//
	// TRUE (0x00000001)
	//
	// FALSE (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	IsServerGC(context.Context, *IsServerGCRequest, ...dcerpc.CallOption) (*IsServerGCResponse, error)

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// The S_DSGetGCListInDomain method returns the list of Global Catalog Servers in the
	// specified domain.
	//
	// Return Values:  If the method succeeds, the return value is MQ_OK (0x00000000).
	// If the method fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, Remote Procedure Call Protocol Extensions, as specified in [MS-RPCE].
	GetGCListInDomain(context.Context, *GetGCListInDomainRequest, ...dcerpc.CallOption) (*GetGCListInDomainResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultDscomm2Client struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultDscomm2Client) GetComputerSites(ctx context.Context, in *GetComputerSitesRequest, opts ...dcerpc.CallOption) (*GetComputerSitesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetComputerSitesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscomm2Client) GetPropertiesEx(ctx context.Context, in *GetPropertiesExRequest, opts ...dcerpc.CallOption) (*GetPropertiesExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPropertiesExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscomm2Client) GetPropertiesGUIDEx(ctx context.Context, in *GetPropertiesGUIDExRequest, opts ...dcerpc.CallOption) (*GetPropertiesGUIDExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPropertiesGUIDExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscomm2Client) BeginDeleteNotification(ctx context.Context, in *BeginDeleteNotificationRequest, opts ...dcerpc.CallOption) (*BeginDeleteNotificationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BeginDeleteNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscomm2Client) NotifyDelete(ctx context.Context, in *NotifyDeleteRequest, opts ...dcerpc.CallOption) (*NotifyDeleteResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NotifyDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscomm2Client) EndDeleteNotification(ctx context.Context, in *EndDeleteNotificationRequest, opts ...dcerpc.CallOption) (*EndDeleteNotificationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EndDeleteNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultDscomm2Client) IsServerGC(ctx context.Context, in *IsServerGCRequest, opts ...dcerpc.CallOption) (*IsServerGCResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IsServerGCResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscomm2Client) GetGCListInDomain(ctx context.Context, in *GetGCListInDomainRequest, opts ...dcerpc.CallOption) (*GetGCListInDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetGCListInDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscomm2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDscomm2Client) Conn() dcerpc.Conn {
	return o.cc
}

func NewDscomm2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Dscomm2Client, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Dscomm2SyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultDscomm2Client{cc: cc}, nil
}

// xxx_GetComputerSitesOperation structure represents the S_DSGetComputerSites operation
type xxx_GetComputerSitesOperation struct {
	PathName            string                 `idl:"name:pwcsPathName;string;pointer:unique" json:"path_name"`
	NumberOfSites       uint32                 `idl:"name:pdwNumberOfSites" json:"number_of_sites"`
	Sites               []*dtyp.GUID           `idl:"name:ppguidSites;size_is:(, pdwNumberOfSites);length_is:(, pdwNumberOfSites)" json:"sites"`
	ServerAuth          *dscomm.ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	ServerSignature     []byte                 `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	ServerSignatureSize uint32                 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	Return              int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetComputerSitesOperation) OpNum() int { return 0 }

func (o *xxx_GetComputerSitesOperation) OpName() string { return "/dscomm2/v1/S_DSGetComputerSites" }

func (o *xxx_GetComputerSitesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComputerSitesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pwcsPathName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.PathName != "" {
			_ptr_pwcsPathName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.PathName, _ptr_pwcsPathName); err != nil {
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
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dscomm.ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComputerSitesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pwcsPathName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwcsPathName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
				return err
			}
			return nil
		})
		_s_pwcsPathName := func(ptr interface{}) { o.PathName = *ptr.(*string) }
		if err := w.ReadPointer(&o.PathName, _s_pwcsPathName, _ptr_pwcsPathName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &dscomm.ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComputerSitesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Sites != nil && o.NumberOfSites == 0 {
		o.NumberOfSites = uint32(len(o.Sites))
	}
	if o.Sites != nil && o.NumberOfSites == 0 {
		o.NumberOfSites = uint32(len(o.Sites))
	}
	if o.ServerSignature != nil && o.ServerSignatureSize == 0 {
		o.ServerSignatureSize = uint32(len(o.ServerSignature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComputerSitesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwNumberOfSites {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NumberOfSites); err != nil {
			return err
		}
	}
	// ppguidSites {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pdwNumberOfSites,length_is=pdwNumberOfSites](struct))
	{
		if o.Sites != nil || o.NumberOfSites > 0 {
			_ptr_ppguidSites := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfSites)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				dimLength1 := uint64(o.NumberOfSites)
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
				for i1 := range o.Sites {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Sites[i1] != nil {
						if err := o.Sites[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Sites); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Sites, _ptr_ppguidSites); err != nil {
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
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
	{
		dimSize1 := uint64(o.ServerSignatureSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ServerSignature {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ServerSignature[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerSignature); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
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

func (o *xxx_GetComputerSitesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwNumberOfSites {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NumberOfSites); err != nil {
			return err
		}
	}
	// ppguidSites {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pdwNumberOfSites,length_is=pdwNumberOfSites](struct))
	{
		_ptr_ppguidSites := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
				return fmt.Errorf("buffer overflow for size %d of array o.Sites", sizeInfo[0])
			}
			o.Sites = make([]*dtyp.GUID, sizeInfo[0])
			for i1 := range o.Sites {
				i1 := i1
				if o.Sites[i1] == nil {
					o.Sites[i1] = &dtyp.GUID{}
				}
				if err := o.Sites[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppguidSites := func(ptr interface{}) { o.Sites = *ptr.(*[]*dtyp.GUID) }
		if err := w.ReadPointer(&o.Sites, _s_ppguidSites, _ptr_ppguidSites); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ServerSignature", sizeInfo[0])
		}
		o.ServerSignature = make([]byte, sizeInfo[0])
		for i1 := range o.ServerSignature {
			i1 := i1
			if err := w.ReadData(&o.ServerSignature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
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

// GetComputerSitesRequest structure represents the S_DSGetComputerSites operation request
type GetComputerSitesRequest struct {
	// pwcsPathName:  MUST be set by the client to a pointer to a NULL-terminated 16-bit
	// Unicode string that contains the directory service pathname, as specified in section
	// 2.2.9, of the object in the Directory Service. The pathname MUST be the pathname
	// of an object of type MQDS_MACHINE.<154>
	PathName string `idl:"name:pwcsPathName;string;pointer:unique" json:"path_name"`
	// phServerAuth:  A PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle acquired from
	// the pphServerAuth parameter in a previous call to S_DSValidateServer. The server
	// MUST use this parameter as a key to locate the GSS security context used to compute
	// the signature returned in pbServerSignature. See section 3.1.4.2.
	ServerAuth *dscomm.ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	// pdwServerSignatureSize: Contains the maximum length of the server signature in bytes
	// to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
}

func (o *GetComputerSitesRequest) xxx_ToOp(ctx context.Context) *xxx_GetComputerSitesOperation {
	if o == nil {
		return &xxx_GetComputerSitesOperation{}
	}
	return &xxx_GetComputerSitesOperation{
		PathName:            o.PathName,
		ServerAuth:          o.ServerAuth,
		ServerSignatureSize: o.ServerSignatureSize,
	}
}

func (o *GetComputerSitesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetComputerSitesOperation) {
	if o == nil {
		return
	}
	o.PathName = op.PathName
	o.ServerAuth = op.ServerAuth
	o.ServerSignatureSize = op.ServerSignatureSize
}
func (o *GetComputerSitesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetComputerSitesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetComputerSitesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetComputerSitesResponse structure represents the S_DSGetComputerSites operation response
type GetComputerSitesResponse struct {
	// pdwNumberOfSites: Number of site identifiers returned in the ppguidSites parameter.
	NumberOfSites uint32 `idl:"name:pdwNumberOfSites" json:"number_of_sites"`
	// ppguidSites: An array of pointers to the GUID site identifiers of the sites of which
	// the computer is a member. Each GUID referenced by the array MUST be the site identifier
	// for a site that the machine object specified by pwcsPathName is a member of. These
	// are obtained from the PROPID_QM_SITE_IDS property of the MQDS_MACHINE object. The
	// array MUST contain the site identifier for every site to which the object specified
	// by pwcsPathName is a member.
	Sites []*dtyp.GUID `idl:"name:ppguidSites;size_is:(, pdwNumberOfSites);length_is:(, pdwNumberOfSites)" json:"sites"`
	// pbServerSignature: A buffer that contains a signed hash over the array of site GUIDs
	// returned in ppguidSites.
	ServerSignature []byte `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	// pdwServerSignatureSize: Contains the maximum length of the server signature in bytes
	// to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	// Return: The S_DSGetComputerSites return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetComputerSitesResponse) xxx_ToOp(ctx context.Context) *xxx_GetComputerSitesOperation {
	if o == nil {
		return &xxx_GetComputerSitesOperation{}
	}
	return &xxx_GetComputerSitesOperation{
		NumberOfSites:       o.NumberOfSites,
		Sites:               o.Sites,
		ServerSignature:     o.ServerSignature,
		ServerSignatureSize: o.ServerSignatureSize,
		Return:              o.Return,
	}
}

func (o *GetComputerSitesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetComputerSitesOperation) {
	if o == nil {
		return
	}
	o.NumberOfSites = op.NumberOfSites
	o.Sites = op.Sites
	o.ServerSignature = op.ServerSignature
	o.ServerSignatureSize = op.ServerSignatureSize
	o.Return = op.Return
}
func (o *GetComputerSitesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetComputerSitesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetComputerSitesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertiesExOperation structure represents the S_DSGetPropsEx operation
type xxx_GetPropertiesExOperation struct {
	ObjectType          uint32                  `idl:"name:dwObjectType" json:"object_type"`
	PathName            string                  `idl:"name:pwcsPathName;string" json:"path_name"`
	CreatePartition     uint32                  `idl:"name:cp" json:"create_partition"`
	Property            []uint32                `idl:"name:aProp;size_is:(cp)" json:"property"`
	Var                 []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	ServerAuth          *dscomm.ServerAuthType  `idl:"name:phServerAuth" json:"server_auth"`
	ServerSignature     []byte                  `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	ServerSignatureSize uint32                  `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	Return              int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesExOperation) OpNum() int { return 1 }

func (o *xxx_GetPropertiesExOperation) OpName() string { return "/dscomm2/v1/S_DSGetPropsEx" }

func (o *xxx_GetPropertiesExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Property != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Property))
	}
	if o.Var != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Var))
	}
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
	}
	if o.CreatePartition < uint32(1) || o.CreatePartition > uint32(128) {
		return fmt.Errorf("CreatePartition is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp])(2:{alias=PROPID}(uint32))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Property {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Property[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Property); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		}
	}
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Var {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Var[i1] != nil {
				if err := o.Var[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Var); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dscomm.ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp])(2:{alias=PROPID}(uint32))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Property", sizeInfo[0])
		}
		o.Property = make([]uint32, sizeInfo[0])
		for i1 := range o.Property {
			i1 := i1
			if err := w.ReadData(&o.Property[i1]); err != nil {
				return err
			}
		}
	}
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Var", sizeInfo[0])
		}
		o.Var = make([]*mqmq.PropertyVariant, sizeInfo[0])
		for i1 := range o.Var {
			i1 := i1
			if o.Var[i1] == nil {
				o.Var[i1] = &mqmq.PropertyVariant{}
			}
			if err := o.Var[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &dscomm.ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ServerSignature != nil && o.ServerSignatureSize == 0 {
		o.ServerSignatureSize = uint32(len(o.ServerSignature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Var {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Var[i1] != nil {
				if err := o.Var[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Var); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
	{
		dimSize1 := uint64(o.ServerSignatureSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ServerSignature {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ServerSignature[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerSignature); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
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

func (o *xxx_GetPropertiesExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Var", sizeInfo[0])
		}
		o.Var = make([]*mqmq.PropertyVariant, sizeInfo[0])
		for i1 := range o.Var {
			i1 := i1
			if o.Var[i1] == nil {
				o.Var[i1] = &mqmq.PropertyVariant{}
			}
			if err := o.Var[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ServerSignature", sizeInfo[0])
		}
		o.ServerSignature = make([]byte, sizeInfo[0])
		for i1 := range o.ServerSignature {
			i1 := i1
			if err := w.ReadData(&o.ServerSignature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
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

// GetPropertiesExRequest structure represents the S_DSGetPropsEx operation request
type GetPropertiesExRequest struct {
	// dwObjectType:  Specifies the type of object for which properties are to be retrieved.
	// MUST be set to one of the object types, as specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pwcsPathName: MUST be set by the client to a pointer to a NULL-terminated 16-bit
	// Unicode string that contains the directory service pathname, as specified in section
	// 2.2.9, of the object in the Directory Service from which to retrieve the properties.
	PathName string `idl:"name:pwcsPathName;string" json:"path_name"`
	// cp:  MUST be set to the size (in elements) of the arrays aProp and apVar, which
	// for this method MUST be one (0x00000001). The arrays aProp and apVar MUST have an
	// identical number of elements, and MUST each contain exactly one element.
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// aProp:  An array of identifiers of properties to retrieve from the object. Each
	// element MUST specify a value from the property identifiers table for the object type
	// specified in dwObjectType. Each element MUST specify the property identifier for
	// the corresponding property value at the same element index in apVar. The array MUST
	// contain exactly one element.
	Property []uint32 `idl:"name:aProp;size_is:(cp)" json:"property"`
	// apVar:  MUST be set by the client to an array that holds the property values retrieved
	// from the object. Each element MUST be set by the server to the property value for
	// the corresponding property identifier at the same element index in aProp. The array
	// MUST contain exactly one element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	// phServerAuth:  A PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle acquired from
	// the pphServerAuth parameter in a previous call to S_DSValidateServer (section 3.1.4.2).
	// The server MUST use this parameter as a key to locate the GSS security context used
	// to compute the signature returned in pbServerSignature. See section 3.1.4.2.
	ServerAuth *dscomm.ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	// pdwServerSignatureSize: A DWORD that contains the maximum length of the server signature
	// in bytes to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
}

func (o *GetPropertiesExRequest) xxx_ToOp(ctx context.Context) *xxx_GetPropertiesExOperation {
	if o == nil {
		return &xxx_GetPropertiesExOperation{}
	}
	return &xxx_GetPropertiesExOperation{
		ObjectType:          o.ObjectType,
		PathName:            o.PathName,
		CreatePartition:     o.CreatePartition,
		Property:            o.Property,
		Var:                 o.Var,
		ServerAuth:          o.ServerAuth,
		ServerSignatureSize: o.ServerSignatureSize,
	}
}

func (o *GetPropertiesExRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesExOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.PathName = op.PathName
	o.CreatePartition = op.CreatePartition
	o.Property = op.Property
	o.Var = op.Var
	o.ServerAuth = op.ServerAuth
	o.ServerSignatureSize = op.ServerSignatureSize
}
func (o *GetPropertiesExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesExResponse structure represents the S_DSGetPropsEx operation response
type GetPropertiesExResponse struct {
	// apVar:  MUST be set by the client to an array that holds the property values retrieved
	// from the object. Each element MUST be set by the server to the property value for
	// the corresponding property identifier at the same element index in aProp. The array
	// MUST contain exactly one element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	// pbServerSignature: A buffer that contains a signed hash over the returned property
	// values. See the pbServerSignature parameter description in section 3.1.4.7.
	ServerSignature []byte `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	// pdwServerSignatureSize: A DWORD that contains the maximum length of the server signature
	// in bytes to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	// Return: The S_DSGetPropsEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesExResponse) xxx_ToOp(ctx context.Context) *xxx_GetPropertiesExOperation {
	if o == nil {
		return &xxx_GetPropertiesExOperation{}
	}
	return &xxx_GetPropertiesExOperation{
		Var:                 o.Var,
		ServerSignature:     o.ServerSignature,
		ServerSignatureSize: o.ServerSignatureSize,
		Return:              o.Return,
	}
}

func (o *GetPropertiesExResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesExOperation) {
	if o == nil {
		return
	}
	o.Var = op.Var
	o.ServerSignature = op.ServerSignature
	o.ServerSignatureSize = op.ServerSignatureSize
	o.Return = op.Return
}
func (o *GetPropertiesExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertiesGUIDExOperation structure represents the S_DSGetPropsGuidEx operation
type xxx_GetPropertiesGUIDExOperation struct {
	ObjectType          uint32                  `idl:"name:dwObjectType" json:"object_type"`
	GUID                *dtyp.GUID              `idl:"name:pGuid;pointer:unique" json:"guid"`
	CreatePartition     uint32                  `idl:"name:cp" json:"create_partition"`
	Property            []uint32                `idl:"name:aProp;size_is:(cp)" json:"property"`
	Var                 []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	ServerAuth          *dscomm.ServerAuthType  `idl:"name:phServerAuth" json:"server_auth"`
	ServerSignature     []byte                  `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	ServerSignatureSize uint32                  `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	Return              int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesGUIDExOperation) OpNum() int { return 2 }

func (o *xxx_GetPropertiesGUIDExOperation) OpName() string { return "/dscomm2/v1/S_DSGetPropsGuidEx" }

func (o *xxx_GetPropertiesGUIDExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Property != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Property))
	}
	if o.Var != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Var))
	}
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
	}
	if o.CreatePartition < uint32(1) || o.CreatePartition > uint32(128) {
		return fmt.Errorf("CreatePartition is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesGUIDExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			_ptr_pGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.GUID != nil {
					if err := o.GUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.GUID, _ptr_pGuid); err != nil {
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
	// cp {in} (1:{range=(1,128), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp])(2:{alias=PROPID}(uint32))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Property {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Property[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Property); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		}
	}
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Var {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Var[i1] != nil {
				if err := o.Var[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Var); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dscomm.ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesGUIDExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_pGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.GUID == nil {
				o.GUID = &dtyp.GUID{}
			}
			if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pGuid := func(ptr interface{}) { o.GUID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.GUID, _s_pGuid, _ptr_pGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp])(2:{alias=PROPID}(uint32))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Property", sizeInfo[0])
		}
		o.Property = make([]uint32, sizeInfo[0])
		for i1 := range o.Property {
			i1 := i1
			if err := w.ReadData(&o.Property[i1]); err != nil {
				return err
			}
		}
	}
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Var", sizeInfo[0])
		}
		o.Var = make([]*mqmq.PropertyVariant, sizeInfo[0])
		for i1 := range o.Var {
			i1 := i1
			if o.Var[i1] == nil {
				o.Var[i1] = &mqmq.PropertyVariant{}
			}
			if err := o.Var[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &dscomm.ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesGUIDExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ServerSignature != nil && o.ServerSignatureSize == 0 {
		o.ServerSignatureSize = uint32(len(o.ServerSignature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesGUIDExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Var {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Var[i1] != nil {
				if err := o.Var[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Var); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
	{
		dimSize1 := uint64(o.ServerSignatureSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ServerSignature {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ServerSignature[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerSignature); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
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

func (o *xxx_GetPropertiesGUIDExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Var", sizeInfo[0])
		}
		o.Var = make([]*mqmq.PropertyVariant, sizeInfo[0])
		for i1 := range o.Var {
			i1 := i1
			if o.Var[i1] == nil {
				o.Var[i1] = &mqmq.PropertyVariant{}
			}
			if err := o.Var[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ServerSignature", sizeInfo[0])
		}
		o.ServerSignature = make([]byte, sizeInfo[0])
		for i1 := range o.ServerSignature {
			i1 := i1
			if err := w.ReadData(&o.ServerSignature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
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

// GetPropertiesGUIDExRequest structure represents the S_DSGetPropsGuidEx operation request
type GetPropertiesGUIDExRequest struct {
	// dwObjectType:  Specifies the type of object for which properties are to be retrieved.
	// MUST be set to one of the object types, as specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pGuid:  MUST specify a pointer to the object identifier of the object for which
	// properties are to be retrieved.
	GUID *dtyp.GUID `idl:"name:pGuid;pointer:unique" json:"guid"`
	// cp:  MUST be set to the size (in elements) of the arrays aProp and apVar, which
	// for this method MUST be one (0x00000001). The arrays aProp and apVar MUST have an
	// identical number of elements, and MUST each contain exactly one element.
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// aProp:  An array of identifiers of properties to retrieve from the object. Each
	// element MUST specify a value from the property identifiers for the object type specified
	// in dwObjectType. Each element MUST specify the property identifier for the corresponding
	// property value at the same element index in apVar. The array MUST contain exactly
	// one element.
	Property []uint32 `idl:"name:aProp;size_is:(cp)" json:"property"`
	// apVar:  MUST be set by the client to an array that holds the property values retrieved
	// from the object. Each element MUST be set by the server to the property value for
	// the corresponding property identifier at the same element index in aProp. The array
	// MUST contain exactly one element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	// phServerAuth:  A PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle acquired from
	// the pphServerAuth parameter in a previous call to S_DSValidateServer. The server
	// MUST use this parameter as a key to locate the GSS security context used to compute
	// the signature returned in pbServerSignature. See section 3.1.4.2.
	ServerAuth *dscomm.ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	// pdwServerSignatureSize:  A DWORD that contains the maximum length of the server
	// signature in bytes to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
}

func (o *GetPropertiesGUIDExRequest) xxx_ToOp(ctx context.Context) *xxx_GetPropertiesGUIDExOperation {
	if o == nil {
		return &xxx_GetPropertiesGUIDExOperation{}
	}
	return &xxx_GetPropertiesGUIDExOperation{
		ObjectType:          o.ObjectType,
		GUID:                o.GUID,
		CreatePartition:     o.CreatePartition,
		Property:            o.Property,
		Var:                 o.Var,
		ServerAuth:          o.ServerAuth,
		ServerSignatureSize: o.ServerSignatureSize,
	}
}

func (o *GetPropertiesGUIDExRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesGUIDExOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.GUID = op.GUID
	o.CreatePartition = op.CreatePartition
	o.Property = op.Property
	o.Var = op.Var
	o.ServerAuth = op.ServerAuth
	o.ServerSignatureSize = op.ServerSignatureSize
}
func (o *GetPropertiesGUIDExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesGUIDExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesGUIDExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesGUIDExResponse structure represents the S_DSGetPropsGuidEx operation response
type GetPropertiesGUIDExResponse struct {
	// apVar:  MUST be set by the client to an array that holds the property values retrieved
	// from the object. Each element MUST be set by the server to the property value for
	// the corresponding property identifier at the same element index in aProp. The array
	// MUST contain exactly one element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	// pbServerSignature:  A buffer that contains a signed hash over the returned property
	// values. See the pbServerSignature parameter description in section 3.1.4.7.
	ServerSignature []byte `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	// pdwServerSignatureSize:  A DWORD that contains the maximum length of the server
	// signature in bytes to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	// Return: The S_DSGetPropsGuidEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesGUIDExResponse) xxx_ToOp(ctx context.Context) *xxx_GetPropertiesGUIDExOperation {
	if o == nil {
		return &xxx_GetPropertiesGUIDExOperation{}
	}
	return &xxx_GetPropertiesGUIDExOperation{
		Var:                 o.Var,
		ServerSignature:     o.ServerSignature,
		ServerSignatureSize: o.ServerSignatureSize,
		Return:              o.Return,
	}
}

func (o *GetPropertiesGUIDExResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesGUIDExOperation) {
	if o == nil {
		return
	}
	o.Var = op.Var
	o.ServerSignature = op.ServerSignature
	o.ServerSignatureSize = op.ServerSignatureSize
	o.Return = op.Return
}
func (o *GetPropertiesGUIDExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesGUIDExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesGUIDExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BeginDeleteNotificationOperation structure represents the S_DSBeginDeleteNotification operation
type xxx_BeginDeleteNotificationOperation struct {
	PathName   string                 `idl:"name:pwcsPathName;string" json:"path_name"`
	Handle     *dscomm.DeleteType     `idl:"name:pHandle" json:"handle"`
	ServerAuth *dscomm.ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	Return     int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_BeginDeleteNotificationOperation) OpNum() int { return 3 }

func (o *xxx_BeginDeleteNotificationOperation) OpName() string {
	return "/dscomm2/v1/S_DSBeginDeleteNotification"
}

func (o *xxx_BeginDeleteNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginDeleteNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dscomm.ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_BeginDeleteNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &dscomm.ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginDeleteNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginDeleteNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pHandle {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_DELETE_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_DELETE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dscomm.DeleteType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_BeginDeleteNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pHandle {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_DELETE_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_DELETE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &dscomm.DeleteType{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
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

// BeginDeleteNotificationRequest structure represents the S_DSBeginDeleteNotification operation request
type BeginDeleteNotificationRequest struct {
	// pwcsPathName:  MUST be set by the client to a pointer to a NULL-terminated 16-bit
	// Unicode string that contains the directory service pathname, as specified in section
	// 2.2.9, for an object of type MQDS_MACHINE or MQDS_QUEUE.<157>
	PathName string `idl:"name:pwcsPathName;string" json:"path_name"`
	// phServerAuth:  A PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle acquired from
	// the pphServerAuth parameter in a previous call to S_DSValidateServer. The server
	// MUST use this parameter as a key to locate the GSS security context used to compute
	// the signature returned in pbServerSignature. See section 3.1.4.2.
	ServerAuth *dscomm.ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
}

func (o *BeginDeleteNotificationRequest) xxx_ToOp(ctx context.Context) *xxx_BeginDeleteNotificationOperation {
	if o == nil {
		return &xxx_BeginDeleteNotificationOperation{}
	}
	return &xxx_BeginDeleteNotificationOperation{
		PathName:   o.PathName,
		ServerAuth: o.ServerAuth,
	}
}

func (o *BeginDeleteNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_BeginDeleteNotificationOperation) {
	if o == nil {
		return
	}
	o.PathName = op.PathName
	o.ServerAuth = op.ServerAuth
}
func (o *BeginDeleteNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BeginDeleteNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BeginDeleteNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BeginDeleteNotificationResponse structure represents the S_DSBeginDeleteNotification operation response
type BeginDeleteNotificationResponse struct {
	// pHandle:  MUST be set by the server to a pointer to a unique RPC context_handle
	// representing the delete notification. This handle is used by the client in subsequent
	// calls to S_DSNotifyDelete.
	Handle *dscomm.DeleteType `idl:"name:pHandle" json:"handle"`
	// Return: The S_DSBeginDeleteNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BeginDeleteNotificationResponse) xxx_ToOp(ctx context.Context) *xxx_BeginDeleteNotificationOperation {
	if o == nil {
		return &xxx_BeginDeleteNotificationOperation{}
	}
	return &xxx_BeginDeleteNotificationOperation{
		Handle: o.Handle,
		Return: o.Return,
	}
}

func (o *BeginDeleteNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_BeginDeleteNotificationOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Return = op.Return
}
func (o *BeginDeleteNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BeginDeleteNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BeginDeleteNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NotifyDeleteOperation structure represents the S_DSNotifyDelete operation
type xxx_NotifyDeleteOperation struct {
	Handle *dscomm.DeleteType `idl:"name:Handle" json:"handle"`
	Return int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_NotifyDeleteOperation) OpNum() int { return 4 }

func (o *xxx_NotifyDeleteOperation) OpName() string { return "/dscomm2/v1/S_DSNotifyDelete" }

func (o *xxx_NotifyDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NotifyDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_DELETE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dscomm.DeleteType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_NotifyDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_DELETE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &dscomm.DeleteType{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NotifyDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NotifyDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NotifyDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NotifyDeleteRequest structure represents the S_DSNotifyDelete operation request
type NotifyDeleteRequest struct {
	// Handle:  MUST be set by the client to a pointer to an RPC context_handle acquired
	// from a previous call to S_DSBeginDeleteNotification. This RPC context handle MUST
	// NOT have been used in a previous call to S_DSEndDeleteNotification.
	Handle *dscomm.DeleteType `idl:"name:Handle" json:"handle"`
}

func (o *NotifyDeleteRequest) xxx_ToOp(ctx context.Context) *xxx_NotifyDeleteOperation {
	if o == nil {
		return &xxx_NotifyDeleteOperation{}
	}
	return &xxx_NotifyDeleteOperation{
		Handle: o.Handle,
	}
}

func (o *NotifyDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_NotifyDeleteOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
}
func (o *NotifyDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *NotifyDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NotifyDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NotifyDeleteResponse structure represents the S_DSNotifyDelete operation response
type NotifyDeleteResponse struct {
	// Return: The S_DSNotifyDelete return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *NotifyDeleteResponse) xxx_ToOp(ctx context.Context) *xxx_NotifyDeleteOperation {
	if o == nil {
		return &xxx_NotifyDeleteOperation{}
	}
	return &xxx_NotifyDeleteOperation{
		Return: o.Return,
	}
}

func (o *NotifyDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_NotifyDeleteOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *NotifyDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *NotifyDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NotifyDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EndDeleteNotificationOperation structure represents the S_DSEndDeleteNotification operation
type xxx_EndDeleteNotificationOperation struct {
	Handle *dscomm.DeleteType `idl:"name:pHandle" json:"handle"`
}

func (o *xxx_EndDeleteNotificationOperation) OpNum() int { return 5 }

func (o *xxx_EndDeleteNotificationOperation) OpName() string {
	return "/dscomm2/v1/S_DSEndDeleteNotification"
}

func (o *xxx_EndDeleteNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndDeleteNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pHandle {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_DELETE_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_DELETE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dscomm.DeleteType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EndDeleteNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pHandle {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_DELETE_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_DELETE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &dscomm.DeleteType{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndDeleteNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndDeleteNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pHandle {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_DELETE_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_DELETE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dscomm.DeleteType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EndDeleteNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pHandle {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_DELETE_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_DELETE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &dscomm.DeleteType{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// EndDeleteNotificationRequest structure represents the S_DSEndDeleteNotification operation request
type EndDeleteNotificationRequest struct {
	// pHandle:  MUST be set by the client to a pointer to an RPC context_handle returned
	// by a previous call to S_DSBeginDeleteNotification. The RPC context handle MUST NOT
	// have been used in a previous call to S_DSEndDeleteNotification. The server MUST set
	// this parameter to NULL.
	Handle *dscomm.DeleteType `idl:"name:pHandle" json:"handle"`
}

func (o *EndDeleteNotificationRequest) xxx_ToOp(ctx context.Context) *xxx_EndDeleteNotificationOperation {
	if o == nil {
		return &xxx_EndDeleteNotificationOperation{}
	}
	return &xxx_EndDeleteNotificationOperation{
		Handle: o.Handle,
	}
}

func (o *EndDeleteNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_EndDeleteNotificationOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
}
func (o *EndDeleteNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EndDeleteNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EndDeleteNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EndDeleteNotificationResponse structure represents the S_DSEndDeleteNotification operation response
type EndDeleteNotificationResponse struct {
	// pHandle:  MUST be set by the client to a pointer to an RPC context_handle returned
	// by a previous call to S_DSBeginDeleteNotification. The RPC context handle MUST NOT
	// have been used in a previous call to S_DSEndDeleteNotification. The server MUST set
	// this parameter to NULL.
	Handle *dscomm.DeleteType `idl:"name:pHandle" json:"handle"`
}

func (o *EndDeleteNotificationResponse) xxx_ToOp(ctx context.Context) *xxx_EndDeleteNotificationOperation {
	if o == nil {
		return &xxx_EndDeleteNotificationOperation{}
	}
	return &xxx_EndDeleteNotificationOperation{
		Handle: o.Handle,
	}
}

func (o *EndDeleteNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_EndDeleteNotificationOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
}
func (o *EndDeleteNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EndDeleteNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EndDeleteNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsServerGCOperation structure represents the S_DSIsServerGC operation
type xxx_IsServerGCOperation struct {
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *xxx_IsServerGCOperation) OpNum() int { return 6 }

func (o *xxx_IsServerGCOperation) OpName() string { return "/dscomm2/v1/S_DSIsServerGC" }

func (o *xxx_IsServerGCOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsServerGCOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_IsServerGCOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_IsServerGCOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsServerGCOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsServerGCOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsServerGCRequest structure represents the S_DSIsServerGC operation request
type IsServerGCRequest struct {
}

func (o *IsServerGCRequest) xxx_ToOp(ctx context.Context) *xxx_IsServerGCOperation {
	if o == nil {
		return &xxx_IsServerGCOperation{}
	}
	return &xxx_IsServerGCOperation{}
}

func (o *IsServerGCRequest) xxx_FromOp(ctx context.Context, op *xxx_IsServerGCOperation) {
	if o == nil {
		return
	}
}
func (o *IsServerGCRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsServerGCRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsServerGCOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsServerGCResponse structure represents the S_DSIsServerGC operation response
type IsServerGCResponse struct {
	// Return: The S_DSIsServerGC return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsServerGCResponse) xxx_ToOp(ctx context.Context) *xxx_IsServerGCOperation {
	if o == nil {
		return &xxx_IsServerGCOperation{}
	}
	return &xxx_IsServerGCOperation{
		Return: o.Return,
	}
}

func (o *IsServerGCResponse) xxx_FromOp(ctx context.Context, op *xxx_IsServerGCOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *IsServerGCResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsServerGCResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsServerGCOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetGCListInDomainOperation structure represents the S_DSGetGCListInDomain operation
type xxx_GetGCListInDomainOperation struct {
	ComputerName        string                 `idl:"name:lpwszComputerName;string;pointer:ptr" json:"computer_name"`
	DomainName          string                 `idl:"name:lpwszDomainName;string;pointer:ptr" json:"domain_name"`
	GCList              string                 `idl:"name:lplpwszGCList;string" json:"gc_list"`
	ServerAuth          *dscomm.ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	ServerSignature     []byte                 `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	ServerSignatureSize uint32                 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	Return              int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetGCListInDomainOperation) OpNum() int { return 8 }

func (o *xxx_GetGCListInDomainOperation) OpName() string { return "/dscomm2/v1/S_DSGetGCListInDomain" }

func (o *xxx_GetGCListInDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGCListInDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpwszComputerName {in} (1:{string, pointer=ptr}*(1)[dim:0,string,null](wchar))
	{
		if o.ComputerName != "" {
			_ptr_lpwszComputerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ComputerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ComputerName, _ptr_lpwszComputerName); err != nil {
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
	// lpwszDomainName {in} (1:{string, pointer=ptr}*(1)[dim:0,string,null](wchar))
	{
		if o.DomainName != "" {
			_ptr_lpwszDomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DomainName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainName, _ptr_lpwszDomainName); err != nil {
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
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dscomm.ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGCListInDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpwszComputerName {in} (1:{string, pointer=ptr}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpwszComputerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ComputerName); err != nil {
				return err
			}
			return nil
		})
		_s_lpwszComputerName := func(ptr interface{}) { o.ComputerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ComputerName, _s_lpwszComputerName, _ptr_lpwszComputerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpwszDomainName {in} (1:{string, pointer=ptr}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpwszDomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DomainName); err != nil {
				return err
			}
			return nil
		})
		_s_lpwszDomainName := func(ptr interface{}) { o.DomainName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DomainName, _s_lpwszDomainName, _ptr_lpwszDomainName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &dscomm.ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGCListInDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ServerSignature != nil && o.ServerSignatureSize == 0 {
		o.ServerSignatureSize = uint32(len(o.ServerSignature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGCListInDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lplpwszGCList {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.GCList != "" {
			_ptr_lplpwszGCList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.GCList); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.GCList, _ptr_lplpwszGCList); err != nil {
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
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
	{
		dimSize1 := uint64(o.ServerSignatureSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ServerSignature {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ServerSignature[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerSignature); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
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

func (o *xxx_GetGCListInDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lplpwszGCList {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_lplpwszGCList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.GCList); err != nil {
				return err
			}
			return nil
		})
		_s_lplpwszGCList := func(ptr interface{}) { o.GCList = *ptr.(*string) }
		if err := w.ReadPointer(&o.GCList, _s_lplpwszGCList, _ptr_lplpwszGCList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ServerSignature", sizeInfo[0])
		}
		o.ServerSignature = make([]byte, sizeInfo[0])
		for i1 := range o.ServerSignature {
			i1 := i1
			if err := w.ReadData(&o.ServerSignature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
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

// GetGCListInDomainRequest structure represents the S_DSGetGCListInDomain operation request
type GetGCListInDomainRequest struct {
	// lpwszComputerName:  MUST be set by the client to NULL. The server will validate
	// the NULL setting.
	ComputerName string `idl:"name:lpwszComputerName;string;pointer:ptr" json:"computer_name"`
	// lpwszDomainName:  MUST be set by the client to the domain name of the domain to
	// query.
	DomainName string `idl:"name:lpwszDomainName;string;pointer:ptr" json:"domain_name"`
	// phServerAuth:  A PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle acquired from
	// the pphServerAuth parameter in a previous call to S_DSValidateServer. The server
	// MUST use this parameter as a key to locate the GSS security context used to compute
	// the signature returned in pbServerSignature.
	ServerAuth *dscomm.ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	// pdwServerSignatureSize:  MUST be set by the client to point to a DWORD that contains
	// the maximum size in bytes of the server signature to return, and MUST be set by the
	// server to contain the actual size in bytes of the server signature on output.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
}

func (o *GetGCListInDomainRequest) xxx_ToOp(ctx context.Context) *xxx_GetGCListInDomainOperation {
	if o == nil {
		return &xxx_GetGCListInDomainOperation{}
	}
	return &xxx_GetGCListInDomainOperation{
		ComputerName:        o.ComputerName,
		DomainName:          o.DomainName,
		ServerAuth:          o.ServerAuth,
		ServerSignatureSize: o.ServerSignatureSize,
	}
}

func (o *GetGCListInDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_GetGCListInDomainOperation) {
	if o == nil {
		return
	}
	o.ComputerName = op.ComputerName
	o.DomainName = op.DomainName
	o.ServerAuth = op.ServerAuth
	o.ServerSignatureSize = op.ServerSignatureSize
}
func (o *GetGCListInDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetGCListInDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetGCListInDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetGCListInDomainResponse structure represents the S_DSGetGCListInDomain operation response
type GetGCListInDomainResponse struct {
	// lplpwszGCList:  MUST be set by the server to the list of Global Catalog Servers.
	// The format of the list is a Server Specification List String.
	GCList string `idl:"name:lplpwszGCList;string" json:"gc_list"`
	// pbServerSignature:  MUST be set by the server to a buffer that contains a signed
	// hash over the returned list of Global Catalog Servers (lplpwszGCList) calculated
	// by using the MD5 algorithm (as specified in [RFC1321]) and the GSS security context,
	// as specified by the following pseudocode.
	ServerSignature []byte `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	// pdwServerSignatureSize:  MUST be set by the client to point to a DWORD that contains
	// the maximum size in bytes of the server signature to return, and MUST be set by the
	// server to contain the actual size in bytes of the server signature on output.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	// Return: The S_DSGetGCListInDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetGCListInDomainResponse) xxx_ToOp(ctx context.Context) *xxx_GetGCListInDomainOperation {
	if o == nil {
		return &xxx_GetGCListInDomainOperation{}
	}
	return &xxx_GetGCListInDomainOperation{
		GCList:              o.GCList,
		ServerSignature:     o.ServerSignature,
		ServerSignatureSize: o.ServerSignatureSize,
		Return:              o.Return,
	}
}

func (o *GetGCListInDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_GetGCListInDomainOperation) {
	if o == nil {
		return
	}
	o.GCList = op.GCList
	o.ServerSignature = op.ServerSignature
	o.ServerSignatureSize = op.ServerSignatureSize
	o.Return = op.Return
}
func (o *GetGCListInDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetGCListInDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetGCListInDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
