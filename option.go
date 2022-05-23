package rusty

import (
	"errors"
)

type Option[T any] struct {
	data *T
}

func ToOption[T any](data *T) *Option[T] {
	return &Option[T]{
		data: data,
	}
}

func ToNone[T any]() *Option[T] {
	return &Option[T]{}
}

func (o *Option[T]) Expect(message string) T {
	if o.data == nil {
		panic(errors.New(message))
	}

	return *o.data
}

func (o *Option[T]) Unwrap() T {
	if o.data == nil {
		panic("Option is nil")
	}

	return *o.data
}

func (o *Option[T]) UnwrapOr(value T) T {
	if o.data == nil {
		return value
	}

	return *o.data
}

func (o *Option[T]) UnwrapOrElse(fn func() T) T {
	if o.data == nil {
		return fn()
	}

	return *o.data
}

func (o *Option[T]) IsSome() bool {
	if o.data == nil {
		return false
	}

	return true
}

func (o *Option[T]) IsSomeWith(fn func(value T) bool) bool {
	if o.data == nil {
		return false
	}

	return fn(*o.data)
}

func (o *Option[T]) IsNone() bool {
	return !o.IsSome()
}
