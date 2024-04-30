package errors

import (
	"context"
	"fmt"
	"sync"
)

// Mapper represents the error mapping engine.
type Mapper interface {
	// MapValue maps the given value to the registered error.
	MapValue(context.Context, any) error
}

type MapperStore struct {
	mu      sync.Mutex
	mappers []Mapper
}

func (ms *MapperStore) AddMapper(m Mapper) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.mappers = append(ms.mappers, m)
}

func (ms *MapperStore) NewError(ctx context.Context, value any) error {
	for _, mapper := range ms.mappers {
		if err := mapper.MapValue(ctx, value); err != nil {
			return err
		}
	}
	return &Error{Value: value}
}

var (
	errorMapperStore = new(MapperStore)
)

func AddMapper(m Mapper) {
	errorMapperStore.AddMapper(m)
}

func New(ctx context.Context, value any) error {
	return errorMapperStore.NewError(ctx, value)
}

type Error struct {
	Value any
}

func (err *Error) Error() string {
	if value, ok := err.Value.(uint32); ok {
		return fmt.Sprintf("error: code: 0x%08x", value)
	}
	return fmt.Sprintf("error: %v", err.Value)
}
