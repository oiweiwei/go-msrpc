package logon

import (
	"context"
	"fmt"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
)

func (o *xxx_SecureChannelClient) AccountDeltas(ctx context.Context, in *AccountDeltasRequest, opts ...dcerpc.CallOption) (*AccountDeltasResponse, error) {

	if in == nil {
		in = &AccountDeltasRequest{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	ret, err := o.LogonClient.AccountDeltas(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	return ret, nil
}

func (o *xxx_SecureChannelClient) AccountSync(ctx context.Context, in *AccountSyncRequest, opts ...dcerpc.CallOption) (*AccountSyncResponse, error) {

	if in == nil {
		in = &AccountSyncRequest{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	ret, err := o.LogonClient.AccountSync(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	return ret, nil
}

func (o *xxx_SecureChannelClient) ChainSetClientAttributes(ctx context.Context, in *ChainSetClientAttributesRequest, opts ...dcerpc.CallOption) (*ChainSetClientAttributesResponse, error) {

	if in == nil {
		in = &ChainSetClientAttributesRequest{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	ret, err := o.LogonClient.ChainSetClientAttributes(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	return ret, nil
}

func (o *xxx_SecureChannelClient) DatabaseDeltas(ctx context.Context, in *DatabaseDeltasRequest, opts ...dcerpc.CallOption) (*DatabaseDeltasResponse, error) {

	if in == nil {
		in = &DatabaseDeltasRequest{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	ret, err := o.LogonClient.DatabaseDeltas(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	return ret, nil
}

func (o *xxx_SecureChannelClient) DatabaseRedo(ctx context.Context, in *DatabaseRedoRequest, opts ...dcerpc.CallOption) (*DatabaseRedoResponse, error) {

	if in == nil {
		in = &DatabaseRedoRequest{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	ret, err := o.LogonClient.DatabaseRedo(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	return ret, nil
}

func (o *xxx_SecureChannelClient) DatabaseSync(ctx context.Context, in *DatabaseSyncRequest, opts ...dcerpc.CallOption) (*DatabaseSyncResponse, error) {

	if in == nil {
		in = &DatabaseSyncRequest{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	ret, err := o.LogonClient.DatabaseSync(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	return ret, nil
}

func (o *xxx_SecureChannelClient) DatabaseSync2(ctx context.Context, in *DatabaseSync2Request, opts ...dcerpc.CallOption) (*DatabaseSync2Response, error) {

	if in == nil {
		in = &DatabaseSync2Request{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	ret, err := o.LogonClient.DatabaseSync2(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	return ret, nil
}

func (o *xxx_SecureChannelClient) GetCapabilities(ctx context.Context, in *GetCapabilitiesRequest, opts ...dcerpc.CallOption) (*GetCapabilitiesResponse, error) {

	if in == nil {
		in = &GetCapabilitiesRequest{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	ret, err := o.LogonClient.GetCapabilities(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	return ret, nil
}

func (o *xxx_SecureChannelClient) GetDomainInfo(ctx context.Context, in *GetDomainInfoRequest, opts ...dcerpc.CallOption) (*GetDomainInfoResponse, error) {

	if in == nil {
		in = &GetDomainInfoRequest{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	ret, err := o.LogonClient.GetDomainInfo(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	return ret, nil
}

func (o *xxx_SecureChannelClient) SAMLogoff(ctx context.Context, in *SAMLogoffRequest, opts ...dcerpc.CallOption) (*SAMLogoffResponse, error) {

	if in == nil {
		in = &SAMLogoffRequest{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	ret, err := o.LogonClient.SAMLogoff(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	return ret, nil
}

func (o *xxx_SecureChannelClient) SAMLogon(ctx context.Context, in *SAMLogonRequest, opts ...dcerpc.CallOption) (*SAMLogonResponse, error) {

	if in == nil {
		in = &SAMLogonRequest{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	ret, err := o.LogonClient.SAMLogon(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	return ret, nil
}

func (o *xxx_SecureChannelClient) SAMLogonWithFlags(ctx context.Context, in *SAMLogonWithFlagsRequest, opts ...dcerpc.CallOption) (*SAMLogonWithFlagsResponse, error) {

	if in == nil {
		in = &SAMLogonWithFlagsRequest{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	ret, err := o.LogonClient.SAMLogonWithFlags(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx, nil).OpName(), err)
	}

	return ret, nil
}
