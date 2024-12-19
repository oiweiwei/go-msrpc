package dcerpc

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
)

// The protocol sequence.
type ProtocolSequence int

func (p ProtocolSequence) String() string {

	switch p {
	case ProtocolSequenceIPTCP:
		return "ncacn_ip_tcp"
	case ProtocolSequenceNamedPipe:
		return "ncacn_np"
	case ProtocolSequenceHTTP:
		return "ncacn_http"
	case ProtocolSequenceLRPC:
		return "ncacnlrpc"
	}

	return "unknown"
}

type Binding struct {
	// The syntax identifier for the binding.
	SyntaxID SyntaxID
	// The transfer syntax.
	TransferSyntaxID SyntaxID
	// The string binding.
	StringBinding StringBinding
}

// The UUID convert to string function type.
type ConvertUUID func(*uuid.UUID) string

// The default UUID covert function.
var DefaultConvertUUID = ConvertUUID(func(u *uuid.UUID) string { return u.String() })

// URL function returns the URL for the Binding.
// Use UUID converter to change the UUID to human-readable string.
func (s Binding) URL(uConvs ...ConvertUUID) *url.URL {
	return s.StringBinding.URLWithSyntax(s.SyntaxID, uConvs...)
}

var (
	// The ncacn_ip_tcp identifies TCP/IP as the protocol
	// family for the endpoint.
	// Example: object@ncacn_ip_tcp:server-name[port-name]
	// URL: tcp://username:password@server-name:port/object/interface/version
	// Example URL: tcp://guest:guest@127.0.0.1:54212/00000000-0000-0000-0000000000000000/12345678-1234-abcd-ef0001234567cffb/v1.0
	ProtocolSequenceIPTCP ProtocolSequence = 1
	// The ncadg_ip_udp identifies UDP as the protocol
	// family for the endpoint.
	// Example: object@ncadg_ip_udp:server-name[port-name]
	// URL (?): udp://username:password@server-name:port/object/interface/version
	ProtocolSequenceIPUDP ProtocolSequence = 2
	// The ncacn_np identifies named pipes as the protocol family
	// for the endpoint.
	// Example: object@ncacn_np:server-name[\\pipe\\pipe-name]
	// URL (?): smb://username:password@server-name:smb_port/@computer/pipe/object/interface/version
	// Example URL: smb://guest:guest@127.0.0.1:445/@WIN_PC/winreg/00000000-0000-0000-0000000000000000/338cd001-2244-31f1-aaaa900038001003/v1.0
	ProtocolSequenceNamedPipe ProtocolSequence = 3
	// The ncalrpc identifies local interprocess communication
	// as the protocol family for the endpoint.
	// Example: object@ncalrpc:[port-name]
	// URL (?): alpc://port-name/object/interface/version
	ProtocolSequenceLRPC ProtocolSequence = 4
	// The ncacn_http identifies the Microsoft Internet Information
	// Server (IIS) as the protocol family for the endpoint.
	// Example: object@ncacn_http:rpc_server[endpoint]
	// URL (?): http://username:password@server-name:port/object/interface/version
	ProtocolSequenceHTTP ProtocolSequence = 5
)

// The string binding is an unsigned character string composed
// of strings that represent the binding object UUID, the RPC
// protocol sequence, the network address, and the endpoint and
// endpoint options.
type StringBinding struct {
	// The username and password for the binding.
	Username, Password string
	// The object UUID.
	ObjectUUID *uuid.UUID
	// The protocol sequence.
	ProtocolSequence ProtocolSequence
	// The IP/FQDN for the binding.
	NetworkAddress string
	// The NetBIOS computer name for named pipes.
	ComputerName string
	// The endpoint: port number for TCP/IP or HTTP,
	// named pipe name for SMB, or local IPC port for ALPC.
	Endpoint string
	// The extra endpoint options.
	Extra []string
}

func (s StringBinding) Complete() bool {
	return s.Endpoint != ""
}

