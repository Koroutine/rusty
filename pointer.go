package rusty

func ToPointer[T comparable](data T) *T {
	return &data
}
