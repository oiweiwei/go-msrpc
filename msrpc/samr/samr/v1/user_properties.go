package samr

import (
	"context"
)

func (o *UserProperties) AfterUnmarshalNDR(ctx context.Context) error {

	if o == nil {
		return nil
	}

	if o.UserPropertiesList != nil {
		o.UserProperties = o.UserPropertiesList.UserProperties
		o.PropertyCount = o.UserPropertiesList.PropertyCount
	}

	for i1 := range o.UserProperties {
		switch v := o.UserProperties[i1].PropertyValue.GetValue().(type) {
		case *KerberosStoredCredential:
			for _, cred := range v.Credentials {
				cred.KeyData = o.UserProperties[i1].PropertyValueRaw[cred.KeyOffset:][:cred.KeyLength]
			}
			for _, cred := range v.OldCredentials {
				cred.KeyData = o.UserProperties[i1].PropertyValueRaw[cred.KeyOffset:][:cred.KeyLength]
			}
		case *KerberosStoredCredentialNew:
			for _, cred := range v.Credentials {
				cred.KeyData = o.UserProperties[i1].PropertyValueRaw[cred.KeyOffset:][:cred.KeyLength]
			}
			for _, cred := range v.ServiceCredentials {
				cred.KeyData = o.UserProperties[i1].PropertyValueRaw[cred.KeyOffset:][:cred.KeyLength]
			}
			for _, cred := range v.OldCredentials {
				cred.KeyData = o.UserProperties[i1].PropertyValueRaw[cred.KeyOffset:][:cred.KeyLength]
			}
			for _, cred := range v.OlderCredentials {
				cred.KeyData = o.UserProperties[i1].PropertyValueRaw[cred.KeyOffset:][:cred.KeyLength]
			}
		}
	}

	return nil
}
