//go:generate go run ../gen.go -o win32_errors.go -pkg win32 -url https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-erref/18d8fbe8-a967-4f1c-ae50-99ca8e491d2d
package win32

import (
	"context"
	"fmt"

	"github.com/oiweiwei/go-msrpc/dcerpc/errors"
)

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

type Error struct {
	Code    uint32
	Name    string
	Details string
}

func (e *Error) Error() string {
	return fmt.Sprintf("win32: %s (0x%08x): %s", e.Name, e.Code, e.Details)
}
