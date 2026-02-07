package dcerpc

import (
	"context"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/oiweiwei/go-msrpc/ndr"
)

// SyntaxID ...
type SyntaxID struct {
	IfUUID         *uuid.UUID
	IfVersionMajor uint16
	IfVersionMinor uint16
}

func (v *SyntaxID) Size() int {
	return 16 + 2 + 2
}

func (v *SyntaxID) Is(other *SyntaxID) bool {
	if v == nil || other == nil {
		return false
	}
	return v.IfUUID.Equals(other.IfUUID) &&
		v.IfVersionMajor == other.IfVersionMajor &&
		v.IfVersionMinor == other.IfVersionMinor
}

// marshal function ...
func (v *SyntaxID) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.SetErr(v.IfUUID.Write(w))
	w.WriteData(v.IfVersionMajor)
	w.WriteData(v.IfVersionMinor)
	return w.Err()
}

// ReadFrom function ...
func (v *SyntaxID) ReadFrom(ctx context.Context, r ndr.Reader) error {
	v.IfUUID = &uuid.UUID{}
	r.SetErr(v.IfUUID.Read(r))
	r.ReadData(&v.IfVersionMajor)
	r.ReadData(&v.IfVersionMinor)
	return r.Err()
}

type ContextList struct {
	NContextElem uint8
	_            uint8  // reserved.
	_            uint16 // reserved2.
	ContextElem  []*Context
}

func (c *ContextList) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(c.NContextElem)
	w.WriteData(uint8(0))  // reserved.
	w.WriteData(uint16(0)) // reserved2.
	for i1 := 0; i1 < int(c.NContextElem); i1++ {
		if c.ContextElem[i1] != nil {
			c.ContextElem[i1].WriteTo(ctx, w)
		} else {
			(&Context{}).WriteTo(ctx, w)
		}
	}
	return w.Err()
}

func (c *ContextList) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&c.NContextElem)
	_reserved := uint8(0)
	r.ReadData(&_reserved)
	_reserved2 := uint16(0)
	r.ReadData(&_reserved2)
	c.ContextElem = make([]*Context, c.NContextElem)
	for i1 := 0; i1 < int(c.NContextElem); i1++ {
		c.ContextElem[i1] = &Context{}
		c.ContextElem[i1].ReadFrom(ctx, r)
	}
	return r.Err()
}

// The ContextElement struct represents one element in a presentation context
// list.
type Context struct {
	ContextID        uint16
	_                uint8 // n_transfer_syn.
	_                uint8 // reserved.
	AbstractSyntax   *SyntaxID
	TransferSyntaxes []*SyntaxID
}

func (c *Context) Size() int {
	size := 4 + c.AbstractSyntax.Size()
	for i := range c.TransferSyntaxes {
		size += c.TransferSyntaxes[i].Size()
	}
	return size
}

// marshal function ...
func (c *Context) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(c.ContextID)
	w.WriteData((uint8)(len(c.TransferSyntaxes)))
	w.WriteData((uint8)(0))
	if c.AbstractSyntax != nil {
		c.AbstractSyntax.WriteTo(ctx, w)
	} else {
		(&SyntaxID{}).WriteTo(ctx, w)
	}
	for i := range c.TransferSyntaxes {
		if c.TransferSyntaxes[i] != nil {
			c.TransferSyntaxes[i].WriteTo(ctx, w)
		} else {
			(&SyntaxID{}).WriteTo(ctx, w)
		}
	}
	return w.Err()
}

// ReadFrom function ...
func (c *Context) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&c.ContextID)
	_transferSyntaxes_len := uint8(0)
	r.ReadData(&_transferSyntaxes_len)
	_reserved := uint8(0)
	r.ReadData(&_reserved)
	c.AbstractSyntax = &SyntaxID{}
	c.AbstractSyntax.ReadFrom(ctx, r)
	c.TransferSyntaxes = make([]*SyntaxID, _transferSyntaxes_len)
	for i := range c.TransferSyntaxes {
		c.TransferSyntaxes[i] = &SyntaxID{}
		c.TransferSyntaxes[i].ReadFrom(ctx, r)
	}
	return r.Err()
}

// DefResults represents the results of a presentation context negotiation.
type DefResult uint16

const (
	Acceptance        DefResult = 0x0000
	UserRejection     DefResult = 0x0001
	ProviderRejection DefResult = 0x0002
	NegotiateAck      DefResult = 0x0003
)

// ProviderReason represents the reasons for rejection of a context element.
type ProviderReason uint16

