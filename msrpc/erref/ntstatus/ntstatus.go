//go:generate go run ../gen.go -o ntstatus_errors.go -pkg ntstatus -url https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-erref/596a1078-e883-4972-9bbc-49e60bebca55
package ntstatus

import (
	"context"
	"fmt"

	"github.com/oiweiwei/go-msrpc/dcerpc/errors"
)

type Error struct {
	Code    uint32
	Name    string
	Details string
}

func init() {
	errors.AddMapper(Mapper{})
}

type Mapper struct{}

func (Mapper) MapValue(ctx context.Context, value any) error {
	switch code := value.(type) {
	case uint32:
		return FromCode(code)
	case int32:
		return FromCode(uint32(code))
	}
	return nil
}

func (e *Error) Error() string {
	return fmt.Sprintf("ntstatus: %s (0x%08x): %s", e.Name, e.Code, e.Details)
}
