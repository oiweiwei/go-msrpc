package smb2

import (
	"fmt"

	"github.com/hirochachacha/go-smb2"
	smb2_fork "github.com/oiweiwei/go-smb2.fork"
)

func CompatDialer(d any) (*smb2_fork.Dialer, error) {

	if d == nil {
		return nil, nil
	}

	switch d := d.(type) {
	case *smb2.Dialer:
		d2 := &smb2_fork.Dialer{
			MaxCreditBalance: d.MaxCreditBalance,
			Negotiator: smb2_fork.Negotiator{
				RequireMessageSigning: d.Negotiator.RequireMessageSigning,
				ClientGuid:            d.Negotiator.ClientGuid,
				SpecifiedDialect:      d.Negotiator.SpecifiedDialect,
			},
		}

		if i, ok := d.Initiator.(*smb2.NTLMInitiator); ok {
			d2.Initiator = &smb2_fork.NTLMInitiator{
				User:        i.User,
				Password:    i.Password,
				Hash:        i.Hash,
				Domain:      i.Domain,
				Workstation: i.Workstation,
				TargetSPN:   i.TargetSPN,
			}
		}
		return d2, nil
	case *smb2_fork.Dialer:
		return d, nil
	default:
		return nil, fmt.Errorf("unknown dialer type: %T", d)
	}

}
