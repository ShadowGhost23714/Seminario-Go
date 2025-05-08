package tree

import (
	"sync"
)

type Tree[T any] struct {
	val         T
	left, right *Tree[T]
}

func (t *Tree[T]) Insert(v T, f func(T, T) bool) *Tree[T] {
	switch {
	case t == nil:
		t = &Tree[T]{val: v}
	case f(v, t.val):
		t.left = t.left.Insert(v, f)
	default:
		t.right = t.right.Insert(v, f)
	}
	return t
}

func (t *Tree[T]) GetAll() []T {
	var elems []T
	if t != nil {
		elems = append(elems, t.left.GetAll()...)
		elems = append(elems, t.val)
		elems = append(elems, t.right.GetAll()...)
	}
	return elems
}

func (t *Tree[T]) finder(p []T, out chan<- []T) {
	if t == nil {
		return
	}

	path := make([]T, len(p))
	copy(path, p)
	path = append(path, t.val)

	if t.left == nil || t.right == nil {
		out <- path
	}

	var wgf sync.WaitGroup

	if t.left != nil {
		wgf.Add(1)
		go func() {
			t.left.finder(path, out)
			wgf.Done()
		}()
	}
	if t.right != nil {
		wgf.Add(1)
		go func() {
			t.right.finder(path, out)
			wgf.Done()
		}()
	}
	wgf.Wait()
}

func (t *Tree[T]) AllPaths() [][]T {
	var paths [][]T
	var wg sync.WaitGroup

	ch := make(chan []T)

	wg.Add(2)

	go func(in <-chan []T) {
		for path := range in {
			paths = append(paths, path)
		}
		wg.Done()
	}(ch)

	go func(ch chan []T) {
		t.finder([]T{}, ch)
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
	return paths
}
