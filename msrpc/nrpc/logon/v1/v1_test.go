package logon

import (
	"context"
	"errors"
	"testing"

	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/ndr"
	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/netlogon"
)

// dcStub plays the domain controller of a secure channel: it verifies the authenticator of the client with its own copy of
// the credential, advances it as the DC does and answers with the return authenticator (also for a failed call).
type dcStub struct {
	LogonClient
	dc       *netlogon.SecureCredential
	status   error // the status of the call
	noReturn bool  // the DC rejects the authenticator: it does not advance its credential, it sends no return authenticator
	rejected int
}

func (s *dcStub) SAMLogonWithFlags(ctx context.Context, in *SAMLogonWithFlagsRequest, _ ...dcerpc.CallOption) (*SAMLogonWithFlagsResponse, error) {
	if s.noReturn {
		s.rejected++
		return &SAMLogonWithFlagsResponse{}, s.status
	}
	if err := s.dc.Verify(context.Background(), in.Authenticator.Timestamp, in.Authenticator.Credential.Data); err != nil {
		return nil, err // the client is out of step: STATUS_ACCESS_DENIED
	}
	ra, err := s.dc.Next(ctx, 1)
	if err != nil {
		return nil, err
	}
	return &SAMLogonWithFlagsResponse{ReturnAuthenticator: &Authenticator{Credential: &Credential{Data: ra}}}, s.status
}

// TestSecureChannelFailedCall: after a call that fails with a status, the client follows the credential of the DC.
func TestSecureChannelFailedCall(t *testing.T) {
	ctx := context.Background()
	errLogon := errors.New("STATUS_NO_SUCH_USER")
	for _, tc := range []struct {
		name  string
		calls []struct {
			status   error
			noReturn bool
		}
	}{
		{"success twice", []struct {
			status   error
			noReturn bool
		}{{nil, false}, {nil, false}}},
		{"failure, then success", []struct {
			status   error
			noReturn bool
		}{{errLogon, false}, {nil, false}}},
		{"failure twice, then success", []struct {
			status   error
			noReturn bool
		}{{errLogon, false}, {errLogon, false}, {nil, false}}},
		{"authenticator rejected (no return authenticator): the error is kept", []struct {
			status   error
			noReturn bool
		}{{errLogon, true}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cfg := netlogon.NewConfig()
			cfg.Credential = credential.NewFromNTHash("COMPUTER$", "00112233445566778899aabbccddeeff")
			cfg.ClientChallenge, cfg.ServerChallenge = []byte("12345678"), []byte("abcdefgh")
			client, err := netlogon.NewSecureCredential(ctx, cfg)
			if err != nil {
				t.Fatal(err)
			}
			dc, err := netlogon.NewSecureCredential(ctx, cfg)
			if err != nil {
				t.Fatal(err)
			}
			stub := &dcStub{dc: dc}
			cli := &xxx_SecureChannelClient{LogonClient: stub, sCred: client}
			for i, c := range tc.calls {
				stub.status, stub.noReturn = c.status, c.noReturn
				_, err := cli.SAMLogonWithFlags(ctx, &SAMLogonWithFlagsRequest{})
				if c.status == nil && err != nil {
					t.Fatalf("call %d: %v", i, err)
				}
				if c.status != nil && err == nil {
					t.Fatalf("call %d: the failure of the call is lost", i)
				}
			}
		})
	}
}

// TestValidationGenericRoundTrip: the DC answers the validation levels 4 (NetlogonValidationGenericInfo) and 5
// (NetlogonValidationGenericInfo2) with the NETLOGON_VALIDATION_GENERIC_INFO2 layout (EXP-760: level 4 failed to decode with
// "unread portion of 12 bytes", an empty payload: referent id, data length 0, null data pointer).
func TestValidationGenericRoundTrip(t *testing.T) {
	ctx := context.Background()
	for _, tc := range []struct {
		name  string
		level ValidationInfoClass
		in    *Validation
		data  []byte
	}{
		{"level 4, data", ValidationInfoClassGenericInfo, &Validation{Value: &Validation_Generic{ValidationGeneric: &ValidationGenericInfo2{DataLength: 3, ValidationData: []byte{1, 2, 3}}}}, []byte{1, 2, 3}},
		{"level 4, empty", ValidationInfoClassGenericInfo, &Validation{Value: &Validation_Generic{ValidationGeneric: &ValidationGenericInfo2{}}}, nil},
		{"level 5, data", ValidationInfoClassGenericInfo2, &Validation{Value: &Validation_Generic2{ValidationGeneric2: &ValidationGenericInfo2{DataLength: 3, ValidationData: []byte{1, 2, 3}}}}, []byte{1, 2, 3}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b, err := ndr.MarshalResponse(&xxx_SAMLogonWithFlagsOperation{ValidationLevel: tc.level, ValidationInformation: tc.in})
			if err != nil {
				t.Fatal(err)
			}
			out := &xxx_SAMLogonWithFlagsOperation{}
			if err := out.UnmarshalNDRResponse(ctx, ndr.NDR20(b)); err != nil {
				t.Fatal(err)
			}
			info, ok := out.ValidationInformation.GetValue().(*ValidationGenericInfo2)
			if !ok || info == nil {
				t.Fatalf("decoded %T, want *ValidationGenericInfo2", out.ValidationInformation.GetValue())
			}
			if string(info.ValidationData) != string(tc.data) {
				t.Fatalf("data %v, want %v", info.ValidationData, tc.data)
			}
		})
	}
}
