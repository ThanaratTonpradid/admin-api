package helper

import "golang.org/x/exp/constraints"

func ToPtr[T any](n T) *T {
	return &n
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func BoolToInt(value bool) int8 {
	if value {
		return 1
	}
	return 0
}
