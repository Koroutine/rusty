package rusty

type Either[L any, R any] struct {
	left  L
	right R
}

func ToEither[L any, R any](left L, right R) *Either[L, R] {
	return &Either[L, R]{
		left, right,
	}
}

func (e *Either[L, R]) Left() L {
	return e.left
}

func (e *Either[L, R]) Right() R {
	return e.right
}
