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

func SliceAppend[T any](slice []T, data ...T) []T {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]T, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
