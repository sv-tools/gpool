package gpool

import "sync"

// Pool is generic drop-in replacement of `sync.Pool`
type Pool[T any] struct {
	sp sync.Pool

	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New func() T
}

// Put adds x to the pool.
func (p *Pool[T]) Put(x T) {
	p.sp.Put(x)
}

// Get selects an arbitrary item from the Pool, removes it from the
// Pool, and returns it to the caller.
// Get may choose to ignore the pool and treat it as empty.
// Callers should not assume any relation between values passed to Put and
// the values returned by Get.
//
// If Get would otherwise return nil and p.New is non-nil, Get returns
// the result of calling p.New.
func (p *Pool[T]) Get() (item T) {
	x := p.sp.Get()
	if x == nil {
		if p.New != nil {
			return p.New()
		}
		return item
	}
	return x.(T)
}
