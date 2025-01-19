package commonutils

func SliceContains[T comparable](slice []T, element T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func IsSliceEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func MapSlice[T interface{}, K interface{}](s []T, f func(T) K) []K {
	mapped := make([]K, len(s))

	for i, e := range s {
		mapped[i] = f(e)
	}

	return mapped
}
