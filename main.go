package hltrnty

import "sync"

func Map[T any, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for k, v := range s {
		r[k] = f(v)
	}
	return r
}

func ConcurMap[T any, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	wg := sync.WaitGroup{}
	for k, v := range s {
		wg.Add(1)
		go func(i int, v T) {
			defer wg.Done()
			r[i] = f(v)
		}(k, v)
	}
	wg.Wait()
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	r := make([]T, 0)
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Reduce[T any, U any](s []T, f func(acc U, cur T) U, zero U) U {
	for _, v := range s {
		zero = f(zero, v)
	}
	return zero
}

func Some[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}
