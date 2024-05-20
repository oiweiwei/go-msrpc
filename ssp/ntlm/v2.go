package ntlm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"strings"

	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/crypto"
	"github.com/oiweiwei/go-msrpc/ssp/ntlm/internal"
	"github.com/oiweiwei/go-msrpc/text/encoding/utf16le"
)

// The LMv2Response structure defines the NTLM v2 authentication
// LmChallengeResponse in the AuthenticateMessage. This response
// is used only when NTLM v2 authentication is configured.
type LMv2Response struct {
	// A 16-byte array of unsigned char that contains the client's
	// LM challenge-response. This is the portion of the LmChallengeResponse
	// field to which the HMAC_MD5 algorithm has been applied.
	// Specifically, Response corresponds to the result of applying the
	// HMAC_MD5 algorithm, using the key ResponseKeyLM, to a message consisting
	// of the concatenation of the ResponseKeyLM, ServerChallenge and
	// ClientChallenge.
	Response []byte
	// An 8-byte array of unsigned char that contains the client's
	// ClientChallenge.
	ChallengeFromClient []byte
}

func (m *LMv2Response) Marshal(ctx context.Context) ([]byte, error) {

	e := internal.NewCodec(ctx, nil)

	// response.
	e.WriteBytes(ctx, m.Response, 16)
	// challenge_from_client.
	e.WriteBytes(ctx, m.ChallengeFromClient, 8)

	return e.Bytes(ctx)
}

func (m *LMv2Response) Unmarshal(ctx context.Context, b []byte) error {

	e := internal.NewCodec(ctx, b)

	// response.
	e.ReadBytes(ctx, &m.Response, 16)
	// challenge_from_client.
	e.ReadBytes(ctx, &m.ChallengeFromClient, 8)

	return e.ReadAll(ctx)
}

// The NTLMv2ClientChallenge structure defines the client challenge
// in the AuthenticateMessage. This structure is used only when NTLM v2
// authentication is configured and is transported in the NTLMv2
// structure.
type NTLMv2ClientChallenge struct {
	// An 8-bit unsigned char that contains the current version of the
	// challenge response type. This field MUST be 0x01.
	RespType uint8
	// An 8-bit unsigned char that contains the maximum supported version
	// of the challenge response type. This field MUST be 0x01.
	HiRespType uint8
	// A 16-bit unsigned integer that SHOULD be 0x0000 and MUST be
	// ignored on receipt.
	_ uint16
	// A 32-bit unsigned integer that SHOULD be 0x00000000 and MUST be
	// ignored on receipt.
	_ uint32
	// A 64-bit unsigned integer that contains the current system time,
	// represented as the number of 100 nanosecond ticks elapsed since
	// midnight of January 1, 1601 (UTC).
	Timestamp Filetime
	// An 8-byte array of unsigned char that contains the client's
	// ClientChallenge.
	ChallengeFromClient []byte
	// A 32-bit unsigned integer that SHOULD be 0x00000000 and MUST
	// be ignored on receipt.
	_ uint32
	// A byte array that contains a sequence of AttrValue structures.
	// The sequence contains the server-naming context and is terminated
	// by an AttrValue structure with an AvId field of MsvAvEOL.
	AttrValues AttrValues
}

func (m *NTLMv2ClientChallenge) Marshal(ctx context.Context) ([]byte, error) {

	e := internal.NewCodec(ctx, nil)

	// resp_type.
	e.WriteData(ctx, m.RespType)
	// hi_resp_type.
	e.WriteData(ctx, m.HiRespType)
	// pad.
	e.WriteData(ctx, uint16(0))
	// pad.
	e.WriteData(ctx, uint32(0))
	// timestamp.
	e.WriteData(ctx, (uint64)(m.Timestamp))
	// challenge_from_client.
	e.WriteBytes(ctx, m.ChallengeFromClient, 8)
	// pad.
	e.WriteData(ctx, uint32(0))
	// attr_values.
	e.WriteData(ctx, &m.AttrValues)

	return e.Bytes(ctx)
}

