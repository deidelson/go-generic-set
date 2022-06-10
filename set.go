package set

type SetEntry[V comparable] interface {
	GetKey() V
}

type HashSet[V comparable, T SetEntry[V]] struct {
	entries map[V]T
}

func NewHashSet[V comparable, T SetEntry[V]]() *HashSet[V, T] {
	return &HashSet[V, T]{
		entries: make(map[V]T),
	}
}

func (s *HashSet[V, T]) Add(entry T) {
	s.entries[entry.GetKey()] = entry
}

func (s *HashSet[V, T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *HashSet[V, T]) Size() int {
	return len(s.entries)
}

func (s *HashSet[V, T]) Contains(value T) bool {
	_, existsKey := s.entries[value.GetKey()]
	return existsKey
}

func (s *HashSet[V, T]) Remove(value T) {
	delete(s.entries, value.GetKey())
}

func (s *HashSet[V, T]) Clear() {
	for key := range s.entries {
		delete(s.entries, key)
	}
}
