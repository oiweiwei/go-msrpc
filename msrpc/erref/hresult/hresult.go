//go:generate go run ../gen.go -o hresult_errors.go -pkg hresult -url https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-erref/705fb797-2175-4a90-b5a3-3918024b10b8 -url https://learn.microsoft.com/en-us/previous-versions/windows/internet-explorer/ie-developer/platform-apis/ms775145(v=vs.85) -url https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-csra/f46c3c08-6566-407e-a93e-b0f5e91010f7
package hresult

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
	return fmt.Sprintf("hresult: %s (0x%08x): %s", e.Name, e.Code, e.Details)
}
