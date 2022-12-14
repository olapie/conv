package conv

import (
	"fmt"
)

func MustSlice[A any, B any](a []A, f func(A) B) []B {
	b := make([]B, len(a))
	for i := range a {
		b[i] = f(a[i])
	}
	return b
}

func Slice[A any, B any](a []A, f func(A) (B, error)) ([]B, error) {
	b := make([]B, len(a))
	var err error
	for i := range a {
		b[i], err = f(a[i])
		if err != nil {
			return nil, fmt.Errorf("index %d: %w", i, err)
		}
	}
	return b, nil
}

func MustSliceToSet[A any, B comparable](a []A, f func(A) B) map[B]bool {
	m := make(map[B]bool, len(a))
	for _, v := range a {
		if f == nil {
			m[any(a).(B)] = true
		} else {
			m[f(v)] = true
		}
	}
	return m
}

func SliceToSet[A any, B comparable](a []A, f func(A) (B, error)) (map[B]bool, error) {
	m := make(map[B]bool, len(a))
	for i, v := range a {
		if f == nil {
			m[any(a).(B)] = true
		} else {
			b, err := f(v)
			if err != nil {
				return nil, fmt.Errorf("index %d: %w", i, err)
			}
			m[b] = true
		}
	}
	return m, nil
}

func UniqueSlice[E comparable](a []E) []E {
	m := make(map[E]struct{}, len(a))
	l := make([]E, 0, len(a))
	for _, v := range a {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		l = append(l, v)
	}
	return l
}

func ReverseSlice[E comparable](a []E) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func RemoveElement[E comparable](a []E, v E) []E {
	for i, e := range a {
		if e == v {
			a = append(a[:i], a[i+1:]...)
			break
		}
	}
	return a
}
