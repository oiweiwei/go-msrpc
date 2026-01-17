package trkwks

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dltw "github.com/oiweiwei/go-msrpc/msrpc/dltw"
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
	_ = dltw.GoPackage
)

var (
	// import guard
	GoPackage = "dltw"
)

var (
	// Syntax UUID
	TrkwksSyntaxUUID = &uuid.UUID{TimeLow: 0x300f3532, TimeMid: 0x38cc, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0xf0, Node: [6]uint8{0x0, 0x20, 0xaf, 0x6b, 0xa, 0xdd}}
	// Syntax ID
	TrkwksSyntaxV1_2 = &dcerpc.SyntaxID{IfUUID: TrkwksSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 2}
)

// trkwks interface.
type TrkwksClient interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

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

	// The LnkSearchMachine method searches for a file object on the specified computer.
	// If information on the file is found, the method attempts to return it. If the file
	// has been moved, the method returns information about its new location.
	//
	// Return Values: A 32-bit integer that indicates success or failure: A value of 0 or
	// any positive value indicates success; a negative value indicates failure. Some of
	// the possible return codes are listed in the following table.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8DEAD106 TRK_E_POTENTIAL_FILE_FOUND | A file was found by a DLT Workstation server processing a LnkSearchMachine call, |
	//	|                                       | and the file meets some—but not all—of the necessary criteria to be considered   |
	//	|                                       | the correct file. More details are presented later in this section.              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8DEAD101 TRK_E_REFERRAL             | A file could not be found by a DLT Workstation server processing a               |
	//	|                                       | LnkSearchMachine call, and the file meets some—but not all—of the necessary      |
	//	|                                       | criteria to be considered the correct file. More details are presented later in  |
	//	|                                       | this section.                                                                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: None.
	//
	// The server responds to this call by attempting to find the requested file in one
	// of its volumes, and either returns the file's new location (in the form of a UNC),
	// or returns information about the location to which the file has moved (see description
	// of TRK_E_REFERRAL in the following bulleted list).
	//
	// The server SHOULD<10> return a failure value, if the file is found, but the length
	// of the path to be returned in the pstszPath argument exceeds 261 characters (where
	// the 261 characters does not include a string terminator character).
	//
	// For each of the possible return values, the output parameters of this call MUST be
	// set as specified in the following list. In these specifications, the requested ObjectID
	// MUST correspond to the Object field of the pDroidLast parameter, the requested VolumeID
	// MUST correspond to the Volume field of the pDroidLast parameter, and the FileID requested
	// MUST correspond to the value of the pDroidBirthLast parameter. The numeric values
	// corresponding to these returns are defined in section 3.1.1:
	//
	// * Success: A success value MUST be returned if the server finds the file, and the
	// client is authorized to get information about the file via a UNC path. The file MUST
	// be found by searching all volumes on the server computer for a file whose ObjectID
	// is equal to the ObjectID of the request, and whose FileID is equal to the FileID
	// of the request. To perform the authorization check, the server MUST use the client's
	// identity (obtained as specified in [MS-RPCE] ( ../ms-rpce/290c38b1-92fe-4229-91e6-4fc376610c15
	// ) section 3.3.3.4.3 ( ../ms-rpce/29b8217a-0bda-4fdb-a3ea-48560125ae8d ) ) to determine,
	// based on local policy, whether or not the client is authorized to get the UNC of
	// the file.
	//
	// If there is more than one file on the server computer that satisfies these conditions,
	// the file MUST be selected as follows:
	//
	// * If one of the matching files is on the volume whose VolumeID equals the VolumeID
	// of the request, that file is selected.
	//
	// * Otherwise, the behavior is arbitrary.
	//
	// The server MUST set the output parameters as follows:
	//
	// pdroidBirthNext : The value of the pdroidBirthLast that is specified by the client.
	//
	// pdroidNext : The current FileLocation value for the file.
	//
	// pmcidNext : The MachineID of the server computer.
	//
	// ptszPath : The UNC of the file on the server computer. If the file can be located
	// under multiple UNC paths, one of those UNC paths MUST be returned, but the server
	// MAY return any of those paths. <11> ( caa74a29-27fb-4f50-afa5-bbf152fe5d27#Appendix_A_11
	// )
	//
	// * TRK_E_REFERRAL: This value MUST be returned if the file is no longer stored on
	// any of the volumes of the server computer, but if the MoveTable of the last volume
	// has an entry for the file. (Note that there is no additional per-file access check
	// in this case.) The last volume MUST be determined by using the volume on the server
	// whose VolumeID equals the VolumeID of the request. The entry in the MoveTable for
	// the file, as defined in section 3.1.1, MUST be that entry whose ObjectID equals the
	// ObjectID of the request.
	//
	// The server MUST set the output parameters as follows:
	//
	// pdroidBirthNext : The value of the FileID specified by the client in the *pdroidBirthLast*
	// field of the call.
	//
	// pdroidNext : The value of the FileLocation field for the file's entry in the *MoveTable*.
	//
	// pmcidNext : The value of the MachineID field for the file's entry in the *MoveTable*.
	//
	// ptszPath : The server MUST NOT modify this value.
	//
	// * TRK_E_POTENTIAL_FILE_FOUND: This value MUST be returned if there is no file on
	// any of the server's volumes with the requested FileID and ObjectID, and there is
	// no entry for this file's ObjectID in the MoveTable, but if there is a file with the
	// requested ObjectID on one of the server's volumes, and the FileID on that file is
	// all zeros. <12> ( caa74a29-27fb-4f50-afa5-bbf152fe5d27#Appendix_A_12 )
	//
	// The server MUST set the output parameters as follows:
	//
	// pdroidBirthNext : The value of the FileID for the found file.
	//
	// pdroidNext : The FileLocation of the file found by the server.
	//
	// pmcidNext : The MachineID of the server computer.
	//
	// ptszPath : A UNC of the file on the server computer.
	//
	// * Other negative return values: A negative value other than those mentioned above
	// MUST be returned if none of the preceding cases is met. The server MUST NOT modify
	// any of the output parameters.
	SearchMachine(context.Context, *SearchMachineRequest, ...dcerpc.CallOption) (*SearchMachineResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultTrkwksClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultTrkwksClient) SearchMachine(ctx context.Context, in *SearchMachineRequest, opts ...dcerpc.CallOption) (*SearchMachineResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SearchMachineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTrkwksClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTrkwksClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewTrkwksClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TrkwksClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TrkwksSyntaxV1_2))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultTrkwksClient{cc: cc}, nil
}

