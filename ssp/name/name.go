package name

import "strings"

// NameType represents the type of naming convention used.
type NameType int

const (
	UnknownNameType NameType = iota
	// DNSNameType represents a DNS name, e.g. host.domain.tld
	DNSNameType
	// NetBIOSNameType represents a NetBIOS name, e.g. DOMAIN\host
	NetBIOSNameType
)

// Name represents a computer name with optional domain and forest components.
type Name struct {
	// Type indicates the naming convention used (DNS or NetBIOS).
	Type NameType
	// Name is the primary identifier, either the hostname or NetBIOS name.
	Name string
	// Domain is the domain part of the name, if applicable.
	Domain string
	// Forest is an optional higher-level grouping, often used in Active Directory contexts.
	Forest string
}

type NameOption func(*Name)

// Forest sets the Forest field of the Name struct.
func Forest(forest string) NameOption {
	return func(n *Name) {
		n.Forest = forest
	}
}

// Domain sets the Domain field of the Name struct.
func Domain(domain string) NameOption {
	return func(n *Name) {
		n.Domain = domain
	}
}

// NetBIOS creates a Name instance using the NetBIOS naming convention.
func NetBIOS(name string, o ...NameOption) *Name {

	n := &Name{Type: NetBIOSNameType, Name: name}

	if n.Name = strings.TrimLeft(n.Name, "\\"); strings.Contains(n.Name, "\\") {
		parts := strings.SplitN(n.Name, "\\", 2)
		n.Name = parts[1]
		n.Domain = parts[0]
	}

	for _, opt := range o {
		opt(n)
	}

	return n
}

// DNS creates a Name instance using the DNS naming convention.
func DNS(name string, o ...NameOption) *Name {

	n := &Name{Type: DNSNameType, Name: name}

	if strings.Contains(n.Name, ".") {
		parts := strings.SplitN(n.Name, ".", 2)
		n.Name = parts[0]
		n.Domain = parts[1]
	}

	for _, opt := range o {
		opt(n)
	}

	return n
}
