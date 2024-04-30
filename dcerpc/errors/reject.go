package errors

import (
	"context"
	"fmt"
)

func init() {
	// register default mapper.
	AddMapper(RPCMapper{})
}

type RPCMapper struct{}

func (RPCMapper) MapValue(ctx context.Context, value any) error {
	switch code := value.(type) {
	case uint32:
		return FromCode(code)
	case int32:
		return FromCode(uint32(code))
	}
	return nil
}

type RPCError struct {
	Code    uint32
	Name    string
	Details string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("fault: %s (0x%08x): %s", e.Name, e.Code, e.Details)
}
