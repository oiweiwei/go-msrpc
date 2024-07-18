package drsr

import (
	"context"
	"fmt"

	"github.com/oiweiwei/go-msrpc/dcerpc/errors"
	"github.com/oiweiwei/go-msrpc/msrpc/drsr/drsuapi/v4"
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

	st := drsuapi.DSNameError(code)
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
	return fmt.Sprintf("drsuapi: %s (0x%08x)", e.Name, e.Code)
}
