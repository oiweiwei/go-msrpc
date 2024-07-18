package ad

import (
	"encoding/asn1"
	"encoding/json"
	"fmt"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
)

type Converter interface {
	Size() int
	Convert([]byte, ...any) (any, error)
}

type Syntax struct {
	Name      string
	Converter Converter
}

func (s Syntax) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Name)
}

// Attr is an Active Directory attribute.
type Attr struct {
	// CN is the common name of the attribute.
	CN string `json:"cn"`
	// LDAPName is the LDAP display name of the attribute.
	LDAPName string `json:"ldap_name"`
	// AttributeID is the ASN.1 object identifier of the attribute.
	AttributeID asn1.ObjectIdentifier `json:"attribute_id"`
	// The SystemID is the GUID of the attribute.
	SystemID *uuid.UUID `json:"system_id"`
	// Syntax is the syntax of the attribute.
	Syntax Syntax `json:"syntax"`
	// Description is the description of the attribute.
	Description string `json:"description"`
}

var oidToAttr = map[string]string{}

func init() {
	for k, attr := range attributes {
		oidToAttr[attr.AttributeID.String()] = k
	}
}

func LookupByOID(oid asn1.ObjectIdentifier) (Attr, bool) {
	n, ok := oidToAttr[oid.String()]
	if !ok {
		return Attr{}, false
	}
	return attributes[n], true
}

func ParseNameAndValue(oid asn1.ObjectIdentifier, b []byte, opts ...any) (string, any, error) {

	attr, ok := LookupByOID(oid)
	if !ok {
		return "", nil, fmt.Errorf("attribute not found")
	}

	if sz := attr.Syntax.Converter.Size(); sz > 0 {
		if len(b) < sz {
			return attr.LDAPName, nil, fmt.Errorf("value is too short")
		}
	}

	value, err := attr.Syntax.Converter.Convert(b, opts...)
	if err != nil {
		return attr.LDAPName, nil, err
	}

	return attr.LDAPName, value, nil
}
