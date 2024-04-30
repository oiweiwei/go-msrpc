package dcerpc

// security.go contains the definitions and implementation
// for the DCE/RPC security context.

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/oiweiwei/go-msrpc/ndr"
	"github.com/oiweiwei/go-msrpc/ssp"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

var (
	// the security context id generator.
	sContextID = new(atomic.Uint32)
)

func init() {
	val := uint32(0)
	binary.Read(rand.Reader, binary.LittleEndian, &val)
	sContextID.Store(val & 0x0f0f0f0f /* use half of values */)
}

// SecurityContextID function returns the unique identifier for
// the next security context.
func SecurityContextID() uint32 {
	return sContextID.Add(1)
}

// The security provider underlying protocol and implementation
// defines the number of legs and whether the number of legs is
// odd or even that are used in the token exchange process that
// builds a security context.
type Legs int

const (
	// Even number of legs.
	LegsEven = 2
	// Odd number of legs.
	LegsOdd = 3
	// Unknown.
	LegsUnknown = 0
)

type ImpersonationLevel int

const (

	// The client is anonymous to the server. The server process
	// can impersonate the client, but the impersonation token does
	// not contain any information about the client. This level is
	// only supported over the local interprocess communication transport.
	// All other transports silently promote this level to identify.
	ImpersonationLevelAnonymous = 1

	// The system default level. The server can obtain the client's
	// identity, and the server can impersonate the client to do ACL
	// checks.
	ImpersonationLevelIdentify = 2

	// The server can impersonate the client's security context while
	// acting on behalf of the client. The server can access local
	// resources as the client. If the server is local, it can access
	// network resources as the client. If the server is remote, it can
	// access only resources that are on the same computer as the server.
	ImpersonationLevelImpersonate = 3

	// The most powerful impersonation level. When this level is selected,
	// the server (whether local or remote) can impersonate the client's
	// security context while acting on behalf of the client. During
	// impersonation, the client's credentials (both local and network) can
	// be passed to any number of computers.
	ImpersonationLevelDelegate = 4
)

type AuthType uint8

const (
	// No authentication.
	AuthTypeNone AuthType = 0x00 // 0
	// Use the Microsoft Negotiate SSP. This SSP negotiates
	// between the use of the NTLM and Kerberos protocol Security
	// Support Providers (SSP).
	AuthTypeGSSNegotiate AuthType = 0x09 // 9
	// Use the Microsoft NT LAN Manager (NTLM) SSP.
	AuthTypeWinNT AuthType = 0x0A // 10
	// Use the Schannel SSP. This SSP supports Secure Socket Layer
	// (SSL), private communication technology (PCT), and transport
	// level security (TLS).
	AuthTypeGSSChannel AuthType = 0x0E // 14
	// Use the Microsoft Kerberos SSP.
	AuthTypeKerberos AuthType = 0x10 // 16
	// The Netlogon Secure Channel.
	AuthTypeNetLogon AuthType = 0x44 // 68
	// Use the default authentication service.
	AuthTypeDefault AuthType = 0xFF // 255
)

type SecurityTrailer struct {
	AuthType      AuthType
	AuthLevel     AuthLevel
	AuthPadLength uint8
	_             uint8
	AuthContextID uint32
}

func (st *SecurityTrailer) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData((*uint8)(&st.AuthType))
	r.ReadData((*uint8)(&st.AuthLevel))
	r.ReadData(&st.AuthPadLength)
	_reserved := uint8(0)
	r.ReadData(&_reserved)
	r.ReadData(&st.AuthContextID)
	return r.Err()
}

func (st *SecurityTrailer) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData((uint8)(st.AuthType))
	w.WriteData((uint8)(st.AuthLevel))
	w.WriteData(st.AuthPadLength)
	w.WriteData(uint8(0))
	w.WriteData(st.AuthContextID)
	return w.Err()
}

func (v AuthType) Legs() int {
	switch v {
	case AuthTypeNone, AuthTypeGSSNegotiate, AuthTypeGSSChannel:
		return LegsEven
	case AuthTypeWinNT, AuthTypeKerberos, AuthTypeNetLogon:
		return LegsOdd
	}

	return LegsUnknown
}

// The authentication-level constants represent authentication levels
// passed to various run-time functions. These levels are listed in
// order of increasing authentication. Each new level adds to the
// authentication provided by the previous level.
type AuthLevel uint8