func (s StringBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// URLWithSyntax function returns the URL with syntax interface identifier
// and version.
func (s StringBinding) URLWithSyntax(syntax SyntaxID, uConvs ...ConvertUUID) *url.URL {

	u := s.URL()

	uConv := DefaultConvertUUID
	if len(uConvs) > 0 && uConvs[0] != nil {
		uConv = uConvs[0]
	}

	u.Path, _ = url.JoinPath(u.Path, uConv(syntax.IfUUID), fmt.Sprintf("v%d.%d",
		syntax.IfVersionMajor, syntax.IfVersionMinor))

	return u
}

func afterSym(s string, sym string) string { return s[strings.LastIndex(s, sym)+1:] }

func afterSlash(s string) string { return afterSym(s, "\\") }

func (s StringBinding) ShareName() string {
	if s.ComputerName != "" {
		return "\\\\" + strings.TrimLeft(s.ComputerName, "\\") + "\\IPC$"
	}
	return "IPC$"
}

// MatchTarget function returns true if the target name of the string binding
// matches the provided target name n.
func (s StringBinding) MatchTarget(n string) bool {
	if target := s.TargetName(); target != "" && n != "" {
		return strings.Contains(strings.ToLower(n), strings.ToLower(target))
	}
	return false
}

func (s StringBinding) TargetName() string {
	if s.ComputerName != "" {
		return afterSlash(s.ComputerName)
	}
	return s.NetworkAddress
}

func (s StringBinding) NamedPipe() string {
	return afterSlash(s.Endpoint)
}

func (s StringBinding) URL() *url.URL {

	u := url.URL{}

	switch s.ProtocolSequence {
	case ProtocolSequenceIPTCP:
		u.Scheme, u.Host = "tcp", net.JoinHostPort(s.NetworkAddress, s.Endpoint)
	case ProtocolSequenceNamedPipe:
		if u.Scheme, u.Host = "smb", afterSlash(s.NetworkAddress); u.Host == "" {
			u.Host = afterSlash(s.ComputerName)
		} else if s.ComputerName != "" {
			u.Path, _ = url.JoinPath(u.Path, "@"+afterSlash(s.ComputerName))
		}
		u.Path, _ = url.JoinPath(u.Path, afterSlash(s.Endpoint))
	case ProtocolSequenceHTTP:
		u.Scheme, u.Host = "http", net.JoinHostPort(s.NetworkAddress, s.Endpoint)
	case ProtocolSequenceLRPC:
		u.Scheme, u.Host = "alpc", afterSlash(s.Endpoint)
	}

	if s.ObjectUUID != nil {
		u.Path, _ = url.JoinPath(u.Path, "object", s.ObjectUUID.String())
	}

	return &u
}

func (s StringBinding) String() string {

	var ret string

	if s.ObjectUUID != nil {
		ret += s.ObjectUUID.String() + "@"
	}

	if ret += s.ProtocolSequence.String() + ":" + s.NetworkAddress; s.NetworkAddress == "" {
		ret += s.ComputerName
	}

	if s.Endpoint != "" {
		ret += "[" + s.Endpoint + "]"
	}

	return ret
}

func ProtocolSequenceFromString(s string) ProtocolSequence {
	var p ProtocolSequence
	switch strings.ToLower(s) {
	case "ncacn_ip_tcp":
		p = ProtocolSequenceIPTCP
	case "ncacn_np":
		p = ProtocolSequenceNamedPipe
	case "ncacn_http":
		p = ProtocolSequenceHTTP
	case "ncacnlrpc":
		p = ProtocolSequenceLRPC
	}
	return p
}

func ParseSyntaxID(s string) (*SyntaxID, error) {

	var (
		id  SyntaxID
		err error
	)

	s, ver, ok := strings.Cut(s, "/")
	if ok {
		vMajStr, vMinStr, ok := strings.Cut(strings.TrimLeft(ver, "v"), ".")
		if ok {
			vMin, err := strconv.ParseUint(vMinStr, 10, 16)
			if err != nil {
				return nil, fmt.Errorf("parse syntax: version major: %w", err)
			}
			id.IfVersionMinor = uint16(vMin)
		}
		vMaj, err := strconv.ParseUint(vMajStr, 10, 16)
		if err != nil {
			return nil, fmt.Errorf("parse syntax: version minor: %w", err)
		}
		id.IfVersionMajor = uint16(vMaj)
	}

	if id.IfUUID, err = uuid.Parse(s); err != nil {
		return nil, fmt.Errorf("parse syntax id: %w", err)
	}

	return &id, nil
}

func ParseBindingURL(s string) (*Binding, error) {

	u, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("parse string binding url: %w", err)
	}

	var url Binding

	url.StringBinding.Username = u.User.Username()
	url.StringBinding.Password, _ = u.User.Password()

	for extra := range u.Query() {
		switch extra {
		case "username":
			url.StringBinding.Username = u.Query().Get(extra)
		case "password":
			url.StringBinding.Password = u.Query().Get(extra)
		default:
			if v := u.Query().Get(extra); v != "" {
				url.StringBinding.Extra = append(url.StringBinding.Extra, extra+"="+v)
			} else {
				url.StringBinding.Extra = append(url.StringBinding.Extra, extra)
			}
		}
	}

	p := strings.Split(u.Path, "/")

	switch u.Scheme {
	case "smb":
		url.StringBinding.ProtocolSequence = ProtocolSequenceNamedPipe
		if len(p) == 0 {
			return nil, fmt.Errorf("invalid binding url: named pipe is missing")
		}
		if !strings.HasPrefix(p[0], "@") {
			url.StringBinding.ComputerName = u.Hostname()
		} else {
			url.StringBinding.NetworkAddress, url.StringBinding.ComputerName, p = u.Hostname(), p[0], p[1:]
		}
		if len(p) == 0 {
			return nil, fmt.Errorf("invalid binding url: named pipe is missing")
		}
		url.StringBinding.Endpoint, p = p[0], p[1:]
	case "alpc":
		url.StringBinding.ProtocolSequence = ProtocolSequenceLRPC
		url.StringBinding.Endpoint = u.Hostname()
	case "http":
		url.StringBinding.ProtocolSequence = ProtocolSequenceHTTP
		url.StringBinding.NetworkAddress, url.StringBinding.Endpoint = u.Hostname(), u.Port()
	case "udp":
		url.StringBinding.ProtocolSequence = ProtocolSequenceIPUDP
		url.StringBinding.NetworkAddress, url.StringBinding.Endpoint = u.Hostname(), u.Port()
	case "tcp":
		url.StringBinding.ProtocolSequence = ProtocolSequenceIPTCP
		url.StringBinding.NetworkAddress, url.StringBinding.Endpoint = u.Hostname(), u.Port()
	}

	switch {
	case len(p) > 2:
		if url.StringBinding.ObjectUUID, err = uuid.Parse(p[0]); err != nil {
			return nil, err
		}
		p = p[1:]
		fallthrough
	case len(p) > 1:
		if url.SyntaxID.IfUUID, err = uuid.Parse(p[0]); err != nil {
			return nil, fmt.Errorf("invalid binding url: parse syntax: %w", err)
		}

		if _, err := fmt.Sscanf(p[1], "v%d.%d", &url.SyntaxID.IfVersionMajor, &url.SyntaxID.IfVersionMinor); err != nil {
			return nil, fmt.Errorf("invalid binding url: parse version: %w", err)
		}
	case len(p) > 0:
		if url.StringBinding.ObjectUUID, err = uuid.Parse(p[0]); err != nil {
			return nil, err
		}
	}

	return &url, nil
}

