package dcom

import (
	"context"
)

func (o *ErrorInfoString) AfterPreparePayload(ctx context.Context) error {

	if o.Actual == 0 {
		o.Actual = o.Max
	}

	return nil
}
