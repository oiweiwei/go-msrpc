package epm

import (
	"context"
	"fmt"

	"github.com/oiweiwei/go-msrpc/dcerpc"

	"github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	"github.com/oiweiwei/go-msrpc/msrpc/well_known"
)

func EndpointMapper(ctx context.Context, addr string, opts ...dcerpc.Option) dcerpc.Option {
	return dcerpc.WithEndpointMapper(NewMapper(ctx, addr, opts...))
}

func NewMapper(ctx context.Context, addr string, opts ...dcerpc.Option) dcerpc.EndpointMapper {

	m := &Mapper{}

	binding, err := dcerpc.ParseStringBinding(addr)
	if err != nil {
		binding = &dcerpc.StringBinding{}
	} else {
		if addr = binding.NetworkAddress; addr == "" {
			addr = binding.ComputerName
		}
	}

	bindings, err := m.WellKnown.Map(ctx, &dcerpc.Binding{StringBinding: *binding, SyntaxID: *EpmSyntaxV3_0})
	if err != nil {
		return m.WithErr(err)
	}

	o, err := dcerpc.ParseOptions(ctx, opts...)
	if err == dcerpc.ErrNoSecurityContext {
		opts = append(opts, dcerpc.WithInsecure())
	}

	var conn dcerpc.Conn

	dialOpts := []dcerpc.Option{}

	for _, opt := range opts {
		if opt, ok := opt.(dcerpc.ConnectOption); ok {
			dialOpts = append(dialOpts, opt)
		}
	}

	for _, binding := range bindings {
		binding.NetworkAddress = addr
		o.Logger.Debug().Msgf("endpoint mapper: dialing %s", binding)
		if conn, err = dcerpc.Dial(ctx, binding.String(), dialOpts...); err != nil {
			o.Logger.Error().Err(err).Msgf("endpoint mapper: dial %s", binding)
			continue
		}
		o.Logger.Debug().Msgf("endpoint mapper: dialing %s done", binding)
		break
	}

	if conn == nil {
		return m.WithErr(err)
	}

	o.Logger.Debug().Msgf("endpoint mapper: dialing client %s", binding)

	if m.cli, err = NewEpmClient(ctx, conn, opts...); err != nil {
		return m.WithErr(err)
	}

	return m
}

type Mapper struct {
	WellKnown well_known.WellKnownMapper
	cli       EpmClient
	err       error
}

func (m *Mapper) WithErr(err error) *Mapper {
	if m.err == nil {
		m.err = err
	}
	return m
}

func (m *Mapper) Map(ctx context.Context, in *dcerpc.Binding) ([]dcerpc.StringBinding, error) {

	var ret []dcerpc.StringBinding

	if m.err != nil {
		bindings, _ := m.WellKnown.Map(ctx, in)
		for _, binding := range bindings {
			ret = append(ret, binding)
		}
		return ret, fmt.Errorf("endpoint mapper: falling back to well known endpoints: %w", m.err)
	}

	resp, err := m.cli.Lookup(ctx, &LookupRequest{
		InquiryType: 0x00000001, // RPC_C_EP_MATCH_BY_IF
		VersOption:  0x00000003, // RPC_C_VERS_EXACT
		InterfaceID: &dcetypes.InterfaceID{
			UUID:      dtyp.GUIDFromUUID(in.SyntaxID.IfUUID),
			VersMajor: in.SyntaxID.IfVersionMajor,
			VersMinor: in.SyntaxID.IfVersionMinor,
		},
		MaxEntries: 500,
	})

	if err != nil {
		return nil, fmt.Errorf("endpoint mapper: lookup: %w", err)
	}

	for _, entry := range resp.Entries {
		binding := entry.Tower.Binding().StringBinding
		if in.StringBinding.ProtocolSequence == 0 || in.StringBinding.ProtocolSequence == binding.ProtocolSequence {
			ret = append(ret, entry.Tower.Binding().StringBinding)
		}
	}

	bindings, err := m.WellKnown.Map(ctx, in)
	if err != nil {
		return nil, err
	}

	for _, binding := range bindings {
		ret = append(ret, binding)
	}

	return ret, nil
}
