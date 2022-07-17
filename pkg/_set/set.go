package set

import (
	"sync"
)

// Set is a structure that can contain a single item of any type compataible type.
// It allows fast insert and fast lookup.
type Set[T comparable] struct {
	once    sync.Once
	mut     sync.RWMutex
	content map[T]struct{}
}

func (s *Set[T]) With(items ...T) *Set[T] {
	s.Add(items...)
	return s
}

// Add adds an item to the set if it does not exist
func (s *Set[T]) Add(items ...T) {
	s.once.Do(func() {
		if s == nil {
			s = &Set[T]{}
		}
		s.content = make(map[T]struct{})
	})
	s.mut.Lock()
	defer s.mut.Unlock()
	for _, item := range items {
		s.content[item] = struct{}{}
	}
}

// Contains return true if the item is contained in the set.
func (s *Set[T]) Contains(items ...T) bool {
	if s == nil {
		return false
	}
	s.mut.RLock()
	defer s.mut.RUnlock()
	for _, item := range items {
		if _, ok := s.content[item]; !ok {
			return false
		}
	}
	return true
}

// Remove removes an item from the set.
// If the item is not present, it's a no-op.
func (s *Set[T]) Remove(items ...T) {
	if s == nil {
		return
	}
	s.mut.Lock()
	defer s.mut.Unlock()
	for _, item := range items {
		delete(s.content, item)
	}
}

// Values returns an array of the values contained in the set
func (s *Set[T]) Values() []T {
	var values = make([]T, len(s.content))
	s.mut.Lock()
	defer s.mut.Unlock()
	for k := range s.content {
		values = append(values, k)
	}
	return values
}
