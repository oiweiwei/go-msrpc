package wmi

import (
	"context"
	"fmt"

	"github.com/oiweiwei/go-msrpc/dcerpc/errors"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi"
)

func init() {
	errors.AddMapper(Mapper{})
}

type Mapper struct{}

func (Mapper) MapValue(ctx context.Context, value any) error {
	switch code := value.(type) {
	case int32:
		return FromCode(code)
	}
	return nil
}

func FromCode(code int32) error {

	st := wmi.Status(code)
	if st.String() == "Invalid" {
		return nil
	}

	return &Error{Code: uint32(code), Name: st.String()}
}

type Error struct {
	Code uint32
	Name string
}

func (e *Error) Error() string {
	return fmt.Sprintf("wmi: %s (0x%08x)", e.Name, e.Code)
}
