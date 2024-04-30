// The cmrp package implements the CMRP client protocol.
//
// # Introduction
//
// The Failover Cluster: Management API (ClusAPI) Protocol is a remote procedure call
// (RPC)–based protocol that is used for remotely managing a failover cluster. Managing
// a failover cluster involves managing the data that represents the configuration of
// the cluster, the entities that constitute the cluster, and the applications and services
// that run in a cluster. For example, the ClusAPI Protocol is used to determine the
// version of the cluster and to read and write data in the cluster registry. This protocol
// is also used to determine whether a computer system is currently acting as a member
// of a failover cluster. Lastly, the ClusAPI Protocol is used to control and query
// a service or application that is hosted by a failover cluster.
//
// # Overview
//
// The ClusAPI Protocol is used to remotely manage a cluster. For example, this protocol
// can be used for the following purposes:
//
// * Determining whether a computer system is configured or active as a node ( 694e5e7a-5833-4f3d-b47e-323ee1d452c2#gt_762051d8-4fdc-437e-af9d-3f4da77c3c7d
// ) in a cluster.
//
// * Querying the configuration of the cluster; for example, cluster and node versions.
//
// * Storing data in and retrieving data from the cluster registry ( 694e5e7a-5833-4f3d-b47e-323ee1d452c2#gt_13de67f6-ac9d-491f-8dfb-12601a8b0838
// ).
//
// * Querying the configuration of applications and services that are hosted by the
// cluster.
//
// * Starting and stopping applications and services that are hosted by the cluster.
//
// An implementation can include methods that are executed using implementation-specific
// methods between servers. These include adding nodes to a cluster, changing the configuration
// of the cluster (for example, quorum policies or cluster version), restoring the cluster,
// and configuring applications and services to be hosted by a cluster. Such methods
// are specific to cluster-server implementations.
//
// A cluster is composed of computer systems that are called nodes. Before a computer
// can participate in a cluster as a node, it is configured as a cluster node. A node
// can be configured as a member of only one cluster at a time. After it is configured,
// a node can actively participate in its cluster.
//
// The nodes of a cluster are interconnected by one or more cluster networks and their
// corresponding cluster network interfaces. A cluster network represents a distinct
// communication path between a set of nodes and typically represents a subnet in the
// underlying network infrastructure. A cluster network interface is an instance of
// a connection point on a cluster network and is associated with a specific node. Thus,
// a given cluster network has a set of interfaces that defines the set of nodes that
// are reachable on that cluster network.
//
// Applications and services that are hosted by a cluster are represented as cluster
// resources. A resource can be started and stopped, consequently starting and stopping
// the application or service that the resource represents. Resources are contained
// in logical units called groups. A resource can be configured to depend on other resources
// in the same group. Resources are started and stopped in dependency order. A group
// is owned by one cluster node at a time, and a client can request that a group be
// moved from one node to another node. Moving a group in this manner stops all the
// resources on one node and starts them on the other node.
//
// Each resource is of one resource type. The resource type codifies how the resource
// is hosted by the cluster; for example, the semantics of starting or stopping it.
// Resource types that codify similar functionality can be grouped into a resource class.
// For example, resources that represent data storage devices can be grouped into a
// storage class even if they are of different resource types. The configuration of
// an object includes its common and private properties if such properties are part
// of the object's configuration and have been defined.
//
// Applications are made aware of changes in both the non-volatile and volatile cluster
// state through a notification port. The application can subscribe to a variety of
// events, such as the creation and deletion of objects and changes in object state
// and property values.
//
// The cluster registry is organized in a hierarchical tree structure that consists
// of keys and values. The cluster registry is rooted at a single key. Each object is
// associated with a key in the cluster registry, and the object's properties are stored
// under this key.
//
// In a typical ClusAPI Protocol session, the client connects to the server and requests
// to open a cluster object on the server. If the server accepts the request, it responds
// with an RPC context handle that refers to the cluster object. The client uses this
// RPC context handle to operate on that cluster object. Typically, the client then
// sends another request to the server and specifies the type of operation to perform
// and any specific parameters that are associated with that operation. If the server
// accepts this request, it attempts to query or change the state of the cluster object
// based on the request and responds to the client with the result of the operation.
// After the client is finished operating on the server cluster object, it terminates
// the protocol by sending a request to close the RPC context handle.
//
// The ClusAPI Protocol is an RPC-based protocol. For every method that the server receives,
// it executes the method against the current server configuration and cluster state.
// The server maintains client state information, and in some cases, protocol methods
// are executed in a particular order.
package cmrp

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
	GoPackage = "cmrp"
)
