//go:generate go run gen.go -o well_known_endpoints.go
package well_known

import (
	"context"
	"strings"

	"github.com/oiweiwei/go-msrpc/dcerpc"
)

// The WellKnownMapper represents the static endpoint mapper.
type WellKnownMapper struct{}

func EndpointMapper() dcerpc.Option {
	return dcerpc.WithEndpointMapper(&WellKnownMapper{})
}

// Map function maps the provided syntax identifier to the string binding.
func (m *WellKnownMapper) Map(ctx context.Context, in *dcerpc.Binding) ([]dcerpc.StringBinding, error) {

	var ret []dcerpc.StringBinding

	for _, binding := range (*UUID)(in.SyntaxID.IfUUID).WellKnownEndpoint() {

		before, after, ok := strings.Cut(binding, ":")
		if !ok {
			continue
		}

		var binding dcerpc.StringBinding

		binding.Endpoint = after

		binding.ProtocolSequence = dcerpc.ProtocolSequenceFromString(before)

		if in.StringBinding.ProtocolSequence == 0 || in.StringBinding.ProtocolSequence == binding.ProtocolSequence {
			ret = append(ret, binding)
		}
	}

	return ret, nil
}

var (
	// XXX: interface guard.
	_ dcerpc.EndpointMapper = (*WellKnownMapper)(nil)
)