const (
	// Uses the default authentication level for the specified
	// authentication service.
	AuthLevelDefault AuthLevel = 0x00
	// Performs no authentication.
	AuthLevelNone AuthLevel = 0x01
	// Authenticates only when the client establishes a relationship
	// with a server.
	AuthLevelConnect AuthLevel = 0x02
	// Authenticates only at the beginning of each remote procedure call
	// when the server receives the request. Does not apply to remote
	// procedure calls made using the connection-based protocol sequences.
	AuthLevelCall AuthLevel = 0x03
	// Authenticates only that all data received is from the expected client.
	// Does not validate the data itself.
	AuthLevelPkt AuthLevel = 0x04
	// Authenticates and verifies that none of the data transferred between
	// client and server has been modified.
	AuthLevelPktIntegrity AuthLevel = 0x05
	// Includes all previous levels, and ensures clear text data can only
	// be seen by the sender and the receiver. This involves encrypting
	// the argument value of each remote procedure call.
	AuthLevelPktPrivacy AuthLevel = 0x06
)

// The DCE/RPC Security Context.
type Security struct {
	// The security mutex.
	mu sync.Mutex
	// The security context identifier.
	id uint32
	// The flag that indicates whether the security context
	// is established.
	established bool
	// The GSSAPI options.
	opts []gssapi.Option
	// The GSSAPI security context.
	ctx context.Context
	// The impersonation level.
	Impersonation ImpersonationLevel
	// SupportHeaderSign.
	RequestHeaderSign PacketFlag
	// The security context authentication type.
	Type AuthType
	// The security context authentication level.
	Level AuthLevel
	// The flag that indicates whether the header signing is
	// supported.
	SignHeader bool
	// The flag that indicates whether the security context
	// multiplexing is supported.
	Multiplexing bool
	// The target name.
	TargetName string
}

// ID returns the security context identifier.
func (o *Security) ID() uint32 {
	return o.id
}

// NewSecurity function returns the new security context
// and saves the GSSAPI context storied in `ctx` variable.
func NewSecurity(ctx context.Context, opts ...gssapi.ContextOption) *Security {

	if _, ok := gssapi.GetAttribute(ctx, gssapi.AttributeRPCContext); ok || len(opts) > 0 {
		ctx = gssapi.NewSecurityContext(ctx, opts...)
	}

	return &Security{
		id:                SecurityContextID(),
		ctx:               ctx,
		Impersonation:     ImpersonationLevelImpersonate,
		Type:              MechanismToAuthType(ssp.MechanismTypeDefault(ctx)),
		RequestHeaderSign: PacketFlagSupportHeaderSign,
	}
}

// MechanismToAuthType function converts the mechanism OID to the
// DCE/RPC authentication type.
func MechanismToAuthType(mech gssapi.OID) AuthType {

	switch {
	case mech.Equal(ssp.MechanismTypeKRB5):
		return AuthTypeKerberos
	case mech.Equal(ssp.MechanismTypeSPNEGO):
		return AuthTypeGSSNegotiate
	case mech.Equal(ssp.MechanismTypeNetlogon):
		return AuthTypeNetLogon
	case mech.Equal(ssp.MechanismTypeNTLM):
		return AuthTypeWinNT
	}

	return AuthTypeWinNT
}

// SecurityTrailer function returns the DCE/RPC security trailer.
func (cc *Security) SecurityTrailer() SecurityTrailer {
	return SecurityTrailer{
		AuthType:      cc.Type,
		AuthLevel:     cc.Level,
		AuthContextID: cc.id,
	}
}

// Established function returns `true` if the security context was established
// and can be used for Wrap/Unwrap functions.
func (cc *Security) Established() bool {

	if cc == nil {
		return false
	}

	cc.mu.Lock()
	defer cc.mu.Unlock()

	return cc.established
}

// Init function inits the security context.
func (cc *Security) Init(ctx context.Context, b []byte) ([]byte, error) {

	if cc == nil {
		return []byte{}, nil
	}

	cc.mu.Lock()
	defer cc.mu.Unlock()

	if cc.established {
		return []byte{}, nil
	}

	if cc.Level == AuthLevelNone || cc.Type == AuthTypeNone {
		cc.established = true
		return []byte{}, nil
	}

	tok, err := gssapi.InitSecurityContext(cc.ctx, &gssapi.Token{Payload: b}, cc.options()...)
	if err != nil {
		return nil, fmt.Errorf("init security context: %w", err)
	}

	if cc.established = gssapi.IsComplete(cc.ctx); cc.established {
		gssapi.SetAttribute(cc.ctx, gssapi.AttributeRPCContext, cc) // save established security context.
	}

	return tok.Payload, nil
}

// AuthLength function returns the expected length for the authentication
// trailer.
func (cc *Security) AuthLength(ctx context.Context, pkt *Packet) int {

	if !cc.Established() {
		return 0
	}

	cc.mu.Lock()
	defer cc.mu.Unlock()

	return cc.authLength(ctx, pkt)
}

