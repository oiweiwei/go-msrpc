package ad

import (
	"encoding/asn1"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
)

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
	Syntax string `json:"syntax"`
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
