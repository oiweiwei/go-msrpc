// The trp package implements the TRP client protocol.
//
// # Introduction
//
// The Microsoft Telephony Application Programming Interface (TAPI) enables implementation
// of communications applications ranging from voice mail to call centers with multiple
// agents and switches. The Microsoft Telephony programming model abstracts communications
// control from device control, freeing end-user applications and device manufacturers
// from the need to conform to the others' requirements. Using this model, an end-user
// or server application does not require detailed information about device control
// and the device does not need to be tailored to the application. Applications and
// devices can undergo innovation and change independently. Possible TAPI applications
// can include:
//
// * Basic voice calls on the public switched telephone network (PSTN).
//
// * Call center applications for tracking multiple agents.
//
// * Private branch exchange (PBX) control.
//
// * Interactive voice response (IVR) computing systems.
//
// * Voice mail.
//
// * Detailed phone device control.
//
// # Overview
//
// The Telephony Remote Protocol enables a client to control telephony devices on the
// server through TAPI, and manage or administer them. The server software can be modeled
// as:
//
// * TAPI service, which is independent of device specifics and depends on device-specific
// software for actual device control.
//
// * Telephony service provider (TSP), which is device-specific software (including
// the device driver software). For more information, see [MSDN-TAPI-SP] ( https://go.microsoft.com/fwlink/?LinkId=120037
// ).
//
// The TAPI service and the TSP can communicate with each other according to a well-defined
// interface, the Telephony Service Provider Interface (TSPI).
//
// An Automatic Call Distribution (ACD) server is a combination of hardware and software
// that classifies, queues, and distributes incoming calls to agents or outgoing calls
// to lines.
//
// The Server ACD application is a TAPI proxy application, which runs on the same server
// as the TSP. With a traditional ACD switch, the proxy application would interface
// to the switch's internal ACD and expose its operation. A software-based or "virtual"
// ACD proxy application would be fully responsible for the tracking of calls, queues,
// groups, and agents and would use the standard TAPI interfaces to control the switching
// hardware. Agent client applications will typically run on the individual agent's
// workstations and make use of the TAPI Remote Service Provider to communicate with
// the TAPISRV on the server machine, and hence the proxy application.
//
// The Agent object represents an agent that is capable of handling calls. This agent
// is usually a person but can be an interactive voice response (IVR) or some other
// combination of software and hardware. Agents are vital to a call center; they are
// responsible for receiving and processing incoming calls and at times, for making
// outgoing calls to customers or prospects.
//
// An Agent Handler represents software or hardware that is capable of passing calls
// to a group of agents. Typically, this is a proprietary switch that connects outside
// lines to telephones at agent stations.
//
// An Agent Session represents an agent who has logged on and is qualified to handle
// calls for a particular ACD Group. An agent session is a dynamically created object
// that relates an agent to an ACD group for which the group will provide service, and
// also to the address where they will receive calls (turret, station, phone, and so
// on). Applications can use the agent session object to track agent activity in a particular
// ACD group.
//
// An ACD group represents a class of calls that requires a particular type of handling.
// An ACD group services one or more queues. As incoming calls are classified, they
// are passed to queues that are associated with the relevant ACD group. A call coming
// off the queue is passed to an agent who has created an agent session object, indicating
// the agent is able to handle calls from that ACD group.
//
// The Queue object represents a point in the ACD system where calls are temporarily
// held pending action. Access to a queue object allows an application to read a variety
// of standard statistics that relate to queue usage; however, access does not give
// an application the ability to control calls on the queue. Only applications that
// have access to the associated addresses and lines are able to control the calls on
// the queue.
//
// Monitoring and control of ACD agent status on stations is supported through these
// functions: GetAgentCaps, GetAgentStatus, GetAgentGroupList, GetAgentActivityList,
// SetAgentGroup, SetAgentState, and SetAgentActivity.
//
// Architecturally, ACD functionality is implemented in a server-based application.
// The client functions mentioned above, rather than mapping to the telephony service
// provider, are conveyed to a server application that has registered (using an option
// of Open) as a handler for such functions.
//
// A line device represents a physical device such as a modem, voice board, fax board,
// or an Integrated Services Digital Network (ISDN) card that is connected to a network.
// Line devices support communications capabilities by allowing applications to send
// information to, or receive information from, a network. A line device contains a
// set of one or more homogeneous channels that can be used to establish calls. In Plain
// Old Telephone Service (POTS), exactly one channel exists on a line, and the channel
// is used exclusively for voice. Other technologies, such as ISDN, can support more
// than one channel on a single line.
//
// An address represents a location on a network. The address itself is a string that
// identifies a location on a network. In the case of a telephone network, the address
// is a telephone number. Each channel can have its own address, which means a line
// could have as many addresses as it has channels. The exact relationship between channels
// and addresses depends on the underlying TSP implementation.
//
// A client can obtain the number of addresses that are present on a line by using the
// GetDevCaps packet, which also provides information that is specific to the line device
// and common to all addresses on that line. Different addresses have different features,
// capabilities, and states. The client can access this information by sending the GetAddressCaps
// packet to the server.
//
// A phone device represents characteristics of the phone device hardware rather than
// the connection to the network itself. Thus, operations such as increasing or decreasing
// the volume of audio that is sent or received, changing the ring mode, and so on are
// carried out by using phone device operations.
//
// Many TAPI operations take a device ID or address ID parameter. The device ID can
// range from 0 to one less than the total number of devices that are reported by the
// corresponding Initialize packet. The address ID can range from 0 to one less than
// the number of addresses on that line device. The number of addresses on a line is
// obtained by sending the GetDevCaps packet for that line device.
//
// This protocol consists of two interfaces: the tapsrv interface and the remotesp interface.
//
// The tapsrv interface allows the client to send RPC packets to the server, causing
// TAPI operations to be executed on the server. The RPC packets in this specification
// are named for the specific TAPI operation that will be executed and are specified
// in section 2.2.
//
// TAPI operations can complete either synchronously or asynchronously.
//
// * Synchronous completion occurs when the requested TAPI operation is completely executed
// before the RPC function call returns to the client. This includes the case when the
// operation was not executed and an error is synchronously returned to the client.
//
// In Synchronous calls the client sends a TAPI32_MSG ( 2b49d075-4b32-4d60-8567-0c522b95f165
// ) packet through the ClientRequest ( c28c36d0-579b-4e20-a100-f4e71012d5f8 ) method
// with appropriate parameters in the packet. Depending on the request, the server fills
// the required values and sends back to client.
//
// For example, the client sends the GetDevCaps packet through the ClientRequest method
// to get the telephony capabilities of a specified line device. The GetDevCaps packet
// follows the structure of a TAPI32_MSG. The server fills the *Req_Func* field and
// *VarData* field of TAPI32_MSG with the result of the encapsulated telephony request
// and LINEDEVCAPS ( f236ea7a-c8a2-4681-b87c-9f0e07a01dc6 ).
//
// [Synchronous Completion](ms-trp_files/image001.png)
package trp

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
	GoPackage = "trp"
)