// authLength function returns the expected length for the authentication
// trailer.
func (cc *Security) authLength(ctx context.Context, pkt *Packet) int {

	switch cc.Level {
	case AuthLevelNone, AuthLevelConnect, AuthLevelDefault:
		return 0
	case AuthLevelCall:
		if !pkt.Header.PacketFlags.IsSet(PacketFlagFirstFrag) {
			return 0
		}
	}

	return -gssapi.WrapSizeLimit(cc.ctx, 0, cc.options()...)
}

// CanWrap function returns true if security context can be applied to the
// packet (context is established and packet is request or response).
func (cc *Security) CanWrap(ctx context.Context, pkt *Packet) bool {
	return cc.Established() && (pkt.Header.PacketType == PacketTypeRequest || pkt.Header.PacketType == PacketTypeResponse)
}

// Wrap function depending on the security level encrypts and computes the
// checksum for the packet.
func (cc *Security) Wrap(ctx context.Context, pkt *Packet) error {

	cc.mu.Lock()
	defer cc.mu.Unlock()

	switch raw, l := pkt.Bytes(), cc.authLength(ctx, pkt); cc.Level {

	case AuthLevelCall:

		if !pkt.Header.PacketFlags.IsSet(PacketFlagFirstFrag) {
			// call level applies only to the first frag.
			break
		}

		fallthrough

	case AuthLevelPktIntegrity, AuthLevelPkt:

		tokEx := &gssapi.MessageTokenEx{}

		caps := gssapi.Cap(0)
		if cc.SignHeader {
			caps |= gssapi.Integrity
		}

		tokEx.Payloads = append(tokEx.Payloads, &gssapi.PayloadEx{
			Payload:      pkt.HeaderBytes(),
			Capabilities: caps,
		})

		tokEx.Payloads = append(tokEx.Payloads, &gssapi.PayloadEx{
			Payload:      pkt.StubDataBytes(),
			Capabilities: caps | gssapi.Integrity,
		})

		tokEx.Payloads = append(tokEx.Payloads, &gssapi.PayloadEx{
			Payload:      pkt.SecurityTrailerBytes(),
			Capabilities: caps,
		})

		tokEx, err := gssapi.MakeSignatureEx(cc.ctx, tokEx, cc.options(SkipSealOpt|SkipSignOpt)...)
		if err != nil {
			return fmt.Errorf("make_signature_ex: %w", err)
		}

		copy(raw[len(raw)-l:], tokEx.Signature)

	case AuthLevelPktPrivacy:

		tokEx := &gssapi.MessageTokenEx{}

		caps := gssapi.Cap(0)
		if cc.SignHeader {
			caps |= gssapi.Integrity
		}

		tokEx.Payloads = append(tokEx.Payloads, &gssapi.PayloadEx{
			Payload:      pkt.HeaderBytes(),
			Capabilities: caps,
		})

		tokEx.Payloads = append(tokEx.Payloads, &gssapi.PayloadEx{
			Payload:      pkt.StubDataBytes(),
			Capabilities: caps | gssapi.Integrity | gssapi.Confidentiality,
		})

		tokEx.Payloads = append(tokEx.Payloads, &gssapi.PayloadEx{
			Payload:      pkt.SecurityTrailerBytes(),
			Capabilities: caps,
		})

		tokEx, err := gssapi.WrapEx(cc.ctx, tokEx, cc.options(SkipSealOpt|SkipSignOpt)...)
		if err != nil {
			return fmt.Errorf("wrap_ex: %w", err)
		}

		copy(raw[len(raw)-l:], tokEx.Signature)

	}

	return nil
}