// xxx_SearchMachineOperation structure represents the LnkSearchMachine operation
type xxx_SearchMachineOperation struct {
	Restrictions  uint32                       `idl:"name:Restrictions" json:"restrictions"`
	BirthLast     *dltw.DomainRelativeObjectID `idl:"name:pdroidBirthLast" json:"birth_last"`
	LastObjectID  *dltw.DomainRelativeObjectID `idl:"name:pdroidLast" json:"last_object_id"`
	BirthNext     *dltw.DomainRelativeObjectID `idl:"name:pdroidBirthNext" json:"birth_next"`
	NextObjectID  *dltw.DomainRelativeObjectID `idl:"name:pdroidNext" json:"next_object_id"`
	NextMachineID *dltw.MachineID              `idl:"name:pmcidNext" json:"next_machine_id"`
	Path          string                       `idl:"name:ptszPath;max_is:(261);string" json:"path"`
	Return        int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_SearchMachineOperation) OpNum() int { return 12 }

func (o *xxx_SearchMachineOperation) OpName() string { return "/trkwks/v1.2/LnkSearchMachine" }

func (o *xxx_SearchMachineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SearchMachineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Restrictions {in} (1:(uint32))
	{
		if err := w.WriteData(o.Restrictions); err != nil {
			return err
		}
	}
	// pdroidBirthLast {in} (1:{pointer=ref}*(1))(2:{alias=CDomainRelativeObjId}(struct))
	{
		if o.BirthLast != nil {
			if err := o.BirthLast.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dltw.DomainRelativeObjectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdroidLast {in} (1:{pointer=ref}*(1))(2:{alias=CDomainRelativeObjId}(struct))
	{
		if o.LastObjectID != nil {
			if err := o.LastObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dltw.DomainRelativeObjectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SearchMachineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Restrictions {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Restrictions); err != nil {
			return err
		}
	}
	// pdroidBirthLast {in} (1:{pointer=ref}*(1))(2:{alias=CDomainRelativeObjId}(struct))
	{
		if o.BirthLast == nil {
			o.BirthLast = &dltw.DomainRelativeObjectID{}
		}
		if err := o.BirthLast.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdroidLast {in} (1:{pointer=ref}*(1))(2:{alias=CDomainRelativeObjId}(struct))
	{
		if o.LastObjectID == nil {
			o.LastObjectID = &dltw.DomainRelativeObjectID{}
		}
		if err := o.LastObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SearchMachineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	// cannot evaluate expression 262
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SearchMachineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdroidBirthNext {out} (1:{pointer=ref}*(1))(2:{alias=CDomainRelativeObjId}(struct))
	{
		if o.BirthNext != nil {
			if err := o.BirthNext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dltw.DomainRelativeObjectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdroidNext {out} (1:{pointer=ref}*(1))(2:{alias=CDomainRelativeObjId}(struct))
	{
		if o.NextObjectID != nil {
			if err := o.NextObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dltw.DomainRelativeObjectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pmcidNext {out} (1:{pointer=ref}*(1))(2:{alias=CMachineId}(struct))
	{
		if o.NextMachineID != nil {
			if err := o.NextMachineID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dltw.MachineID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ptszPath {out} (1:{string, pointer=ref}*(1)[dim:0,max_is=261,string,null](wchar))
	{
		dimSize1 := uint64(262)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.UTF16NLen(o.Path)
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
		_Path_buf := utf16.Encode([]rune(o.Path))
		if uint64(len(_Path_buf)) > sizeInfo[0]-1 {
			_Path_buf = _Path_buf[:sizeInfo[0]-1]
		}
		if o.Path != ndr.ZeroString {
			_Path_buf = append(_Path_buf, uint16(0))
		}
		for i1 := range _Path_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Path_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Path_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
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

func (o *xxx_SearchMachineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdroidBirthNext {out} (1:{pointer=ref}*(1))(2:{alias=CDomainRelativeObjId}(struct))
	{
		if o.BirthNext == nil {
			o.BirthNext = &dltw.DomainRelativeObjectID{}
		}
		if err := o.BirthNext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdroidNext {out} (1:{pointer=ref}*(1))(2:{alias=CDomainRelativeObjId}(struct))
	{
		if o.NextObjectID == nil {
			o.NextObjectID = &dltw.DomainRelativeObjectID{}
		}
		if err := o.NextObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pmcidNext {out} (1:{pointer=ref}*(1))(2:{alias=CMachineId}(struct))
	{
		if o.NextMachineID == nil {
			o.NextMachineID = &dltw.MachineID{}
		}
		if err := o.NextMachineID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ptszPath {out} (1:{string, pointer=ref}*(1)[dim:0,max_is=261,string,null](wchar))
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
		var _Path_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Path_buf", sizeInfo[0])
		}
		_Path_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Path_buf {
			i1 := i1
			if err := w.ReadData(&_Path_buf[i1]); err != nil {
				return err
			}
		}
		o.Path = strings.TrimRight(string(utf16.Decode(_Path_buf)), ndr.ZeroString)
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SearchMachineRequest structure represents the LnkSearchMachine operation request
type SearchMachineRequest struct {
	// Restrictions: SHOULD be set to zero.<9>
	Restrictions uint32 `idl:"name:Restrictions" json:"restrictions"`
	// pdroidBirthLast: FileID of the file for which the server is to search.
	BirthLast *dltw.DomainRelativeObjectID `idl:"name:pdroidBirthLast" json:"birth_last"`
	// pdroidLast: FileLocation of the file for which the server is to search. This is the
	// last known FileLocation by the client.
	LastObjectID *dltw.DomainRelativeObjectID `idl:"name:pdroidLast" json:"last_object_id"`
}

func (o *SearchMachineRequest) xxx_ToOp(ctx context.Context, op *xxx_SearchMachineOperation) *xxx_SearchMachineOperation {
	if op == nil {
		op = &xxx_SearchMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.Restrictions = o.Restrictions
	op.BirthLast = o.BirthLast
	op.LastObjectID = o.LastObjectID
	return op
}

func (o *SearchMachineRequest) xxx_FromOp(ctx context.Context, op *xxx_SearchMachineOperation) {
	if o == nil {
		return
	}
	o.Restrictions = op.Restrictions
	o.BirthLast = op.BirthLast
	o.LastObjectID = op.LastObjectID
}
func (o *SearchMachineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SearchMachineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SearchMachineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SearchMachineResponse structure represents the LnkSearchMachine operation response
type SearchMachineResponse struct {
	// pdroidBirthNext: A new FileID returned from the server. This indicates that the server
	// could not find the file, but that a different file that might be correct has been
	// found. This is associated with the TRK_E_POTENTIAL_FILE_FOUND return value, which
	// is defined below. How the server responds to that error is specified later in this
	// section.
	BirthNext *dltw.DomainRelativeObjectID `idl:"name:pdroidBirthNext" json:"birth_next"`
	// pdroidNext: New FileLocation for the file returned from the server.
	NextObjectID *dltw.DomainRelativeObjectID `idl:"name:pdroidNext" json:"next_object_id"`
	// pmcidNext: New MachineID for the computer that holds the file.
	NextMachineID *dltw.MachineID `idl:"name:pmcidNext" json:"next_machine_id"`
	// ptszPath: New UNC for file.
	Path string `idl:"name:ptszPath;max_is:(261);string" json:"path"`
	// Return: The LnkSearchMachine return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SearchMachineResponse) xxx_ToOp(ctx context.Context, op *xxx_SearchMachineOperation) *xxx_SearchMachineOperation {
	if op == nil {
		op = &xxx_SearchMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.BirthNext = o.BirthNext
	op.NextObjectID = o.NextObjectID
	op.NextMachineID = o.NextMachineID
	op.Path = o.Path
	op.Return = o.Return
	return op
}

func (o *SearchMachineResponse) xxx_FromOp(ctx context.Context, op *xxx_SearchMachineOperation) {
	if o == nil {
		return
	}
	o.BirthNext = op.BirthNext
	o.NextObjectID = op.NextObjectID
	o.NextMachineID = op.NextMachineID
	o.Path = op.Path
	o.Return = op.Return
}
func (o *SearchMachineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SearchMachineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SearchMachineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
