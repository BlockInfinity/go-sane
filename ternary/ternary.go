package ternary

func If[T any](condition bool, vTrue T, vFalse T) T {
	if condition {
		return vTrue
	}

	return vFalse
}
