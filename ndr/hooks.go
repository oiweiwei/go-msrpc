package ndr

import "context"

// BeforePreparePayloadHook is an interface that defines a hook to be called before
// preparing the payload.
type BeforePreparePayloadHook interface {
	BeforePreparePayload(context.Context) error
}

// AfterPreparePayloadHook is an interface that defines a hook to be called after
// preparing the payload.
type AfterPreparePayloadHook interface {
	AfterPreparePayload(context.Context) error
}

// BeforePreparePayload calls the BeforePreparePayload method on the Marshaler if it implements
// the BeforePreparePayloadHook interface.
func BeforePreparePayload(ctx context.Context, o Marshaler) error {
	if hook, ok := o.(BeforePreparePayloadHook); ok {
		return hook.BeforePreparePayload(ctx)
	}
	return nil
}

// AfterPreparePayload calls the AfterPreparePayload method on the Marshaler if it implements
// the AfterPreparePayloadHook interface.
func AfterPreparePayload(ctx context.Context, o Marshaler) error {
	if hook, ok := o.(AfterPreparePayloadHook); ok {
		return hook.AfterPreparePayload(ctx)
	}
	return nil
}
