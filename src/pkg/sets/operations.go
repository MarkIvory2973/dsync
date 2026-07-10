package sets

import "slices"

func Intersection[T comparable](A []T, B []T) []T {
	C := []T{}
	for _, a := range A {
		if !slices.Contains(B, a) {
			continue
		}

		C = append(C, a)
	}

	return C
}

func SetDifference[T comparable](A []T, B []T) []T {
	C := []T{}
	for _, a := range A {
		if slices.Contains(B, a) {
			continue
		}

		C = append(C, a)
	}

	return C
}
