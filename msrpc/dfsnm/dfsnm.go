// The dfsnm package implements the DFSNM client protocol.
//
// # Introduction
//
// The Distributed File System (DFS): Namespace Management Protocol provides a remote
// procedure call (RPC) interface for administering DFS configurations. The client is
// an application that issues method calls on the RPC interface to administer DFS. The
// server is a DFS service that implements support for this RPC interface for administering
// DFS.
//
// # Overview
//
// The DFS: Namespace Management Protocol is one of a collection of protocols that group
// shares that are located on different servers by combining various storage media into
// a single logical namespace. The DFS namespace is a virtual view of the share. When
// a user views the namespace, the directories and files in it appear to reside on a
// single share. Users can navigate the namespace without needing to know the server
// names or shares hosting the data. DFS also provides redundancy of namespace service.
//
// Access to a DFS namespace requires the DFS client. The DFS client uses the DFS Referral
// Protocol, as specified in [MS-DFSC], to ascertain the existence of the DFS namespace
// and to determine the shares to access on servers that participate in the DFS namespace.
// The DFS Referral Protocol navigates through the DFS namespace by appropriately issuing
// referral requests to a domain controller (DC) or to a DFS root target server to resolve
// the original path to a share on a server that contains the data being accessed. For
// more information on DFS and the DFS client, see [MSDFS]. For more information on
// how the DFS Referral Protocol operates within the context of the Server Message Block
// (SMB) Protocol, as specified in [MS-SMB], which is the transport for DFS referrals,
// see [MS-DFSC] section 2.
//
// DFS namespace information, such as name, DFS link name, DFS link target, and so on,
// is stored in the DFS metadata of the namespace. Depending on where the DFS metadata
// is stored, the DFS namespace is "domain-based" or "stand-alone".
//
// * Domain-Based DFS Namespace ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_133ab305-ad75-4b38-8ec9-40bc0dc3b594
// ) : A well-known container in the domain directory, known as the DFS configuration
// container, holds the DFS metadata for a domain-based DFS namespace. An object ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_8bb43a65-7a8c-4585-a7ed-23044772f8ca
// ) exists for each domain-based DFS namespace in the DFS configuration container.
// DFS metadata of a domain-based DFS namespace is stored as a binary large object (BLOB)
// ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_ad861812-8cb0-497a-80bb-13c95aa4e425 )
// in an attribute of the DFS namespace object. A domain-based DFS namespace can have
// multiple DFS root targets, which offer high availability and load sharing at the
// DFS root level. The DFS root ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_639b7503-b879-4ef7-98a8-14adf85bc16d
// ) name of a domain-based DFS namespace has the domain ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_b0276eb2-4e65-4cf1-a718-e0920a614aca
// ) as its first component. A DFS client issues a referral request to a DC in order
// to identify the DFS root targets of the DFS namespace.
//
// * Stand-Alone DFS Namespace ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_701a2cc1-eb5a-400f-b394-5bad264ec8f6
// ) : DFS metadata is stored in an implementation-specific format on the DFS root target
// server itself. A stand-alone DFS namespace supports only one DFS root target. The
// DFS root name of a stand-alone DFS namespace has a host name ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_fe6fd875-b30a-44bb-850e-10733990f04c
// ) as its first component. A DFS client issues referral requests to the DFS root target
// server to access the DFS namespace. A stand-alone DFS namespace can be clustered
// to provide high availability of the DFS namespace. <1> ( 3a466440-1ef6-4439-b4e2-be9eaddeb511#Appendix_A_1
// ) The server hosting a stand-alone DFS namespace can be promoted to a Domain Controller,
// but the namespace cannot be converted to a domain-based namespace, and it will continue
// as a stand-alone namespace.
//
// A server cannot host both domain-based and stand-alone namespace roots with the same
// name.
//
// The DFS: Namespace Management Protocol is used to configure DFS services. This protocol
// is used primarily by administrative applications that run on client computers to
// connect and configure Distributed File System (DFS) servers. It consists of the RPC
// methods that can be issued from an administrative client computer to the protocol
// server on a DC or a Distributed File System (DFS) root target server. An administrator
// can use this protocol to perform various Distributed File System (DFS) namespace
// administration operations, such as creating or deleting a DFS namespace, adding or
// removing DFS root targets, adding or removing DFS links, and adding or removing targets
// to an existing link. The DFS: Namespace Management Protocol includes the following:
//
// * Eleven basic methods for configuring stand-alone DFS namespaces ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_ce6af522-ba70-4ba1-a684-b98b809c72ad
// ) and domain-based DFS namespaces ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_c37de1c8-4bd3-406f-ad8c-50c877666f91
// ) , as specified in section 3.1.4.1 ( 7f56ca02-70ff-45ed-b4d3-0e679907d1ed ).
//
// * Four methods that support extended access to configurations of a DFS namespace,
// as specified in section 3.1.4.2 ( b8f452a8-2a48-4c3e-934c-ab5f2f1a2748 ).
//
// * Three methods for configuring root targets in a domainv1-based DFS namespace (
// a9bc4403-a862-48b9-b99b-1b44a887d177#gt_2aad5e61-ffe9-406e-a192-328c5327ee72 ) ,
// as specified in section 3.1.4.3 ( 98df4e5d-7092-4cb2-80cf-48e535e132b6 ).
//
// * Three methods for configuring a stand-alone DFS namespace, as specified in section
// 3.1.4.4 ( f639efff-4116-4556-b71b-9c006e8f265b ).
//
// * Two methods relating to the association between a DFS server ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_3de8e640-501a-4b25-80a7-0ba769f3c0b9
// ) and the DC used by a domain-based DFS namespace, as specified in section 3.1.4.5
// ( ed6babdd-d313-486b-a606-22f938afe988 ).
//
// Much of the configuration information that is communicated through this protocol
// is marshaled through two unions: DFS_INFO_STRUCT and DFS_INFO_ENUM_STRUCT. The usage
// model of these unions is for the client to specify a Level parameter to determine
// which union case to use. Each level corresponds to a specific DFS_INFO_n structure,
// where n is the level number. Arrays of DFS_INFO_n structures are marshaled using
// DFS_INFO_n_CONTAINER structures. Levels 1, 2, 3, 4, 5, 6, 8, and 9 are common, and
// are shared across both the DFS_INFO_STRUCT and DFS_INFO_ENUM_STRUCT unions. Levels
// 7, 50, 100, 101, 102, 103, 104, 105, 106, 107, and 150 are unique to the DFS_INFO_STRUCT
// union, and Levels 200 and 300 are unique to the DFS_INFO_ENUM_STRUCT union.
//
// While a number of methods use the common configuration information structures, not
// all methods support all levels. The following table lists the levels used in the
// DFS_INFO_STRUCT and DFS_INFO_ENUM_STRUCT unions, their singleton and array structures,
// and the methods with which the level can be used.
//
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|       |              |          ARRAY          |     NETRDFS     |   NETRDFS    |     NETRDFS     |     NETRDFS      |    NETRDFS     |
//	| LEVEL |  STRUCTURE   |        STRUCTURE        |     GETINFO     |     ENUM     |     SETINFO     |     SETINFO2     |     ENUMEX     |
//	|       |              |                         |                 |              |                 |                  |                |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|     1 | DFS_INFO_1   | DFS_INFO_1_ CONTAINER   | X               | X            |                 |                  | X              |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|     2 | DFS_INFO_2   | DFS_INFO_2_ CONTAINER   | X               | X            |                 |                  | X              |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|     3 | DFS_INFO_3   | DFS_INFO_3_ CONTAINER   | X               | X            |                 |                  | X              |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|     4 | DFS_INFO_4   | DFS_INFO_4_ CONTAINER   | X               | X            |                 |                  | X              |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|     5 | DFS_INFO_5   | DFS_INFO_5_ CONTAINER   | X               | X            |                 |                  | X              |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|     6 | DFS_INFO_6   | DFS_INFO_6_ CONTAINER   | X               | X            |                 |                  | X              |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|     7 | DFS_INFO_7   | N/A                     | X               |              |                 |                  |                |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|     8 | DFS_INFO_8   | DFS_INFO_8_ CONTAINER   | X               | X            |                 |                  | X              |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|     9 | DFS_INFO_9   | DFS_INFO_9_ CONTAINER   | X               | X            |                 |                  | X              |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|    50 | DFS_INFO_50  | N/A                     | X               |              |                 |                  |                |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|   100 | DFS_INFO_100 | N/A                     | X               |              | X               | X                |                |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|   101 | DFS_INFO_101 | N/A                     |                 |              | X               | X                |                |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|   102 | DFS_INFO_102 | N/A                     |                 |              | X               | X                |                |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|   103 | DFS_INFO_103 | N/A                     |                 |              | X               | X                |                |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|   104 | DFS_INFO_104 | N/A                     |                 |              | X               | X                |                |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|   105 | DFS_INFO_105 | N/A                     |                 |              | X               | X                |                |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|   106 | DFS_INFO_106 | N/A                     |                 |              | X               | X                |                |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|   107 | DFS_INFO_107 | N/A                     |                 |              | X               | X                |                |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|   150 | DFS_INFO_150 | N/A                     | X               |              | X               | X                |                |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|   200 | DFS_INFO_200 | DFS_INFO_200_ CONTAINER |                 |              |                 |                  | X              |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
//	|   300 | DFS_INFO_300 | DFS_INFO_300_ CONTAINER |                 | X            |                 |                  | X              |
//	+-------+--------------+-------------------------+-----------------+--------------+-----------------+------------------+----------------+
package dfsnm

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

var (
	// import guard
	GoPackage = "dfsnm"
)