// Unwrap function unwraps the packet and verifies the signature. The `afterLock` parameter
// is to ensure that any lock that must be released, will be released once the Security lock
// will be acquired.
func (cc *Security) Unwrap(ctx context.Context, pkt *Packet, afterLock func(context.Context)) error {

	cc.mu.Lock()
	defer cc.mu.Unlock()

	afterLock(ctx)

	switch cc.Level {

	case AuthLevelCall:

		if !pkt.Header.PacketFlags.IsSet(PacketFlagFirstFrag) {
			break
		}

		fallthrough

	case AuthLevelPktIntegrity, AuthLevelPkt:

		tokEx := &gssapi.MessageTokenEx{}

		caps := gssapi.Cap(0)
		if cc.SignHeader {
			caps |= gssapi.Integrity
		}

		tokEx.Payloads = append(tokEx.Payloads, &gssapi.PayloadEx{
			Payload:      pkt.HeaderBytes(),
			Capabilities: caps,
		})

		tokEx.Payloads = append(tokEx.Payloads, &gssapi.PayloadEx{
			Payload:      pkt.StubDataBytes(),
			Capabilities: caps | gssapi.Integrity | gssapi.Confidentiality,
		})

		tokEx.Payloads = append(tokEx.Payloads, &gssapi.PayloadEx{
			Payload:      pkt.SecurityTrailerBytes(),
			Capabilities: caps,
		})

		tokEx.Signature = pkt.AuthDataBytes()

		err := gssapi.VerifySignatureEx(cc.ctx, tokEx, cc.options(SkipSealOpt|SkipSignOpt)...)
		if err != nil {
			return fmt.Errorf("verify_signature_ex: %w", err)
		}

	case AuthLevelPktPrivacy:

		tokEx := &gssapi.MessageTokenEx{}

		caps := gssapi.Cap(0)
		if cc.SignHeader {
			caps |= gssapi.Integrity
		}

		tokEx.Payloads = append(tokEx.Payloads, &gssapi.PayloadEx{
			Payload:      pkt.HeaderBytes(),
			Capabilities: caps,
		})

		tokEx.Payloads = append(tokEx.Payloads, &gssapi.PayloadEx{
			Payload:      pkt.StubDataBytes(),
			Capabilities: caps | gssapi.Integrity | gssapi.Confidentiality,
		})

		tokEx.Payloads = append(tokEx.Payloads, &gssapi.PayloadEx{
			Payload:      pkt.SecurityTrailerBytes(),
			Capabilities: caps,
		})

		tokEx.Signature = pkt.AuthDataBytes()

		_, err := gssapi.UnwrapEx(cc.ctx, tokEx, cc.options(SkipSealOpt|SkipSignOpt)...)
		if err != nil {
			return fmt.Errorf("unwrap_ex: %w", err)
		}
	}

	return nil
}

type SecurityMask uint8

var (
	// SkipSealOpt option skips the Confidentiality option as it must
	// be defined on individual buffer level.
	SkipSealOpt SecurityMask = 0x01
	// SkipSignOpt option skips the Integrity option as it must be
	// defined on individual buffer level.
	SkipSignOpt SecurityMask = 0x02
)

// optSkipSeal function returns `true` if security mask contains SkipSealOpt.
func optSkipSeal(mask []SecurityMask) bool {
	return len(mask) > 0 && mask[0]&SkipSealOpt != 0
}

// optSkipSign function returns `true` if security mask contains SkipSignOpt.
func optSkipSign(mask []SecurityMask) bool {
	return len(mask) > 0 && mask[0]&SkipSignOpt != 0
}

// options function returns the set of SSP options derived from the level.
func (cc *Security) options(mask ...SecurityMask) []gssapi.Option {

	opts := []gssapi.Option{}

	if cc.TargetName != "" {
		opts = append(opts, gssapi.WithTargetName(cc.TargetName))
	}

	switch cc.Level {
	case AuthLevelPktPrivacy:
		if !optSkipSeal(mask) {
			opts = append(opts, gssapi.WithRequest(gssapi.Confidentiality))
		}
		fallthrough
	case AuthLevelPktIntegrity, AuthLevelPkt, AuthLevelCall, AuthLevelDefault:
		if !optSkipSign(mask) {
			opts = append(opts, gssapi.WithRequest(gssapi.Integrity|gssapi.ReplayDetection|gssapi.Sequencing))
		}
	}

	switch cc.Impersonation {
	case ImpersonationLevelDelegate:
		opts = append(opts, gssapi.WithRequest(gssapi.Delegation))
	case ImpersonationLevelIdentify:
		opts = append(opts, gssapi.WithRequest(gssapi.Identify))
	case ImpersonationLevelAnonymous:
		opts = append(opts, gssapi.WithRequest(gssapi.Anonymity))
	}

	switch cc.Type {
	case AuthTypeWinNT:
		opts = append(opts, gssapi.WithMechanismType(ssp.MechanismTypeNTLM))
	case AuthTypeKerberos:
		opts = append(opts, gssapi.WithMechanismType(ssp.MechanismTypeKRB5))
	case AuthTypeGSSNegotiate:
		opts = append(opts, gssapi.WithMechanismType(ssp.MechanismTypeSPNEGO))
	case AuthTypeNetLogon:
		opts = append(opts, gssapi.WithMechanismType(ssp.MechanismTypeNetlogon))
	case AuthTypeDefault:
		opts = append(opts, gssapi.WithMechanismType(ssp.MechanismTypeDefault(cc.ctx)))
	}

	return append(opts, cc.opts...)
}