const (
	ReasonNotSpecified                   ProviderReason = 0x0000
	AbstractSyntaxNotSupported           ProviderReason = 0x0001
	ProposedTransferSyntaxesNotSupported ProviderReason = 0x0002
	LocalLimitExceeded                   ProviderReason = 0x0003
	ProtocolVersionNotSupported          ProviderReason = 0x0004
	AuthTypeNotRecognized                ProviderReason = 0x0008
	InvalidChecksum                      ProviderReason = 0x0009

	// bind time feature negotiation flags.
	SecurityContextMultiplexing ProviderReason = 0x0001
	KeepConnOpenOnOrphaned      ProviderReason = 0x0002
)

// The Result struct represents one element in the results of the context
// negotiation.
type Result struct {
	DefResult      DefResult
	ProviderReason ProviderReason
	TransferSyntax *SyntaxID
}

func (result *Result) HasError() bool {
	return result.DefResult != Acceptance && result.DefResult != NegotiateAck
}

func (result *Result) Error() string {

	def := ""
	switch result.DefResult {
	case UserRejection:
		def = "user_rejection"
	case ProviderRejection:
		def = "provider_rejection"
	}

	switch result.ProviderReason {
	case AbstractSyntaxNotSupported:
		def += ": abstract syntax not supported"
	case ProposedTransferSyntaxesNotSupported:
		def += ": proposed transfer syntax not supported"
	case LocalLimitExceeded:
		def += ": local limit exceeded"
	default:
		def += ": unknown reason"
	}

	return def
}

func (result *Result) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData((uint16)(result.DefResult))
	w.WriteData((uint16)(result.ProviderReason))
	result.TransferSyntax.WriteTo(ctx, w)
	return w.Err()
}

func (result *Result) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData((*uint16)(&result.DefResult))
	r.ReadData((*uint16)(&result.ProviderReason))
	result.TransferSyntax = &SyntaxID{}
	result.TransferSyntax.ReadFrom(ctx, r)
	return r.Err()
}

type Feature Result

func (f *Feature) SecurityContextMultiplexing() bool {
	return f != nil && f.ProviderReason&SecurityContextMultiplexing != 0
}

func (f *Feature) KeepConnOpenOnOrphaned() bool {
	return f != nil && f.ProviderReason&KeepConnOpenOnOrphaned != 0
}

type ResultList struct {
	NResults uint8
	_        uint8
	_        uint16
	Results  []*Result
}

func (rl *ResultList) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(rl.NResults)
	w.WriteData(uint8(0))
	w.WriteData(uint16(0))
	for i1 := 0; i1 < int(rl.NResults); i1++ {
		if rl.Results[i1] != nil {
			rl.Results[i1].WriteTo(ctx, w)
		} else {
			(&Result{}).WriteTo(ctx, w)
		}
	}
	return w.Err()
}

func (rl *ResultList) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&rl.NResults)
	_reserved := uint8(0)
	r.ReadData(&_reserved)
	_reserved2 := uint16(0)
	r.ReadData(&_reserved2)
	rl.Results = make([]*Result, rl.NResults)
	for i1 := 0; i1 < int(rl.NResults); i1++ {
		rl.Results[i1] = &Result{}
		rl.Results[i1].ReadFrom(ctx, r)
	}
	return r.Err()
}

// Version structure specifies the protocol version.
type Version struct {
	Major uint8
	Minor uint8
}

func (version *Version) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(version.Major)
	w.WriteData(version.Minor)
	return w.Err()
}

func (version *Version) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&version.Major)
	r.ReadData(&version.Minor)
	return r.Err()
}

type VersionList struct {
	NProtocols uint8
	Versions   []*Version
}

func (vl *VersionList) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(vl.NProtocols)
	for i1 := 0; i1 < int(vl.NProtocols); i1++ {
		vl.Versions[i1].WriteTo(ctx, w)
	}
	return w.Err()
}

func (vl *VersionList) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&vl.NProtocols)
	vl.Versions = make([]*Version, vl.NProtocols)
	for i1 := range vl.Versions {
		vl.Versions[i1] = &Version{}
		vl.Versions[i1].ReadFrom(ctx, r)
	}
	return r.Err()
}

type PortAny struct {
	PortSpec string
}

func (pa *PortAny) WriteTo(ctx context.Context, w ndr.Writer) error {
	if len(pa.PortSpec) != 0 {
		w.WriteData((uint16)(len(pa.PortSpec) + 1))
		w.Write(append([]byte(pa.PortSpec), 0))
	} else {
		w.WriteData((uint16)(0))
	}
	return w.Err()
}

func (pa *PortAny) ReadFrom(ctx context.Context, r ndr.Reader) error {
	_portSpec_len := uint16(0)
	r.ReadData(&_portSpec_len)
	if _portSpec_len > 0 {
		b := make([]byte, _portSpec_len)
		r.Read(b)
		pa.PortSpec = string(b[:len(b)-1])
	}
	return r.Err()
}
