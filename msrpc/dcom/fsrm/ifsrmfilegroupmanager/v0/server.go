package ifsrmfilegroupmanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
	_ = idispatch.GoPackage
)

// IFsrmFileGroupManager server interface.
type FileGroupManagerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The CreateFileGroup method creates a blank Non-Persisted File Group Instance (section
	// 3.2.1.3.4.2).
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------+
	//	|         RETURN          |                                  |
	//	|       VALUE/CODE        |           DESCRIPTION            |
	//	|                         |                                  |
	//	+-------------------------+----------------------------------+
	//	+-------------------------+----------------------------------+
	//	| 0x80070057 E_INVALIDARG | The fileGroup parameter is NULL. |
	//	+-------------------------+----------------------------------+
	CreateFileGroup(context.Context, *CreateFileGroupRequest) (*CreateFileGroupResponse, error)

	// The GetFileGroup method returns a pointer to the file group with the specified Name
	// property from the List of Persisted File Groups (section 3.2.1.3).
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND    | The specified file group could not be found.                                     |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045308 FSRM_E_INVALID_NAME | The specified name is not valid.                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE | The content of the name parameter exceeds the maximum length of 4,000            |
	//	|                                | characters.                                                                      |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | The fileGroup parameter is NULL.                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	GetFileGroup(context.Context, *GetFileGroupRequest) (*GetFileGroupResponse, error)

	// The EnumFileGroups method returns all the file groups from the List of Persisted
	// File Groups (section 3.2.1.3) of the server.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+-----------------------------------------------------------------+
	//	|             RETURN              |                                                                 |
	//	|           VALUE/CODE            |                           DESCRIPTION                           |
	//	|                                 |                                                                 |
	//	+---------------------------------+-----------------------------------------------------------------+
	//	+---------------------------------+-----------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | This options parameter contains invalid FsrmEnumOptions values. |
	//	+---------------------------------+-----------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | The fileGroups parameter is NULL.                               |
	//	+---------------------------------+-----------------------------------------------------------------+
	EnumFileGroups(context.Context, *EnumFileGroupsRequest) (*EnumFileGroupsResponse, error)

	// The ExportFileGroups method exports an XML string representation of the File Server
	// Resource Manager Protocol file groups from the List of Persisted File Groups (section
	// 3.2.1.3).
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|           RETURN            |                                                                                  |
	//	|         VALUE/CODE          |                                   DESCRIPTION                                    |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | The specified file group could not be found.                                     |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG     | This code is returned for the following reasons: The serializedFileGroups        |
	//	|                             | parameter is NULL. The fileGroupNamesArray parameter is not a variant array of   |
	//	|                             | BSTRs.                                                                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	ExportFileGroups(context.Context, *ExportFileGroupsRequest) (*ExportFileGroupsResponse, error)

	// The ImportFileGroups method imports file groups from the XML string of file groups.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND              | This specified file group could not be found.                                    |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004530B FSRM_E_INVALID_IMPORT_VERSION | The version of the imported objects is not valid.                                |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                  | This code is returned for the following reasons: The serializedFileGroups        |
	//	|                                          | parameter is NULL. The fileGroups parameter is NULL. The fileGroupNamesArray     |
	//	|                                          | parameter is not a variant array of BSTRs.                                       |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	ImportFileGroups(context.Context, *ImportFileGroupsRequest) (*ImportFileGroupsResponse, error)
}

func RegisterFileGroupManagerServer(conn dcerpc.Conn, o FileGroupManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFileGroupManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FileGroupManagerSyntaxV0_0))...)
}

func NewFileGroupManagerServerHandle(o FileGroupManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FileGroupManagerServerHandle(ctx, o, opNum, r)
	}
}

func FileGroupManagerServerHandle(ctx context.Context, o FileGroupManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CreateFileGroup
		in := &CreateFileGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateFileGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // GetFileGroup
		in := &GetFileGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // EnumFileGroups
		in := &EnumFileGroupsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumFileGroups(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // ExportFileGroups
		in := &ExportFileGroupsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExportFileGroups(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // ImportFileGroups
		in := &ImportFileGroupsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ImportFileGroups(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
