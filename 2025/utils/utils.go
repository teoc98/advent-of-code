package utils

import (
	"fmt"
)

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func Mod(a, b int) int {
	return (a%b + b) % b
}

func Map[T any, U any](in []T, f func(T) U) []U {
	out := make([]U, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}

func MapErr[T any, U any](in []T, f func(T) (U, error)) ([]U, error) {
	out := make([]U, 0, len(in))
	for i, v := range in {
		u, err := f(v)
		if err != nil {
			return nil, fmt.Errorf("index %d: %w", i, err)
		}
		out = append(out, u)
	}
	return out, nil
}