func (m *NTLMv2ClientChallenge) Unmarshal(ctx context.Context, b []byte) error {

	e := internal.NewCodec(ctx, b)

	// resp_type.
	e.ReadData(ctx, &m.RespType)
	// hi_resp_type.
	e.ReadData(ctx, &m.HiRespType)
	// pad.
	_pad16 := uint16(0)
	e.ReadData(ctx, &_pad16)
	// pad.
	_pad32 := uint32(0)
	e.ReadData(ctx, &_pad32)
	// timestamp.
	e.ReadData(ctx, (*uint64)(&m.Timestamp))
	// challenge_from_client.
	e.ReadBytes(ctx, &m.ChallengeFromClient, 8)
	// pad.
	e.ReadData(ctx, &_pad32)
	// attr_values.
	e.ReadData(ctx, &m.AttrValues)

	return e.ReadAll(ctx)
}

// The NTLMv2Response structure defines the NTLMv2 authentication
// NtChallengeResponse in the AuthenticateMessage. This response
// is used only when NTLMv2 authentication is configured.
type NTLMv2Response struct {
	// A 16-byte array of unsigned char that contains the client's
	// NTChallengeResponse.
	Response []byte
	// A variable-length byte array, that contains the ClientChallenge.
	NTLMv2ClientChallenge *NTLMv2ClientChallenge
}

func (m *NTLMv2Response) Marshal(ctx context.Context) ([]byte, error) {

	e := internal.NewCodec(ctx, nil)

	// response.
	e.WriteBytes(ctx, m.Response, 16)
	// ntlmv2_client_challenge.
	e.WriteData(ctx, m.NTLMv2ClientChallenge)

	return e.Bytes(ctx)
}

func (m *NTLMv2Response) Unmarshal(ctx context.Context, b []byte) error {

	e := internal.NewCodec(ctx, b)

	// response.
	e.ReadBytes(ctx, &m.Response, 16)
	// ntlmv2_client_challenge.
	m.NTLMv2ClientChallenge = &NTLMv2ClientChallenge{}
	e.ReadData(ctx, m.NTLMv2ClientChallenge)

	return e.ReadAll(ctx)
}

type V2 struct {
	*Config
	*SecurityParameters
}

func (v2 *V2) WithConfig(ctx context.Context, config *Config) NTLMVersion {
	v2.Config = config.clone()
	return v2
}

var (
	// interface guard.
	_ NTLMVersion = (*V2)(nil)
)

// LMOWF function is a NT LAN Manager (LM) one-way function used to create
// a hash based on the user's password.
func (v2 *V2) LMOWF(ctx context.Context, cred Credential) ([]byte, error) {
	return v2.NTOWF(ctx, cred)
}

// NTOWF function is a NT LAN Manager (NT) one-way function used to create
// a hash based on user's password.
func (v2 *V2) NTOWF(ctx context.Context, cred Credential) ([]byte, error) {

	if cred == nil {
		return nil, nil
	}

	var (
		k   []byte
		err error
	)

	if cred, ok := cred.(credential.Password); ok {

		pass, err := utf16le.Encode(cred.Password())
		if err != nil {
			return nil, fmt.Errorf("v2: ntowf: encode password: %w", err)
		}

		k, err = crypto.MD4(pass)
		if err != nil {
			return nil, fmt.Errorf("v2: ntowf: md4 password: %w", err)
		}
	}

	if cred, ok := cred.(credential.NTHash); ok {
		k = cred.NTHash()
	}

	user, err := utf16le.Encode(strings.ToUpper(cred.UserName()) + cred.DomainName())
	if err != nil {
		return nil, fmt.Errorf("v2: ntowf: encode username: %w", err)
	}

	b, err := crypto.HMAC_MD5(k, user)
	if err != nil {
		return nil, fmt.Errorf("v2: ntowf: derive key: %w", err)
	}

	return b, nil
}

var (
	ErrLogonFailure = errors.New("logon failure")
)