// ParseBinding function parses the string binding of format:
// [ 'ObjectUUID' '@' ] ProtocolSequence ':' NetworkAddress '[' Endpoint ']'
// [ NetworkAddress ] ':' Port
// [ 'Username' '%' 'Password' '@' ] ProtocolSequence ':' 'ComputerName' '[' Endpoint ']'
func ParseStringBinding(s string) (*StringBinding, error) {

	var (
		url StringBinding
		err error
	)

	if strings.Contains(s, "://") {
		// parse url.
		b, err := ParseBindingURL(s)
		if err != nil {
			return nil, fmt.Errorf("parse string binding: %w", err)
		}
		return &b.StringBinding, nil
	}

	before, s, ok := CutRight(s, "@")
	if ok {
		if strings.Contains(before, "%") { // XXX: special case to pass username%password
			url.Username, url.Password, _ = strings.Cut(before, "%")
		} else {
			if url.ObjectUUID, err = uuid.Parse(before); err != nil {
				return nil, fmt.Errorf("parse string binding: %v", err)
			}
		}
	} else {
		s = before
	}

	if before, s, ok = strings.Cut(s, ":"); !ok {

		if url.Username != "" || url.Password != "" {
			// it's a username%password@host.
			return &StringBinding{
				Username:         url.Username,
				Password:         url.Password,
				NetworkAddress:   before,
				ProtocolSequence: ProtocolSequenceIPTCP,
			}, nil
		}

		return nil, fmt.Errorf("parse string binding: malformed binding url: protocol sequence is missing")
	}

	if url.ProtocolSequence = ProtocolSequenceFromString(before); url.ProtocolSequence == 0 {

		if _, err := strconv.ParseUint(s, 10, 16); err != nil || url.ObjectUUID != nil {
			// return error if s is not a port number, or we have object_uuid.
			return nil, fmt.Errorf("parse string binding: unknown protocol sequence")
		}

		// it's plain host:port. (assume tcp).

		return &StringBinding{
			Username:         url.Username,
			Password:         url.Password,
			NetworkAddress:   before,
			ProtocolSequence: ProtocolSequenceIPTCP,
			Endpoint:         s,
		}, nil
	}

	switch url.ProtocolSequence {
	case ProtocolSequenceNamedPipe:
		url.ComputerName, s, ok = strings.Cut(s, "[")
		// XXX: any other way to pass the computer ip inside the ncacn_np
		// string binding. (seems to be not needed, as server IP is outside
		// of the scope of the binding).
	default:
		url.NetworkAddress, s, ok = strings.Cut(s, "[")
	}

	if !ok {
		return &url, nil
	}

	if url.Endpoint, s, ok = strings.Cut(s, "]"); !ok || s != "" {
		return nil, fmt.Errorf("parse string binding: malformed binding url: endpoint is malformed")
	}

	if url.Endpoint != "" {
		// parse endpoint options.
		extras := strings.Split(url.Endpoint, ",")
		if url.Endpoint, url.Extra = extras[0], extras[1:]; url.Endpoint == "*" || url.Endpoint == "?" {
			// special case for named pipe.
			url.Endpoint, url.Extra = "", extras
		}
		switch url.ProtocolSequence {
		case ProtocolSequenceIPTCP:
			// parse the port number.
			if _, err := strconv.ParseUint(url.Endpoint, 10, 16); err != nil {
				url.Endpoint, url.Extra = "", extras
			}
		}
	}

	return &url, nil
}

func CutRight(s string, sym string) (string, string, bool) {
	i := strings.LastIndex(s, sym)
	if i == -1 {
		return s, "", false
	}
	return s[:i], s[i+1:], true
}
