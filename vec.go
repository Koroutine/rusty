package rusty

type Vec[T comparable] struct {
	data []T
}

func ToVec[T comparable](data []T) *Vec[T] {
	return &Vec[T]{
		data: data,
	}
}

func (o *Vec[T]) Contains(value T) bool {
	for _, v := range o.data {
		if v == value {
			return true
		}
	}

	return false
}
