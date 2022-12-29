package rusty

import "fmt"

type Result[T any] struct {
	data T
	err  error
}

func ToResult[T any](data T, err error) *Result[T] {
	return &Result[T]{
		data: data,
		err:  err,
	}
}

func ToOk[T any](data T) *Result[T] {
	return &Result[T]{
		data: data,
		err:  nil,
	}
}

func ToError(err error) *Result[bool] {
	return &Result[bool]{
		err:  err,
		data: err == nil,
	}
}

func (o *Result[T]) Unwrap() T {
	if o.err != nil {
		panic(o.err)
	}

	return o.data
}

func (o *Result[T]) UnwrapSafe(fn func(err error)) T {
	if o.err != nil {
		fn(o.err)
	}

	return o.data
}

func (o *Result[T]) Expect(message string) T {
	if o.err != nil {
		panic(fmt.Errorf("%s: %w", message, o.err))
	}

	return o.data
}

func (o *Result[T]) UnwrapOr(value T) T {
	if o.err != nil {
		return value
	}

	return o.data
}

func (o *Result[T]) UnwrapOrElse(fn func() T) T {
	if o.err != nil {
		return fn()
	}

	return o.data
}

func (o *Result[T]) UnwrapOrElseE(fn func(err error) T) T {
	if o.err != nil {
		return fn(o.err)
	}

	return o.data
}

func (o *Result[T]) MapErr(fn func(err error) error) *Result[T] {
	if o.err != nil {
		o.err = fn(o.err)
	}

	return o
}

func (o *Result[T]) Map(fn func(value T) T) *Result[T] {
	if o.err == nil {
		o.data = fn(o.data)
	}

	return o
}

func (o *Result[T]) IsOk() bool {
	if o.err != nil {
		return false
	}

	return true
}

func (o *Result[T]) Ok() *Option[T] {
	if o.err != nil {
		return ToNone[T]()
	}

	return ToOption(&o.data)
}

func (o *Result[T]) IsOkWith(fn func(value T) bool) bool {
	if o.err != nil {
		return false
	}

	return fn(o.data)
}

func (o *Result[T]) IsErr() bool {
	return !o.IsOk()
}

func (o *Result[T]) IsErrWith(fn func(err error) bool) bool {
	if o.err == nil {
		return false
	}

	return fn(o.err)
}

func MapResult[T any, N any](result *Result[T], fn func(value T) N) *Result[N] {

	if result.IsOk() {
		data := fn(result.Unwrap())
		return ToResult(data, nil)
	}

	return &Result[N]{
		err: result.err,
	}
}