// ChallengeResponse function computes the response for the server challenge message
// based on client challenge `nonce` parameter and set of negotiated flags.
func (v2 *V2) ChallengeResponse(ctx context.Context, cred Credential, c *ChallengeMessage, nonce []byte) (*ChallengeResponse, error) {

	if !c.TargetInfo.HasAttr(AttrNetBIOSDomainName) && !c.TargetInfo.HasAttr(AttrNetBIOSComputerName) &&
		(v2.Config.Integrity || v2.Config.Confidentiality) /* any of */ {
		return nil, fmt.Errorf("v2: %w", ErrLogonFailure)
	}

	if cred == nil || (cred.UserName() == "" && IsCredentialEmpty(cred)) {
		// anonymous case.
		return &ChallengeResponse{LM: []byte{0}, IsAnonymous: true}, nil
	}

	var (
		resp = new(ChallengeResponse)
		err  error
	)

	// force set mic.
	resp.RequestMIC = v2.RequestMIC

	// compute nt key.
	if resp.KeyNT, err = v2.NTOWF(ctx, cred); err != nil {
		return nil, err
	}

	// "compute" lm key.
	resp.KeyLM = resp.KeyNT

	respNT, respLM := new(NTLMv2Response), new(LMv2Response)

	respNT.NTLMv2ClientChallenge = &NTLMv2ClientChallenge{
		RespType:            0x01,
		HiRespType:          0x01,
		ChallengeFromClient: nonce,
	}

	attrs := AttrValues{}

	// copy server naming context.
	for k, v := range c.TargetInfo {
		attrs[k] = v
	}

	// provide hashed channel binding.
	if attrs[AttrChannelBindings] = new(Value); v2.ChannelBindings != nil {
		if attrs[AttrChannelBindings].ChannelBindings, err = crypto.MD5(v2.ChannelBindings); err != nil {
			return nil, fmt.Errorf("v2: channel_binding: md5: %w", err)
		}
	}

	// provide target name.
	attrs[AttrTargetName] = &Value{TargetName: v2.TargetName}

	if v2.TargetName != "" && v2.UnverifiedTargetName {
		attrs[AttrFlags] = &Value{Flag: attrs.Flag() | SPNFromUntrustedSource}
	}

	if value, ok := c.TargetInfo[AttrTimestamp]; ok {
		resp.RequestMIC = true
		attrs[AttrFlags] = &Value{Flag: c.TargetInfo.Flag() | MICProvided}
		respNT.NTLMv2ClientChallenge.Timestamp = value.Timestamp
	} else {
		respNT.NTLMv2ClientChallenge.Timestamp = TimeToFiletime(time.Now())
		if respLM.Response, err = crypto.HMAC_MD5(resp.KeyLM, c.ServerChallenge, nonce); err != nil {
			return nil, fmt.Errorf("v2: compute response lm: %v", err)
		}
		respLM.ChallengeFromClient = nonce
	}

	respNT.NTLMv2ClientChallenge.AttrValues = attrs

	if resp.Tmp, err = respNT.NTLMv2ClientChallenge.Marshal(ctx); err != nil {
		return nil, fmt.Errorf("v2: marshal response nt: %v", err)
	}

	if respNT.Response, err = crypto.HMAC_MD5(resp.KeyNT, c.ServerChallenge, resp.Tmp); err != nil {
		return nil, fmt.Errorf("v2: compute response nt: %v", err)
	}

	if resp.SessionBaseKey, err = crypto.HMAC_MD5(resp.KeyNT, respNT.Response); err != nil {
		return nil, fmt.Errorf("v2: compute session base key: %v", err)
	}

	if resp.LM, err = respLM.Marshal(ctx); err != nil {
		return nil, fmt.Errorf("v2: marshal lm: %v", err)
	}

	if resp.NT, err = respNT.Marshal(ctx); err != nil {
		return nil, fmt.Errorf("v2: marshal nt: %v", err)
	}

	return resp, nil
}

// KeyExchangeKey function returns the key used to protect the session key that is
// generated by the client.
func (v2 *V2) KeyExchangeKey(ctx context.Context, c *ChallengeMessage, chal *ChallengeResponse) ([]byte, error) {
	return chal.SessionBaseKey, nil
}
