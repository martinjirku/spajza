package pointer

func Dereference[T comparable](ptr *T) T {
	var result T
	if ptr != nil {
		result = *ptr
	}
	return result
}
