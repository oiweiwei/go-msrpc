// The frs1 package implements the FRS1 client protocol.
//
// # Introduction
//
// The File Replication Service (FRS) Protocol is used to replicate files and folders
// among servers on the network. This protocol enables duplicate files and folders to
// be maintained on multiple servers. FRS  is used to maintain duplicate copies of
// data files in the system volume (SYSVOL) system folders on multiple domain controllers
// in a domain. FRS is used to replicate data files among Distributed File System (DFS)
// shares. Detailed technical reference material for the Remote Procedure Call (RPC)
// interfaces, packet formats, and data structures required for interoperation using
// FRS are provided.
//
// # Overview
//
// The File Replication Service (FRS) Protocol is a multimaster replication protocol
// that is used to replicate files and folders across one or more members in an Active
// Directory domain. It works to keep copies of a file system tree up to date on all
// members of a replication group, while allowing any member of the group to change
// the contents at any time. A particular file system tree being replicated, along with
// the set of machines to which it is replicated, and the topology of connections between
// those machines used for replication, is known as a replica set.
//
// The topology of a replica set is a directed graph over the machines in the set. Because
// the graph is directed, data only flows in one direction on any given connection.
// All machines in a replica set participate as a client on some connections and a server
// on others. As the machines in a replica set update the contents of the replicated
// folder, they are responsible for generating change orders that propagate around the
// FRS topology. This causes the other members of the replica set to be aware of and
// (absent a conflict) replicate the update.
//
// Each machine in a replication set keeps a volume sequence number (VSN) that is incremented
// each time it generates a change order. Each member of a replica set keeps track of
// the highest VSN that it knows about for each member of the replica set; together
// they are known as a version vector. By sending its version vector to its upstream
// partners in the replication topology, the upstream partners can efficiently determine
// what changes need to be sent and what changes are already known, and send only the
// appropriate set of change orders back to the downstream partner. For more information
// on version vectors see section 3.1.1.11.
//
// On a given machine, FRS learns about all replica sets that it is part of, along with
// all its immediate partners, through a set of Active Directory objects. Replication
// topology is defined by two types of objects: Member objects represent a given participant,
// and connection objects that connect two endpoint member objects and define the direction
// of data flow along with the replication schedule.
//
// FRS detects changes made to any file or folder underneath the replica tree root.
// Details of a given local change are captured in a change order construct. A change
// order represents an action that took place on the local file system, such as file
// write, creation, deletion, or rename. In addition, FRS maintains a single ID record
// for every file or folder underneath the replica tree root in the file system. The
// ID record provides the information FRS needs to locate the file on the file system.
// The ID record also stores any extra properties for the resource, such as file attributes.
//
// A new FRS participant goes through a process called initial sync. This process creates
// the initial content on the new member by requesting all the data from upstream partners.
// New participants cannot replicate local changes until the initial sync is concluded.
//
// At sync time, which is defined by the connection schedule, FRS establishes a connection
// with its upstream partner through the remote procedure call (RPC) interface exposed
// by every running FRS instance (one FRS instance per server). The connection is directed,
// so changes flow from the upstream partner to the downstream partner. FRS receives
// the version vector from its downstream partner in a process called Version Vector
// Join (or VVJoin), during which the upstream partner determines the changes that it
// needs to send to the downstream partner. For every such change, the upstream partner
// passes an appropriate change order to the downstream partner. The downstream partner
// inspects every change order it receives and decides, based on its local changes,
// to accept or reject the change. Typically, a change order is rejected if the local
// version supersedes the remote version of the resource. On accepting the change order,
// the downstream partner fetches the resource via one or more stage packets that carry
// the data as part of their payload in FrsRpcSendCommPkt method requests. Large files
// are partitioned into several stage packets that are serialized so that the downstream
// partner can reconstruct the file after receiving all the pieces.
//
// File contents are marshaled before transfer over the wire to capture file system-specific
// metadata along with file data in one binary stream. The marshaled representation
// of a file is known as a staging file. The receiving partner has to be able to unmarshal
// the file at its end prior to placing it in the target location on the file system.
// The replicated file can also be compressed when it is marshaled to save bandwidth.
//
// If two or more users are creating files with the same file name on different replica
// set members, these files will have name collisions with each other as they replicate
// to other members. Each of these created files is a distinct object with unique content,
// but only one can be kept. FRS detects that a name collision has occurred when the
// second file is replicated to a member after a previous file has arrived. FRS then
// performs the last-writer-wins reconciliation between the two distinct objects. The
// loser gets deleted, and the delete is propagated out to the other members. The winner
// keeps the name and gets installed on the member.
//
// For folders, things are a bit different because they can have files and folders underneath
// them. In this case, FRS again detects the name collision when the second folder is
// replicated to a member, and it performs last-writer-wins reconciliation; except,
// in this case, the winner gets a new non-conflicting file name (referred to as a morphed
// name), and the loser gets to keep the original folder name. The rename is replicated
// out so all copies of the renamed folder object get the same new name.
//
// FRS supports four types of replica sets:
//
// * FRS_RSTYPE_ENTERPRISE_SYSVOL (1)
//
// * FRS_RSTYPE_DOMAIN_SYSVOL (2)
//
// * FRS_RSTYPE_DFS (3)
//
// * FRS_RSTYPE_OTHER (4)
//
// FRS_RSTYPE_ENTERPRISE_SYSVOL and FRS_RSTYPE_DOMAIN_SYSVOL are used for SYSVOL replication.
//
// FRS_RSTYPE_DFS is used for DFS replication. FRS_RSTYPE_OTHER is used only for testing.
//
// The SYSVOL replica set is an FRS Replica set that has all the DCs in the domain as
// its members. It is created by default when a new domain is created. Every DC that
// is added to the domain is automatically joined as a member of this replica set. The
// SYSVOL replica set is mainly responsible for replicating policy data between the
// domain controllers.
//
// FRS exposes two sets of RPC interfaces:
//
// * Communication Interface—Exposes functions to implement the FRS replication protocol.
//
// * Programming Interface—Exposes functions to implement administrative and monitoring
// tasks.
//
// Each file or folder is assigned a GUID when it is first added to a replica set. All
// replicas in the replica set use the same GUID to refer to the file or folder.<1>
package frs1

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
)

var (
	// import guard
	GoPackage = "frs1"
)
