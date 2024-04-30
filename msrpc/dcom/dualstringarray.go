package dcom

import (
	"encoding/json"
	"strconv"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
)

type StringBinding struct {
	TowerID     uint16 `json:"tower_id"`
	NetworkAddr string `json:"network_addr"`
}

func (o *StringBinding) String() string {

	switch int(o.TowerID) {
	case dcetypes.ProtocolTCP:
		return "ncacn_ip_tcp:" + o.NetworkAddr
	case dcetypes.ProtocolNamedPipe:
		return "ncacn_np:" + o.NetworkAddr
	case dcetypes.ProtocolHTTP:
		return "ncacn_http:" + o.NetworkAddr
	case dcetypes.ProtocolLRPC:
		return "ncacnlrpc:" + o.NetworkAddr
	}

	return "nca_unknown_" + strconv.Itoa(int(o.TowerID)) + ":" + o.NetworkAddr
}

type SecurityBinding struct {
	AuthnType     uint16 `json:"authn_type"`
	AuthzService  uint16 `json:"authz_service"`
	PrincipalName string `json:"principal_name"`
}

func (o *DualStringArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		StringBindings   []*StringBinding   `json:"string_bindings"`
		SecurityBindings []*SecurityBinding `json:"security_bindings"`
	}{
		o.GetStringBindings(),
		o.GetSecurityBindings(),
	})
}

func (o *DualStringArray) GetSecurityBindings() []*SecurityBinding {

	if o == nil {
		return nil
	}

	ret := []*SecurityBinding{}

	for buf := o.StringArray[o.SecurityOffset:]; len(buf) > 2; {
		i := 2
		for buf[i] != 0x0000 && i < len(buf) {
			i++
		}
		ret, buf = append(ret, &SecurityBinding{
			AuthnType:     buf[0],
			AuthzService:  buf[1],
			PrincipalName: string(utf16.Decode(buf[2:i])),
		}), buf[i+1:]
	}

	return ret
}

func (o *DualStringArray) Endpoints() []dcerpc.Option {

	eps := []dcerpc.Option{}

	for _, binding := range o.GetStringBindings() {
		eps = append(eps, dcerpc.WithEndpoint(binding.String()))
	}

	return eps
}

func (o *DualStringArray) GetStringBindings() []*StringBinding {

	if o == nil {
		return nil
	}

	ret := []*StringBinding{}

	for buf := o.StringArray[:o.SecurityOffset]; len(buf) > 1; {
		i := 1
		for buf[i] != 0x0000 && i < len(buf) {
			i++
		}
		ret, buf = append(ret, &StringBinding{
			TowerID:     buf[0],
			NetworkAddr: string(utf16.Decode(buf[1:i])),
		}), buf[i+1:]
	}

	return ret
}
